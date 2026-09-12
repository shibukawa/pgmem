package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

// request is one line on stdin. Fields not used by an op are ignored.
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
}

type snapshotEntry struct {
	snap     *pgmem.Snapshot
	user     string
	database string
}

// controller owns every server, snapshot and fork of the process and
// serves the control protocol on one reader/writer pair.
type controller struct {
	base pgmem.Options // Params and Log for servers started later

	mu        sync.Mutex
	servers   map[string]*serverEntry
	snapshots map[string]*snapshotEntry
	seq       int
	closed    bool

	wmu sync.Mutex
	w   io.Writer
}

func newController(base pgmem.Options, w io.Writer) *controller {
	return &controller{
		base:      base,
		servers:   map[string]*serverEntry{},
		snapshots: map[string]*snapshotEntry{},
		w:         w,
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
func (c *controller) writeLine(v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.wmu.Lock()
	defer c.wmu.Unlock()
	_, err = c.w.Write(append(b, '\n'))
	return err
}

// ready writes the first stdout line: the legacy flat fields plus the
// event fields wrappers use to negotiate the protocol.
func (c *controller) ready(pid int, version string, tmpl endpoint) error {
	return c.writeLine(map[string]any{
		"event": "ready", "protocol": protocolVersion, "version": version, "pid": pid,
		"server": tmpl,
		// pre-protocol wrappers read these
		"host": tmpl.Host, "port": tmpl.Port, "user": tmpl.User, "database": tmpl.Database, "dsn": tmpl.DSN,
	})
}

func (c *controller) fatal(msg string) {
	c.writeLine(map[string]any{"event": "fatal", "message": msg})
}

// serve reads requests until r hits EOF or a shutdown request arrives,
// then closes everything. It returns when all handlers have finished.
func (c *controller) serve(r io.Reader) {
	var handlers sync.WaitGroup
	shutdown := make(chan struct{})
	lines := make(chan []byte)
	go func() {
		sc := bufio.NewScanner(r)
		sc.Buffer(make([]byte, 0, 64*1024), 16*1024*1024)
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())
			if line == "" {
				continue
			}
			select {
			case lines <- []byte(line):
			case <-shutdown:
				return
			}
		}
		close(lines)
	}()
loop:
	for {
		select {
		case line, ok := <-lines:
			if !ok {
				break loop // stdin closed: the parent is gone or done
			}
			var req request
			if err := json.Unmarshal(line, &req); err != nil {
				c.writeLine(map[string]any{"id": nil, "ok": false, "error": protoError{"protocol", "malformed request: " + err.Error()}})
				continue
			}
			if req.Op == "shutdown" {
				c.reply(req.ID, map[string]any{}, nil)
				break loop
			}
			handlers.Add(1)
			go func() {
				defer handlers.Done()
				res, err := c.handle(req)
				c.reply(req.ID, res, err)
			}()
		}
	}
	close(shutdown)
	c.closeAll()
	handlers.Wait()
}

func (c *controller) reply(id *json.RawMessage, res map[string]any, err error) {
	var rawID any
	if id != nil {
		rawID = id
	}
	if err != nil {
		var pe *protoError
		if !errors.As(err, &pe) {
			pe = &protoError{Code: "internal", Message: err.Error()}
		}
		c.writeLine(map[string]any{"id": rawID, "ok": false, "error": pe})
		return
	}
	out := map[string]any{"id": rawID, "ok": true}
	for k, v := range res {
		out[k] = v
	}
	c.writeLine(out)
}

func (c *controller) handle(req request) (map[string]any, error) {
	switch req.Op {
	case "start":
		return c.opStart(req)
	case "snapshot":
		return c.opSnapshot(req)
	case "fork":
		return c.opFork(req)
	case "close":
		return c.opClose(req)
	case "":
		return nil, perr("protocol", "missing op")
	default:
		return nil, perr("unknown_op", "unknown op %q", req.Op)
	}
}

func (c *controller) opStart(req request) (map[string]any, error) {
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
	s, err := pgmem.Start(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	ep := c.add("", s, opts.User, opts.Database)
	return map[string]any{"server": ep}, nil
}

func (c *controller) opSnapshot(req request) (map[string]any, error) {
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
	ctx := context.Background()
	if req.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(req.Timeout)*time.Millisecond)
		defer cancel()
	}
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
	return map[string]any{"snapshot": sid}, nil
}

func (c *controller) opFork(req request) (map[string]any, error) {
	c.mu.Lock()
	sn, ok := c.snapshots[req.Snapshot]
	c.mu.Unlock()
	if !ok {
		return nil, perr("unknown_id", "no snapshot %q", req.Snapshot)
	}
	ctx := context.Background()
	if req.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(req.Timeout)*time.Millisecond)
		defer cancel()
	}
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
	c.servers[fid] = &serverEntry{srv: s, user: sn.user, database: sn.database}
	c.mu.Unlock()
	return map[string]any{"server": newEndpoint(fid, s, sn.user, sn.database)}, nil
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

func (c *controller) isClosed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.closed
}

// closeAll closes forks first, then snapshots, then templates.
func (c *controller) closeAll() {
	c.mu.Lock()
	c.closed = true
	servers, snaps := c.servers, c.snapshots
	c.servers, c.snapshots = map[string]*serverEntry{}, map[string]*snapshotEntry{}
	c.mu.Unlock()
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
