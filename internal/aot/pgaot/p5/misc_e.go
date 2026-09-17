package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EOH_init_header(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = int32(513)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v4)
	v6 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v6)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l0)+14)) = l0
	return
}
func F_ExecASDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v5 == int32(0) {
		return
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+22)))
		if v8 != int32(1) {
			return
		} else {
			v11 = int32(0)
			F_AfterTriggerSaveEvent(m, l0, l1, v11, v11, int32(1), v11, v11, v11, v11, v11, l2, v11)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ExecBSTruncateTriggers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v19 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L20
	}
L2:
	;
	m.G0 = v9 + int32(48)
	return
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+23)))
	if v22 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(47244640698)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v29 <= int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v37 = v3
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v41 = v38 + v37*int32(60)
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
	if v42&int32(99) != int32(34) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L2
L8:
	;
	v69 = v37 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v69 < v70 {
		v37 = v69
		goto L6
	} else {
		goto L19
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v48 = int32(0)
	v51 = F_TriggerEnabled(m, l0, l1, v41, v47, v48, v48, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v51 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v41
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v60 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v63 = v60
	goto L15
L14:
	;
	v61 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v65 = F_ExecCallTriggerFunc(m, v9+int32(4), v37, v58, v59, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v63 = v61
	goto L15
L17:
	;
	if v65 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L8
L19:
	;
	goto L7
L20:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_ExecBSTruncateTriggers_0), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ExecBSTruncateTriggers_1), int32(3323), int32(_a_F_ExecBSTruncateTriggers_2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecFindJunkAttributeInTlist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	if l0 == int32(0) {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = int32(0)
	if v14 < v11 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v78 = int32(0)
	goto L6
L6:
	;
	return base.I32_extend16_s(v78)
L7:
	;
	v17 = v11
	goto L9
L8:
	;
	v17 = v14
	goto L9
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = int32(0)
	goto L11
L10:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+8)))
	v78 = v70
	goto L6
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(int32(2))%32))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+26)))
	if v30 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return int32(0)
L13:
	;
	v66 = v20 + int32(1)
	if v66 != v17 {
		v20 = v66
		goto L11
	} else {
		goto L24
	}
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v33 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v38 == int32(0))|base.B2i32(v38 != v41) != 0 {
		v59 = v38
		v60 = v41
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v59-v60 == int32(0) {
		goto L10
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v44 = v33
	v45 = l1
	goto L19
L19:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v49 == int32(0) {
		v59 = v49
		v60 = v48
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v59 = v49
	v60 = v48
	goto L17
L21:
	;
	v52 = int32(1)
	if v49 == v48 {
		v44 = v44 + v52
		v45 = v45 + v52
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L13
L24:
	;
	goto L12
}
func F_ExecFindMatchingSubPlans(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v4
	v15 = int32(_a_F_ExecFindMatchingSubPlans_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindMatchingSubPlans[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindMatchingSubPlans[0])) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 < v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = v4
	goto L4
L2:
	;
	v66 = int32(0)
	goto L3
L3:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v68 = F_bms_add_members(m, v66, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L13
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(24)+v30<<(uint(int32(2))%32))))
	F_find_matching_subplans_recurse(m, v36, v36+int32(4), l1, v11+int32(12), l2)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v66 = v56
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v53 = v30 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v53 < v54 {
		v30 = v53
		goto L4
	} else {
		goto L12
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	if v45 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+116))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	F_MemoryContextReset(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L5
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindMatchingSubPlans[0])) = v16
	v72 = F_bms_copy(m, v68)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v75 = F_bms_copy(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_MemoryContextReset(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v75
	goto L17
L19:
	;
	m.G0 = v11 + int32(16)
	return v72
}
func F_ExecGetReturningSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v5 == int32(0) {
		v8 = int32(_a_F_ExecGetReturningSlot_0)
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGetReturningSlot[0]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecGetReturningSlot[0])) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
		v15 = F_table_slot_callbacks(m, v10)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_ExecInitExtraTupleSlot(m, l0, v14, v15)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v19
				*(*int32)(unsafe.Add(mBase, _c_F_ExecGetReturningSlot[0])) = v9
				v24 = v19
				return v24
			}
		}
	} else {
		v24 = v5
		return v24
	}
}
func F_ExecInitExtraTupleSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_MakeTupleTableSlot(m, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v9 = F_lappend(m, v8, v4)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v9
			return v4
		}
	}
}
func F_ExecLockRows(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
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
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
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
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
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
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[0]))
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v25 = l0 + int32(108)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L6
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	if v41 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v16 + int32(80)
	return v371
L8:
	;
	F_ExecReScan(m, v26)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v45 = m.T0[v44].(func(*base.Module, int32) int32)(m, v26)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	goto L7
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v55 == int32(0) {
		v371 = v45
		goto L12
	} else {
		goto L20
	}
L14:
	;
	if v45 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+4)))
	if v47&int32(2) == int32(0) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_EvalPlanQualEnd(m, v25)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v371 = int32(0)
	goto L12
L20:
	;
	v58 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v60 <= v58 {
		v371 = v45
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v71 = v58
	v75 = v58
	goto L22
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v75<<(uint(int32(2))%32))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v84 = F_EvalPlanQualSlot(m, v25, v82, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	if v338&int32(1) == int32(0) {
		v371 = v45
		goto L12
	} else {
		goto L96
	}
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	m.T0[v87].(func(*base.Module, int32))(m, v84)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v90 == v91 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v340 = v75 + int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v340 < v341 {
		v71 = v338
		v75 = v340
		goto L22
	} else {
		goto L95
	}
L27:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L92
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L89
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L85
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L82
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L79
	}
L32:
	;
	v119 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+32)) = uint8(v119)
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80)+4)))
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+6)))
	if v122 < v121 {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80)+6)))
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+6)))
	if v94 < v93 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_slot_getsomeattrs_int(m, v45, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v98 = int32(1)
	v99 = v93 - v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v100))))
	if v102 == v98 {
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v99<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v109 == v110 {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	v112 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+38)) = uint16(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+34)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+32)) = uint8(v112)
	v338 = v71
	goto L26
L40:
	;
	F_slot_getsomeattrs_int(m, v45, v121)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v126 = int32(1)
	v127 = v121 - v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+v128))))
	if v130 == v126 {
		goto L30
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v127<<(uint(int32(2))%32))))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+48))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+119)))
	if v140 == int32(102) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v143 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+51)) = uint8(v143)
	v146 = F_GetFdwRoutineForRelation(m, v138, v143)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+76)) = uint16(v160)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if base.Ui32(int32(4)) <= base.Ui32(v164) {
		goto L28
	} else {
		goto L52
	}
L48:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v146)+108))
	if v148 == int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	m.T0[v148].(func(*base.Module, int32, int32, int32, int32, int32))(m, v27, v81, v137, v84, v16+int32(51))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
	if v155&int32(2) != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+51)))
	v338 = v158 | v71
	goto L26
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	v172 = int32(3)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v81)+28))
	v175 = int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[1]))
	if v175 < v178 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L76
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L73
	}
L55:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[1]))
	if v217 < int32(2) {
		goto L6
	} else {
		goto L68
	}
L56:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[1]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L64
	}
L57:
	;
	v181 = v175
	goto L59
L58:
	;
	v181 = v172
	goto L59
L59:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v167)+188))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+104))
	v186 = m.T0[v185].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v167, v16+int32(72), v170, v84, v171, v172-v164, v174, v181, v16+int32(52))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	if v186 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	switch v186 - int32(1) {
	case 0:
		goto L54
	case 1, 5:
		goto L6
	case 2:
		goto L56
	case 3:
		goto L55
	default:
		goto L53
	}
L62:
	;
	goto L63
L63:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+68)))
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+38)) = uint16(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+34)) = v193
	v338 = v190 | v71
	goto L26
L64:
	;
	if int32(2) <= v197 {
		goto L27
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(3)
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_0), v16+int32(32))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(230), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(_a_F_ExecLockRows_3), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(237), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_4), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(242), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v186
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_5), v16+int32(16))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(247), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_6), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(102), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_7), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(122), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v298 + int32(4)
	F_errmsg(m, int32(_a_F_ExecLockRows_8), v16)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(136), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_9), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(176), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errmsg(m, int32(_a_F_ExecLockRows_3), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(228), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	goto L23
L96:
	;
	F_EvalPlanQualBegin(m, v25)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v45
	v350 = int32(_a_F_ExecLockRows_10)
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[2]))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[2])) = v354
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+52))
	if v357 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_ExecReScan(m, v356)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v361 = m.T0[v360].(func(*base.Module, int32) int32)(m, v356)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[2])) = v351
	if v361 == int32(0) {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+4)))
	if v367&int32(2) != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v371 = v361
	goto L12
}
func F_ExecModifyTable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
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
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
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
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
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
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
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
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
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
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
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
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v517 int32
	_ = v517
	var v546 int32
	_ = v546
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v811 int32
	_ = v811
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
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
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v992 float64
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1005 int32
	_ = v1005
	var v1006 float64
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 float64
	_ = v1092
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1114 int64
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1233 int32
	_ = v1233
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
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
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 float64
	_ = v1402
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1491 int32
	_ = v1491
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1527 int32
	_ = v1527
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1575 int32
	_ = v1575
	var v1580 int32
	_ = v1580
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1631 int32
	_ = v1631
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1796 int32
	_ = v1796
	var v1800 int32
	_ = v1800
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1882 int32
	_ = v1882
	var v1886 int32
	_ = v1886
	var v1890 int32
	_ = v1890
	var v1895 int32
	_ = v1895
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1915 int32
	_ = v1915
	var v1918 int32
	_ = v1918
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1934 int32
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1950 int32
	_ = v1950
	var v1954 int32
	_ = v1954
	var v1958 int32
	_ = v1958
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1974 int32
	_ = v1974
	var v1978 int32
	_ = v1978
	var v1983 int32
	_ = v1983
	var v1986 int32
	_ = v1986
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v2002 int32
	_ = v2002
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	v27 = m.G0
	v29 = v27 - int32(176)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[0]))
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+156))
	if v39 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L4
	} else {
		goto L587
	}
L7:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L4
	} else {
		goto L583
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L4
	} else {
		goto L578
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L4
	} else {
		goto L575
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L4
	} else {
		goto L571
	}
L11:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L4
	} else {
		goto L567
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L4
	} else {
		goto L562
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L4
	} else {
		goto L559
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L4
	} else {
		goto L556
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L4
	} else {
		goto L553
	}
L16:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)))
	if v42 != 0 {
		v1815 = int32(0)
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L4
	} else {
		goto L550
	}
L19:
	;
	m.G0 = v29 + int32(176)
	return v1815
L20:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)))
	if v43 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	switch v47 - int32(2) {
	case 0:
		goto L25
	case 1:
		goto L29
	case 2:
		goto L28
	case 3:
		goto L27
	default:
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+116)) = v32
	v110 = int32(124)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = l0 + v110
	*(*int32)(unsafe.Add(mBase, uint32(v29)+108)) = l0
	v118 = v29 + v110
	v127 = v107 + v108*int32(216)
	goto L48
L24:
	;
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)) = uint8(v102)
	goto L23
L25:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSUpdateTriggers(m, v98, v46)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L47
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L44
	}
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	if v63&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSDeleteTriggers(m, v60, v46)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L33
	}
L29:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSInsertTriggers(m, v51, v46)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+132))
	if v54 != int32(2) {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSUpdateTriggers(m, v57, v46)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L24
L33:
	;
	goto L24
L34:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSInsertTriggers(m, v66, v46)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v70 = v63
	goto L36
L36:
	;
	if v70&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v70 = v69
	goto L36
L38:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSUpdateTriggers(m, v73, v46)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	v77 = v70
	goto L40
L40:
	;
	if v77&int32(4) == int32(0) {
		goto L24
	} else {
		goto L42
	}
L41:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v77 = v76
	goto L40
L42:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSDeleteTriggers(m, v82, v46)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L24
L44:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_0), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(4008), int32(_a_F_ExecModifyTable_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	goto L24
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v32)+152))
	if v149 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v32)+188))
	if v1745 != 0 {
		goto L520
	} else {
		goto L521
	}
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	F_MemoryContextReset(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v153 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+20))
	F_MemoryContextReset(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v157 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+144)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v157
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v165 = F_ExecMergeNotMatched(m, v29+int32(108), v163, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v106)+52))
	if v171 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v167 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v167
	if v165 == v167 {
		goto L48
	} else {
		goto L62
	}
L62:
	;
	v1815 = v165
	goto L19
L63:
	;
	F_ExecReScan(m, v106)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v175 = m.T0[v174].(func(*base.Module, int32) int32)(m, v106)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v177 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+144)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v175
	if v175 == v177 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	goto L49
L69:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
	if v182&int32(2) != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v185 == int32(0) {
		v238 = v127
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+92)))
	if v240 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L72:
	;
	v188 = base.I32_extend16_s(v185)
	v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(v175)+6)))
	if v189 < v188 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_slot_getsomeattrs_int(m, v175, v188)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v193 = int32(1)
	v194 = v188 - v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v195))))
	if v197 == v193 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	if v31 == int32(5) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v194<<(uint(int32(2))%32))))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v229 == v230 {
		v238 = v127
		goto L71
	} else {
		goto L88
	}
L80:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v202
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v208 = F_ExecMergeNotMatched(m, v29+int32(108), v206, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	if v208 == int32(0) {
		goto L48
	} else {
		goto L84
	}
L84:
	;
	v1815 = v208
	goto L19
L85:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_3), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_4), int32(_a_F_ExecModifyTable_5))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v234 = F_ExecLookupResultRelByOid(m, l0, v229, int32(0), int32(1))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v238 = v234
	goto L71
L90:
	;
	v245 = int32(0)
	v247 = F_ExecProcessReturning(m, v29+int32(108), v238, v31, v245, v245, v239)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v239
	v250 = int32(0)
	if base.Ui32(int32(5)) < base.Ui32(v31) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v1815 = v247
	goto L19
L94:
	;
	switch v31 - int32(2) {
	case 0:
		goto L143
	case 1:
		goto L144
	case 2:
		goto L140
	case 3:
		goto L142
	default:
		goto L141
	}
L95:
	;
	v389 = int32(0)
	v391 = v250
	goto L94
L96:
	;
	goto L97
L97:
	;
	if int32(1)<<(uint(v31)%32)&int32(52) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v389 = int32(0)
	v391 = v250
	goto L94
L99:
	;
	goto L100
L100:
	;
	v259 = int32(*(*int16)(unsafe.Add(mBase, uint32(v238)+24)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+48))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+119)))
	v264 = v262 - int32(109)
	v271 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v264))|base.B2i32(int32(1)<<(uint(v264)%32)&int32(41) == v271) == v271 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v276 = int32(*(*int16)(unsafe.Add(mBase, uint32(v239)+6)))
	if v276 < v259 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v323 = int32(0)
	if v259 == v323 {
		v389 = v323
		v391 = v250
		goto L94
	} else {
		goto L119
	}
L104:
	;
	F_slot_getsomeattrs_int(m, v239, v259)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v280 = int32(1)
	v281 = v259 - v280
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281+v282))))
	if v284 == v280 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L106
L108:
	;
	if v31 == int32(5) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	goto L110
L110:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v312+v281<<(uint(int32(2))%32))))
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+104)) = uint16(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v319
	v389 = v29 + int32(100)
	v391 = v250
	goto L94
L111:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v289
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v295 = F_ExecMergeNotMatched(m, v29+int32(108), v293, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L116
	}
L114:
	;
	if v295 == int32(0) {
		v127 = v238
		goto L48
	} else {
		goto L115
	}
L115:
	;
	v1815 = v295
	goto L19
L116:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_6), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_7), int32(_a_F_ExecModifyTable_5))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v326 = int32(*(*int16)(unsafe.Add(mBase, uint32(v239)+6)))
	if v326 < v259 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_slot_getsomeattrs_int(m, v239, v259)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v330 = int32(1)
	v331 = v259 - v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v332))))
	if v334 == v330 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L122
L124:
	;
	if v31 == int32(5) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362+v331<<(uint(int32(2))%32))))
	v367 = F_pg_detoast_datum(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L135
	}
L127:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v339
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v345 = F_ExecMergeNotMatched(m, v29+int32(108), v343, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L4
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L132
	}
L130:
	;
	if v345 == int32(0) {
		v127 = v238
		goto L48
	} else {
		goto L131
	}
L131:
	;
	v1815 = v345
	goto L19
L132:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_8), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_9), int32(_a_F_ExecModifyTable_5))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v367
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v371 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+88)) = uint16(v371)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = int32(base.Ui32(v370) >> (uint(int32(2)) % 32))
	if v262 != int32(118) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+56))
	v383 = v381
	goto L138
L137:
	;
	v383 = int32(0)
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+92)) = v383
	v389 = v323
	v391 = v29 + int32(80)
	goto L94
L139:
	;
	if v1718 == int32(0) {
		v127 = v238
		goto L48
	} else {
		goto L519
	}
L140:
	;
	v1710 = int32(0)
	v1711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v1715 = F_ExecDelete(m, v29+int32(108), v238, v389, v391, int32(1), v1710, v1711, v1710, v1710, v1710)
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L4
	} else {
		goto L518
	}
L141:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L4
	} else {
		goto L515
	}
L142:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v389|v391 != 0 {
		goto L211
	} else {
		goto L212
	}
L143:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+48)))
	if v627 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L144:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+48)))
	if v392 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v395 = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+52))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+44))
	if v401 != 0 {
		goto L151
	} else {
		goto L152
	}
L146:
	;
	goto L147
L147:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	if v575 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L148:
	;
	v546 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+48)) = uint8(v546)
	goto L147
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v492+v238))) = v517
	goto L148
L150:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	F_ExecCheckPlanOutput(m, v481, int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L173
	}
L151:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	if v402 <= int32(0) {
		goto L150
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	F_ExecCheckPlanOutput(m, v471, int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L171
	}
L154:
	;
	v406 = v395
	v411 = v395
	v421 = v395
	goto L155
L155:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v431+v406<<(uint(int32(2))%32))))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+26)))
	if v436 != 0 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	F_ExecCheckPlanOutput(m, v446, v441)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L163
	}
L157:
	;
	v443 = v406 + int32(1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	if v443 < v444 {
		v406 = v443
		v411 = v440
		v421 = v441
		goto L155
	} else {
		goto L162
	}
L158:
	;
	v440 = int32(1)
	v441 = v421
	goto L157
L159:
	;
	goto L160
L160:
	;
	v438 = F_lappend(m, v421, v435)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v440 = v411
	v441 = v438
	goto L157
L162:
	;
	goto L156
L163:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v452 = F_table_slot_create(m, v449, v398+int32(104))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+40)) = v452
	if v440 == int32(0) {
		goto L148
	} else {
		goto L165
	}
L165:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)+52))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v459 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_ExecAssignExprContext(m, v398, l0)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L169
	}
L167:
	;
	v466 = v452
	v467 = v459
	goto L168
L168:
	;
	v469 = F_ExecBuildProjectionInfo(m, v441, v467, v466, l0, v458)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L170
	}
L169:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v466 = v465
	v467 = v464
	goto L168
L170:
	;
	v492 = int32(36)
	v517 = v469
	goto L149
L171:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v479 = F_table_slot_create(m, v476, v398+int32(104))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	v492 = int32(40)
	v517 = v479
	goto L149
L173:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v489 = F_table_slot_create(m, v486, v398+int32(104))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	v492 = int32(40)
	v517 = v489
	goto L149
L175:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v621 = int32(0)
	v623 = F_ExecInsert(m, v29+int32(108), v238, v614, v620, v621, v621)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L4
	} else {
		goto L183
	}
L176:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+8))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v574)+8))
	if v579 == v580 {
		v614 = v574
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v575)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v586)+12)) = v574
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v575)+72))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v575)+16))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)+8))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+12))
	m.T0[v591].(func(*base.Module, int32))(m, v589)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L181
	}
L179:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v579)+32))
	m.T0[v582].(func(*base.Module, int32, int32))(m, v578, v574)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	v614 = v585
	goto L175
L181:
	;
	v594 = int32(_a_F_ExecModifyTable_10)
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v588)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v597
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v575)+24))
	v603 = m.T0[v602].(func(*base.Module, int32, int32, int32) int32)(m, v575+int32(4), v588, int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v595
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+4)))
	v609 = v607 & int32(_a_F_ExecModifyTable_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v589)+4)) = uint16(v609)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	*(*uint16)(unsafe.Add(mBase, uint32(v589)+6)) = uint16(v612)
	v614 = v589
	goto L175
L183:
	;
	if v623 == int32(0) {
		v127 = v238
		goto L48
	} else {
		goto L184
	}
L184:
	;
	v1815 = v623
	goto L19
L185:
	;
	F_ExecInitUpdateProjection(m, l0, v238)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L4
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	if v391 != 0 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L187
L189:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v662)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v663)+4)) = v632
	*(*int32)(unsafe.Add(mBase, uint32(v663)+12)) = v661
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v662)+72))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v662)+16))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)+8))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)+12))
	m.T0[v669].(func(*base.Module, int32))(m, v667)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L4
	} else {
		goto L204
	}
L190:
	;
	v633 = int32(0)
	F_ExecForceStoreHeapTuple(m, v391, v632, v633)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+49)))
	if v638 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v660 = v633
	goto L189
L194:
	;
	F_LockTuple(m, v637, v389, int32(7))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L4
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[2]))
	if v645 != 0 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L196
L198:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecModifyTable[3])))
	if v647&int32(1) == int32(0) {
		goto L6
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v637)+188))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v653)+60))
	v655 = m.T0[v654].(func(*base.Module, int32, int32, int32, int32) int32)(m, v637, v389, int32(_a_F_ExecModifyTable_12), v632)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L4
	} else {
		goto L202
	}
L201:
	;
	goto L200
L202:
	;
	if v655 == int32(0) {
		goto L15
	} else {
		goto L203
	}
L203:
	;
	v660 = v638
	goto L189
L204:
	;
	v672 = int32(_a_F_ExecModifyTable_10)
	v673 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v666)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v675
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v662)+24))
	v681 = m.T0[v680].(func(*base.Module, int32, int32, int32) int32)(m, v662+int32(4), v666, int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v673
	v685 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v667)+4)))
	v687 = v685 & int32(_a_F_ExecModifyTable_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v667)+4)) = uint16(v687)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v667)+12))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)))
	*(*uint16)(unsafe.Add(mBase, uint32(v667)+6)) = uint16(v690)
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v695 = F_ExecUpdate(m, v29+int32(108), v238, v389, v391, v632, v667, v694)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	if v660 == int32(0) {
		v1718 = v695
		goto L139
	} else {
		goto L207
	}
L207:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	F_UnlockTuple(m, v699, v389, int32(7))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L208
	}
L208:
	;
	if v695 == int32(0) {
		v127 = v238
		goto L48
	} else {
		goto L209
	}
L209:
	;
	v1815 = v695
	goto L19
L210:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1691)+220)) = v1692
	v1815 = v1625
	goto L19
L211:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+64))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v238)+164))
	if v710 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	v1687 = F_ExecMergeNotMatched(m, v29+int32(108), v238, v705&int32(1))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L4
	} else {
		goto L513
	}
L214:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v238)+168))
	if v713 == int32(0) {
		v127 = v238
		goto L48
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+4)) = v716
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	v719 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+12)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v708)+8)) = v718
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+172)) = uint16(v719)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+168)) = int32(-1)
	if v391 != 0 {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	goto L216
L218:
	;
	v760 = v238 + int32(164)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v238)+176))
	if v761 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L219:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	F_ExecForceStoreHeapTuple(m, v391, v726, int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L4
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+49)))
	if v730 == int32(1) {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	goto L218
L223:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	F_LockTuple(m, v733, v389, int32(7))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L4
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v742 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[2]))
	if v742 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v737 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+172)) = uint16(v737)
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+168)) = v739
	goto L225
L227:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecModifyTable[3])))
	if v744&int32(1) == int32(0) {
		goto L6
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v749)+188))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v752)+60))
	v754 = m.T0[v753].(func(*base.Module, int32, int32, int32, int32) int32)(m, v749, v389, int32(_a_F_ExecModifyTable_12), v751)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L4
	} else {
		goto L231
	}
L230:
	;
	goto L229
L231:
	;
	if v754 == int32(0) {
		goto L14
	} else {
		goto L232
	}
L232:
	;
	goto L218
L233:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	if v785 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L234:
	;
	v781 = v760
	v784 = v238 + int32(168)
	goto L233
L235:
	;
	goto L236
L236:
	;
	v766 = int32(_a_F_ExecModifyTable_10)
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v708)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v769
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v761)+20))
	v774 = m.T0[v773].(func(*base.Module, int32, int32, int32) int32)(m, v761, v708, v29+int32(152))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v767
	v779 = v238 + int32(168)
	if v774 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v780 = v760
	goto L240
L239:
	;
	v780 = v779
	goto L240
L240:
	;
	v781 = v780
	v784 = v779
	goto L233
L241:
	;
	v1650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+172)))
	if v1650 != 0 {
		goto L507
	} else {
		goto L508
	}
L242:
	;
	v1625 = int32(0)
	v1631 = int32(1)
	goto L241
L243:
	;
	goto L244
L244:
	;
	v791 = v707 + int32(124)
	v800 = int32(0)
	v801 = int32(1)
	v811 = v785
	goto L246
L245:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	v1622 = F_ExecProcessReturning(m, v29+int32(108), v238, int32(2), v1620, v1585, v1621)
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L4
	} else {
		goto L506
	}
L246:
	;
	v820 = int32(0)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	if v821 <= v820 {
		v1625 = v820
		v1631 = v801
		goto L241
	} else {
		goto L248
	}
L247:
	;
	v1625 = int32(0)
	v1631 = v801
	goto L241
L248:
	;
	v825 = v820
	goto L249
L249:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v811)+12))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v850+v825<<(uint(int32(2))%32))))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+4))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v855)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+160)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+152)) = int64(0)
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v854)+12))
	if v861 != 0 {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	goto L247
L251:
	;
	v1613 = v825 + int32(1)
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	if v1613 < v1614 {
		v825 = v1613
		goto L249
	} else {
		goto L505
	}
L252:
	;
	v862 = int32(_a_F_ExecModifyTable_10)
	v863 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v708)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v865
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v861)+20))
	v870 = m.T0[v869].(func(*base.Module, int32, int32, int32) int32)(m, v861, v708, v29+int32(175))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L4
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v238)+116))
	v878 = int32(0)
	if base.B2i32(v877 == v878)|base.B2i32(v856 == int32(7)) == v878 {
		goto L257
	} else {
		goto L258
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v863
	if v870 == int32(0) {
		goto L251
	} else {
		goto L256
	}
L256:
	;
	goto L254
L257:
	;
	if v856 == int32(2) {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L259
L259:
	;
	v896 = v856 - int32(2)
	switch v896 {
	case 0:
		goto L269
	default:
		goto L13
	case 2:
		goto L268
	case 5:
		goto L265
	}
L260:
	;
	v889 = int32(4)
	goto L262
L261:
	;
	v889 = int32(5)
	goto L262
L262:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v891)+8))
	F_ExecWithCheckOptions(m, v889, v238, v890, v892)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L4
	} else {
		goto L263
	}
L263:
	;
	goto L259
L264:
	;
	v1587 = int32(0)
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v238)+152))
	if v1588 == v1587 {
		v1625 = v1587
		v1631 = v801
		goto L241
	} else {
		goto L498
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = int32(0)
	v1585 = v800
	goto L264
L266:
	;
	if v1101 != int32(3) {
		goto L336
	} else {
		goto L337
	}
L267:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v29)+164))
	v1101 = v1100
	v1103 = v1098
	goto L266
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v707)+216)) = v854
	v1011 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = v1011
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	if v1013 == v1011 {
		goto L303
	} else {
		goto L304
	}
L269:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v854)+8))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)+72))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v897)+16))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v899)+8))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)+12))
	m.T0[v901].(func(*base.Module, int32))(m, v899)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	v904 = int32(_a_F_ExecModifyTable_10)
	v905 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v898)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v907
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v897)+24))
	v913 = m.T0[v912].(func(*base.Module, int32, int32, int32) int32)(m, v897+int32(4), v898, int32(0))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v905
	v917 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v899)+4)))
	v919 = v917 & int32(_a_F_ExecModifyTable_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v899)+4)) = uint16(v919)
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v899)+12))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v921)))
	*(*uint16)(unsafe.Add(mBase, uint32(v899)+6)) = uint16(v922)
	*(*int32)(unsafe.Add(mBase, uint32(v707)+216)) = v854
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = int32(0)
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v899)+8))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v928)+28))
	m.T0[v929].(func(*base.Module, int32))(m, v899)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v925)+48))
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932)+116)))
	if v933 != int32(1) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	if v940 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L274:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	if v936 != 0 {
		goto L273
	} else {
		goto L275
	}
L275:
	;
	F_ExecOpenIndices(m, v238, int32(0))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L4
	} else {
		goto L276
	}
L276:
	;
	goto L273
L277:
	;
	if v998 != 0 {
		v1098 = v899
		goto L267
	} else {
		goto L299
	}
L278:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v29)+164))
	v998 = v997
	goto L277
L279:
	;
	v984 = F_ExecUpdateAct(m, v29+int32(108), v238, v389, int32(0), v899, v705&int32(1), v29+int32(152))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L4
	} else {
		goto L297
	}
L280:
	;
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v940)+13)))
	if v943 == int32(1) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v946)+188))
	if v947 != 0 {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	v969 = v940
	goto L283
L283:
	;
	v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969)+15)))
	if v970 != int32(1) {
		goto L279
	} else {
		goto L294
	}
L284:
	;
	F_ExecPendingInserts(m, v946)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L4
	} else {
		goto L287
	}
L285:
	;
	v951 = v946
	goto L286
L286:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v956)+104))
	v960 = F_ExecBRUpdateTriggers(m, v951, v952, v238, v389, int32(0), v899, v29+int32(164), v118, base.B2i32(v957 == int32(5)))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L4
	} else {
		goto L288
	}
L287:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	v951 = v950
	goto L286
L288:
	;
	if v960 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v29)+164))
	if v964 != 0 {
		v1101 = v964
		v1103 = v899
		goto L266
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	if v966 == int32(0) {
		goto L279
	} else {
		goto L293
	}
L292:
	;
	v1625 = int32(0)
	v1631 = v801
	goto L241
L293:
	;
	v969 = v966
	goto L283
L294:
	;
	v973 = F_ExecIRUpdateTriggers(m, v709, v238, v391, v899)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L4
	} else {
		goto L295
	}
L295:
	;
	if v973 != 0 {
		goto L278
	} else {
		goto L296
	}
L296:
	;
	v1625 = int32(0)
	v1631 = v801
	goto L241
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = v984
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+152)))
	if v987&int32(1) == int32(0) {
		v998 = v984
		goto L277
	} else {
		goto L298
	}
L298:
	;
	v992 = *(*float64)(unsafe.Add(mBase, uint32(v707)+232))
	*(*float64)(unsafe.Add(mBase, uint32(v707)+232)) = base.F64_add(v992, float64(1))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v29)+148))
	v1625 = v996
	v1631 = v801
	goto L241
L299:
	;
	F_ExecUpdateEpilogue(m, v29+int32(108), v29+int32(152), v238, v389, int32(0), v899)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L300
	}
L300:
	;
	v1006 = *(*float64)(unsafe.Add(mBase, uint32(v707)+232))
	*(*float64)(unsafe.Add(mBase, uint32(v707)+232)) = base.F64_add(v1006, float64(1))
	v1098 = v899
	goto L267
L301:
	;
	if v1064 != 0 {
		v1098 = v800
		goto L267
	} else {
		goto L322
	}
L302:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v29)+164))
	v1064 = v1063
	goto L301
L303:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+64))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+8))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+12))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+188))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+96))
	v1060 = m.T0[v1059].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1051, v389, v1053, v1054, v1055, int32(1), v118, int32(0))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L4
	} else {
		goto L321
	}
L304:
	;
	v1016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1013)+18)))
	if v1016 == int32(1) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+188))
	if v1020 != 0 {
		goto L308
	} else {
		goto L309
	}
L306:
	;
	v1043 = v1013
	goto L307
L307:
	;
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+20)))
	if v1044 != int32(1) {
		goto L303
	} else {
		goto L318
	}
L308:
	;
	F_ExecPendingInserts(m, v1019)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L4
	} else {
		goto L311
	}
L309:
	;
	v1024 = v1019
	goto L310
L310:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	v1026 = int32(0)
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+104))
	v1034 = F_ExecBRDeleteTriggers(m, v1024, v1025, v238, v389, v1026, v1026, v29+int32(164), v118, base.B2i32(v1031 == int32(5)))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L4
	} else {
		goto L312
	}
L311:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	v1024 = v1023
	goto L310
L312:
	;
	if v1034 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v29)+164))
	if v1038 != 0 {
		v1101 = v1038
		v1103 = v800
		goto L266
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v238)+52))
	if v1040 == int32(0) {
		goto L303
	} else {
		goto L317
	}
L316:
	;
	v1625 = int32(0)
	v1631 = v801
	goto L241
L317:
	;
	v1043 = v1040
	goto L307
L318:
	;
	v1047 = F_ExecIRDeleteTriggers(m, v709, v238, v391)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L4
	} else {
		goto L319
	}
L319:
	;
	if v1047 != 0 {
		goto L302
	} else {
		goto L320
	}
L320:
	;
	v1625 = int32(0)
	v1631 = v801
	goto L241
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = v1060
	v1064 = v1060
	goto L301
L322:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1066)+204))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1066)+104))
	if v1069 != int32(2) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1088 = int32(0)
	F_ExecARDeleteTriggers(m, v1068, v238, v389, v1088, v1087, v1088)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L4
	} else {
		goto L334
	}
L324:
	;
	v1087 = v1067
	goto L323
L325:
	;
	goto L326
L326:
	;
	if v1067 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1087 = int32(0)
	goto L323
L328:
	;
	goto L329
L329:
	;
	v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+1)))
	if v1075 != int32(1) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1087 = v1067
	goto L323
L331:
	;
	goto L332
L332:
	;
	v1078 = int32(0)
	F_ExecARUpdateTriggers(m, v1068, v238, v1078, v1078, v389, v1078, v1078, v1078, v1067, v1078)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L4
	} else {
		goto L333
	}
L333:
	;
	v1087 = v1078
	goto L323
L334:
	;
	v1092 = *(*float64)(unsafe.Add(mBase, uint32(v707)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v707)+240)) = base.F64_add(v1092, float64(1))
	v1098 = v800
	goto L267
L335:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L4
	} else {
		goto L495
	}
L336:
	;
	switch v1101 {
	case 0:
		goto L341
	case 1, 5, 6:
		goto L335
	case 2:
		goto L340
	default:
		v1585 = v1103
		goto L264
	case 4:
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v854)+4))
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+4))
	v1280 = F_ExecUpdateLockMode(m, v709, v238)
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L4
	} else {
		goto L393
	}
L339:
	;
	v1255 = int32(0)
	v1258 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[4]))
	if v1258 < int32(2) {
		v1625 = v1255
		v1631 = v1255
		goto L241
	} else {
		goto L388
	}
L340:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v29)+136))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v709)+64))
	if v1118 != v1119 {
		goto L12
	} else {
		goto L343
	}
L341:
	;
	if base.B2i32(v705&int32(1) == int32(0))|base.B2i32(v856 == int32(7)) != 0 {
		v1585 = v1103
		goto L264
	} else {
		goto L342
	}
L342:
	;
	v1114 = *(*int64)(unsafe.Add(mBase, uint32(v709)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v709)+112)) = v1114 + int64(1)
	v1585 = v1103
	goto L264
L343:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v29)+132))
	if base.Ui32(v1121) < base.Ui32(int32(3)) {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L4
	} else {
		goto L384
	}
L345:
	;
	v1241 = int32(0)
	goto L344
L346:
	;
	goto L347
L347:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[5]))
	if v1132 == v1121 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1241 = int32(1)
	goto L344
L349:
	;
	goto L350
L350:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[6]))
	if v1136 <= int32(0) {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v1241 = v1233
	goto L344
L352:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[7]))
	if v1140 == int32(0) {
		v1233 = int32(0)
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[8]))
	v1204 = int32(0)
	v1206 = v1136 - int32(1)
	goto L374
L355:
	;
	v1145 = v1140
	goto L356
L356:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+20))
	if v1150 == int32(4) {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v1233 = int32(0)
	goto L351
L358:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+80))
	if v1197 != 0 {
		v1145 = v1197
		goto L356
	} else {
		goto L373
	}
L359:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1145)))
	if v1153 == int32(0) {
		goto L358
	} else {
		goto L360
	}
L360:
	;
	v1156 = int32(1)
	if v1121 == v1153 {
		v1233 = v1156
		goto L351
	} else {
		goto L361
	}
L361:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+52))
	v1160 = v1158 - int32(1)
	if v1160 < int32(0) {
		goto L358
	} else {
		goto L362
	}
L362:
	;
	v1165 = int32(0)
	v1167 = v1160
	goto L363
L363:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+48))
	v1173 = int32(2)
	v1174 = base.I32_div_s(v1167-v1165, v1173)
	v1175 = v1174 + v1165
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1171+v1175<<(uint(v1173)%32))))
	if v1179 == v1121 {
		v1233 = v1156
		goto L351
	} else {
		goto L365
	}
L364:
	;
	goto L358
L365:
	;
	v1183 = F_TransactionIdPrecedes(m, v1179, v1121)
	mBase = m.M
	if v1183 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1184 = v1175 + int32(1)
	goto L368
L367:
	;
	v1184 = v1165
	goto L368
L368:
	;
	if v1183 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1187 = v1167
	goto L371
L370:
	;
	v1187 = v1175 - int32(1)
	goto L371
L371:
	;
	if v1184 <= v1187 {
		v1165 = v1184
		v1167 = v1187
		goto L363
	} else {
		goto L372
	}
L372:
	;
	goto L364
L373:
	;
	goto L357
L374:
	;
	v1211 = int32(2)
	v1212 = base.I32_div_s(v1206-v1204, v1211)
	v1213 = v1212 + v1204
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1202+v1213<<(uint(v1211)%32))))
	v1218 = base.B2i32(v1217 == v1121)
	if v1217 == v1121 {
		v1233 = v1218
		goto L351
	} else {
		goto L376
	}
L375:
	;
	v1233 = v1218
	goto L351
L376:
	;
	v1221 = base.B2i32(base.Ui32(v1217) < base.Ui32(v1121))
	if base.Ui32(v1217) < base.Ui32(v1121) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1222 = v1213 + int32(1)
	goto L379
L378:
	;
	v1222 = v1204
	goto L379
L379:
	;
	if base.Ui32(v1217) < base.Ui32(v1121) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1225 = v1206
	goto L382
L381:
	;
	v1225 = v1213 - int32(1)
	goto L382
L382:
	;
	if v1222 <= v1225 {
		v1204 = v1222
		v1206 = v1225
		goto L374
	} else {
		goto L383
	}
L383:
	;
	goto L375
L384:
	;
	if v1241 != 0 {
		goto L11
	} else {
		goto L385
	}
L385:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_13), int32(0))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3333), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L4
	} else {
		goto L389
	}
L389:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L4
	} else {
		goto L390
	}
L390:
	;
	F_errmsg(m, int32(_a_F_ExecModifyTable_15), int32(0))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L4
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3340), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L4
	} else {
		goto L392
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	if v1279 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L394:
	;
	v1289 = int32(0)
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v709)+8))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v709)+64))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+188))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+104))
	v1296 = m.T0[v1295].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1277, v389, v1290, v1288, v1291, v1280, v1289, int32(2), v118)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L4
	} else {
		goto L399
	}
L395:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v1285 = F_EvalPlanQualSlot(m, v791, v1277, v1284)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L4
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	v1288 = v1287
	goto L394
L398:
	;
	v1288 = v1285
	goto L394
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = v1296
	if v1296 != 0 {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L4
	} else {
		goto L492
	}
L401:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v29)+136))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v709)+64))
	if v1412 != v1413 {
		goto L8
	} else {
		goto L447
	}
L402:
	;
	switch v1296 - int32(2) {
	case 0:
		goto L401
	default:
		goto L400
	case 2:
		v1625 = v1289
		v1631 = int32(0)
		goto L241
	}
L403:
	;
	goto L404
L404:
	;
	v1302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+4)))
	if v1302 == int32(_a_F_ExecModifyTable_11) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389))))
	v1306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+2)))
	if v1305&v1306 == int32(_a_F_ExecModifyTable_16) {
		goto L10
	} else {
		goto L408
	}
L406:
	;
	goto L407
L407:
	;
	if v1279 != 0 {
		v1408 = v801
		v1410 = v811
		goto L409
	} else {
		goto L410
	}
L408:
	;
	goto L407
L409:
	;
	if v1410 != 0 {
		v800 = v1103
		v801 = v1408
		v811 = v1410
		goto L246
	} else {
		goto L446
	}
L410:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v1311 = F_EvalPlanQual(m, v791, v1277, v1310, v1288)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L4
	} else {
		goto L411
	}
L411:
	;
	if v1311 == int32(0) {
		v1625 = v1289
		v1631 = v801
		goto L241
	} else {
		goto L412
	}
L412:
	;
	v1315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+4)))
	if v1315&int32(2) != 0 {
		v1625 = v1289
		v1631 = v801
		goto L241
	} else {
		goto L413
	}
L413:
	;
	v1318 = int32(*(*int16)(unsafe.Add(mBase, uint32(v238)+24)))
	v1319 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1311)+6)))
	if v1319 < v1318 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	F_slot_getsomeattrs_int(m, v1311, v1318)
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L4
	} else {
		goto L417
	}
L415:
	;
	goto L416
L416:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+20))
	v1325 = int32(1)
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1323+v1318-v1325))))
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+49)))
	if v1328 == v1325 {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	goto L416
L418:
	;
	v1331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+172)))
	if v1331 != 0 {
		goto L421
	} else {
		goto L422
	}
L419:
	;
	goto L420
L420:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[2]))
	if v1347 != 0 {
		goto L426
	} else {
		goto L427
	}
L421:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	F_UnlockTuple(m, v1332, v29+int32(168), int32(7))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L4
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	F_LockTuple(m, v1338, v389, int32(7))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L4
	} else {
		goto L425
	}
L424:
	;
	goto L423
L425:
	;
	v1342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+172)) = uint16(v1342)
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+168)) = v1344
	goto L420
L426:
	;
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecModifyTable[3])))
	if v1349&int32(1) == int32(0) {
		goto L6
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+188))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+60))
	v1358 = m.T0[v1357].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1277, v389, int32(_a_F_ExecModifyTable_12), v1355)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L4
	} else {
		goto L430
	}
L429:
	;
	goto L428
L430:
	;
	if v1358 == int32(0) {
		goto L9
	} else {
		goto L431
	}
L431:
	;
	if v801&(v1327^int32(-1)) == int32(0) {
		goto L432
	} else {
		goto L433
	}
L432:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v784)))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v707)+36))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+20))
	if v1390 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L433:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v238)+176))
	if v1367 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1408 = int32(1)
	v1410 = v811
	goto L409
L435:
	;
	goto L436
L436:
	;
	v1371 = int32(_a_F_ExecModifyTable_10)
	v1372 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v708)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v1374
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+20))
	v1379 = m.T0[v1378].(func(*base.Module, int32, int32, int32) int32)(m, v1367, v708, v29+int32(175))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L4
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v1372
	if v1379 == int32(0) {
		goto L432
	} else {
		goto L438
	}
L438:
	;
	v1408 = int32(1)
	v1410 = v811
	goto L409
L439:
	;
	v1408 = int32(0)
	v1410 = v1388
	goto L409
L440:
	;
	goto L441
L441:
	;
	if v1388 == int32(0) {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1396 = int32(0)
	v1625 = v1396
	v1631 = v1396
	goto L241
L443:
	;
	goto L444
L444:
	;
	v1398 = int32(0)
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v238)+172))
	if v1399 == v1398 {
		v1408 = v1398
		v1410 = v1388
		goto L409
	} else {
		goto L445
	}
L445:
	;
	v1402 = *(*float64)(unsafe.Add(mBase, uint32(v1390)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1390)+32)) = base.F64_add(v1402, float64(1))
	v1408 = v1398
	v1410 = v1388
	goto L409
L446:
	;
	v1625 = int32(0)
	v1631 = v1408
	goto L241
L447:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v29)+132))
	if base.Ui32(v1415) < base.Ui32(int32(3)) {
		goto L449
	} else {
		goto L450
	}
L448:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L4
	} else {
		goto L488
	}
L449:
	;
	v1535 = int32(0)
	goto L448
L450:
	;
	goto L451
L451:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[5]))
	if v1426 == v1415 {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v1535 = int32(1)
	goto L448
L453:
	;
	goto L454
L454:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[6]))
	if v1430 <= int32(0) {
		goto L456
	} else {
		goto L457
	}
L455:
	;
	v1535 = v1527
	goto L448
L456:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[7]))
	if v1434 == int32(0) {
		v1527 = int32(0)
		goto L455
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[8]))
	v1498 = int32(0)
	v1500 = v1430 - int32(1)
	goto L478
L459:
	;
	v1439 = v1434
	goto L460
L460:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+20))
	if v1444 == int32(4) {
		goto L462
	} else {
		goto L463
	}
L461:
	;
	v1527 = int32(0)
	goto L455
L462:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+80))
	if v1491 != 0 {
		v1439 = v1491
		goto L460
	} else {
		goto L477
	}
L463:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1439)))
	if v1447 == int32(0) {
		goto L462
	} else {
		goto L464
	}
L464:
	;
	v1450 = int32(1)
	if v1415 == v1447 {
		v1527 = v1450
		goto L455
	} else {
		goto L465
	}
L465:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+52))
	v1454 = v1452 - int32(1)
	if v1454 < int32(0) {
		goto L462
	} else {
		goto L466
	}
L466:
	;
	v1459 = int32(0)
	v1461 = v1454
	goto L467
L467:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+48))
	v1467 = int32(2)
	v1468 = base.I32_div_s(v1461-v1459, v1467)
	v1469 = v1468 + v1459
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1465+v1469<<(uint(v1467)%32))))
	if v1473 == v1415 {
		v1527 = v1450
		goto L455
	} else {
		goto L469
	}
L468:
	;
	goto L462
L469:
	;
	v1477 = F_TransactionIdPrecedes(m, v1473, v1415)
	mBase = m.M
	if v1477 != 0 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v1478 = v1469 + int32(1)
	goto L472
L471:
	;
	v1478 = v1459
	goto L472
L472:
	;
	if v1477 != 0 {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1481 = v1461
	goto L475
L474:
	;
	v1481 = v1469 - int32(1)
	goto L475
L475:
	;
	if v1478 <= v1481 {
		v1459 = v1478
		v1461 = v1481
		goto L467
	} else {
		goto L476
	}
L476:
	;
	goto L468
L477:
	;
	goto L461
L478:
	;
	v1505 = int32(2)
	v1506 = base.I32_div_s(v1500-v1498, v1505)
	v1507 = v1506 + v1498
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1496+v1507<<(uint(v1505)%32))))
	v1512 = base.B2i32(v1511 == v1415)
	if v1511 == v1415 {
		v1527 = v1512
		goto L455
	} else {
		goto L480
	}
L479:
	;
	v1527 = v1512
	goto L455
L480:
	;
	v1515 = base.B2i32(base.Ui32(v1511) < base.Ui32(v1415))
	if base.Ui32(v1511) < base.Ui32(v1415) {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	v1516 = v1507 + int32(1)
	goto L483
L482:
	;
	v1516 = v1498
	goto L483
L483:
	;
	if base.Ui32(v1511) < base.Ui32(v1415) {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v1519 = v1500
	goto L486
L485:
	;
	v1519 = v1507 - int32(1)
	goto L486
L486:
	;
	if v1516 <= v1519 {
		v1498 = v1516
		v1500 = v1519
		goto L478
	} else {
		goto L487
	}
L487:
	;
	goto L479
L488:
	;
	if v1535 != 0 {
		goto L7
	} else {
		goto L489
	}
L489:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_13), int32(0))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L4
	} else {
		goto L490
	}
L490:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3525), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L4
	} else {
		goto L491
	}
L491:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L492:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v29)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1553
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_17), v29+int32(32))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L4
	} else {
		goto L493
	}
L493:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3531), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L4
	} else {
		goto L494
	}
L494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L495:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v29)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v1569
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_18), v29-int32(-64))
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L4
	} else {
		goto L496
	}
L496:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3540), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L4
	} else {
		goto L497
	}
L497:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	switch v896 {
	case 0:
		goto L245
	default:
		goto L499
	case 2:
		goto L500
	case 5:
		v1625 = v1587
		v1631 = v801
		goto L241
	}
L499:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L4
	} else {
		goto L502
	}
L500:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	v1597 = F_ExecProcessReturning(m, v29+int32(108), v238, int32(4), v1594, int32(0), v1596)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L4
	} else {
		goto L501
	}
L501:
	;
	v1625 = v1597
	v1631 = v801
	goto L241
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v856
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_19), v29)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L4
	} else {
		goto L503
	}
L503:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3572), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L4
	} else {
		goto L504
	}
L504:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L505:
	;
	goto L250
L506:
	;
	v1625 = v1622
	v1631 = v801
	goto L241
L507:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	F_UnlockTuple(m, v1651, v29+int32(168), int32(7))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L4
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	if v1631 != 0 {
		v1718 = v1625
		goto L139
	} else {
		goto L511
	}
L510:
	;
	goto L509
L511:
	;
	if v1625 != 0 {
		goto L210
	} else {
		goto L512
	}
L512:
	;
	goto L213
L513:
	;
	if v1687 == int32(0) {
		v127 = v238
		goto L48
	} else {
		goto L514
	}
L514:
	;
	v1815 = v1687
	goto L19
L515:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_0), int32(0))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L4
	} else {
		goto L516
	}
L516:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_20), int32(_a_F_ExecModifyTable_5))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L4
	} else {
		goto L517
	}
L517:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L518:
	;
	v1718 = v1715
	goto L139
L519:
	;
	v1815 = v1718
	goto L19
L520:
	;
	F_ExecPendingInserts(m, v32)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L4
	} else {
		goto L523
	}
L521:
	;
	goto L522
L522:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	switch v1749 - int32(2) {
	case 0:
		goto L525
	case 1:
		goto L529
	case 2:
		goto L528
	case 3:
		goto L527
	default:
		goto L526
	}
L523:
	;
	goto L522
L524:
	;
	v1811 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v1811)
	v1815 = int32(0)
	goto L19
L525:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASUpdateTriggers(m, v1806, v1748, v1807)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L4
	} else {
		goto L549
	}
L526:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L4
	} else {
		goto L546
	}
L527:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	if v1768&int32(4) != 0 {
		goto L536
	} else {
		goto L537
	}
L528:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASDeleteTriggers(m, v1764, v1748, v1765)
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L4
	} else {
		goto L535
	}
L529:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1752)+132))
	if v1753 == int32(2) {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	F_ExecASUpdateTriggers(m, v1756, v1748, v1757)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L4
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASInsertTriggers(m, v1760, v1748, v1761)
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L4
	} else {
		goto L534
	}
L533:
	;
	goto L532
L534:
	;
	goto L524
L535:
	;
	goto L524
L536:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASDeleteTriggers(m, v1771, v1748, v1772)
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L4
	} else {
		goto L539
	}
L537:
	;
	v1776 = v1768
	goto L538
L538:
	;
	if v1776&int32(2) != 0 {
		goto L540
	} else {
		goto L541
	}
L539:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v1776 = v1775
	goto L538
L540:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASUpdateTriggers(m, v1779, v1748, v1780)
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L4
	} else {
		goto L543
	}
L541:
	;
	v1784 = v1776
	goto L542
L542:
	;
	if v1784&int32(1) == int32(0) {
		goto L524
	} else {
		goto L544
	}
L543:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v1784 = v1783
	goto L542
L544:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASInsertTriggers(m, v1789, v1748, v1790)
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L4
	} else {
		goto L545
	}
L545:
	;
	goto L524
L546:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_0), int32(0))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L4
	} else {
		goto L547
	}
L547:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(4052), int32(_a_F_ExecModifyTable_21))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L4
	} else {
		goto L548
	}
L548:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L549:
	;
	goto L524
L550:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_22), int32(0))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L4
	} else {
		goto L551
	}
L551:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_23), int32(_a_F_ExecModifyTable_5))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L4
	} else {
		goto L552
	}
L552:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L553:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_24), int32(0))
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L4
	} else {
		goto L554
	}
L554:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_25), int32(_a_F_ExecModifyTable_5))
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L4
	} else {
		goto L555
	}
L555:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L556:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_26), int32(0))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L4
	} else {
		goto L557
	}
L557:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3124), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L4
	} else {
		goto L558
	}
L558:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L559:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_27), int32(0))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L4
	} else {
		goto L560
	}
L560:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3286), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L4
	} else {
		goto L561
	}
L561:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L562:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L4
	} else {
		goto L563
	}
L563:
	;
	F_errmsg(m, int32(_a_F_ExecModifyTable_28), int32(0))
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L4
	} else {
		goto L564
	}
L564:
	;
	F_errhint(m, int32(_a_F_ExecModifyTable_29), int32(0))
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L4
	} else {
		goto L565
	}
L565:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3322), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L4
	} else {
		goto L566
	}
L566:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = int32(_a_F_ExecModifyTable_30)
	F_errmsg(m, int32(_a_F_ExecModifyTable_31), v29+int32(16))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L4
	} else {
		goto L568
	}
L568:
	;
	F_errhint(m, int32(_a_F_ExecModifyTable_32), int32(0))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L4
	} else {
		goto L569
	}
L569:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3330), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L4
	} else {
		goto L570
	}
L570:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L571:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L4
	} else {
		goto L572
	}
L572:
	;
	F_errmsg(m, int32(_a_F_ExecModifyTable_33), int32(0))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L4
	} else {
		goto L573
	}
L573:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3399), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L4
	} else {
		goto L574
	}
L574:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L575:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_26), int32(0))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L4
	} else {
		goto L576
	}
L576:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3453), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L4
	} else {
		goto L577
	}
L577:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L578:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L4
	} else {
		goto L579
	}
L579:
	;
	F_errmsg(m, int32(_a_F_ExecModifyTable_28), int32(0))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L4
	} else {
		goto L580
	}
L580:
	;
	F_errhint(m, int32(_a_F_ExecModifyTable_29), int32(0))
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L4
	} else {
		goto L581
	}
L581:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3514), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L4
	} else {
		goto L582
	}
L582:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L583:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = int32(_a_F_ExecModifyTable_30)
	F_errmsg(m, int32(_a_F_ExecModifyTable_31), v29+int32(48))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L4
	} else {
		goto L584
	}
L584:
	;
	F_errhint(m, int32(_a_F_ExecModifyTable_32), int32(0))
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L4
	} else {
		goto L585
	}
L585:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3522), int32(_a_F_ExecModifyTable_14))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L4
	} else {
		goto L586
	}
L586:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L587:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_34), int32(0))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L4
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_35), int32(1264), int32(_a_F_ExecModifyTable_36))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L4
	} else {
		goto L589
	}
L589:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecProjectSRF(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
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
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v182 int32
	_ = v182
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v318 int64
	_ = v318
	var v320 int64
	_ = v320
	var v325 int64
	_ = v325
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	v3 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	m.T0[v23].(func(*base.Module, int32))(m, v21)
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
	v28 = int32(_a_F_ExecProjectSRF_0)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v31
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v33 < v35 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return int32(0)
L4:
	;
	v49 = v3
	v52 = v3
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v29
	goto L3
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v58 = v57 + v49
	v60 = v49 << (uint(int32(2)) % 32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v62 = v60 + v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v64 = v63 + v60
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65+v60)))
	if l1 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v29
	if v571&int32(1) == int32(0) {
		goto L3
	} else {
		goto L116
	}
L9:
	;
	v577 = v49 + int32(1)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v577 < v578 {
		v49 = v577
		v52 = v571
		goto L7
	} else {
		goto L115
	}
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v77 == int32(391) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v70 != int32(2) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = int32(0)
	v75 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v75)
	v571 = v52
	goto L9
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v81 = m.G0
	v83 = v81 - int32(80)
	m.G0 = v83
	F_check_stack_depth(m)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v552 = m.T0[v551].(func(*base.Module, int32, int32, int32) int32)(m, v67, v20, v58)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L114
	}
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	if v87 != 0 {
		v460 = v87
		goto L18
	} else {
		goto L19
	}
L17:
	;
	m.G0 = v83 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v524
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v546 = base.B2i32(v543 != int32(2)) | v52
	if v543 != int32(1) {
		v571 = v546
		goto L9
	} else {
		goto L113
	}
L18:
	;
	v476 = int32(0)
	v477 = int32(_a_F_ExecProjectSRF_0)
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0]))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v481
	v485 = F_tuplestore_gettupleslot(m, v460, int32(1), v476, v480)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L100
	}
L19:
	;
	goto L23
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L96
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L92
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L88
	}
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v67)+60))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+58)))
	if v108 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(2)
	v405 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v405)
	v524 = int32(0)
	goto L17
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v83 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = int32(383)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v83)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v83)+28)) = int64(4294967299)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v209
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+26)))
	if v215 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v111 = int32(_a_F_ExecProjectSRF_0)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v80
	if v113 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+58)) = uint8(v182)
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v112
	goto L25
L30:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v118 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v128 = int32(0)
	goto L32
L32:
	;
	v145 = v107 + int32(20) + v128<<(uint(int32(3))%32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v128<<(uint(int32(2))%32))))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+20))
	v154 = m.T0[v153].(func(*base.Module, int32, int32, int32) int32)(m, v150, v20, v145+int32(4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v154
	v158 = v128 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v158 < v159 {
		v128 = v158
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v275 = v83 + int32(48)
	F_pgstat_init_function_usage(m, v107, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L45
	}
L37:
	;
	v218 = int32(0)
	v219 = int32(*(*int16)(unsafe.Add(mBase, uint32(v107)+18)))
	if v219 <= v218 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v226 = v218
	goto L39
L39:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v226<<(uint(int32(3))%32))+24)))
	if v244 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v250 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v250)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(2)
	v524 = int32(0)
	goto L17
L41:
	;
	v248 = v226 + int32(1)
	if v219 != v248 {
		v226 = v248
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	goto L36
L45:
	;
	v278 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+16)) = uint8(v278)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+36)) = v278
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v284 = m.T0[v283].(func(*base.Module, int32) int32)(m, v107)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v286)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v83)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v288
	v298 = m.G0
	v300 = v298 - int32(16)
	m.G0 = v300
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v302 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	if v337 != int32(2) {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	F___clock_gettime(m, int32(1), v300)
	mBase = m.M
	v305 = int32(_a_F_ExecProjectSRF_1)
	v306 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProjectSRF[1]))
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v275)+16))
	v309 = int64(*(*int32)(unsafe.Add(mBase, uint32(v300)+8)))
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v275)+24))
	v315 = v309 + v310*int64(1000000000) - v314
	*(*int64)(unsafe.Add(mBase, _c_F_ExecProjectSRF[1])) = v308 + v315
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v275)+8))
	if v288 != int32(1) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	m.G0 = v300 + int32(16)
	goto L47
L51:
	;
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v302)))
	*(*int64)(unsafe.Add(mBase, uint32(v302))) = v320 + int64(1)
	goto L53
L52:
	;
	goto L53
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v302)+8)) = v318 + v315
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v302)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v302)+16)) = v325 + (v315 - v306 + v308)
	goto L50
L54:
	;
	if v337 != int32(1) {
		goto L22
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v83)+36))
	if v353 != 0 {
		goto L21
	} else {
		goto L61
	}
L57:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v342 != int32(1) {
		v524 = v284
		goto L17
	} else {
		goto L58
	}
L58:
	;
	v345 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+58)) = uint8(v345)
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+59)))
	if v347 != 0 {
		v524 = v284
		goto L17
	} else {
		goto L59
	}
L59:
	;
	F_RegisterExprContextCallback(m, v20, int32(635), v67)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v351 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+59)) = uint8(v351)
	v524 = v284
	goto L17
L61:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	if v354 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v83)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+44)) = v354
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
	if v357 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	goto L24
L65:
	;
	v360 = int32(_a_F_ExecProjectSRF_0)
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0]))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v67)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v363
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
	if v365 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if v355 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	v370 = v365
	goto L70
L69:
	;
	if v355 == int32(0) {
		goto L20
	} else {
		goto L71
	}
L70:
	;
	v372 = F_MakeTupleTableSlot(m, v370, int32(_a_F_ExecProjectSRF_2))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	v368 = F_CreateTupleDescCopy(m, v355)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v370 = v368
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+48)) = v372
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v361
	goto L67
L74:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+59)))
	if v390 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L75:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
	if v381 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_tupledesc_match(m, v381, v355)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v355)+12))
	if v384 != int32(-1) {
		goto L74
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	F_FreeTupleDesc(m, v355)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L74
L82:
	;
	F_RegisterExprContextCallback(m, v20, int32(635), v67)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	v396 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+59)) = uint8(v396)
	goto L84
L86:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	if v400 == int32(0) {
		goto L23
	} else {
		goto L87
	}
L87:
	;
	v460 = v400
	goto L18
L88:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v415
	F_errmsg(m, int32(_a_F_ExecProjectSRF_3), v83)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_ExecProjectSRF_4), int32(686), int32(_a_F_ExecProjectSRF_5))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_ExecProjectSRF_6), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_ExecProjectSRF_4), int32(667), int32(_a_F_ExecProjectSRF_5))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_ExecProjectSRF_7), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_ExecProjectSRF_4), int32(895), int32(_a_F_ExecProjectSRF_8))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v478
	if v485 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v489 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v489
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+56)))
	if v491 == v489 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	F_tuplestore_end(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L112
	}
L104:
	;
	v494 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v494)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
	v497 = F_ExecFetchSlotHeapTupleDatum(m, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
	v500 = int32(*(*int16)(unsafe.Add(mBase, uint32(v499)+6)))
	if v500 <= int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v524 = v497
	goto L17
L108:
	;
	F_slot_getsomeattrs_int(m, v499, int32(1))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v499)+20))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v507)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v499)+16))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v524 = v510
	goto L17
L111:
	;
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(2)
	v518 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v518)
	v524 = v476
	goto L17
L113:
	;
	v549 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v549)
	v571 = v546
	goto L9
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v552
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(0)
	v571 = v52
	goto L9
L115:
	;
	goto L8
L116:
	;
	v586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v588 = v586 & int32(_a_F_ExecProjectSRF_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v588)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v591)
	goto L117
L117:
	;
	return v21
}
func F_ExecSort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
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
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v219 int64
	_ = v219
	var v223 int64
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSort[0]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v16 != 0 {
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
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+149)))
	if v247 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v241 = v17
	goto L6
L8:
	;
	goto L9
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v29 = v23 | v24<<(uint(v19)%32)&int32(2)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+149)))
	if v30 == v19 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	if v59 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22+v33<<(uint(int32(4))%32))+88))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSort[1]))
	v46 = F_tuplesort_begin_datum(m, v37, v39, v41, v43, v45, v29)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSort[1]))
	v56 = F_tuplesort_begin_heap(m, v22, v48, v49, v50, v51, v52, v54, int32(0), v29)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	v58 = v46
	goto L10
L15:
	;
	v58 = v56
	goto L10
L16:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+236))
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v58
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+149)))
	if v92 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L19:
	;
	goto L18
L20:
	;
	goto L19
L21:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+72)) = uint32(v62)
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+68)) = uint8(v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
	if v80 != 0 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	if int64(1073741823) < v62 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if int64(1073741823) < v62 {
		goto L20
	} else {
		goto L27
	}
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+232))
	if v68 != int32(-1) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	goto L21
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
	v83 = v82
	goto L30
L29:
	;
	v83 = v79
	goto L30
L30:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v84
	goto L20
L31:
	;
	F_tuplesort_performsort(m, v58)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L59
	}
L32:
	;
	goto L35
L33:
	;
	goto L34
L34:
	;
	goto L49
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v102 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v106 = m.T0[v105].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	if v106 == int32(0) {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
	if v110&int32(2) != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v106)+6)))
	if v113 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_slot_getsomeattrs_int(m, v106, int32(1))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	F_tuplesort_putdatum(m, v58, v120, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	goto L35
L49:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v132 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v136 = m.T0[v135].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	if v136 == int32(0) {
		goto L31
	} else {
		goto L56
	}
L56:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+4)))
	if v140&int32(2) != 0 {
		goto L31
	} else {
		goto L57
	}
L57:
	;
	F_tuplesort_puttupleslot(m, v58, v136)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	goto L49
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v15
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v155)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+129)) = uint8(v157)
	v159 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v161 == int32(0) {
		v241 = v58
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v164 != int32(1) {
		v241 = v58
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSort[2]))
	v173 = v161 + v168<<(uint(int32(4))%32) + int32(8)
	v174 = int32(0)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v58)+128))
	if v178 == v174 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v241 = v58
	goto L6
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v216
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v58)+112))
	v223 = base.I64_div_s(v219+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+8)) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v58)+124))
	switch v225 - int32(3) {
	case 0:
		goto L78
	case 1:
		v236 = v225
		goto L75
	case 2:
		goto L77
	default:
		goto L76
	}
L64:
	;
	if v195&int32(255) != base.B2i32(v178 != int32(0)) {
		goto L70
	} else {
		goto L71
	}
L65:
	;
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v58)+96))
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v58)+88))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+120)))
	v195 = v184
	v196 = v181 - v182
	goto L64
L66:
	;
	goto L67
L67:
	;
	v185 = F_LogicalTapeSetBlocks(m, v178)
	mBase = m.M
	v187 = v185 << (uint(int64(13)) % 64)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+120)))
	if v189 != 0 {
		v195 = int32(1)
		v196 = v187
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v190 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+120)) = uint8(v190)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+112)) = v187
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v58)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+124)) = v193
	v216 = v174
	goto L63
L69:
	;
	v216 = int32(1)
	goto L63
L70:
	;
	if v195&int32(1) != 0 {
		v216 = v174
		goto L63
	} else {
		goto L74
	}
L71:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v58)+112))
	if v196 <= v202 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+120)) = uint8(v195)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+112)) = v196
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v58)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+124)) = v206
	if v195&int32(1) == int32(0) {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v216 = v174
	goto L63
L74:
	;
	goto L69
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v236
	goto L62
L76:
	;
	v236 = int32(0)
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = int32(8)
	goto L62
L78:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+69)))
	if v230 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v231 = int32(1)
	goto L81
L80:
	;
	v231 = int32(2)
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v231
	goto L62
L82:
	;
	return v246
L83:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	m.T0[v251].(func(*base.Module, int32))(m, v246)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v274 = int32(0)
	v276 = F_tuplesort_gettupleslot(m, v241, base.B2i32(v15 == int32(1)), v274, v246, v274)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L90
	}
L86:
	;
	v256 = int32(0)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v246)+20))
	v260 = F_tuplesort_getdatum(m, v241, base.B2i32(v15 == int32(1)), v256, v257, v258, v256)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	if v260 == int32(0) {
		goto L82
	} else {
		goto L88
	}
L88:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v246)+4)))
	v266 = v264 & int32(_a_F_ExecSort_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v246)+4)) = uint16(v266)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	*(*uint16)(unsafe.Add(mBase, uint32(v246)+6)) = uint16(v269)
	goto L89
L89:
	;
	return v246
L90:
	;
	goto L82
}
func F_ExecWithCheckOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v227 int64
	_ = v227
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(144)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	if v21 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = F_MakePerTupleExprContext(m, l3)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v26 = v21
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = l2
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v37 = v5
	goto L9
L4:
	;
	return
L5:
	;
	v26 = v24
	goto L3
L6:
	;
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = base.I64_rotl(v255, int64(32))
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_0), v17+int32(128))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L75
	}
L7:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = base.I64_rotl(v241, int64(32))
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_1), v17+int32(96))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L73
	}
L8:
	;
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = base.I64_rotl(v227, int64(32))
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_2), v17-int32(-64))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L71
	}
L9:
	;
	v44 = int32(0)
	if v29 == v44 {
		v54 = v44
		goto L11
	} else {
		goto L12
	}
L10:
	;
	m.G0 = v17 + int32(144)
	return
L11:
	;
	if v28 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v48 <= v37 {
		v54 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v54 = v50 + v37<<(uint(int32(2))%32)
	goto L11
L14:
	;
	goto L10
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if base.B2i32(v54 == int32(0))|base.B2i32(v59 <= v37) != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v62 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v66 != l0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v37 = v37 + int32(1)
	goto L9
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v62+v37<<(uint(int32(2))%32))))
	if v71 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v74 = int32(_a_F_ExecWithCheckOptions_3)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWithCheckOptions[0]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWithCheckOptions[0])) = v77
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v82 = m.T0[v81].(func(*base.Module, int32, int32, int32) int32)(m, v71, v26, v17+int32(143))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWithCheckOptions[0])) = v75
	if v82 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	switch v86 {
	case 0:
		goto L27
	case 1, 2:
		goto L26
	case 3:
		goto L24
	case 4, 5:
		goto L25
	default:
		goto L23
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L68
	}
L24:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L63
	}
L25:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L58
	}
L26:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L53
	}
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v87 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+56))
	v118 = F_ExecBuildSlotValueDescription(m, v117, v114, v116, v113)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L44
	}
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+52))
	v92 = F_build_attrmap_by_name_if_req(m, v88, v90, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v107 = F_ExecGetInsertedCols(m, l1, l3)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L41
	}
L32:
	;
	if v92 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v95 = F_MakeTupleTableSlot(m, v90, int32(_a_F_ExecWithCheckOptions_4))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	v99 = l2
	goto L35
L35:
	;
	v100 = F_ExecGetInsertedCols(m, v87, l3)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	v97 = F_execute_attr_map_slot(m, v92, l2, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v99 = v97
	goto L35
L38:
	;
	v102 = F_ExecGetUpdatedCols(m, v87, l3)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v104 = F_bms_union(m, v100, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v113 = v104
	v114 = v99
	v115 = v106
	v116 = v90
	goto L28
L41:
	;
	v109 = F_ExecGetUpdatedCols(m, l1, l3)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v111 = F_bms_union(m, v107, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v113 = v111
	v114 = l2
	v115 = v19
	v116 = v20
	goto L28
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(260))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v127
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_5), v17+int32(32))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	if v118 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v118
	F_errdetail(m, int32(_a_F_ExecWithCheckOptions_6), v17+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2327), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	if v145 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v153
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_9), v17+int32(48))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2340), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	if v165 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v173
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_10), v17+int32(80))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2353), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	if v185 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v193
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_11), v17+int32(112))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2365), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v209
	F_errmsg_internal(m, int32(_a_F_ExecWithCheckOptions_12), v17)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2368), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2335), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
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
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2348), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2360), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__equalCaseWhen(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = F_equal(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v11 = F_equal(m, v9, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v14 = v11
				return v14
			}
		} else {
			v14 = int32(0)
			return v14
		}
	}
}
func F_each_worker_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_MemoryContextDelete(m, v30)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L52
	}
L2:
	;
	v102 = v97
	v103 = v94
	v105 = v95
	goto L29
L3:
	;
	v94 = v4
	v95 = v4
	v97 = int32(2)
	goto L2
L4:
	;
	v94 = v90
	v95 = v45
	v97 = int32(1)
	goto L2
L5:
	;
	return
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	if v17&int32(32) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(2))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
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
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L25
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0]))
	v30 = F_AllocSetContextCreateInternal(m, v25, int32(_a_F_each_worker_jsonb_0), int32(0), int32(_a_F_each_worker_jsonb_1), int32(_a_F_each_worker_jsonb_2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v34 = F_JsonbIteratorInit(m, v15+int32(4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v34
	v42 = F_JsonbIteratorNext(m, v12+int32(44), v12+int32(24), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	v44 = int32(_a_F_each_worker_jsonb_3)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0])) = v30
	v48 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)) = uint16(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v52 = F_cstring_to_text_with_len(m, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	switch v42 {
	case 0:
		goto L1
	case 1:
		goto L13
	default:
		goto L3
	}
L15:
	;
	v57 = v12 + int32(24)
	v59 = F_JsonbIteratorNext(m, v12+int32(44), v57, int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v52
	if l2 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v64 = F_JsonbValueToJsonb(m, v57)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v66 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v90 = v64
	goto L4
L21:
	;
	v69 = F_JsonbValueAsText(m, v12+int32(24))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v71)
	v94 = v4
	v95 = v45
	v97 = int32(0)
	goto L2
L24:
	;
	v90 = v69
	goto L4
L25:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	F_errmsg(m, int32(_a_F_each_worker_jsonb_4), v12)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_each_worker_jsonb_5), int32(1989), int32(_a_F_each_worker_jsonb_6))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	switch v102 {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		goto L33
	}
L31:
	;
	v102 = int32(0)
	v105 = v139
	goto L29
L32:
	;
	v102 = int32(1)
	v103 = v178
	v105 = v176
	goto L29
L33:
	;
	goto L38
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v103
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	F_tuplestore_putvalues(m, v109, v110, v12+int32(16), v12+int32(14))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L36
	}
L35:
	;
	v176 = v105
	v178 = int32(0)
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0])) = v105
	F_MemoryContextReset(m, v30)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v102 = int32(2)
	goto L29
L38:
	;
	v136 = F_JsonbIteratorNext(m, v12+int32(44), v12+int32(24), int32(1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	v138 = int32(_a_F_each_worker_jsonb_3)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0])) = v30
	v142 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)) = uint16(v142)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v146 = F_cstring_to_text_with_len(m, v144, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	switch v136 {
	case 0:
		goto L1
	case 1:
		goto L40
	default:
		goto L38
	}
L42:
	;
	v153 = F_JsonbIteratorNext(m, v12+int32(44), v12+int32(24), int32(1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v146
	if l2 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v156 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v167 = F_JsonbValueToJsonb(m, v12+int32(24))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L51
	}
L47:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v159)
	goto L31
L48:
	;
	goto L49
L49:
	;
	v163 = F_JsonbValueAsText(m, v12+int32(24))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v176 = v139
	v178 = v163
	goto L32
L51:
	;
	v176 = v139
	v178 = v167
	goto L32
L52:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v192)
	m.G0 = v12 + int32(48)
	return
}
func F_ec_search_derived_clause_for_ems(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 float64
	_ = v33
	var v36 float64
	_ = v36
	var v39 float64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v54 int64
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v76 int64
	_ = v76
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
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
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
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
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
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
	var v322 int32
	_ = v322
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
	var v331 int32
	_ = v331
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
	var v344 int32
	_ = v344
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
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v16 != 0 {
		v188 = v16
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L9
	} else {
		goto L114
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return v513
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l4
	v195 = base.B2i32(base.Ui32(l2) < base.Ui32(l3))
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L53
	} else {
		goto L54
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v17 == int32(0) {
		v513 = v6
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if int32(32) <= v20 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v25 = F_MemoryContextAllocZero(m, v23, int32(32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v137 == int32(0) {
		v513 = v6
		goto L2
	} else {
		goto L36
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v23
	v33 = float64(4.294967296e+09)
	v36 = base.F64_div(base.F64_convert_i32_u(v20), float64(0.9))
	if base.F64_ge(v36, v33) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = v33
	goto L13
L12:
	;
	v39 = v36
	goto L13
L13:
	;
	v40 = base.I64_trunc_sat_f64_u(v39)
	if base.Ui64(v40) <= base.Ui64(int64(2)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v43 = int64(2)
	goto L16
L15:
	;
	v43 = v40
	goto L16
L16:
	;
	v44 = int64(1)
	if v43&(v43-v44) == int64(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v54 = v43
	goto L19
L18:
	;
	v54 = v44 << (uint(int64(64)-base.I64_clz(v43)) % 64)
	goto L19
L19:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v54*int64(20)) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v63 = F_MemoryContextAllocExtended(m, v23, base.I32_wrap_i64(v54)*int32(20), int32(5))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v63
	v66 = int64(1)
	if v54&(v54-v66) == int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v76 = v54
	goto L24
L23:
	;
	v76 = v66 << (uint(int64(64)-base.I64_clz(v54)) % 64)
	goto L24
L24:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v76*int64(20)) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = base.I32_wrap_i64(v76) - int32(1)
	if v76 == int64(4294967296) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v93 = int32(-85899346)
	goto L28
L27:
	;
	v93 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v76), float64(0.9)))
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v25
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v96 == int32(0) {
		v188 = v25
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 <= int32(0) {
		v188 = v25
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v103 = int32(0)
	goto L31
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v103<<(uint(int32(2))%32))))
	F_ec_add_clause_to_derives_hash(m, l1, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L33
	}
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v125 != 0 {
		v188 = v125
		goto L3
	} else {
		goto L35
	}
L33:
	;
	v122 = v103 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v122 < v123 {
		v103 = v122
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L8
L36:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v140 <= int32(0) {
		v513 = v6
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v143 = int32(0)
	if v143 < v140 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v146 = v140
	goto L40
L39:
	;
	v146 = v143
	goto L40
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v150 = int32(0)
	goto L41
L41:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v147+v150<<(uint(int32(2))%32))))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+108))
	if base.B2i32(l3 == int32(0))&base.B2i32(l2 == v166) != 0 {
		v513 = v165
		goto L2
	} else {
		goto L43
	}
L42:
	;
	v513 = int32(0)
	goto L2
L43:
	;
	if v166 != l2 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v166 != l3 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+112))
	if v170 != l3 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v165)+60))
	if v172 == l4 {
		v513 = v165
		goto L2
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v181 = v150 + int32(1)
	if v181 != v146 {
		v150 = v181
		goto L41
	} else {
		goto L52
	}
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v165)+112))
	if v175 != l2 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v165)+60))
	if v177 == l4 {
		v513 = v165
		goto L2
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	goto L42
L53:
	;
	v196 = l3
	goto L55
L54:
	;
	v196 = l2
	goto L55
L55:
	;
	if l3 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v197 = v196
	goto L58
L57:
	;
	v197 = l2
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v197
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v199 = l2
	goto L61
L60:
	;
	v199 = l3
	goto L61
L61:
	;
	if l3 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v201 = v199
	goto L64
L63:
	;
	v201 = int32(0)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v201
	v203 = int32(12)
	v209 = int32(-1636608420)
	if v14&int32(3) != 0 {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v470 = (v463 ^ v455 - base.I32_rotl(v463, int32(24))) & v469
	v473 = v468 + v470*int32(20)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	if v474 == int32(0) {
		v513 = v6
		goto L2
	} else {
		goto L105
	}
L66:
	;
	v441 = int32(14)
	v443 = v437 ^ v438 - base.I32_rotl(v437, v441)
	v447 = v443 ^ v436 - base.I32_rotl(v443, int32(11))
	v451 = v447 ^ v437 - base.I32_rotl(v447, int32(25))
	v455 = v451 ^ v443 - base.I32_rotl(v451, int32(16))
	v459 = v455 ^ v447 - base.I32_rotl(v455, int32(4))
	v463 = v459 ^ v451 - base.I32_rotl(v459, v441)
	goto L65
L67:
	;
	switch v363 - int32(1) {
	case 0:
		v429 = v354
		v430 = v355
		v431 = v359
		goto L94
	case 1:
		v422 = v354
		v423 = v355
		v424 = v359
		goto L95
	case 2:
		v415 = v354
		v416 = v355
		v417 = v359
		goto L96
	case 3:
		v409 = v355
		v410 = v359
		goto L97
	case 4:
		v405 = v355
		v406 = v359
		goto L98
	case 5:
		v399 = v355
		v400 = v359
		goto L99
	case 6:
		v393 = v355
		v394 = v359
		goto L100
	case 7:
		v388 = v359
		goto L101
	case 8:
		v383 = v359
		goto L102
	case 9:
		v378 = v359
		goto L103
	case 10:
		goto L104
	default:
		v436 = v354
		v437 = v355
		v438 = v359
		goto L66
	}
L68:
	;
	v318 = v14
	v319 = v203
	v320 = v209
	v321 = v209
	v322 = v209
	goto L91
L69:
	;
	goto L68
L70:
	;
	goto L71
L71:
	;
	goto L75
L73:
	;
	switch v261 - int32(1) {
	case 0:
		v315 = v252
		goto L80
	case 1:
		v310 = v252
		goto L81
	case 2:
		goto L82
	case 3:
		v303 = v253
		goto L83
	case 4:
		v300 = v253
		goto L84
	case 5:
		v295 = v253
		goto L85
	case 6:
		goto L86
	case 7:
		v286 = v257
		goto L87
	case 8:
		v281 = v257
		goto L88
	case 9:
		v276 = v257
		goto L89
	case 10:
		goto L90
	default:
		v436 = v252
		v437 = v253
		v438 = v257
		goto L66
	}
L75:
	;
	goto L76
L76:
	;
	v216 = v14
	v217 = v203
	v218 = v209
	v219 = v209
	v220 = v209
	goto L77
L77:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v223 = v222 + v219
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v227 = v226 + v220
	v229 = int32(4)
	v231 = v224 + v218 - v227 ^ base.I32_rotl(v227, v229)
	v235 = v223 - v231 ^ base.I32_rotl(v231, int32(6))
	v236 = v227 + v223
	v237 = v231 + v236
	v238 = v235 + v237
	v242 = v236 - v235 ^ base.I32_rotl(v235, int32(8))
	v246 = v237 - v242 ^ base.I32_rotl(v242, int32(16))
	v250 = v238 - v246 ^ base.I32_rotl(v246, int32(19))
	v251 = v242 + v238
	v252 = v246 + v251
	v253 = v250 + v252
	v257 = v251 - v250 ^ base.I32_rotl(v250, v229)
	v258 = int32(12)
	v259 = v216 + v258
	v261 = v217 - v258
	if base.Ui32(int32(11)) < base.Ui32(v261) {
		v216 = v259
		v217 = v261
		v218 = v252
		v219 = v253
		v220 = v257
		goto L77
	} else {
		goto L79
	}
L78:
	;
	goto L73
L79:
	;
	goto L78
L80:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	v436 = v315 + v316
	v437 = v253
	v438 = v257
	goto L66
L81:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	v315 = v311<<(uint(int32(8))%32) + v310
	goto L80
L82:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+2)))
	v310 = v306<<(uint(int32(16))%32) + v252
	goto L81
L83:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v436 = v304 + v252
	v437 = v303
	v438 = v257
	goto L66
L84:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+4)))
	v303 = v300 + v301
	goto L83
L85:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+5)))
	v300 = v296<<(uint(int32(8))%32) + v295
	goto L84
L86:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+6)))
	v295 = v291<<(uint(int32(16))%32) + v253
	goto L85
L87:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	v436 = v287 + v252
	v437 = v289 + v253
	v438 = v286
	goto L66
L88:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+8)))
	v286 = v282<<(uint(int32(8))%32) + v281
	goto L87
L89:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+9)))
	v281 = v277<<(uint(int32(16))%32) + v276
	goto L88
L90:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+10)))
	v276 = v272<<(uint(int32(24))%32) + v257
	goto L89
L91:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v325 = v324 + v321
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v329 = v328 + v322
	v331 = int32(4)
	v333 = v326 + v320 - v329 ^ base.I32_rotl(v329, v331)
	v337 = v325 - v333 ^ base.I32_rotl(v333, int32(6))
	v338 = v329 + v325
	v339 = v333 + v338
	v340 = v337 + v339
	v344 = v338 - v337 ^ base.I32_rotl(v337, int32(8))
	v348 = v339 - v344 ^ base.I32_rotl(v344, int32(16))
	v352 = v340 - v348 ^ base.I32_rotl(v348, int32(19))
	v353 = v344 + v340
	v354 = v348 + v353
	v355 = v352 + v354
	v359 = v353 - v352 ^ base.I32_rotl(v352, v331)
	v360 = int32(12)
	v361 = v318 + v360
	v363 = v319 - v360
	if base.Ui32(int32(11)) < base.Ui32(v363) {
		v318 = v361
		v319 = v363
		v320 = v354
		v321 = v355
		v322 = v359
		goto L91
	} else {
		goto L93
	}
L92:
	;
	goto L67
L93:
	;
	goto L92
L94:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	v436 = v429 + v432
	v437 = v430
	v438 = v431
	goto L66
L95:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+1)))
	v429 = v425<<(uint(int32(8))%32) + v422
	v430 = v423
	v431 = v424
	goto L94
L96:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+2)))
	v422 = v418<<(uint(int32(16))%32) + v415
	v423 = v416
	v424 = v417
	goto L95
L97:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+3)))
	v415 = v411<<(uint(int32(24))%32) + v354
	v416 = v409
	v417 = v410
	goto L96
L98:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+4)))
	v409 = v405 + v407
	v410 = v406
	goto L97
L99:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+5)))
	v405 = v401<<(uint(int32(8))%32) + v399
	v406 = v400
	goto L98
L100:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+6)))
	v399 = v395<<(uint(int32(16))%32) + v393
	v400 = v394
	goto L99
L101:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+7)))
	v393 = v389<<(uint(int32(24))%32) + v355
	v394 = v388
	goto L100
L102:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+8)))
	v388 = v384<<(uint(int32(8))%32) + v383
	goto L101
L103:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+9)))
	v383 = v379<<(uint(int32(16))%32) + v378
	goto L102
L104:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+10)))
	v378 = v374<<(uint(int32(24))%32) + v359
	goto L103
L105:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v482 = v470
	v483 = v473
	goto L106
L106:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	if v491 != v479 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	v513 = v504
	goto L2
L108:
	;
	goto L107
L109:
	;
	v499 = (v482 + int32(1)) & v469
	v502 = v468 + v499*int32(20)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	if v503 != 0 {
		v482 = v499
		v483 = v502
		goto L106
	} else {
		goto L113
	}
L110:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v483)+8))
	if v493 != v478 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v483)+12))
	if v495 == v477 {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	v513 = v6
	goto L2
L114:
	;
	F_errmsg_internal(m, int32(_a_F_ec_search_derived_clause_for_ems_0), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L9
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_ec_search_derived_clause_for_ems_1), int32(327), int32(_a_F_ec_search_derived_clause_for_ems_2))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L9
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_element_hash(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_element_hash[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_FunctionCall1Coll(m, v5, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_elements_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v4 == int32(1) {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v7 != int32(1) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v17
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
			if v10 != int32(1) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v17
				return int32(0)
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v13)
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_elements_worker_jsonb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
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
	var v25 int32
	_ = v25
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_MemoryContextDelete(m, v35)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L54
	}
L2:
	;
	v110 = v106
	v112 = v104
	v114 = v105
	goto L33
L3:
	;
	v104 = v3
	v105 = v3
	v106 = int32(2)
	goto L2
L4:
	;
	v104 = v101
	v105 = v50
	v106 = int32(1)
	goto L2
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L29
	}
L6:
	;
	return
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v16&int32(268435456) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v16&int32(1073741824) == int32(0) {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L25
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0]))
	v35 = F_AllocSetContextCreateInternal(m, v30, int32(_a_F_elements_worker_jsonb_0), int32(0), int32(_a_F_elements_worker_jsonb_1), int32(_a_F_elements_worker_jsonb_2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v39 = F_JsonbIteratorInit(m, v14+int32(4))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v39
	v47 = F_JsonbIteratorNext(m, v11+int32(28), v11+int32(8), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L16
	}
L15:
	;
	v49 = int32(_a_F_elements_worker_jsonb_3)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0])) = v35
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)) = uint8(v53)
	if l1 == v53 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	switch v47 {
	case 0:
		goto L1
	default:
		goto L3
	case 3:
		goto L15
	}
L17:
	;
	v59 = F_JsonbValueToJsonb(m, v11+int32(8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v61 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v101 = v59
	goto L4
L21:
	;
	v64 = F_JsonbValueAsText(m, v11+int32(8))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)) = uint8(v66)
	v104 = v3
	v105 = v50
	v106 = int32(0)
	goto L2
L24:
	;
	v101 = v64
	goto L4
L25:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_elements_worker_jsonb_4), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_elements_worker_jsonb_5), int32(2235), int32(_a_F_elements_worker_jsonb_6))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_elements_worker_jsonb_7), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_elements_worker_jsonb_5), int32(2239), int32(_a_F_elements_worker_jsonb_6))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
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
	switch v110 {
	case 0:
		goto L39
	case 1:
		goto L38
	default:
		goto L37
	}
L35:
	;
	v110 = int32(0)
	v114 = v146
	goto L33
L36:
	;
	v110 = int32(1)
	v112 = v172
	v114 = v171
	goto L33
L37:
	;
	goto L42
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v112
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	F_tuplestore_putvalues(m, v117, v118, v11+int32(4), v11+int32(3))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L40
	}
L39:
	;
	v171 = v114
	v172 = int32(0)
	goto L36
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0])) = v114
	F_MemoryContextReset(m, v35)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v110 = int32(2)
	goto L33
L42:
	;
	v143 = F_JsonbIteratorNext(m, v11+int32(28), v11+int32(8), int32(1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	v145 = int32(_a_F_elements_worker_jsonb_3)
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0])) = v35
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)) = uint8(v149)
	if l1 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L43
L45:
	;
	switch v143 {
	case 0:
		goto L1
	default:
		goto L42
	case 3:
		goto L44
	}
L46:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v151 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v162 = F_JsonbValueToJsonb(m, v11+int32(8))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L53
	}
L49:
	;
	v154 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)) = uint8(v154)
	goto L35
L50:
	;
	goto L51
L51:
	;
	v158 = F_JsonbValueAsText(m, v11+int32(8))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	v171 = v146
	v172 = v158
	goto L36
L53:
	;
	v171 = v146
	v172 = v162
	goto L36
L54:
	;
	v185 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v185)
	m.G0 = v11 + int32(32)
	return
}
func F_encrypt_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
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
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	F_init_work(m, v9+int32(-12), l1, l4, v9+int32(-56))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(0)
	if l1 == v21 {
		v81 = l2
		v82 = v21
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v85 = int32(1)
	v86 = v81 + v85
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v91 = v89 & v85
	if v91 != 0 {
		goto L33
	} else {
		goto L34
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	if v25 == int32(0) {
		v81 = l2
		v82 = v21
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_encrypt_internal[1]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	goto L6
L6:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v31 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v61 = int32(1)
	if v31&v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v37 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v48 = int32(1)
	if v31&v48 != 0 {
		v60 = int32(base.Ui32(v31)>>(uint(v48)%32)) - v48
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v40 = int32(16)
	goto L13
L12:
	;
	v40 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = int32(4)
	goto L16
L15:
	;
	v47 = v40
	goto L16
L16:
	;
	v60 = v47
	goto L7
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v65 = v61
	goto L20
L19:
	;
	v65 = int32(4)
	goto L20
L20:
	;
	v66 = v65 + l2
	v68 = F_pg_do_encoding_conversion(m, v66, v60, v30, int32(6))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v68 != v66 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = F_cstring_to_text(m, v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v75 = l2
	goto L24
L24:
	;
	v76 = base.B2i32(l2 == v75)
	if l2 == v75 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	F_pfree(m, v68)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v75 = v71
	goto L24
L27:
	;
	v77 = l2
	goto L29
L28:
	;
	v77 = v75
	goto L29
L29:
	;
	if l2 == v75 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v79 = int32(0)
	goto L32
L31:
	;
	v79 = v75
	goto L32
L32:
	;
	v81 = v77
	v82 = v79
	goto L3
L33:
	;
	v92 = v86
	goto L35
L34:
	;
	v92 = v81 + int32(4)
	goto L35
L35:
	;
	if v89 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v120 = F_mbuf_create_from_data(m, v92, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L47
	}
L37:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v98 == int32(18) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v109 = int32(1)
	if v91 != 0 {
		v119 = int32(base.Ui32(v89)>>(uint(v109)%32)) - v109
		goto L36
	} else {
		goto L46
	}
L40:
	;
	v101 = int32(16)
	goto L42
L41:
	;
	v101 = int32(0)
	goto L42
L42:
	;
	if base.Ui32((v98-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v108 = int32(4)
	goto L45
L44:
	;
	v108 = v101
	goto L45
L45:
	;
	v119 = v108
	goto L36
L46:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v119 = int32(base.Ui32(v113)>>(uint(int32(2))%32)) - int32(4)
	goto L36
L47:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v122 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v150 = F_mbuf_create(m, v147+int32(128))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L59
	}
L49:
	;
	v126 = int32(18)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v128 == v126 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v139 = int32(1)
	if v122&v139 != 0 {
		v147 = int32(base.Ui32(v122) >> (uint(v139) % 32))
		goto L48
	} else {
		goto L58
	}
L52:
	;
	v131 = v126
	goto L54
L53:
	;
	v131 = int32(2)
	goto L54
L54:
	;
	if base.Ui32((v128-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v138 = int32(6)
	goto L57
L56:
	;
	v138 = v131
	goto L57
L57:
	;
	v147 = v138
	goto L48
L58:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v147 = int32(base.Ui32(v143) >> (uint(int32(2)) % 32))
	goto L48
L59:
	;
	v155 = F_mbuf_append(m, v150, v9+int32(-4), int32(4))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if l0 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if int32(0) <= v253 {
		goto L101
	} else {
		goto L102
	}
L62:
	;
	v157 = int32(1)
	v158 = l3 + v157
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v163 = v161 & v157
	if v163 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v203 = int32(1)
	v204 = l3 + v203
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v209 = v207 & v203
	if v209 != 0 {
		goto L82
	} else {
		goto L83
	}
L65:
	;
	v164 = v158
	goto L67
L66:
	;
	v164 = l3 + int32(4)
	goto L67
L67:
	;
	if v161 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v192 = F_mbuf_create_from_data(m, v164, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L79
	}
L69:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v170 == int32(18) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v181 = int32(1)
	if v163 != 0 {
		v191 = int32(base.Ui32(v161)>>(uint(v181)%32)) - v181
		goto L68
	} else {
		goto L78
	}
L72:
	;
	v173 = int32(16)
	goto L74
L73:
	;
	v173 = int32(0)
	goto L74
L74:
	;
	if base.Ui32((v170-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v180 = int32(4)
	goto L77
L76:
	;
	v180 = v173
	goto L77
L77:
	;
	v191 = v180
	goto L68
L78:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v191 = int32(base.Ui32(v185)>>(uint(int32(2))%32)) - int32(4)
	goto L68
L79:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v195 = int32(0)
	v198 = F_pgp_set_pubkey(m, v194, v192, v195, v195, v195)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v200 = F_mbuf_free(m, v192)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v253 = v198
	goto L61
L82:
	;
	v210 = v204
	goto L84
L83:
	;
	v210 = l3 + int32(4)
	goto L84
L84:
	;
	if v207 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v238 = int32(0)
	if base.B2i32(v210 == v238)|base.B2i32(v237 <= v238) != 0 {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v216 == int32(18) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v227 = int32(1)
	if v209 != 0 {
		v237 = int32(base.Ui32(v207)>>(uint(v227)%32)) - v227
		goto L85
	} else {
		goto L95
	}
L89:
	;
	v219 = int32(16)
	goto L91
L90:
	;
	v219 = int32(0)
	goto L91
L91:
	;
	if base.Ui32((v216-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v226 = int32(4)
	goto L94
L93:
	;
	v226 = v219
	goto L94
L94:
	;
	v237 = v226
	goto L85
L95:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v237 = int32(base.Ui32(v231)>>(uint(int32(2))%32)) - int32(4)
	goto L85
L96:
	;
	v253 = v249
	goto L61
L97:
	;
	v249 = int32(-13)
	goto L99
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+128)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v202)+124)) = v210
	v249 = int32(0)
	goto L99
L99:
	;
	goto L96
L100:
	;
	v309 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v150)+16)) = uint16(v309)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(-8)))) = v312
	v314 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v150)+8)) = v314
	*(*int64)(unsafe.Add(mBase, uint32(v150))) = v314
	goto L133
L101:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v257 = F_pgp_encrypt(m, v256, v120, v150)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	v261 = v253
	goto L103
L103:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v262 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	if v257 == int32(0) {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v261 = v257
	goto L103
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_encrypt_internal[0])) = int32(0)
	goto L109
L107:
	;
	goto L108
L108:
	;
	if v82 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v267 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	goto L112
L112:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v298 = F_pgp_free(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L129
	}
L113:
	;
	if v292 != 0 {
		goto L125
	} else {
		goto L126
	}
L114:
	;
	v271 = int32(18)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v273 == v271 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	v284 = int32(1)
	if v267&v284 != 0 {
		v292 = int32(base.Ui32(v267) >> (uint(v284) % 32))
		goto L113
	} else {
		goto L123
	}
L117:
	;
	v276 = v271
	goto L119
L118:
	;
	v276 = int32(2)
	goto L119
L119:
	;
	if base.Ui32((v273-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v283 = int32(6)
	goto L122
L121:
	;
	v283 = v276
	goto L122
L122:
	;
	v292 = v283
	goto L113
L123:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v292 = int32(base.Ui32(v288) >> (uint(int32(2)) % 32))
	goto L113
L124:
	;
	F_pfree(m, v82)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L128
	}
L125:
	;
	base.MemoryFill(m, v82, int32(0), v292)
	goto L127
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	goto L112
L129:
	;
	v300 = F_mbuf_free(m, v120)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v302 = F_mbuf_free(m, v150)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_px_THROW_ERROR(m, v261)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
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
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = (v311 - v312) << (uint(int32(2)) % 32)
	if v82 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v324 == int32(1) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	goto L136
L136:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v355 = F_pgp_free(m, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L153
	}
L137:
	;
	if v349 != 0 {
		goto L149
	} else {
		goto L150
	}
L138:
	;
	v328 = int32(18)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v330 == v328 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	v341 = int32(1)
	if v324&v341 != 0 {
		v349 = int32(base.Ui32(v324) >> (uint(v341) % 32))
		goto L137
	} else {
		goto L147
	}
L141:
	;
	v333 = v328
	goto L143
L142:
	;
	v333 = int32(2)
	goto L143
L143:
	;
	if base.Ui32((v330-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v340 = int32(6)
	goto L146
L145:
	;
	v340 = v333
	goto L146
L146:
	;
	v349 = v340
	goto L137
L147:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v349 = int32(base.Ui32(v345) >> (uint(int32(2)) % 32))
	goto L137
L148:
	;
	F_pfree(m, v82)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L152
	}
L149:
	;
	base.MemoryFill(m, v82, int32(0), v349)
	goto L151
L150:
	;
	goto L151
L151:
	;
	goto L148
L152:
	;
	goto L136
L153:
	;
	v357 = F_mbuf_free(m, v120)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v359 = F_mbuf_free(m, v150)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_encrypt_internal[0])) = int32(0)
	goto L156
L156:
	;
	m.G0 = v11 - int32(-64)
	return v319
}
func F_encrypt_password(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
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
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int64
	_ = v916
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v4
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v19 != int32(109) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v15 + int32(112)
	return v881
L2:
	;
	v1054 = F_parse_scram_secret(m, v881, v15+int32(104), v15+int32(96), v15+int32(100), v15+int32(108), v15-int32(-64), v15+int32(32))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L33
	} else {
		goto L261
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L33
	} else {
		goto L256
	}
L4:
	;
	v890 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[0])))
	if v890 != int32(1) {
		goto L1
	} else {
		goto L224
	}
L5:
	;
	if v864 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(0)
	switch l0 {
	case 0:
		goto L38
	case 1:
		goto L39
	case 2:
		goto L37
	default:
		v881 = v4
		goto L4
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(0)
	v139 = F_pstrdup(m, l2)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L33
	} else {
		goto L36
	}
L8:
	;
	v131 = F_parse_scram_secret(m, l2, v15+int32(104), v15+int32(96), v15+int32(100), v15+int32(108), v15-int32(-64), v15+int32(32))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v22 != int32(100) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
	if v25 != int32(53) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v28 = F_strlen(m, l2)
	mBase = m.M
	if v28 != int32(35) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v32 = l2 + int32(3)
	v33 = int32(_a_F_encrypt_password_0)
	v37 = m.G0
	v39 = v37 - int32(32)
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+24)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v39)+8)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v40
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[1])))
	if v48 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v116 == int32(32) {
		goto L7
	} else {
		goto L32
	}
L14:
	;
	v116 = int32(0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[2])))
	if v52 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v56 = v32
	goto L20
L18:
	;
	goto L19
L19:
	;
	v66 = v33
	v67 = v48
	goto L23
L20:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v62 == v48 {
		v56 = v56 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v116 = v56 - v32
	goto L13
L22:
	;
	goto L21
L23:
	;
	v74 = v39 + int32(base.Ui32(v67)>>(uint(int32(3))%32))&int32(28)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v76 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75 | v76<<(uint(v67)%32)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v80 != 0 {
		v66 = v66 + v76
		v67 = v80
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v83 == int32(0) {
		v106 = v32
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v116 = v106 - v32
	goto L13
L27:
	;
	v87 = v32
	v88 = v83
	goto L28
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(base.Ui32(v88)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v96)>>(uint(v88)%32))&int32(1) == int32(0) {
		v106 = v87
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v106 = v104
	goto L26
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v104 = v87 + int32(1)
	if v102 != 0 {
		v87 = v104
		v88 = v102
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L8
L33:
	;
	return int32(0)
L34:
	;
	if v131 == int32(0) {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L7
L36:
	;
	v864 = v139
	goto L5
L37:
	;
	v180 = m.G0
	v182 = v180 - int32(48)
	m.G0 = v182
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = int32(0)
	v188 = F_pg_saslprep(m, l2, v182+int32(44))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L33
	} else {
		goto L50
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L33
	} else {
		goto L46
	}
L39:
	;
	v144 = F_palloc(m, int32(36))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v146 = F_strlen(m, l1)
	mBase = m.M
	v149 = F_pg_md5_encrypt(m, l2, l1, v146, v144, v15+int32(28))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L33
	} else {
		goto L41
	}
L41:
	;
	if v149 != 0 {
		v864 = v144
		goto L5
	} else {
		goto L42
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L33
	} else {
		goto L43
	}
L43:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v155
	F_errmsg_internal(m, int32(_a_F_encrypt_password_1), v15+int32(16))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_2), int32(141), int32(_a_F_encrypt_password_3))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L33
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errmsg_internal(m, int32(_a_F_encrypt_password_4), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L33
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_2), int32(149), int32(_a_F_encrypt_password_3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L33
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v864 = v397
	goto L5
L50:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v182)+44))
	v191 = int32(16)
	v192 = v182 + v191
	v194 = int32(0)
	v198 = m.G0
	v200 = v198 - v191
	m.G0 = v200
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v194
	v206 = F_open(m, int32(_a_F_encrypt_password_5), v194, v200)
	mBase = m.M
	if v206 != int32(-1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v239 != 0 {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	goto L56
L53:
	;
	v239 = v194
	goto L54
L54:
	;
	m.G0 = v200 + int32(16)
	goto L51
L55:
	;
	v234 = F_close(m, v206)
	mBase = m.M
	v239 = v232
	goto L54
L56:
	;
	v212 = v192
	v213 = v191
	goto L57
L57:
	;
	v218 = F_read(m, v206, v212, v213)
	mBase = m.M
	if v218 <= int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v232 = int32(1)
	goto L55
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_encrypt_password[3]))
	if v222 == int32(27) {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v227 = v213 - v218
	if v227 != 0 {
		v212 = v212 + v218
		v213 = v227
		goto L57
	} else {
		goto L63
	}
L62:
	;
	v232 = int32(0)
	goto L55
L63:
	;
	goto L58
L64:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_encrypt_password[4]))
	v246 = m.G0
	v248 = v246 - int32(176)
	m.G0 = v248
	if v188 != 0 {
		goto L72
	} else {
		goto L73
	}
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L33
	} else {
		goto L216
	}
L67:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v182)+44))
	if v839 != 0 {
		goto L212
	} else {
		goto L213
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(_a_F_encrypt_password_6)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L33
	} else {
		goto L209
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(_a_F_encrypt_password_7)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L33
	} else {
		goto L206
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(_a_F_encrypt_password_8)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L33
	} else {
		goto L203
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L33
	} else {
		goto L200
	}
L72:
	;
	v250 = l2
	goto L74
L73:
	;
	v250 = v190
	goto L74
L74:
	;
	v257 = v182 + int32(12)
	v258 = F_scram_SaltedPassword(m, v250, int32(3), int32(32), v192, int32(16), v245, v248+int32(144), v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L33
	} else {
		goto L75
	}
L75:
	;
	if v258 < int32(0) {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v263 = F_pg_hmac_create(m, int32(3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L33
	} else {
		goto L78
	}
L77:
	;
	if v351 < int32(0) {
		goto L71
	} else {
		goto L124
	}
L78:
	;
	if v263 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	goto L83
L80:
	;
	goto L81
L81:
	;
	v295 = F_pg_hmac_init(m, v263, v248+int32(144), int32(32))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L33
	} else {
		goto L97
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(_a_F_encrypt_password_9)
	v351 = int32(-1)
	goto L77
L83:
	;
	goto L82
L95:
	;
	F_pg_hmac_free(m, v263)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L33
	} else {
		goto L123
	}
L96:
	;
	if v263 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L97:
	;
	if v295 < int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	if v263 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v315 < int32(0) {
		goto L96
	} else {
		goto L106
	}
L100:
	;
	v315 = int32(-1)
	goto L99
L101:
	;
	goto L102
L102:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v305 = F_pg_cryptohash_update(m, v304, int32(_a_F_encrypt_password_10), int32(10))
	mBase = m.M
	if int32(0) <= v305 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v315 = int32(0)
	goto L99
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+8)) = int32(2)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v312 = F_pg_cryptohash_error(m, v311)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v263)+12)) = v312
	v315 = int32(-1)
	goto L99
L106:
	;
	v319 = F_pg_hmac_final(m, v263, v248+int32(112), int32(32))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L33
	} else {
		goto L107
	}
L107:
	;
	if int32(0) <= v319 {
		goto L95
	} else {
		goto L108
	}
L108:
	;
	goto L96
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v342
	F_pg_hmac_free(m, v263)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L33
	} else {
		goto L122
	}
L110:
	;
	v342 = int32(_a_F_encrypt_password_9)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	if v327 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v339 = v327
	goto L115
L114:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	if v331 == int32(2) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v342 = v339
	goto L109
L116:
	;
	v334 = int32(_a_F_encrypt_password_11)
	goto L118
L117:
	;
	v334 = int32(_a_F_encrypt_password_12)
	goto L118
L118:
	;
	if v331 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v337 = int32(_a_F_encrypt_password_9)
	goto L121
L120:
	;
	v337 = v334
	goto L121
L121:
	;
	v339 = v337
	goto L115
L122:
	;
	v351 = int32(-1)
	goto L77
L123:
	;
	v351 = int32(0)
	goto L77
L124:
	;
	v355 = v248 + int32(112)
	v358 = F_scram_H(m, v355, int32(3), int32(32), v355, v257)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L33
	} else {
		goto L125
	}
L125:
	;
	if v358 < int32(0) {
		goto L71
	} else {
		goto L126
	}
L126:
	;
	v367 = v248 + int32(80)
	v368 = F_scram_ServerKey(m, v248+int32(144), int32(3), int32(32), v367, v257)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L33
	} else {
		goto L127
	}
L127:
	;
	if v368 < int32(0) {
		goto L71
	} else {
		goto L128
	}
L128:
	;
	v376 = base.I32_div_s(int32(18), int32(3))
	v378 = v376 << (uint(int32(2)) % 32)
	goto L129
L129:
	;
	v383 = base.I32_div_s(int32(34), int32(3))
	v385 = v383 << (uint(int32(2)) % 32)
	goto L130
L130:
	;
	v391 = base.I32_div_s(int32(34), int32(3))
	v393 = v391 << (uint(int32(2)) % 32)
	goto L131
L131:
	;
	v397 = F_palloc(m, v378+v385+v393+int32(28))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L33
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+64)) = v245
	v404 = F_pg_sprintf(m, v397, int32(_a_F_encrypt_password_13), v248-int32(-64))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L33
	} else {
		goto L133
	}
L133:
	;
	v406 = v404 + v397
	goto L137
L134:
	;
	if v518 < int32(0) {
		goto L70
	} else {
		goto L155
	}
L135:
	;
	if v378 != 0 {
		goto L152
	} else {
		goto L153
	}
L136:
	;
	if v378 < v460-v406+int32(4) {
		goto L135
	} else {
		goto L148
	}
L137:
	;
	v415 = v192
	v416 = int32(0)
	v419 = v406
	v420 = int32(2)
	goto L140
L139:
	;
	v518 = v460 - v406
	goto L134
L140:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	v426 = v422<<(uint(v420<<(uint(int32(3))%32))%32) | v416
	if int32(0) < v420 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	if v461 != int32(2) {
		goto L136
	} else {
		goto L147
	}
L142:
	;
	v459 = v426
	v460 = v419
	v461 = v420 - int32(1)
	goto L144
L143:
	;
	if v378 < v419-v406+int32(4) {
		goto L135
	} else {
		goto L145
	}
L144:
	;
	v463 = v415 + int32(1)
	if base.Ui32(v463) < base.Ui32(v182+int32(32)) {
		v415 = v463
		v416 = v459
		v419 = v460
		v420 = v461
		goto L140
	} else {
		goto L146
	}
L145:
	;
	v435 = int32(63)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426&v435)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+3)) = uint8(v437)
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v426)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v419))) = uint8(v441)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v426)>>(uint(int32(6))%32))&v435)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+2)) = uint8(v447)
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v426)>>(uint(int32(12))%32))&v435)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)) = uint8(v453)
	v459 = int32(0)
	v460 = v419 + int32(4)
	v461 = int32(2)
	goto L144
L146:
	;
	goto L141
L147:
	;
	goto L139
L148:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v459)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v481)
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v459)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v460)+1)) = uint8(v487)
	if v461 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v459)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	v497 = v496
	goto L151
L150:
	;
	v497 = int32(61)
	goto L151
L151:
	;
	v498 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v460)+3)) = uint8(v498)
	*(*uint8)(unsafe.Add(mBase, uint32(v460)+2)) = uint8(v497)
	v518 = v460 + int32(4) - v406
	goto L134
L152:
	;
	base.MemoryFill(m, v406, int32(0), v378)
	goto L154
L153:
	;
	goto L154
L154:
	;
	v518 = int32(-1)
	goto L134
L155:
	;
	v521 = v406 + v518
	v522 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v521))) = uint8(v522)
	v526 = v521 + int32(1)
	goto L159
L156:
	;
	if v638 < int32(0) {
		goto L69
	} else {
		goto L177
	}
L157:
	;
	if v385 != 0 {
		goto L174
	} else {
		goto L175
	}
L158:
	;
	if v385 < v580-v526+int32(4) {
		goto L157
	} else {
		goto L170
	}
L159:
	;
	v535 = v355
	v536 = int32(0)
	v539 = v526
	v540 = int32(2)
	goto L162
L161:
	;
	v638 = v580 - v526
	goto L156
L162:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	v546 = v542<<(uint(v540<<(uint(int32(3))%32))%32) | v536
	if int32(0) < v540 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v581 != int32(2) {
		goto L158
	} else {
		goto L169
	}
L164:
	;
	v579 = v546
	v580 = v539
	v581 = v540 - int32(1)
	goto L166
L165:
	;
	if v385 < v539-v526+int32(4) {
		goto L157
	} else {
		goto L167
	}
L166:
	;
	v583 = v535 + int32(1)
	if base.Ui32(v583) < base.Ui32(v248+int32(144)) {
		v535 = v583
		v536 = v579
		v539 = v580
		v540 = v581
		goto L162
	} else {
		goto L168
	}
L167:
	;
	v555 = int32(63)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546&v555)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+3)) = uint8(v557)
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v546)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v539))) = uint8(v561)
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v546)>>(uint(int32(6))%32))&v555)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+2)) = uint8(v567)
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v546)>>(uint(int32(12))%32))&v555)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+1)) = uint8(v573)
	v579 = int32(0)
	v580 = v539 + int32(4)
	v581 = int32(2)
	goto L166
L168:
	;
	goto L163
L169:
	;
	goto L161
L170:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v579)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v580))) = uint8(v601)
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v579)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+1)) = uint8(v607)
	if v581 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v579)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	v617 = v616
	goto L173
L172:
	;
	v617 = int32(61)
	goto L173
L173:
	;
	v618 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+3)) = uint8(v618)
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+2)) = uint8(v617)
	v638 = v580 + int32(4) - v526
	goto L156
L174:
	;
	base.MemoryFill(m, v526, int32(0), v385)
	goto L176
L175:
	;
	goto L176
L176:
	;
	v638 = int32(-1)
	goto L156
L177:
	;
	v641 = v526 + v638
	v642 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v641))) = uint8(v642)
	v646 = v641 + int32(1)
	goto L181
L178:
	;
	if v758 < int32(0) {
		goto L68
	} else {
		goto L199
	}
L179:
	;
	if v393 != 0 {
		goto L196
	} else {
		goto L197
	}
L180:
	;
	if v393 < v700-v646+int32(4) {
		goto L179
	} else {
		goto L192
	}
L181:
	;
	v655 = v367
	v656 = int32(0)
	v659 = v646
	v660 = int32(2)
	goto L184
L183:
	;
	v758 = v700 - v646
	goto L178
L184:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655))))
	v666 = v662<<(uint(v660<<(uint(int32(3))%32))%32) | v656
	if int32(0) < v660 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	if v701 != int32(2) {
		goto L180
	} else {
		goto L191
	}
L186:
	;
	v699 = v666
	v700 = v659
	v701 = v660 - int32(1)
	goto L188
L187:
	;
	if v393 < v659-v646+int32(4) {
		goto L179
	} else {
		goto L189
	}
L188:
	;
	v703 = v655 + int32(1)
	if base.Ui32(v703) < base.Ui32(v248+int32(112)) {
		v655 = v703
		v656 = v699
		v659 = v700
		v660 = v701
		goto L184
	} else {
		goto L190
	}
L189:
	;
	v675 = int32(63)
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666&v675)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+3)) = uint8(v677)
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v666)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v659))) = uint8(v681)
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v666)>>(uint(int32(6))%32))&v675)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+2)) = uint8(v687)
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v666)>>(uint(int32(12))%32))&v675)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+1)) = uint8(v693)
	v699 = int32(0)
	v700 = v659 + int32(4)
	v701 = int32(2)
	goto L188
L190:
	;
	goto L185
L191:
	;
	goto L183
L192:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v699)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v700))) = uint8(v721)
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v699)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v700)+1)) = uint8(v727)
	if v701 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v699)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	v737 = v736
	goto L195
L194:
	;
	v737 = int32(61)
	goto L195
L195:
	;
	v738 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v700)+3)) = uint8(v738)
	*(*uint8)(unsafe.Add(mBase, uint32(v700)+2)) = uint8(v737)
	v758 = v700 + int32(4) - v646
	goto L178
L196:
	;
	base.MemoryFill(m, v646, int32(0), v393)
	goto L198
L197:
	;
	goto L198
L198:
	;
	v758 = int32(-1)
	goto L178
L199:
	;
	v762 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v646+v758))) = uint8(v762)
	m.G0 = v248 + int32(176)
	goto L67
L200:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v775
	F_errmsg_internal(m, int32(_a_F_encrypt_password_14), v248)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L33
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_15), int32(245), int32(_a_F_encrypt_password_16))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L33
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+16)) = v791
	F_errmsg_internal(m, int32(_a_F_encrypt_password_17), v248+int32(16))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L33
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_15), int32(286), int32(_a_F_encrypt_password_16))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L33
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+32)) = v809
	F_errmsg_internal(m, int32(_a_F_encrypt_password_17), v248+int32(32))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L33
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_15), int32(302), int32(_a_F_encrypt_password_16))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L33
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+48)) = v827
	F_errmsg_internal(m, int32(_a_F_encrypt_password_17), v248+int32(48))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L33
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_15), int32(319), int32(_a_F_encrypt_password_16))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L33
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	F_pfree(m, v839)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L33
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	m.G0 = v182 + int32(48)
	goto L49
L215:
	;
	goto L214
L216:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L33
	} else {
		goto L217
	}
L217:
	;
	F_errmsg(m, int32(_a_F_encrypt_password_18), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L33
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_19), int32(504), int32(_a_F_encrypt_password_20))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L33
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	v881 = int32(0)
	goto L4
L221:
	;
	goto L222
L222:
	;
	v875 = F_strlen(m, v864)
	mBase = m.M
	if base.Ui32(int32(513)) <= base.Ui32(v875) {
		goto L3
	} else {
		goto L223
	}
L223:
	;
	v881 = v864
	goto L4
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = int32(0)
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881))))
	if v895 != int32(109) {
		goto L2
	} else {
		goto L225
	}
L225:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881)+1)))
	if v898 != int32(100) {
		goto L2
	} else {
		goto L226
	}
L226:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881)+2)))
	if v901 != int32(53) {
		goto L2
	} else {
		goto L227
	}
L227:
	;
	v904 = F_strlen(m, v881)
	mBase = m.M
	if v904 != int32(35) {
		goto L2
	} else {
		goto L228
	}
L228:
	;
	v908 = v881 + int32(3)
	v909 = int32(_a_F_encrypt_password_0)
	v913 = m.G0
	v915 = v913 - int32(32)
	v916 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v915)+24)) = v916
	*(*int64)(unsafe.Add(mBase, uint32(v915)+16)) = v916
	*(*int64)(unsafe.Add(mBase, uint32(v915)+8)) = v916
	*(*int64)(unsafe.Add(mBase, uint32(v915))) = v916
	v924 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[1])))
	if v924 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	if v992 != int32(32) {
		goto L2
	} else {
		goto L248
	}
L230:
	;
	v992 = int32(0)
	goto L229
L231:
	;
	goto L232
L232:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[2])))
	if v928 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v932 = v908
	goto L236
L234:
	;
	goto L235
L235:
	;
	v942 = v909
	v943 = v924
	goto L239
L236:
	;
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932))))
	if v938 == v924 {
		v932 = v932 + int32(1)
		goto L236
	} else {
		goto L238
	}
L237:
	;
	v992 = v932 - v908
	goto L229
L238:
	;
	goto L237
L239:
	;
	v950 = v915 + int32(base.Ui32(v943)>>(uint(int32(3))%32))&int32(28)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v950)))
	v952 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v950))) = v951 | v952<<(uint(v943)%32)
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942)+1)))
	if v956 != 0 {
		v942 = v942 + v952
		v943 = v956
		goto L239
	} else {
		goto L241
	}
L240:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	if v959 == int32(0) {
		v982 = v908
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L240
L242:
	;
	v992 = v982 - v908
	goto L229
L243:
	;
	v963 = v908
	v964 = v959
	goto L244
L244:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v915+int32(base.Ui32(v964)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v972)>>(uint(v964)%32))&int32(1) == int32(0) {
		v982 = v963
		goto L242
	} else {
		goto L246
	}
L245:
	;
	v982 = v980
	goto L242
L246:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963)+1)))
	v980 = v963 + int32(1)
	if v978 != 0 {
		v963 = v980
		v964 = v978
		goto L244
	} else {
		goto L247
	}
L247:
	;
	goto L245
L248:
	;
	v997 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L33
	} else {
		goto L249
	}
L249:
	;
	if v997 == int32(0) {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_errcode(m, int32(16908352))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L33
	} else {
		goto L251
	}
L251:
	;
	F_errmsg(m, int32(_a_F_encrypt_password_21), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L33
	} else {
		goto L252
	}
L252:
	;
	F_errdetail(m, int32(_a_F_encrypt_password_22), int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L33
	} else {
		goto L253
	}
L253:
	;
	F_errhint(m, int32(_a_F_encrypt_password_23), int32(0))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L33
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_2), int32(185), int32(_a_F_encrypt_password_3))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L33
	} else {
		goto L255
	}
L255:
	;
	goto L1
L256:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L33
	} else {
		goto L257
	}
L257:
	;
	F_errmsg(m, int32(_a_F_encrypt_password_24), int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L33
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(512)
	F_errdetail(m, int32(_a_F_encrypt_password_25), v15)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L33
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_2), int32(176), int32(_a_F_encrypt_password_3))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L33
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	goto L1
}
func F_encrypt_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	if l3 <= int32(0) {
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
	v13 = l1 + int32(4)
	v16 = l2
	v17 = l3
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = int32(_a_F_encrypt_process_0)
	if base.Ui32(v22) <= base.Ui32(v17) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return v43
L6:
	;
	goto L5
L7:
	;
	v25 = v22
	goto L9
L8:
	;
	v25 = v17
	goto L9
L9:
	;
	v26 = F_pgp_cfb_encrypt(m, v21, v16, v25, v13)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v26 < int32(0) {
		v43 = v26
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v32 = F_pushf_write(m, l0, v13, v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v32 < int32(0) {
		v43 = v32
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(0)
	v38 = v17 - v25
	if v37 < v38 {
		v16 = v16 + v25
		v17 = v38
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v43 = v37
	goto L6
}
func F_end_MultiFuncCall(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_UnregisterExprContextCallback(m, v4, int32(1613), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
		F_MemoryContextDelete(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
		}
	}
}
func F_end_tup_output(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	m.T0[v4].(func(*base.Module, int32))(m, v3)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_ExecDropSingleTupleTableSlot(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_english_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v765 int32
	_ = v765
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v984 int32
	_ = v984
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1209 int32
	_ = v1209
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1425 int32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1445 int32
	_ = v1445
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1551 int32
	_ = v1551
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	v9 = v6 + int32(2)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 <= v9 {
		v106 = v10
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v1551
L2:
	;
	v1551 = int32(1)
	goto L1
L3:
	;
	v110 = v6 + int32(3)
	v111 = base.B2i32(v106 < v110)
	if v106 < v110 {
		goto L43
	} else {
		goto L44
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v9))))
	if base.B2i32(v14&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v14)%32)&int32(42750482) == int32(0)) != 0 {
		v106 = v10
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = F_find_among(m, l0, int32(_a_F_english_ISO_8859_1_stem_0), int32(18))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v28 == int32(0) {
		v106 = v32
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v35
	if v35 < v32 {
		v106 = v32
		goto L3
	} else {
		goto L9
	}
L9:
	;
	switch v28 - int32(1) {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	case 5:
		goto L15
	case 6:
		goto L14
	case 7:
		goto L13
	case 8:
		goto L12
	case 9:
		goto L11
	case 10:
		goto L10
	default:
		goto L2
	}
L10:
	;
	v102 = F_slice_from_s(m, l0, int32(5), int32(_a_F_english_ISO_8859_1_stem_1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L41
	}
L11:
	;
	v96 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L39
	}
L12:
	;
	v90 = F_slice_from_s(m, l0, int32(5), int32(_a_F_english_ISO_8859_1_stem_3))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L37
	}
L13:
	;
	v84 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_4))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L35
	}
L14:
	;
	v78 = F_slice_from_s(m, l0, int32(5), int32(_a_F_english_ISO_8859_1_stem_5))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L33
	}
L15:
	;
	v72 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_6))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L31
	}
L16:
	;
	v66 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_7))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L29
	}
L17:
	;
	v60 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_8))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L27
	}
L18:
	;
	v54 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_9))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L25
	}
L19:
	;
	v48 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_10))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L23
	}
L20:
	;
	v42 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_11))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	if int32(0) <= v42 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v1551 = v42
	goto L1
L23:
	;
	if int32(0) <= v48 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v1551 = v48
	goto L1
L25:
	;
	if int32(0) <= v54 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v1551 = v54
	goto L1
L27:
	;
	if int32(0) <= v60 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v1551 = v60
	goto L1
L29:
	;
	if int32(0) <= v66 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v1551 = v66
	goto L1
L31:
	;
	if int32(0) <= v72 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v1551 = v72
	goto L1
L33:
	;
	if int32(0) <= v78 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v1551 = v78
	goto L1
L35:
	;
	if int32(0) <= v84 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v1551 = v84
	goto L1
L37:
	;
	if int32(0) <= v90 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v1551 = v90
	goto L1
L39:
	;
	if int32(0) <= v96 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v1551 = v96
	goto L1
L41:
	;
	if int32(0) <= v102 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v1551 = v102
	goto L1
L43:
	;
	v112 = v6
	goto L45
L44:
	;
	v112 = v110
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v112
	if v106 < v110 {
		v1551 = int32(1)
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v119 == v121 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119
	v166 = v119
	goto L58
L48:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v119))))
	if v125 == int32(39) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v129 = v119 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v129
	v132 = F_slice_del(m, l0)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	v144 = v125
	goto L51
L51:
	;
	if v144&int32(255) != int32(121) {
		goto L47
	} else {
		goto L55
	}
L52:
	;
	if v132 < int32(0) {
		v1551 = v132
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v119 == v138 {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+v119))))
	v144 = v142
	goto L51
L55:
	;
	v149 = int32(1)
	v150 = v119 + v149
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v150
	v155 = F_slice_from_s(m, l0, v149, int32(_a_F_english_ISO_8859_1_stem_12))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	if v155 < int32(0) {
		v1551 = v155
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+8)) = int32(1)
	goto L47
L58:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v178 < v177 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v222
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v256 = v254 + int32(4)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v257 <= v256 {
		goto L85
	} else {
		goto L86
	}
L60:
	;
	goto L59
L61:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v221 != 0 {
		goto L76
	} else {
		goto L77
	}
L62:
	;
	v180 = v177
	goto L64
L63:
	;
	v180 = v178
	goto L64
L64:
	;
	goto L66
L65:
	;
	v221 = v216
	goto L61
L66:
	;
	if v177 == v180 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v216 = int32(0)
	goto L65
L68:
	;
	v221 = int32(-1)
	goto L61
L69:
	;
	goto L70
L70:
	;
	v192 = int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v177))))
	if int32(121) < v195 {
		v216 = v192
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v197 = v195 - int32(97)
	if v197 < int32(0) {
		v216 = v192
		goto L65
	} else {
		goto L72
	}
L72:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v197)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v203)>>(uint(v197&int32(7))%32))&int32(1) == int32(0) {
		v216 = v192
		goto L65
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v177 + int32(1)
	goto L74
L74:
	;
	goto L67
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v166
	v237 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223 + v237
	v242 = F_slice_from_s(m, l0, v237, int32(_a_F_english_ISO_8859_1_stem_13))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L6
	} else {
		goto L81
	}
L76:
	;
	if v222 <= v166 {
		goto L60
	} else {
		goto L80
	}
L77:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v223
	if v222 == v223 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v223))))
	if v228 == int32(121) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	v234 = v166 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v234
	v166 = v234
	goto L58
L81:
	;
	if v242 < int32(0) {
		v1551 = v242
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = int32(1)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v166 = v249
	goto L58
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v254
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v512
	if v512 <= v254 {
		v542 = v512
		goto L152
	} else {
		goto L153
	}
L84:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+4)) = v394
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v405 < v404 {
		goto L122
	} else {
		goto L123
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v254
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v290 < v254 {
		goto L91
	} else {
		goto L92
	}
L86:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259+v256))))
	if base.B2i32(v261&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v261)%32)&int32(_a_F_english_ISO_8859_1_stem_14) == int32(0)) != 0 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v275 = F_find_among(m, l0, int32(_a_F_english_ISO_8859_1_stem_15), int32(3))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	if v275 == int32(0) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v394 = v279
	goto L84
L90:
	;
	if v330 < int32(0) {
		goto L83
	} else {
		goto L105
	}
L91:
	;
	v292 = v254
	goto L93
L92:
	;
	v292 = v290
	goto L93
L93:
	;
	v299 = v254
	goto L95
L94:
	;
	v330 = v310
	goto L90
L95:
	;
	if v299 == v292 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v330 = int32(-1)
	goto L90
L98:
	;
	goto L99
L99:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v299))))
	if int32(121) < v305 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v322 = v299 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v322
	v299 = v322
	goto L95
L101:
	;
	v307 = v305 - int32(97)
	if v307 < int32(0) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v310 = int32(1)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v307)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v314)>>(uint(v307&int32(7))%32))&v310 != 0 {
		goto L94
	} else {
		goto L103
	}
L103:
	;
	goto L100
L105:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v334 = v333 + v330
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v334
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v345 < v334 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v388 < int32(0) {
		goto L83
	} else {
		goto L120
	}
L107:
	;
	v347 = v334
	goto L109
L108:
	;
	v347 = v345
	goto L109
L109:
	;
	v353 = v334
	goto L111
L110:
	;
	v388 = int32(1)
	goto L106
L111:
	;
	if v353 == v347 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v388 = int32(-1)
	goto L106
L114:
	;
	goto L115
L115:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360+v353))))
	if int32(121) < v362 {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v364 = v362 - int32(97)
	if v364 < int32(0) {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v364)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v370)>>(uint(v364&int32(7))%32))&int32(1) == int32(0) {
		goto L110
	} else {
		goto L118
	}
L118:
	;
	v379 = v353 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v379
	v353 = v379
	goto L111
L120:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v392 = v391 + v388
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v392
	v394 = v392
	goto L84
L121:
	;
	if v445 < int32(0) {
		goto L83
	} else {
		goto L136
	}
L122:
	;
	v407 = v404
	goto L124
L123:
	;
	v407 = v405
	goto L124
L124:
	;
	v414 = v404
	goto L126
L125:
	;
	v445 = v425
	goto L121
L126:
	;
	if v414 == v407 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v445 = int32(-1)
	goto L121
L129:
	;
	goto L130
L130:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418+v414))))
	if int32(121) < v420 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v437 = v414 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v437
	v414 = v437
	goto L126
L132:
	;
	v422 = v420 - int32(97)
	if v422 < int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v425 = int32(1)
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v422)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v429)>>(uint(v422&int32(7))%32))&v425 != 0 {
		goto L125
	} else {
		goto L134
	}
L134:
	;
	goto L131
L136:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v449 = v448 + v445
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v449
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v460 < v449 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v503 < int32(0) {
		goto L83
	} else {
		goto L151
	}
L138:
	;
	v462 = v449
	goto L140
L139:
	;
	v462 = v460
	goto L140
L140:
	;
	v468 = v449
	goto L142
L141:
	;
	v503 = int32(1)
	goto L137
L142:
	;
	if v468 == v462 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v503 = int32(-1)
	goto L137
L145:
	;
	goto L146
L146:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475+v468))))
	if int32(121) < v477 {
		goto L141
	} else {
		goto L147
	}
L147:
	;
	v479 = v477 - int32(97)
	if v479 < int32(0) {
		goto L141
	} else {
		goto L148
	}
L148:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v479)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v485)>>(uint(v479&int32(7))%32))&int32(1) == int32(0) {
		goto L141
	} else {
		goto L149
	}
L149:
	;
	v494 = v468 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v494
	v468 = v494
	goto L142
L151:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v506))) = v507 + v503
	goto L83
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v542
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v542 <= v544 {
		goto L161
	} else {
		goto L162
	}
L153:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516+v512-int32(1)))))
	if base.B2i32(v520 != int32(115))&base.B2i32(v520 != int32(39)) != 0 {
		v542 = v512
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v528 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_16), int32(3))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L6
	} else {
		goto L155
	}
L155:
	;
	if v528 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v532
	v542 = v532
	goto L152
L157:
	;
	goto L158
L158:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v534
	v536 = F_slice_del(m, l0)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	if v536 < int32(0) {
		v1551 = v536
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v542 = v540
	goto L152
L161:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v659
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v659
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v659-int32(5) <= v662 {
		v684 = v662
		goto L195
	} else {
		goto L196
	}
L162:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v542-int32(1)))))
	v552 = v550 - int32(100)
	v553 = int32(0)
	if base.B2i32(v552 == v553)|base.B2i32(v552 == int32(15)) == v553 {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v562 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_17), int32(6))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	if v562 == int32(0) {
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v566
	switch v562 - int32(1) {
	case 0:
		goto L168
	case 1:
		goto L167
	case 2:
		goto L166
	default:
		goto L161
	}
L166:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v566 <= v594 {
		goto L161
	} else {
		goto L178
	}
L167:
	;
	v577 = v566 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v577
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v579 <= v577 {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v572 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_18))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	if int32(0) <= v572 {
		goto L161
	} else {
		goto L170
	}
L170:
	;
	v1551 = v572
	goto L1
L171:
	;
	v583 = F_slice_from_s(m, l0, int32(1), int32(_a_F_english_ISO_8859_1_stem_19))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L6
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v566
	v590 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_20))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L6
	} else {
		goto L176
	}
L174:
	;
	if int32(0) <= v583 {
		goto L161
	} else {
		goto L175
	}
L175:
	;
	v1551 = v583
	goto L1
L176:
	;
	if int32(0) <= v590 {
		goto L161
	} else {
		goto L177
	}
L177:
	;
	v1551 = v590
	goto L1
L178:
	;
	v597 = v566 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v597
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v613 = v597
	goto L181
L179:
	;
	if v647 < int32(0) {
		goto L161
	} else {
		goto L191
	}
L180:
	;
	v647 = v627
	goto L179
L181:
	;
	if v613 <= v607 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v647 = int32(-1)
	goto L179
L184:
	;
	goto L185
L185:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618+v613-int32(1)))))
	if int32(121) < v622 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v639 = v613 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v639
	v613 = v639
	goto L181
L187:
	;
	v624 = v622 - int32(97)
	if v624 < int32(0) {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v627 = int32(1)
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v624)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v631)>>(uint(v624&int32(7))%32))&v627 != 0 {
		goto L180
	} else {
		goto L189
	}
L189:
	;
	goto L186
L191:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v650 - v647
	v653 = F_slice_del(m, l0)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	if v653 < int32(0) {
		v1551 = v653
		goto L1
	} else {
		goto L193
	}
L193:
	;
	goto L161
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1496
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+8))
	if v1500 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L195:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v685
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v685
	v689 = v685 - int32(1)
	if v689 <= v684 {
		goto L201
	} else {
		goto L202
	}
L196:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666+v659-int32(1)))))
	switch v670 - int32(100) {
	case 0, 3:
		goto L197
	default:
		v684 = v662
		goto L195
	}
L197:
	;
	v675 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_21), int32(8))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v675 == int32(0) {
		v684 = v677
		goto L195
	} else {
		goto L199
	}
L199:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v680
	if v680 <= v677 {
		v1496 = v677
		goto L194
	} else {
		goto L200
	}
L200:
	;
	v684 = v677
	goto L195
L201:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v956
	v958 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v956
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v956 <= v961 {
		v1040 = v958
		goto L263
	} else {
		goto L264
	}
L202:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691+v689))))
	if base.B2i32(v693&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v693)%32)&int32(33554576) == int32(0)) != 0 {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v707 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_22), int32(6))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L6
	} else {
		goto L204
	}
L204:
	;
	if v707 == int32(0) {
		goto L201
	} else {
		goto L205
	}
L205:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v711
	switch v707 - int32(1) {
	case 0:
		goto L207
	case 1:
		goto L206
	default:
		goto L201
	}
L206:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v739 = v732
	goto L213
L207:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+4))
	if v711 < v716 {
		goto L201
	} else {
		goto L208
	}
L208:
	;
	v720 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_23))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L209
	}
L209:
	;
	if int32(0) <= v720 {
		goto L201
	} else {
		goto L210
	}
L210:
	;
	v1551 = v720
	goto L1
L211:
	;
	if v773 < int32(0) {
		goto L201
	} else {
		goto L223
	}
L212:
	;
	v773 = v753
	goto L211
L213:
	;
	if v739 <= v733 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v773 = int32(-1)
	goto L211
L216:
	;
	goto L217
L217:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744+v739-int32(1)))))
	if int32(121) < v748 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v765 = v739 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v765
	v739 = v765
	goto L213
L219:
	;
	v750 = v748 - int32(97)
	if v750 < int32(0) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v753 = int32(1)
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v750)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v757)>>(uint(v750&int32(7))%32))&v753 != 0 {
		goto L212
	} else {
		goto L221
	}
L221:
	;
	goto L218
L223:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v776 + (v711 - v724)
	v780 = F_slice_del(m, l0)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L6
	} else {
		goto L224
	}
L224:
	;
	if v780 < int32(0) {
		v1551 = v780
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v784
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v784
	v788 = v784 - int32(1)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v788 <= v789 {
		v871 = v784
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v941 = v939 + (v784 - v805)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v941
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v941
	if v941 <= v938 {
		goto L201
	} else {
		goto L260
	}
L227:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v938 = v937
	goto L226
L228:
	;
	v929 = F_slice_from_s(m, l0, int32(1), v925)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L6
	} else {
		goto L258
	}
L229:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	if v871 != v874 {
		goto L201
	} else {
		goto L248
	}
L230:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791+v788))))
	if base.B2i32(v793&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v793)%32)&int32(68514004) == int32(0)) != 0 {
		v871 = v784
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v809 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_24), int32(13))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L6
	} else {
		goto L234
	}
L232:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v871 = v869
	goto L229
L233:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L237
L234:
	;
	switch v809 - int32(1) {
	case 0:
		v925 = int32(_a_F_english_ISO_8859_1_stem_25)
		goto L228
	case 1:
		goto L233
	case 2:
		goto L232
	default:
		goto L227
	}
L235:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v865 != 0 {
		v938 = v866
		goto L226
	} else {
		goto L246
	}
L236:
	;
	v865 = v861
	goto L235
L237:
	;
	if v821 <= v822 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v861 = int32(0)
	goto L236
L239:
	;
	v865 = int32(-1)
	goto L235
L240:
	;
	goto L241
L241:
	;
	v834 = int32(1)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835+v821-v834))))
	if int32(111) < v839 {
		v861 = v834
		goto L236
	} else {
		goto L242
	}
L242:
	;
	v841 = v839 - int32(97)
	if v841 < int32(0) {
		v861 = v834
		goto L236
	} else {
		goto L243
	}
L243:
	;
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v841)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v847)>>(uint(v841&int32(7))%32))&int32(1) == int32(0) {
		v861 = v834
		goto L236
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v821 - int32(1)
	goto L245
L245:
	;
	goto L238
L246:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v866 < v867 {
		v938 = v866
		goto L226
	} else {
		goto L247
	}
L247:
	;
	goto L201
L248:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v877 = int32(0)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v886 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(89), int32(121), v877)
	mBase = m.M
	if v886 != 0 {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	if v917 == int32(0) {
		goto L201
	} else {
		goto L257
	}
L250:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v898 + (v881 - v876)
	v906 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_27), int32(97), int32(121), int32(0))
	mBase = m.M
	if v906 != 0 {
		v915 = v877
		goto L254
	} else {
		goto L255
	}
L251:
	;
	v891 = F_in_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_27), int32(97), int32(121), int32(0))
	mBase = m.M
	if v891 != 0 {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v896 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_27), int32(97), int32(121), int32(0))
	mBase = m.M
	if v896 != 0 {
		goto L250
	} else {
		goto L253
	}
L253:
	;
	v917 = int32(1)
	goto L249
L254:
	;
	v917 = v915
	goto L249
L255:
	;
	v911 = F_in_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_27), int32(97), int32(121), int32(0))
	mBase = m.M
	if v911 != 0 {
		v915 = v877
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v915 = base.B2i32(v912 <= v913)
	goto L254
L257:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v920 + (v871 - v876)
	v925 = int32(_a_F_english_ISO_8859_1_stem_28)
	goto L228
L258:
	;
	if int32(0) <= v929 {
		goto L201
	} else {
		goto L259
	}
L259:
	;
	return v929 >> (uint(int32(31)) % 32) & v929
L260:
	;
	v946 = v941 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v946
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v946
	v949 = F_slice_del(m, l0)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L6
	} else {
		goto L261
	}
L261:
	;
	if v949 < int32(0) {
		v1551 = v949
		goto L1
	} else {
		goto L262
	}
L262:
	;
	goto L201
L263:
	;
	if v1040 < int32(0) {
		v1551 = v1040
		goto L1
	} else {
		goto L284
	}
L264:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963+v956-int32(1)))))
	if v967|int32(32) != int32(121) {
		v1040 = v958
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v973 = v956 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v973
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v973
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L268
L266:
	;
	if v1024 != 0 {
		v1040 = v958
		goto L263
	} else {
		goto L278
	}
L267:
	;
	v1024 = v1021
	goto L266
L268:
	;
	if v973 <= v984 {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1021 = int32(0)
	goto L267
L270:
	;
	v1024 = int32(-1)
	goto L266
L271:
	;
	goto L272
L272:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995+v973-int32(1)))))
	if int32(121) < v999 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v973 - int32(1)
	goto L277
L274:
	;
	v1001 = v999 - int32(97)
	if v1001 < int32(0) {
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v1004 = int32(1)
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1001)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v1008)>>(uint(v1001&int32(7))%32))&v1004 != 0 {
		v1021 = v1004
		goto L267
	} else {
		goto L276
	}
L276:
	;
	goto L273
L277:
	;
	goto L269
L278:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1025 <= v1026 {
		v1040 = v958
		goto L263
	} else {
		goto L279
	}
L279:
	;
	v1028 = int32(1)
	v1031 = F_slice_from_s(m, l0, v1028, int32(_a_F_english_ISO_8859_1_stem_29))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L6
	} else {
		goto L280
	}
L280:
	;
	if int32(0) <= v1031 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1038 = v1028
	goto L283
L282:
	;
	v1038 = v1031 >> (uint(int32(31)) % 32) & v1031
	goto L283
L283:
	;
	v1040 = v1038
	goto L263
L284:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1043
	v1045 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1043
	v1049 = v1043 - int32(1)
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1049 <= v1050 {
		v1235 = v1045
		goto L285
	} else {
		goto L286
	}
L285:
	;
	if v1235 < int32(0) {
		v1551 = v1235
		goto L1
	} else {
		goto L351
	}
L286:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052+v1049))))
	if base.B2i32(v1054&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1054)%32)&int32(_a_F_english_ISO_8859_1_stem_30) == int32(0)) != 0 {
		v1235 = v1045
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v1068 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_31), int32(24))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L6
	} else {
		goto L288
	}
L288:
	;
	if v1068 == int32(0) {
		v1235 = v1045
		goto L285
	} else {
		goto L289
	}
L289:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1072
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+4))
	if v1072 < v1075 {
		v1235 = v1045
		goto L285
	} else {
		goto L290
	}
L290:
	;
	switch v1068 - int32(1) {
	case 0:
		goto L306
	case 1:
		goto L305
	case 2:
		goto L304
	case 3:
		goto L303
	case 4:
		goto L302
	case 5:
		goto L301
	case 6:
		goto L300
	case 7:
		goto L299
	case 8:
		goto L298
	case 9:
		goto L297
	case 10:
		goto L296
	case 11:
		goto L295
	case 12:
		goto L294
	case 13:
		goto L293
	case 14:
		goto L292
	default:
		goto L291
	}
L291:
	;
	v1235 = int32(1)
	goto L285
L292:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L339
L293:
	;
	v1171 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_32))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L6
	} else {
		goto L335
	}
L294:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1072 <= v1151 {
		v1235 = v1045
		goto L285
	} else {
		goto L331
	}
L295:
	;
	v1147 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_33))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L6
	} else {
		goto L329
	}
L296:
	;
	v1141 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_34))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L6
	} else {
		goto L327
	}
L297:
	;
	v1135 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_35))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L6
	} else {
		goto L325
	}
L298:
	;
	v1129 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_36))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L6
	} else {
		goto L323
	}
L299:
	;
	v1123 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_37))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L6
	} else {
		goto L321
	}
L300:
	;
	v1117 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_38))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L6
	} else {
		goto L319
	}
L301:
	;
	v1111 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_39))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L6
	} else {
		goto L317
	}
L302:
	;
	v1105 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_40))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L6
	} else {
		goto L315
	}
L303:
	;
	v1099 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_41))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L6
	} else {
		goto L313
	}
L304:
	;
	v1093 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_42))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L6
	} else {
		goto L311
	}
L305:
	;
	v1087 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_43))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L6
	} else {
		goto L309
	}
L306:
	;
	v1081 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_44))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L6
	} else {
		goto L307
	}
L307:
	;
	if int32(0) <= v1081 {
		goto L291
	} else {
		goto L308
	}
L308:
	;
	v1235 = v1081
	goto L285
L309:
	;
	if int32(0) <= v1087 {
		goto L291
	} else {
		goto L310
	}
L310:
	;
	v1235 = v1087
	goto L285
L311:
	;
	if int32(0) <= v1093 {
		goto L291
	} else {
		goto L312
	}
L312:
	;
	v1235 = v1093
	goto L285
L313:
	;
	if int32(0) <= v1099 {
		goto L291
	} else {
		goto L314
	}
L314:
	;
	v1235 = v1099
	goto L285
L315:
	;
	if int32(0) <= v1105 {
		goto L291
	} else {
		goto L316
	}
L316:
	;
	v1235 = v1105
	goto L285
L317:
	;
	if int32(0) <= v1111 {
		goto L291
	} else {
		goto L318
	}
L318:
	;
	v1235 = v1111
	goto L285
L319:
	;
	if int32(0) <= v1117 {
		goto L291
	} else {
		goto L320
	}
L320:
	;
	v1235 = v1117
	goto L285
L321:
	;
	if int32(0) <= v1123 {
		goto L291
	} else {
		goto L322
	}
L322:
	;
	v1235 = v1123
	goto L285
L323:
	;
	if int32(0) <= v1129 {
		goto L291
	} else {
		goto L324
	}
L324:
	;
	v1235 = v1129
	goto L285
L325:
	;
	if int32(0) <= v1135 {
		goto L291
	} else {
		goto L326
	}
L326:
	;
	v1235 = v1135
	goto L285
L327:
	;
	if int32(0) <= v1141 {
		goto L291
	} else {
		goto L328
	}
L328:
	;
	v1235 = v1141
	goto L285
L329:
	;
	if int32(0) <= v1147 {
		goto L291
	} else {
		goto L330
	}
L330:
	;
	v1235 = v1147
	goto L285
L331:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1153+v1072-int32(1)))))
	if v1157 != int32(108) {
		v1235 = v1045
		goto L285
	} else {
		goto L332
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1072 - int32(1)
	v1165 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_45))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L6
	} else {
		goto L333
	}
L333:
	;
	if int32(0) <= v1165 {
		goto L291
	} else {
		goto L334
	}
L334:
	;
	v1235 = v1165
	goto L285
L335:
	;
	if int32(0) <= v1171 {
		goto L291
	} else {
		goto L336
	}
L336:
	;
	v1235 = v1171
	goto L285
L337:
	;
	if v1227 != 0 {
		v1235 = v1045
		goto L285
	} else {
		goto L348
	}
L338:
	;
	v1227 = v1223
	goto L337
L339:
	;
	if v1183 <= v1184 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	v1223 = int32(0)
	goto L338
L341:
	;
	v1227 = int32(-1)
	goto L337
L342:
	;
	goto L343
L343:
	;
	v1196 = int32(1)
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197+v1183-v1196))))
	if int32(116) < v1201 {
		v1223 = v1196
		goto L338
	} else {
		goto L344
	}
L344:
	;
	v1203 = v1201 - int32(99)
	if v1203 < int32(0) {
		v1223 = v1196
		goto L338
	} else {
		goto L345
	}
L345:
	;
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1203)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v1209)>>(uint(v1203&int32(7))%32))&int32(1) == int32(0) {
		v1223 = v1196
		goto L338
	} else {
		goto L346
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1183 - int32(1)
	goto L347
L347:
	;
	goto L340
L348:
	;
	v1228 = F_slice_del(m, l0)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L6
	} else {
		goto L349
	}
L349:
	;
	if v1228 < int32(0) {
		v1235 = v1228
		goto L285
	} else {
		goto L350
	}
L350:
	;
	goto L291
L351:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1239
	v1241 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1239
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1239-int32(2) <= v1244 {
		v1314 = v1241
		goto L352
	} else {
		goto L353
	}
L352:
	;
	if v1314 < int32(0) {
		v1551 = v1314
		goto L1
	} else {
		goto L378
	}
L353:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1250 = int32(1)
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248+v1239-v1250))))
	if base.B2i32(v1252&int32(224) != int32(96))|base.B2i32(v1250<<(uint(v1252)%32)&int32(_a_F_english_ISO_8859_1_stem_46) == int32(0)) != 0 {
		v1314 = v1241
		goto L352
	} else {
		goto L354
	}
L354:
	;
	v1266 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_47), int32(9))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L6
	} else {
		goto L355
	}
L355:
	;
	if v1266 == int32(0) {
		v1314 = v1241
		goto L352
	} else {
		goto L356
	}
L356:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1270
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+4))
	if v1270 < v1273 {
		v1314 = v1241
		goto L352
	} else {
		goto L357
	}
L357:
	;
	switch v1266 - int32(1) {
	case 0:
		goto L364
	case 1:
		goto L363
	case 2:
		goto L362
	case 3:
		goto L361
	case 4:
		goto L360
	case 5:
		goto L359
	default:
		goto L358
	}
L358:
	;
	v1314 = int32(1)
	goto L352
L359:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1272)))
	if v1270 < v1305 {
		v1314 = v1241
		goto L352
	} else {
		goto L375
	}
L360:
	;
	v1301 = F_slice_del(m, l0)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L6
	} else {
		goto L373
	}
L361:
	;
	v1297 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_48))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L6
	} else {
		goto L371
	}
L362:
	;
	v1291 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_49))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L6
	} else {
		goto L369
	}
L363:
	;
	v1285 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_50))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L6
	} else {
		goto L367
	}
L364:
	;
	v1279 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_51))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L6
	} else {
		goto L365
	}
L365:
	;
	if int32(0) <= v1279 {
		goto L358
	} else {
		goto L366
	}
L366:
	;
	v1314 = v1279
	goto L352
L367:
	;
	if int32(0) <= v1285 {
		goto L358
	} else {
		goto L368
	}
L368:
	;
	v1314 = v1285
	goto L352
L369:
	;
	if int32(0) <= v1291 {
		goto L358
	} else {
		goto L370
	}
L370:
	;
	v1314 = v1291
	goto L352
L371:
	;
	if int32(0) <= v1297 {
		goto L358
	} else {
		goto L372
	}
L372:
	;
	v1314 = v1297
	goto L352
L373:
	;
	if int32(0) <= v1301 {
		goto L358
	} else {
		goto L374
	}
L374:
	;
	v1314 = v1301
	goto L352
L375:
	;
	v1307 = F_slice_del(m, l0)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L6
	} else {
		goto L376
	}
L376:
	;
	if v1307 < int32(0) {
		v1314 = v1307
		goto L352
	} else {
		goto L377
	}
L377:
	;
	goto L358
L378:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1319
	v1321 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1319
	v1325 = v1319 - int32(1)
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1325 <= v1326 {
		v1383 = v1321
		goto L379
	} else {
		goto L380
	}
L379:
	;
	if v1383 < int32(0) {
		v1551 = v1383
		goto L1
	} else {
		goto L394
	}
L380:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1328+v1325))))
	if base.B2i32(v1330&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1330)%32)&int32(_a_F_english_ISO_8859_1_stem_52) == int32(0)) != 0 {
		v1383 = v1321
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1344 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_53), int32(18))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L6
	} else {
		goto L382
	}
L382:
	;
	if v1344 == int32(0) {
		v1383 = v1321
		goto L379
	} else {
		goto L383
	}
L383:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1348
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1350)))
	if v1348 < v1351 {
		v1383 = v1321
		goto L379
	} else {
		goto L384
	}
L384:
	;
	switch v1344 - int32(1) {
	case 0:
		goto L387
	case 1:
		goto L386
	default:
		goto L385
	}
L385:
	;
	v1383 = int32(1)
	goto L379
L386:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1348 <= v1359 {
		v1383 = v1321
		goto L379
	} else {
		goto L390
	}
L387:
	;
	v1355 = F_slice_del(m, l0)
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L6
	} else {
		goto L388
	}
L388:
	;
	if int32(0) <= v1355 {
		goto L385
	} else {
		goto L389
	}
L389:
	;
	v1383 = v1355
	goto L379
L390:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1363 = int32(1)
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1361+v1348-v1363))))
	if base.Ui32(v1363) < base.Ui32((v1365-int32(115))&int32(255)) {
		v1383 = v1321
		goto L379
	} else {
		goto L391
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1348 - int32(1)
	v1375 = F_slice_del(m, l0)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L6
	} else {
		goto L392
	}
L392:
	;
	if v1375 < int32(0) {
		v1383 = v1375
		goto L379
	} else {
		goto L393
	}
L393:
	;
	goto L385
L394:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1386
	v1388 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1386
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1386 <= v1391 {
		v1488 = v1388
		goto L395
	} else {
		goto L396
	}
L395:
	;
	if v1488 < int32(0) {
		v1551 = v1488
		goto L1
	} else {
		goto L423
	}
L396:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1393+v1386-int32(1)))))
	switch v1397 - int32(101) {
	case 0, 7:
		goto L397
	default:
		v1488 = v1388
		goto L395
	}
L397:
	;
	v1402 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_54), int32(2))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L6
	} else {
		goto L398
	}
L398:
	;
	if v1402 == int32(0) {
		v1488 = v1388
		goto L395
	} else {
		goto L399
	}
L399:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1406
	switch v1402 - int32(1) {
	case 0:
		goto L402
	case 1:
		goto L401
	default:
		goto L400
	}
L400:
	;
	v1488 = int32(1)
	goto L395
L401:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1466)))
	if v1406 < v1467 {
		v1488 = v1388
		goto L395
	} else {
		goto L418
	}
L402:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1410)))
	if v1406 < v1411 {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1410)+4))
	if v1406 < v1413 {
		v1488 = v1388
		goto L395
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v1462 = F_slice_del(m, l0)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L6
	} else {
		goto L416
	}
L406:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1416 = int32(0)
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1425 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(89), int32(121), v1416)
	mBase = m.M
	if v1425 != 0 {
		goto L408
	} else {
		goto L409
	}
L407:
	;
	if v1456 != 0 {
		v1488 = v1388
		goto L395
	} else {
		goto L415
	}
L408:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1437 + (v1420 - v1415)
	v1445 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_27), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1445 != 0 {
		v1454 = v1416
		goto L412
	} else {
		goto L413
	}
L409:
	;
	v1430 = F_in_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_27), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1430 != 0 {
		goto L408
	} else {
		goto L410
	}
L410:
	;
	v1435 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_27), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1435 != 0 {
		goto L408
	} else {
		goto L411
	}
L411:
	;
	v1456 = int32(1)
	goto L407
L412:
	;
	v1456 = v1454
	goto L407
L413:
	;
	v1450 = F_in_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_27), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1450 != 0 {
		v1454 = v1416
		goto L412
	} else {
		goto L414
	}
L414:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1454 = base.B2i32(v1451 <= v1452)
	goto L412
L415:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1457 + (v1406 - v1415)
	goto L405
L416:
	;
	if int32(0) <= v1462 {
		goto L400
	} else {
		goto L417
	}
L417:
	;
	v1488 = v1462
	goto L395
L418:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1406 <= v1469 {
		v1488 = v1388
		goto L395
	} else {
		goto L419
	}
L419:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1471+v1406-int32(1)))))
	if v1475 != int32(108) {
		v1488 = v1388
		goto L395
	} else {
		goto L420
	}
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1406 - int32(1)
	v1481 = F_slice_del(m, l0)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L6
	} else {
		goto L421
	}
L421:
	;
	if v1481 < int32(0) {
		v1488 = v1481
		goto L395
	} else {
		goto L422
	}
L422:
	;
	goto L400
L423:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1496 = v1493
	goto L194
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1496
	goto L2
L425:
	;
	goto L426
L426:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1509 < v1508 {
		goto L428
	} else {
		goto L429
	}
L427:
	;
	v1551 = v1534
	goto L1
L428:
	;
	v1511 = v1508
	goto L430
L429:
	;
	v1511 = v1509
	goto L430
L430:
	;
	v1514 = v1508
	goto L431
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1514
	if v1509 != v1514 {
		goto L434
	} else {
		goto L435
	}
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1514
	v1529 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1514 + v1529
	v1534 = F_slice_from_s(m, l0, v1529, int32(_a_F_english_ISO_8859_1_stem_55))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L6
	} else {
		goto L439
	}
L433:
	;
	goto L432
L434:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1519+v1514))))
	if v1521 == int32(89) {
		goto L433
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	if v1514 == v1511 {
		goto L424
	} else {
		goto L438
	}
L437:
	;
	goto L436
L438:
	;
	v1526 = v1514 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1526
	v1514 = v1526
	goto L431
L439:
	;
	if int32(0) <= v1534 {
		goto L426
	} else {
		goto L440
	}
L440:
	;
	goto L427
}
func F_equality_ops_are_compatible(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	if l0 != l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(0)
	v13 = F_SearchSysCacheList(m, int32(3), int32(1), l0, v8, v8)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v64 = int32(1)
	goto L3
L3:
	;
	return v64
L4:
	;
	F_ReleaseCatCacheList(m, v13)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L16
	}
L5:
	;
	return int32(0)
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if v17 <= int32(0) {
		v57 = v8
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v24 = v8
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(48)+v24<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v35 = v33 + v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v38 = F_SearchSysCacheExists(m, int32(3), l1, int32(115), v36, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v57 = int32(0)
	goto L4
L10:
	;
	v51 = v24 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if v51 < v52 {
		v24 = v51
		goto L8
	} else {
		goto L15
	}
L11:
	;
	if v38 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	v44 = F_GetIndexAmRoutineByAmId(m, v42, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+13)))
	if v46 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v57 = int32(1)
	goto L4
L15:
	;
	goto L9
L16:
	;
	v64 = v57
	goto L3
}
func F_errcode_for_socket_access(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_errcode_for_socket_access[0]))
	if int32(0) <= v5 {
		v9 = v5 * int32(100)
		v12 = int32(100663808)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_errcode_for_socket_access[1])))
		switch v13 - int32(13) {
		case 0, 2, 10, 25, 26, 27, 51, 60:
			v19 = v12
		case 1, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 52, 53, 54, 55, 56, 57, 58, 59:
			v19 = int32(2600)
		default:
			if v13 == int32(142) {
				v19 = v12
			} else {
				v19 = int32(2600)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_errcode_for_socket_access[2]))) = v19
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_errcode_for_socket_access[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_errcode_for_socket_access_0), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_errcode_for_socket_access_1), int32(959), int32(_a_F_errcode_for_socket_access_2))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
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
func F_errdatatype(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_errdatatype_0), v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_errdatatype_1), int32(414), int32(_a_F_errdatatype_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
			v29 = v27 + v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
			v31 = F_get_namespace_name(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_err_generic_string(m, int32(115), v31)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_err_generic_string(m, int32(100), v29+int32(4))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_ReleaseCatCache(m, v9)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_errorMissingRTE(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v113 int32
	_ = v113
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v520 int32
	_ = v520
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l0 != 0 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	v154 = int32(0)
	v158 = F_RangeVarGetRelidExtended(m, l1, v154, int32(1), v154, v154)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L28
	} else {
		goto L29
	}
L3:
	;
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = l0
	v31 = v3
	goto L7
L5:
	;
	goto L6
L6:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v130 = F_get_visible_ENR_metadata(m, v129, v20)
	mBase = m.M
	goto L26
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v37 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v113 != 0 {
		v24 = v113
		v31 = v31 + int32(1)
		goto L7
	} else {
		goto L25
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v48 = int32(0)
	goto L12
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v43+v48<<(uint(int32(2))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v67 == int32(0))|base.B2i32(v67 != v70) != 0 {
		v88 = v67
		v89 = v70
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v164 = int32(0)
	v168 = v3
	v171 = v31
	v177 = int32(1)
	goto L1
L14:
	;
	if v88-v89 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	v73 = v64
	v74 = v20
	goto L17
L17:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v78 == int32(0) {
		v88 = v78
		v89 = v77
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v88 = v78
	v89 = v77
	goto L15
L19:
	;
	v81 = int32(1)
	if v78 == v77 {
		v73 = v73 + v81
		v74 = v74 + v81
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v92 = v48 + int32(1)
	if v92 != v40 {
		v48 = v92
		goto L12
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L13
L24:
	;
	goto L9
L25:
	;
	goto L8
L26:
	;
	if base.B2i32(v130 != int32(0)) == int32(0) {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v135 = int32(0)
	v164 = v135
	v168 = int32(1)
	v171 = v135
	v177 = v135
	goto L1
L28:
	;
	return
L29:
	;
	v160 = int32(0)
	v164 = v158
	v168 = v3
	v171 = v160
	v177 = v160
	goto L1
L30:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L28
	} else {
		goto L124
	}
L31:
	;
	v178 = int32(1)
	v188 = l0
	v192 = v3
	goto L34
L32:
	;
	goto L33
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L28
	} else {
		goto L119
	}
L34:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	if v197 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	if v520 != 0 {
		v188 = v520
		v192 = v192 + int32(1)
		goto L34
	} else {
		goto L118
	}
L37:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v200 <= int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v208 = int32(0)
	goto L39
L39:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v203+v208<<(uint(int32(2))%32))))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v225 = int32(0)
	if v224|base.B2i32(v164 == v225) == v225 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v332 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L41:
	;
	goto L40
L42:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v303 == int32(0))|base.B2i32(v303 != v306) != 0 {
		v324 = v303
		v325 = v306
		goto L69
	} else {
		goto L70
	}
L43:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	if v230 != v164 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v177^v178|base.B2i32(v224 != int32(6)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L41
L47:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v223)+88))
	if v237+v192 != v171 {
		goto L42
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v168^v178|base.B2i32(v224 != int32(7)) != 0 {
		goto L42
	} else {
		goto L59
	}
L50:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v223)+84))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v243 == int32(0))|base.B2i32(v243 != v246) != 0 {
		v264 = v243
		v265 = v246
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v264-v265 != 0 {
		goto L42
	} else {
		goto L58
	}
L52:
	;
	goto L51
L53:
	;
	v249 = v240
	v250 = v20
	goto L54
L54:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	if v254 == int32(0) {
		v264 = v254
		v265 = v253
		goto L52
	} else {
		goto L56
	}
L55:
	;
	v264 = v254
	v265 = v253
	goto L52
L56:
	;
	v257 = int32(1)
	if v254 == v253 {
		v249 = v249 + v257
		v250 = v250 + v257
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L41
L59:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v223)+108))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v273 == int32(0))|base.B2i32(v273 != v276) != 0 {
		v294 = v273
		v295 = v276
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v294-v295 == int32(0) {
		goto L41
	} else {
		goto L67
	}
L61:
	;
	goto L60
L62:
	;
	v279 = v270
	v280 = v20
	goto L63
L63:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	if v284 == int32(0) {
		v294 = v284
		v295 = v283
		goto L61
	} else {
		goto L65
	}
L64:
	;
	v294 = v284
	v295 = v283
	goto L61
L65:
	;
	v287 = int32(1)
	if v284 == v283 {
		v279 = v279 + v287
		v280 = v280 + v287
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L42
L68:
	;
	if v324-v325 == int32(0) {
		goto L41
	} else {
		goto L75
	}
L69:
	;
	goto L68
L70:
	;
	v309 = v300
	v310 = v20
	goto L71
L71:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	if v314 == int32(0) {
		v324 = v314
		v325 = v313
		goto L69
	} else {
		goto L73
	}
L72:
	;
	v324 = v314
	v325 = v313
	goto L69
L73:
	;
	v317 = int32(1)
	if v314 == v313 {
		v309 = v309 + v317
		v310 = v310 + v317
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v330 = v208 + int32(1)
	if v330 != v200 {
		v208 = v330
		goto L39
	} else {
		goto L76
	}
L76:
	;
	goto L36
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L28
	} else {
		goto L112
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L28
	} else {
		goto L92
	}
L79:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	if base.B2i32(v340 == int32(0))|base.B2i32(v340 != v343) != 0 {
		v361 = v340
		v362 = v343
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v361-v362 == int32(0) {
		goto L78
	} else {
		goto L87
	}
L81:
	;
	goto L80
L82:
	;
	v346 = v336
	v347 = v337
	goto L83
L83:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+1)))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+1)))
	if v351 == int32(0) {
		v361 = v351
		v362 = v350
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v361 = v351
	v362 = v350
	goto L81
L85:
	;
	v354 = int32(1)
	if v351 == v350 {
		v346 = v346 + v354
		v347 = v347 + v354
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v370 = F_refnameNamespaceItem(m, l0, int32(0), v336, v367, v18+int32(76))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L28
	} else {
		goto L88
	}
L88:
	;
	if v370 == int32(0) {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v374 != v223 {
		goto L78
	} else {
		goto L90
	}
L90:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v377 != 0 {
		goto L77
	} else {
		goto L91
	}
L91:
	;
	goto L78
L92:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L28
	} else {
		goto L93
	}
L93:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v386
	F_errmsg(m, int32(_a_F_errorMissingRTE_0), v18+int32(32))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L28
	} else {
		goto L94
	}
L94:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v394
	F_errdetail(m, int32(_a_F_errorMissingRTE_1), v18+int32(16))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L28
	} else {
		goto L95
	}
L95:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v401 != 0 {
		goto L30
	} else {
		goto L96
	}
L96:
	;
	v408 = l0
	goto L97
L97:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v408)+28))
	if v417 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L30
L99:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	if v474 != 0 {
		v408 = v474
		goto L97
	} else {
		goto L111
	}
L100:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	if v420 <= int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v428 = int32(0)
	goto L102
L102:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v423+v428<<(uint(int32(2))%32))))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v223 != v444 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+22)))
	if v449 != int32(1) {
		goto L30
	} else {
		goto L108
	}
L104:
	;
	v447 = v428 + int32(1)
	if v447 != v420 {
		v428 = v447
		goto L102
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	goto L103
L107:
	;
	goto L99
L108:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+23)))
	if v452 != int32(1) {
		goto L30
	} else {
		goto L109
	}
L109:
	;
	F_errhint(m, int32(_a_F_errorMissingRTE_2), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L28
	} else {
		goto L110
	}
L110:
	;
	goto L30
L111:
	;
	goto L98
L112:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L28
	} else {
		goto L113
	}
L113:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v482
	F_errmsg(m, int32(_a_F_errorMissingRTE_0), v18-int32(-64))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L28
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v377
	F_errhint(m, int32(_a_F_errorMissingRTE_3), v18+int32(48))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L28
	} else {
		goto L115
	}
L115:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L28
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_errorMissingRTE_4), int32(3743), int32(_a_F_errorMissingRTE_5))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L28
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	goto L35
L119:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L28
	} else {
		goto L120
	}
L120:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v543
	F_errmsg(m, int32(_a_F_errorMissingRTE_6), v18)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L28
	} else {
		goto L121
	}
L121:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L28
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_errorMissingRTE_4), int32(3761), int32(_a_F_errorMissingRTE_5))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L28
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errfinish(m, int32(_a_F_errorMissingRTE_4), int32(3754), int32(_a_F_errorMissingRTE_5))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L28
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_escape_json_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int64
	_ = v71
	var v77 int64
	_ = v77
	var v98 int64
	_ = v98
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v167 int64
	_ = v167
	var v170 int64
	_ = v170
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
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	F_enlargeStringInfo(m, l0, l2+int32(2))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 <= v22+int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v43 = l2 & int32(-8)
	v48 = int32(0)
	goto L8
L4:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v29+v22))) = uint8(v31)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = v33 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37+v35))) = uint8(v39)
	goto L3
L7:
	;
	goto L3
L8:
	;
	if v48 < v43 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v61 = v48
	v62 = v48
	goto L13
L11:
	;
	v227 = v48
	goto L12
L12:
	;
	v237 = v227 + int32(8)
	v241 = v227
	goto L34
L13:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(l1+v62)))
	if int64(0) <= v71 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v214 < v215 {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	goto L14
L16:
	;
	v204 = v62 - v61
	if int32(512) <= v204 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v77 = v71&int64(36170086419038336) ^ int64(-9187201950435737472)
	if v77&(v71-int64(2314885530818453536))|v77&(v71^int64(2459565876494606882)-int64(72340172838076673)) != int64(0) {
		v214 = v61
		v215 = v62
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v98 = int64(0)
	if base.B2i32(v71&int64(63050394783186944) == v98)|base.B2i32(v71&int64(246290604621824) == v98)|(base.B2i32(v71&int64(962072674304) == v98)|base.B2i32(v71&int64(3758096384) == v98))|(base.B2i32(v71&int64(57344) == v98)|(base.B2i32(v71&int64(14680064) == v98)|base.B2i32(v71&int64(224) == v98))) != 0 {
		v214 = v61
		v215 = v62
		goto L15
	} else {
		goto L22
	}
L20:
	;
	if v77&(v71^int64(6655295901103053916)-int64(72340172838076673)) == int64(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v214 = v61
	v215 = v62
	goto L15
L22:
	;
	v131 = v71 ^ int64(2459565876494606882)
	v134 = int64(0)
	if base.B2i32(v131&int64(71776119061217280) == v134)|base.B2i32(v131&int64(280375465082880) == v134)|(base.B2i32(v131&int64(1095216660480) == v134)|base.B2i32(v131&int64(4278190080) == v134))|(base.B2i32(v131&int64(65280) == v134)|(base.B2i32(v131&int64(16711680) == v134)|base.B2i32(v131&int64(255) == v134))) != 0 {
		v214 = v61
		v215 = v62
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v167 = v71 ^ int64(6655295901103053916)
	v170 = int64(0)
	if base.B2i32(v167&int64(71776119061217280) == v170)|base.B2i32(v167&int64(280375465082880) == v170)|(base.B2i32(v167&int64(1095216660480) == v170)|base.B2i32(v167&int64(4278190080) == v170))|(base.B2i32(v167&int64(65280) == v170)|(base.B2i32(v167&int64(16711680) == v170)|base.B2i32(v167&int64(255) == v170))) != 0 {
		v214 = v61
		v215 = v62
		goto L15
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	F_appendBinaryStringInfo(m, l0, l1+v61, v204)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v210 = v61
	goto L27
L27:
	;
	v212 = v62 + int32(8)
	if v212 < v43 {
		v61 = v210
		v62 = v212
		goto L13
	} else {
		goto L29
	}
L28:
	;
	v210 = v62
	goto L27
L29:
	;
	v214 = v210
	v215 = v212
	goto L15
L30:
	;
	F_appendBinaryStringInfo(m, l0, l1+v214, v215-v214)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v227 = v215
	goto L12
L33:
	;
	goto L32
L34:
	;
	if l2 != v241 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v48 = v237
	goto L8
L36:
	;
	v331 = v241 + int32(1)
	if v331 != v237 {
		v241 = v331
		goto L34
	} else {
		goto L70
	}
L37:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L69
	}
L38:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v241))))
	switch v252 - int32(8) {
	case 0:
		goto L48
	case 1:
		goto L44
	case 2:
		goto L46
	case 3, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		goto L41
	case 4:
		goto L47
	case 5:
		goto L45
	case 26:
		goto L43
	default:
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v300 <= v301+int32(1) {
		goto L65
	} else {
		goto L66
	}
L41:
	;
	v275 = base.I32_extend8_s(v252)
	if base.Ui32(v252) <= base.Ui32(int32(31)) {
		goto L56
	} else {
		goto L57
	}
L42:
	;
	if v252 == int32(92) {
		goto L37
	} else {
		goto L55
	}
L43:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L54
	}
L44:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_2))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L53
	}
L45:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_3))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L52
	}
L46:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_4))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L51
	}
L47:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_5))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_6))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L36
L50:
	;
	goto L36
L51:
	;
	goto L36
L52:
	;
	goto L36
L53:
	;
	goto L36
L54:
	;
	goto L36
L55:
	;
	goto L41
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v275
	F_appendStringInfo(m, l0, int32(_a_F_escape_json_with_len_7), v15)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v282 <= v283+int32(1) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L36
L60:
	;
	F_appendStringInfoChar(m, l0, v275)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v289+v283))) = uint8(v252)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v294 = v292 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v298 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v296+v294))) = uint8(v298)
	goto L36
L63:
	;
	goto L36
L64:
	;
	m.G0 = v15 + int32(16)
	return
L65:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v310 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v308+v301))) = uint8(v310)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v314 = v312 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v318 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v316+v314))) = uint8(v318)
	goto L64
L68:
	;
	goto L64
L69:
	;
	goto L36
L70:
	;
	goto L35
}
func F_eval_const_expressions_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
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
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
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
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
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
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
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
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
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
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int64
	_ = v889
	var v891 int64
	_ = v891
	var v893 int64
	_ = v893
	var v895 int64
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1029 int32
	_ = v1029
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1138 int32
	_ = v1138
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1354 int32
	_ = v1354
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1553 int32
	_ = v1553
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1598 int32
	_ = v1598
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1639 int32
	_ = v1639
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1716 int32
	_ = v1716
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(112)
	m.G0 = v17
	F_check_stack_depth(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l0 == int32(0) {
		v1755 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v17 + int32(112)
	return v1755
L4:
	;
	v25 = l0
	goto L17
L5:
	;
	v1751 = F_copyObjectImpl(m, v25)
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L1
	} else {
		goto L522
	}
L6:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1706)+20)) = base.B2i32(v1743 == int32(0))
	v1755 = v1706
	goto L3
L7:
	;
	v1724 = F_palloc0(m, int32(36))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L1
	} else {
		goto L521
	}
L8:
	;
	if v1655&int32(1) != 0 {
		goto L514
	} else {
		goto L515
	}
L9:
	;
	v1684 = int32(0)
	v1686 = F_makeBoolConst(m, v1684, v1684)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L1
	} else {
		goto L513
	}
L10:
	;
	if v1660 != 0 {
		v1716 = v290
		goto L7
	} else {
		goto L511
	}
L11:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1639+v1627<<(uint(int32(2))%32))))
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1643)))
	if v1644 != int32(7) {
		v1716 = v290
		goto L7
	} else {
		goto L510
	}
L12:
	;
	if v307&int32(1) == int32(0) {
		v1655 = v355
		v1657 = v356
		v1660 = v357
		goto L10
	} else {
		goto L509
	}
L13:
	;
	v1612 = F_palloc0(m, int32(20))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L1
	} else {
		goto L508
	}
L14:
	;
	if v1281|base.B2i32(v1277 == int32(0)) != 0 {
		goto L13
	} else {
		goto L498
	}
L15:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1512 = F_eval_const_expressions_mutator(m, v1511, l1)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L477
	}
L16:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1453 = F_eval_const_expressions_mutator(m, v1452, l1)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L1
	} else {
		goto L458
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v39 != int32(319) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v1450 = F_expression_tree_mutator_impl(m, v25, int32(871), l1)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L1
	} else {
		goto L457
	}
L19:
	;
	goto L18
L20:
	;
	switch v39 - int32(8) {
	case 0:
		goto L45
	default:
		goto L19
	case 3:
		goto L44
	case 6, 27, 28, 31:
		goto L30
	case 7:
		goto L43
	case 9:
		goto L42
	case 10:
		goto L41
	case 11:
		goto L40
	case 12:
		goto L39
	case 13:
		goto L38
	case 15, 16:
		v1755 = v25
		goto L3
	case 17:
		goto L27
	case 19:
		goto L36
	case 20:
		goto L35
	case 21:
		goto L34
	case 22:
		goto L23
	case 23:
		goto L33
	case 24:
		goto L32
	case 26:
		goto L31
	case 30:
		goto L29
	case 32:
		goto L28
	case 36:
		goto L37
	case 44:
		goto L26
	case 45:
		goto L15
	case 47:
		goto L16
	}
L21:
	;
	goto L22
L22:
	;
	v1442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1442 != int32(1) {
		goto L19
	} else {
		goto L454
	}
L23:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1401 = F_eval_const_expressions_mutator(m, v1400, l1)
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L1
	} else {
		goto L436
	}
L24:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	v1755 = v1399
	goto L3
L25:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	v1388 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+8)))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+28))
	v1393 = F_makeVar(m, v1387, v1388, v1389, v1390, v1391, v1392)
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L1
	} else {
		goto L435
	}
L26:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1277 = F_eval_const_expressions_mutator(m, v1276, l1)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L404
	}
L27:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1186 = F_eval_const_expressions_mutator(m, v1185, l1)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L374
	}
L28:
	;
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1177 != int32(1) {
		goto L5
	} else {
		goto L371
	}
L29:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v1091 != 0 {
		goto L351
	} else {
		goto L352
	}
L30:
	;
	v1077 = F_expression_tree_mutator_impl(m, v25, int32(871), l1)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L343
	}
L31:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1071 == int32(0) {
		goto L5
	} else {
		goto L341
	}
L32:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v946 = F_eval_const_expressions_mutator(m, v945, l1)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L316
	}
L33:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v933 = F_eval_const_expressions_mutator(m, v932, l1)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L310
	}
L34:
	;
	v885 = F_palloc0(m, int32(32))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L297
	}
L35:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v783
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v783
	v789 = F_list_make1_impl(m, int32(1), v17+int32(44))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L283
	}
L36:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v773 = F_eval_const_expressions_mutator(m, v772, l1)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L281
	}
L37:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v760 = F_eval_const_expressions_mutator(m, v759, l1)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L273
	}
L38:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	switch v476 {
	case 0:
		goto L174
	case 1:
		goto L175
	case 2:
		goto L173
	default:
		goto L172
	}
L39:
	;
	v448 = F_expression_tree_mutator_impl(m, v25, int32(871), l1)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L158
	}
L40:
	;
	v365 = F_expression_tree_mutator_impl(m, v25, int32(871), l1)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L132
	}
L41:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v290 = F_expression_tree_mutator_impl(m, v288, int32(871), l1)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L108
	}
L42:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v208
	F_set_opfuncid(m, v25)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L83
	}
L43:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v174 = F_exprTypmod(m, v25)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L79
	}
L44:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v115 = F_SearchSysCache1(m, int32(47), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L67
	}
L45:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v44 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v45 == int32(0) {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v48 <= int32(0) {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+28))
	if v51 < v48 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v53 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v65 == int32(0) {
		goto L5
	} else {
		goto L55
	}
L51:
	;
	v57 = m.T0[v53].(func(*base.Module, int32, int32, int32, int32) int32)(m, v45, v48, int32(1), v17+int32(100))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v64 = v45 + v48*int32(12) + int32(20)
	goto L50
L54:
	;
	v64 = v57
	goto L50
L55:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v65 != v68 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v70 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+6)))
	if v73&int32(1) == int32(0) {
		goto L5
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_get_typlenbyval(m, v65, v17+int32(96), v17+int32(88))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)))
	if v85|v86&int32(1) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+96)))
	v94 = F_datumCopy(m, v84, int32(0), v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v98 = v86
	v99 = v84
	v100 = v85
	goto L64
L64:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+96)))
	v105 = int32(1)
	v109 = F_makeConst(m, v101, v102, v103, v104, v99, v100&v105, v98&v105)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)))
	v98 = v97
	v99 = v94
	v100 = v96
	goto L64
L66:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = v111
	v1755 = v109
	goto L3
L67:
	;
	if v115 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v135 = F_expand_function_arguments(m, v132, int32(0), v134, v115)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L74
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v114
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_0), v17)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(2544), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_ReleaseCatCache(m, v115)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v140 = F_expression_tree_mutator_impl(m, v135, int32(871), l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v143 = F_eval_const_expressions_mutator(m, v142, l1)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v146 = F_palloc0(m, int32(44))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = int32(11)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+8)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+12)) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+24)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v146)+20)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v146)+16)) = v156
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+28)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+32)) = v162
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v146)+36)) = uint8(v164)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v146)+37)) = uint8(v166)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+40)) = v168
	v1755 = v146
	goto L3
L79:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+13)))
	v181 = int32(1)
	v183 = F_simplify_function(m, v172, v173, v174, v176, v177, v17+int32(100), v180, v181, v181, l1)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v183 != 0 {
		v1755 = v183
		goto L3
	} else {
		goto L81
	}
L81:
	;
	v186 = F_palloc0(m, int32(36))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = int32(15)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+8)) = v192
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+12)) = uint8(v194)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+13)) = uint8(v196)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+16)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+24)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+28)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+32)) = v206
	v1755 = v186
	goto L3
L83:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v220 = int32(1)
	v222 = F_simplify_function(m, v212, v213, int32(-1), v215, v216, v17+int32(100), int32(0), v220, v220, l1)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v222 != 0 {
		v1755 = v222
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	switch v225 - int32(85) {
	case 0, 6:
		goto L87
	default:
		goto L86
	}
L86:
	;
	v269 = F_palloc0(m, int32(36))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L107
	}
L87:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	if v230 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	if v261 != 0 {
		v1755 = v261
		goto L3
	} else {
		goto L106
	}
L89:
	;
	v257 = F_negate_clause(m, v254)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L105
	}
L90:
	;
	if v229 == int32(0) {
		v261 = v3
		goto L88
	} else {
		goto L98
	}
L91:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v233 != int32(7) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	if v225 == int32(91) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v236 == int32(0) {
		v254 = v229
		goto L89
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v236 != 0 {
		v254 = v229
		goto L89
	} else {
		goto L97
	}
L96:
	;
	v261 = v229
	goto L88
L97:
	;
	v261 = v229
	goto L88
L98:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	if v243 != int32(7) {
		v261 = v3
		goto L88
	} else {
		goto L99
	}
L99:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	if v225 == int32(91) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	if v246 == int32(0) {
		v254 = v230
		goto L89
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v246 == int32(0) {
		v261 = v230
		goto L88
	} else {
		goto L104
	}
L103:
	;
	v261 = v230
	goto L88
L104:
	;
	v254 = v230
	goto L89
L105:
	;
	v261 = v257
	goto L88
L106:
	;
	goto L86
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = int32(17)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+4)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+8)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+12)) = v277
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v269)+16)) = uint8(v279)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+20)) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+28)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v269)+24)) = v283
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+32)) = v286
	v1755 = v269
	goto L3
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v290
	if v290 == int32(0) {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	v295 = int32(0)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	if v296 <= v295 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v1655 = v295
	v1657 = int32(1)
	v1660 = v3
	goto L10
L111:
	;
	goto L112
L112:
	;
	v300 = int32(1)
	if v296 == v300 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v1627 = int32(0)
	v1628 = v295
	v1630 = v300
	v1633 = v3
	goto L11
L114:
	;
	goto L115
L115:
	;
	v304 = int32(0)
	if v304 < v296 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v307 = v296
	goto L118
L117:
	;
	v307 = v304
	goto L118
L118:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	v316 = int32(0)
	v317 = v295
	v319 = v300
	v320 = v3
	v322 = v3
	goto L119
L119:
	;
	v330 = v312 + v316<<(uint(int32(2))%32)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	if v332 != int32(7) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L12
L121:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	if v346 != int32(7) {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	v341 = v317
	v342 = v319
	v343 = int32(1)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+24)))
	v341 = (v317 | v336) & int32(1)
	v342 = v319 & v336
	v343 = v322
	goto L121
L125:
	;
	v359 = int32(2)
	v360 = v316 + v359
	v362 = v320 + v359
	if v307&int32(2147483646) != v362 {
		v316 = v360
		v317 = v355
		v319 = v356
		v320 = v362
		v322 = v357
		goto L119
	} else {
		goto L129
	}
L126:
	;
	v355 = v341
	v356 = v342
	v357 = int32(1)
	goto L125
L127:
	;
	goto L128
L128:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+24)))
	v355 = (v341 | v350) & int32(1)
	v356 = v342 & v350
	v357 = v343
	goto L125
L129:
	;
	goto L120
L130:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
	v428 = F_func_volatile(m, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L148
	}
L131:
	;
	F_set_opfuncid(m, v365)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L147
	}
L132:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+28))
	if v367 == int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v370 <= int32(0) {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	v373 = int32(0)
	if v373 < v370 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v376 = v370
	goto L137
L136:
	;
	v376 = v373
	goto L137
L137:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v379 = int32(0)
	v384 = v3
	goto L138
L138:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v377+v379<<(uint(int32(2))%32))))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	if v398 == int32(7) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	F_set_opfuncid(m, v365)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L145
	}
L140:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+24)))
	if v401 != 0 {
		goto L24
	} else {
		goto L143
	}
L141:
	;
	v402 = int32(1)
	goto L142
L142:
	;
	v404 = v379 + int32(1)
	if v404 != v376 {
		v379 = v404
		v384 = v402
		goto L138
	} else {
		goto L144
	}
L143:
	;
	v402 = v384
	goto L142
L144:
	;
	goto L139
L145:
	;
	if v402&int32(1) != 0 {
		v1755 = v365
		goto L3
	} else {
		goto L146
	}
L146:
	;
	goto L130
L147:
	;
	goto L130
L148:
	;
	if v428 != int32(105) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	if v428 != int32(115) {
		v1755 = v365
		goto L3
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v439 = F_exprType(m, v365)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v434&int32(1) == int32(0) {
		v1755 = v365
		goto L3
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v441 = F_exprTypmod(m, v365)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v443 = F_exprCollation(m, v365)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v445 = F_evaluate_expr(m, v365, v439, v441, v443)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v1755 = v445
	goto L3
L158:
	;
	F_set_opfuncid(m, v448)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v454 = F_expression_tree_walker_impl(m, v448, int32(872), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	if v454 != 0 {
		v1755 = v448
		goto L3
	} else {
		goto L161
	}
L161:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v457 = F_func_volatile(m, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	if v457 != int32(105) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	if v457 != int32(115) {
		v1755 = v448
		goto L3
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v468 = F_exprType(m, v448)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v463&int32(1) == int32(0) {
		v1755 = v448
		goto L3
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v470 = F_exprTypmod(m, v448)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v472 = F_exprCollation(m, v448)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v474 = F_evaluate_expr(m, v448, v468, v470, v472)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v1755 = v474
	goto L3
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L270
	}
L173:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)+12))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v736)))
	v738 = F_eval_const_expressions_mutator(m, v737, l1)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L268
	}
L174:
	;
	v609 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+100)) = uint8(v609)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+96)) = uint8(v609)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v615 = F_list_copy(m, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L223
	}
L175:
	;
	v477 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+100)) = uint8(v477)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+96)) = uint8(v477)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v483 = F_list_copy(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L177
	}
L176:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+96)))
	if v579 == int32(1) {
		goto L205
	} else {
		goto L206
	}
L177:
	;
	if v483 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v489 = v483
	v491 = v477
	goto L181
L179:
	;
	v552 = v477
	goto L180
L180:
	;
	v578 = v552
	goto L176
L181:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v489)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	v505 = F_list_delete_first(m, v489)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L183
	}
L182:
	;
	v552 = v547
	goto L180
L183:
	;
	if v504 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v546 != 0 {
		v489 = v546
		v491 = v547
		goto L181
	} else {
		goto L204
	}
L185:
	;
	v520 = F_eval_const_expressions_mutator(m, v504, l1)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L192
	}
L186:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	if v509 != int32(21) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	if v512 != int32(1) {
		goto L185
	} else {
		goto L188
	}
L188:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v504)+8))
	v516 = F_list_concat_copy(m, v515, v505)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_list_free(m, v505)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v546 = v516
	v547 = v491
	goto L184
L191:
	;
	v544 = F_lappend(m, v491, v520)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L203
	}
L192:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v524 = v522 - int32(7)
	if v524 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if v524 != int32(14) {
		goto L191
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+24)))
	if v533 == int32(1) {
		goto L199
	} else {
		goto L200
	}
L196:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v527 != int32(1) {
		goto L191
	} else {
		goto L197
	}
L197:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v520)+8))
	v531 = F_list_concat_copy(m, v530, v505)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v546 = v531
	v547 = v491
	goto L184
L199:
	;
	v536 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(100)))) = uint8(v536)
	v546 = v505
	v547 = v491
	goto L184
L200:
	;
	goto L201
L201:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v520)+20))
	if v538 == int32(0) {
		v546 = v505
		v547 = v491
		goto L184
	} else {
		goto L202
	}
L202:
	;
	v541 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(96)))) = uint8(v541)
	v578 = int32(0)
	goto L176
L203:
	;
	v546 = v505
	v547 = v544
	goto L184
L204:
	;
	goto L182
L205:
	;
	v584 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+100)))
	if v586 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v1755 = v584
	goto L3
L209:
	;
	v591 = F_makeBoolConst(m, int32(0), int32(1))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L212
	}
L210:
	;
	v595 = v578
	goto L211
L211:
	;
	if v595 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v593 = F_lappend(m, v578, v591)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v595 = v593
	goto L211
L214:
	;
	v598 = int32(0)
	v600 = F_makeBoolConst(m, v598, v598)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v602 == int32(1) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1755 = v600
	goto L3
L218:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v595)+12))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	v1755 = v606
	goto L3
L219:
	;
	goto L220
L220:
	;
	v607 = F_make_orclause(m, v595)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v1755 = v607
	goto L3
L222:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+96)))
	if v705 == int32(1) {
		goto L251
	} else {
		goto L252
	}
L223:
	;
	if v615 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v621 = v615
	v623 = v609
	goto L227
L225:
	;
	v678 = v609
	goto L226
L226:
	;
	v704 = v678
	goto L222
L227:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v621)+12))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	v637 = F_list_delete_first(m, v621)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L229
	}
L228:
	;
	v678 = v673
	goto L226
L229:
	;
	if v636 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	if v672 != 0 {
		v621 = v672
		v623 = v673
		goto L227
	} else {
		goto L250
	}
L231:
	;
	v650 = F_eval_const_expressions_mutator(m, v636, l1)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L238
	}
L232:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	if v641 != int32(21) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	if v644 != 0 {
		goto L231
	} else {
		goto L234
	}
L234:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v636)+8))
	v646 = F_list_concat_copy(m, v645, v637)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_list_free(m, v637)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v672 = v646
	v673 = v623
	goto L230
L237:
	;
	v670 = F_lappend(m, v623, v650)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L249
	}
L238:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v650)))
	v654 = v652 - int32(7)
	if v654 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	if v654 != int32(14) {
		goto L237
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650)+24)))
	if v661 == int32(1) {
		goto L245
	} else {
		goto L246
	}
L242:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v650)+4))
	if v657 != 0 {
		goto L237
	} else {
		goto L243
	}
L243:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v650)+8))
	v659 = F_list_concat_copy(m, v658, v637)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v672 = v659
	v673 = v623
	goto L230
L245:
	;
	v664 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(100)))) = uint8(v664)
	v672 = v637
	v673 = v623
	goto L230
L246:
	;
	goto L247
L247:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v650)+20))
	if v666 != 0 {
		v672 = v637
		v673 = v623
		goto L230
	} else {
		goto L248
	}
L248:
	;
	v667 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(96)))) = uint8(v667)
	v704 = int32(0)
	goto L222
L249:
	;
	v672 = v637
	v673 = v670
	goto L230
L250:
	;
	goto L228
L251:
	;
	v708 = int32(0)
	v710 = F_makeBoolConst(m, v708, v708)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+100)))
	if v712 == int32(1) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1755 = v710
	goto L3
L255:
	;
	v717 = F_makeBoolConst(m, int32(0), int32(1))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	v721 = v704
	goto L257
L257:
	;
	if v721 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	v719 = F_lappend(m, v704, v717)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v721 = v719
	goto L257
L260:
	;
	v726 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v721)+4))
	if v728 == int32(1) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1755 = v726
	goto L3
L264:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v721)+12))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v1755 = v732
	goto L3
L265:
	;
	goto L266
L266:
	;
	v733 = F_make_andclause(m, v721)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	v1755 = v733
	goto L3
L268:
	;
	v740 = F_negate_clause(m, v738)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	v1755 = v740
	goto L3
L270:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v746
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_3), v17+int32(16))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(2910), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	if v760 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	if v762 == int32(7) {
		v1755 = v760
		goto L3
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v765 = F_eval_const_expressions_mutator(m, v758, l1)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L278
	}
L277:
	;
	goto L276
L278:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v768 = F_copyObjectImpl(m, v767)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v770 = F_makeJsonValueExpr(m, v765, v760, v768)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v1755 = v770
	goto L3
L281:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v781 = F_applyRelabelType(m, v773, v775, v776, v777, v778, v779, int32(1))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	v1755 = v781
	goto L3
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v789
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v793 = F_exprType(m, v792)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	F_getTypeOutputInfo(m, v793, v17+int32(96), v17+int32(95))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	F_getTypeInputInfo(m, v801, v17+int32(88), v17+int32(84))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v17)+96))
	v811 = int32(0)
	v814 = v17 + int32(100)
	v816 = int32(1)
	v818 = F_simplify_function(m, v808, int32(2275), int32(-1), v811, v811, v814, v811, v816, v816, l1)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	if v818 != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v818
	v823 = int32(0)
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v17)+84))
	v828 = F_makeConst(m, int32(26), int32(-1), v823, int32(4), v825, v823, int32(1))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v868 = F_palloc0(m, int32(24))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L296
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v828
	v832 = int32(-1)
	v833 = int32(0)
	v838 = F_makeConst(m, int32(23), v832, v833, int32(4), v832, v833, int32(1))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v838
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v842
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v844
	v852 = F_list_make3_impl(m, v17+int32(40), v17+int32(36), v17+int32(32))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v852
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v859 = int32(0)
	v863 = F_simplify_function(m, v855, v856, int32(-1), v858, v859, v814, v859, v859, int32(1), l1)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	if v863 != 0 {
		v1755 = v863
		goto L3
	} else {
		goto L295
	}
L295:
	;
	goto L290
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v868))) = int32(28)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v872)+12))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	*(*int32)(unsafe.Add(mBase, uint32(v868)+4)) = v874
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v868)+8)) = v876
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v868)+12)) = v878
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v868)+16)) = v880
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v868)+20)) = v882
	v1755 = v868
	goto L3
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v885))) = int32(29)
	v889 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v885))) = v889
	v891 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v885)+8)) = v891
	v893 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v885)+16)) = v893
	v895 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v885)+24)) = v895
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v885)+4))
	v898 = F_eval_const_expressions_mutator(m, v897, l1)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v885)+4)) = v898
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v885)+8))
	v905 = F_eval_const_expressions_mutator(m, v904, l1)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v885)+8)) = v905
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v901
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v885)+4))
	if v909 == int32(0) {
		v1755 = v885
		goto L3
	} else {
		goto L300
	}
L300:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v909)))
	if v912 != int32(7) {
		v1755 = v885
		goto L3
	} else {
		goto L301
	}
L301:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v885)+8))
	if v915 == int32(0) {
		v1755 = v885
		goto L3
	} else {
		goto L302
	}
L302:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	if v918 == int32(55) {
		v1755 = v885
		goto L3
	} else {
		goto L303
	}
L303:
	;
	v922 = F_contain_mutable_functions_walker(m, v915, int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	if v922 != 0 {
		v1755 = v885
		goto L3
	} else {
		goto L305
	}
L305:
	;
	v924 = F_exprType(m, v885)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	v926 = F_exprTypmod(m, v885)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	v928 = F_exprCollation(m, v885)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v930 = F_evaluate_expr(m, v885, v924, v926, v928)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v1755 = v930
	goto L3
L310:
	;
	v935 = F_exprType(m, v933)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	v937 = F_exprTypmod(m, v933)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v943 = F_applyRelabelType(m, v933, v935, v937, v939, int32(2), v941, int32(1))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	v1755 = v943
	goto L3
L314:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v955
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v959 == int32(0) {
		v1029 = v3
		goto L320
	} else {
		goto L321
	}
L315:
	;
	v955 = int32(0)
	v956 = v946
	goto L314
L316:
	;
	if v946 == int32(0) {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	if v950 != int32(7) {
		goto L315
	} else {
		goto L318
	}
L318:
	;
	v955 = v946
	v956 = int32(0)
	goto L314
L319:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1050)))
	v1052 = F_eval_const_expressions_mutator(m, v1051, l1)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L336
	}
L320:
	;
	v1045 = v1029
	v1050 = v25 + int32(20)
	goto L319
L321:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	if v962 <= int32(0) {
		v1029 = v3
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v969 = int32(0)
	v975 = v3
	goto L323
L323:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v959)+12))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v980+v969<<(uint(int32(2))%32))))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)+4))
	v986 = F_eval_const_expressions_mutator(m, v985, l1)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L327
	}
L324:
	;
	v1029 = v1014
	goto L320
L325:
	;
	v1017 = v969 + int32(1)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	if v1017 < v1018 {
		v969 = v1017
		v975 = v1014
		goto L323
	} else {
		goto L335
	}
L326:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v984)+8))
	v1000 = F_eval_const_expressions_mutator(m, v999, l1)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L332
	}
L327:
	;
	if v986 == int32(0) {
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	if v990 != int32(7) {
		goto L326
	} else {
		goto L329
	}
L329:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986)+24)))
	if v993 != 0 {
		v1014 = v975
		goto L325
	} else {
		goto L330
	}
L330:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v986)+20))
	if v994 == int32(0) {
		v1014 = v975
		goto L325
	} else {
		goto L331
	}
L331:
	;
	v1045 = v975
	v1050 = v984 + int32(8)
	goto L319
L332:
	;
	v1003 = F_palloc0(m, int32(16))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+8)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+4)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v1003))) = int32(33)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v984)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+12)) = v1009
	v1011 = F_lappend(m, v975, v1003)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	v1014 = v1011
	goto L325
L335:
	;
	goto L324
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v957
	if v1045 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1755 = v1052
	goto L3
L338:
	;
	goto L339
L339:
	;
	v1058 = F_palloc0(m, int32(28))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1058))) = int32(32)
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1058)+4)) = v1062
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1058)+20)) = v1052
	*(*int32)(unsafe.Add(mBase, uint32(v1058)+16)) = v1045
	*(*int32)(unsafe.Add(mBase, uint32(v1058)+12)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v1058)+8)) = v1064
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1058)+24)) = v1069
	v1755 = v1058
	goto L3
L341:
	;
	v1074 = F_copyObjectImpl(m, v1071)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v1755 = v1074
	goto L3
L343:
	;
	v1081 = F_expression_tree_walker_impl(m, v1077, int32(872), int32(0))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	if v1081 != 0 {
		v1755 = v1077
		goto L3
	} else {
		goto L345
	}
L345:
	;
	v1083 = F_exprType(m, v1077)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v1085 = F_exprTypmod(m, v1077)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	v1087 = F_exprCollation(m, v1077)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	v1089 = F_evaluate_expr(m, v1077, v1083, v1085, v1087)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	v1755 = v1089
	goto L3
L350:
	;
	v1166 = F_palloc0(m, int32(20))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L370
	}
L351:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+4))
	if v1092 <= int32(0) {
		v1138 = v3
		goto L354
	} else {
		goto L355
	}
L352:
	;
	goto L353
L353:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1163 = F_makeNullConst(m, v1160, int32(-1), v1162)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L369
	}
L354:
	;
	if v1138 != 0 {
		goto L350
	} else {
		goto L368
	}
L355:
	;
	v1099 = int32(0)
	v1102 = v3
	goto L356
L356:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+12))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1110+v1099<<(uint(int32(2))%32))))
	v1115 = F_eval_const_expressions_mutator(m, v1114, l1)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L359
	}
L357:
	;
	v1138 = v1127
	goto L354
L358:
	;
	v1129 = v1099 + int32(1)
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+4))
	if v1129 < v1130 {
		v1099 = v1129
		v1102 = v1127
		goto L356
	} else {
		goto L367
	}
L359:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1115)))
	if v1117 == int32(7) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115)+24)))
	if v1120 != 0 {
		v1127 = v1102
		goto L358
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	v1125 = F_lappend(m, v1102, v1115)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L366
	}
L363:
	;
	if v1102 == int32(0) {
		v1755 = v1115
		goto L3
	} else {
		goto L364
	}
L364:
	;
	v1123 = F_lappend(m, v1102, v1115)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1138 = v1123
	goto L354
L366:
	;
	v1127 = v1125
	goto L358
L367:
	;
	goto L357
L368:
	;
	goto L353
L369:
	;
	v1755 = v1163
	goto L3
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1166))) = int32(38)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+4)) = v1170
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+12)) = v1138
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+8)) = v1172
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+16)) = v1175
	v1755 = v1166
	goto L3
L371:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1183 = F_evaluate_expr(m, v25, v1180, v1181, int32(0))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	v1755 = v1183
	goto L3
L373:
	;
	v1245 = F_palloc0(m, int32(24))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L395
	}
L374:
	;
	if v1186 == int32(0) {
		goto L373
	} else {
		goto L375
	}
L375:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	if v1190 == int32(6) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1186)+8)))
	if v1193 != 0 {
		goto L373
	} else {
		goto L379
	}
L377:
	;
	v1203 = v1190
	goto L378
L378:
	;
	if v1203 != int32(36) {
		goto L373
	} else {
		goto L383
	}
L379:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+28))
	if v1194 != 0 {
		goto L373
	} else {
		goto L380
	}
L380:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+12))
	v1196 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+8)))
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1200 = F_rowtype_field_matches(m, v1195, v1196, v1197, v1198, v1199)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	if v1200 != 0 {
		goto L25
	} else {
		goto L382
	}
L382:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	v1203 = v1202
	goto L378
L383:
	;
	v1206 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+8)))
	if v1206 <= int32(0) {
		goto L373
	} else {
		goto L384
	}
L384:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	if v1209 == int32(0) {
		goto L373
	} else {
		goto L385
	}
L385:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+4))
	if v1212 < v1206 {
		goto L373
	} else {
		goto L386
	}
L386:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+12))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1214+v1206<<(uint(int32(2))%32)-int32(4))))
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+8))
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1225 = F_rowtype_field_matches(m, v1221, v1206, v1222, v1223, v1224)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	if v1225 == int32(0) {
		goto L373
	} else {
		goto L388
	}
L388:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1230 = F_exprType(m, v1220)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	if v1229 != v1230 {
		goto L373
	} else {
		goto L390
	}
L390:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v1234 = F_exprTypmod(m, v1220)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	if v1233 != v1234 {
		goto L373
	} else {
		goto L392
	}
L392:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1238 = F_exprCollation(m, v1220)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	if v1237 == v1238 {
		v1755 = v1220
		goto L3
	} else {
		goto L394
	}
L394:
	;
	goto L373
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+4)) = v1186
	*(*int32)(unsafe.Add(mBase, uint32(v1245))) = int32(25)
	v1250 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1245)+8)) = uint16(v1250)
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+12)) = v1252
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+16)) = v1254
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+20)) = v1256
	if v1186 == int32(0) {
		v1755 = v1245
		goto L3
	} else {
		goto L396
	}
L396:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	if v1260 != int32(7) {
		v1755 = v1245
		goto L3
	} else {
		goto L397
	}
L397:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	v1264 = F_rowtype_field_matches(m, v1263, v1250, v1252, v1254, v1256)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	if v1264 == int32(0) {
		v1755 = v1245
		goto L3
	} else {
		goto L399
	}
L399:
	;
	v1268 = F_exprType(m, v1245)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	v1270 = F_exprTypmod(m, v1245)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	v1272 = F_exprCollation(m, v1245)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v1274 = F_evaluate_expr(m, v1245, v1268, v1270, v1272)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	v1755 = v1274
	goto L3
L404:
	;
	v1281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)))
	if base.B2i32(v1277 == int32(0))|base.B2i32(v1281 != int32(1)) != 0 {
		goto L14
	} else {
		goto L405
	}
L405:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	if v1285 != int32(36) {
		goto L13
	} else {
		goto L406
	}
L406:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+4))
	if v1288 != 0 {
		goto L408
	} else {
		goto L409
	}
L407:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+4))
	if v1380 == int32(1) {
		goto L431
	} else {
		goto L432
	}
L408:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1288)+4))
	if int32(0) < v1289 {
		goto L411
	} else {
		goto L412
	}
L409:
	;
	goto L410
L410:
	;
	v1378 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L1
	} else {
		goto L430
	}
L411:
	;
	v1294 = int32(0)
	v1299 = v3
	goto L414
L412:
	;
	v1354 = v3
	goto L413
L413:
	;
	if v1354 != 0 {
		goto L407
	} else {
		goto L429
	}
L414:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1288)+12))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1307+v1294<<(uint(int32(2))%32))))
	if v1311 == int32(0) {
		goto L419
	} else {
		goto L420
	}
L415:
	;
	v1354 = v1343
	goto L413
L416:
	;
	v1345 = v1294 + int32(1)
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1288)+4))
	if v1345 < v1346 {
		v1294 = v1345
		v1299 = v1343
		goto L414
	} else {
		goto L428
	}
L417:
	;
	v1337 = int32(0)
	v1339 = F_makeBoolConst(m, v1337, v1337)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L1
	} else {
		goto L427
	}
L418:
	;
	if v1317 != 0 {
		v1343 = v1299
		goto L416
	} else {
		goto L426
	}
L419:
	;
	v1324 = F_palloc0(m, int32(20))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L1
	} else {
		goto L424
	}
L420:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1311)))
	if v1314 != int32(7) {
		goto L419
	} else {
		goto L421
	}
L421:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+24)))
	if v1318 != int32(1) {
		goto L418
	} else {
		goto L422
	}
L422:
	;
	if v1317 != int32(1) {
		v1343 = v1299
		goto L416
	} else {
		goto L423
	}
L423:
	;
	goto L417
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1324)+4)) = v1311
	*(*int32)(unsafe.Add(mBase, uint32(v1324))) = int32(52)
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1330 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1324)+12)) = uint8(v1330)
	*(*int32)(unsafe.Add(mBase, uint32(v1324)+8)) = v1329
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1324)+16)) = v1333
	v1335 = F_lappend(m, v1299, v1324)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	v1343 = v1335
	goto L416
L426:
	;
	goto L417
L427:
	;
	v1755 = v1339
	goto L3
L428:
	;
	goto L415
L429:
	;
	goto L410
L430:
	;
	v1755 = v1378
	goto L3
L431:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+12))
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1383)))
	v1755 = v1384
	goto L3
L432:
	;
	goto L433
L433:
	;
	v1385 = F_make_andclause(m, v1354)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	v1755 = v1385
	goto L3
L435:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1393)+32)) = v1395
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1393)+24)) = v1397
	v1755 = v1393
	goto L3
L436:
	;
	v1404 = F_palloc0(m, int32(20))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L1
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1404))) = int32(30)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1404)+8)) = v1408
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1404)+12)) = v1410
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1404)+16)) = v1412
	if v1401 == int32(0) {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1404)+4)) = int32(0)
	v1755 = v1404
	goto L3
L439:
	;
	goto L440
L440:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1401)))
	if v1418 != int32(30) {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1430)))
	if v1431 != int32(7) {
		v1755 = v1404
		goto L3
	} else {
		goto L449
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1404)+4)) = v1401
	v1430 = v1401
	goto L441
L443:
	;
	goto L444
L444:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+4))
	if v1410 == int32(2) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1404)+12)) = v1425
	goto L447
L446:
	;
	goto L447
L447:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1404)+4)) = v1422
	if v1422 == int32(0) {
		v1755 = v1404
		goto L3
	} else {
		goto L448
	}
L448:
	;
	v1430 = v1422
	goto L441
L449:
	;
	v1434 = F_exprType(m, v1404)
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	v1436 = F_exprTypmod(m, v1404)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	v1438 = F_exprCollation(m, v1404)
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	v1440 = F_evaluate_expr(m, v1404, v1434, v1436, v1438)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	v1755 = v1440
	goto L3
L454:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	if v1445 != 0 {
		v25 = v1445
		goto L17
	} else {
		goto L456
	}
L456:
	;
	v1755 = v3
	goto L3
L457:
	;
	v1755 = v1450
	goto L3
L458:
	;
	v1455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1455 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	v1496 = F_palloc0(m, int32(28))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L1
	} else {
		goto L475
	}
L460:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1459 = F_DomainHasConstraints(m, v1458)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L463
	}
L461:
	;
	goto L462
L462:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1461 == int32(0) {
		goto L465
	} else {
		goto L466
	}
L463:
	;
	if v1459 != 0 {
		goto L459
	} else {
		goto L464
	}
L464:
	;
	goto L462
L465:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v1493 = F_applyRelabelType(m, v1453, v1487, v1488, v1489, v1490, v1491, int32(1))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L1
	} else {
		goto L474
	}
L466:
	;
	v1464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1464 != 0 {
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if base.Ui32(int32(_a_F_eval_const_expressions_mutator_4)) <= base.Ui32(v1465) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v1469 = F_palloc0(m, int32(12))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	goto L465
L471:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1469))) = int64(352187318651)
	v1475 = F_GetSysCacheHashValue(m, int32(82), v1465, int32(0))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L1
	} else {
		goto L472
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1469)+8)) = v1475
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+8))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+60))
	v1480 = F_lappend(m, v1479, v1469)
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L1
	} else {
		goto L473
	}
L473:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1482)+60)) = v1480
	goto L470
L474:
	;
	v1755 = v1493
	goto L3
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+4)) = v1453
	*(*int32)(unsafe.Add(mBase, uint32(v1496))) = int32(55)
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+8)) = v1501
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+12)) = v1503
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+16)) = v1505
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+20)) = v1507
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+24)) = v1509
	v1755 = v1496
	goto L3
L476:
	;
	v1569 = F_palloc0(m, int32(16))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L1
	} else {
		goto L497
	}
L477:
	;
	if v1512 == int32(0) {
		goto L476
	} else {
		goto L478
	}
L478:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1512)))
	if v1516 != int32(7) {
		goto L476
	} else {
		goto L479
	}
L479:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	switch v1519 {
	case 0:
		goto L488
	case 1:
		goto L487
	case 2:
		goto L486
	case 3:
		goto L485
	case 4:
		goto L484
	case 5:
		goto L483
	default:
		goto L482
	}
L480:
	;
	v1566 = F_makeBoolConst(m, v1562&int32(1), int32(0))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L1
	} else {
		goto L496
	}
L481:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1512)+20))
	v1562 = base.B2i32(v1559 != int32(0))
	goto L480
L482:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L493
	}
L483:
	;
	v1540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512)+24)))
	v1562 = v1540 ^ int32(1)
	goto L480
L484:
	;
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512)+24)))
	v1562 = v1539
	goto L480
L485:
	;
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512)+24)))
	if v1535 != 0 {
		v1562 = int32(1)
		goto L480
	} else {
		goto L492
	}
L486:
	;
	v1530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512)+24)))
	if v1530 != 0 {
		v1562 = int32(0)
		goto L480
	} else {
		goto L491
	}
L487:
	;
	v1525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512)+24)))
	if v1525 != 0 {
		v1562 = int32(1)
		goto L480
	} else {
		goto L490
	}
L488:
	;
	v1520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512)+24)))
	if v1520 == int32(0) {
		goto L481
	} else {
		goto L489
	}
L489:
	;
	v1562 = int32(0)
	goto L480
L490:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1512)+20))
	v1562 = base.B2i32(v1526 == int32(0))
	goto L480
L491:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1512)+20))
	v1562 = base.B2i32(v1531 == int32(0))
	goto L480
L492:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1512)+20))
	v1562 = base.B2i32(v1536 != int32(0))
	goto L480
L493:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v1547
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_5), v17-int32(-64))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(3593), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L496:
	;
	v1755 = v1566
	goto L3
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1569)+4)) = v1512
	*(*int32)(unsafe.Add(mBase, uint32(v1569))) = int32(53)
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1569)+8)) = v1574
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1569)+12)) = v1576
	v1755 = v1569
	goto L3
L498:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	if v1581 != int32(7) {
		goto L13
	} else {
		goto L499
	}
L499:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	switch v1584 {
	case 0:
		goto L501
	case 1:
		goto L503
	default:
		goto L502
	}
L500:
	;
	v1609 = F_makeBoolConst(m, v1605&int32(1), int32(0))
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L1
	} else {
		goto L507
	}
L501:
	;
	v1604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1277)+24)))
	v1605 = v1604
	goto L500
L502:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L1
	} else {
		goto L504
	}
L503:
	;
	v1585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1277)+24)))
	v1605 = v1585 ^ int32(1)
	goto L500
L504:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v1592
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_6), v17+int32(48))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(3532), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L507:
	;
	v1755 = v1609
	goto L3
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1612)+4)) = v1277
	*(*int32)(unsafe.Add(mBase, uint32(v1612))) = int32(52)
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1612)+8)) = v1617
	v1619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1612)+12)) = uint8(v1619)
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1612)+16)) = v1621
	v1755 = v1612
	goto L3
L509:
	;
	v1627 = v360
	v1628 = v355
	v1630 = v356
	v1633 = v357
	goto L11
L510:
	;
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1643)+24)))
	v1655 = (v1647 | v1628) & int32(1)
	v1657 = v1630 & v1647
	v1660 = v1633
	goto L10
L511:
	;
	if v1657&int32(1) == int32(0) {
		goto L8
	} else {
		goto L512
	}
L512:
	;
	goto L9
L513:
	;
	v1755 = v1686
	goto L3
L514:
	;
	v1692 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L1
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	F_set_opfuncid(m, v25)
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L1
	} else {
		goto L518
	}
L517:
	;
	v1755 = v1692
	goto L3
L518:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v1703 = int32(0)
	v1706 = F_simplify_function(m, v1696, v1697, int32(-1), v1699, v1700, v17+int32(100), v1703, v1703, v1703, l1)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	if v1706 != 0 {
		goto L6
	} else {
		goto L520
	}
L520:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v1716 = v1708
	goto L7
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1724))) = int32(18)
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1724)+4)) = v1728
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1724)+8)) = v1730
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1724)+12)) = v1732
	v1734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1724)+16)) = uint8(v1734)
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1724)+20)) = v1736
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1724)+28)) = v1716
	*(*int32)(unsafe.Add(mBase, uint32(v1724)+24)) = v1738
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1724)+32)) = v1741
	v1755 = v1724
	goto L3
L522:
	;
	v1755 = v1751
	goto L3
}
func F_exec_dynquery_with_params(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v14 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		v22 = F_AllocSetContextCreateInternal(m, v17, int32(_a_F_exec_dynquery_with_params_0), int32(0), int32(_a_F_exec_dynquery_with_params_1), int32(_a_F_exec_dynquery_with_params_2))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v22
			v27 = v22
			v34 = F_exec_eval_expr(m, l0, l1, v12+int32(30), v12+int32(24), v12+int32(20))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)))
				if v36 != int32(1) {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					v40 = int32(_a_F_exec_dynquery_with_params_3)
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0]))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v44
					F_getTypeOutputInfo(m, v39, v12+int32(8), v12+int32(31))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						v53 = F_OidOutputFunctionCall(m, v52, v34)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v41
							v57 = F_MemoryContextStrdup(m, v27, v53)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
								if v59 != 0 {
									F_SPI_freetuptable(m, v59)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
										if v64 != 0 {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
											F_MemoryContextReset(m, v65)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
												v70 = F_exec_eval_using_params(m, l0, l2)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
													v76 = m.G0
													v78 = v76 - int32(48)
													m.G0 = v78
													v80 = int32(0)
													v83 = v12 + int32(8)
													if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
														v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														if v90 == int32(0) {
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1545), int32(_a_F_exec_dynquery_with_params_6))
																	mBase = m.M
																	v173 = m.ExcPending
																	if v173 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
															v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
															*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
															v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
															v103 = int32(0)
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
															v105 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
															*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
															*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
															*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
															*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
															*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
															if v119 != 0 {
																v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
																*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
																v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
																*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
															} else {
															}
															v125 = v78 + int32(8)
															F__SPI_prepare_plan(m, v57, v125)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
																v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
																v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
																	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
																	*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
																	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
																	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
																	F_MemoryContextReset(m, v139)
																	mBase = m.M
																	v141 = m.ExcPending
																	if v141 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v78 + int32(48)
																		if v130 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																			mBase = m.M
																			v201 = m.ExcPending
																			if v201 != 0 {
																				return int32(0)
																			} else {
																				v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																				v204 = F_SPI_result_code_string(m, v203)
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																					F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																					mBase = m.M
																					v210 = m.ExcPending
																					if v210 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																						mBase = m.M
																						v215 = m.ExcPending
																						if v215 != 0 {
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
																			F_MemoryContextReset(m, v27)
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v12 + int32(32)
																				return v130
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1541), int32(_a_F_exec_dynquery_with_params_6))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
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
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
											v70 = F_exec_eval_using_params(m, l0, l2)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
												v76 = m.G0
												v78 = v76 - int32(48)
												m.G0 = v78
												v80 = int32(0)
												v83 = v12 + int32(8)
												if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													if v90 == int32(0) {
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1545), int32(_a_F_exec_dynquery_with_params_6))
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
														v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
														v103 = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
														v105 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
														*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														if v119 != 0 {
															v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
														} else {
														}
														v125 = v78 + int32(8)
														F__SPI_prepare_plan(m, v57, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
															v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
															v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
																*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
																*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
																F_MemoryContextReset(m, v139)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v78 + int32(48)
																	if v130 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																			v204 = F_SPI_result_code_string(m, v203)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																				F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																				mBase = m.M
																				v210 = m.ExcPending
																				if v210 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
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
																		F_MemoryContextReset(m, v27)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v12 + int32(32)
																			return v130
																		}
																	}
																}
															}
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1541), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
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
									*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									if v64 != 0 {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
										F_MemoryContextReset(m, v65)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
											v70 = F_exec_eval_using_params(m, l0, l2)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
												v76 = m.G0
												v78 = v76 - int32(48)
												m.G0 = v78
												v80 = int32(0)
												v83 = v12 + int32(8)
												if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													if v90 == int32(0) {
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1545), int32(_a_F_exec_dynquery_with_params_6))
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
														v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
														v103 = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
														v105 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
														*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														if v119 != 0 {
															v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
														} else {
														}
														v125 = v78 + int32(8)
														F__SPI_prepare_plan(m, v57, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
															v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
															v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
																*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
																*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
																F_MemoryContextReset(m, v139)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v78 + int32(48)
																	if v130 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																			v204 = F_SPI_result_code_string(m, v203)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																				F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																				mBase = m.M
																				v210 = m.ExcPending
																				if v210 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
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
																		F_MemoryContextReset(m, v27)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v12 + int32(32)
																			return v130
																		}
																	}
																}
															}
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1541), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
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
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
										v70 = F_exec_eval_using_params(m, l0, l2)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
											v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
											v76 = m.G0
											v78 = v76 - int32(48)
											m.G0 = v78
											v80 = int32(0)
											v83 = v12 + int32(8)
											if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
												if v90 == int32(0) {
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1545), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
													v103 = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
													v105 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
													*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
													*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
													if v119 != 0 {
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
													} else {
													}
													v125 = v78 + int32(8)
													F__SPI_prepare_plan(m, v57, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
														v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
															v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
															*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
															v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
															F_MemoryContextReset(m, v139)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																m.G0 = v78 + int32(48)
																if v130 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																		v204 = F_SPI_result_code_string(m, v203)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																			F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
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
																	F_MemoryContextReset(m, v27)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v12 + int32(32)
																		return v130
																	}
																}
															}
														}
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1541), int32(_a_F_exec_dynquery_with_params_6))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
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
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v188 = m.ExcPending
						if v188 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_exec_dynquery_with_params_13), int32(0))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_14), int32(_a_F_exec_dynquery_with_params_11))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
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
		v27 = v14
		v34 = F_exec_eval_expr(m, l0, l1, v12+int32(30), v12+int32(24), v12+int32(20))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)))
			if v36 != int32(1) {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
				v40 = int32(_a_F_exec_dynquery_with_params_3)
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0]))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
				*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v44
				F_getTypeOutputInfo(m, v39, v12+int32(8), v12+int32(31))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					v53 = F_OidOutputFunctionCall(m, v52, v34)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v41
						v57 = F_MemoryContextStrdup(m, v27, v53)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							if v59 != 0 {
								F_SPI_freetuptable(m, v59)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									if v64 != 0 {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
										F_MemoryContextReset(m, v65)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
											v70 = F_exec_eval_using_params(m, l0, l2)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
												v76 = m.G0
												v78 = v76 - int32(48)
												m.G0 = v78
												v80 = int32(0)
												v83 = v12 + int32(8)
												if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													if v90 == int32(0) {
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1545), int32(_a_F_exec_dynquery_with_params_6))
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
														v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
														v103 = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
														v105 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
														*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														if v119 != 0 {
															v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
														} else {
														}
														v125 = v78 + int32(8)
														F__SPI_prepare_plan(m, v57, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
															v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
															v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
																*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
																*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
																F_MemoryContextReset(m, v139)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v78 + int32(48)
																	if v130 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																			v204 = F_SPI_result_code_string(m, v203)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																				F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																				mBase = m.M
																				v210 = m.ExcPending
																				if v210 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
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
																		F_MemoryContextReset(m, v27)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v12 + int32(32)
																			return v130
																		}
																	}
																}
															}
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1541), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
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
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
										v70 = F_exec_eval_using_params(m, l0, l2)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
											v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
											v76 = m.G0
											v78 = v76 - int32(48)
											m.G0 = v78
											v80 = int32(0)
											v83 = v12 + int32(8)
											if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
												if v90 == int32(0) {
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1545), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
													v103 = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
													v105 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
													*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
													*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
													if v119 != 0 {
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
													} else {
													}
													v125 = v78 + int32(8)
													F__SPI_prepare_plan(m, v57, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
														v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
															v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
															*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
															v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
															F_MemoryContextReset(m, v139)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																m.G0 = v78 + int32(48)
																if v130 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																		v204 = F_SPI_result_code_string(m, v203)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																			F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
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
																	F_MemoryContextReset(m, v27)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v12 + int32(32)
																		return v130
																	}
																}
															}
														}
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1541), int32(_a_F_exec_dynquery_with_params_6))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
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
								*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								if v64 != 0 {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
									F_MemoryContextReset(m, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
										v70 = F_exec_eval_using_params(m, l0, l2)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
											v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
											v76 = m.G0
											v78 = v76 - int32(48)
											m.G0 = v78
											v80 = int32(0)
											v83 = v12 + int32(8)
											if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
												if v90 == int32(0) {
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1545), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
													v103 = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
													v105 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
													*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
													*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
													if v119 != 0 {
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
													} else {
													}
													v125 = v78 + int32(8)
													F__SPI_prepare_plan(m, v57, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
														v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
															v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
															*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
															v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
															F_MemoryContextReset(m, v139)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																m.G0 = v78 + int32(48)
																if v130 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																		v204 = F_SPI_result_code_string(m, v203)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																			F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
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
																	F_MemoryContextReset(m, v27)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v12 + int32(32)
																		return v130
																	}
																}
															}
														}
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1541), int32(_a_F_exec_dynquery_with_params_6))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
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
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
									v70 = F_exec_eval_using_params(m, l0, l2)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
										v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
										v76 = m.G0
										v78 = v76 - int32(48)
										m.G0 = v78
										v80 = int32(0)
										v83 = v12 + int32(8)
										if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
											v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
											if v90 == int32(0) {
												*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1545), int32(_a_F_exec_dynquery_with_params_6))
														mBase = m.M
														v173 = m.ExcPending
														if v173 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
												*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
												*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
												v103 = int32(0)
												*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
												v105 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
												*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
												*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
												*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
												*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
												*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
												if v119 != 0 {
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
												} else {
												}
												v125 = v78 + int32(8)
												F__SPI_prepare_plan(m, v57, v125)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
													v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
													v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return int32(0)
													} else {
														v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
														*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
														F_MemoryContextReset(m, v139)
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int32(0)
														} else {
															m.G0 = v78 + int32(48)
															if v130 == int32(0) {
																F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																	v204 = F_SPI_result_code_string(m, v203)
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																		F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																		mBase = m.M
																		v210 = m.ExcPending
																		if v210 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																			mBase = m.M
																			v215 = m.ExcPending
																			if v215 != 0 {
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
																F_MemoryContextReset(m, v27)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v12 + int32(32)
																	return v130
																}
															}
														}
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
												mBase = m.M
												v152 = m.ExcPending
												if v152 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1541), int32(_a_F_exec_dynquery_with_params_6))
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
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
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67108994))
					mBase = m.M
					v188 = m.ExcPending
					if v188 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_exec_dynquery_with_params_13), int32(0))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_14), int32(_a_F_exec_dynquery_with_params_11))
							mBase = m.M
							v197 = m.ExcPending
							if v197 != 0 {
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
func F_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v6 = int32(0)
	F_check_stack_depth(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v13 == int32(2) {
		v76 = l0
		v82 = v6
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return (v91 ^ v92) & int32(1)
L4:
	;
	v84 = m.T0[l4].(func(*base.Module, int32, int32, int32) int32)(m, l1, v76, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L5:
	;
	v16 = l0
	v19 = l3
	v22 = v6
	goto L6
L6:
	;
	v26 = v16
	goto L8
L7:
	;
	v76 = v72
	v82 = v22
	goto L4
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	switch v34 - int32(33) {
	case 0:
		goto L13
	default:
		goto L11
	case 5:
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v59 = int32(1)
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+2)))
	v66 = F_execute(m, v26+v60<<(uint(int32(3))%32), l1, l2, v19&v59, l4)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L19
	}
L12:
	;
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+2)))
	v56 = F_execute(m, v26+v52<<(uint(int32(3))%32), l1, l2, v19&int32(1), l4)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	v37 = int32(1)
	if v19&v37 == int32(0) {
		v91 = v37
		v92 = v22
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v42 = int32(1)
	v44 = v22 ^ v42
	F_check_stack_depth(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v48 = v26 - int32(8)
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48))))
	if v49 != int32(2) {
		v16 = v48
		v19 = v42
		v22 = v44
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v76 = v48
	v82 = v44
	goto L4
L17:
	;
	if v56 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v91 = int32(0)
	v92 = v22
	goto L3
L19:
	;
	if v66 != 0 {
		v91 = v59
		v92 = v22
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	v72 = v26 - int32(8)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72))))
	if v73 != int32(2) {
		v26 = v72
		goto L8
	} else {
		goto L22
	}
L22:
	;
	goto L9
L23:
	;
	v91 = v84
	v92 = v82
	goto L3
}
func F_existsTimeLineHistory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v3 = m.G0
	v5 = v3 - int32(1136)
	m.G0 = v5
	if l0 != int32(1) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_existsTimeLineHistory[0])))
		if v10 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
			v15 = v5 + int32(48)
			v20 = F_pg_snprintf(m, v15, int32(64), int32(_a_F_existsTimeLineHistory_0), v5+int32(16))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v29 = F_RestoreArchivedFile(m, v5+int32(112), v15, int32(_a_F_existsTimeLineHistory_1), int64(0), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v44 = F_AllocateFile(m, v5+int32(112), int32(_a_F_existsTimeLineHistory_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						if v44 != 0 {
							v46 = F_FreeFile(m, v44)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v56 = int32(1)
								m.G0 = v5 + int32(1136)
								return v56
							}
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_existsTimeLineHistory[1]))
							if v50 != int32(44) {
								F_errstart_cold(m, int32(22), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(112)
										F_errmsg(m, int32(_a_F_existsTimeLineHistory_3), v5)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_existsTimeLineHistory_4), int32(251), int32(_a_F_existsTimeLineHistory_5))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
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
								v56 = int32(0)
								m.G0 = v5 + int32(1136)
								return v56
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = l0
			v38 = F_pg_snprintf(m, v5+int32(112), int32(1024), int32(_a_F_existsTimeLineHistory_6), v5+int32(32))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v44 = F_AllocateFile(m, v5+int32(112), int32(_a_F_existsTimeLineHistory_2))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					if v44 != 0 {
						v46 = F_FreeFile(m, v44)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v56 = int32(1)
							m.G0 = v5 + int32(1136)
							return v56
						}
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, _c_F_existsTimeLineHistory[1]))
						if v50 != int32(44) {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(112)
									F_errmsg(m, int32(_a_F_existsTimeLineHistory_3), v5)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_existsTimeLineHistory_4), int32(251), int32(_a_F_existsTimeLineHistory_5))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
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
							v56 = int32(0)
							m.G0 = v5 + int32(1136)
							return v56
						}
					}
				}
			}
		}
	} else {
		v56 = int32(0)
		m.G0 = v5 + int32(1136)
		return v56
	}
}
func F_expand_partitioned_rtentry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v93 int32
	_ = v93
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
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
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
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
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
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
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v407 int32
	_ = v407
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v453 int32
	_ = v453
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v518 int32
	_ = v518
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	F_check_stack_depth(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
	v28 = F_PartitionDirectoryLookup(m, v27, l4)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+372)))
	if v30 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = F_has_partition_attrs(m, l4, l5, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v37 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+372)) = uint8(v34)
	goto L6
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L140
	}
L9:
	;
	m.G0 = v22 + int32(16)
	return
L10:
	;
	v40 = m.G0
	v42 = v40 + int32(-64)
	m.G0 = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if v44 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	m.G0 = v42 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+256)) = v123
	v130 = int32(0)
	if v123 == v130 {
		goto L34
	} else {
		goto L35
	}
L12:
	;
	v123 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_expand_partitioned_rtentry[0])))
	if v49 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+60)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v42)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v42)+44)) = l1
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
	if v69 == int32(-1) {
		v77 = v52
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v52 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v54 = int32(0)
	v58 = F_bms_add_range(m, v54, v54, v44-int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v123 = v58
	goto L11
L21:
	;
	v81 = F_gen_partprune_steps_internal(m, v40+int32(-20), v77)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L25
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	if v72 == int32(0) {
		v77 = v52
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v75 = F_list_concat_copy(m, v52, v72)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v77 = v75
	goto L21
L25:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+59)))
	if v83 != 0 {
		v123 = v60
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	if v84 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v87 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	v92 = F_bms_add_range(m, v87, v87, v89-int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v95)
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v111 = F_palloc0(m, v97*v107*int32(28))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v123 = v92
	goto L11
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v42)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+40)) = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_expand_partitioned_rtentry[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = v119
	v121 = F_get_matching_partitions(m, v42, v84)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v123 = v121
	goto L11
L33:
	;
	if int32(0) < v165 {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v165 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v137 = int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v138 <= v137 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v141 = v137
	goto L39
L38:
	;
	v141 = v138
	goto L39
L39:
	;
	v145 = int32(0)
	v147 = v130
	goto L40
L40:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v123+int32(8)+v145<<(uint(int32(2))%32))))
	if v153 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v165 = v156
	goto L33
L42:
	;
	v156 = v147 + base.I32_popcnt(v153)
	goto L44
L43:
	;
	v156 = v147
	goto L44
L44:
	;
	v158 = v145 + int32(1)
	if v158 != v141 {
		v145 = v158
		v147 = v156
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	F_expand_planner_arrays(m, l0, v165)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	v173 = F_palloc0(m, v170<<(uint(int32(2))%32))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+252)) = v173
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v176 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if v233 < int32(0) {
		goto L9
	} else {
		goto L62
	}
L52:
	;
	v233 = base.I32_ctz(v219) | v220<<(uint(int32(5))%32)
	goto L51
L53:
	;
	v233 = int32(-2)
	goto L51
L54:
	;
	v186 = base.I32_div_s(int32(0), int32(32))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v187 <= v186 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v190 = v176 + int32(8)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190+v186<<(uint(int32(2))%32))))
	v197 = v194 & int32(-1)
	if v197 != 0 {
		v219 = v197
		v220 = v186
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v199 = v186 + int32(1)
	if v199 == v187 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v202 = v199
	goto L58
L58:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v190+v202<<(uint(int32(2))%32))))
	if v209 != 0 {
		v219 = v209
		v220 = v202
		goto L52
	} else {
		goto L60
	}
L59:
	;
	goto L53
L60:
	;
	v211 = v202 + int32(1)
	if v211 != v187 {
		v202 = v211
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v246 = v233
	goto L63
L63:
	;
	v256 = v246 << (uint(int32(2)) % 32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256+v257)))
	v260 = F_try_table_open(m, v259, l7)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L9
L65:
	;
	if v453 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L66:
	;
	if v260 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	v265 = F_bms_del_member(m, v264, v246)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+48))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+118)))
	if v269 == int32(116) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+256)) = v265
	v453 = v265
	goto L65
L71:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+24)))
	if v272 == int32(0) {
		goto L8
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_expand_single_inheritance_child(m, l0, l2, l3, l4, l6, v260, v22+int32(12), v22+int32(8))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v282 = F_build_simple_rel(m, l0, v281, l1)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	*(*int32)(unsafe.Add(mBase, uint32(v284+v256))) = v282
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v289 = F_bms_add_members(m, v287, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+260)) = v289
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v260)+48))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+119)))
	if v293 == int32(112) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v296+v281<<(uint(int32(2))%32))))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+20))
	v302 = int32(0)
	v305 = F_bms_is_member(m, int32(1), l5)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_relation_close(m, v260, int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L127
	}
L81:
	;
	if v305 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v309 = F_bms_add_member(m, int32(0), int32(1))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	v311 = v302
	goto L84
L84:
	;
	v313 = F_bms_is_member(m, int32(2), l5)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	v311 = v309
	goto L84
L86:
	;
	if v313 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v316 = F_bms_add_member(m, v311, int32(2))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	v318 = v311
	goto L89
L89:
	;
	v320 = F_bms_is_member(m, int32(3), l5)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	v318 = v316
	goto L89
L91:
	;
	if v320 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v323 = F_bms_add_member(m, v318, int32(3))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	v325 = v318
	goto L94
L94:
	;
	v327 = F_bms_is_member(m, int32(4), l5)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	v325 = v323
	goto L94
L96:
	;
	if v327 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v330 = F_bms_add_member(m, v325, int32(4))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	v332 = v325
	goto L99
L99:
	;
	v334 = F_bms_is_member(m, int32(5), l5)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L101
	}
L100:
	;
	v332 = v330
	goto L99
L101:
	;
	if v334 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v337 = F_bms_add_member(m, v332, int32(5))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	v339 = v332
	goto L104
L104:
	;
	v341 = F_bms_is_member(m, int32(6), l5)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	v339 = v337
	goto L104
L106:
	;
	if v341 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v344 = F_bms_add_member(m, v339, int32(6))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	v346 = v339
	goto L109
L109:
	;
	v348 = F_bms_is_member(m, int32(7), l5)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L111
	}
L110:
	;
	v346 = v344
	goto L109
L111:
	;
	if v301 == int32(0) {
		v407 = v346
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	F_expand_partitioned_rtentry(m, l0, v282, v418, v281, v260, v407, l6, l7)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L126
	}
L113:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v352 <= int32(0) {
		v407 = v346
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v363 = v346
	v368 = v302
	goto L115
L115:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374+v368<<(uint(int32(2))%32))))
	if v378 == int32(0) {
		v394 = v363
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v407 = v394
	goto L112
L117:
	;
	v396 = v368 + int32(1)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v396 < v397 {
		v363 = v394
		v368 = v396
		goto L115
	} else {
		goto L125
	}
L118:
	;
	if v348 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v385 = F_bms_is_member(m, v368+int32(8), l5)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v389 = int32(*(*int16)(unsafe.Add(mBase, uint32(v378)+8)))
	v392 = F_bms_add_member(m, v363, v389+int32(7))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	if v385 == int32(0) {
		v394 = v363
		goto L117
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v394 = v392
	goto L117
L125:
	;
	goto L116
L126:
	;
	goto L80
L127:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	v453 = v443
	goto L65
L128:
	;
	if int32(0) <= v518 {
		v246 = v518
		goto L63
	} else {
		goto L139
	}
L129:
	;
	v518 = base.I32_ctz(v504) | v505<<(uint(int32(5))%32)
	goto L128
L130:
	;
	v518 = int32(-2)
	goto L128
L131:
	;
	v469 = v246 + int32(1)
	v471 = base.I32_div_s(v469, int32(32))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if v472 <= v471 {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v475 = v453 + int32(8)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v475+v471<<(uint(int32(2))%32))))
	v482 = v479 & (int32(-1) << (uint(v469) % 32))
	if v482 != 0 {
		v504 = v482
		v505 = v471
		goto L129
	} else {
		goto L133
	}
L133:
	;
	v484 = v471 + int32(1)
	if v484 == v472 {
		goto L130
	} else {
		goto L134
	}
L134:
	;
	v487 = v484
	goto L135
L135:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v475+v487<<(uint(int32(2))%32))))
	if v494 != 0 {
		v504 = v494
		v505 = v487
		goto L129
	} else {
		goto L137
	}
L136:
	;
	goto L130
L137:
	;
	v496 = v487 + int32(1)
	if v496 != v472 {
		v487 = v496
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	goto L64
L140:
	;
	F_errmsg_internal(m, int32(_a_F_expand_partitioned_rtentry_0), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_expand_partitioned_rtentry_1), int32(406), int32(_a_F_expand_partitioned_rtentry_2))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_expm1(m *base.Module, l0 float64) float64 {
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v90 float64
	_ = v90
	var v93 float64
	_ = v93
	var v99 float64
	_ = v99
	var v109 float64
	_ = v109
	var v126 float64
	_ = v126
	var v136 float64
	_ = v136
	var v141 float64
	_ = v141
	var v148 float64
	_ = v148
	var v152 float64
	_ = v152
	var v158 float64
	_ = v158
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	v8 = base.I64_reinterpret_f64(l0)
	v13 = base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1078159482)) <= base.Ui32(v13) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
			v170 = l0
			return v170
		} else {
			if v8 < int64(0) {
				return float64(-1)
			} else {
				if base.F64_gt(l0, float64(709.782712893384)) == int32(0) {
					v51 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(l0, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), l0)))
					v52 = base.F64_convert_i32_s(v51)
					v58 = v51
					v59 = base.F64_mul(v52, float64(1.9082149292705877e-10))
					v61 = base.F64_add(l0, base.F64_mul(v52, float64(-0.6931471803691238)))
					v62 = base.F64_sub(v61, v59)
					v68 = v62
					v69 = v58
					v70 = base.F64_sub(base.F64_sub(v61, v62), v59)
					v73 = base.F64_mul(v68, float64(0.5))
					v74 = base.F64_mul(v68, v73)
					v90 = base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v93 = base.F64_sub(float64(3), base.F64_mul(v90, v73))
					v99 = base.F64_mul(v74, base.F64_div(base.F64_sub(v90, v93), base.F64_sub(float64(6), base.F64_mul(v68, v93))))
					if v69 == int32(0) {
						return base.F64_sub(v68, base.F64_sub(base.F64_mul(v68, v99), v74))
					} else {
						v109 = base.F64_sub(base.F64_sub(base.F64_mul(v68, base.F64_sub(v99, v70)), v70), v74)
						switch v69 + int32(1) {
						case 0:
							return base.F64_add(base.F64_mul(base.F64_sub(v68, v109), float64(0.5)), float64(-0.5))
						default:
							v136 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v69+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v69) {
								v141 = base.F64_add(base.F64_sub(v68, v109), float64(1))
								if v69 == int32(1024) {
									v148 = base.F64_mul(base.F64_add(v141, v141), float64(8.98846567431158e+307))
								} else {
									v148 = base.F64_mul(v141, v136)
								}
								return base.F64_add(v148, float64(-1))
							} else {
								v152 = float64(1)
								v158 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v69) << (uint(int64(52)) % 64))
								if base.Ui32(v69) <= base.Ui32(int32(19)) {
									v168 = base.F64_add(base.F64_sub(v152, v158), base.F64_sub(v68, v109))
								} else {
									v168 = base.F64_add(base.F64_sub(v68, base.F64_add(v109, v158)), v152)
								}
								v170 = base.F64_mul(v168, v136)
								return v170
							}
						case 2:
							if base.F64_lt(v68, float64(-0.25)) != 0 {
								return base.F64_mul(base.F64_sub(v109, base.F64_add(v68, float64(0.5))), float64(-2))
							} else {
								v126 = base.F64_sub(v68, v109)
								return base.F64_add(base.F64_add(v126, v126), float64(1))
							}
						}
					}
				} else {
					return base.F64_mul(l0, float64(8.98846567431158e+307))
				}
			}
		}
	} else {
		if base.Ui32(v13) < base.Ui32(int32(1071001155)) {
			if base.Ui32(v13) < base.Ui32(int32(1016070144)) {
				v170 = l0
				return v170
			} else {
				v68 = l0
				v69 = int32(0)
				v70 = float64(0)
				v73 = base.F64_mul(v68, float64(0.5))
				v74 = base.F64_mul(v68, v73)
				v90 = base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
				v93 = base.F64_sub(float64(3), base.F64_mul(v90, v73))
				v99 = base.F64_mul(v74, base.F64_div(base.F64_sub(v90, v93), base.F64_sub(float64(6), base.F64_mul(v68, v93))))
				if v69 == int32(0) {
					return base.F64_sub(v68, base.F64_sub(base.F64_mul(v68, v99), v74))
				} else {
					v109 = base.F64_sub(base.F64_sub(base.F64_mul(v68, base.F64_sub(v99, v70)), v70), v74)
					switch v69 + int32(1) {
					case 0:
						return base.F64_add(base.F64_mul(base.F64_sub(v68, v109), float64(0.5)), float64(-0.5))
					default:
						v136 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v69+int32(1023)) << (uint(int64(52)) % 64))
						if base.Ui32(int32(57)) <= base.Ui32(v69) {
							v141 = base.F64_add(base.F64_sub(v68, v109), float64(1))
							if v69 == int32(1024) {
								v148 = base.F64_mul(base.F64_add(v141, v141), float64(8.98846567431158e+307))
							} else {
								v148 = base.F64_mul(v141, v136)
							}
							return base.F64_add(v148, float64(-1))
						} else {
							v152 = float64(1)
							v158 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v69) << (uint(int64(52)) % 64))
							if base.Ui32(v69) <= base.Ui32(int32(19)) {
								v168 = base.F64_add(base.F64_sub(v152, v158), base.F64_sub(v68, v109))
							} else {
								v168 = base.F64_add(base.F64_sub(v68, base.F64_add(v109, v158)), v152)
							}
							v170 = base.F64_mul(v168, v136)
							return v170
						}
					case 2:
						if base.F64_lt(v68, float64(-0.25)) != 0 {
							return base.F64_mul(base.F64_sub(v109, base.F64_add(v68, float64(0.5))), float64(-2))
						} else {
							v126 = base.F64_sub(v68, v109)
							return base.F64_add(base.F64_add(v126, v126), float64(1))
						}
					}
				}
			}
		} else {
			if base.Ui32(int32(1072734897)) < base.Ui32(v13) {
				v51 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(l0, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), l0)))
				v52 = base.F64_convert_i32_s(v51)
				v58 = v51
				v59 = base.F64_mul(v52, float64(1.9082149292705877e-10))
				v61 = base.F64_add(l0, base.F64_mul(v52, float64(-0.6931471803691238)))
			} else {
				if int64(0) <= v8 {
					v58 = int32(1)
					v59 = float64(1.9082149292705877e-10)
					v61 = base.F64_add(l0, float64(-0.6931471803691238))
				} else {
					v58 = int32(-1)
					v59 = float64(-1.9082149292705877e-10)
					v61 = base.F64_add(l0, float64(0.6931471803691238))
				}
			}
			v62 = base.F64_sub(v61, v59)
			v68 = v62
			v69 = v58
			v70 = base.F64_sub(base.F64_sub(v61, v62), v59)
			v73 = base.F64_mul(v68, float64(0.5))
			v74 = base.F64_mul(v68, v73)
			v90 = base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
			v93 = base.F64_sub(float64(3), base.F64_mul(v90, v73))
			v99 = base.F64_mul(v74, base.F64_div(base.F64_sub(v90, v93), base.F64_sub(float64(6), base.F64_mul(v68, v93))))
			if v69 == int32(0) {
				return base.F64_sub(v68, base.F64_sub(base.F64_mul(v68, v99), v74))
			} else {
				v109 = base.F64_sub(base.F64_sub(base.F64_mul(v68, base.F64_sub(v99, v70)), v70), v74)
				switch v69 + int32(1) {
				case 0:
					return base.F64_add(base.F64_mul(base.F64_sub(v68, v109), float64(0.5)), float64(-0.5))
				default:
					v136 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v69+int32(1023)) << (uint(int64(52)) % 64))
					if base.Ui32(int32(57)) <= base.Ui32(v69) {
						v141 = base.F64_add(base.F64_sub(v68, v109), float64(1))
						if v69 == int32(1024) {
							v148 = base.F64_mul(base.F64_add(v141, v141), float64(8.98846567431158e+307))
						} else {
							v148 = base.F64_mul(v141, v136)
						}
						return base.F64_add(v148, float64(-1))
					} else {
						v152 = float64(1)
						v158 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v69) << (uint(int64(52)) % 64))
						if base.Ui32(v69) <= base.Ui32(int32(19)) {
							v168 = base.F64_add(base.F64_sub(v152, v158), base.F64_sub(v68, v109))
						} else {
							v168 = base.F64_add(base.F64_sub(v68, base.F64_add(v109, v158)), v152)
						}
						v170 = base.F64_mul(v168, v136)
						return v170
					}
				case 2:
					if base.F64_lt(v68, float64(-0.25)) != 0 {
						return base.F64_mul(base.F64_sub(v109, base.F64_add(v68, float64(0.5))), float64(-2))
					} else {
						v126 = base.F64_sub(v68, v109)
						return base.F64_add(base.F64_add(v126, v126), float64(1))
					}
				}
			}
		}
	}
}
func F_extra_field_used(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	var v34 int32
	_ = v34
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if l1 == v5 {
		v34 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v34
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v7 {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	case 4:
		goto L4
	default:
		goto L3
	}
L3:
	;
	v20 = l0 + int32(56)
	goto L14
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if l1 == v16 {
		v34 = v4
		goto L1
	} else {
		goto L13
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if l1 != v14 {
		goto L3
	} else {
		goto L12
	}
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if l1 != v12 {
		goto L3
	} else {
		goto L11
	}
L7:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if l1 != v10 {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if l1 != v8 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v34 = v4
	goto L1
L10:
	;
	v34 = v4
	goto L1
L11:
	;
	v34 = v4
	goto L1
L12:
	;
	v34 = v4
	goto L1
L13:
	;
	goto L3
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v24 = int32(0)
	v25 = base.B2i32(v23 != v24)
	if v23 == v24 {
		v34 = v25
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v34 = v25
	goto L1
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	if l1 == v28 {
		v34 = v25
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	if l1 != v30 {
		v20 = v23
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
}
func F_extractModify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_defGetString(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L29
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return v100
L3:
	;
	return int32(0)
L4:
	;
	v13 = int32(_a_F_extractModify_0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_extractModify[0])))
	if base.B2i32(v16 == int32(0))|base.B2i32(v16 != v19) != 0 {
		v37 = v16
		v38 = v19
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v37-v38 == int32(0) {
		v100 = int32(114)
		goto L2
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v22 = v9
	v23 = v13
	goto L8
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v27
		v38 = v26
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v37 = v27
	v38 = v26
	goto L6
L10:
	;
	v30 = int32(1)
	if v27 == v26 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v43 = int32(_a_F_extractModify_1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_extractModify[1])))
	if base.B2i32(v46 == int32(0))|base.B2i32(v46 != v49) != 0 {
		v67 = v46
		v68 = v49
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v67-v68 == int32(0) {
		v100 = int32(115)
		goto L2
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v52 = v9
	v53 = v43
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v57
		v68 = v56
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v67 = v57
	v68 = v56
	goto L14
L18:
	;
	v60 = int32(1)
	if v57 == v56 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v72 = int32(_a_F_extractModify_2)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_extractModify[2])))
	if base.B2i32(v75 == int32(0))|base.B2i32(v75 != v78) != 0 {
		v96 = v75
		v97 = v78
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v96-v97 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v81 = v9
	v82 = v72
	goto L24
L24:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v86 == int32(0) {
		v96 = v86
		v97 = v85
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v86
	v97 = v85
	goto L22
L26:
	;
	v89 = int32(1)
	if v86 == v85 {
		v81 = v81 + v89
		v82 = v82 + v89
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v100 = int32(119)
	goto L2
L29:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v112
	F_errmsg(m, int32(_a_F_extractModify_3), v6)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_extractModify_4), int32(491), int32(_a_F_extractModify_5))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_extract_actual_clauses(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	v3 = int32(0)
	if l0 == v3 {
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v10 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v3
	v17 = v3
	goto L7
L5:
	;
	v45 = v3
	goto L6
L6:
	;
	return v45
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v16<<(uint(int32(2))%32))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
	if v23 != l1 {
		v36 = v17
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v45 = v36
	goto L6
L9:
	;
	v38 = v16 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 < v39 {
		v16 = v38
		v17 = v36
		goto L7
	} else {
		goto L17
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 != int32(7) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = F_lappend(m, v17, v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+24)))
	if v29 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	if v30 != 0 {
		v36 = v17
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	return int32(0)
L16:
	;
	v36 = v31
	goto L9
L17:
	;
	goto L8
}
func F_extract_or_clause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
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
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
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
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
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
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v10 == v3 {
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if int32(0) < v15 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v247
L5:
	;
	v22 = v3
	v25 = v3
	goto L8
L6:
	;
	v236 = v3
	goto L7
L7:
	;
	v240 = int32(0)
	if v236 == v240 {
		v247 = v240
		goto L4
	} else {
		goto L75
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v25<<(uint(int32(2))%32))))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v236 = v227
	goto L7
L10:
	;
	if v203 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L11:
	;
	v136 = int32(0)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+10)))
	if v137 != 0 {
		v247 = v136
		goto L4
	} else {
		goto L47
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v33 != int32(21) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v36 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v37 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L17
L17:
	;
	v42 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v44 <= v42 {
		v203 = v42
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v47 = v42
	v50 = v42
	goto L19
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v47<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	goto L23
L20:
	;
	v203 = v131
	goto L10
L21:
	;
	v133 = v47 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v133 < v134 {
		v47 = v133
		v50 = v131
		goto L19
	} else {
		goto L46
	}
L22:
	;
	v128 = F_lappend(m, v50, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L27
	} else {
		goto L45
	}
L23:
	;
	if v60 != int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v63 = F_extract_or_clause(m, v59, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+10)))
	if v67 != 0 {
		v131 = v50
		goto L21
	} else {
		goto L30
	}
L27:
	;
	return int32(0)
L28:
	;
	if v63 != 0 {
		v127 = v63
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v131 = v50
	goto L21
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v70 = int32(0)
	if base.B2i32(v68 == v70)|base.B2i32(v69 == v70) != 0 {
		v116 = base.B2i32(v68|v69 == v70)
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v116 == int32(0) {
		v131 = v50
		goto L21
	} else {
		goto L42
	}
L32:
	;
	goto L31
L33:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v84 != v85 {
		v116 = int32(0)
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v87 = int32(1)
	if v84 <= v87 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v90 = v87
	goto L37
L36:
	;
	v90 = v84
	goto L37
L37:
	;
	v91 = int32(8)
	v96 = int32(0)
	goto L38
L38:
	;
	v104 = v96 << (uint(int32(2)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v68+v91+v104)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v69+v91+v104)))
	v109 = base.B2i32(v106 == v108)
	if v106 != v108 {
		v116 = v109
		goto L32
	} else {
		goto L40
	}
L39:
	;
	v116 = v109
	goto L32
L40:
	;
	v112 = v96 + int32(1)
	if v112 != v90 {
		v96 = v112
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v124 = F_contain_volatile_functions(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	if v124 != 0 {
		v131 = v50
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v127 = v126
	goto L22
L45:
	;
	v131 = v128
	goto L21
L46:
	;
	goto L20
L47:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v140 = int32(0)
	if base.B2i32(v138 == v140)|base.B2i32(v139 == v140) != 0 {
		v186 = base.B2i32(v138|v139 == v140)
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v186 == int32(0) {
		v247 = v136
		goto L4
	} else {
		goto L59
	}
L49:
	;
	goto L48
L50:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v154 != v155 {
		v186 = int32(0)
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v157 = int32(1)
	if v154 <= v157 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v160 = v157
	goto L54
L53:
	;
	v160 = v154
	goto L54
L54:
	;
	v161 = int32(8)
	v166 = int32(0)
	goto L55
L55:
	;
	v174 = v166 << (uint(int32(2)) % 32)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v138+v161+v174)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v139+v161+v174)))
	v179 = base.B2i32(v176 == v178)
	if v176 != v178 {
		v186 = v179
		goto L49
	} else {
		goto L57
	}
L56:
	;
	v186 = v179
	goto L49
L57:
	;
	v182 = v166 + int32(1)
	if v182 != v160 {
		v166 = v182
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v194 = F_contain_volatile_functions(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L27
	} else {
		goto L60
	}
L60:
	;
	if v194 != 0 {
		v247 = v136
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v198 = F_lappend(m, int32(0), v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L27
	} else {
		goto L62
	}
L62:
	;
	v203 = v198
	goto L10
L63:
	;
	return int32(0)
L64:
	;
	goto L65
L65:
	;
	v212 = F_make_ands_explicit(m, v203)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L27
	} else {
		goto L68
	}
L66:
	;
	v229 = v25 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v229 < v230 {
		v22 = v227
		v25 = v229
		goto L8
	} else {
		goto L74
	}
L67:
	;
	v225 = F_lappend(m, v22, v212)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L27
	} else {
		goto L73
	}
L68:
	;
	if v212 == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v216 != int32(21) {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v219 != int32(1) {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v223 = F_list_concat(m, v22, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L27
	} else {
		goto L72
	}
L72:
	;
	v227 = v223
	goto L66
L73:
	;
	v227 = v225
	goto L66
L74:
	;
	goto L9
L75:
	;
	v243 = F_make_orclause(m, v236)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L27
	} else {
		goto L76
	}
L76:
	;
	v247 = v243
	goto L4
}
func F_extract_variadic_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
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
	var v133 int32
	_ = v133
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
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
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
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == v5 {
		v31 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v34
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v34
	if v31&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	goto L1
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v23 == int32(0) {
		v31 = v5
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v26 != int32(15) {
		v31 = v5
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+13)))
	v31 = v29
	goto L2
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L14
	} else {
		goto L50
	}
L7:
	;
	m.G0 = v17 + int32(32)
	return v252
L8:
	;
	v252 = int32(-1)
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v237
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v235
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v236
	v252 = v234
	goto L7
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v40 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v113
	v115 = F_palloc0(m, v113)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L25
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	F_get_typlenbyvalalign(m, v46, v17+int32(16), v17+int32(19), v17+int32(18))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+16)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+19)))
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+18)))
	F_deconstruct_array(m, v42, v55, v56, v57, v17+int32(28), v17+int32(24), v17+int32(20))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v69 = F_palloc0(m, v66<<(uint(int32(2))%32))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if int32(0) < v71 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v80 = int32(0)
	goto L22
L20:
	;
	v103 = v71
	goto L21
L21:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v234 = v103
	v235 = v111
	v236 = v69
	v237 = v112
	goto L9
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69+v80<<(uint(int32(2))%32)))) = v46
	v94 = v80 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v94 < v95 {
		v80 = v94
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v103 = v95
	goto L21
L24:
	;
	goto L23
L25:
	;
	v118 = v113 << (uint(int32(2)) % 32)
	v119 = F_palloc0(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v121 = F_palloc0(m, v118)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	if v113 <= int32(0) {
		v234 = v113
		v235 = v115
		v236 = v121
		v237 = v119
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v133 = int32(0)
	goto L29
L29:
	;
	v145 = l0 + int32(20) + v133<<(uint(int32(3))%32)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133+v115))) = uint8(v146)
	v149 = v133 << (uint(int32(2)) % 32)
	v150 = v121 + v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v152 = F_get_fn_expr_argtype(m, v151, v133)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L14
	} else {
		goto L31
	}
L30:
	;
	v234 = v226
	v235 = v115
	v236 = v121
	v237 = v119
	goto L9
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v152
	if v152 != int32(705) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119+v149))) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if base.B2i32(v218 == int32(0))|base.B2i32(v218 == int32(705)) != 0 {
		goto L6
	} else {
		goto L48
	}
L33:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v216 = v215
	goto L32
L34:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v159 = int32(0)
	if v158 == v159 {
		v205 = v159
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v205 == int32(0) {
		goto L33
	} else {
		goto L45
	}
L36:
	;
	goto L35
L37:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	if v163 == int32(0) {
		v205 = v159
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v168 = v166 - int32(11)
	v175 = int32(0)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v168))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v168)%32))&int32(1) == v175)|base.B2i32(v133 < v175) != 0 {
		v205 = v159
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v168<<(uint(int32(2))%32))+uint32(_c_F_extract_variadic_args[0])))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v163+v183)))
	if v185 == int32(0) {
		v205 = v159
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v188 <= v133 {
		v205 = v159
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v190 = int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v133<<(uint(int32(2))%32))))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	switch v196 - int32(7) {
	case 0:
		v205 = v190
		goto L36
	case 1:
		goto L43
	default:
		goto L42
	}
L42:
	;
	v205 = int32(0)
	goto L36
L43:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if v199 == int32(0) {
		v205 = v190
		goto L36
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = int32(25)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+4)))
	if v211 != 0 {
		v216 = int32(0)
		goto L32
	} else {
		goto L46
	}
L46:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v213 = F_cstring_to_text(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	v216 = v213
	goto L32
L48:
	;
	v225 = v133 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v225 < v226 {
		v133 = v225
		goto L29
	} else {
		goto L49
	}
L49:
	;
	goto L30
L50:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v133 + int32(1)
	F_errmsg(m, int32(_a_F_extract_variadic_args_0), v17)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_extract_variadic_args_1), int32(2091), int32(_a_F_extract_variadic_args_2))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L14
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
