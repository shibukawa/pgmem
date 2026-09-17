package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BarrierArriveAndWait(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_BarrierArriveAndWait_0), int32(132), int32(_a_F_BarrierArriveAndWait_1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = int32(1)
	v18 = v16 + v17
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = v20 + v17
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v23 == v18 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27
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
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	v40 = l0 + int32(24)
	F_ConditionVariablePrepareToSleep(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
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
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v48 != 0 {
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
	if v22 == v56 {
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
	if v22 != v58 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	F_ConditionVariableSleep(m, v40, l1)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L24
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v22
	goto L22
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	return base.B2i32(v58 != v22)
L24:
	;
	goto L11
}
func F_BarrierDetach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v3 != 0 {
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
					v21 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v21
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 + int32(1)
					F_ConditionVariableBroadcast(m, l0+int32(24))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
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
				v21 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v21
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 + int32(1)
				F_ConditionVariableBroadcast(m, l0+int32(24))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
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
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int64
	_ = v259
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
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
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
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
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
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
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
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
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
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
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
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 float64
	_ = v482
	var v483 int32
	_ = v483
	var v484 float64
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 float64
	_ = v527
	var v529 int32
	_ = v529
	var v533 float64
	_ = v533
	var v535 int32
	_ = v535
	var v553 float64
	_ = v553
	var v554 float64
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int64
	_ = v560
	var v562 int64
	_ = v562
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v711 int64
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v935 int32
	_ = v935
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1018 int64
	_ = v1018
	var v1026 int32
	_ = v1026
	var v1037 int32
	_ = v1037
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1077 int32
	_ = v1077
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1095 int32
	_ = v1095
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1158 int32
	_ = v1158
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int64
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1239 int64
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1265 int32
	_ = v1265
	var v1271 int32
	_ = v1271
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int64
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1314 int64
	_ = v1314
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1332 int64
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1413 int32
	_ = v1413
	var v1420 int32
	_ = v1420
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1523 int32
	_ = v1523
	var v1530 int32
	_ = v1530
	var v1533 int64
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1559 int32
	_ = v1559
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1601 int64
	_ = v1601
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1619 int64
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1653 int32
	_ = v1653
	var v1658 int32
	_ = v1658
	var v1661 int64
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1687 int32
	_ = v1687
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1714 int32
	_ = v1714
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1727 int32
	_ = v1727
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1803 int32
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1821 int32
	_ = v1821
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1911 float64
	_ = v1911
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
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1862 = F_smgr_bulk_get_buf(m, v1861)
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L4
	} else {
		goto L355
	}
L2:
	;
	v1743 = int32(0)
	v1746 = v1743
	v1747 = v1743
	v1748 = v1727
	goto L336
L3:
	;
	F_pfree(m, v1278)
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
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
	v1707 = m.ExcPending
	if v1707 != 0 {
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
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v91 <= int32(0) {
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
	*(*int64)(unsafe.Add(mBase, uint32(v58+int32(80))+232)) = int64(2)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v81 + v70
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v87 - v70
	goto L11
L14:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v411 != 0 {
		goto L97
	} else {
		goto L98
	}
L15:
	;
	v95 = v91 + int32(1)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	v98 = F_palloc0(m, int32(32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[3]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+72)) = v103 + int32(1)
	goto L17
L17:
	;
	v110 = F_CreateParallelContext(m, int32(_a_F_btbuild_1), int32(_a_F_btbuild_2), v91)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v96 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v114 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v118 = int32(_a_F_btbuild_3)
	goto L21
L21:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v121 = F_table_parallelscan_estimate(m, v120, v118)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v116 = F_RegisterSnapshot(m, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v118 = v116
	goto L21
L24:
	;
	v123 = F_add_size(m, int32(96), v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v110)+36))
	v130 = F_add_size(m, v125, (v123+int32(31))&int32(-32))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+36)) = v130
	v133 = F_tuplesort_estimate_shared(m, v95)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v110)+36))
	v139 = (v133 + int32(31)) & int32(-32)
	v140 = F_add_size(m, v135, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+36)) = v140
	v143 = int32(1)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	if v145 == v143 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v148 = F_add_size(m, v140, v139)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v152 = int32(2)
	goto L31
L31:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v110)+40))
	v154 = F_add_size(m, v153, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+36)) = v148
	v152 = int32(3)
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+40)) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v110)+36))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v160 = F_mul_size(m, int32(32), v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v166 = F_add_size(m, v157, (v160+int32(31))&int32(-32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+36)) = v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v110)+40))
	v171 = F_add_size(m, v169, int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+40)) = v171
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v110)+36))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v177 = F_mul_size(m, int32(128), v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v183 = F_add_size(m, v174, (v177+int32(31))&int32(-32))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+36)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v110)+40))
	v188 = F_add_size(m, v186, int32(1))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+40)) = v188
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	if v192 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v110)+36))
	v194 = F_strlen(m, v192)
	mBase = m.M
	v199 = F_add_size(m, v193, v194&int32(-32)+int32(32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	v209 = v143
	goto L42
L42:
	;
	F_InitializeParallelDSM(m, v110)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+36)) = v199
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v110)+40))
	v204 = F_add_size(m, v202, int32(1))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+40)) = v204
	v209 = v194 + int32(1)
	goto L42
L45:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v110)+44))
	if v212 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	switch v215 {
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
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	v228 = F_shm_toc_allocate(m, v227, v123)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L54
	}
L49:
	;
	F_DestroyParallelContext(m, v110)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	F_UnregisterSnapshot(m, v118)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[3]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v222)+72)) = v223 - int32(1)
	goto L53
L53:
	;
	goto L14
L54:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v234
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+8)) = uint8(v236)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+12)) = v95
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+10)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+9)) = uint8(v238)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v244 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v228)+16)) = v249
	v252 = v228 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = int64(-4294967296)
	goto L59
L56:
	;
	v249 = int64(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v244)+392))
	v249 = v248
	goto L55
L59:
	;
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+72)) = uint8(v257)
	v259 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v228)+64)) = v259
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+56)) = uint8(v257)
	*(*int64)(unsafe.Add(mBase, uint32(v228)+48)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v228)+36)) = v259
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	F_table_parallelscan_initialize(m, v267, v228+int32(96), v118)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	v273 = F_shm_toc_allocate(m, v272, v133)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v110)+44))
	F_tuplesort_initialize_shared(m, v273, v95, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	F_shm_toc_insert(m, v278, int64(-6917529027641081855), v228)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	F_shm_toc_insert(m, v282, int64(-6917529027641081854), v273)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	if v287 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	v291 = F_shm_toc_allocate(m, v290, v133)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	v300 = int32(0)
	goto L67
L67:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	if v302 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v110)+44))
	F_tuplesort_initialize_shared(m, v291, v95, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	F_shm_toc_insert(m, v296, int64(-6917529027641081853), v291)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v300 = v291
	goto L67
L71:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	v304 = F_shm_toc_allocate(m, v303, v209)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v317 = F_mul_size(m, int32(32), v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L79
	}
L74:
	;
	if v209 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	base.MemoryCopy(m, v304, v307, v209)
	goto L77
L76:
	;
	goto L77
L77:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	F_shm_toc_insert(m, v309, int64(-6917529027641081852), v304)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	goto L73
L79:
	;
	v319 = F_shm_toc_allocate(m, v314, v317)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	F_shm_toc_insert(m, v321, int64(-6917529027641081851), v319)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v328 = F_mul_size(m, int32(128), v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v330 = F_shm_toc_allocate(m, v325, v328)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	F_shm_toc_insert(m, v332, int64(-6917529027641081850), v330)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_LaunchParallelWorkers(m, v110)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v110
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+28)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v98)+24)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v98)+12)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v339 + int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	if v349 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F__bt_end_parallel(m, v98)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v98
	v356 = F_palloc0(m, int32(16))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L4
	} else {
		goto L90
	}
L89:
	;
	goto L14
L90:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+4)) = v359
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v362
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+12)) = uint8(v365)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+13)) = uint8(v367)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+8)))
	if v371 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v375 = F_palloc0(m, int32(16))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L94
	}
L92:
	;
	v384 = int32(0)
	v386 = v370
	goto L93
L93:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[5]))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v392 = base.I32_div_s(v390, v391)
	F__bt_parallel_scan_and_sort(m, v356, v384, v386, v387, v388, v392, int32(1))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L95
	}
L94:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v375)+4)) = v377
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	v380 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v375)+12)) = uint8(v380)
	*(*int32)(unsafe.Add(mBase, uint32(v375)+8)) = v379
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v384 = v375
	v386 = v383
	goto L93
L95:
	;
	F_WaitForParallelWorkersToAttach(m, v110)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	goto L14
L97:
	;
	v413 = F_palloc0(m, int32(12))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L100
	}
L98:
	;
	v423 = int32(0)
	goto L99
L99:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+16)))
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)))
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[5]))
	v428 = F_tuplesort_begin_index_btree(m, l0, l1, v424, v425, v427, v423)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L101
	}
L100:
	;
	v415 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v413))) = uint8(v415)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+4)) = v418
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+8)) = v421
	v423 = v413
	goto L99
L101:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = v428
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	if v432 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v436 = F_palloc0(m, int32(16))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v468 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L105:
	;
	v438 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+12)) = uint8(v438)
	*(*int32)(unsafe.Add(mBase, uint32(v436)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v436)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v436
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v444 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v446 = F_palloc0(m, int32(12))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L109
	}
L107:
	;
	v457 = v438
	v458 = v436
	goto L108
L108:
	;
	v459 = int32(0)
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[6]))
	v463 = F_tuplesort_begin_index_btree(m, l0, l1, v459, v459, v462, v457)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L110
	}
L109:
	;
	v448 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v446))) = uint8(v448)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v446)+4)) = v451
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v446)+8)) = v454
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v457 = v446
	v458 = v456
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v458))) = v463
	goto L104
L111:
	;
	v555 = int32(0)
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v557
	v560 = *(*int64)(unsafe.Add(mBase, _c_F_btbuild[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v560
	v562 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v562
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = base.I64_trunc_sat_f64_s(v554)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v562
	goto L129
L112:
	;
	v471 = int32(1)
	v472 = int32(0)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+140))
	v482 = m.T0[v481].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v471, v472, v471, v472, int32(-1), int32(238), v22+int32(16), v472)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v468)+8))
	v489 = v485 + int32(36)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	goto L116
L115:
	;
	v484 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v553 = v482
	v554 = v484
	goto L111
L116:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = int32(1)
	if v510 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)) = uint8(v525)
	v527 = *(*float64)(unsafe.Add(mBase, uint32(v485)+64))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+32)) = v527
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485)+72)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)) = uint8(v529)
	*(*int32)(unsafe.Add(mBase, uint32(v485)+36)) = int32(0)
	v533 = *(*float64)(unsafe.Add(mBase, uint32(v485)+48))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L126
	}
L118:
	;
	F_s_lock(m, v489, int32(_a_F_btbuild_4), int32(1664), int32(_a_F_btbuild_5))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v485)+40))
	if v490 != v518 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = int32(0)
	F_ConditionVariableSleep(m, v485+int32(24), int32(134217767))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
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
	v553 = v533
	v554 = v527
	goto L111
L127:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v746 == int32(0) {
		v757 = v555
		goto L144
	} else {
		goto L145
	}
L128:
	;
	goto L127
L129:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v582 == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v586&int32(1) == int32(0) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v591 = int32(_a_F_btbuild_0)
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v594 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v593 + v594
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	*(*int32)(unsafe.Add(mBase, uint32(v582))) = v597 + v594
	goto L133
L132:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v728 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v582))) = v727 + v728
	v731 = int32(_a_F_btbuild_0)
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v733 - v728
	goto L128
L133:
	;
	goto L135
L135:
	;
	goto L136
L136:
	;
	v692 = int32(0)
	v695 = v555
	goto L141
L141:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(80)+v695<<(uint(int32(2))%32))))
	v705 = int32(3)
	v711 = *(*int64)(unsafe.Add(mBase, uint32(v22+int32(48)+v695<<(uint(v705)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v582+int32(232)+v704<<(uint(v705)%32)))) = v711
	v713 = int32(1)
	v716 = v692 + v713
	if v716 != int32(3) {
		v692 = v716
		v695 = v695 + v713
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
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v763 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v763 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L145:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)))
	if v749 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v757 = v746
	goto L144
L147:
	;
	goto L148
L148:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	F_tuplesort_end(m, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_pfree(m, v746)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = int32(0)
	v757 = v555
	goto L144
L151:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	F_tuplesort_performsort(m, v796)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L4
	} else {
		goto L155
	}
L152:
	;
	goto L151
L153:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v767&int32(1) == int32(0) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v772 = int32(_a_F_btbuild_0)
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v775 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v774 + v775
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v763)))
	*(*int32)(unsafe.Add(mBase, uint32(v763))) = v778 + v775
	*(*int64)(unsafe.Add(mBase, uint32(v763+int32(80))+232)) = int64(3)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v763)))
	*(*int32)(unsafe.Add(mBase, uint32(v763))) = v786 + v775
	v792 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v792 - v775
	goto L152
L155:
	;
	if v757 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v803 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v803 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	goto L158
L158:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v758)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v839
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v758)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v841
	v843 = int32(0)
	v845 = F__bt_mkscankey(m, v841, v843)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L4
	} else {
		goto L164
	}
L159:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	F_tuplesort_performsort(m, v836)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L4
	} else {
		goto L163
	}
L160:
	;
	goto L159
L161:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v807&int32(1) == int32(0) {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v812 = int32(_a_F_btbuild_0)
	v814 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v815 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v814 + v815
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v803)))
	*(*int32)(unsafe.Add(mBase, uint32(v803))) = v818 + v815
	*(*int64)(unsafe.Add(mBase, uint32(v803+int32(80))+232)) = int64(4)
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v803)))
	*(*int32)(unsafe.Add(mBase, uint32(v803))) = v826 + v815
	v832 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v832 - v815
	goto L160
L163:
	;
	goto L158
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v845
	v849 = F__bt_allequalimage(m, v841, int32(1))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v845)+1)) = uint8(v849)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(1)
	v858 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v858 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v841)+52))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v841)+192))
	v893 = int32(*(*int16)(unsafe.Add(mBase, uint32(v892)+10)))
	v895 = F_smgr_bulk_start_rel(m, v841, int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L4
	} else {
		goto L170
	}
L167:
	;
	goto L166
L168:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v862&int32(1) == int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v867 = int32(_a_F_btbuild_0)
	v869 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v870 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v869 + v870
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	*(*int32)(unsafe.Add(mBase, uint32(v858))) = v873 + v870
	*(*int64)(unsafe.Add(mBase, uint32(v858+int32(80))+232)) = int64(5)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	*(*int32)(unsafe.Add(mBase, uint32(v858))) = v881 + v870
	v887 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v887 - v870
	goto L167
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v895
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845)+1)))
	if v898 != int32(1) {
		v907 = v843
		goto L176
	} else {
		goto L177
	}
L171:
	;
	F_pfree(m, v925)
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L4
	} else {
		goto L330
	}
L172:
	;
	v1588 = int32(0)
	v1591 = v910
	v1601 = v17
	goto L312
L173:
	;
	v1278 = F_palloc(m, int32(1676))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L4
	} else {
		goto L254
	}
L174:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v918 = F_tuplesort_getheaptuple(m, v917)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L4
	} else {
		goto L185
	}
L175:
	;
	if v757 == int32(0) {
		goto L173
	} else {
		goto L184
	}
L176:
	;
	if v757 != 0 {
		goto L174
	} else {
		goto L180
	}
L177:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758)+12)))
	if v901 != 0 {
		v907 = v843
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v841)+180))
	if v902 == int32(0) {
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902)+16)))
	v907 = v905
	goto L176
L180:
	;
	if v907 != 0 {
		goto L173
	} else {
		goto L181
	}
L181:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v910 = F_tuplesort_getheaptuple(m, v909)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	if v910 != 0 {
		goto L172
	} else {
		goto L183
	}
L183:
	;
	v1843 = int32(0)
	v1844 = int32(0)
	goto L1
L184:
	;
	goto L174
L185:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	v921 = F_tuplesort_getheaptuple(m, v920)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	v925 = F_palloc0(m, v893*int32(36))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	if int32(0) < v893 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v935 = int32(0)
	goto L191
L189:
	;
	goto L190
L190:
	;
	v1004 = v918
	v1005 = int32(0)
	v1008 = v921
	v1018 = v17
	goto L195
L191:
	;
	v953 = v925 + v935*int32(36)
	v955 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v953))) = v955
	v959 = v845 + int32(16) + v935*int32(48)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v959)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v953)+4)) = v960
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	v966 = int32(base.Ui32(v962)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v953)+9)) = uint8(v966)
	v968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v959)+4)))
	v969 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v953)+20)) = uint8(v969)
	*(*uint16)(unsafe.Add(mBase, uint32(v953)+10)) = uint16(v968)
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	F_PrepareSortSupportFromIndexRel(m, v841, int32(base.Ui32(v972&int32(16777216))>>(uint(int32(24))%32)), v953)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L4
	} else {
		goto L193
	}
L192:
	;
	goto L190
L193:
	;
	v980 = v935 + int32(1)
	if v980 != v893 {
		v935 = v980
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	if v1008 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	if v1005 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L198:
	;
	if v1004 == int32(0) {
		goto L171
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v1026 = int32(0)
	if v1004 == v1026 {
		v1158 = v1026
		goto L197
	} else {
		goto L202
	}
L201:
	;
	v1158 = int32(1)
	goto L197
L202:
	;
	if int32(0) < v893 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1037 = int32(1)
	goto L206
L204:
	;
	goto L205
L205:
	;
	v1127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1004)+2)))
	v1128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1004))))
	v1129 = int32(16)
	v1131 = v1127 | v1128<<(uint(v1129)%32)
	v1132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1008)+2)))
	v1133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1008))))
	v1136 = v1132 | v1133<<(uint(v1129)%32)
	if base.Ui32(v1131) < base.Ui32(v1136) {
		v1147 = int32(-1)
		goto L229
	} else {
		goto L230
	}
L206:
	;
	v1053 = v925 + v1037*int32(36)
	v1056 = F_index_getattr_2(m, v1004, v1037, v891, v22+int32(80))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L4
	} else {
		goto L208
	}
L207:
	;
	goto L205
L208:
	;
	v1060 = F_index_getattr_2(m, v1008, v1037, v891, v22+int32(95))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L4
	} else {
		goto L209
	}
L209:
	;
	v1062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+95)))
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+80)))
	if v1063 == int32(1) {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	if v1037 != v893 {
		v1037 = v1037 + int32(1)
		goto L206
	} else {
		goto L227
	}
L211:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1053-int32(20))))
	v1084 = m.T0[v1083].(func(*base.Module, int32, int32, int32) int32)(m, v1056, v1060, v1053-int32(36))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L4
	} else {
		goto L220
	}
L212:
	;
	v1158 = int32(1)
	goto L197
L213:
	;
	if v1062&int32(1) != 0 {
		goto L210
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	if v1062&int32(1) == int32(0) {
		goto L211
	} else {
		goto L218
	}
L216:
	;
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053-int32(27)))))
	if v1070 != 0 {
		goto L212
	} else {
		goto L217
	}
L217:
	;
	v1158 = v1026
	goto L197
L218:
	;
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053-int32(27)))))
	if v1077 != 0 {
		v1158 = v1026
		goto L197
	} else {
		goto L219
	}
L219:
	;
	goto L212
L220:
	;
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053-int32(28)))))
	if v1088 == int32(1) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if v1084 < int32(0) {
		v1158 = v1026
		goto L197
	} else {
		goto L224
	}
L222:
	;
	v1095 = v1084
	goto L223
L223:
	;
	if int32(0) < v1095 {
		v1158 = v1026
		goto L197
	} else {
		goto L225
	}
L224:
	;
	v1095 = int32(0) - v1084
	goto L223
L225:
	;
	if v1095 == int32(0) {
		goto L210
	} else {
		goto L226
	}
L226:
	;
	v1158 = int32(1)
	goto L197
L227:
	;
	goto L207
L228:
	;
	v1158 = base.B2i32(v1147 <= int32(0))
	goto L197
L229:
	;
	goto L228
L230:
	;
	if base.Ui32(v1136) < base.Ui32(v1131) {
		v1147 = int32(1)
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v1141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1004)+4)))
	v1142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1008)+4)))
	if base.Ui32(v1141) < base.Ui32(v1142) {
		v1147 = int32(-1)
		goto L229
	} else {
		goto L232
	}
L232:
	;
	v1147 = base.B2i32(base.Ui32(v1142) < base.Ui32(v1141))
	goto L229
L233:
	;
	v1172 = F_palloc0(m, int32(32))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L4
	} else {
		goto L236
	}
L234:
	;
	v1217 = v1005
	goto L235
L235:
	;
	if v1158 != 0 {
		goto L243
	} else {
		goto L244
	}
L236:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1175 = F_smgr_bulk_get_buf(m, v1174)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	F_PageInit(m, v1175, int32(_a_F_btbuild_6), int32(16))
	mBase = m.M
	goto L238
L238:
	;
	v1180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175)+16)))
	v1181 = v1175 + v1180
	*(*int64)(unsafe.Add(mBase, uint32(v1181)+8)) = int64(4294967296)
	v1184 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1181))) = v1184
	v1186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175)+12)))
	v1188 = v1186 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1175)+12)) = uint16(v1188)
	*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1175
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1192 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1191 + v1192
	*(*int64)(unsafe.Add(mBase, uint32(v1172)+16)) = v1184
	*(*uint16)(unsafe.Add(mBase, uint32(v1172)+12)) = uint16(v1192)
	*(*int32)(unsafe.Add(mBase, uint32(v1172)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1172)+4)) = v1191
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+180))
	if v1203 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+4))
	v1210 = base.I32_div_s(int32(_a_F_btbuild_7)-v1205<<(uint(int32(13))%32), int32(100))
	v1212 = v1210
	goto L241
L240:
	;
	v1212 = int32(819)
	goto L241
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1172)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1172)+24)) = v1212
	v1217 = v1172
	goto L235
L242:
	;
	v1239 = v1018 + int64(1)
	v1242 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1242 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L243:
	;
	F__bt_buildadd(m, v22+int32(48), v1217, v1004, int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L4
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	F__bt_buildadd(m, v22+int32(48), v1217, v1008, int32(0))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L4
	} else {
		goto L248
	}
L246:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v1225 = F_tuplesort_getheaptuple(m, v1224)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L4
	} else {
		goto L247
	}
L247:
	;
	v1235 = v1225
	v1236 = v1008
	goto L242
L248:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	v1233 = F_tuplesort_getheaptuple(m, v1232)
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L4
	} else {
		goto L249
	}
L249:
	;
	v1235 = v1004
	v1236 = v1233
	goto L242
L250:
	;
	v1004 = v1235
	v1005 = v1217
	v1008 = v1236
	v1018 = v1239
	goto L195
L251:
	;
	goto L250
L252:
	;
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1246&int32(1) == int32(0) {
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v1251 = int32(_a_F_btbuild_0)
	v1253 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1254 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1253 + v1254
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1242)))
	*(*int32)(unsafe.Add(mBase, uint32(v1242))) = v1257 + v1254
	*(*int64)(unsafe.Add(mBase, uint32(v1242+int32(96))+232)) = v1239
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1242)))
	*(*int32)(unsafe.Add(mBase, uint32(v1242))) = v1265 + v1254
	v1271 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1271 - v1254
	goto L251
L254:
	;
	v1280 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1278)+4)) = v1280
	v1282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1278))) = uint8(v1282)
	*(*int64)(unsafe.Add(mBase, uint32(v1278)+10)) = v1280
	*(*int64)(unsafe.Add(mBase, uint32(v1278)+20)) = v1280
	*(*int64)(unsafe.Add(mBase, uint32(v1278)+28)) = v1280
	*(*int64)(unsafe.Add(mBase, uint32(v1278)+36)) = v1280
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v1293 = F_tuplesort_getheaptuple(m, v1292)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	if v1293 == int32(0) {
		goto L3
	} else {
		goto L256
	}
L256:
	;
	v1301 = int32(0)
	v1304 = v1293
	v1314 = v17
	goto L257
L257:
	;
	if v1301 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L258:
	;
	F__bt_sort_dedup_finish_pending(m, v22+int32(48), v1530, v1278)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L4
	} else {
		goto L308
	}
L259:
	;
	v1533 = v1314 + int64(1)
	v1536 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1536 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L260:
	;
	v1465 = F_CopyIndexTuple(m, v1304)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L4
	} else {
		goto L291
	}
L261:
	;
	v1320 = F_palloc0(m, int32(32))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L4
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+12))
	v1372 = F__bt_keep_natts_fast(m, v1370, v1371, v1304)
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L4
	} else {
		goto L271
	}
L264:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1323 = F_smgr_bulk_get_buf(m, v1322)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	F_PageInit(m, v1323, int32(_a_F_btbuild_6), int32(16))
	mBase = m.M
	goto L266
L266:
	;
	v1328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1323)+16)))
	v1329 = v1323 + v1328
	*(*int64)(unsafe.Add(mBase, uint32(v1329)+8)) = int64(4294967296)
	v1332 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1329))) = v1332
	v1334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1323)+12)))
	v1336 = v1334 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1323)+12)) = uint16(v1336)
	*(*int32)(unsafe.Add(mBase, uint32(v1320))) = v1323
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1340 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1339 + v1340
	*(*int64)(unsafe.Add(mBase, uint32(v1320)+16)) = v1332
	*(*uint16)(unsafe.Add(mBase, uint32(v1320)+12)) = uint16(v1340)
	*(*int32)(unsafe.Add(mBase, uint32(v1320)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1320)+4)) = v1339
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+180))
	if v1351 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1351)+4))
	v1358 = base.I32_div_s(int32(_a_F_btbuild_7)-v1353<<(uint(int32(13))%32), int32(100))
	v1360 = v1358
	goto L269
L268:
	;
	v1360 = int32(819)
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1320)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1320)+24)) = v1360
	v1364 = int32(812)
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+8)) = v1364
	v1367 = F_palloc(m, v1364)
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+24)) = v1367
	v1464 = v1320
	goto L260
L271:
	;
	if v893 < v1372 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1380 = int32(1)
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304)+7)))
	if v1381&int32(32) == int32(0) {
		v1399 = v1380
		v1401 = v1304
		goto L276
	} else {
		goto L277
	}
L273:
	;
	goto L274
L274:
	;
	F__bt_sort_dedup_finish_pending(m, v22+int32(48), v1301, v1278)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L4
	} else {
		goto L289
	}
L275:
	;
	if base.Ui32((v1403+(v1404+v1399)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1402) {
		v1530 = v1301
		goto L259
	} else {
		goto L288
	}
L276:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+8))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+20))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+28))
	v1413 = base.B2i32(base.Ui32((v1403+(v1404+v1399)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1402))
	if v1413 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L277:
	;
	v1386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1304)+4)))
	if v1386&int32(_a_F_btbuild_6) == int32(0) {
		v1399 = v1380
		v1401 = v1304
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1304)+2)))
	v1394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1304))))
	v1399 = v1386 & int32(4095)
	v1401 = v1393 + (v1304 + v1394<<(uint(int32(16))%32))
	goto L276
L279:
	;
	goto L275
L280:
	;
	v1447 = v1278 + v1444
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1447)))
	*(*int32)(unsafe.Add(mBase, uint32(v1447))) = v1448 + v1446
	goto L279
L281:
	;
	if v1404 <= int32(50) {
		goto L279
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+32)) = v1420 + int32(1)
	v1425 = v1399 * int32(6)
	if v1425 != 0 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	v1444 = int32(4)
	v1446 = int32(1)
	goto L280
L285:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+24))
	base.MemoryCopy(m, v1426+v1404*int32(6), v1401, v1425)
	goto L287
L286:
	;
	goto L287
L287:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+28)) = v1431 + v1399
	v1435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1304)+6)))
	v1444 = int32(36)
	v1446 = (v1435&int32(_a_F_btbuild_8)+int32(7))&int32(_a_F_btbuild_9) | int32(4)
	goto L280
L288:
	;
	goto L274
L289:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+12))
	F_pfree(m, v1459)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L4
	} else {
		goto L290
	}
L290:
	;
	v1464 = v1301
	goto L260
L291:
	;
	v1467 = int32(0)
	v1470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1465)+7)))
	if v1470&int32(32) != 0 {
		goto L295
	} else {
		goto L296
	}
L292:
	;
	v1530 = v1464
	goto L259
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+20)) = v1507
	*(*uint16)(unsafe.Add(mBase, uint32(v1278)+16)) = uint16(v1467)
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+12)) = v1465
	v1513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1465)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+36)) = (v1513&int32(_a_F_btbuild_8)+int32(7))&int32(_a_F_btbuild_9) | int32(4)
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1278+v1523<<(uint(int32(2))%32))+44)) = uint16(v1467)
	goto L292
L294:
	;
	v1488 = v1473 & int32(4095)
	v1490 = v1488 * int32(6)
	if v1490 != 0 {
		goto L299
	} else {
		goto L300
	}
L295:
	;
	v1473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1465)+4)))
	if v1473&int32(_a_F_btbuild_6) != 0 {
		goto L294
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+24))
	v1478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1465)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1477)+4)) = uint16(v1478)
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1465)))
	*(*int32)(unsafe.Add(mBase, uint32(v1477))) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+28)) = int32(1)
	v1484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1465)+6)))
	v1507 = v1484 & int32(_a_F_btbuild_8)
	goto L293
L298:
	;
	goto L297
L299:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+24))
	v1492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1465)+2)))
	v1493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1465))))
	base.MemoryCopy(m, v1491, v1492+(v1465+v1493<<(uint(int32(16))%32)), v1490)
	goto L301
L300:
	;
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1278)+28)) = v1488
	v1500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1465)+2)))
	v1501 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1465))))
	v1507 = v1500 | v1501<<(uint(int32(16))%32)
	goto L293
L302:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v1570 = F_tuplesort_getheaptuple(m, v1569)
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L4
	} else {
		goto L306
	}
L303:
	;
	goto L302
L304:
	;
	v1540 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1540&int32(1) == int32(0) {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1545 = int32(_a_F_btbuild_0)
	v1547 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1548 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1547 + v1548
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	*(*int32)(unsafe.Add(mBase, uint32(v1536))) = v1551 + v1548
	*(*int64)(unsafe.Add(mBase, uint32(v1536+int32(96))+232)) = v1533
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	*(*int32)(unsafe.Add(mBase, uint32(v1536))) = v1559 + v1548
	v1565 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1565 - v1548
	goto L303
L306:
	;
	if v1570 != 0 {
		v1301 = v1530
		v1304 = v1570
		v1314 = v1533
		goto L257
	} else {
		goto L307
	}
L307:
	;
	goto L258
L308:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+12))
	F_pfree(m, v1576)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L4
	} else {
		goto L309
	}
L309:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+24))
	F_pfree(m, v1579)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L4
	} else {
		goto L310
	}
L310:
	;
	F_pfree(m, v1278)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L4
	} else {
		goto L311
	}
L311:
	;
	v1727 = v1530
	goto L2
L312:
	;
	if v1588 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	v1727 = v1653
	goto L2
L314:
	;
	v1607 = F_palloc0(m, int32(32))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L4
	} else {
		goto L317
	}
L315:
	;
	v1653 = v1588
	goto L316
L316:
	;
	F__bt_buildadd(m, v22+int32(48), v1653, v1591, int32(0))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L4
	} else {
		goto L323
	}
L317:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1610 = F_smgr_bulk_get_buf(m, v1609)
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L4
	} else {
		goto L318
	}
L318:
	;
	F_PageInit(m, v1610, int32(_a_F_btbuild_6), int32(16))
	mBase = m.M
	goto L319
L319:
	;
	v1615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1610)+16)))
	v1616 = v1610 + v1615
	*(*int64)(unsafe.Add(mBase, uint32(v1616)+8)) = int64(4294967296)
	v1619 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1616))) = v1619
	v1621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1610)+12)))
	v1623 = v1621 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1610)+12)) = uint16(v1623)
	*(*int32)(unsafe.Add(mBase, uint32(v1607))) = v1610
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1627 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1626 + v1627
	*(*int64)(unsafe.Add(mBase, uint32(v1607)+16)) = v1619
	*(*uint16)(unsafe.Add(mBase, uint32(v1607)+12)) = uint16(v1627)
	*(*int32)(unsafe.Add(mBase, uint32(v1607)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1607)+4)) = v1626
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1637)+180))
	if v1638 != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	v1645 = base.I32_div_s(int32(_a_F_btbuild_7)-v1640<<(uint(int32(13))%32), int32(100))
	v1647 = v1645
	goto L322
L321:
	;
	v1647 = int32(819)
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1607)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1607)+24)) = v1647
	v1653 = v1607
	goto L316
L323:
	;
	v1661 = v1601 + int64(1)
	v1664 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1664 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v1698 = F_tuplesort_getheaptuple(m, v1697)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L4
	} else {
		goto L328
	}
L325:
	;
	goto L324
L326:
	;
	v1668 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1668&int32(1) == int32(0) {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1673 = int32(_a_F_btbuild_0)
	v1675 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1676 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1675 + v1676
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1664)))
	*(*int32)(unsafe.Add(mBase, uint32(v1664))) = v1679 + v1676
	*(*int64)(unsafe.Add(mBase, uint32(v1664+int32(96))+232)) = v1661
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1664)))
	*(*int32)(unsafe.Add(mBase, uint32(v1664))) = v1687 + v1676
	v1693 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1693 - v1676
	goto L325
L328:
	;
	if v1698 != 0 {
		v1588 = v1653
		v1591 = v1698
		v1601 = v1661
		goto L312
	} else {
		goto L329
	}
L329:
	;
	goto L313
L330:
	;
	if v1005 != 0 {
		v1727 = v1005
		goto L2
	} else {
		goto L331
	}
L331:
	;
	v1702 = int32(0)
	v1843 = v1702
	v1844 = v1702
	goto L1
L332:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v1708 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btbuild_10), v22)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L4
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(_a_F_btbuild_4), int32(321), int32(_a_F_btbuild_11))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
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
	v1722 = int32(0)
	v1843 = v1722
	v1844 = v1722
	goto L1
L336:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+4))
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+28))
	if v1765 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	v1843 = v1792
	v1844 = v1793
	goto L1
L338:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1748)))
	v1795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1794)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1795) {
		goto L345
	} else {
		goto L346
	}
L339:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1748)))
	v1769 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1768)+16)))
	v1770 = v1768 + v1769
	v1771 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1770)+12)))
	v1773 = v1771 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1770)+12)) = uint16(v1773)
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+20))
	v1792 = v1775
	v1793 = v1764
	goto L338
L340:
	;
	goto L341
L341:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1776))) = base.I32_rotr(v1764, int32(16))
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+28))
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+8))
	F__bt_buildadd(m, v22+int32(48), v1782, v1783, int32(0))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L4
	} else {
		goto L342
	}
L342:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+8))
	F_pfree(m, v1787)
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L4
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1748)+8)) = int32(0)
	v1792 = v1746
	v1793 = v1747
	goto L338
L344:
	;
	v1831 = v1795 - int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1794)+12)) = uint16(v1831)
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+4))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1748)))
	F_smgr_bulk_write(m, v1833, v1834, v1835, int32(1))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L4
	} else {
		goto L353
	}
L345:
	;
	v1803 = int32(base.Ui32(v1795+int32(_a_F_btbuild_12)) >> (uint(int32(2)) % 32))
	goto L347
L346:
	;
	v1803 = int32(0)
	goto L347
L347:
	;
	if base.Ui32(v1803&int32(_a_F_btbuild_13)) < base.Ui32(int32(2)) {
		goto L344
	} else {
		goto L348
	}
L348:
	;
	v1808 = int32(3)
	v1812 = (v1803 + int32(1)) & int32(_a_F_btbuild_13)
	if base.Ui32(v1812) <= base.Ui32(v1808) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1815 = v1808
	goto L351
L350:
	;
	v1815 = v1812
	goto L351
L351:
	;
	v1816 = int32(2)
	v1821 = (v1815 - v1816) & int32(_a_F_btbuild_13) << (uint(v1816) % 32)
	if v1821 == int32(0) {
		goto L344
	} else {
		goto L352
	}
L352:
	;
	base.MemoryCopy(m, v1794+int32(24), v1794+int32(28), v1821)
	goto L344
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1748))) = int32(0)
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+28))
	if v1841 != 0 {
		v1746 = v1792
		v1747 = v1793
		v1748 = v1841
		goto L336
	} else {
		goto L354
	}
L354:
	;
	goto L337
L355:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	v1865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1864)+1)))
	F_PageInit(m, v1862, int32(_a_F_btbuild_6), int32(16))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1862)+64)) = uint8(v1865)
	*(*int64)(unsafe.Add(mBase, uint32(v1862)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v1862)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1862)+44)) = v1843
	*(*int32)(unsafe.Add(mBase, uint32(v1862)+40)) = v1844
	*(*int32)(unsafe.Add(mBase, uint32(v1862)+36)) = v1843
	*(*int32)(unsafe.Add(mBase, uint32(v1862)+32)) = v1844
	*(*int64)(unsafe.Add(mBase, uint32(v1862)+24)) = int64(17180209506)
	v1880 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1862)+16)))
	v1882 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v1862+v1880)+12)) = uint16(v1882)
	v1884 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v1862)+12)) = uint16(v1884)
	goto L356
L356:
	;
	F_smgr_bulk_write(m, v1861, int32(0), v1862, int32(1))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L4
	} else {
		goto L357
	}
L357:
	;
	F_smgr_bulk_finish(m, v1861)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L4
	} else {
		goto L358
	}
L358:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1892)))
	F_tuplesort_end(m, v1893)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L4
	} else {
		goto L359
	}
L359:
	;
	F_pfree(m, v1892)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L4
	} else {
		goto L360
	}
L360:
	;
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v1898 != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v1898)))
	F_tuplesort_end(m, v1899)
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L4
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v1904 != 0 {
		goto L366
	} else {
		goto L367
	}
L364:
	;
	F_pfree(m, v1898)
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L4
	} else {
		goto L365
	}
L365:
	;
	goto L363
L366:
	;
	F__bt_end_parallel(m, v1904)
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L4
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1908 = F_palloc(m, int32(16))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L4
	} else {
		goto L370
	}
L369:
	;
	goto L368
L370:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1908))) = v553
	v1911 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1908)+8)) = v1911
	m.G0 = v22 + int32(96)
	return v1908
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
	var v15 int32
	_ = v15
	v3 = l0 + int32(12)
	v4 = int32(70)
	*(*uint16)(unsafe.Add(mBase, uint32(v3))) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(-1)
	v15 = l0 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(-4294967296)
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
	v3 = Fn13854(m, l0, int32(19))
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
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int64
	_ = v405
	var v407 int64
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v485 float64
	_ = v485
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v614 int32
	_ = v614
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v692 int32
	_ = v692
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v814 int32
	_ = v814
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v868 float64
	_ = v868
	var v870 float64
	_ = v870
	var v871 int32
	_ = v871
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1028 int32
	_ = v1028
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1214 int32
	_ = v1214
	var v1226 int32
	_ = v1226
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1279 int32
	_ = v1279
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	var v1348 int32
	_ = v1348
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1371 int32
	_ = v1371
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1385 int32
	_ = v1385
	var v1397 int32
	_ = v1397
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1466 int32
	_ = v1466
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1481 int64
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1494 int32
	_ = v1494
	var v1508 int32
	_ = v1508
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1610 float64
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1622 int32
	_ = v1622
	var v1627 int32
	_ = v1627
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1716 int32
	_ = v1716
	var v1755 float64
	_ = v1755
	var v1756 float64
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1791 int32
	_ = v1791
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1830 int32
	_ = v1830
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1856 int32
	_ = v1856
	var v1866 int32
	_ = v1866
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1907 int32
	_ = v1907
	var v1909 int32
	_ = v1909
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1944 int32
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1997 int32
	_ = v1997
	var v2002 int32
	_ = v2002
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2016 int32
	_ = v2016
	var v2029 int32
	_ = v2029
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2046 int32
	_ = v2046
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2080 int32
	_ = v2080
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2091 int32
	_ = v2091
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2110 int32
	_ = v2110
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2175 int32
	_ = v2175
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2199 int32
	_ = v2199
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2242 int32
	_ = v2242
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2273 int32
	_ = v2273
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2296 int32
	_ = v2296
	var v2302 int32
	_ = v2302
	var v2304 int32
	_ = v2304
	var v2310 int32
	_ = v2310
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2325 int32
	_ = v2325
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2340 int32
	_ = v2340
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2371 int32
	_ = v2371
	var v2376 int32
	_ = v2376
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2389 int32
	_ = v2389
	var v2396 int32
	_ = v2396
	var v2402 int32
	_ = v2402
	var v2404 int32
	_ = v2404
	var v2410 int32
	_ = v2410
	var v2414 int32
	_ = v2414
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2445 int32
	_ = v2445
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2454 int32
	_ = v2454
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2467 int32
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2484 int32
	_ = v2484
	var v2488 int32
	_ = v2488
	var v2492 int32
	_ = v2492
	var v2496 int32
	_ = v2496
	var v2502 int32
	_ = v2502
	var v2504 int32
	_ = v2504
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2515 int32
	_ = v2515
	var v2521 int32
	_ = v2521
	var v2524 int64
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2529 int32
	_ = v2529
	var v2535 int32
	_ = v2535
	var v2537 int32
	_ = v2537
	var v2543 int32
	_ = v2543
	var v2544 int64
	_ = v2544
	var v2554 int32
	_ = v2554
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2568 int32
	_ = v2568
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2593 int32
	_ = v2593
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2697 int32
	_ = v2697
	var v2703 int32
	_ = v2703
	var v2705 int32
	_ = v2705
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2717 int32
	_ = v2717
	var v2721 int32
	_ = v2721
	var v2723 int32
	_ = v2723
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2740 int32
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2745 int32
	_ = v2745
	var v2747 int32
	_ = v2747
	var v2749 int32
	_ = v2749
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2757 int32
	_ = v2757
	var v2759 int32
	_ = v2759
	var v2763 int32
	_ = v2763
	var v2769 int32
	_ = v2769
	var v2771 int32
	_ = v2771
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2784 int32
	_ = v2784
	var v2788 int32
	_ = v2788
	var v2790 int32
	_ = v2790
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2808 int32
	_ = v2808
	var v2814 int32
	_ = v2814
	var v2816 int32
	_ = v2816
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2825 int32
	_ = v2825
	var v2826 int32
	_ = v2826
	var v2828 int32
	_ = v2828
	var v2836 int32
	_ = v2836
	var __phi2836 int32
	_ = __phi2836
	var v2837 int32
	_ = v2837
	var __phi2837 int32
	_ = __phi2837
	var v2839 int32
	_ = v2839
	var __phi2839 int32
	_ = __phi2839
	var v2846 int32
	_ = v2846
	var __phi2846 int32
	_ = __phi2846
	var v2883 int32
	_ = v2883
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2919 int32
	_ = v2919
	var v2924 int32
	_ = v2924
	var v2927 int32
	_ = v2927
	var v2930 int32
	_ = v2930
	var v2933 int32
	_ = v2933
	var v2935 int32
	_ = v2935
	var v2937 int32
	_ = v2937
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2944 int32
	_ = v2944
	var v2946 int32
	_ = v2946
	var v2950 int32
	_ = v2950
	var v2956 int32
	_ = v2956
	var v2958 int32
	_ = v2958
	var v2964 int32
	_ = v2964
	var v2965 int32
	_ = v2965
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2970 int32
	_ = v2970
	var v2979 int32
	_ = v2979
	var v2981 int32
	_ = v2981
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3031 int32
	_ = v3031
	var v3037 int32
	_ = v3037
	var v3039 int32
	_ = v3039
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3051 int32
	_ = v3051
	var v3054 int32
	_ = v3054
	var v3056 int32
	_ = v3056
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3084 int32
	_ = v3084
	var v3089 int32
	_ = v3089
	var v3101 int32
	_ = v3101
	var v3104 int32
	_ = v3104
	var v3105 int32
	_ = v3105
	var v3108 int32
	_ = v3108
	var v3109 int32
	_ = v3109
	var v3111 int32
	_ = v3111
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3118 int32
	_ = v3118
	var v3120 int32
	_ = v3120
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3126 int32
	_ = v3126
	var v3132 int32
	_ = v3132
	var v3134 int32
	_ = v3134
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3143 int32
	_ = v3143
	var v3147 int32
	_ = v3147
	var v3148 int32
	_ = v3148
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3167 int32
	_ = v3167
	var v3172 int32
	_ = v3172
	var v3177 int32
	_ = v3177
	var v3179 int32
	_ = v3179
	var v3182 int32
	_ = v3182
	var v3184 int32
	_ = v3184
	var v3187 int32
	_ = v3187
	var v3189 int32
	_ = v3189
	var v3192 int32
	_ = v3192
	var v3193 int32
	_ = v3193
	var v3194 int32
	_ = v3194
	var v3197 int32
	_ = v3197
	var v3202 int32
	_ = v3202
	var v3208 int32
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3225 int32
	_ = v3225
	var v3227 int32
	_ = v3227
	var v3231 int32
	_ = v3231
	var v3237 int32
	_ = v3237
	var v3239 int32
	_ = v3239
	var v3245 int32
	_ = v3245
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3254 int32
	_ = v3254
	var v3256 int32
	_ = v3256
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3261 int32
	_ = v3261
	var v3263 int32
	_ = v3263
	var v3270 int32
	_ = v3270
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3292 int32
	_ = v3292
	var v3298 int32
	_ = v3298
	var v3300 int32
	_ = v3300
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3312 int32
	_ = v3312
	var v3317 int32
	_ = v3317
	var v3319 int32
	_ = v3319
	var v3324 int32
	_ = v3324
	var v3330 int32
	_ = v3330
	var v3332 int32
	_ = v3332
	var v3338 int32
	_ = v3338
	var v3339 int32
	_ = v3339
	var v3340 int64
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3344 int32
	_ = v3344
	var v3348 int32
	_ = v3348
	var v3350 int32
	_ = v3350
	var v3353 int32
	_ = v3353
	var v3356 int32
	_ = v3356
	var v3358 int32
	_ = v3358
	var v3361 int32
	_ = v3361
	var v3369 int32
	_ = v3369
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3378 int32
	_ = v3378
	var v3380 int32
	_ = v3380
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3386 int32
	_ = v3386
	var v3390 int32
	_ = v3390
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3396 int32
	_ = v3396
	var v3400 int32
	_ = v3400
	var v3404 int32
	_ = v3404
	var v3408 int32
	_ = v3408
	var v3414 int32
	_ = v3414
	var v3426 int32
	_ = v3426
	var v3431 int64
	_ = v3431
	var v3432 int32
	_ = v3432
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3439 int32
	_ = v3439
	var v3441 int32
	_ = v3441
	var v3443 int32
	_ = v3443
	var v3445 int32
	_ = v3445
	var v3447 int32
	_ = v3447
	var v3449 int32
	_ = v3449
	var v3456 int32
	_ = v3456
	var v3459 int64
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3464 int64
	_ = v3464
	var v3468 int32
	_ = v3468
	var v3474 int32
	_ = v3474
	var v3476 int32
	_ = v3476
	var v3482 int32
	_ = v3482
	var v3483 int64
	_ = v3483
	var v3488 int32
	_ = v3488
	var v3489 int32
	_ = v3489
	var v3493 int32
	_ = v3493
	var v3499 int32
	_ = v3499
	var v3501 int32
	_ = v3501
	var v3507 int32
	_ = v3507
	var v3513 int32
	_ = v3513
	var v3519 int32
	_ = v3519
	var v3521 int32
	_ = v3521
	var v3527 int32
	_ = v3527
	var v3534 int32
	_ = v3534
	var v3538 int32
	_ = v3538
	var v3540 int32
	_ = v3540
	var v3544 int32
	_ = v3544
	var v3551 int32
	_ = v3551
	var v3553 int32
	_ = v3553
	var v3559 int32
	_ = v3559
	var v3561 int32
	_ = v3561
	var v3564 int32
	_ = v3564
	var v3566 int32
	_ = v3566
	var v3569 int32
	_ = v3569
	var v3571 int32
	_ = v3571
	var v3576 int32
	_ = v3576
	var v3578 int32
	_ = v3578
	var v3579 int32
	_ = v3579
	var v3584 int32
	_ = v3584
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3591 int32
	_ = v3591
	var v3593 int32
	_ = v3593
	var v3595 int32
	_ = v3595
	var v3597 int32
	_ = v3597
	var v3599 int32
	_ = v3599
	var v3602 int32
	_ = v3602
	var v3603 int32
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3608 int32
	_ = v3608
	var v3609 int32
	_ = v3609
	var v3610 int32
	_ = v3610
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3620 int32
	_ = v3620
	var v3628 int32
	_ = v3628
	var v3633 int32
	_ = v3633
	var v3639 int32
	_ = v3639
	var v3652 int32
	_ = v3652
	var v3694 int32
	_ = v3694
	var v3697 int32
	_ = v3697
	var v3699 int32
	_ = v3699
	var v3701 int32
	_ = v3701
	var v3703 int32
	_ = v3703
	var v3706 int32
	_ = v3706
	var v3707 int32
	_ = v3707
	var v3710 int32
	_ = v3710
	var v3712 int32
	_ = v3712
	var v3716 int32
	_ = v3716
	var v3720 int32
	_ = v3720
	var v3725 int32
	_ = v3725
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3740 int32
	_ = v3740
	var v3745 int32
	_ = v3745
	var v3749 int32
	_ = v3749
	var v3752 int32
	_ = v3752
	var v3753 int32
	_ = v3753
	var v3754 int32
	_ = v3754
	var v3765 int32
	_ = v3765
	var v3770 int32
	_ = v3770
	var v3774 int32
	_ = v3774
	var v3775 int32
	_ = v3775
	var v3785 int32
	_ = v3785
	var v3790 int32
	_ = v3790
	var v3795 int32
	_ = v3795
	var v3843 int32
	_ = v3843
	var v3844 int32
	_ = v3844
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3857 int32
	_ = v3857
	var v3862 int32
	_ = v3862
	var v3915 int32
	_ = v3915
	var v3967 int32
	_ = v3967
	var v4052 int32
	_ = v4052
	var v4074 int32
	_ = v4074
	var v4104 int32
	_ = v4104
	var v4129 int32
	_ = v4129
	var v4130 int32
	_ = v4130
	var v4132 int32
	_ = v4132
	var v4133 int32
	_ = v4133
	var v4134 int32
	_ = v4134
	var v4136 int32
	_ = v4136
	var v4141 int32
	_ = v4141
	var v4142 int32
	_ = v4142
	var v4147 int32
	_ = v4147
	var v4148 int32
	_ = v4148
	var v4156 int32
	_ = v4156
	var v4161 int32
	_ = v4161
	var v4165 int32
	_ = v4165
	var v4166 int32
	_ = v4166
	var v4167 int32
	_ = v4167
	var v4178 int32
	_ = v4178
	var v4191 int32
	_ = v4191
	var v4197 int32
	_ = v4197
	var v4200 int32
	_ = v4200
	var v4216 int32
	_ = v4216
	var v4223 int32
	_ = v4223
	var v4227 int32
	_ = v4227
	var v4232 int32
	_ = v4232
	var v4234 int32
	_ = v4234
	var v4235 int32
	_ = v4235
	var v4238 int32
	_ = v4238
	var v4246 int32
	_ = v4246
	var v4252 int32
	_ = v4252
	var v4257 int32
	_ = v4257
	var v4259 int32
	_ = v4259
	var v4261 int32
	_ = v4261
	var v4262 int32
	_ = v4262
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4270 int32
	_ = v4270
	var v4271 int32
	_ = v4271
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4274 int32
	_ = v4274
	var v4279 int32
	_ = v4279
	var v4327 int32
	_ = v4327
	var v4330 int32
	_ = v4330
	var v4331 int32
	_ = v4331
	var v4332 int64
	_ = v4332
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4338 int32
	_ = v4338
	var v4339 int32
	_ = v4339
	var v4340 int32
	_ = v4340
	var v4344 int32
	_ = v4344
	var v4345 int32
	_ = v4345
	var v4397 int32
	_ = v4397
	var v4400 int32
	_ = v4400
	var v4449 int32
	_ = v4449
	var v4500 int32
	_ = v4500
	var v4502 int32
	_ = v4502
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
	v4257 = m.ExcPending
	if v4257 != 0 {
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
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	if base.Ui32(v230) < base.Ui32(v189) {
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
	*(*int64)(unsafe.Add(mBase, uint32(v197+int32(120))+232)) = base.I64_extend_i32_u(v189)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v220 + v209
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v226 - v209
	goto L31
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v189
	v233 = v125
	v234 = v126
	v245 = v137
	v258 = v150
	v264 = v156
	v267 = v159
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
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v287 = F_read_stream_next_buffer(m, v264, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v4216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4166)+9)))
	if v4216 != int32(1) {
		v233 = v4166
		v234 = v4167
		v245 = v4178
		v258 = v4191
		v264 = v4197
		v267 = v4200
		goto L37
	} else {
		goto L736
	}
L41:
	;
	F__bt_relbuf(m, v317)
	mBase = m.M
	v4165 = m.ExcPending
	if v4165 != 0 {
		goto L1
	} else {
		goto L735
	}
L42:
	;
	v4141 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4142 = m.ExcPending
	if v4142 != 0 {
		goto L1
	} else {
		goto L730
	}
L43:
	;
	if v287 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v245)+24))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v245)+36))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v245)+32))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v245)+28))
	if v287 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	F_read_stream_reset(m, v264)
	mBase = m.M
	v4136 = m.ExcPending
	if v4136 != 0 {
		goto L1
	} else {
		goto L729
	}
L47:
	;
	v314 = v233
	v315 = v234
	v317 = v287
	v323 = v291
	v324 = v313
	v326 = v245
	v337 = v294
	v339 = v258
	v341 = v313
	v345 = v264
	v346 = v293
	v348 = v267
	v349 = v289
	v355 = v290
	v356 = v292
	goto L51
L48:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v298+(v287^int32(-1))<<(uint(int32(6))%32))+16))
	v313 = v304
	goto L47
L49:
	;
	goto L50
L50:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306+v287<<(uint(int32(6))%32)+int32(-64))+16))
	v313 = v312
	goto L47
L51:
	;
	F_LockBuffer(m, v317, int32(1))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v317 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	if v4104 == int32(0) {
		v4166 = v314
		v4167 = v315
		v4178 = v326
		v4191 = v339
		v4197 = v345
		v4200 = v348
		goto L40
	} else {
		goto L726
	}
L55:
	;
	F__bt_relbuf(m, v317)
	mBase = m.M
	v4074 = m.ExcPending
	if v4074 != 0 {
		goto L1
	} else {
		goto L725
	}
L56:
	;
	v436 = int32(0)
	v438 = v411 & int32(_a_F_btvacuumscan_5)
	if v438&int32(16) == v436 {
		goto L84
	} else {
		goto L85
	}
L57:
	;
	v4052 = int32(0)
	goto L55
L58:
	;
	F_RecordFreeIndexPage(m, v323, v324)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L83
	}
L59:
	;
	v385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384)+14)))
	if v385 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v370+(v317^int32(-1))<<(uint(int32(2))%32))))
	v384 = v376
	goto L59
L61:
	;
	goto L62
L62:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v384 = v378 + v317<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	F__bt_checkpage(m, v323, v317)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v324 != v341 {
		goto L42
	} else {
		goto L82
	}
L66:
	;
	v388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384)+16)))
	v389 = v384 + v388
	v390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+12)))
	if v324 != v341 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v390&int32(17) != int32(1) {
		goto L42
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v390&int32(4) != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	if v390&int32(4) != 0 {
		goto L41
	} else {
		goto L71
	}
L71:
	;
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+14)))
	v399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326)+40)))
	if v398 != v399 {
		goto L41
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	if v390&int32(256) != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v411 = v390
	goto L75
L75:
	;
	if v411&int32(4) == int32(0) {
		goto L56
	} else {
		goto L81
	}
L76:
	;
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v384)+24))
	v407 = v405
	goto L78
L77:
	;
	v407 = int64(3)
	goto L78
L78:
	;
	v408 = F_GlobalVisCheckRemovableFullXid(m, v355, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v408 != 0 {
		goto L58
	} else {
		goto L80
	}
L80:
	;
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+12)))
	v411 = v410
	goto L75
L81:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v337)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v337)+28)) = v416 + int32(1)
	goto L57
L82:
	;
	goto L58
L83:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v337)+28))
	v426 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v337)+28)) = v425 + v426
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v337)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v337)+32)) = v429 + v426
	goto L57
L84:
	;
	if v438&int32(1) == int32(0) {
		v4052 = v436
		goto L55
	} else {
		goto L87
	}
L85:
	;
	v1791 = v436
	goto L86
L86:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v326)+44))
	F_MemoryContextReset(m, v1812)
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L1
	} else {
		goto L249
	}
L87:
	;
	F_LockBuffer(m, v317, int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_LockBufferForCleanup(m, v317)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v452 = int32(0)
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326)+40)))
	if v454 == v452 {
		v471 = v452
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v474 != 0 {
		goto L100
	} else {
		goto L101
	}
L91:
	;
	v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+14)))
	if v458 != v454 {
		v471 = int32(0)
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+12)))
	if v461&int32(32) != 0 {
		v471 = int32(0)
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if base.Ui32(v464) < base.Ui32(v341) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v467 = v464
	goto L96
L95:
	;
	v467 = int32(0)
	goto L96
L96:
	;
	if v464 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v469 = v467
	goto L99
L98:
	;
	v469 = int32(0)
	goto L99
L99:
	;
	v471 = v469
	goto L90
L100:
	;
	v475 = int32(2)
	goto L102
L101:
	;
	v475 = int32(1)
	goto L102
L102:
	;
	v476 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v476) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v484 = int32(base.Ui32(v476+int32(_a_F_btvacuumscan_6)) >> (uint(int32(2)) % 32))
	goto L105
L104:
	;
	v484 = int32(0)
	goto L105
L105:
	;
	v485 = float64(0)
	if v346 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v871 = int32(0)
	if base.B2i32(v826 <= v871)&base.B2i32(v827 <= v871) == v871 {
		goto L146
	} else {
		goto L147
	}
L107:
	;
	v826 = v452
	v827 = int32(0)
	v868 = v485
	v870 = float64(0)
	goto L106
L108:
	;
	goto L109
L109:
	;
	v490 = int32(0)
	v493 = v484 & int32(_a_F_btvacuumscan_5)
	if base.Ui32(v493) < base.Ui32(v475) {
		v826 = v452
		v827 = v490
		v868 = v485
		v870 = float64(0)
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v497 = int32(0)
	v505 = v452
	v506 = v490
	v507 = v475
	v512 = v497
	v513 = v497
	goto L111
L111:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v384+int32(20)+v507&int32(_a_F_btvacuumscan_5)<<(uint(int32(2))%32))))
	v557 = v384 + v554&int32(_a_F_btvacuumscan_7)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+7)))
	if v558&int32(32) != 0 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v826 = v769
	v827 = v770
	v868 = base.F64_convert_i32_s(v777)
	v870 = base.F64_convert_i32_s(v776)
	goto L106
L113:
	;
	v814 = v507 + int32(1)
	if base.Ui32(v814&int32(_a_F_btvacuumscan_5)) <= base.Ui32(v493) {
		v505 = v769
		v506 = v770
		v507 = v814
		v512 = v776
		v513 = v777
		goto L111
	} else {
		goto L144
	}
L114:
	;
	v580 = v561 & int32(4095)
	if v580 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L115:
	;
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557)+4)))
	if v561&int32(_a_F_btvacuumscan_1) != 0 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v565 = m.T0[v346].(func(*base.Module, int32, int32) int32)(m, v557, v356)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	if v565 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v569 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v326+int32(1696)+v505<<(uint(v569)%32)))) = uint16(v507)
	v769 = v505 + v569
	v770 = v506
	v776 = v512
	v777 = v513 + v569
	goto L113
L121:
	;
	goto L122
L122:
	;
	v769 = v505
	v770 = v506
	v776 = v512 + int32(1)
	v777 = v513
	goto L113
L123:
	;
	v769 = v718
	v770 = v719
	v776 = v512 + v729
	v777 = v726
	goto L113
L124:
	;
	v718 = v505
	v719 = v506
	v726 = v513
	v729 = int32(0)
	goto L123
L125:
	;
	goto L126
L126:
	;
	v584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557)+2)))
	v585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557))))
	v594 = int32(0)
	v599 = v594
	v601 = v594
	v614 = v594
	goto L127
L127:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v326)+36))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v326)+32))
	v652 = m.T0[v651].(func(*base.Module, int32, int32) int32)(m, v584+(v557+v585<<(uint(int32(16))%32))+v599*int32(6), v650)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	if v675 == int32(0) {
		v718 = v505
		v719 = v506
		v726 = v513
		v729 = v676
		goto L123
	} else {
		goto L139
	}
L129:
	;
	v679 = v599 + int32(1)
	if v679 != v580 {
		v599 = v679
		v601 = v675
		v614 = v676
		goto L127
	} else {
		goto L138
	}
L130:
	;
	if v652 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v675 = v601
	v676 = v614 + int32(1)
	goto L129
L132:
	;
	goto L133
L133:
	;
	if v601 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v660 = F_palloc(m, v580<<(uint(int32(1))%32)+int32(8))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v601)+6)))
	v668 = int32(1)
	v669 = v667 + v668
	*(*uint16)(unsafe.Add(mBase, uint32(v601)+6)) = uint16(v669)
	*(*uint16)(unsafe.Add(mBase, uint32(v601+v667<<(uint(v668)%32))+8)) = uint16(v599)
	v675 = v601
	v676 = v614
	goto L129
L137:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v660)+8)) = uint16(v599)
	v663 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v660)+6)) = uint16(v663)
	*(*uint16)(unsafe.Add(mBase, uint32(v660)+4)) = uint16(v507)
	*(*int32)(unsafe.Add(mBase, uint32(v660))) = v557
	v675 = v660
	v676 = v614
	goto L129
L138:
	;
	goto L128
L139:
	;
	if int32(0) < v676 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v326-int32(-64)+v506<<(uint(int32(2))%32)))) = v675
	v692 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557)+4)))
	v718 = v505
	v719 = v506 + int32(1)
	v726 = v513 - v676 + v692&int32(4095)
	v729 = v676
	goto L123
L141:
	;
	goto L142
L142:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v326+int32(1696)+v505<<(uint(int32(1))%32)))) = uint16(v507)
	v704 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557)+4)))
	F_pfree(m, v675)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v718 = v505 + int32(1)
	v719 = v506
	v726 = v513 + v704&int32(4095)
	v729 = v676
	goto L123
L144:
	;
	goto L112
L145:
	;
	if base.Ui32(v475) <= base.Ui32(v1716&int32(_a_F_btvacuumscan_5)) {
		goto L241
	} else {
		goto L242
	}
L146:
	;
	v879 = v326 + int32(1696)
	v881 = v326 - int32(-64)
	v882 = int32(0)
	v887 = m.G0
	v889 = v887 - int32(832)
	m.G0 = v889
	if v317 < v882 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	goto L148
L148:
	;
	v1686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326)+40)))
	if v1686 == int32(0) {
		v1716 = v484
		goto L145
	} else {
		goto L238
	}
L149:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v910)+118)))
	if v911 != int32(112) {
		v924 = int32(0)
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v894 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v894+(v317^int32(-1))<<(uint(int32(2))%32))))
	v908 = v900
	goto L149
L151:
	;
	goto L152
L152:
	;
	v902 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v908 = v902 + v317<<(uint(int32(13))%32) + int32(-8192)
	goto L149
L153:
	;
	if int32(0) < v827 {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v916 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if int32(0) < v916 {
		v924 = int32(1)
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v323)+32))
	if v920 != 0 {
		v924 = int32(0)
		goto L153
	} else {
		goto L156
	}
L156:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v323)+40))
	v924 = base.B2i32(v921 == int32(0))
	goto L153
L157:
	;
	if int32(0) < v826 {
		goto L198
	} else {
		goto L199
	}
L158:
	;
	v935 = v882
	v938 = v882
	goto L161
L159:
	;
	goto L160
L160:
	;
	v1377 = int32(_a_F_btvacuumscan_4)
	v1379 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1379 + int32(1)
	v1385 = v882
	v1397 = v882
	goto L157
L161:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v881+v938<<(uint(int32(2))%32))))
	F__bt_update_posting(m, v980)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	v999 = int32(0)
	if v924 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v980)+6)))
	v986 = int32(1)
	v989 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v980)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v889+int32(16)+v938<<(uint(v986)%32)))) = uint16(v989)
	v995 = v935 + v983<<(uint(v986)%32) + int32(2)
	v997 = v938 + v986
	if v997 != v827 {
		v935 = v995
		v938 = v997
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v1001 = int32(0)
	v1002 = F_palloc(m, v995)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	v1214 = v882
	v1226 = v999
	goto L167
L167:
	;
	v1262 = int32(_a_F_btvacuumscan_4)
	v1264 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1264 + int32(1)
	v1279 = v999
	goto L184
L168:
	;
	if v827 != int32(1) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v1214 = v995
	v1226 = v1002
	goto L167
L170:
	;
	v1012 = v882
	v1014 = v1001
	v1028 = v882
	goto L173
L171:
	;
	v1098 = v882
	v1100 = v1001
	goto L172
L172:
	;
	v1146 = v1098 + v1002
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v881+v1100<<(uint(int32(2))%32))))
	v1151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1150)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1146))) = uint16(v1151)
	v1154 = v1151 << (uint(int32(1)) % 32)
	if v1154 == int32(0) {
		goto L169
	} else {
		goto L183
	}
L173:
	;
	v1061 = int32(2)
	v1063 = v881 + v1014<<(uint(v1061)%32)
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	v1065 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1064)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1012+v1002))) = uint16(v1065)
	v1068 = v1012 + v1061
	v1070 = v1065 << (uint(int32(1)) % 32)
	if v1070 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v827&int32(1) == int32(0) {
		goto L169
	} else {
		goto L182
	}
L175:
	;
	base.MemoryCopy(m, v1068+v1002, v1064+int32(8), v1070)
	goto L177
L176:
	;
	goto L177
L177:
	;
	v1075 = v1068 + v1070
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	v1078 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1077)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1002+v1075))) = uint16(v1078)
	v1081 = v1075 + int32(2)
	v1083 = v1078 << (uint(int32(1)) % 32)
	if v1083 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	base.MemoryCopy(m, v1081+v1002, v1077+int32(8), v1083)
	goto L180
L179:
	;
	goto L180
L180:
	;
	v1088 = v1081 + v1083
	v1089 = int32(2)
	v1090 = v1014 + v1089
	v1092 = v1028 + v1089
	if v1092 != v827&int32(2147483646) {
		v1012 = v1088
		v1014 = v1090
		v1028 = v1092
		goto L173
	} else {
		goto L181
	}
L181:
	;
	goto L174
L182:
	;
	v1098 = v1088
	v1100 = v1090
	goto L172
L183:
	;
	base.MemoryCopy(m, v1146+int32(2), v1150+int32(8), v1154)
	goto L169
L184:
	;
	v1323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v889+int32(16)+v1279<<(uint(int32(1))%32)))))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v881+v1279<<(uint(int32(2))%32))))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1327)))
	v1329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1328)+6)))
	v1336 = F_PageIndexTupleOverwrite(m, v908, v1323, v1328, (v1329&int32(_a_F_btvacuumscan_8)+int32(7))&int32(_a_F_btvacuumscan_9))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L186
	}
L185:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L1
	} else {
		goto L191
	}
L186:
	;
	if v1336 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1339 = v1279 + int32(1)
	if v1339 != v827 {
		v1279 = v1339
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
	v1385 = v1214
	v1397 = v1226
	goto L157
L191:
	;
	if v317 < int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v889))) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v889)+4)) = v1364 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_10), v889)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L1
	} else {
		goto L196
	}
L193:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1348+(v317^int32(-1))<<(uint(int32(6))%32))+16))
	v1363 = v1354
	goto L192
L194:
	;
	goto L195
L195:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1356+v317<<(uint(int32(6))%32)+int32(-64))+16))
	v1363 = v1362
	goto L192
L196:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(1200), int32(_a_F_btvacuumscan_12))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
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
	F_PageIndexMultiDelete(m, v908, v879, v826)
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v1437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v908)+16)))
	v1438 = v908 + v1437
	v1439 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1438)+14)) = uint16(v1439)
	v1441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1438)+12)))
	v1443 = v1441 & int32(_a_F_btvacuumscan_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v1438)+12)) = uint16(v1443)
	F_MarkBufferDirty(m, v317)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L202
	}
L201:
	;
	goto L200
L202:
	;
	if v924 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v889)+14)) = uint16(v827)
	*(*uint16)(unsafe.Add(mBase, uint32(v889)+12)) = uint16(v826)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v1487 = int32(_a_F_btvacuumscan_4)
	v1489 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1489 - int32(1)
	if v1397 != 0 {
		goto L219
	} else {
		goto L220
	}
L206:
	;
	F_XLogRegisterBuffer(m, int32(0), v317, int32(8))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_XLogRegisterData(m, v889+int32(12), int32(4))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	if int32(0) < v826 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_XLogRegisterBufData(m, int32(0), v879, v826<<(uint(int32(1))%32))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L1
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	if int32(0) < v827 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	goto L211
L213:
	;
	F_XLogRegisterBufData(m, int32(0), v889+int32(16), v827<<(uint(int32(1))%32))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1481 = F_XLogInsert(m, int32(11), int32(192))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L1
	} else {
		goto L218
	}
L216:
	;
	F_XLogRegisterBufData(m, int32(0), v1397, v1385)
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	goto L215
L218:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v908))) = base.I64_rotr(v1481, int64(32))
	goto L205
L219:
	;
	F_pfree(m, v1397)
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	if int32(0) < v827 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	goto L221
L223:
	;
	v1508 = int32(0)
	goto L226
L224:
	;
	goto L225
L225:
	;
	m.G0 = v889 + int32(832)
	v1610 = *(*float64)(unsafe.Add(mBase, uint32(v337)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v337)+16)) = base.F64_add(v868, v1610)
	v1613 = int32(0)
	v1614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1614) {
		goto L230
	} else {
		goto L231
	}
L226:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v881+v1508<<(uint(int32(2))%32))))
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1550)))
	F_pfree(m, v1551)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L228
	}
L227:
	;
	goto L225
L228:
	;
	v1555 = v1508 + int32(1)
	if v1555 != v827 {
		v1508 = v1555
		goto L226
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	v1622 = int32(base.Ui32(v1614+int32(_a_F_btvacuumscan_6)) >> (uint(int32(2)) % 32))
	goto L232
L231:
	;
	v1622 = v1613
	goto L232
L232:
	;
	if v827 <= int32(0) {
		v1716 = v1622
		goto L145
	} else {
		goto L233
	}
L233:
	;
	v1627 = v1613
	goto L234
L234:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v326-int32(-64)+v1627<<(uint(int32(2))%32))))
	F_pfree(m, v1680)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L1
	} else {
		goto L236
	}
L235:
	;
	v1716 = v1622
	goto L145
L236:
	;
	v1684 = v1627 + int32(1)
	if v1684 != v827 {
		v1627 = v1684
		goto L234
	} else {
		goto L237
	}
L237:
	;
	goto L235
L238:
	;
	v1689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+14)))
	if v1689 != v1686 {
		v1716 = v484
		goto L145
	} else {
		goto L239
	}
L239:
	;
	v1691 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v389)+14)) = uint16(v1691)
	F_MarkBufferDirtyHint(m, v317, int32(1))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v1716 = v484
	goto L145
L241:
	;
	if v346 != 0 {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	goto L243
L243:
	;
	if v324 != v341 {
		v4052 = v471
		goto L55
	} else {
		goto L248
	}
L244:
	;
	v1755 = v870
	goto L246
L245:
	;
	v1755 = base.F64_convert_i32_u((v1716 - v475 + int32(1)) & int32(_a_F_btvacuumscan_5))
	goto L246
L246:
	;
	v1756 = *(*float64)(unsafe.Add(mBase, uint32(v337)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v337)+8)) = base.F64_add(v1755, v1756)
	F__bt_relbuf(m, v317)
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v4104 = v471
	goto L54
L248:
	;
	v1791 = v471
	goto L86
L249:
	;
	v1815 = int32(_a_F_btvacuumscan_14)
	v1816 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0]))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v326)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0])) = v1818
	v1821 = v326 + int32(24)
	v1822 = int32(0)
	v1823 = m.G0
	v1825 = v1823 - int32(288)
	m.G0 = v1825
	if v317 < v1822 {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v1856 = v317
	v1866 = v1822
	goto L261
L251:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1830+(v317^int32(-1))<<(uint(int32(6))%32))+16))
	v1845 = v1836
	goto L250
L252:
	;
	goto L253
L253:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1838+v317<<(uint(int32(6))%32)+int32(-64))+16))
	v1845 = v1844
	goto L250
L254:
	;
	m.G0 = v1825 + int32(288)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0])) = v1816
	v4104 = v1791
	goto L54
L255:
	;
	F_ReleaseBuffer(m, v1856)
	mBase = m.M
	v3967 = m.ExcPending
	if v3967 != 0 {
		goto L1
	} else {
		goto L724
	}
L256:
	;
	F_LockBuffer(m, v1856, int32(0))
	mBase = m.M
	v3915 = m.ExcPending
	if v3915 != 0 {
		goto L1
	} else {
		goto L723
	}
L257:
	;
	v3843 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3844 = m.ExcPending
	if v3844 != 0 {
		goto L1
	} else {
		goto L718
	}
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3774 = m.ExcPending
	if v3774 != 0 {
		goto L1
	} else {
		goto L715
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3749 = m.ExcPending
	if v3749 != 0 {
		goto L1
	} else {
		goto L711
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L1
	} else {
		goto L708
	}
L261:
	;
	v1896 = int32(0)
	v1897 = base.B2i32(v1896 <= v1856)
	if v1897 == v1896 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3716 = m.ExcPending
	if v3716 != 0 {
		goto L1
	} else {
		goto L705
	}
L263:
	;
	v1916 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1915)+16)))
	v1917 = v1916 + v1915
	v1918 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1917)+12)))
	if v1918&int32(5) != int32(1) {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1901+(v1856^int32(-1))<<(uint(int32(2))%32))))
	v1915 = v1907
	goto L263
L265:
	;
	goto L266
L266:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v1915 = v1909 + v1856<<(uint(int32(13))%32) + int32(-8192)
	goto L263
L267:
	;
	if v1918&int32(16) == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	if v1918&int32(2) != 0 {
		goto L291
	} else {
		goto L292
	}
L270:
	;
	v1954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1917)+12)))
	if v1954&int32(4) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L271:
	;
	v1929 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	if v1929 == int32(0) {
		goto L270
	} else {
		goto L273
	}
L273:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+208)) = v1936 + int32(4)
	F_errmsg(m, int32(_a_F_btvacuumscan_15), v1825+int32(208))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	F_errhint(m, int32(_a_F_btvacuumscan_16), int32(0))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(1863), int32(_a_F_btvacuumscan_0))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	goto L270
L278:
	;
	F_LockBuffer(m, v1856, int32(0))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L1
	} else {
		goto L289
	}
L279:
	;
	v1961 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	if v1961 == int32(0) {
		goto L278
	} else {
		goto L281
	}
L281:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	if v1856 < int32(0) {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+196)) = v1845
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+192)) = v1986
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+200)) = v1987 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_17), v1825+int32(192))
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L1
	} else {
		goto L287
	}
L284:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v1971+(v1856^int32(-1))<<(uint(int32(6))%32))+16))
	v1986 = v1977
	goto L283
L285:
	;
	goto L286
L286:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1979+v1856<<(uint(int32(6))%32)+int32(-64))+16))
	v1986 = v1985
	goto L283
L287:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(1871), int32(_a_F_btvacuumscan_0))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
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
	if v1918&int32(16) == int32(0) {
		goto L297
	} else {
		goto L298
	}
L291:
	;
	F_LockBuffer(m, v1856, int32(0))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L1
	} else {
		goto L295
	}
L292:
	;
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v1917)+4))
	if base.B2i32(v2010 == int32(0))|v1918&int32(128) != 0 {
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v2016 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1915)+12)))
	if base.B2i32(base.Ui32(v2016) < base.Ui32(int32(25)))|base.B2i32((v2016+int32(_a_F_btvacuumscan_6))&int32(_a_F_btvacuumscan_18) == int32(0)) != 0 {
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
	if v1866 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L298:
	;
	v2593 = v1918
	goto L299
L299:
	;
	if v2593&int32(16) != 0 {
		goto L429
	} else {
		goto L430
	}
L300:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v1915)+24))
	v2040 = F_CopyIndexTuple(m, v1915+v2036&int32(_a_F_btvacuumscan_7))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L1
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v1821)))
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v2086)+4))
	if v1897 == int32(0) {
		goto L317
	} else {
		goto L318
	}
L303:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	if v1856 < int32(0) {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	F_LockBuffer(m, v1856, int32(0))
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L1
	} else {
		goto L308
	}
L305:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v2046+(v1856^int32(-1))<<(uint(int32(6))%32))+16))
	v2061 = v2052
	goto L304
L306:
	;
	goto L307
L307:
	;
	v2054 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v2054+v1856<<(uint(int32(6))%32)+int32(-64))+16))
	v2061 = v2060
	goto L304
L308:
	;
	v2065 = F__bt_leftsib_splitflag(m, v323, v2042, v2061)
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	if v2065 != 0 {
		goto L255
	} else {
		goto L310
	}
L310:
	;
	v2067 = F__bt_mkscankey(m, v323, v2040)
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	v2069 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2067)+3)) = uint16(v2069)
	v2075 = F__bt_search(m, v323, int32(0), v2067, v1825+int32(248), int32(1))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v1825)+248))
	F_LockBuffer(m, v2077, int32(0))
	mBase = m.M
	v2080 = m.ExcPending
	if v2080 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_ReleaseBuffer(m, v2077)
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_LockBuffer(m, v1856, int32(2))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1866 = v2075
	goto L261
L316:
	;
	v2106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2105)+16)))
	if v1856 < int32(0) {
		goto L321
	} else {
		goto L322
	}
L317:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2091+(v1856^int32(-1))<<(uint(int32(2))%32))))
	v2105 = v2097
	goto L316
L318:
	;
	goto L319
L319:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2105 = v2099 + v1856<<(uint(int32(13))%32) + int32(-8192)
	goto L316
L320:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2105+v2106)+4))
	v2128 = F_ReadBuffer(m, v323, v2127)
	mBase = m.M
	v2129 = m.ExcPending
	if v2129 != 0 {
		goto L1
	} else {
		goto L324
	}
L321:
	;
	v2110 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2110+(v1856^int32(-1))<<(uint(int32(6))%32))+16))
	v2125 = v2116
	goto L320
L322:
	;
	goto L323
L323:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v2118+v1856<<(uint(int32(6))%32)+int32(-64))+16))
	v2125 = v2124
	goto L320
L324:
	;
	F_LockBuffer(m, v2128, int32(1))
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	F__bt_checkpage(m, v323, v2128)
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	if v2128 < int32(0) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v2153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2152)+16)))
	v2155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2153+v2152)+12)))
	F_LockBuffer(m, v2128, int32(0))
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L1
	} else {
		goto L331
	}
L328:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v2138+(v2128^int32(-1))<<(uint(int32(2))%32))))
	v2152 = v2144
	goto L327
L329:
	;
	goto L330
L330:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2152 = v2146 + v2128<<(uint(int32(13))%32) + int32(-8192)
	goto L327
L331:
	;
	F_ReleaseBuffer(m, v2128)
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	if v2155&int32(16) != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v2165 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2166 = m.ExcPending
	if v2166 != 0 {
		goto L1
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v2181 = F__bt_getstackbuf(m, v323, v2087, v1866, v2125)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L1
	} else {
		goto L340
	}
L336:
	;
	if v2165 == int32(0) {
		goto L256
	} else {
		goto L337
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+180)) = v2127
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+176)) = v2125
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_19), v1825+int32(176))
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2128), int32(_a_F_btvacuumscan_20))
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	goto L256
L340:
	;
	if v2181 == int32(0) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v3795 = v2125
	goto L257
L342:
	;
	goto L343
L343:
	;
	v2188 = v1866
	v2191 = v2181
	v2192 = v2127
	v2199 = v2125
	goto L344
L344:
	;
	v2235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2188)+4)))
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v2188)))
	v2237 = int32(0)
	v2238 = base.B2i32(v2237 <= v2191)
	if v2238 == v2237 {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	if v2238 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L346:
	;
	v2257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2256)+16)))
	v2258 = v2257 + v2256
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v2258)))
	v2260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2256)+12)))
	if base.B2i32(base.Ui32(int32(25)) <= base.Ui32(v2260))&base.B2i32(base.Ui32(v2235) < base.Ui32(int32(base.Ui32(v2260+int32(_a_F_btvacuumscan_6))>>(uint(int32(2))%32))&int32(_a_F_btvacuumscan_5))) == int32(0) {
		goto L350
	} else {
		goto L351
	}
L347:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v2242+(v2191^int32(-1))<<(uint(int32(2))%32))))
	v2256 = v2248
	goto L346
L348:
	;
	goto L349
L349:
	;
	v2250 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2256 = v2250 + v2191<<(uint(int32(13))%32) + int32(-8192)
	goto L346
L350:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2258)+4))
	F_LockBuffer(m, v2191, int32(0))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
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
	F_ReleaseBuffer(m, v2191)
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	if v2273 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v2283 = int32(2)
	goto L357
L356:
	;
	v2283 = int32(1)
	goto L357
L357:
	;
	if base.B2i32(v2273 == int32(0))|base.B2i32(v2283 != v2235) != 0 {
		goto L256
	} else {
		goto L358
	}
L358:
	;
	v2286 = F__bt_leftsib_splitflag(m, v323, v2259, v2236)
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	if v2286 != 0 {
		goto L256
	} else {
		goto L360
	}
L360:
	;
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+8))
	v2289 = F__bt_getstackbuf(m, v323, v2087, v2288, v2236)
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	if v2289 == int32(0) {
		v3795 = v2236
		goto L257
	} else {
		goto L362
	}
L362:
	;
	v2188 = v2288
	v2191 = v2289
	v2192 = v2273
	v2199 = v2236
	goto L344
L363:
	;
	v2314 = (v2235 + int32(1)) & int32(_a_F_btvacuumscan_5)
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v2310+v2314<<(uint(int32(2))%32))+20))
	v2321 = v2318&int32(_a_F_btvacuumscan_7) + v2310
	v2322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2321))))
	v2325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2321)+2)))
	if v2192 != v2322<<(uint(int32(16))%32)|v2325 {
		goto L367
	} else {
		goto L368
	}
L364:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v2296+(v2191^int32(-1))<<(uint(int32(2))%32))))
	v2310 = v2302
	goto L363
L365:
	;
	goto L366
L366:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2310 = v2304 + v2191<<(uint(int32(13))%32) + int32(-8192)
	goto L363
L367:
	;
	v2330 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L1
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	F_PredicateLockPageSplit(m, v323, v2125, v2127)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L1
	} else {
		goto L383
	}
L370:
	;
	if v2330 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L1
	} else {
		goto L374
	}
L372:
	;
	goto L373
L373:
	;
	F_LockBuffer(m, v2191, int32(0))
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L1
	} else {
		goto L381
	}
L374:
	;
	v2335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2321)+2)))
	v2336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2321))))
	if v2191 < int32(0) {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+160)) = v2356 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+156)) = v2355
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+152)) = v2335 | v2336<<(uint(int32(16))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+148)) = v2199
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+144)) = v2192
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_21), v1825+int32(144))
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L1
	} else {
		goto L379
	}
L376:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v2340+(v2191^int32(-1))<<(uint(int32(6))%32))+16))
	v2355 = v2346
	goto L375
L377:
	;
	goto L378
L378:
	;
	v2348 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(v2348+v2191<<(uint(int32(6))%32)+int32(-64))+16))
	v2355 = v2354
	goto L375
L379:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2187), int32(_a_F_btvacuumscan_20))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	goto L373
L381:
	;
	F_ReleaseBuffer(m, v2191)
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	goto L256
L383:
	;
	v2387 = int32(_a_F_btvacuumscan_4)
	v2389 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v2389 + int32(1)
	if v2238 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	v2414 = *(*int32)(unsafe.Add(mBase, uint32(v2410+v2235<<(uint(int32(2))%32))+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2414&int32(_a_F_btvacuumscan_7)+v2410))) = base.I32_rotr(v2192, int32(16))
	F_PageIndexTupleDelete(m, v2410, v2314)
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L1
	} else {
		goto L388
	}
L385:
	;
	v2396 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v2396+(v2191^int32(-1))<<(uint(int32(2))%32))))
	v2410 = v2402
	goto L384
L386:
	;
	goto L387
L387:
	;
	v2404 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2410 = v2404 + v2191<<(uint(int32(13))%32) + int32(-8192)
	goto L384
L388:
	;
	if v1897 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	v2441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2440)+16)))
	v2442 = v2441 + v2440
	v2443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2442)+12)))
	v2445 = v2443 | int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v2442)+12)) = uint16(v2445)
	v2448 = base.B2i32(v2125 == v2199)
	if v2125 == v2199 {
		goto L393
	} else {
		goto L394
	}
L390:
	;
	v2426 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(v2426+(v1856^int32(-1))<<(uint(int32(2))%32))))
	v2440 = v2432
	goto L389
L391:
	;
	goto L392
L392:
	;
	v2434 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2440 = v2434 + v1856<<(uint(int32(13))%32) + int32(-8192)
	goto L389
L393:
	;
	v2449 = int32(-1)
	goto L395
L394:
	;
	v2449 = v2199
	goto L395
L395:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1825)+222)) = uint16(v2449)
	if v2125 == v2199 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v2454 = int32(-1)
	goto L398
L397:
	;
	v2454 = int32(base.Ui32(v2199) >> (uint(int32(16)) % 32))
	goto L398
L398:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1825)+220)) = uint16(v2454)
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+224)) = int32(537395200)
	v2462 = F_PageIndexTupleOverwrite(m, v2440, int32(1), v1825+int32(220), int32(8))
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	if v2462 == int32(0) {
		goto L296
	} else {
		goto L400
	}
L400:
	;
	F_MarkBufferDirty(m, v2191)
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	F_MarkBufferDirty(m, v1856)
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	v2471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2470)+118)))
	if v2471 != int32(112) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v2575 = int32(_a_F_btvacuumscan_4)
	v2577 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v2577 - int32(1)
	F_LockBuffer(m, v2191, int32(0))
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L1
	} else {
		goto L427
	}
L404:
	;
	v2475 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if v2475 <= int32(0) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v2478 = *(*int32)(unsafe.Add(mBase, uint32(v323)+32))
	if v2478 != 0 {
		goto L403
	} else {
		goto L408
	}
L406:
	;
	goto L407
L407:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1825)+248)) = uint16(v2235)
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+252)) = v2125
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+264)) = v2449
	F_XLogBeginInsert(m)
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		goto L1
	} else {
		goto L410
	}
L408:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v323)+40))
	if v2479 != 0 {
		goto L403
	} else {
		goto L409
	}
L409:
	;
	goto L407
L410:
	;
	F_XLogRegisterBuffer(m, int32(0), v1856, int32(6))
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	F_XLogRegisterBuffer(m, int32(1), v2191, int32(8))
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	if v1897 == int32(0) {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	v2511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2510)+16)))
	v2512 = v2511 + v2510
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v2512)))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+256)) = v2513
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v2512)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+260)) = v2515
	F_XLogRegisterData(m, v1825+int32(248), int32(20))
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L1
	} else {
		goto L417
	}
L414:
	;
	v2496 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v2496+(v1856^int32(-1))<<(uint(int32(2))%32))))
	v2510 = v2502
	goto L413
L415:
	;
	goto L416
L416:
	;
	v2504 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2510 = v2504 + v1856<<(uint(int32(13))%32) + int32(-8192)
	goto L413
L417:
	;
	v2524 = F_XLogInsert(m, int32(11), int32(176))
	mBase = m.M
	v2525 = m.ExcPending
	if v2525 != 0 {
		goto L1
	} else {
		goto L418
	}
L418:
	;
	if v2238 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	v2544 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v2543))) = base.I64_rotr(v2524, v2544)
	if v1897 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L420:
	;
	v2529 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2535 = *(*int32)(unsafe.Add(mBase, uint32(v2529+(v2191^int32(-1))<<(uint(int32(2))%32))))
	v2543 = v2535
	goto L419
L421:
	;
	goto L422
L422:
	;
	v2537 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2543 = v2537 + v2191<<(uint(int32(13))%32) + int32(-8192)
	goto L419
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2568)+4)) = base.I32_wrap_i64(v2524)
	*(*int32)(unsafe.Add(mBase, uint32(v2568))) = base.I32_wrap_i64(int64(base.Ui64(v2524) >> (uint(v2544) % 64)))
	goto L403
L424:
	;
	v2554 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v2554+(v1856^int32(-1))<<(uint(int32(2))%32))))
	v2568 = v2560
	goto L423
L425:
	;
	goto L426
L426:
	;
	v2562 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2568 = v2562 + v1856<<(uint(int32(13))%32) + int32(-8192)
	goto L423
L427:
	;
	F_ReleaseBuffer(m, v2191)
	mBase = m.M
	v2585 = m.ExcPending
	if v2585 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	v2586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1917)+12)))
	v2593 = v2586
	goto L299
L429:
	;
	v2641 = v1856 ^ int32(-1)
	v2643 = v1856 << (uint(int32(13)) % 32)
	goto L432
L430:
	;
	v3652 = int32(0)
	goto L431
L431:
	;
	v3694 = *(*int32)(unsafe.Add(mBase, uint32(v1917)+4))
	F_LockBuffer(m, v1856, int32(0))
	mBase = m.M
	v3697 = m.ExcPending
	if v3697 != 0 {
		goto L1
	} else {
		goto L695
	}
L432:
	;
	if v1856 < int32(0) {
		goto L435
	} else {
		goto L436
	}
L433:
	;
	v3633 = int32(2)
	if v3197 != 0 {
		goto L692
	} else {
		goto L693
	}
L434:
	;
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+4))
	if v1897 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L435:
	;
	v2697 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(v2697+(v1856^int32(-1))<<(uint(int32(6))%32))+16))
	v2712 = v2703
	goto L434
L436:
	;
	goto L437
L437:
	;
	v2705 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v2705+v1856<<(uint(int32(6))%32)+int32(-64))+16))
	v2712 = v2711
	goto L434
L438:
	;
	v2728 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2727)+16)))
	v2729 = v2728 + v2727
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2729)+4))
	v2731 = *(*int32)(unsafe.Add(mBase, uint32(v2729)))
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(v2727)+24))
	v2735 = v2727 + v2732&int32(_a_F_btvacuumscan_7)
	v2736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2735)+2)))
	v2737 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2735))))
	F_LockBuffer(m, v1856, int32(0))
	mBase = m.M
	v2740 = m.ExcPending
	if v2740 != 0 {
		goto L1
	} else {
		goto L442
	}
L439:
	;
	v2717 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(v2717+v2641<<(uint(int32(2))%32))))
	v2727 = v2721
	goto L438
L440:
	;
	goto L441
L441:
	;
	v2723 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2727 = v2723 + v2643 + int32(-8192)
	goto L438
L442:
	;
	v2743 = v2736 | v2737<<(uint(int32(16))%32)
	v2745 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v2745 != 0 {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2747 = m.ExcPending
	if v2747 != 0 {
		goto L1
	} else {
		goto L446
	}
L444:
	;
	goto L445
L445:
	;
	v2749 = int32(1)
	if v2743 == int32(-1) {
		goto L448
	} else {
		goto L449
	}
L446:
	;
	goto L445
L447:
	;
	v2795 = int32(0)
	if v2790 == v2795 {
		v2979 = int32(0)
		v2981 = v2795
		goto L463
	} else {
		goto L464
	}
L448:
	;
	v2790 = v2731
	v2791 = v1856
	v2792 = v2712
	v2793 = v2749
	v2794 = int32(0)
	goto L447
L449:
	;
	goto L450
L450:
	;
	v2753 = F_ReadBuffer(m, v323, v2743)
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	F_LockBuffer(m, v2753, int32(1))
	mBase = m.M
	v2757 = m.ExcPending
	if v2757 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	F__bt_checkpage(m, v323, v2753)
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	if v2753 < int32(0) {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	v2778 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2777)+16)))
	v2779 = v2778 + v2777
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v2779)+8))
	v2781 = *(*int32)(unsafe.Add(mBase, uint32(v2779)))
	F_LockBuffer(m, v2753, int32(0))
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L1
	} else {
		goto L458
	}
L455:
	;
	v2763 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2769 = *(*int32)(unsafe.Add(mBase, uint32(v2763+(v2753^int32(-1))<<(uint(int32(2))%32))))
	v2777 = v2769
	goto L454
L456:
	;
	goto L457
L457:
	;
	v2771 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2777 = v2771 + v2753<<(uint(int32(13))%32) + int32(-8192)
	goto L454
L458:
	;
	if v2712 == v2743 {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v2790 = v2781
	v2791 = v2753
	v2792 = v2712
	v2793 = v2749
	v2794 = v2780
	goto L447
L460:
	;
	goto L461
L461:
	;
	F_LockBuffer(m, v1856, int32(2))
	mBase = m.M
	v2788 = m.ExcPending
	if v2788 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	v2790 = v2781
	v2791 = v2753
	v2792 = v2743
	v2793 = int32(0)
	v2794 = v2780
	goto L447
L463:
	;
	F_LockBuffer(m, v2791, int32(2))
	mBase = m.M
	v3025 = m.ExcPending
	if v3025 != 0 {
		goto L1
	} else {
		goto L507
	}
L464:
	;
	v2798 = F_ReadBuffer(m, v323, v2790)
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F_LockBuffer(m, v2798, int32(2))
	mBase = m.M
	v2802 = m.ExcPending
	if v2802 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	F__bt_checkpage(m, v323, v2798)
	mBase = m.M
	v2804 = m.ExcPending
	if v2804 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	if v2798 < int32(0) {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	v2823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2822)+16)))
	v2824 = v2823 + v2822
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v2824)+4))
	v2826 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2824)+12)))
	v2828 = v2826 & int32(4)
	if v2828 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L469:
	;
	v2808 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v2808+(v2798^int32(-1))<<(uint(int32(2))%32))))
	v2822 = v2814
	goto L468
L470:
	;
	goto L471
L471:
	;
	v2816 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2822 = v2816 + v2798<<(uint(int32(13))%32) + int32(-8192)
	goto L468
L472:
	;
	if v2825 == v2792 {
		v2979 = v2798
		v2981 = v2790
		goto L463
	} else {
		goto L475
	}
L473:
	;
	goto L474
L474:
	;
	__phi2836 = v2790
	__phi2837 = v2825
	__phi2839 = v2798
	__phi2846 = v2828
	v2836 = __phi2836
	v2837 = __phi2837
	v2839 = __phi2839
	v2846 = __phi2846
	goto L476
L475:
	;
	goto L474
L476:
	;
	v2883 = int32(0)
	if base.B2i32(base.B2i32(v2837 == v2883)|v2846&int32(_a_F_btvacuumscan_5) == v2883)&base.B2i32(v2836 != v2837) == v2883 {
		goto L478
	} else {
		goto L479
	}
L477:
	;
	v2979 = v2940
	v2981 = v2837
	goto L463
L478:
	;
	F_LockBuffer(m, v2839, int32(0))
	mBase = m.M
	v2896 = m.ExcPending
	if v2896 != 0 {
		goto L1
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	F_LockBuffer(m, v2839, int32(0))
	mBase = m.M
	v2933 = m.ExcPending
	if v2933 != 0 {
		goto L1
	} else {
		goto L493
	}
L481:
	;
	F_ReleaseBuffer(m, v2839)
	mBase = m.M
	v2898 = m.ExcPending
	if v2898 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	v2901 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2902 = m.ExcPending
	if v2902 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	if v2901 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2905 = m.ExcPending
	if v2905 != 0 {
		goto L1
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	F_ReleaseBuffer(m, v2791)
	mBase = m.M
	v2927 = m.ExcPending
	if v2927 != 0 {
		goto L1
	} else {
		goto L490
	}
L487:
	;
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+128)) = v2794
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+132)) = v2906 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+124)) = v1845
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+120)) = v2712
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+116)) = v2792
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+112)) = v2837
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_22), v1825+int32(112))
	mBase = m.M
	v2919 = m.ExcPending
	if v2919 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2439), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v2924 = m.ExcPending
	if v2924 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	goto L486
L490:
	;
	if v2793 != 0 {
		goto L254
	} else {
		goto L491
	}
L491:
	;
	F_LockBuffer(m, v1856, int32(0))
	mBase = m.M
	v2930 = m.ExcPending
	if v2930 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	goto L255
L493:
	;
	F_ReleaseBuffer(m, v2839)
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	v2937 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v2937 != 0 {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2939 = m.ExcPending
	if v2939 != 0 {
		goto L1
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	v2940 = F_ReadBuffer(m, v323, v2837)
	mBase = m.M
	v2941 = m.ExcPending
	if v2941 != 0 {
		goto L1
	} else {
		goto L499
	}
L498:
	;
	goto L497
L499:
	;
	F_LockBuffer(m, v2940, int32(2))
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	F__bt_checkpage(m, v323, v2940)
	mBase = m.M
	v2946 = m.ExcPending
	if v2946 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	if v2940 < int32(0) {
		goto L503
	} else {
		goto L504
	}
L502:
	;
	v2965 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2964)+16)))
	v2966 = v2965 + v2964
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v2966)+4))
	v2968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2966)+12)))
	v2970 = v2968 & int32(4)
	if v2970|base.B2i32(v2967 != v2792) != 0 {
		__phi2836 = v2837
		__phi2837 = v2967
		__phi2839 = v2940
		__phi2846 = v2970
		v2836 = __phi2836
		v2837 = __phi2837
		v2839 = __phi2839
		v2846 = __phi2846
		goto L476
	} else {
		goto L506
	}
L503:
	;
	v2950 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2956 = *(*int32)(unsafe.Add(mBase, uint32(v2950+(v2940^int32(-1))<<(uint(int32(2))%32))))
	v2964 = v2956
	goto L502
L504:
	;
	goto L505
L505:
	;
	v2958 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2964 = v2958 + v2940<<(uint(int32(13))%32) + int32(-8192)
	goto L502
L506:
	;
	goto L477
L507:
	;
	v3026 = int32(0)
	v3027 = base.B2i32(v3026 <= v2791)
	if v3027 == v3026 {
		goto L509
	} else {
		goto L510
	}
L508:
	;
	v3046 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3045)+16)))
	v3047 = v3046 + v3045
	v3048 = *(*int32)(unsafe.Add(mBase, uint32(v3047)+4))
	if v3048 == int32(0) {
		goto L260
	} else {
		goto L512
	}
L509:
	;
	v3031 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3037 = *(*int32)(unsafe.Add(mBase, uint32(v3031+(v2791^int32(-1))<<(uint(int32(2))%32))))
	v3045 = v3037
	goto L508
L510:
	;
	goto L511
L511:
	;
	v3039 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3045 = v3039 + v2791<<(uint(int32(13))%32) + int32(-8192)
	goto L508
L512:
	;
	v3051 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3047)+12)))
	if v3051&int32(6) != 0 {
		goto L260
	} else {
		goto L513
	}
L513:
	;
	v3054 = *(*int32)(unsafe.Add(mBase, uint32(v3047)))
	if v3054 != v2981 {
		goto L259
	} else {
		goto L514
	}
L514:
	;
	v3056 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3045)+12)))
	v3058 = v3056 + int32(_a_F_btvacuumscan_6)
	if v2793 != 0 {
		goto L516
	} else {
		goto L517
	}
L515:
	;
	v3114 = F_ReadBuffer(m, v323, v3048)
	mBase = m.M
	v3115 = m.ExcPending
	if v3115 != 0 {
		goto L1
	} else {
		goto L530
	}
L516:
	;
	v3059 = int32(17)
	if v3051&v3059 == v3059 {
		goto L519
	} else {
		goto L520
	}
L517:
	;
	goto L518
L518:
	;
	if v3051&int32(1)|base.B2i32(base.Ui32(v3056) < base.Ui32(int32(25)))|base.B2i32(v3058&int32(_a_F_btvacuumscan_24) != int32(8)) != 0 {
		goto L258
	} else {
		goto L526
	}
L519:
	;
	if base.B2i32(v3058&int32(_a_F_btvacuumscan_18) == int32(0))|base.B2i32(base.Ui32(v3056) < base.Ui32(int32(25))) != 0 {
		v3113 = int32(-1)
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
	v3074 = m.ExcPending
	if v3074 != 0 {
		goto L1
	} else {
		goto L523
	}
L522:
	;
	goto L521
L523:
	;
	v3075 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+64)) = v2792
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+68)) = v3075 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_25), v1825-int32(-64))
	mBase = m.M
	v3084 = m.ExcPending
	if v3084 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2486), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3089 = m.ExcPending
	if v3089 != 0 {
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
	v3101 = *(*int32)(unsafe.Add(mBase, uint32(v3045)+28))
	v3104 = v3045 + v3101&int32(_a_F_btvacuumscan_7)
	v3105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3104))))
	v3108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3104)+2)))
	v3109 = v3105<<(uint(int32(16))%32) | v3108
	if v2712 == v3109 {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v3111 = int32(-1)
	goto L529
L528:
	;
	v3111 = v3109
	goto L529
L529:
	;
	v3113 = v3111
	goto L515
L530:
	;
	F_LockBuffer(m, v3114, int32(2))
	mBase = m.M
	v3118 = m.ExcPending
	if v3118 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	F__bt_checkpage(m, v323, v3114)
	mBase = m.M
	v3120 = m.ExcPending
	if v3120 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v3121 = int32(0)
	v3122 = base.B2i32(v3121 <= v3114)
	if v3122 == v3121 {
		goto L534
	} else {
		goto L535
	}
L533:
	;
	v3141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3140)+16)))
	v3142 = v3141 + v3140
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(v3142)))
	if v2792 != v3143 {
		goto L537
	} else {
		goto L538
	}
L534:
	;
	v3126 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v3126+(v3114^int32(-1))<<(uint(int32(2))%32))))
	v3140 = v3132
	goto L533
L535:
	;
	goto L536
L536:
	;
	v3134 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3140 = v3134 + v3114<<(uint(int32(13))%32) + int32(-8192)
	goto L533
L537:
	;
	v3147 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3148 = m.ExcPending
	if v3148 != 0 {
		goto L1
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	v3193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3140)+12)))
	v3194 = int32(0)
	v3197 = *(*int32)(unsafe.Add(mBase, uint32(v3142)+4))
	if v2981|v3197 != 0 {
		v3258 = v3194
		v3259 = v3194
		v3260 = v3194
		goto L558
	} else {
		goto L559
	}
L540:
	;
	if v3147 != 0 {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3151 = m.ExcPending
	if v3151 != 0 {
		goto L1
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	if v2979 != 0 {
		goto L547
	} else {
		goto L548
	}
L544:
	;
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v3142)))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+52)) = v2794
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+48)) = v3153
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+56)) = v3152 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+44)) = v1845
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+40)) = v2712
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+36)) = v2792
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+32)) = v3048
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_26), v1825+int32(32))
	mBase = m.M
	v3167 = m.ExcPending
	if v3167 != 0 {
		goto L1
	} else {
		goto L545
	}
L545:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2540), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3172 = m.ExcPending
	if v3172 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	goto L543
L547:
	;
	F_LockBuffer(m, v2979, int32(0))
	mBase = m.M
	v3177 = m.ExcPending
	if v3177 != 0 {
		goto L1
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	F_LockBuffer(m, v3114, int32(0))
	mBase = m.M
	v3182 = m.ExcPending
	if v3182 != 0 {
		goto L1
	} else {
		goto L552
	}
L550:
	;
	F_ReleaseBuffer(m, v2979)
	mBase = m.M
	v3179 = m.ExcPending
	if v3179 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	goto L549
L552:
	;
	F_ReleaseBuffer(m, v3114)
	mBase = m.M
	v3184 = m.ExcPending
	if v3184 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	F_LockBuffer(m, v2791, int32(0))
	mBase = m.M
	v3187 = m.ExcPending
	if v3187 != 0 {
		goto L1
	} else {
		goto L554
	}
L554:
	;
	F_ReleaseBuffer(m, v2791)
	mBase = m.M
	v3189 = m.ExcPending
	if v3189 != 0 {
		goto L1
	} else {
		goto L555
	}
L555:
	;
	if v2793 != 0 {
		goto L254
	} else {
		goto L556
	}
L556:
	;
	F_LockBuffer(m, v1856, int32(0))
	mBase = m.M
	v3192 = m.ExcPending
	if v3192 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	goto L255
L558:
	;
	v3261 = int32(_a_F_btvacuumscan_4)
	v3263 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v3263 + int32(1)
	if v2979 != 0 {
		goto L577
	} else {
		goto L578
	}
L559:
	;
	if v3122 == int32(0) {
		goto L561
	} else {
		goto L562
	}
L560:
	;
	v3217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3216)+16)))
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(v3217+v3216)+4))
	if v3219 != 0 {
		v3258 = v3194
		v3259 = v3194
		v3260 = v3194
		goto L558
	} else {
		goto L564
	}
L561:
	;
	v3202 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(v3202+(v3114^int32(-1))<<(uint(int32(2))%32))))
	v3216 = v3208
	goto L560
L562:
	;
	goto L563
L563:
	;
	v3210 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3216 = v3210 + v3114<<(uint(int32(13))%32) + int32(-8192)
	goto L560
L564:
	;
	v3221 = F_ReadBuffer(m, v323, int32(0))
	mBase = m.M
	v3222 = m.ExcPending
	if v3222 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	F_LockBuffer(m, v3221, int32(2))
	mBase = m.M
	v3225 = m.ExcPending
	if v3225 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	F__bt_checkpage(m, v323, v3221)
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	if v3221 < int32(0) {
		goto L569
	} else {
		goto L570
	}
L568:
	;
	v3247 = v3245 + int32(24)
	v3248 = *(*int32)(unsafe.Add(mBase, uint32(v3245)+44))
	if base.Ui32(v3248) <= base.Ui32(v2794+int32(1)) {
		goto L572
	} else {
		goto L573
	}
L569:
	;
	v3231 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3237 = *(*int32)(unsafe.Add(mBase, uint32(v3231+(v3221^int32(-1))<<(uint(int32(2))%32))))
	v3245 = v3237
	goto L568
L570:
	;
	goto L571
L571:
	;
	v3239 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3245 = v3239 + v3221<<(uint(int32(13))%32) + int32(-8192)
	goto L568
L572:
	;
	v3258 = v3245
	v3259 = v3221
	v3260 = v3247
	goto L558
L573:
	;
	goto L574
L574:
	;
	F_LockBuffer(m, v3221, int32(0))
	mBase = m.M
	v3254 = m.ExcPending
	if v3254 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	F_ReleaseBuffer(m, v3221)
	mBase = m.M
	v3256 = m.ExcPending
	if v3256 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	v3258 = v3245
	v3259 = v3194
	v3260 = v3247
	goto L558
L577:
	;
	if v2979 < int32(0) {
		goto L581
	} else {
		goto L582
	}
L578:
	;
	goto L579
L579:
	;
	if v3122 == int32(0) {
		goto L585
	} else {
		goto L586
	}
L580:
	;
	v3285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3284)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v3285+v3284)+4)) = v3048
	goto L579
L581:
	;
	v3270 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v3270+(v2979^int32(-1))<<(uint(int32(2))%32))))
	v3284 = v3276
	goto L580
L582:
	;
	goto L583
L583:
	;
	v3278 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3284 = v3278 + v2979<<(uint(int32(13))%32) + int32(-8192)
	goto L580
L584:
	;
	v3307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3306)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v3307+v3306))) = v2981
	if v2793 == int32(0) {
		goto L588
	} else {
		goto L589
	}
L585:
	;
	v3292 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3298 = *(*int32)(unsafe.Add(mBase, uint32(v3292+(v3114^int32(-1))<<(uint(int32(2))%32))))
	v3306 = v3298
	goto L584
L586:
	;
	goto L587
L587:
	;
	v3300 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3306 = v3300 + v3114<<(uint(int32(13))%32) + int32(-8192)
	goto L584
L588:
	;
	v3312 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2735)+4)) = uint16(v3312)
	*(*int32)(unsafe.Add(mBase, uint32(v2735))) = base.I32_rotr(v3113, int32(16))
	v3317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2735)+6)))
	v3319 = v3317 | int32(_a_F_btvacuumscan_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2735)+6)) = uint16(v3319)
	goto L590
L589:
	;
	goto L590
L590:
	;
	if v3027 == int32(0) {
		goto L592
	} else {
		goto L593
	}
L591:
	;
	v3339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3338)+16)))
	v3340 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v3341 = m.ExcPending
	if v3341 != 0 {
		goto L1
	} else {
		goto L595
	}
L592:
	;
	v3324 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(v3324+(v2791^int32(-1))<<(uint(int32(2))%32))))
	v3338 = v3330
	goto L591
L593:
	;
	goto L594
L594:
	;
	v3332 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3338 = v3332 + v2791<<(uint(int32(13))%32) + int32(-8192)
	goto L591
L595:
	;
	v3342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3338)+16)))
	v3343 = v3338 + v3342
	v3344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3343)+12)))
	v3348 = v3344&int32(_a_F_btvacuumscan_27) | int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v3343)+12)) = uint16(v3348)
	v3350 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v3338)+12)) = uint16(v3350)
	*(*int64)(unsafe.Add(mBase, uint32(v3338)+24)) = v3340
	v3353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3338)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3338)+14)) = uint16(v3353)
	v3356 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3338+v3339)+14)) = uint16(v3356)
	if v3259 != 0 {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	v3358 = *(*int32)(unsafe.Add(mBase, uint32(v3260)+4))
	if base.Ui32(v3358) <= base.Ui32(int32(2)) {
		goto L599
	} else {
		goto L600
	}
L597:
	;
	goto L598
L598:
	;
	F_MarkBufferDirty(m, v3114)
	mBase = m.M
	v3376 = m.ExcPending
	if v3376 != 0 {
		goto L1
	} else {
		goto L603
	}
L599:
	;
	v3361 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3258)+64)) = uint8(v3361)
	*(*int64)(unsafe.Add(mBase, uint32(v3258)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v3258)+48)) = v3361
	*(*int32)(unsafe.Add(mBase, uint32(v3258)+28)) = int32(3)
	v3369 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v3258)+12)) = uint16(v3369)
	goto L601
L600:
	;
	goto L601
L601:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3260)+20)) = v2794
	*(*int32)(unsafe.Add(mBase, uint32(v3260)+16)) = v3048
	F_MarkBufferDirty(m, v3259)
	mBase = m.M
	v3374 = m.ExcPending
	if v3374 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	goto L598
L603:
	;
	F_MarkBufferDirty(m, v2791)
	mBase = m.M
	v3378 = m.ExcPending
	if v3378 != 0 {
		goto L1
	} else {
		goto L604
	}
L604:
	;
	if v2979 != 0 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	F_MarkBufferDirty(m, v2979)
	mBase = m.M
	v3380 = m.ExcPending
	if v3380 != 0 {
		goto L1
	} else {
		goto L608
	}
L606:
	;
	goto L607
L607:
	;
	if v2793 == int32(0) {
		goto L609
	} else {
		goto L610
	}
L608:
	;
	goto L607
L609:
	;
	F_MarkBufferDirty(m, v1856)
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L1
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	v3385 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	v3386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3385)+118)))
	if v3386 != int32(112) {
		goto L613
	} else {
		goto L614
	}
L612:
	;
	goto L611
L613:
	;
	v3551 = int32(_a_F_btvacuumscan_4)
	v3553 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v3553 - int32(1)
	if v3259 != 0 {
		goto L660
	} else {
		goto L661
	}
L614:
	;
	v3390 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if v3390 <= int32(0) {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v3393 = *(*int32)(unsafe.Add(mBase, uint32(v323)+32))
	if v3393 != 0 {
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
	v3396 = m.ExcPending
	if v3396 != 0 {
		goto L1
	} else {
		goto L620
	}
L618:
	;
	v3394 = *(*int32)(unsafe.Add(mBase, uint32(v323)+40))
	if v3394 != 0 {
		goto L613
	} else {
		goto L619
	}
L619:
	;
	goto L617
L620:
	;
	F_XLogRegisterBuffer(m, int32(0), v2791, int32(6))
	mBase = m.M
	v3400 = m.ExcPending
	if v3400 != 0 {
		goto L1
	} else {
		goto L621
	}
L621:
	;
	if v2979 != 0 {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	F_XLogRegisterBuffer(m, int32(1), v2979, int32(8))
	mBase = m.M
	v3404 = m.ExcPending
	if v3404 != 0 {
		goto L1
	} else {
		goto L625
	}
L623:
	;
	goto L624
L624:
	;
	F_XLogRegisterBuffer(m, int32(2), v3114, int32(8))
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L1
	} else {
		goto L626
	}
L625:
	;
	goto L624
L626:
	;
	if v2793 == int32(0) {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	F_XLogRegisterBuffer(m, int32(3), v1856, int32(6))
	mBase = m.M
	v3414 = m.ExcPending
	if v3414 != 0 {
		goto L1
	} else {
		goto L630
	}
L628:
	;
	goto L629
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+280)) = v3113
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+276)) = v2730
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+272)) = v2731
	*(*int64)(unsafe.Add(mBase, uint32(v1825)+264)) = v3340
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+256)) = v2794
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+252)) = v3048
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+248)) = v2981
	F_XLogRegisterData(m, v1825+int32(248), int32(36))
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L1
	} else {
		goto L631
	}
L630:
	;
	goto L629
L631:
	;
	if v3259 == int32(0) {
		goto L633
	} else {
		goto L634
	}
L632:
	;
	if v3122 == int32(0) {
		goto L641
	} else {
		goto L642
	}
L633:
	;
	v3431 = F_XLogInsert(m, int32(11), int32(128))
	mBase = m.M
	v3432 = m.ExcPending
	if v3432 != 0 {
		goto L1
	} else {
		goto L636
	}
L634:
	;
	goto L635
L635:
	;
	F_XLogRegisterBuffer(m, int32(4), v3259, int32(14))
	mBase = m.M
	v3436 = m.ExcPending
	if v3436 != 0 {
		goto L1
	} else {
		goto L637
	}
L636:
	;
	v3464 = v3431
	goto L632
L637:
	;
	v3437 = *(*int32)(unsafe.Add(mBase, uint32(v3260)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+220)) = v3437
	v3439 = *(*int32)(unsafe.Add(mBase, uint32(v3260)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+224)) = v3439
	v3441 = *(*int32)(unsafe.Add(mBase, uint32(v3260)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+228)) = v3441
	v3443 = *(*int32)(unsafe.Add(mBase, uint32(v3260)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+232)) = v3443
	v3445 = *(*int32)(unsafe.Add(mBase, uint32(v3260)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+236)) = v3445
	v3447 = *(*int32)(unsafe.Add(mBase, uint32(v3260)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+240)) = v3447
	v3449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3260)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1825)+244)) = uint8(v3449)
	F_XLogRegisterBufData(m, int32(4), v1825+int32(220), int32(28))
	mBase = m.M
	v3456 = m.ExcPending
	if v3456 != 0 {
		goto L1
	} else {
		goto L638
	}
L638:
	;
	v3459 = F_XLogInsert(m, int32(11), int32(144))
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3258))) = base.I64_rotr(v3459, int64(32))
	v3464 = v3459
	goto L632
L640:
	;
	v3483 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v3482))) = base.I64_rotr(v3464, v3483)
	v3488 = base.I32_wrap_i64(int64(base.Ui64(v3464) >> (uint(v3483) % 64)))
	v3489 = base.I32_wrap_i64(v3464)
	if v3027 == int32(0) {
		goto L645
	} else {
		goto L646
	}
L641:
	;
	v3468 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3474 = *(*int32)(unsafe.Add(mBase, uint32(v3468+(v3114^int32(-1))<<(uint(int32(2))%32))))
	v3482 = v3474
	goto L640
L642:
	;
	goto L643
L643:
	;
	v3476 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3482 = v3476 + v3114<<(uint(int32(13))%32) + int32(-8192)
	goto L640
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3507)+4)) = v3489
	*(*int32)(unsafe.Add(mBase, uint32(v3507))) = v3488
	if v2979 != 0 {
		goto L648
	} else {
		goto L649
	}
L645:
	;
	v3493 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3499 = *(*int32)(unsafe.Add(mBase, uint32(v3493+(v2791^int32(-1))<<(uint(int32(2))%32))))
	v3507 = v3499
	goto L644
L646:
	;
	goto L647
L647:
	;
	v3501 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3507 = v3501 + v2791<<(uint(int32(13))%32) + int32(-8192)
	goto L644
L648:
	;
	if v2979 < int32(0) {
		goto L652
	} else {
		goto L653
	}
L649:
	;
	goto L650
L650:
	;
	if v2793 != 0 {
		goto L613
	} else {
		goto L655
	}
L651:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3527)+4)) = v3489
	*(*int32)(unsafe.Add(mBase, uint32(v3527))) = v3488
	goto L650
L652:
	;
	v3513 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3519 = *(*int32)(unsafe.Add(mBase, uint32(v3513+(v2979^int32(-1))<<(uint(int32(2))%32))))
	v3527 = v3519
	goto L651
L653:
	;
	goto L654
L654:
	;
	v3521 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3527 = v3521 + v2979<<(uint(int32(13))%32) + int32(-8192)
	goto L651
L655:
	;
	if v1897 == int32(0) {
		goto L657
	} else {
		goto L658
	}
L656:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3544)+4)) = v3489
	*(*int32)(unsafe.Add(mBase, uint32(v3544))) = v3488
	goto L613
L657:
	;
	v3534 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3538 = *(*int32)(unsafe.Add(mBase, uint32(v3534+v2641<<(uint(int32(2))%32))))
	v3544 = v3538
	goto L656
L658:
	;
	goto L659
L659:
	;
	v3540 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3544 = v3540 + v2643 + int32(-8192)
	goto L656
L660:
	;
	F_LockBuffer(m, v3259, int32(0))
	mBase = m.M
	v3559 = m.ExcPending
	if v3559 != 0 {
		goto L1
	} else {
		goto L663
	}
L661:
	;
	goto L662
L662:
	;
	if v2979 != 0 {
		goto L665
	} else {
		goto L666
	}
L663:
	;
	F_ReleaseBuffer(m, v3259)
	mBase = m.M
	v3561 = m.ExcPending
	if v3561 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	goto L662
L665:
	;
	F_LockBuffer(m, v2979, int32(0))
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		goto L1
	} else {
		goto L668
	}
L666:
	;
	goto L667
L667:
	;
	F_LockBuffer(m, v3114, int32(0))
	mBase = m.M
	v3569 = m.ExcPending
	if v3569 != 0 {
		goto L1
	} else {
		goto L670
	}
L668:
	;
	F_ReleaseBuffer(m, v2979)
	mBase = m.M
	v3566 = m.ExcPending
	if v3566 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	goto L667
L670:
	;
	F_ReleaseBuffer(m, v3114)
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		goto L1
	} else {
		goto L671
	}
L671:
	;
	if v2793 == int32(0) {
		goto L672
	} else {
		goto L673
	}
L672:
	;
	F_LockBuffer(m, v2791, int32(0))
	mBase = m.M
	v3576 = m.ExcPending
	if v3576 != 0 {
		goto L1
	} else {
		goto L675
	}
L673:
	;
	goto L674
L674:
	;
	v3579 = *(*int32)(unsafe.Add(mBase, uint32(v2713)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2713)+24)) = v3579 + int32(1)
	if base.Ui32(v2792) <= base.Ui32(v1845) {
		goto L677
	} else {
		goto L678
	}
L675:
	;
	F_ReleaseBuffer(m, v2791)
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L1
	} else {
		goto L676
	}
L676:
	;
	goto L674
L677:
	;
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(v2713)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2713)+28)) = v3584 + int32(1)
	goto L679
L678:
	;
	goto L679
L679:
	;
	v3588 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+36))
	v3589 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+28))
	if v3588 != v3589 {
		goto L680
	} else {
		goto L681
	}
L680:
	;
	v3591 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+24))
	if v3591 != v3588 {
		goto L684
	} else {
		goto L685
	}
L681:
	;
	goto L682
L682:
	;
	v3628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1917)+12)))
	if v3628&int32(16) != 0 {
		goto L432
	} else {
		goto L691
	}
L683:
	;
	v3610 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v3609+v3608<<(uint(v3610)%32)))) = v2792
	v3614 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+32))
	v3615 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v3614+v3615<<(uint(v3610)%32))+8)) = v3340
	v3620 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+36)) = v3620 + int32(1)
	goto L682
L684:
	;
	v3593 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+32))
	v3608 = v3588
	v3609 = v3593
	goto L683
L685:
	;
	goto L686
L686:
	;
	v3595 = v3588 << (uint(int32(1)) % 32)
	if v3595 < v3589 {
		goto L687
	} else {
		goto L688
	}
L687:
	;
	v3597 = v3595
	goto L689
L688:
	;
	v3597 = v3589
	goto L689
L689:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+24)) = v3597
	v3599 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+32))
	v3602 = F_repalloc(m, v3599, v3597<<(uint(int32(4))%32))
	mBase = m.M
	v3603 = m.ExcPending
	if v3603 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+32)) = v3602
	v3605 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+36))
	v3608 = v3605
	v3609 = v3602
	goto L683
L691:
	;
	goto L433
L692:
	;
	v3639 = v3633
	goto L694
L693:
	;
	v3639 = int32(1)
	goto L694
L694:
	;
	v3652 = base.B2i32(base.Ui32(int32(base.Ui32(v3193+int32(_a_F_btvacuumscan_6))>>(uint(v3633)%32))&int32(_a_F_btvacuumscan_5)) < base.Ui32(v3639)) | base.B2i32(base.Ui32(v3193) < base.Ui32(int32(25)))
	goto L431
L695:
	;
	F_ReleaseBuffer(m, v1856)
	mBase = m.M
	v3699 = m.ExcPending
	if v3699 != 0 {
		goto L1
	} else {
		goto L696
	}
L696:
	;
	v3701 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v3701 != 0 {
		goto L697
	} else {
		goto L698
	}
L697:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3703 = m.ExcPending
	if v3703 != 0 {
		goto L1
	} else {
		goto L700
	}
L698:
	;
	goto L699
L699:
	;
	if v3652 == int32(0) {
		goto L254
	} else {
		goto L701
	}
L700:
	;
	goto L699
L701:
	;
	v3706 = F_ReadBuffer(m, v323, v3694)
	mBase = m.M
	v3707 = m.ExcPending
	if v3707 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	F_LockBuffer(m, v3706, int32(2))
	mBase = m.M
	v3710 = m.ExcPending
	if v3710 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	F__bt_checkpage(m, v323, v3706)
	mBase = m.M
	v3712 = m.ExcPending
	if v3712 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	v1856 = v3706
	goto L261
L705:
	;
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_28), int32(0))
	mBase = m.M
	v3720 = m.ExcPending
	if v3720 != 0 {
		goto L1
	} else {
		goto L706
	}
L706:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2244), int32(_a_F_btvacuumscan_20))
	mBase = m.M
	v3725 = m.ExcPending
	if v3725 != 0 {
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
	v3731 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+16)) = v2792
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+20)) = v3731 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_29), v1825+int32(16))
	mBase = m.M
	v3740 = m.ExcPending
	if v3740 != 0 {
		goto L1
	} else {
		goto L709
	}
L709:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2472), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3745 = m.ExcPending
	if v3745 != 0 {
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
	v3752 = m.ExcPending
	if v3752 != 0 {
		goto L1
	} else {
		goto L712
	}
L712:
	;
	v3753 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	v3754 = *(*int32)(unsafe.Add(mBase, uint32(v3047)))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+104)) = v2792
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+100)) = v3754
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+96)) = v2981
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+108)) = v3753 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_30), v1825+int32(96))
	mBase = m.M
	v3765 = m.ExcPending
	if v3765 != 0 {
		goto L1
	} else {
		goto L713
	}
L713:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2479), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3770 = m.ExcPending
	if v3770 != 0 {
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
	v3775 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+84)) = v2792
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+80)) = v2794
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+88)) = v3775 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_31), v1825+int32(80))
	mBase = m.M
	v3785 = m.ExcPending
	if v3785 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2498), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3790 = m.ExcPending
	if v3790 != 0 {
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
	if v3843 == int32(0) {
		goto L256
	} else {
		goto L719
	}
L719:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3849 = m.ExcPending
	if v3849 != 0 {
		goto L1
	} else {
		goto L720
	}
L720:
	;
	v3850 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+4)) = v3795
	*(*int32)(unsafe.Add(mBase, uint32(v1825))) = v3850 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_32), v1825)
	mBase = m.M
	v3857 = m.ExcPending
	if v3857 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2848), int32(_a_F_btvacuumscan_33))
	mBase = m.M
	v3862 = m.ExcPending
	if v3862 != 0 {
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
	v4104 = v4052
	goto L54
L726:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v4129 = m.ExcPending
	if v4129 != 0 {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	v4130 = int32(0)
	v4132 = *(*int32)(unsafe.Add(mBase, uint32(v349)+24))
	v4133 = F_ReadBufferExtended(m, v323, v4130, v4104, v4130, v4132)
	mBase = m.M
	v4134 = m.ExcPending
	if v4134 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	v317 = v4133
	v324 = v4104
	goto L51
L729:
	;
	v125 = v233
	v126 = v234
	v137 = v245
	v150 = v258
	v156 = v264
	v159 = v267
	goto L17
L730:
	;
	if v4141 == int32(0) {
		goto L41
	} else {
		goto L731
	}
L731:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v4147 = m.ExcPending
	if v4147 != 0 {
		goto L1
	} else {
		goto L732
	}
L732:
	;
	v4148 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v326)+4)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v326))) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v326)+8)) = v4148 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_34), v326)
	mBase = m.M
	v4156 = m.ExcPending
	if v4156 != 0 {
		goto L1
	} else {
		goto L733
	}
L733:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_35), int32(1413), int32(_a_F_btvacuumscan_36))
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	goto L41
L735:
	;
	v4166 = v314
	v4167 = v315
	v4178 = v326
	v4191 = v339
	v4197 = v345
	v4200 = v348
	goto L40
L736:
	;
	v4223 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[2]))
	if v4223 == int32(0) {
		goto L738
	} else {
		goto L739
	}
L737:
	;
	v233 = v4166
	v234 = v4167
	v245 = v4178
	v258 = v4191
	v264 = v4197
	v267 = v4200
	goto L37
L738:
	;
	goto L737
L739:
	;
	v4227 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btvacuumscan[3])))
	if v4227&int32(1) == int32(0) {
		goto L738
	} else {
		goto L740
	}
L740:
	;
	v4232 = int32(_a_F_btvacuumscan_4)
	v4234 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	v4235 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v4234 + v4235
	v4238 = *(*int32)(unsafe.Add(mBase, uint32(v4223)))
	*(*int32)(unsafe.Add(mBase, uint32(v4223))) = v4238 + v4235
	*(*int64)(unsafe.Add(mBase, uint32(v4223+int32(128))+232)) = base.I64_extend_i32_u(v341)
	v4246 = *(*int32)(unsafe.Add(mBase, uint32(v4223)))
	*(*int32)(unsafe.Add(mBase, uint32(v4223))) = v4246 + v4235
	v4252 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v4252 - v4235
	goto L738
L741:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v189
	v4259 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	F_MemoryContextDelete(m, v4259)
	mBase = m.M
	v4261 = m.ExcPending
	if v4261 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	v4262 = int32(0)
	v4264 = v137 + int32(24)
	v4265 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+36))
	if v4265 == v4262 {
		goto L745
	} else {
		goto L746
	}
L743:
	;
	v4500 = *(*int32)(unsafe.Add(mBase, uint32(v126)+32))
	if v4500 != 0 {
		goto L759
	} else {
		goto L760
	}
L744:
	;
	F_pfree(m, v4400)
	mBase = m.M
	v4449 = m.ExcPending
	if v4449 != 0 {
		goto L1
	} else {
		goto L758
	}
L745:
	;
	v4268 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+32))
	if v4268 != 0 {
		v4400 = v4268
		goto L744
	} else {
		goto L748
	}
L746:
	;
	goto L747
L747:
	;
	v4269 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+4))
	v4270 = *(*int32)(unsafe.Add(mBase, uint32(v4264)))
	v4271 = *(*int32)(unsafe.Add(mBase, uint32(v4270)+4))
	v4272 = F_GetOldestNonRemovableTransactionId(m, v4271)
	mBase = m.M
	v4273 = m.ExcPending
	if v4273 != 0 {
		goto L1
	} else {
		goto L749
	}
L748:
	;
	goto L743
L749:
	;
	v4274 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+36))
	if v4274 <= int32(0) {
		goto L750
	} else {
		goto L751
	}
L750:
	;
	v4397 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+32))
	v4400 = v4397
	goto L744
L751:
	;
	v4279 = v4262
	goto L752
L752:
	;
	v4327 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+32))
	v4330 = v4327 + v4279<<(uint(int32(4))%32)
	v4331 = *(*int32)(unsafe.Add(mBase, uint32(v4330)))
	v4332 = *(*int64)(unsafe.Add(mBase, uint32(v4330)+8))
	v4333 = F_GlobalVisCheckRemovableFullXid(m, v4271, v4332)
	mBase = m.M
	v4334 = m.ExcPending
	if v4334 != 0 {
		goto L1
	} else {
		goto L754
	}
L753:
	;
	goto L750
L754:
	;
	if v4333 == int32(0) {
		goto L750
	} else {
		goto L755
	}
L755:
	;
	F_RecordFreeIndexPage(m, v150, v4331)
	mBase = m.M
	v4338 = m.ExcPending
	if v4338 != 0 {
		goto L1
	} else {
		goto L756
	}
L756:
	;
	v4339 = *(*int32)(unsafe.Add(mBase, uint32(v4269)+32))
	v4340 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4269)+32)) = v4339 + v4340
	v4344 = v4279 + v4340
	v4345 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+36))
	if v4344 < v4345 {
		v4279 = v4344
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
	v4502 = m.ExcPending
	if v4502 != 0 {
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
	v5 = Fn13857(m, l0, l1, int32(21), int32(1))
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
