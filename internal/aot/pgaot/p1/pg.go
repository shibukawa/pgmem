package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_PGSemaphoreUnlock(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	Fn14220(m, l0, int32(_a_F_PGSemaphoreUnlock_0), int32(354), int32(_a_F_PGSemaphoreUnlock_1), int32(4))
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
	var v58 int32
	_ = v58
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
	*(*int64)(unsafe.Add(mBase, _c_F__PG_init_bloom[1])) = int64(17179869186)
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
	v55 = v34 * int32(12)
	v58 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F__PG_init_bloom[4]))) = v28<<(uint(v58)%32) + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F__PG_init_bloom[1]))) = v58
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
								F_RegisterShmemCallbacks(m, int32(_a_F__PG_init_pg_stat_statements_18))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									v62 = int32(_a_F__PG_init_pg_stat_statements_19)
									v63 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[3]))
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[4])) = v63
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[3])) = int32(_a_F__PG_init_pg_stat_statements_20)
									v68 = int32(_a_F__PG_init_pg_stat_statements_21)
									v69 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[5]))
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[5])) = int32(_a_F__PG_init_pg_stat_statements_22)
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[6])) = v69
									v75 = int32(_a_F__PG_init_pg_stat_statements_23)
									v76 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[7]))
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[7])) = int32(_a_F__PG_init_pg_stat_statements_24)
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[8])) = v76
									v82 = int32(_a_F__PG_init_pg_stat_statements_25)
									v83 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[9]))
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[9])) = int32(_a_F__PG_init_pg_stat_statements_26)
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[10])) = v83
									v89 = int32(_a_F__PG_init_pg_stat_statements_27)
									v90 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[11]))
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[11])) = int32(_a_F__PG_init_pg_stat_statements_28)
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[12])) = v90
									v96 = int32(_a_F__PG_init_pg_stat_statements_29)
									v97 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[13]))
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[13])) = int32(_a_F__PG_init_pg_stat_statements_30)
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[14])) = v97
									v103 = int32(_a_F__PG_init_pg_stat_statements_31)
									v104 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[15]))
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[15])) = int32(_a_F__PG_init_pg_stat_statements_32)
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_pg_stat_statements[16])) = v104
									return
								}
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 == int32(100) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L36
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L33
	}
L3:
	;
	v38 = F_text_to_cstring(m, base.I32_wrap_i64(v36))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L14
	}
L4:
	;
	v14 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_create_pg_locale_builtin[0])))
	v15 = F_SearchSysCache1(m, int32(21), v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v27 = F_SearchSysCache1(m, int32(16), base.I64_extend_i32_u(l0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L11
	}
L7:
	;
	return int32(0)
L8:
	;
	if v15 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v23 = F_SysCacheGetAttrNotNull(m, int32(21), v15, int32(15))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v35 = v15
	v36 = v23
	goto L3
L11:
	;
	if v27 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v33 = F_SysCacheGetAttrNotNull(m, int32(16), v27, int32(10))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v35 = v27
	v36 = v33
	goto L3
L14:
	;
	F_ReleaseCatCache(m, v35)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_builtin[1]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	goto L16
L16:
	;
	v45 = F_builtin_validate_locale(m, v44, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v48 = F_MemoryContextAllocZero(m, l1, int32(20))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v50 = F_MemoryContextStrdup(m, l1, v38)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v50
	v53 = int32(_a_F_create_pg_locale_builtin_0)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_builtin[2])))
	if base.B2i32(v56 == int32(0))|base.B2i32(v56 != v59) != 0 {
		v77 = v56
		v78 = v59
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v80 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+16)) = uint8(base.B2i32(v77-v78 == int32(0)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v85 != int32(67) {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	goto L20
L22:
	;
	v62 = v38
	v63 = v53
	goto L23
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v67
		v78 = v66
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v77 = v67
	v78 = v66
	goto L21
L25:
	;
	v70 = int32(1)
	if v67 == v66 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	m.G0 = v8 + int32(32)
	return v48
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(_a_F_create_pg_locale_builtin_1)
	goto L27
L29:
	;
	v88 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)) = uint8(v88)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v92 = base.B2i32(v90 == int32(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)) = uint8(v92)
	if v90 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_builtin[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v109
	F_errmsg_internal(m, int32(_a_F_create_pg_locale_builtin_2), v8)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_create_pg_locale_builtin_3), int32(250), int32(_a_F_create_pg_locale_builtin_4))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_create_pg_locale_builtin_5), v8+int32(16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_create_pg_locale_builtin_3), int32(263), int32(_a_F_create_pg_locale_builtin_4))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_pg_statistic_ext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v9 = int32(3)
	F_ScanKeyInit(m, v7, v9, v9, int32(62), base.I64_extend_i32_u(l2))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v7+int32(56), int32(4), int32(3), int32(184), base.I64_extend_i32_u(l1))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = int32(0)
			v30 = F_systable_beginscan(m, l0, int32(3997), int32(1), v25, int32(2), v7)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = F_systable_getnext(m, v30)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					if v32 == int32(0) {
						F_systable_endscan(m, v30)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v52 = v25
							m.G0 = v7 + int32(112)
							return v52
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v38+v39)))
						F_systable_endscan(m, v30)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							if v41 == int32(0) {
								v52 = v25
								m.G0 = v7 + int32(112)
								return v52
							} else {
								v49 = F_SearchSysCacheCopy(m, int32(64), base.I64_extend_i32_u(v41), int64(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v52 = v49
									m.G0 = v7 + int32(112)
									return v52
								}
							}
						}
					}
				}
			}
		}
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
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
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
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = int32(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = int32(42)
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
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(0) < v163 {
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
	v144 = int32(0)
	v146 = int32(_a_F_pg_analyze_and_rewrite_varparams_3)
	v147 = base.AtomicRmwOr32(m, v144, v146, v144)
	*(*int64)(unsafe.Add(mBase, uint32(v119)+392)) = v115
	v152 = base.AtomicRmwOr32(m, v144, v146, v144)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v153 + v137
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[10])) = v159 - v137
	goto L35
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L52
	}
L40:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v171 = int32(0)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_varparams[0])))
	if v199 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v166+v171<<(uint(int32(2))%32))))
	if base.B2i32(v180 == int32(705))|base.B2i32(v180 == int32(0)) != 0 {
		goto L39
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	v187 = v171 + int32(1)
	if v187 != v163 {
		v171 = v187
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	F_ShowUsage(m, int32(_a_F_pg_analyze_and_rewrite_varparams_4))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v203 = F_pg_rewrite_query(m, v92)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
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
	return v203
L52:
	;
	F_errcode(m, int32(134611076))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v171 + int32(1)
	F_errmsg(m, int32(_a_F_pg_analyze_and_rewrite_varparams_5), v12)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_pg_analyze_and_rewrite_varparams_6), int32(860), int32(_a_F_pg_analyze_and_rewrite_varparams_7))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
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
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
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
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = l4
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
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(42)
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
	v149 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[0])))
	if v149 == int32(1) {
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
	v129 = int32(0)
	v131 = int32(_a_F_pg_analyze_and_rewrite_withcb_3)
	v132 = base.AtomicRmwOr32(m, v129, v131, v129)
	*(*int64)(unsafe.Add(mBase, uint32(v104)+392)) = v100
	v137 = base.AtomicRmwOr32(m, v129, v131, v129)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v138 + v122
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_withcb[9])) = v144 - v122
	goto L34
L38:
	;
	F_ShowUsage(m, int32(_a_F_pg_analyze_and_rewrite_withcb_4))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v155 = F_pg_rewrite_query(m, v79)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	return v155
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
func F_pg_available_extension_versions(m *base.Module, l0 int32) int64 {
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
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
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
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
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
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int64
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v301 int64
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v345 int32
	_ = v345
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int64
	_ = v564
	var v566 int64
	_ = v566
	var v568 int64
	_ = v568
	var v570 int64
	_ = v570
	var v572 int64
	_ = v572
	var v574 int64
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int64
	_ = v583
	var v585 int64
	_ = v585
	var v587 int64
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v630 int32
	_ = v630
	var v634 int64
	_ = v634
	var v635 int64
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v745 int32
	_ = v745
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(96)
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
	return int64(0)
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
	m.G0 = v824 + int32(96)
	return int64(0)
L4:
	;
	if v33 == int32(0) {
		v824 = v25
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v37 <= int32(0) {
		v824 = v25
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v42 = v25
	v48 = v33
	v55 = v27
	v56 = v2
	v57 = v2
	goto L7
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v56<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v68 = F_AllocateDir(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v824 = v798
	goto L3
L9:
	;
	v819 = v812 + int32(1)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	if v819 < v820 {
		v42 = v798
		v48 = v804
		v55 = v811
		v56 = v819
		v57 = v813
		goto L7
	} else {
		goto L146
	}
L10:
	;
	if v68 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_pg_available_extension_versions[0]))
	if v73 == int32(44) {
		v798 = v42
		v804 = v48
		v811 = v55
		v812 = v56
		v813 = v57
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v77 = F_ReadDir(m, v68, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	if v77 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v79 = v77
	v81 = v42
	v86 = v66
	v87 = v48
	v88 = v68
	v94 = v55
	v95 = v56
	v96 = v57
	goto L19
L17:
	;
	v774 = v42
	v780 = v48
	v781 = v68
	v787 = v55
	v788 = v56
	v789 = v57
	goto L18
L18:
	;
	F_FreeDir(m, v781)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L145
	}
L19:
	;
	v102 = v79 + int32(19)
	v106 = F_strlen(m, v102)
	mBase = m.M
	v113 = v106 + int32(1)
	goto L24
L20:
	;
	v774 = v81
	v780 = v87
	v781 = v88
	v787 = v94
	v788 = v95
	v789 = v764
	goto L18
L21:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v770 = F_ReadDir(m, v88, v769)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L143
	}
L22:
	;
	if v125 == int32(0) {
		v764 = v96
		goto L21
	} else {
		goto L28
	}
L23:
	;
	goto L22
L24:
	;
	v115 = int32(0)
	if v113 == v115 {
		v125 = v115
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v125 = v120
	goto L23
L26:
	;
	v119 = v113 - int32(1)
	v120 = v102 + v119
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v121 != int32(46) {
		v113 = v119
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v128 = int32(_a_F_pg_available_extension_versions_0)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_available_extension_versions[1])))
	if base.B2i32(v131 == int32(0))|base.B2i32(v131 != v134) != 0 {
		v152 = v131
		v153 = v134
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v152-v153 != 0 {
		v764 = v96
		goto L21
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v137 = v125
	v138 = v128
	goto L32
L32:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	if v142 == int32(0) {
		v152 = v142
		v153 = v141
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v152 = v142
	v153 = v141
	goto L30
L34:
	;
	v145 = int32(1)
	if v142 == v141 {
		v137 = v137 + v145
		v138 = v138 + v145
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v155 = F_pstrdup(m, v102)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v160 = F_strlen(m, v155)
	mBase = m.M
	v167 = v160 + int32(1)
	goto L40
L38:
	;
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v180)
	v183 = F_strstr(m, v155, int32(_a_F_pg_available_extension_versions_1))
	mBase = m.M
	if v183 != 0 {
		v764 = v96
		goto L21
	} else {
		goto L44
	}
L39:
	;
	goto L38
L40:
	;
	v169 = int32(0)
	if v167 == v169 {
		v179 = v169
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v179 = v174
	goto L39
L42:
	;
	v173 = v167 - int32(1)
	v174 = v155 + v173
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v175 != int32(46) {
		v167 = v173
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v184 = F_makeString(m, v155)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v186 = F_list_member(m, v96, v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v186 != 0 {
		v764 = v96
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v188 = F_lappend(m, v96, v184)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v191 = F_palloc0(m, int32(48))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v193 = F_pstrdup(m, v155)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+36)) = int32(-1)
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+34)) = uint8(v197)
	v199 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v191)+32)) = uint16(v199)
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v193
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v203 = F_pstrdup(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v203
	F_parse_extension_control_file(m, v191, int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v94)+28))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
	v211 = F_get_ext_ver_list(m, v191)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v211 == int32(0) {
		v764 = v188
		goto L21
	} else {
		goto L54
	}
L54:
	;
	v215 = int32(0)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v216 <= v215 {
		v764 = v188
		goto L21
	} else {
		goto L55
	}
L55:
	;
	v223 = v216
	v229 = v215
	goto L56
L56:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v241+v229<<(uint(int32(2))%32))))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+8)))
	if v246 != int32(1) {
		v726 = v223
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v764 = v188
	goto L21
L58:
	;
	v745 = v229 + int32(1)
	if v745 < v726 {
		v223 = v726
		v229 = v745
		goto L56
	} else {
		goto L142
	}
L59:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v251 = F_palloc(m, int32(48))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v191)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v251)+40)) = v253
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v191)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v251)+32)) = v255
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v191)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v251)+24)) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v191)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v251)+16)) = v259
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v191)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v251)+8)) = v261
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v191)))
	*(*int64)(unsafe.Add(mBase, uint32(v251))) = v263
	F_parse_extension_control_file(m, v251, v249)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v269 = int32(0)
	base.MemoryFill(m, v81+int32(16), v269, int32(72))
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+8)) = uint8(v269)
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = int64(0)
	v278 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v251))))
	v279 = F_DirectFunctionCall1Coll(m, int32(534), v269, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+16)) = v279
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v283 = F_cstring_to_text(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = base.I64_extend_i32_u(v283)
	v287 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v251)+33)))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+32)) = v287
	v289 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v251)+34)))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+40)) = v289
	v291 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v251)+32)))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+48)) = v291
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	if v293 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v251)+40))
	if v304 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L65:
	;
	v296 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+5)) = uint8(v296)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v301 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), base.I64_extend_i32_u(v293))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+56)) = v301
	goto L64
L69:
	;
	v407 = F_superuser(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L83
	}
L70:
	;
	v307 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+6)) = uint8(v307)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	v312 = F_palloc(m, v309<<(uint(int32(3))%32))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v314 = int32(0)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if v314 < v315 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v318 = v314
	goto L77
L75:
	;
	v357 = v314
	goto L76
L76:
	;
	v380 = F_construct_array_builtin(m, v312, v357, int32(19))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L81
	}
L77:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	v349 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v345+v318<<(uint(int32(2))%32)))))
	v350 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	v357 = v354
	goto L76
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v312+v318<<(uint(int32(3))%32)))) = v350
	v354 = v318 + int32(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if v354 < v355 {
		v318 = v354
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+64)) = base.I64_extend_i32_u(v380)
	goto L69
L82:
	;
	v414 = F_cstring_to_text(m, v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L86
	}
L83:
	;
	if v407 == int32(0) {
		v413 = int32(_a_F_pg_available_extension_versions_2)
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v411 != 0 {
		v413 = v411
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v413 = v412
	goto L82
L86:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+72)) = base.I64_extend_i32_u(v414)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v251)+24))
	if v418 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	F_tuplestore_putvalues(m, v210, v209, v81+int32(16), v81)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L92
	}
L88:
	;
	v421 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+8)) = uint8(v421)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v423 = F_cstring_to_text(m, v418)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+80)) = base.I64_extend_i32_u(v423)
	goto L87
L92:
	;
	v431 = int32(0)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v432 <= v431 {
		v726 = v432
		goto L58
	} else {
		goto L93
	}
L93:
	;
	v439 = v432
	v447 = v431
	goto L94
L94:
	;
	if v439 <= int32(0) {
		v701 = v439
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v726 = v701
	goto L58
L96:
	;
	v720 = v447 + int32(1)
	if v720 < v701 {
		v439 = v701
		v447 = v720
		goto L94
	} else {
		goto L141
	}
L97:
	;
	v459 = int32(0)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v462+v447<<(uint(int32(2))%32))))
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466)+8)))
	if v467&int32(1) != 0 {
		v701 = v439
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v470 = v459
	v471 = v459
	v475 = v459
	goto L99
L99:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v492+v470<<(uint(int32(2))%32))))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+8)))
	if v497 != int32(1) {
		v550 = v471
		v551 = v475
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v550 != v245 {
		v701 = v557
		goto L96
	} else {
		goto L123
	}
L101:
	;
	v556 = v470 + int32(1)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v556 < v557 {
		v470 = v556
		v471 = v550
		v475 = v551
		goto L99
	} else {
		goto L122
	}
L102:
	;
	v500 = int32(1)
	v502 = F_find_update_path(m, v211, v496, v466, v500, v500)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v502 == int32(0) {
		v550 = v471
		v551 = v475
		goto L101
	} else {
		goto L104
	}
L104:
	;
	if v471 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v550 = v496
	v551 = v502
	goto L101
L106:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	if v475 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v508 != v516 {
		v550 = v471
		v551 = v475
		goto L101
	} else {
		goto L113
	}
L108:
	;
	v511 = int32(0)
	if v511 <= v508 {
		v516 = v511
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	if v508 < v514 {
		goto L105
	} else {
		goto L112
	}
L111:
	;
	goto L105
L112:
	;
	v516 = v514
	goto L107
L113:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	if base.B2i32(v522 == int32(0))|base.B2i32(v522 != v525) != 0 {
		v543 = v522
		v544 = v525
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if int32(0) <= v543-v544 {
		v550 = v471
		v551 = v475
		goto L101
	} else {
		goto L121
	}
L115:
	;
	goto L114
L116:
	;
	v528 = v518
	v529 = v519
	goto L117
L117:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+1)))
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+1)))
	if v533 == int32(0) {
		v543 = v533
		v544 = v532
		goto L115
	} else {
		goto L119
	}
L118:
	;
	v543 = v533
	v544 = v532
	goto L115
L119:
	;
	v536 = int32(1)
	if v533 == v532 {
		v528 = v528 + v536
		v529 = v529 + v536
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	goto L105
L122:
	;
	goto L100
L123:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	v562 = F_palloc(m, int32(48))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v191)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v562)+40)) = v564
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v191)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v562)+32)) = v566
	v568 = *(*int64)(unsafe.Add(mBase, uint32(v191)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v562)+24)) = v568
	v570 = *(*int64)(unsafe.Add(mBase, uint32(v191)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v562)+16)) = v570
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v191)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v562)+8)) = v572
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v191)))
	*(*int64)(unsafe.Add(mBase, uint32(v562))) = v574
	F_parse_extension_control_file(m, v562, v560)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	v579 = F_cstring_to_text(m, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = base.I64_extend_i32_u(v579)
	v583 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v562)+33)))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+32)) = v583
	v585 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v562)+34)))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+40)) = v585
	v587 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v562)+32)))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+48)) = v587
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v562)+40))
	if v589 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+6)) = uint8(v670)
	F_tuplestore_putvalues(m, v210, v209, v81+int32(16), v81)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L140
	}
L128:
	;
	v670 = int32(1)
	goto L127
L129:
	;
	goto L130
L130:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	v596 = F_palloc(m, v593<<(uint(int32(3))%32))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v598 = int32(0)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	if v598 < v600 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v603 = v598
	goto L135
L133:
	;
	v642 = v598
	goto L134
L134:
	;
	v665 = F_construct_array_builtin(m, v596, v642, int32(19))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L139
	}
L135:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	v634 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v630+v603<<(uint(int32(2))%32)))))
	v635 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), v634)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L137
	}
L136:
	;
	v642 = v639
	goto L134
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v596+v603<<(uint(int32(3))%32)))) = v635
	v639 = v603 + int32(1)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	if v639 < v640 {
		v603 = v639
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+64)) = base.I64_extend_i32_u(v665)
	v670 = v598
	goto L127
L140:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v701 = v696
	goto L96
L141:
	;
	goto L95
L142:
	;
	goto L57
L143:
	;
	if v770 != 0 {
		v79 = v770
		v96 = v764
		goto L19
	} else {
		goto L144
	}
L144:
	;
	goto L20
L145:
	;
	v798 = v774
	v804 = v780
	v811 = v787
	v812 = v788
	v813 = v789
	goto L9
L146:
	;
	goto L8
}
func F_pg_b64_dec_len(m *base.Module, l0 int32) int32 {
	return l0 * int32(3) >> (uint(int32(2)) % 32)
}
func F_pg_base64url_decode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pg_base64_decode_internal(m, l0, l1, l2, int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_pg_check_visible(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14344(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_pg_checksum_page(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = int32(0)
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_checksum_page[0]))
	v10 = m.T0[v9].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v5)
		v16 = int32(_a_F_pg_checksum_page_0)
		v17 = base.I32_rem_u_s(l1^v10, v16)
		return (v17 + int32(1)) & v16
	}
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
func F_pg_collation_for(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_get_fn_expr_argtype(m, v9, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v39 = int64(0)
			m.G0 = v7 + int32(16)
			return v39
		} else {
			v20 = F_type_is_collatable(m, v11)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				if base.B2i32(v20 == int32(0))&base.B2i32(v11 != int32(705)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							v51 = F_format_type_be(m, v11)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v51
								F_errmsg(m, int32(_a_F_pg_collation_for_0), v7)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_collation_for_1), int32(604), int32(_a_F_pg_collation_for_2))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int64(0)
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
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v27 == int32(0) {
						v30 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
						v39 = int64(0)
						m.G0 = v7 + int32(16)
						return v39
					} else {
						v33 = F_generate_collation_name(m, v27)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = F_cstring_to_text(m, v33)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
							} else {
								v39 = base.I64_extend_i32_u(v35)
								m.G0 = v7 + int32(16)
								return v39
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_conf_load_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_pg_conf_load_time[0]))
	return v3
}
func F_pg_control_recovery(m *base.Module, l0 int32) int64 {
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
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v11 = F_get_call_result_type(m, l0, int32(0), v4+int32(-60))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		if v11 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_recovery[0]))
			v22 = F_LWLockAcquire(m, v18+int32(1152), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_recovery[1]))
				v28 = F_get_controlfile(m, v25, v4+int32(-61))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_recovery[0]))
					F_LWLockRelease(m, v31+int32(1152))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
						if v36 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_pg_control_recovery_0), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_control_recovery_1), int32(184), int32(_a_F_pg_control_recovery_2))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v28)+144))
							v40 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v39
							v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+152)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v43
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v47
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v28)+168))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v51
							v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v28)+176)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+48)) = v55
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
							v64 = F_heap_form_tuple(m, v59, v4+int32(-48), v4+int32(-53))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
								v67 = F_HeapTupleHeaderGetDatum(m, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int64(0)
								} else {
									m.G0 = v6 - int32(-64)
									return v67
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_control_recovery_3), int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_control_recovery_1), int32(176), int32(_a_F_pg_control_recovery_2))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int64(0)
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
func F_pg_convert(m *base.Module, l0 int32) int64 {
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
	var v283 int32
	_ = v283
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
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L97
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
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
	v283 = m.ExcPending
	if v283 != 0 {
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
	return base.I64_extend_i32_u(v273)
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
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v18
	F_errmsg(m, int32(_a_F_pg_convert_2), v11)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_pg_convert_3), int32(580), int32(_a_F_pg_convert_4))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
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
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v111
	F_errmsg(m, int32(_a_F_pg_convert_5), v11+int32(16))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_pg_convert_3), int32(585), int32(_a_F_pg_convert_4))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
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
func F_pg_database_collation_actual_version(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v12 = F_SearchSysCache1(m, int32(21), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16+v17)+76)))
			if v19 == int32(99) {
				v25 = int32(13)
			} else {
				v25 = int32(15)
			}
			v26 = F_SysCacheGetAttrNotNull(m, int32(21), v12, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				v29 = F_text_to_cstring(m, base.I32_wrap_i64(v26))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = F_get_collation_actual_version(m, v19, v29)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						F_ReleaseCatCache(m, v12)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							if v31 != 0 {
								v35 = F_cstring_to_text(m, v31)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int64(0)
								} else {
									v41 = base.I64_extend_i32_u(v35)
									m.G0 = v8 + int32(16)
									return v41
								}
							} else {
								v38 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
								v41 = int64(0)
								m.G0 = v8 + int32(16)
								return v41
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v11)
					F_errmsg(m, int32(_a_F_pg_database_collation_actual_version_0), v8)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_database_collation_actual_version_1), int32(2807), int32(_a_F_pg_database_collation_actual_version_2))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
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
func F_pg_database_size_name(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_get_database_oid(m, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = F_calculate_database_size(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			if v9 == int64(0) {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			} else {
			}
			return v9
		}
	}
}
func F_pg_ddl_command_send(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_pg_ddl_command_send_0), int32(359), int32(_a_F_pg_ddl_command_send_1), int32(_a_F_pg_ddl_command_send_2), int32(_a_F_pg_ddl_command_send_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
func F_pg_encrypt(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	F_pfree(m, v137)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L93
	}
L2:
	;
	return int64(0)
L3:
	;
	v19 = int32(1)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v23 = v21 & v19
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v19
	goto L6
L5:
	;
	v24 = int32(4)
	goto L6
L6:
	;
	if v21 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v54 = F_downcase_truncate_identifier(m, v15+v24, v52, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v31 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v42 = int32(1)
	if v23 != 0 {
		v52 = int32(base.Ui32(v21)>>(uint(v42)%32)) - v42
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v34 = int32(16)
	goto L13
L12:
	;
	v34 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v41 = int32(4)
	goto L16
L15:
	;
	v41 = v34
	goto L16
L16:
	;
	v52 = v41
	goto L7
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v58 = F_px_find_combo(m, v54, v12+int32(28))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v58 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v54)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
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
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v66 = F_pg_detoast_datum_packed(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v69 = F_pg_detoast_datum_packed(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v101 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v77 == int32(18) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v88 = int32(1)
	if v71&v88 != 0 {
		v100 = int32(base.Ui32(v71)>>(uint(v88)%32)) - v88
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v80 = int32(16)
	goto L32
L31:
	;
	v80 = int32(0)
	goto L32
L32:
	;
	if base.Ui32((v77-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v87 = int32(4)
	goto L35
L34:
	;
	v87 = v80
	goto L35
L35:
	;
	v100 = v87
	goto L26
L36:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v100 = int32(base.Ui32(v94)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v132 = m.T0[v131].(func(*base.Module, int32, int32) int32)(m, v64, v100)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L48
	}
L38:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v107 == int32(18) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v118 = int32(1)
	if v101&v118 != 0 {
		v130 = int32(base.Ui32(v101)>>(uint(v118)%32)) - v118
		goto L37
	} else {
		goto L47
	}
L41:
	;
	v110 = int32(16)
	goto L43
L42:
	;
	v110 = int32(0)
	goto L43
L43:
	;
	if base.Ui32((v107-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v117 = int32(4)
	goto L46
L45:
	;
	v117 = v110
	goto L46
L46:
	;
	v130 = v117
	goto L37
L47:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v130 = int32(base.Ui32(v124)>>(uint(int32(2))%32)) - int32(4)
	goto L37
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v132
	v137 = F_palloc(m, v132+int32(4))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v139 = int32(1)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v141&v139 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v144 = v139
	goto L52
L51:
	;
	v144 = int32(4)
	goto L52
L52:
	;
	v146 = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v149 = m.T0[v148].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v64, v69+v144, v130, v146, v146)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	if v149 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v153 = int32(1)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v155&v153 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v167 = v149
	goto L56
L56:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	m.T0[v168].(func(*base.Module, int32))(m, v64)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L61
	}
L57:
	;
	v158 = v153
	goto L59
L58:
	;
	v158 = int32(4)
	goto L59
L59:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v165 = m.T0[v164].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v64, v66+v158, v100, v137+int32(4), v12+int32(28))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v167 = v165
	goto L56
L61:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v171 != v66 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_pfree(m, v66)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v175 != v69 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	F_pfree(m, v69)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v179 != v15 {
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
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v167 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v183<<(uint(int32(2))%32) + int32(16)
	m.G0 = v12 + int32(32)
	return base.I64_extend_i32_u(v137)
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
	if v58 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v54
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
	if v58 != v212 {
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
	if v58 != v218 {
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
	if v167 == int32(0) {
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
	if v167 != v263 {
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
	if v167 != v269 {
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
func F_pg_extension_config_dump(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int64
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	v13 = m.G0
	v15 = v13 - int32(224)
	m.G0 = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_extension_config_dump[0])))
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L85
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L82
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L79
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L76
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L73
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L69
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L65
	}
L10:
	;
	v25 = base.I32_wrap_i64(v17)
	v26 = F_get_rel_name(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L61
	}
L13:
	;
	if v26 == int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v31 = F_getExtensionOfObject(m, int32(1259), v25)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_extension_config_dump[1]))
	if v31 != v34 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v38 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v41 = v15 + int32(160)
	v46 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_pg_extension_config_dump[1])))
	F_ScanKeyInit(m, v41, int32(1), int32(3), int32(184), v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v50 = int32(1)
	v53 = F_systable_beginscan(m, v38, int32(3080), v50, int32(0), v50, v41)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v55 = F_systable_getnext(m, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v55 == int32(0) {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v59 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+136)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = int64(281474976710656)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = v17 & int64(4294967295)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v86 = F_heap_getattr_7(m, v55, int32(7), v83, v15+int32(147))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+147)))
	if v88 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = base.I64_extend_i32_u(v19)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = base.I64_extend_i32_u(v175)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v184 = F_heap_getattr_7(m, v55, int32(8), v181, v15+int32(147))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L43
	}
L24:
	;
	v91 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+148)) = v91
	v97 = F_construct_array_builtin(m, v15+int32(152), v91, int32(26))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v100 = F_pg_detoast_datum(m, base.I32_wrap_i64(v86))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v167 = int32(0)
	v175 = v97
	goto L23
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v102 != int32(1) {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	if v105 != int32(1) {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	if v108 < int32(0) {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v111 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	if v112 != int32(26) {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+148)) = v108 + int32(1)
	if v108 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v15)+152))
	v161 = F_array_set(m, v100, v15+int32(148), v158, int32(4), int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L42
	}
L35:
	;
	v123 = int32(0)
	goto L36
L36:
	;
	v136 = v123 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(24)+v123<<(uint(int32(2))%32))))
	if v25 == v140 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L34
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+148)) = v136
	goto L34
L39:
	;
	goto L40
L40:
	;
	if v136 != v108 {
		v123 = v136
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	v167 = v108
	v175 = v161
	goto L23
L43:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+147)))
	if v186 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+71)) = uint8(v219)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+136)) = base.I64_extend_i32_u(v218)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v230 = F_heap_modify_tuple(m, v55, v223, v15+int32(80), v15+int32(72), v15-int32(-64))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L57
	}
L45:
	;
	if v167 != 0 {
		goto L5
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v196 = F_pg_detoast_datum(m, base.I32_wrap_i64(v184))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v193 = F_construct_array_builtin(m, v15+int32(152), int32(1), int32(25))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v218 = v193
	goto L44
L50:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v198 != int32(1) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196)+20))
	if v201 != int32(1) {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	if v204 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	if v205 != int32(25) {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v196)+16))
	if v208 != v167 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+152))
	v215 = F_array_set(m, v196, v15+int32(148), v212, int32(-1), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v218 = v215
	goto L44
L57:
	;
	F_CatalogTupleUpdate(m, v38, v230+int32(4), v230)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_systable_endscan(m, v53)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_relation_close(m, v38, int32(3))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	m.G0 = v15 + int32(224)
	return int64(0)
L61:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(_a_F_pg_extension_config_dump_0)
	F_errmsg(m, int32(_a_F_pg_extension_config_dump_1), v15+int32(48))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2884), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
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
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v25
	F_errmsg(m, int32(_a_F_pg_extension_config_dump_4), v15)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2895), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
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
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v26
	F_errmsg(m, int32(_a_F_pg_extension_config_dump_5), v15+int32(32))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2901), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
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
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_pg_extension_config_dump[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v304
	F_errmsg_internal(m, int32(_a_F_pg_extension_config_dump_6), v15+int32(16))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2926), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
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
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2959), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
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
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(2992), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
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
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(3004), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
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
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_pg_extension_config_dump_2), int32(3006), int32(_a_F_pg_extension_config_dump_3))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
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
						F_errfinish(m, int32(_a_F_pg_file_exists_3), int32(515), int32(_a_F_pg_file_exists_4))
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
							F_errfinish(m, int32(_a_F_pg_file_exists_3), int32(515), int32(_a_F_pg_file_exists_4))
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
func F_pg_gen_salt_rounds(m *base.Module, l0 int32) int64 {
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v41 = m.ExcPending
	if v41 != 0 {
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
	return base.I64_extend_i32_u(v31)
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v74
	F_errmsg(m, int32(_a_F_pg_gen_salt_rounds_0), v8)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L29
	}
L16:
	;
	v74 = int32(_a_F_pg_gen_salt_rounds_1)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v53 = int32(_a_F_pg_gen_salt_rounds_2)
	goto L20
L19:
	;
	v74 = v68
	goto L15
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v21 != v56 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v68 = v66
	goto L19
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	if v58 == int32(0) {
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
	v74 = int32(_a_F_pg_gen_salt_rounds_3)
	goto L15
L26:
	;
	goto L27
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v21 != v62 {
		v53 = v53 + int32(16)
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v68 = v58
	goto L19
L29:
	;
	F_errfinish(m, int32(_a_F_pg_gen_salt_rounds_4), int32(202), int32(_a_F_pg_gen_salt_rounds_5))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
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
func F_pg_get_constraintdef(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_get_constraintdef_worker(m, v3, int32(0), int32(2), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int64(0)
		} else {
			v17 = F_cstring_to_text(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v17)
				}
			}
		}
	}
}
func F_pg_get_dsm_registry_allocations(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(1))
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
	v14 = int32(_a_F_pg_get_dsm_registry_allocations_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_dsm_registry_allocations[0]))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_dsm_registry_allocations[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_dsm_registry_allocations[0])) = v18
	F_init_dsm_registry(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_dsm_registry_allocations[0])) = v15
	v25 = v6 + int32(52)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_dsm_registry_allocations[2]))
	v28 = int32(0)
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v25)+12)) = v29
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+24)) = uint8(v28)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = int32(-1)
	goto L4
L4:
	;
	v37 = F_dshash_seq_next(m, v25)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v39 = v37
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_dshash_seq_term(m, v6+int32(52))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v42)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)) = uint16(v42)
	v46 = F_cstring_to_text(m, v39)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = base.I64_extend_i32_u(v46)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50<<(uint(int32(2))%32))+uint32(_c_F_pg_get_dsm_registry_allocations[3])))
	v54 = F_cstring_to_text(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = base.I64_extend_i32_u(v54)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
	switch v58 {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	default:
		goto L14
	}
L13:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	F_tuplestore_putvalues(m, v82, v83, v6+int32(16), v6+int32(12))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	v79 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v79)
	goto L13
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v71 == int32(0) {
		goto L14
	} else {
		goto L21
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v64 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v59 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+72)))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v62
	goto L13
L19:
	;
	v67 = F_dsa_get_total_size_from_handle(m, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = base.I64_extend_i32_u(v67)
	goto L13
L21:
	;
	v74 = F_dsa_get_total_size_from_handle(m, v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = base.I64_extend_i32_u(v74)
	goto L13
L23:
	;
	v92 = F_dshash_seq_next(m, v6+int32(52))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v92 != 0 {
		v39 = v92
		goto L9
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	m.G0 = v6 + int32(80)
	return int64(0)
}
func F_pg_get_indexdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int64
	_ = v337
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int64
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int64
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int64
	_ = v598
	var v605 int32
	_ = v605
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v667 int32
	_ = v667
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	v10 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(288)
	m.G0 = v34
	v37 = base.I64_extend_i32_u(l0)
	v38 = F_SearchSysCache1(m, int32(34), v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L200
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L6
	} else {
		goto L197
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L194
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L6
	} else {
		goto L191
	}
L5:
	;
	m.G0 = v34 + int32(288)
	return v667
L6:
	;
	return int32(0)
L7:
	;
	if v38 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l8 != 0 {
		v667 = int32(0)
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
	v60 = v58 + v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v64 = F_SysCacheGetAttrNotNull(m, int32(34), v38, int32(17))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_indexdef_worker_0), v34)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(1311), int32(_a_F_pg_get_indexdef_worker_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
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
	v68 = F_SysCacheGetAttrNotNull(m, int32(34), v38, int32(18))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v72 = F_SysCacheGetAttrNotNull(m, int32(34), v38, int32(19))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v75 = F_SearchSysCache1(m, int32(57), v37)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	if v75 == int32(0) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
	v82 = v80 + v81
	v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v82)+84)))
	v84 = F_SearchSysCache1(m, int32(2), v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v84 == int32(0) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+22)))
	v90 = v88 + v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+68))
	v92 = F_GetIndexAmRoutine(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v94 = int32(0)
	v97 = F_heap_attisnull(m, v38, int32(20), v94)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	v118 = F_get_rel_name(m, v61)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L31
	}
L24:
	;
	if v97 != 0 {
		v116 = v10
		v117 = v94
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v101 = F_SysCacheGetAttrNotNull(m, int32(34), v38, int32(20))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v104 = F_text_to_cstring(m, base.I32_wrap_i64(v101))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v106 = F_stringToNode(m, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	F_pfree(m, v104)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v110 = int32(0)
	if v106 == v110 {
		v116 = v10
		v117 = v110
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v116 = v106
	v117 = v113
	goto L23
L31:
	;
	if v118 == int32(0) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v123 = F_palloc0(m, int32(80))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v126 = F_palloc0(m, int32(136))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+24)) = int32(1)
	v130 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+21)) = uint8(v130)
	*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v61
	v133 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v126)+12)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = int32(101)
	v138 = F_makeAlias(m, v118, v133)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v138
	v142 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+124)) = uint16(v142)
	v144 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+20)) = uint8(v144)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v34)+232)) = v126
	v151 = F_list_make1_impl(m, int32(1), v34+int32(204))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v153 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+20)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v123)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v151
	F_set_rtable_names(m, v123, v153, v153)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_set_simple_column_names(m, v123)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v34)+272)) = v123
	v169 = F_list_make1_impl(m, int32(1), v34+int32(200))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_initStringInfo(m, v34+int32(216))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
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
	v233 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+8)))
	if v233 <= int32(0) {
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
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+12)))
	if v179 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v219 = F_quote_identifier(m, v90+int32(4))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L64
	}
L46:
	;
	v180 = int32(_a_F_pg_get_indexdef_worker_3)
	goto L48
L47:
	;
	v180 = int32(_a_F_pg_get_indexdef_worker_4)
	goto L48
L48:
	;
	v183 = F_quote_identifier(m, v82+int32(4))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	v185 = int32(_a_F_pg_get_indexdef_worker_4)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+119)))
	if v188 != int32(73) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v191 = v185
	goto L52
L51:
	;
	v191 = int32(_a_F_pg_get_indexdef_worker_5)
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
	v192 = v185
	goto L55
L54:
	;
	v192 = v191
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
	v203 = F_quote_identifier(m, v90+int32(4))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L6
	} else {
		goto L62
	}
L57:
	;
	v196 = F_generate_relation_name(m, v61, int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v198 = F_generate_qualified_relation_name(m, v61)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L61
	}
L60:
	;
	v200 = v196
	goto L56
L61:
	;
	v200 = v198
	goto L56
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+176)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v34)+172)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v34)+168)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v34)+164)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v34)+160)) = v180
	F_appendStringInfo(m, v34+int32(216), int32(_a_F_pg_get_indexdef_worker_6), v34+int32(160))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	goto L41
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v219
	F_appendStringInfo(m, v34+int32(216), int32(_a_F_pg_get_indexdef_worker_7), v34+int32(192))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
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
	v237 = int32(24)
	v255 = int32(0)
	v259 = int32(_a_F_pg_get_indexdef_worker_4)
	v264 = v117
	goto L68
L68:
	;
	v281 = v255 << (uint(int32(1)) % 32)
	v283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60+int32(48)+v281))))
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
	v284 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+10)))
	if v284 <= v255 {
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
	v289 = v34 + int32(216)
	v290 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+10)))
	if v290 == v255 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v283 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	F_appendStringInfoString(m, v289, int32(_a_F_pg_get_indexdef_worker_8))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	v296 = v259
	goto L79
L79:
	;
	F_appendStringInfoString(m, v289, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L6
	} else {
		goto L81
	}
L80:
	;
	v296 = int32(_a_F_pg_get_indexdef_worker_4)
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
	v301 = F_get_attname(m, v61, v283, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v264 == int32(0) {
		goto L1
	} else {
		goto L96
	}
L86:
	;
	if l1 != v255+int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v307 = l1
	goto L89
L88:
	;
	v307 = int32(0)
	goto L89
L89:
	;
	if v307 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v312 = F_quote_identifier(m, v301)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	F_get_atttypetypmodcoll(m, v61, v283, v34+int32(212), v34+int32(232), v34+int32(208))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L6
	} else {
		goto L95
	}
L93:
	;
	F_appendStringInfoString(m, v34+int32(216), v312)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v396 = v264
	goto L82
L96:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v330 = v34 + int32(272)
	F_initStringInfo(m, v330)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	v333 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+264)) = uint8(v333)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+248)) = v333
	v337 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+240)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v34)+236)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v34)+268)) = v333
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+267)) = uint8(v333)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+256)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v34)+252)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v34)+232)) = v330
	v348 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+265)) = uint16(v348)
	F_get_rule_expr(m, v328, v34+int32(232), v333)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v356 = v264 + int32(4)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v34)+272))
	if l1 != v255+int32(1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if base.Ui32(v356) < base.Ui32(v326+v327<<(uint(int32(2))%32)) {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v366 = l1
	goto L102
L101:
	;
	v366 = int32(0)
	goto L102
L102:
	;
	if v366 != 0 {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	if v328 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+144)) = v361
	F_appendStringInfo(m, v34+int32(216), int32(_a_F_pg_get_indexdef_worker_9), v34+int32(144))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L6
	} else {
		goto L109
	}
L105:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	switch v369 - int32(15) {
	case 0:
		goto L107
	default:
		goto L104
	case 4, 23, 24, 25, 26, 33:
		goto L106
	}
L106:
	;
	F_appendStringInfoString(m, v34+int32(216), v361)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L6
	} else {
		goto L108
	}
L107:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v328)+16))
	switch v372 {
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
	v386 = v356
	goto L112
L111:
	;
	v386 = int32(0)
	goto L112
L112:
	;
	v387 = F_exprType(m, v328)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v387
	v390 = F_exprCollation(m, v328)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v390
	v396 = v386
	goto L82
L115:
	;
	v494 = v255 + int32(1)
	v495 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+8)))
	if v494 < v495 {
		v255 = v494
		v259 = int32(_a_F_pg_get_indexdef_worker_10)
		v264 = v396
		goto L68
	} else {
		goto L151
	}
L116:
	;
	v398 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+10)))
	if v398 <= v255 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v402 = v255 + int32(1)
	if v402 != l1 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v404 = l1
	goto L120
L119:
	;
	v404 = int32(0)
	goto L120
L120:
	;
	if v404 != 0 {
		goto L115
	} else {
		goto L121
	}
L121:
	;
	v406 = v255 << (uint(int32(2)) % 32)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v64)+v237+v406)))
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281+(base.I32_wrap_i64(v72)+v237)))))
	v412 = F_get_attoptions(m, l0, base.I32_extend16_s(v402))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	if v408 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v406+(base.I32_wrap_i64(v68)+v237))))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v34)+212))
	v433 = base.B2i32(v412 == int64(0))
	if v412 == int64(0) {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
	if v408 == v416 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v418 = F_generate_collation_name(m, v408)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = v418
	F_appendStringInfo(m, v34+int32(216), int32(_a_F_pg_get_indexdef_worker_11), v34+int32(128))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	goto L123
L128:
	;
	v434 = v430
	goto L130
L129:
	;
	v434 = int32(0)
	goto L130
L130:
	;
	v436 = v34 + int32(216)
	F_get_opclass_name(m, v429, v434, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	if v433 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	F_appendStringInfoString(m, v436, int32(_a_F_pg_get_indexdef_worker_12))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L6
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+10)))
	if v449 != int32(1) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	F_get_reloptions(m, v436, v412)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	F_appendStringInfoChar(m, v436, int32(41))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
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
	v453 = v34 + int32(216)
	if v410&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	F_appendStringInfoString(m, v453, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L6
	} else {
		goto L147
	}
L141:
	;
	F_appendStringInfoString(m, v453, int32(_a_F_pg_get_indexdef_worker_13))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L6
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if v410&int32(2) == int32(0) {
		goto L138
	} else {
		goto L146
	}
L144:
	;
	if v410&int32(2) != 0 {
		goto L138
	} else {
		goto L145
	}
L145:
	;
	v467 = int32(_a_F_pg_get_indexdef_worker_14)
	goto L140
L146:
	;
	v467 = int32(_a_F_pg_get_indexdef_worker_15)
	goto L140
L147:
	;
	goto L138
L148:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l2+v406)))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v34)+212))
	v476 = F_generate_operator_name(m, v474, v475, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+112)) = v476
	F_appendStringInfo(m, v34+int32(216), int32(_a_F_pg_get_indexdef_worker_16), v34+int32(112))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
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
	F_ReleaseCatCache(m, v38)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L6
	} else {
		goto L188
	}
L153:
	;
	v529 = v34 + int32(216)
	F_appendStringInfoChar(m, v529, int32(41))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L6
	} else {
		goto L154
	}
L154:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+13)))
	if v533 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_appendStringInfoString(m, v529, int32(_a_F_pg_get_indexdef_worker_17))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L6
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v539 = F_flatten_reloptions(m, l0)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L6
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	if v539 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+96)) = v539
	F_appendStringInfo(m, v34+int32(216), int32(_a_F_pg_get_indexdef_worker_18), v34+int32(96))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
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
	F_pfree(m, v539)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v577 = F_heap_attisnull(m, v38, int32(21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L6
	} else {
		goto L176
	}
L166:
	;
	v553 = F_get_rel_tablespace(m, l0)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	if v553 == int32(0) {
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
	F_appendStringInfoString(m, v34+int32(216), int32(_a_F_pg_get_indexdef_worker_19))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L6
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v562 = F_get_tablespace_name(m, v553)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L6
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v564 = F_quote_identifier(m, v562)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+80)) = v564
	F_appendStringInfo(m, v34+int32(216), int32(_a_F_pg_get_indexdef_worker_20), v34+int32(80))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	goto L165
L176:
	;
	if v577 != 0 {
		goto L152
	} else {
		goto L177
	}
L177:
	;
	v581 = F_SysCacheGetAttrNotNull(m, int32(34), v38, int32(21))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	v584 = F_text_to_cstring(m, base.I32_wrap_i64(v581))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	v586 = F_stringToNode(m, v584)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	F_pfree(m, v584)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	v591 = v34 + int32(272)
	F_initStringInfo(m, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L6
	} else {
		goto L182
	}
L182:
	;
	v594 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+264)) = uint8(v594)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+248)) = v594
	v598 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+240)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v34)+236)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v34)+268)) = v594
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+267)) = uint8(v594)
	v605 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+265)) = uint16(v605)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+256)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v34)+252)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v34)+232)) = v591
	F_get_rule_expr(m, v586, v34+int32(232), v594)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L6
	} else {
		goto L183
	}
L183:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v34)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v616
	if l2 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v622 = int32(_a_F_pg_get_indexdef_worker_21)
	goto L186
L185:
	;
	v622 = int32(_a_F_pg_get_indexdef_worker_22)
	goto L186
L186:
	;
	F_appendStringInfo(m, v34+int32(216), v622, v34-int32(-64))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	goto L152
L188:
	;
	F_ReleaseCatCache(m, v75)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	F_ReleaseCatCache(m, v84)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v34)+216))
	v667 = v635
	goto L5
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_indexdef_worker_23), v34+int32(16))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(1336), int32(_a_F_pg_get_indexdef_worker_2))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
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
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v82)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v691
	F_errmsg_internal(m, int32(_a_F_pg_get_indexdef_worker_24), v34+int32(32))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L6
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(1345), int32(_a_F_pg_get_indexdef_worker_2))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v61
	F_errmsg_internal(m, int32(_a_F_pg_get_indexdef_worker_23), v34+int32(48))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(_a_F_pg_get_indexdef_worker_25), int32(_a_F_pg_get_indexdef_worker_26))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
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
	v725 = m.ExcPending
	if v725 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_pg_get_indexdef_worker_1), int32(1443), int32(_a_F_pg_get_indexdef_worker_2))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
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
func F_pg_get_keywords(m *base.Module, l0 int32) int64 {
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
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int64
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
			return int64(0)
		} else {
			v21 = int32(_a_F_pg_get_keywords_0)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[0]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_pg_get_keywords[0])) = v24
			v27 = F_get_call_result_type(m, l0, int32(0), v11)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				if v27 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_get_keywords_1), int32(0))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_get_keywords_2), int32(404), int32(_a_F_pg_get_keywords_3))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int64(0)
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
						return int64(0)
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
								return int64(0)
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
									return int64(0)
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
								return int64(0)
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(2)
								v97 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v97)
								v104 = int64(0)
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
				return int64(0)
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
					return int64(0)
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
				return int64(0)
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(2)
				v97 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v97)
				v104 = int64(0)
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
func F_pg_get_publication_tables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v412 int32
	_ = v412
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v481 int32
	_ = v481
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v532 int32
	_ = v532
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int64
	_ = v722
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int64
	_ = v739
	var v740 int64
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int64
	_ = v746
	var v747 int64
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int64
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v764 int64
	_ = v764
	var v765 int32
	_ = v765
	var v771 int64
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v802 int32
	_ = v802
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int64
	_ = v873
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int64
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v916 int32
	_ = v916
	var v929 int64
	_ = v929
	v6 = int32(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v24 == v6 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v916 - int32(-64)
	return v929
L2:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L7
	} else {
		goto L180
	}
L3:
	;
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L7
	} else {
		goto L179
	}
L4:
	;
	v27 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v688 = v21 + int32(32)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)+16))
	goto L139
L7:
	;
	return int64(0)
L8:
	;
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v33 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v55 = int32(_a_F_pg_get_publication_tables_0)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_tables[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_tables[0])) = v58
	F_deconstruct_array_builtin(m, l1, int32(25), v19+int32(-48), int32(0), v19+int32(-4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L19
	}
L12:
	;
	if v33 == int32(0) {
		v890 = v21
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v39 = v37 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+119)))
	switch v40 - int32(112) {
	case 0, 2:
		goto L14
	default:
		goto L3
	}
L14:
	;
	goto L15
L15:
	;
	if base.B2i32(base.Ui32(l2) < base.Ui32(int32(_a_F_pg_get_publication_tables_1)))|base.B2i32(base.Ui32(l2) < base.Ui32(int32(_a_F_pg_get_publication_tables_2))) != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+118)))
	if v48 != int32(112) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v68 <= int32(0) {
		v532 = v6
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v540 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L7
	} else {
		goto L113
	}
L21:
	;
	v71 = base.I64_extend_i32_u(l2)
	v81 = v6
	v83 = v6
	v84 = v6
	goto L22
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v81<<(uint(int32(3))%32))))
	v95 = F_text_to_cstring(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L25
	}
L23:
	;
	v356 = int32(0)
	if v346&base.B2i32(v345 != v356) == v356 {
		v532 = v345
		goto L20
	} else {
		goto L88
	}
L24:
	;
	v353 = v81 + int32(1)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v353 < v354 {
		v81 = v353
		v83 = v345
		v84 = v346
		goto L22
	} else {
		goto L87
	}
L25:
	;
	v97 = F_get_publication_oid(m, v95, l4)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	if v97 == int32(0) {
		v345 = v83
		v346 = v84
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v101 = F_GetPublication(m, v97)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	if v101 == int32(0) {
		v345 = v83
		v346 = v84
		goto L24
	} else {
		goto L29
	}
L29:
	;
	if l3 != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+10)))
	v345 = v325
	v346 = v332 | v84
	goto L24
L31:
	;
	if v272 == int32(0) {
		v325 = v83
		goto L30
	} else {
		goto L80
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = l2
	v252 = F_list_make1_impl(m, int32(480), v19+int32(-52))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L79
	}
L33:
	;
	v229 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v101))))
	v230 = int64(0)
	v232 = F_SearchSysCacheExists(m, int32(53), v71, v229, v230, v230)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L7
	} else {
		goto L74
	}
L34:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+10)))
	if v105 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+10)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+8)))
	if v149 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L37:
	;
	v108 = F_get_rel_relkind(m, l2)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v112 = F_get_rel_relispartition(m, l2)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L43
	}
L40:
	;
	if v108 == int32(112) {
		v325 = v83
		goto L30
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v142 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v101))))
	v143 = int64(0)
	v145 = F_SearchSysCacheExists(m, int32(53), v140, v142, v143, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L57
	}
L43:
	;
	if v112 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v114 = F_get_partition_ancestors(m, l2)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+8)))
	if v136 == int32(0) {
		goto L33
	} else {
		goto L56
	}
L47:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+8)))
	if v116 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+10)))
	if v117 != 0 {
		v325 = v83
		goto L30
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v129 = F_GetTopMostAncestorInPublication(m, v128, v114)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	if v114 == int32(0) {
		v140 = v71
		goto L42
	} else {
		goto L52
	}
L52:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v127 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v120+v121<<(uint(int32(2))%32)-int32(4)))))
	v140 = v127
	goto L42
L53:
	;
	if v129 == int32(0) {
		goto L33
	} else {
		goto L54
	}
L54:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+10)))
	if v133 == int32(0) {
		goto L32
	} else {
		goto L55
	}
L55:
	;
	v325 = v83
	goto L30
L56:
	;
	v140 = v71
	goto L42
L57:
	;
	if v145 != 0 {
		v325 = v83
		goto L30
	} else {
		goto L58
	}
L58:
	;
	goto L32
L59:
	;
	v155 = F_GetAllPublicationRelations(m, v148, int32(114), v147&int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L7
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v157 = int32(0)
	v161 = F_get_publication_relations(m, v148, v147^int32(1), v157)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L63
	}
L62:
	;
	v272 = v155
	goto L31
L63:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+10)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v165 = F_GetPublicationSchemas(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L7
	} else {
		goto L65
	}
L64:
	;
	v224 = F_list_concat_unique_oid(m, v161, v219)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L7
	} else {
		goto L73
	}
L65:
	;
	if v165 == int32(0) {
		v219 = v157
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v169 = int32(0)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v170 <= v169 {
		v219 = v157
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v176 = v169
	v188 = v157
	goto L68
L68:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193+v176<<(uint(int32(2))%32))))
	v198 = F_GetSchemaPublicationRelations(m, v197, v163^int32(1))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L70
	}
L69:
	;
	v219 = v200
	goto L64
L70:
	;
	v200 = F_list_concat(m, v188, v198)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	v203 = v176 + int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v203 < v204 {
		v176 = v203
		v188 = v200
		goto L68
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v272 = v224
	goto L31
L74:
	;
	if v232 != 0 {
		goto L32
	} else {
		goto L75
	}
L75:
	;
	v235 = F_get_rel_namespace(m, l2)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	v238 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v101))))
	v239 = int64(0)
	v241 = F_SearchSysCacheExists(m, int32(50), base.I64_extend_i32_u(v235), v238, v239, v239)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	if v241 == int32(0) {
		v325 = v83
		goto L30
	} else {
		goto L78
	}
L78:
	;
	goto L32
L79:
	;
	v272 = v252
	goto L31
L80:
	;
	v275 = int32(0)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v276 <= v275 {
		v325 = v83
		goto L30
	} else {
		goto L81
	}
L81:
	;
	v280 = v275
	v290 = v83
	goto L82
L82:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v299 = F_palloc(m, int32(8))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L7
	} else {
		goto L84
	}
L83:
	;
	v325 = v308
	goto L30
L84:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v297+v280<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v299))) = v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v299)+4)) = v306
	v308 = F_lappend(m, v290, v299)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	v311 = v280 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v311 < v312 {
		v280 = v311
		v290 = v308
		goto L82
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	goto L23
L88:
	;
	v364 = v345
	v371 = v356
	goto L89
L89:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	if v380 <= v371 {
		v532 = v345
		goto L20
	} else {
		goto L91
	}
L90:
	;
	v532 = v345
	goto L20
L91:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v382+v371<<(uint(int32(2))%32))))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v388 = F_get_rel_relispartition(m, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L7
	} else {
		goto L94
	}
L92:
	;
	if v503 != 0 {
		v364 = v503
		v371 = v502 + int32(1)
		goto L89
	} else {
		goto L112
	}
L93:
	;
	v502 = v371
	v503 = v364
	goto L92
L94:
	;
	if v388 == int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v393 = F_get_partition_ancestors(m, v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	if v393 == int32(0) {
		goto L93
	} else {
		goto L97
	}
L97:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v397 <= int32(0) {
		goto L93
	} else {
		goto L98
	}
L98:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	v412 = int32(0)
	goto L99
L99:
	;
	if v364 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L93
L101:
	;
	v481 = v412 + int32(1)
	if v397 != v481 {
		v412 = v481
		goto L99
	} else {
		goto L111
	}
L102:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	if v422 <= int32(0) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v400+v412<<(uint(int32(2))%32))))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v432 = int32(0)
	goto L104
L104:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v429+v432<<(uint(int32(2))%32))))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	if v428 != v453 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v460 = F_list_delete_nth_cell(m, v364, v371)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L7
	} else {
		goto L110
	}
L106:
	;
	v456 = v432 + int32(1)
	if v456 != v422 {
		v432 = v456
		goto L104
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	goto L105
L109:
	;
	goto L101
L110:
	;
	v502 = v371 - int32(1)
	v503 = v460
	goto L92
L111:
	;
	goto L100
L112:
	;
	goto L90
L113:
	;
	F_TupleDescInitEntry(m, v540, int32(1), int32(_a_F_pg_get_publication_tables_3), int32(26), int32(-1), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	F_TupleDescInitEntry(m, v540, int32(2), int32(_a_F_pg_get_publication_tables_4), int32(26), int32(-1), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	F_TupleDescInitEntry(m, v540, int32(3), int32(_a_F_pg_get_publication_tables_5), int32(22), int32(-1), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	F_TupleDescInitEntry(m, v540, int32(4), int32(_a_F_pg_get_publication_tables_6), int32(194), int32(-1), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	v570 = int32(0)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	if v570 < v579 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v657 = F_BlessTupleDesc(m, v540)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L7
	} else {
		goto L137
	}
L119:
	;
	v583 = v540 + int32(28)
	v590 = v570
	v591 = v579
	v593 = v570
	goto L123
L120:
	;
	v647 = v570
	v654 = v579
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v540)+20)) = v654
	*(*int32)(unsafe.Add(mBase, uint32(v540)+16)) = v647
	goto L118
L122:
	;
	v647 = v641
	v654 = v620
	goto L121
L123:
	;
	v599 = v583 + v579<<(uint(int32(3))%32) + v590*int32(100)
	v602 = v583 + v590<<(uint(int32(3))%32)
	if v579 != v591 {
		v620 = v591
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v641 = v579
	goto L122
L125:
	;
	v621 = int32(*(*int16)(unsafe.Add(mBase, uint32(v602)+2)))
	if v621 <= int32(0) {
		v641 = v590
		goto L122
	} else {
		goto L133
	}
L126:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+7)))
	if v604 != int32(118) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v620 = v590
	goto L125
L128:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+4)))
	if v607 != int32(1) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+6)))
	if v610&int32(6) != 0 {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	v613 = int32(*(*int16)(unsafe.Add(mBase, uint32(v602)+2)))
	if v613 <= int32(0) {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599)+90)))
	if v616 != int32(118) {
		v620 = v579
		goto L125
	} else {
		goto L132
	}
L132:
	;
	goto L127
L133:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599)+90)))
	if v624 == int32(118) {
		v641 = v590
		goto L122
	} else {
		goto L134
	}
L134:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+5)))
	v633 = (v593 + v627 - int32(1)) & (int32(0) - v627)
	if int32(_a_F_pg_get_publication_tables_7) < v633 {
		v641 = v590
		goto L122
	} else {
		goto L135
	}
L135:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v602))) = uint16(v633)
	v639 = v590 + int32(1)
	if v639 != v579 {
		v590 = v639
		v591 = v620
		v593 = v633 + v621
		goto L123
	} else {
		goto L136
	}
L136:
	;
	goto L124
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v657
	v661 = F_palloc(m, int32(8))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v661)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v661))) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v661
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_tables[0])) = v56
	goto L6
L139:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+16))
	goto L141
L140:
	;
	v890 = v21
	goto L2
L141:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v691)))
	if v710 != 0 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v720)+4))
	v737 = F_GetPublication(m, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L7
	} else {
		goto L149
	}
L143:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)+4))
	v713 = v711
	goto L145
L144:
	;
	v713 = int32(0)
	goto L145
L145:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	if v713 <= v714 {
		goto L140
	} else {
		goto L146
	}
L146:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v710)+12))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v716+v714<<(uint(int32(2))%32))))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	v722 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v688)+8)) = v722
	*(*int64)(unsafe.Add(mBase, uint32(v688))) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = int32(0)
	v728 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v691)+4)) = v714 + v728
	v732 = F_try_table_open(m, v721, v728)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	if v732 == int32(0) {
		goto L141
	} else {
		goto L148
	}
L148:
	;
	goto L142
L149:
	;
	v739 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v737))))
	v740 = base.I64_extend_i32_u(v721)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v740
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v739
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737)+8)))
	if v743 != 0 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	F_relation_close(m, v732, int32(1))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L7
	} else {
		goto L176
	}
L151:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v732)+52))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v782 = F_palloc_mul(m, int32(2), v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L7
	} else {
		goto L162
	}
L152:
	;
	v761 = v21 + int32(60)
	v764 = F_SysCacheGetAttr(m, int32(53), v753, int32(6), v761|int32(2))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L7
	} else {
		goto L159
	}
L153:
	;
	v756 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+62)) = uint16(v756)
	goto L151
L154:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v732)+48))
	v746 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v745)+68)))
	v747 = int64(0)
	v749 = F_SearchSysCacheExists(m, int32(50), v746, v739, v747, v747)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L7
	} else {
		goto L155
	}
L155:
	;
	if v749 != 0 {
		goto L153
	} else {
		goto L156
	}
L156:
	;
	v752 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v737))))
	v753 = F_SearchSysCacheCopy(m, int32(53), v740, v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L7
	} else {
		goto L157
	}
L157:
	;
	if v753 != 0 {
		goto L152
	} else {
		goto L158
	}
L158:
	;
	goto L153
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = v764
	v771 = F_SysCacheGetAttr(m, int32(53), v753, int32(5), v761|int32(3))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L7
	} else {
		goto L160
	}
L160:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = v771
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+62)))
	if v774 != int32(1) {
		goto L150
	} else {
		goto L161
	}
L161:
	;
	goto L151
L162:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	if v784 <= int32(0) {
		goto L150
	} else {
		goto L163
	}
L163:
	;
	v787 = int32(0)
	v790 = v787
	v791 = v787
	v802 = v784
	goto L164
L164:
	;
	v812 = v780 + v802<<(uint(int32(3))%32) + v790*int32(100)
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812)+119)))
	if v813 != 0 {
		v830 = v791
		v833 = v802
		goto L166
	} else {
		goto L167
	}
L165:
	;
	if v830 <= int32(0) {
		goto L150
	} else {
		goto L174
	}
L166:
	;
	v835 = v790 + int32(1)
	if v835 < v833 {
		v790 = v835
		v791 = v830
		v802 = v833
		goto L164
	} else {
		goto L173
	}
L167:
	;
	v815 = v812 + int32(28)
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+90)))
	if v816 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if v816 != int32(115) {
		v830 = v791
		v833 = v802
		goto L166
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v822 = int32(1)
	v825 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v815)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v782+v791<<(uint(v822)%32)))) = uint16(v825)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v830 = v791 + v822
	v833 = v829
	goto L166
L171:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v737)+12))
	if v819 != int32(115) {
		v830 = v791
		v833 = v802
		goto L166
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	goto L165
L174:
	;
	v839 = F_buildint2vector(m, v782, v830)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L7
	} else {
		goto L175
	}
L175:
	;
	v841 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+62)) = uint8(v841)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = base.I64_extend_i32_u(v839)
	goto L150
L176:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v690)+28))
	v871 = F_heap_form_tuple(m, v866, v21+int32(16), v21+int32(60))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	v873 = *(*int64)(unsafe.Add(mBase, uint32(v690)))
	*(*int64)(unsafe.Add(mBase, uint32(v690))) = v873 + int64(1)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v877)+20)) = int32(1)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v871)+16))
	v881 = F_HeapTupleHeaderGetDatum(m, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L7
	} else {
		goto L178
	}
L178:
	;
	v916 = v21
	v929 = v881
	goto L1
L179:
	;
	v890 = v21
	goto L2
L180:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v905)+20)) = int32(2)
	v908 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v908)
	v916 = v890
	v929 = int64(0)
	goto L1
}
func F_pg_get_publication_tables_a(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = int32(0)
		v10 = F_pg_get_publication_tables(m, l0, v3, v7, v7, v7)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return v10
		}
	}
}
func F_pg_get_sequence_data(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	v2 = int32(0)
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+28)) = uint16(v2)
	v24 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_TupleDescInitEntry(m, v24, int32(1), int32(_a_F_pg_get_sequence_data_0), int32(20), int32(-1), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_TupleDescInitEntry(m, v24, int32(2), int32(_a_F_pg_get_sequence_data_1), int32(16), int32(-1), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_TupleDescInitEntry(m, v24, int32(3), int32(_a_F_pg_get_sequence_data_2), int32(3220), int32(-1), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v49 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v49 < v58 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v136 = F_BlessTupleDesc(m, v24)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L25
	}
L7:
	;
	v62 = v24 + int32(28)
	v69 = v49
	v70 = v58
	v72 = v49
	goto L11
L8:
	;
	v126 = v49
	v133 = v58
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v126
	goto L6
L10:
	;
	v126 = v120
	v133 = v99
	goto L9
L11:
	;
	v78 = v62 + v58<<(uint(int32(3))%32) + v69*int32(100)
	v81 = v62 + v69<<(uint(int32(3))%32)
	if v58 != v70 {
		v99 = v70
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v120 = v58
	goto L10
L13:
	;
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v81)+2)))
	if v100 <= int32(0) {
		v120 = v69
		goto L10
	} else {
		goto L21
	}
L14:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+7)))
	if v83 != int32(118) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v99 = v69
	goto L13
L16:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+4)))
	if v86 != int32(1) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+6)))
	if v89&int32(6) != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v81)+2)))
	if v92 <= int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+90)))
	if v95 != int32(118) {
		v99 = v58
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+90)))
	if v103 == int32(118) {
		v120 = v69
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+5)))
	v112 = (v72 + v106 - int32(1)) & (int32(0) - v106)
	if int32(_a_F_pg_get_sequence_data_3) < v112 {
		v120 = v69
		goto L10
	} else {
		goto L23
	}
L23:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v81))) = uint16(v112)
	v118 = v69 + int32(1)
	if v118 != v58 {
		v69 = v118
		v70 = v99
		v72 = v112 + v100
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	v139 = F_try_relation_open(m, v12, int32(1))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L29
	}
L26:
	;
	v226 = F_heap_form_tuple(m, v136, v8+int32(-32), v8+int32(-36))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L52
	}
L27:
	;
	F_relation_close(m, v139, int32(1))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L51
	}
L28:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)) = uint8(v209)
	v211 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+28)) = uint16(v211)
	goto L27
L29:
	;
	if v139 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)+48))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+119)))
	if v142 != int32(83) {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v205 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)) = uint8(v205)
	v207 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+28)) = uint16(v207)
	goto L26
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[0]))
	v148 = F_pg_class_aclcheck(m, v12, v146, int64(2))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v148 != 0 {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v139)+48))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+118)))
	switch v151 - int32(112) {
	case 0:
		goto L36
	default:
		goto L37
	case 4:
		goto L38
	}
L36:
	;
	v174 = F_read_seq_tuple(m, v139, v8+int32(-40), v8+int32(-60))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L45
	}
L37:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[1])))
	if v159 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+24)))
	if v154 != int32(1) {
		goto L28
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	if v169 != 0 {
		goto L28
	} else {
		goto L44
	}
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[2]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+308))
	v167 = base.B2i32(v165 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[1])) = uint8(v167)
	v169 = v167
	goto L43
L42:
	;
	v169 = int32(0)
	goto L43
L43:
	;
	goto L40
L44:
	;
	goto L36
L45:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v176 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v195
	v197 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v174)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v197
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = base.I64_rotl(v199, int64(32))
	F_UnlockReleaseBuffer(m, v176)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L50
	}
L47:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[3]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180+(v176^int32(-1))<<(uint(int32(2))%32))))
	v194 = v186
	goto L46
L48:
	;
	goto L49
L49:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_sequence_data[4]))
	v194 = v188 + v176<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	goto L27
L51:
	;
	goto L26
L52:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
	v229 = F_HeapTupleHeaderGetDatum(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	m.G0 = v10 - int32(-64)
	return v229
}
func F_pg_get_shmem_allocations(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[0]))
	v22 = F_LWLockAcquire(m, v18+int32(16), int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = v9 + int32(60)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[1]))
	F_hash_seq_init(m, v25, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	v32 = F_hash_seq_search(m, v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = v32
	v37 = v2
	goto L9
L7:
	;
	v71 = v2
	goto L8
L8:
	;
	v75 = F_cstring_to_text(m, int32(_a_F_pg_get_shmem_allocations_0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	v40 = F_cstring_to_text(m, v34)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v71 = v63
	goto L8
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = base.I64_extend_i32_u(v40)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = base.I64_extend_i32_s(v44 - v46)
	v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+52)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = base.I64_extend_i32_u(v52)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v55, v56, v9+int32(16), v9+int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v63 = v52 + v37
	v66 = F_hash_seq_search(m, v9+int32(60))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v66 != 0 {
		v34 = v66
		v37 = v63
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+13)) = uint8(v77)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = base.I64_extend_i32_u(v75)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[0]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v85 = base.I64_extend_i32_u(v83 - v71)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v85
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v91 = v9 + int32(16)
	v93 = v9 + int32(12)
	F_tuplestore_putvalues(m, v88, v89, v91, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)) = uint8(v96)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[0]))
	v100 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v99))))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v100
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+13)) = uint8(v102)
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[2]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v109 = base.I64_extend_i32_u(v106 - v107)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v109
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v112, v113, v91, v93)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_shmem_allocations[0]))
	F_LWLockRelease(m, v117+int32(16))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v9 + int32(80)
	return int64(0)
}
func F_pg_get_shmem_allocations_numa(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v2 = m.G0
	m.G0 = v2 - int32(80)
	F_errstart_cold(m, int32(21), int32(0))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		F_errmsg_internal(m, int32(_a_F_pg_get_shmem_allocations_numa_0), int32(0))
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			F_errfinish(m, int32(_a_F_pg_get_shmem_allocations_numa_1), int32(1231), int32(_a_F_pg_get_shmem_allocations_numa_2))
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
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
func F_pg_has_role_id(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_has_role_id[0]))
		v12 = F_convert_any_priv_string(m, v5, int32(_a_F_pg_has_role_id_0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = F_pg_role_aclcheck(m, v3, v10, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v14 ^ int32(1))
			}
		}
	}
}
func F_pg_has_role_id_id(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v10 = F_convert_any_priv_string(m, v5, int32(_a_F_pg_has_role_id_id_0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = F_pg_role_aclcheck(m, v2, v3, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v12 ^ int32(1))
			}
		}
	}
}
func F_pg_has_role_name_id(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int64
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
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v18 = int64(0)
		v21 = F_GetSysCacheOid(m, int32(10), v11, v18, v18, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v11)
						F_errmsg(m, int32(_a_F_pg_has_role_name_id_0), v8)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_has_role_name_id_1), int32(_a_F_pg_has_role_name_id_2), int32(_a_F_pg_has_role_name_id_3))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v43 = F_convert_any_priv_string(m, v13, int32(_a_F_pg_has_role_name_id_4))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = F_pg_role_aclcheck(m, base.I32_wrap_i64(v10), v21, v43)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						m.G0 = v8 + int32(16)
						return base.I64_extend_i32_u(v45 ^ int32(1))
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
func F_pg_jit_available(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_provider_init(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v2)
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
func F_pg_my_temp_schema(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_pg_my_temp_schema[0])))
	return v3
}
func F_pg_node_tree_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_pg_node_tree_in_0), int32(334), int32(_a_F_pg_node_tree_in_1), int32(_a_F_pg_node_tree_in_2), int32(_a_F_pg_node_tree_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_pg_notify(m *base.Module, l0 int32) int64 {
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
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = F_text_to_cstring(m, v10)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v16 = v14
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
				if v17 == int32(0) {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v21 = F_pg_detoast_datum_packed(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v23 = F_text_to_cstring(m, v21)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int64(0)
						} else {
							v25 = v23
							F_PreventCommandDuringRecovery(m, int32(_a_F_pg_notify_1))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int64(0)
							} else {
								F_Async_Notify(m, v16, v25)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int64(0)
								} else {
									return int64(0)
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
						return int64(0)
					} else {
						F_Async_Notify(m, v16, v25)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							return int64(0)
						}
					}
				}
			}
		}
	} else {
		v16 = v4
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v17 == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v21 = F_pg_detoast_datum_packed(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v23 = F_text_to_cstring(m, v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = v23
					F_PreventCommandDuringRecovery(m, int32(_a_F_pg_notify_1))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						F_Async_Notify(m, v16, v25)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							return int64(0)
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
				return int64(0)
			} else {
				F_Async_Notify(m, v16, v25)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					return int64(0)
				}
			}
		}
	}
}
func F_pg_num_nulls(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v12 = F_count_nulls(m, l0, v6+int32(12), v6+int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v22 = int64(0)
		} else {
			v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
			v22 = v21
		}
		m.G0 = v6 + int32(16)
		return v22
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
	F_errfinish(m, int32(_a_F_pg_perm_setlocale_9), int32(253), int32(_a_F_pg_perm_setlocale_10))
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
func F_pg_read_binary_file_all(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_convert_and_check_filename(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v13 = F_read_binary_file(m, v8, int64(0), int64(-1), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				if v13 == int32(0) {
					v17 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v13)
				}
			}
		}
	}
}
func F_pg_read_file_all(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_convert_and_check_filename(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v13 = F_read_binary_file(m, v8, int64(0), int64(-1), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				if v13 == int32(0) {
					v17 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
					return int64(0)
				} else {
					v21 = int32(4)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					F_pg_verifymbstr(m, v13+v21, int32(base.Ui32(v23)>>(uint(int32(2))%32))-v21)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v13)
					}
				}
			}
		}
	}
}
func F_pg_read_file_off_len(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v11 = F_pg_read_file_common(m, v4, v8, v9, int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			if v11 == int32(0) {
				v15 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v15)
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v11)
			}
		}
	}
}
func F_pg_regcomp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int64
	_ = v148
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v712 int32
	_ = v712
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v743 int32
	_ = v743
	var v757 int32
	_ = v757
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v843 int32
	_ = v843
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v863 int32
	_ = v863
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(144)
	m.G0 = v11
	v19 = int32(2)
	if base.B2i32(l0 == v6)|base.B2i32(l1 == v6)|(int32(base.Ui32(l3)>>(uint(v19)%32))&base.B2i32(l3&int32(227) != v6)|base.B2i32(l3&int32(3) == v19)) == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pg_set_regex_collation(m, l4)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v863 = int32(16)
	goto L3
L3:
	;
	m.G0 = v11 + int32(144)
	return v863
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v11 + int32(52)
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l1 + l2<<(uint(int32(2))%32)
	v58 = v43
	goto L6
L6:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v63+v58<<(uint(int32(2))%32)))) = int32(0)
	v70 = v58 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	if base.Ui32(v70) < base.Ui32(v71) {
		v58 = v70
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v73 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v73
	v75 = int32(_a_F_pg_regcomp_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+100)) = uint16(v75)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+92)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v11)+124)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v11)+132)) = v73
	v85 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+140)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(_a_F_pg_regcomp_1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(17179869184)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_pg_regcomp_2)
	v98 = F_palloc_extended(m, int32(432), int32(2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v98
	if v98 == int32(0) {
		v673 = int32(12)
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v681 = v11 + int32(4)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	if v682 != 0 {
		goto L172
	} else {
		goto L173
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+72)) = int32(2166)
	v106 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v98)+200)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v98)+192)) = int64(0)
	v112 = int32(_a_F_pg_regcomp_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+188)) = uint16(v112)
	*(*int64)(unsafe.Add(mBase, uint32(v98)+180)) = int64(4294969344)
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+88)) = uint16(v106)
	*(*int64)(unsafe.Add(mBase, uint32(v98)+80)) = int64(10)
	*(*int32)(unsafe.Add(mBase, uint32(v98)+92)) = v98 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v98)+76)) = v11 + int32(4)
	v128 = F_palloc_extended(m, int32(_a_F_pg_regcomp_3), int32(2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+96)) = v128
	if v128 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v98 + int32(72)
	v187 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v98)+424)) = int64(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	v195 = F_newnfa(m, v11+int32(4), v193, v187)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L27
	}
L14:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+24)) = int32(101)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v139 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	base.MemoryFill(m, v128, int32(0), int32(_a_F_pg_regcomp_3))
	v148 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v98)+156)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v98)+148)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v98)+140)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v98)+132)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v98)+124)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v98)+116)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v98)+108)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v98)+100)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v98)+176)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v98)+168)) = int64(4294967300)
	v170 = F_palloc_extended(m, int32(8), int32(2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L20
	}
L17:
	;
	v141 = v139
	goto L19
L18:
	;
	v141 = int32(12)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v98)+160)) = int64(0)
	goto L13
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+164)) = v170
	if v170 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+24)) = int32(101)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	if v179 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v183 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v170))) = uint16(v183)
	goto L13
L24:
	;
	v181 = v179
	goto L26
L25:
	;
	v181 = int32(12)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+12)) = v181
	goto L13
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v198 != 0 {
		v673 = v198
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v201 = F_palloc_extended(m, int32(588), int32(2))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if v201 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = int32(0)
	v673 = int32(12)
	goto L10
L31:
	;
	goto L32
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v201))) = int64(429496729600)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v201)+12)) = int64(85899345920)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+20)) = v201 + int32(428)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v201 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = v201
	v221 = int32(4)
	v222 = v11 + v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	if v223&v221 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	if v459 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L34:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v226-v227 < int32(13) {
		v283 = v223
		v284 = v227
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v285 = int32(3)
	if v283&v285 != v285 {
		goto L33
	} else {
		goto L50
	}
L36:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if v231 != int32(42) {
		v283 = v223
		v284 = v227
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v234 != int32(42) {
		v283 = v223
		v284 = v227
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	if v237 != int32(42) {
		v283 = v223
		v284 = v227
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	switch v240 - int32(58) {
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
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+8)) = v271 | int32(128)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v277 = v275 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v222)+16)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v281 = v279 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v281
	v283 = v277
	v284 = v281
	goto L35
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+24)) = int32(101)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	if v266 != 0 {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v249)+8)) = v250 | int32(128)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v254 + int32(16)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v222)+16)) = v258&int32(-232) | int32(4)
	goto L33
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+24)) = int32(101)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	if v245 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v247 = v245
	goto L46
L45:
	;
	v247 = int32(2)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+12)) = v247
	goto L33
L47:
	;
	v268 = v266
	goto L49
L48:
	;
	v268 = int32(13)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+12)) = v268
	goto L33
L50:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	if v289-v284 < int32(9) {
		goto L33
	} else {
		goto L51
	}
L51:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	if v293 != int32(40) {
		goto L33
	} else {
		goto L52
	}
L52:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v296 != int32(63) {
		goto L33
	} else {
		goto L53
	}
L53:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_pg_regcomp[0]))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+2)))
	if v302 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v318)+8)) = v319 | int32(128)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v325 = v323 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	if base.Ui32(v327) <= base.Ui32(v325) {
		goto L63
	} else {
		goto L64
	}
L55:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v299) {
		goto L33
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+20))
	v314 = m.T0[v313].(func(*base.Module, int32, int32) int32)(m, v299, v301)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299)+uint32(_c_F_pg_regcomp[1]))))
	if v307&int32(2) == int32(0) {
		goto L33
	} else {
		goto L59
	}
L59:
	;
	goto L54
L60:
	;
	if v314 == int32(0) {
		goto L33
	} else {
		goto L61
	}
L61:
	;
	goto L54
L62:
	;
	v440 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v418 + v440
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	if v443&v440 == int32(0) {
		goto L33
	} else {
		goto L100
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+24)) = int32(101)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	if v436 != 0 {
		goto L97
	} else {
		goto L98
	}
L64:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_pg_regcomp[0]))
	v333 = v325
	v337 = v330
	v338 = v327
	goto L65
L65:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+2)))
	if v340 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	if base.Ui32(v421) <= base.Ui32(v418) {
		goto L63
	} else {
		goto L95
	}
L67:
	;
	goto L66
L68:
	;
	switch v360 - int32(98) {
	case 0:
		goto L79
	case 1:
		goto L90
	default:
		goto L80
	case 3:
		goto L89
	case 7:
		goto L88
	case 11, 12:
		goto L87
	case 14:
		goto L86
	case 15:
		goto L85
	case 17:
		goto L84
	case 18:
		goto L83
	case 21:
		goto L82
	case 22:
		goto L81
	}
L69:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v339) {
		v418 = v333
		v421 = v338
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v337)+8))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+20))
	v350 = m.T0[v349].(func(*base.Module, int32, int32) int32)(m, v339, v337)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+uint32(_c_F_pg_regcomp[1]))))
	if v345&int32(2) != 0 {
		v359 = v333
		v360 = v339
		v361 = v337
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v418 = v333
	v421 = v338
	goto L67
L74:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v350 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	v418 = v352
	v421 = v355
	goto L67
L76:
	;
	goto L77
L77:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_pg_regcomp[0]))
	v359 = v352
	v360 = v356
	v361 = v358
	goto L68
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+16)) = v411
	v414 = v359 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v414
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	if base.Ui32(v414) < base.Ui32(v416) {
		v333 = v414
		v337 = v361
		v338 = v416
		goto L65
	} else {
		goto L94
	}
L79:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v408 & int32(-8)
	goto L78
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+24)) = int32(101)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	if v404 != 0 {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v399 | int32(32)
	goto L78
L82:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v394&int32(-193) | int32(128)
	goto L78
L83:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v391 & int32(-33)
	goto L78
L84:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v388 & int32(-193)
	goto L78
L85:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v383&int32(-8) | int32(4)
	goto L78
L86:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v378&int32(-193) | int32(64)
	goto L78
L87:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v375 | int32(192)
	goto L78
L88:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v372 | int32(8)
	goto L78
L89:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v367&int32(-8) | int32(1)
	goto L78
L90:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v411 = v364 & int32(-9)
	goto L78
L91:
	;
	v406 = v404
	goto L93
L92:
	;
	v406 = int32(18)
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+12)) = v406
	goto L33
L94:
	;
	v418 = v414
	v421 = v416
	goto L67
L95:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	if v423 == int32(41) {
		goto L62
	} else {
		goto L96
	}
L96:
	;
	goto L63
L97:
	;
	v438 = v436
	goto L99
L98:
	;
	v438 = int32(18)
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+12)) = v438
	goto L33
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+16)) = v443 & int32(-225)
	goto L33
L101:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	if v462&int32(4) != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L103
L103:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	if v478&int32(192) != 0 {
		goto L112
	} else {
		goto L113
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+24)) = int32(110)
	v475 = F_next(m, v222)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L111
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+32)) = int32(3)
	goto L104
L106:
	;
	goto L107
L107:
	;
	if v462&int32(1) != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+32)) = int32(1)
	goto L104
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+32)) = int32(2)
	goto L104
L111:
	;
	goto L103
L112:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	v483 = F_subcolor(m, v481, int32(10))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v490 != 0 {
		v673 = v490
		goto L10
	} else {
		goto L117
	}
L115:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+100)) = uint16(v483)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	F_okcolors(m, v486, v487)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v492 = v11 + int32(4)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v495)+8))
	v498 = F_parse(m, v492, int32(101), int32(112), v496, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v498
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v501 != 0 {
		v673 = v501
		goto L10
	} else {
		goto L119
	}
L119:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	F_specialcolors(m, v502)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v505 != 0 {
		v673 = v505
		goto L10
	} else {
		goto L121
	}
L121:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	if v506&int32(16) != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	F_removecaptures(m, v492, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v512 = int32(1)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v513)+4)) = v512
	v518 = int32(2)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v513)+20))
	if v519 != 0 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L124
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+120)) = v527
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+1)))
	v532 = v530 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v529)+1)) = uint8(v532)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v529)+20))
	if v534 != 0 {
		goto L134
	} else {
		goto L135
	}
L127:
	;
	v521 = v519
	v522 = v518
	goto L130
L128:
	;
	v527 = v518
	goto L129
L129:
	;
	goto L126
L130:
	;
	v523 = F_numst(m, v521, v522)
	mBase = m.M
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v521)+24))
	if v524 != 0 {
		v521 = v524
		v522 = v523
		goto L130
	} else {
		goto L132
	}
L131:
	;
	v527 = v523
	goto L129
L132:
	;
	goto L131
L133:
	;
	v540 = v11 + int32(4)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+108))
	if v541 != 0 {
		goto L140
	} else {
		goto L141
	}
L134:
	;
	v535 = v534
	goto L137
L135:
	;
	goto L136
L136:
	;
	goto L133
L137:
	;
	F_markst(m, v535)
	mBase = m.M
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v535)+24))
	if v537 != 0 {
		v535 = v537
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L136
L139:
	;
	goto L138
L140:
	;
	v543 = v541
	goto L143
L141:
	;
	goto L142
L142:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v540)+108)) = int64(0)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v571 = F_nfatree(m, v11+int32(4), v570)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L4
	} else {
		goto L150
	}
L143:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v543)+84))
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+1)))
	if v551&int32(64) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L142
L145:
	;
	F_pfree(m, v543)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	if v550 != 0 {
		v543 = v550
		goto L143
	} else {
		goto L149
	}
L148:
	;
	goto L147
L149:
	;
	goto L144
L150:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v571 | v573
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v576 != 0 {
		v673 = v576
		goto L10
	} else {
		goto L151
	}
L151:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v11)+136))
	if int32(2) <= v577 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v586 = v512
	goto L155
L153:
	;
	goto L154
L154:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+1)))
	if v615&int32(2) != 0 {
		goto L160
	} else {
		goto L161
	}
L155:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	v593 = v590 + v586*int32(88)
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+2)))
	v599 = F_nfanode(m, v11+int32(4), v593, base.B2i32(v594&int32(2) == int32(0)))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L157
	}
L156:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v605 != 0 {
		v673 = v605
		goto L10
	} else {
		goto L159
	}
L157:
	;
	v602 = v586 + int32(1)
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v11)+136))
	if v602 < v603 {
		v586 = v602
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	goto L154
L160:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v618)+8)) = v619 | int32(_a_F_pg_regcomp_4)
	goto L162
L161:
	;
	goto L162
L162:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v625 = F_optimize(m, v624)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v627 != 0 {
		v673 = v627
		goto L10
	} else {
		goto L164
	}
L164:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	F_makesearch(m, v11+int32(4), v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v633 != 0 {
		v673 = v633
		goto L10
	} else {
		goto L166
	}
L166:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	F_compact(m, v634, v98+int32(20))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v639 != 0 {
		v673 = v639
		goto L10
	} else {
		goto L168
	}
L168:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v640
	v642 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = int32(_a_F_pg_regcomp_5)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v647
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v649
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+12)) = v651
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v653
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v642
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+68)) = v657
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v661&int32(8) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v664 = int32(1033)
	goto L171
L170:
	;
	v664 = int32(1034)
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+420)) = v664
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+424)) = v666
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = int32(0)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v11)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+428)) = v670
	v673 = v642
	goto L10
L172:
	;
	F_rfree(m, v682)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L4
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v681)+40))
	if v685 != v11+int32(52) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	goto L174
L176:
	;
	F_pfree(m, v685)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v681)+88))
	if v691 != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L178
L180:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v691)+36))
	if v692 != 0 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	goto L182
L182:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v681)+104))
	if v766 != 0 {
		goto L198
	} else {
		goto L199
	}
L183:
	;
	v693 = v692
	goto L186
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v691)+36)) = int32(0)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v691)+40))
	if v723 != 0 {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v693)))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v691)+76))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+136))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v693)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+136)) = v703 + v704*int32(-36) - int32(8)
	F_pfree(m, v693)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L188
	}
L187:
	;
	goto L185
L188:
	;
	if v701 != 0 {
		v693 = v701
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v724 = v723
	goto L193
L191:
	;
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v691)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v691)+40)) = int32(0)
	F_pfree(m, v691)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L4
	} else {
		goto L197
	}
L193:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v724)))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v691)+76))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)+136))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v724)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v733)+136)) = v734 + v735*int32(-40) - int32(8)
	F_pfree(m, v724)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L4
	} else {
		goto L195
	}
L194:
	;
	goto L192
L195:
	;
	if v732 != 0 {
		v724 = v732
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	goto L182
L198:
	;
	F_freesubre(m, v681, v766)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L4
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v681)+108))
	if v769 != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	goto L200
L202:
	;
	v770 = v769
	goto L205
L203:
	;
	goto L204
L204:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v681)+120))
	if v796 != 0 {
		goto L212
	} else {
		goto L213
	}
L205:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v770)+84))
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770)+1)))
	if v779&int32(64) == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v681)+108)) = int64(0)
	goto L204
L207:
	;
	F_pfree(m, v770)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L4
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	if v778 != 0 {
		v770 = v778
		goto L205
	} else {
		goto L211
	}
L210:
	;
	goto L209
L211:
	;
	goto L206
L212:
	;
	F_pfree(m, v796)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L4
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v681)+124))
	if v799 != 0 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	goto L214
L216:
	;
	F_pfree(m, v799)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L4
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v681)+128))
	if v802 != 0 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	goto L218
L220:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v681)+132))
	v805 = v803 - int32(1)
	if int32(0) < v805 {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v681)+24)) = int32(101)
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v681)+12))
	if v854 != 0 {
		goto L236
	} else {
		goto L237
	}
L223:
	;
	v808 = v802
	v809 = v805
	goto L226
L224:
	;
	goto L225
L225:
	;
	F_pfree(m, v802)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L235
	}
L226:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v808)+124))
	if v816 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	goto L225
L228:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v808)+152))
	F_pfree(m, v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L4
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v830 = int32(1)
	if v830 < v809 {
		v808 = v808 + int32(88)
		v809 = v809 - v830
		goto L226
	} else {
		goto L234
	}
L231:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v808)+156))
	F_pfree(m, v820)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L4
	} else {
		goto L232
	}
L232:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v808)+160))
	F_pfree(m, v823)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v808)+124)) = int32(0)
	goto L230
L234:
	;
	goto L227
L235:
	;
	goto L222
L236:
	;
	v855 = v854
	goto L238
L237:
	;
	v855 = v673
	goto L238
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v681)+12)) = v855
	v863 = v855
	goto L3
}
func F_pg_relpagesbyid_v1_5(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_relation_open(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_pg_relpages_impl(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return v8
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
func F_pg_sleep(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v10 float64
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v69 int64
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v5&int64(9223372036854775807)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v10 = base.F64_reinterpret_i64(v5)
	if base.F64_le(v10, float64(0)) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = m.G0
	v17 = int32(16)
	v18 = v16 - v17
	m.G0 = v18
	F_gettimeofday(m, v18)
	mBase = m.M
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+8)))
	m.G0 = v18 + v17
	goto L4
L4:
	;
	v33 = base.F64_ceil(base.F64_mul(v10, float64(1e+06)))
	v34 = float64(4.611686018427388e+18)
	if base.F64_lt(v33, v34) != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v37 = v33
	goto L7
L6:
	;
	v37 = v34
	goto L7
L7:
	;
	goto L8
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sleep[0]))
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v54 = m.G0
	v55 = int32(16)
	v56 = v54 - v55
	m.G0 = v56
	F_gettimeofday(m, v56)
	mBase = m.M
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v56)))
	v60 = int64(*(*int32)(unsafe.Add(mBase, uint32(v56)+8)))
	m.G0 = v56 + v55
	goto L15
L13:
	;
	return int64(0)
L14:
	;
	goto L12
L15:
	;
	v69 = v22 + v21*int64(1000000) - int64(946684800000000) + base.I64_trunc_sat_f64_s(v37) - (v60 + v59*int64(1000000) - int64(946684800000000))
	if v69 <= int64(599999999) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v69 <= int64(0) {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v79 = int32(_a_F_pg_sleep_0)
	goto L18
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sleep[1]))
	v84 = F_WaitLatch(m, v81, int32(41), v79, int32(150994947))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v78 = base.I32_div_u_s(base.I32_wrap_i64(v69)+int32(999), int32(1000))
	v79 = v78
	goto L18
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sleep[1]))
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v88
	v93 = base.AtomicRmwOr32(m, v88, int32(_a_F_pg_sleep_1), v88)
	goto L21
L21:
	;
	goto L8
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v7 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v83
L2:
	;
	v10 = base.B2i32(base.Ui32(l1) < base.Ui32(l3))
	if base.Ui32(l1) < base.Ui32(l3) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v78 = m.T0[v77].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L27
	} else {
		goto L28
	}
L5:
	;
	v11 = l1
	goto L7
L6:
	;
	v11 = l3
	goto L7
L7:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v11) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v73 != 0 {
		v83 = v73
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v73 = int32(0)
	goto L8
L10:
	;
	v47 = v42
	v48 = v43
	v49 = v44
	goto L20
L11:
	;
	if (l0|l2)&int32(3) != 0 {
		v42 = l0
		v43 = l2
		v44 = v11
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v35 = l0
	v36 = l2
	v37 = v11
	goto L13
L13:
	;
	if v37 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v19 = l0
	v20 = l2
	v21 = v11
	goto L15
L15:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v24 != v25 {
		v42 = v19
		v43 = v20
		v44 = v21
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v35 = v30
	v36 = v28
	v37 = v32
	goto L13
L17:
	;
	v27 = int32(4)
	v28 = v20 + v27
	v30 = v19 + v27
	v32 = v21 - v27
	if base.Ui32(int32(3)) < base.Ui32(v32) {
		v19 = v30
		v20 = v28
		v21 = v32
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v42 = v35
	v43 = v36
	v44 = v37
	goto L10
L20:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v52 == v53 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v73 = v52 - v53
	goto L8
L22:
	;
	v55 = int32(1)
	v60 = v49 - v55
	if v60 != 0 {
		v47 = v47 + v55
		v48 = v48 + v55
		v49 = v60
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L9
L26:
	;
	return base.B2i32(base.Ui32(l3) < base.Ui32(l1)) - v10
L27:
	;
	return int32(0)
L28:
	;
	v83 = v78
	goto L1
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
func F_pg_switch_wal(m *base.Module, l0 int32) int64 {
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_switch_wal[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_switch_wal[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+308))
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
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_switch_wal_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					F_errhint(m, int32(_a_F_pg_switch_wal_1), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_switch_wal_2), int32(215), int32(_a_F_pg_switch_wal_3))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
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
			return int64(0)
		} else {
			return v38
		}
	}
}
func F_pg_tablespace_databases(m *base.Module, l0 int32) int64 {
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L39
	}
L4:
	;
	m.G0 = v10 + int32(80)
	return int64(0)
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
	F_errfinish(m, int32(_a_F_pg_tablespace_databases_5), int32(253), int32(_a_F_pg_tablespace_databases_6))
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
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_pg_tablespace_databases[0]))
	if v104 != int32(44) {
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
	v102 = m.ExcPending
	if v102 != 0 {
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
	v92 = F_ReadDir(m, v48, v47)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+71)) = uint8(v79)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = base.I64_extend_i32_u(v65)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v83, v84, v10+int32(72), v10+int32(71))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L23
L31:
	;
	if v92 != 0 {
		v52 = v92
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
	v109 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v109 == int32(0) {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	F_errmsg(m, int32(_a_F_pg_tablespace_databases_8), v10)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_pg_tablespace_databases_5), int32(275), int32(_a_F_pg_tablespace_databases_6))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
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
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v47
	F_errmsg(m, int32(_a_F_pg_tablespace_databases_9), v10+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_pg_tablespace_databases_5), int32(273), int32(_a_F_pg_tablespace_databases_6))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
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
func F_pg_u_isalnum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
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
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v107
L2:
	;
	v107 = base.B2i32(v96&int32(255) == int32(9))
	goto L1
L3:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_u_isalnum[0]))))
	v96 = v89
	goto L2
L4:
	;
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
L5:
	;
	v12 = int32(1201)
	v13 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v68 = int32(1)
	v70 = l0 << (uint(v68) % 32)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_u_isalnum[1]))))
	if v71&v68 != 0 {
		v107 = v68
		goto L1
	} else {
		goto L28
	}
L8:
	;
	v18 = base.I32_div_s(v12+v13, int32(2))
	v20 = v18 << (uint(int32(3)) % 32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_pg_u_isalnum[2])))
	if base.Ui32(v23) < base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L18
	}
L10:
	;
	if v36 <= v35 {
		v12 = v35
		v13 = v36
		goto L8
	} else {
		goto L17
	}
L11:
	;
	v35 = v12
	v36 = v18 + int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_pg_u_isalnum[3])))
	if base.Ui32(v29) <= base.Ui32(l0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(1)
L15:
	;
	goto L16
L16:
	;
	v35 = v18 - int32(1)
	v36 = v13
	goto L10
L17:
	;
	goto L9
L18:
	;
	v42 = int32(3408)
	v43 = int32(0)
	goto L20
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_pg_u_isalnum[4]))))
	v96 = v67
	goto L2
L20:
	;
	v48 = base.I32_div_s(v42+v43, int32(2))
	v50 = v48 * int32(12)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_pg_u_isalnum[5])))
	if base.Ui32(v53) < base.Ui32(l0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v96 = int32(0)
	goto L2
L22:
	;
	if v64 <= v63 {
		v42 = v63
		v43 = v64
		goto L20
	} else {
		goto L27
	}
L23:
	;
	v63 = v42
	v64 = v48 + int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_pg_u_isalnum[6])))
	if base.Ui32(v59) <= base.Ui32(l0) {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v63 = v48 - int32(1)
	v64 = v43
	goto L22
L27:
	;
	goto L21
L28:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L4
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
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
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
					v223 = m.ExcPending
					if v223 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v226 = m.ExcPending
						if v226 != 0 {
							return
						} else {
							v228 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server[0]))
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
							v231 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v231
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v229
							F_errmsg(m, int32(_a_F_pg_unicode_to_server_2), v8)
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_pg_unicode_to_server_3), int32(913), int32(_a_F_pg_unicode_to_server_4))
								mBase = m.M
								v241 = m.ExcPending
								if v241 != 0 {
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
					v196 = F_FunctionCall6Coll(m, v189, int64(6), base.I64_extend_i32_s(v21), base.I64_extend_i32_u(v159), base.I64_extend_i32_u(l1), base.I64_extend_i32_s(v184), int64(0))
					mBase = m.M
					v197 = m.ExcPending
					if v197 != 0 {
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
		v207 = m.ExcPending
		if v207 != 0 {
			return
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v210 = m.ExcPending
			if v210 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_pg_unicode_to_server_5), int32(0))
				mBase = m.M
				v214 = m.ExcPending
				if v214 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_pg_unicode_to_server_3), int32(888), int32(_a_F_pg_unicode_to_server_4))
					mBase = m.M
					v219 = m.ExcPending
					if v219 != 0 {
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
	var v198 int64
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	v1 = l0
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.Ui32(int32(_a_F_pg_unicode_to_server_noerror_0)) < base.Ui32(v1-int32(1)) {
		v205 = v3
		m.G0 = v8 + int32(16)
		return v205
	} else {
		if base.Ui32(v1) <= base.Ui32(int32(127)) {
			v16 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v16)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v1)
			v205 = int32(1)
			m.G0 = v8 + int32(16)
			return v205
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
				v205 = int32(1)
				m.G0 = v8 + int32(16)
				return v205
			} else {
				v105 = *(*int32)(unsafe.Add(mBase, _c_F_pg_unicode_to_server_noerror[1]))
				if v105 == int32(0) {
					v205 = v3
					m.G0 = v8 + int32(16)
					return v205
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
					v198 = F_FunctionCall6Coll(m, v191, int64(6), base.I64_extend_i32_s(v22), base.I64_extend_i32_u(v161), base.I64_extend_i32_u(l1), base.I64_extend_i32_s(v186), int64(1))
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return int32(0)
					} else {
						v205 = base.B2i32(v186 == base.I32_wrap_i64(v198))
						m.G0 = v8 + int32(16)
						return v205
					}
				}
			}
		}
	}
}
func F_pg_visibility_map(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v2
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)) = uint16(v2)
	v18 = F_relation_open(m, v12, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L44
	}
L2:
	;
	return int64(0)
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+119)))
	v25 = v23 - int32(109)
	v32 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v25))|base.B2i32(int32(1)<<(uint(v25)%32)&int32(161) == v32) == v32 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if base.Ui64(int64(4294967295)) <= base.Ui64(v11) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L39
	}
L7:
	;
	v40 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v40, int32(1), int32(_a_F_pg_visibility_map_0), int32(16), int32(-1), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v40, int32(2), int32(_a_F_pg_visibility_map_1), int32(16), int32(-1), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v56 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v56 < v65 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v143 = F_BlessTupleDesc(m, v40)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L30
	}
L12:
	;
	v69 = v40 + int32(28)
	v76 = v56
	v77 = v65
	v79 = v56
	goto L16
L13:
	;
	v133 = v56
	v140 = v65
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v133
	goto L11
L15:
	;
	v133 = v127
	v140 = v106
	goto L14
L16:
	;
	v85 = v69 + v65<<(uint(int32(3))%32) + v76*int32(100)
	v88 = v69 + v76<<(uint(int32(3))%32)
	if v65 != v77 {
		v106 = v77
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v127 = v65
	goto L15
L18:
	;
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+2)))
	if v107 <= int32(0) {
		v127 = v76
		goto L15
	} else {
		goto L26
	}
L19:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+7)))
	if v90 != int32(118) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v106 = v76
	goto L18
L21:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
	if v93 != int32(1) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+6)))
	if v96&int32(6) != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+2)))
	if v99 <= int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+90)))
	if v102 != int32(118) {
		v106 = v65
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L20
L26:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+90)))
	if v110 == int32(118) {
		v127 = v76
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+5)))
	v119 = (v79 + v113 - int32(1)) & (int32(0) - v113)
	if int32(_a_F_pg_visibility_map_2) < v119 {
		v127 = v76
		goto L15
	} else {
		goto L28
	}
L28:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v88))) = uint16(v119)
	v125 = v76 + int32(1)
	if v125 != v65 {
		v76 = v125
		v77 = v106
		v79 = v119 + v107
		goto L16
	} else {
		goto L29
	}
L29:
	;
	goto L17
L30:
	;
	v148 = F_visibilitymap_get_status(m, v18, base.I32_wrap_i64(v11), v9+int32(44))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	if v150 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_ReleaseBuffer(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v153 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = base.I64_extend_i32_u(v148 & v153)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = base.I64_extend_i32_u(int32(base.Ui32(v148)>>(uint(v153)%32)) & v153)
	F_relation_close(m, v18, v153)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v170 = F_heap_form_tuple(m, v143, v9+int32(16), v9+int32(14))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170)+16))
	v173 = F_HeapTupleHeaderGetDatum(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	m.G0 = v9 + int32(48)
	return v173
L39:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v186 + int32(4)
	F_errmsg(m, int32(_a_F_pg_visibility_map_3), v9)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v194 = int32(*(*int8)(unsafe.Add(mBase, uint32(v193)+119)))
	F_errdetail_relkind_not_supported(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_pg_visibility_map_4), int32(932), int32(_a_F_pg_visibility_map_5))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_pg_visibility_map_6), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_visibility_map_4), int32(103), int32(_a_F_pg_visibility_map_7))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_visibility_map_rel(m *base.Module, l0 int32) int64 {
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int64
	_ = v203
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	goto L32
L4:
	;
	return int64(0)
L5:
	;
	v21 = int32(_a_F_pg_visibility_map_rel_0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_map_rel[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_map_rel[0])) = v24
	v27 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v27, int32(1), int32(_a_F_pg_visibility_map_rel_1), int32(20), int32(-1), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v27, int32(2), int32(_a_F_pg_visibility_map_rel_2), int32(16), int32(-1), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v27, int32(3), int32(_a_F_pg_visibility_map_rel_3), int32(16), int32(-1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v50 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v50 < v59 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v137 = F_BlessTupleDesc(m, v27)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L29
	}
L11:
	;
	v63 = v27 + int32(28)
	v70 = v50
	v71 = v59
	v73 = v50
	goto L15
L12:
	;
	v127 = v50
	v134 = v59
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v127
	goto L10
L14:
	;
	v127 = v121
	v134 = v100
	goto L13
L15:
	;
	v79 = v63 + v59<<(uint(int32(3))%32) + v70*int32(100)
	v82 = v63 + v70<<(uint(int32(3))%32)
	if v59 != v71 {
		v100 = v71
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v121 = v59
	goto L14
L17:
	;
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+2)))
	if v101 <= int32(0) {
		v121 = v70
		goto L14
	} else {
		goto L25
	}
L18:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+7)))
	if v84 != int32(118) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v100 = v70
	goto L17
L20:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	if v87 != int32(1) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+6)))
	if v90&int32(6) != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+2)))
	if v93 <= int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+90)))
	if v96 != int32(118) {
		v100 = v59
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L19
L25:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+90)))
	if v104 == int32(118) {
		v121 = v70
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+5)))
	v113 = (v73 + v107 - int32(1)) & (int32(0) - v107)
	if int32(_a_F_pg_visibility_map_rel_4) < v113 {
		v121 = v70
		goto L14
	} else {
		goto L27
	}
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v82))) = uint16(v113)
	v119 = v70 + int32(1)
	if v119 != v59 {
		v70 = v119
		v71 = v100
		v73 = v113 + v101
		goto L15
	} else {
		goto L28
	}
L28:
	;
	goto L16
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v137
	v141 = F_collect_visibility_data(m, v16, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v141
	*(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_map_rel[0])) = v22
	goto L3
L31:
	;
	m.G0 = v10 + int32(48)
	return v203
L32:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if base.Ui32(v153) < base.Ui32(v154) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v156)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v156)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = base.I64_extend_i32_u(v153)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v153)+8)))
	v164 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = base.I64_extend_i32_u(v163 & v164)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(int32(base.Ui32(v163)>>(uint(v164)%32)) & v164)
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v153 + v164
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v151)+28))
	v182 = F_heap_form_tuple(m, v177, v10+int32(16), v10+int32(12))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
	*(*int64)(unsafe.Add(mBase, uint32(v151))) = v184 + int64(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+20)) = int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	v192 = F_HeapTupleHeaderGetDatum(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v203 = v192
	goto L31
L38:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+20)) = int32(2)
	v199 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v199)
	v203 = int64(0)
	goto L31
}
func F_pg_visible_in_snapshot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v63 int64
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v11
	v18 = int64(1)
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	if base.Ui64(v11) < base.Ui64(v19) {
		v63 = v18
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(16)
	return v63
L4:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	if base.Ui64(v21) <= base.Ui64(v11) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v63 = int64(0)
	goto L3
L6:
	;
	v24 = v13 + int32(24)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if base.Ui32(v25) <= base.Ui32(int32(30)) {
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
	if v25 == int32(0) {
		v63 = v18
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v31 = int32(8)
	v35 = F_bsearch(m, v9+v31, v24, v25, v31, int32(1759))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v63 = base.I64_extend_i32_u(base.B2i32(v35 == int32(0)))
	goto L3
L13:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v24+v40<<(uint(int32(3))%32))))
	if v11 == v49 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v63 = v18
	goto L3
L15:
	;
	v52 = v40 + int32(1)
	if v25 != v52 {
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
