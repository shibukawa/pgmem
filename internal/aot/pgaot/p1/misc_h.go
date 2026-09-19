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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_HandleConcurrentAbort[0]))
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
				v9 = *(*int32)(unsafe.Add(mBase, _c_F_HandleConcurrentAbort[0]))
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
								F_errmsg(m, int32(_a_F_HandleConcurrentAbort_0), int32(0))
								mBase = m.M
								v22 = m.ExcPending
								if v22 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_HandleConcurrentAbort_1), int32(498), int32(_a_F_HandleConcurrentAbort_2))
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
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_HistoricSnapshotActive[0]))
	return base.B2i32(v2 != int32(0))
}
func F_handle_pm_pmsignal_signal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[1]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[2]))
	if v17 == v13 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[3]))
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v47 = F_pgmem_kill(m, v13, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v21 + int32(16)
	goto L1
L10:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v27)
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[4]))
	v35 = F_write(m, v31, v21+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[5]))
	if v39 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
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
	var v105 int32
	_ = v105
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
		v105 = v3
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
		goto L9
	} else {
		goto L26
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return v105
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 == int32(0) {
		v105 = v3
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 <= v18 {
		v105 = v3
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v23 = v18
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v23<<(uint(int32(2))%32))))
	v34 = F_has_dangerous_join_using(m, l0, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v105 = v34
	goto L2
L9:
	;
	return int32(0)
L10:
	;
	if v34 != 0 {
		v105 = v34
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v39 = v23 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v39 < v40 {
		v23 = v39
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v105 = int32(1)
	goto L2
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v86 = F_has_dangerous_join_using(m, l0, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L22
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v43 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48<<(uint(int32(2))%32)-int32(4))))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	if v55 <= int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v63 = int32(0)
	goto L18
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v59+v63<<(uint(int32(2))%32))))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72 != int32(6) {
		goto L13
	} else {
		goto L20
	}
L19:
	;
	goto L14
L20:
	;
	v76 = v63 + int32(1)
	if v76 != v55 {
		v63 = v76
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if v86 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v89 = F_has_dangerous_join_using(m, l0, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	if v89 == int32(0) {
		v105 = v3
		goto L2
	} else {
		goto L25
	}
L25:
	;
	goto L13
L26:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v116
	F_errmsg_internal(m, int32(_a_F_has_dangerous_join_using_0), v10)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_has_dangerous_join_using_1), int32(_a_F_has_dangerous_join_using_2), int32(_a_F_has_dangerous_join_using_3))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
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
		v14 = F_convert_any_priv_string(m, v9, int32(_a_F_has_largeobject_privilege_id_id_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v14&int64(4) == int64(0) {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_id_id[0]))
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
					v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_id_id[1])))
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
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	v17 = float64(1024)
	v21 = *(*float64)(unsafe.Add(mBase, _c_F_hashagg_spill_init[0]))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_hashagg_spill_init[1]))
	v27 = base.F64_mul(base.F64_mul(v21, base.F64_convert_i32_s(v23)), v17)
	v28 = float64(4.294967295e+09)
	if base.F64_lt(v27, v28) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v31))
	v39 = base.F64_mul(base.F64_add(base.F64_mul(v33, float64(0.25)), float64(-8192)), float64(0.0001220703125))
	v45 = base.F64_add(base.F64_div(base.F64_mul(base.F64_mul(l3, float64(1.5)), l4), v33), float64(1))
	if base.F64_gt(v45, v39) != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v31 = v27
	goto L4
L3:
	;
	v31 = v28
	goto L4
L4:
	;
	goto L1
L5:
	;
	v47 = v39
	goto L7
L6:
	;
	v47 = v45
	goto L7
L7:
	;
	if base.F64_lt(v47, float64(4)) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v50 = float64(4)
	goto L10
L9:
	;
	v50 = v47
	goto L10
L10:
	;
	if base.F64_gt(v50, float64(1024)) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v17
	goto L13
L12:
	;
	v53 = v50
	goto L13
L13:
	;
	v54 = base.I32_trunc_sat_f64_s(v53)
	v56 = int32(1073741823)
	if v56 <= v54 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if int32(31) < l2+v67 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v59 = v56
	goto L17
L16:
	;
	v59 = v54
	goto L17
L17:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v59) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = int32(32) - base.I32_clz(v59-int32(1))
	goto L20
L19:
	;
	v67 = int32(0)
	goto L20
L20:
	;
	goto L14
L21:
	;
	v71 = int32(32) - l2
	goto L23
L22:
	;
	v71 = v67
	goto L23
L23:
	;
	v73 = F_palloc0(m, int32(4)<<(uint(v71)%32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v73
	v78 = F_palloc0(m, int32(8)<<(uint(v71)%32))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78
	v83 = F_palloc0(m, int32(24)<<(uint(v71)%32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v83
	v87 = int32(1) << (uint(v71) % 32)
	v89 = base.B2i32(v71 == int32(31))
	if v89 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v92 = int32(1)
	if v87 <= v92 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v87
	v134 = int32(32)
	v136 = v134 - (l2 + v71)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v136
	v138 = int32(0)
	if v136 < v134 {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	v95 = v92
	goto L33
L32:
	;
	v95 = v87
	goto L33
L33:
	;
	v102 = int32(0)
	goto L34
L34:
	;
	v110 = F_LogicalTapeCreate(m, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L24
	} else {
		goto L36
	}
L35:
	;
	goto L30
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v112+v102<<(uint(int32(2))%32)))) = v110
	v118 = v102 + int32(1)
	if v118 != v95 {
		v102 = v118
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v145 = (v87 - int32(1)) << (uint(v136) % 32)
	goto L40
L39:
	;
	v145 = v138
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v145
	if v89 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v149 = int32(1)
	if v87 <= v149 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	return
L44:
	;
	v152 = v149
	goto L46
L45:
	;
	v152 = v87
	goto L46
L46:
	;
	v158 = v138
	goto L47
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_initHyperLogLog(m, v166+v158*int32(24), int32(5))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L24
	} else {
		goto L49
	}
L48:
	;
	goto L43
L49:
	;
	v174 = v158 + int32(1)
	if v174 != v152 {
		v158 = v174
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
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
	var v17 int32
	_ = v17
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
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
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
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
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
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
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
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
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
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L138
	}
L4:
	;
	v15 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v19 = v17 & v15
	if v19 != 0 {
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
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L133
	}
L7:
	;
	v20 = v15
	goto L9
L8:
	;
	v20 = int32(4)
	goto L9
L9:
	;
	v21 = v20 + v10
	v23 = v10 + int32(1)
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = v23
	goto L12
L11:
	;
	v26 = v10 + int32(4)
	goto L12
L12:
	;
	if v17 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v58 = v53
	goto L24
L14:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v32 == int32(18) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v43 = int32(1)
	if v19 != 0 {
		v53 = int32(base.Ui32(v17)>>(uint(v43)%32)) - v43
		goto L13
	} else {
		goto L23
	}
L17:
	;
	v35 = int32(16)
	goto L19
L18:
	;
	v35 = int32(0)
	goto L19
L19:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v42 = int32(4)
	goto L22
L21:
	;
	v42 = v35
	goto L22
L22:
	;
	v53 = v42
	goto L13
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L13
L24:
	;
	if v58 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v75 = F_pg_newlocale_from_collation(m, v14)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	v74 = v53 & (v53 >> (uint(int32(31)) % 32))
	goto L26
L28:
	;
	goto L29
L29:
	;
	v68 = v58 - int32(1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v68))))
	if v70 == int32(32) {
		v58 = v68
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v74 = v58
	goto L26
L31:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v735 != v10 {
		goto L129
	} else {
		goto L130
	}
L32:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v77 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
	v87 = v74 - int32(1636608432)
	if v81 == int64(0) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	goto L35
L35:
	;
	v399 = int32(0)
	v401 = F_pg_strnxfrm(m, v399, v399, v21, v74, v75)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L80
	}
L36:
	;
	v397 = F_Int64GetDatum(m, base.I64_extend_i32_u(v387)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v387^v379-base.I32_rotl(v387, int32(24))))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L79
	}
L37:
	;
	if v21&int32(3) != 0 {
		goto L53
	} else {
		goto L54
	}
L38:
	;
	v124 = v87
	v126 = v87
	v128 = v87
	goto L37
L39:
	;
	goto L40
L40:
	;
	v91 = v87 + base.I32_wrap_i64(v81)
	v92 = v91 + v87
	v96 = int32(4)
	v98 = base.I32_wrap_i64(int64(base.Ui64(v81)>>(uint(int64(32))%64))) ^ base.I32_rotl(v87, v96)
	v102 = v91 - v98 ^ base.I32_rotl(v98, int32(6))
	v106 = v92 - v102 ^ base.I32_rotl(v102, int32(8))
	v107 = v92 + v98
	v108 = v102 + v107
	v109 = v106 + v108
	v113 = v107 - v106 ^ base.I32_rotl(v106, int32(16))
	v117 = v108 - v113 ^ base.I32_rotl(v113, int32(19))
	v122 = v109 + v113
	v124 = v122
	v126 = v109 - v117 ^ base.I32_rotl(v117, v96)
	v128 = v117 + v122
	goto L37
L41:
	;
	v365 = int32(14)
	v367 = v361 ^ v362 - base.I32_rotl(v361, v365)
	v371 = v367 ^ v360 - base.I32_rotl(v367, int32(11))
	v375 = v371 ^ v361 - base.I32_rotl(v371, int32(25))
	v379 = v375 ^ v367 - base.I32_rotl(v375, int32(16))
	v383 = v379 ^ v371 - base.I32_rotl(v379, int32(4))
	v387 = v383 ^ v375 - base.I32_rotl(v383, v365)
	goto L36
L42:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v360 = v352 + v355
	v361 = v353
	v362 = v354
	goto L41
L43:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
	v352 = v348<<(uint(int32(8))%32) + v345
	v353 = v346
	v354 = v347
	goto L42
L44:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
	v345 = v341<<(uint(int32(16))%32) + v338
	v346 = v339
	v347 = v340
	goto L43
L45:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
	v338 = v334<<(uint(int32(24))%32) + v185
	v339 = v332
	v340 = v333
	goto L44
L46:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
	v332 = v328 + v330
	v333 = v329
	goto L45
L47:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
	v328 = v324<<(uint(int32(8))%32) + v322
	v329 = v323
	goto L46
L48:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+6)))
	v322 = v318<<(uint(int32(16))%32) + v316
	v323 = v317
	goto L47
L49:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
	v316 = v312<<(uint(int32(24))%32) + v186
	v317 = v311
	goto L48
L50:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
	v311 = v307<<(uint(int32(8))%32) + v306
	goto L49
L51:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+9)))
	v306 = v302<<(uint(int32(16))%32) + v301
	goto L50
L52:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+10)))
	v301 = v297<<(uint(int32(24))%32) + v187
	goto L51
L53:
	;
	if base.Ui32(int32(11)) < base.Ui32(v74) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v74) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	v133 = v21
	v134 = v74
	v136 = v124
	v137 = v128
	v138 = v126
	goto L59
L57:
	;
	v182 = v21
	v183 = v74
	v185 = v124
	v186 = v128
	v187 = v126
	goto L58
L58:
	;
	switch v183 - int32(1) {
	case 0:
		v352 = v185
		v353 = v186
		v354 = v187
		goto L42
	case 1:
		v345 = v185
		v346 = v186
		v347 = v187
		goto L43
	case 2:
		v338 = v185
		v339 = v186
		v340 = v187
		goto L44
	case 3:
		v332 = v186
		v333 = v187
		goto L45
	case 4:
		v328 = v186
		v329 = v187
		goto L46
	case 5:
		v322 = v186
		v323 = v187
		goto L47
	case 6:
		v316 = v186
		v317 = v187
		goto L48
	case 7:
		v311 = v187
		goto L49
	case 8:
		v306 = v187
		goto L50
	case 9:
		v301 = v187
		goto L51
	case 10:
		goto L52
	default:
		v360 = v185
		v361 = v186
		v362 = v187
		goto L41
	}
L59:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v141 = v140 + v137
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v145 = v144 + v138
	v147 = int32(4)
	v149 = v142 + v136 - v145 ^ base.I32_rotl(v145, v147)
	v153 = v141 - v149 ^ base.I32_rotl(v149, int32(6))
	v154 = v145 + v141
	v155 = v149 + v154
	v156 = v153 + v155
	v160 = v154 - v153 ^ base.I32_rotl(v153, int32(8))
	v164 = v155 - v160 ^ base.I32_rotl(v160, int32(16))
	v168 = v156 - v164 ^ base.I32_rotl(v164, int32(19))
	v169 = v160 + v156
	v170 = v164 + v169
	v171 = v168 + v170
	v175 = v169 - v168 ^ base.I32_rotl(v168, v147)
	v176 = int32(12)
	v177 = v133 + v176
	v179 = v134 - v176
	if base.Ui32(int32(11)) < base.Ui32(v179) {
		v133 = v177
		v134 = v179
		v136 = v170
		v137 = v171
		v138 = v175
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v182 = v177
	v183 = v179
	v185 = v170
	v186 = v171
	v187 = v175
	goto L58
L61:
	;
	goto L60
L62:
	;
	v193 = v21
	v194 = v74
	v196 = v124
	v197 = v128
	v198 = v126
	goto L65
L63:
	;
	v242 = v21
	v243 = v74
	v245 = v124
	v246 = v128
	v247 = v126
	goto L64
L64:
	;
	switch v243 - int32(1) {
	case 0:
		v294 = v245
		goto L68
	case 1:
		v289 = v245
		goto L69
	case 2:
		goto L70
	case 3:
		v282 = v246
		goto L71
	case 4:
		v279 = v246
		goto L72
	case 5:
		v274 = v246
		goto L73
	case 6:
		goto L74
	case 7:
		v265 = v247
		goto L75
	case 8:
		v260 = v247
		goto L76
	case 9:
		v255 = v247
		goto L77
	case 10:
		goto L78
	default:
		v360 = v245
		v361 = v246
		v362 = v247
		goto L41
	}
L65:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v201 = v200 + v197
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	v205 = v204 + v198
	v207 = int32(4)
	v209 = v202 + v196 - v205 ^ base.I32_rotl(v205, v207)
	v213 = v201 - v209 ^ base.I32_rotl(v209, int32(6))
	v214 = v205 + v201
	v215 = v209 + v214
	v216 = v213 + v215
	v220 = v214 - v213 ^ base.I32_rotl(v213, int32(8))
	v224 = v215 - v220 ^ base.I32_rotl(v220, int32(16))
	v228 = v216 - v224 ^ base.I32_rotl(v224, int32(19))
	v229 = v220 + v216
	v230 = v224 + v229
	v231 = v228 + v230
	v235 = v229 - v228 ^ base.I32_rotl(v228, v207)
	v236 = int32(12)
	v237 = v193 + v236
	v239 = v194 - v236
	if base.Ui32(int32(11)) < base.Ui32(v239) {
		v193 = v237
		v194 = v239
		v196 = v230
		v197 = v231
		v198 = v235
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v242 = v237
	v243 = v239
	v245 = v230
	v246 = v231
	v247 = v235
	goto L64
L67:
	;
	goto L66
L68:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v360 = v294 + v295
	v361 = v246
	v362 = v247
	goto L41
L69:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	v294 = v290<<(uint(int32(8))%32) + v289
	goto L68
L70:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+2)))
	v289 = v285<<(uint(int32(16))%32) + v245
	goto L69
L71:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v360 = v283 + v245
	v361 = v282
	v362 = v247
	goto L41
L72:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+4)))
	v282 = v279 + v280
	goto L71
L73:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+5)))
	v279 = v275<<(uint(int32(8))%32) + v274
	goto L72
L74:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+6)))
	v274 = v270<<(uint(int32(16))%32) + v246
	goto L73
L75:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	v360 = v266 + v245
	v361 = v268 + v246
	v362 = v265
	goto L41
L76:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+8)))
	v265 = v261<<(uint(int32(8))%32) + v260
	goto L75
L77:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+9)))
	v260 = v256<<(uint(int32(16))%32) + v255
	goto L76
L78:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+10)))
	v255 = v251<<(uint(int32(24))%32) + v247
	goto L77
L79:
	;
	v731 = v397
	goto L31
L80:
	;
	v404 = v401 + int32(1)
	v405 = F_palloc(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v407 = F_pg_strnxfrm(m, v405, v404, v21, v74, v75)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if base.Ui32(v401) < base.Ui32(v407) {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v410)))
	v417 = v404 - int32(1636608432)
	if v411 == int64(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v727 = F_Int64GetDatum(m, base.I64_extend_i32_u(v717)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v717^v709-base.I32_rotl(v717, int32(24))))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L127
	}
L85:
	;
	if v405&int32(3) != 0 {
		goto L101
	} else {
		goto L102
	}
L86:
	;
	v454 = v417
	v456 = v417
	v458 = v417
	goto L85
L87:
	;
	goto L88
L88:
	;
	v421 = v417 + base.I32_wrap_i64(v411)
	v422 = v421 + v417
	v426 = int32(4)
	v428 = base.I32_wrap_i64(int64(base.Ui64(v411)>>(uint(int64(32))%64))) ^ base.I32_rotl(v417, v426)
	v432 = v421 - v428 ^ base.I32_rotl(v428, int32(6))
	v436 = v422 - v432 ^ base.I32_rotl(v432, int32(8))
	v437 = v422 + v428
	v438 = v432 + v437
	v439 = v436 + v438
	v443 = v437 - v436 ^ base.I32_rotl(v436, int32(16))
	v447 = v438 - v443 ^ base.I32_rotl(v443, int32(19))
	v452 = v439 + v443
	v454 = v452
	v456 = v439 - v447 ^ base.I32_rotl(v447, v426)
	v458 = v447 + v452
	goto L85
L89:
	;
	v695 = int32(14)
	v697 = v691 ^ v692 - base.I32_rotl(v691, v695)
	v701 = v697 ^ v690 - base.I32_rotl(v697, int32(11))
	v705 = v701 ^ v691 - base.I32_rotl(v701, int32(25))
	v709 = v705 ^ v697 - base.I32_rotl(v705, int32(16))
	v713 = v709 ^ v701 - base.I32_rotl(v709, int32(4))
	v717 = v713 ^ v705 - base.I32_rotl(v713, v695)
	goto L84
L90:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	v690 = v682 + v685
	v691 = v683
	v692 = v684
	goto L89
L91:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+1)))
	v682 = v678<<(uint(int32(8))%32) + v675
	v683 = v676
	v684 = v677
	goto L90
L92:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+2)))
	v675 = v671<<(uint(int32(16))%32) + v668
	v676 = v669
	v677 = v670
	goto L91
L93:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+3)))
	v668 = v664<<(uint(int32(24))%32) + v515
	v669 = v662
	v670 = v663
	goto L92
L94:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+4)))
	v662 = v658 + v660
	v663 = v659
	goto L93
L95:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+5)))
	v658 = v654<<(uint(int32(8))%32) + v652
	v659 = v653
	goto L94
L96:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+6)))
	v652 = v648<<(uint(int32(16))%32) + v646
	v653 = v647
	goto L95
L97:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+7)))
	v646 = v642<<(uint(int32(24))%32) + v516
	v647 = v641
	goto L96
L98:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+8)))
	v641 = v637<<(uint(int32(8))%32) + v636
	goto L97
L99:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+9)))
	v636 = v632<<(uint(int32(16))%32) + v631
	goto L98
L100:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+10)))
	v631 = v627<<(uint(int32(24))%32) + v517
	goto L99
L101:
	;
	if base.Ui32(int32(11)) < base.Ui32(v404) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v404) {
		goto L110
	} else {
		goto L111
	}
L104:
	;
	v463 = v405
	v464 = v404
	v466 = v454
	v467 = v458
	v468 = v456
	goto L107
L105:
	;
	v512 = v405
	v513 = v404
	v515 = v454
	v516 = v458
	v517 = v456
	goto L106
L106:
	;
	switch v513 - int32(1) {
	case 0:
		v682 = v515
		v683 = v516
		v684 = v517
		goto L90
	case 1:
		v675 = v515
		v676 = v516
		v677 = v517
		goto L91
	case 2:
		v668 = v515
		v669 = v516
		v670 = v517
		goto L92
	case 3:
		v662 = v516
		v663 = v517
		goto L93
	case 4:
		v658 = v516
		v659 = v517
		goto L94
	case 5:
		v652 = v516
		v653 = v517
		goto L95
	case 6:
		v646 = v516
		v647 = v517
		goto L96
	case 7:
		v641 = v517
		goto L97
	case 8:
		v636 = v517
		goto L98
	case 9:
		v631 = v517
		goto L99
	case 10:
		goto L100
	default:
		v690 = v515
		v691 = v516
		v692 = v517
		goto L89
	}
L107:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	v471 = v470 + v467
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v463)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v463)+8))
	v475 = v474 + v468
	v477 = int32(4)
	v479 = v472 + v466 - v475 ^ base.I32_rotl(v475, v477)
	v483 = v471 - v479 ^ base.I32_rotl(v479, int32(6))
	v484 = v475 + v471
	v485 = v479 + v484
	v486 = v483 + v485
	v490 = v484 - v483 ^ base.I32_rotl(v483, int32(8))
	v494 = v485 - v490 ^ base.I32_rotl(v490, int32(16))
	v498 = v486 - v494 ^ base.I32_rotl(v494, int32(19))
	v499 = v490 + v486
	v500 = v494 + v499
	v501 = v498 + v500
	v505 = v499 - v498 ^ base.I32_rotl(v498, v477)
	v506 = int32(12)
	v507 = v463 + v506
	v509 = v464 - v506
	if base.Ui32(int32(11)) < base.Ui32(v509) {
		v463 = v507
		v464 = v509
		v466 = v500
		v467 = v501
		v468 = v505
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v512 = v507
	v513 = v509
	v515 = v500
	v516 = v501
	v517 = v505
	goto L106
L109:
	;
	goto L108
L110:
	;
	v523 = v405
	v524 = v404
	v526 = v454
	v527 = v458
	v528 = v456
	goto L113
L111:
	;
	v572 = v405
	v573 = v404
	v575 = v454
	v576 = v458
	v577 = v456
	goto L112
L112:
	;
	switch v573 - int32(1) {
	case 0:
		v624 = v575
		goto L116
	case 1:
		v619 = v575
		goto L117
	case 2:
		goto L118
	case 3:
		v612 = v576
		goto L119
	case 4:
		v609 = v576
		goto L120
	case 5:
		v604 = v576
		goto L121
	case 6:
		goto L122
	case 7:
		v595 = v577
		goto L123
	case 8:
		v590 = v577
		goto L124
	case 9:
		v585 = v577
		goto L125
	case 10:
		goto L126
	default:
		v690 = v575
		v691 = v576
		v692 = v577
		goto L89
	}
L113:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	v531 = v530 + v527
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v523)+8))
	v535 = v534 + v528
	v537 = int32(4)
	v539 = v532 + v526 - v535 ^ base.I32_rotl(v535, v537)
	v543 = v531 - v539 ^ base.I32_rotl(v539, int32(6))
	v544 = v535 + v531
	v545 = v539 + v544
	v546 = v543 + v545
	v550 = v544 - v543 ^ base.I32_rotl(v543, int32(8))
	v554 = v545 - v550 ^ base.I32_rotl(v550, int32(16))
	v558 = v546 - v554 ^ base.I32_rotl(v554, int32(19))
	v559 = v550 + v546
	v560 = v554 + v559
	v561 = v558 + v560
	v565 = v559 - v558 ^ base.I32_rotl(v558, v537)
	v566 = int32(12)
	v567 = v523 + v566
	v569 = v524 - v566
	if base.Ui32(int32(11)) < base.Ui32(v569) {
		v523 = v567
		v524 = v569
		v526 = v560
		v527 = v561
		v528 = v565
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v572 = v567
	v573 = v569
	v575 = v560
	v576 = v561
	v577 = v565
	goto L112
L115:
	;
	goto L114
L116:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572))))
	v690 = v624 + v625
	v691 = v576
	v692 = v577
	goto L89
L117:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+1)))
	v624 = v620<<(uint(int32(8))%32) + v619
	goto L116
L118:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+2)))
	v619 = v615<<(uint(int32(16))%32) + v575
	goto L117
L119:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v690 = v613 + v575
	v691 = v612
	v692 = v577
	goto L89
L120:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+4)))
	v612 = v609 + v610
	goto L119
L121:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+5)))
	v609 = v605<<(uint(int32(8))%32) + v604
	goto L120
L122:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+6)))
	v604 = v600<<(uint(int32(16))%32) + v576
	goto L121
L123:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v690 = v596 + v575
	v691 = v598 + v576
	v692 = v595
	goto L89
L124:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+8)))
	v595 = v591<<(uint(int32(8))%32) + v590
	goto L123
L125:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+9)))
	v590 = v586<<(uint(int32(16))%32) + v585
	goto L124
L126:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+10)))
	v585 = v581<<(uint(int32(24))%32) + v577
	goto L125
L127:
	;
	F_pfree(m, v405)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v731 = v727
	goto L31
L129:
	;
	F_pfree(m, v10)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	return v731
L132:
	;
	goto L131
L133:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(_a_F_hashbpcharextended_0), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errhint(m, int32(_a_F_hashbpcharextended_1), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_hashbpcharextended_2), int32(1057), int32(_a_F_hashbpcharextended_3))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_hashbpcharextended_4), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_hashbpcharextended_2), int32(1082), int32(_a_F_hashbpcharextended_3))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v40 float64
	_ = v40
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v14
	F_genericcostestimate(m, l0, l1, l2, v12)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return
	} else {
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = v32
		v34 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v34
		v36 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v36
		v38 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v38
		v40 = *(*float64)(unsafe.Add(mBase, uint32(v12)+32))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v40
		m.G0 = v12 - int32(-64)
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
			v55 = v40 + v46
			v56 = v50 + v55
			v57 = v54 + v56
			v61 = v55 - v54 ^ base.I32_rotl(v54, int32(16))
			v65 = v56 - v61 ^ base.I32_rotl(v61, int32(19))
			v70 = v57 + v61
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
		v345 = F_Int64GetDatum(m, base.I64_extend_i32_u(v335)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v335^v327-base.I32_rotl(v335, int32(24))))
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
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(12)
	v16 = v11 + int32(11)
	v17 = F__hash_convert_tuple(m, l0, l1, l2, v14, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v22 = F_index_form_tuple(m, v21, v14, v16)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)) = uint16(v24)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v26
				F__hash_doinsert(m, l0, v22, l4, int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v22)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
		v31 = v16 + v22
		v32 = v26 + v31
		v33 = v30 + v32
		v37 = v31 - v30 ^ base.I32_rotl(v30, int32(16))
		v41 = v32 - v37 ^ base.I32_rotl(v37, int32(19))
		v46 = v33 + v37
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
	v321 = F_Int64GetDatum(m, base.I64_extend_i32_u(v311)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v311^v303-base.I32_rotl(v311, int32(24))))
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
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
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
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
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
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_strlen(m, v2)
	mBase = m.M
	v9 = v3 - int32(1636608432)
	if v2&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v3) {
			v118 = v2
			v119 = v3
			v120 = v9
			v121 = v9
			v122 = v9
			for {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
				v125 = v124 + v121
				v126 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
				v129 = v128 + v122
				v131 = int32(4)
				v133 = v126 + v120 - v129 ^ base.I32_rotl(v129, v131)
				v137 = v125 - v133 ^ base.I32_rotl(v133, int32(6))
				v138 = v129 + v125
				v139 = v133 + v138
				v140 = v137 + v139
				v144 = v138 - v137 ^ base.I32_rotl(v137, int32(8))
				v148 = v139 - v144 ^ base.I32_rotl(v144, int32(16))
				v152 = v140 - v148 ^ base.I32_rotl(v148, int32(19))
				v153 = v144 + v140
				v154 = v148 + v153
				v155 = v152 + v154
				v159 = v153 - v152 ^ base.I32_rotl(v152, v131)
				v160 = int32(12)
				v161 = v118 + v160
				v163 = v119 - v160
				if base.Ui32(int32(11)) < base.Ui32(v163) {
					v118 = v161
					v119 = v163
					v120 = v154
					v121 = v155
					v122 = v159
					continue
				} else {
					break
				}
				break
			}
			v166 = v161
			v167 = v163
			v168 = v154
			v169 = v155
			v170 = v159
		} else {
			v166 = v2
			v167 = v3
			v168 = v9
			v169 = v9
			v170 = v9
		}
		switch v167 - int32(1) {
		case 0:
			v229 = v168
			v230 = v169
			v231 = v170
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 1:
			v222 = v168
			v223 = v169
			v224 = v170
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 2:
			v215 = v168
			v216 = v169
			v217 = v170
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 3:
			v209 = v169
			v210 = v170
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)))
			v215 = v211<<(uint(int32(24))%32) + v168
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 4:
			v205 = v169
			v206 = v170
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)))
			v215 = v211<<(uint(int32(24))%32) + v168
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 5:
			v199 = v169
			v200 = v170
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)))
			v215 = v211<<(uint(int32(24))%32) + v168
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 6:
			v193 = v169
			v194 = v170
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)))
			v215 = v211<<(uint(int32(24))%32) + v168
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 7:
			v188 = v170
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+7)))
			v193 = v189<<(uint(int32(24))%32) + v169
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)))
			v215 = v211<<(uint(int32(24))%32) + v168
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 8:
			v183 = v170
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+7)))
			v193 = v189<<(uint(int32(24))%32) + v169
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)))
			v215 = v211<<(uint(int32(24))%32) + v168
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 9:
			v178 = v170
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+7)))
			v193 = v189<<(uint(int32(24))%32) + v169
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)))
			v215 = v211<<(uint(int32(24))%32) + v168
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 10:
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+10)))
			v178 = v174<<(uint(int32(24))%32) + v170
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+7)))
			v193 = v189<<(uint(int32(24))%32) + v169
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)))
			v215 = v211<<(uint(int32(24))%32) + v168
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		default:
			v236 = v168
			v237 = v169
			v238 = v170
		}
	} else {
		if base.Ui32(v3) < base.Ui32(int32(12)) {
			v64 = v2
			v65 = v3
			v66 = v9
			v67 = v9
			v68 = v9
		} else {
			v16 = v2
			v17 = v3
			v18 = v9
			v19 = v9
			v20 = v9
			for {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				v23 = v22 + v19
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
				v27 = v26 + v20
				v29 = int32(4)
				v31 = v24 + v18 - v27 ^ base.I32_rotl(v27, v29)
				v35 = v23 - v31 ^ base.I32_rotl(v31, int32(6))
				v36 = v27 + v23
				v37 = v31 + v36
				v38 = v35 + v37
				v42 = v36 - v35 ^ base.I32_rotl(v35, int32(8))
				v46 = v37 - v42 ^ base.I32_rotl(v42, int32(16))
				v50 = v38 - v46 ^ base.I32_rotl(v46, int32(19))
				v51 = v42 + v38
				v52 = v46 + v51
				v53 = v50 + v52
				v57 = v51 - v50 ^ base.I32_rotl(v50, v29)
				v58 = int32(12)
				v59 = v16 + v58
				v61 = v17 - v58
				if base.Ui32(int32(11)) < base.Ui32(v61) {
					v16 = v59
					v17 = v61
					v18 = v52
					v19 = v53
					v20 = v57
					continue
				} else {
					break
				}
				break
			}
			v64 = v59
			v65 = v61
			v66 = v52
			v67 = v53
			v68 = v57
		}
		switch v65 - int32(1) {
		case 0:
			v115 = v66
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
			v236 = v115 + v116
			v237 = v67
			v238 = v68
		case 1:
			v110 = v66
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
			v236 = v115 + v116
			v237 = v67
			v238 = v68
		case 2:
			v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+2)))
			v110 = v106<<(uint(int32(16))%32) + v66
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
			v236 = v115 + v116
			v237 = v67
			v238 = v68
		case 3:
			v103 = v67
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v236 = v104 + v66
			v237 = v103
			v238 = v68
		case 4:
			v100 = v67
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v236 = v104 + v66
			v237 = v103
			v238 = v68
		case 5:
			v95 = v67
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v236 = v104 + v66
			v237 = v103
			v238 = v68
		case 6:
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+6)))
			v95 = v91<<(uint(int32(16))%32) + v67
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v236 = v104 + v66
			v237 = v103
			v238 = v68
		case 7:
			v86 = v68
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
			v236 = v87 + v66
			v237 = v89 + v67
			v238 = v86
		case 8:
			v81 = v68
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
			v236 = v87 + v66
			v237 = v89 + v67
			v238 = v86
		case 9:
			v76 = v68
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
			v236 = v87 + v66
			v237 = v89 + v67
			v238 = v86
		case 10:
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+10)))
			v76 = v72<<(uint(int32(24))%32) + v68
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
			v236 = v87 + v66
			v237 = v89 + v67
			v238 = v86
		default:
			v236 = v66
			v237 = v67
			v238 = v68
		}
	}
	v241 = int32(14)
	v243 = v237 ^ v238 - base.I32_rotl(v237, v241)
	v247 = v243 ^ v236 - base.I32_rotl(v243, int32(11))
	v251 = v247 ^ v237 - base.I32_rotl(v247, int32(25))
	v255 = v251 ^ v243 - base.I32_rotl(v251, int32(16))
	v259 = v255 ^ v247 - base.I32_rotl(v255, int32(4))
	v263 = v259 ^ v251 - base.I32_rotl(v259, v241)
	return v263 ^ v255 - base.I32_rotl(v263, int32(24))
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
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
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
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
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
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v711 int32
	_ = v711
	v4 = int32(0)
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	if base.B2i32(v9 == v4)|base.B2i32(v10 == v4) != 0 {
		v56 = v4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v711
L2:
	;
	if v56 != 0 {
		v711 = v8
		goto L1
	} else {
		goto L15
	}
L3:
	;
	goto L2
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v21 < v22 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = v21
	goto L7
L6:
	;
	v24 = v22
	goto L7
L7:
	;
	if v24 <= int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = int32(1)
	goto L10
L9:
	;
	v27 = v24
	goto L10
L10:
	;
	v28 = int32(8)
	v33 = int32(0)
	goto L11
L11:
	;
	v40 = v33 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10+v28+v40)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v9+v28+v40)))
	v45 = v42 & v44
	v47 = base.B2i32(v45 != int32(0))
	if v45 != 0 {
		v56 = v47
		goto L3
	} else {
		goto L13
	}
L12:
	;
	v56 = v47
	goto L3
L13:
	;
	v49 = v33 + int32(1)
	if v49 != v27 {
		v33 = v49
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v59 = int32(0)
	if base.B2i32(v57 == v59)|base.B2i32(v58 == v59) != 0 {
		v104 = v59
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v104 != 0 {
		v711 = v8
		goto L1
	} else {
		goto L29
	}
L17:
	;
	goto L16
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v69 < v70 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = v69
	goto L21
L20:
	;
	v72 = v70
	goto L21
L21:
	;
	if v72 <= int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v75 = int32(1)
	goto L24
L23:
	;
	v75 = v72
	goto L24
L24:
	;
	v76 = int32(8)
	v81 = int32(0)
	goto L25
L25:
	;
	v88 = v81 << (uint(int32(2)) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v58+v76+v88)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v57+v76+v88)))
	v93 = v90 & v92
	v95 = base.B2i32(v93 != int32(0))
	if v93 != 0 {
		v104 = v95
		goto L17
	} else {
		goto L27
	}
L26:
	;
	v104 = v95
	goto L17
L27:
	;
	v97 = v81 + int32(1)
	if v97 != v75 {
		v81 = v97
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v105 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v711 = int32(1)
	goto L1
L31:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v246 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v108 <= int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v115 = v4
	goto L34
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v115<<(uint(int32(2))%32))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v125 = int32(0)
	if v118 == v125 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L31
L36:
	;
	if v178 != 0 {
		goto L50
	} else {
		goto L51
	}
L37:
	;
	v178 = int32(1)
	goto L36
L38:
	;
	goto L39
L39:
	;
	if v124 == int32(0) {
		v171 = v125
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v178 = v171
	goto L36
L41:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v135 < v134 {
		v171 = v125
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v137 = int32(1)
	if v134 <= v137 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v140 = v137
	goto L45
L44:
	;
	v140 = v134
	goto L45
L45:
	;
	v141 = int32(8)
	v146 = int32(0)
	goto L46
L46:
	;
	v153 = v146 << (uint(int32(2)) % 32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v118+v141+v153)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v124+v141+v153)))
	v160 = v155 & (v157 ^ int32(-1))
	v162 = base.B2i32(v160 == int32(0))
	if v160 != 0 {
		v171 = v162
		goto L40
	} else {
		goto L48
	}
L47:
	;
	v171 = v162
	goto L40
L48:
	;
	v164 = v146 + int32(1)
	if v164 != v140 {
		v146 = v164
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v181 = int32(0)
	if v179 == v181 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	goto L52
L52:
	;
	v236 = v115 + int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v236 < v237 {
		v115 = v236
		goto L34
	} else {
		goto L68
	}
L53:
	;
	if v234 != 0 {
		goto L30
	} else {
		goto L67
	}
L54:
	;
	v234 = int32(1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	if v180 == int32(0) {
		v227 = v181
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v234 = v227
	goto L53
L58:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v191 < v190 {
		v227 = v181
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v193 = int32(1)
	if v190 <= v193 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v196 = v193
	goto L62
L61:
	;
	v196 = v190
	goto L62
L62:
	;
	v197 = int32(8)
	v202 = int32(0)
	goto L63
L63:
	;
	v209 = v202 << (uint(int32(2)) % 32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v179+v197+v209)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v180+v197+v209)))
	v216 = v211 & (v213 ^ int32(-1))
	v218 = base.B2i32(v216 == int32(0))
	if v216 != 0 {
		v227 = v218
		goto L57
	} else {
		goto L65
	}
L64:
	;
	v227 = v218
	goto L57
L65:
	;
	v220 = v202 + int32(1)
	if v220 != v196 {
		v202 = v220
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L52
L68:
	;
	goto L35
L69:
	;
	return int32(0)
L70:
	;
	goto L71
L71:
	;
	v251 = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v252 <= v251 {
		v711 = v251
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v261 = v4
	goto L73
L73:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v263 = int32(2)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262+v261<<(uint(v263)%32))))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	if v267 == v263 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v692 = F_has_legal_joinclause(m, l0, l1)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L201
	} else {
		goto L202
	}
L75:
	;
	goto L74
L76:
	;
	v689 = v261 + int32(1)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v689 < v690 {
		v261 = v689
		goto L73
	} else {
		goto L200
	}
L77:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v272 = int32(0)
	if v270 == v272 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v325 != 0 {
		goto L92
	} else {
		goto L93
	}
L79:
	;
	v325 = int32(1)
	goto L78
L80:
	;
	goto L81
L81:
	;
	if v271 == int32(0) {
		v318 = v272
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v325 = v318
	goto L78
L83:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if v282 < v281 {
		v318 = v272
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v284 = int32(1)
	if v281 <= v284 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v287 = v284
	goto L87
L86:
	;
	v287 = v281
	goto L87
L87:
	;
	v288 = int32(8)
	v293 = int32(0)
	goto L88
L88:
	;
	v300 = v293 << (uint(int32(2)) % 32)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v270+v288+v300)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v271+v288+v300)))
	v307 = v302 & (v304 ^ int32(-1))
	v309 = base.B2i32(v307 == int32(0))
	if v307 != 0 {
		v318 = v309
		goto L82
	} else {
		goto L90
	}
L89:
	;
	v318 = v309
	goto L82
L90:
	;
	v311 = v293 + int32(1)
	if v311 != v287 {
		v293 = v311
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v328 = int32(0)
	if v326 == v328 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	goto L94
L94:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v384 = int32(0)
	if v382 == v384 {
		goto L111
	} else {
		goto L112
	}
L95:
	;
	if v381 != 0 {
		goto L75
	} else {
		goto L109
	}
L96:
	;
	v381 = int32(1)
	goto L95
L97:
	;
	goto L98
L98:
	;
	if v327 == int32(0) {
		v374 = v328
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v381 = v374
	goto L95
L100:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v338 < v337 {
		v374 = v328
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v340 = int32(1)
	if v337 <= v340 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v343 = v340
	goto L104
L103:
	;
	v343 = v337
	goto L104
L104:
	;
	v344 = int32(8)
	v349 = int32(0)
	goto L105
L105:
	;
	v356 = v349 << (uint(int32(2)) % 32)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v326+v344+v356)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v327+v344+v356)))
	v363 = v358 & (v360 ^ int32(-1))
	v365 = base.B2i32(v363 == int32(0))
	if v363 != 0 {
		v374 = v365
		goto L99
	} else {
		goto L107
	}
L106:
	;
	v374 = v365
	goto L99
L107:
	;
	v367 = v349 + int32(1)
	if v367 != v343 {
		v349 = v367
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	goto L94
L110:
	;
	if v437 != 0 {
		goto L124
	} else {
		goto L125
	}
L111:
	;
	v437 = int32(1)
	goto L110
L112:
	;
	goto L113
L113:
	;
	if v383 == int32(0) {
		v430 = v384
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v437 = v430
	goto L110
L115:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v394 < v393 {
		v430 = v384
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v396 = int32(1)
	if v393 <= v396 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v399 = v396
	goto L119
L118:
	;
	v399 = v393
	goto L119
L119:
	;
	v400 = int32(8)
	v405 = int32(0)
	goto L120
L120:
	;
	v412 = v405 << (uint(int32(2)) % 32)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v382+v400+v412)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v383+v400+v412)))
	v419 = v414 & (v416 ^ int32(-1))
	v421 = base.B2i32(v419 == int32(0))
	if v419 != 0 {
		v430 = v421
		goto L114
	} else {
		goto L122
	}
L121:
	;
	v430 = v421
	goto L114
L122:
	;
	v423 = v405 + int32(1)
	if v423 != v399 {
		v405 = v423
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v440 = int32(0)
	if v438 == v440 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	goto L126
L126:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v496 = int32(0)
	if base.B2i32(v494 == v496)|base.B2i32(v495 == v496) != 0 {
		v541 = v496
		goto L143
	} else {
		goto L144
	}
L127:
	;
	if v493 != 0 {
		goto L75
	} else {
		goto L141
	}
L128:
	;
	v493 = int32(1)
	goto L127
L129:
	;
	goto L130
L130:
	;
	if v439 == int32(0) {
		v486 = v440
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v493 = v486
	goto L127
L132:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v450 < v449 {
		v486 = v440
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v452 = int32(1)
	if v449 <= v452 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v455 = v452
	goto L136
L135:
	;
	v455 = v449
	goto L136
L136:
	;
	v456 = int32(8)
	v461 = int32(0)
	goto L137
L137:
	;
	v468 = v461 << (uint(int32(2)) % 32)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v438+v456+v468)))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v439+v456+v468)))
	v475 = v470 & (v472 ^ int32(-1))
	v477 = base.B2i32(v475 == int32(0))
	if v475 != 0 {
		v486 = v477
		goto L131
	} else {
		goto L139
	}
L138:
	;
	v486 = v477
	goto L131
L139:
	;
	v479 = v461 + int32(1)
	if v479 != v455 {
		v461 = v479
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	goto L126
L142:
	;
	if v541 != 0 {
		goto L155
	} else {
		goto L156
	}
L143:
	;
	goto L142
L144:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	if v506 < v507 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v509 = v506
	goto L147
L146:
	;
	v509 = v507
	goto L147
L147:
	;
	if v509 <= int32(1) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v512 = int32(1)
	goto L150
L149:
	;
	v512 = v509
	goto L150
L150:
	;
	v513 = int32(8)
	v518 = int32(0)
	goto L151
L151:
	;
	v525 = v518 << (uint(int32(2)) % 32)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v495+v513+v525)))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v494+v513+v525)))
	v530 = v527 & v529
	v532 = base.B2i32(v530 != int32(0))
	if v530 != 0 {
		v541 = v532
		goto L143
	} else {
		goto L153
	}
L152:
	;
	v541 = v532
	goto L143
L153:
	;
	v534 = v518 + int32(1)
	if v534 != v512 {
		v518 = v534
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v544 = int32(0)
	if base.B2i32(v542 == v544)|base.B2i32(v543 == v544) != 0 {
		v589 = v544
		goto L159
	} else {
		goto L160
	}
L156:
	;
	goto L157
L157:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v592 = int32(0)
	if base.B2i32(v590 == v592)|base.B2i32(v591 == v592) != 0 {
		v637 = v592
		goto L173
	} else {
		goto L174
	}
L158:
	;
	if v589 != 0 {
		goto L75
	} else {
		goto L171
	}
L159:
	;
	goto L158
L160:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	if v554 < v555 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v557 = v554
	goto L163
L162:
	;
	v557 = v555
	goto L163
L163:
	;
	if v557 <= int32(1) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v560 = int32(1)
	goto L166
L165:
	;
	v560 = v557
	goto L166
L166:
	;
	v561 = int32(8)
	v566 = int32(0)
	goto L167
L167:
	;
	v573 = v566 << (uint(int32(2)) % 32)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v543+v561+v573)))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v542+v561+v573)))
	v578 = v575 & v577
	v580 = base.B2i32(v578 != int32(0))
	if v578 != 0 {
		v589 = v580
		goto L159
	} else {
		goto L169
	}
L168:
	;
	v589 = v580
	goto L159
L169:
	;
	v582 = v566 + int32(1)
	if v582 != v560 {
		v566 = v582
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	goto L157
L172:
	;
	if v637 == int32(0) {
		goto L76
	} else {
		goto L185
	}
L173:
	;
	goto L172
L174:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v590)+4))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	if v602 < v603 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v605 = v602
	goto L177
L176:
	;
	v605 = v603
	goto L177
L177:
	;
	if v605 <= int32(1) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v608 = int32(1)
	goto L180
L179:
	;
	v608 = v605
	goto L180
L180:
	;
	v609 = int32(8)
	v614 = int32(0)
	goto L181
L181:
	;
	v621 = v614 << (uint(int32(2)) % 32)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v591+v609+v621)))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v590+v609+v621)))
	v626 = v623 & v625
	v628 = base.B2i32(v626 != int32(0))
	if v626 != 0 {
		v637 = v628
		goto L173
	} else {
		goto L183
	}
L182:
	;
	v637 = v628
	goto L173
L183:
	;
	v630 = v614 + int32(1)
	if v630 != v608 {
		v614 = v630
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v642 = int32(0)
	if base.B2i32(v640 == v642)|base.B2i32(v641 == v642) != 0 {
		v687 = v642
		goto L187
	} else {
		goto L188
	}
L186:
	;
	if v687 != 0 {
		goto L75
	} else {
		goto L199
	}
L187:
	;
	goto L186
L188:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	if v652 < v653 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v655 = v652
	goto L191
L190:
	;
	v655 = v653
	goto L191
L191:
	;
	if v655 <= int32(1) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v658 = int32(1)
	goto L194
L193:
	;
	v658 = v655
	goto L194
L194:
	;
	v659 = int32(8)
	v664 = int32(0)
	goto L195
L195:
	;
	v671 = v664 << (uint(int32(2)) % 32)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v641+v659+v671)))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v640+v659+v671)))
	v676 = v673 & v675
	v678 = base.B2i32(v676 != int32(0))
	if v676 != 0 {
		v687 = v678
		goto L187
	} else {
		goto L197
	}
L196:
	;
	v687 = v678
	goto L187
L197:
	;
	v680 = v664 + int32(1)
	if v680 != v658 {
		v664 = v680
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	goto L76
L200:
	;
	v711 = v251
	goto L1
L201:
	;
	return int32(0)
L202:
	;
	if v692 != 0 {
		v711 = v251
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v696 = F_has_legal_joinclause(m, l0, l2)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L201
	} else {
		goto L204
	}
L204:
	;
	if v696 != 0 {
		v711 = v251
		goto L1
	} else {
		goto L205
	}
L205:
	;
	goto L30
}
func F_heapcheck_read_stream_next_unskippable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	goto L1
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9 + int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v13) <= base.Ui32(v9) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v9
L3:
	;
	return int32(-1)
L4:
	;
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = F_visibilitymap_get_status(m, v17, v9, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v19&int32(2) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = v23
	goto L10
L9:
	;
	v27 = int32(1)
	goto L10
L10:
	;
	v30 = int32(1)
	if base.B2i32(v27 == int32(0))|v19&v30&base.B2i32(v23 == v30) != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
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
	var v83 int32
	_ = v83
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L4
L1:
	;
	return v83
L2:
	;
	if v58 < int32(0) {
		v83 = v2
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
	v67 = F_find_among_b(m, l0, int32(_a_F_hindi_UTF_8_stem_0), int32(132))
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
		v83 = v2
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
		v83 = v75
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
	v83 = int32(1)
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
func F_hnswinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v13 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0]))
		v22 = F_AllocSetContextCreateInternal(m, v17, int32(_a_F_hnswinsert_0), int32(0), int32(_a_F_hnswinsert_1), int32(_a_F_hnswinsert_2))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(_a_F_hnswinsert_3)
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0])) = v22
			v30 = F_HnswGetTypeInfo(m, l0)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_HnswInitSupport(m, v11, l0)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v36 = F_HnswFormIndexValue(m, v11+int32(12), l1, l2, v30, v11)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v36 != 0 {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							v40 = F_HnswInsertTupleOnDisk(m, l0, v11, v38, l3, int32(0))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0])) = v27
								F_MemoryContextDelete(m, v22)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(16)
									return int32(0)
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0])) = v27
							F_MemoryContextDelete(m, v22)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(16)
								return int32(0)
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v11 + int32(16)
		return int32(0)
	}
}
func F_hnswvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	if l1 == int32(0) {
		return l1
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v5&int32(1) != 0 {
			return l1
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v10 = F_RelationGetNumberOfBlocksInFork(m, v8, int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10
				return l1
			}
		}
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
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
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
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
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
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
	v273 = v269 - int32(1)
	if v273 <= v8 {
		goto L76
	} else {
		goto L77
	}
L2:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v258
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v164 < v12 {
		goto L44
	} else {
		goto L45
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
	v65 = v60
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
	v60 = int32(0)
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
		v60 = v36
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v41 = v39 - int32(97)
	if v41 < int32(0) {
		v60 = v36
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v41)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_ISO_8859_2_stem[0]))))
	if int32(base.Ui32(v47)>>(uint(v41&int32(7))%32))&int32(1) == int32(0) {
		v60 = v36
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
	v83 = v74
	goto L24
L23:
	;
	v118 = int32(1)
	goto L19
L24:
	;
	if v83 == v77 {
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
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v83))))
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
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v94)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_ISO_8859_2_stem[0]))))
	if int32(base.Ui32(v100)>>(uint(v94&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v109 = v83 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
	v83 = v109
	goto L24
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v123 = v121 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v124 <= v123 {
		v148 = v124
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v121 < v148 {
		v258 = v123
		goto L2
	} else {
		goto L42
	}
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v123))))
	if base.B2i32(v128&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v128)%32)&int32(101187584) == int32(0)) != 0 {
		v148 = v124
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v142 = F_find_among(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_0), int32(8))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(0)
L38:
	;
	if v142 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v258 = v146
	goto L2
L40:
	;
	goto L41
L41:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v148 = v147
	goto L34
L42:
	;
	goto L3
L43:
	;
	if v204 != 0 {
		goto L1
	} else {
		goto L58
	}
L44:
	;
	v166 = v12
	goto L46
L45:
	;
	v166 = v164
	goto L46
L46:
	;
	goto L48
L47:
	;
	v204 = v201
	goto L43
L48:
	;
	if v12 == v166 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v201 = int32(0)
	goto L47
L50:
	;
	v204 = int32(-1)
	goto L43
L51:
	;
	goto L52
L52:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v12))))
	if int32(252) < v179 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 + int32(1)
	goto L57
L54:
	;
	v181 = v179 - int32(97)
	if v181 < int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v184 = int32(1)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v181)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_ISO_8859_2_stem[0]))))
	if int32(base.Ui32(v188)>>(uint(v181&int32(7))%32))&v184 != 0 {
		v201 = v184
		goto L47
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	goto L49
L58:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v213 < v212 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v253 < int32(0) {
		goto L1
	} else {
		goto L74
	}
L60:
	;
	v215 = v212
	goto L62
L61:
	;
	v215 = v213
	goto L62
L62:
	;
	v222 = v212
	goto L64
L63:
	;
	v253 = v233
	goto L59
L64:
	;
	if v222 == v215 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v253 = int32(-1)
	goto L59
L67:
	;
	goto L68
L68:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v222))))
	if int32(252) < v228 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v245 = v222 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v245
	v222 = v245
	goto L64
L70:
	;
	v230 = v228 - int32(97)
	if v230 < int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v233 = int32(1)
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v230)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_ISO_8859_2_stem[0]))))
	if int32(base.Ui32(v237)>>(uint(v230&int32(7))%32))&v233 != 0 {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	goto L69
L74:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v258 = v256 + v253
	goto L2
L75:
	;
	return v743
L76:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v330
	v335 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_1), int32(44))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L37
	} else {
		goto L91
	}
L77:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275+v273))))
	if v277 != int32(108) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v282 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_2), int32(2))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L37
	} else {
		goto L79
	}
L79:
	;
	if v282 == int32(0) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v286 < v289 {
		goto L76
	} else {
		goto L81
	}
L81:
	;
	v292 = v286 - int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v292 <= v293 {
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v292))))
	if base.B2i32(v297&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v297)%32)&int32(106790108) == int32(0)) != 0 {
		goto L76
	} else {
		goto L83
	}
L83:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v312 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_3), int32(23))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L37
	} else {
		goto L84
	}
L84:
	;
	if v312 == int32(0) {
		goto L76
	} else {
		goto L85
	}
L85:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v316 + (v286 - v309)
	v320 = F_slice_del(m, l0)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L37
	} else {
		goto L86
	}
L86:
	;
	if v320 < int32(0) {
		v743 = v320
		goto L75
	} else {
		goto L87
	}
L87:
	;
	v324 = F_r_undouble_2(m, l0)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L37
	} else {
		goto L88
	}
L88:
	;
	if v324 < int32(0) {
		v743 = v324
		goto L75
	} else {
		goto L89
	}
L89:
	;
	goto L76
L90:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v386
	v390 = v386 - int32(1)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v390 <= v391 {
		goto L107
	} else {
		goto L108
	}
L91:
	;
	if v335 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v339
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	if v339 < v342 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v344 = F_slice_del(m, l0)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L37
	} else {
		goto L94
	}
L94:
	;
	if v344 < int32(0) {
		v743 = v344
		goto L75
	} else {
		goto L95
	}
L95:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v348
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v348 <= v350 {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352+v348-int32(1)))))
	switch v356 - int32(225) {
	case 0, 8:
		goto L97
	default:
		goto L90
	}
L97:
	;
	v361 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_4), int32(2))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L37
	} else {
		goto L98
	}
L98:
	;
	if v361 == int32(0) {
		goto L90
	} else {
		goto L99
	}
L99:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	if v365 < v368 {
		goto L90
	} else {
		goto L100
	}
L100:
	;
	switch v361 - int32(1) {
	case 0:
		goto L102
	case 1:
		goto L101
	default:
		goto L90
	}
L101:
	;
	v380 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_5))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L37
	} else {
		goto L105
	}
L102:
	;
	v374 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_6))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L37
	} else {
		goto L103
	}
L103:
	;
	if int32(0) <= v374 {
		goto L90
	} else {
		goto L104
	}
L104:
	;
	v743 = v374
	goto L75
L105:
	;
	if v380 < int32(0) {
		v743 = v380
		goto L75
	} else {
		goto L106
	}
L106:
	;
	goto L90
L107:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v425
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v425
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v425-int32(3) <= v428 {
		goto L119
	} else {
		goto L120
	}
L108:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393+v390))))
	switch v395 - int32(110) {
	case 0, 6:
		goto L109
	default:
		goto L107
	}
L109:
	;
	v400 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_7), int32(3))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L37
	} else {
		goto L110
	}
L110:
	;
	if v400 == int32(0) {
		goto L107
	} else {
		goto L111
	}
L111:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v404
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	if v404 < v407 {
		goto L107
	} else {
		goto L112
	}
L112:
	;
	switch v400 - int32(1) {
	case 0:
		goto L114
	case 1:
		goto L113
	default:
		goto L107
	}
L113:
	;
	v419 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_8))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L37
	} else {
		goto L117
	}
L114:
	;
	v413 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_9))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L37
	} else {
		goto L115
	}
L115:
	;
	if int32(0) <= v413 {
		goto L107
	} else {
		goto L116
	}
L116:
	;
	v743 = v413
	goto L75
L117:
	;
	if v419 < int32(0) {
		v743 = v419
		goto L75
	} else {
		goto L118
	}
L118:
	;
	goto L107
L119:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v470
	v472 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v470
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v470 <= v475 {
		v534 = v472
		goto L134
	} else {
		goto L135
	}
L120:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432+v425-int32(1)))))
	if v436 != int32(108) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v441 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_10), int32(6))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L37
	} else {
		goto L122
	}
L122:
	;
	if v441 == int32(0) {
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v445
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	if v445 < v448 {
		goto L119
	} else {
		goto L124
	}
L124:
	;
	switch v441 - int32(1) {
	case 0:
		goto L127
	case 1:
		goto L126
	case 2:
		goto L125
	default:
		goto L119
	}
L125:
	;
	v464 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_11))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L37
	} else {
		goto L132
	}
L126:
	;
	v458 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_12))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L37
	} else {
		goto L130
	}
L127:
	;
	v452 = F_slice_del(m, l0)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L37
	} else {
		goto L128
	}
L128:
	;
	if int32(0) <= v452 {
		goto L119
	} else {
		goto L129
	}
L129:
	;
	v743 = v452
	goto L75
L130:
	;
	if int32(0) <= v458 {
		goto L119
	} else {
		goto L131
	}
L131:
	;
	v743 = v458
	goto L75
L132:
	;
	if v464 < int32(0) {
		v743 = v464
		goto L75
	} else {
		goto L133
	}
L133:
	;
	goto L119
L134:
	;
	if v534 < int32(0) {
		v743 = v534
		goto L75
	} else {
		goto L150
	}
L135:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477+v470-int32(1)))))
	switch v481 - int32(225) {
	case 0, 8:
		goto L136
	default:
		v534 = v472
		goto L134
	}
L136:
	;
	v486 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_13), int32(2))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L37
	} else {
		goto L137
	}
L137:
	;
	if v486 == int32(0) {
		v534 = v472
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v490
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	if v490 < v493 {
		v534 = v472
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v496 = v490 - int32(1)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v496 <= v497 {
		v534 = v472
		goto L134
	} else {
		goto L140
	}
L140:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499+v496))))
	if base.B2i32(v501&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v501)%32)&int32(106790108) == int32(0)) != 0 {
		v534 = v472
		goto L134
	} else {
		goto L141
	}
L141:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v516 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_3), int32(23))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L37
	} else {
		goto L142
	}
L142:
	;
	if v516 == int32(0) {
		v534 = v472
		goto L134
	} else {
		goto L143
	}
L143:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v520 + (v490 - v513)
	v524 = F_slice_del(m, l0)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L37
	} else {
		goto L144
	}
L144:
	;
	if v524 < int32(0) {
		v534 = v524
		goto L134
	} else {
		goto L145
	}
L145:
	;
	v529 = F_r_undouble_2(m, l0)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L37
	} else {
		goto L146
	}
L146:
	;
	if int32(0) < v529 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v533 = int32(1)
	goto L149
L148:
	;
	v533 = v529
	goto L149
L149:
	;
	v534 = v533
	goto L134
L150:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v539
	v541 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v539
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v539 <= v544 {
		v586 = v541
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if v586 < int32(0) {
		v743 = v586
		goto L75
	} else {
		goto L167
	}
L152:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v539-int32(1)))))
	if v550|int32(128) != int32(233) {
		v586 = v541
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v557 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_14), int32(12))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L37
	} else {
		goto L154
	}
L154:
	;
	if v557 == int32(0) {
		v586 = v541
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v561
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	if v561 < v564 {
		v586 = v541
		goto L151
	} else {
		goto L156
	}
L156:
	;
	switch v557 - int32(1) {
	case 0:
		goto L160
	case 1:
		goto L159
	case 2:
		goto L158
	default:
		goto L157
	}
L157:
	;
	v586 = int32(1)
	goto L151
L158:
	;
	v580 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_15))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L37
	} else {
		goto L165
	}
L159:
	;
	v574 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_16))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L37
	} else {
		goto L163
	}
L160:
	;
	v568 = F_slice_del(m, l0)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L37
	} else {
		goto L161
	}
L161:
	;
	if int32(0) <= v568 {
		goto L157
	} else {
		goto L162
	}
L162:
	;
	v586 = v568
	goto L151
L163:
	;
	if int32(0) <= v574 {
		goto L157
	} else {
		goto L164
	}
L164:
	;
	v586 = v574
	goto L151
L165:
	;
	if v580 < int32(0) {
		v586 = v580
		goto L151
	} else {
		goto L166
	}
L166:
	;
	goto L157
L167:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v591
	v593 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v591
	v598 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_17), int32(31))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L37
	} else {
		goto L169
	}
L168:
	;
	if v627 < int32(0) {
		v743 = v627
		goto L75
	} else {
		goto L182
	}
L169:
	;
	if v598 == int32(0) {
		v627 = v593
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v602
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)))
	if v602 < v605 {
		v627 = v593
		goto L168
	} else {
		goto L171
	}
L171:
	;
	switch v598 - int32(1) {
	case 0:
		goto L175
	case 1:
		goto L174
	case 2:
		goto L173
	default:
		goto L172
	}
L172:
	;
	v627 = int32(1)
	goto L168
L173:
	;
	v621 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_18))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L37
	} else {
		goto L180
	}
L174:
	;
	v615 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_19))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L37
	} else {
		goto L178
	}
L175:
	;
	v609 = F_slice_del(m, l0)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L37
	} else {
		goto L176
	}
L176:
	;
	if int32(0) <= v609 {
		goto L172
	} else {
		goto L177
	}
L177:
	;
	v627 = v609
	goto L168
L178:
	;
	if int32(0) <= v615 {
		goto L172
	} else {
		goto L179
	}
L179:
	;
	v627 = v615
	goto L168
L180:
	;
	if v621 < int32(0) {
		v627 = v621
		goto L168
	} else {
		goto L181
	}
L181:
	;
	goto L172
L182:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v631
	v633 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v631
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v631 <= v636 {
		v685 = v633
		goto L183
	} else {
		goto L184
	}
L183:
	;
	if v685 < int32(0) {
		v743 = v685
		goto L75
	} else {
		goto L199
	}
L184:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v640 = int32(1)
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638+v631-v640))))
	if base.B2i32(v642&int32(224) != int32(96))|base.B2i32(v640<<(uint(v642)%32)&int32(_a_F_hungarian_ISO_8859_2_stem_20) == int32(0)) != 0 {
		v685 = v633
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v656 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_21), int32(42))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L37
	} else {
		goto L186
	}
L186:
	;
	if v656 == int32(0) {
		v685 = v633
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v660
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v662)))
	if v660 < v663 {
		v685 = v633
		goto L183
	} else {
		goto L188
	}
L188:
	;
	switch v656 - int32(1) {
	case 0:
		goto L192
	case 1:
		goto L191
	case 2:
		goto L190
	default:
		goto L189
	}
L189:
	;
	v685 = int32(1)
	goto L183
L190:
	;
	v679 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_22))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L37
	} else {
		goto L197
	}
L191:
	;
	v673 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_23))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L37
	} else {
		goto L195
	}
L192:
	;
	v667 = F_slice_del(m, l0)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L37
	} else {
		goto L193
	}
L193:
	;
	if int32(0) <= v667 {
		goto L189
	} else {
		goto L194
	}
L194:
	;
	v685 = v667
	goto L183
L195:
	;
	if int32(0) <= v673 {
		goto L189
	} else {
		goto L196
	}
L196:
	;
	v685 = v673
	goto L183
L197:
	;
	if v679 < int32(0) {
		v685 = v679
		goto L183
	} else {
		goto L198
	}
L198:
	;
	goto L189
L199:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v690
	v692 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v690
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v690 <= v695 {
		v735 = v692
		goto L200
	} else {
		goto L201
	}
L200:
	;
	if v735 < int32(0) {
		v743 = v735
		goto L75
	} else {
		goto L216
	}
L201:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697+v690-int32(1)))))
	if v701 != int32(107) {
		v735 = v692
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v706 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_24), int32(7))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L37
	} else {
		goto L203
	}
L203:
	;
	if v706 == int32(0) {
		v735 = v692
		goto L200
	} else {
		goto L204
	}
L204:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v710
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	if v710 < v713 {
		v735 = v692
		goto L200
	} else {
		goto L205
	}
L205:
	;
	switch v706 - int32(1) {
	case 0:
		goto L209
	case 1:
		goto L208
	case 2:
		goto L207
	default:
		goto L206
	}
L206:
	;
	v735 = int32(1)
	goto L200
L207:
	;
	v729 = F_slice_del(m, l0)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L37
	} else {
		goto L214
	}
L208:
	;
	v725 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_25))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L37
	} else {
		goto L212
	}
L209:
	;
	v719 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_26))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L37
	} else {
		goto L210
	}
L210:
	;
	if int32(0) <= v719 {
		goto L206
	} else {
		goto L211
	}
L211:
	;
	v735 = v719
	goto L200
L212:
	;
	if int32(0) <= v725 {
		goto L206
	} else {
		goto L213
	}
L213:
	;
	v735 = v725
	goto L200
L214:
	;
	if v729 < int32(0) {
		v735 = v729
		goto L200
	} else {
		goto L215
	}
L215:
	;
	goto L206
L216:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v740
	v743 = int32(1)
	goto L75
}
