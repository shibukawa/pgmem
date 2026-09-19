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
	"runtime"
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
	// CREATE DATABASE); the backend serves one database at a time and is
	// restarted on another one when a connection to it needs it.
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
	// SingleUser runs PostgreSQL in single-user mode: one backend session
	// that every connection shares, multiplexed by pgmem at transaction
	// boundaries (the pre-0.2 model). It starts and forks faster than the
	// default multi-process model but cannot run two sessions at once:
	// no lock waits between connections, no deadlock detection, session
	// state shared between connections. See the README for the trade-offs.
	SingleUser bool
	// WaitTimeout bounds how long a connection waits for the shared session
	// while the connection holding it sits idle inside a transaction. The
	// waiting connection is then terminated with SQLSTATE 55P03 and a
	// message naming the holder, instead of waiting forever: the usual
	// cause is a query sent through another pooled connection from inside
	// a transaction callback. 0 means 2s; negative waits forever.
	WaitTimeout time.Duration
}

// Server is a running in-memory PostgreSQL.
type Server struct {
	opts Options
	e    *engine.Engine
	ln   net.Listener
	port int

	// fs and b are the data directory and the backend. Restore and
	// database switches replace them while holding sem and bmu; Close
	// takes bmu to shut b down without waiting for a connection that
	// keeps sem.
	bmu sync.Mutex
	fs  *vfs.FS
	b   *engine.Backend
	// cl is the postmaster and its processes in the default multi-process
	// model (b is then nil); see serveCluster.
	cl *engine.Cluster

	// sem serializes use of the single backend. A connection holds it for
	// one message batch normally, and across batches while it is inside a
	// transaction (or mid-pipeline), so pooled connections interleave only
	// at transaction boundaries. holder is the id of the connection that
	// has it, for diagnostics. idleSince is when the holder last finished
	// a batch without giving the backend back (0 while a batch runs or
	// nobody holds it), which tells a waiter an idle transaction from a
	// slow statement.
	sem       chan struct{}
	holder    atomic.Int64
	idleSince atomic.Int64

	// current is the database the backend serves and dbs the session
	// bookkeeping of every database connected to (see databases.go).
	// Guarded by sem.
	current string
	dbs     map[string]*dbState

	// LISTEN/NOTIFY. The backend's own listen set is the union over the
	// connections to its database; these map each channel to the sessions
	// that asked for it, fed by the commit hook (Backend.Listen), so
	// NotifyResponse messages can be routed to the right connections.
	// relisten holds channels the backend dropped (UNLISTEN by one
	// connection) while others still want them. All guarded by sem.
	sessions    map[int64]*session
	listeners   map[listenKey]map[int64]*session
	relisten    map[listenKey]bool
	quietListen bool // pgmem is re-issuing LISTEN; ignore hook events

	// origin is the snapshot a fork was started from, which Reset returns
	// to. restored is the snapshot the data directory was last copied from
	// and dirty whether a client has run anything since, so a Restore with
	// nothing to undo returns at once. Guarded by sem.
	origin   *Snapshot
	restored *Snapshot
	dirty    bool

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
	if !o.SingleUser && runtime.GOOS == "windows" {
		// the shared memory of a cluster is mapped with mmap(MAP_FIXED);
		// Windows has no equivalent here yet
		o.SingleUser = true
	}
	if o.Database == "" {
		o.Database = "postgres"
	}
	if o.User == "" {
		o.User = "postgres"
	}
	if !hasSetting(o.Params, "shared_buffers") {
		o.Params = append([]string{"-c", "shared_buffers=32MB"}, o.Params...)
	}
	// PGlite runs the backend with IsUnderPostmaster set, so PostgreSQL
	// 18's default io_method=worker hands batched reads (sequential scans
	// through read_stream) to IO workers that do not exist and waits for
	// them forever. sync executes every IO in the backend itself.
	if !hasSetting(o.Params, "io_method") {
		o.Params = append([]string{"-c", "io_method=sync"}, o.Params...)
	}
	// For the same reason a parallel query or index build registers
	// workers that no postmaster will ever start, and the leader waits
	// for them to attach forever. With no worker slots the planner and
	// CREATE INDEX fall back to the leader doing all the work. A cluster
	// has a postmaster and starts them.
	if o.SingleUser && !hasSetting(o.Params, "max_parallel_workers") {
		o.Params = append([]string{"-c", "max_parallel_workers=0"}, o.Params...)
	}
	if o.WaitTimeout == 0 {
		o.WaitTimeout = 2 * time.Second
	}
	return o
}

// boot starts a backend on a filesystem that already holds a data
// directory and begins serving it. origin is the snapshot fs was copied
// from (nil for Start); onClose runs at the end of Close.
func boot(e *engine.Engine, opts Options, fs *vfs.FS, origin *Snapshot, onClose func()) (*Server, error) {
	s := &Server{
		opts:      opts,
		e:         e,
		conns:     map[net.Conn]struct{}{},
		sem:       make(chan struct{}, 1),
		done:      make(chan struct{}),
		current:   opts.Database,
		dbs:       map[string]*dbState{},
		sessions:  map[int64]*session{},
		listeners: map[listenKey]map[int64]*session{},
		relisten:  map[listenKey]bool{},
		origin:    origin,
		restored:  origin,
		onClose:   onClose,
	}
	if opts.SingleUser {
		b, err := s.startBackend(fs, opts.Database)
		if err != nil {
			return nil, err
		}
		s.b = b
	} else {
		cl, err := s.startCluster(fs)
		if err != nil {
			return nil, err
		}
		s.cl = cl
	}
	s.fs = fs
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

// startBackend boots a backend on fs serving database name, with the
// server's options and hooks.
func (s *Server) startBackend(fs *vfs.FS, name string) (*engine.Backend, error) {
	b, err := s.e.Start(fs, engine.StartOptions{User: s.opts.User, Database: name, Params: s.opts.Params})
	if err != nil {
		return nil, err
	}
	if s.opts.Log != nil {
		b.Stderr = func(p []byte) { s.opts.Log("%s", bytes.TrimRight(p, "\n")) }
	}
	b.Listen = s.onListen
	return b, nil
}

// startCluster boots a postmaster on fs with the server's options.
func (s *Server) startCluster(fs *vfs.FS) (*engine.Cluster, error) {
	var stderr func([]byte)
	if s.opts.Log != nil {
		stderr = func(p []byte) { s.opts.Log("%s", bytes.TrimRight(p, "\n")) }
	}
	return s.e.StartCluster(context.Background(), fs, engine.StartOptions{User: s.opts.User, Database: s.opts.Database, Params: s.opts.Params}, stderr)
}

// stopEngine shuts the backend or the cluster down.
func (s *Server) stopEngine() error {
	if s.cl != nil {
		return s.cl.Shutdown(context.Background())
	}
	if s.b != nil {
		return s.b.Close()
	}
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
	err := s.stopEngine() // single user: waits for a batch that is running
	s.bmu.Unlock()
	if s.onClose != nil {
		s.onClose()
	}
	return err
}

// lockWarnAfter is how long a connection waits for the backend before a
// diagnostic is logged. The usual cause is another connection sitting idle
// in a transaction, which serializes everyone else behind it.
const lockWarnAfter = 5 * time.Second

var errServerClosed = errors.New("pgmem: server is closed")

// idleHolderError reports a wait abandoned under Options.WaitTimeout.
type idleHolderError struct {
	waiter, holder int64
	waited         time.Duration
}

func (e *idleHolderError) Error() string {
	return fmt.Sprintf("connection %d waited %s for the session, which connection %d holds idle inside a transaction; "+
		"commit or roll back that transaction first (inside a transaction callback, send queries through the transaction, not another pooled connection)",
		e.waiter, e.waited.Round(time.Millisecond), e.holder)
}

// acquire takes exclusive use of the backend for connection id (0 for
// pgmem itself). It fails when the server closes or ctx ends and, with
// failFast, when the connection holding the backend stays idle inside a
// transaction for Options.WaitTimeout while this one waits: if the
// holder's client is waiting for this connection, nothing else would ever
// end the wait.
func (s *Server) acquire(ctx context.Context, id int64, failFast bool) error {
	select {
	case s.sem <- struct{}{}:
		s.holder.Store(id)
		return nil
	default:
	}
	start := time.Now()
	tick := time.NewTicker(50 * time.Millisecond)
	defer tick.Stop()
	warned := false
	for {
		select {
		case s.sem <- struct{}{}:
			s.holder.Store(id)
			return nil
		case <-s.done:
			return errServerClosed
		case <-ctx.Done():
			return ctx.Err()
		case now := <-tick.C:
			if limit := s.opts.WaitTimeout; failFast && limit > 0 {
				if idle := s.idleSince.Load(); idle != 0 {
					from := time.Unix(0, idle)
					if from.Before(start) {
						from = start
					}
					if now.Sub(from) >= limit {
						return &idleHolderError{waiter: id, holder: s.holder.Load(), waited: now.Sub(start)}
					}
				}
			}
			if !warned && now.Sub(start) >= lockWarnAfter {
				s.logf("pgmem: connection %d has waited %s for the backend held by connection %d (idle in transaction?)", id, lockWarnAfter, s.holder.Load())
				warned = true
			}
		}
	}
}

func (s *Server) release() {
	s.idleSince.Store(0)
	s.holder.Store(0)
	<-s.sem
}

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
		s.serve(c)
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

// session is one client connection.
type session struct {
	id       int64
	c        net.Conn
	database string
	wmu      sync.Mutex // serializes writes: batch replies and routed notifications
	// notify queues NotifyResponse messages routed from other connections;
	// a writer goroutine drains it so delivery never blocks the backend
	// (a net.Pipe client that is not reading would otherwise stall it).
	notify chan []byte
	// parses holds the Parse message of every named prepared statement the
	// connection has open (names already prefixed), to drop them when it
	// ends and to re-create them on a restarted backend. Guarded by the
	// server's sem.
	parses map[string][]byte
}

// record notes the named statements a batch creates and closes.
func (ss *session) record(batch []byte) {
	for len(batch) >= 5 {
		n := binary.BigEndian.Uint32(batch[1:5])
		if n < 4 || int(n)+1 > len(batch) {
			break
		}
		msg, body := batch[:1+n], batch[5:1+n]
		switch msg[0] {
		case 'P': // Parse: stmt\0 query\0 ...
			if z := bytes.IndexByte(body, 0); z > 0 {
				ss.parses[string(body[:z])] = append([]byte(nil), msg...)
			}
		case 'C': // Close: 'S' stmt\0 or 'P' portal\0
			if len(body) > 1 && body[0] == 'S' {
				if z := bytes.IndexByte(body[1:], 0); z > 0 {
					delete(ss.parses, string(body[1:1+z]))
				}
			}
		}
		batch = batch[1+n:]
	}
}

func (ss *session) write(p []byte) error {
	ss.wmu.Lock()
	defer ss.wmu.Unlock()
	_, err := ss.c.Write(p)
	return err
}

func (s *Server) serve(c net.Conn) {
	if s.cl != nil {
		s.serveCluster(c)
		return
	}
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
	db := startupParam(pkt, "database")
	if db == "" {
		db = s.opts.Database
	}
	id := s.connSeq.Add(1)
	prefix := "pgmem" + strconv.FormatInt(id, 36) + "_"

	sess := &session{id: id, c: c, database: db, notify: make(chan []byte, 256), parses: map[string][]byte{}}
	go func() {
		for msg := range sess.notify {
			sess.write(msg)
		}
	}()
	defer close(sess.notify)

	if err := s.acquire(context.Background(), id, true); err != nil {
		c.Write(s.acquireError(err))
		return
	}
	if err := s.use(db); err != nil {
		s.release()
		c.Write(s.useError(err))
		return
	}
	resp, err := s.startSession(pkt, db)
	if err != nil {
		s.release()
		s.logf("pgmem: startup failed: %v", err)
		c.Write(errorResponse("08006", err.Error()))
		return
	}
	s.db(db).live++
	s.sessions[id] = sess
	s.release()
	if err := sess.write(resp); err != nil {
		s.endSession(sess, false)
		return
	}

	// held is true while this connection keeps the backend across batches:
	// inside a transaction block, or between a pipelined message and its
	// Sync.
	held := false
	for {
		batch, terminate, err := readBatch(r)
		if s.closed.Load() {
			s.linger(sess, r, len(batch) > 0)
			return
		}
		if err != nil {
			s.endSession(sess, held)
			return
		}
		if len(batch) > 0 {
			batch = rewriteNames(batch, prefix)
			if !held {
				if err := s.acquire(context.Background(), id, true); err != nil {
					sess.write(s.acquireError(err))
					s.endSession(sess, false)
					return
				}
				held = true
				if err := s.use(sess.database); err != nil {
					sess.write(s.useError(err))
					s.endSession(sess, true)
					return
				}
			}
			s.idleSince.Store(0)
			s.dirty = true
			sess.record(batch)
			// A COPY FROM STDIN reads its data inside this Exec, so the
			// backend may ask for more of this connection's input than
			// the batch holds. If the client goes away or sends
			// Terminate in the middle, CopyFail ends the COPY with an
			// error instead of the fatal "protocol synchronization
			// was lost" that an EOF mid-message would cause.
			s.b.More = func() []byte {
				more, term, err := readBatch(r)
				if term {
					terminate = true
				}
				if err != nil || (term && len(more) == 0) {
					return copyFail("client disconnected")
				}
				more = rewriteNames(more, prefix)
				sess.record(more)
				return more
			}
			out, err := s.b.Exec(batch)
			s.b.More = nil
			out = s.afterExec(sess, out)
			if err == nil && backendIdle(out) {
				s.release()
				held = false
			} else if err == nil {
				s.idleSince.Store(time.Now().UnixNano())
			}
			if len(out) > 0 {
				if werr := sess.write(out); werr != nil {
					s.endSession(sess, held)
					return
				}
			}
			if err != nil {
				if s.closed.Load() { // Close shut the backend down under this batch
					sess.write(errorResponse("57P01", "terminating connection because the pgmem server was closed"))
					return
				}
				s.logf("pgmem: exec failed: %v", err)
				c.Write(errorResponse("XX000", err.Error()))
				s.endSession(sess, held)
				return
			}
		}
		if terminate {
			s.endSession(sess, held)
			return
		}
	}
}

// acquireError is what a connection is sent when it cannot have the
// backend for its next batch; the connection ends after it.
func (s *Server) acquireError(err error) []byte {
	var idle *idleHolderError
	if errors.As(err, &idle) {
		s.logf("pgmem: %v", err)
		return errorResponse("55P03", "pgmem: "+err.Error())
	}
	return errorResponse("57P01", "terminating connection because the pgmem server was closed")
}

// useError is what a connection is sent when the backend cannot serve its
// database; the connection ends after it.
func (s *Server) useError(err error) []byte {
	var missing *noDatabaseError
	switch {
	case errors.As(err, &missing):
		return errorResponse("3D000", err.Error())
	case errors.Is(err, errServerClosed):
		return errorResponse("57P01", "terminating connection because the pgmem server was closed")
	}
	s.logf("pgmem: %v", err)
	return errorResponse("XX000", err.Error())
}

// linger keeps a connection of a closed server open until its client
// closes it or closeGrace passes. A message from the client, or one that
// was already read (pending), is answered with 57P01, ending it.
func (s *Server) linger(sess *session, r *bufio.Reader, pending bool) {
	if !pending {
		sess.c.SetReadDeadline(time.Now().Add(closeGrace))
		if _, err := r.ReadByte(); err != nil {
			return
		}
	}
	sess.write(errorResponse("57P01", "terminating connection because the pgmem server was closed"))
}

// endSession cleans up after a connection: aborts a transaction it left
// open, drops its prepared statements and returns the backend.
func (s *Server) endSession(sess *session, held bool) {
	if !held && s.acquire(context.Background(), sess.id, false) != nil {
		return // closed: the backend and the bookkeeping are gone
	}
	defer s.release()
	s.db(sess.database).live--
	delete(s.sessions, sess.id)
	for k, m := range s.listeners {
		if k.database != sess.database {
			continue
		}
		delete(m, sess.id)
		if len(m) == 0 {
			delete(s.listeners, k)
		}
	}
	if s.closed.Load() || sess.database != s.current {
		return // the session's state went with its backend
	}
	if held {
		// The client went away mid-transaction, mid-pipeline or mid-COPY.
		// CopyFail is ignored outside COPY mode and aborts it inside; Sync
		// clears a pipelined error state; the abort ends the transaction
		// without a ROLLBACK statement anyone could observe.
		s.b.Exec(copyFail("client disconnected"))
		s.b.Exec(syncMessage)
		s.b.ResetSession(false)
	}
	if len(sess.parses) > 0 {
		var msg bytes.Buffer
		for name := range sess.parses {
			msg.WriteByte('C')
			binary.Write(&msg, binary.BigEndian, uint32(4+1+len(name)+1))
			msg.WriteByte('S')
			msg.WriteString(name)
			msg.WriteByte(0)
		}
		msg.Write(syncMessage)
		s.b.Exec(msg.Bytes())
	}
}

// onListen is the backend's commit hook for LISTEN/UNLISTEN. It runs inside
// Exec, so the current holder of the backend is the connection whose
// transaction just committed.
func (s *Server) onListen(channel string, op int) {
	if s.quietListen {
		return
	}
	id := s.holder.Load()
	switch op {
	case 1:
		sess := s.sessions[id]
		if sess == nil {
			return
		}
		key := listenKey{s.current, channel}
		m := s.listeners[key]
		if m == nil {
			m = map[int64]*session{}
			s.listeners[key] = m
		}
		m[id] = sess
	case 0:
		s.dropListener(listenKey{s.current, channel}, id)
	case 2:
		for k := range s.listeners {
			if k.database == s.current {
				s.dropListener(k, id)
			}
		}
	}
}

// dropListener records that the backend stopped listening on a channel
// because connection id unlistened. If other connections still listen,
// the backend has to be re-subscribed once it is idle again.
func (s *Server) dropListener(key listenKey, id int64) {
	m := s.listeners[key]
	if m == nil {
		return
	}
	delete(m, id)
	if len(m) == 0 {
		delete(s.listeners, key)
		return
	}
	s.relisten[key] = true
}

// afterExec post-processes the backend's reply to a batch from sess: it
// routes NotifyResponse messages to the connections listening on their
// channel (keeping only sess's own in the reply) and re-issues LISTEN for
// channels other connections still want. Runs with the backend held.
func (s *Server) afterExec(sess *session, out []byte) []byte {
	if len(s.listeners) == 0 && len(s.relisten) == 0 {
		return out
	}
	kept := make([]byte, 0, len(out))
	rest := out
	for len(rest) >= 5 {
		n := binary.BigEndian.Uint32(rest[1:5])
		if n < 4 || int(n)+1 > len(rest) {
			break
		}
		msg := rest[:1+n]
		rest = rest[1+n:]
		if msg[0] != 'A' || n < 9 {
			kept = append(kept, msg...)
			continue
		}
		body := msg[9:] // pid, channel\0, payload\0
		z := bytes.IndexByte(body, 0)
		if z < 0 {
			continue
		}
		for id, target := range s.listeners[listenKey{s.current, string(body[:z])}] {
			if id == sess.id {
				kept = append(kept, msg...)
				continue
			}
			select {
			case target.notify <- msg:
			default:
				s.logf("pgmem: connection %d is not draining notifications; dropped one", id)
			}
		}
	}
	out = append(kept, rest...)
	if len(s.relisten) > 0 && backendIdle(out) {
		var q strings.Builder
		for k := range s.relisten {
			if k.database == s.current {
				q.WriteString("LISTEN " + quoteIdent(k.channel) + ";")
				delete(s.relisten, k)
			}
		}
		if q.Len() > 0 {
			s.quietListen = true
			s.b.Exec(simpleQuery(q.String()))
			s.quietListen = false
		}
	}
	return out
}

// startSession runs the startup handshake for the first connection to a
// database and replays it (after resetting session state) for later ones.
// The backend serves the database already (see use).
func (s *Server) startSession(pkt []byte, name string) ([]byte, error) {
	st := s.db(name)
	if st.startupResp == nil {
		resp, err := s.b.Startup(pkt)
		if err != nil {
			return nil, err
		}
		if err := checkNoError(resp); err != nil {
			return nil, err
		}
		st.startupPkt = append([]byte(nil), pkt...)
		st.startupResp = resp
		return resp, s.applyUser()
	}
	// Fresh session semantics for a new connection, but only when no other
	// connection to the database is alive: the reset would drop the
	// prepared statements and temp tables of connections still using the
	// shared session. It is done in C rather than as ROLLBACK and DISCARD
	// ALL statements, so pg_stat_statements and the log do not see it.
	if st.live == 0 {
		if err := s.b.ResetSession(true); err != nil {
			return nil, err
		}
	}
	return st.startupResp, s.applyUser()
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

// backendIdle reports whether out, the backend's reply to a batch, ends
// with ReadyForQuery in the idle state. No ReadyForQuery means the batch
// was pipelined without a Sync (or opened COPY IN), and 'T' or 'E' means a
// transaction block is open; in both cases the connection keeps the backend.
func backendIdle(out []byte) bool {
	status := byte(0)
	for len(out) >= 5 {
		n := binary.BigEndian.Uint32(out[1:5])
		if n < 4 || int(n)+1 > len(out) {
			break
		}
		if out[0] == 'Z' && n == 5 {
			status = out[5]
		}
		out = out[1+n:]
	}
	return status == 'I'
}

// startupParam returns the value of key in a StartupMessage (length,
// protocol version, then NUL-terminated key/value pairs), or "".
func startupParam(pkt []byte, key string) string {
	if len(pkt) < 8 {
		return ""
	}
	fields := bytes.Split(pkt[8:], []byte{0})
	for i := 0; i+1 < len(fields); i += 2 {
		if len(fields[i]) == 0 {
			break
		}
		if string(fields[i]) == key {
			return string(fields[i+1])
		}
	}
	return ""
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

// syncMessage is a Sync message: it ends an extended-protocol batch.
var syncMessage = []byte{'S', 0, 0, 0, 4}

// copyFail is a CopyFail message: ignored outside COPY mode, it aborts a
// COPY FROM STDIN with an error inside it.
func copyFail(reason string) []byte {
	var pkt bytes.Buffer
	pkt.WriteByte('f')
	binary.Write(&pkt, binary.BigEndian, uint32(len(reason)+1+4))
	pkt.WriteString(reason)
	pkt.WriteByte(0)
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
	_, ok := settingValue(params, name)
	return ok
}

func quoteIdent(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
