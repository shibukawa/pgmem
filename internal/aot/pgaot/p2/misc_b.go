package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BarrierArriveAndWait(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	v7 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_BarrierArriveAndWait_0), int32(132), int32(_a_F_BarrierArriveAndWait_1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = int32(1)
	v17 = v15 + v16
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = v19 + v16
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 == v17 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
	v26 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v26
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v26))
	F_ConditionVariableBroadcast(m, l0+int32(24))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v37 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v37))
	v41 = l0 + int32(24)
	F_ConditionVariablePrepareToSleep(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	return int32(1)
L10:
	;
	goto L11
L11:
	;
	v50 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_s_lock(m, l0, int32(_a_F_BarrierArriveAndWait_0), int32(173), int32(_a_F_BarrierArriveAndWait_1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 == v56 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v21 != v58 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v68 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v68))
	F_ConditionVariableSleep(m, v41, l1)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L24
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v21
	goto L22
L21:
	;
	goto L22
L22:
	;
	v61 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v61))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	return base.B2i32(v58 != v21)
L24:
	;
	goto L11
}
func F_BarrierDetach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		F_s_lock(m, l0, int32(_a_F_BarrierDetach_0), int32(307), int32(_a_F_BarrierDetach_1))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = v11 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13
			if int32(0) < v13 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v17 == v13 {
					v22 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v22
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v24 + int32(1)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v22))
					F_ConditionVariableBroadcast(m, l0+int32(24))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						return
					}
				} else {
					v19 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v19))
					return
				}
			} else {
				v19 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v19))
				return
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = v11 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13
		if int32(0) < v13 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v17 == v13 {
				v22 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v24 + int32(1)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v22))
				F_ConditionVariableBroadcast(m, l0+int32(24))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					return
				}
			} else {
				v19 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v19))
				return
			}
		} else {
			v19 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v19))
			return
		}
	}
}
func F_BasicOpenFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
	v14 = F_open(m, l0, l1, v9+int32(16))
	mBase = m.M
	if int32(0) <= v14 {
		v68 = v14
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v68
L2:
	;
	goto L3
L3:
	;
	v23 = int32(-1)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_BasicOpenFilePerm[0]))
	switch v25 - int32(33) {
	case 0, 8:
		goto L5
	default:
		v68 = v23
		goto L1
	}
L4:
	;
	v68 = v61
	goto L1
L5:
	;
	v30 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BasicOpenFilePerm[0])) = v47
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_BasicOpenFilePerm[1]))
	if v50 <= v47 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	F_errmsg(m, int32(_a_F_BasicOpenFilePerm_0), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_BasicOpenFilePerm_1), int32(1169), int32(_a_F_BasicOpenFilePerm_2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BasicOpenFilePerm[0])) = v25
	v68 = v23
	goto L1
L15:
	;
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_BasicOpenFilePerm[2]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	F_LruDelete(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
	v61 = F_open(m, l0, l1, v9)
	mBase = m.M
	if v61 < int32(0) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L4
}
func F_BlessTupleDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2 != int32(2249) {
		return l0
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if int32(0) <= v5 {
			return l0
		} else {
			F_assign_record_type_typmod(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return l0
			}
		}
	}
}
func F_BogusFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_errmsg_internal(m, int32(_a_F_BogusFree_0), v5)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_errfinish(m, int32(_a_F_BogusFree_1), int32(292), int32(_a_F_BogusFree_2))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_BuildCallback_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
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
	var v44 int32
	_ = v44
	var v47 float64
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v9 = int32(0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v12 == v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L32
	}
L2:
	;
	v15 = int32(_a_F_BuildCallback_2_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_2[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l5)+168))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_2[0])) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+68))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	return
L5:
	;
	return
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l5)+52))
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_2[0])) = v16
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l5)+168))
	F_MemoryContextReset(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L31
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
	v27 = F_IvfflatCheckNorm(m, v25, v26, v23)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v35 = v23
	goto L10
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if int32(0) < v36 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if v27 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
	v33 = F_HnswNormValue(m, v31, v32, v23)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v35 = v33
	goto L10
L14:
	;
	v44 = int32(0)
	v47 = float64(1.7976931348623157e+308)
	v49 = v9
	goto L17
L15:
	;
	v78 = v9
	goto L16
L16:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	m.T0[v82].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L28
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v52 <= v44 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v78 = v65
	goto L16
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v60 = F_FunctionCall2Coll(m, v54, v55, v35, v56+v57*v44)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
	v63 = base.F64_gt(v47, v62)
	if v63 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v64 = v62
	goto L23
L22:
	;
	v64 = v47
	goto L23
L23:
	;
	if v63 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v65 = v44
	goto L26
L25:
	;
	v65 = v49
	goto L26
L26:
	;
	v67 = v44 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v67 < v68 {
		v44 = v67
		v47 = v64
		v49 = v65
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v78
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v88 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v35
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+2)) = uint8(v88)
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	v98 = v96 & int32(_a_F_BuildCallback_2_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)) = uint16(v98)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+6)) = uint16(v101)
	goto L29
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l5)+152))
	F_tuplesort_puttupleslot(m, v103, v20)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v106 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+32)) = base.F64_add(v106, float64(1))
	goto L7
L31:
	;
	goto L4
L32:
	;
	F_errmsg_internal(m, int32(_a_F_BuildCallback_2_2), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_BuildCallback_2_3), int32(326), int32(_a_F_BuildCallback_2_4))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BuildDescForRelation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l0 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v172
L2:
	;
	v19 = F_CreateTemplateTupleDesc(m, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = F_CreateTemplateTupleDesc(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v172 = v19
	goto L1
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 <= int32(0) {
		v172 = v24
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v35 = v2
	v37 = v2
	goto L10
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L40
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v37<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	F_typenameTypeIdAndMod(m, int32(0), v47, v14+int32(12), v14+int32(8))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L12
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L36
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_BuildDescForRelation[0]))
	v59 = F_object_aclcheck(m, int32(1247), v55, v57, int64(256))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	if v59 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_aclcheck_error_type(m, v59, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v64 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v67 = F_GetColumnDefCollation(m, v64, v44, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L11
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if int32(_a_F_BuildDescForRelation_0) <= v71 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	v74 = v64
	goto L22
L22:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+12)))
	if v75 == int32(1) {
		goto L9
	} else {
		goto L24
	}
L23:
	;
	v74 = v71
	goto L22
L24:
	;
	v80 = base.I32_extend16_s(v35 + int32(1))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	F_TupleDescInitEntry(m, v24, v80, v45, v81, v82, v74)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v24+v85<<(uint(int32(4))%32)+v80*int32(100))+16)) = v67
	goto L26
L26:
	;
	v99 = v24 + v85<<(uint(int32(4))%32) + v80*int32(100)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+6)) = uint8(v100)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+12)) = uint8(v102)
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+14)) = uint16(v104)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+9)) = uint8(v106)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+10)) = uint8(v108)
	v111 = v99 - int32(80)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+68))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v114 = F_GetAttributeCompression(m, v112, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+5)) = uint8(v114)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+21)))
	if v119 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_populate_compact_attribute(m, v24, v80-int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L34
	}
L29:
	;
	v127 = v119
	goto L31
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	if v120 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+84)) = uint8(v127)
	goto L28
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v111)+68))
	v124 = F_GetAttributeStorage(m, v123, v120)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v127 = v124
	goto L31
L34:
	;
	v133 = v37 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v134 <= v133 {
		v172 = v24
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v35 = v80
	v37 = v133
	goto L10
L36:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_BuildDescForRelation_1), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_BuildDescForRelation_2), int32(1425), int32(_a_F_BuildDescForRelation_3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v45
	F_errmsg(m, int32(_a_F_BuildDescForRelation_4), v14)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_BuildDescForRelation_2), int32(1431), int32(_a_F_BuildDescForRelation_3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BuildParameterizedTidPaths(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v13 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = v4
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v21<<(uint(int32(2))%32))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+10)))
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v69 = v21 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v69 < v70 {
		v21 = v69
		goto L4
	} else {
		goto L24
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	if base.Ui32(v29) < base.Ui32(v28) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v33&int32(1) == int32(0) {
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+13)))
	v33 = v31
	goto L11
L10:
	;
	v33 = int32(1)
	goto L11
L11:
	;
	goto L8
L12:
	;
	v38 = F_IsBinaryTidClause(m, v26, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v38 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v43 != int32(387) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v46 = F_join_clause_is_movable_to(m, v26, l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	if v46 == int32(0) {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v26
	v55 = F_list_make1_impl(m, int32(1), v9+int32(8))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v59 = F_bms_union(m, v57, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v62 = F_bms_del_member(m, v59, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v64 = F_create_tidscan_path(m, l0, l1, v55, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	F_add_path(m, l1, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	goto L6
L24:
	;
	goto L5
}
func F_begin_tup_output_tupdesc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v6 = F_palloc(m, int32(8))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_MakeTupleTableSlot(m, l1, l2)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			m.T0[v15].(func(*base.Module, int32, int32, int32))(m, l0, int32(1), l1)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v6
			}
		}
	}
}
func F_bitfromint4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = v12 - int32(2147483641)
	if base.Ui32(v14) < base.Ui32(int32(-2147483640)) {
		v17 = int32(1)
	} else {
		v17 = v12
	}
	v20 = int32(8)
	v21 = base.I32_div_s(v17+int32(7), v20)
	v23 = v21 + v20
	v24 = F_palloc(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v24))) = v23 << (uint(int32(2)) % 32)
		v33 = v24 + int32(8)
		v34 = int32(32)
		if v34 <= v17 {
			v37 = v34
		} else {
			v37 = v17
		}
		v39 = v37 + int32(8)
		if v39 <= v17 {
			v41 = int32(-2147483640)
			if base.Ui32(v14) <= base.Ui32(v41) {
				v44 = v41
			} else {
				v44 = v14
			}
			v46 = v44 + int32(2147483633)
			v47 = v46 - v37
			v49 = int32(base.Ui32(v47) >> (uint(int32(3)) % 32))
			v51 = v49 + int32(1)
			if v51 != 0 {
				base.MemoryFill(m, v33, v10>>(uint(int32(31))%32), v51)
			} else {
			}
			v61 = v24 + v49 + int32(9)
			v62 = v46 - v47&int32(-8)
		} else {
			v61 = v33
			v62 = v17
		}
		if v37 < v62 {
			v74 = v62 - int32(8)
			v76 = int32(-1)<<(uint(v39-v62)%32)&(v10>>(uint(int32(31))%32)) | v10>>(uint(v74)%32)
			*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v76)
			v80 = v61 + int32(1)
			v81 = v74
		} else {
			v80 = v61
			v81 = v62
		}
		if v81 < int32(8) {
			v171 = v80
			v172 = v81
		} else {
			v85 = v81 - int32(8)
			v86 = int32(56)
			if v85&v86 != v86 {
				v97 = v80
				v98 = v81
				v103 = int32(0)
				for {
					v107 = v98 - int32(8)
					v108 = v10 >> (uint(v107) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v108)
					v110 = int32(1)
					v111 = v97 + v110
					v113 = v103 + v110
					if v113 != (int32(base.Ui32(v85)>>(uint(int32(3))%32))+int32(1))&int32(7) {
						v97 = v111
						v98 = v107
						v103 = v113
						continue
					} else {
						break
					}
					break
				}
				v115 = v111
				v116 = v107
			} else {
				v115 = v80
				v116 = v81
			}
			if base.Ui32(v85) < base.Ui32(int32(56)) {
				v171 = v115
				v172 = v116
			} else {
				v126 = v115
				v127 = v116
				for {
					v136 = v127 + int32(-64)
					v137 = v10 >> (uint(v136) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v126)+7)) = uint8(v137)
					v141 = v10 >> (uint(v127-int32(56)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v126)+6)) = uint8(v141)
					v145 = v10 >> (uint(v127-int32(48)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v126)+5)) = uint8(v145)
					v149 = v10 >> (uint(v127-int32(40)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v126)+4)) = uint8(v149)
					v153 = v10 >> (uint(v127-int32(32)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v126)+3)) = uint8(v153)
					v157 = v10 >> (uint(v127-int32(24)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v126)+2)) = uint8(v157)
					v161 = v10 >> (uint(v127-int32(16)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)) = uint8(v161)
					v163 = int32(8)
					v165 = v10 >> (uint(v127-v163) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v165)
					v168 = v126 + v163
					if int32(71) < v127 {
						v126 = v168
						v127 = v136
						continue
					} else {
						break
					}
					break
				}
				v171 = v168
				v172 = v136
			}
		}
		if int32(0) < v172 {
			v184 = v10 << (uint(int32(8)-v172) % 32)
			*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v184)
		} else {
		}
		return v24
	}
}
func F_bitfromint8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v101 int64
	_ = v101
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v128 int64
	_ = v128
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = v15 - int32(2147483641)
	if base.Ui32(v17) < base.Ui32(int32(-2147483640)) {
		v20 = int32(1)
	} else {
		v20 = v15
	}
	v23 = int32(8)
	v24 = base.I32_div_s(v20+int32(7), v23)
	v26 = v24 + v23
	v27 = F_palloc(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v27))) = v26 << (uint(int32(2)) % 32)
		v36 = v27 + int32(8)
		v37 = int32(64)
		if v37 <= v20 {
			v40 = v37
		} else {
			v40 = v20
		}
		v42 = v40 + int32(8)
		if v42 <= v20 {
			v44 = int32(-2147483640)
			if base.Ui32(v17) <= base.Ui32(v44) {
				v47 = v44
			} else {
				v47 = v17
			}
			v49 = v47 + int32(2147483633)
			v50 = v49 - v40
			v52 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
			v54 = v52 + int32(1)
			if v54 != 0 {
				base.MemoryFill(m, v36, base.I32_wrap_i64(v13>>(uint(int64(63))%64)), v54)
			} else {
			}
			v65 = v27 + v52 + int32(9)
			v66 = v49 - v50&int32(-8)
		} else {
			v65 = v36
			v66 = v20
		}
		if v40 < v66 {
			v79 = v66 - int32(8)
			v83 = base.I32_wrap_i64(v13>>(uint(int64(63))%64))&(int32(-1)<<(uint(v42-v66)%32)) | base.I32_wrap_i64(v13>>(uint(base.I64_extend_i32_u(v79))%64))
			*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v83)
			v87 = v65 + int32(1)
			v88 = v79
		} else {
			v87 = v65
			v88 = v66
		}
		if int32(8) <= v88 {
			v92 = v87
			v101 = base.I64_extend_i32_u(v88)
			for {
				v104 = v101 - int64(8)
				v105 = v13 >> (uint(v104) % 64)
				*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v105)
				v108 = v92 + int32(1)
				if base.Ui64(int64(15)) < base.Ui64(v101) {
					v92 = v108
					v101 = v104
					continue
				} else {
					break
				}
				break
			}
			v112 = v108
			v113 = base.I32_wrap_i64(v104)
		} else {
			v112 = v87
			v113 = v88
		}
		if int32(0) < v113 {
			v128 = v13 << (uint(base.I64_extend_i32_u(int32(8)-v113)) % 64)
			*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v128)
		} else {
		}
		return v27
	}
}
func F_bitgt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int32(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = v15 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v20 = int32(2)
	v21 = int32(base.Ui32(v19) >> (uint(v20) % 32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v24 = int32(base.Ui32(v22) >> (uint(v20) % 32))
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v21
	goto L7
L6:
	;
	v26 = v24
	goto L7
L7:
	;
	v28 = v26 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v28) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v90 != 0 {
		v99 = v90
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v90 = int32(0)
	goto L8
L10:
	;
	v64 = v59
	v65 = v60
	v66 = v61
	goto L20
L11:
	;
	if (v13|v18)&int32(3) != 0 {
		v59 = v13
		v60 = v18
		v61 = v28
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v52 = v13
	v53 = v18
	v54 = v28
	goto L13
L13:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v36 = v13
	v37 = v18
	v38 = v28
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v41 != v42 {
		v59 = v36
		v60 = v37
		v61 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v52 = v47
	v53 = v45
	v54 = v49
	goto L13
L17:
	;
	v44 = int32(4)
	v45 = v37 + v44
	v47 = v36 + v44
	v49 = v38 - v44
	if base.Ui32(int32(3)) < base.Ui32(v49) {
		v36 = v47
		v37 = v45
		v38 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v59 = v52
	v60 = v53
	v61 = v54
	goto L10
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == v70 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v90 = v69 - v70
	goto L8
L22:
	;
	v72 = int32(1)
	v77 = v66 - v72
	if v77 != 0 {
		v64 = v64 + v72
		v65 = v65 + v72
		v66 = v77
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
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v92 == v93 {
		v99 = int32(0)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v92 < v93 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = int32(-1)
	goto L30
L29:
	;
	v98 = int32(1)
	goto L30
L30:
	;
	v99 = v98
	goto L1
L31:
	;
	F_pfree(m, v8)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v106 != v15 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_pfree(m, v15)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	return base.B2i32(int32(0) < v99)
L38:
	;
	goto L37
}
func F_bitle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int32(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = v15 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v20 = int32(2)
	v21 = int32(base.Ui32(v19) >> (uint(v20) % 32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v24 = int32(base.Ui32(v22) >> (uint(v20) % 32))
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v21
	goto L7
L6:
	;
	v26 = v24
	goto L7
L7:
	;
	v28 = v26 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v28) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v90 != 0 {
		v99 = v90
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v90 = int32(0)
	goto L8
L10:
	;
	v64 = v59
	v65 = v60
	v66 = v61
	goto L20
L11:
	;
	if (v13|v18)&int32(3) != 0 {
		v59 = v13
		v60 = v18
		v61 = v28
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v52 = v13
	v53 = v18
	v54 = v28
	goto L13
L13:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v36 = v13
	v37 = v18
	v38 = v28
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v41 != v42 {
		v59 = v36
		v60 = v37
		v61 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v52 = v47
	v53 = v45
	v54 = v49
	goto L13
L17:
	;
	v44 = int32(4)
	v45 = v37 + v44
	v47 = v36 + v44
	v49 = v38 - v44
	if base.Ui32(int32(3)) < base.Ui32(v49) {
		v36 = v47
		v37 = v45
		v38 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v59 = v52
	v60 = v53
	v61 = v54
	goto L10
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == v70 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v90 = v69 - v70
	goto L8
L22:
	;
	v72 = int32(1)
	v77 = v66 - v72
	if v77 != 0 {
		v64 = v64 + v72
		v65 = v65 + v72
		v66 = v77
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
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v92 == v93 {
		v99 = int32(0)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v92 < v93 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = int32(-1)
	goto L30
L29:
	;
	v98 = int32(1)
	goto L30
L30:
	;
	v99 = v98
	goto L1
L31:
	;
	F_pfree(m, v8)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v106 != v15 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_pfree(m, v15)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	return base.B2i32(v99 <= int32(0))
L38:
	;
	goto L37
}
func F_bitshiftleft(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v16 < int32(0) {
			v20 = int32(0)
			v22 = int32(-2147483640)
			if base.Ui32(v16) <= base.Ui32(v22) {
				v25 = v22
			} else {
				v25 = v16
			}
			v27 = F_DirectFunctionCall2Coll(m, int32(1533), v20, v12, v20-v25)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				return v27
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v33 = F_palloc(m, int32(base.Ui32(v30)>>(uint(int32(2))%32)))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v37 = v35 & int32(-4)
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v39
				v42 = v33 + int32(8)
				if v39 <= v16 {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v46 = int32(base.Ui32(v44) >> (uint(int32(2)) % 32))
					v48 = v46 - int32(8)
					if v42&int32(3)|v44&int32(12)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v48)) == int32(0) {
						v59 = v33 + v46
						if base.Ui32(v59) <= base.Ui32(v42) {
							return v33
						} else {
							v62 = v33 + int32(12)
							if base.Ui32(v62) < base.Ui32(v59) {
								v64 = v59
							} else {
								v64 = v62
							}
							v71 = (v64-v33-int32(9))&int32(-4) + int32(4)
							if v71 == int32(0) {
								return v33
							} else {
								v204 = v42
								v205 = v71
								base.MemoryFill(m, v204, int32(0), v205)
								return v33
							}
						}
					} else {
						if v48 == int32(0) {
							return v33
						} else {
							v204 = v42
							v205 = v48
							base.MemoryFill(m, v204, int32(0), v205)
							return v33
						}
					}
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v78 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
					v80 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
					v82 = v80 + int32(8)
					v83 = v12 + v82
					v85 = v16 & int32(7)
					if v85 != 0 {
						if base.Ui32(v82) < base.Ui32(v78) {
							v89 = v42
							v93 = v83
							for {
								v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
								v100 = v99 << (uint(v85) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v100)
								v103 = v93 + int32(1)
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v106 = int32(base.Ui32(v104) >> (uint(int32(2)) % 32))
								if base.Ui32(v103) < base.Ui32(v12+v106) {
									v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
									v111 = int32(base.Ui32(v109)>>(uint(int32(8)-v85)%32)) | v100
									*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v111)
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v116 = int32(base.Ui32(v113) >> (uint(int32(2)) % 32))
								} else {
									v116 = v106
								}
								v118 = v89 + int32(1)
								if base.Ui32(v103) < base.Ui32(v12+v116) {
									v89 = v118
									v93 = v103
									continue
								} else {
									break
								}
								break
							}
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							v122 = v118
							v132 = v121
						} else {
							v122 = v42
							v132 = v37
						}
						if base.Ui32(v33+int32(base.Ui32(v132)>>(uint(int32(2))%32))) <= base.Ui32(v122) {
						} else {
							v137 = v122
							for {
								v147 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v147)
								v150 = v137 + int32(1)
								v151 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
								if base.Ui32(v150) < base.Ui32(v33+int32(base.Ui32(v151)>>(uint(int32(2))%32))) {
									v137 = v150
									continue
								} else {
									break
								}
								break
							}
						}
						return v33
					} else {
						v156 = v78 - v80
						v158 = v156 - int32(8)
						if v158 != 0 {
							base.MemoryCopy(m, v42, v83, v158)
						} else {
						}
						v162 = v156 + v33
						if v16&int32(24)|(v162&int32(3)|base.B2i32(base.Ui32(int32(_a_F_bitshiftleft_0)) < base.Ui32(v16))) == int32(0) {
							v171 = v33 + v78
							if base.Ui32(v171) <= base.Ui32(v162) {
								return v33
							} else {
								v175 = v171 - v80 + int32(4)
								if base.Ui32(v171) < base.Ui32(v175) {
									v177 = v175
								} else {
									v177 = v171
								}
								v186 = (v177+v80+(v33^int32(-1))-v78)&int32(-4) + int32(4)
								if v186 == int32(0) {
									return v33
								} else {
									v204 = v162
									v205 = v186
									base.MemoryFill(m, v204, int32(0), v205)
									return v33
								}
							}
						} else {
							if v80 == int32(0) {
							} else {
								base.MemoryFill(m, v162, int32(0), v80)
							}
							return v33
						}
					}
				}
			}
		}
	}
}
func F_bittypmodout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_palloc(m, int32(64))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v8 {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
			v19 = F_pg_snprintf(m, v10, int32(64), int32(_a_F_bittypmodout_0), v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v10
			}
		} else {
			v21 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v21)
			m.G0 = v6 + int32(16)
			return v10
		}
	}
}
func F_blbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v8 = m.G0
	v10 = v8 - int32(_a_F_blbuild_0)
	m.G0 = v10
	v13 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			F_BloomInitMetapage(m, l1, int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = v10 + int32(8)
				base.MemoryFill(m, v23, int32(0), int32(_a_F_blbuild_1))
				F_initBloomState(m, v23, l1)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_blbuild[0]))
					v35 = F_AllocSetContextCreateInternal(m, v30, int32(_a_F_blbuild_2), int32(0), int32(_a_F_blbuild_3), int32(_a_F_blbuild_4))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+1184)) = v35
						v39 = v10 + int32(1192)
						v40 = int32(0)
						F_PageInit(m, v39, int32(_a_F_blbuild_3), int32(8))
						mBase = m.M
						v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+16)))
						v45 = v39 + v44
						v46 = int32(_a_F_blbuild_5)
						*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)) = uint16(v46)
						*(*uint16)(unsafe.Add(mBase, uint32(v45)+2)) = uint16(v40)
						v49 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_blbuild[1]))) = v49
						v51 = int32(1)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+140))
						v60 = m.T0[v59].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v51, v49, v51, v49, int32(-1), int32(_a_F_blbuild_6), v23, v49)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_blbuild[1])))
							if int32(0) < v62 {
								v65 = F_BloomNewBuffer(m, l1)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = F_GenericXLogStart(m, l1)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v70 = F_GenericXLogRegisterBuffer(m, v67, v65, int32(1))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											base.MemoryCopy(m, v70, v39, int32(_a_F_blbuild_3))
											F_GenericXLogFinish(m, v67)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												F_UnlockReleaseBuffer(m, v65)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return int32(0)
												} else {
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1184))
													F_MemoryContextDelete(m, v80)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int32(0)
													} else {
														v84 = F_palloc(m, int32(16))
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															*(*float64)(unsafe.Add(mBase, uint32(v84))) = v60
															v87 = *(*int64)(unsafe.Add(mBase, uint32(v10)+1176))
															*(*float64)(unsafe.Add(mBase, uint32(v84)+8)) = base.F64_convert_i64_s(v87)
															m.G0 = v10 + int32(_a_F_blbuild_0)
															return v84
														}
													}
												}
											}
										}
									}
								}
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1184))
								F_MemoryContextDelete(m, v80)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									v84 = F_palloc(m, int32(16))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v84))) = v60
										v87 = *(*int64)(unsafe.Add(mBase, uint32(v10)+1176))
										*(*float64)(unsafe.Add(mBase, uint32(v84)+8)) = base.F64_convert_i64_s(v87)
										m.G0 = v10 + int32(_a_F_blbuild_0)
										return v84
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
			v97 = m.ExcPending
			if v97 != 0 {
				return int32(0)
			} else {
				v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v98 + int32(4)
				F_errmsg_internal(m, int32(_a_F_blbuild_7), v10)
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_blbuild_8), int32(130), int32(_a_F_blbuild_9))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
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
func F_blhandler(m *base.Module, l0 int32) int32 {
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
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v59 int32
	_ = v59
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+10)) = v7
		v9 = int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+8)) = uint16(v9)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(562954248389046)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+13)) = v7
		v15 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+19)) = v15
		v17 = int32(257)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+17)) = uint16(v17)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+27)) = uint8(v7)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(_a_F_blhandler_0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(_a_F_blhandler_1)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(_a_F_blhandler_2)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = int32(_a_F_blhandler_3)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = int32(_a_F_blhandler_4)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+76)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(_a_F_blhandler_5)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(_a_F_blhandler_6)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(_a_F_blhandler_7)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(_a_F_blhandler_8)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(_a_F_blhandler_9)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(_a_F_blhandler_10)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(_a_F_blhandler_11)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = v7
		v59 = int32(5)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+29)) = uint8(v59)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(v3)+128)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v3)+120)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v3)+112)) = v15
		return v3
	}
}
func F_blrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
		F_pfree(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
			if l1 == v10 {
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v14 <= int32(0) {
				} else {
					v18 = v14 * int32(48)
					if v18 == int32(0) {
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						base.MemoryCopy(m, v21, l1, v18)
					}
				}
			}
			return
		}
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
		if l1 == v10 {
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v14 <= int32(0) {
			} else {
				v18 = v14 * int32(48)
				if v18 == int32(0) {
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					base.MemoryCopy(m, v21, l1, v18)
				}
			}
		}
		return
	}
}
func F_booleq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.B2i32(v2 == v3) ^ base.B2i32(v5 != v3)
}
func F_boolge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.B2i32(v2 == v3) | base.B2i32(v5 != v3)
}
func F_boolout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(2))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v9)
		if v3 != 0 {
			v13 = int32(116)
		} else {
			v13 = int32(102)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v13)
		return v5
	}
}
func F_boot_get_type_io_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v9 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_boot_get_type_io_data[0]))
	if v18 == v9 {
		goto L29
	} else {
		goto L30
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L53
	} else {
		goto L57
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L53
	} else {
		goto L54
	}
L3:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v140
	m.G0 = v15 + int32(32)
	return
L4:
	;
	v105 = v103 * int32(92)
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v106)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v108)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v110)
	v112 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v112)
	if int32(1)<<(uint(v103)%32)&int32(_a_F_boot_get_type_io_data_0) != 0 {
		goto L50
	} else {
		goto L51
	}
L5:
	;
	v103 = int32(24)
	goto L4
L6:
	;
	v103 = int32(23)
	goto L4
L7:
	;
	v103 = int32(22)
	goto L4
L8:
	;
	v103 = int32(21)
	goto L4
L9:
	;
	v103 = int32(20)
	goto L4
L10:
	;
	v103 = int32(19)
	goto L4
L11:
	;
	v103 = int32(18)
	goto L4
L12:
	;
	v103 = int32(17)
	goto L4
L13:
	;
	v103 = int32(16)
	goto L4
L14:
	;
	v103 = int32(15)
	goto L4
L15:
	;
	v103 = int32(14)
	goto L4
L16:
	;
	v103 = int32(13)
	goto L4
L17:
	;
	v103 = int32(12)
	goto L4
L18:
	;
	v103 = int32(11)
	goto L4
L19:
	;
	v103 = int32(10)
	goto L4
L20:
	;
	v103 = int32(9)
	goto L4
L21:
	;
	v103 = int32(8)
	goto L4
L22:
	;
	v103 = int32(7)
	goto L4
L23:
	;
	v103 = int32(6)
	goto L4
L24:
	;
	v103 = int32(5)
	goto L4
L25:
	;
	v103 = int32(4)
	goto L4
L26:
	;
	v103 = int32(3)
	goto L4
L27:
	;
	v103 = int32(2)
	goto L4
L28:
	;
	v103 = int32(1)
	goto L4
L29:
	;
	if l0 <= int32(1001) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v35 <= int32(0) {
		goto L2
	} else {
		goto L40
	}
L32:
	;
	switch l0 - int32(16) {
	case 0:
		v103 = v9
		goto L4
	case 1:
		goto L28
	case 2:
		goto L27
	case 3:
		goto L23
	case 4:
		goto L1
	case 5:
		goto L26
	case 6:
		goto L11
	case 7:
		goto L25
	case 8:
		goto L21
	case 9:
		goto L17
	case 10:
		goto L16
	case 11:
		goto L15
	case 12:
		goto L14
	case 13:
		goto L13
	case 14:
		goto L10
	default:
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	switch l0 - int32(1002) {
	case 0:
		goto L6
	case 1, 2, 3, 4, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 27, 28, 29, 30, 31:
		goto L1
	case 5:
		goto L9
	case 7:
		goto L8
	case 26:
		goto L7
	case 32:
		goto L5
	default:
		goto L38
	}
L35:
	;
	if l0 == int32(194) {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	if l0 == int32(700) {
		goto L24
	} else {
		goto L37
	}
L37:
	;
	goto L1
L38:
	;
	switch l0 - int32(2205) {
	case 0:
		goto L22
	case 1:
		goto L20
	default:
		goto L39
	}
L39:
	;
	switch l0 - int32(4089) {
	case 0:
		goto L18
	default:
		goto L1
	case 7:
		goto L19
	}
L40:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v47 = int32(0)
	goto L42
L41:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v64)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v66)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+132)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v68)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+87)))
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v70)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)+96))
	if v72 != 0 {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38+v47<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 == l0 {
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v61 != l0 {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	v59 = v47 + int32(1)
	if v59 != v35 {
		v47 = v59
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	goto L41
L47:
	;
	v73 = v72
	goto L49
L48:
	;
	v73 = l0
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v55)+104))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v75
	v139 = v55 + int32(108)
	goto L3
L50:
	;
	v121 = l0
	goto L52
L51:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[4])))
	v121 = v120
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[5])))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v123
	v139 = v105 + int32(_a_F_boot_get_type_io_data_1)
	goto L3
L53:
	;
	return
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_boot_get_type_io_data_2), v15+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_boot_get_type_io_data_3), int32(860), int32(_a_F_boot_get_type_io_data_4))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg_internal(m, int32(_a_F_boot_get_type_io_data_5), v15)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_boot_get_type_io_data_3), int32(887), int32(_a_F_boot_get_type_io_data_4))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bpchareq(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
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
	var v58 int32
	_ = v58
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
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
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = int32(1)
	v20 = v11 + v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v25 = v23 & v19
	if v25 != 0 {
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
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L94
	}
L7:
	;
	v26 = v20
	goto L9
L8:
	;
	v26 = v11 + int32(4)
	goto L9
L9:
	;
	if v23 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v58 = v53
	goto L21
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v32 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v35 = int32(16)
	goto L16
L15:
	;
	v35 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v42 = int32(4)
	goto L19
L18:
	;
	v42 = v35
	goto L19
L19:
	;
	v53 = v42
	goto L10
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	if v58 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v76 = int32(1)
	v77 = v16 + v76
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v82 = v80 & v76
	if v82 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v75 = v53 & (v53 >> (uint(int32(31)) % 32))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v69 = v58 - int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v69))))
	if v71 == int32(32) {
		v58 = v69
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v75 = v58
	goto L23
L28:
	;
	v83 = v77
	goto L30
L29:
	;
	v83 = v16 + int32(4)
	goto L30
L30:
	;
	if v80 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v115 = v110
	goto L42
L32:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v89 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v100 = int32(1)
	if v82 != 0 {
		v110 = int32(base.Ui32(v80)>>(uint(v100)%32)) - v100
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v92 = int32(16)
	goto L37
L36:
	;
	v92 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v99 = int32(4)
	goto L40
L39:
	;
	v99 = v92
	goto L40
L40:
	;
	v110 = v99
	goto L31
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	if v115 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v133 = F_pg_newlocale_from_collation(m, v18)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	goto L43
L45:
	;
	v132 = v110 & (v110 >> (uint(int32(31)) % 32))
	goto L44
L46:
	;
	goto L47
L47:
	;
	v126 = v115 - int32(1)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v126))))
	if v128 == int32(32) {
		v115 = v126
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v132 = v115
	goto L44
L49:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v237 != v11 {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	if v135 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v132 != v75 {
		v236 = int32(0)
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v218 = int32(1)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v220&v218 != 0 {
		goto L79
	} else {
		goto L80
	}
L54:
	;
	v140 = int32(1)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v142&v140 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v145 = v140
	goto L57
L56:
	;
	v145 = int32(4)
	goto L57
L57:
	;
	v146 = v11 + v145
	v147 = int32(1)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v149&v147 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v152 = v147
	goto L60
L59:
	;
	v152 = int32(4)
	goto L60
L60:
	;
	v153 = v16 + v152
	if base.Ui32(int32(4)) <= base.Ui32(v75) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v236 = base.B2i32(v215 == int32(0))
	goto L49
L62:
	;
	v215 = int32(0)
	goto L61
L63:
	;
	v189 = v184
	v190 = v185
	v191 = v186
	goto L73
L64:
	;
	if (v146|v153)&int32(3) != 0 {
		v184 = v146
		v185 = v153
		v186 = v75
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v177 = v146
	v178 = v153
	v179 = v75
	goto L66
L66:
	;
	if v179 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v161 = v146
	v162 = v153
	v163 = v75
	goto L68
L68:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v166 != v167 {
		v184 = v161
		v185 = v162
		v186 = v163
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v177 = v172
	v178 = v170
	v179 = v174
	goto L66
L70:
	;
	v169 = int32(4)
	v170 = v162 + v169
	v172 = v161 + v169
	v174 = v163 - v169
	if base.Ui32(int32(3)) < base.Ui32(v174) {
		v161 = v172
		v162 = v170
		v163 = v174
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v184 = v177
	v185 = v178
	v186 = v179
	goto L63
L73:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v194 == v195 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v215 = v194 - v195
	goto L61
L75:
	;
	v197 = int32(1)
	v202 = v191 - v197
	if v202 != 0 {
		v189 = v189 + v197
		v190 = v190 + v197
		v191 = v202
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	goto L62
L79:
	;
	v223 = v218
	goto L81
L80:
	;
	v223 = int32(4)
	goto L81
L81:
	;
	v225 = int32(1)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v227&v225 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v230 = v225
	goto L84
L83:
	;
	v230 = int32(4)
	goto L84
L84:
	;
	v232 = F_varstr_cmp(m, v11+v223, v75, v16+v230, v132, v18)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v236 = base.B2i32(v232 == int32(0))
	goto L49
L86:
	;
	F_pfree(m, v11)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v241 != v16 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	F_pfree(m, v16)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	return v236
L93:
	;
	goto L92
L94:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_bpchareq_0), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errhint(m, int32(_a_F_bpchareq_1), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_bpchareq_2), int32(738), int32(_a_F_bpchareq_3))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bpcharge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
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
	v17 = v12 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v25 = v23 & int32(1)
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v17
	goto L6
L5:
	;
	v26 = v12 + int32(4)
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
	v58 = v53
	goto L18
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v58 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v77 = int32(1)
	v78 = v19 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v83 = v81 & v77
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v76 = v53 & (v53 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v70 = v58 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v70))))
	if v72 == int32(32) {
		v58 = v70
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v76 = v58
	goto L20
L25:
	;
	v84 = v78
	goto L27
L26:
	;
	v84 = v19 + int32(4)
	goto L27
L27:
	;
	if v81 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v116 = v111
	goto L39
L29:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v90 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v101 = int32(1)
	if v83 != 0 {
		v111 = int32(base.Ui32(v81)>>(uint(v101)%32)) - v101
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v93 = int32(16)
	goto L34
L33:
	;
	v93 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v100 = int32(4)
	goto L37
L36:
	;
	v100 = v93
	goto L37
L37:
	;
	v111 = v100
	goto L28
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v116 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v135 = int32(1)
	if v23&v135 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v134 = v111 & (v111 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v128 = v116 - int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v128))))
	if v130 == int32(32) {
		v116 = v128
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v134 = v116
	goto L41
L46:
	;
	v139 = v135
	goto L48
L47:
	;
	v139 = int32(4)
	goto L48
L48:
	;
	v141 = int32(1)
	if v81&v141 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v145 = v141
	goto L51
L50:
	;
	v145 = int32(4)
	goto L51
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v148 = F_varstr_cmp(m, v12+v139, v76, v19+v145, v134, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v150 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v154 != v19 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v19)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return int32(base.Ui32(v148^int32(-1)) >> (uint(int32(31)) % 32))
L60:
	;
	goto L59
}
func F_bpcharrecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v15 = F_pq_getmsgtext(m, v9, v10-v11, v6+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v21 = F_bpchar_input(m, v15, v19, v8, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v15)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v21
			}
		}
	}
}
func F_bpchartypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_anychar_typmodin(m, v3, int32(_a_F_bpchartypmodin_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_bqarr_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+28)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v13
	v22 = F_makepol_3(m, v10+int32(24))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v113
L2:
	;
	return int32(0)
L3:
	;
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
	v113 = v2
	goto L1
L5:
	;
	goto L6
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	if v28 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = F_errsave_start(m, v13)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(int32(134217727)) <= base.Ui32(v28) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	if v31 == int32(0) {
		v113 = v2
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errmsg(m, int32(_a_F_bqarr_in_0), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, v13, int32(_a_F_bqarr_in_1), int32(504), int32(_a_F_bqarr_in_2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v113 = v2
	goto L1
L15:
	;
	v49 = F_errsave_start(m, v13)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v70 = v28<<(uint(int32(3))%32) + int32(8)
	v71 = F_palloc(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L23
	}
L18:
	;
	if v49 == int32(0) {
		v113 = v2
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(134217726)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v28
	F_errmsg(m, int32(_a_F_bqarr_in_3), v10)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_errsave_finish(m, v13, int32(_a_F_bqarr_in_1), int32(510), int32(_a_F_bqarr_in_2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v113 = v2
	goto L1
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v70 << (uint(int32(2)) % 32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v80 = v79
	v82 = v28
	goto L24
L24:
	;
	v89 = v71 + v82<<(uint(int32(3))%32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v90)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	F_pfree(m, v80)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L26
	}
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v101 - int32(1)
	F_findoprnd_2(m, v71+int32(8), v10+int32(20))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L28
	}
L26:
	;
	v97 = int32(1)
	if base.Ui32(v97) < base.Ui32(v82) {
		v80 = v94
		v82 = v82 - v97
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = v71
	goto L1
}
func F_brinbuildempty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = l0
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v6)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v13
	v20 = F_ExtendBufferedRel(m, v6+int32(8), int32(3), int32(0), int32(9))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v22 = int32(_a_F_brinbuildempty_0)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuildempty[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_brinbuildempty[0])) = v24 + int32(1)
		if v20 < int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuildempty[1]))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v31+(v20^int32(-1))<<(uint(int32(2))%32))))
			v45 = v37
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuildempty[2]))
			v45 = v39 + v20<<(uint(int32(13))%32) + int32(-8192)
		}
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
		if v46 != 0 {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
			v49 = v47
		} else {
			v49 = int32(128)
		}
		F_PageInit(m, v45, int32(_a_F_brinbuildempty_1), int32(8))
		mBase = m.M
		v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+16)))
		v56 = int32(_a_F_brinbuildempty_2)
		*(*uint16)(unsafe.Add(mBase, uint32(v45+v54)+6)) = uint16(v56)
		*(*int32)(unsafe.Add(mBase, uint32(v45)+36)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v45)+32)) = v49
		*(*int32)(unsafe.Add(mBase, uint32(v45)+28)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v45)+24)) = int32(-1475306246)
		v64 = int32(40)
		*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)) = uint16(v64)
		F_MarkBufferDirty(m, v20)
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			F_log_newpage_buffer(m, v20, int32(1))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				v71 = int32(_a_F_brinbuildempty_0)
				v73 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuildempty[0]))
				*(*int32)(unsafe.Add(mBase, _c_F_brinbuildempty[0])) = v73 - int32(1)
				F_UnlockReleaseBuffer(m, v20)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					m.G0 = v6 + int32(32)
					return
				}
			}
		}
	}
}
func F_brininsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
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
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
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
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	v9 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l7)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v9
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)))
	v29 = v28
	goto L3
L2:
	;
	v29 = v9
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[0]))
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l7)+140))
	*(*int32)(unsafe.Add(mBase, _c_F_brininsert[0])) = v35
	v38 = F_palloc0(m, int32(12))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v53 = v23
	goto L6
L6:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v58 = v54 | v55<<(uint(int32(16))%32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v60 = base.I32_rem_u_s(v58, v59)
	v61 = v58 - v60
	v63 = v61 - int32(1)
	v64 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v78 = v64
	goto L11
L7:
	;
	return int32(0)
L8:
	;
	v42 = F_brin_build_desc(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v42
	v47 = F_brinRevmapInitialize(m, l0, v38+int32(8))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l7)+136)) = v38
	*(*int32)(unsafe.Add(mBase, _c_F_brininsert[0])) = v31
	v53 = v38
	goto L6
L11:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[1]))
	if v92 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	if v339 != 0 {
		goto L76
	} else {
		goto L77
	}
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v29&(base.B2i32(v60 == v64)&base.B2i32(v58 != v64)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v236 = F_brinGetTupleForHeapBlock(m, v71, v61, v21+int32(28), v21+int32(26), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L7
	} else {
		goto L46
	}
L18:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	if v97 != int32(1) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v105 = F_brinGetTupleForHeapBlock(m, v71, v63, v21+int32(28), v21+int32(26), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	if v105 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v110 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[2]))
	v116 = F_LWLockAcquire(m, v112+int32(2816), v110)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_LockBuffer(m, v209, int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L44
	}
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[3]))
	v121 = v119 + int32(36)
	v130 = v110
	goto L30
L25:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[2]))
	F_LWLockRelease(m, v182+int32(2816))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L7
	} else {
		goto L37
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = int32(0)
	v170 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v167)+4)) = uint16(v170)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v173
	v180 = v170
	goto L25
L27:
	;
	v167 = v146 + int32(20)
	goto L26
L28:
	;
	v167 = v146 + int32(40)
	goto L26
L29:
	;
	v167 = v146 + int32(60)
	goto L26
L30:
	;
	v141 = v130 * int32(20)
	v142 = v121 + v141
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+4)))
	if v143 == int32(0) {
		v167 = v142
		goto L26
	} else {
		goto L32
	}
L31:
	;
	v180 = int32(0)
	goto L25
L32:
	;
	v146 = v121 + v141
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+24)))
	if v147 != int32(1) {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+44)))
	if v150 != int32(1) {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+64)))
	if v153 != int32(1) {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v157 = v130 + int32(4)
	if v157 != int32(256) {
		v130 = v157
		goto L30
	} else {
		goto L36
	}
L36:
	;
	goto L31
L37:
	;
	if v180 != 0 {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	v189 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	if v189 == int32(0) {
		goto L17
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v196 + int32(4)
	F_errmsg(m, int32(_a_F_brininsert_0), v21)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_brininsert_1), int32(416), int32(_a_F_brininsert_2))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	goto L17
L44:
	;
	goto L17
L45:
	;
	goto L12
L46:
	;
	if v236 == int32(0) {
		v334 = v78
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if v78 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v242 = int32(_a_F_brininsert_3)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[0]))
	v249 = F_AllocSetContextCreateInternal(m, v244, int32(_a_F_brininsert_4), int32(0), int32(_a_F_brininsert_5), int32(_a_F_brininsert_6))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	v252 = v78
	goto L50
L50:
	;
	v254 = F_brin_deform_tuple(m, v70, v236, int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brininsert[0])) = v249
	v252 = v249
	goto L50
L52:
	;
	v256 = F_add_values_to_range(m, l0, v70, v254, l1, l2)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	if v256 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_LockBuffer(m, v258, int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L7
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v258 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v334 = v252
	goto L45
L58:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+26)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v281+v282<<(uint(int32(2))%32))+20))
	v288 = int32(base.Ui32(v286) >> (uint(int32(17)) % 32))
	v289 = int32(0)
	v291 = F_brin_copy_tuple(m, v236, v288, v289, v289)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L7
	} else {
		goto L62
	}
L59:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[5]))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v267+(v258^int32(-1))<<(uint(int32(2))%32))))
	v281 = v273
	goto L58
L60:
	;
	goto L61
L61:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[6]))
	v281 = v275 + v258<<(uint(int32(13))%32) + int32(-8192)
	goto L58
L62:
	;
	v295 = F_brin_form_tuple(m, v70, v61, v254, v21+int32(20))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if base.Ui32(v288) < base.Ui32(v298) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_LockBuffer(m, v323, int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L7
	} else {
		goto L72
	}
L65:
	;
	if v297 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v322 = int32(1)
	goto L67
L67:
	;
	goto L64
L68:
	;
	v318 = F_PageGetExactFreeSpace(m, v317)
	mBase = m.M
	v322 = base.B2i32(base.Ui32(v298-v288) <= base.Ui32(v318))
	goto L67
L69:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[5]))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v303+(v297^int32(-1))<<(uint(int32(2))%32))))
	v317 = v309
	goto L68
L70:
	;
	goto L71
L71:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[6]))
	v317 = v311 + v297<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+26)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v330 = F_brin_doupdate(m, l0, v59, v71, v61, v327, v328, v291, v288, v295, v329, v322)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	if v330 != 0 {
		v334 = v252
		goto L45
	} else {
		goto L74
	}
L74:
	;
	F_MemoryContextReset(m, v252)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v78 = v252
	goto L11
L76:
	;
	F_ReleaseBuffer(m, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L7
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brininsert[0])) = v31
	if v334 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	F_MemoryContextDelete(m, v334)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L7
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	m.G0 = v21 + int32(32)
	return int32(0)
L83:
	;
	goto L82
}
func F_brinsummarize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 float64
	_ = v257
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
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
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 float64
	_ = v363
	var v367 float64
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	v7 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v26 = F_brinRevmapInitialize(m, l0, v22+int32(12))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l2 == int32(-1) {
		v43 = v7
		v44 = v29
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L99
	}
L5:
	;
	m.G0 = v22 + int32(32)
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(0)
	if base.Ui32(v44) <= base.Ui32(v43) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v34 = base.I32_rem_u_s(l2, v33)
	v35 = l2 - v34
	v36 = v33 + v35
	if base.Ui32(v29) < base.Ui32(v36) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v38 = v29
	goto L10
L9:
	;
	v38 = v36
	goto L10
L10:
	;
	if base.Ui32(v35) <= base.Ui32(v38) {
		v43 = v35
		v44 = v38
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_brinRevmapTerminate(m, v26)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	F_brinRevmapTerminate(m, v26)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v53 = int32(0)
	v58 = v43
	v64 = v7
	goto L17
L16:
	;
	goto L5
L17:
	;
	if l3 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v416 != 0 {
		goto L91
	} else {
		goto L92
	}
L19:
	;
	goto L18
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if base.Ui32(v44) < base.Ui32(v72+v58) {
		v399 = v53
		v410 = v64
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[0]))
	if v76 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v84 = F_brinGetTupleForHeapBlock(m, v26, v58, v22+int32(8), v22+int32(6), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v395 = v394 + v58
	if base.Ui32(v395) < base.Ui32(v44) {
		v53 = v377
		v58 = v395
		v64 = v388
		goto L17
	} else {
		goto L90
	}
L29:
	;
	if v84 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v53 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if l5 != 0 {
		goto L86
	} else {
		goto L87
	}
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v92 = F_palloc(m, int32(80))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v129 = v53
	v132 = v64
	goto L35
L35:
	;
	v133 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v133
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v129)+44))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v146 = base.I32_div_s(v140<<(uint(int32(1))%32)+int32(7), int32(8))
	v148 = v146 + int32(12)
	v150 = v148 & int32(-8)
	v151 = F_palloc0(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L40
	}
L36:
	;
	v94 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v94
	v99 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v92)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v92)+28)) = v90
	v105 = F_brin_build_desc(m, l0)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+44)) = v105
	v108 = F_brin_new_memtuple(m, v105)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+72)) = int32(0)
	v112 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v92)+64)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v92)+48)) = v108
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v92)+52)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v92)+60)) = v116
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	v122 = base.I32_rem_u_s(int32(-2), v90)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+36)) = v120 - v122 - int32(2)
	v127 = F_BuildIndexInfo(m, l0)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v129 = v92
	v132 = v127
	goto L35
L40:
	;
	v156 = v148&int32(24) | int32(-32)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)) = uint8(v156)
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v58
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if int32(0) < v160 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v174 = int32(128)
	v176 = v151 + int32(4)
	v180 = v156
	v181 = v133
	goto L44
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(24)))) = v150
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v129)+28))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v129)+40))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v231 = F_brin_doinsert(m, v225, v226, v227, v22+int32(28), v58, v151, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	if v174 != int32(128) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L43
L46:
	;
	v195 = v176
	v196 = v180
	v197 = v174 << (uint(int32(1)) % 32)
	goto L48
L47:
	;
	v189 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)) = uint8(v189)
	v192 = int32(1)
	v195 = v176 + v192
	v196 = v189
	v197 = v192
	goto L48
L48:
	;
	v198 = v197 | v196
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v198)
	v201 = v181 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	if v201 < v203 {
		v174 = v197
		v176 = v195
		v180 = v198
		v181 = v201
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)) = uint16(v231)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v129)+28))
	if base.Ui32(v234+v58) <= base.Ui32(v44) {
		v247 = v234
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+32)) = v58
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v250 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+140))
	v257 = m.T0[v256].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l1, v249, v132, v250, int32(1), v250, v58, v247, int32(14), v129, v250)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L56
	}
L52:
	;
	v238 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v129)+28))
	if base.Ui32(v241) <= base.Ui32(v238-v58) {
		v247 = v241
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v244 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v247 = v244 - v58
	goto L51
L56:
	;
	v267 = v151
	goto L57
L57:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[0]))
	if v279 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_ReleaseBuffer(m, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L83
	}
L59:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v129)+44))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	v286 = F_brin_form_tuple(m, v282, v58, v283, v22+int32(16))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if base.Ui32(v289) < base.Ui32(v290) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v129)+28))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v129)+40))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v322 = F_brin_doupdate(m, v315, v316, v317, v58, v318, v319, v267, v320, v286, v321, v314)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L72
	}
L65:
	;
	if v288 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v314 = int32(1)
	goto L67
L67:
	;
	goto L64
L68:
	;
	v310 = F_PageGetExactFreeSpace(m, v309)
	mBase = m.M
	v314 = base.B2i32(base.Ui32(v290-v289) <= base.Ui32(v310))
	goto L67
L69:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[2]))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v295+(v288^int32(-1))<<(uint(int32(2))%32))))
	v309 = v301
	goto L68
L70:
	;
	goto L71
L71:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[3]))
	v309 = v303 + v288<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	F_pfree(m, v267)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_pfree(m, v286)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v322 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v129)+40))
	v337 = F_brinGetTupleForHeapBlock(m, v330, v58, v22+int32(28), v22+int32(22), v22+int32(24))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L58
L78:
	;
	if v337 == int32(0) {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v342 = int32(0)
	v344 = F_brin_copy_tuple(m, v337, v341, v342, v342)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_LockBuffer(m, v346, int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v129)+44))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	F_union_tuples(m, v350, v351, v344)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v267 = v344
	goto L57
L83:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v129)+44))
	F_brin_memtuple_initialize(m, v357, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if l4 == int32(0) {
		v377 = v129
		v388 = v132
		goto L28
	} else {
		goto L85
	}
L85:
	;
	v363 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(v363, float64(1))
	v377 = v129
	v388 = v132
	goto L28
L86:
	;
	v367 = *(*float64)(unsafe.Add(mBase, uint32(l5)))
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = base.F64_add(v367, float64(1))
	goto L88
L87:
	;
	goto L88
L88:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	F_LockBuffer(m, v371, int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v377 = v53
	v388 = v64
	goto L28
L90:
	;
	v399 = v377
	v410 = v388
	goto L19
L91:
	;
	F_ReleaseBuffer(m, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	F_brinRevmapTerminate(m, v26)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	if v399 == int32(0) {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	F_terminate_brin_buildstate(m, v399)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_pfree(m, v410)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	goto L5
L99:
	;
	F_errmsg_internal(m, int32(_a_F_brinsummarize_0), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_brinsummarize_1), int32(1864), int32(_a_F_brinsummarize_2))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v49
L2:
	;
	v10 = l1
	v11 = l2
	goto L5
L3:
	;
	goto L4
L4:
	;
	v49 = int32(0)
	goto L1
L5:
	;
	v18 = int32(base.Ui32(v11) >> (uint(int32(1)) % 32))
	v20 = v10 + v18*l3
	v21 = m.T0[l4].(func(*base.Module, int32, int32) int32)(m, l0, v20)
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	if v34 != 0 {
		v10 = v33
		v11 = v34
		goto L5
	} else {
		goto L14
	}
L8:
	;
	return int32(0)
L9:
	;
	if v21 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = v10
	v34 = v18
	goto L7
L11:
	;
	goto L12
L12:
	;
	if v21 == int32(0) {
		v49 = v20
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v33 = l3 + v20
	v34 = v11 + (v18 ^ int32(-1))
	goto L7
L14:
	;
	goto L6
}
func F_btbpchar_pattern_cmp(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(1)
	v21 = v6 + v20
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v26 = v24 & v20
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v6 {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	v27 = v21
	goto L7
L6:
	;
	v27 = v6 + int32(4)
	goto L7
L7:
	;
	if v24 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v61 = v54
	goto L19
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v33 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v44 = int32(1)
	if v26 != 0 {
		v54 = int32(base.Ui32(v24)>>(uint(v44)%32)) - v44
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v36 = int32(16)
	goto L14
L13:
	;
	v36 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = int32(4)
	goto L17
L16:
	;
	v43 = v36
	goto L17
L17:
	;
	v54 = v43
	goto L8
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	if v61 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v77 = int32(1)
	v78 = v11 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v83 = v81 & v77
	if v83 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v76 = v54 & (v54 >> (uint(int32(31)) % 32))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v70 = v61 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v70))))
	if v72 == int32(32) {
		v61 = v70
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v76 = v61
	goto L21
L26:
	;
	v84 = v78
	goto L28
L27:
	;
	v84 = v11 + int32(4)
	goto L28
L28:
	;
	if v81 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v118 = v111
	goto L40
L30:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v90 == int32(18) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v101 = int32(1)
	if v83 != 0 {
		v111 = int32(base.Ui32(v81)>>(uint(v101)%32)) - v101
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v93 = int32(16)
	goto L35
L34:
	;
	v93 = int32(0)
	goto L35
L35:
	;
	if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v100 = int32(4)
	goto L38
L37:
	;
	v100 = v93
	goto L38
L38:
	;
	v111 = v100
	goto L29
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	if v118 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v134 = int32(1)
	if v24&v134 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v132 = v111 & (v111 >> (uint(int32(31)) % 32))
	goto L42
L44:
	;
	goto L45
L45:
	;
	v127 = v118 - int32(1)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v127))))
	if v129 == int32(32) {
		v118 = v127
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v132 = v118
	goto L42
L47:
	;
	goto L4
L48:
	;
	v138 = v134
	goto L50
L49:
	;
	v138 = int32(4)
	goto L50
L50:
	;
	v140 = int32(1)
	if v81&v140 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v144 = v140
	goto L53
L52:
	;
	v144 = int32(4)
	goto L53
L53:
	;
	v146 = base.B2i32(v76 < v132)
	if v76 < v132 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v147 = v76
	goto L56
L55:
	;
	v147 = v132
	goto L56
L56:
	;
	v148 = F_memcmp(m, v6+v138, v11+v144, v147)
	mBase = m.M
	if v148 != 0 {
		v151 = v148
		goto L47
	} else {
		goto L57
	}
L57:
	;
	if v76 < v132 {
		v151 = int32(-1)
		goto L47
	} else {
		goto L58
	}
L58:
	;
	v151 = base.B2i32(v132 < v76)
	goto L47
L59:
	;
	F_pfree(m, v6)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v11 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v11)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	return v151
L66:
	;
	goto L65
}
func F_btbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
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
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
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
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
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
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
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
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v256 int64
	_ = v256
	var v257 int64
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v271 int64
	_ = v271
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
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
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
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
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 float64
	_ = v494
	var v495 int32
	_ = v495
	var v496 float64
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 float64
	_ = v540
	var v542 int32
	_ = v542
	var v544 float64
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v567 float64
	_ = v567
	var v568 float64
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int64
	_ = v574
	var v576 int64
	_ = v576
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v729 int64
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v981 int32
	_ = v981
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1064 int64
	_ = v1064
	var v1072 int32
	_ = v1072
	var v1083 int32
	_ = v1083
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1141 int32
	_ = v1141
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1204 int32
	_ = v1204
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1230 int64
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1285 int64
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1325 int32
	_ = v1325
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int64
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1368 int64
	_ = v1368
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1386 int64
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1467 int32
	_ = v1467
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1561 int32
	_ = v1561
	var v1567 int32
	_ = v1567
	var v1577 int32
	_ = v1577
	var v1584 int32
	_ = v1584
	var v1587 int64
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1594 int32
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1663 int64
	_ = v1663
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1681 int64
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1723 int64
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1763 int32
	_ = v1763
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1784 int32
	_ = v1784
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1797 int32
	_ = v1797
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1891 int32
	_ = v1891
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1981 float64
	_ = v1981
	v4 = int32(0)
	v17 = int64(0)
	v20 = m.G0
	v22 = v20 - int32(96)
	m.G0 = v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+16)) = uint8(v24)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+117)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)) = uint8(v4)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)) = uint8(v26)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v4
	v38 = F_RelationGetNumberOfBlocksInFork(m, l1, v4)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1932 = F_smgr_bulk_get_buf(m, v1931)
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L4
	} else {
		goto L355
	}
L2:
	;
	v1813 = int32(0)
	v1816 = v1813
	v1817 = v1813
	v1818 = v1797
	goto L336
L3:
	;
	F_pfree(m, v1332)
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L4
	} else {
		goto L335
	}
L4:
	;
	return int32(0)
L5:
	;
	if v38 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v45 = F_palloc0(m, int32(16))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L4
	} else {
		goto L332
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = l0
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)) = uint8(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+117)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+13)) = uint8(v51)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v45
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v58 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v99 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v62&int32(1) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v67 = int32(_a_F_btbuild_0)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v70 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v69 + v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v73 + v70
	v77 = int32(0)
	v79 = int32(_a_F_btbuild_1)
	v80 = base.AtomicRmwOr32(m, v77, v79, v77)
	*(*int64)(unsafe.Add(mBase, uint32(v58+int32(80))+232)) = int64(2)
	v88 = base.AtomicRmwOr32(m, v77, v79, v77)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v89 + v70
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v95 - v70
	goto L11
L14:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v423 != 0 {
		goto L97
	} else {
		goto L98
	}
L15:
	;
	v103 = v99 + int32(1)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	v106 = F_palloc0(m, int32(32))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[3]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+72)) = v111 + int32(1)
	goto L17
L17:
	;
	v118 = F_CreateParallelContext(m, int32(_a_F_btbuild_2), int32(_a_F_btbuild_3), v99)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v104 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v122 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v126 = int32(_a_F_btbuild_4)
	goto L21
L21:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v129 = F_table_parallelscan_estimate(m, v128, v126)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v124 = F_RegisterSnapshot(m, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v126 = v124
	goto L21
L24:
	;
	v131 = F_add_size(m, int32(96), v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
	v138 = F_add_size(m, v133, (v131+int32(31))&int32(-32))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+36)) = v138
	v141 = F_tuplesort_estimate_shared(m, v103)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
	v147 = (v141 + int32(31)) & int32(-32)
	v148 = F_add_size(m, v143, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+36)) = v148
	v151 = int32(1)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	if v153 == v151 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v156 = F_add_size(m, v148, v147)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v160 = int32(2)
	goto L31
L31:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v118)+40))
	v162 = F_add_size(m, v161, v160)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+36)) = v156
	v160 = int32(3)
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+40)) = v162
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v168 = F_mul_size(m, int32(32), v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v174 = F_add_size(m, v165, (v168+int32(31))&int32(-32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+36)) = v174
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v118)+40))
	v179 = F_add_size(m, v177, int32(1))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+40)) = v179
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v185 = F_mul_size(m, int32(128), v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v191 = F_add_size(m, v182, (v185+int32(31))&int32(-32))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+36)) = v191
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v118)+40))
	v196 = F_add_size(m, v194, int32(1))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+40)) = v196
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	if v200 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
	v202 = F_strlen(m, v200)
	mBase = m.M
	v207 = F_add_size(m, v201, v202&int32(-32)+int32(32))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	v217 = v151
	goto L42
L42:
	;
	F_InitializeParallelDSM(m, v118)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+36)) = v207
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v118)+40))
	v212 = F_add_size(m, v210, int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+40)) = v212
	v217 = v202 + int32(1)
	goto L42
L45:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v118)+44))
	if v220 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	switch v223 {
	case 0, 5:
		goto L50
	default:
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	v236 = F_shm_toc_allocate(m, v235, v131)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L54
	}
L49:
	;
	F_DestroyParallelContext(m, v118)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	F_UnregisterSnapshot(m, v126)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[3]))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+72)) = v231 - int32(1)
	goto L53
L53:
	;
	goto L14
L54:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = v242
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+8)) = uint8(v244)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v103
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+10)) = uint8(v104)
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+9)) = uint8(v246)
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v252 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v257
	v260 = v236 + int32(24)
	v261 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v260))), uint32(v261))
	*(*int64)(unsafe.Add(mBase, uint32(v260)+4)) = int64(-1)
	goto L59
L56:
	;
	v257 = int64(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v252)+392))
	v257 = v256
	goto L55
L59:
	;
	v266 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v236)+36)), uint32(v266))
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+72)) = uint8(v266)
	v271 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = v271
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+56)) = uint8(v266)
	*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v236)+40)) = v266
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	F_table_parallelscan_initialize(m, v279, v236+int32(96), v126)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	v285 = F_shm_toc_allocate(m, v284, v141)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v118)+44))
	F_tuplesort_initialize_shared(m, v285, v103, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	F_shm_toc_insert(m, v290, int64(-6917529027641081855), v236)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	F_shm_toc_insert(m, v294, int64(-6917529027641081854), v285)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	if v299 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	v303 = F_shm_toc_allocate(m, v302, v141)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	v312 = int32(0)
	goto L67
L67:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	if v314 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v118)+44))
	F_tuplesort_initialize_shared(m, v303, v103, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	F_shm_toc_insert(m, v308, int64(-6917529027641081853), v303)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v312 = v303
	goto L67
L71:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	v316 = F_shm_toc_allocate(m, v315, v217)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v329 = F_mul_size(m, int32(32), v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L79
	}
L74:
	;
	if v217 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	base.MemoryCopy(m, v316, v319, v217)
	goto L77
L76:
	;
	goto L77
L77:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	F_shm_toc_insert(m, v321, int64(-6917529027641081852), v316)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	goto L73
L79:
	;
	v331 = F_shm_toc_allocate(m, v326, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	F_shm_toc_insert(m, v333, int64(-6917529027641081851), v331)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v340 = F_mul_size(m, int32(128), v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v342 = F_shm_toc_allocate(m, v337, v340)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
	F_shm_toc_insert(m, v344, int64(-6917529027641081850), v342)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_LaunchParallelWorkers(m, v118)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v118
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v106)+8)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v351 + int32(1)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	if v361 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F__bt_end_parallel(m, v106)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v106
	v368 = F_palloc0(m, int32(16))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L90
	}
L89:
	;
	goto L14
L90:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+4)) = v371
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+8)) = v374
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v368)+12)) = uint8(v377)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v368)+13)) = uint8(v379)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+8)))
	if v383 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v387 = F_palloc0(m, int32(16))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L94
	}
L92:
	;
	v396 = int32(0)
	v398 = v382
	goto L93
L93:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[5]))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v404 = base.I32_div_s(v402, v403)
	F__bt_parallel_scan_and_sort(m, v368, v396, v398, v399, v400, v404, int32(1))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L95
	}
L94:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v389
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	v392 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v387)+12)) = uint8(v392)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+8)) = v391
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	v396 = v387
	v398 = v395
	goto L93
L95:
	;
	F_WaitForParallelWorkersToAttach(m, v118)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	goto L14
L97:
	;
	v425 = F_palloc0(m, int32(12))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L100
	}
L98:
	;
	v435 = int32(0)
	goto L99
L99:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+16)))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)))
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[5]))
	v440 = F_tuplesort_begin_index_btree(m, l0, l1, v436, v437, v439, v435)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L101
	}
L100:
	;
	v427 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v427)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v425)+4)) = v430
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v425)+8)) = v433
	v435 = v425
	goto L99
L101:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v440
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	if v444 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v448 = F_palloc0(m, int32(16))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v480 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L105:
	;
	v450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v448)+12)) = uint8(v450)
	*(*int32)(unsafe.Add(mBase, uint32(v448)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v448
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v456 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v458 = F_palloc0(m, int32(12))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L109
	}
L107:
	;
	v469 = v450
	v470 = v448
	goto L108
L108:
	;
	v471 = int32(0)
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[6]))
	v475 = F_tuplesort_begin_index_btree(m, l0, l1, v471, v471, v474, v469)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L110
	}
L109:
	;
	v460 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v460)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v458)+4)) = v463
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v458)+8)) = v466
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v469 = v458
	v470 = v468
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v475
	goto L104
L111:
	;
	v569 = int32(0)
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v571
	v574 = *(*int64)(unsafe.Add(mBase, _c_F_btbuild[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v574
	v576 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v576
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = base.I64_trunc_sat_f64_s(v568)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v576
	goto L129
L112:
	;
	v483 = int32(1)
	v484 = int32(0)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+140))
	v494 = m.T0[v493].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v483, v484, v483, v484, int32(-1), int32(238), v22+int32(16), v484)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v480)+8))
	v501 = v497 + int32(36)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	goto L116
L115:
	;
	v496 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v567 = v494
	v568 = v496
	goto L111
L116:
	;
	v524 = base.AtomicRmwXchg32(m, v501, int32(0), int32(1))
	if v524 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)) = uint8(v538)
	v540 = *(*float64)(unsafe.Add(mBase, uint32(v497)+64))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+32)) = v540
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+72)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)) = uint8(v542)
	v544 = *(*float64)(unsafe.Add(mBase, uint32(v497)+48))
	v545 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v497)+36)), uint32(v545))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L126
	}
L118:
	;
	F_s_lock(m, v501, int32(_a_F_btbuild_5), int32(1664), int32(_a_F_btbuild_6))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v497)+40))
	if v502 != v530 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	v532 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v501))), uint32(v532))
	F_ConditionVariableSleep(m, v497+int32(24), int32(134217767))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	goto L117
L125:
	;
	goto L116
L126:
	;
	v567 = v544
	v568 = v540
	goto L111
L127:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v768 == int32(0) {
		v779 = v569
		goto L144
	} else {
		goto L145
	}
L128:
	;
	goto L127
L129:
	;
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v596 == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v600&int32(1) == int32(0) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v605 = int32(_a_F_btbuild_0)
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v608 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v607 + v608
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	*(*int32)(unsafe.Add(mBase, uint32(v596))) = v611 + v608
	v615 = int32(0)
	v618 = base.AtomicRmwOr32(m, v615, int32(_a_F_btbuild_1), v615)
	goto L133
L132:
	;
	v745 = int32(0)
	v748 = base.AtomicRmwOr32(m, v745, int32(_a_F_btbuild_1), v745)
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	v750 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v596))) = v749 + v750
	v753 = int32(_a_F_btbuild_0)
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v755 - v750
	goto L128
L133:
	;
	goto L135
L135:
	;
	goto L136
L136:
	;
	v710 = int32(0)
	v713 = v569
	goto L141
L141:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(80)+v713<<(uint(int32(2))%32))))
	v723 = int32(3)
	v729 = *(*int64)(unsafe.Add(mBase, uint32(v22+int32(48)+v713<<(uint(v723)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v596+int32(232)+v722<<(uint(v723)%32)))) = v729
	v731 = int32(1)
	v734 = v710 + v731
	if v734 != int32(3) {
		v710 = v734
		v713 = v713 + v731
		goto L141
	} else {
		goto L143
	}
L142:
	;
	goto L132
L143:
	;
	goto L142
L144:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v785 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L145:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)))
	if v771 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v779 = v768
	goto L144
L147:
	;
	goto L148
L148:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v768)))
	F_tuplesort_end(m, v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_pfree(m, v768)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = int32(0)
	v779 = v569
	goto L144
L151:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	F_tuplesort_performsort(m, v826)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L4
	} else {
		goto L155
	}
L152:
	;
	goto L151
L153:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v789&int32(1) == int32(0) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v794 = int32(_a_F_btbuild_0)
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v797 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v796 + v797
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	*(*int32)(unsafe.Add(mBase, uint32(v785))) = v800 + v797
	v804 = int32(0)
	v806 = int32(_a_F_btbuild_1)
	v807 = base.AtomicRmwOr32(m, v804, v806, v804)
	*(*int64)(unsafe.Add(mBase, uint32(v785+int32(80))+232)) = int64(3)
	v815 = base.AtomicRmwOr32(m, v804, v806, v804)
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	*(*int32)(unsafe.Add(mBase, uint32(v785))) = v816 + v797
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v822 - v797
	goto L152
L155:
	;
	if v779 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v833 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	goto L158
L158:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v780)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v877
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v780)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v879
	v881 = int32(0)
	v883 = F__bt_mkscankey(m, v879, v881)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L4
	} else {
		goto L164
	}
L159:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	F_tuplesort_performsort(m, v874)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L4
	} else {
		goto L163
	}
L160:
	;
	goto L159
L161:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v837&int32(1) == int32(0) {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v842 = int32(_a_F_btbuild_0)
	v844 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v845 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v844 + v845
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v833)))
	*(*int32)(unsafe.Add(mBase, uint32(v833))) = v848 + v845
	v852 = int32(0)
	v854 = int32(_a_F_btbuild_1)
	v855 = base.AtomicRmwOr32(m, v852, v854, v852)
	*(*int64)(unsafe.Add(mBase, uint32(v833+int32(80))+232)) = int64(4)
	v863 = base.AtomicRmwOr32(m, v852, v854, v852)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v833)))
	*(*int32)(unsafe.Add(mBase, uint32(v833))) = v864 + v845
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v870 - v845
	goto L160
L163:
	;
	goto L158
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v883
	v887 = F__bt_allequalimage(m, v879, int32(1))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v883)+1)) = uint8(v887)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(1)
	v896 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v896 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v879)+52))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v879)+192))
	v939 = int32(*(*int16)(unsafe.Add(mBase, uint32(v938)+10)))
	v941 = F_smgr_bulk_start_rel(m, v879, int32(0))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L4
	} else {
		goto L170
	}
L167:
	;
	goto L166
L168:
	;
	v900 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v900&int32(1) == int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v905 = int32(_a_F_btbuild_0)
	v907 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v908 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v907 + v908
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v896)))
	*(*int32)(unsafe.Add(mBase, uint32(v896))) = v911 + v908
	v915 = int32(0)
	v917 = int32(_a_F_btbuild_1)
	v918 = base.AtomicRmwOr32(m, v915, v917, v915)
	*(*int64)(unsafe.Add(mBase, uint32(v896+int32(80))+232)) = int64(5)
	v926 = base.AtomicRmwOr32(m, v915, v917, v915)
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v896)))
	*(*int32)(unsafe.Add(mBase, uint32(v896))) = v927 + v908
	v933 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v933 - v908
	goto L167
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v941
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v883)+1)))
	if v944 != int32(1) {
		v953 = v881
		goto L176
	} else {
		goto L177
	}
L171:
	;
	F_pfree(m, v971)
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L4
	} else {
		goto L330
	}
L172:
	;
	v1650 = int32(0)
	v1652 = v956
	v1663 = v17
	goto L312
L173:
	;
	v1332 = F_palloc(m, int32(1676))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L4
	} else {
		goto L254
	}
L174:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v964 = F_tuplesort_getheaptuple(m, v963)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L4
	} else {
		goto L185
	}
L175:
	;
	if v779 == int32(0) {
		goto L173
	} else {
		goto L184
	}
L176:
	;
	if v779 != 0 {
		goto L174
	} else {
		goto L180
	}
L177:
	;
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780)+12)))
	if v947 != 0 {
		v953 = v881
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v879)+180))
	if v948 == int32(0) {
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+16)))
	v953 = v951
	goto L176
L180:
	;
	if v953 != 0 {
		goto L173
	} else {
		goto L181
	}
L181:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v956 = F_tuplesort_getheaptuple(m, v955)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	if v956 != 0 {
		goto L172
	} else {
		goto L183
	}
L183:
	;
	v1913 = int32(0)
	v1914 = int32(0)
	goto L1
L184:
	;
	goto L174
L185:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	v967 = F_tuplesort_getheaptuple(m, v966)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	v971 = F_palloc0(m, v939*int32(36))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	if int32(0) < v939 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v981 = int32(0)
	goto L191
L189:
	;
	goto L190
L190:
	;
	v1049 = v964
	v1051 = int32(0)
	v1058 = v967
	v1064 = v17
	goto L195
L191:
	;
	v999 = v971 + v981*int32(36)
	v1001 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v999))) = v1001
	v1005 = v883 + int32(16) + v981*int32(48)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1005)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v999)+4)) = v1006
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1005)))
	v1012 = int32(base.Ui32(v1008)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+9)) = uint8(v1012)
	v1014 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1005)+4)))
	v1015 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+20)) = uint8(v1015)
	*(*uint16)(unsafe.Add(mBase, uint32(v999)+10)) = uint16(v1014)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1005)))
	F_PrepareSortSupportFromIndexRel(m, v879, int32(base.Ui32(v1018&int32(16777216))>>(uint(int32(24))%32)), v999)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L4
	} else {
		goto L193
	}
L192:
	;
	goto L190
L193:
	;
	v1026 = v981 + int32(1)
	if v1026 != v939 {
		v981 = v1026
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	if v1058 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	if v1051 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L198:
	;
	if v1049 == int32(0) {
		goto L171
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v1072 = int32(0)
	if v1049 == v1072 {
		v1204 = v1072
		goto L197
	} else {
		goto L202
	}
L201:
	;
	v1204 = int32(1)
	goto L197
L202:
	;
	if int32(0) < v939 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1083 = int32(1)
	goto L206
L204:
	;
	goto L205
L205:
	;
	v1173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1049)+2)))
	v1174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1049))))
	v1175 = int32(16)
	v1177 = v1173 | v1174<<(uint(v1175)%32)
	v1178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058)+2)))
	v1179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058))))
	v1182 = v1178 | v1179<<(uint(v1175)%32)
	if base.Ui32(v1177) < base.Ui32(v1182) {
		v1193 = int32(-1)
		goto L229
	} else {
		goto L230
	}
L206:
	;
	v1099 = v971 + v1083*int32(36)
	v1102 = F_index_getattr_2(m, v1049, v1083, v937, v22+int32(80))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L4
	} else {
		goto L208
	}
L207:
	;
	goto L205
L208:
	;
	v1106 = F_index_getattr_2(m, v1058, v1083, v937, v22+int32(95))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L4
	} else {
		goto L209
	}
L209:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+95)))
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+80)))
	if v1109 == int32(1) {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	if v1083 != v939 {
		v1083 = v1083 + int32(1)
		goto L206
	} else {
		goto L227
	}
L211:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1099-int32(20))))
	v1130 = m.T0[v1129].(func(*base.Module, int32, int32, int32) int32)(m, v1102, v1106, v1099-int32(36))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L4
	} else {
		goto L220
	}
L212:
	;
	v1204 = int32(1)
	goto L197
L213:
	;
	if v1108&int32(1) != 0 {
		goto L210
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	if v1108&int32(1) == int32(0) {
		goto L211
	} else {
		goto L218
	}
L216:
	;
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099-int32(27)))))
	if v1116 != 0 {
		goto L212
	} else {
		goto L217
	}
L217:
	;
	v1204 = v1072
	goto L197
L218:
	;
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099-int32(27)))))
	if v1123 != 0 {
		v1204 = v1072
		goto L197
	} else {
		goto L219
	}
L219:
	;
	goto L212
L220:
	;
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099-int32(28)))))
	if v1134 == int32(1) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if v1130 < int32(0) {
		v1204 = v1072
		goto L197
	} else {
		goto L224
	}
L222:
	;
	v1141 = v1130
	goto L223
L223:
	;
	if int32(0) < v1141 {
		v1204 = v1072
		goto L197
	} else {
		goto L225
	}
L224:
	;
	v1141 = int32(0) - v1130
	goto L223
L225:
	;
	if v1141 == int32(0) {
		goto L210
	} else {
		goto L226
	}
L226:
	;
	v1204 = int32(1)
	goto L197
L227:
	;
	goto L207
L228:
	;
	v1204 = base.B2i32(v1193 <= int32(0))
	goto L197
L229:
	;
	goto L228
L230:
	;
	if base.Ui32(v1182) < base.Ui32(v1177) {
		v1193 = int32(1)
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v1187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1049)+4)))
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058)+4)))
	if base.Ui32(v1187) < base.Ui32(v1188) {
		v1193 = int32(-1)
		goto L229
	} else {
		goto L232
	}
L232:
	;
	v1193 = base.B2i32(base.Ui32(v1188) < base.Ui32(v1187))
	goto L229
L233:
	;
	v1218 = F_palloc0(m, int32(32))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L4
	} else {
		goto L236
	}
L234:
	;
	v1263 = v1051
	goto L235
L235:
	;
	if v1204 != 0 {
		goto L243
	} else {
		goto L244
	}
L236:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1221 = F_smgr_bulk_get_buf(m, v1220)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	F_PageInit(m, v1221, int32(_a_F_btbuild_7), int32(16))
	mBase = m.M
	goto L238
L238:
	;
	v1226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1221)+16)))
	v1227 = v1221 + v1226
	*(*int64)(unsafe.Add(mBase, uint32(v1227)+8)) = int64(4294967296)
	v1230 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1227))) = v1230
	v1232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1221)+12)))
	v1234 = v1232 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1221)+12)) = uint16(v1234)
	*(*int32)(unsafe.Add(mBase, uint32(v1218))) = v1221
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1238 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1237 + v1238
	*(*int64)(unsafe.Add(mBase, uint32(v1218)+16)) = v1230
	*(*uint16)(unsafe.Add(mBase, uint32(v1218)+12)) = uint16(v1238)
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+4)) = v1237
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1248)+180))
	if v1249 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1249)+4))
	v1256 = base.I32_div_s(int32(_a_F_btbuild_8)-v1251<<(uint(int32(13))%32), int32(100))
	v1258 = v1256
	goto L241
L240:
	;
	v1258 = int32(819)
	goto L241
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+24)) = v1258
	v1263 = v1218
	goto L235
L242:
	;
	v1285 = v1064 + int64(1)
	v1288 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1288 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L243:
	;
	F__bt_buildadd(m, v22+int32(48), v1263, v1049, int32(0))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L4
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	F__bt_buildadd(m, v22+int32(48), v1263, v1058, int32(0))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L4
	} else {
		goto L248
	}
L246:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v1271 = F_tuplesort_getheaptuple(m, v1270)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L4
	} else {
		goto L247
	}
L247:
	;
	v1281 = v1271
	v1282 = v1058
	goto L242
L248:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	v1279 = F_tuplesort_getheaptuple(m, v1278)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L4
	} else {
		goto L249
	}
L249:
	;
	v1281 = v1049
	v1282 = v1279
	goto L242
L250:
	;
	v1049 = v1281
	v1051 = v1263
	v1058 = v1282
	v1064 = v1285
	goto L195
L251:
	;
	goto L250
L252:
	;
	v1292 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1292&int32(1) == int32(0) {
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v1297 = int32(_a_F_btbuild_0)
	v1299 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1300 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1299 + v1300
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1288)))
	*(*int32)(unsafe.Add(mBase, uint32(v1288))) = v1303 + v1300
	v1307 = int32(0)
	v1309 = int32(_a_F_btbuild_1)
	v1310 = base.AtomicRmwOr32(m, v1307, v1309, v1307)
	*(*int64)(unsafe.Add(mBase, uint32(v1288+int32(96))+232)) = v1285
	v1318 = base.AtomicRmwOr32(m, v1307, v1309, v1307)
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1288)))
	*(*int32)(unsafe.Add(mBase, uint32(v1288))) = v1319 + v1300
	v1325 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1325 - v1300
	goto L251
L254:
	;
	v1334 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1332)+4)) = v1334
	v1336 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1332))) = uint8(v1336)
	*(*int64)(unsafe.Add(mBase, uint32(v1332)+10)) = v1334
	*(*int64)(unsafe.Add(mBase, uint32(v1332)+20)) = v1334
	*(*int64)(unsafe.Add(mBase, uint32(v1332)+28)) = v1334
	*(*int64)(unsafe.Add(mBase, uint32(v1332)+36)) = v1334
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v1347 = F_tuplesort_getheaptuple(m, v1346)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	if v1347 == int32(0) {
		goto L3
	} else {
		goto L256
	}
L256:
	;
	v1355 = int32(0)
	v1357 = v1347
	v1368 = v17
	goto L257
L257:
	;
	if v1355 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L258:
	;
	F__bt_sort_dedup_finish_pending(m, v22+int32(48), v1584, v1332)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L4
	} else {
		goto L308
	}
L259:
	;
	v1587 = v1368 + int64(1)
	v1590 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1590 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L260:
	;
	v1519 = F_CopyIndexTuple(m, v1357)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L4
	} else {
		goto L291
	}
L261:
	;
	v1374 = F_palloc0(m, int32(32))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L4
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+12))
	v1426 = F__bt_keep_natts_fast(m, v1424, v1425, v1357)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L4
	} else {
		goto L271
	}
L264:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1377 = F_smgr_bulk_get_buf(m, v1376)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	F_PageInit(m, v1377, int32(_a_F_btbuild_7), int32(16))
	mBase = m.M
	goto L266
L266:
	;
	v1382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1377)+16)))
	v1383 = v1377 + v1382
	*(*int64)(unsafe.Add(mBase, uint32(v1383)+8)) = int64(4294967296)
	v1386 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1383))) = v1386
	v1388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1377)+12)))
	v1390 = v1388 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1377)+12)) = uint16(v1390)
	*(*int32)(unsafe.Add(mBase, uint32(v1374))) = v1377
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1394 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1393 + v1394
	*(*int64)(unsafe.Add(mBase, uint32(v1374)+16)) = v1386
	*(*uint16)(unsafe.Add(mBase, uint32(v1374)+12)) = uint16(v1394)
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+4)) = v1393
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+180))
	if v1405 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+4))
	v1412 = base.I32_div_s(int32(_a_F_btbuild_8)-v1407<<(uint(int32(13))%32), int32(100))
	v1414 = v1412
	goto L269
L268:
	;
	v1414 = int32(819)
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+24)) = v1414
	v1418 = int32(812)
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+8)) = v1418
	v1421 = F_palloc(m, v1418)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+24)) = v1421
	v1518 = v1374
	goto L260
L271:
	;
	if v939 < v1426 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1434 = int32(1)
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1357)+7)))
	if v1435&int32(32) == int32(0) {
		v1453 = v1434
		v1455 = v1357
		goto L276
	} else {
		goto L277
	}
L273:
	;
	goto L274
L274:
	;
	F__bt_sort_dedup_finish_pending(m, v22+int32(48), v1355, v1332)
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L4
	} else {
		goto L289
	}
L275:
	;
	if base.Ui32((v1457+(v1458+v1453)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1456) {
		v1584 = v1355
		goto L259
	} else {
		goto L288
	}
L276:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+8))
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+20))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+28))
	v1467 = base.B2i32(base.Ui32((v1457+(v1458+v1453)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1456))
	if v1467 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L277:
	;
	v1440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1357)+4)))
	if v1440&int32(_a_F_btbuild_7) == int32(0) {
		v1453 = v1434
		v1455 = v1357
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1357)+2)))
	v1448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1357))))
	v1453 = v1440 & int32(4095)
	v1455 = v1447 + (v1357 + v1448<<(uint(int32(16))%32))
	goto L276
L279:
	;
	goto L275
L280:
	;
	v1501 = v1332 + v1498
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1501)))
	*(*int32)(unsafe.Add(mBase, uint32(v1501))) = v1502 + v1500
	goto L279
L281:
	;
	if v1458 <= int32(50) {
		goto L279
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+32)) = v1474 + int32(1)
	v1479 = v1453 * int32(6)
	if v1479 != 0 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	v1498 = int32(4)
	v1500 = int32(1)
	goto L280
L285:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+24))
	base.MemoryCopy(m, v1480+v1458*int32(6), v1455, v1479)
	goto L287
L286:
	;
	goto L287
L287:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+28)) = v1485 + v1453
	v1489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1357)+6)))
	v1498 = int32(36)
	v1500 = (v1489&int32(_a_F_btbuild_9)+int32(7))&int32(_a_F_btbuild_10) | int32(4)
	goto L280
L288:
	;
	goto L274
L289:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+12))
	F_pfree(m, v1513)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L4
	} else {
		goto L290
	}
L290:
	;
	v1518 = v1355
	goto L260
L291:
	;
	v1521 = int32(0)
	v1524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1519)+7)))
	if v1524&int32(32) != 0 {
		goto L295
	} else {
		goto L296
	}
L292:
	;
	v1584 = v1518
	goto L259
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+20)) = v1561
	*(*uint16)(unsafe.Add(mBase, uint32(v1332)+16)) = uint16(v1521)
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+12)) = v1519
	v1567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+36)) = (v1567&int32(_a_F_btbuild_9)+int32(7))&int32(_a_F_btbuild_10) | int32(4)
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1332+v1577<<(uint(int32(2))%32))+44)) = uint16(v1521)
	goto L292
L294:
	;
	v1542 = v1527 & int32(4095)
	v1544 = v1542 * int32(6)
	if v1544 != 0 {
		goto L299
	} else {
		goto L300
	}
L295:
	;
	v1527 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519)+4)))
	if v1527&int32(_a_F_btbuild_7) != 0 {
		goto L294
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+24))
	v1532 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1531)+4)) = uint16(v1532)
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1519)))
	*(*int32)(unsafe.Add(mBase, uint32(v1531))) = v1534
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+28)) = int32(1)
	v1538 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519)+6)))
	v1561 = v1538 & int32(_a_F_btbuild_9)
	goto L293
L298:
	;
	goto L297
L299:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+24))
	v1546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519)+2)))
	v1547 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519))))
	base.MemoryCopy(m, v1545, v1546+(v1519+v1547<<(uint(int32(16))%32)), v1544)
	goto L301
L300:
	;
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+28)) = v1542
	v1554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519)+2)))
	v1555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519))))
	v1561 = v1554 | v1555<<(uint(int32(16))%32)
	goto L293
L302:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v1632 = F_tuplesort_getheaptuple(m, v1631)
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L4
	} else {
		goto L306
	}
L303:
	;
	goto L302
L304:
	;
	v1594 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1594&int32(1) == int32(0) {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1599 = int32(_a_F_btbuild_0)
	v1601 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1602 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1601 + v1602
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1590)))
	*(*int32)(unsafe.Add(mBase, uint32(v1590))) = v1605 + v1602
	v1609 = int32(0)
	v1611 = int32(_a_F_btbuild_1)
	v1612 = base.AtomicRmwOr32(m, v1609, v1611, v1609)
	*(*int64)(unsafe.Add(mBase, uint32(v1590+int32(96))+232)) = v1587
	v1620 = base.AtomicRmwOr32(m, v1609, v1611, v1609)
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1590)))
	*(*int32)(unsafe.Add(mBase, uint32(v1590))) = v1621 + v1602
	v1627 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1627 - v1602
	goto L303
L306:
	;
	if v1632 != 0 {
		v1355 = v1584
		v1357 = v1632
		v1368 = v1587
		goto L257
	} else {
		goto L307
	}
L307:
	;
	goto L258
L308:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+12))
	F_pfree(m, v1638)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L4
	} else {
		goto L309
	}
L309:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+24))
	F_pfree(m, v1641)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L4
	} else {
		goto L310
	}
L310:
	;
	F_pfree(m, v1332)
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L4
	} else {
		goto L311
	}
L311:
	;
	v1797 = v1584
	goto L2
L312:
	;
	if v1650 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	v1797 = v1715
	goto L2
L314:
	;
	v1669 = F_palloc0(m, int32(32))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L4
	} else {
		goto L317
	}
L315:
	;
	v1715 = v1650
	goto L316
L316:
	;
	F__bt_buildadd(m, v22+int32(48), v1715, v1652, int32(0))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L4
	} else {
		goto L323
	}
L317:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1672 = F_smgr_bulk_get_buf(m, v1671)
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L4
	} else {
		goto L318
	}
L318:
	;
	F_PageInit(m, v1672, int32(_a_F_btbuild_7), int32(16))
	mBase = m.M
	goto L319
L319:
	;
	v1677 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1672)+16)))
	v1678 = v1672 + v1677
	*(*int64)(unsafe.Add(mBase, uint32(v1678)+8)) = int64(4294967296)
	v1681 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1678))) = v1681
	v1683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1672)+12)))
	v1685 = v1683 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1672)+12)) = uint16(v1685)
	*(*int32)(unsafe.Add(mBase, uint32(v1669))) = v1672
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1689 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1688 + v1689
	*(*int64)(unsafe.Add(mBase, uint32(v1669)+16)) = v1681
	*(*uint16)(unsafe.Add(mBase, uint32(v1669)+12)) = uint16(v1689)
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+4)) = v1688
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1699)+180))
	if v1700 != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1700)+4))
	v1707 = base.I32_div_s(int32(_a_F_btbuild_8)-v1702<<(uint(int32(13))%32), int32(100))
	v1709 = v1707
	goto L322
L321:
	;
	v1709 = int32(819)
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+24)) = v1709
	v1715 = v1669
	goto L316
L323:
	;
	v1723 = v1663 + int64(1)
	v1726 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1726 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v1768 = F_tuplesort_getheaptuple(m, v1767)
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L4
	} else {
		goto L328
	}
L325:
	;
	goto L324
L326:
	;
	v1730 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1730&int32(1) == int32(0) {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1735 = int32(_a_F_btbuild_0)
	v1737 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1738 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1737 + v1738
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1726)))
	*(*int32)(unsafe.Add(mBase, uint32(v1726))) = v1741 + v1738
	v1745 = int32(0)
	v1747 = int32(_a_F_btbuild_1)
	v1748 = base.AtomicRmwOr32(m, v1745, v1747, v1745)
	*(*int64)(unsafe.Add(mBase, uint32(v1726+int32(96))+232)) = v1723
	v1756 = base.AtomicRmwOr32(m, v1745, v1747, v1745)
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1726)))
	*(*int32)(unsafe.Add(mBase, uint32(v1726))) = v1757 + v1738
	v1763 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1763 - v1738
	goto L325
L328:
	;
	if v1768 != 0 {
		v1650 = v1715
		v1652 = v1768
		v1663 = v1723
		goto L312
	} else {
		goto L329
	}
L329:
	;
	goto L313
L330:
	;
	if v1051 != 0 {
		v1797 = v1051
		goto L2
	} else {
		goto L331
	}
L331:
	;
	v1772 = int32(0)
	v1913 = v1772
	v1914 = v1772
	goto L1
L332:
	;
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v1778 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btbuild_11), v22)
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L4
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(_a_F_btbuild_5), int32(321), int32(_a_F_btbuild_12))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L4
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	v1792 = int32(0)
	v1913 = v1792
	v1914 = v1792
	goto L1
L336:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+4))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+28))
	if v1835 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	v1913 = v1862
	v1914 = v1863
	goto L1
L338:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1865 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1864)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1865) {
		goto L345
	} else {
		goto L346
	}
L339:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1838)+16)))
	v1840 = v1838 + v1839
	v1841 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1840)+12)))
	v1843 = v1841 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1840)+12)) = uint16(v1843)
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+20))
	v1862 = v1834
	v1863 = v1845
	goto L338
L340:
	;
	goto L341
L341:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1846))) = base.I32_rotr(v1834, int32(16))
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+28))
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+8))
	F__bt_buildadd(m, v22+int32(48), v1852, v1853, int32(0))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L4
	} else {
		goto L342
	}
L342:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+8))
	F_pfree(m, v1857)
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L4
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+8)) = int32(0)
	v1862 = v1816
	v1863 = v1817
	goto L338
L344:
	;
	v1901 = v1865 - int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1864)+12)) = uint16(v1901)
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+4))
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	F_smgr_bulk_write(m, v1903, v1904, v1905, int32(1))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L4
	} else {
		goto L353
	}
L345:
	;
	v1873 = int32(base.Ui32(v1865+int32(_a_F_btbuild_13)) >> (uint(int32(2)) % 32))
	goto L347
L346:
	;
	v1873 = int32(0)
	goto L347
L347:
	;
	if base.Ui32(v1873&int32(_a_F_btbuild_14)) < base.Ui32(int32(2)) {
		goto L344
	} else {
		goto L348
	}
L348:
	;
	v1878 = int32(3)
	v1882 = (v1873 + int32(1)) & int32(_a_F_btbuild_14)
	if base.Ui32(v1882) <= base.Ui32(v1878) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1885 = v1878
	goto L351
L350:
	;
	v1885 = v1882
	goto L351
L351:
	;
	v1886 = int32(2)
	v1891 = (v1885 - v1886) & int32(_a_F_btbuild_14) << (uint(v1886) % 32)
	if v1891 == int32(0) {
		goto L344
	} else {
		goto L352
	}
L352:
	;
	base.MemoryCopy(m, v1864+int32(24), v1864+int32(28), v1891)
	goto L344
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1818))) = int32(0)
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+28))
	if v1911 != 0 {
		v1816 = v1862
		v1817 = v1863
		v1818 = v1911
		goto L336
	} else {
		goto L354
	}
L354:
	;
	goto L337
L355:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	v1935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1934)+1)))
	F_PageInit(m, v1932, int32(_a_F_btbuild_7), int32(16))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1932)+64)) = uint8(v1935)
	*(*int64)(unsafe.Add(mBase, uint32(v1932)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v1932)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1932)+44)) = v1914
	*(*int32)(unsafe.Add(mBase, uint32(v1932)+40)) = v1913
	*(*int32)(unsafe.Add(mBase, uint32(v1932)+36)) = v1914
	*(*int32)(unsafe.Add(mBase, uint32(v1932)+32)) = v1913
	*(*int64)(unsafe.Add(mBase, uint32(v1932)+24)) = int64(17180209506)
	v1950 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1932)+16)))
	v1952 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v1932+v1950)+12)) = uint16(v1952)
	v1954 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v1932)+12)) = uint16(v1954)
	goto L356
L356:
	;
	F_smgr_bulk_write(m, v1931, int32(0), v1932, int32(1))
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L4
	} else {
		goto L357
	}
L357:
	;
	F_smgr_bulk_finish(m, v1931)
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L4
	} else {
		goto L358
	}
L358:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v1962)))
	F_tuplesort_end(m, v1963)
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L4
	} else {
		goto L359
	}
L359:
	;
	F_pfree(m, v1962)
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L4
	} else {
		goto L360
	}
L360:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v1968 != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1968)))
	F_tuplesort_end(m, v1969)
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L4
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v1974 != 0 {
		goto L366
	} else {
		goto L367
	}
L364:
	;
	F_pfree(m, v1968)
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L4
	} else {
		goto L365
	}
L365:
	;
	goto L363
L366:
	;
	F__bt_end_parallel(m, v1974)
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L4
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1978 = F_palloc(m, int32(16))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L4
	} else {
		goto L370
	}
L369:
	;
	goto L368
L370:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1978))) = v567
	v1981 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1978)+8)) = v1981
	m.G0 = v22 + int32(96)
	return v1978
}
func F_btbuildphasename(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = l0 - int64(1)
	if base.Ui64(v3) <= base.Ui64(int64(4)) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v3)<<(uint(int32(2))%32))+uint32(_c_F_btbuildphasename[0])))
		v11 = v9
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_btcharskipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(206)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(207)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(1095216660480)
	return int32(0)
}
func F_btfloat4fastcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 float32
	_ = v8
	var v9 float32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	v6 = int32(2147483647)
	v7 = l0 & v6
	v8 = base.F32_reinterpret_i32(l1)
	v9 = base.F32_reinterpret_i32(l0)
	v11 = l1 & v6
	if base.Ui32(int32(2139095041)) <= base.Ui32(v11) {
		v21 = base.B2i32(base.Ui32(v7) < base.Ui32(int32(2139095041)))
		v30 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v11))|base.F32_gt(v8, v9))&v21
	} else {
		v16 = int32(1)
		if base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v7))|base.F32_lt(v8, v9) != 0 {
			v30 = v16
		} else {
			v21 = v16
			v30 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v11))|base.F32_gt(v8, v9))&v21
		}
	}
	return v30
}
func F_btfloat84cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 float32
	_ = v11
	var v12 float64
	_ = v12
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v7) & v9
	v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = base.F64_promote_f32(v11)
	v15 = base.I64_reinterpret_f64(v12) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v15) {
		v25 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v34 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v25
	} else {
		v20 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v10))|base.F64_gt(v7, v12) != 0 {
			v34 = v20
		} else {
			v25 = v20
			v34 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v25
		}
	}
	return v34
}
func F_bthandler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v75 int32
	_ = v75
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+22)) = int32(16843009)
		v9 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+21)) = uint8(v9)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+13)) = int64(72340172838076673)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+12)) = uint8(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(_a_F_bthandler_0)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(1688871335100854)
		v19 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+26)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = int32(212)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+132)) = int32(213)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+128)) = int32(214)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+124)) = int32(215)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+120)) = int32(216)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+116)) = int32(217)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+112)) = int32(218)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(219)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(220)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = int32(221)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(222)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = int32(223)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = int32(224)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = int32(225)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+80)) = int32(226)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+76)) = int32(227)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(228)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = int32(229)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(230)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = int32(231)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(232)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(233)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v9
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(234)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(235)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(236)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = v9
		v75 = int32(3)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+29)) = uint8(v75)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+27)) = uint16(v9)
		return v3
	}
}
func F_btinitparallelscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	v3 = l0 + int32(12)
	v4 = int32(70)
	*(*uint16)(unsafe.Add(mBase, uint32(v3))) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = int64(-1)
	v10 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10
	v12 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v12
	v15 = l0 + int32(28)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v15))), uint32(v10))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = v12
	return
}
func F_btint2sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(194)
	return int32(0)
}
func F_btnamesortsupport(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13856(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_btoidfastcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return base.B2i32(base.Ui32(l1) < base.Ui32(l0)) - base.B2i32(base.Ui32(l0) < base.Ui32(l1))
}
func F_btrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	if v7 != int32(-1) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
		if int32(0) < v10 {
			F__bt_killitems(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
				if v15 != 0 {
					F_ReleaseBuffer(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
						v21 = int32(0)
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v22 != 0 {
							v39 = v21
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
							switch v24 {
							case 0, 5:
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
								v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
								if v27 != int32(112) {
									v39 = v21
								} else {
									v31 = *(*int32)(unsafe.Add(mBase, _c_F_btrescan[0]))
									if v31 <= int32(0) {
										v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
										if v34 != 0 {
											v39 = v21
										} else {
											v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
											if v35 != 0 {
												v39 = v21
											} else {
												v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v39 = base.B2i32(v36 != int32(0))
											}
										}
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v39 = base.B2i32(v36 != int32(0))
									}
								}
							default:
								v39 = v21
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v39)
						v44 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v44)
						*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v44)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1])))
						if v48 != 0 {
							F_ReleaseBuffer(m, v48)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
								if v53 != int32(1) {
									if l1 == int32(0) {
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v67 <= int32(0) {
										} else {
											v71 = v67 * int32(48)
											if v71 == int32(0) {
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v74, l1, v71)
											}
										}
									}
									v77 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
									return
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
									if v56 != 0 {
										if l1 == int32(0) {
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v67 <= int32(0) {
											} else {
												v71 = v67 * int32(48)
												if v71 == int32(0) {
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													base.MemoryCopy(m, v74, l1, v71)
												}
											}
										}
										v77 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
										*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
										return
									} else {
										v58 = F_palloc(m, int32(_a_F_btrescan_0))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
											*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
											if l1 == int32(0) {
											} else {
												v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v67 <= int32(0) {
												} else {
													v71 = v67 * int32(48)
													if v71 == int32(0) {
													} else {
														v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														base.MemoryCopy(m, v74, l1, v71)
													}
												}
											}
											v77 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
											*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
											return
										}
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v53 != int32(1) {
								if l1 == int32(0) {
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v67 <= int32(0) {
									} else {
										v71 = v67 * int32(48)
										if v71 == int32(0) {
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v74, l1, v71)
										}
									}
								}
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
								return
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v56 != 0 {
									if l1 == int32(0) {
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v67 <= int32(0) {
										} else {
											v71 = v67 * int32(48)
											if v71 == int32(0) {
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v74, l1, v71)
											}
										}
									}
									v77 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
									return
								} else {
									v58 = F_palloc(m, int32(_a_F_btrescan_0))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
										*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
										if l1 == int32(0) {
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v67 <= int32(0) {
											} else {
												v71 = v67 * int32(48)
												if v71 == int32(0) {
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													base.MemoryCopy(m, v74, l1, v71)
												}
											}
										}
										v77 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
										*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
										return
									}
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
					v21 = int32(0)
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v22 != 0 {
						v39 = v21
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						switch v24 {
						case 0, 5:
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
							v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
							if v27 != int32(112) {
								v39 = v21
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, _c_F_btrescan[0]))
								if v31 <= int32(0) {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
									if v34 != 0 {
										v39 = v21
									} else {
										v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
										if v35 != 0 {
											v39 = v21
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v39 = base.B2i32(v36 != int32(0))
										}
									}
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v39 = base.B2i32(v36 != int32(0))
								}
							}
						default:
							v39 = v21
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v39)
					v44 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v44)
					*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v44)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1])))
					if v48 != 0 {
						F_ReleaseBuffer(m, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v53 != int32(1) {
								if l1 == int32(0) {
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v67 <= int32(0) {
									} else {
										v71 = v67 * int32(48)
										if v71 == int32(0) {
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v74, l1, v71)
										}
									}
								}
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
								return
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v56 != 0 {
									if l1 == int32(0) {
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v67 <= int32(0) {
										} else {
											v71 = v67 * int32(48)
											if v71 == int32(0) {
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v74, l1, v71)
											}
										}
									}
									v77 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
									return
								} else {
									v58 = F_palloc(m, int32(_a_F_btrescan_0))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
										*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
										if l1 == int32(0) {
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v67 <= int32(0) {
											} else {
												v71 = v67 * int32(48)
												if v71 == int32(0) {
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													base.MemoryCopy(m, v74, l1, v71)
												}
											}
										}
										v77 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
										*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
										return
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v53 != int32(1) {
							if l1 == int32(0) {
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v67 <= int32(0) {
								} else {
									v71 = v67 * int32(48)
									if v71 == int32(0) {
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v74, l1, v71)
									}
								}
							}
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
							return
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							if v56 != 0 {
								if l1 == int32(0) {
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v67 <= int32(0) {
									} else {
										v71 = v67 * int32(48)
										if v71 == int32(0) {
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v74, l1, v71)
										}
									}
								}
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
								return
							} else {
								v58 = F_palloc(m, int32(_a_F_btrescan_0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
									*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
									if l1 == int32(0) {
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v67 <= int32(0) {
										} else {
											v71 = v67 * int32(48)
											if v71 == int32(0) {
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v74, l1, v71)
											}
										}
									}
									v77 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
									return
								}
							}
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
			if v15 != 0 {
				F_ReleaseBuffer(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
					v21 = int32(0)
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v22 != 0 {
						v39 = v21
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						switch v24 {
						case 0, 5:
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
							v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
							if v27 != int32(112) {
								v39 = v21
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, _c_F_btrescan[0]))
								if v31 <= int32(0) {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
									if v34 != 0 {
										v39 = v21
									} else {
										v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
										if v35 != 0 {
											v39 = v21
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v39 = base.B2i32(v36 != int32(0))
										}
									}
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v39 = base.B2i32(v36 != int32(0))
								}
							}
						default:
							v39 = v21
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v39)
					v44 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v44)
					*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v44)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1])))
					if v48 != 0 {
						F_ReleaseBuffer(m, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v53 != int32(1) {
								if l1 == int32(0) {
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v67 <= int32(0) {
									} else {
										v71 = v67 * int32(48)
										if v71 == int32(0) {
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v74, l1, v71)
										}
									}
								}
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
								return
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v56 != 0 {
									if l1 == int32(0) {
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v67 <= int32(0) {
										} else {
											v71 = v67 * int32(48)
											if v71 == int32(0) {
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v74, l1, v71)
											}
										}
									}
									v77 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
									return
								} else {
									v58 = F_palloc(m, int32(_a_F_btrescan_0))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
										*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
										if l1 == int32(0) {
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v67 <= int32(0) {
											} else {
												v71 = v67 * int32(48)
												if v71 == int32(0) {
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													base.MemoryCopy(m, v74, l1, v71)
												}
											}
										}
										v77 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
										*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
										return
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v53 != int32(1) {
							if l1 == int32(0) {
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v67 <= int32(0) {
								} else {
									v71 = v67 * int32(48)
									if v71 == int32(0) {
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v74, l1, v71)
									}
								}
							}
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
							return
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							if v56 != 0 {
								if l1 == int32(0) {
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v67 <= int32(0) {
									} else {
										v71 = v67 * int32(48)
										if v71 == int32(0) {
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v74, l1, v71)
										}
									}
								}
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
								return
							} else {
								v58 = F_palloc(m, int32(_a_F_btrescan_0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
									*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
									if l1 == int32(0) {
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v67 <= int32(0) {
										} else {
											v71 = v67 * int32(48)
											if v71 == int32(0) {
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v74, l1, v71)
											}
										}
									}
									v77 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
									return
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
				v21 = int32(0)
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v22 != 0 {
					v39 = v21
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					switch v24 {
					case 0, 5:
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
						if v27 != int32(112) {
							v39 = v21
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, _c_F_btrescan[0]))
							if v31 <= int32(0) {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
								if v34 != 0 {
									v39 = v21
								} else {
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
									if v35 != 0 {
										v39 = v21
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v39 = base.B2i32(v36 != int32(0))
									}
								}
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v39 = base.B2i32(v36 != int32(0))
							}
						}
					default:
						v39 = v21
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v39)
				v44 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v44)
				*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v44)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1])))
				if v48 != 0 {
					F_ReleaseBuffer(m, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v53 != int32(1) {
							if l1 == int32(0) {
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v67 <= int32(0) {
								} else {
									v71 = v67 * int32(48)
									if v71 == int32(0) {
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v74, l1, v71)
									}
								}
							}
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
							return
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							if v56 != 0 {
								if l1 == int32(0) {
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v67 <= int32(0) {
									} else {
										v71 = v67 * int32(48)
										if v71 == int32(0) {
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v74, l1, v71)
										}
									}
								}
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
								return
							} else {
								v58 = F_palloc(m, int32(_a_F_btrescan_0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
									*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
									if l1 == int32(0) {
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v67 <= int32(0) {
										} else {
											v71 = v67 * int32(48)
											if v71 == int32(0) {
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v74, l1, v71)
											}
										}
									}
									v77 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
									return
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v53 != int32(1) {
						if l1 == int32(0) {
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v67 <= int32(0) {
							} else {
								v71 = v67 * int32(48)
								if v71 == int32(0) {
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									base.MemoryCopy(m, v74, l1, v71)
								}
							}
						}
						v77 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
						return
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
						if v56 != 0 {
							if l1 == int32(0) {
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v67 <= int32(0) {
								} else {
									v71 = v67 * int32(48)
									if v71 == int32(0) {
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v74, l1, v71)
									}
								}
							}
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
							return
						} else {
							v58 = F_palloc(m, int32(_a_F_btrescan_0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
								*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
								if l1 == int32(0) {
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v67 <= int32(0) {
									} else {
										v71 = v67 * int32(48)
										if v71 == int32(0) {
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v74, l1, v71)
										}
									}
								}
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
								return
							}
						}
					}
				}
			}
		}
	} else {
		v21 = int32(0)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
		if v22 != 0 {
			v39 = v21
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			switch v24 {
			case 0, 5:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
				if v27 != int32(112) {
					v39 = v21
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_btrescan[0]))
					if v31 <= int32(0) {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
						if v34 != 0 {
							v39 = v21
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
							if v35 != 0 {
								v39 = v21
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v39 = base.B2i32(v36 != int32(0))
							}
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v39 = base.B2i32(v36 != int32(0))
					}
				}
			default:
				v39 = v21
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v39)
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v44)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v44)
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1])))
		if v48 != 0 {
			F_ReleaseBuffer(m, v48)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v53 != int32(1) {
					if l1 == int32(0) {
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v67 <= int32(0) {
						} else {
							v71 = v67 * int32(48)
							if v71 == int32(0) {
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								base.MemoryCopy(m, v74, l1, v71)
							}
						}
					}
					v77 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
					return
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
					if v56 != 0 {
						if l1 == int32(0) {
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v67 <= int32(0) {
							} else {
								v71 = v67 * int32(48)
								if v71 == int32(0) {
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									base.MemoryCopy(m, v74, l1, v71)
								}
							}
						}
						v77 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
						return
					} else {
						v58 = F_palloc(m, int32(_a_F_btrescan_0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
							*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
							if l1 == int32(0) {
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v67 <= int32(0) {
								} else {
									v71 = v67 * int32(48)
									if v71 == int32(0) {
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v74, l1, v71)
									}
								}
							}
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
							return
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			if v53 != int32(1) {
				if l1 == int32(0) {
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v67 <= int32(0) {
					} else {
						v71 = v67 * int32(48)
						if v71 == int32(0) {
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							base.MemoryCopy(m, v74, l1, v71)
						}
					}
				}
				v77 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
				return
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
				if v56 != 0 {
					if l1 == int32(0) {
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v67 <= int32(0) {
						} else {
							v71 = v67 * int32(48)
							if v71 == int32(0) {
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								base.MemoryCopy(m, v74, l1, v71)
							}
						}
					}
					v77 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
					return
				} else {
					v58 = F_palloc(m, int32(_a_F_btrescan_0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
						*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
						if l1 == int32(0) {
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v67 <= int32(0) {
							} else {
								v71 = v67 * int32(48)
								if v71 == int32(0) {
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									base.MemoryCopy(m, v74, l1, v71)
								}
							}
						}
						v77 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v77
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v77
						return
					}
				}
			}
		}
	}
}
func F_btrestrpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
	if int32(0) <= v7 {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+100)) = v7
		return
	} else {
		v12 = v6 + int32(56)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
		if v13 == int32(-1) {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[0])))
			if v29 != int32(-1) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[1])))
				if v34 != 0 {
					F_IncrBufferRefCount(m, v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
						v41 = v37*int32(10) + int32(58)
						if v41 != 0 {
							base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
						} else {
						}
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
						if v43 == int32(0) {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
							if v46 == int32(0) {
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
								base.MemoryCopy(m, v43, v49, v46)
							}
						}
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						if v52 == int32(0) {
							return
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							F__bt_start_array_keys(m, l0, v55)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								v58 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
								return
							}
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
					v41 = v37*int32(10) + int32(58)
					if v41 != 0 {
						base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
					} else {
					}
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
					if v43 == int32(0) {
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
						if v46 == int32(0) {
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
							base.MemoryCopy(m, v43, v49, v46)
						}
					}
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					if v52 == int32(0) {
						return
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						F__bt_start_array_keys(m, l0, v55)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							v58 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
							return
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
				return
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
			if int32(0) < v16 {
				F__bt_killitems(m, l0)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v21 == int32(0) {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[0])))
						if v29 != int32(-1) {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[1])))
							if v34 != 0 {
								F_IncrBufferRefCount(m, v34)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
									v41 = v37*int32(10) + int32(58)
									if v41 != 0 {
										base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
									} else {
									}
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
									if v43 == int32(0) {
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
										if v46 == int32(0) {
										} else {
											v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
											base.MemoryCopy(m, v43, v49, v46)
										}
									}
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									if v52 == int32(0) {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
										F__bt_start_array_keys(m, l0, v55)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											v58 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
											return
										}
									}
								}
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
								v41 = v37*int32(10) + int32(58)
								if v41 != 0 {
									base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
								} else {
								}
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v43 == int32(0) {
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
									if v46 == int32(0) {
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
										base.MemoryCopy(m, v43, v49, v46)
									}
								}
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
								if v52 == int32(0) {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
									F__bt_start_array_keys(m, l0, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										v58 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
										return
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
							return
						}
					} else {
						F_ReleaseBuffer(m, v21)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[0])))
							if v29 != int32(-1) {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[1])))
								if v34 != 0 {
									F_IncrBufferRefCount(m, v34)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
										v41 = v37*int32(10) + int32(58)
										if v41 != 0 {
											base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
										} else {
										}
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
										if v43 == int32(0) {
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
											if v46 == int32(0) {
											} else {
												v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
												base.MemoryCopy(m, v43, v49, v46)
											}
										}
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
										if v52 == int32(0) {
											return
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
											F__bt_start_array_keys(m, l0, v55)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												v58 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
												return
											}
										}
									}
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
									v41 = v37*int32(10) + int32(58)
									if v41 != 0 {
										base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
									} else {
									}
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
									if v43 == int32(0) {
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
										if v46 == int32(0) {
										} else {
											v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
											base.MemoryCopy(m, v43, v49, v46)
										}
									}
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									if v52 == int32(0) {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
										F__bt_start_array_keys(m, l0, v55)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											v58 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
											return
										}
									}
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
								return
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v21 == int32(0) {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[0])))
					if v29 != int32(-1) {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[1])))
						if v34 != 0 {
							F_IncrBufferRefCount(m, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
								v41 = v37*int32(10) + int32(58)
								if v41 != 0 {
									base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
								} else {
								}
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v43 == int32(0) {
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
									if v46 == int32(0) {
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
										base.MemoryCopy(m, v43, v49, v46)
									}
								}
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
								if v52 == int32(0) {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
									F__bt_start_array_keys(m, l0, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										v58 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
										return
									}
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
							v41 = v37*int32(10) + int32(58)
							if v41 != 0 {
								base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
							} else {
							}
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							if v43 == int32(0) {
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
								if v46 == int32(0) {
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
									base.MemoryCopy(m, v43, v49, v46)
								}
							}
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
							if v52 == int32(0) {
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
								F__bt_start_array_keys(m, l0, v55)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v58 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
									return
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
						return
					}
				} else {
					F_ReleaseBuffer(m, v21)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[0])))
						if v29 != int32(-1) {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[1])))
							if v34 != 0 {
								F_IncrBufferRefCount(m, v34)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
									v41 = v37*int32(10) + int32(58)
									if v41 != 0 {
										base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
									} else {
									}
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
									if v43 == int32(0) {
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
										if v46 == int32(0) {
										} else {
											v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
											base.MemoryCopy(m, v43, v49, v46)
										}
									}
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									if v52 == int32(0) {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
										F__bt_start_array_keys(m, l0, v55)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											v58 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
											return
										}
									}
								}
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
								v41 = v37*int32(10) + int32(58)
								if v41 != 0 {
									base.MemoryCopy(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
								} else {
								}
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v43 == int32(0) {
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
									if v46 == int32(0) {
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
										base.MemoryCopy(m, v43, v49, v46)
									}
								}
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
								if v52 == int32(0) {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
									F__bt_start_array_keys(m, l0, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										v58 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v58)
										return
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
							return
						}
					}
				}
			}
		}
	}
}
func F_btvacuumscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
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
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v293 int32
	_ = v293
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
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int64
	_ = v413
	var v415 int64
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v493 float64
	_ = v493
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v622 int32
	_ = v622
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v700 int32
	_ = v700
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v822 int32
	_ = v822
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v876 float64
	_ = v876
	var v878 float64
	_ = v878
	var v879 int32
	_ = v879
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1036 int32
	_ = v1036
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1222 int32
	_ = v1222
	var v1234 int32
	_ = v1234
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1287 int32
	_ = v1287
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1356 int32
	_ = v1356
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1405 int32
	_ = v1405
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1467 int32
	_ = v1467
	var v1474 int32
	_ = v1474
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1489 int64
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1502 int32
	_ = v1502
	var v1516 int32
	_ = v1516
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1618 float64
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1724 int32
	_ = v1724
	var v1763 float64
	_ = v1763
	var v1764 float64
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1799 int32
	_ = v1799
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1838 int32
	_ = v1838
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1864 int32
	_ = v1864
	var v1874 int32
	_ = v1874
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1979 int32
	_ = v1979
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v2005 int32
	_ = v2005
	var v2010 int32
	_ = v2010
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2024 int32
	_ = v2024
	var v2037 int32
	_ = v2037
	var v2044 int32
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2054 int32
	_ = v2054
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2099 int32
	_ = v2099
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2146 int32
	_ = v2146
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2183 int32
	_ = v2183
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2196 int32
	_ = v2196
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2207 int32
	_ = v2207
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2291 int32
	_ = v2291
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2304 int32
	_ = v2304
	var v2310 int32
	_ = v2310
	var v2312 int32
	_ = v2312
	var v2318 int32
	_ = v2318
	var v2322 int32
	_ = v2322
	var v2326 int32
	_ = v2326
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2348 int32
	_ = v2348
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2379 int32
	_ = v2379
	var v2384 int32
	_ = v2384
	var v2390 int32
	_ = v2390
	var v2392 int32
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2397 int32
	_ = v2397
	var v2404 int32
	_ = v2404
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2418 int32
	_ = v2418
	var v2422 int32
	_ = v2422
	var v2430 int32
	_ = v2430
	var v2434 int32
	_ = v2434
	var v2440 int32
	_ = v2440
	var v2442 int32
	_ = v2442
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2462 int32
	_ = v2462
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2483 int32
	_ = v2483
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2492 int32
	_ = v2492
	var v2496 int32
	_ = v2496
	var v2500 int32
	_ = v2500
	var v2504 int32
	_ = v2504
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2523 int32
	_ = v2523
	var v2529 int32
	_ = v2529
	var v2532 int64
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2543 int32
	_ = v2543
	var v2545 int32
	_ = v2545
	var v2551 int32
	_ = v2551
	var v2552 int64
	_ = v2552
	var v2562 int32
	_ = v2562
	var v2568 int32
	_ = v2568
	var v2570 int32
	_ = v2570
	var v2576 int32
	_ = v2576
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2601 int32
	_ = v2601
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2705 int32
	_ = v2705
	var v2711 int32
	_ = v2711
	var v2713 int32
	_ = v2713
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2721 int32
	_ = v2721
	var v2725 int32
	_ = v2725
	var v2729 int32
	_ = v2729
	var v2731 int32
	_ = v2731
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2748 int32
	_ = v2748
	var v2751 int32
	_ = v2751
	var v2753 int32
	_ = v2753
	var v2755 int32
	_ = v2755
	var v2757 int32
	_ = v2757
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2765 int32
	_ = v2765
	var v2767 int32
	_ = v2767
	var v2771 int32
	_ = v2771
	var v2777 int32
	_ = v2777
	var v2779 int32
	_ = v2779
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2792 int32
	_ = v2792
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2812 int32
	_ = v2812
	var v2816 int32
	_ = v2816
	var v2822 int32
	_ = v2822
	var v2824 int32
	_ = v2824
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2836 int32
	_ = v2836
	var v2844 int32
	_ = v2844
	var __phi2844 int32
	_ = __phi2844
	var v2845 int32
	_ = v2845
	var __phi2845 int32
	_ = __phi2845
	var v2847 int32
	_ = v2847
	var __phi2847 int32
	_ = __phi2847
	var v2854 int32
	_ = v2854
	var __phi2854 int32
	_ = __phi2854
	var v2891 int32
	_ = v2891
	var v2904 int32
	_ = v2904
	var v2906 int32
	_ = v2906
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2913 int32
	_ = v2913
	var v2914 int32
	_ = v2914
	var v2927 int32
	_ = v2927
	var v2932 int32
	_ = v2932
	var v2935 int32
	_ = v2935
	var v2938 int32
	_ = v2938
	var v2941 int32
	_ = v2941
	var v2943 int32
	_ = v2943
	var v2945 int32
	_ = v2945
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2952 int32
	_ = v2952
	var v2954 int32
	_ = v2954
	var v2958 int32
	_ = v2958
	var v2964 int32
	_ = v2964
	var v2966 int32
	_ = v2966
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2978 int32
	_ = v2978
	var v2987 int32
	_ = v2987
	var v2989 int32
	_ = v2989
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3039 int32
	_ = v3039
	var v3045 int32
	_ = v3045
	var v3047 int32
	_ = v3047
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3059 int32
	_ = v3059
	var v3062 int32
	_ = v3062
	var v3064 int32
	_ = v3064
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3092 int32
	_ = v3092
	var v3097 int32
	_ = v3097
	var v3109 int32
	_ = v3109
	var v3112 int32
	_ = v3112
	var v3113 int32
	_ = v3113
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3119 int32
	_ = v3119
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3126 int32
	_ = v3126
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3134 int32
	_ = v3134
	var v3140 int32
	_ = v3140
	var v3142 int32
	_ = v3142
	var v3148 int32
	_ = v3148
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3151 int32
	_ = v3151
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3175 int32
	_ = v3175
	var v3180 int32
	_ = v3180
	var v3185 int32
	_ = v3185
	var v3187 int32
	_ = v3187
	var v3190 int32
	_ = v3190
	var v3192 int32
	_ = v3192
	var v3195 int32
	_ = v3195
	var v3197 int32
	_ = v3197
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3205 int32
	_ = v3205
	var v3210 int32
	_ = v3210
	var v3216 int32
	_ = v3216
	var v3218 int32
	_ = v3218
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3227 int32
	_ = v3227
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3233 int32
	_ = v3233
	var v3235 int32
	_ = v3235
	var v3239 int32
	_ = v3239
	var v3245 int32
	_ = v3245
	var v3247 int32
	_ = v3247
	var v3253 int32
	_ = v3253
	var v3255 int32
	_ = v3255
	var v3256 int32
	_ = v3256
	var v3262 int32
	_ = v3262
	var v3264 int32
	_ = v3264
	var v3266 int32
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3271 int32
	_ = v3271
	var v3278 int32
	_ = v3278
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3300 int32
	_ = v3300
	var v3306 int32
	_ = v3306
	var v3308 int32
	_ = v3308
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3320 int32
	_ = v3320
	var v3325 int32
	_ = v3325
	var v3327 int32
	_ = v3327
	var v3332 int32
	_ = v3332
	var v3338 int32
	_ = v3338
	var v3340 int32
	_ = v3340
	var v3346 int32
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3348 int64
	_ = v3348
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3356 int32
	_ = v3356
	var v3358 int32
	_ = v3358
	var v3361 int32
	_ = v3361
	var v3364 int32
	_ = v3364
	var v3366 int32
	_ = v3366
	var v3369 int32
	_ = v3369
	var v3377 int32
	_ = v3377
	var v3382 int32
	_ = v3382
	var v3384 int32
	_ = v3384
	var v3386 int32
	_ = v3386
	var v3388 int32
	_ = v3388
	var v3392 int32
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3398 int32
	_ = v3398
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3404 int32
	_ = v3404
	var v3408 int32
	_ = v3408
	var v3412 int32
	_ = v3412
	var v3416 int32
	_ = v3416
	var v3422 int32
	_ = v3422
	var v3434 int32
	_ = v3434
	var v3439 int64
	_ = v3439
	var v3440 int32
	_ = v3440
	var v3444 int32
	_ = v3444
	var v3445 int32
	_ = v3445
	var v3447 int32
	_ = v3447
	var v3449 int32
	_ = v3449
	var v3451 int32
	_ = v3451
	var v3453 int32
	_ = v3453
	var v3455 int32
	_ = v3455
	var v3457 int32
	_ = v3457
	var v3464 int32
	_ = v3464
	var v3467 int64
	_ = v3467
	var v3468 int32
	_ = v3468
	var v3472 int64
	_ = v3472
	var v3476 int32
	_ = v3476
	var v3482 int32
	_ = v3482
	var v3484 int32
	_ = v3484
	var v3490 int32
	_ = v3490
	var v3491 int64
	_ = v3491
	var v3496 int32
	_ = v3496
	var v3497 int32
	_ = v3497
	var v3501 int32
	_ = v3501
	var v3507 int32
	_ = v3507
	var v3509 int32
	_ = v3509
	var v3515 int32
	_ = v3515
	var v3521 int32
	_ = v3521
	var v3527 int32
	_ = v3527
	var v3529 int32
	_ = v3529
	var v3535 int32
	_ = v3535
	var v3542 int32
	_ = v3542
	var v3546 int32
	_ = v3546
	var v3548 int32
	_ = v3548
	var v3552 int32
	_ = v3552
	var v3559 int32
	_ = v3559
	var v3561 int32
	_ = v3561
	var v3567 int32
	_ = v3567
	var v3569 int32
	_ = v3569
	var v3572 int32
	_ = v3572
	var v3574 int32
	_ = v3574
	var v3577 int32
	_ = v3577
	var v3579 int32
	_ = v3579
	var v3584 int32
	_ = v3584
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3592 int32
	_ = v3592
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3599 int32
	_ = v3599
	var v3601 int32
	_ = v3601
	var v3603 int32
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3607 int32
	_ = v3607
	var v3610 int32
	_ = v3610
	var v3611 int32
	_ = v3611
	var v3613 int32
	_ = v3613
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3628 int32
	_ = v3628
	var v3636 int32
	_ = v3636
	var v3641 int32
	_ = v3641
	var v3647 int32
	_ = v3647
	var v3660 int32
	_ = v3660
	var v3702 int32
	_ = v3702
	var v3705 int32
	_ = v3705
	var v3707 int32
	_ = v3707
	var v3709 int32
	_ = v3709
	var v3711 int32
	_ = v3711
	var v3714 int32
	_ = v3714
	var v3715 int32
	_ = v3715
	var v3718 int32
	_ = v3718
	var v3720 int32
	_ = v3720
	var v3724 int32
	_ = v3724
	var v3728 int32
	_ = v3728
	var v3733 int32
	_ = v3733
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3748 int32
	_ = v3748
	var v3753 int32
	_ = v3753
	var v3757 int32
	_ = v3757
	var v3760 int32
	_ = v3760
	var v3761 int32
	_ = v3761
	var v3762 int32
	_ = v3762
	var v3773 int32
	_ = v3773
	var v3778 int32
	_ = v3778
	var v3782 int32
	_ = v3782
	var v3783 int32
	_ = v3783
	var v3793 int32
	_ = v3793
	var v3798 int32
	_ = v3798
	var v3803 int32
	_ = v3803
	var v3851 int32
	_ = v3851
	var v3852 int32
	_ = v3852
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3865 int32
	_ = v3865
	var v3870 int32
	_ = v3870
	var v3923 int32
	_ = v3923
	var v3975 int32
	_ = v3975
	var v4060 int32
	_ = v4060
	var v4082 int32
	_ = v4082
	var v4112 int32
	_ = v4112
	var v4137 int32
	_ = v4137
	var v4138 int32
	_ = v4138
	var v4140 int32
	_ = v4140
	var v4141 int32
	_ = v4141
	var v4142 int32
	_ = v4142
	var v4144 int32
	_ = v4144
	var v4149 int32
	_ = v4149
	var v4150 int32
	_ = v4150
	var v4155 int32
	_ = v4155
	var v4156 int32
	_ = v4156
	var v4164 int32
	_ = v4164
	var v4169 int32
	_ = v4169
	var v4173 int32
	_ = v4173
	var v4174 int32
	_ = v4174
	var v4175 int32
	_ = v4175
	var v4186 int32
	_ = v4186
	var v4199 int32
	_ = v4199
	var v4205 int32
	_ = v4205
	var v4208 int32
	_ = v4208
	var v4224 int32
	_ = v4224
	var v4231 int32
	_ = v4231
	var v4235 int32
	_ = v4235
	var v4240 int32
	_ = v4240
	var v4242 int32
	_ = v4242
	var v4243 int32
	_ = v4243
	var v4246 int32
	_ = v4246
	var v4250 int32
	_ = v4250
	var v4252 int32
	_ = v4252
	var v4253 int32
	_ = v4253
	var v4261 int32
	_ = v4261
	var v4262 int32
	_ = v4262
	var v4268 int32
	_ = v4268
	var v4273 int32
	_ = v4273
	var v4275 int32
	_ = v4275
	var v4277 int32
	_ = v4277
	var v4278 int32
	_ = v4278
	var v4280 int32
	_ = v4280
	var v4281 int32
	_ = v4281
	var v4284 int32
	_ = v4284
	var v4285 int32
	_ = v4285
	var v4286 int32
	_ = v4286
	var v4287 int32
	_ = v4287
	var v4288 int32
	_ = v4288
	var v4289 int32
	_ = v4289
	var v4290 int32
	_ = v4290
	var v4295 int32
	_ = v4295
	var v4343 int32
	_ = v4343
	var v4346 int32
	_ = v4346
	var v4347 int32
	_ = v4347
	var v4348 int64
	_ = v4348
	var v4349 int32
	_ = v4349
	var v4350 int32
	_ = v4350
	var v4354 int32
	_ = v4354
	var v4355 int32
	_ = v4355
	var v4356 int32
	_ = v4356
	var v4360 int32
	_ = v4360
	var v4361 int32
	_ = v4361
	var v4413 int32
	_ = v4413
	var v4416 int32
	_ = v4416
	var v4465 int32
	_ = v4465
	var v4516 int32
	_ = v4516
	var v4518 int32
	_ = v4518
	v5 = l4
	v6 = int32(0)
	v47 = int64(0)
	v51 = m.G0
	v53 = v51 - int32(2512)
	m.G0 = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+28)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+40)) = uint16(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v53)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v53)+32)) = l2
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0]))
	v73 = F_AllocSetContextCreateInternal(m, v68, int32(_a_F_btvacuumscan_0), v6, int32(_a_F_btvacuumscan_1), int32(_a_F_btvacuumscan_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v75 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+48)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v53)+44)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v53)+56)) = v75
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v81 = v53 + int32(24)
	v82 = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = v82
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[1]))
	v91 = v87 << (uint(int32(6)) % 32) & int32(268435392)
	if base.Ui32(v91) <= base.Ui32(v82) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+24)))
	if v107 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v94 = v82
	goto L8
L7:
	;
	v94 = v91
	goto L8
L8:
	;
	if base.Ui32(int32(67108863)) <= base.Ui32(v94) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v97 = int32(67108863)
	goto L11
L10:
	;
	v97 = v94
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+28)) = v97
	v100 = F_palloc(m, int32(_a_F_btvacuumscan_3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+32)) = v100
	goto L5
L13:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
	v113 = base.B2i32(v110 == int32(0))
	goto L15
L14:
	;
	v113 = v6
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v118 = int32(0)
	v123 = F_read_stream_begin_relation(m, int32(13), v117, v55, v118, int32(120), v53+int32(16), v118)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v125 = l0
	v126 = l1
	v137 = v53
	v150 = v55
	v156 = v123
	v159 = v113
	goto L17
L17:
	;
	if v159 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	F_read_stream_end(m, v156)
	mBase = m.M
	v4273 = m.ExcPending
	if v4273 != 0 {
		goto L1
	} else {
		goto L741
	}
L19:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+9)))
	if v190 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v178 = F_RelationGetNumberOfBlocksInFork(m, v150, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_LockRelationForExtension(m, v150, int32(7))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v189 = v178
	goto L19
L24:
	;
	v184 = F_RelationGetNumberOfBlocksInFork(m, v150, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_UnlockRelationForExtension(m, v150, int32(7))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v189 = v184
	goto L19
L27:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[2]))
	if v197 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	if base.Ui32(v238) < base.Ui32(v189) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	goto L29
L31:
	;
	goto L30
L32:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btvacuumscan[3])))
	if v201&int32(1) == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v206 = int32(_a_F_btvacuumscan_4)
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	v209 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v208 + v209
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v212 + v209
	v216 = int32(0)
	v218 = int32(_a_F_btvacuumscan_5)
	v219 = base.AtomicRmwOr32(m, v216, v218, v216)
	*(*int64)(unsafe.Add(mBase, uint32(v197+int32(120))+232)) = base.I64_extend_i32_u(v189)
	v227 = base.AtomicRmwOr32(m, v216, v218, v216)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v228 + v209
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v234 - v209
	goto L31
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v189
	v241 = v125
	v242 = v126
	v253 = v137
	v266 = v150
	v272 = v156
	v275 = v159
	goto L37
L35:
	;
	goto L36
L36:
	;
	goto L18
L37:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v295 = F_read_stream_next_buffer(m, v272, int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v4224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4174)+9)))
	if v4224 != int32(1) {
		v241 = v4174
		v242 = v4175
		v253 = v4186
		v266 = v4199
		v272 = v4205
		v275 = v4208
		goto L37
	} else {
		goto L736
	}
L41:
	;
	F__bt_relbuf(m, v325)
	mBase = m.M
	v4173 = m.ExcPending
	if v4173 != 0 {
		goto L1
	} else {
		goto L735
	}
L42:
	;
	v4149 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4150 = m.ExcPending
	if v4150 != 0 {
		goto L1
	} else {
		goto L730
	}
L43:
	;
	if v295 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v253)+24))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v253)+36))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v253)+32))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v253)+28))
	if v295 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	F_read_stream_reset(m, v272)
	mBase = m.M
	v4144 = m.ExcPending
	if v4144 != 0 {
		goto L1
	} else {
		goto L729
	}
L47:
	;
	v322 = v241
	v323 = v242
	v325 = v295
	v331 = v299
	v332 = v321
	v334 = v253
	v345 = v302
	v347 = v266
	v349 = v321
	v353 = v272
	v354 = v301
	v356 = v275
	v357 = v297
	v363 = v298
	v364 = v300
	goto L51
L48:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306+(v295^int32(-1))<<(uint(int32(6))%32))+16))
	v321 = v312
	goto L47
L49:
	;
	goto L50
L50:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v314+v295<<(uint(int32(6))%32)+int32(-64))+16))
	v321 = v320
	goto L47
L51:
	;
	F_LockBuffer(m, v325, int32(1))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v325 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	if v4112 == int32(0) {
		v4174 = v322
		v4175 = v323
		v4186 = v334
		v4199 = v347
		v4205 = v353
		v4208 = v356
		goto L40
	} else {
		goto L726
	}
L55:
	;
	F__bt_relbuf(m, v325)
	mBase = m.M
	v4082 = m.ExcPending
	if v4082 != 0 {
		goto L1
	} else {
		goto L725
	}
L56:
	;
	v444 = int32(0)
	v446 = v419 & int32(_a_F_btvacuumscan_6)
	if v446&int32(16) == v444 {
		goto L84
	} else {
		goto L85
	}
L57:
	;
	v4060 = int32(0)
	goto L55
L58:
	;
	F_RecordFreeIndexPage(m, v331, v332)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L83
	}
L59:
	;
	v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392)+14)))
	if v393 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v378+(v325^int32(-1))<<(uint(int32(2))%32))))
	v392 = v384
	goto L59
L61:
	;
	goto L62
L62:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v392 = v386 + v325<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	F__bt_checkpage(m, v331, v325)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v332 != v349 {
		goto L42
	} else {
		goto L82
	}
L66:
	;
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392)+16)))
	v397 = v392 + v396
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+12)))
	if v332 != v349 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v398&int32(17) != int32(1) {
		goto L42
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v398&int32(4) != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	if v398&int32(4) != 0 {
		goto L41
	} else {
		goto L71
	}
L71:
	;
	v406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+14)))
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+40)))
	if v406 != v407 {
		goto L41
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	if v398&int32(256) != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v419 = v398
	goto L75
L75:
	;
	if v419&int32(4) == int32(0) {
		goto L56
	} else {
		goto L81
	}
L76:
	;
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v392)+24))
	v415 = v413
	goto L78
L77:
	;
	v415 = int64(3)
	goto L78
L78:
	;
	v416 = F_GlobalVisCheckRemovableFullXid(m, v363, v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v416 != 0 {
		goto L58
	} else {
		goto L80
	}
L80:
	;
	v418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+12)))
	v419 = v418
	goto L75
L81:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v345)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v345)+28)) = v424 + int32(1)
	goto L57
L82:
	;
	goto L58
L83:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v345)+28))
	v434 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v345)+28)) = v433 + v434
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v345)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v345)+32)) = v437 + v434
	goto L57
L84:
	;
	if v446&int32(1) == int32(0) {
		v4060 = v444
		goto L55
	} else {
		goto L87
	}
L85:
	;
	v1799 = v444
	goto L86
L86:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v334)+44))
	F_MemoryContextReset(m, v1820)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L1
	} else {
		goto L249
	}
L87:
	;
	F_LockBuffer(m, v325, int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_LockBufferForCleanup(m, v325)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v460 = int32(0)
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+40)))
	if v462 == v460 {
		v479 = v460
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v482 != 0 {
		goto L100
	} else {
		goto L101
	}
L91:
	;
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+14)))
	if v466 != v462 {
		v479 = int32(0)
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+12)))
	if v469&int32(32) != 0 {
		v479 = int32(0)
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if base.Ui32(v472) < base.Ui32(v349) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v475 = v472
	goto L96
L95:
	;
	v475 = int32(0)
	goto L96
L96:
	;
	if v472 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v477 = v475
	goto L99
L98:
	;
	v477 = int32(0)
	goto L99
L99:
	;
	v479 = v477
	goto L90
L100:
	;
	v483 = int32(2)
	goto L102
L101:
	;
	v483 = int32(1)
	goto L102
L102:
	;
	v484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v484) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v492 = int32(base.Ui32(v484+int32(_a_F_btvacuumscan_7)) >> (uint(int32(2)) % 32))
	goto L105
L104:
	;
	v492 = int32(0)
	goto L105
L105:
	;
	v493 = float64(0)
	if v354 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v879 = int32(0)
	if base.B2i32(v834 <= v879)&base.B2i32(v835 <= v879) == v879 {
		goto L146
	} else {
		goto L147
	}
L107:
	;
	v834 = v460
	v835 = int32(0)
	v876 = v493
	v878 = float64(0)
	goto L106
L108:
	;
	goto L109
L109:
	;
	v498 = int32(0)
	v501 = v492 & int32(_a_F_btvacuumscan_6)
	if base.Ui32(v501) < base.Ui32(v483) {
		v834 = v460
		v835 = v498
		v876 = v493
		v878 = float64(0)
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v505 = int32(0)
	v513 = v460
	v514 = v498
	v515 = v483
	v520 = v505
	v521 = v505
	goto L111
L111:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v392+int32(20)+v515&int32(_a_F_btvacuumscan_6)<<(uint(int32(2))%32))))
	v565 = v392 + v562&int32(_a_F_btvacuumscan_8)
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+7)))
	if v566&int32(32) != 0 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v834 = v777
	v835 = v778
	v876 = base.F64_convert_i32_s(v785)
	v878 = base.F64_convert_i32_s(v784)
	goto L106
L113:
	;
	v822 = v515 + int32(1)
	if base.Ui32(v822&int32(_a_F_btvacuumscan_6)) <= base.Ui32(v501) {
		v513 = v777
		v514 = v778
		v515 = v822
		v520 = v784
		v521 = v785
		goto L111
	} else {
		goto L144
	}
L114:
	;
	v588 = v569 & int32(4095)
	if v588 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L115:
	;
	v569 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565)+4)))
	if v569&int32(_a_F_btvacuumscan_1) != 0 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v573 = m.T0[v354].(func(*base.Module, int32, int32) int32)(m, v565, v364)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	if v573 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v577 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v334+int32(1696)+v513<<(uint(v577)%32)))) = uint16(v515)
	v777 = v513 + v577
	v778 = v514
	v784 = v520
	v785 = v521 + v577
	goto L113
L121:
	;
	goto L122
L122:
	;
	v777 = v513
	v778 = v514
	v784 = v520 + int32(1)
	v785 = v521
	goto L113
L123:
	;
	v777 = v726
	v778 = v727
	v784 = v520 + v737
	v785 = v734
	goto L113
L124:
	;
	v726 = v513
	v727 = v514
	v734 = v521
	v737 = int32(0)
	goto L123
L125:
	;
	goto L126
L126:
	;
	v592 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565)+2)))
	v593 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565))))
	v602 = int32(0)
	v607 = v602
	v609 = v602
	v622 = v602
	goto L127
L127:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v334)+36))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v334)+32))
	v660 = m.T0[v659].(func(*base.Module, int32, int32) int32)(m, v592+(v565+v593<<(uint(int32(16))%32))+v607*int32(6), v658)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	if v683 == int32(0) {
		v726 = v513
		v727 = v514
		v734 = v521
		v737 = v684
		goto L123
	} else {
		goto L139
	}
L129:
	;
	v687 = v607 + int32(1)
	if v687 != v588 {
		v607 = v687
		v609 = v683
		v622 = v684
		goto L127
	} else {
		goto L138
	}
L130:
	;
	if v660 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v683 = v609
	v684 = v622 + int32(1)
	goto L129
L132:
	;
	goto L133
L133:
	;
	if v609 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v668 = F_palloc(m, v588<<(uint(int32(1))%32)+int32(8))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v675 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v609)+6)))
	v676 = int32(1)
	v677 = v675 + v676
	*(*uint16)(unsafe.Add(mBase, uint32(v609)+6)) = uint16(v677)
	*(*uint16)(unsafe.Add(mBase, uint32(v609+v675<<(uint(v676)%32))+8)) = uint16(v607)
	v683 = v609
	v684 = v622
	goto L129
L137:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v668)+8)) = uint16(v607)
	v671 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v668)+6)) = uint16(v671)
	*(*uint16)(unsafe.Add(mBase, uint32(v668)+4)) = uint16(v515)
	*(*int32)(unsafe.Add(mBase, uint32(v668))) = v565
	v683 = v668
	v684 = v622
	goto L129
L138:
	;
	goto L128
L139:
	;
	if int32(0) < v684 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334-int32(-64)+v514<<(uint(int32(2))%32)))) = v683
	v700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565)+4)))
	v726 = v513
	v727 = v514 + int32(1)
	v734 = v521 - v684 + v700&int32(4095)
	v737 = v684
	goto L123
L141:
	;
	goto L142
L142:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v334+int32(1696)+v513<<(uint(int32(1))%32)))) = uint16(v515)
	v712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565)+4)))
	F_pfree(m, v683)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v726 = v513 + int32(1)
	v727 = v514
	v734 = v521 + v712&int32(4095)
	v737 = v684
	goto L123
L144:
	;
	goto L112
L145:
	;
	if base.Ui32(v483) <= base.Ui32(v1724&int32(_a_F_btvacuumscan_6)) {
		goto L241
	} else {
		goto L242
	}
L146:
	;
	v887 = v334 + int32(1696)
	v889 = v334 - int32(-64)
	v890 = int32(0)
	v895 = m.G0
	v897 = v895 - int32(832)
	m.G0 = v897
	if v325 < v890 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	goto L148
L148:
	;
	v1694 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+40)))
	if v1694 == int32(0) {
		v1724 = v492
		goto L145
	} else {
		goto L238
	}
L149:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918)+118)))
	if v919 != int32(112) {
		v932 = int32(0)
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v902 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v902+(v325^int32(-1))<<(uint(int32(2))%32))))
	v916 = v908
	goto L149
L151:
	;
	goto L152
L152:
	;
	v910 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v916 = v910 + v325<<(uint(int32(13))%32) + int32(-8192)
	goto L149
L153:
	;
	if int32(0) < v835 {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v924 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if int32(0) < v924 {
		v932 = int32(1)
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v331)+32))
	if v928 != 0 {
		v932 = int32(0)
		goto L153
	} else {
		goto L156
	}
L156:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v331)+40))
	v932 = base.B2i32(v929 == int32(0))
	goto L153
L157:
	;
	if int32(0) < v834 {
		goto L198
	} else {
		goto L199
	}
L158:
	;
	v943 = v890
	v946 = v890
	goto L161
L159:
	;
	goto L160
L160:
	;
	v1385 = int32(_a_F_btvacuumscan_4)
	v1387 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1387 + int32(1)
	v1393 = v890
	v1405 = v890
	goto L157
L161:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v889+v946<<(uint(int32(2))%32))))
	F__bt_update_posting(m, v988)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	v1007 = int32(0)
	if v932 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v991 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v988)+6)))
	v994 = int32(1)
	v997 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v988)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v897+int32(16)+v946<<(uint(v994)%32)))) = uint16(v997)
	v1003 = v943 + v991<<(uint(v994)%32) + int32(2)
	v1005 = v946 + v994
	if v1005 != v835 {
		v943 = v1003
		v946 = v1005
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v1009 = int32(0)
	v1010 = F_palloc(m, v1003)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	v1222 = v890
	v1234 = v1007
	goto L167
L167:
	;
	v1270 = int32(_a_F_btvacuumscan_4)
	v1272 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1272 + int32(1)
	v1287 = v1007
	goto L184
L168:
	;
	if v835 != int32(1) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v1222 = v1003
	v1234 = v1010
	goto L167
L170:
	;
	v1020 = v890
	v1022 = v1009
	v1036 = v890
	goto L173
L171:
	;
	v1106 = v890
	v1108 = v1009
	goto L172
L172:
	;
	v1154 = v1106 + v1010
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v889+v1108<<(uint(int32(2))%32))))
	v1159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1158)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1154))) = uint16(v1159)
	v1162 = v1159 << (uint(int32(1)) % 32)
	if v1162 == int32(0) {
		goto L169
	} else {
		goto L183
	}
L173:
	;
	v1069 = int32(2)
	v1071 = v889 + v1022<<(uint(v1069)%32)
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1071)))
	v1073 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1072)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1020+v1010))) = uint16(v1073)
	v1076 = v1020 + v1069
	v1078 = v1073 << (uint(int32(1)) % 32)
	if v1078 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v835&int32(1) == int32(0) {
		goto L169
	} else {
		goto L182
	}
L175:
	;
	base.MemoryCopy(m, v1076+v1010, v1072+int32(8), v1078)
	goto L177
L176:
	;
	goto L177
L177:
	;
	v1083 = v1076 + v1078
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+4))
	v1086 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1085)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1010+v1083))) = uint16(v1086)
	v1089 = v1083 + int32(2)
	v1091 = v1086 << (uint(int32(1)) % 32)
	if v1091 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	base.MemoryCopy(m, v1089+v1010, v1085+int32(8), v1091)
	goto L180
L179:
	;
	goto L180
L180:
	;
	v1096 = v1089 + v1091
	v1097 = int32(2)
	v1098 = v1022 + v1097
	v1100 = v1036 + v1097
	if v1100 != v835&int32(2147483646) {
		v1020 = v1096
		v1022 = v1098
		v1036 = v1100
		goto L173
	} else {
		goto L181
	}
L181:
	;
	goto L174
L182:
	;
	v1106 = v1096
	v1108 = v1098
	goto L172
L183:
	;
	base.MemoryCopy(m, v1154+int32(2), v1158+int32(8), v1162)
	goto L169
L184:
	;
	v1331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v897+int32(16)+v1287<<(uint(int32(1))%32)))))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v889+v1287<<(uint(int32(2))%32))))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1335)))
	v1337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1336)+6)))
	v1344 = F_PageIndexTupleOverwrite(m, v916, v1331, v1336, (v1337&int32(_a_F_btvacuumscan_9)+int32(7))&int32(_a_F_btvacuumscan_10))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L186
	}
L185:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L1
	} else {
		goto L191
	}
L186:
	;
	if v1344 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1347 = v1287 + int32(1)
	if v1347 != v835 {
		v1287 = v1347
		goto L184
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	goto L185
L190:
	;
	v1393 = v1222
	v1405 = v1234
	goto L157
L191:
	;
	if v325 < int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v897))) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v897)+4)) = v1372 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_11), v897)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L1
	} else {
		goto L196
	}
L193:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1356+(v325^int32(-1))<<(uint(int32(6))%32))+16))
	v1371 = v1362
	goto L192
L194:
	;
	goto L195
L195:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1364+v325<<(uint(int32(6))%32)+int32(-64))+16))
	v1371 = v1370
	goto L192
L196:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(1200), int32(_a_F_btvacuumscan_13))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	F_PageIndexMultiDelete(m, v916, v887, v834)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v1445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v916)+16)))
	v1446 = v916 + v1445
	v1447 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1446)+14)) = uint16(v1447)
	v1449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1446)+12)))
	v1451 = v1449 & int32(_a_F_btvacuumscan_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v1446)+12)) = uint16(v1451)
	F_MarkBufferDirty(m, v325)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L1
	} else {
		goto L202
	}
L201:
	;
	goto L200
L202:
	;
	if v932 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v897)+14)) = uint16(v835)
	*(*uint16)(unsafe.Add(mBase, uint32(v897)+12)) = uint16(v834)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v1495 = int32(_a_F_btvacuumscan_4)
	v1497 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1497 - int32(1)
	if v1405 != 0 {
		goto L219
	} else {
		goto L220
	}
L206:
	;
	F_XLogRegisterBuffer(m, int32(0), v325, int32(8))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_XLogRegisterData(m, v897+int32(12), int32(4))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	if int32(0) < v834 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_XLogRegisterBufData(m, int32(0), v887, v834<<(uint(int32(1))%32))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L1
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	if int32(0) < v835 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	goto L211
L213:
	;
	F_XLogRegisterBufData(m, int32(0), v897+int32(16), v835<<(uint(int32(1))%32))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1489 = F_XLogInsert(m, int32(11), int32(192))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L1
	} else {
		goto L218
	}
L216:
	;
	F_XLogRegisterBufData(m, int32(0), v1405, v1393)
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	goto L215
L218:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v916))) = base.I64_rotr(v1489, int64(32))
	goto L205
L219:
	;
	F_pfree(m, v1405)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	if int32(0) < v835 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	goto L221
L223:
	;
	v1516 = int32(0)
	goto L226
L224:
	;
	goto L225
L225:
	;
	m.G0 = v897 + int32(832)
	v1618 = *(*float64)(unsafe.Add(mBase, uint32(v345)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v345)+16)) = base.F64_add(v876, v1618)
	v1621 = int32(0)
	v1622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1622) {
		goto L230
	} else {
		goto L231
	}
L226:
	;
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v889+v1516<<(uint(int32(2))%32))))
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1558)))
	F_pfree(m, v1559)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L1
	} else {
		goto L228
	}
L227:
	;
	goto L225
L228:
	;
	v1563 = v1516 + int32(1)
	if v1563 != v835 {
		v1516 = v1563
		goto L226
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	v1630 = int32(base.Ui32(v1622+int32(_a_F_btvacuumscan_7)) >> (uint(int32(2)) % 32))
	goto L232
L231:
	;
	v1630 = v1621
	goto L232
L232:
	;
	if v835 <= int32(0) {
		v1724 = v1630
		goto L145
	} else {
		goto L233
	}
L233:
	;
	v1635 = v1621
	goto L234
L234:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v334-int32(-64)+v1635<<(uint(int32(2))%32))))
	F_pfree(m, v1688)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L1
	} else {
		goto L236
	}
L235:
	;
	v1724 = v1630
	goto L145
L236:
	;
	v1692 = v1635 + int32(1)
	if v1692 != v835 {
		v1635 = v1692
		goto L234
	} else {
		goto L237
	}
L237:
	;
	goto L235
L238:
	;
	v1697 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+14)))
	if v1697 != v1694 {
		v1724 = v492
		goto L145
	} else {
		goto L239
	}
L239:
	;
	v1699 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v397)+14)) = uint16(v1699)
	F_MarkBufferDirtyHint(m, v325, int32(1))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v1724 = v492
	goto L145
L241:
	;
	if v354 != 0 {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	goto L243
L243:
	;
	if v332 != v349 {
		v4060 = v479
		goto L55
	} else {
		goto L248
	}
L244:
	;
	v1763 = v878
	goto L246
L245:
	;
	v1763 = base.F64_convert_i32_u((v1724 - v483 + int32(1)) & int32(_a_F_btvacuumscan_6))
	goto L246
L246:
	;
	v1764 = *(*float64)(unsafe.Add(mBase, uint32(v345)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v345)+8)) = base.F64_add(v1763, v1764)
	F__bt_relbuf(m, v325)
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v4112 = v479
	goto L54
L248:
	;
	v1799 = v479
	goto L86
L249:
	;
	v1823 = int32(_a_F_btvacuumscan_15)
	v1824 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0]))
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v334)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0])) = v1826
	v1829 = v334 + int32(24)
	v1830 = int32(0)
	v1831 = m.G0
	v1833 = v1831 - int32(288)
	m.G0 = v1833
	if v325 < v1830 {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v1864 = v325
	v1874 = v1830
	goto L261
L251:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1838+(v325^int32(-1))<<(uint(int32(6))%32))+16))
	v1853 = v1844
	goto L250
L252:
	;
	goto L253
L253:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1846+v325<<(uint(int32(6))%32)+int32(-64))+16))
	v1853 = v1852
	goto L250
L254:
	;
	m.G0 = v1833 + int32(288)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0])) = v1824
	v4112 = v1799
	goto L54
L255:
	;
	F_ReleaseBuffer(m, v1864)
	mBase = m.M
	v3975 = m.ExcPending
	if v3975 != 0 {
		goto L1
	} else {
		goto L724
	}
L256:
	;
	F_LockBuffer(m, v1864, int32(0))
	mBase = m.M
	v3923 = m.ExcPending
	if v3923 != 0 {
		goto L1
	} else {
		goto L723
	}
L257:
	;
	v3851 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3852 = m.ExcPending
	if v3852 != 0 {
		goto L1
	} else {
		goto L718
	}
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3782 = m.ExcPending
	if v3782 != 0 {
		goto L1
	} else {
		goto L715
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3757 = m.ExcPending
	if v3757 != 0 {
		goto L1
	} else {
		goto L711
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3738 = m.ExcPending
	if v3738 != 0 {
		goto L1
	} else {
		goto L708
	}
L261:
	;
	v1904 = int32(0)
	v1905 = base.B2i32(v1904 <= v1864)
	if v1905 == v1904 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3724 = m.ExcPending
	if v3724 != 0 {
		goto L1
	} else {
		goto L705
	}
L263:
	;
	v1924 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1923)+16)))
	v1925 = v1924 + v1923
	v1926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1925)+12)))
	if v1926&int32(5) != int32(1) {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1909+(v1864^int32(-1))<<(uint(int32(2))%32))))
	v1923 = v1915
	goto L263
L265:
	;
	goto L266
L266:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v1923 = v1917 + v1864<<(uint(int32(13))%32) + int32(-8192)
	goto L263
L267:
	;
	if v1926&int32(16) == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	if v1926&int32(2) != 0 {
		goto L291
	} else {
		goto L292
	}
L270:
	;
	v1962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1925)+12)))
	if v1962&int32(4) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L271:
	;
	v1937 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	if v1937 == int32(0) {
		goto L270
	} else {
		goto L273
	}
L273:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+208)) = v1944 + int32(4)
	F_errmsg(m, int32(_a_F_btvacuumscan_16), v1833+int32(208))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	F_errhint(m, int32(_a_F_btvacuumscan_17), int32(0))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(1863), int32(_a_F_btvacuumscan_0))
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	goto L270
L278:
	;
	F_LockBuffer(m, v1864, int32(0))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L289
	}
L279:
	;
	v1969 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	if v1969 == int32(0) {
		goto L278
	} else {
		goto L281
	}
L281:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	if v1864 < int32(0) {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+196)) = v1853
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+192)) = v1994
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+200)) = v1995 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_18), v1833+int32(192))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L1
	} else {
		goto L287
	}
L284:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1979+(v1864^int32(-1))<<(uint(int32(6))%32))+16))
	v1994 = v1985
	goto L283
L285:
	;
	goto L286
L286:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1987+v1864<<(uint(int32(6))%32)+int32(-64))+16))
	v1994 = v1993
	goto L283
L287:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(1871), int32(_a_F_btvacuumscan_0))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	goto L278
L289:
	;
	goto L255
L290:
	;
	if v1926&int32(16) == int32(0) {
		goto L297
	} else {
		goto L298
	}
L291:
	;
	F_LockBuffer(m, v1864, int32(0))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L1
	} else {
		goto L295
	}
L292:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+4))
	if base.B2i32(v2018 == int32(0))|v1926&int32(128) != 0 {
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v2024 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1923)+12)))
	if base.B2i32(base.Ui32(v2024) < base.Ui32(int32(25)))|base.B2i32((v2024+int32(_a_F_btvacuumscan_7))&int32(_a_F_btvacuumscan_19) == int32(0)) != 0 {
		goto L290
	} else {
		goto L294
	}
L294:
	;
	goto L291
L295:
	;
	goto L255
L296:
	;
	goto L262
L297:
	;
	if v1874 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L298:
	;
	v2601 = v1926
	goto L299
L299:
	;
	if v2601&int32(16) != 0 {
		goto L429
	} else {
		goto L430
	}
L300:
	;
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v1923)+24))
	v2048 = F_CopyIndexTuple(m, v1923+v2044&int32(_a_F_btvacuumscan_8))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L1
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v1829)))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2094)+4))
	if v1905 == int32(0) {
		goto L317
	} else {
		goto L318
	}
L303:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v1925)))
	if v1864 < int32(0) {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	F_LockBuffer(m, v1864, int32(0))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L1
	} else {
		goto L308
	}
L305:
	;
	v2054 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v2054+(v1864^int32(-1))<<(uint(int32(6))%32))+16))
	v2069 = v2060
	goto L304
L306:
	;
	goto L307
L307:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2062+v1864<<(uint(int32(6))%32)+int32(-64))+16))
	v2069 = v2068
	goto L304
L308:
	;
	v2073 = F__bt_leftsib_splitflag(m, v331, v2050, v2069)
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	if v2073 != 0 {
		goto L255
	} else {
		goto L310
	}
L310:
	;
	v2075 = F__bt_mkscankey(m, v331, v2048)
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	v2077 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2075)+3)) = uint16(v2077)
	v2083 = F__bt_search(m, v331, int32(0), v2075, v1833+int32(248), int32(1))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+248))
	F_LockBuffer(m, v2085, int32(0))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_ReleaseBuffer(m, v2085)
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_LockBuffer(m, v1864, int32(2))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1874 = v2083
	goto L261
L316:
	;
	v2114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2113)+16)))
	if v1864 < int32(0) {
		goto L321
	} else {
		goto L322
	}
L317:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v2099+(v1864^int32(-1))<<(uint(int32(2))%32))))
	v2113 = v2105
	goto L316
L318:
	;
	goto L319
L319:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2113 = v2107 + v1864<<(uint(int32(13))%32) + int32(-8192)
	goto L316
L320:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v2113+v2114)+4))
	v2136 = F_ReadBuffer(m, v331, v2135)
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L1
	} else {
		goto L324
	}
L321:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v2118+(v1864^int32(-1))<<(uint(int32(6))%32))+16))
	v2133 = v2124
	goto L320
L322:
	;
	goto L323
L323:
	;
	v2126 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2126+v1864<<(uint(int32(6))%32)+int32(-64))+16))
	v2133 = v2132
	goto L320
L324:
	;
	F_LockBuffer(m, v2136, int32(1))
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	F__bt_checkpage(m, v331, v2136)
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	if v2136 < int32(0) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v2161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2160)+16)))
	v2163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2161+v2160)+12)))
	F_LockBuffer(m, v2136, int32(0))
	mBase = m.M
	v2166 = m.ExcPending
	if v2166 != 0 {
		goto L1
	} else {
		goto L331
	}
L328:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2146+(v2136^int32(-1))<<(uint(int32(2))%32))))
	v2160 = v2152
	goto L327
L329:
	;
	goto L330
L330:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2160 = v2154 + v2136<<(uint(int32(13))%32) + int32(-8192)
	goto L327
L331:
	;
	F_ReleaseBuffer(m, v2136)
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	if v2163&int32(16) != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v2173 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L1
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v2189 = F__bt_getstackbuf(m, v331, v2095, v1874, v2133)
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L1
	} else {
		goto L340
	}
L336:
	;
	if v2173 == int32(0) {
		goto L256
	} else {
		goto L337
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+180)) = v2135
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+176)) = v2133
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_20), v1833+int32(176))
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2128), int32(_a_F_btvacuumscan_21))
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	goto L256
L340:
	;
	if v2189 == int32(0) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v3803 = v2133
	goto L257
L342:
	;
	goto L343
L343:
	;
	v2196 = v1874
	v2199 = v2189
	v2200 = v2135
	v2207 = v2133
	goto L344
L344:
	;
	v2243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2196)+4)))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2196)))
	v2245 = int32(0)
	v2246 = base.B2i32(v2245 <= v2199)
	if v2246 == v2245 {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	if v2246 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L346:
	;
	v2265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2264)+16)))
	v2266 = v2265 + v2264
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v2266)))
	v2268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2264)+12)))
	if base.B2i32(base.Ui32(int32(25)) <= base.Ui32(v2268))&base.B2i32(base.Ui32(v2243) < base.Ui32(int32(base.Ui32(v2268+int32(_a_F_btvacuumscan_7))>>(uint(int32(2))%32))&int32(_a_F_btvacuumscan_6))) == int32(0) {
		goto L350
	} else {
		goto L351
	}
L347:
	;
	v2250 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v2250+(v2199^int32(-1))<<(uint(int32(2))%32))))
	v2264 = v2256
	goto L346
L348:
	;
	goto L349
L349:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2264 = v2258 + v2199<<(uint(int32(13))%32) + int32(-8192)
	goto L346
L350:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v2266)+4))
	F_LockBuffer(m, v2199, int32(0))
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L1
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	goto L345
L353:
	;
	F_ReleaseBuffer(m, v2199)
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	if v2281 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v2291 = int32(2)
	goto L357
L356:
	;
	v2291 = int32(1)
	goto L357
L357:
	;
	if base.B2i32(v2281 == int32(0))|base.B2i32(v2291 != v2243) != 0 {
		goto L256
	} else {
		goto L358
	}
L358:
	;
	v2294 = F__bt_leftsib_splitflag(m, v331, v2267, v2244)
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	if v2294 != 0 {
		goto L256
	} else {
		goto L360
	}
L360:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2196)+8))
	v2297 = F__bt_getstackbuf(m, v331, v2095, v2296, v2244)
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	if v2297 == int32(0) {
		v3803 = v2244
		goto L257
	} else {
		goto L362
	}
L362:
	;
	v2196 = v2296
	v2199 = v2297
	v2200 = v2281
	v2207 = v2244
	goto L344
L363:
	;
	v2322 = (v2243 + int32(1)) & int32(_a_F_btvacuumscan_6)
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v2318+v2322<<(uint(int32(2))%32))+20))
	v2329 = v2326&int32(_a_F_btvacuumscan_8) + v2318
	v2330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2329))))
	v2333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2329)+2)))
	if v2200 != v2330<<(uint(int32(16))%32)|v2333 {
		goto L367
	} else {
		goto L368
	}
L364:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v2304+(v2199^int32(-1))<<(uint(int32(2))%32))))
	v2318 = v2310
	goto L363
L365:
	;
	goto L366
L366:
	;
	v2312 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2318 = v2312 + v2199<<(uint(int32(13))%32) + int32(-8192)
	goto L363
L367:
	;
	v2338 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L1
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	F_PredicateLockPageSplit(m, v331, v2133, v2135)
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L1
	} else {
		goto L383
	}
L370:
	;
	if v2338 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L1
	} else {
		goto L374
	}
L372:
	;
	goto L373
L373:
	;
	F_LockBuffer(m, v2199, int32(0))
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L1
	} else {
		goto L381
	}
L374:
	;
	v2343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2329)+2)))
	v2344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2329))))
	if v2199 < int32(0) {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v2364 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+160)) = v2364 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+156)) = v2363
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+152)) = v2343 | v2344<<(uint(int32(16))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+148)) = v2207
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+144)) = v2200
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_22), v1833+int32(144))
	mBase = m.M
	v2379 = m.ExcPending
	if v2379 != 0 {
		goto L1
	} else {
		goto L379
	}
L376:
	;
	v2348 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(v2348+(v2199^int32(-1))<<(uint(int32(6))%32))+16))
	v2363 = v2354
	goto L375
L377:
	;
	goto L378
L378:
	;
	v2356 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v2356+v2199<<(uint(int32(6))%32)+int32(-64))+16))
	v2363 = v2362
	goto L375
L379:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2187), int32(_a_F_btvacuumscan_21))
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	goto L373
L381:
	;
	F_ReleaseBuffer(m, v2199)
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	goto L256
L383:
	;
	v2395 = int32(_a_F_btvacuumscan_4)
	v2397 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v2397 + int32(1)
	if v2246 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(v2418+v2243<<(uint(int32(2))%32))+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2422&int32(_a_F_btvacuumscan_8)+v2418))) = base.I32_rotr(v2200, int32(16))
	F_PageIndexTupleDelete(m, v2418, v2322)
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L1
	} else {
		goto L388
	}
L385:
	;
	v2404 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2410 = *(*int32)(unsafe.Add(mBase, uint32(v2404+(v2199^int32(-1))<<(uint(int32(2))%32))))
	v2418 = v2410
	goto L384
L386:
	;
	goto L387
L387:
	;
	v2412 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2418 = v2412 + v2199<<(uint(int32(13))%32) + int32(-8192)
	goto L384
L388:
	;
	if v1905 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	v2449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2448)+16)))
	v2450 = v2449 + v2448
	v2451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2450)+12)))
	v2453 = v2451 | int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v2450)+12)) = uint16(v2453)
	v2456 = base.B2i32(v2133 == v2207)
	if v2133 == v2207 {
		goto L393
	} else {
		goto L394
	}
L390:
	;
	v2434 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v2434+(v1864^int32(-1))<<(uint(int32(2))%32))))
	v2448 = v2440
	goto L389
L391:
	;
	goto L392
L392:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2448 = v2442 + v1864<<(uint(int32(13))%32) + int32(-8192)
	goto L389
L393:
	;
	v2457 = int32(-1)
	goto L395
L394:
	;
	v2457 = v2207
	goto L395
L395:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1833)+222)) = uint16(v2457)
	if v2133 == v2207 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v2462 = int32(-1)
	goto L398
L397:
	;
	v2462 = int32(base.Ui32(v2207) >> (uint(int32(16)) % 32))
	goto L398
L398:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1833)+220)) = uint16(v2462)
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+224)) = int32(537395200)
	v2470 = F_PageIndexTupleOverwrite(m, v2448, int32(1), v1833+int32(220), int32(8))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	if v2470 == int32(0) {
		goto L296
	} else {
		goto L400
	}
L400:
	;
	F_MarkBufferDirty(m, v2199)
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	F_MarkBufferDirty(m, v1864)
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v2478 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	v2479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2478)+118)))
	if v2479 != int32(112) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v2583 = int32(_a_F_btvacuumscan_4)
	v2585 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v2585 - int32(1)
	F_LockBuffer(m, v2199, int32(0))
	mBase = m.M
	v2591 = m.ExcPending
	if v2591 != 0 {
		goto L1
	} else {
		goto L427
	}
L404:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if v2483 <= int32(0) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(v331)+32))
	if v2486 != 0 {
		goto L403
	} else {
		goto L408
	}
L406:
	;
	goto L407
L407:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1833)+248)) = uint16(v2243)
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+252)) = v2133
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+264)) = v2457
	F_XLogBeginInsert(m)
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L1
	} else {
		goto L410
	}
L408:
	;
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v331)+40))
	if v2487 != 0 {
		goto L403
	} else {
		goto L409
	}
L409:
	;
	goto L407
L410:
	;
	F_XLogRegisterBuffer(m, int32(0), v1864, int32(6))
	mBase = m.M
	v2496 = m.ExcPending
	if v2496 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	F_XLogRegisterBuffer(m, int32(1), v2199, int32(8))
	mBase = m.M
	v2500 = m.ExcPending
	if v2500 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	if v1905 == int32(0) {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	v2519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2518)+16)))
	v2520 = v2519 + v2518
	v2521 = *(*int32)(unsafe.Add(mBase, uint32(v2520)))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+256)) = v2521
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(v2520)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+260)) = v2523
	F_XLogRegisterData(m, v1833+int32(248), int32(20))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L1
	} else {
		goto L417
	}
L414:
	;
	v2504 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v2504+(v1864^int32(-1))<<(uint(int32(2))%32))))
	v2518 = v2510
	goto L413
L415:
	;
	goto L416
L416:
	;
	v2512 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2518 = v2512 + v1864<<(uint(int32(13))%32) + int32(-8192)
	goto L413
L417:
	;
	v2532 = F_XLogInsert(m, int32(11), int32(176))
	mBase = m.M
	v2533 = m.ExcPending
	if v2533 != 0 {
		goto L1
	} else {
		goto L418
	}
L418:
	;
	if v2246 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	v2552 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v2551))) = base.I64_rotr(v2532, v2552)
	if v1905 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L420:
	;
	v2537 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v2537+(v2199^int32(-1))<<(uint(int32(2))%32))))
	v2551 = v2543
	goto L419
L421:
	;
	goto L422
L422:
	;
	v2545 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2551 = v2545 + v2199<<(uint(int32(13))%32) + int32(-8192)
	goto L419
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2576)+4)) = base.I32_wrap_i64(v2532)
	*(*int32)(unsafe.Add(mBase, uint32(v2576))) = base.I32_wrap_i64(int64(base.Ui64(v2532) >> (uint(v2552) % 64)))
	goto L403
L424:
	;
	v2562 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v2562+(v1864^int32(-1))<<(uint(int32(2))%32))))
	v2576 = v2568
	goto L423
L425:
	;
	goto L426
L426:
	;
	v2570 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2576 = v2570 + v1864<<(uint(int32(13))%32) + int32(-8192)
	goto L423
L427:
	;
	F_ReleaseBuffer(m, v2199)
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	v2594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1925)+12)))
	v2601 = v2594
	goto L299
L429:
	;
	v2649 = v1864 ^ int32(-1)
	v2651 = v1864 << (uint(int32(13)) % 32)
	goto L432
L430:
	;
	v3660 = int32(0)
	goto L431
L431:
	;
	v3702 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+4))
	F_LockBuffer(m, v1864, int32(0))
	mBase = m.M
	v3705 = m.ExcPending
	if v3705 != 0 {
		goto L1
	} else {
		goto L695
	}
L432:
	;
	if v1864 < int32(0) {
		goto L435
	} else {
		goto L436
	}
L433:
	;
	v3641 = int32(2)
	if v3205 != 0 {
		goto L692
	} else {
		goto L693
	}
L434:
	;
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+4))
	if v1905 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L435:
	;
	v2705 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v2705+(v1864^int32(-1))<<(uint(int32(6))%32))+16))
	v2720 = v2711
	goto L434
L436:
	;
	goto L437
L437:
	;
	v2713 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(v2713+v1864<<(uint(int32(6))%32)+int32(-64))+16))
	v2720 = v2719
	goto L434
L438:
	;
	v2736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2735)+16)))
	v2737 = v2736 + v2735
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v2737)+4))
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2737)))
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v2735)+24))
	v2743 = v2735 + v2740&int32(_a_F_btvacuumscan_8)
	v2744 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2743)+2)))
	v2745 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2743))))
	F_LockBuffer(m, v1864, int32(0))
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L1
	} else {
		goto L442
	}
L439:
	;
	v2725 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2729 = *(*int32)(unsafe.Add(mBase, uint32(v2725+v2649<<(uint(int32(2))%32))))
	v2735 = v2729
	goto L438
L440:
	;
	goto L441
L441:
	;
	v2731 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2735 = v2731 + v2651 + int32(-8192)
	goto L438
L442:
	;
	v2751 = v2744 | v2745<<(uint(int32(16))%32)
	v2753 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v2753 != 0 {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2755 = m.ExcPending
	if v2755 != 0 {
		goto L1
	} else {
		goto L446
	}
L444:
	;
	goto L445
L445:
	;
	v2757 = int32(1)
	if v2751 == int32(-1) {
		goto L448
	} else {
		goto L449
	}
L446:
	;
	goto L445
L447:
	;
	v2803 = int32(0)
	if v2798 == v2803 {
		v2987 = int32(0)
		v2989 = v2803
		goto L463
	} else {
		goto L464
	}
L448:
	;
	v2798 = v2739
	v2799 = v1864
	v2800 = v2720
	v2801 = v2757
	v2802 = int32(0)
	goto L447
L449:
	;
	goto L450
L450:
	;
	v2761 = F_ReadBuffer(m, v331, v2751)
	mBase = m.M
	v2762 = m.ExcPending
	if v2762 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	F_LockBuffer(m, v2761, int32(1))
	mBase = m.M
	v2765 = m.ExcPending
	if v2765 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	F__bt_checkpage(m, v331, v2761)
	mBase = m.M
	v2767 = m.ExcPending
	if v2767 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	if v2761 < int32(0) {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	v2786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2785)+16)))
	v2787 = v2786 + v2785
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(v2787)+8))
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v2787)))
	F_LockBuffer(m, v2761, int32(0))
	mBase = m.M
	v2792 = m.ExcPending
	if v2792 != 0 {
		goto L1
	} else {
		goto L458
	}
L455:
	;
	v2771 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v2771+(v2761^int32(-1))<<(uint(int32(2))%32))))
	v2785 = v2777
	goto L454
L456:
	;
	goto L457
L457:
	;
	v2779 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2785 = v2779 + v2761<<(uint(int32(13))%32) + int32(-8192)
	goto L454
L458:
	;
	if v2720 == v2751 {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v2798 = v2789
	v2799 = v2761
	v2800 = v2720
	v2801 = v2757
	v2802 = v2788
	goto L447
L460:
	;
	goto L461
L461:
	;
	F_LockBuffer(m, v1864, int32(2))
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	v2798 = v2789
	v2799 = v2761
	v2800 = v2751
	v2801 = int32(0)
	v2802 = v2788
	goto L447
L463:
	;
	F_LockBuffer(m, v2799, int32(2))
	mBase = m.M
	v3033 = m.ExcPending
	if v3033 != 0 {
		goto L1
	} else {
		goto L507
	}
L464:
	;
	v2806 = F_ReadBuffer(m, v331, v2798)
	mBase = m.M
	v2807 = m.ExcPending
	if v2807 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F_LockBuffer(m, v2806, int32(2))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	F__bt_checkpage(m, v331, v2806)
	mBase = m.M
	v2812 = m.ExcPending
	if v2812 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	if v2806 < int32(0) {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	v2831 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2830)+16)))
	v2832 = v2831 + v2830
	v2833 = *(*int32)(unsafe.Add(mBase, uint32(v2832)+4))
	v2834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2832)+12)))
	v2836 = v2834 & int32(4)
	if v2836 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L469:
	;
	v2816 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v2816+(v2806^int32(-1))<<(uint(int32(2))%32))))
	v2830 = v2822
	goto L468
L470:
	;
	goto L471
L471:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2830 = v2824 + v2806<<(uint(int32(13))%32) + int32(-8192)
	goto L468
L472:
	;
	if v2833 == v2800 {
		v2987 = v2806
		v2989 = v2798
		goto L463
	} else {
		goto L475
	}
L473:
	;
	goto L474
L474:
	;
	__phi2844 = v2798
	__phi2845 = v2833
	__phi2847 = v2806
	__phi2854 = v2836
	v2844 = __phi2844
	v2845 = __phi2845
	v2847 = __phi2847
	v2854 = __phi2854
	goto L476
L475:
	;
	goto L474
L476:
	;
	v2891 = int32(0)
	if base.B2i32(base.B2i32(v2845 == v2891)|v2854&int32(_a_F_btvacuumscan_6) == v2891)&base.B2i32(v2844 != v2845) == v2891 {
		goto L478
	} else {
		goto L479
	}
L477:
	;
	v2987 = v2948
	v2989 = v2845
	goto L463
L478:
	;
	F_LockBuffer(m, v2847, int32(0))
	mBase = m.M
	v2904 = m.ExcPending
	if v2904 != 0 {
		goto L1
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	F_LockBuffer(m, v2847, int32(0))
	mBase = m.M
	v2941 = m.ExcPending
	if v2941 != 0 {
		goto L1
	} else {
		goto L493
	}
L481:
	;
	F_ReleaseBuffer(m, v2847)
	mBase = m.M
	v2906 = m.ExcPending
	if v2906 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	v2909 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	if v2909 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2913 = m.ExcPending
	if v2913 != 0 {
		goto L1
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	F_ReleaseBuffer(m, v2799)
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L1
	} else {
		goto L490
	}
L487:
	;
	v2914 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+128)) = v2802
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+132)) = v2914 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+124)) = v1853
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+120)) = v2720
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+116)) = v2800
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+112)) = v2845
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_23), v1833+int32(112))
	mBase = m.M
	v2927 = m.ExcPending
	if v2927 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2439), int32(_a_F_btvacuumscan_24))
	mBase = m.M
	v2932 = m.ExcPending
	if v2932 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	goto L486
L490:
	;
	if v2801 != 0 {
		goto L254
	} else {
		goto L491
	}
L491:
	;
	F_LockBuffer(m, v1864, int32(0))
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	goto L255
L493:
	;
	F_ReleaseBuffer(m, v2847)
	mBase = m.M
	v2943 = m.ExcPending
	if v2943 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	v2945 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v2945 != 0 {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2947 = m.ExcPending
	if v2947 != 0 {
		goto L1
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	v2948 = F_ReadBuffer(m, v331, v2845)
	mBase = m.M
	v2949 = m.ExcPending
	if v2949 != 0 {
		goto L1
	} else {
		goto L499
	}
L498:
	;
	goto L497
L499:
	;
	F_LockBuffer(m, v2948, int32(2))
	mBase = m.M
	v2952 = m.ExcPending
	if v2952 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	F__bt_checkpage(m, v331, v2948)
	mBase = m.M
	v2954 = m.ExcPending
	if v2954 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	if v2948 < int32(0) {
		goto L503
	} else {
		goto L504
	}
L502:
	;
	v2973 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2972)+16)))
	v2974 = v2973 + v2972
	v2975 = *(*int32)(unsafe.Add(mBase, uint32(v2974)+4))
	v2976 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2974)+12)))
	v2978 = v2976 & int32(4)
	if v2978|base.B2i32(v2975 != v2800) != 0 {
		__phi2844 = v2845
		__phi2845 = v2975
		__phi2847 = v2948
		__phi2854 = v2978
		v2844 = __phi2844
		v2845 = __phi2845
		v2847 = __phi2847
		v2854 = __phi2854
		goto L476
	} else {
		goto L506
	}
L503:
	;
	v2958 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2964 = *(*int32)(unsafe.Add(mBase, uint32(v2958+(v2948^int32(-1))<<(uint(int32(2))%32))))
	v2972 = v2964
	goto L502
L504:
	;
	goto L505
L505:
	;
	v2966 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2972 = v2966 + v2948<<(uint(int32(13))%32) + int32(-8192)
	goto L502
L506:
	;
	goto L477
L507:
	;
	v3034 = int32(0)
	v3035 = base.B2i32(v3034 <= v2799)
	if v3035 == v3034 {
		goto L509
	} else {
		goto L510
	}
L508:
	;
	v3054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3053)+16)))
	v3055 = v3054 + v3053
	v3056 = *(*int32)(unsafe.Add(mBase, uint32(v3055)+4))
	if v3056 == int32(0) {
		goto L260
	} else {
		goto L512
	}
L509:
	;
	v3039 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3045 = *(*int32)(unsafe.Add(mBase, uint32(v3039+(v2799^int32(-1))<<(uint(int32(2))%32))))
	v3053 = v3045
	goto L508
L510:
	;
	goto L511
L511:
	;
	v3047 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3053 = v3047 + v2799<<(uint(int32(13))%32) + int32(-8192)
	goto L508
L512:
	;
	v3059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3055)+12)))
	if v3059&int32(6) != 0 {
		goto L260
	} else {
		goto L513
	}
L513:
	;
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(v3055)))
	if v3062 != v2989 {
		goto L259
	} else {
		goto L514
	}
L514:
	;
	v3064 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3053)+12)))
	v3066 = v3064 + int32(_a_F_btvacuumscan_7)
	if v2801 != 0 {
		goto L516
	} else {
		goto L517
	}
L515:
	;
	v3122 = F_ReadBuffer(m, v331, v3056)
	mBase = m.M
	v3123 = m.ExcPending
	if v3123 != 0 {
		goto L1
	} else {
		goto L530
	}
L516:
	;
	v3067 = int32(17)
	if v3059&v3067 == v3067 {
		goto L519
	} else {
		goto L520
	}
L517:
	;
	goto L518
L518:
	;
	if v3059&int32(1)|base.B2i32(base.Ui32(v3064) < base.Ui32(int32(25)))|base.B2i32(v3066&int32(_a_F_btvacuumscan_25) != int32(8)) != 0 {
		goto L258
	} else {
		goto L526
	}
L519:
	;
	if base.B2i32(v3066&int32(_a_F_btvacuumscan_19) == int32(0))|base.B2i32(base.Ui32(v3064) < base.Ui32(int32(25))) != 0 {
		v3121 = int32(-1)
		goto L515
	} else {
		goto L522
	}
L520:
	;
	goto L521
L521:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3082 = m.ExcPending
	if v3082 != 0 {
		goto L1
	} else {
		goto L523
	}
L522:
	;
	goto L521
L523:
	;
	v3083 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+64)) = v2800
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+68)) = v3083 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_26), v1833-int32(-64))
	mBase = m.M
	v3092 = m.ExcPending
	if v3092 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2486), int32(_a_F_btvacuumscan_24))
	mBase = m.M
	v3097 = m.ExcPending
	if v3097 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L526:
	;
	v3109 = *(*int32)(unsafe.Add(mBase, uint32(v3053)+28))
	v3112 = v3053 + v3109&int32(_a_F_btvacuumscan_8)
	v3113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3112))))
	v3116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3112)+2)))
	v3117 = v3113<<(uint(int32(16))%32) | v3116
	if v2720 == v3117 {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v3119 = int32(-1)
	goto L529
L528:
	;
	v3119 = v3117
	goto L529
L529:
	;
	v3121 = v3119
	goto L515
L530:
	;
	F_LockBuffer(m, v3122, int32(2))
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	F__bt_checkpage(m, v331, v3122)
	mBase = m.M
	v3128 = m.ExcPending
	if v3128 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v3129 = int32(0)
	v3130 = base.B2i32(v3129 <= v3122)
	if v3130 == v3129 {
		goto L534
	} else {
		goto L535
	}
L533:
	;
	v3149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3148)+16)))
	v3150 = v3149 + v3148
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(v3150)))
	if v2800 != v3151 {
		goto L537
	} else {
		goto L538
	}
L534:
	;
	v3134 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3140 = *(*int32)(unsafe.Add(mBase, uint32(v3134+(v3122^int32(-1))<<(uint(int32(2))%32))))
	v3148 = v3140
	goto L533
L535:
	;
	goto L536
L536:
	;
	v3142 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3148 = v3142 + v3122<<(uint(int32(13))%32) + int32(-8192)
	goto L533
L537:
	;
	v3155 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3156 = m.ExcPending
	if v3156 != 0 {
		goto L1
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	v3201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3148)+12)))
	v3202 = int32(0)
	v3205 = *(*int32)(unsafe.Add(mBase, uint32(v3150)+4))
	if v2989|v3205 != 0 {
		v3266 = v3202
		v3267 = v3202
		v3268 = v3202
		goto L558
	} else {
		goto L559
	}
L540:
	;
	if v3155 != 0 {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3159 = m.ExcPending
	if v3159 != 0 {
		goto L1
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	if v2987 != 0 {
		goto L547
	} else {
		goto L548
	}
L544:
	;
	v3160 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	v3161 = *(*int32)(unsafe.Add(mBase, uint32(v3150)))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+52)) = v2802
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+48)) = v3161
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+56)) = v3160 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+44)) = v1853
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+40)) = v2720
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+36)) = v2800
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+32)) = v3056
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_27), v1833+int32(32))
	mBase = m.M
	v3175 = m.ExcPending
	if v3175 != 0 {
		goto L1
	} else {
		goto L545
	}
L545:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2540), int32(_a_F_btvacuumscan_24))
	mBase = m.M
	v3180 = m.ExcPending
	if v3180 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	goto L543
L547:
	;
	F_LockBuffer(m, v2987, int32(0))
	mBase = m.M
	v3185 = m.ExcPending
	if v3185 != 0 {
		goto L1
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	F_LockBuffer(m, v3122, int32(0))
	mBase = m.M
	v3190 = m.ExcPending
	if v3190 != 0 {
		goto L1
	} else {
		goto L552
	}
L550:
	;
	F_ReleaseBuffer(m, v2987)
	mBase = m.M
	v3187 = m.ExcPending
	if v3187 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	goto L549
L552:
	;
	F_ReleaseBuffer(m, v3122)
	mBase = m.M
	v3192 = m.ExcPending
	if v3192 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	F_LockBuffer(m, v2799, int32(0))
	mBase = m.M
	v3195 = m.ExcPending
	if v3195 != 0 {
		goto L1
	} else {
		goto L554
	}
L554:
	;
	F_ReleaseBuffer(m, v2799)
	mBase = m.M
	v3197 = m.ExcPending
	if v3197 != 0 {
		goto L1
	} else {
		goto L555
	}
L555:
	;
	if v2801 != 0 {
		goto L254
	} else {
		goto L556
	}
L556:
	;
	F_LockBuffer(m, v1864, int32(0))
	mBase = m.M
	v3200 = m.ExcPending
	if v3200 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	goto L255
L558:
	;
	v3269 = int32(_a_F_btvacuumscan_4)
	v3271 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v3271 + int32(1)
	if v2987 != 0 {
		goto L577
	} else {
		goto L578
	}
L559:
	;
	if v3130 == int32(0) {
		goto L561
	} else {
		goto L562
	}
L560:
	;
	v3225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3224)+16)))
	v3227 = *(*int32)(unsafe.Add(mBase, uint32(v3225+v3224)+4))
	if v3227 != 0 {
		v3266 = v3202
		v3267 = v3202
		v3268 = v3202
		goto L558
	} else {
		goto L564
	}
L561:
	;
	v3210 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(v3210+(v3122^int32(-1))<<(uint(int32(2))%32))))
	v3224 = v3216
	goto L560
L562:
	;
	goto L563
L563:
	;
	v3218 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3224 = v3218 + v3122<<(uint(int32(13))%32) + int32(-8192)
	goto L560
L564:
	;
	v3229 = F_ReadBuffer(m, v331, int32(0))
	mBase = m.M
	v3230 = m.ExcPending
	if v3230 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	F_LockBuffer(m, v3229, int32(2))
	mBase = m.M
	v3233 = m.ExcPending
	if v3233 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	F__bt_checkpage(m, v331, v3229)
	mBase = m.M
	v3235 = m.ExcPending
	if v3235 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	if v3229 < int32(0) {
		goto L569
	} else {
		goto L570
	}
L568:
	;
	v3255 = v3253 + int32(24)
	v3256 = *(*int32)(unsafe.Add(mBase, uint32(v3253)+44))
	if base.Ui32(v3256) <= base.Ui32(v2802+int32(1)) {
		goto L572
	} else {
		goto L573
	}
L569:
	;
	v3239 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3245 = *(*int32)(unsafe.Add(mBase, uint32(v3239+(v3229^int32(-1))<<(uint(int32(2))%32))))
	v3253 = v3245
	goto L568
L570:
	;
	goto L571
L571:
	;
	v3247 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3253 = v3247 + v3229<<(uint(int32(13))%32) + int32(-8192)
	goto L568
L572:
	;
	v3266 = v3253
	v3267 = v3229
	v3268 = v3255
	goto L558
L573:
	;
	goto L574
L574:
	;
	F_LockBuffer(m, v3229, int32(0))
	mBase = m.M
	v3262 = m.ExcPending
	if v3262 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	F_ReleaseBuffer(m, v3229)
	mBase = m.M
	v3264 = m.ExcPending
	if v3264 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	v3266 = v3253
	v3267 = v3202
	v3268 = v3255
	goto L558
L577:
	;
	if v2987 < int32(0) {
		goto L581
	} else {
		goto L582
	}
L578:
	;
	goto L579
L579:
	;
	if v3130 == int32(0) {
		goto L585
	} else {
		goto L586
	}
L580:
	;
	v3293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3292)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v3293+v3292)+4)) = v3056
	goto L579
L581:
	;
	v3278 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(v3278+(v2987^int32(-1))<<(uint(int32(2))%32))))
	v3292 = v3284
	goto L580
L582:
	;
	goto L583
L583:
	;
	v3286 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3292 = v3286 + v2987<<(uint(int32(13))%32) + int32(-8192)
	goto L580
L584:
	;
	v3315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3314)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v3315+v3314))) = v2989
	if v2801 == int32(0) {
		goto L588
	} else {
		goto L589
	}
L585:
	;
	v3300 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3306 = *(*int32)(unsafe.Add(mBase, uint32(v3300+(v3122^int32(-1))<<(uint(int32(2))%32))))
	v3314 = v3306
	goto L584
L586:
	;
	goto L587
L587:
	;
	v3308 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3314 = v3308 + v3122<<(uint(int32(13))%32) + int32(-8192)
	goto L584
L588:
	;
	v3320 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2743)+4)) = uint16(v3320)
	*(*int32)(unsafe.Add(mBase, uint32(v2743))) = base.I32_rotr(v3121, int32(16))
	v3325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2743)+6)))
	v3327 = v3325 | int32(_a_F_btvacuumscan_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2743)+6)) = uint16(v3327)
	goto L590
L589:
	;
	goto L590
L590:
	;
	if v3035 == int32(0) {
		goto L592
	} else {
		goto L593
	}
L591:
	;
	v3347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3346)+16)))
	v3348 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v3349 = m.ExcPending
	if v3349 != 0 {
		goto L1
	} else {
		goto L595
	}
L592:
	;
	v3332 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3338 = *(*int32)(unsafe.Add(mBase, uint32(v3332+(v2799^int32(-1))<<(uint(int32(2))%32))))
	v3346 = v3338
	goto L591
L593:
	;
	goto L594
L594:
	;
	v3340 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3346 = v3340 + v2799<<(uint(int32(13))%32) + int32(-8192)
	goto L591
L595:
	;
	v3350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3346)+16)))
	v3351 = v3346 + v3350
	v3352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3351)+12)))
	v3356 = v3352&int32(_a_F_btvacuumscan_28) | int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v3351)+12)) = uint16(v3356)
	v3358 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v3346)+12)) = uint16(v3358)
	*(*int64)(unsafe.Add(mBase, uint32(v3346)+24)) = v3348
	v3361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3346)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3346)+14)) = uint16(v3361)
	v3364 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3346+v3347)+14)) = uint16(v3364)
	if v3267 != 0 {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	v3366 = *(*int32)(unsafe.Add(mBase, uint32(v3268)+4))
	if base.Ui32(v3366) <= base.Ui32(int32(2)) {
		goto L599
	} else {
		goto L600
	}
L597:
	;
	goto L598
L598:
	;
	F_MarkBufferDirty(m, v3122)
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L1
	} else {
		goto L603
	}
L599:
	;
	v3369 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3266)+64)) = uint8(v3369)
	*(*int64)(unsafe.Add(mBase, uint32(v3266)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v3266)+48)) = v3369
	*(*int32)(unsafe.Add(mBase, uint32(v3266)+28)) = int32(3)
	v3377 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v3266)+12)) = uint16(v3377)
	goto L601
L600:
	;
	goto L601
L601:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3268)+20)) = v2802
	*(*int32)(unsafe.Add(mBase, uint32(v3268)+16)) = v3056
	F_MarkBufferDirty(m, v3267)
	mBase = m.M
	v3382 = m.ExcPending
	if v3382 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	goto L598
L603:
	;
	F_MarkBufferDirty(m, v2799)
	mBase = m.M
	v3386 = m.ExcPending
	if v3386 != 0 {
		goto L1
	} else {
		goto L604
	}
L604:
	;
	if v2987 != 0 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	F_MarkBufferDirty(m, v2987)
	mBase = m.M
	v3388 = m.ExcPending
	if v3388 != 0 {
		goto L1
	} else {
		goto L608
	}
L606:
	;
	goto L607
L607:
	;
	if v2801 == int32(0) {
		goto L609
	} else {
		goto L610
	}
L608:
	;
	goto L607
L609:
	;
	F_MarkBufferDirty(m, v1864)
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L1
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	v3393 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	v3394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3393)+118)))
	if v3394 != int32(112) {
		goto L613
	} else {
		goto L614
	}
L612:
	;
	goto L611
L613:
	;
	v3559 = int32(_a_F_btvacuumscan_4)
	v3561 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v3561 - int32(1)
	if v3267 != 0 {
		goto L660
	} else {
		goto L661
	}
L614:
	;
	v3398 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if v3398 <= int32(0) {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v3401 = *(*int32)(unsafe.Add(mBase, uint32(v331)+32))
	if v3401 != 0 {
		goto L613
	} else {
		goto L618
	}
L616:
	;
	goto L617
L617:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v3404 = m.ExcPending
	if v3404 != 0 {
		goto L1
	} else {
		goto L620
	}
L618:
	;
	v3402 = *(*int32)(unsafe.Add(mBase, uint32(v331)+40))
	if v3402 != 0 {
		goto L613
	} else {
		goto L619
	}
L619:
	;
	goto L617
L620:
	;
	F_XLogRegisterBuffer(m, int32(0), v2799, int32(6))
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L1
	} else {
		goto L621
	}
L621:
	;
	if v2987 != 0 {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	F_XLogRegisterBuffer(m, int32(1), v2987, int32(8))
	mBase = m.M
	v3412 = m.ExcPending
	if v3412 != 0 {
		goto L1
	} else {
		goto L625
	}
L623:
	;
	goto L624
L624:
	;
	F_XLogRegisterBuffer(m, int32(2), v3122, int32(8))
	mBase = m.M
	v3416 = m.ExcPending
	if v3416 != 0 {
		goto L1
	} else {
		goto L626
	}
L625:
	;
	goto L624
L626:
	;
	if v2801 == int32(0) {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	F_XLogRegisterBuffer(m, int32(3), v1864, int32(6))
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L1
	} else {
		goto L630
	}
L628:
	;
	goto L629
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+280)) = v3121
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+276)) = v2738
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+272)) = v2739
	*(*int64)(unsafe.Add(mBase, uint32(v1833)+264)) = v3348
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+256)) = v2802
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+252)) = v3056
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+248)) = v2989
	F_XLogRegisterData(m, v1833+int32(248), int32(36))
	mBase = m.M
	v3434 = m.ExcPending
	if v3434 != 0 {
		goto L1
	} else {
		goto L631
	}
L630:
	;
	goto L629
L631:
	;
	if v3267 == int32(0) {
		goto L633
	} else {
		goto L634
	}
L632:
	;
	if v3130 == int32(0) {
		goto L641
	} else {
		goto L642
	}
L633:
	;
	v3439 = F_XLogInsert(m, int32(11), int32(128))
	mBase = m.M
	v3440 = m.ExcPending
	if v3440 != 0 {
		goto L1
	} else {
		goto L636
	}
L634:
	;
	goto L635
L635:
	;
	F_XLogRegisterBuffer(m, int32(4), v3267, int32(14))
	mBase = m.M
	v3444 = m.ExcPending
	if v3444 != 0 {
		goto L1
	} else {
		goto L637
	}
L636:
	;
	v3472 = v3439
	goto L632
L637:
	;
	v3445 = *(*int32)(unsafe.Add(mBase, uint32(v3268)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+220)) = v3445
	v3447 = *(*int32)(unsafe.Add(mBase, uint32(v3268)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+224)) = v3447
	v3449 = *(*int32)(unsafe.Add(mBase, uint32(v3268)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+228)) = v3449
	v3451 = *(*int32)(unsafe.Add(mBase, uint32(v3268)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+232)) = v3451
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(v3268)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+236)) = v3453
	v3455 = *(*int32)(unsafe.Add(mBase, uint32(v3268)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+240)) = v3455
	v3457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3268)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1833)+244)) = uint8(v3457)
	F_XLogRegisterBufData(m, int32(4), v1833+int32(220), int32(28))
	mBase = m.M
	v3464 = m.ExcPending
	if v3464 != 0 {
		goto L1
	} else {
		goto L638
	}
L638:
	;
	v3467 = F_XLogInsert(m, int32(11), int32(144))
	mBase = m.M
	v3468 = m.ExcPending
	if v3468 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3266))) = base.I64_rotr(v3467, int64(32))
	v3472 = v3467
	goto L632
L640:
	;
	v3491 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v3490))) = base.I64_rotr(v3472, v3491)
	v3496 = base.I32_wrap_i64(int64(base.Ui64(v3472) >> (uint(v3491) % 64)))
	v3497 = base.I32_wrap_i64(v3472)
	if v3035 == int32(0) {
		goto L645
	} else {
		goto L646
	}
L641:
	;
	v3476 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3482 = *(*int32)(unsafe.Add(mBase, uint32(v3476+(v3122^int32(-1))<<(uint(int32(2))%32))))
	v3490 = v3482
	goto L640
L642:
	;
	goto L643
L643:
	;
	v3484 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3490 = v3484 + v3122<<(uint(int32(13))%32) + int32(-8192)
	goto L640
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3515)+4)) = v3497
	*(*int32)(unsafe.Add(mBase, uint32(v3515))) = v3496
	if v2987 != 0 {
		goto L648
	} else {
		goto L649
	}
L645:
	;
	v3501 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3507 = *(*int32)(unsafe.Add(mBase, uint32(v3501+(v2799^int32(-1))<<(uint(int32(2))%32))))
	v3515 = v3507
	goto L644
L646:
	;
	goto L647
L647:
	;
	v3509 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3515 = v3509 + v2799<<(uint(int32(13))%32) + int32(-8192)
	goto L644
L648:
	;
	if v2987 < int32(0) {
		goto L652
	} else {
		goto L653
	}
L649:
	;
	goto L650
L650:
	;
	if v2801 != 0 {
		goto L613
	} else {
		goto L655
	}
L651:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3535)+4)) = v3497
	*(*int32)(unsafe.Add(mBase, uint32(v3535))) = v3496
	goto L650
L652:
	;
	v3521 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3527 = *(*int32)(unsafe.Add(mBase, uint32(v3521+(v2987^int32(-1))<<(uint(int32(2))%32))))
	v3535 = v3527
	goto L651
L653:
	;
	goto L654
L654:
	;
	v3529 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3535 = v3529 + v2987<<(uint(int32(13))%32) + int32(-8192)
	goto L651
L655:
	;
	if v1905 == int32(0) {
		goto L657
	} else {
		goto L658
	}
L656:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3552)+4)) = v3497
	*(*int32)(unsafe.Add(mBase, uint32(v3552))) = v3496
	goto L613
L657:
	;
	v3542 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3546 = *(*int32)(unsafe.Add(mBase, uint32(v3542+v2649<<(uint(int32(2))%32))))
	v3552 = v3546
	goto L656
L658:
	;
	goto L659
L659:
	;
	v3548 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3552 = v3548 + v2651 + int32(-8192)
	goto L656
L660:
	;
	F_LockBuffer(m, v3267, int32(0))
	mBase = m.M
	v3567 = m.ExcPending
	if v3567 != 0 {
		goto L1
	} else {
		goto L663
	}
L661:
	;
	goto L662
L662:
	;
	if v2987 != 0 {
		goto L665
	} else {
		goto L666
	}
L663:
	;
	F_ReleaseBuffer(m, v3267)
	mBase = m.M
	v3569 = m.ExcPending
	if v3569 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	goto L662
L665:
	;
	F_LockBuffer(m, v2987, int32(0))
	mBase = m.M
	v3572 = m.ExcPending
	if v3572 != 0 {
		goto L1
	} else {
		goto L668
	}
L666:
	;
	goto L667
L667:
	;
	F_LockBuffer(m, v3122, int32(0))
	mBase = m.M
	v3577 = m.ExcPending
	if v3577 != 0 {
		goto L1
	} else {
		goto L670
	}
L668:
	;
	F_ReleaseBuffer(m, v2987)
	mBase = m.M
	v3574 = m.ExcPending
	if v3574 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	goto L667
L670:
	;
	F_ReleaseBuffer(m, v3122)
	mBase = m.M
	v3579 = m.ExcPending
	if v3579 != 0 {
		goto L1
	} else {
		goto L671
	}
L671:
	;
	if v2801 == int32(0) {
		goto L672
	} else {
		goto L673
	}
L672:
	;
	F_LockBuffer(m, v2799, int32(0))
	mBase = m.M
	v3584 = m.ExcPending
	if v3584 != 0 {
		goto L1
	} else {
		goto L675
	}
L673:
	;
	goto L674
L674:
	;
	v3587 = *(*int32)(unsafe.Add(mBase, uint32(v2721)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2721)+24)) = v3587 + int32(1)
	if base.Ui32(v2800) <= base.Ui32(v1853) {
		goto L677
	} else {
		goto L678
	}
L675:
	;
	F_ReleaseBuffer(m, v2799)
	mBase = m.M
	v3586 = m.ExcPending
	if v3586 != 0 {
		goto L1
	} else {
		goto L676
	}
L676:
	;
	goto L674
L677:
	;
	v3592 = *(*int32)(unsafe.Add(mBase, uint32(v2721)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2721)+28)) = v3592 + int32(1)
	goto L679
L678:
	;
	goto L679
L679:
	;
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+36))
	v3597 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+28))
	if v3596 != v3597 {
		goto L680
	} else {
		goto L681
	}
L680:
	;
	v3599 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+24))
	if v3599 != v3596 {
		goto L684
	} else {
		goto L685
	}
L681:
	;
	goto L682
L682:
	;
	v3636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1925)+12)))
	if v3636&int32(16) != 0 {
		goto L432
	} else {
		goto L691
	}
L683:
	;
	v3618 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v3617+v3616<<(uint(v3618)%32)))) = v2800
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+32))
	v3623 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v3622+v3623<<(uint(v3618)%32))+8)) = v3348
	v3628 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1829)+36)) = v3628 + int32(1)
	goto L682
L684:
	;
	v3601 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+32))
	v3616 = v3596
	v3617 = v3601
	goto L683
L685:
	;
	goto L686
L686:
	;
	v3603 = v3596 << (uint(int32(1)) % 32)
	if v3603 < v3597 {
		goto L687
	} else {
		goto L688
	}
L687:
	;
	v3605 = v3603
	goto L689
L688:
	;
	v3605 = v3597
	goto L689
L689:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1829)+24)) = v3605
	v3607 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+32))
	v3610 = F_repalloc(m, v3607, v3605<<(uint(int32(4))%32))
	mBase = m.M
	v3611 = m.ExcPending
	if v3611 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1829)+32)) = v3610
	v3613 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+36))
	v3616 = v3613
	v3617 = v3610
	goto L683
L691:
	;
	goto L433
L692:
	;
	v3647 = v3641
	goto L694
L693:
	;
	v3647 = int32(1)
	goto L694
L694:
	;
	v3660 = base.B2i32(base.Ui32(int32(base.Ui32(v3201+int32(_a_F_btvacuumscan_7))>>(uint(v3641)%32))&int32(_a_F_btvacuumscan_6)) < base.Ui32(v3647)) | base.B2i32(base.Ui32(v3201) < base.Ui32(int32(25)))
	goto L431
L695:
	;
	F_ReleaseBuffer(m, v1864)
	mBase = m.M
	v3707 = m.ExcPending
	if v3707 != 0 {
		goto L1
	} else {
		goto L696
	}
L696:
	;
	v3709 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v3709 != 0 {
		goto L697
	} else {
		goto L698
	}
L697:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3711 = m.ExcPending
	if v3711 != 0 {
		goto L1
	} else {
		goto L700
	}
L698:
	;
	goto L699
L699:
	;
	if v3660 == int32(0) {
		goto L254
	} else {
		goto L701
	}
L700:
	;
	goto L699
L701:
	;
	v3714 = F_ReadBuffer(m, v331, v3702)
	mBase = m.M
	v3715 = m.ExcPending
	if v3715 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	F_LockBuffer(m, v3714, int32(2))
	mBase = m.M
	v3718 = m.ExcPending
	if v3718 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	F__bt_checkpage(m, v331, v3714)
	mBase = m.M
	v3720 = m.ExcPending
	if v3720 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	v1864 = v3714
	goto L261
L705:
	;
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_29), int32(0))
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L1
	} else {
		goto L706
	}
L706:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2244), int32(_a_F_btvacuumscan_21))
	mBase = m.M
	v3733 = m.ExcPending
	if v3733 != 0 {
		goto L1
	} else {
		goto L707
	}
L707:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L708:
	;
	v3739 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+16)) = v2800
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+20)) = v3739 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_30), v1833+int32(16))
	mBase = m.M
	v3748 = m.ExcPending
	if v3748 != 0 {
		goto L1
	} else {
		goto L709
	}
L709:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2472), int32(_a_F_btvacuumscan_24))
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L711:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3760 = m.ExcPending
	if v3760 != 0 {
		goto L1
	} else {
		goto L712
	}
L712:
	;
	v3761 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	v3762 = *(*int32)(unsafe.Add(mBase, uint32(v3055)))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+104)) = v2800
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+100)) = v3762
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+96)) = v2989
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+108)) = v3761 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_31), v1833+int32(96))
	mBase = m.M
	v3773 = m.ExcPending
	if v3773 != 0 {
		goto L1
	} else {
		goto L713
	}
L713:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2479), int32(_a_F_btvacuumscan_24))
	mBase = m.M
	v3778 = m.ExcPending
	if v3778 != 0 {
		goto L1
	} else {
		goto L714
	}
L714:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L715:
	;
	v3783 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+84)) = v2800
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+80)) = v2802
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+88)) = v3783 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_32), v1833+int32(80))
	mBase = m.M
	v3793 = m.ExcPending
	if v3793 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2498), int32(_a_F_btvacuumscan_24))
	mBase = m.M
	v3798 = m.ExcPending
	if v3798 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L718:
	;
	if v3851 == int32(0) {
		goto L256
	} else {
		goto L719
	}
L719:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3857 = m.ExcPending
	if v3857 != 0 {
		goto L1
	} else {
		goto L720
	}
L720:
	;
	v3858 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+4)) = v3803
	*(*int32)(unsafe.Add(mBase, uint32(v1833))) = v3858 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_33), v1833)
	mBase = m.M
	v3865 = m.ExcPending
	if v3865 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_12), int32(2848), int32(_a_F_btvacuumscan_34))
	mBase = m.M
	v3870 = m.ExcPending
	if v3870 != 0 {
		goto L1
	} else {
		goto L722
	}
L722:
	;
	goto L256
L723:
	;
	goto L255
L724:
	;
	goto L254
L725:
	;
	v4112 = v4060
	goto L54
L726:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v4137 = m.ExcPending
	if v4137 != 0 {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	v4138 = int32(0)
	v4140 = *(*int32)(unsafe.Add(mBase, uint32(v357)+24))
	v4141 = F_ReadBufferExtended(m, v331, v4138, v4112, v4138, v4140)
	mBase = m.M
	v4142 = m.ExcPending
	if v4142 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	v325 = v4141
	v332 = v4112
	goto L51
L729:
	;
	v125 = v241
	v126 = v242
	v137 = v253
	v150 = v266
	v156 = v272
	v159 = v275
	goto L17
L730:
	;
	if v4149 == int32(0) {
		goto L41
	} else {
		goto L731
	}
L731:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v4155 = m.ExcPending
	if v4155 != 0 {
		goto L1
	} else {
		goto L732
	}
L732:
	;
	v4156 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+4)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v334))) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v334)+8)) = v4156 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_35), v334)
	mBase = m.M
	v4164 = m.ExcPending
	if v4164 != 0 {
		goto L1
	} else {
		goto L733
	}
L733:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_36), int32(1413), int32(_a_F_btvacuumscan_37))
	mBase = m.M
	v4169 = m.ExcPending
	if v4169 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	goto L41
L735:
	;
	v4174 = v322
	v4175 = v323
	v4186 = v334
	v4199 = v347
	v4205 = v353
	v4208 = v356
	goto L40
L736:
	;
	v4231 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[2]))
	if v4231 == int32(0) {
		goto L738
	} else {
		goto L739
	}
L737:
	;
	v241 = v4174
	v242 = v4175
	v253 = v4186
	v266 = v4199
	v272 = v4205
	v275 = v4208
	goto L37
L738:
	;
	goto L737
L739:
	;
	v4235 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btvacuumscan[3])))
	if v4235&int32(1) == int32(0) {
		goto L738
	} else {
		goto L740
	}
L740:
	;
	v4240 = int32(_a_F_btvacuumscan_4)
	v4242 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	v4243 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v4242 + v4243
	v4246 = *(*int32)(unsafe.Add(mBase, uint32(v4231)))
	*(*int32)(unsafe.Add(mBase, uint32(v4231))) = v4246 + v4243
	v4250 = int32(0)
	v4252 = int32(_a_F_btvacuumscan_5)
	v4253 = base.AtomicRmwOr32(m, v4250, v4252, v4250)
	*(*int64)(unsafe.Add(mBase, uint32(v4231+int32(128))+232)) = base.I64_extend_i32_u(v349)
	v4261 = base.AtomicRmwOr32(m, v4250, v4252, v4250)
	v4262 = *(*int32)(unsafe.Add(mBase, uint32(v4231)))
	*(*int32)(unsafe.Add(mBase, uint32(v4231))) = v4262 + v4243
	v4268 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v4268 - v4243
	goto L738
L741:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v189
	v4275 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	F_MemoryContextDelete(m, v4275)
	mBase = m.M
	v4277 = m.ExcPending
	if v4277 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	v4278 = int32(0)
	v4280 = v137 + int32(24)
	v4281 = *(*int32)(unsafe.Add(mBase, uint32(v4280)+36))
	if v4281 == v4278 {
		goto L745
	} else {
		goto L746
	}
L743:
	;
	v4516 = *(*int32)(unsafe.Add(mBase, uint32(v126)+32))
	if v4516 != 0 {
		goto L759
	} else {
		goto L760
	}
L744:
	;
	F_pfree(m, v4416)
	mBase = m.M
	v4465 = m.ExcPending
	if v4465 != 0 {
		goto L1
	} else {
		goto L758
	}
L745:
	;
	v4284 = *(*int32)(unsafe.Add(mBase, uint32(v4280)+32))
	if v4284 != 0 {
		v4416 = v4284
		goto L744
	} else {
		goto L748
	}
L746:
	;
	goto L747
L747:
	;
	v4285 = *(*int32)(unsafe.Add(mBase, uint32(v4280)+4))
	v4286 = *(*int32)(unsafe.Add(mBase, uint32(v4280)))
	v4287 = *(*int32)(unsafe.Add(mBase, uint32(v4286)+4))
	v4288 = F_GetOldestNonRemovableTransactionId(m, v4287)
	mBase = m.M
	v4289 = m.ExcPending
	if v4289 != 0 {
		goto L1
	} else {
		goto L749
	}
L748:
	;
	goto L743
L749:
	;
	v4290 = *(*int32)(unsafe.Add(mBase, uint32(v4280)+36))
	if v4290 <= int32(0) {
		goto L750
	} else {
		goto L751
	}
L750:
	;
	v4413 = *(*int32)(unsafe.Add(mBase, uint32(v4280)+32))
	v4416 = v4413
	goto L744
L751:
	;
	v4295 = v4278
	goto L752
L752:
	;
	v4343 = *(*int32)(unsafe.Add(mBase, uint32(v4280)+32))
	v4346 = v4343 + v4295<<(uint(int32(4))%32)
	v4347 = *(*int32)(unsafe.Add(mBase, uint32(v4346)))
	v4348 = *(*int64)(unsafe.Add(mBase, uint32(v4346)+8))
	v4349 = F_GlobalVisCheckRemovableFullXid(m, v4287, v4348)
	mBase = m.M
	v4350 = m.ExcPending
	if v4350 != 0 {
		goto L1
	} else {
		goto L754
	}
L753:
	;
	goto L750
L754:
	;
	if v4349 == int32(0) {
		goto L750
	} else {
		goto L755
	}
L755:
	;
	F_RecordFreeIndexPage(m, v150, v4347)
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L1
	} else {
		goto L756
	}
L756:
	;
	v4355 = *(*int32)(unsafe.Add(mBase, uint32(v4285)+32))
	v4356 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+32)) = v4355 + v4356
	v4360 = v4295 + v4356
	v4361 = *(*int32)(unsafe.Add(mBase, uint32(v4280)+36))
	if v4360 < v4361 {
		v4295 = v4360
		goto L752
	} else {
		goto L757
	}
L757:
	;
	goto L753
L758:
	;
	goto L743
L759:
	;
	F_FreeSpaceMapVacuum(m, v150)
	mBase = m.M
	v4518 = m.ExcPending
	if v4518 != 0 {
		goto L1
	} else {
		goto L762
	}
L760:
	;
	goto L761
L761:
	;
	m.G0 = v137 + int32(2512)
	return
L762:
	;
	goto L761
}
func F_build_attrmap_by_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v206 int32
	_ = v206
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
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v24 = F_palloc0(m, int32(8))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v22
	v31 = F_palloc0(m, v22<<(uint(int32(1))%32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v31
	if int32(0) < v22 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L43
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L36
	}
L6:
	;
	v43 = int32(0)
	v46 = v31
	v48 = int32(-1)
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v19 + int32(32)
	return v24
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v59 = l1 + v53<<(uint(int32(4))%32) + v43*int32(100)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+111)))
	if v60 != 0 {
		v173 = v46
		v175 = v48
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v181 = v43 + int32(1)
	if v181 != v22 {
		v43 = v181
		v46 = v173
		v48 = v175
		goto L9
	} else {
		goto L35
	}
L12:
	;
	v62 = v59 + int32(20)
	v64 = v59 + int32(24)
	if v21 <= int32(0) {
		v151 = v46
		v153 = v48
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if l2 != 0 {
		v173 = v151
		v175 = v153
		goto L11
	} else {
		goto L33
	}
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+76))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+68))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = v48
	v88 = int32(0)
	goto L15
L15:
	;
	v93 = v87 + int32(1)
	if v93 < v21 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v151 = v46
	v153 = v96
	goto L13
L17:
	;
	v140 = v88 + int32(1)
	if v140 != v21 {
		v87 = v96
		v88 = v140
		goto L15
	} else {
		goto L32
	}
L18:
	;
	v96 = v93
	goto L20
L19:
	;
	v96 = int32(0)
	goto L20
L20:
	;
	v99 = l0 + v69<<(uint(int32(4))%32) + int32(20) + v96*int32(100)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+91)))
	if v100 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v102 = v99 + int32(4)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if base.B2i32(v105 == int32(0))|base.B2i32(v105 != v108) != 0 {
		v126 = v105
		v127 = v108
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v126-v127 != 0 {
		goto L17
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v111 = v64
	v112 = v102
	goto L25
L25:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	if v116 == int32(0) {
		v126 = v116
		v127 = v115
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v126 = v116
	v127 = v115
	goto L23
L27:
	;
	v119 = int32(1)
	if v116 == v115 {
		v111 = v111 + v119
		v112 = v112 + v119
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	if v68 != v129 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v99)+76))
	if v67 != v131 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v46+v43<<(uint(int32(1))%32)))) = uint16(v136)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v151 = v138
	v153 = v96
	goto L13
L32:
	;
	goto L16
L33:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151+v43<<(uint(int32(1))%32)))))
	if v161 == int32(0) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v173 = v151
	v175 = v153
	goto L11
L35:
	;
	goto L10
L36:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_build_attrmap_by_name_0), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v215 = F_format_type_be(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v218 = F_format_type_be(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v64
	F_errdetail(m, int32(_a_F_build_attrmap_by_name_1), v19+int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_build_attrmap_by_name_2), int32(235), int32(_a_F_build_attrmap_by_name_3))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_build_attrmap_by_name_0), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v245 = F_format_type_be(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v248 = F_format_type_be(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v64
	F_errdetail(m, int32(_a_F_build_attrmap_by_name_4), v19)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_build_attrmap_by_name_2), int32(247), int32(_a_F_build_attrmap_by_name_3))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_build_joinrel_partition_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v56 int32
	_ = v56
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
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
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
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
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
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
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
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
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v635 int32
	_ = v635
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v698 int32
	_ = v698
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v733 int32
	_ = v733
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v801 int32
	_ = v801
	var v810 int32
	_ = v810
	var v817 int32
	_ = v817
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v873 int32
	_ = v873
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1088 int32
	_ = v1088
	var v1116 int32
	_ = v1116
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	v7 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_joinrel_partition_info[0])))
	if v26 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L102
	} else {
		goto L244
	}
L2:
	;
	m.G0 = v23 + int32(80)
	return
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+232))
	if v29 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)+232))
	if v32 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+217)))
	if base.B2i32(v35 != int32(1))|base.B2i32(v32 != v29) != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+217)))
	if v40&int32(1) == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v46 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v46
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v46
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v46
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v46
	if l5 == int32(0) {
		v621 = v7
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l2)+232))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+232)) = v839
	v841 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v842 = int32(*(*int16)(unsafe.Add(mBase, uint32(v839)+2)))
	v844 = v842 << (uint(int32(2)) % 32)
	v845 = F_palloc0(m, v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L102
	} else {
		goto L200
	}
L9:
	;
	v627 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	if v627 <= int32(0) {
		goto L2
	} else {
		goto L166
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v56 <= int32(0) {
		v621 = v7
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v76 = v7
	v77 = v7
	goto L12
L12:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83+v76<<(uint(int32(2))%32))))
	if int32(1)<<(uint(v45)%32)&int32(174) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v621 = v602
	goto L9
L14:
	;
	v604 = v76 + int32(1)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v604 < v605 {
		v76 = v604
		v77 = v602
		goto L12
	} else {
		goto L165
	}
L15:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+8)))
	if v88 != 0 {
		v602 = v77
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+9)))
	if v147 != int32(1) {
		v602 = v77
		goto L14
	} else {
		goto L34
	}
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+32))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v91 = int32(0)
	if v89 == v91 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v144 == int32(0) {
		v602 = v77
		goto L14
	} else {
		goto L33
	}
L20:
	;
	v144 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	if v90 == int32(0) {
		v137 = v91
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v144 = v137
	goto L19
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v101 < v100 {
		v137 = v91
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v103 = int32(1)
	if v100 <= v103 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v106 = v103
	goto L28
L27:
	;
	v106 = v100
	goto L28
L28:
	;
	v107 = int32(8)
	v112 = int32(0)
	goto L29
L29:
	;
	v119 = v112 << (uint(int32(2)) % 32)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v89+v107+v119)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v90+v107+v119)))
	v126 = v121 & (v123 ^ int32(-1))
	v128 = base.B2i32(v126 == int32(0))
	if v126 != 0 {
		v137 = v128
		goto L23
	} else {
		goto L31
	}
L30:
	;
	v137 = v128
	goto L23
L31:
	;
	v130 = v112 + int32(1)
	if v130 != v106 {
		v112 = v130
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	goto L17
L34:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v87)+96))
	if v150 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v87)+124))
	if v153 == int32(0) {
		v602 = v77
		goto L14
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v159 = int32(0)
	if v157 == v159 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	goto L37
L39:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v402 = F_op_strict(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L102
	} else {
		goto L103
	}
L40:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v279 = int32(0)
	if v277 == v279 {
		goto L72
	} else {
		goto L73
	}
L41:
	;
	if v212 == int32(0) {
		goto L40
	} else {
		goto L55
	}
L42:
	;
	v212 = int32(1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	if v158 == int32(0) {
		v205 = v159
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v212 = v205
	goto L41
L46:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v169 < v168 {
		v205 = v159
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v171 = int32(1)
	if v168 <= v171 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v174 = v171
	goto L50
L49:
	;
	v174 = v168
	goto L50
L50:
	;
	v175 = int32(8)
	v180 = int32(0)
	goto L51
L51:
	;
	v187 = v180 << (uint(int32(2)) % 32)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v157+v175+v187)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v158+v175+v187)))
	v194 = v189 & (v191 ^ int32(-1))
	v196 = base.B2i32(v194 == int32(0))
	if v194 != 0 {
		v205 = v196
		goto L45
	} else {
		goto L53
	}
L52:
	;
	v205 = v196
	goto L45
L53:
	;
	v198 = v180 + int32(1)
	if v198 != v174 {
		v180 = v198
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v87)+48))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v217 = int32(0)
	if v215 == v217 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v270 == int32(0) {
		goto L40
	} else {
		goto L70
	}
L57:
	;
	v270 = int32(1)
	goto L56
L58:
	;
	goto L59
L59:
	;
	if v216 == int32(0) {
		v263 = v217
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v270 = v263
	goto L56
L61:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v227 < v226 {
		v263 = v217
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v229 = int32(1)
	if v226 <= v229 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v232 = v229
	goto L65
L64:
	;
	v232 = v226
	goto L65
L65:
	;
	v233 = int32(8)
	v238 = int32(0)
	goto L66
L66:
	;
	v245 = v238 << (uint(int32(2)) % 32)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v215+v233+v245)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v216+v233+v245)))
	v252 = v247 & (v249 ^ int32(-1))
	v254 = base.B2i32(v252 == int32(0))
	if v252 != 0 {
		v263 = v254
		goto L60
	} else {
		goto L68
	}
L67:
	;
	v263 = v254
	goto L60
L68:
	;
	v256 = v238 + int32(1)
	if v256 != v232 {
		v238 = v256
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v156)+28))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v397 = v274 + int32(4)
	v398 = v274
	goto L39
L71:
	;
	if v332 == int32(0) {
		v602 = v77
		goto L14
	} else {
		goto L85
	}
L72:
	;
	v332 = int32(1)
	goto L71
L73:
	;
	goto L74
L74:
	;
	if v278 == int32(0) {
		v325 = v279
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v332 = v325
	goto L71
L76:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v289 < v288 {
		v325 = v279
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v291 = int32(1)
	if v288 <= v291 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v294 = v291
	goto L80
L79:
	;
	v294 = v288
	goto L80
L80:
	;
	v295 = int32(8)
	v300 = int32(0)
	goto L81
L81:
	;
	v307 = v300 << (uint(int32(2)) % 32)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v277+v295+v307)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v278+v295+v307)))
	v314 = v309 & (v311 ^ int32(-1))
	v316 = base.B2i32(v314 == int32(0))
	if v314 != 0 {
		v325 = v316
		goto L75
	} else {
		goto L83
	}
L82:
	;
	v325 = v316
	goto L75
L83:
	;
	v318 = v300 + int32(1)
	if v318 != v294 {
		v300 = v318
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v87)+48))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v337 = int32(0)
	if v335 == v337 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v390 == int32(0) {
		v602 = v77
		goto L14
	} else {
		goto L100
	}
L87:
	;
	v390 = int32(1)
	goto L86
L88:
	;
	goto L89
L89:
	;
	if v336 == int32(0) {
		v383 = v337
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v390 = v383
	goto L86
L91:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
	if v347 < v346 {
		v383 = v337
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v349 = int32(1)
	if v346 <= v349 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v352 = v349
	goto L95
L94:
	;
	v352 = v346
	goto L95
L95:
	;
	v353 = int32(8)
	v358 = int32(0)
	goto L96
L96:
	;
	v365 = v358 << (uint(int32(2)) % 32)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v335+v353+v365)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v336+v353+v365)))
	v372 = v367 & (v369 ^ int32(-1))
	v374 = base.B2i32(v372 == int32(0))
	if v372 != 0 {
		v383 = v374
		goto L90
	} else {
		goto L98
	}
L97:
	;
	v383 = v374
	goto L90
L98:
	;
	v376 = v358 + int32(1)
	if v376 != v352 {
		v358 = v376
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v156)+28))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	v397 = v394
	v398 = v394 + int32(4)
	goto L39
L101:
	;
	v515 = F_match_expr_to_partition_keys(m, v513, l2, v402)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L102
	} else {
		goto L137
	}
L102:
	;
	return
L103:
	;
	if v402 == int32(0) {
		v513 = v400
		v514 = v399
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v408 = int32(0)
	if base.B2i32(v406 == v408)|base.B2i32(v407 == v408) != 0 {
		v453 = v408
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v453 != 0 {
		goto L118
	} else {
		goto L119
	}
L106:
	;
	goto L105
L107:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v418 < v419 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v421 = v418
	goto L110
L109:
	;
	v421 = v419
	goto L110
L110:
	;
	if v421 <= int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v424 = int32(1)
	goto L113
L112:
	;
	v424 = v421
	goto L113
L113:
	;
	v425 = int32(8)
	v430 = int32(0)
	goto L114
L114:
	;
	v437 = v430 << (uint(int32(2)) % 32)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v407+v425+v437)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v406+v425+v437)))
	v442 = v439 & v441
	v444 = base.B2i32(v442 != int32(0))
	if v442 != 0 {
		v453 = v444
		goto L106
	} else {
		goto L116
	}
L115:
	;
	v453 = v444
	goto L106
L116:
	;
	v446 = v430 + int32(1)
	if v446 != v424 {
		v430 = v446
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v456 = F_remove_nulling_relids(m, v400, v454, int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L102
	} else {
		goto L121
	}
L119:
	;
	v458 = v400
	goto L120
L120:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v461 = int32(0)
	if base.B2i32(v459 == v461)|base.B2i32(v460 == v461) != 0 {
		v506 = v461
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v458 = v456
	goto L120
L122:
	;
	if v506 == int32(0) {
		v513 = v458
		v514 = v399
		goto L101
	} else {
		goto L135
	}
L123:
	;
	goto L122
L124:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	if v471 < v472 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v474 = v471
	goto L127
L126:
	;
	v474 = v472
	goto L127
L127:
	;
	if v474 <= int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v477 = int32(1)
	goto L130
L129:
	;
	v477 = v474
	goto L130
L130:
	;
	v478 = int32(8)
	v483 = int32(0)
	goto L131
L131:
	;
	v490 = v483 << (uint(int32(2)) % 32)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v460+v478+v490)))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v459+v478+v490)))
	v495 = v492 & v494
	v497 = base.B2i32(v495 != int32(0))
	if v495 != 0 {
		v506 = v497
		goto L123
	} else {
		goto L133
	}
L132:
	;
	v506 = v497
	goto L123
L133:
	;
	v499 = v483 + int32(1)
	if v499 != v477 {
		v483 = v499
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v511 = F_remove_nulling_relids(m, v399, v509, int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L102
	} else {
		goto L136
	}
L136:
	;
	v513 = v458
	v514 = v511
	goto L101
L137:
	;
	if v515 < int32(0) {
		v602 = v77
		goto L14
	} else {
		goto L138
	}
L138:
	;
	v519 = F_match_expr_to_partition_keys(m, v514, l3, v402)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L102
	} else {
		goto L139
	}
L139:
	;
	if v519 != v515 {
		v602 = v77
		goto L14
	} else {
		goto L140
	}
L140:
	;
	v524 = v23 + int32(32) + v515
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	if v525 != 0 {
		v602 = v77
		goto L14
	} else {
		goto L141
	}
L141:
	;
	v527 = v515 << (uint(int32(2)) % 32)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l2)+232))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+12))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v527+v529)))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v156)+24))
	if v531 != v532 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v534 == int32(104) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v591 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v591)
	v594 = v77 + v591
	v595 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	if v594 == v595 {
		goto L8
	} else {
		goto L164
	}
L144:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v87)+124))
	if v537 == int32(0) {
		v602 = v77
		goto L14
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v87)+96))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v546+v527)))
	v549 = int32(0)
	if v545 == v549 {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v540+v527)))
	v543 = F_op_in_opfamily(m, v537, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L102
	} else {
		goto L148
	}
L148:
	;
	if v543 != 0 {
		goto L143
	} else {
		goto L149
	}
L149:
	;
	v602 = v77
	goto L14
L150:
	;
	if v587 == int32(0) {
		v602 = v77
		goto L14
	} else {
		goto L163
	}
L151:
	;
	v587 = int32(0)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	if v555 <= int32(0) {
		v581 = v549
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v587 = v581
	goto L150
L155:
	;
	v558 = int32(0)
	if v558 < v555 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v561 = v555
	goto L158
L157:
	;
	v561 = v558
	goto L158
L158:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v545)+12))
	v564 = int32(0)
	goto L159
L159:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v562+v564<<(uint(int32(2))%32))))
	v573 = base.B2i32(v572 == v548)
	if v572 == v548 {
		v581 = v573
		goto L154
	} else {
		goto L161
	}
L160:
	;
	v581 = v573
	goto L154
L161:
	;
	v575 = v564 + int32(1)
	if v575 != v561 {
		v564 = v575
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	goto L143
L164:
	;
	v602 = v594
	goto L14
L165:
	;
	goto L13
L166:
	;
	v635 = v627
	v644 = v621
	v645 = v7
	goto L167
L167:
	;
	v652 = v23 + int32(32) + v645
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652))))
	if v653 == int32(1) {
		v801 = v635
		v810 = v644
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L2
L169:
	;
	v817 = v645 + int32(1)
	if v817 < v801 {
		v635 = v801
		v644 = v810
		v645 = v817
		goto L167
	} else {
		goto L199
	}
L170:
	;
	v657 = v645 << (uint(int32(2)) % 32)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v659 = v657 + v658
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v660 == int32(104) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v664+v657)))
	v668 = F_get_opfamily_member(m, v663, v666, v666, int32(1))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L102
	} else {
		goto L174
	}
L172:
	;
	v677 = v659
	goto L173
L173:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l2)+264))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v678+v657)))
	if v680 == int32(0) {
		goto L2
	} else {
		goto L178
	}
L174:
	;
	if v668 == int32(0) {
		goto L2
	} else {
		goto L175
	}
L175:
	;
	v672 = F_get_mergejoin_opfamilies(m, v668)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L102
	} else {
		goto L176
	}
L176:
	;
	if v672 == int32(0) {
		goto L2
	} else {
		goto L177
	}
L177:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v677 = v676
	goto L173
L178:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	if v683 <= int32(0) {
		goto L2
	} else {
		goto L179
	}
L179:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v677)))
	v698 = int32(0)
	goto L180
L180:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l2)+232))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v708)+12))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v709+v657)))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v680)+12))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v712+v698<<(uint(int32(2))%32))))
	v717 = F_exprCollation(m, v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L102
	} else {
		goto L182
	}
L181:
	;
	goto L2
L182:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l3)+264))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v719+v657)))
	if v721 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v793 = v698 + int32(1)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	if v793 < v794 {
		v698 = v793
		goto L180
	} else {
		goto L198
	}
L184:
	;
	v724 = int32(0)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v721)+4))
	if v725 <= v724 {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v733 = v724
	goto L186
L186:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v721)+12))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v748+v733<<(uint(int32(2))%32))))
	v753 = F_exprs_known_equal(m, l0, v716, v752, v686)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L102
	} else {
		goto L188
	}
L187:
	;
	v764 = F_exprCollation(m, v752)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L102
	} else {
		goto L196
	}
L188:
	;
	if v711 == v717 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v757 = v753
	goto L191
L190:
	;
	v757 = int32(0)
	goto L191
L191:
	;
	if v757 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v761 = v733 + int32(1)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v721)+4))
	if v761 < v762 {
		v733 = v761
		goto L186
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	goto L187
L195:
	;
	goto L183
L196:
	;
	v766 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v652))) = uint8(v766)
	v768 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	v770 = v644 + v766
	if v768 == v770 {
		goto L8
	} else {
		goto L197
	}
L197:
	;
	v801 = v768
	v810 = v770
	goto L169
L198:
	;
	goto L181
L199:
	;
	goto L168
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+264)) = v845
	v848 = F_palloc0(m, v844)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L102
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+268)) = v848
	if int32(0) < v842 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	if base.B2i32(int32(base.Ui32(int32(55))>>(uint(v841)%32))&int32(1) == int32(0))|base.B2i32(base.Ui32(int32(5)) < base.Ui32(v841)) != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v1088 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+217)) = uint8(v1088)
	goto L2
L205:
	;
	v873 = int32(0)
	goto L206
L206:
	;
	v884 = v873 << (uint(int32(2)) % 32)
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l3)+268))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v884+v885)))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l3)+264))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v888+v884)))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l2)+268))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v891+v884)))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l2)+264))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v894+v884)))
	switch v841 {
	case 0:
		goto L209
	case 1:
		goto L211
	default:
		goto L210
	case 4, 5:
		goto L212
	}
L207:
	;
	goto L204
L208:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v1059+v884))) = v1045
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l1)+268))
	*(*int32)(unsafe.Add(mBase, uint32(v1062+v884))) = v1043
	v1066 = v873 + int32(1)
	if v1066 != v842 {
		v873 = v1066
		goto L206
	} else {
		goto L243
	}
L209:
	;
	v1035 = F_list_concat_copy(m, v896, v890)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L102
	} else {
		goto L241
	}
L210:
	;
	v907 = F_list_concat_copy(m, v896, v890)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L102
	} else {
		goto L218
	}
L211:
	;
	v901 = F_list_copy(m, v896)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L102
	} else {
		goto L215
	}
L212:
	;
	v897 = F_list_copy(m, v896)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L102
	} else {
		goto L213
	}
L213:
	;
	v899 = F_list_copy(m, v893)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L102
	} else {
		goto L214
	}
L214:
	;
	v1043 = v899
	v1045 = v897
	goto L208
L215:
	;
	v903 = F_list_concat_copy(m, v890, v893)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L102
	} else {
		goto L216
	}
L216:
	;
	v905 = F_list_concat(m, v903, v887)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L102
	} else {
		goto L217
	}
L217:
	;
	v1043 = v905
	v1045 = v901
	goto L208
L218:
	;
	v909 = F_list_concat(m, v907, v893)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L102
	} else {
		goto L219
	}
L219:
	;
	v911 = F_list_concat(m, v909, v887)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L102
	} else {
		goto L220
	}
L220:
	;
	v913 = F_list_concat_copy(m, v896, v893)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L102
	} else {
		goto L221
	}
L221:
	;
	if v913 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1043 = v911
	v1045 = int32(0)
	goto L208
L223:
	;
	goto L224
L224:
	;
	v918 = int32(0)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v913)+4))
	if v920 <= v918 {
		v1043 = v911
		v1045 = v918
		goto L208
	} else {
		goto L225
	}
L225:
	;
	v927 = v911
	v930 = v918
	goto L226
L226:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v913)+12))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v943+v930<<(uint(int32(2))%32))))
	v948 = F_list_concat_copy(m, v890, v887)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L102
	} else {
		goto L229
	}
L227:
	;
	v1043 = v1015
	v1045 = v918
	goto L208
L228:
	;
	v1032 = v930 + int32(1)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v913)+4))
	if v1032 < v1033 {
		v927 = v1015
		v930 = v1032
		goto L226
	} else {
		goto L240
	}
L229:
	;
	if v948 == int32(0) {
		v1015 = v927
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v952 = int32(0)
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if v953 <= v952 {
		v1015 = v927
		goto L228
	} else {
		goto L231
	}
L231:
	;
	v956 = v952
	v960 = v927
	goto L232
L232:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v948)+12))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v976+v956<<(uint(int32(2))%32))))
	v982 = F_palloc0(m, int32(20))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L102
	} else {
		goto L234
	}
L233:
	;
	v1015 = v1005
	goto L228
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v982))) = int32(38)
	v986 = F_exprType(m, v947)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L102
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v982)+4)) = v986
	v989 = F_exprCollation(m, v947)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L102
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v982)+8)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v980
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v947
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v947
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v980
	v1000 = F_list_make2_impl(m, v23+int32(12), v23+int32(8))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L102
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v982)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v982)+12)) = v1000
	v1005 = F_lappend(m, v960, v982)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L102
	} else {
		goto L238
	}
L238:
	;
	v1008 = v956 + int32(1)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if v1008 < v1009 {
		v956 = v1008
		v960 = v1005
		goto L232
	} else {
		goto L239
	}
L239:
	;
	goto L233
L240:
	;
	goto L227
L241:
	;
	v1037 = F_list_concat_copy(m, v893, v887)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L102
	} else {
		goto L242
	}
L242:
	;
	v1043 = v1037
	v1045 = v1035
	goto L208
L243:
	;
	goto L207
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v841
	F_errmsg_internal(m, int32(_a_F_build_joinrel_partition_info_0), v23+int32(16))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L102
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_build_joinrel_partition_info_1), int32(2491), int32(_a_F_build_joinrel_partition_info_2))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L102
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_build_reloptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	v7 = int32(0)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_reloptions[0])))
	if v15 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_initialize_reloptions(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_build_reloptions[1]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	if l0 != 0 {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v103 = int32(0)
	v104 = v7
	goto L6
L8:
	;
	goto L9
L9:
	;
	v34 = v24
	v36 = v7
	v38 = v7
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v45 = v36 + base.B2i32(v41&l2 != int32(0))
	v47 = v38 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23+v47<<(uint(int32(2))%32))))
	if v51 != 0 {
		v34 = v51
		v36 = v45
		v38 = v47
		goto L10
	} else {
		goto L12
	}
L11:
	;
	if v45 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v103 = int32(0)
	v104 = v45
	goto L6
L14:
	;
	goto L15
L15:
	;
	v57 = F_palloc(m, v45<<(uint(int32(4))%32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_build_reloptions[1]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v61 == int32(0) {
		v103 = v57
		v104 = v45
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v71 = v61
	v74 = v7
	v75 = int32(0)
	goto L18
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v78&l2 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v103 = v57
	v104 = v45
	goto L6
L20:
	;
	v82 = v57 + v74<<(uint(int32(4))%32)
	v83 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v71
	v88 = v74 + int32(1)
	goto L22
L21:
	;
	v88 = v74
	goto L22
L22:
	;
	v91 = v75 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v60+v91<<(uint(int32(2))%32))))
	if v95 != 0 {
		v71 = v95
		v74 = v88
		v75 = v91
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	F_parseRelOptionsInternal(m, l0, l1, v103, v104)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v104 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	return int32(0)
L29:
	;
	goto L30
L30:
	;
	if int32(0) < v104 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v124 = int32(0)
	v127 = l3
	goto L34
L32:
	;
	v178 = l3
	goto L33
L33:
	;
	v182 = F_palloc0(m, v178)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L56
	}
L34:
	;
	v131 = int32(4)
	v133 = v103 + v124<<(uint(v131)%32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v135 == v131 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v178 = v163
	goto L33
L36:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+4)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+36))
	if v139 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v163 = v127
	goto L38
L38:
	;
	v167 = v124 + int32(1)
	if v167 != v104 {
		v124 = v167
		v127 = v163
		goto L34
	} else {
		goto L55
	}
L39:
	;
	v163 = v161 + v127
	goto L38
L40:
	;
	if v138&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v138&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v144 = m.T0[v139].(func(*base.Module, int32, int32) int32)(m, v142, int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+28)))
	if v146 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v161 = v144
	goto L39
L47:
	;
	v149 = int32(0)
	goto L49
L48:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	v149 = v148
	goto L49
L49:
	;
	v151 = m.T0[v139].(func(*base.Module, int32, int32) int32)(m, v149, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v161 = v151
	goto L39
L51:
	;
	v161 = v158 + int32(1)
	goto L39
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v156 = F_strlen(m, v155)
	mBase = m.M
	v158 = v156
	goto L51
L53:
	;
	goto L54
L54:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	v158 = v157
	goto L51
L55:
	;
	goto L35
L56:
	;
	F_fillRelOptions(m, v182, l3, v103, v104, l1, l4, l5)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v103)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	return v182
}
func F_buildint2vector(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13859(m, l0, l1, int32(21), int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_byteane(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_toast_raw_datum_size(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v109
L2:
	;
	return int32(0)
L3:
	;
	v13 = F_toast_raw_datum_size(m, v6)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v9 != v13 {
		v109 = int32(1)
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v16 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v18 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v20 = int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v22&v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v20
	goto L10
L9:
	;
	v25 = int32(4)
	goto L10
L10:
	;
	v26 = v16 + v25
	v27 = int32(1)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v29&v27 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = v27
	goto L13
L12:
	;
	v32 = int32(4)
	goto L13
L13:
	;
	v33 = v18 + v32
	v34 = int32(4)
	v35 = v9 - v34
	if base.Ui32(v34) <= base.Ui32(v35) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v98 != v16 {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v97 = int32(0)
	goto L14
L16:
	;
	v71 = v66
	v72 = v67
	v73 = v68
	goto L26
L17:
	;
	if (v26|v33)&int32(3) != 0 {
		v66 = v26
		v67 = v33
		v68 = v35
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v59 = v26
	v60 = v33
	v61 = v35
	goto L19
L19:
	;
	if v61 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v43 = v26
	v44 = v33
	v45 = v35
	goto L21
L21:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v48 != v49 {
		v66 = v43
		v67 = v44
		v68 = v45
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v59 = v54
	v60 = v52
	v61 = v56
	goto L19
L23:
	;
	v51 = int32(4)
	v52 = v44 + v51
	v54 = v43 + v51
	v56 = v45 - v51
	if base.Ui32(int32(3)) < base.Ui32(v56) {
		v43 = v54
		v44 = v52
		v45 = v56
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v66 = v59
	v67 = v60
	v68 = v61
	goto L16
L26:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v76 == v77 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v97 = v76 - v77
	goto L14
L28:
	;
	v79 = int32(1)
	v84 = v73 - v79
	if v84 != 0 {
		v71 = v71 + v79
		v72 = v72 + v79
		v73 = v84
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	F_pfree(m, v16)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v103 = base.B2i32(v97 != int32(0))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v18 == v104 {
		v109 = v103
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	F_pfree(m, v18)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v109 = v103
	goto L1
}
func F_byteanlike(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = v7 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			v20 = v18 & int32(1)
			if v20 != 0 {
				v21 = v12
			} else {
				v21 = v7 + int32(4)
			}
			if v18 == int32(1) {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
				if v27 == int32(18) {
					v30 = int32(16)
				} else {
					v30 = int32(0)
				}
				if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v37 = int32(4)
				} else {
					v37 = v30
				}
				v48 = v37
			} else {
				v38 = int32(1)
				if v20 != 0 {
					v48 = int32(base.Ui32(v18)>>(uint(v38)%32)) - v38
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v49 = int32(1)
			v50 = v14 + v49
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v55 = v53 & v49
			if v55 != 0 {
				v56 = v50
			} else {
				v56 = v14 + int32(4)
			}
			if v53 == int32(1) {
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
				if v62 == int32(18) {
					v65 = int32(16)
				} else {
					v65 = int32(0)
				}
				if base.Ui32((v62-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v72 = int32(4)
				} else {
					v72 = v65
				}
				v83 = v72
			} else {
				v73 = int32(1)
				if v55 != 0 {
					v83 = int32(base.Ui32(v53)>>(uint(v73)%32)) - v73
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v83 = int32(base.Ui32(v77)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v85 = F_SB_MatchText(m, v21, v48, v56, v83, int32(0))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v85 != int32(1))
			}
		}
	}
}
func F_bytes_to_mpi(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 == v3 {
		v26 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v54 = int32(0)
	v57 = F_pgp_mpi_alloc(m, v52, v9+int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	if l1 == v26 {
		v52 = v3
		goto L1
	} else {
		goto L8
	}
L3:
	;
	v15 = v3
	goto L4
L4:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v15))))
	if v20 != 0 {
		v26 = v15
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v52 = v3
	goto L1
L6:
	;
	v22 = v15 + int32(1)
	if v22 != l1 {
		v15 = v22
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v35 = (l1 + (v26 ^ int32(-1))) << (uint(int32(3)) % 32)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v26))))
	if v37 == int32(0) {
		v52 = v35
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v52 = v35 + (int32(8)-base.I32_clz(v37<<(uint(int32(24))%32)))&int32(255)
	goto L1
L10:
	;
	m.G0 = v9 + int32(16)
	return v77
L11:
	;
	return int32(0)
L12:
	;
	if v57 < int32(0) {
		v77 = v54
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v64 != l1 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_px_debug(m, int32(_a_F_bytes_to_mpi_0), v9)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if l1 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v72 = F_pgp_mpi_free(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v77 = v54
	goto L10
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	base.MemoryCopy(m, v74, l0, l1)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v77 = v76
	goto L10
}
