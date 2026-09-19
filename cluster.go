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
	"strings"
	"sync"
	"time"

	"github.com/shibukawa/pgmem/internal/host"
	"github.com/shibukawa/pgmem/internal/vfs"
)

// The multi-process model: every client connection is handed to the
// postmaster, which starts a backend process for it, exactly as a
// PostgreSQL server does. Sessions, locks, transactions and their
// interactions are PostgreSQL's own.
//
// pgmem still sits between the client and the backend, for one reason: a
// Restore (or Snapshot) stops the cluster and starts it again, which ends
// every backend, while the promise to clients is that their connections
// survive. So a session remembers its startup packet, its named prepared
// statements and its LISTEN set; when its backend is gone and the client
// sends its next message, the session quietly gets a new backend, replays
// them and carries on. Between such events bytes are copied unchanged in
// both directions.

// serveCluster runs one client connection.
func (s *Server) serveCluster(c net.Conn) {
	sess := &csession{s: s, c: c, r: bufio.NewReaderSize(c, 64*1024), parses: map[string][]byte{}, listens: map[string]bool{}}
	sess.relay = &relay{c: c, sess: sess, got: make(chan struct{}, 1)}
	sess.run()
}

// csession is one client connection of a cluster server.
type csession struct {
	s       *Server
	c       net.Conn
	r       *bufio.Reader
	relay   *relay
	startup []byte
	parses  map[string][]byte // named Parse messages, by statement name
	listens map[string]bool

	// attachMu serializes attaching (the client's next message and an
	// eager re-attach after a restart both want to).
	attachMu sync.Mutex

	mu       sync.Mutex
	cs       *host.ConnSock // the backend side; nil while detached
	closedCh chan struct{}  // closed when the current backend side closes
	// idle is what the last ReadyForQuery said; busy is set from the
	// first client message after it until the next ReadyForQuery.
	idle bool
	busy bool
	// gone is set when the backend side closed for good: the client
	// terminated, or the server closed.
	gone bool
}

func (ss *csession) run() {
	c := ss.c
	s := ss.s
	defer s.untrackSession(ss)
	// the startup phase: SSL/GSS requests get 'N', a CancelRequest is
	// handed to a backend of its own, a StartupMessage opens the session
	for {
		pkt, err := readStartupPacket(ss.r)
		if err != nil {
			c.Close()
			return
		}
		code := binary.BigEndian.Uint32(pkt[4:8])
		switch code {
		case sslRequestCode, gssEncRequestCode:
			if _, err := c.Write([]byte{'N'}); err != nil {
				c.Close()
				return
			}
			continue
		case cancelRequestCode:
			// a backend processes it and closes; nothing comes back
			if cs, err := s.cl.Connect(io.Discard, nil); err == nil {
				cs.Push(pkt)
				cs.EOF()
			}
			c.Close()
			return
		}
		ss.startup = pkt
		break
	}
	s.trackSession(ss)
	if err := ss.attach(true); err != nil {
		if !errors.Is(err, errBackendRefused) {
			c.Write(errorResponse("08006", err.Error()))
		}
		c.Close()
		return
	}
	for {
		msg, err := readMessage(ss.r)
		if err != nil {
			if s.closed.Load() && !isClosedConn(err) {
				// Close woke us with a read deadline: keep the socket
				// open for the client (see Server.Close)
				c.SetReadDeadline(time.Time{})
				ss.linger(false)
				return
			}
			ss.clientGone()
			return
		}
		if s.closed.Load() {
			ss.linger(true)
			return
		}
		if ss.isGone() {
			// PostgreSQL ended the backend; the client has its FATAL and
			// will hang up
			ss.drain()
			return
		}
		if msg[0] == 'X' {
			ss.terminate(msg)
			return
		}
		if err := ss.send(msg); err != nil {
			if s.closed.Load() {
				ss.linger(true)
				return
			}
			s.logf("pgmem: connection lost its backend: %v", err)
			c.Write(errorResponse("08006", err.Error()))
			c.Close()
			return
		}
		ss.record(msg) // after send: a replay must not include this message
	}
}

// errBackendRefused is an attach that ended with the backend's own
// ErrorResponse, which the client has been sent.
var errBackendRefused = errors.New("backend refused the connection")

// send forwards one client message, attaching to a new backend first if
// the current one is gone.
func (ss *csession) send(msg []byte) error {
	ss.attachMu.Lock()
	defer ss.attachMu.Unlock()
	for attempt := 0; attempt < 2; attempt++ {
		ss.mu.Lock()
		cs := ss.cs
		ss.busy = true
		ss.mu.Unlock()
		if cs == nil {
			if err := ss.attach(false); err != nil {
				return err
			}
			continue
		}
		if cs.Push(msg) {
			return nil
		}
		// the backend went away under us: detach and try once more
		ss.mu.Lock()
		if ss.cs == cs {
			ss.cs = nil
		}
		ss.mu.Unlock()
	}
	return errors.New("pgmem: no backend for the connection")
}

// attach gets the session a backend: the startup packet is sent and its
// reply awaited. The first time the reply goes to the client; later (a
// new backend after a Restore) it is swallowed and the session's prepared
// statements and LISTEN registrations are re-created.
func (ss *csession) attach(first bool) error {
	s := ss.s
	closed := make(chan struct{})
	var once sync.Once
	var cs *host.ConnSock
	cs, err := s.cl.Connect(ss.relay, func() {
		once.Do(func() { close(closed) })
		ss.mu.Lock()
		current := ss.closedCh == closed
		if current {
			ss.cs = nil
		}
		gone := ss.gone
		ss.mu.Unlock()
		// A backend pgmem itself stopped (muted for a Restore or Close)
		// is replaced on the client's next message. One that PostgreSQL
		// ended (FATAL, crash) ends the connection: the client has been
		// sent the message. The socket is not closed here, though: a
		// client that writes its next query before reading the FATAL
		// would get a reset on Windows and lose the message. It stays
		// open until the client hangs up (or closeGrace passes), and
		// whatever the client still sends is swallowed (see run).
		if current && !gone && !cs.Muted() && !s.closed.Load() && !ss.relay.capturing() {
			ss.mu.Lock()
			ss.gone = true
			ss.mu.Unlock()
			ss.c.SetReadDeadline(time.Now().Add(closeGrace))
		}
	})
	if err != nil {
		return err
	}
	cs.Listen = func(channel string, op int) {
		// a backend pgmem is stopping unlistens everything on its way
		// out; that is not the client's doing
		if cs.Muted() {
			return
		}
		ss.onListen(channel, op)
	}
	ss.mu.Lock()
	ss.cs, ss.closedCh = cs, closed
	ss.mu.Unlock()
	ss.relay.startCapture()
	cs.Push(ss.startup)
	resp, err := ss.relay.waitReady(closed)
	if err != nil {
		ss.relay.stopCapture()
		return fmt.Errorf("pgmem: starting the session: %w", err)
	}
	if e := checkNoError(resp); e != nil {
		ss.relay.stopCapture()
		ss.c.Write(resp)
		if first {
			return errBackendRefused
		}
		return fmt.Errorf("pgmem: the restored server refused the session: %w", e)
	}
	if first {
		ss.relay.stopCapture()
		if _, err := ss.c.Write(resp); err != nil {
			return err
		}
		return nil
	}
	// a replacement backend: replay what the client believes exists
	if len(ss.parses) > 0 {
		var batch bytes.Buffer
		for _, p := range ss.parses {
			batch.Write(p)
		}
		batch.Write(syncMessage)
		cs.Push(batch.Bytes())
		if out, err := ss.relay.waitReady(closed); err != nil {
			ss.relay.stopCapture()
			return err
		} else if e := checkNoError(out); e != nil {
			s.logf("pgmem: re-creating prepared statements after restore: %v", e)
		}
	}
	if len(ss.listens) > 0 {
		var q strings.Builder
		for ch := range ss.listens {
			q.WriteString("LISTEN " + quoteIdent(ch) + ";")
		}
		cs.Push(simpleQuery(q.String()))
		if _, err := ss.relay.waitReady(closed); err != nil {
			ss.relay.stopCapture()
			return err
		}
	}
	ss.relay.stopCapture()
	return nil
}

// reattach gives a detached session a backend right away (after a
// restart), so that LISTEN registrations are live again before anyone
// notifies, without waiting for the client's next message.
func (ss *csession) reattach() {
	ss.attachMu.Lock()
	defer ss.attachMu.Unlock()
	ss.mu.Lock()
	detached := ss.cs == nil && !ss.gone
	ss.mu.Unlock()
	if !detached {
		return
	}
	if err := ss.attach(false); err != nil {
		ss.s.logf("pgmem: re-attaching a session after restart: %v", err)
	}
}

// record notes the named statements the client creates and closes.
func (ss *csession) record(msg []byte) {
	if len(msg) < 5 {
		return
	}
	body := msg[5:]
	switch msg[0] {
	case 'P': // Parse: stmt\0 query\0 ...
		if z := bytes.IndexByte(body, 0); z > 0 {
			ss.parses[string(body[:z])] = append([]byte(nil), msg...)
		}
	case 'C': // Close: 'S' stmt\0
		if len(body) > 1 && body[0] == 'S' {
			if z := bytes.IndexByte(body[1:], 0); z > 0 {
				delete(ss.parses, string(body[1:1+z]))
			}
		}
	case 'Q': // DEALLOCATE / DISCARD drop them too; be approximate
		q := strings.ToUpper(strings.TrimSpace(string(body)))
		if strings.HasPrefix(q, "DEALLOCATE ALL") || strings.HasPrefix(q, "DISCARD ALL") {
			ss.parses = map[string][]byte{}
		}
	}
}

// onListen is the backend's commit hook for this session's LISTEN set.
func (ss *csession) onListen(channel string, op int) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	switch op {
	case 1:
		ss.listens[channel] = true
	case 0:
		delete(ss.listens, channel)
	case 2:
		ss.listens = map[string]bool{}
	}
}

// terminate forwards the client's Terminate and closes once the backend
// has let go.
func (ss *csession) terminate(msg []byte) {
	ss.mu.Lock()
	cs, closed := ss.cs, ss.closedCh
	ss.gone = true
	ss.mu.Unlock()
	if cs != nil && cs.Push(msg) {
		cs.EOF()
		select {
		case <-closed:
		case <-time.After(5 * time.Second):
		}
	}
	ss.c.Close()
}

// clientGone ends the backend side after the client hung up.
func (ss *csession) clientGone() {
	ss.mu.Lock()
	cs := ss.cs
	ss.gone = true
	ss.mu.Unlock()
	if cs != nil {
		cs.EOF()
	}
	ss.c.Close()
}

// linger keeps the socket of a closed server open until the client sends
// something (or already has, pending) or closeGrace passes, then answers
// with 57P01 and closes. The backend's own shutdown messages were muted.
func (ss *csession) linger(pending bool) {
	c := ss.c
	if !pending {
		c.SetReadDeadline(time.Now().Add(closeGrace))
		if _, err := ss.r.ReadByte(); err != nil {
			c.Close()
			return
		}
	}
	c.Write(errorResponse("57P01", "terminating connection because the pgmem server was closed"))
	c.Close()
}

// isGone reports whether the session's backend ended for good.
func (ss *csession) isGone() bool {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	return ss.gone
}

// drain swallows what a client whose backend PostgreSQL ended still
// sends, until it hangs up or closeGrace passes, then closes the socket.
func (ss *csession) drain() {
	ss.c.SetReadDeadline(time.Now().Add(closeGrace))
	for {
		if _, err := readMessage(ss.r); err != nil {
			break
		}
	}
	ss.c.Close()
}

// inTransaction reports whether the session is inside a transaction or a
// batch, as far as the last ReadyForQuery says.
func (ss *csession) inTransaction() bool {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	return ss.cs != nil && (ss.busy || !ss.idle)
}

// mute silences the current backend (a closing or restarting server).
func (ss *csession) mute() {
	ss.mu.Lock()
	cs := ss.cs
	ss.mu.Unlock()
	if cs != nil {
		cs.Mute()
	}
}

// ---- backend -> client ----

// relay is the writer the backend side sends to. Normally it forwards to
// the client and watches ReadyForQuery to know the transaction status;
// while attaching it captures instead.
type relay struct {
	c    net.Conn
	sess *csession

	mu      sync.Mutex
	capture bool
	buf     bytes.Buffer
	got     chan struct{}
	partial []byte // an incomplete frame carried between writes
}

func (r *relay) Write(p []byte) (int, error) {
	r.mu.Lock()
	if r.capture {
		r.buf.Write(p)
		r.mu.Unlock()
		select {
		case r.got <- struct{}{}:
		default:
		}
		return len(p), nil
	}
	r.scan(p)
	r.mu.Unlock()
	return r.c.Write(p)
}

// scan tracks ReadyForQuery frames in the forwarded stream. Locked.
func (r *relay) scan(p []byte) {
	data := p
	if len(r.partial) > 0 {
		data = append(r.partial, p...)
		r.partial = nil
	}
	for len(data) >= 5 {
		n := binary.BigEndian.Uint32(data[1:5])
		if n < 4 || n > 1<<30 {
			return // not a frame boundary we understand; stop tracking this run
		}
		if uint64(len(data)) < uint64(n)+1 {
			break
		}
		if data[0] == 'Z' && n == 5 {
			ss := r.sess
			ss.mu.Lock()
			ss.idle = data[5] == 'I'
			ss.busy = false
			ss.mu.Unlock()
		}
		data = data[1+n:]
	}
	if len(data) > 0 {
		r.partial = append([]byte(nil), data...)
	}
}

func (r *relay) capturing() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.capture
}

func (r *relay) startCapture() {
	r.mu.Lock()
	r.capture = true
	r.buf.Reset()
	r.mu.Unlock()
}

func (r *relay) stopCapture() {
	r.mu.Lock()
	r.capture = false
	r.buf.Reset()
	r.mu.Unlock()
}

// waitReady waits until the captured reply ends with ReadyForQuery (or an
// ErrorResponse of a refused connection, which the backend follows by
// closing) and returns it.
func (r *relay) waitReady(closed chan struct{}) ([]byte, error) {
	timeout := time.NewTimer(60 * time.Second)
	defer timeout.Stop()
	for {
		r.mu.Lock()
		out := append([]byte(nil), r.buf.Bytes()...)
		r.mu.Unlock()
		if endsWithReadyForQuery(out) {
			r.mu.Lock()
			r.buf.Reset()
			r.mu.Unlock()
			return out, nil
		}
		select {
		case <-r.got:
		case <-closed:
			r.mu.Lock()
			out := append([]byte(nil), r.buf.Bytes()...)
			r.buf.Reset()
			r.mu.Unlock()
			if checkNoError(out) != nil {
				return out, nil // the backend refused and closed
			}
			return out, errors.New("the backend closed the connection")
		case <-timeout.C:
			return nil, errors.New("timed out waiting for the backend")
		}
	}
}

// endsWithReadyForQuery reports whether out is a whole sequence of frames
// whose last one is ReadyForQuery.
func endsWithReadyForQuery(out []byte) bool {
	last := byte(0)
	for len(out) >= 5 {
		n := binary.BigEndian.Uint32(out[1:5])
		if n < 4 || uint64(len(out)) < uint64(n)+1 {
			return false
		}
		last = out[0]
		out = out[1+n:]
	}
	return len(out) == 0 && last == 'Z'
}

// readMessage reads one complete frontend message.
func readMessage(r *bufio.Reader) ([]byte, error) {
	var hdr [5]byte
	if _, err := io.ReadFull(r, hdr[:]); err != nil {
		return nil, err
	}
	n := binary.BigEndian.Uint32(hdr[1:])
	if n < 4 || n > 1<<30 {
		return nil, fmt.Errorf("bad message length %d", n)
	}
	msg := make([]byte, 1+n)
	copy(msg, hdr[:])
	if _, err := io.ReadFull(r, msg[5:]); err != nil {
		return nil, err
	}
	return msg, nil
}

// isClosedConn reports whether a read error means the client went away
// (as opposed to a deadline set by Close).
func isClosedConn(err error) bool {
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return false
	}
	return true
}

// ---- sessions of the server ----

func (s *Server) trackSession(ss *csession) {
	s.connsMu.Lock()
	defer s.connsMu.Unlock()
	if s.csessions == nil {
		s.csessions = map[*csession]struct{}{}
	}
	s.csessions[ss] = struct{}{}
}

func (s *Server) untrackSession(ss *csession) {
	s.connsMu.Lock()
	defer s.connsMu.Unlock()
	delete(s.csessions, ss)
}

func (s *Server) eachSession(fn func(*csession)) {
	s.connsMu.Lock()
	list := make([]*csession, 0, len(s.csessions))
	for ss := range s.csessions {
		list = append(list, ss)
	}
	s.connsMu.Unlock()
	for _, ss := range list {
		fn(ss)
	}
}

// muteSessions silences every backend before the cluster is stopped.
func (s *Server) muteSessions() {
	s.eachSession(func(ss *csession) { ss.mute() })
}

// waitSessionsIdle waits until no session is inside a transaction or a
// batch, so a restart happens between transactions; ctx bounds the wait.
func (s *Server) waitSessionsIdle(ctx context.Context) error {
	tick := time.NewTicker(5 * time.Millisecond)
	defer tick.Stop()
	for {
		busy := false
		s.eachSession(func(ss *csession) {
			if ss.inTransaction() {
				busy = true
			}
		})
		if !busy {
			return nil
		}
		select {
		case <-tick.C:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// restartCluster stops the cluster and starts one on fs (a Restore, or a
// Snapshot that needs a clean shutdown). Client sessions keep their
// sockets and get a new backend on their next message.
func (s *Server) restartCluster(ctx context.Context, fs *vfs.FS) error {
	if err := s.waitSessionsIdle(ctx); err != nil {
		return fmt.Errorf("pgmem: restore waited for an open transaction to end (commit or close every connection first): %w", err)
	}
	s.bmu.Lock()
	defer s.bmu.Unlock()
	if s.closed.Load() {
		return errServerClosed
	}
	s.muteSessions()
	s.cl.Kill() // the old data directory is discarded: no shutdown needed
	cl, err := s.startCluster(fs)
	if err != nil {
		return fmt.Errorf("pgmem: restart: %w", err)
	}
	s.fs, s.cl = fs, cl
	s.reattachSessions()
	return nil
}

// reattachSessions gives every detached session a new backend, in
// parallel, and waits for them.
func (s *Server) reattachSessions() {
	var wg sync.WaitGroup
	s.eachSession(func(ss *csession) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ss.reattach()
		}()
	})
	wg.Wait()
}
