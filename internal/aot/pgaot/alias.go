package pgaot

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	_ "unsafe"
)
//go:linkname F___wasm_call_ctors github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___wasm_call_ctors
func F___wasm_call_ctors(m *base.Module)
//go:linkname F_pq_buffer_remaining_data github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_buffer_remaining_data
func F_pq_buffer_remaining_data(m *base.Module) int32
//go:linkname F_ProcessStartupPacket github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessStartupPacket
func F_ProcessStartupPacket(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgl_startPGlite github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgl_startPGlite
func F_pgl_startPGlite(m *base.Module)
//go:linkname F_pgl_pq_flush github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgl_pq_flush
func F_pgl_pq_flush(m *base.Module)
//go:linkname F_pgl_getMyProcPort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgl_getMyProcPort
func F_pgl_getMyProcPort(m *base.Module) int32
//go:linkname F_pgl_sendConnData github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgl_sendConnData
func F_pgl_sendConnData(m *base.Module)
//go:linkname F_PostgresMainLongJmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PostgresMainLongJmp
func F_PostgresMainLongJmp(m *base.Module)
//go:linkname F_PostgresMainLoopOnce github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PostgresMainLoopOnce
func F_PostgresMainLoopOnce(m *base.Module)
//go:linkname F_PostgresSendReadyForQueryIfNecessary github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PostgresSendReadyForQueryIfNecessary
func F_PostgresSendReadyForQueryIfNecessary(m *base.Module)
//go:linkname F_pgmem_reset_session github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgmem_reset_session
func F_pgmem_reset_session(m *base.Module, l0 int32)
//go:linkname F_pgl_setPGliteActive github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_setPGliteActive
func F_pgl_setPGliteActive(m *base.Module, l0 int32) int32
//go:linkname F_pgl_setPGliteExitStatus github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_setPGliteExitStatus
func F_pgl_setPGliteExitStatus(m *base.Module, l0 int32) int32
//go:linkname F_pgl_run_atexit_funcs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_run_atexit_funcs
func F_pgl_run_atexit_funcs(m *base.Module)
//go:linkname F_pgl_freopen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgl_freopen
func F_pgl_freopen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgmem_poll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgmem_poll
func F_pgmem_poll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgmem_init github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgmem_init
func F_pgmem_init(m *base.Module)
//go:linkname F_pgmem_main github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgmem_main
func F_pgmem_main(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgmem_call_sighandler github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgmem_call_sighandler
func F_pgmem_call_sighandler(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgmem_module_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgmem_module_name
func F_pgmem_module_name(m *base.Module, l0 int32) int32
//go:linkname F_fflush github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F__emscripten_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_timeout
func F__emscripten_timeout(m *base.Module, l0 int32, l1 float64)
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F_emscripten_builtin_memalign github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_emscripten_builtin_memalign
func F_emscripten_builtin_memalign(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__emscripten_stack_restore github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_stack_restore
func F__emscripten_stack_restore(m *base.Module, l0 int32)
//go:linkname F_emscripten_stack_get_current github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_emscripten_stack_get_current
func F_emscripten_stack_get_current(m *base.Module) int32
//go:linkname InitElemSeg_0_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.InitElemSeg_0_0
func InitElemSeg_0_0(m *base.Module)
//go:linkname InitElemSeg_0_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.InitElemSeg_0_1
func InitElemSeg_0_1(m *base.Module)
//go:linkname InitElemSeg_0_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.InitElemSeg_0_2
func InitElemSeg_0_2(m *base.Module)
//go:linkname InitElemSeg_0_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.InitElemSeg_0_3
func InitElemSeg_0_3(m *base.Module)
//go:linkname InitElemSeg_0_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.InitElemSeg_0_4
func InitElemSeg_0_4(m *base.Module)
//go:linkname InitElemSeg_1_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.InitElemSeg_1_0
func InitElemSeg_1_0(m *base.Module)
//go:linkname InitElemSeg_1_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.InitElemSeg_1_1
func InitElemSeg_1_1(m *base.Module)
//go:linkname InitElemSeg_1_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.InitElemSeg_1_2
func InitElemSeg_1_2(m *base.Module)
//go:linkname InitElemSeg_1_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.InitElemSeg_1_3
func InitElemSeg_1_3(m *base.Module)
//go:linkname InitElemSeg_1_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.InitElemSeg_1_4
func InitElemSeg_1_4(m *base.Module)
//go:linkname InitElemSeg_2_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.InitElemSeg_2_0
func InitElemSeg_2_0(m *base.Module)
//go:linkname InitElemSeg_2_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.InitElemSeg_2_1
func InitElemSeg_2_1(m *base.Module)
//go:linkname InitElemSeg_2_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.InitElemSeg_2_2
func InitElemSeg_2_2(m *base.Module)
//go:linkname InitElemSeg_2_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.InitElemSeg_2_3
func InitElemSeg_2_3(m *base.Module)
//go:linkname InitElemSeg_2_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.InitElemSeg_2_4
func InitElemSeg_2_4(m *base.Module)
//go:linkname InitElemSeg_3_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.InitElemSeg_3_0
func InitElemSeg_3_0(m *base.Module)
//go:linkname InitElemSeg_3_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.InitElemSeg_3_1
func InitElemSeg_3_1(m *base.Module)
//go:linkname InitElemSeg_3_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.InitElemSeg_3_2
func InitElemSeg_3_2(m *base.Module)
//go:linkname InitElemSeg_3_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.InitElemSeg_3_3
func InitElemSeg_3_3(m *base.Module)
//go:linkname InitElemSeg_3_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.InitElemSeg_3_4
func InitElemSeg_3_4(m *base.Module)
//go:linkname InitElemSeg_3_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.InitElemSeg_3_5
func InitElemSeg_3_5(m *base.Module)
//go:linkname InitElemSeg_4_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_0
func InitElemSeg_4_0(m *base.Module)
//go:linkname InitElemSeg_4_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_1
func InitElemSeg_4_1(m *base.Module)
//go:linkname InitElemSeg_4_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_2
func InitElemSeg_4_2(m *base.Module)
//go:linkname InitElemSeg_4_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_3
func InitElemSeg_4_3(m *base.Module)
//go:linkname InitElemSeg_4_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_4
func InitElemSeg_4_4(m *base.Module)
//go:linkname InitElemSeg_4_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_5
func InitElemSeg_4_5(m *base.Module)
//go:linkname InitElemSeg_5_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_0
func InitElemSeg_5_0(m *base.Module)
//go:linkname InitElemSeg_5_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_1
func InitElemSeg_5_1(m *base.Module)
//go:linkname InitElemSeg_5_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_2
func InitElemSeg_5_2(m *base.Module)
//go:linkname InitElemSeg_5_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_3
func InitElemSeg_5_3(m *base.Module)
//go:linkname InitElemSeg_5_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_4
func InitElemSeg_5_4(m *base.Module)
