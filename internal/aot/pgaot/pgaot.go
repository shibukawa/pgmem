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
	m.T0 = make([]any, 5454)
	m.G0 = int32(13007856)
	InitElemSeg_0_0(m)
	InitElemSeg_0_1(m)
	InitElemSeg_0_2(m)
	InitElemSeg_0_3(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_3_2(m)
	InitElemSeg_3_3(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	InitElemSeg_5_2(m)
	InitElemSeg_5_3(m)
	m.DataEnd = 4336955
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
	m.T0 = make([]any, 5454)
	m.G0 = int32(13007856)
	InitElemSeg_0_0(m)
	InitElemSeg_0_1(m)
	InitElemSeg_0_2(m)
	InitElemSeg_0_3(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_3_2(m)
	InitElemSeg_3_3(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	InitElemSeg_5_2(m)
	InitElemSeg_5_3(m)
	m.DataEnd = 4336955
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
	m.T0 = make([]any, 5454)
	m.G0 = int32(13007856)
	InitElemSeg_0_0(m)
	InitElemSeg_0_1(m)
	InitElemSeg_0_2(m)
	InitElemSeg_0_3(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_3_2(m)
	InitElemSeg_3_3(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	InitElemSeg_5_2(m)
	InitElemSeg_5_3(m)
	m.DataEnd = 4336955
	base.RestoreGlobals(m, globals)
	return m
}
func initData_0(m *base.Module) {
	copy(m.Memory[4096:], wasm2goData_data_bin[0:1566784])
	copy(m.Memory[1572112:], wasm2goData_data_bin[1566784:1566793])
	copy(m.Memory[1574160:], wasm2goData_data_bin[1566793:1579721])
	copy(m.Memory[1597444:], wasm2goData_data_bin[1579721:1672962])
	copy(m.Memory[1691712:], wasm2goData_data_bin[1672962:1784618])
	copy(m.Memory[1814584:], wasm2goData_data_bin[1784618:4024176])
	copy(m.Memory[4061632:], wasm2goData_data_bin[4024176:4299499])
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
func PgMagicFuncUtf8AndIso8859(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_iso8859_1(m)
}
func Iso8859ToUtf8(m *base.Module, l0 int32) int32 {
	return F_iso8859_to_utf8(m, l0)
}
func PgFinfoIso8859ToUtf8(m *base.Module) int32 {
	return F_pg_finfo_iso8859_to_utf8(m)
}
func PgFinfoUtf8ToIso8859(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_iso8859_1(m)
}
func Utf8ToIso8859(m *base.Module, l0 int32) int32 {
	return F_utf8_to_iso8859_1(m, l0)
}
func PgMagicFuncUtf8AndIso8859_1(m *base.Module) int32 {
	return F_Pg_magic_func_utf8_and_iso8859_2(m)
}
func Iso8859_1_toUtf8(m *base.Module, l0 int32) int32 {
	return F_iso8859_1_to_utf8(m, l0)
}
func PgFinfoIso8859_1_toUtf8(m *base.Module) int32 {
	return F_pg_finfo_iso8859_1_to_utf8(m)
}
func PgFinfoUtf8ToIso8859_1(m *base.Module) int32 {
	return F_pg_finfo_utf8_to_iso8859_2(m)
}
func Utf8ToIso8859_1(m *base.Module, l0 int32) int32 {
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
