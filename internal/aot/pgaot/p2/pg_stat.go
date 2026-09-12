package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_stat_get_autovacuum_count(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+144))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_backend_client_addr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(256)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pgstat_get_beentry_by_proc_number(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(256)
	return v184
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
	v184 = v2
	goto L1
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v25 = F_has_privs_of_role(m, v23, int32(3375))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v35 = v14 + int32(56)
	if v35&int32(3) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	if v25 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v30 = F_has_privs_of_role(m, v28, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v30 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
	v184 = v2
	goto L1
L12:
	;
	v180 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v180)
	v184 = int32(0)
	goto L1
L13:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	switch v148 - int32(2) {
	case 0, 8:
		goto L40
	default:
		goto L41
	}
L14:
	;
	v63 = v35 + (int32(-56)-v14)&int32(3)
	v65 = v14 + int32(188)
	v67 = v65 & int32(-4)
	v69 = v67 - int32(28)
	if base.Ui32(v63) < base.Ui32(v69) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+56)))
	if v40 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if (v14+int32(57))&int32(3) == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+57)))
	if v47 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if (v14+int32(58))&int32(3) == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+58)))
	if v54 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	if (v14-int32(1))&int32(3) != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	v72 = v63
	goto L25
L23:
	;
	v98 = v63
	goto L24
L24:
	;
	if base.Ui32(v98) < base.Ui32(v67) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v79|(v80|(v81|(v82|(v83|(v84|(v85|v86)))))) != 0 {
		goto L13
	} else {
		goto L27
	}
L26:
	;
	v98 = v95
	goto L24
L27:
	;
	v95 = v72 + int32(32)
	if base.Ui32(v95) < base.Ui32(v69) {
		v72 = v95
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v107 = v98
	goto L32
L30:
	;
	v119 = v98
	goto L31
L31:
	;
	v127 = v119
	goto L36
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v114 != 0 {
		goto L13
	} else {
		goto L34
	}
L33:
	;
	v119 = v116
	goto L31
L34:
	;
	v116 = v107 + int32(4)
	if base.Ui32(v116) < base.Ui32(v67) {
		v107 = v116
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	if base.Ui32(v65) <= base.Ui32(v127) {
		goto L12
	} else {
		goto L38
	}
L37:
	;
	goto L13
L38:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v135 == int32(0) {
		v127 = v127 + int32(1)
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v154)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v14)+184))
	v162 = F_pg_getnameinfo_all(m, v35, v157, v11, int32(255), v154, v154, int32(3))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L42
	}
L41:
	;
	v151 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v151)
	v184 = int32(0)
	goto L1
L42:
	;
	if v162 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v164 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v164)
	v184 = v154
	goto L1
L44:
	;
	goto L45
L45:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	if v166 != int32(10) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v178 = F_DirectFunctionCall1Coll(m, int32(1482), int32(0), v11)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L50
	}
L47:
	;
	goto L46
L48:
	;
	v170 = F_strchr(m, v11, int32(37))
	mBase = m.M
	if v170 == int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v173)
	goto L47
L50:
	;
	v184 = v178
	goto L1
}
func F_pg_stat_get_backend_dbid(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pgstat_get_beentry_by_proc_number(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+48))
			return v14
		}
	}
}
func F_pg_stat_get_blocks_fetched(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+104))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_blocks_hit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+112))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_checkpointer_restartpoints_timed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+24))
		v7 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_pg_stat_get_checkpointer_slru_written(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+72))
		v7 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_pg_stat_get_checkpointer_stat_reset_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+80))
		v7 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_pg_stat_get_db_conflict_tablespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+80))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_sessions_killed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+232))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_tuples_returned(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+32))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_io(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v2 = int32(0)
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_pgstat_snapshot_fixed(m, int32(10))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v2
	goto L4
L4:
	;
	if int32(base.Ui32(int32(130554))>>(uint(v18)%32))&base.B2i32(base.Ui32(v18) < base.Ui32(int32(17))) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v29 = *(*int64)(unsafe.Add(mBase, _consts[1131]))
	F_pg_stat_io_build_tuples(m, v10, int32(4478328)+v18*int32(2880), v18, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v33 = v18 + int32(1)
	if v33 != int32(18) {
		v18 = v33
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	goto L5
}
func F_pg_stat_get_last_autoanalyze_time(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+168))
			if v9 != int64(0) {
				v17 = F_Int64GetDatum(m, v9)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		}
	}
}
func F_pg_stat_get_last_vacuum_time(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+120))
			if v9 != int64(0) {
				v17 = F_Int64GetDatum(m, v9)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		}
	}
}
func F_pg_stat_get_total_autoanalyze_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Float8GetDatum(m, float64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+208))
			v15 = F_Float8GetDatum(m, base.F64_convert_i64_s(v13))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_pg_stat_get_total_vacuum_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Float8GetDatum(m, float64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+184))
			v15 = F_Float8GetDatum(m, base.F64_convert_i64_s(v13))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_pg_stat_get_wal(m *base.Module, l0 int32) int32 {
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
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	F_pgstat_snapshot_fixed(m, int32(12))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(4530680)
		v14 = *(*int64)(unsafe.Add(mBase, _consts[1135]))
		v15 = *(*int64)(unsafe.Add(mBase, _consts[1136]))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v15
		v17 = *(*int64)(unsafe.Add(mBase, _consts[1137]))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v17
		v19 = *(*int64)(unsafe.Add(mBase, _consts[1138]))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v19
		v21 = *(*int64)(unsafe.Add(mBase, _consts[1139]))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v21
		v23 = F_pg_stat_wal_build_tuple(m, v6, v14)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(32)
			return v23
		}
	}
}
func F_pg_stat_get_xact_function_self_time(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_find_funcstat_entry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)+16))
			v18 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v14), float64(1e+06)))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
func F_pg_stat_get_xact_function_total_time(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_find_funcstat_entry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
			v18 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v14), float64(1e+06)))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
func F_pg_stat_get_xact_tuples_fetched(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+32))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_xact_tuples_returned(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+24))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_xact_tuples_updated(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+48))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_io_build_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) {
	mBase := m.M
	_ = mBase
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v135 int32
	_ = v135
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int64
	_ = v360
	var v398 int32
	_ = v398
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v452 int32
	_ = v452
	var v503 int32
	_ = v503
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v568 int32
	_ = v568
	var v631 int32
	_ = v631
	var v647 int32
	_ = v647
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v684 int32
	_ = v684
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v707 int64
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v719 int64
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v729 int64
	_ = v729
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	v60 = m.G0
	v62 = v60 - int32(384)
	m.G0 = v62
	v85 = v62 + int32(300)
	v96 = int32(12)
	v105 = v62 + int32(272)
	v135 = v62 + int32(271)
	if base.Ui32(l2) <= base.Ui32(int32(17)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v174 = F_cstring_to_text(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[1132])))
	v173 = v172
	goto L4
L3:
	;
	v173 = int32(385464)
	goto L4
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v186 = int32(0)
	goto L7
L7:
	;
	v235 = m.G0
	v237 = v235 - int32(16)
	m.G0 = v237
	if base.Ui32(int32(3)) <= base.Ui32(v186) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v62 + int32(384)
	return
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v186<<(uint(int32(2))%32))+uint32(_consts[1133])))
	m.G0 = v237 + int32(16)
	v263 = v186 * int32(320)
	v271 = int32(0)
	goto L15
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v186
	F_errmsg_internal(m, int32(505282), v237)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(519156), int32(273), int32(396236))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v325 = m.G0
	v327 = v325 - int32(16)
	m.G0 = v327
	if base.Ui32(int32(5)) <= base.Ui32(v271) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v832 = v186 + int32(1)
	if v832 != int32(3) {
		v186 = v832
		goto L7
	} else {
		goto L144
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v271<<(uint(int32(2))%32))+uint32(_consts[1134])))
	m.G0 = v327 + int32(16)
	v357 = F__emscripten_memset_bulkmem(m, v62+int32(304), base.I32_extend8_s(int32(0)), int32(80))
	mBase = m.M
	goto L23
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v271
	F_errmsg_internal(m, int32(505249), v327)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(519156), int32(256), int32(396051))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v358 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+288)) = v358
	v360 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v62)+280)) = v360
	*(*int64)(unsafe.Add(mBase, uint32(v62)+272)) = v360
	if base.Ui32(int32(16)) < base.Ui32(l2) {
		v428 = v358
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v428 != 0 {
		goto L46
	} else {
		goto L47
	}
L25:
	;
	if int32(1)<<(uint(l2)%32)&int32(130554) == int32(0) {
		v428 = v358
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if base.B2i32(v186 == int32(2))&base.B2i32(base.Ui32(v271-int32(4)) < base.Ui32(int32(-2))) != 0 {
		v428 = v358
		goto L24
	} else {
		goto L27
	}
L27:
	;
	if base.B2i32(v186 == int32(1))&base.B2i32(v271 != int32(3)) != 0 {
		v428 = v358
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(int32(16)) < base.Ui32(l2) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v398 = l2 & int32(-2)
	if v186 != int32(2) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	if int32(1)<<(uint(l2)%32)&int32(126232) == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v186 != int32(1) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v271 == int32(3) {
		v428 = v358
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	if base.B2i32(l2 == int32(16))|base.B2i32(v398 == int32(14)) != 0 {
		v428 = v358
		goto L24
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v398 == int32(10) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L36
L38:
	;
	v428 = base.B2i32(v271 != int32(1)) | base.B2i32(base.Ui32(l2-int32(5)) < base.Ui32(int32(-2)))
	goto L24
L39:
	;
	if base.Ui32(int32(4)) < base.Ui32(v271) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if l2 != int32(3) {
		goto L38
	} else {
		goto L44
	}
L42:
	;
	if int32(1)<<(uint(v271)%32)&int32(19) == int32(0) {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v428 = v358
	goto L24
L44:
	;
	if v271 == int32(4) {
		v428 = v358
		goto L24
	} else {
		goto L45
	}
L45:
	;
	goto L38
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+304)) = v174
	v430 = F_cstring_to_text(m, v348)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v828 = v271 + int32(1)
	if v828 != int32(5) {
		v271 = v828
		goto L15
	} else {
		goto L143
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+312)) = v430
	v433 = F_cstring_to_text(m, v258)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+308)) = v433
	if l3 != int64(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v452 = int32(0)
	goto L56
L52:
	;
	v438 = F_Int64GetDatum(m, l3)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L5
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v441 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+291)) = uint8(v441)
	goto L51
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+380)) = v438
	goto L51
L56:
	;
	v503 = int32(0)
	switch v452 - int32(1) {
	case 0:
		goto L64
	case 1:
		goto L68
	case 2:
		goto L67
	case 3:
		goto L66
	case 4:
		goto L65
	case 5:
		v527 = v62 + int32(304) | v96
		v528 = v62 + int32(324)
		v529 = v105 | int32(5)
		v530 = v105 | int32(3)
		v531 = v503
		v532 = v105 | int32(4)
		v533 = v62 + int32(320)
		v534 = v503
		goto L58
	case 6:
		goto L62
	default:
		goto L63
	}
L57:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_tuplestore_putvalues(m, v760, v761, v62+int32(304), v62+int32(272))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L5
	} else {
		goto L142
	}
L58:
	;
	if base.Ui32(int32(16)) < base.Ui32(l2) {
		goto L72
	} else {
		goto L73
	}
L59:
	;
	v527 = v523
	v528 = v85
	v529 = v135
	v530 = v524
	v531 = v525
	v532 = v135
	v533 = v85
	v534 = int32(1)
	goto L58
L60:
	;
	v527 = v518
	v528 = v519
	v529 = v520
	v530 = v521
	v531 = v503
	v532 = v135
	v533 = v85
	v534 = v522
	goto L58
L61:
	;
	v527 = v512
	v528 = v513
	v529 = v514
	v530 = v517
	v531 = v503
	v532 = v515
	v533 = v516
	v534 = v503
	goto L58
L62:
	;
	v512 = v62 + int32(328)
	v513 = v62 + int32(336)
	v514 = v105 | int32(8)
	v515 = v105 | int32(7)
	v516 = v62 + int32(332)
	v517 = v105 | int32(6)
	goto L61
L63:
	;
	v523 = v62 + int32(364)
	v524 = v105 | int32(15)
	v525 = int32(1)
	goto L59
L64:
	;
	v518 = v62 + int32(372)
	v519 = v62 + int32(376)
	v520 = v62 + int32(290)
	v521 = v62 + int32(289)
	v522 = int32(1)
	goto L60
L65:
	;
	v512 = v62 + int32(348)
	v513 = v62 + int32(356)
	v514 = v105 | int32(13)
	v515 = v105 | v96
	v516 = v62 + int32(352)
	v517 = v105 | int32(11)
	goto L61
L66:
	;
	v518 = v62 + int32(340)
	v519 = v62 + int32(344)
	v520 = v105 | int32(10)
	v521 = v105 | int32(9)
	v522 = int32(1)
	goto L60
L67:
	;
	v523 = v62 + int32(368)
	v524 = v62 + int32(288)
	v525 = int32(1)
	goto L59
L68:
	;
	v523 = v62 + int32(360)
	v524 = v105 | int32(14)
	v525 = int32(1)
	goto L59
L69:
	;
	v757 = v452 + int32(1)
	if v757 != int32(8) {
		v452 = v757
		goto L56
	} else {
		goto L141
	}
L70:
	;
	if v531 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L71:
	;
	if v696 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L72:
	;
	v696 = int32(0)
	goto L71
L73:
	;
	if int32(1)<<(uint(l2)%32)&int32(130554) == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	if base.B2i32(v186 == int32(2))&base.B2i32(base.Ui32(v271-int32(4)) < base.Ui32(int32(-2))) != 0 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	if base.B2i32(v186 == int32(1))&base.B2i32(v271 != int32(3)) != 0 {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	if base.Ui32(int32(16)) < base.Ui32(l2) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v568 = l2 & int32(-2)
	if v186 != int32(2) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	if int32(1)<<(uint(l2)%32)&int32(126232) == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	if v186 != int32(1) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	if v271 == int32(3) {
		goto L72
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	if base.B2i32(l2 == int32(16))|base.B2i32(v568 == int32(14)) != 0 {
		goto L72
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v568 == int32(10) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L84
L86:
	;
	switch l2 - int32(10) {
	case 0:
		goto L97
	case 1:
		goto L96
	default:
		goto L95
	}
L87:
	;
	if base.Ui32(int32(4)) < base.Ui32(v271) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if base.B2i32(l2 == int32(3))&base.B2i32(v271 == int32(4)) != 0 {
		goto L72
	} else {
		goto L92
	}
L90:
	;
	if int32(1)<<(uint(v271)%32)&int32(19) == int32(0) {
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L72
L92:
	;
	if base.Ui32(l2-int32(5)) < base.Ui32(int32(-2)) {
		goto L86
	} else {
		goto L93
	}
L93:
	;
	if v271 == int32(1) {
		goto L72
	} else {
		goto L94
	}
L94:
	;
	goto L86
L95:
	;
	if base.B2i32(base.B2i32(l2 != int32(3))&base.B2i32(v568 != int32(10)) == int32(0))&base.B2i32(v452 == int32(5)) != 0 {
		goto L72
	} else {
		goto L102
	}
L96:
	;
	if base.B2i32(v452 == int32(6))&base.B2i32(v186 != int32(2)) != 0 {
		goto L72
	} else {
		goto L100
	}
L97:
	;
	if base.Ui32(int32(6)) < base.Ui32(v452) {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	if int32(1)<<(uint(v452)%32)&int32(69) == int32(0) {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	goto L72
L100:
	;
	if v452&int32(-3) == int32(0) {
		goto L72
	} else {
		goto L101
	}
L101:
	;
	goto L95
L102:
	;
	if v186 != int32(2) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v669 = base.B2i32(v186 != int32(2))
	if v186 != int32(2) {
		goto L118
	} else {
		goto L119
	}
L104:
	;
	v667 = base.B2i32(v271 == int32(4)) | base.B2i32(base.Ui32(v271) < base.Ui32(int32(2)))
	goto L103
L105:
	;
	if int32(base.Ui32(int32(10371))>>(uint(v631)%32))&int32(1) != 0 {
		goto L72
	} else {
		goto L116
	}
L106:
	;
	if v186 != int32(1) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	if v452 != int32(6) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v631 = l2 - int32(3)
	if base.Ui32(v631) < base.Ui32(int32(14)) {
		goto L105
	} else {
		goto L109
	}
L109:
	;
	goto L104
L110:
	;
	if base.B2i32(v271 == int32(0))&base.B2i32(v452 == int32(5)) != 0 {
		goto L72
	} else {
		goto L112
	}
L111:
	;
	switch v452 - int32(1) {
	case 0, 3:
		goto L72
	default:
		goto L110
	}
L112:
	;
	v647 = base.B2i32(v271 == int32(4)) | base.B2i32(base.Ui32(v271) < base.Ui32(int32(2)))
	if v452 != int32(3) {
		v667 = v647
		goto L103
	} else {
		goto L113
	}
L113:
	;
	if base.Ui32(int32(4)) < base.Ui32(v271) {
		goto L72
	} else {
		goto L114
	}
L114:
	;
	if int32(1)<<(uint(v271)%32)&int32(19) == int32(0) {
		goto L72
	} else {
		goto L115
	}
L115:
	;
	v667 = v647
	goto L103
L116:
	;
	goto L104
L117:
	;
	v684 = int32(1)
	v696 = v667 ^ v684 | base.B2i32(v452 != v684)
	goto L71
L118:
	;
	if v186 != int32(2) {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	if v271 != int32(2) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	switch v452 - int32(1) {
	case 0, 6:
		goto L117
	default:
		goto L72
	}
L121:
	;
	if v271 != int32(3) {
		goto L117
	} else {
		goto L122
	}
L122:
	;
	if base.Ui32(int32(7)) < base.Ui32(v452) {
		goto L72
	} else {
		goto L123
	}
L123:
	;
	if int32(1)<<(uint(v452)%32)&int32(194) == int32(0) {
		goto L72
	} else {
		goto L124
	}
L124:
	;
	goto L117
L125:
	;
	v699 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v530))) = uint8(v699)
	goto L70
L126:
	;
	goto L127
L127:
	;
	v702 = v452 << (uint(int32(3)) % 32)
	v704 = v271 << (uint(int32(6)) % 32)
	v707 = *(*int64)(unsafe.Add(mBase, uint32(v702+(l1+int32(960)+v263+v704))))
	v708 = F_Int64GetDatum(m, v707)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v527))) = v708
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	if v711&int32(1) != 0 {
		goto L70
	} else {
		goto L129
	}
L129:
	;
	if v531 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v719 = *(*int64)(unsafe.Add(mBase, uint32(v263+(l1+int32(1920))+v704+v702)))
	v723 = F_Float8GetDatum(m, base.F64_mul(base.F64_convert_i64_s(v719), float64(0.001)))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L5
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if v534 != 0 {
		goto L69
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = v723
	goto L132
L134:
	;
	v729 = *(*int64)(unsafe.Add(mBase, uint32(l1+v263+v704+v702)))
	*(*int64)(unsafe.Add(mBase, uint32(v62))) = v729
	v735 = F_pg_snprintf(m, v62+int32(16), int32(256), int32(450270), v62)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L5
	} else {
		goto L135
	}
L135:
	;
	v738 = int32(0)
	v743 = F_DirectFunctionCall3Coll(m, int32(408), v738, v62+int32(16), v738, int32(-1))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v743
	goto L69
L137:
	;
	v750 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v529))) = uint8(v750)
	goto L139
L138:
	;
	goto L139
L139:
	;
	if v534 != 0 {
		goto L69
	} else {
		goto L140
	}
L140:
	;
	v752 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v532))) = uint8(v752)
	goto L69
L141:
	;
	goto L57
L142:
	;
	goto L48
L143:
	;
	goto L16
L144:
	;
	goto L8
}
func F_pg_stat_reset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	F___gettimeofday(m, v8)
	mBase = m.M
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	m.G0 = v8 + v7
	v23 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	F_pgstat_reset_matching_entries(m, int32(1214), v23, v12+v11*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_reset_single_table_counters(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v46 int32
	_ = v46
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(1)
	if v3 <= int32(3591) {
		if v3 <= int32(2670) {
			switch v3 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v74 = v6
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v74 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(v3-int32(2396)) {
					v74 = int32(0)
				} else {
					v74 = v6
				}
			}
		} else {
			v18 = v3 - int32(2671)
			if base.Ui32(int32(27)) < base.Ui32(v18) {
				if base.Ui32(v3-int32(2964)) < base.Ui32(int32(4)) {
					v74 = v6
				} else {
					if base.Ui32(v3-int32(2846)) < base.Ui32(int32(2)) {
						v74 = v6
					} else {
						v74 = int32(0)
					}
				}
			} else {
				if int32(1)<<(uint(v18)%32)&int32(226492515) == int32(0) {
					if base.Ui32(v3-int32(2964)) < base.Ui32(int32(4)) {
						v74 = v6
					} else {
						if base.Ui32(v3-int32(2846)) < base.Ui32(int32(2)) {
							v74 = v6
						} else {
							v74 = int32(0)
						}
					}
				} else {
					v74 = v6
				}
			}
		}
	} else {
		if v3 <= int32(5999) {
			v30 = v3 - int32(4177)
			if base.Ui32(int32(9)) < base.Ui32(v30) {
				if base.Ui32(v3-int32(3592)) < base.Ui32(int32(2)) {
					v74 = v6
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(v3-int32(4060)) {
						v74 = int32(0)
					} else {
						v74 = v6
					}
				}
			} else {
				if int32(1)<<(uint(v30)%32)&int32(963) == int32(0) {
					if base.Ui32(v3-int32(3592)) < base.Ui32(int32(2)) {
						v74 = v6
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(v3-int32(4060)) {
							v74 = int32(0)
						} else {
							v74 = v6
						}
					}
				} else {
					v74 = v6
				}
			}
		} else {
			switch v3 - int32(6243) {
			case 0, 1, 2, 3, 4, 59, 60:
				v74 = v6
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v74 = int32(0)
			default:
				if base.Ui32(v3-int32(6000)) < base.Ui32(int32(3)) {
					v74 = v6
				} else {
					v46 = v3 - int32(6100)
					if base.Ui32(int32(15)) < base.Ui32(v46) {
						v74 = int32(0)
					} else {
						if int32(1)<<(uint(v46)%32)&int32(49153) != 0 {
							v74 = v6
						} else {
							v74 = int32(0)
						}
					}
				}
			}
		}
	}
	v78 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v74 != 0 {
		v79 = int32(0)
	} else {
		v79 = v78
	}
	F_pgstat_reset(m, int32(2), v79, base.I64_extend_i32_u(v3))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
