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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
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
	return v181
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
	v181 = v2
	goto L1
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_client_addr[0]))
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
	v39 = (int32(-56) - v14) & int32(3)
	if v39 == int32(0) {
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
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_client_addr[0]))
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
	v181 = v2
	goto L1
L12:
	;
	v177 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v177)
	v181 = int32(0)
	goto L1
L13:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	switch v145 - int32(2) {
	case 0, 8:
		goto L42
	default:
		goto L43
	}
L14:
	;
	v52 = v39 + v35
	v56 = (v14 + int32(188)) & int32(-4)
	v58 = v56 - int32(28)
	if base.Ui32(v52) < base.Ui32(v58) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v42 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v39 == int32(1) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+57)))
	if v45 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v39 == int32(2) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+58)))
	if v48|base.B2i32(v39 != int32(3)) != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	v61 = v52
	v62 = v39
	goto L24
L22:
	;
	v89 = v39
	goto L23
L23:
	;
	v95 = v89 + v35
	if base.Ui32(v95) < base.Ui32(v56) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v68|(v69|(v70|(v71|(v72|(v73|(v74|v75)))))) != 0 {
		goto L13
	} else {
		goto L26
	}
L25:
	;
	v89 = v84
	goto L23
L26:
	;
	v84 = v62 + int32(32)
	v85 = v35 + v84
	if base.Ui32(v85) < base.Ui32(v58) {
		v61 = v85
		v62 = v84
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v98 = v95
	v99 = v89
	goto L31
L29:
	;
	v112 = v89
	goto L30
L30:
	;
	v118 = int32(132)
	if base.Ui32(v112) <= base.Ui32(v118) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v105 != 0 {
		goto L13
	} else {
		goto L33
	}
L32:
	;
	v112 = v107
	goto L30
L33:
	;
	v107 = v99 + int32(4)
	v108 = v35 + v107
	if base.Ui32(v108) < base.Ui32(v56) {
		v98 = v108
		v99 = v107
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v121 = v118
	goto L37
L36:
	;
	v121 = v112
	goto L37
L37:
	;
	v124 = v112
	goto L38
L38:
	;
	if v121 == v124 {
		goto L12
	} else {
		goto L40
	}
L39:
	;
	goto L13
L40:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v35))))
	if v134 == int32(0) {
		v124 = v124 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v151)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v14)+184))
	v159 = F_pg_getnameinfo_all(m, v35, v154, v11, int32(255), v151, v151, int32(3))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L44
	}
L43:
	;
	v148 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v148)
	v181 = int32(0)
	goto L1
L44:
	;
	if v159 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v161 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v161)
	v181 = v151
	goto L1
L46:
	;
	goto L47
L47:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	if v163 != int32(10) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v175 = F_DirectFunctionCall1Coll(m, int32(1466), int32(0), v11)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L52
	}
L49:
	;
	goto L48
L50:
	;
	v167 = F_strchr(m, v11, int32(37))
	mBase = m.M
	if v167 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v170 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v170)
	goto L49
L52:
	;
	v181 = v175
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v2 = int32(0)
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_pgstat_snapshot_fixed(m, int32(10))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v2
	goto L4
L4:
	;
	if int32(base.Ui32(int32(_a_F_pg_stat_get_io_0))>>(uint(v13)%32))&base.B2i32(base.Ui32(v13) < base.Ui32(int32(17))) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v24 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_io[0]))
	F_pg_stat_io_build_tuples(m, v8, v13*int32(2880)+int32(_a_F_pg_stat_get_io_1), v13, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v28 = v13 + int32(1)
	if v28 != int32(18) {
		v13 = v28
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
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	F_pgstat_snapshot_fixed(m, int32(12))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[0]))
		v16 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v16
		v19 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v19
		v22 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v22
		v25 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v25
		v27 = F_pg_stat_wal_build_tuple(m, v6, v14)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(32)
			return v27
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v128 int32
	_ = v128
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v345 int64
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v367 int32
	_ = v367
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v440 int32
	_ = v440
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
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
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v546 int32
	_ = v546
	var v567 int32
	_ = v567
	var v629 int32
	_ = v629
	var v655 int32
	_ = v655
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v705 int64
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v713 int64
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int64
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	v65 = m.G0
	v67 = v65 - int32(384)
	m.G0 = v67
	v90 = v67 + int32(300)
	v99 = int32(12)
	v110 = v67 + int32(272)
	v128 = v67 + int32(271)
	if base.Ui32(l2) <= base.Ui32(int32(17)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v152 = F_cstring_to_text(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_io_build_tuples[0])))
	v151 = v149
	goto L4
L3:
	;
	v151 = int32(_a_F_pg_stat_io_build_tuples_0)
	goto L4
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v163 = int32(0)
	goto L7
L7:
	;
	v219 = v163 * int32(320)
	v221 = m.G0
	v223 = v221 - int32(16)
	m.G0 = v223
	if base.Ui32(int32(3)) <= base.Ui32(v163) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v67 + int32(384)
	return
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v163<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_io_build_tuples[1])))
	m.G0 = v223 + int32(16)
	v255 = int32(0)
	goto L15
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v163
	F_errmsg_internal(m, int32(_a_F_pg_stat_io_build_tuples_1), v223)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_pg_stat_io_build_tuples_2), int32(273), int32(_a_F_pg_stat_io_build_tuples_3))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
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
	v313 = m.G0
	v315 = v313 - int32(16)
	m.G0 = v315
	if base.Ui32(int32(5)) <= base.Ui32(v255) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v825 = v163 + int32(1)
	if v825 != int32(3) {
		v163 = v825
		goto L7
	} else {
		goto L118
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v332 = int32(2)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v255<<(uint(v332)%32))+uint32(_c_F_pg_stat_io_build_tuples[2])))
	v335 = int32(16)
	m.G0 = v315 + v335
	v340 = int32(0)
	base.MemoryFill(m, v67+int32(304), v340, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+288)) = v340
	v345 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v67)+280)) = v345
	*(*int64)(unsafe.Add(mBase, uint32(v67)+272)) = v345
	v351 = base.B2i32(base.Ui32(v335) < base.Ui32(l2))
	v352 = int32(1)
	v353 = v352 << (uint(l2) % 32)
	v367 = base.B2i32(v163 == v352)
	if v351|base.B2i32(v353&int32(_a_F_pg_stat_io_build_tuples_4) == v340)|(base.B2i32(v163 == v332)&base.B2i32(base.Ui32(v255-int32(4)) < base.Ui32(int32(-2)))|v367&base.B2i32(v255 != int32(3))) != 0 {
		v411 = v340
		goto L23
	} else {
		goto L24
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = v255
	F_errmsg_internal(m, int32(_a_F_pg_stat_io_build_tuples_5), v315)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_pg_stat_io_build_tuples_2), int32(256), int32(_a_F_pg_stat_io_build_tuples_6))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
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
	if v411 != 0 {
		goto L37
	} else {
		goto L38
	}
L24:
	;
	if v351|base.B2i32(v353&int32(_a_F_pg_stat_io_build_tuples_7) == int32(0)) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if l2&int32(-2) == int32(10) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	if base.B2i32(v255 == int32(3))&v367 != 0 {
		v411 = v340
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(int32(2)) < base.Ui32(l2-int32(14)) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if v163 != int32(2) {
		v411 = v340
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v411 = base.B2i32(v255 != int32(1)) | base.B2i32(base.Ui32(l2-int32(5)) < base.Ui32(int32(-2)))
	goto L23
L31:
	;
	if base.B2i32(int32(1)<<(uint(v255)%32)&int32(19) == int32(0))|base.B2i32(base.Ui32(int32(4)) < base.Ui32(v255)) != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if l2 != int32(3) {
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v411 = v340
	goto L23
L35:
	;
	if v255 == int32(4) {
		v411 = v340
		goto L23
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+304)) = v152
	v413 = F_cstring_to_text(m, v334)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v821 = v255 + int32(1)
	if v821 != int32(5) {
		v255 = v821
		goto L15
	} else {
		goto L117
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+312)) = v413
	v416 = F_cstring_to_text(m, v244)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+308)) = v416
	if l3 != int64(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v427 = v255 << (uint(int32(6)) % 32)
	v440 = int32(0)
	goto L47
L43:
	;
	v421 = F_Int64GetDatum(m, l3)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L5
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v424 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+291)) = uint8(v424)
	goto L42
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+380)) = v421
	goto L42
L47:
	;
	v496 = int32(0)
	switch v440 - int32(1) {
	case 0:
		goto L55
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	case 5:
		v520 = v67 + int32(304) | v99
		v521 = v67 + int32(324)
		v522 = v110 | int32(5)
		v523 = v110 | int32(3)
		v524 = v110 | int32(4)
		v525 = v496
		v526 = v67 + int32(320)
		v527 = v496
		goto L49
	case 6:
		goto L53
	default:
		goto L54
	}
L48:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_tuplestore_putvalues(m, v748, v749, v67+int32(304), v67+int32(272))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L5
	} else {
		goto L116
	}
L49:
	;
	v528 = int32(0)
	v530 = base.B2i32(base.Ui32(int32(16)) < base.Ui32(l2))
	v531 = int32(1)
	v532 = v531 << (uint(l2) % 32)
	v546 = base.B2i32(v163 == v531)
	if v530|base.B2i32(v532&int32(_a_F_pg_stat_io_build_tuples_4) == v528)|(base.B2i32(v163 == int32(2))&base.B2i32(base.Ui32(v255-int32(4)) < base.Ui32(int32(-2)))|v546&base.B2i32(v255 != int32(3))) != 0 {
		v697 = v528
		goto L60
	} else {
		goto L61
	}
L50:
	;
	v520 = v516
	v521 = v90
	v522 = v128
	v523 = v517
	v524 = v128
	v525 = v518
	v526 = v90
	v527 = int32(1)
	goto L49
L51:
	;
	v520 = v511
	v521 = v512
	v522 = v513
	v523 = v514
	v524 = v128
	v525 = v496
	v526 = v90
	v527 = v515
	goto L49
L52:
	;
	v520 = v505
	v521 = v506
	v522 = v507
	v523 = v510
	v524 = v508
	v525 = v496
	v526 = v509
	v527 = v496
	goto L49
L53:
	;
	v505 = v67 + int32(328)
	v506 = v67 + int32(336)
	v507 = v110 | int32(8)
	v508 = v110 | int32(7)
	v509 = v67 + int32(332)
	v510 = v110 | int32(6)
	goto L52
L54:
	;
	v516 = v67 + int32(364)
	v517 = v110 | int32(15)
	v518 = int32(1)
	goto L50
L55:
	;
	v511 = v67 + int32(372)
	v512 = v67 + int32(376)
	v513 = v67 + int32(290)
	v514 = v67 + int32(289)
	v515 = int32(1)
	goto L51
L56:
	;
	v505 = v67 + int32(348)
	v506 = v67 + int32(356)
	v507 = v110 | int32(13)
	v508 = v110 | v99
	v509 = v67 + int32(352)
	v510 = v110 | int32(11)
	goto L52
L57:
	;
	v511 = v67 + int32(340)
	v512 = v67 + int32(344)
	v513 = v110 | int32(10)
	v514 = v110 | int32(9)
	v515 = int32(1)
	goto L51
L58:
	;
	v516 = v67 + int32(368)
	v517 = v67 + int32(288)
	v518 = int32(1)
	goto L50
L59:
	;
	v516 = v67 + int32(360)
	v517 = v110 | int32(14)
	v518 = int32(1)
	goto L50
L60:
	;
	if v697 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L61:
	;
	if v530|base.B2i32(v532&int32(_a_F_pg_stat_io_build_tuples_7) == int32(0)) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v567 = l2 & int32(-2)
	if v567 == int32(10) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	if base.B2i32(v255 == int32(3))&v546 != 0 {
		v697 = v528
		goto L60
	} else {
		goto L64
	}
L64:
	;
	if base.Ui32(int32(2)) < base.Ui32(l2-int32(14)) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	if v163 != int32(2) {
		v697 = v528
		goto L60
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	switch l2 - int32(10) {
	case 0:
		goto L77
	case 1:
		goto L76
	default:
		goto L75
	}
L68:
	;
	if base.B2i32(int32(1)<<(uint(v255)%32)&int32(19) == int32(0))|base.B2i32(base.Ui32(int32(4)) < base.Ui32(v255)) != 0 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if base.B2i32(l2 == int32(3))&base.B2i32(v255 == int32(4)) != 0 {
		v697 = v528
		goto L60
	} else {
		goto L72
	}
L71:
	;
	v697 = v528
	goto L60
L72:
	;
	if base.Ui32(l2-int32(5)) < base.Ui32(int32(-2)) {
		goto L67
	} else {
		goto L73
	}
L73:
	;
	if v255 == int32(1) {
		v697 = v528
		goto L60
	} else {
		goto L74
	}
L74:
	;
	goto L67
L75:
	;
	if base.B2i32(base.B2i32(l2 != int32(3))&base.B2i32(v567 != int32(10)) == int32(0))&base.B2i32(v440 == int32(5)) != 0 {
		v697 = v528
		goto L60
	} else {
		goto L80
	}
L76:
	;
	if base.B2i32(v440&int32(-3) == int32(0))|base.B2i32(v440 == int32(6))&base.B2i32(v163 != int32(2)) != 0 {
		v697 = v528
		goto L60
	} else {
		goto L79
	}
L77:
	;
	if base.B2i32(int32(1)<<(uint(v440)%32)&int32(69) == int32(0))|base.B2i32(base.Ui32(int32(6)) < base.Ui32(v440)) != 0 {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v697 = v528
	goto L60
L79:
	;
	goto L75
L80:
	;
	if base.B2i32(v163 != int32(2))|base.B2i32(v440 != int32(6)) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v669 = int32(2)
	v670 = base.B2i32(v163 != v669)
	if v670|base.B2i32(v255 != v669) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L82:
	;
	v629 = l2 - int32(3)
	if base.B2i32(base.Ui32(v629) <= base.Ui32(int32(13)))&(int32(base.Ui32(int32(_a_F_pg_stat_io_build_tuples_8))>>(uint(v629)%32))&int32(1)) != 0 {
		v697 = v528
		goto L60
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v163 != int32(1) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v668 = base.B2i32(v255 == int32(4)) | base.B2i32(base.Ui32(v255) < base.Ui32(int32(2)))
	goto L81
L86:
	;
	if base.B2i32(v255 == int32(0))&base.B2i32(v440 == int32(5)) != 0 {
		v697 = v528
		goto L60
	} else {
		goto L88
	}
L87:
	;
	switch v440 - int32(1) {
	case 0, 3:
		v697 = v528
		goto L60
	default:
		goto L86
	}
L88:
	;
	v655 = base.B2i32(v255 == int32(4)) | base.B2i32(base.Ui32(v255) < base.Ui32(int32(2)))
	if v440 != int32(3) {
		v668 = v655
		goto L81
	} else {
		goto L89
	}
L89:
	;
	if base.B2i32(int32(1)<<(uint(v255)%32)&int32(19) == int32(0))|base.B2i32(base.Ui32(int32(4)) < base.Ui32(v255)) != 0 {
		v697 = v528
		goto L60
	} else {
		goto L90
	}
L90:
	;
	v668 = v655
	goto L81
L91:
	;
	v697 = base.B2i32(v668 == int32(0)) | base.B2i32(v440 != int32(1))
	goto L60
L92:
	;
	switch v440 - int32(1) {
	case 0, 6:
		goto L91
	default:
		v697 = v528
		goto L60
	}
L93:
	;
	goto L94
L94:
	;
	if v670|base.B2i32(v255 != int32(3)) != 0 {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	if base.B2i32(int32(1)<<(uint(v440)%32)&int32(194) == int32(0))|base.B2i32(base.Ui32(int32(7)) < base.Ui32(v440)) != 0 {
		v697 = v528
		goto L60
	} else {
		goto L96
	}
L96:
	;
	goto L91
L97:
	;
	v745 = v440 + int32(1)
	if v745 != int32(8) {
		v440 = v745
		goto L47
	} else {
		goto L115
	}
L98:
	;
	if v525 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L99:
	;
	v700 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v523))) = uint8(v700)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v703 = v440 << (uint(int32(3)) % 32)
	v705 = *(*int64)(unsafe.Add(mBase, uint32(v427+(v219+(l1+int32(960)))+v703)))
	v706 = F_Int64GetDatum(m, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = v706
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	if v709 != 0 {
		goto L98
	} else {
		goto L103
	}
L103:
	;
	if v525 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v713 = *(*int64)(unsafe.Add(mBase, uint32(v703+(v427+(v219+(l1+int32(1920)))))))
	v717 = F_Float8GetDatum(m, base.F64_mul(base.F64_convert_i64_s(v713), float64(0.001)))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L5
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v527 != 0 {
		goto L97
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521))) = v717
	goto L106
L108:
	;
	v721 = *(*int64)(unsafe.Add(mBase, uint32(v703+(l1+v219+v427))))
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = v721
	v724 = v67 + int32(16)
	v727 = F_pg_snprintf(m, v724, int32(256), int32(_a_F_pg_stat_io_build_tuples_9), v67)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v730 = int32(0)
	v733 = F_DirectFunctionCall3Coll(m, int32(408), v730, v724, v730, int32(-1))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526))) = v733
	goto L97
L111:
	;
	v739 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v522))) = uint8(v739)
	goto L113
L112:
	;
	goto L113
L113:
	;
	if v527 != 0 {
		goto L97
	} else {
		goto L114
	}
L114:
	;
	v741 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v741)
	goto L97
L115:
	;
	goto L48
L116:
	;
	goto L39
L117:
	;
	goto L16
L118:
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
	F_gettimeofday(m, v8)
	mBase = m.M
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	m.G0 = v8 + v7
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset[0]))
	F_pgstat_reset_matching_entries(m, int32(1198), v23, v12+v11*int64(1000000)-int64(946684800000000))
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
	var v31 int32
	_ = v31
	var v48 int32
	_ = v48
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(1)
	if v3 <= int32(3591) {
		if v3 <= int32(2670) {
			switch v3 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v77 = v6
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v77 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(v3-int32(2396)) {
					v77 = int32(0)
				} else {
					v77 = v6
				}
			}
		} else {
			v18 = v3 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v18))|base.B2i32(int32(1)<<(uint(v18)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(v3-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v3-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v77 = v6
				} else {
					v77 = int32(0)
				}
			} else {
				v77 = v6
			}
		}
	} else {
		if v3 <= int32(_a_F_pg_stat_reset_single_table_counters_0) {
			v31 = v3 - int32(_a_F_pg_stat_reset_single_table_counters_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v31))|base.B2i32(int32(1)<<(uint(v31)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(v3-int32(3592)) < base.Ui32(int32(2)) {
					v77 = v6
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(v3-int32(4060)) {
						v77 = int32(0)
					} else {
						v77 = v6
					}
				}
			} else {
				v77 = v6
			}
		} else {
			switch v3 - int32(_a_F_pg_stat_reset_single_table_counters_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v77 = v6
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v77 = int32(0)
			default:
				if base.Ui32(v3-int32(_a_F_pg_stat_reset_single_table_counters_3)) < base.Ui32(int32(3)) {
					v77 = v6
				} else {
					v48 = v3 - int32(_a_F_pg_stat_reset_single_table_counters_4)
					if base.Ui32(int32(15)) < base.Ui32(v48) {
						v77 = int32(0)
					} else {
						if int32(1)<<(uint(v48)%32)&int32(_a_F_pg_stat_reset_single_table_counters_5) != 0 {
							v77 = v6
						} else {
							v77 = int32(0)
						}
					}
				}
			}
		}
	}
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_single_table_counters[0]))
	if v77 != 0 {
		v82 = int32(0)
	} else {
		v82 = v81
	}
	F_pgstat_reset(m, int32(2), v82, base.I64_extend_i32_u(v3))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_statements_1_11(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_stat_statements_internal(m, l0, int32(7), base.B2i32(v3 != int32(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
