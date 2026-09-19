// Package aot runs the PostgreSQL module as Go code generated ahead of time
// by wasm2go (internal/aot/pgaot) instead of through wazero. It plugs the
// generated package's import interfaces into the same host table the
// wazero binding uses, so the engine does not know the difference.
//
//go:generate python3 gen.py ../../wasm/out/postgres.wasm
package aot

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/shibukawa/pgmem/internal/aot/pgaot"
	"github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"github.com/shibukawa/pgmem/internal/guest"
	"github.com/shibukawa/pgmem/internal/host"
	"github.com/shibukawa/pgmem/internal/vfs"
)

// memoryMax is the linear-memory address space reserved per instance
// (wasm32 allows up to 4 GiB; 2 GiB matches the module's declared
// maximum). Only touched pages become resident.
const memoryMax = 2 << 30

func mustLookup(module, name string) host.Fn {
	fn, ok := host.Lookup(module, name)
	if !ok {
		panic("aot: no host implementation for " + module + "." + name)
	}
	return fn
}

// Factory creates instances of the generated module.
//
// Where mmap is available, every instance starts from one shared image of
// the module's data segments (4.4 MB), mapped copy-on-write: nothing is
// copied per instance and the pages a process never writes stay shared.
// The constructors still run per instance: they initialize the C
// environment from the process's own variables, so a snapshot taken
// after them would fix one process's environment for all. Elsewhere each
// instance is built from scratch.
type Factory struct {
	once sync.Once
	img  *base.SharedImage
}

// image builds the shared snapshot on first use.
func (f *Factory) image() *base.SharedImage {
	f.once.Do(func() {
		f.img = base.NewSharedImageInPlace(memoryMax, func(mem []byte) (*base.Module, error) {
			h := host.New(vfs.New())
			inst := &instance{h: h}
			imp := &imports{h: h, mem: inst}
			inst.m = pgaot.NewWithMemory(imp, imp, mem, pgaot.InitialMemoryBytes)
			pgaot.InitData(inst.m)
			return inst.m, nil
		})
	})
	return f.img
}

// Instantiate implements guest.Factory.
func (f *Factory) Instantiate(ctx context.Context, h *host.Host) (guest.Instance, error) {
	inst := &instance{h: h}
	imp := &imports{h: h, mem: inst}
	if img := f.image(); img.Err() == nil {
		// Linear memory is a private copy-on-write view of the image, with
		// the data segments in place and zero pages above, sparse up to
		// the whole growable range; Close unmaps it.
		mem, err := img.Memory(memoryMax)
		if err != nil {
			return nil, fmt.Errorf("aot: map the data image: %w", err)
		}
		inst.release = func() { base.UnmapMemory(mem) }
		inst.m = pgaot.NewWithMemory(imp, imp, mem, img.Size())
		if h.Sys != nil {
			base.ForceContendedAtomics(inst.m)
		}
		h.Guest = inst
		if err := inst.callVoid(func() { pgaot.WasmCallCtors(inst.m) }); err != nil {
			return nil, fmt.Errorf("aot: __wasm_call_ctors: %w", err)
		}
		return inst, nil
	}
	// Linear memory lives outside the Go heap: the whole growable range is
	// mapped up front and materialized lazily, growth is a bookkeeping
	// change, and Close unmaps it immediately instead of waiting for the GC.
	mem, release, err := reserveMemory(memoryMax)
	if err != nil {
		return nil, fmt.Errorf("aot: reserve linear memory: %w", err)
	}
	if err := commitMemory(mem, 0, pgaot.InitialMemoryBytes); err != nil {
		release()
		return nil, fmt.Errorf("aot: commit linear memory: %w", err)
	}
	inst.release = release
	inst.m = pgaot.NewWithMemory(imp, imp, mem, pgaot.InitialMemoryBytes)
	pgaot.InitData(inst.m) // NewWithMemory leaves the data segments to the caller
	if h.Sys != nil {
		// A process of a cluster shares memory segments with the other
		// processes' instances: its atomics must be real ones.
		base.ForceContendedAtomics(inst.m)
	}
	h.Guest = inst
	if err := inst.callVoid(func() { pgaot.WasmCallCtors(inst.m) }); err != nil {
		return nil, fmt.Errorf("aot: __wasm_call_ctors: %w", err)
	}
	return inst, nil
}

// imports serves the generated module's env and WASI interfaces.
type imports struct {
	h   *host.Host
	mem host.Memory
}

// instance is one running copy of the generated module.
type instance struct {
	m       *base.Module
	h       *host.Host
	release func() // unmaps the linear memory
	mu      sync.Mutex
}

// ---- host.Memory over the module's linear memory ----
//
// m.Memory spans the whole reservation; the guest-visible size is
// m.MemSize (what memory.size reports and what the heap grows within).
// Host reads and writes are bounded by the reservation, not MemSize: the
// shared memory segments of a cluster are mapped above the heap (see
// host.Cluster), and PostgreSQL hands buffers in them to the host (a
// page written from shared_buffers to a file).

func (in *instance) Read(off, n uint32) ([]byte, bool) {
	end := uint64(off) + uint64(n)
	if end > uint64(len(in.m.Memory)) {
		return nil, false
	}
	return in.m.Memory[off:end:end], true
}

func (in *instance) Write(off uint32, b []byte) bool {
	end := uint64(off) + uint64(len(b))
	if end > uint64(len(in.m.Memory)) {
		return false
	}
	copy(in.m.Memory[off:], b)
	return true
}

func (in *instance) Size() uint32 { return uint32(in.m.MemSize.Load()) }

// Grow adds delta pages within the reservation.
func (in *instance) Grow(delta uint32) (uint32, bool) {
	m := in.m
	m.MemMu.Lock()
	defer m.MemMu.Unlock()
	cur := m.MemSize.Load()
	want := cur + uint64(delta)*65536
	if want > uint64(len(m.Memory)) {
		return 0, false
	}
	if err := commitMemory(m.Memory, cur, want); err != nil {
		return 0, false
	}
	m.MemSize.Store(want)
	return uint32(cur / 65536), true
}

// ---- guest.Instance ----

func (in *instance) Memory() host.Memory { return in }

// callVoid runs f with exit/abort panics converted to errors and the wasm
// stack pointer restored after an unwind.
func (in *instance) callVoid(f func()) (err error) {
	sp := pgaot.EmscriptenStackGetCurrent(in.m)
	defer func() {
		if r := recover(); r != nil {
			pgaot.EmscriptenStackRestore(in.m, sp)
			switch e := r.(type) {
			case *host.ExitError:
				err = e
			case *host.AbortError:
				err = e
			case error:
				err = fmt.Errorf("aot: trap: %w", e)
			default:
				err = fmt.Errorf("aot: trap: %v", r)
			}
		}
	}()
	f()
	return nil
}

// Call implements guest.Instance.
func (in *instance) Call(name string, args ...uint64) ([]uint64, error) {
	var res []uint64
	var ok bool
	err := in.callVoid(func() { res, ok = callExport(in.m, name, args) })
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("aot: no export %q", name)
	}
	return res, nil
}

// CallI32 implements guest.Instance.
func (in *instance) CallI32(name string, args ...uint64) (int32, error) {
	res, err := in.Call(name, args...)
	if err != nil {
		return 0, err
	}
	if len(res) == 0 {
		return 0, nil
	}
	return int32(uint32(res[0])), nil
}

// Malloc implements guest.Instance and host.Guest.
func (in *instance) Malloc(n uint32) (uint32, error) {
	var p int32
	if err := in.callVoid(func() { p = pgaot.Malloc(in.m, int32(n)) }); err != nil {
		return 0, err
	}
	if p == 0 {
		return 0, errors.New("aot: malloc failed")
	}
	return uint32(p), nil
}

// Free implements guest.Instance.
func (in *instance) Free(p uint32) { in.callVoid(func() { pgaot.Free(in.m, int32(p)) }) }

// WriteCString implements guest.Instance.
func (in *instance) WriteCString(s string) (uint32, error) {
	p, err := in.Malloc(uint32(len(s) + 1))
	if err != nil {
		return 0, err
	}
	in.Write(p, append([]byte(s), 0))
	return p, nil
}

// Close implements guest.Instance: the linear memory is unmapped at once.
func (in *instance) Close(ctx context.Context) error {
	if in.m != nil {
		in.m.Memory = nil
		in.m.M = nil
		in.m = nil
	}
	if in.release != nil {
		in.release()
		in.release = nil
	}
	return nil
}

// ---- host.Guest ----

// CallSighandler implements host.Guest.
func (in *instance) CallSighandler(fp, sig int32) error {
	return in.callVoid(func() { pgaot.PgmemCallSighandler(in.m, fp, sig) })
}

// Memalign implements host.Guest.
func (in *instance) Memalign(align, size uint32) (uint32, error) {
	var p int32
	err := in.callVoid(func() { p = pgaot.EmscriptenBuiltinMemalign(in.m, int32(align), int32(size)) })
	return uint32(p), err
}

// Timeout implements host.Guest.
func (in *instance) Timeout(which int32, now float64) error {
	return in.callVoid(func() { pgaot.EmscriptenTimeout(in.m, which, now) })
}

// Raise implements host.Guest.
func (in *instance) Raise(sig int32) error {
	return in.callVoid(func() { pgaot.PgmemRaise(in.m, sig) })
}

// MapShared implements host.Guest: the segment file replaces the private
// pages at off, so every instance that maps it sees the same bytes at
// the same address.
func (in *instance) MapShared(off uint32, f *os.File, size uint32) error {
	if uint64(off)+uint64(size) > uint64(len(in.m.Memory)) {
		return errors.New("aot: shared segment outside the linear memory")
	}
	return mapShared(in.m.Memory, off, f, size)
}

// UnmapShared implements host.Guest.
func (in *instance) UnmapShared(off, size uint32) error {
	return unmapShared(in.m.Memory, off, size)
}
