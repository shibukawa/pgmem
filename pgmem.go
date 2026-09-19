// Package pgmem runs a real PostgreSQL server entirely in memory, inside the
// test process, and exposes it over the normal wire protocol on a loopback
// TCP port. Any PostgreSQL driver can connect to it.
//
// The server is PostgreSQL (the PGlite fork) compiled to WebAssembly and then
// translated to Go ahead of time, running on an in-memory filesystem. It is a
// single-user, single-session backend: every TCP connection shares one
// backend session. Connections are multiplexed onto it the way a
// transaction-mode pooler does: a connection owns the backend for the
// duration of a transaction block, and others wait. Pools of any size work,
// but they serialize rather than run in parallel, and session-level state
// (SET, temp tables) is shared between live connections.
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
	"time"

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
	// Connections may also name other databases of the server (made with
	// CREATE DATABASE).
	Database string
	User     string
	// Port to listen on (0 = pick a free port). Always binds 127.0.0.1.
	Port int
	// Params are extra "postgres" command-line arguments, e.g.
	// []string{"-c", "log_statement=all"}. Unless it sets shared_buffers,
	// pgmem uses 32MB instead of initdb's 128MB: a test database does not
	// need a large buffer cache and every megabyte here is resident memory.
	// io_method defaults to sync (see Options.withDefaults).
	Params []string
	// Log receives server log output and host diagnostics (nil = discard).
	Log func(format string, args ...any)
}

// Server is a running in-memory PostgreSQL.
type Server struct {
	opts Options
	e    *engine.Engine
	ln   net.Listener
	port int

	// fs is the data directory and cl the postmaster with its processes.
	// Restore replaces them while holding bmu; Close takes bmu to stop cl.
	bmu sync.Mutex
	fs  *vfs.FS
	cl  *engine.Cluster

	// origin is the snapshot a fork was started from, which Reset returns
	// to.
	origin *Snapshot

	onClose   func()        // set for forks: returns the Snapshot's slot
	done      chan struct{} // closed by Close
	connSeq   atomic.Int64
	closed    atomic.Bool
	connsMu   sync.Mutex
	conns     map[net.Conn]struct{}
	csessions map[*csession]struct{} // cluster model: the live sessions
}

// Start boots a fresh server: compiles the module, runs initdb into memory,
// starts the backend and listens on 127.0.0.1.
func Start(ctx context.Context, opts Options) (*Server, error) {
	opts = opts.withDefaults()
	e, err := sharedEngine()
	if err != nil {
		return nil, err
	}
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
	// Unpacking the share tree and data directory leaves the decompressor
	// window and tar buffers as garbage; give that back to the OS now so a
	// server's resident size is what it actually uses.
	debug.FreeOSMemory()
	return boot(e, opts, fs, nil, nil)
}

func (o Options) withDefaults() Options {
	if o.Database == "" {
		o.Database = "postgres"
	}
	if o.User == "" {
		o.User = "postgres"
	}
	if !hasSetting(o.Params, "shared_buffers") {
		o.Params = append([]string{"-c", "shared_buffers=32MB"}, o.Params...)
	}
	// PostgreSQL 18's default io_method=worker would add three I/O worker
	// processes for no gain here; sync does every I/O in the backend.
	if !hasSetting(o.Params, "io_method") {
		o.Params = append([]string{"-c", "io_method=sync"}, o.Params...)
	}
	return o
}

// boot starts a cluster on a filesystem that already holds a data
// directory and begins serving it. origin is the snapshot fs was copied
// from (nil for Start); onClose runs at the end of Close.
func boot(e *engine.Engine, opts Options, fs *vfs.FS, origin *Snapshot, onClose func()) (*Server, error) {
	s := &Server{
		opts:    opts,
		e:       e,
		conns:   map[net.Conn]struct{}{},
		done:    make(chan struct{}),
		origin:  origin,
		onClose: onClose,
	}
	cl, err := s.startCluster(fs)
	if err != nil {
		return nil, err
	}
	s.cl, s.fs = cl, fs
	ln, err := net.Listen("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(opts.Port)))
	if err != nil {
		s.stopEngine()
		return nil, err
	}
	s.ln = ln
	s.port = ln.Addr().(*net.TCPAddr).Port
	go s.acceptLoop()
	return s, nil
}

// startCluster boots a postmaster on fs with the server's options.
func (s *Server) startCluster(fs *vfs.FS) (*engine.Cluster, error) {
	var stderr func([]byte)
	if s.opts.Log != nil {
		stderr = func(p []byte) { s.opts.Log("%s", bytes.TrimRight(p, "\n")) }
	}
	return s.e.StartCluster(context.Background(), fs, engine.StartOptions{User: s.opts.User, Database: s.opts.Database, Params: s.opts.Params}, stderr)
}

// stopEngine kills the cluster: its data directory is not used again.
func (s *Server) stopEngine() error {
	s.cl.Kill()
	return nil
}

// Dial returns an in-process connection to the server, bypassing TCP. It
// has the signature drivers expect for a custom dialer: with pgx set
// ConnConfig.DialFunc = s.Dial (the network and address are ignored), or
// register such a config with pgx's stdlib adapter for database/sql.
func (s *Server) Dial(ctx context.Context, network, addr string) (net.Conn, error) {
	if s.closed.Load() {
		return nil, errServerClosed
	}
	client, server := net.Pipe()
	s.track(server)
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

// closeGrace is how long a client connection of a closed server stays open
// while its client does nothing. Pools keep idle connections, and
// node-postgres among others raises an unhandled error event when the
// server drops one, so a closed server leaves the socket to the client
// instead: it answers the next message with 57P01 and closes.
const closeGrace = 30 * time.Second

// Close stops the listener and shuts the backend down. Client connections
// stay open until their client closes them, sends a message (answered with
// SQLSTATE 57P01) or closeGrace passes.
func (s *Server) Close() error {
	if !s.closed.CompareAndSwap(false, true) {
		return nil
	}
	close(s.done)
	s.ln.Close()
	s.connsMu.Lock()
	for ss := range s.csessions {
		ss.mute() // the backends' shutdown messages must not reach idle clients
	}
	for c := range s.conns {
		c.SetReadDeadline(time.Now()) // wake the reader; serve lingers from there
	}
	s.connsMu.Unlock()
	s.bmu.Lock()
	err := s.stopEngine()
	s.bmu.Unlock()
	if s.onClose != nil {
		s.onClose()
	}
	return err
}

var errServerClosed = errors.New("pgmem: server is closed")

func (s *Server) acceptLoop() {
	for {
		c, err := s.ln.Accept()
		if err != nil {
			return
		}
		s.track(c)
	}
}

// track registers a client connection and serves it in its own goroutine.
// A connection that arrives while Close runs is woken like the others.
func (s *Server) track(c net.Conn) {
	s.connsMu.Lock()
	s.conns[c] = struct{}{}
	if s.closed.Load() {
		c.SetReadDeadline(time.Now())
	}
	s.connsMu.Unlock()
	go func() {
		s.serveCluster(c)
		s.connsMu.Lock()
		delete(s.conns, c)
		s.connsMu.Unlock()
	}()
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

// syncMessage is a Sync message: it ends an extended-protocol batch.
var syncMessage = []byte{'S', 0, 0, 0, 4}

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
	_, ok := settingValue(params, name)
	return ok
}

func quoteIdent(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
