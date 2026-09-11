// Package pgmem runs a real PostgreSQL server entirely in memory, inside the
// test process, and exposes it over the normal wire protocol on a loopback
// TCP port. Any PostgreSQL driver can connect to it.
//
// The server is PostgreSQL (the PGlite fork) compiled to WebAssembly and then
// translated to Go ahead of time, running on an in-memory filesystem. It is a
// single-user,
// single-session backend: every TCP connection shares one backend session,
// so run tests against it with at most one connection at a time (for
// pgxpool set MaxConns to 1, for database/sql call SetMaxOpenConns(1)).
package pgmem

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/klauspost/compress/zstd"

	"github.com/shibukawa/pgmem/internal/engine"
	"github.com/shibukawa/pgmem/internal/pgdata"
	"github.com/shibukawa/pgmem/internal/vfs"
)

// The compiled wasm module is shared by every Server in the process; each
// Server still gets its own filesystem and backend instance.
var (
	engineOnce sync.Once
	sharedEng  *engine.Engine
	engineErr  error
)

func sharedEngine() (*engine.Engine, error) {
	engineOnce.Do(func() {
		sharedEng, engineErr = engine.New(context.Background())
	})
	return sharedEng, engineErr
}

// Options configures Start.
type Options struct {
	// Database and User for the DSN. The user is the initdb superuser.
	Database string
	User     string
	// Port to listen on (0 = pick a free port). Always binds 127.0.0.1.
	Port int
	// Params are extra "postgres" command-line arguments, e.g.
	// []string{"-c", "log_statement=all"}. Unless it sets shared_buffers,
	// pgmem uses 32MB instead of initdb's 128MB: a test database does not
	// need a large buffer cache and every megabyte here is resident memory.
	Params []string
	// Log receives server log output and host diagnostics (nil = discard).
	Log func(format string, args ...any)
}

// Server is a running in-memory PostgreSQL.
type Server struct {
	opts Options
	e    *engine.Engine
	fs   *vfs.FS
	b    *engine.Backend
	ln   net.Listener
	port int

	mu          sync.Mutex // serializes sessions on the single backend
	startupResp []byte
	connSeq     atomic.Int64
	closed      atomic.Bool
	wg          sync.WaitGroup
	connsMu     sync.Mutex
	conns       map[net.Conn]struct{}
}

// Start boots a fresh server: compiles the module, runs initdb into memory,
// starts the backend and listens on 127.0.0.1.
func Start(ctx context.Context, opts Options) (*Server, error) {
	if opts.Database == "" {
		opts.Database = "postgres"
	}
	if opts.User == "" {
		opts.User = "postgres"
	}
	if !hasSetting(opts.Params, "shared_buffers") {
		opts.Params = append([]string{"-c", "shared_buffers=32MB"}, opts.Params...)
	}
	e, err := sharedEngine()
	if err != nil {
		return nil, err
	}
	s := &Server{opts: opts, e: e, conns: map[net.Conn]struct{}{}}
	fs, err := e.BaseFS()
	if err != nil {
		return nil, err
	}
	zr, err := zstd.NewReader(bytes.NewReader(pgdata.TarZst))
	if err != nil {
		return nil, err
	}
	err = engine.Untar(fs, zr, engine.PGData)
	zr.Close()
	if err != nil {
		return nil, fmt.Errorf("unpack pgdata: %w", err)
	}
	// A live single-user session cannot run CREATE DATABASE, so extra
	// setup goes through one standalone child, the way initdb does it.
	var setup []string
	if opts.User != "postgres" {
		setup = append(setup, "CREATE ROLE "+quoteIdent(opts.User)+" SUPERUSER LOGIN;")
	}
	if opts.Database != "postgres" {
		setup = append(setup, "CREATE DATABASE "+quoteIdent(opts.Database)+" OWNER "+quoteIdent(opts.User)+";")
	}
	if len(setup) > 0 {
		// postgres --single -j ends a command at ";\n\n", and CREATE DATABASE
		// must be its own transaction.
		if err := e.ExecStandalone(fs, "template1", strings.Join(setup, "\n\n")+"\n"); err != nil {
			return nil, fmt.Errorf("setup: %w", err)
		}
	}
	b, err := e.Start(fs, engine.StartOptions{User: opts.User, Database: opts.Database, Params: opts.Params})
	if err != nil {
		return nil, err
	}
	if opts.Log != nil {
		b.Stderr = func(p []byte) { opts.Log("%s", bytes.TrimRight(p, "\n")) }
	}
	s.fs, s.b = fs, b
	// Unpacking the share tree and data directory leaves the decompressor
	// window and tar buffers as garbage; give that back to the OS now so a
	// server's resident size is what it actually uses.
	debug.FreeOSMemory()
	ln, err := net.Listen("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(opts.Port)))
	if err != nil {
		b.Close()
		return nil, err
	}
	s.ln = ln
	s.port = ln.Addr().(*net.TCPAddr).Port
	s.wg.Add(1)
	go s.acceptLoop()
	return s, nil
}

// Dial returns an in-process connection to the server, bypassing TCP. It
// has the signature drivers expect for a custom dialer: with pgx set
// ConnConfig.DialFunc = s.Dial (the network and address are ignored), or
// register such a config with pgx's stdlib adapter for database/sql.
func (s *Server) Dial(ctx context.Context, network, addr string) (net.Conn, error) {
	if s.closed.Load() {
		return nil, errors.New("pgmem: server is closed")
	}
	client, server := net.Pipe()
	s.connsMu.Lock()
	s.conns[server] = struct{}{}
	s.connsMu.Unlock()
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.serve(server)
		s.connsMu.Lock()
		delete(s.conns, server)
		s.connsMu.Unlock()
	}()
	return client, nil
}

// Port returns the TCP port the server listens on.
func (s *Server) Port() int { return s.port }

// Addr returns "127.0.0.1:port".
func (s *Server) Addr() string { return net.JoinHostPort("127.0.0.1", strconv.Itoa(s.port)) }

// DSN returns a connection URL for the server.
func (s *Server) DSN() string {
	return fmt.Sprintf("postgres://%s@127.0.0.1:%d/%s?sslmode=disable", s.opts.User, s.port, s.opts.Database)
}

// Close stops the listener, drops connections and shuts the backend down.
func (s *Server) Close() error {
	if !s.closed.CompareAndSwap(false, true) {
		return nil
	}
	s.ln.Close()
	s.connsMu.Lock()
	for c := range s.conns {
		c.Close()
	}
	s.connsMu.Unlock()
	s.wg.Wait()
	s.mu.Lock()
	err := s.b.Close()
	s.mu.Unlock()
	return err
}

func (s *Server) acceptLoop() {
	defer s.wg.Done()
	for {
		c, err := s.ln.Accept()
		if err != nil {
			return
		}
		s.connsMu.Lock()
		s.conns[c] = struct{}{}
		s.connsMu.Unlock()
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.serve(c)
			s.connsMu.Lock()
			delete(s.conns, c)
			s.connsMu.Unlock()
		}()
	}
}

const (
	sslRequestCode    = 80877103
	gssEncRequestCode = 80877104
	cancelRequestCode = 80877102
)

func (s *Server) logf(format string, args ...any) {
	if s.opts.Log != nil {
		s.opts.Log(format, args...)
	}
}

func (s *Server) serve(c net.Conn) {
	defer c.Close()
	r := bufio.NewReaderSize(c, 64*1024)
	var pkt []byte
	for {
		var err error
		pkt, err = readStartupPacket(r)
		if err != nil {
			return
		}
		code := binary.BigEndian.Uint32(pkt[4:8])
		switch code {
		case sslRequestCode, gssEncRequestCode:
			if _, err := c.Write([]byte{'N'}); err != nil {
				return
			}
			continue
		case cancelRequestCode:
			return
		}
		break
	}
	id := s.connSeq.Add(1)
	prefix := "pgmem" + strconv.FormatInt(id, 36) + "_"

	s.mu.Lock()
	resp, err := s.startSession(pkt)
	s.mu.Unlock()
	if err != nil {
		s.logf("pgmem: startup failed: %v", err)
		c.Write(errorResponse("08006", err.Error()))
		return
	}
	if _, err := c.Write(resp); err != nil {
		return
	}

	for {
		batch, terminate, err := readBatch(r)
		if err != nil {
			return
		}
		if len(batch) > 0 {
			batch = rewriteNames(batch, prefix)
			s.mu.Lock()
			out, err := s.b.Exec(batch)
			s.mu.Unlock()
			if len(out) > 0 {
				if _, werr := c.Write(out); werr != nil {
					return
				}
			}
			if err != nil {
				s.logf("pgmem: exec failed: %v", err)
				c.Write(errorResponse("XX000", err.Error()))
				return
			}
		}
		if terminate {
			return
		}
	}
}

// startSession runs the startup handshake for the first connection and
// replays it (after resetting session state) for later ones.
func (s *Server) startSession(pkt []byte) ([]byte, error) {
	if s.startupResp == nil {
		resp, err := s.b.Startup(pkt)
		if err != nil {
			return nil, err
		}
		if err := checkNoError(resp); err != nil {
			return nil, err
		}
		s.startupResp = resp
		return resp, s.applyUser()
	}
	// Fresh session semantics for a new connection.
	if _, err := s.b.Exec(simpleQuery("ROLLBACK")); err != nil {
		return nil, err
	}
	if _, err := s.b.Exec(simpleQuery("DISCARD ALL")); err != nil {
		return nil, err
	}
	return s.startupResp, s.applyUser()
}

// applyUser switches the session to the configured user. The single-user
// backend always starts as the initdb superuser, and DISCARD ALL resets
// the session authorization, so this runs at every session start.
func (s *Server) applyUser() error {
	if s.opts.User == "postgres" {
		return nil
	}
	out, err := s.b.Exec(simpleQuery("SET SESSION AUTHORIZATION " + quoteIdent(s.opts.User)))
	if err != nil {
		return err
	}
	return checkNoError(out)
}

// ---- protocol helpers ----

func readStartupPacket(r *bufio.Reader) ([]byte, error) {
	var hdr [4]byte
	if _, err := io.ReadFull(r, hdr[:]); err != nil {
		return nil, err
	}
	n := binary.BigEndian.Uint32(hdr[:])
	if n < 8 || n > 1<<20 {
		return nil, fmt.Errorf("bad startup packet length %d", n)
	}
	pkt := make([]byte, n)
	copy(pkt, hdr[:])
	if _, err := io.ReadFull(r, pkt[4:]); err != nil {
		return nil, err
	}
	return pkt, nil
}

// readBatch reads at least one complete message and then every further
// message already buffered, so a pipelined Parse/Bind/Execute/Sync arrives
// together. A Terminate message ends the batch.
func readBatch(r *bufio.Reader) (batch []byte, terminate bool, err error) {
	for {
		var hdr [5]byte
		if _, err := io.ReadFull(r, hdr[:]); err != nil {
			if len(batch) > 0 && (err == io.EOF || errors.Is(err, io.ErrUnexpectedEOF)) {
				return batch, true, nil
			}
			return nil, false, err
		}
		n := binary.BigEndian.Uint32(hdr[1:])
		if n < 4 || n > 1<<30 {
			return nil, false, fmt.Errorf("bad message length %d", n)
		}
		msg := make([]byte, 1+n)
		copy(msg, hdr[:])
		if _, err := io.ReadFull(r, msg[5:]); err != nil {
			return nil, false, err
		}
		if msg[0] == 'X' {
			return batch, true, nil
		}
		batch = append(batch, msg...)
		if r.Buffered() == 0 {
			return batch, false, nil
		}
	}
}

// rewriteNames prefixes named prepared statements and portals so that
// several client connections sharing one backend session do not collide.
func rewriteNames(batch []byte, prefix string) []byte {
	var out bytes.Buffer
	for len(batch) >= 5 {
		typ := batch[0]
		n := binary.BigEndian.Uint32(batch[1:5])
		msg := batch[:1+n]
		batch = batch[1+n:]
		body := msg[5:]
		var nb []byte
		switch typ {
		case 'P': // Parse: stmt\0 query\0 ...
			nb = prefixCString(body, prefix, 0)
		case 'B': // Bind: portal\0 stmt\0 ...
			nb = prefixCString(prefixCString(body, prefix, 0), prefix, 1)
		case 'D', 'C': // Describe/Close: kind name\0
			if len(body) > 0 {
				nb = append([]byte{body[0]}, prefixCString(body[1:], prefix, 0)...)
			}
		case 'E': // Execute: portal\0 maxrows
			nb = prefixCString(body, prefix, 0)
		}
		if nb == nil {
			out.Write(msg)
			continue
		}
		out.WriteByte(typ)
		var l [4]byte
		binary.BigEndian.PutUint32(l[:], uint32(len(nb)+4))
		out.Write(l[:])
		out.Write(nb)
	}
	return out.Bytes()
}

// prefixCString prefixes the idx-th NUL-terminated string in body if it is
// non-empty.
func prefixCString(body []byte, prefix string, idx int) []byte {
	start := 0
	for i := 0; i < idx; i++ {
		z := bytes.IndexByte(body[start:], 0)
		if z < 0 {
			return body
		}
		start += z + 1
	}
	z := bytes.IndexByte(body[start:], 0)
	if z <= 0 {
		return body
	}
	out := make([]byte, 0, len(body)+len(prefix))
	out = append(out, body[:start]...)
	out = append(out, prefix...)
	out = append(out, body[start:]...)
	return out
}

func startupPacket(user, database string) []byte {
	var body bytes.Buffer
	binary.Write(&body, binary.BigEndian, uint32(196608))
	for _, kv := range [][2]string{{"user", user}, {"database", database}, {"client_encoding", "UTF8"}} {
		body.WriteString(kv[0])
		body.WriteByte(0)
		body.WriteString(kv[1])
		body.WriteByte(0)
	}
	body.WriteByte(0)
	var pkt bytes.Buffer
	binary.Write(&pkt, binary.BigEndian, uint32(body.Len()+4))
	pkt.Write(body.Bytes())
	return pkt.Bytes()
}

func simpleQuery(q string) []byte {
	var pkt bytes.Buffer
	pkt.WriteByte('Q')
	binary.Write(&pkt, binary.BigEndian, uint32(len(q)+1+4))
	pkt.WriteString(q)
	pkt.WriteByte(0)
	return pkt.Bytes()
}

func errorResponse(code, msg string) []byte {
	var body bytes.Buffer
	body.WriteString("SFATAL\x00")
	body.WriteString("C" + code + "\x00")
	body.WriteString("M" + msg + "\x00")
	body.WriteByte(0)
	var pkt bytes.Buffer
	pkt.WriteByte('E')
	binary.Write(&pkt, binary.BigEndian, uint32(body.Len()+4))
	pkt.Write(body.Bytes())
	return pkt.Bytes()
}

// checkNoError scans backend output for an ErrorResponse and returns it as
// a Go error.
func checkNoError(out []byte) error {
	for len(out) >= 5 {
		typ := out[0]
		n := binary.BigEndian.Uint32(out[1:5])
		if int(1+n) > len(out) {
			break
		}
		if typ == 'E' {
			fields := bytes.Split(out[5:1+n], []byte{0})
			var msg, code string
			for _, f := range fields {
				if len(f) == 0 {
					continue
				}
				switch f[0] {
				case 'M':
					msg = string(f[1:])
				case 'C':
					code = string(f[1:])
				}
			}
			return fmt.Errorf("%s: %s", code, msg)
		}
		out = out[1+n:]
	}
	return nil
}

// hasSetting reports whether params already carry "-c name=...".
func hasSetting(params []string, name string) bool {
	for _, p := range params {
		if strings.HasPrefix(strings.TrimSpace(p), name+"=") {
			return true
		}
	}
	return false
}

func quoteIdent(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
