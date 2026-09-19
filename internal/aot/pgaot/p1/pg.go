package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_PGSemaphoreUnlock(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	Fn13857(m, l0, int32(_a_F_PGSemaphoreUnlock_0), int32(360), int32(_a_F_PGSemaphoreUnlock_1), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_SetPGVariable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v4 = F_flatten_set_variable_args(m, l0, l1)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v8 = F_superuser(m)
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			if v8 != 0 {
				v10 = int32(5)
			} else {
				v10 = int32(6)
			}
			F_set_config_option(m, l0, v4, v10, int32(13), l2, int32(1))
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F__PG_init_bloom(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = F_add_reloption_kind(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__PG_init_bloom[0])) = v11
	F_add_int_reloption(m, v11, int32(_a_F__PG_init_bloom_0), int32(_a_F__PG_init_bloom_1), int32(80), int32(1), int32(_a_F__PG_init_bloom_2))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, _c_F__PG_init_bloom[1])) = int64(17179869185)
	*(*int32)(unsafe.Add(mBase, _c_F__PG_init_bloom[2])) = int32(_a_F__PG_init_bloom_0)
	v28 = int32(0)
	goto L4
L4:
	;
	v34 = v28 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v34
	v36 = int32(16)
	v37 = v8 + v36
	v40 = F_pg_snprintf(m, v37, v36, int32(_a_F__PG_init_bloom_3), v8)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v8 + int32(32)
	return
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_bloom[0]))
	F_add_int_reloption(m, v43, v37, int32(_a_F__PG_init_bloom_4), int32(2), int32(1), int32(4095))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_bloom[3]))
	v52 = F_MemoryContextStrdup(m, v51, v37)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v55 = v34 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F__PG_init_bloom[4]))) = v28<<(uint(int32(2))%32) + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F__PG_init_bloom[1]))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F__PG_init_bloom[2]))) = v52
	if v34 != int32(32) {
		v28 = v34
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F__PG_init_pg_stat_statements(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[0])))
	if v3 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[1]))
		if v7 != 0 {
			v9 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[2])) = uint8(v9)
		} else {
		}
		v13 = int32(0)
		F_DefineCustomIntVariable(m, int32(_a_F__PG_init_pg_stat_statements_0), int32(_a_F__PG_init_pg_stat_statements_1), v13, int32(_a_F__PG_init_pg_stat_statements_2), int32(_a_F__PG_init_pg_stat_statements_3), int32(100), int32(1073741823), int32(1), v13)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_DefineCustomEnumVariable(m, int32(_a_F__PG_init_pg_stat_statements_4), int32(_a_F__PG_init_pg_stat_statements_5), int32(0), int32(_a_F__PG_init_pg_stat_statements_6), int32(1), int32(_a_F__PG_init_pg_stat_statements_7), int32(5))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_pg_stat_statements_8), int32(_a_F__PG_init_pg_stat_statements_9), int32(0), int32(_a_F__PG_init_pg_stat_statements_10), int32(1), int32(5))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v41 = int32(0)
					F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_pg_stat_statements_11), int32(_a_F__PG_init_pg_stat_statements_12), v41, int32(_a_F__PG_init_pg_stat_statements_13), v41, int32(5))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_pg_stat_statements_14), int32(_a_F__PG_init_pg_stat_statements_15), int32(0), int32(_a_F__PG_init_pg_stat_statements_16), int32(1), int32(2))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_pg_stat_statements_17))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								v59 = int32(_a_F__PG_init_pg_stat_statements_18)
								v60 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[3]))
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[4])) = v60
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[3])) = int32(_a_F__PG_init_pg_stat_statements_19)
								v65 = int32(_a_F__PG_init_pg_stat_statements_20)
								v66 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[5]))
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[5])) = int32(_a_F__PG_init_pg_stat_statements_21)
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[6])) = v66
								v72 = int32(_a_F__PG_init_pg_stat_statements_22)
								v73 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[7]))
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[7])) = int32(_a_F__PG_init_pg_stat_statements_23)
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[8])) = v73
								v79 = int32(_a_F__PG_init_pg_stat_statements_24)
								v80 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[9]))
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[9])) = int32(_a_F__PG_init_pg_stat_statements_25)
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[10])) = v80
								v86 = int32(_a_F__PG_init_pg_stat_statements_26)
								v87 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[11]))
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[11])) = int32(_a_F__PG_init_pg_stat_statements_27)
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[12])) = v87
								v93 = int32(_a_F__PG_init_pg_stat_statements_28)
								v94 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[13]))
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[13])) = int32(_a_F__PG_init_pg_stat_statements_29)
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[14])) = v94
								v100 = int32(_a_F__PG_init_pg_stat_statements_30)
								v101 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[15]))
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[15])) = int32(_a_F__PG_init_pg_stat_statements_31)
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[16])) = v101
								v107 = int32(_a_F__PG_init_pg_stat_statements_32)
								v108 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[17]))
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[17])) = int32(_a_F__PG_init_pg_stat_statements_33)
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[18])) = v108
								v114 = int32(_a_F__PG_init_pg_stat_statements_34)
								v115 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[19]))
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[19])) = int32(_a_F__PG_init_pg_stat_statements_35)
								*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[20])) = v115
								return
							}
						}
					}
				}
			}
		}
	} else {
		return
	}
}
func F__PG_init_vector(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	*(*int32)(unsafe.Add(mBase, _c_F__PG_init_vector[0])) = int32(_a_F__PG_init_vector_0)
	*(*int32)(unsafe.Add(mBase, _c_F__PG_init_vector[1])) = int32(_a_F__PG_init_vector_1)
	*(*int32)(unsafe.Add(mBase, _c_F__PG_init_vector[2])) = int32(_a_F__PG_init_vector_2)
	*(*int32)(unsafe.Add(mBase, _c_F__PG_init_vector[3])) = int32(_a_F__PG_init_vector_3)
	*(*int32)(unsafe.Add(mBase, _c_F__PG_init_vector[4])) = int32(_a_F__PG_init_vector_4)
	*(*int32)(unsafe.Add(mBase, _c_F__PG_init_vector[5])) = int32(_a_F__PG_init_vector_5)
	F_HnswInit(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		F_IvfflatInit(m)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			return
		}
	}
}
func F_create_pg_locale_builtin(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 == int32(100) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L33
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L30
	}
L3:
	;
	v35 = F_text_to_cstring(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L14
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_builtin[0]))
	v14 = F_SearchSysCache1(m, int32(21), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v25 = F_SearchSysCache1(m, int32(16), l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L11
	}
L7:
	;
	return int32(0)
L8:
	;
	if v14 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v22 = F_SysCacheGetAttrNotNull(m, int32(21), v14, int32(15))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v33 = v14
	v34 = v22
	goto L3
L11:
	;
	if v25 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v31 = F_SysCacheGetAttrNotNull(m, int32(16), v25, int32(10))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v33 = v25
	v34 = v31
	goto L3
L14:
	;
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_builtin[1]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	goto L16
L16:
	;
	v42 = F_builtin_validate_locale(m, v41, v35)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v45 = F_MemoryContextAllocZero(m, l1, int32(20))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v47 = F_MemoryContextStrdup(m, l1, v35)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v47
	v50 = int32(_a_F_create_pg_locale_builtin_0)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_builtin[2])))
	if base.B2i32(v53 == int32(0))|base.B2i32(v53 != v56) != 0 {
		v74 = v53
		v75 = v56
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)) = uint8(v77)
	v79 = int32(354)
	*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v79)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+16)) = uint8(base.B2i32(v74-v75 == int32(0)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v84 == int32(67) {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	v59 = v35
	v60 = v50
	goto L23
L23:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v64 == int32(0) {
		v74 = v64
		v75 = v63
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v74 = v64
	v75 = v63
	goto L21
L25:
	;
	v67 = int32(1)
	if v64 == v63 {
		v59 = v59 + v67
		v60 = v60 + v67
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	v89 = v87
	goto L29
L28:
	;
	v89 = int32(1)
	goto L29
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+3)) = uint8(base.B2i32(v89 == int32(0)))
	m.G0 = v7 + int32(32)
	return v45
L30:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_builtin[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v102
	F_errmsg_internal(m, int32(_a_F_create_pg_locale_builtin_1), v7)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_create_pg_locale_builtin_2), int32(135), int32(_a_F_create_pg_locale_builtin_3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_create_pg_locale_builtin_4), v7+int32(16))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_create_pg_locale_builtin_2), int32(148), int32(_a_F_create_pg_locale_builtin_3))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_analyze_and_rewrite_varparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int64
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[0])))
	if v15 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(_a_F_pg_analyze_and_rewrite_varparams_0)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[1])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[2])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[3])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[4])) = int64(1)
	v28 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L4
L2:
	;
	goto L3
L3:
	;
	v32 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_gettimeofday(m, int32(_a_F_pg_analyze_and_rewrite_varparams_1))
	mBase = m.M
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = l1
	F_setup_parse_variable_parameters(m, v32, l2, l3)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 != int32(141) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v92 = F_transformStmt(m, v32, v84)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L22
	}
L9:
	;
	v84 = v41
	goto L8
L10:
	;
	goto L11
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = v41
	goto L15
L13:
	;
	v61 = v41
	goto L14
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v66 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+76))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+68))
	if v56 != 0 {
		v50 = v55
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v61 = v55
	goto L14
L17:
	;
	goto L16
L18:
	;
	v84 = v41
	goto L8
L19:
	;
	goto L20
L20:
	;
	v70 = F_palloc0(m, int32(20))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(242)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+16)) = uint8(v76)
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = int32(0)
	v84 = v70
	goto L8
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+156)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+160)) = v96
	F_check_variable_parameters(m, v32, v92)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[5]))
	switch v101 {
	case 0:
		v108 = v5
		goto L24
	case 1:
		goto L25
	default:
		goto L26
	}
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[6]))
	if v110 != 0 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v106 = F_JumbleQuery(m, v92)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[7])))
	if v103 != int32(1) {
		v108 = v5
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v108 = v106
	goto L24
L29:
	;
	m.T0[v110].(func(*base.Module, int32, int32, int32))(m, v32, v92, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_free_parsestate(m, v32)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v92)+16))
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[8]))
	if v119 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(0) < v155 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	goto L34
L36:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[9])))
	if v123&int32(1) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v119)+392))
	if int32(1)&base.B2i32(v130 != int64(0)) != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v134 = int32(_a_F_pg_analyze_and_rewrite_varparams_2)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[10]))
	v137 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[10])) = v136 + v137
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v140 + v137
	*(*int64)(unsafe.Add(mBase, uint32(v119)+392)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v140 + int32(2)
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[10])) = v151 - v137
	goto L35
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L52
	}
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v163 = int32(0)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[0])))
	if v191 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v158+v163<<(uint(int32(2))%32))))
	if base.B2i32(v172 == int32(705))|base.B2i32(v172 == int32(0)) != 0 {
		goto L39
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	v179 = v163 + int32(1)
	if v179 != v155 {
		v163 = v179
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	F_ShowUsage(m, int32(_a_F_pg_analyze_and_rewrite_varparams_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v195 = F_pg_rewrite_query(m, v92)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	m.G0 = v12 + int32(16)
	return v195
L52:
	;
	F_errcode(m, int32(134611076))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v163 + int32(1)
	F_errmsg(m, int32(_a_F_pg_analyze_and_rewrite_varparams_4), v12)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_pg_analyze_and_rewrite_varparams_5), int32(842), int32(_a_F_pg_analyze_and_rewrite_varparams_6))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_analyze_and_rewrite_withcb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	v6 = int32(0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[0])))
	if v9 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = int32(_a_F_pg_analyze_and_rewrite_withcb_0)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[1])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[2])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[3])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[4])) = int64(1)
	v22 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L4
L2:
	;
	goto L3
L3:
	;
	v26 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_gettimeofday(m, int32(_a_F_pg_analyze_and_rewrite_withcb_1))
	mBase = m.M
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = l1
	m.T0[l2].(func(*base.Module, int32, int32))(m, v26, l3)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v35 != int32(141) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v79 = F_transformStmt(m, v26, v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L22
	}
L9:
	;
	v75 = v34
	goto L8
L10:
	;
	goto L11
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = v34
	goto L15
L13:
	;
	v52 = v34
	goto L14
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v55 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	if v47 != 0 {
		v43 = v46
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v52 = v46
	goto L14
L17:
	;
	goto L16
L18:
	;
	v75 = v34
	goto L8
L19:
	;
	goto L20
L20:
	;
	v59 = F_palloc0(m, int32(20))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(242)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v65 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+16)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = int32(0)
	v75 = v59
	goto L8
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+156)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+160)) = v83
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[5]))
	switch v86 {
	case 0:
		v93 = v6
		goto L23
	case 1:
		goto L24
	default:
		goto L25
	}
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[6]))
	if v95 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v91 = F_JumbleQuery(m, v79)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[10])))
	if v88 != int32(1) {
		v93 = v6
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v93 = v91
	goto L23
L28:
	;
	m.T0[v95].(func(*base.Module, int32, int32, int32))(m, v26, v79, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_free_parsestate(m, v26)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[7]))
	if v104 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[0])))
	if v141 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	goto L33
L35:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[8])))
	if v108&int32(1) == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v104)+392))
	if int32(1)&base.B2i32(v115 != int64(0)) != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v119 = int32(_a_F_pg_analyze_and_rewrite_withcb_2)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[9]))
	v122 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[9])) = v121 + v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v125 + v122
	*(*int64)(unsafe.Add(mBase, uint32(v104)+392)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v125 + int32(2)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[9])) = v136 - v122
	goto L34
L38:
	;
	F_ShowUsage(m, int32(_a_F_pg_analyze_and_rewrite_withcb_3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v147 = F_pg_rewrite_query(m, v79)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	return v147
}
func F_pg_ascii2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v8
	return v8
L2:
	;
	goto L3
L3:
	;
	v12 = l0
	v13 = l1
	v15 = v4
	goto L5
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
	return v30
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v17 == int32(0) {
		v29 = v13
		v30 = v15
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v29 = v22
	v30 = l2
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v17
	v22 = v13 + int32(4)
	v23 = int32(1)
	v26 = v15 + v23
	if v26 != l2 {
		v12 = v12 + v23
		v13 = v22
		v15 = v26
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_pg_available_extension_versions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int64
	_ = v546
	var v548 int64
	_ = v548
	var v550 int64
	_ = v550
	var v552 int64
	_ = v552
	var v554 int64
	_ = v554
	var v556 int64
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(48)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = F_get_extension_control_directories(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v756 + int32(48)
	return int32(0)
L4:
	;
	if v33 == int32(0) {
		v756 = v25
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v37 <= int32(0) {
		v756 = v25
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v41 = v25
	v48 = v33
	v54 = v27
	v55 = v2
	v58 = v2
	goto L7
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v55<<(uint(int32(2))%32))))
	v67 = F_AllocateDir(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v756 = v730
	goto L3
L9:
	;
	v752 = v55 + int32(1)
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v737)+4))
	if v752 < v753 {
		v41 = v730
		v48 = v737
		v54 = v743
		v55 = v752
		v58 = v747
		goto L7
	} else {
		goto L138
	}
L10:
	;
	if v67 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_available_extension_versions[0]))
	if v72 == int32(44) {
		v730 = v41
		v737 = v48
		v743 = v54
		v747 = v58
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v93 = v58
	goto L15
L14:
	;
	goto L13
L15:
	;
	v97 = F_ReadDir(m, v67, v66)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	F_FreeDir(m, v67)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L137
	}
L17:
	;
	if v97 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v100 = v97 + int32(19)
	v104 = F_strlen(m, v100)
	mBase = m.M
	v111 = v104 + int32(1)
	goto L23
L19:
	;
	goto L20
L20:
	;
	goto L16
L21:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	v113 = int32(0)
	if v111 == v113 {
		v123 = v113
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v123 = v118
	goto L22
L25:
	;
	v117 = v111 - int32(1)
	v118 = v100 + v117
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v119 != int32(46) {
		v111 = v117
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v126 = int32(_a_F_pg_available_extension_versions_0)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_available_extension_versions[1])))
	if base.B2i32(v129 == int32(0))|base.B2i32(v129 != v132) != 0 {
		v150 = v129
		v151 = v132
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v150-v151 != 0 {
		goto L15
	} else {
		goto L35
	}
L29:
	;
	goto L28
L30:
	;
	v135 = v123
	v136 = v126
	goto L31
L31:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v140 == int32(0) {
		v150 = v140
		v151 = v139
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v150 = v140
	v151 = v139
	goto L29
L33:
	;
	v143 = int32(1)
	if v140 == v139 {
		v135 = v135 + v143
		v136 = v136 + v143
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v153 = F_pstrdup(m, v100)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v158 = F_strlen(m, v153)
	mBase = m.M
	v165 = v158 + int32(1)
	goto L39
L37:
	;
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v178)
	v181 = F_strstr(m, v153, int32(_a_F_pg_available_extension_versions_1))
	mBase = m.M
	if v181 != 0 {
		goto L15
	} else {
		goto L43
	}
L38:
	;
	goto L37
L39:
	;
	v167 = int32(0)
	if v165 == v167 {
		v177 = v167
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v177 = v172
	goto L38
L41:
	;
	v171 = v165 - int32(1)
	v172 = v153 + v171
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v173 != int32(46) {
		v165 = v171
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v182 = F_makeString(m, v153)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v184 = F_list_member(m, v93, v182)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v184 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	v186 = F_lappend(m, v93, v182)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v189 = F_palloc0(m, int32(48))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v191 = F_pstrdup(m, v153)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+36)) = int32(-1)
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+34)) = uint8(v195)
	v197 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v189)+32)) = uint16(v197)
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v191
	v200 = F_pstrdup(m, v66)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+8)) = v200
	F_parse_extension_control_file(m, v189, int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	v208 = F_get_ext_ver_list(m, v189)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v208 == int32(0) {
		v93 = v186
		goto L15
	} else {
		goto L53
	}
L53:
	;
	v212 = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v213 <= v212 {
		v93 = v186
		goto L15
	} else {
		goto L54
	}
L54:
	;
	v220 = v213
	v225 = v212
	goto L55
L55:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238+v225<<(uint(int32(2))%32))))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+8)))
	if v243 != int32(1) {
		v706 = v220
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v93 = v186
	goto L15
L57:
	;
	v725 = v225 + int32(1)
	if v725 < v706 {
		v220 = v706
		v225 = v725
		goto L55
	} else {
		goto L136
	}
L58:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v248 = F_palloc(m, int32(48))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v189)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+40)) = v250
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v189)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+32)) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v189)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+24)) = v254
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v189)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+16)) = v256
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v189)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+8)) = v258
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
	*(*int64)(unsafe.Add(mBase, uint32(v248))) = v260
	F_parse_extension_control_file(m, v248, v246)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v264 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+40)) = v264
	*(*int64)(unsafe.Add(mBase, uint32(v41)+32)) = v264
	*(*int64)(unsafe.Add(mBase, uint32(v41)+24)) = v264
	*(*int64)(unsafe.Add(mBase, uint32(v41)+16)) = v264
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = v264
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v277 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v277
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v281 = F_cstring_to_text(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v281
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+33)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v284
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+34)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v286
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+32)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v248)+28))
	if v290 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v248)+40))
	if v300 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v293 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+13)) = uint8(v293)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v297 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v290)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v297
	goto L63
L68:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v248)+24))
	if v399 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	v303 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+14)) = uint8(v303)
	goto L68
L70:
	;
	goto L71
L71:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v308 = F_palloc(m, v305<<(uint(int32(2))%32))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v310 = int32(0)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	if v310 < v311 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v314 = v310
	goto L76
L74:
	;
	v351 = v310
	goto L75
L75:
	;
	v374 = F_construct_array_builtin(m, v308, v351, int32(19))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L80
	}
L76:
	;
	v337 = v314 << (uint(int32(2)) % 32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v341+v337)))
	v344 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v351 = v348
	goto L75
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v308+v337))) = v344
	v348 = v314 + int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	if v348 < v349 {
		v314 = v348
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v374
	goto L68
L81:
	;
	F_tuplestore_putvalues(m, v207, v206, v41+int32(16), v41+int32(8))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L86
	}
L82:
	;
	v402 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)) = uint8(v402)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v404 = F_cstring_to_text(m, v399)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v404
	goto L81
L86:
	;
	v413 = int32(0)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v414 <= v413 {
		v706 = v414
		goto L57
	} else {
		goto L87
	}
L87:
	;
	v421 = v414
	v428 = v413
	goto L88
L88:
	;
	if v421 <= int32(0) {
		v681 = v421
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v706 = v681
	goto L57
L90:
	;
	v700 = v428 + int32(1)
	if v700 < v681 {
		v421 = v681
		v428 = v700
		goto L88
	} else {
		goto L135
	}
L91:
	;
	v441 = int32(0)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v444+v428<<(uint(int32(2))%32))))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+8)))
	if v449&int32(1) != 0 {
		v681 = v421
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v452 = v441
	v454 = v441
	v457 = v441
	goto L93
L93:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v474+v452<<(uint(int32(2))%32))))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+8)))
	if v479 != int32(1) {
		v532 = v454
		v533 = v457
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v533 != v242 {
		v681 = v539
		goto L90
	} else {
		goto L117
	}
L95:
	;
	v538 = v452 + int32(1)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v538 < v539 {
		v452 = v538
		v454 = v532
		v457 = v533
		goto L93
	} else {
		goto L116
	}
L96:
	;
	v482 = int32(1)
	v484 = F_find_update_path(m, v208, v478, v448, v482, v482)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v484 == int32(0) {
		v532 = v454
		v533 = v457
		goto L95
	} else {
		goto L98
	}
L98:
	;
	if v457 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v532 = v484
	v533 = v478
	goto L95
L100:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	if v454 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v490 != v498 {
		v532 = v454
		v533 = v457
		goto L95
	} else {
		goto L107
	}
L102:
	;
	v493 = int32(0)
	if v493 <= v490 {
		v498 = v493
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	if v490 < v496 {
		goto L99
	} else {
		goto L106
	}
L105:
	;
	goto L99
L106:
	;
	v498 = v496
	goto L101
L107:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if base.B2i32(v504 == int32(0))|base.B2i32(v504 != v507) != 0 {
		v525 = v504
		v526 = v507
		goto L109
	} else {
		goto L110
	}
L108:
	;
	if int32(0) <= v525-v526 {
		v532 = v454
		v533 = v457
		goto L95
	} else {
		goto L115
	}
L109:
	;
	goto L108
L110:
	;
	v510 = v500
	v511 = v501
	goto L111
L111:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511)+1)))
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)))
	if v515 == int32(0) {
		v525 = v515
		v526 = v514
		goto L109
	} else {
		goto L113
	}
L112:
	;
	v525 = v515
	v526 = v514
	goto L109
L113:
	;
	v518 = int32(1)
	if v515 == v514 {
		v510 = v510 + v518
		v511 = v511 + v518
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	goto L99
L116:
	;
	goto L94
L117:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v544 = F_palloc(m, int32(48))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v189)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v544)+40)) = v546
	v548 = *(*int64)(unsafe.Add(mBase, uint32(v189)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v544)+32)) = v548
	v550 = *(*int64)(unsafe.Add(mBase, uint32(v189)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v544)+24)) = v550
	v552 = *(*int64)(unsafe.Add(mBase, uint32(v189)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v544)+16)) = v552
	v554 = *(*int64)(unsafe.Add(mBase, uint32(v189)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v544)+8)) = v554
	v556 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
	*(*int64)(unsafe.Add(mBase, uint32(v544))) = v556
	F_parse_extension_control_file(m, v544, v542)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v561 = F_cstring_to_text(m, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v561
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+33)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v564
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+34)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v566
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+32)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v568
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v544)+40))
	if v570 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+14)) = uint8(v649)
	F_tuplestore_putvalues(m, v207, v206, v41+int32(16), v41+int32(8))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L134
	}
L122:
	;
	v649 = int32(1)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	v577 = F_palloc(m, v574<<(uint(int32(2))%32))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v579 = int32(0)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v579 < v581 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v584 = v579
	goto L129
L127:
	;
	v621 = v579
	goto L128
L128:
	;
	v644 = F_construct_array_builtin(m, v577, v621, int32(19))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L133
	}
L129:
	;
	v607 = v584 << (uint(int32(2)) % 32)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v611+v607)))
	v614 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	v621 = v618
	goto L128
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v577+v607))) = v614
	v618 = v584 + int32(1)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v618 < v619 {
		v584 = v618
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v644
	v649 = v579
	goto L121
L134:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v681 = v676
	goto L90
L135:
	;
	goto L89
L136:
	;
	goto L56
L137:
	;
	v730 = v41
	v737 = v48
	v743 = v54
	v747 = v93
	goto L9
L138:
	;
	goto L8
}
func F_pg_b64_dec_len(m *base.Module, l0 int32) int32 {
	return l0 * int32(3) >> (uint(int32(2)) % 32)
}
func F_pg_buffercache_evict_relation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+46)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+44)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v2
	v30 = F_get_call_result_type(m, l0, v2, v13+int32(-4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L3
	} else {
		goto L68
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L3
	} else {
		goto L64
	}
L3:
	;
	return int32(0)
L4:
	;
	if v30 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v36 = F_superuser(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L3
	} else {
		goto L61
	}
L8:
	;
	if v36 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = F_relation_open(m, v40, int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+118)))
	if v45 == int32(116) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v48 = m.G0
	v50 = v48 - int32(32)
	m.G0 = v50
	v53 = v13 + int32(-32)
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v54
	v57 = v13 + int32(-24)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v54
	v61 = v13 + int32(-28)
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v54
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[0]))
	if v54 < v65 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v74 = int32(1)
	goto L15
L13:
	;
	goto L14
L14:
	;
	m.G0 = v50 + int32(32)
	F_relation_close(m, v42, int32(1))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L3
	} else {
		goto L58
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[1]))
	v85 = v82 + v74<<(uint(int32(6))%32)
	v87 = v85 - int32(40)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[2]))
	if v90 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v88&int32(16777216) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v234 = v74 + int32(1)
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[0]))
	if v234 <= v236 {
		v74 = v234
		goto L15
	} else {
		goto L57
	}
L22:
	;
	v98 = v85 + int32(-64)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v99 != v100 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v103 = v85 - int32(60)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v104 != v105 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v108 = v85 - int32(56)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v109 != v110 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[3]))
	F_ResourceOwnerEnlarge(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = int32(_a_F_pg_buffercache_evict_relation_0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = int32(_a_F_pg_buffercache_evict_relation_1)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = int32(_a_F_pg_buffercache_evict_relation_2)
	v124 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v124
	*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = int64(0)
	v128 = int32(_a_F_pg_buffercache_evict_relation_3)
	v130 = base.AtomicRmwOr32(m, v87, v124, v128)
	if v130&v128 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	goto L31
L29:
	;
	v156 = v130
	goto L30
L30:
	;
	v169 = int32(_a_F_pg_buffercache_evict_relation_4)
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[4]))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(8))+8))
	if v172 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	F_perform_spin_delay(m, v50+int32(8))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L33
	}
L32:
	;
	v156 = v151
	goto L30
L33:
	;
	v149 = int32(_a_F_pg_buffercache_evict_relation_3)
	v151 = base.AtomicRmwOr32(m, v87, int32(0), v149)
	if v151&v149 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if v156&int32(16777216) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_relation[4])) = v187
	goto L36
L38:
	;
	if int32(999) < v170 {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v170 < int32(11) {
		goto L36
	} else {
		goto L45
	}
L41:
	;
	v177 = int32(900)
	if v177 <= v170 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v180 = v177
	goto L44
L43:
	;
	v180 = v170
	goto L44
L44:
	;
	v187 = v180 + int32(100)
	goto L37
L45:
	;
	v187 = v170 - int32(1)
	goto L37
L46:
	;
	v207 = F_EvictUnpinnedBufferInternal(m, v98, v50+int32(8))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L52
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v156 & int32(-4194305)
	goto L21
L48:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v193 != v194 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v196 != v197 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v199 == v200 {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	if v207 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v209 = v57
	goto L55
L54:
	;
	v209 = v53
	goto L55
L55:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v211 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v210 + v211
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+8)))
	if v214 != v211 {
		goto L21
	} else {
		goto L56
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v217 + int32(1)
	goto L21
L57:
	;
	goto L16
L58:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v260
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v267 = F_heap_form_tuple(m, v262, v13+int32(-16), v13+int32(-20))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v267)+16))
	v270 = F_HeapTupleHeaderGetDatum(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	m.G0 = v15 - int32(-64)
	return v270
L61:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_evict_relation_5), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_relation_6), int32(727), int32(_a_F_pg_buffercache_evict_relation_7))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L3
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
	F_errcode(m, int32(16797828))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_pg_buffercache_evict_relation_7)
	F_errmsg(m, int32(_a_F_pg_buffercache_evict_relation_8), v13+int32(-48))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_relation_6), int32(672), int32(_a_F_pg_buffercache_evict_relation_9))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_pg_buffercache_evict_relation_7)
	F_errmsg(m, int32(_a_F_pg_buffercache_evict_relation_10), v15)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_relation_6), int32(739), int32(_a_F_pg_buffercache_evict_relation_7))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_buffercache_summary(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 float64
	_ = v102
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v17 = F_get_call_result_type(m, l0, v2, v12+int32(44))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v95
	if v95 != 0 {
		goto L24
	} else {
		goto L25
	}
L2:
	;
	v42 = int32(0)
	v44 = v2
	v46 = v2
	v47 = v2
	v48 = v2
	v49 = int64(0)
	goto L12
L3:
	;
	return int32(0)
L4:
	;
	if v17 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_summary[0]))
	if int32(0) < v24 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v95 = v2
	v97 = v2
	v98 = v2
	v99 = v2
	v102 = float64(0)
	goto L1
L9:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_summary_0), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_summary_1), int32(568), int32(_a_F_pg_buffercache_summary_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_summary[1]))
	if v52 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v95 = v78
	v97 = v79
	v98 = v80
	v99 = v86
	v102 = base.F64_convert_i64_s(v81)
	goto L1
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_summary[2]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v42<<(uint(int32(6))%32))+24))
	if v60&int32(16777216) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	v86 = v48 + base.B2i32(v60&int32(_a_F_pg_buffercache_summary_3) != int32(0))
	v88 = v42 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_summary[0]))
	if v88 < v90 {
		v42 = v88
		v44 = v78
		v46 = v79
		v47 = v80
		v48 = v86
		v49 = v81
		goto L12
	} else {
		goto L22
	}
L19:
	;
	v63 = int32(1)
	v78 = v44 + v63
	v79 = int32(base.Ui32(v60)>>(uint(int32(23))%32))&v63 + v46
	v80 = v47
	v81 = v49 + base.I64_extend_i32_u(int32(base.Ui32(v60)>>(uint(int32(18))%32))&int32(15))
	goto L18
L20:
	;
	goto L21
L21:
	;
	v78 = v44
	v79 = v46
	v80 = v47 + int32(1)
	v81 = v49
	goto L18
L22:
	;
	goto L13
L23:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v123 = F_heap_form_tuple(m, v118, v12+int32(16), v12+int32(8))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L28
	}
L24:
	;
	v113 = F_Float8GetDatum(m, base.F64_div(v102, base.F64_convert_i32_s(v95)))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v116)
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v113
	goto L23
L28:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v126 = F_HeapTupleHeaderGetDatum(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	m.G0 = v12 + int32(48)
	return v126
}
func F_pg_check_visible(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13981(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pg_checksum_page(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
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
	var v326 int32
	_ = v326
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	v3 = int32(0)
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	v39 = m.G0
	v40 = int32(128)
	v41 = v39 - v40
	base.MemoryCopy(m, v41, int32(_a_F_pg_checksum_page_0), v40)
	v48 = v3
	for {
		v82 = l0 + v48<<(uint(int32(7))%32)
		v89 = int32(0)
		for {
			v119 = int32(2)
			v120 = v89 << (uint(v119) % 32)
			v121 = v41 + v120
			v123 = *(*int32)(unsafe.Add(mBase, uint32(v82+v120)))
			v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
			v125 = v123 ^ v124
			v126 = int32(16777619)
			v128 = int32(17)
			*(*int32)(unsafe.Add(mBase, uint32(v121))) = v125*v126 ^ int32(base.Ui32(v125)>>(uint(v128)%32))
			v133 = v120 | int32(4)
			v134 = v41 + v133
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v82+v133)))
			v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
			v138 = v136 ^ v137
			*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138*v126 ^ int32(base.Ui32(v138)>>(uint(v128)%32))
			v146 = v89 + v119
			if v146 != int32(32) {
				v89 = v146
				continue
			} else {
				break
			}
			break
		}
		v150 = v48 + int32(1)
		if v150 != int32(64) {
			v48 = v150
			continue
		} else {
			break
		}
		break
	}
	v159 = int32(0)
	for {
		v191 = v41 + v159<<(uint(int32(2))%32)
		v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
		v193 = int32(16777619)
		v195 = int32(17)
		*(*int32)(unsafe.Add(mBase, uint32(v191))) = v192*v193 ^ int32(base.Ui32(v192)>>(uint(v195)%32))
		v199 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v191)+4)) = v199*v193 ^ int32(base.Ui32(v199)>>(uint(v195)%32))
		v206 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v206*v193 ^ int32(base.Ui32(v206)>>(uint(v195)%32))
		v213 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v191)+12)) = v213*v193 ^ int32(base.Ui32(v213)>>(uint(v195)%32))
		v221 = v159 + int32(4)
		if v221 != int32(32) {
			v159 = v221
			continue
		} else {
			break
		}
		break
	}
	v230 = int32(0)
	for {
		v262 = v41 + v230<<(uint(int32(2))%32)
		v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
		v264 = int32(16777619)
		v266 = int32(17)
		*(*int32)(unsafe.Add(mBase, uint32(v262))) = v263*v264 ^ int32(base.Ui32(v263)>>(uint(v266)%32))
		v270 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v270*v264 ^ int32(base.Ui32(v270)>>(uint(v266)%32))
		v277 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v262)+8)) = v277*v264 ^ int32(base.Ui32(v277)>>(uint(v266)%32))
		v284 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v262)+12)) = v284*v264 ^ int32(base.Ui32(v284)>>(uint(v266)%32))
		v292 = v230 + int32(4)
		if v292 != int32(32) {
			v230 = v292
			continue
		} else {
			break
		}
		break
	}
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v41)+124))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v41)+120))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v41)+116))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v41)+112))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v41)+108))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v41)+100))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v41)+92))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v41)+84))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v41)+76))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v41)+64))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v41)+60))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v36)
	v360 = int32(_a_F_pg_checksum_page_1)
	v361 = base.I32_rem_u_s(v295^(v296^(v297^(v298^(v299^(v300^(v301^(v302^(v303^(v304^(v305^(v306^(v307^(v308^(v309^(v310^(v311^(v312^(v313^(v314^(v315^(v316^(v317^(v318^(v319^(v320^(v321^(v322^(v323^(v324^(v325^(l1^v326))))))))))))))))))))))))))))))), v360)
	return (v361 + int32(1)) & v360
}
func F_pg_client_to_server(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_client_to_server[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = F_pg_any_to_server(m, l0, l1, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_pg_collation_for(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_get_fn_expr_argtype(m, v9, v2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v36 = v2
			m.G0 = v7 + int32(16)
			return v36
		} else {
			v19 = F_type_is_collatable(m, v11)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if base.B2i32(v19 == int32(0))&base.B2i32(v11 != int32(705)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = F_format_type_be(m, v11)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v48
								F_errmsg(m, int32(_a_F_pg_collation_for_0), v7)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_collation_for_1), int32(631), int32(_a_F_pg_collation_for_2))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v26 == int32(0) {
						v29 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
						v36 = v2
						m.G0 = v7 + int32(16)
						return v36
					} else {
						v31 = F_generate_collation_name(m, v26)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = F_cstring_to_text(m, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v36 = v33
								m.G0 = v7 + int32(16)
								return v36
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_conf_load_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_pg_conf_load_time[0]))
	v4 = F_Int64GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pg_control_recovery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
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
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v11 = F_get_call_result_type(m, l0, int32(0), v6+int32(4))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_recovery[0]))
			v22 = F_LWLockAcquire(m, v18+int32(1152), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_recovery[1]))
				v28 = F_get_controlfile(m, v25, v6+int32(3))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_recovery[0]))
					F_LWLockRelease(m, v31+int32(1152))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
						if v36 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_pg_control_recovery_0), int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_control_recovery_1), int32(181), int32(_a_F_pg_control_recovery_2))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v28)+136))
							v40 = F_Int64GetDatum(m, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)) = uint8(v42)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v40
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+144))
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v42)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v45
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v28)+152))
								v50 = F_Int64GetDatum(m, v49)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v52)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v50
									v55 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
									v56 = F_Int64GetDatum(m, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v58)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v56
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+168)))
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v58)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v61
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
										v70 = F_heap_form_tuple(m, v65, v6+int32(16), v6+int32(11))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
											v73 = F_HeapTupleHeaderGetDatum(m, v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												m.G0 = v6 + int32(48)
												return v73
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_control_recovery_3), int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_control_recovery_1), int32(173), int32(_a_F_pg_control_recovery_2))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_pg_convert(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = m.G0
	v28 = v26 + int32(-64)
	m.G0 = v28
	v30 = int32(-1)
	if v18 == int32(0) {
		v105 = v30
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v119 = m.G0
	v121 = v119 + int32(-64)
	m.G0 = v121
	v123 = int32(-1)
	if v111 == int32(0) {
		v198 = v123
		goto L30
	} else {
		goto L31
	}
L4:
	;
	m.G0 = v28 - int32(-64)
	goto L3
L5:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v33 == int32(0) {
		v105 = v30
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = F_strlen(m, v18)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v36) {
		v105 = v30
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v39 = v18
	v40 = v33
	v41 = v28
	goto L8
L8:
	;
	v49 = F_isalnum(m, v40&int32(255))
	mBase = m.M
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v66)
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
	v71 = int32(_a_F_pg_convert_0)
	v72 = int32(_a_F_pg_convert_1)
	goto L17
L10:
	;
	if base.Ui32((v40-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v62 = v41
	goto L12
L12:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v63 != 0 {
		v39 = v39 + int32(1)
		v40 = v63
		v41 = v62
		goto L8
	} else {
		goto L16
	}
L13:
	;
	v58 = v40 | int32(32)
	goto L15
L14:
	;
	v58 = v40
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v58)
	v62 = v41 + int32(1)
	goto L12
L16:
	;
	goto L9
L17:
	;
	v84 = v72 + (v71-v72)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85))))
	v87 = v70 - v86
	if v87 != 0 {
		v90 = v87
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v105 = v30
	goto L4
L19:
	;
	v94 = base.B2i32(v90 < int32(0))
	if v90 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v88 = F_strcmp(m, v28, v85)
	mBase = m.M
	if v88 != 0 {
		v90 = v88
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v105 = v89
	goto L4
L22:
	;
	v95 = v84 - int32(8)
	goto L24
L23:
	;
	v95 = v71
	goto L24
L24:
	;
	if v90 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v98 = v72
	goto L27
L26:
	;
	v98 = v84 + int32(8)
	goto L27
L27:
	;
	if base.Ui32(v98) <= base.Ui32(v95) {
		v71 = v95
		v72 = v98
		goto L17
	} else {
		goto L28
	}
L28:
	;
	goto L18
L29:
	;
	if int32(0) <= v105 {
		goto L57
	} else {
		goto L58
	}
L30:
	;
	m.G0 = v121 - int32(-64)
	goto L29
L31:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v126 == int32(0) {
		v198 = v123
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v129 = F_strlen(m, v111)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v129) {
		v198 = v123
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v132 = v111
	v133 = v126
	v134 = v121
	goto L34
L34:
	;
	v142 = F_isalnum(m, v133&int32(255))
	mBase = m.M
	if v142 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v159)
	v163 = int32(*(*int8)(unsafe.Add(mBase, uint32(v121))))
	v164 = int32(_a_F_pg_convert_0)
	v165 = int32(_a_F_pg_convert_1)
	goto L43
L36:
	;
	if base.Ui32((v133-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v155 = v134
	goto L38
L38:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v156 != 0 {
		v132 = v132 + int32(1)
		v133 = v156
		v134 = v155
		goto L34
	} else {
		goto L42
	}
L39:
	;
	v151 = v133 | int32(32)
	goto L41
L40:
	;
	v151 = v133
	goto L41
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v151)
	v155 = v134 + int32(1)
	goto L38
L42:
	;
	goto L35
L43:
	;
	v177 = v165 + (v164-v165)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v179 = int32(*(*int8)(unsafe.Add(mBase, uint32(v178))))
	v180 = v163 - v179
	if v180 != 0 {
		v183 = v180
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v198 = v123
	goto L30
L45:
	;
	v187 = base.B2i32(v183 < int32(0))
	if v183 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v181 = F_strcmp(m, v121, v178)
	mBase = m.M
	if v181 != 0 {
		v183 = v181
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	v198 = v182
	goto L30
L48:
	;
	v188 = v177 - int32(8)
	goto L50
L49:
	;
	v188 = v164
	goto L50
L50:
	;
	if v183 < int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v191 = v165
	goto L53
L52:
	;
	v191 = v177 + int32(8)
	goto L53
L53:
	;
	if base.Ui32(v191) <= base.Ui32(v188) {
		v164 = v188
		v165 = v191
		goto L43
	} else {
		goto L54
	}
L54:
	;
	goto L44
L55:
	;
	F_report_invalid_encoding(m, v105, v243+v249, v237-v249)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L97
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L93
	}
L57:
	;
	if v198 < int32(0) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L89
	}
L60:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v208 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v238 = int32(1)
	if v208&v238 != 0 {
		goto L72
	} else {
		goto L73
	}
L62:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v214 == int32(18) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	v225 = int32(1)
	if v208&v225 != 0 {
		v237 = int32(base.Ui32(v208)>>(uint(v225)%32)) - v225
		goto L61
	} else {
		goto L71
	}
L65:
	;
	v217 = int32(16)
	goto L67
L66:
	;
	v217 = int32(0)
	goto L67
L67:
	;
	if base.Ui32((v214-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v224 = int32(4)
	goto L70
L69:
	;
	v224 = v217
	goto L70
L70:
	;
	v237 = v224
	goto L61
L71:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v237 = int32(base.Ui32(v231)>>(uint(int32(2))%32)) - int32(4)
	goto L61
L72:
	;
	v242 = v238
	goto L74
L73:
	;
	v242 = int32(4)
	goto L74
L74:
	;
	v243 = v14 + v242
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v105*int32(28))+uint32(_c_F_pg_convert[0])))
	v249 = m.T0[v248].(func(*base.Module, int32, int32) int32)(m, v243, v237)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v237 != v249 {
		goto L55
	} else {
		goto L76
	}
L76:
	;
	v252 = F_pg_do_encoding_conversion(m, v243, v237, v105, v198)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	m.G0 = v11 + int32(32)
	return v273
L78:
	;
	if v243 == v252 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v273 = v14
	goto L77
L80:
	;
	goto L81
L81:
	;
	v255 = F_strlen(m, v252)
	mBase = m.M
	v257 = v255 + int32(4)
	v258 = F_palloc(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v257 << (uint(int32(2)) % 32)
	if v255 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	base.MemoryCopy(m, v258+int32(4), v252, v255)
	goto L85
L84:
	;
	goto L85
L85:
	;
	F_pfree(m, v252)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 == v268 {
		v273 = v258
		goto L77
	} else {
		goto L87
	}
L87:
	;
	F_pfree(m, v14)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v273 = v258
	goto L77
L89:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v18
	F_errmsg(m, int32(_a_F_pg_convert_2), v11)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_pg_convert_3), int32(578), int32(_a_F_pg_convert_4))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v111
	F_errmsg(m, int32(_a_F_pg_convert_5), v11+int32(16))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_pg_convert_3), int32(583), int32(_a_F_pg_convert_4))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_ctype_get_cache(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ctype_get_cache[0]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ctype_get_cache[1]))
	v13 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v32 = F_emscripten_builtin_malloc(m, int32(40))
	mBase = m.M
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != l0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	if v24 != 0 {
		v13 = v24
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19 != v10 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	return v13 + int32(8)
L9:
	;
	goto L5
L10:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	F_emscripten_builtin_free(m, v303)
	mBase = m.M
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	F_emscripten_builtin_free(m, v305)
	mBase = m.M
	F_emscripten_builtin_free(m, v32)
	mBase = m.M
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = l0
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ctype_get_cache[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = int64(549755813888)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
	v39 = int32(512)
	v40 = F_emscripten_builtin_malloc(m, v39)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v32)+20)) = int64(274877906944)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v40
	v45 = F_emscripten_builtin_malloc(m, v39)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v45
	v47 = int32(0)
	if base.B2i32(v40 == v47)|base.B2i32(v45 == v47) != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v295 = int32(0)
	goto L13
L13:
	;
	return v295
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = l1
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ctype_get_cache[2]))
	switch v57 - int32(1) {
	case 0, 1:
		v64 = int32(2048)
		goto L15
	case 2:
		goto L17
	default:
		v61 = int32(128)
		goto L16
	}
L15:
	;
	v67 = v32 + int32(8)
	v71 = int32(0)
	v73 = int32(0)
	goto L19
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(-1)
	v64 = v61
	goto L15
L17:
	;
	v61 = int32(256)
	goto L16
L18:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v241 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L19:
	;
	v75 = m.T0[l0].(func(*base.Module, int32) int32)(m, v71)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v160 <= int32(0) {
		goto L18
	} else {
		goto L49
	}
L21:
	;
	v162 = v71 + int32(1)
	if v162 != v64 {
		v71 = v162
		v73 = v160
		goto L19
	} else {
		goto L48
	}
L22:
	;
	return int32(0)
L23:
	;
	if v75 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v160 = v73 + int32(1)
	goto L21
L25:
	;
	goto L26
L26:
	;
	if v73 <= int32(0) {
		v160 = v73
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v83 = v71 - v73
	if base.Ui32(int32(2)) <= base.Ui32(v73) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v153 == int32(0) {
		goto L10
	} else {
		goto L46
	}
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v88 < v89 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v124 < v125 {
		goto L40
	} else {
		goto L41
	}
L32:
	;
	v106 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v105+v104<<(uint(v106)%32)))) = v83
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v116 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v110+v111<<(uint(v106)%32))+4)) = v83 + v73 - v116
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v119 + v116
	v153 = v116
	goto L28
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v104 = v88
	v105 = v91
	goto L32
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v89 << (uint(int32(1)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v98 = F_emscripten_builtin_realloc(m, v95, v89<<(uint(int32(4))%32))
	mBase = m.M
	if v98 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v153 = int32(0)
	goto L28
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v98
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v104 = v103
	v105 = v98
	goto L32
L39:
	;
	v142 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v140 + v142
	*(*int32)(unsafe.Add(mBase, uint32(v141+v140<<(uint(int32(2))%32)))) = v83
	v153 = v142
	goto L28
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v140 = v124
	v141 = v127
	goto L39
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v125 << (uint(int32(1)) % 32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v134 = F_emscripten_builtin_realloc(m, v131, v125<<(uint(int32(3))%32))
	mBase = m.M
	if v134 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v153 = int32(0)
	goto L28
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v134
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v140 = v139
	v141 = v134
	goto L39
L46:
	;
	v158 = v71 + int32(1)
	if v158 != v64 {
		v71 = v158
		v73 = int32(0)
		goto L19
	} else {
		goto L47
	}
L47:
	;
	goto L18
L48:
	;
	goto L20
L49:
	;
	v166 = v64 - v160
	if base.Ui32(int32(2)) <= base.Ui32(v160) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v236 == int32(0) {
		goto L10
	} else {
		goto L68
	}
L51:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v171 < v172 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v207 < v208 {
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v189 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v188+v187<<(uint(v189)%32)))) = v166
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v199 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v189)%32))+4)) = v166 + v160 - v199
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v202 + v199
	v236 = v199
	goto L50
L55:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v187 = v171
	v188 = v174
	goto L54
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v172 << (uint(int32(1)) % 32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v181 = F_emscripten_builtin_realloc(m, v178, v172<<(uint(int32(4))%32))
	mBase = m.M
	if v181 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v236 = int32(0)
	goto L50
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v181
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v187 = v186
	v188 = v181
	goto L54
L61:
	;
	v225 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v223 + v225
	*(*int32)(unsafe.Add(mBase, uint32(v224+v223<<(uint(int32(2))%32)))) = v166
	v236 = v225
	goto L50
L62:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v223 = v207
	v224 = v210
	goto L61
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v208 << (uint(int32(1)) % 32)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v217 = F_emscripten_builtin_realloc(m, v214, v208<<(uint(int32(3))%32))
	mBase = m.M
	if v217 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v236 = int32(0)
	goto L50
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v217
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v223 = v222
	v224 = v217
	goto L61
L68:
	;
	goto L18
L69:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v262 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v256
	goto L69
L71:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	F_emscripten_builtin_free(m, v244)
	mBase = m.M
	v246 = int32(0)
	v256 = v246
	v257 = v246
	goto L70
L72:
	;
	goto L73
L73:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v248 <= v241 {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v253 = F_emscripten_builtin_realloc(m, v250, v241<<(uint(int32(2))%32))
	mBase = m.M
	if v253 == int32(0) {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	v256 = v253
	v257 = v241
	goto L70
L76:
	;
	v283 = int32(_a_F_pg_ctype_get_cache_0)
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ctype_get_cache[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v284
	*(*int32)(unsafe.Add(mBase, _c_F_pg_ctype_get_cache[0])) = v32
	v295 = v67
	goto L13
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v277
	goto L76
L78:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	F_emscripten_builtin_free(m, v265)
	mBase = m.M
	v267 = int32(0)
	v277 = v267
	v278 = v267
	goto L77
L79:
	;
	goto L80
L80:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v269 <= v262 {
		goto L76
	} else {
		goto L81
	}
L81:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v274 = F_emscripten_builtin_realloc(m, v271, v262<<(uint(int32(3))%32))
	mBase = m.M
	if v274 == int32(0) {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	v277 = v274
	v278 = v262
	goto L77
}
func F_pg_database_collation_actual_version(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_SearchSysCache1(m, int32(21), v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v18 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15+v16)+76)))
			if v18 == int32(99) {
				v24 = int32(13)
			} else {
				v24 = int32(15)
			}
			v25 = F_SysCacheGetAttrNotNull(m, int32(21), v11, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_text_to_cstring(m, v25)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = F_get_collation_actual_version(m, v18, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v11)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v33 = F_cstring_to_text(m, v29)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									v38 = v33
									m.G0 = v7 + int32(16)
									return v38
								}
							} else {
								v35 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
								v38 = int32(0)
								m.G0 = v7 + int32(16)
								return v38
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
					F_errmsg(m, int32(_a_F_pg_database_collation_actual_version_0), v7)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_database_collation_actual_version_1), int32(2789), int32(_a_F_pg_database_collation_actual_version_2))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_database_size_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_get_database_oid(m, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_calculate_database_size(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v9 == int64(0) {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			} else {
				v17 = F_Int64GetDatum(m, v9)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_pg_ddl_command_send(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_pg_ddl_command_send_0), int32(359), int32(_a_F_pg_ddl_command_send_1), int32(_a_F_pg_ddl_command_send_2), int32(_a_F_pg_ddl_command_send_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_detoast_datum_packed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v3 != int32(1))&base.B2i32(v3&int32(3) != int32(2)) != 0 {
		v15 = l0
		return v15
	} else {
		v11 = F_detoast_attr(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			return v15
		}
	}
}
func F_pg_encrypt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_pfree(m, v138)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L93
	}
L2:
	;
	return int32(0)
L3:
	;
	v19 = int32(1)
	v20 = v15 + v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v25 = v23 & v19
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v20
	goto L6
L5:
	;
	v26 = v15 + int32(4)
	goto L6
L6:
	;
	if v23 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v55 = F_downcase_truncate_identifier(m, v26, v53, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v32 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v35 = int32(16)
	goto L13
L12:
	;
	v35 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = int32(4)
	goto L16
L15:
	;
	v42 = v35
	goto L16
L16:
	;
	v53 = v42
	goto L7
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v59 = F_px_find_combo(m, v55, v12+int32(28))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v59 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v55)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L75
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v67 = F_pg_detoast_datum_packed(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v70 = F_pg_detoast_datum_packed(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v72 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v102 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v78 == int32(18) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v89 = int32(1)
	if v72&v89 != 0 {
		v101 = int32(base.Ui32(v72)>>(uint(v89)%32)) - v89
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v81 = int32(16)
	goto L32
L31:
	;
	v81 = int32(0)
	goto L32
L32:
	;
	if base.Ui32((v78-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v88 = int32(4)
	goto L35
L34:
	;
	v88 = v81
	goto L35
L35:
	;
	v101 = v88
	goto L26
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v101 = int32(base.Ui32(v95)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v133 = m.T0[v132].(func(*base.Module, int32, int32) int32)(m, v65, v101)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L48
	}
L38:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v108 == int32(18) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v119 = int32(1)
	if v102&v119 != 0 {
		v131 = int32(base.Ui32(v102)>>(uint(v119)%32)) - v119
		goto L37
	} else {
		goto L47
	}
L41:
	;
	v111 = int32(16)
	goto L43
L42:
	;
	v111 = int32(0)
	goto L43
L43:
	;
	if base.Ui32((v108-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v118 = int32(4)
	goto L46
L45:
	;
	v118 = v111
	goto L46
L46:
	;
	v131 = v118
	goto L37
L47:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v131 = int32(base.Ui32(v125)>>(uint(int32(2))%32)) - int32(4)
	goto L37
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v133
	v138 = F_palloc(m, v133+int32(4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v140 = int32(1)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v142&v140 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v145 = v140
	goto L52
L51:
	;
	v145 = int32(4)
	goto L52
L52:
	;
	v147 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v150 = m.T0[v149].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v65, v70+v145, v131, v147, v147)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	if v150 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v154 = int32(1)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v156&v154 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v168 = v150
	goto L56
L56:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	m.T0[v169].(func(*base.Module, int32))(m, v65)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L61
	}
L57:
	;
	v159 = v154
	goto L59
L58:
	;
	v159 = int32(4)
	goto L59
L59:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v166 = m.T0[v165].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v65, v67+v159, v101, v138+int32(4), v12+int32(28))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v168 = v166
	goto L56
L61:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v172 != v67 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_pfree(m, v67)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v176 != v70 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	F_pfree(m, v70)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v180 != v15 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	F_pfree(m, v15)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v168 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v184<<(uint(int32(2))%32) + int32(16)
	m.G0 = v12 + int32(32)
	return v138
L75:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	if v59 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
	F_errmsg(m, int32(_a_F_pg_encrypt_0), v12+int32(16))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L2
	} else {
		goto L91
	}
L78:
	;
	v230 = int32(_a_F_pg_encrypt_1)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v209 = int32(_a_F_pg_encrypt_2)
	goto L82
L81:
	;
	v230 = v224
	goto L77
L82:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+8))
	if v59 != v212 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v224 = v222
	goto L81
L84:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
	if v214 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	v230 = int32(_a_F_pg_encrypt_3)
	goto L77
L88:
	;
	goto L89
L89:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v209)+16))
	if v59 != v218 {
		v209 = v209 + int32(16)
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v224 = v214
	goto L81
L91:
	;
	F_errfinish(m, int32(_a_F_pg_encrypt_4), int32(513), int32(_a_F_pg_encrypt_5))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	if v168 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v281
	F_errmsg(m, int32(_a_F_pg_encrypt_6), v12)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L110
	}
L97:
	;
	v281 = int32(_a_F_pg_encrypt_1)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v260 = int32(_a_F_pg_encrypt_2)
	goto L101
L100:
	;
	v281 = v275
	goto L96
L101:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	if v168 != v263 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v275 = v273
	goto L100
L103:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	if v265 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	v281 = int32(_a_F_pg_encrypt_3)
	goto L96
L107:
	;
	goto L108
L108:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	if v168 != v269 {
		v260 = v260 + int32(16)
		goto L101
	} else {
		goto L109
	}
L109:
	;
	v275 = v265
	goto L100
L110:
	;
	F_errfinish(m, int32(_a_F_pg_encrypt_4), int32(290), int32(_a_F_pg_encrypt_7))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_euccn_dsplen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if v2 < int32(0) {
		return int32(2)
	} else {
		v7 = int32(-1)
		if v2 == int32(127) {
			v12 = v7
		} else {
			v12 = int32(1)
		}
		if base.Ui32(v2) < base.Ui32(int32(32)) {
			v15 = v7
		} else {
			v15 = v12
		}
		if v2 != 0 {
			v17 = v15
		} else {
			v17 = int32(0)
		}
		return v17
	}
}
func F_pg_extension_config_dump(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	v12 = m.G0
	v14 = v12 - int32(176)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_extension_config_dump[0])))
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L85
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L82
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L79
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L76
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L73
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L69
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L65
	}
L10:
	;
	v24 = F_get_rel_name(m, v16)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L61
	}
L13:
	;
	if v24 == int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v29 = F_getExtensionOfObject(m, int32(1259), v16)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_extension_config_dump[1]))
	if v29 != v32 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v36 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v39 = v14 + int32(128)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_pg_extension_config_dump[1]))
	F_ScanKeyInit(m, v39, int32(1), int32(3), int32(184), v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v48 = int32(1)
	v51 = F_systable_beginscan(m, v36, int32(3080), v48, int32(0), v48, v39)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v53 = F_systable_getnext(m, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v53 == int32(0) {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v57 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = int64(281474976710656)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v16
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v74 = F_heap_getattr_7(m, v53, int32(7), v71, v14+int32(119))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+119)))
	if v76 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v159
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v166 = F_heap_getattr_7(m, v53, int32(8), v163, v14+int32(119))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L43
	}
L24:
	;
	v79 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v79
	v85 = F_construct_array_builtin(m, v14+int32(124), v79, int32(26))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v87 = F_pg_detoast_datum(m, v74)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v153 = int32(0)
	v159 = v85
	goto L23
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v89 != int32(1) {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	if v92 != int32(1) {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	if v95 < int32(0) {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	if v98 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	if v99 != int32(26) {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v95 + int32(1)
	if v95 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	v146 = F_array_set(m, v87, v14+int32(120), v143, int32(4), int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L42
	}
L35:
	;
	v110 = int32(0)
	goto L36
L36:
	;
	v122 = v110 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v87+int32(24)+v110<<(uint(int32(2))%32))))
	if v16 == v126 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L34
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v122
	goto L34
L39:
	;
	goto L40
L40:
	;
	if v122 != v95 {
		v110 = v122
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	v153 = v95
	v159 = v146
	goto L23
L43:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+119)))
	if v168 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+71)) = uint8(v200)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v199
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v210 = F_heap_modify_tuple(m, v53, v203, v14+int32(80), v14+int32(72), v14-int32(-64))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L57
	}
L45:
	;
	if v153 != 0 {
		goto L5
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v177 = F_pg_detoast_datum(m, v166)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v175 = F_construct_array_builtin(m, v14+int32(124), int32(1), int32(25))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v199 = v175
	goto L44
L50:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v179 != int32(1) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	if v182 != int32(1) {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	if v185 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	if v186 != int32(25) {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	if v189 != v153 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	v196 = F_array_set(m, v177, v14+int32(120), v193, int32(-1), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v199 = v196
	goto L44
L57:
	;
	F_CatalogTupleUpdate(m, v36, v210+int32(4), v210)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_systable_endscan(m, v51)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_relation_close(m, v36, int32(3))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	m.G0 = v14 + int32(176)
	return int32(0)
L61:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(_a_F_pg_extension_config_dump_0)
	F_errmsg(m, int32(_a_F_pg_extension_config_dump_1), v14+int32(48))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2819), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	F_errmsg(m, int32(_a_F_pg_extension_config_dump_4), v14)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2830), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v24
	F_errmsg(m, int32(_a_F_pg_extension_config_dump_5), v14+int32(32))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2836), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_pg_extension_config_dump[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v284
	F_errmsg_internal(m, int32(_a_F_pg_extension_config_dump_6), v14+int32(16))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2861), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_pg_extension_config_dump_7), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2894), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errmsg_internal(m, int32(_a_F_pg_extension_config_dump_8), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2927), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_pg_extension_config_dump_9), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2939), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_pg_extension_config_dump_8), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2941), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_file_exists(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v13 = F___fstatat(m, int32(-100), l0, v7+int32(16), v2)
	mBase = m.M
	if v13 == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
		v44 = base.B2i32(v16&int32(_a_F_pg_file_exists_0) != int32(_a_F_pg_file_exists_1))
		m.G0 = v7 + int32(112)
		return v44
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_file_exists[0]))
		switch v22 - int32(44) {
		case 0, 10:
			v44 = v2
			m.G0 = v7 + int32(112)
			return v44
		case 1, 2, 3, 4, 5, 6, 7, 8, 9:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg(m, int32(_a_F_pg_file_exists_2), v7)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_file_exists_3), int32(514), int32(_a_F_pg_file_exists_4))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		default:
			if v22 == int32(2) {
				v44 = v2
				m.G0 = v7 + int32(112)
				return v44
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(_a_F_pg_file_exists_2), v7)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_file_exists_3), int32(514), int32(_a_F_pg_file_exists_4))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_gen_salt_rounds(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v6 = m.G0
	v8 = v6 - int32(160)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = v8 + int32(16)
	F_text_to_cstring_buffer(m, v11, v17, int32(129))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = F_px_gen_salt(m, v17, v17, v15)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if int32(0) <= v21 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v25 != v11 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	F_pfree(m, v11)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v31 = F_cstring_to_text_with_len(m, v8+int32(16), v21)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	m.G0 = v8 + int32(160)
	return v31
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v21 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v73
	F_errmsg(m, int32(_a_F_pg_gen_salt_rounds_0), v8)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L29
	}
L16:
	;
	v73 = int32(_a_F_pg_gen_salt_rounds_1)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v52 = int32(_a_F_pg_gen_salt_rounds_2)
	goto L20
L19:
	;
	v73 = v67
	goto L15
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v21 != v55 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v67 = v65
	goto L19
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	if v57 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v73 = int32(_a_F_pg_gen_salt_rounds_3)
	goto L15
L26:
	;
	goto L27
L27:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v21 != v61 {
		v52 = v52 + int32(16)
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v67 = v57
	goto L19
L29:
	;
	F_errfinish(m, int32(_a_F_pg_gen_salt_rounds_4), int32(202), int32(_a_F_pg_gen_salt_rounds_5))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_constraintdef(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_get_constraintdef_worker(m, v3, int32(0), int32(2), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			v17 = F_cstring_to_text(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_pg_get_indexdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int64
	_ = v324
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int64
	_ = v575
	var v582 int32
	_ = v582
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	v10 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(288)
	m.G0 = v30
	v33 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L6
	} else {
		goto L200
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L6
	} else {
		goto L197
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L6
	} else {
		goto L194
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L6
	} else {
		goto L191
	}
L5:
	;
	m.G0 = v30 + int32(288)
	return v640
L6:
	;
	return int32(0)
L7:
	;
	if v33 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l8 != 0 {
		v640 = int32(0)
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	v55 = v53 + v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v59 = F_SysCacheGetAttrNotNull(m, int32(34), v33, int32(17))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_indexdef_worker_0), v30)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(1308), int32(_a_F_pg_get_indexdef_worker_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
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
	v63 = F_SysCacheGetAttrNotNull(m, int32(34), v33, int32(18))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v67 = F_SysCacheGetAttrNotNull(m, int32(34), v33, int32(19))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v70 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	if v70 == int32(0) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
	v77 = v75 + v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+84))
	v79 = F_SearchSysCache1(m, int32(2), v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v79 == int32(0) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+22)))
	v85 = v83 + v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+68))
	v87 = F_GetIndexAmRoutine(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v89 = int32(0)
	v92 = F_heap_attisnull(m, v33, int32(20), v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	v112 = F_get_rel_name(m, v56)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L31
	}
L24:
	;
	if v92 != 0 {
		v110 = v10
		v111 = v89
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v96 = F_SysCacheGetAttrNotNull(m, int32(34), v33, int32(20))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v98 = F_text_to_cstring(m, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v100 = F_stringToNode(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	F_pfree(m, v98)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v104 = int32(0)
	if v100 == v104 {
		v110 = v10
		v111 = v104
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v110 = v100
	v111 = v107
	goto L23
L31:
	;
	if v112 == int32(0) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v117 = F_palloc0(m, int32(80))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v120 = F_palloc0(m, int32(136))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+24)) = int32(1)
	v124 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+21)) = uint8(v124)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v56
	v127 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+12)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(101)
	v132 = F_makeAlias(m, v112, v127)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v132
	v136 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+124)) = uint16(v136)
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+20)) = uint8(v138)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+204)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v30)+232)) = v120
	v145 = F_list_make1_impl(m, int32(1), v30+int32(204))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v147 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+20)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v117)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v145
	F_set_rtable_names(m, v117, v147, v147)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_set_simple_column_names(m, v117)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+200)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v117
	v163 = F_list_make1_impl(m, int32(1), v30+int32(200))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_initStringInfo(m, v30+int32(216))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	if l3 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v227 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+8)))
	if v227 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L42:
	;
	if l2 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+12)))
	if v173 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v213 = F_quote_identifier(m, v85+int32(4))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L6
	} else {
		goto L64
	}
L46:
	;
	v174 = int32(_a_F_pg_get_indexdef_worker_3)
	goto L48
L47:
	;
	v174 = int32(_a_F_pg_get_indexdef_worker_4)
	goto L48
L48:
	;
	v177 = F_quote_identifier(m, v77+int32(4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	v179 = int32(_a_F_pg_get_indexdef_worker_4)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+119)))
	if v182 != int32(73) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v185 = v179
	goto L52
L51:
	;
	v185 = int32(_a_F_pg_get_indexdef_worker_5)
	goto L52
L52:
	;
	if l6 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v186 = v179
	goto L55
L54:
	;
	v186 = v185
	goto L55
L55:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l7) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v197 = F_quote_identifier(m, v85+int32(4))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L62
	}
L57:
	;
	v190 = F_generate_relation_name(m, v56, int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v192 = F_generate_qualified_relation_name(m, v56)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L61
	}
L60:
	;
	v194 = v190
	goto L56
L61:
	;
	v194 = v192
	goto L56
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+176)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v30)+172)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v30)+168)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v30)+164)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v30)+160)) = v174
	F_appendStringInfo(m, v30+int32(216), int32(_a_F_pg_get_indexdef_worker_6), v30+int32(160))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	goto L41
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+192)) = v213
	F_appendStringInfo(m, v30+int32(216), int32(_a_F_pg_get_indexdef_worker_7), v30+int32(192))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	goto L41
L66:
	;
	if l3 != 0 {
		goto L152
	} else {
		goto L153
	}
L67:
	;
	v230 = int32(24)
	v246 = int32(0)
	v250 = int32(_a_F_pg_get_indexdef_worker_4)
	v256 = v111
	goto L68
L68:
	;
	v268 = v246 << (uint(int32(1)) % 32)
	v270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55+int32(48)+v268))))
	if l4 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L66
L70:
	;
	v271 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+10)))
	if v271 <= v246 {
		goto L66
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if l1 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	v276 = v30 + int32(216)
	v277 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+10)))
	if v277 == v246 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v270 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	F_appendStringInfoString(m, v276, int32(_a_F_pg_get_indexdef_worker_8))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	v283 = v250
	goto L79
L79:
	;
	F_appendStringInfoString(m, v276, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L81
	}
L80:
	;
	v283 = int32(_a_F_pg_get_indexdef_worker_4)
	goto L79
L81:
	;
	goto L76
L82:
	;
	if l3 != 0 {
		goto L115
	} else {
		goto L116
	}
L83:
	;
	v288 = F_get_attname(m, v56, v270, int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v256 == int32(0) {
		goto L1
	} else {
		goto L96
	}
L86:
	;
	if l1 != v246+int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v294 = l1
	goto L89
L88:
	;
	v294 = int32(0)
	goto L89
L89:
	;
	if v294 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v299 = F_quote_identifier(m, v288)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	F_get_atttypetypmodcoll(m, v56, v270, v30+int32(212), v30+int32(232), v30+int32(208))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L6
	} else {
		goto L95
	}
L93:
	;
	F_appendStringInfoString(m, v30+int32(216), v299)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v383 = v256
	goto L82
L96:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v317 = v30 + int32(272)
	F_initStringInfo(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	v320 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+264)) = uint8(v320)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+248)) = v320
	v324 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+240)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v30)+236)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v320
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+267)) = uint8(v320)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+256)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v30)+232)) = v317
	v335 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+265)) = uint16(v335)
	F_get_rule_expr(m, v315, v30+int32(232), v320)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v343 = v256 + int32(4)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v30)+272))
	if l1 != v246+int32(1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if base.Ui32(v343) < base.Ui32(v313+v314<<(uint(int32(2))%32)) {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v353 = l1
	goto L102
L101:
	;
	v353 = int32(0)
	goto L102
L102:
	;
	if v353 != 0 {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	if v315 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+144)) = v348
	F_appendStringInfo(m, v30+int32(216), int32(_a_F_pg_get_indexdef_worker_9), v30+int32(144))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L6
	} else {
		goto L109
	}
L105:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	switch v356 - int32(15) {
	case 0:
		goto L107
	default:
		goto L104
	case 4, 23, 24, 25, 26, 33:
		goto L106
	}
L106:
	;
	F_appendStringInfoString(m, v30+int32(216), v348)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L6
	} else {
		goto L108
	}
L107:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v315)+16))
	switch v359 {
	case 0, 3:
		goto L106
	default:
		goto L104
	}
L108:
	;
	goto L99
L109:
	;
	goto L99
L110:
	;
	v373 = v343
	goto L112
L111:
	;
	v373 = int32(0)
	goto L112
L112:
	;
	v374 = F_exprType(m, v315)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+212)) = v374
	v377 = F_exprCollation(m, v315)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+208)) = v377
	v383 = v373
	goto L82
L115:
	;
	v476 = v246 + int32(1)
	v477 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+8)))
	if v476 < v477 {
		v246 = v476
		v250 = int32(_a_F_pg_get_indexdef_worker_10)
		v256 = v383
		goto L68
	} else {
		goto L151
	}
L116:
	;
	v385 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+10)))
	if v385 <= v246 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v389 = v246 + int32(1)
	if v389 != l1 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v391 = l1
	goto L120
L119:
	;
	v391 = int32(0)
	goto L120
L120:
	;
	if v391 != 0 {
		goto L115
	} else {
		goto L121
	}
L121:
	;
	v393 = v246 << (uint(int32(2)) % 32)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v59+v230+v393)))
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v268+(v67+v230)))))
	v399 = F_get_attoptions(m, l0, base.I32_extend16_s(v389))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	if v395 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v393+(v63+v230))))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v30)+212))
	if v399 != 0 {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v30)+208))
	if v395 == v403 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v405 = F_generate_collation_name(m, v395)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+128)) = v405
	F_appendStringInfo(m, v30+int32(216), int32(_a_F_pg_get_indexdef_worker_11), v30+int32(128))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	goto L123
L128:
	;
	v419 = int32(0)
	goto L130
L129:
	;
	v419 = v418
	goto L130
L130:
	;
	v421 = v30 + int32(216)
	F_get_opclass_name(m, v416, v419, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	if v399 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	F_appendStringInfoString(m, v421, int32(_a_F_pg_get_indexdef_worker_12))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L6
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)))
	if v432 != int32(1) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	F_get_reloptions(m, v421, v399)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	F_appendStringInfoChar(m, v421, int32(41))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	if l2 == int32(0) {
		goto L115
	} else {
		goto L148
	}
L139:
	;
	v436 = v30 + int32(216)
	if v397&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	F_appendStringInfoString(m, v436, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L6
	} else {
		goto L147
	}
L141:
	;
	F_appendStringInfoString(m, v436, int32(_a_F_pg_get_indexdef_worker_13))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L6
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if v397&int32(2) == int32(0) {
		goto L138
	} else {
		goto L146
	}
L144:
	;
	if v397&int32(2) != 0 {
		goto L138
	} else {
		goto L145
	}
L145:
	;
	v450 = int32(_a_F_pg_get_indexdef_worker_14)
	goto L140
L146:
	;
	v450 = int32(_a_F_pg_get_indexdef_worker_15)
	goto L140
L147:
	;
	goto L138
L148:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l2+v393)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v30)+212))
	v459 = F_generate_operator_name(m, v457, v458, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v459
	F_appendStringInfo(m, v30+int32(216), int32(_a_F_pg_get_indexdef_worker_16), v30+int32(112))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	goto L115
L151:
	;
	goto L69
L152:
	;
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L188
	}
L153:
	;
	v507 = v30 + int32(216)
	F_appendStringInfoChar(m, v507, int32(41))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L6
	} else {
		goto L154
	}
L154:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+13)))
	if v511 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_appendStringInfoString(m, v507, int32(_a_F_pg_get_indexdef_worker_17))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L6
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v517 = F_flatten_reloptions(m, l0)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L6
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	if v517 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v517
	F_appendStringInfo(m, v30+int32(216), int32(_a_F_pg_get_indexdef_worker_18), v30+int32(96))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L6
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if l5 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	F_pfree(m, v517)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v555 = F_heap_attisnull(m, v33, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L6
	} else {
		goto L176
	}
L166:
	;
	v531 = F_get_rel_tablespace(m, l0)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	if v531 == int32(0) {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	if l2 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	F_appendStringInfoString(m, v30+int32(216), int32(_a_F_pg_get_indexdef_worker_19))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L6
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v540 = F_get_tablespace_name(m, v531)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L6
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v542 = F_quote_identifier(m, v540)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v542
	F_appendStringInfo(m, v30+int32(216), int32(_a_F_pg_get_indexdef_worker_20), v30+int32(80))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	goto L165
L176:
	;
	if v555 != 0 {
		goto L152
	} else {
		goto L177
	}
L177:
	;
	v559 = F_SysCacheGetAttrNotNull(m, int32(34), v33, int32(21))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	v561 = F_text_to_cstring(m, v559)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	v563 = F_stringToNode(m, v561)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	F_pfree(m, v561)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	v568 = v30 + int32(272)
	F_initStringInfo(m, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L6
	} else {
		goto L182
	}
L182:
	;
	v571 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+264)) = uint8(v571)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+248)) = v571
	v575 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+240)) = v575
	*(*int32)(unsafe.Add(mBase, uint32(v30)+236)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v571
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+267)) = uint8(v571)
	v582 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+265)) = uint16(v582)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+256)) = v575
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v30)+232)) = v568
	F_get_rule_expr(m, v563, v30+int32(232), v571)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L6
	} else {
		goto L183
	}
L183:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v30)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v593
	if l2 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v599 = int32(_a_F_pg_get_indexdef_worker_21)
	goto L186
L185:
	;
	v599 = int32(_a_F_pg_get_indexdef_worker_22)
	goto L186
L186:
	;
	F_appendStringInfo(m, v30+int32(216), v599, v30-int32(-64))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	goto L152
L188:
	;
	F_ReleaseCatCache(m, v70)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	F_ReleaseCatCache(m, v79)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v30)+216))
	v640 = v612
	goto L5
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_indexdef_worker_23), v30+int32(16))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(1333), int32(_a_F_pg_get_indexdef_worker_2))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L6
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v77)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v664
	F_errmsg_internal(m, int32(_a_F_pg_get_indexdef_worker_24), v30+int32(32))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L6
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(1342), int32(_a_F_pg_get_indexdef_worker_2))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v56
	F_errmsg_internal(m, int32(_a_F_pg_get_indexdef_worker_23), v30+int32(48))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(_a_F_pg_get_indexdef_worker_25), int32(_a_F_pg_get_indexdef_worker_26))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_indexdef_worker_27), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(1440), int32(_a_F_pg_get_indexdef_worker_2))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L6
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_keywords(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == v2 {
		v17 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(_a_F_pg_get_keywords_0)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[0]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[0])) = v24
			v27 = F_get_call_result_type(m, l0, int32(0), v11)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_get_keywords_1), int32(0))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_get_keywords_2), int32(431), int32(_a_F_pg_get_keywords_3))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v31
					v33 = F_TupleDescGetAttInMetadata(m, v31)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v33
						*(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[0])) = v22
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
						v45 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[1])))
						if base.Ui64(v43) < base.Ui64(v45) {
							v49 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[2]))
							v51 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[3]))
							v52 = base.I32_wrap_i64(v43)
							v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(1))%32)))))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v49 + v56
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_pg_get_keywords[4]))))
							if base.Ui32(v59) <= base.Ui32(int32(3)) {
								v63 = v59 << (uint(int32(2)) % 32)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_get_keywords[5])))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_get_keywords[6])))
								v66 = v65
								v67 = v64
							} else {
								v66 = int32(0)
								v67 = v2
							}
							*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v66
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_pg_get_keywords[7]))))
							if v72 != 0 {
								v73 = int32(_a_F_pg_get_keywords_4)
							} else {
								v73 = int32(_a_F_pg_get_keywords_5)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v73
							if v72 != 0 {
								v77 = int32(_a_F_pg_get_keywords_6)
							} else {
								v77 = int32(_a_F_pg_get_keywords_7)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v77
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
							v80 = F_BuildTupleFromCStrings(m, v79, v11)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
								*(*int64)(unsafe.Add(mBase, uint32(v42))) = v82 + int64(1)
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = int32(1)
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								v90 = F_HeapTupleHeaderGetDatum(m, v89)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v104 = v90
									m.G0 = v11 + int32(32)
									return v104
								}
							}
						} else {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(2)
								v97 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v97)
								v104 = int32(0)
								m.G0 = v11 + int32(32)
								return v104
							}
						}
					}
				}
			}
		}
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
		v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
		v45 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[1])))
		if base.Ui64(v43) < base.Ui64(v45) {
			v49 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[2]))
			v51 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[3]))
			v52 = base.I32_wrap_i64(v43)
			v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(1))%32)))))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v49 + v56
			v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_pg_get_keywords[4]))))
			if base.Ui32(v59) <= base.Ui32(int32(3)) {
				v63 = v59 << (uint(int32(2)) % 32)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_get_keywords[5])))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_get_keywords[6])))
				v66 = v65
				v67 = v64
			} else {
				v66 = int32(0)
				v67 = v2
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v66
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_pg_get_keywords[7]))))
			if v72 != 0 {
				v73 = int32(_a_F_pg_get_keywords_4)
			} else {
				v73 = int32(_a_F_pg_get_keywords_5)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v73
			if v72 != 0 {
				v77 = int32(_a_F_pg_get_keywords_6)
			} else {
				v77 = int32(_a_F_pg_get_keywords_7)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v77
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
			v80 = F_BuildTupleFromCStrings(m, v79, v11)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				v82 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
				*(*int64)(unsafe.Add(mBase, uint32(v42))) = v82 + int64(1)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = int32(1)
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
				v90 = F_HeapTupleHeaderGetDatum(m, v89)
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					v104 = v90
					m.G0 = v11 + int32(32)
					return v104
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(2)
				v97 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v97)
				v104 = int32(0)
				m.G0 = v11 + int32(32)
				return v104
			}
		}
	}
}
func F_pg_get_next_timezone_abbrev(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 < v3 {
		v21 = v3
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+268))
		if v8 <= v5 {
			v21 = v3
		} else {
			v11 = int32(_a_F_pg_get_next_timezone_abbrev_0)
			v13 = F_strlen(m, l1+v5+v11)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v13 + v5 + int32(1)
			v21 = l1 + v11 + v5
		}
	}
	return v21
}
func F_pg_get_publication_tables(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v451 int32
	_ = v451
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v516 int32
	_ = v516
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int64
	_ = v626
	var v627 int64
	_ = v627
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int64
	_ = v638
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v769 int32
	_ = v769
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int64
	_ = v791
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v826 int32
	_ = v826
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v22 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)+16))
	goto L123
L4:
	;
	return int32(0)
L5:
	;
	v29 = int32(_a_F_pg_get_publication_tables_0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_tables[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_tables[0])) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = F_pg_detoast_datum(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_deconstruct_array_builtin(m, v35, int32(25), v19, int32(0), v19+int32(28))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v43 <= int32(0) {
		v562 = v2
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v569 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L4
	} else {
		goto L115
	}
L9:
	;
	v54 = v2
	v56 = v2
	v59 = v2
	goto L10
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v54<<(uint(int32(2))%32))))
	v68 = F_text_to_cstring(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	v399 = int32(0)
	if v394&base.B2i32(v387 != v399) == v399 {
		v562 = v387
		goto L8
	} else {
		goto L90
	}
L12:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	v394 = v393 | v59
	v396 = v54 + int32(1)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v396 < v397 {
		v54 = v396
		v56 = v387
		v59 = v394
		goto L10
	} else {
		goto L89
	}
L13:
	;
	if v337 == int32(0) {
		v387 = v56
		goto L12
	} else {
		goto L82
	}
L14:
	;
	v71 = F_get_publication_oid(m, v68, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v71 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v73 = F_GetPublication(m, v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	v75 = int32(0)
	goto L18
L18:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+8)))
	if v76 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v75 = v73
	goto L18
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	v80 = int32(0)
	v81 = m.G0
	v83 = v81 - int32(48)
	m.G0 = v83
	v87 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	v257 = F_GetPublicationRelations(m, v253, v254^int32(1))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L67
	}
L23:
	;
	F_ScanKeyInit(m, v83, int32(18), int32(3), int32(61), int32(114))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v96 = F_table_beginscan_catalog(m, v87, int32(1), v83)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v98 = F_heap_getnext(m, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	if v98 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v101 = v98
	v103 = v80
	goto L30
L28:
	;
	v143 = v80
	goto L29
L29:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+188))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	m.T0[v158].(func(*base.Module, int32))(m, v96)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L44
	}
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+22)))
	v118 = v116 + v117
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+119)))
	switch v119 - int32(112) {
	case 0, 2:
		goto L33
	default:
		v136 = v103
		goto L32
	}
L31:
	;
	v143 = v136
	goto L29
L32:
	;
	v138 = F_heap_getnext(m, v96)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L42
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	goto L34
L34:
	;
	if base.B2i32(base.Ui32(v122) < base.Ui32(int32(_a_F_pg_get_publication_tables_1)))|base.B2i32(base.Ui32(v122) < base.Ui32(int32(_a_F_pg_get_publication_tables_2))) != 0 {
		v136 = v103
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+118)))
	if v128 != int32(112) {
		v136 = v103
		goto L32
	} else {
		goto L36
	}
L36:
	;
	if v79 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+131)))
	if v131&int32(1) != 0 {
		v136 = v103
		goto L32
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v134 = F_lappend_oid(m, v103, v122)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v136 = v134
	goto L32
L42:
	;
	if v138 != 0 {
		v101 = v138
		v103 = v136
		goto L30
	} else {
		goto L43
	}
L43:
	;
	goto L31
L44:
	;
	if v79 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_ScanKeyInit(m, v83, int32(18), int32(3), int32(61), int32(112))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	v234 = v143
	goto L47
L47:
	;
	F_relation_close(m, v87, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L66
	}
L48:
	;
	v168 = F_table_beginscan_catalog(m, v87, int32(1), v83)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v170 = F_heap_getnext(m, v168)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	if v170 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v173 = v170
	v175 = v143
	goto L54
L52:
	;
	v213 = v143
	goto L53
L53:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+188))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	m.T0[v228].(func(*base.Module, int32))(m, v168)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L65
	}
L54:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+22)))
	v190 = v188 + v189
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+119)))
	switch v191 - int32(112) {
	case 0, 2:
		goto L57
	default:
		v206 = v175
		goto L56
	}
L55:
	;
	v213 = v206
	goto L53
L56:
	;
	v208 = F_heap_getnext(m, v168)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L63
	}
L57:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	goto L58
L58:
	;
	if base.B2i32(base.Ui32(v194) < base.Ui32(int32(_a_F_pg_get_publication_tables_1)))|base.B2i32(base.Ui32(v194) < base.Ui32(int32(_a_F_pg_get_publication_tables_2))) != 0 {
		v206 = v175
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+118)))
	if v200 != int32(112) {
		v206 = v175
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+131)))
	if v203 != 0 {
		v206 = v175
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v204 = F_lappend_oid(m, v175, v194)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v206 = v204
	goto L56
L63:
	;
	if v208 != 0 {
		v173 = v208
		v175 = v206
		goto L54
	} else {
		goto L64
	}
L64:
	;
	goto L55
L65:
	;
	v234 = v213
	goto L47
L66:
	;
	m.G0 = v83 + int32(48)
	v337 = v234
	goto L13
L67:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v261 = F_GetPublicationSchemas(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L69
	}
L68:
	;
	v319 = F_list_concat_unique_oid(m, v257, v306)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L81
	}
L69:
	;
	if v261 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v306 = int32(0)
	goto L68
L71:
	;
	goto L72
L72:
	;
	v266 = int32(0)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v267 <= v266 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v306 = int32(0)
	goto L68
L74:
	;
	goto L75
L75:
	;
	v275 = v266
	v277 = int32(0)
	goto L76
L76:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290+v275<<(uint(int32(2))%32))))
	v295 = F_GetSchemaPublicationRelations(m, v294, v259^int32(1))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L78
	}
L77:
	;
	v306 = v297
	goto L68
L78:
	;
	v297 = F_list_concat(m, v277, v295)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v300 = v275 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v300 < v301 {
		v275 = v300
		v277 = v297
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	v337 = v319
	goto L13
L82:
	;
	v340 = int32(0)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v341 <= v340 {
		v387 = v56
		goto L12
	} else {
		goto L83
	}
L83:
	;
	v345 = v340
	v354 = v56
	goto L84
L84:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v362 = F_palloc(m, int32(8))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
	} else {
		goto L86
	}
L85:
	;
	v387 = v371
	goto L12
L86:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v360+v345<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v367
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v369
	v371 = F_lappend(m, v354, v362)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v374 = v345 + int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v374 < v375 {
		v345 = v374
		v354 = v371
		goto L84
	} else {
		goto L88
	}
L88:
	;
	goto L85
L89:
	;
	goto L11
L90:
	;
	v407 = v387
	v408 = v399
	goto L91
L91:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v421 <= v408 {
		v562 = v387
		goto L8
	} else {
		goto L93
	}
L92:
	;
	v562 = v387
	goto L8
L93:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v407)+12))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v423+v408<<(uint(int32(2))%32))))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v429 = F_get_rel_relispartition(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L96
	}
L94:
	;
	if v536 != 0 {
		v407 = v536
		v408 = v535 + int32(1)
		goto L91
	} else {
		goto L114
	}
L95:
	;
	v535 = v408
	v536 = v407
	goto L94
L96:
	;
	if v429 == int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v434 = F_get_partition_ancestors(m, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	if v434 == int32(0) {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	if v438 <= int32(0) {
		goto L95
	} else {
		goto L100
	}
L100:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v434)+12))
	v451 = int32(0)
	goto L101
L101:
	;
	if v407 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L95
L103:
	;
	v516 = v451 + int32(1)
	if v438 != v516 {
		v451 = v516
		goto L101
	} else {
		goto L113
	}
L104:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v461 <= int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v441+v451<<(uint(int32(2))%32))))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v407)+12))
	v471 = int32(0)
	goto L106
L106:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v468+v471<<(uint(int32(2))%32))))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	if v467 != v490 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v497 = F_list_delete_nth_cell(m, v407, v408)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L112
	}
L108:
	;
	v493 = v471 + int32(1)
	if v493 != v461 {
		v471 = v493
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	goto L107
L111:
	;
	goto L103
L112:
	;
	v535 = v408 - int32(1)
	v536 = v497
	goto L94
L113:
	;
	goto L102
L114:
	;
	goto L92
L115:
	;
	F_TupleDescInitEntry(m, v569, int32(1), int32(_a_F_pg_get_publication_tables_3), int32(26), int32(-1), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	F_TupleDescInitEntry(m, v569, int32(2), int32(_a_F_pg_get_publication_tables_4), int32(26), int32(-1), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	F_TupleDescInitEntry(m, v569, int32(3), int32(_a_F_pg_get_publication_tables_5), int32(22), int32(-1), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	F_TupleDescInitEntry(m, v569, int32(4), int32(_a_F_pg_get_publication_tables_6), int32(194), int32(-1), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	v599 = F_BlessTupleDesc(m, v569)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v599
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_tables[0])) = v30
	goto L3
L121:
	;
	m.G0 = v19 + int32(32)
	return v826
L122:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L4
	} else {
		goto L159
	}
L123:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)+16))
	if v623 == int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v626 = *(*int64)(unsafe.Add(mBase, uint32(v622)))
	v627 = int64(*(*int32)(unsafe.Add(mBase, uint32(v623)+4)))
	if base.Ui64(v627) <= base.Ui64(v626) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v623)+12))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v629+base.I32_wrap_i64(v626)<<(uint(int32(2))%32))))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	v636 = F_get_rel_namespace(m, v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v638 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v638
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = int32(0)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v634)+4))
	v645 = F_GetPublication(m, v644)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v645)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v647
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+8)))
	if v650 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v622)+28))
	v789 = F_heap_form_tuple(m, v786, v19, v19+int32(28))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L157
	}
L129:
	;
	v685 = F_table_open(m, v635, int32(1))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L140
	}
L130:
	;
	v666 = v19 + int32(28)
	v669 = F_SysCacheGetAttr(m, int32(53), v658, int32(5), v666|int32(2))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L137
	}
L131:
	;
	v661 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+30)) = uint16(v661)
	goto L129
L132:
	;
	v652 = int32(0)
	v654 = F_SearchSysCacheExists(m, int32(50), v636, v647, v652, v652)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	if v654 != 0 {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v645)))
	v658 = F_SearchSysCacheCopy(m, int32(53), v635, v657)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	if v658 != 0 {
		goto L130
	} else {
		goto L136
	}
L136:
	;
	goto L131
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v669
	v676 = F_SysCacheGetAttr(m, int32(53), v658, int32(4), v666|int32(3))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v676
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)))
	if v679 != int32(1) {
		goto L128
	} else {
		goto L139
	}
L139:
	;
	goto L129
L140:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v685)+52))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v687)))
	v691 = F_palloc(m, v688<<(uint(int32(1))%32))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v687)))
	if v693 <= int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	F_relation_close(m, v685, int32(1))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L156
	}
L143:
	;
	v696 = int32(0)
	v699 = v696
	v700 = v696
	v701 = v693
	goto L144
L144:
	;
	v719 = v687 + v701<<(uint(int32(4))%32) + v699*int32(100)
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719)+111)))
	if v720 != 0 {
		v737 = v700
		v738 = v701
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if v737 <= int32(0) {
		goto L142
	} else {
		goto L154
	}
L146:
	;
	v742 = v699 + int32(1)
	if v742 < v738 {
		v699 = v742
		v700 = v737
		v701 = v738
		goto L144
	} else {
		goto L153
	}
L147:
	;
	v722 = v719 + int32(20)
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+90)))
	if v723 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	if v723 != int32(115) {
		v737 = v700
		v738 = v701
		goto L146
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v729 = int32(1)
	v732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v722)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v691+v700<<(uint(v729)%32)))) = uint16(v732)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v687)))
	v737 = v700 + v729
	v738 = v736
	goto L146
L151:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v645)+12))
	if v726 != int32(115) {
		v737 = v700
		v738 = v701
		goto L146
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	goto L145
L154:
	;
	v746 = F_buildint2vector(m, v691, v737)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v748 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)) = uint8(v748)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v746
	goto L142
L156:
	;
	goto L128
L157:
	;
	v791 = *(*int64)(unsafe.Add(mBase, uint32(v622)))
	*(*int64)(unsafe.Add(mBase, uint32(v622))) = v791 + int64(1)
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v795)+20)) = int32(1)
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v789)+16))
	v799 = F_HeapTupleHeaderGetDatum(m, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	v826 = v799
	goto L121
L159:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v804)+20)) = int32(2)
	v807 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v807)
	v826 = int32(0)
	goto L121
}
func F_pg_get_sequence_data(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v2)
	v15 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_TupleDescInitEntry(m, v15, int32(1), int32(_a_F_pg_get_sequence_data_0), int32(20), int32(-1), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v15, int32(2), int32(_a_F_pg_get_sequence_data_1), int32(16), int32(-1), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = F_BlessTupleDesc(m, v15)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v36 = F_try_relation_open(m, v9, int32(1))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v36 != 0 {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+119)))
							if v39 != int32(83) {
								v84 = int32(257)
								*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
								F_relation_close(m, v36, int32(1))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
										v98 = F_HeapTupleHeaderGetDatum(m, v97)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											m.G0 = v7 + int32(48)
											return v98
										}
									}
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[0]))
								v45 = F_pg_class_aclcheck(m, v9, v43, int64(2))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									if v45 != 0 {
										v84 = int32(257)
										*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
										F_relation_close(m, v36, int32(1))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
												v98 = F_HeapTupleHeaderGetDatum(m, v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 + int32(48)
													return v98
												}
											}
										}
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
										v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+118)))
										switch v48 - int32(112) {
										case 0:
											v71 = F_read_seq_tuple(m, v36, v7+int32(32), v7+int32(12))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												v73 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
												v74 = F_Int64GetDatum(m, v73)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v74
													v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+16)))
													*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v77
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
													F_UnlockReleaseBuffer(m, v79)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v36, int32(1))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
																v98 = F_HeapTupleHeaderGetDatum(m, v97)
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v7 + int32(48)
																	return v98
																}
															}
														}
													}
												}
											}
										default:
											v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[1])))
											if v56 == int32(1) {
												v61 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[2]))
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+316))
												v64 = base.B2i32(v62 != int32(2))
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[1])) = uint8(v64)
												v66 = v64
											} else {
												v66 = int32(0)
											}
											if v66 != 0 {
												v84 = int32(257)
												*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
												F_relation_close(m, v36, int32(1))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
														v98 = F_HeapTupleHeaderGetDatum(m, v97)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 + int32(48)
															return v98
														}
													}
												}
											} else {
												v71 = F_read_seq_tuple(m, v36, v7+int32(32), v7+int32(12))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													v73 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
													v74 = F_Int64GetDatum(m, v73)
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v74
														v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+16)))
														*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v77
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
														F_UnlockReleaseBuffer(m, v79)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v36, int32(1))
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int32(0)
															} else {
																v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return int32(0)
																} else {
																	v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
																	v98 = F_HeapTupleHeaderGetDatum(m, v97)
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v7 + int32(48)
																		return v98
																	}
																}
															}
														}
													}
												}
											}
										case 4:
											v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
											if v51 != int32(1) {
												v84 = int32(257)
												*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
												F_relation_close(m, v36, int32(1))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
														v98 = F_HeapTupleHeaderGetDatum(m, v97)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 + int32(48)
															return v98
														}
													}
												}
											} else {
												v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[1])))
												if v56 == int32(1) {
													v61 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[2]))
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+316))
													v64 = base.B2i32(v62 != int32(2))
													*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[1])) = uint8(v64)
													v66 = v64
												} else {
													v66 = int32(0)
												}
												if v66 != 0 {
													v84 = int32(257)
													*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v84)
													F_relation_close(m, v36, int32(1))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int32(0)
														} else {
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
															v98 = F_HeapTupleHeaderGetDatum(m, v97)
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																m.G0 = v7 + int32(48)
																return v98
															}
														}
													}
												} else {
													v71 = F_read_seq_tuple(m, v36, v7+int32(32), v7+int32(12))
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														v73 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
														v74 = F_Int64GetDatum(m, v73)
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v74
															v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+16)))
															*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v77
															v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
															F_UnlockReleaseBuffer(m, v79)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v36, int32(1))
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return int32(0)
																} else {
																	v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int32(0)
																	} else {
																		v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
																		v98 = F_HeapTupleHeaderGetDatum(m, v97)
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v7 + int32(48)
																			return v98
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
						} else {
							v82 = int32(257)
							*(*uint16)(unsafe.Add(mBase, uint32(v7)+38)) = uint16(v82)
							v95 = F_heap_form_tuple(m, v33, v7+int32(40), v7+int32(38))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
								v98 = F_HeapTupleHeaderGetDatum(m, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(48)
									return v98
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_get_shmem_allocations(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[0]))
	v21 = F_LWLockAcquire(m, v17+int32(128), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v6 + int32(-20)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[1]))
	F_hash_seq_init(m, v24, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	v31 = F_hash_seq_search(m, v24)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v33 = v31
	v36 = v2
	goto L9
L7:
	;
	v74 = v2
	goto L8
L8:
	;
	v77 = F_cstring_to_text(m, int32(_a_F_pg_get_shmem_allocations_0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L9:
	;
	v38 = F_cstring_to_text(m, v33)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v74 = v66
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[2]))
	v46 = F_Int64GetDatum(m, base.I64_extend_i32_s(v41-v43))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v46
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+52)))
	v50 = F_Int64GetDatum(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v50
	v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+56)))
	v54 = F_Int64GetDatum(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v33)+56))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_tuplestore_putvalues(m, v58, v59, v6+int32(-48), v6+int32(-52))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v66 = v57 + v36
	v69 = F_hash_seq_search(m, v6+int32(-20))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v69 != 0 {
		v33 = v69
		v36 = v66
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L10
L18:
	;
	v79 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v77
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[2]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v87 = F_Int64GetDatum(m, base.I64_extend_i32_u(v84-v74))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v87
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v94 = v6 + int32(-48)
	v96 = v6 + int32(-52)
	F_tuplestore_putvalues(m, v91, v92, v94, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v99)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[2]))
	v103 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v102)+12)))
	v104 = F_Int64GetDatum(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v104
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[2]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v111-v112))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v115
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_tuplestore_putvalues(m, v119, v120, v94, v96)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[0]))
	F_LWLockRelease(m, v124+int32(128))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v8 - int32(-64)
	return int32(0)
}
func F_pg_get_shmem_allocations_numa(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v2 = m.G0
	m.G0 = v2 + int32(-64)
	F_errstart_cold(m, int32(21), int32(0))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errmsg_internal(m, int32(_a_F_pg_get_shmem_allocations_numa_0), int32(0))
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(_a_F_pg_get_shmem_allocations_numa_1), int32(601), int32(_a_F_pg_get_shmem_allocations_numa_2))
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_pg_getaddrinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v205
L2:
	;
	v12 = int32(-4)
	v13 = F_strlen(m, l1)
	mBase = m.M
	if base.Ui32(int32(107)) < base.Ui32(v13) {
		v205 = v12
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if l0 != 0 {
		goto L56
	} else {
		goto L57
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v16 != int32(1) {
		v205 = v12
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	goto L10
L7:
	;
	if v44 == int32(0) {
		v205 = int32(-10)
		goto L1
	} else {
		goto L17
	}
L8:
	;
	goto L7
L9:
	;
	v44 = F_emscripten_builtin_malloc(m, v32)
	mBase = m.M
	if v44 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L10:
	;
	v32 = base.I32_wrap_i64(base.I64_extend_i32_u(int32(1)) * base.I64_extend_i32_u(int32(32)))
	goto L9
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44-int32(4)))))
	if v49&int32(3) == int32(0) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	F___memset(m, v44, int32(0), v32)
	mBase = m.M
	goto L8
L17:
	;
	goto L21
L18:
	;
	if v80 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	goto L18
L20:
	;
	v80 = F_emscripten_builtin_malloc(m, v68)
	mBase = m.M
	if v80 == int32(0) {
		goto L19
	} else {
		goto L26
	}
L21:
	;
	v68 = base.I32_wrap_i64(base.I64_extend_i32_u(int32(1)) * base.I64_extend_i32_u(int32(110)))
	goto L20
L26:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80-int32(4)))))
	if v85&int32(3) == int32(0) {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	F___memset(m, v80, int32(0), v68)
	mBase = m.M
	goto L19
L28:
	;
	F_emscripten_builtin_free(m, v44)
	mBase = m.M
	return int32(-10)
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v20
	v98 = int32(1)
	if base.Ui32(v19) <= base.Ui32(v98) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v101 = v98
	goto L33
L32:
	;
	v101 = v19
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v101
	v103 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v44
	*(*uint16)(unsafe.Add(mBase, uint32(v80))) = uint16(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = int32(110)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v80
	v112 = v80 + int32(2)
	if (l1^v112)&int32(3) != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v188 != int32(64) {
		v205 = int32(0)
		goto L1
	} else {
		goto L55
	}
L35:
	;
	goto L34
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v166)
	if v166&int32(255) == int32(0) {
		goto L35
	} else {
		goto L51
	}
L37:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v165 = l1
	v166 = v118
	v167 = v112
	goto L36
L38:
	;
	goto L39
L39:
	;
	if l1&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v122 = l1
	v124 = v112
	goto L43
L41:
	;
	v136 = l1
	v138 = v112
	goto L42
L42:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v143 = int32(-2139062144)
	if (int32(16843008)-v140|v140)&v143 != v143 {
		v165 = v136
		v166 = v140
		v167 = v138
		goto L36
	} else {
		goto L47
	}
L43:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v125)
	if v125 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L44:
	;
	v136 = v132
	v138 = v130
	goto L42
L45:
	;
	v129 = int32(1)
	v130 = v124 + v129
	v132 = v122 + v129
	if v132&int32(3) != 0 {
		v122 = v132
		v124 = v130
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v148 = v136
	v149 = v140
	v150 = v138
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v149
	v152 = int32(4)
	v153 = v150 + v152
	v155 = v148 + v152
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v160 = int32(-2139062144)
	if (int32(16843008)-v157|v157)&v160 == v160 {
		v148 = v155
		v149 = v157
		v150 = v153
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v165 = v155
	v166 = v157
	v167 = v153
	goto L36
L50:
	;
	goto L49
L51:
	;
	v174 = v165
	v176 = v167
	goto L52
L52:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)) = uint8(v177)
	v179 = int32(1)
	if v177 != 0 {
		v174 = v174 + v179
		v176 = v176 + v179
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L35
L54:
	;
	goto L53
L55:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v191)
	v193 = F_strlen(m, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v193 + int32(2)
	return v191
L56:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v200 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v203 = int32(0)
	goto L58
L58:
	;
	v204 = m.Env.Getaddrinfo(m, v203, l1, l2, l3)
	mBase = m.M
	v205 = v204
	goto L1
L59:
	;
	v201 = l0
	goto L61
L60:
	;
	v201 = int32(0)
	goto L61
L61:
	;
	v203 = v201
	goto L58
}
func F_pg_has_role_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_has_role_id[0]))
		v12 = F_convert_any_priv_string(m, v5, int32(_a_F_pg_has_role_id_0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_pg_role_aclcheck(m, v3, v10, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14 ^ int32(1)
			}
		}
	}
}
func F_pg_has_role_id_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = F_convert_any_priv_string(m, v5, int32(_a_F_pg_has_role_id_id_0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_pg_role_aclcheck(m, v2, v3, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12 ^ int32(1)
			}
		}
	}
}
func F_pg_has_role_name_id(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v21 = F_GetSysCacheOid(m, int32(10), v11, v18, v18, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
						F_errmsg(m, int32(_a_F_pg_has_role_name_id_0), v8)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_has_role_name_id_1), int32(_a_F_pg_has_role_name_id_2), int32(_a_F_pg_has_role_name_id_3))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v42 = F_convert_any_priv_string(m, v13, int32(_a_F_pg_has_role_name_id_4))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = F_pg_role_aclcheck(m, v10, v21, v42)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v44 ^ int32(1)
					}
				}
			}
		}
	}
}
func F_pg_hmac_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	if l0 == int32(0) {
		return int32(-1)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = F_pg_cryptohash_update(m, v8, l1, l2)
		mBase = m.M
		if int32(0) <= v9 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v16 == int32(0) {
				v31 = int32(_a_F_pg_hmac_update_0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				if v23 == int32(1) {
					v26 = int32(_a_F_pg_hmac_update_1)
				} else {
					v26 = int32(_a_F_pg_hmac_update_2)
				}
				if v23 == int32(2) {
					v29 = int32(_a_F_pg_hmac_update_0)
				} else {
					v29 = v26
				}
				v31 = v29
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
			return int32(-1)
		}
	}
}
func F_pg_jit_available(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_provider_init(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_pg_johab_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	if l1 <= int32(0) {
		v46 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v46 - l0
L2:
	;
	v9 = l1
	v10 = l0
	goto L3
L3:
	;
	v13 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	if int32(0) <= v13 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v46 = v40
	goto L1
L5:
	;
	v40 = v10 + v39
	v41 = v9 - v39
	if int32(0) < v41 {
		v9 = v41
		v10 = v40
		goto L3
	} else {
		goto L17
	}
L6:
	;
	if v13 != 0 {
		v39 = int32(1)
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v13 == int32(-113) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v46 = v10
	goto L1
L10:
	;
	v21 = int32(3)
	goto L12
L11:
	;
	v21 = int32(2)
	goto L12
L12:
	;
	if base.Ui32(v9) < base.Ui32(v21) {
		v46 = v10
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v23+int32(95))&int32(255)) {
		v46 = v10
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v13 != int32(-113) {
		v39 = v21
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+2)))
	if base.Ui32(int32(93)) < base.Ui32((v32+int32(95))&int32(255)) {
		v46 = v10
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v39 = v21
	goto L5
L17:
	;
	goto L4
}
func F_pg_mb2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mb2wchar_with_len[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6*int32(28))+uint32(_c_F_pg_mb2wchar_with_len[1])))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_pg_mblen_unbounded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mblen_unbounded[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4*int32(28))+uint32(_c_F_pg_mblen_unbounded[1])))
	v10 = m.T0[v9].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_pg_md5_hash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v15 = F_pg_cryptohash_create(m, v5)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_pg_md5_hash_0)
			v101 = v5
			m.G0 = v10 + int32(16)
			return v101
		} else {
			v38 = F_pg_cryptohash_init(m, v15)
			mBase = m.M
			if v38 < int32(0) {
				if v15 == int32(0) {
					v92 = int32(_a_F_pg_md5_hash_0)
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					if v84 == int32(1) {
						v87 = int32(_a_F_pg_md5_hash_1)
					} else {
						v87 = int32(_a_F_pg_md5_hash_2)
					}
					if v84 == int32(2) {
						v90 = int32(_a_F_pg_md5_hash_0)
					} else {
						v90 = v87
					}
					v92 = v90
				}
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v92
				F_pg_cryptohash_free(m, v15)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					v101 = v5
					m.G0 = v10 + int32(16)
					return v101
				}
			} else {
				v41 = F_pg_cryptohash_update(m, v15, l0, l1)
				mBase = m.M
				if v41 < int32(0) {
					if v15 == int32(0) {
						v92 = int32(_a_F_pg_md5_hash_0)
					} else {
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						if v84 == int32(1) {
							v87 = int32(_a_F_pg_md5_hash_1)
						} else {
							v87 = int32(_a_F_pg_md5_hash_2)
						}
						if v84 == int32(2) {
							v90 = int32(_a_F_pg_md5_hash_0)
						} else {
							v90 = v87
						}
						v92 = v90
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v92
					F_pg_cryptohash_free(m, v15)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v101 = v5
						m.G0 = v10 + int32(16)
						return v101
					}
				} else {
					v45 = F_pg_cryptohash_final(m, v15, v10, int32(16))
					mBase = m.M
					if v45 < int32(0) {
						if v15 == int32(0) {
							v92 = int32(_a_F_pg_md5_hash_0)
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
							if v84 == int32(1) {
								v87 = int32(_a_F_pg_md5_hash_1)
							} else {
								v87 = int32(_a_F_pg_md5_hash_2)
							}
							if v84 == int32(2) {
								v90 = int32(_a_F_pg_md5_hash_0)
							} else {
								v90 = v87
							}
							v92 = v90
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v92
						F_pg_cryptohash_free(m, v15)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v101 = v5
							m.G0 = v10 + int32(16)
							return v101
						}
					} else {
						v52 = int32(0)
						v54 = v5
						for {
							v56 = l2 + v54
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v10))))
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58&int32(15))+uint32(_c_F_pg_md5_hash[0]))))
							*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v61)
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v58)>>(uint(int32(4))%32)))+uint32(_c_F_pg_md5_hash[0]))))
							*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v65)
							v70 = v52 + int32(1)
							if v70 != int32(16) {
								v52 = v70
								v54 = v54 + int32(2)
								continue
							} else {
								break
							}
							break
						}
						v73 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)) = uint8(v73)
						F_pg_cryptohash_free(m, v15)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							v101 = int32(1)
							m.G0 = v10 + int32(16)
							return v101
						}
					}
				}
			}
		}
	}
}
func F_pg_mule2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v19 == int32(0) {
		v107 = v14
		v111 = v18
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(0)
	return v111
L6:
	;
	goto L5
L7:
	;
	if base.Ui32((v19+int32(127))&int32(255)) <= base.Ui32(int32(12)) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v95
	v100 = v18 + int32(1)
	v102 = v14 + int32(4)
	v103 = v15 + v96
	if int32(0) < v103 {
		v13 = v97
		v14 = v102
		v15 = v103
		v18 = v100
		goto L4
	} else {
		goto L25
	}
L9:
	;
	if v15 == int32(1) {
		v107 = v14
		v111 = v18
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v39 = v19 & int32(254)
	if v39 == int32(154) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v31 = v19 << (uint(int32(16)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v95 = v31 | v33
	v96 = int32(-2)
	v97 = v13 + int32(2)
	goto L8
L13:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v107 = v14
		v111 = v18
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if base.Ui32((v19+int32(112))&int32(255)) <= base.Ui32(int32(9)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v46 = v44 << (uint(int32(16)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v95 = v46 | v48
	v96 = int32(-3)
	v97 = v13 + int32(3)
	goto L8
L17:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v107 = v14
		v111 = v18
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v39 == int32(156) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v62 = v19 << (uint(int32(16)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v62
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v67 = v64<<(uint(int32(8))%32) | v62
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v95 = v67 | v69
	v96 = int32(-3)
	v97 = v13 + int32(3)
	goto L8
L21:
	;
	if base.Ui32(v15) < base.Ui32(int32(4)) {
		v107 = v14
		v111 = v18
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v95 = v19
	v96 = int32(-1)
	v97 = v13 + int32(1)
	goto L8
L24:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v80 = v78 << (uint(int32(16)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v85 = v82<<(uint(int32(8))%32) | v80
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v85
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	v95 = v85 | v87
	v96 = int32(-4)
	v97 = v13 + int32(4)
	goto L8
L25:
	;
	v107 = v102
	v111 = v100
	goto L6
}
func F_pg_my_temp_schema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_my_temp_schema[0]))
	return v3
}
func F_pg_node_tree_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_pg_node_tree_in_0), int32(334), int32(_a_F_pg_node_tree_in_1), int32(_a_F_pg_node_tree_in_2), int32(_a_F_pg_node_tree_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_notify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v4 = int32(_a_F_pg_notify_0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_text_to_cstring(m, v10)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = v14
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v17 == int32(0) {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v21 = F_pg_detoast_datum_packed(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = F_text_to_cstring(m, v21)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v25 = v23
							F_PreventCommandDuringRecovery(m, int32(_a_F_pg_notify_1))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								F_Async_Notify(m, v16, v25)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						}
					}
				} else {
					v25 = v4
					F_PreventCommandDuringRecovery(m, int32(_a_F_pg_notify_1))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_Async_Notify(m, v16, v25)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v16 = v4
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v17 == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v21 = F_pg_detoast_datum_packed(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_text_to_cstring(m, v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = v23
					F_PreventCommandDuringRecovery(m, int32(_a_F_pg_notify_1))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_Async_Notify(m, v16, v25)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			}
		} else {
			v25 = v4
			F_PreventCommandDuringRecovery(m, int32(_a_F_pg_notify_1))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_Async_Notify(m, v16, v25)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_pg_num_nulls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v11 = F_count_nulls(m, l0, v5+int32(12), v5+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v21 = int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			v21 = v20
		}
		m.G0 = v5 + int32(16)
		return v21
	}
}
func F_pg_parse_json_or_errsave(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v4 = F_pg_parse_json(m, l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			F_json_errsave_error(m, v4, l0, l2)
			v9 = m.ExcPending
			if v9 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v4 == int32(0))
			}
		} else {
			return base.B2i32(v4 == int32(0))
		}
	}
}
func F_pg_perm_setlocale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
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
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v143 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v143 != 0 {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v17 + int32(48)
	goto L1
L3:
	;
	if l0 == int32(6) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v104 = int32(0)
	v105 = int32(_a_F_pg_perm_setlocale_0)
	v110 = v3
	goto L28
L5:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if l1 != 0 {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v26
	v29 = *(*int64)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v29
	v32 = *(*int64)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v32
	v35 = int32(0)
	v36 = l1
	goto L10
L9:
	;
	v143 = int32(0)
	goto L2
L10:
	;
	v44 = F___strchrnul(m, v36, int32(59))
	mBase = m.M
	v45 = v44 - v36
	if v45 <= int32(23) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int64)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[3])) = v71
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[4])) = v74
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[5])) = v77
	goto L4
L12:
	;
	v48 = F___memcpy(m, v17, v36, v45)
	mBase = m.M
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+v45))) = uint8(v50)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v56 = v36
	goto L14
L14:
	;
	v57 = F___get_locale(m, v35, v17)
	mBase = m.M
	if v57 == int32(-1) {
		goto L9
	} else {
		goto L18
	}
L15:
	;
	v55 = v44 + int32(1)
	goto L17
L16:
	;
	v55 = v36
	goto L17
L17:
	;
	v56 = v55
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(24)+v35<<(uint(int32(2))%32)))) = v57
	v67 = v35 + int32(1)
	if v67 != int32(6) {
		v35 = v67
		v36 = v56
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L11
L20:
	;
	if v89 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v80 = F___get_locale(m, l0, l1)
	mBase = m.M
	if v80 == int32(-1) {
		v143 = v3
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_pg_perm_setlocale[5])))
	v89 = v88
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_pg_perm_setlocale[5]))) = v80
	v89 = v80
	goto L20
L25:
	;
	v93 = v89 + int32(8)
	goto L27
L26:
	;
	v93 = int32(_a_F_pg_perm_setlocale_1)
	goto L27
L27:
	;
	v143 = v93
	goto L2
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[5]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v104<<(uint(int32(2))%32))+uint32(_c_F_pg_perm_setlocale[5])))
	if v116 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v134 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v134)
	if v129 != int32(6) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v120 = v116 + int32(8)
	goto L32
L31:
	;
	v120 = int32(_a_F_pg_perm_setlocale_1)
	goto L32
L32:
	;
	v121 = F_strlen(m, v120)
	mBase = m.M
	v122 = F___memcpy(m, v105, v120, v121)
	mBase = m.M
	v123 = v105 + v121
	v124 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v124)
	v126 = int32(1)
	v129 = v110 + base.B2i32(v116 == v113)
	v131 = v104 + v126
	if v131 != int32(6) {
		v104 = v131
		v105 = v123 + v126
		v110 = v129
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v139 = int32(_a_F_pg_perm_setlocale_0)
	goto L36
L35:
	;
	v139 = v120
	goto L36
L36:
	;
	v143 = v139
	goto L2
L37:
	;
	switch l0 {
	case 0:
		goto L41
	case 1:
		goto L44
	case 2:
		goto L43
	case 3:
		v302 = v143
		v303 = int32(_a_F_pg_perm_setlocale_2)
		goto L40
	case 4:
		goto L45
	case 5:
		goto L46
	default:
		goto L42
	}
L38:
	;
	v430 = int32(0)
	goto L39
L39:
	;
	m.G0 = v7 + int32(16)
	return v430
L40:
	;
	v304 = int32(0)
	if v303 == v304 {
		goto L87
	} else {
		goto L88
	}
L41:
	;
	v171 = int32(_a_F_pg_perm_setlocale_3)
	goto L54
L42:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v302 = v143
	v303 = int32(_a_F_pg_perm_setlocale_4)
	goto L40
L44:
	;
	v302 = v143
	v303 = int32(_a_F_pg_perm_setlocale_5)
	goto L40
L45:
	;
	v302 = v143
	v303 = int32(_a_F_pg_perm_setlocale_6)
	goto L40
L46:
	;
	v302 = v143
	v303 = int32(_a_F_pg_perm_setlocale_7)
	goto L40
L47:
	;
	return int32(0)
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_perm_setlocale_8), v7)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_pg_perm_setlocale_9), int32(279), int32(_a_F_pg_perm_setlocale_10))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[6]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	goto L82
L52:
	;
	v288 = F_strlen(m, v277)
	mBase = m.M
	goto L51
L54:
	;
	goto L55
L55:
	;
	v178 = int32(127)
	if (v171^v143)&int32(3) != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v281 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v278))) = uint8(v281)
	goto L52
L57:
	;
	v262 = v257
	v263 = v258
	v264 = v259
	goto L78
L58:
	;
	if v252 == int32(0) {
		v277 = v250
		v278 = v251
		goto L56
	} else {
		goto L77
	}
L59:
	;
	v250 = v143
	v251 = v171
	v252 = v178
	goto L58
L60:
	;
	goto L61
L61:
	;
	v182 = int32(0)
	if base.B2i32(v143&int32(3) == v182)|int32(0) == v182 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v218 == int32(0) {
		v277 = v215
		v278 = v216
		goto L56
	} else {
		goto L71
	}
L63:
	;
	v194 = v143
	v195 = v171
	v196 = v178
	goto L66
L64:
	;
	goto L65
L65:
	;
	v215 = v143
	v216 = v171
	v217 = v178
	v218 = int32(1)
	goto L62
L66:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v198)
	if v198 == int32(0) {
		v257 = v194
		v258 = v195
		v259 = v196
		goto L57
	} else {
		goto L68
	}
L67:
	;
	v215 = v209
	v216 = v203
	v217 = v205
	v218 = v207
	goto L62
L68:
	;
	v202 = int32(1)
	v203 = v195 + v202
	v205 = v196 - v202
	v206 = int32(0)
	v207 = base.B2i32(v205 != v206)
	v209 = v194 + v202
	if v209&int32(3) == v206 {
		v215 = v209
		v216 = v203
		v217 = v205
		v218 = v207
		goto L62
	} else {
		goto L69
	}
L69:
	;
	if v205 != 0 {
		v194 = v209
		v195 = v203
		v196 = v205
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if base.B2i32(v221 == int32(0))|base.B2i32(base.Ui32(v217) < base.Ui32(int32(4))) != 0 {
		v250 = v215
		v251 = v216
		v252 = v217
		goto L58
	} else {
		goto L72
	}
L72:
	;
	v228 = v215
	v229 = v216
	v230 = v217
	goto L73
L73:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v236 = int32(-2139062144)
	if (int32(16843008)-v233|v233)&v236 != v236 {
		v257 = v228
		v258 = v229
		v259 = v230
		goto L57
	} else {
		goto L75
	}
L74:
	;
	v250 = v244
	v251 = v242
	v252 = v246
	goto L58
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v233
	v241 = int32(4)
	v242 = v229 + v241
	v244 = v228 + v241
	v246 = v230 - v241
	if base.Ui32(int32(3)) < base.Ui32(v246) {
		v228 = v244
		v229 = v242
		v230 = v246
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v257 = v250
	v258 = v251
	v259 = v252
	goto L57
L78:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	*(*uint8)(unsafe.Add(mBase, uint32(v263))) = uint8(v266)
	if v266 == int32(0) {
		v277 = v262
		v278 = v263
		goto L56
	} else {
		goto L80
	}
L79:
	;
	v277 = v273
	v278 = v271
	goto L56
L80:
	;
	v270 = int32(1)
	v271 = v263 + v270
	v273 = v262 + v270
	v275 = v264 - v270
	if v275 != 0 {
		v262 = v273
		v263 = v271
		v264 = v275
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[7])) = v294<<(uint(int32(3))%32) + int32(_a_F_pg_perm_setlocale_11)
	v302 = int32(_a_F_pg_perm_setlocale_3)
	v303 = int32(_a_F_pg_perm_setlocale_12)
	goto L40
L83:
	;
	if v425 != 0 {
		goto L122
	} else {
		goto L123
	}
L84:
	;
	v335 = F___memcpy(m, v330, v303, v313)
	mBase = m.M
	v336 = v330 + v313
	v337 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v336))) = uint8(v337)
	v339 = int32(1)
	v343 = F___memcpy(m, v336+v339, v302, v326+v339)
	mBase = m.M
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[8]))
	if v345 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L85:
	;
	v425 = int32(-1)
	goto L83
L86:
	;
	goto L91
L87:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[9])) = int32(28)
	goto L85
L88:
	;
	v311 = F___strchrnul(m, v303, int32(61))
	mBase = m.M
	if v311 == v303 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v313 = v311 - v303
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v313))))
	if v315 == int32(0) {
		goto L86
	} else {
		goto L90
	}
L90:
	;
	goto L87
L91:
	;
	v326 = F_strlen(m, v302)
	mBase = m.M
	v330 = F_emscripten_builtin_malloc(m, v313+v326+int32(2))
	mBase = m.M
	if v330 != 0 {
		goto L84
	} else {
		goto L94
	}
L94:
	;
	goto L85
L95:
	;
	v425 = v419
	goto L83
L96:
	;
	v381 = v376 << (uint(int32(2)) % 32)
	v383 = v381 + int32(8)
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[10]))
	if v375 == v385 {
		goto L111
	} else {
		goto L112
	}
L97:
	;
	v356 = v345
	v357 = int32(0)
	v360 = v349
	goto L103
L98:
	;
	v375 = v350
	v376 = int32(0)
	goto L96
L99:
	;
	v350 = int32(0)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	if v349 != 0 {
		goto L97
	} else {
		goto L102
	}
L102:
	;
	v350 = v345
	goto L98
L103:
	;
	v361 = F_strncmp(m, v330, v360, v313+int32(1))
	mBase = m.M
	if v361 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[8]))
	v375 = v374
	v376 = v369
	goto L96
L105:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v330
	F___env_rm_add(m, v364, v330)
	mBase = m.M
	v419 = int32(0)
	goto L95
L106:
	;
	goto L107
L107:
	;
	v369 = v357 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	if v370 != 0 {
		v356 = v356 + int32(4)
		v357 = v369
		v360 = v370
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	F_emscripten_builtin_free(m, v330)
	mBase = m.M
	v419 = int32(-1)
	goto L95
L110:
	;
	v400 = v397 + v376<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v400))) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v400)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[8])) = v397
	*(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[10])) = v397
	if v330 != 0 {
		goto L119
	} else {
		goto L120
	}
L111:
	;
	v387 = F_emscripten_builtin_realloc(m, v385, v383)
	mBase = m.M
	if v387 != 0 {
		v397 = v387
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v388 = F_emscripten_builtin_malloc(m, v383)
	mBase = m.M
	if v388 == int32(0) {
		goto L109
	} else {
		goto L115
	}
L114:
	;
	goto L109
L115:
	;
	if v376 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[8]))
	v393 = F___memcpy(m, v388, v392, v381)
	mBase = m.M
	goto L118
L117:
	;
	goto L118
L118:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_pg_perm_setlocale[10]))
	F_emscripten_builtin_free(m, v395)
	mBase = m.M
	v397 = v388
	goto L110
L119:
	;
	F___env_rm_add(m, int32(0), v330)
	mBase = m.M
	goto L121
L120:
	;
	goto L121
L121:
	;
	v419 = int32(0)
	goto L95
L122:
	;
	v426 = v304
	goto L124
L123:
	;
	v426 = v302
	goto L124
L124:
	;
	v430 = v426
	goto L39
}
func F_pg_range_sockaddr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	switch v6 - int32(2) {
	case 0:
		goto L3
	default:
		goto L1
	case 8:
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = int32(8)
	v18 = l2 + v17
	v20 = l1 + v17
	v22 = l0 + v17
	v24 = int32(0)
	goto L4
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return base.B2i32(v9&(v10^v11) == int32(0))
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v18))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v20))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v22))))
	if v30&(v32^v34) != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	return int32(1)
L6:
	;
	v38 = v24 | int32(1)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v38))))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v38))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v38))))
	if (v40^v42)&v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v48 = v24 + int32(2)
	if v48 != int32(16) {
		v24 = v48
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_pg_read_binary_file_all(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_convert_and_check_filename(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v13 = F_read_binary_file(m, v8, int64(0), int64(-1), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v13 == int32(0) {
					v17 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
					return int32(0)
				} else {
					return v13
				}
			}
		}
	}
}
func F_pg_read_file_all(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_convert_and_check_filename(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v13 = F_read_binary_file(m, v8, int64(0), int64(-1), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v13 == int32(0) {
					v17 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
					return int32(0)
				} else {
					v21 = int32(4)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					F_pg_verifymbstr(m, v13+v21, int32(base.Ui32(v23)>>(uint(int32(2))%32))-v21)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						return v13
					}
				}
			}
		}
	}
}
func F_pg_read_file_off_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		v13 = F_pg_read_file_common(m, v4, v9, v11, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v13 == int32(0) {
				v17 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_pg_regcomp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int64
	_ = v146
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v852 int32
	_ = v852
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v881 int32
	_ = v881
	var v894 int32
	_ = v894
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v975 int32
	_ = v975
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v994 int32
	_ = v994
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v18 = int32(2)
	if base.B2i32(l0 == v6)|base.B2i32(l1 == v6)|(int32(base.Ui32(l3)>>(uint(v18)%32))&base.B2i32(l3&int32(227) != v6)|base.B2i32(l3&int32(3) == v18)) == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pg_set_regex_collation(m, l4)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v994 = int32(16)
	goto L3
L3:
	;
	m.G0 = v10 + int32(144)
	return v994
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v10 + int32(52)
	v42 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1 + l2<<(uint(int32(2))%32)
	v57 = v42
	goto L6
L6:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v61+v57<<(uint(int32(2))%32)))) = int32(0)
	v68 = v57 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	if base.Ui32(v68) < base.Ui32(v69) {
		v57 = v68
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v71 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v71
	v73 = int32(_a_F_pg_regcomp_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+100)) = uint16(v73)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+92)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v10)+124)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v10)+132)) = v71
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+140)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(_a_F_pg_regcomp_1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(17179869184)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_pg_regcomp_2)
	v96 = F_palloc_extended(m, int32(432), int32(2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v96
	if v96 == int32(0) {
		v815 = int32(12)
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v822 = v10 + int32(4)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v822)))
	if v823 != 0 {
		goto L210
	} else {
		goto L211
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+72)) = int32(2166)
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+16)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v96)+200)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v96)+192)) = int64(0)
	v110 = int32(_a_F_pg_regcomp_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+188)) = uint16(v110)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+180)) = int64(4294969344)
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+88)) = uint16(v104)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+80)) = int64(10)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+92)) = v96 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+76)) = v10 + int32(4)
	v126 = F_palloc_extended(m, int32(_a_F_pg_regcomp_3), int32(2))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+96)) = v126
	if v126 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v96 + int32(72)
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+20)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v96)+424)) = int64(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	v193 = F_newnfa(m, v10+int32(4), v191, v185)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L27
	}
L14:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v96)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+24)) = int32(101)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v96)+76))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	if v137 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	base.MemoryFill(m, v126, int32(0), int32(_a_F_pg_regcomp_3))
	v146 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+156)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v96)+148)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v96)+140)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v96)+132)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v96)+124)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v96)+116)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v96)+108)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v96)+100)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v96)+176)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+168)) = int64(4294967300)
	v168 = F_palloc_extended(m, int32(8), int32(2))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L20
	}
L17:
	;
	v139 = v137
	goto L19
L18:
	;
	v139 = int32(12)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+12)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v96)+160)) = int64(0)
	goto L13
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+164)) = v168
	if v168 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v96)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = int32(101)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v96)+76))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	if v177 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v181 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v168))) = uint16(v181)
	goto L13
L24:
	;
	v179 = v177
	goto L26
L25:
	;
	v179 = int32(12)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+12)) = v179
	goto L13
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = v193
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v196 != 0 {
		v815 = v196
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v199 = F_palloc_extended(m, int32(588), int32(2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if v199 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = int32(0)
	v815 = int32(12)
	goto L10
L31:
	;
	goto L32
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v199))) = int64(429496729600)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v199)+12)) = int64(85899345920)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+20)) = v199 + int32(428)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v199 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = v199
	v219 = int32(4)
	v220 = v10 + v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	if v221&v219 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	if v605 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L34:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v224-v225 < int32(13) {
		v281 = v221
		v282 = v225
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v283 = int32(3)
	if v281&v283 != v283 {
		goto L33
	} else {
		goto L50
	}
L36:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v229 != int32(42) {
		v281 = v221
		v282 = v225
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v232 != int32(42) {
		v281 = v221
		v282 = v225
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	if v235 != int32(42) {
		v281 = v221
		v282 = v225
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	switch v238 - int32(58) {
	case 0:
		goto L40
	default:
		goto L41
	case 3:
		goto L42
	case 5:
		goto L43
	}
L40:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v268)+8)) = v269 | int32(128)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v275 = v273 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+16)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v279 = v277 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v279
	v281 = v275
	v282 = v279
	goto L35
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = int32(101)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	if v264 != 0 {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+8)) = v248 | int32(128)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v252 + int32(16)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+16)) = v256&int32(-232) | int32(4)
	goto L33
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = int32(101)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	if v243 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v245 = v243
	goto L46
L45:
	;
	v245 = int32(2)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = v245
	goto L33
L47:
	;
	v266 = v264
	goto L49
L48:
	;
	v266 = int32(13)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = v266
	goto L33
L50:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	if v287-v282 < int32(9) {
		goto L33
	} else {
		goto L51
	}
L51:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v291 != int32(40) {
		goto L33
	} else {
		goto L52
	}
L52:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v294 != int32(63) {
		goto L33
	} else {
		goto L53
	}
L53:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_pg_regcomp[0]))
	switch v299 - int32(1) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	default:
		goto L58
	}
L54:
	;
	if v391 == int32(0) {
		goto L33
	} else {
		goto L80
	}
L55:
	;
	if base.Ui32(int32(255)) < base.Ui32(v297) {
		goto L33
	} else {
		goto L78
	}
L56:
	;
	if base.Ui32(v297) <= base.Ui32(int32(_a_F_pg_regcomp_4)) {
		goto L75
	} else {
		goto L76
	}
L57:
	;
	if base.Ui32(int32(127)) < base.Ui32(v297) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	if base.Ui32(int32(127)) < base.Ui32(v297) {
		goto L33
	} else {
		goto L59
	}
L59:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+uint32(_c_F_pg_regcomp[1]))))
	v305 = int32(1)
	v391 = int32(base.Ui32(v304)>>(uint(v305)%32)) & v305
	goto L54
L60:
	;
	v391 = v353
	goto L54
L61:
	;
	v317 = int32(1178)
	v318 = int32(0)
	goto L64
L62:
	;
	goto L63
L63:
	;
	v343 = int32(1)
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297<<(uint(v343)%32))+uint32(_c_F_pg_regcomp[2]))))
	v353 = v345 & v343
	goto L60
L64:
	;
	v323 = base.I32_div_s(v317+v318, int32(2))
	v325 = v323 << (uint(int32(3)) % 32)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v325)+uint32(_c_F_pg_regcomp[3])))
	if base.Ui32(v328) < base.Ui32(v297) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v353 = int32(0)
	goto L60
L66:
	;
	if v340 <= v339 {
		v317 = v339
		v318 = v340
		goto L64
	} else {
		goto L73
	}
L67:
	;
	v339 = v317
	v340 = v323 + int32(1)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v325)+uint32(_c_F_pg_regcomp[4])))
	if base.Ui32(v334) <= base.Ui32(v297) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v353 = int32(1)
	goto L60
L71:
	;
	goto L72
L72:
	;
	v339 = v323 - int32(1)
	v340 = v318
	goto L66
L73:
	;
	goto L65
L74:
	;
	v391 = v377
	goto L54
L75:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v297)>>(uint(int32(8))%32)))+uint32(_c_F_pg_regcomp[5]))))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v297)>>(uint(int32(3))%32))&int32(31)|v365<<(uint(int32(5))%32))+uint32(_c_F_pg_regcomp[5]))))
	v377 = int32(base.Ui32(v369)>>(uint(v297&int32(7))%32)) & int32(1)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v377 = base.B2i32(base.Ui32(v297) < base.Ui32(int32(_a_F_pg_regcomp_5)))
	goto L74
L78:
	;
	goto L79
L79:
	;
	v391 = base.B2i32(base.B2i32(base.Ui32(v297|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L54
L80:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+8)) = v395 | int32(128)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v401 = v399 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v401
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	if base.Ui32(v403) <= base.Ui32(v401) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v587 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v568 + v587
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	if v590&v587 == int32(0) {
		goto L33
	} else {
		goto L138
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = int32(101)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	if v583 != 0 {
		goto L135
	} else {
		goto L136
	}
L83:
	;
	v407 = v401
	v411 = v403
	goto L84
L84:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_pg_regcomp[0]))
	switch v414 - int32(1) {
	case 0:
		goto L90
	case 1:
		goto L89
	case 2:
		goto L88
	default:
		goto L91
	}
L85:
	;
	if base.Ui32(v569) <= base.Ui32(v568) {
		goto L82
	} else {
		goto L133
	}
L86:
	;
	goto L85
L87:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v506 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L88:
	;
	if base.Ui32(int32(255)) < base.Ui32(v412) {
		v568 = v407
		v569 = v411
		goto L86
	} else {
		goto L111
	}
L89:
	;
	if base.Ui32(v412) <= base.Ui32(int32(_a_F_pg_regcomp_4)) {
		goto L108
	} else {
		goto L109
	}
L90:
	;
	if base.Ui32(int32(127)) < base.Ui32(v412) {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	if base.Ui32(int32(127)) < base.Ui32(v412) {
		v568 = v407
		v569 = v411
		goto L86
	} else {
		goto L92
	}
L92:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+uint32(_c_F_pg_regcomp[1]))))
	v420 = int32(1)
	v506 = int32(base.Ui32(v419)>>(uint(v420)%32)) & v420
	goto L87
L93:
	;
	v506 = v468
	goto L87
L94:
	;
	v432 = int32(1178)
	v433 = int32(0)
	goto L97
L95:
	;
	goto L96
L96:
	;
	v458 = int32(1)
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412<<(uint(v458)%32))+uint32(_c_F_pg_regcomp[2]))))
	v468 = v460 & v458
	goto L93
L97:
	;
	v438 = base.I32_div_s(v432+v433, int32(2))
	v440 = v438 << (uint(int32(3)) % 32)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v440)+uint32(_c_F_pg_regcomp[3])))
	if base.Ui32(v443) < base.Ui32(v412) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v468 = int32(0)
	goto L93
L99:
	;
	if v455 <= v454 {
		v432 = v454
		v433 = v455
		goto L97
	} else {
		goto L106
	}
L100:
	;
	v454 = v432
	v455 = v438 + int32(1)
	goto L99
L101:
	;
	goto L102
L102:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v440)+uint32(_c_F_pg_regcomp[4])))
	if base.Ui32(v449) <= base.Ui32(v412) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v468 = int32(1)
	goto L93
L104:
	;
	goto L105
L105:
	;
	v454 = v438 - int32(1)
	v455 = v433
	goto L99
L106:
	;
	goto L98
L107:
	;
	v506 = v492
	goto L87
L108:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v412)>>(uint(int32(8))%32)))+uint32(_c_F_pg_regcomp[5]))))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v412)>>(uint(int32(3))%32))&int32(31)|v480<<(uint(int32(5))%32))+uint32(_c_F_pg_regcomp[5]))))
	v492 = int32(base.Ui32(v484)>>(uint(v412&int32(7))%32)) & int32(1)
	goto L107
L109:
	;
	goto L110
L110:
	;
	v492 = base.B2i32(base.Ui32(v412) < base.Ui32(int32(_a_F_pg_regcomp_5)))
	goto L107
L111:
	;
	goto L112
L112:
	;
	v506 = base.B2i32(base.B2i32(base.Ui32(v412|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L87
L113:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	v568 = v507
	v569 = v510
	goto L86
L114:
	;
	goto L115
L115:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	switch v511 - int32(98) {
	case 0:
		goto L117
	case 1:
		goto L128
	default:
		goto L118
	case 3:
		goto L127
	case 7:
		goto L126
	case 11, 12:
		goto L125
	case 14:
		goto L124
	case 15:
		goto L123
	case 17:
		goto L122
	case 18:
		goto L121
	case 21:
		goto L120
	case 22:
		goto L119
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+16)) = v561
	v564 = v507 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v564
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	if base.Ui32(v564) < base.Ui32(v566) {
		v407 = v564
		v411 = v566
		goto L84
	} else {
		goto L132
	}
L117:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v558 & int32(-8)
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = int32(101)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	if v554 != 0 {
		goto L129
	} else {
		goto L130
	}
L119:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v549 | int32(32)
	goto L116
L120:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v544&int32(-193) | int32(128)
	goto L116
L121:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v541 & int32(-33)
	goto L116
L122:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v538 & int32(-193)
	goto L116
L123:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v533&int32(-8) | int32(4)
	goto L116
L124:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v528&int32(-193) | int32(64)
	goto L116
L125:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v525 | int32(192)
	goto L116
L126:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v522 | int32(8)
	goto L116
L127:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v517&int32(-8) | int32(1)
	goto L116
L128:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v561 = v514 & int32(-9)
	goto L116
L129:
	;
	v556 = v554
	goto L131
L130:
	;
	v556 = int32(18)
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = v556
	goto L33
L132:
	;
	v568 = v564
	v569 = v566
	goto L86
L133:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	if v571 == int32(41) {
		goto L81
	} else {
		goto L134
	}
L134:
	;
	goto L82
L135:
	;
	v585 = v583
	goto L137
L136:
	;
	v585 = int32(18)
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = v585
	goto L33
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+16)) = v590 & int32(-225)
	goto L33
L139:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	if v608&int32(4) != 0 {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	goto L141
L141:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v624&int32(192) != 0 {
		goto L150
	} else {
		goto L151
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = int32(110)
	v621 = F_next(m, v220)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L4
	} else {
		goto L149
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+32)) = int32(3)
	goto L142
L144:
	;
	goto L145
L145:
	;
	if v608&int32(1) != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+32)) = int32(1)
	goto L142
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+32)) = int32(2)
	goto L142
L149:
	;
	goto L141
L150:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	v629 = F_subcolor(m, v627, int32(10))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L4
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v636 != 0 {
		v815 = v636
		goto L10
	} else {
		goto L155
	}
L153:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+100)) = uint16(v629)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	F_okcolors(m, v632, v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v638 = v10 + int32(4)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v641)+8))
	v644 = F_parse(m, v638, int32(101), int32(112), v642, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v644
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v647 != 0 {
		v815 = v647
		goto L10
	} else {
		goto L157
	}
L157:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	F_specialcolors(m, v648)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v651 != 0 {
		v815 = v651
		goto L10
	} else {
		goto L159
	}
L159:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v652&int32(16) != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	F_removecaptures(m, v638, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L4
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v658 = int32(1)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v659)+4)) = v658
	v664 = int32(2)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v659)+20))
	if v665 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L162
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v673
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+1)))
	v678 = v676 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v675)+1)) = uint8(v678)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v675)+20))
	if v680 != 0 {
		goto L172
	} else {
		goto L173
	}
L165:
	;
	v667 = v665
	v668 = v664
	goto L168
L166:
	;
	v673 = v664
	goto L167
L167:
	;
	goto L164
L168:
	;
	v669 = F_numst(m, v667, v668)
	mBase = m.M
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v667)+24))
	if v670 != 0 {
		v667 = v670
		v668 = v669
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v673 = v669
	goto L167
L170:
	;
	goto L169
L171:
	;
	v686 = v10 + int32(4)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v686)+108))
	if v687 != 0 {
		goto L178
	} else {
		goto L179
	}
L172:
	;
	v681 = v680
	goto L175
L173:
	;
	goto L174
L174:
	;
	goto L171
L175:
	;
	F_markst(m, v681)
	mBase = m.M
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v681)+24))
	if v683 != 0 {
		v681 = v683
		goto L175
	} else {
		goto L177
	}
L176:
	;
	goto L174
L177:
	;
	goto L176
L178:
	;
	v689 = v687
	goto L181
L179:
	;
	goto L180
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v686)+108)) = int64(0)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	v715 = F_nfatree(m, v10+int32(4), v714)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L4
	} else {
		goto L188
	}
L181:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v689)+84))
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689)+1)))
	if v696&int32(64) == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L180
L183:
	;
	F_pfree(m, v689)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	if v695 != 0 {
		v689 = v695
		goto L181
	} else {
		goto L187
	}
L186:
	;
	goto L185
L187:
	;
	goto L182
L188:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v715 | v717
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v720 != 0 {
		v815 = v720
		goto L10
	} else {
		goto L189
	}
L189:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v10)+136))
	if int32(2) <= v721 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v730 = v658
	goto L193
L191:
	;
	goto L192
L192:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756)+1)))
	if v757&int32(2) != 0 {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	v736 = v733 + v730*int32(88)
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+2)))
	v742 = F_nfanode(m, v10+int32(4), v736, base.B2i32(v737&int32(2) == int32(0)))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L4
	} else {
		goto L195
	}
L194:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v748 != 0 {
		v815 = v748
		goto L10
	} else {
		goto L197
	}
L195:
	;
	v745 = v730 + int32(1)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v10)+136))
	if v745 < v746 {
		v730 = v745
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	goto L192
L198:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v760)+8)) = v761 | int32(_a_F_pg_regcomp_6)
	goto L200
L199:
	;
	goto L200
L200:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v767 = F_optimize(m, v766)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L4
	} else {
		goto L201
	}
L201:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v769 != 0 {
		v815 = v769
		goto L10
	} else {
		goto L202
	}
L202:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	F_makesearch(m, v10+int32(4), v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L4
	} else {
		goto L203
	}
L203:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v775 != 0 {
		v815 = v775
		goto L10
	} else {
		goto L204
	}
L204:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	F_compact(m, v776, v96+int32(20))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v781 != 0 {
		v815 = v781
		goto L10
	} else {
		goto L206
	}
L206:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v782
	v784 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v784
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(_a_F_pg_regcomp_7)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v789
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v791
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v793
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+16)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v784
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v10)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+68)) = v799
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v803&int32(8) != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v806 = int32(968)
	goto L209
L208:
	;
	v806 = int32(969)
	goto L209
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+420)) = v806
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+424)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = int32(0)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v10)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+428)) = v812
	v815 = v784
	goto L10
L210:
	;
	F_rfree(m, v823)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L4
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v822)+40))
	if v826 != v10+int32(52) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	goto L212
L214:
	;
	F_pfree(m, v826)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L4
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v822)+88))
	if v832 != 0 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	goto L216
L218:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v832)+36))
	if v833 != 0 {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v822)+104))
	if v902 != 0 {
		goto L236
	} else {
		goto L237
	}
L221:
	;
	v834 = v833
	goto L224
L222:
	;
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+36)) = int32(0)
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v832)+40))
	if v862 != 0 {
		goto L228
	} else {
		goto L229
	}
L224:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v832)+76))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v842)+136))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v842)+136)) = v843 + v844*int32(-36) - int32(8)
	F_pfree(m, v834)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L4
	} else {
		goto L226
	}
L225:
	;
	goto L223
L226:
	;
	if v841 != 0 {
		v834 = v841
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v863 = v862
	goto L231
L229:
	;
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v832)+40)) = int32(0)
	F_pfree(m, v832)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L4
	} else {
		goto L235
	}
L231:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v832)+76))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v871)+136))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v863)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v871)+136)) = v872 + v873*int32(-40) - int32(8)
	F_pfree(m, v863)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L4
	} else {
		goto L233
	}
L232:
	;
	goto L230
L233:
	;
	if v870 != 0 {
		v863 = v870
		goto L231
	} else {
		goto L234
	}
L234:
	;
	goto L232
L235:
	;
	goto L220
L236:
	;
	F_freesubre(m, v822, v902)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L4
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v822)+108))
	if v905 != 0 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	goto L238
L240:
	;
	v906 = v905
	goto L243
L241:
	;
	goto L242
L242:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v822)+120))
	if v930 != 0 {
		goto L250
	} else {
		goto L251
	}
L243:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v906)+84))
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906)+1)))
	if v914&int32(64) == int32(0) {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v822)+108)) = int64(0)
	goto L242
L245:
	;
	F_pfree(m, v906)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L4
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	if v913 != 0 {
		v906 = v913
		goto L243
	} else {
		goto L249
	}
L248:
	;
	goto L247
L249:
	;
	goto L244
L250:
	;
	F_pfree(m, v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L4
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v822)+124))
	if v933 != 0 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	goto L252
L254:
	;
	F_pfree(m, v933)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L4
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v822)+128))
	if v936 != 0 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	goto L256
L258:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v822)+132))
	v939 = v937 - int32(1)
	if int32(0) < v939 {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	goto L260
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+24)) = int32(101)
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v822)+12))
	if v985 != 0 {
		goto L274
	} else {
		goto L275
	}
L261:
	;
	v942 = v936
	v943 = v939
	goto L264
L262:
	;
	goto L263
L263:
	;
	F_pfree(m, v936)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L4
	} else {
		goto L273
	}
L264:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v942)+124))
	if v949 != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	goto L263
L266:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v942)+152))
	F_pfree(m, v950)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L4
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v963 = int32(1)
	if v963 < v943 {
		v942 = v942 + int32(88)
		v943 = v943 - v963
		goto L264
	} else {
		goto L272
	}
L269:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v942)+156))
	F_pfree(m, v953)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v942)+160))
	F_pfree(m, v956)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v942)+124)) = int32(0)
	goto L268
L272:
	;
	goto L265
L273:
	;
	goto L260
L274:
	;
	v986 = v985
	goto L276
L275:
	;
	v986 = v815
	goto L276
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+12)) = v986
	v994 = v986
	goto L3
}
func F_pg_relpagesbyid_v1_5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_relation_open(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_pg_relpages_impl(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_Int64GetDatum(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_pg_rewrite_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[0])))
	if v8 == int32(1) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[1])))
		F_elog_node_display(m, int32(_a_F_pg_rewrite_query_0), l0, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[2])))
			if v19 == int32(1) {
				v22 = int32(_a_F_pg_rewrite_query_1)
				*(*int64)(unsafe.Add(mBase, _c_F_pg_rewrite_query[3])) = int64(4)
				*(*int64)(unsafe.Add(mBase, _c_F_pg_rewrite_query[4])) = int64(3)
				*(*int64)(unsafe.Add(mBase, _c_F_pg_rewrite_query[5])) = int64(2)
				*(*int64)(unsafe.Add(mBase, _c_F_pg_rewrite_query[6])) = int64(1)
				v32 = F___syscall_ret(m, int32(0))
				mBase = m.M
				F_gettimeofday(m, int32(_a_F_pg_rewrite_query_2))
				mBase = m.M
			} else {
			}
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v35 == int32(6) {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = l0
				v43 = F_list_make1_impl(m, int32(1), v5+int32(8))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v47 = v43
					v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[2])))
					if v49 == int32(1) {
						F_ShowUsage(m, int32(_a_F_pg_rewrite_query_3))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[7])))
							if v56 == int32(1) {
								v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[1])))
								F_elog_node_display(m, int32(_a_F_pg_rewrite_query_4), v47, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									m.G0 = v5 + int32(16)
									return v47
								}
							} else {
								m.G0 = v5 + int32(16)
								return v47
							}
						}
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[7])))
						if v56 == int32(1) {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[1])))
							F_elog_node_display(m, int32(_a_F_pg_rewrite_query_4), v47, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(16)
								return v47
							}
						} else {
							m.G0 = v5 + int32(16)
							return v47
						}
					}
				}
			} else {
				v45 = F_QueryRewrite(m, l0)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = v45
					v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[2])))
					if v49 == int32(1) {
						F_ShowUsage(m, int32(_a_F_pg_rewrite_query_3))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[7])))
							if v56 == int32(1) {
								v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[1])))
								F_elog_node_display(m, int32(_a_F_pg_rewrite_query_4), v47, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									m.G0 = v5 + int32(16)
									return v47
								}
							} else {
								m.G0 = v5 + int32(16)
								return v47
							}
						}
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[7])))
						if v56 == int32(1) {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[1])))
							F_elog_node_display(m, int32(_a_F_pg_rewrite_query_4), v47, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(16)
								return v47
							}
						} else {
							m.G0 = v5 + int32(16)
							return v47
						}
					}
				}
			}
		}
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[2])))
		if v19 == int32(1) {
			v22 = int32(_a_F_pg_rewrite_query_1)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_rewrite_query[3])) = int64(4)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_rewrite_query[4])) = int64(3)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_rewrite_query[5])) = int64(2)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_rewrite_query[6])) = int64(1)
			v32 = F___syscall_ret(m, int32(0))
			mBase = m.M
			F_gettimeofday(m, int32(_a_F_pg_rewrite_query_2))
			mBase = m.M
		} else {
		}
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v35 == int32(6) {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = l0
			v43 = F_list_make1_impl(m, int32(1), v5+int32(8))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v47 = v43
				v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[2])))
				if v49 == int32(1) {
					F_ShowUsage(m, int32(_a_F_pg_rewrite_query_3))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[7])))
						if v56 == int32(1) {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[1])))
							F_elog_node_display(m, int32(_a_F_pg_rewrite_query_4), v47, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(16)
								return v47
							}
						} else {
							m.G0 = v5 + int32(16)
							return v47
						}
					}
				} else {
					v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[7])))
					if v56 == int32(1) {
						v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[1])))
						F_elog_node_display(m, int32(_a_F_pg_rewrite_query_4), v47, v61)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							m.G0 = v5 + int32(16)
							return v47
						}
					} else {
						m.G0 = v5 + int32(16)
						return v47
					}
				}
			}
		} else {
			v45 = F_QueryRewrite(m, l0)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = v45
				v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[2])))
				if v49 == int32(1) {
					F_ShowUsage(m, int32(_a_F_pg_rewrite_query_3))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[7])))
						if v56 == int32(1) {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[1])))
							F_elog_node_display(m, int32(_a_F_pg_rewrite_query_4), v47, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(16)
								return v47
							}
						} else {
							m.G0 = v5 + int32(16)
							return v47
						}
					}
				} else {
					v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[7])))
					if v56 == int32(1) {
						v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rewrite_query[1])))
						F_elog_node_display(m, int32(_a_F_pg_rewrite_query_4), v47, v61)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							m.G0 = v5 + int32(16)
							return v47
						}
					} else {
						m.G0 = v5 + int32(16)
						return v47
					}
				}
			}
		}
	}
}
func F_pg_role_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l2&int64(2199023255552) == int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v179
L2:
	;
	if l2&int64(512) != int64(0) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v14 = F_superuser_arg(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v14 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if l0 == l1 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v22 = F_roles_is_member_of(m, l1, int32(0), l0, v8+int32(12))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v24 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	if l0 == l1 {
		v179 = v4
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if l2&int64(256) != int64(0) {
		goto L31
	} else {
		goto L32
	}
L13:
	;
	v30 = F_superuser_arg(m, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v30 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(0)
	v35 = F_roles_is_member_of(m, l1, v32, v32, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v37 = int32(0)
	if v35 == v37 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v75 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L30
	}
L18:
	;
	v75 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v43 <= int32(0) {
		v69 = v37
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v75 = v69
	goto L17
L22:
	;
	v46 = int32(0)
	if v46 < v43 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v49 = v43
	goto L25
L24:
	;
	v49 = v46
	goto L25
L25:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v52 = int32(0)
	goto L26
L26:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50+v52<<(uint(int32(2))%32))))
	v61 = base.B2i32(v60 == l0)
	if v60 == l0 {
		v69 = v61
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v69 = v61
	goto L21
L28:
	;
	v63 = v52 + int32(1)
	if v63 != v49 {
		v52 = v63
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L12
L31:
	;
	if l0 == l1 {
		v179 = v4
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if l2&int64(4096) != int64(0) {
		goto L52
	} else {
		goto L53
	}
L34:
	;
	v81 = F_superuser_arg(m, l1)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v81 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v84 = int32(0)
	v86 = F_roles_is_member_of(m, l1, int32(1), v84, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v88 = int32(0)
	if v86 == v88 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v126 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L51
	}
L39:
	;
	v126 = int32(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v94 <= int32(0) {
		v120 = v88
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v126 = v120
	goto L38
L43:
	;
	v97 = int32(0)
	if v97 < v94 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v100 = v94
	goto L46
L45:
	;
	v100 = v97
	goto L46
L46:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v103 = int32(0)
	goto L47
L47:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
	v112 = base.B2i32(v111 == l0)
	if v111 == l0 {
		v120 = v112
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v120 = v112
	goto L42
L49:
	;
	v114 = v103 + int32(1)
	if v114 != v100 {
		v103 = v114
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L33
L52:
	;
	if l0 == l1 {
		v179 = v4
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v179 = int32(1)
	goto L1
L55:
	;
	v132 = F_superuser_arg(m, l1)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	if v132 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v135 = int32(0)
	v137 = F_roles_is_member_of(m, l1, int32(2), v135, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v139 = int32(0)
	if v137 == v139 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v177 != 0 {
		v179 = v4
		goto L1
	} else {
		goto L72
	}
L60:
	;
	v177 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v145 <= int32(0) {
		v171 = v139
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v177 = v171
	goto L59
L64:
	;
	v148 = int32(0)
	if v148 < v145 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v151 = v145
	goto L67
L66:
	;
	v151 = v148
	goto L67
L67:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v154 = int32(0)
	goto L68
L68:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v152+v154<<(uint(int32(2))%32))))
	v163 = base.B2i32(v162 == l0)
	if v162 == l0 {
		v171 = v163
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v171 = v163
	goto L63
L70:
	;
	v165 = v154 + int32(1)
	if v165 != v151 {
		v154 = v165
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L54
}
func F_pg_rusage_show(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	v12 = m.G0
	v14 = v12 - int32(192)
	m.G0 = v14
	v17 = v14 + int32(40)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = int64(3)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(1)
	v27 = F___syscall_ret(m, int32(0))
	mBase = m.M
	F_gettimeofday(m, v14+int32(24))
	mBase = m.M
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v31 < v32 {
		v35 = v31 + int32(_a_F_pg_rusage_show_0)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v35
		v37 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v37 - int64(1)
		v41 = v35
	} else {
		v41 = v31
	}
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v42 < v43 {
		v46 = v42 + int32(_a_F_pg_rusage_show_0)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v46
		v48 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
		*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v48 - int64(1)
		v52 = v46
	} else {
		v52 = v42
	}
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v54 < v55 {
		v58 = v54 + int32(_a_F_pg_rusage_show_0)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v58
		v61 = v53 - int64(1)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v61
		v63 = v58
		v64 = v61
	} else {
		v63 = v54
		v64 = v53
	}
	v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v69 = v67 - v68
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+16)) = uint32(v69)
	v72 = int32(_a_F_pg_rusage_show_1)
	v73 = base.I32_div_s(v41-v32, v72)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v73
	v75 = v64 - v66
	*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v75)
	v79 = base.I32_div_s(v63-v55, v72)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v79
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
	v82 = v81 - v65
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v82)
	v86 = base.I32_div_s(v52-v43, v72)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v86
	v91 = F_pg_snprintf(m, int32(_a_F_pg_rusage_show_2), int32(100), int32(_a_F_pg_rusage_show_3), v14)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		return int32(0)
	} else {
		m.G0 = v14 + int32(192)
		return int32(_a_F_pg_rusage_show_2)
	}
}
func F_pg_sjis_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v7 = int32(1)
	v10 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v19 = base.B2i32(int32(0) <= v10) | base.B2i32(base.Ui32((v10+int32(95))&int32(255)) < base.Ui32(int32(63)))
	if v19 != 0 {
		v20 = v7
	} else {
		v20 = int32(2)
	}
	v21 = base.B2i32(l1 < v20)
	if l1 < v20 {
		v22 = int32(-1)
	} else {
		v22 = v7
	}
	if v19|v21 != 0 {
		v53 = v22
	} else {
		if base.Ui32(int32(31)) <= base.Ui32((v10+int32(127))&int32(255)) {
			if base.Ui32(int32(28)) < base.Ui32((v10+int32(32))&int32(255)) {
				v53 = int32(-1)
			} else {
				v38 = int32(2)
				v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
				if base.Ui32((v41+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
					v48 = v38
				} else {
					v48 = int32(-1)
				}
				if v41 < int32(-3) {
					v51 = v38
				} else {
					v51 = v48
				}
				v53 = v51
			}
		} else {
			v38 = int32(2)
			v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32((v41+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
				v48 = v38
			} else {
				v48 = int32(-1)
			}
			if v41 < int32(-3) {
				v51 = v38
			} else {
				v51 = v48
			}
			v53 = v51
		}
	}
	return v53
}
func F_pg_sleep(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v59 float64
	_ = v59
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	F_gettimeofday(m, v11)
	mBase = m.M
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
	m.G0 = v11 + v10
	goto L1
L1:
	;
	goto L2
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sleep[0]))
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return int32(0)
L4:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v41 = m.G0
	v42 = int32(16)
	v43 = v41 - v42
	m.G0 = v43
	F_gettimeofday(m, v43)
	mBase = m.M
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+8)))
	m.G0 = v43 + v42
	goto L10
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	goto L3
L10:
	;
	v59 = base.F64_add(base.F64_add(v5, base.F64_div(base.F64_convert_i64_s(v15+v14*int64(1000000)-int64(946684800000000)), float64(1e+06))), base.F64_div(base.F64_convert_i64_s(v47+v46*int64(1000000)-int64(946684800000000)), float64(-1e+06)))
	if base.F64_ge(v59, float64(600)) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if base.F64_gt(v59, float64(0)) == int32(0) {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v72 = int32(_a_F_pg_sleep_0)
	goto L13
L13:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sleep[1]))
	v77 = F_WaitLatch(m, v74, int32(41), v72, int32(150994946))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	v72 = base.I32_trunc_sat_f64_s(base.F64_ceil(base.F64_mul(v59, float64(1000))))
	goto L13
L15:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sleep[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(0)
	goto L16
L16:
	;
	goto L2
}
func F_pg_strfromd(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v21 int64
	_ = v21
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10)+100)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+108)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = l0 + int32(31)
	v21 = base.I64_reinterpret_f64(l2)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v21&int64(9223372036854775807)) {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_pg_strfromd_0)
		v88 = int32(3)
		F_dostr(m, v10+int32(16), v88, v10+int32(88))
		mBase = m.M
		v96 = m.ExcPending
		if v96 != 0 {
			return
		} else {
			v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
			v101 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v101)
			m.G0 = v10 + int32(112)
			return
		}
	} else {
		if base.B2i32(base.B2i32(base.Ui64(v21) < base.Ui64(int64(9218868437227405313)))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v21&int64(9223372036854775807))) == int32(0))|base.F64_lt(l2, float64(0)) != 0 {
			v45 = base.F64_neg(l2)
			v46 = int32(45)
			v47 = int32(0)
		} else {
			v45 = l2
			v46 = v4
			v47 = int32(1)
		}
		if base.F64_eq(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_strfromd[0])))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v52)
			v55 = *(*int64)(unsafe.Add(mBase, _c_F_pg_strfromd[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v55
			v83 = int32(8)
			if v47 != 0 {
				v88 = v83
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = l0 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v46)
				v88 = v83
			}
			F_dostr(m, v10+int32(16), v88, v10+int32(88))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return
			} else {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
				v101 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v101)
				m.G0 = v10 + int32(112)
				return
			}
		} else {
			v58 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+84)) = uint8(v58)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = int32(1730817573)
			v63 = int32(1)
			if l1 <= v63 {
				v66 = v63
			} else {
				v66 = l1
			}
			if int32(32) <= v66 {
				v69 = int32(32)
			} else {
				v69 = v66
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v69
			*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v45
			v77 = F_snprintf(m, v10+int32(16), int32(64), v10+int32(80), v10)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				if int32(0) <= v77 {
					v83 = v77
					if v47 != 0 {
						v88 = v83
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = l0 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v46)
						v88 = v83
					}
					F_dostr(m, v10+int32(16), v88, v10+int32(88))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
						v101 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v101)
						m.G0 = v10 + int32(112)
						return
					}
				} else {
					v81 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v81)
					m.G0 = v10 + int32(112)
					return
				}
			}
		}
	}
}
func F_pg_strncoll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_pg_strong_random(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v3
	v15 = F_open(m, int32(_a_F_pg_strong_random_0), v3, v9)
	mBase = m.M
	if v15 != int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(1)
	if l1 == int32(0) {
		v41 = v18
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v48 = v3
	goto L3
L3:
	;
	m.G0 = v9 + int32(16)
	return v48
L4:
	;
	v43 = F_close(m, v15)
	mBase = m.M
	v48 = v41
	goto L3
L5:
	;
	v21 = l0
	v22 = l1
	goto L6
L6:
	;
	v27 = F_read(m, v15, v21, v22)
	mBase = m.M
	if v27 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v41 = v18
	goto L4
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strong_random[0]))
	if v31 == int32(27) {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v36 = v22 - v27
	if v36 != 0 {
		v21 = v21 + v27
		v22 = v36
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v41 = int32(0)
	goto L4
L12:
	;
	goto L7
}
func F_pg_switch_wal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_switch_wal[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_switch_wal[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_switch_wal[0])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_switch_wal_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_pg_switch_wal_1), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_switch_wal_2), int32(185), int32(_a_F_pg_switch_wal_3))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v38 = F_RequestXLogSwitch(m, int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = F_Int64GetDatum(m, v38)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				return v40
			}
		}
	}
}
func F_pg_tablespace_databases(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_InitMaterializedSRF(m, l0, int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	switch v13 - int32(1663) {
	case 0:
		v47 = int32(_a_F_pg_tablespace_databases_0)
		goto L5
	case 1:
		goto L7
	default:
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L39
	}
L4:
	;
	m.G0 = v10 + int32(80)
	return int32(0)
L5:
	;
	v48 = F_AllocateDir(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = int32(_a_F_pg_tablespace_databases_1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(_a_F_pg_tablespace_databases_2)
	v45 = F_psprintf(m, int32(_a_F_pg_tablespace_databases_3), v10+int32(48))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v24 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v24 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(_a_F_pg_tablespace_databases_4), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_pg_tablespace_databases_5), int32(237), int32(_a_F_pg_tablespace_databases_6))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v47 = v45
	goto L5
L13:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v50 = F_ReadDir(m, v48, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_pg_tablespace_databases[0]))
	if v103 != int32(44) {
		goto L3
	} else {
		goto L34
	}
L17:
	;
	if v50 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = v50
	goto L21
L19:
	;
	goto L20
L20:
	;
	F_FreeDir(m, v48)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L33
	}
L21:
	;
	v60 = v52 + int32(19)
	v64 = F_strtox_2(m, v60, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	v65 = base.I32_wrap_i64(v64)
	goto L24
L22:
	;
	goto L20
L23:
	;
	v91 = F_ReadDir(m, v48, v47)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	if v65 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v47
	v73 = F_psprintf(m, int32(_a_F_pg_tablespace_databases_7), v10+int32(32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v75 = F_directory_is_empty(m, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, v73)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v75 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+75)) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = v65
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v82, v83, v10+int32(76), v10+int32(75))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L23
L31:
	;
	if v91 != 0 {
		v52 = v91
		goto L21
	} else {
		goto L32
	}
L32:
	;
	goto L22
L33:
	;
	goto L4
L34:
	;
	v108 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v108 == int32(0) {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	F_errmsg(m, int32(_a_F_pg_tablespace_databases_8), v10)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_pg_tablespace_databases_5), int32(259), int32(_a_F_pg_tablespace_databases_6))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L4
L39:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v47
	F_errmsg(m, int32(_a_F_pg_tablespace_databases_9), v10+int32(16))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_pg_tablespace_databases_5), int32(257), int32(_a_F_pg_tablespace_databases_6))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_toupper(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	if base.Ui32((l0-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v10 = l0 - int32(32)
	} else {
		v10 = l0
	}
	return v10 & int32(255)
}
func F_pg_tz_acceptable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(946684800)
	v13 = F_localsub(m, l0+int32(256), v5+int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v19 = v17
		} else {
			v19 = int32(1)
		}
		m.G0 = v5 + int32(16)
		return base.B2i32(v19 == int32(0))
	}
}
func F_pg_ultostr_zeropad(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	if base.B2i32(l2 != int32(2))|base.B2i32(base.Ui32(int32(99)) < base.Ui32(l1)) == int32(0) {
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_c_F_pg_ultostr_zeropad[0]))))
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v14)
		return l0 + int32(2)
	} else {
		v19 = int32(0)
		if l1 == v19 {
			v28 = int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v28)
			v138 = int32(1)
		} else {
			v34 = int32(1233)
			v39 = int32(base.Ui32((base.I32_clz(l1)^int32(31))*v34+v34) >> (uint(int32(12)) % 32))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v39<<(uint(int32(2))%32))+uint32(_c_F_pg_ultostr_zeropad[1])))
			v44 = v39 + base.B2i32(base.Ui32(v42) <= base.Ui32(l1))
			if base.Ui32(int32(_a_F_pg_ultostr_zeropad_0)) <= base.Ui32(l1) {
				v48 = l1
				v50 = v19
				for {
					v57 = l0 + v44 - v50
					v58 = int32(4)
					v61 = base.I32_div_u_s(v48, int32(_a_F_pg_ultostr_zeropad_0))
					v64 = v48 + v61*int32(-10000)
					v65 = int32(100)
					v66 = base.I32_div_u_s(v64, v65)
					v67 = int32(1)
					v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66<<(uint(v67)%32))+uint32(_c_F_pg_ultostr_zeropad[0]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v57-v58))) = uint16(v69)
					v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v64-v66*v65)<<(uint(v67)%32))+uint32(_c_F_pg_ultostr_zeropad[0]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v57-int32(2)))) = uint16(v78)
					v81 = v50 + v58
					if base.Ui32(int32(99999999)) < base.Ui32(v48) {
						v48 = v61
						v50 = v81
						continue
					} else {
						break
					}
					break
				}
				v84 = v61
				v86 = v81
			} else {
				v84 = l1
				v86 = v19
			}
			if base.Ui32(int32(100)) <= base.Ui32(v84) {
				v97 = int32(2)
				v99 = int32(_a_F_pg_ultostr_zeropad_1)
				v101 = int32(100)
				v102 = base.I32_div_u_s(v84&v99, v101)
				v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v84-v102*v101)&v99<<(uint(int32(1))%32))+uint32(_c_F_pg_ultostr_zeropad[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(l0+v44-v86-v97))) = uint16(v110)
				v114 = v102
				v115 = v86 | v97
			} else {
				v114 = v84
				v115 = v86
			}
			if base.Ui32(int32(10)) <= base.Ui32(v114) {
				v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_c_F_pg_ultostr_zeropad[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(l0+v44-v115-int32(2)))) = uint16(v124)
				v138 = v44
			} else {
				v127 = v114 | int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v127)
				v138 = v44
			}
		}
		if l2 <= v138 {
			return l0 + v138
		} else {
			v142 = l0 + l2
			if v138 != 0 {
				base.MemoryCopy(m, v142-v138, l0, v138)
			} else {
			}
			v145 = l2 - v138
			if v145 != 0 {
				base.MemoryFill(m, l0, int32(48), v145)
			} else {
			}
			return v142
		}
	}
}
func F_pg_unicode_to_server(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	v1 = l0
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.Ui32(v1-int32(1)) < base.Ui32(int32(_a_F_pg_unicode_to_server_0)) {
		if base.Ui32(v1) <= base.Ui32(int32(127)) {
			v16 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v16)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v1)
			m.G0 = v8 + int32(16)
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server[0]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			if v21 == int32(6) {
				if base.Ui32(v1) <= base.Ui32(int32(2047)) {
					v29 = int32(base.Ui32(v1)>>(uint(int32(6))%32)) | int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v29)
					v67 = int32(1)
				} else {
					if base.Ui32(v1) <= base.Ui32(int32(_a_F_pg_unicode_to_server_1)) {
						v37 = int32(base.Ui32(v1)>>(uint(int32(12))%32)) | int32(224)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v37)
						v44 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&int32(63) | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v44)
						v67 = int32(2)
					} else {
						v50 = int32(base.Ui32(v1)>>(uint(int32(18))%32)) | int32(240)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v50)
						v54 = int32(63)
						v56 = int32(128)
						v57 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&v54 | v56
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v57)
						v64 = int32(base.Ui32(v1)>>(uint(int32(12))%32))&v54 | v56
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v64)
						v67 = int32(3)
					}
				}
				v72 = v1&int32(63) | int32(128)
				*(*uint8)(unsafe.Add(mBase, uint32(v67+l1))) = uint8(v72)
				v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
				if int32(0) <= v74 {
					v98 = int32(1)
				} else {
					v79 = v74 & int32(255)
					if v79&int32(224) == int32(192) {
						v98 = int32(2)
					} else {
						if v79&int32(240) == int32(224) {
							v98 = int32(3)
						} else {
							if v79&int32(248) == int32(240) {
								v96 = int32(4)
							} else {
								v96 = int32(1)
							}
							v98 = v96
						}
					}
				}
				v100 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v98+l1))) = uint8(v100)
				m.G0 = v8 + int32(16)
				return
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server[1]))
				if v103 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v219 = m.ExcPending
					if v219 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v222 = m.ExcPending
						if v222 != 0 {
							return
						} else {
							v224 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server[0]))
							v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
							v227 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v227
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v225
							F_errmsg(m, int32(_a_F_pg_unicode_to_server_2), v8)
							mBase = m.M
							v232 = m.ExcPending
							if v232 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_pg_unicode_to_server_3), int32(911), int32(_a_F_pg_unicode_to_server_4))
								mBase = m.M
								v237 = m.ExcPending
								if v237 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if base.Ui32(v1) <= base.Ui32(int32(2047)) {
						v111 = int32(base.Ui32(v1)>>(uint(int32(6))%32)) | int32(192)
						*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v111)
						v152 = v8 + int32(12)
					} else {
						if base.Ui32(v1) <= base.Ui32(int32(_a_F_pg_unicode_to_server_1)) {
							v120 = int32(base.Ui32(v1)>>(uint(int32(12))%32)) | int32(224)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v120)
							v127 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&int32(63) | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v127)
							v152 = v8 + int32(13)
						} else {
							v134 = int32(base.Ui32(v1)>>(uint(int32(18))%32)) | int32(240)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v134)
							v138 = int32(63)
							v140 = int32(128)
							v141 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&v138 | v140
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)) = uint8(v141)
							v148 = int32(base.Ui32(v1)>>(uint(int32(12))%32))&v138 | v140
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v148)
							v152 = v8 + int32(14)
						}
					}
					v156 = v1&int32(63) | int32(128)
					*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v156)
					v159 = v8 + int32(11)
					v160 = int32(*(*int8)(unsafe.Add(mBase, uint32(v159))))
					if int32(0) <= v160 {
						v184 = int32(1)
					} else {
						v165 = v160 & int32(255)
						if v165&int32(224) == int32(192) {
							v184 = int32(2)
						} else {
							if v165&int32(240) == int32(224) {
								v184 = int32(3)
							} else {
								if v165&int32(248) == int32(240) {
									v182 = int32(4)
								} else {
									v182 = int32(1)
								}
								v184 = v182
							}
						}
					}
					v186 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v184+v159))) = uint8(v186)
					v189 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server[1]))
					v192 = F_FunctionCall6Coll(m, v189, int32(6), v21, v159, l1, v184, v186)
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v203 = m.ExcPending
		if v203 != 0 {
			return
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v206 = m.ExcPending
			if v206 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_pg_unicode_to_server_5), int32(0))
				mBase = m.M
				v210 = m.ExcPending
				if v210 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_pg_unicode_to_server_3), int32(886), int32(_a_F_pg_unicode_to_server_4))
					mBase = m.M
					v215 = m.ExcPending
					if v215 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_pg_unicode_to_server_noerror(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	v1 = l0
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.Ui32(int32(_a_F_pg_unicode_to_server_noerror_0)) < base.Ui32(v1-int32(1)) {
		v200 = v3
		m.G0 = v8 + int32(16)
		return v200
	} else {
		if base.Ui32(v1) <= base.Ui32(int32(127)) {
			v16 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v16)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v1)
			v200 = int32(1)
			m.G0 = v8 + int32(16)
			return v200
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server_noerror[0]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			if v22 == int32(6) {
				if base.Ui32(v1) <= base.Ui32(int32(2047)) {
					v30 = int32(base.Ui32(v1)>>(uint(int32(6))%32)) | int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
					v68 = int32(1)
				} else {
					if base.Ui32(v1) <= base.Ui32(int32(_a_F_pg_unicode_to_server_noerror_1)) {
						v38 = int32(base.Ui32(v1)>>(uint(int32(12))%32)) | int32(224)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
						v45 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&int32(63) | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v45)
						v68 = int32(2)
					} else {
						v51 = int32(base.Ui32(v1)>>(uint(int32(18))%32)) | int32(240)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v51)
						v55 = int32(63)
						v57 = int32(128)
						v58 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&v55 | v57
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v58)
						v65 = int32(base.Ui32(v1)>>(uint(int32(12))%32))&v55 | v57
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v65)
						v68 = int32(3)
					}
				}
				v73 = v1&int32(63) | int32(128)
				*(*uint8)(unsafe.Add(mBase, uint32(v68+l1))) = uint8(v73)
				v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
				if int32(0) <= v75 {
					v99 = int32(1)
				} else {
					v80 = v75 & int32(255)
					if v80&int32(224) == int32(192) {
						v99 = int32(2)
					} else {
						if v80&int32(240) == int32(224) {
							v99 = int32(3)
						} else {
							if v80&int32(248) == int32(240) {
								v97 = int32(4)
							} else {
								v97 = int32(1)
							}
							v99 = v97
						}
					}
				}
				v101 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v99+l1))) = uint8(v101)
				v200 = int32(1)
				m.G0 = v8 + int32(16)
				return v200
			} else {
				v105 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server_noerror[1]))
				if v105 == int32(0) {
					v200 = v3
					m.G0 = v8 + int32(16)
					return v200
				} else {
					if base.Ui32(v1) <= base.Ui32(int32(2047)) {
						v113 = int32(base.Ui32(v1)>>(uint(int32(6))%32)) | int32(192)
						*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v113)
						v154 = v8 + int32(12)
					} else {
						if base.Ui32(v1) <= base.Ui32(int32(_a_F_pg_unicode_to_server_noerror_1)) {
							v122 = int32(base.Ui32(v1)>>(uint(int32(12))%32)) | int32(224)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v122)
							v129 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&int32(63) | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v129)
							v154 = v8 + int32(13)
						} else {
							v136 = int32(base.Ui32(v1)>>(uint(int32(18))%32)) | int32(240)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v136)
							v140 = int32(63)
							v142 = int32(128)
							v143 = int32(base.Ui32(v1)>>(uint(int32(6))%32))&v140 | v142
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)) = uint8(v143)
							v150 = int32(base.Ui32(v1)>>(uint(int32(12))%32))&v140 | v142
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v150)
							v154 = v8 + int32(14)
						}
					}
					v158 = v1&int32(63) | int32(128)
					*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v158)
					v161 = v8 + int32(11)
					v162 = int32(*(*int8)(unsafe.Add(mBase, uint32(v161))))
					if int32(0) <= v162 {
						v186 = int32(1)
					} else {
						v167 = v162 & int32(255)
						if v167&int32(224) == int32(192) {
							v186 = int32(2)
						} else {
							if v167&int32(240) == int32(224) {
								v186 = int32(3)
							} else {
								if v167&int32(248) == int32(240) {
									v184 = int32(4)
								} else {
									v184 = int32(1)
								}
								v186 = v184
							}
						}
					}
					v188 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v186+v161))) = uint8(v188)
					v191 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server_noerror[1]))
					v194 = F_FunctionCall6Coll(m, v191, int32(6), v22, v161, l1, v186, int32(1))
					mBase = m.M
					v197 = m.ExcPending
					if v197 != 0 {
						return int32(0)
					} else {
						v200 = base.B2i32(v194 == v186)
						m.G0 = v8 + int32(16)
						return v200
					}
				}
			}
		}
	}
}
func F_pg_visibility_map(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
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
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v2
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)) = uint16(v2)
	v19 = F_relation_open(m, v13, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+119)))
		v26 = v24 - int32(109)
		v33 = int32(0)
		if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v26))|base.B2i32(int32(1)<<(uint(v26)%32)&int32(161) == v33) == v33 {
			if base.Ui64(int64(4294967295)) <= base.Ui64(v12) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_visibility_map_0), int32(0))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_visibility_map_1), int32(103), int32(_a_F_pg_visibility_map_2))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v41 = F_CreateTemplateTupleDesc(m, int32(2))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v41, int32(1), int32(_a_F_pg_visibility_map_3), int32(16), int32(-1), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v41, int32(2), int32(_a_F_pg_visibility_map_4), int32(16), int32(-1), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v57 = F_BlessTupleDesc(m, v41)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v62 = F_visibilitymap_get_status(m, v19, base.I32_wrap_i64(v12), v9+int32(28))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									if v64 != 0 {
										F_ReleaseBuffer(m, v64)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v67 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v62 & v67
											*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(base.Ui32(v62)>>(uint(v67)%32)) & v67
											F_relation_close(m, v19, v67)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												v82 = F_heap_form_tuple(m, v57, v9+int32(20), v9+int32(18))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
													v85 = F_HeapTupleHeaderGetDatum(m, v84)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														m.G0 = v9 + int32(32)
														return v85
													}
												}
											}
										}
									} else {
										v67 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v62 & v67
										*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(base.Ui32(v62)>>(uint(v67)%32)) & v67
										F_relation_close(m, v19, v67)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v82 = F_heap_form_tuple(m, v57, v9+int32(20), v9+int32(18))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
												v85 = F_HeapTupleHeaderGetDatum(m, v84)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(32)
													return v85
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v98 + int32(4)
					F_errmsg(m, int32(_a_F_pg_visibility_map_5), v9)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
						v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105)+119)))
						F_errdetail_relkind_not_supported(m, v106)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_visibility_map_1), int32(951), int32(_a_F_pg_visibility_map_6))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_visibility_map_rel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(_a_F_pg_visibility_map_rel_0)
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_map_rel[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_map_rel[0])) = v23
			v26 = F_CreateTemplateTupleDesc(m, int32(3))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v26, int32(1), int32(_a_F_pg_visibility_map_rel_1), int32(20), int32(-1), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v26, int32(2), int32(_a_F_pg_visibility_map_rel_2), int32(16), int32(-1), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v26, int32(3), int32(_a_F_pg_visibility_map_rel_3), int32(16), int32(-1), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v49 = F_BlessTupleDesc(m, v26)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v49
								v53 = F_collect_visibility_data(m, v15, int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v53
									*(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_map_rel[0])) = v21
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
									if base.Ui32(v65) < base.Ui32(v66) {
										v68 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)) = uint8(v68)
										*(*uint16)(unsafe.Add(mBase, uint32(v9))) = uint16(v68)
										v73 = F_Int64GetDatum(m, base.I64_extend_i32_u(v65))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v73
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
											v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v76)+8)))
											v79 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v78 & v79
											*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(base.Ui32(v78)>>(uint(v79)%32)) & v79
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v76 + v79
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
											v93 = F_heap_form_tuple(m, v90, v9+int32(4), v9)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												v95 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
												*(*int64)(unsafe.Add(mBase, uint32(v63))) = v95 + int64(1)
												v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = int32(1)
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
												v103 = F_HeapTupleHeaderGetDatum(m, v102)
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return int32(0)
												} else {
													v115 = v103
													m.G0 = v9 + int32(16)
													return v115
												}
											}
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = int32(2)
											v110 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v110)
											v115 = int32(0)
											m.G0 = v9 + int32(16)
											return v115
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
		if base.Ui32(v65) < base.Ui32(v66) {
			v68 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)) = uint8(v68)
			*(*uint16)(unsafe.Add(mBase, uint32(v9))) = uint16(v68)
			v73 = F_Int64GetDatum(m, base.I64_extend_i32_u(v65))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v73
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
				v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v76)+8)))
				v79 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v78 & v79
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(base.Ui32(v78)>>(uint(v79)%32)) & v79
				*(*int32)(unsafe.Add(mBase, uint32(v64))) = v76 + v79
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
				v93 = F_heap_form_tuple(m, v90, v9+int32(4), v9)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					v95 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
					*(*int64)(unsafe.Add(mBase, uint32(v63))) = v95 + int64(1)
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = int32(1)
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
					v103 = F_HeapTupleHeaderGetDatum(m, v102)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						v115 = v103
						m.G0 = v9 + int32(16)
						return v115
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = int32(2)
				v110 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v110)
				v115 = int32(0)
				m.G0 = v9 + int32(16)
				return v115
			}
		}
	}
}
func F_pg_visible_in_snapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v12
	v19 = int32(1)
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui64(v12) < base.Ui64(v20) {
		v64 = v19
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(16)
	return v64
L4:
	;
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	if base.Ui64(v22) <= base.Ui64(v12) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v64 = int32(0)
	goto L3
L6:
	;
	v25 = v14 + int32(24)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if base.Ui32(v26) <= base.Ui32(int32(30)) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v40 = int32(0)
	goto L13
L8:
	;
	if v26 == int32(0) {
		v64 = v19
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v32 = int32(8)
	v36 = F_bsearch(m, v9+v32, v25, v26, v32, int32(1546))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v64 = base.B2i32(v36 == int32(0))
	goto L3
L13:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v25+v40<<(uint(int32(3))%32))))
	if v12 == v49 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v64 = v19
	goto L3
L15:
	;
	v52 = v40 + int32(1)
	if v26 != v52 {
		v40 = v52
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
}
func F_pg_wchar2euc_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v9)
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v74)
	return v73
L6:
	;
	v21 = int32(base.Ui32(v19) >> (uint(int32(24)) % 32))
	if v21 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v69 = v14
	v73 = v18
	goto L8
L8:
	;
	goto L5
L9:
	;
	v61 = int32(1)
	v65 = v59 + v18
	if v61 < v15 {
		v13 = v13 + int32(4)
		v14 = v60
		v15 = v15 - v61
		v18 = v65
		goto L4
	} else {
		goto L19
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v21)
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v23)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v27 = int32(base.Ui32(v25) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)) = uint8(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)) = uint8(v29)
	v31 = int32(4)
	v59 = v31
	v60 = v14 + v31
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = int32(base.Ui32(v19) >> (uint(int32(16)) % 32))
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v35)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v39 = int32(base.Ui32(v37) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v39)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)) = uint8(v41)
	v43 = int32(3)
	v59 = v43
	v60 = v14 + v43
	goto L9
L14:
	;
	goto L15
L15:
	;
	v47 = int32(base.Ui32(v19) >> (uint(int32(8)) % 32))
	if v47 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v47)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v49)
	v51 = int32(2)
	v59 = v51
	v60 = v14 + v51
	goto L9
L17:
	;
	goto L18
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v19)
	v55 = int32(1)
	v59 = v55
	v60 = v14 + v55
	goto L9
L19:
	;
	v69 = v60
	v73 = v65
	goto L8
}
func F_pg_wchar2mb_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wchar2mb_with_len[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6*int32(28))+uint32(_c_F_pg_wchar2mb_with_len[1])))
	v12 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_pg_wchar2mule_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v10)
	return v10
L2:
	;
	goto L3
L3:
	;
	v14 = l0
	v15 = l1
	v16 = l2
	v20 = v4
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v114)
	return v113
L6:
	;
	v23 = int32(base.Ui32(v21) >> (uint(int32(16)) % 32))
	v25 = v23 & int32(255)
	if base.Ui32(v25-int32(129)) <= base.Ui32(int32(12)) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v109 = v15
	v113 = v20
	goto L8
L8:
	;
	goto L5
L9:
	;
	v100 = v99 + v15
	v101 = int32(1)
	v105 = v99 + v20
	if v101 < v16 {
		v14 = v14 + int32(4)
		v15 = v100
		v16 = v16 - v101
		v20 = v105
		goto L4
	} else {
		goto L28
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v23)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v31)
	v99 = int32(2)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v25-int32(144)) <= base.Ui32(int32(9)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v23)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v41 = int32(base.Ui32(v39) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v41)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v43)
	v99 = int32(3)
	goto L9
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(v25-int32(160)) <= base.Ui32(int32(63)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v23)
	v51 = int32(154)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v53)
	v99 = int32(3)
	goto L9
L17:
	;
	goto L18
L18:
	;
	if v21&int32(15728640) == int32(14680064) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v23)
	v61 = int32(155)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v61)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v63)
	v99 = int32(3)
	goto L9
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(v25-int32(240)) <= base.Ui32(int32(4)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v23)
	v71 = int32(156)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v71)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v73) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)) = uint8(v77)
	v99 = int32(4)
	goto L9
L23:
	;
	goto L24
L24:
	;
	if base.B2i32(v25 == int32(255))|base.B2i32(base.Ui32(v25) < base.Ui32(int32(245))) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v23)
	v88 = int32(157)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v88)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v92 = int32(base.Ui32(v90) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v92)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)) = uint8(v94)
	v99 = int32(4)
	goto L9
L26:
	;
	goto L27
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v21)
	v99 = int32(1)
	goto L9
L28:
	;
	v109 = v100
	v113 = v105
	goto L8
}
