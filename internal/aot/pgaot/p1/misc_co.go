package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(-4294967296)
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
	var v3 int32
	_ = v3
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v65 int32
	_ = v65
	v3 = int32(12)
	goto L4
L1:
	;
	return v65
L2:
	;
	v65 = int32(0)
	goto L1
L3:
	;
	v39 = v34
	v40 = v35
	v41 = v36
	goto L13
L4:
	;
	if (l0|l1)&int32(3) != 0 {
		v34 = l0
		v35 = l1
		v36 = v3
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v24 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v11 = l0
	v12 = l1
	v13 = v3
	goto L8
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v16 != v17 {
		v34 = v11
		v35 = v12
		v36 = v13
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v19 = int32(4)
	v20 = v12 + v19
	v22 = v11 + v19
	v24 = v13 - v19
	if base.Ui32(int32(3)) < base.Ui32(v24) {
		v11 = v22
		v12 = v20
		v13 = v24
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v34 = v22
	v35 = v20
	v36 = v24
	goto L3
L13:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v44 == v45 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v65 = v44 - v45
	goto L1
L15:
	;
	v47 = int32(1)
	v52 = v41 - v47
	if v52 != 0 {
		v39 = v39 + v47
		v40 = v40 + v47
		v41 = v52
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
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
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v7
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
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v7
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(993)
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(993)
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
					F_errfinish(m, int32(_a_F_commit_prepared_cb_wrapper_4), int32(1031), int32(_a_F_commit_prepared_cb_wrapper_5))
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
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_gt(v7, v8) != 0 {
		v10 = int32(1)
	} else {
		v10 = int32(-1)
	}
	if base.F64_ne(v7, v8) != 0 {
		v13 = v10
	} else {
		v13 = int32(0)
	}
	return v13
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	v11 = int32(0)
	goto L6
L4:
	;
	if v51 != 0 {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v45 = int32(0)
	v51 = base.B2i32(v24 == v45)
	v53 = base.B2i32(v33 != v45) << (uint(int32(1)) % 32)
	goto L4
L6:
	;
	v14 = int32(0)
	if l0 == v14 {
		v24 = v14
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 <= v11 {
		v24 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = v20 + v11<<(uint(int32(2))%32)
	goto L8
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v33 = v30 + v11<<(uint(int32(2))%32)
	if v24 == int32(0) {
		goto L5
	} else {
		goto L16
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 < v25 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v27 = int32(0)
	v51 = base.B2i32(v24 == v27)
	v53 = v27
	goto L4
L15:
	;
	goto L14
L16:
	;
	if v33 == int32(0) {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v40 == v41 {
		v11 = v11 + int32(1)
		goto L6
	} else {
		goto L18
	}
L18:
	;
	goto L7
L19:
	;
	v55 = v53
	goto L21
L20:
	;
	v55 = int32(1)
	goto L21
L21:
	;
	return v55
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		m.Env.Pgmem_zstream_free(m, v5)
		mBase = m.M
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		if v7 != 0 {
			F_ResourceOwnerForget(m, v7, v4, int32(_a_F_compress_free_0))
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_pfree(m, v4)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v16 = F___memset(m, l0, int32(0), int32(_a_F_compress_free_1))
					mBase = m.M
					F_pfree(m, l0)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_pfree(m, v4)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v16 = F___memset(m, l0, int32(0), int32(_a_F_compress_free_1))
				mBase = m.M
				F_pfree(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v16 = F___memset(m, l0, int32(0), int32(_a_F_compress_free_1))
		mBase = m.M
		F_pfree(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
				v20 = F_palloc(m, v11<<(uint(int32(2))%32))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v20
					v23 = F_palloc(m, v11)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v23
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v28 = v26 + int32(1)
						v31 = F_palloc(m, v28<<(uint(int32(2))%32))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v31
							v34 = F_palloc(m, v28)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v34
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(0)
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
													F_errfinish(m, int32(_a_F_cookDefault_2), int32(3385), int32(_a_F_cookDefault_3))
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
									F_errfinish(m, int32(_a_F_cookDefault_2), int32(3348), int32(_a_F_cookDefault_3))
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
																	F_errfinish(m, int32(_a_F_cookDefault_2), int32(3385), int32(_a_F_cookDefault_3))
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
																		F_errfinish(m, int32(_a_F_cookDefault_2), int32(3385), int32(_a_F_cookDefault_3))
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
	var v16 int32
	_ = v16
	v4 = l3
	v6 = l5
	if base.Ui32(v4) <= base.Ui32(int32(24)) {
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)) = uint8(v4)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v6)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l4
		v12 = int32(17)
		*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v12)
		if v4 != 0 {
			v16 = F__emscripten_memcpy_bulkmem(m, l1+int32(12), l2, v4)
			mBase = m.M
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	} else {
	}
	return
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
func F_cost_subplan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
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
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v63 float64
	_ = v63
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v92 float64
	_ = v92
	var v96 float64
	_ = v96
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 float64
	_ = v122
	var v127 float64
	_ = v127
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	v6 = float64(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = F_make_ands_implicit(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v19
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v23
	if v17 == v23 {
		v63 = v6
		v69 = float64(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v70 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v71 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v29 <= int32(0) {
		v63 = v6
		v69 = float64(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v36 = int32(0)
	goto L6
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v36<<(uint(int32(2))%32))))
	v50 = F_cost_qual_eval_walker(m, v47, v14+int32(8))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v63 = v56
	v69 = v57
	goto L3
L8:
	;
	v53 = v36 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v53 < v54 {
		v36 = v53
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+64)) = v130
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v132
	m.G0 = v14 + int32(32)
	return
L11:
	;
	v75 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subplan[0]))
	v76 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v130 = v63
	v132 = base.F64_add(v69, base.F64_add(base.F64_mul(v75, v76), v70))
	goto L10
L12:
	;
	goto L13
L13:
	;
	v80 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v81 = base.F64_sub(v70, v80)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v82 {
	case 0:
		goto L17
	case 1, 2:
		goto L16
	default:
		goto L15
	}
L14:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v113 != 0 {
		v127 = v80
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v112 = base.F64_add(v63, v81)
	goto L14
L16:
	;
	v99 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v100 = float64(0.5)
	v103 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subplan[0]))
	v112 = base.F64_add(base.F64_mul(base.F64_mul(v99, v100), v103), base.F64_add(base.F64_mul(v81, v100), v63))
	goto L14
L17:
	;
	v83 = float64(1e+100)
	v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.F64_gt(v84, v83) != 0 {
		v96 = v83
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v112 = base.F64_add(v63, base.F64_div(v81, v96))
	goto L14
L19:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v84)&int64(9223372036854775807)) {
		v96 = v83
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v92 = float64(1)
	if base.F64_le(v84, v92) != 0 {
		v96 = v92
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v96 = base.F64_nearest(v84)
	goto L18
L22:
	;
	v130 = base.F64_add(v112, v127)
	v132 = v69
	goto L10
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v116 = v114 - int32(348)
	goto L24
L24:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v116) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_cost_subplan_0))>>(uint(v116)%32)) == int32(0) {
		v127 = v122
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v130 = v112
	v132 = base.F64_add(v69, v122)
	goto L10
}
