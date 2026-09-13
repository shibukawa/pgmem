package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shibukawa/pgmem"
)

// protocolVersion is the integer a wrapper checks in the ready event. It
// changes only when a message becomes incompatible.
const protocolVersion = 1

// endpoint describes one listening server (the template, an extra
// template started with "start", or a fork) in the same shape the ready
// line used before the control protocol existed.
type endpoint struct {
	ID       string `json:"id"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Database string `json:"database"`
	DSN      string `json:"dsn"`
}

func newEndpoint(id string, s *pgmem.Server, user, database string) endpoint {
	return endpoint{ID: id, Host: "127.0.0.1", Port: s.Port(), User: user, Database: database, DSN: s.DSN()}
}

// request is one line from a client. Fields not used by an op are ignored.
type request struct {
	ID       *json.RawMessage `json:"id"`
	Op       string           `json:"op"`
	Server   string           `json:"server"`
	Snapshot string           `json:"snapshot"`
	Database string           `json:"database"`
	User     string           `json:"user"`
	Params   []string         `json:"params"`
	MaxForks int              `json:"max_forks"`
	Timeout  int              `json:"timeout_ms"`
	Token    string           `json:"token"`
}

type protoError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *protoError) Error() string { return e.Code + ": " + e.Message }

func perr(code, format string, args ...any) error {
	return &protoError{Code: code, Message: fmt.Sprintf(format, args...)}
}

type serverEntry struct {
	srv      *pgmem.Server
	user     string
	database string
	forked   bool // started by fork, so reset without a snapshot has a target
}

type snapshotEntry struct {
	snap     *pgmem.Snapshot
	user     string
	database string
}

// controlInfo is the "control" object of the ready event.
type controlInfo struct {
	Addr  string `json:"addr"`
	Token string `json:"token"`
	URL   string `json:"url"`
}

// controller owns every server, snapshot and fork of the process. It serves
// the control protocol on stdin/stdout for the process that started it and,
// after listen, on a loopback socket for other processes such as test
// workers.
type controller struct {
	base pgmem.Options // Params and Log for servers started later

	mu        sync.Mutex
	servers   map[string]*serverEntry
	snapshots map[string]*snapshotEntry
	seq       int
	closed    bool
	sockets   map[net.Conn]struct{}

	stdio *client
	ln    net.Listener
	token string
}

// client is one control channel: stdin/stdout, or a socket connection.
type client struct {
	wmu    sync.Mutex
	w      io.Writer
	socket bool
	authed bool            // socket clients must say hello with the token first
	ctx    context.Context // ends when the client goes away

	mu    sync.Mutex
	owned []string // ids of servers and snapshots created on a socket, closed with it
}

func newController(base pgmem.Options, w io.Writer) *controller {
	return &controller{
		base:      base,
		servers:   map[string]*serverEntry{},
		snapshots: map[string]*snapshotEntry{},
		sockets:   map[net.Conn]struct{}{},
		stdio:     &client{w: w, authed: true, ctx: context.Background()},
	}
}

// add registers a server under id (empty = next "t<n>").
func (c *controller) add(id string, s *pgmem.Server, user, database string) endpoint {
	c.mu.Lock()
	defer c.mu.Unlock()
	if id == "" {
		c.seq++
		id = "t" + strconv.Itoa(c.seq+1)
	}
	c.servers[id] = &serverEntry{srv: s, user: user, database: database}
	return newEndpoint(id, s, user, database)
}

// writeLine writes one JSON object followed by a newline, serialized so
// concurrent handlers never interleave.
func (cl *client) writeLine(v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	cl.wmu.Lock()
	defer cl.wmu.Unlock()
	_, err = cl.w.Write(append(b, '\n'))
	return err
}

// own records that id was created on a socket connection.
func (cl *client) own(id string) {
	if !cl.socket {
		return
	}
	cl.mu.Lock()
	cl.owned = append(cl.owned, id)
	cl.mu.Unlock()
}

// listen opens the control socket on addr, which must be a loopback
// address, and starts accepting connections.
func (c *controller) listen(addr string) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	if ip := net.ParseIP(host); host != "localhost" && (ip == nil || !ip.IsLoopback()) {
		return fmt.Errorf("control address %q is not a loopback address", addr)
	}
	var secret [16]byte
	if _, err := rand.Read(secret[:]); err != nil {
		return err
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	c.ln, c.token = ln, hex.EncodeToString(secret[:])
	go c.acceptControl()
	return nil
}

// control describes the control socket, or is nil when there is none.
func (c *controller) control() *controlInfo {
	if c.ln == nil {
		return nil
	}
	addr := c.ln.Addr().String()
	return &controlInfo{Addr: addr, Token: c.token, URL: "pgmem-control://" + c.token + "@" + addr}
}

// ready writes the first stdout line: the legacy flat fields plus the
// event fields wrappers use to negotiate the protocol.
func (c *controller) ready(pid int, version string, tmpl endpoint) error {
	ev := map[string]any{
		"event": "ready", "protocol": protocolVersion, "version": version, "pid": pid,
		"server": tmpl,
		// pre-protocol wrappers read these
		"host": tmpl.Host, "port": tmpl.Port, "user": tmpl.User, "database": tmpl.Database, "dsn": tmpl.DSN,
	}
	if info := c.control(); info != nil {
		ev["control"] = info
	}
	return c.stdio.writeLine(ev)
}

func (c *controller) fatal(msg string) {
	c.stdio.writeLine(map[string]any{"event": "fatal", "message": msg})
}

// serve reads requests from stdin until EOF or a shutdown request arrives,
// then closes everything. It returns when all handlers have finished.
func (c *controller) serve(r io.Reader) {
	ctx, cancel := context.WithCancel(context.Background())
	c.stdio.ctx = ctx
	c.run(c.stdio, r, func() {
		cancel()
		c.closeAll()
	})
}

func (c *controller) acceptControl() {
	for {
		conn, err := c.ln.Accept()
		if err != nil {
			return
		}
		c.mu.Lock()
		if c.closed {
			c.mu.Unlock()
			conn.Close()
			return
		}
		c.sockets[conn] = struct{}{}
		c.mu.Unlock()
		go c.serveSocket(conn)
	}
}

// serveSocket serves one control connection. Servers and snapshots created
// on it are closed when it ends, so a test worker that exits or crashes
// cannot keep fork slots.
func (c *controller) serveSocket(conn net.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
	cl := &client{w: conn, socket: true, ctx: ctx}
	c.run(cl, conn, cancel)
	conn.Close()
	c.mu.Lock()
	delete(c.sockets, conn)
	c.mu.Unlock()
	c.release(cl)
}

// run reads requests from r and handles each in its own goroutine, so a
// fork waiting for a slot never delays the close that frees one. It stops
// at the end of r, at shutdown (stdin only) or when a socket client fails
// to authenticate, calls end, and returns once every handler finished.
func (c *controller) run(cl *client, r io.Reader, end func()) {
	var handlers sync.WaitGroup
	stop := make(chan struct{})
	lines := make(chan []byte)
	go func() {
		defer close(lines)
		sc := bufio.NewScanner(r)
		sc.Buffer(make([]byte, 0, 64*1024), 16*1024*1024)
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())
			if line == "" {
				continue
			}
			select {
			case lines <- []byte(line):
			case <-stop:
				return
			}
		}
	}()
loop:
	for line := range lines {
		var req request
		if err := json.Unmarshal(line, &req); err != nil {
			cl.writeLine(map[string]any{"id": nil, "ok": false, "error": protoError{"protocol", "malformed request: " + err.Error()}})
			continue
		}
		switch {
		case req.Op == "hello":
			if cl.socket && subtle.ConstantTimeCompare([]byte(req.Token), []byte(c.token)) != 1 {
				c.reply(cl, req.ID, nil, perr("unauthorized", "wrong control token"))
				break loop
			}
			cl.authed = true
			c.reply(cl, req.ID, map[string]any{"protocol": protocolVersion, "version": buildVersion()}, nil)
			continue
		case !cl.authed:
			c.reply(cl, req.ID, nil, perr("unauthorized", "send hello with the control token first"))
			break loop
		case req.Op == "shutdown":
			if cl.socket {
				c.reply(cl, req.ID, nil, perr("forbidden", "shutdown is accepted on stdin only"))
				continue
			}
			c.reply(cl, req.ID, map[string]any{}, nil)
			break loop
		}
		handlers.Add(1)
		go func() {
			defer handlers.Done()
			res, err := c.handle(cl, req)
			c.reply(cl, req.ID, res, err)
		}()
	}
	close(stop)
	end()
	handlers.Wait()
}

func (c *controller) reply(cl *client, id *json.RawMessage, res map[string]any, err error) {
	var rawID any
	if id != nil {
		rawID = id
	}
	if err != nil {
		var pe *protoError
		if !errors.As(err, &pe) {
			pe = &protoError{Code: "internal", Message: err.Error()}
		}
		cl.writeLine(map[string]any{"id": rawID, "ok": false, "error": pe})
		return
	}
	out := map[string]any{"id": rawID, "ok": true}
	for k, v := range res {
		out[k] = v
	}
	cl.writeLine(out)
}

func (c *controller) handle(cl *client, req request) (map[string]any, error) {
	switch req.Op {
	case "start":
		return c.opStart(cl, req)
	case "snapshot":
		return c.opSnapshot(cl, req)
	case "fork":
		return c.opFork(cl, req)
	case "reset":
		return c.opReset(cl, req)
	case "close":
		return c.opClose(req)
	case "":
		return nil, perr("protocol", "missing op")
	default:
		return nil, perr("unknown_op", "unknown op %q", req.Op)
	}
}

// timeout bounds ctx by a request's timeout_ms, if it has one.
func timeout(ctx context.Context, ms int) (context.Context, context.CancelFunc) {
	if ms > 0 {
		return context.WithTimeout(ctx, time.Duration(ms)*time.Millisecond)
	}
	return context.WithCancel(ctx)
}

func (c *controller) opStart(cl *client, req request) (map[string]any, error) {
	if c.isClosed() {
		return nil, perr("internal", "shutting down")
	}
	opts := c.base
	opts.Database = req.Database
	opts.User = req.User
	if opts.Database == "" {
		opts.Database = "postgres"
	}
	if opts.User == "" {
		opts.User = "postgres"
	}
	for _, p := range req.Params {
		opts.Params = append(opts.Params, "-c", p)
	}
	s, err := pgmem.Start(cl.ctx, opts)
	if err != nil {
		return nil, err
	}
	ep := c.add("", s, opts.User, opts.Database)
	cl.own(ep.ID)
	return map[string]any{"server": ep}, nil
}

func (c *controller) opSnapshot(cl *client, req request) (map[string]any, error) {
	id := req.Server
	if id == "" {
		id = "template"
	}
	c.mu.Lock()
	se, ok := c.servers[id]
	c.mu.Unlock()
	if !ok {
		return nil, perr("unknown_id", "no server %q", id)
	}
	ctx, cancel := timeout(cl.ctx, req.Timeout)
	defer cancel()
	snap, err := se.srv.Snapshot(ctx, pgmem.SnapshotOptions{MaxForks: req.MaxForks})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, perr("busy", "a connection to %s is still in a transaction after %dms; commit or close every connection before snapshot", id, req.Timeout)
		}
		return nil, err
	}
	c.mu.Lock()
	c.seq++
	sid := "s" + strconv.Itoa(c.seq)
	c.snapshots[sid] = &snapshotEntry{snap: snap, user: se.user, database: se.database}
	c.mu.Unlock()
	cl.own(sid)
	return map[string]any{"snapshot": sid}, nil
}

func (c *controller) opFork(cl *client, req request) (map[string]any, error) {
	c.mu.Lock()
	sn, ok := c.snapshots[req.Snapshot]
	c.mu.Unlock()
	if !ok {
		return nil, perr("unknown_id", "no snapshot %q", req.Snapshot)
	}
	ctx, cancel := timeout(cl.ctx, req.Timeout)
	defer cancel()
	s, err := sn.snap.Fork(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, perr("pool_timeout", "no fork slot became free within %dms", req.Timeout)
		}
		if strings.Contains(err.Error(), "snapshot is closed") {
			return nil, perr("snapshot_closed", "snapshot %q is closed", req.Snapshot)
		}
		return nil, err
	}
	c.mu.Lock()
	c.seq++
	fid := "f" + strconv.Itoa(c.seq)
	c.servers[fid] = &serverEntry{srv: s, user: sn.user, database: sn.database, forked: true}
	c.mu.Unlock()
	cl.own(fid)
	return map[string]any{"server": newEndpoint(fid, s, sn.user, sn.database)}, nil
}

// opReset puts a server back to a snapshot in place: the fork's own
// snapshot by default, or the one named. Its port and client connections
// survive.
func (c *controller) opReset(cl *client, req request) (map[string]any, error) {
	c.mu.Lock()
	se, ok := c.servers[req.Server]
	sn := c.snapshots[req.Snapshot]
	c.mu.Unlock()
	if !ok {
		return nil, perr("unknown_id", "no server %q", req.Server)
	}
	if req.Snapshot != "" && sn == nil {
		return nil, perr("unknown_id", "no snapshot %q", req.Snapshot)
	}
	if sn == nil && !se.forked {
		return nil, perr("protocol", "server %q was not started by fork; name the snapshot to reset to", req.Server)
	}
	ctx, cancel := timeout(cl.ctx, req.Timeout)
	defer cancel()
	var err error
	if sn != nil {
		err = se.srv.Restore(ctx, sn.snap)
	} else {
		err = se.srv.Reset(ctx)
	}
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, perr("busy", "a connection to %s is still in a transaction after %dms; commit or close it before reset", req.Server, req.Timeout)
		}
		return nil, err
	}
	return map[string]any{}, nil
}

// opClose is idempotent: an id that is unknown (already closed) succeeds.
func (c *controller) opClose(req request) (map[string]any, error) {
	if req.Server == "" && req.Snapshot == "" {
		return nil, perr("protocol", "close needs server or snapshot")
	}
	c.mu.Lock()
	se := c.servers[req.Server]
	delete(c.servers, req.Server)
	sn := c.snapshots[req.Snapshot]
	delete(c.snapshots, req.Snapshot)
	c.mu.Unlock()
	if se != nil {
		se.srv.Close()
	}
	if sn != nil {
		sn.snap.Close()
	}
	return map[string]any{}, nil
}

// release closes what a socket client created: servers first, then
// snapshots.
func (c *controller) release(cl *client) {
	cl.mu.Lock()
	owned := cl.owned
	cl.owned = nil
	cl.mu.Unlock()
	var servers []*serverEntry
	var snaps []*snapshotEntry
	c.mu.Lock()
	for _, id := range owned {
		if se := c.servers[id]; se != nil {
			servers = append(servers, se)
			delete(c.servers, id)
		}
		if sn := c.snapshots[id]; sn != nil {
			snaps = append(snaps, sn)
			delete(c.snapshots, id)
		}
	}
	c.mu.Unlock()
	for _, se := range servers {
		se.srv.Close()
	}
	for _, sn := range snaps {
		sn.snap.Close()
	}
}

func (c *controller) isClosed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.closed
}

// closeAll stops the control socket and closes forks first, then
// snapshots, then templates.
func (c *controller) closeAll() {
	c.mu.Lock()
	c.closed = true
	servers, snaps := c.servers, c.snapshots
	c.servers, c.snapshots = map[string]*serverEntry{}, map[string]*snapshotEntry{}
	for conn := range c.sockets {
		conn.Close()
	}
	c.mu.Unlock()
	if c.ln != nil {
		c.ln.Close()
	}
	for id, se := range servers {
		if strings.HasPrefix(id, "f") {
			se.srv.Close()
		}
	}
	for _, sn := range snaps {
		sn.snap.Close()
	}
	for id, se := range servers {
		if !strings.HasPrefix(id, "f") {
			se.srv.Close()
		}
	}
}
