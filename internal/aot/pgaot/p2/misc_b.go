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
func F_BuildDescForRelation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
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
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v174
L2:
	;
	v20 = F_CreateTemplateTupleDesc(m, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = F_CreateTemplateTupleDesc(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v174 = v20
	goto L1
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v27 <= int32(0) {
		v174 = v25
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v39 = v2
	v40 = v2
	goto L10
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L40
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v40<<(uint(int32(2))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	F_typenameTypeIdAndMod(m, int32(0), v51, v15+int32(12), v15+int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L12
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L36
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_BuildDescForRelation[0]))
	v63 = F_object_aclcheck(m, int32(1247), v59, v61, int64(256))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	if v63 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_aclcheck_error_type(m, v63, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v68 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v71 = F_GetColumnDefCollation(m, v68, v48, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
	if v74 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L11
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if int32(_a_F_BuildDescForRelation_0) <= v75 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	v78 = v68
	goto L22
L22:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+12)))
	if v79 == int32(1) {
		goto L9
	} else {
		goto L24
	}
L23:
	;
	v78 = v75
	goto L22
L24:
	;
	v84 = base.I32_extend16_s(v39 + int32(1))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_TupleDescInitEntry(m, v25, v84, v49, v85, v86, v78)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v25+v89<<(uint(int32(4))%32)+v84*int32(100))+16)) = v71
	goto L26
L26:
	;
	v102 = v84 - int32(1)
	v105 = v25 + int32(20) + v89<<(uint(int32(4))%32) + v102*int32(100)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+86)) = uint8(v106)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+92)) = uint8(v108)
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v105)+94)) = uint16(v110)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+89)) = uint8(v112)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+90)) = uint8(v114)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v105)+68))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v118 = F_GetAttributeCompression(m, v116, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+85)) = uint8(v118)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+21)))
	if v121 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_populate_compact_attribute(m, v25, v102)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L34
	}
L29:
	;
	v129 = v121
	goto L31
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v122 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+84)) = uint8(v129)
	goto L28
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v105)+68))
	v126 = F_GetAttributeStorage(m, v125, v122)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v129 = v126
	goto L31
L34:
	;
	v135 = v40 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v136 <= v135 {
		v174 = v25
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v39 = v84
	v40 = v135
	goto L10
L36:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_BuildDescForRelation_1), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_BuildDescForRelation_2), int32(1425), int32(_a_F_BuildDescForRelation_3))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
	v160 = m.ExcPending
	if v160 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
	F_errmsg(m, int32(_a_F_BuildDescForRelation_4), v15)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_BuildDescForRelation_2), int32(1431), int32(_a_F_BuildDescForRelation_3))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	v70 = v21 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v70 < v71 {
		v21 = v70
		goto L4
	} else {
		goto L24
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	if base.Ui32(v31) < base.Ui32(v30) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v34&int32(1) == int32(0) {
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+13)))
	v34 = v33
	goto L11
L10:
	;
	v34 = int32(1)
	goto L11
L11:
	;
	goto L8
L12:
	;
	v39 = F_IsBinaryTidClause(m, v26, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v39 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v44 != int32(387) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v47 = F_join_clause_is_movable_to(m, v26, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	if v47 == int32(0) {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v26
	v56 = F_list_make1_impl(m, int32(1), v9+int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v60 = F_bms_union(m, v58, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v63 = F_bms_del_member(m, v60, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v65 = F_create_tidscan_path(m, l0, l1, v56, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	F_add_path(m, l1, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
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
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = v10 - int32(2147483641)
	if base.Ui32(v12) < base.Ui32(int32(-2147483640)) {
		v15 = int32(1)
	} else {
		v15 = v10
	}
	v18 = int32(8)
	v19 = base.I32_div_s(v15+int32(7), v18)
	v21 = v19 + v18
	v22 = F_palloc(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v21 << (uint(int32(2)) % 32)
		v31 = v22 + int32(8)
		v32 = int32(32)
		if v32 <= v15 {
			v35 = v32
		} else {
			v35 = v15
		}
		v37 = v35 + int32(8)
		if v37 <= v15 {
			v41 = int32(-2147483640)
			if base.Ui32(v12) <= base.Ui32(v41) {
				v44 = v41
			} else {
				v44 = v12
			}
			v46 = v44 + int32(2147483633)
			v47 = v46 - v35
			v49 = int32(base.Ui32(v47) >> (uint(int32(3)) % 32))
			v53 = F__emscripten_memset_bulkmem(m, v31, base.I32_extend8_s(v8>>(uint(int32(31))%32)), v49+int32(1))
			mBase = m.M
			v60 = v49 + v22 + int32(9)
			v61 = v46 - v47&int32(-8)
		} else {
			v60 = v31
			v61 = v15
		}
		if v35 < v61 {
			v71 = v61 - int32(8)
			v73 = int32(-1)<<(uint(v37-v61)%32)&(v8>>(uint(int32(31))%32)) | v8>>(uint(v71)%32)
			*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v73)
			v77 = v60 + int32(1)
			v78 = v71
		} else {
			v77 = v60
			v78 = v61
		}
		if v78 < int32(8) {
			v162 = v77
			v163 = v78
		} else {
			v82 = v78 - int32(8)
			v83 = int32(56)
			if v82&v83 != v83 {
				v94 = v77
				v95 = v78
				v96 = int32(0)
				for {
					v102 = v95 - int32(8)
					v103 = v8 >> (uint(v102) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v103)
					v105 = int32(1)
					v106 = v94 + v105
					v108 = v96 + v105
					if v108 != (int32(base.Ui32(v82)>>(uint(int32(3))%32))+int32(1))&int32(7) {
						v94 = v106
						v95 = v102
						v96 = v108
						continue
					} else {
						break
					}
					break
				}
				v110 = v106
				v111 = v102
			} else {
				v110 = v77
				v111 = v78
			}
			if base.Ui32(v82) < base.Ui32(int32(56)) {
				v162 = v110
				v163 = v111
			} else {
				v119 = v110
				v121 = v111
				for {
					v127 = v121 + int32(-64)
					v128 = v8 >> (uint(v127) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)) = uint8(v128)
					v131 = v121 - int32(56)
					v132 = v8 >> (uint(v131) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)) = uint8(v132)
					v136 = v8 >> (uint(v121-int32(48)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v119)+5)) = uint8(v136)
					v140 = v8 >> (uint(v121-int32(40)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)) = uint8(v140)
					v144 = v8 >> (uint(v121-int32(32)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)) = uint8(v144)
					v148 = v8 >> (uint(v121-int32(24)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)) = uint8(v148)
					v152 = v8 >> (uint(v121-int32(16)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)) = uint8(v152)
					v154 = int32(8)
					v156 = v8 >> (uint(v121-v154) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v156)
					v159 = v119 + v154
					if base.Ui32(int32(15)) < base.Ui32(v131) {
						v119 = v159
						v121 = v127
						continue
					} else {
						break
					}
					break
				}
				v162 = v159
				v163 = v127
			}
		}
		if int32(0) < v163 {
			v173 = v8 << (uint(int32(8)-v163) % 32)
			*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v173)
		} else {
		}
		return v22
	}
}
func F_bitfromint8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v121 int64
	_ = v121
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = v13 - int32(2147483641)
	if base.Ui32(v15) < base.Ui32(int32(-2147483640)) {
		v18 = int32(1)
	} else {
		v18 = v13
	}
	v21 = int32(8)
	v22 = base.I32_div_s(v18+int32(7), v21)
	v24 = v22 + v21
	v25 = F_palloc(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v25))) = v24 << (uint(int32(2)) % 32)
		v34 = v25 + int32(8)
		v35 = int32(64)
		if v35 <= v18 {
			v38 = v35
		} else {
			v38 = v18
		}
		v40 = v38 + int32(8)
		if v40 <= v18 {
			v45 = int32(-2147483640)
			if base.Ui32(v15) <= base.Ui32(v45) {
				v48 = v45
			} else {
				v48 = v15
			}
			v50 = v48 + int32(2147483633)
			v51 = v50 - v38
			v53 = int32(base.Ui32(v51) >> (uint(int32(3)) % 32))
			v57 = F__emscripten_memset_bulkmem(m, v34, base.I32_extend8_s(base.I32_wrap_i64(v11>>(uint(int64(63))%64))), v53+int32(1))
			mBase = m.M
			v64 = v53 + v25 + int32(9)
			v65 = v50 - v51&int32(-8)
		} else {
			v64 = v34
			v65 = v18
		}
		if v38 < v65 {
			v76 = v65 - int32(8)
			v80 = base.I32_wrap_i64(v11>>(uint(int64(63))%64))&(int32(-1)<<(uint(v40-v65)%32)) | base.I32_wrap_i64(v11>>(uint(base.I64_extend_i32_u(v76))%64))
			*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v80)
			v84 = v64 + int32(1)
			v85 = v76
		} else {
			v84 = v64
			v85 = v65
		}
		if int32(8) <= v85 {
			v89 = v84
			v96 = base.I64_extend_i32_u(v85)
			for {
				v99 = v96 - int64(8)
				v100 = v11 >> (uint(v99) % 64)
				*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v100)
				v103 = v89 + int32(1)
				if base.Ui64(int64(15)) < base.Ui64(v96) {
					v89 = v103
					v96 = v99
					continue
				} else {
					break
				}
				break
			}
			v107 = v103
			v108 = base.I32_wrap_i64(v99)
		} else {
			v107 = v84
			v108 = v85
		}
		if int32(0) < v108 {
			v121 = v11 << (uint(base.I64_extend_i32_u(int32(8)-v108)) % 64)
			*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v121)
		} else {
		}
		return v25
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v15 < int32(0) {
			v19 = int32(0)
			v21 = int32(-2147483640)
			if base.Ui32(v15) <= base.Ui32(v21) {
				v24 = v21
			} else {
				v24 = v15
			}
			v26 = F_DirectFunctionCall2Coll(m, int32(1548), v19, v11, v19-v24)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v26
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v32 = F_palloc(m, int32(base.Ui32(v29)>>(uint(int32(2))%32)))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v36 = v34 & int32(-4)
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v38
				v41 = v32 + int32(8)
				if v38 <= v15 {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
					v47 = v45 - int32(8)
					if v41&int32(3) != 0 {
						v73 = F__emscripten_memset_bulkmem(m, v41, base.I32_extend8_s(int32(0)), v47)
						mBase = m.M
						return v32
					} else {
						if v43&int32(12) != 0 {
							v73 = F__emscripten_memset_bulkmem(m, v41, base.I32_extend8_s(int32(0)), v47)
							mBase = m.M
							return v32
						} else {
							if base.Ui32(int32(1024)) < base.Ui32(v47) {
								v73 = F__emscripten_memset_bulkmem(m, v41, base.I32_extend8_s(int32(0)), v47)
								mBase = m.M
								return v32
							} else {
								v54 = v32 + v45
								if base.Ui32(v54) <= base.Ui32(v41) {
									return v32
								} else {
									v58 = v32 + int32(12)
									if base.Ui32(v58) < base.Ui32(v54) {
										v60 = v54
									} else {
										v60 = v58
									}
									v69 = F__emscripten_memset_bulkmem(m, v41, base.I32_extend8_s(int32(0)), (v60-v32-int32(9))&int32(-4)+int32(4))
									mBase = m.M
									return v32
								}
							}
						}
					}
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v77 = int32(base.Ui32(v75) >> (uint(int32(2)) % 32))
					v79 = int32(base.Ui32(v15) >> (uint(int32(3)) % 32))
					v82 = v11 + v79 + int32(8)
					v84 = v15 & int32(7)
					if v84 != 0 {
						if base.Ui32(v82) < base.Ui32(v11+v77) {
							v89 = v41
							v92 = v82
							for {
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
								v99 = v98 << (uint(v84) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v99)
								v102 = v92 + int32(1)
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								v105 = int32(base.Ui32(v103) >> (uint(int32(2)) % 32))
								if base.Ui32(v102) < base.Ui32(v11+v105) {
									v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
									v110 = int32(base.Ui32(v108)>>(uint(int32(8)-v84)%32)) | v99
									*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v110)
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
									v115 = int32(base.Ui32(v112) >> (uint(int32(2)) % 32))
								} else {
									v115 = v105
								}
								v117 = v89 + int32(1)
								if base.Ui32(v102) < base.Ui32(v11+v115) {
									v89 = v117
									v92 = v102
									continue
								} else {
									break
								}
								break
							}
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
							v121 = v117
							v130 = v120
						} else {
							v121 = v41
							v130 = v36
						}
						if base.Ui32(v32+int32(base.Ui32(v130)>>(uint(int32(2))%32))) <= base.Ui32(v121) {
						} else {
							v135 = v121
							for {
								v144 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v144)
								v147 = v135 + int32(1)
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
								if base.Ui32(v147) < base.Ui32(v32+int32(base.Ui32(v148)>>(uint(int32(2))%32))) {
									v135 = v147
									continue
								} else {
									break
								}
								break
							}
						}
						return v32
					} else {
						v153 = v77 - v79
						v155 = v153 - int32(8)
						if v155 != 0 {
							v156 = F__emscripten_memcpy_bulkmem(m, v41, v82, v155)
							mBase = m.M
						} else {
						}
						v158 = v153 + v32
						if v158&int32(3) != 0 {
							v187 = F__emscripten_memset_bulkmem(m, v158, base.I32_extend8_s(int32(0)), v79)
							mBase = m.M
							return v32
						} else {
							if base.Ui32(int32(_a_F_bitshiftleft_0)) < base.Ui32(v15) {
								v187 = F__emscripten_memset_bulkmem(m, v158, base.I32_extend8_s(int32(0)), v79)
								mBase = m.M
								return v32
							} else {
								if v15&int32(24) != 0 {
									v187 = F__emscripten_memset_bulkmem(m, v158, base.I32_extend8_s(int32(0)), v79)
									mBase = m.M
									return v32
								} else {
									v165 = v32 + v77
									if base.Ui32(v165) <= base.Ui32(v158) {
										return v32
									} else {
										v170 = v165 - v79 + int32(4)
										if base.Ui32(v165) < base.Ui32(v170) {
											v172 = v170
										} else {
											v172 = v165
										}
										v183 = F__emscripten_memset_bulkmem(m, v158, base.I32_extend8_s(int32(0)), (v172+v79+(v32^int32(-1))-v77)&int32(-4)+int32(4))
										mBase = m.M
										return v32
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(2))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			v11 = int32(116)
		} else {
			v11 = int32(102)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v11)
		v13 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v13)
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
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
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
	v179 = m.ExcPending
	if v179 != 0 {
		goto L53
	} else {
		goto L57
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L53
	} else {
		goto L54
	}
L3:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v144
	m.G0 = v15 + int32(32)
	return
L4:
	;
	v105 = v103 * int32(92)
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v108)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v112)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v116)
	v118 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v118)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[4])))
	if v122 != 0 {
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
	v143 = v55 + int32(108)
	goto L3
L50:
	;
	v123 = v122
	goto L52
L51:
	;
	v123 = l0
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v123
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v105)+uint32(_c_F_boot_get_type_io_data[5])))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v127
	v143 = v105 + int32(_a_F_boot_get_type_io_data_0)
	goto L3
L53:
	;
	return
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_boot_get_type_io_data_1), v15+int32(16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_boot_get_type_io_data_2), int32(860), int32(_a_F_boot_get_type_io_data_3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_boot_get_type_io_data_4), v15)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_boot_get_type_io_data_2), int32(887), int32(_a_F_boot_get_type_io_data_3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
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
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
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
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v23 = v21 & v19
	if v21 == v19 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L94
	}
L7:
	;
	if v23 != 0 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v28&int32(254) == int32(2) {
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
	if v23 != 0 {
		v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v37 = v26
	goto L13
L12:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L13
L13:
	;
	if v28 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = v26
	goto L16
L15:
	;
	v40 = v37
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
	v54 = v20
	goto L20
L19:
	;
	v54 = v11 + int32(4)
	goto L20
L20:
	;
	v59 = v51
	goto L21
L21:
	;
	if v59 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v77 = int32(1)
	v78 = v16 + v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v81 = v79 & v77
	if v79 == v77 {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	v76 = v51 >> (uint(int32(31)) % 32) & v51
	goto L23
L25:
	;
	goto L26
L26:
	;
	v70 = v59 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v70))))
	if v72 == int32(32) {
		v59 = v70
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v76 = v59
	goto L23
L28:
	;
	if v81 != 0 {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	v84 = int32(4)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v86&int32(254) == int32(2) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v99 = int32(1)
	if v81 != 0 {
		v109 = int32(base.Ui32(v79)>>(uint(v99)%32)) - v99
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v95 = v84
	goto L34
L33:
	;
	v95 = base.B2i32(v86 == int32(18)) << (uint(v84) % 32)
	goto L34
L34:
	;
	if v86 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v98 = v84
	goto L37
L36:
	;
	v98 = v95
	goto L37
L37:
	;
	v109 = v98
	goto L28
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	v112 = v78
	goto L41
L40:
	;
	v112 = v16 + int32(4)
	goto L41
L41:
	;
	v117 = v109
	goto L42
L42:
	;
	if v117 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v135 = F_pg_newlocale_from_collation(m, v18)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	goto L43
L45:
	;
	v134 = v109 >> (uint(int32(31)) % 32) & v109
	goto L44
L46:
	;
	goto L47
L47:
	;
	v128 = v117 - int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v128))))
	if v130 == int32(32) {
		v117 = v128
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v134 = v117
	goto L44
L49:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v239 != v11 {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v137 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v134 != v76 {
		v238 = int32(0)
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v220 = int32(1)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v222&v220 != 0 {
		goto L79
	} else {
		goto L80
	}
L54:
	;
	v142 = int32(1)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v144&v142 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v147 = v142
	goto L57
L56:
	;
	v147 = int32(4)
	goto L57
L57:
	;
	v148 = v11 + v147
	v149 = int32(1)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v151&v149 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v154 = v149
	goto L60
L59:
	;
	v154 = int32(4)
	goto L60
L60:
	;
	v155 = v16 + v154
	if base.Ui32(int32(4)) <= base.Ui32(v76) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v238 = base.B2i32(v217 == int32(0))
	goto L49
L62:
	;
	v217 = int32(0)
	goto L61
L63:
	;
	v191 = v186
	v192 = v187
	v193 = v188
	goto L73
L64:
	;
	if (v148|v155)&int32(3) != 0 {
		v186 = v148
		v187 = v155
		v188 = v76
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v179 = v148
	v180 = v155
	v181 = v76
	goto L66
L66:
	;
	if v181 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v163 = v148
	v164 = v155
	v165 = v76
	goto L68
L68:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v168 != v169 {
		v186 = v163
		v187 = v164
		v188 = v165
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v179 = v174
	v180 = v172
	v181 = v176
	goto L66
L70:
	;
	v171 = int32(4)
	v172 = v164 + v171
	v174 = v163 + v171
	v176 = v165 - v171
	if base.Ui32(int32(3)) < base.Ui32(v176) {
		v163 = v174
		v164 = v172
		v165 = v176
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v186 = v179
	v187 = v180
	v188 = v181
	goto L63
L73:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v196 == v197 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v217 = v196 - v197
	goto L61
L75:
	;
	v199 = int32(1)
	v204 = v193 - v199
	if v204 != 0 {
		v191 = v191 + v199
		v192 = v192 + v199
		v193 = v204
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
	v225 = v220
	goto L81
L80:
	;
	v225 = int32(4)
	goto L81
L81:
	;
	v227 = int32(1)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v229&v227 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v232 = v227
	goto L84
L83:
	;
	v232 = int32(4)
	goto L84
L84:
	;
	v234 = F_varstr_cmp(m, v11+v225, v76, v16+v232, v134, v18)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v238 = base.B2i32(v234 == int32(0))
	goto L49
L86:
	;
	F_pfree(m, v11)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v243 != v16 {
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
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	return v238
L93:
	;
	goto L92
L94:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_bpchareq_0), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errhint(m, int32(_a_F_bpchareq_1), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_bpchareq_2), int32(738), int32(_a_F_bpchareq_3))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v22 = int32(1)
	v23 = v21 & v22
	if v21 == v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v23 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v23 != 0 {
		v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v37 = v26
	goto L10
L9:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v26
	goto L13
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v51 = v40
	goto L4
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v54 = v17
	goto L17
L16:
	;
	v54 = v12 + int32(4)
	goto L17
L17:
	;
	v59 = v51
	goto L18
L18:
	;
	if v59 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v78 = int32(1)
	v79 = v19 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v82 = v80 & v78
	if v80 == v78 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	v77 = v51 >> (uint(int32(31)) % 32) & v51
	goto L20
L22:
	;
	goto L23
L23:
	;
	v71 = v59 - int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v71))))
	if v73 == int32(32) {
		v59 = v71
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v77 = v59
	goto L20
L25:
	;
	if v82 != 0 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v85 = int32(4)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v87&int32(254) == int32(2) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v100 = int32(1)
	if v82 != 0 {
		v110 = int32(base.Ui32(v80)>>(uint(v100)%32)) - v100
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v96 = v85
	goto L31
L30:
	;
	v96 = base.B2i32(v87 == int32(18)) << (uint(v85) % 32)
	goto L31
L31:
	;
	if v87 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v85
	goto L34
L33:
	;
	v99 = v96
	goto L34
L34:
	;
	v110 = v99
	goto L25
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	v113 = v79
	goto L38
L37:
	;
	v113 = v19 + int32(4)
	goto L38
L38:
	;
	v118 = v110
	goto L39
L39:
	;
	if v118 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v137 = int32(1)
	if v21&v137 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v136 = v110 >> (uint(int32(31)) % 32) & v110
	goto L41
L43:
	;
	goto L44
L44:
	;
	v130 = v118 - int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v130))))
	if v132 == int32(32) {
		v118 = v130
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v136 = v118
	goto L41
L46:
	;
	v141 = v137
	goto L48
L47:
	;
	v141 = int32(4)
	goto L48
L48:
	;
	v143 = int32(1)
	if v80&v143 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v147 = v143
	goto L51
L50:
	;
	v147 = int32(4)
	goto L51
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = F_varstr_cmp(m, v12+v141, v77, v19+v147, v136, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v19 {
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
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return int32(base.Ui32(v150^int32(-1)) >> (uint(int32(31)) % 32))
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
func F_brinbuildempty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int64
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = l0
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v6)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v13
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
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
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	v9 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l7)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v9
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)))
	v28 = v27
	goto L3
L2:
	;
	v28 = v9
	goto L3
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[0]))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l7)+140))
	*(*int32)(unsafe.Add(mBase, _c_F_brininsert[0])) = v34
	v37 = F_palloc0(m, int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v52 = v22
	goto L6
L6:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v57 = v53 | v54<<(uint(int32(16))%32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v59 = base.I32_rem_u_s(v57, v58)
	v60 = v57 - v59
	v62 = v60 - int32(1)
	v63 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v82 = v63
	goto L11
L7:
	;
	return int32(0)
L8:
	;
	v41 = F_brin_build_desc(m, l0)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v41
	v46 = F_brinRevmapInitialize(m, l0, v37+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v46
	*(*int32)(unsafe.Add(mBase, uint32(l7)+136)) = v37
	*(*int32)(unsafe.Add(mBase, _c_F_brininsert[0])) = v30
	v52 = v37
	goto L6
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[1]))
	if v90 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v345 != 0 {
		goto L73
	} else {
		goto L74
	}
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v28&(base.B2i32(v59 == v63)&base.B2i32(v57 != v63)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v241 = F_brinGetTupleForHeapBlock(m, v70, v60, v20+int32(28), v20+int32(26), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L43
	}
L18:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	if v95 != int32(1) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v103 = F_brinGetTupleForHeapBlock(m, v70, v62, v20+int32(28), v20+int32(26), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	if v103 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v108 = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[2]))
	v114 = F_LWLockAcquire(m, v110+int32(2816), v108)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	F_LockBuffer(m, v215, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L41
	}
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[3]))
	v119 = v117 + int32(36)
	v126 = v108
	goto L27
L25:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[2]))
	F_LWLockRelease(m, v188+int32(2816))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L34
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = int32(0)
	v175 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v172)+4)) = uint16(v175)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v172)+16)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v172)+12)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v172)+8)) = v178
	v186 = v175
	goto L25
L27:
	;
	v139 = v119 + v126*int32(20)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+4)))
	if v140 == int32(0) {
		v172 = v139
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v186 = int32(0)
	goto L25
L29:
	;
	v143 = int32(1)
	v147 = v119 + (v126|v143)*int32(20)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+4)))
	if v148 != v143 {
		v172 = v147
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v155 = v119 + (v126|int32(2))*int32(20)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
	if v156 != int32(1) {
		v172 = v155
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v163 = v119 + (v126|int32(3))*int32(20)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+4)))
	if v164 != int32(1) {
		v172 = v163
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v168 = v126 + int32(4)
	if v168 != int32(256) {
		v126 = v168
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	if v186 != 0 {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	v195 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	if v195 == int32(0) {
		goto L17
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v202 + int32(4)
	F_errmsg(m, int32(_a_F_brininsert_0), v20)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_brininsert_1), int32(416), int32(_a_F_brininsert_2))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	goto L17
L41:
	;
	goto L17
L42:
	;
	goto L12
L43:
	;
	if v241 == int32(0) {
		v343 = v82
		goto L42
	} else {
		goto L44
	}
L44:
	;
	if v82 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v247 = int32(_a_F_brininsert_3)
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[0]))
	v254 = F_AllocSetContextCreateInternal(m, v249, int32(_a_F_brininsert_4), int32(0), int32(_a_F_brininsert_5), int32(_a_F_brininsert_6))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L48
	}
L46:
	;
	v257 = v82
	goto L47
L47:
	;
	v259 = F_brin_deform_tuple(m, v69, v241, int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L7
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brininsert[0])) = v254
	v257 = v254
	goto L47
L49:
	;
	v261 = F_add_values_to_range(m, l0, v69, v259, l1, l2)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v261 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_LockBuffer(m, v263, int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L7
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v263 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v343 = v257
	goto L42
L55:
	;
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v286+v287<<(uint(int32(2))%32))+20))
	v293 = int32(base.Ui32(v291) >> (uint(int32(17)) % 32))
	v294 = int32(0)
	v296 = F_brin_copy_tuple(m, v241, v293, v294, v294)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L59
	}
L56:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[5]))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v272+(v263^int32(-1))<<(uint(int32(2))%32))))
	v286 = v278
	goto L55
L57:
	;
	goto L58
L58:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[6]))
	v286 = v280 + v263<<(uint(int32(13))%32) + int32(-8192)
	goto L55
L59:
	;
	v300 = F_brin_form_tuple(m, v69, v60, v259, v20+int32(20))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if base.Ui32(v293) < base.Ui32(v303) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	F_LockBuffer(m, v329, int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L7
	} else {
		goto L69
	}
L62:
	;
	if v302 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v328 = int32(1)
	goto L64
L64:
	;
	goto L61
L65:
	;
	v325 = F_PageGetExactFreeSpace(m, v324)
	mBase = m.M
	v328 = base.B2i32(base.Ui32(v303-v293) <= base.Ui32(v325))
	goto L64
L66:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[5]))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v310+(v302^int32(-1))<<(uint(int32(2))%32))))
	v324 = v316
	goto L65
L67:
	;
	goto L68
L68:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_brininsert[6]))
	v324 = v318 + v302<<(uint(int32(13))%32) + int32(-8192)
	goto L65
L69:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v336 = F_brin_doupdate(m, l0, v58, v70, v60, v333, v334, v296, v293, v300, v335, v328)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	if v336 != 0 {
		v343 = v257
		goto L42
	} else {
		goto L71
	}
L71:
	;
	F_MemoryContextReset(m, v257)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	v82 = v257
	goto L11
L73:
	;
	F_ReleaseBuffer(m, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L7
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brininsert[0])) = v30
	if v343 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	F_MemoryContextDelete(m, v343)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	m.G0 = v20 + int32(32)
	return int32(0)
L80:
	;
	goto L79
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int64
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
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
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 float64
	_ = v258
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
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
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
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
	var v362 int32
	_ = v362
	var v365 float64
	_ = v365
	var v369 float64
	_ = v369
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
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
		v44 = v7
		v45 = v29
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
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
	if base.Ui32(v45) <= base.Ui32(v44) {
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
		v44 = v35
		v45 = v38
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
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v54 = int32(0)
	v61 = v44
	v65 = v7
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
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v418 != 0 {
		goto L91
	} else {
		goto L92
	}
L19:
	;
	goto L18
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if base.Ui32(v45) < base.Ui32(v73+v61) {
		v401 = v54
		v412 = v65
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[0]))
	if v77 != 0 {
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
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v85 = F_brinGetTupleForHeapBlock(m, v26, v61, v22+int32(8), v22+int32(6), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v397 = v396 + v61
	if base.Ui32(v397) < base.Ui32(v45) {
		v54 = v379
		v61 = v397
		v65 = v390
		goto L17
	} else {
		goto L90
	}
L29:
	;
	if v85 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v54 == int32(0) {
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
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v93 = F_palloc(m, int32(80))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v130 = v54
	v133 = v65
	goto L35
L35:
	;
	v134 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v134
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v130)+44))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v147 = base.I32_div_s(v141<<(uint(int32(1))%32)+int32(7), int32(8))
	v149 = v147 + int32(12)
	v151 = v149 & int32(-8)
	v152 = F_palloc0(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L40
	}
L36:
	;
	v95 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+8)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v93)+40)) = v26
	v99 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v93)+16)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v93)+24)) = v99
	v106 = F_brin_build_desc(m, l0)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+44)) = v106
	v109 = F_brin_new_memtuple(m, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+72)) = int32(0)
	v113 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+64)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v93)+48)) = v109
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v93)+52)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v93)+60)) = v117
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v123 = base.I32_rem_u_s(int32(-2), v91)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v121 - v123 - int32(2)
	v128 = F_BuildIndexInfo(m, l0)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v130 = v93
	v133 = v128
	goto L35
L40:
	;
	v157 = v149&int32(24) | int32(-32)
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+4)) = uint8(v157)
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v61
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if int32(0) < v161 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v174 = int32(128)
	v175 = v157
	v181 = v152 + int32(4)
	v182 = v134
	goto L44
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(24)))) = v151
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v130)+28))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v130)+40))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v232 = F_brin_doinsert(m, v226, v227, v228, v22+int32(28), v61, v152, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
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
	v196 = v175
	v197 = v181
	v198 = v174 << (uint(int32(1)) % 32)
	goto L48
L47:
	;
	v190 = int32(0)
	v191 = int32(1)
	v192 = v181 + v191
	*(*uint8)(unsafe.Add(mBase, uint32(v192))) = uint8(v190)
	v196 = v190
	v197 = v192
	v198 = v191
	goto L48
L48:
	;
	v199 = v198 | v196
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v199)
	v202 = v182 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v202 < v204 {
		v174 = v198
		v175 = v199
		v181 = v197
		v182 = v202
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)) = uint16(v232)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v130)+28))
	if base.Ui32(v235+v61) <= base.Ui32(v45) {
		v248 = v235
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+32)) = v61
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v251 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+140))
	v258 = m.T0[v257].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l1, v250, v133, v251, int32(1), v251, v61, v248, int32(14), v130, v251)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L56
	}
L52:
	;
	v239 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v130)+28))
	if base.Ui32(v242) <= base.Ui32(v239-v61) {
		v248 = v242
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v245 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v248 = v245 - v61
	goto L51
L56:
	;
	v267 = v152
	goto L57
L57:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[0]))
	if v280 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_ReleaseBuffer(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L83
	}
L59:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v130)+44))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v130)+48))
	v287 = F_brin_form_tuple(m, v283, v61, v284, v22+int32(16))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if base.Ui32(v290) < base.Ui32(v291) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v130)+28))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v130)+40))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v324 = F_brin_doupdate(m, v317, v318, v319, v61, v320, v321, v267, v322, v287, v323, v316)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L72
	}
L65:
	;
	if v289 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v316 = int32(1)
	goto L67
L67:
	;
	goto L64
L68:
	;
	v313 = F_PageGetExactFreeSpace(m, v312)
	mBase = m.M
	v316 = base.B2i32(base.Ui32(v291-v290) <= base.Ui32(v313))
	goto L67
L69:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[2]))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v298+(v289^int32(-1))<<(uint(int32(2))%32))))
	v312 = v304
	goto L68
L70:
	;
	goto L71
L71:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_brinsummarize[3]))
	v312 = v306 + v289<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	F_pfree(m, v267)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_pfree(m, v287)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v324 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v130)+40))
	v339 = F_brinGetTupleForHeapBlock(m, v332, v61, v22+int32(28), v22+int32(22), v22+int32(24))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
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
	if v339 == int32(0) {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v344 = int32(0)
	v346 = F_brin_copy_tuple(m, v339, v343, v344, v344)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_LockBuffer(m, v348, int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v130)+44))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v130)+48))
	F_union_tuples(m, v352, v353, v346)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v267 = v346
	goto L57
L83:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v130)+48))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v130)+44))
	F_brin_memtuple_initialize(m, v359, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if l4 == int32(0) {
		v379 = v130
		v390 = v133
		goto L28
	} else {
		goto L85
	}
L85:
	;
	v365 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(v365, float64(1))
	v379 = v130
	v390 = v133
	goto L28
L86:
	;
	v369 = *(*float64)(unsafe.Add(mBase, uint32(l5)))
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = base.F64_add(v369, float64(1))
	goto L88
L87:
	;
	goto L88
L88:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	F_LockBuffer(m, v373, int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v379 = v54
	v390 = v65
	goto L28
L90:
	;
	v401 = v379
	v412 = v390
	goto L19
L91:
	;
	F_ReleaseBuffer(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
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
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	if v401 == int32(0) {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	F_terminate_brin_buildstate(m, v401)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_pfree(m, v412)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
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
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_brinsummarize_1), int32(1864), int32(_a_F_brinsummarize_2))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
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
	var v164 int32
	_ = v164
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = l0 + int32(28)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(1)
	v24 = v7 + v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v27 = v25 & v23
	if v25 == v23 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v157 != v7 {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	if v27 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v30 = int32(4)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v32&int32(254) == int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v45 = int32(1)
	if v27 != 0 {
		v55 = int32(base.Ui32(v25)>>(uint(v45)%32)) - v45
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v41 = v30
	goto L11
L10:
	;
	v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
	goto L11
L11:
	;
	if v32 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v30
	goto L14
L13:
	;
	v44 = v41
	goto L14
L14:
	;
	v55 = v44
	goto L5
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	v58 = v24
	goto L18
L17:
	;
	v58 = v7 + int32(4)
	goto L18
L18:
	;
	v64 = v55
	goto L19
L19:
	;
	if v64 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v81 = int32(1)
	v82 = v14 + v81
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v85 = v83 & v81
	if v83 == v81 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	v80 = v55 >> (uint(int32(31)) % 32) & v55
	goto L21
L23:
	;
	goto L24
L24:
	;
	v74 = v64 - int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v74))))
	if v76 == int32(32) {
		v64 = v74
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v80 = v64
	goto L21
L26:
	;
	if v85 != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v88 = int32(4)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v90&int32(254) == int32(2) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v103 = int32(1)
	if v85 != 0 {
		v113 = int32(base.Ui32(v83)>>(uint(v103)%32)) - v103
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v99 = v88
	goto L32
L31:
	;
	v99 = base.B2i32(v90 == int32(18)) << (uint(v88) % 32)
	goto L32
L32:
	;
	if v90 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = v88
	goto L35
L34:
	;
	v102 = v99
	goto L35
L35:
	;
	v113 = v102
	goto L26
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v113 = int32(base.Ui32(v107)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v116 = v82
	goto L39
L38:
	;
	v116 = v14 + int32(4)
	goto L39
L39:
	;
	v122 = v113
	goto L40
L40:
	;
	if v122 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v139 = int32(1)
	if v25&v139 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v138 = v113 >> (uint(int32(31)) % 32) & v113
	goto L42
L44:
	;
	goto L45
L45:
	;
	v132 = v122 - int32(1)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v132))))
	if v134 == int32(32) {
		v122 = v132
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v138 = v122
	goto L42
L47:
	;
	goto L4
L48:
	;
	v143 = v139
	goto L50
L49:
	;
	v143 = int32(4)
	goto L50
L50:
	;
	v145 = int32(1)
	if v83&v145 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v149 = v145
	goto L53
L52:
	;
	v149 = int32(4)
	goto L53
L53:
	;
	v151 = base.B2i32(v80 < v138)
	if v80 < v138 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v152 = v80
	goto L56
L55:
	;
	v152 = v138
	goto L56
L56:
	;
	v153 = F_memcmp(m, v7+v143, v14+v149, v152)
	mBase = m.M
	if v153 != 0 {
		v156 = v153
		goto L47
	} else {
		goto L57
	}
L57:
	;
	if v80 < v138 {
		v156 = int32(-1)
		goto L47
	} else {
		goto L58
	}
L58:
	;
	v156 = base.B2i32(v138 < v80)
	goto L47
L59:
	;
	F_pfree(m, v7)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v161 != v14 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v14)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	return v156
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
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
	var v117 int32
	_ = v117
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
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
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
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
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int64
	_ = v313
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
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
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
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
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
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
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
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
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
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
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
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
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 float64
	_ = v536
	var v537 int32
	_ = v537
	var v538 float64
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 float64
	_ = v581
	var v583 int32
	_ = v583
	var v587 float64
	_ = v587
	var v589 int32
	_ = v589
	var v607 float64
	_ = v607
	var v608 float64
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v614 int64
	_ = v614
	var v618 int64
	_ = v618
	var v625 int64
	_ = v625
	var v627 int64
	_ = v627
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v772 int64
	_ = v772
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v936 int32
	_ = v936
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
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
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v989 int32
	_ = v989
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
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1071 int64
	_ = v1071
	var v1079 int32
	_ = v1079
	var v1088 int32
	_ = v1088
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1130 int32
	_ = v1130
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1148 int32
	_ = v1148
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1212 int32
	_ = v1212
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int64
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1292 int64
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1316 int32
	_ = v1316
	var v1322 int32
	_ = v1322
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int64
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1364 int64
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1382 int64
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1462 int32
	_ = v1462
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1541 int32
	_ = v1541
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1558 int32
	_ = v1558
	var v1568 int32
	_ = v1568
	var v1575 int32
	_ = v1575
	var v1578 int64
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1602 int32
	_ = v1602
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1644 int64
	_ = v1644
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1662 int64
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1696 int32
	_ = v1696
	var v1701 int32
	_ = v1701
	var v1704 int64
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1728 int32
	_ = v1728
	var v1734 int32
	_ = v1734
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1780 int32
	_ = v1780
	var v1786 int32
	_ = v1786
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1862 int32
	_ = v1862
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1884 int32
	_ = v1884
	var v1888 int32
	_ = v1888
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1925 int32
	_ = v1925
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1940 int32
	_ = v1940
	var v1947 int32
	_ = v1947
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1960 int32
	_ = v1960
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2017 int32
	_ = v2017
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2082 int32
	_ = v2082
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2113 float64
	_ = v2113
	v4 = int32(0)
	v17 = int64(0)
	v20 = m.G0
	v22 = v20 - int32(96)
	m.G0 = v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+16)) = uint8(v24)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+117)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)) = uint8(v4)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)) = uint8(v26)
	v38 = F_RelationGetNumberOfBlocksInFork(m, l1, v4)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v2062 = F_smgr_bulk_get_buf(m, v2061)
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L5
	} else {
		goto L418
	}
L2:
	;
	v1801 = int32(0)
	v1804 = v1801
	v1805 = v1801
	v1807 = v1786
	goto L353
L3:
	;
	v1780 = int32(0)
	v2043 = v1780
	v2044 = v1780
	goto L1
L4:
	;
	F_pfree(m, v1328)
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L5
	} else {
		goto L352
	}
L5:
	;
	return int32(0)
L6:
	;
	if v38 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = F_palloc0(m, int32(16))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L5
	} else {
		goto L349
	}
L10:
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
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v89 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L11
L13:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v62 != int32(1) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v65 = int32(_a_F_btbuild_0)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v68 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v67 + v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v71 + v68
	*(*int64)(unsafe.Add(mBase, uint32(v58+int32(80))+232)) = int64(2)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v79 + v68
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v85 - v68
	goto L12
L15:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v465 != 0 {
		goto L116
	} else {
		goto L117
	}
L16:
	;
	v93 = v89 + int32(1)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	v96 = F_palloc0(m, int32(32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[3]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+72)) = v101 + int32(1)
	goto L18
L18:
	;
	v107 = F_CreateParallelContext(m, int32(_a_F_btbuild_1), v89)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	if v94 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v111 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	v115 = int32(_a_F_btbuild_2)
	goto L22
L22:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v118 = F_table_parallelscan_estimate(m, v117, v115)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L25
	}
L23:
	;
	v113 = F_RegisterSnapshot(m, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v115 = v113
	goto L22
L25:
	;
	v120 = F_add_size(m, int32(96), v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v107)+36))
	v127 = F_add_size(m, v122, (v120+int32(31))&int32(-32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+36)) = v127
	v130 = F_tuplesort_estimate_shared(m, v93)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v107)+36))
	v136 = (v130 + int32(31)) & int32(-32)
	v137 = F_add_size(m, v132, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+36)) = v137
	v140 = int32(1)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	if v142 == v140 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v145 = F_add_size(m, v137, v136)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v149 = int32(2)
	goto L32
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v107)+40))
	v151 = F_add_size(m, v150, v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+36)) = v145
	v149 = int32(3)
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+40)) = v151
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v107)+36))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v157 = F_mul_size(m, int32(32), v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v163 = F_add_size(m, v154, (v157+int32(31))&int32(-32))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+36)) = v163
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v107)+40))
	v168 = F_add_size(m, v166, int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+40)) = v168
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v107)+36))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v174 = F_mul_size(m, int32(128), v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v180 = F_add_size(m, v171, (v174+int32(31))&int32(-32))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+36)) = v180
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v107)+40))
	v185 = F_add_size(m, v183, int32(1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+40)) = v185
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	if v189 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v107)+36))
	if v189&int32(3) == int32(0) {
		v214 = v189
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v263 = v140
	goto L43
L43:
	;
	F_InitializeParallelDSM(m, v107)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L5
	} else {
		goto L63
	}
L44:
	;
	v252 = F_add_size(m, v190, v247&int32(-32)+int32(32))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L61
	}
L45:
	;
	v247 = v239 - v189
	goto L44
L46:
	;
	v218 = v214
	goto L55
L47:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v198 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v247 = int32(0)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v203 = v189
	goto L51
L51:
	;
	v207 = v203 + int32(1)
	if v207&int32(3) == int32(0) {
		v214 = v207
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v239 = v207
	goto L45
L53:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v212 != 0 {
		v203 = v207
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v227 = int32(-2139062144)
	if (int32(16843008)-v224|v224)&v227 == v227 {
		v218 = v218 + int32(4)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v233 = v218
	goto L58
L57:
	;
	goto L56
L58:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v237 != 0 {
		v233 = v233 + int32(1)
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v239 = v233
	goto L45
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+36)) = v252
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v107)+40))
	v257 = F_add_size(m, v255, int32(1))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+40)) = v257
	v263 = v247 + int32(1)
	goto L43
L63:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	if v266 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	switch v269 {
	case 0, 5:
		goto L68
	default:
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	v282 = F_shm_toc_allocate(m, v281, v120)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L72
	}
L67:
	;
	F_DestroyParallelContext(m, v107)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L70
	}
L68:
	;
	F_UnregisterSnapshot(m, v115)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[3]))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+72)) = v277 - int32(1)
	goto L71
L71:
	;
	goto L15
L72:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v288
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+8)) = uint8(v290)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = v93
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+10)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+9)) = uint8(v292)
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v298 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v282)+16)) = v303
	v306 = v282 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = int64(-4294967296)
	goto L77
L74:
	;
	v303 = int64(0)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v298)+392))
	v303 = v302
	goto L73
L77:
	;
	v311 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+72)) = uint8(v311)
	v313 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v282)+64)) = v313
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+56)) = uint8(v311)
	*(*int64)(unsafe.Add(mBase, uint32(v282)+48)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v282)+36)) = v313
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	F_table_parallelscan_initialize(m, v321, v282+int32(96), v115)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	v327 = F_shm_toc_allocate(m, v326, v130)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	F_tuplesort_initialize_shared(m, v327, v93, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	F_shm_toc_insert(m, v332, int64(-6917529027641081855), v282)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	F_shm_toc_insert(m, v336, int64(-6917529027641081854), v327)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	if v341 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	v345 = F_shm_toc_allocate(m, v344, v130)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L5
	} else {
		goto L86
	}
L84:
	;
	v354 = int32(0)
	goto L85
L85:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	if v356 != 0 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	F_tuplesort_initialize_shared(m, v345, v93, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	F_shm_toc_insert(m, v350, int64(-6917529027641081853), v345)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	v354 = v345
	goto L85
L89:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	v358 = F_shm_toc_allocate(m, v357, v263)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v372 = F_mul_size(m, int32(32), v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L98
	}
L92:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[4]))
	if v263 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	F_shm_toc_insert(m, v364, int64(-6917529027641081852), v363)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L97
	}
L94:
	;
	v362 = F__emscripten_memcpy_bulkmem(m, v358, v361, v263)
	mBase = m.M
	v363 = v362
	goto L96
L95:
	;
	v363 = v358
	goto L96
L96:
	;
	goto L93
L97:
	;
	goto L91
L98:
	;
	v374 = F_shm_toc_allocate(m, v369, v372)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	F_shm_toc_insert(m, v376, int64(-6917529027641081851), v374)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v383 = F_mul_size(m, int32(128), v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v385 = F_shm_toc_allocate(m, v380, v383)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	F_shm_toc_insert(m, v387, int64(-6917529027641081850), v385)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_LaunchParallelWorkers(m, v107)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+28)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v96)+24)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v96)+20)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v96)+16)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v394 + int32(1)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	if v404 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F__bt_end_parallel(m, v96)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v96
	v411 = F_palloc0(m, int32(16))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L109
	}
L108:
	;
	goto L15
L109:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v411)+4)) = v414
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v411)+8)) = v417
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v411)+12)) = uint8(v420)
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v411)+13)) = uint8(v422)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+8)))
	if v426 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v430 = F_palloc0(m, int32(16))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L113
	}
L111:
	;
	v439 = int32(0)
	v440 = v425
	goto L112
L112:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[5]))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v446 = base.I32_div_s(v444, v445)
	F__bt_parallel_scan_and_sort(m, v411, v439, v440, v441, v442, v446, int32(1))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L114
	}
L113:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v411)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v430)+4)) = v432
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v411)+8))
	v435 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+12)) = uint8(v435)
	*(*int32)(unsafe.Add(mBase, uint32(v430)+8)) = v434
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v439 = v430
	v440 = v438
	goto L112
L114:
	;
	F_WaitForParallelWorkersToAttach(m, v107)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	goto L15
L116:
	;
	v467 = F_palloc0(m, int32(12))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L5
	} else {
		goto L119
	}
L117:
	;
	v477 = int32(0)
	goto L118
L118:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+16)))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)))
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[5]))
	v482 = F_tuplesort_begin_index_btree(m, l0, l1, v478, v479, v481, v477)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L5
	} else {
		goto L120
	}
L119:
	;
	v469 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v467))) = uint8(v469)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v467)+4)) = v472
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v467)+8)) = v475
	v477 = v467
	goto L118
L120:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = v482
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	if v486 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v490 = F_palloc0(m, int32(16))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L5
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v522 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L124:
	;
	v492 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v490)+12)) = uint8(v492)
	*(*int32)(unsafe.Add(mBase, uint32(v490)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v490)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v490
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v498 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v500 = F_palloc0(m, int32(12))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L5
	} else {
		goto L128
	}
L126:
	;
	v511 = v492
	v512 = v490
	goto L127
L127:
	;
	v513 = int32(0)
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[6]))
	v517 = F_tuplesort_begin_index_btree(m, l0, l1, v513, v513, v516, v511)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L5
	} else {
		goto L129
	}
L128:
	;
	v502 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v500))) = uint8(v502)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v500)+4)) = v505
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v500)+8)) = v508
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v511 = v500
	v512 = v510
	goto L127
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v512))) = v517
	goto L123
L130:
	;
	v609 = int32(0)
	v611 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v611
	v614 = *(*int64)(unsafe.Add(mBase, _c_F_btbuild[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v614
	v618 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22-int32(-64)))) = v618
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v618
	if base.F64_lt(base.F64_abs(v607), float64(9.223372036854776e+18)) != 0 {
		goto L147
	} else {
		goto L148
	}
L131:
	;
	v525 = int32(1)
	v526 = int32(0)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v534)+140))
	v536 = m.T0[v535].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v525, v526, v525, v526, int32(-1), int32(238), v22+int32(16), v526)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L5
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	v543 = v539 + int32(36)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	goto L135
L134:
	;
	v538 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v607 = v538
	v608 = v536
	goto L130
L135:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = int32(1)
	if v564 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)) = uint8(v579)
	v581 = *(*float64)(unsafe.Add(mBase, uint32(v539)+64))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+32)) = v581
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+72)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)) = uint8(v583)
	*(*int32)(unsafe.Add(mBase, uint32(v539)+36)) = int32(0)
	v587 = *(*float64)(unsafe.Add(mBase, uint32(v539)+48))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L5
	} else {
		goto L145
	}
L137:
	;
	F_s_lock(m, v543, int32(_a_F_btbuild_3), int32(1664), int32(_a_F_btbuild_4))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L5
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v539)+40))
	if v544 != v572 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L139
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = int32(0)
	F_ConditionVariableSleep(m, v539+int32(24), int32(134217767))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L5
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	goto L136
L144:
	;
	goto L135
L145:
	;
	v607 = v581
	v608 = v587
	goto L130
L146:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v627
	v634 = int32(0)
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v641 == v634 {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	v625 = base.I64_trunc_f64_s(v607)
	v627 = v625
	goto L146
L148:
	;
	goto L149
L149:
	;
	v627 = int64(-9223372036854775807 - 1)
	goto L146
L150:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v807 == int32(0) {
		v818 = v609
		goto L167
	} else {
		goto L168
	}
L151:
	;
	goto L150
L152:
	;
	goto L153
L153:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v647&int32(1) == int32(0) {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v652 = int32(_a_F_btbuild_0)
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v655 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v654 + v655
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	*(*int32)(unsafe.Add(mBase, uint32(v641))) = v658 + v655
	goto L156
L155:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	v789 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v641))) = v788 + v789
	v792 = int32(_a_F_btbuild_0)
	v794 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v794 - v789
	goto L151
L156:
	;
	goto L158
L158:
	;
	goto L159
L159:
	;
	goto L163
L163:
	;
	v753 = int32(0)
	v756 = v634
	goto L164
L164:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(80)+v756<<(uint(int32(2))%32))))
	v766 = int32(3)
	v772 = *(*int64)(unsafe.Add(mBase, uint32(v22+int32(48)+v756<<(uint(v766)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v641+int32(232)+v765<<(uint(v766)%32)))) = v772
	v774 = int32(1)
	v777 = v753 + v774
	if v777 != int32(3) {
		v753 = v777
		v756 = v756 + v774
		goto L164
	} else {
		goto L166
	}
L165:
	;
	goto L155
L166:
	;
	goto L165
L167:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v824 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v824 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L168:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)))
	if v810 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v818 = v807
	goto L167
L170:
	;
	goto L171
L171:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v807)))
	F_tuplesort_end(m, v811)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L5
	} else {
		goto L172
	}
L172:
	;
	F_pfree(m, v807)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = int32(0)
	v818 = v609
	goto L167
L174:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	F_tuplesort_performsort(m, v855)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L5
	} else {
		goto L178
	}
L175:
	;
	goto L174
L176:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v828 != int32(1) {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v831 = int32(_a_F_btbuild_0)
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v834 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v833 + v834
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v837 + v834
	*(*int64)(unsafe.Add(mBase, uint32(v824+int32(80))+232)) = int64(3)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v845 + v834
	v851 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v851 - v834
	goto L175
L178:
	;
	if v818 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v862 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v862 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	goto L181
L181:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v896
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v819)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v898
	v900 = int32(0)
	v902 = F__bt_mkscankey(m, v898, v900)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L5
	} else {
		goto L187
	}
L182:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	F_tuplesort_performsort(m, v893)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L5
	} else {
		goto L186
	}
L183:
	;
	goto L182
L184:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v866 != int32(1) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v869 = int32(_a_F_btbuild_0)
	v871 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v872 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v871 + v872
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v875 + v872
	*(*int64)(unsafe.Add(mBase, uint32(v862+int32(80))+232)) = int64(4)
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v883 + v872
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v889 - v872
	goto L183
L186:
	;
	goto L181
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v902
	v906 = F__bt_allequalimage(m, v898, int32(1))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v902)+1)) = uint8(v906)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(1)
	v915 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v915 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v898)+52))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v898)+192))
	v948 = int32(*(*int16)(unsafe.Add(mBase, uint32(v947)+10)))
	v950 = F_smgr_bulk_start_rel(m, v898, int32(0))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L5
	} else {
		goto L193
	}
L190:
	;
	goto L189
L191:
	;
	v919 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v919 != int32(1) {
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v922 = int32(_a_F_btbuild_0)
	v924 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v925 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v924 + v925
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	*(*int32)(unsafe.Add(mBase, uint32(v915))) = v928 + v925
	*(*int64)(unsafe.Add(mBase, uint32(v915+int32(80))+232)) = int64(5)
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	*(*int32)(unsafe.Add(mBase, uint32(v915))) = v936 + v925
	v942 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v942 - v925
	goto L190
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v950
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902)+1)))
	if v953 != int32(1) {
		v961 = v900
		goto L199
	} else {
		goto L200
	}
L194:
	;
	F_pfree(m, v978)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L5
	} else {
		goto L347
	}
L195:
	;
	v1632 = int32(0)
	v1634 = v964
	v1644 = v17
	goto L329
L196:
	;
	v1328 = F_palloc(m, int32(1676))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L5
	} else {
		goto L277
	}
L197:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	v971 = F_tuplesort_getheaptuple(m, v970)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L5
	} else {
		goto L208
	}
L198:
	;
	if v818 == int32(0) {
		goto L196
	} else {
		goto L207
	}
L199:
	;
	if v818 != 0 {
		goto L197
	} else {
		goto L203
	}
L200:
	;
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v819)+12)))
	if v956 != 0 {
		v961 = v900
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v898)+180))
	if v957 == int32(0) {
		goto L198
	} else {
		goto L202
	}
L202:
	;
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v957)+16)))
	v961 = v960
	goto L199
L203:
	;
	if v961 != 0 {
		goto L196
	} else {
		goto L204
	}
L204:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	v964 = F_tuplesort_getheaptuple(m, v963)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L5
	} else {
		goto L205
	}
L205:
	;
	if v964 != 0 {
		goto L195
	} else {
		goto L206
	}
L206:
	;
	v2043 = int32(0)
	v2044 = int32(0)
	goto L1
L207:
	;
	goto L197
L208:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	v974 = F_tuplesort_getheaptuple(m, v973)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v978 = F_palloc0(m, v948*int32(36))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L5
	} else {
		goto L210
	}
L210:
	;
	if int32(0) < v948 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v989 = int32(0)
	goto L214
L212:
	;
	goto L213
L213:
	;
	v1057 = v971
	v1059 = int32(0)
	v1063 = v974
	v1071 = v17
	goto L218
L214:
	;
	v1006 = v978 + v989*int32(36)
	v1008 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1006))) = v1008
	v1012 = v902 + int32(16) + v989*int32(48)
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
	F_PrepareSortSupportFromIndexRel(m, v898, int32(base.Ui32(v1025&int32(16777216))>>(uint(int32(24))%32)), v1006)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L5
	} else {
		goto L216
	}
L215:
	;
	goto L213
L216:
	;
	v1033 = v989 + int32(1)
	if v1033 != v948 {
		v989 = v1033
		goto L214
	} else {
		goto L217
	}
L217:
	;
	goto L215
L218:
	;
	if v1063 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	if v1059 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L221:
	;
	if v1057 == int32(0) {
		goto L194
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v1079 = int32(0)
	if v1057 == v1079 {
		v1212 = v1079
		goto L220
	} else {
		goto L225
	}
L224:
	;
	v1212 = int32(1)
	goto L220
L225:
	;
	if int32(0) < v948 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1088 = int32(1)
	goto L229
L227:
	;
	goto L228
L228:
	;
	v1180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1057)+2)))
	v1181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1057))))
	v1182 = int32(16)
	v1184 = v1180 | v1181<<(uint(v1182)%32)
	v1185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063)+2)))
	v1186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063))))
	v1189 = v1185 | v1186<<(uint(v1182)%32)
	if base.Ui32(v1184) < base.Ui32(v1189) {
		v1200 = int32(-1)
		goto L252
	} else {
		goto L253
	}
L229:
	;
	v1106 = v978 + v1088*int32(36)
	v1109 = F_index_getattr_2(m, v1057, v1088, v946, v22+int32(80))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L5
	} else {
		goto L231
	}
L230:
	;
	goto L228
L231:
	;
	v1113 = F_index_getattr_2(m, v1063, v1088, v946, v22+int32(95))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L5
	} else {
		goto L232
	}
L232:
	;
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+95)))
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+80)))
	if v1116 == int32(1) {
		goto L236
	} else {
		goto L237
	}
L233:
	;
	if v1088 != v948 {
		v1088 = v1088 + int32(1)
		goto L229
	} else {
		goto L250
	}
L234:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1106-int32(20))))
	v1137 = m.T0[v1136].(func(*base.Module, int32, int32, int32) int32)(m, v1109, v1113, v1106-int32(36))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L5
	} else {
		goto L243
	}
L235:
	;
	v1212 = int32(1)
	goto L220
L236:
	;
	if v1115&int32(1) != 0 {
		goto L233
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	if v1115&int32(1) == int32(0) {
		goto L234
	} else {
		goto L241
	}
L239:
	;
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1106-int32(27)))))
	if v1123 != 0 {
		goto L235
	} else {
		goto L240
	}
L240:
	;
	v1212 = v1079
	goto L220
L241:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1106-int32(27)))))
	if v1130 != 0 {
		v1212 = v1079
		goto L220
	} else {
		goto L242
	}
L242:
	;
	goto L235
L243:
	;
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1106-int32(28)))))
	if v1141 == int32(1) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	if v1137 < int32(0) {
		v1212 = v1079
		goto L220
	} else {
		goto L247
	}
L245:
	;
	v1148 = v1137
	goto L246
L246:
	;
	if int32(0) < v1148 {
		v1212 = v1079
		goto L220
	} else {
		goto L248
	}
L247:
	;
	v1148 = int32(0) - v1137
	goto L246
L248:
	;
	if v1148 == int32(0) {
		goto L233
	} else {
		goto L249
	}
L249:
	;
	v1212 = int32(1)
	goto L220
L250:
	;
	goto L230
L251:
	;
	v1212 = base.B2i32(v1200 <= int32(0))
	goto L220
L252:
	;
	goto L251
L253:
	;
	if base.Ui32(v1189) < base.Ui32(v1184) {
		v1200 = int32(1)
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v1194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1057)+4)))
	v1195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063)+4)))
	if base.Ui32(v1194) < base.Ui32(v1195) {
		v1200 = int32(-1)
		goto L252
	} else {
		goto L255
	}
L255:
	;
	v1200 = base.B2i32(base.Ui32(v1195) < base.Ui32(v1194))
	goto L252
L256:
	;
	v1225 = F_palloc0(m, int32(32))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L5
	} else {
		goto L259
	}
L257:
	;
	v1270 = v1059
	goto L258
L258:
	;
	if v1212 != 0 {
		goto L266
	} else {
		goto L267
	}
L259:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1228 = F_smgr_bulk_get_buf(m, v1227)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	F_PageInit(m, v1228, int32(_a_F_btbuild_5), int32(16))
	mBase = m.M
	goto L261
L261:
	;
	v1233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1228)+16)))
	v1234 = v1228 + v1233
	*(*int64)(unsafe.Add(mBase, uint32(v1234)+8)) = int64(4294967296)
	v1237 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1234))) = v1237
	v1239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1228)+12)))
	v1241 = v1239 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1228)+12)) = uint16(v1241)
	*(*int32)(unsafe.Add(mBase, uint32(v1225))) = v1228
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1245 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1244 + v1245
	*(*int64)(unsafe.Add(mBase, uint32(v1225)+16)) = v1237
	*(*uint16)(unsafe.Add(mBase, uint32(v1225)+12)) = uint16(v1245)
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+4)) = v1244
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+180))
	if v1256 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+4))
	v1263 = base.I32_div_s(int32(_a_F_btbuild_6)-v1258<<(uint(int32(13))%32), int32(100))
	v1265 = v1263
	goto L264
L263:
	;
	v1265 = int32(819)
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+24)) = v1265
	v1270 = v1225
	goto L258
L265:
	;
	v1292 = v1071 + int64(1)
	v1295 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1295 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L266:
	;
	F__bt_buildadd(m, v22+int32(48), v1270, v1057, int32(0))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L5
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	F__bt_buildadd(m, v22+int32(48), v1270, v1063, int32(0))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L5
	} else {
		goto L271
	}
L269:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	v1278 = F_tuplesort_getheaptuple(m, v1277)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L5
	} else {
		goto L270
	}
L270:
	;
	v1288 = v1278
	v1289 = v1063
	goto L265
L271:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	v1286 = F_tuplesort_getheaptuple(m, v1285)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L5
	} else {
		goto L272
	}
L272:
	;
	v1288 = v1057
	v1289 = v1286
	goto L265
L273:
	;
	v1057 = v1288
	v1059 = v1270
	v1063 = v1289
	v1071 = v1292
	goto L218
L274:
	;
	goto L273
L275:
	;
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1299 != int32(1) {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1302 = int32(_a_F_btbuild_0)
	v1304 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1305 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1304 + v1305
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	*(*int32)(unsafe.Add(mBase, uint32(v1295))) = v1308 + v1305
	*(*int64)(unsafe.Add(mBase, uint32(v1295+int32(96))+232)) = v1292
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	*(*int32)(unsafe.Add(mBase, uint32(v1295))) = v1316 + v1305
	v1322 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1322 - v1305
	goto L274
L277:
	;
	v1330 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1328)+4)) = v1330
	v1332 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1328))) = uint8(v1332)
	*(*int64)(unsafe.Add(mBase, uint32(v1328)+20)) = v1330
	*(*int64)(unsafe.Add(mBase, uint32(v1328)+10)) = v1330
	*(*int64)(unsafe.Add(mBase, uint32(v1328)+28)) = v1330
	*(*int64)(unsafe.Add(mBase, uint32(v1328)+36)) = v1330
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	v1343 = F_tuplesort_getheaptuple(m, v1342)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L5
	} else {
		goto L278
	}
L278:
	;
	if v1343 == int32(0) {
		goto L4
	} else {
		goto L279
	}
L279:
	;
	v1352 = int32(0)
	v1354 = v1343
	v1364 = v17
	goto L280
L280:
	;
	if v1352 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L281:
	;
	F__bt_sort_dedup_finish_pending(m, v22+int32(48), v1575, v1328)
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L5
	} else {
		goto L325
	}
L282:
	;
	v1578 = v1364 + int64(1)
	v1581 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1581 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L283:
	;
	v1512 = F_CopyIndexTuple(m, v1354)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L5
	} else {
		goto L311
	}
L284:
	;
	v1370 = F_palloc0(m, int32(32))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L5
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+12))
	v1422 = F__bt_keep_natts_fast(m, v1420, v1421, v1354)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L5
	} else {
		goto L294
	}
L287:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1373 = F_smgr_bulk_get_buf(m, v1372)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L5
	} else {
		goto L288
	}
L288:
	;
	F_PageInit(m, v1373, int32(_a_F_btbuild_5), int32(16))
	mBase = m.M
	goto L289
L289:
	;
	v1378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1373)+16)))
	v1379 = v1373 + v1378
	*(*int64)(unsafe.Add(mBase, uint32(v1379)+8)) = int64(4294967296)
	v1382 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1379))) = v1382
	v1384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1373)+12)))
	v1386 = v1384 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1373)+12)) = uint16(v1386)
	*(*int32)(unsafe.Add(mBase, uint32(v1370))) = v1373
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1390 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1389 + v1390
	*(*int64)(unsafe.Add(mBase, uint32(v1370)+16)) = v1382
	*(*uint16)(unsafe.Add(mBase, uint32(v1370)+12)) = uint16(v1390)
	*(*int32)(unsafe.Add(mBase, uint32(v1370)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1370)+4)) = v1389
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+180))
	if v1401 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+4))
	v1408 = base.I32_div_s(int32(_a_F_btbuild_6)-v1403<<(uint(int32(13))%32), int32(100))
	v1410 = v1408
	goto L292
L291:
	;
	v1410 = int32(819)
	goto L292
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1370)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1370)+24)) = v1410
	v1414 = int32(812)
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+8)) = v1414
	v1417 = F_palloc(m, v1414)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L5
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+24)) = v1417
	v1511 = v1370
	goto L283
L294:
	;
	if v948 < v1422 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1429 = int32(1)
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1354)+7)))
	if v1430&int32(32) == int32(0) {
		v1448 = v1429
		v1450 = v1354
		goto L299
	} else {
		goto L300
	}
L296:
	;
	goto L297
L297:
	;
	F__bt_sort_dedup_finish_pending(m, v22+int32(48), v1352, v1328)
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L5
	} else {
		goto L309
	}
L298:
	;
	if base.Ui32((v1452+(v1453+v1448)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1451) {
		v1575 = v1352
		goto L282
	} else {
		goto L308
	}
L299:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+8))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+20))
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+28))
	v1462 = base.B2i32(base.Ui32((v1452+(v1453+v1448)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1451))
	if v1462 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	v1435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1354)+4)))
	if v1435&int32(_a_F_btbuild_5) == int32(0) {
		v1448 = v1429
		v1450 = v1354
		goto L299
	} else {
		goto L301
	}
L301:
	;
	v1442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1354)+2)))
	v1443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1354))))
	v1448 = v1435 & int32(4095)
	v1450 = v1354 + (v1442 | v1443<<(uint(int32(16))%32))
	goto L299
L302:
	;
	goto L298
L303:
	;
	v1495 = v1328 + v1493
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1495)))
	*(*int32)(unsafe.Add(mBase, uint32(v1495))) = v1496 + v1494
	goto L302
L304:
	;
	if v1453 <= int32(50) {
		goto L302
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+32)) = v1469 + int32(1)
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+24))
	v1474 = int32(6)
	v1479 = F___memcpy(m, v1473+v1453*v1474, v1450, v1448*v1474)
	mBase = m.M
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+28)) = v1480 + v1448
	v1484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1354)+6)))
	v1493 = int32(36)
	v1494 = (v1484&int32(_a_F_btbuild_7)+int32(7))&int32(_a_F_btbuild_8) | int32(4)
	goto L303
L307:
	;
	v1493 = int32(4)
	v1494 = int32(1)
	goto L303
L308:
	;
	goto L297
L309:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+12))
	F_pfree(m, v1506)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L5
	} else {
		goto L310
	}
L310:
	;
	v1511 = v1352
	goto L283
L311:
	;
	v1514 = int32(0)
	v1516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512)+7)))
	if v1516&int32(32) != 0 {
		goto L315
	} else {
		goto L316
	}
L312:
	;
	v1575 = v1511
	goto L282
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+20)) = v1552
	*(*uint16)(unsafe.Add(mBase, uint32(v1328)+16)) = uint16(v1514)
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+12)) = v1512
	v1558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1512)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+36)) = (v1558&int32(_a_F_btbuild_7)+int32(7))&int32(_a_F_btbuild_8) | int32(4)
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1328+v1568<<(uint(int32(2))%32))+44)) = uint16(v1514)
	goto L312
L314:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+24))
	v1534 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1512)+2)))
	v1535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1512))))
	v1536 = int32(16)
	v1541 = v1519 & int32(4095)
	v1544 = F___memcpy(m, v1533, v1512+(v1534|v1535<<(uint(v1536)%32)), v1541*int32(6))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+28)) = v1541
	v1546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1512)+2)))
	v1547 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1512))))
	v1552 = v1546 | v1547<<(uint(v1536)%32)
	goto L313
L315:
	;
	v1519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1512)+4)))
	if v1519&int32(_a_F_btbuild_5) != 0 {
		goto L314
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+24))
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1512)))
	*(*int32)(unsafe.Add(mBase, uint32(v1523))) = v1524
	v1526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1512)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1523)+4)) = uint16(v1526)
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+28)) = int32(1)
	v1530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1512)+6)))
	v1552 = v1530 & int32(_a_F_btbuild_7)
	goto L313
L318:
	;
	goto L317
L319:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	v1613 = F_tuplesort_getheaptuple(m, v1612)
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L5
	} else {
		goto L323
	}
L320:
	;
	goto L319
L321:
	;
	v1585 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1585 != int32(1) {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1588 = int32(_a_F_btbuild_0)
	v1590 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1591 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1590 + v1591
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1581)))
	*(*int32)(unsafe.Add(mBase, uint32(v1581))) = v1594 + v1591
	*(*int64)(unsafe.Add(mBase, uint32(v1581+int32(96))+232)) = v1578
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1581)))
	*(*int32)(unsafe.Add(mBase, uint32(v1581))) = v1602 + v1591
	v1608 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1608 - v1591
	goto L320
L323:
	;
	if v1613 != 0 {
		v1352 = v1575
		v1354 = v1613
		v1364 = v1578
		goto L280
	} else {
		goto L324
	}
L324:
	;
	goto L281
L325:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+12))
	F_pfree(m, v1619)
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L5
	} else {
		goto L326
	}
L326:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+24))
	F_pfree(m, v1622)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L5
	} else {
		goto L327
	}
L327:
	;
	F_pfree(m, v1328)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L5
	} else {
		goto L328
	}
L328:
	;
	v1786 = v1575
	goto L2
L329:
	;
	if v1632 == int32(0) {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1786 = v1696
	goto L2
L331:
	;
	v1650 = F_palloc0(m, int32(32))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L5
	} else {
		goto L334
	}
L332:
	;
	v1696 = v1632
	goto L333
L333:
	;
	F__bt_buildadd(m, v22+int32(48), v1696, v1634, int32(0))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L5
	} else {
		goto L340
	}
L334:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1653 = F_smgr_bulk_get_buf(m, v1652)
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	F_PageInit(m, v1653, int32(_a_F_btbuild_5), int32(16))
	mBase = m.M
	goto L336
L336:
	;
	v1658 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1653)+16)))
	v1659 = v1653 + v1658
	*(*int64)(unsafe.Add(mBase, uint32(v1659)+8)) = int64(4294967296)
	v1662 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1659))) = v1662
	v1664 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1653)+12)))
	v1666 = v1664 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1653)+12)) = uint16(v1666)
	*(*int32)(unsafe.Add(mBase, uint32(v1650))) = v1653
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1670 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1669 + v1670
	*(*int64)(unsafe.Add(mBase, uint32(v1650)+16)) = v1662
	*(*uint16)(unsafe.Add(mBase, uint32(v1650)+12)) = uint16(v1670)
	*(*int32)(unsafe.Add(mBase, uint32(v1650)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1650)+4)) = v1669
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+180))
	if v1681 != 0 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+4))
	v1688 = base.I32_div_s(int32(_a_F_btbuild_6)-v1683<<(uint(int32(13))%32), int32(100))
	v1690 = v1688
	goto L339
L338:
	;
	v1690 = int32(819)
	goto L339
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1650)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1650)+24)) = v1690
	v1696 = v1650
	goto L333
L340:
	;
	v1704 = v1644 + int64(1)
	v1707 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[0]))
	if v1707 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	v1739 = F_tuplesort_getheaptuple(m, v1738)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L5
	} else {
		goto L345
	}
L342:
	;
	goto L341
L343:
	;
	v1711 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btbuild[1])))
	if v1711 != int32(1) {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1714 = int32(_a_F_btbuild_0)
	v1716 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	v1717 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1716 + v1717
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1707)))
	*(*int32)(unsafe.Add(mBase, uint32(v1707))) = v1720 + v1717
	*(*int64)(unsafe.Add(mBase, uint32(v1707+int32(96))+232)) = v1704
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1707)))
	*(*int32)(unsafe.Add(mBase, uint32(v1707))) = v1728 + v1717
	v1734 = *(*int32)(unsafe.Add(mBase, _c_F_btbuild[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btbuild[2])) = v1734 - v1717
	goto L342
L345:
	;
	if v1739 != 0 {
		v1632 = v1696
		v1634 = v1739
		v1644 = v1704
		goto L329
	} else {
		goto L346
	}
L346:
	;
	goto L330
L347:
	;
	if v1059 != 0 {
		v1786 = v1059
		goto L2
	} else {
		goto L348
	}
L348:
	;
	goto L3
L349:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v1747 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btbuild_9), v22)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L5
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_btbuild_3), int32(321), int32(_a_F_btbuild_10))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L5
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	goto L3
L353:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+4))
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+28))
	if v1823 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L354:
	;
	v2043 = v1850
	v2044 = v1851
	goto L1
L355:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1807)))
	v1854 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1853)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1854) {
		goto L361
	} else {
		goto L362
	}
L356:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1807)))
	v1827 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1826)+16)))
	v1828 = v1826 + v1827
	v1829 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1828)+12)))
	v1831 = v1829 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1828)+12)) = uint16(v1831)
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+20))
	v1850 = v1833
	v1851 = v1822
	goto L355
L357:
	;
	goto L358
L358:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1834))) = base.I32_rotr(v1822, int32(16))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+28))
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+8))
	F__bt_buildadd(m, v22+int32(48), v1840, v1841, int32(0))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L5
	} else {
		goto L359
	}
L359:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+8))
	F_pfree(m, v1845)
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L5
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1807)+8)) = int32(0)
	v1850 = v1804
	v1851 = v1805
	goto L355
L361:
	;
	v1862 = int32(base.Ui32(v1854+int32(_a_F_btbuild_11)) >> (uint(int32(2)) % 32))
	goto L363
L362:
	;
	v1862 = int32(0)
	goto L363
L363:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1862&int32(_a_F_btbuild_12)) {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1868 = v1853 + int32(24)
	v1870 = v1853 + int32(28)
	v1871 = int32(3)
	v1875 = (v1862 + int32(1)) & int32(_a_F_btbuild_12)
	if base.Ui32(v1875) <= base.Ui32(v1871) {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	goto L366
L366:
	;
	v2031 = v1854 - int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1853)+12)) = uint16(v2031)
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+4))
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1807)))
	F_smgr_bulk_write(m, v2033, v2034, v2035, int32(1))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L5
	} else {
		goto L416
	}
L367:
	;
	v1878 = v1871
	goto L369
L368:
	;
	v1878 = v1875
	goto L369
L369:
	;
	v1879 = int32(2)
	v1884 = (v1878 - v1879) & int32(_a_F_btbuild_12) << (uint(v1879) % 32)
	if v1868 == v1870 {
		goto L371
	} else {
		goto L372
	}
L370:
	;
	goto L366
L371:
	;
	goto L370
L372:
	;
	v1888 = v1868 + v1884
	if base.Ui32(v1870-v1888) <= base.Ui32(int32(0)-v1884<<(uint(int32(1))%32)) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1895 = F___memcpy(m, v1868, v1870, v1884)
	mBase = m.M
	goto L370
L374:
	;
	goto L375
L375:
	;
	v1898 = (v1868 ^ v1870) & int32(3)
	if base.Ui32(v1868) < base.Ui32(v1870) {
		goto L378
	} else {
		goto L379
	}
L376:
	;
	if v2000 == int32(0) {
		goto L371
	} else {
		goto L412
	}
L377:
	;
	if base.Ui32(v1978) <= base.Ui32(int32(3)) {
		v1999 = v1977
		v2000 = v1978
		v2001 = v1979
		goto L376
	} else {
		goto L408
	}
L378:
	;
	if v1898 != 0 {
		goto L381
	} else {
		goto L382
	}
L379:
	;
	goto L380
L380:
	;
	if v1898 != 0 {
		v1960 = v1884
		goto L391
	} else {
		goto L392
	}
L381:
	;
	v1999 = v1870
	v2000 = v1884
	v2001 = v1868
	goto L376
L382:
	;
	goto L383
L383:
	;
	if v1868&int32(3) == int32(0) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1977 = v1870
	v1978 = v1884
	v1979 = v1868
	goto L377
L385:
	;
	goto L386
L386:
	;
	v1905 = v1870
	v1906 = v1884
	v1907 = v1868
	goto L387
L387:
	;
	if v1906 == int32(0) {
		goto L371
	} else {
		goto L389
	}
L388:
	;
	v1977 = v1914
	v1978 = v1916
	v1979 = v1918
	goto L377
L389:
	;
	v1911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1905))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1907))) = uint8(v1911)
	v1913 = int32(1)
	v1914 = v1905 + v1913
	v1916 = v1906 - v1913
	v1918 = v1907 + v1913
	if v1918&int32(3) != 0 {
		v1905 = v1914
		v1906 = v1916
		v1907 = v1918
		goto L387
	} else {
		goto L390
	}
L390:
	;
	goto L388
L391:
	;
	if v1960 == int32(0) {
		goto L371
	} else {
		goto L404
	}
L392:
	;
	if v1888&int32(3) != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1925 = v1884
	goto L396
L394:
	;
	v1940 = v1884
	goto L395
L395:
	;
	if base.Ui32(v1940) <= base.Ui32(int32(3)) {
		v1960 = v1940
		goto L391
	} else {
		goto L400
	}
L396:
	;
	if v1925 == int32(0) {
		goto L371
	} else {
		goto L398
	}
L397:
	;
	v1940 = v1931
	goto L395
L398:
	;
	v1931 = v1925 - int32(1)
	v1932 = v1868 + v1931
	v1934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1870+v1931))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1932))) = uint8(v1934)
	if v1932&int32(3) != 0 {
		v1925 = v1931
		goto L396
	} else {
		goto L399
	}
L399:
	;
	goto L397
L400:
	;
	v1947 = v1940
	goto L401
L401:
	;
	v1951 = v1947 - int32(4)
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1870+v1951)))
	*(*int32)(unsafe.Add(mBase, uint32(v1868+v1951))) = v1954
	if base.Ui32(int32(3)) < base.Ui32(v1951) {
		v1947 = v1951
		goto L401
	} else {
		goto L403
	}
L402:
	;
	v1960 = v1951
	goto L391
L403:
	;
	goto L402
L404:
	;
	v1967 = v1960
	goto L405
L405:
	;
	v1971 = v1967 - int32(1)
	v1974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1870+v1971))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1868+v1971))) = uint8(v1974)
	if v1971 != 0 {
		v1967 = v1971
		goto L405
	} else {
		goto L407
	}
L406:
	;
	goto L371
L407:
	;
	goto L406
L408:
	;
	v1984 = v1977
	v1985 = v1978
	v1986 = v1979
	goto L409
L409:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1984)))
	*(*int32)(unsafe.Add(mBase, uint32(v1986))) = v1988
	v1990 = int32(4)
	v1991 = v1984 + v1990
	v1993 = v1986 + v1990
	v1995 = v1985 - v1990
	if base.Ui32(int32(3)) < base.Ui32(v1995) {
		v1984 = v1991
		v1985 = v1995
		v1986 = v1993
		goto L409
	} else {
		goto L411
	}
L410:
	;
	v1999 = v1991
	v2000 = v1995
	v2001 = v1993
	goto L376
L411:
	;
	goto L410
L412:
	;
	v2006 = v1999
	v2007 = v2000
	v2008 = v2001
	goto L413
L413:
	;
	v2010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2006))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2008))) = uint8(v2010)
	v2012 = int32(1)
	v2017 = v2007 - v2012
	if v2017 != 0 {
		v2006 = v2006 + v2012
		v2007 = v2017
		v2008 = v2008 + v2012
		goto L413
	} else {
		goto L415
	}
L414:
	;
	goto L371
L415:
	;
	goto L414
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1807))) = int32(0)
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+28))
	if v2041 != 0 {
		v1804 = v1850
		v1805 = v1851
		v1807 = v2041
		goto L353
	} else {
		goto L417
	}
L417:
	;
	goto L354
L418:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	v2065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2064)+1)))
	F_PageInit(m, v2062, int32(_a_F_btbuild_5), int32(16))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v2062-int32(-64)))) = uint8(v2065)
	*(*int64)(unsafe.Add(mBase, uint32(v2062)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v2062)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2062)+44)) = v2043
	*(*int32)(unsafe.Add(mBase, uint32(v2062)+40)) = v2044
	*(*int32)(unsafe.Add(mBase, uint32(v2062)+36)) = v2043
	*(*int32)(unsafe.Add(mBase, uint32(v2062)+32)) = v2044
	*(*int64)(unsafe.Add(mBase, uint32(v2062)+24)) = int64(17180209506)
	v2082 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2062)+16)))
	v2084 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v2062+v2082)+12)) = uint16(v2084)
	v2086 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v2062)+12)) = uint16(v2086)
	goto L419
L419:
	;
	F_smgr_bulk_write(m, v2061, int32(0), v2062, int32(1))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L5
	} else {
		goto L420
	}
L420:
	;
	F_smgr_bulk_finish(m, v2061)
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L5
	} else {
		goto L421
	}
L421:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2094)))
	F_tuplesort_end(m, v2095)
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L5
	} else {
		goto L422
	}
L422:
	;
	F_pfree(m, v2094)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L5
	} else {
		goto L423
	}
L423:
	;
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v2100 != 0 {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v2100)))
	F_tuplesort_end(m, v2101)
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L5
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v2106 != 0 {
		goto L429
	} else {
		goto L430
	}
L427:
	;
	F_pfree(m, v2100)
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L5
	} else {
		goto L428
	}
L428:
	;
	goto L426
L429:
	;
	F__bt_end_parallel(m, v2106)
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L5
	} else {
		goto L432
	}
L430:
	;
	goto L431
L431:
	;
	v2110 = F_palloc(m, int32(16))
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L5
	} else {
		goto L433
	}
L432:
	;
	goto L431
L433:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v2110))) = v608
	v2113 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2110)+8)) = v2113
	m.G0 = v22 + int32(96)
	return v2110
}
func F_btbuildphasename(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = l0 - int64(1)
	if base.Ui64(v4) <= base.Ui64(int64(4)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)<<(uint(int32(2))%32))+uint32(_c_F_btbuildphasename[0])))
		v13 = v12
	} else {
		v13 = int32(0)
	}
	return v13
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
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	v6 = int32(2147483647)
	v7 = l0 & v6
	v8 = base.F32_reinterpret_i32(l1)
	v9 = base.F32_reinterpret_i32(l0)
	v11 = l1 & v6
	if base.Ui32(int32(2139095041)) <= base.Ui32(v11) {
		v20 = base.B2i32(base.Ui32(v7) < base.Ui32(int32(2139095041)))
		v28 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v11))|base.F32_gt(v8, v9))&v20
	} else {
		v16 = int32(1)
		if base.F32_lt(v8, v9) != 0 {
			v28 = v16
		} else {
			if base.Ui32(int32(2139095040)) < base.Ui32(v7) {
				v28 = v16
			} else {
				v20 = v16
				v28 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v11))|base.F32_gt(v8, v9))&v20
			}
		}
	}
	return v28
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
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v7) & v9
	v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = base.F64_promote_f32(v11)
	v15 = base.I64_reinterpret_f64(v12) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v15) {
		v24 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v32 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v24
	} else {
		v20 = int32(1)
		if base.F64_gt(v7, v12) != 0 {
			v32 = v20
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10) {
				v32 = v20
			} else {
				v24 = v20
				v32 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v24
			}
		}
	}
	return v32
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
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
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
		v73 = int32(3)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+29)) = uint8(v73)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+27)) = uint16(v9)
		v77 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+26)) = uint8(v77)
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
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	v4 = int32(_a_F_btnamesortsupport_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_btnamesortsupport[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, _c_F_btnamesortsupport[0])) = v9
	F_varstr_sortsupport(m, v6, int32(19), v7)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_btnamesortsupport[0])) = v5
		return int32(0)
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	if v7 != int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	if int32(0) < v10 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v21 = int32(0)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v22 != 0 {
		v39 = v21
		goto L13
	} else {
		goto L14
	}
L4:
	;
	F__bt_killitems(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	F_ReleaseBuffer(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(-4294967296)
	goto L3
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v39)
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+19)) = uint8(v44)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+17)) = uint16(v44)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1])))
	if v48 != 0 {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	switch v24 {
	case 0, 5:
		goto L15
	default:
		v39 = v21
		goto L13
	}
L15:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
	if v27 != int32(112) {
		v39 = v21
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_btrescan[0]))
	if v31 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v34 != 0 {
		v39 = v21
		goto L13
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = base.B2i32(v36 != int32(0))
	goto L13
L20:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v35 != 0 {
		v39 = v21
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_ReleaseBuffer(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrescan[1]))) = int64(-4294967296)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v53 != int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	if l1 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	if v56 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v58 = F_palloc(m, int32(_a_F_btrescan_0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v58 - int32(-8192)
	goto L26
L30:
	;
	v76 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v76
	return
L31:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v67 <= int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v72 = v67 * int32(48)
	if v72 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L30
L34:
	;
	v73 = F__emscripten_memcpy_bulkmem(m, v70, l1, v72)
	mBase = m.M
	goto L36
L35:
	;
	goto L36
L36:
	;
	goto L33
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
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
							v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
							mBase = m.M
						} else {
						}
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
						if v44 != 0 {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
							if v46 != 0 {
								v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
								mBase = m.M
							} else {
							}
						} else {
						}
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						if v49 == int32(0) {
							return
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							F__bt_start_array_keys(m, l0, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
								return
							}
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
					v41 = v37*int32(10) + int32(58)
					if v41 != 0 {
						v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
						mBase = m.M
					} else {
					}
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
					if v44 != 0 {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
						if v46 != 0 {
							v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
							mBase = m.M
						} else {
						}
					} else {
					}
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					if v49 == int32(0) {
						return
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						F__bt_start_array_keys(m, l0, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
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
										v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
										mBase = m.M
									} else {
									}
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
									if v44 != 0 {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
										if v46 != 0 {
											v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
											mBase = m.M
										} else {
										}
									} else {
									}
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									if v49 == int32(0) {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
										F__bt_start_array_keys(m, l0, v52)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											v55 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
											return
										}
									}
								}
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
								v41 = v37*int32(10) + int32(58)
								if v41 != 0 {
									v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
									mBase = m.M
								} else {
								}
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v44 != 0 {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
									if v46 != 0 {
										v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
										mBase = m.M
									} else {
									}
								} else {
								}
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
								if v49 == int32(0) {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
									F__bt_start_array_keys(m, l0, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v55 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
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
											v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
											mBase = m.M
										} else {
										}
										v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
										if v44 != 0 {
											v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
											if v46 != 0 {
												v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
												mBase = m.M
											} else {
											}
										} else {
										}
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
										if v49 == int32(0) {
											return
										} else {
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
											F__bt_start_array_keys(m, l0, v52)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												v55 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
												return
											}
										}
									}
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
									v41 = v37*int32(10) + int32(58)
									if v41 != 0 {
										v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
										mBase = m.M
									} else {
									}
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
									if v44 != 0 {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
										if v46 != 0 {
											v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
											mBase = m.M
										} else {
										}
									} else {
									}
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									if v49 == int32(0) {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
										F__bt_start_array_keys(m, l0, v52)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											v55 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
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
									v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
									mBase = m.M
								} else {
								}
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v44 != 0 {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
									if v46 != 0 {
										v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
										mBase = m.M
									} else {
									}
								} else {
								}
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
								if v49 == int32(0) {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
									F__bt_start_array_keys(m, l0, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v55 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
										return
									}
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
							v41 = v37*int32(10) + int32(58)
							if v41 != 0 {
								v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
								mBase = m.M
							} else {
							}
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							if v44 != 0 {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
								if v46 != 0 {
									v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
									mBase = m.M
								} else {
								}
							} else {
							}
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
							if v49 == int32(0) {
								return
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
								F__bt_start_array_keys(m, l0, v52)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v55 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
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
										v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
										mBase = m.M
									} else {
									}
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
									if v44 != 0 {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
										if v46 != 0 {
											v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
											mBase = m.M
										} else {
										}
									} else {
									}
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									if v49 == int32(0) {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
										F__bt_start_array_keys(m, l0, v52)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											v55 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
											return
										}
									}
								}
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[2])))
								v41 = v37*int32(10) + int32(58)
								if v41 != 0 {
									v42 = F__emscripten_memcpy_bulkmem(m, v12, v6+int32(_a_F_btrestrpos_0), v41)
									mBase = m.M
								} else {
								}
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								if v44 != 0 {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F_btrestrpos[3])))
									if v46 != 0 {
										v47 = F__emscripten_memcpy_bulkmem(m, v44, v45, v46)
										mBase = m.M
									} else {
									}
								} else {
								}
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
								if v49 == int32(0) {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
									F__bt_start_array_keys(m, l0, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v55 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v55)
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
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
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
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
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
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int64
	_ = v412
	var v414 int64
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v492 float64
	_ = v492
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v621 int32
	_ = v621
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v703 int32
	_ = v703
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v827 int32
	_ = v827
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v882 float64
	_ = v882
	var v884 float64
	_ = v884
	var v885 int32
	_ = v885
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1127 int32
	_ = v1127
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1247 int32
	_ = v1247
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1356 int32
	_ = v1356
	var v1362 int32
	_ = v1362
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1438 int32
	_ = v1438
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1453 int64
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1466 int32
	_ = v1466
	var v1478 int32
	_ = v1478
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1584 float64
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1596 int32
	_ = v1596
	var v1601 int32
	_ = v1601
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1692 int32
	_ = v1692
	var v1731 float64
	_ = v1731
	var v1732 float64
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1768 int32
	_ = v1768
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1807 int32
	_ = v1807
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1833 int32
	_ = v1833
	var v1842 int32
	_ = v1842
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1975 int32
	_ = v1975
	var v1980 int32
	_ = v1980
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1993 int32
	_ = v1993
	var v2006 int32
	_ = v2006
	var v2013 int32
	_ = v2013
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2023 int32
	_ = v2023
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2087 int32
	_ = v2087
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2152 int32
	_ = v2152
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2220 int32
	_ = v2220
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2258 int32
	_ = v2258
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2268 int32
	_ = v2268
	var v2272 int32
	_ = v2272
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2310 int32
	_ = v2310
	var v2316 int32
	_ = v2316
	var v2318 int32
	_ = v2318
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2341 int32
	_ = v2341
	var v2346 int32
	_ = v2346
	var v2352 int32
	_ = v2352
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2366 int32
	_ = v2366
	var v2372 int32
	_ = v2372
	var v2374 int32
	_ = v2374
	var v2380 int32
	_ = v2380
	var v2382 int32
	_ = v2382
	var v2394 int32
	_ = v2394
	var v2398 int32
	_ = v2398
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2426 int32
	_ = v2426
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2439 int32
	_ = v2439
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2456 int32
	_ = v2456
	var v2460 int32
	_ = v2460
	var v2464 int32
	_ = v2464
	var v2468 int32
	_ = v2468
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2487 int32
	_ = v2487
	var v2493 int32
	_ = v2493
	var v2496 int64
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2501 int32
	_ = v2501
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2515 int32
	_ = v2515
	var v2522 int32
	_ = v2522
	var v2528 int32
	_ = v2528
	var v2530 int32
	_ = v2530
	var v2536 int32
	_ = v2536
	var v2539 int64
	_ = v2539
	var v2543 int32
	_ = v2543
	var v2545 int32
	_ = v2545
	var v2551 int32
	_ = v2551
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2560 int32
	_ = v2560
	var v2610 int32
	_ = v2610
	var v2612 int32
	_ = v2612
	var v2667 int32
	_ = v2667
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2687 int32
	_ = v2687
	var v2691 int32
	_ = v2691
	var v2693 int32
	_ = v2693
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2710 int32
	_ = v2710
	var v2713 int32
	_ = v2713
	var v2715 int32
	_ = v2715
	var v2717 int32
	_ = v2717
	var v2719 int32
	_ = v2719
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2733 int32
	_ = v2733
	var v2739 int32
	_ = v2739
	var v2741 int32
	_ = v2741
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2754 int32
	_ = v2754
	var v2758 int32
	_ = v2758
	var v2760 int32
	_ = v2760
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2773 int32
	_ = v2773
	var v2775 int32
	_ = v2775
	var v2779 int32
	_ = v2779
	var v2785 int32
	_ = v2785
	var v2787 int32
	_ = v2787
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2797 int32
	_ = v2797
	var v2799 int32
	_ = v2799
	var v2807 int32
	_ = v2807
	var __phi2807 int32
	_ = __phi2807
	var v2808 int32
	_ = v2808
	var __phi2808 int32
	_ = __phi2808
	var v2809 int32
	_ = v2809
	var __phi2809 int32
	_ = __phi2809
	var v2818 int32
	_ = v2818
	var __phi2818 int32
	_ = __phi2818
	var v2862 int32
	_ = v2862
	var v2864 int32
	_ = v2864
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2885 int32
	_ = v2885
	var v2890 int32
	_ = v2890
	var v2893 int32
	_ = v2893
	var v2896 int32
	_ = v2896
	var v2899 int32
	_ = v2899
	var v2901 int32
	_ = v2901
	var v2903 int32
	_ = v2903
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2910 int32
	_ = v2910
	var v2912 int32
	_ = v2912
	var v2916 int32
	_ = v2916
	var v2922 int32
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2936 int32
	_ = v2936
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2997 int32
	_ = v2997
	var v3003 int32
	_ = v3003
	var v3005 int32
	_ = v3005
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3017 int32
	_ = v3017
	var v3020 int32
	_ = v3020
	var v3022 int32
	_ = v3022
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3050 int32
	_ = v3050
	var v3055 int32
	_ = v3055
	var v3065 int32
	_ = v3065
	var v3068 int32
	_ = v3068
	var v3069 int32
	_ = v3069
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3075 int32
	_ = v3075
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3082 int32
	_ = v3082
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3090 int32
	_ = v3090
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3104 int32
	_ = v3104
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3131 int32
	_ = v3131
	var v3136 int32
	_ = v3136
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3148 int32
	_ = v3148
	var v3151 int32
	_ = v3151
	var v3153 int32
	_ = v3153
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3161 int32
	_ = v3161
	var v3166 int32
	_ = v3166
	var v3172 int32
	_ = v3172
	var v3174 int32
	_ = v3174
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3183 int32
	_ = v3183
	var v3185 int32
	_ = v3185
	var v3186 int32
	_ = v3186
	var v3189 int32
	_ = v3189
	var v3191 int32
	_ = v3191
	var v3195 int32
	_ = v3195
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3209 int32
	_ = v3209
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3218 int32
	_ = v3218
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3227 int32
	_ = v3227
	var v3234 int32
	_ = v3234
	var v3240 int32
	_ = v3240
	var v3242 int32
	_ = v3242
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3256 int32
	_ = v3256
	var v3262 int32
	_ = v3262
	var v3264 int32
	_ = v3264
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3276 int32
	_ = v3276
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3288 int32
	_ = v3288
	var v3294 int32
	_ = v3294
	var v3296 int32
	_ = v3296
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3304 int64
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3312 int32
	_ = v3312
	var v3314 int32
	_ = v3314
	var v3317 int32
	_ = v3317
	var v3320 int32
	_ = v3320
	var v3322 int32
	_ = v3322
	var v3325 int32
	_ = v3325
	var v3329 int32
	_ = v3329
	var v3340 int32
	_ = v3340
	var v3342 int32
	_ = v3342
	var v3344 int32
	_ = v3344
	var v3346 int32
	_ = v3346
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3356 int32
	_ = v3356
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3362 int32
	_ = v3362
	var v3366 int32
	_ = v3366
	var v3370 int32
	_ = v3370
	var v3374 int32
	_ = v3374
	var v3380 int32
	_ = v3380
	var v3392 int32
	_ = v3392
	var v3397 int64
	_ = v3397
	var v3398 int32
	_ = v3398
	var v3402 int32
	_ = v3402
	var v3403 int32
	_ = v3403
	var v3405 int32
	_ = v3405
	var v3407 int32
	_ = v3407
	var v3409 int32
	_ = v3409
	var v3411 int32
	_ = v3411
	var v3413 int32
	_ = v3413
	var v3415 int32
	_ = v3415
	var v3422 int32
	_ = v3422
	var v3425 int64
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3430 int64
	_ = v3430
	var v3434 int32
	_ = v3434
	var v3440 int32
	_ = v3440
	var v3442 int32
	_ = v3442
	var v3448 int32
	_ = v3448
	var v3455 int32
	_ = v3455
	var v3461 int32
	_ = v3461
	var v3463 int32
	_ = v3463
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3474 int32
	_ = v3474
	var v3479 int32
	_ = v3479
	var v3485 int32
	_ = v3485
	var v3487 int32
	_ = v3487
	var v3493 int32
	_ = v3493
	var v3500 int32
	_ = v3500
	var v3504 int32
	_ = v3504
	var v3506 int32
	_ = v3506
	var v3510 int32
	_ = v3510
	var v3517 int32
	_ = v3517
	var v3519 int32
	_ = v3519
	var v3525 int32
	_ = v3525
	var v3527 int32
	_ = v3527
	var v3530 int32
	_ = v3530
	var v3532 int32
	_ = v3532
	var v3535 int32
	_ = v3535
	var v3537 int32
	_ = v3537
	var v3542 int32
	_ = v3542
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3550 int32
	_ = v3550
	var v3554 int32
	_ = v3554
	var v3555 int32
	_ = v3555
	var v3557 int32
	_ = v3557
	var v3559 int32
	_ = v3559
	var v3561 int32
	_ = v3561
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3568 int32
	_ = v3568
	var v3569 int32
	_ = v3569
	var v3571 int32
	_ = v3571
	var v3572 int32
	_ = v3572
	var v3573 int32
	_ = v3573
	var v3574 int32
	_ = v3574
	var v3578 int32
	_ = v3578
	var v3579 int32
	_ = v3579
	var v3584 int32
	_ = v3584
	var v3590 int32
	_ = v3590
	var v3598 int32
	_ = v3598
	var v3604 int32
	_ = v3604
	var v3606 int32
	_ = v3606
	var v3613 int32
	_ = v3613
	var v3658 int32
	_ = v3658
	var v3661 int32
	_ = v3661
	var v3663 int32
	_ = v3663
	var v3665 int32
	_ = v3665
	var v3667 int32
	_ = v3667
	var v3670 int32
	_ = v3670
	var v3671 int32
	_ = v3671
	var v3674 int32
	_ = v3674
	var v3676 int32
	_ = v3676
	var v3680 int32
	_ = v3680
	var v3684 int32
	_ = v3684
	var v3689 int32
	_ = v3689
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3704 int32
	_ = v3704
	var v3709 int32
	_ = v3709
	var v3713 int32
	_ = v3713
	var v3716 int32
	_ = v3716
	var v3717 int32
	_ = v3717
	var v3718 int32
	_ = v3718
	var v3729 int32
	_ = v3729
	var v3734 int32
	_ = v3734
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3749 int32
	_ = v3749
	var v3754 int32
	_ = v3754
	var v3759 int32
	_ = v3759
	var v3808 int32
	_ = v3808
	var v3809 int32
	_ = v3809
	var v3814 int32
	_ = v3814
	var v3815 int32
	_ = v3815
	var v3822 int32
	_ = v3822
	var v3827 int32
	_ = v3827
	var v3881 int32
	_ = v3881
	var v3934 int32
	_ = v3934
	var v4021 int32
	_ = v4021
	var v4043 int32
	_ = v4043
	var v4074 int32
	_ = v4074
	var v4099 int32
	_ = v4099
	var v4100 int32
	_ = v4100
	var v4102 int32
	_ = v4102
	var v4103 int32
	_ = v4103
	var v4104 int32
	_ = v4104
	var v4106 int32
	_ = v4106
	var v4111 int32
	_ = v4111
	var v4112 int32
	_ = v4112
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4126 int32
	_ = v4126
	var v4131 int32
	_ = v4131
	var v4135 int32
	_ = v4135
	var v4187 int32
	_ = v4187
	var v4194 int32
	_ = v4194
	var v4198 int32
	_ = v4198
	var v4201 int32
	_ = v4201
	var v4203 int32
	_ = v4203
	var v4204 int32
	_ = v4204
	var v4207 int32
	_ = v4207
	var v4215 int32
	_ = v4215
	var v4221 int32
	_ = v4221
	var v4226 int32
	_ = v4226
	var v4228 int32
	_ = v4228
	var v4230 int32
	_ = v4230
	var v4231 int32
	_ = v4231
	var v4233 int32
	_ = v4233
	var v4234 int32
	_ = v4234
	var v4237 int32
	_ = v4237
	var v4238 int32
	_ = v4238
	var v4239 int32
	_ = v4239
	var v4240 int32
	_ = v4240
	var v4241 int32
	_ = v4241
	var v4242 int32
	_ = v4242
	var v4243 int32
	_ = v4243
	var v4248 int32
	_ = v4248
	var v4297 int32
	_ = v4297
	var v4300 int32
	_ = v4300
	var v4301 int32
	_ = v4301
	var v4302 int64
	_ = v4302
	var v4303 int32
	_ = v4303
	var v4304 int32
	_ = v4304
	var v4308 int32
	_ = v4308
	var v4309 int32
	_ = v4309
	var v4310 int32
	_ = v4310
	var v4314 int32
	_ = v4314
	var v4315 int32
	_ = v4315
	var v4368 int32
	_ = v4368
	var v4371 int32
	_ = v4371
	var v4421 int32
	_ = v4421
	var v4473 int32
	_ = v4473
	var v4475 int32
	_ = v4475
	v5 = l4
	v6 = int32(0)
	v48 = int64(0)
	v52 = m.G0
	v54 = v52 - int32(2512)
	m.G0 = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+28)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+40)) = uint16(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v54)+32)) = l2
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0]))
	v74 = F_AllocSetContextCreateInternal(m, v69, int32(_a_F_btvacuumscan_0), v6, int32(_a_F_btvacuumscan_1), int32(_a_F_btvacuumscan_2))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v76 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v54)+56)) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v54)+48)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v54)+44)) = v74
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v82 = v54 + int32(24)
	v83 = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = v83
	v86 = int32(67108863)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[1]))
	v92 = v88 << (uint(int32(6)) % 32) & int32(268435392)
	if base.Ui32(v86) <= base.Ui32(v92) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)))
	if v108 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v95 = v86
	goto L8
L7:
	;
	v95 = v92
	goto L8
L8:
	;
	if base.Ui32(v95) <= base.Ui32(int32(256)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v98 = v83
	goto L11
L10:
	;
	v98 = v95
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v98
	v101 = F_palloc(m, int32(_a_F_btvacuumscan_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+32)) = v101
	goto L5
L13:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
	v114 = base.B2i32(v111 == int32(0))
	goto L15
L14:
	;
	v114 = v6
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v123 = F_read_stream_begin_relation(m, int32(13), v118, v56, int32(120), v54+int32(16), int32(0))
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
	v136 = v54
	v150 = v56
	v158 = v123
	v161 = v114
	goto L17
L17:
	;
	if v161 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	F_read_stream_end(m, v158)
	mBase = m.M
	v4226 = m.ExcPending
	if v4226 != 0 {
		goto L1
	} else {
		goto L762
	}
L19:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+9)))
	if v191 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v179 = F_RelationGetNumberOfBlocksInFork(m, v150, int32(0))
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
	F_LockRelationForExtension(m, v150, int32(7))
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
	v185 = F_RelationGetNumberOfBlocksInFork(m, v150, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_UnlockRelationForExtension(m, v150, int32(7))
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
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	if base.Ui32(v229) < base.Ui32(v190) {
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
	if v202 != int32(1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v205 = int32(_a_F_btvacuumscan_4)
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	v208 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v207 + v208
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v211 + v208
	*(*int64)(unsafe.Add(mBase, uint32(v198+int32(120))+232)) = base.I64_extend_i32_u(v190)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v219 + v208
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v225 - v208
	goto L31
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = v190
	v232 = v125
	v233 = v126
	v243 = v136
	v257 = v150
	v265 = v158
	v268 = v161
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
	v287 = F_read_stream_next_buffer(m, v265, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v4187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+9)))
	if v4187 != int32(1) {
		goto L37
	} else {
		goto L757
	}
L41:
	;
	F__bt_relbuf(m, v317)
	mBase = m.M
	v4135 = m.ExcPending
	if v4135 != 0 {
		goto L1
	} else {
		goto L756
	}
L42:
	;
	v4111 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4112 = m.ExcPending
	if v4112 != 0 {
		goto L1
	} else {
		goto L751
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
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v243)+24))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v243)+36))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v243)+32))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v243)+28))
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
	F_read_stream_reset(m, v265)
	mBase = m.M
	v4106 = m.ExcPending
	if v4106 != 0 {
		goto L1
	} else {
		goto L750
	}
L47:
	;
	v317 = v287
	v320 = v313
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
	F__bt_lockbuf(m, v317, int32(1))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
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
	if v4074 == int32(0) {
		goto L40
	} else {
		goto L747
	}
L55:
	;
	F__bt_relbuf(m, v317)
	mBase = m.M
	v4043 = m.ExcPending
	if v4043 != 0 {
		goto L1
	} else {
		goto L746
	}
L56:
	;
	v443 = int32(0)
	v445 = v418 & int32(_a_F_btvacuumscan_5)
	if v445&int32(16) == v443 {
		goto L87
	} else {
		goto L88
	}
L57:
	;
	v4021 = int32(0)
	goto L55
L58:
	;
	F_RecordFreeIndexPage(m, v291, v320)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L86
	}
L59:
	;
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385)+14)))
	if v386 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v371+(v317^int32(-1))<<(uint(int32(2))%32))))
	v385 = v377
	goto L59
L61:
	;
	goto L62
L62:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v385 = v379 + v317<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	F__bt_checkpage(m, v291, v317)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v320 != v313 {
		goto L42
	} else {
		goto L85
	}
L66:
	;
	v389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385)+16)))
	v390 = v385 + v389
	if v320 != v313 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v407&int32(4) != 0 {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	if v390 == int32(0) {
		goto L42
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v390 == int32(0) {
		goto L58
	} else {
		goto L75
	}
L71:
	;
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+12)))
	if v394&int32(17) != int32(1) {
		goto L42
	} else {
		goto L72
	}
L72:
	;
	if v394&int32(4) != 0 {
		goto L41
	} else {
		goto L73
	}
L73:
	;
	v401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+14)))
	v402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243)+40)))
	if v401 == v402 {
		v407 = v394
		goto L67
	} else {
		goto L74
	}
L74:
	;
	goto L41
L75:
	;
	v406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+12)))
	v407 = v406
	goto L67
L76:
	;
	if v407&int32(256) != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v418 = v407
	goto L78
L78:
	;
	if v418&int32(4) == int32(0) {
		goto L56
	} else {
		goto L84
	}
L79:
	;
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v385)+24))
	v414 = v412
	goto L81
L80:
	;
	v414 = int64(3)
	goto L81
L81:
	;
	v415 = F_GlobalVisCheckRemovableFullXid(m, v290, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v415 != 0 {
		goto L58
	} else {
		goto L83
	}
L83:
	;
	v417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+12)))
	v418 = v417
	goto L78
L84:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v294)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+28)) = v423 + int32(1)
	goto L57
L85:
	;
	goto L58
L86:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v294)+28))
	v433 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+28)) = v432 + v433
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v294)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+32)) = v436 + v433
	goto L57
L87:
	;
	if v445&int32(1) == int32(0) {
		v4021 = v443
		goto L55
	} else {
		goto L90
	}
L88:
	;
	v1768 = v443
	goto L89
L89:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	F_MemoryContextReset(m, v1789)
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L1
	} else {
		goto L258
	}
L90:
	;
	F_LockBuffer(m, v317, int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_LockBufferForCleanup(m, v317)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v459 = int32(0)
	v461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243)+40)))
	if v461 == v459 {
		v478 = v459
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v481 != 0 {
		goto L103
	} else {
		goto L104
	}
L94:
	;
	v465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+14)))
	if v465 != v461 {
		v478 = int32(0)
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+12)))
	if v468&int32(32) != 0 {
		v478 = int32(0)
		goto L93
	} else {
		goto L96
	}
L96:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if base.Ui32(v471) < base.Ui32(v313) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v474 = v471
	goto L99
L98:
	;
	v474 = int32(0)
	goto L99
L99:
	;
	if v471 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v476 = v474
	goto L102
L101:
	;
	v476 = int32(0)
	goto L102
L102:
	;
	v478 = v476
	goto L93
L103:
	;
	v482 = int32(2)
	goto L105
L104:
	;
	v482 = int32(1)
	goto L105
L105:
	;
	v483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v483) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v491 = int32(base.Ui32(v483+int32(_a_F_btvacuumscan_6)) >> (uint(int32(2)) % 32))
	goto L108
L107:
	;
	v491 = int32(0)
	goto L108
L108:
	;
	v492 = float64(0)
	if v293 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v885 = int32(0)
	if base.B2i32(v843 <= v885)&base.B2i32(v845 <= v885) == v885 {
		goto L149
	} else {
		goto L150
	}
L110:
	;
	v843 = v459
	v845 = int32(0)
	v882 = v492
	v884 = float64(0)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v497 = int32(0)
	v500 = v491 & int32(_a_F_btvacuumscan_5)
	if base.Ui32(v500) < base.Ui32(v482) {
		v843 = v459
		v845 = v497
		v882 = v492
		v884 = float64(0)
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v504 = int32(0)
	v514 = v504
	v516 = v459
	v518 = v497
	v519 = v482
	v525 = v504
	goto L114
L114:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v519&int32(_a_F_btvacuumscan_5)<<(uint(int32(2))%32)+(v385+int32(24))-int32(4))))
	v567 = v385 + v564&int32(_a_F_btvacuumscan_7)
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+7)))
	if v568&int32(32) != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v843 = v785
	v845 = v787
	v882 = base.F64_convert_i32_s(v783)
	v884 = base.F64_convert_i32_s(v794)
	goto L109
L116:
	;
	v827 = v519 + int32(1)
	if base.Ui32(v827&int32(_a_F_btvacuumscan_5)) <= base.Ui32(v500) {
		v514 = v783
		v516 = v785
		v518 = v787
		v519 = v827
		v525 = v794
		goto L114
	} else {
		goto L147
	}
L117:
	;
	v590 = v571 & int32(4095)
	if v590 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L118:
	;
	v571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v567)+4)))
	if v571&int32(_a_F_btvacuumscan_1) != 0 {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v575 = m.T0[v293].(func(*base.Module, int32, int32) int32)(m, v567, v292)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	if v575 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v579 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v243+int32(1696)+v516<<(uint(v579)%32)))) = uint16(v519)
	v783 = v514 + v579
	v785 = v516 + v579
	v787 = v518
	v794 = v525
	goto L116
L124:
	;
	goto L125
L125:
	;
	v783 = v514
	v785 = v516
	v787 = v518
	v794 = v525 + int32(1)
	goto L116
L126:
	;
	v783 = v731
	v785 = v733
	v787 = v735
	v794 = v737 + v525
	goto L116
L127:
	;
	v731 = v514
	v733 = v516
	v735 = v518
	v737 = int32(0)
	goto L126
L128:
	;
	goto L129
L129:
	;
	v594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v567)+2)))
	v595 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v567))))
	v604 = int32(0)
	v609 = v604
	v611 = v604
	v621 = v604
	goto L130
L130:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v243)+36))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v243)+32))
	v663 = m.T0[v662].(func(*base.Module, int32, int32) int32)(m, v567+(v594|v595<<(uint(int32(16))%32))+v609*int32(6), v661)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	if v686 == int32(0) {
		v731 = v514
		v733 = v516
		v735 = v518
		v737 = v687
		goto L126
	} else {
		goto L142
	}
L132:
	;
	v690 = v609 + int32(1)
	if v690 != v590 {
		v609 = v690
		v611 = v686
		v621 = v687
		goto L130
	} else {
		goto L141
	}
L133:
	;
	if v663 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v686 = v611
	v687 = v621 + int32(1)
	goto L132
L135:
	;
	goto L136
L136:
	;
	if v611 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v671 = F_palloc(m, v590<<(uint(int32(1))%32)+int32(8))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v611)+6)))
	v679 = int32(1)
	v680 = v678 + v679
	*(*uint16)(unsafe.Add(mBase, uint32(v611)+6)) = uint16(v680)
	*(*uint16)(unsafe.Add(mBase, uint32(v611+v678<<(uint(v679)%32))+8)) = uint16(v609)
	v686 = v611
	v687 = v621
	goto L132
L140:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v671)+8)) = uint16(v609)
	v674 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v671)+6)) = uint16(v674)
	*(*uint16)(unsafe.Add(mBase, uint32(v671)+4)) = uint16(v519)
	*(*int32)(unsafe.Add(mBase, uint32(v671))) = v567
	v686 = v671
	v687 = v621
	goto L132
L141:
	;
	goto L131
L142:
	;
	if int32(0) < v687 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243-int32(-64)+v518<<(uint(int32(2))%32)))) = v686
	v703 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v567)+4)))
	v731 = v514 - v687 + v703&int32(4095)
	v733 = v516
	v735 = v518 + int32(1)
	v737 = v687
	goto L126
L144:
	;
	goto L145
L145:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v243+int32(1696)+v516<<(uint(int32(1))%32)))) = uint16(v519)
	v715 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v567)+4)))
	F_pfree(m, v686)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v731 = v514 + v715&int32(4095)
	v733 = v516 + int32(1)
	v735 = v518
	v737 = v687
	goto L126
L147:
	;
	goto L115
L148:
	;
	if base.Ui32(v482) <= base.Ui32(v1692&int32(_a_F_btvacuumscan_5)) {
		goto L250
	} else {
		goto L251
	}
L149:
	;
	v893 = v243 + int32(1696)
	v895 = v243 - int32(-64)
	v896 = int32(0)
	v901 = m.G0
	v903 = v901 - int32(832)
	m.G0 = v903
	if v317 < v896 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	goto L151
L151:
	;
	v1661 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243)+40)))
	if v1661 == int32(0) {
		v1692 = v491
		goto L148
	} else {
		goto L247
	}
L152:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924)+118)))
	if v925 != int32(112) {
		v938 = int32(0)
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v908 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v908+(v317^int32(-1))<<(uint(int32(2))%32))))
	v922 = v914
	goto L152
L154:
	;
	goto L155
L155:
	;
	v916 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v922 = v916 + v317<<(uint(int32(13))%32) + int32(-8192)
	goto L152
L156:
	;
	if int32(0) < v845 {
		goto L161
	} else {
		goto L162
	}
L157:
	;
	v930 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if int32(0) < v930 {
		v938 = int32(1)
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v291)+32))
	if v934 != 0 {
		v938 = int32(0)
		goto L156
	} else {
		goto L159
	}
L159:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v291)+40))
	v938 = base.B2i32(v935 == int32(0))
	goto L156
L160:
	;
	if int32(0) < v843 {
		goto L207
	} else {
		goto L208
	}
L161:
	;
	v950 = v896
	v954 = v896
	goto L164
L162:
	;
	goto L163
L163:
	;
	v1348 = int32(_a_F_btvacuumscan_4)
	v1350 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1350 + int32(1)
	v1356 = v896
	v1362 = v896
	goto L160
L164:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v895+v950<<(uint(int32(2))%32))))
	F__bt_update_posting(m, v995)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L166
	}
L165:
	;
	v1014 = int32(0)
	if v938 != 0 {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	v998 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+6)))
	v1001 = int32(1)
	v1004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v903+int32(16)+v950<<(uint(v1001)%32)))) = uint16(v1004)
	v1010 = v954 + v998<<(uint(v1001)%32) + int32(2)
	v1012 = v950 + v1001
	if v1012 != v845 {
		v950 = v1012
		v954 = v1010
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v1018 = int32(0)
	v1019 = F_palloc(m, v1010)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	v1183 = v896
	v1189 = v1014
	goto L170
L170:
	;
	v1232 = int32(_a_F_btvacuumscan_4)
	v1234 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1234 + int32(1)
	v1247 = v1014
	goto L193
L171:
	;
	if v845 != int32(1) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v1027 = v896
	v1040 = v1018
	v1042 = v896
	goto L175
L173:
	;
	v1114 = v896
	v1127 = v1018
	goto L174
L174:
	;
	if v845&int32(1) != 0 {
		goto L186
	} else {
		goto L187
	}
L175:
	;
	v1077 = int32(2)
	v1079 = v895 + v1040<<(uint(v1077)%32)
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1079)))
	v1081 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1080)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1027+v1019))) = uint16(v1081)
	v1084 = v1027 + v1077
	v1089 = v1081 << (uint(int32(1)) % 32)
	if v1089 != 0 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v1114 = v1106
	v1127 = v1108
	goto L174
L177:
	;
	v1092 = v1084 + v1089
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+4))
	v1095 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1094)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1019+v1092))) = uint16(v1095)
	v1098 = v1092 + int32(2)
	v1103 = v1095 << (uint(int32(1)) % 32)
	if v1103 != 0 {
		goto L182
	} else {
		goto L183
	}
L178:
	;
	v1090 = F__emscripten_memcpy_bulkmem(m, v1019+v1084, v1080+int32(8), v1089)
	mBase = m.M
	goto L180
L179:
	;
	goto L180
L180:
	;
	goto L177
L181:
	;
	v1106 = v1098 + v1103
	v1107 = int32(2)
	v1108 = v1040 + v1107
	v1110 = v1042 + v1107
	if v1110 != v845&int32(2147483646) {
		v1027 = v1106
		v1040 = v1108
		v1042 = v1110
		goto L175
	} else {
		goto L185
	}
L182:
	;
	v1104 = F__emscripten_memcpy_bulkmem(m, v1019+v1098, v1094+int32(8), v1103)
	mBase = m.M
	goto L184
L183:
	;
	goto L184
L184:
	;
	goto L181
L185:
	;
	goto L176
L186:
	;
	v1163 = v1114 + v1019
	v1164 = int32(2)
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v895+v1127<<(uint(v1164)%32))))
	v1168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1167)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1163))) = uint16(v1168)
	v1175 = v1168 << (uint(int32(1)) % 32)
	if v1175 != 0 {
		goto L190
	} else {
		goto L191
	}
L187:
	;
	goto L188
L188:
	;
	v1183 = v1010
	v1189 = v1019
	goto L170
L189:
	;
	goto L188
L190:
	;
	v1176 = F__emscripten_memcpy_bulkmem(m, v1163+v1164, v1167+int32(8), v1175)
	mBase = m.M
	goto L192
L191:
	;
	goto L192
L192:
	;
	goto L189
L193:
	;
	v1294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v903+int32(16)+v1247<<(uint(int32(1))%32)))))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v895+v1247<<(uint(int32(2))%32))))
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1298)))
	v1300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1299)+6)))
	v1307 = F_PageIndexTupleOverwrite(m, v922, v1294, v1299, (v1300&int32(_a_F_btvacuumscan_8)+int32(7))&int32(_a_F_btvacuumscan_9))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L1
	} else {
		goto L195
	}
L194:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L1
	} else {
		goto L200
	}
L195:
	;
	if v1307 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1310 = v1247 + int32(1)
	if v1310 != v845 {
		v1247 = v1310
		goto L193
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	goto L194
L199:
	;
	v1356 = v1183
	v1362 = v1189
	goto L160
L200:
	;
	if v317 < int32(0) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v903))) = v1334
	*(*int32)(unsafe.Add(mBase, uint32(v903)+4)) = v1335 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_10), v903)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L205
	}
L202:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1319+(v317^int32(-1))<<(uint(int32(6))%32))+16))
	v1334 = v1325
	goto L201
L203:
	;
	goto L204
L204:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1327+v317<<(uint(int32(6))%32)+int32(-64))+16))
	v1334 = v1333
	goto L201
L205:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(1200), int32(_a_F_btvacuumscan_12))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	F_PageIndexMultiDelete(m, v922, v893, v843)
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L1
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v1409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v922)+16)))
	v1410 = v922 + v1409
	v1411 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1410)+14)) = uint16(v1411)
	v1413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1410)+12)))
	v1415 = v1413 & int32(_a_F_btvacuumscan_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v1410)+12)) = uint16(v1415)
	F_MarkBufferDirty(m, v317)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L1
	} else {
		goto L211
	}
L210:
	;
	goto L209
L211:
	;
	if v938 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v903)+14)) = uint16(v845)
	*(*uint16)(unsafe.Add(mBase, uint32(v903)+12)) = uint16(v843)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L1
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v1459 = int32(_a_F_btvacuumscan_4)
	v1461 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v1461 - int32(1)
	if v1362 != 0 {
		goto L228
	} else {
		goto L229
	}
L215:
	;
	F_XLogRegisterBuffer(m, int32(0), v317, int32(8))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_XLogRegisterData(m, v903+int32(12), int32(4))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	if int32(0) < v843 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	F_XLogRegisterBufData(m, int32(0), v893, v843<<(uint(int32(1))%32))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	if int32(0) < v845 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	goto L220
L222:
	;
	F_XLogRegisterBufData(m, int32(0), v903+int32(16), v845<<(uint(int32(1))%32))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1453 = F_XLogInsert(m, int32(11), int32(192))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L1
	} else {
		goto L227
	}
L225:
	;
	F_XLogRegisterBufData(m, int32(0), v1362, v1356)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	goto L224
L227:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v922))) = base.I64_rotr(v1453, int64(32))
	goto L214
L228:
	;
	F_pfree(m, v1362)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L1
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	if int32(0) < v845 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	v1478 = int32(0)
	goto L235
L233:
	;
	goto L234
L234:
	;
	m.G0 = v903 + int32(832)
	v1584 = *(*float64)(unsafe.Add(mBase, uint32(v294)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v294)+16)) = base.F64_add(v882, v1584)
	v1587 = int32(0)
	v1588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1588) {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v895+v1478<<(uint(int32(2))%32))))
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1523)))
	F_pfree(m, v1524)
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L1
	} else {
		goto L237
	}
L236:
	;
	goto L234
L237:
	;
	v1528 = v1478 + int32(1)
	if v1528 != v845 {
		v1478 = v1528
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	v1596 = int32(base.Ui32(v1588+int32(_a_F_btvacuumscan_6)) >> (uint(int32(2)) % 32))
	goto L241
L240:
	;
	v1596 = v1587
	goto L241
L241:
	;
	if v845 <= int32(0) {
		v1692 = v1596
		goto L148
	} else {
		goto L242
	}
L242:
	;
	v1601 = v1587
	goto L243
L243:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v243-int32(-64)+v1601<<(uint(int32(2))%32))))
	F_pfree(m, v1655)
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L1
	} else {
		goto L245
	}
L244:
	;
	v1692 = v1596
	goto L148
L245:
	;
	v1659 = v1601 + int32(1)
	if v1659 != v845 {
		v1601 = v1659
		goto L243
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	v1664 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+14)))
	if v1664 != v1661 {
		v1692 = v491
		goto L148
	} else {
		goto L248
	}
L248:
	;
	v1666 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v390)+14)) = uint16(v1666)
	F_MarkBufferDirtyHint(m, v317, int32(1))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	v1692 = v491
	goto L148
L250:
	;
	if v293 != 0 {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	goto L252
L252:
	;
	if v320 != v313 {
		v4021 = v478
		goto L55
	} else {
		goto L257
	}
L253:
	;
	v1731 = v884
	goto L255
L254:
	;
	v1731 = base.F64_convert_i32_u((v1692-v482)&int32(_a_F_btvacuumscan_5) + int32(1))
	goto L255
L255:
	;
	v1732 = *(*float64)(unsafe.Add(mBase, uint32(v294)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v294)+8)) = base.F64_add(v1731, v1732)
	F__bt_relbuf(m, v317)
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v4074 = v478
	goto L54
L257:
	;
	v1768 = v478
	goto L89
L258:
	;
	v1792 = int32(_a_F_btvacuumscan_14)
	v1793 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0]))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0])) = v1795
	v1798 = v243 + int32(24)
	v1799 = int32(0)
	v1800 = m.G0
	v1802 = v1800 - int32(288)
	m.G0 = v1802
	if v317 < v1799 {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	v1833 = v317
	v1842 = v1799
	goto L270
L260:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1807+(v317^int32(-1))<<(uint(int32(6))%32))+16))
	v1822 = v1813
	goto L259
L261:
	;
	goto L262
L262:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1815+v317<<(uint(int32(6))%32)+int32(-64))+16))
	v1822 = v1821
	goto L259
L263:
	;
	m.G0 = v1802 + int32(288)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[0])) = v1793
	v4074 = v1768
	goto L54
L264:
	;
	F_ReleaseBuffer(m, v1833)
	mBase = m.M
	v3934 = m.ExcPending
	if v3934 != 0 {
		goto L1
	} else {
		goto L745
	}
L265:
	;
	F_LockBuffer(m, v1833, int32(0))
	mBase = m.M
	v3881 = m.ExcPending
	if v3881 != 0 {
		goto L1
	} else {
		goto L744
	}
L266:
	;
	v3808 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3809 = m.ExcPending
	if v3809 != 0 {
		goto L1
	} else {
		goto L739
	}
L267:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3738 = m.ExcPending
	if v3738 != 0 {
		goto L1
	} else {
		goto L736
	}
L268:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3713 = m.ExcPending
	if v3713 != 0 {
		goto L1
	} else {
		goto L732
	}
L269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3694 = m.ExcPending
	if v3694 != 0 {
		goto L1
	} else {
		goto L729
	}
L270:
	;
	v1874 = int32(0)
	v1875 = base.B2i32(v1874 <= v1833)
	if v1875 == v1874 {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3680 = m.ExcPending
	if v3680 != 0 {
		goto L1
	} else {
		goto L726
	}
L272:
	;
	v1894 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1893)+16)))
	v1895 = v1894 + v1893
	v1896 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1895)+12)))
	if v1896&int32(5) != int32(1) {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1879+(v1833^int32(-1))<<(uint(int32(2))%32))))
	v1893 = v1885
	goto L272
L274:
	;
	goto L275
L275:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v1893 = v1887 + v1833<<(uint(int32(13))%32) + int32(-8192)
	goto L272
L276:
	;
	if v1896&int32(16) == int32(0) {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	goto L278
L278:
	;
	if v1896&int32(2) != 0 {
		goto L300
	} else {
		goto L301
	}
L279:
	;
	v1932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895)+12)))
	if v1932&int32(4) == int32(0) {
		goto L287
	} else {
		goto L288
	}
L280:
	;
	v1907 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	if v1907 == int32(0) {
		goto L279
	} else {
		goto L282
	}
L282:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+208)) = v1914 + int32(4)
	F_errmsg(m, int32(_a_F_btvacuumscan_15), v1802+int32(208))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	F_errhint(m, int32(_a_F_btvacuumscan_16), int32(0))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(1863), int32(_a_F_btvacuumscan_0))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	goto L279
L287:
	;
	F_LockBuffer(m, v1833, int32(0))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L1
	} else {
		goto L298
	}
L288:
	;
	v1939 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	if v1939 == int32(0) {
		goto L287
	} else {
		goto L290
	}
L290:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	if v1833 < int32(0) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+196)) = v1822
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+192)) = v1964
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+200)) = v1965 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_17), v1802+int32(192))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L1
	} else {
		goto L296
	}
L293:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1949+(v1833^int32(-1))<<(uint(int32(6))%32))+16))
	v1964 = v1955
	goto L292
L294:
	;
	goto L295
L295:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v1957+v1833<<(uint(int32(6))%32)+int32(-64))+16))
	v1964 = v1963
	goto L292
L296:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(1871), int32(_a_F_btvacuumscan_0))
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	goto L287
L298:
	;
	goto L264
L299:
	;
	if v1896&int32(16) == int32(0) {
		goto L307
	} else {
		goto L308
	}
L300:
	;
	F_LockBuffer(m, v1833, int32(0))
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L1
	} else {
		goto L305
	}
L301:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1895)+4))
	if v1988 == int32(0) {
		goto L300
	} else {
		goto L302
	}
L302:
	;
	if v1896&int32(128) != 0 {
		goto L300
	} else {
		goto L303
	}
L303:
	;
	v1993 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1893)+12)))
	if base.B2i32(base.Ui32(v1993) < base.Ui32(int32(25)))|base.B2i32((v1993+int32(_a_F_btvacuumscan_6))&int32(_a_F_btvacuumscan_18) == int32(0)) != 0 {
		goto L299
	} else {
		goto L304
	}
L304:
	;
	goto L300
L305:
	;
	goto L264
L306:
	;
	goto L271
L307:
	;
	if v1842 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L308:
	;
	v2560 = v1896
	goto L309
L309:
	;
	if v2560&int32(16) != 0 {
		goto L442
	} else {
		goto L443
	}
L310:
	;
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+24))
	v2017 = F_CopyIndexTuple(m, v1893+v2013&int32(_a_F_btvacuumscan_7))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L1
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v1798)))
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v2063)+4))
	if v1875 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L313:
	;
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v1895)))
	if v1833 < int32(0) {
		goto L315
	} else {
		goto L316
	}
L314:
	;
	F_LockBuffer(m, v1833, int32(0))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L1
	} else {
		goto L318
	}
L315:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v2023+(v1833^int32(-1))<<(uint(int32(6))%32))+16))
	v2038 = v2029
	goto L314
L316:
	;
	goto L317
L317:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v2031+v1833<<(uint(int32(6))%32)+int32(-64))+16))
	v2038 = v2037
	goto L314
L318:
	;
	v2042 = F__bt_leftsib_splitflag(m, v291, v2019, v2038)
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	if v2042 != 0 {
		goto L264
	} else {
		goto L320
	}
L320:
	;
	v2044 = F__bt_mkscankey(m, v291, v2017)
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v2046 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2044)+3)) = uint16(v2046)
	v2052 = F__bt_search(m, v291, int32(0), v2044, v1802+int32(248), int32(1))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v1802)+248))
	F_LockBuffer(m, v2054, int32(0))
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	F_ReleaseBuffer(m, v2054)
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_LockBuffer(m, v1833, int32(2))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	v1842 = v2052
	goto L270
L326:
	;
	v2083 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2082)+16)))
	if v1833 < int32(0) {
		goto L331
	} else {
		goto L332
	}
L327:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v2068+(v1833^int32(-1))<<(uint(int32(2))%32))))
	v2082 = v2074
	goto L326
L328:
	;
	goto L329
L329:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2082 = v2076 + v1833<<(uint(int32(13))%32) + int32(-8192)
	goto L326
L330:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v2082+v2083)+4))
	v2105 = F_ReadBuffer(m, v291, v2104)
	mBase = m.M
	v2106 = m.ExcPending
	if v2106 != 0 {
		goto L1
	} else {
		goto L334
	}
L331:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2087+(v1833^int32(-1))<<(uint(int32(6))%32))+16))
	v2102 = v2093
	goto L330
L332:
	;
	goto L333
L333:
	;
	v2095 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v2095+v1833<<(uint(int32(6))%32)+int32(-64))+16))
	v2102 = v2101
	goto L330
L334:
	;
	F_LockBuffer(m, v2105, int32(1))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	F__bt_checkpage(m, v291, v2105)
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	if v2105 < int32(0) {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v2130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2129)+16)))
	v2132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2130+v2129)+12)))
	F_LockBuffer(m, v2105, int32(0))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L1
	} else {
		goto L341
	}
L338:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2115+(v2105^int32(-1))<<(uint(int32(2))%32))))
	v2129 = v2121
	goto L337
L339:
	;
	goto L340
L340:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2129 = v2123 + v2105<<(uint(int32(13))%32) + int32(-8192)
	goto L337
L341:
	;
	F_ReleaseBuffer(m, v2105)
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	if v2132&int32(16) != 0 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v2142 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2143 = m.ExcPending
	if v2143 != 0 {
		goto L1
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	v2158 = F__bt_getstackbuf(m, v291, v2064, v1842, v2102)
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L1
	} else {
		goto L350
	}
L346:
	;
	if v2142 == int32(0) {
		goto L265
	} else {
		goto L347
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+180)) = v2104
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+176)) = v2102
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_19), v1802+int32(176))
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2128), int32(_a_F_btvacuumscan_20))
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	goto L265
L350:
	;
	if v2158 == int32(0) {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v3759 = v2102
	goto L266
L352:
	;
	goto L353
L353:
	;
	v2165 = v1842
	v2167 = v2158
	v2174 = v2104
	v2177 = v2102
	goto L354
L354:
	;
	v2213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2165)+4)))
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v2165)))
	v2215 = int32(0)
	v2216 = base.B2i32(v2215 <= v2167)
	if v2216 == v2215 {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	v2268 = v2213 << (uint(int32(2)) % 32)
	if v2216 == int32(0) {
		goto L377
	} else {
		goto L378
	}
L356:
	;
	v2235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2234)+16)))
	v2236 = v2235 + v2234
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2236)))
	v2238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2234)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v2238) {
		goto L361
	} else {
		goto L362
	}
L357:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2220+(v2167^int32(-1))<<(uint(int32(2))%32))))
	v2234 = v2226
	goto L356
L358:
	;
	goto L359
L359:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2234 = v2228 + v2167<<(uint(int32(13))%32) + int32(-8192)
	goto L356
L360:
	;
	goto L355
L361:
	;
	if base.Ui32(v2213) < base.Ui32(int32(base.Ui32(v2238+int32(_a_F_btvacuumscan_6))>>(uint(int32(2))%32))&int32(_a_F_btvacuumscan_5)) {
		goto L360
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+4))
	F_LockBuffer(m, v2167, int32(0))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L1
	} else {
		goto L365
	}
L364:
	;
	goto L363
L365:
	;
	F_ReleaseBuffer(m, v2167)
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	if v2248 == int32(0) {
		goto L265
	} else {
		goto L367
	}
L367:
	;
	if v2248 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v2258 = int32(2)
	goto L370
L369:
	;
	v2258 = int32(1)
	goto L370
L370:
	;
	if v2258 != v2213 {
		goto L265
	} else {
		goto L371
	}
L371:
	;
	v2260 = F__bt_leftsib_splitflag(m, v291, v2237, v2214)
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	if v2260 != 0 {
		goto L265
	} else {
		goto L373
	}
L373:
	;
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v2165)+8))
	v2263 = F__bt_getstackbuf(m, v291, v2064, v2262, v2214)
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	if v2263 == int32(0) {
		v3759 = v2214
		goto L266
	} else {
		goto L375
	}
L375:
	;
	v2165 = v2262
	v2167 = v2263
	v2174 = v2248
	v2177 = v2214
	goto L354
L376:
	;
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v2268+v2286)+24))
	v2291 = v2288&int32(_a_F_btvacuumscan_7) + v2286
	v2292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2291))))
	v2295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2291)+2)))
	if v2174 != v2292<<(uint(int32(16))%32)|v2295 {
		goto L380
	} else {
		goto L381
	}
L377:
	;
	v2272 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2272+(v2167^int32(-1))<<(uint(int32(2))%32))))
	v2286 = v2278
	goto L376
L378:
	;
	goto L379
L379:
	;
	v2280 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2286 = v2280 + v2167<<(uint(int32(13))%32) + int32(-8192)
	goto L376
L380:
	;
	v2300 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L1
	} else {
		goto L383
	}
L381:
	;
	goto L382
L382:
	;
	F_PredicateLockPageCombine(m, v291, v2102, v2104)
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L1
	} else {
		goto L396
	}
L383:
	;
	if v2300 != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2304 = m.ExcPending
	if v2304 != 0 {
		goto L1
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	F_LockBuffer(m, v2167, int32(0))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L1
	} else {
		goto L394
	}
L387:
	;
	v2305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2291)+2)))
	v2306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2291))))
	if v2167 < int32(0) {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+160)) = v2326 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+156)) = v2325
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+152)) = v2305 | v2306<<(uint(int32(16))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+148)) = v2177
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+144)) = v2174
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_21), v1802+int32(144))
	mBase = m.M
	v2341 = m.ExcPending
	if v2341 != 0 {
		goto L1
	} else {
		goto L392
	}
L389:
	;
	v2310 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(v2310+(v2167^int32(-1))<<(uint(int32(6))%32))+16))
	v2325 = v2316
	goto L388
L390:
	;
	goto L391
L391:
	;
	v2318 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(v2318+v2167<<(uint(int32(6))%32)+int32(-64))+16))
	v2325 = v2324
	goto L388
L392:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2187), int32(_a_F_btvacuumscan_20))
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	goto L386
L394:
	;
	F_ReleaseBuffer(m, v2167)
	mBase = m.M
	v2354 = m.ExcPending
	if v2354 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	goto L265
L396:
	;
	v2357 = int32(_a_F_btvacuumscan_4)
	v2359 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v2359 + int32(1)
	if v2216 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v2268+v2380)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2382&int32(_a_F_btvacuumscan_7)+v2380))) = base.I32_rotr(v2174, int32(16))
	F_PageIndexTupleDelete(m, v2380, (v2213+int32(1))&int32(_a_F_btvacuumscan_5))
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L1
	} else {
		goto L401
	}
L398:
	;
	v2366 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v2366+(v2167^int32(-1))<<(uint(int32(2))%32))))
	v2380 = v2372
	goto L397
L399:
	;
	goto L400
L400:
	;
	v2374 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2380 = v2374 + v2167<<(uint(int32(13))%32) + int32(-8192)
	goto L397
L401:
	;
	if v1875 == int32(0) {
		goto L403
	} else {
		goto L404
	}
L402:
	;
	v2413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2412)+16)))
	v2414 = v2413 + v2412
	v2415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2414)+12)))
	v2417 = v2415 | int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v2414)+12)) = uint16(v2417)
	v2420 = base.B2i32(v2102 == v2177)
	if v2102 == v2177 {
		goto L406
	} else {
		goto L407
	}
L403:
	;
	v2398 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(v2398+(v1833^int32(-1))<<(uint(int32(2))%32))))
	v2412 = v2404
	goto L402
L404:
	;
	goto L405
L405:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2412 = v2406 + v1833<<(uint(int32(13))%32) + int32(-8192)
	goto L402
L406:
	;
	v2421 = int32(-1)
	goto L408
L407:
	;
	v2421 = v2177
	goto L408
L408:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1802)+222)) = uint16(v2421)
	if v2102 == v2177 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v2426 = int32(-1)
	goto L411
L410:
	;
	v2426 = int32(base.Ui32(v2177) >> (uint(int32(16)) % 32))
	goto L411
L411:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1802)+220)) = uint16(v2426)
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+224)) = int32(537395200)
	v2434 = F_PageIndexTupleOverwrite(m, v2412, int32(1), v1802+int32(220), int32(8))
	mBase = m.M
	v2435 = m.ExcPending
	if v2435 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	if v2434 == int32(0) {
		goto L306
	} else {
		goto L413
	}
L413:
	;
	F_MarkBufferDirty(m, v2167)
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	F_MarkBufferDirty(m, v1833)
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	v2443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2442)+118)))
	if v2443 != int32(112) {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v2543 = int32(_a_F_btvacuumscan_4)
	v2545 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v2545 - int32(1)
	F_LockBuffer(m, v2167, int32(0))
	mBase = m.M
	v2551 = m.ExcPending
	if v2551 != 0 {
		goto L1
	} else {
		goto L440
	}
L417:
	;
	v2447 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if v2447 <= int32(0) {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v291)+32))
	if v2450 != 0 {
		goto L416
	} else {
		goto L421
	}
L419:
	;
	goto L420
L420:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1802)+248)) = uint16(v2213)
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+252)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+264)) = v2421
	F_XLogBeginInsert(m)
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L1
	} else {
		goto L423
	}
L421:
	;
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v291)+40))
	if v2451 != 0 {
		goto L416
	} else {
		goto L422
	}
L422:
	;
	goto L420
L423:
	;
	F_XLogRegisterBuffer(m, int32(0), v1833, int32(6))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	F_XLogRegisterBuffer(m, int32(1), v2167, int32(8))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	if v1875 == int32(0) {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	v2483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2482)+16)))
	v2484 = v2483 + v2482
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v2484)))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+256)) = v2485
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v2484)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+260)) = v2487
	F_XLogRegisterData(m, v1802+int32(248), int32(20))
	mBase = m.M
	v2493 = m.ExcPending
	if v2493 != 0 {
		goto L1
	} else {
		goto L430
	}
L427:
	;
	v2468 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v2468+(v1833^int32(-1))<<(uint(int32(2))%32))))
	v2482 = v2474
	goto L426
L428:
	;
	goto L429
L429:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2482 = v2476 + v1833<<(uint(int32(13))%32) + int32(-8192)
	goto L426
L430:
	;
	v2496 = F_XLogInsert(m, int32(11), int32(176))
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	if v2216 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2515))) = base.I64_rotr(v2496, int64(32))
	if v1875 == int32(0) {
		goto L437
	} else {
		goto L438
	}
L433:
	;
	v2501 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v2501+(v2167^int32(-1))<<(uint(int32(2))%32))))
	v2515 = v2507
	goto L432
L434:
	;
	goto L435
L435:
	;
	v2509 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2515 = v2509 + v2167<<(uint(int32(13))%32) + int32(-8192)
	goto L432
L436:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v2536)+4)) = uint32(v2496)
	v2539 = int64(base.Ui64(v2496) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2536))) = uint32(v2539)
	goto L416
L437:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2522+(v1833^int32(-1))<<(uint(int32(2))%32))))
	v2536 = v2528
	goto L436
L438:
	;
	goto L439
L439:
	;
	v2530 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2536 = v2530 + v1833<<(uint(int32(13))%32) + int32(-8192)
	goto L436
L440:
	;
	F_ReleaseBuffer(m, v2167)
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	v2554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1895)+12)))
	v2560 = v2554
	goto L309
L442:
	;
	v2610 = v1833 ^ int32(-1)
	v2612 = v1833 << (uint(int32(13)) % 32)
	goto L445
L443:
	;
	v3613 = int32(0)
	goto L444
L444:
	;
	v3658 = *(*int32)(unsafe.Add(mBase, uint32(v1895)+4))
	F_LockBuffer(m, v1833, int32(0))
	mBase = m.M
	v3661 = m.ExcPending
	if v3661 != 0 {
		goto L1
	} else {
		goto L716
	}
L445:
	;
	if v1833 < int32(0) {
		goto L448
	} else {
		goto L449
	}
L446:
	;
	if base.Ui32(v3157) < base.Ui32(int32(25)) {
		goto L710
	} else {
		goto L711
	}
L447:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+4))
	if v1875 == int32(0) {
		goto L452
	} else {
		goto L453
	}
L448:
	;
	v2667 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[5]))
	v2673 = *(*int32)(unsafe.Add(mBase, uint32(v2667+(v1833^int32(-1))<<(uint(int32(6))%32))+16))
	v2682 = v2673
	goto L447
L449:
	;
	goto L450
L450:
	;
	v2675 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[6]))
	v2681 = *(*int32)(unsafe.Add(mBase, uint32(v2675+v1833<<(uint(int32(6))%32)+int32(-64))+16))
	v2682 = v2681
	goto L447
L451:
	;
	v2698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2697)+16)))
	v2699 = v2698 + v2697
	v2700 = *(*int32)(unsafe.Add(mBase, uint32(v2699)+4))
	v2701 = *(*int32)(unsafe.Add(mBase, uint32(v2699)))
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(v2697)+24))
	v2705 = v2697 + v2702&int32(_a_F_btvacuumscan_7)
	v2706 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2705)+2)))
	v2707 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2705))))
	F_LockBuffer(m, v1833, int32(0))
	mBase = m.M
	v2710 = m.ExcPending
	if v2710 != 0 {
		goto L1
	} else {
		goto L455
	}
L452:
	;
	v2687 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2687+v2610<<(uint(int32(2))%32))))
	v2697 = v2691
	goto L451
L453:
	;
	goto L454
L454:
	;
	v2693 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2697 = v2693 + v2612 + int32(-8192)
	goto L451
L455:
	;
	v2713 = v2706 | v2707<<(uint(int32(16))%32)
	v2715 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v2715 != 0 {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2717 = m.ExcPending
	if v2717 != 0 {
		goto L1
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	v2719 = int32(1)
	if v2713 == int32(-1) {
		goto L461
	} else {
		goto L462
	}
L459:
	;
	goto L458
L460:
	;
	v2766 = int32(0)
	if v2760 == v2766 {
		v2943 = int32(0)
		v2944 = v2766
		goto L476
	} else {
		goto L477
	}
L461:
	;
	v2760 = v2701
	v2762 = v1833
	v2763 = v2682
	v2764 = v2719
	v2765 = int32(0)
	goto L460
L462:
	;
	goto L463
L463:
	;
	v2723 = F_ReadBuffer(m, v291, v2713)
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	F_LockBuffer(m, v2723, int32(1))
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F__bt_checkpage(m, v291, v2723)
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	if v2723 < int32(0) {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	v2748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2747)+16)))
	v2749 = v2748 + v2747
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(v2749)+8))
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2749)))
	F_LockBuffer(m, v2723, int32(0))
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		goto L1
	} else {
		goto L471
	}
L468:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2733+(v2723^int32(-1))<<(uint(int32(2))%32))))
	v2747 = v2739
	goto L467
L469:
	;
	goto L470
L470:
	;
	v2741 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2747 = v2741 + v2723<<(uint(int32(13))%32) + int32(-8192)
	goto L467
L471:
	;
	if v2682 == v2713 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v2760 = v2751
	v2762 = v2723
	v2763 = v2682
	v2764 = v2719
	v2765 = v2750
	goto L460
L473:
	;
	goto L474
L474:
	;
	F_LockBuffer(m, v1833, int32(2))
	mBase = m.M
	v2758 = m.ExcPending
	if v2758 != 0 {
		goto L1
	} else {
		goto L475
	}
L475:
	;
	v2760 = v2751
	v2762 = v2723
	v2763 = v2713
	v2764 = int32(0)
	v2765 = v2750
	goto L460
L476:
	;
	F_LockBuffer(m, v2762, int32(2))
	mBase = m.M
	v2991 = m.ExcPending
	if v2991 != 0 {
		goto L1
	} else {
		goto L523
	}
L477:
	;
	v2769 = F_ReadBuffer(m, v291, v2760)
	mBase = m.M
	v2770 = m.ExcPending
	if v2770 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	F_LockBuffer(m, v2769, int32(2))
	mBase = m.M
	v2773 = m.ExcPending
	if v2773 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	F__bt_checkpage(m, v291, v2769)
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	if v2769 < int32(0) {
		goto L482
	} else {
		goto L483
	}
L481:
	;
	v2794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2793)+16)))
	v2795 = v2794 + v2793
	v2796 = *(*int32)(unsafe.Add(mBase, uint32(v2795)+4))
	v2797 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2795)+12)))
	v2799 = v2797 & int32(4)
	if v2799 == int32(0) {
		goto L485
	} else {
		goto L486
	}
L482:
	;
	v2779 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(v2779+(v2769^int32(-1))<<(uint(int32(2))%32))))
	v2793 = v2785
	goto L481
L483:
	;
	goto L484
L484:
	;
	v2787 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2793 = v2787 + v2769<<(uint(int32(13))%32) + int32(-8192)
	goto L481
L485:
	;
	if v2796 == v2763 {
		v2943 = v2769
		v2944 = v2760
		goto L476
	} else {
		goto L488
	}
L486:
	;
	goto L487
L487:
	;
	__phi2807 = v2760
	__phi2808 = v2796
	__phi2809 = v2769
	__phi2818 = v2799
	v2807 = __phi2807
	v2808 = __phi2808
	v2809 = __phi2809
	v2818 = __phi2818
	goto L489
L488:
	;
	goto L487
L489:
	;
	if v2808 == int32(0) {
		goto L492
	} else {
		goto L493
	}
L490:
	;
	v2943 = v2906
	v2944 = v2808
	goto L476
L491:
	;
	F_LockBuffer(m, v2809, int32(0))
	mBase = m.M
	v2899 = m.ExcPending
	if v2899 != 0 {
		goto L1
	} else {
		goto L508
	}
L492:
	;
	F_LockBuffer(m, v2809, int32(0))
	mBase = m.M
	v2862 = m.ExcPending
	if v2862 != 0 {
		goto L1
	} else {
		goto L496
	}
L493:
	;
	if v2818&int32(_a_F_btvacuumscan_5) != 0 {
		goto L492
	} else {
		goto L494
	}
L494:
	;
	if v2807 != v2808 {
		goto L491
	} else {
		goto L495
	}
L495:
	;
	goto L492
L496:
	;
	F_ReleaseBuffer(m, v2809)
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v2867 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2868 = m.ExcPending
	if v2868 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	if v2867 != 0 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2871 = m.ExcPending
	if v2871 != 0 {
		goto L1
	} else {
		goto L502
	}
L500:
	;
	goto L501
L501:
	;
	F_ReleaseBuffer(m, v2762)
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L1
	} else {
		goto L505
	}
L502:
	;
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+128)) = v2765
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+132)) = v2872 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+124)) = v1822
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+120)) = v2682
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+116)) = v2763
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+112)) = v2808
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_22), v1802+int32(112))
	mBase = m.M
	v2885 = m.ExcPending
	if v2885 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2439), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v2890 = m.ExcPending
	if v2890 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	goto L501
L505:
	;
	if v2764 != 0 {
		goto L263
	} else {
		goto L506
	}
L506:
	;
	F_LockBuffer(m, v1833, int32(0))
	mBase = m.M
	v2896 = m.ExcPending
	if v2896 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	goto L264
L508:
	;
	F_ReleaseBuffer(m, v2809)
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	v2903 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v2903 != 0 {
		goto L510
	} else {
		goto L511
	}
L510:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2905 = m.ExcPending
	if v2905 != 0 {
		goto L1
	} else {
		goto L513
	}
L511:
	;
	goto L512
L512:
	;
	v2906 = F_ReadBuffer(m, v291, v2808)
	mBase = m.M
	v2907 = m.ExcPending
	if v2907 != 0 {
		goto L1
	} else {
		goto L514
	}
L513:
	;
	goto L512
L514:
	;
	F_LockBuffer(m, v2906, int32(2))
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	F__bt_checkpage(m, v291, v2906)
	mBase = m.M
	v2912 = m.ExcPending
	if v2912 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	if v2906 < int32(0) {
		goto L518
	} else {
		goto L519
	}
L517:
	;
	v2931 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2930)+16)))
	v2932 = v2931 + v2930
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v2932)+4))
	v2934 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2932)+12)))
	v2936 = v2934 & int32(4)
	if v2936 != 0 {
		__phi2807 = v2808
		__phi2808 = v2933
		__phi2809 = v2906
		__phi2818 = v2936
		v2807 = __phi2807
		v2808 = __phi2808
		v2809 = __phi2809
		v2818 = __phi2818
		goto L489
	} else {
		goto L521
	}
L518:
	;
	v2916 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v2922 = *(*int32)(unsafe.Add(mBase, uint32(v2916+(v2906^int32(-1))<<(uint(int32(2))%32))))
	v2930 = v2922
	goto L517
L519:
	;
	goto L520
L520:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v2930 = v2924 + v2906<<(uint(int32(13))%32) + int32(-8192)
	goto L517
L521:
	;
	if v2933 != v2763 {
		__phi2807 = v2808
		__phi2808 = v2933
		__phi2809 = v2906
		__phi2818 = v2936
		v2807 = __phi2807
		v2808 = __phi2808
		v2809 = __phi2809
		v2818 = __phi2818
		goto L489
	} else {
		goto L522
	}
L522:
	;
	goto L490
L523:
	;
	v2992 = int32(0)
	v2993 = base.B2i32(v2992 <= v2762)
	if v2993 == v2992 {
		goto L525
	} else {
		goto L526
	}
L524:
	;
	v3012 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3011)+16)))
	v3013 = v3012 + v3011
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(v3013)+4))
	if v3014 == int32(0) {
		goto L269
	} else {
		goto L528
	}
L525:
	;
	v2997 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3003 = *(*int32)(unsafe.Add(mBase, uint32(v2997+(v2762^int32(-1))<<(uint(int32(2))%32))))
	v3011 = v3003
	goto L524
L526:
	;
	goto L527
L527:
	;
	v3005 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3011 = v3005 + v2762<<(uint(int32(13))%32) + int32(-8192)
	goto L524
L528:
	;
	v3017 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3013)+12)))
	if v3017&int32(6) != 0 {
		goto L269
	} else {
		goto L529
	}
L529:
	;
	v3020 = *(*int32)(unsafe.Add(mBase, uint32(v3013)))
	if v3020 != v2944 {
		goto L268
	} else {
		goto L530
	}
L530:
	;
	v3022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3011)+12)))
	v3024 = v3022 + int32(_a_F_btvacuumscan_6)
	if v2764 != 0 {
		goto L532
	} else {
		goto L533
	}
L531:
	;
	v3078 = F_ReadBuffer(m, v291, v3014)
	mBase = m.M
	v3079 = m.ExcPending
	if v3079 != 0 {
		goto L1
	} else {
		goto L548
	}
L532:
	;
	v3025 = int32(17)
	if v3017&v3025 == v3025 {
		goto L535
	} else {
		goto L536
	}
L533:
	;
	goto L534
L534:
	;
	if v3017&int32(1) != 0 {
		goto L267
	} else {
		goto L542
	}
L535:
	;
	if base.B2i32(v3024&int32(_a_F_btvacuumscan_18) == int32(0))|base.B2i32(base.Ui32(v3022) < base.Ui32(int32(25))) != 0 {
		v3077 = int32(-1)
		goto L531
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3040 = m.ExcPending
	if v3040 != 0 {
		goto L1
	} else {
		goto L539
	}
L538:
	;
	goto L537
L539:
	;
	v3041 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+64)) = v2763
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+68)) = v3041 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_24), v1802-int32(-64))
	mBase = m.M
	v3050 = m.ExcPending
	if v3050 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2486), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L542:
	;
	if base.Ui32(v3022) < base.Ui32(int32(25)) {
		goto L267
	} else {
		goto L543
	}
L543:
	;
	if v3024&int32(_a_F_btvacuumscan_25) != int32(8) {
		goto L267
	} else {
		goto L544
	}
L544:
	;
	v3065 = *(*int32)(unsafe.Add(mBase, uint32(v3011)+28))
	v3068 = v3011 + v3065&int32(_a_F_btvacuumscan_7)
	v3069 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3068))))
	v3072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3068)+2)))
	v3073 = v3069<<(uint(int32(16))%32) | v3072
	if v3073 == v2682 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v3075 = int32(-1)
	goto L547
L546:
	;
	v3075 = v3073
	goto L547
L547:
	;
	v3077 = v3075
	goto L531
L548:
	;
	F_LockBuffer(m, v3078, int32(2))
	mBase = m.M
	v3082 = m.ExcPending
	if v3082 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	F__bt_checkpage(m, v291, v3078)
	mBase = m.M
	v3084 = m.ExcPending
	if v3084 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	v3085 = int32(0)
	v3086 = base.B2i32(v3085 <= v3078)
	if v3086 == v3085 {
		goto L552
	} else {
		goto L553
	}
L551:
	;
	v3105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3104)+16)))
	v3106 = v3105 + v3104
	v3107 = *(*int32)(unsafe.Add(mBase, uint32(v3106)))
	if v2763 != v3107 {
		goto L555
	} else {
		goto L556
	}
L552:
	;
	v3090 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3096 = *(*int32)(unsafe.Add(mBase, uint32(v3090+(v3078^int32(-1))<<(uint(int32(2))%32))))
	v3104 = v3096
	goto L551
L553:
	;
	goto L554
L554:
	;
	v3098 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3104 = v3098 + v3078<<(uint(int32(13))%32) + int32(-8192)
	goto L551
L555:
	;
	v3111 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3112 = m.ExcPending
	if v3112 != 0 {
		goto L1
	} else {
		goto L558
	}
L556:
	;
	goto L557
L557:
	;
	v3157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3104)+12)))
	v3158 = int32(0)
	v3161 = *(*int32)(unsafe.Add(mBase, uint32(v3106)+4))
	if v2944|v3161 != 0 {
		v3222 = v3158
		v3223 = v3158
		v3224 = v3158
		goto L576
	} else {
		goto L577
	}
L558:
	;
	if v3111 != 0 {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3115 = m.ExcPending
	if v3115 != 0 {
		goto L1
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	if v2943 != 0 {
		goto L565
	} else {
		goto L566
	}
L562:
	;
	v3116 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(v3106)))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+52)) = v2765
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+48)) = v3117
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+56)) = v3116 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+44)) = v1822
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+40)) = v2682
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+36)) = v2763
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+32)) = v3014
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_26), v1802+int32(32))
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2540), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3136 = m.ExcPending
	if v3136 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	goto L561
L565:
	;
	F_LockBuffer(m, v2943, int32(0))
	mBase = m.M
	v3141 = m.ExcPending
	if v3141 != 0 {
		goto L1
	} else {
		goto L568
	}
L566:
	;
	goto L567
L567:
	;
	F_LockBuffer(m, v3078, int32(0))
	mBase = m.M
	v3146 = m.ExcPending
	if v3146 != 0 {
		goto L1
	} else {
		goto L570
	}
L568:
	;
	F_ReleaseBuffer(m, v2943)
	mBase = m.M
	v3143 = m.ExcPending
	if v3143 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	goto L567
L570:
	;
	F_ReleaseBuffer(m, v3078)
	mBase = m.M
	v3148 = m.ExcPending
	if v3148 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	F_LockBuffer(m, v2762, int32(0))
	mBase = m.M
	v3151 = m.ExcPending
	if v3151 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	F_ReleaseBuffer(m, v2762)
	mBase = m.M
	v3153 = m.ExcPending
	if v3153 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	if v2764 != 0 {
		goto L263
	} else {
		goto L574
	}
L574:
	;
	F_LockBuffer(m, v1833, int32(0))
	mBase = m.M
	v3156 = m.ExcPending
	if v3156 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	goto L264
L576:
	;
	v3225 = int32(_a_F_btvacuumscan_4)
	v3227 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v3227 + int32(1)
	if v2943 != 0 {
		goto L595
	} else {
		goto L596
	}
L577:
	;
	if v3086 == int32(0) {
		goto L579
	} else {
		goto L580
	}
L578:
	;
	v3181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3180)+16)))
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v3181+v3180)+4))
	if v3183 != 0 {
		v3222 = v3158
		v3223 = v3158
		v3224 = v3158
		goto L576
	} else {
		goto L582
	}
L579:
	;
	v3166 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3172 = *(*int32)(unsafe.Add(mBase, uint32(v3166+(v3078^int32(-1))<<(uint(int32(2))%32))))
	v3180 = v3172
	goto L578
L580:
	;
	goto L581
L581:
	;
	v3174 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3180 = v3174 + v3078<<(uint(int32(13))%32) + int32(-8192)
	goto L578
L582:
	;
	v3185 = F_ReadBuffer(m, v291, int32(0))
	mBase = m.M
	v3186 = m.ExcPending
	if v3186 != 0 {
		goto L1
	} else {
		goto L583
	}
L583:
	;
	F_LockBuffer(m, v3185, int32(2))
	mBase = m.M
	v3189 = m.ExcPending
	if v3189 != 0 {
		goto L1
	} else {
		goto L584
	}
L584:
	;
	F__bt_checkpage(m, v291, v3185)
	mBase = m.M
	v3191 = m.ExcPending
	if v3191 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	if v3185 < int32(0) {
		goto L587
	} else {
		goto L588
	}
L586:
	;
	v3211 = v3209 + int32(24)
	v3212 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+44))
	if base.Ui32(v3212) <= base.Ui32(v2765+int32(1)) {
		goto L590
	} else {
		goto L591
	}
L587:
	;
	v3195 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v3195+(v3185^int32(-1))<<(uint(int32(2))%32))))
	v3209 = v3201
	goto L586
L588:
	;
	goto L589
L589:
	;
	v3203 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3209 = v3203 + v3185<<(uint(int32(13))%32) + int32(-8192)
	goto L586
L590:
	;
	v3222 = v3209
	v3223 = v3211
	v3224 = v3185
	goto L576
L591:
	;
	goto L592
L592:
	;
	F_LockBuffer(m, v3185, int32(0))
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	F_ReleaseBuffer(m, v3185)
	mBase = m.M
	v3220 = m.ExcPending
	if v3220 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	v3222 = v3209
	v3223 = v3211
	v3224 = v3158
	goto L576
L595:
	;
	if v2943 < int32(0) {
		goto L599
	} else {
		goto L600
	}
L596:
	;
	goto L597
L597:
	;
	if v3086 == int32(0) {
		goto L603
	} else {
		goto L604
	}
L598:
	;
	v3249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3248)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v3249+v3248)+4)) = v3014
	goto L597
L599:
	;
	v3234 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3240 = *(*int32)(unsafe.Add(mBase, uint32(v3234+(v2943^int32(-1))<<(uint(int32(2))%32))))
	v3248 = v3240
	goto L598
L600:
	;
	goto L601
L601:
	;
	v3242 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3248 = v3242 + v2943<<(uint(int32(13))%32) + int32(-8192)
	goto L598
L602:
	;
	v3271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3270)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v3271+v3270))) = v2944
	if v2764 == int32(0) {
		goto L606
	} else {
		goto L607
	}
L603:
	;
	v3256 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3262 = *(*int32)(unsafe.Add(mBase, uint32(v3256+(v3078^int32(-1))<<(uint(int32(2))%32))))
	v3270 = v3262
	goto L602
L604:
	;
	goto L605
L605:
	;
	v3264 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3270 = v3264 + v3078<<(uint(int32(13))%32) + int32(-8192)
	goto L602
L606:
	;
	v3276 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2705)+4)) = uint16(v3276)
	*(*int32)(unsafe.Add(mBase, uint32(v2705))) = base.I32_rotr(v3077, int32(16))
	v3281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2705)+6)))
	v3283 = v3281 | int32(_a_F_btvacuumscan_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2705)+6)) = uint16(v3283)
	goto L608
L607:
	;
	goto L608
L608:
	;
	if v2993 == int32(0) {
		goto L610
	} else {
		goto L611
	}
L609:
	;
	v3303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3302)+16)))
	v3304 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		goto L1
	} else {
		goto L613
	}
L610:
	;
	v3288 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(v3288+(v2762^int32(-1))<<(uint(int32(2))%32))))
	v3302 = v3294
	goto L609
L611:
	;
	goto L612
L612:
	;
	v3296 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3302 = v3296 + v2762<<(uint(int32(13))%32) + int32(-8192)
	goto L609
L613:
	;
	v3306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3302)+16)))
	v3307 = v3302 + v3306
	v3308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3307)+12)))
	v3312 = v3308&int32(_a_F_btvacuumscan_27) | int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v3307)+12)) = uint16(v3312)
	v3314 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v3302)+12)) = uint16(v3314)
	*(*int64)(unsafe.Add(mBase, uint32(v3302)+24)) = v3304
	v3317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3302)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3302)+14)) = uint16(v3317)
	v3320 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3302+v3303)+14)) = uint16(v3320)
	if v3224 != 0 {
		goto L614
	} else {
		goto L615
	}
L614:
	;
	v3322 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+4))
	if base.Ui32(v3322) <= base.Ui32(int32(2)) {
		goto L617
	} else {
		goto L618
	}
L615:
	;
	goto L616
L616:
	;
	F_MarkBufferDirty(m, v3078)
	mBase = m.M
	v3342 = m.ExcPending
	if v3342 != 0 {
		goto L1
	} else {
		goto L621
	}
L617:
	;
	v3325 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v3222)+12)) = uint16(v3325)
	v3329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3222-int32(-64)))) = uint8(v3329)
	*(*int64)(unsafe.Add(mBase, uint32(v3222)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v3222)+48)) = v3329
	*(*int32)(unsafe.Add(mBase, uint32(v3222)+28)) = int32(3)
	goto L619
L618:
	;
	goto L619
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3223)+20)) = v2765
	*(*int32)(unsafe.Add(mBase, uint32(v3223)+16)) = v3014
	F_MarkBufferDirty(m, v3224)
	mBase = m.M
	v3340 = m.ExcPending
	if v3340 != 0 {
		goto L1
	} else {
		goto L620
	}
L620:
	;
	goto L616
L621:
	;
	F_MarkBufferDirty(m, v2762)
	mBase = m.M
	v3344 = m.ExcPending
	if v3344 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	if v2943 != 0 {
		goto L623
	} else {
		goto L624
	}
L623:
	;
	F_MarkBufferDirty(m, v2943)
	mBase = m.M
	v3346 = m.ExcPending
	if v3346 != 0 {
		goto L1
	} else {
		goto L626
	}
L624:
	;
	goto L625
L625:
	;
	if v2764 == int32(0) {
		goto L627
	} else {
		goto L628
	}
L626:
	;
	goto L625
L627:
	;
	F_MarkBufferDirty(m, v1833)
	mBase = m.M
	v3350 = m.ExcPending
	if v3350 != 0 {
		goto L1
	} else {
		goto L630
	}
L628:
	;
	goto L629
L629:
	;
	v3351 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	v3352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3351)+118)))
	if v3352 != int32(112) {
		goto L631
	} else {
		goto L632
	}
L630:
	;
	goto L629
L631:
	;
	v3517 = int32(_a_F_btvacuumscan_4)
	v3519 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v3519 - int32(1)
	if v3224 != 0 {
		goto L678
	} else {
		goto L679
	}
L632:
	;
	v3356 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[9]))
	if v3356 <= int32(0) {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v3359 = *(*int32)(unsafe.Add(mBase, uint32(v291)+32))
	if v3359 != 0 {
		goto L631
	} else {
		goto L636
	}
L634:
	;
	goto L635
L635:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v3362 = m.ExcPending
	if v3362 != 0 {
		goto L1
	} else {
		goto L638
	}
L636:
	;
	v3360 = *(*int32)(unsafe.Add(mBase, uint32(v291)+40))
	if v3360 != 0 {
		goto L631
	} else {
		goto L637
	}
L637:
	;
	goto L635
L638:
	;
	F_XLogRegisterBuffer(m, int32(0), v2762, int32(6))
	mBase = m.M
	v3366 = m.ExcPending
	if v3366 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	if v2943 != 0 {
		goto L640
	} else {
		goto L641
	}
L640:
	;
	F_XLogRegisterBuffer(m, int32(1), v2943, int32(8))
	mBase = m.M
	v3370 = m.ExcPending
	if v3370 != 0 {
		goto L1
	} else {
		goto L643
	}
L641:
	;
	goto L642
L642:
	;
	F_XLogRegisterBuffer(m, int32(2), v3078, int32(8))
	mBase = m.M
	v3374 = m.ExcPending
	if v3374 != 0 {
		goto L1
	} else {
		goto L644
	}
L643:
	;
	goto L642
L644:
	;
	if v2764 == int32(0) {
		goto L645
	} else {
		goto L646
	}
L645:
	;
	F_XLogRegisterBuffer(m, int32(3), v1833, int32(6))
	mBase = m.M
	v3380 = m.ExcPending
	if v3380 != 0 {
		goto L1
	} else {
		goto L648
	}
L646:
	;
	goto L647
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+280)) = v3077
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+276)) = v2700
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+272)) = v2701
	*(*int64)(unsafe.Add(mBase, uint32(v1802)+264)) = v3304
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+256)) = v2765
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+252)) = v3014
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+248)) = v2944
	F_XLogRegisterData(m, v1802+int32(248), int32(36))
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L1
	} else {
		goto L649
	}
L648:
	;
	goto L647
L649:
	;
	if v3224 == int32(0) {
		goto L651
	} else {
		goto L652
	}
L650:
	;
	if v3086 == int32(0) {
		goto L659
	} else {
		goto L660
	}
L651:
	;
	v3397 = F_XLogInsert(m, int32(11), int32(128))
	mBase = m.M
	v3398 = m.ExcPending
	if v3398 != 0 {
		goto L1
	} else {
		goto L654
	}
L652:
	;
	goto L653
L653:
	;
	F_XLogRegisterBuffer(m, int32(4), v3224, int32(14))
	mBase = m.M
	v3402 = m.ExcPending
	if v3402 != 0 {
		goto L1
	} else {
		goto L655
	}
L654:
	;
	v3430 = v3397
	goto L650
L655:
	;
	v3403 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+220)) = v3403
	v3405 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+224)) = v3405
	v3407 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+228)) = v3407
	v3409 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+232)) = v3409
	v3411 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+236)) = v3411
	v3413 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+240)) = v3413
	v3415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3223)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1802)+244)) = uint8(v3415)
	F_XLogRegisterBufData(m, int32(4), v1802+int32(220), int32(28))
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L1
	} else {
		goto L656
	}
L656:
	;
	v3425 = F_XLogInsert(m, int32(11), int32(144))
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L1
	} else {
		goto L657
	}
L657:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3222))) = base.I64_rotr(v3425, int64(32))
	v3430 = v3425
	goto L650
L658:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3448))) = base.I64_rotr(v3430, int64(32))
	if v2993 == int32(0) {
		goto L663
	} else {
		goto L664
	}
L659:
	;
	v3434 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3440 = *(*int32)(unsafe.Add(mBase, uint32(v3434+(v3078^int32(-1))<<(uint(int32(2))%32))))
	v3448 = v3440
	goto L658
L660:
	;
	goto L661
L661:
	;
	v3442 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3448 = v3442 + v3078<<(uint(int32(13))%32) + int32(-8192)
	goto L658
L662:
	;
	v3470 = base.I32_wrap_i64(v3430)
	*(*int32)(unsafe.Add(mBase, uint32(v3469)+4)) = v3470
	v3474 = base.I32_wrap_i64(int64(base.Ui64(v3430) >> (uint(int64(32)) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v3469))) = v3474
	if v2943 != 0 {
		goto L666
	} else {
		goto L667
	}
L663:
	;
	v3455 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3461 = *(*int32)(unsafe.Add(mBase, uint32(v3455+(v2762^int32(-1))<<(uint(int32(2))%32))))
	v3469 = v3461
	goto L662
L664:
	;
	goto L665
L665:
	;
	v3463 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3469 = v3463 + v2762<<(uint(int32(13))%32) + int32(-8192)
	goto L662
L666:
	;
	if v2943 < int32(0) {
		goto L670
	} else {
		goto L671
	}
L667:
	;
	goto L668
L668:
	;
	if v2764 != 0 {
		goto L631
	} else {
		goto L673
	}
L669:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3493)+4)) = v3470
	*(*int32)(unsafe.Add(mBase, uint32(v3493))) = v3474
	goto L668
L670:
	;
	v3479 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3479+(v2943^int32(-1))<<(uint(int32(2))%32))))
	v3493 = v3485
	goto L669
L671:
	;
	goto L672
L672:
	;
	v3487 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3493 = v3487 + v2943<<(uint(int32(13))%32) + int32(-8192)
	goto L669
L673:
	;
	if v1875 == int32(0) {
		goto L675
	} else {
		goto L676
	}
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3510)+4)) = v3470
	*(*int32)(unsafe.Add(mBase, uint32(v3510))) = v3474
	goto L631
L675:
	;
	v3500 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[7]))
	v3504 = *(*int32)(unsafe.Add(mBase, uint32(v3500+v2610<<(uint(int32(2))%32))))
	v3510 = v3504
	goto L674
L676:
	;
	goto L677
L677:
	;
	v3506 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[8]))
	v3510 = v3506 + v2612 + int32(-8192)
	goto L674
L678:
	;
	F_LockBuffer(m, v3224, int32(0))
	mBase = m.M
	v3525 = m.ExcPending
	if v3525 != 0 {
		goto L1
	} else {
		goto L681
	}
L679:
	;
	goto L680
L680:
	;
	if v2943 != 0 {
		goto L683
	} else {
		goto L684
	}
L681:
	;
	F_ReleaseBuffer(m, v3224)
	mBase = m.M
	v3527 = m.ExcPending
	if v3527 != 0 {
		goto L1
	} else {
		goto L682
	}
L682:
	;
	goto L680
L683:
	;
	F_LockBuffer(m, v2943, int32(0))
	mBase = m.M
	v3530 = m.ExcPending
	if v3530 != 0 {
		goto L1
	} else {
		goto L686
	}
L684:
	;
	goto L685
L685:
	;
	F_LockBuffer(m, v3078, int32(0))
	mBase = m.M
	v3535 = m.ExcPending
	if v3535 != 0 {
		goto L1
	} else {
		goto L688
	}
L686:
	;
	F_ReleaseBuffer(m, v2943)
	mBase = m.M
	v3532 = m.ExcPending
	if v3532 != 0 {
		goto L1
	} else {
		goto L687
	}
L687:
	;
	goto L685
L688:
	;
	F_ReleaseBuffer(m, v3078)
	mBase = m.M
	v3537 = m.ExcPending
	if v3537 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	if v2764 == int32(0) {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	F_LockBuffer(m, v2762, int32(0))
	mBase = m.M
	v3542 = m.ExcPending
	if v3542 != 0 {
		goto L1
	} else {
		goto L693
	}
L691:
	;
	goto L692
L692:
	;
	v3545 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2683)+24)) = v3545 + int32(1)
	if base.Ui32(v2763) <= base.Ui32(v1822) {
		goto L695
	} else {
		goto L696
	}
L693:
	;
	F_ReleaseBuffer(m, v2762)
	mBase = m.M
	v3544 = m.ExcPending
	if v3544 != 0 {
		goto L1
	} else {
		goto L694
	}
L694:
	;
	goto L692
L695:
	;
	v3550 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2683)+28)) = v3550 + int32(1)
	goto L697
L696:
	;
	goto L697
L697:
	;
	v3554 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+36))
	v3555 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+28))
	if v3554 != v3555 {
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+24))
	if v3557 != v3554 {
		goto L702
	} else {
		goto L703
	}
L699:
	;
	goto L700
L700:
	;
	v3590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895)+12)))
	if v3590&int32(16) != 0 {
		goto L445
	} else {
		goto L709
	}
L701:
	;
	v3574 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v3573+v3572<<(uint(v3574)%32)))) = v2763
	v3578 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+32))
	v3579 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v3578+v3579<<(uint(v3574)%32))+8)) = v3304
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+36)) = v3584 + int32(1)
	goto L700
L702:
	;
	v3559 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+32))
	v3572 = v3554
	v3573 = v3559
	goto L701
L703:
	;
	goto L704
L704:
	;
	v3561 = v3554 << (uint(int32(1)) % 32)
	if v3561 < v3555 {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	v3563 = v3561
	goto L707
L706:
	;
	v3563 = v3555
	goto L707
L707:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+24)) = v3563
	v3565 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+32))
	v3568 = F_repalloc(m, v3565, v3563<<(uint(int32(4))%32))
	mBase = m.M
	v3569 = m.ExcPending
	if v3569 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+32)) = v3568
	v3571 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+36))
	v3572 = v3571
	v3573 = v3568
	goto L701
L709:
	;
	goto L446
L710:
	;
	v3606 = int32(1)
	goto L712
L711:
	;
	v3598 = int32(2)
	if v3161 != 0 {
		goto L713
	} else {
		goto L714
	}
L712:
	;
	v3613 = v3606
	goto L444
L713:
	;
	v3604 = v3598
	goto L715
L714:
	;
	v3604 = int32(1)
	goto L715
L715:
	;
	v3606 = base.B2i32(base.Ui32(int32(base.Ui32(v3157+int32(_a_F_btvacuumscan_6))>>(uint(v3598)%32))&int32(_a_F_btvacuumscan_5)) < base.Ui32(v3604))
	goto L712
L716:
	;
	F_ReleaseBuffer(m, v1833)
	mBase = m.M
	v3663 = m.ExcPending
	if v3663 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	v3665 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[10]))
	if v3665 != 0 {
		goto L718
	} else {
		goto L719
	}
L718:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3667 = m.ExcPending
	if v3667 != 0 {
		goto L1
	} else {
		goto L721
	}
L719:
	;
	goto L720
L720:
	;
	if v3613 == int32(0) {
		goto L263
	} else {
		goto L722
	}
L721:
	;
	goto L720
L722:
	;
	v3670 = F_ReadBuffer(m, v291, v3658)
	mBase = m.M
	v3671 = m.ExcPending
	if v3671 != 0 {
		goto L1
	} else {
		goto L723
	}
L723:
	;
	F_LockBuffer(m, v3670, int32(2))
	mBase = m.M
	v3674 = m.ExcPending
	if v3674 != 0 {
		goto L1
	} else {
		goto L724
	}
L724:
	;
	F__bt_checkpage(m, v291, v3670)
	mBase = m.M
	v3676 = m.ExcPending
	if v3676 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	v1833 = v3670
	goto L270
L726:
	;
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_28), int32(0))
	mBase = m.M
	v3684 = m.ExcPending
	if v3684 != 0 {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2244), int32(_a_F_btvacuumscan_20))
	mBase = m.M
	v3689 = m.ExcPending
	if v3689 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L729:
	;
	v3695 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+16)) = v2763
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+20)) = v3695 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_29), v1802+int32(16))
	mBase = m.M
	v3704 = m.ExcPending
	if v3704 != 0 {
		goto L1
	} else {
		goto L730
	}
L730:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2472), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3709 = m.ExcPending
	if v3709 != 0 {
		goto L1
	} else {
		goto L731
	}
L731:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L732:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3716 = m.ExcPending
	if v3716 != 0 {
		goto L1
	} else {
		goto L733
	}
L733:
	;
	v3717 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	v3718 = *(*int32)(unsafe.Add(mBase, uint32(v3013)))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+104)) = v2763
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+100)) = v3718
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+96)) = v2944
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+108)) = v3717 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_30), v1802+int32(96))
	mBase = m.M
	v3729 = m.ExcPending
	if v3729 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2479), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3734 = m.ExcPending
	if v3734 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L736:
	;
	v3739 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+84)) = v2763
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+80)) = v2765
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+88)) = v3739 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_31), v1802+int32(80))
	mBase = m.M
	v3749 = m.ExcPending
	if v3749 != 0 {
		goto L1
	} else {
		goto L737
	}
L737:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2498), int32(_a_F_btvacuumscan_23))
	mBase = m.M
	v3754 = m.ExcPending
	if v3754 != 0 {
		goto L1
	} else {
		goto L738
	}
L738:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L739:
	;
	if v3808 == int32(0) {
		goto L265
	} else {
		goto L740
	}
L740:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3814 = m.ExcPending
	if v3814 != 0 {
		goto L1
	} else {
		goto L741
	}
L741:
	;
	v3815 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+4)) = v3759
	*(*int32)(unsafe.Add(mBase, uint32(v1802))) = v3815 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_32), v1802)
	mBase = m.M
	v3822 = m.ExcPending
	if v3822 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_11), int32(2848), int32(_a_F_btvacuumscan_33))
	mBase = m.M
	v3827 = m.ExcPending
	if v3827 != 0 {
		goto L1
	} else {
		goto L743
	}
L743:
	;
	goto L265
L744:
	;
	goto L264
L745:
	;
	goto L263
L746:
	;
	v4074 = v4021
	goto L54
L747:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v4099 = m.ExcPending
	if v4099 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	v4100 = int32(0)
	v4102 = *(*int32)(unsafe.Add(mBase, uint32(v289)+24))
	v4103 = F_ReadBufferExtended(m, v291, v4100, v4074, v4100, v4102)
	mBase = m.M
	v4104 = m.ExcPending
	if v4104 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	v317 = v4103
	v320 = v4074
	goto L51
L750:
	;
	v125 = v232
	v126 = v233
	v136 = v243
	v150 = v257
	v158 = v265
	v161 = v268
	goto L17
L751:
	;
	if v4111 == int32(0) {
		goto L41
	} else {
		goto L752
	}
L752:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v4117 = m.ExcPending
	if v4117 != 0 {
		goto L1
	} else {
		goto L753
	}
L753:
	;
	v4118 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v4118 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btvacuumscan_34), v243)
	mBase = m.M
	v4126 = m.ExcPending
	if v4126 != 0 {
		goto L1
	} else {
		goto L754
	}
L754:
	;
	F_errfinish(m, int32(_a_F_btvacuumscan_35), int32(1413), int32(_a_F_btvacuumscan_36))
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L1
	} else {
		goto L755
	}
L755:
	;
	goto L41
L756:
	;
	goto L40
L757:
	;
	v4194 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[2]))
	if v4194 == int32(0) {
		goto L759
	} else {
		goto L760
	}
L758:
	;
	goto L37
L759:
	;
	goto L758
L760:
	;
	v4198 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_btvacuumscan[3])))
	if v4198 != int32(1) {
		goto L759
	} else {
		goto L761
	}
L761:
	;
	v4201 = int32(_a_F_btvacuumscan_4)
	v4203 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	v4204 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v4203 + v4204
	v4207 = *(*int32)(unsafe.Add(mBase, uint32(v4194)))
	*(*int32)(unsafe.Add(mBase, uint32(v4194))) = v4207 + v4204
	*(*int64)(unsafe.Add(mBase, uint32(v4194+int32(128))+232)) = base.I64_extend_i32_u(v313)
	v4215 = *(*int32)(unsafe.Add(mBase, uint32(v4194)))
	*(*int32)(unsafe.Add(mBase, uint32(v4194))) = v4215 + v4204
	v4221 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumscan[4])) = v4221 - v4204
	goto L759
L762:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v190
	v4228 = *(*int32)(unsafe.Add(mBase, uint32(v136)+44))
	F_MemoryContextDelete(m, v4228)
	mBase = m.M
	v4230 = m.ExcPending
	if v4230 != 0 {
		goto L1
	} else {
		goto L763
	}
L763:
	;
	v4231 = int32(0)
	v4233 = v136 + int32(24)
	v4234 = *(*int32)(unsafe.Add(mBase, uint32(v4233)+36))
	if v4234 == v4231 {
		goto L766
	} else {
		goto L767
	}
L764:
	;
	v4473 = *(*int32)(unsafe.Add(mBase, uint32(v126)+32))
	if v4473 != 0 {
		goto L780
	} else {
		goto L781
	}
L765:
	;
	F_pfree(m, v4371)
	mBase = m.M
	v4421 = m.ExcPending
	if v4421 != 0 {
		goto L1
	} else {
		goto L779
	}
L766:
	;
	v4237 = *(*int32)(unsafe.Add(mBase, uint32(v4233)+32))
	if v4237 != 0 {
		v4371 = v4237
		goto L765
	} else {
		goto L769
	}
L767:
	;
	goto L768
L768:
	;
	v4238 = *(*int32)(unsafe.Add(mBase, uint32(v4233)+4))
	v4239 = *(*int32)(unsafe.Add(mBase, uint32(v4233)))
	v4240 = *(*int32)(unsafe.Add(mBase, uint32(v4239)+4))
	v4241 = F_GetOldestNonRemovableTransactionId(m, v4240)
	mBase = m.M
	v4242 = m.ExcPending
	if v4242 != 0 {
		goto L1
	} else {
		goto L770
	}
L769:
	;
	goto L764
L770:
	;
	v4243 = *(*int32)(unsafe.Add(mBase, uint32(v4233)+36))
	if v4243 <= int32(0) {
		goto L771
	} else {
		goto L772
	}
L771:
	;
	v4368 = *(*int32)(unsafe.Add(mBase, uint32(v4233)+32))
	v4371 = v4368
	goto L765
L772:
	;
	v4248 = v4231
	goto L773
L773:
	;
	v4297 = *(*int32)(unsafe.Add(mBase, uint32(v4233)+32))
	v4300 = v4297 + v4248<<(uint(int32(4))%32)
	v4301 = *(*int32)(unsafe.Add(mBase, uint32(v4300)))
	v4302 = *(*int64)(unsafe.Add(mBase, uint32(v4300)+8))
	v4303 = F_GlobalVisCheckRemovableFullXid(m, v4240, v4302)
	mBase = m.M
	v4304 = m.ExcPending
	if v4304 != 0 {
		goto L1
	} else {
		goto L775
	}
L774:
	;
	goto L771
L775:
	;
	if v4303 == int32(0) {
		goto L771
	} else {
		goto L776
	}
L776:
	;
	F_RecordFreeIndexPage(m, v150, v4301)
	mBase = m.M
	v4308 = m.ExcPending
	if v4308 != 0 {
		goto L1
	} else {
		goto L777
	}
L777:
	;
	v4309 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+32))
	v4310 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+32)) = v4309 + v4310
	v4314 = v4248 + v4310
	v4315 = *(*int32)(unsafe.Add(mBase, uint32(v4233)+36))
	if v4314 < v4315 {
		v4248 = v4314
		goto L773
	} else {
		goto L778
	}
L778:
	;
	goto L774
L779:
	;
	goto L764
L780:
	;
	F_FreeSpaceMapVacuum(m, v150)
	mBase = m.M
	v4475 = m.ExcPending
	if v4475 != 0 {
		goto L1
	} else {
		goto L783
	}
L781:
	;
	goto L782
L782:
	;
	m.G0 = v136 + int32(2512)
	return
L783:
	;
	goto L782
}
func F_build_attrmap_by_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
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
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
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
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v217 int32
	_ = v217
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
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
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
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = F_palloc0(m, int32(8))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v24
	v33 = F_palloc0(m, v24<<(uint(int32(1))%32))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v33
	if int32(0) < v24 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L44
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L37
	}
L6:
	;
	v38 = int32(20)
	v49 = int32(0)
	v51 = int32(-1)
	v54 = v33
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v21 + int32(32)
	return v26
L9:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v67 = l1 + v38 + v61<<(uint(int32(4))%32) + v49*int32(100)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+91)))
	if v68 != 0 {
		v179 = v51
		v182 = v54
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v190 = v49 + int32(1)
	if v190 != v24 {
		v49 = v190
		v51 = v179
		v54 = v182
		goto L9
	} else {
		goto L36
	}
L12:
	;
	v70 = v67 + int32(4)
	if v23 <= int32(0) {
		v155 = v51
		v158 = v54
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if l2 != 0 {
		v179 = v155
		v182 = v158
		goto L11
	} else {
		goto L34
	}
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v85 = int32(0)
	v88 = v51
	goto L15
L15:
	;
	v99 = v88 + int32(1)
	if v99 < v23 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v155 = v102
	v158 = v54
	goto L13
L17:
	;
	v145 = v85 + int32(1)
	if v145 != v23 {
		v85 = v145
		v88 = v102
		goto L15
	} else {
		goto L33
	}
L18:
	;
	v102 = v99
	goto L20
L19:
	;
	v102 = int32(0)
	goto L20
L20:
	;
	v105 = l0 + v38 + v75<<(uint(int32(4))%32) + v102*int32(100)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+91)))
	if v106 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v108 = v105 + int32(4)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v112 == int32(0) {
		v131 = v111
		v132 = v112
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v132-v131 != 0 {
		goto L17
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	if v111 != v112 {
		v131 = v111
		v132 = v112
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v116 = v70
	v117 = v108
	goto L26
L26:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v121 == int32(0) {
		v131 = v120
		v132 = v121
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v131 = v120
	v132 = v121
	goto L23
L28:
	;
	v124 = int32(1)
	if v120 == v121 {
		v116 = v116 + v124
		v117 = v117 + v124
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v105)+68))
	if v74 != v134 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v105)+76))
	if v73 != v136 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v54+v49<<(uint(int32(1))%32)))) = uint16(v141)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v155 = v102
	v158 = v143
	goto L13
L33:
	;
	goto L16
L34:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158+v49<<(uint(int32(1))%32)))))
	if v168 == int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v179 = v155
	v182 = v158
	goto L11
L36:
	;
	goto L10
L37:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_build_attrmap_by_name_0), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v226 = F_format_type_be(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v229 = F_format_type_be(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v70
	F_errdetail(m, int32(_a_F_build_attrmap_by_name_1), v21+int32(16))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_build_attrmap_by_name_2), int32(235), int32(_a_F_build_attrmap_by_name_3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_build_attrmap_by_name_0), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v256 = F_format_type_be(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v259 = F_format_type_be(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v70
	F_errdetail(m, int32(_a_F_build_attrmap_by_name_4), v21)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_build_attrmap_by_name_2), int32(247), int32(_a_F_build_attrmap_by_name_3))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
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
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v55 int32
	_ = v55
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
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
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
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
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v633 int32
	_ = v633
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v731 int32
	_ = v731
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v872 int32
	_ = v872
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
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
	var v895 int32
	_ = v895
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
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v964 int32
	_ = v964
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1019 int32
	_ = v1019
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1085 int32
	_ = v1085
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
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
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L107
	} else {
		goto L254
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
	if v35 != int32(1) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v32 != v29 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+217)))
	if v39&int32(1) == int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v45 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v45
	if l5 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l2)+232))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+232)) = v837
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v840 = int32(*(*int16)(unsafe.Add(mBase, uint32(v837)+2)))
	v842 = v840 << (uint(int32(2)) % 32)
	v843 = F_palloc0(m, v842)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L107
	} else {
		goto L209
	}
L10:
	;
	v624 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	if v624 <= int32(0) {
		goto L2
	} else {
		goto L173
	}
L11:
	;
	v622 = v7
	goto L10
L12:
	;
	goto L13
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v55 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v622 = v7
	goto L10
L15:
	;
	goto L16
L16:
	;
	v75 = v7
	v80 = v7
	goto L17
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v75<<(uint(int32(2))%32))))
	if int32(1)<<(uint(v44)%32)&int32(174) != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v622 = v599
	goto L10
L19:
	;
	v601 = v75 + int32(1)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v601 < v602 {
		v75 = v601
		v80 = v599
		goto L17
	} else {
		goto L172
	}
L20:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+8)))
	if v87 != 0 {
		v599 = v80
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+9)))
	if v146 != int32(1) {
		v599 = v80
		goto L19
	} else {
		goto L39
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+32))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v90 = int32(0)
	if v88 == v90 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v143 == int32(0) {
		v599 = v80
		goto L19
	} else {
		goto L38
	}
L25:
	;
	v143 = int32(1)
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v89 == int32(0) {
		v134 = v90
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v143 = v134
	goto L24
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v100 < v99 {
		v134 = v90
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v102 = int32(1)
	if v99 <= v102 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v105 = v102
	goto L33
L32:
	;
	v105 = v99
	goto L33
L33:
	;
	v106 = int32(8)
	v111 = int32(0)
	goto L34
L34:
	;
	v118 = v111 << (uint(int32(2)) % 32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v88+v106+v118)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+(v89+v106))))
	v125 = v120 & (v122 ^ int32(-1))
	v127 = base.B2i32(v125 == int32(0))
	if v125 != 0 {
		v134 = v127
		goto L28
	} else {
		goto L36
	}
L35:
	;
	v134 = v127
	goto L28
L36:
	;
	v129 = v111 + int32(1)
	if v129 != v105 {
		v111 = v129
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L22
L39:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
	if v149 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v86)+124))
	if v152 == int32(0) {
		v599 = v80
		goto L19
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v158 = int32(0)
	if v156 == v158 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	goto L42
L44:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v401 = F_op_strict(m, v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L107
	} else {
		goto L108
	}
L45:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v278 = int32(0)
	if v276 == v278 {
		goto L77
	} else {
		goto L78
	}
L46:
	;
	if v211 == int32(0) {
		goto L45
	} else {
		goto L60
	}
L47:
	;
	v211 = int32(1)
	goto L46
L48:
	;
	goto L49
L49:
	;
	if v157 == int32(0) {
		v202 = v158
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v211 = v202
	goto L46
L51:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v168 < v167 {
		v202 = v158
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v170 = int32(1)
	if v167 <= v170 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v173 = v170
	goto L55
L54:
	;
	v173 = v167
	goto L55
L55:
	;
	v174 = int32(8)
	v179 = int32(0)
	goto L56
L56:
	;
	v186 = v179 << (uint(int32(2)) % 32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v156+v174+v186)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186+(v157+v174))))
	v193 = v188 & (v190 ^ int32(-1))
	v195 = base.B2i32(v193 == int32(0))
	if v193 != 0 {
		v202 = v195
		goto L50
	} else {
		goto L58
	}
L57:
	;
	v202 = v195
	goto L50
L58:
	;
	v197 = v179 + int32(1)
	if v197 != v173 {
		v179 = v197
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v86)+48))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v216 = int32(0)
	if v214 == v216 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v269 == int32(0) {
		goto L45
	} else {
		goto L75
	}
L62:
	;
	v269 = int32(1)
	goto L61
L63:
	;
	goto L64
L64:
	;
	if v215 == int32(0) {
		v260 = v216
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v269 = v260
	goto L61
L66:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v226 < v225 {
		v260 = v216
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v228 = int32(1)
	if v225 <= v228 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v231 = v228
	goto L70
L69:
	;
	v231 = v225
	goto L70
L70:
	;
	v232 = int32(8)
	v237 = int32(0)
	goto L71
L71:
	;
	v244 = v237 << (uint(int32(2)) % 32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v214+v232+v244)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+(v215+v232))))
	v251 = v246 & (v248 ^ int32(-1))
	v253 = base.B2i32(v251 == int32(0))
	if v251 != 0 {
		v260 = v253
		goto L65
	} else {
		goto L73
	}
L72:
	;
	v260 = v253
	goto L65
L73:
	;
	v255 = v237 + int32(1)
	if v255 != v231 {
		v237 = v255
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v396 = v273 + int32(4)
	v397 = v273
	goto L44
L76:
	;
	if v331 == int32(0) {
		v599 = v80
		goto L19
	} else {
		goto L90
	}
L77:
	;
	v331 = int32(1)
	goto L76
L78:
	;
	goto L79
L79:
	;
	if v277 == int32(0) {
		v322 = v278
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v331 = v322
	goto L76
L81:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v288 < v287 {
		v322 = v278
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v290 = int32(1)
	if v287 <= v290 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v293 = v290
	goto L85
L84:
	;
	v293 = v287
	goto L85
L85:
	;
	v294 = int32(8)
	v299 = int32(0)
	goto L86
L86:
	;
	v306 = v299 << (uint(int32(2)) % 32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v276+v294+v306)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306+(v277+v294))))
	v313 = v308 & (v310 ^ int32(-1))
	v315 = base.B2i32(v313 == int32(0))
	if v313 != 0 {
		v322 = v315
		goto L80
	} else {
		goto L88
	}
L87:
	;
	v322 = v315
	goto L80
L88:
	;
	v317 = v299 + int32(1)
	if v317 != v293 {
		v299 = v317
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v86)+48))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v336 = int32(0)
	if v334 == v336 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v389 == int32(0) {
		v599 = v80
		goto L19
	} else {
		goto L105
	}
L92:
	;
	v389 = int32(1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	if v335 == int32(0) {
		v380 = v336
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v389 = v380
	goto L91
L96:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v346 < v345 {
		v380 = v336
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v348 = int32(1)
	if v345 <= v348 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v351 = v348
	goto L100
L99:
	;
	v351 = v345
	goto L100
L100:
	;
	v352 = int32(8)
	v357 = int32(0)
	goto L101
L101:
	;
	v364 = v357 << (uint(int32(2)) % 32)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v334+v352+v364)))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364+(v335+v352))))
	v371 = v366 & (v368 ^ int32(-1))
	v373 = base.B2i32(v371 == int32(0))
	if v371 != 0 {
		v380 = v373
		goto L95
	} else {
		goto L103
	}
L102:
	;
	v380 = v373
	goto L95
L103:
	;
	v375 = v357 + int32(1)
	if v375 != v351 {
		v357 = v375
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	v396 = v393
	v397 = v393 + int32(4)
	goto L44
L106:
	;
	v512 = F_match_expr_to_partition_keys(m, v510, l2, v401)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L107
	} else {
		goto L144
	}
L107:
	;
	return
L108:
	;
	if v401 == int32(0) {
		v510 = v399
		v511 = v398
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v407 = int32(0)
	if v405 == v407 {
		v448 = v407
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v448 != 0 {
		goto L124
	} else {
		goto L125
	}
L111:
	;
	goto L110
L112:
	;
	if v406 == int32(0) {
		v448 = v407
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v416 < v417 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v419 = v416
	goto L116
L115:
	;
	v419 = v417
	goto L116
L116:
	;
	if v419 <= int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v422 = int32(1)
	goto L119
L118:
	;
	v422 = v419
	goto L119
L119:
	;
	v423 = int32(8)
	v428 = int32(0)
	goto L120
L120:
	;
	v435 = v428 << (uint(int32(2)) % 32)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v406+v423+v435)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v435+(v405+v423))))
	v440 = v437 & v439
	v442 = base.B2i32(v440 != int32(0))
	if v440 != 0 {
		v448 = v442
		goto L111
	} else {
		goto L122
	}
L121:
	;
	v448 = v442
	goto L111
L122:
	;
	v444 = v428 + int32(1)
	if v444 != v422 {
		v428 = v444
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v454 = F_remove_nulling_relids(m, v399, v452, int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L107
	} else {
		goto L127
	}
L125:
	;
	v456 = v399
	goto L126
L126:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v459 = int32(0)
	if v457 == v459 {
		v500 = v459
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v456 = v454
	goto L126
L128:
	;
	if v500 == int32(0) {
		v510 = v456
		v511 = v398
		goto L106
	} else {
		goto L142
	}
L129:
	;
	goto L128
L130:
	;
	if v458 == int32(0) {
		v500 = v459
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	if v468 < v469 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v471 = v468
	goto L134
L133:
	;
	v471 = v469
	goto L134
L134:
	;
	if v471 <= int32(1) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v474 = int32(1)
	goto L137
L136:
	;
	v474 = v471
	goto L137
L137:
	;
	v475 = int32(8)
	v480 = int32(0)
	goto L138
L138:
	;
	v487 = v480 << (uint(int32(2)) % 32)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v458+v475+v487)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v487+(v457+v475))))
	v492 = v489 & v491
	v494 = base.B2i32(v492 != int32(0))
	if v492 != 0 {
		v500 = v494
		goto L129
	} else {
		goto L140
	}
L139:
	;
	v500 = v494
	goto L129
L140:
	;
	v496 = v480 + int32(1)
	if v496 != v474 {
		v480 = v496
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v508 = F_remove_nulling_relids(m, v398, v506, int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L107
	} else {
		goto L143
	}
L143:
	;
	v510 = v456
	v511 = v508
	goto L106
L144:
	;
	if v512 < int32(0) {
		v599 = v80
		goto L19
	} else {
		goto L145
	}
L145:
	;
	v516 = F_match_expr_to_partition_keys(m, v511, l3, v401)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L107
	} else {
		goto L146
	}
L146:
	;
	if v516 != v512 {
		v599 = v80
		goto L19
	} else {
		goto L147
	}
L147:
	;
	v521 = v23 + int32(32) + v512
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521))))
	if v522 != 0 {
		v599 = v80
		goto L19
	} else {
		goto L148
	}
L148:
	;
	v524 = v512 << (uint(int32(2)) % 32)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l2)+232))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+12))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v524+v526)))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v155)+24))
	if v528 != v529 {
		goto L2
	} else {
		goto L149
	}
L149:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v531 == int32(104) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v588 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v521))) = uint8(v588)
	v591 = v80 + v588
	v592 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	if v591 == v592 {
		goto L9
	} else {
		goto L171
	}
L151:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v86)+124))
	if v534 == int32(0) {
		v599 = v80
		goto L19
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v543+v524)))
	v546 = int32(0)
	if v542 == v546 {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v537+v524)))
	v540 = F_op_in_opfamily(m, v534, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L107
	} else {
		goto L155
	}
L155:
	;
	if v540 != 0 {
		goto L150
	} else {
		goto L156
	}
L156:
	;
	v599 = v80
	goto L19
L157:
	;
	if v584 == int32(0) {
		v599 = v80
		goto L19
	} else {
		goto L170
	}
L158:
	;
	v584 = int32(0)
	goto L157
L159:
	;
	goto L160
L160:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v552 <= int32(0) {
		v577 = v546
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v584 = v577
	goto L157
L162:
	;
	v555 = int32(0)
	if v555 < v552 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v558 = v552
	goto L165
L164:
	;
	v558 = v555
	goto L165
L165:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
	v561 = int32(0)
	goto L166
L166:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v559+v561<<(uint(int32(2))%32))))
	v570 = base.B2i32(v569 == v545)
	if v569 == v545 {
		v577 = v570
		goto L161
	} else {
		goto L168
	}
L167:
	;
	v577 = v570
	goto L161
L168:
	;
	v572 = v561 + int32(1)
	if v572 != v558 {
		v561 = v572
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	goto L150
L171:
	;
	v599 = v591
	goto L19
L172:
	;
	goto L18
L173:
	;
	v633 = v624
	v641 = v7
	v645 = v622
	goto L174
L174:
	;
	v649 = v23 + int32(32) + v641
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649))))
	if v650 == int32(1) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L2
L176:
	;
	v815 = v641 + int32(1)
	if v815 < v799 {
		v633 = v800
		v641 = v815
		v645 = v812
		goto L174
	} else {
		goto L208
	}
L177:
	;
	v799 = base.I32_extend16_s(v633)
	v800 = v633
	v812 = v645
	goto L176
L178:
	;
	goto L179
L179:
	;
	v655 = v641 << (uint(int32(2)) % 32)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v657 = v655 + v656
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v658 == int32(104) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v657)))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v662+v655)))
	v666 = F_get_opfamily_member(m, v661, v664, v664, int32(1))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L107
	} else {
		goto L183
	}
L181:
	;
	v675 = v657
	goto L182
L182:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l2)+264))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v676+v655)))
	if v678 == int32(0) {
		goto L2
	} else {
		goto L187
	}
L183:
	;
	if v666 == int32(0) {
		goto L2
	} else {
		goto L184
	}
L184:
	;
	v670 = F_get_mergejoin_opfamilies(m, v666)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L107
	} else {
		goto L185
	}
L185:
	;
	if v670 == int32(0) {
		goto L2
	} else {
		goto L186
	}
L186:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v670)+12))
	v675 = v674
	goto L182
L187:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	if v681 <= int32(0) {
		goto L2
	} else {
		goto L188
	}
L188:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	v698 = int32(0)
	goto L189
L189:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l2)+232))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+12))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v707+v655)))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v678)+12))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v710+v698<<(uint(int32(2))%32))))
	v715 = F_exprCollation(m, v714)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L107
	} else {
		goto L191
	}
L190:
	;
	goto L2
L191:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l3)+264))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v717+v655)))
	if v719 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v791 = v698 + int32(1)
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	if v791 < v792 {
		v698 = v791
		goto L189
	} else {
		goto L207
	}
L193:
	;
	v722 = int32(0)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	if v723 <= v722 {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v731 = v722
	goto L195
L195:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v719)+12))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v746+v731<<(uint(int32(2))%32))))
	v751 = F_exprs_known_equal(m, l0, v714, v750, v684)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L107
	} else {
		goto L197
	}
L196:
	;
	v762 = F_exprCollation(m, v750)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L107
	} else {
		goto L205
	}
L197:
	;
	if v715 == v709 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v755 = v751
	goto L200
L199:
	;
	v755 = int32(0)
	goto L200
L200:
	;
	if v755 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v759 = v731 + int32(1)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	if v759 < v760 {
		v731 = v759
		goto L195
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	goto L196
L204:
	;
	goto L192
L205:
	;
	v764 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v649))) = uint8(v764)
	v766 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	v768 = v645 + v764
	if v768 == v766 {
		goto L9
	} else {
		goto L206
	}
L206:
	;
	v799 = v766
	v800 = v766
	v812 = v768
	goto L176
L207:
	;
	goto L190
L208:
	;
	goto L175
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+264)) = v843
	v846 = F_palloc0(m, v842)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L107
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+268)) = v846
	if int32(0) < v840 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	if base.Ui32(int32(5)) < base.Ui32(v839) {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v1085 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+217)) = uint8(v1085)
	goto L2
L214:
	;
	if int32(base.Ui32(int32(55))>>(uint(v839)%32))&int32(1) == int32(0) {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v872 = int32(0)
	goto L216
L216:
	;
	v881 = v872 << (uint(int32(2)) % 32)
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l3)+268))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v881+v882)))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l3)+264))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v885+v881)))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l2)+268))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v888+v881)))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l2)+264))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v891+v881)))
	switch v839 {
	case 0:
		goto L219
	case 1:
		goto L221
	default:
		goto L220
	case 4, 5:
		goto L222
	}
L217:
	;
	goto L213
L218:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v1056+v881))) = v1051
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l1)+268))
	*(*int32)(unsafe.Add(mBase, uint32(v1059+v881))) = v1047
	v1063 = v872 + int32(1)
	if v1063 != v840 {
		v872 = v1063
		goto L216
	} else {
		goto L253
	}
L219:
	;
	v1032 = F_list_concat_copy(m, v893, v887)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L107
	} else {
		goto L251
	}
L220:
	;
	v904 = F_list_concat_copy(m, v893, v887)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L107
	} else {
		goto L228
	}
L221:
	;
	v898 = F_list_copy(m, v893)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L107
	} else {
		goto L225
	}
L222:
	;
	v894 = F_list_copy(m, v893)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L107
	} else {
		goto L223
	}
L223:
	;
	v896 = F_list_copy(m, v890)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L107
	} else {
		goto L224
	}
L224:
	;
	v1047 = v896
	v1051 = v894
	goto L218
L225:
	;
	v900 = F_list_concat_copy(m, v887, v890)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L107
	} else {
		goto L226
	}
L226:
	;
	v902 = F_list_concat(m, v900, v884)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L107
	} else {
		goto L227
	}
L227:
	;
	v1047 = v902
	v1051 = v898
	goto L218
L228:
	;
	v906 = F_list_concat(m, v904, v890)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L107
	} else {
		goto L229
	}
L229:
	;
	v908 = F_list_concat(m, v906, v884)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L107
	} else {
		goto L230
	}
L230:
	;
	v910 = F_list_concat_copy(m, v893, v890)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L107
	} else {
		goto L231
	}
L231:
	;
	if v910 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1047 = v908
	v1051 = int32(0)
	goto L218
L233:
	;
	goto L234
L234:
	;
	v915 = int32(0)
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v910)+4))
	if v917 <= v915 {
		v1047 = v908
		v1051 = v915
		goto L218
	} else {
		goto L235
	}
L235:
	;
	v927 = v915
	v931 = v908
	goto L236
L236:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v910)+12))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v940+v927<<(uint(int32(2))%32))))
	v945 = F_list_concat_copy(m, v887, v884)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L107
	} else {
		goto L239
	}
L237:
	;
	v1047 = v1019
	v1051 = v915
	goto L218
L238:
	;
	v1029 = v927 + int32(1)
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v910)+4))
	if v1029 < v1030 {
		v927 = v1029
		v931 = v1019
		goto L236
	} else {
		goto L250
	}
L239:
	;
	if v945 == int32(0) {
		v1019 = v931
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v949 = int32(0)
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v945)+4))
	if v950 <= v949 {
		v1019 = v931
		goto L238
	} else {
		goto L241
	}
L241:
	;
	v953 = v949
	v964 = v931
	goto L242
L242:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v945)+12))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v973+v953<<(uint(int32(2))%32))))
	v979 = F_palloc0(m, int32(20))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L107
	} else {
		goto L244
	}
L243:
	;
	v1019 = v1002
	goto L238
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v979))) = int32(38)
	v983 = F_exprType(m, v944)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L107
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v979)+4)) = v983
	v986 = F_exprCollation(m, v944)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L107
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v979)+8)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v977
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v977
	v997 = F_list_make2_impl(m, v23+int32(12), v23+int32(8))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L107
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v979)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v979)+12)) = v997
	v1002 = F_lappend(m, v964, v979)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L107
	} else {
		goto L248
	}
L248:
	;
	v1005 = v953 + int32(1)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v945)+4))
	if v1005 < v1006 {
		v953 = v1005
		v964 = v1002
		goto L242
	} else {
		goto L249
	}
L249:
	;
	goto L243
L250:
	;
	goto L237
L251:
	;
	v1034 = F_list_concat_copy(m, v890, v884)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L107
	} else {
		goto L252
	}
L252:
	;
	v1047 = v1034
	v1051 = v1032
	goto L218
L253:
	;
	goto L217
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v839
	F_errmsg_internal(m, int32(_a_F_build_joinrel_partition_info_0), v23+int32(16))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L107
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_build_joinrel_partition_info_1), int32(2491), int32(_a_F_build_joinrel_partition_info_2))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L107
	} else {
		goto L256
	}
L256:
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v73 int32
	_ = v73
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
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
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
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v115 = v7
	v118 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v34 = v7
	v35 = v7
	v36 = v24
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v45 = v35 + base.B2i32(v41&l2 != int32(0))
	v47 = v34 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23+v47<<(uint(int32(2))%32))))
	if v51 != 0 {
		v34 = v47
		v35 = v45
		v36 = v51
		goto L10
	} else {
		goto L12
	}
L11:
	;
	if v45 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v115 = v45
	v118 = v105
	goto L6
L14:
	;
	v105 = int32(0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v57 = F_palloc(m, v45<<(uint(int32(4))%32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_build_reloptions[1]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v61 == int32(0) {
		v105 = v57
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v71 = v61
	v73 = int32(0)
	v75 = v7
	goto L19
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v78&l2 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v105 = v57
	goto L13
L21:
	;
	v82 = v57 + v75<<(uint(int32(4))%32)
	v83 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v71
	v88 = v75 + int32(1)
	goto L23
L22:
	;
	v88 = v75
	goto L23
L23:
	;
	v91 = v73 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v60+v91<<(uint(int32(2))%32))))
	if v95 != 0 {
		v71 = v95
		v73 = v91
		v75 = v88
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	F_parseRelOptionsInternal(m, l0, l1, v118, v115)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v115 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	return int32(0)
L30:
	;
	goto L31
L31:
	;
	if int32(0) < v115 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v138 = int32(0)
	v141 = l3
	goto L35
L33:
	;
	v250 = l3
	goto L34
L34:
	;
	v253 = F_palloc0(m, v250)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L74
	}
L35:
	;
	v144 = int32(4)
	v146 = v118 + v138<<(uint(v144)%32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	if v148 == v144 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v250 = v234
	goto L34
L37:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)+36))
	if v152 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v234 = v141
	goto L39
L39:
	;
	v238 = v138 + int32(1)
	if v238 != v115 {
		v138 = v238
		v141 = v234
		goto L35
	} else {
		goto L73
	}
L40:
	;
	v234 = v231 + v141
	goto L39
L41:
	;
	if v151&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if v151&int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v157 = m.T0[v152].(func(*base.Module, int32, int32) int32)(m, v155, int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+28)))
	if v160 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v231 = v157
	goto L40
L48:
	;
	v162 = int32(0)
	goto L50
L49:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	v162 = v161
	goto L50
L50:
	;
	v164 = m.T0[v152].(func(*base.Module, int32, int32) int32)(m, v162, int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v231 = v164
	goto L40
L52:
	;
	v231 = v227 + int32(1)
	goto L40
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	if v168&int32(3) == int32(0) {
		v192 = v168
		goto L58
	} else {
		goto L59
	}
L54:
	;
	goto L55
L55:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v147)+24))
	v227 = v226
	goto L52
L56:
	;
	v227 = v225
	goto L52
L57:
	;
	v225 = v217 - v168
	goto L56
L58:
	;
	v196 = v192
	goto L67
L59:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v176 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v225 = int32(0)
	goto L56
L61:
	;
	goto L62
L62:
	;
	v181 = v168
	goto L63
L63:
	;
	v185 = v181 + int32(1)
	if v185&int32(3) == int32(0) {
		v192 = v185
		goto L58
	} else {
		goto L65
	}
L64:
	;
	v217 = v185
	goto L57
L65:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v190 != 0 {
		v181 = v185
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v205 = int32(-2139062144)
	if (int32(16843008)-v202|v202)&v205 == v205 {
		v196 = v196 + int32(4)
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v211 = v196
	goto L70
L69:
	;
	goto L68
L70:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v215 != 0 {
		v211 = v211 + int32(1)
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v217 = v211
	goto L57
L72:
	;
	goto L71
L73:
	;
	goto L36
L74:
	;
	F_fillRelOptions(m, v253, l3, v118, v115, l1, l4, l5)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_pfree(m, v118)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	return v253
}
func F_buildint2vector(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	v7 = l1 << (uint(int32(1)) % 32)
	v9 = v7 + int32(24)
	v10 = F_palloc0(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
		} else {
			if l1 <= int32(0) {
			} else {
				if v7 != 0 {
					v20 = F__emscripten_memcpy_bulkmem(m, v10+int32(24), l0, v7)
					mBase = m.M
				} else {
				}
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(21)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(1)
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v9 << (uint(int32(2)) % 32)
		return v10
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
				v24 = int32(4)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
				if v26&int32(254) == int32(2) {
					v35 = v24
				} else {
					v35 = base.B2i32(v26 == int32(18)) << (uint(v24) % 32)
				}
				if v26 == int32(1) {
					v38 = v24
				} else {
					v38 = v35
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v20 != 0 {
					v49 = int32(base.Ui32(v18)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = int32(1)
			v51 = v14 + v50
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v56 = v54 & v50
			if v56 != 0 {
				v57 = v51
			} else {
				v57 = v14 + int32(4)
			}
			if v54 == int32(1) {
				v60 = int32(4)
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				if v62&int32(254) == int32(2) {
					v71 = v60
				} else {
					v71 = base.B2i32(v62 == int32(18)) << (uint(v60) % 32)
				}
				if v62 == int32(1) {
					v74 = v60
				} else {
					v74 = v71
				}
				v85 = v74
			} else {
				v75 = int32(1)
				if v56 != 0 {
					v85 = int32(base.Ui32(v54)>>(uint(v75)%32)) - v75
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = F_SB_MatchText(m, v21, v49, v57, v85, int32(0))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v87 != int32(1))
			}
		}
	}
}
