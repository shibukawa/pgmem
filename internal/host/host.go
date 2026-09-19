// Package host implements the Emscripten "env" and WASI imports that the
// PostgreSQL wasm module needs, on top of the in-memory vfs. It is engine
// agnostic: every import is a function over a Memory and a raw argument
// slice, collected in a table that a wazero (or wasm2go) binding walks.
package host

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/crc32"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/shibukawa/pgmem/internal/vfs"
	"golang.org/x/crypto/ripemd160"
)

// Memory is the linear memory of the guest.
type Memory interface {
	Read(off, n uint32) ([]byte, bool)
	Write(off uint32, b []byte) bool
	Size() uint32
	// Grow adds delta pages (64 KiB) and returns the previous size in pages.
	Grow(delta uint32) (uint32, bool)
}

// Guest exposes the few guest exports the host needs to call back into.
type Guest interface {
	// CallSighandler invokes the guest function pointer fp with sig.
	CallSighandler(fp, sig int32) error
	// Memalign calls emscripten_builtin_memalign(align, size).
	Memalign(align, size uint32) (uint32, error)
	// Timeout calls _emscripten_timeout(which, now).
	Timeout(which int32, now float64) error
	// Malloc calls the guest's malloc.
	Malloc(size uint32) (uint32, error)
	// Raise runs the guest's handler for sig (pgmem_raise), honouring the
	// process's signal mask.
	Raise(sig int32) error
	// MapShared maps f (size bytes) at offset off of the linear memory,
	// shared with every other instance that maps it; UnmapShared puts
	// private zero pages back.
	MapShared(off uint32, f *os.File, size uint32) error
	UnmapShared(off, size uint32) error
}

// ExitError is panicked by exit-like imports to unwind the guest stack.
type ExitError struct {
	Code int32
	// Live is true for emscripten_exit_with_live_runtime.
	Live bool
}

func (e *ExitError) Error() string {
	if e.Live {
		return "exit with live runtime"
	}
	return fmt.Sprintf("exit(%d)", e.Code)
}

// AbortError is panicked by abort()/assert failures.
type AbortError struct{ Msg string }

func (e *AbortError) Error() string { return "abort: " + e.Msg }

// Host holds the state behind the imports.
type Host struct {
	FS    *vfs.FS
	Env   []string // "KEY=VALUE"
	Guest Guest

	// Run is called by pgmem_run to execute "postgres ..." as a child
	// (initdb support). It returns the exit code.
	Run func(cmd, stdinPath, stdoutPath string) int
	// Log receives diagnostics from the host itself.
	Log func(format string, args ...any)
	// Listen is called by pgmem_listen when the session's LISTEN set
	// changes at commit: op 1 = LISTEN channel, 0 = UNLISTEN channel,
	// 2 = UNLISTEN * (channel is empty).
	Listen func(channel string, op int)

	// Process identity in the multi-process model (see cluster.go). Sys
	// is nil for a standalone process (initdb's children, the setup
	// child), whose process calls are answered with ENOSYS or served
	// locally (single.go).
	Pid  int32
	Sys  *Cluster
	proc *Process
	// HeapMax caps the heap below the shared memory window (0 = no cap).
	HeapMax uint32
	// pending is the bitmask of queued signals, timerDue that of itimers
	// whose deadline passed; wake is how another goroutine interrupts a
	// blocking host call on this process. killed ends the process at its
	// next host call (SIGKILL).
	pending   atomic.Uint32
	timerDue  atomic.Uint32
	wake      chan struct{}
	inDeliver bool
	killed    atomic.Bool
	attached  map[uint32]*Segment // shared memory segments, by address
	// single-user mode: shared memory and semaphores served locally
	localSegs    map[int32]*localSeg
	nextLocalSeg int32
	localSems    map[uint32]int32

	timers   [3]time.Time   // itimer deadlines (zero = unset)
	timerFns [3]*time.Timer // fire timerDue at the deadline
	inTimer  bool           // a timer handler is running (no re-entry)
	start    time.Time
	epochSec int64

	// hashes holds the guest's live cryptohash contexts (see
	// wasm/pgmem_cryptohash.inc), keyed by handle.
	hashes   map[int32]hash.Hash
	hashType map[int32]int32
	hashFree map[int32][]hash.Hash // freed contexts per type, reused by create
	nextHash int32

	// ciphers holds pgcrypto's live cipher contexts (see
	// wasm/pgmem_pgcrypto_openssl.inc and crypto.go), keyed by handle.
	ciphers    map[int32]*hostCipher
	nextCipher int32

	// zstreams holds pgcrypto's live deflate/inflate contexts (see
	// wasm/pgmem_pgcrypto_compress.inc and crypto.go), keyed by handle.
	zstreams    map[int32]*zstream
	nextZstream int32
}

// newHash maps a hash type to a Go hash: 0-5 are PostgreSQL's
// pg_cryptohash_type (src/common/cryptohash.c), 6 and up are extra digests
// pgcrypto's digest()/hmac() accept (wasm/pgmem_pgcrypto_openssl.inc).
func newHash(typ int32) hash.Hash {
	switch typ {
	case 0:
		return md5.New()
	case 1:
		return sha1.New()
	case 2:
		return sha256.New224()
	case 3:
		return sha256.New()
	case 4:
		return sha512.New384()
	case 5:
		return sha512.New()
	case 6:
		return ripemd160.New()
	case 7:
		return sha3.New224()
	case 8:
		return sha3.New256()
	case 9:
		return sha3.New384()
	case 10:
		return sha3.New512()
	}
	return nil
}

var castagnoli = crc32.MakeTable(crc32.Castagnoli)

// New creates a host over fs.
func New(fs *vfs.FS) *Host {
	return &Host{FS: fs, start: time.Now(), wake: make(chan struct{}, 1)}
}

// Every import first lets the process notice queued signals, expired
// timers and its own death (Host.enter): a host call is where a kernel
// would interrupt it.
func init() {
	for i := range table {
		call := table[i].Call
		table[i].Call = func(h *Host, m Memory, a []uint64) uint64 {
			h.enter()
			return call(h, m, a)
		}
	}
}

// Fn is one host import.
type Fn struct {
	Module  string
	Name    string
	Params  string // one char per param: i=i32 j=i64 f=f32 d=f64
	Results string
	Call    func(h *Host, m Memory, args []uint64) uint64
}

// Table returns every import the module may need. Bindings register the
// subset actually imported by a given module.
func Table() []Fn {
	return table
}

// Lookup finds an import by module and name.
func Lookup(module, name string) (Fn, bool) {
	for _, f := range table {
		if f.Module == module && f.Name == name {
			return f, true
		}
	}
	return Fn{}, false
}

// ---- argument helpers ----

func i32(v uint64) int32   { return int32(uint32(v)) }
func u32(v uint64) uint32  { return uint32(v) }
func i64(v uint64) int64   { return int64(v) }
func f64(v uint64) float64 { return math.Float64frombits(v) }
func ret32(v int32) uint64 { return uint64(uint32(v)) }
func retf64(v float64) uint64 {
	return math.Float64bits(v)
}
func errno(e vfs.Errno) uint64 { return ret32(-int32(e)) }

func (h *Host) str(m Memory, ptr uint32) string {
	if ptr == 0 {
		return ""
	}
	var sb strings.Builder
	for {
		b, ok := m.Read(ptr, 256)
		if !ok {
			// near the end of memory: read byte by byte
			b, ok = m.Read(ptr, 1)
			if !ok {
				return sb.String()
			}
		}
		for i, c := range b {
			if c == 0 {
				sb.Write(b[:i])
				return sb.String()
			}
		}
		sb.Write(b)
		ptr += uint32(len(b))
	}
}

func rdU32(m Memory, off uint32) uint32 {
	b, ok := m.Read(off, 4)
	if !ok {
		return 0
	}
	return binary.LittleEndian.Uint32(b)
}

func wrU32(m Memory, off uint32, v uint32) {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], v)
	m.Write(off, b[:])
}

func wrU64(m Memory, off uint32, v uint64) {
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], v)
	m.Write(off, b[:])
}

func rdI32(m Memory, off uint32) int32 { return int32(rdU32(m, off)) }
func wrI32(m Memory, off uint32, v int32) {
	wrU32(m, off, uint32(v))
}

func (h *Host) logf(format string, args ...any) {
	if h.Log != nil {
		h.Log(format, args...)
	}
}

// ---- stat / dirent encoding (Emscripten wasm32 layout) ----

func writeStat(m Memory, buf uint32, st vfs.Stat) {
	var b [96]byte
	le := binary.LittleEndian
	le.PutUint32(b[0:], 1)        // st_dev
	le.PutUint32(b[4:], st.Mode)  // st_mode
	le.PutUint32(b[8:], st.Nlink) // st_nlink
	le.PutUint32(b[12:], 0)       // st_uid
	le.PutUint32(b[16:], 0)       // st_gid
	le.PutUint32(b[20:], 0)       // st_rdev
	le.PutUint64(b[24:], uint64(st.Size))
	le.PutUint32(b[32:], 4096)                      // st_blksize
	le.PutUint32(b[36:], uint32((st.Size+511)/512)) // st_blocks
	sec := uint64(st.Mtime.Unix())
	nsec := uint32(st.Mtime.Nanosecond())
	for _, off := range []int{40, 56, 72} {
		le.PutUint64(b[off:], sec)
		le.PutUint32(b[off+8:], nsec)
	}
	le.PutUint64(b[88:], st.Ino)
	m.Write(buf, b[:])
}

const direntSize = 280

// ---- the table ----

var table = []Fn{
	// ===== pgmem bridge =====
	// recv/send on a socket fd (a ConnSock of the cluster).
	{"env", "pgmem_sock_recv", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, ptr, n := i32(a[0]), u32(a[1]), u32(a[2])
		dst, ok := m.Read(ptr, n)
		if !ok {
			return errno(vfs.EFAULT)
		}
		if h.Sys == nil {
			return errno(vfs.ENOTSOCK)
		}
		s, err := h.connSock(fd)
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(s.recv(dst))
	}},
	{"env", "pgmem_sock_send", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, ptr, n := i32(a[0]), u32(a[1]), u32(a[2])
		b, ok := m.Read(ptr, n)
		if !ok {
			return errno(vfs.EFAULT)
		}
		if h.Sys == nil {
			return errno(vfs.ENOTSOCK)
		}
		s, err := h.connSock(fd)
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(s.send(b))
	}},
	{"env", "pgmem_getpid", "", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Pid == 0 {
			return ret32(42) // what Emscripten's getpid says
		}
		return ret32(h.Pid)
	}},
	{"env", "pgmem_kill", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		pid, sig := i32(a[0]), i32(a[1])
		if h.Sys == nil {
			// standalone: only oneself exists (SetLatch sends SIGURG)
			if pid != 42 && pid != h.Pid {
				return errno(vfs.ESRCH)
			}
			if sig != 0 && h.Guest != nil {
				if err := h.Guest.Raise(sig); err != nil {
					panic(err)
				}
			}
			return 0
		}
		return ret32(h.Sys.Kill(h, pid, sig))
	}},
	{"env", "pgmem_waitpid", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return errno(vfs.ECHILD)
		}
		pid, status := h.Sys.Waitpid(h, i32(a[0]), i32(a[2]))
		if pid > 0 && u32(a[1]) != 0 {
			wrI32(m, u32(a[1]), status)
		}
		return ret32(pid)
	}},
	{"env", "pgmem_usleep", "i", "", func(h *Host, m Memory, a []uint64) uint64 {
		us := i32(a[0])
		if h.Sys == nil {
			h.sleep((us + 999) / 1000)
			return 0
		}
		h.usleep(us)
		return 0
	}},
	{"env", "pgmem_spawn", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return errno(vfs.ENOSYS)
		}
		return ret32(h.Sys.spawn(h, h.str(m, u32(a[0])), h.str(m, u32(a[1]))))
	}},
	{"env", "pgmem_shmget", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return ret32(h.localShmget(m, i32(a[0]), u32(a[1]), i32(a[2])))
		}
		return ret32(h.Sys.shmget(i32(a[0]), u32(a[1]), i32(a[2])))
	}},
	{"env", "pgmem_shmat", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return ret32(h.localShmat(i32(a[0])))
		}
		return ret32(h.Sys.shmat(h, i32(a[0]), u32(a[1])))
	}},
	{"env", "pgmem_shmdt", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return 0
		}
		return ret32(h.Sys.shmdt(h, u32(a[0])))
	}},
	{"env", "pgmem_shmctl", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		var segsz, nattch, r int32
		if h.Sys == nil {
			segsz, nattch, r = h.localShmctl(i32(a[0]), i32(a[1]))
		} else {
			segsz, nattch, r = h.Sys.shmctl(i32(a[0]), i32(a[1]))
		}
		if r == 0 {
			wrI32(m, u32(a[2]), segsz)
			wrI32(m, u32(a[3]), nattch)
		}
		return ret32(r)
	}},
	{"env", "pgmem_sem", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return ret32(h.localSem(i32(a[0]), u32(a[1]), i32(a[2])))
		}
		return ret32(h.Sys.sem(h, i32(a[0]), u32(a[1]), i32(a[2])))
	}},
	{"env", "pgmem_run", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		cmd := h.str(m, u32(a[0]))
		in := h.str(m, u32(a[1]))
		out := h.str(m, u32(a[2]))
		if h.Run == nil {
			return ret32(127)
		}
		return ret32(int32(h.Run(cmd, in, out)))
	}},

	{"env", "pgmem_listen", "ii", "", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Listen != nil {
			h.Listen(h.str(m, u32(a[0])), int(i32(a[1])))
		}
		return 0
	}},
	{"env", "pgmem_poll", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fds, nfds, timeout := u32(a[0]), i32(a[1]), i32(a[2])
		if h.Sys == nil {
			// standalone: nothing can become ready; sleep for the timeout
			for i := int32(0); i < nfds; i++ {
				m.Write(fds+uint32(i)*8+6, []byte{0, 0})
			}
			h.sleep(timeout)
			return 0
		}
		return h.poll(m, fds, nfds, timeout)
	}},
	// pgmem_crc32c(crc, data, len): PostgreSQL's raw CRC-32C state update
	// (INIT/FIN are applied by the caller), computed with Go's hardware
	// accelerated hash/crc32.
	{"env", "pgmem_crc32c", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		crc := u32(a[0])
		b, ok := m.Read(u32(a[1]), u32(a[2]))
		if !ok {
			return ret32(int32(crc))
		}
		return ret32(int32(^crc32.Update(^crc, castagnoli, b)))
	}},
	// ===== cryptohash (md5(), sha256(), SCRAM HMAC ...) on Go's crypto/* =====
	{"env", "pgmem_hash_create", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		typ := i32(a[0])
		if h.hashes == nil {
			h.hashes = map[int32]hash.Hash{}
			h.hashType = map[int32]int32{}
			h.hashFree = map[int32][]hash.Hash{}
		}
		var hh hash.Hash
		if free := h.hashFree[typ]; len(free) > 0 {
			hh = free[len(free)-1]
			h.hashFree[typ] = free[:len(free)-1]
			hh.Reset()
		} else if hh = newHash(typ); hh == nil {
			return 0
		}
		h.nextHash++
		if h.nextHash <= 0 {
			h.nextHash = 1
		}
		h.hashes[h.nextHash] = hh
		h.hashType[h.nextHash] = typ
		return ret32(h.nextHash)
	}},
	{"env", "pgmem_hash_reset", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if hh := h.hashes[i32(a[0])]; hh != nil {
			hh.Reset()
			return 0
		}
		return ret32(-1)
	}},
	{"env", "pgmem_hash_update", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		hh := h.hashes[i32(a[0])]
		if hh == nil {
			return ret32(-1)
		}
		b, ok := m.Read(u32(a[1]), u32(a[2]))
		if !ok {
			return ret32(-1)
		}
		hh.Write(b)
		return 0
	}},
	{"env", "pgmem_hash_final", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		hh := h.hashes[i32(a[0])]
		if hh == nil {
			return ret32(-1)
		}
		if uint32(hh.Size()) > u32(a[2]) {
			return ret32(-1)
		}
		m.Write(u32(a[1]), hh.Sum(nil))
		return ret32(int32(hh.Size()))
	}},
	{"env", "pgmem_hash_free", "i", "", func(h *Host, m Memory, a []uint64) uint64 {
		id := i32(a[0])
		if hh := h.hashes[id]; hh != nil {
			typ := h.hashType[id]
			if len(h.hashFree[typ]) < 8 {
				h.hashFree[typ] = append(h.hashFree[typ], hh)
			}
			delete(h.hashes, id)
			delete(h.hashType, id)
		}
		return 0
	}},
	// pgmem_random_bytes(buf, n): crypto/rand straight into guest memory.
	{"env", "pgmem_random_bytes", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		buf := make([]byte, u32(a[1]))
		if _, err := rand.Read(buf); err != nil {
			return 0
		}
		if !m.Write(u32(a[0]), buf) {
			return 0
		}
		return 1
	}},

	// ===== process / runtime =====
	{"env", "__assert_fail", "iiii", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&AbortError{Msg: fmt.Sprintf("assertion failed: %s at %s:%d (%s)", h.str(m, u32(a[0])), h.str(m, u32(a[1])), i32(a[2]), h.str(m, u32(a[3])))})
	}},
	{"env", "_abort_js", "", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&AbortError{Msg: "abort()"})
	}},
	{"env", "exit", "i", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&ExitError{Code: i32(a[0])})
	}},
	{"wasi_snapshot_preview1", "proc_exit", "i", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&ExitError{Code: i32(a[0])})
	}},
	{"env", "emscripten_exit_with_live_runtime", "", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&ExitError{Live: true})
	}},
	{"env", "_emscripten_runtime_keepalive_clear", "", "", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "__call_sighandler", "ii", "", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Guest != nil {
			if err := h.Guest.CallSighandler(i32(a[0]), i32(a[1])); err != nil {
				panic(err)
			}
		}
		return 0
	}},
	{"env", "emscripten_resize_heap", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		req := u32(a[0])
		cur := m.Size()
		if req <= cur {
			return 1
		}
		// Grow geometrically like Emscripten does, capped at 4 GiB.
		want := uint64(req)
		overGrown := uint64(float64(cur) * 1.2)
		if want < overGrown {
			want = overGrown
		}
		if want > 0xFFFF0000 {
			want = 0xFFFF0000
		}
		if h.HeapMax != 0 && want > uint64(h.HeapMax) {
			want = uint64(h.HeapMax)
		}
		if want < uint64(req) {
			return 0
		}
		pages := (want + 65535) / 65536
		curPages := uint64(cur) / 65536
		if _, ok := m.Grow(uint32(pages - curPages)); !ok {
			// try exact
			exact := (uint64(req) + 65535) / 65536
			if _, ok := m.Grow(uint32(exact - curPages)); !ok {
				return 0
			}
		}
		return 1
	}},
	{"env", "emscripten_get_heap_max", "", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.HeapMax != 0 {
			return ret32(int32(h.HeapMax))
		}
		return ret32(int32(0xFFFF0000 - 0x100000000)) // 4 GiB - 64 KiB as i32
	}},

	// ===== time =====
	{"env", "emscripten_date_now", "", "d", func(h *Host, m Memory, a []uint64) uint64 {
		h.fireTimers()
		return retf64(float64(time.Now().UnixNano()) / 1e6)
	}},
	{"env", "emscripten_get_now", "", "d", func(h *Host, m Memory, a []uint64) uint64 {
		h.fireTimers()
		return retf64(float64(time.Since(h.start).Nanoseconds()) / 1e6)
	}},
	{"wasi_snapshot_preview1", "clock_time_get", "iji", "i", func(h *Host, m Memory, a []uint64) uint64 {
		h.fireTimers()
		id := i32(a[0])
		var ns int64
		switch id {
		case 0: // realtime
			ns = time.Now().UnixNano()
		default: // monotonic and friends
			ns = time.Since(h.start).Nanoseconds()
		}
		wrU64(m, u32(a[2]), uint64(ns))
		return 0
	}},
	{"env", "_tzset_js", "iiii", "", func(h *Host, m Memory, a []uint64) uint64 {
		// timezone (seconds west of UTC), daylight, std_name, dst_name. We
		// always run the guest in UTC.
		wrU32(m, u32(a[0]), 0)
		wrU32(m, u32(a[1]), 0)
		m.Write(u32(a[2]), []byte("UTC\x00"))
		m.Write(u32(a[3]), []byte("UTC\x00"))
		return 0
	}},
	{"env", "_gmtime_js", "ji", "i", func(h *Host, m Memory, a []uint64) uint64 {
		writeTm(m, u32(a[1]), time.Unix(i64(a[0]), 0).UTC())
		return 0
	}},
	{"env", "_localtime_js", "ji", "i", func(h *Host, m Memory, a []uint64) uint64 {
		writeTm(m, u32(a[1]), time.Unix(i64(a[0]), 0).UTC())
		return 0
	}},
	{"env", "_mktime_js", "i", "j", func(h *Host, m Memory, a []uint64) uint64 {
		p := u32(a[0])
		t := time.Date(int(rdI32(m, p+20))+1900, time.Month(rdI32(m, p+16)+1), int(rdI32(m, p+12)),
			int(rdI32(m, p+8)), int(rdI32(m, p+4)), int(rdI32(m, p)), 0, time.UTC)
		writeTm(m, p, t)
		return uint64(t.Unix())
	}},
	{"env", "_setitimer_js", "id", "i", func(h *Host, m Memory, a []uint64) uint64 {
		which := i32(a[0])
		ms := f64(a[1])
		if which < 0 || int(which) >= len(h.timers) {
			return errno(vfs.EINVAL)
		}
		if t := h.timerFns[which]; t != nil {
			t.Stop()
			h.timerFns[which] = nil
		}
		if ms == 0 {
			h.timers[which] = time.Time{}
		} else {
			d := time.Duration(ms * float64(time.Millisecond))
			h.timers[which] = time.Now().Add(d)
			w := uint(which)
			h.timerFns[which] = time.AfterFunc(d, func() {
				h.timerDue.Or(1 << w)
				h.Wake()
			})
		}
		return 0
	}},

	// ===== environment =====
	{"wasi_snapshot_preview1", "environ_sizes_get", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		size := 0
		for _, e := range h.Env {
			size += len(e) + 1
		}
		wrU32(m, u32(a[0]), uint32(len(h.Env)))
		wrU32(m, u32(a[1]), uint32(size))
		return 0
	}},
	{"wasi_snapshot_preview1", "environ_get", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		envp, buf := u32(a[0]), u32(a[1])
		for i, e := range h.Env {
			wrU32(m, envp+uint32(i)*4, buf)
			m.Write(buf, append([]byte(e), 0))
			buf += uint32(len(e) + 1)
		}
		return 0
	}},
	{"wasi_snapshot_preview1", "random_get", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		buf := make([]byte, u32(a[1]))
		rand.Read(buf)
		m.Write(u32(a[0]), buf)
		return 0
	}},

	// ===== WASI fd ops =====
	{"wasi_snapshot_preview1", "fd_close", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return uint64(h.FS.Close(i32(a[0])))
	}},
	{"wasi_snapshot_preview1", "fd_sync", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return uint64(h.FS.Fsync(i32(a[0])))
	}},
	{"wasi_snapshot_preview1", "fd_fdstat_get", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, buf := i32(a[0]), u32(a[1])
		kind, err := h.FS.Kind(fd)
		if err != vfs.OK {
			return uint64(err)
		}
		var ftype byte
		switch kind {
		case vfs.KindDir:
			ftype = 3
		case vfs.KindReg:
			ftype = 4
		case vfs.KindLink:
			ftype = 7
		default:
			ftype = 2 // character device
		}
		flags, _ := h.FS.Flags(fd)
		var fdflags uint16
		if flags&vfs.O_APPEND != 0 {
			fdflags |= 1
		}
		if flags&vfs.O_NONBLOCK != 0 {
			fdflags |= 4
		}
		var b [24]byte
		b[0] = ftype
		binary.LittleEndian.PutUint16(b[2:], fdflags)
		binary.LittleEndian.PutUint64(b[8:], ^uint64(0))  // rights base
		binary.LittleEndian.PutUint64(b[16:], ^uint64(0)) // rights inheriting
		m.Write(buf, b[:])
		return 0
	}},
	{"wasi_snapshot_preview1", "fd_read", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.readv(m, i32(a[0]), u32(a[1]), u32(a[2]), u32(a[3]), -1)
	}},
	{"wasi_snapshot_preview1", "fd_pread", "iiiji", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.readv(m, i32(a[0]), u32(a[1]), u32(a[2]), u32(a[4]), i64(a[3]))
	}},
	{"wasi_snapshot_preview1", "fd_write", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.writev(m, i32(a[0]), u32(a[1]), u32(a[2]), u32(a[3]), -1)
	}},
	{"wasi_snapshot_preview1", "fd_pwrite", "iiiji", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.writev(m, i32(a[0]), u32(a[1]), u32(a[2]), u32(a[4]), i64(a[3]))
	}},
	{"wasi_snapshot_preview1", "fd_seek", "ijii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		np, err := h.FS.Seek(i32(a[0]), i64(a[1]), i32(a[2]))
		if err != vfs.OK {
			return uint64(err)
		}
		wrU64(m, u32(a[3]), uint64(np))
		return 0
	}},

	// ===== Emscripten syscalls (return -errno) =====
	{"env", "__syscall_openat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		dirfd, path, flags, varargs := i32(a[0]), h.str(m, u32(a[1])), i32(a[2]), u32(a[3])
		var mode uint32
		if varargs != 0 {
			mode = rdU32(m, varargs)
		}
		fd, err := h.FS.Openat(dirfd, path, flags, mode)
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(fd)
	}},
	{"env", "__syscall_mkdirat", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Mkdirat(i32(a[0]), h.str(m, u32(a[1])), u32(a[2])))
	}},
	{"env", "__syscall_unlinkat", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Unlinkat(i32(a[0]), h.str(m, u32(a[1])), i32(a[2])))
	}},
	{"env", "__syscall_rmdir", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Unlinkat(vfs.AT_FDCWD, h.str(m, u32(a[0])), vfs.AT_REMOVEDIR))
	}},
	{"env", "__syscall_renameat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Renameat(i32(a[0]), h.str(m, u32(a[1])), i32(a[2]), h.str(m, u32(a[3]))))
	}},
	{"env", "__syscall_stat64", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		st, err := h.FS.Statat(vfs.AT_FDCWD, h.str(m, u32(a[0])), true)
		if err != vfs.OK {
			return errno(err)
		}
		writeStat(m, u32(a[1]), st)
		return 0
	}},
	{"env", "__syscall_lstat64", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		st, err := h.FS.Statat(vfs.AT_FDCWD, h.str(m, u32(a[0])), false)
		if err != vfs.OK {
			return errno(err)
		}
		writeStat(m, u32(a[1]), st)
		return 0
	}},
	{"env", "__syscall_fstat64", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		st, err := h.FS.Fstat(i32(a[0]))
		if err != vfs.OK {
			return errno(err)
		}
		writeStat(m, u32(a[1]), st)
		return 0
	}},
	{"env", "__syscall_newfstatat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		dirfd, path, buf, flags := i32(a[0]), h.str(m, u32(a[1])), u32(a[2]), i32(a[3])
		follow := flags&vfs.AT_SYMLINK_NOFOLLOW == 0
		if path == "" && flags&vfs.AT_EMPTY_PATH == 0 {
			return errno(vfs.ENOENT)
		}
		st, err := h.FS.Statat(dirfd, path, follow)
		if err != vfs.OK {
			return errno(err)
		}
		writeStat(m, buf, st)
		return 0
	}},
	{"env", "__syscall_faccessat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Accessat(i32(a[0]), h.str(m, u32(a[1]))))
	}},
	{"env", "__syscall_readlinkat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		target, err := h.FS.Readlinkat(i32(a[0]), h.str(m, u32(a[1])))
		if err != vfs.OK {
			return errno(err)
		}
		buf, size := u32(a[2]), i32(a[3])
		if size <= 0 {
			return errno(vfs.EINVAL)
		}
		if len(target) > int(size) {
			target = target[:size]
		}
		m.Write(buf, []byte(target))
		return ret32(int32(len(target)))
	}},
	{"env", "__syscall_symlinkat", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Symlinkat(h.str(m, u32(a[0])), i32(a[1]), h.str(m, u32(a[2]))))
	}},
	{"env", "__syscall_truncate64", "ij", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Truncate(h.str(m, u32(a[0])), i64(a[1])))
	}},
	{"env", "__syscall_umask", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return 0
	}},
	{"env", "__syscall_ftruncate64", "ij", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Ftruncate(i32(a[0]), i64(a[1])))
	}},
	{"env", "__syscall_chmod", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Chmodat(vfs.AT_FDCWD, h.str(m, u32(a[0])), u32(a[1]), true))
	}},
	{"env", "__syscall_fchmod", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Fchmod(i32(a[0]), u32(a[1])))
	}},
	{"env", "__syscall_fchmodat2", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Chmodat(i32(a[0]), h.str(m, u32(a[1])), u32(a[2]), i32(a[3])&vfs.AT_SYMLINK_NOFOLLOW == 0))
	}},
	{"env", "__syscall_fchown32", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "__syscall_fchownat", "iiiii", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "__syscall_utimensat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Utimensat(i32(a[0]), h.str(m, u32(a[1])), i32(a[3])&vfs.AT_SYMLINK_NOFOLLOW == 0))
	}},
	{"env", "__syscall_chdir", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Chdir(h.str(m, u32(a[0]))))
	}},
	{"env", "__syscall_getcwd", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		buf, size := u32(a[0]), u32(a[1])
		cwd := h.FS.Getcwd()
		if size == 0 && buf != 0 {
			return errno(vfs.EINVAL)
		}
		if uint32(len(cwd)+1) > size {
			return errno(vfs.ERANGE)
		}
		m.Write(buf, append([]byte(cwd), 0))
		return ret32(int32(len(cwd) + 1))
	}},
	// Emscripten 6 emits this identity query for a few libc paths. The
	// in-memory runtime has no guest identity separation, so root is correct.
	{"env", "__syscall_getegid32", "", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return 0
	}},
	{"env", "__syscall_getdents64", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, dirp, count := i32(a[0]), u32(a[1]), u32(a[2])
		ents, err := h.FS.Getdents(fd, int(count/direntSize), direntSize)
		if err != vfs.OK {
			return errno(err)
		}
		pos := uint32(0)
		for i, e := range ents {
			var b [direntSize]byte
			le := binary.LittleEndian
			le.PutUint64(b[0:], e.Ino)
			le.PutUint64(b[8:], uint64(i+1)*direntSize)
			le.PutUint16(b[16:], direntSize)
			b[18] = e.Type
			name := e.Name
			if len(name) > 255 {
				name = name[:255]
			}
			copy(b[19:], name)
			m.Write(dirp+pos, b[:])
			pos += direntSize
		}
		return ret32(int32(pos))
	}},
	{"env", "__syscall_dup", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, err := h.FS.Dup(i32(a[0]), 0)
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(fd)
	}},
	{"env", "__syscall_dup3", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, err := h.FS.Dup3(i32(a[0]), i32(a[1]))
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(fd)
	}},
	{"env", "__syscall_fcntl64", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, cmd, varargs := i32(a[0]), i32(a[1]), u32(a[2])
		switch cmd {
		case 0, 1030: // F_DUPFD, F_DUPFD_CLOEXEC
			min := rdI32(m, varargs)
			if min < 0 {
				return errno(vfs.EINVAL)
			}
			nfd, err := h.FS.Dup(fd, min)
			if err != vfs.OK {
				return errno(err)
			}
			return ret32(nfd)
		case 1, 2: // F_GETFD, F_SETFD
			if _, err := h.FS.Flags(fd); err != vfs.OK {
				return errno(err)
			}
			return 0
		case 3: // F_GETFL
			fl, err := h.FS.Flags(fd)
			if err != vfs.OK {
				return errno(err)
			}
			return ret32(fl)
		case 4: // F_SETFL
			return errno(h.FS.SetFlags(fd, rdI32(m, varargs)))
		case 5, 6, 7, 12, 13, 14: // F_GETLK/F_SETLK/F_SETLKW (both variants): pretend unlocked
			return 0
		}
		return errno(vfs.EINVAL)
	}},
	{"env", "__syscall_ioctl", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd := i32(a[0])
		if _, err := h.FS.Flags(fd); err != vfs.OK {
			return errno(err)
		}
		return errno(vfs.ENOTTY)
	}},
	{"env", "__syscall_pipe", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		p := u32(a[0])
		if p == 0 {
			return errno(vfs.EFAULT)
		}
		r, w := h.FS.Pipe()
		wrI32(m, p, r)
		wrI32(m, p+4, w)
		return 0
	}},
	{"env", "__syscall_pipe2", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		p := u32(a[0])
		if p == 0 {
			return errno(vfs.EFAULT)
		}
		r, w := h.FS.Pipe()
		wrI32(m, p, r)
		wrI32(m, p+4, w)
		return 0
	}},
	{"env", "__syscall_poll", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		// The pgmem bridge owns the actual client wire. For libc poll calls,
		// preserve the timeout behavior and report no filesystem readiness.
		h.sleep(i32(a[2]))
		return 0
	}},
	{"env", "__syscall_fdatasync", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Fsync(i32(a[0])))
	}},
	{"env", "__syscall_fadvise64", "ijji", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "__syscall_fallocate", "iijj", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, off, ln := i32(a[0]), i64(a[2]), i64(a[3])
		st, err := h.FS.Fstat(fd)
		if err != vfs.OK {
			return errno(err)
		}
		if off+ln > st.Size {
			return errno(h.FS.Ftruncate(fd, off+ln))
		}
		return 0
	}},
	{"env", "__syscall_statfs64", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		var b [64]byte
		le := binary.LittleEndian
		le.PutUint32(b[4:], 4096)   // f_bsize
		le.PutUint32(b[8:], 1<<20)  // f_blocks
		le.PutUint32(b[12:], 1<<20) // f_bfree
		le.PutUint32(b[16:], 1<<20) // f_bavail
		le.PutUint32(b[20:], 1<<20) // f_files
		le.PutUint32(b[24:], 1<<20) // f_ffree
		le.PutUint32(b[36:], 255)   // f_namelen
		le.PutUint32(b[40:], 4096)  // f_frsize
		m.Write(u32(a[2]), b[:])
		return 0
	}},
	{"env", "_mmap_js", "iiiijii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		// (len, prot, flags, fd, offset, allocated*, addr*)
		ln, fd, off := u32(a[0]), i32(a[3]), i64(a[4])
		allocatedP, addrP := u32(a[5]), u32(a[6])
		if h.Guest == nil {
			return errno(vfs.ENOSYS)
		}
		ptr, err := h.Guest.Memalign(65536, ln)
		if err != nil || ptr == 0 {
			return errno(vfs.ENOMEM)
		}
		// Copy file contents in, like MEMFS does.
		st, e := h.FS.Fstat(fd)
		if e != vfs.OK {
			return errno(e)
		}
		if off < st.Size {
			buf := make([]byte, ln)
			cur, _ := h.FS.Seek(fd, 0, vfs.SEEK_CUR)
			h.FS.Seek(fd, off, vfs.SEEK_SET)
			n, _ := h.FS.Read(fd, buf)
			h.FS.Seek(fd, cur, vfs.SEEK_SET)
			m.Write(ptr, buf[:n])
		}
		wrI32(m, allocatedP, 1)
		wrU32(m, addrP, ptr)
		return 0
	}},
	{"env", "_munmap_js", "iiiiij", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},

	// ===== sockets (never really used; the wire goes through pgmem_recv/send) =====
	// The postmaster's listening sockets (AF_UNIX or AF_INET: the
	// address is only a name, every connection comes from pgmem through
	// the cluster's accept queue) and the accepted client sockets.
	{"env", "__syscall_socket", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return errno(vfs.ENOSYS)
		}
		return ret32(h.FS.NewSocket(&Listener{}, nil))
	}},
	{"env", "__syscall_bind", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return errno(vfs.ENOSYS)
		}
		fd, addr, ln := i32(a[0]), u32(a[1]), u32(a[2])
		sock, err := h.FS.Socket(fd)
		if err != vfs.OK {
			return errno(err)
		}
		l, ok := sock.(*Listener)
		if !ok {
			return errno(vfs.EINVAL)
		}
		if ln >= 2 && rdU32(m, addr)&0xffff == 1 { // AF_UNIX: sun_path follows
			path := h.str(m, addr+2)
			if err := h.FS.BindSocket(fd, path); err != vfs.OK {
				return errno(err)
			}
			l.path = path
		}
		return 0
	}},
	{"env", "__syscall_listen", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return errno(vfs.ENOSYS)
		}
		if _, err := h.FS.Socket(i32(a[0])); err != vfs.OK {
			return errno(err)
		}
		return 0
	}},
	{"env", "__syscall_accept4", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Sys == nil {
			return errno(vfs.ENOSYS)
		}
		return h.accept(m, i32(a[0]), u32(a[1]), u32(a[2]))
	}},
	{"env", "__syscall_connect", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 { return errno(vfs.ENOSYS) }},
	{"env", "__syscall_recvfrom", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 { return errno(vfs.ENOSYS) }},
	{"env", "__syscall_sendto", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 { return errno(vfs.ENOSYS) }},
	{"env", "__syscall__newselect", "iiiii", "i", func(h *Host, m Memory, a []uint64) uint64 { return errno(vfs.ENOSYS) }},
	{"env", "getaddrinfo", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return ret32(h.getaddrinfo(m, h.str(m, u32(a[0])), u32(a[0]) != 0, h.str(m, u32(a[1])), u32(a[2]), u32(a[3])))
	}},
	{"env", "getnameinfo", "iiiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 { return ret32(-2) }},

	// ===== dynamic linking: not supported in the static build =====
	{"env", "_dlopen_js", "i", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "_dlsym_js", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "_emscripten_throw_longjmp", "", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&AbortError{Msg: "longjmp in emscripten mode is not supported by this host"})
	}},
}

// connSock returns the client connection behind a socket descriptor.
func (h *Host) connSock(fd int32) (*ConnSock, vfs.Errno) {
	sock, err := h.FS.Socket(fd)
	if err != vfs.OK {
		return nil, err
	}
	s, ok := sock.(*ConnSock)
	if !ok {
		return nil, vfs.ENOTCONN
	}
	return s, vfs.OK
}

// getaddrinfo resolves numeric hosts (and "localhost") the way
// Emscripten's JS implementation does: ai and ai_addr are separate malloc
// blocks, which is what the guest's freeaddrinfo expects.
func (h *Host) getaddrinfo(m Memory, node string, hasNode bool, service string, hint, out uint32) int32 {
	const (
		afUnspec   = 0
		afInet     = 2
		afInet6    = 10
		aiPassive  = 1
		eaiNoname  = -2
		eaiService = -8
		eaiFamily  = -6
	)
	var flags, family, socktype, proto int32
	if hint != 0 {
		flags, family, socktype, proto = rdI32(m, hint), rdI32(m, hint+4), rdI32(m, hint+8), rdI32(m, hint+12)
	}
	if socktype == 0 {
		socktype = 1 // SOCK_STREAM
	}
	if proto == 0 {
		proto = 6 // IPPROTO_TCP
	}
	var port int
	if service != "" {
		p, err := strconv.Atoi(service)
		if err != nil {
			return eaiService
		}
		port = p
	}
	var ip net.IP
	switch {
	case !hasNode || node == "":
		if family == afInet6 {
			ip = net.IPv6zero
			if flags&aiPassive == 0 {
				ip = net.IPv6loopback
			}
		} else {
			ip = net.IPv4zero
			if flags&aiPassive == 0 {
				ip = net.IPv4(127, 0, 0, 1)
			}
		}
	case node == "localhost":
		if family == afInet6 {
			ip = net.IPv6loopback
		} else {
			ip = net.IPv4(127, 0, 0, 1)
		}
	default:
		ip = net.ParseIP(node)
		if ip == nil {
			return eaiNoname
		}
	}
	var sa []byte
	if v4 := ip.To4(); v4 != nil && family != afInet6 {
		family = afInet
		sa = make([]byte, 16)
		binary.LittleEndian.PutUint16(sa[0:], afInet)
		binary.BigEndian.PutUint16(sa[2:], uint16(port))
		copy(sa[4:8], v4)
	} else if family != afInet {
		family = afInet6
		sa = make([]byte, 28)
		binary.LittleEndian.PutUint16(sa[0:], afInet6)
		binary.BigEndian.PutUint16(sa[2:], uint16(port))
		copy(sa[8:24], ip.To16())
	} else {
		return eaiFamily
	}
	if h.Guest == nil {
		return eaiNoname
	}
	saPtr, err := h.Guest.Malloc(uint32(len(sa)))
	if err != nil {
		return eaiNoname
	}
	aiPtr, err := h.Guest.Malloc(32)
	if err != nil {
		return eaiNoname
	}
	m.Write(saPtr, sa)
	var ai [32]byte
	le := binary.LittleEndian
	le.PutUint32(ai[4:], uint32(family))
	le.PutUint32(ai[8:], uint32(socktype))
	le.PutUint32(ai[12:], uint32(proto))
	le.PutUint32(ai[16:], uint32(len(sa)))
	le.PutUint32(ai[20:], saPtr)
	m.Write(aiPtr, ai[:])
	wrU32(m, out, aiPtr)
	return 0
}

func (h *Host) readv(m Memory, fd int32, iov, iovcnt, pnum uint32, offset int64) uint64 {
	total := 0
	var saved int64
	if offset >= 0 {
		saved, _ = h.FS.Seek(fd, 0, vfs.SEEK_CUR)
		if _, err := h.FS.Seek(fd, offset, vfs.SEEK_SET); err != vfs.OK {
			return uint64(err)
		}
	}
	for i := uint32(0); i < iovcnt; i++ {
		base := rdU32(m, iov+i*8)
		ln := rdU32(m, iov+i*8+4)
		// Read directly into linear memory: no intermediate buffer.
		dst, ok := m.Read(base, ln)
		if !ok {
			return uint64(vfs.EFAULT)
		}
		n, err := h.FS.Read(fd, dst)
		if err != vfs.OK {
			if total > 0 {
				break
			}
			if offset >= 0 {
				h.FS.Seek(fd, saved, vfs.SEEK_SET)
			}
			return uint64(err)
		}
		total += n
		if uint32(n) < ln {
			break
		}
	}
	if offset >= 0 {
		h.FS.Seek(fd, saved, vfs.SEEK_SET)
	}
	wrU32(m, pnum, uint32(total))
	return 0
}

func (h *Host) writev(m Memory, fd int32, iov, iovcnt, pnum uint32, offset int64) uint64 {
	total := 0
	var saved int64
	if offset >= 0 {
		saved, _ = h.FS.Seek(fd, 0, vfs.SEEK_CUR)
		if _, err := h.FS.Seek(fd, offset, vfs.SEEK_SET); err != vfs.OK {
			return uint64(err)
		}
	}
	for i := uint32(0); i < iovcnt; i++ {
		base := rdU32(m, iov+i*8)
		ln := rdU32(m, iov+i*8+4)
		b, ok := m.Read(base, ln)
		if !ok {
			return uint64(vfs.EFAULT)
		}
		n, err := h.FS.Write(fd, b)
		if err != vfs.OK {
			if total > 0 {
				break
			}
			if offset >= 0 {
				h.FS.Seek(fd, saved, vfs.SEEK_SET)
			}
			return uint64(err)
		}
		total += n
	}
	if offset >= 0 {
		h.FS.Seek(fd, saved, vfs.SEEK_SET)
	}
	wrU32(m, pnum, uint32(total))
	return 0
}

func writeTm(m Memory, p uint32, t time.Time) {
	wrI32(m, p, int32(t.Second()))
	wrI32(m, p+4, int32(t.Minute()))
	wrI32(m, p+8, int32(t.Hour()))
	wrI32(m, p+12, int32(t.Day()))
	wrI32(m, p+16, int32(t.Month())-1)
	wrI32(m, p+20, int32(t.Year())-1900)
	wrI32(m, p+24, int32(t.Weekday()))
	wrI32(m, p+28, int32(t.YearDay())-1)
	wrI32(m, p+32, 0) // isdst
	wrI32(m, p+36, 0) // gmtoff
	// tm_zone left untouched (guest keeps its own pointer)
}

// CheckTimers fires any expired itimer by calling the guest's timeout
// handler (which runs PostgreSQL's SIGALRM handler). It is called between
// guest invocations and, through fireTimers, from inside host calls the
// guest makes while running, so statement_timeout and friends can
// interrupt pg_sleep and anything that reads the clock.
func (h *Host) CheckTimers() error {
	_, err := h.checkTimers()
	return err
}

// checkTimers reports whether any handler ran.
func (h *Host) checkTimers() (bool, error) {
	if h.inTimer {
		return false, nil
	}
	fired := false
	now := time.Now()
	for which := range h.timers {
		if !h.timers[which].IsZero() && !now.Before(h.timers[which]) {
			h.timers[which] = time.Time{}
			if h.Guest != nil {
				ms := float64(time.Since(h.start).Nanoseconds()) / 1e6
				h.inTimer = true
				err := h.Guest.Timeout(int32(which), ms)
				h.inTimer = false
				fired = true
				if err != nil {
					return fired, err
				}
			}
		}
	}
	return fired, nil
}

// fireTimers is CheckTimers for use inside host calls: a failure in the
// handler (an exit unwinding through it) is propagated as a panic so the
// guest call unwinds the same way it would for any other host exit. It
// reports whether a handler ran.
func (h *Host) fireTimers() bool {
	if !h.TimerPending() {
		return false
	}
	fired, err := h.checkTimers()
	if err != nil {
		panic(err)
	}
	return fired
}

// sleep implements the poll() replacement: wait for timeoutMs (or a bounded
// slice of "forever"), waking early to run an expired timer handler.
func (h *Host) sleep(timeoutMs int32) {
	const slice = 50 * time.Millisecond
	var deadline time.Time
	if timeoutMs < 0 {
		deadline = time.Now().Add(time.Second) // bounded; the caller loops
	} else {
		deadline = time.Now().Add(time.Duration(timeoutMs) * time.Millisecond)
	}
	for {
		// A signal handler ran (statement_timeout etc.): it has set the
		// latch, so return like poll() would when the self-pipe is written.
		if h.fireTimers() {
			return
		}
		now := time.Now()
		if !now.Before(deadline) {
			return
		}
		d := deadline.Sub(now)
		if d > slice {
			d = slice
		}
		// Wake for the nearest timer if it is sooner.
		for _, t := range h.timers {
			if !t.IsZero() && t.After(now) && t.Sub(now) < d {
				d = t.Sub(now)
			}
		}
		time.Sleep(d)
	}
}

// TimerPending reports whether an itimer is armed.
func (h *Host) TimerPending() bool {
	for _, t := range h.timers {
		if !t.IsZero() {
			return true
		}
	}
	return false
}
