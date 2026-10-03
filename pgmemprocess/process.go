// Package pgmemprocess starts the pgmem binary and controls its databases.
// It imports no embedded PostgreSQL engine. SQL uses ordinary TCP or Unix sockets.
package pgmemprocess

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Options struct {
	Binary          string // explicit path, PGMEM_BINARY, then pgmem on PATH; no runtime download
	Database        string
	User            string
	Params          []string      // postgres settings, e.g. log_statement=all
	Transport       string        // "tcp" (default) or "unix"
	SocketDir       string        // Unix socket parent; /tmp on Unix, os.TempDir() on Windows
	Log             io.Writer     // child stderr; nil discards
	StartupTimeout  time.Duration // default 30s
	ShutdownTimeout time.Duration // default 10s
}

type ProtocolError struct{ Code, Message string }

func (e *ProtocolError) Error() string { return "pgmem: " + e.Code + ": " + e.Message }

type endpoint struct {
	ID   string `json:"id"`
	DSN  string `json:"dsn"`
	Host string `json:"host"`
	Port int    `json:"port"`
}
type response struct {
	ID       uint64         `json:"id"`
	OK       bool           `json:"ok"`
	Event    string         `json:"event"`
	Protocol int            `json:"protocol"`
	Version  string         `json:"version"`
	Server   endpoint       `json:"server"`
	Snapshot string         `json:"snapshot"`
	Error    *ProtocolError `json:"error"`
}
type result struct {
	response response
	err      error
}

type Process struct {
	Template        *Server
	Version         string
	cmd             *exec.Cmd
	stdin           io.WriteCloser
	writeMu         sync.Mutex
	mu              sync.Mutex
	seq             uint64
	pending         map[uint64]chan result
	failure         error
	ready           chan result
	exited          chan struct{}
	exitErr         error
	closeOnce       sync.Once
	closeErr        error
	shutdownTimeout time.Duration
	stderr          diagnosticWriter
}

type diagnosticWriter struct {
	mu   sync.Mutex
	data []byte
}

func (w *diagnosticWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.data = append(w.data, p...)
	if len(w.data) > 16384 {
		w.data = append([]byte(nil), w.data[len(w.data)-16384:]...)
	}
	return len(p), nil
}
func (w *diagnosticWriter) text() string {
	w.mu.Lock()
	defer w.mu.Unlock()
	return string(bytes.TrimSpace(w.data))
}

// Start waits for readiness. ctx bounds startup, not the lifetime of the returned process.
func Start(ctx context.Context, opts Options) (*Process, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	binary := opts.Binary
	if binary == "" {
		binary = os.Getenv("PGMEM_BINARY")
	}
	if binary == "" {
		binary = "pgmem"
	}
	args := []string{}
	if opts.Database != "" {
		args = append(args, "-database", opts.Database)
	}
	if opts.User != "" {
		args = append(args, "-user", opts.User)
	}
	if len(opts.Params) > 0 {
		args = append(args, "-params", strings.Join(opts.Params, ","))
	}
	switch opts.Transport {
	case "", "tcp":
		if opts.SocketDir != "" {
			return nil, errors.New("pgmem: SocketDir requires Transport unix")
		}
	case "unix":
		dir := opts.SocketDir
		if dir == "" {
			dir = "/tmp"
			if runtime.GOOS == "windows" {
				dir = os.TempDir()
			}
		}
		absolute, err := filepath.Abs(dir)
		if err != nil {
			return nil, err
		}
		if len(absolute)+1+len("pgmem-")+10+1+len(".s.PGSQL.5432") > 103 {
			return nil, errors.New("pgmem: Unix socket path is too long; choose a shorter SocketDir")
		}
		args = append(args, "-socket-dir", dir)
	default:
		return nil, fmt.Errorf("pgmem: unknown transport %q", opts.Transport)
	}
	if opts.Log != nil {
		args = append(args, "-log")
	}
	cmd := exec.Command(binary, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		stdin.Close()
		return nil, err
	}
	p := &Process{cmd: cmd, stdin: stdin, pending: make(map[uint64]chan result), ready: make(chan result, 1), exited: make(chan struct{}), shutdownTimeout: opts.ShutdownTimeout}
	cmd.Stderr = &p.stderr
	if opts.Log != nil {
		cmd.Stderr = io.MultiWriter(&p.stderr, opts.Log)
	}
	if p.shutdownTimeout <= 0 {
		p.shutdownTimeout = 10 * time.Second
	}
	if err := cmd.Start(); err != nil {
		stdin.Close()
		stdout.Close()
		return nil, fmt.Errorf("pgmem: start %q (set Binary or PGMEM_BINARY): %w", binary, err)
	}
	// Read all stdout before Wait closes the pipe; failed requests are released on EOF.
	go func() {
		p.read(stdout)
		p.exitErr = cmd.Wait()
		close(p.exited)
	}()
	duration := opts.StartupTimeout
	if duration <= 0 {
		duration = 30 * time.Second
	}
	startup, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	select {
	case r := <-p.ready:
		if r.err != nil {
			p.Close()
			if detail := p.stderr.text(); detail != "" {
				return nil, fmt.Errorf("%w: %s", r.err, detail)
			}
			return nil, r.err
		}
		if r.response.Protocol != 1 {
			p.Close()
			return nil, fmt.Errorf("pgmem: protocol %d, expected 1", r.response.Protocol)
		}
		p.Version = r.response.Version
		p.Template = &Server{process: p, endpoint: r.response.Server}
		return p, nil
	case <-startup.Done():
		p.cmd.Process.Kill()
		p.Close()
		return nil, fmt.Errorf("pgmem: startup: %w", startup.Err())
	}
}

func (p *Process) read(stdout io.Reader) {
	scanner := bufio.NewScanner(stdout)
	scanner.Buffer(make([]byte, 4096), 1<<20)
	for scanner.Scan() {
		var msg response
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			p.fail(fmt.Errorf("pgmem: malformed response: %w", err))
			return
		}
		if msg.Event == "ready" {
			p.ready <- result{response: msg}
			continue
		}
		if msg.Event == "fatal" {
			p.fail(errors.New("pgmem: child reported a fatal error"))
			continue
		}
		p.mu.Lock()
		ch := p.pending[msg.ID]
		delete(p.pending, msg.ID)
		p.mu.Unlock()
		if ch != nil {
			var err error
			if !msg.OK {
				if msg.Error != nil {
					err = msg.Error
				} else {
					err = errors.New("pgmem: unsuccessful response")
				}
			}
			ch <- result{response: msg, err: err}
		}
	}
	err := scanner.Err()
	if err == nil {
		err = errors.New("pgmem: child exited or closed stdout")
	}
	p.fail(err)
}

func (p *Process) fail(err error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.failure != nil {
		return
	}
	p.failure = err
	select {
	case p.ready <- result{err: err}:
	default:
	}
	for id, ch := range p.pending {
		ch <- result{err: err}
		delete(p.pending, id)
	}
}

func (p *Process) request(ctx context.Context, op string, fields map[string]any) (response, error) {
	if err := ctx.Err(); err != nil {
		return response{}, err
	}
	p.mu.Lock()
	if p.failure != nil {
		err := p.failure
		p.mu.Unlock()
		return response{}, err
	}
	p.seq++
	id := p.seq
	ch := make(chan result, 1)
	p.pending[id] = ch
	p.mu.Unlock()
	req := make(map[string]any, len(fields)+2)
	for k, v := range fields {
		req[k] = v
	}
	req["id"], req["op"] = id, op
	if deadline, ok := ctx.Deadline(); ok {
		ms := time.Until(deadline).Milliseconds()
		if ms < 1 {
			ms = 1
		}
		req["timeout_ms"] = ms
	}
	data, err := json.Marshal(req)
	if err == nil {
		p.writeMu.Lock()
		_, err = p.stdin.Write(append(data, '\n'))
		p.writeMu.Unlock()
	}
	if err != nil {
		p.fail(err)
	}
	select {
	case r := <-ch:
		return r.response, r.err
	case <-ctx.Done():
		// A cancelled allocation may complete later. Close its resource instead of leaking a slot.
		go func() {
			r := <-ch
			if r.err != nil {
				return
			}
			cleanup, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if r.response.Server.ID != "" {
				p.request(cleanup, "close", map[string]any{"server": r.response.Server.ID})
			}
			if r.response.Snapshot != "" {
				p.request(cleanup, "close", map[string]any{"snapshot": r.response.Snapshot})
			}
		}()
		return response{}, ctx.Err()
	}
}

func (p *Process) PID() int { return p.cmd.Process.Pid }

// Close closes stdin, waits for graceful cleanup and kills a child that exceeds the deadline.
func (p *Process) Close() error {
	p.closeOnce.Do(func() {
		p.stdin.Close()
		timer := time.NewTimer(p.shutdownTimeout)
		defer timer.Stop()
		select {
		case <-p.exited:
		case <-timer.C:
			p.cmd.Process.Kill()
			<-p.exited
		}
		p.closeErr = p.exitErr
	})
	return p.closeErr
}

type Server struct {
	process   *Process
	endpoint  endpoint
	closeOnce sync.Once
	closeErr  error
}

func (s *Server) DSN() string  { return s.endpoint.DSN }
func (s *Server) Host() string { return s.endpoint.Host }
func (s *Server) Port() int    { return s.endpoint.Port }
func (s *Server) Dial(ctx context.Context, network, addr string) (net.Conn, error) {
	return (&net.Dialer{}).DialContext(ctx, network, addr)
}
func (s *Server) Snapshot(ctx context.Context, maxForks int) (*Snapshot, error) {
	msg, err := s.process.request(ctx, "snapshot", map[string]any{"server": s.endpoint.ID, "max_forks": maxForks})
	if err != nil {
		return nil, err
	}
	return &Snapshot{process: s.process, id: msg.Snapshot}, nil
}
func (s *Server) Reset(ctx context.Context) error {
	_, err := s.process.request(ctx, "reset", map[string]any{"server": s.endpoint.ID})
	return err
}
func (s *Server) Restore(ctx context.Context, snapshot *Snapshot) error {
	if snapshot == nil || snapshot.process != s.process {
		return errors.New("pgmem: snapshot belongs to another process")
	}
	_, err := s.process.request(ctx, "reset", map[string]any{"server": s.endpoint.ID, "snapshot": snapshot.id})
	return err
}
func (s *Server) Close() error {
	s.closeOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, s.closeErr = s.process.request(ctx, "close", map[string]any{"server": s.endpoint.ID})
	})
	return s.closeErr
}

type Snapshot struct {
	process *Process
	id      string
}

func (s *Snapshot) Fork(ctx context.Context) (*Server, error) {
	msg, err := s.process.request(ctx, "fork", map[string]any{"snapshot": s.id})
	if err != nil {
		return nil, err
	}
	return &Server{process: s.process, endpoint: msg.Server}, nil
}
func (s *Snapshot) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := s.process.request(ctx, "close", map[string]any{"snapshot": s.id})
	return err
}

// StartServer creates another template in this process.
func (p *Process) StartServer(ctx context.Context, database, user string) (*Server, error) {
	msg, err := p.request(ctx, "start", map[string]any{"database": database, "user": user})
	if err != nil {
		return nil, err
	}
	return &Server{process: p, endpoint: msg.Server}, nil
}

// SocketPath is the PostgreSQL socket filename for a Unix endpoint.
func (s *Server) SocketPath() string {
	if !filepath.IsAbs(s.Host()) {
		return ""
	}
	return filepath.Join(s.Host(), ".s.PGSQL."+strconv.Itoa(s.Port()))
}
