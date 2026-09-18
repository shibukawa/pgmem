package pgaot

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "github.com/shibukawa/pgmem/internal/aot/pgaot/p5"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	__memcap := reserveBytes
	if __memcap < 33554432 {
		__memcap = 33554432
	}
	m.Memory = make([]byte, 33554432, __memcap)
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(33554432)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 2147483648
	m.T0 = make([]any, 7740)
	m.G0 = int32(13128304)
	m.G1 = int32(0)
	m.G2 = int32(0)
	InitElemSeg_0_0(m)
	InitElemSeg_0_1(m)
	InitElemSeg_0_2(m)
	InitElemSeg_0_3(m)
	InitElemSeg_0_4(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_3_2(m)
	InitElemSeg_3_3(m)
	InitElemSeg_3_4(m)
	InitElemSeg_3_5(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_4_4(m)
	InitElemSeg_4_5(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	InitElemSeg_5_2(m)
	InitElemSeg_5_3(m)
	InitElemSeg_5_4(m)
	m.DataEnd = 4453631
	initData_0(m)
	return m
}
// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, env, 41943040)
}
// New constructs a *Module using DefaultWASI() for the
// wasi_snapshot_preview1 import. Use NewWithWASI to plug in a
// custom implementation (sandboxed FS, captured stdout, ...).
func New(env base.EnvImports) *base.Module {
	return NewWithWASI(base.DefaultWASI(), env)
}

const InitialMemoryBytes = 33554432

func NewWithMemory(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, memory []byte, memSize uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 4294836224 {
		panic("wasm2go: memory size exceeds the implementation limit (4294836224 bytes)")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 7740)
	m.G0 = int32(13128304)
	m.G1 = int32(0)
	m.G2 = int32(0)
	InitElemSeg_0_0(m)
	InitElemSeg_0_1(m)
	InitElemSeg_0_2(m)
	InitElemSeg_0_3(m)
	InitElemSeg_0_4(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_3_2(m)
	InitElemSeg_3_3(m)
	InitElemSeg_3_4(m)
	InitElemSeg_3_5(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_4_4(m)
	InitElemSeg_4_5(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	InitElemSeg_5_2(m)
	InitElemSeg_5_3(m)
	InitElemSeg_5_4(m)
	m.DataEnd = 4453631
	return m
}
func NewFromSnapshot(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, memory []byte, memSize uint64, globals []uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 4294836224 {
		panic("wasm2go: memory size exceeds the implementation limit (4294836224 bytes)")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 7740)
	m.G0 = int32(13128304)
	m.G1 = int32(0)
	m.G2 = int32(0)
	InitElemSeg_0_0(m)
	InitElemSeg_0_1(m)
	InitElemSeg_0_2(m)
	InitElemSeg_0_3(m)
	InitElemSeg_0_4(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_3_2(m)
	InitElemSeg_3_3(m)
	InitElemSeg_3_4(m)
	InitElemSeg_3_5(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_4_4(m)
	InitElemSeg_4_5(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	InitElemSeg_5_2(m)
	InitElemSeg_5_3(m)
	InitElemSeg_5_4(m)
	m.DataEnd = 4453631
	base.RestoreGlobals(m, globals)
	return m
}
func initData_0(m *base.Module) {
	copy(m.Memory[4096:], wasm2goData_data_bin[0:1643504])
	copy(m.Memory[1648832:], wasm2goData_data_bin[1643504:1643513])
	copy(m.Memory[1650880:], wasm2goData_data_bin[1643513:1656441])
	copy(m.Memory[1675268:], wasm2goData_data_bin[1656441:1749650])
	copy(m.Memory[1769504:], wasm2goData_data_bin[1749650:1861306])
	copy(m.Memory[1892408:], wasm2goData_data_bin[1861306:4125264])
	copy(m.Memory[4163856:], wasm2goData_data_bin[4125264:4415039])
}
func InitData(m *base.Module) {
	initData_0(m)
}
func WasmCallCtors(m *base.Module) {
	F___wasm_call_ctors(m)
}
func Fflush(m *base.Module, l0 int32) int32 {
	return F_fflush(m, l0)
}
func PqBufferRemainingData(m *base.Module) int32 {
	return F_pq_buffer_remaining_data(m)
}
func Malloc(m *base.Module, l0 int32) int32 {
	return F_emscripten_builtin_malloc(m, l0)
}
func Free(m *base.Module, l0 int32) {
	F_emscripten_builtin_free(m, l0)
}
func PgmemPoll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_pgmem_poll(m, l0, l1, l2)
}
func ProcessStartupPacket(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_ProcessStartupPacket(m, l0, l1, l2)
}
func PglStartPGlite(m *base.Module) {
	F_pgl_startPGlite(m)
}
func PglPqFlush(m *base.Module) {
	F_pgl_pq_flush(m)
}
func PglGetMyProcPort(m *base.Module) int32 {
	return F_pgl_getMyProcPort(m)
}
func PglSendConnData(m *base.Module) {
	F_pgl_sendConnData(m)
}
func PostgresMainLongJmp(m *base.Module) {
	F_PostgresMainLongJmp(m)
}
func PostgresMainLoopOnce(m *base.Module) {
	F_PostgresMainLoopOnce(m)
}
func PostgresSendReadyForQueryIfNecessary(m *base.Module) {
	F_PostgresSendReadyForQueryIfNecessary(m)
}
func PgmemResetSession(m *base.Module, l0 int32) {
	F_pgmem_reset_session(m, l0)
}
func PglSetPGliteExitStatus(m *base.Module, l0 int32) int32 {
	return F_pgl_setPGliteExitStatus(m, l0)
}
func PglSetPGliteActive(m *base.Module, l0 int32) int32 {
	return F_pgl_setPGliteActive(m, l0)
}
func PglRunAtexitFuncs(m *base.Module) {
	F_pgl_run_atexit_funcs(m)
}
func PglFreopen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_pgl_freopen(m, l0, l1, l2)
}
func PgmemInit(m *base.Module) {
	F_pgmem_init(m)
}
func PgmemMain(m *base.Module, l0 int32, l1 int32) int32 {
	return F_pgmem_main(m, l0, l1)
}
func PgmemCallSighandler(m *base.Module, l0 int32, l1 int32) {
	F_pgmem_call_sighandler(m, l0, l1)
}
func EmscriptenBuiltinMemalign(m *base.Module, l0 int32, l1 int32) int32 {
	return F_emscripten_builtin_memalign(m, l0, l1)
}
func PgmemModuleName(m *base.Module, l0 int32) int32 {
	return F_pgmem_module_name(m, l0)
}
func EmscriptenTimeout(m *base.Module, l0 int32, l1 float64) {
	F__emscripten_timeout(m, l0, l1)
}
func EmscriptenStackRestore(m *base.Module, l0 int32) {
	F__emscripten_stack_restore(m, l0)
}
func EmscriptenStackGetCurrent(m *base.Module) int32 {
	return F_emscripten_stack_get_current(m)
}
func Memory(m *base.Module) []byte {
	return m.Memory
}
//go:embed data.bin
var wasm2goData_data_bin []byte
