// Package engine drives the PostgreSQL wasm modules: it prepares the virtual
// filesystem, runs initdb, and hosts a single-user backend that speaks the
// frontend/backend protocol through in-memory buffers.
package engine

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/shibukawa/pgmem/internal/assets"
	"github.com/shibukawa/pgmem/internal/guest"
	"github.com/shibukawa/pgmem/internal/host"
	"github.com/shibukawa/pgmem/internal/vfs"
)

// Guest filesystem layout (mirrors PGlite so PostgreSQL's relative-path
// discovery works: bin/ -> share/postgresql, lib/postgresql).
const (
	PGRoot   = "/pglite"
	PGData   = PGRoot + "/data"
	BinDir   = PGRoot + "/bin"
	ShareDir = PGRoot + "/share/postgresql"
	LibDir   = PGRoot + "/lib/postgresql"
	PGMemDir = "/pgmem"

	exitAlive   = 99  // PGLITE_EXIT_ALIVE
	exitLongjmp = 100 // POSTGRES_MAIN_LONGJMP
)

// Engine owns the backend factory.
type Engine struct {
	ctx      context.Context
	postgres guest.Factory
	modules  []string // statically linked loadable modules
	// Log receives host diagnostics (nil = silent).
	Log func(format string, args ...any)
}

// PostgresFactory must be set before New; the root package sets it to the
// ahead-of-time compiled backend.
var PostgresFactory guest.Factory

// New prepares the backend.
func New(ctx context.Context) (*Engine, error) {
	if PostgresFactory == nil {
		return nil, fmt.Errorf("engine: no backend registered")
	}
	e := &Engine{ctx: ctx, postgres: PostgresFactory}
	if err := e.loadModuleNames(); err != nil {
		return nil, err
	}
	return e, nil
}

// Close releases everything.
func (e *Engine) Close() error { return nil }

func (e *Engine) logf(format string, args ...any) {
	if e.Log != nil {
		e.Log(format, args...)
	}
}

// loadModuleNames asks the backend module which loadable modules were
// linked in, so placeholder .so files can be created for dfmgr's stat().
func (e *Engine) loadModuleNames() error {
	h := host.New(vfs.New())
	mod, err := e.postgres.Instantiate(e.ctx, h)
	if err != nil {
		return err
	}
	defer mod.Close(e.ctx)
	for i := int32(0); ; i++ {
		p, err := mod.CallI32("pgmem_module_name", uint64(uint32(i)))
		if err != nil || p == 0 {
			break
		}
		e.modules = append(e.modules, readCString(mod, uint32(p)))
	}
	return nil
}

func readCString(m guest.Instance, p uint32) string {
	var sb strings.Builder
	for {
		b, ok := m.Memory().Read(p, 1)
		if !ok || b[0] == 0 {
			return sb.String()
		}
		sb.WriteByte(b[0])
		p++
	}
}

// BaseFS creates a filesystem with the share tree and executable/library
// placeholders, but no data directory.
func (e *Engine) BaseFS() (*vfs.FS, error) {
	fs := vfs.New()
	if err := untarGz(fs, assets.ShareTarGz, ShareDir); err != nil {
		return nil, fmt.Errorf("unpack share: %w", err)
	}
	for _, d := range []string{BinDir, LibDir, PGMemDir, "/home/postgres", "/tmp"} {
		if _, err := fs.MkdirAll(d, 0o755); err != vfs.OK {
			return nil, err
		}
	}
	for _, b := range []string{"postgres", "initdb", "pg_dump"} {
		fs.WriteFile(BinDir+"/"+b, nil, 0o755)
	}
	for _, m := range e.modules {
		fs.WriteFile(LibDir+"/"+m+".so", nil, 0o644)
	}
	return fs, nil
}

func untarGz(fs *vfs.FS, data []byte, root string) error {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return err
	}
	return Untar(fs, gz, root)
}

// Untar unpacks a tar stream under root.
func Untar(fs *vfs.FS, r io.Reader, root string) error {
	tr := tar.NewReader(r)
	if _, err := fs.MkdirAll(root, 0o755); err != vfs.OK {
		return err
	}
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		name := strings.TrimPrefix(strings.TrimPrefix(hdr.Name, "./"), "/")
		if name == "" || name == "." {
			continue
		}
		p := root + "/" + name
		switch hdr.Typeflag {
		case tar.TypeDir:
			if _, err := fs.MkdirAll(p, uint32(hdr.Mode)); err != vfs.OK {
				return fmt.Errorf("mkdir %s: %w", p, err)
			}
		case tar.TypeReg:
			// Exact-size buffer handed to the VFS as is: one copy, no slack.
			b := make([]byte, hdr.Size)
			if _, err := io.ReadFull(tr, b); err != nil {
				return err
			}
			if err := fs.PutFile(p, b, uint32(hdr.Mode)); err != vfs.OK {
				return fmt.Errorf("write %s: %w", p, err)
			}
		case tar.TypeLink:
			// A hard link (zic makes the timezone aliases this way):
			// store a copy of the target already unpacked.
			target := root + "/" + strings.TrimPrefix(strings.TrimPrefix(hdr.Linkname, "./"), "/")
			b, errno := fs.ReadFile(target)
			if errno != vfs.OK {
				return fmt.Errorf("link %s -> %s: %w", p, target, errno)
			}
			if err := fs.PutFile(p, append([]byte(nil), b...), uint32(hdr.Mode)); err != vfs.OK {
				return fmt.Errorf("write %s: %w", p, err)
			}
		case tar.TypeSymlink:
			fs.Symlink(hdr.Linkname, p)
		}
	}
}

// Tar serializes a subtree of fs (paths relative to root) to w.
func Tar(fs *vfs.FS, root string, w io.Writer) error {
	tw := tar.NewWriter(w)
	err := fs.Walk(root, func(p string, st vfs.Stat, data []byte, target string) error {
		rel := strings.TrimPrefix(strings.TrimPrefix(p, root), "/")
		if rel == "" {
			return nil
		}
		hdr := &tar.Header{Name: rel, Mode: int64(st.Mode & 0o777), ModTime: st.Mtime}
		switch st.Mode & vfs.S_IFMT {
		case vfs.S_IFDIR:
			hdr.Typeflag = tar.TypeDir
			hdr.Name += "/"
		case vfs.S_IFLNK:
			hdr.Typeflag = tar.TypeSymlink
			hdr.Linkname = target
		case vfs.S_IFREG:
			hdr.Typeflag = tar.TypeReg
			hdr.Size = int64(len(data))
		default:
			return nil
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if hdr.Typeflag == tar.TypeReg {
			if _, err := tw.Write(data); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return tw.Close()
}

func guestEnv(user, database string) []string {
	return []string{
		"PGDATA=" + PGData,
		"HOME=/home/postgres",
		"USER=postgres",
		"LOGNAME=postgres",
		"PGUSER=" + user,
		"PGDATABASE=" + database,
		"LANG=C.UTF-8",
		"LC_ALL=C.UTF-8",
		"LC_COLLATE=C.UTF-8",
		"LC_CTYPE=C.UTF-8",
		"TZ=UTC",
		"PGTZ=UTC",
		"PGCLIENTENCODING=UTF8",
		"PATH=" + BinDir,
	}
}

// callMain runs pgmem_main(argv) and returns the process exit code. A
// "live" exit (emscripten_exit_with_live_runtime) is reported as
// (code, true, nil).
func callMain(mod guest.Instance, argv []string) (code int32, live bool, err error) {
	var sb strings.Builder
	for _, a := range argv {
		sb.WriteString(a)
		sb.WriteByte(0)
	}
	buf := sb.String()
	p, err := mod.Malloc(uint32(len(buf)))
	if err != nil {
		return 0, false, err
	}
	mod.Memory().Write(p, []byte(buf))
	res, err := mod.Call("pgmem_main", uint64(p), uint64(len(buf)))
	if err != nil {
		var ee *host.ExitError
		if errors.As(err, &ee) {
			if ee.Live {
				return 0, true, nil
			}
			return ee.Code, false, nil
		}
		return 0, false, err
	}
	if len(res) > 0 {
		return int32(uint32(res[0])), false, nil
	}
	return 0, false, nil
}

// InitdbOptions configures Initdb.
type InitdbOptions struct {
	// Initdb provides the initdb program (the wasm module run under wazero;
	// only cmd/pgmem-mkdata needs it, so it is not part of the library).
	Initdb guest.Factory
	User   string   // superuser name (default postgres)
	Args   []string // extra initdb arguments
}

// Initdb runs initdb into PGData on fs. stderr output from initdb and its
// postgres children is returned with any error.
func (e *Engine) Initdb(fs *vfs.FS, opts InitdbOptions) error {
	if opts.User == "" {
		opts.User = "postgres"
	}
	var stderr, stdout bytes.Buffer
	pfs := fs.Fork()
	pfs.Stderr = func(b []byte) { stderr.Write(b) }
	pfs.Stdout = func(b []byte) { stdout.Write(b) }
	h := host.New(pfs)
	h.Env = guestEnv(opts.User, "postgres")
	h.Log = e.Log
	h.Run = func(cmd, in, out string) int {
		return e.runChild(fs, h.Env, cmd, in, out, &stderr)
	}
	if opts.Initdb == nil {
		return fmt.Errorf("initdb: no initdb module provided")
	}
	mod, err := opts.Initdb.Instantiate(e.ctx, h)
	if err != nil {
		return err
	}
	defer mod.Close(e.ctx)
	if _, err := mod.Call("pgmem_init"); err != nil {
		return err
	}
	argv := append([]string{
		BinDir + "/initdb",
		"--allow-group-access",
		"--encoding", "UTF8",
		"--locale=C.UTF-8",
		"--locale-provider=libc",
		"--auth=trust",
		"--no-sync",
		"--no-instructions",
		"-U", opts.User,
		"-D", PGData,
	}, opts.Args...)
	code, _, err := callMain(mod, argv)
	if err != nil {
		return fmt.Errorf("initdb: %w\n%s%s", err, stdout.String(), stderr.String())
	}
	if code != 0 {
		return fmt.Errorf("initdb exited with %d\n%s%s", code, stdout.String(), stderr.String())
	}
	e.logf("initdb ok\n%s", stdout.String())
	return nil
}

// runChild executes "postgres ..." on behalf of initdb (popen/system).
func (e *Engine) runChild(fs *vfs.FS, env []string, cmd, stdinPath, stdoutPath string, stderr *bytes.Buffer) int {
	argv, in, out := parseCommand(cmd)
	if stdinPath != "" {
		in = stdinPath
	}
	if stdoutPath != "" {
		out = stdoutPath
	}
	if len(argv) == 0 {
		return 127
	}
	e.logf("child: %v stdin=%q stdout=%q", argv, in, out)
	cfs := fs.Fork()
	cfs.Stderr = func(b []byte) { stderr.Write(b) }
	cfs.Stdout = func(b []byte) { stderr.Write(b) }
	h := host.New(cfs)
	h.Env = env
	h.Log = e.Log
	mod, err := e.postgres.Instantiate(e.ctx, h)
	if err != nil {
		e.logf("child instantiate: %v", err)
		return 126
	}
	defer mod.Close(e.ctx)
	if _, err := mod.Call("pgmem_init"); err != nil {
		return 126
	}
	if in != "" {
		if err := freopen(mod, in, "r", 0); err != nil {
			e.logf("child freopen stdin: %v", err)
			return 126
		}
	}
	if out != "" {
		if err := freopen(mod, out, "w", 1); err != nil {
			e.logf("child freopen stdout: %v", err)
			return 126
		}
	}
	code, live, err := callMain(mod, argv)
	if err != nil {
		e.logf("child %v: %v", argv, err)
		fmt.Fprintf(stderr, "pgmem: child %v failed: %v\n", argv, err)
		return 125
	}
	if live {
		return 0
	}
	return int(code)
}

func freopen(mod guest.Instance, path, mode string, stream int32) error {
	p, err := mod.WriteCString(path)
	if err != nil {
		return err
	}
	m, err := mod.WriteCString(mode)
	if err != nil {
		return err
	}
	r, err := mod.CallI32("pgl_freopen", uint64(p), uint64(m), uint64(uint32(stream)))
	mod.Free(p)
	mod.Free(m)
	if err != nil {
		return err
	}
	if r == 0 {
		return fmt.Errorf("freopen(%s) failed", path)
	}
	return nil
}

// parseCommand splits a shell-ish command line into argv and stdin/stdout
// redirections. Quotes are honoured; "2>&1" is ignored.
func parseCommand(cmd string) (argv []string, stdin, stdout string) {
	var toks []string
	var cur strings.Builder
	inTok := false
	var quote byte
	for i := 0; i < len(cmd); i++ {
		c := cmd[i]
		switch {
		case quote != 0:
			if c == quote {
				quote = 0
			} else {
				cur.WriteByte(c)
			}
		case c == '"' || c == '\'':
			quote = c
			inTok = true
		case c == ' ' || c == '\t' || c == '\n':
			if inTok {
				toks = append(toks, cur.String())
				cur.Reset()
				inTok = false
			}
		default:
			cur.WriteByte(c)
			inTok = true
		}
	}
	if inTok {
		toks = append(toks, cur.String())
	}
	for i := 0; i < len(toks); i++ {
		t := toks[i]
		switch {
		case t == "<" && i+1 < len(toks):
			stdin = toks[i+1]
			i++
		case t == ">" && i+1 < len(toks):
			stdout = toks[i+1]
			i++
		case strings.HasPrefix(t, "<"):
			stdin = t[1:]
		case strings.HasPrefix(t, "2>"):
			// ignore stderr redirection
		case strings.HasPrefix(t, ">"):
			stdout = t[1:]
		default:
			argv = append(argv, t)
		}
	}
	return
}

// ExecStandalone runs SQL through a throw-away "postgres --single" child
// (the same way initdb creates the postgres database). It is the way to
// run commands that a live session cannot, such as CREATE DATABASE.
func (e *Engine) ExecStandalone(fs *vfs.FS, database, sql string) error {
	const in, out = PGMemDir + "/standalone.in", PGMemDir + "/standalone.out"
	if err := fs.WriteFile(in, []byte(sql+"\n"), 0o600); err != vfs.OK {
		return err
	}
	var stderr bytes.Buffer
	argv := []string{BinDir + "/postgres", "--single", "-F", "-O", "-j", "-c", "search_path=pg_catalog", "-c", "exit_on_error=true", "-c", "log_checkpoints=false", "-c", "shared_buffers=16MB", database}
	code := e.runChild(fs, guestEnv("postgres", database), strings.Join(argv, " "), in, out, &stderr)
	fs.RemoveAll(in)
	fs.RemoveAll(out)
	// The child's linear memory is garbage now; return it to the OS before
	// the caller starts the next instance, so peak RSS stays at one copy.
	debug.FreeOSMemory()
	if code != 0 {
		return fmt.Errorf("standalone postgres exited with %d\n%s", code, stderr.String())
	}
	return nil
}

// StartOptions configures a backend session.
type StartOptions struct {
	User     string
	Database string
	// Params are extra "postgres" command line arguments (e.g. "-c", "x=y").
	Params []string
}

// Backend is a running single-user PostgreSQL session.
type Backend struct {
	e    *Engine
	fs   *vfs.FS
	h    *host.Host
	mod  guest.Instance
	mu   sync.Mutex
	in   []byte
	off  int
	out  bytes.Buffer
	log  bytes.Buffer
	dead bool
	// Stderr receives server log lines (nil = keep in Log()).
	Stderr func([]byte)
	// Listen receives the session's LISTEN set changes at commit time
	// (see host.Host.Listen). It runs inside Exec.
	Listen func(channel string, op int)
	// More is called, inside Exec, when the backend wants input beyond
	// the batch Exec was given: a COPY FROM STDIN whose data spans
	// batches. It returns the next messages from the client, or nil when
	// there are none (the backend then sees end of stream).
	More func() []byte
}

// Start boots a backend on a filesystem that already has a data directory.
func (e *Engine) Start(fs *vfs.FS, opts StartOptions) (*Backend, error) {
	if opts.User == "" {
		opts.User = "postgres"
	}
	if opts.Database == "" {
		opts.Database = "postgres"
	}
	b := &Backend{e: e, fs: fs}
	pfs := fs.Fork()
	pfs.Stderr = func(p []byte) {
		if b.Stderr != nil {
			b.Stderr(p)
		} else {
			b.log.Write(p)
		}
	}
	pfs.Stdout = pfs.Stderr
	h := host.New(pfs)
	h.Env = guestEnv(opts.User, opts.Database)
	h.Log = e.Log
	h.Recv = b.recv
	h.Send = b.send
	h.Listen = func(channel string, op int) {
		if b.Listen != nil {
			b.Listen(channel, op)
		}
	}
	h.Run = func(cmd, _, _ string) int { e.logf("backend tried to run %q", cmd); return 1 }
	b.h = h
	mod, err := e.postgres.Instantiate(e.ctx, h)
	if err != nil {
		return nil, err
	}
	b.mod = mod
	if _, err := mod.Call("pgmem_init"); err != nil {
		mod.Close(e.ctx)
		return nil, err
	}
	if _, err := mod.Call("pgl_setPGliteActive", 1); err != nil {
		mod.Close(e.ctx)
		return nil, err
	}
	argv := append([]string{
		BinDir + "/postgres",
		"--single",
		"-F",
		"-O",
		"-j",
		"-c", "search_path=public",
	}, opts.Params...)
	argv = append(argv, "-D", PGData, opts.Database)
	code, live, err := callMain(mod, argv)
	if err != nil {
		mod.Close(e.ctx)
		return nil, fmt.Errorf("postgres: %w\n%s", err, b.log.String())
	}
	status, _ := mod.CallI32("pgl_setPGliteExitStatus", uint64(uint32(0xFFFFFFFD))) // -3
	if !live || status != exitAlive {
		mod.Close(e.ctx)
		return nil, fmt.Errorf("postgres failed to start (exit=%d live=%v status=%d)\n%s", code, live, status, b.log.String())
	}
	if _, err := mod.Call("pgl_startPGlite"); err != nil {
		mod.Close(e.ctx)
		return nil, fmt.Errorf("pgl_startPGlite: %w\n%s", err, b.log.String())
	}
	return b, nil
}

func (b *Backend) recv(buf []byte) int {
	if b.off >= len(b.in) && b.More != nil {
		if more := b.More(); len(more) > 0 {
			b.in, b.off = more, 0
		}
	}
	n := copy(buf, b.in[b.off:])
	b.off += n
	return n
}

func (b *Backend) send(p []byte) int {
	b.out.Write(p)
	return len(p)
}

func (b *Backend) takeOutput() []byte {
	out := append([]byte(nil), b.out.Bytes()...)
	b.out.Reset()
	return out
}

// LinearMemorySize returns the current size of the backend's wasm memory.
func (b *Backend) LinearMemorySize() uint64 { return uint64(b.mod.Memory().Size()) }

// Log returns and clears accumulated server log output.
func (b *Backend) Log() string {
	s := b.log.String()
	b.log.Reset()
	return s
}

// Startup processes a StartupMessage and returns the backend's response
// (AuthenticationOk, ParameterStatus..., BackendKeyData, ReadyForQuery).
func (b *Backend) Startup(pkt []byte) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.dead {
		return nil, errors.New("backend is closed")
	}
	b.in, b.off = pkt, 0
	b.out.Reset()
	port, err := b.mod.CallI32("pgl_getMyProcPort")
	if err != nil {
		return nil, err
	}
	r, err := b.mod.CallI32("ProcessStartupPacket", uint64(uint32(port)), 1, 1)
	if err != nil {
		return nil, fmt.Errorf("ProcessStartupPacket: %w\n%s", err, b.log.String())
	}
	if r != 0 {
		return nil, fmt.Errorf("ProcessStartupPacket returned %d\n%s", r, b.log.String())
	}
	if _, err := b.mod.Call("pgl_sendConnData"); err != nil {
		return nil, fmt.Errorf("pgl_sendConnData: %w\n%s", err, b.log.String())
	}
	if _, err := b.mod.Call("pgl_pq_flush"); err != nil {
		return nil, err
	}
	return b.takeOutput(), nil
}

// Exec feeds one or more complete frontend messages to the backend and
// returns everything it wrote back.
func (b *Backend) Exec(msg []byte) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.dead {
		return nil, errors.New("backend is closed")
	}
	b.in, b.off = msg, 0
	b.out.Reset()
	var loopErr error
	for b.off < len(b.in) || b.remaining() > 0 {
		if err := b.h.CheckTimers(); err != nil {
			loopErr = err
			break
		}
		_, err := b.mod.Call("PostgresMainLoopOnce")
		if err == nil {
			continue
		}
		var ee *host.ExitError
		if !errors.As(err, &ee) {
			loopErr = err
			break
		}
		status, _ := b.mod.CallI32("pgl_setPGliteExitStatus", uint64(uint32(0xFFFFFFFE))) // -2
		switch {
		case status == exitLongjmp:
			// ereport(ERROR): run the recovery block and keep going.
			if _, err := b.mod.Call("PostgresMainLongJmp"); err != nil {
				loopErr = err
			}
		case status == exitAlive:
			// Client sent Terminate.
			b.dead = true
		default:
			b.dead = true
			loopErr = fmt.Errorf("backend exited (code=%d status=%d)\n%s", ee.Code, status, b.log.String())
		}
		if loopErr != nil || b.dead {
			break
		}
	}
	if !b.dead {
		if _, err := b.mod.Call("PostgresSendReadyForQueryIfNecessary"); err != nil && loopErr == nil {
			loopErr = err
		}
		if _, err := b.mod.Call("pgl_pq_flush"); err != nil && loopErr == nil {
			loopErr = err
		}
	}
	out := b.takeOutput()
	if loopErr != nil {
		return out, loopErr
	}
	return out, nil
}

// ResetSession aborts any open transaction and, with discard, gives the
// session the state a new client expects (what DISCARD ALL leaves, and
// no temp namespace) without running a statement, so pg_stat_statements
// and the log do not see pgmem's own housekeeping.
func (b *Backend) ResetSession(discard bool) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.dead {
		return errors.New("backend is closed")
	}
	b.in, b.off = nil, 0
	b.out.Reset()
	flag := uint64(0)
	if discard {
		flag = 1
	}
	_, err := b.mod.Call("pgmem_reset_session", flag)
	if err == nil {
		return nil
	}
	var ee *host.ExitError
	if !errors.As(err, &ee) {
		return err
	}
	status, _ := b.mod.CallI32("pgl_setPGliteExitStatus", uint64(uint32(0xFFFFFFFE)))
	if status != exitLongjmp {
		b.dead = true
		return fmt.Errorf("backend exited during session reset (code=%d status=%d)\n%s", ee.Code, status, b.log.String())
	}
	// ereport(ERROR) inside the reset: run the recovery block; the
	// session is then at least out of any transaction.
	if _, err := b.mod.Call("PostgresMainLongJmp"); err != nil {
		return err
	}
	return fmt.Errorf("session reset failed: %s", strings.TrimSpace(b.log.String()))
}

func (b *Backend) remaining() int32 {
	n, err := b.mod.CallI32("pq_buffer_remaining_data")
	if err != nil {
		return 0
	}
	return n
}

// Close shuts the session down and frees the instance.
func (b *Backend) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.mod == nil {
		return nil
	}
	if !b.dead {
		b.mod.Call("pgl_setPGliteActive", 0)
		b.in, b.off = []byte{'X', 0, 0, 0, 4}, 0
		b.mod.Call("PostgresMainLoopOnce") // exits via proc_exit(0)
		b.mod.Call("pgl_run_atexit_funcs")
		b.dead = true
	}
	err := b.mod.Close(b.e.ctx)
	b.mod = nil
	return err
}
