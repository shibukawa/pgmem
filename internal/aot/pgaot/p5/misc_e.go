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
	var v36 int32
	_ = v36
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
	v36 = v3
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v41 = v38 + v36*int32(60)
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
	v69 = v36 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v69 < v70 {
		v36 = v69
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
	v65 = F_ExecCallTriggerFunc(m, v9+int32(4), v36, v58, v59, v64)
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
	F_errmsg(m, int32(340934), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(487030), int32(3323), int32(132557))
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
	var v3 int32
	_ = v3
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
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
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
	v76 = v3
	goto L6
L6:
	;
	return base.I32_extend16_s(v76)
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
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+8)))
	v76 = v69
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
	v65 = v20 + int32(1)
	if v65 != v17 {
		v20 = v65
		goto L11
	} else {
		goto L25
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
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v39 == int32(0) {
		v58 = v38
		v59 = v39
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v59-v58 == int32(0) {
		goto L10
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	if v38 != v39 {
		v58 = v38
		v59 = v39
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v43 = v33
	v44 = l1
	goto L20
L20:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v48 == int32(0) {
		v58 = v47
		v59 = v48
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v58 = v47
	v59 = v48
	goto L17
L22:
	;
	v51 = int32(1)
	if v47 == v48 {
		v43 = v43 + v51
		v44 = v44 + v51
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L13
L25:
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
	var v28 int32
	_ = v28
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
	var v65 int32
	_ = v65
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v4
	v15 = int32(4470560)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 < v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = v4
	goto L4
L2:
	;
	v65 = v4
	goto L3
L3:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v67 = F_bms_add_members(m, v65, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L13
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(24)+v28<<(uint(int32(2))%32))))
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
	v65 = v56
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
	v53 = v28 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v53 < v54 {
		v28 = v53
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
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v16
	v71 = F_bms_copy(m, v67)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
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
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v74 = F_bms_copy(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_MemoryContextReset(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v74
	goto L17
L19:
	;
	m.G0 = v11 + int32(16)
	return v71
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
		v8 = int32(4470560)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, _consts[9])) = v12
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
				*(*int32)(unsafe.Add(mBase, _consts[9])) = v9
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
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
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v26 = l0 + int32(108)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L6
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v17 + int32(80)
	return v379
L8:
	;
	F_ExecReScan(m, v27)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v47 = m.T0[v46].(func(*base.Module, int32) int32)(m, v27)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v57 == int32(0) {
		v379 = v47
		goto L12
	} else {
		goto L20
	}
L14:
	;
	if v47 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v49&int32(2) == int32(0) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_EvalPlanQualEnd(m, v26)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v379 = int32(0)
	goto L12
L20:
	;
	v60 = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v62 <= v60 {
		v379 = v47
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v72 = v60
	v77 = v60
	goto L22
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v77<<(uint(int32(2))%32))))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v87 = F_EvalPlanQualSlot(m, v26, v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	if v343&int32(1) == int32(0) {
		v379 = v47
		goto L12
	} else {
		goto L96
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	m.T0[v90].(func(*base.Module, int32))(m, v87)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	if v93 == v94 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v346 = v77 + int32(1)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v346 < v347 {
		v72 = v343
		v77 = v346
		goto L22
	} else {
		goto L95
	}
L27:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L92
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L89
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L85
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L82
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L79
	}
L32:
	;
	v122 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+32)) = uint8(v122)
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+4)))
	v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+6)))
	if v125 < v124 {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+6)))
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+6)))
	if v97 < v96 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_slot_getsomeattrs_int(m, v47, v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v101 = int32(1)
	v102 = v96 - v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v103))))
	if v105 == v101 {
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v102<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v112 == v113 {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	v115 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+38)) = uint16(v115)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+34)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+32)) = uint8(v115)
	v343 = v72
	goto L26
L40:
	;
	F_slot_getsomeattrs_int(m, v47, v124)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v129 = int32(1)
	v130 = v124 - v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v131))))
	if v133 == v129 {
		goto L30
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v130<<(uint(int32(2))%32))))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+48))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+119)))
	if v143 == int32(102) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v146 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+51)) = uint8(v146)
	v149 = F_GetFdwRoutineForRelation(m, v141, v146)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v164 = v17 + int32(76)
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v164))) = uint16(v165)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if base.Ui32(int32(4)) <= base.Ui32(v169) {
		goto L28
	} else {
		goto L52
	}
L48:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+108))
	if v151 == int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	m.T0[v151].(func(*base.Module, int32, int32, int32, int32, int32))(m, v28, v84, v140, v87, v17+int32(51))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v158&int32(2) != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+51)))
	v343 = v161 | v72
	goto L26
L52:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	v177 = int32(3)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
	v180 = int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, _consts[322]))
	if v180 < v183 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L76
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L73
	}
L55:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[322]))
	if v222 < int32(2) {
		goto L6
	} else {
		goto L68
	}
L56:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[322]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L64
	}
L57:
	;
	v186 = v180
	goto L59
L58:
	;
	v186 = v177
	goto L59
L59:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v172)+188))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+104))
	v191 = m.T0[v190].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v172, v17+int32(72), v175, v87, v176, v177-v169, v179, v186, v17+int32(52))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	if v191 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	switch v191 - int32(1) {
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
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+68)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+34)) = v196
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164))))
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+38)) = uint16(v198)
	v343 = v195 | v72
	goto L26
L64:
	;
	if int32(2) <= v202 {
		goto L27
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(3)
	F_errmsg_internal(m, int32(56969), v17+int32(32))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(485136), int32(230), int32(112124))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
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
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(350022), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(485136), int32(237), int32(112124))
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
	F_errmsg_internal(m, int32(377852), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(485136), int32(242), int32(112124))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v191
	F_errmsg_internal(m, int32(56928), v17+int32(16))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(485136), int32(247), int32(112124))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
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
	F_errmsg_internal(m, int32(523401), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(485136), int32(102), int32(112124))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
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
	F_errmsg_internal(m, int32(523388), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(485136), int32(122), int32(112124))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
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
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v303 + int32(4)
	F_errmsg(m, int32(693781), v17)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(485136), int32(136), int32(112124))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
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
	F_errmsg_internal(m, int32(363010), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(485136), int32(176), int32(112124))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
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
	F_errmsg(m, int32(350022), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(485136), int32(228), int32(112124))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
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
	F_EvalPlanQualBegin(m, v26)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v47
	v356 = int32(4470560)
	v357 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+100))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+52))
	if v363 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_ExecReScan(m, v362)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	v367 = m.T0[v366].(func(*base.Module, int32) int32)(m, v362)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v357
	if v367 == int32(0) {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367)+4)))
	if v373&int32(2) != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v379 = v367
	goto L12
}
func F_ExecModifyTable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
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
	var v203 int32
	_ = v203
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
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
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
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
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
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
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
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v846 int32
	_ = v846
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
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
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v956 float64
	_ = v956
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
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
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 float64
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1072 int32
	_ = v1072
	var v1073 float64
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1101 int64
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1220 int32
	_ = v1220
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
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
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1385 float64
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1468 int32
	_ = v1468
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1504 int32
	_ = v1504
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1536 int32
	_ = v1536
	var v1541 int32
	_ = v1541
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1644 int32
	_ = v1644
	var v1649 int32
	_ = v1649
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1731 int32
	_ = v1731
	var v1749 int32
	_ = v1749
	var v1760 int32
	_ = v1760
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1782 int32
	_ = v1782
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1986 int32
	_ = v1986
	var v2020 int32
	_ = v2020
	var v2024 int32
	_ = v2024
	var v2029 int32
	_ = v2029
	var v2033 int32
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2042 int32
	_ = v2042
	v29 = m.G0
	v31 = v29 - int32(176)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+156))
	if v41 == int32(0) {
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L4
	} else {
		goto L588
	}
L7:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)))
	if v44 != 0 {
		v1986 = int32(0)
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L4
	} else {
		goto L585
	}
L10:
	;
	m.G0 = v31 + int32(176)
	return v1986
L11:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)))
	if v45 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	switch v49 - int32(2) {
	case 0:
		goto L16
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	default:
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v34
	v112 = int32(124)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = l0 + v112
	*(*int32)(unsafe.Add(mBase, uint32(v31)+108)) = l0
	v120 = v31 + v112
	v130 = v109 + v110*int32(216)
	goto L39
L15:
	;
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)) = uint8(v104)
	goto L14
L16:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSUpdateTriggers(m, v100, v48)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L38
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L35
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	if v65&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSDeleteTriggers(m, v62, v48)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L24
	}
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSInsertTriggers(m, v53, v48)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+132))
	if v56 != int32(2) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSUpdateTriggers(m, v59, v48)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	goto L15
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSInsertTriggers(m, v68, v48)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	v72 = v65
	goto L27
L27:
	;
	if v72&int32(2) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v72 = v71
	goto L27
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSUpdateTriggers(m, v75, v48)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v79 = v72
	goto L31
L31:
	;
	if v79&int32(4) == int32(0) {
		goto L15
	} else {
		goto L33
	}
L32:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v79 = v78
	goto L31
L33:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSDeleteTriggers(m, v84, v48)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L15
L35:
	;
	F_errmsg_internal(m, int32(255606), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(490519), int32(4008), int32(132602))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	goto L15
L39:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v155 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v1986 = v1956
	goto L10
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	F_MemoryContextReset(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v159 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
	F_MemoryContextReset(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v163 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+144)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v163
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v171 = F_ExecMergeNotMatched(m, v31+int32(108), v169, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	if v177 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v173 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v173
	if v171 == v173 {
		goto L39
	} else {
		goto L53
	}
L53:
	;
	v1986 = v171
	goto L10
L54:
	;
	F_ExecReScan(m, v108)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v181 = m.T0[v180].(func(*base.Module, int32) int32)(m, v108)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v183 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+144)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v181
	if v181 == v183 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	if v1956 == int32(0) {
		v130 = v243
		goto L39
	} else {
		goto L584
	}
L60:
	;
	v1948 = int32(0)
	v1949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v1953 = F_ExecDelete(m, v31+int32(108), v243, v393, v395, int32(1), v1948, v1949, v1948, v1948, v1948)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L4
	} else {
		goto L583
	}
L61:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v34)+188))
	if v1876 != 0 {
		goto L553
	} else {
		goto L554
	}
L62:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+4)))
	if v188&int32(2) != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v191 == int32(0) {
		v243 = v130
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+92)))
	if v246 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L65:
	;
	v194 = base.I32_extend16_s(v191)
	v195 = int32(*(*int16)(unsafe.Add(mBase, uint32(v181)+6)))
	if v195 < v194 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_slot_getsomeattrs_int(m, v181, v194)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v199 = int32(1)
	v200 = v194 - v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v201))))
	if v203 == v199 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	if v33 == int32(5) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v200<<(uint(int32(2))%32))))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v235 == v236 {
		v243 = v130
		goto L64
	} else {
		goto L81
	}
L73:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v208
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v214 = F_ExecMergeNotMatched(m, v31+int32(108), v212, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L78
	}
L76:
	;
	if v214 == int32(0) {
		goto L39
	} else {
		goto L77
	}
L77:
	;
	v1986 = v214
	goto L10
L78:
	;
	F_errmsg_internal(m, int32(523401), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(490519), int32(4303), int32(390459))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v240 = F_ExecLookupResultRelByOid(m, l0, v235, int32(0), int32(1))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v243 = v240
	goto L64
L83:
	;
	v251 = int32(0)
	v253 = F_ExecProcessReturning(m, v31+int32(108), v243, v33, v251, v251, v245)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v245
	v256 = int32(0)
	if base.Ui32(int32(5)) < base.Ui32(v33) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v1986 = v253
	goto L10
L87:
	;
	switch v33 - int32(2) {
	case 0:
		goto L135
	case 1:
		goto L136
	case 2:
		goto L60
	case 3:
		goto L134
	default:
		goto L6
	}
L88:
	;
	v393 = int32(0)
	v395 = v256
	goto L87
L89:
	;
	goto L90
L90:
	;
	if int32(1)<<(uint(v33)%32)&int32(52) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v393 = int32(0)
	v395 = v256
	goto L87
L92:
	;
	goto L93
L93:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+48))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+119)))
	v269 = v267 - int32(109)
	if base.Ui32(int32(5)) < base.Ui32(v269) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v326 = int32(0)
	v327 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+24)))
	if v327 == v326 {
		v393 = v326
		v395 = v256
		goto L87
	} else {
		goto L112
	}
L95:
	;
	if int32(1)<<(uint(v269)%32)&int32(41) == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v278 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+24)))
	v279 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+6)))
	if v279 < v278 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_slot_getsomeattrs_int(m, v245, v278)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v283 = int32(1)
	v284 = v278 - v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v245)+20))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v285))))
	if v287 == v283 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	if v33 == int32(5) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315+v284<<(uint(int32(2))%32))))
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v319)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+104)) = uint16(v320)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v322
	v393 = v31 + int32(100)
	v395 = v256
	goto L87
L104:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v292
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v298 = F_ExecMergeNotMatched(m, v31+int32(108), v296, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L109
	}
L107:
	;
	if v298 == int32(0) {
		v130 = v243
		goto L39
	} else {
		goto L108
	}
L108:
	;
	v1986 = v298
	goto L10
L109:
	;
	F_errmsg_internal(m, int32(523388), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(490519), int32(4400), int32(390459))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	v330 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+6)))
	if v330 < v327 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_slot_getsomeattrs_int(m, v245, v327)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v334 = int32(1)
	v335 = v327 - v334
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v245)+20))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335+v336))))
	if v338 == v334 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L115
L117:
	;
	if v33 == int32(5) {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366+v335<<(uint(int32(2))%32))))
	v371 = F_pg_detoast_datum(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L128
	}
L120:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v343
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v349 = F_ExecMergeNotMatched(m, v31+int32(108), v347, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	if v349 == int32(0) {
		v130 = v243
		goto L39
	} else {
		goto L124
	}
L124:
	;
	v1986 = v349
	goto L10
L125:
	;
	F_errmsg_internal(m, int32(523371), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(490519), int32(4457), int32(390459))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v371
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	v375 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+88)) = uint16(v375)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = int32(base.Ui32(v374) >> (uint(int32(2)) % 32))
	if v267 != int32(118) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)+56))
	v387 = v385
	goto L131
L130:
	;
	v387 = int32(0)
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+92)) = v387
	v393 = v326
	v395 = v31 + int32(80)
	goto L87
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L4
	} else {
		goto L550
	}
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L4
	} else {
		goto L547
	}
L134:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v393|v395 != 0 {
		goto L199
	} else {
		goto L200
	}
L135:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+48)))
	if v589 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L136:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+48)))
	if v396 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v399 = int32(0)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+52))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+44))
	if v405 == v399 {
		v458 = v399
		v474 = v399
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v243)+36))
	if v537 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L140:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	F_ExecCheckPlanOutput(m, v482, v474)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L151
	}
L141:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v408 <= int32(0) {
		v458 = v399
		v474 = v399
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v415 = v399
	v427 = v399
	v431 = v399
	goto L143
L143:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v439+v427<<(uint(int32(2))%32))))
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+26)))
	if v444 != 0 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v458 = v448
	v474 = v449
	goto L140
L145:
	;
	v451 = v427 + int32(1)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v451 < v452 {
		v415 = v448
		v427 = v451
		v431 = v449
		goto L143
	} else {
		goto L150
	}
L146:
	;
	v448 = int32(1)
	v449 = v431
	goto L145
L147:
	;
	goto L148
L148:
	;
	v446 = F_lappend(m, v431, v443)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v448 = v415
	v449 = v446
	goto L145
L150:
	;
	goto L144
L151:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v488 = F_table_slot_create(m, v485, v402+int32(104))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+40)) = v488
	if v458 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+52))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v493 != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	v506 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+48)) = uint8(v506)
	goto L139
L156:
	;
	v498 = v488
	v499 = v493
	goto L158
L157:
	;
	F_ExecAssignExprContext(m, v402, l0)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L159
	}
L158:
	;
	v500 = F_ExecBuildProjectionInfo(m, v474, v499, v498, l0, v492)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L4
	} else {
		goto L160
	}
L159:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v243)+40))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v498 = v496
	v499 = v497
	goto L158
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+36)) = v500
	goto L155
L161:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v583 = int32(0)
	v585 = F_ExecInsert(m, v31+int32(108), v243, v576, v582, v583, v583)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L4
	} else {
		goto L169
	}
L162:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v243)+40))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v536)+8))
	if v541 == v542 {
		v576 = v536
		goto L161
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v537)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v548)+12)) = v536
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v537)+72))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v537)+16))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+8))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+12))
	m.T0[v553].(func(*base.Module, int32))(m, v551)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L167
	}
L165:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v541)+32))
	m.T0[v544].(func(*base.Module, int32, int32))(m, v540, v536)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v243)+40))
	v576 = v547
	goto L161
L167:
	;
	v556 = int32(4470560)
	v557 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v550)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v559
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v537)+24))
	v565 = m.T0[v564].(func(*base.Module, int32, int32, int32) int32)(m, v537+int32(4), v550, int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v557
	v569 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551)+4)))
	v571 = v569 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v551)+4)) = uint16(v571)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v551)+12))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	*(*uint16)(unsafe.Add(mBase, uint32(v551)+6)) = uint16(v574)
	v576 = v551
	goto L161
L169:
	;
	if v585 == int32(0) {
		v130 = v243
		goto L39
	} else {
		goto L170
	}
L170:
	;
	v1986 = v585
	goto L10
L171:
	;
	F_ExecInitUpdateProjection(m, l0, v243)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	if v395 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L173
L175:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v243)+36))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v624)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v625)+4)) = v594
	*(*int32)(unsafe.Add(mBase, uint32(v625)+12)) = v623
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v624)+72))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v624)+16))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+8))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	m.T0[v631].(func(*base.Module, int32))(m, v629)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L190
	}
L176:
	;
	v595 = int32(0)
	F_ExecForceStoreHeapTuple(m, v395, v594, v595)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+49)))
	if v600 == int32(1) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v622 = v595
	goto L175
L180:
	;
	F_LockTuple(m, v599, v393, int32(7))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L4
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v607 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	goto L182
L184:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v609&int32(1) == int32(0) {
		goto L132
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v599)+188))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)+60))
	v617 = m.T0[v616].(func(*base.Module, int32, int32, int32, int32) int32)(m, v599, v393, int32(4140944), v594)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L4
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	if v617 == int32(0) {
		goto L133
	} else {
		goto L189
	}
L189:
	;
	v622 = v600
	goto L175
L190:
	;
	v634 = int32(4470560)
	v635 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v628)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v637
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v624)+24))
	v643 = m.T0[v642].(func(*base.Module, int32, int32, int32) int32)(m, v624+int32(4), v628, int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v635
	v647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v629)+4)))
	v649 = v647 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v629)+4)) = uint16(v649)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v629)+12))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)))
	*(*uint16)(unsafe.Add(mBase, uint32(v629)+6)) = uint16(v652)
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v657 = F_ExecUpdate(m, v31+int32(108), v243, v393, v395, v594, v629, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	if v622 == int32(0) {
		v1956 = v657
		goto L59
	} else {
		goto L193
	}
L193:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	F_UnlockTuple(m, v661, v393, int32(7))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	if v657 == int32(0) {
		v130 = v243
		goto L39
	} else {
		goto L195
	}
L195:
	;
	v1986 = v657
	goto L10
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L4
	} else {
		goto L544
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L4
	} else {
		goto L541
	}
L198:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+220)) = v1822
	v1986 = v1749
	goto L10
L199:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)+64))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v243)+164))
	if v672 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	v1817 = F_ExecMergeNotMatched(m, v31+int32(108), v243, v667&int32(1))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L4
	} else {
		goto L539
	}
L202:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v243)+168))
	if v675 == int32(0) {
		v130 = v243
		goto L39
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+4)) = v678
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	v681 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v670)+12)) = v681
	*(*int32)(unsafe.Add(mBase, uint32(v670)+8)) = v680
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+172)) = uint16(v681)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+168)) = int32(-1)
	if v395 != 0 {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	goto L204
L206:
	;
	v722 = v243 + int32(164)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v243)+176))
	if v723 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L207:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	F_ExecForceStoreHeapTuple(m, v395, v688, int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L4
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+49)))
	if v692 == int32(1) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	goto L206
L211:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	F_LockTuple(m, v695, v393, int32(7))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L4
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v704 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v704 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+172)) = uint16(v699)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+168)) = v701
	goto L213
L215:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v706&int32(1) == int32(0) {
		goto L196
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v711)+188))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+60))
	v716 = m.T0[v715].(func(*base.Module, int32, int32, int32, int32) int32)(m, v711, v393, int32(4140944), v713)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L219
	}
L218:
	;
	goto L217
L219:
	;
	if v716 == int32(0) {
		goto L197
	} else {
		goto L220
	}
L220:
	;
	goto L206
L221:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	if v747 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L222:
	;
	v743 = v722
	v746 = v243 + int32(168)
	goto L221
L223:
	;
	goto L224
L224:
	;
	v728 = int32(4470560)
	v729 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v670)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v731
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v723)+20))
	v736 = m.T0[v735].(func(*base.Module, int32, int32, int32) int32)(m, v723, v670, v31+int32(152))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L4
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v729
	v741 = v243 + int32(168)
	if v736 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v742 = v722
	goto L228
L227:
	;
	v742 = v741
	goto L228
L228:
	;
	v743 = v742
	v746 = v741
	goto L221
L229:
	;
	v1776 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+172)))
	if v1776 != 0 {
		goto L533
	} else {
		goto L534
	}
L230:
	;
	v1749 = int32(0)
	v1760 = int32(1)
	goto L229
L231:
	;
	goto L232
L232:
	;
	v753 = v669 + int32(124)
	v766 = v747
	v769 = int32(0)
	v770 = int32(1)
	goto L233
L233:
	;
	v786 = int32(0)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v766)+4))
	if v787 <= v786 {
		v1731 = v770
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v1749 = int32(0)
	v1760 = v1731
	goto L229
L235:
	;
	goto L234
L236:
	;
	v791 = v786
	goto L253
L237:
	;
	if v766 != 0 {
		v769 = v1092
		v770 = v1718
		goto L233
	} else {
		goto L532
	}
L238:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L4
	} else {
		goto L528
	}
L239:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L4
	} else {
		goto L523
	}
L240:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L4
	} else {
		goto L520
	}
L241:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L4
	} else {
		goto L517
	}
L242:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L4
	} else {
		goto L513
	}
L243:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L4
	} else {
		goto L509
	}
L244:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L4
	} else {
		goto L504
	}
L245:
	;
	v1749 = int32(0)
	v1760 = v1372
	goto L229
L246:
	;
	v1562 = int32(0)
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v243)+152))
	if v1563 == v1562 {
		v1749 = v1562
		v1760 = v770
		goto L229
	} else {
		goto L495
	}
L247:
	;
	if v1089 != int32(3) {
		goto L332
	} else {
		goto L333
	}
L248:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v31)+164))
	if v1086 == int32(0) {
		v1731 = v770
		goto L235
	} else {
		goto L330
	}
L249:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v31)+164))
	v1089 = v1081
	v1092 = v1080
	goto L247
L250:
	;
	if v1065 != 0 {
		v1080 = v865
		goto L249
	} else {
		goto L328
	}
L251:
	;
	if v1033 != 0 {
		v1080 = v769
		goto L249
	} else {
		goto L315
	}
L252:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L4
	} else {
		goto L312
	}
L253:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v766)+12))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v818+v791<<(uint(int32(2))%32))))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v822)+4))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v823)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(160)))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+152)) = int64(0)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v822)+12))
	if v829 != 0 {
		goto L257
	} else {
		goto L258
	}
L254:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+64))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+8))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+12))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+188))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+96))
	v1017 = m.T0[v1016].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1008, v393, v1010, v1011, v1012, int32(1), v120, int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L4
	} else {
		goto L311
	}
L255:
	;
	goto L254
L256:
	;
	v1004 = v791 + int32(1)
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v766)+4))
	if v1004 < v1005 {
		v791 = v1004
		goto L253
	} else {
		goto L310
	}
L257:
	;
	v830 = int32(4470560)
	v831 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v670)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v833
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v829)+20))
	v838 = m.T0[v837].(func(*base.Module, int32, int32, int32) int32)(m, v829, v670, v31+int32(175))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L4
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v243)+116))
	if v846 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v831
	if v838 == int32(0) {
		goto L256
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	v862 = v824 - int32(2)
	switch v862 {
	case 0:
		goto L271
	default:
		goto L252
	case 2:
		goto L270
	case 5:
		goto L269
	}
L263:
	;
	if v824 == int32(7) {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	if v824 == int32(2) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v855 = int32(4)
	goto L267
L266:
	;
	v855 = int32(5)
	goto L267
L267:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v857)+8))
	F_ExecWithCheckOptions(m, v855, v243, v856, v858)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	goto L262
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = int32(0)
	v1561 = v769
	goto L246
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+216)) = v822
	v962 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = v962
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v243)+52))
	if v964 == v962 {
		goto L255
	} else {
		goto L296
	}
L271:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v822)+8))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v863)+72))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v863)+16))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v865)+8))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v866)+12))
	m.T0[v867].(func(*base.Module, int32))(m, v865)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	v870 = int32(4470560)
	v871 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v864)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v873
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v863)+24))
	v879 = m.T0[v878].(func(*base.Module, int32, int32, int32) int32)(m, v863+int32(4), v864, int32(0))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v871
	v883 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v865)+4)))
	v885 = v883 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v865)+4)) = uint16(v885)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v865)+12))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v887)))
	*(*uint16)(unsafe.Add(mBase, uint32(v865)+6)) = uint16(v888)
	*(*int32)(unsafe.Add(mBase, uint32(v669)+216)) = v822
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = int32(0)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v865)+8))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v894)+28))
	m.T0[v895].(func(*base.Module, int32))(m, v865)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v891)+48))
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+116)))
	if v899 != int32(1) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v243)+52))
	if v906 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v243)+16))
	if v902 != 0 {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	F_ExecOpenIndices(m, v243, int32(0))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L4
	} else {
		goto L278
	}
L278:
	;
	goto L275
L279:
	;
	v950 = F_ExecUpdateAct(m, v31+int32(108), v243, v393, int32(0), v865, v667&int32(1), v31+int32(152))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L4
	} else {
		goto L294
	}
L280:
	;
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906)+13)))
	if v909 == int32(1) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)+188))
	if v913 != 0 {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	v933 = v906
	goto L283
L283:
	;
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933)+15)))
	if v934 != int32(1) {
		goto L279
	} else {
		goto L291
	}
L284:
	;
	F_ExecPendingInserts(m, v912)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L4
	} else {
		goto L287
	}
L285:
	;
	v917 = v912
	goto L286
L286:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v31)+112))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v922)+104))
	v926 = F_ExecBRUpdateTriggers(m, v917, v918, v243, v393, int32(0), v865, v31+int32(164), v120, base.B2i32(v923 == int32(5)))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L4
	} else {
		goto L288
	}
L287:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	v917 = v916
	goto L286
L288:
	;
	if v926 == int32(0) {
		v1085 = v865
		goto L248
	} else {
		goto L289
	}
L289:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v243)+52))
	if v930 == int32(0) {
		goto L279
	} else {
		goto L290
	}
L290:
	;
	v933 = v930
	goto L283
L291:
	;
	v937 = F_ExecIRUpdateTriggers(m, v671, v243, v395, v865)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L4
	} else {
		goto L292
	}
L292:
	;
	if v937 == int32(0) {
		v1731 = v770
		goto L235
	} else {
		goto L293
	}
L293:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v31)+164))
	v1065 = v941
	goto L250
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = v950
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+152)))
	if v953 != int32(1) {
		v1065 = v950
		goto L250
	} else {
		goto L295
	}
L295:
	;
	v956 = *(*float64)(unsafe.Add(mBase, uint32(v669)+232))
	*(*float64)(unsafe.Add(mBase, uint32(v669)+232)) = base.F64_add(v956, float64(1))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v31)+148))
	v1749 = v960
	v1760 = v770
	goto L229
L296:
	;
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964)+18)))
	if v967 == int32(1) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v970)+188))
	if v971 != 0 {
		goto L300
	} else {
		goto L301
	}
L298:
	;
	v992 = v964
	goto L299
L299:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+20)))
	if v993 != int32(1) {
		goto L255
	} else {
		goto L307
	}
L300:
	;
	F_ExecPendingInserts(m, v970)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L4
	} else {
		goto L303
	}
L301:
	;
	v975 = v970
	goto L302
L302:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v31)+112))
	v977 = int32(0)
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v981)+104))
	v985 = F_ExecBRDeleteTriggers(m, v975, v976, v243, v393, v977, v977, v31+int32(164), v120, base.B2i32(v982 == int32(5)))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L4
	} else {
		goto L304
	}
L303:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	v975 = v974
	goto L302
L304:
	;
	if v985 == int32(0) {
		v1085 = v769
		goto L248
	} else {
		goto L305
	}
L305:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v243)+52))
	if v989 == int32(0) {
		goto L255
	} else {
		goto L306
	}
L306:
	;
	v992 = v989
	goto L299
L307:
	;
	v996 = F_ExecIRDeleteTriggers(m, v671, v243, v395)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L4
	} else {
		goto L308
	}
L308:
	;
	if v996 == int32(0) {
		v1731 = v770
		goto L235
	} else {
		goto L309
	}
L309:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v31)+164))
	v1033 = v1000
	goto L251
L310:
	;
	v1731 = v770
	goto L235
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = v1017
	v1033 = v1017
	goto L251
L312:
	;
	F_errmsg_internal(m, int32(352865), int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L4
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(490519), int32(3286), int32(451697))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L4
	} else {
		goto L314
	}
L314:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L315:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+204))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+104))
	if v1038 != int32(2) {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	v1057 = int32(0)
	F_ExecARDeleteTriggers(m, v1037, v243, v393, v1057, v1056, v1057)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L4
	} else {
		goto L327
	}
L317:
	;
	v1056 = v1036
	goto L316
L318:
	;
	goto L319
L319:
	;
	if v1036 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1056 = int32(0)
	goto L316
L321:
	;
	goto L322
L322:
	;
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1036)+1)))
	if v1044 != int32(1) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1056 = v1036
	goto L316
L324:
	;
	goto L325
L325:
	;
	v1047 = int32(0)
	F_ExecARUpdateTriggers(m, v1037, v243, v1047, v1047, v393, v1047, v1047, v1047, v1036, v1047)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L4
	} else {
		goto L326
	}
L326:
	;
	v1056 = v1047
	goto L316
L327:
	;
	v1061 = *(*float64)(unsafe.Add(mBase, uint32(v669)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v669)+240)) = base.F64_add(v1061, float64(1))
	v1080 = v769
	goto L249
L328:
	;
	F_ExecUpdateEpilogue(m, v31+int32(108), v31+int32(152), v243, v393, int32(0), v865)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L4
	} else {
		goto L329
	}
L329:
	;
	v1073 = *(*float64)(unsafe.Add(mBase, uint32(v669)+232))
	*(*float64)(unsafe.Add(mBase, uint32(v669)+232)) = base.F64_add(v1073, float64(1))
	v1080 = v865
	goto L249
L330:
	;
	v1089 = v1086
	v1092 = v1085
	goto L247
L331:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L4
	} else {
		goto L492
	}
L332:
	;
	switch v1089 {
	case 0:
		goto L337
	case 1, 5, 6:
		goto L331
	case 2:
		goto L336
	default:
		v1561 = v1092
		goto L246
	case 4:
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v822)+4))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+4))
	v1267 = F_ExecUpdateLockMode(m, v671, v243)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L4
	} else {
		goto L390
	}
L335:
	;
	v1242 = int32(0)
	v1245 = *(*int32)(unsafe.Add(mBase, _consts[322]))
	if v1245 < int32(2) {
		v1749 = v1242
		v1760 = v1242
		goto L229
	} else {
		goto L385
	}
L336:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v671)+64))
	if v1105 != v1106 {
		goto L244
	} else {
		goto L340
	}
L337:
	;
	if v667&int32(1) == int32(0) {
		v1561 = v1092
		goto L246
	} else {
		goto L338
	}
L338:
	;
	if v824 == int32(7) {
		v1561 = v1092
		goto L246
	} else {
		goto L339
	}
L339:
	;
	v1101 = *(*int64)(unsafe.Add(mBase, uint32(v671)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v671)+112)) = v1101 + int64(1)
	v1561 = v1092
	goto L246
L340:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v31)+132))
	if base.Ui32(v1108) < base.Ui32(int32(3)) {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L4
	} else {
		goto L381
	}
L342:
	;
	v1228 = int32(0)
	goto L341
L343:
	;
	goto L344
L344:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	if v1119 == v1108 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1228 = int32(1)
	goto L341
L346:
	;
	goto L347
L347:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v1123 <= int32(0) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1228 = v1220
	goto L341
L349:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v1127 == int32(0) {
		v1220 = int32(0)
		goto L348
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v1191 = int32(0)
	v1193 = v1123 - int32(1)
	goto L371
L352:
	;
	v1132 = v1127
	goto L353
L353:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+20))
	if v1137 == int32(4) {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v1220 = int32(0)
	goto L348
L355:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+80))
	if v1184 != 0 {
		v1132 = v1184
		goto L353
	} else {
		goto L370
	}
L356:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1132)))
	if v1140 == int32(0) {
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v1143 = int32(1)
	if v1108 == v1140 {
		v1220 = v1143
		goto L348
	} else {
		goto L358
	}
L358:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+52))
	v1147 = v1145 - int32(1)
	if v1147 < int32(0) {
		goto L355
	} else {
		goto L359
	}
L359:
	;
	v1152 = int32(0)
	v1154 = v1147
	goto L360
L360:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+48))
	v1160 = int32(2)
	v1161 = base.I32_div_s(v1154-v1152, v1160)
	v1162 = v1161 + v1152
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1158+v1162<<(uint(v1160)%32))))
	if v1166 == v1108 {
		v1220 = v1143
		goto L348
	} else {
		goto L362
	}
L361:
	;
	goto L355
L362:
	;
	v1170 = F_TransactionIdPrecedes(m, v1166, v1108)
	mBase = m.M
	if v1170 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1171 = v1162 + int32(1)
	goto L365
L364:
	;
	v1171 = v1152
	goto L365
L365:
	;
	if v1170 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1174 = v1154
	goto L368
L367:
	;
	v1174 = v1162 - int32(1)
	goto L368
L368:
	;
	if v1171 <= v1174 {
		v1152 = v1171
		v1154 = v1174
		goto L360
	} else {
		goto L369
	}
L369:
	;
	goto L361
L370:
	;
	goto L354
L371:
	;
	v1198 = int32(2)
	v1199 = base.I32_div_s(v1193-v1191, v1198)
	v1200 = v1199 + v1191
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1189+v1200<<(uint(v1198)%32))))
	v1205 = base.B2i32(v1204 == v1108)
	if v1204 == v1108 {
		v1220 = v1205
		goto L348
	} else {
		goto L373
	}
L372:
	;
	v1220 = v1205
	goto L348
L373:
	;
	v1208 = base.B2i32(base.Ui32(v1204) < base.Ui32(v1108))
	if base.Ui32(v1204) < base.Ui32(v1108) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1209 = v1200 + int32(1)
	goto L376
L375:
	;
	v1209 = v1191
	goto L376
L376:
	;
	if base.Ui32(v1204) < base.Ui32(v1108) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1212 = v1193
	goto L379
L378:
	;
	v1212 = v1200 - int32(1)
	goto L379
L379:
	;
	if v1209 <= v1212 {
		v1191 = v1209
		v1193 = v1212
		goto L371
	} else {
		goto L380
	}
L380:
	;
	goto L372
L381:
	;
	if v1228 != 0 {
		goto L243
	} else {
		goto L382
	}
L382:
	;
	F_errmsg_internal(m, int32(377925), int32(0))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L4
	} else {
		goto L383
	}
L383:
	;
	F_errfinish(m, int32(490519), int32(3333), int32(451697))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	F_errmsg(m, int32(345366), int32(0))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L4
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(490519), int32(3340), int32(451697))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L4
	} else {
		goto L389
	}
L389:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L390:
	;
	if v1266 == int32(0) {
		goto L392
	} else {
		goto L393
	}
L391:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v671)+8))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v671)+64))
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+188))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+104))
	v1282 = m.T0[v1281].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1264, v393, v1276, v1275, v1277, v1267, int32(0), int32(2), v120)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L4
	} else {
		goto L396
	}
L392:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v1272 = F_EvalPlanQualSlot(m, v753, v1264, v1271)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L4
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	v1275 = v1274
	goto L391
L395:
	;
	v1275 = v1272
	goto L391
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = v1282
	if v1282 != 0 {
		goto L399
	} else {
		goto L400
	}
L397:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L4
	} else {
		goto L489
	}
L398:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v671)+64))
	if v1389 != v1390 {
		goto L239
	} else {
		goto L444
	}
L399:
	;
	v1285 = int32(0)
	switch v1282 - int32(2) {
	case 0:
		goto L398
	default:
		goto L397
	case 2:
		v1749 = v1285
		v1760 = v1285
		goto L229
	}
L400:
	;
	goto L401
L401:
	;
	v1289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+4)))
	if v1289 == int32(65533) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393))))
	v1293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+2)))
	if v1292&v1293 == int32(65535) {
		goto L242
	} else {
		goto L405
	}
L403:
	;
	goto L404
L404:
	;
	if v1266 != 0 {
		v1718 = v770
		goto L237
	} else {
		goto L406
	}
L405:
	;
	goto L404
L406:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v1298 = F_EvalPlanQual(m, v753, v1264, v1297, v1275)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L4
	} else {
		goto L407
	}
L407:
	;
	if v1298 == int32(0) {
		v1731 = v770
		goto L235
	} else {
		goto L408
	}
L408:
	;
	v1302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1298)+4)))
	if v1302&int32(2) != 0 {
		v1731 = v770
		goto L235
	} else {
		goto L409
	}
L409:
	;
	v1305 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+24)))
	v1306 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1298)+6)))
	if v1306 < v1305 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	F_slot_getsomeattrs_int(m, v1298, v1305)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L4
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1298)+20))
	v1312 = int32(1)
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1310+v1305-v1312))))
	v1315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+49)))
	if v1315 == v1312 {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	goto L412
L414:
	;
	v1318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+172)))
	if v1318 != 0 {
		goto L417
	} else {
		goto L418
	}
L415:
	;
	goto L416
L416:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v1334 != 0 {
		goto L422
	} else {
		goto L423
	}
L417:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	F_UnlockTuple(m, v1319, v31+int32(168), int32(7))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L4
	} else {
		goto L420
	}
L418:
	;
	goto L419
L419:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	F_LockTuple(m, v1325, v393, int32(7))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L4
	} else {
		goto L421
	}
L420:
	;
	goto L419
L421:
	;
	v1329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+172)) = uint16(v1329)
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+168)) = v1331
	goto L416
L422:
	;
	v1336 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v1336&int32(1) == int32(0) {
		goto L241
	} else {
		goto L425
	}
L423:
	;
	goto L424
L424:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+188))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+60))
	v1345 = m.T0[v1344].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1264, v393, int32(4140944), v1342)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L4
	} else {
		goto L426
	}
L425:
	;
	goto L424
L426:
	;
	if v1345 == int32(0) {
		goto L240
	} else {
		goto L427
	}
L427:
	;
	if v1314 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1350 = int32(0)
	goto L430
L429:
	;
	v1350 = v770
	goto L430
L430:
	;
	if v1350&int32(1) != 0 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v243)+176))
	if v1353 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L432:
	;
	v1372 = v1350
	goto L433
L433:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v669)+36))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	if v1377 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L434:
	;
	v1718 = int32(1)
	goto L237
L435:
	;
	goto L436
L436:
	;
	v1358 = int32(4470560)
	v1359 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v670)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1361
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1353)+20))
	v1366 = m.T0[v1365].(func(*base.Module, int32, int32, int32) int32)(m, v1353, v670, v31+int32(175))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L4
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1359
	if v1366 != 0 {
		v1718 = int32(1)
		goto L237
	} else {
		goto L438
	}
L438:
	;
	v1372 = int32(0)
	goto L433
L439:
	;
	if v1375 != 0 {
		v766 = v1375
		v769 = v1092
		v770 = v1372
		goto L233
	} else {
		goto L443
	}
L440:
	;
	if v1375 == int32(0) {
		goto L245
	} else {
		goto L441
	}
L441:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v243+int32(172))))
	if v1382 == int32(0) {
		goto L439
	} else {
		goto L442
	}
L442:
	;
	v1385 = *(*float64)(unsafe.Add(mBase, uint32(v1377)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1377)+32)) = base.F64_add(v1385, float64(1))
	goto L439
L443:
	;
	v1731 = v1372
	goto L235
L444:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v31)+132))
	if base.Ui32(v1392) < base.Ui32(int32(3)) {
		goto L446
	} else {
		goto L447
	}
L445:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L4
	} else {
		goto L485
	}
L446:
	;
	v1512 = int32(0)
	goto L445
L447:
	;
	goto L448
L448:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	if v1403 == v1392 {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v1512 = int32(1)
	goto L445
L450:
	;
	goto L451
L451:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v1407 <= int32(0) {
		goto L453
	} else {
		goto L454
	}
L452:
	;
	v1512 = v1504
	goto L445
L453:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v1411 == int32(0) {
		v1504 = int32(0)
		goto L452
	} else {
		goto L456
	}
L454:
	;
	goto L455
L455:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v1475 = int32(0)
	v1477 = v1407 - int32(1)
	goto L475
L456:
	;
	v1416 = v1411
	goto L457
L457:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1416)+20))
	if v1421 == int32(4) {
		goto L459
	} else {
		goto L460
	}
L458:
	;
	v1504 = int32(0)
	goto L452
L459:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1416)+80))
	if v1468 != 0 {
		v1416 = v1468
		goto L457
	} else {
		goto L474
	}
L460:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1416)))
	if v1424 == int32(0) {
		goto L459
	} else {
		goto L461
	}
L461:
	;
	v1427 = int32(1)
	if v1392 == v1424 {
		v1504 = v1427
		goto L452
	} else {
		goto L462
	}
L462:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1416)+52))
	v1431 = v1429 - int32(1)
	if v1431 < int32(0) {
		goto L459
	} else {
		goto L463
	}
L463:
	;
	v1436 = int32(0)
	v1438 = v1431
	goto L464
L464:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1416)+48))
	v1444 = int32(2)
	v1445 = base.I32_div_s(v1438-v1436, v1444)
	v1446 = v1445 + v1436
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1442+v1446<<(uint(v1444)%32))))
	if v1450 == v1392 {
		v1504 = v1427
		goto L452
	} else {
		goto L466
	}
L465:
	;
	goto L459
L466:
	;
	v1454 = F_TransactionIdPrecedes(m, v1450, v1392)
	mBase = m.M
	if v1454 != 0 {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v1455 = v1446 + int32(1)
	goto L469
L468:
	;
	v1455 = v1436
	goto L469
L469:
	;
	if v1454 != 0 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v1458 = v1438
	goto L472
L471:
	;
	v1458 = v1446 - int32(1)
	goto L472
L472:
	;
	if v1455 <= v1458 {
		v1436 = v1455
		v1438 = v1458
		goto L464
	} else {
		goto L473
	}
L473:
	;
	goto L465
L474:
	;
	goto L458
L475:
	;
	v1482 = int32(2)
	v1483 = base.I32_div_s(v1477-v1475, v1482)
	v1484 = v1483 + v1475
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1473+v1484<<(uint(v1482)%32))))
	v1489 = base.B2i32(v1488 == v1392)
	if v1488 == v1392 {
		v1504 = v1489
		goto L452
	} else {
		goto L477
	}
L476:
	;
	v1504 = v1489
	goto L452
L477:
	;
	v1492 = base.B2i32(base.Ui32(v1488) < base.Ui32(v1392))
	if base.Ui32(v1488) < base.Ui32(v1392) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v1493 = v1484 + int32(1)
	goto L480
L479:
	;
	v1493 = v1475
	goto L480
L480:
	;
	if base.Ui32(v1488) < base.Ui32(v1392) {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	v1496 = v1477
	goto L483
L482:
	;
	v1496 = v1484 - int32(1)
	goto L483
L483:
	;
	if v1493 <= v1496 {
		v1475 = v1493
		v1477 = v1496
		goto L475
	} else {
		goto L484
	}
L484:
	;
	goto L476
L485:
	;
	if v1512 != 0 {
		goto L238
	} else {
		goto L486
	}
L486:
	;
	F_errmsg_internal(m, int32(377925), int32(0))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L4
	} else {
		goto L487
	}
L487:
	;
	F_errfinish(m, int32(490519), int32(3525), int32(451697))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L4
	} else {
		goto L488
	}
L488:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L489:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v31)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v1530
	F_errmsg_internal(m, int32(56969), v31+int32(32))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L4
	} else {
		goto L490
	}
L490:
	;
	F_errfinish(m, int32(490519), int32(3531), int32(451697))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
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
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v31)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v1546
	F_errmsg_internal(m, int32(473291), v31-int32(-64))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L4
	} else {
		goto L493
	}
L493:
	;
	F_errfinish(m, int32(490519), int32(3540), int32(451697))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
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
	switch v862 {
	case 0:
		goto L498
	default:
		goto L496
	case 2:
		goto L497
	case 5:
		v1749 = v1562
		v1760 = v770
		goto L229
	}
L496:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L4
	} else {
		goto L501
	}
L497:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	v1579 = F_ExecProcessReturning(m, v31+int32(108), v243, int32(4), v1576, int32(0), v1578)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L4
	} else {
		goto L500
	}
L498:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	v1571 = F_ExecProcessReturning(m, v31+int32(108), v243, int32(2), v1569, v1561, v1570)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L4
	} else {
		goto L499
	}
L499:
	;
	v1749 = v1571
	v1760 = v770
	goto L229
L500:
	;
	v1749 = v1579
	v1760 = v770
	goto L229
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v824
	F_errmsg_internal(m, int32(478062), v31)
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L4
	} else {
		goto L502
	}
L502:
	;
	F_errfinish(m, int32(490519), int32(3572), int32(451697))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L4
	} else {
		goto L503
	}
L503:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L504:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L4
	} else {
		goto L505
	}
L505:
	;
	F_errmsg(m, int32(420688), int32(0))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L4
	} else {
		goto L506
	}
L506:
	;
	F_errhint(m, int32(558588), int32(0))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L4
	} else {
		goto L507
	}
L507:
	;
	F_errfinish(m, int32(490519), int32(3322), int32(451697))
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L4
	} else {
		goto L508
	}
L508:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = int32(532234)
	F_errmsg(m, int32(369505), v31+int32(16))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L4
	} else {
		goto L510
	}
L510:
	;
	F_errhint(m, int32(552313), int32(0))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L4
	} else {
		goto L511
	}
L511:
	;
	F_errfinish(m, int32(490519), int32(3330), int32(451697))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L4
	} else {
		goto L512
	}
L512:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L513:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L4
	} else {
		goto L514
	}
L514:
	;
	F_errmsg(m, int32(350157), int32(0))
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L4
	} else {
		goto L515
	}
L515:
	;
	F_errfinish(m, int32(490519), int32(3399), int32(451697))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L4
	} else {
		goto L516
	}
L516:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L517:
	;
	F_errmsg_internal(m, int32(330566), int32(0))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L4
	} else {
		goto L518
	}
L518:
	;
	F_errfinish(m, int32(321185), int32(1264), int32(266896))
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L4
	} else {
		goto L519
	}
L519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L520:
	;
	F_errmsg_internal(m, int32(377324), int32(0))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L4
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(490519), int32(3453), int32(451697))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L4
	} else {
		goto L522
	}
L522:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L523:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L4
	} else {
		goto L524
	}
L524:
	;
	F_errmsg(m, int32(420688), int32(0))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L4
	} else {
		goto L525
	}
L525:
	;
	F_errhint(m, int32(558588), int32(0))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L4
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(490519), int32(3514), int32(451697))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L4
	} else {
		goto L527
	}
L527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = int32(532234)
	F_errmsg(m, int32(369505), v31+int32(48))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L4
	} else {
		goto L529
	}
L529:
	;
	F_errhint(m, int32(552313), int32(0))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L4
	} else {
		goto L530
	}
L530:
	;
	F_errfinish(m, int32(490519), int32(3522), int32(451697))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L4
	} else {
		goto L531
	}
L531:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L532:
	;
	v1731 = v1718
	goto L235
L533:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	F_UnlockTuple(m, v1777, v31+int32(168), int32(7))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L4
	} else {
		goto L536
	}
L534:
	;
	goto L535
L535:
	;
	if v1760&int32(1) != 0 {
		v1956 = v1749
		goto L59
	} else {
		goto L537
	}
L536:
	;
	goto L535
L537:
	;
	if v1749 != 0 {
		goto L198
	} else {
		goto L538
	}
L538:
	;
	goto L201
L539:
	;
	if v1817 == int32(0) {
		v130 = v243
		goto L39
	} else {
		goto L540
	}
L540:
	;
	v1986 = v1817
	goto L10
L541:
	;
	F_errmsg_internal(m, int32(377324), int32(0))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L4
	} else {
		goto L542
	}
L542:
	;
	F_errfinish(m, int32(490519), int32(3124), int32(451697))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L4
	} else {
		goto L543
	}
L543:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L544:
	;
	F_errmsg_internal(m, int32(330566), int32(0))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L4
	} else {
		goto L545
	}
L545:
	;
	F_errfinish(m, int32(321185), int32(1264), int32(266896))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L4
	} else {
		goto L546
	}
L546:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L547:
	;
	F_errmsg_internal(m, int32(441335), int32(0))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L4
	} else {
		goto L548
	}
L548:
	;
	F_errfinish(m, int32(490519), int32(4520), int32(390459))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L4
	} else {
		goto L549
	}
L549:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L550:
	;
	F_errmsg_internal(m, int32(330566), int32(0))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L4
	} else {
		goto L551
	}
L551:
	;
	F_errfinish(m, int32(321185), int32(1264), int32(266896))
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
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
	F_ExecPendingInserts(m, v34)
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L4
	} else {
		goto L556
	}
L554:
	;
	goto L555
L555:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	switch v1880 - int32(2) {
	case 0:
		goto L558
	case 1:
		goto L562
	case 2:
		goto L561
	case 3:
		goto L560
	default:
		goto L559
	}
L556:
	;
	goto L555
L557:
	;
	v1942 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v1942)
	v1986 = int32(0)
	goto L10
L558:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASUpdateTriggers(m, v1937, v1879, v1938)
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L4
	} else {
		goto L582
	}
L559:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L4
	} else {
		goto L579
	}
L560:
	;
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	if v1899&int32(4) != 0 {
		goto L569
	} else {
		goto L570
	}
L561:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASDeleteTriggers(m, v1895, v1879, v1896)
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		goto L4
	} else {
		goto L568
	}
L562:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1883)+132))
	if v1884 == int32(2) {
		goto L563
	} else {
		goto L564
	}
L563:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	F_ExecASUpdateTriggers(m, v1887, v1879, v1888)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L4
	} else {
		goto L566
	}
L564:
	;
	goto L565
L565:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASInsertTriggers(m, v1891, v1879, v1892)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L4
	} else {
		goto L567
	}
L566:
	;
	goto L565
L567:
	;
	goto L557
L568:
	;
	goto L557
L569:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASDeleteTriggers(m, v1902, v1879, v1903)
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L4
	} else {
		goto L572
	}
L570:
	;
	v1907 = v1899
	goto L571
L571:
	;
	if v1907&int32(2) != 0 {
		goto L573
	} else {
		goto L574
	}
L572:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v1907 = v1906
	goto L571
L573:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASUpdateTriggers(m, v1910, v1879, v1911)
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L4
	} else {
		goto L576
	}
L574:
	;
	v1915 = v1907
	goto L575
L575:
	;
	if v1915&int32(1) == int32(0) {
		goto L557
	} else {
		goto L577
	}
L576:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v1915 = v1914
	goto L575
L577:
	;
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASInsertTriggers(m, v1920, v1879, v1921)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L4
	} else {
		goto L578
	}
L578:
	;
	goto L557
L579:
	;
	F_errmsg_internal(m, int32(255606), int32(0))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L4
	} else {
		goto L580
	}
L580:
	;
	F_errfinish(m, int32(490519), int32(4052), int32(132617))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L4
	} else {
		goto L581
	}
L581:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L582:
	;
	goto L557
L583:
	;
	v1956 = v1953
	goto L59
L584:
	;
	goto L40
L585:
	;
	F_errmsg_internal(m, int32(304731), int32(0))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L4
	} else {
		goto L586
	}
L586:
	;
	F_errfinish(m, int32(490519), int32(4179), int32(390459))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L4
	} else {
		goto L587
	}
L587:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L588:
	;
	F_errmsg_internal(m, int32(255606), int32(0))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L4
	} else {
		goto L589
	}
L589:
	;
	F_errfinish(m, int32(490519), int32(4544), int32(390459))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L4
	} else {
		goto L590
	}
L590:
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
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
	var v127 int32
	_ = v127
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
	var v227 int32
	_ = v227
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v329 int64
	_ = v329
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
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
	v28 = int32(4470560)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v31
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
	v50 = v3
	v53 = v3
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v29
	goto L3
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v58 = v57 + v50
	v60 = v50 << (uint(int32(2)) % 32)
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
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v29
	if v576&int32(1) == int32(0) {
		goto L3
	} else {
		goto L116
	}
L9:
	;
	v581 = v50 + int32(1)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v581 < v582 {
		v50 = v581
		v53 = v576
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
	v576 = v53
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
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v556 = m.T0[v555].(func(*base.Module, int32, int32, int32) int32)(m, v67, v20, v58)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L114
	}
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	if v87 != 0 {
		v466 = v87
		goto L18
	} else {
		goto L19
	}
L17:
	;
	m.G0 = v83 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v527
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v550 = base.B2i32(v547 != int32(2)) | v53
	if v547 != int32(1) {
		v576 = v550
		goto L9
	} else {
		goto L113
	}
L18:
	;
	v480 = int32(0)
	v481 = int32(4470560)
	v482 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+24))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v485
	v489 = F_tuplestore_gettupleslot(m, v466, int32(1), v480, v484)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
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
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L96
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L92
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
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
	v409 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v409)
	v527 = int32(0)
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
	v111 = int32(4470560)
	v112 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v80
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
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v112
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
	v127 = int32(0)
	goto L32
L32:
	;
	v145 = v107 + int32(20) + v127<<(uint(int32(3))%32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v127<<(uint(int32(2))%32))))
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
	v158 = v127 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v158 < v159 {
		v127 = v158
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	F_pgstat_init_function_usage(m, v107, v83+int32(48))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L45
	}
L37:
	;
	v218 = int32(*(*int16)(unsafe.Add(mBase, uint32(v107)+18)))
	if v218 <= int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v227 = int32(0)
	goto L39
L39:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+int32(24)+v227<<(uint(int32(3))%32)))))
	if v246 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v252 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(2)
	v527 = int32(0)
	goto L17
L41:
	;
	v250 = v227 + int32(1)
	if v218 != v250 {
		v227 = v250
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
	v280 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+16)) = uint8(v280)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+36)) = v280
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v286 = m.T0[v285].(func(*base.Module, int32) int32)(m, v107)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v288)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v83)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v290
	v293 = v83 + int32(48)
	v302 = m.G0
	v304 = v302 - int32(16)
	m.G0 = v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	if v306 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	if v341 != int32(2) {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	F___clock_gettime(m, int32(1), v304)
	mBase = m.M
	v309 = int32(4450336)
	v310 = *(*int64)(unsafe.Add(mBase, _consts[315]))
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v293)+16))
	v313 = int64(*(*int32)(unsafe.Add(mBase, uint32(v304)+8)))
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v293)+24))
	v319 = v313 + v314*int64(1000000000) - v318
	*(*int64)(unsafe.Add(mBase, _consts[315])) = v312 + v319
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v293)+8))
	if v290 != int32(1) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	m.G0 = v304 + int32(16)
	goto L47
L51:
	;
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v306)))
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = v324 + int64(1)
	goto L53
L52:
	;
	goto L53
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v306)+8)) = v322 + v319
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v306)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v306)+16)) = v329 + (v319 - v310 + v312)
	goto L50
L54:
	;
	if v341 != int32(1) {
		goto L22
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v83)+36))
	if v357 != 0 {
		goto L21
	} else {
		goto L61
	}
L57:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v346 != int32(1) {
		v527 = v286
		goto L17
	} else {
		goto L58
	}
L58:
	;
	v349 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+58)) = uint8(v349)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+59)))
	if v351 != 0 {
		v527 = v286
		goto L17
	} else {
		goto L59
	}
L59:
	;
	F_RegisterExprContextCallback(m, v20, int32(634), v67)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v355 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+59)) = uint8(v355)
	v527 = v286
	goto L17
L61:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	if v358 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v83)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+44)) = v358
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
	if v361 == int32(0) {
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
	v364 = int32(4470560)
	v365 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v67)+36))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v367
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
	if v369 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if v359 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	v374 = v369
	goto L70
L69:
	;
	if v359 == int32(0) {
		goto L20
	} else {
		goto L71
	}
L70:
	;
	v376 = F_MakeSingleTupleTableSlot(m, v374, int32(1592204))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	v372 = F_CreateTupleDescCopy(m, v359)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v374 = v372
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+48)) = v376
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v365
	goto L67
L74:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+59)))
	if v394 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L75:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
	if v385 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_tupledesc_match(m, v385, v359)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	if v388 != int32(-1) {
		goto L74
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	F_FreeTupleDesc(m, v359)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L74
L82:
	;
	F_RegisterExprContextCallback(m, v20, int32(634), v67)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
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
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	v400 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+59)) = uint8(v400)
	goto L84
L86:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	if v404 == int32(0) {
		goto L23
	} else {
		goto L87
	}
L87:
	;
	v466 = v404
	goto L18
L88:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v419
	F_errmsg(m, int32(478640), v83)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(491565), int32(686), int32(106601))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
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
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(431671), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(491565), int32(667), int32(106601))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
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
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(414266), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(491565), int32(895), int32(96735))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v482
	if v489 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v493 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v493
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+56)))
	if v495 == v493 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	F_tuplestore_end(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L112
	}
L104:
	;
	v498 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v498)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
	v501 = F_ExecFetchSlotHeapTupleDatum(m, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
	v504 = int32(*(*int16)(unsafe.Add(mBase, uint32(v503)+6)))
	if v504 <= int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v527 = v501
	goto L17
L108:
	;
	F_slot_getsomeattrs_int(m, v503, int32(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v503)+20))
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510))))
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v511)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v503)+16))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v527 = v514
	goto L17
L111:
	;
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(2)
	v522 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v522)
	v527 = v480
	goto L17
L113:
	;
	v553 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v553)
	v576 = v550
	goto L9
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v556
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(0)
	v576 = v53
	goto L9
L115:
	;
	goto L8
L116:
	;
	v590 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v592 = v590 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v592)
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v595)
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
	var v21 int32
	_ = v21
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
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
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v198 int64
	_ = v198
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v215 int64
	_ = v215
	var v219 int64
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
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
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	v9 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+149)))
	if v243 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v238 = v17
	goto L6
L8:
	;
	goto L9
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	if v26 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v27 = v23 | int32(2)
	goto L12
L11:
	;
	v27 = v23
	goto L12
L12:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+149)))
	if v28 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	if v56 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22+v31<<(uint(int32(4))%32))+88))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v43 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v44 = F_tuplesort_begin_datum(m, v35, v37, v39, v41, v43, v27)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	v52 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v53 = F_tuplesort_begin_heap(m, v22, v46, v47, v48, v49, v50, v52, v27)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v55 = v44
	goto L13
L18:
	;
	v55 = v53
	goto L13
L19:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+236))
	if v62 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v55
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+149)))
	if v89 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L22:
	;
	goto L21
L23:
	;
	goto L22
L24:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v55)+72)) = uint32(v59)
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v55)+68)) = uint8(v71)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+32))
	if v77 != 0 {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	if int64(1073741823) < v59 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if int64(1073741823) < v59 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v55)+232))
	if v65 != int32(-1) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	goto L24
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
	v80 = v79
	goto L33
L32:
	;
	v80 = v76
	goto L33
L33:
	;
	v81 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v81
	goto L23
L34:
	;
	F_tuplesort_performsort(m, v55)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L62
	}
L35:
	;
	goto L38
L36:
	;
	goto L37
L37:
	;
	goto L52
L38:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v99 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v103 = m.T0[v102].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	if v103 == int32(0) {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
	if v107&int32(2) != 0 {
		goto L34
	} else {
		goto L46
	}
L46:
	;
	v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+6)))
	if v110 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_slot_getsomeattrs_int(m, v103, int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	F_tuplesort_putdatum(m, v55, v117, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	goto L38
L52:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v129 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v133 = m.T0[v132].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	if v133 == int32(0) {
		goto L34
	} else {
		goto L59
	}
L59:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+4)))
	if v137&int32(2) != 0 {
		goto L34
	} else {
		goto L60
	}
L60:
	;
	F_tuplesort_puttupleslot(m, v55, v133)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	goto L52
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v15
	v152 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v152)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+129)) = uint8(v154)
	v156 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v158 == int32(0) {
		v238 = v55
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v161 != int32(1) {
		v238 = v55
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	v170 = v158 + v165<<(uint(int32(4))%32) + int32(8)
	v171 = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v55)+128))
	if v175 == v171 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v238 = v55
	goto L6
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v212
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v55)+112))
	v219 = base.I64_div_s(v215+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v55)+124))
	switch v221 - int32(3) {
	case 0:
		goto L81
	case 1:
		v232 = v221
		goto L78
	case 2:
		goto L80
	default:
		goto L79
	}
L67:
	;
	if v191&int32(255) != base.B2i32(v175 != int32(0)) {
		goto L73
	} else {
		goto L74
	}
L68:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v55)+96))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v55)+88))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+120)))
	v191 = v181
	v192 = v178 - v179
	goto L67
L69:
	;
	goto L70
L70:
	;
	v182 = F_LogicalTapeSetBlocks(m, v175)
	mBase = m.M
	v184 = v182 << (uint(int64(13)) % 64)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+120)))
	if v185 != 0 {
		v191 = v185
		v192 = v184
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v55)+120)) = uint8(v186)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+112)) = v184
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v55)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+124)) = v189
	v212 = v171
	goto L66
L72:
	;
	v212 = int32(1)
	goto L66
L73:
	;
	if v191&int32(1) != 0 {
		v212 = v171
		goto L66
	} else {
		goto L77
	}
L74:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v55)+112))
	if v192 <= v198 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v55)+120)) = uint8(v191)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+112)) = v192
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v55)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+124)) = v202
	if v191&int32(1) == int32(0) {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v212 = v171
	goto L66
L77:
	;
	goto L72
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v232
	goto L65
L79:
	;
	v232 = int32(0)
	goto L78
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = int32(8)
	goto L65
L81:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+69)))
	if v226 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v227 = int32(1)
	goto L84
L83:
	;
	v227 = int32(2)
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v227
	goto L65
L85:
	;
	return v242
L86:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+8))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	m.T0[v247].(func(*base.Module, int32))(m, v242)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v270 = int32(0)
	v272 = F_tuplesort_gettupleslot(m, v238, base.B2i32(v15 == int32(1)), v270, v242, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L93
	}
L89:
	;
	v252 = int32(0)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v242)+20))
	v256 = F_tuplesort_getdatum(m, v238, base.B2i32(v15 == int32(1)), v252, v253, v254, v252)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	if v256 == int32(0) {
		goto L85
	} else {
		goto L91
	}
L91:
	;
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242)+4)))
	v262 = v260 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v242)+4)) = uint16(v262)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	*(*uint16)(unsafe.Add(mBase, uint32(v242)+6)) = uint16(v265)
	goto L92
L92:
	;
	return v242
L93:
	;
	goto L85
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
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v226 int64
	_ = v226
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int64
	_ = v254
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
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
	v35 = v5
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
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = base.I64_rotl(v254, int64(32))
	F_errmsg(m, int32(692676), v17+int32(128))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L76
	}
L7:
	;
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = base.I64_rotl(v240, int64(32))
	F_errmsg(m, int32(692758), v17+int32(96))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L74
	}
L8:
	;
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = base.I64_rotl(v226, int64(32))
	F_errmsg(m, int32(692843), v17-int32(-64))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L72
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
	if v48 <= v35 {
		v54 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v54 = v50 + v35<<(uint(int32(2))%32)
	goto L11
L14:
	;
	goto L10
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v57 <= v35 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if v54 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v64 = v61 + v35<<(uint(int32(2))%32)
	if v64 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v68 != l0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v35 = v35 + int32(1)
	goto L9
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v70 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v73 = int32(4470560)
	v74 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v76
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v81 = m.T0[v80].(func(*base.Module, int32, int32, int32) int32)(m, v70, v26, v17+int32(143))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v74
	if v81 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	switch v85 {
	case 0:
		goto L28
	case 1, 2:
		goto L27
	case 3:
		goto L25
	case 4, 5:
		goto L26
	default:
		goto L24
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L69
	}
L25:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L64
	}
L26:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L59
	}
L27:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L54
	}
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v86 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+56))
	v117 = F_ExecBuildSlotValueDescription(m, v116, v112, v115, v113)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L45
	}
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+52))
	v91 = F_build_attrmap_by_name_if_req(m, v87, v89, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v106 = F_ExecGetInsertedCols(m, l1, l3)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L42
	}
L33:
	;
	if v91 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v94 = F_MakeTupleTableSlot(m, v89, int32(1592100))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v98 = l2
	goto L36
L36:
	;
	v99 = F_ExecGetInsertedCols(m, v86, l3)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	v96 = F_execute_attr_map_slot(m, v91, l2, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v98 = v96
	goto L36
L39:
	;
	v101 = F_ExecGetUpdatedCols(m, v86, l3)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v103 = F_bms_union(m, v99, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v112 = v98
	v113 = v103
	v114 = v105
	v115 = v89
	goto L29
L42:
	;
	v108 = F_ExecGetUpdatedCols(m, l1, l3)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v110 = F_bms_union(m, v106, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v112 = l2
	v113 = v110
	v114 = v19
	v115 = v20
	goto L29
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(260))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v126
	F_errmsg(m, int32(671511), v17+int32(32))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	if v117 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v117
	F_errdetail(m, int32(579397), v17+int32(16))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_errfinish(m, int32(488162), int32(2327), int32(136148))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v144 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v152
	F_errmsg(m, int32(692084), v17+int32(48))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(488162), int32(2340), int32(136148))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	if v164 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v172
	F_errmsg(m, int32(692596), v17+int32(80))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(488162), int32(2353), int32(136148))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	if v184 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v192
	F_errmsg(m, int32(692519), v17+int32(112))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(488162), int32(2365), int32(136148))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v208
	F_errmsg_internal(m, int32(58160), v17)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(488162), int32(2368), int32(136148))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errfinish(m, int32(488162), int32(2335), int32(136148))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
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
	F_errfinish(m, int32(488162), int32(2348), int32(136148))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
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
	F_errfinish(m, int32(488162), int32(2360), int32(136148))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__equalCaseWhen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v12 = F_equal(m, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = v12
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v119 int32
	_ = v119
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_MemoryContextDelete(m, v29)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L52
	}
L2:
	;
	v102 = v97
	v103 = v95
	v105 = v96
	goto L29
L3:
	;
	v95 = v4
	v96 = v4
	v97 = int32(2)
	goto L2
L4:
	;
	v95 = v91
	v96 = v44
	v97 = int32(1)
	goto L2
L5:
	;
	return
L6:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+7)))
	if v16&int32(32) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
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
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L25
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v29 = F_AllocSetContextCreateInternal(m, v24, int32(63380), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v33 = F_JsonbIteratorInit(m, v14+int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v33
	v41 = F_JsonbIteratorNext(m, v11+int32(44), v11+int32(24), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	v43 = int32(4470560)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v29
	v47 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+14)) = uint16(v47)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v51 = F_cstring_to_text_with_len(m, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	switch v41 {
	case 0:
		goto L1
	case 1:
		goto L13
	default:
		goto L3
	}
L15:
	;
	v58 = F_JsonbIteratorNext(m, v11+int32(44), v11+int32(24), int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v51
	if l2 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v65 = F_JsonbValueToJsonb(m, v11+int32(24))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v91 = v65
	goto L4
L21:
	;
	v70 = F_JsonbValueAsText(m, v11+int32(24))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v72)
	v95 = v4
	v96 = v44
	v97 = int32(0)
	goto L2
L24:
	;
	v91 = v70
	goto L4
L25:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	F_errmsg(m, int32(108656), v11)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(486214), int32(1989), int32(494115))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
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
	v105 = v137
	goto L29
L32:
	;
	v102 = int32(1)
	v103 = v175
	v105 = v174
	goto L29
L33:
	;
	goto L38
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v103
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	F_tuplestore_putvalues(m, v108, v109, v11+int32(16), v11+int32(14))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L36
	}
L35:
	;
	v174 = v105
	v175 = int32(0)
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v105
	F_MemoryContextReset(m, v29)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
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
	v134 = F_JsonbIteratorNext(m, v11+int32(44), v11+int32(24), int32(1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	v136 = int32(4470560)
	v137 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v29
	v140 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+14)) = uint16(v140)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v144 = F_cstring_to_text_with_len(m, v142, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	switch v134 {
	case 0:
		goto L1
	case 1:
		goto L40
	default:
		goto L38
	}
L42:
	;
	v151 = F_JsonbIteratorNext(m, v11+int32(44), v11+int32(24), int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v144
	if l2 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v154 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v165 = F_JsonbValueToJsonb(m, v11+int32(24))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L51
	}
L47:
	;
	v157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v157)
	goto L31
L48:
	;
	goto L49
L49:
	;
	v161 = F_JsonbValueAsText(m, v11+int32(24))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v174 = v137
	v175 = v161
	goto L32
L51:
	;
	v174 = v137
	v175 = v165
	goto L32
L52:
	;
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v188)
	m.G0 = v11 + int32(48)
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
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v61 int64
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v83 int64
	_ = v83
	var v98 float64
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
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
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
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
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
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
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
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
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
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
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v527 int32
	_ = v527
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v16 != 0 {
		v147 = v16
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L9
	} else {
		goto L125
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L9
	} else {
		goto L122
	}
L3:
	;
	m.G0 = v14 + int32(16)
	return v527
L4:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v475 == int32(0) {
		v527 = v6
		goto L3
	} else {
		goto L105
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l4
	v154 = base.B2i32(base.Ui32(l2) < base.Ui32(l3))
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L44
	} else {
		goto L45
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v17 == int32(0) {
		v527 = v6
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 < int32(32) {
		goto L4
	} else {
		goto L8
	}
L8:
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
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if base.Ui64(v47) <= base.Ui64(int64(2)) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	v39 = v33
	goto L14
L13:
	;
	v39 = v36
	goto L14
L14:
	;
	if base.F64_lt(v39, float64(1.8446744073709552e+19))&base.F64_ge(v39, float64(0)) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v45 = base.I64_trunc_f64_u(v39)
	v47 = v45
	goto L11
L16:
	;
	goto L17
L17:
	;
	v47 = int64(0)
	goto L11
L18:
	;
	v50 = int64(2)
	goto L20
L19:
	;
	v50 = v47
	goto L20
L20:
	;
	v51 = int64(1)
	if v50&(v50-v51) == int64(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = v50
	goto L23
L22:
	;
	v61 = v51 << (uint(int64(64)-base.I64_clz(v50)) % 64)
	goto L23
L23:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v61*int64(20)) {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v70 = F_MemoryContextAllocExtended(m, v23, base.I32_wrap_i64(v61)*int32(20), int32(5))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v70
	v73 = int64(1)
	if v61&(v61-v73) == int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v83 = v61
	goto L28
L27:
	;
	v83 = v73 << (uint(int64(64)-base.I64_clz(v61)) % 64)
	goto L28
L28:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v83*int64(20)) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = base.I32_wrap_i64(v83) - int32(1)
	v98 = base.F64_mul(base.F64_convert_i64_u(v83), float64(0.9))
	if base.F64_lt(v98, float64(4.294967296e+09))&base.F64_ge(v98, float64(0)) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v83 == int64(4294967296) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v104 = base.I32_trunc_f64_u(v98)
	v106 = v104
	goto L30
L32:
	;
	goto L33
L33:
	;
	v106 = int32(0)
	goto L30
L34:
	;
	v107 = int32(-85899346)
	goto L36
L35:
	;
	v107 = v106
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v25
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v110 == int32(0) {
		v147 = v25
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v113 <= int32(0) {
		v147 = v25
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v117 = int32(0)
	goto L39
L39:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v117<<(uint(int32(2))%32))))
	F_ec_add_clause_to_derives_hash(m, l1, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L41
	}
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v139 == int32(0) {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	v136 = v117 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v136 < v137 {
		v117 = v136
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v147 = v139
	goto L5
L44:
	;
	v155 = l3
	goto L46
L45:
	;
	v155 = l2
	goto L46
L46:
	;
	if l3 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v156 = v155
	goto L49
L48:
	;
	v156 = l2
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v156
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v158 = l2
	goto L52
L51:
	;
	v158 = l3
	goto L52
L52:
	;
	if l3 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v160 = v158
	goto L55
L54:
	;
	v160 = int32(0)
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v160
	v162 = int32(12)
	v168 = int32(-1636608420)
	if v14&int32(3) != 0 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	v429 = (v422 ^ v414 - base.I32_rotl(v422, int32(24))) & v428
	v432 = v427 + v429*int32(20)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	if v433 == int32(0) {
		v527 = v6
		goto L3
	} else {
		goto L96
	}
L57:
	;
	v400 = int32(14)
	v402 = v396 ^ v397 - base.I32_rotl(v396, v400)
	v406 = v402 ^ v395 - base.I32_rotl(v402, int32(11))
	v410 = v406 ^ v396 - base.I32_rotl(v406, int32(25))
	v414 = v410 ^ v402 - base.I32_rotl(v410, int32(16))
	v418 = v414 ^ v406 - base.I32_rotl(v414, int32(4))
	v422 = v418 ^ v410 - base.I32_rotl(v418, v400)
	goto L56
L58:
	;
	switch v322 - int32(1) {
	case 0:
		v388 = v313
		v389 = v314
		v390 = v318
		goto L85
	case 1:
		v381 = v313
		v382 = v314
		v383 = v318
		goto L86
	case 2:
		v374 = v313
		v375 = v314
		v376 = v318
		goto L87
	case 3:
		v368 = v314
		v369 = v318
		goto L88
	case 4:
		v364 = v314
		v365 = v318
		goto L89
	case 5:
		v358 = v314
		v359 = v318
		goto L90
	case 6:
		v352 = v314
		v353 = v318
		goto L91
	case 7:
		v347 = v318
		goto L92
	case 8:
		v342 = v318
		goto L93
	case 9:
		v337 = v318
		goto L94
	case 10:
		goto L95
	default:
		v395 = v313
		v396 = v314
		v397 = v318
		goto L57
	}
L59:
	;
	v277 = v14
	v278 = v162
	v279 = v168
	v280 = v168
	v281 = v168
	goto L82
L60:
	;
	goto L59
L61:
	;
	goto L62
L62:
	;
	goto L66
L64:
	;
	switch v220 - int32(1) {
	case 0:
		v274 = v211
		goto L71
	case 1:
		v269 = v211
		goto L72
	case 2:
		goto L73
	case 3:
		v262 = v212
		goto L74
	case 4:
		v259 = v212
		goto L75
	case 5:
		v254 = v212
		goto L76
	case 6:
		goto L77
	case 7:
		v245 = v216
		goto L78
	case 8:
		v240 = v216
		goto L79
	case 9:
		v235 = v216
		goto L80
	case 10:
		goto L81
	default:
		v395 = v211
		v396 = v212
		v397 = v216
		goto L57
	}
L66:
	;
	goto L67
L67:
	;
	v175 = v14
	v176 = v162
	v177 = v168
	v178 = v168
	v179 = v168
	goto L68
L68:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v182 = v181 + v178
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	v186 = v185 + v179
	v188 = int32(4)
	v190 = v183 + v177 - v186 ^ base.I32_rotl(v186, v188)
	v194 = v182 - v190 ^ base.I32_rotl(v190, int32(6))
	v195 = v186 + v182
	v196 = v190 + v195
	v197 = v194 + v196
	v201 = v195 - v194 ^ base.I32_rotl(v194, int32(8))
	v205 = v196 - v201 ^ base.I32_rotl(v201, int32(16))
	v209 = v197 - v205 ^ base.I32_rotl(v205, int32(19))
	v210 = v201 + v197
	v211 = v205 + v210
	v212 = v209 + v211
	v216 = v210 - v209 ^ base.I32_rotl(v209, v188)
	v217 = int32(12)
	v218 = v175 + v217
	v220 = v176 - v217
	if base.Ui32(int32(11)) < base.Ui32(v220) {
		v175 = v218
		v176 = v220
		v177 = v211
		v178 = v212
		v179 = v216
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L64
L70:
	;
	goto L69
L71:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	v395 = v274 + v275
	v396 = v212
	v397 = v216
	goto L57
L72:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	v274 = v270<<(uint(int32(8))%32) + v269
	goto L71
L73:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+2)))
	v269 = v265<<(uint(int32(16))%32) + v211
	goto L72
L74:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v395 = v263 + v211
	v396 = v262
	v397 = v216
	goto L57
L75:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
	v262 = v259 + v260
	goto L74
L76:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+5)))
	v259 = v255<<(uint(int32(8))%32) + v254
	goto L75
L77:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+6)))
	v254 = v250<<(uint(int32(16))%32) + v212
	goto L76
L78:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	v395 = v246 + v211
	v396 = v248 + v212
	v397 = v245
	goto L57
L79:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
	v245 = v241<<(uint(int32(8))%32) + v240
	goto L78
L80:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+9)))
	v240 = v236<<(uint(int32(16))%32) + v235
	goto L79
L81:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+10)))
	v235 = v231<<(uint(int32(24))%32) + v216
	goto L80
L82:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v284 = v283 + v280
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	v288 = v287 + v281
	v290 = int32(4)
	v292 = v285 + v279 - v288 ^ base.I32_rotl(v288, v290)
	v296 = v284 - v292 ^ base.I32_rotl(v292, int32(6))
	v297 = v288 + v284
	v298 = v292 + v297
	v299 = v296 + v298
	v303 = v297 - v296 ^ base.I32_rotl(v296, int32(8))
	v307 = v298 - v303 ^ base.I32_rotl(v303, int32(16))
	v311 = v299 - v307 ^ base.I32_rotl(v307, int32(19))
	v312 = v303 + v299
	v313 = v307 + v312
	v314 = v311 + v313
	v318 = v312 - v311 ^ base.I32_rotl(v311, v290)
	v319 = int32(12)
	v320 = v277 + v319
	v322 = v278 - v319
	if base.Ui32(int32(11)) < base.Ui32(v322) {
		v277 = v320
		v278 = v322
		v279 = v313
		v280 = v314
		v281 = v318
		goto L82
	} else {
		goto L84
	}
L83:
	;
	goto L58
L84:
	;
	goto L83
L85:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	v395 = v388 + v391
	v396 = v389
	v397 = v390
	goto L57
L86:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+1)))
	v388 = v384<<(uint(int32(8))%32) + v381
	v389 = v382
	v390 = v383
	goto L85
L87:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+2)))
	v381 = v377<<(uint(int32(16))%32) + v374
	v382 = v375
	v383 = v376
	goto L86
L88:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+3)))
	v374 = v370<<(uint(int32(24))%32) + v313
	v375 = v368
	v376 = v369
	goto L87
L89:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+4)))
	v368 = v364 + v366
	v369 = v365
	goto L88
L90:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+5)))
	v364 = v360<<(uint(int32(8))%32) + v358
	v365 = v359
	goto L89
L91:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+6)))
	v358 = v354<<(uint(int32(16))%32) + v352
	v359 = v353
	goto L90
L92:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+7)))
	v352 = v348<<(uint(int32(24))%32) + v314
	v353 = v347
	goto L91
L93:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+8)))
	v347 = v343<<(uint(int32(8))%32) + v342
	goto L92
L94:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+9)))
	v342 = v338<<(uint(int32(16))%32) + v337
	goto L93
L95:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+10)))
	v337 = v333<<(uint(int32(24))%32) + v318
	goto L94
L96:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v441 = v429
	v442 = v432
	goto L97
L97:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	if v450 != v438 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v442)+16))
	v527 = v463
	goto L3
L99:
	;
	goto L98
L100:
	;
	v458 = (v441 + int32(1)) & v428
	v461 = v427 + v458*int32(20)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	if v462 != 0 {
		v441 = v458
		v442 = v461
		goto L97
	} else {
		goto L104
	}
L101:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v442)+8))
	if v452 != v437 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v442)+12))
	if v454 == v436 {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	v527 = v6
	goto L3
L105:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	if v478 <= int32(0) {
		v527 = v6
		goto L3
	} else {
		goto L106
	}
L106:
	;
	v481 = int32(0)
	if v481 < v478 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v484 = v478
	goto L109
L108:
	;
	v484 = v481
	goto L109
L109:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v475)+12))
	v488 = int32(0)
	goto L110
L110:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v485+v488<<(uint(int32(2))%32))))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+108))
	if base.B2i32(l3 == int32(0))&base.B2i32(l2 == v504) != 0 {
		v527 = v503
		goto L3
	} else {
		goto L112
	}
L111:
	;
	v527 = int32(0)
	goto L3
L112:
	;
	if v504 != l2 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	if v504 != l3 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v503)+112))
	if v508 != l3 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v503)+60))
	if v510 == l4 {
		v527 = v503
		goto L3
	} else {
		goto L116
	}
L116:
	;
	goto L113
L117:
	;
	v519 = v488 + int32(1)
	if v519 != v484 {
		v488 = v519
		goto L110
	} else {
		goto L121
	}
L118:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v503)+112))
	if v513 != l2 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v503)+60))
	if v515 == l4 {
		v527 = v503
		goto L3
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	goto L111
L122:
	;
	F_errmsg_internal(m, int32(394012), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(321223), int32(327), int32(335250))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L9
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errmsg_internal(m, int32(394012), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L9
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(321223), int32(327), int32(335250))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L9
	} else {
		goto L127
	}
L127:
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1036]))
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
	v30 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v35 = F_AllocSetContextCreateInternal(m, v30, int32(63321), int32(0), int32(8192), int32(8388608))
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
	v49 = int32(4470560)
	v50 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v35
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
	F_errmsg(m, int32(226177), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(486214), int32(2235), int32(494093))
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
	F_errmsg(m, int32(109155), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(486214), int32(2239), int32(494093))
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
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v114
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
	v145 = int32(4470560)
	v146 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v35
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
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
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
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
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int64
	_ = v321
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
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
	var v369 int32
	_ = v369
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	F_init_work(m, v10+int32(-12), l1, l4, v10+int32(-56))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = int32(0)
	if l1 == v22 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v86 = int32(1)
	v87 = v81 + v86
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v92 = v90 & v86
	if v92 != 0 {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v81 = l2
	v82 = v22
	goto L3
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v81 = l2
	v82 = v22
	goto L3
L8:
	;
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	goto L10
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v32 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v63 = int32(1)
	if v32&v63 != 0 {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v35 = int32(4)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v37&int32(254) == int32(2) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v50 = int32(1)
	if v32&v50 != 0 {
		v62 = int32(base.Ui32(v32)>>(uint(v50)%32)) - v50
		goto L11
	} else {
		goto L21
	}
L15:
	;
	v46 = v35
	goto L17
L16:
	;
	v46 = base.B2i32(v37 == int32(18)) << (uint(v35) % 32)
	goto L17
L17:
	;
	if v37 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = v35
	goto L20
L19:
	;
	v49 = v46
	goto L20
L20:
	;
	v62 = v49
	goto L11
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v62 = int32(base.Ui32(v56)>>(uint(int32(2))%32)) - int32(4)
	goto L11
L22:
	;
	v67 = v63
	goto L24
L23:
	;
	v67 = int32(4)
	goto L24
L24:
	;
	v68 = l2 + v67
	v70 = F_pg_do_encoding_conversion(m, v68, v62, v31, int32(6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v70 != v68 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v73 = F_cstring_to_text(m, v70)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v77 = l2
	goto L28
L28:
	;
	if v77 != l2 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F_pfree(m, v70)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v77 = v73
	goto L28
L31:
	;
	v80 = v77
	goto L33
L32:
	;
	v80 = int32(0)
	goto L33
L33:
	;
	v81 = v77
	v82 = v80
	goto L3
L34:
	;
	v93 = v87
	goto L36
L35:
	;
	v93 = v81 + int32(4)
	goto L36
L36:
	;
	if v90 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v122 = F_mbuf_create_from_data(m, v93, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L48
	}
L38:
	;
	v96 = int32(4)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v98&int32(254) == int32(2) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v111 = int32(1)
	if v92 != 0 {
		v121 = int32(base.Ui32(v90)>>(uint(v111)%32)) - v111
		goto L37
	} else {
		goto L47
	}
L41:
	;
	v107 = v96
	goto L43
L42:
	;
	v107 = base.B2i32(v98 == int32(18)) << (uint(v96) % 32)
	goto L43
L43:
	;
	if v98 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v110 = v96
	goto L46
L45:
	;
	v110 = v107
	goto L46
L46:
	;
	v121 = v110
	goto L37
L47:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v121 = int32(base.Ui32(v115)>>(uint(int32(2))%32)) - int32(4)
	goto L37
L48:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v124 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v154 = F_mbuf_create(m, v151+int32(128))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	v127 = int32(6)
	v129 = int32(18)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v131 == v129 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v143 = int32(1)
	if v124&v143 != 0 {
		v151 = int32(base.Ui32(v124) >> (uint(v143) % 32))
		goto L49
	} else {
		goto L62
	}
L53:
	;
	v134 = v129
	goto L55
L54:
	;
	v134 = int32(2)
	goto L55
L55:
	;
	if v131&int32(254) == int32(2) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v139 = v127
	goto L58
L57:
	;
	v139 = v134
	goto L58
L58:
	;
	if v131 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v142 = v127
	goto L61
L60:
	;
	v142 = v139
	goto L61
L61:
	;
	v151 = v142
	goto L49
L62:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v151 = int32(base.Ui32(v147) >> (uint(int32(2)) % 32))
	goto L49
L63:
	;
	v159 = F_mbuf_append(m, v154, v10+int32(-4), int32(4))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if l0 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if int32(0) <= v256 {
		goto L105
	} else {
		goto L106
	}
L66:
	;
	v161 = int32(1)
	v162 = l3 + v161
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v167 = v165 & v161
	if v167 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v208 = int32(1)
	v209 = l3 + v208
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v214 = v212 & v208
	if v214 != 0 {
		goto L86
	} else {
		goto L87
	}
L69:
	;
	v168 = v162
	goto L71
L70:
	;
	v168 = l3 + int32(4)
	goto L71
L71:
	;
	if v165 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v197 = F_mbuf_create_from_data(m, v168, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L83
	}
L73:
	;
	v171 = int32(4)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v173&int32(254) == int32(2) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v186 = int32(1)
	if v167 != 0 {
		v196 = int32(base.Ui32(v165)>>(uint(v186)%32)) - v186
		goto L72
	} else {
		goto L82
	}
L76:
	;
	v182 = v171
	goto L78
L77:
	;
	v182 = base.B2i32(v173 == int32(18)) << (uint(v171) % 32)
	goto L78
L78:
	;
	if v173 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v185 = v171
	goto L81
L80:
	;
	v185 = v182
	goto L81
L81:
	;
	v196 = v185
	goto L72
L82:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v196 = int32(base.Ui32(v190)>>(uint(int32(2))%32)) - int32(4)
	goto L72
L83:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v200 = int32(0)
	v203 = F_pgp_set_pubkey(m, v199, v197, v200, v200, v200)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v205 = F_mbuf_free(m, v197)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v256 = v203
	goto L65
L86:
	;
	v215 = v209
	goto L88
L87:
	;
	v215 = l3 + int32(4)
	goto L88
L88:
	;
	if v212 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v245 = int32(-13)
	if v215 == int32(0) {
		v253 = v245
		goto L101
	} else {
		goto L102
	}
L90:
	;
	v218 = int32(4)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v220&int32(254) == int32(2) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v233 = int32(1)
	if v214 != 0 {
		v243 = int32(base.Ui32(v212)>>(uint(v233)%32)) - v233
		goto L89
	} else {
		goto L99
	}
L93:
	;
	v229 = v218
	goto L95
L94:
	;
	v229 = base.B2i32(v220 == int32(18)) << (uint(v218) % 32)
	goto L95
L95:
	;
	if v220 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v232 = v218
	goto L98
L97:
	;
	v232 = v229
	goto L98
L98:
	;
	v243 = v232
	goto L89
L99:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v243 = int32(base.Ui32(v237)>>(uint(int32(2))%32)) - int32(4)
	goto L89
L100:
	;
	v256 = v253
	goto L65
L101:
	;
	goto L100
L102:
	;
	if v243 <= int32(0) {
		v253 = v245
		goto L101
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+128)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v207)+124)) = v215
	v253 = int32(0)
	goto L101
L104:
	;
	v316 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)) = uint16(v316)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(-8)))) = v319
	v321 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v154)+8)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v154))) = v321
	goto L137
L105:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v262 = F_pgp_encrypt(m, v261, v122, v154)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v266 = v256
	goto L107
L107:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v267 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	if v262 == int32(0) {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v266 = v262
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1454])) = int32(0)
	goto L113
L111:
	;
	goto L112
L112:
	;
	if v82 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L112
L114:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v272 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	goto L116
L116:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v305 = F_pgp_free(m, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L133
	}
L117:
	;
	v300 = F___memset(m, v82, int32(0), v299)
	mBase = m.M
	goto L131
L118:
	;
	v275 = int32(6)
	v277 = int32(18)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v279 == v277 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v291 = int32(1)
	if v272&v291 != 0 {
		v299 = int32(base.Ui32(v272) >> (uint(v291) % 32))
		goto L117
	} else {
		goto L130
	}
L121:
	;
	v282 = v277
	goto L123
L122:
	;
	v282 = int32(2)
	goto L123
L123:
	;
	if v279&int32(254) == int32(2) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v287 = v275
	goto L126
L125:
	;
	v287 = v282
	goto L126
L126:
	;
	if v279 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v290 = v275
	goto L129
L128:
	;
	v290 = v287
	goto L129
L129:
	;
	v299 = v290
	goto L117
L130:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v299 = int32(base.Ui32(v295) >> (uint(int32(2)) % 32))
	goto L117
L131:
	;
	F_pfree(m, v82)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	goto L116
L133:
	;
	v307 = F_mbuf_free(m, v122)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v309 = F_mbuf_free(m, v154)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_px_THROW_ERROR(m, v266)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v326))) = (v318 - v319) << (uint(int32(2)) % 32)
	if v82 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v331 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	goto L140
L140:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v364 = F_pgp_free(m, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L157
	}
L141:
	;
	v359 = F___memset(m, v82, int32(0), v358)
	mBase = m.M
	goto L155
L142:
	;
	v334 = int32(6)
	v336 = int32(18)
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v338 == v336 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	v350 = int32(1)
	if v331&v350 != 0 {
		v358 = int32(base.Ui32(v331) >> (uint(v350) % 32))
		goto L141
	} else {
		goto L154
	}
L145:
	;
	v341 = v336
	goto L147
L146:
	;
	v341 = int32(2)
	goto L147
L147:
	;
	if v338&int32(254) == int32(2) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v346 = v334
	goto L150
L149:
	;
	v346 = v341
	goto L150
L150:
	;
	if v338 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v349 = v334
	goto L153
L152:
	;
	v349 = v346
	goto L153
L153:
	;
	v358 = v349
	goto L141
L154:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v358 = int32(base.Ui32(v354) >> (uint(int32(2)) % 32))
	goto L141
L155:
	;
	F_pfree(m, v82)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	goto L140
L157:
	;
	v366 = F_mbuf_free(m, v122)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v368 = F_mbuf_free(m, v154)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1454])) = int32(0)
	goto L160
L160:
	;
	m.G0 = v12 - int32(-64)
	return v326
}
func F_encrypt_password(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
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
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v615 int32
	_ = v615
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v639 int32
	_ = v639
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v732 int32
	_ = v732
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v778 int32
	_ = v778
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v845 int32
	_ = v845
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v871 int32
	_ = v871
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v893 int32
	_ = v893
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v917 int32
	_ = v917
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1096 int32
	_ = v1096
	var v1102 int32
	_ = v1102
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int64
	_ = v1191
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1247 int32
	_ = v1247
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1267 int32
	_ = v1267
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v4
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v17 != int32(109) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v13 + int32(112)
	return v1102
L2:
	;
	v1329 = F_parse_scram_secret(m, v1102, v13+int32(104), v13+int32(96), v13+int32(100), v13+int32(108), v13-int32(-64), v13+int32(32))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L52
	} else {
		goto L324
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L52
	} else {
		goto L319
	}
L4:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[373])))
	if v1109 != int32(1) {
		goto L1
	} else {
		goto L268
	}
L5:
	;
	if v1031 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
	switch l0 {
	case 0:
		goto L57
	case 1:
		goto L58
	case 2:
		goto L56
	default:
		v1102 = v4
		goto L4
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
	v193 = F_pstrdup(m, l2)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L52
	} else {
		goto L55
	}
L8:
	;
	v185 = F_parse_scram_secret(m, l2, v13+int32(104), v13+int32(96), v13+int32(100), v13+int32(108), v13-int32(-64), v13+int32(32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L52
	} else {
		goto L53
	}
L9:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v20 != int32(100) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
	if v23 != int32(53) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if l2&int32(3) == int32(0) {
		v49 = l2
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v82 != int32(35) {
		goto L8
	} else {
		goto L29
	}
L13:
	;
	v82 = v74 - l2
	goto L12
L14:
	;
	v53 = v49
	goto L23
L15:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v33 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v82 = int32(0)
	goto L12
L17:
	;
	goto L18
L18:
	;
	v38 = l2
	goto L19
L19:
	;
	v42 = v38 + int32(1)
	if v42&int32(3) == int32(0) {
		v49 = v42
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v74 = v42
	goto L13
L21:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v47 != 0 {
		v38 = v42
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v62 = int32(-2139062144)
	if (int32(16843008)-v59|v59)&v62 == v62 {
		v53 = v53 + int32(4)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v68 = v53
	goto L26
L25:
	;
	goto L24
L26:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v72 != 0 {
		v68 = v68 + int32(1)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v74 = v68
	goto L13
L28:
	;
	goto L27
L29:
	;
	v86 = l2 + int32(3)
	v87 = int32(333853)
	v91 = m.G0
	v93 = v91 - int32(32)
	v94 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+24)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v93)+16)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v93)+8)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v93))) = v94
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _consts[371])))
	if v102 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v170 == int32(32) {
		goto L7
	} else {
		goto L51
	}
L31:
	;
	v170 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, _consts[372])))
	if v106 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v110 = v86
	goto L37
L35:
	;
	goto L36
L36:
	;
	v120 = v87
	v121 = v102
	goto L40
L37:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v116 == v102 {
		v110 = v110 + int32(1)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v170 = v110 - v86
	goto L30
L39:
	;
	goto L38
L40:
	;
	v128 = v93 + int32(base.Ui32(v121)>>(uint(int32(3))%32))&int32(28)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v130 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v129 | v130<<(uint(v121)%32)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	if v134 != 0 {
		v120 = v120 + v130
		v121 = v134
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v137 == int32(0) {
		v162 = v86
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v170 = v162 - v86
	goto L30
L44:
	;
	v141 = v86
	v142 = v137
	goto L45
L45:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v93+int32(base.Ui32(v142)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v150)>>(uint(v142)%32))&int32(1) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v162 = v158
	goto L43
L47:
	;
	v162 = v141
	goto L43
L48:
	;
	goto L49
L49:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	v158 = v141 + int32(1)
	if v156 != 0 {
		v141 = v158
		v142 = v156
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	goto L8
L52:
	;
	return int32(0)
L53:
	;
	if v185 == int32(0) {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	goto L7
L55:
	;
	v1031 = v193
	goto L5
L56:
	;
	v290 = m.G0
	v292 = v290 - int32(48)
	m.G0 = v292
	*(*int32)(unsafe.Add(mBase, uint32(v292)+12)) = int32(0)
	v298 = F_pg_saslprep(m, l2, v292+int32(44))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L52
	} else {
		goto L86
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L52
	} else {
		goto L82
	}
L58:
	;
	v198 = F_palloc(m, int32(36))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L52
	} else {
		goto L59
	}
L59:
	;
	if l1&int32(3) == int32(0) {
		v223 = l1
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v259 = F_pg_md5_encrypt(m, l2, l1, v256, v198, v13+int32(28))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L52
	} else {
		goto L77
	}
L61:
	;
	v256 = v248 - l1
	goto L60
L62:
	;
	v227 = v223
	goto L71
L63:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v207 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v256 = int32(0)
	goto L60
L65:
	;
	goto L66
L66:
	;
	v212 = l1
	goto L67
L67:
	;
	v216 = v212 + int32(1)
	if v216&int32(3) == int32(0) {
		v223 = v216
		goto L62
	} else {
		goto L69
	}
L68:
	;
	v248 = v216
	goto L61
L69:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v221 != 0 {
		v212 = v216
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v236 = int32(-2139062144)
	if (int32(16843008)-v233|v233)&v236 == v236 {
		v227 = v227 + int32(4)
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v242 = v227
	goto L74
L73:
	;
	goto L72
L74:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v246 != 0 {
		v242 = v242 + int32(1)
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v248 = v242
	goto L61
L76:
	;
	goto L75
L77:
	;
	if v259 != 0 {
		v1031 = v198
		goto L5
	} else {
		goto L78
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L52
	} else {
		goto L79
	}
L79:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v265
	F_errmsg_internal(m, int32(200936), v13+int32(16))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L52
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(484591), int32(141), int32(413014))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L52
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
	F_errmsg_internal(m, int32(661391), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L52
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(484591), int32(149), int32(413014))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L52
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
	v1031 = v510
	goto L5
L86:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v292)+44))
	v301 = int32(16)
	v304 = int32(0)
	v308 = m.G0
	v310 = v308 - v301
	m.G0 = v310
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v304
	v316 = F_open(m, int32(283766), v304, v310)
	mBase = m.M
	if v316 != int32(-1) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v349 != 0 {
		goto L100
	} else {
		goto L101
	}
L88:
	;
	goto L92
L89:
	;
	v349 = v304
	goto L90
L90:
	;
	m.G0 = v310 + int32(16)
	goto L87
L91:
	;
	v344 = F_close(m, v316)
	mBase = m.M
	v349 = v342
	goto L90
L92:
	;
	v322 = v292 + v301
	v323 = v301
	goto L93
L93:
	;
	v328 = F_read(m, v316, v322, v323)
	mBase = m.M
	if v328 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v342 = int32(1)
	goto L91
L95:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v332 == int32(27) {
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v337 = v323 - v328
	if v337 != 0 {
		v322 = v322 + v328
		v323 = v337
		goto L93
	} else {
		goto L99
	}
L98:
	;
	v342 = int32(0)
	goto L91
L99:
	;
	goto L94
L100:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _consts[374]))
	v356 = m.G0
	v358 = v356 - int32(176)
	m.G0 = v358
	if v298 != 0 {
		goto L108
	} else {
		goto L109
	}
L101:
	;
	goto L102
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L52
	} else {
		goto L243
	}
L103:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v292)+44))
	if v1006 != 0 {
		goto L239
	} else {
		goto L240
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = int32(21285)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L52
	} else {
		goto L236
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = int32(22143)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L52
	} else {
		goto L233
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = int32(97304)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L52
	} else {
		goto L230
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L52
	} else {
		goto L227
	}
L108:
	;
	v360 = l2
	goto L110
L109:
	;
	v360 = v300
	goto L110
L110:
	;
	v363 = int32(16)
	v364 = v292 + v363
	v369 = v292 + int32(12)
	v370 = F_scram_SaltedPassword(m, v360, int32(3), int32(32), v364, v363, v355, v358+int32(144), v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L52
	} else {
		goto L111
	}
L111:
	;
	if v370 < int32(0) {
		goto L107
	} else {
		goto L112
	}
L112:
	;
	v379 = F_pg_hmac_create(m, int32(3))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L52
	} else {
		goto L114
	}
L113:
	;
	if v462 < int32(0) {
		goto L107
	} else {
		goto L160
	}
L114:
	;
	if v379 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	goto L119
L116:
	;
	goto L117
L117:
	;
	v407 = F_pg_hmac_init(m, v379, v358+int32(144), int32(32))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L52
	} else {
		goto L133
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = int32(13796)
	v462 = int32(-1)
	goto L113
L119:
	;
	goto L118
L131:
	;
	F_pg_hmac_free(m, v379)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L52
	} else {
		goto L159
	}
L132:
	;
	if v379 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L133:
	;
	if v407 < int32(0) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	if v379 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v427 < int32(0) {
		goto L132
	} else {
		goto L142
	}
L136:
	;
	v427 = int32(-1)
	goto L135
L137:
	;
	goto L138
L138:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v417 = F_pg_cryptohash_update(m, v416, int32(22427), int32(10))
	mBase = m.M
	if int32(0) <= v417 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v427 = int32(0)
	goto L135
L140:
	;
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+8)) = int32(2)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v424 = F_pg_cryptohash_error(m, v423)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v379)+12)) = v424
	v427 = int32(-1)
	goto L135
L142:
	;
	v431 = F_pg_hmac_final(m, v379, v358+int32(112), int32(32))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L52
	} else {
		goto L143
	}
L143:
	;
	if int32(0) <= v431 {
		goto L131
	} else {
		goto L144
	}
L144:
	;
	goto L132
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v454
	F_pg_hmac_free(m, v379)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L52
	} else {
		goto L158
	}
L146:
	;
	v454 = int32(13796)
	goto L145
L147:
	;
	goto L148
L148:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	if v439 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v451 = v439
	goto L151
L150:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	if v443 == int32(2) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v454 = v451
	goto L145
L152:
	;
	v446 = int32(209285)
	goto L154
L153:
	;
	v446 = int32(127867)
	goto L154
L154:
	;
	if v443 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v449 = int32(13796)
	goto L157
L156:
	;
	v449 = v446
	goto L157
L157:
	;
	v451 = v449
	goto L151
L158:
	;
	v462 = int32(-1)
	goto L113
L159:
	;
	v462 = int32(0)
	goto L113
L160:
	;
	v466 = v358 + int32(112)
	v471 = F_scram_H(m, v466, int32(3), int32(32), v466, v369)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L52
	} else {
		goto L161
	}
L161:
	;
	if v471 < int32(0) {
		goto L107
	} else {
		goto L162
	}
L162:
	;
	v481 = F_scram_ServerKey(m, v358+int32(144), int32(3), int32(32), v358+int32(80), v369)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L52
	} else {
		goto L163
	}
L163:
	;
	if v481 < int32(0) {
		goto L107
	} else {
		goto L164
	}
L164:
	;
	v489 = base.I32_div_s(int32(18), int32(3))
	v491 = v489 << (uint(int32(2)) % 32)
	goto L165
L165:
	;
	v496 = base.I32_div_s(int32(34), int32(3))
	v498 = v496 << (uint(int32(2)) % 32)
	goto L166
L166:
	;
	v504 = base.I32_div_s(int32(34), int32(3))
	v506 = v504 << (uint(int32(2)) % 32)
	goto L167
L167:
	;
	v510 = F_palloc(m, v491+v498+v506+int32(28))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L52
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+64)) = v355
	v517 = F_pg_sprintf(m, v510, int32(538226), v358-int32(-64))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L52
	} else {
		goto L169
	}
L169:
	;
	v519 = v517 + v510
	v523 = v292 + int32(32)
	if base.Ui32(v364) < base.Ui32(v523) {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	if v648 < int32(0) {
		goto L106
	} else {
		goto L188
	}
L171:
	;
	v639 = F___memset(m, v519, int32(0), v491)
	mBase = m.M
	v648 = int32(-1)
	goto L170
L172:
	;
	if v491 < v582-v519+int32(4) {
		goto L171
	} else {
		goto L184
	}
L173:
	;
	v527 = v364
	v528 = int32(0)
	v531 = v519
	v532 = int32(2)
	goto L176
L174:
	;
	v593 = v519
	goto L175
L175:
	;
	v648 = v593 - v519
	goto L170
L176:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527))))
	v538 = v534<<(uint(v532<<(uint(int32(3))%32))%32) | v528
	if int32(0) < v532 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	if v583 != int32(2) {
		goto L172
	} else {
		goto L183
	}
L178:
	;
	v581 = v538
	v582 = v531
	v583 = v532 - int32(1)
	goto L180
L179:
	;
	if v491 < v531-v519+int32(4) {
		goto L171
	} else {
		goto L181
	}
L180:
	;
	v585 = v527 + int32(1)
	if v585 != v523 {
		v527 = v585
		v528 = v581
		v531 = v582
		v532 = v583
		goto L176
	} else {
		goto L182
	}
L181:
	;
	v547 = int32(63)
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538&v547)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+3)) = uint8(v551)
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v538)>>(uint(int32(6))%32))&v547)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+2)) = uint8(v559)
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v538)>>(uint(int32(12))%32))&v547)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+1)) = uint8(v567)
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v538)>>(uint(int32(18))%32))&v547)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v531))) = uint8(v575)
	v581 = int32(0)
	v582 = v531 + int32(4)
	v583 = int32(2)
	goto L180
L182:
	;
	goto L177
L183:
	;
	v593 = v582
	goto L175
L184:
	;
	v603 = int32(63)
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v581)>>(uint(int32(12))%32))&v603)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+1)) = uint8(v607)
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v581)>>(uint(int32(18))%32))&v603)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v582))) = uint8(v615)
	if v583 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v581)>>(uint(int32(6))%32))&int32(63))+uint32(_consts[375]))))
	v627 = v626
	goto L187
L186:
	;
	v627 = int32(61)
	goto L187
L187:
	;
	v628 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+3)) = uint8(v628)
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+2)) = uint8(v627)
	v648 = v582 + int32(4) - v519
	goto L170
L188:
	;
	v651 = v648 + v519
	v652 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v651))) = uint8(v652)
	v655 = v358 + int32(112)
	v658 = v651 + int32(1)
	v662 = v358 + int32(144)
	if base.Ui32(v655) < base.Ui32(v662) {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	if v787 < int32(0) {
		goto L105
	} else {
		goto L207
	}
L190:
	;
	v778 = F___memset(m, v658, int32(0), v498)
	mBase = m.M
	v787 = int32(-1)
	goto L189
L191:
	;
	if v498 < v721-v658+int32(4) {
		goto L190
	} else {
		goto L203
	}
L192:
	;
	v666 = v655
	v667 = int32(0)
	v670 = v658
	v671 = int32(2)
	goto L195
L193:
	;
	v732 = v658
	goto L194
L194:
	;
	v787 = v732 - v658
	goto L189
L195:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	v677 = v673<<(uint(v671<<(uint(int32(3))%32))%32) | v667
	if int32(0) < v671 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	if v722 != int32(2) {
		goto L191
	} else {
		goto L202
	}
L197:
	;
	v720 = v677
	v721 = v670
	v722 = v671 - int32(1)
	goto L199
L198:
	;
	if v498 < v670-v658+int32(4) {
		goto L190
	} else {
		goto L200
	}
L199:
	;
	v724 = v666 + int32(1)
	if v724 != v662 {
		v666 = v724
		v667 = v720
		v670 = v721
		v671 = v722
		goto L195
	} else {
		goto L201
	}
L200:
	;
	v686 = int32(63)
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677&v686)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v670)+3)) = uint8(v690)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v677)>>(uint(int32(6))%32))&v686)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v670)+2)) = uint8(v698)
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v677)>>(uint(int32(12))%32))&v686)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v670)+1)) = uint8(v706)
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v677)>>(uint(int32(18))%32))&v686)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v670))) = uint8(v714)
	v720 = int32(0)
	v721 = v670 + int32(4)
	v722 = int32(2)
	goto L199
L201:
	;
	goto L196
L202:
	;
	v732 = v721
	goto L194
L203:
	;
	v742 = int32(63)
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v720)>>(uint(int32(12))%32))&v742)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v721)+1)) = uint8(v746)
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v720)>>(uint(int32(18))%32))&v742)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v721))) = uint8(v754)
	if v722 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v720)>>(uint(int32(6))%32))&int32(63))+uint32(_consts[375]))))
	v766 = v765
	goto L206
L205:
	;
	v766 = int32(61)
	goto L206
L206:
	;
	v767 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v721)+3)) = uint8(v767)
	*(*uint8)(unsafe.Add(mBase, uint32(v721)+2)) = uint8(v766)
	v787 = v721 + int32(4) - v658
	goto L189
L207:
	;
	v790 = v658 + v787
	v791 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v790))) = uint8(v791)
	v794 = v358 + int32(80)
	v797 = v790 + int32(1)
	v801 = v358 + int32(112)
	if base.Ui32(v794) < base.Ui32(v801) {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	if v926 < int32(0) {
		goto L104
	} else {
		goto L226
	}
L209:
	;
	v917 = F___memset(m, v797, int32(0), v506)
	mBase = m.M
	v926 = int32(-1)
	goto L208
L210:
	;
	if v506 < v860-v797+int32(4) {
		goto L209
	} else {
		goto L222
	}
L211:
	;
	v805 = v794
	v806 = int32(0)
	v809 = v797
	v810 = int32(2)
	goto L214
L212:
	;
	v871 = v797
	goto L213
L213:
	;
	v926 = v871 - v797
	goto L208
L214:
	;
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805))))
	v816 = v812<<(uint(v810<<(uint(int32(3))%32))%32) | v806
	if int32(0) < v810 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	if v861 != int32(2) {
		goto L210
	} else {
		goto L221
	}
L216:
	;
	v859 = v816
	v860 = v809
	v861 = v810 - int32(1)
	goto L218
L217:
	;
	if v506 < v809-v797+int32(4) {
		goto L209
	} else {
		goto L219
	}
L218:
	;
	v863 = v805 + int32(1)
	if v863 != v801 {
		v805 = v863
		v806 = v859
		v809 = v860
		v810 = v861
		goto L214
	} else {
		goto L220
	}
L219:
	;
	v825 = int32(63)
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816&v825)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v809)+3)) = uint8(v829)
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v816)>>(uint(int32(6))%32))&v825)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v809)+2)) = uint8(v837)
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v816)>>(uint(int32(12))%32))&v825)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v809)+1)) = uint8(v845)
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v816)>>(uint(int32(18))%32))&v825)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v809))) = uint8(v853)
	v859 = int32(0)
	v860 = v809 + int32(4)
	v861 = int32(2)
	goto L218
L220:
	;
	goto L215
L221:
	;
	v871 = v860
	goto L213
L222:
	;
	v881 = int32(63)
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v859)>>(uint(int32(12))%32))&v881)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v860)+1)) = uint8(v885)
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v859)>>(uint(int32(18))%32))&v881)+uint32(_consts[375]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v860))) = uint8(v893)
	if v861 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v859)>>(uint(int32(6))%32))&int32(63))+uint32(_consts[375]))))
	v905 = v904
	goto L225
L224:
	;
	v905 = int32(61)
	goto L225
L225:
	;
	v906 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v860)+3)) = uint8(v906)
	*(*uint8)(unsafe.Add(mBase, uint32(v860)+2)) = uint8(v905)
	v926 = v860 + int32(4) - v797
	goto L208
L226:
	;
	v930 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v797+v926))) = uint8(v930)
	m.G0 = v358 + int32(176)
	goto L103
L227:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	*(*int32)(unsafe.Add(mBase, uint32(v358))) = v942
	F_errmsg_internal(m, int32(196007), v358)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L52
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(487737), int32(245), int32(105467))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L52
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+16)) = v958
	F_errmsg_internal(m, int32(203086), v358+int32(16))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L52
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(487737), int32(286), int32(105467))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L52
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+32)) = v976
	F_errmsg_internal(m, int32(203086), v358+int32(32))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L52
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(487737), int32(302), int32(105467))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L52
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+48)) = v994
	F_errmsg_internal(m, int32(203086), v358+int32(48))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L52
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(487737), int32(319), int32(105467))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L52
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_pfree(m, v1006)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L52
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	m.G0 = v292 + int32(48)
	goto L85
L242:
	;
	goto L241
L243:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L52
	} else {
		goto L244
	}
L244:
	;
	F_errmsg(m, int32(97250), int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L52
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(488710), int32(504), int32(105461))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L52
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	v1102 = int32(0)
	goto L4
L248:
	;
	goto L249
L249:
	;
	if v1031&int32(3) == int32(0) {
		v1063 = v1031
		goto L252
	} else {
		goto L253
	}
L250:
	;
	if base.Ui32(int32(513)) <= base.Ui32(v1096) {
		goto L3
	} else {
		goto L267
	}
L251:
	;
	v1096 = v1088 - v1031
	goto L250
L252:
	;
	v1067 = v1063
	goto L261
L253:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1031))))
	if v1047 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1096 = int32(0)
	goto L250
L255:
	;
	goto L256
L256:
	;
	v1052 = v1031
	goto L257
L257:
	;
	v1056 = v1052 + int32(1)
	if v1056&int32(3) == int32(0) {
		v1063 = v1056
		goto L252
	} else {
		goto L259
	}
L258:
	;
	v1088 = v1056
	goto L251
L259:
	;
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056))))
	if v1061 != 0 {
		v1052 = v1056
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1067)))
	v1076 = int32(-2139062144)
	if (int32(16843008)-v1073|v1073)&v1076 == v1076 {
		v1067 = v1067 + int32(4)
		goto L261
	} else {
		goto L263
	}
L262:
	;
	v1082 = v1067
	goto L264
L263:
	;
	goto L262
L264:
	;
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082))))
	if v1086 != 0 {
		v1082 = v1082 + int32(1)
		goto L264
	} else {
		goto L266
	}
L265:
	;
	v1088 = v1082
	goto L251
L266:
	;
	goto L265
L267:
	;
	v1102 = v1031
	goto L4
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = int32(0)
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102))))
	if v1114 != int32(109) {
		goto L2
	} else {
		goto L269
	}
L269:
	;
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102)+1)))
	if v1117 != int32(100) {
		goto L2
	} else {
		goto L270
	}
L270:
	;
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102)+2)))
	if v1120 != int32(53) {
		goto L2
	} else {
		goto L271
	}
L271:
	;
	if v1102&int32(3) == int32(0) {
		v1146 = v1102
		goto L274
	} else {
		goto L275
	}
L272:
	;
	if v1179 != int32(35) {
		goto L2
	} else {
		goto L289
	}
L273:
	;
	v1179 = v1171 - v1102
	goto L272
L274:
	;
	v1150 = v1146
	goto L283
L275:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102))))
	if v1130 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1179 = int32(0)
	goto L272
L277:
	;
	goto L278
L278:
	;
	v1135 = v1102
	goto L279
L279:
	;
	v1139 = v1135 + int32(1)
	if v1139&int32(3) == int32(0) {
		v1146 = v1139
		goto L274
	} else {
		goto L281
	}
L280:
	;
	v1171 = v1139
	goto L273
L281:
	;
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1139))))
	if v1144 != 0 {
		v1135 = v1139
		goto L279
	} else {
		goto L282
	}
L282:
	;
	goto L280
L283:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1150)))
	v1159 = int32(-2139062144)
	if (int32(16843008)-v1156|v1156)&v1159 == v1159 {
		v1150 = v1150 + int32(4)
		goto L283
	} else {
		goto L285
	}
L284:
	;
	v1165 = v1150
	goto L286
L285:
	;
	goto L284
L286:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165))))
	if v1169 != 0 {
		v1165 = v1165 + int32(1)
		goto L286
	} else {
		goto L288
	}
L287:
	;
	v1171 = v1165
	goto L273
L288:
	;
	goto L287
L289:
	;
	v1183 = v1102 + int32(3)
	v1184 = int32(333853)
	v1188 = m.G0
	v1190 = v1188 - int32(32)
	v1191 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1190)+24)) = v1191
	*(*int64)(unsafe.Add(mBase, uint32(v1190)+16)) = v1191
	*(*int64)(unsafe.Add(mBase, uint32(v1190)+8)) = v1191
	*(*int64)(unsafe.Add(mBase, uint32(v1190))) = v1191
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, _consts[371])))
	if v1199 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	if v1267 != int32(32) {
		goto L2
	} else {
		goto L311
	}
L291:
	;
	v1267 = int32(0)
	goto L290
L292:
	;
	goto L293
L293:
	;
	v1203 = int32(*(*uint8)(unsafe.Add(mBase, _consts[372])))
	if v1203 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1207 = v1183
	goto L297
L295:
	;
	goto L296
L296:
	;
	v1217 = v1184
	v1218 = v1199
	goto L300
L297:
	;
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1207))))
	if v1213 == v1199 {
		v1207 = v1207 + int32(1)
		goto L297
	} else {
		goto L299
	}
L298:
	;
	v1267 = v1207 - v1183
	goto L290
L299:
	;
	goto L298
L300:
	;
	v1225 = v1190 + int32(base.Ui32(v1218)>>(uint(int32(3))%32))&int32(28)
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1225)))
	v1227 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1225))) = v1226 | v1227<<(uint(v1218)%32)
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1217)+1)))
	if v1231 != 0 {
		v1217 = v1217 + v1227
		v1218 = v1231
		goto L300
	} else {
		goto L302
	}
L301:
	;
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1183))))
	if v1234 == int32(0) {
		v1259 = v1183
		goto L303
	} else {
		goto L304
	}
L302:
	;
	goto L301
L303:
	;
	v1267 = v1259 - v1183
	goto L290
L304:
	;
	v1238 = v1183
	v1239 = v1234
	goto L305
L305:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1190+int32(base.Ui32(v1239)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v1247)>>(uint(v1239)%32))&int32(1) == int32(0) {
		goto L307
	} else {
		goto L308
	}
L306:
	;
	v1259 = v1255
	goto L303
L307:
	;
	v1259 = v1238
	goto L303
L308:
	;
	goto L309
L309:
	;
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1238)+1)))
	v1255 = v1238 + int32(1)
	if v1253 != 0 {
		v1238 = v1255
		v1239 = v1253
		goto L305
	} else {
		goto L310
	}
L310:
	;
	goto L306
L311:
	;
	v1272 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L52
	} else {
		goto L312
	}
L312:
	;
	if v1272 == int32(0) {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_errcode(m, int32(16908352))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L52
	} else {
		goto L314
	}
L314:
	;
	F_errmsg(m, int32(413184), int32(0))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L52
	} else {
		goto L315
	}
L315:
	;
	F_errdetail(m, int32(632505), int32(0))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L52
	} else {
		goto L316
	}
L316:
	;
	F_errhint(m, int32(608567), int32(0))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L52
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(484591), int32(185), int32(413014))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L52
	} else {
		goto L318
	}
L318:
	;
	goto L1
L319:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L52
	} else {
		goto L320
	}
L320:
	;
	F_errmsg(m, int32(322290), int32(0))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L52
	} else {
		goto L321
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(512)
	F_errdetail(m, int32(572554), v13)
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L52
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(484591), int32(176), int32(413014))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L52
	} else {
		goto L323
	}
L323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L324:
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
	v22 = int32(8192)
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
	F_UnregisterExprContextCallback(m, v4, int32(1628), v6)
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v629 int32
	_ = v629
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v974 int32
	_ = v974
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1198 int32
	_ = v1198
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1432 int32
	_ = v1432
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
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1539 int32
	_ = v1539
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	v9 = v6 + int32(2)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 <= v9 {
		v105 = v10
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v1539
L2:
	;
	v1539 = int32(1)
	goto L1
L3:
	;
	v109 = v6 + int32(3)
	v110 = base.B2i32(v105 < v109)
	if v105 < v109 {
		goto L44
	} else {
		goto L45
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v9))))
	if v14&int32(224) != int32(96) {
		v105 = v10
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if int32(1)<<(uint(v14)%32)&int32(42750482) == int32(0) {
		v105 = v10
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v27 = F_find_among(m, l0, int32(4165760), int32(18))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v27 == int32(0) {
		v105 = v31
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v34
	if v34 < v31 {
		v105 = v31
		goto L3
	} else {
		goto L10
	}
L10:
	;
	switch v27 - int32(1) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	case 5:
		goto L16
	case 6:
		goto L15
	case 7:
		goto L14
	case 8:
		goto L13
	case 9:
		goto L12
	case 10:
		goto L11
	default:
		goto L2
	}
L11:
	;
	v101 = F_slice_from_s(m, l0, int32(5), int32(2148829))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L42
	}
L12:
	;
	v95 = F_slice_from_s(m, l0, int32(4), int32(2148825))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L40
	}
L13:
	;
	v89 = F_slice_from_s(m, l0, int32(5), int32(2148820))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L38
	}
L14:
	;
	v83 = F_slice_from_s(m, l0, int32(4), int32(2148816))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L7
	} else {
		goto L36
	}
L15:
	;
	v77 = F_slice_from_s(m, l0, int32(5), int32(2148811))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L34
	}
L16:
	;
	v71 = F_slice_from_s(m, l0, int32(3), int32(2148808))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L32
	}
L17:
	;
	v65 = F_slice_from_s(m, l0, int32(3), int32(2148805))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L30
	}
L18:
	;
	v59 = F_slice_from_s(m, l0, int32(3), int32(2148802))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L28
	}
L19:
	;
	v53 = F_slice_from_s(m, l0, int32(3), int32(2148799))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L26
	}
L20:
	;
	v47 = F_slice_from_s(m, l0, int32(3), int32(2148796))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L24
	}
L21:
	;
	v41 = F_slice_from_s(m, l0, int32(3), int32(2148793))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	if int32(0) <= v41 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v1539 = v41
	goto L1
L24:
	;
	if int32(0) <= v47 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v1539 = v47
	goto L1
L26:
	;
	if int32(0) <= v53 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v1539 = v53
	goto L1
L28:
	;
	if int32(0) <= v59 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v1539 = v59
	goto L1
L30:
	;
	if int32(0) <= v65 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v1539 = v65
	goto L1
L32:
	;
	if int32(0) <= v71 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v1539 = v71
	goto L1
L34:
	;
	if int32(0) <= v77 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v1539 = v77
	goto L1
L36:
	;
	if int32(0) <= v83 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v1539 = v83
	goto L1
L38:
	;
	if int32(0) <= v89 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v1539 = v89
	goto L1
L40:
	;
	if int32(0) <= v95 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v1539 = v95
	goto L1
L42:
	;
	if int32(0) <= v101 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v1539 = v101
	goto L1
L44:
	;
	v111 = v6
	goto L46
L45:
	;
	v111 = v109
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	if v105 < v109 {
		v1539 = int32(1)
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v118 == v120 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118
	v163 = v118
	goto L59
L49:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v118))))
	if v124 == int32(39) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v128 = v118 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v128
	v131 = F_slice_del(m, l0)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	v136 = v120
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118
	if v136 == v118 {
		goto L48
	} else {
		goto L55
	}
L53:
	;
	if v131 < int32(0) {
		v1539 = v131
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v136 = v135
	goto L52
L55:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v118))))
	if v143 != int32(121) {
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v146 = int32(1)
	v147 = v118 + v146
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v147
	v152 = F_slice_from_s(m, l0, v146, int32(2148918))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	if v152 < int32(0) {
		v1539 = v152
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+8)) = int32(1)
	goto L48
L59:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v176 < v175 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v249)+4)) = v220
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v254 = v252 + int32(4)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v255 <= v254 {
		goto L86
	} else {
		goto L87
	}
L61:
	;
	goto L60
L62:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v219 != 0 {
		goto L77
	} else {
		goto L78
	}
L63:
	;
	v178 = v175
	goto L65
L64:
	;
	v178 = v176
	goto L65
L65:
	;
	goto L67
L66:
	;
	v219 = v215
	goto L62
L67:
	;
	if v175 == v178 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v215 = int32(0)
	goto L66
L69:
	;
	v219 = int32(-1)
	goto L62
L70:
	;
	goto L71
L71:
	;
	v190 = int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191+v175))))
	if int32(121) < v193 {
		v215 = v190
		goto L66
	} else {
		goto L72
	}
L72:
	;
	v195 = v193 - int32(97)
	if v195 < int32(0) {
		v215 = v190
		goto L66
	} else {
		goto L73
	}
L73:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v195)>>(uint(int32(3))%32)))+uint32(_consts[1433]))))
	if int32(base.Ui32(v201)>>(uint(v195&int32(7))%32))&int32(1) == int32(0) {
		v215 = v190
		goto L66
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v175 + int32(1)
	goto L75
L75:
	;
	goto L68
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v163
	v235 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v221 + v235
	v240 = F_slice_from_s(m, l0, v235, int32(2148923))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L7
	} else {
		goto L82
	}
L77:
	;
	if v220 <= v163 {
		goto L61
	} else {
		goto L81
	}
L78:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v221
	if v220 == v221 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v221))))
	if v226 == int32(121) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	v232 = v163 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v232
	v163 = v232
	goto L59
L82:
	;
	if v240 < int32(0) {
		v1539 = v240
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v163 = v247
	goto L59
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v252
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v509
	if v509 <= v252 {
		v538 = v509
		goto L154
	} else {
		goto L155
	}
L85:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+4)) = v391
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v402 < v401 {
		goto L124
	} else {
		goto L125
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v287 < v252 {
		goto L93
	} else {
		goto L94
	}
L87:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257+v254))))
	if v259&int32(224) != int32(96) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	if int32(1)<<(uint(v259)%32)&int32(2375680) == int32(0) {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v272 = F_find_among(m, l0, int32(4166128), int32(3))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	if v272 == int32(0) {
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v391 = v276
	goto L85
L92:
	;
	if v327 < int32(0) {
		goto L84
	} else {
		goto L107
	}
L93:
	;
	v289 = v252
	goto L95
L94:
	;
	v289 = v287
	goto L95
L95:
	;
	v296 = v252
	goto L97
L96:
	;
	v327 = v307
	goto L92
L97:
	;
	if v296 == v289 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v327 = int32(-1)
	goto L92
L100:
	;
	goto L101
L101:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300+v296))))
	if int32(121) < v302 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v319 = v296 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v319
	v296 = v319
	goto L97
L103:
	;
	v304 = v302 - int32(97)
	if v304 < int32(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v307 = int32(1)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v304)>>(uint(int32(3))%32)))+uint32(_consts[1433]))))
	if int32(base.Ui32(v311)>>(uint(v304&int32(7))%32))&v307 != 0 {
		goto L96
	} else {
		goto L105
	}
L105:
	;
	goto L102
L107:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v331 = v330 + v327
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v342 < v331 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	if v385 < int32(0) {
		goto L84
	} else {
		goto L122
	}
L109:
	;
	v344 = v331
	goto L111
L110:
	;
	v344 = v342
	goto L111
L111:
	;
	v351 = v331
	goto L113
L112:
	;
	v385 = int32(1)
	goto L108
L113:
	;
	if v351 == v344 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v385 = int32(-1)
	goto L108
L116:
	;
	goto L117
L117:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357+v351))))
	if int32(121) < v359 {
		goto L112
	} else {
		goto L118
	}
L118:
	;
	v361 = v359 - int32(97)
	if v361 < int32(0) {
		goto L112
	} else {
		goto L119
	}
L119:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v361)>>(uint(int32(3))%32)))+uint32(_consts[1433]))))
	if int32(base.Ui32(v367)>>(uint(v361&int32(7))%32))&int32(1) == int32(0) {
		goto L112
	} else {
		goto L120
	}
L120:
	;
	v376 = v351 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v376
	v351 = v376
	goto L113
L122:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v389 = v388 + v385
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v389
	v391 = v389
	goto L85
L123:
	;
	if v442 < int32(0) {
		goto L84
	} else {
		goto L138
	}
L124:
	;
	v404 = v401
	goto L126
L125:
	;
	v404 = v402
	goto L126
L126:
	;
	v411 = v401
	goto L128
L127:
	;
	v442 = v422
	goto L123
L128:
	;
	if v411 == v404 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v442 = int32(-1)
	goto L123
L131:
	;
	goto L132
L132:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v411))))
	if int32(121) < v417 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v434 = v411 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v434
	v411 = v434
	goto L128
L134:
	;
	v419 = v417 - int32(97)
	if v419 < int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v422 = int32(1)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v419)>>(uint(int32(3))%32)))+uint32(_consts[1433]))))
	if int32(base.Ui32(v426)>>(uint(v419&int32(7))%32))&v422 != 0 {
		goto L127
	} else {
		goto L136
	}
L136:
	;
	goto L133
L138:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v446 = v445 + v442
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v446
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v457 < v446 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v500 < int32(0) {
		goto L84
	} else {
		goto L153
	}
L140:
	;
	v459 = v446
	goto L142
L141:
	;
	v459 = v457
	goto L142
L142:
	;
	v466 = v446
	goto L144
L143:
	;
	v500 = int32(1)
	goto L139
L144:
	;
	if v466 == v459 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v500 = int32(-1)
	goto L139
L147:
	;
	goto L148
L148:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v466))))
	if int32(121) < v474 {
		goto L143
	} else {
		goto L149
	}
L149:
	;
	v476 = v474 - int32(97)
	if v476 < int32(0) {
		goto L143
	} else {
		goto L150
	}
L150:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v476)>>(uint(int32(3))%32)))+uint32(_consts[1433]))))
	if int32(base.Ui32(v482)>>(uint(v476&int32(7))%32))&int32(1) == int32(0) {
		goto L143
	} else {
		goto L151
	}
L151:
	;
	v491 = v466 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v491
	v466 = v491
	goto L144
L153:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v503))) = v504 + v500
	goto L84
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v538
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v538 <= v541 {
		goto L163
	} else {
		goto L164
	}
L155:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513+v509-int32(1)))))
	if base.B2i32(v517 != int32(115))&base.B2i32(v517 != int32(39)) != 0 {
		v538 = v509
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v525 = F_find_among_b(m, l0, int32(4166192), int32(3))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L7
	} else {
		goto L157
	}
L157:
	;
	if v525 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v529
	v538 = v529
	goto L154
L159:
	;
	goto L160
L160:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v531
	v533 = F_slice_del(m, l0)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L7
	} else {
		goto L161
	}
L161:
	;
	if v533 < int32(0) {
		v1539 = v533
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v538 = v537
	goto L154
L163:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v649
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v649
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v649-int32(5) <= v652 {
		v674 = v652
		goto L197
	} else {
		goto L198
	}
L164:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v538-int32(1)))))
	switch v547 - int32(100) {
	case 0, 15:
		goto L165
	default:
		goto L163
	}
L165:
	;
	v552 = F_find_among_b(m, l0, int32(4166256), int32(6))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L7
	} else {
		goto L166
	}
L166:
	;
	if v552 == int32(0) {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v556
	switch v552 - int32(1) {
	case 0:
		goto L170
	case 1:
		goto L169
	case 2:
		goto L168
	default:
		goto L163
	}
L168:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v556 <= v584 {
		goto L163
	} else {
		goto L180
	}
L169:
	;
	v567 = v556 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v567
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v569 <= v567 {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	v562 = F_slice_from_s(m, l0, int32(2), int32(2148940))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L7
	} else {
		goto L171
	}
L171:
	;
	if int32(0) <= v562 {
		goto L163
	} else {
		goto L172
	}
L172:
	;
	v1539 = v562
	goto L1
L173:
	;
	v573 = F_slice_from_s(m, l0, int32(1), int32(2148942))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L7
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v556
	v580 = F_slice_from_s(m, l0, int32(2), int32(2148943))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L7
	} else {
		goto L178
	}
L176:
	;
	if int32(0) <= v573 {
		goto L163
	} else {
		goto L177
	}
L177:
	;
	v1539 = v573
	goto L1
L178:
	;
	if int32(0) <= v580 {
		goto L163
	} else {
		goto L179
	}
L179:
	;
	v1539 = v580
	goto L1
L180:
	;
	v587 = v556 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v587
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v603 = v587
	goto L183
L181:
	;
	if v637 < int32(0) {
		goto L163
	} else {
		goto L193
	}
L182:
	;
	v637 = v617
	goto L181
L183:
	;
	if v603 <= v597 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v637 = int32(-1)
	goto L181
L186:
	;
	goto L187
L187:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608+v603-int32(1)))))
	if int32(121) < v612 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v629 = v603 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v629
	v603 = v629
	goto L183
L189:
	;
	v614 = v612 - int32(97)
	if v614 < int32(0) {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v617 = int32(1)
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v614)>>(uint(int32(3))%32)))+uint32(_consts[1433]))))
	if int32(base.Ui32(v621)>>(uint(v614&int32(7))%32))&v617 != 0 {
		goto L182
	} else {
		goto L191
	}
L191:
	;
	goto L188
L193:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v640 - v637
	v643 = F_slice_del(m, l0)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L7
	} else {
		goto L194
	}
L194:
	;
	if v643 < int32(0) {
		v1539 = v643
		goto L1
	} else {
		goto L195
	}
L195:
	;
	goto L163
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1483
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+8))
	if v1487 == int32(0) {
		goto L431
	} else {
		goto L432
	}
L197:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v675
	v679 = v675 - int32(1)
	if v679 <= v674 {
		goto L203
	} else {
		goto L204
	}
L198:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656+v649-int32(1)))))
	switch v660 - int32(100) {
	case 0, 3:
		goto L199
	default:
		v674 = v652
		goto L197
	}
L199:
	;
	v665 = F_find_among_b(m, l0, int32(4166384), int32(8))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L7
	} else {
		goto L200
	}
L200:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v665 == int32(0) {
		v674 = v667
		goto L197
	} else {
		goto L201
	}
L201:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v670
	if v670 <= v667 {
		v1483 = v667
		goto L196
	} else {
		goto L202
	}
L202:
	;
	v674 = v667
	goto L197
L203:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v946
	v948 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v946
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v946 <= v951 {
		v1030 = v948
		goto L267
	} else {
		goto L268
	}
L204:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681+v679))))
	if v683&int32(224) != int32(96) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	if int32(1)<<(uint(v683)%32)&int32(33554576) == int32(0) {
		goto L203
	} else {
		goto L206
	}
L206:
	;
	v696 = F_find_among_b(m, l0, int32(4166544), int32(6))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L7
	} else {
		goto L207
	}
L207:
	;
	if v696 == int32(0) {
		goto L203
	} else {
		goto L208
	}
L208:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v700
	switch v696 - int32(1) {
	case 0:
		goto L210
	case 1:
		goto L209
	default:
		goto L203
	}
L209:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v728 = v721
	goto L216
L210:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)+4))
	if v700 < v705 {
		goto L203
	} else {
		goto L211
	}
L211:
	;
	v709 = F_slice_from_s(m, l0, int32(2), int32(2149019))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L7
	} else {
		goto L212
	}
L212:
	;
	if int32(0) <= v709 {
		goto L203
	} else {
		goto L213
	}
L213:
	;
	v1539 = v709
	goto L1
L214:
	;
	if v762 < int32(0) {
		goto L203
	} else {
		goto L226
	}
L215:
	;
	v762 = v742
	goto L214
L216:
	;
	if v728 <= v722 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v762 = int32(-1)
	goto L214
L219:
	;
	goto L220
L220:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733+v728-int32(1)))))
	if int32(121) < v737 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v754 = v728 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v754
	v728 = v754
	goto L216
L222:
	;
	v739 = v737 - int32(97)
	if v739 < int32(0) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v742 = int32(1)
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v739)>>(uint(int32(3))%32)))+uint32(_consts[1433]))))
	if int32(base.Ui32(v746)>>(uint(v739&int32(7))%32))&v742 != 0 {
		goto L215
	} else {
		goto L224
	}
L224:
	;
	goto L221
L226:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v765 + (v700 - v713)
	v769 = F_slice_del(m, l0)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L7
	} else {
		goto L227
	}
L227:
	;
	if v769 < int32(0) {
		v1539 = v769
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v773
	v777 = v773 - int32(1)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v777 <= v778 {
		v858 = v773
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v930 = v928 + (v773 - v793)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v930
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v930
	if v930 <= v926 {
		goto L203
	} else {
		goto L264
	}
L230:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v926 = v925
	goto L229
L231:
	;
	v917 = F_slice_from_s(m, l0, int32(1), v914)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L7
	} else {
		goto L262
	}
L232:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	if v858 != v862 {
		goto L203
	} else {
		goto L252
	}
L233:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780+v777))))
	if v782&int32(224) != int32(96) {
		v858 = v773
		goto L232
	} else {
		goto L234
	}
L234:
	;
	if int32(1)<<(uint(v782)%32)&int32(68514004) == int32(0) {
		v858 = v773
		goto L232
	} else {
		goto L235
	}
L235:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v797 = F_find_among_b(m, l0, int32(4166672), int32(13))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L7
	} else {
		goto L238
	}
L236:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v858 = v857
	goto L232
L237:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L241
L238:
	;
	switch v797 - int32(1) {
	case 0:
		v914 = int32(2149021)
		goto L231
	case 1:
		goto L237
	case 2:
		goto L236
	default:
		goto L230
	}
L239:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v853 != 0 {
		v926 = v854
		goto L229
	} else {
		goto L250
	}
L240:
	;
	v853 = v849
	goto L239
L241:
	;
	if v809 <= v810 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v849 = int32(0)
	goto L240
L243:
	;
	v853 = int32(-1)
	goto L239
L244:
	;
	goto L245
L245:
	;
	v822 = int32(1)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823+v809-v822))))
	if int32(111) < v827 {
		v849 = v822
		goto L240
	} else {
		goto L246
	}
L246:
	;
	v829 = v827 - int32(97)
	if v829 < int32(0) {
		v849 = v822
		goto L240
	} else {
		goto L247
	}
L247:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v829)>>(uint(int32(3))%32)))+uint32(_consts[1434]))))
	if int32(base.Ui32(v835)>>(uint(v829&int32(7))%32))&int32(1) == int32(0) {
		v849 = v822
		goto L240
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v809 - int32(1)
	goto L249
L249:
	;
	goto L242
L250:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v854 < v855 {
		v926 = v854
		goto L229
	} else {
		goto L251
	}
L251:
	;
	goto L203
L252:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v865 = int32(0)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v874 = F_out_grouping_b(m, l0, int32(2149071), int32(89), int32(121), v865)
	mBase = m.M
	if v874 != 0 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	if v905 == int32(0) {
		goto L203
	} else {
		goto L261
	}
L254:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v886 + (v869 - v864)
	v894 = F_out_grouping_b(m, l0, int32(2148919), int32(97), int32(121), int32(0))
	mBase = m.M
	if v894 != 0 {
		v903 = v865
		goto L258
	} else {
		goto L259
	}
L255:
	;
	v879 = F_in_grouping_b(m, l0, int32(2148919), int32(97), int32(121), int32(0))
	mBase = m.M
	if v879 != 0 {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v884 = F_out_grouping_b(m, l0, int32(2148919), int32(97), int32(121), int32(0))
	mBase = m.M
	if v884 != 0 {
		goto L254
	} else {
		goto L257
	}
L257:
	;
	v905 = int32(1)
	goto L253
L258:
	;
	v905 = v903
	goto L253
L259:
	;
	v899 = F_in_grouping_b(m, l0, int32(2148919), int32(97), int32(121), int32(0))
	mBase = m.M
	if v899 != 0 {
		v903 = v865
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v903 = base.B2i32(v900 <= v901)
	goto L258
L261:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v908 + (v858 - v864)
	v914 = int32(2149024)
	goto L231
L262:
	;
	if int32(0) <= v917 {
		goto L203
	} else {
		goto L263
	}
L263:
	;
	return v917 >> (uint(int32(31)) % 32) & v917
L264:
	;
	v935 = v930 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v935
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v935
	v938 = F_slice_del(m, l0)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L7
	} else {
		goto L265
	}
L265:
	;
	if v938 < int32(0) {
		v1539 = v938
		goto L1
	} else {
		goto L266
	}
L266:
	;
	goto L203
L267:
	;
	if v1030 < int32(0) {
		v1539 = v1030
		goto L1
	} else {
		goto L288
	}
L268:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v953+v946-int32(1)))))
	if v957|int32(32) != int32(121) {
		v1030 = v948
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v963 = v946 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v963
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v963
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L272
L270:
	;
	if v1014 != 0 {
		v1030 = v948
		goto L267
	} else {
		goto L282
	}
L271:
	;
	v1014 = v1011
	goto L270
L272:
	;
	if v963 <= v974 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	v1011 = int32(0)
	goto L271
L274:
	;
	v1014 = int32(-1)
	goto L270
L275:
	;
	goto L276
L276:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985+v963-int32(1)))))
	if int32(121) < v989 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v963 - int32(1)
	goto L281
L278:
	;
	v991 = v989 - int32(97)
	if v991 < int32(0) {
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v994 = int32(1)
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v991)>>(uint(int32(3))%32)))+uint32(_consts[1433]))))
	if int32(base.Ui32(v998)>>(uint(v991&int32(7))%32))&v994 != 0 {
		v1011 = v994
		goto L271
	} else {
		goto L280
	}
L280:
	;
	goto L277
L281:
	;
	goto L273
L282:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1015 <= v1016 {
		v1030 = v948
		goto L267
	} else {
		goto L283
	}
L283:
	;
	v1018 = int32(1)
	v1021 = F_slice_from_s(m, l0, v1018, int32(2149076))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L7
	} else {
		goto L284
	}
L284:
	;
	if int32(0) <= v1021 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1028 = v1018
	goto L287
L286:
	;
	v1028 = v1021 >> (uint(int32(31)) % 32) & v1021
	goto L287
L287:
	;
	v1030 = v1028
	goto L267
L288:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1033
	v1035 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1033
	v1039 = v1033 - int32(1)
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1039 <= v1040 {
		v1223 = v1035
		goto L289
	} else {
		goto L290
	}
L289:
	;
	if v1223 < int32(0) {
		v1539 = v1223
		goto L1
	} else {
		goto L356
	}
L290:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042+v1039))))
	if v1044&int32(224) != int32(96) {
		v1223 = v1035
		goto L289
	} else {
		goto L291
	}
L291:
	;
	if int32(1)<<(uint(v1044)%32)&int32(815616) == int32(0) {
		v1223 = v1035
		goto L289
	} else {
		goto L292
	}
L292:
	;
	v1057 = F_find_among_b(m, l0, int32(4166944), int32(24))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L7
	} else {
		goto L293
	}
L293:
	;
	if v1057 == int32(0) {
		v1223 = v1035
		goto L289
	} else {
		goto L294
	}
L294:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1061
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	if v1061 < v1064 {
		v1223 = v1035
		goto L289
	} else {
		goto L295
	}
L295:
	;
	switch v1057 - int32(1) {
	case 0:
		goto L311
	case 1:
		goto L310
	case 2:
		goto L309
	case 3:
		goto L308
	case 4:
		goto L307
	case 5:
		goto L306
	case 6:
		goto L305
	case 7:
		goto L304
	case 8:
		goto L303
	case 9:
		goto L302
	case 10:
		goto L301
	case 11:
		goto L300
	case 12:
		goto L299
	case 13:
		goto L298
	case 14:
		goto L297
	default:
		goto L296
	}
L296:
	;
	v1223 = int32(1)
	goto L289
L297:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L344
L298:
	;
	v1160 = F_slice_from_s(m, l0, int32(4), int32(2149118))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L7
	} else {
		goto L340
	}
L299:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1061 <= v1140 {
		v1223 = v1035
		goto L289
	} else {
		goto L336
	}
L300:
	;
	v1136 = F_slice_from_s(m, l0, int32(3), int32(2149113))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L7
	} else {
		goto L334
	}
L301:
	;
	v1130 = F_slice_from_s(m, l0, int32(3), int32(2149110))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L7
	} else {
		goto L332
	}
L302:
	;
	v1124 = F_slice_from_s(m, l0, int32(3), int32(2149107))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L7
	} else {
		goto L330
	}
L303:
	;
	v1118 = F_slice_from_s(m, l0, int32(3), int32(2149104))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L7
	} else {
		goto L328
	}
L304:
	;
	v1112 = F_slice_from_s(m, l0, int32(2), int32(2149102))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L7
	} else {
		goto L326
	}
L305:
	;
	v1106 = F_slice_from_s(m, l0, int32(3), int32(2149099))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L7
	} else {
		goto L324
	}
L306:
	;
	v1100 = F_slice_from_s(m, l0, int32(3), int32(2149096))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L7
	} else {
		goto L322
	}
L307:
	;
	v1094 = F_slice_from_s(m, l0, int32(3), int32(2149093))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L7
	} else {
		goto L320
	}
L308:
	;
	v1088 = F_slice_from_s(m, l0, int32(4), int32(2149089))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L7
	} else {
		goto L318
	}
L309:
	;
	v1082 = F_slice_from_s(m, l0, int32(4), int32(2149085))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L7
	} else {
		goto L316
	}
L310:
	;
	v1076 = F_slice_from_s(m, l0, int32(4), int32(2149081))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L7
	} else {
		goto L314
	}
L311:
	;
	v1070 = F_slice_from_s(m, l0, int32(4), int32(2149077))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L7
	} else {
		goto L312
	}
L312:
	;
	if int32(0) <= v1070 {
		goto L296
	} else {
		goto L313
	}
L313:
	;
	v1223 = v1070
	goto L289
L314:
	;
	if int32(0) <= v1076 {
		goto L296
	} else {
		goto L315
	}
L315:
	;
	v1223 = v1076
	goto L289
L316:
	;
	if int32(0) <= v1082 {
		goto L296
	} else {
		goto L317
	}
L317:
	;
	v1223 = v1082
	goto L289
L318:
	;
	if int32(0) <= v1088 {
		goto L296
	} else {
		goto L319
	}
L319:
	;
	v1223 = v1088
	goto L289
L320:
	;
	if int32(0) <= v1094 {
		goto L296
	} else {
		goto L321
	}
L321:
	;
	v1223 = v1094
	goto L289
L322:
	;
	if int32(0) <= v1100 {
		goto L296
	} else {
		goto L323
	}
L323:
	;
	v1223 = v1100
	goto L289
L324:
	;
	if int32(0) <= v1106 {
		goto L296
	} else {
		goto L325
	}
L325:
	;
	v1223 = v1106
	goto L289
L326:
	;
	if int32(0) <= v1112 {
		goto L296
	} else {
		goto L327
	}
L327:
	;
	v1223 = v1112
	goto L289
L328:
	;
	if int32(0) <= v1118 {
		goto L296
	} else {
		goto L329
	}
L329:
	;
	v1223 = v1118
	goto L289
L330:
	;
	if int32(0) <= v1124 {
		goto L296
	} else {
		goto L331
	}
L331:
	;
	v1223 = v1124
	goto L289
L332:
	;
	if int32(0) <= v1130 {
		goto L296
	} else {
		goto L333
	}
L333:
	;
	v1223 = v1130
	goto L289
L334:
	;
	if int32(0) <= v1136 {
		goto L296
	} else {
		goto L335
	}
L335:
	;
	v1223 = v1136
	goto L289
L336:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1142+v1061-int32(1)))))
	if v1146 != int32(108) {
		v1223 = v1035
		goto L289
	} else {
		goto L337
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1061 - int32(1)
	v1154 = F_slice_from_s(m, l0, int32(2), int32(2149116))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L7
	} else {
		goto L338
	}
L338:
	;
	if int32(0) <= v1154 {
		goto L296
	} else {
		goto L339
	}
L339:
	;
	v1223 = v1154
	goto L289
L340:
	;
	if int32(0) <= v1160 {
		goto L296
	} else {
		goto L341
	}
L341:
	;
	v1223 = v1160
	goto L289
L342:
	;
	if v1216 != 0 {
		v1223 = v1035
		goto L289
	} else {
		goto L353
	}
L343:
	;
	v1216 = v1212
	goto L342
L344:
	;
	if v1172 <= v1173 {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	v1212 = int32(0)
	goto L343
L346:
	;
	v1216 = int32(-1)
	goto L342
L347:
	;
	goto L348
L348:
	;
	v1185 = int32(1)
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186+v1172-v1185))))
	if int32(116) < v1190 {
		v1212 = v1185
		goto L343
	} else {
		goto L349
	}
L349:
	;
	v1192 = v1190 - int32(99)
	if v1192 < int32(0) {
		v1212 = v1185
		goto L343
	} else {
		goto L350
	}
L350:
	;
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1192)>>(uint(int32(3))%32)))+uint32(_consts[1435]))))
	if int32(base.Ui32(v1198)>>(uint(v1192&int32(7))%32))&int32(1) == int32(0) {
		v1212 = v1185
		goto L343
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1172 - int32(1)
	goto L352
L352:
	;
	goto L345
L353:
	;
	v1217 = F_slice_del(m, l0)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L7
	} else {
		goto L354
	}
L354:
	;
	if v1217 < int32(0) {
		v1223 = v1217
		goto L289
	} else {
		goto L355
	}
L355:
	;
	goto L296
L356:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1228
	v1230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1228
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1228-int32(2) <= v1233 {
		v1301 = v1230
		goto L357
	} else {
		goto L358
	}
L357:
	;
	if v1301 < int32(0) {
		v1539 = v1301
		goto L1
	} else {
		goto L384
	}
L358:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1237+v1228-int32(1)))))
	if v1241&int32(224) != int32(96) {
		v1301 = v1230
		goto L357
	} else {
		goto L359
	}
L359:
	;
	if int32(1)<<(uint(v1241)%32)&int32(528928) == int32(0) {
		v1301 = v1230
		goto L357
	} else {
		goto L360
	}
L360:
	;
	v1254 = F_find_among_b(m, l0, int32(4167424), int32(9))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L7
	} else {
		goto L361
	}
L361:
	;
	if v1254 == int32(0) {
		v1301 = v1230
		goto L357
	} else {
		goto L362
	}
L362:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1258
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+4))
	if v1258 < v1261 {
		v1301 = v1230
		goto L357
	} else {
		goto L363
	}
L363:
	;
	switch v1254 - int32(1) {
	case 0:
		goto L370
	case 1:
		goto L369
	case 2:
		goto L368
	case 3:
		goto L367
	case 4:
		goto L366
	case 5:
		goto L365
	default:
		goto L364
	}
L364:
	;
	v1301 = int32(1)
	goto L357
L365:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1260)))
	if v1258 < v1293 {
		v1301 = v1230
		goto L357
	} else {
		goto L381
	}
L366:
	;
	v1289 = F_slice_del(m, l0)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L7
	} else {
		goto L379
	}
L367:
	;
	v1285 = F_slice_from_s(m, l0, int32(2), int32(2149254))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L7
	} else {
		goto L377
	}
L368:
	;
	v1279 = F_slice_from_s(m, l0, int32(2), int32(2149252))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L7
	} else {
		goto L375
	}
L369:
	;
	v1273 = F_slice_from_s(m, l0, int32(3), int32(2149249))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L7
	} else {
		goto L373
	}
L370:
	;
	v1267 = F_slice_from_s(m, l0, int32(4), int32(2149245))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L7
	} else {
		goto L371
	}
L371:
	;
	if int32(0) <= v1267 {
		goto L364
	} else {
		goto L372
	}
L372:
	;
	v1301 = v1267
	goto L357
L373:
	;
	if int32(0) <= v1273 {
		goto L364
	} else {
		goto L374
	}
L374:
	;
	v1301 = v1273
	goto L357
L375:
	;
	if int32(0) <= v1279 {
		goto L364
	} else {
		goto L376
	}
L376:
	;
	v1301 = v1279
	goto L357
L377:
	;
	if int32(0) <= v1285 {
		goto L364
	} else {
		goto L378
	}
L378:
	;
	v1301 = v1285
	goto L357
L379:
	;
	if int32(0) <= v1289 {
		goto L364
	} else {
		goto L380
	}
L380:
	;
	v1301 = v1289
	goto L357
L381:
	;
	v1295 = F_slice_del(m, l0)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L7
	} else {
		goto L382
	}
L382:
	;
	if v1295 < int32(0) {
		v1301 = v1295
		goto L357
	} else {
		goto L383
	}
L383:
	;
	goto L364
L384:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1307
	v1309 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1307
	v1313 = v1307 - int32(1)
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1313 <= v1314 {
		v1370 = v1309
		goto L385
	} else {
		goto L386
	}
L385:
	;
	if v1370 < int32(0) {
		v1539 = v1370
		goto L1
	} else {
		goto L401
	}
L386:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1316+v1313))))
	if v1318&int32(224) != int32(96) {
		v1370 = v1309
		goto L385
	} else {
		goto L387
	}
L387:
	;
	if int32(1)<<(uint(v1318)%32)&int32(1864232) == int32(0) {
		v1370 = v1309
		goto L385
	} else {
		goto L388
	}
L388:
	;
	v1331 = F_find_among_b(m, l0, int32(4167616), int32(18))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L7
	} else {
		goto L389
	}
L389:
	;
	if v1331 == int32(0) {
		v1370 = v1309
		goto L385
	} else {
		goto L390
	}
L390:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1335
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1337)))
	if v1335 < v1338 {
		v1370 = v1309
		goto L385
	} else {
		goto L391
	}
L391:
	;
	switch v1331 - int32(1) {
	case 0:
		goto L394
	case 1:
		goto L393
	default:
		goto L392
	}
L392:
	;
	v1370 = int32(1)
	goto L385
L393:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1335 <= v1346 {
		v1370 = v1309
		goto L385
	} else {
		goto L397
	}
L394:
	;
	v1342 = F_slice_del(m, l0)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L7
	} else {
		goto L395
	}
L395:
	;
	if int32(0) <= v1342 {
		goto L392
	} else {
		goto L396
	}
L396:
	;
	v1370 = v1342
	goto L385
L397:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1350 = int32(1)
	v1352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1348+v1335-v1350))))
	if base.Ui32(v1350) < base.Ui32((v1352-int32(115))&int32(255)) {
		v1370 = v1309
		goto L385
	} else {
		goto L398
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1335 - int32(1)
	v1362 = F_slice_del(m, l0)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L7
	} else {
		goto L399
	}
L399:
	;
	if v1362 < int32(0) {
		v1370 = v1362
		goto L385
	} else {
		goto L400
	}
L400:
	;
	goto L392
L401:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1373
	v1375 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1373
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1373 <= v1378 {
		v1476 = v1375
		goto L402
	} else {
		goto L403
	}
L402:
	;
	if v1476 < int32(0) {
		v1539 = v1476
		goto L1
	} else {
		goto L430
	}
L403:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1380+v1373-int32(1)))))
	switch v1384 - int32(101) {
	case 0, 7:
		goto L404
	default:
		v1476 = v1375
		goto L402
	}
L404:
	;
	v1389 = F_find_among_b(m, l0, int32(4167984), int32(2))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L7
	} else {
		goto L405
	}
L405:
	;
	if v1389 == int32(0) {
		v1476 = v1375
		goto L402
	} else {
		goto L406
	}
L406:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1393
	switch v1389 - int32(1) {
	case 0:
		goto L409
	case 1:
		goto L408
	default:
		goto L407
	}
L407:
	;
	v1476 = int32(1)
	goto L402
L408:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1453)))
	if v1393 < v1454 {
		v1476 = v1375
		goto L402
	} else {
		goto L425
	}
L409:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1397)))
	if v1393 < v1398 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+4))
	if v1393 < v1400 {
		v1476 = v1375
		goto L402
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v1449 = F_slice_del(m, l0)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L7
	} else {
		goto L423
	}
L413:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1403 = int32(0)
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1412 = F_out_grouping_b(m, l0, int32(2149071), int32(89), int32(121), v1403)
	mBase = m.M
	if v1412 != 0 {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	if v1443 != 0 {
		v1476 = v1375
		goto L402
	} else {
		goto L422
	}
L415:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1424 + (v1407 - v1402)
	v1432 = F_out_grouping_b(m, l0, int32(2148919), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1432 != 0 {
		v1441 = v1403
		goto L419
	} else {
		goto L420
	}
L416:
	;
	v1417 = F_in_grouping_b(m, l0, int32(2148919), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1417 != 0 {
		goto L415
	} else {
		goto L417
	}
L417:
	;
	v1422 = F_out_grouping_b(m, l0, int32(2148919), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1422 != 0 {
		goto L415
	} else {
		goto L418
	}
L418:
	;
	v1443 = int32(1)
	goto L414
L419:
	;
	v1443 = v1441
	goto L414
L420:
	;
	v1437 = F_in_grouping_b(m, l0, int32(2148919), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1437 != 0 {
		v1441 = v1403
		goto L419
	} else {
		goto L421
	}
L421:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1441 = base.B2i32(v1438 <= v1439)
	goto L419
L422:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1444 + (v1393 - v1402)
	goto L412
L423:
	;
	if int32(0) <= v1449 {
		goto L407
	} else {
		goto L424
	}
L424:
	;
	v1476 = v1449
	goto L402
L425:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1393 <= v1456 {
		v1476 = v1375
		goto L402
	} else {
		goto L426
	}
L426:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1458+v1393-int32(1)))))
	if v1462 != int32(108) {
		v1476 = v1375
		goto L402
	} else {
		goto L427
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1393 - int32(1)
	v1468 = F_slice_del(m, l0)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L7
	} else {
		goto L428
	}
L428:
	;
	if v1468 < int32(0) {
		v1476 = v1468
		goto L402
	} else {
		goto L429
	}
L429:
	;
	goto L407
L430:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1483 = v1480
	goto L196
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1483
	goto L2
L432:
	;
	goto L433
L433:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1496 < v1495 {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	v1539 = v1521
	goto L1
L435:
	;
	v1498 = v1495
	goto L437
L436:
	;
	v1498 = v1496
	goto L437
L437:
	;
	v1500 = v1495
	goto L438
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1500
	if v1500 != v1496 {
		goto L441
	} else {
		goto L442
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1500
	v1516 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1500 + v1516
	v1521 = F_slice_from_s(m, l0, v1516, int32(2149360))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L7
	} else {
		goto L446
	}
L440:
	;
	goto L439
L441:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1506+v1500))))
	if v1508 == int32(89) {
		goto L440
	} else {
		goto L444
	}
L442:
	;
	goto L443
L443:
	;
	if v1500 == v1498 {
		goto L431
	} else {
		goto L445
	}
L444:
	;
	goto L443
L445:
	;
	v1513 = v1500 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1513
	v1500 = v1513
	goto L438
L446:
	;
	if int32(0) <= v1521 {
		goto L433
	} else {
		goto L447
	}
L447:
	;
	goto L434
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1146]))
	if int32(0) <= v5 {
		v8 = int32(100663808)
		v10 = v5 * int32(100)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[1209])))
		switch v13 - int32(13) {
		case 0, 2, 10, 25, 26, 27, 51, 60:
			v19 = v8
		case 1, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 52, 53, 54, 55, 56, 57, 58, 59:
			v19 = int32(2600)
		default:
			if v13 == int32(142) {
				v19 = v8
			} else {
				v19 = int32(2600)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[1157]))) = v19
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[1146])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(446655), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errfinish(m, int32(489588), int32(959), int32(127875))
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
				F_errmsg_internal(m, int32(49727), v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errfinish(m, int32(485404), int32(414), int32(359949))
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v114 int32
	_ = v114
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
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
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
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
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v506 int32
	_ = v506
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v21 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l0 != 0 {
		goto L33
	} else {
		goto L34
	}
L2:
	;
	if l0 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v138 = v3
	v145 = v3
	goto L4
L4:
	;
	v151 = int32(1)
	v152 = int32(0)
	v156 = F_RangeVarGetRelidExtended(m, l1, v152, v151, v152, v152)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L30
	} else {
		goto L31
	}
L5:
	;
	v29 = l0
	v33 = v3
	goto L8
L6:
	;
	goto L7
L7:
	;
	v130 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v133 = F_get_visible_ENR_metadata(m, v132, v20)
	mBase = m.M
	goto L28
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v39 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v114 != 0 {
		v29 = v114
		v33 = v33 + int32(1)
		goto L8
	} else {
		goto L27
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v42 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v49 = int32(0)
	goto L13
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v45+v49<<(uint(int32(2))%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v70 == int32(0) {
		v89 = v69
		v90 = v70
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v95 = int32(1)
	v160 = v95
	v165 = v3
	v167 = v33
	v168 = v95
	goto L1
L15:
	;
	if v90-v89 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	goto L15
L17:
	;
	if v69 != v70 {
		v89 = v69
		v90 = v70
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v74 = v66
	v75 = v20
	goto L19
L19:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v79 == int32(0) {
		v89 = v78
		v90 = v79
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v89 = v78
	v90 = v79
	goto L16
L21:
	;
	v82 = int32(1)
	if v78 == v79 {
		v74 = v74 + v82
		v75 = v75 + v82
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v93 = v49 + int32(1)
	if v93 != v42 {
		v49 = v93
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L14
L26:
	;
	goto L10
L27:
	;
	goto L9
L28:
	;
	if v133 != v130 {
		v160 = v130
		v165 = v3
		v167 = v130
		v168 = v3
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v138 = v130
	v145 = v130
	goto L4
L30:
	;
	return
L31:
	;
	v160 = v138
	v165 = v156
	v167 = v145
	v168 = v151
	goto L1
L32:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L30
	} else {
		goto L130
	}
L33:
	;
	v186 = v3
	v187 = l0
	goto L36
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L30
	} else {
		goto L125
	}
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	if v190 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v506 != 0 {
		v186 = v186 + int32(1)
		v187 = v506
		goto L36
	} else {
		goto L124
	}
L39:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v193 <= int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v200 = int32(0)
	goto L41
L41:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v196+v200<<(uint(int32(2))%32))))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	if v217 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v319 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L43:
	;
	goto L42
L44:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if v292 == int32(0) {
		v311 = v291
		v312 = v292
		goto L73
	} else {
		goto L74
	}
L45:
	;
	if v160^int32(1)|base.B2i32(v217 != int32(6)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	if v165 == int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	if v220 != v165 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v216)+88))
	if v227+v186 != v167 {
		goto L44
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v168|base.B2i32(v217 != int32(7)) != 0 {
		goto L44
	} else {
		goto L62
	}
L52:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v216)+84))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v234 == int32(0) {
		v253 = v233
		v254 = v234
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v254-v253 != 0 {
		goto L44
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	if v233 != v234 {
		v253 = v233
		v254 = v234
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v238 = v230
	v239 = v20
	goto L57
L57:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	if v243 == int32(0) {
		v253 = v242
		v254 = v243
		goto L54
	} else {
		goto L59
	}
L58:
	;
	v253 = v242
	v254 = v243
	goto L54
L59:
	;
	v246 = int32(1)
	if v242 == v243 {
		v238 = v238 + v246
		v239 = v239 + v246
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	goto L43
L62:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v216)+108))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v263 == int32(0) {
		v282 = v262
		v283 = v263
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v283-v282 == int32(0) {
		goto L43
	} else {
		goto L71
	}
L64:
	;
	goto L63
L65:
	;
	if v262 != v263 {
		v282 = v262
		v283 = v263
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v267 = v259
	v268 = v20
	goto L67
L67:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+1)))
	if v272 == int32(0) {
		v282 = v271
		v283 = v272
		goto L64
	} else {
		goto L69
	}
L68:
	;
	v282 = v271
	v283 = v272
	goto L64
L69:
	;
	v275 = int32(1)
	if v271 == v272 {
		v267 = v267 + v275
		v268 = v268 + v275
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	goto L44
L72:
	;
	if v312-v311 == int32(0) {
		goto L43
	} else {
		goto L80
	}
L73:
	;
	goto L72
L74:
	;
	if v291 != v292 {
		v311 = v291
		v312 = v292
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v296 = v288
	v297 = v20
	goto L76
L76:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+1)))
	if v301 == int32(0) {
		v311 = v300
		v312 = v301
		goto L73
	} else {
		goto L78
	}
L77:
	;
	v311 = v300
	v312 = v301
	goto L73
L78:
	;
	v304 = int32(1)
	if v300 == v301 {
		v296 = v296 + v304
		v297 = v297 + v304
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v317 = v200 + int32(1)
	if v317 != v193 {
		v200 = v317
		goto L41
	} else {
		goto L81
	}
L81:
	;
	goto L38
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L30
	} else {
		goto L118
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L30
	} else {
		goto L98
	}
L84:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	if v328 == int32(0) {
		v347 = v327
		v348 = v328
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v348-v347 == int32(0) {
		goto L83
	} else {
		goto L93
	}
L86:
	;
	goto L85
L87:
	;
	if v327 != v328 {
		v347 = v327
		v348 = v328
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v332 = v323
	v333 = v324
	goto L89
L89:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+1)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+1)))
	if v337 == int32(0) {
		v347 = v336
		v348 = v337
		goto L86
	} else {
		goto L91
	}
L90:
	;
	v347 = v336
	v348 = v337
	goto L86
L91:
	;
	v340 = int32(1)
	if v336 == v337 {
		v332 = v332 + v340
		v333 = v333 + v340
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v356 = F_refnameNamespaceItem(m, l0, int32(0), v323, v353, v18+int32(76))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L30
	} else {
		goto L94
	}
L94:
	;
	if v356 == int32(0) {
		goto L83
	} else {
		goto L95
	}
L95:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	if v360 != v216 {
		goto L83
	} else {
		goto L96
	}
L96:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	if v363 != 0 {
		goto L82
	} else {
		goto L97
	}
L97:
	;
	goto L83
L98:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L30
	} else {
		goto L99
	}
L99:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v372
	F_errmsg(m, int32(691921), v18+int32(32))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L30
	} else {
		goto L100
	}
L100:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v380
	F_errdetail(m, int32(549314), v18+int32(16))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L30
	} else {
		goto L101
	}
L101:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v387 != 0 {
		goto L32
	} else {
		goto L102
	}
L102:
	;
	v393 = l0
	goto L103
L103:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v393)+28))
	if v403 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L32
L105:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	if v460 != 0 {
		v393 = v460
		goto L103
	} else {
		goto L117
	}
L106:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v406 <= int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	v413 = int32(0)
	goto L108
L108:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v409+v413<<(uint(int32(2))%32))))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	if v216 != v430 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+22)))
	if v435 != int32(1) {
		goto L32
	} else {
		goto L114
	}
L110:
	;
	v433 = v413 + int32(1)
	if v433 != v406 {
		v413 = v433
		goto L108
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	goto L109
L113:
	;
	goto L105
L114:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+23)))
	if v438 != int32(1) {
		goto L32
	} else {
		goto L115
	}
L115:
	;
	F_errhint(m, int32(633226), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L30
	} else {
		goto L116
	}
L116:
	;
	goto L32
L117:
	;
	goto L104
L118:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L30
	} else {
		goto L119
	}
L119:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v468
	F_errmsg(m, int32(691921), v18-int32(-64))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L30
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v363
	F_errhint(m, int32(639691), v18+int32(48))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L30
	} else {
		goto L121
	}
L121:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L30
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(487844), int32(3743), int32(528813))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L30
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
	goto L37
L125:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L30
	} else {
		goto L126
	}
L126:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v529
	F_errmsg(m, int32(691975), v18)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L30
	} else {
		goto L127
	}
L127:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L30
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(487844), int32(3761), int32(528813))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L30
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errfinish(m, int32(487844), int32(3754), int32(528813))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L30
	} else {
		goto L131
	}
L131:
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
	var v126 int64
	_ = v126
	var v156 int64
	_ = v156
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
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
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
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
	v210 = v48
	goto L12
L12:
	;
	v220 = v210 + int32(8)
	v224 = v210
	goto L53
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
	if v197 < v198 {
		goto L49
	} else {
		goto L50
	}
L15:
	;
	goto L14
L16:
	;
	v187 = v62 - v61
	if int32(512) <= v187 {
		goto L44
	} else {
		goto L45
	}
L17:
	;
	v77 = v71&int64(36170086419038336) ^ int64(-9187201950435737472)
	if v77&(v71-int64(2314885530818453536)) != int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v71&int64(63050394783186944) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L23
	}
L20:
	;
	if v77&(v71^int64(2459565876494606882)-int64(72340172838076673)) != int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L21
	}
L21:
	;
	if v77&(v71^int64(6655295901103053916)-int64(72340172838076673)) == int64(0) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v197 = v61
	v198 = v62
	goto L15
L23:
	;
	if v71&int64(246290604621824) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L24
	}
L24:
	;
	if v71&int64(962072674304) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L25
	}
L25:
	;
	if v71&int64(3758096384) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L26
	}
L26:
	;
	if v71&int64(14680064) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L27
	}
L27:
	;
	if v71&int64(224) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L28
	}
L28:
	;
	if v71&int64(57344) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v126 = v71 ^ int64(2459565876494606882)
	if v126&int64(71776119061217280) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L30
	}
L30:
	;
	if v126&int64(280375465082880) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L31
	}
L31:
	;
	if v126&int64(1095216660480) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L32
	}
L32:
	;
	if v126&int64(4278190080) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L33
	}
L33:
	;
	if v126&int64(16711680) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L34
	}
L34:
	;
	if v126&int64(255) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L35
	}
L35:
	;
	if v126&int64(65280) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L36
	}
L36:
	;
	v156 = v71 ^ int64(6655295901103053916)
	if v156&int64(71776119061217280) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L37
	}
L37:
	;
	if v156&int64(280375465082880) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L38
	}
L38:
	;
	if v156&int64(1095216660480) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L39
	}
L39:
	;
	if v156&int64(4278190080) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L40
	}
L40:
	;
	if v156&int64(16711680) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L41
	}
L41:
	;
	if v156&int64(255) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L42
	}
L42:
	;
	if v156&int64(65280) == int64(0) {
		v197 = v61
		v198 = v62
		goto L15
	} else {
		goto L43
	}
L43:
	;
	goto L16
L44:
	;
	F_appendBinaryStringInfo(m, l0, l1+v61, v187)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	v193 = v61
	goto L46
L46:
	;
	v195 = v62 + int32(8)
	if v195 < v43 {
		v61 = v193
		v62 = v195
		goto L13
	} else {
		goto L48
	}
L47:
	;
	v193 = v62
	goto L46
L48:
	;
	v197 = v193
	v198 = v195
	goto L15
L49:
	;
	F_appendBinaryStringInfo(m, l0, l1+v197, v198-v197)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v210 = v198
	goto L12
L52:
	;
	goto L51
L53:
	;
	if l2 != v224 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v48 = v220
	goto L8
L55:
	;
	v314 = v224 + int32(1)
	if v314 != v220 {
		v224 = v314
		goto L53
	} else {
		goto L89
	}
L56:
	;
	F_appendStringInfoString(m, l0, int32(498997))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L88
	}
L57:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v224))))
	switch v235 - int32(8) {
	case 0:
		goto L67
	case 1:
		goto L63
	case 2:
		goto L65
	case 3, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		goto L60
	case 4:
		goto L66
	case 5:
		goto L64
	case 26:
		goto L62
	default:
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v283 <= v284+int32(1) {
		goto L84
	} else {
		goto L85
	}
L60:
	;
	v258 = base.I32_extend8_s(v235)
	if base.Ui32(v235) <= base.Ui32(int32(31)) {
		goto L75
	} else {
		goto L76
	}
L61:
	;
	if v235 == int32(92) {
		goto L56
	} else {
		goto L74
	}
L62:
	;
	F_appendStringInfoString(m, l0, int32(704462))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L73
	}
L63:
	;
	F_appendStringInfoString(m, l0, int32(110663))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L72
	}
L64:
	;
	F_appendStringInfoString(m, l0, int32(227111))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L71
	}
L65:
	;
	F_appendStringInfoString(m, l0, int32(280895))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L70
	}
L66:
	;
	F_appendStringInfoString(m, l0, int32(334163))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	F_appendStringInfoString(m, l0, int32(495216))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	goto L55
L69:
	;
	goto L55
L70:
	;
	goto L55
L71:
	;
	goto L55
L72:
	;
	goto L55
L73:
	;
	goto L55
L74:
	;
	goto L60
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v258
	F_appendStringInfo(m, l0, int32(29136), v15)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v265 <= v266+int32(1) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L55
L79:
	;
	F_appendStringInfoChar(m, l0, v258)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v272+v266))) = uint8(v235)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v277 = v275 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v281 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v279+v277))) = uint8(v281)
	goto L55
L82:
	;
	goto L55
L83:
	;
	m.G0 = v15 + int32(16)
	return
L84:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v293 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v291+v284))) = uint8(v293)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v297 = v295 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v297
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v301 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v299+v297))) = uint8(v301)
	goto L83
L87:
	;
	goto L83
L88:
	;
	goto L55
L89:
	;
	goto L54
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
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
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
	var v221 int32
	_ = v221
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
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
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
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
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
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
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
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
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
	var v466 int32
	_ = v466
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
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
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
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
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
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
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
	var v701 int32
	_ = v701
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
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
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
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
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
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int64
	_ = v887
	var v890 int32
	_ = v890
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
	var v963 int32
	_ = v963
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
	var v1093 int32
	_ = v1093
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
	var v1210 int32
	_ = v1210
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
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
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
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1352 int32
	_ = v1352
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
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
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
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
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1596 int32
	_ = v1596
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1662 int32
	_ = v1662
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
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
		v1743 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v17 + int32(112)
	return v1743
L4:
	;
	v25 = l0
	goto L15
L5:
	;
	v1739 = F_copyObjectImpl(m, v25)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L1
	} else {
		goto L527
	}
L6:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1707)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+20)) = base.B2i32(v1732 == int32(0))
	v1743 = v1707
	goto L3
L7:
	;
	v1713 = F_palloc0(m, int32(36))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L1
	} else {
		goto L526
	}
L8:
	;
	if v1654&int32(1) != 0 {
		goto L519
	} else {
		goto L520
	}
L9:
	;
	v1685 = int32(0)
	v1687 = F_makeBoolConst(m, v1685, v1685)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L518
	}
L10:
	;
	if v1662&int32(1) != 0 {
		v1711 = v290
		goto L7
	} else {
		goto L516
	}
L11:
	;
	if v304&v305 == int32(0) {
		v1654 = v1624
		v1656 = v1626
		v1662 = v1632
		goto L10
	} else {
		goto L512
	}
L12:
	;
	if v1278 != 0 {
		goto L499
	} else {
		goto L500
	}
L13:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1511 = F_eval_const_expressions_mutator(m, v1510, l1)
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L478
	}
L14:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1451 = F_eval_const_expressions_mutator(m, v1450, l1)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L1
	} else {
		goto L459
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v39 != int32(319) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v1448 = F_expression_tree_mutator_impl(m, v25, int32(870), l1)
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L1
	} else {
		goto L458
	}
L17:
	;
	goto L16
L18:
	;
	switch v39 - int32(8) {
	case 0:
		goto L43
	default:
		goto L17
	case 3:
		goto L42
	case 6, 27, 28, 31:
		goto L28
	case 7:
		goto L41
	case 9:
		goto L40
	case 10:
		goto L39
	case 11:
		goto L38
	case 12:
		goto L37
	case 13:
		goto L36
	case 15, 16:
		v1743 = v25
		goto L3
	case 17:
		goto L25
	case 19:
		goto L34
	case 20:
		goto L33
	case 21:
		goto L32
	case 22:
		goto L21
	case 23:
		goto L31
	case 24:
		goto L30
	case 26:
		goto L29
	case 30:
		goto L27
	case 32:
		goto L26
	case 36:
		goto L35
	case 44:
		goto L24
	case 45:
		goto L13
	case 47:
		goto L14
	}
L19:
	;
	goto L20
L20:
	;
	v1440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1440 != int32(1) {
		goto L17
	} else {
		goto L455
	}
L21:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1399 = F_eval_const_expressions_mutator(m, v1398, l1)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L437
	}
L22:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v1743 = v1397
	goto L3
L23:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	v1386 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+8)))
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+28))
	v1391 = F_makeVar(m, v1385, v1386, v1387, v1388, v1389, v1390)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L436
	}
L24:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1276 = F_eval_const_expressions_mutator(m, v1275, l1)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L404
	}
L25:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1186 = F_eval_const_expressions_mutator(m, v1185, l1)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L372
	}
L26:
	;
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1177 != int32(1) {
		goto L5
	} else {
		goto L369
	}
L27:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v1091 != 0 {
		goto L349
	} else {
		goto L350
	}
L28:
	;
	v1077 = F_expression_tree_mutator_impl(m, v25, int32(870), l1)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L341
	}
L29:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1071 == int32(0) {
		goto L5
	} else {
		goto L339
	}
L30:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v946 = F_eval_const_expressions_mutator(m, v945, l1)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L314
	}
L31:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v933 = F_eval_const_expressions_mutator(m, v932, l1)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L308
	}
L32:
	;
	v883 = F_palloc0(m, int32(32))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L295
	}
L33:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v780
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v780
	v786 = F_list_make1_impl(m, int32(1), v17+int32(44))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L281
	}
L34:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v770 = F_eval_const_expressions_mutator(m, v769, l1)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L279
	}
L35:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v757 = F_eval_const_expressions_mutator(m, v756, l1)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L271
	}
L36:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	switch v479 {
	case 0:
		goto L176
	case 1:
		goto L177
	case 2:
		goto L175
	default:
		goto L174
	}
L37:
	;
	v451 = F_expression_tree_mutator_impl(m, v25, int32(870), l1)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L160
	}
L38:
	;
	v367 = F_expression_tree_mutator_impl(m, v25, int32(870), l1)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L134
	}
L39:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v290 = F_expression_tree_mutator_impl(m, v288, int32(870), l1)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L110
	}
L40:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v205
	F_set_opfuncid(m, v25)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L84
	}
L41:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v171 = F_exprTypmod(m, v25)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L80
	}
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v112 = F_SearchSysCache1(m, int32(47), v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L68
	}
L43:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v44 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v45 == int32(0) {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v48 <= int32(0) {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+28))
	if v51 < v48 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v53 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v65 == int32(0) {
		goto L5
	} else {
		goto L53
	}
L49:
	;
	v57 = m.T0[v53].(func(*base.Module, int32, int32, int32, int32) int32)(m, v45, v48, int32(1), v17+int32(100))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v64 = v48*int32(12) + v45 + int32(20)
	goto L48
L52:
	;
	v64 = v57
	goto L48
L53:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v65 != v68 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v70 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+6)))
	if v73&int32(1) == int32(0) {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_get_typlenbyval(m, v65, v17+int32(96), v17+int32(88))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
	if v84 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+96)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)))
	v106 = F_makeConst(m, v99, v100, v101, v102, v98, v97&int32(1), v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L67
	}
L61:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+96)))
	v94 = F_datumCopy(m, v91, int32(0), v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L66
	}
L62:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)))
	if v87 != int32(1) {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v97 = v84
	v98 = v90
	goto L60
L65:
	;
	goto L64
L66:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
	v97 = v96
	v98 = v94
	goto L60
L67:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = v108
	v1743 = v106
	goto L3
L68:
	;
	if v112 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v132 = F_expand_function_arguments(m, v129, int32(0), v131, v112)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L75
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v111
	F_errmsg_internal(m, int32(44089), v17)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(485655), int32(2544), int32(205697))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
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
	F_ReleaseCatCache(m, v112)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v137 = F_expression_tree_mutator_impl(m, v132, int32(870), l1)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v140 = F_eval_const_expressions_mutator(m, v139, l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v143 = F_palloc0(m, int32(44))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = int32(11)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+8)) = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+12)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+24)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v143)+20)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v143)+16)) = v153
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+28)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+32)) = v159
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+36)) = uint8(v161)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+37)) = uint8(v163)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+40)) = v165
	v1743 = v143
	goto L3
L80:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+13)))
	v178 = int32(1)
	v180 = F_simplify_function(m, v169, v170, v171, v173, v174, v17+int32(100), v177, v178, v178, l1)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if v180 != 0 {
		v1743 = v180
		goto L3
	} else {
		goto L82
	}
L82:
	;
	v183 = F_palloc0(m, int32(36))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = int32(15)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = v189
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+12)) = uint8(v191)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+13)) = uint8(v193)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+16)) = v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+20)) = v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+24)) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+28)) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+32)) = v203
	v1743 = v183
	goto L3
L84:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v217 = int32(1)
	v219 = F_simplify_function(m, v209, v210, int32(-1), v212, v213, v17+int32(100), int32(0), v217, v217, l1)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v219 != 0 {
		v1743 = v219
		goto L3
	} else {
		goto L86
	}
L86:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	switch v221 - int32(85) {
	case 0, 6:
		goto L88
	default:
		goto L89
	}
L87:
	;
	v269 = F_palloc0(m, int32(36))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L109
	}
L88:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	if v228 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v262 = v224
	goto L87
L90:
	;
	if v258 != 0 {
		v1743 = v258
		goto L3
	} else {
		goto L108
	}
L91:
	;
	v256 = F_negate_clause(m, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L107
	}
L92:
	;
	v239 = int32(0)
	if v227 == v239 {
		v258 = v239
		goto L90
	} else {
		goto L100
	}
L93:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	if v231 != int32(7) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v228)+20))
	if v221 == int32(91) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if v234 == int32(0) {
		v254 = v227
		goto L91
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v234 != 0 {
		v254 = v227
		goto L91
	} else {
		goto L99
	}
L98:
	;
	v258 = v227
	goto L90
L99:
	;
	v258 = v227
	goto L90
L100:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if v242 != int32(7) {
		v258 = v239
		goto L90
	} else {
		goto L101
	}
L101:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v227)+20))
	if v221 == int32(91) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if v245 == int32(0) {
		v254 = v228
		goto L91
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if v245 == int32(0) {
		v258 = v228
		goto L90
	} else {
		goto L106
	}
L105:
	;
	v258 = v228
	goto L90
L106:
	;
	v254 = v228
	goto L91
L107:
	;
	v258 = v256
	goto L90
L108:
	;
	v262 = v225
	goto L87
L109:
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
	*(*int32)(unsafe.Add(mBase, uint32(v269)+28)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v269)+24)) = v283
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+32)) = v286
	v1743 = v269
	goto L3
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v290
	if v290 == int32(0) {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	v295 = int32(0)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	if v296 <= v295 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v1654 = v295
	v1656 = int32(1)
	v1662 = v3
	goto L10
L113:
	;
	goto L114
L114:
	;
	v300 = int32(1)
	v301 = int32(0)
	if v301 < v296 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v304 = v296
	goto L117
L116:
	;
	v304 = v301
	goto L117
L117:
	;
	v305 = int32(1)
	if v296 == v305 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v309 = int32(0)
	v1623 = v309
	v1624 = v309
	v1626 = v300
	v1632 = v3
	goto L11
L119:
	;
	goto L120
L120:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	v314 = int32(0)
	v318 = v314
	v319 = v314
	v321 = v300
	v322 = v3
	v327 = v3
	goto L121
L121:
	;
	v332 = v313 + v318<<(uint(int32(2))%32)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	if v334 != int32(7) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	v1623 = v362
	v1624 = v357
	v1626 = v358
	v1632 = v360
	goto L11
L123:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	if v348 != int32(7) {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v343 = v319
	v344 = v321
	v346 = int32(1)
	goto L123
L125:
	;
	goto L126
L126:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+24)))
	v343 = (v319 | v338) & int32(1)
	v344 = v321 & v338
	v346 = v327
	goto L123
L127:
	;
	v361 = int32(2)
	v362 = v318 + v361
	v364 = v322 + v361
	if v304&int32(2147483646) != v364 {
		v318 = v362
		v319 = v357
		v321 = v358
		v322 = v364
		v327 = v360
		goto L121
	} else {
		goto L131
	}
L128:
	;
	v357 = v343
	v358 = v344
	v360 = int32(1)
	goto L127
L129:
	;
	goto L130
L130:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+24)))
	v357 = (v343 | v352) & int32(1)
	v358 = v344 & v352
	v360 = v346
	goto L127
L131:
	;
	goto L122
L132:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	v431 = F_func_volatile(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L150
	}
L133:
	;
	F_set_opfuncid(m, v367)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L149
	}
L134:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v367)+28))
	if v369 == int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v372 <= int32(0) {
		goto L133
	} else {
		goto L136
	}
L136:
	;
	v375 = int32(0)
	if v375 < v372 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v379 = v372
	goto L139
L138:
	;
	v379 = v375
	goto L139
L139:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v382 = v375
	v387 = int32(0)
	goto L140
L140:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v380+v382<<(uint(int32(2))%32))))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	if v401 == int32(7) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	F_set_opfuncid(m, v367)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L147
	}
L142:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+24)))
	if v404 != 0 {
		goto L22
	} else {
		goto L145
	}
L143:
	;
	v405 = int32(1)
	goto L144
L144:
	;
	v407 = v382 + int32(1)
	if v407 != v379 {
		v382 = v407
		v387 = v405
		goto L140
	} else {
		goto L146
	}
L145:
	;
	v405 = v387
	goto L144
L146:
	;
	goto L141
L147:
	;
	if v405&int32(1) != 0 {
		v1743 = v367
		goto L3
	} else {
		goto L148
	}
L148:
	;
	goto L132
L149:
	;
	goto L132
L150:
	;
	if v431 != int32(105) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if v431 != int32(115) {
		v1743 = v367
		goto L3
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v442 = F_exprType(m, v367)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v437&int32(1) == int32(0) {
		v1743 = v367
		goto L3
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v444 = F_exprTypmod(m, v367)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v446 = F_exprCollation(m, v367)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v448 = F_evaluate_expr(m, v367, v442, v444, v446)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v1743 = v448
	goto L3
L160:
	;
	F_set_opfuncid(m, v451)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v457 = F_expression_tree_walker_impl(m, v451, int32(871), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	if v457 != 0 {
		v1743 = v451
		goto L3
	} else {
		goto L163
	}
L163:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v451)+8))
	v460 = F_func_volatile(m, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	if v460 != int32(105) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	if v460 != int32(115) {
		v1743 = v451
		goto L3
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v471 = F_exprType(m, v451)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v466&int32(1) == int32(0) {
		v1743 = v451
		goto L3
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v473 = F_exprTypmod(m, v451)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v475 = F_exprCollation(m, v451)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v477 = F_evaluate_expr(m, v451, v471, v473, v475)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v1743 = v477
	goto L3
L174:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L268
	}
L175:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)+12))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)))
	v735 = F_eval_const_expressions_mutator(m, v734, l1)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L266
	}
L176:
	;
	v609 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+100)) = uint8(v609)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+96)) = uint8(v609)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v619 = F_list_copy(m, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L223
	}
L177:
	;
	v480 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+100)) = uint8(v480)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+96)) = uint8(v480)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v490 = F_list_copy(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L179
	}
L178:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+96)))
	if v579 == int32(1) {
		goto L205
	} else {
		goto L206
	}
L179:
	;
	if v490 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v492 = v490
	v494 = v480
	goto L183
L181:
	;
	v552 = v480
	goto L182
L182:
	;
	v578 = v552
	goto L178
L183:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)))
	v508 = F_list_delete_first(m, v492)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L185
	}
L184:
	;
	v552 = v548
	goto L182
L185:
	;
	if v507 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	if v547 != 0 {
		v492 = v547
		v494 = v548
		goto L183
	} else {
		goto L204
	}
L187:
	;
	v523 = F_eval_const_expressions_mutator(m, v507, l1)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L196
	}
L188:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	if v512 != int32(21) {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	if v515 != int32(1) {
		goto L187
	} else {
		goto L190
	}
L190:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	v519 = F_list_concat_copy(m, v518, v508)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_list_free(m, v508)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v547 = v519
	v548 = v494
	goto L186
L193:
	;
	v545 = F_lappend(m, v494, v523)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L203
	}
L194:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+24)))
	if v534 == int32(1) {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if v528 != int32(1) {
		goto L193
	} else {
		goto L197
	}
L196:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	switch v525 - int32(7) {
	case 0:
		goto L194
	default:
		goto L193
	case 14:
		goto L195
	}
L197:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v523)+8))
	v532 = F_list_concat_copy(m, v531, v508)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v547 = v532
	v548 = v494
	goto L186
L199:
	;
	v537 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(100)))) = uint8(v537)
	v547 = v508
	v548 = v494
	goto L186
L200:
	;
	goto L201
L201:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v523)+20))
	if v539 == int32(0) {
		v547 = v508
		v548 = v494
		goto L186
	} else {
		goto L202
	}
L202:
	;
	v542 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(96)))) = uint8(v542)
	v578 = int32(0)
	goto L178
L203:
	;
	v547 = v508
	v548 = v545
	goto L186
L204:
	;
	goto L184
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
	v1743 = v584
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
	v1743 = v600
	goto L3
L218:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v595)+12))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	v1743 = v606
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
	v1743 = v607
	goto L3
L222:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+96)))
	if v702 == int32(1) {
		goto L249
	} else {
		goto L250
	}
L223:
	;
	if v619 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v621 = v619
	v623 = v609
	goto L227
L225:
	;
	v675 = v609
	goto L226
L226:
	;
	v701 = v675
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
	v675 = v671
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
	if v670 != 0 {
		v621 = v670
		v623 = v671
		goto L227
	} else {
		goto L248
	}
L231:
	;
	v650 = F_eval_const_expressions_mutator(m, v636, l1)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L240
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
	v670 = v646
	v671 = v623
	goto L230
L237:
	;
	v668 = F_lappend(m, v623, v650)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L247
	}
L238:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650)+24)))
	if v659 == int32(1) {
		goto L243
	} else {
		goto L244
	}
L239:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v650)+4))
	if v655 != 0 {
		goto L237
	} else {
		goto L241
	}
L240:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v650)))
	switch v652 - int32(7) {
	case 0:
		goto L238
	default:
		goto L237
	case 14:
		goto L239
	}
L241:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v650)+8))
	v657 = F_list_concat_copy(m, v656, v637)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v670 = v657
	v671 = v623
	goto L230
L243:
	;
	v662 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(100)))) = uint8(v662)
	v670 = v637
	v671 = v623
	goto L230
L244:
	;
	goto L245
L245:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v650)+20))
	if v664 != 0 {
		v670 = v637
		v671 = v623
		goto L230
	} else {
		goto L246
	}
L246:
	;
	v665 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(96)))) = uint8(v665)
	v701 = int32(0)
	goto L222
L247:
	;
	v670 = v637
	v671 = v668
	goto L230
L248:
	;
	goto L228
L249:
	;
	v705 = int32(0)
	v707 = F_makeBoolConst(m, v705, v705)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+100)))
	if v709 == int32(1) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1743 = v707
	goto L3
L253:
	;
	v714 = F_makeBoolConst(m, int32(0), int32(1))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L256
	}
L254:
	;
	v718 = v701
	goto L255
L255:
	;
	if v718 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	v716 = F_lappend(m, v701, v714)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v718 = v716
	goto L255
L258:
	;
	v723 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	if v725 == int32(1) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v1743 = v723
	goto L3
L262:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v718)+12))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	v1743 = v729
	goto L3
L263:
	;
	goto L264
L264:
	;
	v730 = F_make_andclause(m, v718)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v1743 = v730
	goto L3
L266:
	;
	v737 = F_negate_clause(m, v735)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	v1743 = v737
	goto L3
L268:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v743
	F_errmsg_internal(m, int32(474480), v17+int32(16))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(485655), int32(2910), int32(205697))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	if v757 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	if v759 == int32(7) {
		v1743 = v757
		goto L3
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v762 = F_eval_const_expressions_mutator(m, v755, l1)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L276
	}
L275:
	;
	goto L274
L276:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v765 = F_copyObjectImpl(m, v764)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	v767 = F_makeJsonValueExpr(m, v762, v757, v765)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	v1743 = v767
	goto L3
L279:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v778 = F_applyRelabelType(m, v770, v772, v773, v774, v775, v776, int32(1))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v1743 = v778
	goto L3
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v786
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v790 = F_exprType(m, v789)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	F_getTypeOutputInfo(m, v790, v17+int32(96), v17+int32(95))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	F_getTypeInputInfo(m, v798, v17+int32(88), v17+int32(84))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v17)+96))
	v808 = int32(0)
	v813 = int32(1)
	v815 = F_simplify_function(m, v805, int32(2275), int32(-1), v808, v808, v17+int32(100), v808, v813, v813, l1)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	if v815 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v815
	v820 = int32(0)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v17)+84))
	v825 = F_makeConst(m, int32(26), int32(-1), v820, int32(4), v822, v820, int32(1))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v866 = F_palloc0(m, int32(24))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L294
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v825
	v829 = int32(-1)
	v830 = int32(0)
	v835 = F_makeConst(m, int32(23), v829, v830, int32(4), v829, v830, int32(1))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v835
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v839
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v841
	v849 = F_list_make3_impl(m, v17+int32(40), v17+int32(36), v17+int32(32))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v849
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v856 = int32(0)
	v862 = F_simplify_function(m, v852, v853, int32(-1), v855, v856, v17+int32(100), v856, v856, int32(1), l1)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	if v862 != 0 {
		v1743 = v862
		goto L3
	} else {
		goto L293
	}
L293:
	;
	goto L288
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v866))) = int32(28)
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)+12))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v871)))
	*(*int32)(unsafe.Add(mBase, uint32(v866)+4)) = v872
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v866)+8)) = v874
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v866)+12)) = v876
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v866)+16)) = v878
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v866)+20)) = v880
	v1743 = v866
	goto L3
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v883))) = int32(29)
	v887 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v883))) = v887
	v890 = v883 + int32(8)
	v891 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v890))) = v891
	v893 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v883)+16)) = v893
	v895 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v883)+24)) = v895
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v883)+4))
	v898 = F_eval_const_expressions_mutator(m, v897, l1)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v883)+4)) = v898
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	v905 = F_eval_const_expressions_mutator(m, v904, l1)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v890))) = v905
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v901
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v883)+4))
	if v909 == int32(0) {
		v1743 = v883
		goto L3
	} else {
		goto L298
	}
L298:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v909)))
	if v912 != int32(7) {
		v1743 = v883
		goto L3
	} else {
		goto L299
	}
L299:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v883)+8))
	if v915 == int32(0) {
		v1743 = v883
		goto L3
	} else {
		goto L300
	}
L300:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	if v918 == int32(55) {
		v1743 = v883
		goto L3
	} else {
		goto L301
	}
L301:
	;
	v922 = F_contain_mutable_functions_walker(m, v915, int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	if v922 != 0 {
		v1743 = v883
		goto L3
	} else {
		goto L303
	}
L303:
	;
	v924 = F_exprType(m, v883)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v926 = F_exprTypmod(m, v883)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v928 = F_exprCollation(m, v883)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	v930 = F_evaluate_expr(m, v883, v924, v926, v928)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	v1743 = v930
	goto L3
L308:
	;
	v935 = F_exprType(m, v933)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v937 = F_exprTypmod(m, v933)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v943 = F_applyRelabelType(m, v933, v935, v937, v939, int32(2), v941, int32(1))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	v1743 = v943
	goto L3
L312:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v955
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v959 == int32(0) {
		v1029 = v3
		goto L318
	} else {
		goto L319
	}
L313:
	;
	v955 = int32(0)
	v956 = v946
	goto L312
L314:
	;
	if v946 == int32(0) {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	if v950 != int32(7) {
		goto L313
	} else {
		goto L316
	}
L316:
	;
	v955 = v946
	v956 = int32(0)
	goto L312
L317:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1050)))
	v1052 = F_eval_const_expressions_mutator(m, v1051, l1)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L334
	}
L318:
	;
	v1045 = v1029
	v1050 = v25 + int32(20)
	goto L317
L319:
	;
	v962 = int32(0)
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	if v963 <= v962 {
		v1029 = v3
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v969 = v962
	v975 = v3
	goto L321
L321:
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
		goto L325
	}
L322:
	;
	v1029 = v1014
	goto L318
L323:
	;
	v1017 = v969 + int32(1)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	if v1017 < v1018 {
		v969 = v1017
		v975 = v1014
		goto L321
	} else {
		goto L333
	}
L324:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v984)+8))
	v1000 = F_eval_const_expressions_mutator(m, v999, l1)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L330
	}
L325:
	;
	if v986 == int32(0) {
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	if v990 != int32(7) {
		goto L324
	} else {
		goto L327
	}
L327:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986)+24)))
	if v993 != 0 {
		v1014 = v975
		goto L323
	} else {
		goto L328
	}
L328:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v986)+20))
	if v994 == int32(0) {
		v1014 = v975
		goto L323
	} else {
		goto L329
	}
L329:
	;
	v1045 = v975
	v1050 = v984 + int32(8)
	goto L317
L330:
	;
	v1003 = F_palloc0(m, int32(16))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
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
		goto L332
	}
L332:
	;
	v1014 = v1011
	goto L323
L333:
	;
	goto L322
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v957
	if v1045 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1743 = v1052
	goto L3
L336:
	;
	goto L337
L337:
	;
	v1058 = F_palloc0(m, int32(28))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
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
	v1743 = v1058
	goto L3
L339:
	;
	v1074 = F_copyObjectImpl(m, v1071)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	v1743 = v1074
	goto L3
L341:
	;
	v1081 = F_expression_tree_walker_impl(m, v1077, int32(871), int32(0))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	if v1081 != 0 {
		v1743 = v1077
		goto L3
	} else {
		goto L343
	}
L343:
	;
	v1083 = F_exprType(m, v1077)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	v1085 = F_exprTypmod(m, v1077)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	v1087 = F_exprCollation(m, v1077)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v1089 = F_evaluate_expr(m, v1077, v1083, v1085, v1087)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	v1743 = v1089
	goto L3
L348:
	;
	v1166 = F_palloc0(m, int32(20))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L368
	}
L349:
	;
	v1092 = int32(0)
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+4))
	if v1093 <= v1092 {
		v1138 = v3
		goto L352
	} else {
		goto L353
	}
L350:
	;
	goto L351
L351:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1163 = F_makeNullConst(m, v1160, int32(-1), v1162)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L367
	}
L352:
	;
	if v1138 != 0 {
		goto L348
	} else {
		goto L366
	}
L353:
	;
	v1099 = v1092
	v1102 = v3
	goto L354
L354:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+12))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1110+v1099<<(uint(int32(2))%32))))
	v1115 = F_eval_const_expressions_mutator(m, v1114, l1)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L357
	}
L355:
	;
	v1138 = v1127
	goto L352
L356:
	;
	v1129 = v1099 + int32(1)
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+4))
	if v1129 < v1130 {
		v1099 = v1129
		v1102 = v1127
		goto L354
	} else {
		goto L365
	}
L357:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1115)))
	if v1117 == int32(7) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115)+24)))
	if v1120 != 0 {
		v1127 = v1102
		goto L356
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1125 = F_lappend(m, v1102, v1115)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L364
	}
L361:
	;
	if v1102 == int32(0) {
		v1743 = v1115
		goto L3
	} else {
		goto L362
	}
L362:
	;
	v1123 = F_lappend(m, v1102, v1115)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	v1138 = v1123
	goto L352
L364:
	;
	v1127 = v1125
	goto L356
L365:
	;
	goto L355
L366:
	;
	goto L351
L367:
	;
	v1743 = v1163
	goto L3
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1166))) = int32(38)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+4)) = v1170
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+12)) = v1138
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+8)) = v1172
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+16)) = v1175
	v1743 = v1166
	goto L3
L369:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1183 = F_evaluate_expr(m, v25, v1180, v1181, int32(0))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	v1743 = v1183
	goto L3
L371:
	;
	v1244 = F_palloc0(m, int32(24))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L1
	} else {
		goto L395
	}
L372:
	;
	if v1186 == int32(0) {
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	if v1190 == int32(6) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1186)+8)))
	if v1193 != 0 {
		goto L371
	} else {
		goto L377
	}
L375:
	;
	v1203 = v1190
	goto L376
L376:
	;
	if v1203 != int32(36) {
		goto L371
	} else {
		goto L381
	}
L377:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+28))
	if v1194 != 0 {
		goto L371
	} else {
		goto L378
	}
L378:
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
		goto L379
	}
L379:
	;
	if v1200 != 0 {
		goto L23
	} else {
		goto L380
	}
L380:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	v1203 = v1202
	goto L376
L381:
	;
	v1206 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+8)))
	if v1206 <= int32(0) {
		goto L371
	} else {
		goto L382
	}
L382:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	if v1209 != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+4))
	v1212 = v1210
	goto L385
L384:
	;
	v1212 = int32(0)
	goto L385
L385:
	;
	if v1212 < v1206 {
		goto L371
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
		goto L371
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
		goto L371
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
		goto L371
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
		v1743 = v1220
		goto L3
	} else {
		goto L394
	}
L394:
	;
	goto L371
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+4)) = v1186
	*(*int32)(unsafe.Add(mBase, uint32(v1244))) = int32(25)
	v1249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1244)+8)) = uint16(v1249)
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+12)) = v1251
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+16)) = v1253
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+20)) = v1255
	if v1186 == int32(0) {
		v1743 = v1244
		goto L3
	} else {
		goto L396
	}
L396:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	if v1259 != int32(7) {
		v1743 = v1244
		goto L3
	} else {
		goto L397
	}
L397:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	v1263 = F_rowtype_field_matches(m, v1262, v1249, v1251, v1253, v1255)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	if v1263 == int32(0) {
		v1743 = v1244
		goto L3
	} else {
		goto L399
	}
L399:
	;
	v1267 = F_exprType(m, v1244)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	v1269 = F_exprTypmod(m, v1244)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	v1271 = F_exprCollation(m, v1244)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v1273 = F_evaluate_expr(m, v1244, v1267, v1269, v1271)
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	v1743 = v1273
	goto L3
L404:
	;
	v1278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)))
	if v1278 != int32(1) {
		goto L12
	} else {
		goto L405
	}
L405:
	;
	if v1276 == int32(0) {
		goto L12
	} else {
		goto L406
	}
L406:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1276)))
	if v1283 != int32(36) {
		goto L12
	} else {
		goto L407
	}
L407:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+4))
	if v1286 != 0 {
		goto L409
	} else {
		goto L410
	}
L408:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1352)+4))
	if v1378 == int32(1) {
		goto L432
	} else {
		goto L433
	}
L409:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+4))
	if int32(0) < v1287 {
		goto L412
	} else {
		goto L413
	}
L410:
	;
	goto L411
L411:
	;
	v1376 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L1
	} else {
		goto L431
	}
L412:
	;
	v1292 = int32(0)
	v1297 = v3
	goto L415
L413:
	;
	v1352 = v3
	goto L414
L414:
	;
	if v1352 != 0 {
		goto L408
	} else {
		goto L430
	}
L415:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+12))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1305+v1292<<(uint(int32(2))%32))))
	if v1309 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L416:
	;
	v1352 = v1341
	goto L414
L417:
	;
	v1343 = v1292 + int32(1)
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+4))
	if v1343 < v1344 {
		v1292 = v1343
		v1297 = v1341
		goto L415
	} else {
		goto L429
	}
L418:
	;
	v1335 = int32(0)
	v1337 = F_makeBoolConst(m, v1335, v1335)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L1
	} else {
		goto L428
	}
L419:
	;
	if v1315 != 0 {
		v1341 = v1297
		goto L417
	} else {
		goto L427
	}
L420:
	;
	v1322 = F_palloc0(m, int32(20))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L425
	}
L421:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1309)))
	if v1312 != int32(7) {
		goto L420
	} else {
		goto L422
	}
L422:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1309)+24)))
	if v1316 != int32(1) {
		goto L419
	} else {
		goto L423
	}
L423:
	;
	if v1315 != int32(1) {
		v1341 = v1297
		goto L417
	} else {
		goto L424
	}
L424:
	;
	goto L418
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1322)+4)) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(v1322))) = int32(52)
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1322)+12)) = uint8(v1328)
	*(*int32)(unsafe.Add(mBase, uint32(v1322)+8)) = v1327
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1322)+16)) = v1331
	v1333 = F_lappend(m, v1297, v1322)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	v1341 = v1333
	goto L417
L427:
	;
	goto L418
L428:
	;
	v1743 = v1337
	goto L3
L429:
	;
	goto L416
L430:
	;
	goto L411
L431:
	;
	v1743 = v1376
	goto L3
L432:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1352)+12))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1381)))
	v1743 = v1382
	goto L3
L433:
	;
	goto L434
L434:
	;
	v1383 = F_make_andclause(m, v1352)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	v1743 = v1383
	goto L3
L436:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1391)+32)) = v1393
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1391)+24)) = v1395
	v1743 = v1391
	goto L3
L437:
	;
	v1402 = F_palloc0(m, int32(20))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1402))) = int32(30)
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+8)) = v1406
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+12)) = v1408
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+16)) = v1410
	if v1399 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+4)) = int32(0)
	v1743 = v1402
	goto L3
L440:
	;
	goto L441
L441:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1399)))
	if v1416 != int32(30) {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1428)))
	if v1429 != int32(7) {
		v1743 = v1402
		goto L3
	} else {
		goto L450
	}
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+4)) = v1399
	v1428 = v1399
	goto L442
L444:
	;
	goto L445
L445:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1399)+4))
	if v1408 == int32(2) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1399)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+12)) = v1423
	goto L448
L447:
	;
	goto L448
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+4)) = v1420
	if v1420 == int32(0) {
		v1743 = v1402
		goto L3
	} else {
		goto L449
	}
L449:
	;
	v1428 = v1420
	goto L442
L450:
	;
	v1432 = F_exprType(m, v1402)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	v1434 = F_exprTypmod(m, v1402)
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	v1436 = F_exprCollation(m, v1402)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	v1438 = F_evaluate_expr(m, v1402, v1432, v1434, v1436)
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	v1743 = v1438
	goto L3
L455:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	if v1443 != 0 {
		v25 = v1443
		goto L15
	} else {
		goto L457
	}
L457:
	;
	v1743 = v3
	goto L3
L458:
	;
	v1743 = v1448
	goto L3
L459:
	;
	v1453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1453 == int32(0) {
		goto L461
	} else {
		goto L462
	}
L460:
	;
	v1495 = F_palloc0(m, int32(28))
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L1
	} else {
		goto L476
	}
L461:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1457 = F_DomainHasConstraints(m, v1456)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L464
	}
L462:
	;
	goto L463
L463:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1459 == int32(0) {
		goto L466
	} else {
		goto L467
	}
L464:
	;
	if v1457 != 0 {
		goto L460
	} else {
		goto L465
	}
L465:
	;
	goto L463
L466:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v1492 = F_applyRelabelType(m, v1451, v1486, v1487, v1488, v1489, v1490, int32(1))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L1
	} else {
		goto L475
	}
L467:
	;
	v1462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1462 != 0 {
		goto L466
	} else {
		goto L468
	}
L468:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if base.Ui32(int32(12000)) <= base.Ui32(v1463) {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v1467 = F_palloc0(m, int32(12))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L472
	}
L470:
	;
	goto L471
L471:
	;
	goto L466
L472:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1467))) = int64(352187318651)
	v1473 = F_GetSysCacheHashValue(m, int32(82), v1463, int32(0))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L1
	} else {
		goto L473
	}
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+8)) = v1473
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+8))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1476)+60))
	v1478 = F_lappend(m, v1477, v1467)
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L1
	} else {
		goto L474
	}
L474:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1480)+60)) = v1478
	goto L471
L475:
	;
	v1743 = v1492
	goto L3
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1495)+4)) = v1451
	*(*int32)(unsafe.Add(mBase, uint32(v1495))) = int32(55)
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1495)+8)) = v1500
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1495)+12)) = v1502
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1495)+16)) = v1504
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1495)+20)) = v1506
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1495)+24)) = v1508
	v1743 = v1495
	goto L3
L477:
	;
	v1568 = F_palloc0(m, int32(16))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L1
	} else {
		goto L498
	}
L478:
	;
	if v1511 == int32(0) {
		goto L477
	} else {
		goto L479
	}
L479:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	if v1515 != int32(7) {
		goto L477
	} else {
		goto L480
	}
L480:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	switch v1518 {
	case 0:
		goto L489
	case 1:
		goto L488
	case 2:
		goto L487
	case 3:
		goto L486
	case 4:
		goto L485
	case 5:
		goto L484
	default:
		goto L483
	}
L481:
	;
	v1565 = F_makeBoolConst(m, v1561&int32(1), int32(0))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L1
	} else {
		goto L497
	}
L482:
	;
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+20))
	v1561 = base.B2i32(v1558 != int32(0))
	goto L481
L483:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L1
	} else {
		goto L494
	}
L484:
	;
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511)+24)))
	v1561 = v1539 ^ int32(1)
	goto L481
L485:
	;
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511)+24)))
	v1561 = v1538
	goto L481
L486:
	;
	v1534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511)+24)))
	if v1534 != 0 {
		v1561 = int32(1)
		goto L481
	} else {
		goto L493
	}
L487:
	;
	v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511)+24)))
	if v1529 != 0 {
		v1561 = int32(0)
		goto L481
	} else {
		goto L492
	}
L488:
	;
	v1524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511)+24)))
	if v1524 != 0 {
		v1561 = int32(1)
		goto L481
	} else {
		goto L491
	}
L489:
	;
	v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511)+24)))
	if v1519 == int32(0) {
		goto L482
	} else {
		goto L490
	}
L490:
	;
	v1561 = int32(0)
	goto L481
L491:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+20))
	v1561 = base.B2i32(v1525 == int32(0))
	goto L481
L492:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+20))
	v1561 = base.B2i32(v1530 == int32(0))
	goto L481
L493:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+20))
	v1561 = base.B2i32(v1535 != int32(0))
	goto L481
L494:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v1546
	F_errmsg_internal(m, int32(475707), v17-int32(-64))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	F_errfinish(m, int32(485655), int32(3593), int32(205697))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L497:
	;
	v1743 = v1565
	goto L3
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+4)) = v1511
	*(*int32)(unsafe.Add(mBase, uint32(v1568))) = int32(53)
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+8)) = v1573
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+12)) = v1575
	v1743 = v1568
	goto L3
L499:
	;
	v1610 = F_palloc0(m, int32(20))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L1
	} else {
		goto L511
	}
L500:
	;
	if v1276 == int32(0) {
		goto L499
	} else {
		goto L501
	}
L501:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1276)))
	if v1579 != int32(7) {
		goto L499
	} else {
		goto L502
	}
L502:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	switch v1582 {
	case 0:
		goto L504
	case 1:
		goto L506
	default:
		goto L505
	}
L503:
	;
	v1607 = F_makeBoolConst(m, v1603&int32(1), int32(0))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L1
	} else {
		goto L510
	}
L504:
	;
	v1602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1276)+24)))
	v1603 = v1602
	goto L503
L505:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L1
	} else {
		goto L507
	}
L506:
	;
	v1583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1276)+24)))
	v1603 = v1583 ^ int32(1)
	goto L503
L507:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v1590
	F_errmsg_internal(m, int32(475737), v17+int32(48))
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	F_errfinish(m, int32(485655), int32(3532), int32(205697))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L510:
	;
	v1743 = v1607
	goto L3
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+4)) = v1276
	*(*int32)(unsafe.Add(mBase, uint32(v1610))) = int32(52)
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+8)) = v1615
	v1617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1610)+12)) = uint8(v1617)
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+16)) = v1619
	v1743 = v1610
	goto L3
L512:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1637+v1623<<(uint(int32(2))%32))))
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1641)))
	if v1642 != int32(7) {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v1654 = v1624
	v1656 = v1626
	v1662 = int32(1)
	goto L10
L514:
	;
	goto L515
L515:
	;
	v1646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1641)+24)))
	v1654 = (v1646 | v1624) & int32(1)
	v1656 = v1626 & v1646
	v1662 = v1632
	goto L10
L516:
	;
	if v1656&int32(1) == int32(0) {
		goto L8
	} else {
		goto L517
	}
L517:
	;
	goto L9
L518:
	;
	v1743 = v1687
	goto L3
L519:
	;
	v1693 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L1
	} else {
		goto L522
	}
L520:
	;
	goto L521
L521:
	;
	F_set_opfuncid(m, v25)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L1
	} else {
		goto L523
	}
L522:
	;
	v1743 = v1693
	goto L3
L523:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v1704 = int32(0)
	v1707 = F_simplify_function(m, v1697, v1698, int32(-1), v1700, v1701, v17+int32(100), v1704, v1704, v1704, l1)
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	if v1707 != 0 {
		goto L6
	} else {
		goto L525
	}
L525:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v1711 = v1709
	goto L7
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1713))) = int32(18)
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+4)) = v1717
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+8)) = v1719
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+12)) = v1721
	v1723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1713)+16)) = uint8(v1723)
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+20)) = v1725
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+28)) = v1711
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+24)) = v1727
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+32)) = v1730
	v1743 = v1713
	goto L3
L527:
	;
	v1743 = v1739
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v14 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		v22 = F_AllocSetContextCreateInternal(m, v17, int32(495698), int32(0), int32(8192), int32(8388608))
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
					v40 = int32(4470560)
					v41 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
					*(*int32)(unsafe.Add(mBase, _consts[9])) = v44
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
							*(*int32)(unsafe.Add(mBase, _consts[9])) = v41
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
													v77 = v12 + int32(8)
													v78 = m.G0
													v80 = v78 - int32(48)
													m.G0 = v80
													if v57 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(119827), int32(0))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(489192), int32(1541), int32(277050))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														if v77 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(119827), int32(0))
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(489192), int32(1541), int32(277050))
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, _consts[362]))
															if v87 == int32(0) {
																*(*int32)(unsafe.Add(mBase, _consts[365])) = int32(-4)
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v165 = m.ExcPending
																if v165 != 0 {
																	return int32(0)
																} else {
																	F_errmsg_internal(m, int32(440278), int32(0))
																	mBase = m.M
																	v169 = m.ExcPending
																	if v169 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(489192), int32(1545), int32(277050))
																		mBase = m.M
																		v174 = m.ExcPending
																		if v174 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															} else {
																v91 = *(*int32)(unsafe.Add(mBase, _consts[39]))
																v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
																v94 = *(*int32)(unsafe.Add(mBase, _consts[362]))
																*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v92
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
																*(*int32)(unsafe.Add(mBase, _consts[9])) = v97
																v99 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v80)+20)) = v99
																v102 = v80 + int32(28)
																*(*int64)(unsafe.Add(mBase, uint32(v102))) = v99
																*(*int64)(unsafe.Add(mBase, uint32(v80)+36)) = v99
																v107 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v107
																*(*int32)(unsafe.Add(mBase, _consts[365])) = v107
																*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v99
																*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(569278163)
																v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
																if v118 != 0 {
																	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
																	*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v119
																	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
																	*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v121
																} else {
																}
																F__SPI_prepare_plan(m, v57, v80+int32(8))
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return int32(0)
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
																	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
																	v131 = F_SPI_cursor_open_internal(m, l3, v80+int32(8), v129, v130)
																	mBase = m.M
																	v132 = m.ExcPending
																	if v132 != 0 {
																		return int32(0)
																	} else {
																		v135 = *(*int32)(unsafe.Add(mBase, _consts[362]))
																		v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
																		*(*int32)(unsafe.Add(mBase, _consts[9])) = v136
																		*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(0)
																		v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
																		F_MemoryContextReset(m, v140)
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v80 + int32(48)
																			if v131 == int32(0) {
																				F_errstart_cold(m, int32(21), int32(540539))
																				mBase = m.M
																				v202 = m.ExcPending
																				if v202 != 0 {
																					return int32(0)
																				} else {
																					v204 = *(*int32)(unsafe.Add(mBase, _consts[365]))
																					v205 = F_SPI_result_code_string(m, v204)
																					mBase = m.M
																					v206 = m.ExcPending
																					if v206 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v205
																						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																						F_errmsg_internal(m, int32(201994), v12)
																						mBase = m.M
																						v211 = m.ExcPending
																						if v211 != 0 {
																							return int32(0)
																						} else {
																							F_errfinish(m, int32(491440), int32(8998), int32(148481))
																							mBase = m.M
																							v216 = m.ExcPending
																							if v216 != 0 {
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
																				v178 = m.ExcPending
																				if v178 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v12 + int32(32)
																					return v131
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
												v77 = v12 + int32(8)
												v78 = m.G0
												v80 = v78 - int32(48)
												m.G0 = v80
												if v57 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(119827), int32(0))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(489192), int32(1541), int32(277050))
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													if v77 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(119827), int32(0))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(489192), int32(1541), int32(277050))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, _consts[362]))
														if v87 == int32(0) {
															*(*int32)(unsafe.Add(mBase, _consts[365])) = int32(-4)
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v165 = m.ExcPending
															if v165 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(440278), int32(0))
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(489192), int32(1545), int32(277050))
																	mBase = m.M
																	v174 = m.ExcPending
																	if v174 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, _consts[39]))
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
															v94 = *(*int32)(unsafe.Add(mBase, _consts[362]))
															*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v92
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
															*(*int32)(unsafe.Add(mBase, _consts[9])) = v97
															v99 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v80)+20)) = v99
															v102 = v80 + int32(28)
															*(*int64)(unsafe.Add(mBase, uint32(v102))) = v99
															*(*int64)(unsafe.Add(mBase, uint32(v80)+36)) = v99
															v107 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v107
															*(*int32)(unsafe.Add(mBase, _consts[365])) = v107
															*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v99
															*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(569278163)
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
															if v118 != 0 {
																v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
																*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v119
																v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
																*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v121
															} else {
															}
															F__SPI_prepare_plan(m, v57, v80+int32(8))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
																v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
																v131 = F_SPI_cursor_open_internal(m, l3, v80+int32(8), v129, v130)
																mBase = m.M
																v132 = m.ExcPending
																if v132 != 0 {
																	return int32(0)
																} else {
																	v135 = *(*int32)(unsafe.Add(mBase, _consts[362]))
																	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
																	*(*int32)(unsafe.Add(mBase, _consts[9])) = v136
																	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(0)
																	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
																	F_MemoryContextReset(m, v140)
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v80 + int32(48)
																		if v131 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(540539))
																			mBase = m.M
																			v202 = m.ExcPending
																			if v202 != 0 {
																				return int32(0)
																			} else {
																				v204 = *(*int32)(unsafe.Add(mBase, _consts[365]))
																				v205 = F_SPI_result_code_string(m, v204)
																				mBase = m.M
																				v206 = m.ExcPending
																				if v206 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v205
																					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																					F_errmsg_internal(m, int32(201994), v12)
																					mBase = m.M
																					v211 = m.ExcPending
																					if v211 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(491440), int32(8998), int32(148481))
																						mBase = m.M
																						v216 = m.ExcPending
																						if v216 != 0 {
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
																			v178 = m.ExcPending
																			if v178 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v12 + int32(32)
																				return v131
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
												v77 = v12 + int32(8)
												v78 = m.G0
												v80 = v78 - int32(48)
												m.G0 = v80
												if v57 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(119827), int32(0))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(489192), int32(1541), int32(277050))
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													if v77 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(119827), int32(0))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(489192), int32(1541), int32(277050))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, _consts[362]))
														if v87 == int32(0) {
															*(*int32)(unsafe.Add(mBase, _consts[365])) = int32(-4)
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v165 = m.ExcPending
															if v165 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(440278), int32(0))
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(489192), int32(1545), int32(277050))
																	mBase = m.M
																	v174 = m.ExcPending
																	if v174 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, _consts[39]))
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
															v94 = *(*int32)(unsafe.Add(mBase, _consts[362]))
															*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v92
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
															*(*int32)(unsafe.Add(mBase, _consts[9])) = v97
															v99 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v80)+20)) = v99
															v102 = v80 + int32(28)
															*(*int64)(unsafe.Add(mBase, uint32(v102))) = v99
															*(*int64)(unsafe.Add(mBase, uint32(v80)+36)) = v99
															v107 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v107
															*(*int32)(unsafe.Add(mBase, _consts[365])) = v107
															*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v99
															*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(569278163)
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
															if v118 != 0 {
																v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
																*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v119
																v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
																*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v121
															} else {
															}
															F__SPI_prepare_plan(m, v57, v80+int32(8))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
																v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
																v131 = F_SPI_cursor_open_internal(m, l3, v80+int32(8), v129, v130)
																mBase = m.M
																v132 = m.ExcPending
																if v132 != 0 {
																	return int32(0)
																} else {
																	v135 = *(*int32)(unsafe.Add(mBase, _consts[362]))
																	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
																	*(*int32)(unsafe.Add(mBase, _consts[9])) = v136
																	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(0)
																	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
																	F_MemoryContextReset(m, v140)
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v80 + int32(48)
																		if v131 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(540539))
																			mBase = m.M
																			v202 = m.ExcPending
																			if v202 != 0 {
																				return int32(0)
																			} else {
																				v204 = *(*int32)(unsafe.Add(mBase, _consts[365]))
																				v205 = F_SPI_result_code_string(m, v204)
																				mBase = m.M
																				v206 = m.ExcPending
																				if v206 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v205
																					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																					F_errmsg_internal(m, int32(201994), v12)
																					mBase = m.M
																					v211 = m.ExcPending
																					if v211 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(491440), int32(8998), int32(148481))
																						mBase = m.M
																						v216 = m.ExcPending
																						if v216 != 0 {
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
																			v178 = m.ExcPending
																			if v178 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v12 + int32(32)
																				return v131
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
											v77 = v12 + int32(8)
											v78 = m.G0
											v80 = v78 - int32(48)
											m.G0 = v80
											if v57 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(119827), int32(0))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(489192), int32(1541), int32(277050))
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												if v77 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(119827), int32(0))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(489192), int32(1541), int32(277050))
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, _consts[362]))
													if v87 == int32(0) {
														*(*int32)(unsafe.Add(mBase, _consts[365])) = int32(-4)
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(440278), int32(0))
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(489192), int32(1545), int32(277050))
																mBase = m.M
																v174 = m.ExcPending
																if v174 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, _consts[39]))
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
														v94 = *(*int32)(unsafe.Add(mBase, _consts[362]))
														*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v92
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														*(*int32)(unsafe.Add(mBase, _consts[9])) = v97
														v99 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v80)+20)) = v99
														v102 = v80 + int32(28)
														*(*int64)(unsafe.Add(mBase, uint32(v102))) = v99
														*(*int64)(unsafe.Add(mBase, uint32(v80)+36)) = v99
														v107 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v107
														*(*int32)(unsafe.Add(mBase, _consts[365])) = v107
														*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v99
														*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(569278163)
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
														if v118 != 0 {
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v119
															v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v121
														} else {
														}
														F__SPI_prepare_plan(m, v57, v80+int32(8))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
															v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
															v131 = F_SPI_cursor_open_internal(m, l3, v80+int32(8), v129, v130)
															mBase = m.M
															v132 = m.ExcPending
															if v132 != 0 {
																return int32(0)
															} else {
																v135 = *(*int32)(unsafe.Add(mBase, _consts[362]))
																v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
																*(*int32)(unsafe.Add(mBase, _consts[9])) = v136
																*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(0)
																v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
																F_MemoryContextReset(m, v140)
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v80 + int32(48)
																	if v131 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(540539))
																		mBase = m.M
																		v202 = m.ExcPending
																		if v202 != 0 {
																			return int32(0)
																		} else {
																			v204 = *(*int32)(unsafe.Add(mBase, _consts[365]))
																			v205 = F_SPI_result_code_string(m, v204)
																			mBase = m.M
																			v206 = m.ExcPending
																			if v206 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v205
																				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																				F_errmsg_internal(m, int32(201994), v12)
																				mBase = m.M
																				v211 = m.ExcPending
																				if v211 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(491440), int32(8998), int32(148481))
																					mBase = m.M
																					v216 = m.ExcPending
																					if v216 != 0 {
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
																		v178 = m.ExcPending
																		if v178 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v12 + int32(32)
																			return v131
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
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(540539))
					mBase = m.M
					v186 = m.ExcPending
					if v186 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(297399), int32(0))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(491440), int32(8974), int32(148481))
								mBase = m.M
								v198 = m.ExcPending
								if v198 != 0 {
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
				v40 = int32(4470560)
				v41 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
				*(*int32)(unsafe.Add(mBase, _consts[9])) = v44
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
						*(*int32)(unsafe.Add(mBase, _consts[9])) = v41
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
												v77 = v12 + int32(8)
												v78 = m.G0
												v80 = v78 - int32(48)
												m.G0 = v80
												if v57 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(119827), int32(0))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(489192), int32(1541), int32(277050))
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													if v77 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(119827), int32(0))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(489192), int32(1541), int32(277050))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, _consts[362]))
														if v87 == int32(0) {
															*(*int32)(unsafe.Add(mBase, _consts[365])) = int32(-4)
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v165 = m.ExcPending
															if v165 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(440278), int32(0))
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(489192), int32(1545), int32(277050))
																	mBase = m.M
																	v174 = m.ExcPending
																	if v174 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, _consts[39]))
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
															v94 = *(*int32)(unsafe.Add(mBase, _consts[362]))
															*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v92
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
															*(*int32)(unsafe.Add(mBase, _consts[9])) = v97
															v99 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v80)+20)) = v99
															v102 = v80 + int32(28)
															*(*int64)(unsafe.Add(mBase, uint32(v102))) = v99
															*(*int64)(unsafe.Add(mBase, uint32(v80)+36)) = v99
															v107 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v107
															*(*int32)(unsafe.Add(mBase, _consts[365])) = v107
															*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v99
															*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(569278163)
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
															if v118 != 0 {
																v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
																*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v119
																v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
																*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v121
															} else {
															}
															F__SPI_prepare_plan(m, v57, v80+int32(8))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
																v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
																v131 = F_SPI_cursor_open_internal(m, l3, v80+int32(8), v129, v130)
																mBase = m.M
																v132 = m.ExcPending
																if v132 != 0 {
																	return int32(0)
																} else {
																	v135 = *(*int32)(unsafe.Add(mBase, _consts[362]))
																	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
																	*(*int32)(unsafe.Add(mBase, _consts[9])) = v136
																	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(0)
																	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
																	F_MemoryContextReset(m, v140)
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v80 + int32(48)
																		if v131 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(540539))
																			mBase = m.M
																			v202 = m.ExcPending
																			if v202 != 0 {
																				return int32(0)
																			} else {
																				v204 = *(*int32)(unsafe.Add(mBase, _consts[365]))
																				v205 = F_SPI_result_code_string(m, v204)
																				mBase = m.M
																				v206 = m.ExcPending
																				if v206 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v205
																					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																					F_errmsg_internal(m, int32(201994), v12)
																					mBase = m.M
																					v211 = m.ExcPending
																					if v211 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(491440), int32(8998), int32(148481))
																						mBase = m.M
																						v216 = m.ExcPending
																						if v216 != 0 {
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
																			v178 = m.ExcPending
																			if v178 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v12 + int32(32)
																				return v131
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
											v77 = v12 + int32(8)
											v78 = m.G0
											v80 = v78 - int32(48)
											m.G0 = v80
											if v57 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(119827), int32(0))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(489192), int32(1541), int32(277050))
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												if v77 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(119827), int32(0))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(489192), int32(1541), int32(277050))
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, _consts[362]))
													if v87 == int32(0) {
														*(*int32)(unsafe.Add(mBase, _consts[365])) = int32(-4)
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(440278), int32(0))
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(489192), int32(1545), int32(277050))
																mBase = m.M
																v174 = m.ExcPending
																if v174 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, _consts[39]))
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
														v94 = *(*int32)(unsafe.Add(mBase, _consts[362]))
														*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v92
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														*(*int32)(unsafe.Add(mBase, _consts[9])) = v97
														v99 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v80)+20)) = v99
														v102 = v80 + int32(28)
														*(*int64)(unsafe.Add(mBase, uint32(v102))) = v99
														*(*int64)(unsafe.Add(mBase, uint32(v80)+36)) = v99
														v107 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v107
														*(*int32)(unsafe.Add(mBase, _consts[365])) = v107
														*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v99
														*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(569278163)
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
														if v118 != 0 {
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v119
															v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v121
														} else {
														}
														F__SPI_prepare_plan(m, v57, v80+int32(8))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
															v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
															v131 = F_SPI_cursor_open_internal(m, l3, v80+int32(8), v129, v130)
															mBase = m.M
															v132 = m.ExcPending
															if v132 != 0 {
																return int32(0)
															} else {
																v135 = *(*int32)(unsafe.Add(mBase, _consts[362]))
																v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
																*(*int32)(unsafe.Add(mBase, _consts[9])) = v136
																*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(0)
																v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
																F_MemoryContextReset(m, v140)
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v80 + int32(48)
																	if v131 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(540539))
																		mBase = m.M
																		v202 = m.ExcPending
																		if v202 != 0 {
																			return int32(0)
																		} else {
																			v204 = *(*int32)(unsafe.Add(mBase, _consts[365]))
																			v205 = F_SPI_result_code_string(m, v204)
																			mBase = m.M
																			v206 = m.ExcPending
																			if v206 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v205
																				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																				F_errmsg_internal(m, int32(201994), v12)
																				mBase = m.M
																				v211 = m.ExcPending
																				if v211 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(491440), int32(8998), int32(148481))
																					mBase = m.M
																					v216 = m.ExcPending
																					if v216 != 0 {
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
																		v178 = m.ExcPending
																		if v178 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v12 + int32(32)
																			return v131
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
											v77 = v12 + int32(8)
											v78 = m.G0
											v80 = v78 - int32(48)
											m.G0 = v80
											if v57 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(119827), int32(0))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(489192), int32(1541), int32(277050))
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												if v77 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(119827), int32(0))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(489192), int32(1541), int32(277050))
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, _consts[362]))
													if v87 == int32(0) {
														*(*int32)(unsafe.Add(mBase, _consts[365])) = int32(-4)
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(440278), int32(0))
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(489192), int32(1545), int32(277050))
																mBase = m.M
																v174 = m.ExcPending
																if v174 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, _consts[39]))
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
														v94 = *(*int32)(unsafe.Add(mBase, _consts[362]))
														*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v92
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														*(*int32)(unsafe.Add(mBase, _consts[9])) = v97
														v99 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v80)+20)) = v99
														v102 = v80 + int32(28)
														*(*int64)(unsafe.Add(mBase, uint32(v102))) = v99
														*(*int64)(unsafe.Add(mBase, uint32(v80)+36)) = v99
														v107 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v107
														*(*int32)(unsafe.Add(mBase, _consts[365])) = v107
														*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v99
														*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(569278163)
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
														if v118 != 0 {
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v119
															v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v121
														} else {
														}
														F__SPI_prepare_plan(m, v57, v80+int32(8))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
															v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
															v131 = F_SPI_cursor_open_internal(m, l3, v80+int32(8), v129, v130)
															mBase = m.M
															v132 = m.ExcPending
															if v132 != 0 {
																return int32(0)
															} else {
																v135 = *(*int32)(unsafe.Add(mBase, _consts[362]))
																v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
																*(*int32)(unsafe.Add(mBase, _consts[9])) = v136
																*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(0)
																v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
																F_MemoryContextReset(m, v140)
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v80 + int32(48)
																	if v131 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(540539))
																		mBase = m.M
																		v202 = m.ExcPending
																		if v202 != 0 {
																			return int32(0)
																		} else {
																			v204 = *(*int32)(unsafe.Add(mBase, _consts[365]))
																			v205 = F_SPI_result_code_string(m, v204)
																			mBase = m.M
																			v206 = m.ExcPending
																			if v206 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v205
																				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																				F_errmsg_internal(m, int32(201994), v12)
																				mBase = m.M
																				v211 = m.ExcPending
																				if v211 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(491440), int32(8998), int32(148481))
																					mBase = m.M
																					v216 = m.ExcPending
																					if v216 != 0 {
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
																		v178 = m.ExcPending
																		if v178 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v12 + int32(32)
																			return v131
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
										v77 = v12 + int32(8)
										v78 = m.G0
										v80 = v78 - int32(48)
										m.G0 = v80
										if v57 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(119827), int32(0))
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(489192), int32(1541), int32(277050))
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											if v77 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(119827), int32(0))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(489192), int32(1541), int32(277050))
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, _consts[362]))
												if v87 == int32(0) {
													*(*int32)(unsafe.Add(mBase, _consts[365])) = int32(-4)
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(440278), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(489192), int32(1545), int32(277050))
															mBase = m.M
															v174 = m.ExcPending
															if v174 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, _consts[39]))
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
													v94 = *(*int32)(unsafe.Add(mBase, _consts[362]))
													*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v92
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
													*(*int32)(unsafe.Add(mBase, _consts[9])) = v97
													v99 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v80)+20)) = v99
													v102 = v80 + int32(28)
													*(*int64)(unsafe.Add(mBase, uint32(v102))) = v99
													*(*int64)(unsafe.Add(mBase, uint32(v80)+36)) = v99
													v107 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v107
													*(*int32)(unsafe.Add(mBase, _consts[365])) = v107
													*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v99
													*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(569278163)
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
													if v118 != 0 {
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v119
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v121
													} else {
													}
													F__SPI_prepare_plan(m, v57, v80+int32(8))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
														v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
														v131 = F_SPI_cursor_open_internal(m, l3, v80+int32(8), v129, v130)
														mBase = m.M
														v132 = m.ExcPending
														if v132 != 0 {
															return int32(0)
														} else {
															v135 = *(*int32)(unsafe.Add(mBase, _consts[362]))
															v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
															*(*int32)(unsafe.Add(mBase, _consts[9])) = v136
															*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(0)
															v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
															F_MemoryContextReset(m, v140)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																m.G0 = v80 + int32(48)
																if v131 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(540539))
																	mBase = m.M
																	v202 = m.ExcPending
																	if v202 != 0 {
																		return int32(0)
																	} else {
																		v204 = *(*int32)(unsafe.Add(mBase, _consts[365]))
																		v205 = F_SPI_result_code_string(m, v204)
																		mBase = m.M
																		v206 = m.ExcPending
																		if v206 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v205
																			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																			F_errmsg_internal(m, int32(201994), v12)
																			mBase = m.M
																			v211 = m.ExcPending
																			if v211 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(491440), int32(8998), int32(148481))
																				mBase = m.M
																				v216 = m.ExcPending
																				if v216 != 0 {
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
																	v178 = m.ExcPending
																	if v178 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v12 + int32(32)
																		return v131
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
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(540539))
				mBase = m.M
				v186 = m.ExcPending
				if v186 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67108994))
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(297399), int32(0))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(491440), int32(8974), int32(148481))
							mBase = m.M
							v198 = m.ExcPending
							if v198 != 0 {
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
func F_existsTimeLineHistory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v3 = m.G0
	v5 = v3 - int32(1136)
	m.G0 = v5
	if l0 != int32(1) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[95])))
		if v10 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
			v20 = F_pg_snprintf(m, v5+int32(48), int32(64), int32(12756), v5+int32(16))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v31 = F_RestoreArchivedFile(m, v5+int32(112), v5+int32(48), int32(499792), int64(0), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v45 = F_AllocateFile(m, v5+int32(112), int32(227248))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						if v45 != 0 {
							v47 = F_FreeFile(m, v45)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v57 = int32(1)
								m.G0 = v5 + int32(1136)
								return v57
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, _consts[86]))
							if v51 != int32(44) {
								F_errstart_cold(m, int32(22), int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(112)
										F_errmsg(m, int32(293799), v5)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490293), int32(251), int32(13068))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
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
								v57 = int32(0)
								m.G0 = v5 + int32(1136)
								return v57
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = l0
			v40 = F_pg_snprintf(m, v5+int32(112), int32(1024), int32(12749), v5+int32(32))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v45 = F_AllocateFile(m, v5+int32(112), int32(227248))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					if v45 != 0 {
						v47 = F_FreeFile(m, v45)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v57 = int32(1)
							m.G0 = v5 + int32(1136)
							return v57
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _consts[86]))
						if v51 != int32(44) {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(112)
									F_errmsg(m, int32(293799), v5)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490293), int32(251), int32(13068))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
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
							v57 = int32(0)
							m.G0 = v5 + int32(1136)
							return v57
						}
					}
				}
			}
		}
	} else {
		v57 = int32(0)
		m.G0 = v5 + int32(1136)
		return v57
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
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[396])))
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
	v119 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
	F_sequence_close(m, v260, int32(0))
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
	F_errmsg_internal(m, int32(244718), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(484731), int32(406), int32(11992))
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
	var v50 float64
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v95 float64
	_ = v95
	var v98 float64
	_ = v98
	var v104 float64
	_ = v104
	var v114 float64
	_ = v114
	var v131 float64
	_ = v131
	var v141 float64
	_ = v141
	var v146 float64
	_ = v146
	var v153 float64
	_ = v153
	var v157 float64
	_ = v157
	var v163 float64
	_ = v163
	var v173 float64
	_ = v173
	var v175 float64
	_ = v175
	v8 = base.I64_reinterpret_f64(l0)
	v13 = base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1078159482)) <= base.Ui32(v13) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
			v175 = l0
			return v175
		} else {
			if v8 < int64(0) {
				return float64(-1)
			} else {
				if base.F64_gt(l0, float64(709.782712893384)) == int32(0) {
					v50 = base.F64_add(base.F64_mul(l0, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), l0))
					if base.F64_lt(base.F64_abs(v50), float64(2.147483648e+09)) != 0 {
						v54 = base.I32_trunc_f64_s(v50)
						v56 = v54
					} else {
						v56 = int32(-2147483648)
					}
					v57 = base.F64_convert_i32_s(v56)
					v64 = base.F64_mul(v57, float64(1.9082149292705877e-10))
					v65 = v56
					v66 = base.F64_add(l0, base.F64_mul(v57, float64(-0.6931471803691238)))
					v67 = base.F64_sub(v66, v64)
					v73 = v67
					v75 = base.F64_sub(base.F64_sub(v66, v67), v64)
					v76 = v65
					v78 = base.F64_mul(v73, float64(0.5))
					v79 = base.F64_mul(v73, v78)
					v95 = base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v98 = base.F64_sub(float64(3), base.F64_mul(v95, v78))
					v104 = base.F64_mul(v79, base.F64_div(base.F64_sub(v95, v98), base.F64_sub(float64(6), base.F64_mul(v73, v98))))
					if v76 == int32(0) {
						return base.F64_sub(v73, base.F64_sub(base.F64_mul(v73, v104), v79))
					} else {
						v114 = base.F64_sub(base.F64_sub(base.F64_mul(v73, base.F64_sub(v104, v75)), v75), v79)
						switch v76 + int32(1) {
						case 0:
							return base.F64_add(base.F64_mul(base.F64_sub(v73, v114), float64(0.5)), float64(-0.5))
						default:
							v141 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v76+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v76) {
								v146 = base.F64_add(base.F64_sub(v73, v114), float64(1))
								if v76 == int32(1024) {
									v153 = base.F64_mul(base.F64_add(v146, v146), float64(8.98846567431158e+307))
								} else {
									v153 = base.F64_mul(v146, v141)
								}
								return base.F64_add(v153, float64(-1))
							} else {
								v157 = float64(1)
								v163 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v76) << (uint(int64(52)) % 64))
								if base.Ui32(v76) <= base.Ui32(int32(19)) {
									v173 = base.F64_add(base.F64_sub(v157, v163), base.F64_sub(v73, v114))
								} else {
									v173 = base.F64_add(base.F64_sub(v73, base.F64_add(v114, v163)), v157)
								}
								v175 = base.F64_mul(v173, v141)
								return v175
							}
						case 2:
							if base.F64_lt(v73, float64(-0.25)) != 0 {
								return base.F64_mul(base.F64_sub(v114, base.F64_add(v73, float64(0.5))), float64(-2))
							} else {
								v131 = base.F64_sub(v73, v114)
								return base.F64_add(base.F64_add(v131, v131), float64(1))
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
				v175 = l0
				return v175
			} else {
				v73 = l0
				v75 = float64(0)
				v76 = int32(0)
				v78 = base.F64_mul(v73, float64(0.5))
				v79 = base.F64_mul(v73, v78)
				v95 = base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
				v98 = base.F64_sub(float64(3), base.F64_mul(v95, v78))
				v104 = base.F64_mul(v79, base.F64_div(base.F64_sub(v95, v98), base.F64_sub(float64(6), base.F64_mul(v73, v98))))
				if v76 == int32(0) {
					return base.F64_sub(v73, base.F64_sub(base.F64_mul(v73, v104), v79))
				} else {
					v114 = base.F64_sub(base.F64_sub(base.F64_mul(v73, base.F64_sub(v104, v75)), v75), v79)
					switch v76 + int32(1) {
					case 0:
						return base.F64_add(base.F64_mul(base.F64_sub(v73, v114), float64(0.5)), float64(-0.5))
					default:
						v141 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v76+int32(1023)) << (uint(int64(52)) % 64))
						if base.Ui32(int32(57)) <= base.Ui32(v76) {
							v146 = base.F64_add(base.F64_sub(v73, v114), float64(1))
							if v76 == int32(1024) {
								v153 = base.F64_mul(base.F64_add(v146, v146), float64(8.98846567431158e+307))
							} else {
								v153 = base.F64_mul(v146, v141)
							}
							return base.F64_add(v153, float64(-1))
						} else {
							v157 = float64(1)
							v163 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v76) << (uint(int64(52)) % 64))
							if base.Ui32(v76) <= base.Ui32(int32(19)) {
								v173 = base.F64_add(base.F64_sub(v157, v163), base.F64_sub(v73, v114))
							} else {
								v173 = base.F64_add(base.F64_sub(v73, base.F64_add(v114, v163)), v157)
							}
							v175 = base.F64_mul(v173, v141)
							return v175
						}
					case 2:
						if base.F64_lt(v73, float64(-0.25)) != 0 {
							return base.F64_mul(base.F64_sub(v114, base.F64_add(v73, float64(0.5))), float64(-2))
						} else {
							v131 = base.F64_sub(v73, v114)
							return base.F64_add(base.F64_add(v131, v131), float64(1))
						}
					}
				}
			}
		} else {
			if base.Ui32(int32(1072734897)) < base.Ui32(v13) {
				v50 = base.F64_add(base.F64_mul(l0, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), l0))
				if base.F64_lt(base.F64_abs(v50), float64(2.147483648e+09)) != 0 {
					v54 = base.I32_trunc_f64_s(v50)
					v56 = v54
				} else {
					v56 = int32(-2147483648)
				}
				v57 = base.F64_convert_i32_s(v56)
				v64 = base.F64_mul(v57, float64(1.9082149292705877e-10))
				v65 = v56
				v66 = base.F64_add(l0, base.F64_mul(v57, float64(-0.6931471803691238)))
			} else {
				if int64(0) <= v8 {
					v64 = float64(1.9082149292705877e-10)
					v65 = int32(1)
					v66 = base.F64_add(l0, float64(-0.6931471803691238))
				} else {
					v64 = float64(-1.9082149292705877e-10)
					v65 = int32(-1)
					v66 = base.F64_add(l0, float64(0.6931471803691238))
				}
			}
			v67 = base.F64_sub(v66, v64)
			v73 = v67
			v75 = base.F64_sub(base.F64_sub(v66, v67), v64)
			v76 = v65
			v78 = base.F64_mul(v73, float64(0.5))
			v79 = base.F64_mul(v73, v78)
			v95 = base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
			v98 = base.F64_sub(float64(3), base.F64_mul(v95, v78))
			v104 = base.F64_mul(v79, base.F64_div(base.F64_sub(v95, v98), base.F64_sub(float64(6), base.F64_mul(v73, v98))))
			if v76 == int32(0) {
				return base.F64_sub(v73, base.F64_sub(base.F64_mul(v73, v104), v79))
			} else {
				v114 = base.F64_sub(base.F64_sub(base.F64_mul(v73, base.F64_sub(v104, v75)), v75), v79)
				switch v76 + int32(1) {
				case 0:
					return base.F64_add(base.F64_mul(base.F64_sub(v73, v114), float64(0.5)), float64(-0.5))
				default:
					v141 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v76+int32(1023)) << (uint(int64(52)) % 64))
					if base.Ui32(int32(57)) <= base.Ui32(v76) {
						v146 = base.F64_add(base.F64_sub(v73, v114), float64(1))
						if v76 == int32(1024) {
							v153 = base.F64_mul(base.F64_add(v146, v146), float64(8.98846567431158e+307))
						} else {
							v153 = base.F64_mul(v146, v141)
						}
						return base.F64_add(v153, float64(-1))
					} else {
						v157 = float64(1)
						v163 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v76) << (uint(int64(52)) % 64))
						if base.Ui32(v76) <= base.Ui32(int32(19)) {
							v173 = base.F64_add(base.F64_sub(v157, v163), base.F64_sub(v73, v114))
						} else {
							v173 = base.F64_add(base.F64_sub(v73, base.F64_add(v114, v163)), v157)
						}
						v175 = base.F64_mul(v173, v141)
						return v175
					}
				case 2:
					if base.F64_lt(v73, float64(-0.25)) != 0 {
						return base.F64_mul(base.F64_sub(v114, base.F64_add(v73, float64(0.5))), float64(-2))
					} else {
						v131 = base.F64_sub(v73, v114)
						return base.F64_add(base.F64_add(v131, v131), float64(1))
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
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
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L32
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return v97
L3:
	;
	return int32(0)
L4:
	;
	v13 = int32(19320)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[275])))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v17 == int32(0) {
		v36 = v16
		v37 = v17
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v37-v36 == int32(0) {
		v97 = int32(114)
		goto L2
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	if v16 != v17 {
		v36 = v16
		v37 = v17
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v21 = v9
	v22 = v13
	goto L9
L9:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v36 = v25
		v37 = v26
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v36 = v25
	v37 = v26
	goto L6
L11:
	;
	v29 = int32(1)
	if v25 == v26 {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v42 = int32(390386)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[276])))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v46 == int32(0) {
		v65 = v45
		v66 = v46
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v66-v65 == int32(0) {
		v97 = int32(115)
		goto L2
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	if v45 != v46 {
		v65 = v45
		v66 = v46
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = v9
	v51 = v42
	goto L18
L18:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v55 == int32(0) {
		v65 = v54
		v66 = v55
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v65 = v54
	v66 = v55
	goto L15
L20:
	;
	v58 = int32(1)
	if v54 == v55 {
		v50 = v50 + v58
		v51 = v51 + v58
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v70 = int32(343826)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[277])))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v74 == int32(0) {
		v93 = v73
		v94 = v74
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v94-v93 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	if v73 != v74 {
		v93 = v73
		v94 = v74
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v78 = v9
	v79 = v70
	goto L27
L27:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v93 = v82
		v94 = v83
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v93 = v82
	v94 = v83
	goto L24
L29:
	;
	v86 = int32(1)
	if v82 == v83 {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v97 = int32(119)
	goto L2
L32:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v109
	F_errmsg(m, int32(528851), v6)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(486007), int32(491), int32(20572))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
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
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v3 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v40
L2:
	;
	v12 = v3
	v13 = v3
	goto L7
L3:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v6 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v40 = v3
	goto L1
L6:
	;
	goto L5
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14+v13<<(uint(int32(2))%32))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+10)))
	if v19 != l1 {
		v32 = v12
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v40 = v32
	goto L1
L9:
	;
	v34 = v13 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v34 < v35 {
		v12 = v32
		v13 = v34
		goto L7
	} else {
		goto L17
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 != int32(7) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v27 = F_lappend(m, v12, v21)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)))
	if v25 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if v26 != 0 {
		v32 = v12
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
	v32 = v27
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v51 int32
	_ = v51
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
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
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
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
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
	return v245
L5:
	;
	v21 = v3
	v24 = v3
	goto L8
L6:
	;
	v233 = v3
	goto L7
L7:
	;
	v238 = int32(0)
	if v233 == v238 {
		v245 = v238
		goto L4
	} else {
		goto L77
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(int32(2))%32))))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v233 = v225
	goto L7
L10:
	;
	if v202 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L11:
	;
	v135 = int32(0)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+10)))
	if v136 != 0 {
		v245 = v135
		goto L4
	} else {
		goto L48
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
		v202 = v42
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v47 = v42
	v51 = v42
	goto L19
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v47<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	goto L23
L20:
	;
	v202 = v130
	goto L10
L21:
	;
	v132 = v47 + int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v132 < v133 {
		v47 = v132
		v51 = v130
		goto L19
	} else {
		goto L47
	}
L22:
	;
	v127 = F_lappend(m, v51, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L27
	} else {
		goto L46
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
		v130 = v51
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
		v126 = v63
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v130 = v51
	goto L21
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v70 = int32(0)
	v77 = base.B2i32(v68|v69 == v70)
	if v68 == v70 {
		v116 = v77
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v116 == int32(0) {
		v130 = v51
		goto L21
	} else {
		goto L43
	}
L32:
	;
	goto L31
L33:
	;
	if v69 == int32(0) {
		v116 = v77
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v83 != v84 {
		v116 = int32(0)
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v86 = int32(1)
	if v83 <= v86 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v89 = v86
	goto L38
L37:
	;
	v89 = v83
	goto L38
L38:
	;
	v90 = int32(8)
	v95 = int32(0)
	goto L39
L39:
	;
	v103 = v95 << (uint(int32(2)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v68+v90+v103)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+(v69+v90))))
	v108 = base.B2i32(v105 == v107)
	if v107 != v105 {
		v116 = v108
		goto L32
	} else {
		goto L41
	}
L40:
	;
	v116 = v108
	goto L32
L41:
	;
	v111 = v95 + int32(1)
	if v111 != v89 {
		v95 = v111
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v123 = F_contain_volatile_functions(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	if v123 != 0 {
		v130 = v51
		goto L21
	} else {
		goto L45
	}
L45:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v126 = v125
	goto L22
L46:
	;
	v130 = v127
	goto L21
L47:
	;
	goto L20
L48:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v139 = int32(0)
	v146 = base.B2i32(v137|v138 == v139)
	if v137 == v139 {
		v185 = v146
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v185 == int32(0) {
		v245 = v135
		goto L4
	} else {
		goto L61
	}
L50:
	;
	goto L49
L51:
	;
	if v138 == int32(0) {
		v185 = v146
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v152 != v153 {
		v185 = int32(0)
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v155 = int32(1)
	if v152 <= v155 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v158 = v155
	goto L56
L55:
	;
	v158 = v152
	goto L56
L56:
	;
	v159 = int32(8)
	v164 = int32(0)
	goto L57
L57:
	;
	v172 = v164 << (uint(int32(2)) % 32)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v137+v159+v172)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172+(v138+v159))))
	v177 = base.B2i32(v174 == v176)
	if v176 != v174 {
		v185 = v177
		goto L50
	} else {
		goto L59
	}
L58:
	;
	v185 = v177
	goto L50
L59:
	;
	v180 = v164 + int32(1)
	if v180 != v158 {
		v164 = v180
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v192 = F_contain_volatile_functions(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L27
	} else {
		goto L62
	}
L62:
	;
	if v192 != 0 {
		v245 = v135
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v196 = F_lappend(m, int32(0), v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L27
	} else {
		goto L64
	}
L64:
	;
	v202 = v196
	goto L10
L65:
	;
	return int32(0)
L66:
	;
	goto L67
L67:
	;
	v210 = F_make_ands_explicit(m, v202)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L27
	} else {
		goto L70
	}
L68:
	;
	v227 = v24 + int32(1)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v227 < v228 {
		v21 = v225
		v24 = v227
		goto L8
	} else {
		goto L76
	}
L69:
	;
	v223 = F_lappend(m, v21, v210)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L27
	} else {
		goto L75
	}
L70:
	;
	if v210 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	if v214 != int32(21) {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v217 != int32(1) {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	v221 = F_list_concat(m, v21, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L27
	} else {
		goto L74
	}
L74:
	;
	v225 = v221
	goto L68
L75:
	;
	v225 = v223
	goto L68
L76:
	;
	goto L9
L77:
	;
	v241 = F_make_orclause(m, v233)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L27
	} else {
		goto L78
	}
L78:
	;
	v245 = v241
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v135 int32
	_ = v135
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
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
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
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
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
		goto L9
	} else {
		goto L10
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
	v268 = m.ExcPending
	if v268 != 0 {
		goto L15
	} else {
		goto L54
	}
L7:
	;
	m.G0 = v17 + int32(32)
	return v253
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v240
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v239
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v238
	v253 = v236
	goto L7
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v40 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v116 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v116
	v118 = F_palloc0(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L15
	} else {
		goto L26
	}
L12:
	;
	v253 = int32(-1)
	goto L7
L13:
	;
	goto L14
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v45 = F_pg_detoast_datum(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	F_get_typlenbyvalalign(m, v49, v17+int32(16), v17+int32(19), v17+int32(18))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+16)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+19)))
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+18)))
	F_deconstruct_array(m, v45, v58, v59, v60, v17+int32(28), v17+int32(24), v17+int32(20))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v72 = F_palloc0(m, v69<<(uint(int32(2))%32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if int32(0) < v74 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v82 = int32(0)
	goto L23
L21:
	;
	v106 = v74
	goto L22
L22:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v236 = v106
	v238 = v72
	v239 = v114
	v240 = v115
	goto L8
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72+v82<<(uint(int32(2))%32)))) = v49
	v97 = v82 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v97 < v98 {
		v82 = v97
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v106 = v98
	goto L22
L25:
	;
	goto L24
L26:
	;
	v121 = v116 << (uint(int32(2)) % 32)
	v122 = F_palloc0(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v124 = F_palloc0(m, v121)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	if v116 <= int32(0) {
		v236 = v116
		v238 = v124
		v239 = v118
		v240 = v122
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v135 = int32(0)
	goto L30
L30:
	;
	v148 = l0 + int32(20) + v135<<(uint(int32(3))%32)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v118))) = uint8(v149)
	v152 = v135 << (uint(int32(2)) % 32)
	v153 = v124 + v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = F_get_fn_expr_argtype(m, v154, v135)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L15
	} else {
		goto L32
	}
L31:
	;
	v236 = v228
	v238 = v124
	v239 = v118
	v240 = v122
	goto L8
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v155
	if v155 != int32(705) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122+v152))) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	if v221 == int32(0) {
		goto L6
	} else {
		goto L51
	}
L34:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v219 = v218
	goto L33
L35:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v162 = int32(0)
	if v161 == v162 {
		v207 = v162
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v207 == int32(0) {
		goto L34
	} else {
		goto L48
	}
L37:
	;
	goto L36
L38:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	if v166 == int32(0) {
		v207 = v162
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v171 = v169 - int32(11)
	if base.Ui32(int32(9)) < base.Ui32(v171) {
		v207 = v162
		goto L37
	} else {
		goto L40
	}
L40:
	;
	if int32(base.Ui32(int32(977))>>(uint(v171)%32))&int32(1) == int32(0) {
		v207 = v162
		goto L37
	} else {
		goto L41
	}
L41:
	;
	if v135 < int32(0) {
		v207 = v162
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v171<<(uint(int32(2))%32))+uint32(_consts[1213])))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v166+v186)))
	if v188 == int32(0) {
		v207 = v162
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v191 <= v135 {
		v207 = v162
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v193 = int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194+v135<<(uint(int32(2))%32))))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	switch v199 - int32(7) {
	case 0:
		v207 = v193
		goto L37
	case 1:
		goto L46
	default:
		goto L45
	}
L45:
	;
	v207 = int32(0)
	goto L37
L46:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v202 == int32(0) {
		v207 = v193
		goto L37
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = int32(25)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
	if v214 != 0 {
		v219 = int32(0)
		goto L33
	} else {
		goto L49
	}
L49:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v216 = F_cstring_to_text(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	v219 = v216
	goto L33
L51:
	;
	if v221 == int32(705) {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	v227 = v135 + int32(1)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v227 < v228 {
		v135 = v227
		goto L30
	} else {
		goto L53
	}
L53:
	;
	goto L31
L54:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v135 + int32(1)
	F_errmsg(m, int32(460273), v17)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(489211), int32(2091), int32(152872))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L15
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
