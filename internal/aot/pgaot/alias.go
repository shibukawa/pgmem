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
//go:linkname InitElemSeg_4_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_0
func InitElemSeg_4_0(m *base.Module)
//go:linkname InitElemSeg_4_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_1
func InitElemSeg_4_1(m *base.Module)
//go:linkname InitElemSeg_4_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_2
func InitElemSeg_4_2(m *base.Module)
//go:linkname InitElemSeg_4_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.InitElemSeg_4_3
func InitElemSeg_4_3(m *base.Module)
//go:linkname InitElemSeg_5_0 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_0
func InitElemSeg_5_0(m *base.Module)
//go:linkname InitElemSeg_5_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_1
func InitElemSeg_5_1(m *base.Module)
//go:linkname InitElemSeg_5_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_2
func InitElemSeg_5_2(m *base.Module)
//go:linkname InitElemSeg_5_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.InitElemSeg_5_3
func InitElemSeg_5_3(m *base.Module)
