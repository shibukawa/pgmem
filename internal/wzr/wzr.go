// Package wzr binds the engine-agnostic host imports to wazero and exposes a
// small typed surface over a module instance.
package wzr

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental"

	"github.com/shibukawa/pgmem/internal/guest"
	"github.com/shibukawa/pgmem/internal/host"
)

// Runtime owns a wazero runtime and compiled modules.
type Runtime struct {
	rt wazero.Runtime
}

// compilationCache returns an on-disk cache for compiled modules so only
// the first process on a machine pays the ~2s compile. PGMEM_CACHE_DIR
// overrides the location; PGMEM_CACHE_DIR=off disables it.
func compilationCache() wazero.CompilationCache {
	dir := os.Getenv("PGMEM_CACHE_DIR")
	if dir == "off" {
		return nil
	}
	if dir == "" {
		base, err := os.UserCacheDir()
		if err != nil {
			return nil
		}
		dir = filepath.Join(base, "pgmem", "wazero")
	}
	cache, err := wazero.NewCompilationCacheWithDir(dir)
	if err != nil {
		return nil
	}
	return cache
}

type hostKey struct{}

// traceHost logs every host call to stderr (PGMEM_TRACE=1).
var traceHost = os.Getenv("PGMEM_TRACE") != ""

// hostFrom returns the Host bound to the calling instance's context.
func hostFrom(ctx context.Context) *host.Host {
	h, _ := ctx.Value(hostKey{}).(*host.Host)
	return h
}

// NewRuntime creates a wazero runtime with the features the PostgreSQL
// module needs (exception handling is used for setjmp/longjmp) and
// registers every host import once. Each instance carries its own Host in
// the context it is called with, so many instances can coexist.
func NewRuntime(ctx context.Context) (*Runtime, error) {
	cfg := wazero.NewRuntimeConfigCompiler().
		WithCoreFeatures(api.CoreFeaturesV2 | experimental.CoreFeaturesExceptionHandling).
		WithCloseOnContextDone(true)
	if cache := compilationCache(); cache != nil {
		cfg = cfg.WithCompilationCache(cache)
	}
	r := &Runtime{rt: wazero.NewRuntimeWithConfig(ctx, cfg)}
	byModule := map[string][]host.Fn{}
	for _, fn := range host.Table() {
		byModule[fn.Module] = append(byModule[fn.Module], fn)
	}
	for modName, fns := range byModule {
		b := r.rt.NewHostModuleBuilder(modName)
		for _, fn := range fns {
			fn := fn
			b.NewFunctionBuilder().WithGoModuleFunction(api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				h := hostFrom(ctx)
				if h == nil {
					panic("wzr: host function called without a bound Host")
				}
				in := make([]uint64, len(fn.Params))
				copy(in, stack)
				if traceHost {
					fmt.Fprintf(os.Stderr, "[host] %s.%s %v\n", fn.Module, fn.Name, in)
				}
				res := fn.Call(h, memAdapter{mod.Memory()}, in)
				if traceHost && len(fn.Results) > 0 {
					fmt.Fprintf(os.Stderr, "[host]   -> %s = %d\n", fn.Name, int32(uint32(res)))
				}
				if len(fn.Results) > 0 {
					stack[0] = res
				}
			}), valTypes(fn.Params), valTypes(fn.Results)).Export(fn.Name)
		}
		if _, err := b.Instantiate(ctx); err != nil {
			r.rt.Close(ctx)
			return nil, fmt.Errorf("wzr: register host module %s: %w", modName, err)
		}
	}
	return r, nil
}

// Close releases the runtime and every module compiled from it.
func (r *Runtime) Close(ctx context.Context) error { return r.rt.Close(ctx) }

// Compiled is a compiled module.
type Compiled struct {
	rt *Runtime
	cm wazero.CompiledModule
}

// Compile compiles wasm bytes and checks every function import is known
// to the host table.
func (r *Runtime) Compile(ctx context.Context, wasm []byte) (*Compiled, error) {
	cm, err := r.rt.CompileModule(ctx, wasm)
	if err != nil {
		return nil, err
	}
	for _, imp := range cm.ImportedFunctions() {
		mod, name, _ := imp.Import()
		if _, ok := host.Lookup(mod, name); !ok {
			cm.Close(ctx)
			return nil, fmt.Errorf("wzr: unsupported import %s.%s", mod, name)
		}
	}
	return &Compiled{rt: r, cm: cm}, nil
}

// Close releases the compiled module.
func (c *Compiled) Close(ctx context.Context) error { return c.cm.Close(ctx) }

// Module is an instantiated module bound to one host.
type Module struct {
	ctx  context.Context
	mod  api.Module
	host *host.Host
}

type memAdapter struct{ m api.Memory }

func (a memAdapter) Read(off, n uint32) ([]byte, bool) { return a.m.Read(off, n) }
func (a memAdapter) Write(off uint32, b []byte) bool   { return a.m.Write(off, b) }
func (a memAdapter) Size() uint32                      { return a.m.Size() }
func (a memAdapter) Grow(delta uint32) (uint32, bool)  { return a.m.Grow(delta) }

func valTypes(s string) []api.ValueType {
	out := make([]api.ValueType, 0, len(s))
	for _, c := range s {
		switch c {
		case 'i':
			out = append(out, api.ValueTypeI32)
		case 'j':
			out = append(out, api.ValueTypeI64)
		case 'f':
			out = append(out, api.ValueTypeF32)
		case 'd':
			out = append(out, api.ValueTypeF64)
		}
	}
	return out
}

// Instantiate creates a fresh instance whose imports are served by h.
// Instances are single-use: each PostgreSQL "process" gets its own.
func (c *Compiled) Instantiate(ctx context.Context, h *host.Host) (guest.Instance, error) {
	ctx = context.WithValue(ctx, hostKey{}, h)
	m := &Module{ctx: ctx, host: h}
	cfg := wazero.NewModuleConfig().WithName("").WithStartFunctions() // we call __wasm_call_ctors ourselves
	mod, err := c.rt.rt.InstantiateModule(ctx, c.cm, cfg)
	if err != nil {
		return nil, fmt.Errorf("wzr: instantiate: %w", err)
	}
	m.mod = mod
	h.Guest = m
	if fn := mod.ExportedFunction("__wasm_call_ctors"); fn != nil {
		if _, err := fn.Call(ctx); err != nil {
			mod.Close(ctx)
			return nil, fmt.Errorf("wzr: __wasm_call_ctors: %w", err)
		}
	}
	return m, nil
}

// Close tears down the instance.
func (m *Module) Close(ctx context.Context) error {
	if m.mod == nil {
		return nil
	}
	err := m.mod.Close(ctx)
	m.mod = nil
	return err
}

// Memory returns the instance's linear memory.
func (m *Module) Memory() host.Memory { return memAdapter{m.mod.Memory()} }

// Call invokes an export. A guest exit (exit(), proc_exit,
// emscripten_exit_with_live_runtime) is returned as *host.ExitError; other
// traps and aborts as ordinary errors. The wasm stack pointer is restored
// after an unwinding call so the instance stays usable.
func (m *Module) Call(name string, args ...uint64) ([]uint64, error) {
	fn := m.mod.ExportedFunction(name)
	if fn == nil {
		return nil, fmt.Errorf("wzr: no export %q", name)
	}
	sp, hasSP := m.stackSave()
	res, err := fn.Call(m.ctx, args...)
	if err != nil {
		if hasSP {
			m.stackRestore(sp)
		}
		var ee *host.ExitError
		if errors.As(err, &ee) {
			return nil, ee
		}
		var ae *host.AbortError
		if errors.As(err, &ae) {
			return nil, ae
		}
		return nil, err
	}
	return res, nil
}

func (m *Module) stackSave() (uint32, bool) {
	fn := m.mod.ExportedFunction("emscripten_stack_get_current")
	if fn == nil {
		return 0, false
	}
	res, err := fn.Call(m.ctx)
	if err != nil || len(res) == 0 {
		return 0, false
	}
	return uint32(res[0]), true
}

func (m *Module) stackRestore(sp uint32) {
	fn := m.mod.ExportedFunction("_emscripten_stack_restore")
	if fn == nil {
		return
	}
	fn.Call(m.ctx, uint64(sp))
}

// CallI32 invokes an export returning a single i32.
func (m *Module) CallI32(name string, args ...uint64) (int32, error) {
	res, err := m.Call(name, args...)
	if err != nil {
		return 0, err
	}
	if len(res) == 0 {
		return 0, nil
	}
	return int32(uint32(res[0])), nil
}

// Malloc allocates guest memory.
func (m *Module) Malloc(n uint32) (uint32, error) {
	r, err := m.CallI32("malloc", uint64(n))
	if err != nil {
		return 0, err
	}
	if r == 0 {
		return 0, errors.New("wzr: malloc failed")
	}
	return uint32(r), nil
}

// Free releases guest memory.
func (m *Module) Free(p uint32) { m.Call("free", uint64(p)) }

// WriteCString copies s (NUL terminated) into freshly allocated guest memory.
func (m *Module) WriteCString(s string) (uint32, error) {
	p, err := m.Malloc(uint32(len(s) + 1))
	if err != nil {
		return 0, err
	}
	m.mod.Memory().Write(p, append([]byte(s), 0))
	return p, nil
}

// ---- host.Guest ----

// CallSighandler implements host.Guest.
func (m *Module) CallSighandler(fp, sig int32) error {
	fn := m.mod.ExportedFunction("pgmem_call_sighandler")
	if fn == nil {
		return errors.New("wzr: pgmem_call_sighandler not exported")
	}
	_, err := fn.Call(m.ctx, uint64(uint32(fp)), uint64(uint32(sig)))
	return err
}

// Memalign implements host.Guest.
func (m *Module) Memalign(align, size uint32) (uint32, error) {
	fn := m.mod.ExportedFunction("emscripten_builtin_memalign")
	if fn == nil {
		return 0, errors.New("wzr: emscripten_builtin_memalign not exported")
	}
	res, err := fn.Call(m.ctx, uint64(align), uint64(size))
	if err != nil {
		return 0, err
	}
	return uint32(res[0]), nil
}

// Raise implements host.Guest.
func (m *Module) Raise(sig int32) error {
	fn := m.mod.ExportedFunction("pgmem_raise")
	if fn == nil {
		return errors.New("wzr: pgmem_raise not exported")
	}
	_, err := fn.Call(m.ctx, uint64(uint32(sig)))
	return err
}

// MapShared implements host.Guest; wazero instances (initdb) never share memory.
func (m *Module) MapShared(off uint32, f *os.File, size uint32) error {
	return errors.New("wzr: shared memory is not supported under wazero")
}

// UnmapShared implements host.Guest.
func (m *Module) UnmapShared(off, size uint32) error { return nil }

// Timeout implements host.Guest.
func (m *Module) Timeout(which int32, now float64) error {
	fn := m.mod.ExportedFunction("_emscripten_timeout")
	if fn == nil {
		return nil
	}
	_, err := fn.Call(m.ctx, uint64(uint32(which)), math.Float64bits(now))
	return err
}
