package host

import (
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shibukawa/pgmem/internal/vfs"
)

// The multi-process model.
//
// PostgreSQL is built with EXEC_BACKEND: the postmaster starts every child
// as a new program image ("postgres --forkchild=<kind> <file>") that
// re-attaches to the shared memory. pgmem gives each such process its own
// module instance on its own goroutine, and a Cluster is the "kernel" they
// share: process table and signals, System V shared memory (one file per
// segment, mapped at the same address into every instance's linear
// memory), POSIX semaphores, and the sockets that carry client
// connections to the postmaster and its backends.
//
// A process only runs guest code on its own goroutine, so the host calls
// back into a guest (a signal handler) only from that goroutine: signals
// are queued on the target's Host and delivered when the target next
// enters a host import, or when it is blocked in one (poll, semaphore
// wait, sleep), which is where a kernel would interrupt it too.

// Signal numbers (Linux/musl values, which Emscripten uses).
const (
	SIGHUP  = 1
	SIGINT  = 2
	SIGQUIT = 3
	SIGABRT = 6
	SIGKILL = 9
	SIGUSR1 = 10
	SIGUSR2 = 12
	SIGTERM = 15
	SIGCHLD = 17
	SIGURG  = 23
)

// exit status encodings of wait(2)
func exitStatus(code int32) int32  { return (code & 0xff) << 8 }
func signalStatus(sig int32) int32 { return sig & 0x7f }

// pidSeq numbers processes across every cluster of the Go process, so a
// stale pid in a copied data directory's lock file never names a live
// process of another cluster.
var pidSeq atomic.Int32

func init() { pidSeq.Store(1000) }

// Process is one PostgreSQL process: a module instance to be.
type Process struct {
	Pid  int32
	PPid int32
	Argv []string
	H    *Host
	FS   *vfs.FS
	// Done is closed when the process has exited; Status is then its
	// wait(2) status.
	Done   chan struct{}
	Status int32
	c      *Cluster
}

// Cluster is the process group of one postmaster.
type Cluster struct {
	// Exec runs p on its own goroutine: instantiate the module on p.H,
	// run main with p.Argv and call p.Exit with the wait status. Set by
	// the engine, which owns the module factory.
	Exec func(p *Process)
	Log  func(format string, args ...any)
	// ShmBase is the linear-memory address the shared segments are mapped
	// at (the same in every process); the heap of every process is capped
	// below it. Segments are placed from ShmBase up to ShmLimit.
	ShmBase, ShmLimit uint32

	mu         sync.Mutex
	procs      map[int32]*Process
	exited     []*Process // waitpid queue
	postmaster *Process
	dead       bool // the postmaster has exited

	// shared memory
	segs      map[int32]*Segment
	segByKey  map[int32]*Segment
	nextSegID int32
	shmFree   []span // free address ranges in [ShmBase, ShmLimit)

	// semaphores, by sem_t address
	sems map[uint32]*sema

	// listen queue: client connections handed to the postmaster
	acceptQ       []*ConnSock
	acceptWaiters map[*Host]struct{}
}

type span struct{ addr, size uint32 }

// NewCluster creates an empty cluster. shmBase/shmLimit bound the shared
// memory window in every process's linear memory.
func NewCluster(shmBase, shmLimit uint32) *Cluster {
	return &Cluster{
		ShmBase:       shmBase,
		ShmLimit:      shmLimit,
		procs:         map[int32]*Process{},
		segs:          map[int32]*Segment{},
		segByKey:      map[int32]*Segment{},
		nextSegID:     1,
		shmFree:       []span{{shmBase, shmLimit - shmBase}},
		sems:          map[uint32]*sema{},
		acceptWaiters: map[*Host]struct{}{},
	}
}

func (c *Cluster) logf(format string, args ...any) {
	if c.Log != nil {
		c.Log(format, args...)
	}
}

// Start runs the postmaster: the first process, whose file descriptor
// table is fs's. It returns once the goroutine is started; the caller
// waits for readiness by other means (the pid file).
func (c *Cluster) Start(fs *vfs.FS, argv []string, env []string) *Process {
	p := c.newProcess(0, fs, argv, env)
	c.mu.Lock()
	c.postmaster = p
	c.mu.Unlock()
	c.Exec(p)
	return p
}

func (c *Cluster) newProcess(ppid int32, fs *vfs.FS, argv, env []string) *Process {
	h := New(fs)
	h.Env = env
	h.Log = c.Log
	h.Sys = c
	h.HeapMax = c.ShmBase
	p := &Process{Pid: pidSeq.Add(1), PPid: ppid, Argv: argv, H: h, FS: fs, Done: make(chan struct{}), c: c}
	h.Pid = p.Pid
	h.proc = p
	c.mu.Lock()
	c.procs[p.Pid] = p
	c.mu.Unlock()
	return p
}

// Postmaster returns the first process.
func (c *Cluster) Postmaster() *Process {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.postmaster
}

// Dead reports whether the postmaster has exited.
func (c *Cluster) Dead() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.dead
}

// Processes returns the live processes.
func (c *Cluster) Processes() []*Process {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]*Process, 0, len(c.procs))
	for _, p := range c.procs {
		out = append(out, p)
	}
	return out
}

// spawn is the pgmem_spawn import: start "postgres <forkarg> <paramfile>"
// as a child of parent, inheriting its open descriptors.
func (c *Cluster) spawn(parent *Host, forkarg, paramfile string) int32 {
	argv := []string{parent.proc.Argv[0], forkarg, paramfile}
	fs := parent.FS.ForkFDs()
	p := c.newProcess(parent.Pid, fs, argv, parent.Env)
	c.Exec(p)
	return p.Pid
}

// Exit records that p has finished: its descriptors close (a pipe end
// held only by it reads as EOF elsewhere, a client socket it owned is
// closed), its shared memory attachments end, its parent gets SIGCHLD
// and can waitpid it. The postmaster's exit wakes every process, whose
// postmaster-death pipe now reads as closed.
func (p *Process) Exit(status int32) {
	c := p.c
	p.FS.CloseAll()
	c.detachAll(p.H)
	c.mu.Lock()
	p.Status = status
	delete(c.procs, p.Pid)
	c.exited = append(c.exited, p)
	parent := c.procs[p.PPid]
	isPM := p == c.postmaster
	if isPM {
		c.dead = true
	}
	var all []*Process
	if isPM {
		for _, q := range c.procs {
			all = append(all, q)
		}
	}
	c.mu.Unlock()
	close(p.Done)
	if parent != nil {
		parent.H.Signal(SIGCHLD)
	}
	for _, q := range all {
		q.H.Wake()
	}
}

// Kill implements kill(2) for pid from h. A signal to oneself runs the
// handler right away, as the kernel would; SIGKILL ends the target at its
// next host call.
func (c *Cluster) Kill(from *Host, pid int32, sig int32) int32 {
	if pid < 0 {
		// a process group: every process is its own (setsid), so this is
		// the process itself
		pid = -pid
	}
	if pid == 0 {
		return -int32(vfs.ENOSYS)
	}
	c.mu.Lock()
	p := c.procs[pid]
	c.mu.Unlock()
	if p == nil {
		return -int32(vfs.ESRCH)
	}
	if sig == 0 {
		return 0
	}
	if sig == SIGKILL {
		p.H.killed.Store(true)
		p.H.Wake()
		return 0
	}
	if p.H == from {
		if err := from.Guest.Raise(sig); err != nil {
			panic(err)
		}
		return 0
	}
	p.H.Signal(sig)
	return 0
}

// KillAll sends sig to every live process.
func (c *Cluster) KillAll(sig int32) {
	for _, p := range c.Processes() {
		c.Kill(nil, p.Pid, sig)
	}
}

// Waitpid implements waitpid(2) with WNOHANG (what PostgreSQL uses).
func (c *Cluster) Waitpid(from *Host, pid int32, options int32) (int32, int32) {
	const wnohang = 1
	for {
		c.mu.Lock()
		for i, p := range c.exited {
			if p.PPid != from.Pid {
				continue
			}
			if pid == -1 || pid == p.Pid {
				c.exited = append(c.exited[:i], c.exited[i+1:]...)
				c.mu.Unlock()
				return p.Pid, p.Status
			}
		}
		anyChild := false
		for _, p := range c.procs {
			if p.PPid == from.Pid && (pid == -1 || pid == p.Pid) {
				anyChild = true
				break
			}
		}
		c.mu.Unlock()
		if !anyChild {
			return -int32(vfs.ECHILD), 0
		}
		if options&wnohang != 0 {
			return 0, 0
		}
		from.wait(-1)
		if from.interrupted() {
			return -int32(vfs.EINTR), 0
		}
	}
}

// ---- signals and blocking on the Host side ----

// Signal queues sig for delivery to h's process and wakes it if it is
// blocked in a host call.
func (h *Host) Signal(sig int32) {
	if sig <= 0 || sig >= 32 {
		return
	}
	h.pending.Or(1 << uint(sig))
	h.Wake()
}

// Wake makes a blocked host call on h's goroutine return and look around.
func (h *Host) Wake() {
	select {
	case h.wake <- struct{}{}:
	default:
	}
}

// wait blocks h's goroutine until it is woken or d passes (d < 0: no
// timeout). It reports whether the timeout expired.
func (h *Host) wait(d time.Duration) bool {
	if d < 0 {
		<-h.wake
		return false
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-h.wake:
		return false
	case <-t.C:
		return true
	}
}

// interrupted reports whether a signal or timer is waiting for delivery,
// or the process was killed; if so the pending work is delivered first
// (a killed process does not return from here).
func (h *Host) interrupted() bool {
	if h.killed.Load() {
		panic(&ExitError{Code: 128 + SIGKILL})
	}
	if h.pending.Load() == 0 && h.timerDue.Load() == 0 {
		return false
	}
	h.deliver()
	return true
}

// enter runs at every host import: it is where the process notices
// signals, expired timers and its own death.
func (h *Host) enter() {
	if h.inDeliver {
		return
	}
	if h.killed.Load() {
		panic(&ExitError{Code: 128 + SIGKILL})
	}
	if h.pending.Load() != 0 || h.timerDue.Load() != 0 {
		h.deliver()
	}
}

// deliver runs the guest's handlers for every queued signal and expired
// timer. A handler that exits the process unwinds through here.
func (h *Host) deliver() {
	if h.inDeliver || h.Guest == nil {
		return
	}
	h.inDeliver = true
	defer func() { h.inDeliver = false }()
	if h.timerDue.Swap(0) != 0 {
		if _, err := h.checkTimers(); err != nil {
			panic(err)
		}
	}
	for {
		p := h.pending.Swap(0)
		if p == 0 {
			return
		}
		for sig := int32(1); sig < 32; sig++ {
			if p&(1<<uint(sig)) == 0 {
				continue
			}
			if err := h.Guest.Raise(sig); err != nil {
				panic(err)
			}
		}
	}
}

// ---- shared memory ----

// Segment is one System V shared memory segment: a file mapped at the
// same address into every attached process.
type Segment struct {
	id     int32
	key    int32
	size   uint32 // rounded up to 64 KiB
	addr   uint32
	f      *os.File
	nattch int
	rmid   bool
}

const (
	ipcCreat  = 0o1000
	ipcExcl   = 0o2000
	ipcRMID   = 0
	ipcSet    = 1
	ipcStat   = 2
	shmAlign  = 65536
	shmMaxKey = 1<<31 - 1
)

func (c *Cluster) shmget(key int32, size uint32, flags int32) int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	if key != 0 {
		if seg := c.segByKey[key]; seg != nil {
			if flags&(ipcCreat|ipcExcl) == ipcCreat|ipcExcl {
				return -int32(vfs.EEXIST)
			}
			if size > seg.size {
				return -int32(vfs.EINVAL)
			}
			return seg.id
		}
		if flags&ipcCreat == 0 {
			return -int32(vfs.ENOENT)
		}
	}
	rounded := (size + shmAlign - 1) &^ (shmAlign - 1)
	if rounded == 0 {
		rounded = shmAlign
	}
	addr, ok := c.allocShm(rounded)
	if !ok {
		return -int32(vfs.ENOMEM)
	}
	f, err := newSegmentFile(rounded)
	if err != nil {
		c.logf("shmget: %v", err)
		c.freeShm(addr, rounded)
		return -int32(vfs.ENOMEM)
	}
	seg := &Segment{id: c.nextSegID, key: key, size: rounded, addr: addr, f: f}
	c.nextSegID++
	c.segs[seg.id] = seg
	if key != 0 {
		c.segByKey[key] = seg
	}
	return seg.id
}

func (c *Cluster) allocShm(size uint32) (uint32, bool) {
	for i, s := range c.shmFree {
		if s.size >= size {
			addr := s.addr
			if s.size == size {
				c.shmFree = append(c.shmFree[:i], c.shmFree[i+1:]...)
			} else {
				c.shmFree[i] = span{s.addr + size, s.size - size}
			}
			return addr, true
		}
	}
	return 0, false
}

func (c *Cluster) freeShm(addr, size uint32) {
	// keep the list sorted and merged
	i := 0
	for i < len(c.shmFree) && c.shmFree[i].addr < addr {
		i++
	}
	c.shmFree = append(c.shmFree, span{})
	copy(c.shmFree[i+1:], c.shmFree[i:])
	c.shmFree[i] = span{addr, size}
	for j := 0; j+1 < len(c.shmFree); {
		a, b := c.shmFree[j], c.shmFree[j+1]
		if a.addr+a.size == b.addr {
			c.shmFree[j] = span{a.addr, a.size + b.size}
			c.shmFree = append(c.shmFree[:j+1], c.shmFree[j+2:]...)
			continue
		}
		j++
	}
}

func (c *Cluster) shmat(h *Host, id int32, addr uint32) int32 {
	c.mu.Lock()
	seg := c.segs[id]
	if seg == nil || seg.rmid && seg.nattch == 0 {
		c.mu.Unlock()
		return -int32(vfs.EINVAL)
	}
	if addr != 0 && addr != seg.addr {
		c.mu.Unlock()
		return -int32(vfs.EINVAL)
	}
	if h.attached[seg.addr] != nil {
		c.mu.Unlock()
		return int32(seg.addr) // already attached here
	}
	seg.nattch++
	c.mu.Unlock()
	if err := h.Guest.MapShared(seg.addr, seg.f, seg.size); err != nil {
		c.mu.Lock()
		seg.nattch--
		c.mu.Unlock()
		h.logf("shmat: %v", err)
		return -int32(vfs.ENOMEM)
	}
	if h.attached == nil {
		h.attached = map[uint32]*Segment{}
	}
	h.attached[seg.addr] = seg
	return int32(seg.addr)
}

func (c *Cluster) shmdt(h *Host, addr uint32) int32 {
	seg := h.attached[addr]
	if seg == nil {
		return -int32(vfs.EINVAL)
	}
	delete(h.attached, addr)
	if err := h.Guest.UnmapShared(seg.addr, seg.size); err != nil {
		h.logf("shmdt: %v", err)
	}
	c.mu.Lock()
	seg.nattch--
	c.maybeDestroy(seg)
	c.mu.Unlock()
	return 0
}

// detachAll ends every attachment of an exiting process; its memory is
// released whole by the instance, so nothing is unmapped here.
func (c *Cluster) detachAll(h *Host) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for addr, seg := range h.attached {
		delete(h.attached, addr)
		seg.nattch--
		c.maybeDestroy(seg)
	}
}

// maybeDestroy frees a removed segment nobody is attached to. Locked.
func (c *Cluster) maybeDestroy(seg *Segment) {
	if !seg.rmid || seg.nattch > 0 || c.segs[seg.id] != seg {
		return
	}
	delete(c.segs, seg.id)
	if seg.key != 0 && c.segByKey[seg.key] == seg {
		delete(c.segByKey, seg.key)
	}
	seg.f.Close()
	c.freeShm(seg.addr, seg.size)
}

func (c *Cluster) shmctl(id int32, cmd int32) (segsz, nattch, ret int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	seg := c.segs[id]
	if seg == nil {
		return 0, 0, -int32(vfs.EINVAL)
	}
	switch cmd {
	case ipcRMID:
		seg.rmid = true
		if seg.key != 0 && c.segByKey[seg.key] == seg {
			delete(c.segByKey, seg.key) // the key can be reused at once
		}
		c.maybeDestroy(seg)
		return 0, 0, 0
	case ipcStat:
		return int32(seg.size), int32(seg.nattch), 0
	case ipcSet:
		return 0, 0, 0
	}
	return 0, 0, -int32(vfs.EINVAL)
}

// Close releases every segment (the cluster is gone).
func (c *Cluster) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for id, seg := range c.segs {
		seg.f.Close()
		delete(c.segs, id)
	}
	c.segByKey = map[int32]*Segment{}
}

// ---- semaphores ----

// sema is a counting semaphore in shared memory (identified by the sem_t
// address, the same in every process) whose count lives here.
type sema struct {
	mu      sync.Mutex
	count   int32
	waiters []chan struct{}
}

const (
	semInit = iota
	semDestroy
	semWait
	semTryWait
	semPost
)

func (c *Cluster) sem(h *Host, op int32, addr uint32, arg int32) int32 {
	c.mu.Lock()
	s := c.sems[addr]
	if s == nil && op != semInit && op != semDestroy {
		c.mu.Unlock()
		return -int32(vfs.EINVAL)
	}
	switch op {
	case semInit:
		c.sems[addr] = &sema{count: arg}
		c.mu.Unlock()
		return 0
	case semDestroy:
		delete(c.sems, addr)
		c.mu.Unlock()
		return 0
	}
	c.mu.Unlock()
	switch op {
	case semPost:
		s.mu.Lock()
		s.count++
		if len(s.waiters) > 0 {
			w := s.waiters[0]
			s.waiters = s.waiters[1:]
			close(w)
		}
		s.mu.Unlock()
		return 0
	case semTryWait:
		s.mu.Lock()
		defer s.mu.Unlock()
		if s.count > 0 {
			s.count--
			return 0
		}
		return -int32(vfs.EAGAIN)
	case semWait:
		for {
			s.mu.Lock()
			if s.count > 0 {
				s.count--
				s.mu.Unlock()
				return 0
			}
			ch := make(chan struct{})
			s.waiters = append(s.waiters, ch)
			s.mu.Unlock()
			woken := false
			select {
			case <-ch:
				woken = true
			case <-h.wake:
			}
			if !woken {
				// leave the queue; if a post picked us in the meantime,
				// the count it handed us is still in the semaphore
				s.mu.Lock()
				for i, w := range s.waiters {
					if w == ch {
						s.waiters = append(s.waiters[:i], s.waiters[i+1:]...)
						break
					}
				}
				s.mu.Unlock()
				if h.interrupted() {
					return -int32(vfs.EINTR)
				}
			}
		}
	}
	return -int32(vfs.EINVAL)
}

// ---- sockets ----

// Listener is a listening socket: every one of a cluster shares one
// accept queue, because pgmem hands connections to the cluster, not to
// an address.
type Listener struct {
	path string
}

// ConnSock is an accepted client connection as the backend sees it: bytes
// the client sent wait in buf; what the backend sends goes to w.
type ConnSock struct {
	mu      sync.Mutex
	cond    *sync.Cond
	buf     []byte
	eof     bool
	closed  bool
	muted   bool // drop what the backend sends (a closing server)
	w       io.Writer
	waiters map[*Host]struct{}
	// OnClose runs when the backend side has closed its last descriptor
	// (the process exited or ended the session).
	OnClose func()
	// Listen receives the backend's LISTEN set changes at commit (see
	// Host.Listen); pgmem replays them when the session gets a new
	// backend after a Restore.
	Listen func(channel string, op int)
}

// ClientSock returns the client connection a backend process serves, if
// it has one among its descriptors.
func (h *Host) ClientSock() *ConnSock {
	var found *ConnSock
	h.FS.Sockets(func(sock any) {
		if s, ok := sock.(*ConnSock); ok && found == nil {
			found = s
		}
	})
	return found
}

// pushLimit bounds how far a client can run ahead of its backend.
const pushLimit = 4 << 20

// NewConnSock wraps the client side of a connection: w receives what the
// backend sends.
func NewConnSock(w io.Writer) *ConnSock {
	s := &ConnSock{w: w, waiters: map[*Host]struct{}{}}
	s.cond = sync.NewCond(&s.mu)
	return s
}

// Push delivers bytes from the client. It blocks while the backend is
// far behind, and returns false once the backend side is closed.
func (s *ConnSock) Push(b []byte) bool {
	s.mu.Lock()
	for len(s.buf) > pushLimit && !s.closed {
		s.cond.Wait()
	}
	if s.closed {
		s.mu.Unlock()
		return false
	}
	s.buf = append(s.buf, b...)
	ws := s.snapshotWaiters()
	s.mu.Unlock()
	for _, h := range ws {
		h.Wake()
	}
	return true
}

// EOF marks the client side closed.
func (s *ConnSock) EOF() {
	s.mu.Lock()
	s.eof = true
	ws := s.snapshotWaiters()
	s.mu.Unlock()
	for _, h := range ws {
		h.Wake()
	}
}

func (s *ConnSock) snapshotWaiters() []*Host {
	ws := make([]*Host, 0, len(s.waiters))
	for h := range s.waiters {
		ws = append(ws, h)
	}
	return ws
}

func (s *ConnSock) readable() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.buf) > 0 || s.eof
}

func (s *ConnSock) recv(dst []byte) int32 {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.buf) == 0 {
		if s.eof {
			return 0
		}
		return -int32(vfs.EAGAIN)
	}
	n := copy(dst, s.buf)
	s.buf = s.buf[n:]
	if len(s.buf) == 0 {
		s.buf = nil
	}
	s.cond.Broadcast()
	return int32(n)
}

// Mute drops everything the backend sends from now on: a server being
// closed keeps its client sockets open and quiet (see pgmem.Server.Close).
func (s *ConnSock) Mute() {
	s.mu.Lock()
	s.muted = true
	s.mu.Unlock()
}

// Muted reports whether Mute was called.
func (s *ConnSock) Muted() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.muted
}

func (s *ConnSock) send(b []byte) int32 {
	s.mu.Lock()
	w, closed, muted := s.w, s.closed, s.muted
	s.mu.Unlock()
	if closed || w == nil {
		return -int32(vfs.EPIPE)
	}
	if muted {
		return int32(len(b))
	}
	if _, err := w.Write(b); err != nil {
		return -int32(vfs.EPIPE)
	}
	return int32(len(b))
}

// backendClosed is the vfs close callback: the last descriptor went away.
func (s *ConnSock) backendClosed() {
	s.mu.Lock()
	s.closed = true
	s.cond.Broadcast()
	cb := s.OnClose
	s.mu.Unlock()
	if cb != nil {
		cb()
	}
}

func (s *ConnSock) addWaiter(h *Host) {
	s.mu.Lock()
	s.waiters[h] = struct{}{}
	s.mu.Unlock()
}

func (s *ConnSock) removeWaiter(h *Host) {
	s.mu.Lock()
	delete(s.waiters, h)
	s.mu.Unlock()
}

// Connect hands a client connection to the postmaster's accept queue.
func (c *Cluster) Connect(s *ConnSock) error {
	c.mu.Lock()
	if c.dead || c.postmaster == nil {
		c.mu.Unlock()
		return fmt.Errorf("pgmem: the postmaster is not running")
	}
	c.acceptQ = append(c.acceptQ, s)
	ws := make([]*Host, 0, len(c.acceptWaiters))
	for h := range c.acceptWaiters {
		ws = append(ws, h)
	}
	c.mu.Unlock()
	for _, h := range ws {
		h.Wake()
	}
	return nil
}

func (c *Cluster) acceptPending() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.acceptQ) > 0
}

func (c *Cluster) popAccept() *ConnSock {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.acceptQ) == 0 {
		return nil
	}
	s := c.acceptQ[0]
	c.acceptQ = c.acceptQ[1:]
	return s
}

func (c *Cluster) setAcceptWaiter(h *Host, on bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if on {
		c.acceptWaiters[h] = struct{}{}
	} else {
		delete(c.acceptWaiters, h)
	}
}

// accept implements accept(2) on a listening socket: it blocks until a
// connection is queued (or a signal arrives) and returns a descriptor
// on it. The peer address is a nameless AF_UNIX one.
func (h *Host) accept(m Memory, fd int32, addrP, lenP uint32) uint64 {
	c := h.Sys
	sock, err := h.FS.Socket(fd)
	if err != vfs.OK {
		return errno(err)
	}
	if _, ok := sock.(*Listener); !ok {
		return errno(vfs.EINVAL)
	}
	for {
		s := c.popAccept()
		if s != nil {
			nfd := h.FS.NewSocket(s, s.backendClosed)
			if addrP != 0 && lenP != 0 {
				n := rdU32(m, lenP)
				var sa [3]byte
				sa[0] = 1 // AF_UNIX, little endian u16
				if n > 3 {
					n = 3
				}
				m.Write(addrP, sa[:n])
				wrU32(m, lenP, 3)
			}
			return ret32(nfd)
		}
		c.setAcceptWaiter(h, true)
		if !c.acceptPending() {
			h.wait(-1)
		}
		c.setAcceptWaiter(h, false)
		if h.interrupted() {
			return errno(vfs.EINTR)
		}
	}
}

// ---- poll ----

const (
	pollIn   = 0x1
	pollOut  = 0x4
	pollErr  = 0x8
	pollHup  = 0x10
	pollNval = 0x20
)

// poll implements poll(2) over pipes, sockets and files. It is the
// blocking point of a PostgreSQL process, so signals and timers are
// delivered here and reported as EINTR, which WaitEventSetWait retries.
func (h *Host) poll(m Memory, fdsP uint32, nfds int32, timeout int32) uint64 {
	c := h.Sys
	var deadline time.Time
	if timeout > 0 {
		deadline = time.Now().Add(time.Duration(timeout) * time.Millisecond)
	}
	var socks []*ConnSock
	listening := false
	for {
		ready := int32(0)
		socks = socks[:0]
		listening = false
		for i := int32(0); i < nfds; i++ {
			p := fdsP + uint32(i)*8
			fd := rdI32(m, p)
			events := int16(rdU32(m, p+4) & 0xffff)
			var revents int16
			if fd < 0 {
				m.Write(p+6, []byte{0, 0})
				continue
			}
			pi, err := h.FS.Poll(fd)
			if err != vfs.OK {
				revents = pollNval
			} else {
				switch pi.Kind {
				case vfs.KindPipe:
					if pi.Readable {
						revents |= pollIn
					}
					if pi.Hup {
						revents |= pollHup
					}
					if events&pollOut != 0 {
						revents |= pollOut
					}
				case vfs.KindSocket:
					switch s := pi.Sock.(type) {
					case *Listener:
						listening = true
						if c.acceptPending() {
							revents |= pollIn
						}
					case *ConnSock:
						socks = append(socks, s)
						if s.readable() {
							revents |= pollIn
						}
						if events&pollOut != 0 {
							revents |= pollOut
						}
					}
				default:
					revents = events & (pollIn | pollOut)
				}
				revents &= events | pollHup | pollErr | pollNval
			}
			m.Write(p+6, []byte{byte(revents), byte(revents >> 8)})
			if revents != 0 {
				ready++
			}
		}
		if ready > 0 {
			return ret32(ready)
		}
		if timeout == 0 {
			return 0
		}
		d := time.Duration(-1)
		if timeout > 0 {
			d = time.Until(deadline)
			if d <= 0 {
				return 0
			}
		}
		for _, s := range socks {
			s.addWaiter(h)
		}
		if listening {
			c.setAcceptWaiter(h, true)
		}
		// something may have become ready between the scan and the
		// registration; a wake is cheap, a missed one is a hang
		recheck := listening && c.acceptPending()
		for _, s := range socks {
			if s.readable() {
				recheck = true
			}
		}
		if !recheck {
			h.wait(d)
		}
		for _, s := range socks {
			s.removeWaiter(h)
		}
		if listening {
			c.setAcceptWaiter(h, false)
		}
		if h.interrupted() {
			return errno(vfs.EINTR)
		}
	}
}

// usleep sleeps for us microseconds, returning early on a signal.
func (h *Host) usleep(us int32) {
	if us <= 0 {
		return
	}
	h.wait(time.Duration(us) * time.Microsecond)
	h.interrupted()
}
