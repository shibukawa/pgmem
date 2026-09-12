package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GinPageDeletePostingItem(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
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
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v5)+4)))
	if v7 != l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(10)
	v11 = l0 + l1*v9
	v13 = v11 + int32(22)
	v15 = v11 + int32(32)
	v18 = (v7 - l1) * v9
	if v13 == v15 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v165 = v5
	goto L3
L3:
	;
	v168 = v7 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v165+l0)+4)) = uint16(v168)
	v173 = v7*int32(10) + int32(22)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v173)
	return
L4:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v165 = v163
	goto L3
L5:
	;
	goto L4
L6:
	;
	v22 = v13 + v18
	if base.Ui32(v15-v22) <= base.Ui32(int32(0)-v18<<(uint(int32(1))%32)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v29 = F___memcpy(m, v13, v15, v18)
	mBase = m.M
	goto L4
L8:
	;
	goto L9
L9:
	;
	v32 = (v13 ^ v15) & int32(3)
	if base.Ui32(v13) < base.Ui32(v15) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v134 == int32(0) {
		goto L5
	} else {
		goto L46
	}
L11:
	;
	if base.Ui32(v112) <= base.Ui32(int32(3)) {
		v133 = v111
		v134 = v112
		v135 = v113
		goto L10
	} else {
		goto L42
	}
L12:
	;
	if v32 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v32 != 0 {
		v94 = v18
		goto L25
	} else {
		goto L26
	}
L15:
	;
	v133 = v15
	v134 = v18
	v135 = v13
	goto L10
L16:
	;
	goto L17
L17:
	;
	if v13&int32(3) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v111 = v15
	v112 = v18
	v113 = v13
	goto L11
L19:
	;
	goto L20
L20:
	;
	v39 = v15
	v40 = v18
	v41 = v13
	goto L21
L21:
	;
	if v40 == int32(0) {
		goto L5
	} else {
		goto L23
	}
L22:
	;
	v111 = v48
	v112 = v50
	v113 = v52
	goto L11
L23:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v45)
	v47 = int32(1)
	v48 = v39 + v47
	v50 = v40 - v47
	v52 = v41 + v47
	if v52&int32(3) != 0 {
		v39 = v48
		v40 = v50
		v41 = v52
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	if v94 == int32(0) {
		goto L5
	} else {
		goto L38
	}
L26:
	;
	if v22&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v59 = v18
	goto L30
L28:
	;
	v74 = v18
	goto L29
L29:
	;
	if base.Ui32(v74) <= base.Ui32(int32(3)) {
		v94 = v74
		goto L25
	} else {
		goto L34
	}
L30:
	;
	if v59 == int32(0) {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v74 = v65
	goto L29
L32:
	;
	v65 = v59 - int32(1)
	v66 = v13 + v65
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v65))))
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v68)
	if v66&int32(3) != 0 {
		v59 = v65
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v81 = v74
	goto L35
L35:
	;
	v85 = v81 - int32(4)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v15+v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v13+v85))) = v88
	if base.Ui32(int32(3)) < base.Ui32(v85) {
		v81 = v85
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v94 = v85
	goto L25
L37:
	;
	goto L36
L38:
	;
	v101 = v94
	goto L39
L39:
	;
	v105 = v101 - int32(1)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v105))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13+v105))) = uint8(v108)
	if v105 != 0 {
		v101 = v105
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L5
L41:
	;
	goto L40
L42:
	;
	v118 = v111
	v119 = v112
	v120 = v113
	goto L43
L43:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v122
	v124 = int32(4)
	v125 = v118 + v124
	v127 = v120 + v124
	v129 = v119 - v124
	if base.Ui32(int32(3)) < base.Ui32(v129) {
		v118 = v125
		v119 = v129
		v120 = v127
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v133 = v125
	v134 = v129
	v135 = v127
	goto L10
L45:
	;
	goto L44
L46:
	;
	v140 = v133
	v141 = v134
	v142 = v135
	goto L47
L47:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v144)
	v146 = int32(1)
	v151 = v141 - v146
	if v151 != 0 {
		v140 = v140 + v146
		v141 = v151
		v142 = v142 + v146
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L5
L49:
	;
	goto L48
}
func F_ginBuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 float64
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(4487040)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[34])))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if int32(0) < v23 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[37])))
	v94 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	if base.Ui32(v94<<(uint(int32(10))%32)) <= base.Ui32(v92) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v33))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2+v33<<(uint(int32(2))%32))))
	v47 = int32(4487040)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[51])))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[33])))
	v54 = v33 + int32(1)
	v56 = v54 & int32(65535)
	v61 = F_ginExtractEntries(m, v52, v56, v46, v42, v15+int32(8), v15+int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v48
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_ginInsertBAEntries(m, l5+int32(5728), l1, v56, v61, v65, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[41])))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[41]))) = base.F64_add(v69, base.F64_convert_i32_s(v70))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[51])))
	F_MemoryContextReset(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v54 < v78 {
		v33 = v54
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	v99 = l5 + int32(5728)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[52])))
	v103 = m.G0
	v104 = int32(16)
	v105 = v103 - v104
	m.G0 = v105
	*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[53]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[54]))) = v102
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[55]))) = uint8(base.B2i32(v112 == int32(4101768)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[56]))) = int32(790)
	m.G0 = v105 + v104
	goto L14
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v18
	m.G0 = v15 + int32(16)
	return
L14:
	;
	v129 = F_ginGetBAEntry(m, v99, v15+int32(4), v15+int32(12), v15+int32(7), v15+int32(8))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if v129 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v137 = v129
	goto L19
L17:
	;
	goto L18
L18:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_consts[34])))
	F_MemoryContextReset(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L28
	}
L19:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v146 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v151 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+7)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_ginEntryInsert(m, l5, v149, v150, v151, v137, v152, l5+int32(5688))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v163 = F_ginGetBAEntry(m, v99, v15+int32(4), v15+int32(12), v15+int32(7), v15+int32(8))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	if v163 != 0 {
		v137 = v163
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L20
L28:
	;
	F_ginInitBA(m, v99)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	goto L13
}
func F_ginCompareEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	if l3 != l5 {
		if l3 < l5 {
			v11 = int32(-1)
		} else {
			v11 = int32(1)
		}
		return v11
	} else {
		if l3 != 0 {
			v29 = int32(0)
			return v29
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32)+l0)+uint32(_consts[43])))
			v25 = F_FunctionCall2Coll(m, l1*int32(28)+l0+int32(112), v24, l2, l4)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = v25
				return v29
			}
		}
	}
}
func F_ginDataFillRoot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v14
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v16)
	v19 = l1 + int32(32)
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v21 = l1 + v20
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v23 = int32(10)
	v24 = v22 * v23
	v25 = v19 + v24
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = base.I32_rotr(l2, v12)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v29
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+8)) = uint16(v31)
	v33 = int32(1)
	v34 = v22 + v33
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v34)
	v36 = int32(42)
	v37 = v24 + v36
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v37)
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v39)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v44 = l1 + v43
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	v47 = v45 * v23
	v48 = v19 + v47
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = base.I32_rotr(l4, v12)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v52
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+8)) = uint16(v54)
	v57 = v45 + v33
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)) = uint16(v57)
	v60 = v47 + v36
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v60)
	return
}
func F_ginExtractEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
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
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
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
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v291
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	v18 = F_palloc(m, int32(4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
	v38 = l1 - int32(1)
	v41 = l0 + v38<<(uint(int32(2))%32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[46])))
	v47 = F_FunctionCall3Coll(m, l1*int32(28)+l0+int32(1008), v44, l2, l4, v13+int32(12))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L9
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	v25 = F_palloc(m, int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v25
	v28 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v28)
	v291 = v18
	goto L1
L8:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v66 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if int32(0) < v49 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	v56 = F_palloc(m, int32(4))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(0)
	v61 = F_palloc(m, int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v61
	v64 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v64)
	v291 = v56
	goto L1
L16:
	;
	v69 = F_palloc0(m, v49)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	v73 = v49
	goto L18
L18:
	;
	if int32(2) <= v73 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v69
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v73 = v72
	goto L18
L20:
	;
	v78 = F_palloc(m, v73<<(uint(int32(3))%32))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	v260 = v73
	goto L22
L22:
	;
	v261 = F_palloc0(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L56
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if int32(0) < v80 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v85 = int32(0)
	goto L27
L25:
	;
	v117 = v80
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0 + v38*int32(28) + int32(140)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[46])))
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v127)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v126
	F_qsort_arg(m, v78, v117, int32(8), int32(75), v13)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L30
	}
L27:
	;
	v96 = v78 + v85<<(uint(int32(3))%32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v47+v85<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v85))))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)) = uint8(v104)
	v107 = v85 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v107 < v108 {
		v85 = v107
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v117 = v108
	goto L26
L29:
	;
	goto L28
L30:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
	if v135 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	F_pfree(m, v78)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L55
	}
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v138 <= int32(0) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v170)
	v172 = int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if int32(2) <= v173 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v142 = v127
	goto L36
L36:
	;
	v156 = v78 + v142<<(uint(int32(3))%32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*int32)(unsafe.Add(mBase, uint32(v47+v142<<(uint(int32(2))%32)))) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v159+v142))) = uint8(v161)
	v164 = v142 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v164 < v165 {
		v142 = v164
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L31
L38:
	;
	goto L37
L39:
	;
	v177 = v172
	v184 = int32(1)
	goto L42
L40:
	;
	v226 = v172
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v226
	goto L31
L42:
	;
	v189 = v78 + v184<<(uint(int32(3))%32)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+4)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189-int32(4)))))
	if v193 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v226 = v221
	goto L41
L44:
	;
	v223 = v184 + int32(1)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v223 < v224 {
		v177 = v221
		v184 = v223
		goto L42
	} else {
		goto L54
	}
L45:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	*(*int32)(unsafe.Add(mBase, uint32(v47+v177<<(uint(int32(2))%32)))) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v215+v177))) = uint8(v217)
	v221 = v177 + int32(1)
	goto L44
L46:
	;
	v208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v208)
	v221 = v177
	goto L44
L47:
	;
	if v190&int32(1) != 0 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v190&int32(1) != 0 {
		goto L45
	} else {
		goto L51
	}
L50:
	;
	goto L45
L51:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v189-int32(8))))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v206 = F_FunctionCall2Coll(m, v200, v201, v204, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	if v206 != 0 {
		goto L45
	} else {
		goto L53
	}
L53:
	;
	goto L46
L54:
	;
	goto L43
L55:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v260 = v249
	goto L22
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v261
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v264 <= int32(0) {
		v291 = v47
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v269 = int32(0)
	goto L58
L58:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v269))))
	*(*uint8)(unsafe.Add(mBase, uint32(v278+v269))) = uint8(v282)
	v285 = v269 + int32(1)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v285 < v286 {
		v269 = v285
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v291 = v47
	goto L1
L60:
	;
	goto L59
}
func F_ginFinishSplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
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
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
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
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
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
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
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
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
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
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v481 int32
	_ = v481
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
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
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v19 = l1
	v29 = int32(0)
	goto L2
L1:
	;
	m.G0 = v16 - int32(-64)
	return
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	F_LockBuffer(m, v32, int32(2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	F_LockBuffer(m, v561, int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L145
	}
L4:
	;
	return
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v36 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v108 = m.T0[v107].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v104, v105, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L24
	}
L7:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+16)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v54)+6)))
	if v57&int32(64) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(v36^int32(-1))<<(uint(int32(2))%32))))
	v54 = v46
	goto L7
L9:
	;
	goto L10
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v54 = v48 + v36<<(uint(int32(13))%32) + int32(-8192)
	goto L7
L11:
	;
	v62 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	v86 = v36
	goto L13
L13:
	;
	if v86 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	if v62 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v65 + int32(4)
	F_errmsg_internal(m, int32(674912), v14+int32(-16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_ginFinishSplit(m, l0, v31, int32(0), l3)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	F_errfinish(m, int32(496259), int32(783), int32(102171))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v86 = v85
	goto L13
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90+(v86^int32(-1))<<(uint(int32(2))%32))))
	v104 = v96
	goto L6
L22:
	;
	goto L23
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v104 = v98 + v86<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L24:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)) = uint16(v108)
	if v108 != 0 {
		v481 = v31
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v492 = m.T0[v491].(func(*base.Module, int32, int32) int32)(m, l0, v490)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L118
	}
L26:
	;
	v117 = v104
	goto L28
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+20)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v469
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+8)) = uint16(v468)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v171
	v481 = v171
	goto L25
L28:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v117+v125)))
	if v127 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L115
	}
L30:
	;
	goto L29
L31:
	;
	F_LockBuffer(m, v124, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v365 = F_ginStepRight(m, v124, v363, int32(2))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L94
	}
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+20))
	if v134 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v143 = v133
	goto L38
L36:
	;
	v161 = v133
	goto L37
L37:
	;
	v166 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+8)) = uint16(v166)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v171 = F_palloc(m, int32(24))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L42
	}
L38:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	F_ReleaseBuffer(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L40
	}
L39:
	;
	v161 = v151
	goto L37
L40:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
	if v152 != 0 {
		v143 = v151
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v178 = v168
	v183 = v169
	goto L43
L43:
	;
	F_LockBuffer(m, v178, int32(2))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	if v178 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)))
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+v206)+6)))
	if v209&int32(2) != 0 {
		goto L30
	} else {
		goto L50
	}
L47:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192+(v178^int32(-1))<<(uint(int32(2))%32))))
	v206 = v198
	goto L46
L48:
	;
	goto L49
L49:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v206 = v200 + v178<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	if v209&int32(64) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+20)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v183
	v217 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+8)) = uint16(v217)
	v221 = F_errstart(m, int32(14), v217)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v247 = m.T0[v246].(func(*base.Module, int32, int32) int32)(m, l0, v206)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L61
	}
L54:
	;
	if v221 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+48))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v224 + int32(4)
	F_errmsg_internal(m, int32(674912), v14+int32(-48))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v241 = int32(0)
	F_ginFinishSplit(m, l0, v171, v241, v241)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	F_errfinish(m, int32(496259), int32(783), int32(102171))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	goto L53
L61:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v252 = m.T0[v251].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v206, v249, int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v361 = F_ReadBuffer(m, v360, v247)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L93
	}
L63:
	;
	if v252 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v261 = v178
	v262 = v206
	goto L67
L65:
	;
	goto L66
L66:
	;
	if v183 != int32(-1) {
		v464 = v178
		v468 = v252
		v469 = v183
		goto L27
	} else {
		goto L92
	}
L67:
	;
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+16)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v262+v269)))
	if v271 == int32(-1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_LockBuffer(m, v261, int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v283 = F_ginStepRight(m, v261, v281, int32(2))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L76
	}
L72:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v261 == v277 {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	F_ReleaseBuffer(m, v261)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	goto L62
L75:
	;
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v302)+16)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v302)+6)))
	if v305&int32(64) != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	if v283 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288+(v283^int32(-1))<<(uint(int32(2))%32))))
	v302 = v294
	goto L75
L78:
	;
	goto L79
L79:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v302 = v296 + v283<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+20)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v271
	v311 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+8)) = uint16(v311)
	v315 = F_errstart(m, int32(14), v311)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v341 = m.T0[v340].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v302, v338, int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L90
	}
L83:
	;
	if v315 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+48))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v318 + int32(4)
	F_errmsg_internal(m, int32(674912), v16)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v333 = int32(0)
	F_ginFinishSplit(m, l0, v171, v333, v333)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	F_errfinish(m, int32(496259), int32(783), int32(102171))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	goto L82
L90:
	;
	if v341 == int32(0) {
		v261 = v283
		v262 = v302
		goto L67
	} else {
		goto L91
	}
L91:
	;
	v464 = v283
	v468 = v341
	v469 = v271
	goto L27
L92:
	;
	goto L62
L93:
	;
	v178 = v361
	v183 = v247
	goto L43
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v365
	if v365 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v388 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v371+(v365^int32(-1))<<(uint(int32(6))%32))+16))
	v386 = v377
	goto L95
L97:
	;
	goto L98
L98:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v379+v365<<(uint(int32(6))%32)+int32(-64))+16))
	v386 = v385
	goto L95
L99:
	;
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v406)+16)))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407+v406)+6)))
	if v409&int32(64) != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v392+(v388^int32(-1))<<(uint(int32(2))%32))))
	v406 = v398
	goto L99
L101:
	;
	goto L102
L102:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v406 = v400 + v388<<(uint(int32(13))%32) + int32(-8192)
	goto L99
L103:
	;
	v414 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v441 = m.T0[v440].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v406, v438, v439)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L113
	}
L106:
	;
	if v414 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+48))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v417 + int32(4)
	F_errmsg_internal(m, int32(674912), v14+int32(-32))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	F_ginFinishSplit(m, l0, v31, int32(0), l3)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	F_errfinish(m, int32(496259), int32(783), int32(102171))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	goto L105
L113:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)) = uint16(v441)
	if v441 == int32(0) {
		v117 = v406
		goto L28
	} else {
		goto L114
	}
L114:
	;
	v481 = v31
	goto L25
L115:
	;
	F_errmsg_internal(m, int32(319919), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(496259), int32(256), int32(120004))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
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
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v494 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v512)+16)))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v513+v512)))
	v516 = F_ginPlaceToPage(m, l0, v481, v492, v515, v494, l3)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L123
	}
L120:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v498+(v494^int32(-1))<<(uint(int32(2))%32))))
	v512 = v504
	goto L119
L121:
	;
	goto L122
L122:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v512 = v506 + v494<<(uint(int32(13))%32) + int32(-8192)
	goto L119
L123:
	;
	F_pfree(m, v492)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	if (l2|v29)&int32(1) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	F_LockBuffer(m, v523, int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L4
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	if l2 != 0 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	goto L127
L129:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	F_ReleaseBuffer(m, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	if v516 == int32(0) {
		v19 = v481
		v29 = int32(1)
		goto L2
	} else {
		goto L144
	}
L132:
	;
	F_pfree(m, v19)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	if v516 == int32(0) {
		v19 = v481
		v29 = int32(1)
		goto L2
	} else {
		goto L134
	}
L134:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	F_LockBuffer(m, v535, int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v543 = v481
	goto L136
L136:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v543)+20))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	if v553 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	goto L1
L138:
	;
	F_ReleaseBuffer(m, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	F_pfree(m, v543)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	if v552 != 0 {
		v543 = v552
		goto L136
	} else {
		goto L143
	}
L143:
	;
	goto L137
L144:
	;
	goto L3
L145:
	;
	goto L1
}
func F_ginFlushBuildState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[57])))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v19 = l0 + int32(5728)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[52])))
	v23 = m.G0
	v24 = int32(16)
	v25 = v23 - v24
	m.G0 = v25
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[53]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[54]))) = v22
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[55]))) = uint8(base.B2i32(v32 == int32(4101768)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[56]))) = int32(790)
	m.G0 = v25 + v24
	goto L1
L1:
	;
	v44 = base.I32_div_u_s(int32(178956970), v16<<(uint(int32(1))%32))
	v53 = F_ginGetBAEntry(m, v19, v14+int32(18), v14+int32(28), v14+int32(27), v14+int32(20))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	if v53 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v64 = v53
	goto L7
L5:
	;
	goto L6
L6:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[34])))
	F_MemoryContextReset(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L27
	}
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v68 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+18)))
	v76 = v17 - int32(80) + v69<<(uint(int32(4))%32) + v73*int32(100)
	v79 = int32(0)
	v81 = v68
	goto L12
L10:
	;
	goto L11
L11:
	;
	v136 = F_ginGetBAEntry(m, v19, v14+int32(18), v14+int32(28), v14+int32(27), v14+int32(20))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L25
	}
L12:
	;
	v89 = v81 - v79
	if base.Ui32(v44) < base.Ui32(v89) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v91 = v44
	goto L16
L15:
	;
	v91 = v89
	goto L16
L16:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v93 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+18)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76)+72)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+82)))
	v106 = F__gin_build_tuple(m, v96, v97, v98, v99, v100, v64+v79*int32(6), v91, v14+int32(12))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[58])))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_tuplesort_putgintuple(m, v108, v106, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_pfree(m, v106)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v114 = v79 + v91
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if base.Ui32(v114) < base.Ui32(v115) {
		v79 = v114
		v81 = v115
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	if v136 != 0 {
		v64 = v136
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L8
L27:
	;
	F_ginInitBA(m, v19)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	m.G0 = v14 + int32(32)
	return
}
func F_ginPostingListDecodeAllSegments(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
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
	var v91 int64
	_ = v91
	var v96 int32
	_ = v96
	var v102 int64
	_ = v102
	var v107 int32
	_ = v107
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v124 int64
	_ = v124
	var v129 int32
	_ = v129
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v146 int64
	_ = v146
	var v151 int64
	_ = v151
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	v4 = int32(0)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v13 = int32(1)
	v16 = v12<<(uint(v13)%32) | v13
	v19 = F_palloc(m, v16*int32(6))
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
	v23 = l0 + l1
	if base.Ui32(l0) < base.Ui32(v23) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = l0
	v28 = v16
	v29 = v4
	v30 = v19
	goto L6
L4:
	;
	v201 = v4
	v202 = v19
	goto L5
L5:
	;
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L6:
	;
	if v29 < v28 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v201 = v182
	v202 = v183
	goto L5
L8:
	;
	v47 = v43 + v29*int32(6)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v48
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)) = uint16(v50)
	v53 = v29 + int32(1)
	v55 = v25 + int32(8)
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
	v57 = v55 + v56
	if base.Ui32(v55) < base.Ui32(v57) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v43 = v30
	v44 = v28
	goto L8
L10:
	;
	goto L11
L11:
	;
	v41 = F_repalloc(m, v30, v28*int32(12))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v43 = v41
	v44 = v28 << (uint(int32(1)) % 32)
	goto L8
L13:
	;
	v59 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
	v60 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v25)+2)))
	v63 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v71 = v55
	v72 = v53
	v73 = v43
	v74 = v44
	v78 = v59 | (v60<<(uint(int64(11))%64) | v63<<(uint(int64(27))%64))
	goto L16
L14:
	;
	v179 = v56
	v182 = v53
	v183 = v43
	v184 = v44
	goto L15
L15:
	;
	v195 = v25 + (v179+int32(1))&int32(131070) + int32(8)
	if base.Ui32(v195) < base.Ui32(v23) {
		v25 = v195
		v28 = v184
		v29 = v182
		v30 = v183
		goto L6
	} else {
		goto L30
	}
L16:
	;
	if v74 <= v72 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
	v179 = v177
	v182 = v175
	v183 = v86
	v184 = v87
	goto L15
L18:
	;
	v82 = F_repalloc(m, v73, v74*int32(12))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v86 = v73
	v87 = v74
	goto L20
L20:
	;
	v88 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71))))
	v91 = base.I64_extend_i32_u(v88 & int32(127))
	if int32(0) <= v88 {
		v158 = v91
		v159 = v71 + int32(1)
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v86 = v82
	v87 = v74 << (uint(int32(1)) % 32)
	goto L20
L22:
	;
	v162 = v86 + v72*int32(6)
	v163 = v158 + v78
	v165 = int64(base.Ui64(v163) >> (uint(int64(11)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+2)) = uint16(v165)
	v168 = int64(base.Ui64(v163) >> (uint(int64(27)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v162))) = uint16(v168)
	v172 = base.I32_wrap_i64(v163) & int32(2047)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+4)) = uint16(v172)
	v175 = v72 + int32(1)
	if base.Ui32(v159) < base.Ui32(v57) {
		v71 = v159
		v72 = v175
		v73 = v86
		v74 = v87
		v78 = v163
		goto L16
	} else {
		goto L29
	}
L23:
	;
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+1)))
	v102 = base.I64_extend_i32_u(v96<<(uint(int32(7))%32))&int64(16256) | v91
	if int32(0) <= v96 {
		v158 = v102
		v159 = v71 + int32(2)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+2)))
	v113 = base.I64_extend_i32_u(v107<<(uint(int32(14))%32))&int64(2080768) | v102
	if int32(0) <= v107 {
		v158 = v113
		v159 = v71 + int32(3)
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v118 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+3)))
	v124 = base.I64_extend_i32_u(v118<<(uint(int32(21))%32))&int64(266338304) | v113
	if int32(0) <= v118 {
		v158 = v124
		v159 = v71 + int32(4)
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v129 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+4)))
	v135 = base.I64_extend_i32_u(v129)<<(uint(int64(28))%64)&int64(34091302912) | v124
	if int32(0) <= v129 {
		v158 = v135
		v159 = v71 + int32(5)
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v140 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+5)))
	v146 = base.I64_extend_i32_u(v140)<<(uint(int64(35))%64)&int64(4363686772736) | v135
	if int32(0) <= v140 {
		v158 = v146
		v159 = v71 + int32(6)
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v71)+6)))
	v158 = v151<<(uint(int64(42))%64) | v146
	v159 = v71 + int32(7)
	goto L22
L29:
	;
	goto L17
L30:
	;
	goto L7
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v201
	goto L33
L32:
	;
	goto L33
L33:
	;
	return v202
}
func F_ginScanToDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
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
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
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
	var v197 int32
	_ = v197
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
	var v204 int32
	_ = v204
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
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int64
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v384 int64
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	v5 = l4
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v49 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[39])))
	v46 = F_ReadBufferExtended(m, v42, v43, l1, v43, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L12
	}
L5:
	;
	v24 = F_palloc0(m, int32(20))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v32 = v20
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[39])))
	v37 = F_ReadBufferExtended(m, v33, v34, l1, v34, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
	v32 = v24
	goto L7
L10:
	;
	F_LockBuffer(m, v37, int32(2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v48 = v32
	v49 = v37
	goto L1
L12:
	;
	v48 = l3
	v49 = v46
	goto L1
L13:
	;
	v156 = v67 + v147&int32(65535)
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)))
	if v157&int32(2) != 0 {
		goto L32
	} else {
		goto L33
	}
L14:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v67)+6)))
	if v70&int32(2) != 0 {
		v147 = v68
		goto L13
	} else {
		goto L18
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(v49^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L14
L16:
	;
	goto L17
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v67 = v61 + v49<<(uint(int32(13))%32) + int32(-8192)
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = l1
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v75 = v67 + v74
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
	if v76 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v83 = int32(1)
	goto L22
L20:
	;
	v121 = v74
	v123 = v75
	goto L21
L21:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v128 != int32(-1) {
		v147 = v121
		goto L13
	} else {
		goto L26
	}
L22:
	;
	v94 = v83 & int32(65535)
	v97 = v67 + int32(22) + v94*int32(10)
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+2)))
	v104 = F_ginScanToDelete(m, l0, v98<<(uint(int32(16))%32)|v101, int32(0), v48, v94)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L8
	} else {
		goto L24
	}
L23:
	;
	v121 = v111
	v123 = v112
	goto L21
L24:
	;
	v108 = v83 - v104 + int32(1)
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v112 = v67 + v111
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+4)))
	if base.Ui32(v108&int32(65535)) <= base.Ui32(v113) {
		v83 = v108
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	if v132 == int32(0) {
		v147 = v121
		goto L13
	} else {
		goto L27
	}
L27:
	;
	F_UnlockReleaseBuffer(m, v132)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+12)) = int32(0)
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v147 = v140
	goto L13
L29:
	;
	if l2 != 0 {
		goto L104
	} else {
		goto L105
	}
L30:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v448 != 0 {
		goto L100
	} else {
		goto L101
	}
L31:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v167 == int32(0) {
		goto L30
	} else {
		goto L41
	}
L32:
	;
	if v157&int32(128) != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
	if v166 != 0 {
		goto L30
	} else {
		goto L40
	}
L35:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	if v162 == int32(32) {
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
	if v165 != 0 {
		goto L30
	} else {
		goto L39
	}
L38:
	;
	goto L30
L39:
	;
	goto L31
L40:
	;
	goto L31
L41:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v170 == int32(-1) {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	if v167 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[39])))
	v198 = F_ReadBufferExtended(m, v194, v195, v191, v195, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L8
	} else {
		goto L47
	}
L44:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v176+(v167^int32(-1))<<(uint(int32(6))%32))+16))
	v191 = v182
	goto L43
L45:
	;
	goto L46
L46:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184+v167<<(uint(int32(6))%32)+int32(-64))+16))
	v191 = v190
	goto L43
L47:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v201 = int32(0)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[39])))
	v204 = F_ReadBufferExtended(m, v200, v201, l1, v201, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v207 = int32(0)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[39])))
	v210 = F_ReadBufferExtended(m, v206, v207, v193, v207, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v204 < int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230)+16)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231+v230)))
	F_PredicateLockPageCombine(m, v212, l1, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L8
	} else {
		goto L54
	}
L51:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v216+(v204^int32(-1))<<(uint(int32(2))%32))))
	v230 = v222
	goto L50
L52:
	;
	goto L53
L53:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v230 = v224 + v204<<(uint(int32(13))%32) + int32(-8192)
	goto L50
L54:
	;
	v236 = int32(4481700)
	v238 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v238 + int32(1)
	if v198 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v260+v259))) = v233
	if v210 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245+(v198^int32(-1))<<(uint(int32(2))%32))))
	v259 = v251
	goto L55
L57:
	;
	goto L58
L58:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v259 = v253 + v198<<(uint(int32(13))%32) + int32(-8192)
	goto L55
L59:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+16)))
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280+v283)+4)))
	if v285 != v5 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266+(v210^int32(-1))<<(uint(int32(2))%32))))
	v280 = v272
	goto L59
L61:
	;
	goto L62
L62:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v280 = v274 + v210<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	if v204 < int32(0) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v287 = int32(10)
	v289 = v280 + v5*v287
	v297 = F_memmove(m, v289+int32(22), v289+int32(32), (v285-v5)*v287)
	mBase = m.M
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+16)))
	v300 = v298
	goto L66
L65:
	;
	v300 = v283
	goto L66
L66:
	;
	v303 = v285 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v300+v280)+4)) = uint16(v303)
	v308 = v285*int32(10) + int32(22)
	*(*uint16)(unsafe.Add(mBase, uint32(v280)+12)) = uint16(v308)
	goto L63
L67:
	;
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v327)+16)))
	v329 = v328 + v327
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+6)))
	v332 = v330 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+6)) = uint16(v332)
	v334 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L8
	} else {
		goto L71
	}
L68:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v313+(v204^int32(-1))<<(uint(int32(2))%32))))
	v327 = v319
	goto L67
L69:
	;
	goto L70
L70:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v327 = v321 + v204<<(uint(int32(13))%32) + int32(-8192)
	goto L67
L71:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v327)+20)) = uint32(v334)
	F_MarkBufferDirty(m, v210)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	F_MarkBufferDirty(m, v198)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	F_MarkBufferDirty(m, v204)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+48))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+118)))
	if v345 != int32(112) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_ReleaseBuffer(m, v210)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L8
	} else {
		goto L92
	}
L76:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if v349 <= int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v343)+32))
	if v352 != 0 {
		goto L75
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L8
	} else {
		goto L82
	}
L80:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v343)+40))
	if v353 != 0 {
		goto L75
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v356 = int32(0)
	F_XLogRegisterBuffer(m, v356, v204, v356)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	F_XLogRegisterBuffer(m, int32(1), v210, int32(8))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	F_XLogRegisterBuffer(m, int32(2), v198, int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)) = uint16(v5)
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v327)+16)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v327+v369)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v371
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v327)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v373
	F_XLogRegisterData(m, v16+int32(4), int32(12))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v382 = F_XLogInsert(m, int32(13), int32(80))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	v384 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v327))) = base.I64_rotr(v382, v384)
	v387 = base.I32_wrap_i64(v382)
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v387
	v391 = base.I32_wrap_i64(int64(base.Ui64(v382) >> (uint(v384) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = v391
	if v198 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+4)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = v391
	goto L75
L89:
	;
	v396 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v396+(v198^int32(-1))<<(uint(int32(2))%32))))
	v410 = v402
	goto L88
L90:
	;
	goto L91
L91:
	;
	v404 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v410 = v404 + v198<<(uint(int32(13))%32) + int32(-8192)
	goto L88
L92:
	;
	F_ReleaseBuffer(m, v198)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_ReleaseBuffer(m, v204)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	v423 = int32(4481700)
	v425 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v426 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v425 - v426
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v430)+24)) = v431 + v426
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v435)+28)) = v436 + v426
	if l2 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	F_LockBuffer(m, v49, int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L8
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	F_ReleaseBuffer(m, v49)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L8
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v454 = v426
	goto L29
L100:
	;
	F_UnlockReleaseBuffer(m, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L8
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v49
	v454 = int32(0)
	goto L29
L103:
	;
	goto L102
L104:
	;
	F_ReleaseBuffer(m, v49)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L8
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	m.G0 = v16 + int32(16)
	return v454
L107:
	;
	goto L106
}
func F_gin_clean_pending_list(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v5 = m.G0
	v7 = v5 - int32(5744)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_index_open(m, v9, int32(3))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
		if v17 == int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[30]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+316))
			v25 = base.B2i32(v23 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(v25)
			v27 = v25
		} else {
			v27 = int32(0)
		}
		if v27 == int32(0) {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
			if v31 != int32(105) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return int32(0)
					} else {
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v150 + int32(4)
						F_errmsg(m, int32(29130), v7+int32(16))
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(489722), int32(1049), int32(73637))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
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
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+84))
				if v34 != int32(2742) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int32(0)
						} else {
							v150 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v150 + int32(4)
							F_errmsg(m, int32(29130), v7+int32(16))
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(489722), int32(1049), int32(73637))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
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
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+118)))
					if v37 == int32(116) {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
						if v40 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(142934), int32(0))
									mBase = m.M
									v174 = m.ExcPending
									if v174 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(489722), int32(1059), int32(73637))
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
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
							v45 = *(*int32)(unsafe.Add(mBase, _consts[31]))
							v46 = F_object_ownercheck(m, int32(1259), v9, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								if v46 == int32(0) {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
									F_aclcheck_error(m, int32(2), int32(20), v52+int32(4))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v59 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[32]))) = v59
										*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[33]))) = v59
										*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[34]))) = v59
										*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[35]))) = v59
										*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[36]))) = v59
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+192))
										v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+18)))
										if v76 == int32(1) {
											F_initGinState(m, v7+int32(28), v11)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v85 = int32(1)
												F_ginInsertCleanup(m, v7+int32(28), v85, v85, v85, v7+int32(5704))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													F_relation_close(m, v11, int32(3))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
														v117 = F_Int64GetDatum(m, v116)
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 + int32(5744)
															return v117
														}
													}
												}
											}
										} else {
											v94 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												if v94 == int32(0) {
													F_relation_close(m, v11, int32(3))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
														v117 = F_Int64GetDatum(m, v116)
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 + int32(5744)
															return v117
														}
													}
												} else {
													F_errcode(m, int32(325))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v7))) = v101 + int32(4)
														F_errmsg(m, int32(433333), v7)
														mBase = m.M
														v107 = m.ExcPending
														if v107 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(489722), int32(1086), int32(73637))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v11, int32(3))
																mBase = m.M
																v115 = m.ExcPending
																if v115 != 0 {
																	return int32(0)
																} else {
																	v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
																	v117 = F_Int64GetDatum(m, v116)
																	mBase = m.M
																	v118 = m.ExcPending
																	if v118 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v7 + int32(5744)
																		return v117
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
									v59 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[32]))) = v59
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[33]))) = v59
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[34]))) = v59
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[35]))) = v59
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[36]))) = v59
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+192))
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+18)))
									if v76 == int32(1) {
										F_initGinState(m, v7+int32(28), v11)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											v85 = int32(1)
											F_ginInsertCleanup(m, v7+int32(28), v85, v85, v85, v7+int32(5704))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v11, int32(3))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
													v117 = F_Int64GetDatum(m, v116)
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(5744)
														return v117
													}
												}
											}
										}
									} else {
										v94 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											if v94 == int32(0) {
												F_relation_close(m, v11, int32(3))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
													v117 = F_Int64GetDatum(m, v116)
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(5744)
														return v117
													}
												}
											} else {
												F_errcode(m, int32(325))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return int32(0)
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v7))) = v101 + int32(4)
													F_errmsg(m, int32(433333), v7)
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(489722), int32(1086), int32(73637))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v11, int32(3))
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return int32(0)
															} else {
																v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
																v117 = F_Int64GetDatum(m, v116)
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v7 + int32(5744)
																	return v117
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
						v45 = *(*int32)(unsafe.Add(mBase, _consts[31]))
						v46 = F_object_ownercheck(m, int32(1259), v9, v45)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v46 == int32(0) {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
								F_aclcheck_error(m, int32(2), int32(20), v52+int32(4))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v59 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[32]))) = v59
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[33]))) = v59
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[34]))) = v59
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[35]))) = v59
									*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[36]))) = v59
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+192))
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+18)))
									if v76 == int32(1) {
										F_initGinState(m, v7+int32(28), v11)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											v85 = int32(1)
											F_ginInsertCleanup(m, v7+int32(28), v85, v85, v85, v7+int32(5704))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v11, int32(3))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
													v117 = F_Int64GetDatum(m, v116)
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(5744)
														return v117
													}
												}
											}
										}
									} else {
										v94 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											if v94 == int32(0) {
												F_relation_close(m, v11, int32(3))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
													v117 = F_Int64GetDatum(m, v116)
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(5744)
														return v117
													}
												}
											} else {
												F_errcode(m, int32(325))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return int32(0)
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v7))) = v101 + int32(4)
													F_errmsg(m, int32(433333), v7)
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(489722), int32(1086), int32(73637))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v11, int32(3))
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return int32(0)
															} else {
																v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
																v117 = F_Int64GetDatum(m, v116)
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v7 + int32(5744)
																	return v117
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
								v59 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[32]))) = v59
								*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[33]))) = v59
								*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[34]))) = v59
								*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[35]))) = v59
								*(*int64)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[36]))) = v59
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+192))
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+18)))
								if v76 == int32(1) {
									F_initGinState(m, v7+int32(28), v11)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										v85 = int32(1)
										F_ginInsertCleanup(m, v7+int32(28), v85, v85, v85, v7+int32(5704))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											F_relation_close(m, v11, int32(3))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
												v117 = F_Int64GetDatum(m, v116)
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 + int32(5744)
													return v117
												}
											}
										}
									}
								} else {
									v94 = F_errstart(m, int32(14), int32(0))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										if v94 == int32(0) {
											F_relation_close(m, v11, int32(3))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
												v117 = F_Int64GetDatum(m, v116)
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 + int32(5744)
													return v117
												}
											}
										} else {
											F_errcode(m, int32(325))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = v101 + int32(4)
												F_errmsg(m, int32(433333), v7)
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(489722), int32(1086), int32(73637))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v11, int32(3))
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[37]))))
															v117 = F_Int64GetDatum(m, v116)
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return int32(0)
															} else {
																m.G0 = v7 + int32(5744)
																return v117
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(127299), int32(0))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(554555), int32(0))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(489722), int32(1041), int32(73637))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
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
func F_gin_compare_jsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(4)
			v22 = int32(1)
			v23 = v18 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			v30 = v13 + v22
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v33 = v31 & v22
			if v31 == v22 {
				v36 = int32(4)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
				if v38&int32(254) == int32(2) {
					v47 = v36
				} else {
					v47 = base.B2i32(v38 == int32(18)) << (uint(v36) % 32)
				}
				if v38 == int32(1) {
					v50 = v36
				} else {
					v50 = v47
				}
				v61 = v50
			} else {
				v51 = int32(1)
				if v33 != 0 {
					v61 = int32(base.Ui32(v31)>>(uint(v51)%32)) - v51
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v33 != 0 {
				v62 = v30
			} else {
				v62 = v13 + v20
			}
			if v24&v22 != 0 {
				v63 = v23
			} else {
				v63 = v18 + v20
			}
			if v24 == int32(1) {
				v66 = int32(4)
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				if v68&int32(254) == int32(2) {
					v77 = v66
				} else {
					v77 = base.B2i32(v68 == int32(18)) << (uint(v66) % 32)
				}
				if v68 == int32(1) {
					v80 = v66
				} else {
					v80 = v77
				}
				v93 = v80
			} else {
				v81 = int32(1)
				if v24&v81 != 0 {
					v93 = int32(base.Ui32(v24)>>(uint(v81)%32)) - v81
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v93 = int32(base.Ui32(v87)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v95 = F_varstr_cmp(m, v62, v61, v63, v93, int32(950))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return int32(0)
			} else {
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v97 != v13 {
					F_pfree(m, v13)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v101 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								return v95
							}
						} else {
							return v95
						}
					}
				} else {
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v101 != v18 {
						F_pfree(m, v18)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							return v95
						}
					} else {
						return v95
					}
				}
			}
		}
	}
}
func F_gin_consistent_hstore(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	switch v14 - int32(7) {
	case 0:
		goto L6
	default:
		goto L4
	case 2, 3:
		goto L3
	case 4:
		goto L5
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v77
L2:
	;
	v77 = int32(1)
	goto L1
L3:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v72)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v35 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v35)
	if v12 <= v35 {
		goto L2
	} else {
		goto L14
	}
L6:
	;
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v17)
	v20 = int32(0)
	if v12 <= v20 {
		v77 = v17
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v23 = v20
	goto L8
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v13))))
	if v30 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v77 = int32(0)
	goto L1
L10:
	;
	v32 = v23 + int32(1)
	if v12 != v32 {
		v23 = v32
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v77 = v17
	goto L1
L14:
	;
	v40 = v35
	goto L15
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v13))))
	if v47 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v77 = int32(0)
	goto L1
L17:
	;
	v48 = int32(1)
	v50 = v40 + v48
	if v12 != v50 {
		v40 = v50
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	v77 = v48
	goto L1
L21:
	;
	return int32(0)
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	F_errmsg_internal(m, int32(479325), v9)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(493451), int32(207), int32(361793))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_jsonb_path(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v21 = v19 & int32(268435455)
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(48)
	return v120
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	v120 = v2
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = F_palloc(m, v21<<(uint(int32(3))%32))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	v34 = F_JsonbIteratorInit(m, v14+int32(4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v34
	v41 = v11 + int32(16)
	v44 = v28
	v45 = v21 << (uint(int32(1)) % 32)
	v46 = v2
	goto L10
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v46
	v120 = v44
	goto L3
L10:
	;
	v54 = F_JsonbIteratorNext(m, v11+int32(44), v11+int32(24), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L34
	}
L12:
	;
	goto L11
L13:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	F_pfree(m, v41)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L30
	}
L14:
	;
	F_JsonbHashScalarValue(m, v11+int32(24), v41)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	if v54 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	switch v54 {
	case 0:
		goto L9
	default:
		goto L12
	case 2, 3:
		goto L14
	case 4, 6:
		goto L19
	case 5, 7:
		goto L13
	}
L17:
	;
	goto L18
L18:
	;
	F_JsonbHashScalarValue(m, v11+int32(24), v41)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v59 = F_palloc(m, int32(8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v61
	v41 = v59
	goto L10
L21:
	;
	goto L10
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v46 < v45 {
		v84 = v44
		v85 = v45
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84+v46<<(uint(int32(2))%32)))) = v72
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v91
	v44 = v84
	v45 = v85
	v46 = v46 + int32(1)
	goto L10
L24:
	;
	if v45 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = F_repalloc(m, v44, v45<<(uint(int32(3))%32))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v82 = F_palloc(m, int32(32))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v84 = v76
	v85 = v45 << (uint(int32(1)) % 32)
	goto L23
L29:
	;
	v84 = v82
	v85 = int32(8)
	goto L23
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v98 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v99
	v41 = v95
	goto L10
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(0)
	v41 = v95
	goto L10
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v54
	F_errmsg_internal(m, int32(484631), v11)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(493476), int32(1170), int32(319896))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_query_bit(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(6941), int32(2661))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2112)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(-32768)
						v71 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v71)
						m.G0 = v11 + int32(16)
						return v19
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(479325), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493464), int32(97), int32(15698))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
func F_gin_extract_query_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2288)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						v70 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v70
							v73 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v73)
							m.G0 = v11 + int32(16)
							return v19
						}
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(479325), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493464), int32(97), int32(15698))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
func F_gin_extract_query_interval(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(0), int32(6938), int32(2540))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				v35 = F_pg_detoast_datum(m, v15)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(6247)
					v38 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v38)
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v35
					*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
					v43 = F_palloc(m, int32(4))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v43))) = v24
						v48 = v14 & int32(65535)
						switch v48 - int32(1) {
						case 0, 1:
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
							v73 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v73)
							m.G0 = v11 + int32(16)
							return v19
						case 2:
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v35
							m.G0 = v11 + int32(16)
							return v19
						case 3, 4:
							v51 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v51)
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v35
							m.G0 = v11 + int32(16)
							return v19
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v48
								F_errmsg_internal(m, int32(479325), v11)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493464), int32(97), int32(15698))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
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
func F_gin_extract_query_text(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(6940), int32(2120))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_timetz(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(0), int32(6937), int32(1443))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_uuid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(3400)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						v70 = F_palloc0(m, int32(16))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v70
							v73 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v73)
							m.G0 = v11 + int32(16)
							return v19
						}
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(479325), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493464), int32(97), int32(15698))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
func F_gin_extract_tsquery_5args(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v2 <= int32(6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(120946), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(489400), int32(319), int32(154621))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v20 = F_gin_extract_tsquery(m, l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return v20
		}
	}
}
func F_gin_extract_tsquery_oldsig(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_gin_extract_tsquery(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_gin_extract_tsvector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v14
	if v14 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v65 != v9 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v62 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v20 = F_palloc(m, v14<<(uint(int32(2))%32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v22 <= int32(0) {
		v62 = v20
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v27 = v9 + int32(8)
	v30 = int32(0)
	v31 = v22
	v33 = v27
	goto L9
L9:
	;
	v35 = int32(2)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v49 = F_cstring_to_text_with_len(m, v27+v31<<(uint(v35)%32)+int32(base.Ui32(v41)>>(uint(int32(12))%32)), int32(base.Ui32(v41)>>(uint(int32(1))%32))&int32(2047))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v62 = v20
	goto L3
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v30<<(uint(v35)%32)))) = v49
	v55 = v30 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v55 < v56 {
		v30 = v55
		v31 = v56
		v33 = v33 + int32(4)
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_pfree(m, v9)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	return v62
L16:
	;
	goto L15
}
func F_gin_triconsistent_jsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	switch v12 - int32(7) {
	case 0, 4:
		goto L3
	default:
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v91
L2:
	;
	if base.Ui32((v12-int32(9))&int32(65535)) <= base.Ui32(int32(1)) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	v15 = int32(0)
	v16 = int32(2)
	if v10 <= v15 {
		v91 = v16
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = v15
	goto L5
L5:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v11))))
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v91 = int32(0)
	goto L1
L7:
	;
	v27 = v19 + int32(1)
	if v10 != v27 {
		v19 = v27
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	v91 = v16
	goto L1
L11:
	;
	if v10 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if base.Ui32((v12-int32(15))&int32(65535)) <= base.Ui32(int32(1)) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v91 = int32(0)
	goto L1
L15:
	;
	goto L16
L16:
	;
	v40 = int32(0)
	goto L17
L17:
	;
	v45 = int32(2)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v11))))
	if base.Ui32((v47-int32(1))&int32(255)) < base.Ui32(v45) {
		v91 = v45
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v91 = int32(0)
	goto L1
L19:
	;
	v55 = v40 + int32(1)
	if v55 != v10 {
		v40 = v55
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = F_execute_jsp_gin_node(m, v84, v11)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L26
	} else {
		goto L30
	}
L22:
	;
	if int32(0) < v10 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v91 = int32(2)
	goto L1
L26:
	;
	return int32(0)
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
	F_errmsg_internal(m, int32(479325), v8)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(493476), int32(1072), int32(499457))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	if v85 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v89 = int32(2)
	goto L33
L32:
	;
	v89 = v85
	goto L33
L33:
	;
	v91 = v89
	goto L1
}
func F_gin_triconsistent_jsonb_path(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	if v12 == int32(7) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v63
L2:
	;
	v15 = int32(0)
	v16 = int32(2)
	if v10 <= v15 {
		v63 = v16
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if base.Ui32((v12-int32(15))&int32(65535)) <= base.Ui32(int32(1)) {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v19 = v15
	goto L6
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v11))))
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v63 = int32(0)
	goto L1
L8:
	;
	v27 = v19 + int32(1)
	if v10 != v27 {
		v19 = v27
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	v63 = v16
	goto L1
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = F_execute_jsp_gin_node(m, v56, v11)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L17
	} else {
		goto L21
	}
L13:
	;
	if int32(0) < v10 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v63 = int32(2)
	goto L1
L17:
	;
	return int32(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
	F_errmsg_internal(m, int32(479325), v8)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(493476), int32(1314), int32(319841))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	if v57 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v61 = int32(2)
	goto L24
L23:
	;
	v61 = v57
	goto L24
L24:
	;
	v63 = v61
	goto L1
}
func F_initGinState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
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
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v168 int32
	_ = v168
	var v170 int64
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v219 int32
	_ = v219
	var v221 int64
	_ = v221
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int32
	_ = v239
	var v241 int64
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v275 int32
	_ = v275
	var v277 int64
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int32
	_ = v311
	var v313 int64
	_ = v313
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int64
	_ = v352
	var v354 int64
	_ = v354
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	v3 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(32)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v33 = F__emscripten_memset_bulkmem(m, l0+int32(4), base.I32_extend8_s(v3), int32(5672))
	mBase = m.M
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(base.B2i32(v35 == int32(1)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if int32(0) < v40 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L13
	} else {
		goto L63
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L13
	} else {
		goto L58
	}
L4:
	;
	v64 = l0 + int32(140)
	v69 = v40
	v71 = v3
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v25 + int32(32)
	return
L7:
	;
	v94 = v27 + int32(20) + v69<<(uint(int32(4))%32) + v71*int32(100)
	v96 = v71 << (uint(int32(2)) % 32)
	v97 = l0 + int32(12) + v96
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v98 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v134 = int32(1)
	v135 = v71 + v134
	v136 = base.I32_extend16_s(v135)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+6)))
	v148 = int32(4)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v138+v140*(v136-v134)<<(uint(int32(2))%32)+v148-v148)))
	goto L19
L10:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v101
	goto L9
L11:
	;
	goto L12
L12:
	;
	v104 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v104
	v108 = int32(0)
	F_TupleDescInitEntry(m, v104, int32(1), v108, int32(21), int32(-1), v108)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v94)+68))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v94)+76))
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+80)))
	F_TupleDescInitEntry(m, v114, int32(2), int32(0), v117, v118, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v94)+96))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	*(*int32)(unsafe.Add(mBase, uint32(v122+v125<<(uint(int32(4))%32)+int32(200))+16)) = v124
	goto L17
L17:
	;
	goto L9
L18:
	;
	v205 = v71 * int32(28)
	v206 = l0 + int32(1036) + v205
	v208 = F_index_getprocinfo(m, l1, v136, int32(2))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L13
	} else {
		goto L28
	}
L19:
	;
	if v152 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v155 = v64 + v71*int32(28)
	v157 = F_index_getprocinfo(m, l1, v136, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v94)+68))
	v177 = F_lookup_type_cache(m, v175, int32(64))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L25
	}
L23:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v163 = v155 + int32(16)
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v157)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = v164
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	*(*int64)(unsafe.Add(mBase, uint32(v155))) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v157)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+24)) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v157)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v155)+8)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v155)+20)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = int32(0)
	goto L24
L24:
	;
	goto L18
L25:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)+108))
	if v179 == int32(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v184 = v64 + v71*int32(28)
	v186 = v177 + int32(104)
	v188 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v191 = v184 + int32(16)
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v186)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v192
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v186)))
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v186)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v184)+24)) = v196
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v186)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v184)+8)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v184)+20)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = int32(0)
	goto L27
L27:
	;
	goto L18
L28:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v214 = v206 + int32(16)
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v208)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v214))) = v215
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v208)))
	*(*int64)(unsafe.Add(mBase, uint32(v206))) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v208)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+24)) = v219
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v208)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v206)+8)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v206)+20)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = int32(0)
	goto L29
L29:
	;
	v226 = v205 + (l0 + int32(1932))
	v228 = F_index_getprocinfo(m, l1, v136, int32(3))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v234 = v226 + int32(16)
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v228)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v234))) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	*(*int64)(unsafe.Add(mBase, uint32(v226))) = v237
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v228)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+24)) = v239
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v228)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v226)+8)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v226)+20)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(0)
	goto L31
L31:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248)+6)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v247+v249*(v136-int32(1))<<(uint(int32(2))%32)+int32(24)-int32(4))))
	goto L32
L32:
	;
	if v261 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v262 = v205 + (l0 + int32(3724))
	v264 = F_index_getprocinfo(m, l1, v136, int32(6))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+6)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v283+v285*(v136-int32(1))<<(uint(int32(2))%32)+int32(16)-int32(4))))
	goto L38
L36:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v270 = v262 + int32(16)
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v264)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v270))) = v271
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v264)))
	*(*int64)(unsafe.Add(mBase, uint32(v262))) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v264)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+24)) = v275
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v264)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v262)+8)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v262)+20)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = int32(0)
	goto L37
L37:
	;
	goto L35
L38:
	;
	if v297 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v298 = v205 + (l0 + int32(2828))
	v300 = F_index_getprocinfo(m, l1, v136, int32(4))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L13
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v205+(l0+int32(2832)))))
	if v319 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v306 = v298 + int32(16)
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = v307
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
	*(*int64)(unsafe.Add(mBase, uint32(v298))) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+24)) = v311
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v298)+8)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v298)+20)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = int32(0)
	goto L43
L43:
	;
	goto L41
L44:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v205+(l0+int32(3728)))))
	if v323 == int32(0) {
		goto L2
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+6)))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v328+v330*(v136-int32(1))<<(uint(int32(2))%32)+int32(20)-int32(4))))
	goto L48
L47:
	;
	goto L46
L48:
	;
	if v342 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v343 = v205 + (l0 + int32(4620))
	v345 = F_index_getprocinfo(m, l1, v136, int32(5))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L13
	} else {
		goto L52
	}
L50:
	;
	v365 = int32(0)
	goto L51
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v71+(l0+int32(5516))))) = uint8(v365)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v368+v96)))
	if v370 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v351 = v343 + int32(16)
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v345)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v352
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v345)))
	*(*int64)(unsafe.Add(mBase, uint32(v343))) = v354
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v345)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v343)+24)) = v356
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v345)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v343)+8)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v343)+20)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = int32(0)
	goto L53
L53:
	;
	v365 = int32(1)
	goto L51
L54:
	;
	v372 = v370
	goto L56
L55:
	;
	v372 = int32(100)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96+(l0+int32(5548))))) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v135 < v374 {
		v69 = v374
		v71 = v135
		goto L7
	} else {
		goto L57
	}
L57:
	;
	goto L8
L58:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v94)+68))
	v409 = F_format_type_be(m, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v409
	F_errmsg(m, int32(188669), v25)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(494271), int32(153), int32(351613))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L13
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
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v135
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = int64(25769803780)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v424 + int32(4)
	F_errmsg_internal(m, int32(678122), v25+int32(16))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(494271), int32(190), int32(351613))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
