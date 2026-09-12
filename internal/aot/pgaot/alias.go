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
//go:linkname F_pgmem_reset_session github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgmem_reset_session
func F_pgmem_reset_session(m *base.Module, l0 int32)
//go:linkname F_float4up github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_float4up
func F_float4up(m *base.Module, l0 int32) int32
//go:linkname F_gtsvector_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gtsvector_decompress
func F_gtsvector_decompress(m *base.Module, l0 int32) int32
//go:linkname F_tsquery_numnode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tsquery_numnode
func F_tsquery_numnode(m *base.Module, l0 int32) int32
//go:linkname F_xid8toxid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_xid8toxid
func F_xid8toxid(m *base.Module, l0 int32) int32
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
//go:linkname F_pg_finfo_iso8859_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_iso8859_to_utf8
func F_pg_finfo_iso8859_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_iso8859_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_utf8_to_iso8859_1
func F_pg_finfo_utf8_to_iso8859_1(m *base.Module) int32
//go:linkname F_iso8859_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_iso8859_to_utf8
func F_iso8859_to_utf8(m *base.Module, l0 int32) int32
//go:linkname F_utf8_to_iso8859_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_utf8_to_iso8859_1
func F_utf8_to_iso8859_1(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_utf8_and_iso8859_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_Pg_magic_func_utf8_and_iso8859_2
func F_Pg_magic_func_utf8_and_iso8859_2(m *base.Module) int32
//go:linkname F_pg_finfo_iso8859_1_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_iso8859_1_to_utf8
func F_pg_finfo_iso8859_1_to_utf8(m *base.Module) int32
//go:linkname F_pg_finfo_utf8_to_iso8859_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_utf8_to_iso8859_2
func F_pg_finfo_utf8_to_iso8859_2(m *base.Module) int32
//go:linkname F_iso8859_1_to_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_iso8859_1_to_utf8
func F_iso8859_1_to_utf8(m *base.Module, l0 int32) int32
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
//go:linkname F_pg_finfo_gbt_bit_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_bit_compress
func F_pg_finfo_gbt_bit_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bit_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_bit_union
func F_pg_finfo_gbt_bit_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bit_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_bit_picksplit
func F_pg_finfo_gbt_bit_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bit_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_bit_consistent
func F_pg_finfo_gbt_bit_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bit_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_bit_penalty
func F_pg_finfo_gbt_bit_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bit_same github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_bit_same
func F_pg_finfo_gbt_bit_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bit_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_bit_sortsupport
func F_pg_finfo_gbt_bit_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_varbit_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_varbit_sortsupport
func F_pg_finfo_gbt_varbit_sortsupport(m *base.Module) int32
//go:linkname F_gbt_bit_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_bit_compress
func F_gbt_bit_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bit_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_bit_consistent
func F_gbt_bit_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bit_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_bit_union
func F_gbt_bit_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bit_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_bit_picksplit
func F_gbt_bit_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bit_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_bit_same
func F_gbt_bit_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bit_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_bit_penalty
func F_gbt_bit_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bit_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_bit_sortsupport
func F_gbt_bit_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_bool_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_bool_compress
func F_pg_finfo_gbt_bool_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bool_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_bool_fetch
func F_pg_finfo_gbt_bool_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bool_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_bool_union
func F_pg_finfo_gbt_bool_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bool_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_bool_picksplit
func F_pg_finfo_gbt_bool_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bool_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_bool_consistent
func F_pg_finfo_gbt_bool_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bool_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_bool_penalty
func F_pg_finfo_gbt_bool_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bool_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_bool_same
func F_pg_finfo_gbt_bool_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bool_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_bool_sortsupport
func F_pg_finfo_gbt_bool_sortsupport(m *base.Module) int32
//go:linkname F_gbt_bool_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_bool_compress
func F_gbt_bool_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bool_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_bool_fetch
func F_gbt_bool_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bool_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_bool_consistent
func F_gbt_bool_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bool_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_bool_union
func F_gbt_bool_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bool_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_bool_penalty
func F_gbt_bool_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bool_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_bool_picksplit
func F_gbt_bool_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bool_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_bool_same
func F_gbt_bool_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bool_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_bool_sortsupport
func F_gbt_bool_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_bytea_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_bytea_compress
func F_pg_finfo_gbt_bytea_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bytea_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_bytea_union
func F_pg_finfo_gbt_bytea_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bytea_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_bytea_picksplit
func F_pg_finfo_gbt_bytea_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bytea_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_bytea_consistent
func F_pg_finfo_gbt_bytea_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bytea_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_bytea_penalty
func F_pg_finfo_gbt_bytea_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bytea_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_bytea_same
func F_pg_finfo_gbt_bytea_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bytea_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_bytea_sortsupport
func F_pg_finfo_gbt_bytea_sortsupport(m *base.Module) int32
//go:linkname F_gbt_bytea_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_bytea_compress
func F_gbt_bytea_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bytea_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_bytea_consistent
func F_gbt_bytea_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bytea_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_bytea_union
func F_gbt_bytea_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bytea_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_bytea_picksplit
func F_gbt_bytea_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bytea_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_bytea_same
func F_gbt_bytea_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bytea_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_bytea_penalty
func F_gbt_bytea_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bytea_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_bytea_sortsupport
func F_gbt_bytea_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_cash_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_cash_compress
func F_pg_finfo_gbt_cash_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_cash_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_cash_fetch
func F_pg_finfo_gbt_cash_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_cash_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_cash_union
func F_pg_finfo_gbt_cash_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_cash_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_cash_picksplit
func F_pg_finfo_gbt_cash_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_cash_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_cash_consistent
func F_pg_finfo_gbt_cash_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_cash_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_cash_distance
func F_pg_finfo_gbt_cash_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_cash_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_cash_penalty
func F_pg_finfo_gbt_cash_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_cash_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_cash_same
func F_pg_finfo_gbt_cash_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_cash_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_cash_sortsupport
func F_pg_finfo_gbt_cash_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_cash_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_cash_dist
func F_pg_finfo_cash_dist(m *base.Module) int32
//go:linkname F_cash_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_cash_dist
func F_cash_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_cash_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_cash_compress
func F_gbt_cash_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_cash_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_cash_fetch
func F_gbt_cash_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_cash_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_cash_consistent
func F_gbt_cash_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_cash_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_cash_distance
func F_gbt_cash_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_cash_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_cash_union
func F_gbt_cash_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_cash_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_cash_penalty
func F_gbt_cash_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_cash_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_cash_picksplit
func F_gbt_cash_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_cash_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_cash_same
func F_gbt_cash_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_cash_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_cash_sortsupport
func F_gbt_cash_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_date_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_date_compress
func F_pg_finfo_gbt_date_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_date_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_date_fetch
func F_pg_finfo_gbt_date_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_date_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_date_union
func F_pg_finfo_gbt_date_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_date_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_date_picksplit
func F_pg_finfo_gbt_date_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_date_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_date_consistent
func F_pg_finfo_gbt_date_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_date_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_date_distance
func F_pg_finfo_gbt_date_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_date_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_date_penalty
func F_pg_finfo_gbt_date_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_date_same github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_date_same
func F_pg_finfo_gbt_date_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_date_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_date_sortsupport
func F_pg_finfo_gbt_date_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_date_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_date_dist
func F_pg_finfo_date_dist(m *base.Module) int32
//go:linkname F_date_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_date_dist
func F_date_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_date_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_date_compress
func F_gbt_date_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_date_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_date_fetch
func F_gbt_date_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_date_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_date_consistent
func F_gbt_date_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_date_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_date_distance
func F_gbt_date_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_date_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_date_union
func F_gbt_date_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_date_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_date_penalty
func F_gbt_date_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_date_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_date_picksplit
func F_gbt_date_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_date_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_date_same
func F_gbt_date_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_date_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_date_sortsupport
func F_gbt_date_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_enum_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_enum_compress
func F_pg_finfo_gbt_enum_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_enum_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_enum_fetch
func F_pg_finfo_gbt_enum_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_enum_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_enum_union
func F_pg_finfo_gbt_enum_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_enum_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_enum_picksplit
func F_pg_finfo_gbt_enum_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_enum_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_enum_consistent
func F_pg_finfo_gbt_enum_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_enum_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_enum_penalty
func F_pg_finfo_gbt_enum_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_enum_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_enum_same
func F_pg_finfo_gbt_enum_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_enum_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_enum_sortsupport
func F_pg_finfo_gbt_enum_sortsupport(m *base.Module) int32
//go:linkname F_gbt_enum_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_enum_compress
func F_gbt_enum_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_enum_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_enum_fetch
func F_gbt_enum_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_enum_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_enum_consistent
func F_gbt_enum_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_enum_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_enum_union
func F_gbt_enum_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_enum_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_enum_penalty
func F_gbt_enum_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_enum_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_enum_picksplit
func F_gbt_enum_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_enum_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_enum_same
func F_gbt_enum_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_enum_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_enum_sortsupport
func F_gbt_enum_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_float4_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_float4_compress
func F_pg_finfo_gbt_float4_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float4_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_float4_fetch
func F_pg_finfo_gbt_float4_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float4_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_float4_union
func F_pg_finfo_gbt_float4_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float4_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_float4_picksplit
func F_pg_finfo_gbt_float4_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float4_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_float4_consistent
func F_pg_finfo_gbt_float4_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float4_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_float4_distance
func F_pg_finfo_gbt_float4_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float4_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_float4_penalty
func F_pg_finfo_gbt_float4_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float4_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_float4_same
func F_pg_finfo_gbt_float4_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float4_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_float4_sortsupport
func F_pg_finfo_gbt_float4_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_float4_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_float4_dist
func F_pg_finfo_float4_dist(m *base.Module) int32
//go:linkname F_float4_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_float4_dist
func F_float4_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float4_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_float4_compress
func F_gbt_float4_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float4_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_float4_fetch
func F_gbt_float4_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float4_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_float4_consistent
func F_gbt_float4_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float4_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_float4_distance
func F_gbt_float4_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float4_union github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_float4_union
func F_gbt_float4_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float4_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_float4_penalty
func F_gbt_float4_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float4_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_float4_picksplit
func F_gbt_float4_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float4_same github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_float4_same
func F_gbt_float4_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float4_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_float4_sortsupport
func F_gbt_float4_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_float8_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_float8_compress
func F_pg_finfo_gbt_float8_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float8_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_float8_fetch
func F_pg_finfo_gbt_float8_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float8_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_float8_union
func F_pg_finfo_gbt_float8_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float8_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_float8_picksplit
func F_pg_finfo_gbt_float8_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float8_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_float8_consistent
func F_pg_finfo_gbt_float8_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float8_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_float8_distance
func F_pg_finfo_gbt_float8_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float8_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_float8_penalty
func F_pg_finfo_gbt_float8_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float8_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_float8_same
func F_pg_finfo_gbt_float8_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_float8_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_float8_sortsupport
func F_pg_finfo_gbt_float8_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_float8_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_float8_dist
func F_pg_finfo_float8_dist(m *base.Module) int32
//go:linkname F_float8_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_float8_dist
func F_float8_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float8_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_float8_compress
func F_gbt_float8_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float8_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_float8_fetch
func F_gbt_float8_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float8_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_float8_consistent
func F_gbt_float8_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float8_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_float8_distance
func F_gbt_float8_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float8_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_float8_union
func F_gbt_float8_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float8_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_float8_penalty
func F_gbt_float8_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float8_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_float8_picksplit
func F_gbt_float8_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float8_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_float8_same
func F_gbt_float8_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_float8_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_float8_sortsupport
func F_gbt_float8_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_btree_gist github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_Pg_magic_func_btree_gist
func F_Pg_magic_func_btree_gist(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_decompress
func F_pg_finfo_gbt_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_gbtreekey_in github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbtreekey_in
func F_pg_finfo_gbtreekey_in(m *base.Module) int32
//go:linkname F_pg_finfo_gbtreekey_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbtreekey_out
func F_pg_finfo_gbtreekey_out(m *base.Module) int32
//go:linkname F_pg_finfo_gist_translate_cmptype_btree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gist_translate_cmptype_btree
func F_pg_finfo_gist_translate_cmptype_btree(m *base.Module) int32
//go:linkname F_gbtreekey_in github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbtreekey_in
func F_gbtreekey_in(m *base.Module, l0 int32) int32
//go:linkname F_gbtreekey_out github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbtreekey_out
func F_gbtreekey_out(m *base.Module, l0 int32) int32
//go:linkname F_gist_translate_cmptype_btree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gist_translate_cmptype_btree
func F_gist_translate_cmptype_btree(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_inet_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_inet_compress
func F_pg_finfo_gbt_inet_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_inet_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_inet_union
func F_pg_finfo_gbt_inet_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_inet_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_inet_picksplit
func F_pg_finfo_gbt_inet_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_inet_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_inet_consistent
func F_pg_finfo_gbt_inet_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_inet_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_inet_penalty
func F_pg_finfo_gbt_inet_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_inet_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_inet_same
func F_pg_finfo_gbt_inet_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_inet_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_inet_sortsupport
func F_pg_finfo_gbt_inet_sortsupport(m *base.Module) int32
//go:linkname F_gbt_inet_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_inet_compress
func F_gbt_inet_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_inet_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_inet_consistent
func F_gbt_inet_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_inet_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_inet_union
func F_gbt_inet_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_inet_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_inet_picksplit
func F_gbt_inet_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_inet_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_inet_same
func F_gbt_inet_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_inet_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_inet_sortsupport
func F_gbt_inet_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_int2_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_int2_compress
func F_pg_finfo_gbt_int2_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int2_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_int2_fetch
func F_pg_finfo_gbt_int2_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int2_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_int2_union
func F_pg_finfo_gbt_int2_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int2_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_int2_picksplit
func F_pg_finfo_gbt_int2_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int2_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_int2_consistent
func F_pg_finfo_gbt_int2_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int2_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_int2_distance
func F_pg_finfo_gbt_int2_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int2_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_int2_penalty
func F_pg_finfo_gbt_int2_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int2_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_int2_same
func F_pg_finfo_gbt_int2_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int2_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_int2_sortsupport
func F_pg_finfo_gbt_int2_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_int2_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_int2_dist
func F_pg_finfo_int2_dist(m *base.Module) int32
//go:linkname F_int2_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_int2_dist
func F_int2_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int2_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_int2_compress
func F_gbt_int2_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int2_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_int2_fetch
func F_gbt_int2_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int2_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_int2_consistent
func F_gbt_int2_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int2_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_int2_distance
func F_gbt_int2_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int2_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_int2_union
func F_gbt_int2_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int2_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_int2_penalty
func F_gbt_int2_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int2_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_int2_picksplit
func F_gbt_int2_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int2_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_int2_same
func F_gbt_int2_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int2_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_int2_sortsupport
func F_gbt_int2_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_int4_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_int4_compress
func F_pg_finfo_gbt_int4_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int4_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_int4_fetch
func F_pg_finfo_gbt_int4_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int4_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_int4_union
func F_pg_finfo_gbt_int4_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int4_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_int4_picksplit
func F_pg_finfo_gbt_int4_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int4_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_int4_consistent
func F_pg_finfo_gbt_int4_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int4_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_int4_distance
func F_pg_finfo_gbt_int4_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int4_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_int4_penalty
func F_pg_finfo_gbt_int4_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int4_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_int4_same
func F_pg_finfo_gbt_int4_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int4_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_int4_sortsupport
func F_pg_finfo_gbt_int4_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_int4_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_int4_dist
func F_pg_finfo_int4_dist(m *base.Module) int32
//go:linkname F_int4_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_int4_dist
func F_int4_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int4_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_int4_compress
func F_gbt_int4_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int4_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_int4_fetch
func F_gbt_int4_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int4_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_int4_consistent
func F_gbt_int4_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int4_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_int4_distance
func F_gbt_int4_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int4_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_int4_union
func F_gbt_int4_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int4_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_int4_penalty
func F_gbt_int4_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int4_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_int4_picksplit
func F_gbt_int4_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int4_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_int4_same
func F_gbt_int4_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int4_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_int4_sortsupport
func F_gbt_int4_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_int8_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_int8_compress
func F_pg_finfo_gbt_int8_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int8_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_int8_fetch
func F_pg_finfo_gbt_int8_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int8_union github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_int8_union
func F_pg_finfo_gbt_int8_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int8_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_int8_picksplit
func F_pg_finfo_gbt_int8_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int8_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_int8_consistent
func F_pg_finfo_gbt_int8_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int8_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_int8_distance
func F_pg_finfo_gbt_int8_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int8_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_int8_penalty
func F_pg_finfo_gbt_int8_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int8_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_int8_same
func F_pg_finfo_gbt_int8_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_int8_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_int8_sortsupport
func F_pg_finfo_gbt_int8_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_int8_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_int8_dist
func F_pg_finfo_int8_dist(m *base.Module) int32
//go:linkname F_int8_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_int8_dist
func F_int8_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int8_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_int8_compress
func F_gbt_int8_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int8_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_int8_fetch
func F_gbt_int8_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int8_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_int8_consistent
func F_gbt_int8_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int8_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_int8_distance
func F_gbt_int8_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int8_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_int8_union
func F_gbt_int8_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int8_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_int8_picksplit
func F_gbt_int8_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int8_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_int8_same
func F_gbt_int8_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_int8_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_int8_sortsupport
func F_gbt_int8_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_intv_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_intv_compress
func F_pg_finfo_gbt_intv_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_intv_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_intv_fetch
func F_pg_finfo_gbt_intv_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_intv_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_intv_decompress
func F_pg_finfo_gbt_intv_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_intv_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_intv_union
func F_pg_finfo_gbt_intv_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_intv_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_intv_picksplit
func F_pg_finfo_gbt_intv_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_intv_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_intv_consistent
func F_pg_finfo_gbt_intv_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_intv_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_intv_distance
func F_pg_finfo_gbt_intv_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_intv_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_intv_penalty
func F_pg_finfo_gbt_intv_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_intv_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_intv_same
func F_pg_finfo_gbt_intv_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_intv_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_intv_sortsupport
func F_pg_finfo_gbt_intv_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_interval_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_interval_dist
func F_pg_finfo_interval_dist(m *base.Module) int32
//go:linkname F_interval_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_interval_dist
func F_interval_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_intv_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_intv_compress
func F_gbt_intv_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_intv_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_intv_fetch
func F_gbt_intv_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_intv_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_intv_consistent
func F_gbt_intv_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_intv_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_intv_distance
func F_gbt_intv_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_intv_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_intv_union
func F_gbt_intv_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_intv_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_intv_penalty
func F_gbt_intv_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_intv_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_intv_picksplit
func F_gbt_intv_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_intv_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_intv_same
func F_gbt_intv_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_intv_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_intv_sortsupport
func F_gbt_intv_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_macad_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_macad_compress
func F_pg_finfo_gbt_macad_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_macad_fetch
func F_pg_finfo_gbt_macad_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_macad_union
func F_pg_finfo_gbt_macad_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_macad_picksplit
func F_pg_finfo_gbt_macad_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_macad_consistent
func F_pg_finfo_gbt_macad_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_macad_penalty
func F_pg_finfo_gbt_macad_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_macad_same
func F_pg_finfo_gbt_macad_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macaddr_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_macaddr_sortsupport
func F_pg_finfo_gbt_macaddr_sortsupport(m *base.Module) int32
//go:linkname F_gbt_macad_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_macad_compress
func F_gbt_macad_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_macad_fetch
func F_gbt_macad_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_macad_consistent
func F_gbt_macad_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_macad_union
func F_gbt_macad_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_macad_penalty
func F_gbt_macad_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_macad_picksplit
func F_gbt_macad_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_macad_same
func F_gbt_macad_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macaddr_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_macaddr_sortsupport
func F_gbt_macaddr_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_macad8_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_macad8_compress
func F_pg_finfo_gbt_macad8_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad8_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_macad8_fetch
func F_pg_finfo_gbt_macad8_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad8_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_macad8_union
func F_pg_finfo_gbt_macad8_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad8_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_macad8_picksplit
func F_pg_finfo_gbt_macad8_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad8_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_macad8_consistent
func F_pg_finfo_gbt_macad8_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad8_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_macad8_penalty
func F_pg_finfo_gbt_macad8_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad8_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_macad8_same
func F_pg_finfo_gbt_macad8_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_macad8_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_macad8_sortsupport
func F_pg_finfo_gbt_macad8_sortsupport(m *base.Module) int32
//go:linkname F_gbt_macad8_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_macad8_compress
func F_gbt_macad8_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad8_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_macad8_fetch
func F_gbt_macad8_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad8_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_macad8_consistent
func F_gbt_macad8_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad8_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_macad8_union
func F_gbt_macad8_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad8_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_macad8_penalty
func F_gbt_macad8_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad8_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_macad8_picksplit
func F_gbt_macad8_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad8_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_macad8_same
func F_gbt_macad8_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_macad8_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_macad8_sortsupport
func F_gbt_macad8_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_numeric_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_numeric_compress
func F_pg_finfo_gbt_numeric_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_numeric_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_numeric_union
func F_pg_finfo_gbt_numeric_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_numeric_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_numeric_picksplit
func F_pg_finfo_gbt_numeric_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_numeric_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_numeric_consistent
func F_pg_finfo_gbt_numeric_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_numeric_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_numeric_penalty
func F_pg_finfo_gbt_numeric_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_numeric_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_numeric_same
func F_pg_finfo_gbt_numeric_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_numeric_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_numeric_sortsupport
func F_pg_finfo_gbt_numeric_sortsupport(m *base.Module) int32
//go:linkname F_gbt_numeric_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_numeric_compress
func F_gbt_numeric_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_numeric_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_numeric_consistent
func F_gbt_numeric_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_numeric_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_numeric_union
func F_gbt_numeric_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_numeric_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_numeric_same
func F_gbt_numeric_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_numeric_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_numeric_penalty
func F_gbt_numeric_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_numeric_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_numeric_picksplit
func F_gbt_numeric_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_numeric_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_numeric_sortsupport
func F_gbt_numeric_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_oid_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_oid_compress
func F_pg_finfo_gbt_oid_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_oid_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_oid_fetch
func F_pg_finfo_gbt_oid_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_oid_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_oid_union
func F_pg_finfo_gbt_oid_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_oid_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_oid_picksplit
func F_pg_finfo_gbt_oid_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_oid_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_oid_consistent
func F_pg_finfo_gbt_oid_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_oid_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_oid_distance
func F_pg_finfo_gbt_oid_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_oid_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_oid_penalty
func F_pg_finfo_gbt_oid_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_oid_same github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_oid_same
func F_pg_finfo_gbt_oid_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_oid_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_oid_sortsupport
func F_pg_finfo_gbt_oid_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_oid_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_oid_dist
func F_pg_finfo_oid_dist(m *base.Module) int32
//go:linkname F_oid_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_oid_dist
func F_oid_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_oid_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_oid_compress
func F_gbt_oid_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_oid_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_oid_fetch
func F_gbt_oid_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_oid_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_oid_consistent
func F_gbt_oid_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_oid_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_oid_distance
func F_gbt_oid_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_oid_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_oid_union
func F_gbt_oid_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_oid_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_oid_picksplit
func F_gbt_oid_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_oid_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_oid_same
func F_gbt_oid_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_oid_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_oid_sortsupport
func F_gbt_oid_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_text_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_text_compress
func F_pg_finfo_gbt_text_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bpchar_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_bpchar_compress
func F_pg_finfo_gbt_bpchar_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_text_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_text_union
func F_pg_finfo_gbt_text_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_text_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_text_picksplit
func F_pg_finfo_gbt_text_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_text_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_text_consistent
func F_pg_finfo_gbt_text_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bpchar_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_bpchar_consistent
func F_pg_finfo_gbt_bpchar_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_text_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_text_penalty
func F_pg_finfo_gbt_text_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_text_same github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_text_same
func F_pg_finfo_gbt_text_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_text_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_text_sortsupport
func F_pg_finfo_gbt_text_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_bpchar_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_bpchar_sortsupport
func F_pg_finfo_gbt_bpchar_sortsupport(m *base.Module) int32
//go:linkname F_gbt_text_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_text_compress
func F_gbt_text_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_text_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_text_consistent
func F_gbt_text_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bpchar_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_bpchar_consistent
func F_gbt_bpchar_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_text_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_text_union
func F_gbt_text_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_text_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_text_picksplit
func F_gbt_text_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_text_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_text_same
func F_gbt_text_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_text_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_text_penalty
func F_gbt_text_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_text_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_text_sortsupport
func F_gbt_text_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_gbt_bpchar_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_bpchar_sortsupport
func F_gbt_bpchar_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_time_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_time_compress
func F_pg_finfo_gbt_time_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_timetz_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_timetz_compress
func F_pg_finfo_gbt_timetz_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_time_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_time_fetch
func F_pg_finfo_gbt_time_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_time_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_time_union
func F_pg_finfo_gbt_time_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_time_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_time_picksplit
func F_pg_finfo_gbt_time_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_time_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_time_consistent
func F_pg_finfo_gbt_time_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_time_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_time_distance
func F_pg_finfo_gbt_time_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_timetz_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_timetz_consistent
func F_pg_finfo_gbt_timetz_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_time_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_time_penalty
func F_pg_finfo_gbt_time_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_time_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_time_same
func F_pg_finfo_gbt_time_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_time_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_time_sortsupport
func F_pg_finfo_gbt_time_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_timetz_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_timetz_sortsupport
func F_pg_finfo_gbt_timetz_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_time_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_time_dist
func F_pg_finfo_time_dist(m *base.Module) int32
//go:linkname F_time_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_time_dist
func F_time_dist(m *base.Module, l0 int32) int32
//go:linkname F_gbt_time_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_time_compress
func F_gbt_time_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_timetz_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_timetz_compress
func F_gbt_timetz_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_time_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_time_fetch
func F_gbt_time_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_time_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_time_consistent
func F_gbt_time_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_time_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_time_distance
func F_gbt_time_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_timetz_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_timetz_consistent
func F_gbt_timetz_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_time_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_time_union
func F_gbt_time_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_time_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_time_penalty
func F_gbt_time_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_time_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_time_picksplit
func F_gbt_time_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_time_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_time_same
func F_gbt_time_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_time_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_time_sortsupport
func F_gbt_time_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_ts_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_ts_compress
func F_pg_finfo_gbt_ts_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_tstz_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_tstz_compress
func F_pg_finfo_gbt_tstz_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_ts_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_ts_fetch
func F_pg_finfo_gbt_ts_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_ts_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_ts_union
func F_pg_finfo_gbt_ts_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_ts_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_ts_picksplit
func F_pg_finfo_gbt_ts_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_ts_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_ts_consistent
func F_pg_finfo_gbt_ts_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_ts_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_ts_distance
func F_pg_finfo_gbt_ts_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_tstz_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_tstz_consistent
func F_pg_finfo_gbt_tstz_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_tstz_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_tstz_distance
func F_pg_finfo_gbt_tstz_distance(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_ts_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_ts_penalty
func F_pg_finfo_gbt_ts_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_ts_same github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_ts_same
func F_pg_finfo_gbt_ts_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_ts_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_ts_sortsupport
func F_pg_finfo_gbt_ts_sortsupport(m *base.Module) int32
//go:linkname F_pg_finfo_ts_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_ts_dist
func F_pg_finfo_ts_dist(m *base.Module) int32
//go:linkname F_ts_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ts_dist
func F_ts_dist(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_tstz_dist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_tstz_dist
func F_pg_finfo_tstz_dist(m *base.Module) int32
//go:linkname F_gbt_ts_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_ts_compress
func F_gbt_ts_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_tstz_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_tstz_compress
func F_gbt_tstz_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_ts_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_ts_fetch
func F_gbt_ts_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_ts_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_ts_consistent
func F_gbt_ts_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_ts_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_ts_distance
func F_gbt_ts_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_tstz_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_tstz_consistent
func F_gbt_tstz_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_tstz_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_tstz_distance
func F_gbt_tstz_distance(m *base.Module, l0 int32) int32
//go:linkname F_gbt_ts_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_ts_union
func F_gbt_ts_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_ts_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_ts_penalty
func F_gbt_ts_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_ts_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_ts_picksplit
func F_gbt_ts_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_ts_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_ts_same
func F_gbt_ts_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_ts_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_ts_sortsupport
func F_gbt_ts_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_var_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_var_decompress
func F_pg_finfo_gbt_var_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_var_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_var_fetch
func F_pg_finfo_gbt_var_fetch(m *base.Module) int32
//go:linkname F_gbt_var_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_var_fetch
func F_gbt_var_fetch(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gbt_uuid_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_uuid_compress
func F_pg_finfo_gbt_uuid_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_uuid_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gbt_uuid_fetch
func F_pg_finfo_gbt_uuid_fetch(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_uuid_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_uuid_union
func F_pg_finfo_gbt_uuid_union(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_uuid_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gbt_uuid_picksplit
func F_pg_finfo_gbt_uuid_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_uuid_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gbt_uuid_consistent
func F_pg_finfo_gbt_uuid_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_uuid_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gbt_uuid_penalty
func F_pg_finfo_gbt_uuid_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_uuid_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gbt_uuid_same
func F_pg_finfo_gbt_uuid_same(m *base.Module) int32
//go:linkname F_pg_finfo_gbt_uuid_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gbt_uuid_sortsupport
func F_pg_finfo_gbt_uuid_sortsupport(m *base.Module) int32
//go:linkname F_gbt_uuid_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_uuid_compress
func F_gbt_uuid_compress(m *base.Module, l0 int32) int32
//go:linkname F_gbt_uuid_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_uuid_fetch
func F_gbt_uuid_fetch(m *base.Module, l0 int32) int32
//go:linkname F_gbt_uuid_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_uuid_consistent
func F_gbt_uuid_consistent(m *base.Module, l0 int32) int32
//go:linkname F_gbt_uuid_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_uuid_union
func F_gbt_uuid_union(m *base.Module, l0 int32) int32
//go:linkname F_gbt_uuid_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_uuid_penalty
func F_gbt_uuid_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gbt_uuid_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_uuid_picksplit
func F_gbt_uuid_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gbt_uuid_same github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_uuid_same
func F_gbt_uuid_same(m *base.Module, l0 int32) int32
//go:linkname F_gbt_uuid_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_uuid_sortsupport
func F_gbt_uuid_sortsupport(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_btree_gin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_btree_gin
func F_Pg_magic_func_btree_gin(m *base.Module) int32
//go:linkname F_pg_finfo_gin_btree_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_btree_consistent
func F_pg_finfo_gin_btree_consistent(m *base.Module) int32
//go:linkname F_gin_btree_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_btree_consistent
func F_gin_btree_consistent(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_extract_value_int2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_extract_value_int2
func F_pg_finfo_gin_extract_value_int2(m *base.Module) int32
//go:linkname F_gin_extract_value_int2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_value_int2
func F_gin_extract_value_int2(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_extract_query_int2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_query_int2
func F_pg_finfo_gin_extract_query_int2(m *base.Module) int32
//go:linkname F_gin_extract_query_int2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_extract_query_int2
func F_gin_extract_query_int2(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_int2 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_compare_prefix_int2
func F_pg_finfo_gin_compare_prefix_int2(m *base.Module) int32
//go:linkname F_gin_compare_prefix_int2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gin_compare_prefix_int2
func F_gin_compare_prefix_int2(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_extract_value_int4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_value_int4
func F_pg_finfo_gin_extract_value_int4(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_int4 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_query_int4
func F_pg_finfo_gin_extract_query_int4(m *base.Module) int32
//go:linkname F_gin_extract_query_int4 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gin_extract_query_int4
func F_gin_extract_query_int4(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_int4 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_compare_prefix_int4
func F_pg_finfo_gin_compare_prefix_int4(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_int8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_value_int8
func F_pg_finfo_gin_extract_value_int8(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_int8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_query_int8
func F_pg_finfo_gin_extract_query_int8(m *base.Module) int32
//go:linkname F_gin_extract_query_int8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_extract_query_int8
func F_gin_extract_query_int8(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_int8 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_compare_prefix_int8
func F_pg_finfo_gin_compare_prefix_int8(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_float4 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_value_float4
func F_pg_finfo_gin_extract_value_float4(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_float4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_query_float4
func F_pg_finfo_gin_extract_query_float4(m *base.Module) int32
//go:linkname F_gin_extract_query_float4 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_extract_query_float4
func F_gin_extract_query_float4(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_float4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_compare_prefix_float4
func F_pg_finfo_gin_compare_prefix_float4(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_float8 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_extract_value_float8
func F_pg_finfo_gin_extract_value_float8(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_float8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_query_float8
func F_pg_finfo_gin_extract_query_float8(m *base.Module) int32
//go:linkname F_gin_extract_query_float8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_extract_query_float8
func F_gin_extract_query_float8(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_float8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_compare_prefix_float8
func F_pg_finfo_gin_compare_prefix_float8(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_money github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_value_money
func F_pg_finfo_gin_extract_value_money(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_money github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_query_money
func F_pg_finfo_gin_extract_query_money(m *base.Module) int32
//go:linkname F_gin_extract_query_money github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_query_money
func F_gin_extract_query_money(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_money github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_compare_prefix_money
func F_pg_finfo_gin_compare_prefix_money(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_value_oid
func F_pg_finfo_gin_extract_value_oid(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_query_oid
func F_pg_finfo_gin_extract_query_oid(m *base.Module) int32
//go:linkname F_gin_extract_query_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_query_oid
func F_gin_extract_query_oid(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_compare_prefix_oid
func F_pg_finfo_gin_compare_prefix_oid(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_timestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_value_timestamp
func F_pg_finfo_gin_extract_value_timestamp(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_timestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_query_timestamp
func F_pg_finfo_gin_extract_query_timestamp(m *base.Module) int32
//go:linkname F_gin_extract_query_timestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_extract_query_timestamp
func F_gin_extract_query_timestamp(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_timestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_compare_prefix_timestamp
func F_pg_finfo_gin_compare_prefix_timestamp(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_timestamptz github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_value_timestamptz
func F_pg_finfo_gin_extract_value_timestamptz(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_timestamptz github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_query_timestamptz
func F_pg_finfo_gin_extract_query_timestamptz(m *base.Module) int32
//go:linkname F_pg_finfo_gin_compare_prefix_timestamptz github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_compare_prefix_timestamptz
func F_pg_finfo_gin_compare_prefix_timestamptz(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_time github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_value_time
func F_pg_finfo_gin_extract_value_time(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_time github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_query_time
func F_pg_finfo_gin_extract_query_time(m *base.Module) int32
//go:linkname F_gin_extract_query_time github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gin_extract_query_time
func F_gin_extract_query_time(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_time github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_compare_prefix_time
func F_pg_finfo_gin_compare_prefix_time(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_timetz github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_value_timetz
func F_pg_finfo_gin_extract_value_timetz(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_timetz github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_query_timetz
func F_pg_finfo_gin_extract_query_timetz(m *base.Module) int32
//go:linkname F_gin_extract_query_timetz github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_extract_query_timetz
func F_gin_extract_query_timetz(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_timetz github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_compare_prefix_timetz
func F_pg_finfo_gin_compare_prefix_timetz(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_date github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_extract_value_date
func F_pg_finfo_gin_extract_value_date(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_date github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_extract_query_date
func F_pg_finfo_gin_extract_query_date(m *base.Module) int32
//go:linkname F_gin_extract_query_date github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_query_date
func F_gin_extract_query_date(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_date github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_compare_prefix_date
func F_pg_finfo_gin_compare_prefix_date(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_interval github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_value_interval
func F_pg_finfo_gin_extract_value_interval(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_interval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_query_interval
func F_pg_finfo_gin_extract_query_interval(m *base.Module) int32
//go:linkname F_gin_extract_query_interval github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_extract_query_interval
func F_gin_extract_query_interval(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_interval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_compare_prefix_interval
func F_pg_finfo_gin_compare_prefix_interval(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_macaddr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_extract_value_macaddr
func F_pg_finfo_gin_extract_value_macaddr(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_macaddr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_query_macaddr
func F_pg_finfo_gin_extract_query_macaddr(m *base.Module) int32
//go:linkname F_gin_extract_query_macaddr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_query_macaddr
func F_gin_extract_query_macaddr(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_macaddr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_compare_prefix_macaddr
func F_pg_finfo_gin_compare_prefix_macaddr(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_macaddr8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_value_macaddr8
func F_pg_finfo_gin_extract_value_macaddr8(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_macaddr8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_query_macaddr8
func F_pg_finfo_gin_extract_query_macaddr8(m *base.Module) int32
//go:linkname F_gin_extract_query_macaddr8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_extract_query_macaddr8
func F_gin_extract_query_macaddr8(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_macaddr8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_compare_prefix_macaddr8
func F_pg_finfo_gin_compare_prefix_macaddr8(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_inet github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_value_inet
func F_pg_finfo_gin_extract_value_inet(m *base.Module) int32
//go:linkname F_gin_extract_value_inet github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gin_extract_value_inet
func F_gin_extract_value_inet(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_extract_query_inet github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_query_inet
func F_pg_finfo_gin_extract_query_inet(m *base.Module) int32
//go:linkname F_gin_extract_query_inet github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_query_inet
func F_gin_extract_query_inet(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_inet github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_compare_prefix_inet
func F_pg_finfo_gin_compare_prefix_inet(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_cidr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_value_cidr
func F_pg_finfo_gin_extract_value_cidr(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_cidr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_query_cidr
func F_pg_finfo_gin_extract_query_cidr(m *base.Module) int32
//go:linkname F_pg_finfo_gin_compare_prefix_cidr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_compare_prefix_cidr
func F_pg_finfo_gin_compare_prefix_cidr(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_text github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_value_text
func F_pg_finfo_gin_extract_value_text(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_text github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_query_text
func F_pg_finfo_gin_extract_query_text(m *base.Module) int32
//go:linkname F_gin_extract_query_text github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_extract_query_text
func F_gin_extract_query_text(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_compare_prefix_text
func F_pg_finfo_gin_compare_prefix_text(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_bpchar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_extract_value_bpchar
func F_pg_finfo_gin_extract_value_bpchar(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_bpchar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_query_bpchar
func F_pg_finfo_gin_extract_query_bpchar(m *base.Module) int32
//go:linkname F_gin_extract_query_bpchar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_query_bpchar
func F_gin_extract_query_bpchar(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_bpchar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_compare_prefix_bpchar
func F_pg_finfo_gin_compare_prefix_bpchar(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_char github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_value_char
func F_pg_finfo_gin_extract_value_char(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_char github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_query_char
func F_pg_finfo_gin_extract_query_char(m *base.Module) int32
//go:linkname F_gin_extract_query_char github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_extract_query_char
func F_gin_extract_query_char(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_char github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_compare_prefix_char
func F_pg_finfo_gin_compare_prefix_char(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_value_bytea
func F_pg_finfo_gin_extract_value_bytea(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_query_bytea
func F_pg_finfo_gin_extract_query_bytea(m *base.Module) int32
//go:linkname F_gin_extract_query_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gin_extract_query_bytea
func F_gin_extract_query_bytea(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_compare_prefix_bytea
func F_pg_finfo_gin_compare_prefix_bytea(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_bit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_value_bit
func F_pg_finfo_gin_extract_value_bit(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_bit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_query_bit
func F_pg_finfo_gin_extract_query_bit(m *base.Module) int32
//go:linkname F_gin_extract_query_bit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_extract_query_bit
func F_gin_extract_query_bit(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_bit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_compare_prefix_bit
func F_pg_finfo_gin_compare_prefix_bit(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_varbit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_value_varbit
func F_pg_finfo_gin_extract_value_varbit(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_varbit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_query_varbit
func F_pg_finfo_gin_extract_query_varbit(m *base.Module) int32
//go:linkname F_gin_extract_query_varbit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gin_extract_query_varbit
func F_gin_extract_query_varbit(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_varbit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_compare_prefix_varbit
func F_pg_finfo_gin_compare_prefix_varbit(m *base.Module) int32
//go:linkname F_pg_finfo_gin_numeric_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_numeric_cmp
func F_pg_finfo_gin_numeric_cmp(m *base.Module) int32
//go:linkname F_gin_numeric_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gin_numeric_cmp
func F_gin_numeric_cmp(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_extract_value_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_extract_value_numeric
func F_pg_finfo_gin_extract_value_numeric(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_extract_query_numeric
func F_pg_finfo_gin_extract_query_numeric(m *base.Module) int32
//go:linkname F_gin_extract_query_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_extract_query_numeric
func F_gin_extract_query_numeric(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_compare_prefix_numeric
func F_pg_finfo_gin_compare_prefix_numeric(m *base.Module) int32
//go:linkname F_pg_finfo_gin_enum_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_enum_cmp
func F_pg_finfo_gin_enum_cmp(m *base.Module) int32
//go:linkname F_gin_enum_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gin_enum_cmp
func F_gin_enum_cmp(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_extract_value_anyenum github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_value_anyenum
func F_pg_finfo_gin_extract_value_anyenum(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_anyenum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_query_anyenum
func F_pg_finfo_gin_extract_query_anyenum(m *base.Module) int32
//go:linkname F_gin_extract_query_anyenum github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gin_extract_query_anyenum
func F_gin_extract_query_anyenum(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_anyenum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_compare_prefix_anyenum
func F_pg_finfo_gin_compare_prefix_anyenum(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_uuid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_extract_value_uuid
func F_pg_finfo_gin_extract_value_uuid(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_uuid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_extract_query_uuid
func F_pg_finfo_gin_extract_query_uuid(m *base.Module) int32
//go:linkname F_gin_extract_query_uuid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_extract_query_uuid
func F_gin_extract_query_uuid(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_uuid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_compare_prefix_uuid
func F_pg_finfo_gin_compare_prefix_uuid(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_extract_value_name
func F_pg_finfo_gin_extract_value_name(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_gin_extract_query_name
func F_pg_finfo_gin_extract_query_name(m *base.Module) int32
//go:linkname F_gin_extract_query_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gin_extract_query_name
func F_gin_extract_query_name(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_compare_prefix_name
func F_pg_finfo_gin_compare_prefix_name(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_value_bool github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gin_extract_value_bool
func F_pg_finfo_gin_extract_value_bool(m *base.Module) int32
//go:linkname F_pg_finfo_gin_extract_query_bool github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gin_extract_query_bool
func F_pg_finfo_gin_extract_query_bool(m *base.Module) int32
//go:linkname F_gin_extract_query_bool github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_extract_query_bool
func F_gin_extract_query_bool(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_compare_prefix_bool github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_compare_prefix_bool
func F_pg_finfo_gin_compare_prefix_bool(m *base.Module) int32
//go:linkname F_Pg_magic_func_unaccent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_unaccent
func F_Pg_magic_func_unaccent(m *base.Module) int32
//go:linkname F_pg_finfo_unaccent_init github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_unaccent_init
func F_pg_finfo_unaccent_init(m *base.Module) int32
//go:linkname F_unaccent_init github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_unaccent_init
func F_unaccent_init(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_unaccent_lexize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_unaccent_lexize
func F_pg_finfo_unaccent_lexize(m *base.Module) int32
//go:linkname F_unaccent_lexize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_unaccent_lexize
func F_unaccent_lexize(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_unaccent_dict github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_unaccent_dict
func F_pg_finfo_unaccent_dict(m *base.Module) int32
//go:linkname F_unaccent_dict github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_unaccent_dict
func F_unaccent_dict(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_tablefunc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Pg_magic_func_tablefunc
func F_Pg_magic_func_tablefunc(m *base.Module) int32
//go:linkname F_pg_finfo_normal_rand github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_normal_rand
func F_pg_finfo_normal_rand(m *base.Module) int32
//go:linkname F_normal_rand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_normal_rand
func F_normal_rand(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_crosstab github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_crosstab
func F_pg_finfo_crosstab(m *base.Module) int32
//go:linkname F_crosstab github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_crosstab
func F_crosstab(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_crosstab_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_crosstab_hash
func F_pg_finfo_crosstab_hash(m *base.Module) int32
//go:linkname F_crosstab_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_crosstab_hash
func F_crosstab_hash(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_connectby_text github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_connectby_text
func F_pg_finfo_connectby_text(m *base.Module) int32
//go:linkname F_connectby_text github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_connectby_text
func F_connectby_text(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_connectby_text_serial github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_connectby_text_serial
func F_pg_finfo_connectby_text_serial(m *base.Module) int32
//go:linkname F_connectby_text_serial github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_connectby_text_serial
func F_connectby_text_serial(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_bqarr_in github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_bqarr_in
func F_pg_finfo_bqarr_in(m *base.Module) int32
//go:linkname F_pg_finfo_bqarr_out github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_bqarr_out
func F_pg_finfo_bqarr_out(m *base.Module) int32
//go:linkname F_pg_finfo_boolop github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_boolop
func F_pg_finfo_boolop(m *base.Module) int32
//go:linkname F_pg_finfo_rboolop github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_rboolop
func F_pg_finfo_rboolop(m *base.Module) int32
//go:linkname F_pg_finfo_querytree github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_querytree
func F_pg_finfo_querytree(m *base.Module) int32
//go:linkname F_rboolop github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_rboolop
func F_rboolop(m *base.Module, l0 int32) int32
//go:linkname F_boolop github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_boolop
func F_boolop(m *base.Module, l0 int32) int32
//go:linkname F_bqarr_in github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bqarr_in
func F_bqarr_in(m *base.Module, l0 int32) int32
//go:linkname F_bqarr_out github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bqarr_out
func F_bqarr_out(m *base.Module, l0 int32) int32
//go:linkname F_querytree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_querytree
func F_querytree(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ginint4_queryextract github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ginint4_queryextract
func F_pg_finfo_ginint4_queryextract(m *base.Module) int32
//go:linkname F_ginint4_queryextract github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginint4_queryextract
func F_ginint4_queryextract(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ginint4_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_ginint4_consistent
func F_pg_finfo_ginint4_consistent(m *base.Module) int32
//go:linkname F_ginint4_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ginint4_consistent
func F_ginint4_consistent(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_g_int_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_g_int_consistent
func F_pg_finfo_g_int_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_g_int_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_g_int_compress
func F_pg_finfo_g_int_compress(m *base.Module) int32
//go:linkname F_pg_finfo_g_int_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_g_int_decompress
func F_pg_finfo_g_int_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_g_int_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_g_int_penalty
func F_pg_finfo_g_int_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_g_int_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_g_int_picksplit
func F_pg_finfo_g_int_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_g_int_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_g_int_union
func F_pg_finfo_g_int_union(m *base.Module) int32
//go:linkname F_pg_finfo_g_int_same github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_g_int_same
func F_pg_finfo_g_int_same(m *base.Module) int32
//go:linkname F_pg_finfo_g_int_options github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_g_int_options
func F_pg_finfo_g_int_options(m *base.Module) int32
//go:linkname F_g_int_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_g_int_consistent
func F_g_int_consistent(m *base.Module, l0 int32) int32
//go:linkname F_g_int_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_g_int_same
func F_g_int_same(m *base.Module, l0 int32) int32
//go:linkname F_g_int_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_g_int_union
func F_g_int_union(m *base.Module, l0 int32) int32
//go:linkname F_g_int_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_g_int_compress
func F_g_int_compress(m *base.Module, l0 int32) int32
//go:linkname F_g_int_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_g_int_decompress
func F_g_int_decompress(m *base.Module, l0 int32) int32
//go:linkname F_g_int_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_g_int_penalty
func F_g_int_penalty(m *base.Module, l0 int32) int32
//go:linkname F_g_int_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_g_int_picksplit
func F_g_int_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_g_int_options github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_g_int_options
func F_g_int_options(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func__int github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func__int
func F_Pg_magic_func__int(m *base.Module) int32
//go:linkname F_pg_finfo__int_different github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo__int_different
func F_pg_finfo__int_different(m *base.Module) int32
//go:linkname F_pg_finfo__int_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo__int_same
func F_pg_finfo__int_same(m *base.Module) int32
//go:linkname F_pg_finfo__int_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo__int_contains
func F_pg_finfo__int_contains(m *base.Module) int32
//go:linkname F_pg_finfo__int_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo__int_contained
func F_pg_finfo__int_contained(m *base.Module) int32
//go:linkname F_pg_finfo__int_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo__int_overlap
func F_pg_finfo__int_overlap(m *base.Module) int32
//go:linkname F_pg_finfo__int_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo__int_union
func F_pg_finfo__int_union(m *base.Module) int32
//go:linkname F_pg_finfo__int_inter github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo__int_inter
func F_pg_finfo__int_inter(m *base.Module) int32
//go:linkname F__int_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__int_contained
func F__int_contained(m *base.Module, l0 int32) int32
//go:linkname F__int_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__int_contains
func F__int_contains(m *base.Module, l0 int32) int32
//go:linkname F__int_different github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__int_different
func F__int_different(m *base.Module, l0 int32) int32
//go:linkname F__int_same github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__int_same
func F__int_same(m *base.Module, l0 int32) int32
//go:linkname F__int_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__int_overlap
func F__int_overlap(m *base.Module, l0 int32) int32
//go:linkname F__int_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__int_union
func F__int_union(m *base.Module, l0 int32) int32
//go:linkname F__int_inter github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__int_inter
func F__int_inter(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_intset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_intset
func F_pg_finfo_intset(m *base.Module) int32
//go:linkname F_pg_finfo_icount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_icount
func F_pg_finfo_icount(m *base.Module) int32
//go:linkname F_pg_finfo_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_sort
func F_pg_finfo_sort(m *base.Module) int32
//go:linkname F_pg_finfo_sort_asc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_sort_asc
func F_pg_finfo_sort_asc(m *base.Module) int32
//go:linkname F_pg_finfo_sort_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_sort_desc
func F_pg_finfo_sort_desc(m *base.Module) int32
//go:linkname F_pg_finfo_uniq github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_uniq
func F_pg_finfo_uniq(m *base.Module) int32
//go:linkname F_pg_finfo_idx github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_idx
func F_pg_finfo_idx(m *base.Module) int32
//go:linkname F_pg_finfo_subarray github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_subarray
func F_pg_finfo_subarray(m *base.Module) int32
//go:linkname F_pg_finfo_intarray_push_elem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_intarray_push_elem
func F_pg_finfo_intarray_push_elem(m *base.Module) int32
//go:linkname F_pg_finfo_intarray_push_array github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_intarray_push_array
func F_pg_finfo_intarray_push_array(m *base.Module) int32
//go:linkname F_pg_finfo_intarray_del_elem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_intarray_del_elem
func F_pg_finfo_intarray_del_elem(m *base.Module) int32
//go:linkname F_pg_finfo_intset_union_elem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_intset_union_elem
func F_pg_finfo_intset_union_elem(m *base.Module) int32
//go:linkname F_pg_finfo_intset_subtract github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_intset_subtract
func F_pg_finfo_intset_subtract(m *base.Module) int32
//go:linkname F_intset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_intset
func F_intset(m *base.Module, l0 int32) int32
//go:linkname F_icount github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_icount
func F_icount(m *base.Module, l0 int32) int32
//go:linkname F_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sort
func F_sort(m *base.Module, l0 int32) int32
//go:linkname F_sort_asc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sort_asc
func F_sort_asc(m *base.Module, l0 int32) int32
//go:linkname F_sort_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sort_desc
func F_sort_desc(m *base.Module, l0 int32) int32
//go:linkname F_uniq github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_uniq
func F_uniq(m *base.Module, l0 int32) int32
//go:linkname F_idx github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_idx
func F_idx(m *base.Module, l0 int32) int32
//go:linkname F_subarray github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_subarray
func F_subarray(m *base.Module, l0 int32) int32
//go:linkname F_intarray_push_elem github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_intarray_push_elem
func F_intarray_push_elem(m *base.Module, l0 int32) int32
//go:linkname F_intarray_push_array github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_intarray_push_array
func F_intarray_push_array(m *base.Module, l0 int32) int32
//go:linkname F_intarray_del_elem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_intarray_del_elem
func F_intarray_del_elem(m *base.Module, l0 int32) int32
//go:linkname F_intset_union_elem github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_intset_union_elem
func F_intset_union_elem(m *base.Module, l0 int32) int32
//go:linkname F_intset_subtract github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_intset_subtract
func F_intset_subtract(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo__int_overlap_sel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo__int_overlap_sel
func F_pg_finfo__int_overlap_sel(m *base.Module) int32
//go:linkname F_pg_finfo__int_contains_sel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo__int_contains_sel
func F_pg_finfo__int_contains_sel(m *base.Module) int32
//go:linkname F_pg_finfo__int_contained_sel github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo__int_contained_sel
func F_pg_finfo__int_contained_sel(m *base.Module) int32
//go:linkname F_pg_finfo__int_overlap_joinsel github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo__int_overlap_joinsel
func F_pg_finfo__int_overlap_joinsel(m *base.Module) int32
//go:linkname F_pg_finfo__int_contains_joinsel github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo__int_contains_joinsel
func F_pg_finfo__int_contains_joinsel(m *base.Module) int32
//go:linkname F_pg_finfo__int_contained_joinsel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo__int_contained_joinsel
func F_pg_finfo__int_contained_joinsel(m *base.Module) int32
//go:linkname F_pg_finfo__int_matchsel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo__int_matchsel
func F_pg_finfo__int_matchsel(m *base.Module) int32
//go:linkname F__int_overlap_sel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__int_overlap_sel
func F__int_overlap_sel(m *base.Module, l0 int32) int32
//go:linkname F__int_contains_sel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__int_contains_sel
func F__int_contains_sel(m *base.Module, l0 int32) int32
//go:linkname F__int_contained_sel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__int_contained_sel
func F__int_contained_sel(m *base.Module, l0 int32) int32
//go:linkname F__int_overlap_joinsel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__int_overlap_joinsel
func F__int_overlap_joinsel(m *base.Module, l0 int32) int32
//go:linkname F__int_contains_joinsel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__int_contains_joinsel
func F__int_contains_joinsel(m *base.Module, l0 int32) int32
//go:linkname F__int_contained_joinsel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__int_contained_joinsel
func F__int_contained_joinsel(m *base.Module, l0 int32) int32
//go:linkname F__int_matchsel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__int_matchsel
func F__int_matchsel(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_g_intbig_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_g_intbig_consistent
func F_pg_finfo_g_intbig_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_g_intbig_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_g_intbig_compress
func F_pg_finfo_g_intbig_compress(m *base.Module) int32
//go:linkname F_pg_finfo_g_intbig_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_g_intbig_decompress
func F_pg_finfo_g_intbig_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_g_intbig_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_g_intbig_penalty
func F_pg_finfo_g_intbig_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_g_intbig_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_g_intbig_picksplit
func F_pg_finfo_g_intbig_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_g_intbig_union github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_g_intbig_union
func F_pg_finfo_g_intbig_union(m *base.Module) int32
//go:linkname F_pg_finfo_g_intbig_same github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_g_intbig_same
func F_pg_finfo_g_intbig_same(m *base.Module) int32
//go:linkname F_pg_finfo_g_intbig_options github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_g_intbig_options
func F_pg_finfo_g_intbig_options(m *base.Module) int32
//go:linkname F_pg_finfo__intbig_in github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo__intbig_in
func F_pg_finfo__intbig_in(m *base.Module) int32
//go:linkname F_pg_finfo__intbig_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo__intbig_out
func F_pg_finfo__intbig_out(m *base.Module) int32
//go:linkname F__intbig_in github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__intbig_in
func F__intbig_in(m *base.Module, l0 int32) int32
//go:linkname F__intbig_out github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__intbig_out
func F__intbig_out(m *base.Module, l0 int32) int32
//go:linkname F_g_intbig_same github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_g_intbig_same
func F_g_intbig_same(m *base.Module, l0 int32) int32
//go:linkname F_g_intbig_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_g_intbig_compress
func F_g_intbig_compress(m *base.Module, l0 int32) int32
//go:linkname F_g_intbig_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_g_intbig_union
func F_g_intbig_union(m *base.Module, l0 int32) int32
//go:linkname F_g_intbig_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_g_intbig_penalty
func F_g_intbig_penalty(m *base.Module, l0 int32) int32
//go:linkname F_g_intbig_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_g_intbig_picksplit
func F_g_intbig_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_g_intbig_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_g_intbig_consistent
func F_g_intbig_consistent(m *base.Module, l0 int32) int32
//go:linkname F_g_intbig_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_g_intbig_options
func F_g_intbig_options(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_daitch_mokotoff github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_daitch_mokotoff
func F_pg_finfo_daitch_mokotoff(m *base.Module) int32
//go:linkname F_daitch_mokotoff github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_daitch_mokotoff
func F_daitch_mokotoff(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_dmetaphone github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_dmetaphone
func F_pg_finfo_dmetaphone(m *base.Module) int32
//go:linkname F_dmetaphone github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dmetaphone
func F_dmetaphone(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_dmetaphone_alt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_dmetaphone_alt
func F_pg_finfo_dmetaphone_alt(m *base.Module) int32
//go:linkname F_dmetaphone_alt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dmetaphone_alt
func F_dmetaphone_alt(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_fuzzystrmatch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_fuzzystrmatch
func F_Pg_magic_func_fuzzystrmatch(m *base.Module) int32
//go:linkname F_pg_finfo_levenshtein_with_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_levenshtein_with_costs
func F_pg_finfo_levenshtein_with_costs(m *base.Module) int32
//go:linkname F_levenshtein_with_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_levenshtein_with_costs
func F_levenshtein_with_costs(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_levenshtein github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_levenshtein
func F_pg_finfo_levenshtein(m *base.Module) int32
//go:linkname F_levenshtein github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_levenshtein
func F_levenshtein(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_levenshtein_less_equal_with_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_levenshtein_less_equal_with_costs
func F_pg_finfo_levenshtein_less_equal_with_costs(m *base.Module) int32
//go:linkname F_levenshtein_less_equal_with_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_levenshtein_less_equal_with_costs
func F_levenshtein_less_equal_with_costs(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_levenshtein_less_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_levenshtein_less_equal
func F_pg_finfo_levenshtein_less_equal(m *base.Module) int32
//go:linkname F_levenshtein_less_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_levenshtein_less_equal
func F_levenshtein_less_equal(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_metaphone github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_metaphone
func F_pg_finfo_metaphone(m *base.Module) int32
//go:linkname F_metaphone github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_metaphone
func F_metaphone(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_soundex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_soundex
func F_pg_finfo_soundex(m *base.Module) int32
//go:linkname F_soundex github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_soundex
func F_soundex(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_difference github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_difference
func F_pg_finfo_difference(m *base.Module) int32
//go:linkname F_difference github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_difference
func F_difference(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_cube github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_cube
func F_Pg_magic_func_cube(m *base.Module) int32
//go:linkname F_pg_finfo_cube_in github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_cube_in
func F_pg_finfo_cube_in(m *base.Module) int32
//go:linkname F_pg_finfo_cube_a_f8_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_cube_a_f8_f8
func F_pg_finfo_cube_a_f8_f8(m *base.Module) int32
//go:linkname F_pg_finfo_cube_a_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_a_f8
func F_pg_finfo_cube_a_f8(m *base.Module) int32
//go:linkname F_pg_finfo_cube_out github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_out
func F_pg_finfo_cube_out(m *base.Module) int32
//go:linkname F_pg_finfo_cube_send github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_send
func F_pg_finfo_cube_send(m *base.Module) int32
//go:linkname F_pg_finfo_cube_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_cube_recv
func F_pg_finfo_cube_recv(m *base.Module) int32
//go:linkname F_pg_finfo_cube_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_cube_f8
func F_pg_finfo_cube_f8(m *base.Module) int32
//go:linkname F_pg_finfo_cube_f8_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_cube_f8_f8
func F_pg_finfo_cube_f8_f8(m *base.Module) int32
//go:linkname F_pg_finfo_cube_c_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_c_f8
func F_pg_finfo_cube_c_f8(m *base.Module) int32
//go:linkname F_pg_finfo_cube_c_f8_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_cube_c_f8_f8
func F_pg_finfo_cube_c_f8_f8(m *base.Module) int32
//go:linkname F_pg_finfo_cube_dim github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_dim
func F_pg_finfo_cube_dim(m *base.Module) int32
//go:linkname F_pg_finfo_cube_ll_coord github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_ll_coord
func F_pg_finfo_cube_ll_coord(m *base.Module) int32
//go:linkname F_pg_finfo_cube_ur_coord github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_cube_ur_coord
func F_pg_finfo_cube_ur_coord(m *base.Module) int32
//go:linkname F_pg_finfo_cube_coord github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_cube_coord
func F_pg_finfo_cube_coord(m *base.Module) int32
//go:linkname F_pg_finfo_cube_coord_llur github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_cube_coord_llur
func F_pg_finfo_cube_coord_llur(m *base.Module) int32
//go:linkname F_pg_finfo_cube_subset github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_subset
func F_pg_finfo_cube_subset(m *base.Module) int32
//go:linkname F_pg_finfo_g_cube_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_g_cube_consistent
func F_pg_finfo_g_cube_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_g_cube_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_g_cube_compress
func F_pg_finfo_g_cube_compress(m *base.Module) int32
//go:linkname F_pg_finfo_g_cube_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_g_cube_decompress
func F_pg_finfo_g_cube_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_g_cube_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_g_cube_penalty
func F_pg_finfo_g_cube_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_g_cube_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_g_cube_picksplit
func F_pg_finfo_g_cube_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_g_cube_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_g_cube_union
func F_pg_finfo_g_cube_union(m *base.Module) int32
//go:linkname F_pg_finfo_g_cube_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_g_cube_same
func F_pg_finfo_g_cube_same(m *base.Module) int32
//go:linkname F_pg_finfo_g_cube_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_g_cube_distance
func F_pg_finfo_g_cube_distance(m *base.Module) int32
//go:linkname F_pg_finfo_cube_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_cube_eq
func F_pg_finfo_cube_eq(m *base.Module) int32
//go:linkname F_pg_finfo_cube_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_cube_ne
func F_pg_finfo_cube_ne(m *base.Module) int32
//go:linkname F_pg_finfo_cube_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_lt
func F_pg_finfo_cube_lt(m *base.Module) int32
//go:linkname F_pg_finfo_cube_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_cube_gt
func F_pg_finfo_cube_gt(m *base.Module) int32
//go:linkname F_pg_finfo_cube_le github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_cube_le
func F_pg_finfo_cube_le(m *base.Module) int32
//go:linkname F_pg_finfo_cube_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_cube_ge
func F_pg_finfo_cube_ge(m *base.Module) int32
//go:linkname F_pg_finfo_cube_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_cube_cmp
func F_pg_finfo_cube_cmp(m *base.Module) int32
//go:linkname F_pg_finfo_cube_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_cube_contains
func F_pg_finfo_cube_contains(m *base.Module) int32
//go:linkname F_pg_finfo_cube_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_cube_contained
func F_pg_finfo_cube_contained(m *base.Module) int32
//go:linkname F_pg_finfo_cube_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_cube_overlap
func F_pg_finfo_cube_overlap(m *base.Module) int32
//go:linkname F_pg_finfo_cube_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_cube_union
func F_pg_finfo_cube_union(m *base.Module) int32
//go:linkname F_pg_finfo_cube_inter github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_inter
func F_pg_finfo_cube_inter(m *base.Module) int32
//go:linkname F_pg_finfo_cube_size github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_cube_size
func F_pg_finfo_cube_size(m *base.Module) int32
//go:linkname F_pg_finfo_distance_taxicab github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_distance_taxicab
func F_pg_finfo_distance_taxicab(m *base.Module) int32
//go:linkname F_pg_finfo_cube_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_cube_distance
func F_pg_finfo_cube_distance(m *base.Module) int32
//go:linkname F_pg_finfo_distance_chebyshev github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_distance_chebyshev
func F_pg_finfo_distance_chebyshev(m *base.Module) int32
//go:linkname F_pg_finfo_cube_is_point github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_cube_is_point
func F_pg_finfo_cube_is_point(m *base.Module) int32
//go:linkname F_pg_finfo_cube_enlarge github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_cube_enlarge
func F_pg_finfo_cube_enlarge(m *base.Module) int32
//go:linkname F_cube_in github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cube_in
func F_cube_in(m *base.Module, l0 int32) int32
//go:linkname F_cube_a_f8_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cube_a_f8_f8
func F_cube_a_f8_f8(m *base.Module, l0 int32) int32
//go:linkname F_cube_a_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cube_a_f8
func F_cube_a_f8(m *base.Module, l0 int32) int32
//go:linkname F_cube_subset github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_cube_subset
func F_cube_subset(m *base.Module, l0 int32) int32
//go:linkname F_cube_out github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_out
func F_cube_out(m *base.Module, l0 int32) int32
//go:linkname F_cube_send github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_send
func F_cube_send(m *base.Module, l0 int32) int32
//go:linkname F_cube_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_recv
func F_cube_recv(m *base.Module, l0 int32) int32
//go:linkname F_g_cube_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_g_cube_consistent
func F_g_cube_consistent(m *base.Module, l0 int32) int32
//go:linkname F_g_cube_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_g_cube_union
func F_g_cube_union(m *base.Module, l0 int32) int32
//go:linkname F_g_cube_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_g_cube_decompress
func F_g_cube_decompress(m *base.Module, l0 int32) int32
//go:linkname F_g_cube_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_g_cube_penalty
func F_g_cube_penalty(m *base.Module, l0 int32) int32
//go:linkname F_g_cube_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_g_cube_picksplit
func F_g_cube_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_cube_inter github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_inter
func F_cube_inter(m *base.Module, l0 int32) int32
//go:linkname F_g_cube_same github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_g_cube_same
func F_g_cube_same(m *base.Module, l0 int32) int32
//go:linkname F_cube_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cube_union
func F_cube_union(m *base.Module, l0 int32) int32
//go:linkname F_cube_size github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cube_size
func F_cube_size(m *base.Module, l0 int32) int32
//go:linkname F_cube_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cube_cmp
func F_cube_cmp(m *base.Module, l0 int32) int32
//go:linkname F_cube_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_eq
func F_cube_eq(m *base.Module, l0 int32) int32
//go:linkname F_cube_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cube_ne
func F_cube_ne(m *base.Module, l0 int32) int32
//go:linkname F_cube_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_cube_lt
func F_cube_lt(m *base.Module, l0 int32) int32
//go:linkname F_cube_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cube_gt
func F_cube_gt(m *base.Module, l0 int32) int32
//go:linkname F_cube_le github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cube_le
func F_cube_le(m *base.Module, l0 int32) int32
//go:linkname F_cube_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cube_ge
func F_cube_ge(m *base.Module, l0 int32) int32
//go:linkname F_cube_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cube_contains
func F_cube_contains(m *base.Module, l0 int32) int32
//go:linkname F_cube_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cube_contained
func F_cube_contained(m *base.Module, l0 int32) int32
//go:linkname F_cube_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cube_overlap
func F_cube_overlap(m *base.Module, l0 int32) int32
//go:linkname F_cube_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cube_distance
func F_cube_distance(m *base.Module, l0 int32) int32
//go:linkname F_distance_taxicab github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_distance_taxicab
func F_distance_taxicab(m *base.Module, l0 int32) int32
//go:linkname F_distance_chebyshev github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_distance_chebyshev
func F_distance_chebyshev(m *base.Module, l0 int32) int32
//go:linkname F_g_cube_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_g_cube_distance
func F_g_cube_distance(m *base.Module, l0 int32) int32
//go:linkname F_cube_is_point github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cube_is_point
func F_cube_is_point(m *base.Module, l0 int32) int32
//go:linkname F_cube_dim github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_dim
func F_cube_dim(m *base.Module, l0 int32) int32
//go:linkname F_cube_ll_coord github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_ll_coord
func F_cube_ll_coord(m *base.Module, l0 int32) int32
//go:linkname F_cube_ur_coord github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cube_ur_coord
func F_cube_ur_coord(m *base.Module, l0 int32) int32
//go:linkname F_cube_coord github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cube_coord
func F_cube_coord(m *base.Module, l0 int32) int32
//go:linkname F_cube_coord_llur github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cube_coord_llur
func F_cube_coord_llur(m *base.Module, l0 int32) int32
//go:linkname F_cube_enlarge github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_cube_enlarge
func F_cube_enlarge(m *base.Module, l0 int32) int32
//go:linkname F_cube_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cube_f8
func F_cube_f8(m *base.Module, l0 int32) int32
//go:linkname F_cube_f8_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cube_f8_f8
func F_cube_f8_f8(m *base.Module, l0 int32) int32
//go:linkname F_cube_c_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cube_c_f8
func F_cube_c_f8(m *base.Module, l0 int32) int32
//go:linkname F_cube_c_f8_f8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cube_c_f8_f8
func F_cube_c_f8_f8(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_earthdistance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_earthdistance
func F_Pg_magic_func_earthdistance(m *base.Module) int32
//go:linkname F_pg_finfo_geo_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_geo_distance
func F_pg_finfo_geo_distance(m *base.Module) int32
//go:linkname F_geo_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_geo_distance
func F_geo_distance(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_seg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_seg
func F_Pg_magic_func_seg(m *base.Module) int32
//go:linkname F_pg_finfo_seg_in github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_seg_in
func F_pg_finfo_seg_in(m *base.Module) int32
//go:linkname F_pg_finfo_seg_out github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_seg_out
func F_pg_finfo_seg_out(m *base.Module) int32
//go:linkname F_pg_finfo_seg_size github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_seg_size
func F_pg_finfo_seg_size(m *base.Module) int32
//go:linkname F_pg_finfo_seg_lower github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_seg_lower
func F_pg_finfo_seg_lower(m *base.Module) int32
//go:linkname F_pg_finfo_seg_upper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_seg_upper
func F_pg_finfo_seg_upper(m *base.Module) int32
//go:linkname F_pg_finfo_seg_center github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_seg_center
func F_pg_finfo_seg_center(m *base.Module) int32
//go:linkname F_pg_finfo_gseg_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gseg_consistent
func F_pg_finfo_gseg_consistent(m *base.Module) int32
//go:linkname F_pg_finfo_gseg_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gseg_compress
func F_pg_finfo_gseg_compress(m *base.Module) int32
//go:linkname F_pg_finfo_gseg_decompress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gseg_decompress
func F_pg_finfo_gseg_decompress(m *base.Module) int32
//go:linkname F_pg_finfo_gseg_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gseg_picksplit
func F_pg_finfo_gseg_picksplit(m *base.Module) int32
//go:linkname F_pg_finfo_gseg_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gseg_penalty
func F_pg_finfo_gseg_penalty(m *base.Module) int32
//go:linkname F_pg_finfo_gseg_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_gseg_union
func F_pg_finfo_gseg_union(m *base.Module) int32
//go:linkname F_pg_finfo_gseg_same github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_gseg_same
func F_pg_finfo_gseg_same(m *base.Module) int32
//go:linkname F_pg_finfo_seg_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_seg_same
func F_pg_finfo_seg_same(m *base.Module) int32
//go:linkname F_pg_finfo_seg_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_seg_contains
func F_pg_finfo_seg_contains(m *base.Module) int32
//go:linkname F_pg_finfo_seg_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_seg_contained
func F_pg_finfo_seg_contained(m *base.Module) int32
//go:linkname F_pg_finfo_seg_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_seg_overlap
func F_pg_finfo_seg_overlap(m *base.Module) int32
//go:linkname F_pg_finfo_seg_left github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_seg_left
func F_pg_finfo_seg_left(m *base.Module) int32
//go:linkname F_pg_finfo_seg_over_left github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_seg_over_left
func F_pg_finfo_seg_over_left(m *base.Module) int32
//go:linkname F_pg_finfo_seg_right github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_seg_right
func F_pg_finfo_seg_right(m *base.Module) int32
//go:linkname F_pg_finfo_seg_over_right github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_seg_over_right
func F_pg_finfo_seg_over_right(m *base.Module) int32
//go:linkname F_pg_finfo_seg_union github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_seg_union
func F_pg_finfo_seg_union(m *base.Module) int32
//go:linkname F_pg_finfo_seg_inter github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_seg_inter
func F_pg_finfo_seg_inter(m *base.Module) int32
//go:linkname F_pg_finfo_seg_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_seg_cmp
func F_pg_finfo_seg_cmp(m *base.Module) int32
//go:linkname F_pg_finfo_seg_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_seg_lt
func F_pg_finfo_seg_lt(m *base.Module) int32
//go:linkname F_pg_finfo_seg_le github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_seg_le
func F_pg_finfo_seg_le(m *base.Module) int32
//go:linkname F_pg_finfo_seg_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_seg_gt
func F_pg_finfo_seg_gt(m *base.Module) int32
//go:linkname F_pg_finfo_seg_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_seg_ge
func F_pg_finfo_seg_ge(m *base.Module) int32
//go:linkname F_pg_finfo_seg_different github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_seg_different
func F_pg_finfo_seg_different(m *base.Module) int32
//go:linkname F_seg_in github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_seg_in
func F_seg_in(m *base.Module, l0 int32) int32
//go:linkname F_seg_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_seg_out
func F_seg_out(m *base.Module, l0 int32) int32
//go:linkname F_seg_center github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_seg_center
func F_seg_center(m *base.Module, l0 int32) int32
//go:linkname F_gseg_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gseg_consistent
func F_gseg_consistent(m *base.Module, l0 int32) int32
//go:linkname F_seg_over_right github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_seg_over_right
func F_seg_over_right(m *base.Module, l0 int32) int32
//go:linkname F_seg_right github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_seg_right
func F_seg_right(m *base.Module, l0 int32) int32
//go:linkname F_seg_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_seg_overlap
func F_seg_overlap(m *base.Module, l0 int32) int32
//go:linkname F_seg_left github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_seg_left
func F_seg_left(m *base.Module, l0 int32) int32
//go:linkname F_seg_over_left github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_seg_over_left
func F_seg_over_left(m *base.Module, l0 int32) int32
//go:linkname F_seg_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_seg_contains
func F_seg_contains(m *base.Module, l0 int32) int32
//go:linkname F_gseg_union github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gseg_union
func F_gseg_union(m *base.Module, l0 int32) int32
//go:linkname F_seg_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_seg_union
func F_seg_union(m *base.Module, l0 int32) int32
//go:linkname F_gseg_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gseg_penalty
func F_gseg_penalty(m *base.Module, l0 int32) int32
//go:linkname F_gseg_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gseg_picksplit
func F_gseg_picksplit(m *base.Module, l0 int32) int32
//go:linkname F_gseg_same github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gseg_same
func F_gseg_same(m *base.Module, l0 int32) int32
//go:linkname F_seg_same github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_seg_same
func F_seg_same(m *base.Module, l0 int32) int32
//go:linkname F_seg_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_seg_cmp
func F_seg_cmp(m *base.Module, l0 int32) int32
//go:linkname F_seg_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_seg_contained
func F_seg_contained(m *base.Module, l0 int32) int32
//go:linkname F_seg_inter github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_seg_inter
func F_seg_inter(m *base.Module, l0 int32) int32
//go:linkname F_seg_size github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_seg_size
func F_seg_size(m *base.Module, l0 int32) int32
//go:linkname F_seg_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_seg_lt
func F_seg_lt(m *base.Module, l0 int32) int32
//go:linkname F_seg_le github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_seg_le
func F_seg_le(m *base.Module, l0 int32) int32
//go:linkname F_seg_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_seg_gt
func F_seg_gt(m *base.Module, l0 int32) int32
//go:linkname F_seg_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_seg_ge
func F_seg_ge(m *base.Module, l0 int32) int32
//go:linkname F_seg_different github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_seg_different
func F_seg_different(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_bloom github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_bloom
func F_Pg_magic_func_bloom(m *base.Module) int32
//go:linkname F_pg_finfo_blhandler github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_blhandler
func F_pg_finfo_blhandler(m *base.Module) int32
//go:linkname F__PG_init_bloom github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__PG_init_bloom
func F__PG_init_bloom(m *base.Module)
//go:linkname F_blhandler github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_blhandler
func F_blhandler(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_isn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_Pg_magic_func_isn
func F_Pg_magic_func_isn(m *base.Module) int32
//go:linkname F__PG_init_isn github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__PG_init_isn
func F__PG_init_isn(m *base.Module)
//go:linkname F_pg_finfo_isn_out github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_isn_out
func F_pg_finfo_isn_out(m *base.Module) int32
//go:linkname F_isn_out github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_isn_out
func F_isn_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ean13_out github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_ean13_out
func F_pg_finfo_ean13_out(m *base.Module) int32
//go:linkname F_ean13_out github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ean13_out
func F_ean13_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ean13_in github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_ean13_in
func F_pg_finfo_ean13_in(m *base.Module) int32
//go:linkname F_ean13_in github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ean13_in
func F_ean13_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_isbn_in github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_isbn_in
func F_pg_finfo_isbn_in(m *base.Module) int32
//go:linkname F_isbn_in github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_isbn_in
func F_isbn_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ismn_in github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ismn_in
func F_pg_finfo_ismn_in(m *base.Module) int32
//go:linkname F_ismn_in github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ismn_in
func F_ismn_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_issn_in github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_issn_in
func F_pg_finfo_issn_in(m *base.Module) int32
//go:linkname F_issn_in github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_issn_in
func F_issn_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_upc_in github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_upc_in
func F_pg_finfo_upc_in(m *base.Module) int32
//go:linkname F_upc_in github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_upc_in
func F_upc_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_isbn_cast_from_ean13 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_isbn_cast_from_ean13
func F_pg_finfo_isbn_cast_from_ean13(m *base.Module) int32
//go:linkname F_isbn_cast_from_ean13 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_isbn_cast_from_ean13
func F_isbn_cast_from_ean13(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ismn_cast_from_ean13 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_ismn_cast_from_ean13
func F_pg_finfo_ismn_cast_from_ean13(m *base.Module) int32
//go:linkname F_ismn_cast_from_ean13 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ismn_cast_from_ean13
func F_ismn_cast_from_ean13(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_issn_cast_from_ean13 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_issn_cast_from_ean13
func F_pg_finfo_issn_cast_from_ean13(m *base.Module) int32
//go:linkname F_issn_cast_from_ean13 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_issn_cast_from_ean13
func F_issn_cast_from_ean13(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_upc_cast_from_ean13 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_upc_cast_from_ean13
func F_pg_finfo_upc_cast_from_ean13(m *base.Module) int32
//go:linkname F_upc_cast_from_ean13 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_upc_cast_from_ean13
func F_upc_cast_from_ean13(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_is_valid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_is_valid
func F_pg_finfo_is_valid(m *base.Module) int32
//go:linkname F_is_valid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_is_valid
func F_is_valid(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_make_valid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_make_valid
func F_pg_finfo_make_valid(m *base.Module) int32
//go:linkname F_make_valid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_valid
func F_make_valid(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_accept_weak_input github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_accept_weak_input
func F_pg_finfo_accept_weak_input(m *base.Module) int32
//go:linkname F_accept_weak_input github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_accept_weak_input
func F_accept_weak_input(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_weak_input_status github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_weak_input_status
func F_pg_finfo_weak_input_status(m *base.Module) int32
//go:linkname F_weak_input_status github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_weak_input_status
func F_weak_input_status(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_dict_int github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Pg_magic_func_dict_int
func F_Pg_magic_func_dict_int(m *base.Module) int32
//go:linkname F_pg_finfo_dintdict_init github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_dintdict_init
func F_pg_finfo_dintdict_init(m *base.Module) int32
//go:linkname F_pg_finfo_dintdict_lexize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_dintdict_lexize
func F_pg_finfo_dintdict_lexize(m *base.Module) int32
//go:linkname F_dintdict_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dintdict_init
func F_dintdict_init(m *base.Module, l0 int32) int32
//go:linkname F_dintdict_lexize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dintdict_lexize
func F_dintdict_lexize(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_dict_xsyn github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_dict_xsyn
func F_Pg_magic_func_dict_xsyn(m *base.Module) int32
//go:linkname F_pg_finfo_dxsyn_init github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_dxsyn_init
func F_pg_finfo_dxsyn_init(m *base.Module) int32
//go:linkname F_pg_finfo_dxsyn_lexize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_dxsyn_lexize
func F_pg_finfo_dxsyn_lexize(m *base.Module) int32
//go:linkname F_dxsyn_init github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dxsyn_init
func F_dxsyn_init(m *base.Module, l0 int32) int32
//go:linkname F_dxsyn_lexize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dxsyn_lexize
func F_dxsyn_lexize(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_lo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Pg_magic_func_lo
func F_Pg_magic_func_lo(m *base.Module) int32
//go:linkname F_pg_finfo_lo_manage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_lo_manage
func F_pg_finfo_lo_manage(m *base.Module) int32
//go:linkname F_lo_manage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lo_manage
func F_lo_manage(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_tsm_system_rows github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_tsm_system_rows
func F_Pg_magic_func_tsm_system_rows(m *base.Module) int32
//go:linkname F_pg_finfo_tsm_system_rows_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_tsm_system_rows_handler
func F_pg_finfo_tsm_system_rows_handler(m *base.Module) int32
//go:linkname F_tsm_system_rows_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tsm_system_rows_handler
func F_tsm_system_rows_handler(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_tsm_system_time github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Pg_magic_func_tsm_system_time
func F_Pg_magic_func_tsm_system_time(m *base.Module) int32
//go:linkname F_pg_finfo_tsm_system_time_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_tsm_system_time_handler
func F_pg_finfo_tsm_system_time_handler(m *base.Module) int32
//go:linkname F_tsm_system_time_handler github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tsm_system_time_handler
func F_tsm_system_time_handler(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pgstattuple_approx github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_pgstattuple_approx
func F_pg_finfo_pgstattuple_approx(m *base.Module) int32
//go:linkname F_pg_finfo_pgstattuple_approx_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_pgstattuple_approx_v1_5
func F_pg_finfo_pgstattuple_approx_v1_5(m *base.Module) int32
//go:linkname F_pgstattuple_approx github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstattuple_approx
func F_pgstattuple_approx(m *base.Module, l0 int32) int32
//go:linkname F_pgstattuple_approx_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstattuple_approx_v1_5
func F_pgstattuple_approx_v1_5(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_pgstatindex github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pgstatindex
func F_pg_finfo_pgstatindex(m *base.Module) int32
//go:linkname F_pg_finfo_pgstatindexbyid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pgstatindexbyid
func F_pg_finfo_pgstatindexbyid(m *base.Module) int32
//go:linkname F_pg_finfo_pg_relpages github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pg_relpages
func F_pg_finfo_pg_relpages(m *base.Module) int32
//go:linkname F_pg_finfo_pg_relpagesbyid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_relpagesbyid
func F_pg_finfo_pg_relpagesbyid(m *base.Module) int32
//go:linkname F_pg_finfo_pgstatginindex github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pgstatginindex
func F_pg_finfo_pgstatginindex(m *base.Module) int32
//go:linkname F_pg_finfo_pgstathashindex github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pgstathashindex
func F_pg_finfo_pgstathashindex(m *base.Module) int32
//go:linkname F_pg_finfo_pgstatindex_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pgstatindex_v1_5
func F_pg_finfo_pgstatindex_v1_5(m *base.Module) int32
//go:linkname F_pg_finfo_pgstatindexbyid_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pgstatindexbyid_v1_5
func F_pg_finfo_pgstatindexbyid_v1_5(m *base.Module) int32
//go:linkname F_pg_finfo_pg_relpages_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pg_relpages_v1_5
func F_pg_finfo_pg_relpages_v1_5(m *base.Module) int32
//go:linkname F_pg_finfo_pg_relpagesbyid_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_relpagesbyid_v1_5
func F_pg_finfo_pg_relpagesbyid_v1_5(m *base.Module) int32
//go:linkname F_pg_finfo_pgstatginindex_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pgstatginindex_v1_5
func F_pg_finfo_pgstatginindex_v1_5(m *base.Module) int32
//go:linkname F_pgstatindex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstatindex
func F_pgstatindex(m *base.Module, l0 int32) int32
//go:linkname F_pgstatindex_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstatindex_v1_5
func F_pgstatindex_v1_5(m *base.Module, l0 int32) int32
//go:linkname F_pgstatindexbyid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstatindexbyid
func F_pgstatindexbyid(m *base.Module, l0 int32) int32
//go:linkname F_pgstatindexbyid_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstatindexbyid_v1_5
func F_pgstatindexbyid_v1_5(m *base.Module, l0 int32) int32
//go:linkname F_pg_relpages github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_relpages
func F_pg_relpages(m *base.Module, l0 int32) int32
//go:linkname F_pg_relpages_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_relpages_v1_5
func F_pg_relpages_v1_5(m *base.Module, l0 int32) int32
//go:linkname F_pg_relpagesbyid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_relpagesbyid
func F_pg_relpagesbyid(m *base.Module, l0 int32) int32
//go:linkname F_pg_relpagesbyid_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_relpagesbyid_v1_5
func F_pg_relpagesbyid_v1_5(m *base.Module, l0 int32) int32
//go:linkname F_pgstatginindex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstatginindex
func F_pgstatginindex(m *base.Module, l0 int32) int32
//go:linkname F_pgstatginindex_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstatginindex_v1_5
func F_pgstatginindex_v1_5(m *base.Module, l0 int32) int32
//go:linkname F_pgstathashindex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstathashindex
func F_pgstathashindex(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_pgstattuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_pgstattuple
func F_Pg_magic_func_pgstattuple(m *base.Module) int32
//go:linkname F_pg_finfo_pgstattuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pgstattuple
func F_pg_finfo_pgstattuple(m *base.Module) int32
//go:linkname F_pg_finfo_pgstattuple_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pgstattuple_v1_5
func F_pg_finfo_pgstattuple_v1_5(m *base.Module) int32
//go:linkname F_pg_finfo_pgstattuplebyid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pgstattuplebyid
func F_pg_finfo_pgstattuplebyid(m *base.Module) int32
//go:linkname F_pg_finfo_pgstattuplebyid_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pgstattuplebyid_v1_5
func F_pg_finfo_pgstattuplebyid_v1_5(m *base.Module) int32
//go:linkname F_pgstattuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstattuple
func F_pgstattuple(m *base.Module, l0 int32) int32
//go:linkname F_pgstattuple_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstattuple_v1_5
func F_pgstattuple_v1_5(m *base.Module, l0 int32) int32
//go:linkname F_pgstattuplebyid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstattuplebyid
func F_pgstattuplebyid(m *base.Module, l0 int32) int32
//go:linkname F_pgstattuplebyid_v1_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstattuplebyid_v1_5
func F_pgstattuplebyid_v1_5(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_uuid_ossp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_uuid_ossp
func F_Pg_magic_func_uuid_ossp(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_nil github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_uuid_nil
func F_pg_finfo_uuid_nil(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_ns_dns github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_uuid_ns_dns
func F_pg_finfo_uuid_ns_dns(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_ns_url github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_uuid_ns_url
func F_pg_finfo_uuid_ns_url(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_ns_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_uuid_ns_oid
func F_pg_finfo_uuid_ns_oid(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_ns_x500 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_uuid_ns_x500
func F_pg_finfo_uuid_ns_x500(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_generate_v1 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_uuid_generate_v1
func F_pg_finfo_uuid_generate_v1(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_generate_v1mc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_uuid_generate_v1mc
func F_pg_finfo_uuid_generate_v1mc(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_generate_v3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_uuid_generate_v3
func F_pg_finfo_uuid_generate_v3(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_generate_v4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_uuid_generate_v4
func F_pg_finfo_uuid_generate_v4(m *base.Module) int32
//go:linkname F_pg_finfo_uuid_generate_v5 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_uuid_generate_v5
func F_pg_finfo_uuid_generate_v5(m *base.Module) int32
//go:linkname F_uuid_nil github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_uuid_nil
func F_uuid_nil(m *base.Module, l0 int32) int32
//go:linkname F_uuid_ns_dns github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_uuid_ns_dns
func F_uuid_ns_dns(m *base.Module, l0 int32) int32
//go:linkname F_uuid_ns_url github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_uuid_ns_url
func F_uuid_ns_url(m *base.Module, l0 int32) int32
//go:linkname F_uuid_ns_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_uuid_ns_oid
func F_uuid_ns_oid(m *base.Module, l0 int32) int32
//go:linkname F_uuid_ns_x500 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_uuid_ns_x500
func F_uuid_ns_x500(m *base.Module, l0 int32) int32
//go:linkname F_uuid_generate_v1 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_uuid_generate_v1
func F_uuid_generate_v1(m *base.Module, l0 int32) int32
//go:linkname F_uuid_generate_v1mc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_uuid_generate_v1mc
func F_uuid_generate_v1mc(m *base.Module, l0 int32) int32
//go:linkname F_uuid_generate_v3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_uuid_generate_v3
func F_uuid_generate_v3(m *base.Module, l0 int32) int32
//go:linkname F_uuid_generate_v4 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_uuid_generate_v4
func F_uuid_generate_v4(m *base.Module, l0 int32) int32
//go:linkname F_uuid_generate_v5 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_uuid_generate_v5
func F_uuid_generate_v5(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_index_check github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gin_index_check
func F_pg_finfo_gin_index_check(m *base.Module) int32
//go:linkname F_gin_index_check github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gin_index_check
func F_gin_index_check(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_verify_heapam github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_verify_heapam
func F_pg_finfo_verify_heapam(m *base.Module) int32
//go:linkname F_verify_heapam github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_verify_heapam
func F_verify_heapam(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_amcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_amcheck
func F_Pg_magic_func_amcheck(m *base.Module) int32
//go:linkname F_pg_finfo_bt_index_check github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_bt_index_check
func F_pg_finfo_bt_index_check(m *base.Module) int32
//go:linkname F_pg_finfo_bt_index_parent_check github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_bt_index_parent_check
func F_pg_finfo_bt_index_parent_check(m *base.Module) int32
//go:linkname F_bt_index_check github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bt_index_check
func F_bt_index_check(m *base.Module, l0 int32) int32
//go:linkname F_bt_index_parent_check github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bt_index_parent_check
func F_bt_index_parent_check(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_pg_visibility github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_Pg_magic_func_pg_visibility
func F_Pg_magic_func_pg_visibility(m *base.Module) int32
//go:linkname F_pg_finfo_pg_visibility_map github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_visibility_map
func F_pg_finfo_pg_visibility_map(m *base.Module) int32
//go:linkname F_pg_finfo_pg_visibility_map_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_visibility_map_rel
func F_pg_finfo_pg_visibility_map_rel(m *base.Module) int32
//go:linkname F_pg_finfo_pg_visibility github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_pg_visibility
func F_pg_finfo_pg_visibility(m *base.Module) int32
//go:linkname F_pg_finfo_pg_visibility_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_pg_visibility_rel
func F_pg_finfo_pg_visibility_rel(m *base.Module) int32
//go:linkname F_pg_finfo_pg_visibility_map_summary github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_pg_visibility_map_summary
func F_pg_finfo_pg_visibility_map_summary(m *base.Module) int32
//go:linkname F_pg_finfo_pg_check_frozen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_check_frozen
func F_pg_finfo_pg_check_frozen(m *base.Module) int32
//go:linkname F_pg_finfo_pg_check_visible github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_check_visible
func F_pg_finfo_pg_check_visible(m *base.Module) int32
//go:linkname F_pg_finfo_pg_truncate_visibility_map github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_pg_truncate_visibility_map
func F_pg_finfo_pg_truncate_visibility_map(m *base.Module) int32
//go:linkname F_pg_visibility_map github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_visibility_map
func F_pg_visibility_map(m *base.Module, l0 int32) int32
//go:linkname F_pg_visibility github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_visibility
func F_pg_visibility(m *base.Module, l0 int32) int32
//go:linkname F_pg_visibility_map_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_visibility_map_rel
func F_pg_visibility_map_rel(m *base.Module, l0 int32) int32
//go:linkname F_pg_visibility_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_visibility_rel
func F_pg_visibility_rel(m *base.Module, l0 int32) int32
//go:linkname F_pg_visibility_map_summary github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_visibility_map_summary
func F_pg_visibility_map_summary(m *base.Module, l0 int32) int32
//go:linkname F_pg_check_frozen github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_check_frozen
func F_pg_check_frozen(m *base.Module, l0 int32) int32
//go:linkname F_pg_check_visible github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_check_visible
func F_pg_check_visible(m *base.Module, l0 int32) int32
//go:linkname F_pg_truncate_visibility_map github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_truncate_visibility_map
func F_pg_truncate_visibility_map(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_brin_page_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_brin_page_type
func F_pg_finfo_brin_page_type(m *base.Module) int32
//go:linkname F_pg_finfo_brin_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_brin_page_items
func F_pg_finfo_brin_page_items(m *base.Module) int32
//go:linkname F_pg_finfo_brin_metapage_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_brin_metapage_info
func F_pg_finfo_brin_metapage_info(m *base.Module) int32
//go:linkname F_pg_finfo_brin_revmap_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_brin_revmap_data
func F_pg_finfo_brin_revmap_data(m *base.Module) int32
//go:linkname F_brin_page_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_brin_page_type
func F_brin_page_type(m *base.Module, l0 int32) int32
//go:linkname F_brin_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_brin_page_items
func F_brin_page_items(m *base.Module, l0 int32) int32
//go:linkname F_brin_metapage_info github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_brin_metapage_info
func F_brin_metapage_info(m *base.Module, l0 int32) int32
//go:linkname F_brin_revmap_data github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_brin_revmap_data
func F_brin_revmap_data(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_bt_metap github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_bt_metap
func F_pg_finfo_bt_metap(m *base.Module) int32
//go:linkname F_pg_finfo_bt_page_items_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_bt_page_items_1_9
func F_pg_finfo_bt_page_items_1_9(m *base.Module) int32
//go:linkname F_pg_finfo_bt_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_bt_page_items
func F_pg_finfo_bt_page_items(m *base.Module) int32
//go:linkname F_pg_finfo_bt_page_items_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_bt_page_items_bytea
func F_pg_finfo_bt_page_items_bytea(m *base.Module) int32
//go:linkname F_pg_finfo_bt_page_stats_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_bt_page_stats_1_9
func F_pg_finfo_bt_page_stats_1_9(m *base.Module) int32
//go:linkname F_pg_finfo_bt_page_stats github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_bt_page_stats
func F_pg_finfo_bt_page_stats(m *base.Module) int32
//go:linkname F_pg_finfo_bt_multi_page_stats github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_bt_multi_page_stats
func F_pg_finfo_bt_multi_page_stats(m *base.Module) int32
//go:linkname F_bt_page_stats_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bt_page_stats_1_9
func F_bt_page_stats_1_9(m *base.Module, l0 int32) int32
//go:linkname F_bt_page_stats github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bt_page_stats
func F_bt_page_stats(m *base.Module, l0 int32) int32
//go:linkname F_bt_multi_page_stats github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bt_multi_page_stats
func F_bt_multi_page_stats(m *base.Module, l0 int32) int32
//go:linkname F_bt_page_items_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bt_page_items_1_9
func F_bt_page_items_1_9(m *base.Module, l0 int32) int32
//go:linkname F_bt_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bt_page_items
func F_bt_page_items(m *base.Module, l0 int32) int32
//go:linkname F_bt_page_items_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bt_page_items_bytea
func F_bt_page_items_bytea(m *base.Module, l0 int32) int32
//go:linkname F_bt_metap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bt_metap
func F_bt_metap(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_fsm_page_contents github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_fsm_page_contents
func F_pg_finfo_fsm_page_contents(m *base.Module) int32
//go:linkname F_fsm_page_contents github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fsm_page_contents
func F_fsm_page_contents(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gin_metapage_info github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_gin_metapage_info
func F_pg_finfo_gin_metapage_info(m *base.Module) int32
//go:linkname F_pg_finfo_gin_page_opaque_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_page_opaque_info
func F_pg_finfo_gin_page_opaque_info(m *base.Module) int32
//go:linkname F_pg_finfo_gin_leafpage_items github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_gin_leafpage_items
func F_pg_finfo_gin_leafpage_items(m *base.Module) int32
//go:linkname F_gin_metapage_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_metapage_info
func F_gin_metapage_info(m *base.Module, l0 int32) int32
//go:linkname F_gin_page_opaque_info github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gin_page_opaque_info
func F_gin_page_opaque_info(m *base.Module, l0 int32) int32
//go:linkname F_gin_leafpage_items github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gin_leafpage_items
func F_gin_leafpage_items(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_gist_page_opaque_info github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gist_page_opaque_info
func F_pg_finfo_gist_page_opaque_info(m *base.Module) int32
//go:linkname F_pg_finfo_gist_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gist_page_items
func F_pg_finfo_gist_page_items(m *base.Module) int32
//go:linkname F_pg_finfo_gist_page_items_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_gist_page_items_bytea
func F_pg_finfo_gist_page_items_bytea(m *base.Module) int32
//go:linkname F_gist_page_opaque_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gist_page_opaque_info
func F_gist_page_opaque_info(m *base.Module, l0 int32) int32
//go:linkname F_gist_page_items_bytea github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gist_page_items_bytea
func F_gist_page_items_bytea(m *base.Module, l0 int32) int32
//go:linkname F_gist_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gist_page_items
func F_gist_page_items(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hash_page_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_hash_page_type
func F_pg_finfo_hash_page_type(m *base.Module) int32
//go:linkname F_pg_finfo_hash_page_stats github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_hash_page_stats
func F_pg_finfo_hash_page_stats(m *base.Module) int32
//go:linkname F_pg_finfo_hash_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_hash_page_items
func F_pg_finfo_hash_page_items(m *base.Module) int32
//go:linkname F_pg_finfo_hash_bitmap_info github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_hash_bitmap_info
func F_pg_finfo_hash_bitmap_info(m *base.Module) int32
//go:linkname F_pg_finfo_hash_metapage_info github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hash_metapage_info
func F_pg_finfo_hash_metapage_info(m *base.Module) int32
//go:linkname F_hash_page_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hash_page_type
func F_hash_page_type(m *base.Module, l0 int32) int32
//go:linkname F_hash_page_stats github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hash_page_stats
func F_hash_page_stats(m *base.Module, l0 int32) int32
//go:linkname F_hash_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_page_items
func F_hash_page_items(m *base.Module, l0 int32) int32
//go:linkname F_hash_bitmap_info github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hash_bitmap_info
func F_hash_bitmap_info(m *base.Module, l0 int32) int32
//go:linkname F_hash_metapage_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hash_metapage_info
func F_hash_metapage_info(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_heap_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_heap_page_items
func F_pg_finfo_heap_page_items(m *base.Module) int32
//go:linkname F_heap_page_items github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_page_items
func F_heap_page_items(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_tuple_data_split github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_tuple_data_split
func F_pg_finfo_tuple_data_split(m *base.Module) int32
//go:linkname F_tuple_data_split github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuple_data_split
func F_tuple_data_split(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_heap_tuple_infomask_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_heap_tuple_infomask_flags
func F_pg_finfo_heap_tuple_infomask_flags(m *base.Module) int32
//go:linkname F_heap_tuple_infomask_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_tuple_infomask_flags
func F_heap_tuple_infomask_flags(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_pageinspect github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Pg_magic_func_pageinspect
func F_Pg_magic_func_pageinspect(m *base.Module) int32
//go:linkname F_pg_finfo_get_raw_page_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_get_raw_page_1_9
func F_pg_finfo_get_raw_page_1_9(m *base.Module) int32
//go:linkname F_get_raw_page_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_raw_page_1_9
func F_get_raw_page_1_9(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_get_raw_page github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_get_raw_page
func F_pg_finfo_get_raw_page(m *base.Module) int32
//go:linkname F_get_raw_page github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_raw_page
func F_get_raw_page(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_get_raw_page_fork_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_get_raw_page_fork_1_9
func F_pg_finfo_get_raw_page_fork_1_9(m *base.Module) int32
//go:linkname F_get_raw_page_fork_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_raw_page_fork_1_9
func F_get_raw_page_fork_1_9(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_get_raw_page_fork github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_get_raw_page_fork
func F_pg_finfo_get_raw_page_fork(m *base.Module) int32
//go:linkname F_get_raw_page_fork github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_raw_page_fork
func F_get_raw_page_fork(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_page_header github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_page_header
func F_pg_finfo_page_header(m *base.Module) int32
//go:linkname F_page_header github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_page_header
func F_page_header(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_page_checksum_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_page_checksum_1_9
func F_pg_finfo_page_checksum_1_9(m *base.Module) int32
//go:linkname F_pg_finfo_page_checksum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_page_checksum
func F_pg_finfo_page_checksum(m *base.Module) int32
//go:linkname F_page_checksum_1_9 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_page_checksum_1_9
func F_page_checksum_1_9(m *base.Module, l0 int32) int32
//go:linkname F_page_checksum github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_page_checksum
func F_page_checksum(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_pg_buffercache github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_Pg_magic_func_pg_buffercache
func F_Pg_magic_func_pg_buffercache(m *base.Module) int32
//go:linkname F_pg_finfo_pg_buffercache_pages github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_pg_buffercache_pages
func F_pg_finfo_pg_buffercache_pages(m *base.Module) int32
//go:linkname F_pg_finfo_pg_buffercache_numa_pages github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pg_buffercache_numa_pages
func F_pg_finfo_pg_buffercache_numa_pages(m *base.Module) int32
//go:linkname F_pg_finfo_pg_buffercache_summary github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pg_buffercache_summary
func F_pg_finfo_pg_buffercache_summary(m *base.Module) int32
//go:linkname F_pg_finfo_pg_buffercache_usage_counts github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_pg_buffercache_usage_counts
func F_pg_finfo_pg_buffercache_usage_counts(m *base.Module) int32
//go:linkname F_pg_finfo_pg_buffercache_evict github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_pg_buffercache_evict
func F_pg_finfo_pg_buffercache_evict(m *base.Module) int32
//go:linkname F_pg_finfo_pg_buffercache_evict_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_pg_buffercache_evict_relation
func F_pg_finfo_pg_buffercache_evict_relation(m *base.Module) int32
//go:linkname F_pg_finfo_pg_buffercache_evict_all github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_pg_buffercache_evict_all
func F_pg_finfo_pg_buffercache_evict_all(m *base.Module) int32
//go:linkname F_pg_buffercache_pages github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_buffercache_pages
func F_pg_buffercache_pages(m *base.Module, l0 int32) int32
//go:linkname F_pg_buffercache_numa_pages github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_buffercache_numa_pages
func F_pg_buffercache_numa_pages(m *base.Module, l0 int32) int32
//go:linkname F_pg_buffercache_summary github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_buffercache_summary
func F_pg_buffercache_summary(m *base.Module, l0 int32) int32
//go:linkname F_pg_buffercache_usage_counts github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_buffercache_usage_counts
func F_pg_buffercache_usage_counts(m *base.Module, l0 int32) int32
//go:linkname F_pg_buffercache_evict github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_buffercache_evict
func F_pg_buffercache_evict(m *base.Module, l0 int32) int32
//go:linkname F_pg_buffercache_evict_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_buffercache_evict_relation
func F_pg_buffercache_evict_relation(m *base.Module, l0 int32) int32
//go:linkname F_pg_buffercache_evict_all github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_buffercache_evict_all
func F_pg_buffercache_evict_all(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_pg_freespacemap github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Pg_magic_func_pg_freespacemap
func F_Pg_magic_func_pg_freespacemap(m *base.Module) int32
//go:linkname F_pg_finfo_pg_freespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_pg_freespace
func F_pg_finfo_pg_freespace(m *base.Module) int32
//go:linkname F_pg_freespace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_freespace
func F_pg_freespace(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_autoprewarm_start_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_autoprewarm_start_worker
func F_pg_finfo_autoprewarm_start_worker(m *base.Module) int32
//go:linkname F_pg_finfo_autoprewarm_dump_now github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_autoprewarm_dump_now
func F_pg_finfo_autoprewarm_dump_now(m *base.Module) int32
//go:linkname F__PG_init_pg_prewarm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__PG_init_pg_prewarm
func F__PG_init_pg_prewarm(m *base.Module)
//go:linkname F_autoprewarm_main github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_autoprewarm_main
func F_autoprewarm_main(m *base.Module, l0 int32)
//go:linkname F_autoprewarm_database_main github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_autoprewarm_database_main
func F_autoprewarm_database_main(m *base.Module, l0 int32)
//go:linkname F_autoprewarm_start_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_autoprewarm_start_worker
func F_autoprewarm_start_worker(m *base.Module, l0 int32) int32
//go:linkname F_autoprewarm_dump_now github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_autoprewarm_dump_now
func F_autoprewarm_dump_now(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_pg_prewarm github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Pg_magic_func_pg_prewarm
func F_Pg_magic_func_pg_prewarm(m *base.Module) int32
//go:linkname F_pg_finfo_pg_prewarm github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_pg_prewarm
func F_pg_finfo_pg_prewarm(m *base.Module) int32
//go:linkname F_pg_prewarm github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_prewarm
func F_pg_prewarm(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hamming_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_hamming_distance
func F_pg_finfo_hamming_distance(m *base.Module) int32
//go:linkname F_hamming_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hamming_distance
func F_hamming_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_jaccard_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_jaccard_distance
func F_pg_finfo_jaccard_distance(m *base.Module) int32
//go:linkname F_jaccard_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_jaccard_distance
func F_jaccard_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_in github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_in
func F_pg_finfo_halfvec_in(m *base.Module) int32
//go:linkname F_halfvec_in github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_halfvec_in
func F_halfvec_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_out github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_halfvec_out
func F_pg_finfo_halfvec_out(m *base.Module) int32
//go:linkname F_halfvec_out github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_halfvec_out
func F_halfvec_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_typmod_in github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_typmod_in
func F_pg_finfo_halfvec_typmod_in(m *base.Module) int32
//go:linkname F_halfvec_typmod_in github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_halfvec_typmod_in
func F_halfvec_typmod_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_halfvec_recv
func F_pg_finfo_halfvec_recv(m *base.Module) int32
//go:linkname F_halfvec_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_halfvec_recv
func F_halfvec_recv(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_send github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_halfvec_send
func F_pg_finfo_halfvec_send(m *base.Module) int32
//go:linkname F_halfvec_send github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_halfvec_send
func F_halfvec_send(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_halfvec
func F_pg_finfo_halfvec(m *base.Module) int32
//go:linkname F_halfvec github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_halfvec
func F_halfvec(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_array_to_halfvec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_array_to_halfvec
func F_pg_finfo_array_to_halfvec(m *base.Module) int32
//go:linkname F_array_to_halfvec github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_to_halfvec
func F_array_to_halfvec(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_to_float4 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_halfvec_to_float4
func F_pg_finfo_halfvec_to_float4(m *base.Module) int32
//go:linkname F_halfvec_to_float4 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_halfvec_to_float4
func F_halfvec_to_float4(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_to_halfvec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_vector_to_halfvec
func F_pg_finfo_vector_to_halfvec(m *base.Module) int32
//go:linkname F_vector_to_halfvec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_vector_to_halfvec
func F_vector_to_halfvec(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_l2_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_halfvec_l2_distance
func F_pg_finfo_halfvec_l2_distance(m *base.Module) int32
//go:linkname F_halfvec_l2_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_halfvec_l2_distance
func F_halfvec_l2_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_l2_squared_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_halfvec_l2_squared_distance
func F_pg_finfo_halfvec_l2_squared_distance(m *base.Module) int32
//go:linkname F_halfvec_l2_squared_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_halfvec_l2_squared_distance
func F_halfvec_l2_squared_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_halfvec_inner_product
func F_pg_finfo_halfvec_inner_product(m *base.Module) int32
//go:linkname F_halfvec_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_halfvec_inner_product
func F_halfvec_inner_product(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_negative_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_halfvec_negative_inner_product
func F_pg_finfo_halfvec_negative_inner_product(m *base.Module) int32
//go:linkname F_halfvec_negative_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_halfvec_negative_inner_product
func F_halfvec_negative_inner_product(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_cosine_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_cosine_distance
func F_pg_finfo_halfvec_cosine_distance(m *base.Module) int32
//go:linkname F_halfvec_cosine_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_halfvec_cosine_distance
func F_halfvec_cosine_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_spherical_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_spherical_distance
func F_pg_finfo_halfvec_spherical_distance(m *base.Module) int32
//go:linkname F_halfvec_spherical_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_halfvec_spherical_distance
func F_halfvec_spherical_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_l1_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_l1_distance
func F_pg_finfo_halfvec_l1_distance(m *base.Module) int32
//go:linkname F_halfvec_l1_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_halfvec_l1_distance
func F_halfvec_l1_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_vector_dims github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_halfvec_vector_dims
func F_pg_finfo_halfvec_vector_dims(m *base.Module) int32
//go:linkname F_halfvec_vector_dims github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_halfvec_vector_dims
func F_halfvec_vector_dims(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_l2_norm github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_l2_norm
func F_pg_finfo_halfvec_l2_norm(m *base.Module) int32
//go:linkname F_halfvec_l2_norm github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_halfvec_l2_norm
func F_halfvec_l2_norm(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_l2_normalize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_halfvec_l2_normalize
func F_pg_finfo_halfvec_l2_normalize(m *base.Module) int32
//go:linkname F_halfvec_l2_normalize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_halfvec_l2_normalize
func F_halfvec_l2_normalize(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_add github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_add
func F_pg_finfo_halfvec_add(m *base.Module) int32
//go:linkname F_halfvec_add github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_halfvec_add
func F_halfvec_add(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_sub github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_halfvec_sub
func F_pg_finfo_halfvec_sub(m *base.Module) int32
//go:linkname F_halfvec_sub github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_halfvec_sub
func F_halfvec_sub(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_mul github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_halfvec_mul
func F_pg_finfo_halfvec_mul(m *base.Module) int32
//go:linkname F_halfvec_mul github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_halfvec_mul
func F_halfvec_mul(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_halfvec_concat
func F_pg_finfo_halfvec_concat(m *base.Module) int32
//go:linkname F_halfvec_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_halfvec_concat
func F_halfvec_concat(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_binary_quantize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_binary_quantize
func F_pg_finfo_halfvec_binary_quantize(m *base.Module) int32
//go:linkname F_halfvec_binary_quantize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_halfvec_binary_quantize
func F_halfvec_binary_quantize(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_subvector github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_subvector
func F_pg_finfo_halfvec_subvector(m *base.Module) int32
//go:linkname F_halfvec_subvector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_halfvec_subvector
func F_halfvec_subvector(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_halfvec_lt
func F_pg_finfo_halfvec_lt(m *base.Module) int32
//go:linkname F_halfvec_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_halfvec_lt
func F_halfvec_lt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_le github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_halfvec_le
func F_pg_finfo_halfvec_le(m *base.Module) int32
//go:linkname F_halfvec_le github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_halfvec_le
func F_halfvec_le(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_halfvec_eq
func F_pg_finfo_halfvec_eq(m *base.Module) int32
//go:linkname F_halfvec_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_halfvec_eq
func F_halfvec_eq(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_halfvec_ne
func F_pg_finfo_halfvec_ne(m *base.Module) int32
//go:linkname F_halfvec_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_halfvec_ne
func F_halfvec_ne(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_halfvec_ge
func F_pg_finfo_halfvec_ge(m *base.Module) int32
//go:linkname F_halfvec_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_halfvec_ge
func F_halfvec_ge(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_halfvec_gt
func F_pg_finfo_halfvec_gt(m *base.Module) int32
//go:linkname F_halfvec_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_halfvec_gt
func F_halfvec_gt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_halfvec_cmp
func F_pg_finfo_halfvec_cmp(m *base.Module) int32
//go:linkname F_halfvec_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_halfvec_cmp
func F_halfvec_cmp(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_accum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_halfvec_accum
func F_pg_finfo_halfvec_accum(m *base.Module) int32
//go:linkname F_halfvec_accum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_halfvec_accum
func F_halfvec_accum(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_avg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_halfvec_avg
func F_pg_finfo_halfvec_avg(m *base.Module) int32
//go:linkname F_halfvec_avg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_halfvec_avg
func F_halfvec_avg(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_to_halfvec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_sparsevec_to_halfvec
func F_pg_finfo_sparsevec_to_halfvec(m *base.Module) int32
//go:linkname F_sparsevec_to_halfvec github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sparsevec_to_halfvec
func F_sparsevec_to_halfvec(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hnswhandler github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hnswhandler
func F_pg_finfo_hnswhandler(m *base.Module) int32
//go:linkname F_hnswhandler github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hnswhandler
func F_hnswhandler(m *base.Module, l0 int32) int32
//go:linkname F_HnswParallelBuildMain github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HnswParallelBuildMain
func F_HnswParallelBuildMain(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_finfo_hnsw_halfvec_support github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_hnsw_halfvec_support
func F_pg_finfo_hnsw_halfvec_support(m *base.Module) int32
//go:linkname F_hnsw_halfvec_support github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hnsw_halfvec_support
func F_hnsw_halfvec_support(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hnsw_bit_support github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_hnsw_bit_support
func F_pg_finfo_hnsw_bit_support(m *base.Module) int32
//go:linkname F_hnsw_bit_support github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hnsw_bit_support
func F_hnsw_bit_support(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_hnsw_sparsevec_support github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_hnsw_sparsevec_support
func F_pg_finfo_hnsw_sparsevec_support(m *base.Module) int32
//go:linkname F_hnsw_sparsevec_support github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hnsw_sparsevec_support
func F_hnsw_sparsevec_support(m *base.Module, l0 int32) int32
//go:linkname F_IvfflatParallelBuildMain github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_IvfflatParallelBuildMain
func F_IvfflatParallelBuildMain(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_finfo_ivfflathandler github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ivfflathandler
func F_pg_finfo_ivfflathandler(m *base.Module) int32
//go:linkname F_ivfflathandler github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ivfflathandler
func F_ivfflathandler(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ivfflat_halfvec_support github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_ivfflat_halfvec_support
func F_pg_finfo_ivfflat_halfvec_support(m *base.Module) int32
//go:linkname F_ivfflat_halfvec_support github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ivfflat_halfvec_support
func F_ivfflat_halfvec_support(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_ivfflat_bit_support github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_ivfflat_bit_support
func F_pg_finfo_ivfflat_bit_support(m *base.Module) int32
//go:linkname F_ivfflat_bit_support github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ivfflat_bit_support
func F_ivfflat_bit_support(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_in github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_sparsevec_in
func F_pg_finfo_sparsevec_in(m *base.Module) int32
//go:linkname F_sparsevec_in github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sparsevec_in
func F_sparsevec_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_out github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_sparsevec_out
func F_pg_finfo_sparsevec_out(m *base.Module) int32
//go:linkname F_sparsevec_out github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_sparsevec_out
func F_sparsevec_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_typmod_in github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_sparsevec_typmod_in
func F_pg_finfo_sparsevec_typmod_in(m *base.Module) int32
//go:linkname F_sparsevec_typmod_in github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sparsevec_typmod_in
func F_sparsevec_typmod_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_sparsevec_recv
func F_pg_finfo_sparsevec_recv(m *base.Module) int32
//go:linkname F_sparsevec_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_sparsevec_recv
func F_sparsevec_recv(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_send github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_sparsevec_send
func F_pg_finfo_sparsevec_send(m *base.Module) int32
//go:linkname F_sparsevec_send github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_sparsevec_send
func F_sparsevec_send(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_sparsevec
func F_pg_finfo_sparsevec(m *base.Module) int32
//go:linkname F_sparsevec github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sparsevec
func F_sparsevec(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_to_sparsevec github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_vector_to_sparsevec
func F_pg_finfo_vector_to_sparsevec(m *base.Module) int32
//go:linkname F_vector_to_sparsevec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vector_to_sparsevec
func F_vector_to_sparsevec(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_to_sparsevec github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_halfvec_to_sparsevec
func F_pg_finfo_halfvec_to_sparsevec(m *base.Module) int32
//go:linkname F_halfvec_to_sparsevec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_halfvec_to_sparsevec
func F_halfvec_to_sparsevec(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_array_to_sparsevec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_array_to_sparsevec
func F_pg_finfo_array_to_sparsevec(m *base.Module) int32
//go:linkname F_array_to_sparsevec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_array_to_sparsevec
func F_array_to_sparsevec(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_l2_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_sparsevec_l2_distance
func F_pg_finfo_sparsevec_l2_distance(m *base.Module) int32
//go:linkname F_sparsevec_l2_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_sparsevec_l2_distance
func F_sparsevec_l2_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_l2_squared_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_sparsevec_l2_squared_distance
func F_pg_finfo_sparsevec_l2_squared_distance(m *base.Module) int32
//go:linkname F_sparsevec_l2_squared_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sparsevec_l2_squared_distance
func F_sparsevec_l2_squared_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_sparsevec_inner_product
func F_pg_finfo_sparsevec_inner_product(m *base.Module) int32
//go:linkname F_sparsevec_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_sparsevec_inner_product
func F_sparsevec_inner_product(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_negative_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_sparsevec_negative_inner_product
func F_pg_finfo_sparsevec_negative_inner_product(m *base.Module) int32
//go:linkname F_sparsevec_negative_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_sparsevec_negative_inner_product
func F_sparsevec_negative_inner_product(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_cosine_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_sparsevec_cosine_distance
func F_pg_finfo_sparsevec_cosine_distance(m *base.Module) int32
//go:linkname F_sparsevec_cosine_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sparsevec_cosine_distance
func F_sparsevec_cosine_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_l1_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_sparsevec_l1_distance
func F_pg_finfo_sparsevec_l1_distance(m *base.Module) int32
//go:linkname F_sparsevec_l1_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sparsevec_l1_distance
func F_sparsevec_l1_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_l2_norm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_sparsevec_l2_norm
func F_pg_finfo_sparsevec_l2_norm(m *base.Module) int32
//go:linkname F_sparsevec_l2_norm github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sparsevec_l2_norm
func F_sparsevec_l2_norm(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_l2_normalize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_sparsevec_l2_normalize
func F_pg_finfo_sparsevec_l2_normalize(m *base.Module) int32
//go:linkname F_sparsevec_l2_normalize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_sparsevec_l2_normalize
func F_sparsevec_l2_normalize(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_sparsevec_lt
func F_pg_finfo_sparsevec_lt(m *base.Module) int32
//go:linkname F_sparsevec_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_sparsevec_lt
func F_sparsevec_lt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_le github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_sparsevec_le
func F_pg_finfo_sparsevec_le(m *base.Module) int32
//go:linkname F_sparsevec_le github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sparsevec_le
func F_sparsevec_le(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_sparsevec_eq
func F_pg_finfo_sparsevec_eq(m *base.Module) int32
//go:linkname F_sparsevec_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_sparsevec_eq
func F_sparsevec_eq(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_sparsevec_ne
func F_pg_finfo_sparsevec_ne(m *base.Module) int32
//go:linkname F_sparsevec_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sparsevec_ne
func F_sparsevec_ne(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_sparsevec_ge
func F_pg_finfo_sparsevec_ge(m *base.Module) int32
//go:linkname F_sparsevec_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_sparsevec_ge
func F_sparsevec_ge(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_sparsevec_gt
func F_pg_finfo_sparsevec_gt(m *base.Module) int32
//go:linkname F_sparsevec_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sparsevec_gt
func F_sparsevec_gt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_sparsevec_cmp
func F_pg_finfo_sparsevec_cmp(m *base.Module) int32
//go:linkname F_sparsevec_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sparsevec_cmp
func F_sparsevec_cmp(m *base.Module, l0 int32) int32
//go:linkname F_Pg_magic_func_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Pg_magic_func_vector
func F_Pg_magic_func_vector(m *base.Module) int32
//go:linkname F__PG_init_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__PG_init_vector
func F__PG_init_vector(m *base.Module)
//go:linkname F_pg_finfo_vector_in github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_vector_in
func F_pg_finfo_vector_in(m *base.Module) int32
//go:linkname F_vector_in github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vector_in
func F_vector_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_out github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_vector_out
func F_pg_finfo_vector_out(m *base.Module) int32
//go:linkname F_vector_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vector_out
func F_vector_out(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_typmod_in github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_vector_typmod_in
func F_pg_finfo_vector_typmod_in(m *base.Module) int32
//go:linkname F_vector_typmod_in github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vector_typmod_in
func F_vector_typmod_in(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_vector_recv
func F_pg_finfo_vector_recv(m *base.Module) int32
//go:linkname F_vector_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vector_recv
func F_vector_recv(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_send github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_vector_send
func F_pg_finfo_vector_send(m *base.Module) int32
//go:linkname F_vector_send github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vector_send
func F_vector_send(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_vector
func F_pg_finfo_vector(m *base.Module) int32
//go:linkname F_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_vector
func F_vector(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_array_to_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_array_to_vector
func F_pg_finfo_array_to_vector(m *base.Module) int32
//go:linkname F_array_to_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_array_to_vector
func F_array_to_vector(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_to_float4 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_vector_to_float4
func F_pg_finfo_vector_to_float4(m *base.Module) int32
//go:linkname F_vector_to_float4 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vector_to_float4
func F_vector_to_float4(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_halfvec_to_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_halfvec_to_vector
func F_pg_finfo_halfvec_to_vector(m *base.Module) int32
//go:linkname F_halfvec_to_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_halfvec_to_vector
func F_halfvec_to_vector(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_l2_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_l2_distance
func F_pg_finfo_l2_distance(m *base.Module) int32
//go:linkname F_l2_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_l2_distance
func F_l2_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_l2_squared_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_vector_l2_squared_distance
func F_pg_finfo_vector_l2_squared_distance(m *base.Module) int32
//go:linkname F_vector_l2_squared_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vector_l2_squared_distance
func F_vector_l2_squared_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_inner_product
func F_pg_finfo_inner_product(m *base.Module) int32
//go:linkname F_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_inner_product
func F_inner_product(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_negative_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_vector_negative_inner_product
func F_pg_finfo_vector_negative_inner_product(m *base.Module) int32
//go:linkname F_vector_negative_inner_product github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_vector_negative_inner_product
func F_vector_negative_inner_product(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_cosine_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_cosine_distance
func F_pg_finfo_cosine_distance(m *base.Module) int32
//go:linkname F_cosine_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cosine_distance
func F_cosine_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_spherical_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_vector_spherical_distance
func F_pg_finfo_vector_spherical_distance(m *base.Module) int32
//go:linkname F_vector_spherical_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_vector_spherical_distance
func F_vector_spherical_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_l1_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_l1_distance
func F_pg_finfo_l1_distance(m *base.Module) int32
//go:linkname F_l1_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_l1_distance
func F_l1_distance(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_dims github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_finfo_vector_dims
func F_pg_finfo_vector_dims(m *base.Module) int32
//go:linkname F_pg_finfo_vector_norm github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_vector_norm
func F_pg_finfo_vector_norm(m *base.Module) int32
//go:linkname F_vector_norm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vector_norm
func F_vector_norm(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_l2_normalize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_l2_normalize
func F_pg_finfo_l2_normalize(m *base.Module) int32
//go:linkname F_l2_normalize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_l2_normalize
func F_l2_normalize(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_add github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_vector_add
func F_pg_finfo_vector_add(m *base.Module) int32
//go:linkname F_vector_add github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_vector_add
func F_vector_add(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_sub github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_vector_sub
func F_pg_finfo_vector_sub(m *base.Module) int32
//go:linkname F_vector_sub github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vector_sub
func F_vector_sub(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_mul github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_finfo_vector_mul
func F_pg_finfo_vector_mul(m *base.Module) int32
//go:linkname F_vector_mul github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_vector_mul
func F_vector_mul(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_vector_concat
func F_pg_finfo_vector_concat(m *base.Module) int32
//go:linkname F_vector_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vector_concat
func F_vector_concat(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_binary_quantize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_binary_quantize
func F_pg_finfo_binary_quantize(m *base.Module) int32
//go:linkname F_binary_quantize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_binary_quantize
func F_binary_quantize(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_subvector github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_subvector
func F_pg_finfo_subvector(m *base.Module) int32
//go:linkname F_subvector github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_subvector
func F_subvector(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_vector_lt
func F_pg_finfo_vector_lt(m *base.Module) int32
//go:linkname F_vector_lt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_vector_lt
func F_vector_lt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_le github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_finfo_vector_le
func F_pg_finfo_vector_le(m *base.Module) int32
//go:linkname F_vector_le github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_vector_le
func F_vector_le(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_vector_eq
func F_pg_finfo_vector_eq(m *base.Module) int32
//go:linkname F_vector_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_vector_eq
func F_vector_eq(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_vector_ne
func F_pg_finfo_vector_ne(m *base.Module) int32
//go:linkname F_vector_ne github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_vector_ne
func F_vector_ne(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_finfo_vector_ge
func F_pg_finfo_vector_ge(m *base.Module) int32
//go:linkname F_vector_ge github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vector_ge
func F_vector_ge(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_vector_gt
func F_pg_finfo_vector_gt(m *base.Module) int32
//go:linkname F_vector_gt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vector_gt
func F_vector_gt(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_vector_cmp
func F_pg_finfo_vector_cmp(m *base.Module) int32
//go:linkname F_vector_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vector_cmp
func F_vector_cmp(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_accum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_vector_accum
func F_pg_finfo_vector_accum(m *base.Module) int32
//go:linkname F_vector_accum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_vector_accum
func F_vector_accum(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_combine github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_finfo_vector_combine
func F_pg_finfo_vector_combine(m *base.Module) int32
//go:linkname F_vector_combine github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_vector_combine
func F_vector_combine(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_vector_avg github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_vector_avg
func F_pg_finfo_vector_avg(m *base.Module) int32
//go:linkname F_vector_avg github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vector_avg
func F_vector_avg(m *base.Module, l0 int32) int32
//go:linkname F_pg_finfo_sparsevec_to_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_finfo_sparsevec_to_vector
func F_pg_finfo_sparsevec_to_vector(m *base.Module) int32
//go:linkname F_sparsevec_to_vector github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sparsevec_to_vector
func F_sparsevec_to_vector(m *base.Module, l0 int32) int32
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
