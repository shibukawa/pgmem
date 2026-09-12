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
	m.T0 = make([]any, 7579)
	m.G0 = int32(13098112)
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
	m.DataEnd = 4426555
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
	m.T0 = make([]any, 7579)
	m.G0 = int32(13098112)
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
	m.DataEnd = 4426555
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
	m.T0 = make([]any, 7579)
	m.G0 = int32(13098112)
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
	m.DataEnd = 4426555
	base.RestoreGlobals(m, globals)
	return m
}
func initData_0(m *base.Module) {
	copy(m.Memory[4096:], wasm2goData_data_bin[0:1620352])
	copy(m.Memory[1625680:], wasm2goData_data_bin[1620352:1620361])
	copy(m.Memory[1627728:], wasm2goData_data_bin[1620361:1633289])
	copy(m.Memory[1650692:], wasm2goData_data_bin[1633289:1726530])
	copy(m.Memory[1744960:], wasm2goData_data_bin[1726530:1838186])
	copy(m.Memory[1867832:], wasm2goData_data_bin[1838186:4100480])
	copy(m.Memory[4137616:], wasm2goData_data_bin[4100480:4389419])
}
func InitData(m *base.Module) {
	initData_0(m)
}
func WasmCallCtors(m *base.Module) {
	F___wasm_call_ctors(m)
}
func ErrnoLocation(m *base.Module) int32 {
	return F___errno_location(m)
}
func IsTransactionBlock(m *base.Module) int32 {
	return F_IsTransactionBlock(m)
}
func Fflush(m *base.Module, l0 int32) int32 {
	return F_fflush(m, l0)
}
func PglSystem(m *base.Module, l0 int32) int32 {
	return F_pgl_system(m, l0)
}
func PglGeteuid(m *base.Module) int32 {
	return F_pgl_geteuid(m)
}
func PglExit(m *base.Module, l0 int32) {
	F_pgl_exit(m, l0)
}
func PglConnect(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_comparetup_index_hash_tiebreak(m, l0, l1, l2)
}
func PglSend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return F_pgl_send(m, l0, l1, l2, l3)
}
func PglRecv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return F_pgl_recv(m, l0, l1, l2, l3)
}
func PglGetsockname(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_comparetup_index_hash_tiebreak(m, l0, l1, l2)
}
func PglSetsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	return F_pgl_setsockopt(m, l0, l1, l2, l3, l4)
}
func PglFcntl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_comparetup_index_hash_tiebreak(m, l0, l1, l2)
}
func PqBufferRemainingData(m *base.Module) int32 {
	return F_pq_buffer_remaining_data(m)
}
func PglGetsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	return F_pgl_setsockopt(m, l0, l1, l2, l3, l4)
}
func MainArgcArgv(m *base.Module, l0 int32, l1 int32) int32 {
	return F_main(m, l0, l1)
}
func PglGetuid(m *base.Module) int32 {
	return F_pgl_geteuid(m)
}
func PglShmdt(m *base.Module, l0 int32) int32 {
	return F_pgl_shmdt(m, l0)
}
func PglShmctl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_pgl_shmctl(m, l0, l1, l2)
}
func PglShmat(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_pgl_shmat(m, l0, l1, l2)
}
func PglShmget(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_pgl_shmget(m, l0, l1, l2)
}
func PglMunmap(m *base.Module, l0 int32, l1 int32) int32 {
	return F_gist_bbox_zorder_abbrev_abort(m, l0, l1)
}
func Malloc(m *base.Module, l0 int32) int32 {
	return F_emscripten_builtin_malloc(m, l0)
}
func Free(m *base.Module, l0 int32) {
	F_emscripten_builtin_free(m, l0)
}
func PglPopen(m *base.Module, l0 int32, l1 int32) int32 {
	return F_pgl_popen(m, l0, l1)
}
func PglPclose(m *base.Module, l0 int32) int32 {
	return F_pgl_pclose(m, l0)
}
func PglAtexit(m *base.Module, l0 int32) int32 {
	return F_pgl_atexit(m, l0)
}
func PgmemPoll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_pgmem_poll(m, l0, l1, l2)
}
func ProcessStartupPacket(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_ProcessStartupPacket(m, l0, l1, l2)
}
func Htons(m *base.Module, l0 int32) int32 {
	return F_htons(m, l0)
}
func Htonl(m *base.Module, l0 int32) int32 {
	return F_htonl(m, l0)
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
func PglSetPGliteExitStatus(m *base.Module, l0 int32) int32 {
	return F_pgl_setPGliteExitStatus(m, l0)
}
func PglLongjmp(m *base.Module, l0 int32, l1 int32) {
	F_pgl_longjmp(m, l0, l1)
}
func PglGetpwuid(m *base.Module, l0 int32) int32 {
	return F_pgl_getpwuid(m, l0)
}
func Strerror(m *base.Module, l0 int32) int32 {
	return F_strerror(m, l0)
}
func ClearSetitimer(m *base.Module) {
	F_clear_setitimer(m)
}
func PglSetPGliteActive(m *base.Module, l0 int32) int32 {
	return F_pgl_setPGliteActive(m, l0)
}
func PglGetPGliteExitStatus(m *base.Module) int32 {
	return F_pgl_getPGliteExitStatus(m)
}
func PglSiglongjmp(m *base.Module, l0 int32, l1 int32) {
	F_pgl_longjmp(m, l0, l1)
}
func PglSetSystemFn(m *base.Module, l0 int32) {
	F_pgl_set_system_fn(m, l0)
}
func PglSetPopenFn(m *base.Module, l0 int32) {
	F_pgl_set_popen_fn(m, l0)
}
func PglSetPcloseFn(m *base.Module, l0 int32) {
	F_pgl_set_pclose_fn(m, l0)
}
func PglRunAtexitFuncs(m *base.Module) {
	F_pgl_run_atexit_funcs(m)
}
func PglFreopen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_pgl_freopen(m, l0, l1, l2)
}
func PglSetRwCbs(m *base.Module, l0 int32, l1 int32) {
	F_pgl_set_rw_cbs(m, l0, l1)
}
func PglPoll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_pgl_poll(m, l0, l1, l2)
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
func PgmemModuleName(m *base.Module, l0 int32) int32 {
	return F_pgmem_module_name(m, l0)
}
func PgMagicFuncPlpgsql(m *base.Module) int32 {
	return F_Pg_magic_func_plpgsql(m)
}
func PGInitPlpgsql(m *base.Module) {
	F__PG_init_plpgsql(m)
}
func PgFinfoPlpgsqlCallHandler(m *base.Module) int32 {
	return F_pg_finfo_plpgsql_call_handler(m)
}
func PgFinfoPlpgsqlInlineHandler(m *base.Module) int32 {
	return F_pg_finfo_plpgsql_inline_handler(m)
}
func PgFinfoPlpgsqlValidator(m *base.Module) int32 {
	return F_pg_finfo_plpgsql_validator(m)
}
func PlpgsqlBuildDatatype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return F_plpgsql_build_datatype(m, l0, l1, l2, l3)
}
func PlpgsqlCallHandler(m *base.Module, l0 int32) int32 {
	return F_plpgsql_call_handler(m, l0)
}
func PlpgsqlCompile(m *base.Module, l0 int32, l1 int32) int32 {
	return F_plpgsql_compile(m, l0, l1)
}
func PlpgsqlExecGetDatumType(m *base.Module, l0 int32, l1 int32) int32 {
	return F_plpgsql_exec_get_datum_type(m, l0, l1)
}
func PlpgsqlInlineHandler(m *base.Module, l0 int32) int32 {
	return F_plpgsql_inline_handler(m, l0)
}
func PlpgsqlNsLookup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return F_plpgsql_ns_lookup(m, l0, l1, l2, l3, l4, l5)
}
func PlpgsqlParserSetup(m *base.Module, l0 int32, l1 int32) {
	F_plpgsql_parser_setup(m, l0, l1)
}
func PlpgsqlRecognizeErrCondition(m *base.Module, l0 int32, l1 int32) int32 {
	return F_plpgsql_recognize_err_condition(m, l0, l1)
}
func PlpgsqlStmtTypename(m *base.Module, l0 int32) int32 {
	return F_plpgsql_stmt_typename(m, l0)
}
func PlpgsqlValidator(m *base.Module, l0 int32) int32 {
	return F_plpgsql_validator(m, l0)
}
func PgMagicFuncDictSnowball(m *base.Module) int32 {
	return F_Pg_magic_func_dict_snowball(m)
}
func DsnowballInit(m *base.Module, l0 int32) int32 {
	return F_dsnowball_init(m, l0)
}
func DsnowballLexize(m *base.Module, l0 int32) int32 {
	return F_dsnowball_lexize(m, l0)
}
func PgFinfoDsnowballInit(m *base.Module) int32 {
	return F_pg_finfo_dsnowball_init(m)
}
func PgFinfoDsnowballLexize(m *base.Module) int32 {
	return F_pg_finfo_dsnowball_lexize(m)
}
func PgMagicFuncCyrillicAndMic(m *base.Module) int32 {
	return F_Pg_magic_func_cyrillic_and_mic(m)
}
func IsoToKoi8r(m *base.Module, l0 int32) int32 {
	return F_iso_to_koi8r(m, l0)
}
func IsoToMic(m *base.Module, l0 int32) int32 {
	return F_iso_to_mic(m, l0)
}
func IsoToWin1251(m *base.Module, l0 int32) int32 {
	return F_iso_to_win1251(m, l0)
}
func IsoToWin866(m *base.Module, l0 int32) int32 {
	return F_iso_to_win866(m, l0)
}
func Koi8rToIso(m *base.Module, l0 int32) int32 {
	return F_koi8r_to_iso(m, l0)
}
func Koi8rToMic(m *base.Module, l0 int32) int32 {
	return F_koi8r_to_mic(m, l0)
}
func Koi8rToWin1251(m *base.Module, l0 int32) int32 {
	return F_koi8r_to_win1251(m, l0)
}
func Koi8rToWin866(m *base.Module, l0 int32) int32 {
	return F_koi8r_to_win866(m, l0)
}
func MicToIso(m *base.Module, l0 int32) int32 {
	return F_mic_to_iso(m, l0)
}
func MicToKoi8r(m *base.Module, l0 int32) int32 {
	return F_mic_to_koi8r(m, l0)
}
func MicToWin1251(m *base.Module, l0 int32) int32 {
	return F_mic_to_win1251(m, l0)
}
func MicToWin866(m *base.Module, l0 int32) int32 {
	return F_mic_to_win866(m, l0)
}
func PgFinfoIsoToKoi8r(m *base.Module) int32 {
	return F_pg_finfo_iso_to_koi8r(m)
}
func PgFinfoIsoToMic(m *base.Module) int32 {
	return F_pg_finfo_iso_to_mic(m)
}
func PgFinfoIsoToWin1251(m *base.Module) int32 {
	return F_pg_finfo_iso_to_win1251(m)
}
func PgFinfoIsoToWin866(m *base.Module) int32 {
	return F_pg_finfo_iso_to_win866(m)
}
func PgFinfoKoi8rToIso(m *base.Module) int32 {
	return F_pg_finfo_koi8r_to_iso(m)
}
func PgFinfoKoi8rToMic(m *base.Module) int32 {
	return F_pg_finfo_koi8r_to_mic(m)
}
func PgFinfoKoi8rToWin1251(m *base.Module) int32 {
	return F_pg_finfo_koi8r_to_win1251(m)
}
func PgFinfoKoi8rToWin866(m *base.Module) int32 {
	return F_pg_finfo_koi8r_to_win866(m)
}
func PgFinfoMicToIso(m *base.Module) int32 {
	return F_pg_finfo_mic_to_iso(m)
}
func PgFinfoMicToKoi8r(m *base.Module) int32 {
	return F_pg_finfo_mic_to_koi8r(m)
}
func PgFinfoMicToWin1251(m *base.Module) int32 {
	return F_pg_finfo_mic_to_win1251(m)
}
func PgFinfoMicToWin866(m *base.Module) int32 {
	return F_pg_finfo_mic_to_win866(m)
}
func PgFinfoWin1251ToIso(m *base.Module) int32 {
	return F_pg_finfo_win1251_to_iso(m)
}
func PgFinfoWin1251ToKoi8r(m *base.Module) int32 {
	return F_pg_finfo_win1251_to_koi8r(m)
}
func PgFinfoWin1251ToMic(m *base.Module) int32 {
	return F_pg_finfo_win1251_to_mic(m)
}
func PgFinfoWin1251ToWin866(m *base.Module) int32 {
	return F_pg_finfo_win1251_to_win866(m)
}
func PgFinfoWin866ToIso(m *base.Module) int32 {
	return F_pg_finfo_win866_to_iso(m)
}
func PgFinfoWin866ToKoi8r(m *base.Module) int32 {
	return F_pg_finfo_win866_to_koi8r(m)
}
func PgFinfoWin866ToMic(m *base.Module) int32 {
	return F_pg_finfo_win866_to_mic(m)
}
func PgFinfoWin866ToWin1251(m *base.Module) int32 {
	return F_pg_finfo_win866_to_win1251(m)
}
func Win1251ToIso(m *base.Module, l0 int32) int32 {
	return F_win1251_to_iso(m, l0)
}
func Win1251ToKoi8r(m *base.Module, l0 int32) int32 {
	return F_win1251_to_koi8r(m, l0)
}
func Win1251ToMic(m *base.Module, l0 int32) int32 {
	return F_win1251_to_mic(m, l0)
}
func Win1251ToWin866(m *base.Module, l0 int32) int32 {
	return F_win1251_to_win866(m, l0)
}
func Win866ToIso(m *base.Module, l0 int32) int32 {
	return F_win866_to_iso(m, l0)
}
func Win866ToKoi8r(m *base.Module, l0 int32) int32 {
	return F_win866_to_koi8r(m, l0)
}
func Win866ToMic(m *base.Module, l0 int32) int32 {
	return F_win866_to_mic(m, l0)
}
func Win866ToWin1251(m *base.Module, l0 int32) int32 {
	return F_win866_to_win1251(m, l0)
}
func PgMagicFuncEucCnAndMic(m *base.Module) int32 {
	return F_Pg_magic_func_euc_cn_and_mic(m)
}
func EucCnToMic(m *base.Module, l0 int32) int32 {
	return F_euc_cn_to_mic(m, l0)
}
func MicToEucCn(m *base.Module, l0 int32) int32 {
	return F_mic_to_euc_cn(m, l0)
}
func PgFinfoEucCnToMic(m *base.Module) int32 {
	return F_pg_finfo_euc_cn_to_mic(m)
}
func PgFinfoMicToEucCn(m *base.Module) int32 {
	return F_pg_finfo_mic_to_euc_cn(m)
}
func PgMagicFuncEucJpAndSjis(m *base.Module) int32 {
	return F_Pg_magic_func_euc_jp_and_sjis(m)
}
func EucJpToMic(m *base.Module, l0 int32) int32 {
	return F_euc_jp_to_mic(m, l0)
}
func EucJpToSjis(m *base.Module, l0 int32) int32 {
	return F_euc_jp_to_sjis(m, l0)
}
func MicToEucJp(m *base.Module, l0 int32) int32 {
	return F_mic_to_euc_jp(m, l0)
}
func MicToSjis(m *base.Module, l0 int32) int32 {
	return F_mic_to_sjis(m, l0)
}
func PgFinfoEucJpToMic(m *base.Module) int32 {
	return F_pg_finfo_euc_jp_to_mic(m)
}
func PgFinfoEucJpToSjis(m *base.Module) int32 {
	return F_pg_finfo_euc_jp_to_sjis(m)
}
func PgFinfoMicToEucJp(m *base.Module) int32 {
	return F_pg_finfo_mic_to_euc_jp(m)
}
func PgFinfoMicToSjis(m *base.Module) int32 {
	return F_pg_finfo_mic_to_sjis(m)
}
func PgFinfoSjisToEucJp(m *base.Module) int32 {
	return F_pg_finfo_sjis_to_euc_jp(m)
}
func PgFinfoSjisToMic(m *base.Module) int32 {
	return F_pg_finfo_sjis_to_mic(m)
}
func SjisToEucJp(m *base.Module, l0 int32) int32 {
	return F_sjis_to_euc_jp(m, l0)
}
func SjisToMic(m *base.Module, l0 int32) int32 {
	return F_sjis_to_mic(m, l0)
}
func PgMagicFuncEucKrAndMic(m *base.Module) int32 {
	return F_Pg_magic_func_euc_kr_and_mic(m)
}
func EucKrToMic(m *base.Module, l0 int32) int32 {
	return F_euc_kr_to_mic(m, l0)
}
func MicToEucKr(m *base.Module, l0 int32) int32 {
	return F_mic_to_euc_kr(m, l0)
}
func PgFinfoEucKrToMic(m *base.Module) int32 {
	return F_pg_finfo_euc_kr_to_mic(m)
}
func PgFinfoMicToEucKr(m *base.Module) int32 {
	return F_pg_finfo_mic_to_euc_kr(m)
}
func PgMagicFuncEucTwAndBig5(m *base.Module) int32 {
	return F_Pg_magic_func_euc_tw_and_big5(m)
}
func Big5ToEucTw(m *base.Module, l0 int32) int32 {
	return F_big5_to_euc_tw(m, l0)
}
func Big5ToMic(m *base.Module, l0 int32) int32 {
	return F_big5_to_mic(m, l0)
}
func EucTwToBig5(m *base.Module, l0 int32) int32 {
	return F_euc_tw_to_big5(m, l0)
}
func EucTwToMic(m *base.Module, l0 int32) int32 {
	return F_euc_tw_to_mic(m, l0)
}
func MicToBig5(m *base.Module, l0 int32) int32 {
	return F_mic_to_big5(m, l0)
}
func MicToEucTw(m *base.Module, l0 int32) int32 {
	return F_mic_to_euc_tw(m, l0)
}
func PgFinfoBig5ToEucTw(m *base.Module) int32 {
	return F_pg_finfo_big5_to_euc_tw(m)
}
func PgFinfoBig5ToMic(m *base.Module) int32 {
	return F_pg_finfo_big5_to_mic(m)
}
func PgFinfoEucTwToBig5(m *base.Module) int32 {
	return F_pg_finfo_euc_tw_to_big5(m)
}
func PgFinfoEucTwToMic(m *base.Module) int32 {
	return F_pg_finfo_euc_tw_to_mic(m)
}
func PgFinfoMicToBig5(m *base.Module) int32 {
	return F_pg_finfo_mic_to_big5(m)
}
func PgFinfoMicToEucTw(m *base.Module) int32 {
	return F_pg_finfo_mic_to_euc_tw(m)
}
func PgMagicFuncEuc2004Sjis2004(m *base.Module) int32 {
	return F_Pg_magic_func_euc2004_sjis2004(m)
}
func EucJis_2004_toShiftJis_2004(m *base.Module, l0 int32) int32 {
	return F_euc_jis_2004_to_shift_jis_2004(m, l0)
}
func PgFinfoEucJis_2004_toShiftJis_2004(m *base.Module) int32 {
	return F_pg_finfo_euc_jis_2004_to_shift_jis_2004(m)
}
func PgFinfoShiftJis_2004_toEucJis_2004(m *base.Module) int32 {
	return F_pg_finfo_shift_jis_2004_to_euc_jis_2004(m)
}
func ShiftJis_2004_toEucJis_2004(m *base.Module, l0 int32) int32 {
	return F_shift_jis_2004_to_euc_jis_2004(m, l0)
}
func PgMagicFuncLatinAndMic(m *base.Module) int32 {
	return F_Pg_magic_func_latin_and_mic(m)
}
func Latin1ToMic(m *base.Module, l0 int32) int32 {
	return F_latin1_to_mic(m, l0)
}
func Latin3ToMic(m *base.Module, l0 int32) int32 {
	return F_latin3_to_mic(m, l0)
}
func Latin4ToMic(m *base.Module, l0 int32) int32 {
	return F_latin4_to_mic(m, l0)
}
func MicToLatin1(m *base.Module, l0 int32) int32 {
	return F_mic_to_latin1(m, l0)
}
func MicToLatin3(m *base.Module, l0 int32) int32 {
	return F_mic_to_latin3(m, l0)
}
func MicToLatin4(m *base.Module, l0 int32) int32 {
	return F_mic_to_latin4(m, l0)
}
func PgFinfoLatin1ToMic(m *base.Module) int32 {
	return F_pg_finfo_latin1_to_mic(m)
}
func PgFinfoLatin3ToMic(m *base.Module) int32 {
	return F_pg_finfo_latin3_to_mic(m)
}
func PgFinfoLatin4ToMic(m *base.Module) int32 {
	return F_pg_finfo_latin4_to_mic(m)
}
func PgFinfoMicToLatin1(m *base.Module) int32 {
	return F_pg_finfo_mic_to_latin1(m)
}
func PgFinfoMicToLatin3(m *base.Module) int32 {
	return F_pg_finfo_mic_to_latin3(m)
}
func PgFinfoMicToLatin4(m *base.Module) int32 {
	return F_pg_finfo_mic_to_latin4(m)
}
func PgMagicFuncLatin2AndWin1250(m *base.Module) int32 {
	return F_Pg_magic_func_latin2_and_win1250(m)
}
func Latin2ToMic(m *base.Module, l0 int32) int32 {
	return F_latin2_to_mic(m, l0)
}
func Latin2ToWin1250(m *base.Module, l0 int32) int32 {
	return F_latin2_to_win1250(m, l0)
}
func MicToLatin2(m *base.Module, l0 int32) int32 {
	return F_mic_to_latin2(m, l0)
}
func MicToWin1250(m *base.Module, l0 int32) int32 {
	return F_mic_to_win1250(m, l0)
}
func PgFinfoLatin2ToMic(m *base.Module) int32 {
	return F_pg_finfo_latin2_to_mic(m)
}
func PgFinfoLatin2ToWin1250(m *base.Module) int32 {
	return F_pg_finfo_latin2_to_win1250(m)
}
func PgFinfoMicToLatin2(m *base.Module) int32 {
	return F_pg_finfo_mic_to_latin2(m)
}
func PgFinfoMicToWin1250(m *base.Module) int32 {
	return F_pg_finfo_mic_to_win1250(m)
}
func PgFinfoWin1250ToLatin2(m *base.Module) int32 {
	return F_pg_finfo_win1250_to_latin2(m)
}
func PgFinfoWin1250ToMic(m *base.Module) int32 {
	return F_pg_finfo_win1250_to_mic(m)
}
func Win1250ToLatin2(m *base.Module, l0 int32) int32 {
	return F_win1250_to_latin2(m, l0)
}
func Win1250ToMic(m *base.Module, l0 int32) int32 {
	return F_win1250_to_mic(m, l0)
}
func PgMagicFuncUtf8AndBig5(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_big5(m)
}
func Big5ToUtf8(m *base.Module, l0 int32) int32 {
	return F_big5_to_utf8(m, l0)
}
func PgFinfoBig5ToUtf8(m *base.Module) int32 {
	return F_pg_finfo_big5_to_utf8(m)
}
func PgFinfoUtf8ToBig5(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_big5(m)
}
func Utf8ToBig5(m *base.Module, l0 int32) int32 {
	return F_utf8_to_big5(m, l0)
}
func PgMagicFuncUtf8AndCyrillic(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_cyrillic(m)
}
func Koi8rToUtf8(m *base.Module, l0 int32) int32 {
	return F_koi8r_to_utf8(m, l0)
}
func Koi8uToUtf8(m *base.Module, l0 int32) int32 {
	return F_koi8u_to_utf8(m, l0)
}
func PgFinfoKoi8rToUtf8(m *base.Module) int32 {
	return F_pg_finfo_koi8r_to_utf8(m)
}
func PgFinfoKoi8uToUtf8(m *base.Module) int32 {
	return F_pg_finfo_koi8u_to_utf8(m)
}
func PgFinfoUtf8ToKoi8r(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_koi8r(m)
}
func PgFinfoUtf8ToKoi8u(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_koi8u(m)
}
func Utf8ToKoi8r(m *base.Module, l0 int32) int32 {
	return F_utf8_to_koi8r(m, l0)
}
func Utf8ToKoi8u(m *base.Module, l0 int32) int32 {
	return F_utf8_to_koi8u(m, l0)
}
func PgMagicFuncUtf8AndEucCn(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_euc_cn(m)
}
func EucCnToUtf8(m *base.Module, l0 int32) int32 {
	return F_euc_cn_to_utf8(m, l0)
}
func PgFinfoEucCnToUtf8(m *base.Module) int32 {
	return F_pg_finfo_euc_cn_to_utf8(m)
}
func PgFinfoUtf8ToEucCn(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_euc_cn(m)
}
func Utf8ToEucCn(m *base.Module, l0 int32) int32 {
	return F_utf8_to_euc_cn(m, l0)
}
func PgMagicFuncUtf8AndEucJp(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_euc_jp(m)
}
func EucJpToUtf8(m *base.Module, l0 int32) int32 {
	return F_euc_jp_to_utf8(m, l0)
}
func PgFinfoEucJpToUtf8(m *base.Module) int32 {
	return F_pg_finfo_euc_jp_to_utf8(m)
}
func PgFinfoUtf8ToEucJp(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_euc_jp(m)
}
func Utf8ToEucJp(m *base.Module, l0 int32) int32 {
	return F_utf8_to_euc_jp(m, l0)
}
func PgMagicFuncUtf8AndEucKr(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_euc_kr(m)
}
func EucKrToUtf8(m *base.Module, l0 int32) int32 {
	return F_euc_kr_to_utf8(m, l0)
}
func PgFinfoEucKrToUtf8(m *base.Module) int32 {
	return F_pg_finfo_euc_kr_to_utf8(m)
}
func PgFinfoUtf8ToEucKr(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_euc_kr(m)
}
func Utf8ToEucKr(m *base.Module, l0 int32) int32 {
	return F_utf8_to_euc_kr(m, l0)
}
func PgMagicFuncUtf8AndEucTw(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_euc_tw(m)
}
func EucTwToUtf8(m *base.Module, l0 int32) int32 {
	return F_euc_tw_to_utf8(m, l0)
}
func PgFinfoEucTwToUtf8(m *base.Module) int32 {
	return F_pg_finfo_euc_tw_to_utf8(m)
}
func PgFinfoUtf8ToEucTw(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_euc_tw(m)
}
func Utf8ToEucTw(m *base.Module, l0 int32) int32 {
	return F_utf8_to_euc_tw(m, l0)
}
func PgMagicFuncUtf8AndEuc2004(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_euc2004(m)
}
func EucJis_2004_toUtf8(m *base.Module, l0 int32) int32 {
	return F_euc_jis_2004_to_utf8(m, l0)
}
func PgFinfoEucJis_2004_toUtf8(m *base.Module) int32 {
	return F_pg_finfo_euc_jis_2004_to_utf8(m)
}
func PgFinfoUtf8ToEucJis_2004(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_euc_jis_2004(m)
}
func Utf8ToEucJis_2004(m *base.Module, l0 int32) int32 {
	return F_utf8_to_euc_jis_2004(m, l0)
}
func PgMagicFuncUtf8AndGb18030(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_gb18030(m)
}
func Gb18030ToUtf8(m *base.Module, l0 int32) int32 {
	return F_gb18030_to_utf8(m, l0)
}
func PgFinfoGb18030ToUtf8(m *base.Module) int32 {
	return F_pg_finfo_gb18030_to_utf8(m)
}
func PgFinfoUtf8ToGb18030(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_gb18030(m)
}
func Utf8ToGb18030(m *base.Module, l0 int32) int32 {
	return F_utf8_to_gb18030(m, l0)
}
func PgMagicFuncUtf8AndGbk(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_gbk(m)
}
func GbkToUtf8(m *base.Module, l0 int32) int32 {
	return F_gbk_to_utf8(m, l0)
}
func PgFinfoGbkToUtf8(m *base.Module) int32 {
	return F_pg_finfo_gbk_to_utf8(m)
}
func PgFinfoUtf8ToGbk(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_gbk(m)
}
func Utf8ToGbk(m *base.Module, l0 int32) int32 {
	return F_utf8_to_gbk(m, l0)
}
func PgMagicFuncUtf8AndIso8859_1(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_iso8859_1(m)
}
func Iso8859_1_toUtf8(m *base.Module, l0 int32) int32 {
	return F_iso8859_1_to_utf8(m, l0)
}
func PgFinfoIso8859_1_toUtf8(m *base.Module) int32 {
	return F_pg_finfo_iso8859_1_to_utf8(m)
}
func PgFinfoUtf8ToIso8859_1(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_iso8859_1(m)
}
func Utf8ToIso8859_1(m *base.Module, l0 int32) int32 {
	return F_utf8_to_iso8859_1(m, l0)
}
func PgMagicFuncUtf8AndIso8859(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_iso8859_2(m)
}
func Iso8859ToUtf8(m *base.Module, l0 int32) int32 {
	return F_iso8859_to_utf8(m, l0)
}
func PgFinfoIso8859ToUtf8(m *base.Module) int32 {
	return F_pg_finfo_iso8859_to_utf8(m)
}
func PgFinfoUtf8ToIso8859(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_iso8859_2(m)
}
func Utf8ToIso8859(m *base.Module, l0 int32) int32 {
	return F_utf8_to_iso8859_2(m, l0)
}
func PgMagicFuncUtf8AndJohab(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_johab(m)
}
func JohabToUtf8(m *base.Module, l0 int32) int32 {
	return F_johab_to_utf8(m, l0)
}
func PgFinfoJohabToUtf8(m *base.Module) int32 {
	return F_pg_finfo_johab_to_utf8(m)
}
func PgFinfoUtf8ToJohab(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_johab(m)
}
func Utf8ToJohab(m *base.Module, l0 int32) int32 {
	return F_utf8_to_johab(m, l0)
}
func PgMagicFuncUtf8AndSjis(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_sjis(m)
}
func PgFinfoSjisToUtf8(m *base.Module) int32 {
	return F_pg_finfo_sjis_to_utf8(m)
}
func PgFinfoUtf8ToSjis(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_sjis(m)
}
func SjisToUtf8(m *base.Module, l0 int32) int32 {
	return F_sjis_to_utf8(m, l0)
}
func Utf8ToSjis(m *base.Module, l0 int32) int32 {
	return F_utf8_to_sjis(m, l0)
}
func PgMagicFuncUtf8AndSjis2004(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_sjis2004(m)
}
func PgFinfoShiftJis_2004_toUtf8(m *base.Module) int32 {
	return F_pg_finfo_shift_jis_2004_to_utf8(m)
}
func PgFinfoUtf8ToShiftJis_2004(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_shift_jis_2004(m)
}
func ShiftJis_2004_toUtf8(m *base.Module, l0 int32) int32 {
	return F_shift_jis_2004_to_utf8(m, l0)
}
func Utf8ToShiftJis_2004(m *base.Module, l0 int32) int32 {
	return F_utf8_to_shift_jis_2004(m, l0)
}
func PgMagicFuncUtf8AndUhc(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_uhc(m)
}
func PgFinfoUhcToUtf8(m *base.Module) int32 {
	return F_pg_finfo_uhc_to_utf8(m)
}
func PgFinfoUtf8ToUhc(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_uhc(m)
}
func UhcToUtf8(m *base.Module, l0 int32) int32 {
	return F_uhc_to_utf8(m, l0)
}
func Utf8ToUhc(m *base.Module, l0 int32) int32 {
	return F_utf8_to_uhc(m, l0)
}
func PgMagicFuncUtf8AndWin(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_win(m)
}
func PgFinfoUtf8ToWin(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_win(m)
}
func PgFinfoWinToUtf8(m *base.Module) int32 {
	return F_pg_finfo_win_to_utf8(m)
}
func Utf8ToWin(m *base.Module, l0 int32) int32 {
	return F_utf8_to_win(m, l0)
}
func WinToUtf8(m *base.Module, l0 int32) int32 {
	return F_win_to_utf8(m, l0)
}
func PgMagicFuncPgcrypto(m *base.Module) int32 {
	return F_Pg_magic_func_pgcrypto(m)
}
func PGInitPgcrypto(m *base.Module) {
	F__PG_init_pgcrypto(m)
}
func PgArmor(m *base.Module, l0 int32) int32 {
	return F_pg_armor(m, l0)
}
func PgCheckFipsmode(m *base.Module, l0 int32) int32 {
	return F_pg_numa_available(m, l0)
}
func PgCrypt(m *base.Module, l0 int32) int32 {
	return F_pg_crypt(m, l0)
}
func PgDearmor(m *base.Module, l0 int32) int32 {
	return F_pg_dearmor(m, l0)
}
func PgDecrypt(m *base.Module, l0 int32) int32 {
	return F_pg_decrypt(m, l0)
}
func PgDecryptIv(m *base.Module, l0 int32) int32 {
	return F_pg_decrypt_iv(m, l0)
}
func PgDigest(m *base.Module, l0 int32) int32 {
	return F_pg_digest(m, l0)
}
func PgEncrypt(m *base.Module, l0 int32) int32 {
	return F_pg_encrypt(m, l0)
}
func PgEncryptIv(m *base.Module, l0 int32) int32 {
	return F_pg_encrypt_iv(m, l0)
}
func PgFinfoPgArmor(m *base.Module) int32 {
	return F_pg_finfo_pg_armor(m)
}
func PgFinfoPgCheckFipsmode(m *base.Module) int32 {
	return F_pg_finfo_pg_check_fipsmode(m)
}
func PgFinfoPgCrypt(m *base.Module) int32 {
	return F_pg_finfo_pg_crypt(m)
}
func PgFinfoPgDearmor(m *base.Module) int32 {
	return F_pg_finfo_pg_dearmor(m)
}
func PgFinfoPgDecrypt(m *base.Module) int32 {
	return F_pg_finfo_pg_decrypt(m)
}
func PgFinfoPgDecryptIv(m *base.Module) int32 {
	return F_pg_finfo_pg_decrypt_iv(m)
}
func PgFinfoPgDigest(m *base.Module) int32 {
	return F_pg_finfo_pg_digest(m)
}
func PgFinfoPgEncrypt(m *base.Module) int32 {
	return F_pg_finfo_pg_encrypt(m)
}
func PgFinfoPgEncryptIv(m *base.Module) int32 {
	return F_pg_finfo_pg_encrypt_iv(m)
}
func PgFinfoPgGenSalt(m *base.Module) int32 {
	return F_pg_finfo_pg_gen_salt(m)
}
func PgFinfoPgGenSaltRounds(m *base.Module) int32 {
	return F_pg_finfo_pg_gen_salt_rounds(m)
}
func PgFinfoPgHmac(m *base.Module) int32 {
	return F_pg_finfo_pg_hmac(m)
}
func PgFinfoPgRandomBytes(m *base.Module) int32 {
	return F_pg_finfo_pg_random_bytes(m)
}
func PgFinfoPgRandomUuid(m *base.Module) int32 {
	return F_pg_finfo_pg_random_uuid(m)
}
func PgFinfoPgpArmorHeaders(m *base.Module) int32 {
	return F_pg_finfo_pgp_armor_headers(m)
}
func PgFinfoPgpKeyIdW(m *base.Module) int32 {
	return F_pg_finfo_pgp_key_id_w(m)
}
func PgFinfoPgpPubDecryptBytea(m *base.Module) int32 {
	return F_pg_finfo_pgp_pub_decrypt_bytea(m)
}
func PgFinfoPgpPubDecryptText(m *base.Module) int32 {
	return F_pg_finfo_pgp_pub_decrypt_text(m)
}
func PgFinfoPgpPubEncryptBytea(m *base.Module) int32 {
	return F_pg_finfo_pgp_pub_encrypt_bytea(m)
}
func PgFinfoPgpPubEncryptText(m *base.Module) int32 {
	return F_pg_finfo_pgp_pub_encrypt_text(m)
}
func PgFinfoPgpSymDecryptBytea(m *base.Module) int32 {
	return F_pg_finfo_pgp_sym_decrypt_bytea(m)
}
func PgFinfoPgpSymDecryptText(m *base.Module) int32 {
	return F_pg_finfo_pgp_sym_decrypt_text(m)
}
func PgFinfoPgpSymEncryptBytea(m *base.Module) int32 {
	return F_pg_finfo_pgp_sym_encrypt_bytea(m)
}
func PgFinfoPgpSymEncryptText(m *base.Module) int32 {
	return F_pg_finfo_pgp_sym_encrypt_text(m)
}
func PgGenSalt(m *base.Module, l0 int32) int32 {
	return F_pg_gen_salt(m, l0)
}
func PgGenSaltRounds(m *base.Module, l0 int32) int32 {
	return F_pg_gen_salt_rounds(m, l0)
}
func PgHmac(m *base.Module, l0 int32) int32 {
	return F_pg_hmac(m, l0)
}
func PgRandomBytes(m *base.Module, l0 int32) int32 {
	return F_pg_random_bytes(m, l0)
}
func PgRandomUuid(m *base.Module, l0 int32) int32 {
	return F_pg_random_uuid(m, l0)
}
func PgpArmorHeaders(m *base.Module, l0 int32) int32 {
	return F_pgp_armor_headers(m, l0)
}
func PgpKeyIdW(m *base.Module, l0 int32) int32 {
	return F_pgp_key_id_w(m, l0)
}
func PgpPubDecryptBytea(m *base.Module, l0 int32) int32 {
	return F_pgp_pub_decrypt_bytea(m, l0)
}
func PgpPubDecryptText(m *base.Module, l0 int32) int32 {
	return F_pgp_pub_decrypt_text(m, l0)
}
func PgpPubEncryptBytea(m *base.Module, l0 int32) int32 {
	return F_pgp_pub_encrypt_bytea(m, l0)
}
func PgpPubEncryptText(m *base.Module, l0 int32) int32 {
	return F_pgp_pub_encrypt_text(m, l0)
}
func PgpSymDecryptBytea(m *base.Module, l0 int32) int32 {
	return F_pgp_sym_decrypt_bytea(m, l0)
}
func PgpSymDecryptText(m *base.Module, l0 int32) int32 {
	return F_pgp_sym_decrypt_text(m, l0)
}
func PgpSymEncryptBytea(m *base.Module, l0 int32) int32 {
	return F_pgp_sym_encrypt_bytea(m, l0)
}
func PgpSymEncryptText(m *base.Module, l0 int32) int32 {
	return F_pgp_sym_encrypt_text(m, l0)
}
func PgMagicFuncCitext(m *base.Module) int32 {
	return F_Pg_magic_func_citext(m)
}
func CitextCmp(m *base.Module, l0 int32) int32 {
	return F_citext_cmp(m, l0)
}
func CitextEq(m *base.Module, l0 int32) int32 {
	return F_citext_eq(m, l0)
}
func CitextGe(m *base.Module, l0 int32) int32 {
	return F_citext_ge(m, l0)
}
func CitextGt(m *base.Module, l0 int32) int32 {
	return F_citext_gt(m, l0)
}
func CitextHash(m *base.Module, l0 int32) int32 {
	return F_citext_hash(m, l0)
}
func CitextHashExtended(m *base.Module, l0 int32) int32 {
	return F_citext_hash_extended(m, l0)
}
func CitextLarger(m *base.Module, l0 int32) int32 {
	return F_citext_larger(m, l0)
}
func CitextLe(m *base.Module, l0 int32) int32 {
	return F_citext_le(m, l0)
}
func CitextLt(m *base.Module, l0 int32) int32 {
	return F_citext_lt(m, l0)
}
func CitextNe(m *base.Module, l0 int32) int32 {
	return F_citext_ne(m, l0)
}
func CitextPatternCmp(m *base.Module, l0 int32) int32 {
	return F_citext_pattern_cmp(m, l0)
}
func CitextPatternGe(m *base.Module, l0 int32) int32 {
	return F_citext_pattern_ge(m, l0)
}
func CitextPatternGt(m *base.Module, l0 int32) int32 {
	return F_citext_pattern_gt(m, l0)
}
func CitextPatternLe(m *base.Module, l0 int32) int32 {
	return F_citext_pattern_le(m, l0)
}
func CitextPatternLt(m *base.Module, l0 int32) int32 {
	return F_citext_pattern_lt(m, l0)
}
func CitextSmaller(m *base.Module, l0 int32) int32 {
	return F_citext_smaller(m, l0)
}
func PgFinfoCitextCmp(m *base.Module) int32 {
	return F_pg_finfo_citext_cmp(m)
}
func PgFinfoCitextEq(m *base.Module) int32 {
	return F_pg_finfo_citext_eq(m)
}
func PgFinfoCitextGe(m *base.Module) int32 {
	return F_pg_finfo_citext_ge(m)
}
func PgFinfoCitextGt(m *base.Module) int32 {
	return F_pg_finfo_citext_gt(m)
}
func PgFinfoCitextHash(m *base.Module) int32 {
	return F_pg_finfo_citext_hash(m)
}
func PgFinfoCitextHashExtended(m *base.Module) int32 {
	return F_pg_finfo_citext_hash_extended(m)
}
func PgFinfoCitextLarger(m *base.Module) int32 {
	return F_pg_finfo_citext_larger(m)
}
func PgFinfoCitextLe(m *base.Module) int32 {
	return F_pg_finfo_citext_le(m)
}
func PgFinfoCitextLt(m *base.Module) int32 {
	return F_pg_finfo_citext_lt(m)
}
func PgFinfoCitextNe(m *base.Module) int32 {
	return F_pg_finfo_citext_ne(m)
}
func PgFinfoCitextPatternCmp(m *base.Module) int32 {
	return F_pg_finfo_citext_pattern_cmp(m)
}
func PgFinfoCitextPatternGe(m *base.Module) int32 {
	return F_pg_finfo_citext_pattern_ge(m)
}
func PgFinfoCitextPatternGt(m *base.Module) int32 {
	return F_pg_finfo_citext_pattern_gt(m)
}
func PgFinfoCitextPatternLe(m *base.Module) int32 {
	return F_pg_finfo_citext_pattern_le(m)
}
func PgFinfoCitextPatternLt(m *base.Module) int32 {
	return F_pg_finfo_citext_pattern_lt(m)
}
func PgFinfoCitextSmaller(m *base.Module) int32 {
	return F_pg_finfo_citext_smaller(m)
}
func PgMagicFuncPgTrgm(m *base.Module) int32 {
	return F_Pg_magic_func_pg_trgm(m)
}
func PGInitPgTrgm(m *base.Module) {
	F__PG_init_pg_trgm(m)
}
func GinExtractQueryTrgm(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_trgm(m, l0)
}
func GinExtractTrgm(m *base.Module, l0 int32) int32 {
	return F_gin_extract_trgm(m, l0)
}
func GinExtractValueTrgm(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_trgm(m, l0)
}
func GinTrgmConsistent(m *base.Module, l0 int32) int32 {
	return F_gin_trgm_consistent(m, l0)
}
func GinTrgmTriconsistent(m *base.Module, l0 int32) int32 {
	return F_gin_trgm_triconsistent(m, l0)
}
func GtrgmCompress(m *base.Module, l0 int32) int32 {
	return F_gtrgm_compress(m, l0)
}
func GtrgmConsistent(m *base.Module, l0 int32) int32 {
	return F_gtrgm_consistent(m, l0)
}
func GtrgmDecompress(m *base.Module, l0 int32) int32 {
	return F_gtrgm_decompress(m, l0)
}
func GtrgmDistance(m *base.Module, l0 int32) int32 {
	return F_gtrgm_distance(m, l0)
}
func GtrgmIn(m *base.Module, l0 int32) int32 {
	return F_gtrgm_in(m, l0)
}
func GtrgmOptions(m *base.Module, l0 int32) int32 {
	return F_gtrgm_options(m, l0)
}
func GtrgmOut(m *base.Module, l0 int32) int32 {
	return F_gtrgm_out(m, l0)
}
func GtrgmPenalty(m *base.Module, l0 int32) int32 {
	return F_gtrgm_penalty(m, l0)
}
func GtrgmPicksplit(m *base.Module, l0 int32) int32 {
	return F_gtrgm_picksplit(m, l0)
}
func GtrgmSame(m *base.Module, l0 int32) int32 {
	return F_gtrgm_same(m, l0)
}
func GtrgmUnion(m *base.Module, l0 int32) int32 {
	return F_gtrgm_union(m, l0)
}
func PgFinfoGinExtractQueryTrgm(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_trgm(m)
}
func PgFinfoGinExtractTrgm(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_trgm(m)
}
func PgFinfoGinExtractValueTrgm(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_trgm(m)
}
func PgFinfoGinTrgmConsistent(m *base.Module) int32 {
	return F_pg_finfo_gin_trgm_consistent(m)
}
func PgFinfoGinTrgmTriconsistent(m *base.Module) int32 {
	return F_pg_finfo_gin_trgm_triconsistent(m)
}
func PgFinfoGtrgmCompress(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_compress(m)
}
func PgFinfoGtrgmConsistent(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_consistent(m)
}
func PgFinfoGtrgmDecompress(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_decompress(m)
}
func PgFinfoGtrgmDistance(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_distance(m)
}
func PgFinfoGtrgmIn(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_in(m)
}
func PgFinfoGtrgmOptions(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_options(m)
}
func PgFinfoGtrgmOut(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_out(m)
}
func PgFinfoGtrgmPenalty(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_penalty(m)
}
func PgFinfoGtrgmPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_picksplit(m)
}
func PgFinfoGtrgmSame(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_same(m)
}
func PgFinfoGtrgmUnion(m *base.Module) int32 {
	return F_pg_finfo_gtrgm_union(m)
}
func PgFinfoSetLimit(m *base.Module) int32 {
	return F_pg_finfo_set_limit(m)
}
func PgFinfoShowLimit(m *base.Module) int32 {
	return F_pg_finfo_show_limit(m)
}
func PgFinfoShowTrgm(m *base.Module) int32 {
	return F_pg_finfo_show_trgm(m)
}
func PgFinfoSimilarity(m *base.Module) int32 {
	return F_pg_finfo_similarity(m)
}
func PgFinfoSimilarityDist(m *base.Module) int32 {
	return F_pg_finfo_similarity_dist(m)
}
func PgFinfoSimilarityOp(m *base.Module) int32 {
	return F_pg_finfo_similarity_op(m)
}
func PgFinfoStrictWordSimilarity(m *base.Module) int32 {
	return F_pg_finfo_strict_word_similarity(m)
}
func PgFinfoStrictWordSimilarityCommutatorOp(m *base.Module) int32 {
	return F_pg_finfo_strict_word_similarity_commutator_op(m)
}
func PgFinfoStrictWordSimilarityDistCommutatorOp(m *base.Module) int32 {
	return F_pg_finfo_strict_word_similarity_dist_commutator_op(m)
}
func PgFinfoStrictWordSimilarityDistOp(m *base.Module) int32 {
	return F_pg_finfo_strict_word_similarity_dist_op(m)
}
func PgFinfoStrictWordSimilarityOp(m *base.Module) int32 {
	return F_pg_finfo_strict_word_similarity_op(m)
}
func PgFinfoWordSimilarity(m *base.Module) int32 {
	return F_pg_finfo_word_similarity(m)
}
func PgFinfoWordSimilarityCommutatorOp(m *base.Module) int32 {
	return F_pg_finfo_word_similarity_commutator_op(m)
}
func PgFinfoWordSimilarityDistCommutatorOp(m *base.Module) int32 {
	return F_pg_finfo_word_similarity_dist_commutator_op(m)
}
func PgFinfoWordSimilarityDistOp(m *base.Module) int32 {
	return F_pg_finfo_word_similarity_dist_op(m)
}
func PgFinfoWordSimilarityOp(m *base.Module) int32 {
	return F_pg_finfo_word_similarity_op(m)
}
func SetLimit(m *base.Module, l0 int32) int32 {
	return F_set_limit(m, l0)
}
func ShowLimit(m *base.Module, l0 int32) int32 {
	return F_show_limit(m, l0)
}
func ShowTrgm(m *base.Module, l0 int32) int32 {
	return F_show_trgm(m, l0)
}
func Similarity(m *base.Module, l0 int32) int32 {
	return F_similarity(m, l0)
}
func SimilarityDist(m *base.Module, l0 int32) int32 {
	return F_similarity_dist(m, l0)
}
func SimilarityOp(m *base.Module, l0 int32) int32 {
	return F_similarity_op(m, l0)
}
func StrictWordSimilarity(m *base.Module, l0 int32) int32 {
	return F_strict_word_similarity(m, l0)
}
func StrictWordSimilarityCommutatorOp(m *base.Module, l0 int32) int32 {
	return F_strict_word_similarity_commutator_op(m, l0)
}
func StrictWordSimilarityDistCommutatorOp(m *base.Module, l0 int32) int32 {
	return F_strict_word_similarity_dist_commutator_op(m, l0)
}
func StrictWordSimilarityDistOp(m *base.Module, l0 int32) int32 {
	return F_strict_word_similarity_dist_op(m, l0)
}
func StrictWordSimilarityOp(m *base.Module, l0 int32) int32 {
	return F_strict_word_similarity_op(m, l0)
}
func WordSimilarity(m *base.Module, l0 int32) int32 {
	return F_word_similarity(m, l0)
}
func WordSimilarityCommutatorOp(m *base.Module, l0 int32) int32 {
	return F_word_similarity_commutator_op(m, l0)
}
func WordSimilarityDistCommutatorOp(m *base.Module, l0 int32) int32 {
	return F_word_similarity_dist_commutator_op(m, l0)
}
func WordSimilarityDistOp(m *base.Module, l0 int32) int32 {
	return F_word_similarity_dist_op(m, l0)
}
func WordSimilarityOp(m *base.Module, l0 int32) int32 {
	return F_word_similarity_op(m, l0)
}
func PgMagicFuncHstore(m *base.Module) int32 {
	return F_Pg_magic_func_hstore(m)
}
func Akeys(m *base.Module, l0 int32) int32 {
	return F_akeys(m, l0)
}
func Avals(m *base.Module, l0 int32) int32 {
	return F_avals(m, l0)
}
func Defined(m *base.Module, l0 int32) int32 {
	return F_defined(m, l0)
}
func Delete(m *base.Module, l0 int32) int32 {
	return F_delete(m, l0)
}
func Each(m *base.Module, l0 int32) int32 {
	return F_each(m, l0)
}
func Exists(m *base.Module, l0 int32) int32 {
	return F_exists(m, l0)
}
func Fetchval(m *base.Module, l0 int32) int32 {
	return F_fetchval(m, l0)
}
func GhstoreCompress(m *base.Module, l0 int32) int32 {
	return F_ghstore_compress(m, l0)
}
func GhstoreConsistent(m *base.Module, l0 int32) int32 {
	return F_ghstore_consistent(m, l0)
}
func GhstoreDecompress(m *base.Module, l0 int32) int32 {
	return F_float4up(m, l0)
}
func GhstoreIn(m *base.Module, l0 int32) int32 {
	return F_ghstore_in(m, l0)
}
func GhstoreOptions(m *base.Module, l0 int32) int32 {
	return F_ghstore_options(m, l0)
}
func GhstoreOut(m *base.Module, l0 int32) int32 {
	return F_ghstore_out(m, l0)
}
func GhstorePenalty(m *base.Module, l0 int32) int32 {
	return F_ghstore_penalty(m, l0)
}
func GhstorePicksplit(m *base.Module, l0 int32) int32 {
	return F_ghstore_picksplit(m, l0)
}
func GhstoreSame(m *base.Module, l0 int32) int32 {
	return F_ghstore_same(m, l0)
}
func GhstoreUnion(m *base.Module, l0 int32) int32 {
	return F_ghstore_union(m, l0)
}
func GinConsistentHstore(m *base.Module, l0 int32) int32 {
	return F_gin_consistent_hstore(m, l0)
}
func GinExtractHstore(m *base.Module, l0 int32) int32 {
	return F_gin_extract_hstore(m, l0)
}
func GinExtractHstoreQuery(m *base.Module, l0 int32) int32 {
	return F_gin_extract_hstore_query(m, l0)
}
func HsConcat(m *base.Module, l0 int32) int32 {
	return F_hs_concat(m, l0)
}
func HsContained(m *base.Module, l0 int32) int32 {
	return F_hs_contained(m, l0)
}
func HsContains(m *base.Module, l0 int32) int32 {
	return F_hs_contains(m, l0)
}
func HstoreArrayToPairs(m *base.Module, l0 int32, l1 int32) int32 {
	return F_hstoreArrayToPairs(m, l0, l1)
}
func HstoreCheckKeyLen(m *base.Module, l0 int32) int32 {
	return F_hstoreCheckKeyLen(m, l0)
}
func HstoreCheckValLen(m *base.Module, l0 int32) int32 {
	return F_hstoreCheckValLen(m, l0)
}
func HstoreFindKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return F_hstoreFindKey(m, l0, l1, l2, l3)
}
func HstorePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_hstorePairs(m, l0, l1, l2)
}
func HstoreUniquePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F_hstoreUniquePairs(m, l0, l1, l2)
}
func HstoreUpgrade(m *base.Module, l0 int32) int32 {
	return F_hstoreUpgrade(m, l0)
}
func HstoreAkeys(m *base.Module, l0 int32) int32 {
	return F_hstore_akeys(m, l0)
}
func HstoreAvals(m *base.Module, l0 int32) int32 {
	return F_hstore_avals(m, l0)
}
func HstoreCmp(m *base.Module, l0 int32) int32 {
	return F_hstore_cmp(m, l0)
}
func HstoreConcat(m *base.Module, l0 int32) int32 {
	return F_hstore_concat(m, l0)
}
func HstoreContained(m *base.Module, l0 int32) int32 {
	return F_hs_contained(m, l0)
}
func HstoreContains(m *base.Module, l0 int32) int32 {
	return F_hstore_contains(m, l0)
}
func HstoreDefined(m *base.Module, l0 int32) int32 {
	return F_hstore_defined(m, l0)
}
func HstoreDelete(m *base.Module, l0 int32) int32 {
	return F_hstore_delete(m, l0)
}
func HstoreDeleteArray(m *base.Module, l0 int32) int32 {
	return F_hstore_delete_array(m, l0)
}
func HstoreDeleteHstore(m *base.Module, l0 int32) int32 {
	return F_hstore_delete_hstore(m, l0)
}
func HstoreEach(m *base.Module, l0 int32) int32 {
	return F_hstore_each(m, l0)
}
func HstoreEq(m *base.Module, l0 int32) int32 {
	return F_hstore_eq(m, l0)
}
func HstoreExists(m *base.Module, l0 int32) int32 {
	return F_hstore_exists(m, l0)
}
func HstoreExistsAll(m *base.Module, l0 int32) int32 {
	return F_hstore_exists_all(m, l0)
}
func HstoreExistsAny(m *base.Module, l0 int32) int32 {
	return F_hstore_exists_any(m, l0)
}
func HstoreFetchval(m *base.Module, l0 int32) int32 {
	return F_hstore_fetchval(m, l0)
}
func HstoreFromArray(m *base.Module, l0 int32) int32 {
	return F_hstore_from_array(m, l0)
}
func HstoreFromArrays(m *base.Module, l0 int32) int32 {
	return F_hstore_from_arrays(m, l0)
}
func HstoreFromRecord(m *base.Module, l0 int32) int32 {
	return F_hstore_from_record(m, l0)
}
func HstoreFromText(m *base.Module, l0 int32) int32 {
	return F_hstore_from_text(m, l0)
}
func HstoreGe(m *base.Module, l0 int32) int32 {
	return F_hstore_ge(m, l0)
}
func HstoreGt(m *base.Module, l0 int32) int32 {
	return F_hstore_gt(m, l0)
}
func HstoreHash(m *base.Module, l0 int32) int32 {
	return F_hstore_hash(m, l0)
}
func HstoreHashExtended(m *base.Module, l0 int32) int32 {
	return F_hstore_hash_extended(m, l0)
}
func HstoreIn(m *base.Module, l0 int32) int32 {
	return F_hstore_in(m, l0)
}
func HstoreLe(m *base.Module, l0 int32) int32 {
	return F_hstore_le(m, l0)
}
func HstoreLt(m *base.Module, l0 int32) int32 {
	return F_hstore_lt(m, l0)
}
func HstoreNe(m *base.Module, l0 int32) int32 {
	return F_hstore_ne(m, l0)
}
func HstoreOut(m *base.Module, l0 int32) int32 {
	return F_hstore_out(m, l0)
}
func HstorePopulateRecord(m *base.Module, l0 int32) int32 {
	return F_hstore_populate_record(m, l0)
}
func HstoreRecv(m *base.Module, l0 int32) int32 {
	return F_hstore_recv(m, l0)
}
func HstoreSend(m *base.Module, l0 int32) int32 {
	return F_hstore_send(m, l0)
}
func HstoreSkeys(m *base.Module, l0 int32) int32 {
	return F_hstore_skeys(m, l0)
}
func HstoreSliceToArray(m *base.Module, l0 int32) int32 {
	return F_hstore_slice_to_array(m, l0)
}
func HstoreSliceToHstore(m *base.Module, l0 int32) int32 {
	return F_hstore_slice_to_hstore(m, l0)
}
func HstoreSubscriptHandler(m *base.Module, l0 int32) int32 {
	return F_hstore_subscript_handler(m, l0)
}
func HstoreSvals(m *base.Module, l0 int32) int32 {
	return F_hstore_svals(m, l0)
}
func HstoreToArray(m *base.Module, l0 int32) int32 {
	return F_hstore_to_array(m, l0)
}
func HstoreToJson(m *base.Module, l0 int32) int32 {
	return F_hstore_to_json(m, l0)
}
func HstoreToJsonLoose(m *base.Module, l0 int32) int32 {
	return F_hstore_to_json_loose(m, l0)
}
func HstoreToJsonb(m *base.Module, l0 int32) int32 {
	return F_hstore_to_jsonb(m, l0)
}
func HstoreToJsonbLoose(m *base.Module, l0 int32) int32 {
	return F_hstore_to_jsonb_loose(m, l0)
}
func HstoreToMatrix(m *base.Module, l0 int32) int32 {
	return F_hstore_to_matrix(m, l0)
}
func HstoreVersionDiag(m *base.Module, l0 int32) int32 {
	return F_hstore_version_diag(m, l0)
}
func PgFinfoAkeys(m *base.Module) int32 {
	return F_pg_finfo_akeys(m)
}
func PgFinfoAvals(m *base.Module) int32 {
	return F_pg_finfo_avals(m)
}
func PgFinfoDefined(m *base.Module) int32 {
	return F_pg_finfo_defined(m)
}
func PgFinfoDelete(m *base.Module) int32 {
	return F_pg_finfo_delete(m)
}
func PgFinfoEach(m *base.Module) int32 {
	return F_pg_finfo_each(m)
}
func PgFinfoExists(m *base.Module) int32 {
	return F_pg_finfo_exists(m)
}
func PgFinfoFetchval(m *base.Module) int32 {
	return F_pg_finfo_fetchval(m)
}
func PgFinfoGhstoreCompress(m *base.Module) int32 {
	return F_pg_finfo_ghstore_compress(m)
}
func PgFinfoGhstoreConsistent(m *base.Module) int32 {
	return F_pg_finfo_ghstore_consistent(m)
}
func PgFinfoGhstoreDecompress(m *base.Module) int32 {
	return F_pg_finfo_ghstore_decompress(m)
}
func PgFinfoGhstoreIn(m *base.Module) int32 {
	return F_pg_finfo_ghstore_in(m)
}
func PgFinfoGhstoreOptions(m *base.Module) int32 {
	return F_pg_finfo_ghstore_options(m)
}
func PgFinfoGhstoreOut(m *base.Module) int32 {
	return F_pg_finfo_ghstore_out(m)
}
func PgFinfoGhstorePenalty(m *base.Module) int32 {
	return F_pg_finfo_ghstore_penalty(m)
}
func PgFinfoGhstorePicksplit(m *base.Module) int32 {
	return F_pg_finfo_ghstore_picksplit(m)
}
func PgFinfoGhstoreSame(m *base.Module) int32 {
	return F_pg_finfo_ghstore_same(m)
}
func PgFinfoGhstoreUnion(m *base.Module) int32 {
	return F_pg_finfo_ghstore_union(m)
}
func PgFinfoGinConsistentHstore(m *base.Module) int32 {
	return F_pg_finfo_gin_consistent_hstore(m)
}
func PgFinfoGinExtractHstore(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_hstore(m)
}
func PgFinfoGinExtractHstoreQuery(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_hstore_query(m)
}
func PgFinfoHsConcat(m *base.Module) int32 {
	return F_pg_finfo_hs_concat(m)
}
func PgFinfoHsContained(m *base.Module) int32 {
	return F_pg_finfo_hs_contained(m)
}
func PgFinfoHsContains(m *base.Module) int32 {
	return F_pg_finfo_hs_contains(m)
}
func PgFinfoHstoreAkeys(m *base.Module) int32 {
	return F_pg_finfo_hstore_akeys(m)
}
func PgFinfoHstoreAvals(m *base.Module) int32 {
	return F_pg_finfo_hstore_avals(m)
}
func PgFinfoHstoreCmp(m *base.Module) int32 {
	return F_pg_finfo_hstore_cmp(m)
}
func PgFinfoHstoreConcat(m *base.Module) int32 {
	return F_pg_finfo_hstore_concat(m)
}
func PgFinfoHstoreContained(m *base.Module) int32 {
	return F_pg_finfo_hstore_contained(m)
}
func PgFinfoHstoreContains(m *base.Module) int32 {
	return F_pg_finfo_hstore_contains(m)
}
func PgFinfoHstoreDefined(m *base.Module) int32 {
	return F_pg_finfo_hstore_defined(m)
}
func PgFinfoHstoreDelete(m *base.Module) int32 {
	return F_pg_finfo_hstore_delete(m)
}
func PgFinfoHstoreDeleteArray(m *base.Module) int32 {
	return F_pg_finfo_hstore_delete_array(m)
}
func PgFinfoHstoreDeleteHstore(m *base.Module) int32 {
	return F_pg_finfo_hstore_delete_hstore(m)
}
func PgFinfoHstoreEach(m *base.Module) int32 {
	return F_pg_finfo_hstore_each(m)
}
func PgFinfoHstoreEq(m *base.Module) int32 {
	return F_pg_finfo_hstore_eq(m)
}
func PgFinfoHstoreExists(m *base.Module) int32 {
	return F_pg_finfo_hstore_exists(m)
}
func PgFinfoHstoreExistsAll(m *base.Module) int32 {
	return F_pg_finfo_hstore_exists_all(m)
}
func PgFinfoHstoreExistsAny(m *base.Module) int32 {
	return F_pg_finfo_hstore_exists_any(m)
}
func PgFinfoHstoreFetchval(m *base.Module) int32 {
	return F_pg_finfo_hstore_fetchval(m)
}
func PgFinfoHstoreFromArray(m *base.Module) int32 {
	return F_pg_finfo_hstore_from_array(m)
}
func PgFinfoHstoreFromArrays(m *base.Module) int32 {
	return F_pg_finfo_hstore_from_arrays(m)
}
func PgFinfoHstoreFromRecord(m *base.Module) int32 {
	return F_pg_finfo_hstore_from_record(m)
}
func PgFinfoHstoreFromText(m *base.Module) int32 {
	return F_pg_finfo_hstore_from_text(m)
}
func PgFinfoHstoreGe(m *base.Module) int32 {
	return F_pg_finfo_hstore_ge(m)
}
func PgFinfoHstoreGt(m *base.Module) int32 {
	return F_pg_finfo_hstore_gt(m)
}
func PgFinfoHstoreHash(m *base.Module) int32 {
	return F_pg_finfo_hstore_hash(m)
}
func PgFinfoHstoreHashExtended(m *base.Module) int32 {
	return F_pg_finfo_hstore_hash_extended(m)
}
func PgFinfoHstoreIn(m *base.Module) int32 {
	return F_pg_finfo_hstore_in(m)
}
func PgFinfoHstoreLe(m *base.Module) int32 {
	return F_pg_finfo_hstore_le(m)
}
func PgFinfoHstoreLt(m *base.Module) int32 {
	return F_pg_finfo_hstore_lt(m)
}
func PgFinfoHstoreNe(m *base.Module) int32 {
	return F_pg_finfo_hstore_ne(m)
}
func PgFinfoHstoreOut(m *base.Module) int32 {
	return F_pg_finfo_hstore_out(m)
}
func PgFinfoHstorePopulateRecord(m *base.Module) int32 {
	return F_pg_finfo_hstore_populate_record(m)
}
func PgFinfoHstoreRecv(m *base.Module) int32 {
	return F_pg_finfo_hstore_recv(m)
}
func PgFinfoHstoreSend(m *base.Module) int32 {
	return F_pg_finfo_hstore_send(m)
}
func PgFinfoHstoreSkeys(m *base.Module) int32 {
	return F_pg_finfo_hstore_skeys(m)
}
func PgFinfoHstoreSliceToArray(m *base.Module) int32 {
	return F_pg_finfo_hstore_slice_to_array(m)
}
func PgFinfoHstoreSliceToHstore(m *base.Module) int32 {
	return F_pg_finfo_hstore_slice_to_hstore(m)
}
func PgFinfoHstoreSubscriptHandler(m *base.Module) int32 {
	return F_pg_finfo_hstore_subscript_handler(m)
}
func PgFinfoHstoreSvals(m *base.Module) int32 {
	return F_pg_finfo_hstore_svals(m)
}
func PgFinfoHstoreToArray(m *base.Module) int32 {
	return F_pg_finfo_hstore_to_array(m)
}
func PgFinfoHstoreToJson(m *base.Module) int32 {
	return F_pg_finfo_hstore_to_json(m)
}
func PgFinfoHstoreToJsonLoose(m *base.Module) int32 {
	return F_pg_finfo_hstore_to_json_loose(m)
}
func PgFinfoHstoreToJsonb(m *base.Module) int32 {
	return F_pg_finfo_hstore_to_jsonb(m)
}
func PgFinfoHstoreToJsonbLoose(m *base.Module) int32 {
	return F_pg_finfo_hstore_to_jsonb_loose(m)
}
func PgFinfoHstoreToMatrix(m *base.Module) int32 {
	return F_pg_finfo_hstore_to_matrix(m)
}
func PgFinfoHstoreVersionDiag(m *base.Module) int32 {
	return F_pg_finfo_hstore_version_diag(m)
}
func PgFinfoSkeys(m *base.Module) int32 {
	return F_pg_finfo_skeys(m)
}
func PgFinfoSvals(m *base.Module) int32 {
	return F_pg_finfo_svals(m)
}
func PgFinfoTconvert(m *base.Module) int32 {
	return F_pg_finfo_tconvert(m)
}
func Skeys(m *base.Module, l0 int32) int32 {
	return F_skeys(m, l0)
}
func Svals(m *base.Module, l0 int32) int32 {
	return F_svals(m, l0)
}
func Tconvert(m *base.Module, l0 int32) int32 {
	return F_tconvert(m, l0)
}
func PgMagicFuncLtree(m *base.Module) int32 {
	return F_Pg_magic_func_ltree(m)
}
func Lca(m *base.Module, l0 int32) int32 {
	return F__lca(m, l0)
}
func LtQRegex(m *base.Module, l0 int32) int32 {
	return F__lt_q_regex(m, l0)
}
func LtQRregex(m *base.Module, l0 int32) int32 {
	return F__lt_q_rregex(m, l0)
}
func LtqExtractRegex(m *base.Module, l0 int32) int32 {
	return F__ltq_extract_regex(m, l0)
}
func LtqRegex(m *base.Module, l0 int32) int32 {
	return F__ltq_regex(m, l0)
}
func LtqRregex(m *base.Module, l0 int32) int32 {
	return F__ltq_rregex(m, l0)
}
func LtreeCompress(m *base.Module, l0 int32) int32 {
	return F__ltree_compress(m, l0)
}
func LtreeConsistent(m *base.Module, l0 int32) int32 {
	return F__ltree_consistent(m, l0)
}
func LtreeExtractIsparent(m *base.Module, l0 int32) int32 {
	return F__ltree_extract_isparent(m, l0)
}
func LtreeExtractRisparent(m *base.Module, l0 int32) int32 {
	return F__ltree_extract_risparent(m, l0)
}
func LtreeGistOptions(m *base.Module, l0 int32) int32 {
	return F__ltree_gist_options(m, l0)
}
func LtreeIsparent(m *base.Module, l0 int32) int32 {
	return F__ltree_isparent(m, l0)
}
func LtreePenalty(m *base.Module, l0 int32) int32 {
	return F__ltree_penalty(m, l0)
}
func LtreePicksplit(m *base.Module, l0 int32) int32 {
	return F__ltree_picksplit(m, l0)
}
func LtreeRIsparent(m *base.Module, l0 int32) int32 {
	return F__ltree_r_isparent(m, l0)
}
func LtreeRRisparent(m *base.Module, l0 int32) int32 {
	return F__ltree_r_risparent(m, l0)
}
func LtreeRisparent(m *base.Module, l0 int32) int32 {
	return F__ltree_risparent(m, l0)
}
func LtreeSame(m *base.Module, l0 int32) int32 {
	return F__ltree_same(m, l0)
}
func LtreeUnion(m *base.Module, l0 int32) int32 {
	return F__ltree_union(m, l0)
}
func LtxtqExec(m *base.Module, l0 int32) int32 {
	return F__ltxtq_exec(m, l0)
}
func LtxtqExtractExec(m *base.Module, l0 int32) int32 {
	return F__ltxtq_extract_exec(m, l0)
}
func LtxtqRexec(m *base.Module, l0 int32) int32 {
	return F__ltxtq_rexec(m, l0)
}
func HashLtree(m *base.Module, l0 int32) int32 {
	return F_hash_ltree(m, l0)
}
func HashLtreeExtended(m *base.Module, l0 int32) int32 {
	return F_hash_ltree_extended(m, l0)
}
func Lca_lca(m *base.Module, l0 int32) int32 {
	return F_lca(m, l0)
}
func LqueryIn(m *base.Module, l0 int32) int32 {
	return F_lquery_in(m, l0)
}
func LqueryOut(m *base.Module, l0 int32) int32 {
	return F_lquery_out(m, l0)
}
func LqueryRecv(m *base.Module, l0 int32) int32 {
	return F_lquery_recv(m, l0)
}
func LquerySend(m *base.Module, l0 int32) int32 {
	return F_lquery_send(m, l0)
}
func LtQRegex_lt_q_regex(m *base.Module, l0 int32) int32 {
	return F_lt_q_regex(m, l0)
}
func LtQRregex_lt_q_rregex(m *base.Module, l0 int32) int32 {
	return F_lt_q_rregex(m, l0)
}
func LtqRegex_ltq_regex(m *base.Module, l0 int32) int32 {
	return F_ltq_regex(m, l0)
}
func LtqRregex_ltq_rregex(m *base.Module, l0 int32) int32 {
	return F_ltq_rregex(m, l0)
}
func Ltree2text(m *base.Module, l0 int32) int32 {
	return F_ltree2text(m, l0)
}
func LtreeAddltree(m *base.Module, l0 int32) int32 {
	return F_ltree_addltree(m, l0)
}
func LtreeAddtext(m *base.Module, l0 int32) int32 {
	return F_ltree_addtext(m, l0)
}
func LtreeCmp(m *base.Module, l0 int32) int32 {
	return F_ltree_cmp(m, l0)
}
func LtreeCompress_ltree_compress(m *base.Module, l0 int32) int32 {
	return F_ltree_compress(m, l0)
}
func LtreeConsistent_ltree_consistent(m *base.Module, l0 int32) int32 {
	return F_ltree_consistent(m, l0)
}
func LtreeDecompress(m *base.Module, l0 int32) int32 {
	return F_gtsvector_decompress(m, l0)
}
func LtreeEq(m *base.Module, l0 int32) int32 {
	return F_ltree_eq(m, l0)
}
func LtreeGe(m *base.Module, l0 int32) int32 {
	return F_ltree_ge(m, l0)
}
func LtreeGistIn(m *base.Module, l0 int32) int32 {
	return F_ltree_gist_in(m, l0)
}
func LtreeGistOptions_ltree_gist_options(m *base.Module, l0 int32) int32 {
	return F_ltree_gist_options(m, l0)
}
func LtreeGistOut(m *base.Module, l0 int32) int32 {
	return F_ltree_gist_out(m, l0)
}
func LtreeGt(m *base.Module, l0 int32) int32 {
	return F_ltree_gt(m, l0)
}
func LtreeIn(m *base.Module, l0 int32) int32 {
	return F_ltree_in(m, l0)
}
func LtreeIndex(m *base.Module, l0 int32) int32 {
	return F_ltree_index(m, l0)
}
func LtreeIsparent_ltree_isparent(m *base.Module, l0 int32) int32 {
	return F_ltree_isparent(m, l0)
}
func LtreeLe(m *base.Module, l0 int32) int32 {
	return F_ltree_le(m, l0)
}
func LtreeLt(m *base.Module, l0 int32) int32 {
	return F_ltree_lt(m, l0)
}
func LtreeNe(m *base.Module, l0 int32) int32 {
	return F_ltree_ne(m, l0)
}
func LtreeOut(m *base.Module, l0 int32) int32 {
	return F_ltree_out(m, l0)
}
func LtreePenalty_ltree_penalty(m *base.Module, l0 int32) int32 {
	return F_ltree_penalty(m, l0)
}
func LtreePicksplit_ltree_picksplit(m *base.Module, l0 int32) int32 {
	return F_ltree_picksplit(m, l0)
}
func LtreeRecv(m *base.Module, l0 int32) int32 {
	return F_ltree_recv(m, l0)
}
func LtreeRisparent_ltree_risparent(m *base.Module, l0 int32) int32 {
	return F_ltree_risparent(m, l0)
}
func LtreeSame_ltree_same(m *base.Module, l0 int32) int32 {
	return F_ltree_same(m, l0)
}
func LtreeSend(m *base.Module, l0 int32) int32 {
	return F_ltree_send(m, l0)
}
func LtreeTextadd(m *base.Module, l0 int32) int32 {
	return F_ltree_textadd(m, l0)
}
func LtreeUnion_ltree_union(m *base.Module, l0 int32) int32 {
	return F_ltree_union(m, l0)
}
func Ltreeparentsel(m *base.Module, l0 int32) int32 {
	return F_ltreeparentsel(m, l0)
}
func LtxtqExec_ltxtq_exec(m *base.Module, l0 int32) int32 {
	return F_ltxtq_exec(m, l0)
}
func LtxtqIn(m *base.Module, l0 int32) int32 {
	return F_ltxtq_in(m, l0)
}
func LtxtqOut(m *base.Module, l0 int32) int32 {
	return F_ltxtq_out(m, l0)
}
func LtxtqRecv(m *base.Module, l0 int32) int32 {
	return F_ltxtq_recv(m, l0)
}
func LtxtqRexec_ltxtq_rexec(m *base.Module, l0 int32) int32 {
	return F_ltxtq_rexec(m, l0)
}
func LtxtqSend(m *base.Module, l0 int32) int32 {
	return F_ltxtq_send(m, l0)
}
func Nlevel(m *base.Module, l0 int32) int32 {
	return F_nlevel(m, l0)
}
func PgFinfoLca(m *base.Module) int32 {
	return F_pg_finfo__lca(m)
}
func PgFinfoLtQRegex(m *base.Module) int32 {
	return F_pg_finfo__lt_q_regex(m)
}
func PgFinfoLtQRregex(m *base.Module) int32 {
	return F_pg_finfo__lt_q_rregex(m)
}
func PgFinfoLtqExtractRegex(m *base.Module) int32 {
	return F_pg_finfo__ltq_extract_regex(m)
}
func PgFinfoLtqRegex(m *base.Module) int32 {
	return F_pg_finfo__ltq_regex(m)
}
func PgFinfoLtqRregex(m *base.Module) int32 {
	return F_pg_finfo__ltq_rregex(m)
}
func PgFinfoLtreeCompress(m *base.Module) int32 {
	return F_pg_finfo__ltree_compress(m)
}
func PgFinfoLtreeConsistent(m *base.Module) int32 {
	return F_pg_finfo__ltree_consistent(m)
}
func PgFinfoLtreeExtractIsparent(m *base.Module) int32 {
	return F_pg_finfo__ltree_extract_isparent(m)
}
func PgFinfoLtreeExtractRisparent(m *base.Module) int32 {
	return F_pg_finfo__ltree_extract_risparent(m)
}
func PgFinfoLtreeGistOptions(m *base.Module) int32 {
	return F_pg_finfo__ltree_gist_options(m)
}
func PgFinfoLtreeIsparent(m *base.Module) int32 {
	return F_pg_finfo__ltree_isparent(m)
}
func PgFinfoLtreePenalty(m *base.Module) int32 {
	return F_pg_finfo__ltree_penalty(m)
}
func PgFinfoLtreePicksplit(m *base.Module) int32 {
	return F_pg_finfo__ltree_picksplit(m)
}
func PgFinfoLtreeRIsparent(m *base.Module) int32 {
	return F_pg_finfo__ltree_r_isparent(m)
}
func PgFinfoLtreeRRisparent(m *base.Module) int32 {
	return F_pg_finfo__ltree_r_risparent(m)
}
func PgFinfoLtreeRisparent(m *base.Module) int32 {
	return F_pg_finfo__ltree_risparent(m)
}
func PgFinfoLtreeSame(m *base.Module) int32 {
	return F_pg_finfo__ltree_same(m)
}
func PgFinfoLtreeUnion(m *base.Module) int32 {
	return F_pg_finfo__ltree_union(m)
}
func PgFinfoLtxtqExec(m *base.Module) int32 {
	return F_pg_finfo__ltxtq_exec(m)
}
func PgFinfoLtxtqExtractExec(m *base.Module) int32 {
	return F_pg_finfo__ltxtq_extract_exec(m)
}
func PgFinfoLtxtqRexec(m *base.Module) int32 {
	return F_pg_finfo__ltxtq_rexec(m)
}
func PgFinfoHashLtree(m *base.Module) int32 {
	return F_pg_finfo_hash_ltree(m)
}
func PgFinfoHashLtreeExtended(m *base.Module) int32 {
	return F_pg_finfo_hash_ltree_extended(m)
}
func PgFinfoLca_pg_finfo_lca(m *base.Module) int32 {
	return F_pg_finfo_lca(m)
}
func PgFinfoLqueryIn(m *base.Module) int32 {
	return F_pg_finfo_lquery_in(m)
}
func PgFinfoLqueryOut(m *base.Module) int32 {
	return F_pg_finfo_lquery_out(m)
}
func PgFinfoLqueryRecv(m *base.Module) int32 {
	return F_pg_finfo_lquery_recv(m)
}
func PgFinfoLquerySend(m *base.Module) int32 {
	return F_pg_finfo_lquery_send(m)
}
func PgFinfoLtQRegex_pg_finfo_lt_q_regex(m *base.Module) int32 {
	return F_pg_finfo_lt_q_regex(m)
}
func PgFinfoLtQRregex_pg_finfo_lt_q_rregex(m *base.Module) int32 {
	return F_pg_finfo_lt_q_rregex(m)
}
func PgFinfoLtqRegex_pg_finfo_ltq_regex(m *base.Module) int32 {
	return F_pg_finfo_ltq_regex(m)
}
func PgFinfoLtqRregex_pg_finfo_ltq_rregex(m *base.Module) int32 {
	return F_pg_finfo_ltq_rregex(m)
}
func PgFinfoLtree2text(m *base.Module) int32 {
	return F_pg_finfo_ltree2text(m)
}
func PgFinfoLtreeAddltree(m *base.Module) int32 {
	return F_pg_finfo_ltree_addltree(m)
}
func PgFinfoLtreeAddtext(m *base.Module) int32 {
	return F_pg_finfo_ltree_addtext(m)
}
func PgFinfoLtreeCmp(m *base.Module) int32 {
	return F_pg_finfo_ltree_cmp(m)
}
func PgFinfoLtreeCompress_pg_finfo_ltree_compress(m *base.Module) int32 {
	return F_pg_finfo_ltree_compress(m)
}
func PgFinfoLtreeConsistent_pg_finfo_ltree_consistent(m *base.Module) int32 {
	return F_pg_finfo_ltree_consistent(m)
}
func PgFinfoLtreeDecompress(m *base.Module) int32 {
	return F_pg_finfo_ltree_decompress(m)
}
func PgFinfoLtreeEq(m *base.Module) int32 {
	return F_pg_finfo_ltree_eq(m)
}
func PgFinfoLtreeGe(m *base.Module) int32 {
	return F_pg_finfo_ltree_ge(m)
}
func PgFinfoLtreeGistIn(m *base.Module) int32 {
	return F_pg_finfo_ltree_gist_in(m)
}
func PgFinfoLtreeGistOptions_pg_finfo_ltree_gist_options(m *base.Module) int32 {
	return F_pg_finfo_ltree_gist_options(m)
}
func PgFinfoLtreeGistOut(m *base.Module) int32 {
	return F_pg_finfo_ltree_gist_out(m)
}
func PgFinfoLtreeGt(m *base.Module) int32 {
	return F_pg_finfo_ltree_gt(m)
}
func PgFinfoLtreeIn(m *base.Module) int32 {
	return F_pg_finfo_ltree_in(m)
}
func PgFinfoLtreeIndex(m *base.Module) int32 {
	return F_pg_finfo_ltree_index(m)
}
func PgFinfoLtreeIsparent_pg_finfo_ltree_isparent(m *base.Module) int32 {
	return F_pg_finfo_ltree_isparent(m)
}
func PgFinfoLtreeLe(m *base.Module) int32 {
	return F_pg_finfo_ltree_le(m)
}
func PgFinfoLtreeLt(m *base.Module) int32 {
	return F_pg_finfo_ltree_lt(m)
}
func PgFinfoLtreeNe(m *base.Module) int32 {
	return F_pg_finfo_ltree_ne(m)
}
func PgFinfoLtreeOut(m *base.Module) int32 {
	return F_pg_finfo_ltree_out(m)
}
func PgFinfoLtreePenalty_pg_finfo_ltree_penalty(m *base.Module) int32 {
	return F_pg_finfo_ltree_penalty(m)
}
func PgFinfoLtreePicksplit_pg_finfo_ltree_picksplit(m *base.Module) int32 {
	return F_pg_finfo_ltree_picksplit(m)
}
func PgFinfoLtreeRecv(m *base.Module) int32 {
	return F_pg_finfo_ltree_recv(m)
}
func PgFinfoLtreeRisparent_pg_finfo_ltree_risparent(m *base.Module) int32 {
	return F_pg_finfo_ltree_risparent(m)
}
func PgFinfoLtreeSame_pg_finfo_ltree_same(m *base.Module) int32 {
	return F_pg_finfo_ltree_same(m)
}
func PgFinfoLtreeSend(m *base.Module) int32 {
	return F_pg_finfo_ltree_send(m)
}
func PgFinfoLtreeTextadd(m *base.Module) int32 {
	return F_pg_finfo_ltree_textadd(m)
}
func PgFinfoLtreeUnion_pg_finfo_ltree_union(m *base.Module) int32 {
	return F_pg_finfo_ltree_union(m)
}
func PgFinfoLtreeparentsel(m *base.Module) int32 {
	return F_pg_finfo_ltreeparentsel(m)
}
func PgFinfoLtxtqExec_pg_finfo_ltxtq_exec(m *base.Module) int32 {
	return F_pg_finfo_ltxtq_exec(m)
}
func PgFinfoLtxtqIn(m *base.Module) int32 {
	return F_pg_finfo_ltxtq_in(m)
}
func PgFinfoLtxtqOut(m *base.Module) int32 {
	return F_pg_finfo_ltxtq_out(m)
}
func PgFinfoLtxtqRecv(m *base.Module) int32 {
	return F_pg_finfo_ltxtq_recv(m)
}
func PgFinfoLtxtqRexec_pg_finfo_ltxtq_rexec(m *base.Module) int32 {
	return F_pg_finfo_ltxtq_rexec(m)
}
func PgFinfoLtxtqSend(m *base.Module) int32 {
	return F_pg_finfo_ltxtq_send(m)
}
func PgFinfoNlevel(m *base.Module) int32 {
	return F_pg_finfo_nlevel(m)
}
func PgFinfoSubltree(m *base.Module) int32 {
	return F_pg_finfo_subltree(m)
}
func PgFinfoSubpath(m *base.Module) int32 {
	return F_pg_finfo_subpath(m)
}
func PgFinfoText2ltree(m *base.Module) int32 {
	return F_pg_finfo_text2ltree(m)
}
func Subltree(m *base.Module, l0 int32) int32 {
	return F_subltree(m, l0)
}
func Subpath(m *base.Module, l0 int32) int32 {
	return F_subpath(m, l0)
}
func Text2ltree(m *base.Module, l0 int32) int32 {
	return F_text2ltree(m, l0)
}
func PgMagicFuncBtreeGist(m *base.Module) int32 {
	return F_Pg_magic_func_btree_gist(m)
}
func CashDist(m *base.Module, l0 int32) int32 {
	return F_cash_dist(m, l0)
}
func DateDist(m *base.Module, l0 int32) int32 {
	return F_date_dist(m, l0)
}
func Float4Dist(m *base.Module, l0 int32) int32 {
	return F_float4_dist(m, l0)
}
func Float8Dist(m *base.Module, l0 int32) int32 {
	return F_float8_dist(m, l0)
}
func GbtBitCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_bit_compress(m, l0)
}
func GbtBitConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_bit_consistent(m, l0)
}
func GbtBitPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_bit_penalty(m, l0)
}
func GbtBitPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_bit_picksplit(m, l0)
}
func GbtBitSame(m *base.Module, l0 int32) int32 {
	return F_gbt_bit_same(m, l0)
}
func GbtBitSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_bit_sortsupport(m, l0)
}
func GbtBitUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_bit_union(m, l0)
}
func GbtBoolCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_bool_compress(m, l0)
}
func GbtBoolConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_bool_consistent(m, l0)
}
func GbtBoolFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_bool_fetch(m, l0)
}
func GbtBoolPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_bool_penalty(m, l0)
}
func GbtBoolPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_bool_picksplit(m, l0)
}
func GbtBoolSame(m *base.Module, l0 int32) int32 {
	return F_gbt_bool_same(m, l0)
}
func GbtBoolSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_bool_sortsupport(m, l0)
}
func GbtBoolUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_bool_union(m, l0)
}
func GbtBpcharCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_text_compress(m, l0)
}
func GbtBpcharConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_bpchar_consistent(m, l0)
}
func GbtBpcharSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_bpchar_sortsupport(m, l0)
}
func GbtByteaCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_bytea_compress(m, l0)
}
func GbtByteaConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_bytea_consistent(m, l0)
}
func GbtByteaPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_bytea_penalty(m, l0)
}
func GbtByteaPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_bytea_picksplit(m, l0)
}
func GbtByteaSame(m *base.Module, l0 int32) int32 {
	return F_gbt_bytea_same(m, l0)
}
func GbtByteaSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_bytea_sortsupport(m, l0)
}
func GbtByteaUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_bytea_union(m, l0)
}
func GbtCashCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_compress(m, l0)
}
func GbtCashConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_consistent(m, l0)
}
func GbtCashDistance(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_distance(m, l0)
}
func GbtCashFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_fetch(m, l0)
}
func GbtCashPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_penalty(m, l0)
}
func GbtCashPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_picksplit(m, l0)
}
func GbtCashSame(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_same(m, l0)
}
func GbtCashSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_sortsupport(m, l0)
}
func GbtCashUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_union(m, l0)
}
func GbtDateCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_date_compress(m, l0)
}
func GbtDateConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_date_consistent(m, l0)
}
func GbtDateDistance(m *base.Module, l0 int32) int32 {
	return F_gbt_date_distance(m, l0)
}
func GbtDateFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_date_fetch(m, l0)
}
func GbtDatePenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_date_penalty(m, l0)
}
func GbtDatePicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_date_picksplit(m, l0)
}
func GbtDateSame(m *base.Module, l0 int32) int32 {
	return F_gbt_date_same(m, l0)
}
func GbtDateSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_date_sortsupport(m, l0)
}
func GbtDateUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_date_union(m, l0)
}
func GbtDecompress(m *base.Module, l0 int32) int32 {
	return F_float4up(m, l0)
}
func GbtEnumCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_enum_compress(m, l0)
}
func GbtEnumConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_enum_consistent(m, l0)
}
func GbtEnumFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_enum_fetch(m, l0)
}
func GbtEnumPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_enum_penalty(m, l0)
}
func GbtEnumPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_enum_picksplit(m, l0)
}
func GbtEnumSame(m *base.Module, l0 int32) int32 {
	return F_gbt_enum_same(m, l0)
}
func GbtEnumSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_enum_sortsupport(m, l0)
}
func GbtEnumUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_enum_union(m, l0)
}
func GbtFloat4Compress(m *base.Module, l0 int32) int32 {
	return F_gbt_float4_compress(m, l0)
}
func GbtFloat4Consistent(m *base.Module, l0 int32) int32 {
	return F_gbt_float4_consistent(m, l0)
}
func GbtFloat4Distance(m *base.Module, l0 int32) int32 {
	return F_gbt_float4_distance(m, l0)
}
func GbtFloat4Fetch(m *base.Module, l0 int32) int32 {
	return F_gbt_float4_fetch(m, l0)
}
func GbtFloat4Penalty(m *base.Module, l0 int32) int32 {
	return F_gbt_float4_penalty(m, l0)
}
func GbtFloat4Picksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_float4_picksplit(m, l0)
}
func GbtFloat4Same(m *base.Module, l0 int32) int32 {
	return F_gbt_float4_same(m, l0)
}
func GbtFloat4Sortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_float4_sortsupport(m, l0)
}
func GbtFloat4Union(m *base.Module, l0 int32) int32 {
	return F_gbt_float4_union(m, l0)
}
func GbtFloat8Compress(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_compress(m, l0)
}
func GbtFloat8Consistent(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_consistent(m, l0)
}
func GbtFloat8Distance(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_distance(m, l0)
}
func GbtFloat8Fetch(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_fetch(m, l0)
}
func GbtFloat8Penalty(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_penalty(m, l0)
}
func GbtFloat8Picksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_picksplit(m, l0)
}
func GbtFloat8Same(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_same(m, l0)
}
func GbtFloat8Sortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_sortsupport(m, l0)
}
func GbtFloat8Union(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_union(m, l0)
}
func GbtInetCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_inet_compress(m, l0)
}
func GbtInetConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_inet_consistent(m, l0)
}
func GbtInetPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_float8_penalty(m, l0)
}
func GbtInetPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_inet_picksplit(m, l0)
}
func GbtInetSame(m *base.Module, l0 int32) int32 {
	return F_gbt_inet_same(m, l0)
}
func GbtInetSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_inet_sortsupport(m, l0)
}
func GbtInetUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_inet_union(m, l0)
}
func GbtInt2Compress(m *base.Module, l0 int32) int32 {
	return F_gbt_int2_compress(m, l0)
}
func GbtInt2Consistent(m *base.Module, l0 int32) int32 {
	return F_gbt_int2_consistent(m, l0)
}
func GbtInt2Distance(m *base.Module, l0 int32) int32 {
	return F_gbt_int2_distance(m, l0)
}
func GbtInt2Fetch(m *base.Module, l0 int32) int32 {
	return F_gbt_int2_fetch(m, l0)
}
func GbtInt2Penalty(m *base.Module, l0 int32) int32 {
	return F_gbt_int2_penalty(m, l0)
}
func GbtInt2Picksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_int2_picksplit(m, l0)
}
func GbtInt2Same(m *base.Module, l0 int32) int32 {
	return F_gbt_int2_same(m, l0)
}
func GbtInt2Sortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_int2_sortsupport(m, l0)
}
func GbtInt2Union(m *base.Module, l0 int32) int32 {
	return F_gbt_int2_union(m, l0)
}
func GbtInt4Compress(m *base.Module, l0 int32) int32 {
	return F_gbt_int4_compress(m, l0)
}
func GbtInt4Consistent(m *base.Module, l0 int32) int32 {
	return F_gbt_int4_consistent(m, l0)
}
func GbtInt4Distance(m *base.Module, l0 int32) int32 {
	return F_gbt_int4_distance(m, l0)
}
func GbtInt4Fetch(m *base.Module, l0 int32) int32 {
	return F_gbt_int4_fetch(m, l0)
}
func GbtInt4Penalty(m *base.Module, l0 int32) int32 {
	return F_gbt_int4_penalty(m, l0)
}
func GbtInt4Picksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_int4_picksplit(m, l0)
}
func GbtInt4Same(m *base.Module, l0 int32) int32 {
	return F_gbt_int4_same(m, l0)
}
func GbtInt4Sortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_int4_sortsupport(m, l0)
}
func GbtInt4Union(m *base.Module, l0 int32) int32 {
	return F_gbt_int4_union(m, l0)
}
func GbtInt8Compress(m *base.Module, l0 int32) int32 {
	return F_gbt_int8_compress(m, l0)
}
func GbtInt8Consistent(m *base.Module, l0 int32) int32 {
	return F_gbt_int8_consistent(m, l0)
}
func GbtInt8Distance(m *base.Module, l0 int32) int32 {
	return F_gbt_int8_distance(m, l0)
}
func GbtInt8Fetch(m *base.Module, l0 int32) int32 {
	return F_gbt_int8_fetch(m, l0)
}
func GbtInt8Penalty(m *base.Module, l0 int32) int32 {
	return F_gbt_cash_penalty(m, l0)
}
func GbtInt8Picksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_int8_picksplit(m, l0)
}
func GbtInt8Same(m *base.Module, l0 int32) int32 {
	return F_gbt_int8_same(m, l0)
}
func GbtInt8Sortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_int8_sortsupport(m, l0)
}
func GbtInt8Union(m *base.Module, l0 int32) int32 {
	return F_gbt_int8_union(m, l0)
}
func GbtIntvCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_intv_compress(m, l0)
}
func GbtIntvConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_intv_consistent(m, l0)
}
func GbtIntvDecompress(m *base.Module, l0 int32) int32 {
	return F_float4up(m, l0)
}
func GbtIntvDistance(m *base.Module, l0 int32) int32 {
	return F_gbt_intv_distance(m, l0)
}
func GbtIntvFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_intv_fetch(m, l0)
}
func GbtIntvPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_intv_penalty(m, l0)
}
func GbtIntvPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_intv_picksplit(m, l0)
}
func GbtIntvSame(m *base.Module, l0 int32) int32 {
	return F_gbt_intv_same(m, l0)
}
func GbtIntvSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_intv_sortsupport(m, l0)
}
func GbtIntvUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_intv_union(m, l0)
}
func GbtMacad8Compress(m *base.Module, l0 int32) int32 {
	return F_gbt_macad8_compress(m, l0)
}
func GbtMacad8Consistent(m *base.Module, l0 int32) int32 {
	return F_gbt_macad8_consistent(m, l0)
}
func GbtMacad8Fetch(m *base.Module, l0 int32) int32 {
	return F_gbt_macad8_fetch(m, l0)
}
func GbtMacad8Penalty(m *base.Module, l0 int32) int32 {
	return F_gbt_macad8_penalty(m, l0)
}
func GbtMacad8Picksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_macad8_picksplit(m, l0)
}
func GbtMacad8Same(m *base.Module, l0 int32) int32 {
	return F_gbt_macad8_same(m, l0)
}
func GbtMacad8Sortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_macad8_sortsupport(m, l0)
}
func GbtMacad8Union(m *base.Module, l0 int32) int32 {
	return F_gbt_macad8_union(m, l0)
}
func GbtMacadCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_macad_compress(m, l0)
}
func GbtMacadConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_macad_consistent(m, l0)
}
func GbtMacadFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_macad_fetch(m, l0)
}
func GbtMacadPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_macad_penalty(m, l0)
}
func GbtMacadPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_macad_picksplit(m, l0)
}
func GbtMacadSame(m *base.Module, l0 int32) int32 {
	return F_gbt_macad_same(m, l0)
}
func GbtMacadUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_macad_union(m, l0)
}
func GbtMacaddrSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_macaddr_sortsupport(m, l0)
}
func GbtNumericCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_numeric_compress(m, l0)
}
func GbtNumericConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_numeric_consistent(m, l0)
}
func GbtNumericPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_numeric_penalty(m, l0)
}
func GbtNumericPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_numeric_picksplit(m, l0)
}
func GbtNumericSame(m *base.Module, l0 int32) int32 {
	return F_gbt_numeric_same(m, l0)
}
func GbtNumericSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_numeric_sortsupport(m, l0)
}
func GbtNumericUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_numeric_union(m, l0)
}
func GbtOidCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_oid_compress(m, l0)
}
func GbtOidConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_oid_consistent(m, l0)
}
func GbtOidDistance(m *base.Module, l0 int32) int32 {
	return F_gbt_oid_distance(m, l0)
}
func GbtOidFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_oid_fetch(m, l0)
}
func GbtOidPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_enum_penalty(m, l0)
}
func GbtOidPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_oid_picksplit(m, l0)
}
func GbtOidSame(m *base.Module, l0 int32) int32 {
	return F_gbt_oid_same(m, l0)
}
func GbtOidSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_oid_sortsupport(m, l0)
}
func GbtOidUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_oid_union(m, l0)
}
func GbtTextCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_text_compress(m, l0)
}
func GbtTextConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_text_consistent(m, l0)
}
func GbtTextPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_text_penalty(m, l0)
}
func GbtTextPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_text_picksplit(m, l0)
}
func GbtTextSame(m *base.Module, l0 int32) int32 {
	return F_gbt_text_same(m, l0)
}
func GbtTextSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_text_sortsupport(m, l0)
}
func GbtTextUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_text_union(m, l0)
}
func GbtTimeCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_time_compress(m, l0)
}
func GbtTimeConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_time_consistent(m, l0)
}
func GbtTimeDistance(m *base.Module, l0 int32) int32 {
	return F_gbt_time_distance(m, l0)
}
func GbtTimeFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_time_fetch(m, l0)
}
func GbtTimePenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_time_penalty(m, l0)
}
func GbtTimePicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_time_picksplit(m, l0)
}
func GbtTimeSame(m *base.Module, l0 int32) int32 {
	return F_gbt_time_same(m, l0)
}
func GbtTimeSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_time_sortsupport(m, l0)
}
func GbtTimeUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_time_union(m, l0)
}
func GbtTimetzCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_timetz_compress(m, l0)
}
func GbtTimetzConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_timetz_consistent(m, l0)
}
func GbtTsCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_ts_compress(m, l0)
}
func GbtTsConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_ts_consistent(m, l0)
}
func GbtTsDistance(m *base.Module, l0 int32) int32 {
	return F_gbt_ts_distance(m, l0)
}
func GbtTsFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_ts_fetch(m, l0)
}
func GbtTsPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_ts_penalty(m, l0)
}
func GbtTsPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_ts_picksplit(m, l0)
}
func GbtTsSame(m *base.Module, l0 int32) int32 {
	return F_gbt_ts_same(m, l0)
}
func GbtTsSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_ts_sortsupport(m, l0)
}
func GbtTsUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_ts_union(m, l0)
}
func GbtTstzCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_tstz_compress(m, l0)
}
func GbtTstzConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_tstz_consistent(m, l0)
}
func GbtTstzDistance(m *base.Module, l0 int32) int32 {
	return F_gbt_tstz_distance(m, l0)
}
func GbtUuidCompress(m *base.Module, l0 int32) int32 {
	return F_gbt_uuid_compress(m, l0)
}
func GbtUuidConsistent(m *base.Module, l0 int32) int32 {
	return F_gbt_uuid_consistent(m, l0)
}
func GbtUuidFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_uuid_fetch(m, l0)
}
func GbtUuidPenalty(m *base.Module, l0 int32) int32 {
	return F_gbt_uuid_penalty(m, l0)
}
func GbtUuidPicksplit(m *base.Module, l0 int32) int32 {
	return F_gbt_uuid_picksplit(m, l0)
}
func GbtUuidSame(m *base.Module, l0 int32) int32 {
	return F_gbt_uuid_same(m, l0)
}
func GbtUuidSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_uuid_sortsupport(m, l0)
}
func GbtUuidUnion(m *base.Module, l0 int32) int32 {
	return F_gbt_uuid_union(m, l0)
}
func GbtVarDecompress(m *base.Module, l0 int32) int32 {
	return F_gtsvector_decompress(m, l0)
}
func GbtVarFetch(m *base.Module, l0 int32) int32 {
	return F_gbt_var_fetch(m, l0)
}
func GbtVarbitSortsupport(m *base.Module, l0 int32) int32 {
	return F_gbt_bit_sortsupport(m, l0)
}
func GbtreekeyIn(m *base.Module, l0 int32) int32 {
	return F_gbtreekey_in(m, l0)
}
func GbtreekeyOut(m *base.Module, l0 int32) int32 {
	return F_gbtreekey_out(m, l0)
}
func GistTranslateCmptypeBtree(m *base.Module, l0 int32) int32 {
	return F_gist_translate_cmptype_btree(m, l0)
}
func Int2Dist(m *base.Module, l0 int32) int32 {
	return F_int2_dist(m, l0)
}
func Int4Dist(m *base.Module, l0 int32) int32 {
	return F_int4_dist(m, l0)
}
func Int8Dist(m *base.Module, l0 int32) int32 {
	return F_int8_dist(m, l0)
}
func IntervalDist(m *base.Module, l0 int32) int32 {
	return F_interval_dist(m, l0)
}
func OidDist(m *base.Module, l0 int32) int32 {
	return F_oid_dist(m, l0)
}
func PgFinfoCashDist(m *base.Module) int32 {
	return F_pg_finfo_cash_dist(m)
}
func PgFinfoDateDist(m *base.Module) int32 {
	return F_pg_finfo_date_dist(m)
}
func PgFinfoFloat4Dist(m *base.Module) int32 {
	return F_pg_finfo_float4_dist(m)
}
func PgFinfoFloat8Dist(m *base.Module) int32 {
	return F_pg_finfo_float8_dist(m)
}
func PgFinfoGbtBitCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_bit_compress(m)
}
func PgFinfoGbtBitConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_bit_consistent(m)
}
func PgFinfoGbtBitPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_bit_penalty(m)
}
func PgFinfoGbtBitPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_bit_picksplit(m)
}
func PgFinfoGbtBitSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_bit_same(m)
}
func PgFinfoGbtBitSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_bit_sortsupport(m)
}
func PgFinfoGbtBitUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_bit_union(m)
}
func PgFinfoGbtBoolCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_bool_compress(m)
}
func PgFinfoGbtBoolConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_bool_consistent(m)
}
func PgFinfoGbtBoolFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_bool_fetch(m)
}
func PgFinfoGbtBoolPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_bool_penalty(m)
}
func PgFinfoGbtBoolPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_bool_picksplit(m)
}
func PgFinfoGbtBoolSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_bool_same(m)
}
func PgFinfoGbtBoolSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_bool_sortsupport(m)
}
func PgFinfoGbtBoolUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_bool_union(m)
}
func PgFinfoGbtBpcharCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_bpchar_compress(m)
}
func PgFinfoGbtBpcharConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_bpchar_consistent(m)
}
func PgFinfoGbtBpcharSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_bpchar_sortsupport(m)
}
func PgFinfoGbtByteaCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_bytea_compress(m)
}
func PgFinfoGbtByteaConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_bytea_consistent(m)
}
func PgFinfoGbtByteaPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_bytea_penalty(m)
}
func PgFinfoGbtByteaPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_bytea_picksplit(m)
}
func PgFinfoGbtByteaSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_bytea_same(m)
}
func PgFinfoGbtByteaSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_bytea_sortsupport(m)
}
func PgFinfoGbtByteaUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_bytea_union(m)
}
func PgFinfoGbtCashCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_cash_compress(m)
}
func PgFinfoGbtCashConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_cash_consistent(m)
}
func PgFinfoGbtCashDistance(m *base.Module) int32 {
	return F_pg_finfo_gbt_cash_distance(m)
}
func PgFinfoGbtCashFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_cash_fetch(m)
}
func PgFinfoGbtCashPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_cash_penalty(m)
}
func PgFinfoGbtCashPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_cash_picksplit(m)
}
func PgFinfoGbtCashSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_cash_same(m)
}
func PgFinfoGbtCashSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_cash_sortsupport(m)
}
func PgFinfoGbtCashUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_cash_union(m)
}
func PgFinfoGbtDateCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_date_compress(m)
}
func PgFinfoGbtDateConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_date_consistent(m)
}
func PgFinfoGbtDateDistance(m *base.Module) int32 {
	return F_pg_finfo_gbt_date_distance(m)
}
func PgFinfoGbtDateFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_date_fetch(m)
}
func PgFinfoGbtDatePenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_date_penalty(m)
}
func PgFinfoGbtDatePicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_date_picksplit(m)
}
func PgFinfoGbtDateSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_date_same(m)
}
func PgFinfoGbtDateSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_date_sortsupport(m)
}
func PgFinfoGbtDateUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_date_union(m)
}
func PgFinfoGbtDecompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_decompress(m)
}
func PgFinfoGbtEnumCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_enum_compress(m)
}
func PgFinfoGbtEnumConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_enum_consistent(m)
}
func PgFinfoGbtEnumFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_enum_fetch(m)
}
func PgFinfoGbtEnumPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_enum_penalty(m)
}
func PgFinfoGbtEnumPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_enum_picksplit(m)
}
func PgFinfoGbtEnumSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_enum_same(m)
}
func PgFinfoGbtEnumSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_enum_sortsupport(m)
}
func PgFinfoGbtEnumUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_enum_union(m)
}
func PgFinfoGbtFloat4Compress(m *base.Module) int32 {
	return F_pg_finfo_gbt_float4_compress(m)
}
func PgFinfoGbtFloat4Consistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_float4_consistent(m)
}
func PgFinfoGbtFloat4Distance(m *base.Module) int32 {
	return F_pg_finfo_gbt_float4_distance(m)
}
func PgFinfoGbtFloat4Fetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_float4_fetch(m)
}
func PgFinfoGbtFloat4Penalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_float4_penalty(m)
}
func PgFinfoGbtFloat4Picksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_float4_picksplit(m)
}
func PgFinfoGbtFloat4Same(m *base.Module) int32 {
	return F_pg_finfo_gbt_float4_same(m)
}
func PgFinfoGbtFloat4Sortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_float4_sortsupport(m)
}
func PgFinfoGbtFloat4Union(m *base.Module) int32 {
	return F_pg_finfo_gbt_float4_union(m)
}
func PgFinfoGbtFloat8Compress(m *base.Module) int32 {
	return F_pg_finfo_gbt_float8_compress(m)
}
func PgFinfoGbtFloat8Consistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_float8_consistent(m)
}
func PgFinfoGbtFloat8Distance(m *base.Module) int32 {
	return F_pg_finfo_gbt_float8_distance(m)
}
func PgFinfoGbtFloat8Fetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_float8_fetch(m)
}
func PgFinfoGbtFloat8Penalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_float8_penalty(m)
}
func PgFinfoGbtFloat8Picksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_float8_picksplit(m)
}
func PgFinfoGbtFloat8Same(m *base.Module) int32 {
	return F_pg_finfo_gbt_float8_same(m)
}
func PgFinfoGbtFloat8Sortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_float8_sortsupport(m)
}
func PgFinfoGbtFloat8Union(m *base.Module) int32 {
	return F_pg_finfo_gbt_float8_union(m)
}
func PgFinfoGbtInetCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_inet_compress(m)
}
func PgFinfoGbtInetConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_inet_consistent(m)
}
func PgFinfoGbtInetPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_inet_penalty(m)
}
func PgFinfoGbtInetPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_inet_picksplit(m)
}
func PgFinfoGbtInetSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_inet_same(m)
}
func PgFinfoGbtInetSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_inet_sortsupport(m)
}
func PgFinfoGbtInetUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_inet_union(m)
}
func PgFinfoGbtInt2Compress(m *base.Module) int32 {
	return F_pg_finfo_gbt_int2_compress(m)
}
func PgFinfoGbtInt2Consistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_int2_consistent(m)
}
func PgFinfoGbtInt2Distance(m *base.Module) int32 {
	return F_pg_finfo_gbt_int2_distance(m)
}
func PgFinfoGbtInt2Fetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_int2_fetch(m)
}
func PgFinfoGbtInt2Penalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_int2_penalty(m)
}
func PgFinfoGbtInt2Picksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_int2_picksplit(m)
}
func PgFinfoGbtInt2Same(m *base.Module) int32 {
	return F_pg_finfo_gbt_int2_same(m)
}
func PgFinfoGbtInt2Sortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_int2_sortsupport(m)
}
func PgFinfoGbtInt2Union(m *base.Module) int32 {
	return F_pg_finfo_gbt_int2_union(m)
}
func PgFinfoGbtInt4Compress(m *base.Module) int32 {
	return F_pg_finfo_gbt_int4_compress(m)
}
func PgFinfoGbtInt4Consistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_int4_consistent(m)
}
func PgFinfoGbtInt4Distance(m *base.Module) int32 {
	return F_pg_finfo_gbt_int4_distance(m)
}
func PgFinfoGbtInt4Fetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_int4_fetch(m)
}
func PgFinfoGbtInt4Penalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_int4_penalty(m)
}
func PgFinfoGbtInt4Picksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_int4_picksplit(m)
}
func PgFinfoGbtInt4Same(m *base.Module) int32 {
	return F_pg_finfo_gbt_int4_same(m)
}
func PgFinfoGbtInt4Sortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_int4_sortsupport(m)
}
func PgFinfoGbtInt4Union(m *base.Module) int32 {
	return F_pg_finfo_gbt_int4_union(m)
}
func PgFinfoGbtInt8Compress(m *base.Module) int32 {
	return F_pg_finfo_gbt_int8_compress(m)
}
func PgFinfoGbtInt8Consistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_int8_consistent(m)
}
func PgFinfoGbtInt8Distance(m *base.Module) int32 {
	return F_pg_finfo_gbt_int8_distance(m)
}
func PgFinfoGbtInt8Fetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_int8_fetch(m)
}
func PgFinfoGbtInt8Penalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_int8_penalty(m)
}
func PgFinfoGbtInt8Picksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_int8_picksplit(m)
}
func PgFinfoGbtInt8Same(m *base.Module) int32 {
	return F_pg_finfo_gbt_int8_same(m)
}
func PgFinfoGbtInt8Sortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_int8_sortsupport(m)
}
func PgFinfoGbtInt8Union(m *base.Module) int32 {
	return F_pg_finfo_gbt_int8_union(m)
}
func PgFinfoGbtIntvCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_compress(m)
}
func PgFinfoGbtIntvConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_consistent(m)
}
func PgFinfoGbtIntvDecompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_decompress(m)
}
func PgFinfoGbtIntvDistance(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_distance(m)
}
func PgFinfoGbtIntvFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_fetch(m)
}
func PgFinfoGbtIntvPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_penalty(m)
}
func PgFinfoGbtIntvPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_picksplit(m)
}
func PgFinfoGbtIntvSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_same(m)
}
func PgFinfoGbtIntvSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_sortsupport(m)
}
func PgFinfoGbtIntvUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_intv_union(m)
}
func PgFinfoGbtMacad8Compress(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad8_compress(m)
}
func PgFinfoGbtMacad8Consistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad8_consistent(m)
}
func PgFinfoGbtMacad8Fetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad8_fetch(m)
}
func PgFinfoGbtMacad8Penalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad8_penalty(m)
}
func PgFinfoGbtMacad8Picksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad8_picksplit(m)
}
func PgFinfoGbtMacad8Same(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad8_same(m)
}
func PgFinfoGbtMacad8Sortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad8_sortsupport(m)
}
func PgFinfoGbtMacad8Union(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad8_union(m)
}
func PgFinfoGbtMacadCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad_compress(m)
}
func PgFinfoGbtMacadConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad_consistent(m)
}
func PgFinfoGbtMacadFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad_fetch(m)
}
func PgFinfoGbtMacadPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad_penalty(m)
}
func PgFinfoGbtMacadPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad_picksplit(m)
}
func PgFinfoGbtMacadSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad_same(m)
}
func PgFinfoGbtMacadUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_macad_union(m)
}
func PgFinfoGbtMacaddrSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_macaddr_sortsupport(m)
}
func PgFinfoGbtNumericCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_numeric_compress(m)
}
func PgFinfoGbtNumericConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_numeric_consistent(m)
}
func PgFinfoGbtNumericPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_numeric_penalty(m)
}
func PgFinfoGbtNumericPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_numeric_picksplit(m)
}
func PgFinfoGbtNumericSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_numeric_same(m)
}
func PgFinfoGbtNumericSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_numeric_sortsupport(m)
}
func PgFinfoGbtNumericUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_numeric_union(m)
}
func PgFinfoGbtOidCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_oid_compress(m)
}
func PgFinfoGbtOidConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_oid_consistent(m)
}
func PgFinfoGbtOidDistance(m *base.Module) int32 {
	return F_pg_finfo_gbt_oid_distance(m)
}
func PgFinfoGbtOidFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_oid_fetch(m)
}
func PgFinfoGbtOidPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_oid_penalty(m)
}
func PgFinfoGbtOidPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_oid_picksplit(m)
}
func PgFinfoGbtOidSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_oid_same(m)
}
func PgFinfoGbtOidSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_oid_sortsupport(m)
}
func PgFinfoGbtOidUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_oid_union(m)
}
func PgFinfoGbtTextCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_text_compress(m)
}
func PgFinfoGbtTextConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_text_consistent(m)
}
func PgFinfoGbtTextPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_text_penalty(m)
}
func PgFinfoGbtTextPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_text_picksplit(m)
}
func PgFinfoGbtTextSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_text_same(m)
}
func PgFinfoGbtTextSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_text_sortsupport(m)
}
func PgFinfoGbtTextUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_text_union(m)
}
func PgFinfoGbtTimeCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_time_compress(m)
}
func PgFinfoGbtTimeConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_time_consistent(m)
}
func PgFinfoGbtTimeDistance(m *base.Module) int32 {
	return F_pg_finfo_gbt_time_distance(m)
}
func PgFinfoGbtTimeFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_time_fetch(m)
}
func PgFinfoGbtTimePenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_time_penalty(m)
}
func PgFinfoGbtTimePicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_time_picksplit(m)
}
func PgFinfoGbtTimeSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_time_same(m)
}
func PgFinfoGbtTimeSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_time_sortsupport(m)
}
func PgFinfoGbtTimeUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_time_union(m)
}
func PgFinfoGbtTimetzCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_timetz_compress(m)
}
func PgFinfoGbtTimetzConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_timetz_consistent(m)
}
func PgFinfoGbtTimetzSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_timetz_sortsupport(m)
}
func PgFinfoGbtTsCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_ts_compress(m)
}
func PgFinfoGbtTsConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_ts_consistent(m)
}
func PgFinfoGbtTsDistance(m *base.Module) int32 {
	return F_pg_finfo_gbt_ts_distance(m)
}
func PgFinfoGbtTsFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_ts_fetch(m)
}
func PgFinfoGbtTsPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_ts_penalty(m)
}
func PgFinfoGbtTsPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_ts_picksplit(m)
}
func PgFinfoGbtTsSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_ts_same(m)
}
func PgFinfoGbtTsSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_ts_sortsupport(m)
}
func PgFinfoGbtTsUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_ts_union(m)
}
func PgFinfoGbtTstzCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_tstz_compress(m)
}
func PgFinfoGbtTstzConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_tstz_consistent(m)
}
func PgFinfoGbtTstzDistance(m *base.Module) int32 {
	return F_pg_finfo_gbt_tstz_distance(m)
}
func PgFinfoGbtUuidCompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_uuid_compress(m)
}
func PgFinfoGbtUuidConsistent(m *base.Module) int32 {
	return F_pg_finfo_gbt_uuid_consistent(m)
}
func PgFinfoGbtUuidFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_uuid_fetch(m)
}
func PgFinfoGbtUuidPenalty(m *base.Module) int32 {
	return F_pg_finfo_gbt_uuid_penalty(m)
}
func PgFinfoGbtUuidPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gbt_uuid_picksplit(m)
}
func PgFinfoGbtUuidSame(m *base.Module) int32 {
	return F_pg_finfo_gbt_uuid_same(m)
}
func PgFinfoGbtUuidSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_uuid_sortsupport(m)
}
func PgFinfoGbtUuidUnion(m *base.Module) int32 {
	return F_pg_finfo_gbt_uuid_union(m)
}
func PgFinfoGbtVarDecompress(m *base.Module) int32 {
	return F_pg_finfo_gbt_var_decompress(m)
}
func PgFinfoGbtVarFetch(m *base.Module) int32 {
	return F_pg_finfo_gbt_var_fetch(m)
}
func PgFinfoGbtVarbitSortsupport(m *base.Module) int32 {
	return F_pg_finfo_gbt_varbit_sortsupport(m)
}
func PgFinfoGbtreekeyIn(m *base.Module) int32 {
	return F_pg_finfo_gbtreekey_in(m)
}
func PgFinfoGbtreekeyOut(m *base.Module) int32 {
	return F_pg_finfo_gbtreekey_out(m)
}
func PgFinfoGistTranslateCmptypeBtree(m *base.Module) int32 {
	return F_pg_finfo_gist_translate_cmptype_btree(m)
}
func PgFinfoInt2Dist(m *base.Module) int32 {
	return F_pg_finfo_int2_dist(m)
}
func PgFinfoInt4Dist(m *base.Module) int32 {
	return F_pg_finfo_int4_dist(m)
}
func PgFinfoInt8Dist(m *base.Module) int32 {
	return F_pg_finfo_int8_dist(m)
}
func PgFinfoIntervalDist(m *base.Module) int32 {
	return F_pg_finfo_interval_dist(m)
}
func PgFinfoOidDist(m *base.Module) int32 {
	return F_pg_finfo_oid_dist(m)
}
func PgFinfoTimeDist(m *base.Module) int32 {
	return F_pg_finfo_time_dist(m)
}
func PgFinfoTsDist(m *base.Module) int32 {
	return F_pg_finfo_ts_dist(m)
}
func PgFinfoTstzDist(m *base.Module) int32 {
	return F_pg_finfo_tstz_dist(m)
}
func TimeDist(m *base.Module, l0 int32) int32 {
	return F_time_dist(m, l0)
}
func TsDist(m *base.Module, l0 int32) int32 {
	return F_ts_dist(m, l0)
}
func TstzDist(m *base.Module, l0 int32) int32 {
	return F_ts_dist(m, l0)
}
func PgMagicFuncBtreeGin(m *base.Module) int32 {
	return F_Pg_magic_func_btree_gin(m)
}
func GinBtreeConsistent(m *base.Module, l0 int32) int32 {
	return F_gin_btree_consistent(m, l0)
}
func GinComparePrefixAnyenum(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixBit(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixBool(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixBpchar(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixBytea(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixChar(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixCidr(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixDate(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixFloat4(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixFloat8(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixInet(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixInt2(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixInt4(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixInt8(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixInterval(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixMacaddr(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixMacaddr8(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixMoney(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixName(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixNumeric(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixOid(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixText(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixTime(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixTimestamp(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixTimestamptz(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixTimetz(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixUuid(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinComparePrefixVarbit(m *base.Module, l0 int32) int32 {
	return F_gin_compare_prefix_int2(m, l0)
}
func GinEnumCmp(m *base.Module, l0 int32) int32 {
	return F_gin_enum_cmp(m, l0)
}
func GinExtractQueryAnyenum(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_anyenum(m, l0)
}
func GinExtractQueryBit(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_bit(m, l0)
}
func GinExtractQueryBool(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_bool(m, l0)
}
func GinExtractQueryBpchar(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_bpchar(m, l0)
}
func GinExtractQueryBytea(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_bytea(m, l0)
}
func GinExtractQueryChar(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_char(m, l0)
}
func GinExtractQueryCidr(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_inet(m, l0)
}
func GinExtractQueryDate(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_date(m, l0)
}
func GinExtractQueryFloat4(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_float4(m, l0)
}
func GinExtractQueryFloat8(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_float8(m, l0)
}
func GinExtractQueryInet(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_inet(m, l0)
}
func GinExtractQueryInt2(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_int2(m, l0)
}
func GinExtractQueryInt4(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_int4(m, l0)
}
func GinExtractQueryInt8(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_int8(m, l0)
}
func GinExtractQueryInterval(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_interval(m, l0)
}
func GinExtractQueryMacaddr(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_macaddr(m, l0)
}
func GinExtractQueryMacaddr8(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_macaddr8(m, l0)
}
func GinExtractQueryMoney(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_money(m, l0)
}
func GinExtractQueryName(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_name(m, l0)
}
func GinExtractQueryNumeric(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_numeric(m, l0)
}
func GinExtractQueryOid(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_oid(m, l0)
}
func GinExtractQueryText(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_text(m, l0)
}
func GinExtractQueryTime(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_time(m, l0)
}
func GinExtractQueryTimestamp(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_timestamp(m, l0)
}
func GinExtractQueryTimestamptz(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_timestamp(m, l0)
}
func GinExtractQueryTimetz(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_timetz(m, l0)
}
func GinExtractQueryUuid(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_uuid(m, l0)
}
func GinExtractQueryVarbit(m *base.Module, l0 int32) int32 {
	return F_gin_extract_query_varbit(m, l0)
}
func GinExtractValueAnyenum(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueBit(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_inet(m, l0)
}
func GinExtractValueBool(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueBpchar(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_inet(m, l0)
}
func GinExtractValueBytea(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_inet(m, l0)
}
func GinExtractValueChar(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueCidr(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_inet(m, l0)
}
func GinExtractValueDate(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueFloat4(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueFloat8(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueInet(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_inet(m, l0)
}
func GinExtractValueInt2(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueInt4(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueInt8(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueInterval(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueMacaddr(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueMacaddr8(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueMoney(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueName(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueNumeric(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_inet(m, l0)
}
func GinExtractValueOid(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueText(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_inet(m, l0)
}
func GinExtractValueTime(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueTimestamp(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueTimestamptz(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueTimetz(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueUuid(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_int2(m, l0)
}
func GinExtractValueVarbit(m *base.Module, l0 int32) int32 {
	return F_gin_extract_value_inet(m, l0)
}
func GinNumericCmp(m *base.Module, l0 int32) int32 {
	return F_gin_numeric_cmp(m, l0)
}
func PgFinfoGinBtreeConsistent(m *base.Module) int32 {
	return F_pg_finfo_gin_btree_consistent(m)
}
func PgFinfoGinComparePrefixAnyenum(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_anyenum(m)
}
func PgFinfoGinComparePrefixBit(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_bit(m)
}
func PgFinfoGinComparePrefixBool(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_bool(m)
}
func PgFinfoGinComparePrefixBpchar(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_bpchar(m)
}
func PgFinfoGinComparePrefixBytea(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_bytea(m)
}
func PgFinfoGinComparePrefixChar(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_char(m)
}
func PgFinfoGinComparePrefixCidr(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_cidr(m)
}
func PgFinfoGinComparePrefixDate(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_date(m)
}
func PgFinfoGinComparePrefixFloat4(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_float4(m)
}
func PgFinfoGinComparePrefixFloat8(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_float8(m)
}
func PgFinfoGinComparePrefixInet(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_inet(m)
}
func PgFinfoGinComparePrefixInt2(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_int2(m)
}
func PgFinfoGinComparePrefixInt4(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_int4(m)
}
func PgFinfoGinComparePrefixInt8(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_int8(m)
}
func PgFinfoGinComparePrefixInterval(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_interval(m)
}
func PgFinfoGinComparePrefixMacaddr(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_macaddr(m)
}
func PgFinfoGinComparePrefixMacaddr8(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_macaddr8(m)
}
func PgFinfoGinComparePrefixMoney(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_money(m)
}
func PgFinfoGinComparePrefixName(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_name(m)
}
func PgFinfoGinComparePrefixNumeric(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_numeric(m)
}
func PgFinfoGinComparePrefixOid(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_oid(m)
}
func PgFinfoGinComparePrefixText(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_text(m)
}
func PgFinfoGinComparePrefixTime(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_time(m)
}
func PgFinfoGinComparePrefixTimestamp(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_timestamp(m)
}
func PgFinfoGinComparePrefixTimestamptz(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_timestamptz(m)
}
func PgFinfoGinComparePrefixTimetz(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_timetz(m)
}
func PgFinfoGinComparePrefixUuid(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_uuid(m)
}
func PgFinfoGinComparePrefixVarbit(m *base.Module) int32 {
	return F_pg_finfo_gin_compare_prefix_varbit(m)
}
func PgFinfoGinEnumCmp(m *base.Module) int32 {
	return F_pg_finfo_gin_enum_cmp(m)
}
func PgFinfoGinExtractQueryAnyenum(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_anyenum(m)
}
func PgFinfoGinExtractQueryBit(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_bit(m)
}
func PgFinfoGinExtractQueryBool(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_bool(m)
}
func PgFinfoGinExtractQueryBpchar(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_bpchar(m)
}
func PgFinfoGinExtractQueryBytea(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_bytea(m)
}
func PgFinfoGinExtractQueryChar(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_char(m)
}
func PgFinfoGinExtractQueryCidr(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_cidr(m)
}
func PgFinfoGinExtractQueryDate(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_date(m)
}
func PgFinfoGinExtractQueryFloat4(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_float4(m)
}
func PgFinfoGinExtractQueryFloat8(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_float8(m)
}
func PgFinfoGinExtractQueryInet(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_inet(m)
}
func PgFinfoGinExtractQueryInt2(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_int2(m)
}
func PgFinfoGinExtractQueryInt4(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_int4(m)
}
func PgFinfoGinExtractQueryInt8(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_int8(m)
}
func PgFinfoGinExtractQueryInterval(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_interval(m)
}
func PgFinfoGinExtractQueryMacaddr(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_macaddr(m)
}
func PgFinfoGinExtractQueryMacaddr8(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_macaddr8(m)
}
func PgFinfoGinExtractQueryMoney(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_money(m)
}
func PgFinfoGinExtractQueryName(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_name(m)
}
func PgFinfoGinExtractQueryNumeric(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_numeric(m)
}
func PgFinfoGinExtractQueryOid(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_oid(m)
}
func PgFinfoGinExtractQueryText(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_text(m)
}
func PgFinfoGinExtractQueryTime(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_time(m)
}
func PgFinfoGinExtractQueryTimestamp(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_timestamp(m)
}
func PgFinfoGinExtractQueryTimestamptz(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_timestamptz(m)
}
func PgFinfoGinExtractQueryTimetz(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_timetz(m)
}
func PgFinfoGinExtractQueryUuid(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_uuid(m)
}
func PgFinfoGinExtractQueryVarbit(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_query_varbit(m)
}
func PgFinfoGinExtractValueAnyenum(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_anyenum(m)
}
func PgFinfoGinExtractValueBit(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_bit(m)
}
func PgFinfoGinExtractValueBool(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_bool(m)
}
func PgFinfoGinExtractValueBpchar(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_bpchar(m)
}
func PgFinfoGinExtractValueBytea(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_bytea(m)
}
func PgFinfoGinExtractValueChar(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_char(m)
}
func PgFinfoGinExtractValueCidr(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_cidr(m)
}
func PgFinfoGinExtractValueDate(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_date(m)
}
func PgFinfoGinExtractValueFloat4(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_float4(m)
}
func PgFinfoGinExtractValueFloat8(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_float8(m)
}
func PgFinfoGinExtractValueInet(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_inet(m)
}
func PgFinfoGinExtractValueInt2(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_int2(m)
}
func PgFinfoGinExtractValueInt4(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_int4(m)
}
func PgFinfoGinExtractValueInt8(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_int8(m)
}
func PgFinfoGinExtractValueInterval(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_interval(m)
}
func PgFinfoGinExtractValueMacaddr(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_macaddr(m)
}
func PgFinfoGinExtractValueMacaddr8(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_macaddr8(m)
}
func PgFinfoGinExtractValueMoney(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_money(m)
}
func PgFinfoGinExtractValueName(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_name(m)
}
func PgFinfoGinExtractValueNumeric(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_numeric(m)
}
func PgFinfoGinExtractValueOid(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_oid(m)
}
func PgFinfoGinExtractValueText(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_text(m)
}
func PgFinfoGinExtractValueTime(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_time(m)
}
func PgFinfoGinExtractValueTimestamp(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_timestamp(m)
}
func PgFinfoGinExtractValueTimestamptz(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_timestamptz(m)
}
func PgFinfoGinExtractValueTimetz(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_timetz(m)
}
func PgFinfoGinExtractValueUuid(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_uuid(m)
}
func PgFinfoGinExtractValueVarbit(m *base.Module) int32 {
	return F_pg_finfo_gin_extract_value_varbit(m)
}
func PgFinfoGinNumericCmp(m *base.Module) int32 {
	return F_pg_finfo_gin_numeric_cmp(m)
}
func PgMagicFuncUnaccent(m *base.Module) int32 {
	return F_Pg_magic_func_unaccent(m)
}
func PgFinfoUnaccentDict(m *base.Module) int32 {
	return F_pg_finfo_unaccent_dict(m)
}
func PgFinfoUnaccentInit(m *base.Module) int32 {
	return F_pg_finfo_unaccent_init(m)
}
func PgFinfoUnaccentLexize(m *base.Module) int32 {
	return F_pg_finfo_unaccent_lexize(m)
}
func UnaccentDict(m *base.Module, l0 int32) int32 {
	return F_unaccent_dict(m, l0)
}
func UnaccentInit(m *base.Module, l0 int32) int32 {
	return F_unaccent_init(m, l0)
}
func UnaccentLexize(m *base.Module, l0 int32) int32 {
	return F_unaccent_lexize(m, l0)
}
func PgMagicFuncTablefunc(m *base.Module) int32 {
	return F_Pg_magic_func_tablefunc(m)
}
func ConnectbyText(m *base.Module, l0 int32) int32 {
	return F_connectby_text(m, l0)
}
func ConnectbyTextSerial(m *base.Module, l0 int32) int32 {
	return F_connectby_text_serial(m, l0)
}
func Crosstab(m *base.Module, l0 int32) int32 {
	return F_crosstab(m, l0)
}
func CrosstabHash(m *base.Module, l0 int32) int32 {
	return F_crosstab_hash(m, l0)
}
func NormalRand(m *base.Module, l0 int32) int32 {
	return F_normal_rand(m, l0)
}
func PgFinfoConnectbyText(m *base.Module) int32 {
	return F_pg_finfo_connectby_text(m)
}
func PgFinfoConnectbyTextSerial(m *base.Module) int32 {
	return F_pg_finfo_connectby_text_serial(m)
}
func PgFinfoCrosstab(m *base.Module) int32 {
	return F_pg_finfo_crosstab(m)
}
func PgFinfoCrosstabHash(m *base.Module) int32 {
	return F_pg_finfo_crosstab_hash(m)
}
func PgFinfoNormalRand(m *base.Module) int32 {
	return F_pg_finfo_normal_rand(m)
}
func PgMagicFuncInt(m *base.Module) int32 {
	return F_Pg_magic_func__int(m)
}
func IntContained(m *base.Module, l0 int32) int32 {
	return F__int_contained(m, l0)
}
func IntContainedJoinsel(m *base.Module, l0 int32) int32 {
	return F__int_contained_joinsel(m, l0)
}
func IntContainedSel(m *base.Module, l0 int32) int32 {
	return F__int_contained_sel(m, l0)
}
func IntContains(m *base.Module, l0 int32) int32 {
	return F__int_contains(m, l0)
}
func IntContainsJoinsel(m *base.Module, l0 int32) int32 {
	return F__int_contains_joinsel(m, l0)
}
func IntContainsSel(m *base.Module, l0 int32) int32 {
	return F__int_contains_sel(m, l0)
}
func IntDifferent(m *base.Module, l0 int32) int32 {
	return F__int_different(m, l0)
}
func IntInter(m *base.Module, l0 int32) int32 {
	return F__int_inter(m, l0)
}
func IntMatchsel(m *base.Module, l0 int32) int32 {
	return F__int_matchsel(m, l0)
}
func IntOverlap(m *base.Module, l0 int32) int32 {
	return F__int_overlap(m, l0)
}
func IntOverlapJoinsel(m *base.Module, l0 int32) int32 {
	return F__int_overlap_joinsel(m, l0)
}
func IntOverlapSel(m *base.Module, l0 int32) int32 {
	return F__int_overlap_sel(m, l0)
}
func IntSame(m *base.Module, l0 int32) int32 {
	return F__int_same(m, l0)
}
func IntUnion(m *base.Module, l0 int32) int32 {
	return F__int_union(m, l0)
}
func IntbigIn(m *base.Module, l0 int32) int32 {
	return F__intbig_in(m, l0)
}
func IntbigOut(m *base.Module, l0 int32) int32 {
	return F__intbig_out(m, l0)
}
func Boolop(m *base.Module, l0 int32) int32 {
	return F_boolop(m, l0)
}
func BqarrIn(m *base.Module, l0 int32) int32 {
	return F_bqarr_in(m, l0)
}
func BqarrOut(m *base.Module, l0 int32) int32 {
	return F_bqarr_out(m, l0)
}
func GIntCompress(m *base.Module, l0 int32) int32 {
	return F_g_int_compress(m, l0)
}
func GIntConsistent(m *base.Module, l0 int32) int32 {
	return F_g_int_consistent(m, l0)
}
func GIntDecompress(m *base.Module, l0 int32) int32 {
	return F_g_int_decompress(m, l0)
}
func GIntOptions(m *base.Module, l0 int32) int32 {
	return F_g_int_options(m, l0)
}
func GIntPenalty(m *base.Module, l0 int32) int32 {
	return F_g_int_penalty(m, l0)
}
func GIntPicksplit(m *base.Module, l0 int32) int32 {
	return F_g_int_picksplit(m, l0)
}
func GIntSame(m *base.Module, l0 int32) int32 {
	return F_g_int_same(m, l0)
}
func GIntUnion(m *base.Module, l0 int32) int32 {
	return F_g_int_union(m, l0)
}
func GIntbigCompress(m *base.Module, l0 int32) int32 {
	return F_g_intbig_compress(m, l0)
}
func GIntbigConsistent(m *base.Module, l0 int32) int32 {
	return F_g_intbig_consistent(m, l0)
}
func GIntbigDecompress(m *base.Module, l0 int32) int32 {
	return F_float4up(m, l0)
}
func GIntbigOptions(m *base.Module, l0 int32) int32 {
	return F_g_intbig_options(m, l0)
}
func GIntbigPenalty(m *base.Module, l0 int32) int32 {
	return F_g_intbig_penalty(m, l0)
}
func GIntbigPicksplit(m *base.Module, l0 int32) int32 {
	return F_g_intbig_picksplit(m, l0)
}
func GIntbigSame(m *base.Module, l0 int32) int32 {
	return F_g_intbig_same(m, l0)
}
func GIntbigUnion(m *base.Module, l0 int32) int32 {
	return F_g_intbig_union(m, l0)
}
func Ginint4Consistent(m *base.Module, l0 int32) int32 {
	return F_ginint4_consistent(m, l0)
}
func Ginint4Queryextract(m *base.Module, l0 int32) int32 {
	return F_ginint4_queryextract(m, l0)
}
func Icount(m *base.Module, l0 int32) int32 {
	return F_icount(m, l0)
}
func Idx(m *base.Module, l0 int32) int32 {
	return F_idx(m, l0)
}
func IntarrayDelElem(m *base.Module, l0 int32) int32 {
	return F_intarray_del_elem(m, l0)
}
func IntarrayPushArray(m *base.Module, l0 int32) int32 {
	return F_intarray_push_array(m, l0)
}
func IntarrayPushElem(m *base.Module, l0 int32) int32 {
	return F_intarray_push_elem(m, l0)
}
func Intset(m *base.Module, l0 int32) int32 {
	return F_intset(m, l0)
}
func IntsetSubtract(m *base.Module, l0 int32) int32 {
	return F_intset_subtract(m, l0)
}
func IntsetUnionElem(m *base.Module, l0 int32) int32 {
	return F_intset_union_elem(m, l0)
}
func PgFinfoIntContained(m *base.Module) int32 {
	return F_pg_finfo__int_contained(m)
}
func PgFinfoIntContainedJoinsel(m *base.Module) int32 {
	return F_pg_finfo__int_contained_joinsel(m)
}
func PgFinfoIntContainedSel(m *base.Module) int32 {
	return F_pg_finfo__int_contained_sel(m)
}
func PgFinfoIntContains(m *base.Module) int32 {
	return F_pg_finfo__int_contains(m)
}
func PgFinfoIntContainsJoinsel(m *base.Module) int32 {
	return F_pg_finfo__int_contains_joinsel(m)
}
func PgFinfoIntContainsSel(m *base.Module) int32 {
	return F_pg_finfo__int_contains_sel(m)
}
func PgFinfoIntDifferent(m *base.Module) int32 {
	return F_pg_finfo__int_different(m)
}
func PgFinfoIntInter(m *base.Module) int32 {
	return F_pg_finfo__int_inter(m)
}
func PgFinfoIntMatchsel(m *base.Module) int32 {
	return F_pg_finfo__int_matchsel(m)
}
func PgFinfoIntOverlap(m *base.Module) int32 {
	return F_pg_finfo__int_overlap(m)
}
func PgFinfoIntOverlapJoinsel(m *base.Module) int32 {
	return F_pg_finfo__int_overlap_joinsel(m)
}
func PgFinfoIntOverlapSel(m *base.Module) int32 {
	return F_pg_finfo__int_overlap_sel(m)
}
func PgFinfoIntSame(m *base.Module) int32 {
	return F_pg_finfo__int_same(m)
}
func PgFinfoIntUnion(m *base.Module) int32 {
	return F_pg_finfo__int_union(m)
}
func PgFinfoIntbigIn(m *base.Module) int32 {
	return F_pg_finfo__intbig_in(m)
}
func PgFinfoIntbigOut(m *base.Module) int32 {
	return F_pg_finfo__intbig_out(m)
}
func PgFinfoBoolop(m *base.Module) int32 {
	return F_pg_finfo_boolop(m)
}
func PgFinfoBqarrIn(m *base.Module) int32 {
	return F_pg_finfo_bqarr_in(m)
}
func PgFinfoBqarrOut(m *base.Module) int32 {
	return F_pg_finfo_bqarr_out(m)
}
func PgFinfoGIntCompress(m *base.Module) int32 {
	return F_pg_finfo_g_int_compress(m)
}
func PgFinfoGIntConsistent(m *base.Module) int32 {
	return F_pg_finfo_g_int_consistent(m)
}
func PgFinfoGIntDecompress(m *base.Module) int32 {
	return F_pg_finfo_g_int_decompress(m)
}
func PgFinfoGIntOptions(m *base.Module) int32 {
	return F_pg_finfo_g_int_options(m)
}
func PgFinfoGIntPenalty(m *base.Module) int32 {
	return F_pg_finfo_g_int_penalty(m)
}
func PgFinfoGIntPicksplit(m *base.Module) int32 {
	return F_pg_finfo_g_int_picksplit(m)
}
func PgFinfoGIntSame(m *base.Module) int32 {
	return F_pg_finfo_g_int_same(m)
}
func PgFinfoGIntUnion(m *base.Module) int32 {
	return F_pg_finfo_g_int_union(m)
}
func PgFinfoGIntbigCompress(m *base.Module) int32 {
	return F_pg_finfo_g_intbig_compress(m)
}
func PgFinfoGIntbigConsistent(m *base.Module) int32 {
	return F_pg_finfo_g_intbig_consistent(m)
}
func PgFinfoGIntbigDecompress(m *base.Module) int32 {
	return F_pg_finfo_g_intbig_decompress(m)
}
func PgFinfoGIntbigOptions(m *base.Module) int32 {
	return F_pg_finfo_g_intbig_options(m)
}
func PgFinfoGIntbigPenalty(m *base.Module) int32 {
	return F_pg_finfo_g_intbig_penalty(m)
}
func PgFinfoGIntbigPicksplit(m *base.Module) int32 {
	return F_pg_finfo_g_intbig_picksplit(m)
}
func PgFinfoGIntbigSame(m *base.Module) int32 {
	return F_pg_finfo_g_intbig_same(m)
}
func PgFinfoGIntbigUnion(m *base.Module) int32 {
	return F_pg_finfo_g_intbig_union(m)
}
func PgFinfoGinint4Consistent(m *base.Module) int32 {
	return F_pg_finfo_ginint4_consistent(m)
}
func PgFinfoGinint4Queryextract(m *base.Module) int32 {
	return F_pg_finfo_ginint4_queryextract(m)
}
func PgFinfoIcount(m *base.Module) int32 {
	return F_pg_finfo_icount(m)
}
func PgFinfoIdx(m *base.Module) int32 {
	return F_pg_finfo_idx(m)
}
func PgFinfoIntarrayDelElem(m *base.Module) int32 {
	return F_pg_finfo_intarray_del_elem(m)
}
func PgFinfoIntarrayPushArray(m *base.Module) int32 {
	return F_pg_finfo_intarray_push_array(m)
}
func PgFinfoIntarrayPushElem(m *base.Module) int32 {
	return F_pg_finfo_intarray_push_elem(m)
}
func PgFinfoIntset(m *base.Module) int32 {
	return F_pg_finfo_intset(m)
}
func PgFinfoIntsetSubtract(m *base.Module) int32 {
	return F_pg_finfo_intset_subtract(m)
}
func PgFinfoIntsetUnionElem(m *base.Module) int32 {
	return F_pg_finfo_intset_union_elem(m)
}
func PgFinfoQuerytree(m *base.Module) int32 {
	return F_pg_finfo_querytree(m)
}
func PgFinfoRboolop(m *base.Module) int32 {
	return F_pg_finfo_rboolop(m)
}
func PgFinfoSort(m *base.Module) int32 {
	return F_pg_finfo_sort(m)
}
func PgFinfoSortAsc(m *base.Module) int32 {
	return F_pg_finfo_sort_asc(m)
}
func PgFinfoSortDesc(m *base.Module) int32 {
	return F_pg_finfo_sort_desc(m)
}
func PgFinfoSubarray(m *base.Module) int32 {
	return F_pg_finfo_subarray(m)
}
func PgFinfoUniq(m *base.Module) int32 {
	return F_pg_finfo_uniq(m)
}
func Querytree(m *base.Module, l0 int32) int32 {
	return F_querytree(m, l0)
}
func Rboolop(m *base.Module, l0 int32) int32 {
	return F_rboolop(m, l0)
}
func Sort(m *base.Module, l0 int32) int32 {
	return F_sort(m, l0)
}
func SortAsc(m *base.Module, l0 int32) int32 {
	return F_sort_asc(m, l0)
}
func SortDesc(m *base.Module, l0 int32) int32 {
	return F_sort_desc(m, l0)
}
func Subarray(m *base.Module, l0 int32) int32 {
	return F_subarray(m, l0)
}
func Uniq(m *base.Module, l0 int32) int32 {
	return F_uniq(m, l0)
}
func PgMagicFuncFuzzystrmatch(m *base.Module) int32 {
	return F_Pg_magic_func_fuzzystrmatch(m)
}
func DaitchMokotoff(m *base.Module, l0 int32) int32 {
	return F_daitch_mokotoff(m, l0)
}
func Difference(m *base.Module, l0 int32) int32 {
	return F_difference(m, l0)
}
func Dmetaphone(m *base.Module, l0 int32) int32 {
	return F_dmetaphone(m, l0)
}
func DmetaphoneAlt(m *base.Module, l0 int32) int32 {
	return F_dmetaphone_alt(m, l0)
}
func Levenshtein(m *base.Module, l0 int32) int32 {
	return F_levenshtein(m, l0)
}
func LevenshteinLessEqual(m *base.Module, l0 int32) int32 {
	return F_levenshtein_less_equal(m, l0)
}
func LevenshteinLessEqualWithCosts(m *base.Module, l0 int32) int32 {
	return F_levenshtein_less_equal_with_costs(m, l0)
}
func LevenshteinWithCosts(m *base.Module, l0 int32) int32 {
	return F_levenshtein_with_costs(m, l0)
}
func Metaphone(m *base.Module, l0 int32) int32 {
	return F_metaphone(m, l0)
}
func PgFinfoDaitchMokotoff(m *base.Module) int32 {
	return F_pg_finfo_daitch_mokotoff(m)
}
func PgFinfoDifference(m *base.Module) int32 {
	return F_pg_finfo_difference(m)
}
func PgFinfoDmetaphone(m *base.Module) int32 {
	return F_pg_finfo_dmetaphone(m)
}
func PgFinfoDmetaphoneAlt(m *base.Module) int32 {
	return F_pg_finfo_dmetaphone_alt(m)
}
func PgFinfoLevenshtein(m *base.Module) int32 {
	return F_pg_finfo_levenshtein(m)
}
func PgFinfoLevenshteinLessEqual(m *base.Module) int32 {
	return F_pg_finfo_levenshtein_less_equal(m)
}
func PgFinfoLevenshteinLessEqualWithCosts(m *base.Module) int32 {
	return F_pg_finfo_levenshtein_less_equal_with_costs(m)
}
func PgFinfoLevenshteinWithCosts(m *base.Module) int32 {
	return F_pg_finfo_levenshtein_with_costs(m)
}
func PgFinfoMetaphone(m *base.Module) int32 {
	return F_pg_finfo_metaphone(m)
}
func PgFinfoSoundex(m *base.Module) int32 {
	return F_pg_finfo_soundex(m)
}
func Soundex(m *base.Module, l0 int32) int32 {
	return F_soundex(m, l0)
}
func PgMagicFuncCube(m *base.Module) int32 {
	return F_Pg_magic_func_cube(m)
}
func CubeAF8(m *base.Module, l0 int32) int32 {
	return F_cube_a_f8(m, l0)
}
func CubeAF8F8(m *base.Module, l0 int32) int32 {
	return F_cube_a_f8_f8(m, l0)
}
func CubeCF8(m *base.Module, l0 int32) int32 {
	return F_cube_c_f8(m, l0)
}
func CubeCF8F8(m *base.Module, l0 int32) int32 {
	return F_cube_c_f8_f8(m, l0)
}
func CubeCmp(m *base.Module, l0 int32) int32 {
	return F_cube_cmp(m, l0)
}
func CubeContained(m *base.Module, l0 int32) int32 {
	return F_cube_contained(m, l0)
}
func CubeContains(m *base.Module, l0 int32) int32 {
	return F_cube_contains(m, l0)
}
func CubeCoord(m *base.Module, l0 int32) int32 {
	return F_cube_coord(m, l0)
}
func CubeCoordLlur(m *base.Module, l0 int32) int32 {
	return F_cube_coord_llur(m, l0)
}
func CubeDim(m *base.Module, l0 int32) int32 {
	return F_cube_dim(m, l0)
}
func CubeDistance(m *base.Module, l0 int32) int32 {
	return F_cube_distance(m, l0)
}
func CubeEnlarge(m *base.Module, l0 int32) int32 {
	return F_cube_enlarge(m, l0)
}
func CubeEq(m *base.Module, l0 int32) int32 {
	return F_cube_eq(m, l0)
}
func CubeF8(m *base.Module, l0 int32) int32 {
	return F_cube_f8(m, l0)
}
func CubeF8F8(m *base.Module, l0 int32) int32 {
	return F_cube_f8_f8(m, l0)
}
func CubeGe(m *base.Module, l0 int32) int32 {
	return F_cube_ge(m, l0)
}
func CubeGt(m *base.Module, l0 int32) int32 {
	return F_cube_gt(m, l0)
}
func CubeIn(m *base.Module, l0 int32) int32 {
	return F_cube_in(m, l0)
}
func CubeInter(m *base.Module, l0 int32) int32 {
	return F_cube_inter(m, l0)
}
func CubeIsPoint(m *base.Module, l0 int32) int32 {
	return F_cube_is_point(m, l0)
}
func CubeLe(m *base.Module, l0 int32) int32 {
	return F_cube_le(m, l0)
}
func CubeLlCoord(m *base.Module, l0 int32) int32 {
	return F_cube_ll_coord(m, l0)
}
func CubeLt(m *base.Module, l0 int32) int32 {
	return F_cube_lt(m, l0)
}
func CubeNe(m *base.Module, l0 int32) int32 {
	return F_cube_ne(m, l0)
}
func CubeOut(m *base.Module, l0 int32) int32 {
	return F_cube_out(m, l0)
}
func CubeOverlap(m *base.Module, l0 int32) int32 {
	return F_cube_overlap(m, l0)
}
func CubeRecv(m *base.Module, l0 int32) int32 {
	return F_cube_recv(m, l0)
}
func CubeSend(m *base.Module, l0 int32) int32 {
	return F_cube_send(m, l0)
}
func CubeSize(m *base.Module, l0 int32) int32 {
	return F_cube_size(m, l0)
}
func CubeSubset(m *base.Module, l0 int32) int32 {
	return F_cube_subset(m, l0)
}
func CubeUnion(m *base.Module, l0 int32) int32 {
	return F_cube_union(m, l0)
}
func CubeUrCoord(m *base.Module, l0 int32) int32 {
	return F_cube_ur_coord(m, l0)
}
func DistanceChebyshev(m *base.Module, l0 int32) int32 {
	return F_distance_chebyshev(m, l0)
}
func DistanceTaxicab(m *base.Module, l0 int32) int32 {
	return F_distance_taxicab(m, l0)
}
func GCubeCompress(m *base.Module, l0 int32) int32 {
	return F_float4up(m, l0)
}
func GCubeConsistent(m *base.Module, l0 int32) int32 {
	return F_g_cube_consistent(m, l0)
}
func GCubeDecompress(m *base.Module, l0 int32) int32 {
	return F_g_cube_decompress(m, l0)
}
func GCubeDistance(m *base.Module, l0 int32) int32 {
	return F_g_cube_distance(m, l0)
}
func GCubePenalty(m *base.Module, l0 int32) int32 {
	return F_g_cube_penalty(m, l0)
}
func GCubePicksplit(m *base.Module, l0 int32) int32 {
	return F_g_cube_picksplit(m, l0)
}
func GCubeSame(m *base.Module, l0 int32) int32 {
	return F_g_cube_same(m, l0)
}
func GCubeUnion(m *base.Module, l0 int32) int32 {
	return F_g_cube_union(m, l0)
}
func PgFinfoCubeAF8(m *base.Module) int32 {
	return F_pg_finfo_cube_a_f8(m)
}
func PgFinfoCubeAF8F8(m *base.Module) int32 {
	return F_pg_finfo_cube_a_f8_f8(m)
}
func PgFinfoCubeCF8(m *base.Module) int32 {
	return F_pg_finfo_cube_c_f8(m)
}
func PgFinfoCubeCF8F8(m *base.Module) int32 {
	return F_pg_finfo_cube_c_f8_f8(m)
}
func PgFinfoCubeCmp(m *base.Module) int32 {
	return F_pg_finfo_cube_cmp(m)
}
func PgFinfoCubeContained(m *base.Module) int32 {
	return F_pg_finfo_cube_contained(m)
}
func PgFinfoCubeContains(m *base.Module) int32 {
	return F_pg_finfo_cube_contains(m)
}
func PgFinfoCubeCoord(m *base.Module) int32 {
	return F_pg_finfo_cube_coord(m)
}
func PgFinfoCubeCoordLlur(m *base.Module) int32 {
	return F_pg_finfo_cube_coord_llur(m)
}
func PgFinfoCubeDim(m *base.Module) int32 {
	return F_pg_finfo_cube_dim(m)
}
func PgFinfoCubeDistance(m *base.Module) int32 {
	return F_pg_finfo_cube_distance(m)
}
func PgFinfoCubeEnlarge(m *base.Module) int32 {
	return F_pg_finfo_cube_enlarge(m)
}
func PgFinfoCubeEq(m *base.Module) int32 {
	return F_pg_finfo_cube_eq(m)
}
func PgFinfoCubeF8(m *base.Module) int32 {
	return F_pg_finfo_cube_f8(m)
}
func PgFinfoCubeF8F8(m *base.Module) int32 {
	return F_pg_finfo_cube_f8_f8(m)
}
func PgFinfoCubeGe(m *base.Module) int32 {
	return F_pg_finfo_cube_ge(m)
}
func PgFinfoCubeGt(m *base.Module) int32 {
	return F_pg_finfo_cube_gt(m)
}
func PgFinfoCubeIn(m *base.Module) int32 {
	return F_pg_finfo_cube_in(m)
}
func PgFinfoCubeInter(m *base.Module) int32 {
	return F_pg_finfo_cube_inter(m)
}
func PgFinfoCubeIsPoint(m *base.Module) int32 {
	return F_pg_finfo_cube_is_point(m)
}
func PgFinfoCubeLe(m *base.Module) int32 {
	return F_pg_finfo_cube_le(m)
}
func PgFinfoCubeLlCoord(m *base.Module) int32 {
	return F_pg_finfo_cube_ll_coord(m)
}
func PgFinfoCubeLt(m *base.Module) int32 {
	return F_pg_finfo_cube_lt(m)
}
func PgFinfoCubeNe(m *base.Module) int32 {
	return F_pg_finfo_cube_ne(m)
}
func PgFinfoCubeOut(m *base.Module) int32 {
	return F_pg_finfo_cube_out(m)
}
func PgFinfoCubeOverlap(m *base.Module) int32 {
	return F_pg_finfo_cube_overlap(m)
}
func PgFinfoCubeRecv(m *base.Module) int32 {
	return F_pg_finfo_cube_recv(m)
}
func PgFinfoCubeSend(m *base.Module) int32 {
	return F_pg_finfo_cube_send(m)
}
func PgFinfoCubeSize(m *base.Module) int32 {
	return F_pg_finfo_cube_size(m)
}
func PgFinfoCubeSubset(m *base.Module) int32 {
	return F_pg_finfo_cube_subset(m)
}
func PgFinfoCubeUnion(m *base.Module) int32 {
	return F_pg_finfo_cube_union(m)
}
func PgFinfoCubeUrCoord(m *base.Module) int32 {
	return F_pg_finfo_cube_ur_coord(m)
}
func PgFinfoDistanceChebyshev(m *base.Module) int32 {
	return F_pg_finfo_distance_chebyshev(m)
}
func PgFinfoDistanceTaxicab(m *base.Module) int32 {
	return F_pg_finfo_distance_taxicab(m)
}
func PgFinfoGCubeCompress(m *base.Module) int32 {
	return F_pg_finfo_g_cube_compress(m)
}
func PgFinfoGCubeConsistent(m *base.Module) int32 {
	return F_pg_finfo_g_cube_consistent(m)
}
func PgFinfoGCubeDecompress(m *base.Module) int32 {
	return F_pg_finfo_g_cube_decompress(m)
}
func PgFinfoGCubeDistance(m *base.Module) int32 {
	return F_pg_finfo_g_cube_distance(m)
}
func PgFinfoGCubePenalty(m *base.Module) int32 {
	return F_pg_finfo_g_cube_penalty(m)
}
func PgFinfoGCubePicksplit(m *base.Module) int32 {
	return F_pg_finfo_g_cube_picksplit(m)
}
func PgFinfoGCubeSame(m *base.Module) int32 {
	return F_pg_finfo_g_cube_same(m)
}
func PgFinfoGCubeUnion(m *base.Module) int32 {
	return F_pg_finfo_g_cube_union(m)
}
func PgMagicFuncEarthdistance(m *base.Module) int32 {
	return F_Pg_magic_func_earthdistance(m)
}
func GeoDistance(m *base.Module, l0 int32) int32 {
	return F_geo_distance(m, l0)
}
func PgFinfoGeoDistance(m *base.Module) int32 {
	return F_pg_finfo_geo_distance(m)
}
func PgMagicFuncSeg(m *base.Module) int32 {
	return F_Pg_magic_func_seg(m)
}
func GsegCompress(m *base.Module, l0 int32) int32 {
	return F_float4up(m, l0)
}
func GsegConsistent(m *base.Module, l0 int32) int32 {
	return F_gseg_consistent(m, l0)
}
func GsegDecompress(m *base.Module, l0 int32) int32 {
	return F_float4up(m, l0)
}
func GsegPenalty(m *base.Module, l0 int32) int32 {
	return F_gseg_penalty(m, l0)
}
func GsegPicksplit(m *base.Module, l0 int32) int32 {
	return F_gseg_picksplit(m, l0)
}
func GsegSame(m *base.Module, l0 int32) int32 {
	return F_gseg_same(m, l0)
}
func GsegUnion(m *base.Module, l0 int32) int32 {
	return F_gseg_union(m, l0)
}
func PgFinfoGsegCompress(m *base.Module) int32 {
	return F_pg_finfo_gseg_compress(m)
}
func PgFinfoGsegConsistent(m *base.Module) int32 {
	return F_pg_finfo_gseg_consistent(m)
}
func PgFinfoGsegDecompress(m *base.Module) int32 {
	return F_pg_finfo_gseg_decompress(m)
}
func PgFinfoGsegPenalty(m *base.Module) int32 {
	return F_pg_finfo_gseg_penalty(m)
}
func PgFinfoGsegPicksplit(m *base.Module) int32 {
	return F_pg_finfo_gseg_picksplit(m)
}
func PgFinfoGsegSame(m *base.Module) int32 {
	return F_pg_finfo_gseg_same(m)
}
func PgFinfoGsegUnion(m *base.Module) int32 {
	return F_pg_finfo_gseg_union(m)
}
func PgFinfoSegCenter(m *base.Module) int32 {
	return F_pg_finfo_seg_center(m)
}
func PgFinfoSegCmp(m *base.Module) int32 {
	return F_pg_finfo_seg_cmp(m)
}
func PgFinfoSegContained(m *base.Module) int32 {
	return F_pg_finfo_seg_contained(m)
}
func PgFinfoSegContains(m *base.Module) int32 {
	return F_pg_finfo_seg_contains(m)
}
func PgFinfoSegDifferent(m *base.Module) int32 {
	return F_pg_finfo_seg_different(m)
}
func PgFinfoSegGe(m *base.Module) int32 {
	return F_pg_finfo_seg_ge(m)
}
func PgFinfoSegGt(m *base.Module) int32 {
	return F_pg_finfo_seg_gt(m)
}
func PgFinfoSegIn(m *base.Module) int32 {
	return F_pg_finfo_seg_in(m)
}
func PgFinfoSegInter(m *base.Module) int32 {
	return F_pg_finfo_seg_inter(m)
}
func PgFinfoSegLe(m *base.Module) int32 {
	return F_pg_finfo_seg_le(m)
}
func PgFinfoSegLeft(m *base.Module) int32 {
	return F_pg_finfo_seg_left(m)
}
func PgFinfoSegLower(m *base.Module) int32 {
	return F_pg_finfo_seg_lower(m)
}
func PgFinfoSegLt(m *base.Module) int32 {
	return F_pg_finfo_seg_lt(m)
}
func PgFinfoSegOut(m *base.Module) int32 {
	return F_pg_finfo_seg_out(m)
}
func PgFinfoSegOverLeft(m *base.Module) int32 {
	return F_pg_finfo_seg_over_left(m)
}
func PgFinfoSegOverRight(m *base.Module) int32 {
	return F_pg_finfo_seg_over_right(m)
}
func PgFinfoSegOverlap(m *base.Module) int32 {
	return F_pg_finfo_seg_overlap(m)
}
func PgFinfoSegRight(m *base.Module) int32 {
	return F_pg_finfo_seg_right(m)
}
func PgFinfoSegSame(m *base.Module) int32 {
	return F_pg_finfo_seg_same(m)
}
func PgFinfoSegSize(m *base.Module) int32 {
	return F_pg_finfo_seg_size(m)
}
func PgFinfoSegUnion(m *base.Module) int32 {
	return F_pg_finfo_seg_union(m)
}
func PgFinfoSegUpper(m *base.Module) int32 {
	return F_pg_finfo_seg_upper(m)
}
func SegCenter(m *base.Module, l0 int32) int32 {
	return F_seg_center(m, l0)
}
func SegCmp(m *base.Module, l0 int32) int32 {
	return F_seg_cmp(m, l0)
}
func SegContained(m *base.Module, l0 int32) int32 {
	return F_seg_contained(m, l0)
}
func SegContains(m *base.Module, l0 int32) int32 {
	return F_seg_contains(m, l0)
}
func SegDifferent(m *base.Module, l0 int32) int32 {
	return F_seg_different(m, l0)
}
func SegGe(m *base.Module, l0 int32) int32 {
	return F_seg_ge(m, l0)
}
func SegGt(m *base.Module, l0 int32) int32 {
	return F_seg_gt(m, l0)
}
func SegIn(m *base.Module, l0 int32) int32 {
	return F_seg_in(m, l0)
}
func SegInter(m *base.Module, l0 int32) int32 {
	return F_seg_inter(m, l0)
}
func SegLe(m *base.Module, l0 int32) int32 {
	return F_seg_le(m, l0)
}
func SegLeft(m *base.Module, l0 int32) int32 {
	return F_seg_left(m, l0)
}
func SegLower(m *base.Module, l0 int32) int32 {
	return F_xid8toxid(m, l0)
}
func SegLt(m *base.Module, l0 int32) int32 {
	return F_seg_lt(m, l0)
}
func SegOut(m *base.Module, l0 int32) int32 {
	return F_seg_out(m, l0)
}
func SegOverLeft(m *base.Module, l0 int32) int32 {
	return F_seg_over_left(m, l0)
}
func SegOverRight(m *base.Module, l0 int32) int32 {
	return F_seg_over_right(m, l0)
}
func SegOverlap(m *base.Module, l0 int32) int32 {
	return F_seg_overlap(m, l0)
}
func SegRight(m *base.Module, l0 int32) int32 {
	return F_seg_right(m, l0)
}
func SegSame(m *base.Module, l0 int32) int32 {
	return F_seg_same(m, l0)
}
func SegSize(m *base.Module, l0 int32) int32 {
	return F_seg_size(m, l0)
}
func SegUnion(m *base.Module, l0 int32) int32 {
	return F_seg_union(m, l0)
}
func SegUpper(m *base.Module, l0 int32) int32 {
	return F_tsquery_numnode(m, l0)
}
func PgMagicFuncBloom(m *base.Module) int32 {
	return F_Pg_magic_func_bloom(m)
}
func PGInitBloom(m *base.Module) {
	F__PG_init_bloom(m)
}
func Blhandler(m *base.Module, l0 int32) int32 {
	return F_blhandler(m, l0)
}
func PgFinfoBlhandler(m *base.Module) int32 {
	return F_pg_finfo_blhandler(m)
}
func PgMagicFuncIsn(m *base.Module) int32 {
	return F_Pg_magic_func_isn(m)
}
func PGInitIsn(m *base.Module) {
	F__PG_init_isn(m)
}
func AcceptWeakInput(m *base.Module, l0 int32) int32 {
	return F_accept_weak_input(m, l0)
}
func Ean13In(m *base.Module, l0 int32) int32 {
	return F_ean13_in(m, l0)
}
func Ean13Out(m *base.Module, l0 int32) int32 {
	return F_ean13_out(m, l0)
}
func IsValid(m *base.Module, l0 int32) int32 {
	return F_is_valid(m, l0)
}
func IsbnCastFromEan13(m *base.Module, l0 int32) int32 {
	return F_isbn_cast_from_ean13(m, l0)
}
func IsbnIn(m *base.Module, l0 int32) int32 {
	return F_isbn_in(m, l0)
}
func IsmnCastFromEan13(m *base.Module, l0 int32) int32 {
	return F_ismn_cast_from_ean13(m, l0)
}
func IsmnIn(m *base.Module, l0 int32) int32 {
	return F_ismn_in(m, l0)
}
func IsnOut(m *base.Module, l0 int32) int32 {
	return F_isn_out(m, l0)
}
func IssnCastFromEan13(m *base.Module, l0 int32) int32 {
	return F_issn_cast_from_ean13(m, l0)
}
func IssnIn(m *base.Module, l0 int32) int32 {
	return F_issn_in(m, l0)
}
func MakeValid(m *base.Module, l0 int32) int32 {
	return F_make_valid(m, l0)
}
func PgFinfoAcceptWeakInput(m *base.Module) int32 {
	return F_pg_finfo_accept_weak_input(m)
}
func PgFinfoEan13In(m *base.Module) int32 {
	return F_pg_finfo_ean13_in(m)
}
func PgFinfoEan13Out(m *base.Module) int32 {
	return F_pg_finfo_ean13_out(m)
}
func PgFinfoIsValid(m *base.Module) int32 {
	return F_pg_finfo_is_valid(m)
}
func PgFinfoIsbnCastFromEan13(m *base.Module) int32 {
	return F_pg_finfo_isbn_cast_from_ean13(m)
}
func PgFinfoIsbnIn(m *base.Module) int32 {
	return F_pg_finfo_isbn_in(m)
}
func PgFinfoIsmnCastFromEan13(m *base.Module) int32 {
	return F_pg_finfo_ismn_cast_from_ean13(m)
}
func PgFinfoIsmnIn(m *base.Module) int32 {
	return F_pg_finfo_ismn_in(m)
}
func PgFinfoIsnOut(m *base.Module) int32 {
	return F_pg_finfo_isn_out(m)
}
func PgFinfoIssnCastFromEan13(m *base.Module) int32 {
	return F_pg_finfo_issn_cast_from_ean13(m)
}
func PgFinfoIssnIn(m *base.Module) int32 {
	return F_pg_finfo_issn_in(m)
}
func PgFinfoMakeValid(m *base.Module) int32 {
	return F_pg_finfo_make_valid(m)
}
func PgFinfoUpcCastFromEan13(m *base.Module) int32 {
	return F_pg_finfo_upc_cast_from_ean13(m)
}
func PgFinfoUpcIn(m *base.Module) int32 {
	return F_pg_finfo_upc_in(m)
}
func PgFinfoWeakInputStatus(m *base.Module) int32 {
	return F_pg_finfo_weak_input_status(m)
}
func UpcCastFromEan13(m *base.Module, l0 int32) int32 {
	return F_upc_cast_from_ean13(m, l0)
}
func UpcIn(m *base.Module, l0 int32) int32 {
	return F_upc_in(m, l0)
}
func WeakInputStatus(m *base.Module, l0 int32) int32 {
	return F_weak_input_status(m, l0)
}
func PgMagicFuncDictInt(m *base.Module) int32 {
	return F_Pg_magic_func_dict_int(m)
}
func DintdictInit(m *base.Module, l0 int32) int32 {
	return F_dintdict_init(m, l0)
}
func DintdictLexize(m *base.Module, l0 int32) int32 {
	return F_dintdict_lexize(m, l0)
}
func PgFinfoDintdictInit(m *base.Module) int32 {
	return F_pg_finfo_dintdict_init(m)
}
func PgFinfoDintdictLexize(m *base.Module) int32 {
	return F_pg_finfo_dintdict_lexize(m)
}
func PgMagicFuncDictXsyn(m *base.Module) int32 {
	return F_Pg_magic_func_dict_xsyn(m)
}
func DxsynInit(m *base.Module, l0 int32) int32 {
	return F_dxsyn_init(m, l0)
}
func DxsynLexize(m *base.Module, l0 int32) int32 {
	return F_dxsyn_lexize(m, l0)
}
func PgFinfoDxsynInit(m *base.Module) int32 {
	return F_pg_finfo_dxsyn_init(m)
}
func PgFinfoDxsynLexize(m *base.Module) int32 {
	return F_pg_finfo_dxsyn_lexize(m)
}
func PgMagicFuncLo(m *base.Module) int32 {
	return F_Pg_magic_func_lo(m)
}
func LoManage(m *base.Module, l0 int32) int32 {
	return F_lo_manage(m, l0)
}
func PgFinfoLoManage(m *base.Module) int32 {
	return F_pg_finfo_lo_manage(m)
}
func PgMagicFuncTsmSystemRows(m *base.Module) int32 {
	return F_Pg_magic_func_tsm_system_rows(m)
}
func PgFinfoTsmSystemRowsHandler(m *base.Module) int32 {
	return F_pg_finfo_tsm_system_rows_handler(m)
}
func TsmSystemRowsHandler(m *base.Module, l0 int32) int32 {
	return F_tsm_system_rows_handler(m, l0)
}
func PgMagicFuncTsmSystemTime(m *base.Module) int32 {
	return F_Pg_magic_func_tsm_system_time(m)
}
func PgFinfoTsmSystemTimeHandler(m *base.Module) int32 {
	return F_pg_finfo_tsm_system_time_handler(m)
}
func TsmSystemTimeHandler(m *base.Module, l0 int32) int32 {
	return F_tsm_system_time_handler(m, l0)
}
func PgMagicFuncPgstattuple(m *base.Module) int32 {
	return F_Pg_magic_func_pgstattuple(m)
}
func PgFinfoPgRelpages(m *base.Module) int32 {
	return F_pg_finfo_pg_relpages(m)
}
func PgFinfoPgRelpagesV1_5(m *base.Module) int32 {
	return F_pg_finfo_pg_relpages_v1_5(m)
}
func PgFinfoPgRelpagesbyid(m *base.Module) int32 {
	return F_pg_finfo_pg_relpagesbyid(m)
}
func PgFinfoPgRelpagesbyidV1_5(m *base.Module) int32 {
	return F_pg_finfo_pg_relpagesbyid_v1_5(m)
}
func PgFinfoPgstatginindex(m *base.Module) int32 {
	return F_pg_finfo_pgstatginindex(m)
}
func PgFinfoPgstatginindexV1_5(m *base.Module) int32 {
	return F_pg_finfo_pgstatginindex_v1_5(m)
}
func PgFinfoPgstathashindex(m *base.Module) int32 {
	return F_pg_finfo_pgstathashindex(m)
}
func PgFinfoPgstatindex(m *base.Module) int32 {
	return F_pg_finfo_pgstatindex(m)
}
func PgFinfoPgstatindexV1_5(m *base.Module) int32 {
	return F_pg_finfo_pgstatindex_v1_5(m)
}
func PgFinfoPgstatindexbyid(m *base.Module) int32 {
	return F_pg_finfo_pgstatindexbyid(m)
}
func PgFinfoPgstatindexbyidV1_5(m *base.Module) int32 {
	return F_pg_finfo_pgstatindexbyid_v1_5(m)
}
func PgFinfoPgstattuple(m *base.Module) int32 {
	return F_pg_finfo_pgstattuple(m)
}
func PgFinfoPgstattupleApprox(m *base.Module) int32 {
	return F_pg_finfo_pgstattuple_approx(m)
}
func PgFinfoPgstattupleApproxV1_5(m *base.Module) int32 {
	return F_pg_finfo_pgstattuple_approx_v1_5(m)
}
func PgFinfoPgstattupleV1_5(m *base.Module) int32 {
	return F_pg_finfo_pgstattuple_v1_5(m)
}
func PgFinfoPgstattuplebyid(m *base.Module) int32 {
	return F_pg_finfo_pgstattuplebyid(m)
}
func PgFinfoPgstattuplebyidV1_5(m *base.Module) int32 {
	return F_pg_finfo_pgstattuplebyid_v1_5(m)
}
func PgRelpages(m *base.Module, l0 int32) int32 {
	return F_pg_relpages(m, l0)
}
func PgRelpagesV1_5(m *base.Module, l0 int32) int32 {
	return F_pg_relpages_v1_5(m, l0)
}
func PgRelpagesbyid(m *base.Module, l0 int32) int32 {
	return F_pg_relpagesbyid(m, l0)
}
func PgRelpagesbyidV1_5(m *base.Module, l0 int32) int32 {
	return F_pg_relpagesbyid_v1_5(m, l0)
}
func Pgstatginindex(m *base.Module, l0 int32) int32 {
	return F_pgstatginindex(m, l0)
}
func PgstatginindexV1_5(m *base.Module, l0 int32) int32 {
	return F_pgstatginindex_v1_5(m, l0)
}
func Pgstathashindex(m *base.Module, l0 int32) int32 {
	return F_pgstathashindex(m, l0)
}
func Pgstatindex(m *base.Module, l0 int32) int32 {
	return F_pgstatindex(m, l0)
}
func PgstatindexV1_5(m *base.Module, l0 int32) int32 {
	return F_pgstatindex_v1_5(m, l0)
}
func Pgstatindexbyid(m *base.Module, l0 int32) int32 {
	return F_pgstatindexbyid(m, l0)
}
func PgstatindexbyidV1_5(m *base.Module, l0 int32) int32 {
	return F_pgstatindexbyid_v1_5(m, l0)
}
func Pgstattuple(m *base.Module, l0 int32) int32 {
	return F_pgstattuple(m, l0)
}
func PgstattupleApprox(m *base.Module, l0 int32) int32 {
	return F_pgstattuple_approx(m, l0)
}
func PgstattupleApproxV1_5(m *base.Module, l0 int32) int32 {
	return F_pgstattuple_approx_v1_5(m, l0)
}
func PgstattupleV1_5(m *base.Module, l0 int32) int32 {
	return F_pgstattuple_v1_5(m, l0)
}
func Pgstattuplebyid(m *base.Module, l0 int32) int32 {
	return F_pgstattuplebyid(m, l0)
}
func PgstattuplebyidV1_5(m *base.Module, l0 int32) int32 {
	return F_pgstattuplebyid_v1_5(m, l0)
}
func PgMagicFuncUuidOssp(m *base.Module) int32 {
	return F_Pg_magic_func_uuid_ossp(m)
}
func PgFinfoUuidGenerateV1(m *base.Module) int32 {
	return F_pg_finfo_uuid_generate_v1(m)
}
func PgFinfoUuidGenerateV1mc(m *base.Module) int32 {
	return F_pg_finfo_uuid_generate_v1mc(m)
}
func PgFinfoUuidGenerateV3(m *base.Module) int32 {
	return F_pg_finfo_uuid_generate_v3(m)
}
func PgFinfoUuidGenerateV4(m *base.Module) int32 {
	return F_pg_finfo_uuid_generate_v4(m)
}
func PgFinfoUuidGenerateV5(m *base.Module) int32 {
	return F_pg_finfo_uuid_generate_v5(m)
}
func PgFinfoUuidNil(m *base.Module) int32 {
	return F_pg_finfo_uuid_nil(m)
}
func PgFinfoUuidNsDns(m *base.Module) int32 {
	return F_pg_finfo_uuid_ns_dns(m)
}
func PgFinfoUuidNsOid(m *base.Module) int32 {
	return F_pg_finfo_uuid_ns_oid(m)
}
func PgFinfoUuidNsUrl(m *base.Module) int32 {
	return F_pg_finfo_uuid_ns_url(m)
}
func PgFinfoUuidNsX500(m *base.Module) int32 {
	return F_pg_finfo_uuid_ns_x500(m)
}
func UuidGenerateV1(m *base.Module, l0 int32) int32 {
	return F_uuid_generate_v1(m, l0)
}
func UuidGenerateV1mc(m *base.Module, l0 int32) int32 {
	return F_uuid_generate_v1mc(m, l0)
}
func UuidGenerateV3(m *base.Module, l0 int32) int32 {
	return F_uuid_generate_v3(m, l0)
}
func UuidGenerateV4(m *base.Module, l0 int32) int32 {
	return F_uuid_generate_v4(m, l0)
}
func UuidGenerateV5(m *base.Module, l0 int32) int32 {
	return F_uuid_generate_v5(m, l0)
}
func UuidNil(m *base.Module, l0 int32) int32 {
	return F_uuid_nil(m, l0)
}
func UuidNsDns(m *base.Module, l0 int32) int32 {
	return F_uuid_ns_dns(m, l0)
}
func UuidNsOid(m *base.Module, l0 int32) int32 {
	return F_uuid_ns_oid(m, l0)
}
func UuidNsUrl(m *base.Module, l0 int32) int32 {
	return F_uuid_ns_url(m, l0)
}
func UuidNsX500(m *base.Module, l0 int32) int32 {
	return F_uuid_ns_x500(m, l0)
}
func HnswParallelBuildMain(m *base.Module, l0 int32, l1 int32) {
	F_HnswParallelBuildMain(m, l0, l1)
}
func IvfflatParallelBuildMain(m *base.Module, l0 int32, l1 int32) {
	F_IvfflatParallelBuildMain(m, l0, l1)
}
func PgMagicFuncVector(m *base.Module) int32 {
	return F_Pg_magic_func_vector(m)
}
func PGInitVector(m *base.Module) {
	F__PG_init_vector(m)
}
func ArrayToHalfvec(m *base.Module, l0 int32) int32 {
	return F_array_to_halfvec(m, l0)
}
func ArrayToSparsevec(m *base.Module, l0 int32) int32 {
	return F_array_to_sparsevec(m, l0)
}
func ArrayToVector(m *base.Module, l0 int32) int32 {
	return F_array_to_vector(m, l0)
}
func BinaryQuantize(m *base.Module, l0 int32) int32 {
	return F_binary_quantize(m, l0)
}
func CosineDistance(m *base.Module, l0 int32) int32 {
	return F_cosine_distance(m, l0)
}
func Halfvec(m *base.Module, l0 int32) int32 {
	return F_halfvec(m, l0)
}
func HalfvecAccum(m *base.Module, l0 int32) int32 {
	return F_halfvec_accum(m, l0)
}
func HalfvecAdd(m *base.Module, l0 int32) int32 {
	return F_halfvec_add(m, l0)
}
func HalfvecAvg(m *base.Module, l0 int32) int32 {
	return F_halfvec_avg(m, l0)
}
func HalfvecBinaryQuantize(m *base.Module, l0 int32) int32 {
	return F_halfvec_binary_quantize(m, l0)
}
func HalfvecCmp(m *base.Module, l0 int32) int32 {
	return F_halfvec_cmp(m, l0)
}
func HalfvecConcat(m *base.Module, l0 int32) int32 {
	return F_halfvec_concat(m, l0)
}
func HalfvecCosineDistance(m *base.Module, l0 int32) int32 {
	return F_halfvec_cosine_distance(m, l0)
}
func HalfvecEq(m *base.Module, l0 int32) int32 {
	return F_halfvec_eq(m, l0)
}
func HalfvecGe(m *base.Module, l0 int32) int32 {
	return F_halfvec_ge(m, l0)
}
func HalfvecGt(m *base.Module, l0 int32) int32 {
	return F_halfvec_gt(m, l0)
}
func HalfvecIn(m *base.Module, l0 int32) int32 {
	return F_halfvec_in(m, l0)
}
func HalfvecInnerProduct(m *base.Module, l0 int32) int32 {
	return F_halfvec_inner_product(m, l0)
}
func HalfvecL1Distance(m *base.Module, l0 int32) int32 {
	return F_halfvec_l1_distance(m, l0)
}
func HalfvecL2Distance(m *base.Module, l0 int32) int32 {
	return F_halfvec_l2_distance(m, l0)
}
func HalfvecL2Norm(m *base.Module, l0 int32) int32 {
	return F_halfvec_l2_norm(m, l0)
}
func HalfvecL2Normalize(m *base.Module, l0 int32) int32 {
	return F_halfvec_l2_normalize(m, l0)
}
func HalfvecL2SquaredDistance(m *base.Module, l0 int32) int32 {
	return F_halfvec_l2_squared_distance(m, l0)
}
func HalfvecLe(m *base.Module, l0 int32) int32 {
	return F_halfvec_le(m, l0)
}
func HalfvecLt(m *base.Module, l0 int32) int32 {
	return F_halfvec_lt(m, l0)
}
func HalfvecMul(m *base.Module, l0 int32) int32 {
	return F_halfvec_mul(m, l0)
}
func HalfvecNe(m *base.Module, l0 int32) int32 {
	return F_halfvec_ne(m, l0)
}
func HalfvecNegativeInnerProduct(m *base.Module, l0 int32) int32 {
	return F_halfvec_negative_inner_product(m, l0)
}
func HalfvecOut(m *base.Module, l0 int32) int32 {
	return F_halfvec_out(m, l0)
}
func HalfvecRecv(m *base.Module, l0 int32) int32 {
	return F_halfvec_recv(m, l0)
}
func HalfvecSend(m *base.Module, l0 int32) int32 {
	return F_halfvec_send(m, l0)
}
func HalfvecSphericalDistance(m *base.Module, l0 int32) int32 {
	return F_halfvec_spherical_distance(m, l0)
}
func HalfvecSub(m *base.Module, l0 int32) int32 {
	return F_halfvec_sub(m, l0)
}
func HalfvecSubvector(m *base.Module, l0 int32) int32 {
	return F_halfvec_subvector(m, l0)
}
func HalfvecToFloat4(m *base.Module, l0 int32) int32 {
	return F_halfvec_to_float4(m, l0)
}
func HalfvecToSparsevec(m *base.Module, l0 int32) int32 {
	return F_halfvec_to_sparsevec(m, l0)
}
func HalfvecToVector(m *base.Module, l0 int32) int32 {
	return F_halfvec_to_vector(m, l0)
}
func HalfvecTypmodIn(m *base.Module, l0 int32) int32 {
	return F_halfvec_typmod_in(m, l0)
}
func HalfvecVectorDims(m *base.Module, l0 int32) int32 {
	return F_halfvec_vector_dims(m, l0)
}
func HammingDistance(m *base.Module, l0 int32) int32 {
	return F_hamming_distance(m, l0)
}
func HnswBitSupport(m *base.Module, l0 int32) int32 {
	return F_hnsw_bit_support(m, l0)
}
func HnswHalfvecSupport(m *base.Module, l0 int32) int32 {
	return F_hnsw_halfvec_support(m, l0)
}
func HnswSparsevecSupport(m *base.Module, l0 int32) int32 {
	return F_hnsw_sparsevec_support(m, l0)
}
func Hnswhandler(m *base.Module, l0 int32) int32 {
	return F_hnswhandler(m, l0)
}
func InnerProduct(m *base.Module, l0 int32) int32 {
	return F_inner_product(m, l0)
}
func IvfflatBitSupport(m *base.Module, l0 int32) int32 {
	return F_ivfflat_bit_support(m, l0)
}
func IvfflatHalfvecSupport(m *base.Module, l0 int32) int32 {
	return F_ivfflat_halfvec_support(m, l0)
}
func Ivfflathandler(m *base.Module, l0 int32) int32 {
	return F_ivfflathandler(m, l0)
}
func JaccardDistance(m *base.Module, l0 int32) int32 {
	return F_jaccard_distance(m, l0)
}
func L1Distance(m *base.Module, l0 int32) int32 {
	return F_l1_distance(m, l0)
}
func L2Distance(m *base.Module, l0 int32) int32 {
	return F_l2_distance(m, l0)
}
func L2Normalize(m *base.Module, l0 int32) int32 {
	return F_l2_normalize(m, l0)
}
func PgFinfoArrayToHalfvec(m *base.Module) int32 {
	return F_pg_finfo_array_to_halfvec(m)
}
func PgFinfoArrayToSparsevec(m *base.Module) int32 {
	return F_pg_finfo_array_to_sparsevec(m)
}
func PgFinfoArrayToVector(m *base.Module) int32 {
	return F_pg_finfo_array_to_vector(m)
}
func PgFinfoBinaryQuantize(m *base.Module) int32 {
	return F_pg_finfo_binary_quantize(m)
}
func PgFinfoCosineDistance(m *base.Module) int32 {
	return F_pg_finfo_cosine_distance(m)
}
func PgFinfoHalfvec(m *base.Module) int32 {
	return F_pg_finfo_halfvec(m)
}
func PgFinfoHalfvecAccum(m *base.Module) int32 {
	return F_pg_finfo_halfvec_accum(m)
}
func PgFinfoHalfvecAdd(m *base.Module) int32 {
	return F_pg_finfo_halfvec_add(m)
}
func PgFinfoHalfvecAvg(m *base.Module) int32 {
	return F_pg_finfo_halfvec_avg(m)
}
func PgFinfoHalfvecBinaryQuantize(m *base.Module) int32 {
	return F_pg_finfo_halfvec_binary_quantize(m)
}
func PgFinfoHalfvecCmp(m *base.Module) int32 {
	return F_pg_finfo_halfvec_cmp(m)
}
func PgFinfoHalfvecConcat(m *base.Module) int32 {
	return F_pg_finfo_halfvec_concat(m)
}
func PgFinfoHalfvecCosineDistance(m *base.Module) int32 {
	return F_pg_finfo_halfvec_cosine_distance(m)
}
func PgFinfoHalfvecEq(m *base.Module) int32 {
	return F_pg_finfo_halfvec_eq(m)
}
func PgFinfoHalfvecGe(m *base.Module) int32 {
	return F_pg_finfo_halfvec_ge(m)
}
func PgFinfoHalfvecGt(m *base.Module) int32 {
	return F_pg_finfo_halfvec_gt(m)
}
func PgFinfoHalfvecIn(m *base.Module) int32 {
	return F_pg_finfo_halfvec_in(m)
}
func PgFinfoHalfvecInnerProduct(m *base.Module) int32 {
	return F_pg_finfo_halfvec_inner_product(m)
}
func PgFinfoHalfvecL1Distance(m *base.Module) int32 {
	return F_pg_finfo_halfvec_l1_distance(m)
}
func PgFinfoHalfvecL2Distance(m *base.Module) int32 {
	return F_pg_finfo_halfvec_l2_distance(m)
}
func PgFinfoHalfvecL2Norm(m *base.Module) int32 {
	return F_pg_finfo_halfvec_l2_norm(m)
}
func PgFinfoHalfvecL2Normalize(m *base.Module) int32 {
	return F_pg_finfo_halfvec_l2_normalize(m)
}
func PgFinfoHalfvecL2SquaredDistance(m *base.Module) int32 {
	return F_pg_finfo_halfvec_l2_squared_distance(m)
}
func PgFinfoHalfvecLe(m *base.Module) int32 {
	return F_pg_finfo_halfvec_le(m)
}
func PgFinfoHalfvecLt(m *base.Module) int32 {
	return F_pg_finfo_halfvec_lt(m)
}
func PgFinfoHalfvecMul(m *base.Module) int32 {
	return F_pg_finfo_halfvec_mul(m)
}
func PgFinfoHalfvecNe(m *base.Module) int32 {
	return F_pg_finfo_halfvec_ne(m)
}
func PgFinfoHalfvecNegativeInnerProduct(m *base.Module) int32 {
	return F_pg_finfo_halfvec_negative_inner_product(m)
}
func PgFinfoHalfvecOut(m *base.Module) int32 {
	return F_pg_finfo_halfvec_out(m)
}
func PgFinfoHalfvecRecv(m *base.Module) int32 {
	return F_pg_finfo_halfvec_recv(m)
}
func PgFinfoHalfvecSend(m *base.Module) int32 {
	return F_pg_finfo_halfvec_send(m)
}
func PgFinfoHalfvecSphericalDistance(m *base.Module) int32 {
	return F_pg_finfo_halfvec_spherical_distance(m)
}
func PgFinfoHalfvecSub(m *base.Module) int32 {
	return F_pg_finfo_halfvec_sub(m)
}
func PgFinfoHalfvecSubvector(m *base.Module) int32 {
	return F_pg_finfo_halfvec_subvector(m)
}
func PgFinfoHalfvecToFloat4(m *base.Module) int32 {
	return F_pg_finfo_halfvec_to_float4(m)
}
func PgFinfoHalfvecToSparsevec(m *base.Module) int32 {
	return F_pg_finfo_halfvec_to_sparsevec(m)
}
func PgFinfoHalfvecToVector(m *base.Module) int32 {
	return F_pg_finfo_halfvec_to_vector(m)
}
func PgFinfoHalfvecTypmodIn(m *base.Module) int32 {
	return F_pg_finfo_halfvec_typmod_in(m)
}
func PgFinfoHalfvecVectorDims(m *base.Module) int32 {
	return F_pg_finfo_halfvec_vector_dims(m)
}
func PgFinfoHammingDistance(m *base.Module) int32 {
	return F_pg_finfo_hamming_distance(m)
}
func PgFinfoHnswBitSupport(m *base.Module) int32 {
	return F_pg_finfo_hnsw_bit_support(m)
}
func PgFinfoHnswHalfvecSupport(m *base.Module) int32 {
	return F_pg_finfo_hnsw_halfvec_support(m)
}
func PgFinfoHnswSparsevecSupport(m *base.Module) int32 {
	return F_pg_finfo_hnsw_sparsevec_support(m)
}
func PgFinfoHnswhandler(m *base.Module) int32 {
	return F_pg_finfo_hnswhandler(m)
}
func PgFinfoInnerProduct(m *base.Module) int32 {
	return F_pg_finfo_inner_product(m)
}
func PgFinfoIvfflatBitSupport(m *base.Module) int32 {
	return F_pg_finfo_ivfflat_bit_support(m)
}
func PgFinfoIvfflatHalfvecSupport(m *base.Module) int32 {
	return F_pg_finfo_ivfflat_halfvec_support(m)
}
func PgFinfoIvfflathandler(m *base.Module) int32 {
	return F_pg_finfo_ivfflathandler(m)
}
func PgFinfoJaccardDistance(m *base.Module) int32 {
	return F_pg_finfo_jaccard_distance(m)
}
func PgFinfoL1Distance(m *base.Module) int32 {
	return F_pg_finfo_l1_distance(m)
}
func PgFinfoL2Distance(m *base.Module) int32 {
	return F_pg_finfo_l2_distance(m)
}
func PgFinfoL2Normalize(m *base.Module) int32 {
	return F_pg_finfo_l2_normalize(m)
}
func PgFinfoSparsevec(m *base.Module) int32 {
	return F_pg_finfo_sparsevec(m)
}
func PgFinfoSparsevecCmp(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_cmp(m)
}
func PgFinfoSparsevecCosineDistance(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_cosine_distance(m)
}
func PgFinfoSparsevecEq(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_eq(m)
}
func PgFinfoSparsevecGe(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_ge(m)
}
func PgFinfoSparsevecGt(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_gt(m)
}
func PgFinfoSparsevecIn(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_in(m)
}
func PgFinfoSparsevecInnerProduct(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_inner_product(m)
}
func PgFinfoSparsevecL1Distance(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_l1_distance(m)
}
func PgFinfoSparsevecL2Distance(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_l2_distance(m)
}
func PgFinfoSparsevecL2Norm(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_l2_norm(m)
}
func PgFinfoSparsevecL2Normalize(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_l2_normalize(m)
}
func PgFinfoSparsevecL2SquaredDistance(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_l2_squared_distance(m)
}
func PgFinfoSparsevecLe(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_le(m)
}
func PgFinfoSparsevecLt(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_lt(m)
}
func PgFinfoSparsevecNe(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_ne(m)
}
func PgFinfoSparsevecNegativeInnerProduct(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_negative_inner_product(m)
}
func PgFinfoSparsevecOut(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_out(m)
}
func PgFinfoSparsevecRecv(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_recv(m)
}
func PgFinfoSparsevecSend(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_send(m)
}
func PgFinfoSparsevecToHalfvec(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_to_halfvec(m)
}
func PgFinfoSparsevecToVector(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_to_vector(m)
}
func PgFinfoSparsevecTypmodIn(m *base.Module) int32 {
	return F_pg_finfo_sparsevec_typmod_in(m)
}
func PgFinfoSubvector(m *base.Module) int32 {
	return F_pg_finfo_subvector(m)
}
func PgFinfoVector(m *base.Module) int32 {
	return F_pg_finfo_vector(m)
}
func PgFinfoVectorAccum(m *base.Module) int32 {
	return F_pg_finfo_vector_accum(m)
}
func PgFinfoVectorAdd(m *base.Module) int32 {
	return F_pg_finfo_vector_add(m)
}
func PgFinfoVectorAvg(m *base.Module) int32 {
	return F_pg_finfo_vector_avg(m)
}
func PgFinfoVectorCmp(m *base.Module) int32 {
	return F_pg_finfo_vector_cmp(m)
}
func PgFinfoVectorCombine(m *base.Module) int32 {
	return F_pg_finfo_vector_combine(m)
}
func PgFinfoVectorConcat(m *base.Module) int32 {
	return F_pg_finfo_vector_concat(m)
}
func PgFinfoVectorDims(m *base.Module) int32 {
	return F_pg_finfo_vector_dims(m)
}
func PgFinfoVectorEq(m *base.Module) int32 {
	return F_pg_finfo_vector_eq(m)
}
func PgFinfoVectorGe(m *base.Module) int32 {
	return F_pg_finfo_vector_ge(m)
}
func PgFinfoVectorGt(m *base.Module) int32 {
	return F_pg_finfo_vector_gt(m)
}
func PgFinfoVectorIn(m *base.Module) int32 {
	return F_pg_finfo_vector_in(m)
}
func PgFinfoVectorL2SquaredDistance(m *base.Module) int32 {
	return F_pg_finfo_vector_l2_squared_distance(m)
}
func PgFinfoVectorLe(m *base.Module) int32 {
	return F_pg_finfo_vector_le(m)
}
func PgFinfoVectorLt(m *base.Module) int32 {
	return F_pg_finfo_vector_lt(m)
}
func PgFinfoVectorMul(m *base.Module) int32 {
	return F_pg_finfo_vector_mul(m)
}
func PgFinfoVectorNe(m *base.Module) int32 {
	return F_pg_finfo_vector_ne(m)
}
func PgFinfoVectorNegativeInnerProduct(m *base.Module) int32 {
	return F_pg_finfo_vector_negative_inner_product(m)
}
func PgFinfoVectorNorm(m *base.Module) int32 {
	return F_pg_finfo_vector_norm(m)
}
func PgFinfoVectorOut(m *base.Module) int32 {
	return F_pg_finfo_vector_out(m)
}
func PgFinfoVectorRecv(m *base.Module) int32 {
	return F_pg_finfo_vector_recv(m)
}
func PgFinfoVectorSend(m *base.Module) int32 {
	return F_pg_finfo_vector_send(m)
}
func PgFinfoVectorSphericalDistance(m *base.Module) int32 {
	return F_pg_finfo_vector_spherical_distance(m)
}
func PgFinfoVectorSub(m *base.Module) int32 {
	return F_pg_finfo_vector_sub(m)
}
func PgFinfoVectorToFloat4(m *base.Module) int32 {
	return F_pg_finfo_vector_to_float4(m)
}
func PgFinfoVectorToHalfvec(m *base.Module) int32 {
	return F_pg_finfo_vector_to_halfvec(m)
}
func PgFinfoVectorToSparsevec(m *base.Module) int32 {
	return F_pg_finfo_vector_to_sparsevec(m)
}
func PgFinfoVectorTypmodIn(m *base.Module) int32 {
	return F_pg_finfo_vector_typmod_in(m)
}
func Sparsevec(m *base.Module, l0 int32) int32 {
	return F_sparsevec(m, l0)
}
func SparsevecCmp(m *base.Module, l0 int32) int32 {
	return F_sparsevec_cmp(m, l0)
}
func SparsevecCosineDistance(m *base.Module, l0 int32) int32 {
	return F_sparsevec_cosine_distance(m, l0)
}
func SparsevecEq(m *base.Module, l0 int32) int32 {
	return F_sparsevec_eq(m, l0)
}
func SparsevecGe(m *base.Module, l0 int32) int32 {
	return F_sparsevec_ge(m, l0)
}
func SparsevecGt(m *base.Module, l0 int32) int32 {
	return F_sparsevec_gt(m, l0)
}
func SparsevecIn(m *base.Module, l0 int32) int32 {
	return F_sparsevec_in(m, l0)
}
func SparsevecInnerProduct(m *base.Module, l0 int32) int32 {
	return F_sparsevec_inner_product(m, l0)
}
func SparsevecL1Distance(m *base.Module, l0 int32) int32 {
	return F_sparsevec_l1_distance(m, l0)
}
func SparsevecL2Distance(m *base.Module, l0 int32) int32 {
	return F_sparsevec_l2_distance(m, l0)
}
func SparsevecL2Norm(m *base.Module, l0 int32) int32 {
	return F_sparsevec_l2_norm(m, l0)
}
func SparsevecL2Normalize(m *base.Module, l0 int32) int32 {
	return F_sparsevec_l2_normalize(m, l0)
}
func SparsevecL2SquaredDistance(m *base.Module, l0 int32) int32 {
	return F_sparsevec_l2_squared_distance(m, l0)
}
func SparsevecLe(m *base.Module, l0 int32) int32 {
	return F_sparsevec_le(m, l0)
}
func SparsevecLt(m *base.Module, l0 int32) int32 {
	return F_sparsevec_lt(m, l0)
}
func SparsevecNe(m *base.Module, l0 int32) int32 {
	return F_sparsevec_ne(m, l0)
}
func SparsevecNegativeInnerProduct(m *base.Module, l0 int32) int32 {
	return F_sparsevec_negative_inner_product(m, l0)
}
func SparsevecOut(m *base.Module, l0 int32) int32 {
	return F_sparsevec_out(m, l0)
}
func SparsevecRecv(m *base.Module, l0 int32) int32 {
	return F_sparsevec_recv(m, l0)
}
func SparsevecSend(m *base.Module, l0 int32) int32 {
	return F_sparsevec_send(m, l0)
}
func SparsevecToHalfvec(m *base.Module, l0 int32) int32 {
	return F_sparsevec_to_halfvec(m, l0)
}
func SparsevecToVector(m *base.Module, l0 int32) int32 {
	return F_sparsevec_to_vector(m, l0)
}
func SparsevecTypmodIn(m *base.Module, l0 int32) int32 {
	return F_sparsevec_typmod_in(m, l0)
}
func Subvector(m *base.Module, l0 int32) int32 {
	return F_subvector(m, l0)
}
func Vector(m *base.Module, l0 int32) int32 {
	return F_vector(m, l0)
}
func VectorAccum(m *base.Module, l0 int32) int32 {
	return F_vector_accum(m, l0)
}
func VectorAdd(m *base.Module, l0 int32) int32 {
	return F_vector_add(m, l0)
}
func VectorAvg(m *base.Module, l0 int32) int32 {
	return F_vector_avg(m, l0)
}
func VectorCmp(m *base.Module, l0 int32) int32 {
	return F_vector_cmp(m, l0)
}
func VectorCombine(m *base.Module, l0 int32) int32 {
	return F_vector_combine(m, l0)
}
func VectorConcat(m *base.Module, l0 int32) int32 {
	return F_vector_concat(m, l0)
}
func VectorDims(m *base.Module, l0 int32) int32 {
	return F_halfvec_vector_dims(m, l0)
}
func VectorEq(m *base.Module, l0 int32) int32 {
	return F_vector_eq(m, l0)
}
func VectorGe(m *base.Module, l0 int32) int32 {
	return F_vector_ge(m, l0)
}
func VectorGt(m *base.Module, l0 int32) int32 {
	return F_vector_gt(m, l0)
}
func VectorIn(m *base.Module, l0 int32) int32 {
	return F_vector_in(m, l0)
}
func VectorL2SquaredDistance(m *base.Module, l0 int32) int32 {
	return F_vector_l2_squared_distance(m, l0)
}
func VectorLe(m *base.Module, l0 int32) int32 {
	return F_vector_le(m, l0)
}
func VectorLt(m *base.Module, l0 int32) int32 {
	return F_vector_lt(m, l0)
}
func VectorMul(m *base.Module, l0 int32) int32 {
	return F_vector_mul(m, l0)
}
func VectorNe(m *base.Module, l0 int32) int32 {
	return F_vector_ne(m, l0)
}
func VectorNegativeInnerProduct(m *base.Module, l0 int32) int32 {
	return F_vector_negative_inner_product(m, l0)
}
func VectorNorm(m *base.Module, l0 int32) int32 {
	return F_vector_norm(m, l0)
}
func VectorOut(m *base.Module, l0 int32) int32 {
	return F_vector_out(m, l0)
}
func VectorRecv(m *base.Module, l0 int32) int32 {
	return F_vector_recv(m, l0)
}
func VectorSend(m *base.Module, l0 int32) int32 {
	return F_vector_send(m, l0)
}
func VectorSphericalDistance(m *base.Module, l0 int32) int32 {
	return F_vector_spherical_distance(m, l0)
}
func VectorSub(m *base.Module, l0 int32) int32 {
	return F_vector_sub(m, l0)
}
func VectorToFloat4(m *base.Module, l0 int32) int32 {
	return F_vector_to_float4(m, l0)
}
func VectorToHalfvec(m *base.Module, l0 int32) int32 {
	return F_vector_to_halfvec(m, l0)
}
func VectorToSparsevec(m *base.Module, l0 int32) int32 {
	return F_vector_to_sparsevec(m, l0)
}
func VectorTypmodIn(m *base.Module, l0 int32) int32 {
	return F_vector_typmod_in(m, l0)
}
func EmscriptenMemcpyBulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F__emscripten_memcpy_bulkmem(m, l0, l1, l2)
}
func EmscriptenMemsetBulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return F__emscripten_memset_bulkmem(m, l0, l1, l2)
}
func FuncsOnExit(m *base.Module) {
	F___funcs_on_exit(m)
}
func EmscriptenBuiltinFree(m *base.Module, l0 int32) {
	F_emscripten_builtin_free(m, l0)
}
func EmscriptenBuiltinMemalign(m *base.Module, l0 int32, l1 int32) int32 {
	return F_emscripten_builtin_memalign(m, l0, l1)
}
func Ntohs(m *base.Module, l0 int32) int32 {
	return F_htons(m, l0)
}
func EmscriptenTimeout(m *base.Module, l0 int32, l1 float64) {
	F__emscripten_timeout(m, l0, l1)
}
func EmscriptenStackRestore(m *base.Module, l0 int32) {
	F__emscripten_stack_restore(m, l0)
}
func EmscriptenStackAlloc(m *base.Module, l0 int32) int32 {
	return F__emscripten_stack_alloc(m, l0)
}
func EmscriptenStackGetCurrent(m *base.Module) int32 {
	return F_emscripten_stack_get_current(m)
}
func Memory(m *base.Module) []byte {
	return m.Memory
}
//go:embed data.bin
var wasm2goData_data_bin []byte
