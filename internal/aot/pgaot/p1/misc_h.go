package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HandleConcurrentAbort(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if v3 == int32(0) {
		return
	} else {
		v6 = F_TransactionIdIsInProgress(m, v3)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			if v6 != 0 {
				return
			} else {
				v9 = *(*int32)(unsafe.Add(mBase, _consts[114]))
				v10 = F_TransactionIdDidCommit(m, v9)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					if v10 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v15 = m.ExcPending
						if v15 != 0 {
							return
						} else {
							F_errcode(m, int32(4))
							mBase = m.M
							v18 = m.ExcPending
							if v18 != 0 {
								return
							} else {
								F_errmsg(m, int32(274496), int32(0))
								mBase = m.M
								v22 = m.ExcPending
								if v22 != 0 {
									return
								} else {
									F_errfinish(m, int32(480409), int32(498), int32(78182))
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
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
			}
		}
	}
}
func F_HistoricSnapshotActive(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[655]))
	return base.B2i32(v2 != int32(0))
}
func F_handle_pm_pmsignal_signal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _consts[575])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_has_createrole_privilege(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	v4 = F_superuser_arg(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v11 = F_SearchSysCache1(m, int32(11), l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 == int32(0) {
					return int32(0)
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v18)+70)))
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v24 = v20
						return v24 & int32(1)
					}
				}
			}
		} else {
			v24 = int32(1)
			return v24 & int32(1)
		}
	}
}
func F_has_dangerous_join_using(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
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
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v12 - int32(63) {
	case 0:
		v104 = v3
		goto L2
	case 1:
		goto L3
	case 2:
		goto L4
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L11
	} else {
		goto L28
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return v104
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 == int32(0) {
		v104 = v3
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 <= v18 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v104 = v3
	goto L2
L7:
	;
	goto L8
L8:
	;
	v23 = v18
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v23<<(uint(int32(2))%32))))
	v34 = F_has_dangerous_join_using(m, l0, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v104 = v34
	goto L2
L11:
	;
	return int32(0)
L12:
	;
	if v34 != 0 {
		v104 = v34
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v39 = v23 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v39 < v40 {
		v23 = v39
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v104 = int32(1)
	goto L2
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v86 = F_has_dangerous_join_using(m, l0, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L11
	} else {
		goto L24
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v43 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48<<(uint(int32(2))%32)-int32(4))))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	if v55 <= int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v63 = int32(0)
	goto L20
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v59+v63<<(uint(int32(2))%32))))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72 != int32(6) {
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L16
L22:
	;
	v76 = v63 + int32(1)
	if v76 != v55 {
		v63 = v76
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v86 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v89 = F_has_dangerous_join_using(m, l0, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	if v89 == int32(0) {
		v104 = v3
		goto L2
	} else {
		goto L27
	}
L27:
	;
	goto L15
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v116
	F_errmsg_internal(m, int32(469568), v10)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(477182), int32(4189), int32(317104))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_has_largeobject_privilege_id_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = F_convert_any_priv_string(m, v9, int32(1624736))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v14&int64(4) == int64(0) {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[26]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v23 = v22
			} else {
				v23 = int32(0)
			}
			v24 = F_LargeObjectExistsWithSnapshot(m, v7, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 != 0 {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1052])))
					if v28 != 0 {
						v37 = int32(1)
						return v37
					} else {
						v29 = F_pg_largeobject_aclcheck_snapshot(m, v7, v6, v14, v23)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v29 == int32(0))
						}
					}
				} else {
					v34 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
					v37 = int32(0)
					return v37
				}
			}
		}
	}
}
func F_hashagg_spill_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64) {
	mBase := m.M
	_ = mBase
	var v17 float64
	_ = v17
	var v21 float64
	_ = v21
	var v23 int32
	_ = v23
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v31 float64
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v46 float64
	_ = v46
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	v17 = float64(1024)
	v21 = *(*float64)(unsafe.Add(mBase, _consts[433]))
	v23 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	v27 = base.F64_mul(base.F64_mul(v21, base.F64_convert_i32_s(v23)), v17)
	v28 = float64(4.294967295e+09)
	if base.F64_lt(v27, v28) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v68 = int32(1073741823)
	if v68 <= v66 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v40 = base.F64_convert_i32_u(v39)
	v46 = base.F64_mul(base.F64_add(base.F64_mul(v40, float64(0.25)), float64(-8192)), float64(0.0001220703125))
	v52 = base.F64_add(base.F64_div(base.F64_mul(base.F64_mul(l3, float64(1.5)), l4), v40), float64(1))
	if base.F64_gt(v52, v46) != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v31 = v27
	goto L5
L4:
	;
	v31 = v28
	goto L5
L5:
	;
	if base.F64_lt(v31, float64(4.294967296e+09))&base.F64_ge(v31, float64(0)) != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = base.I32_trunc_f64_u(v31)
	v39 = v37
	goto L2
L7:
	;
	goto L8
L8:
	;
	v39 = int32(0)
	goto L2
L9:
	;
	v54 = v46
	goto L11
L10:
	;
	v54 = v52
	goto L11
L11:
	;
	if base.F64_lt(v54, float64(4)) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v57 = float64(4)
	goto L14
L13:
	;
	v57 = v54
	goto L14
L14:
	;
	if base.F64_gt(v57, float64(1024)) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v60 = v17
	goto L17
L16:
	;
	v60 = v57
	goto L17
L17:
	;
	if base.F64_lt(base.F64_abs(v60), float64(2.147483648e+09)) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v64 = base.I32_trunc_f64_s(v60)
	v66 = v64
	goto L1
L19:
	;
	goto L20
L20:
	;
	v66 = int32(-2147483648)
	goto L1
L21:
	;
	if int32(31) < l2+v79 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	v71 = v68
	goto L24
L23:
	;
	v71 = v66
	goto L24
L24:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v71) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v79 = int32(32) - base.I32_clz(v71-int32(1))
	goto L27
L26:
	;
	v79 = int32(0)
	goto L27
L27:
	;
	goto L21
L28:
	;
	v83 = int32(32) - l2
	goto L30
L29:
	;
	v83 = v79
	goto L30
L30:
	;
	v85 = F_palloc0(m, int32(4)<<(uint(v83)%32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
	v90 = F_palloc0(m, int32(8)<<(uint(v83)%32))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v90
	v95 = F_palloc0(m, int32(24)<<(uint(v83)%32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v95
	v99 = int32(1) << (uint(v83) % 32)
	v101 = base.B2i32(v83 == int32(31))
	if v101 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v104 = int32(1)
	if v99 <= v104 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v99
	v146 = int32(32)
	v148 = v146 - (l2 + v83)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v148
	v150 = int32(0)
	if v148 < v146 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v107 = v104
	goto L40
L39:
	;
	v107 = v99
	goto L40
L40:
	;
	v114 = int32(0)
	goto L41
L41:
	;
	v122 = F_LogicalTapeCreate(m, l1)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L31
	} else {
		goto L43
	}
L42:
	;
	goto L37
L43:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v124+v114<<(uint(int32(2))%32)))) = v122
	v130 = v114 + int32(1)
	if v130 != v107 {
		v114 = v130
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v157 = (v99 - int32(1)) << (uint(v148) % 32)
	goto L47
L46:
	;
	v157 = v150
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v157
	if v101 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v161 = int32(1)
	if v99 <= v161 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	return
L51:
	;
	v164 = v161
	goto L53
L52:
	;
	v164 = v99
	goto L53
L53:
	;
	v170 = v150
	goto L54
L54:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_initHyperLogLog(m, v178+v170*int32(24), int32(5))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L31
	} else {
		goto L56
	}
L55:
	;
	goto L50
L56:
	;
	v186 = v170 + int32(1)
	if v186 != v164 {
		v170 = v186
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
}
func F_hashbpcharextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
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
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int64
	_ = v412
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
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
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L138
	}
L4:
	;
	v15 = int32(1)
	v16 = v10 + v15
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v21 = v19 & v15
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L133
	}
L7:
	;
	v22 = v15
	goto L9
L8:
	;
	v22 = int32(4)
	goto L9
L9:
	;
	if v19 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v51 = v10 + v22
	if v21 != 0 {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v25 = int32(4)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v27&int32(254) == int32(2) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v40 = int32(1)
	if v21 != 0 {
		v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v36 = v25
	goto L16
L15:
	;
	v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
	goto L16
L16:
	;
	if v27 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v39 = v25
	goto L19
L18:
	;
	v39 = v36
	goto L19
L19:
	;
	v50 = v39
	goto L10
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	v54 = v16
	goto L23
L22:
	;
	v54 = v10 + int32(4)
	goto L23
L23:
	;
	v59 = v50
	goto L24
L24:
	;
	if v59 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v76 = F_pg_newlocale_from_collation(m, v14)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	v75 = v50 >> (uint(int32(31)) % 32) & v50
	goto L26
L28:
	;
	goto L29
L29:
	;
	v69 = v59 - int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v69))))
	if v71 == int32(32) {
		v59 = v69
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v75 = v59
	goto L26
L31:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v736 != v10 {
		goto L129
	} else {
		goto L130
	}
L32:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v78 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	v88 = v75 - int32(1636608432)
	if v82 == int64(0) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	goto L35
L35:
	;
	v400 = int32(0)
	v402 = F_pg_strnxfrm(m, v400, v400, v51, v75, v76)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L80
	}
L36:
	;
	v398 = F_Int64GetDatum(m, base.I64_extend_i32_u(v388)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v380^v388-base.I32_rotl(v388, int32(24))))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L79
	}
L37:
	;
	if v51&int32(3) != 0 {
		goto L53
	} else {
		goto L54
	}
L38:
	;
	v125 = v88
	v127 = v88
	v129 = v88
	goto L37
L39:
	;
	goto L40
L40:
	;
	v92 = v88 + base.I32_wrap_i64(v82)
	v93 = v92 + v88
	v97 = int32(4)
	v99 = base.I32_wrap_i64(int64(base.Ui64(v82)>>(uint(int64(32))%64))) ^ base.I32_rotl(v88, v97)
	v103 = v92 - v99 ^ base.I32_rotl(v99, int32(6))
	v107 = v93 - v103 ^ base.I32_rotl(v103, int32(8))
	v108 = v99 + v93
	v109 = v103 + v108
	v110 = v107 + v109
	v114 = v108 - v107 ^ base.I32_rotl(v107, int32(16))
	v118 = v109 - v114 ^ base.I32_rotl(v114, int32(19))
	v123 = v114 + v110
	v125 = v123
	v127 = v110 - v118 ^ base.I32_rotl(v118, v97)
	v129 = v118 + v123
	goto L37
L41:
	;
	v366 = int32(14)
	v368 = v362 ^ v363 - base.I32_rotl(v362, v366)
	v372 = v368 ^ v361 - base.I32_rotl(v368, int32(11))
	v376 = v372 ^ v362 - base.I32_rotl(v372, int32(25))
	v380 = v376 ^ v368 - base.I32_rotl(v376, int32(16))
	v384 = v380 ^ v372 - base.I32_rotl(v380, int32(4))
	v388 = v384 ^ v376 - base.I32_rotl(v384, v366)
	goto L36
L42:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v361 = v353 + v356
	v362 = v354
	v363 = v355
	goto L41
L43:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	v353 = v349<<(uint(int32(8))%32) + v346
	v354 = v347
	v355 = v348
	goto L42
L44:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+2)))
	v346 = v342<<(uint(int32(16))%32) + v339
	v347 = v340
	v348 = v341
	goto L43
L45:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+3)))
	v339 = v335<<(uint(int32(24))%32) + v186
	v340 = v333
	v341 = v334
	goto L44
L46:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
	v333 = v329 + v331
	v334 = v330
	goto L45
L47:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+5)))
	v329 = v325<<(uint(int32(8))%32) + v323
	v330 = v324
	goto L46
L48:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+6)))
	v323 = v319<<(uint(int32(16))%32) + v317
	v324 = v318
	goto L47
L49:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+7)))
	v317 = v313<<(uint(int32(24))%32) + v187
	v318 = v312
	goto L48
L50:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+8)))
	v312 = v308<<(uint(int32(8))%32) + v307
	goto L49
L51:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+9)))
	v307 = v303<<(uint(int32(16))%32) + v302
	goto L50
L52:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+10)))
	v302 = v298<<(uint(int32(24))%32) + v188
	goto L51
L53:
	;
	if base.Ui32(int32(11)) < base.Ui32(v75) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v75) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	v134 = v51
	v135 = v75
	v137 = v125
	v138 = v129
	v139 = v127
	goto L59
L57:
	;
	v183 = v51
	v184 = v75
	v186 = v125
	v187 = v129
	v188 = v127
	goto L58
L58:
	;
	switch v184 - int32(1) {
	case 0:
		v353 = v186
		v354 = v187
		v355 = v188
		goto L42
	case 1:
		v346 = v186
		v347 = v187
		v348 = v188
		goto L43
	case 2:
		v339 = v186
		v340 = v187
		v341 = v188
		goto L44
	case 3:
		v333 = v187
		v334 = v188
		goto L45
	case 4:
		v329 = v187
		v330 = v188
		goto L46
	case 5:
		v323 = v187
		v324 = v188
		goto L47
	case 6:
		v317 = v187
		v318 = v188
		goto L48
	case 7:
		v312 = v188
		goto L49
	case 8:
		v307 = v188
		goto L50
	case 9:
		v302 = v188
		goto L51
	case 10:
		goto L52
	default:
		v361 = v186
		v362 = v187
		v363 = v188
		goto L41
	}
L59:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	v142 = v141 + v138
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v146 = v145 + v139
	v148 = int32(4)
	v150 = v143 + v137 - v146 ^ base.I32_rotl(v146, v148)
	v154 = v142 - v150 ^ base.I32_rotl(v150, int32(6))
	v155 = v146 + v142
	v156 = v150 + v155
	v157 = v154 + v156
	v161 = v155 - v154 ^ base.I32_rotl(v154, int32(8))
	v165 = v156 - v161 ^ base.I32_rotl(v161, int32(16))
	v169 = v157 - v165 ^ base.I32_rotl(v165, int32(19))
	v170 = v161 + v157
	v171 = v165 + v170
	v172 = v169 + v171
	v176 = v170 - v169 ^ base.I32_rotl(v169, v148)
	v177 = int32(12)
	v178 = v134 + v177
	v180 = v135 - v177
	if base.Ui32(int32(11)) < base.Ui32(v180) {
		v134 = v178
		v135 = v180
		v137 = v171
		v138 = v172
		v139 = v176
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v183 = v178
	v184 = v180
	v186 = v171
	v187 = v172
	v188 = v176
	goto L58
L61:
	;
	goto L60
L62:
	;
	v194 = v51
	v195 = v75
	v197 = v125
	v198 = v129
	v199 = v127
	goto L65
L63:
	;
	v243 = v51
	v244 = v75
	v246 = v125
	v247 = v129
	v248 = v127
	goto L64
L64:
	;
	switch v244 - int32(1) {
	case 0:
		v295 = v246
		goto L68
	case 1:
		v290 = v246
		goto L69
	case 2:
		goto L70
	case 3:
		v283 = v247
		goto L71
	case 4:
		v280 = v247
		goto L72
	case 5:
		v275 = v247
		goto L73
	case 6:
		goto L74
	case 7:
		v266 = v248
		goto L75
	case 8:
		v261 = v248
		goto L76
	case 9:
		v256 = v248
		goto L77
	case 10:
		goto L78
	default:
		v361 = v246
		v362 = v247
		v363 = v248
		goto L41
	}
L65:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	v202 = v201 + v198
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	v206 = v205 + v199
	v208 = int32(4)
	v210 = v203 + v197 - v206 ^ base.I32_rotl(v206, v208)
	v214 = v202 - v210 ^ base.I32_rotl(v210, int32(6))
	v215 = v206 + v202
	v216 = v210 + v215
	v217 = v214 + v216
	v221 = v215 - v214 ^ base.I32_rotl(v214, int32(8))
	v225 = v216 - v221 ^ base.I32_rotl(v221, int32(16))
	v229 = v217 - v225 ^ base.I32_rotl(v225, int32(19))
	v230 = v221 + v217
	v231 = v225 + v230
	v232 = v229 + v231
	v236 = v230 - v229 ^ base.I32_rotl(v229, v208)
	v237 = int32(12)
	v238 = v194 + v237
	v240 = v195 - v237
	if base.Ui32(int32(11)) < base.Ui32(v240) {
		v194 = v238
		v195 = v240
		v197 = v231
		v198 = v232
		v199 = v236
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v243 = v238
	v244 = v240
	v246 = v231
	v247 = v232
	v248 = v236
	goto L64
L67:
	;
	goto L66
L68:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v361 = v295 + v296
	v362 = v247
	v363 = v248
	goto L41
L69:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)))
	v295 = v291<<(uint(int32(8))%32) + v290
	goto L68
L70:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+2)))
	v290 = v286<<(uint(int32(16))%32) + v246
	goto L69
L71:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v361 = v284 + v246
	v362 = v283
	v363 = v248
	goto L41
L72:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+4)))
	v283 = v280 + v281
	goto L71
L73:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+5)))
	v280 = v276<<(uint(int32(8))%32) + v275
	goto L72
L74:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+6)))
	v275 = v271<<(uint(int32(16))%32) + v247
	goto L73
L75:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v361 = v267 + v246
	v362 = v269 + v247
	v363 = v266
	goto L41
L76:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+8)))
	v266 = v262<<(uint(int32(8))%32) + v261
	goto L75
L77:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+9)))
	v261 = v257<<(uint(int32(16))%32) + v256
	goto L76
L78:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+10)))
	v256 = v252<<(uint(int32(24))%32) + v248
	goto L77
L79:
	;
	v732 = v398
	goto L31
L80:
	;
	v405 = v402 + int32(1)
	v406 = F_palloc(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v408 = F_pg_strnxfrm(m, v406, v405, v51, v75, v76)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if base.Ui32(v402) < base.Ui32(v408) {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v411)))
	v418 = v405 - int32(1636608432)
	if v412 == int64(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v728 = F_Int64GetDatum(m, base.I64_extend_i32_u(v718)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v710^v718-base.I32_rotl(v718, int32(24))))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L127
	}
L85:
	;
	if v406&int32(3) != 0 {
		goto L101
	} else {
		goto L102
	}
L86:
	;
	v455 = v418
	v457 = v418
	v459 = v418
	goto L85
L87:
	;
	goto L88
L88:
	;
	v422 = v418 + base.I32_wrap_i64(v412)
	v423 = v422 + v418
	v427 = int32(4)
	v429 = base.I32_wrap_i64(int64(base.Ui64(v412)>>(uint(int64(32))%64))) ^ base.I32_rotl(v418, v427)
	v433 = v422 - v429 ^ base.I32_rotl(v429, int32(6))
	v437 = v423 - v433 ^ base.I32_rotl(v433, int32(8))
	v438 = v429 + v423
	v439 = v433 + v438
	v440 = v437 + v439
	v444 = v438 - v437 ^ base.I32_rotl(v437, int32(16))
	v448 = v439 - v444 ^ base.I32_rotl(v444, int32(19))
	v453 = v444 + v440
	v455 = v453
	v457 = v440 - v448 ^ base.I32_rotl(v448, v427)
	v459 = v448 + v453
	goto L85
L89:
	;
	v696 = int32(14)
	v698 = v692 ^ v693 - base.I32_rotl(v692, v696)
	v702 = v698 ^ v691 - base.I32_rotl(v698, int32(11))
	v706 = v702 ^ v692 - base.I32_rotl(v702, int32(25))
	v710 = v706 ^ v698 - base.I32_rotl(v706, int32(16))
	v714 = v710 ^ v702 - base.I32_rotl(v710, int32(4))
	v718 = v714 ^ v706 - base.I32_rotl(v714, v696)
	goto L84
L90:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	v691 = v683 + v686
	v692 = v684
	v693 = v685
	goto L89
L91:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+1)))
	v683 = v679<<(uint(int32(8))%32) + v676
	v684 = v677
	v685 = v678
	goto L90
L92:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+2)))
	v676 = v672<<(uint(int32(16))%32) + v669
	v677 = v670
	v678 = v671
	goto L91
L93:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+3)))
	v669 = v665<<(uint(int32(24))%32) + v516
	v670 = v663
	v671 = v664
	goto L92
L94:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+4)))
	v663 = v659 + v661
	v664 = v660
	goto L93
L95:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+5)))
	v659 = v655<<(uint(int32(8))%32) + v653
	v660 = v654
	goto L94
L96:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+6)))
	v653 = v649<<(uint(int32(16))%32) + v647
	v654 = v648
	goto L95
L97:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+7)))
	v647 = v643<<(uint(int32(24))%32) + v517
	v648 = v642
	goto L96
L98:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+8)))
	v642 = v638<<(uint(int32(8))%32) + v637
	goto L97
L99:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+9)))
	v637 = v633<<(uint(int32(16))%32) + v632
	goto L98
L100:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+10)))
	v632 = v628<<(uint(int32(24))%32) + v518
	goto L99
L101:
	;
	if base.Ui32(int32(11)) < base.Ui32(v405) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v405) {
		goto L110
	} else {
		goto L111
	}
L104:
	;
	v464 = v406
	v465 = v405
	v467 = v455
	v468 = v459
	v469 = v457
	goto L107
L105:
	;
	v513 = v406
	v514 = v405
	v516 = v455
	v517 = v459
	v518 = v457
	goto L106
L106:
	;
	switch v514 - int32(1) {
	case 0:
		v683 = v516
		v684 = v517
		v685 = v518
		goto L90
	case 1:
		v676 = v516
		v677 = v517
		v678 = v518
		goto L91
	case 2:
		v669 = v516
		v670 = v517
		v671 = v518
		goto L92
	case 3:
		v663 = v517
		v664 = v518
		goto L93
	case 4:
		v659 = v517
		v660 = v518
		goto L94
	case 5:
		v653 = v517
		v654 = v518
		goto L95
	case 6:
		v647 = v517
		v648 = v518
		goto L96
	case 7:
		v642 = v518
		goto L97
	case 8:
		v637 = v518
		goto L98
	case 9:
		v632 = v518
		goto L99
	case 10:
		goto L100
	default:
		v691 = v516
		v692 = v517
		v693 = v518
		goto L89
	}
L107:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	v472 = v471 + v468
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v464)+8))
	v476 = v475 + v469
	v478 = int32(4)
	v480 = v473 + v467 - v476 ^ base.I32_rotl(v476, v478)
	v484 = v472 - v480 ^ base.I32_rotl(v480, int32(6))
	v485 = v476 + v472
	v486 = v480 + v485
	v487 = v484 + v486
	v491 = v485 - v484 ^ base.I32_rotl(v484, int32(8))
	v495 = v486 - v491 ^ base.I32_rotl(v491, int32(16))
	v499 = v487 - v495 ^ base.I32_rotl(v495, int32(19))
	v500 = v491 + v487
	v501 = v495 + v500
	v502 = v499 + v501
	v506 = v500 - v499 ^ base.I32_rotl(v499, v478)
	v507 = int32(12)
	v508 = v464 + v507
	v510 = v465 - v507
	if base.Ui32(int32(11)) < base.Ui32(v510) {
		v464 = v508
		v465 = v510
		v467 = v501
		v468 = v502
		v469 = v506
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v513 = v508
	v514 = v510
	v516 = v501
	v517 = v502
	v518 = v506
	goto L106
L109:
	;
	goto L108
L110:
	;
	v524 = v406
	v525 = v405
	v527 = v455
	v528 = v459
	v529 = v457
	goto L113
L111:
	;
	v573 = v406
	v574 = v405
	v576 = v455
	v577 = v459
	v578 = v457
	goto L112
L112:
	;
	switch v574 - int32(1) {
	case 0:
		v625 = v576
		goto L116
	case 1:
		v620 = v576
		goto L117
	case 2:
		goto L118
	case 3:
		v613 = v577
		goto L119
	case 4:
		v610 = v577
		goto L120
	case 5:
		v605 = v577
		goto L121
	case 6:
		goto L122
	case 7:
		v596 = v578
		goto L123
	case 8:
		v591 = v578
		goto L124
	case 9:
		v586 = v578
		goto L125
	case 10:
		goto L126
	default:
		v691 = v576
		v692 = v577
		v693 = v578
		goto L89
	}
L113:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	v532 = v531 + v528
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v524)+8))
	v536 = v535 + v529
	v538 = int32(4)
	v540 = v533 + v527 - v536 ^ base.I32_rotl(v536, v538)
	v544 = v532 - v540 ^ base.I32_rotl(v540, int32(6))
	v545 = v536 + v532
	v546 = v540 + v545
	v547 = v544 + v546
	v551 = v545 - v544 ^ base.I32_rotl(v544, int32(8))
	v555 = v546 - v551 ^ base.I32_rotl(v551, int32(16))
	v559 = v547 - v555 ^ base.I32_rotl(v555, int32(19))
	v560 = v551 + v547
	v561 = v555 + v560
	v562 = v559 + v561
	v566 = v560 - v559 ^ base.I32_rotl(v559, v538)
	v567 = int32(12)
	v568 = v524 + v567
	v570 = v525 - v567
	if base.Ui32(int32(11)) < base.Ui32(v570) {
		v524 = v568
		v525 = v570
		v527 = v561
		v528 = v562
		v529 = v566
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v573 = v568
	v574 = v570
	v576 = v561
	v577 = v562
	v578 = v566
	goto L112
L115:
	;
	goto L114
L116:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573))))
	v691 = v625 + v626
	v692 = v577
	v693 = v578
	goto L89
L117:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+1)))
	v625 = v621<<(uint(int32(8))%32) + v620
	goto L116
L118:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+2)))
	v620 = v616<<(uint(int32(16))%32) + v576
	goto L117
L119:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v691 = v614 + v576
	v692 = v613
	v693 = v578
	goto L89
L120:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+4)))
	v613 = v610 + v611
	goto L119
L121:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+5)))
	v610 = v606<<(uint(int32(8))%32) + v605
	goto L120
L122:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+6)))
	v605 = v601<<(uint(int32(16))%32) + v577
	goto L121
L123:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	v691 = v597 + v576
	v692 = v599 + v577
	v693 = v596
	goto L89
L124:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+8)))
	v596 = v592<<(uint(int32(8))%32) + v591
	goto L123
L125:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+9)))
	v591 = v587<<(uint(int32(16))%32) + v586
	goto L124
L126:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+10)))
	v586 = v582<<(uint(int32(24))%32) + v578
	goto L125
L127:
	;
	F_pfree(m, v406)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v732 = v728
	goto L31
L129:
	;
	F_pfree(m, v10)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	return v732
L132:
	;
	goto L131
L133:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(323474), int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errhint(m, int32(541241), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(478807), int32(1057), int32(444894))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errmsg_internal(m, int32(93809), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(478807), int32(1082), int32(444894))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v18
	v25 = v14 + int32(-32)
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v18
	v29 = v14 + int32(-40)
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v18
	v33 = v14 + int32(-48)
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v18
	v37 = v14 + int32(-56)
	*(*int64)(unsafe.Add(mBase, uint32(v37))) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v18
	F_genericcostestimate(m, l0, l1, l2, v16)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		return
	} else {
		v44 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = v44
		v46 = *(*float64)(unsafe.Add(mBase, uint32(v37)))
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v46
		v48 = *(*float64)(unsafe.Add(mBase, uint32(v33)))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v48
		v50 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v50
		v52 = *(*float64)(unsafe.Add(mBase, uint32(v25)))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v52
		m.G0 = v16 - int32(-64)
		return
	}
}
func F_hashfloat8extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
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
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	if base.F64_eq(v10, float64(0)) != 0 {
		v16 = F_Int64GetDatum(m, v13)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v347 = v16
			m.G0 = v7 + int32(16)
			return v347
		}
	} else {
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)) {
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(9221120237041090560)
		} else {
		}
		v28 = v7 + int32(8)
		v35 = int32(-1636608424)
		if v13 == int64(0) {
			v72 = v35
			v74 = v35
			v76 = v35
		} else {
			v38 = base.I32_wrap_i64(v13)
			v40 = v38 + int32(1021750448)
			v44 = int32(4)
			v46 = base.I32_wrap_i64(int64(base.Ui64(v13)>>(uint(int64(32))%64))) ^ base.I32_rotl(v35, v44)
			v50 = v35 + v38 - v46 ^ base.I32_rotl(v46, int32(6))
			v54 = v40 - v50 ^ base.I32_rotl(v50, int32(8))
			v55 = v46 + v40
			v56 = v50 + v55
			v57 = v54 + v56
			v61 = v55 - v54 ^ base.I32_rotl(v54, int32(16))
			v65 = v56 - v61 ^ base.I32_rotl(v61, int32(19))
			v70 = v61 + v57
			v72 = v70
			v74 = v57 - v65 ^ base.I32_rotl(v65, v44)
			v76 = v65 + v70
		}
		if v28&int32(3) != 0 {
			switch int32(7) {
			case 0:
				v300 = v72
				v301 = v76
				v302 = v74
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 1:
				v293 = v72
				v294 = v76
				v295 = v74
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 2:
				v286 = v72
				v287 = v76
				v288 = v74
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v293 = v289<<(uint(int32(16))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 3:
				v280 = v76
				v281 = v74
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
				v286 = v282<<(uint(int32(24))%32) + v72
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v293 = v289<<(uint(int32(16))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 4:
				v276 = v76
				v277 = v74
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v280 = v276 + v278
				v281 = v277
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
				v286 = v282<<(uint(int32(24))%32) + v72
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v293 = v289<<(uint(int32(16))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 5:
				v270 = v76
				v271 = v74
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
				v276 = v272<<(uint(int32(8))%32) + v270
				v277 = v271
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v280 = v276 + v278
				v281 = v277
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
				v286 = v282<<(uint(int32(24))%32) + v72
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v293 = v289<<(uint(int32(16))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 6:
				v264 = v76
				v265 = v74
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
				v270 = v266<<(uint(int32(16))%32) + v264
				v271 = v265
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
				v276 = v272<<(uint(int32(8))%32) + v270
				v277 = v271
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v280 = v276 + v278
				v281 = v277
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
				v286 = v282<<(uint(int32(24))%32) + v72
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v293 = v289<<(uint(int32(16))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 7:
				v259 = v74
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
				v264 = v260<<(uint(int32(24))%32) + v76
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
				v270 = v266<<(uint(int32(16))%32) + v264
				v271 = v265
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
				v276 = v272<<(uint(int32(8))%32) + v270
				v277 = v271
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v280 = v276 + v278
				v281 = v277
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
				v286 = v282<<(uint(int32(24))%32) + v72
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v293 = v289<<(uint(int32(16))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 8:
				v254 = v74
				v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
				v259 = v255<<(uint(int32(8))%32) + v254
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
				v264 = v260<<(uint(int32(24))%32) + v76
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
				v270 = v266<<(uint(int32(16))%32) + v264
				v271 = v265
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
				v276 = v272<<(uint(int32(8))%32) + v270
				v277 = v271
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v280 = v276 + v278
				v281 = v277
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
				v286 = v282<<(uint(int32(24))%32) + v72
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v293 = v289<<(uint(int32(16))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 9:
				v249 = v74
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
				v254 = v250<<(uint(int32(16))%32) + v249
				v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
				v259 = v255<<(uint(int32(8))%32) + v254
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
				v264 = v260<<(uint(int32(24))%32) + v76
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
				v270 = v266<<(uint(int32(16))%32) + v264
				v271 = v265
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
				v276 = v272<<(uint(int32(8))%32) + v270
				v277 = v271
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v280 = v276 + v278
				v281 = v277
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
				v286 = v282<<(uint(int32(24))%32) + v72
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v293 = v289<<(uint(int32(16))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			case 10:
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
				v249 = v245<<(uint(int32(24))%32) + v74
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
				v254 = v250<<(uint(int32(16))%32) + v249
				v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
				v259 = v255<<(uint(int32(8))%32) + v254
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
				v264 = v260<<(uint(int32(24))%32) + v76
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
				v270 = v266<<(uint(int32(16))%32) + v264
				v271 = v265
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
				v276 = v272<<(uint(int32(8))%32) + v270
				v277 = v271
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v280 = v276 + v278
				v281 = v277
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
				v286 = v282<<(uint(int32(24))%32) + v72
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v293 = v289<<(uint(int32(16))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v300 = v296<<(uint(int32(8))%32) + v293
				v301 = v294
				v302 = v295
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v300 + v303
				v309 = v301
				v310 = v302
			default:
				v308 = v72
				v309 = v76
				v310 = v74
			}
		} else {
			switch int32(7) {
			case 0:
				v242 = v72
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v242 + v243
				v309 = v76
				v310 = v74
			case 1:
				v237 = v72
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v242 = v238<<(uint(int32(8))%32) + v237
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v242 + v243
				v309 = v76
				v310 = v74
			case 2:
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				v237 = v233<<(uint(int32(16))%32) + v72
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
				v242 = v238<<(uint(int32(8))%32) + v237
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				v308 = v242 + v243
				v309 = v76
				v310 = v74
			case 3:
				v230 = v76
				v231 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v308 = v231 + v72
				v309 = v230
				v310 = v74
			case 4:
				v227 = v76
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v230 = v227 + v228
				v231 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v308 = v231 + v72
				v309 = v230
				v310 = v74
			case 5:
				v222 = v76
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
				v227 = v223<<(uint(int32(8))%32) + v222
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v230 = v227 + v228
				v231 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v308 = v231 + v72
				v309 = v230
				v310 = v74
			case 6:
				v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
				v222 = v218<<(uint(int32(16))%32) + v76
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
				v227 = v223<<(uint(int32(8))%32) + v222
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
				v230 = v227 + v228
				v231 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v308 = v231 + v72
				v309 = v230
				v310 = v74
			case 7:
				v213 = v74
				v214 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v308 = v214 + v72
				v309 = v216 + v76
				v310 = v213
			case 8:
				v208 = v74
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
				v213 = v209<<(uint(int32(8))%32) + v208
				v214 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v308 = v214 + v72
				v309 = v216 + v76
				v310 = v213
			case 9:
				v203 = v74
				v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
				v208 = v204<<(uint(int32(16))%32) + v203
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
				v213 = v209<<(uint(int32(8))%32) + v208
				v214 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v308 = v214 + v72
				v309 = v216 + v76
				v310 = v213
			case 10:
				v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
				v203 = v199<<(uint(int32(24))%32) + v74
				v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
				v208 = v204<<(uint(int32(16))%32) + v203
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
				v213 = v209<<(uint(int32(8))%32) + v208
				v214 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v308 = v214 + v72
				v309 = v216 + v76
				v310 = v213
			default:
				v308 = v72
				v309 = v76
				v310 = v74
			}
		}
		v313 = int32(14)
		v315 = v309 ^ v310 - base.I32_rotl(v309, v313)
		v319 = v315 ^ v308 - base.I32_rotl(v315, int32(11))
		v323 = v319 ^ v309 - base.I32_rotl(v319, int32(25))
		v327 = v323 ^ v315 - base.I32_rotl(v323, int32(16))
		v331 = v327 ^ v319 - base.I32_rotl(v327, int32(4))
		v335 = v331 ^ v323 - base.I32_rotl(v331, v313)
		v345 = F_Int64GetDatum(m, base.I64_extend_i32_u(v335)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v327^v335-base.I32_rotl(v335, int32(24))))
		mBase = m.M
		v346 = m.ExcPending
		if v346 != 0 {
			return int32(0)
		} else {
			v347 = v345
			m.G0 = v7 + int32(16)
			return v347
		}
	}
}
func F_hashinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v17 = F__hash_convert_tuple(m, l0, l1, l2, v11+int32(12), v11+int32(11))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v26 = F_index_form_tuple(m, v21, v11+int32(12), v11+int32(11))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)) = uint16(v28)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = v30
				F__hash_doinsert(m, l0, v26, l4, int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v26)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						m.G0 = v11 + int32(16)
						return int32(0)
					}
				}
			}
		} else {
			m.G0 = v11 + int32(16)
			return int32(0)
		}
	}
}
func F_hashint2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v7 = int32(711645284)
	v10 = v2 - int32(1636608428) ^ v7 - int32(1455628627)
	v15 = v10 ^ int32(-1636608428) - base.I32_rotl(v10, int32(25))
	v20 = v15 ^ v7 - base.I32_rotl(v15, int32(16))
	v24 = v20 ^ v10 - base.I32_rotl(v20, int32(4))
	v28 = v24 ^ v15 - base.I32_rotl(v24, int32(14))
	return v28 ^ v20 - base.I32_rotl(v28, int32(24))
}
func F_hashmacaddrextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v11 = int32(-1636608426)
	if v5 == int64(0) {
		v48 = v11
		v50 = v11
		v52 = v11
	} else {
		v14 = base.I32_wrap_i64(v5)
		v16 = v14 + int32(1021750444)
		v20 = int32(4)
		v22 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) ^ base.I32_rotl(v11, v20)
		v26 = v11 + v14 - v22 ^ base.I32_rotl(v22, int32(6))
		v30 = v16 - v26 ^ base.I32_rotl(v26, int32(8))
		v31 = v22 + v16
		v32 = v26 + v31
		v33 = v30 + v32
		v37 = v31 - v30 ^ base.I32_rotl(v30, int32(16))
		v41 = v32 - v37 ^ base.I32_rotl(v37, int32(19))
		v46 = v37 + v33
		v48 = v46
		v50 = v33 - v41 ^ base.I32_rotl(v41, v20)
		v52 = v41 + v46
	}
	if v2&int32(3) != 0 {
		switch int32(5) {
		case 0:
			v276 = v48
			v277 = v52
			v278 = v50
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 1:
			v269 = v48
			v270 = v52
			v271 = v50
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 2:
			v262 = v48
			v263 = v52
			v264 = v50
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 3:
			v256 = v52
			v257 = v50
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 4:
			v252 = v52
			v253 = v50
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 5:
			v246 = v52
			v247 = v50
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 6:
			v240 = v52
			v241 = v50
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 7:
			v235 = v50
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v240 = v236<<(uint(int32(24))%32) + v52
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 8:
			v230 = v50
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v240 = v236<<(uint(int32(24))%32) + v52
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 9:
			v225 = v50
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v230 = v226<<(uint(int32(16))%32) + v225
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v240 = v236<<(uint(int32(24))%32) + v52
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 10:
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v225 = v221<<(uint(int32(24))%32) + v50
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v230 = v226<<(uint(int32(16))%32) + v225
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v240 = v236<<(uint(int32(24))%32) + v52
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		default:
			v284 = v48
			v285 = v52
			v286 = v50
		}
	} else {
		switch int32(5) {
		case 0:
			v218 = v48
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v218 + v219
			v285 = v52
			v286 = v50
		case 1:
			v213 = v48
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v218 = v214<<(uint(int32(8))%32) + v213
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v218 + v219
			v285 = v52
			v286 = v50
		case 2:
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v213 = v209<<(uint(int32(16))%32) + v48
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v218 = v214<<(uint(int32(8))%32) + v213
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v218 + v219
			v285 = v52
			v286 = v50
		case 3:
			v206 = v52
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v284 = v207 + v48
			v285 = v206
			v286 = v50
		case 4:
			v203 = v52
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v284 = v207 + v48
			v285 = v206
			v286 = v50
		case 5:
			v198 = v52
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v284 = v207 + v48
			v285 = v206
			v286 = v50
		case 6:
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v198 = v194<<(uint(int32(16))%32) + v52
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v284 = v207 + v48
			v285 = v206
			v286 = v50
		case 7:
			v189 = v50
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v284 = v190 + v48
			v285 = v192 + v52
			v286 = v189
		case 8:
			v184 = v50
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v284 = v190 + v48
			v285 = v192 + v52
			v286 = v189
		case 9:
			v179 = v50
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v284 = v190 + v48
			v285 = v192 + v52
			v286 = v189
		case 10:
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v179 = v175<<(uint(int32(24))%32) + v50
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v284 = v190 + v48
			v285 = v192 + v52
			v286 = v189
		default:
			v284 = v48
			v285 = v52
			v286 = v50
		}
	}
	v289 = int32(14)
	v291 = v285 ^ v286 - base.I32_rotl(v285, v289)
	v295 = v291 ^ v284 - base.I32_rotl(v291, int32(11))
	v299 = v295 ^ v285 - base.I32_rotl(v295, int32(25))
	v303 = v299 ^ v291 - base.I32_rotl(v299, int32(16))
	v307 = v303 ^ v295 - base.I32_rotl(v303, int32(4))
	v311 = v307 ^ v299 - base.I32_rotl(v307, v289)
	v321 = F_Int64GetDatum(m, base.I64_extend_i32_u(v311)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v303^v311-base.I32_rotl(v311, int32(24))))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		return int32(0)
	} else {
		return v321
	}
}
func F_hashname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
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
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
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
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2&int32(3) == int32(0) {
		v26 = v2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v65 = v59 - int32(1636608432)
	if v2&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v59 = v51 - v2
	goto L1
L3:
	;
	v30 = v26
	goto L12
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v59 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v15 = v2
	goto L8
L8:
	;
	v19 = v15 + int32(1)
	if v19&int32(3) == int32(0) {
		v26 = v19
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v51 = v19
	goto L2
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 != 0 {
		v15 = v19
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 == v39 {
		v30 = v30 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v45 = v30
	goto L15
L14:
	;
	goto L13
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		v45 = v45 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v51 = v45
	goto L2
L17:
	;
	goto L16
L18:
	;
	return v319 ^ v311 - base.I32_rotl(v319, int32(24))
L19:
	;
	v297 = int32(14)
	v299 = v293 ^ v294 - base.I32_rotl(v293, v297)
	v303 = v299 ^ v292 - base.I32_rotl(v299, int32(11))
	v307 = v303 ^ v293 - base.I32_rotl(v303, int32(25))
	v311 = v307 ^ v299 - base.I32_rotl(v307, int32(16))
	v315 = v311 ^ v303 - base.I32_rotl(v311, int32(4))
	v319 = v315 ^ v307 - base.I32_rotl(v315, v297)
	goto L18
L20:
	;
	switch v223 - int32(1) {
	case 0:
		v285 = v224
		v286 = v225
		v287 = v226
		goto L47
	case 1:
		v278 = v224
		v279 = v225
		v280 = v226
		goto L48
	case 2:
		v271 = v224
		v272 = v225
		v273 = v226
		goto L49
	case 3:
		v265 = v225
		v266 = v226
		goto L50
	case 4:
		v261 = v225
		v262 = v226
		goto L51
	case 5:
		v255 = v225
		v256 = v226
		goto L52
	case 6:
		v249 = v225
		v250 = v226
		goto L53
	case 7:
		v244 = v226
		goto L54
	case 8:
		v239 = v226
		goto L55
	case 9:
		v234 = v226
		goto L56
	case 10:
		goto L57
	default:
		v292 = v224
		v293 = v225
		v294 = v226
		goto L19
	}
L21:
	;
	v174 = v2
	v175 = v59
	v176 = v65
	v177 = v65
	v178 = v65
	goto L44
L22:
	;
	if base.Ui32(int32(11)) < base.Ui32(v59) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v59) < base.Ui32(int32(12)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v222 = v2
	v223 = v59
	v224 = v65
	v225 = v65
	v226 = v65
	goto L20
L26:
	;
	switch v121 - int32(1) {
	case 0:
		v171 = v122
		goto L33
	case 1:
		v166 = v122
		goto L34
	case 2:
		goto L35
	case 3:
		v159 = v123
		goto L36
	case 4:
		v156 = v123
		goto L37
	case 5:
		v151 = v123
		goto L38
	case 6:
		goto L39
	case 7:
		v142 = v124
		goto L40
	case 8:
		v137 = v124
		goto L41
	case 9:
		v132 = v124
		goto L42
	case 10:
		goto L43
	default:
		v292 = v122
		v293 = v123
		v294 = v124
		goto L19
	}
L27:
	;
	v120 = v2
	v121 = v59
	v122 = v65
	v123 = v65
	v124 = v65
	goto L26
L28:
	;
	goto L29
L29:
	;
	v72 = v2
	v73 = v59
	v74 = v65
	v75 = v65
	v76 = v65
	goto L30
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v79 = v78 + v75
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v83 = v82 + v76
	v85 = int32(4)
	v87 = v80 + v74 - v83 ^ base.I32_rotl(v83, v85)
	v91 = v79 - v87 ^ base.I32_rotl(v87, int32(6))
	v92 = v83 + v79
	v93 = v87 + v92
	v94 = v91 + v93
	v98 = v92 - v91 ^ base.I32_rotl(v91, int32(8))
	v102 = v93 - v98 ^ base.I32_rotl(v98, int32(16))
	v106 = v94 - v102 ^ base.I32_rotl(v102, int32(19))
	v107 = v98 + v94
	v108 = v102 + v107
	v109 = v106 + v108
	v113 = v107 - v106 ^ base.I32_rotl(v106, v85)
	v114 = int32(12)
	v115 = v72 + v114
	v117 = v73 - v114
	if base.Ui32(int32(11)) < base.Ui32(v117) {
		v72 = v115
		v73 = v117
		v74 = v108
		v75 = v109
		v76 = v113
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v120 = v115
	v121 = v117
	v122 = v108
	v123 = v109
	v124 = v113
	goto L26
L32:
	;
	goto L31
L33:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v292 = v171 + v172
	v293 = v123
	v294 = v124
	goto L19
L34:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	v171 = v167<<(uint(int32(8))%32) + v166
	goto L33
L35:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+2)))
	v166 = v162<<(uint(int32(16))%32) + v122
	goto L34
L36:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v292 = v160 + v122
	v293 = v159
	v294 = v124
	goto L19
L37:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+4)))
	v159 = v156 + v157
	goto L36
L38:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)))
	v156 = v152<<(uint(int32(8))%32) + v151
	goto L37
L39:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+6)))
	v151 = v147<<(uint(int32(16))%32) + v123
	goto L38
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v292 = v143 + v122
	v293 = v145 + v123
	v294 = v142
	goto L19
L41:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)))
	v142 = v138<<(uint(int32(8))%32) + v137
	goto L40
L42:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+9)))
	v137 = v133<<(uint(int32(16))%32) + v132
	goto L41
L43:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+10)))
	v132 = v128<<(uint(int32(24))%32) + v124
	goto L42
L44:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	v181 = v180 + v177
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	v185 = v184 + v178
	v187 = int32(4)
	v189 = v182 + v176 - v185 ^ base.I32_rotl(v185, v187)
	v193 = v181 - v189 ^ base.I32_rotl(v189, int32(6))
	v194 = v185 + v181
	v195 = v189 + v194
	v196 = v193 + v195
	v200 = v194 - v193 ^ base.I32_rotl(v193, int32(8))
	v204 = v195 - v200 ^ base.I32_rotl(v200, int32(16))
	v208 = v196 - v204 ^ base.I32_rotl(v204, int32(19))
	v209 = v200 + v196
	v210 = v204 + v209
	v211 = v208 + v210
	v215 = v209 - v208 ^ base.I32_rotl(v208, v187)
	v216 = int32(12)
	v217 = v174 + v216
	v219 = v175 - v216
	if base.Ui32(int32(11)) < base.Ui32(v219) {
		v174 = v217
		v175 = v219
		v176 = v210
		v177 = v211
		v178 = v215
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v222 = v217
	v223 = v219
	v224 = v210
	v225 = v211
	v226 = v215
	goto L20
L46:
	;
	goto L45
L47:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v292 = v285 + v288
	v293 = v286
	v294 = v287
	goto L19
L48:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	v285 = v281<<(uint(int32(8))%32) + v278
	v286 = v279
	v287 = v280
	goto L47
L49:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+2)))
	v278 = v274<<(uint(int32(16))%32) + v271
	v279 = v272
	v280 = v273
	goto L48
L50:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+3)))
	v271 = v267<<(uint(int32(24))%32) + v224
	v272 = v265
	v273 = v266
	goto L49
L51:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+4)))
	v265 = v261 + v263
	v266 = v262
	goto L50
L52:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+5)))
	v261 = v257<<(uint(int32(8))%32) + v255
	v262 = v256
	goto L51
L53:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+6)))
	v255 = v251<<(uint(int32(16))%32) + v249
	v256 = v250
	goto L52
L54:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+7)))
	v249 = v245<<(uint(int32(24))%32) + v225
	v250 = v244
	goto L53
L55:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+8)))
	v244 = v240<<(uint(int32(8))%32) + v239
	goto L54
L56:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+9)))
	v239 = v235<<(uint(int32(16))%32) + v234
	goto L55
L57:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)))
	v234 = v230<<(uint(int32(24))%32) + v226
	goto L56
}
func F_hashoidvector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_valid_oidvector(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = v2 + int32(24)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
		v11 = v9 << (uint(int32(2)) % 32)
		v17 = v11 - int32(1636608432)
		if v8&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v11) {
				v126 = v8
				v127 = v11
				v128 = v17
				v129 = v17
				v130 = v17
				for {
					v132 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
					v133 = v132 + v129
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
					v137 = v136 + v130
					v139 = int32(4)
					v141 = v134 + v128 - v137 ^ base.I32_rotl(v137, v139)
					v145 = v133 - v141 ^ base.I32_rotl(v141, int32(6))
					v146 = v137 + v133
					v147 = v141 + v146
					v148 = v145 + v147
					v152 = v146 - v145 ^ base.I32_rotl(v145, int32(8))
					v156 = v147 - v152 ^ base.I32_rotl(v152, int32(16))
					v160 = v148 - v156 ^ base.I32_rotl(v156, int32(19))
					v161 = v152 + v148
					v162 = v156 + v161
					v163 = v160 + v162
					v167 = v161 - v160 ^ base.I32_rotl(v160, v139)
					v168 = int32(12)
					v169 = v126 + v168
					v171 = v127 - v168
					if base.Ui32(int32(11)) < base.Ui32(v171) {
						v126 = v169
						v127 = v171
						v128 = v162
						v129 = v163
						v130 = v167
						continue
					} else {
						break
					}
					break
				}
				v174 = v169
				v175 = v171
				v176 = v162
				v177 = v163
				v178 = v167
			} else {
				v174 = v8
				v175 = v11
				v176 = v17
				v177 = v17
				v178 = v17
			}
			switch v175 - int32(1) {
			case 0:
				v237 = v176
				v238 = v177
				v239 = v178
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 1:
				v230 = v176
				v231 = v177
				v232 = v178
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 2:
				v223 = v176
				v224 = v177
				v225 = v178
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v230 = v226<<(uint(int32(16))%32) + v223
				v231 = v224
				v232 = v225
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 3:
				v217 = v177
				v218 = v178
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
				v223 = v219<<(uint(int32(24))%32) + v176
				v224 = v217
				v225 = v218
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v230 = v226<<(uint(int32(16))%32) + v223
				v231 = v224
				v232 = v225
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 4:
				v213 = v177
				v214 = v178
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v217 = v213 + v215
				v218 = v214
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
				v223 = v219<<(uint(int32(24))%32) + v176
				v224 = v217
				v225 = v218
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v230 = v226<<(uint(int32(16))%32) + v223
				v231 = v224
				v232 = v225
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 5:
				v207 = v177
				v208 = v178
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
				v213 = v209<<(uint(int32(8))%32) + v207
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v217 = v213 + v215
				v218 = v214
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
				v223 = v219<<(uint(int32(24))%32) + v176
				v224 = v217
				v225 = v218
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v230 = v226<<(uint(int32(16))%32) + v223
				v231 = v224
				v232 = v225
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 6:
				v201 = v177
				v202 = v178
				v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
				v207 = v203<<(uint(int32(16))%32) + v201
				v208 = v202
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
				v213 = v209<<(uint(int32(8))%32) + v207
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v217 = v213 + v215
				v218 = v214
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
				v223 = v219<<(uint(int32(24))%32) + v176
				v224 = v217
				v225 = v218
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v230 = v226<<(uint(int32(16))%32) + v223
				v231 = v224
				v232 = v225
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 7:
				v196 = v178
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+7)))
				v201 = v197<<(uint(int32(24))%32) + v177
				v202 = v196
				v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
				v207 = v203<<(uint(int32(16))%32) + v201
				v208 = v202
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
				v213 = v209<<(uint(int32(8))%32) + v207
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v217 = v213 + v215
				v218 = v214
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
				v223 = v219<<(uint(int32(24))%32) + v176
				v224 = v217
				v225 = v218
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v230 = v226<<(uint(int32(16))%32) + v223
				v231 = v224
				v232 = v225
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 8:
				v191 = v178
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
				v196 = v192<<(uint(int32(8))%32) + v191
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+7)))
				v201 = v197<<(uint(int32(24))%32) + v177
				v202 = v196
				v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
				v207 = v203<<(uint(int32(16))%32) + v201
				v208 = v202
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
				v213 = v209<<(uint(int32(8))%32) + v207
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v217 = v213 + v215
				v218 = v214
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
				v223 = v219<<(uint(int32(24))%32) + v176
				v224 = v217
				v225 = v218
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v230 = v226<<(uint(int32(16))%32) + v223
				v231 = v224
				v232 = v225
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 9:
				v186 = v178
				v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+9)))
				v191 = v187<<(uint(int32(16))%32) + v186
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
				v196 = v192<<(uint(int32(8))%32) + v191
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+7)))
				v201 = v197<<(uint(int32(24))%32) + v177
				v202 = v196
				v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
				v207 = v203<<(uint(int32(16))%32) + v201
				v208 = v202
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
				v213 = v209<<(uint(int32(8))%32) + v207
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v217 = v213 + v215
				v218 = v214
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
				v223 = v219<<(uint(int32(24))%32) + v176
				v224 = v217
				v225 = v218
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v230 = v226<<(uint(int32(16))%32) + v223
				v231 = v224
				v232 = v225
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			case 10:
				v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+10)))
				v186 = v182<<(uint(int32(24))%32) + v178
				v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+9)))
				v191 = v187<<(uint(int32(16))%32) + v186
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
				v196 = v192<<(uint(int32(8))%32) + v191
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+7)))
				v201 = v197<<(uint(int32(24))%32) + v177
				v202 = v196
				v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
				v207 = v203<<(uint(int32(16))%32) + v201
				v208 = v202
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
				v213 = v209<<(uint(int32(8))%32) + v207
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v217 = v213 + v215
				v218 = v214
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
				v223 = v219<<(uint(int32(24))%32) + v176
				v224 = v217
				v225 = v218
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v230 = v226<<(uint(int32(16))%32) + v223
				v231 = v224
				v232 = v225
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v237 = v233<<(uint(int32(8))%32) + v230
				v238 = v231
				v239 = v232
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v244 = v237 + v240
				v245 = v238
				v246 = v239
			default:
				v244 = v176
				v245 = v177
				v246 = v178
			}
		} else {
			if base.Ui32(v11) < base.Ui32(int32(12)) {
				v72 = v8
				v73 = v11
				v74 = v17
				v75 = v17
				v76 = v17
			} else {
				v24 = v8
				v25 = v11
				v26 = v17
				v27 = v17
				v28 = v17
				for {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					v31 = v30 + v27
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
					v35 = v34 + v28
					v37 = int32(4)
					v39 = v32 + v26 - v35 ^ base.I32_rotl(v35, v37)
					v43 = v31 - v39 ^ base.I32_rotl(v39, int32(6))
					v44 = v35 + v31
					v45 = v39 + v44
					v46 = v43 + v45
					v50 = v44 - v43 ^ base.I32_rotl(v43, int32(8))
					v54 = v45 - v50 ^ base.I32_rotl(v50, int32(16))
					v58 = v46 - v54 ^ base.I32_rotl(v54, int32(19))
					v59 = v50 + v46
					v60 = v54 + v59
					v61 = v58 + v60
					v65 = v59 - v58 ^ base.I32_rotl(v58, v37)
					v66 = int32(12)
					v67 = v24 + v66
					v69 = v25 - v66
					if base.Ui32(int32(11)) < base.Ui32(v69) {
						v24 = v67
						v25 = v69
						v26 = v60
						v27 = v61
						v28 = v65
						continue
					} else {
						break
					}
					break
				}
				v72 = v67
				v73 = v69
				v74 = v60
				v75 = v61
				v76 = v65
			}
			switch v73 - int32(1) {
			case 0:
				v123 = v74
				v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
				v244 = v123 + v124
				v245 = v75
				v246 = v76
			case 1:
				v118 = v74
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
				v123 = v119<<(uint(int32(8))%32) + v118
				v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
				v244 = v123 + v124
				v245 = v75
				v246 = v76
			case 2:
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+2)))
				v118 = v114<<(uint(int32(16))%32) + v74
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
				v123 = v119<<(uint(int32(8))%32) + v118
				v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
				v244 = v123 + v124
				v245 = v75
				v246 = v76
			case 3:
				v111 = v75
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				v244 = v112 + v74
				v245 = v111
				v246 = v76
			case 4:
				v108 = v75
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)))
				v111 = v108 + v109
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				v244 = v112 + v74
				v245 = v111
				v246 = v76
			case 5:
				v103 = v75
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+5)))
				v108 = v104<<(uint(int32(8))%32) + v103
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)))
				v111 = v108 + v109
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				v244 = v112 + v74
				v245 = v111
				v246 = v76
			case 6:
				v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+6)))
				v103 = v99<<(uint(int32(16))%32) + v75
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+5)))
				v108 = v104<<(uint(int32(8))%32) + v103
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)))
				v111 = v108 + v109
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				v244 = v112 + v74
				v245 = v111
				v246 = v76
			case 7:
				v94 = v76
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
				v244 = v95 + v74
				v245 = v97 + v75
				v246 = v94
			case 8:
				v89 = v76
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+8)))
				v94 = v90<<(uint(int32(8))%32) + v89
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
				v244 = v95 + v74
				v245 = v97 + v75
				v246 = v94
			case 9:
				v84 = v76
				v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+9)))
				v89 = v85<<(uint(int32(16))%32) + v84
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+8)))
				v94 = v90<<(uint(int32(8))%32) + v89
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
				v244 = v95 + v74
				v245 = v97 + v75
				v246 = v94
			case 10:
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+10)))
				v84 = v80<<(uint(int32(24))%32) + v76
				v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+9)))
				v89 = v85<<(uint(int32(16))%32) + v84
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+8)))
				v94 = v90<<(uint(int32(8))%32) + v89
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
				v244 = v95 + v74
				v245 = v97 + v75
				v246 = v94
			default:
				v244 = v74
				v245 = v75
				v246 = v76
			}
		}
		v249 = int32(14)
		v251 = v245 ^ v246 - base.I32_rotl(v245, v249)
		v255 = v251 ^ v244 - base.I32_rotl(v251, int32(11))
		v259 = v255 ^ v245 - base.I32_rotl(v255, int32(25))
		v263 = v259 ^ v251 - base.I32_rotl(v259, int32(16))
		v267 = v263 ^ v255 - base.I32_rotl(v263, int32(4))
		v271 = v267 ^ v259 - base.I32_rotl(v267, v249)
		return v271 ^ v263 - base.I32_rotl(v271, int32(24))
	}
}
func F_hashvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	if l1 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v5 = F_RelationGetNumberOfBlocksInFork(m, v3, int32(0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
			return l1
		}
	} else {
		return l1
	}
}
func F_have_join_order_restriction(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v707 int32
	_ = v707
	v4 = int32(0)
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	if v9 == v4 {
		v52 = v4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v707
L2:
	;
	if v52 != 0 {
		v707 = v8
		goto L1
	} else {
		goto L16
	}
L3:
	;
	goto L2
L4:
	;
	if v10 == int32(0) {
		v52 = v4
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v20 < v21 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = v20
	goto L8
L7:
	;
	v23 = v21
	goto L8
L8:
	;
	if v23 <= int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = int32(1)
	goto L11
L10:
	;
	v26 = v23
	goto L11
L11:
	;
	v27 = int32(8)
	v32 = int32(0)
	goto L12
L12:
	;
	v39 = v32 << (uint(int32(2)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10+v27+v39)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v9+v27))))
	v44 = v41 & v43
	v46 = base.B2i32(v44 != int32(0))
	if v44 != 0 {
		v52 = v46
		goto L3
	} else {
		goto L14
	}
L13:
	;
	v52 = v46
	goto L3
L14:
	;
	v48 = v32 + int32(1)
	if v48 != v26 {
		v32 = v48
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v58 = int32(0)
	if v56 == v58 {
		v99 = v58
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v99 != 0 {
		v707 = v8
		goto L1
	} else {
		goto L31
	}
L18:
	;
	goto L17
L19:
	;
	if v57 == int32(0) {
		v99 = v58
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v67 < v68 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = v67
	goto L23
L22:
	;
	v70 = v68
	goto L23
L23:
	;
	if v70 <= int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v73 = int32(1)
	goto L26
L25:
	;
	v73 = v70
	goto L26
L26:
	;
	v74 = int32(8)
	v79 = int32(0)
	goto L27
L27:
	;
	v86 = v79 << (uint(int32(2)) % 32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v57+v74+v86)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+(v56+v74))))
	v91 = v88 & v90
	v93 = base.B2i32(v91 != int32(0))
	if v91 != 0 {
		v99 = v93
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v99 = v93
	goto L18
L29:
	;
	v95 = v79 + int32(1)
	if v95 != v73 {
		v79 = v95
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v103 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v707 = int32(1)
	goto L1
L33:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v244 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L34:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v106 <= int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v112 = v4
	goto L36
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v112<<(uint(int32(2))%32))))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v123 = int32(0)
	if v116 == v123 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L33
L38:
	;
	if v176 != 0 {
		goto L52
	} else {
		goto L53
	}
L39:
	;
	v176 = int32(1)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v122 == int32(0) {
		v167 = v123
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v176 = v167
	goto L38
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v133 < v132 {
		v167 = v123
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v135 = int32(1)
	if v132 <= v135 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v138 = v135
	goto L47
L46:
	;
	v138 = v132
	goto L47
L47:
	;
	v139 = int32(8)
	v144 = int32(0)
	goto L48
L48:
	;
	v151 = v144 << (uint(int32(2)) % 32)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v116+v139+v151)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+(v122+v139))))
	v158 = v153 & (v155 ^ int32(-1))
	v160 = base.B2i32(v158 == int32(0))
	if v158 != 0 {
		v167 = v160
		goto L42
	} else {
		goto L50
	}
L49:
	;
	v167 = v160
	goto L42
L50:
	;
	v162 = v144 + int32(1)
	if v162 != v138 {
		v144 = v162
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v179 = int32(0)
	if v177 == v179 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	v234 = v112 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v234 < v235 {
		v112 = v234
		goto L36
	} else {
		goto L70
	}
L55:
	;
	if v232 != 0 {
		goto L32
	} else {
		goto L69
	}
L56:
	;
	v232 = int32(1)
	goto L55
L57:
	;
	goto L58
L58:
	;
	if v178 == int32(0) {
		v223 = v179
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v232 = v223
	goto L55
L60:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if v189 < v188 {
		v223 = v179
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v191 = int32(1)
	if v188 <= v191 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v194 = v191
	goto L64
L63:
	;
	v194 = v188
	goto L64
L64:
	;
	v195 = int32(8)
	v200 = int32(0)
	goto L65
L65:
	;
	v207 = v200 << (uint(int32(2)) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v177+v195+v207)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207+(v178+v195))))
	v214 = v209 & (v211 ^ int32(-1))
	v216 = base.B2i32(v214 == int32(0))
	if v214 != 0 {
		v223 = v216
		goto L59
	} else {
		goto L67
	}
L66:
	;
	v223 = v216
	goto L59
L67:
	;
	v218 = v200 + int32(1)
	if v218 != v194 {
		v200 = v218
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	goto L54
L70:
	;
	goto L37
L71:
	;
	return int32(0)
L72:
	;
	goto L73
L73:
	;
	v249 = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v250 <= v249 {
		v707 = v249
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v258 = int32(0)
	goto L75
L75:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v262 = int32(2)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261+v258<<(uint(v262)%32))))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+20))
	if v266 == v262 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v687 = F_has_legal_joinclause(m, l0, l1)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L207
	} else {
		goto L208
	}
L77:
	;
	goto L76
L78:
	;
	v684 = v258 + int32(1)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v684 < v685 {
		v258 = v684
		goto L75
	} else {
		goto L206
	}
L79:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v271 = int32(0)
	if v269 == v271 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v324 != 0 {
		goto L94
	} else {
		goto L95
	}
L81:
	;
	v324 = int32(1)
	goto L80
L82:
	;
	goto L83
L83:
	;
	if v270 == int32(0) {
		v315 = v271
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v324 = v315
	goto L80
L85:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v281 < v280 {
		v315 = v271
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v283 = int32(1)
	if v280 <= v283 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v286 = v283
	goto L89
L88:
	;
	v286 = v280
	goto L89
L89:
	;
	v287 = int32(8)
	v292 = int32(0)
	goto L90
L90:
	;
	v299 = v292 << (uint(int32(2)) % 32)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v269+v287+v299)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v299+(v270+v287))))
	v306 = v301 & (v303 ^ int32(-1))
	v308 = base.B2i32(v306 == int32(0))
	if v306 != 0 {
		v315 = v308
		goto L84
	} else {
		goto L92
	}
L91:
	;
	v315 = v308
	goto L84
L92:
	;
	v310 = v292 + int32(1)
	if v310 != v286 {
		v292 = v310
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v327 = int32(0)
	if v325 == v327 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	goto L96
L96:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v383 = int32(0)
	if v381 == v383 {
		goto L113
	} else {
		goto L114
	}
L97:
	;
	if v380 != 0 {
		goto L77
	} else {
		goto L111
	}
L98:
	;
	v380 = int32(1)
	goto L97
L99:
	;
	goto L100
L100:
	;
	if v326 == int32(0) {
		v371 = v327
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v380 = v371
	goto L97
L102:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v337 < v336 {
		v371 = v327
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v339 = int32(1)
	if v336 <= v339 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v342 = v339
	goto L106
L105:
	;
	v342 = v336
	goto L106
L106:
	;
	v343 = int32(8)
	v348 = int32(0)
	goto L107
L107:
	;
	v355 = v348 << (uint(int32(2)) % 32)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v325+v343+v355)))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v355+(v326+v343))))
	v362 = v357 & (v359 ^ int32(-1))
	v364 = base.B2i32(v362 == int32(0))
	if v362 != 0 {
		v371 = v364
		goto L101
	} else {
		goto L109
	}
L108:
	;
	v371 = v364
	goto L101
L109:
	;
	v366 = v348 + int32(1)
	if v366 != v342 {
		v348 = v366
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	goto L96
L112:
	;
	if v436 != 0 {
		goto L126
	} else {
		goto L127
	}
L113:
	;
	v436 = int32(1)
	goto L112
L114:
	;
	goto L115
L115:
	;
	if v382 == int32(0) {
		v427 = v383
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v436 = v427
	goto L112
L117:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v393 < v392 {
		v427 = v383
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v395 = int32(1)
	if v392 <= v395 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v398 = v395
	goto L121
L120:
	;
	v398 = v392
	goto L121
L121:
	;
	v399 = int32(8)
	v404 = int32(0)
	goto L122
L122:
	;
	v411 = v404 << (uint(int32(2)) % 32)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v381+v399+v411)))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v411+(v382+v399))))
	v418 = v413 & (v415 ^ int32(-1))
	v420 = base.B2i32(v418 == int32(0))
	if v418 != 0 {
		v427 = v420
		goto L116
	} else {
		goto L124
	}
L123:
	;
	v427 = v420
	goto L116
L124:
	;
	v422 = v404 + int32(1)
	if v422 != v398 {
		v404 = v422
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v439 = int32(0)
	if v437 == v439 {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v495 = int32(0)
	if v493 == v495 {
		v536 = v495
		goto L145
	} else {
		goto L146
	}
L129:
	;
	if v492 != 0 {
		goto L77
	} else {
		goto L143
	}
L130:
	;
	v492 = int32(1)
	goto L129
L131:
	;
	goto L132
L132:
	;
	if v438 == int32(0) {
		v483 = v439
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v492 = v483
	goto L129
L134:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v449 < v448 {
		v483 = v439
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v451 = int32(1)
	if v448 <= v451 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v454 = v451
	goto L138
L137:
	;
	v454 = v448
	goto L138
L138:
	;
	v455 = int32(8)
	v460 = int32(0)
	goto L139
L139:
	;
	v467 = v460 << (uint(int32(2)) % 32)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v437+v455+v467)))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v467+(v438+v455))))
	v474 = v469 & (v471 ^ int32(-1))
	v476 = base.B2i32(v474 == int32(0))
	if v474 != 0 {
		v483 = v476
		goto L133
	} else {
		goto L141
	}
L140:
	;
	v483 = v476
	goto L133
L141:
	;
	v478 = v460 + int32(1)
	if v478 != v454 {
		v460 = v478
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	goto L128
L144:
	;
	if v536 != 0 {
		goto L158
	} else {
		goto L159
	}
L145:
	;
	goto L144
L146:
	;
	if v494 == int32(0) {
		v536 = v495
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	if v504 < v505 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v507 = v504
	goto L150
L149:
	;
	v507 = v505
	goto L150
L150:
	;
	if v507 <= int32(1) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v510 = int32(1)
	goto L153
L152:
	;
	v510 = v507
	goto L153
L153:
	;
	v511 = int32(8)
	v516 = int32(0)
	goto L154
L154:
	;
	v523 = v516 << (uint(int32(2)) % 32)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v494+v511+v523)))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v523+(v493+v511))))
	v528 = v525 & v527
	v530 = base.B2i32(v528 != int32(0))
	if v528 != 0 {
		v536 = v530
		goto L145
	} else {
		goto L156
	}
L155:
	;
	v536 = v530
	goto L145
L156:
	;
	v532 = v516 + int32(1)
	if v532 != v510 {
		v516 = v532
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v542 = int32(0)
	if v540 == v542 {
		v583 = v542
		goto L162
	} else {
		goto L163
	}
L159:
	;
	goto L160
L160:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v589 = int32(0)
	if v587 == v589 {
		v630 = v589
		goto L177
	} else {
		goto L178
	}
L161:
	;
	if v583 != 0 {
		goto L77
	} else {
		goto L175
	}
L162:
	;
	goto L161
L163:
	;
	if v541 == int32(0) {
		v583 = v542
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	if v551 < v552 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v554 = v551
	goto L167
L166:
	;
	v554 = v552
	goto L167
L167:
	;
	if v554 <= int32(1) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v557 = int32(1)
	goto L170
L169:
	;
	v557 = v554
	goto L170
L170:
	;
	v558 = int32(8)
	v563 = int32(0)
	goto L171
L171:
	;
	v570 = v563 << (uint(int32(2)) % 32)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v541+v558+v570)))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v570+(v540+v558))))
	v575 = v572 & v574
	v577 = base.B2i32(v575 != int32(0))
	if v575 != 0 {
		v583 = v577
		goto L162
	} else {
		goto L173
	}
L172:
	;
	v583 = v577
	goto L162
L173:
	;
	v579 = v563 + int32(1)
	if v579 != v557 {
		v563 = v579
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	goto L160
L176:
	;
	if v630 == int32(0) {
		goto L78
	} else {
		goto L190
	}
L177:
	;
	goto L176
L178:
	;
	if v588 == int32(0) {
		v630 = v589
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	if v598 < v599 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v601 = v598
	goto L182
L181:
	;
	v601 = v599
	goto L182
L182:
	;
	if v601 <= int32(1) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v604 = int32(1)
	goto L185
L184:
	;
	v604 = v601
	goto L185
L185:
	;
	v605 = int32(8)
	v610 = int32(0)
	goto L186
L186:
	;
	v617 = v610 << (uint(int32(2)) % 32)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v588+v605+v617)))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v617+(v587+v605))))
	v622 = v619 & v621
	v624 = base.B2i32(v622 != int32(0))
	if v622 != 0 {
		v630 = v624
		goto L177
	} else {
		goto L188
	}
L187:
	;
	v630 = v624
	goto L177
L188:
	;
	v626 = v610 + int32(1)
	if v626 != v604 {
		v610 = v626
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v638 = int32(0)
	if v636 == v638 {
		v679 = v638
		goto L192
	} else {
		goto L193
	}
L191:
	;
	if v679 != 0 {
		goto L77
	} else {
		goto L205
	}
L192:
	;
	goto L191
L193:
	;
	if v637 == int32(0) {
		v679 = v638
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	if v647 < v648 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v650 = v647
	goto L197
L196:
	;
	v650 = v648
	goto L197
L197:
	;
	if v650 <= int32(1) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v653 = int32(1)
	goto L200
L199:
	;
	v653 = v650
	goto L200
L200:
	;
	v654 = int32(8)
	v659 = int32(0)
	goto L201
L201:
	;
	v666 = v659 << (uint(int32(2)) % 32)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v637+v654+v666)))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v666+(v636+v654))))
	v671 = v668 & v670
	v673 = base.B2i32(v671 != int32(0))
	if v671 != 0 {
		v679 = v673
		goto L192
	} else {
		goto L203
	}
L202:
	;
	v679 = v673
	goto L192
L203:
	;
	v675 = v659 + int32(1)
	if v675 != v653 {
		v659 = v675
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	goto L78
L206:
	;
	v707 = v249
	goto L1
L207:
	;
	return int32(0)
L208:
	;
	if v687 != 0 {
		v707 = v249
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v691 = F_has_legal_joinclause(m, l0, l2)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	if v691 != 0 {
		v707 = v249
		goto L1
	} else {
		goto L211
	}
L211:
	;
	goto L32
}
func F_hindi_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L4
L1:
	;
	return v82
L2:
	;
	if v58 < int32(0) {
		v82 = v2
		goto L1
	} else {
		goto L22
	}
L4:
	;
	goto L5
L5:
	;
	goto L6
L6:
	;
	v13 = v5
	v15 = int32(1)
	goto L9
L8:
	;
	v58 = v43
	goto L2
L9:
	;
	if v6 <= v13 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v58 = int32(-1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v20 = v13 + int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+v13))))
	if base.Ui32(v22) < base.Ui32(int32(192)) {
		v43 = v20
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = int32(1)
	if v44 < v15 {
		v13 = v43
		v15 = v15 - v44
		goto L9
	} else {
		goto L21
	}
L15:
	;
	if v6 <= v20 {
		v43 = v20
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v29 = v20
	goto L17
L17:
	;
	v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4+v29))))
	if int32(-65) < v32 {
		v43 = v29
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v43 = v6
	goto L14
L19:
	;
	v36 = v29 + int32(1)
	if v36 != v6 {
		v29 = v36
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L10
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v58
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62
	v67 = F_find_among_b(m, l0, int32(4246448), int32(132))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v67 == int32(0) {
		v82 = v2
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v73
	v75 = F_slice_del(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	if v75 < int32(0) {
		v82 = v75
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
	v82 = int32(1)
	goto L1
}
func F_hmac_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = m.T0[v5].(func(*base.Module, int32) int32)(m, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		m.T0[v8].(func(*base.Module, int32))(m, v4)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
			m.T0[v12].(func(*base.Module, int32, int32, int32))(m, v4, v11, v6)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_hmac_update(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	m.T0[v5].(func(*base.Module, int32, int32, int32))(m, v4, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_hungarian_ISO_8859_2_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 < v12 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v268
	v272 = v268 - int32(1)
	if v272 <= v8 {
		goto L77
	} else {
		goto L78
	}
L2:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v257
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v163 < v12 {
		goto L45
	} else {
		goto L46
	}
L4:
	;
	if v65 != 0 {
		goto L3
	} else {
		goto L18
	}
L5:
	;
	v24 = v12
	goto L7
L6:
	;
	v24 = v22
	goto L7
L7:
	;
	goto L9
L8:
	;
	v65 = v61
	goto L4
L9:
	;
	if v12 == v24 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v61 = int32(0)
	goto L8
L11:
	;
	v65 = int32(-1)
	goto L4
L12:
	;
	goto L13
L13:
	;
	v36 = int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v12))))
	if int32(252) < v39 {
		v61 = v36
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v41 = v39 - int32(97)
	if v41 < int32(0) {
		v61 = v36
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v41)>>(uint(int32(3))%32)))+uint32(_consts[1292]))))
	if int32(base.Ui32(v47)>>(uint(v41&int32(7))%32))&int32(1) == int32(0) {
		v61 = v36
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 + int32(1)
	goto L17
L17:
	;
	goto L10
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v75 < v74 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v118 < int32(0) {
		goto L3
	} else {
		goto L33
	}
L20:
	;
	v77 = v74
	goto L22
L21:
	;
	v77 = v75
	goto L22
L22:
	;
	v84 = v74
	goto L24
L23:
	;
	v118 = int32(1)
	goto L19
L24:
	;
	if v84 == v77 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v118 = int32(-1)
	goto L19
L27:
	;
	goto L28
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v84))))
	if int32(252) < v92 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v94 = v92 - int32(97)
	if v94 < int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v94)>>(uint(int32(3))%32)))+uint32(_consts[1292]))))
	if int32(base.Ui32(v100)>>(uint(v94&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v109 = v84 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
	v84 = v109
	goto L24
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v123 = v121 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v124 <= v123 {
		v147 = v124
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v121 < v147 {
		v257 = v123
		goto L2
	} else {
		goto L43
	}
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v123))))
	if v128&int32(224) != int32(96) {
		v147 = v124
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if int32(1)<<(uint(v128)%32)&int32(101187584) == int32(0) {
		v147 = v124
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v141 = F_find_among(m, l0, int32(4176608), int32(8))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	if v141 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v257 = v145
	goto L2
L41:
	;
	goto L42
L42:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v147 = v146
	goto L34
L43:
	;
	goto L3
L44:
	;
	if v203 != 0 {
		goto L1
	} else {
		goto L59
	}
L45:
	;
	v165 = v12
	goto L47
L46:
	;
	v165 = v163
	goto L47
L47:
	;
	goto L49
L48:
	;
	v203 = v200
	goto L44
L49:
	;
	if v12 == v165 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v200 = int32(0)
	goto L48
L51:
	;
	v203 = int32(-1)
	goto L44
L52:
	;
	goto L53
L53:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v12))))
	if int32(252) < v178 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 + int32(1)
	goto L58
L55:
	;
	v180 = v178 - int32(97)
	if v180 < int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v183 = int32(1)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v180)>>(uint(int32(3))%32)))+uint32(_consts[1292]))))
	if int32(base.Ui32(v187)>>(uint(v180&int32(7))%32))&v183 != 0 {
		v200 = v183
		goto L48
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	goto L50
L59:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v212 < v211 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v252 < int32(0) {
		goto L1
	} else {
		goto L75
	}
L61:
	;
	v214 = v211
	goto L63
L62:
	;
	v214 = v212
	goto L63
L63:
	;
	v221 = v211
	goto L65
L64:
	;
	v252 = v232
	goto L60
L65:
	;
	if v221 == v214 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v252 = int32(-1)
	goto L60
L68:
	;
	goto L69
L69:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+v221))))
	if int32(252) < v227 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v244 = v221 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v244
	v221 = v244
	goto L65
L71:
	;
	v229 = v227 - int32(97)
	if v229 < int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v232 = int32(1)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v229)>>(uint(int32(3))%32)))+uint32(_consts[1292]))))
	if int32(base.Ui32(v236)>>(uint(v229&int32(7))%32))&v232 != 0 {
		goto L64
	} else {
		goto L73
	}
L73:
	;
	goto L70
L75:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v257 = v255 + v252
	goto L2
L76:
	;
	return v757
L77:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v348
	v353 = F_find_among_b(m, l0, int32(4177280), int32(44))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L38
	} else {
		goto L99
	}
L78:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v272))))
	if v276 != int32(108) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v281 = F_find_among_b(m, l0, int32(4176768), int32(2))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L38
	} else {
		goto L80
	}
L80:
	;
	if v281 == int32(0) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	if v285 < v288 {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v291 = v285 - int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v291 <= v292 {
		goto L77
	} else {
		goto L83
	}
L83:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v291))))
	if v296&int32(224) != int32(96) {
		goto L77
	} else {
		goto L84
	}
L84:
	;
	if int32(1)<<(uint(v296)%32)&int32(106790108) == int32(0) {
		goto L77
	} else {
		goto L85
	}
L85:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v310 = F_find_among_b(m, l0, int32(4176816), int32(23))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L38
	} else {
		goto L86
	}
L86:
	;
	if v310 == int32(0) {
		goto L77
	} else {
		goto L87
	}
L87:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v314 + (v285 - v307)
	v318 = F_slice_del(m, l0)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L38
	} else {
		goto L88
	}
L88:
	;
	if v318 < int32(0) {
		v757 = v318
		goto L76
	} else {
		goto L89
	}
L89:
	;
	v322 = int32(0)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v323 <= v324 {
		v341 = v322
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v341 < int32(0) {
		v757 = v341
		goto L76
	} else {
		goto L97
	}
L91:
	;
	v327 = v323 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v323 - int32(1)
	if v327 < v324 {
		v341 = v322
		goto L90
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v327
	v335 = F_slice_del(m, l0)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L38
	} else {
		goto L93
	}
L93:
	;
	if int32(0) <= v335 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v339 = int32(1)
	goto L96
L95:
	;
	v339 = v335
	goto L96
L96:
	;
	v341 = v339
	goto L90
L97:
	;
	goto L77
L98:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v404
	v408 = v404 - int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v408 <= v409 {
		goto L115
	} else {
		goto L116
	}
L99:
	;
	if v353 == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v357
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	if v357 < v360 {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v362 = F_slice_del(m, l0)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L38
	} else {
		goto L102
	}
L102:
	;
	if v362 < int32(0) {
		v757 = v362
		goto L76
	} else {
		goto L103
	}
L103:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v366 <= v368 {
		goto L98
	} else {
		goto L104
	}
L104:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370+v366-int32(1)))))
	switch v374 - int32(225) {
	case 0, 8:
		goto L105
	default:
		goto L98
	}
L105:
	;
	v379 = F_find_among_b(m, l0, int32(4178160), int32(2))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L38
	} else {
		goto L106
	}
L106:
	;
	if v379 == int32(0) {
		goto L98
	} else {
		goto L107
	}
L107:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v383 < v386 {
		goto L98
	} else {
		goto L108
	}
L108:
	;
	switch v379 - int32(1) {
	case 0:
		goto L110
	case 1:
		goto L109
	default:
		goto L98
	}
L109:
	;
	v398 = F_slice_from_s(m, l0, int32(1), int32(2141436))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L38
	} else {
		goto L113
	}
L110:
	;
	v392 = F_slice_from_s(m, l0, int32(1), int32(2141435))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L38
	} else {
		goto L111
	}
L111:
	;
	if int32(0) <= v392 {
		goto L98
	} else {
		goto L112
	}
L112:
	;
	v757 = v392
	goto L76
L113:
	;
	if v398 < int32(0) {
		v757 = v398
		goto L76
	} else {
		goto L114
	}
L114:
	;
	goto L98
L115:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v443
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v443-int32(3) <= v446 {
		goto L127
	} else {
		goto L128
	}
L116:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411+v408))))
	switch v413 - int32(110) {
	case 0, 6:
		goto L117
	default:
		goto L115
	}
L117:
	;
	v418 = F_find_among_b(m, l0, int32(4178208), int32(3))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L38
	} else {
		goto L118
	}
L118:
	;
	if v418 == int32(0) {
		goto L115
	} else {
		goto L119
	}
L119:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v422
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	if v422 < v425 {
		goto L115
	} else {
		goto L120
	}
L120:
	;
	switch v418 - int32(1) {
	case 0:
		goto L122
	case 1:
		goto L121
	default:
		goto L115
	}
L121:
	;
	v437 = F_slice_from_s(m, l0, int32(1), int32(2141440))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L38
	} else {
		goto L125
	}
L122:
	;
	v431 = F_slice_from_s(m, l0, int32(1), int32(2141439))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L38
	} else {
		goto L123
	}
L123:
	;
	if int32(0) <= v431 {
		goto L115
	} else {
		goto L124
	}
L124:
	;
	v757 = v431
	goto L76
L125:
	;
	if v437 < int32(0) {
		v757 = v437
		goto L76
	} else {
		goto L126
	}
L126:
	;
	goto L115
L127:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v488
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v488 <= v491 {
		goto L142
	} else {
		goto L143
	}
L128:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450+v443-int32(1)))))
	if v454 != int32(108) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v459 = F_find_among_b(m, l0, int32(4178272), int32(6))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L38
	} else {
		goto L130
	}
L130:
	;
	if v459 == int32(0) {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v463
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	if v463 < v466 {
		goto L127
	} else {
		goto L132
	}
L132:
	;
	switch v459 - int32(1) {
	case 0:
		goto L135
	case 1:
		goto L134
	case 2:
		goto L133
	default:
		goto L127
	}
L133:
	;
	v482 = F_slice_from_s(m, l0, int32(1), int32(2141452))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L38
	} else {
		goto L140
	}
L134:
	;
	v476 = F_slice_from_s(m, l0, int32(1), int32(2141451))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L38
	} else {
		goto L138
	}
L135:
	;
	v470 = F_slice_del(m, l0)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L38
	} else {
		goto L136
	}
L136:
	;
	if int32(0) <= v470 {
		goto L127
	} else {
		goto L137
	}
L137:
	;
	v757 = v470
	goto L76
L138:
	;
	if int32(0) <= v476 {
		goto L127
	} else {
		goto L139
	}
L139:
	;
	v757 = v476
	goto L76
L140:
	;
	if v482 < int32(0) {
		v757 = v482
		goto L76
	} else {
		goto L141
	}
L141:
	;
	goto L127
L142:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v561
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v561 <= v564 {
		goto L159
	} else {
		goto L160
	}
L143:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493+v488-int32(1)))))
	switch v497 - int32(225) {
	case 0, 8:
		goto L144
	default:
		goto L142
	}
L144:
	;
	v502 = F_find_among_b(m, l0, int32(4178400), int32(2))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L38
	} else {
		goto L145
	}
L145:
	;
	if v502 == int32(0) {
		goto L142
	} else {
		goto L146
	}
L146:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v506
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	if v506 < v509 {
		goto L142
	} else {
		goto L147
	}
L147:
	;
	v512 = v506 - int32(1)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v512 <= v513 {
		goto L142
	} else {
		goto L148
	}
L148:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515+v512))))
	if v517&int32(224) != int32(96) {
		goto L142
	} else {
		goto L149
	}
L149:
	;
	if int32(1)<<(uint(v517)%32)&int32(106790108) == int32(0) {
		goto L142
	} else {
		goto L150
	}
L150:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v531 = F_find_among_b(m, l0, int32(4176816), int32(23))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L38
	} else {
		goto L151
	}
L151:
	;
	if v531 == int32(0) {
		goto L142
	} else {
		goto L152
	}
L152:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v535 + (v506 - v528)
	v539 = F_slice_del(m, l0)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L38
	} else {
		goto L153
	}
L153:
	;
	if v539 < int32(0) {
		v757 = v539
		goto L76
	} else {
		goto L154
	}
L154:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v543 <= v544 {
		goto L142
	} else {
		goto L155
	}
L155:
	;
	v547 = v543 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v547
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v543 - int32(1)
	if v547 < v544 {
		goto L142
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v547
	v554 = F_slice_del(m, l0)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L38
	} else {
		goto L157
	}
L157:
	;
	if v554 < int32(0) {
		v757 = v554
		goto L76
	} else {
		goto L158
	}
L158:
	;
	goto L142
L159:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v606
	v608 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v606
	v613 = F_find_among_b(m, l0, int32(4178688), int32(31))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L38
	} else {
		goto L175
	}
L160:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566+v561-int32(1)))))
	if v570|int32(128) != int32(233) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v577 = F_find_among_b(m, l0, int32(4178448), int32(12))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L38
	} else {
		goto L162
	}
L162:
	;
	if v577 == int32(0) {
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v581
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	if v581 < v584 {
		goto L159
	} else {
		goto L164
	}
L164:
	;
	switch v577 - int32(1) {
	case 0:
		goto L167
	case 1:
		goto L166
	case 2:
		goto L165
	default:
		goto L159
	}
L165:
	;
	v600 = F_slice_from_s(m, l0, int32(1), int32(2141484))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L38
	} else {
		goto L172
	}
L166:
	;
	v594 = F_slice_from_s(m, l0, int32(1), int32(2141483))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L38
	} else {
		goto L170
	}
L167:
	;
	v588 = F_slice_del(m, l0)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L38
	} else {
		goto L168
	}
L168:
	;
	if int32(0) <= v588 {
		goto L159
	} else {
		goto L169
	}
L169:
	;
	v757 = v588
	goto L76
L170:
	;
	if int32(0) <= v594 {
		goto L159
	} else {
		goto L171
	}
L171:
	;
	v757 = v594
	goto L76
L172:
	;
	if v600 < int32(0) {
		v757 = v600
		goto L76
	} else {
		goto L173
	}
L173:
	;
	goto L159
L174:
	;
	if v642 < int32(0) {
		v757 = v642
		goto L76
	} else {
		goto L188
	}
L175:
	;
	if v613 == int32(0) {
		v642 = v608
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v617
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	if v617 < v620 {
		v642 = v608
		goto L174
	} else {
		goto L177
	}
L177:
	;
	switch v613 - int32(1) {
	case 0:
		goto L181
	case 1:
		goto L180
	case 2:
		goto L179
	default:
		goto L178
	}
L178:
	;
	v642 = int32(1)
	goto L174
L179:
	;
	v636 = F_slice_from_s(m, l0, int32(1), int32(2141517))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L38
	} else {
		goto L186
	}
L180:
	;
	v630 = F_slice_from_s(m, l0, int32(1), int32(2141516))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L38
	} else {
		goto L184
	}
L181:
	;
	v624 = F_slice_del(m, l0)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L38
	} else {
		goto L182
	}
L182:
	;
	if int32(0) <= v624 {
		goto L178
	} else {
		goto L183
	}
L183:
	;
	v642 = v624
	goto L174
L184:
	;
	if int32(0) <= v630 {
		goto L178
	} else {
		goto L185
	}
L185:
	;
	v642 = v630
	goto L174
L186:
	;
	if v636 < int32(0) {
		v642 = v636
		goto L174
	} else {
		goto L187
	}
L187:
	;
	goto L178
L188:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v646
	v648 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v646
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v646 <= v651 {
		v700 = v648
		goto L189
	} else {
		goto L190
	}
L189:
	;
	if v700 < int32(0) {
		v757 = v700
		goto L76
	} else {
		goto L206
	}
L190:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653+v646-int32(1)))))
	if v657&int32(224) != int32(96) {
		v700 = v648
		goto L189
	} else {
		goto L191
	}
L191:
	;
	if int32(1)<<(uint(v657)%32)&int32(10768) == int32(0) {
		v700 = v648
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v670 = F_find_among_b(m, l0, int32(4179312), int32(42))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L38
	} else {
		goto L193
	}
L193:
	;
	if v670 == int32(0) {
		v700 = v648
		goto L189
	} else {
		goto L194
	}
L194:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v674
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v676)))
	if v674 < v677 {
		v700 = v648
		goto L189
	} else {
		goto L195
	}
L195:
	;
	switch v670 - int32(1) {
	case 0:
		goto L199
	case 1:
		goto L198
	case 2:
		goto L197
	default:
		goto L196
	}
L196:
	;
	v700 = int32(1)
	goto L189
L197:
	;
	v693 = F_slice_from_s(m, l0, int32(1), int32(2141584))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L38
	} else {
		goto L204
	}
L198:
	;
	v687 = F_slice_from_s(m, l0, int32(1), int32(2141583))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L38
	} else {
		goto L202
	}
L199:
	;
	v681 = F_slice_del(m, l0)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L38
	} else {
		goto L200
	}
L200:
	;
	if int32(0) <= v681 {
		goto L196
	} else {
		goto L201
	}
L201:
	;
	v700 = v681
	goto L189
L202:
	;
	if int32(0) <= v687 {
		goto L196
	} else {
		goto L203
	}
L203:
	;
	v700 = v687
	goto L189
L204:
	;
	if v693 < int32(0) {
		v700 = v693
		goto L189
	} else {
		goto L205
	}
L205:
	;
	goto L196
L206:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v704
	v706 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v704
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v704 <= v709 {
		v749 = v706
		goto L207
	} else {
		goto L208
	}
L207:
	;
	if v749 < int32(0) {
		v757 = v749
		goto L76
	} else {
		goto L223
	}
L208:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711+v704-int32(1)))))
	if v715 != int32(107) {
		v749 = v706
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v720 = F_find_among_b(m, l0, int32(4180160), int32(7))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L38
	} else {
		goto L210
	}
L210:
	;
	if v720 == int32(0) {
		v749 = v706
		goto L207
	} else {
		goto L211
	}
L211:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v724
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	if v724 < v727 {
		v749 = v706
		goto L207
	} else {
		goto L212
	}
L212:
	;
	switch v720 - int32(1) {
	case 0:
		goto L216
	case 1:
		goto L215
	case 2:
		goto L214
	default:
		goto L213
	}
L213:
	;
	v749 = int32(1)
	goto L207
L214:
	;
	v743 = F_slice_del(m, l0)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L38
	} else {
		goto L221
	}
L215:
	;
	v739 = F_slice_from_s(m, l0, int32(1), int32(2141732))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L38
	} else {
		goto L219
	}
L216:
	;
	v733 = F_slice_from_s(m, l0, int32(1), int32(2141731))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L38
	} else {
		goto L217
	}
L217:
	;
	if int32(0) <= v733 {
		goto L213
	} else {
		goto L218
	}
L218:
	;
	v749 = v733
	goto L207
L219:
	;
	if int32(0) <= v739 {
		goto L213
	} else {
		goto L220
	}
L220:
	;
	v749 = v739
	goto L207
L221:
	;
	if v743 < int32(0) {
		v749 = v743
		goto L207
	} else {
		goto L222
	}
L222:
	;
	goto L213
L223:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v754
	v757 = int32(1)
	goto L76
}
