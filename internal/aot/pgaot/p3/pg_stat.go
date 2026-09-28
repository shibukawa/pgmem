package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pg_stat_clear_snapshot(m *base.Module, l0 int32) int64 {
	var v5 int32
	_ = v5
	F_pgstat_clear_snapshot(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_file_1arg(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_pg_stat_file(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_pg_stat_get_autoanalyze_count(m *base.Module, l0 int32) int64 {
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
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+176))
			return v11
		}
	}
}
func F_pg_stat_get_backend_io(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v17 = F_pgstat_fetch_stat_backend_by_pid(m, v14, v6+int32(12))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if v17 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				F_pg_stat_io_build_tuples(m, v13, v17+int32(8), v21, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					m.G0 = v6 + int32(16)
					return int64(0)
				}
			} else {
				m.G0 = v6 + int32(16)
				return int64(0)
			}
		}
	}
}
func F_pg_stat_get_backend_subxact(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	v2 = int32(0)
	v4 = int64(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v4
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v2)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_TupleDescInitEntry(m, v17, int32(1), int32(_a_F_pg_stat_get_backend_subxact_0), int32(23), int32(-1), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_TupleDescInitEntry(m, v17, int32(2), int32(_a_F_pg_stat_get_backend_subxact_1), int32(16), int32(-1), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v35 < v44 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v122 = F_BlessTupleDesc(m, v17)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L24
	}
L6:
	;
	v48 = v17 + int32(28)
	v55 = v35
	v56 = v44
	v58 = v35
	goto L10
L7:
	;
	v112 = v35
	v119 = v44
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v112
	goto L5
L9:
	;
	v112 = v106
	v119 = v85
	goto L8
L10:
	;
	v64 = v48 + v44<<(uint(int32(3))%32) + v55*int32(100)
	v67 = v48 + v55<<(uint(int32(3))%32)
	if v44 != v56 {
		v85 = v56
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v106 = v44
	goto L9
L12:
	;
	v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+2)))
	if v86 <= int32(0) {
		v106 = v55
		goto L9
	} else {
		goto L20
	}
L13:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+7)))
	if v69 != int32(118) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v85 = v55
	goto L12
L15:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)))
	if v72 != int32(1) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+6)))
	if v75&int32(6) != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+2)))
	if v78 <= int32(0) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+90)))
	if v81 != int32(118) {
		v85 = v44
		goto L12
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+90)))
	if v89 == int32(118) {
		v106 = v55
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+5)))
	v98 = (v58 + v92 - int32(1)) & (int32(0) - v92)
	if int32(_a_F_pg_stat_get_backend_subxact_2) < v98 {
		v106 = v55
		goto L9
	} else {
		goto L22
	}
L22:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v67))) = uint16(v98)
	v104 = v55 + int32(1)
	if v104 != v44 {
		v55 = v104
		v56 = v85
		v58 = v98 + v86
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L11
L24:
	;
	v124 = F_pgstat_get_beentry_by_proc_number(m, v15)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v136 = F_heap_form_tuple(m, v17, v7+int32(16), v7+int32(14))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L30
	}
L26:
	;
	if v124 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v126 = int64(*(*int32)(unsafe.Add(mBase, uint32(v124)+420)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v126
	v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v124)+424)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v128
	goto L25
L28:
	;
	goto L29
L29:
	;
	v130 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v130)
	goto L25
L30:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v139 = F_HeapTupleHeaderGetDatum(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	m.G0 = v7 + int32(32)
	return v139
}
func F_pg_stat_get_backend_userid(m *base.Module, l0 int32) int64 {
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
	v4 = F_pgstat_get_beentry_by_proc_number(m, v3)
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
			v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4)+52)))
			return v14
		}
	}
}
func F_pg_stat_get_backend_wait_event(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v79 = F_cstring_to_text(m, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L31
	}
L2:
	;
	return int64(0)
L3:
	;
	if v5 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v77 = int32(_a_F_pg_stat_get_backend_wait_event_0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event[0]))
	v15 = F_has_privs_of_role(m, v13, int32(3375))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v24 = F_BackendPidGetProc(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L13
	}
L8:
	;
	if v15 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
	v20 = F_has_privs_of_role(m, v18, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v20 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v77 = int32(_a_F_pg_stat_get_backend_wait_event_1)
	goto L1
L12:
	;
	v73 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v73)
	return int64(0)
L13:
	;
	if v24 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v29 = int32(0)
	if v28 == v29 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v67 = v24
	goto L16
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+648))
	v69 = F_pgstat_get_wait_event(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L29
	}
L17:
	;
	if v64 == int32(0) {
		goto L12
	} else {
		goto L28
	}
L18:
	;
	v64 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event[1]))
	v41 = v29
	goto L23
L21:
	;
	v64 = v58
	goto L17
L22:
	;
	v58 = v48 + int32(768)
	goto L21
L23:
	;
	v44 = v41 * int32(768)
	v45 = v37 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if v46 == v28 {
		v58 = v45
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v64 = int32(0)
	goto L17
L25:
	;
	v48 = v37 + v44
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+780))
	if v49 == v28 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v52 = v41 + int32(2)
	if v52 != int32(38) {
		v41 = v52
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v67 = v64
	goto L16
L29:
	;
	if v69 != 0 {
		v77 = v69
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L12
L31:
	;
	return base.I64_extend_i32_u(v79)
}
func F_pg_stat_get_buf_alloc(m *base.Module, l0 int32) int64 {
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
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+16))
		return v6
	}
}
func F_pg_stat_get_checkpointer_write_time(m *base.Module, l0 int32) int64 {
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
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+48))
		return base.I64_reinterpret_f64(base.F64_convert_i64_s(v6))
	}
}
func F_pg_stat_get_db_blk_read_time(m *base.Module, l0 int32) int64 {
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
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+168))
			return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v11), float64(1000)))
		}
	}
}
func F_pg_stat_get_db_conflict_lock(m *base.Module, l0 int32) int64 {
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
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+88))
			return v11
		}
	}
}
func F_pg_stat_get_db_conflict_logicalslot(m *base.Module, l0 int32) int64 {
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
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+104))
			return v11
		}
	}
}
func F_pg_stat_get_db_session_time(m *base.Module, l0 int32) int64 {
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
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+192))
			return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v11), float64(1000)))
		}
	}
}
func F_pg_stat_get_db_stat_reset_time(m *base.Module, l0 int32) int64 {
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
	v5 = F_pgstat_fetch_stat_dbentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+256))
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
func F_pg_stat_get_db_temp_bytes(m *base.Module, l0 int32) int64 {
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
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+136))
			return v11
		}
	}
}
func F_pg_stat_get_db_tuples_inserted(m *base.Module, l0 int32) int64 {
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
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+48))
			return v11
		}
	}
}
func F_pg_stat_get_lock(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_pgstat_snapshot_fixed(m, int32(11))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(0)
	goto L4
L4:
	;
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v23
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v33
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v19<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_lock[0])))
	v40 = F_cstring_to_text(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v7 - int32(-64)
	return int64(0)
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = base.I64_extend_i32_u(v40)
	v45 = v19 * int32(24)
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_pg_stat_get_lock[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_pg_stat_get_lock[2])))
	*(*float64)(unsafe.Add(mBase, uint32(v7)+32)) = base.F64_mul(base.F64_convert_i64_s(v50), float64(0.001))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_pg_stat_get_lock[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v55
	v58 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_lock[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	F_tuplestore_putvalues(m, v60, v61, v5+int32(-48), v5+int32(-56))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v69 = v19 + int32(1)
	if v69 != int32(12) {
		v19 = v69
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_pg_stat_get_recovery(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v144 int64
	_ = v144
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery[0])))
	if v19 == int32(1) {
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery[1]))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+308))
		v27 = base.B2i32(v25 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery[0])) = uint8(v27)
		v29 = v27
	} else {
		v29 = int32(0)
	}
	if v29 == int32(0) {
		v32 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
		v144 = int64(0)
		m.G0 = v15 + int32(16)
		return v144
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery[2]))
		v38 = F_has_privs_of_role(m, v36, int32(3375))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int64(0)
		} else {
			if v38 == int32(0) {
				v44 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
				v144 = int64(0)
				m.G0 = v15 + int32(16)
				return v144
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery[3]))
				v51 = base.AtomicRmwXchg32(m, v48, int32(96), int32(1))
				if v51 != 0 {
					F_s_lock(m, v48+int32(96), int32(_a_F_pg_stat_get_recovery_0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery[3]))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+80))
						v60 = *(*int64)(unsafe.Add(mBase, uint32(v58)+72))
						v61 = *(*int64)(unsafe.Add(mBase, uint32(v58)+64))
						v62 = int64(*(*int32)(unsafe.Add(mBase, uint32(v58)+56)))
						v63 = *(*int64)(unsafe.Add(mBase, uint32(v58)+48))
						v64 = int64(*(*int32)(unsafe.Add(mBase, uint32(v58)+40)))
						v65 = *(*int64)(unsafe.Add(mBase, uint32(v58)+32))
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v58)+24))
						v67 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
						v68 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v58)+96)), uint32(v68))
						v74 = F_get_call_result_type(m, l0, v68, v15+int32(12))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int64(0)
						} else {
							if v74 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return int64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_pg_stat_get_recovery_1), int32(0))
									mBase = m.M
									v156 = m.ExcPending
									if v156 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_stat_get_recovery_2), int32(815), int32(_a_F_pg_stat_get_recovery_3))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								v81 = F_palloc0_mul(m, int32(8), v80)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int64(0)
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
									v86 = F_palloc0_mul(m, int32(1), v85)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v81))) = v67
										if v66 != int64(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v66
										} else {
											v92 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)) = uint8(v92)
										}
										if v65 != int64(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = v64
											*(*int64)(unsafe.Add(mBase, uint32(v81)+16)) = v65
										} else {
											v98 = int32(257)
											*(*uint16)(unsafe.Add(mBase, uint32(v86)+2)) = uint16(v98)
										}
										if v63 != int64(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v81)+40)) = v62
											*(*int64)(unsafe.Add(mBase, uint32(v81)+32)) = v63
										} else {
											v104 = int32(257)
											*(*uint16)(unsafe.Add(mBase, uint32(v86)+4)) = uint16(v104)
										}
										if v61 != int64(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v81)+48)) = v61
										} else {
											v109 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v86)+6)) = uint8(v109)
										}
										if v60 != int64(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v81)+56)) = v60
										} else {
											v114 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v86)+7)) = uint8(v114)
										}
										if base.Ui32(v59) <= base.Ui32(int32(2)) {
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v59<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_recovery[4])))
											v122 = v120
										} else {
											v122 = int32(0)
										}
										v123 = F_cstring_to_text(m, v122)
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v81)+64)) = base.I64_extend_i32_u(v123)
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
											v128 = F_heap_form_tuple(m, v127, v81, v86)
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int64(0)
											} else {
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
												v131 = F_HeapTupleHeaderGetDatum(m, v130)
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int64(0)
												} else {
													v144 = v131
													m.G0 = v15 + int32(16)
													return v144
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery[3]))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+80))
					v60 = *(*int64)(unsafe.Add(mBase, uint32(v58)+72))
					v61 = *(*int64)(unsafe.Add(mBase, uint32(v58)+64))
					v62 = int64(*(*int32)(unsafe.Add(mBase, uint32(v58)+56)))
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v58)+48))
					v64 = int64(*(*int32)(unsafe.Add(mBase, uint32(v58)+40)))
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v58)+32))
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v58)+24))
					v67 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
					v68 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v58)+96)), uint32(v68))
					v74 = F_get_call_result_type(m, l0, v68, v15+int32(12))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						if v74 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_pg_stat_get_recovery_1), int32(0))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_stat_get_recovery_2), int32(815), int32(_a_F_pg_stat_get_recovery_3))
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							v81 = F_palloc0_mul(m, int32(8), v80)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int64(0)
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
								v86 = F_palloc0_mul(m, int32(1), v85)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v81))) = v67
									if v66 != int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v66
									} else {
										v92 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)) = uint8(v92)
									}
									if v65 != int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = v64
										*(*int64)(unsafe.Add(mBase, uint32(v81)+16)) = v65
									} else {
										v98 = int32(257)
										*(*uint16)(unsafe.Add(mBase, uint32(v86)+2)) = uint16(v98)
									}
									if v63 != int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v81)+40)) = v62
										*(*int64)(unsafe.Add(mBase, uint32(v81)+32)) = v63
									} else {
										v104 = int32(257)
										*(*uint16)(unsafe.Add(mBase, uint32(v86)+4)) = uint16(v104)
									}
									if v61 != int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v81)+48)) = v61
									} else {
										v109 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+6)) = uint8(v109)
									}
									if v60 != int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v81)+56)) = v60
									} else {
										v114 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+7)) = uint8(v114)
									}
									if base.Ui32(v59) <= base.Ui32(int32(2)) {
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v59<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_recovery[4])))
										v122 = v120
									} else {
										v122 = int32(0)
									}
									v123 = F_cstring_to_text(m, v122)
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v81)+64)) = base.I64_extend_i32_u(v123)
										v127 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
										v128 = F_heap_form_tuple(m, v127, v81, v86)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int64(0)
										} else {
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
											v131 = F_HeapTupleHeaderGetDatum(m, v130)
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int64(0)
											} else {
												v144 = v131
												m.G0 = v15 + int32(16)
												return v144
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_stat_get_recovery_prefetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	v4 = m.G0
	v6 = v4 - int32(96)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+8)) = uint16(v14)
		v16 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v16
		v18 = int32(_a_F_pg_stat_get_recovery_prefetch_0)
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
		v23 = base.AtomicRmwCmpxchg64(m, v19, v14, v16, v16)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v23
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
		v30 = base.AtomicRmwCmpxchg64(m, v26, int32(8), v16, v16)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v30
		v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
		v36 = int32(16)
		v37 = base.AtomicRmwCmpxchg64(m, v33, v36, v16, v16)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v37
		v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
		v44 = base.AtomicRmwCmpxchg64(m, v40, int32(24), v16, v16)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v44
		v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
		v51 = base.AtomicRmwCmpxchg64(m, v47, int32(32), v16, v16)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+48)) = v51
		v54 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
		v58 = base.AtomicRmwCmpxchg64(m, v54, int32(40), v16, v16)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = v58
		v61 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
		v65 = base.AtomicRmwCmpxchg64(m, v61, int32(48), v16, v16)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+64)) = v65
		v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
		v69 = int64(*(*int32)(unsafe.Add(mBase, uint32(v68)+56)))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+72)) = v69
		v71 = int64(*(*int32)(unsafe.Add(mBase, uint32(v68)+60)))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+80)) = v71
		v73 = int64(*(*int32)(unsafe.Add(mBase, uint32(v68)+64)))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+88)) = v73
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		F_tuplestore_putvalues(m, v75, v76, v6+v36, v6)
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return int64(0)
		} else {
			m.G0 = v6 + int32(96)
			return int64(0)
		}
	}
}
func F_pg_stat_get_replication_slot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v229 int64
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int64
	_ = v279
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v302 int32
	_ = v302
	var v305 int64
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(352)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
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
	base.MemoryFill(m, v8+int32(176), int32(0), int32(104))
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+165)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = v20
	v25 = F_CreateTemplateTupleDesc(m, int32(13))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_TupleDescInitEntry(m, v25, int32(1), int32(_a_F_pg_stat_get_replication_slot_0), int32(25), int32(-1), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_TupleDescInitEntry(m, v25, int32(2), int32(_a_F_pg_stat_get_replication_slot_1), int32(20), int32(-1), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_TupleDescInitEntry(m, v25, int32(3), int32(_a_F_pg_stat_get_replication_slot_2), int32(20), int32(-1), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v25, int32(4), int32(_a_F_pg_stat_get_replication_slot_3), int32(20), int32(-1), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v25, int32(5), int32(_a_F_pg_stat_get_replication_slot_4), int32(20), int32(-1), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v25, int32(6), int32(_a_F_pg_stat_get_replication_slot_5), int32(20), int32(-1), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v25, int32(7), int32(_a_F_pg_stat_get_replication_slot_6), int32(20), int32(-1), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v25, int32(8), int32(_a_F_pg_stat_get_replication_slot_7), int32(20), int32(-1), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v25, int32(9), int32(_a_F_pg_stat_get_replication_slot_8), int32(20), int32(-1), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v25, int32(10), int32(_a_F_pg_stat_get_replication_slot_9), int32(20), int32(-1), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v25, int32(11), int32(_a_F_pg_stat_get_replication_slot_10), int32(20), int32(-1), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v25, int32(12), int32(_a_F_pg_stat_get_replication_slot_11), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v25, int32(13), int32(_a_F_pg_stat_get_replication_slot_12), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v118 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v118 < v127 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v205 = F_BlessTupleDesc(m, v25)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L36
	}
L18:
	;
	v131 = v25 + int32(28)
	v138 = v118
	v139 = v127
	v141 = v118
	goto L22
L19:
	;
	v195 = v118
	v202 = v127
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v195
	goto L17
L21:
	;
	v195 = v189
	v202 = v168
	goto L20
L22:
	;
	v147 = v131 + v127<<(uint(int32(3))%32) + v138*int32(100)
	v150 = v131 + v138<<(uint(int32(3))%32)
	if v127 != v139 {
		v168 = v139
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v189 = v127
	goto L21
L24:
	;
	v169 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+2)))
	if v169 <= int32(0) {
		v189 = v138
		goto L21
	} else {
		goto L32
	}
L25:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+7)))
	if v152 != int32(118) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v168 = v138
	goto L24
L27:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+4)))
	if v155 != int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+6)))
	if v158&int32(6) != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v161 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+2)))
	if v161 <= int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+90)))
	if v164 != int32(118) {
		v168 = v127
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+90)))
	if v172 == int32(118) {
		v189 = v138
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+5)))
	v181 = (v141 + v175 - int32(1)) & (int32(0) - v175)
	if int32(_a_F_pg_stat_get_replication_slot_13) < v181 {
		v189 = v138
		goto L21
	} else {
		goto L34
	}
L34:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v150))) = uint16(v181)
	v187 = v138 + int32(1)
	if v187 != v127 {
		v138 = v187
		v139 = v168
		v141 = v181 + v169
		goto L22
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	v209 = F_text_to_cstring(m, v11)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v212 = F_strncpy(m, v8+int32(288), v209, int32(64))
	mBase = m.M
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+63)) = uint8(v213)
	goto L38
L38:
	;
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v8)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v215
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v8)+336))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v217
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v8)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v219
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v8)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v221
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v8)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v223
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v8)+304))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v225
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v8)+296))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v227
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v8)+288))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v229
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_replication_slot[0]))
	v236 = F_LWLockAcquire(m, v232+int32(_a_F_pg_stat_get_replication_slot_14), int32(1))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v239 = F_SearchNamedReplicationSlot(m, v8, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_replication_slot[0]))
	F_LWLockRelease(m, v259+int32(_a_F_pg_stat_get_replication_slot_14))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L46
	}
L41:
	;
	if v239 == int32(0) {
		v257 = v2
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_replication_slot[1]))
	v247 = base.I32_div_s(v239-v244, int32(296))
	goto L43
L43:
	;
	if v247 == int32(-1) {
		v257 = v2
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v251 = int32(0)
	v254 = F_pgstat_fetch_entry(m, int32(4), v251, base.I64_extend_i32_s(v247), v251)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v257 = v254
	goto L40
L46:
	;
	if v257 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v267 = v8 - int32(-64)
	base.MemoryFill(m, v267, int32(0), int32(96))
	v272 = v267
	goto L49
L48:
	;
	v272 = v257
	goto L49
L49:
	;
	v275 = F_cstring_to_text(m, v8+int32(288))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+176)) = base.I64_extend_i32_u(v275)
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v272)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+184)) = v279
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v272)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+192)) = v281
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v272)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+200)) = v283
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v272)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+208)) = v285
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v272)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+216)) = v287
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v272)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+224)) = v289
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v272)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+232)) = v291
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v272)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+240)) = v293
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v272)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+248)) = v295
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v272)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+256)) = v297
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v272)+80))
	if v299 == int64(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v272)+88))
	if v305 == int64(0) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v302 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+171)) = uint8(v302)
	goto L51
L53:
	;
	goto L54
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+264)) = v299
	goto L51
L55:
	;
	v315 = F_heap_form_tuple(m, v25, v8+int32(176), v8+int32(160))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L59
	}
L56:
	;
	v308 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+172)) = uint8(v308)
	goto L55
L57:
	;
	goto L58
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+272)) = v305
	goto L55
L59:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v315)+16))
	v318 = F_HeapTupleHeaderGetDatum(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	m.G0 = v8 + int32(352)
	return v318
}
func F_pg_stat_get_snapshot_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[0])))
	if v10 != 0 {
		v12 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[1])) = v12
		*(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[2])) = v12
		*(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[3])) = v12
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[4])) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[5])) = v21
		*(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[6])) = v21
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[7]))
		if v30 != 0 {
			F_MemoryContextDelete(m, v30)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[7])) = int32(0)
				F_pgstat_clear_backend_activity_snapshot(m)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					v41 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[0])) = uint8(v41)
					v45 = v7 + int32(15)
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[6]))
					if v47 == int32(2) {
						v50 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v50)
						v53 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[8]))
						v57 = v53
					} else {
						v54 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v54)
						v57 = int64(0)
					}
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
					if v58 == int32(0) {
						v61 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v61)
						v64 = int64(0)
					} else {
						v64 = v57
					}
					m.G0 = v7 + int32(16)
					return v64
				}
			}
		} else {
			F_pgstat_clear_backend_activity_snapshot(m)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int64(0)
			} else {
				v41 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[0])) = uint8(v41)
				v45 = v7 + int32(15)
				v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[6]))
				if v47 == int32(2) {
					v50 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v50)
					v53 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[8]))
					v57 = v53
				} else {
					v54 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v54)
					v57 = int64(0)
				}
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				if v58 == int32(0) {
					v61 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v61)
					v64 = int64(0)
				} else {
					v64 = v57
				}
				m.G0 = v7 + int32(16)
				return v64
			}
		}
	} else {
		v45 = v7 + int32(15)
		v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[6]))
		if v47 == int32(2) {
			v50 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v50)
			v53 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[8]))
			v57 = v53
		} else {
			v54 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v54)
			v57 = int64(0)
		}
		v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v58 == int32(0) {
			v61 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v61)
			v64 = int64(0)
		} else {
			v64 = v57
		}
		m.G0 = v7 + int32(16)
		return v64
	}
}
func F_pg_stat_get_subscription_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int64
	_ = v214
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	v2 = int32(0)
	v4 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(224)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	base.MemoryFill(m, v8+int32(112), v2, int32(104))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+101)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v8)+96)) = v4
	v25 = F_pgstat_fetch_entry(m, int32(5), v2, base.I64_extend_i32_u(base.I32_wrap_i64(v10)), v2)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v30 = F_CreateTemplateTupleDesc(m, int32(13))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_TupleDescInitEntry(m, v30, int32(1), int32(_a_F_pg_stat_get_subscription_stats_0), int32(26), int32(-1), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_TupleDescInitEntry(m, v30, int32(2), int32(_a_F_pg_stat_get_subscription_stats_1), int32(20), int32(-1), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_TupleDescInitEntry(m, v30, int32(3), int32(_a_F_pg_stat_get_subscription_stats_2), int32(20), int32(-1), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v30, int32(4), int32(_a_F_pg_stat_get_subscription_stats_3), int32(20), int32(-1), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v30, int32(5), int32(_a_F_pg_stat_get_subscription_stats_4), int32(20), int32(-1), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v30, int32(6), int32(_a_F_pg_stat_get_subscription_stats_5), int32(20), int32(-1), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v30, int32(7), int32(_a_F_pg_stat_get_subscription_stats_6), int32(20), int32(-1), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v30, int32(8), int32(_a_F_pg_stat_get_subscription_stats_7), int32(20), int32(-1), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v30, int32(9), int32(_a_F_pg_stat_get_subscription_stats_8), int32(20), int32(-1), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v30, int32(10), int32(_a_F_pg_stat_get_subscription_stats_9), int32(20), int32(-1), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v30, int32(11), int32(_a_F_pg_stat_get_subscription_stats_10), int32(20), int32(-1), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v30, int32(12), int32(_a_F_pg_stat_get_subscription_stats_11), int32(20), int32(-1), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v30, int32(13), int32(_a_F_pg_stat_get_subscription_stats_12), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v123 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v123 < v132 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v210 = F_BlessTupleDesc(m, v30)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L36
	}
L18:
	;
	v136 = v30 + int32(28)
	v143 = v123
	v144 = v132
	v146 = v123
	goto L22
L19:
	;
	v200 = v123
	v207 = v132
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v200
	goto L17
L21:
	;
	v200 = v194
	v207 = v173
	goto L20
L22:
	;
	v152 = v136 + v132<<(uint(int32(3))%32) + v143*int32(100)
	v155 = v136 + v143<<(uint(int32(3))%32)
	if v132 != v144 {
		v173 = v144
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v194 = v132
	goto L21
L24:
	;
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+2)))
	if v174 <= int32(0) {
		v194 = v143
		goto L21
	} else {
		goto L32
	}
L25:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+7)))
	if v157 != int32(118) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v173 = v143
	goto L24
L27:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
	if v160 != int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+6)))
	if v163&int32(6) != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+2)))
	if v166 <= int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+90)))
	if v169 != int32(118) {
		v173 = v132
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+90)))
	if v177 == int32(118) {
		v194 = v143
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+5)))
	v186 = (v146 + v180 - int32(1)) & (int32(0) - v180)
	if int32(_a_F_pg_stat_get_subscription_stats_13) < v186 {
		v194 = v143
		goto L21
	} else {
		goto L34
	}
L34:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v155))) = uint16(v186)
	v192 = v143 + int32(1)
	if v192 != v132 {
		v143 = v192
		v144 = v173
		v146 = v186 + v174
		goto L22
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	if v25 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+120)) = v219
	*(*int64)(unsafe.Add(mBase, uint32(v8)+112)) = v10 & int64(4294967295)
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+128)) = v222
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v218)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+136)) = v224
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v218)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+144)) = v226
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v218)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+152)) = v228
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v218)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = v230
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v218)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+168)) = v232
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v218)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+176)) = v234
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v218)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+184)) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v218)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+192)) = v238
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v218)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+200)) = v240
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v218)+88))
	if v242 == int64(0) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v218 = v25
	v219 = v214
	goto L37
L39:
	;
	goto L40
L40:
	;
	base.MemoryFill(m, v8, int32(0), int32(96))
	v218 = v8
	v219 = v4
	goto L37
L41:
	;
	v252 = F_heap_form_tuple(m, v30, v8+int32(112), v8+int32(96))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L45
	}
L42:
	;
	v245 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+108)) = uint8(v245)
	goto L41
L43:
	;
	goto L44
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+208)) = v242
	goto L41
L45:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
	v255 = F_HeapTupleHeaderGetDatum(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	m.G0 = v8 + int32(224)
	return v255
}
func F_pg_stat_get_tuples_deleted(m *base.Module, l0 int32) int64 {
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
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+48))
			return v11
		}
	}
}
func F_pg_stat_get_wal_senders(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
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
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(128)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v32 = F_SyncRepGetCandidateStandbys(m, v22+int32(124))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[0]))
	if int32(0) < v35 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v39 = v22 + int32(1)
	v44 = v2
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v22 + int32(128)
	return int64(0)
L7:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = int64(0)
	v67 = v60 + v44*int32(96)
	v68 = int32(164)
	v69 = v67 + v68
	v72 = base.AtomicRmwXchg32(m, v67, v68, int32(1))
	if v72 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_s_lock(m, v69, int32(_a_F_pg_stat_get_wal_senders_0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v77 = v67 + int32(88)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v78 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v291 = v44 + int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[0]))
	if v291 < v293 {
		v44 = v291
		goto L7
	} else {
		goto L78
	}
L14:
	;
	v81 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v69))), uint32(v81))
	goto L13
L15:
	;
	goto L16
L16:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v77)+80))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v77)+64))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v77)+56))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v77)+48))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v77)+32))
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v77)+24))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
	v94 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v77)+76)), uint32(v94))
	if v32 <= v94 {
		v135 = int32(1)
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = base.I64_extend_i32_s(v78)
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[2]))
	v157 = F_has_privs_of_role(m, v155, int32(3375))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L26
	}
L18:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v102 = int32(0)
	goto L19
L19:
	;
	v123 = v101 + v102*int32(48)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	if v124 != v44 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v135 = v129
	goto L17
L21:
	;
	v129 = int32(1)
	v131 = v102 + v129
	if v131 != v32 {
		v102 = v131
		goto L19
	} else {
		goto L24
	}
L22:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v126 != v78 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v135 = int32(0)
	goto L17
L24:
	;
	goto L20
L25:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_tuplestore_putvalues(m, v265, v266, v22+int32(16), v22)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L77
	}
L26:
	;
	if v157 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+7)) = int32(16843009)
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(72340172838076673)
	goto L25
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(v92) <= base.Ui32(int32(4)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v92<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_wal_senders[3])))
	v171 = v169
	goto L32
L31:
	;
	v171 = int32(_a_F_pg_stat_get_wal_senders_1)
	goto L32
L32:
	;
	v172 = F_cstring_to_text(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = base.I64_extend_i32_u(v172)
	if v93 == int64(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)) = uint8(v178)
	goto L36
L35:
	;
	goto L36
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v93
	if v91 == int64(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v183 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)) = uint8(v183)
	goto L39
L38:
	;
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v91
	if v90 == int64(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)) = uint8(v188)
	goto L42
L41:
	;
	goto L42
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v90
	if v89 == int64(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)) = uint8(v193)
	goto L45
L44:
	;
	goto L45
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v89
	if v88 < int64(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v209 = int64(0)
	if v87 < v209 {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	v198 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)) = uint8(v198)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v201 = F_palloc(m, int32(16))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v201)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v201))) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = base.I64_extend_i32_u(v201)
	goto L46
L51:
	;
	if v90 == v209 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)) = uint8(v213)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v216 = F_palloc(m, int32(16))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v216)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v216))) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = base.I64_extend_i32_u(v216)
	goto L51
L56:
	;
	v225 = int32(0)
	goto L58
L57:
	;
	v225 = v85
	goto L58
L58:
	;
	if v86 < int64(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = base.I64_extend_i32_s(v225)
	if base.B2i32(v225 == int32(0))|v135 != 0 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v228)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v231 = F_palloc(m, int32(16))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v231)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v231))) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = base.I64_extend_i32_u(v231)
	goto L59
L64:
	;
	if v225 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[4]))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+8)))
	if v251 != 0 {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	v254 = F_cstring_to_text(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L73
	}
L67:
	;
	v246 = int32(_a_F_pg_stat_get_wal_senders_2)
	goto L69
L68:
	;
	v246 = int32(_a_F_pg_stat_get_wal_senders_3)
	goto L69
L69:
	;
	v253 = v246
	goto L66
L70:
	;
	v252 = int32(_a_F_pg_stat_get_wal_senders_4)
	goto L72
L71:
	;
	v252 = int32(_a_F_pg_stat_get_wal_senders_5)
	goto L72
L72:
	;
	v253 = v252
	goto L66
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = base.I64_extend_i32_u(v254)
	if v84 == int64(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v260 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+11)) = uint8(v260)
	goto L25
L75:
	;
	goto L76
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v84
	goto L25
L77:
	;
	goto L13
L78:
	;
	goto L8
}
func F_pg_stat_get_xact_blocks_fetched(m *base.Module, l0 int32) int64 {
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
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+112))
			return v11
		}
	}
}
func F_pg_stat_have_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v488 int32
	_ = v488
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v533 int32
	_ = v533
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v677 int32
	_ = v677
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v726 int32
	_ = v726
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v775 int32
	_ = v775
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v824 int32
	_ = v824
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v873 int32
	_ = v873
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v922 int32
	_ = v922
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v971 int32
	_ = v971
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1023 int32
	_ = v1023
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1046 int32
	_ = v1046
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1081 int32
	_ = v1081
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v12 = F_text_to_cstring(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v23 = v12
	v24 = int32(_a_F_pg_stat_have_stats_0)
	goto L6
L4:
	;
	m.G0 = v18 + int32(16)
	if base.Ui32(v1067-int32(1)) <= base.Ui32(int32(12)) {
		goto L386
	} else {
		goto L387
	}
L5:
	;
	if v61 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == v28 {
		v50 = v27
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v61 = int32(0)
	goto L5
L8:
	;
	v52 = int32(1)
	if v50 != 0 {
		v23 = v23 + v52
		v24 = v24 + v52
		goto L6
	} else {
		goto L17
	}
L9:
	;
	if base.Ui32((v27-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = v27 | int32(32)
	goto L12
L11:
	;
	v38 = v27
	goto L12
L12:
	;
	if base.Ui32((v28-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v28 | int32(32)
	goto L15
L14:
	;
	v47 = v28
	goto L15
L15:
	;
	if v38 == v47 {
		v50 = v38
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v61 = v38 - v47
	goto L5
L17:
	;
	goto L7
L18:
	;
	v1067 = int32(1)
	goto L4
L19:
	;
	goto L20
L20:
	;
	v68 = v12
	v69 = int32(_a_F_pg_stat_have_stats_1)
	goto L22
L21:
	;
	if v106 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v72 == v73 {
		v95 = v72
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v106 = int32(0)
	goto L21
L24:
	;
	v97 = int32(1)
	if v95 != 0 {
		v68 = v68 + v97
		v69 = v69 + v97
		goto L22
	} else {
		goto L33
	}
L25:
	;
	if base.Ui32((v72-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v83 = v72 | int32(32)
	goto L28
L27:
	;
	v83 = v72
	goto L28
L28:
	;
	if base.Ui32((v73-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v92 = v73 | int32(32)
	goto L31
L30:
	;
	v92 = v73
	goto L31
L31:
	;
	if v83 == v92 {
		v95 = v83
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v106 = v83 - v92
	goto L21
L33:
	;
	goto L23
L34:
	;
	v1067 = int32(2)
	goto L4
L35:
	;
	goto L36
L36:
	;
	v113 = v12
	v114 = int32(_a_F_pg_stat_have_stats_2)
	goto L38
L37:
	;
	if v151 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v117 == v118 {
		v140 = v117
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v151 = int32(0)
	goto L37
L40:
	;
	v142 = int32(1)
	if v140 != 0 {
		v113 = v113 + v142
		v114 = v114 + v142
		goto L38
	} else {
		goto L49
	}
L41:
	;
	if base.Ui32((v117-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v128 = v117 | int32(32)
	goto L44
L43:
	;
	v128 = v117
	goto L44
L44:
	;
	if base.Ui32((v118-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v137 = v118 | int32(32)
	goto L47
L46:
	;
	v137 = v118
	goto L47
L47:
	;
	if v128 == v137 {
		v140 = v128
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v151 = v128 - v137
	goto L37
L49:
	;
	goto L39
L50:
	;
	v1067 = int32(3)
	goto L4
L51:
	;
	goto L52
L52:
	;
	v158 = v12
	v159 = int32(_a_F_pg_stat_have_stats_3)
	goto L54
L53:
	;
	if v196 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v162 == v163 {
		v185 = v162
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v196 = int32(0)
	goto L53
L56:
	;
	v187 = int32(1)
	if v185 != 0 {
		v158 = v158 + v187
		v159 = v159 + v187
		goto L54
	} else {
		goto L65
	}
L57:
	;
	if base.Ui32((v162-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v173 = v162 | int32(32)
	goto L60
L59:
	;
	v173 = v162
	goto L60
L60:
	;
	if base.Ui32((v163-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v182 = v163 | int32(32)
	goto L63
L62:
	;
	v182 = v163
	goto L63
L63:
	;
	if v173 == v182 {
		v185 = v173
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v196 = v173 - v182
	goto L53
L65:
	;
	goto L55
L66:
	;
	v1067 = int32(4)
	goto L4
L67:
	;
	goto L68
L68:
	;
	v203 = v12
	v204 = int32(_a_F_pg_stat_have_stats_4)
	goto L70
L69:
	;
	if v241 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L70:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v207 == v208 {
		v230 = v207
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v241 = int32(0)
	goto L69
L72:
	;
	v232 = int32(1)
	if v230 != 0 {
		v203 = v203 + v232
		v204 = v204 + v232
		goto L70
	} else {
		goto L81
	}
L73:
	;
	if base.Ui32((v207-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v218 = v207 | int32(32)
	goto L76
L75:
	;
	v218 = v207
	goto L76
L76:
	;
	if base.Ui32((v208-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v227 = v208 | int32(32)
	goto L79
L78:
	;
	v227 = v208
	goto L79
L79:
	;
	if v218 == v227 {
		v230 = v218
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v241 = v218 - v227
	goto L69
L81:
	;
	goto L71
L82:
	;
	v1067 = int32(5)
	goto L4
L83:
	;
	goto L84
L84:
	;
	v248 = v12
	v249 = int32(_a_F_pg_stat_have_stats_5)
	goto L86
L85:
	;
	if v286 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L86:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v252 == v253 {
		v275 = v252
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v286 = int32(0)
	goto L85
L88:
	;
	v277 = int32(1)
	if v275 != 0 {
		v248 = v248 + v277
		v249 = v249 + v277
		goto L86
	} else {
		goto L97
	}
L89:
	;
	if base.Ui32((v252-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v263 = v252 | int32(32)
	goto L92
L91:
	;
	v263 = v252
	goto L92
L92:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v272 = v253 | int32(32)
	goto L95
L94:
	;
	v272 = v253
	goto L95
L95:
	;
	if v263 == v272 {
		v275 = v263
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v286 = v263 - v272
	goto L85
L97:
	;
	goto L87
L98:
	;
	v1067 = int32(6)
	goto L4
L99:
	;
	goto L100
L100:
	;
	v293 = v12
	v294 = int32(_a_F_pg_stat_have_stats_6)
	goto L102
L101:
	;
	if v331 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L102:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v297 == v298 {
		v320 = v297
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v331 = int32(0)
	goto L101
L104:
	;
	v322 = int32(1)
	if v320 != 0 {
		v293 = v293 + v322
		v294 = v294 + v322
		goto L102
	} else {
		goto L113
	}
L105:
	;
	if base.Ui32((v297-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v308 = v297 | int32(32)
	goto L108
L107:
	;
	v308 = v297
	goto L108
L108:
	;
	if base.Ui32((v298-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v317 = v298 | int32(32)
	goto L111
L110:
	;
	v317 = v298
	goto L111
L111:
	;
	if v308 == v317 {
		v320 = v308
		goto L104
	} else {
		goto L112
	}
L112:
	;
	v331 = v308 - v317
	goto L101
L113:
	;
	goto L103
L114:
	;
	v1067 = int32(7)
	goto L4
L115:
	;
	goto L116
L116:
	;
	v338 = v12
	v339 = int32(_a_F_pg_stat_have_stats_7)
	goto L118
L117:
	;
	if v376 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L118:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if v342 == v343 {
		v365 = v342
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v376 = int32(0)
	goto L117
L120:
	;
	v367 = int32(1)
	if v365 != 0 {
		v338 = v338 + v367
		v339 = v339 + v367
		goto L118
	} else {
		goto L129
	}
L121:
	;
	if base.Ui32((v342-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v353 = v342 | int32(32)
	goto L124
L123:
	;
	v353 = v342
	goto L124
L124:
	;
	if base.Ui32((v343-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v362 = v343 | int32(32)
	goto L127
L126:
	;
	v362 = v343
	goto L127
L127:
	;
	if v353 == v362 {
		v365 = v353
		goto L120
	} else {
		goto L128
	}
L128:
	;
	v376 = v353 - v362
	goto L117
L129:
	;
	goto L119
L130:
	;
	v1067 = int32(8)
	goto L4
L131:
	;
	goto L132
L132:
	;
	v383 = v12
	v384 = int32(_a_F_pg_stat_have_stats_8)
	goto L134
L133:
	;
	if v421 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L134:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v387 == v388 {
		v410 = v387
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v421 = int32(0)
	goto L133
L136:
	;
	v412 = int32(1)
	if v410 != 0 {
		v383 = v383 + v412
		v384 = v384 + v412
		goto L134
	} else {
		goto L145
	}
L137:
	;
	if base.Ui32((v387-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v398 = v387 | int32(32)
	goto L140
L139:
	;
	v398 = v387
	goto L140
L140:
	;
	if base.Ui32((v388-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v407 = v388 | int32(32)
	goto L143
L142:
	;
	v407 = v388
	goto L143
L143:
	;
	if v398 == v407 {
		v410 = v398
		goto L136
	} else {
		goto L144
	}
L144:
	;
	v421 = v398 - v407
	goto L133
L145:
	;
	goto L135
L146:
	;
	v1067 = int32(9)
	goto L4
L147:
	;
	goto L148
L148:
	;
	v428 = v12
	v429 = int32(_a_F_pg_stat_have_stats_9)
	goto L150
L149:
	;
	if v466 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L150:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	if v432 == v433 {
		v455 = v432
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v466 = int32(0)
	goto L149
L152:
	;
	v457 = int32(1)
	if v455 != 0 {
		v428 = v428 + v457
		v429 = v429 + v457
		goto L150
	} else {
		goto L161
	}
L153:
	;
	if base.Ui32((v432-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v443 = v432 | int32(32)
	goto L156
L155:
	;
	v443 = v432
	goto L156
L156:
	;
	if base.Ui32((v433-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v452 = v433 | int32(32)
	goto L159
L158:
	;
	v452 = v433
	goto L159
L159:
	;
	if v443 == v452 {
		v455 = v443
		goto L152
	} else {
		goto L160
	}
L160:
	;
	v466 = v443 - v452
	goto L149
L161:
	;
	goto L151
L162:
	;
	v1067 = int32(10)
	goto L4
L163:
	;
	goto L164
L164:
	;
	v473 = v12
	v474 = int32(_a_F_pg_stat_have_stats_10)
	goto L166
L165:
	;
	if v511 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L166:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473))))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	if v477 == v478 {
		v500 = v477
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v511 = int32(0)
	goto L165
L168:
	;
	v502 = int32(1)
	if v500 != 0 {
		v473 = v473 + v502
		v474 = v474 + v502
		goto L166
	} else {
		goto L177
	}
L169:
	;
	if base.Ui32((v477-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v488 = v477 | int32(32)
	goto L172
L171:
	;
	v488 = v477
	goto L172
L172:
	;
	if base.Ui32((v478-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v497 = v478 | int32(32)
	goto L175
L174:
	;
	v497 = v478
	goto L175
L175:
	;
	if v488 == v497 {
		v500 = v488
		goto L168
	} else {
		goto L176
	}
L176:
	;
	v511 = v488 - v497
	goto L165
L177:
	;
	goto L167
L178:
	;
	v1067 = int32(11)
	goto L4
L179:
	;
	goto L180
L180:
	;
	v518 = v12
	v519 = int32(_a_F_pg_stat_have_stats_11)
	goto L182
L181:
	;
	if v556 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L182:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	if v522 == v523 {
		v545 = v522
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v556 = int32(0)
	goto L181
L184:
	;
	v547 = int32(1)
	if v545 != 0 {
		v518 = v518 + v547
		v519 = v519 + v547
		goto L182
	} else {
		goto L193
	}
L185:
	;
	if base.Ui32((v522-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v533 = v522 | int32(32)
	goto L188
L187:
	;
	v533 = v522
	goto L188
L188:
	;
	if base.Ui32((v523-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v542 = v523 | int32(32)
	goto L191
L190:
	;
	v542 = v523
	goto L191
L191:
	;
	if v533 == v542 {
		v545 = v533
		goto L184
	} else {
		goto L192
	}
L192:
	;
	v556 = v533 - v542
	goto L181
L193:
	;
	goto L183
L194:
	;
	v1067 = int32(12)
	goto L4
L195:
	;
	goto L196
L196:
	;
	v563 = v12
	v564 = int32(_a_F_pg_stat_have_stats_12)
	goto L198
L197:
	;
	if v601 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L198:
	;
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	if v567 == v568 {
		v590 = v567
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v601 = int32(0)
	goto L197
L200:
	;
	v592 = int32(1)
	if v590 != 0 {
		v563 = v563 + v592
		v564 = v564 + v592
		goto L198
	} else {
		goto L209
	}
L201:
	;
	if base.Ui32((v567-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v578 = v567 | int32(32)
	goto L204
L203:
	;
	v578 = v567
	goto L204
L204:
	;
	if base.Ui32((v568-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v587 = v568 | int32(32)
	goto L207
L206:
	;
	v587 = v568
	goto L207
L207:
	;
	if v578 == v587 {
		v590 = v578
		goto L200
	} else {
		goto L208
	}
L208:
	;
	v601 = v578 - v587
	goto L197
L209:
	;
	goto L199
L210:
	;
	v1067 = int32(13)
	goto L4
L211:
	;
	goto L212
L212:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	if v606 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L382
	}
L214:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v606)))
	if v609 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)+80))
	v613 = v12
	v614 = v610
	goto L219
L216:
	;
	v657 = v606
	goto L217
L217:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)+4))
	if v658 != 0 {
		goto L234
	} else {
		goto L235
	}
L218:
	;
	if v651 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L219:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v617 == v618 {
		v640 = v617
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v651 = int32(0)
	goto L218
L221:
	;
	v642 = int32(1)
	if v640 != 0 {
		v613 = v613 + v642
		v614 = v614 + v642
		goto L219
	} else {
		goto L230
	}
L222:
	;
	if base.Ui32((v617-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v628 = v617 | int32(32)
	goto L225
L224:
	;
	v628 = v617
	goto L225
L225:
	;
	if base.Ui32((v618-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v637 = v618 | int32(32)
	goto L228
L227:
	;
	v637 = v618
	goto L228
L228:
	;
	if v628 == v637 {
		v640 = v628
		goto L221
	} else {
		goto L229
	}
L229:
	;
	v651 = v628 - v637
	goto L218
L230:
	;
	goto L220
L231:
	;
	v1067 = int32(24)
	goto L4
L232:
	;
	goto L233
L233:
	;
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v657 = v656
	goto L217
L234:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)+80))
	v662 = v12
	v663 = v659
	goto L238
L235:
	;
	v706 = v657
	goto L236
L236:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+8))
	if v707 != 0 {
		goto L253
	} else {
		goto L254
	}
L237:
	;
	if v700 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L238:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663))))
	if v666 == v667 {
		v689 = v666
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v700 = int32(0)
	goto L237
L240:
	;
	v691 = int32(1)
	if v689 != 0 {
		v662 = v662 + v691
		v663 = v663 + v691
		goto L238
	} else {
		goto L249
	}
L241:
	;
	if base.Ui32((v666-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v677 = v666 | int32(32)
	goto L244
L243:
	;
	v677 = v666
	goto L244
L244:
	;
	if base.Ui32((v667-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v686 = v667 | int32(32)
	goto L247
L246:
	;
	v686 = v667
	goto L247
L247:
	;
	if v677 == v686 {
		v689 = v677
		goto L240
	} else {
		goto L248
	}
L248:
	;
	v700 = v677 - v686
	goto L237
L249:
	;
	goto L239
L250:
	;
	v1067 = int32(25)
	goto L4
L251:
	;
	goto L252
L252:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v706 = v705
	goto L236
L253:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+80))
	v711 = v12
	v712 = v708
	goto L257
L254:
	;
	v755 = v706
	goto L255
L255:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)+12))
	if v756 != 0 {
		goto L272
	} else {
		goto L273
	}
L256:
	;
	if v749 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L257:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711))))
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712))))
	if v715 == v716 {
		v738 = v715
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v749 = int32(0)
	goto L256
L259:
	;
	v740 = int32(1)
	if v738 != 0 {
		v711 = v711 + v740
		v712 = v712 + v740
		goto L257
	} else {
		goto L268
	}
L260:
	;
	if base.Ui32((v715-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v726 = v715 | int32(32)
	goto L263
L262:
	;
	v726 = v715
	goto L263
L263:
	;
	if base.Ui32((v716-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v735 = v716 | int32(32)
	goto L266
L265:
	;
	v735 = v716
	goto L266
L266:
	;
	if v726 == v735 {
		v738 = v726
		goto L259
	} else {
		goto L267
	}
L267:
	;
	v749 = v726 - v735
	goto L256
L268:
	;
	goto L258
L269:
	;
	v1067 = int32(26)
	goto L4
L270:
	;
	goto L271
L271:
	;
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v755 = v754
	goto L255
L272:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)+80))
	v760 = v12
	v761 = v757
	goto L276
L273:
	;
	v804 = v755
	goto L274
L274:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v804)+16))
	if v805 != 0 {
		goto L291
	} else {
		goto L292
	}
L275:
	;
	if v798 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L276:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760))))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761))))
	if v764 == v765 {
		v787 = v764
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v798 = int32(0)
	goto L275
L278:
	;
	v789 = int32(1)
	if v787 != 0 {
		v760 = v760 + v789
		v761 = v761 + v789
		goto L276
	} else {
		goto L287
	}
L279:
	;
	if base.Ui32((v764-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v775 = v764 | int32(32)
	goto L282
L281:
	;
	v775 = v764
	goto L282
L282:
	;
	if base.Ui32((v765-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v784 = v765 | int32(32)
	goto L285
L284:
	;
	v784 = v765
	goto L285
L285:
	;
	if v775 == v784 {
		v787 = v775
		goto L278
	} else {
		goto L286
	}
L286:
	;
	v798 = v775 - v784
	goto L275
L287:
	;
	goto L277
L288:
	;
	v1067 = int32(27)
	goto L4
L289:
	;
	goto L290
L290:
	;
	v803 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v804 = v803
	goto L274
L291:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v805)+80))
	v809 = v12
	v810 = v806
	goto L295
L292:
	;
	v853 = v804
	goto L293
L293:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v853)+20))
	if v854 != 0 {
		goto L310
	} else {
		goto L311
	}
L294:
	;
	if v847 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L295:
	;
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809))))
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810))))
	if v813 == v814 {
		v836 = v813
		goto L297
	} else {
		goto L298
	}
L296:
	;
	v847 = int32(0)
	goto L294
L297:
	;
	v838 = int32(1)
	if v836 != 0 {
		v809 = v809 + v838
		v810 = v810 + v838
		goto L295
	} else {
		goto L306
	}
L298:
	;
	if base.Ui32((v813-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v824 = v813 | int32(32)
	goto L301
L300:
	;
	v824 = v813
	goto L301
L301:
	;
	if base.Ui32((v814-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v833 = v814 | int32(32)
	goto L304
L303:
	;
	v833 = v814
	goto L304
L304:
	;
	if v824 == v833 {
		v836 = v824
		goto L297
	} else {
		goto L305
	}
L305:
	;
	v847 = v824 - v833
	goto L294
L306:
	;
	goto L296
L307:
	;
	v1067 = int32(28)
	goto L4
L308:
	;
	goto L309
L309:
	;
	v852 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v853 = v852
	goto L293
L310:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+80))
	v858 = v12
	v859 = v855
	goto L314
L311:
	;
	v902 = v853
	goto L312
L312:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v902)+24))
	if v903 != 0 {
		goto L329
	} else {
		goto L330
	}
L313:
	;
	if v896 == int32(0) {
		goto L326
	} else {
		goto L327
	}
L314:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v858))))
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859))))
	if v862 == v863 {
		v885 = v862
		goto L316
	} else {
		goto L317
	}
L315:
	;
	v896 = int32(0)
	goto L313
L316:
	;
	v887 = int32(1)
	if v885 != 0 {
		v858 = v858 + v887
		v859 = v859 + v887
		goto L314
	} else {
		goto L325
	}
L317:
	;
	if base.Ui32((v862-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v873 = v862 | int32(32)
	goto L320
L319:
	;
	v873 = v862
	goto L320
L320:
	;
	if base.Ui32((v863-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v882 = v863 | int32(32)
	goto L323
L322:
	;
	v882 = v863
	goto L323
L323:
	;
	if v873 == v882 {
		v885 = v873
		goto L316
	} else {
		goto L324
	}
L324:
	;
	v896 = v873 - v882
	goto L313
L325:
	;
	goto L315
L326:
	;
	v1067 = int32(29)
	goto L4
L327:
	;
	goto L328
L328:
	;
	v901 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v902 = v901
	goto L312
L329:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v903)+80))
	v907 = v12
	v908 = v904
	goto L333
L330:
	;
	v951 = v902
	goto L331
L331:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)+28))
	if v952 != 0 {
		goto L348
	} else {
		goto L349
	}
L332:
	;
	if v945 == int32(0) {
		goto L345
	} else {
		goto L346
	}
L333:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907))))
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	if v911 == v912 {
		v934 = v911
		goto L335
	} else {
		goto L336
	}
L334:
	;
	v945 = int32(0)
	goto L332
L335:
	;
	v936 = int32(1)
	if v934 != 0 {
		v907 = v907 + v936
		v908 = v908 + v936
		goto L333
	} else {
		goto L344
	}
L336:
	;
	if base.Ui32((v911-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v922 = v911 | int32(32)
	goto L339
L338:
	;
	v922 = v911
	goto L339
L339:
	;
	if base.Ui32((v912-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v931 = v912 | int32(32)
	goto L342
L341:
	;
	v931 = v912
	goto L342
L342:
	;
	if v922 == v931 {
		v934 = v922
		goto L335
	} else {
		goto L343
	}
L343:
	;
	v945 = v922 - v931
	goto L332
L344:
	;
	goto L334
L345:
	;
	v1067 = int32(30)
	goto L4
L346:
	;
	goto L347
L347:
	;
	v950 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v951 = v950
	goto L331
L348:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v952)+80))
	v956 = v12
	v957 = v953
	goto L352
L349:
	;
	v1000 = v951
	goto L350
L350:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+32))
	if v1002 == int32(0) {
		goto L213
	} else {
		goto L367
	}
L351:
	;
	if v994 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L352:
	;
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956))))
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v957))))
	if v960 == v961 {
		v983 = v960
		goto L354
	} else {
		goto L355
	}
L353:
	;
	v994 = int32(0)
	goto L351
L354:
	;
	v985 = int32(1)
	if v983 != 0 {
		v956 = v956 + v985
		v957 = v957 + v985
		goto L352
	} else {
		goto L363
	}
L355:
	;
	if base.Ui32((v960-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v971 = v960 | int32(32)
	goto L358
L357:
	;
	v971 = v960
	goto L358
L358:
	;
	if base.Ui32((v961-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v980 = v961 | int32(32)
	goto L361
L360:
	;
	v980 = v961
	goto L361
L361:
	;
	if v971 == v980 {
		v983 = v971
		goto L354
	} else {
		goto L362
	}
L362:
	;
	v994 = v971 - v980
	goto L351
L363:
	;
	goto L353
L364:
	;
	v1067 = int32(31)
	goto L4
L365:
	;
	goto L366
L366:
	;
	v999 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v1000 = v999
	goto L350
L367:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+80))
	v1008 = v12
	v1009 = v1005
	goto L369
L368:
	;
	if v1046 == int32(0) {
		v1067 = int32(32)
		goto L4
	} else {
		goto L381
	}
L369:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008))))
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009))))
	if v1012 == v1013 {
		v1035 = v1012
		goto L371
	} else {
		goto L372
	}
L370:
	;
	v1046 = int32(0)
	goto L368
L371:
	;
	v1037 = int32(1)
	if v1035 != 0 {
		v1008 = v1008 + v1037
		v1009 = v1009 + v1037
		goto L369
	} else {
		goto L380
	}
L372:
	;
	if base.Ui32((v1012-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1023 = v1012 | int32(32)
	goto L375
L374:
	;
	v1023 = v1012
	goto L375
L375:
	;
	if base.Ui32((v1013-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1032 = v1013 | int32(32)
	goto L378
L377:
	;
	v1032 = v1013
	goto L378
L378:
	;
	if v1023 == v1032 {
		v1035 = v1023
		goto L371
	} else {
		goto L379
	}
L379:
	;
	v1046 = v1023 - v1032
	goto L368
L380:
	;
	goto L370
L381:
	;
	goto L213
L382:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v12
	F_errmsg(m, int32(_a_F_pg_stat_have_stats_13), v18)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	F_errfinish(m, int32(_a_F_pg_stat_have_stats_14), int32(1469), int32(_a_F_pg_stat_have_stats_15))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L386:
	;
	v1088 = v1067*int32(84) + int32(_a_F_pg_stat_have_stats_16)
	goto L388
L387:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1081+v1067<<(uint(int32(2))%32)-int32(96))))
	v1088 = v1087
	goto L388
L388:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1088))))
	if v1089&int32(1) != 0 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1099 = int32(1)
	goto L391
L390:
	;
	v1093 = int32(0)
	v1095 = F_pgstat_get_entry_ref(m, v1067, v15, v14, v1093, v1093)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L392
	}
L391:
	;
	return base.I64_extend_i32_u(v1099)
L392:
	;
	v1099 = base.B2i32(v1095 != int32(0))
	goto L391
}
func F_pg_stat_statements(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_pg_stat_statements_internal(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_statements_1_3(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pg_stat_statements_internal(m, l0, int32(3), base.B2i32(v3 != int64(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_statements_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v491 int64
	_ = v491
	var v497 int64
	_ = v497
	var v511 int32
	_ = v511
	var v514 int64
	_ = v514
	var v517 int64
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v605 int64
	_ = v605
	var v606 int64
	_ = v606
	var v607 int64
	_ = v607
	var v608 int64
	_ = v608
	var v609 int64
	_ = v609
	var v610 int64
	_ = v610
	var v611 int64
	_ = v611
	var v612 int64
	_ = v612
	var v613 int64
	_ = v613
	var v614 int64
	_ = v614
	var v615 int64
	_ = v615
	var v616 int64
	_ = v616
	var v617 int64
	_ = v617
	var v618 int64
	_ = v618
	var v619 int64
	_ = v619
	var v620 int64
	_ = v620
	var v621 int64
	_ = v621
	var v622 int64
	_ = v622
	var v623 int64
	_ = v623
	var v624 int64
	_ = v624
	var v625 int64
	_ = v625
	var v626 int64
	_ = v626
	var v627 int64
	_ = v627
	var v628 int64
	_ = v628
	var v629 int64
	_ = v629
	var v630 int64
	_ = v630
	var v631 int64
	_ = v631
	var v632 int64
	_ = v632
	var v633 int64
	_ = v633
	var v634 int64
	_ = v634
	var v635 int64
	_ = v635
	var v636 int64
	_ = v636
	var v637 int64
	_ = v637
	var v638 int64
	_ = v638
	var v639 int64
	_ = v639
	var v640 float64
	_ = v640
	var v641 float64
	_ = v641
	var v642 int64
	_ = v642
	var v643 int64
	_ = v643
	var v644 int64
	_ = v644
	var v645 int64
	_ = v645
	var v646 int64
	_ = v646
	var v647 int64
	_ = v647
	var v648 int64
	_ = v648
	var v649 int64
	_ = v649
	var v650 int64
	_ = v650
	var v651 int64
	_ = v651
	var v652 int32
	_ = v652
	var v658 int64
	_ = v658
	var v659 int64
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v681 float64
	_ = v681
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v715 float64
	_ = v715
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v796 int64
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	v4 = int32(0)
	v73 = m.G0
	v75 = v73 - int32(800)
	m.G0 = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+796)) = v4
	v83 = F_has_privs_of_role(m, v79, int32(3375))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	if v86 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	if l2 != 0 {
		goto L92
	} else {
		goto L93
	}
L4:
	;
	v339 = v336
	v340 = v337
	v341 = v4
	v342 = v4
	v343 = v4
	v344 = v4
	v345 = v4
	v346 = v4
	v347 = int32(0)
	goto L3
L5:
	;
	v339 = v327
	v340 = v328
	v341 = v329
	v342 = v330
	v343 = v331
	v344 = v332
	v345 = v334
	v346 = v333
	v347 = int32(1)
	goto L3
L6:
	;
	v327 = v321
	v328 = v322
	v329 = v325
	v330 = v323
	v331 = v4
	v332 = v324
	v333 = v4
	v334 = int32(0)
	goto L5
L7:
	;
	v321 = v316
	v322 = v317
	v323 = v318
	v324 = v319
	v325 = int32(1)
	goto L6
L8:
	;
	v313 = int32(1)
	v316 = v313
	v317 = v313
	v318 = v312
	v319 = int32(0)
	goto L7
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L85
	}
L10:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[2]))
	if v90 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	switch v97 - int32(14) {
	case 0:
		goto L23
	default:
		goto L13
	case 4:
		goto L22
	case 5:
		goto L21
	case 9:
		goto L20
	case 18:
		goto L19
	case 19:
		goto L18
	case 29:
		goto L17
	case 35:
		goto L16
	case 38:
		goto L15
	case 40:
		goto L14
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L82
	}
L14:
	;
	if l1 == int32(9) {
		goto L76
	} else {
		goto L77
	}
L15:
	;
	if l1 == int32(8) {
		goto L70
	} else {
		goto L71
	}
L16:
	;
	if l1 == int32(7) {
		goto L64
	} else {
		goto L65
	}
L17:
	;
	if l1 == int32(6) {
		goto L58
	} else {
		goto L59
	}
L18:
	;
	if l1 == int32(5) {
		goto L52
	} else {
		goto L53
	}
L19:
	;
	if l1 == int32(4) {
		goto L46
	} else {
		goto L47
	}
L20:
	;
	if l1 == int32(3) {
		goto L40
	} else {
		goto L41
	}
L21:
	;
	if l1 == int32(2) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	if l1 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	if l1 == int32(0) {
		v336 = v4
		v337 = v4
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1704), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v336 = int32(1)
	v337 = v4
	goto L4
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1709), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v133 = int32(1)
	v336 = v133
	v337 = v133
	goto L4
L35:
	;
	goto L36
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1714), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v150 = int32(1)
	v321 = v150
	v322 = v150
	v323 = v4
	v324 = v4
	v325 = int32(0)
	goto L6
L41:
	;
	goto L42
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1718), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v312 = v4
	goto L8
L47:
	;
	goto L48
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1722), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v312 = int32(1)
	goto L8
L53:
	;
	goto L54
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1726), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v199 = int32(1)
	v316 = v199
	v317 = v199
	v318 = v199
	v319 = v199
	goto L7
L59:
	;
	goto L60
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1730), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v218 = int32(1)
	v327 = v218
	v328 = v218
	v329 = v218
	v330 = v218
	v331 = v218
	v332 = v218
	v333 = v4
	v334 = int32(0)
	goto L5
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1734), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v240 = int32(1)
	v327 = v240
	v328 = v240
	v329 = v240
	v330 = v240
	v331 = v240
	v332 = v240
	v333 = v4
	v334 = v240
	goto L5
L71:
	;
	goto L72
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1738), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v262 = int32(1)
	v327 = v262
	v328 = v262
	v329 = v262
	v330 = v262
	v331 = v262
	v332 = v262
	v333 = v262
	v334 = v262
	goto L5
L77:
	;
	goto L78
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1742), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1745), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_pg_stat_statements_internal_3), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1691), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v404 = v75 + int32(776)
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[2]))
	F_hash_seq_init(m, v404, v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L111
	}
L90:
	;
	v397 = F_qtext_load_file(m, v75+int32(796))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L110
	}
L91:
	;
	v376 = F_qtext_load_file(m, v75+int32(796))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L102
	}
L92:
	;
	v352 = base.AtomicRmwXchg32(m, v349, int32(140), int32(1))
	if v352 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v372 = F_LWLockAcquire(m, v349, int32(1))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L101
	}
L95:
	;
	F_s_lock(m, v349+int32(140), int32(_a_F_pg_stat_statements_internal_4))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+152))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v359)+148))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v359)+144))
	v363 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v359)+140)), uint32(v363))
	if v361 == v363 {
		goto L91
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v369 = F_LWLockAcquire(m, v359, int32(1))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	goto L90
L101:
	;
	v401 = v4
	goto L89
L102:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	v381 = F_LWLockAcquire(m, v379, int32(1))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v376 == int32(0) {
		goto L90
	} else {
		goto L104
	}
L104:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+144))
	if v362 == v387 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v386)+152))
	if v389 == v360 {
		v401 = v376
		goto L89
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	F_pfree(m, v376)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	goto L90
L110:
	;
	v401 = v397
	goto L89
L111:
	;
	v409 = F_hash_seq_search(m, v404)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if v409 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v419 = v409
	goto L116
L114:
	;
	goto L115
L115:
	;
	v961 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	F_LWLockRelease(m, v961)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L202
	}
L116:
	;
	v491 = *(*int64)(unsafe.Add(mBase, uint32(v419)+8))
	base.MemoryFill(m, v75+int32(336), int32(0), int32(432))
	v497 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v75)+318)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v75)+312)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v75)+304)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v75)+296)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v75)+288)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v75)+280)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v75)+272)) = v497
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+336)) = base.I64_extend_i32_u(v511)
	v514 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v419)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+344)) = v514
	if v342 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L115
L118:
	;
	v517 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v419)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+352)) = v517
	v520 = v75 + int32(360)
	v521 = int32(3)
	goto L120
L119:
	;
	v520 = v75 + int32(352)
	v521 = int32(2)
	goto L120
L120:
	;
	if v83|base.B2i32(v511 == v79) == int32(1) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v599 = base.AtomicRmwXchg32(m, v419, int32(440), int32(1))
	if v599 != 0 {
		goto L147
	} else {
		goto L148
	}
L122:
	;
	if v340 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	if v340 != 0 {
		goto L140
	} else {
		goto L141
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v520))) = v491
	v529 = v521 + int32(1)
	goto L127
L126:
	;
	v529 = v521
	goto L127
L127:
	;
	if l2 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if v401 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	v568 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(272)|v529))) = uint8(v568)
	v594 = v529
	goto L121
L131:
	;
	v563 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(272)|v529))) = uint8(v563)
	v594 = v529
	goto L121
L132:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v419)+412))
	if v532 < int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v419)+408))
	v536 = v532 + v535
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v75)+796))
	if base.Ui32(v537) <= base.Ui32(v536) {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401+v536))))
	if v540 != 0 {
		goto L131
	} else {
		goto L135
	}
L135:
	;
	v546 = v535 + v401
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v419)+416))
	v548 = F_pg_any_to_server(m, v546, v532, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v550 = F_cstring_to_text(m, v548)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v75+int32(336)+v529<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v550)
	if v548 == v546 {
		v594 = v529
		goto L121
	} else {
		goto L138
	}
L138:
	;
	F_pfree(m, v548)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v594 = v529
	goto L121
L140:
	;
	v573 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(272)|v521))) = uint8(v573)
	v577 = v521 + v573
	goto L142
L141:
	;
	v577 = v521
	goto L142
L142:
	;
	if l2 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v584 = F_cstring_to_text(m, int32(_a_F_pg_stat_statements_internal_5))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v591 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(272)|v577))) = uint8(v591)
	v594 = v577
	goto L121
L146:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v75+int32(336)+v577<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v584)
	v594 = v577
	goto L121
L147:
	;
	F_s_lock(m, v419+int32(440), int32(_a_F_pg_stat_statements_internal_4))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v605 = *(*int64)(unsafe.Add(mBase, uint32(v419)+400))
	v606 = *(*int64)(unsafe.Add(mBase, uint32(v419)+392))
	v607 = *(*int64)(unsafe.Add(mBase, uint32(v419)+384))
	v608 = *(*int64)(unsafe.Add(mBase, uint32(v419)+376))
	v609 = *(*int64)(unsafe.Add(mBase, uint32(v419)+368))
	v610 = *(*int64)(unsafe.Add(mBase, uint32(v419)+360))
	v611 = *(*int64)(unsafe.Add(mBase, uint32(v419)+352))
	v612 = *(*int64)(unsafe.Add(mBase, uint32(v419)+344))
	v613 = *(*int64)(unsafe.Add(mBase, uint32(v419)+336))
	v614 = *(*int64)(unsafe.Add(mBase, uint32(v419)+328))
	v615 = *(*int64)(unsafe.Add(mBase, uint32(v419)+320))
	v616 = *(*int64)(unsafe.Add(mBase, uint32(v419)+312))
	v617 = *(*int64)(unsafe.Add(mBase, uint32(v419)+304))
	v618 = *(*int64)(unsafe.Add(mBase, uint32(v419)+296))
	v619 = *(*int64)(unsafe.Add(mBase, uint32(v419)+288))
	v620 = *(*int64)(unsafe.Add(mBase, uint32(v419)+280))
	v621 = *(*int64)(unsafe.Add(mBase, uint32(v419)+272))
	v622 = *(*int64)(unsafe.Add(mBase, uint32(v419)+264))
	v623 = *(*int64)(unsafe.Add(mBase, uint32(v419)+248))
	v624 = *(*int64)(unsafe.Add(mBase, uint32(v419)+240))
	v625 = *(*int64)(unsafe.Add(mBase, uint32(v419)+232))
	v626 = *(*int64)(unsafe.Add(mBase, uint32(v419)+224))
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v419)+216))
	v628 = *(*int64)(unsafe.Add(mBase, uint32(v419)+208))
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v419)+200))
	v630 = *(*int64)(unsafe.Add(mBase, uint32(v419)+192))
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v419)+184))
	v632 = *(*int64)(unsafe.Add(mBase, uint32(v419)+176))
	v633 = *(*int64)(unsafe.Add(mBase, uint32(v419)+168))
	v634 = *(*int64)(unsafe.Add(mBase, uint32(v419)+160))
	v635 = *(*int64)(unsafe.Add(mBase, uint32(v419)+152))
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v419)+144))
	v637 = *(*int64)(unsafe.Add(mBase, uint32(v419)+136))
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v419)+128))
	v639 = *(*int64)(unsafe.Add(mBase, uint32(v419)+120))
	v640 = *(*float64)(unsafe.Add(mBase, uint32(v419)+112))
	v641 = *(*float64)(unsafe.Add(mBase, uint32(v419)+104))
	v642 = *(*int64)(unsafe.Add(mBase, uint32(v419)+96))
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v419)+88))
	v644 = *(*int64)(unsafe.Add(mBase, uint32(v419)+80))
	v645 = *(*int64)(unsafe.Add(mBase, uint32(v419)+72))
	v646 = *(*int64)(unsafe.Add(mBase, uint32(v419)+64))
	v647 = *(*int64)(unsafe.Add(mBase, uint32(v419)+56))
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v419)+48))
	v649 = *(*int64)(unsafe.Add(mBase, uint32(v419)+40))
	v650 = *(*int64)(unsafe.Add(mBase, uint32(v419)+24))
	v651 = *(*int64)(unsafe.Add(mBase, uint32(v419)+32))
	v652 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v419)+440)), uint32(v652))
	if int64(0)-v651 != v650 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L149
L151:
	;
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v419)+432))
	v659 = *(*int64)(unsafe.Add(mBase, uint32(v419)+424))
	v661 = v75 + int32(336)
	v663 = v594 + int32(1)
	v666 = v661 + v663<<(uint(int32(3))%32)
	if v341 != 0 {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	goto L153
L153:
	;
	v886 = F_hash_seq_search(m, v75+int32(776))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L200
	}
L154:
	;
	v725 = v722<<(uint(int32(3))%32) + v661
	*(*int64)(unsafe.Add(mBase, uint32(v725))) = v639
	*(*int64)(unsafe.Add(mBase, uint32(v725)+16)) = v637
	*(*int64)(unsafe.Add(mBase, uint32(v725)+8)) = v638
	if v339 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L155:
	;
	v705 = v75 + int32(336) + v700<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v705))) = v646
	*(*int64)(unsafe.Add(mBase, uint32(v705)+16)) = v642
	*(*int64)(unsafe.Add(mBase, uint32(v705)+8)) = v644
	if int64(2) <= v651 {
		goto L163
	} else {
		goto L164
	}
L156:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v666))) = v650
	*(*int64)(unsafe.Add(mBase, uint32(v666)+8)) = v649
	v671 = v594<<(uint(int32(3))%32) + v661
	*(*int64)(unsafe.Add(mBase, uint32(v671)+24)) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v671)+40)) = v643
	*(*int64)(unsafe.Add(mBase, uint32(v671)+32)) = v645
	if int64(2) <= v650 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v666))) = v651
	*(*int64)(unsafe.Add(mBase, uint32(v666)+8)) = v648
	v695 = v594 + int32(3)
	if v347|v341 == int32(0) {
		v722 = v695
		goto L154
	} else {
		goto L162
	}
L159:
	;
	v681 = base.F64_sqrt(base.F64_div(v641, base.F64_convert_i64_u(v650)))
	goto L161
L160:
	;
	v681 = float64(0)
	goto L161
L161:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v671)+48)) = v681
	v684 = v594 + int32(7)
	v687 = v684<<(uint(int32(3))%32) + v661
	*(*int64)(unsafe.Add(mBase, uint32(v687))) = v651
	*(*int64)(unsafe.Add(mBase, uint32(v687)+8)) = v648
	v699 = v684
	v700 = v594 + int32(9)
	goto L155
L162:
	;
	v699 = v663
	v700 = v695
	goto L155
L163:
	;
	v715 = base.F64_sqrt(base.F64_div(v640, base.F64_convert_i64_u(v651)))
	goto L165
L164:
	;
	v715 = float64(0)
	goto L165
L165:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v705)+24)) = v715
	v722 = v699 + int32(6)
	goto L154
L166:
	;
	if v343 != 0 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v725)+40)) = v633
	*(*int64)(unsafe.Add(mBase, uint32(v725)+32)) = v634
	*(*int64)(unsafe.Add(mBase, uint32(v725)+24)) = v635
	*(*int64)(unsafe.Add(mBase, uint32(v725-int32(-64)))) = v629
	*(*int64)(unsafe.Add(mBase, uint32(v725)+56)) = v630
	*(*int64)(unsafe.Add(mBase, uint32(v725)+48)) = v631
	v755 = v722 + int32(9)
	goto L166
L168:
	;
	goto L169
L169:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v725)+24)) = v636
	*(*int64)(unsafe.Add(mBase, uint32(v725)+96)) = v627
	*(*int64)(unsafe.Add(mBase, uint32(v725)+88)) = v628
	*(*int64)(unsafe.Add(mBase, uint32(v725)+80)) = v629
	*(*int64)(unsafe.Add(mBase, uint32(v725)+72)) = v630
	*(*int64)(unsafe.Add(mBase, uint32(v725-int32(-64)))) = v631
	*(*int64)(unsafe.Add(mBase, uint32(v725)+56)) = v632
	*(*int64)(unsafe.Add(mBase, uint32(v725)+48)) = v633
	*(*int64)(unsafe.Add(mBase, uint32(v725)+40)) = v634
	*(*int64)(unsafe.Add(mBase, uint32(v725)+32)) = v635
	v755 = v722 + int32(13)
	goto L166
L170:
	;
	v760 = v75 + int32(336) + v755<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v760))) = v626
	*(*int64)(unsafe.Add(mBase, uint32(v760)+8)) = v625
	v765 = v755 + int32(2)
	goto L172
L171:
	;
	v765 = v755
	goto L172
L172:
	;
	if v344 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v771 = v75 + int32(336) + v765<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v771))) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v771)+8)) = v623
	v776 = v765 + int32(2)
	goto L175
L174:
	;
	v776 = v765
	goto L175
L175:
	;
	if v341 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v782 = v75 + int32(336) + v776<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v782))) = v622
	*(*int64)(unsafe.Add(mBase, uint32(v782)+8)) = v621
	*(*int64)(unsafe.Add(mBase, uint32(v75))) = v620
	v790 = F_pg_snprintf(m, v75+int32(16), int32(256), int32(_a_F_pg_stat_statements_internal_6), v75)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	v801 = v776
	goto L178
L178:
	;
	if v345 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v796 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), base.I64_extend_i32_u(v75+int32(16)), int64(0), int64(-1))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v782)+16)) = v796
	v801 = v776 + int32(3)
	goto L178
L181:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v75+int32(336)+v801<<(uint(int32(3))%32)))) = v619
	v811 = v801 + int32(1)
	goto L183
L182:
	;
	v811 = v801
	goto L183
L183:
	;
	if v344 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v816 = v75 + int32(336) + v811<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v816))) = v618
	*(*int64)(unsafe.Add(mBase, uint32(v816)+56)) = v609
	*(*int64)(unsafe.Add(mBase, uint32(v816)+48)) = v610
	*(*int64)(unsafe.Add(mBase, uint32(v816)+40)) = v611
	*(*int64)(unsafe.Add(mBase, uint32(v816)+32)) = v612
	*(*int64)(unsafe.Add(mBase, uint32(v816)+24)) = v613
	*(*int64)(unsafe.Add(mBase, uint32(v816)+16)) = v616
	*(*int64)(unsafe.Add(mBase, uint32(v816)+8)) = v617
	v827 = v811 + int32(8)
	goto L186
L185:
	;
	v827 = v811
	goto L186
L186:
	;
	if v343 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v833 = v75 + int32(336) + v827<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v833))) = v614
	*(*int64)(unsafe.Add(mBase, uint32(v833)+8)) = v615
	v838 = v827 + int32(2)
	goto L189
L188:
	;
	v838 = v827
	goto L189
L189:
	;
	if v345 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v844 = v75 + int32(336) + v838<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v844))) = v608
	*(*int64)(unsafe.Add(mBase, uint32(v844)+8)) = v607
	v849 = v838 + int32(2)
	goto L192
L191:
	;
	v849 = v838
	goto L192
L192:
	;
	if v346 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v855 = v75 + int32(336) + v849<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v855))) = v606
	*(*int64)(unsafe.Add(mBase, uint32(v855)+8)) = v605
	v860 = v849 + int32(2)
	goto L195
L194:
	;
	v860 = v849
	goto L195
L195:
	;
	if v343 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v866 = v75 + int32(336) + v860<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v866))) = v659
	*(*int64)(unsafe.Add(mBase, uint32(v866)+8)) = v658
	goto L198
L197:
	;
	goto L198
L198:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	F_tuplestore_putvalues(m, v870, v871, v75+int32(336), v75+int32(272))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	goto L153
L200:
	;
	if v886 != 0 {
		v419 = v886
		goto L116
	} else {
		goto L201
	}
L201:
	;
	goto L117
L202:
	;
	if v401 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	F_pfree(m, v401)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	m.G0 = v75 + int32(800)
	return
L206:
	;
	goto L205
}
func F_pg_stat_statements_reset_1_7(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v6 = F_entry_reset(m, v2, v3, v4, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
