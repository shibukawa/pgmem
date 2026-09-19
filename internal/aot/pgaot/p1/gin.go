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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v5)+4)))
	if v7 != l1 {
		v11 = (v7 - l1) * int32(10)
		if v11 != 0 {
			v14 = l0 + l1*int32(10)
			base.MemoryCopy(m, v14+int32(22), v14+int32(32), v11)
		} else {
		}
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v24 = v21
	} else {
		v24 = v5
	}
	v27 = v7 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24+l0)+4)) = uint16(v27)
	v32 = v7*int32(10) + int32(22)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v32)
	return
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
	v17 = int32(_a_F_ginBuildCallback_0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[0]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[0])) = v20
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
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[2])))
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[3]))
	if base.Ui32(v94<<(uint(int32(10))%32)) <= base.Ui32(v92) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v33))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2+v33<<(uint(int32(2))%32))))
	v47 = int32(_a_F_ginBuildCallback_0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[0]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[0])) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[5])))
	v54 = v33 + int32(1)
	v56 = v54 & int32(_a_F_ginBuildCallback_1)
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
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[0])) = v48
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_ginInsertBAEntries(m, l5+int32(_a_F_ginBuildCallback_2), l1, v56, v61, v65, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[6])))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[6]))) = base.F64_add(v69, base.F64_convert_i32_s(v70))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[4])))
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
	v99 = l5 + int32(_a_F_ginBuildCallback_2)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[7])))
	v103 = m.G0
	v104 = int32(16)
	v105 = v103 - v104
	m.G0 = v105
	*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[8]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[9]))) = v102
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[10]))) = uint8(base.B2i32(v112 == int32(_a_F_ginBuildCallback_3)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[11]))) = int32(790)
	m.G0 = v105 + v104
	goto L14
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[0])) = v18
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
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[1])))
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
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[12]))
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
	F_ginEntryInsert(m, l5, v149, v150, v151, v137, v152, l5+int32(_a_F_ginBuildCallback_4))
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
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32)+l0)+uint32(_c_F_ginCompareEntries[0])))
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_ginExtractEntries[0])))
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
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_ginExtractEntries[0])))
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
	var v30 int32
	_ = v30
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
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
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
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
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
	var v163 int32
	_ = v163
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
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
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
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
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
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
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
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
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
	v30 = l2
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
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v109 = m.T0[v108].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v105, v106, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(v36^int32(-1))<<(uint(int32(2))%32))))
	v54 = v46
	goto L7
L9:
	;
	goto L10
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
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
	v87 = v36
	goto L13
L13:
	;
	if v87 < int32(0) {
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
	F_errmsg_internal(m, int32(_a_F_ginFinishSplit_0), v14+int32(-16))
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
	F_errfinish(m, int32(_a_F_ginFinishSplit_1), int32(783), int32(_a_F_ginFinishSplit_2))
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
	v87 = v85
	goto L13
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91+(v87^int32(-1))<<(uint(int32(2))%32))))
	v105 = v97
	goto L6
L22:
	;
	goto L23
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
	v105 = v99 + v87<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L24:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)) = uint16(v109)
	if v109 != 0 {
		v482 = v31
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v493 = m.T0[v492].(func(*base.Module, int32, int32) int32)(m, l0, v491)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L118
	}
L26:
	;
	v118 = v105
	goto L28
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v465
	*(*uint16)(unsafe.Add(mBase, uint32(v172)+8)) = uint16(v471)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v172
	v482 = v172
	goto L25
L28:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118+v126)))
	if v128 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L115
	}
L30:
	;
	goto L29
L31:
	;
	F_LockBuffer(m, v125, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v366 = F_ginStepRight(m, v125, v364, int32(2))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L94
	}
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v135 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v145 = v134
	goto L38
L36:
	;
	v163 = v134
	goto L37
L37:
	;
	v167 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v163)+8)) = uint16(v167)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v172 = F_palloc(m, int32(24))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L42
	}
L38:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	F_ReleaseBuffer(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L40
	}
L39:
	;
	v163 = v152
	goto L37
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v145)+20))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	if v153 != 0 {
		v145 = v152
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v179 = v170
	v181 = v169
	goto L43
L43:
	;
	F_LockBuffer(m, v181, int32(2))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	if v181 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+16)))
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208+v207)+6)))
	if v210&int32(2) != 0 {
		goto L30
	} else {
		goto L50
	}
L47:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v193+(v181^int32(-1))<<(uint(int32(2))%32))))
	v207 = v199
	goto L46
L48:
	;
	goto L49
L49:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
	v207 = v201 + v181<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	if v210&int32(64) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v179
	v218 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v172)+8)) = uint16(v218)
	v222 = F_errstart(m, int32(14), v218)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v248 = m.T0[v247].(func(*base.Module, int32, int32) int32)(m, l0, v207)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L61
	}
L54:
	;
	if v222 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+48))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v225 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginFinishSplit_0), v14+int32(-48))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v242 = int32(0)
	F_ginFinishSplit(m, l0, v172, v242, v242)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	F_errfinish(m, int32(_a_F_ginFinishSplit_1), int32(783), int32(_a_F_ginFinishSplit_2))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
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
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v253 = m.T0[v252].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v207, v250, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v362 = F_ReadBuffer(m, v361, v248)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
	} else {
		goto L93
	}
L63:
	;
	if v253 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v263 = v207
	v264 = v181
	goto L67
L65:
	;
	goto L66
L66:
	;
	if v179 != int32(-1) {
		v465 = v179
		v467 = v181
		v471 = v253
		goto L27
	} else {
		goto L92
	}
L67:
	;
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263)+16)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v263+v270)))
	if v272 == int32(-1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_LockBuffer(m, v264, int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v284 = F_ginStepRight(m, v264, v282, int32(2))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L76
	}
L72:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v264 == v278 {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	F_ReleaseBuffer(m, v264)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	goto L62
L75:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v303)+16)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304+v303)+6)))
	if v306&int32(64) != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	if v284 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v289+(v284^int32(-1))<<(uint(int32(2))%32))))
	v303 = v295
	goto L75
L78:
	;
	goto L79
L79:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
	v303 = v297 + v284<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v272
	v312 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v172)+8)) = uint16(v312)
	v316 = F_errstart(m, int32(14), v312)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v342 = m.T0[v341].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v303, v339, int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L90
	}
L83:
	;
	if v316 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+48))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v319 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginFinishSplit_0), v16)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v334 = int32(0)
	F_ginFinishSplit(m, l0, v172, v334, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ginFinishSplit_1), int32(783), int32(_a_F_ginFinishSplit_2))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
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
	if v342 == int32(0) {
		v263 = v303
		v264 = v284
		goto L67
	} else {
		goto L91
	}
L91:
	;
	v465 = v272
	v467 = v284
	v471 = v342
	goto L27
L92:
	;
	goto L62
L93:
	;
	v179 = v248
	v181 = v362
	goto L43
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v366
	if v366 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v387
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v389 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[2]))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v372+(v366^int32(-1))<<(uint(int32(6))%32))+16))
	v387 = v378
	goto L95
L97:
	;
	goto L98
L98:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[3]))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v380+v366<<(uint(int32(6))%32)+int32(-64))+16))
	v387 = v386
	goto L95
L99:
	;
	v408 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v407)+16)))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+v407)+6)))
	if v410&int32(64) != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v393+(v389^int32(-1))<<(uint(int32(2))%32))))
	v407 = v399
	goto L99
L101:
	;
	goto L102
L102:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
	v407 = v401 + v389<<(uint(int32(13))%32) + int32(-8192)
	goto L99
L103:
	;
	v415 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v442 = m.T0[v441].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v407, v439, v440)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L113
	}
L106:
	;
	if v415 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+48))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v418 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginFinishSplit_0), v14+int32(-32))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
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
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	F_errfinish(m, int32(_a_F_ginFinishSplit_1), int32(783), int32(_a_F_ginFinishSplit_2))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
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
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)) = uint16(v442)
	if v442 == int32(0) {
		v118 = v407
		goto L28
	} else {
		goto L114
	}
L114:
	;
	v482 = v31
	goto L25
L115:
	;
	F_errmsg_internal(m, int32(_a_F_ginFinishSplit_3), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ginFinishSplit_1), int32(256), int32(_a_F_ginFinishSplit_4))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
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
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v495 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v513)+16)))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v514+v513)))
	v517 = F_ginPlaceToPage(m, l0, v482, v493, v516, v495, l3)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L123
	}
L120:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v499+(v495^int32(-1))<<(uint(int32(2))%32))))
	v513 = v505
	goto L119
L121:
	;
	goto L122
L122:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
	v513 = v507 + v495<<(uint(int32(13))%32) + int32(-8192)
	goto L119
L123:
	;
	F_pfree(m, v493)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	if v30&int32(1) != 0 {
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
	if v517 == int32(0) {
		v19 = v482
		v30 = int32(1)
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
	if v517 == int32(0) {
		v19 = v482
		v30 = int32(1)
		goto L2
	} else {
		goto L134
	}
L134:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
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
	v543 = v482
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v112 int32
	_ = v112
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[0])))
	v18 = l0 + int32(_a_F_ginFlushBuildState_0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[1])))
	v22 = m.G0
	v23 = int32(16)
	v24 = v22 - v23
	m.G0 = v24
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[2]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[3]))) = v21
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[4]))) = uint8(base.B2i32(v31 == int32(_a_F_ginFlushBuildState_1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[5]))) = int32(790)
	m.G0 = v24 + v23
	goto L1
L1:
	;
	v43 = base.I32_div_u_s(int32(178956970), v16<<(uint(int32(1))%32))
	v52 = F_ginGetBAEntry(m, v18, v13+int32(18), v13+int32(28), v13+int32(27), v13+int32(20))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	if v52 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v61 = v52
	goto L7
L5:
	;
	goto L6
L6:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[6])))
	F_MemoryContextReset(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L27
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v64 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+18)))
	v74 = v15 + v65<<(uint(int32(4))%32) + v69*int32(100) - int32(80)
	v77 = int32(0)
	v79 = v64
	goto L12
L10:
	;
	goto L11
L11:
	;
	v132 = F_ginGetBAEntry(m, v18, v13+int32(18), v13+int32(28), v13+int32(27), v13+int32(20))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L25
	}
L12:
	;
	v86 = v79 - v77
	if base.Ui32(v43) < base.Ui32(v86) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v88 = v43
	goto L16
L15:
	;
	v88 = v86
	goto L16
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_ginFlushBuildState[7]))
	if v90 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+18)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74)+72)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+82)))
	v103 = F__gin_build_tuple(m, v93, v94, v95, v96, v97, v61+v77*int32(6), v88, v13+int32(12))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[8])))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_tuplesort_putgintuple(m, v105, v103, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_pfree(m, v103)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v111 = v77 + v88
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if base.Ui32(v111) < base.Ui32(v112) {
		v77 = v111
		v79 = v112
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	if v132 != 0 {
		v61 = v132
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L8
L27:
	;
	F_ginInitBA(m, v18)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	m.G0 = v13 + int32(32)
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
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
	if int32(0) < l1 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = l0
	v29 = v16
	v31 = v4
	v32 = v19
	goto L6
L4:
	;
	v204 = v4
	v205 = v19
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
	if v31 < v29 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v204 = v188
	v205 = v189
	goto L5
L8:
	;
	v48 = v45 + v31*int32(6)
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)) = uint16(v49)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v51
	v54 = v31 + int32(1)
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
	if v55 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v44 = v29
	v45 = v32
	goto L8
L10:
	;
	goto L11
L11:
	;
	v42 = F_repalloc(m, v32, v29*int32(12))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v44 = v29 << (uint(int32(1)) % 32)
	v45 = v42
	goto L8
L13:
	;
	v57 = v26 + int32(8)
	v59 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	v60 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v26)+2)))
	v63 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
	v69 = v44
	v71 = v57
	v73 = v54
	v74 = v45
	v78 = v59 | (v60<<(uint(int64(11))%64) | v63<<(uint(int64(27))%64))
	goto L16
L14:
	;
	v184 = v44
	v188 = v54
	v189 = v45
	v194 = int32(0)
	goto L15
L15:
	;
	v197 = v26 + v194 + int32(8)
	if base.Ui32(v197) < base.Ui32(l0+l1) {
		v26 = v197
		v29 = v184
		v31 = v188
		v32 = v189
		goto L6
	} else {
		goto L30
	}
L16:
	;
	if v69 <= v73 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
	v184 = v86
	v188 = v175
	v189 = v87
	v194 = (v177 + int32(1)) & int32(_a_F_ginPostingListDecodeAllSegments_0)
	goto L15
L18:
	;
	v82 = F_repalloc(m, v74, v69*int32(12))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v86 = v69
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
	v86 = v69 << (uint(int32(1)) % 32)
	v87 = v82
	goto L20
L22:
	;
	v162 = v87 + v73*int32(6)
	v163 = v158 + v78
	v165 = int64(base.Ui64(v163) >> (uint(int64(11)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+2)) = uint16(v165)
	v168 = int64(base.Ui64(v163) >> (uint(int64(27)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v162))) = uint16(v168)
	v172 = base.I32_wrap_i64(v163) & int32(2047)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+4)) = uint16(v172)
	v175 = v73 + int32(1)
	if base.Ui32(v159) < base.Ui32(v55+v57) {
		v69 = v86
		v71 = v159
		v73 = v175
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
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v204
	goto L33
L32:
	;
	goto L33
L33:
	;
	return v205
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int64
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v388 int64
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
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
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginScanToDelete[0])))
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginScanToDelete[0])))
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
	v154 = v67 + v148
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+6)))
	if v155&int32(2) != 0 {
		goto L32
	} else {
		goto L33
	}
L14:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v67)+6)))
	if v70&int32(2) != 0 {
		v148 = v68
		goto L13
	} else {
		goto L18
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[1]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(v49^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L14
L16:
	;
	goto L17
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[2]))
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
	v122 = v74
	v124 = v75
	goto L21
L21:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v128 != int32(-1) {
		v148 = v122
		goto L13
	} else {
		goto L26
	}
L22:
	;
	v94 = v83 & int32(_a_F_ginScanToDelete_0)
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
	v122 = v111
	v124 = v112
	goto L21
L24:
	;
	v108 = v83 - v104 + int32(1)
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v112 = v67 + v111
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+4)))
	if base.Ui32(v108&int32(_a_F_ginScanToDelete_0)) <= base.Ui32(v113) {
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
		v148 = v122
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
	v148 = v140
	goto L13
L29:
	;
	if l2 != 0 {
		goto L107
	} else {
		goto L108
	}
L30:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v452 != 0 {
		goto L103
	} else {
		goto L104
	}
L31:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v169 == int32(0) {
		goto L30
	} else {
		goto L41
	}
L32:
	;
	if v155&int32(128) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+4)))
	if v168 != 0 {
		goto L30
	} else {
		goto L40
	}
L35:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+4)))
	if v162 == int32(0) {
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	if v165 != int32(32) {
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
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v172 == int32(-1) {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	if v169 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v197 = int32(0)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginScanToDelete[0])))
	v200 = F_ReadBufferExtended(m, v196, v197, v193, v197, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L47
	}
L44:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[3]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v178+(v169^int32(-1))<<(uint(int32(6))%32))+16))
	v193 = v184
	goto L43
L45:
	;
	goto L46
L46:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[4]))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186+v169<<(uint(int32(6))%32)+int32(-64))+16))
	v193 = v192
	goto L43
L47:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v203 = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginScanToDelete[0])))
	v206 = F_ReadBufferExtended(m, v202, v203, l1, v203, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v209 = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginScanToDelete[0])))
	v212 = F_ReadBufferExtended(m, v208, v209, v195, v209, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v206 < int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232)+16)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233+v232)))
	F_PredicateLockPageSplit(m, v214, l1, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L54
	}
L51:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[1]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v218+(v206^int32(-1))<<(uint(int32(2))%32))))
	v232 = v224
	goto L50
L52:
	;
	goto L53
L53:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[2]))
	v232 = v226 + v206<<(uint(int32(13))%32) + int32(-8192)
	goto L50
L54:
	;
	v238 = int32(_a_F_ginScanToDelete_1)
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[5])) = v240 + int32(1)
	if v200 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v262+v261))) = v235
	if v212 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[1]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v247+(v200^int32(-1))<<(uint(int32(2))%32))))
	v261 = v253
	goto L55
L57:
	;
	goto L58
L58:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[2]))
	v261 = v255 + v200<<(uint(int32(13))%32) + int32(-8192)
	goto L55
L59:
	;
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282)+16)))
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282+v285)+4)))
	if v287 != v5 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[1]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v268+(v212^int32(-1))<<(uint(int32(2))%32))))
	v282 = v274
	goto L59
L61:
	;
	goto L62
L62:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[2]))
	v282 = v276 + v212<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	if v206 < int32(0) {
		goto L71
	} else {
		goto L72
	}
L64:
	;
	v291 = (v287 - v5) * int32(10)
	if v291 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v304 = v285
	goto L66
L66:
	;
	v307 = v287 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v304+v282)+4)) = uint16(v307)
	v312 = v287*int32(10) + int32(22)
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+12)) = uint16(v312)
	goto L63
L67:
	;
	v294 = v282 + v5*int32(10)
	base.MemoryCopy(m, v294+int32(22), v294+int32(32), v291)
	goto L69
L68:
	;
	goto L69
L69:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282)+16)))
	v304 = v301
	goto L66
L70:
	;
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331)+16)))
	v333 = v332 + v331
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v333)+6)))
	v336 = v334 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v333)+6)) = uint16(v336)
	v338 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L8
	} else {
		goto L74
	}
L71:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[1]))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v317+(v206^int32(-1))<<(uint(int32(2))%32))))
	v331 = v323
	goto L70
L72:
	;
	goto L73
L73:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[2]))
	v331 = v325 + v206<<(uint(int32(13))%32) + int32(-8192)
	goto L70
L74:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v331)+20)) = uint32(v338)
	F_MarkBufferDirty(m, v212)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	F_MarkBufferDirty(m, v200)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	F_MarkBufferDirty(m, v206)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+48))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+118)))
	if v349 != int32(112) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_ReleaseBuffer(m, v212)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L8
	} else {
		goto L95
	}
L79:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[6]))
	if v353 <= int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v347)+32))
	if v356 != 0 {
		goto L78
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L8
	} else {
		goto L85
	}
L83:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v347)+40))
	if v357 != 0 {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v360 = int32(0)
	F_XLogRegisterBuffer(m, v360, v206, v360)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	F_XLogRegisterBuffer(m, int32(1), v212, int32(8))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	F_XLogRegisterBuffer(m, int32(2), v200, int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)) = uint16(v5)
	v373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331)+16)))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v331+v373)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v375
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v331)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v377
	F_XLogRegisterData(m, v16+int32(4), int32(12))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v386 = F_XLogInsert(m, int32(13), int32(80))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v388 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v331))) = base.I64_rotr(v386, v388)
	v391 = base.I32_wrap_i64(v386)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v391
	v395 = base.I32_wrap_i64(int64(base.Ui64(v386) >> (uint(v388) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v395
	if v200 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414)+4)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v414))) = v395
	goto L78
L92:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[1]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v400+(v200^int32(-1))<<(uint(int32(2))%32))))
	v414 = v406
	goto L91
L93:
	;
	goto L94
L94:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[2]))
	v414 = v408 + v200<<(uint(int32(13))%32) + int32(-8192)
	goto L91
L95:
	;
	F_ReleaseBuffer(m, v200)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_ReleaseBuffer(m, v206)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	v427 = int32(_a_F_ginScanToDelete_1)
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[5]))
	v430 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginScanToDelete[5])) = v429 - v430
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v434)+24)) = v435 + v430
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v439)+28)) = v440 + v430
	if l2 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_LockBuffer(m, v49, int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L8
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_ReleaseBuffer(m, v49)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L8
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	v459 = v430
	goto L29
L103:
	;
	F_UnlockReleaseBuffer(m, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L8
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v49
	v459 = int32(0)
	goto L29
L106:
	;
	goto L105
L107:
	;
	F_ReleaseBuffer(m, v49)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L8
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	m.G0 = v16 + int32(16)
	return v459
L110:
	;
	goto L109
}
func F_gin_clean_pending_list(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(_a_F_gin_clean_pending_list_0)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_index_open(m, v10, int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gin_clean_pending_list[0])))
		if v18 == int32(1) {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_gin_clean_pending_list[1]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+316))
			v26 = base.B2i32(v24 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_gin_clean_pending_list[0])) = uint8(v26)
			v28 = v26
		} else {
			v28 = int32(0)
		}
		if v28 == int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
			if v32 != int32(105) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return int32(0)
					} else {
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v143 + int32(4)
						F_errmsg(m, int32(_a_F_gin_clean_pending_list_1), v8+int32(16))
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1049), int32(_a_F_gin_clean_pending_list_3))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
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
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+84))
				if v35 != int32(2742) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v143 + int32(4)
							F_errmsg(m, int32(_a_F_gin_clean_pending_list_1), v8+int32(16))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1049), int32(_a_F_gin_clean_pending_list_3))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
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
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+118)))
					if v38 == int32(116) {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
						if v41 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_gin_clean_pending_list_4), int32(0))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1059), int32(_a_F_gin_clean_pending_list_3))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
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
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_gin_clean_pending_list[2]))
							v47 = F_object_ownercheck(m, int32(1259), v10, v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								if v47 == int32(0) {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
									F_aclcheck_error(m, int32(2), int32(20), v53+int32(4))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[3]))) = v58
										*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[4]))) = v58
										*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[5]))) = v58
										*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[6]))) = v58
										*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[7]))) = v58
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+192))
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+18)))
										if v69 == int32(1) {
											v73 = v8 + int32(28)
											F_initGinState(m, v73, v12)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												v76 = int32(1)
												F_ginInsertCleanup(m, v73, v76, v76, v76, v8+int32(_a_F_gin_clean_pending_list_5))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[8]))))
													v106 = v83
													F_relation_close(m, v12, int32(3))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int32(0)
													} else {
														v110 = F_Int64GetDatum(m, v106)
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
															return v110
														}
													}
												}
											}
										} else {
											v86 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v86 == int32(0) {
													v106 = v5
													F_relation_close(m, v12, int32(3))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int32(0)
													} else {
														v110 = F_Int64GetDatum(m, v106)
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
															return v110
														}
													}
												} else {
													F_errcode(m, int32(325))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int32(0)
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93 + int32(4)
														F_errmsg(m, int32(_a_F_gin_clean_pending_list_6), v8)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1086), int32(_a_F_gin_clean_pending_list_3))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																v106 = v5
																F_relation_close(m, v12, int32(3))
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return int32(0)
																} else {
																	v110 = F_Int64GetDatum(m, v106)
																	mBase = m.M
																	v111 = m.ExcPending
																	if v111 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
																		return v110
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
									v58 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[3]))) = v58
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[4]))) = v58
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[5]))) = v58
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[6]))) = v58
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[7]))) = v58
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+192))
									v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+18)))
									if v69 == int32(1) {
										v73 = v8 + int32(28)
										F_initGinState(m, v73, v12)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											v76 = int32(1)
											F_ginInsertCleanup(m, v73, v76, v76, v76, v8+int32(_a_F_gin_clean_pending_list_5))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[8]))))
												v106 = v83
												F_relation_close(m, v12, int32(3))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int32(0)
												} else {
													v110 = F_Int64GetDatum(m, v106)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
														return v110
													}
												}
											}
										}
									} else {
										v86 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v86 == int32(0) {
												v106 = v5
												F_relation_close(m, v12, int32(3))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int32(0)
												} else {
													v110 = F_Int64GetDatum(m, v106)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
														return v110
													}
												}
											} else {
												F_errcode(m, int32(325))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int32(0)
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93 + int32(4)
													F_errmsg(m, int32(_a_F_gin_clean_pending_list_6), v8)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1086), int32(_a_F_gin_clean_pending_list_3))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int32(0)
														} else {
															v106 = v5
															F_relation_close(m, v12, int32(3))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return int32(0)
															} else {
																v110 = F_Int64GetDatum(m, v106)
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
																	return v110
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
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_gin_clean_pending_list[2]))
						v47 = F_object_ownercheck(m, int32(1259), v10, v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							if v47 == int32(0) {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
								F_aclcheck_error(m, int32(2), int32(20), v53+int32(4))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v58 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[3]))) = v58
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[4]))) = v58
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[5]))) = v58
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[6]))) = v58
									*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[7]))) = v58
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+192))
									v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+18)))
									if v69 == int32(1) {
										v73 = v8 + int32(28)
										F_initGinState(m, v73, v12)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											v76 = int32(1)
											F_ginInsertCleanup(m, v73, v76, v76, v76, v8+int32(_a_F_gin_clean_pending_list_5))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[8]))))
												v106 = v83
												F_relation_close(m, v12, int32(3))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int32(0)
												} else {
													v110 = F_Int64GetDatum(m, v106)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
														return v110
													}
												}
											}
										}
									} else {
										v86 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v86 == int32(0) {
												v106 = v5
												F_relation_close(m, v12, int32(3))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int32(0)
												} else {
													v110 = F_Int64GetDatum(m, v106)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
														return v110
													}
												}
											} else {
												F_errcode(m, int32(325))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int32(0)
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93 + int32(4)
													F_errmsg(m, int32(_a_F_gin_clean_pending_list_6), v8)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1086), int32(_a_F_gin_clean_pending_list_3))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int32(0)
														} else {
															v106 = v5
															F_relation_close(m, v12, int32(3))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return int32(0)
															} else {
																v110 = F_Int64GetDatum(m, v106)
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
																	return v110
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
								v58 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[3]))) = v58
								*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[4]))) = v58
								*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[5]))) = v58
								*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[6]))) = v58
								*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[7]))) = v58
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+192))
								v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+18)))
								if v69 == int32(1) {
									v73 = v8 + int32(28)
									F_initGinState(m, v73, v12)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v76 = int32(1)
										F_ginInsertCleanup(m, v73, v76, v76, v76, v8+int32(_a_F_gin_clean_pending_list_5))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[8]))))
											v106 = v83
											F_relation_close(m, v12, int32(3))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int32(0)
											} else {
												v110 = F_Int64GetDatum(m, v106)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
													return v110
												}
											}
										}
									}
								} else {
									v86 = F_errstart(m, int32(14), int32(0))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										if v86 == int32(0) {
											v106 = v5
											F_relation_close(m, v12, int32(3))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int32(0)
											} else {
												v110 = F_Int64GetDatum(m, v106)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
													return v110
												}
											}
										} else {
											F_errcode(m, int32(325))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93 + int32(4)
												F_errmsg(m, int32(_a_F_gin_clean_pending_list_6), v8)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1086), int32(_a_F_gin_clean_pending_list_3))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int32(0)
													} else {
														v106 = v5
														F_relation_close(m, v12, int32(3))
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return int32(0)
														} else {
															v110 = F_Int64GetDatum(m, v106)
															mBase = m.M
															v111 = m.ExcPending
															if v111 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
																return v110
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
			v119 = m.ExcPending
			if v119 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_gin_clean_pending_list_7), int32(0))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_gin_clean_pending_list_8), int32(0))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1041), int32(_a_F_gin_clean_pending_list_3))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(4)
			v21 = int32(1)
			v22 = v17 + v21
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v27 = v12 + v21
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			v32 = v30 & v21
			if v32 != 0 {
				v33 = v27
			} else {
				v33 = v12 + v19
			}
			if v30 == int32(1) {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
				if v39 == int32(18) {
					v42 = int32(16)
				} else {
					v42 = int32(0)
				}
				if base.Ui32((v39-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v49 = int32(4)
				} else {
					v49 = v42
				}
				v60 = v49
			} else {
				v50 = int32(1)
				if v32 != 0 {
					v60 = int32(base.Ui32(v30)>>(uint(v50)%32)) - v50
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v23&v21 != 0 {
				v61 = v22
			} else {
				v61 = v17 + v19
			}
			if v23 == int32(1) {
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				if v67 == int32(18) {
					v70 = int32(16)
				} else {
					v70 = int32(0)
				}
				if base.Ui32((v67-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v77 = int32(4)
				} else {
					v77 = v70
				}
				v90 = v77
			} else {
				v78 = int32(1)
				if v23&v78 != 0 {
					v90 = int32(base.Ui32(v23)>>(uint(v78)%32)) - v78
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v92 = F_varstr_cmp(m, v33, v60, v61, v90, int32(950))
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v94 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v98 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								return v92
							}
						} else {
							return v92
						}
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v98 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							return v92
						}
					} else {
						return v92
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
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
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
	return v73
L2:
	;
	v73 = int32(1)
	goto L1
L3:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v68)
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
		v73 = v17
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
	v73 = int32(0)
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
	v73 = v17
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
	v73 = int32(0)
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
	v73 = v48
	goto L1
L21:
	;
	return int32(0)
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	F_errmsg_internal(m, int32(_a_F_gin_consistent_hstore_0), v9)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_gin_consistent_hstore_1), int32(207), int32(_a_F_gin_consistent_hstore_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
	var v121 int32
	_ = v121
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
	return v121
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	v121 = v2
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
	v44 = v21 << (uint(int32(1)) % 32)
	v45 = v28
	v46 = v2
	goto L10
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v46
	v121 = v45
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
	if v46 < v44 {
		v84 = v44
		v85 = v45
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85+v46<<(uint(int32(2))%32)))) = v72
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v91
	v44 = v84
	v45 = v85
	v46 = v46 + int32(1)
	goto L10
L24:
	;
	if v44 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = F_repalloc(m, v45, v44<<(uint(int32(3))%32))
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
	v84 = v44 << (uint(int32(1)) % 32)
	v85 = v76
	goto L23
L29:
	;
	v84 = int32(8)
	v85 = v82
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
	goto L33
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(0)
	goto L33
L33:
	;
	v41 = v95
	goto L10
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v54
	F_errmsg_internal(m, int32(_a_F_gin_extract_jsonb_path_0), v11)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_gin_extract_jsonb_path_1), int32(1170), int32(_a_F_gin_extract_jsonb_path_2))
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
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(_a_F_gin_extract_query_bit_0), int32(2645))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_int2(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13927(m, l0, int32(-32768), int32(2096))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_int8(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13928(m, l0, int64(-9223372036854775807-1), int32(2272))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_interval(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(0), int32(_a_F_gin_extract_query_interval_0), int32(2524))
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
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
				v34 = F_pg_detoast_datum(m, v15)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(_a_F_gin_extract_query_numeric_0)
					v38 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v38)
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v34
					*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
					v43 = F_palloc(m, int32(4))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v43))) = v24
						v48 = v14 & int32(_a_F_gin_extract_query_numeric_1)
						switch v48 - int32(1) {
						case 0, 1:
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
							v69 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v69)
							m.G0 = v11 + int32(16)
							return v19
						case 2:
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v34
							m.G0 = v11 + int32(16)
							return v19
						case 3, 4:
							v51 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v51)
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v34
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
								F_errmsg_internal(m, int32(_a_F_gin_extract_query_numeric_2), v11)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_gin_extract_query_numeric_3), int32(97), int32(_a_F_gin_extract_query_numeric_4))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
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
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(_a_F_gin_extract_query_text_0), int32(2104))
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
	v5 = F_gin_btree_extract_query(m, l0, int32(0), int32(_a_F_gin_extract_query_timetz_0), int32(1427))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_uuid(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13929(m, l0, int32(16), int32(3384))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
			F_errmsg_internal(m, int32(_a_F_gin_extract_tsquery_5args_0), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_gin_extract_tsquery_5args_1), int32(319), int32(_a_F_gin_extract_tsquery_5args_2))
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
	var v32 int32
	_ = v32
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
	var v63 int32
	_ = v63
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
		v63 = int32(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v65 != v9 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v20 = F_palloc(m, v14<<(uint(int32(2))%32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v22 <= int32(0) {
		v63 = v20
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v27 = v9 + int32(8)
	v30 = int32(0)
	v31 = v22
	v32 = v27
	goto L7
L7:
	;
	v35 = int32(2)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v49 = F_cstring_to_text_with_len(m, v27+v31<<(uint(v35)%32)+int32(base.Ui32(v41)>>(uint(int32(12))%32)), int32(base.Ui32(v41)>>(uint(int32(1))%32))&int32(2047))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v63 = v20
	goto L3
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v30<<(uint(v35)%32)))) = v49
	v55 = v30 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v55 < v56 {
		v30 = v55
		v31 = v56
		v32 = v32 + int32(4)
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_pfree(m, v9)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	return v63
L14:
	;
	goto L13
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
	if base.Ui32((v12-int32(9))&int32(_a_F_gin_triconsistent_jsonb_0)) <= base.Ui32(int32(1)) {
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
	if base.Ui32((v12-int32(15))&int32(_a_F_gin_triconsistent_jsonb_0)) <= base.Ui32(int32(1)) {
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
	F_errmsg_internal(m, int32(_a_F_gin_triconsistent_jsonb_1), v8)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_gin_triconsistent_jsonb_2), int32(1072), int32(_a_F_gin_triconsistent_jsonb_3))
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
	if base.Ui32((v12-int32(15))&int32(_a_F_gin_triconsistent_jsonb_path_0)) <= base.Ui32(int32(1)) {
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
	F_errmsg_internal(m, int32(_a_F_gin_triconsistent_jsonb_path_1), v8)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_gin_triconsistent_jsonb_path_2), int32(1314), int32(_a_F_gin_triconsistent_jsonb_path_3))
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
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
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v199 int64
	_ = v199
	var v201 int64
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
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
	var v245 int64
	_ = v245
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v251 int64
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v280 int32
	_ = v280
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
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
	var v323 int64
	_ = v323
	var v325 int32
	_ = v325
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	v3 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	base.MemoryFill(m, l0+int32(4), v3, int32(_a_F_initGinState_0))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(base.B2i32(v31 == int32(1)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v3 < v36 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L12
	} else {
		goto L62
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L12
	} else {
		goto L57
	}
L3:
	;
	v54 = l0 + int32(140)
	v59 = v36
	v62 = v3
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v22 + int32(32)
	return
L6:
	;
	v83 = v24 + v59<<(uint(int32(4))%32) + v62*int32(100) + int32(20)
	v85 = v62 << (uint(int32(2)) % 32)
	v86 = l0 + int32(12) + v85
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v87 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v123 = int32(1)
	v124 = v62 + v123
	v125 = base.I32_extend16_s(v124)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+6)))
	v137 = int32(4)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v127+v129*(v125-v123)<<(uint(int32(2))%32)+v137-v137)))
	goto L18
L9:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v90
	goto L8
L10:
	;
	goto L11
L11:
	;
	v93 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v93
	v97 = int32(0)
	F_TupleDescInitEntry(m, v93, int32(1), v97, int32(21), int32(-1), v97)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v83)+68))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+80)))
	F_TupleDescInitEntry(m, v103, int32(2), int32(0), v106, v107, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v83)+96))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v111+v114<<(uint(int32(4))%32)+int32(200))+16)) = v113
	goto L16
L16:
	;
	goto L8
L17:
	;
	v188 = v62 * int32(28)
	v189 = l0 + int32(1036) + v188
	v191 = F_index_getprocinfo(m, l1, v125, int32(2))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L12
	} else {
		goto L27
	}
L18:
	;
	if v141 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v144 = v54 + v62*int32(28)
	v146 = F_index_getprocinfo(m, l1, v125, int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v83)+68))
	v163 = F_lookup_type_cache(m, v161, int32(64))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L24
	}
L22:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v146)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v144)+16)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = v152
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v146)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v144)+8)) = v154
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v146)))
	*(*int64)(unsafe.Add(mBase, uint32(v144))) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v144)+16)) = int32(0)
	goto L23
L23:
	;
	goto L17
L24:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+108))
	if v165 == int32(0) {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v170 = v54 + v62*int32(28)
	v172 = v163 + int32(104)
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v172)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+16)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = v177
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v172)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = v179
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	*(*int64)(unsafe.Add(mBase, uint32(v170))) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v170)+16)) = int32(0)
	goto L26
L26:
	;
	goto L17
L27:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v191)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v189)+16)) = v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+24)) = v197
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v191)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v189)+8)) = v199
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v191)))
	*(*int64)(unsafe.Add(mBase, uint32(v189))) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v189)+20)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v189)+16)) = int32(0)
	goto L28
L28:
	;
	v206 = v188 + (l0 + int32(1932))
	v208 = F_index_getprocinfo(m, l1, v125, int32(3))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v208)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v206)+16)) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v208)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+24)) = v214
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v208)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v206)+8)) = v216
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v208)))
	*(*int64)(unsafe.Add(mBase, uint32(v206))) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v206)+20)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v206)+16)) = int32(0)
	goto L30
L30:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225)+6)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v224+v226*(v125-int32(1))<<(uint(int32(2))%32)+int32(24)-int32(4))))
	goto L31
L31:
	;
	if v238 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v239 = v188 + (l0 + int32(3724))
	v241 = F_index_getprocinfo(m, l1, v125, int32(6))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+6)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v257+v259*(v125-int32(1))<<(uint(int32(2))%32)+int32(16)-int32(4))))
	goto L37
L35:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v241)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v239)+16)) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v239)+24)) = v247
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v241)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v239)+8)) = v249
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v241)))
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v239)+20)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v239)+16)) = int32(0)
	goto L36
L36:
	;
	goto L34
L37:
	;
	if v271 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v272 = v188 + (l0 + int32(2828))
	v274 = F_index_getprocinfo(m, l1, v125, int32(4))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v289 = l0 + v188
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v289+int32(2832))))
	if v292 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v274)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+16)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v274)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v272)+24)) = v280
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v274)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+8)) = v282
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v274)))
	*(*int64)(unsafe.Add(mBase, uint32(v272))) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v272)+20)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v272)+16)) = int32(0)
	goto L42
L42:
	;
	goto L40
L43:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v289+int32(3728))))
	if v297 == int32(0) {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v303)+6)))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v302+v304*(v125-int32(1))<<(uint(int32(2))%32)+int32(20)-int32(4))))
	goto L47
L46:
	;
	goto L45
L47:
	;
	if v316 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v317 = v188 + (l0 + int32(_a_F_initGinState_1))
	v319 = F_index_getprocinfo(m, l1, v125, int32(5))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L12
	} else {
		goto L51
	}
L49:
	;
	v336 = int32(0)
	goto L50
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62+(l0+int32(_a_F_initGinState_2))))) = uint8(v336)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v339+v85)))
	if v341 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v319)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v317)+16)) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v319)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+24)) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v319)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v317)+8)) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v319)))
	*(*int64)(unsafe.Add(mBase, uint32(v317))) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v317)+20)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v317)+16)) = int32(0)
	goto L52
L52:
	;
	v336 = int32(1)
	goto L50
L53:
	;
	v343 = v341
	goto L55
L54:
	;
	v343 = int32(100)
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85+(l0+int32(_a_F_initGinState_3))))) = v343
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v124 < v345 {
		v59 = v345
		v62 = v124
		goto L6
	} else {
		goto L56
	}
L56:
	;
	goto L7
L57:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v83)+68))
	v377 = F_format_type_be(m, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v377
	F_errmsg(m, int32(_a_F_initGinState_4), v22)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_initGinState_5), int32(153), int32(_a_F_initGinState_6))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v124
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = int64(25769803780)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v392 + int32(4)
	F_errmsg_internal(m, int32(_a_F_initGinState_7), v22+int32(16))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_initGinState_5), int32(190), int32(_a_F_initGinState_6))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
