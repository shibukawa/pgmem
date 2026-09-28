package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CompareTSQ(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 != v7 {
		if v6 < v7 {
			v12 = int32(-1)
		} else {
			v12 = int32(1)
		}
		return v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = int32(2)
		v16 = int32(base.Ui32(v14) >> (uint(v15) % 32))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v19 = int32(base.Ui32(v17) >> (uint(v15) % 32))
		if v16 != v19 {
			if base.Ui32(v16) < base.Ui32(v19) {
				v24 = int32(-1)
			} else {
				v24 = int32(1)
			}
			return v24
		} else {
			if v6 == int32(0) {
				return int32(0)
			} else {
				v31 = l0 + int32(8)
				v35 = F_QT2QTN(m, v31, v31+v6*int32(12))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v40 = l1 + int32(8)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v45 = F_QT2QTN(m, v40, v40+v41*int32(12))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = F_QTNodeCompare(m, v35, v45)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_QTNFree(m, v35)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_QTNFree(m, v45)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									return v47
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ConditionVariableInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v2))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(-1)
	return
}
func F_ConditionalLockRelation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v9
	v21 = F_LockAcquireExtended(m, v5+int32(16), int32(8), v2, int32(1), v5+int32(12), v2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		switch v21 {
		case 0, 3:
			m.G0 = v5 + int32(32)
			return base.B2i32(v21 != int32(0))
		default:
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+53)) = uint8(v28)
				m.G0 = v5 + int32(32)
				return base.B2i32(v21 != int32(0))
			}
		}
	}
}
func F_ConfigurePostmasterWaitSet(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[0]))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_FreeWaitEventSet(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v7 = int32(_a_F_ConfigurePostmasterWaitSet_0)
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[0])) = v8
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[1]))
	v16 = F_CreateWaitEventSet(m, v8, v13+int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[0])) = v16
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[2]))
	F_AddWaitEventToSet(m, v16, int32(1), int32(-1), v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[1]))
	if int32(0) < v26 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v30 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	return
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[0]))
	v33 = int32(2)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[3]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v30<<(uint(v33)%32))))
	F_AddWaitEventToSet(m, v32, v33, v39, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v44 = v30 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[1]))
	if v44 < v46 {
		v30 = v44
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F___cos(m *base.Module, l0 float64, l1 float64) float64 {
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v22 float64
	_ = v22
	v6 = float64(1)
	v7 = base.F64_mul(l0, l0)
	v9 = base.F64_mul(v7, float64(0.5))
	v10 = base.F64_sub(v6, v9)
	v22 = base.F64_mul(v7, v7)
	return base.F64_add(v10, base.F64_add(base.F64_sub(base.F64_sub(v6, v10), v9), base.F64_sub(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v22, v22), base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(l0, l1))))
}
func F_collprovider_name(m *base.Module, l0 int32) int32 {
	var v11 int32
	_ = v11
	switch l0 - int32(98) {
	case 0:
		v11 = int32(_a_F_collprovider_name_0)
		return v11
	case 1:
		return int32(_a_F_collprovider_name_1)
	default:
		v11 = int32(_a_F_collprovider_name_2)
		return v11
	case 7:
		return int32(_a_F_collprovider_name_3)
	}
}
func F_colorTrgmInfoCmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v6 = int64(56)
	v8 = int64(65280)
	v10 = int64(40)
	v13 = int64(16711680)
	v15 = int64(24)
	v17 = int64(4278190080)
	v19 = int64(8)
	v40 = v5<<(uint(v6)%64) | v5&v8<<(uint(v10)%64) | (v5&v13<<(uint(v15)%64) | v5&v17<<(uint(v19)%64)) | (int64(base.Ui64(v5)>>(uint(v19)%64))&v17 | int64(base.Ui64(v5)>>(uint(v15)%64))&v13 | (int64(base.Ui64(v5)>>(uint(v10)%64))&v8 | int64(base.Ui64(v5)>>(uint(v6)%64))))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v76 = v41<<(uint(v6)%64) | v41&v8<<(uint(v10)%64) | (v41&v13<<(uint(v15)%64) | v41&v17<<(uint(v19)%64)) | (int64(base.Ui64(v41)>>(uint(v19)%64))&v17 | int64(base.Ui64(v41)>>(uint(v15)%64))&v13 | (int64(base.Ui64(v41)>>(uint(v10)%64))&v8 | int64(base.Ui64(v41)>>(uint(v6)%64))))
	if v40 == v76 {
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v80 = int32(16711935)
		v82 = int32(8)
		v84 = int32(24)
		v89 = base.I64_extend_i32_u(base.I32_rotr(v79&v80, v82) | base.I32_rotr(v79, v84)&v80)
		v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v100 = base.I64_extend_i32_u(base.I32_rotr(v90&v80, v82) | base.I32_rotr(v90, v84)&v80)
		if v89 == v100 {
			v112 = int32(0)
		} else {
			v103 = v100
			v104 = v89
			if base.Ui64(v104) < base.Ui64(v103) {
				v108 = int32(-1)
			} else {
				v108 = int32(1)
			}
			v112 = v108
		}
	} else {
		v103 = v76
		v104 = v40
		if base.Ui64(v104) < base.Ui64(v103) {
			v108 = int32(-1)
		} else {
			v108 = int32(1)
		}
		v112 = v108
	}
	return v112
}
func F_combo_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
		m.T0[v4].(func(*base.Module, int32))(m, v3)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			v7 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
			F_pfree(m, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v7 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
		F_pfree(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F_commit_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_commit_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1058)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v15
	v19 = int32(_a_F_commit_cb_wrapper_1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_commit_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_commit_cb_wrapper[0])) = v9 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+147)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+164)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	m.T0[v37].(func(*base.Module, int32, int32, int64))(m, v11, l1, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		*(*int32)(unsafe.Add(mBase, _c_F_commit_cb_wrapper[0])) = v41
		m.G0 = v9 + int32(32)
		return
	}
}
func F_commit_prepared_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_commit_prepared_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1058)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v15
	v19 = int32(_a_F_commit_prepared_cb_wrapper_1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_commit_prepared_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_commit_prepared_cb_wrapper[0])) = v9 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+147)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+164)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v37 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_commit_prepared_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_commit_prepared_cb_wrapper_3), v9)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_commit_prepared_cb_wrapper_4), int32(1111), int32(_a_F_commit_prepared_cb_wrapper_5))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
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
		m.T0[v37].(func(*base.Module, int32, int32, int64))(m, v11, l1, l2)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_commit_prepared_cb_wrapper[0])) = v60
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_committssyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_SlruSyncFileTag(m, int32(_a_F_committssyncfiletag_0), l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v23 int32
	_ = v23
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v7 != v8 {
		if v7 < v8 {
			v13 = int32(-1)
		} else {
			v13 = int32(1)
		}
		return v13
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
		if base.F64_eq(v15, v16) != 0 {
			return int32(0)
		} else {
			if base.F64_lt(v15, v16) != 0 {
				v23 = int32(-1)
			} else {
				v23 = int32(1)
			}
			return v23
		}
	}
}
func F_compare_mcvs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return v4 - v5
}
func F_compare_pathkeys(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	if l0 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = int32(0)
	goto L6
L4:
	;
	if v53 != 0 {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v48 = int32(0)
	if v34 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v16 = int32(0)
	if l0 == v16 {
		v26 = v16
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(3)
L8:
	;
	if l1 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 <= v12 {
		v26 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = v22 + v12<<(uint(int32(2))%32)
	goto L8
L11:
	;
	v32 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.B2i32(v26 == v32)|base.B2i32(v34 == v32) != 0 {
		goto L5
	} else {
		goto L16
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 < v27 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v29 = int32(0)
	v53 = base.B2i32(v26 == v29)
	v55 = v29
	goto L4
L15:
	;
	goto L14
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34+v12<<(uint(int32(2))%32))))
	if v42 == v44 {
		v12 = v12 + int32(1)
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L7
L18:
	;
	v52 = int32(2)
	goto L20
L19:
	;
	v52 = v48
	goto L20
L20:
	;
	v53 = base.B2i32(v26 == v48)
	v55 = v52
	goto L4
L21:
	;
	v57 = v55
	goto L23
L22:
	;
	v57 = int32(1)
	goto L23
L23:
	;
	return v57
}
func F_compare_rows(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
	v8 = int32(16)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)))
	v11 = v7<<(uint(v8)%32) | v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+6)))
	v17 = v13<<(uint(v8)%32) | v16
	if base.Ui32(v11) < base.Ui32(v17) {
		v28 = int32(-1)
	} else {
		if base.Ui32(v17) < base.Ui32(v11) {
			v28 = int32(1)
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+8)))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)))
			if base.Ui32(v22) < base.Ui32(v23) {
				v28 = int32(-1)
			} else {
				v28 = base.B2i32(base.Ui32(v23) < base.Ui32(v22))
			}
		}
	}
	return v28
}
func F_comparecost_2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return v3 - v4
}
func F_complete_direction(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v6 = int32(0)
	v8 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		switch v8 - int32(324) {
		case 0, 5:
			v38 = v6
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
			return
		case 1, 2, 3, 4:
			F_plpgsql_push_back_token(m, v8, l2, l3, l4)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v25 = int32(0)
				v28 = int32(1)
				v32 = F_read_sql_construct(m, int32(324), int32(329), v25, int32(_a_F_complete_direction_0), int32(2), v28, v28, v25, v25, l2, l3, l4)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
					v35 = v6
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v36)
					v38 = v35
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
					return
				}
			}
		default:
			if v8 != int32(282) {
				if v8 != 0 {
					F_plpgsql_push_back_token(m, v8, l2, l3, l4)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v25 = int32(0)
						v28 = int32(1)
						v32 = F_read_sql_construct(m, int32(324), int32(329), v25, int32(_a_F_complete_direction_0), int32(2), v28, v28, v25, v25, l2, l3, l4)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
							v35 = v6
							v36 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v36)
							v38 = v35
							*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
							return
						}
					}
				} else {
					F_plpgsql_yyerror(m, l3, int32(0), l4, int32(_a_F_complete_direction_1))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(2147483647)
				v35 = int32(1)
				v36 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v36)
				v38 = v35
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
				return
			}
		}
	}
}
func F_compress_flush(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = m.Env.Pgmem_deflate_finish(m, v7)
	mBase = m.M
	if v8 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-105)
L2:
	;
	goto L3
L3:
	;
	v14 = l1 + int32(8)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v23 = m.Env.Pgmem_zstream_read(m, v22, v14, v15)
	mBase = m.M
	if v23 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v32
L6:
	;
	return int32(-105)
L7:
	;
	goto L8
L8:
	;
	if v23 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v32 = F_pushf_write(m, l0, v14, v23)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if int32(0) <= v32 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
}
func F_compress_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		m.Env.Pgmem_zstream_free(m, v5)
		mBase = m.M
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		if v7 != 0 {
			F_ResourceOwnerForget(m, v7, base.I64_extend_i32_u(v4), int32(_a_F_compress_free_0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_pfree(m, v4)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					base.MemoryFill(m, l0, int32(0), int32(_a_F_compress_free_1))
					F_pfree(m, l0)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_pfree(m, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				base.MemoryFill(m, l0, int32(0), int32(_a_F_compress_free_1))
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		base.MemoryFill(m, l0, int32(0), int32(_a_F_compress_free_1))
		F_pfree(m, l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			return
		}
	}
}
func F_contains_required_value(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	F_check_stack_depth(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v6 == int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(1)
L4:
	;
	goto L5
L5:
	;
	v11 = l0
	goto L6
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v12 - int32(33) {
	case 0:
		goto L9
	default:
		goto L10
	case 5:
		goto L11
	}
L7:
	;
	return int32(1)
L8:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L16
	}
L9:
	;
	return int32(0)
L10:
	;
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+2)))
	v29 = F_contains_required_value(m, v11+v25<<(uint(int32(3))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+2)))
	v19 = F_contains_required_value(m, v11+v15<<(uint(int32(3))%32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v19 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	return int32(1)
L14:
	;
	if v29 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	v36 = v11 - int32(8)
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if v37 != int32(2) {
		v11 = v36
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L7
}
func F_contains_user_functions_checker(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(base.Ui32(int32(_a_F_contains_user_functions_checker_0)) < base.Ui32(l0))
}
func F_convert_tuples_by_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v3 = int32(0)
	v7 = F_build_attrmap_by_name_if_req(m, l0, l1, v3)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v13 = F_palloc(m, int32(28))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
				v19 = F_palloc_mul(m, int32(8), v11)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v19
					v23 = F_palloc_mul(m, int32(1), v11)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v23
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v29 = v27 + int32(1)
						v30 = F_palloc_mul(m, int32(8), v29)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v30
							v34 = F_palloc_mul(m, int32(1), v29)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v34
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
								*(*int64)(unsafe.Add(mBase, uint32(v37))) = int64(0)
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
								v41 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v41)
								v44 = v13
								return v44
							}
						}
					}
				}
			}
		} else {
			v44 = v3
			return v44
		}
	}
}
func F_cookDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
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
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l5 != 0 {
		v14 = int32(43)
	} else {
		v14 = int32(30)
	}
	v15 = F_transformExpr(m, l0, l1, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if l5 == int32(0) {
			if l2 != 0 {
				v29 = F_exprType(m, v15)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v34 = F_coerce_to_target_type(m, l0, v15, v29, l2, l3, int32(1), int32(2), int32(-1))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v69 = F_format_type_be(m, l2)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = F_format_type_be(m, v29)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v71
											*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
											F_errmsg(m, int32(_a_F_cookDefault_0), v10)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_errhint(m, int32(_a_F_cookDefault_1), int32(0))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_cookDefault_2), int32(3421), int32(_a_F_cookDefault_3))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
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
						} else {
							v38 = v34
							F_assign_expr_collations(m, l0, v38)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v38
							}
						}
					}
				}
			} else {
				v38 = v15
				F_assign_expr_collations(m, l0, v38)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(16)
					return v38
				}
			}
		} else {
			v21 = F_check_nested_generated_walker(m, v15, l0)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_contain_mutable_functions_after_planning(m, v15)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_cookDefault_4), int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_cookDefault_2), int32(3384), int32(_a_F_cookDefault_3))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
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
						if l5 != int32(118) {
							if l2 != 0 {
								v29 = F_exprType(m, v15)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									v34 = F_coerce_to_target_type(m, l0, v15, v29, l2, l3, int32(1), int32(2), int32(-1))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										if v34 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(67141764))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v69 = F_format_type_be(m, l2)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = F_format_type_be(m, v29)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v71
															*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
															*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
															F_errmsg(m, int32(_a_F_cookDefault_0), v10)
															mBase = m.M
															v78 = m.ExcPending
															if v78 != 0 {
																return int32(0)
															} else {
																F_errhint(m, int32(_a_F_cookDefault_1), int32(0))
																mBase = m.M
																v82 = m.ExcPending
																if v82 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_cookDefault_2), int32(3421), int32(_a_F_cookDefault_3))
																	mBase = m.M
																	v87 = m.ExcPending
																	if v87 != 0 {
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
										} else {
											v38 = v34
											F_assign_expr_collations(m, l0, v38)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v38
											}
										}
									}
								}
							} else {
								v38 = v15
								F_assign_expr_collations(m, l0, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return v38
								}
							}
						} else {
							v27 = F_check_virtual_generated_security_walker(m, v15, l0)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									v29 = F_exprType(m, v15)
									mBase = m.M
									v30 = m.ExcPending
									if v30 != 0 {
										return int32(0)
									} else {
										v34 = F_coerce_to_target_type(m, l0, v15, v29, l2, l3, int32(1), int32(2), int32(-1))
										mBase = m.M
										v35 = m.ExcPending
										if v35 != 0 {
											return int32(0)
										} else {
											if v34 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(67141764))
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														v69 = F_format_type_be(m, l2)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v71 = F_format_type_be(m, v29)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v71
																*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
																*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
																F_errmsg(m, int32(_a_F_cookDefault_0), v10)
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return int32(0)
																} else {
																	F_errhint(m, int32(_a_F_cookDefault_1), int32(0))
																	mBase = m.M
																	v82 = m.ExcPending
																	if v82 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_cookDefault_2), int32(3421), int32(_a_F_cookDefault_3))
																		mBase = m.M
																		v87 = m.ExcPending
																		if v87 != 0 {
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
											} else {
												v38 = v34
												F_assign_expr_collations(m, l0, v38)
												mBase = m.M
												v41 = m.ExcPending
												if v41 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(16)
													return v38
												}
											}
										}
									}
								} else {
									v38 = v15
									F_assign_expr_collations(m, l0, v38)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v38
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
func F_copy_lladdr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	v4 = l3
	v6 = l5
	if base.Ui32(v4) <= base.Ui32(int32(24)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)) = uint8(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v6)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l4
	v12 = int32(17)
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v12)
	v15 = l1 + int32(12)
	if base.Ui32(int32(512)) <= base.Ui32(v4) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	goto L3
L5:
	;
	if v4 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v22 = v15 + v4
	if (v15^l2)&int32(3) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	base.MemoryCopy(m, v15, l2, v4)
	goto L10
L9:
	;
	goto L10
L10:
	;
	goto L4
L11:
	;
	if base.Ui32(v154) < base.Ui32(v22) {
		goto L45
	} else {
		goto L46
	}
L12:
	;
	if v15&int32(3) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(v22) < base.Ui32(int32(4)) {
		goto L36
	} else {
		goto L37
	}
L15:
	;
	v58 = v22 & int32(-4)
	if base.Ui32(v22) < base.Ui32(int32(64)) {
		v108 = v52
		v109 = v53
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = l2
	v53 = v15
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v4 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v52 = l2
	v53 = v15
	goto L15
L20:
	;
	goto L21
L21:
	;
	v35 = l2
	v36 = v15
	goto L22
L22:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v40)
	v42 = int32(1)
	v43 = v35 + v42
	v45 = v36 + v42
	if v45&int32(3) == int32(0) {
		v52 = v43
		v53 = v45
		goto L15
	} else {
		goto L24
	}
L23:
	;
	v52 = v43
	v53 = v45
	goto L15
L24:
	;
	if base.Ui32(v45) < base.Ui32(v22) {
		v35 = v43
		v36 = v45
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if base.Ui32(v58) <= base.Ui32(v109) {
		v153 = v108
		v154 = v109
		goto L11
	} else {
		goto L32
	}
L27:
	;
	v62 = v58 + int32(-64)
	if base.Ui32(v62) < base.Ui32(v53) {
		v108 = v52
		v109 = v53
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v65 = v52
	v66 = v53
	goto L29
L29:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+16)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v65)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+48)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v65)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+52)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v65)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+56)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v65)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+60)) = v100
	v102 = int32(-64)
	v103 = v65 - v102
	v105 = v66 - v102
	if base.Ui32(v105) <= base.Ui32(v62) {
		v65 = v103
		v66 = v105
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v108 = v103
	v109 = v105
	goto L26
L31:
	;
	goto L30
L32:
	;
	v115 = v108
	v116 = v109
	goto L33
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v120
	v122 = int32(4)
	v123 = v115 + v122
	v125 = v116 + v122
	if base.Ui32(v125) < base.Ui32(v58) {
		v115 = v123
		v116 = v125
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v153 = v123
	v154 = v125
	goto L11
L35:
	;
	goto L34
L36:
	;
	v153 = l2
	v154 = v15
	goto L11
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(v4) < base.Ui32(int32(4)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v153 = l2
	v154 = v15
	goto L11
L40:
	;
	goto L41
L41:
	;
	v134 = l2
	v135 = v15
	goto L42
L42:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v139)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)) = uint8(v141)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)) = uint8(v143)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+3)) = uint8(v145)
	v147 = int32(4)
	v148 = v134 + v147
	v150 = v135 + v147
	if base.Ui32(v150) <= base.Ui32(v22-int32(4)) {
		v134 = v148
		v135 = v150
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v153 = v148
	v154 = v150
	goto L11
L44:
	;
	goto L43
L45:
	;
	v160 = v153
	v161 = v154
	goto L48
L46:
	;
	goto L47
L47:
	;
	goto L4
L48:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	*(*uint8)(unsafe.Add(mBase, uint32(v161))) = uint8(v165)
	v167 = int32(1)
	v170 = v161 + v167
	if v170 != v22 {
		v160 = v160 + v167
		v161 = v170
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	goto L49
}
func F_cos(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v40 float64
	_ = v40
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v85 float64
	_ = v85
	var v105 float64
	_ = v105
	var v121 float64
	_ = v121
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v147 float64
	_ = v147
	var v148 float64
	_ = v148
	var v160 float64
	_ = v160
	var v181 float64
	_ = v181
	var v197 float64
	_ = v197
	var v219 float64
	_ = v219
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v14) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v14) < base.Ui32(int32(1044816030)) {
			v219 = float64(1)
		} else {
			v24 = float64(1)
			v25 = base.F64_mul(l0, l0)
			v27 = base.F64_mul(v25, float64(0.5))
			v28 = base.F64_sub(v24, v27)
			v40 = base.F64_mul(v25, v25)
			v219 = base.F64_add(v28, base.F64_add(base.F64_sub(base.F64_sub(v24, v28), v27), base.F64_sub(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v40, v40), base.F64_add(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(l0, float64(0)))))
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v14) {
			v219 = base.F64_sub(l0, l0)
		} else {
			v59 = F___rem_pio2(m, l0, v7)
			mBase = m.M
			v60 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v61 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			switch v59&int32(3) - int32(1) {
			case 0:
				v105 = base.F64_mul(v61, v61)
				v121 = base.F64_mul(v61, v105)
				v219 = base.F64_neg(base.F64_sub(v61, base.F64_add(base.F64_sub(base.F64_mul(v105, base.F64_sub(base.F64_mul(v60, float64(0.5)), base.F64_mul(v121, base.F64_add(base.F64_mul(base.F64_mul(v105, base.F64_mul(v105, v105)), base.F64_add(base.F64_mul(v105, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v105, base.F64_add(base.F64_mul(v105, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v60), base.F64_mul(v121, float64(0.16666666666666632)))))
			case 1:
				v144 = float64(1)
				v145 = base.F64_mul(v61, v61)
				v147 = base.F64_mul(v145, float64(0.5))
				v148 = base.F64_sub(v144, v147)
				v160 = base.F64_mul(v145, v145)
				v219 = base.F64_neg(base.F64_add(v148, base.F64_add(base.F64_sub(base.F64_sub(v144, v148), v147), base.F64_sub(base.F64_mul(v145, base.F64_add(base.F64_mul(v145, base.F64_add(base.F64_mul(v145, base.F64_add(base.F64_mul(v145, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v160, v160), base.F64_add(base.F64_mul(v145, base.F64_add(base.F64_mul(v145, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v61, v60)))))
			case 2:
				v181 = base.F64_mul(v61, v61)
				v197 = base.F64_mul(v61, v181)
				v219 = base.F64_sub(v61, base.F64_add(base.F64_sub(base.F64_mul(v181, base.F64_sub(base.F64_mul(v60, float64(0.5)), base.F64_mul(v197, base.F64_add(base.F64_mul(base.F64_mul(v181, base.F64_mul(v181, v181)), base.F64_add(base.F64_mul(v181, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v181, base.F64_add(base.F64_mul(v181, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v60), base.F64_mul(v197, float64(0.16666666666666632))))
			default:
				v69 = float64(1)
				v70 = base.F64_mul(v61, v61)
				v72 = base.F64_mul(v70, float64(0.5))
				v73 = base.F64_sub(v69, v72)
				v85 = base.F64_mul(v70, v70)
				v219 = base.F64_add(v73, base.F64_add(base.F64_sub(base.F64_sub(v69, v73), v72), base.F64_sub(base.F64_mul(v70, base.F64_add(base.F64_mul(v70, base.F64_add(base.F64_mul(v70, base.F64_add(base.F64_mul(v70, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v85, v85), base.F64_add(base.F64_mul(v70, base.F64_add(base.F64_mul(v70, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v61, v60))))
			}
		}
	}
	m.G0 = v7 + int32(16)
	return v219
}
func F_cosine_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 float32
	_ = v10
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 float32
	_ = v58
	var v59 float32
	_ = v59
	var v60 float32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 float32
	_ = v69
	var v71 float32
	_ = v71
	var v74 float32
	_ = v74
	var v76 float32
	_ = v76
	var v79 float32
	_ = v79
	var v83 float32
	_ = v83
	var v87 float32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v108 float32
	_ = v108
	var v109 float32
	_ = v109
	var v110 float32
	_ = v110
	var v115 int32
	_ = v115
	var v117 float32
	_ = v117
	var v119 float32
	_ = v119
	var v137 float32
	_ = v137
	var v138 float32
	_ = v138
	var v139 float32
	_ = v139
	var v143 float64
	_ = v143
	var v149 float64
	_ = v149
	var v174 float64
	_ = v174
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	v10 = float32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int64(0)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v28 = F_pg_detoast_datum(m, v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
			if v30 == v31 {
				v34 = base.I32_extend16_s(v30)
				if v34 <= int32(0) {
					v174 = math.Float64frombits(uint64(0x7ff8000000000000))
				} else {
					v37 = int32(8)
					v38 = v28 + v37
					v40 = v23 + v37
					if v34 == int32(1) {
						v97 = int32(0)
						v108 = v10
						v109 = v10
						v110 = v10
						v115 = v97 << (uint(int32(2)) % 32)
						v117 = *(*float32)(unsafe.Add(mBase, uint32(v40+v115)))
						v119 = *(*float32)(unsafe.Add(mBase, uint32(v115+v38)))
						v137 = base.F32_add(base.F32_mul(v117, v119), v108)
						v138 = base.F32_add(base.F32_mul(v119, v119), v109)
						v139 = base.F32_add(base.F32_mul(v117, v117), v110)
					} else {
						v47 = int32(0)
						v55 = int32(0)
						v58 = v10
						v59 = v10
						v60 = v10
						for {
							v64 = int32(2)
							v65 = v47 << (uint(v64) % 32)
							v67 = v65 | int32(4)
							v69 = *(*float32)(unsafe.Add(mBase, uint32(v40+v67)))
							v71 = *(*float32)(unsafe.Add(mBase, uint32(v38+v67)))
							v74 = *(*float32)(unsafe.Add(mBase, uint32(v40+v65)))
							v76 = *(*float32)(unsafe.Add(mBase, uint32(v38+v65)))
							v79 = base.F32_add(base.F32_mul(v69, v71), base.F32_add(base.F32_mul(v74, v76), v58))
							v83 = base.F32_add(base.F32_mul(v71, v71), base.F32_add(base.F32_mul(v76, v76), v59))
							v87 = base.F32_add(base.F32_mul(v69, v69), base.F32_add(base.F32_mul(v74, v74), v60))
							v89 = v47 + v64
							v91 = v55 + v64
							if v91 != v34&int32(_a_F_cosine_distance_0) {
								v47 = v89
								v55 = v91
								v58 = v79
								v59 = v83
								v60 = v87
								continue
							} else {
								break
							}
							break
						}
						if v34&int32(1) == int32(0) {
							v137 = v79
							v138 = v83
							v139 = v87
						} else {
							v97 = v89
							v108 = v79
							v109 = v83
							v110 = v87
							v115 = v97 << (uint(int32(2)) % 32)
							v117 = *(*float32)(unsafe.Add(mBase, uint32(v40+v115)))
							v119 = *(*float32)(unsafe.Add(mBase, uint32(v115+v38)))
							v137 = base.F32_add(base.F32_mul(v117, v119), v108)
							v138 = base.F32_add(base.F32_mul(v119, v119), v109)
							v139 = base.F32_add(base.F32_mul(v117, v117), v110)
						}
					}
					v143 = float64(1)
					v149 = base.F64_div(base.F64_promote_f32(v137), base.F64_sqrt(base.F64_mul(base.F64_promote_f32(v138), base.F64_promote_f32(v139))))
					if base.F64_gt(v149, v143) != 0 {
						v174 = v143
					} else {
						if base.F64_lt(v149, float64(-1)) == int32(0) {
							v174 = v149
						} else {
							v174 = float64(-1)
						}
					}
				}
				m.G0 = v20 + int32(16)
				return base.I64_reinterpret_f64(base.F64_sub(float64(1), v174))
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v188 = m.ExcPending
					if v188 != 0 {
						return int64(0)
					} else {
						v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
						v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v190
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v189
						F_errmsg(m, int32(_a_F_cosine_distance_1), v20)
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_cosine_distance_2), int32(76), int32(_a_F_cosine_distance_3))
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
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
}
func F_cost_subplan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v87 float64
	_ = v87
	var v91 float64
	_ = v91
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v98 float64
	_ = v98
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 float64
	_ = v116
	var v120 float64
	_ = v120
	var v122 float64
	_ = v122
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
	v3 = float64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = F_make_ands_implicit(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v17
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v17
	if v15 == v19 {
		v56 = v3
		v63 = float64(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v64 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v65 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v27 <= int32(0) {
		v56 = v3
		v63 = float64(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v38 = int32(0)
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v38<<(uint(int32(2))%32))))
	v46 = F_cost_qual_eval_walker(m, v43, v12+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v56 = v52
	v63 = v53
	goto L3
L8:
	;
	v49 = v38 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v49 < v50 {
		v38 = v49
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+64)) = v122
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v126
	m.G0 = v12 + int32(32)
	return
L11:
	;
	v69 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subplan[0]))
	v70 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v122 = v56
	v124 = base.F64_add(v63, base.F64_add(base.F64_mul(v69, v70), v64))
	goto L10
L12:
	;
	goto L13
L13:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v75 = base.F64_sub(v64, v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v76 {
	case 0:
		goto L17
	case 1, 2:
		goto L16
	default:
		goto L15
	}
L14:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v107 != 0 {
		v120 = v74
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v106 = base.F64_add(v56, v75)
	goto L14
L16:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v95 = float64(0.5)
	v98 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subplan[0]))
	v106 = base.F64_add(base.F64_mul(base.F64_mul(v94, v95), v98), base.F64_add(base.F64_mul(v75, v95), v56))
	goto L14
L17:
	;
	v77 = float64(1e+100)
	v78 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.F64_gt(v78, v77)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v78)&int64(9223372036854775807))) != 0 {
		v91 = v77
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v106 = base.F64_add(v56, base.F64_div(v75, v91))
	goto L14
L19:
	;
	v87 = float64(1)
	if base.F64_le(v78, v87) != 0 {
		v91 = v87
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v91 = base.F64_nearest(v78)
	goto L18
L21:
	;
	v122 = base.F64_add(v106, v120)
	v124 = v63
	goto L10
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v110 = v108 - int32(352)
	goto L23
L23:
	;
	v116 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v110) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_cost_subplan_0))>>(uint(v110)%32)) == int32(0) {
		v120 = v116
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v122 = v106
	v124 = base.F64_add(v63, v116)
	goto L10
}
