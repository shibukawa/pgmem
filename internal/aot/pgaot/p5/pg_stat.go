package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_stat_file(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v201 int64
	_ = v201
	var v207 int64
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v232 int32
	_ = v232
	var v235 int64
	_ = v235
	v6 = m.G0
	v8 = v6 - int32(160)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v15 == int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = base.B2i32(v18 != int64(0))
	goto L5
L4:
	;
	v21 = int32(0)
	goto L5
L5:
	;
	v22 = F_convert_and_check_filename(m, v11)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	m.G0 = v8 + int32(160)
	return v235
L7:
	;
	v28 = F___fstatat(m, int32(-100), v22, v8-int32(-64), int32(0))
	mBase = m.M
	goto L8
L8:
	;
	if v28 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v21 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v56 = F_CreateTemplateTupleDesc(m, int32(6))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L19
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_file[0]))
	if v34 != int32(44) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
	v235 = int64(0)
	goto L6
L15:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
	F_errmsg(m, int32(_a_F_pg_stat_file_0), v8)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_pg_stat_file_1), int32(437), int32(_a_F_pg_stat_file_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	F_TupleDescInitEntry(m, v56, int32(1), int32(_a_F_pg_stat_file_3), int32(20), int32(-1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_TupleDescInitEntry(m, v56, int32(2), int32(_a_F_pg_stat_file_4), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_TupleDescInitEntry(m, v56, int32(3), int32(_a_F_pg_stat_file_5), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_TupleDescInitEntry(m, v56, int32(4), int32(_a_F_pg_stat_file_6), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_TupleDescInitEntry(m, v56, int32(5), int32(_a_F_pg_stat_file_7), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_TupleDescInitEntry(m, v56, int32(6), int32(_a_F_pg_stat_file_8), int32(16), int32(-1), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v100 = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v100 < v109 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v187 = F_BlessTupleDesc(m, v56)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L45
	}
L27:
	;
	v113 = v56 + int32(28)
	v120 = v100
	v121 = v109
	v123 = v100
	goto L31
L28:
	;
	v177 = v100
	v184 = v109
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v177
	goto L26
L30:
	;
	v177 = v171
	v184 = v150
	goto L29
L31:
	;
	v129 = v113 + v109<<(uint(int32(3))%32) + v120*int32(100)
	v132 = v113 + v120<<(uint(int32(3))%32)
	if v109 != v121 {
		v150 = v121
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v171 = v109
	goto L30
L33:
	;
	v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v132)+2)))
	if v151 <= int32(0) {
		v171 = v120
		goto L30
	} else {
		goto L41
	}
L34:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+7)))
	if v134 != int32(118) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v150 = v120
	goto L33
L36:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+4)))
	if v137 != int32(1) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+6)))
	if v140&int32(6) != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(v132)+2)))
	if v143 <= int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+90)))
	if v146 != int32(118) {
		v150 = v109
		goto L33
	} else {
		goto L40
	}
L40:
	;
	goto L35
L41:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+90)))
	if v154 == int32(118) {
		v171 = v120
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+5)))
	v163 = (v123 + v157 - int32(1)) & (int32(0) - v157)
	if int32(_a_F_pg_stat_file_9) < v163 {
		v171 = v120
		goto L30
	} else {
		goto L43
	}
L43:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v132))) = uint16(v163)
	v169 = v120 + int32(1)
	if v169 != v109 {
		v120 = v169
		v121 = v150
		v123 = v163 + v151
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L32
L45:
	;
	v189 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v189)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v189
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v8)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v193
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v8)+104))
	goto L46
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v195*int64(1000000) - int64(946684800000000)
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v8)+120))
	goto L47
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v201*int64(1000000) - int64(946684800000000)
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v8)+136))
	goto L48
L48:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v212)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v207*int64(1000000) - int64(946684800000000)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = base.I64_extend_i32_u(base.B2i32(v215&int32(_a_F_pg_stat_file_10) == int32(_a_F_pg_stat_file_11)))
	v226 = F_heap_form_tuple(m, v56, v8+int32(16), v8+int32(8))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_pfree(m, v22)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
	v231 = F_HeapTupleHeaderGetDatum(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v235 = v231
	goto L6
}
func F_pg_stat_get_backend_wal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pgstat_fetch_stat_backend_by_pid(m, v9, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v36 = int64(0)
			m.G0 = v7 + int32(48)
			return v36
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v11)+2920))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v21
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v11)+2912))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v23
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v11)+2904))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v25
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v11)+2896))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v27
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v11)+2888))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v29
			v33 = F_pg_stat_wal_build_tuple(m, v7+int32(8), v20)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				v36 = v33
				m.G0 = v7 + int32(48)
				return v36
			}
		}
	}
}
func F_pg_stat_get_bgwriter_stat_reset_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_bgwriter(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+24))
		return v6
	}
}
func F_pg_stat_get_checkpointer_buffers_written(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+64))
		return v6
	}
}
func F_pg_stat_get_checkpointer_restartpoints_requested(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+32))
		return v6
	}
}
func F_pg_stat_get_checkpointer_sync_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+56))
		return base.I64_reinterpret_f64(base.F64_convert_i64_s(v6))
	}
}
func F_pg_stat_get_db_active_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+200))
			return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v11), float64(1000)))
		}
	}
}
func F_pg_stat_get_db_blocks_fetched(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+16))
			return v11
		}
	}
}
func F_pg_stat_get_db_conflict_all(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+120))
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v3)+112))
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+104))
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v3)+96))
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v3)+88))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v3)+80))
			return v11 + (v12 + (v13 + (v14 + (v15 + v16))))
		}
	}
}
func F_pg_stat_get_db_conflict_bufferpin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+112))
			return v11
		}
	}
}
func F_pg_stat_get_db_parallel_workers_launched(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+248))
			return v11
		}
	}
}
func F_pg_stat_get_function_total_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pgstat_fetch_stat_funcentry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int64(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
			return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v14), float64(1000)))
		}
	}
}
func F_pg_stat_get_last_analyze_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+152))
			if v9 != int64(0) {
				v16 = v9
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				v16 = int64(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			v16 = int64(0)
		}
		return v16
	}
}
func F_pg_stat_get_last_autovacuum_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+136))
			if v9 != int64(0) {
				v16 = v9
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				v16 = int64(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			v16 = int64(0)
		}
		return v16
	}
}
func F_pg_stat_get_subscription(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(96)
	m.G0 = v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v26 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v30 = v29
	goto L3
L2:
	;
	v30 = v2
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[0]))
	v42 = F_LWLockAcquire(m, v38+int32(_a_F_pg_stat_get_subscription_0), int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[1]))
	if v45 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[0]))
	F_LWLockRelease(m, v342+int32(_a_F_pg_stat_get_subscription_0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L4
	} else {
		goto L65
	}
L8:
	;
	v55 = v2
	goto L9
L9:
	;
	v71 = int32(0)
	base.MemoryFill(m, v24+int32(16), v71, int32(80))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)) = uint16(v71)
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = int64(0)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[2]))
	v82 = v79 + v55<<(uint(int32(7))%32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+36))
	if v83 == v71 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	v316 = v55 + int32(1)
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[1]))
	if v316 < v318 {
		v55 = v316
		goto L9
	} else {
		goto L64
	}
L12:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v82)+128))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v82)+120))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v82)+112))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v82)+104))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v82)+96))
	v91 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+80)))
	v92 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v82)+52)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+32)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	v96 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v98 == v96 {
		v203 = v96
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v204 = int32(0)
	if base.B2i32(v203 == v204)|base.B2i32(base.B2i32(v30 == v204)|base.B2i32(v30 == v93) == v204) != 0 {
		goto L11
	} else {
		goto L25
	}
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[0]))
	v106 = F_LWLockAcquire(m, v102+int32(512), int32(1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[3]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if int32(0) < v110 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[4]))
	v117 = v96
	goto L20
L17:
	;
	v155 = v96
	goto L18
L18:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[0]))
	F_LWLockRelease(m, v177+int32(512))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L24
	}
L19:
	;
	v155 = base.B2i32(v152 != int32(0))
	goto L18
L20:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(36)+v117<<(uint(int32(2))%32))))
	v144 = v116 + v141*int32(768)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	if v145 == v98 {
		v152 = v144
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v152 = int32(0)
	goto L19
L22:
	;
	v148 = v117 + int32(1)
	if v148 != v110 {
		v117 = v148
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v203 = v155
	goto L13
L25:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = base.I64_extend_i32_u(v93)
	v216 = int32(1)
	v217 = v94 & v216
	v218 = int32(0)
	if base.B2i32(v217 == v218)|base.B2i32(v95 != v216) == v218 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v90 == int64(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)) = uint8(v238)
	goto L26
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = base.I64_extend_i32_s(v213)
	goto L27
L29:
	;
	goto L30
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = base.I64_extend_i32_s(v213)
	v230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v230)
	if base.B2i32(v217 == int32(0))|base.B2i32(v95 != int32(4)) != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+40)) = v91
	goto L26
L32:
	;
	if v89 == int64(0) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v242 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)) = uint8(v242)
	goto L32
L34:
	;
	goto L35
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = v90
	goto L32
L36:
	;
	if v88 == int64(0) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v247 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)) = uint8(v247)
	goto L36
L38:
	;
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+56)) = v89
	goto L36
L40:
	;
	if v87 == int64(0) {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v252 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)) = uint8(v252)
	goto L40
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+64)) = v88
	goto L40
L44:
	;
	if v86 == int64(0) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v257 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)) = uint8(v257)
	goto L44
L46:
	;
	goto L47
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+72)) = v87
	goto L44
L48:
	;
	switch v95 {
	case 0:
		goto L55
	case 1:
		goto L56
	case 2:
		goto L57
	case 3:
		v282 = int32(_a_F_pg_stat_get_subscription_1)
		goto L53
	case 4:
		goto L54
	default:
		goto L52
	}
L49:
	;
	v262 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v262)
	goto L48
L50:
	;
	goto L51
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+80)) = v86
	goto L48
L52:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	F_tuplestore_putvalues(m, v288, v289, v24+int32(16), v24)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L62
	}
L53:
	;
	v283 = F_cstring_to_text(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L61
	}
L54:
	;
	v282 = int32(_a_F_pg_stat_get_subscription_2)
	goto L53
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	v282 = int32(_a_F_pg_stat_get_subscription_3)
	goto L53
L57:
	;
	v282 = int32(_a_F_pg_stat_get_subscription_4)
	goto L53
L58:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_get_subscription_5), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_pg_stat_get_subscription_6), int32(1771), int32(_a_F_pg_stat_get_subscription_7))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+88)) = base.I64_extend_i32_u(v283)
	goto L52
L62:
	;
	if v30 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	goto L11
L64:
	;
	goto L10
L65:
	;
	m.G0 = v24 + int32(96)
	return int64(0)
}
func F_pg_stat_get_tuples_fetched(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+24))
			return v11
		}
	}
}
func F_pg_stat_get_tuples_hot_updated(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+56))
			return v11
		}
	}
}
func F_pg_stat_get_tuples_updated(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+40))
			return v11
		}
	}
}
func F_pg_stat_get_vacuum_count(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+128))
			return v11
		}
	}
}
func F_pg_stat_get_xact_blocks_hit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+120))
			return v11
		}
	}
}
func F_pg_stat_get_xact_numscans(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+16))
			return v11
		}
	}
}
func F_pg_stat_reset_backend_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_BackendPidGetProc(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int64(0)
L2:
	;
	return int64(0)
L3:
	;
	if v4 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v10 = int32(0)
	if v3 == v10 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v48 = v4
	goto L6
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_backend_stats[0]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v54 = base.I32_div_s(v48-v51, int32(768))
	v55 = F_pgstat_get_beentry_by_proc_number(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L19
	}
L7:
	;
	if v45 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v45 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_backend_stats[1]))
	v22 = v10
	goto L13
L11:
	;
	v45 = v39
	goto L7
L12:
	;
	v39 = v29 + int32(768)
	goto L11
L13:
	;
	v25 = v22 * int32(768)
	v26 = v18 + v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v27 == v3 {
		v39 = v26
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v45 = int32(0)
	goto L7
L15:
	;
	v29 = v18 + v25
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+780))
	if v30 == v3 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v33 = v22 + int32(2)
	if v33 != int32(38) {
		v22 = v33
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v48 = v45
	goto L6
L19:
	;
	if v55 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	goto L21
L21:
	;
	if int32(base.Ui32(int32(_a_F_pg_stat_reset_backend_stats_0))>>(uint(v59)%32))&base.B2i32(base.Ui32(v59) < base.Ui32(int32(17))) == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_pgstat_reset(m, int32(6), int32(0), base.I64_extend_i32_s(v54))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L1
}
func F_pg_stat_reset_slru(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v100 int32
	_ = v100
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v5 == int32(1) {
		F_pgstat_reset_of_kind(m, int32(12))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = F_text_to_cstring(m, v16)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v23 = m.G0
				v24 = int32(16)
				v25 = v23 - v24
				m.G0 = v25
				F_gettimeofday(m, v25)
				mBase = m.M
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
				m.G0 = v25 + v24
				v39 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_0), v18)
				mBase = m.M
				if v39 == int32(0) {
					v73 = int32(0)
				} else {
					v44 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_1), v18)
					mBase = m.M
					if v44 == int32(0) {
						v73 = int32(1)
					} else {
						v49 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_2), v18)
						mBase = m.M
						if v49 == int32(0) {
							v73 = int32(2)
						} else {
							v54 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_3), v18)
							mBase = m.M
							if v54 == int32(0) {
								v73 = int32(3)
							} else {
								v59 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_4), v18)
								mBase = m.M
								if v59 == int32(0) {
									v73 = int32(4)
								} else {
									v64 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_5), v18)
									mBase = m.M
									if v64 == int32(0) {
										v73 = int32(5)
									} else {
										v71 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_6), v18)
										mBase = m.M
										if v71 != 0 {
											v72 = int32(7)
										} else {
											v72 = int32(6)
										}
										v73 = v72
									}
								}
							}
						}
					}
				}
				v75 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_slru[0]))
				v77 = v75 + int32(_a_F_pg_stat_reset_slru_7)
				v79 = F_LWLockAcquire(m, v77, int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int64(0)
				} else {
					v83 = v75 + v73<<(uint(int32(6))%32)
					v84 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[1]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[2]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[3]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[4]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[5]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[6]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[7]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[8]))) = v29 + v28*int64(1000000) - int64(946684800000000)
					F_LWLockRelease(m, v77)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int64(0)
					} else {
						return int64(0)
					}
				}
			}
		}
	}
}
func F_pg_stat_statements_1_12(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pg_stat_statements_internal(m, l0, int32(8), base.B2i32(v3 != int64(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_statements_1_9(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pg_stat_statements_internal(m, l0, int32(5), base.B2i32(v3 != int64(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_statements_reset(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v2 = int32(0)
	v6 = F_entry_reset(m, v2, v2, int64(0), v2)
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_wal_build_tuple(m *base.Module, l0 int32, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int64
	_ = v171
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v205 int32
	_ = v205
	v3 = int32(0)
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(352)
	m.G0 = v9
	*(*int64)(unsafe.Add(mBase, uint32(v9)+344)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v9)+336)) = v6
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+300)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+296)) = v3
	v20 = F_CreateTemplateTupleDesc(m, int32(6))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_TupleDescInitEntry(m, v20, int32(1), int32(_a_F_pg_stat_wal_build_tuple_0), int32(20), int32(-1), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_TupleDescInitEntry(m, v20, int32(2), int32(_a_F_pg_stat_wal_build_tuple_1), int32(20), int32(-1), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_TupleDescInitEntry(m, v20, int32(3), int32(_a_F_pg_stat_wal_build_tuple_2), int32(1700), int32(-1), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_TupleDescInitEntry(m, v20, int32(4), int32(_a_F_pg_stat_wal_build_tuple_3), int32(1700), int32(-1), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v20, int32(5), int32(_a_F_pg_stat_wal_build_tuple_4), int32(20), int32(-1), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v20, int32(6), int32(_a_F_pg_stat_wal_build_tuple_5), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v66 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v66 < v75 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v153 = F_BlessTupleDesc(m, v20)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L28
	}
L10:
	;
	v79 = v20 + int32(28)
	v86 = v66
	v87 = v75
	v89 = v66
	goto L14
L11:
	;
	v143 = v66
	v150 = v75
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v143
	goto L9
L13:
	;
	v143 = v137
	v150 = v116
	goto L12
L14:
	;
	v95 = v79 + v75<<(uint(int32(3))%32) + v86*int32(100)
	v98 = v79 + v86<<(uint(int32(3))%32)
	if v75 != v87 {
		v116 = v87
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v137 = v75
	goto L13
L16:
	;
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+2)))
	if v117 <= int32(0) {
		v137 = v86
		goto L13
	} else {
		goto L24
	}
L17:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+7)))
	if v100 != int32(118) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v116 = v86
	goto L16
L19:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+4)))
	if v103 != int32(1) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+6)))
	if v106&int32(6) != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+2)))
	if v109 <= int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+90)))
	if v112 != int32(118) {
		v116 = v75
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L18
L24:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+90)))
	if v120 == int32(118) {
		v137 = v86
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+5)))
	v129 = (v89 + v123 - int32(1)) & (int32(0) - v123)
	if int32(_a_F_pg_stat_wal_build_tuple_6) < v129 {
		v137 = v86
		goto L13
	} else {
		goto L26
	}
L26:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v98))) = uint16(v129)
	v135 = v86 + int32(1)
	if v135 != v75 {
		v86 = v135
		v87 = v116
		v89 = v129 + v117
		goto L14
	} else {
		goto L27
	}
L27:
	;
	goto L15
L28:
	;
	v155 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+304)) = v155
	v157 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+312)) = v157
	v159 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v159
	v162 = v9 + int32(32)
	v167 = F_pg_snprintf(m, v162, int32(256), int32(_a_F_pg_stat_wal_build_tuple_7), v9+int32(16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v171 = base.I64_extend_i32_u(v162)
	v174 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), v171, int64(0), int64(-1))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+320)) = v174
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v177
	v181 = F_pg_snprintf(m, v162, int32(256), int32(_a_F_pg_stat_wal_build_tuple_7), v9)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v187 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), v171, int64(0), int64(-1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+328)) = v187
	v190 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+336)) = v190
	if l1 != int64(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v201 = F_heap_form_tuple(m, v20, v9+int32(304), v9+int32(296))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L37
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+344)) = l1
	goto L33
L35:
	;
	goto L36
L36:
	;
	v195 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+301)) = uint8(v195)
	goto L33
L37:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	v204 = F_HeapTupleHeaderGetDatum(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	m.G0 = v9 + int32(352)
	return v204
}
