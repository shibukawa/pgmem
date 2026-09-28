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
									F_errfinish(m, int32(_a_F_HandleConcurrentAbort_1), int32(499), int32(_a_F_HandleConcurrentAbort_2))
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
func F_handle_pm_pmsignal_signal(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[1]))
	v8 = int32(0)
	v11 = base.AtomicRmwOr32(m, v8, int32(_a_F_handle_pm_pmsignal_signal_0), v8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v12 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	v15 = int32(0)
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_handle_pm_pmsignal_signal_0), v15)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v19 == v15 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v22 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[2]))
	if v26 == v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[3]))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v56 = F_pgmem_kill(m, v22, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v30 + int32(16)
	goto L1
L10:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)) = uint8(v36)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[4]))
	v44 = F_write(m, v40, v30+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v44 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_pmsignal_signal[5]))
	if v48 == int32(27) {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	v4 = F_superuser_arg(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v12 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(l0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				if v12 == int32(0) {
					return int32(0)
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v19)+70)))
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = v21
						return v25 & int32(1)
					}
				}
			}
		} else {
			v25 = int32(1)
			return v25 & int32(1)
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
func F_has_largeobject_privilege_id_id(m *base.Module, l0 int32) int64 {
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
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v14 = F_convert_any_priv_string(m, v9, int32(_a_F_has_largeobject_privilege_id_id_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
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
				return int64(0)
			} else {
				if v24 != 0 {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_id_id[1])))
					if v28 != 0 {
						v38 = int64(1)
						return v38
					} else {
						v29 = F_pg_largeobject_aclcheck_snapshot(m, v7, v6, v14, v23)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v29 == int32(0)))
						}
					}
				} else {
					v35 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
					v38 = int64(0)
					return v38
				}
			}
		}
	}
}
func F_hashagg_spill_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 float64
	_ = v19
	var v23 float64
	_ = v23
	var v25 int32
	_ = v25
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v41 float64
	_ = v41
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	v16 = int32(32)
	v19 = float64(1024)
	v23 = *(*float64)(unsafe.Add(mBase, _c_F_hashagg_spill_init[0]))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_hashagg_spill_init[1]))
	v29 = base.F64_mul(base.F64_mul(v23, base.F64_convert_i32_s(v25)), v19)
	v30 = float64(4.294967295e+09)
	if base.F64_lt(v29, v30) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v33))
	v41 = base.F64_mul(base.F64_add(base.F64_mul(v35, float64(0.25)), float64(-8192)), float64(0.0001220703125))
	v47 = base.F64_add(base.F64_div(base.F64_mul(base.F64_mul(l3, float64(1.5)), l4), v35), float64(1))
	if base.F64_gt(v47, v41) != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v33 = v29
	goto L4
L3:
	;
	v33 = v30
	goto L4
L4:
	;
	goto L1
L5:
	;
	v49 = v41
	goto L7
L6:
	;
	v49 = v47
	goto L7
L7:
	;
	if base.F64_lt(v49, float64(4)) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v52 = float64(4)
	goto L10
L9:
	;
	v52 = v49
	goto L10
L10:
	;
	if base.F64_gt(v52, float64(1024)) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = v19
	goto L13
L12:
	;
	v55 = v52
	goto L13
L13:
	;
	v56 = base.I32_trunc_sat_f64_s(v55)
	if base.Ui32(int32(2)) <= base.Ui32(v56) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v64 = v16 - base.I32_clz(v56-int32(1))
	goto L16
L15:
	;
	v64 = int32(0)
	goto L16
L16:
	;
	if int32(31) < l2+v64 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v68 = v16 - l2
	goto L19
L18:
	;
	v68 = v64
	goto L19
L19:
	;
	v69 = int32(1) << (uint(v68) % 32)
	v70 = F_palloc0_mul(m, int32(4), v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v70
	v74 = F_palloc0_mul(m, int32(8), v69)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v74
	v78 = F_palloc0_mul(m, int32(24), v69)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v78
	v82 = base.B2i32(v68 == int32(31))
	if v82 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v85 = int32(1)
	if v69 <= v85 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v69
	v126 = int32(32)
	v128 = v126 - (l2 + v68)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128
	v130 = int32(0)
	if v128 < v126 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	v88 = v85
	goto L29
L28:
	;
	v88 = v69
	goto L29
L29:
	;
	v95 = int32(0)
	goto L30
L30:
	;
	v102 = F_LogicalTapeCreate(m, l1)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L32
	}
L31:
	;
	goto L26
L32:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v104+v95<<(uint(int32(2))%32)))) = v102
	v110 = v95 + int32(1)
	if v110 != v88 {
		v95 = v110
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v137 = (v69 - int32(1)) << (uint(v128) % 32)
	goto L36
L35:
	;
	v137 = v130
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	if v82 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v141 = int32(1)
	if v69 <= v141 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	return
L40:
	;
	v144 = v141
	goto L42
L41:
	;
	v144 = v69
	goto L42
L42:
	;
	v151 = v130
	goto L43
L43:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_initHyperLogLog(m, v158+v151*int32(24), int32(5))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L20
	} else {
		goto L45
	}
L44:
	;
	goto L39
L45:
	;
	v166 = v151 + int32(1)
	if v166 != v144 {
		v151 = v166
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
}
func F_hashbpcharextended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
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
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
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
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
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
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
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
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
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
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
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
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
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
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
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
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
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
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v723 int32
	_ = v723
	var v728 int64
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L133
	}
L4:
	;
	v16 = int32(1)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v20 = v18 & v16
	if v20 != 0 {
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
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L128
	}
L7:
	;
	v21 = v16
	goto L9
L8:
	;
	v21 = int32(4)
	goto L9
L9:
	;
	v22 = v21 + v11
	if v18 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v55 = v49
	goto L21
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v28 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v39 = int32(1)
	if v20 != 0 {
		v49 = int32(base.Ui32(v18)>>(uint(v39)%32)) - v39
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v31 = int32(16)
	goto L16
L15:
	;
	v31 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v38 = int32(4)
	goto L19
L18:
	;
	v38 = v31
	goto L19
L19:
	;
	v49 = v38
	goto L10
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	if v55 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v72 = F_pg_newlocale_from_collation(m, v15)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v70 = v49 & (v49 >> (uint(int32(31)) % 32))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v65 = v55 - int32(1)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v65))))
	if v67 == int32(32) {
		v55 = v65
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v70 = v55
	goto L23
L28:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v729 != v11 {
		goto L124
	} else {
		goto L125
	}
L29:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v74 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v83 = v70 - int32(1636608432)
	if v77 == int64(0) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	goto L32
L32:
	;
	v393 = int32(0)
	v395 = F_pg_strnxfrm(m, v393, v393, v22, v70, v72)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L76
	}
L33:
	;
	v728 = base.I64_extend_i32_u(v383)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v383^v375-base.I32_rotl(v383, int32(24)))
	goto L28
L34:
	;
	if v22&int32(3) != 0 {
		goto L50
	} else {
		goto L51
	}
L35:
	;
	v120 = v83
	v122 = v83
	v124 = v83
	goto L34
L36:
	;
	goto L37
L37:
	;
	v87 = v83 + base.I32_wrap_i64(v77)
	v88 = v87 + v83
	v92 = int32(4)
	v94 = base.I32_wrap_i64(int64(base.Ui64(v77)>>(uint(int64(32))%64))) ^ base.I32_rotl(v83, v92)
	v98 = v87 - v94 ^ base.I32_rotl(v94, int32(6))
	v102 = v88 - v98 ^ base.I32_rotl(v98, int32(8))
	v103 = v88 + v94
	v104 = v98 + v103
	v105 = v102 + v104
	v109 = v103 - v102 ^ base.I32_rotl(v102, int32(16))
	v113 = v104 - v109 ^ base.I32_rotl(v109, int32(19))
	v118 = v105 + v109
	v120 = v118
	v122 = v105 - v113 ^ base.I32_rotl(v113, v92)
	v124 = v113 + v118
	goto L34
L38:
	;
	v361 = int32(14)
	v363 = v357 ^ v358 - base.I32_rotl(v357, v361)
	v367 = v363 ^ v356 - base.I32_rotl(v363, int32(11))
	v371 = v367 ^ v357 - base.I32_rotl(v367, int32(25))
	v375 = v371 ^ v363 - base.I32_rotl(v371, int32(16))
	v379 = v375 ^ v367 - base.I32_rotl(v375, int32(4))
	v383 = v379 ^ v371 - base.I32_rotl(v379, v361)
	goto L33
L39:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v356 = v348 + v351
	v357 = v349
	v358 = v350
	goto L38
L40:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	v348 = v344<<(uint(int32(8))%32) + v341
	v349 = v342
	v350 = v343
	goto L39
L41:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
	v341 = v337<<(uint(int32(16))%32) + v334
	v342 = v335
	v343 = v336
	goto L40
L42:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
	v334 = v330<<(uint(int32(24))%32) + v181
	v335 = v328
	v336 = v329
	goto L41
L43:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
	v328 = v324 + v326
	v329 = v325
	goto L42
L44:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
	v324 = v320<<(uint(int32(8))%32) + v318
	v325 = v319
	goto L43
L45:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
	v318 = v314<<(uint(int32(16))%32) + v312
	v319 = v313
	goto L44
L46:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
	v312 = v308<<(uint(int32(24))%32) + v182
	v313 = v307
	goto L45
L47:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
	v307 = v303<<(uint(int32(8))%32) + v302
	goto L46
L48:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
	v302 = v298<<(uint(int32(16))%32) + v297
	goto L47
L49:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+10)))
	v297 = v293<<(uint(int32(24))%32) + v183
	goto L48
L50:
	;
	if base.Ui32(int32(11)) < base.Ui32(v70) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v70) {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	v129 = v22
	v130 = v70
	v132 = v120
	v133 = v124
	v134 = v122
	goto L56
L54:
	;
	v178 = v22
	v179 = v70
	v181 = v120
	v182 = v124
	v183 = v122
	goto L55
L55:
	;
	switch v179 - int32(1) {
	case 0:
		v348 = v181
		v349 = v182
		v350 = v183
		goto L39
	case 1:
		v341 = v181
		v342 = v182
		v343 = v183
		goto L40
	case 2:
		v334 = v181
		v335 = v182
		v336 = v183
		goto L41
	case 3:
		v328 = v182
		v329 = v183
		goto L42
	case 4:
		v324 = v182
		v325 = v183
		goto L43
	case 5:
		v318 = v182
		v319 = v183
		goto L44
	case 6:
		v312 = v182
		v313 = v183
		goto L45
	case 7:
		v307 = v183
		goto L46
	case 8:
		v302 = v183
		goto L47
	case 9:
		v297 = v183
		goto L48
	case 10:
		goto L49
	default:
		v356 = v181
		v357 = v182
		v358 = v183
		goto L38
	}
L56:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v137 = v136 + v133
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	v141 = v140 + v134
	v143 = int32(4)
	v145 = v138 + v132 - v141 ^ base.I32_rotl(v141, v143)
	v149 = v137 - v145 ^ base.I32_rotl(v145, int32(6))
	v150 = v141 + v137
	v151 = v145 + v150
	v152 = v149 + v151
	v156 = v150 - v149 ^ base.I32_rotl(v149, int32(8))
	v160 = v151 - v156 ^ base.I32_rotl(v156, int32(16))
	v164 = v152 - v160 ^ base.I32_rotl(v160, int32(19))
	v165 = v156 + v152
	v166 = v160 + v165
	v167 = v164 + v166
	v171 = v165 - v164 ^ base.I32_rotl(v164, v143)
	v172 = int32(12)
	v173 = v129 + v172
	v175 = v130 - v172
	if base.Ui32(int32(11)) < base.Ui32(v175) {
		v129 = v173
		v130 = v175
		v132 = v166
		v133 = v167
		v134 = v171
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v178 = v173
	v179 = v175
	v181 = v166
	v182 = v167
	v183 = v171
	goto L55
L58:
	;
	goto L57
L59:
	;
	v189 = v22
	v190 = v70
	v192 = v120
	v193 = v124
	v194 = v122
	goto L62
L60:
	;
	v238 = v22
	v239 = v70
	v241 = v120
	v242 = v124
	v243 = v122
	goto L61
L61:
	;
	switch v239 - int32(1) {
	case 0:
		v290 = v241
		goto L65
	case 1:
		v285 = v241
		goto L66
	case 2:
		goto L67
	case 3:
		v278 = v242
		goto L68
	case 4:
		v275 = v242
		goto L69
	case 5:
		v270 = v242
		goto L70
	case 6:
		goto L71
	case 7:
		v261 = v243
		goto L72
	case 8:
		v256 = v243
		goto L73
	case 9:
		v251 = v243
		goto L74
	case 10:
		goto L75
	default:
		v356 = v241
		v357 = v242
		v358 = v243
		goto L38
	}
L62:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v197 = v196 + v193
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v201 = v200 + v194
	v203 = int32(4)
	v205 = v198 + v192 - v201 ^ base.I32_rotl(v201, v203)
	v209 = v197 - v205 ^ base.I32_rotl(v205, int32(6))
	v210 = v201 + v197
	v211 = v205 + v210
	v212 = v209 + v211
	v216 = v210 - v209 ^ base.I32_rotl(v209, int32(8))
	v220 = v211 - v216 ^ base.I32_rotl(v216, int32(16))
	v224 = v212 - v220 ^ base.I32_rotl(v220, int32(19))
	v225 = v216 + v212
	v226 = v220 + v225
	v227 = v224 + v226
	v231 = v225 - v224 ^ base.I32_rotl(v224, v203)
	v232 = int32(12)
	v233 = v189 + v232
	v235 = v190 - v232
	if base.Ui32(int32(11)) < base.Ui32(v235) {
		v189 = v233
		v190 = v235
		v192 = v226
		v193 = v227
		v194 = v231
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v238 = v233
	v239 = v235
	v241 = v226
	v242 = v227
	v243 = v231
	goto L61
L64:
	;
	goto L63
L65:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v356 = v290 + v291
	v357 = v242
	v358 = v243
	goto L38
L66:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	v290 = v286<<(uint(int32(8))%32) + v285
	goto L65
L67:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+2)))
	v285 = v281<<(uint(int32(16))%32) + v241
	goto L66
L68:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v356 = v279 + v241
	v357 = v278
	v358 = v243
	goto L38
L69:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+4)))
	v278 = v275 + v276
	goto L68
L70:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+5)))
	v275 = v271<<(uint(int32(8))%32) + v270
	goto L69
L71:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+6)))
	v270 = v266<<(uint(int32(16))%32) + v242
	goto L70
L72:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v356 = v262 + v241
	v357 = v264 + v242
	v358 = v261
	goto L38
L73:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+8)))
	v261 = v257<<(uint(int32(8))%32) + v256
	goto L72
L74:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+9)))
	v256 = v252<<(uint(int32(16))%32) + v251
	goto L73
L75:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+10)))
	v251 = v247<<(uint(int32(24))%32) + v243
	goto L74
L76:
	;
	v398 = v395 + int32(1)
	v399 = F_palloc(m, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v401 = F_pg_strnxfrm(m, v399, v398, v22, v70, v72)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(v395) < base.Ui32(v401) {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v405 = v401 + int32(1)
	v406 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v412 = v405 - int32(1636608432)
	if v406 == int64(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	F_pfree(m, v399)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L123
	}
L81:
	;
	if v399&int32(3) != 0 {
		goto L97
	} else {
		goto L98
	}
L82:
	;
	v449 = v412
	v451 = v412
	v453 = v412
	goto L81
L83:
	;
	goto L84
L84:
	;
	v416 = v412 + base.I32_wrap_i64(v406)
	v417 = v416 + v412
	v421 = int32(4)
	v423 = base.I32_wrap_i64(int64(base.Ui64(v406)>>(uint(int64(32))%64))) ^ base.I32_rotl(v412, v421)
	v427 = v416 - v423 ^ base.I32_rotl(v423, int32(6))
	v431 = v417 - v427 ^ base.I32_rotl(v427, int32(8))
	v432 = v417 + v423
	v433 = v427 + v432
	v434 = v431 + v433
	v438 = v432 - v431 ^ base.I32_rotl(v431, int32(16))
	v442 = v433 - v438 ^ base.I32_rotl(v438, int32(19))
	v447 = v434 + v438
	v449 = v447
	v451 = v434 - v442 ^ base.I32_rotl(v442, v421)
	v453 = v442 + v447
	goto L81
L85:
	;
	v690 = int32(14)
	v692 = v686 ^ v687 - base.I32_rotl(v686, v690)
	v696 = v692 ^ v685 - base.I32_rotl(v692, int32(11))
	v700 = v696 ^ v686 - base.I32_rotl(v696, int32(25))
	v704 = v700 ^ v692 - base.I32_rotl(v700, int32(16))
	v708 = v704 ^ v696 - base.I32_rotl(v704, int32(4))
	v712 = v708 ^ v700 - base.I32_rotl(v708, v690)
	goto L80
L86:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
	v685 = v677 + v680
	v686 = v678
	v687 = v679
	goto L85
L87:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+1)))
	v677 = v673<<(uint(int32(8))%32) + v670
	v678 = v671
	v679 = v672
	goto L86
L88:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+2)))
	v670 = v666<<(uint(int32(16))%32) + v663
	v671 = v664
	v672 = v665
	goto L87
L89:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+3)))
	v663 = v659<<(uint(int32(24))%32) + v510
	v664 = v657
	v665 = v658
	goto L88
L90:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+4)))
	v657 = v653 + v655
	v658 = v654
	goto L89
L91:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+5)))
	v653 = v649<<(uint(int32(8))%32) + v647
	v654 = v648
	goto L90
L92:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+6)))
	v647 = v643<<(uint(int32(16))%32) + v641
	v648 = v642
	goto L91
L93:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+7)))
	v641 = v637<<(uint(int32(24))%32) + v511
	v642 = v636
	goto L92
L94:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+8)))
	v636 = v632<<(uint(int32(8))%32) + v631
	goto L93
L95:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+9)))
	v631 = v627<<(uint(int32(16))%32) + v626
	goto L94
L96:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+10)))
	v626 = v622<<(uint(int32(24))%32) + v512
	goto L95
L97:
	;
	if base.Ui32(int32(11)) < base.Ui32(v405) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v405) {
		goto L106
	} else {
		goto L107
	}
L100:
	;
	v458 = v399
	v459 = v405
	v461 = v449
	v462 = v453
	v463 = v451
	goto L103
L101:
	;
	v507 = v399
	v508 = v405
	v510 = v449
	v511 = v453
	v512 = v451
	goto L102
L102:
	;
	switch v508 - int32(1) {
	case 0:
		v677 = v510
		v678 = v511
		v679 = v512
		goto L86
	case 1:
		v670 = v510
		v671 = v511
		v672 = v512
		goto L87
	case 2:
		v663 = v510
		v664 = v511
		v665 = v512
		goto L88
	case 3:
		v657 = v511
		v658 = v512
		goto L89
	case 4:
		v653 = v511
		v654 = v512
		goto L90
	case 5:
		v647 = v511
		v648 = v512
		goto L91
	case 6:
		v641 = v511
		v642 = v512
		goto L92
	case 7:
		v636 = v512
		goto L93
	case 8:
		v631 = v512
		goto L94
	case 9:
		v626 = v512
		goto L95
	case 10:
		goto L96
	default:
		v685 = v510
		v686 = v511
		v687 = v512
		goto L85
	}
L103:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	v466 = v465 + v462
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v458)))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v458)+8))
	v470 = v469 + v463
	v472 = int32(4)
	v474 = v467 + v461 - v470 ^ base.I32_rotl(v470, v472)
	v478 = v466 - v474 ^ base.I32_rotl(v474, int32(6))
	v479 = v470 + v466
	v480 = v474 + v479
	v481 = v478 + v480
	v485 = v479 - v478 ^ base.I32_rotl(v478, int32(8))
	v489 = v480 - v485 ^ base.I32_rotl(v485, int32(16))
	v493 = v481 - v489 ^ base.I32_rotl(v489, int32(19))
	v494 = v485 + v481
	v495 = v489 + v494
	v496 = v493 + v495
	v500 = v494 - v493 ^ base.I32_rotl(v493, v472)
	v501 = int32(12)
	v502 = v458 + v501
	v504 = v459 - v501
	if base.Ui32(int32(11)) < base.Ui32(v504) {
		v458 = v502
		v459 = v504
		v461 = v495
		v462 = v496
		v463 = v500
		goto L103
	} else {
		goto L105
	}
L104:
	;
	v507 = v502
	v508 = v504
	v510 = v495
	v511 = v496
	v512 = v500
	goto L102
L105:
	;
	goto L104
L106:
	;
	v518 = v399
	v519 = v405
	v521 = v449
	v522 = v453
	v523 = v451
	goto L109
L107:
	;
	v567 = v399
	v568 = v405
	v570 = v449
	v571 = v453
	v572 = v451
	goto L108
L108:
	;
	switch v568 - int32(1) {
	case 0:
		v619 = v570
		goto L112
	case 1:
		v614 = v570
		goto L113
	case 2:
		goto L114
	case 3:
		v607 = v571
		goto L115
	case 4:
		v604 = v571
		goto L116
	case 5:
		v599 = v571
		goto L117
	case 6:
		goto L118
	case 7:
		v590 = v572
		goto L119
	case 8:
		v585 = v572
		goto L120
	case 9:
		v580 = v572
		goto L121
	case 10:
		goto L122
	default:
		v685 = v570
		v686 = v571
		v687 = v572
		goto L85
	}
L109:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	v526 = v525 + v522
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	v530 = v529 + v523
	v532 = int32(4)
	v534 = v527 + v521 - v530 ^ base.I32_rotl(v530, v532)
	v538 = v526 - v534 ^ base.I32_rotl(v534, int32(6))
	v539 = v530 + v526
	v540 = v534 + v539
	v541 = v538 + v540
	v545 = v539 - v538 ^ base.I32_rotl(v538, int32(8))
	v549 = v540 - v545 ^ base.I32_rotl(v545, int32(16))
	v553 = v541 - v549 ^ base.I32_rotl(v549, int32(19))
	v554 = v545 + v541
	v555 = v549 + v554
	v556 = v553 + v555
	v560 = v554 - v553 ^ base.I32_rotl(v553, v532)
	v561 = int32(12)
	v562 = v518 + v561
	v564 = v519 - v561
	if base.Ui32(int32(11)) < base.Ui32(v564) {
		v518 = v562
		v519 = v564
		v521 = v555
		v522 = v556
		v523 = v560
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v567 = v562
	v568 = v564
	v570 = v555
	v571 = v556
	v572 = v560
	goto L108
L111:
	;
	goto L110
L112:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	v685 = v619 + v620
	v686 = v571
	v687 = v572
	goto L85
L113:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+1)))
	v619 = v615<<(uint(int32(8))%32) + v614
	goto L112
L114:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)))
	v614 = v610<<(uint(int32(16))%32) + v570
	goto L113
L115:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v685 = v608 + v570
	v686 = v607
	v687 = v572
	goto L85
L116:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+4)))
	v607 = v604 + v605
	goto L115
L117:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+5)))
	v604 = v600<<(uint(int32(8))%32) + v599
	goto L116
L118:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+6)))
	v599 = v595<<(uint(int32(16))%32) + v571
	goto L117
L119:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	v685 = v591 + v570
	v686 = v593 + v571
	v687 = v590
	goto L85
L120:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+8)))
	v590 = v586<<(uint(int32(8))%32) + v585
	goto L119
L121:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+9)))
	v585 = v581<<(uint(int32(16))%32) + v580
	goto L120
L122:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+10)))
	v580 = v576<<(uint(int32(24))%32) + v572
	goto L121
L123:
	;
	v728 = base.I64_extend_i32_u(v712)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v712^v704-base.I32_rotl(v712, int32(24)))
	goto L28
L124:
	;
	F_pfree(m, v11)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	return v728
L127:
	;
	goto L126
L128:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_hashbpcharextended_0), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errhint(m, int32(_a_F_hashbpcharextended_1), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_hashbpcharextended_2), int32(1060), int32(_a_F_hashbpcharextended_3))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errmsg_internal(m, int32(_a_F_hashbpcharextended_4), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_hashbpcharextended_2), int32(1085), int32(_a_F_hashbpcharextended_3))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v26 float64
	_ = v26
	var v28 float64
	_ = v28
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v16 = v13 + int32(8)
	base.MemoryFill(m, v16, int32(0), int32(72))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = int32(1)
	F_genericcostestimate(m, l0, l1, l2, v16)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		v24 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = v24
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v26
		v28 = *(*float64)(unsafe.Add(mBase, uint32(v13)+24))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v28
		v30 = *(*float64)(unsafe.Add(mBase, uint32(v13)+32))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v30
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v32
		m.G0 = v13 + int32(80)
		return
	}
}
func F_hashfloat8extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
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
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
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
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
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
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
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
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v338 int64
	_ = v338
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = v9 & int64(9223372036854775807)
	if v13 != int64(0) {
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v13) {
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(9221120237041090560)
		} else {
		}
		v21 = v7 + int32(8)
		v28 = int32(-1636608424)
		if v11 == int64(0) {
			v65 = v28
			v67 = v28
			v69 = v28
		} else {
			v31 = base.I32_wrap_i64(v11)
			v33 = v31 + int32(1021750448)
			v37 = int32(4)
			v39 = base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(int64(32))%64))) ^ base.I32_rotl(v28, v37)
			v43 = v28 + v31 - v39 ^ base.I32_rotl(v39, int32(6))
			v47 = v33 - v43 ^ base.I32_rotl(v43, int32(8))
			v48 = v33 + v39
			v49 = v43 + v48
			v50 = v47 + v49
			v54 = v48 - v47 ^ base.I32_rotl(v47, int32(16))
			v58 = v49 - v54 ^ base.I32_rotl(v54, int32(19))
			v63 = v50 + v54
			v65 = v63
			v67 = v50 - v58 ^ base.I32_rotl(v58, v37)
			v69 = v58 + v63
		}
		if v21&int32(3) != 0 {
			switch int32(7) {
			case 0:
				v293 = v65
				v294 = v69
				v295 = v67
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 1:
				v286 = v65
				v287 = v69
				v288 = v67
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 2:
				v279 = v65
				v280 = v69
				v281 = v67
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 3:
				v273 = v69
				v274 = v67
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
				v279 = v275<<(uint(int32(24))%32) + v65
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 4:
				v269 = v69
				v270 = v67
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
				v279 = v275<<(uint(int32(24))%32) + v65
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 5:
				v263 = v69
				v264 = v67
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
				v279 = v275<<(uint(int32(24))%32) + v65
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 6:
				v257 = v69
				v258 = v67
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
				v279 = v275<<(uint(int32(24))%32) + v65
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 7:
				v252 = v67
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)))
				v257 = v253<<(uint(int32(24))%32) + v69
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
				v279 = v275<<(uint(int32(24))%32) + v65
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 8:
				v247 = v67
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
				v252 = v248<<(uint(int32(8))%32) + v247
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)))
				v257 = v253<<(uint(int32(24))%32) + v69
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
				v279 = v275<<(uint(int32(24))%32) + v65
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 9:
				v242 = v67
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+9)))
				v247 = v243<<(uint(int32(16))%32) + v242
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
				v252 = v248<<(uint(int32(8))%32) + v247
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)))
				v257 = v253<<(uint(int32(24))%32) + v69
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
				v279 = v275<<(uint(int32(24))%32) + v65
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 10:
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+10)))
				v242 = v238<<(uint(int32(24))%32) + v67
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+9)))
				v247 = v243<<(uint(int32(16))%32) + v242
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
				v252 = v248<<(uint(int32(8))%32) + v247
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)))
				v257 = v253<<(uint(int32(24))%32) + v69
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
				v279 = v275<<(uint(int32(24))%32) + v65
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			default:
				v301 = v65
				v302 = v69
				v303 = v67
			}
		} else {
			switch int32(7) {
			case 0:
				v235 = v65
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v235 + v236
				v302 = v69
				v303 = v67
			case 1:
				v230 = v65
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v235 = v231<<(uint(int32(8))%32) + v230
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v235 + v236
				v302 = v69
				v303 = v67
			case 2:
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v230 = v226<<(uint(int32(16))%32) + v65
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v235 = v231<<(uint(int32(8))%32) + v230
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v301 = v235 + v236
				v302 = v69
				v303 = v67
			case 3:
				v223 = v69
				v224 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v301 = v224 + v65
				v302 = v223
				v303 = v67
			case 4:
				v220 = v69
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v223 = v220 + v221
				v224 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v301 = v224 + v65
				v302 = v223
				v303 = v67
			case 5:
				v215 = v69
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
				v220 = v216<<(uint(int32(8))%32) + v215
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v223 = v220 + v221
				v224 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v301 = v224 + v65
				v302 = v223
				v303 = v67
			case 6:
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
				v215 = v211<<(uint(int32(16))%32) + v69
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
				v220 = v216<<(uint(int32(8))%32) + v215
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
				v223 = v220 + v221
				v224 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v301 = v224 + v65
				v302 = v223
				v303 = v67
			case 7:
				v206 = v67
				v207 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v301 = v207 + v65
				v302 = v209 + v69
				v303 = v206
			case 8:
				v201 = v67
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
				v206 = v202<<(uint(int32(8))%32) + v201
				v207 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v301 = v207 + v65
				v302 = v209 + v69
				v303 = v206
			case 9:
				v196 = v67
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+9)))
				v201 = v197<<(uint(int32(16))%32) + v196
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
				v206 = v202<<(uint(int32(8))%32) + v201
				v207 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v301 = v207 + v65
				v302 = v209 + v69
				v303 = v206
			case 10:
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+10)))
				v196 = v192<<(uint(int32(24))%32) + v67
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+9)))
				v201 = v197<<(uint(int32(16))%32) + v196
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
				v206 = v202<<(uint(int32(8))%32) + v201
				v207 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v301 = v207 + v65
				v302 = v209 + v69
				v303 = v206
			default:
				v301 = v65
				v302 = v69
				v303 = v67
			}
		}
		v306 = int32(14)
		v308 = v302 ^ v303 - base.I32_rotl(v302, v306)
		v312 = v308 ^ v301 - base.I32_rotl(v308, int32(11))
		v316 = v312 ^ v302 - base.I32_rotl(v312, int32(25))
		v320 = v316 ^ v308 - base.I32_rotl(v316, int32(16))
		v324 = v320 ^ v312 - base.I32_rotl(v320, int32(4))
		v328 = v324 ^ v316 - base.I32_rotl(v324, v306)
		v338 = base.I64_extend_i32_u(v328)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v328^v320-base.I32_rotl(v328, int32(24)))
	} else {
		v338 = v11
	}
	m.G0 = v7 + int32(16)
	return v338
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
	v14 = v11 + int32(8)
	v16 = v11 + int32(7)
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
func F_hashint2(m *base.Module, l0 int32) int64 {
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
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v7 = int32(711645284)
	v10 = v2 - int32(1636608428) ^ v7 - int32(1455628627)
	v15 = v10 ^ int32(-1636608428) - base.I32_rotl(v10, int32(25))
	v20 = v15 ^ v7 - base.I32_rotl(v15, int32(16))
	v24 = v20 ^ v10 - base.I32_rotl(v20, int32(4))
	v28 = v24 ^ v15 - base.I32_rotl(v24, int32(14))
	return base.I64_extend_i32_u(v28 ^ v20 - base.I32_rotl(v28, int32(24)))
}
func F_hashmacaddrextended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
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
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
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
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
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
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
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
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = int32(-1636608426)
	if v4 == int64(0) {
		v47 = v10
		v49 = v10
		v51 = v10
	} else {
		v13 = base.I32_wrap_i64(v4)
		v15 = v13 + int32(1021750444)
		v19 = int32(4)
		v21 = base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(32))%64))) ^ base.I32_rotl(v10, v19)
		v25 = v10 + v13 - v21 ^ base.I32_rotl(v21, int32(6))
		v29 = v15 - v25 ^ base.I32_rotl(v25, int32(8))
		v30 = v15 + v21
		v31 = v25 + v30
		v32 = v29 + v31
		v36 = v30 - v29 ^ base.I32_rotl(v29, int32(16))
		v40 = v31 - v36 ^ base.I32_rotl(v36, int32(19))
		v45 = v32 + v36
		v47 = v45
		v49 = v32 - v40 ^ base.I32_rotl(v40, v19)
		v51 = v40 + v45
	}
	if v2&int32(3) != 0 {
		switch int32(5) {
		case 0:
			v275 = v47
			v276 = v51
			v277 = v49
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 1:
			v268 = v47
			v269 = v51
			v270 = v49
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 2:
			v261 = v47
			v262 = v51
			v263 = v49
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 3:
			v255 = v51
			v256 = v49
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 4:
			v251 = v51
			v252 = v49
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 5:
			v245 = v51
			v246 = v49
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 6:
			v239 = v51
			v240 = v49
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 7:
			v234 = v49
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v239 = v235<<(uint(int32(24))%32) + v51
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 8:
			v229 = v49
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v234 = v230<<(uint(int32(8))%32) + v229
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v239 = v235<<(uint(int32(24))%32) + v51
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 9:
			v224 = v49
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v229 = v225<<(uint(int32(16))%32) + v224
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v234 = v230<<(uint(int32(8))%32) + v229
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v239 = v235<<(uint(int32(24))%32) + v51
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 10:
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v224 = v220<<(uint(int32(24))%32) + v49
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v229 = v225<<(uint(int32(16))%32) + v224
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v234 = v230<<(uint(int32(8))%32) + v229
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v239 = v235<<(uint(int32(24))%32) + v51
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		default:
			v283 = v47
			v284 = v51
			v285 = v49
		}
	} else {
		switch int32(5) {
		case 0:
			v217 = v47
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v217 + v218
			v284 = v51
			v285 = v49
		case 1:
			v212 = v47
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v217 = v213<<(uint(int32(8))%32) + v212
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v217 + v218
			v284 = v51
			v285 = v49
		case 2:
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v212 = v208<<(uint(int32(16))%32) + v47
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v217 = v213<<(uint(int32(8))%32) + v212
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v217 + v218
			v284 = v51
			v285 = v49
		case 3:
			v205 = v51
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v283 = v206 + v47
			v284 = v205
			v285 = v49
		case 4:
			v202 = v51
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v205 = v202 + v203
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v283 = v206 + v47
			v284 = v205
			v285 = v49
		case 5:
			v197 = v51
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v205 = v202 + v203
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v283 = v206 + v47
			v284 = v205
			v285 = v49
		case 6:
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v197 = v193<<(uint(int32(16))%32) + v51
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v205 = v202 + v203
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v283 = v206 + v47
			v284 = v205
			v285 = v49
		case 7:
			v188 = v49
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v283 = v189 + v47
			v284 = v191 + v51
			v285 = v188
		case 8:
			v183 = v49
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v283 = v189 + v47
			v284 = v191 + v51
			v285 = v188
		case 9:
			v178 = v49
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v283 = v189 + v47
			v284 = v191 + v51
			v285 = v188
		case 10:
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v178 = v174<<(uint(int32(24))%32) + v49
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v283 = v189 + v47
			v284 = v191 + v51
			v285 = v188
		default:
			v283 = v47
			v284 = v51
			v285 = v49
		}
	}
	v288 = int32(14)
	v290 = v284 ^ v285 - base.I32_rotl(v284, v288)
	v294 = v290 ^ v283 - base.I32_rotl(v290, int32(11))
	v298 = v294 ^ v284 - base.I32_rotl(v294, int32(25))
	v302 = v298 ^ v290 - base.I32_rotl(v298, int32(16))
	v306 = v302 ^ v294 - base.I32_rotl(v302, int32(4))
	v310 = v306 ^ v298 - base.I32_rotl(v306, v288)
	return base.I64_extend_i32_u(v310)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v310^v302-base.I32_rotl(v310, int32(24)))
}
func F_hashname(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return base.I64_extend_i32_u(v263 ^ v255 - base.I32_rotl(v263, int32(24)))
}
func F_hashoidvector(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_check_valid_oidvector(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
		return base.I64_extend_i32_u(v271 ^ v263 - base.I32_rotl(v271, int32(24)))
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
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
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
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
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
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
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
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
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
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
	v68 = F_find_among_b(m, l0, int32(_a_F_hindi_UTF_8_stem_0), int32(132), int32(_a_F_hindi_UTF_8_stem_1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v68 == int32(0) {
		v83 = v2
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v74
	v76 = F_slice_del(m, l0)
	mBase = m.M
	if v76 < int32(0) {
		v83 = v76
		goto L1
	} else {
		goto L26
	}
L26:
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v14 == int32(0) {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0]))
		v23 = F_AllocSetContextCreateInternal(m, v18, int32(_a_F_hnswinsert_0), int32(0), int32(_a_F_hnswinsert_1), int32(_a_F_hnswinsert_2))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = int32(_a_F_hnswinsert_3)
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0])) = v23
			v31 = F_HnswGetTypeInfo(m, l0)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v34 = v12 + int32(12)
				F_HnswInitSupport(m, v34, l0)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v39 = F_HnswFormIndexValue(m, v12+int32(24), l1, l2, v31, v34)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 != 0 {
							v41 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
							v43 = F_HnswInsertTupleOnDisk(m, l0, v34, v41, l3, int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0])) = v28
								F_MemoryContextDelete(m, v23)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									m.G0 = v12 + int32(32)
									return int32(0)
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_hnswinsert[0])) = v28
							F_MemoryContextDelete(m, v23)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								m.G0 = v12 + int32(32)
								return int32(0)
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v12 + int32(32)
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
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
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 < v9 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v178
	v182 = v178 - int32(1)
	if v182 <= v9 {
		goto L52
	} else {
		goto L53
	}
L2:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v173 + v172
	goto L1
L3:
	;
	if v62 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v21 = v9
	goto L6
L5:
	;
	v21 = v7
	goto L6
L6:
	;
	goto L8
L7:
	;
	v62 = v57
	goto L3
L8:
	;
	if v9 == v21 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v57 = int32(0)
	goto L7
L10:
	;
	v62 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v33 = int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v9))))
	if int32(252) < v36 {
		v57 = v33
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v38 = v36 - int32(97)
	if v38 < int32(0) {
		v57 = v33
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v38)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_ISO_8859_2_stem[0]))))
	if int32(base.Ui32(v44)>>(uint(v38&int32(7))%32))&int32(1) == int32(0) {
		v57 = v33
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9 + int32(1)
	goto L16
L16:
	;
	goto L9
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v74 < v73 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v129 < v9 {
		goto L36
	} else {
		goto L37
	}
L20:
	;
	if int32(0) <= v117 {
		v172 = v117
		goto L2
	} else {
		goto L34
	}
L21:
	;
	v76 = v73
	goto L23
L22:
	;
	v76 = v74
	goto L23
L23:
	;
	v82 = v73
	goto L25
L24:
	;
	v117 = int32(1)
	goto L20
L25:
	;
	if v82 == v76 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v117 = int32(-1)
	goto L20
L28:
	;
	goto L29
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v82))))
	if int32(252) < v91 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v93 = v91 - int32(97)
	if v93 < int32(0) {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v93)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_ISO_8859_2_stem[0]))))
	if int32(base.Ui32(v99)>>(uint(v93&int32(7))%32))&int32(1) == int32(0) {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v108 = v82 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
	v82 = v108
	goto L25
L34:
	;
	goto L1
L35:
	;
	if v169 < int32(0) {
		goto L1
	} else {
		goto L50
	}
L36:
	;
	v131 = v9
	goto L38
L37:
	;
	v131 = v129
	goto L38
L38:
	;
	v138 = v9
	goto L40
L39:
	;
	v169 = v149
	goto L35
L40:
	;
	if v138 == v131 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v169 = int32(-1)
	goto L35
L43:
	;
	goto L44
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142+v138))))
	if int32(252) < v144 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v161 = v138 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161
	v138 = v161
	goto L40
L46:
	;
	v146 = v144 - int32(97)
	if v146 < int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v149 = int32(1)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v146)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_ISO_8859_2_stem[0]))))
	if int32(base.Ui32(v153)>>(uint(v146&int32(7))%32))&v149 != 0 {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L45
L50:
	;
	v172 = v169
	goto L2
L51:
	;
	return v684
L52:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v264
	v270 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_0), int32(44), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L55
	} else {
		goto L73
	}
L53:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v182))))
	if v186 != int32(108) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v192 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_1), int32(2), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	return int32(0)
L56:
	;
	if v192 == int32(0) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v198 < v200 {
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v203 = v198 - int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v203 <= v204 {
		goto L52
	} else {
		goto L59
	}
L59:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+v203))))
	if base.B2i32(v208&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v208)%32)&int32(106790108) == int32(0)) != 0 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v224 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_2), int32(23), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L55
	} else {
		goto L61
	}
L61:
	;
	if v224 == int32(0) {
		goto L52
	} else {
		goto L62
	}
L62:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v228 + (v198 - v220)
	v232 = F_slice_del(m, l0)
	mBase = m.M
	if v232 < int32(0) {
		v684 = v232
		goto L51
	} else {
		goto L63
	}
L63:
	;
	v235 = int32(0)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v239 <= v240 {
		v259 = v235
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v259 < int32(0) {
		v684 = v259
		goto L51
	} else {
		goto L71
	}
L65:
	;
	goto L64
L66:
	;
	v243 = v239 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v243
	if v243 <= v240 {
		v259 = v235
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v248 = v239 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v248
	v252 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v252 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v255 = int32(1)
	goto L70
L69:
	;
	v255 = v252
	goto L70
L70:
	;
	v259 = v255
	goto L65
L71:
	;
	goto L52
L72:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v319
	v323 = v319 - int32(1)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v323 <= v324 {
		goto L88
	} else {
		goto L89
	}
L73:
	;
	if v270 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v274 < v276 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v278 = F_slice_del(m, l0)
	mBase = m.M
	if v278 < int32(0) {
		v684 = v278
		goto L51
	} else {
		goto L76
	}
L76:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v281 <= v283 {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+v281-int32(1)))))
	switch v289 - int32(225) {
	case 0, 8:
		goto L78
	default:
		goto L72
	}
L78:
	;
	v295 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_3), int32(2), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L55
	} else {
		goto L79
	}
L79:
	;
	if v295 == int32(0) {
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v299 < v301 {
		goto L72
	} else {
		goto L81
	}
L81:
	;
	switch v295 - int32(1) {
	case 0:
		goto L83
	case 1:
		goto L82
	default:
		goto L72
	}
L82:
	;
	v313 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_4))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L55
	} else {
		goto L86
	}
L83:
	;
	v307 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_5))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L55
	} else {
		goto L84
	}
L84:
	;
	if int32(0) <= v307 {
		goto L72
	} else {
		goto L85
	}
L85:
	;
	v684 = v307
	goto L51
L86:
	;
	if v313 < int32(0) {
		v684 = v313
		goto L51
	} else {
		goto L87
	}
L87:
	;
	goto L72
L88:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v358
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v358-int32(3) <= v361 {
		goto L100
	} else {
		goto L101
	}
L89:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v323))))
	switch v328 - int32(110) {
	case 0, 6:
		goto L90
	default:
		goto L88
	}
L90:
	;
	v334 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_6), int32(3), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L55
	} else {
		goto L91
	}
L91:
	;
	if v334 == int32(0) {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v338
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v338 < v340 {
		goto L88
	} else {
		goto L93
	}
L93:
	;
	switch v334 - int32(1) {
	case 0:
		goto L95
	case 1:
		goto L94
	default:
		goto L88
	}
L94:
	;
	v352 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_7))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L55
	} else {
		goto L98
	}
L95:
	;
	v346 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_8))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L55
	} else {
		goto L96
	}
L96:
	;
	if int32(0) <= v346 {
		goto L88
	} else {
		goto L97
	}
L97:
	;
	v684 = v346
	goto L51
L98:
	;
	if v352 < int32(0) {
		v684 = v352
		goto L51
	} else {
		goto L99
	}
L99:
	;
	goto L88
L100:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v402
	v404 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v402
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v402 <= v407 {
		v478 = v404
		goto L114
	} else {
		goto L115
	}
L101:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365+v358-int32(1)))))
	if v369 != int32(108) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v375 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_9), int32(6), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L55
	} else {
		goto L103
	}
L103:
	;
	if v375 == int32(0) {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v379
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v379 < v381 {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	switch v375 - int32(1) {
	case 0:
		goto L108
	case 1:
		goto L107
	case 2:
		goto L106
	default:
		goto L100
	}
L106:
	;
	v396 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_10))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L55
	} else {
		goto L112
	}
L107:
	;
	v390 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_11))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L55
	} else {
		goto L110
	}
L108:
	;
	v385 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v385 {
		goto L100
	} else {
		goto L109
	}
L109:
	;
	v684 = v385
	goto L51
L110:
	;
	if int32(0) <= v390 {
		goto L100
	} else {
		goto L111
	}
L111:
	;
	v684 = v390
	goto L51
L112:
	;
	if v396 < int32(0) {
		v684 = v396
		goto L51
	} else {
		goto L113
	}
L113:
	;
	goto L100
L114:
	;
	if v478 < int32(0) {
		v684 = v478
		goto L51
	} else {
		goto L130
	}
L115:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409+v402-int32(1)))))
	switch v413 - int32(225) {
	case 0, 8:
		goto L116
	default:
		v478 = v404
		goto L114
	}
L116:
	;
	v417 = v402 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v417
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v402 <= v420 {
		v478 = v404
		goto L114
	} else {
		goto L117
	}
L117:
	;
	v423 = v402 - int32(2)
	if v423 <= v407 {
		v478 = v404
		goto L114
	} else {
		goto L118
	}
L118:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423+v409))))
	if base.B2i32(v426&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v426)%32)&int32(106790108) == int32(0)) != 0 {
		v478 = v404
		goto L114
	} else {
		goto L119
	}
L119:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v442 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_2), int32(23), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L55
	} else {
		goto L120
	}
L120:
	;
	if v442 == int32(0) {
		v478 = v404
		goto L114
	} else {
		goto L121
	}
L121:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v446 + (v417 - v438)
	v450 = F_slice_del(m, l0)
	mBase = m.M
	if v450 < int32(0) {
		v478 = v450
		goto L114
	} else {
		goto L122
	}
L122:
	;
	v453 = int32(0)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v457 <= v458 {
		v477 = v453
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v478 = v477
	goto L114
L124:
	;
	goto L123
L125:
	;
	v461 = v457 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v461
	if v461 <= v458 {
		v477 = v453
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v466 = v457 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v466
	v470 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v470 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v473 = int32(1)
	goto L129
L128:
	;
	v473 = v470
	goto L129
L129:
	;
	v477 = v473
	goto L124
L130:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v484
	v486 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v484
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v484 <= v489 {
		v530 = v486
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if v530 < int32(0) {
		v684 = v530
		goto L51
	} else {
		goto L146
	}
L132:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491+v484-int32(1)))))
	if v495|int32(128) != int32(233) {
		v530 = v486
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v503 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_12), int32(12), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L55
	} else {
		goto L134
	}
L134:
	;
	if v503 == int32(0) {
		v530 = v486
		goto L131
	} else {
		goto L135
	}
L135:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v507
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v507 < v509 {
		v530 = v486
		goto L131
	} else {
		goto L136
	}
L136:
	;
	switch v503 - int32(1) {
	case 0:
		goto L140
	case 1:
		goto L139
	case 2:
		goto L138
	default:
		goto L137
	}
L137:
	;
	v530 = int32(1)
	goto L131
L138:
	;
	v524 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_13))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L55
	} else {
		goto L144
	}
L139:
	;
	v518 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_14))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L55
	} else {
		goto L142
	}
L140:
	;
	v513 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v513 {
		goto L137
	} else {
		goto L141
	}
L141:
	;
	v530 = v513
	goto L131
L142:
	;
	if int32(0) <= v518 {
		goto L137
	} else {
		goto L143
	}
L143:
	;
	v530 = v518
	goto L131
L144:
	;
	if v524 < int32(0) {
		v530 = v524
		goto L131
	} else {
		goto L145
	}
L145:
	;
	goto L137
L146:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v535
	v537 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v535
	v543 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_15), int32(31), v537)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L55
	} else {
		goto L148
	}
L147:
	;
	if v570 < int32(0) {
		v684 = v570
		goto L51
	} else {
		goto L160
	}
L148:
	;
	if v543 == int32(0) {
		v570 = v537
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v547
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v547 < v549 {
		v570 = v537
		goto L147
	} else {
		goto L150
	}
L150:
	;
	switch v543 - int32(1) {
	case 0:
		goto L154
	case 1:
		goto L153
	case 2:
		goto L152
	default:
		goto L151
	}
L151:
	;
	v570 = int32(1)
	goto L147
L152:
	;
	v564 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_16))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L55
	} else {
		goto L158
	}
L153:
	;
	v558 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_17))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L55
	} else {
		goto L156
	}
L154:
	;
	v553 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v553 {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v570 = v553
	goto L147
L156:
	;
	if int32(0) <= v558 {
		goto L151
	} else {
		goto L157
	}
L157:
	;
	v570 = v558
	goto L147
L158:
	;
	if v564 < int32(0) {
		v570 = v564
		goto L147
	} else {
		goto L159
	}
L159:
	;
	goto L151
L160:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v574
	v576 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v574
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v574 <= v579 {
		v627 = v576
		goto L161
	} else {
		goto L162
	}
L161:
	;
	if v627 < int32(0) {
		v684 = v627
		goto L51
	} else {
		goto L176
	}
L162:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v583 = int32(1)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581+v574-v583))))
	if base.B2i32(v585&int32(224) != int32(96))|base.B2i32(v583<<(uint(v585)%32)&int32(_a_F_hungarian_ISO_8859_2_stem_18) == int32(0)) != 0 {
		v627 = v576
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v600 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_19), int32(42), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L55
	} else {
		goto L164
	}
L164:
	;
	if v600 == int32(0) {
		v627 = v576
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v604
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v604 < v606 {
		v627 = v576
		goto L161
	} else {
		goto L166
	}
L166:
	;
	switch v600 - int32(1) {
	case 0:
		goto L170
	case 1:
		goto L169
	case 2:
		goto L168
	default:
		goto L167
	}
L167:
	;
	v627 = int32(1)
	goto L161
L168:
	;
	v621 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_20))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L55
	} else {
		goto L174
	}
L169:
	;
	v615 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_21))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L55
	} else {
		goto L172
	}
L170:
	;
	v610 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v610 {
		goto L167
	} else {
		goto L171
	}
L171:
	;
	v627 = v610
	goto L161
L172:
	;
	if int32(0) <= v615 {
		goto L167
	} else {
		goto L173
	}
L173:
	;
	v627 = v615
	goto L161
L174:
	;
	if v621 < int32(0) {
		v627 = v621
		goto L161
	} else {
		goto L175
	}
L175:
	;
	goto L167
L176:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v632
	v634 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v632
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v632 <= v637 {
		v676 = v634
		goto L177
	} else {
		goto L178
	}
L177:
	;
	if v676 < int32(0) {
		v684 = v676
		goto L51
	} else {
		goto L192
	}
L178:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639+v632-int32(1)))))
	if v643 != int32(107) {
		v676 = v634
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v649 = F_find_among_b(m, l0, int32(_a_F_hungarian_ISO_8859_2_stem_22), int32(7), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L55
	} else {
		goto L180
	}
L180:
	;
	if v649 == int32(0) {
		v676 = v634
		goto L177
	} else {
		goto L181
	}
L181:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v653
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v653 < v655 {
		v676 = v634
		goto L177
	} else {
		goto L182
	}
L182:
	;
	switch v649 - int32(1) {
	case 0:
		goto L186
	case 1:
		goto L185
	case 2:
		goto L184
	default:
		goto L183
	}
L183:
	;
	v676 = int32(1)
	goto L177
L184:
	;
	v671 = F_slice_del(m, l0)
	mBase = m.M
	if v671 < int32(0) {
		v676 = v671
		goto L177
	} else {
		goto L191
	}
L185:
	;
	v667 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_23))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L55
	} else {
		goto L189
	}
L186:
	;
	v661 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_ISO_8859_2_stem_24))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L55
	} else {
		goto L187
	}
L187:
	;
	if int32(0) <= v661 {
		goto L183
	} else {
		goto L188
	}
L188:
	;
	v676 = v661
	goto L177
L189:
	;
	if int32(0) <= v667 {
		goto L183
	} else {
		goto L190
	}
L190:
	;
	v676 = v667
	goto L177
L191:
	;
	goto L183
L192:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v681
	v684 = int32(1)
	goto L51
}
