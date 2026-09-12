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
	m.T0 = make([]any, 5886)
	m.G0 = int32(13016400)
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
	InitElemSeg_3_4(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_4_4(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	InitElemSeg_5_2(m)
	InitElemSeg_5_3(m)
	m.DataEnd = 4345499
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
	m.T0 = make([]any, 5886)
	m.G0 = int32(13016400)
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
	InitElemSeg_3_4(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_4_4(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	InitElemSeg_5_2(m)
	InitElemSeg_5_3(m)
	m.DataEnd = 4345499
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
	m.T0 = make([]any, 5886)
	m.G0 = int32(13016400)
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
	InitElemSeg_3_4(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_4_4(m)
	InitElemSeg_5_0(m)
	InitElemSeg_5_1(m)
	InitElemSeg_5_2(m)
	InitElemSeg_5_3(m)
	m.DataEnd = 4345499
	base.RestoreGlobals(m, globals)
	return m
}
func initData_0(m *base.Module) {
	copy(m.Memory[4096:], wasm2goData_data_bin[0:1572496])
	copy(m.Memory[1577824:], wasm2goData_data_bin[1572496:1572505])
	copy(m.Memory[1579872:], wasm2goData_data_bin[1572505:1585433])
	copy(m.Memory[1601540:], wasm2goData_data_bin[1585433:1678674])
	copy(m.Memory[1695808:], wasm2goData_data_bin[1678674:1790330])
	copy(m.Memory[1818680:], wasm2goData_data_bin[1790330:4034032])
	copy(m.Memory[4069872:], wasm2goData_data_bin[4034032:4309659])
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
