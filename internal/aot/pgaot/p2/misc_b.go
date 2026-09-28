package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BTreeShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v3 = F_time(m)
	mBase = m.M
	v4 = int32(_a_F_BTreeShmemInit_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_BTreeShmemInit[0]))
	*(*uint16)(unsafe.Add(mBase, uint32(v5))) = uint16(v3)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_BTreeShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_BTreeShmemInit[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	return
}
func F_BarrierArriveAndWait(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
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
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	v7 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_BarrierArriveAndWait_0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = int32(1)
	v15 = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = v17 + v14
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v20 == v15 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v24
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v24))
	F_ConditionVariableBroadcast(m, l0+int32(24))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v35 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v35))
	v39 = l0 + int32(24)
	F_ConditionVariablePrepareToSleep(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v48 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_s_lock(m, l0, int32(_a_F_BarrierArriveAndWait_0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 == v52 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v19 != v54 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v64 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v64))
	F_ConditionVariableSleep(m, v39, l1)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L24
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19
	goto L22
L21:
	;
	goto L22
L22:
	;
	v57 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v57))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	return base.B2i32(v54 != v19)
L24:
	;
	goto L11
}
func F_BarrierDetach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		F_s_lock(m, l0, int32(_a_F_BarrierDetach_0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v11 = v9 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
			if int32(0) < v11 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v15 == v11 {
					v20 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22 + int32(1)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v20))
					F_ConditionVariableBroadcast(m, l0+int32(24))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						return
					}
				} else {
					v17 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v17))
					return
				}
			} else {
				v17 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v17))
				return
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v11 = v9 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
		if int32(0) < v11 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v15 == v11 {
				v20 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22 + int32(1)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v20))
				F_ConditionVariableBroadcast(m, l0+int32(24))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					return
				}
			} else {
				v17 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v17))
				return
			}
		} else {
			v17 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v17))
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
	F_errfinish(m, int32(_a_F_BasicOpenFilePerm_1), int32(1153), int32(_a_F_BasicOpenFilePerm_2))
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
			F_errfinish(m, int32(_a_F_BogusFree_1), int32(312), int32(_a_F_BogusFree_2))
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
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 float64
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
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
	var v129 float64
	_ = v129
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(_a_F_BuildCallback_2_0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_2[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l5)+168))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_2[0])) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l5)+68))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	return
L5:
	;
	v26 = base.I64_extend_i32_u(v24)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l5)+52))
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_2[0])) = v17
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l5)+168))
	F_MemoryContextReset(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L36
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
	v29 = F_IvfflatCheckNorm(m, v27, v28, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	v37 = v26
	goto L9
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v38 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v29 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
	v35 = F_HnswNormValue(m, v33, v34, v26)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v37 = v35
	goto L9
L13:
	;
	v102 = int64(0)
	goto L15
L14:
	;
	v43 = int32(0)
	v47 = v43
	v48 = v43
	v53 = float64(1.7976931348623157e+308)
	goto L17
L15:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	m.T0[v104].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L33
	}
L16:
	;
	v102 = base.I64_extend_i32_u(v71)
	goto L15
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v48 < v57 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L30
	}
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v66 = F_FunctionCall2Coll(m, v59, v60, v37, base.I64_extend_i32_u(v61+v62*v48))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	v68 = base.F64_reinterpret_i64(v66)
	v69 = base.F64_gt(v53, v68)
	if v69 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v70 = v68
	goto L25
L24:
	;
	v70 = v53
	goto L25
L25:
	;
	if v69 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v71 = v48
	goto L28
L27:
	;
	v71 = v47
	goto L28
L28:
	;
	v73 = v48 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v73 < v74 {
		v47 = v71
		v48 = v73
		v53 = v70
		goto L17
	} else {
		goto L29
	}
L29:
	;
	goto L16
L30:
	;
	F_errmsg_internal(m, int32(_a_F_BuildCallback_2_1), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_BuildCallback_2_2), int32(326), int32(_a_F_BuildCallback_2_3))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
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
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v107))) = v102
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v110)
	*(*int64)(unsafe.Add(mBase, uint32(v107)+8)) = base.I64_extend_i32_u(l1)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)) = uint8(v110)
	*(*int64)(unsafe.Add(mBase, uint32(v107)+16)) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+2)) = uint8(v110)
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v121 = v119 & int32(_a_F_BuildCallback_2_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v121)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v124)
	goto L34
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l5)+152))
	F_tuplesort_puttupleslot(m, v126, v21)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v129 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+32)) = base.F64_add(v129, float64(1))
	goto L6
L36:
	;
	goto L3
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
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
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
	v179 = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v179 < v188 {
		goto L45
	} else {
		goto L46
	}
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
	*(*int32)(unsafe.Add(mBase, uint32(v24+v85<<(uint(int32(3))%32)+v80*int32(100))+24)) = v67
	goto L26
L26:
	;
	v99 = v24 + v85<<(uint(int32(3))%32) + v80*int32(100)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+14)) = uint8(v100)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+20)) = uint8(v102)
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+22)) = uint16(v104)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+17)) = uint8(v106)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+18)) = uint8(v108)
	v111 = v99 - int32(72)
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
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+13)) = uint8(v114)
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
	F_errfinish(m, int32(_a_F_BuildDescForRelation_2), int32(1454), int32(_a_F_BuildDescForRelation_3))
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
	F_errfinish(m, int32(_a_F_BuildDescForRelation_2), int32(1460), int32(_a_F_BuildDescForRelation_3))
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
L44:
	;
	m.G0 = v14 + int32(16)
	return v172
L45:
	;
	v192 = v172 + int32(28)
	v199 = v179
	v200 = v188
	v202 = v179
	goto L49
L46:
	;
	v256 = v179
	v263 = v188
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v172)+16)) = v256
	goto L44
L48:
	;
	v256 = v250
	v263 = v229
	goto L47
L49:
	;
	v208 = v192 + v188<<(uint(int32(3))%32) + v199*int32(100)
	v211 = v192 + v199<<(uint(int32(3))%32)
	if v188 != v200 {
		v229 = v200
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v250 = v188
	goto L48
L51:
	;
	v230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v211)+2)))
	if v230 <= int32(0) {
		v250 = v199
		goto L48
	} else {
		goto L59
	}
L52:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+7)))
	if v213 != int32(118) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v229 = v199
	goto L51
L54:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
	if v216 != int32(1) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+6)))
	if v219&int32(6) != 0 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v222 = int32(*(*int16)(unsafe.Add(mBase, uint32(v211)+2)))
	if v222 <= int32(0) {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+90)))
	if v225 != int32(118) {
		v229 = v188
		goto L51
	} else {
		goto L58
	}
L58:
	;
	goto L53
L59:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+90)))
	if v233 == int32(118) {
		v250 = v199
		goto L48
	} else {
		goto L60
	}
L60:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+5)))
	v242 = (v202 + v236 - int32(1)) & (int32(0) - v236)
	if int32(_a_F_BuildDescForRelation_5) < v242 {
		v250 = v199
		goto L48
	} else {
		goto L61
	}
L61:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v211))) = uint16(v242)
	v248 = v199 + int32(1)
	if v248 != v188 {
		v199 = v248
		v200 = v229
		v202 = v242 + v230
		goto L49
	} else {
		goto L62
	}
L62:
	;
	goto L50
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
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
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
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
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
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
func F_base32hex_enc_len(m *base.Module, l0 int32, l1 int32) int64 {
	var v7 int64
	_ = v7
	v7 = base.I64_div_u_s(base.I64_extend_i32_u(l1)+int64(4), int64(5))
	return v7 << (uint(int64(3)) % 64)
}
func F_begin_tup_output_tupdesc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v9 = F_palloc(m, int32(8))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if l1 == int32(0) {
			v32 = int32(2)
			v33 = v13
		} else {
			v18 = int32(7)
			v20 = int32(-8)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v32 = int32(18)
			v33 = (v13+v18)&v20 + v22<<(uint(int32(3))%32) + (v22+v18)&v20
		}
		v34 = F_palloc0(m, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = l1
			*(*uint16)(unsafe.Add(mBase, uint32(v34)+4)) = uint16(v32)
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(449)
			*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = l2
			v42 = *(*int32)(unsafe.Add(mBase, _c_F_begin_tup_output_tupdesc[0]))
			v43 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v34)+6)) = uint16(v43)
			*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v42
			if l1 != 0 {
				v50 = v34 + (v13+int32(7))&int32(-8)
				*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v50 + v52<<(uint(int32(3))%32)
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if int32(0) <= v57 {
					F_IncrTupleDescRefCount(m, l1)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
						v63 = v62
						*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(0)
						v66 = v63
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
						m.T0[v68].(func(*base.Module, int32))(m, v34)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v34
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							m.T0[v74].(func(*base.Module, int32, int32, int32))(m, l0, int32(1), l1)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								return v9
							}
						}
					}
				} else {
					v63 = l2
					*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(0)
					v66 = v63
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
					m.T0[v68].(func(*base.Module, int32))(m, v34)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v34
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						m.T0[v74].(func(*base.Module, int32, int32, int32))(m, l0, int32(1), l1)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							return v9
						}
					}
				}
			} else {
				v66 = l2
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
				m.T0[v68].(func(*base.Module, int32))(m, v34)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v34
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					m.T0[v74].(func(*base.Module, int32, int32, int32))(m, l0, int32(1), l1)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						return v9
					}
				}
			}
		}
	}
}
func F_bitfromint4(m *base.Module, l0 int32) int64 {
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
		return int64(0)
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
		return base.I64_extend_i32_u(v24)
	}
}
func F_bitfromint8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v127 int64
	_ = v127
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = v14 - int32(2147483641)
	if base.Ui32(v16) < base.Ui32(int32(-2147483640)) {
		v19 = int32(1)
	} else {
		v19 = v14
	}
	v22 = int32(8)
	v23 = base.I32_div_s(v19+int32(7), v22)
	v25 = v23 + v22
	v26 = F_palloc(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = v25 << (uint(int32(2)) % 32)
		v35 = v26 + int32(8)
		v36 = int32(64)
		if v36 <= v19 {
			v39 = v36
		} else {
			v39 = v19
		}
		v41 = v39 + int32(8)
		if v41 <= v19 {
			v43 = int32(-2147483640)
			if base.Ui32(v16) <= base.Ui32(v43) {
				v46 = v43
			} else {
				v46 = v16
			}
			v48 = v46 + int32(2147483633)
			v49 = v48 - v39
			v51 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
			v53 = v51 + int32(1)
			if v53 != 0 {
				base.MemoryFill(m, v35, base.I32_wrap_i64(v12>>(uint(int64(63))%64)), v53)
			} else {
			}
			v64 = v26 + v51 + int32(9)
			v65 = v48 - v49&int32(-8)
		} else {
			v64 = v35
			v65 = v19
		}
		if v39 < v65 {
			v78 = v65 - int32(8)
			v82 = base.I32_wrap_i64(v12>>(uint(int64(63))%64))&(int32(-1)<<(uint(v41-v65)%32)) | base.I32_wrap_i64(v12>>(uint(base.I64_extend_i32_u(v78))%64))
			*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v82)
			v86 = v64 + int32(1)
			v87 = v78
		} else {
			v86 = v64
			v87 = v65
		}
		if int32(8) <= v87 {
			v91 = v86
			v100 = base.I64_extend_i32_u(v87)
			for {
				v103 = v100 - int64(8)
				v104 = v12 >> (uint(v103) % 64)
				*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v104)
				v107 = v91 + int32(1)
				if base.Ui64(int64(15)) < base.Ui64(v100) {
					v91 = v107
					v100 = v103
					continue
				} else {
					break
				}
				break
			}
			v111 = v107
			v112 = base.I32_wrap_i64(v103)
		} else {
			v111 = v86
			v112 = v87
		}
		if int32(0) < v112 {
			v127 = v12 << (uint(base.I64_extend_i32_u(int32(8)-v112)) % 64)
			*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v127)
		} else {
		}
		return base.I64_extend_i32_u(v26)
	}
}
func F_bitgt(m *base.Module, l0 int32) int64 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int64(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(base.B2i32(int32(0) < v99))
L38:
	;
	goto L37
}
func F_bitle(m *base.Module, l0 int32) int64 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int64(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(base.B2i32(v99 <= int32(0)))
L38:
	;
	goto L37
}
func F_bitshiftleft(m *base.Module, l0 int32) int64 {
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v16 < int32(0) {
			v20 = int32(0)
			v23 = int32(-2147483640)
			if base.Ui32(v16) <= base.Ui32(v23) {
				v26 = v23
			} else {
				v26 = v16
			}
			v29 = F_DirectFunctionCall2Coll(m, int32(1746), v20, base.I64_extend_i32_u(v12), base.I64_extend_i32_u(v20-v26))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				return v29
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v35 = F_palloc(m, int32(base.Ui32(v32)>>(uint(int32(2))%32)))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int64(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v39 = v37 & int32(-4)
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v41
				v44 = v35 + int32(8)
				if v41 <= v16 {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v48 = int32(base.Ui32(v46) >> (uint(int32(2)) % 32))
					v50 = v48 - int32(8)
					if v44&int32(3)|v46&int32(12)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v50)) == int32(0) {
						v61 = v35 + v48
						if base.Ui32(v61) <= base.Ui32(v44) {
							return base.I64_extend_i32_u(v35)
						} else {
							v64 = v35 + int32(12)
							if base.Ui32(v64) < base.Ui32(v61) {
								v66 = v61
							} else {
								v66 = v64
							}
							v73 = (v66-v35-int32(9))&int32(-4) + int32(4)
							if v73 == int32(0) {
								return base.I64_extend_i32_u(v35)
							} else {
								v207 = v44
								v208 = v73
								base.MemoryFill(m, v207, int32(0), v208)
								return base.I64_extend_i32_u(v35)
							}
						}
					} else {
						if v50 == int32(0) {
							return base.I64_extend_i32_u(v35)
						} else {
							v207 = v44
							v208 = v50
							base.MemoryFill(m, v207, int32(0), v208)
							return base.I64_extend_i32_u(v35)
						}
					}
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v80 = int32(base.Ui32(v78) >> (uint(int32(2)) % 32))
					v82 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
					v84 = v82 + int32(8)
					v85 = v12 + v84
					v87 = v16 & int32(7)
					if v87 != 0 {
						if base.Ui32(v84) < base.Ui32(v80) {
							v91 = v44
							v95 = v85
							for {
								v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
								v102 = v101 << (uint(v87) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v102)
								v105 = v95 + int32(1)
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v108 = int32(base.Ui32(v106) >> (uint(int32(2)) % 32))
								if base.Ui32(v105) < base.Ui32(v12+v108) {
									v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
									v113 = int32(base.Ui32(v111)>>(uint(int32(8)-v87)%32)) | v102
									*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v113)
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v118 = int32(base.Ui32(v115) >> (uint(int32(2)) % 32))
								} else {
									v118 = v108
								}
								v120 = v91 + int32(1)
								if base.Ui32(v105) < base.Ui32(v12+v118) {
									v91 = v120
									v95 = v105
									continue
								} else {
									break
								}
								break
							}
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
							v124 = v120
							v134 = v123
						} else {
							v124 = v44
							v134 = v39
						}
						if base.Ui32(v35+int32(base.Ui32(v134)>>(uint(int32(2))%32))) <= base.Ui32(v124) {
						} else {
							v139 = v124
							for {
								v149 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v149)
								v152 = v139 + int32(1)
								v153 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
								if base.Ui32(v152) < base.Ui32(v35+int32(base.Ui32(v153)>>(uint(int32(2))%32))) {
									v139 = v152
									continue
								} else {
									break
								}
								break
							}
						}
						return base.I64_extend_i32_u(v35)
					} else {
						v158 = v80 - v82
						v160 = v158 - int32(8)
						if v160 != 0 {
							base.MemoryCopy(m, v44, v85, v160)
						} else {
						}
						v164 = v158 + v35
						if v16&int32(24)|(v164&int32(3)|base.B2i32(base.Ui32(int32(_a_F_bitshiftleft_0)) < base.Ui32(v16))) == int32(0) {
							v173 = v35 + v80
							if base.Ui32(v173) <= base.Ui32(v164) {
								return base.I64_extend_i32_u(v35)
							} else {
								v177 = v173 - v82 + int32(4)
								if base.Ui32(v173) < base.Ui32(v177) {
									v179 = v177
								} else {
									v179 = v173
								}
								v188 = (v179+v82+(v35^int32(-1))-v80)&int32(-4) + int32(4)
								if v188 == int32(0) {
									return base.I64_extend_i32_u(v35)
								} else {
									v207 = v164
									v208 = v188
									base.MemoryFill(m, v207, int32(0), v208)
									return base.I64_extend_i32_u(v35)
								}
							}
						} else {
							if v82 == int32(0) {
							} else {
								base.MemoryFill(m, v164, int32(0), v82)
							}
							return base.I64_extend_i32_u(v35)
						}
					}
				}
			}
		}
	}
}
func F_bittypmodout(m *base.Module, l0 int32) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(64))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		if int32(0) <= v8 {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
			v19 = F_pg_snprintf(m, v10, int32(64), int32(_a_F_bittypmodout_0), v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v10)
			}
		} else {
			v21 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v21)
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v10)
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
func F_blhandler(m *base.Module, l0 int32) int64 {
	return int64(4521744)
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
					v18 = v14 * int32(56)
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
				v18 = v14 * int32(56)
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
func F_booleq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(0)
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3) ^ base.B2i32(v5 != v3))
}
func F_boolge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(0)
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3) | base.B2i32(v5 != v3))
}
func F_boolout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_palloc(m, int32(2))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v9)
		if v3 == int64(0) {
			v15 = int32(102)
		} else {
			v15 = int32(116)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v15)
		return base.I64_extend_i32_u(v5)
	}
}
func F_boot_get_type_io_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	v10 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_boot_get_type_io_data[0]))
	if v19 == v10 {
		goto L27
	} else {
		goto L28
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L48
	} else {
		goto L55
	}
L2:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v158
	m.G0 = v16 + int32(32)
	return
L3:
	;
	v120 = v118 * int32(92)
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_boot_get_type_io_data[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v121)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_boot_get_type_io_data[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v123)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_boot_get_type_io_data[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v125)
	v127 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v127)
	if int32(1)<<(uint(v118)%32)&int32(_a_F_boot_get_type_io_data_0) != 0 {
		goto L52
	} else {
		goto L53
	}
L4:
	;
	v118 = int32(22)
	goto L3
L5:
	;
	v118 = int32(21)
	goto L3
L6:
	;
	v118 = int32(20)
	goto L3
L7:
	;
	v118 = int32(19)
	goto L3
L8:
	;
	v118 = int32(18)
	goto L3
L9:
	;
	v118 = int32(17)
	goto L3
L10:
	;
	v118 = int32(16)
	goto L3
L11:
	;
	v118 = int32(15)
	goto L3
L12:
	;
	v118 = int32(14)
	goto L3
L13:
	;
	v118 = int32(13)
	goto L3
L14:
	;
	v118 = int32(11)
	goto L3
L15:
	;
	v118 = int32(10)
	goto L3
L16:
	;
	v118 = int32(9)
	goto L3
L17:
	;
	v118 = int32(8)
	goto L3
L18:
	;
	v118 = int32(7)
	goto L3
L19:
	;
	v118 = int32(6)
	goto L3
L20:
	;
	v118 = int32(5)
	goto L3
L21:
	;
	v118 = int32(4)
	goto L3
L22:
	;
	v118 = int32(3)
	goto L3
L23:
	;
	v118 = int32(2)
	goto L3
L24:
	;
	v118 = int32(1)
	goto L3
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L48
	} else {
		goto L49
	}
L26:
	;
	if l0 == int32(194) {
		goto L11
	} else {
		goto L47
	}
L27:
	;
	if l0 <= int32(1001) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v35 <= int32(0) {
		goto L1
	} else {
		goto L37
	}
L30:
	;
	switch l0 - int32(16) {
	case 0:
		v118 = v10
		goto L3
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L16
	case 4:
		goto L19
	case 5:
		goto L21
	case 6:
		goto L10
	case 7:
		goto L20
	case 8:
		goto L15
	case 9:
		goto L14
	case 10:
		goto L13
	case 11, 12, 13:
		goto L25
	case 14:
		goto L9
	default:
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	switch l0 - int32(1002) {
	case 0:
		goto L5
	case 1, 2, 3, 4, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 27, 28, 29, 30:
		goto L25
	case 5:
		goto L8
	case 7:
		goto L7
	case 26:
		goto L6
	case 31:
		goto L12
	case 32:
		goto L4
	default:
		goto L34
	}
L33:
	;
	switch l0 - int32(700) {
	case 0:
		goto L18
	case 1:
		goto L17
	default:
		goto L26
	}
L34:
	;
	if l0 == int32(2275) {
		goto L22
	} else {
		goto L35
	}
L35:
	;
	if l0 != int32(3802) {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v118 = int32(12)
	goto L3
L37:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v48 = int32(0)
	goto L39
L38:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v65)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v67)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+132)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v69)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+87)))
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v71)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+96))
	if v73 != 0 {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v38+v48<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v57 == l0 {
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v62 != l0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v60 = v48 + int32(1)
	if v60 != v35 {
		v48 = v60
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L38
L44:
	;
	v74 = v73
	goto L46
L45:
	;
	v74 = l0
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v56)+104))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v56)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v78
	v157 = v56 + int32(148)
	goto L2
L47:
	;
	goto L25
L48:
	;
	return
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg_internal(m, int32(_a_F_boot_get_type_io_data_1), v16)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_boot_get_type_io_data_2), int32(1064), int32(_a_F_boot_get_type_io_data_3))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v136 = l0
	goto L54
L53:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_boot_get_type_io_data[4])))
	v136 = v135
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_boot_get_type_io_data[5])))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_boot_get_type_io_data[6])))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v140
	v157 = v120 + int32(_a_F_boot_get_type_io_data_4)
	goto L2
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_boot_get_type_io_data_5), v16+int32(16))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L48
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_boot_get_type_io_data_2), int32(1035), int32(_a_F_boot_get_type_io_data_3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L48
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bpchareq(m *base.Module, l0 int32) int64 {
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
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
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
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
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
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v23 = v21 & v19
	if v23 != 0 {
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
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L94
	}
L7:
	;
	v24 = v19
	goto L9
L8:
	;
	v24 = int32(4)
	goto L9
L9:
	;
	if v21 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v57 = v52
	goto L21
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v31 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v42 = int32(1)
	if v23 != 0 {
		v52 = int32(base.Ui32(v21)>>(uint(v42)%32)) - v42
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v34 = int32(16)
	goto L16
L15:
	;
	v34 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v41 = int32(4)
	goto L19
L18:
	;
	v41 = v34
	goto L19
L19:
	;
	v52 = v41
	goto L10
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	if v57 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v75 = int32(1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v79 = v77 & v75
	if v79 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v74 = v52 & (v52 >> (uint(int32(31)) % 32))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v68 = v57 - int32(1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v11+v68))))
	if v70 == int32(32) {
		v57 = v68
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v74 = v57
	goto L23
L28:
	;
	v80 = v75
	goto L30
L29:
	;
	v80 = int32(4)
	goto L30
L30:
	;
	if v77 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v113 = v108
	goto L42
L32:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v87 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v98 = int32(1)
	if v79 != 0 {
		v108 = int32(base.Ui32(v77)>>(uint(v98)%32)) - v98
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v90 = int32(16)
	goto L37
L36:
	;
	v90 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v87-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v97 = int32(4)
	goto L40
L39:
	;
	v97 = v90
	goto L40
L40:
	;
	v108 = v97
	goto L31
L41:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v108 = int32(base.Ui32(v102)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	if v113 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v131 = F_pg_newlocale_from_collation(m, v18)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	goto L43
L45:
	;
	v130 = v108 & (v108 >> (uint(int32(31)) % 32))
	goto L44
L46:
	;
	goto L47
L47:
	;
	v124 = v113 - int32(1)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v16+v124))))
	if v126 == int32(32) {
		v113 = v124
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v130 = v113
	goto L44
L49:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v235 != v11 {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v133 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v130 != v74 {
		v234 = int32(0)
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v216 = int32(1)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v218&v216 != 0 {
		goto L79
	} else {
		goto L80
	}
L54:
	;
	v138 = int32(1)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v140&v138 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v143 = v138
	goto L57
L56:
	;
	v143 = int32(4)
	goto L57
L57:
	;
	v144 = v11 + v143
	v145 = int32(1)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v147&v145 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v150 = v145
	goto L60
L59:
	;
	v150 = int32(4)
	goto L60
L60:
	;
	v151 = v16 + v150
	if base.Ui32(int32(4)) <= base.Ui32(v74) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v234 = base.B2i32(v213 == int32(0))
	goto L49
L62:
	;
	v213 = int32(0)
	goto L61
L63:
	;
	v187 = v182
	v188 = v183
	v189 = v184
	goto L73
L64:
	;
	if (v144|v151)&int32(3) != 0 {
		v182 = v144
		v183 = v151
		v184 = v74
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v175 = v144
	v176 = v151
	v177 = v74
	goto L66
L66:
	;
	if v177 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v159 = v144
	v160 = v151
	v161 = v74
	goto L68
L68:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v164 != v165 {
		v182 = v159
		v183 = v160
		v184 = v161
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v175 = v170
	v176 = v168
	v177 = v172
	goto L66
L70:
	;
	v167 = int32(4)
	v168 = v160 + v167
	v170 = v159 + v167
	v172 = v161 - v167
	if base.Ui32(int32(3)) < base.Ui32(v172) {
		v159 = v170
		v160 = v168
		v161 = v172
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v182 = v175
	v183 = v176
	v184 = v177
	goto L63
L73:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v192 == v193 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v213 = v192 - v193
	goto L61
L75:
	;
	v195 = int32(1)
	v200 = v189 - v195
	if v200 != 0 {
		v187 = v187 + v195
		v188 = v188 + v195
		v189 = v200
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
	v221 = v216
	goto L81
L80:
	;
	v221 = int32(4)
	goto L81
L81:
	;
	v223 = int32(1)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v225&v223 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v228 = v223
	goto L84
L83:
	;
	v228 = int32(4)
	goto L84
L84:
	;
	v230 = F_varstr_cmp(m, v11+v221, v74, v16+v228, v130, v18)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v234 = base.B2i32(v230 == int32(0))
	goto L49
L86:
	;
	F_pfree(m, v11)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v239 != v16 {
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
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	return base.I64_extend_i32_u(v234)
L93:
	;
	goto L92
L94:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_bpchareq_0), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errhint(m, int32(_a_F_bpchareq_1), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_bpchareq_2), int32(741), int32(_a_F_bpchareq_3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
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
func F_bpcharge(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
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
	v18 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v22 = v20 & v18
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v18
	goto L6
L5:
	;
	v23 = int32(4)
	goto L6
L6:
	;
	v24 = v23 + v11
	if v20 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = v51
	goto L18
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v30 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v41 = int32(1)
	if v22 != 0 {
		v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v33 = int32(16)
	goto L13
L12:
	;
	v33 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = int32(4)
	goto L16
L15:
	;
	v40 = v33
	goto L16
L16:
	;
	v51 = v40
	goto L7
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v57 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v74 = int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v78 = v76 & v74
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v73 = v51 & (v51 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v67 = v57 - int32(1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v67))))
	if v69 == int32(32) {
		v57 = v67
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v73 = v57
	goto L20
L25:
	;
	v79 = v74
	goto L27
L26:
	;
	v79 = int32(4)
	goto L27
L27:
	;
	v80 = v79 + v16
	if v76 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v113 = v107
	goto L39
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v86 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v97 = int32(1)
	if v78 != 0 {
		v107 = int32(base.Ui32(v76)>>(uint(v97)%32)) - v97
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v89 = int32(16)
	goto L34
L33:
	;
	v89 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v96 = int32(4)
	goto L37
L36:
	;
	v96 = v89
	goto L37
L37:
	;
	v107 = v96
	goto L28
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v113 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = F_varstr_cmp(m, v24, v73, v80, v128, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L46
	}
L41:
	;
	goto L40
L42:
	;
	v128 = v107 & (v107 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v123 = v113 - int32(1)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v123))))
	if v125 == int32(32) {
		v113 = v123
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v128 = v113
	goto L41
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v133 != v11 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_pfree(m, v11)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v137 != v16 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	F_pfree(m, v16)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return base.I64_extend_i32_u(base.B2i32(int32(0) <= v131))
L54:
	;
	goto L53
}
func F_bpcharrecv(m *base.Module, l0 int32) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v15 = F_pq_getmsgtext(m, v9, v10-v11, v6+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v21 = F_bpchar_input(m, v15, v19, v8, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			F_pfree(m, v15)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v21)
			}
		}
	}
}
func F_bpchartypmodin(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_anychar_typmodin(m, v3, int32(_a_F_bpchartypmodin_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_bqarr_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v130 int64
	_ = v130
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+28)) = int64(1)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+24)) = uint32(v14)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v15
	v25 = F_makepol_3(m, v12+int32(24))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return v130
L2:
	;
	return int64(0)
L3:
	;
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
	v130 = v9
	goto L1
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v31 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v34 = F_errsave_start(m, v15)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(int32(134217727)) <= base.Ui32(v31) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	if v34 == int32(0) {
		v130 = v9
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errmsg(m, int32(_a_F_bqarr_in_0), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, v15, int32(_a_F_bqarr_in_1), int32(544), int32(_a_F_bqarr_in_2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v130 = v9
	goto L1
L15:
	;
	v52 = F_errsave_start(m, v15)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v73 = v31<<(uint(int32(3))%32) + int32(8)
	v74 = F_palloc(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L23
	}
L18:
	;
	if v52 == int32(0) {
		v130 = v9
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(134217726)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v31
	F_errmsg(m, int32(_a_F_bqarr_in_3), v12)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_errsave_finish(m, v15, int32(_a_F_bqarr_in_1), int32(550), int32(_a_F_bqarr_in_2))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v130 = v9
	goto L1
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v73 << (uint(int32(2)) % 32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v85 = v82
	v86 = v31
	goto L24
L24:
	;
	v94 = v74 + v86<<(uint(int32(3))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*uint16)(unsafe.Add(mBase, uint32(v94))) = uint16(v95)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	F_pfree(m, v85)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L26
	}
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v106 - int32(1)
	v115 = F_findoprnd_2(m, v12+int32(24), v74+int32(8), v12+int32(20))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L28
	}
L26:
	;
	v102 = int32(1)
	if base.Ui32(v102) < base.Ui32(v86) {
		v85 = v99
		v86 = v86 - v102
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	if v115 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v119 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v119)
	v130 = v9
	goto L1
L30:
	;
	goto L31
L31:
	;
	v130 = base.I64_extend_i32_u(v74)
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
	var v211 int32
	_ = v211
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
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
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
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
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	if v336 != 0 {
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
	v235 = F_brinGetTupleForHeapBlock(m, v71, v61, v21+int32(28), v21+int32(26), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
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
	F_UnlockBuffer(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
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
	F_errfinish(m, int32(_a_F_brininsert_1), int32(421), int32(_a_F_brininsert_2))
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
	if v235 == int32(0) {
		v331 = v78
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
	v241 = int32(_a_F_brininsert_3)
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[0]))
	v248 = F_AllocSetContextCreateInternal(m, v243, int32(_a_F_brininsert_4), int32(0), int32(_a_F_brininsert_5), int32(_a_F_brininsert_6))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	v251 = v78
	goto L50
L50:
	;
	v253 = F_brin_deform_tuple(m, v70, v235, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brininsert[0])) = v248
	v251 = v248
	goto L50
L52:
	;
	v255 = F_add_values_to_range(m, l0, v70, v253, l1, l2)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	if v255 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_UnlockBuffer(m, v257)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L7
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v257 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v331 = v251
	goto L45
L58:
	;
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+26)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v279+v280<<(uint(int32(2))%32))+20))
	v286 = int32(base.Ui32(v284) >> (uint(int32(17)) % 32))
	v287 = int32(0)
	v289 = F_brin_copy_tuple(m, v235, v286, v287, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L62
	}
L59:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[5]))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265+(v257^int32(-1))<<(uint(int32(2))%32))))
	v279 = v271
	goto L58
L60:
	;
	goto L61
L61:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[6]))
	v279 = v273 + v257<<(uint(int32(13))%32) + int32(-8192)
	goto L58
L62:
	;
	v293 = F_brin_form_tuple(m, v70, v61, v253, v21+int32(20))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if base.Ui32(v286) < base.Ui32(v296) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_UnlockBuffer(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L72
	}
L65:
	;
	if v295 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v320 = int32(1)
	goto L67
L67:
	;
	goto L64
L68:
	;
	v316 = F_PageGetExactFreeSpace(m, v315)
	mBase = m.M
	v320 = base.B2i32(base.Ui32(v296-v286) <= base.Ui32(v316))
	goto L67
L69:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[5]))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v301+(v295^int32(-1))<<(uint(int32(2))%32))))
	v315 = v307
	goto L68
L70:
	;
	goto L71
L71:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[6]))
	v315 = v309 + v295<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+26)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v327 = F_brin_doupdate(m, l0, v59, v71, v61, v324, v325, v289, v286, v293, v326, v320)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	if v327 != 0 {
		v331 = v251
		goto L45
	} else {
		goto L74
	}
L74:
	;
	F_MemoryContextReset(m, v251)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v78 = v251
	goto L11
L76:
	;
	F_ReleaseBuffer(m, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
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
	if v331 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	F_MemoryContextDelete(m, v331)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
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
	var v251 int32
	_ = v251
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
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 float64
	_ = v362
	var v366 float64
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
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
	v450 = m.ExcPending
	if v450 != 0 {
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
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v414 != 0 {
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
		v397 = v53
		v408 = v64
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
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v393 = v392 + v58
	if base.Ui32(v393) < base.Ui32(v44) {
		v53 = v375
		v58 = v393
		v64 = v386
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
	v251 = int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+140))
	v257 = m.T0[v256].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l1, v249, v132, v250, v251, v250, v58, v247, v251, v129, v250)
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
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_ReleaseBuffer(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
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
	F_UnlockBuffer(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v129)+44))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	F_union_tuples(m, v349, v350, v344)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
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
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v129)+44))
	F_brin_memtuple_initialize(m, v356, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if l4 == int32(0) {
		v375 = v129
		v386 = v132
		goto L28
	} else {
		goto L85
	}
L85:
	;
	v362 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(v362, float64(1))
	v375 = v129
	v386 = v132
	goto L28
L86:
	;
	v366 = *(*float64)(unsafe.Add(mBase, uint32(l5)))
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = base.F64_add(v366, float64(1))
	goto L88
L87:
	;
	goto L88
L88:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	F_UnlockBuffer(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v375 = v53
	v386 = v64
	goto L28
L90:
	;
	v397 = v375
	v408 = v386
	goto L19
L91:
	;
	F_ReleaseBuffer(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
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
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	if v397 == int32(0) {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	F_terminate_brin_buildstate(m, v397)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_pfree(m, v408)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
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
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_brinsummarize_1), int32(1866), int32(_a_F_brinsummarize_2))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
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
func F_btbpchar_pattern_cmp(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v17 = int32(1)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v21 = v19 & v17
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v129 != v6 {
		goto L53
	} else {
		goto L54
	}
L5:
	;
	v22 = v17
	goto L7
L6:
	;
	v22 = int32(4)
	goto L7
L7:
	;
	v23 = v22 + v6
	if v19 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v56 = v50
	goto L19
L9:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	if v29 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v40 = int32(1)
	if v21 != 0 {
		v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v32 = int32(16)
	goto L14
L13:
	;
	v32 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v39 = int32(4)
	goto L17
L16:
	;
	v39 = v32
	goto L17
L17:
	;
	v50 = v39
	goto L8
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	if v56 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v70 = int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v74 = v72 & v70
	if v74 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v68 = v50 & (v50 >> (uint(int32(31)) % 32))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v63 = v56 - int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v63))))
	if v65 == int32(32) {
		v56 = v63
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v68 = v56
	goto L21
L26:
	;
	v75 = v70
	goto L28
L27:
	;
	v75 = int32(4)
	goto L28
L28:
	;
	v76 = v75 + v11
	if v72 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v109 = v103
	goto L40
L30:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v82 == int32(18) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v93 = int32(1)
	if v74 != 0 {
		v103 = int32(base.Ui32(v72)>>(uint(v93)%32)) - v93
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v85 = int32(16)
	goto L35
L34:
	;
	v85 = int32(0)
	goto L35
L35:
	;
	if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v92 = int32(4)
	goto L38
L37:
	;
	v92 = v85
	goto L38
L38:
	;
	v103 = v92
	goto L29
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v103 = int32(base.Ui32(v97)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	if v109 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v123 = base.B2i32(v68 < v121)
	if v68 < v121 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v121 = v103 & (v103 >> (uint(int32(31)) % 32))
	goto L42
L44:
	;
	goto L45
L45:
	;
	v116 = v109 - int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v116))))
	if v118 == int32(32) {
		v109 = v116
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v121 = v109
	goto L42
L47:
	;
	goto L4
L48:
	;
	v124 = v68
	goto L50
L49:
	;
	v124 = v121
	goto L50
L50:
	;
	v125 = F_memcmp(m, v23, v76, v124)
	mBase = m.M
	if v125 != 0 {
		v128 = v125
		goto L47
	} else {
		goto L51
	}
L51:
	;
	if v68 < v121 {
		v128 = int32(-1)
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v128 = base.B2i32(v121 < v68)
	goto L47
L53:
	;
	F_pfree(m, v6)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v133 != v11 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v11)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return base.I64_extend_i32_s(v128)
L60:
	;
	goto L59
}
func F_btbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
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
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
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
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v275 int64
	_ = v275
	var v283 int32
	_ = v283
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
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
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
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
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
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
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
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
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
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 float64
	_ = v498
	var v499 int32
	_ = v499
	var v500 float64
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 float64
	_ = v544
	var v546 int32
	_ = v546
	var v548 float64
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v573 float64
	_ = v573
	var v574 float64
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v735 int64
	_ = v735
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v986 int32
	_ = v986
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1075 int64
	_ = v1075
	var v1083 int32
	_ = v1083
	var v1094 int32
	_ = v1094
	var v1112 int32
	_ = v1112
	var v1115 int64
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int64
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1154 int32
	_ = v1154
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1220 int32
	_ = v1220
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int64
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int64
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1342 int32
	_ = v1342
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int64
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1385 int64
	_ = v1385
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1405 int64
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1486 int32
	_ = v1486
	var v1493 int32
	_ = v1493
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1504 int32
	_ = v1504
	var v1508 int32
	_ = v1508
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1580 int32
	_ = v1580
	var v1586 int32
	_ = v1586
	var v1596 int32
	_ = v1596
	var v1602 int32
	_ = v1602
	var v1606 int64
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1682 int64
	_ = v1682
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1702 int64
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	var v1744 int64
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1784 int32
	_ = v1784
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1805 int32
	_ = v1805
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1818 int32
	_ = v1818
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1898 int32
	_ = v1898
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1916 int32
	_ = v1916
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2008 float64
	_ = v2008
	v4 = int32(0)
	v17 = int64(0)
	v22 = m.G0
	v24 = v22 - int32(96)
	m.G0 = v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+16)) = uint8(v26)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+117)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+18)) = uint8(v4)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+17)) = uint8(v28)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v4
	v40 = F_RelationGetNumberOfBlocksInFork(m, l1, v4)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v1959 = F_smgr_bulk_get_buf(m, v1958)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L4
	} else {
		goto L356
	}
L2:
	;
	v1836 = int32(0)
	v1839 = v1836
	v1840 = v1836
	v1841 = v1818
	goto L337
L3:
	;
	F_pfree(m, v1349)
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L4
	} else {
		goto L336
	}
L4:
	;
	return int32(0)
L5:
	;
	if v40 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v47 = F_palloc0(m, int32(16))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	v1798 = m.ExcPending
	if v1798 != 0 {
		goto L4
	} else {
		goto L333
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = l0
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+12)) = uint8(v51)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+117)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+13)) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v47
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v60 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v101 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v64&int32(1) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v69 = int32(_a_F_btbuild_0)
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v72 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v71 + v72
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v75 + v72
	v79 = int32(0)
	v81 = int32(_a_F_btbuild_1)
	v82 = base.AtomicRmwOr32(m, v79, v81, v79)
	*(*int64)(unsafe.Add(mBase, uint32(v60+int32(80))+232)) = int64(2)
	v90 = base.AtomicRmwOr32(m, v79, v81, v79)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v91 + v72
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v97 - v72
	goto L11
L14:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v427 != 0 {
		goto L98
	} else {
		goto L99
	}
L15:
	;
	v105 = v101 + int32(1)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	v108 = F_palloc0(m, int32(32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[3]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+72)) = v113 + int32(1)
	goto L17
L17:
	;
	v120 = F_CreateParallelContext(m, int32(_a_F_btbuild_2), int32(_a_F_btbuild_3), v101)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v106 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v124 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v128 = int32(_a_F_btbuild_4)
	goto L21
L21:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v131 = F_table_parallelscan_estimate(m, v130, v128)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v126 = F_RegisterSnapshot(m, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v128 = v126
	goto L21
L24:
	;
	v133 = F_add_size(m, int32(96), v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v120)+36))
	v140 = F_add_size(m, v135, (v133+int32(31))&int32(-32))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+36)) = v140
	v143 = F_tuplesort_estimate_shared(m, v105)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v120)+36))
	v149 = (v143 + int32(31)) & int32(-32)
	v150 = F_add_size(m, v145, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+36)) = v150
	v153 = int32(1)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+12)))
	if v155 == v153 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v158 = F_add_size(m, v150, v149)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v162 = int32(2)
	goto L31
L31:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v120)+40))
	v164 = F_add_size(m, v163, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+36)) = v158
	v162 = int32(3)
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+40)) = v164
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v120)+36))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v170 = F_mul_size(m, int32(40), v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v176 = F_add_size(m, v167, (v170+int32(31))&int32(-32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+36)) = v176
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+40))
	v181 = F_add_size(m, v179, int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+40)) = v181
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v120)+36))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v187 = F_mul_size(m, int32(128), v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v193 = F_add_size(m, v184, (v187+int32(31))&int32(-32))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+36)) = v193
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v120)+40))
	v198 = F_add_size(m, v196, int32(1))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+40)) = v198
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	if v202 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v120)+36))
	v204 = F_strlen(m, v202)
	mBase = m.M
	v209 = F_add_size(m, v203, v204&int32(-32)+int32(32))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	v219 = v153
	goto L42
L42:
	;
	F_InitializeParallelDSM(m, v120)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+36)) = v209
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v120)+40))
	v214 = F_add_size(m, v212, int32(1))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+40)) = v214
	v219 = v204 + int32(1)
	goto L42
L45:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v120)+44))
	if v222 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v225 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	v240 = F_shm_toc_allocate(m, v239, v133)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L55
	}
L49:
	;
	F_UnregisterSnapshot(m, v128)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_DestroyParallelContext(m, v120)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[3]))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+72)) = v235 - int32(1)
	goto L54
L54:
	;
	goto L14
L55:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v246
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+8)) = uint8(v248)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+12)) = v105
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+10)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+9)) = uint8(v250)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v256 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v240)+16)) = v261
	v264 = v240 + int32(24)
	v265 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v264))), uint32(v265))
	*(*int64)(unsafe.Add(mBase, uint32(v264)+4)) = int64(-1)
	goto L60
L57:
	;
	v261 = int64(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v256)+392))
	v261 = v260
	goto L56
L60:
	;
	v270 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v240)+36)), uint32(v270))
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+72)) = uint8(v270)
	v275 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v240)+64)) = v275
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+56)) = uint8(v270)
	*(*int64)(unsafe.Add(mBase, uint32(v240)+48)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v240)+40)) = v270
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	F_table_parallelscan_initialize(m, v283, v240+int32(96), v128)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	v289 = F_shm_toc_allocate(m, v288, v143)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v120)+44))
	F_tuplesort_initialize_shared(m, v289, v105, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	F_shm_toc_insert(m, v294, int64(-6917529027641081855), v240)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	F_shm_toc_insert(m, v298, int64(-6917529027641081854), v289)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+12)))
	if v303 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	v307 = F_shm_toc_allocate(m, v306, v143)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	v316 = int32(0)
	goto L68
L68:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	if v318 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v120)+44))
	F_tuplesort_initialize_shared(m, v307, v105, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	F_shm_toc_insert(m, v312, int64(-6917529027641081853), v307)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v316 = v307
	goto L68
L72:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	v320 = F_shm_toc_allocate(m, v319, v219)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v333 = F_mul_size(m, int32(40), v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L80
	}
L75:
	;
	if v219 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	base.MemoryCopy(m, v320, v323, v219)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	F_shm_toc_insert(m, v325, int64(-6917529027641081852), v320)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	goto L74
L80:
	;
	v335 = F_shm_toc_allocate(m, v330, v333)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	F_shm_toc_insert(m, v337, int64(-6917529027641081851), v335)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v344 = F_mul_size(m, int32(128), v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v346 = F_shm_toc_allocate(m, v341, v344)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
	F_shm_toc_insert(m, v348, int64(-6917529027641081850), v346)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_LaunchParallelWorkers(m, v120)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v120
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+28)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v108)+12)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v355 + int32(1)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	if v365 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F__bt_end_parallel(m, v108)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v108
	v372 = F_palloc0(m, int32(16))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L91
	}
L90:
	;
	goto L14
L91:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v372)+4)) = v375
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v372)+8)) = v378
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v372)+12)) = uint8(v381)
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v372)+13)) = uint8(v383)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+8)))
	if v387 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v391 = F_palloc0(m, int32(16))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L95
	}
L93:
	;
	v400 = int32(0)
	v402 = v386
	goto L94
L94:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[5]))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v408 = base.I32_div_s(v406, v407)
	F__bt_parallel_scan_and_sort(m, v372, v400, v402, v403, v404, v408, int32(1))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L96
	}
L95:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v391)+4)) = v393
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v396 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v391)+12)) = uint8(v396)
	*(*int32)(unsafe.Add(mBase, uint32(v391)+8)) = v395
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v400 = v391
	v402 = v399
	goto L94
L96:
	;
	F_WaitForParallelWorkersToAttach(m, v120)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	goto L14
L98:
	;
	v429 = F_palloc0(m, int32(12))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	v439 = int32(0)
	goto L100
L100:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+16)))
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+17)))
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[5]))
	v444 = F_tuplesort_begin_index_btree(m, l0, l1, v440, v441, v443, v439)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	v431 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v429))) = uint8(v431)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v429)+4)) = v434
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v429)+8)) = v437
	v439 = v429
	goto L100
L102:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v446))) = v444
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	if v448 == int32(1) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v452 = F_palloc0(m, int32(16))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v484 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L106:
	;
	v454 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v452)+12)) = uint8(v454)
	*(*int32)(unsafe.Add(mBase, uint32(v452)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v452)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v452
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v460 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v462 = F_palloc0(m, int32(12))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	v473 = v454
	v474 = v452
	goto L109
L109:
	;
	v475 = int32(0)
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[6]))
	v479 = F_tuplesort_begin_index_btree(m, l0, l1, v475, v475, v478, v473)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L111
	}
L110:
	;
	v464 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v462))) = uint8(v464)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+4)) = v467
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+8)) = v470
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v473 = v462
	v474 = v472
	goto L109
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = v479
	goto L105
L112:
	;
	v575 = int32(0)
	v577 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+88)) = v577
	v580 = *(*int64)(unsafe.Add(mBase, _c_F_btbuild[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+80)) = v580
	v582 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+56)) = v582
	*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = base.I64_trunc_sat_f64_s(v574)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+64)) = v582
	goto L130
L113:
	;
	v487 = int32(1)
	v488 = int32(0)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+140))
	v498 = m.T0[v497].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v487, v488, v487, v488, int32(-1), int32(243), v24+int32(16), v488)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v484)+8))
	v505 = v501 + int32(36)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	goto L117
L116:
	;
	v500 = *(*float64)(unsafe.Add(mBase, uint32(v24)+32))
	v573 = v498
	v574 = v500
	goto L112
L117:
	;
	v530 = base.AtomicRmwXchg32(m, v505, int32(0), int32(1))
	if v530 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+18)) = uint8(v542)
	v544 = *(*float64)(unsafe.Add(mBase, uint32(v501)+64))
	*(*float64)(unsafe.Add(mBase, uint32(v24)+32)) = v544
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+72)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)) = uint8(v546)
	v548 = *(*float64)(unsafe.Add(mBase, uint32(v501)+48))
	v549 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v501)+36)), uint32(v549))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L4
	} else {
		goto L127
	}
L119:
	;
	F_s_lock(m, v505, int32(_a_F_btbuild_5))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v501)+40))
	if v506 != v534 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L121
L123:
	;
	v536 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v505))), uint32(v536))
	F_ConditionVariableSleep(m, v501+int32(24), int32(134217767))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L4
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L118
L126:
	;
	goto L117
L127:
	;
	v573 = v548
	v574 = v544
	goto L112
L128:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v774 == int32(0) {
		v785 = v575
		goto L145
	} else {
		goto L146
	}
L129:
	;
	goto L128
L130:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v602 == int32(0) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v606&int32(1) == int32(0) {
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v611 = int32(_a_F_btbuild_0)
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v614 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v613 + v614
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	*(*int32)(unsafe.Add(mBase, uint32(v602))) = v617 + v614
	v621 = int32(0)
	v624 = base.AtomicRmwOr32(m, v621, int32(_a_F_btbuild_1), v621)
	goto L134
L133:
	;
	v751 = int32(0)
	v754 = base.AtomicRmwOr32(m, v751, int32(_a_F_btbuild_1), v751)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	v756 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v602))) = v755 + v756
	v759 = int32(_a_F_btbuild_0)
	v761 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v761 - v756
	goto L129
L134:
	;
	goto L136
L136:
	;
	goto L137
L137:
	;
	v716 = int32(0)
	v719 = v575
	goto L142
L142:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(80)+v719<<(uint(int32(2))%32))))
	v729 = int32(3)
	v735 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(48)+v719<<(uint(v729)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v602+int32(232)+v728<<(uint(v729)%32)))) = v735
	v737 = int32(1)
	v740 = v716 + v737
	if v740 != int32(3) {
		v716 = v740
		v719 = v719 + v737
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L133
L144:
	;
	goto L143
L145:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v791 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L146:
	;
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+18)))
	if v777 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v785 = v774
	goto L145
L148:
	;
	goto L149
L149:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	F_tuplesort_end(m, v778)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	F_pfree(m, v774)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = int32(0)
	v785 = v575
	goto L145
L152:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	F_tuplesort_performsort(m, v832)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L4
	} else {
		goto L156
	}
L153:
	;
	goto L152
L154:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v795&int32(1) == int32(0) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v800 = int32(_a_F_btbuild_0)
	v802 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v803 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v802 + v803
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v791)))
	*(*int32)(unsafe.Add(mBase, uint32(v791))) = v806 + v803
	v810 = int32(0)
	v812 = int32(_a_F_btbuild_1)
	v813 = base.AtomicRmwOr32(m, v810, v812, v810)
	*(*int64)(unsafe.Add(mBase, uint32(v791+int32(80))+232)) = int64(3)
	v821 = base.AtomicRmwOr32(m, v810, v812, v810)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v791)))
	*(*int32)(unsafe.Add(mBase, uint32(v791))) = v822 + v803
	v828 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v828 - v803
	goto L153
L156:
	;
	if v785 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v839 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L159
L159:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v786)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v883
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v786)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v885
	v887 = int32(0)
	v889 = F__bt_mkscankey(m, v885, v887)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L4
	} else {
		goto L165
	}
L160:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	F_tuplesort_performsort(m, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L4
	} else {
		goto L164
	}
L161:
	;
	goto L160
L162:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v843&int32(1) == int32(0) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v848 = int32(_a_F_btbuild_0)
	v850 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v851 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v850 + v851
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v839)))
	*(*int32)(unsafe.Add(mBase, uint32(v839))) = v854 + v851
	v858 = int32(0)
	v860 = int32(_a_F_btbuild_1)
	v861 = base.AtomicRmwOr32(m, v858, v860, v858)
	*(*int64)(unsafe.Add(mBase, uint32(v839+int32(80))+232)) = int64(4)
	v869 = base.AtomicRmwOr32(m, v858, v860, v858)
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v839)))
	*(*int32)(unsafe.Add(mBase, uint32(v839))) = v870 + v851
	v876 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v876 - v851
	goto L161
L164:
	;
	goto L159
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v889
	v893 = F__bt_allequalimage(m, v885, int32(1))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v889)+1)) = uint8(v893)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = int32(1)
	v902 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v902 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v885)+52))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v885)+192))
	v945 = int32(*(*int16)(unsafe.Add(mBase, uint32(v944)+10)))
	v947 = F_smgr_bulk_start_rel(m, v885, int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L4
	} else {
		goto L171
	}
L168:
	;
	goto L167
L169:
	;
	v906 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v906&int32(1) == int32(0) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v911 = int32(_a_F_btbuild_0)
	v913 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v914 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v913 + v914
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	*(*int32)(unsafe.Add(mBase, uint32(v902))) = v917 + v914
	v921 = int32(0)
	v923 = int32(_a_F_btbuild_1)
	v924 = base.AtomicRmwOr32(m, v921, v923, v921)
	*(*int64)(unsafe.Add(mBase, uint32(v902+int32(80))+232)) = int64(5)
	v932 = base.AtomicRmwOr32(m, v921, v923, v921)
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	*(*int32)(unsafe.Add(mBase, uint32(v902))) = v933 + v914
	v939 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v939 - v914
	goto L168
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v947
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v889)+1)))
	if v950 != int32(1) {
		v959 = v887
		goto L177
	} else {
		goto L178
	}
L172:
	;
	F_pfree(m, v976)
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L4
	} else {
		goto L331
	}
L173:
	;
	v1669 = int32(0)
	v1672 = v962
	v1682 = v17
	goto L313
L174:
	;
	v1349 = F_palloc(m, int32(1676))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L4
	} else {
		goto L255
	}
L175:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v970 = F_tuplesort_getheaptuple(m, v969)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L4
	} else {
		goto L186
	}
L176:
	;
	if v785 == int32(0) {
		goto L174
	} else {
		goto L185
	}
L177:
	;
	if v785 != 0 {
		goto L175
	} else {
		goto L181
	}
L178:
	;
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+12)))
	if v953 != 0 {
		v959 = v887
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v885)+180))
	if v954 == int32(0) {
		goto L176
	} else {
		goto L180
	}
L180:
	;
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v954)+16)))
	v959 = v957
	goto L177
L181:
	;
	if v959 != 0 {
		goto L174
	} else {
		goto L182
	}
L182:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v962 = F_tuplesort_getheaptuple(m, v961)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	if v962 != 0 {
		goto L173
	} else {
		goto L184
	}
L184:
	;
	v1938 = int32(0)
	v1939 = int32(0)
	goto L1
L185:
	;
	goto L175
L186:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	v973 = F_tuplesort_getheaptuple(m, v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	v976 = F_palloc0_mul(m, int32(36), v945)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	if int32(0) < v945 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v986 = int32(0)
	goto L192
L190:
	;
	goto L191
L191:
	;
	v1058 = v970
	v1060 = int32(0)
	v1064 = v973
	v1075 = v17
	goto L196
L192:
	;
	v1006 = v976 + v986*int32(36)
	v1008 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1006))) = v1008
	v1012 = v889 + int32(16) + v986*int32(56)
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1006)+4)) = v1013
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1012)))
	v1019 = int32(base.Ui32(v1015)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1006)+9)) = uint8(v1019)
	v1021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012)+4)))
	v1022 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1006)+20)) = uint8(v1022)
	*(*uint16)(unsafe.Add(mBase, uint32(v1006)+10)) = uint16(v1021)
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1012)))
	F_PrepareSortSupportFromIndexRel(m, v885, int32(base.Ui32(v1025&int32(16777216))>>(uint(int32(24))%32)), v1006)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L4
	} else {
		goto L194
	}
L193:
	;
	goto L191
L194:
	;
	v1033 = v986 + int32(1)
	if v1033 != v945 {
		v986 = v1033
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	if v1064 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v1060 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L199:
	;
	if v1058 == int32(0) {
		goto L172
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1083 = int32(0)
	if v1058 == v1083 {
		v1220 = v1083
		goto L198
	} else {
		goto L203
	}
L202:
	;
	v1220 = int32(1)
	goto L198
L203:
	;
	if int32(0) < v945 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v1094 = int32(1)
	goto L207
L205:
	;
	goto L206
L206:
	;
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058)+2)))
	v1189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058))))
	v1190 = int32(16)
	v1192 = v1188 | v1189<<(uint(v1190)%32)
	v1193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1064)+2)))
	v1194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1064))))
	v1197 = v1193 | v1194<<(uint(v1190)%32)
	if base.Ui32(v1192) < base.Ui32(v1197) {
		v1208 = int32(-1)
		goto L230
	} else {
		goto L231
	}
L207:
	;
	v1112 = v976 + v1094*int32(36)
	v1115 = F_index_getattr_2(m, v1058, v1094, v943, v24+int32(80))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L4
	} else {
		goto L209
	}
L208:
	;
	goto L206
L209:
	;
	v1119 = F_index_getattr_2(m, v1064, v1094, v943, v24+int32(95))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+95)))
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+80)))
	if v1122 == int32(1) {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	if v945 != v1094 {
		v1094 = v1094 + int32(1)
		goto L207
	} else {
		goto L228
	}
L212:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1112-int32(20))))
	v1143 = m.T0[v1142].(func(*base.Module, int64, int64, int32) int32)(m, v1115, v1119, v1112-int32(36))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L4
	} else {
		goto L221
	}
L213:
	;
	v1220 = int32(1)
	goto L198
L214:
	;
	if v1121&int32(1) != 0 {
		goto L211
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	if v1121&int32(1) == int32(0) {
		goto L212
	} else {
		goto L219
	}
L217:
	;
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112-int32(27)))))
	if v1129 != 0 {
		goto L213
	} else {
		goto L218
	}
L218:
	;
	v1220 = v1083
	goto L198
L219:
	;
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112-int32(27)))))
	if v1136 != 0 {
		v1220 = v1083
		goto L198
	} else {
		goto L220
	}
L220:
	;
	goto L213
L221:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112-int32(28)))))
	if v1147 == int32(1) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	if v1143 < int32(0) {
		v1220 = v1083
		goto L198
	} else {
		goto L225
	}
L223:
	;
	v1154 = v1143
	goto L224
L224:
	;
	if int32(0) < v1154 {
		v1220 = v1083
		goto L198
	} else {
		goto L226
	}
L225:
	;
	v1154 = int32(0) - v1143
	goto L224
L226:
	;
	if v1154 == int32(0) {
		goto L211
	} else {
		goto L227
	}
L227:
	;
	v1220 = int32(1)
	goto L198
L228:
	;
	goto L208
L229:
	;
	v1220 = base.B2i32(v1208 <= int32(0))
	goto L198
L230:
	;
	goto L229
L231:
	;
	if base.Ui32(v1197) < base.Ui32(v1192) {
		v1208 = int32(1)
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v1202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058)+4)))
	v1203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1064)+4)))
	if base.Ui32(v1202) < base.Ui32(v1203) {
		v1208 = int32(-1)
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v1208 = base.B2i32(base.Ui32(v1203) < base.Ui32(v1202))
	goto L230
L234:
	;
	v1235 = F_palloc0(m, int32(32))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L4
	} else {
		goto L237
	}
L235:
	;
	v1280 = v1060
	goto L236
L236:
	;
	if v1220 != 0 {
		goto L244
	} else {
		goto L245
	}
L237:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v1238 = F_smgr_bulk_get_buf(m, v1237)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L4
	} else {
		goto L238
	}
L238:
	;
	F_PageInit(m, v1238, int32(_a_F_btbuild_6), int32(16))
	mBase = m.M
	goto L239
L239:
	;
	v1243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1238)+16)))
	v1244 = v1238 + v1243
	*(*int64)(unsafe.Add(mBase, uint32(v1244)+8)) = int64(4294967296)
	v1247 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1244))) = v1247
	v1249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1238)+12)))
	v1251 = v1249 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1238)+12)) = uint16(v1251)
	*(*int32)(unsafe.Add(mBase, uint32(v1235))) = v1238
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v1255 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v1254 + v1255
	*(*int64)(unsafe.Add(mBase, uint32(v1235)+16)) = v1247
	*(*uint16)(unsafe.Add(mBase, uint32(v1235)+12)) = uint16(v1255)
	*(*int32)(unsafe.Add(mBase, uint32(v1235)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1235)+4)) = v1254
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+180))
	if v1266 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+4))
	v1273 = base.I32_div_s(int32(_a_F_btbuild_7)-v1268<<(uint(int32(13))%32), int32(100))
	v1275 = v1273
	goto L242
L241:
	;
	v1275 = int32(819)
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1235)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1235)+24)) = v1275
	v1280 = v1235
	goto L236
L243:
	;
	v1302 = v1075 + int64(1)
	v1305 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1305 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L244:
	;
	F__bt_buildadd(m, v24+int32(48), v1280, v1058, int32(0))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L4
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	F__bt_buildadd(m, v24+int32(48), v1280, v1064, int32(0))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L4
	} else {
		goto L249
	}
L247:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v1288 = F_tuplesort_getheaptuple(m, v1287)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L4
	} else {
		goto L248
	}
L248:
	;
	v1298 = v1288
	v1299 = v1064
	goto L243
L249:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	v1296 = F_tuplesort_getheaptuple(m, v1295)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L4
	} else {
		goto L250
	}
L250:
	;
	v1298 = v1058
	v1299 = v1296
	goto L243
L251:
	;
	v1058 = v1298
	v1060 = v1280
	v1064 = v1299
	v1075 = v1302
	goto L196
L252:
	;
	goto L251
L253:
	;
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1309&int32(1) == int32(0) {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v1314 = int32(_a_F_btbuild_0)
	v1316 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1317 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1316 + v1317
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	*(*int32)(unsafe.Add(mBase, uint32(v1305))) = v1320 + v1317
	v1324 = int32(0)
	v1326 = int32(_a_F_btbuild_1)
	v1327 = base.AtomicRmwOr32(m, v1324, v1326, v1324)
	*(*int64)(unsafe.Add(mBase, uint32(v1305+int32(96))+232)) = v1302
	v1335 = base.AtomicRmwOr32(m, v1324, v1326, v1324)
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	*(*int32)(unsafe.Add(mBase, uint32(v1305))) = v1336 + v1317
	v1342 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1342 - v1317
	goto L252
L255:
	;
	v1351 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1349)+4)) = v1351
	v1353 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1349))) = uint8(v1353)
	*(*int64)(unsafe.Add(mBase, uint32(v1349)+10)) = v1351
	*(*int64)(unsafe.Add(mBase, uint32(v1349)+20)) = v1351
	*(*int64)(unsafe.Add(mBase, uint32(v1349)+28)) = v1351
	*(*int64)(unsafe.Add(mBase, uint32(v1349)+36)) = v1351
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v1364 = F_tuplesort_getheaptuple(m, v1363)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L4
	} else {
		goto L256
	}
L256:
	;
	if v1364 == int32(0) {
		goto L3
	} else {
		goto L257
	}
L257:
	;
	v1372 = int32(0)
	v1375 = v1364
	v1385 = v17
	goto L258
L258:
	;
	if v1372 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L259:
	;
	F__bt_sort_dedup_finish_pending(m, v24+int32(48), v1602, v1349)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L4
	} else {
		goto L309
	}
L260:
	;
	v1606 = v1385 + int64(1)
	v1609 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1609 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L261:
	;
	v1538 = F_CopyIndexTuple(m, v1375)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L4
	} else {
		goto L292
	}
L262:
	;
	v1393 = F_palloc0(m, int32(32))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L4
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+12))
	v1445 = F__bt_keep_natts_fast(m, v1443, v1444, v1375)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L4
	} else {
		goto L272
	}
L265:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v1396 = F_smgr_bulk_get_buf(m, v1395)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	F_PageInit(m, v1396, int32(_a_F_btbuild_6), int32(16))
	mBase = m.M
	goto L267
L267:
	;
	v1401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1396)+16)))
	v1402 = v1396 + v1401
	*(*int64)(unsafe.Add(mBase, uint32(v1402)+8)) = int64(4294967296)
	v1405 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1402))) = v1405
	v1407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1396)+12)))
	v1409 = v1407 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1396)+12)) = uint16(v1409)
	*(*int32)(unsafe.Add(mBase, uint32(v1393))) = v1396
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v1413 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v1412 + v1413
	*(*int64)(unsafe.Add(mBase, uint32(v1393)+16)) = v1405
	*(*uint16)(unsafe.Add(mBase, uint32(v1393)+12)) = uint16(v1413)
	*(*int32)(unsafe.Add(mBase, uint32(v1393)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1393)+4)) = v1412
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1423)+180))
	if v1424 != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+4))
	v1431 = base.I32_div_s(int32(_a_F_btbuild_7)-v1426<<(uint(int32(13))%32), int32(100))
	v1433 = v1431
	goto L270
L269:
	;
	v1433 = int32(819)
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1393)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1393)+24)) = v1433
	v1437 = int32(812)
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+8)) = v1437
	v1440 = F_palloc(m, v1437)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+24)) = v1440
	v1536 = v1393
	goto L261
L272:
	;
	if v945 < v1445 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1453 = int32(1)
	v1454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1375)+7)))
	if v1454&int32(32) == int32(0) {
		v1472 = v1453
		v1474 = v1375
		goto L277
	} else {
		goto L278
	}
L274:
	;
	goto L275
L275:
	;
	F__bt_sort_dedup_finish_pending(m, v24+int32(48), v1372, v1349)
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L4
	} else {
		goto L290
	}
L276:
	;
	if base.Ui32((v1476+(v1477+v1472)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1475) {
		v1602 = v1372
		goto L260
	} else {
		goto L289
	}
L277:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+8))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+20))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+28))
	v1486 = base.B2i32(base.Ui32((v1476+(v1477+v1472)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1475))
	if v1486 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L278:
	;
	v1459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1375)+4)))
	if v1459&int32(_a_F_btbuild_6) == int32(0) {
		v1472 = v1453
		v1474 = v1375
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v1466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1375)+2)))
	v1467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1375))))
	v1472 = v1459 & int32(4095)
	v1474 = v1466 + (v1375 + v1467<<(uint(int32(16))%32))
	goto L277
L280:
	;
	goto L276
L281:
	;
	v1520 = v1349 + v1517
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1520)))
	*(*int32)(unsafe.Add(mBase, uint32(v1520))) = v1521 + v1519
	goto L280
L282:
	;
	if v1477 <= int32(50) {
		goto L280
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+32)) = v1493 + int32(1)
	v1498 = v1472 * int32(6)
	if v1498 != 0 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	v1517 = int32(4)
	v1519 = int32(1)
	goto L281
L286:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+24))
	base.MemoryCopy(m, v1499+v1477*int32(6), v1474, v1498)
	goto L288
L287:
	;
	goto L288
L288:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+28)) = v1504 + v1472
	v1508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1375)+6)))
	v1517 = int32(36)
	v1519 = (v1508&int32(_a_F_btbuild_8)+int32(7))&int32(_a_F_btbuild_9) | int32(4)
	goto L281
L289:
	;
	goto L275
L290:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+12))
	F_pfree(m, v1532)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L4
	} else {
		goto L291
	}
L291:
	;
	v1536 = v1372
	goto L261
L292:
	;
	v1540 = int32(0)
	v1543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1538)+7)))
	if v1543&int32(32) != 0 {
		goto L296
	} else {
		goto L297
	}
L293:
	;
	v1602 = v1536
	goto L260
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+20)) = v1580
	*(*uint16)(unsafe.Add(mBase, uint32(v1349)+16)) = uint16(v1540)
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+12)) = v1538
	v1586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+36)) = (v1586&int32(_a_F_btbuild_8)+int32(7))&int32(_a_F_btbuild_9) | int32(4)
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1349+v1596<<(uint(int32(2))%32))+44)) = uint16(v1540)
	goto L293
L295:
	;
	v1561 = v1546 & int32(4095)
	v1563 = v1561 * int32(6)
	if v1563 != 0 {
		goto L300
	} else {
		goto L301
	}
L296:
	;
	v1546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538)+4)))
	if v1546&int32(_a_F_btbuild_6) != 0 {
		goto L295
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+24))
	v1551 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1550)+4)) = uint16(v1551)
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1538)))
	*(*int32)(unsafe.Add(mBase, uint32(v1550))) = v1553
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+28)) = int32(1)
	v1557 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538)+6)))
	v1580 = v1557 & int32(_a_F_btbuild_8)
	goto L294
L299:
	;
	goto L298
L300:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+24))
	v1565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538)+2)))
	v1566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538))))
	base.MemoryCopy(m, v1564, v1565+(v1538+v1566<<(uint(int32(16))%32)), v1563)
	goto L302
L301:
	;
	goto L302
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+28)) = v1561
	v1573 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538)+2)))
	v1574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538))))
	v1580 = v1573 | v1574<<(uint(int32(16))%32)
	goto L294
L303:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v1651 = F_tuplesort_getheaptuple(m, v1650)
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L4
	} else {
		goto L307
	}
L304:
	;
	goto L303
L305:
	;
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1613&int32(1) == int32(0) {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v1618 = int32(_a_F_btbuild_0)
	v1620 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1621 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1620 + v1621
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1609)))
	*(*int32)(unsafe.Add(mBase, uint32(v1609))) = v1624 + v1621
	v1628 = int32(0)
	v1630 = int32(_a_F_btbuild_1)
	v1631 = base.AtomicRmwOr32(m, v1628, v1630, v1628)
	*(*int64)(unsafe.Add(mBase, uint32(v1609+int32(96))+232)) = v1606
	v1639 = base.AtomicRmwOr32(m, v1628, v1630, v1628)
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1609)))
	*(*int32)(unsafe.Add(mBase, uint32(v1609))) = v1640 + v1621
	v1646 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1646 - v1621
	goto L304
L307:
	;
	if v1651 != 0 {
		v1372 = v1602
		v1375 = v1651
		v1385 = v1606
		goto L258
	} else {
		goto L308
	}
L308:
	;
	goto L259
L309:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+12))
	F_pfree(m, v1657)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L4
	} else {
		goto L310
	}
L310:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+24))
	F_pfree(m, v1660)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L4
	} else {
		goto L311
	}
L311:
	;
	F_pfree(m, v1349)
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L4
	} else {
		goto L312
	}
L312:
	;
	v1818 = v1602
	goto L2
L313:
	;
	if v1669 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L314:
	;
	v1818 = v1736
	goto L2
L315:
	;
	v1690 = F_palloc0(m, int32(32))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L4
	} else {
		goto L318
	}
L316:
	;
	v1736 = v1669
	goto L317
L317:
	;
	F__bt_buildadd(m, v24+int32(48), v1736, v1672, int32(0))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L4
	} else {
		goto L324
	}
L318:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v1693 = F_smgr_bulk_get_buf(m, v1692)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L4
	} else {
		goto L319
	}
L319:
	;
	F_PageInit(m, v1693, int32(_a_F_btbuild_6), int32(16))
	mBase = m.M
	goto L320
L320:
	;
	v1698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+16)))
	v1699 = v1693 + v1698
	*(*int64)(unsafe.Add(mBase, uint32(v1699)+8)) = int64(4294967296)
	v1702 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1699))) = v1702
	v1704 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+12)))
	v1706 = v1704 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1693)+12)) = uint16(v1706)
	*(*int32)(unsafe.Add(mBase, uint32(v1690))) = v1693
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v1710 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v1709 + v1710
	*(*int64)(unsafe.Add(mBase, uint32(v1690)+16)) = v1702
	*(*uint16)(unsafe.Add(mBase, uint32(v1690)+12)) = uint16(v1710)
	*(*int32)(unsafe.Add(mBase, uint32(v1690)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1690)+4)) = v1709
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1720)+180))
	if v1721 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1721)+4))
	v1728 = base.I32_div_s(int32(_a_F_btbuild_7)-v1723<<(uint(int32(13))%32), int32(100))
	v1730 = v1728
	goto L323
L322:
	;
	v1730 = int32(819)
	goto L323
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1690)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1690)+24)) = v1730
	v1736 = v1690
	goto L317
L324:
	;
	v1744 = v1682 + int64(1)
	v1747 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1747 == int32(0) {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v1789 = F_tuplesort_getheaptuple(m, v1788)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L4
	} else {
		goto L329
	}
L326:
	;
	goto L325
L327:
	;
	v1751 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1751&int32(1) == int32(0) {
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1756 = int32(_a_F_btbuild_0)
	v1758 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1759 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1758 + v1759
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1747)))
	*(*int32)(unsafe.Add(mBase, uint32(v1747))) = v1762 + v1759
	v1766 = int32(0)
	v1768 = int32(_a_F_btbuild_1)
	v1769 = base.AtomicRmwOr32(m, v1766, v1768, v1766)
	*(*int64)(unsafe.Add(mBase, uint32(v1747+int32(96))+232)) = v1744
	v1777 = base.AtomicRmwOr32(m, v1766, v1768, v1766)
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1747)))
	*(*int32)(unsafe.Add(mBase, uint32(v1747))) = v1778 + v1759
	v1784 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1784 - v1759
	goto L326
L329:
	;
	if v1789 != 0 {
		v1669 = v1736
		v1672 = v1789
		v1682 = v1744
		goto L313
	} else {
		goto L330
	}
L330:
	;
	goto L314
L331:
	;
	if v1060 != 0 {
		v1818 = v1060
		goto L2
	} else {
		goto L332
	}
L332:
	;
	v1793 = int32(0)
	v1938 = v1793
	v1939 = v1793
	goto L1
L333:
	;
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v1799 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btbuild_10), v24)
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L4
	} else {
		goto L334
	}
L334:
	;
	F_errfinish(m, int32(_a_F_btbuild_11), int32(325), int32(_a_F_btbuild_12))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L4
	} else {
		goto L335
	}
L335:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L336:
	;
	v1813 = int32(0)
	v1938 = v1813
	v1939 = v1813
	goto L1
L337:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+4))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+28))
	if v1860 == int32(0) {
		goto L340
	} else {
		goto L341
	}
L338:
	;
	v1938 = v1887
	v1939 = v1888
	goto L1
L339:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1841)))
	v1890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1889)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1890) {
		goto L346
	} else {
		goto L347
	}
L340:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1841)))
	v1864 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1863)+16)))
	v1865 = v1863 + v1864
	v1866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1865)+12)))
	v1868 = v1866 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1865)+12)) = uint16(v1868)
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+20))
	v1887 = v1859
	v1888 = v1870
	goto L339
L341:
	;
	goto L342
L342:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1871))) = base.I32_rotr(v1859, int32(16))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+28))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+8))
	F__bt_buildadd(m, v24+int32(48), v1877, v1878, int32(0))
	mBase = m.M
	v1881 = m.ExcPending
	if v1881 != 0 {
		goto L4
	} else {
		goto L343
	}
L343:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+8))
	F_pfree(m, v1882)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L4
	} else {
		goto L344
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1841)+8)) = int32(0)
	v1887 = v1839
	v1888 = v1840
	goto L339
L345:
	;
	v1926 = v1890 - int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1889)+12)) = uint16(v1926)
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+4))
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1841)))
	F_smgr_bulk_write(m, v1928, v1929, v1930, int32(1))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L4
	} else {
		goto L354
	}
L346:
	;
	v1898 = int32(base.Ui32(v1890+int32(_a_F_btbuild_13)) >> (uint(int32(2)) % 32))
	goto L348
L347:
	;
	v1898 = int32(0)
	goto L348
L348:
	;
	if base.Ui32(v1898&int32(_a_F_btbuild_14)) < base.Ui32(int32(2)) {
		goto L345
	} else {
		goto L349
	}
L349:
	;
	v1903 = int32(3)
	v1907 = (v1898 + int32(1)) & int32(_a_F_btbuild_14)
	if base.Ui32(v1907) <= base.Ui32(v1903) {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1910 = v1903
	goto L352
L351:
	;
	v1910 = v1907
	goto L352
L352:
	;
	v1911 = int32(2)
	v1916 = (v1910 - v1911) & int32(_a_F_btbuild_14) << (uint(v1911) % 32)
	if v1916 == int32(0) {
		goto L345
	} else {
		goto L353
	}
L353:
	;
	base.MemoryCopy(m, v1889+int32(24), v1889+int32(28), v1916)
	goto L345
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1841))) = int32(0)
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+28))
	if v1936 != 0 {
		v1839 = v1887
		v1840 = v1888
		v1841 = v1936
		goto L337
	} else {
		goto L355
	}
L355:
	;
	goto L338
L356:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	v1962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1961)+1)))
	F_PageInit(m, v1959, int32(_a_F_btbuild_6), int32(16))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1959)+64)) = uint8(v1962)
	*(*int64)(unsafe.Add(mBase, uint32(v1959)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+44)) = v1939
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+40)) = v1938
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+36)) = v1939
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+32)) = v1938
	*(*int64)(unsafe.Add(mBase, uint32(v1959)+24)) = int64(17180209506)
	v1977 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1959)+16)))
	v1979 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v1959+v1977)+12)) = uint16(v1979)
	v1981 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v1959)+12)) = uint16(v1981)
	goto L357
L357:
	;
	F_smgr_bulk_write(m, v1958, int32(0), v1959, int32(1))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L4
	} else {
		goto L358
	}
L358:
	;
	F_smgr_bulk_finish(m, v1958)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L4
	} else {
		goto L359
	}
L359:
	;
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v1989)))
	F_tuplesort_end(m, v1990)
	mBase = m.M
	v1992 = m.ExcPending
	if v1992 != 0 {
		goto L4
	} else {
		goto L360
	}
L360:
	;
	F_pfree(m, v1989)
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L4
	} else {
		goto L361
	}
L361:
	;
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v1995 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v1995)))
	F_tuplesort_end(m, v1996)
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L4
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v2001 != 0 {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	F_pfree(m, v1995)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L4
	} else {
		goto L366
	}
L366:
	;
	goto L364
L367:
	;
	F__bt_end_parallel(m, v2001)
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L4
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v2005 = F_palloc(m, int32(16))
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L4
	} else {
		goto L371
	}
L370:
	;
	goto L369
L371:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v2005))) = v573
	v2008 = *(*float64)(unsafe.Add(mBase, uint32(v24)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2005)+8)) = v2008
	m.G0 = v24 + int32(96)
	return v2005
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
func F_btcharskipsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v6 int32
	_ = v6
	v2 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(211)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(212)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(255)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v2
	return v2
}
func F_btfloat4fastcmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float32
	_ = v12
	var v13 float32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	v8 = base.I32_wrap_i64(l0)
	v9 = int32(2147483647)
	v10 = v8 & v9
	v11 = base.I32_wrap_i64(l1)
	v12 = base.F32_reinterpret_i32(v11)
	v13 = base.F32_reinterpret_i32(v8)
	v15 = v11 & v9
	if base.Ui32(int32(2139095041)) <= base.Ui32(v15) {
		v25 = base.B2i32(base.Ui32(v10) < base.Ui32(int32(2139095041)))
		v34 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v15))|base.F32_gt(v12, v13))&v25
	} else {
		v20 = int32(1)
		if base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v10))|base.F32_lt(v12, v13) != 0 {
			v34 = v20
		} else {
			v25 = v20
			v34 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v15))|base.F32_gt(v12, v13))&v25
		}
	}
	return v34
}
func F_btfloat84cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v36 int64
	_ = v36
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v7) & v9
	v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = base.F64_promote_f32(v11)
	v15 = base.I64_reinterpret_f64(v12) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v15) {
		v25 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v36 = int64(0) - base.I64_extend_i32_u((base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v25)
	} else {
		v20 = int64(1)
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10) {
			v36 = v20
		} else {
			if base.F64_gt(v7, v12) != 0 {
				v36 = v20
			} else {
				v25 = int32(1)
				v36 = int64(0) - base.I64_extend_i32_u((base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v25)
			}
		}
	}
	return v36
}
func F_bthandler(m *base.Module, l0 int32) int64 {
	return int64(830316)
}
func F_btinitparallelscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	F_LWLockInitialize(m, l0+int32(12), int32(74))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v7
		v9 = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v9
		v12 = l0 + int32(28)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v12))), uint32(v7))
		*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = v9
		return
	}
}
func F_btint2sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(196)
	return int64(0)
}
func F_btnamesortsupport(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14238(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_btoid8cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_s(base.B2i32(base.Ui64(v5) < base.Ui64(v4)) - base.B2i32(base.Ui64(v4) < base.Ui64(v5)))
}
func F_btoid8skipsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v6 int32
	_ = v6
	v2 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(209)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(210)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v2
	return v2
}
func F_btoidfastcmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v5 = base.I32_wrap_i64(l0)
	v6 = base.I32_wrap_i64(l1)
	return base.B2i32(base.Ui32(v6) < base.Ui32(v5)) - base.B2i32(base.Ui32(v5) < base.Ui32(v6))
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
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
							v28 = v21
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
							switch v24 {
							case 0, 5:
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v28 = base.B2i32(v25 != int32(0))
							default:
								v28 = v21
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v28)
						v32 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v32)
						*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v32)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0])))
						if v36 != 0 {
							F_ReleaseBuffer(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
								if v41 != int32(1) {
									if l1 == int32(0) {
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v55 <= int32(0) {
										} else {
											v59 = v55 * int32(56)
											if v59 == int32(0) {
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v62, l1, v59)
											}
										}
									}
									v65 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
									return
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
									if v44 != 0 {
										if l1 == int32(0) {
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v55 <= int32(0) {
											} else {
												v59 = v55 * int32(56)
												if v59 == int32(0) {
												} else {
													v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													base.MemoryCopy(m, v62, l1, v59)
												}
											}
										}
										v65 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
										*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
										return
									} else {
										v46 = F_palloc(m, int32(_a_F_btrescan_0))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
											*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
											if l1 == int32(0) {
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v55 <= int32(0) {
												} else {
													v59 = v55 * int32(56)
													if v59 == int32(0) {
													} else {
														v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														base.MemoryCopy(m, v62, l1, v59)
													}
												}
											}
											v65 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
											*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
											return
										}
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v41 != int32(1) {
								if l1 == int32(0) {
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v55 <= int32(0) {
									} else {
										v59 = v55 * int32(56)
										if v59 == int32(0) {
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v62, l1, v59)
										}
									}
								}
								v65 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
								return
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v44 != 0 {
									if l1 == int32(0) {
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v55 <= int32(0) {
										} else {
											v59 = v55 * int32(56)
											if v59 == int32(0) {
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v62, l1, v59)
											}
										}
									}
									v65 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
									return
								} else {
									v46 = F_palloc(m, int32(_a_F_btrescan_0))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
										*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
										if l1 == int32(0) {
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v55 <= int32(0) {
											} else {
												v59 = v55 * int32(56)
												if v59 == int32(0) {
												} else {
													v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													base.MemoryCopy(m, v62, l1, v59)
												}
											}
										}
										v65 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
										*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
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
						v28 = v21
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						switch v24 {
						case 0, 5:
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v28 = base.B2i32(v25 != int32(0))
						default:
							v28 = v21
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v28)
					v32 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v32)
					*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v32)
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0])))
					if v36 != 0 {
						F_ReleaseBuffer(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v41 != int32(1) {
								if l1 == int32(0) {
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v55 <= int32(0) {
									} else {
										v59 = v55 * int32(56)
										if v59 == int32(0) {
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v62, l1, v59)
										}
									}
								}
								v65 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
								return
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v44 != 0 {
									if l1 == int32(0) {
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v55 <= int32(0) {
										} else {
											v59 = v55 * int32(56)
											if v59 == int32(0) {
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v62, l1, v59)
											}
										}
									}
									v65 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
									return
								} else {
									v46 = F_palloc(m, int32(_a_F_btrescan_0))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
										*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
										if l1 == int32(0) {
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v55 <= int32(0) {
											} else {
												v59 = v55 * int32(56)
												if v59 == int32(0) {
												} else {
													v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													base.MemoryCopy(m, v62, l1, v59)
												}
											}
										}
										v65 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
										*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
										return
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v41 != int32(1) {
							if l1 == int32(0) {
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v55 <= int32(0) {
								} else {
									v59 = v55 * int32(56)
									if v59 == int32(0) {
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v62, l1, v59)
									}
								}
							}
							v65 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
							return
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							if v44 != 0 {
								if l1 == int32(0) {
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v55 <= int32(0) {
									} else {
										v59 = v55 * int32(56)
										if v59 == int32(0) {
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v62, l1, v59)
										}
									}
								}
								v65 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
								return
							} else {
								v46 = F_palloc(m, int32(_a_F_btrescan_0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
									*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
									if l1 == int32(0) {
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v55 <= int32(0) {
										} else {
											v59 = v55 * int32(56)
											if v59 == int32(0) {
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v62, l1, v59)
											}
										}
									}
									v65 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
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
						v28 = v21
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						switch v24 {
						case 0, 5:
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v28 = base.B2i32(v25 != int32(0))
						default:
							v28 = v21
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v28)
					v32 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v32)
					*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v32)
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0])))
					if v36 != 0 {
						F_ReleaseBuffer(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v41 != int32(1) {
								if l1 == int32(0) {
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v55 <= int32(0) {
									} else {
										v59 = v55 * int32(56)
										if v59 == int32(0) {
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v62, l1, v59)
										}
									}
								}
								v65 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
								return
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v44 != 0 {
									if l1 == int32(0) {
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v55 <= int32(0) {
										} else {
											v59 = v55 * int32(56)
											if v59 == int32(0) {
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v62, l1, v59)
											}
										}
									}
									v65 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
									return
								} else {
									v46 = F_palloc(m, int32(_a_F_btrescan_0))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
										*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
										if l1 == int32(0) {
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v55 <= int32(0) {
											} else {
												v59 = v55 * int32(56)
												if v59 == int32(0) {
												} else {
													v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													base.MemoryCopy(m, v62, l1, v59)
												}
											}
										}
										v65 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
										*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
										return
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v41 != int32(1) {
							if l1 == int32(0) {
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v55 <= int32(0) {
								} else {
									v59 = v55 * int32(56)
									if v59 == int32(0) {
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v62, l1, v59)
									}
								}
							}
							v65 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
							return
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							if v44 != 0 {
								if l1 == int32(0) {
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v55 <= int32(0) {
									} else {
										v59 = v55 * int32(56)
										if v59 == int32(0) {
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v62, l1, v59)
										}
									}
								}
								v65 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
								return
							} else {
								v46 = F_palloc(m, int32(_a_F_btrescan_0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
									*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
									if l1 == int32(0) {
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v55 <= int32(0) {
										} else {
											v59 = v55 * int32(56)
											if v59 == int32(0) {
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v62, l1, v59)
											}
										}
									}
									v65 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
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
					v28 = v21
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					switch v24 {
					case 0, 5:
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v28 = base.B2i32(v25 != int32(0))
					default:
						v28 = v21
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v28)
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v32)
				*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v32)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0])))
				if v36 != 0 {
					F_ReleaseBuffer(m, v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v41 != int32(1) {
							if l1 == int32(0) {
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v55 <= int32(0) {
								} else {
									v59 = v55 * int32(56)
									if v59 == int32(0) {
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v62, l1, v59)
									}
								}
							}
							v65 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
							return
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							if v44 != 0 {
								if l1 == int32(0) {
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v55 <= int32(0) {
									} else {
										v59 = v55 * int32(56)
										if v59 == int32(0) {
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v62, l1, v59)
										}
									}
								}
								v65 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
								return
							} else {
								v46 = F_palloc(m, int32(_a_F_btrescan_0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
									*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
									if l1 == int32(0) {
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v55 <= int32(0) {
										} else {
											v59 = v55 * int32(56)
											if v59 == int32(0) {
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												base.MemoryCopy(m, v62, l1, v59)
											}
										}
									}
									v65 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
									return
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v41 != int32(1) {
						if l1 == int32(0) {
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v55 <= int32(0) {
							} else {
								v59 = v55 * int32(56)
								if v59 == int32(0) {
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									base.MemoryCopy(m, v62, l1, v59)
								}
							}
						}
						v65 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
						if v44 != 0 {
							if l1 == int32(0) {
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v55 <= int32(0) {
								} else {
									v59 = v55 * int32(56)
									if v59 == int32(0) {
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v62, l1, v59)
									}
								}
							}
							v65 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
							return
						} else {
							v46 = F_palloc(m, int32(_a_F_btrescan_0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
								*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
								if l1 == int32(0) {
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v55 <= int32(0) {
									} else {
										v59 = v55 * int32(56)
										if v59 == int32(0) {
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											base.MemoryCopy(m, v62, l1, v59)
										}
									}
								}
								v65 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
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
			v28 = v21
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			switch v24 {
			case 0, 5:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = base.B2i32(v25 != int32(0))
			default:
				v28 = v21
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v28)
		v32 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v32)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v32)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0])))
		if v36 != 0 {
			F_ReleaseBuffer(m, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v41 != int32(1) {
					if l1 == int32(0) {
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v55 <= int32(0) {
						} else {
							v59 = v55 * int32(56)
							if v59 == int32(0) {
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								base.MemoryCopy(m, v62, l1, v59)
							}
						}
					}
					v65 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
					if v44 != 0 {
						if l1 == int32(0) {
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v55 <= int32(0) {
							} else {
								v59 = v55 * int32(56)
								if v59 == int32(0) {
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									base.MemoryCopy(m, v62, l1, v59)
								}
							}
						}
						v65 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
						return
					} else {
						v46 = F_palloc(m, int32(_a_F_btrescan_0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
							*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
							if l1 == int32(0) {
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v55 <= int32(0) {
								} else {
									v59 = v55 * int32(56)
									if v59 == int32(0) {
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										base.MemoryCopy(m, v62, l1, v59)
									}
								}
							}
							v65 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
							return
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[0]))) = int64(-4294967296)
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			if v41 != int32(1) {
				if l1 == int32(0) {
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v55 <= int32(0) {
					} else {
						v59 = v55 * int32(56)
						if v59 == int32(0) {
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							base.MemoryCopy(m, v62, l1, v59)
						}
					}
				}
				v65 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
				return
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
				if v44 != 0 {
					if l1 == int32(0) {
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v55 <= int32(0) {
						} else {
							v59 = v55 * int32(56)
							if v59 == int32(0) {
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								base.MemoryCopy(m, v62, l1, v59)
							}
						}
					}
					v65 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
					return
				} else {
					v46 = F_palloc(m, int32(_a_F_btrescan_0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v46 - int32(-8192)
						if l1 == int32(0) {
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v55 <= int32(0) {
							} else {
								v59 = v55 * int32(56)
								if v59 == int32(0) {
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									base.MemoryCopy(m, v62, l1, v59)
								}
							}
						}
						v65 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v65
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v65
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
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
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
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
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
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
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int64
	_ = v414
	var v416 int64
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
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
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
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
	var v626 int32
	_ = v626
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
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v700 int32
	_ = v700
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v822 int32
	_ = v822
	var v838 int32
	_ = v838
	var v847 int32
	_ = v847
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
	var v951 int32
	_ = v951
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
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1019 int32
	_ = v1019
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1105 int32
	_ = v1105
	var v1128 int32
	_ = v1128
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1221 int32
	_ = v1221
	var v1233 int32
	_ = v1233
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1291 int32
	_ = v1291
	var v1330 int32
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1386 int32
	_ = v1386
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1406 int32
	_ = v1406
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1468 int32
	_ = v1468
	var v1475 int32
	_ = v1475
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1490 int64
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int64
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int64
	_ = v1494
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1525 int32
	_ = v1525
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1622 float64
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1634 int32
	_ = v1634
	var v1639 int32
	_ = v1639
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1707 int32
	_ = v1707
	var v1728 int32
	_ = v1728
	var v1767 float64
	_ = v1767
	var v1768 float64
	_ = v1768
	var v1772 int32
	_ = v1772
	var v1812 int32
	_ = v1812
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1842 int32
	_ = v1842
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1867 int32
	_ = v1867
	var v1877 int32
	_ = v1877
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1956 int32
	_ = v1956
	var v1960 int32
	_ = v1960
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1979 int32
	_ = v1979
	var v1983 int32
	_ = v1983
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2009 int32
	_ = v2009
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2023 int32
	_ = v2023
	var v2030 int32
	_ = v2030
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2081 int32
	_ = v2081
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2116 int32
	_ = v2116
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2159 int32
	_ = v2159
	var v2163 int32
	_ = v2163
	var v2169 int32
	_ = v2169
	var v2171 int32
	_ = v2171
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2197 int32
	_ = v2197
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2210 int32
	_ = v2210
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2270 int32
	_ = v2270
	var v2272 int32
	_ = v2272
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2295 int32
	_ = v2295
	var v2297 int32
	_ = v2297
	var v2302 int32
	_ = v2302
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2311 int32
	_ = v2311
	var v2313 int32
	_ = v2313
	var v2317 int32
	_ = v2317
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2353 int32
	_ = v2353
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2367 int32
	_ = v2367
	var v2371 int32
	_ = v2371
	var v2375 int32
	_ = v2375
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2382 int32
	_ = v2382
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2397 int32
	_ = v2397
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2428 int32
	_ = v2428
	var v2433 int32
	_ = v2433
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2443 int32
	_ = v2443
	var v2450 int32
	_ = v2450
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2464 int32
	_ = v2464
	var v2468 int32
	_ = v2468
	var v2476 int32
	_ = v2476
	var v2480 int32
	_ = v2480
	var v2486 int32
	_ = v2486
	var v2488 int32
	_ = v2488
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2508 int32
	_ = v2508
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2521 int32
	_ = v2521
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2529 int32
	_ = v2529
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2538 int32
	_ = v2538
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2550 int32
	_ = v2550
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2564 int32
	_ = v2564
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2569 int32
	_ = v2569
	var v2575 int32
	_ = v2575
	var v2578 int64
	_ = v2578
	var v2579 int32
	_ = v2579
	var v2580 int64
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2583 int64
	_ = v2583
	var v2587 int32
	_ = v2587
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2601 int32
	_ = v2601
	var v2603 int64
	_ = v2603
	var v2608 int32
	_ = v2608
	var v2614 int32
	_ = v2614
	var v2616 int32
	_ = v2616
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2626 int32
	_ = v2626
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2639 int32
	_ = v2639
	var v2687 int32
	_ = v2687
	var v2689 int32
	_ = v2689
	var v2743 int32
	_ = v2743
	var v2749 int32
	_ = v2749
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2763 int32
	_ = v2763
	var v2767 int32
	_ = v2767
	var v2769 int32
	_ = v2769
	var v2773 int32
	_ = v2773
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2785 int32
	_ = v2785
	var v2788 int32
	_ = v2788
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2794 int32
	_ = v2794
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
	var v2832 int32
	_ = v2832
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2846 int32
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2852 int32
	_ = v2852
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2872 int32
	_ = v2872
	var v2880 int32
	_ = v2880
	var __phi2880 int32
	_ = __phi2880
	var v2883 int32
	_ = v2883
	var __phi2883 int32
	_ = __phi2883
	var v2885 int32
	_ = v2885
	var __phi2885 int32
	_ = __phi2885
	var v2895 int32
	_ = v2895
	var __phi2895 int32
	_ = __phi2895
	var v2927 int32
	_ = v2927
	var v2939 int32
	_ = v2939
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2960 int32
	_ = v2960
	var v2965 int32
	_ = v2965
	var v2968 int32
	_ = v2968
	var v2972 int32
	_ = v2972
	var v2974 int32
	_ = v2974
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2981 int32
	_ = v2981
	var v2983 int32
	_ = v2983
	var v2987 int32
	_ = v2987
	var v2993 int32
	_ = v2993
	var v2995 int32
	_ = v2995
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3005 int32
	_ = v3005
	var v3007 int32
	_ = v3007
	var v3014 int32
	_ = v3014
	var v3016 int32
	_ = v3016
	var v3062 int32
	_ = v3062
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3068 int32
	_ = v3068
	var v3074 int32
	_ = v3074
	var v3076 int32
	_ = v3076
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3088 int32
	_ = v3088
	var v3091 int32
	_ = v3091
	var v3093 int32
	_ = v3093
	var v3095 int32
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3121 int32
	_ = v3121
	var v3126 int32
	_ = v3126
	var v3138 int32
	_ = v3138
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3145 int32
	_ = v3145
	var v3146 int32
	_ = v3146
	var v3148 int32
	_ = v3148
	var v3150 int32
	_ = v3150
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3155 int32
	_ = v3155
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3163 int32
	_ = v3163
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3177 int32
	_ = v3177
	var v3178 int32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3180 int32
	_ = v3180
	var v3184 int32
	_ = v3184
	var v3185 int32
	_ = v3185
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3204 int32
	_ = v3204
	var v3209 int32
	_ = v3209
	var v3213 int32
	_ = v3213
	var v3215 int32
	_ = v3215
	var v3217 int32
	_ = v3217
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3224 int32
	_ = v3224
	var v3229 int32
	_ = v3229
	var v3235 int32
	_ = v3235
	var v3237 int32
	_ = v3237
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3252 int32
	_ = v3252
	var v3254 int32
	_ = v3254
	var v3258 int32
	_ = v3258
	var v3264 int32
	_ = v3264
	var v3266 int32
	_ = v3266
	var v3272 int32
	_ = v3272
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3280 int32
	_ = v3280
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3288 int32
	_ = v3288
	var v3295 int32
	_ = v3295
	var v3301 int32
	_ = v3301
	var v3303 int32
	_ = v3303
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3317 int32
	_ = v3317
	var v3323 int32
	_ = v3323
	var v3325 int32
	_ = v3325
	var v3331 int32
	_ = v3331
	var v3332 int32
	_ = v3332
	var v3337 int32
	_ = v3337
	var v3342 int32
	_ = v3342
	var v3344 int32
	_ = v3344
	var v3349 int32
	_ = v3349
	var v3355 int32
	_ = v3355
	var v3357 int32
	_ = v3357
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3365 int64
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3373 int32
	_ = v3373
	var v3375 int32
	_ = v3375
	var v3378 int32
	_ = v3378
	var v3381 int32
	_ = v3381
	var v3383 int32
	_ = v3383
	var v3386 int32
	_ = v3386
	var v3394 int32
	_ = v3394
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3403 int32
	_ = v3403
	var v3405 int32
	_ = v3405
	var v3409 int32
	_ = v3409
	var v3410 int32
	_ = v3410
	var v3411 int32
	_ = v3411
	var v3415 int32
	_ = v3415
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3421 int32
	_ = v3421
	var v3425 int32
	_ = v3425
	var v3429 int32
	_ = v3429
	var v3433 int32
	_ = v3433
	var v3439 int32
	_ = v3439
	var v3451 int32
	_ = v3451
	var v3456 int64
	_ = v3456
	var v3457 int32
	_ = v3457
	var v3461 int32
	_ = v3461
	var v3462 int32
	_ = v3462
	var v3464 int32
	_ = v3464
	var v3466 int32
	_ = v3466
	var v3468 int32
	_ = v3468
	var v3470 int32
	_ = v3470
	var v3472 int32
	_ = v3472
	var v3474 int32
	_ = v3474
	var v3481 int32
	_ = v3481
	var v3484 int64
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3486 int64
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3488 int64
	_ = v3488
	var v3495 int32
	_ = v3495
	var v3501 int32
	_ = v3501
	var v3503 int32
	_ = v3503
	var v3509 int32
	_ = v3509
	var v3511 int64
	_ = v3511
	var v3516 int32
	_ = v3516
	var v3522 int32
	_ = v3522
	var v3524 int32
	_ = v3524
	var v3530 int32
	_ = v3530
	var v3535 int32
	_ = v3535
	var v3541 int32
	_ = v3541
	var v3543 int32
	_ = v3543
	var v3549 int32
	_ = v3549
	var v3556 int32
	_ = v3556
	var v3560 int32
	_ = v3560
	var v3562 int32
	_ = v3562
	var v3566 int32
	_ = v3566
	var v3568 int32
	_ = v3568
	var v3570 int32
	_ = v3570
	var v3575 int32
	_ = v3575
	var v3577 int32
	_ = v3577
	var v3579 int32
	_ = v3579
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3589 int32
	_ = v3589
	var v3593 int32
	_ = v3593
	var v3594 int32
	_ = v3594
	var v3596 int32
	_ = v3596
	var v3598 int32
	_ = v3598
	var v3600 int32
	_ = v3600
	var v3602 int32
	_ = v3602
	var v3604 int32
	_ = v3604
	var v3607 int32
	_ = v3607
	var v3608 int32
	_ = v3608
	var v3610 int32
	_ = v3610
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3624 int32
	_ = v3624
	var v3631 int32
	_ = v3631
	var v3636 int32
	_ = v3636
	var v3642 int32
	_ = v3642
	var v3651 int32
	_ = v3651
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
	var v3799 int32
	_ = v3799
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
	var v3865 int32
	_ = v3865
	var v3866 int32
	_ = v3866
	var v3867 int32
	_ = v3867
	var v3872 int32
	_ = v3872
	var v3874 int32
	_ = v3874
	var v3875 int32
	_ = v3875
	var v3876 int32
	_ = v3876
	var v3878 int32
	_ = v3878
	var v3881 int32
	_ = v3881
	var v3933 int32
	_ = v3933
	var v4027 int32
	_ = v4027
	var v4040 int32
	_ = v4040
	var v4079 int32
	_ = v4079
	var v4095 int32
	_ = v4095
	var v4096 int32
	_ = v4096
	var v4098 int32
	_ = v4098
	var v4099 int32
	_ = v4099
	var v4100 int32
	_ = v4100
	var v4102 int32
	_ = v4102
	var v4107 int32
	_ = v4107
	var v4108 int32
	_ = v4108
	var v4113 int32
	_ = v4113
	var v4114 int32
	_ = v4114
	var v4122 int32
	_ = v4122
	var v4127 int32
	_ = v4127
	var v4131 int32
	_ = v4131
	var v4132 int32
	_ = v4132
	var v4133 int32
	_ = v4133
	var v4143 int32
	_ = v4143
	var v4164 int32
	_ = v4164
	var v4172 int32
	_ = v4172
	var v4174 int32
	_ = v4174
	var v4182 int32
	_ = v4182
	var v4189 int32
	_ = v4189
	var v4193 int32
	_ = v4193
	var v4198 int32
	_ = v4198
	var v4200 int32
	_ = v4200
	var v4201 int32
	_ = v4201
	var v4204 int32
	_ = v4204
	var v4208 int32
	_ = v4208
	var v4210 int32
	_ = v4210
	var v4211 int32
	_ = v4211
	var v4219 int32
	_ = v4219
	var v4220 int32
	_ = v4220
	var v4226 int32
	_ = v4226
	var v4231 int32
	_ = v4231
	var v4233 int32
	_ = v4233
	var v4235 int32
	_ = v4235
	var v4236 int32
	_ = v4236
	var v4238 int32
	_ = v4238
	var v4239 int32
	_ = v4239
	var v4242 int32
	_ = v4242
	var v4243 int32
	_ = v4243
	var v4244 int32
	_ = v4244
	var v4245 int32
	_ = v4245
	var v4246 int32
	_ = v4246
	var v4247 int32
	_ = v4247
	var v4248 int32
	_ = v4248
	var v4253 int32
	_ = v4253
	var v4301 int32
	_ = v4301
	var v4304 int32
	_ = v4304
	var v4305 int32
	_ = v4305
	var v4306 int64
	_ = v4306
	var v4307 int32
	_ = v4307
	var v4308 int32
	_ = v4308
	var v4312 int32
	_ = v4312
	var v4313 int32
	_ = v4313
	var v4314 int32
	_ = v4314
	var v4318 int32
	_ = v4318
	var v4319 int32
	_ = v4319
	var v4371 int32
	_ = v4371
	var v4374 int32
	_ = v4374
	var v4423 int32
	_ = v4423
	var v4474 int32
	_ = v4474
	var v4476 int32
	_ = v4476
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
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+24)))
	if v108 == int32(0) {
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
	v101 = F_palloc_mul(m, int32(16), int32(256))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+32)) = v101
	goto L5
L13:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
	v114 = base.B2i32(v111 == int32(0))
	goto L15
L14:
	;
	v114 = v6
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v119 = int32(0)
	v124 = F_read_stream_begin_relation(m, int32(13), v118, v55, v119, int32(3), v53+int32(16), v119)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v126 = l0
	v127 = l1
	v137 = v53
	v158 = v55
	v166 = v124
	v168 = v114
	goto L17
L17:
	;
	if v168 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	F_read_stream_end(m, v166)
	mBase = m.M
	v4231 = m.ExcPending
	if v4231 != 0 {
		goto L1
	} else {
		goto L748
	}
L19:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+9)))
	if v191 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v179 = F_RelationGetNumberOfBlocksInFork(m, v158, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_LockRelationForExtension(m, v158, int32(7))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v190 = v179
	goto L19
L24:
	;
	v185 = F_RelationGetNumberOfBlocksInFork(m, v158, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_UnlockRelationForExtension(m, v158, int32(7))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v190 = v185
	goto L19
L27:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[2]))
	if v198 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	if base.Ui32(v239) < base.Ui32(v190) {
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
	v202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btvacuumscan[3])))
	if v202&int32(1) == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v207 = int32(_a_F_btvacuumscan_3)
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	v210 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v209 + v210
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v213 + v210
	v217 = int32(0)
	v219 = int32(_a_F_btvacuumscan_4)
	v220 = base.AtomicRmwOr32(m, v217, v219, v217)
	*(*int64)(unsafe.Add(mBase, uint32(v198+int32(120))+232)) = base.I64_extend_i32_u(v190)
	v228 = base.AtomicRmwOr32(m, v217, v219, v217)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v229 + v210
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v235 - v210
	goto L31
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v190
	v242 = v126
	v243 = v127
	v253 = v137
	v274 = v158
	v282 = v166
	v284 = v168
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
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v296 = F_read_stream_next_buffer(m, v282, int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v4182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4132)+9)))
	if v4182 != int32(1) {
		v242 = v4132
		v243 = v4133
		v253 = v4143
		v274 = v4164
		v282 = v4172
		v284 = v4174
		goto L37
	} else {
		goto L743
	}
L41:
	;
	F_UnlockReleaseBuffer(m, v326)
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L1
	} else {
		goto L742
	}
L42:
	;
	v4107 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4108 = m.ExcPending
	if v4108 != 0 {
		goto L1
	} else {
		goto L737
	}
L43:
	;
	if v296 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v253)+24))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v253)+36))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v253)+32))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v253)+28))
	if v296 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	F_read_stream_reset(m, v282)
	mBase = m.M
	v4102 = m.ExcPending
	if v4102 != 0 {
		goto L1
	} else {
		goto L736
	}
L47:
	;
	v323 = v242
	v324 = v243
	v326 = v296
	v330 = v300
	v332 = v322
	v334 = v253
	v354 = v303
	v355 = v274
	v358 = v322
	v363 = v282
	v364 = v302
	v365 = v284
	v366 = v298
	v367 = v299
	v368 = v301
	goto L51
L48:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v307+(v296^int32(-1))*int32(56))+16))
	v322 = v313
	goto L47
L49:
	;
	goto L50
L50:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v316 = int32(56)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v315+v296*v316-v316)+16))
	v322 = v321
	goto L47
L51:
	;
	F__bt_lockbuf(m, v326, int32(1))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v326 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	if v4079 == int32(0) {
		v4132 = v323
		v4133 = v324
		v4143 = v334
		v4164 = v355
		v4172 = v363
		v4174 = v365
		goto L40
	} else {
		goto L733
	}
L55:
	;
	F_UnlockReleaseBuffer(m, v326)
	mBase = m.M
	v4040 = m.ExcPending
	if v4040 != 0 {
		goto L1
	} else {
		goto L732
	}
L56:
	;
	v445 = int32(0)
	v447 = v420 & int32(_a_F_btvacuumscan_5)
	if v447&int32(16) == v445 {
		goto L84
	} else {
		goto L85
	}
L57:
	;
	v4027 = int32(0)
	goto L55
L58:
	;
	F_RecordFreeIndexPage(m, v330, v332)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L83
	}
L59:
	;
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+14)))
	if v394 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v379+(v326^int32(-1))<<(uint(int32(2))%32))))
	v393 = v385
	goto L59
L61:
	;
	goto L62
L62:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v393 = v387 + v326<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	F__bt_checkpage(m, v330, v326)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v332 != v358 {
		goto L42
	} else {
		goto L82
	}
L66:
	;
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+16)))
	v398 = v393 + v397
	v399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+12)))
	if v332 != v358 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v399&int32(17) != int32(1) {
		goto L42
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v399&int32(4) != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	if v399&int32(4) != 0 {
		goto L41
	} else {
		goto L71
	}
L71:
	;
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+14)))
	v408 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+40)))
	if v407 != v408 {
		goto L41
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	if v399&int32(256) != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v420 = v399
	goto L75
L75:
	;
	if v420&int32(4) == int32(0) {
		goto L56
	} else {
		goto L81
	}
L76:
	;
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v393)+24))
	v416 = v414
	goto L78
L77:
	;
	v416 = int64(3)
	goto L78
L78:
	;
	v417 = F_GlobalVisCheckRemovableFullXid(m, v367, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v417 != 0 {
		goto L58
	} else {
		goto L80
	}
L80:
	;
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+12)))
	v420 = v419
	goto L75
L81:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v354)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+28)) = v425 + int32(1)
	goto L57
L82:
	;
	goto L58
L83:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v354)+28))
	v435 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v354)+28)) = v434 + v435
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v354)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+32)) = v438 + v435
	goto L57
L84:
	;
	if v447&int32(1) == int32(0) {
		v4027 = v445
		goto L55
	} else {
		goto L87
	}
L85:
	;
	v1812 = v445
	goto L86
L86:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v334)+44))
	F_MemoryContextReset(m, v1824)
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L1
	} else {
		goto L250
	}
L87:
	;
	F_UnlockBuffer(m, v326)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_LockBufferForCleanup(m, v326)
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
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v482 != 0 {
		goto L100
	} else {
		goto L101
	}
L91:
	;
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+14)))
	if v466 != v462 {
		v479 = int32(0)
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398)+12)))
	if v469&int32(32) != 0 {
		v479 = int32(0)
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if base.Ui32(v472) < base.Ui32(v358) {
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
	v484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v484) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v492 = int32(base.Ui32(v484+int32(_a_F_btvacuumscan_6)) >> (uint(int32(2)) % 32))
	goto L105
L104:
	;
	v492 = int32(0)
	goto L105
L105:
	;
	v493 = float64(0)
	if v364 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v879 = int32(0)
	if base.B2i32(v847 <= v879)&base.B2i32(v838 <= v879) == v879 {
		goto L146
	} else {
		goto L147
	}
L107:
	;
	v838 = int32(0)
	v847 = v460
	v876 = v493
	v878 = float64(0)
	goto L106
L108:
	;
	goto L109
L109:
	;
	v498 = int32(0)
	v501 = v492 & int32(_a_F_btvacuumscan_5)
	if base.Ui32(v501) < base.Ui32(v483) {
		v838 = v498
		v847 = v460
		v876 = v493
		v878 = float64(0)
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v505 = int32(0)
	v515 = v483
	v517 = v498
	v521 = v505
	v526 = v460
	v532 = v505
	goto L111
L111:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v393+int32(20)+v515&int32(_a_F_btvacuumscan_5)<<(uint(int32(2))%32))))
	v565 = v393 + v562&int32(_a_F_btvacuumscan_7)
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+7)))
	if v566&int32(32) != 0 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v838 = v781
	v847 = v790
	v876 = base.F64_convert_i32_s(v785)
	v878 = base.F64_convert_i32_s(v796)
	goto L106
L113:
	;
	v822 = v515 + int32(1)
	if base.Ui32(v822&int32(_a_F_btvacuumscan_5)) <= base.Ui32(v501) {
		v515 = v822
		v517 = v781
		v521 = v785
		v526 = v790
		v532 = v796
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
	v573 = m.T0[v364].(func(*base.Module, int32, int32) int32)(m, v565, v368)
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
	*(*uint16)(unsafe.Add(mBase, uint32(v334+int32(1696)+v526<<(uint(v577)%32)))) = uint16(v515)
	v781 = v517
	v785 = v521 + v577
	v790 = v526 + v577
	v796 = v532
	goto L113
L121:
	;
	goto L122
L122:
	;
	v781 = v517
	v785 = v521
	v790 = v526
	v796 = v532 + int32(1)
	goto L113
L123:
	;
	v781 = v730
	v785 = v734
	v790 = v739
	v796 = v741 + v532
	goto L113
L124:
	;
	v730 = v517
	v734 = v521
	v739 = v526
	v741 = int32(0)
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
	v626 = v602
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
		v730 = v517
		v734 = v521
		v739 = v526
		v741 = v685
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
		v626 = v685
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
	v685 = v626 + int32(1)
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
	v685 = v626
	goto L129
L137:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v668)+8)) = uint16(v607)
	v671 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v668)+6)) = uint16(v671)
	*(*uint16)(unsafe.Add(mBase, uint32(v668)+4)) = uint16(v515)
	*(*int32)(unsafe.Add(mBase, uint32(v668))) = v565
	v683 = v668
	v685 = v626
	goto L129
L138:
	;
	goto L128
L139:
	;
	if int32(0) < v685 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334-int32(-64)+v517<<(uint(int32(2))%32)))) = v683
	v700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565)+4)))
	v730 = v517 + int32(1)
	v734 = v521 - v685 + v700&int32(4095)
	v739 = v526
	v741 = v685
	goto L123
L141:
	;
	goto L142
L142:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v334+int32(1696)+v526<<(uint(int32(1))%32)))) = uint16(v515)
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
	v730 = v517
	v734 = v521 + v712&int32(4095)
	v739 = v526 + int32(1)
	v741 = v685
	goto L123
L144:
	;
	goto L112
L145:
	;
	if base.Ui32(v483) <= base.Ui32(v1728&int32(_a_F_btvacuumscan_5)) {
		goto L242
	} else {
		goto L243
	}
L146:
	;
	v887 = v334 + int32(1696)
	v889 = v334 - int32(-64)
	v890 = int32(0)
	v895 = m.G0
	v897 = v895 - int32(832)
	m.G0 = v897
	if v326 < v890 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	goto L148
L148:
	;
	v1698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+40)))
	if v1698 == int32(0) {
		v1728 = v492
		goto L145
	} else {
		goto L239
	}
L149:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
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
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v902+(v326^int32(-1))<<(uint(int32(2))%32))))
	v916 = v908
	goto L149
L151:
	;
	goto L152
L152:
	;
	v910 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v916 = v910 + v326<<(uint(int32(13))%32) + int32(-8192)
	goto L149
L153:
	;
	if int32(0) < v838 {
		goto L159
	} else {
		goto L160
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
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v330)+32))
	if v928 != 0 {
		v932 = int32(0)
		goto L153
	} else {
		goto L156
	}
L156:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v330)+40))
	v932 = base.B2i32(v929 == int32(0))
	goto L153
L157:
	;
	if int32(0) < v847 {
		goto L197
	} else {
		goto L198
	}
L158:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L1
	} else {
		goto L190
	}
L159:
	;
	v943 = v890
	v951 = v890
	goto L162
L160:
	;
	goto L161
L161:
	;
	v1350 = int32(_a_F_btvacuumscan_3)
	v1352 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1352 + int32(1)
	v1394 = v890
	v1406 = v890
	goto L157
L162:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v889+v951<<(uint(int32(2))%32))))
	F__bt_update_posting(m, v988)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L1
	} else {
		goto L164
	}
L163:
	;
	if v932 != 0 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v991 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v988)+6)))
	v994 = int32(1)
	v997 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v988)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v897+int32(16)+v951<<(uint(v994)%32)))) = uint16(v997)
	v1003 = v943 + v991<<(uint(v994)%32) + int32(2)
	v1005 = v951 + v994
	if v1005 != v838 {
		v943 = v1003
		v951 = v1005
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v1008 = int32(0)
	v1009 = F_palloc(m, v1003)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	v1221 = v890
	v1233 = v890
	goto L168
L168:
	;
	v1269 = int32(_a_F_btvacuumscan_3)
	v1271 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1271 + int32(1)
	v1291 = int32(0)
	goto L185
L169:
	;
	if v838 != int32(1) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v1221 = v1003
	v1233 = v1009
	goto L168
L171:
	;
	v1019 = v890
	v1042 = v1008
	v1047 = v890
	goto L174
L172:
	;
	v1105 = v890
	v1128 = v1008
	goto L173
L173:
	;
	v1153 = v1105 + v1009
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v889+v1128<<(uint(int32(2))%32))))
	v1158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1157)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1153))) = uint16(v1158)
	v1161 = v1158 << (uint(int32(1)) % 32)
	if v1161 == int32(0) {
		goto L170
	} else {
		goto L184
	}
L174:
	;
	v1068 = int32(2)
	v1070 = v889 + v1042<<(uint(v1068)%32)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1070)))
	v1072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1071)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1019+v1009))) = uint16(v1072)
	v1075 = v1019 + v1068
	v1077 = v1072 << (uint(int32(1)) % 32)
	if v1077 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	if v838&int32(1) == int32(0) {
		goto L170
	} else {
		goto L183
	}
L176:
	;
	base.MemoryCopy(m, v1009+v1075, v1071+int32(8), v1077)
	goto L178
L177:
	;
	goto L178
L178:
	;
	v1082 = v1077 + v1075
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+4))
	v1085 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1009+v1082))) = uint16(v1085)
	v1088 = v1082 + int32(2)
	v1090 = v1085 << (uint(int32(1)) % 32)
	if v1090 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	base.MemoryCopy(m, v1088+v1009, v1084+int32(8), v1090)
	goto L181
L180:
	;
	goto L181
L181:
	;
	v1095 = v1090 + v1088
	v1096 = int32(2)
	v1097 = v1042 + v1096
	v1099 = v1047 + v1096
	if v1099 != v838&int32(2147483646) {
		v1019 = v1095
		v1042 = v1097
		v1047 = v1099
		goto L174
	} else {
		goto L182
	}
L182:
	;
	goto L175
L183:
	;
	v1105 = v1095
	v1128 = v1097
	goto L173
L184:
	;
	base.MemoryCopy(m, v1153+int32(2), v1157+int32(8), v1161)
	goto L170
L185:
	;
	v1330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v897+int32(16)+v1291<<(uint(int32(1))%32)))))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v889+v1291<<(uint(int32(2))%32))))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1334)))
	v1336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1335)+6)))
	v1343 = F_PageIndexTupleOverwrite(m, v916, v1330, v1335, (v1336&int32(_a_F_btvacuumscan_8)+int32(7))&int32(_a_F_btvacuumscan_9))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L1
	} else {
		goto L187
	}
L186:
	;
	v1394 = v1221
	v1406 = v1233
	goto L157
L187:
	;
	if v1343 == int32(0) {
		goto L158
	} else {
		goto L188
	}
L188:
	;
	v1348 = v1291 + int32(1)
	if v1348 != v838 {
		v1291 = v1348
		goto L185
	} else {
		goto L189
	}
L189:
	;
	goto L186
L190:
	;
	if v326 < int32(0) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v897))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v897)+4)) = v1379 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_10), v897)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L1
	} else {
		goto L195
	}
L192:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1363+(v326^int32(-1))*int32(56))+16))
	v1378 = v1369
	goto L191
L193:
	;
	goto L194
L194:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1372 = int32(56)
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1371+v326*v1372-v1372)+16))
	v1378 = v1377
	goto L191
L195:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(1228), int32(_a_F_btvacuumscan_12))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L1
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
	F_PageIndexMultiDelete(m, v916, v887, v847)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L1
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v1446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v916)+16)))
	v1447 = v916 + v1446
	v1448 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1447)+14)) = uint16(v1448)
	v1450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1447)+12)))
	v1452 = v1450 & int32(_a_F_btvacuumscan_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v1447)+12)) = uint16(v1452)
	F_MarkBufferDirty(m, v326)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L1
	} else {
		goto L201
	}
L200:
	;
	goto L199
L201:
	;
	if v932 != 0 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v916))) = base.I64_rotl(v1494, int64(32))
	v1499 = int32(_a_F_btvacuumscan_3)
	v1501 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1501 - int32(1)
	if v1406 != 0 {
		goto L220
	} else {
		goto L221
	}
L203:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v897)+14)) = uint16(v838)
	*(*uint16)(unsafe.Add(mBase, uint32(v897)+12)) = uint16(v847)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v1492 = F_XLogGetFakeLSN(m, v330)
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L1
	} else {
		goto L219
	}
L206:
	;
	F_XLogRegisterBuffer(m, int32(0), v326, int32(8))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_XLogRegisterData(m, v897+int32(12), int32(4))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	if int32(0) < v847 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_XLogRegisterBufData(m, int32(0), v887, v847<<(uint(int32(1))%32))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L1
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	if int32(0) < v838 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	goto L211
L213:
	;
	F_XLogRegisterBufData(m, int32(0), v897+int32(16), v838<<(uint(int32(1))%32))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1490 = F_XLogInsert(m, int32(11), int32(192))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L1
	} else {
		goto L218
	}
L216:
	;
	F_XLogRegisterBufData(m, int32(0), v1406, v1394)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	goto L215
L218:
	;
	v1494 = v1490
	goto L202
L219:
	;
	v1494 = v1492
	goto L202
L220:
	;
	F_pfree(m, v1406)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L1
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	if int32(0) < v838 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	goto L222
L224:
	;
	v1525 = int32(0)
	goto L227
L225:
	;
	goto L226
L226:
	;
	m.G0 = v897 + int32(832)
	v1622 = *(*float64)(unsafe.Add(mBase, uint32(v354)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v354)+16)) = base.F64_add(v876, v1622)
	v1625 = int32(0)
	v1626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1626) {
		goto L231
	} else {
		goto L232
	}
L227:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v889+v1525<<(uint(int32(2))%32))))
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1562)))
	F_pfree(m, v1563)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L1
	} else {
		goto L229
	}
L228:
	;
	goto L226
L229:
	;
	v1567 = v1525 + int32(1)
	if v1567 != v838 {
		v1525 = v1567
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	v1634 = int32(base.Ui32(v1626+int32(_a_F_btvacuumscan_6)) >> (uint(int32(2)) % 32))
	goto L233
L232:
	;
	v1634 = v1625
	goto L233
L233:
	;
	if v838 <= int32(0) {
		v1728 = v1634
		goto L145
	} else {
		goto L234
	}
L234:
	;
	v1639 = v1625
	goto L235
L235:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v334-int32(-64)+v1639<<(uint(int32(2))%32))))
	F_pfree(m, v1692)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L1
	} else {
		goto L237
	}
L236:
	;
	v1728 = v1634
	goto L145
L237:
	;
	v1696 = v1639 + int32(1)
	if v1696 != v838 {
		v1639 = v1696
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	v1701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+14)))
	if v1701 != v1698 {
		v1728 = v492
		goto L145
	} else {
		goto L240
	}
L240:
	;
	v1703 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v398)+14)) = uint16(v1703)
	F_MarkBufferDirtyHint(m, v326, int32(1))
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v1728 = v492
	goto L145
L242:
	;
	if v364 != 0 {
		goto L245
	} else {
		goto L246
	}
L243:
	;
	goto L244
L244:
	;
	if v332 != v358 {
		v4027 = v479
		goto L55
	} else {
		goto L249
	}
L245:
	;
	v1767 = v878
	goto L247
L246:
	;
	v1767 = base.F64_convert_i32_u((v1728 - v483 + int32(1)) & int32(_a_F_btvacuumscan_5))
	goto L247
L247:
	;
	v1768 = *(*float64)(unsafe.Add(mBase, uint32(v354)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v354)+8)) = base.F64_add(v1767, v1768)
	F_UnlockReleaseBuffer(m, v326)
	mBase = m.M
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	v4079 = v479
	goto L54
L249:
	;
	v1812 = v479
	goto L86
L250:
	;
	v1827 = int32(_a_F_btvacuumscan_14)
	v1828 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0]))
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v334)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0])) = v1830
	v1833 = v334 + int32(24)
	v1834 = int32(0)
	v1835 = m.G0
	v1837 = v1835 - int32(288)
	m.G0 = v1837
	if v326 < v1834 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v1867 = v326
	v1877 = v1834
	goto L256
L252:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1842+(v326^int32(-1))*int32(56))+16))
	v1857 = v1848
	goto L251
L253:
	;
	goto L254
L254:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1851 = int32(56)
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1850+v326*v1851-v1851)+16))
	v1857 = v1856
	goto L251
L255:
	;
	m.G0 = v1837 + int32(288)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0])) = v1828
	v4079 = v1812
	goto L54
L256:
	;
	v1908 = int32(0)
	v1909 = base.B2i32(v1908 <= v1867)
	if v1909 == v1908 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	F_UnlockReleaseBuffer(m, v1867)
	mBase = m.M
	v3933 = m.ExcPending
	if v3933 != 0 {
		goto L1
	} else {
		goto L731
	}
L258:
	;
	goto L257
L259:
	;
	v1928 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1927)+16)))
	v1929 = v1928 + v1927
	v1930 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1929)+12)))
	if v1930&int32(5) != int32(1) {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1913+(v1867^int32(-1))<<(uint(int32(2))%32))))
	v1927 = v1919
	goto L259
L261:
	;
	goto L262
L262:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v1927 = v1921 + v1867<<(uint(int32(13))%32) + int32(-8192)
	goto L259
L263:
	;
	if v1930&int32(16) == int32(0) {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	goto L265
L265:
	;
	if v1930&int32(2) != 0 {
		goto L258
	} else {
		goto L284
	}
L266:
	;
	v1966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1929)+12)))
	if v1966&int32(4) == int32(0) {
		goto L258
	} else {
		goto L274
	}
L267:
	;
	v1941 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	if v1941 == int32(0) {
		goto L266
	} else {
		goto L269
	}
L269:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+208)) = v1948 + int32(4)
	F_errmsg(m, int32(_a_F_btvacuumscan_15), v1837+int32(208))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	F_errhint(m, int32(_a_F_btvacuumscan_16), int32(0))
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(1893), int32(_a_F_btvacuumscan_0))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	goto L266
L274:
	;
	v1973 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	if v1973 == int32(0) {
		goto L258
	} else {
		goto L276
	}
L276:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	if v1867 < int32(0) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+196)) = v1857
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+192)) = v1998
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+200)) = v1999 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_17), v1837+int32(192))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L1
	} else {
		goto L282
	}
L279:
	;
	v1983 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v1983+(v1867^int32(-1))*int32(56))+16))
	v1998 = v1989
	goto L278
L280:
	;
	goto L281
L281:
	;
	v1991 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1992 = int32(56)
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1991+v1867*v1992-v1992)+16))
	v1998 = v1997
	goto L278
L282:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(1901), int32(_a_F_btvacuumscan_0))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	goto L258
L284:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1929)+4))
	if base.B2i32(v2017 == int32(0))|v1930&int32(128) != 0 {
		goto L258
	} else {
		goto L285
	}
L285:
	;
	v2023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1927)+12)))
	v2030 = int32(0)
	if base.B2i32(base.Ui32(v2023) < base.Ui32(int32(25)))|base.B2i32((v2023+int32(_a_F_btvacuumscan_6))&int32(_a_F_btvacuumscan_18) == v2030) == v2030 {
		goto L258
	} else {
		goto L286
	}
L286:
	;
	if v1930&int32(16) == int32(0) {
		goto L293
	} else {
		goto L294
	}
L287:
	;
	v3865 = F__bt_mkscankey(m, v330, v2045)
	mBase = m.M
	v3866 = m.ExcPending
	if v3866 != 0 {
		goto L1
	} else {
		goto L727
	}
L288:
	;
	v3843 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3844 = m.ExcPending
	if v3844 != 0 {
		goto L1
	} else {
		goto L722
	}
L289:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3774 = m.ExcPending
	if v3774 != 0 {
		goto L1
	} else {
		goto L719
	}
L290:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3749 = m.ExcPending
	if v3749 != 0 {
		goto L1
	} else {
		goto L715
	}
L291:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L1
	} else {
		goto L712
	}
L292:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3716 = m.ExcPending
	if v3716 != 0 {
		goto L1
	} else {
		goto L709
	}
L293:
	;
	if v1877 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	v2639 = v1930
	goto L295
L295:
	;
	if v2639&int32(16) != 0 {
		goto L442
	} else {
		goto L443
	}
L296:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v1927)+24))
	v2045 = F_CopyIndexTuple(m, v1927+v2041&int32(_a_F_btvacuumscan_7))
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L1
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v1833)))
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+4))
	if v1909 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L299:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1929)))
	if v1867 < int32(0) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	F_UnlockBuffer(m, v1867)
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L1
	} else {
		goto L304
	}
L301:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v2051+(v1867^int32(-1))*int32(56))+16))
	v2066 = v2057
	goto L300
L302:
	;
	goto L303
L303:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2060 = int32(56)
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v2059+v1867*v2060-v2060)+16))
	v2066 = v2065
	goto L300
L304:
	;
	if v2047 == int32(0) {
		goto L287
	} else {
		goto L305
	}
L305:
	;
	v2071 = F_ReadBuffer(m, v330, v2047)
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_LockBufferInternal(m, v2071, int32(1))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	F__bt_checkpage(m, v330, v2071)
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	if v2071 < int32(0) {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	v2096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2095)+16)))
	v2097 = v2096 + v2095
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2097)+4))
	if v2066 != v2098 {
		goto L313
	} else {
		goto L314
	}
L310:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v2081+(v2071^int32(-1))<<(uint(int32(2))%32))))
	v2095 = v2087
	goto L309
L311:
	;
	goto L312
L312:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2095 = v2089 + v2071<<(uint(int32(13))%32) + int32(-8192)
	goto L309
L313:
	;
	F_UnlockReleaseBuffer(m, v2071)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L1
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	v2102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2097)+12)))
	F_UnlockReleaseBuffer(m, v2071)
	mBase = m.M
	v2104 = m.ExcPending
	if v2104 != 0 {
		goto L1
	} else {
		goto L317
	}
L316:
	;
	goto L287
L317:
	;
	if v2102&int32(128) == int32(0) {
		goto L287
	} else {
		goto L318
	}
L318:
	;
	F_ReleaseBuffer(m, v1867)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	goto L255
L320:
	;
	v2131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2130)+16)))
	if v1867 < int32(0) {
		goto L325
	} else {
		goto L326
	}
L321:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v2116+(v1867^int32(-1))<<(uint(int32(2))%32))))
	v2130 = v2122
	goto L320
L322:
	;
	goto L323
L323:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2130 = v2124 + v1867<<(uint(int32(13))%32) + int32(-8192)
	goto L320
L324:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2131+v2130)+4))
	v2153 = F_ReadBuffer(m, v330, v2152)
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L1
	} else {
		goto L328
	}
L325:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v2135+(v1867^int32(-1))*int32(56))+16))
	v2150 = v2141
	goto L324
L326:
	;
	goto L327
L327:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2144 = int32(56)
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2143+v1867*v2144-v2144)+16))
	v2150 = v2149
	goto L324
L328:
	;
	F_LockBufferInternal(m, v2153, int32(1))
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	F__bt_checkpage(m, v330, v2153)
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	if v2153 < int32(0) {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	v2178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2177)+16)))
	v2180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2178+v2177)+12)))
	F_UnlockReleaseBuffer(m, v2153)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L1
	} else {
		goto L335
	}
L332:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v2163+(v2153^int32(-1))<<(uint(int32(2))%32))))
	v2177 = v2169
	goto L331
L333:
	;
	goto L334
L334:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2177 = v2171 + v2153<<(uint(int32(13))%32) + int32(-8192)
	goto L331
L335:
	;
	if v2180&int32(16) != 0 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v2187 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L1
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	v2203 = F__bt_getstackbuf(m, v330, v2112, v1877, v2150)
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L1
	} else {
		goto L343
	}
L339:
	;
	if v2187 == int32(0) {
		goto L258
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+180)) = v2152
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+176)) = v2150
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_19), v1837+int32(176))
	mBase = m.M
	v2197 = m.ExcPending
	if v2197 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2163), int32(_a_F_btvacuumscan_20))
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	goto L258
L343:
	;
	if v2203 == int32(0) {
		v3799 = v2150
		goto L288
	} else {
		goto L344
	}
L344:
	;
	v2210 = v1877
	v2213 = v2203
	v2215 = v2150
	v2217 = v2152
	goto L345
L345:
	;
	v2257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2210)+4)))
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v2210)))
	v2259 = int32(0)
	v2260 = base.B2i32(v2259 <= v2213)
	if v2260 == v2259 {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	if v2260 == int32(0) {
		goto L377
	} else {
		goto L378
	}
L347:
	;
	v2279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2278)+16)))
	v2280 = v2279 + v2278
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v2280)))
	v2282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2278)+12)))
	if base.B2i32(base.Ui32(int32(25)) <= base.Ui32(v2282))&base.B2i32(base.Ui32(v2257) < base.Ui32(int32(base.Ui32(v2282+int32(_a_F_btvacuumscan_6))>>(uint(int32(2))%32))&int32(_a_F_btvacuumscan_5))) == int32(0) {
		goto L351
	} else {
		goto L352
	}
L348:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v2264+(v2213^int32(-1))<<(uint(int32(2))%32))))
	v2278 = v2270
	goto L347
L349:
	;
	goto L350
L350:
	;
	v2272 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2278 = v2272 + v2213<<(uint(int32(13))%32) + int32(-8192)
	goto L347
L351:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2280)+4))
	F_UnlockReleaseBuffer(m, v2213)
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L1
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	goto L346
L354:
	;
	if v2295 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v2302 = int32(2)
	goto L357
L356:
	;
	v2302 = int32(1)
	goto L357
L357:
	;
	if base.B2i32(v2295 == int32(0))|base.B2i32(v2302 != v2257) != 0 {
		goto L258
	} else {
		goto L358
	}
L358:
	;
	if v2281 == int32(0) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+8))
	v2346 = F__bt_getstackbuf(m, v330, v2112, v2345, v2258)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L1
	} else {
		goto L374
	}
L360:
	;
	v2307 = F_ReadBuffer(m, v330, v2281)
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	F_LockBufferInternal(m, v2307, int32(1))
	mBase = m.M
	v2311 = m.ExcPending
	if v2311 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	F__bt_checkpage(m, v330, v2307)
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	if v2307 < int32(0) {
		goto L365
	} else {
		goto L366
	}
L364:
	;
	v2332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2331)+16)))
	v2333 = v2332 + v2331
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v2333)+4))
	if v2258 != v2334 {
		goto L368
	} else {
		goto L369
	}
L365:
	;
	v2317 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v2317+(v2307^int32(-1))<<(uint(int32(2))%32))))
	v2331 = v2323
	goto L364
L366:
	;
	goto L367
L367:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2331 = v2325 + v2307<<(uint(int32(13))%32) + int32(-8192)
	goto L364
L368:
	;
	F_UnlockReleaseBuffer(m, v2307)
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L1
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	v2338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2333)+12)))
	F_UnlockReleaseBuffer(m, v2307)
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L1
	} else {
		goto L372
	}
L371:
	;
	goto L359
L372:
	;
	if v2338&int32(128) != 0 {
		goto L258
	} else {
		goto L373
	}
L373:
	;
	goto L359
L374:
	;
	if v2346 == int32(0) {
		v3799 = v2258
		goto L288
	} else {
		goto L375
	}
L375:
	;
	v2210 = v2345
	v2213 = v2346
	v2215 = v2258
	v2217 = v2295
	goto L345
L376:
	;
	v2371 = (v2257 + int32(1)) & int32(_a_F_btvacuumscan_5)
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(v2367+v2371<<(uint(int32(2))%32))+20))
	v2378 = v2375&int32(_a_F_btvacuumscan_7) + v2367
	v2379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2378))))
	v2382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2378)+2)))
	if v2217 != v2379<<(uint(int32(16))%32)|v2382 {
		goto L380
	} else {
		goto L381
	}
L377:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2359 = *(*int32)(unsafe.Add(mBase, uint32(v2353+(v2213^int32(-1))<<(uint(int32(2))%32))))
	v2367 = v2359
	goto L376
L378:
	;
	goto L379
L379:
	;
	v2361 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2367 = v2361 + v2213<<(uint(int32(13))%32) + int32(-8192)
	goto L376
L380:
	;
	v2387 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L1
	} else {
		goto L383
	}
L381:
	;
	goto L382
L382:
	;
	F_PredicateLockPageSplit(m, v330, v2150, v2152)
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L1
	} else {
		goto L395
	}
L383:
	;
	if v2387 != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L1
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	F_UnlockReleaseBuffer(m, v2213)
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L1
	} else {
		goto L394
	}
L387:
	;
	v2392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2378)+2)))
	v2393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2378))))
	if v2213 < int32(0) {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+160)) = v2413 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+156)) = v2412
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+152)) = v2392 | v2393<<(uint(int32(16))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+148)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+144)) = v2217
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_21), v1837+int32(144))
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L1
	} else {
		goto L392
	}
L389:
	;
	v2397 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2397+(v2213^int32(-1))*int32(56))+16))
	v2412 = v2403
	goto L388
L390:
	;
	goto L391
L391:
	;
	v2405 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2406 = int32(56)
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v2405+v2213*v2406-v2406)+16))
	v2412 = v2411
	goto L388
L392:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2222), int32(_a_F_btvacuumscan_20))
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	goto L386
L394:
	;
	goto L258
L395:
	;
	v2441 = int32(_a_F_btvacuumscan_3)
	v2443 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v2443 + int32(1)
	if v2260 == int32(0) {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2464+v2257<<(uint(int32(2))%32))+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2468&int32(_a_F_btvacuumscan_7)+v2464))) = base.I32_rotr(v2217, int32(16))
	F_PageIndexTupleDelete(m, v2464, v2371)
	mBase = m.M
	v2476 = m.ExcPending
	if v2476 != 0 {
		goto L1
	} else {
		goto L400
	}
L397:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2450+(v2213^int32(-1))<<(uint(int32(2))%32))))
	v2464 = v2456
	goto L396
L398:
	;
	goto L399
L399:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2464 = v2458 + v2213<<(uint(int32(13))%32) + int32(-8192)
	goto L396
L400:
	;
	if v1909 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L401:
	;
	v2495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2494)+16)))
	v2496 = v2495 + v2494
	v2497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2496)+12)))
	v2499 = v2497 | int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v2496)+12)) = uint16(v2499)
	v2502 = base.B2i32(v2215 == v2150)
	if v2215 == v2150 {
		goto L405
	} else {
		goto L406
	}
L402:
	;
	v2480 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(v2480+(v1867^int32(-1))<<(uint(int32(2))%32))))
	v2494 = v2486
	goto L401
L403:
	;
	goto L404
L404:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2494 = v2488 + v1867<<(uint(int32(13))%32) + int32(-8192)
	goto L401
L405:
	;
	v2503 = int32(-1)
	goto L407
L406:
	;
	v2503 = v2215
	goto L407
L407:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1837)+222)) = uint16(v2503)
	if v2215 == v2150 {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v2508 = int32(-1)
	goto L410
L409:
	;
	v2508 = int32(base.Ui32(v2215) >> (uint(int32(16)) % 32))
	goto L410
L410:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1837)+220)) = uint16(v2508)
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+224)) = int32(537395200)
	v2516 = F_PageIndexTupleOverwrite(m, v2494, int32(1), v1837+int32(220), int32(8))
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	if v2516 == int32(0) {
		goto L292
	} else {
		goto L412
	}
L412:
	;
	F_MarkBufferDirty(m, v2213)
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	F_MarkBufferDirty(m, v1867)
	mBase = m.M
	v2523 = m.ExcPending
	if v2523 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	v2525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2524)+118)))
	if v2525 != int32(112) {
		goto L416
	} else {
		goto L417
	}
L415:
	;
	if v2260 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L416:
	;
	v2580 = F_XLogGetFakeLSN(m, v330)
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		goto L1
	} else {
		goto L432
	}
L417:
	;
	v2529 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if v2529 <= int32(0) {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(v330)+32))
	if v2532 != 0 {
		goto L416
	} else {
		goto L421
	}
L419:
	;
	goto L420
L420:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1837)+248)) = uint16(v2257)
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+252)) = v2150
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+264)) = v2503
	F_XLogBeginInsert(m)
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L1
	} else {
		goto L423
	}
L421:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(v330)+40))
	if v2533 != 0 {
		goto L416
	} else {
		goto L422
	}
L422:
	;
	goto L420
L423:
	;
	F_XLogRegisterBuffer(m, int32(0), v1867, int32(6))
	mBase = m.M
	v2542 = m.ExcPending
	if v2542 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	F_XLogRegisterBuffer(m, int32(1), v2213, int32(8))
	mBase = m.M
	v2546 = m.ExcPending
	if v2546 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	if v1909 == int32(0) {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	v2565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2564)+16)))
	v2566 = v2565 + v2564
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v2566)))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+256)) = v2567
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+260)) = v2569
	F_XLogRegisterData(m, v1837+int32(248), int32(20))
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L1
	} else {
		goto L430
	}
L427:
	;
	v2550 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v2550+(v1867^int32(-1))<<(uint(int32(2))%32))))
	v2564 = v2556
	goto L426
L428:
	;
	goto L429
L429:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2564 = v2558 + v1867<<(uint(int32(13))%32) + int32(-8192)
	goto L426
L430:
	;
	v2578 = F_XLogInsert(m, int32(11), int32(176))
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	v2583 = v2578
	goto L415
L432:
	;
	v2583 = v2580
	goto L415
L433:
	;
	v2603 = base.I64_rotl(v2583, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v2601))) = v2603
	if v1909 == int32(0) {
		goto L438
	} else {
		goto L439
	}
L434:
	;
	v2587 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2593 = *(*int32)(unsafe.Add(mBase, uint32(v2587+(v2213^int32(-1))<<(uint(int32(2))%32))))
	v2601 = v2593
	goto L433
L435:
	;
	goto L436
L436:
	;
	v2595 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2601 = v2595 + v2213<<(uint(int32(13))%32) + int32(-8192)
	goto L433
L437:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2622))) = v2603
	v2624 = int32(_a_F_btvacuumscan_3)
	v2626 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v2626 - int32(1)
	F_UnlockReleaseBuffer(m, v2213)
	mBase = m.M
	v2631 = m.ExcPending
	if v2631 != 0 {
		goto L1
	} else {
		goto L441
	}
L438:
	;
	v2608 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v2608+(v1867^int32(-1))<<(uint(int32(2))%32))))
	v2622 = v2614
	goto L437
L439:
	;
	goto L440
L440:
	;
	v2616 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2622 = v2616 + v1867<<(uint(int32(13))%32) + int32(-8192)
	goto L437
L441:
	;
	v2632 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1929)+12)))
	v2639 = v2632
	goto L295
L442:
	;
	v2687 = v1867 ^ int32(-1)
	v2689 = v1867 << (uint(int32(13)) % 32)
	goto L445
L443:
	;
	v3651 = int32(0)
	goto L444
L444:
	;
	v3697 = *(*int32)(unsafe.Add(mBase, uint32(v1929)+4))
	F_UnlockReleaseBuffer(m, v1867)
	mBase = m.M
	v3699 = m.ExcPending
	if v3699 != 0 {
		goto L1
	} else {
		goto L700
	}
L445:
	;
	if v1867 < int32(0) {
		goto L448
	} else {
		goto L449
	}
L446:
	;
	v3636 = int32(2)
	if v3224 != 0 {
		goto L697
	} else {
		goto L698
	}
L447:
	;
	v2759 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+4))
	if v1909 == int32(0) {
		goto L452
	} else {
		goto L453
	}
L448:
	;
	v2743 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2749 = *(*int32)(unsafe.Add(mBase, uint32(v2743+(v1867^int32(-1))*int32(56))+16))
	v2758 = v2749
	goto L447
L449:
	;
	goto L450
L450:
	;
	v2751 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2752 = int32(56)
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(v2751+v1867*v2752-v2752)+16))
	v2758 = v2757
	goto L447
L451:
	;
	v2774 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2773)+16)))
	v2775 = v2774 + v2773
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v2775)+4))
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v2775)))
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v2773)+24))
	v2781 = v2773 + v2778&int32(_a_F_btvacuumscan_7)
	v2782 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2781)+2)))
	v2783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2781))))
	F_UnlockBuffer(m, v1867)
	mBase = m.M
	v2785 = m.ExcPending
	if v2785 != 0 {
		goto L1
	} else {
		goto L455
	}
L452:
	;
	v2763 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(v2763+v2687<<(uint(int32(2))%32))))
	v2773 = v2767
	goto L451
L453:
	;
	goto L454
L454:
	;
	v2769 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2773 = v2769 + v2689 + int32(-8192)
	goto L451
L455:
	;
	v2788 = v2782 | v2783<<(uint(int32(16))%32)
	v2790 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v2790 != 0 {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2792 = m.ExcPending
	if v2792 != 0 {
		goto L1
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	v2794 = int32(1)
	if v2788 == int32(-1) {
		goto L461
	} else {
		goto L462
	}
L459:
	;
	goto L458
L460:
	;
	v2839 = int32(0)
	if v2834 == v2839 {
		v3014 = v2839
		v3016 = int32(0)
		goto L476
	} else {
		goto L477
	}
L461:
	;
	v2834 = v2777
	v2835 = v1867
	v2836 = v2758
	v2837 = v2794
	v2838 = int32(0)
	goto L460
L462:
	;
	goto L463
L463:
	;
	v2798 = F_ReadBuffer(m, v330, v2788)
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	F_LockBufferInternal(m, v2798, int32(1))
	mBase = m.M
	v2802 = m.ExcPending
	if v2802 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F__bt_checkpage(m, v330, v2798)
	mBase = m.M
	v2804 = m.ExcPending
	if v2804 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	if v2798 < int32(0) {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	v2823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2822)+16)))
	v2824 = v2823 + v2822
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v2824)+8))
	v2826 = *(*int32)(unsafe.Add(mBase, uint32(v2824)))
	F_UnlockBuffer(m, v2798)
	mBase = m.M
	v2828 = m.ExcPending
	if v2828 != 0 {
		goto L1
	} else {
		goto L471
	}
L468:
	;
	v2808 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v2808+(v2798^int32(-1))<<(uint(int32(2))%32))))
	v2822 = v2814
	goto L467
L469:
	;
	goto L470
L470:
	;
	v2816 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2822 = v2816 + v2798<<(uint(int32(13))%32) + int32(-8192)
	goto L467
L471:
	;
	if v2758 == v2788 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v2834 = v2826
	v2835 = v2798
	v2836 = v2758
	v2837 = v2794
	v2838 = v2825
	goto L460
L473:
	;
	goto L474
L474:
	;
	F_LockBufferInternal(m, v1867, int32(3))
	mBase = m.M
	v2832 = m.ExcPending
	if v2832 != 0 {
		goto L1
	} else {
		goto L475
	}
L475:
	;
	v2834 = v2826
	v2835 = v2798
	v2836 = v2788
	v2837 = int32(0)
	v2838 = v2825
	goto L460
L476:
	;
	F_LockBufferInternal(m, v2835, int32(3))
	mBase = m.M
	v3062 = m.ExcPending
	if v3062 != 0 {
		goto L1
	} else {
		goto L517
	}
L477:
	;
	v2842 = F_ReadBuffer(m, v330, v2834)
	mBase = m.M
	v2843 = m.ExcPending
	if v2843 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	F_LockBufferInternal(m, v2842, int32(3))
	mBase = m.M
	v2846 = m.ExcPending
	if v2846 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	F__bt_checkpage(m, v330, v2842)
	mBase = m.M
	v2848 = m.ExcPending
	if v2848 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	if v2842 < int32(0) {
		goto L482
	} else {
		goto L483
	}
L481:
	;
	v2867 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2866)+16)))
	v2868 = v2867 + v2866
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2868)+4))
	v2870 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2868)+12)))
	v2872 = v2870 & int32(4)
	if v2872 == int32(0) {
		goto L485
	} else {
		goto L486
	}
L482:
	;
	v2852 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2858 = *(*int32)(unsafe.Add(mBase, uint32(v2852+(v2842^int32(-1))<<(uint(int32(2))%32))))
	v2866 = v2858
	goto L481
L483:
	;
	goto L484
L484:
	;
	v2860 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2866 = v2860 + v2842<<(uint(int32(13))%32) + int32(-8192)
	goto L481
L485:
	;
	if v2869 == v2836 {
		v3014 = v2834
		v3016 = v2842
		goto L476
	} else {
		goto L488
	}
L486:
	;
	goto L487
L487:
	;
	__phi2880 = v2834
	__phi2883 = v2842
	__phi2885 = v2869
	__phi2895 = v2872
	v2880 = __phi2880
	v2883 = __phi2883
	v2885 = __phi2885
	v2895 = __phi2895
	goto L489
L488:
	;
	goto L487
L489:
	;
	v2927 = int32(0)
	if base.B2i32(base.B2i32(v2885 == v2927)|v2895&int32(_a_F_btvacuumscan_5) == v2927)&base.B2i32(v2880 != v2885) == v2927 {
		goto L491
	} else {
		goto L492
	}
L490:
	;
	v3014 = v2885
	v3016 = v2977
	goto L476
L491:
	;
	F_UnlockReleaseBuffer(m, v2883)
	mBase = m.M
	v2939 = m.ExcPending
	if v2939 != 0 {
		goto L1
	} else {
		goto L494
	}
L492:
	;
	goto L493
L493:
	;
	F_UnlockReleaseBuffer(m, v2883)
	mBase = m.M
	v2972 = m.ExcPending
	if v2972 != 0 {
		goto L1
	} else {
		goto L504
	}
L494:
	;
	v2942 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2943 = m.ExcPending
	if v2943 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	if v2942 != 0 {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2946 = m.ExcPending
	if v2946 != 0 {
		goto L1
	} else {
		goto L499
	}
L497:
	;
	goto L498
L498:
	;
	F_ReleaseBuffer(m, v2835)
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L1
	} else {
		goto L502
	}
L499:
	;
	v2947 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+128)) = v2838
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+132)) = v2947 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+124)) = v1857
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+120)) = v2758
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+116)) = v2836
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+112)) = v2885
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_22), v1837+int32(112))
	mBase = m.M
	v2960 = m.ExcPending
	if v2960 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2477), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v2965 = m.ExcPending
	if v2965 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	goto L498
L502:
	;
	if v2837 == int32(0) {
		goto L258
	} else {
		goto L503
	}
L503:
	;
	goto L255
L504:
	;
	v2974 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v2974 != 0 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2976 = m.ExcPending
	if v2976 != 0 {
		goto L1
	} else {
		goto L508
	}
L506:
	;
	goto L507
L507:
	;
	v2977 = F_ReadBuffer(m, v330, v2885)
	mBase = m.M
	v2978 = m.ExcPending
	if v2978 != 0 {
		goto L1
	} else {
		goto L509
	}
L508:
	;
	goto L507
L509:
	;
	F_LockBufferInternal(m, v2977, int32(3))
	mBase = m.M
	v2981 = m.ExcPending
	if v2981 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	F__bt_checkpage(m, v330, v2977)
	mBase = m.M
	v2983 = m.ExcPending
	if v2983 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	if v2977 < int32(0) {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	v3002 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3001)+16)))
	v3003 = v3002 + v3001
	v3004 = *(*int32)(unsafe.Add(mBase, uint32(v3003)+4))
	v3005 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3003)+12)))
	v3007 = v3005 & int32(4)
	if v3007|base.B2i32(v3004 != v2836) != 0 {
		__phi2880 = v2885
		__phi2883 = v2977
		__phi2885 = v3004
		__phi2895 = v3007
		v2880 = __phi2880
		v2883 = __phi2883
		v2885 = __phi2885
		v2895 = __phi2895
		goto L489
	} else {
		goto L516
	}
L513:
	;
	v2987 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2993 = *(*int32)(unsafe.Add(mBase, uint32(v2987+(v2977^int32(-1))<<(uint(int32(2))%32))))
	v3001 = v2993
	goto L512
L514:
	;
	goto L515
L515:
	;
	v2995 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3001 = v2995 + v2977<<(uint(int32(13))%32) + int32(-8192)
	goto L512
L516:
	;
	goto L490
L517:
	;
	v3063 = int32(0)
	v3064 = base.B2i32(v3063 <= v2835)
	if v3064 == v3063 {
		goto L519
	} else {
		goto L520
	}
L518:
	;
	v3083 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3082)+16)))
	v3084 = v3083 + v3082
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v3084)+4))
	if v3085 == int32(0) {
		goto L291
	} else {
		goto L522
	}
L519:
	;
	v3068 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3074 = *(*int32)(unsafe.Add(mBase, uint32(v3068+(v2835^int32(-1))<<(uint(int32(2))%32))))
	v3082 = v3074
	goto L518
L520:
	;
	goto L521
L521:
	;
	v3076 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3082 = v3076 + v2835<<(uint(int32(13))%32) + int32(-8192)
	goto L518
L522:
	;
	v3088 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3084)+12)))
	if v3088&int32(6) != 0 {
		goto L291
	} else {
		goto L523
	}
L523:
	;
	v3091 = *(*int32)(unsafe.Add(mBase, uint32(v3084)))
	if v3091 != v3014 {
		goto L290
	} else {
		goto L524
	}
L524:
	;
	v3093 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3082)+12)))
	v3095 = v3093 + int32(_a_F_btvacuumscan_6)
	if v2837 != 0 {
		goto L526
	} else {
		goto L527
	}
L525:
	;
	v3151 = F_ReadBuffer(m, v330, v3085)
	mBase = m.M
	v3152 = m.ExcPending
	if v3152 != 0 {
		goto L1
	} else {
		goto L540
	}
L526:
	;
	v3096 = int32(17)
	if v3088&v3096 == v3096 {
		goto L529
	} else {
		goto L530
	}
L527:
	;
	goto L528
L528:
	;
	if v3088&int32(1)|base.B2i32(base.Ui32(v3093) < base.Ui32(int32(25)))|base.B2i32(v3095&int32(_a_F_btvacuumscan_24) != int32(8)) != 0 {
		goto L289
	} else {
		goto L536
	}
L529:
	;
	if base.B2i32(v3095&int32(_a_F_btvacuumscan_18) == int32(0))|base.B2i32(base.Ui32(v3093) < base.Ui32(int32(25))) != 0 {
		v3150 = int32(-1)
		goto L525
	} else {
		goto L532
	}
L530:
	;
	goto L531
L531:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3111 = m.ExcPending
	if v3111 != 0 {
		goto L1
	} else {
		goto L533
	}
L532:
	;
	goto L531
L533:
	;
	v3112 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+64)) = v2836
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+68)) = v3112 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_25), v1837-int32(-64))
	mBase = m.M
	v3121 = m.ExcPending
	if v3121 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2524), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L536:
	;
	v3138 = *(*int32)(unsafe.Add(mBase, uint32(v3082)+28))
	v3141 = v3082 + v3138&int32(_a_F_btvacuumscan_7)
	v3142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3141))))
	v3145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3141)+2)))
	v3146 = v3142<<(uint(int32(16))%32) | v3145
	if v2758 == v3146 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v3148 = int32(-1)
	goto L539
L538:
	;
	v3148 = v3146
	goto L539
L539:
	;
	v3150 = v3148
	goto L525
L540:
	;
	F_LockBufferInternal(m, v3151, int32(3))
	mBase = m.M
	v3155 = m.ExcPending
	if v3155 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	F__bt_checkpage(m, v330, v3151)
	mBase = m.M
	v3157 = m.ExcPending
	if v3157 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	v3158 = int32(0)
	v3159 = base.B2i32(v3158 <= v3151)
	if v3159 == v3158 {
		goto L544
	} else {
		goto L545
	}
L543:
	;
	v3178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3177)+16)))
	v3179 = v3178 + v3177
	v3180 = *(*int32)(unsafe.Add(mBase, uint32(v3179)))
	if v2836 != v3180 {
		goto L547
	} else {
		goto L548
	}
L544:
	;
	v3163 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v3163+(v3151^int32(-1))<<(uint(int32(2))%32))))
	v3177 = v3169
	goto L543
L545:
	;
	goto L546
L546:
	;
	v3171 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3177 = v3171 + v3151<<(uint(int32(13))%32) + int32(-8192)
	goto L543
L547:
	;
	v3184 = F_errstart(m, int32(15), int32(0))
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
	v3220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3177)+12)))
	v3221 = int32(0)
	v3224 = *(*int32)(unsafe.Add(mBase, uint32(v3179)+4))
	if v3014|v3224 != 0 {
		v3283 = v3221
		v3284 = v3221
		v3285 = v3221
		goto L564
	} else {
		goto L565
	}
L550:
	;
	if v3184 != 0 {
		goto L551
	} else {
		goto L552
	}
L551:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3188 = m.ExcPending
	if v3188 != 0 {
		goto L1
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	if v3016 != 0 {
		goto L557
	} else {
		goto L558
	}
L554:
	;
	v3189 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	v3190 = *(*int32)(unsafe.Add(mBase, uint32(v3179)))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+52)) = v2838
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+48)) = v3190
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+56)) = v3189 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+44)) = v1857
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+40)) = v2758
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+36)) = v2836
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+32)) = v3085
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_26), v1837+int32(32))
	mBase = m.M
	v3204 = m.ExcPending
	if v3204 != 0 {
		goto L1
	} else {
		goto L555
	}
L555:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2578), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3209 = m.ExcPending
	if v3209 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	goto L553
L557:
	;
	F_UnlockReleaseBuffer(m, v3016)
	mBase = m.M
	v3213 = m.ExcPending
	if v3213 != 0 {
		goto L1
	} else {
		goto L560
	}
L558:
	;
	goto L559
L559:
	;
	F_UnlockReleaseBuffer(m, v3151)
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L1
	} else {
		goto L561
	}
L560:
	;
	goto L559
L561:
	;
	F_UnlockReleaseBuffer(m, v2835)
	mBase = m.M
	v3217 = m.ExcPending
	if v3217 != 0 {
		goto L1
	} else {
		goto L562
	}
L562:
	;
	if v2837 == int32(0) {
		goto L258
	} else {
		goto L563
	}
L563:
	;
	goto L255
L564:
	;
	v3286 = int32(_a_F_btvacuumscan_3)
	v3288 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v3288 + int32(1)
	if v3016 != 0 {
		goto L580
	} else {
		goto L581
	}
L565:
	;
	if v3159 == int32(0) {
		goto L567
	} else {
		goto L568
	}
L566:
	;
	v3244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3243)+16)))
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(v3244+v3243)+4))
	if v3246 != 0 {
		v3283 = v3221
		v3284 = v3221
		v3285 = v3221
		goto L564
	} else {
		goto L570
	}
L567:
	;
	v3229 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3235 = *(*int32)(unsafe.Add(mBase, uint32(v3229+(v3151^int32(-1))<<(uint(int32(2))%32))))
	v3243 = v3235
	goto L566
L568:
	;
	goto L569
L569:
	;
	v3237 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3243 = v3237 + v3151<<(uint(int32(13))%32) + int32(-8192)
	goto L566
L570:
	;
	v3248 = F_ReadBuffer(m, v330, int32(0))
	mBase = m.M
	v3249 = m.ExcPending
	if v3249 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	F_LockBufferInternal(m, v3248, int32(3))
	mBase = m.M
	v3252 = m.ExcPending
	if v3252 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	F__bt_checkpage(m, v330, v3248)
	mBase = m.M
	v3254 = m.ExcPending
	if v3254 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	if v3248 < int32(0) {
		goto L575
	} else {
		goto L576
	}
L574:
	;
	v3274 = v3272 + int32(24)
	v3275 = *(*int32)(unsafe.Add(mBase, uint32(v3272)+44))
	if base.Ui32(v3275) <= base.Ui32(v2838+int32(1)) {
		v3283 = v3272
		v3284 = v3274
		v3285 = v3248
		goto L564
	} else {
		goto L578
	}
L575:
	;
	v3258 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3264 = *(*int32)(unsafe.Add(mBase, uint32(v3258+(v3248^int32(-1))<<(uint(int32(2))%32))))
	v3272 = v3264
	goto L574
L576:
	;
	goto L577
L577:
	;
	v3266 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3272 = v3266 + v3248<<(uint(int32(13))%32) + int32(-8192)
	goto L574
L578:
	;
	F_UnlockReleaseBuffer(m, v3248)
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	v3283 = v3272
	v3284 = v3274
	v3285 = int32(0)
	goto L564
L580:
	;
	if v3016 < int32(0) {
		goto L584
	} else {
		goto L585
	}
L581:
	;
	goto L582
L582:
	;
	if v3159 == int32(0) {
		goto L588
	} else {
		goto L589
	}
L583:
	;
	v3310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3309)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v3310+v3309)+4)) = v3085
	goto L582
L584:
	;
	v3295 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v3295+(v3016^int32(-1))<<(uint(int32(2))%32))))
	v3309 = v3301
	goto L583
L585:
	;
	goto L586
L586:
	;
	v3303 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3309 = v3303 + v3016<<(uint(int32(13))%32) + int32(-8192)
	goto L583
L587:
	;
	v3332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3331)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v3332+v3331))) = v3014
	if v2837 == int32(0) {
		goto L591
	} else {
		goto L592
	}
L588:
	;
	v3317 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3323 = *(*int32)(unsafe.Add(mBase, uint32(v3317+(v3151^int32(-1))<<(uint(int32(2))%32))))
	v3331 = v3323
	goto L587
L589:
	;
	goto L590
L590:
	;
	v3325 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3331 = v3325 + v3151<<(uint(int32(13))%32) + int32(-8192)
	goto L587
L591:
	;
	v3337 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2781)+4)) = uint16(v3337)
	*(*int32)(unsafe.Add(mBase, uint32(v2781))) = base.I32_rotr(v3150, int32(16))
	v3342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2781)+6)))
	v3344 = v3342 | int32(_a_F_btvacuumscan_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2781)+6)) = uint16(v3344)
	goto L593
L592:
	;
	goto L593
L593:
	;
	if v3064 == int32(0) {
		goto L595
	} else {
		goto L596
	}
L594:
	;
	v3364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3363)+16)))
	v3365 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v3366 = m.ExcPending
	if v3366 != 0 {
		goto L1
	} else {
		goto L598
	}
L595:
	;
	v3349 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v3349+(v2835^int32(-1))<<(uint(int32(2))%32))))
	v3363 = v3355
	goto L594
L596:
	;
	goto L597
L597:
	;
	v3357 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3363 = v3357 + v2835<<(uint(int32(13))%32) + int32(-8192)
	goto L594
L598:
	;
	v3367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3363)+16)))
	v3368 = v3363 + v3367
	v3369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3368)+12)))
	v3373 = v3369&int32(_a_F_btvacuumscan_27) | int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v3368)+12)) = uint16(v3373)
	v3375 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v3363)+12)) = uint16(v3375)
	*(*int64)(unsafe.Add(mBase, uint32(v3363)+24)) = v3365
	v3378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3363)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3363)+14)) = uint16(v3378)
	v3381 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3364+v3363)+14)) = uint16(v3381)
	if v3285 != 0 {
		goto L599
	} else {
		goto L600
	}
L599:
	;
	v3383 = *(*int32)(unsafe.Add(mBase, uint32(v3284)+4))
	if base.Ui32(v3383) <= base.Ui32(int32(2)) {
		goto L602
	} else {
		goto L603
	}
L600:
	;
	goto L601
L601:
	;
	F_MarkBufferDirty(m, v3151)
	mBase = m.M
	v3401 = m.ExcPending
	if v3401 != 0 {
		goto L1
	} else {
		goto L606
	}
L602:
	;
	v3386 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3283)+64)) = uint8(v3386)
	*(*int64)(unsafe.Add(mBase, uint32(v3283)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v3283)+48)) = v3386
	*(*int32)(unsafe.Add(mBase, uint32(v3283)+28)) = int32(3)
	v3394 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v3283)+12)) = uint16(v3394)
	goto L604
L603:
	;
	goto L604
L604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3284)+20)) = v2838
	*(*int32)(unsafe.Add(mBase, uint32(v3284)+16)) = v3085
	F_MarkBufferDirty(m, v3285)
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L1
	} else {
		goto L605
	}
L605:
	;
	goto L601
L606:
	;
	F_MarkBufferDirty(m, v2835)
	mBase = m.M
	v3403 = m.ExcPending
	if v3403 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	if v3016 != 0 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	F_MarkBufferDirty(m, v3016)
	mBase = m.M
	v3405 = m.ExcPending
	if v3405 != 0 {
		goto L1
	} else {
		goto L611
	}
L609:
	;
	goto L610
L610:
	;
	if v2837 == int32(0) {
		goto L612
	} else {
		goto L613
	}
L611:
	;
	goto L610
L612:
	;
	F_MarkBufferDirty(m, v1867)
	mBase = m.M
	v3409 = m.ExcPending
	if v3409 != 0 {
		goto L1
	} else {
		goto L615
	}
L613:
	;
	goto L614
L614:
	;
	v3410 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	v3411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3410)+118)))
	if v3411 != int32(112) {
		goto L617
	} else {
		goto L618
	}
L615:
	;
	goto L614
L616:
	;
	if v3285 != 0 {
		goto L644
	} else {
		goto L645
	}
L617:
	;
	v3486 = F_XLogGetFakeLSN(m, v330)
	mBase = m.M
	v3487 = m.ExcPending
	if v3487 != 0 {
		goto L1
	} else {
		goto L643
	}
L618:
	;
	v3415 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if v3415 <= int32(0) {
		goto L619
	} else {
		goto L620
	}
L619:
	;
	v3418 = *(*int32)(unsafe.Add(mBase, uint32(v330)+32))
	if v3418 != 0 {
		goto L617
	} else {
		goto L622
	}
L620:
	;
	goto L621
L621:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v3421 = m.ExcPending
	if v3421 != 0 {
		goto L1
	} else {
		goto L624
	}
L622:
	;
	v3419 = *(*int32)(unsafe.Add(mBase, uint32(v330)+40))
	if v3419 != 0 {
		goto L617
	} else {
		goto L623
	}
L623:
	;
	goto L621
L624:
	;
	F_XLogRegisterBuffer(m, int32(0), v2835, int32(6))
	mBase = m.M
	v3425 = m.ExcPending
	if v3425 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	if v3016 != 0 {
		goto L626
	} else {
		goto L627
	}
L626:
	;
	F_XLogRegisterBuffer(m, int32(1), v3016, int32(8))
	mBase = m.M
	v3429 = m.ExcPending
	if v3429 != 0 {
		goto L1
	} else {
		goto L629
	}
L627:
	;
	goto L628
L628:
	;
	F_XLogRegisterBuffer(m, int32(2), v3151, int32(8))
	mBase = m.M
	v3433 = m.ExcPending
	if v3433 != 0 {
		goto L1
	} else {
		goto L630
	}
L629:
	;
	goto L628
L630:
	;
	if v2837 == int32(0) {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	F_XLogRegisterBuffer(m, int32(3), v1867, int32(6))
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L1
	} else {
		goto L634
	}
L632:
	;
	goto L633
L633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+280)) = v3150
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+276)) = v2776
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+272)) = v2777
	*(*int64)(unsafe.Add(mBase, uint32(v1837)+264)) = v3365
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+256)) = v2838
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+252)) = v3085
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+248)) = v3014
	F_XLogRegisterData(m, v1837+int32(248), int32(36))
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		goto L1
	} else {
		goto L635
	}
L634:
	;
	goto L633
L635:
	;
	if v3285 == int32(0) {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	v3456 = F_XLogInsert(m, int32(11), int32(128))
	mBase = m.M
	v3457 = m.ExcPending
	if v3457 != 0 {
		goto L1
	} else {
		goto L639
	}
L637:
	;
	goto L638
L638:
	;
	F_XLogRegisterBuffer(m, int32(4), v3285, int32(14))
	mBase = m.M
	v3461 = m.ExcPending
	if v3461 != 0 {
		goto L1
	} else {
		goto L640
	}
L639:
	;
	v3488 = v3456
	goto L616
L640:
	;
	v3462 = *(*int32)(unsafe.Add(mBase, uint32(v3284)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+220)) = v3462
	v3464 = *(*int32)(unsafe.Add(mBase, uint32(v3284)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+224)) = v3464
	v3466 = *(*int32)(unsafe.Add(mBase, uint32(v3284)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+228)) = v3466
	v3468 = *(*int32)(unsafe.Add(mBase, uint32(v3284)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+232)) = v3468
	v3470 = *(*int32)(unsafe.Add(mBase, uint32(v3284)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+236)) = v3470
	v3472 = *(*int32)(unsafe.Add(mBase, uint32(v3284)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+240)) = v3472
	v3474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3284)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1837)+244)) = uint8(v3474)
	F_XLogRegisterBufData(m, int32(4), v1837+int32(220), int32(28))
	mBase = m.M
	v3481 = m.ExcPending
	if v3481 != 0 {
		goto L1
	} else {
		goto L641
	}
L641:
	;
	v3484 = F_XLogInsert(m, int32(11), int32(144))
	mBase = m.M
	v3485 = m.ExcPending
	if v3485 != 0 {
		goto L1
	} else {
		goto L642
	}
L642:
	;
	v3488 = v3484
	goto L616
L643:
	;
	v3488 = v3486
	goto L616
L644:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3283))) = base.I64_rotl(v3488, int64(32))
	goto L646
L645:
	;
	goto L646
L646:
	;
	if v3159 == int32(0) {
		goto L648
	} else {
		goto L649
	}
L647:
	;
	v3511 = base.I64_rotl(v3488, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v3509))) = v3511
	if v3064 == int32(0) {
		goto L652
	} else {
		goto L653
	}
L648:
	;
	v3495 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3501 = *(*int32)(unsafe.Add(mBase, uint32(v3495+(v3151^int32(-1))<<(uint(int32(2))%32))))
	v3509 = v3501
	goto L647
L649:
	;
	goto L650
L650:
	;
	v3503 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3509 = v3503 + v3151<<(uint(int32(13))%32) + int32(-8192)
	goto L647
L651:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3530))) = v3511
	if v3016 != 0 {
		goto L655
	} else {
		goto L656
	}
L652:
	;
	v3516 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3522 = *(*int32)(unsafe.Add(mBase, uint32(v3516+(v2835^int32(-1))<<(uint(int32(2))%32))))
	v3530 = v3522
	goto L651
L653:
	;
	goto L654
L654:
	;
	v3524 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3530 = v3524 + v2835<<(uint(int32(13))%32) + int32(-8192)
	goto L651
L655:
	;
	if v3016 < int32(0) {
		goto L659
	} else {
		goto L660
	}
L656:
	;
	goto L657
L657:
	;
	if v2837 == int32(0) {
		goto L662
	} else {
		goto L663
	}
L658:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3549))) = v3511
	goto L657
L659:
	;
	v3535 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3541 = *(*int32)(unsafe.Add(mBase, uint32(v3535+(v3016^int32(-1))<<(uint(int32(2))%32))))
	v3549 = v3541
	goto L658
L660:
	;
	goto L661
L661:
	;
	v3543 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3549 = v3543 + v3016<<(uint(int32(13))%32) + int32(-8192)
	goto L658
L662:
	;
	if v1909 == int32(0) {
		goto L666
	} else {
		goto L667
	}
L663:
	;
	goto L664
L664:
	;
	v3568 = int32(_a_F_btvacuumscan_3)
	v3570 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v3570 - int32(1)
	if v3285 != 0 {
		goto L669
	} else {
		goto L670
	}
L665:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3566))) = v3511
	goto L664
L666:
	;
	v3556 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3560 = *(*int32)(unsafe.Add(mBase, uint32(v3556+v2687<<(uint(int32(2))%32))))
	v3566 = v3560
	goto L665
L667:
	;
	goto L668
L668:
	;
	v3562 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3566 = v3562 + v2689 + int32(-8192)
	goto L665
L669:
	;
	F_UnlockReleaseBuffer(m, v3285)
	mBase = m.M
	v3575 = m.ExcPending
	if v3575 != 0 {
		goto L1
	} else {
		goto L672
	}
L670:
	;
	goto L671
L671:
	;
	if v3016 != 0 {
		goto L673
	} else {
		goto L674
	}
L672:
	;
	goto L671
L673:
	;
	F_UnlockReleaseBuffer(m, v3016)
	mBase = m.M
	v3577 = m.ExcPending
	if v3577 != 0 {
		goto L1
	} else {
		goto L676
	}
L674:
	;
	goto L675
L675:
	;
	F_UnlockReleaseBuffer(m, v3151)
	mBase = m.M
	v3579 = m.ExcPending
	if v3579 != 0 {
		goto L1
	} else {
		goto L677
	}
L676:
	;
	goto L675
L677:
	;
	if v2837 == int32(0) {
		goto L678
	} else {
		goto L679
	}
L678:
	;
	F_UnlockReleaseBuffer(m, v2835)
	mBase = m.M
	v3583 = m.ExcPending
	if v3583 != 0 {
		goto L1
	} else {
		goto L681
	}
L679:
	;
	goto L680
L680:
	;
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(v2759)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2759)+24)) = v3584 + int32(1)
	if base.Ui32(v2836) <= base.Ui32(v1857) {
		goto L682
	} else {
		goto L683
	}
L681:
	;
	goto L680
L682:
	;
	v3589 = *(*int32)(unsafe.Add(mBase, uint32(v2759)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2759)+28)) = v3589 + int32(1)
	goto L684
L683:
	;
	goto L684
L684:
	;
	v3593 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+36))
	v3594 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+28))
	if v3593 != v3594 {
		goto L685
	} else {
		goto L686
	}
L685:
	;
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+24))
	if v3596 != v3593 {
		goto L689
	} else {
		goto L690
	}
L686:
	;
	goto L687
L687:
	;
	v3631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1929)+12)))
	if v3631&int32(16) != 0 {
		goto L445
	} else {
		goto L696
	}
L688:
	;
	v3614 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v3612+v3613<<(uint(v3614)%32)))) = v2836
	v3618 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+32))
	v3619 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v3618+v3619<<(uint(v3614)%32))+8)) = v3365
	v3624 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+36)) = v3624 + int32(1)
	goto L687
L689:
	;
	v3598 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+32))
	v3612 = v3598
	v3613 = v3593
	goto L688
L690:
	;
	goto L691
L691:
	;
	v3600 = v3593 << (uint(int32(1)) % 32)
	if v3600 < v3594 {
		goto L692
	} else {
		goto L693
	}
L692:
	;
	v3602 = v3600
	goto L694
L693:
	;
	v3602 = v3594
	goto L694
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+24)) = v3602
	v3604 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+32))
	v3607 = F_repalloc(m, v3604, v3602<<(uint(int32(4))%32))
	mBase = m.M
	v3608 = m.ExcPending
	if v3608 != 0 {
		goto L1
	} else {
		goto L695
	}
L695:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+32)) = v3607
	v3610 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+36))
	v3612 = v3607
	v3613 = v3610
	goto L688
L696:
	;
	goto L446
L697:
	;
	v3642 = v3636
	goto L699
L698:
	;
	v3642 = int32(1)
	goto L699
L699:
	;
	v3651 = base.B2i32(base.Ui32(int32(base.Ui32(v3220+int32(_a_F_btvacuumscan_6))>>(uint(v3636)%32))&int32(_a_F_btvacuumscan_5)) < base.Ui32(v3642)) | base.B2i32(base.Ui32(v3220) < base.Ui32(int32(25)))
	goto L444
L700:
	;
	v3701 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v3701 != 0 {
		goto L701
	} else {
		goto L702
	}
L701:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3703 = m.ExcPending
	if v3703 != 0 {
		goto L1
	} else {
		goto L704
	}
L702:
	;
	goto L703
L703:
	;
	if v3651 == int32(0) {
		goto L255
	} else {
		goto L705
	}
L704:
	;
	goto L703
L705:
	;
	v3706 = F_ReadBuffer(m, v330, v3697)
	mBase = m.M
	v3707 = m.ExcPending
	if v3707 != 0 {
		goto L1
	} else {
		goto L706
	}
L706:
	;
	F_LockBufferInternal(m, v3706, int32(3))
	mBase = m.M
	v3710 = m.ExcPending
	if v3710 != 0 {
		goto L1
	} else {
		goto L707
	}
L707:
	;
	F__bt_checkpage(m, v330, v3706)
	mBase = m.M
	v3712 = m.ExcPending
	if v3712 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	v1867 = v3706
	goto L256
L709:
	;
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_28), int32(0))
	mBase = m.M
	v3720 = m.ExcPending
	if v3720 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2278), int32(_a_F_btvacuumscan_20))
	mBase = m.M
	v3725 = m.ExcPending
	if v3725 != 0 {
		goto L1
	} else {
		goto L711
	}
L711:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L712:
	;
	v3731 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+16)) = v2836
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+20)) = v3731 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_29), v1837+int32(16))
	mBase = m.M
	v3740 = m.ExcPending
	if v3740 != 0 {
		goto L1
	} else {
		goto L713
	}
L713:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2510), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3745 = m.ExcPending
	if v3745 != 0 {
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
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3752 = m.ExcPending
	if v3752 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	v3753 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	v3754 = *(*int32)(unsafe.Add(mBase, uint32(v3084)))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+104)) = v2836
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+100)) = v3754
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+96)) = v3014
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+108)) = v3753 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_30), v1837+int32(96))
	mBase = m.M
	v3765 = m.ExcPending
	if v3765 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2517), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3770 = m.ExcPending
	if v3770 != 0 {
		goto L1
	} else {
		goto L718
	}
L718:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L719:
	;
	v3775 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+84)) = v2836
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+80)) = v2838
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+88)) = v3775 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_31), v1837+int32(80))
	mBase = m.M
	v3785 = m.ExcPending
	if v3785 != 0 {
		goto L1
	} else {
		goto L720
	}
L720:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2536), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3790 = m.ExcPending
	if v3790 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L722:
	;
	if v3843 == int32(0) {
		goto L258
	} else {
		goto L723
	}
L723:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3849 = m.ExcPending
	if v3849 != 0 {
		goto L1
	} else {
		goto L724
	}
L724:
	;
	v3850 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+4)) = v3799
	*(*int32)(unsafe.Add(mBase, uint32(v1837))) = v3850 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_32), v1837)
	mBase = m.M
	v3857 = m.ExcPending
	if v3857 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2885), int32(_a_F_btvacuumscan_33))
	mBase = m.M
	v3862 = m.ExcPending
	if v3862 != 0 {
		goto L1
	} else {
		goto L726
	}
L726:
	;
	goto L258
L727:
	;
	v3867 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v3865)+3)) = uint16(v3867)
	v3872 = int32(1)
	v3874 = F__bt_search(m, v330, int32(0), v3865, v1837+int32(248), v3872, v3872)
	mBase = m.M
	v3875 = m.ExcPending
	if v3875 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	v3876 = *(*int32)(unsafe.Add(mBase, uint32(v1837)+248))
	F_UnlockReleaseBuffer(m, v3876)
	mBase = m.M
	v3878 = m.ExcPending
	if v3878 != 0 {
		goto L1
	} else {
		goto L729
	}
L729:
	;
	F_LockBufferInternal(m, v1867, int32(3))
	mBase = m.M
	v3881 = m.ExcPending
	if v3881 != 0 {
		goto L1
	} else {
		goto L730
	}
L730:
	;
	v1877 = v3874
	goto L256
L731:
	;
	goto L255
L732:
	;
	v4079 = v4027
	goto L54
L733:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v4095 = m.ExcPending
	if v4095 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	v4096 = int32(0)
	v4098 = *(*int32)(unsafe.Add(mBase, uint32(v366)+24))
	v4099 = F_ReadBufferExtended(m, v330, v4096, v4079, v4096, v4098)
	mBase = m.M
	v4100 = m.ExcPending
	if v4100 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	v326 = v4099
	v332 = v4079
	goto L51
L736:
	;
	v126 = v242
	v127 = v243
	v137 = v253
	v158 = v274
	v166 = v282
	v168 = v284
	goto L17
L737:
	;
	if v4107 == int32(0) {
		goto L41
	} else {
		goto L738
	}
L738:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v4113 = m.ExcPending
	if v4113 != 0 {
		goto L1
	} else {
		goto L739
	}
L739:
	;
	v4114 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+4)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v334))) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v334)+8)) = v4114 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_34), v334)
	mBase = m.M
	v4122 = m.ExcPending
	if v4122 != 0 {
		goto L1
	} else {
		goto L740
	}
L740:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_35), int32(1470), int32(_a_F_btvacuumscan_36))
	mBase = m.M
	v4127 = m.ExcPending
	if v4127 != 0 {
		goto L1
	} else {
		goto L741
	}
L741:
	;
	goto L41
L742:
	;
	v4132 = v323
	v4133 = v324
	v4143 = v334
	v4164 = v355
	v4172 = v363
	v4174 = v365
	goto L40
L743:
	;
	v4189 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[2]))
	if v4189 == int32(0) {
		goto L745
	} else {
		goto L746
	}
L744:
	;
	v242 = v4132
	v243 = v4133
	v253 = v4143
	v274 = v4164
	v282 = v4172
	v284 = v4174
	goto L37
L745:
	;
	goto L744
L746:
	;
	v4193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btvacuumscan[3])))
	if v4193&int32(1) == int32(0) {
		goto L745
	} else {
		goto L747
	}
L747:
	;
	v4198 = int32(_a_F_btvacuumscan_3)
	v4200 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	v4201 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v4200 + v4201
	v4204 = *(*int32)(unsafe.Add(mBase, uint32(v4189)))
	*(*int32)(unsafe.Add(mBase, uint32(v4189))) = v4204 + v4201
	v4208 = int32(0)
	v4210 = int32(_a_F_btvacuumscan_4)
	v4211 = base.AtomicRmwOr32(m, v4208, v4210, v4208)
	*(*int64)(unsafe.Add(mBase, uint32(v4189+int32(128))+232)) = base.I64_extend_i32_u(v358)
	v4219 = base.AtomicRmwOr32(m, v4208, v4210, v4208)
	v4220 = *(*int32)(unsafe.Add(mBase, uint32(v4189)))
	*(*int32)(unsafe.Add(mBase, uint32(v4189))) = v4220 + v4201
	v4226 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v4226 - v4201
	goto L745
L748:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v190
	v4233 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	F_MemoryContextDelete(m, v4233)
	mBase = m.M
	v4235 = m.ExcPending
	if v4235 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	v4236 = int32(0)
	v4238 = v137 + int32(24)
	v4239 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+36))
	if v4239 == v4236 {
		goto L752
	} else {
		goto L753
	}
L750:
	;
	v4474 = *(*int32)(unsafe.Add(mBase, uint32(v127)+32))
	if v4474 != 0 {
		goto L766
	} else {
		goto L767
	}
L751:
	;
	F_pfree(m, v4374)
	mBase = m.M
	v4423 = m.ExcPending
	if v4423 != 0 {
		goto L1
	} else {
		goto L765
	}
L752:
	;
	v4242 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+32))
	if v4242 != 0 {
		v4374 = v4242
		goto L751
	} else {
		goto L755
	}
L753:
	;
	goto L754
L754:
	;
	v4243 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+4))
	v4244 = *(*int32)(unsafe.Add(mBase, uint32(v4238)))
	v4245 = *(*int32)(unsafe.Add(mBase, uint32(v4244)+4))
	v4246 = F_GetOldestNonRemovableTransactionId(m, v4245)
	mBase = m.M
	v4247 = m.ExcPending
	if v4247 != 0 {
		goto L1
	} else {
		goto L756
	}
L755:
	;
	goto L750
L756:
	;
	v4248 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+36))
	if v4248 <= int32(0) {
		goto L757
	} else {
		goto L758
	}
L757:
	;
	v4371 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+32))
	v4374 = v4371
	goto L751
L758:
	;
	v4253 = v4236
	goto L759
L759:
	;
	v4301 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+32))
	v4304 = v4301 + v4253<<(uint(int32(4))%32)
	v4305 = *(*int32)(unsafe.Add(mBase, uint32(v4304)))
	v4306 = *(*int64)(unsafe.Add(mBase, uint32(v4304)+8))
	v4307 = F_GlobalVisCheckRemovableFullXid(m, v4245, v4306)
	mBase = m.M
	v4308 = m.ExcPending
	if v4308 != 0 {
		goto L1
	} else {
		goto L761
	}
L760:
	;
	goto L757
L761:
	;
	if v4307 == int32(0) {
		goto L757
	} else {
		goto L762
	}
L762:
	;
	F_RecordFreeIndexPage(m, v158, v4305)
	mBase = m.M
	v4312 = m.ExcPending
	if v4312 != 0 {
		goto L1
	} else {
		goto L763
	}
L763:
	;
	v4313 = *(*int32)(unsafe.Add(mBase, uint32(v4243)+32))
	v4314 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4243)+32)) = v4313 + v4314
	v4318 = v4253 + v4314
	v4319 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+36))
	if v4318 < v4319 {
		v4253 = v4318
		goto L759
	} else {
		goto L764
	}
L764:
	;
	goto L760
L765:
	;
	goto L750
L766:
	;
	F_FreeSpaceMapVacuum(m, v158)
	mBase = m.M
	v4476 = m.ExcPending
	if v4476 != 0 {
		goto L1
	} else {
		goto L769
	}
L767:
	;
	goto L768
L768:
	;
	m.G0 = v137 + int32(2512)
	return
L769:
	;
	goto L768
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
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
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
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
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
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
	v30 = F_palloc0_mul(m, int32(2), v22)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v30
	if int32(0) < v22 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L43
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L36
	}
L6:
	;
	v42 = int32(0)
	v45 = v30
	v47 = int32(-1)
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v58 = l1 + v52<<(uint(int32(3))%32) + v42*int32(100)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+119)))
	if v59 != 0 {
		v172 = v45
		v174 = v47
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v180 = v42 + int32(1)
	if v180 != v22 {
		v42 = v180
		v45 = v172
		v47 = v174
		goto L9
	} else {
		goto L35
	}
L12:
	;
	v61 = v58 + int32(28)
	v63 = v58 + int32(32)
	if v21 <= int32(0) {
		v150 = v45
		v152 = v47
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if l2 != 0 {
		v172 = v150
		v174 = v152
		goto L11
	} else {
		goto L33
	}
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+76))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+68))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = v47
	v87 = int32(0)
	goto L15
L15:
	;
	v92 = v86 + int32(1)
	if v92 < v21 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v150 = v45
	v152 = v95
	goto L13
L17:
	;
	v139 = v87 + int32(1)
	if v139 != v21 {
		v86 = v95
		v87 = v139
		goto L15
	} else {
		goto L32
	}
L18:
	;
	v95 = v92
	goto L20
L19:
	;
	v95 = int32(0)
	goto L20
L20:
	;
	v98 = l0 + v68<<(uint(int32(3))%32) + int32(28) + v95*int32(100)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+91)))
	if v99 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v101 = v98 + int32(4)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if base.B2i32(v104 == int32(0))|base.B2i32(v104 != v107) != 0 {
		v125 = v104
		v126 = v107
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v125-v126 != 0 {
		goto L17
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v110 = v63
	v111 = v101
	goto L25
L25:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v115 == int32(0) {
		v125 = v115
		v126 = v114
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v125 = v115
	v126 = v114
	goto L23
L27:
	;
	v118 = int32(1)
	if v115 == v114 {
		v110 = v110 + v118
		v111 = v111 + v118
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
	if v67 != v128 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
	if v66 != v130 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v45+v42<<(uint(int32(1))%32)))) = uint16(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v150 = v137
	v152 = v95
	goto L13
L32:
	;
	goto L16
L33:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150+v42<<(uint(int32(1))%32)))))
	if v160 == int32(0) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v172 = v150
	v174 = v152
	goto L11
L35:
	;
	goto L10
L36:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_build_attrmap_by_name_0), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v214 = F_format_type_be(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v217 = F_format_type_be(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v63
	v225 = F_errdetail(m, int32(_a_F_build_attrmap_by_name_1), v19+int32(16))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_build_attrmap_by_name_2), int32(235), int32(_a_F_build_attrmap_by_name_3))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
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
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_build_attrmap_by_name_0), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v244 = F_format_type_be(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v247 = F_format_type_be(m, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v63
	v253 = F_errdetail(m, int32(_a_F_build_attrmap_by_name_4), v19)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_build_attrmap_by_name_2), int32(247), int32(_a_F_build_attrmap_by_name_3))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
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
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v57 int32
	_ = v57
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
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
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
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
	var v462 int32
	_ = v462
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
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
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v636 int32
	_ = v636
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
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
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v699 int32
	_ = v699
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v734 int32
	_ = v734
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v818 int32
	_ = v818
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v874 int32
	_ = v874
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
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
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1016 int32
	_ = v1016
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1089 int32
	_ = v1089
	var v1117 int32
	_ = v1117
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	v7 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	if v25&int32(2) == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
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
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+256))
	if v30 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)+256))
	if v33 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+233)))
	if base.B2i32(v36 != int32(1))|base.B2i32(v33 != v30) != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+233)))
	if v41&int32(1) == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v47 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v47
	if l5 == int32(0) {
		v622 = v7
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l2)+256))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+256)) = v840
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v844 = int32(*(*int16)(unsafe.Add(mBase, uint32(v840)+2)))
	v845 = F_palloc0_mul(m, int32(4), v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L102
	} else {
		goto L200
	}
L9:
	;
	v628 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+2)))
	if v628 <= int32(0) {
		goto L2
	} else {
		goto L166
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v57 <= int32(0) {
		v622 = v7
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v77 = v7
	v78 = v7
	goto L12
L12:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v77<<(uint(int32(2))%32))))
	if int32(1)<<(uint(v46)%32)&int32(174) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v622 = v603
	goto L9
L14:
	;
	v605 = v77 + int32(1)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v605 < v606 {
		v77 = v605
		v78 = v603
		goto L12
	} else {
		goto L165
	}
L15:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+8)))
	if v89 != 0 {
		v603 = v78
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+9)))
	if v148 != int32(1) {
		v603 = v78
		goto L14
	} else {
		goto L34
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v88)+32))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v92 = int32(0)
	if v90 == v92 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v145 == int32(0) {
		v603 = v78
		goto L14
	} else {
		goto L33
	}
L20:
	;
	v145 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	if v91 == int32(0) {
		v138 = v92
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v145 = v138
	goto L19
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v102 < v101 {
		v138 = v92
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v104 = int32(1)
	if v101 <= v104 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v107 = v104
	goto L28
L27:
	;
	v107 = v101
	goto L28
L28:
	;
	v108 = int32(8)
	v113 = int32(0)
	goto L29
L29:
	;
	v120 = v113 << (uint(int32(2)) % 32)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v90+v108+v120)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v91+v108+v120)))
	v127 = v122 & (v124 ^ int32(-1))
	v129 = base.B2i32(v127 == int32(0))
	if v127 != 0 {
		v138 = v129
		goto L23
	} else {
		goto L31
	}
L30:
	;
	v138 = v129
	goto L23
L31:
	;
	v131 = v113 + int32(1)
	if v131 != v107 {
		v113 = v131
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
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v88)+96))
	if v151 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v88)+124))
	if v154 == int32(0) {
		v603 = v78
		goto L14
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v88)+44))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v160 = int32(0)
	if v158 == v160 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	goto L37
L39:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v403 = F_op_strict(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L102
	} else {
		goto L103
	}
L40:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v88)+44))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v280 = int32(0)
	if v278 == v280 {
		goto L72
	} else {
		goto L73
	}
L41:
	;
	if v213 == int32(0) {
		goto L40
	} else {
		goto L55
	}
L42:
	;
	v213 = int32(1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	if v159 == int32(0) {
		v206 = v160
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v213 = v206
	goto L41
L46:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v170 < v169 {
		v206 = v160
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v172 = int32(1)
	if v169 <= v172 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v175 = v172
	goto L50
L49:
	;
	v175 = v169
	goto L50
L50:
	;
	v176 = int32(8)
	v181 = int32(0)
	goto L51
L51:
	;
	v188 = v181 << (uint(int32(2)) % 32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v158+v176+v188)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v159+v176+v188)))
	v195 = v190 & (v192 ^ int32(-1))
	v197 = base.B2i32(v195 == int32(0))
	if v195 != 0 {
		v206 = v197
		goto L45
	} else {
		goto L53
	}
L52:
	;
	v206 = v197
	goto L45
L53:
	;
	v199 = v181 + int32(1)
	if v199 != v175 {
		v181 = v199
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v88)+48))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v218 = int32(0)
	if v216 == v218 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v271 == int32(0) {
		goto L40
	} else {
		goto L70
	}
L57:
	;
	v271 = int32(1)
	goto L56
L58:
	;
	goto L59
L59:
	;
	if v217 == int32(0) {
		v264 = v218
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v271 = v264
	goto L56
L61:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v228 < v227 {
		v264 = v218
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v230 = int32(1)
	if v227 <= v230 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v233 = v230
	goto L65
L64:
	;
	v233 = v227
	goto L65
L65:
	;
	v234 = int32(8)
	v239 = int32(0)
	goto L66
L66:
	;
	v246 = v239 << (uint(int32(2)) % 32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v216+v234+v246)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v217+v234+v246)))
	v253 = v248 & (v250 ^ int32(-1))
	v255 = base.B2i32(v253 == int32(0))
	if v253 != 0 {
		v264 = v255
		goto L60
	} else {
		goto L68
	}
L67:
	;
	v264 = v255
	goto L60
L68:
	;
	v257 = v239 + int32(1)
	if v257 != v233 {
		v239 = v257
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v157)+28))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v398 = v275 + int32(4)
	v399 = v275
	goto L39
L71:
	;
	if v333 == int32(0) {
		v603 = v78
		goto L14
	} else {
		goto L85
	}
L72:
	;
	v333 = int32(1)
	goto L71
L73:
	;
	goto L74
L74:
	;
	if v279 == int32(0) {
		v326 = v280
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v333 = v326
	goto L71
L76:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v290 < v289 {
		v326 = v280
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v292 = int32(1)
	if v289 <= v292 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v295 = v292
	goto L80
L79:
	;
	v295 = v289
	goto L80
L80:
	;
	v296 = int32(8)
	v301 = int32(0)
	goto L81
L81:
	;
	v308 = v301 << (uint(int32(2)) % 32)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v278+v296+v308)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v279+v296+v308)))
	v315 = v310 & (v312 ^ int32(-1))
	v317 = base.B2i32(v315 == int32(0))
	if v315 != 0 {
		v326 = v317
		goto L75
	} else {
		goto L83
	}
L82:
	;
	v326 = v317
	goto L75
L83:
	;
	v319 = v301 + int32(1)
	if v319 != v295 {
		v301 = v319
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v88)+48))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v338 = int32(0)
	if v336 == v338 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v391 == int32(0) {
		v603 = v78
		goto L14
	} else {
		goto L100
	}
L87:
	;
	v391 = int32(1)
	goto L86
L88:
	;
	goto L89
L89:
	;
	if v337 == int32(0) {
		v384 = v338
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v391 = v384
	goto L86
L91:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v348 < v347 {
		v384 = v338
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v350 = int32(1)
	if v347 <= v350 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v353 = v350
	goto L95
L94:
	;
	v353 = v347
	goto L95
L95:
	;
	v354 = int32(8)
	v359 = int32(0)
	goto L96
L96:
	;
	v366 = v359 << (uint(int32(2)) % 32)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v336+v354+v366)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v337+v354+v366)))
	v373 = v368 & (v370 ^ int32(-1))
	v375 = base.B2i32(v373 == int32(0))
	if v373 != 0 {
		v384 = v375
		goto L90
	} else {
		goto L98
	}
L97:
	;
	v384 = v375
	goto L90
L98:
	;
	v377 = v359 + int32(1)
	if v377 != v353 {
		v359 = v377
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v157)+28))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	v398 = v395
	v399 = v395 + int32(4)
	goto L39
L101:
	;
	v516 = F_match_expr_to_partition_keys(m, v514, l2, v403)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L102
	} else {
		goto L137
	}
L102:
	;
	return
L103:
	;
	if v403 == int32(0) {
		v514 = v401
		v515 = v400
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v409 = int32(0)
	if base.B2i32(v407 == v409)|base.B2i32(v408 == v409) != 0 {
		v454 = v409
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v454 != 0 {
		goto L118
	} else {
		goto L119
	}
L106:
	;
	goto L105
L107:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	if v419 < v420 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v422 = v419
	goto L110
L109:
	;
	v422 = v420
	goto L110
L110:
	;
	if v422 <= int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v425 = int32(1)
	goto L113
L112:
	;
	v425 = v422
	goto L113
L113:
	;
	v426 = int32(8)
	v431 = int32(0)
	goto L114
L114:
	;
	v438 = v431 << (uint(int32(2)) % 32)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v408+v426+v438)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v407+v426+v438)))
	v443 = v440 & v442
	v445 = base.B2i32(v443 != int32(0))
	if v443 != 0 {
		v454 = v445
		goto L106
	} else {
		goto L116
	}
L115:
	;
	v454 = v445
	goto L106
L116:
	;
	v447 = v431 + int32(1)
	if v447 != v425 {
		v431 = v447
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v457 = F_remove_nulling_relids(m, v401, v455, int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L102
	} else {
		goto L121
	}
L119:
	;
	v459 = v401
	goto L120
L120:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v462 = int32(0)
	if base.B2i32(v460 == v462)|base.B2i32(v461 == v462) != 0 {
		v507 = v462
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v459 = v457
	goto L120
L122:
	;
	if v507 == int32(0) {
		v514 = v459
		v515 = v400
		goto L101
	} else {
		goto L135
	}
L123:
	;
	goto L122
L124:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v472 < v473 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v475 = v472
	goto L127
L126:
	;
	v475 = v473
	goto L127
L127:
	;
	if v475 <= int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v478 = int32(1)
	goto L130
L129:
	;
	v478 = v475
	goto L130
L130:
	;
	v479 = int32(8)
	v484 = int32(0)
	goto L131
L131:
	;
	v491 = v484 << (uint(int32(2)) % 32)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v461+v479+v491)))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v460+v479+v491)))
	v496 = v493 & v495
	v498 = base.B2i32(v496 != int32(0))
	if v496 != 0 {
		v507 = v498
		goto L123
	} else {
		goto L133
	}
L132:
	;
	v507 = v498
	goto L123
L133:
	;
	v500 = v484 + int32(1)
	if v500 != v478 {
		v484 = v500
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v512 = F_remove_nulling_relids(m, v400, v510, int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L102
	} else {
		goto L136
	}
L136:
	;
	v514 = v459
	v515 = v512
	goto L101
L137:
	;
	if v516 < int32(0) {
		v603 = v78
		goto L14
	} else {
		goto L138
	}
L138:
	;
	v520 = F_match_expr_to_partition_keys(m, v515, l3, v403)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L102
	} else {
		goto L139
	}
L139:
	;
	if v520 != v516 {
		v603 = v78
		goto L14
	} else {
		goto L140
	}
L140:
	;
	v525 = v23 + int32(32) + v516
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	if v526 != 0 {
		v603 = v78
		goto L14
	} else {
		goto L141
	}
L141:
	;
	v528 = v516 << (uint(int32(2)) % 32)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l2)+256))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+12))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528+v530)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v157)+24))
	if v532 != v533 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v535 == int32(104) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v592 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v525))) = uint8(v592)
	v595 = v78 + v592
	v596 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+2)))
	if v595 == v596 {
		goto L8
	} else {
		goto L164
	}
L144:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v88)+124))
	if v538 == int32(0) {
		v603 = v78
		goto L14
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v88)+96))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v547+v528)))
	v550 = int32(0)
	if v546 == v550 {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v541+v528)))
	v544 = F_op_in_opfamily(m, v538, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L102
	} else {
		goto L148
	}
L148:
	;
	if v544 != 0 {
		goto L143
	} else {
		goto L149
	}
L149:
	;
	v603 = v78
	goto L14
L150:
	;
	if v588 == int32(0) {
		v603 = v78
		goto L14
	} else {
		goto L163
	}
L151:
	;
	v588 = int32(0)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	if v556 <= int32(0) {
		v582 = v550
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v588 = v582
	goto L150
L155:
	;
	v559 = int32(0)
	if v559 < v556 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v562 = v556
	goto L158
L157:
	;
	v562 = v559
	goto L158
L158:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v546)+12))
	v565 = int32(0)
	goto L159
L159:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v563+v565<<(uint(int32(2))%32))))
	v574 = base.B2i32(v573 == v549)
	if v573 == v549 {
		v582 = v574
		goto L154
	} else {
		goto L161
	}
L160:
	;
	v582 = v574
	goto L154
L161:
	;
	v576 = v565 + int32(1)
	if v576 != v562 {
		v565 = v576
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
	v603 = v595
	goto L14
L165:
	;
	goto L13
L166:
	;
	v636 = v628
	v645 = v622
	v646 = v7
	goto L167
L167:
	;
	v653 = v23 + int32(32) + v646
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653))))
	if v654 == int32(1) {
		v802 = v636
		v811 = v645
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L2
L169:
	;
	v818 = v646 + int32(1)
	if v818 < v802 {
		v636 = v802
		v645 = v811
		v646 = v818
		goto L167
	} else {
		goto L199
	}
L170:
	;
	v658 = v646 << (uint(int32(2)) % 32)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v660 = v658 + v659
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v661 == int32(104) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v660)))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v665+v658)))
	v669 = F_get_opfamily_member(m, v664, v667, v667, int32(1))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L102
	} else {
		goto L174
	}
L172:
	;
	v678 = v660
	goto L173
L173:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l2)+288))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v679+v658)))
	if v681 == int32(0) {
		goto L2
	} else {
		goto L178
	}
L174:
	;
	if v669 == int32(0) {
		goto L2
	} else {
		goto L175
	}
L175:
	;
	v673 = F_get_mergejoin_opfamilies(m, v669)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L102
	} else {
		goto L176
	}
L176:
	;
	if v673 == int32(0) {
		goto L2
	} else {
		goto L177
	}
L177:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v673)+12))
	v678 = v677
	goto L173
L178:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	if v684 <= int32(0) {
		goto L2
	} else {
		goto L179
	}
L179:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	v699 = int32(0)
	goto L180
L180:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l2)+256))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v709)+12))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v710+v658)))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v681)+12))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v713+v699<<(uint(int32(2))%32))))
	v718 = F_exprCollation(m, v717)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L102
	} else {
		goto L182
	}
L181:
	;
	goto L2
L182:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l3)+288))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v720+v658)))
	if v722 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v794 = v699 + int32(1)
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	if v794 < v795 {
		v699 = v794
		goto L180
	} else {
		goto L198
	}
L184:
	;
	v725 = int32(0)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
	if v726 <= v725 {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v734 = v725
	goto L186
L186:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v722)+12))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v749+v734<<(uint(int32(2))%32))))
	v754 = F_exprs_known_equal(m, l0, v717, v753, v687)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L102
	} else {
		goto L188
	}
L187:
	;
	v765 = F_exprCollation(m, v753)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L102
	} else {
		goto L196
	}
L188:
	;
	if v712 == v718 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v758 = v754
	goto L191
L190:
	;
	v758 = int32(0)
	goto L191
L191:
	;
	if v758 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v762 = v734 + int32(1)
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
	if v762 < v763 {
		v734 = v762
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
	v767 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v653))) = uint8(v767)
	v769 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+2)))
	v771 = v645 + v767
	if v769 == v771 {
		goto L8
	} else {
		goto L197
	}
L197:
	;
	v802 = v769
	v811 = v771
	goto L169
L198:
	;
	goto L181
L199:
	;
	goto L168
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+288)) = v845
	v849 = F_palloc0_mul(m, int32(4), v844)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L102
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+292)) = v849
	if int32(0) < v844 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	if base.B2i32(int32(base.Ui32(int32(55))>>(uint(v842)%32))&int32(1) == int32(0))|base.B2i32(base.Ui32(int32(5)) < base.Ui32(v842)) != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v1089 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+233)) = uint8(v1089)
	goto L2
L205:
	;
	v874 = int32(0)
	goto L206
L206:
	;
	v885 = v874 << (uint(int32(2)) % 32)
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l3)+292))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v885+v886)))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l3)+288))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v889+v885)))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l2)+292))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v892+v885)))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l2)+288))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v895+v885)))
	switch v842 {
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
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l1)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v1060+v885))) = v1046
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l1)+292))
	*(*int32)(unsafe.Add(mBase, uint32(v1063+v885))) = v1044
	v1067 = v874 + int32(1)
	if v1067 != v844 {
		v874 = v1067
		goto L206
	} else {
		goto L243
	}
L209:
	;
	v1036 = F_list_concat_copy(m, v897, v891)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L102
	} else {
		goto L241
	}
L210:
	;
	v908 = F_list_concat_copy(m, v897, v891)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L102
	} else {
		goto L218
	}
L211:
	;
	v902 = F_list_copy(m, v897)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L102
	} else {
		goto L215
	}
L212:
	;
	v898 = F_list_copy(m, v897)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L102
	} else {
		goto L213
	}
L213:
	;
	v900 = F_list_copy(m, v894)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L102
	} else {
		goto L214
	}
L214:
	;
	v1044 = v900
	v1046 = v898
	goto L208
L215:
	;
	v904 = F_list_concat_copy(m, v891, v894)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L102
	} else {
		goto L216
	}
L216:
	;
	v906 = F_list_concat(m, v904, v888)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L102
	} else {
		goto L217
	}
L217:
	;
	v1044 = v906
	v1046 = v902
	goto L208
L218:
	;
	v910 = F_list_concat(m, v908, v894)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L102
	} else {
		goto L219
	}
L219:
	;
	v912 = F_list_concat(m, v910, v888)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L102
	} else {
		goto L220
	}
L220:
	;
	v914 = F_list_concat_copy(m, v897, v894)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L102
	} else {
		goto L221
	}
L221:
	;
	if v914 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1044 = v912
	v1046 = int32(0)
	goto L208
L223:
	;
	goto L224
L224:
	;
	v919 = int32(0)
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v914)+4))
	if v921 <= v919 {
		v1044 = v912
		v1046 = v919
		goto L208
	} else {
		goto L225
	}
L225:
	;
	v928 = v912
	v931 = v919
	goto L226
L226:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v914)+12))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v944+v931<<(uint(int32(2))%32))))
	v949 = F_list_concat_copy(m, v891, v888)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L102
	} else {
		goto L229
	}
L227:
	;
	v1044 = v1016
	v1046 = v919
	goto L208
L228:
	;
	v1033 = v931 + int32(1)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v914)+4))
	if v1033 < v1034 {
		v928 = v1016
		v931 = v1033
		goto L226
	} else {
		goto L240
	}
L229:
	;
	if v949 == int32(0) {
		v1016 = v928
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v953 = int32(0)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v949)+4))
	if v954 <= v953 {
		v1016 = v928
		goto L228
	} else {
		goto L231
	}
L231:
	;
	v957 = v953
	v961 = v928
	goto L232
L232:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v949)+12))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v977+v957<<(uint(int32(2))%32))))
	v983 = F_palloc0(m, int32(20))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L102
	} else {
		goto L234
	}
L233:
	;
	v1016 = v1006
	goto L228
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v983))) = int32(38)
	v987 = F_exprType(m, v948)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L102
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v983)+4)) = v987
	v990 = F_exprCollation(m, v948)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L102
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v983)+8)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v948
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v948
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v981
	v1001 = F_list_make2_impl(m, v23+int32(12), v23+int32(8))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L102
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v983)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v983)+12)) = v1001
	v1006 = F_lappend(m, v961, v983)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L102
	} else {
		goto L238
	}
L238:
	;
	v1009 = v957 + int32(1)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v949)+4))
	if v1009 < v1010 {
		v957 = v1009
		v961 = v1006
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
	v1038 = F_list_concat_copy(m, v894, v888)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L102
	} else {
		goto L242
	}
L242:
	;
	v1044 = v1038
	v1046 = v1036
	goto L208
L243:
	;
	goto L207
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v842
	F_errmsg_internal(m, int32(_a_F_build_joinrel_partition_info_0), v23+int32(16))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L102
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_build_joinrel_partition_info_1), int32(2696), int32(_a_F_build_joinrel_partition_info_2))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
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
func F_build_reloptions(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v37 int32
	_ = v37
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
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
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
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
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
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
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
	if base.I32_wrap_i64(l0) != 0 {
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
	v37 = v7
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v45 = v36 + base.B2i32(v41&l2 != int32(0))
	v47 = v37 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23+v47<<(uint(int32(2))%32))))
	if v51 != 0 {
		v34 = v51
		v36 = v45
		v37 = v47
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
	v74 = int32(0)
	v75 = v7
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
	v82 = v57 + v75<<(uint(int32(4))%32)
	v83 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v71
	v88 = v75 + int32(1)
	goto L22
L21:
	;
	v88 = v75
	goto L22
L22:
	;
	v91 = v74 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v60+v91<<(uint(int32(2))%32))))
	if v95 != 0 {
		v71 = v95
		v74 = v91
		v75 = v88
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
	v111 = m.ExcPending
	if v111 != 0 {
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
	v125 = int32(0)
	v129 = l3
	goto L34
L32:
	;
	v180 = l3
	goto L33
L33:
	;
	v183 = F_palloc0(m, v180)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L56
	}
L34:
	;
	v134 = v103 + v125<<(uint(int32(4))%32)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
	if v136 == int32(5) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v180 = v164
	goto L33
L36:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+4)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+36))
	if v140 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v164 = v129
	goto L38
L38:
	;
	v168 = v125 + int32(1)
	if v168 != v104 {
		v125 = v168
		v129 = v164
		goto L34
	} else {
		goto L55
	}
L39:
	;
	v164 = v162 + v129
	goto L38
L40:
	;
	if v139&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v139&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v145 = m.T0[v140].(func(*base.Module, int32, int32) int32)(m, v143, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+28)))
	if v147 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v162 = v145
	goto L39
L47:
	;
	v150 = int32(0)
	goto L49
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v135)+40))
	v150 = v149
	goto L49
L49:
	;
	v152 = m.T0[v140].(func(*base.Module, int32, int32) int32)(m, v150, int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v162 = v152
	goto L39
L51:
	;
	v162 = v159 + int32(1)
	goto L39
L52:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v157 = F_strlen(m, v156)
	mBase = m.M
	v159 = v157
	goto L51
L53:
	;
	goto L54
L54:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
	v159 = v158
	goto L51
L55:
	;
	goto L35
L56:
	;
	F_fillRelOptions(m, v183, l3, v103, v104, l1, l4, l5)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v103)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	return v183
}
func F_build_replindex_scan_key(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
	v24 = F_SysCacheGetAttrNotNull(m, int32(34), v22, int32(18))
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+10)))
	if int32(0) < v29 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L23
	}
L4:
	;
	v41 = v28
	v42 = int32(0)
	v45 = v4
	goto L7
L5:
	;
	v127 = v4
	goto L6
L6:
	;
	m.G0 = v18 + int32(16)
	return v127
L7:
	;
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20+int32(48)+v42<<(uint(int32(1))%32)))))
	if v56 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v127 = v112
	goto L6
L9:
	;
	v58 = v42 << (uint(int32(2)) % 32)
	v59 = base.I32_wrap_i64(v24) + int32(24) + v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = F_get_opclass_input_type(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v110 = v41
	v112 = v45
	goto L11
L11:
	;
	v117 = v42 + int32(1)
	v118 = int32(*(*int16)(unsafe.Add(mBase, uint32(v110)+10)))
	if v117 < v118 {
		v41 = v110
		v42 = v117
		v45 = v112
		goto L7
	} else {
		goto L22
	}
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v64 = F_get_opclass_family(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+84))
	v70 = F_IndexAmTranslateCompareType(m, int32(3), v68, v64, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v73 = F_get_opfamily_member(m, v64, v61, v61, base.I32_extend16_s(v70))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v73 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v79 = l0 + v45*int32(56)
	v83 = F_get_opcode(m, v73)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v87 = v56 - int32(1)
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v85+v87<<(uint(int32(3))%32))))
	F_ScanKeyInit(m, v79, base.I32_extend16_s(v42+int32(1)), v70, v83, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v94+v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v87))))
	if v100 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v103 | int32(65)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v110 = v109
	v112 = v45 + int32(1)
	goto L11
L22:
	;
	goto L8
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v70
	F_errmsg_internal(m, int32(_a_F_build_replindex_scan_key_0), v18)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_build_replindex_scan_key_1), int32(104), int32(_a_F_build_replindex_scan_key_2))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_buildint2vector(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14241(m, l0, l1, int32(21), int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_byteane(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
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
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_toast_raw_datum_size(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.I64_extend_i32_u(v111)
L2:
	;
	return int64(0)
L3:
	;
	v14 = F_toast_raw_datum_size(m, v7)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v10 != v14 {
		v111 = int32(1)
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v18 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v9))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v21 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v7))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v25&v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v28 = v23
	goto L10
L9:
	;
	v28 = int32(4)
	goto L10
L10:
	;
	v29 = v18 + v28
	v30 = int32(1)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v32&v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = v30
	goto L13
L12:
	;
	v35 = int32(4)
	goto L13
L13:
	;
	v36 = v21 + v35
	v37 = int32(4)
	v38 = v10 - v37
	if base.Ui32(v37) <= base.Ui32(v38) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v101 != v18 {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v100 = int32(0)
	goto L14
L16:
	;
	v74 = v69
	v75 = v70
	v76 = v71
	goto L26
L17:
	;
	if (v29|v36)&int32(3) != 0 {
		v69 = v29
		v70 = v36
		v71 = v38
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v62 = v29
	v63 = v36
	v64 = v38
	goto L19
L19:
	;
	if v64 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v46 = v29
	v47 = v36
	v48 = v38
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v51 != v52 {
		v69 = v46
		v70 = v47
		v71 = v48
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v62 = v57
	v63 = v55
	v64 = v59
	goto L19
L23:
	;
	v54 = int32(4)
	v55 = v47 + v54
	v57 = v46 + v54
	v59 = v48 - v54
	if base.Ui32(int32(3)) < base.Ui32(v59) {
		v46 = v57
		v47 = v55
		v48 = v59
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v69 = v62
	v70 = v63
	v71 = v64
	goto L16
L26:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 == v80 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v100 = v79 - v80
	goto L14
L28:
	;
	v82 = int32(1)
	v87 = v76 - v82
	if v87 != 0 {
		v74 = v74 + v82
		v75 = v75 + v82
		v76 = v87
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
	F_pfree(m, v18)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v106 = base.B2i32(v100 != int32(0))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v21 == v107 {
		v111 = v106
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	F_pfree(m, v21)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v111 = v106
	goto L1
}
func F_byteanlike(m *base.Module, l0 int32) int64 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v15 == int32(1) {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
				if v21 == int32(18) {
					v24 = int32(16)
				} else {
					v24 = int32(0)
				}
				if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v31 = int32(4)
				} else {
					v31 = v24
				}
				v44 = v31
			} else {
				v32 = int32(1)
				if v15&v32 != 0 {
					v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v45 == int32(1) {
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
				if v51 == int32(18) {
					v54 = int32(16)
				} else {
					v54 = int32(0)
				}
				if base.Ui32((v51-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v61 = int32(4)
				} else {
					v61 = v54
				}
				v74 = v61
			} else {
				v62 = int32(1)
				if v45&v62 != 0 {
					v74 = int32(base.Ui32(v45)>>(uint(v62)%32)) - v62
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v75 = int32(1)
			if v15&v75 != 0 {
				v79 = v75
			} else {
				v79 = int32(4)
			}
			v81 = int32(1)
			if v45&v81 != 0 {
				v85 = v81
			} else {
				v85 = int32(4)
			}
			v88 = F_SB_MatchText(m, v8+v79, v44, v13+v85, v74, int32(0))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(base.B2i32(v88 != int32(1)))
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
