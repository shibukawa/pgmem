package pgaot

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	_ "unsafe"
)
//go:linkname F___wasm_call_ctors github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___wasm_call_ctors
func F___wasm_call_ctors(m *base.Module)
//go:linkname F_gist_bbox_zorder_abbrev_abort github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gist_bbox_zorder_abbrev_abort
func F_gist_bbox_zorder_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_IsTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IsTransactionBlock
func F_IsTransactionBlock(m *base.Module) int32
//go:linkname F_pq_buffer_remaining_data github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_buffer_remaining_data
func F_pq_buffer_remaining_data(m *base.Module) int32
//go:linkname F_main github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_main
func F_main(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_numa_available github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_numa_available
func F_pg_numa_available(m *base.Module, l0 int32) int32
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
//go:linkname F_float4up github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_float4up
func F_float4up(m *base.Module, l0 int32) int32
//go:linkname F_gtsvector_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gtsvector_decompress
func F_gtsvector_decompress(m *base.Module, l0 int32) int32
//go:linkname F_comparetup_index_hash_tiebreak github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_comparetup_index_hash_tiebreak
func F_comparetup_index_hash_tiebreak(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_clear_setitimer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_clear_setitimer
func F_clear_setitimer(m *base.Module)
//go:linkname F_pgl_setPGliteActive github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_setPGliteActive
func F_pgl_setPGliteActive(m *base.Module, l0 int32) int32
//go:linkname F_pgl_getPGliteExitStatus github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgl_getPGliteExitStatus
func F_pgl_getPGliteExitStatus(m *base.Module) int32
//go:linkname F_pgl_setPGliteExitStatus github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_setPGliteExitStatus
func F_pgl_setPGliteExitStatus(m *base.Module, l0 int32) int32
//go:linkname F_pgl_longjmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_longjmp
func F_pgl_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgl_set_system_fn github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_set_system_fn
func F_pgl_set_system_fn(m *base.Module, l0 int32)
//go:linkname F_pgl_system github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_system
func F_pgl_system(m *base.Module, l0 int32) int32
//go:linkname F_pgl_set_popen_fn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_set_popen_fn
func F_pgl_set_popen_fn(m *base.Module, l0 int32)
//go:linkname F_pgl_popen github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_popen
func F_pgl_popen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgl_set_pclose_fn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_set_pclose_fn
func F_pgl_set_pclose_fn(m *base.Module, l0 int32)
//go:linkname F_pgl_pclose github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgl_pclose
func F_pgl_pclose(m *base.Module, l0 int32) int32
//go:linkname F_pgl_geteuid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_geteuid
func F_pgl_geteuid(m *base.Module) int32
//go:linkname F_pgl_getpwuid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_getpwuid
func F_pgl_getpwuid(m *base.Module, l0 int32) int32
//go:linkname F_pgl_atexit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgl_atexit
func F_pgl_atexit(m *base.Module, l0 int32) int32
//go:linkname F_pgl_run_atexit_funcs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_run_atexit_funcs
func F_pgl_run_atexit_funcs(m *base.Module)
//go:linkname F_pgl_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_exit
func F_pgl_exit(m *base.Module, l0 int32)
//go:linkname F_pgl_freopen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgl_freopen
func F_pgl_freopen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgl_shmget github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_shmget
func F_pgl_shmget(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgl_shmat github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgl_shmat
func F_pgl_shmat(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgl_shmdt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgl_shmdt
func F_pgl_shmdt(m *base.Module, l0 int32) int32
//go:linkname F_pgl_shmctl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_shmctl
func F_pgl_shmctl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgl_set_rw_cbs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgl_set_rw_cbs
func F_pgl_set_rw_cbs(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgl_setsockopt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_setsockopt
func F_pgl_setsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pgl_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgl_recv
func F_pgl_recv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgl_send github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_send
func F_pgl_send(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgl_poll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_poll
func F_pgl_poll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
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
//go:linkname F_plpgsql_compile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_compile
func F_plpgsql_compile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_build_datatype github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plpgsql_build_datatype
func F_plpgsql_build_datatype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_plpgsql_parser_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_parser_setup
func F_plpgsql_parser_setup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_plpgsql_recognize_err_condition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_recognize_err_condition
func F_plpgsql_recognize_err_condition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_exec_get_datum_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plpgsql_exec_get_datum_type
func F_plpgsql_exec_get_datum_type(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_ns_lookup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_ns_lookup
func F_plpgsql_ns_lookup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_plpgsql_stmt_typename github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_stmt_typename
func F_plpgsql_stmt_typename(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_plpgsql github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_plpgsql
func F_Pg_magic_func_plpgsql(m *base.Module) int32
//go:linkname F__PG_init_plpgsql github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__PG_init_plpgsql
func F__PG_init_plpgsql(m *base.Module)
//go:linkname F_pg_finfo_plpgsql_call_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_plpgsql_call_handler
func F_pg_finfo_plpgsql_call_handler(m *base.Module) int32
//go:linkname F_plpgsql_call_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plpgsql_call_handler
func F_plpgsql_call_handler(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_plpgsql_inline_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_plpgsql_inline_handler
func F_pg_finfo_plpgsql_inline_handler(m *base.Module) int32
//go:linkname F_plpgsql_inline_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plpgsql_inline_handler
func F_plpgsql_inline_handler(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_plpgsql_validator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_plpgsql_validator
func F_pg_finfo_plpgsql_validator(m *base.Module) int32
//go:linkname F_plpgsql_validator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plpgsql_validator
func F_plpgsql_validator(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_dict_snowball github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_dict_snowball
func F_Pg_magic_func_dict_snowball(m *base.Module) int32
//go:linkname F_pg_finfo_dsnowball_init github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_dsnowball_init
func F_pg_finfo_dsnowball_init(m *base.Module) int32
//go:linkname F_pg_finfo_dsnowball_lexize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_dsnowball_lexize
func F_pg_finfo_dsnowball_lexize(m *base.Module) int32
//go:linkname F_dsnowball_init github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dsnowball_init
func F_dsnowball_init(m *base.Module, l0 int32) int32
//go:linkname F_dsnowball_lexize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dsnowball_lexize
func F_dsnowball_lexize(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_cyrillic_and_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_cyrillic_and_mic
func F_Pg_magic_func_cyrillic_and_mic(m *base.Module) int32
//go:linkname F_pg_finfo_koi8r_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_koi8r_to_mic
func F_pg_finfo_koi8r_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_mic_to_koi8r
func F_pg_finfo_mic_to_koi8r(m *base.Module) int32
//go:linkname F_pg_finfo_iso_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_iso_to_mic
func F_pg_finfo_iso_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_iso github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_mic_to_iso
func F_pg_finfo_mic_to_iso(m *base.Module) int32
//go:linkname F_pg_finfo_win1251_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_win1251_to_mic
func F_pg_finfo_win1251_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_win1251 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_mic_to_win1251
func F_pg_finfo_mic_to_win1251(m *base.Module) int32
//go:linkname F_pg_finfo_win866_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_win866_to_mic
func F_pg_finfo_win866_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_win866 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_mic_to_win866
func F_pg_finfo_mic_to_win866(m *base.Module) int32
//go:linkname F_pg_finfo_koi8r_to_win1251 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_koi8r_to_win1251
func F_pg_finfo_koi8r_to_win1251(m *base.Module) int32
//go:linkname F_pg_finfo_win1251_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_win1251_to_koi8r
func F_pg_finfo_win1251_to_koi8r(m *base.Module) int32
//go:linkname F_pg_finfo_koi8r_to_win866 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_koi8r_to_win866
func F_pg_finfo_koi8r_to_win866(m *base.Module) int32
//go:linkname F_pg_finfo_win866_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_win866_to_koi8r
func F_pg_finfo_win866_to_koi8r(m *base.Module) int32
//go:linkname F_pg_finfo_win866_to_win1251 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_win866_to_win1251
func F_pg_finfo_win866_to_win1251(m *base.Module) int32
//go:linkname F_pg_finfo_win1251_to_win866 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_win1251_to_win866
func F_pg_finfo_win1251_to_win866(m *base.Module) int32
//go:linkname F_pg_finfo_iso_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_iso_to_koi8r
func F_pg_finfo_iso_to_koi8r(m *base.Module) int32
//go:linkname F_pg_finfo_koi8r_to_iso github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_koi8r_to_iso
func F_pg_finfo_koi8r_to_iso(m *base.Module) int32
//go:linkname F_pg_finfo_iso_to_win1251 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_iso_to_win1251
func F_pg_finfo_iso_to_win1251(m *base.Module) int32
//go:linkname F_pg_finfo_win1251_to_iso github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_win1251_to_iso
func F_pg_finfo_win1251_to_iso(m *base.Module) int32
//go:linkname F_pg_finfo_iso_to_win866 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_iso_to_win866
func F_pg_finfo_iso_to_win866(m *base.Module) int32
//go:linkname F_pg_finfo_win866_to_iso github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_win866_to_iso
func F_pg_finfo_win866_to_iso(m *base.Module) int32
//go:linkname F_koi8r_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_koi8r_to_mic
func F_koi8r_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mic_to_koi8r
func F_mic_to_koi8r(m *base.Module, l0 int32) int32
//go:linkname F_iso_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_iso_to_mic
func F_iso_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_iso github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mic_to_iso
func F_mic_to_iso(m *base.Module, l0 int32) int32
//go:linkname F_win1251_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_win1251_to_mic
func F_win1251_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_win1251 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mic_to_win1251
func F_mic_to_win1251(m *base.Module, l0 int32) int32
//go:linkname F_win866_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_win866_to_mic
func F_win866_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_win866 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_mic_to_win866
func F_mic_to_win866(m *base.Module, l0 int32) int32
//go:linkname F_koi8r_to_win1251 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_koi8r_to_win1251
func F_koi8r_to_win1251(m *base.Module, l0 int32) int32
//go:linkname F_win1251_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_win1251_to_koi8r
func F_win1251_to_koi8r(m *base.Module, l0 int32) int32
//go:linkname F_koi8r_to_win866 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_koi8r_to_win866
func F_koi8r_to_win866(m *base.Module, l0 int32) int32
//go:linkname F_win866_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_win866_to_koi8r
func F_win866_to_koi8r(m *base.Module, l0 int32) int32
//go:linkname F_win866_to_win1251 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_win866_to_win1251
func F_win866_to_win1251(m *base.Module, l0 int32) int32
//go:linkname F_win1251_to_win866 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_win1251_to_win866
func F_win1251_to_win866(m *base.Module, l0 int32) int32
//go:linkname F_iso_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_iso_to_koi8r
func F_iso_to_koi8r(m *base.Module, l0 int32) int32
//go:linkname F_koi8r_to_iso github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_koi8r_to_iso
func F_koi8r_to_iso(m *base.Module, l0 int32) int32
//go:linkname F_iso_to_win1251 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_iso_to_win1251
func F_iso_to_win1251(m *base.Module, l0 int32) int32
//go:linkname F_win1251_to_iso github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_win1251_to_iso
func F_win1251_to_iso(m *base.Module, l0 int32) int32
//go:linkname F_iso_to_win866 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_iso_to_win866
func F_iso_to_win866(m *base.Module, l0 int32) int32
//go:linkname F_win866_to_iso github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_win866_to_iso
func F_win866_to_iso(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_euc_cn_and_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_euc_cn_and_mic
func F_Pg_magic_func_euc_cn_and_mic(m *base.Module) int32
//go:linkname F_pg_finfo_euc_cn_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_euc_cn_to_mic
func F_pg_finfo_euc_cn_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_euc_cn github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_mic_to_euc_cn
func F_pg_finfo_mic_to_euc_cn(m *base.Module) int32
//go:linkname F_euc_cn_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_euc_cn_to_mic
func F_euc_cn_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_euc_cn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mic_to_euc_cn
func F_mic_to_euc_cn(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_euc_jp_and_sjis github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_euc_jp_and_sjis
func F_Pg_magic_func_euc_jp_and_sjis(m *base.Module) int32
//go:linkname F_pg_finfo_euc_jp_to_sjis github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_euc_jp_to_sjis
func F_pg_finfo_euc_jp_to_sjis(m *base.Module) int32
//go:linkname F_pg_finfo_sjis_to_euc_jp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_sjis_to_euc_jp
func F_pg_finfo_sjis_to_euc_jp(m *base.Module) int32
//go:linkname F_pg_finfo_euc_jp_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_euc_jp_to_mic
func F_pg_finfo_euc_jp_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_euc_jp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_mic_to_euc_jp
func F_pg_finfo_mic_to_euc_jp(m *base.Module) int32
//go:linkname F_pg_finfo_sjis_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_sjis_to_mic
func F_pg_finfo_sjis_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_sjis github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_mic_to_sjis
func F_pg_finfo_mic_to_sjis(m *base.Module) int32
//go:linkname F_euc_jp_to_sjis github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_euc_jp_to_sjis
func F_euc_jp_to_sjis(m *base.Module, l0 int32) int32
//go:linkname F_sjis_to_euc_jp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sjis_to_euc_jp
func F_sjis_to_euc_jp(m *base.Module, l0 int32) int32
//go:linkname F_euc_jp_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_euc_jp_to_mic
func F_euc_jp_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_euc_jp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_mic_to_euc_jp
func F_mic_to_euc_jp(m *base.Module, l0 int32) int32
//go:linkname F_sjis_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_sjis_to_mic
func F_sjis_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_sjis github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mic_to_sjis
func F_mic_to_sjis(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_euc_kr_and_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_euc_kr_and_mic
func F_Pg_magic_func_euc_kr_and_mic(m *base.Module) int32
//go:linkname F_pg_finfo_euc_kr_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_euc_kr_to_mic
func F_pg_finfo_euc_kr_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_euc_kr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_mic_to_euc_kr
func F_pg_finfo_mic_to_euc_kr(m *base.Module) int32
//go:linkname F_euc_kr_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_euc_kr_to_mic
func F_euc_kr_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_euc_kr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mic_to_euc_kr
func F_mic_to_euc_kr(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_euc_tw_and_big5 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_euc_tw_and_big5
func F_Pg_magic_func_euc_tw_and_big5(m *base.Module) int32
//go:linkname F_pg_finfo_euc_tw_to_big5 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_euc_tw_to_big5
func F_pg_finfo_euc_tw_to_big5(m *base.Module) int32
//go:linkname F_pg_finfo_big5_to_euc_tw github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_big5_to_euc_tw
func F_pg_finfo_big5_to_euc_tw(m *base.Module) int32
//go:linkname F_pg_finfo_euc_tw_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_euc_tw_to_mic
func F_pg_finfo_euc_tw_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_euc_tw github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_mic_to_euc_tw
func F_pg_finfo_mic_to_euc_tw(m *base.Module) int32
//go:linkname F_pg_finfo_big5_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_big5_to_mic
func F_pg_finfo_big5_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_big5 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_mic_to_big5
func F_pg_finfo_mic_to_big5(m *base.Module) int32
//go:linkname F_euc_tw_to_big5 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_euc_tw_to_big5
func F_euc_tw_to_big5(m *base.Module, l0 int32) int32
//go:linkname F_big5_to_euc_tw github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_big5_to_euc_tw
func F_big5_to_euc_tw(m *base.Module, l0 int32) int32
//go:linkname F_euc_tw_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_euc_tw_to_mic
func F_euc_tw_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_euc_tw github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mic_to_euc_tw
func F_mic_to_euc_tw(m *base.Module, l0 int32) int32
//go:linkname F_big5_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_big5_to_mic
func F_big5_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_big5 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mic_to_big5
func F_mic_to_big5(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_euc2004_sjis2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_euc2004_sjis2004
func F_Pg_magic_func_euc2004_sjis2004(m *base.Module) int32
//go:linkname F_pg_finfo_euc_jis_2004_to_shift_jis_2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_euc_jis_2004_to_shift_jis_2004
func F_pg_finfo_euc_jis_2004_to_shift_jis_2004(m *base.Module) int32
//go:linkname F_pg_finfo_shift_jis_2004_to_euc_jis_2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_shift_jis_2004_to_euc_jis_2004
func F_pg_finfo_shift_jis_2004_to_euc_jis_2004(m *base.Module) int32
//go:linkname F_euc_jis_2004_to_shift_jis_2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_euc_jis_2004_to_shift_jis_2004
func F_euc_jis_2004_to_shift_jis_2004(m *base.Module, l0 int32) int32
//go:linkname F_shift_jis_2004_to_euc_jis_2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shift_jis_2004_to_euc_jis_2004
func F_shift_jis_2004_to_euc_jis_2004(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_latin_and_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Pg_magic_func_latin_and_mic
func F_Pg_magic_func_latin_and_mic(m *base.Module) int32
//go:linkname F_pg_finfo_latin1_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_latin1_to_mic
func F_pg_finfo_latin1_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_latin1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_mic_to_latin1
func F_pg_finfo_mic_to_latin1(m *base.Module) int32
//go:linkname F_pg_finfo_latin3_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_latin3_to_mic
func F_pg_finfo_latin3_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_latin3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_mic_to_latin3
func F_pg_finfo_mic_to_latin3(m *base.Module) int32
//go:linkname F_pg_finfo_latin4_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_latin4_to_mic
func F_pg_finfo_latin4_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_latin4 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_mic_to_latin4
func F_pg_finfo_mic_to_latin4(m *base.Module) int32
//go:linkname F_latin1_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_latin1_to_mic
func F_latin1_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_latin1 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mic_to_latin1
func F_mic_to_latin1(m *base.Module, l0 int32) int32
//go:linkname F_latin3_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_latin3_to_mic
func F_latin3_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_latin3 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mic_to_latin3
func F_mic_to_latin3(m *base.Module, l0 int32) int32
//go:linkname F_latin4_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_latin4_to_mic
func F_latin4_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_latin4 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mic_to_latin4
func F_mic_to_latin4(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_latin2_and_win1250 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_latin2_and_win1250
func F_Pg_magic_func_latin2_and_win1250(m *base.Module) int32
//go:linkname F_pg_finfo_latin2_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_latin2_to_mic
func F_pg_finfo_latin2_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_latin2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_mic_to_latin2
func F_pg_finfo_mic_to_latin2(m *base.Module) int32
//go:linkname F_pg_finfo_win1250_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_win1250_to_mic
func F_pg_finfo_win1250_to_mic(m *base.Module) int32
//go:linkname F_pg_finfo_mic_to_win1250 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_mic_to_win1250
func F_pg_finfo_mic_to_win1250(m *base.Module) int32
//go:linkname F_pg_finfo_latin2_to_win1250 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_latin2_to_win1250
func F_pg_finfo_latin2_to_win1250(m *base.Module) int32
//go:linkname F_pg_finfo_win1250_to_latin2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_win1250_to_latin2
func F_pg_finfo_win1250_to_latin2(m *base.Module) int32
//go:linkname F_latin2_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_latin2_to_mic
func F_latin2_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_latin2 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mic_to_latin2
func F_mic_to_latin2(m *base.Module, l0 int32) int32
//go:linkname F_win1250_to_mic github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_win1250_to_mic
func F_win1250_to_mic(m *base.Module, l0 int32) int32
//go:linkname F_mic_to_win1250 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mic_to_win1250
func F_mic_to_win1250(m *base.Module, l0 int32) int32
//go:linkname F_latin2_to_win1250 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_latin2_to_win1250
func F_latin2_to_win1250(m *base.Module, l0 int32) int32
//go:linkname F_win1250_to_latin2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_win1250_to_latin2
func F_win1250_to_latin2(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_big5 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_utf8_and_big5
func F_Pg_magic_func_utf8_and_big5(m *base.Module) int32
//go:linkname F_pg_finfo_big5_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_big5_to_utf8
func F_pg_finfo_big5_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_big5 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_utf8_to_big5
func F_pg_finfo_utf8_to_big5(m *base.Module) int32
//go:linkname F_big5_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_big5_to_utf8
func F_big5_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_big5 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_utf8_to_big5
func F_utf8_to_big5(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_cyrillic github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_utf8_and_cyrillic
func F_Pg_magic_func_utf8_and_cyrillic(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_utf8_to_koi8r
func F_pg_finfo_utf8_to_koi8r(m *base.Module) int32
//go:linkname F_pg_finfo_koi8r_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_koi8r_to_utf8
func F_pg_finfo_koi8r_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_koi8u github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_utf8_to_koi8u
func F_pg_finfo_utf8_to_koi8u(m *base.Module) int32
//go:linkname F_pg_finfo_koi8u_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_koi8u_to_utf8
func F_pg_finfo_koi8u_to_utf8(m *base.Module) int32
//go:linkname F_utf8_to_koi8r github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_utf8_to_koi8r
func F_utf8_to_koi8r(m *base.Module, l0 int32) int32
//go:linkname F_koi8r_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_koi8r_to_utf8
func F_koi8r_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_koi8u github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_utf8_to_koi8u
func F_utf8_to_koi8u(m *base.Module, l0 int32) int32
//go:linkname F_koi8u_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_koi8u_to_utf8
func F_koi8u_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_euc_cn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_utf8_and_euc_cn
func F_Pg_magic_func_utf8_and_euc_cn(m *base.Module) int32
//go:linkname F_pg_finfo_euc_cn_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_euc_cn_to_utf8
func F_pg_finfo_euc_cn_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_euc_cn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_utf8_to_euc_cn
func F_pg_finfo_utf8_to_euc_cn(m *base.Module) int32
//go:linkname F_euc_cn_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_euc_cn_to_utf8
func F_euc_cn_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_euc_cn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_utf8_to_euc_cn
func F_utf8_to_euc_cn(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_euc_jp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_Pg_magic_func_utf8_and_euc_jp
func F_Pg_magic_func_utf8_and_euc_jp(m *base.Module) int32
//go:linkname F_pg_finfo_euc_jp_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_euc_jp_to_utf8
func F_pg_finfo_euc_jp_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_euc_jp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_utf8_to_euc_jp
func F_pg_finfo_utf8_to_euc_jp(m *base.Module) int32
//go:linkname F_euc_jp_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_euc_jp_to_utf8
func F_euc_jp_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_euc_jp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_utf8_to_euc_jp
func F_utf8_to_euc_jp(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_euc_kr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_utf8_and_euc_kr
func F_Pg_magic_func_utf8_and_euc_kr(m *base.Module) int32
//go:linkname F_pg_finfo_euc_kr_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_euc_kr_to_utf8
func F_pg_finfo_euc_kr_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_euc_kr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_utf8_to_euc_kr
func F_pg_finfo_utf8_to_euc_kr(m *base.Module) int32
//go:linkname F_euc_kr_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_euc_kr_to_utf8
func F_euc_kr_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_euc_kr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_utf8_to_euc_kr
func F_utf8_to_euc_kr(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_euc_tw github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_utf8_and_euc_tw
func F_Pg_magic_func_utf8_and_euc_tw(m *base.Module) int32
//go:linkname F_pg_finfo_euc_tw_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_euc_tw_to_utf8
func F_pg_finfo_euc_tw_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_euc_tw github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_utf8_to_euc_tw
func F_pg_finfo_utf8_to_euc_tw(m *base.Module) int32
//go:linkname F_euc_tw_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_euc_tw_to_utf8
func F_euc_tw_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_euc_tw github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_utf8_to_euc_tw
func F_utf8_to_euc_tw(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_euc2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_utf8_and_euc2004
func F_Pg_magic_func_utf8_and_euc2004(m *base.Module) int32
//go:linkname F_pg_finfo_euc_jis_2004_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_euc_jis_2004_to_utf8
func F_pg_finfo_euc_jis_2004_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_euc_jis_2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_utf8_to_euc_jis_2004
func F_pg_finfo_utf8_to_euc_jis_2004(m *base.Module) int32
//go:linkname F_euc_jis_2004_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_euc_jis_2004_to_utf8
func F_euc_jis_2004_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_euc_jis_2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_utf8_to_euc_jis_2004
func F_utf8_to_euc_jis_2004(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_gb18030 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_utf8_and_gb18030
func F_Pg_magic_func_utf8_and_gb18030(m *base.Module) int32
//go:linkname F_pg_finfo_gb18030_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gb18030_to_utf8
func F_pg_finfo_gb18030_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_gb18030 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_utf8_to_gb18030
func F_pg_finfo_utf8_to_gb18030(m *base.Module) int32
//go:linkname F_gb18030_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gb18030_to_utf8
func F_gb18030_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_gb18030 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_utf8_to_gb18030
func F_utf8_to_gb18030(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_gbk github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_utf8_and_gbk
func F_Pg_magic_func_utf8_and_gbk(m *base.Module) int32
//go:linkname F_pg_finfo_gbk_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbk_to_utf8
func F_pg_finfo_gbk_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_gbk github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_utf8_to_gbk
func F_pg_finfo_utf8_to_gbk(m *base.Module) int32
//go:linkname F_gbk_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbk_to_utf8
func F_gbk_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_gbk github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_utf8_to_gbk
func F_utf8_to_gbk(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_iso8859_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_utf8_and_iso8859_1
func F_Pg_magic_func_utf8_and_iso8859_1(m *base.Module) int32
//go:linkname F_pg_finfo_iso8859_1_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_iso8859_1_to_utf8
func F_pg_finfo_iso8859_1_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_iso8859_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_utf8_to_iso8859_1
func F_pg_finfo_utf8_to_iso8859_1(m *base.Module) int32
//go:linkname F_iso8859_1_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_iso8859_1_to_utf8
func F_iso8859_1_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_iso8859_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_utf8_to_iso8859_1
func F_utf8_to_iso8859_1(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_iso8859_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_Pg_magic_func_utf8_and_iso8859_2
func F_Pg_magic_func_utf8_and_iso8859_2(m *base.Module) int32
//go:linkname F_pg_finfo_iso8859_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_iso8859_to_utf8
func F_pg_finfo_iso8859_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_iso8859_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_utf8_to_iso8859_2
func F_pg_finfo_utf8_to_iso8859_2(m *base.Module) int32
//go:linkname F_iso8859_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_iso8859_to_utf8
func F_iso8859_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_iso8859_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_utf8_to_iso8859_2
func F_utf8_to_iso8859_2(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_johab github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Pg_magic_func_utf8_and_johab
func F_Pg_magic_func_utf8_and_johab(m *base.Module) int32
//go:linkname F_pg_finfo_johab_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_johab_to_utf8
func F_pg_finfo_johab_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_johab github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_utf8_to_johab
func F_pg_finfo_utf8_to_johab(m *base.Module) int32
//go:linkname F_johab_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_johab_to_utf8
func F_johab_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_johab github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_utf8_to_johab
func F_utf8_to_johab(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_sjis github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_utf8_and_sjis
func F_Pg_magic_func_utf8_and_sjis(m *base.Module) int32
//go:linkname F_pg_finfo_sjis_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_sjis_to_utf8
func F_pg_finfo_sjis_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_sjis github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_utf8_to_sjis
func F_pg_finfo_utf8_to_sjis(m *base.Module) int32
//go:linkname F_sjis_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sjis_to_utf8
func F_sjis_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_sjis github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_utf8_to_sjis
func F_utf8_to_sjis(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_sjis2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_utf8_and_sjis2004
func F_Pg_magic_func_utf8_and_sjis2004(m *base.Module) int32
//go:linkname F_pg_finfo_shift_jis_2004_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_shift_jis_2004_to_utf8
func F_pg_finfo_shift_jis_2004_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_shift_jis_2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_utf8_to_shift_jis_2004
func F_pg_finfo_utf8_to_shift_jis_2004(m *base.Module) int32
//go:linkname F_shift_jis_2004_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_shift_jis_2004_to_utf8
func F_shift_jis_2004_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_shift_jis_2004 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_utf8_to_shift_jis_2004
func F_utf8_to_shift_jis_2004(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_uhc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_utf8_and_uhc
func F_Pg_magic_func_utf8_and_uhc(m *base.Module) int32
//go:linkname F_pg_finfo_uhc_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_uhc_to_utf8
func F_pg_finfo_uhc_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_uhc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_utf8_to_uhc
func F_pg_finfo_utf8_to_uhc(m *base.Module) int32
//go:linkname F_uhc_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_uhc_to_utf8
func F_uhc_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_uhc github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_utf8_to_uhc
func F_utf8_to_uhc(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_win github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_utf8_and_win
func F_Pg_magic_func_utf8_and_win(m *base.Module) int32
//go:linkname F_pg_finfo_win_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_win_to_utf8
func F_pg_finfo_win_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_win github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_utf8_to_win
func F_pg_finfo_utf8_to_win(m *base.Module) int32
//go:linkname F_win_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_win_to_utf8
func F_win_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_win github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_utf8_to_win
func F_utf8_to_win(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_pgcrypto github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_Pg_magic_func_pgcrypto
func F_Pg_magic_func_pgcrypto(m *base.Module) int32
//go:linkname F__PG_init_pgcrypto github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__PG_init_pgcrypto
func F__PG_init_pgcrypto(m *base.Module)
//go:linkname F_pg_finfo_pg_digest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_pg_digest
func F_pg_finfo_pg_digest(m *base.Module) int32
//go:linkname F_pg_digest github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_digest
func F_pg_digest(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_hmac github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_hmac
func F_pg_finfo_pg_hmac(m *base.Module) int32
//go:linkname F_pg_hmac github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_hmac
func F_pg_hmac(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_gen_salt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_gen_salt
func F_pg_finfo_pg_gen_salt(m *base.Module) int32
//go:linkname F_pg_gen_salt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_gen_salt
func F_pg_gen_salt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_gen_salt_rounds github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_gen_salt_rounds
func F_pg_finfo_pg_gen_salt_rounds(m *base.Module) int32
//go:linkname F_pg_gen_salt_rounds github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_gen_salt_rounds
func F_pg_gen_salt_rounds(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_crypt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_pg_crypt
func F_pg_finfo_pg_crypt(m *base.Module) int32
//go:linkname F_pg_crypt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_crypt
func F_pg_crypt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_encrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_encrypt
func F_pg_finfo_pg_encrypt(m *base.Module) int32
//go:linkname F_pg_encrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_encrypt
func F_pg_encrypt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_decrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pg_decrypt
func F_pg_finfo_pg_decrypt(m *base.Module) int32
//go:linkname F_pg_decrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_decrypt
func F_pg_decrypt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_encrypt_iv github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_encrypt_iv
func F_pg_finfo_pg_encrypt_iv(m *base.Module) int32
//go:linkname F_pg_encrypt_iv github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_encrypt_iv
func F_pg_encrypt_iv(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_decrypt_iv github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pg_decrypt_iv
func F_pg_finfo_pg_decrypt_iv(m *base.Module) int32
//go:linkname F_pg_decrypt_iv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_decrypt_iv
func F_pg_decrypt_iv(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_random_bytes github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pg_random_bytes
func F_pg_finfo_pg_random_bytes(m *base.Module) int32
//go:linkname F_pg_random_bytes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_random_bytes
func F_pg_random_bytes(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_random_uuid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pg_random_uuid
func F_pg_finfo_pg_random_uuid(m *base.Module) int32
//go:linkname F_pg_random_uuid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_random_uuid
func F_pg_random_uuid(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pg_check_fipsmode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_pg_check_fipsmode
func F_pg_finfo_pg_check_fipsmode(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_sym_encrypt_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pgp_sym_encrypt_bytea
func F_pg_finfo_pgp_sym_encrypt_bytea(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_sym_encrypt_text github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pgp_sym_encrypt_text
func F_pg_finfo_pgp_sym_encrypt_text(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_sym_decrypt_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pgp_sym_decrypt_bytea
func F_pg_finfo_pgp_sym_decrypt_bytea(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_sym_decrypt_text github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pgp_sym_decrypt_text
func F_pg_finfo_pgp_sym_decrypt_text(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_pub_encrypt_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pgp_pub_encrypt_bytea
func F_pg_finfo_pgp_pub_encrypt_bytea(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_pub_encrypt_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pgp_pub_encrypt_text
func F_pg_finfo_pgp_pub_encrypt_text(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_pub_decrypt_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pgp_pub_decrypt_bytea
func F_pg_finfo_pgp_pub_decrypt_bytea(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_pub_decrypt_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pgp_pub_decrypt_text
func F_pg_finfo_pgp_pub_decrypt_text(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_key_id_w github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pgp_key_id_w
func F_pg_finfo_pgp_key_id_w(m *base.Module) int32
//go:linkname F_pg_finfo_pg_armor github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_armor
func F_pg_finfo_pg_armor(m *base.Module) int32
//go:linkname F_pg_finfo_pg_dearmor github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_pg_dearmor
func F_pg_finfo_pg_dearmor(m *base.Module) int32
//go:linkname F_pg_finfo_pgp_armor_headers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_pgp_armor_headers
func F_pg_finfo_pgp_armor_headers(m *base.Module) int32
//go:linkname F_pgp_sym_encrypt_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_sym_encrypt_bytea
func F_pgp_sym_encrypt_bytea(m *base.Module, l0 int32) int32
//go:linkname F_pgp_sym_encrypt_text github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgp_sym_encrypt_text
func F_pgp_sym_encrypt_text(m *base.Module, l0 int32) int32
//go:linkname F_pgp_sym_decrypt_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgp_sym_decrypt_bytea
func F_pgp_sym_decrypt_bytea(m *base.Module, l0 int32) int32
//go:linkname F_pgp_sym_decrypt_text github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_sym_decrypt_text
func F_pgp_sym_decrypt_text(m *base.Module, l0 int32) int32
//go:linkname F_pgp_pub_encrypt_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_pub_encrypt_bytea
func F_pgp_pub_encrypt_bytea(m *base.Module, l0 int32) int32
//go:linkname F_pgp_pub_encrypt_text github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_pub_encrypt_text
func F_pgp_pub_encrypt_text(m *base.Module, l0 int32) int32
//go:linkname F_pgp_pub_decrypt_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_pub_decrypt_bytea
func F_pgp_pub_decrypt_bytea(m *base.Module, l0 int32) int32
//go:linkname F_pgp_pub_decrypt_text github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_pub_decrypt_text
func F_pgp_pub_decrypt_text(m *base.Module, l0 int32) int32
//go:linkname F_pg_armor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_armor
func F_pg_armor(m *base.Module, l0 int32) int32
//go:linkname F_pg_dearmor github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_dearmor
func F_pg_dearmor(m *base.Module, l0 int32) int32
//go:linkname F_pgp_armor_headers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgp_armor_headers
func F_pgp_armor_headers(m *base.Module, l0 int32) int32
//go:linkname F_pgp_key_id_w github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgp_key_id_w
func F_pgp_key_id_w(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_citext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_citext
func F_Pg_magic_func_citext(m *base.Module) int32
//go:linkname F_pg_finfo_citext_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_citext_cmp
func F_pg_finfo_citext_cmp(m *base.Module) int32
//go:linkname F_citext_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_citext_cmp
func F_citext_cmp(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_pattern_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_citext_pattern_cmp
func F_pg_finfo_citext_pattern_cmp(m *base.Module) int32
//go:linkname F_citext_pattern_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_citext_pattern_cmp
func F_citext_pattern_cmp(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_citext_hash
func F_pg_finfo_citext_hash(m *base.Module) int32
//go:linkname F_citext_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_citext_hash
func F_citext_hash(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_hash_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_citext_hash_extended
func F_pg_finfo_citext_hash_extended(m *base.Module) int32
//go:linkname F_citext_hash_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_citext_hash_extended
func F_citext_hash_extended(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_citext_eq
func F_pg_finfo_citext_eq(m *base.Module) int32
//go:linkname F_citext_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_citext_eq
func F_citext_eq(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_citext_ne
func F_pg_finfo_citext_ne(m *base.Module) int32
//go:linkname F_citext_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_citext_ne
func F_citext_ne(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_citext_lt
func F_pg_finfo_citext_lt(m *base.Module) int32
//go:linkname F_citext_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_citext_lt
func F_citext_lt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_le github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_citext_le
func F_pg_finfo_citext_le(m *base.Module) int32
//go:linkname F_citext_le github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_citext_le
func F_citext_le(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_citext_gt
func F_pg_finfo_citext_gt(m *base.Module) int32
//go:linkname F_citext_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_citext_gt
func F_citext_gt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_citext_ge
func F_pg_finfo_citext_ge(m *base.Module) int32
//go:linkname F_citext_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_citext_ge
func F_citext_ge(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_pattern_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_citext_pattern_lt
func F_pg_finfo_citext_pattern_lt(m *base.Module) int32
//go:linkname F_citext_pattern_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_citext_pattern_lt
func F_citext_pattern_lt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_pattern_le github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_citext_pattern_le
func F_pg_finfo_citext_pattern_le(m *base.Module) int32
//go:linkname F_citext_pattern_le github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_citext_pattern_le
func F_citext_pattern_le(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_pattern_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_citext_pattern_gt
func F_pg_finfo_citext_pattern_gt(m *base.Module) int32
//go:linkname F_citext_pattern_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_citext_pattern_gt
func F_citext_pattern_gt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_pattern_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_citext_pattern_ge
func F_pg_finfo_citext_pattern_ge(m *base.Module) int32
//go:linkname F_citext_pattern_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_citext_pattern_ge
func F_citext_pattern_ge(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_smaller github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_citext_smaller
func F_pg_finfo_citext_smaller(m *base.Module) int32
//go:linkname F_citext_smaller github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_citext_smaller
func F_citext_smaller(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_citext_larger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_citext_larger
func F_pg_finfo_citext_larger(m *base.Module) int32
//go:linkname F_citext_larger github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_citext_larger
func F_citext_larger(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_extract_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_trgm
func F_pg_finfo_gin_extract_trgm(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_extract_value_trgm
func F_pg_finfo_gin_extract_value_trgm(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_extract_query_trgm
func F_pg_finfo_gin_extract_query_trgm(m *base.Module) int32
//go:linkname F_pg_finfo_gin_trgm_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_trgm_consistent
func F_pg_finfo_gin_trgm_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gin_trgm_triconsistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_trgm_triconsistent
func F_pg_finfo_gin_trgm_triconsistent(m *base.Module) int32
//go:linkname F_gin_extract_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gin_extract_trgm
func F_gin_extract_trgm(m *base.Module, l0 int32) int32
//go:linkname F_gin_extract_query_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_query_trgm
func F_gin_extract_query_trgm(m *base.Module, l0 int32) int32
//go:linkname F_gin_extract_value_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gin_extract_value_trgm
func F_gin_extract_value_trgm(m *base.Module, l0 int32) int32
//go:linkname F_gin_trgm_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_trgm_consistent
func F_gin_trgm_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gin_trgm_triconsistent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gin_trgm_triconsistent
func F_gin_trgm_triconsistent(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gtrgm_in github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gtrgm_in
func F_pg_finfo_gtrgm_in(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_out github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gtrgm_out
func F_pg_finfo_gtrgm_out(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gtrgm_compress
func F_pg_finfo_gtrgm_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gtrgm_decompress
func F_pg_finfo_gtrgm_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gtrgm_consistent
func F_pg_finfo_gtrgm_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gtrgm_distance
func F_pg_finfo_gtrgm_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gtrgm_union
func F_pg_finfo_gtrgm_union(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gtrgm_same
func F_pg_finfo_gtrgm_same(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gtrgm_penalty
func F_pg_finfo_gtrgm_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gtrgm_picksplit
func F_pg_finfo_gtrgm_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gtrgm_options github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gtrgm_options
func F_pg_finfo_gtrgm_options(m *base.Module) int32
//go:linkname F_gtrgm_in github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gtrgm_in
func F_gtrgm_in(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gtrgm_out
func F_gtrgm_out(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gtrgm_compress
func F_gtrgm_compress(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gtrgm_decompress
func F_gtrgm_decompress(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gtrgm_consistent
func F_gtrgm_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gtrgm_distance
func F_gtrgm_distance(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gtrgm_union
func F_gtrgm_union(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gtrgm_same
func F_gtrgm_same(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gtrgm_penalty
func F_gtrgm_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gtrgm_picksplit
func F_gtrgm_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gtrgm_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gtrgm_options
func F_gtrgm_options(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_pg_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_pg_trgm
func F_Pg_magic_func_pg_trgm(m *base.Module) int32
//go:linkname F_pg_finfo_set_limit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_set_limit
func F_pg_finfo_set_limit(m *base.Module) int32
//go:linkname F_pg_finfo_show_limit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_show_limit
func F_pg_finfo_show_limit(m *base.Module) int32
//go:linkname F_pg_finfo_show_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_show_trgm
func F_pg_finfo_show_trgm(m *base.Module) int32
//go:linkname F_pg_finfo_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_similarity
func F_pg_finfo_similarity(m *base.Module) int32
//go:linkname F_pg_finfo_word_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_word_similarity
func F_pg_finfo_word_similarity(m *base.Module) int32
//go:linkname F_pg_finfo_strict_word_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_strict_word_similarity
func F_pg_finfo_strict_word_similarity(m *base.Module) int32
//go:linkname F_pg_finfo_similarity_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_similarity_dist
func F_pg_finfo_similarity_dist(m *base.Module) int32
//go:linkname F_pg_finfo_similarity_op github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_similarity_op
func F_pg_finfo_similarity_op(m *base.Module) int32
//go:linkname F_pg_finfo_word_similarity_op github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_word_similarity_op
func F_pg_finfo_word_similarity_op(m *base.Module) int32
//go:linkname F_pg_finfo_word_similarity_commutator_op github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_word_similarity_commutator_op
func F_pg_finfo_word_similarity_commutator_op(m *base.Module) int32
//go:linkname F_pg_finfo_word_similarity_dist_op github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_word_similarity_dist_op
func F_pg_finfo_word_similarity_dist_op(m *base.Module) int32
//go:linkname F_pg_finfo_word_similarity_dist_commutator_op github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_word_similarity_dist_commutator_op
func F_pg_finfo_word_similarity_dist_commutator_op(m *base.Module) int32
//go:linkname F_pg_finfo_strict_word_similarity_op github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_strict_word_similarity_op
func F_pg_finfo_strict_word_similarity_op(m *base.Module) int32
//go:linkname F_pg_finfo_strict_word_similarity_commutator_op github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_strict_word_similarity_commutator_op
func F_pg_finfo_strict_word_similarity_commutator_op(m *base.Module) int32
//go:linkname F_pg_finfo_strict_word_similarity_dist_op github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_strict_word_similarity_dist_op
func F_pg_finfo_strict_word_similarity_dist_op(m *base.Module) int32
//go:linkname F_pg_finfo_strict_word_similarity_dist_commutator_op github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_strict_word_similarity_dist_commutator_op
func F_pg_finfo_strict_word_similarity_dist_commutator_op(m *base.Module) int32
//go:linkname F__PG_init_pg_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__PG_init_pg_trgm
func F__PG_init_pg_trgm(m *base.Module)
//go:linkname F_set_limit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_limit
func F_set_limit(m *base.Module, l0 int32) int32
//go:linkname F_show_limit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_show_limit
func F_show_limit(m *base.Module, l0 int32) int32
//go:linkname F_show_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_show_trgm
func F_show_trgm(m *base.Module, l0 int32) int32
//go:linkname F_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_similarity
func F_similarity(m *base.Module, l0 int32) int32
//go:linkname F_word_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_word_similarity
func F_word_similarity(m *base.Module, l0 int32) int32
//go:linkname F_strict_word_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strict_word_similarity
func F_strict_word_similarity(m *base.Module, l0 int32) int32
//go:linkname F_similarity_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_similarity_dist
func F_similarity_dist(m *base.Module, l0 int32) int32
//go:linkname F_similarity_op github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_similarity_op
func F_similarity_op(m *base.Module, l0 int32) int32
//go:linkname F_word_similarity_op github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_word_similarity_op
func F_word_similarity_op(m *base.Module, l0 int32) int32
//go:linkname F_word_similarity_commutator_op github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_word_similarity_commutator_op
func F_word_similarity_commutator_op(m *base.Module, l0 int32) int32
//go:linkname F_word_similarity_dist_op github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_word_similarity_dist_op
func F_word_similarity_dist_op(m *base.Module, l0 int32) int32
//go:linkname F_word_similarity_dist_commutator_op github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_word_similarity_dist_commutator_op
func F_word_similarity_dist_commutator_op(m *base.Module, l0 int32) int32
//go:linkname F_strict_word_similarity_op github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strict_word_similarity_op
func F_strict_word_similarity_op(m *base.Module, l0 int32) int32
//go:linkname F_strict_word_similarity_commutator_op github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_strict_word_similarity_commutator_op
func F_strict_word_similarity_commutator_op(m *base.Module, l0 int32) int32
//go:linkname F_strict_word_similarity_dist_op github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_strict_word_similarity_dist_op
func F_strict_word_similarity_dist_op(m *base.Module, l0 int32) int32
//go:linkname F_strict_word_similarity_dist_commutator_op github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strict_word_similarity_dist_commutator_op
func F_strict_word_similarity_dist_commutator_op(m *base.Module, l0 int32) int32
//go:linkname F_hstoreUpgrade github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUpgrade
func F_hstoreUpgrade(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_version_diag github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hstore_version_diag
func F_pg_finfo_hstore_version_diag(m *base.Module) int32
//go:linkname F_hstore_version_diag github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstore_version_diag
func F_hstore_version_diag(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_extract_hstore github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_hstore
func F_pg_finfo_gin_extract_hstore(m *base.Module) int32
//go:linkname F_gin_extract_hstore github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_extract_hstore
func F_gin_extract_hstore(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_extract_hstore_query github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_extract_hstore_query
func F_pg_finfo_gin_extract_hstore_query(m *base.Module) int32
//go:linkname F_gin_extract_hstore_query github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gin_extract_hstore_query
func F_gin_extract_hstore_query(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_consistent_hstore github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_consistent_hstore
func F_pg_finfo_gin_consistent_hstore(m *base.Module) int32
//go:linkname F_gin_consistent_hstore github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_consistent_hstore
func F_gin_consistent_hstore(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ghstore_in github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ghstore_in
func F_pg_finfo_ghstore_in(m *base.Module) int32
//go:linkname F_pg_finfo_ghstore_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ghstore_out
func F_pg_finfo_ghstore_out(m *base.Module) int32
//go:linkname F_ghstore_in github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ghstore_in
func F_ghstore_in(m *base.Module, l0 int32) int32
//go:linkname F_ghstore_out github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ghstore_out
func F_ghstore_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ghstore_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ghstore_consistent
func F_pg_finfo_ghstore_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_ghstore_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ghstore_compress
func F_pg_finfo_ghstore_compress(m *base.Module) int32
//go:linkname F_pg_finfo_ghstore_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ghstore_decompress
func F_pg_finfo_ghstore_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_ghstore_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ghstore_penalty
func F_pg_finfo_ghstore_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_ghstore_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_ghstore_picksplit
func F_pg_finfo_ghstore_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_ghstore_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ghstore_union
func F_pg_finfo_ghstore_union(m *base.Module) int32
//go:linkname F_pg_finfo_ghstore_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_ghstore_same
func F_pg_finfo_ghstore_same(m *base.Module) int32
//go:linkname F_pg_finfo_ghstore_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ghstore_options
func F_pg_finfo_ghstore_options(m *base.Module) int32
//go:linkname F_ghstore_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ghstore_compress
func F_ghstore_compress(m *base.Module, l0 int32) int32
//go:linkname F_ghstore_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ghstore_same
func F_ghstore_same(m *base.Module, l0 int32) int32
//go:linkname F_ghstore_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ghstore_union
func F_ghstore_union(m *base.Module, l0 int32) int32
//go:linkname F_ghstore_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ghstore_penalty
func F_ghstore_penalty(m *base.Module, l0 int32) int32
//go:linkname F_ghstore_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ghstore_picksplit
func F_ghstore_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_ghstore_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ghstore_consistent
func F_ghstore_consistent(m *base.Module, l0 int32) int32
//go:linkname F_ghstore_options github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ghstore_options
func F_ghstore_options(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_hstore github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_hstore
func F_Pg_magic_func_hstore(m *base.Module) int32
//go:linkname F_pg_finfo_tconvert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_tconvert
func F_pg_finfo_tconvert(m *base.Module) int32
//go:linkname F_tconvert github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tconvert
func F_tconvert(m *base.Module, l0 int32) int32
//go:linkname F_hstore_from_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hstore_from_text
func F_hstore_from_text(m *base.Module, l0 int32) int32
//go:linkname F_hstoreUniquePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUniquePairs
func F_hstoreUniquePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hstoreCheckKeyLen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstoreCheckKeyLen
func F_hstoreCheckKeyLen(m *base.Module, l0 int32) int32
//go:linkname F_hstoreCheckValLen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstoreCheckValLen
func F_hstoreCheckValLen(m *base.Module, l0 int32) int32
//go:linkname F_hstorePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstorePairs
func F_hstorePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_finfo_hstore_in github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hstore_in
func F_pg_finfo_hstore_in(m *base.Module) int32
//go:linkname F_hstore_in github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_in
func F_hstore_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hstore_recv
func F_pg_finfo_hstore_recv(m *base.Module) int32
//go:linkname F_hstore_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_recv
func F_hstore_recv(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_from_text github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hstore_from_text
func F_pg_finfo_hstore_from_text(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_from_arrays github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_hstore_from_arrays
func F_pg_finfo_hstore_from_arrays(m *base.Module) int32
//go:linkname F_hstore_from_arrays github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstore_from_arrays
func F_hstore_from_arrays(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_from_array github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hstore_from_array
func F_pg_finfo_hstore_from_array(m *base.Module) int32
//go:linkname F_hstore_from_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hstore_from_array
func F_hstore_from_array(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_from_record github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_hstore_from_record
func F_pg_finfo_hstore_from_record(m *base.Module) int32
//go:linkname F_hstore_from_record github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_from_record
func F_hstore_from_record(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_populate_record github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_hstore_populate_record
func F_pg_finfo_hstore_populate_record(m *base.Module) int32
//go:linkname F_hstore_populate_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstore_populate_record
func F_hstore_populate_record(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_out github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hstore_out
func F_pg_finfo_hstore_out(m *base.Module) int32
//go:linkname F_hstore_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_out
func F_hstore_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_send github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hstore_send
func F_pg_finfo_hstore_send(m *base.Module) int32
//go:linkname F_hstore_send github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstore_send
func F_hstore_send(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_to_json_loose github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hstore_to_json_loose
func F_pg_finfo_hstore_to_json_loose(m *base.Module) int32
//go:linkname F_hstore_to_json_loose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_to_json_loose
func F_hstore_to_json_loose(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_to_json github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hstore_to_json
func F_pg_finfo_hstore_to_json(m *base.Module) int32
//go:linkname F_hstore_to_json github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_to_json
func F_hstore_to_json(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_to_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hstore_to_jsonb
func F_pg_finfo_hstore_to_jsonb(m *base.Module) int32
//go:linkname F_hstore_to_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_to_jsonb
func F_hstore_to_jsonb(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_to_jsonb_loose github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hstore_to_jsonb_loose
func F_pg_finfo_hstore_to_jsonb_loose(m *base.Module) int32
//go:linkname F_hstore_to_jsonb_loose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_to_jsonb_loose
func F_hstore_to_jsonb_loose(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_fetchval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_fetchval
func F_pg_finfo_fetchval(m *base.Module) int32
//go:linkname F_fetchval github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fetchval
func F_fetchval(m *base.Module, l0 int32) int32
//go:linkname F_hstore_fetchval github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstore_fetchval
func F_hstore_fetchval(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_exists
func F_pg_finfo_exists(m *base.Module) int32
//go:linkname F_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exists
func F_exists(m *base.Module, l0 int32) int32
//go:linkname F_hstore_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hstore_exists
func F_hstore_exists(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_defined github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_defined
func F_pg_finfo_defined(m *base.Module) int32
//go:linkname F_defined github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_defined
func F_defined(m *base.Module, l0 int32) int32
//go:linkname F_hstore_defined github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_defined
func F_hstore_defined(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_delete
func F_pg_finfo_delete(m *base.Module) int32
//go:linkname F_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_delete
func F_delete(m *base.Module, l0 int32) int32
//go:linkname F_hstore_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_delete
func F_hstore_delete(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hs_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hs_concat
func F_pg_finfo_hs_concat(m *base.Module) int32
//go:linkname F_hs_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hs_concat
func F_hs_concat(m *base.Module, l0 int32) int32
//go:linkname F_hstore_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_concat
func F_hstore_concat(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hs_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hs_contains
func F_pg_finfo_hs_contains(m *base.Module) int32
//go:linkname F_hs_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hs_contains
func F_hs_contains(m *base.Module, l0 int32) int32
//go:linkname F_hstore_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_contains
func F_hstore_contains(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hs_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hs_contained
func F_pg_finfo_hs_contained(m *base.Module) int32
//go:linkname F_hs_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hs_contained
func F_hs_contained(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_akeys github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_akeys
func F_pg_finfo_akeys(m *base.Module) int32
//go:linkname F_akeys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_akeys
func F_akeys(m *base.Module, l0 int32) int32
//go:linkname F_hstore_akeys github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstore_akeys
func F_hstore_akeys(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_avals github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_avals
func F_pg_finfo_avals(m *base.Module) int32
//go:linkname F_avals github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_avals
func F_avals(m *base.Module, l0 int32) int32
//go:linkname F_hstore_avals github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_avals
func F_hstore_avals(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_skeys github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_skeys
func F_pg_finfo_skeys(m *base.Module) int32
//go:linkname F_skeys github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_skeys
func F_skeys(m *base.Module, l0 int32) int32
//go:linkname F_hstore_skeys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_skeys
func F_hstore_skeys(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_svals github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_svals
func F_pg_finfo_svals(m *base.Module) int32
//go:linkname F_svals github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_svals
func F_svals(m *base.Module, l0 int32) int32
//go:linkname F_hstore_svals github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstore_svals
func F_hstore_svals(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_each github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_each
func F_pg_finfo_each(m *base.Module) int32
//go:linkname F_each github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_each
func F_each(m *base.Module, l0 int32) int32
//go:linkname F_hstore_each github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_each
func F_hstore_each(m *base.Module, l0 int32) int32
//go:linkname F_hstoreFindKey github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstoreFindKey
func F_hstoreFindKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hstoreArrayToPairs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstoreArrayToPairs
func F_hstoreArrayToPairs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_finfo_hstore_fetchval github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_hstore_fetchval
func F_pg_finfo_hstore_fetchval(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hstore_exists
func F_pg_finfo_hstore_exists(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_exists_any github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hstore_exists_any
func F_pg_finfo_hstore_exists_any(m *base.Module) int32
//go:linkname F_hstore_exists_any github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_exists_any
func F_hstore_exists_any(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_exists_all github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_hstore_exists_all
func F_pg_finfo_hstore_exists_all(m *base.Module) int32
//go:linkname F_hstore_exists_all github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_exists_all
func F_hstore_exists_all(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_defined github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hstore_defined
func F_pg_finfo_hstore_defined(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hstore_delete
func F_pg_finfo_hstore_delete(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_delete_array github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hstore_delete_array
func F_pg_finfo_hstore_delete_array(m *base.Module) int32
//go:linkname F_hstore_delete_array github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_delete_array
func F_hstore_delete_array(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_delete_hstore github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_hstore_delete_hstore
func F_pg_finfo_hstore_delete_hstore(m *base.Module) int32
//go:linkname F_hstore_delete_hstore github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_delete_hstore
func F_hstore_delete_hstore(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hstore_concat
func F_pg_finfo_hstore_concat(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_slice_to_array github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hstore_slice_to_array
func F_pg_finfo_hstore_slice_to_array(m *base.Module) int32
//go:linkname F_hstore_slice_to_array github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_slice_to_array
func F_hstore_slice_to_array(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_slice_to_hstore github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hstore_slice_to_hstore
func F_pg_finfo_hstore_slice_to_hstore(m *base.Module) int32
//go:linkname F_hstore_slice_to_hstore github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_slice_to_hstore
func F_hstore_slice_to_hstore(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_akeys github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hstore_akeys
func F_pg_finfo_hstore_akeys(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_avals github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hstore_avals
func F_pg_finfo_hstore_avals(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_to_array github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hstore_to_array
func F_pg_finfo_hstore_to_array(m *base.Module) int32
//go:linkname F_hstore_to_array github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_to_array
func F_hstore_to_array(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_to_matrix github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hstore_to_matrix
func F_pg_finfo_hstore_to_matrix(m *base.Module) int32
//go:linkname F_hstore_to_matrix github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_to_matrix
func F_hstore_to_matrix(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_skeys github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hstore_skeys
func F_pg_finfo_hstore_skeys(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_svals github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_hstore_svals
func F_pg_finfo_hstore_svals(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hstore_contains
func F_pg_finfo_hstore_contains(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_hstore_contained
func F_pg_finfo_hstore_contained(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_each github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_hstore_each
func F_pg_finfo_hstore_each(m *base.Module) int32
//go:linkname F_pg_finfo_hstore_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_hstore_cmp
func F_pg_finfo_hstore_cmp(m *base.Module) int32
//go:linkname F_hstore_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hstore_cmp
func F_hstore_cmp(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hstore_eq
func F_pg_finfo_hstore_eq(m *base.Module) int32
//go:linkname F_hstore_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_eq
func F_hstore_eq(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hstore_ne
func F_pg_finfo_hstore_ne(m *base.Module) int32
//go:linkname F_hstore_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_ne
func F_hstore_ne(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hstore_gt
func F_pg_finfo_hstore_gt(m *base.Module) int32
//go:linkname F_hstore_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_gt
func F_hstore_gt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hstore_ge
func F_pg_finfo_hstore_ge(m *base.Module) int32
//go:linkname F_hstore_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstore_ge
func F_hstore_ge(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hstore_lt
func F_pg_finfo_hstore_lt(m *base.Module) int32
//go:linkname F_hstore_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_lt
func F_hstore_lt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_le github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hstore_le
func F_pg_finfo_hstore_le(m *base.Module) int32
//go:linkname F_hstore_le github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_le
func F_hstore_le(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hstore_hash
func F_pg_finfo_hstore_hash(m *base.Module) int32
//go:linkname F_hstore_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstore_hash
func F_hstore_hash(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_hash_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hstore_hash_extended
func F_pg_finfo_hstore_hash_extended(m *base.Module) int32
//go:linkname F_hstore_hash_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstore_hash_extended
func F_hstore_hash_extended(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hstore_subscript_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hstore_subscript_handler
func F_pg_finfo_hstore_subscript_handler(m *base.Module) int32
//go:linkname F_hstore_subscript_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hstore_subscript_handler
func F_hstore_subscript_handler(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo__ltree_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo__ltree_compress
func F_pg_finfo__ltree_compress(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_same github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo__ltree_same
func F_pg_finfo__ltree_same(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo__ltree_union
func F_pg_finfo__ltree_union(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo__ltree_penalty
func F_pg_finfo__ltree_penalty(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo__ltree_picksplit
func F_pg_finfo__ltree_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo__ltree_consistent
func F_pg_finfo__ltree_consistent(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_gist_options github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo__ltree_gist_options
func F_pg_finfo__ltree_gist_options(m *base.Module) int32
//go:linkname F__ltree_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__ltree_compress
func F__ltree_compress(m *base.Module, l0 int32) int32
//go:linkname F__ltree_same github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__ltree_same
func F__ltree_same(m *base.Module, l0 int32) int32
//go:linkname F__ltree_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__ltree_union
func F__ltree_union(m *base.Module, l0 int32) int32
//go:linkname F__ltree_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__ltree_penalty
func F__ltree_penalty(m *base.Module, l0 int32) int32
//go:linkname F__ltree_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__ltree_picksplit
func F__ltree_picksplit(m *base.Module, l0 int32) int32
//go:linkname F__ltree_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__ltree_consistent
func F__ltree_consistent(m *base.Module, l0 int32) int32
//go:linkname F__ltree_gist_options github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__ltree_gist_options
func F__ltree_gist_options(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo__ltree_isparent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo__ltree_isparent
func F_pg_finfo__ltree_isparent(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_r_isparent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo__ltree_r_isparent
func F_pg_finfo__ltree_r_isparent(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_risparent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo__ltree_risparent
func F_pg_finfo__ltree_risparent(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_r_risparent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo__ltree_r_risparent
func F_pg_finfo__ltree_r_risparent(m *base.Module) int32
//go:linkname F_pg_finfo__ltq_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo__ltq_regex
func F_pg_finfo__ltq_regex(m *base.Module) int32
//go:linkname F_pg_finfo__ltq_rregex github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo__ltq_rregex
func F_pg_finfo__ltq_rregex(m *base.Module) int32
//go:linkname F_pg_finfo__lt_q_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo__lt_q_regex
func F_pg_finfo__lt_q_regex(m *base.Module) int32
//go:linkname F_pg_finfo__lt_q_rregex github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo__lt_q_rregex
func F_pg_finfo__lt_q_rregex(m *base.Module) int32
//go:linkname F_pg_finfo__ltxtq_exec github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo__ltxtq_exec
func F_pg_finfo__ltxtq_exec(m *base.Module) int32
//go:linkname F_pg_finfo__ltxtq_rexec github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo__ltxtq_rexec
func F_pg_finfo__ltxtq_rexec(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_extract_isparent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo__ltree_extract_isparent
func F_pg_finfo__ltree_extract_isparent(m *base.Module) int32
//go:linkname F_pg_finfo__ltree_extract_risparent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo__ltree_extract_risparent
func F_pg_finfo__ltree_extract_risparent(m *base.Module) int32
//go:linkname F_pg_finfo__ltq_extract_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo__ltq_extract_regex
func F_pg_finfo__ltq_extract_regex(m *base.Module) int32
//go:linkname F_pg_finfo__ltxtq_extract_exec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo__ltxtq_extract_exec
func F_pg_finfo__ltxtq_extract_exec(m *base.Module) int32
//go:linkname F_pg_finfo__lca github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo__lca
func F_pg_finfo__lca(m *base.Module) int32
//go:linkname F__ltree_isparent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__ltree_isparent
func F__ltree_isparent(m *base.Module, l0 int32) int32
//go:linkname F__ltree_r_isparent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__ltree_r_isparent
func F__ltree_r_isparent(m *base.Module, l0 int32) int32
//go:linkname F__ltree_risparent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__ltree_risparent
func F__ltree_risparent(m *base.Module, l0 int32) int32
//go:linkname F__ltree_r_risparent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__ltree_r_risparent
func F__ltree_r_risparent(m *base.Module, l0 int32) int32
//go:linkname F__ltq_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__ltq_regex
func F__ltq_regex(m *base.Module, l0 int32) int32
//go:linkname F__ltq_rregex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__ltq_rregex
func F__ltq_rregex(m *base.Module, l0 int32) int32
//go:linkname F__lt_q_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__lt_q_regex
func F__lt_q_regex(m *base.Module, l0 int32) int32
//go:linkname F__lt_q_rregex github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__lt_q_rregex
func F__lt_q_rregex(m *base.Module, l0 int32) int32
//go:linkname F__ltxtq_exec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__ltxtq_exec
func F__ltxtq_exec(m *base.Module, l0 int32) int32
//go:linkname F__ltxtq_rexec github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__ltxtq_rexec
func F__ltxtq_rexec(m *base.Module, l0 int32) int32
//go:linkname F__ltree_extract_isparent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__ltree_extract_isparent
func F__ltree_extract_isparent(m *base.Module, l0 int32) int32
//go:linkname F__ltree_extract_risparent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__ltree_extract_risparent
func F__ltree_extract_risparent(m *base.Module, l0 int32) int32
//go:linkname F__ltq_extract_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__ltq_extract_regex
func F__ltq_extract_regex(m *base.Module, l0 int32) int32
//go:linkname F__ltxtq_extract_exec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__ltxtq_extract_exec
func F__ltxtq_extract_exec(m *base.Module, l0 int32) int32
//go:linkname F__lca github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__lca
func F__lca(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltq_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_ltq_regex
func F_pg_finfo_ltq_regex(m *base.Module) int32
//go:linkname F_pg_finfo_ltq_rregex github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ltq_rregex
func F_pg_finfo_ltq_rregex(m *base.Module) int32
//go:linkname F_pg_finfo_lt_q_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_lt_q_regex
func F_pg_finfo_lt_q_regex(m *base.Module) int32
//go:linkname F_pg_finfo_lt_q_rregex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_lt_q_rregex
func F_pg_finfo_lt_q_rregex(m *base.Module) int32
//go:linkname F_ltq_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltq_regex
func F_ltq_regex(m *base.Module, l0 int32) int32
//go:linkname F_ltq_rregex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ltq_rregex
func F_ltq_rregex(m *base.Module, l0 int32) int32
//go:linkname F_lt_q_regex github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lt_q_regex
func F_lt_q_regex(m *base.Module, l0 int32) int32
//go:linkname F_lt_q_rregex github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lt_q_rregex
func F_lt_q_rregex(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltree_gist_in github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_ltree_gist_in
func F_pg_finfo_ltree_gist_in(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_gist_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ltree_gist_out
func F_pg_finfo_ltree_gist_out(m *base.Module) int32
//go:linkname F_ltree_gist_in github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ltree_gist_in
func F_ltree_gist_in(m *base.Module, l0 int32) int32
//go:linkname F_ltree_gist_out github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltree_gist_out
func F_ltree_gist_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltree_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_ltree_compress
func F_pg_finfo_ltree_compress(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_ltree_decompress
func F_pg_finfo_ltree_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_ltree_same
func F_pg_finfo_ltree_same(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_ltree_union
func F_pg_finfo_ltree_union(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ltree_penalty
func F_pg_finfo_ltree_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ltree_picksplit
func F_pg_finfo_ltree_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ltree_consistent
func F_pg_finfo_ltree_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_gist_options github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_ltree_gist_options
func F_pg_finfo_ltree_gist_options(m *base.Module) int32
//go:linkname F_ltree_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltree_compress
func F_ltree_compress(m *base.Module, l0 int32) int32
//go:linkname F_ltree_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ltree_same
func F_ltree_same(m *base.Module, l0 int32) int32
//go:linkname F_ltree_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltree_union
func F_ltree_union(m *base.Module, l0 int32) int32
//go:linkname F_ltree_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ltree_penalty
func F_ltree_penalty(m *base.Module, l0 int32) int32
//go:linkname F_ltree_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltree_picksplit
func F_ltree_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_ltree_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltree_consistent
func F_ltree_consistent(m *base.Module, l0 int32) int32
//go:linkname F_ltree_gist_options github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltree_gist_options
func F_ltree_gist_options(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltree_in github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_ltree_in
func F_pg_finfo_ltree_in(m *base.Module) int32
//go:linkname F_ltree_in github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltree_in
func F_ltree_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltree_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ltree_out
func F_pg_finfo_ltree_out(m *base.Module) int32
//go:linkname F_ltree_out github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltree_out
func F_ltree_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltree_send github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ltree_send
func F_pg_finfo_ltree_send(m *base.Module) int32
//go:linkname F_ltree_send github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ltree_send
func F_ltree_send(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltree_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ltree_recv
func F_pg_finfo_ltree_recv(m *base.Module) int32
//go:linkname F_ltree_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltree_recv
func F_ltree_recv(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_lquery_in github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_lquery_in
func F_pg_finfo_lquery_in(m *base.Module) int32
//go:linkname F_lquery_in github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lquery_in
func F_lquery_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_lquery_out github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_lquery_out
func F_pg_finfo_lquery_out(m *base.Module) int32
//go:linkname F_lquery_out github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lquery_out
func F_lquery_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_lquery_send github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_lquery_send
func F_pg_finfo_lquery_send(m *base.Module) int32
//go:linkname F_lquery_send github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lquery_send
func F_lquery_send(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_lquery_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_lquery_recv
func F_pg_finfo_lquery_recv(m *base.Module) int32
//go:linkname F_lquery_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lquery_recv
func F_lquery_recv(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_ltree github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_ltree
func F_Pg_magic_func_ltree(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_ltree_cmp
func F_pg_finfo_ltree_cmp(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ltree_lt
func F_pg_finfo_ltree_lt(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_le github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ltree_le
func F_pg_finfo_ltree_le(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ltree_eq
func F_pg_finfo_ltree_eq(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ltree_ne
func F_pg_finfo_ltree_ne(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ltree_ge
func F_pg_finfo_ltree_ge(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_ltree_gt
func F_pg_finfo_ltree_gt(m *base.Module) int32
//go:linkname F_pg_finfo_hash_ltree github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hash_ltree
func F_pg_finfo_hash_ltree(m *base.Module) int32
//go:linkname F_pg_finfo_hash_ltree_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hash_ltree_extended
func F_pg_finfo_hash_ltree_extended(m *base.Module) int32
//go:linkname F_pg_finfo_nlevel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_nlevel
func F_pg_finfo_nlevel(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_isparent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ltree_isparent
func F_pg_finfo_ltree_isparent(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_risparent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ltree_risparent
func F_pg_finfo_ltree_risparent(m *base.Module) int32
//go:linkname F_pg_finfo_subltree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_subltree
func F_pg_finfo_subltree(m *base.Module) int32
//go:linkname F_pg_finfo_subpath github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_subpath
func F_pg_finfo_subpath(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_index github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_ltree_index
func F_pg_finfo_ltree_index(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_addltree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ltree_addltree
func F_pg_finfo_ltree_addltree(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_addtext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_ltree_addtext
func F_pg_finfo_ltree_addtext(m *base.Module) int32
//go:linkname F_pg_finfo_ltree_textadd github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_ltree_textadd
func F_pg_finfo_ltree_textadd(m *base.Module) int32
//go:linkname F_pg_finfo_lca github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_lca
func F_pg_finfo_lca(m *base.Module) int32
//go:linkname F_pg_finfo_ltree2text github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ltree2text
func F_pg_finfo_ltree2text(m *base.Module) int32
//go:linkname F_pg_finfo_text2ltree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_text2ltree
func F_pg_finfo_text2ltree(m *base.Module) int32
//go:linkname F_pg_finfo_ltreeparentsel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ltreeparentsel
func F_pg_finfo_ltreeparentsel(m *base.Module) int32
//go:linkname F_ltree_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltree_cmp
func F_ltree_cmp(m *base.Module, l0 int32) int32
//go:linkname F_ltree_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltree_lt
func F_ltree_lt(m *base.Module, l0 int32) int32
//go:linkname F_ltree_le github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ltree_le
func F_ltree_le(m *base.Module, l0 int32) int32
//go:linkname F_ltree_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltree_eq
func F_ltree_eq(m *base.Module, l0 int32) int32
//go:linkname F_ltree_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ltree_ge
func F_ltree_ge(m *base.Module, l0 int32) int32
//go:linkname F_ltree_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltree_gt
func F_ltree_gt(m *base.Module, l0 int32) int32
//go:linkname F_ltree_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ltree_ne
func F_ltree_ne(m *base.Module, l0 int32) int32
//go:linkname F_hash_ltree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hash_ltree
func F_hash_ltree(m *base.Module, l0 int32) int32
//go:linkname F_hash_ltree_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_ltree_extended
func F_hash_ltree_extended(m *base.Module, l0 int32) int32
//go:linkname F_nlevel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_nlevel
func F_nlevel(m *base.Module, l0 int32) int32
//go:linkname F_ltree_isparent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltree_isparent
func F_ltree_isparent(m *base.Module, l0 int32) int32
//go:linkname F_ltree_risparent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltree_risparent
func F_ltree_risparent(m *base.Module, l0 int32) int32
//go:linkname F_subltree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_subltree
func F_subltree(m *base.Module, l0 int32) int32
//go:linkname F_subpath github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_subpath
func F_subpath(m *base.Module, l0 int32) int32
//go:linkname F_ltree_addltree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ltree_addltree
func F_ltree_addltree(m *base.Module, l0 int32) int32
//go:linkname F_ltree_addtext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltree_addtext
func F_ltree_addtext(m *base.Module, l0 int32) int32
//go:linkname F_ltree_index github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltree_index
func F_ltree_index(m *base.Module, l0 int32) int32
//go:linkname F_ltree_textadd github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ltree_textadd
func F_ltree_textadd(m *base.Module, l0 int32) int32
//go:linkname F_lca github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lca
func F_lca(m *base.Module, l0 int32) int32
//go:linkname F_text2ltree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text2ltree
func F_text2ltree(m *base.Module, l0 int32) int32
//go:linkname F_ltree2text github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltree2text
func F_ltree2text(m *base.Module, l0 int32) int32
//go:linkname F_ltreeparentsel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ltreeparentsel
func F_ltreeparentsel(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltxtq_in github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ltxtq_in
func F_pg_finfo_ltxtq_in(m *base.Module) int32
//go:linkname F_ltxtq_in github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ltxtq_in
func F_ltxtq_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltxtq_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ltxtq_recv
func F_pg_finfo_ltxtq_recv(m *base.Module) int32
//go:linkname F_ltxtq_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltxtq_recv
func F_ltxtq_recv(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltxtq_out github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ltxtq_out
func F_pg_finfo_ltxtq_out(m *base.Module) int32
//go:linkname F_ltxtq_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ltxtq_out
func F_ltxtq_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltxtq_send github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ltxtq_send
func F_pg_finfo_ltxtq_send(m *base.Module) int32
//go:linkname F_ltxtq_send github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ltxtq_send
func F_ltxtq_send(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ltxtq_exec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ltxtq_exec
func F_pg_finfo_ltxtq_exec(m *base.Module) int32
//go:linkname F_pg_finfo_ltxtq_rexec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ltxtq_rexec
func F_pg_finfo_ltxtq_rexec(m *base.Module) int32
//go:linkname F_ltxtq_exec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltxtq_exec
func F_ltxtq_exec(m *base.Module, l0 int32) int32
//go:linkname F_ltxtq_rexec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltxtq_rexec
func F_ltxtq_rexec(m *base.Module, l0 int32) int32
//go:linkname F__emscripten_memcpy_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memcpy_bulkmem
func F__emscripten_memcpy_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memset_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memset_bulkmem
func F__emscripten_memset_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___errno_location github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___errno_location
func F___errno_location(m *base.Module) int32
//go:linkname F_fflush github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F_htonl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_htonl
func F_htonl(m *base.Module, l0 int32) int32
//go:linkname F_htons github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_htons
func F_htons(m *base.Module, l0 int32) int32
//go:linkname F__emscripten_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_timeout
func F__emscripten_timeout(m *base.Module, l0 int32, l1 float64)
//go:linkname F_strerror github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strerror
func F_strerror(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F_emscripten_builtin_memalign github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_emscripten_builtin_memalign
func F_emscripten_builtin_memalign(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___funcs_on_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___funcs_on_exit
func F___funcs_on_exit(m *base.Module)
//go:linkname F__emscripten_stack_restore github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_stack_restore
func F__emscripten_stack_restore(m *base.Module, l0 int32)
//go:linkname F__emscripten_stack_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__emscripten_stack_alloc
func F__emscripten_stack_alloc(m *base.Module, l0 int32) int32
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
//go:linkname InitElemSeg_1_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.InitElemSeg_1_0
func InitElemSeg_1_0(m *base.Module)
//go:linkname InitElemSeg_1_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.InitElemSeg_1_1
func InitElemSeg_1_1(m *base.Module)
//go:linkname InitElemSeg_1_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.InitElemSeg_1_2
func InitElemSeg_1_2(m *base.Module)
//go:linkname InitElemSeg_1_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.InitElemSeg_1_3
func InitElemSeg_1_3(m *base.Module)
//go:linkname InitElemSeg_2_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.InitElemSeg_2_0
func InitElemSeg_2_0(m *base.Module)
//go:linkname InitElemSeg_2_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.InitElemSeg_2_1
func InitElemSeg_2_1(m *base.Module)
//go:linkname InitElemSeg_2_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.InitElemSeg_2_2
func InitElemSeg_2_2(m *base.Module)
//go:linkname InitElemSeg_2_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.InitElemSeg_2_3
func InitElemSeg_2_3(m *base.Module)
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
//go:linkname InitElemSeg_5_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_0
func InitElemSeg_5_0(m *base.Module)
//go:linkname InitElemSeg_5_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_1
func InitElemSeg_5_1(m *base.Module)
//go:linkname InitElemSeg_5_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_2
func InitElemSeg_5_2(m *base.Module)
//go:linkname InitElemSeg_5_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_3
func InitElemSeg_5_3(m *base.Module)
