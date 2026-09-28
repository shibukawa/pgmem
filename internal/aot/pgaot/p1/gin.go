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
	var v46 int64
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
	var v150 int64
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
	v15 = v13 - int32(32)
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
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l2+v33<<(uint(int32(3))%32))))
	v47 = int32(_a_F_ginBuildCallback_0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[0]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[0])) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[5])))
	v54 = v33 + int32(1)
	v56 = v54 & int32(_a_F_ginBuildCallback_1)
	v61 = F_ginExtractEntries(m, v52, v56, v46, v42, v15+int32(28), v15+int32(16))
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
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
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
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
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
	*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallback[11]))) = int32(836)
	m.G0 = v105 + v104
	goto L14
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallback[0])) = v18
	m.G0 = v15 + int32(32)
	return
L14:
	;
	v129 = F_ginGetBAEntry(m, v99, v15+int32(12), v15+int32(16), v15+int32(15), v15+int32(28))
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
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	v151 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+15)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
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
	v163 = F_ginGetBAEntry(m, v99, v15+int32(12), v15+int32(16), v15+int32(15), v15+int32(28))
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
func F_ginExtractEntries(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
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
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int64
	_ = v116
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
	var v133 int64
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int64
	_ = v171
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v257 int64
	_ = v257
	var v258 int64
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int64
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v300 int32
	_ = v300
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(32)
	return v339
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	v24 = F_palloc(m, int32(8))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v36
	v40 = int32(28)
	v46 = l1 - int32(1)
	v49 = l0 + v46<<(uint(int32(2))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_c_F_ginExtractEntries[0])))
	v59 = F_FunctionCall3Coll(m, l1*v40+l0+int32(1008), v52, l2, base.I64_extend_i32_u(v19+int32(24)), base.I64_extend_i32_u(v19+v40))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L9
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = int64(0)
	v31 = F_palloc(m, int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v31
	v34 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v34)
	v339 = v24
	goto L1
L8:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v79 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v61 = base.I32_wrap_i64(v59)
	if v61 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if int32(0) < v62 {
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
	v69 = F_palloc(m, int32(8))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = int64(0)
	v74 = F_palloc(m, int32(1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v74
	v77 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v77)
	v339 = v69
	goto L1
L16:
	;
	if v195 < int32(2) {
		v300 = v195
		goto L39
	} else {
		goto L40
	}
L17:
	;
	v195 = v62
	v202 = v7
	goto L16
L18:
	;
	goto L19
L19:
	;
	v82 = int32(0)
	if v62 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v178
	v195 = v178
	v202 = v185
	goto L16
L21:
	;
	v92 = v82
	v94 = v82
	v99 = int32(0)
	v101 = v7
	goto L24
L22:
	;
	v147 = v82
	v149 = v82
	v156 = v7
	goto L23
L23:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v79))))
	if v163 != 0 {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v79))))
	if v108 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v62&int32(1) == int32(0) {
		v178 = v137
		v185 = v138
		goto L20
	} else {
		goto L35
	}
L26:
	;
	v123 = v92 | int32(1)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v123))))
	if v125 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v120 = v94
	v121 = int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v110 = int32(3)
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v61+v92<<(uint(v110)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v61+v94<<(uint(v110)%32)))) = v116
	v120 = v94 + int32(1)
	v121 = v101
	goto L26
L30:
	;
	v139 = int32(2)
	v140 = v92 + v139
	v142 = v99 + v139
	if v142 != v62&int32(2147483646) {
		v92 = v140
		v94 = v137
		v99 = v142
		v101 = v138
		goto L24
	} else {
		goto L34
	}
L31:
	;
	v137 = v120
	v138 = int32(1)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v127 = int32(3)
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v61+v123<<(uint(v127)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v61+v120<<(uint(v127)%32)))) = v133
	v137 = v120 + int32(1)
	v138 = v121
	goto L30
L34:
	;
	goto L25
L35:
	;
	v147 = v140
	v149 = v137
	v156 = v138
	goto L23
L36:
	;
	v178 = v149
	v185 = int32(1)
	goto L20
L37:
	;
	goto L38
L38:
	;
	v165 = int32(3)
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v61+v147<<(uint(v165)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v61+v149<<(uint(v165)%32)))) = v171
	v178 = v149 + int32(1)
	v185 = v156
	goto L20
L39:
	;
	v315 = F_palloc0_mul(m, int32(1), v300+v202)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L55
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l0 + v46*int32(28) + int32(140)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_c_F_ginExtractEntries[0])))
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+20)) = uint8(v217)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v216
	F_qsort_arg_entries(m, v61, v195, v19+int32(12))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v225 = int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+20)))
	if v227 != v225 {
		v300 = v226
		goto L39
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v226) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v232 = v217
	v233 = v225
	goto L46
L44:
	;
	v283 = v226
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v283
	v300 = v283
	goto L39
L46:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v250 = int32(3)
	v252 = v61 + v233<<(uint(v250)%32)
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v61+v232<<(uint(v250)%32))))
	v258 = F_FunctionCall2Coll(m, v248, v249, v253, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L49
	}
L47:
	;
	v283 = v273 + int32(1)
	goto L45
L48:
	;
	v276 = v233 + int32(1)
	if v276 != v226 {
		v232 = v273
		v233 = v276
		goto L46
	} else {
		goto L54
	}
L49:
	;
	if base.I32_wrap_i64(v258) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v263 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+20)) = uint8(v263)
	v273 = v232
	goto L48
L51:
	;
	goto L52
L52:
	;
	v266 = v232 + int32(1)
	if v233 == v266 {
		v273 = v233
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	*(*int64)(unsafe.Add(mBase, uint32(v61+v266<<(uint(int32(3))%32)))) = v271
	v273 = v266
	goto L48
L54:
	;
	goto L47
L55:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v202 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v61+v317<<(uint(int32(3))%32)))) = int64(0)
	v324 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v317+v315))) = uint8(v324)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v329 = v326 + v324
	goto L58
L57:
	;
	v329 = v317
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v329
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v315
	v339 = v61
	goto L1
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
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
	var v180 int32
	_ = v180
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
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
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
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
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
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
	F_LockBufferInternal(m, v32, int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	F_UnlockBuffer(m, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
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
		v480 = v31
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v491 = m.T0[v490].(func(*base.Module, int32, int32) int32)(m, l0, v489)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v171)+20)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v463
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+8)) = uint16(v469)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v171
	v480 = v171
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
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L115
	}
L30:
	;
	goto L29
L31:
	;
	F_UnlockBuffer(m, v125)
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
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v364 = F_ginStepRight(m, v125, v362, int32(3))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
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
	v144 = v133
	goto L38
L36:
	;
	v162 = v133
	goto L37
L37:
	;
	v166 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+8)) = uint16(v166)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
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
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
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
	v162 = v151
	goto L37
L40:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
	if v152 != 0 {
		v144 = v151
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v178 = v169
	v180 = v168
	goto L43
L43:
	;
	F_LockBufferInternal(m, v180, int32(3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	if v180 < int32(0) {
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
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192+(v180^int32(-1))<<(uint(int32(2))%32))))
	v206 = v198
	goto L46
L48:
	;
	goto L49
L49:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
	v206 = v200 + v180<<(uint(int32(13))%32) + int32(-8192)
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
	*(*int32)(unsafe.Add(mBase, uint32(v171)+20)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v178
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
	F_errmsg_internal(m, int32(_a_F_ginFinishSplit_0), v14+int32(-48))
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
	F_errfinish(m, int32(_a_F_ginFinishSplit_1), int32(783), int32(_a_F_ginFinishSplit_2))
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
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v360 = F_ReadBuffer(m, v359, v247)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
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
	v262 = v206
	v263 = v180
	goto L67
L65:
	;
	goto L66
L66:
	;
	if v178 != int32(-1) {
		v463 = v178
		v465 = v180
		v469 = v252
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
	F_UnlockBuffer(m, v263)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v282 = F_ginStepRight(m, v263, v280, int32(3))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L76
	}
L72:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v263 == v276 {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	F_ReleaseBuffer(m, v263)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	goto L62
L75:
	;
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301)+16)))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302+v301)+6)))
	if v304&int32(64) != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	if v282 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v287+(v282^int32(-1))<<(uint(int32(2))%32))))
	v301 = v293
	goto L75
L78:
	;
	goto L79
L79:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
	v301 = v295 + v282<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+20)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v271
	v310 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+8)) = uint16(v310)
	v314 = F_errstart(m, int32(14), v310)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v340 = m.T0[v339].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v301, v337, int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L90
	}
L83:
	;
	if v314 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+48))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v317 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginFinishSplit_0), v16)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v332 = int32(0)
	F_ginFinishSplit(m, l0, v171, v332, v332)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ginFinishSplit_1), int32(783), int32(_a_F_ginFinishSplit_2))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
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
	if v340 == int32(0) {
		v262 = v301
		v263 = v282
		goto L67
	} else {
		goto L91
	}
L91:
	;
	v463 = v271
	v465 = v282
	v469 = v340
	goto L27
L92:
	;
	goto L62
L93:
	;
	v178 = v247
	v180 = v360
	goto L43
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v364
	if v364 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v387 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[2]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v370+(v364^int32(-1))*int32(56))+16))
	v385 = v376
	goto L95
L97:
	;
	goto L98
L98:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[3]))
	v379 = int32(56)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v378+v364*v379-v379)+16))
	v385 = v384
	goto L95
L99:
	;
	v406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v405)+16)))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v405)+6)))
	if v408&int32(64) != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v391+(v387^int32(-1))<<(uint(int32(2))%32))))
	v405 = v397
	goto L99
L101:
	;
	goto L102
L102:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
	v405 = v399 + v387<<(uint(int32(13))%32) + int32(-8192)
	goto L99
L103:
	;
	v413 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v440 = m.T0[v439].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v405, v437, v438)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L113
	}
L106:
	;
	if v413 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+48))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v416 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginFinishSplit_0), v14+int32(-32))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
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
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	F_errfinish(m, int32(_a_F_ginFinishSplit_1), int32(783), int32(_a_F_ginFinishSplit_2))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
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
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+8)) = uint16(v440)
	if v440 == int32(0) {
		v118 = v405
		goto L28
	} else {
		goto L114
	}
L114:
	;
	v480 = v31
	goto L25
L115:
	;
	F_errmsg_internal(m, int32(_a_F_ginFinishSplit_3), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ginFinishSplit_1), int32(256), int32(_a_F_ginFinishSplit_4))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
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
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v493 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511)+16)))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v512+v511)))
	v515 = F_ginPlaceToPage(m, l0, v480, v491, v514, v493, l3)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L4
	} else {
		goto L123
	}
L120:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[0]))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v497+(v493^int32(-1))<<(uint(int32(2))%32))))
	v511 = v503
	goto L119
L121:
	;
	goto L122
L122:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishSplit[1]))
	v511 = v505 + v493<<(uint(int32(13))%32) + int32(-8192)
	goto L119
L123:
	;
	F_pfree(m, v491)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
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
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	F_UnlockBuffer(m, v521)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
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
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	F_ReleaseBuffer(m, v524)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L4
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	if v515 == int32(0) {
		v19 = v480
		v30 = int32(1)
		goto L2
	} else {
		goto L144
	}
L132:
	;
	F_pfree(m, v19)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	if v515 == int32(0) {
		v19 = v480
		v30 = int32(1)
		goto L2
	} else {
		goto L134
	}
L134:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	F_UnlockBuffer(m, v532)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v539 = v480
	goto L136
L136:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v539)+20))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v549 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	goto L1
L138:
	;
	F_ReleaseBuffer(m, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	F_pfree(m, v539)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L4
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	if v548 != 0 {
		v539 = v548
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[5]))) = int32(836)
	m.G0 = v24 + v23
	goto L1
L1:
	;
	v43 = base.I32_div_u_s(int32(178956970), v16<<(uint(int32(1))%32))
	v52 = F_ginGetBAEntry(m, v18, v13+int32(14), v13+int32(24), v13+int32(23), v13+int32(16))
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
	v62 = v52
	goto L7
L5:
	;
	goto L6
L6:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[6])))
	F_MemoryContextReset(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L27
	}
L7:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v66 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+14)))
	v70 = v15 + int32(20) + v67<<(uint(int32(3))%32)
	v73 = int32(0)
	v75 = v66
	goto L12
L10:
	;
	goto L11
L11:
	;
	v128 = F_ginGetBAEntry(m, v18, v13+int32(14), v13+int32(24), v13+int32(23), v13+int32(16))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L25
	}
L12:
	;
	v82 = v75 - v73
	if base.Ui32(v43) < base.Ui32(v82) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v84 = v43
	goto L16
L15:
	;
	v84 = v82
	goto L16
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_ginFlushBuildState[7]))
	if v86 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+14)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+23)))
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v70)+2)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)))
	v99 = F__gin_build_tuple(m, v89, v90, v91, v92, v93, v62+v73*int32(6), v84, v13+int32(8))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFlushBuildState[8])))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_tuplesort_putgintuple(m, v101, v99, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_pfree(m, v99)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v107 = v73 + v84
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if base.Ui32(v107) < base.Ui32(v108) {
		v73 = v107
		v75 = v108
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	if v128 != 0 {
		v62 = v128
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
	v144 = m.ExcPending
	if v144 != 0 {
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
func F_ginScanPostingTreeToDelete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v70 int32
	_ = v70
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
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
	var v385 int64
	_ = v385
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v15 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v141 = v33 + v134
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+6)))
	if v142&int32(2) != 0 {
		goto L29
	} else {
		goto L30
	}
L2:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+16)))
	v35 = v34 + v33
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+6)))
	if v36&int32(2) != 0 {
		v134 = v34
		goto L1
	} else {
		goto L6
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+(v15^int32(-1))<<(uint(int32(2))%32))))
	v33 = v25
	goto L2
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[1]))
	v33 = v27 + v15<<(uint(int32(13))%32) + int32(-8192)
	goto L2
L6:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+4)))
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v46 = int32(1)
	goto L10
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+16)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v33+v116)))
	if v118 != int32(-1) {
		v134 = v116
		goto L1
	} else {
		goto L21
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v54 = int32(0)
	v59 = v33 + int32(22) + v46&int32(_a_F_ginScanPostingTreeToDelete_0)*int32(10)
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59))))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+2)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginScanPostingTreeToDelete[2])))
	v67 = F_ReadBufferExtended(m, v53, v54, v60<<(uint(int32(16))%32)|v63, v54, v66)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	return int32(0)
L13:
	;
	F_LockBufferInternal(m, v67, int32(3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v74 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v85 = v74
	goto L17
L16:
	;
	v76 = F_palloc0(m, int32(20))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v67
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v88 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+18)) = uint8(v88)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)) = uint16(v46)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v93 = F_ginScanPostingTreeToDelete(m, l0, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = l1
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v85 = v83
	goto L17
L19:
	;
	v97 = v46 + (v93 ^ int32(1))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+16)))
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33+v100)+4)))
	if base.Ui32(v97&int32(_a_F_ginScanPostingTreeToDelete_0)) <= base.Ui32(v102) {
		v46 = v97
		goto L10
	} else {
		goto L20
	}
L20:
	;
	goto L11
L21:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	if v122 == int32(0) {
		v134 = v116
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_UnlockReleaseBuffer(m, v122)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = int32(0)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+16)))
	v134 = v130
	goto L1
L24:
	;
	m.G0 = v13 + int32(16)
	return v435
L25:
	;
	v415 = int32(_a_F_ginScanPostingTreeToDelete_1)
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[3]))
	v418 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[3])) = v417 - v418
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v422)+24)) = v423 + v418
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v427)+28)) = v428 + v418
	F_UnlockReleaseBuffer(m, v15)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L12
	} else {
		goto L94
	}
L26:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v407+v156<<(uint(int32(13))%32))+uint32(_c_F_ginScanPostingTreeToDelete[4]))) = v385
	goto L25
L27:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v401 != 0 {
		goto L90
	} else {
		goto L91
	}
L28:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.B2i32(v156 == int32(0))|v142&int32(64) != 0 {
		goto L27
	} else {
		goto L38
	}
L29:
	;
	if v142&int32(128) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+4)))
	if v155 != 0 {
		goto L27
	} else {
		goto L37
	}
L32:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+4)))
	if v149 == int32(0) {
		goto L28
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)))
	if v152 != int32(32) {
		goto L27
	} else {
		goto L36
	}
L35:
	;
	goto L27
L36:
	;
	goto L28
L37:
	;
	goto L28
L38:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v162 == int32(-1) {
		goto L27
	} else {
		goto L39
	}
L39:
	;
	if v156 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182)+16)))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v182)+6)))
	if v185&int32(64) != 0 {
		goto L27
	} else {
		goto L44
	}
L41:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[0]))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168+(v156^int32(-1))<<(uint(int32(2))%32))))
	v182 = v174
	goto L40
L42:
	;
	goto L43
L43:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[1]))
	v182 = v176 + v156<<(uint(int32(13))%32) + int32(-8192)
	goto L40
L44:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	if v15 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[5]))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194+(v15^int32(-1))*int32(56))+16))
	v209 = v200
	goto L45
L47:
	;
	goto L48
L48:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[6]))
	v203 = int32(56)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+v15*v203-v203)+16))
	v209 = v208
	goto L45
L49:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+16)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v229+v228)))
	F_PredicateLockPageSplit(m, v210, v209, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L12
	} else {
		goto L53
	}
L50:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[0]))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214+(v15^int32(-1))<<(uint(int32(2))%32))))
	v228 = v220
	goto L49
L51:
	;
	goto L52
L52:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[1]))
	v228 = v222 + v15<<(uint(int32(13))%32) + int32(-8192)
	goto L49
L53:
	;
	v234 = int32(_a_F_ginScanPostingTreeToDelete_1)
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[3])) = v236 + int32(1)
	if v156 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v258+v257))) = v231
	if v190 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[0]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v243+(v156^int32(-1))<<(uint(int32(2))%32))))
	v257 = v249
	goto L54
L56:
	;
	goto L57
L57:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[1]))
	v257 = v251 + v156<<(uint(int32(13))%32) + int32(-8192)
	goto L54
L58:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+16)))
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278+v281)+4)))
	if v283 != v188 {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[0]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v264+(v190^int32(-1))<<(uint(int32(2))%32))))
	v278 = v270
	goto L58
L60:
	;
	goto L61
L61:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[1]))
	v278 = v272 + v190<<(uint(int32(13))%32) + int32(-8192)
	goto L58
L62:
	;
	if v15 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	v287 = (v283 - v188) * int32(10)
	if v287 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v300 = v281
	goto L65
L65:
	;
	v303 = v283 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v300+v278)+4)) = uint16(v303)
	v308 = v283*int32(10) + int32(22)
	*(*uint16)(unsafe.Add(mBase, uint32(v278)+12)) = uint16(v308)
	goto L62
L66:
	;
	v290 = v278 + v188*int32(10)
	base.MemoryCopy(m, v290+int32(22), v290+int32(32), v287)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+16)))
	v300 = v297
	goto L65
L69:
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
		goto L12
	} else {
		goto L73
	}
L70:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[0]))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v313+(v15^int32(-1))<<(uint(int32(2))%32))))
	v327 = v319
	goto L69
L71:
	;
	goto L72
L72:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[1]))
	v327 = v321 + v15<<(uint(int32(13))%32) + int32(-8192)
	goto L69
L73:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v327)+20)) = uint32(v334)
	F_MarkBufferDirty(m, v190)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	F_MarkBufferDirty(m, v156)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	F_MarkBufferDirty(m, v15)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+48))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+118)))
	if v345 != int32(112) {
		goto L25
	} else {
		goto L77
	}
L77:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[7]))
	if v349 <= int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v343)+32))
	if v352 != 0 {
		goto L25
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L12
	} else {
		goto L83
	}
L81:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v343)+40))
	if v353 != 0 {
		goto L25
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v356 = int32(0)
	F_XLogRegisterBuffer(m, v356, v15, v356)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L12
	} else {
		goto L84
	}
L84:
	;
	F_XLogRegisterBuffer(m, int32(1), v190, int32(8))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	F_XLogRegisterBuffer(m, int32(2), v156, int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v188)
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v327)+16)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v327+v369)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v371
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v327)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v373
	F_XLogRegisterData(m, v13+int32(4), int32(12))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L12
	} else {
		goto L87
	}
L87:
	;
	v382 = F_XLogInsert(m, int32(13), int32(80))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L12
	} else {
		goto L88
	}
L88:
	;
	v385 = base.I64_rotl(v382, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v327))) = v385
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v385
	if int32(0) <= v156 {
		goto L26
	} else {
		goto L89
	}
L89:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_ginScanPostingTreeToDelete[0]))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v391+(v156^int32(-1))<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v397))) = v385
	goto L25
L90:
	;
	F_UnlockReleaseBuffer(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L12
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v15
	v435 = int32(0)
	goto L24
L93:
	;
	goto L92
L94:
	;
	v435 = v418
	goto L24
}
func F_gin_clean_pending_list(m *base.Module, l0 int32) int64 {
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
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(_a_F_gin_clean_pending_list_0)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_index_open(m, v10, int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gin_clean_pending_list[0])))
		if v18 == int32(1) {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_gin_clean_pending_list[1]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+308))
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
				v137 = m.ExcPending
				if v137 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int64(0)
					} else {
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v141 + int32(4)
						F_errmsg(m, int32(_a_F_gin_clean_pending_list_1), v8+int32(16))
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1052), int32(_a_F_gin_clean_pending_list_3))
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
								return int64(0)
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
					v137 = m.ExcPending
					if v137 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int64(0)
						} else {
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v141 + int32(4)
							F_errmsg(m, int32(_a_F_gin_clean_pending_list_1), v8+int32(16))
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1052), int32(_a_F_gin_clean_pending_list_3))
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return int64(0)
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
							v158 = m.ExcPending
							if v158 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_gin_clean_pending_list_4), int32(0))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1062), int32(_a_F_gin_clean_pending_list_3))
										mBase = m.M
										v170 = m.ExcPending
										if v170 != 0 {
											return int64(0)
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
								return int64(0)
							} else {
								if v47 == int32(0) {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
									F_aclcheck_error(m, int32(2), int32(20), v53+int32(4))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int64(0)
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
												return int64(0)
											} else {
												v76 = int32(1)
												F_ginInsertCleanup(m, v73, v76, v76, v76, v8+int32(_a_F_gin_clean_pending_list_5))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int64(0)
												} else {
													v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[8]))))
													v106 = v83
													F_relation_close(m, v12, int32(3))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int64(0)
													} else {
														m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
														return v106
													}
												}
											}
										} else {
											v86 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int64(0)
											} else {
												if v86 == int32(0) {
													v106 = v5
													F_relation_close(m, v12, int32(3))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int64(0)
													} else {
														m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
														return v106
													}
												} else {
													F_errcode(m, int32(325))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int64(0)
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93 + int32(4)
														F_errmsg(m, int32(_a_F_gin_clean_pending_list_6), v8)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1089), int32(_a_F_gin_clean_pending_list_3))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int64(0)
															} else {
																v106 = v5
																F_relation_close(m, v12, int32(3))
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
																	return v106
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
											return int64(0)
										} else {
											v76 = int32(1)
											F_ginInsertCleanup(m, v73, v76, v76, v76, v8+int32(_a_F_gin_clean_pending_list_5))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int64(0)
											} else {
												v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[8]))))
												v106 = v83
												F_relation_close(m, v12, int32(3))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int64(0)
												} else {
													m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
													return v106
												}
											}
										}
									} else {
										v86 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int64(0)
										} else {
											if v86 == int32(0) {
												v106 = v5
												F_relation_close(m, v12, int32(3))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int64(0)
												} else {
													m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
													return v106
												}
											} else {
												F_errcode(m, int32(325))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int64(0)
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93 + int32(4)
													F_errmsg(m, int32(_a_F_gin_clean_pending_list_6), v8)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1089), int32(_a_F_gin_clean_pending_list_3))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int64(0)
														} else {
															v106 = v5
															F_relation_close(m, v12, int32(3))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return int64(0)
															} else {
																m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
																return v106
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
							return int64(0)
						} else {
							if v47 == int32(0) {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
								F_aclcheck_error(m, int32(2), int32(20), v53+int32(4))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int64(0)
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
											return int64(0)
										} else {
											v76 = int32(1)
											F_ginInsertCleanup(m, v73, v76, v76, v76, v8+int32(_a_F_gin_clean_pending_list_5))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int64(0)
											} else {
												v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[8]))))
												v106 = v83
												F_relation_close(m, v12, int32(3))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int64(0)
												} else {
													m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
													return v106
												}
											}
										}
									} else {
										v86 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int64(0)
										} else {
											if v86 == int32(0) {
												v106 = v5
												F_relation_close(m, v12, int32(3))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int64(0)
												} else {
													m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
													return v106
												}
											} else {
												F_errcode(m, int32(325))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int64(0)
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93 + int32(4)
													F_errmsg(m, int32(_a_F_gin_clean_pending_list_6), v8)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1089), int32(_a_F_gin_clean_pending_list_3))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int64(0)
														} else {
															v106 = v5
															F_relation_close(m, v12, int32(3))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return int64(0)
															} else {
																m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
																return v106
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
										return int64(0)
									} else {
										v76 = int32(1)
										F_ginInsertCleanup(m, v73, v76, v76, v76, v8+int32(_a_F_gin_clean_pending_list_5))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int64(0)
										} else {
											v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_gin_clean_pending_list[8]))))
											v106 = v83
											F_relation_close(m, v12, int32(3))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
												return v106
											}
										}
									}
								} else {
									v86 = F_errstart(m, int32(14), int32(0))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int64(0)
									} else {
										if v86 == int32(0) {
											v106 = v5
											F_relation_close(m, v12, int32(3))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
												return v106
											}
										} else {
											F_errcode(m, int32(325))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int64(0)
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93 + int32(4)
												F_errmsg(m, int32(_a_F_gin_clean_pending_list_6), v8)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1089), int32(_a_F_gin_clean_pending_list_3))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int64(0)
													} else {
														v106 = v5
														F_relation_close(m, v12, int32(3))
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return int64(0)
														} else {
															m.G0 = v8 + int32(_a_F_gin_clean_pending_list_0)
															return v106
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
			v117 = m.ExcPending
			if v117 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_gin_clean_pending_list_7), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int64(0)
					} else {
						F_errhint(m, int32(_a_F_gin_clean_pending_list_8), int32(0))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_gin_clean_pending_list_2), int32(1044), int32(_a_F_gin_clean_pending_list_3))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
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
func F_gin_compare_jsonb(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v18&int32(1) != 0 {
				v21 = int32(1)
			} else {
				v21 = int32(4)
			}
			v22 = int32(1)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v26 = v24 & v22
			if v26 != 0 {
				v27 = v22
			} else {
				v27 = int32(4)
			}
			if v24 == int32(1) {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v34 == int32(18) {
					v37 = int32(16)
				} else {
					v37 = int32(0)
				}
				if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v44 = int32(4)
				} else {
					v44 = v37
				}
				v55 = v44
			} else {
				v45 = int32(1)
				if v26 != 0 {
					v55 = int32(base.Ui32(v24)>>(uint(v45)%32)) - v45
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v18 == int32(1) {
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
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
				v85 = v72
			} else {
				v73 = int32(1)
				if v18&v73 != 0 {
					v85 = int32(base.Ui32(v18)>>(uint(v73)%32)) - v73
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = F_varstr_cmp(m, v27+v9, v55, v16+v21, v85, int32(950))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int64(0)
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v89 != v9 {
					F_pfree(m, v9)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int64(0)
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v93 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_s(v87)
							}
						} else {
							return base.I64_extend_i32_s(v87)
						}
					}
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v93 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(v87)
						}
					} else {
						return base.I64_extend_i32_s(v87)
					}
				}
			}
		}
	}
}
func F_gin_consistent_hstore(m *base.Module, l0 int32) int64 {
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
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int64
	_ = v75
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = base.I32_wrap_i64(v14) & int32(_a_F_gin_consistent_hstore_0)
	switch v17 - int32(7) {
	case 0:
		goto L5
	default:
		goto L3
	case 2, 3:
		goto L2
	case 4:
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v75
L2:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v71)
	v75 = int64(1)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v38)
	v41 = int64(1)
	if v12 <= v38 {
		v75 = v41
		goto L1
	} else {
		goto L13
	}
L5:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v20)
	v22 = int32(0)
	v23 = int64(1)
	if v12 <= v22 {
		v75 = v23
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v26 = v22
	goto L7
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v13))))
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v75 = int64(0)
	goto L1
L9:
	;
	v35 = v26 + int32(1)
	if v12 != v35 {
		v26 = v35
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v75 = v23
	goto L1
L13:
	;
	v44 = v38
	goto L14
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v13))))
	if v51 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v75 = int64(0)
	goto L1
L16:
	;
	v53 = v44 + int32(1)
	if v12 != v53 {
		v44 = v53
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	v75 = v41
	goto L1
L20:
	;
	return int64(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
	F_errmsg_internal(m, int32(_a_F_gin_consistent_hstore_1), v9)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_gin_consistent_hstore_2), int32(209), int32(_a_F_gin_consistent_hstore_3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_jsonb_path(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v129 int64
	_ = v129
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v22 = v20 & int32(268435455)
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v12 - int32(-64)
	return v129
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
	v129 = int64(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v30 = v22 << (uint(int32(1)) % 32)
	v31 = F_palloc_mul(m, int32(8), v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
	v37 = F_JsonbIteratorInit(m, v15+int32(4))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v37
	v42 = v10 + int32(-48)
	v45 = v30
	v46 = v31
	v47 = int32(0)
	goto L10
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v47
	v129 = base.I64_extend_i32_u(v46)
	goto L3
L10:
	;
	v56 = F_JsonbIteratorNext(m, v10+int32(-4), v10+int32(-40), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L34
	}
L12:
	;
	goto L11
L13:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	F_pfree(m, v42)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L30
	}
L14:
	;
	F_JsonbHashScalarValue(m, v10+int32(-40), v42)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	if v56 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	switch v56 {
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
	F_JsonbHashScalarValue(m, v10+int32(-40), v42)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v61 = F_palloc(m, int32(8))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v63
	v42 = v61
	goto L10
L21:
	;
	goto L10
L22:
	;
	v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42))))
	if v47 < v45 {
		v86 = v45
		v87 = v46
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v87+v47<<(uint(int32(3))%32)))) = v74
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v93
	v45 = v86
	v46 = v87
	v47 = v47 + int32(1)
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
	v78 = v45 << (uint(int32(1)) % 32)
	v79 = F_repalloc_mul(m, v46, int32(8), v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v81 = int32(8)
	v84 = F_palloc_mul(m, v81, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v86 = v78
	v87 = v79
	goto L23
L29:
	;
	v86 = v81
	v87 = v84
	goto L23
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v100 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v101
	goto L33
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(0)
	goto L33
L33:
	;
	v42 = v97
	goto L10
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v56
	F_errmsg_internal(m, int32(_a_F_gin_extract_jsonb_path_0), v12)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_gin_extract_jsonb_path_1), int32(1171), int32(_a_F_gin_extract_jsonb_path_2))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
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
func F_gin_extract_query_bit(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_bit_0), int32(_a_F_gin_extract_query_bit_1), int32(0), int32(_a_F_gin_extract_query_bit_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_int2(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_int2_0), int32(_a_F_gin_extract_query_int2_1), int32(_a_F_gin_extract_query_int2_2), int32(_a_F_gin_extract_query_int2_3))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_int8(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_int8_0), int32(_a_F_gin_extract_query_int8_1), int32(_a_F_gin_extract_query_int8_2), int32(_a_F_gin_extract_query_int8_3))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_interval(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_interval_0), int32(_a_F_gin_extract_query_interval_1), int32(0), int32(_a_F_gin_extract_query_interval_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_numeric(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_numeric_0), int32(_a_F_gin_extract_query_numeric_1), int32(0), int32(_a_F_gin_extract_query_numeric_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_text(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_text_0), int32(_a_F_gin_extract_query_text_1), int32(_a_F_gin_extract_query_text_2), int32(_a_F_gin_extract_query_text_3))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_timetz(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_timetz_0), int32(_a_F_gin_extract_query_timetz_1), int32(0), int32(_a_F_gin_extract_query_timetz_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_uuid(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_uuid_0), int32(_a_F_gin_extract_query_uuid_1), int32(0), int32(_a_F_gin_extract_query_uuid_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_tsquery_5args(m *base.Module, l0 int32) int64 {
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
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v2 <= int32(6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_gin_extract_tsquery_5args_0), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_gin_extract_tsquery_5args_1), int32(325), int32(_a_F_gin_extract_tsquery_5args_2))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
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
			return int64(0)
		} else {
			return v20
		}
	}
}
func F_gin_extract_tsquery_oldsig(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_gin_extract_tsquery(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_gin_extract_tsvector(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v15
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v78 = int64(0)
	goto L5
L4:
	;
	v21 = F_palloc_mul(m, int32(8), v15)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v79 != v10 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if int32(0) < v23 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = v10 + int32(8)
	v31 = int32(0)
	v32 = v23
	v33 = v28
	goto L10
L8:
	;
	goto L9
L9:
	;
	v78 = base.I64_extend_i32_u(v21)
	goto L5
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v51 = F_cstring_to_text_with_len(m, v28+v32<<(uint(int32(2))%32)+int32(base.Ui32(v43)>>(uint(int32(12))%32)), int32(base.Ui32(v43)>>(uint(int32(1))%32))&int32(2047))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21+v31<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v51)
	v58 = v31 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v58 < v59 {
		v31 = v58
		v32 = v59
		v33 = v33 + int32(4)
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	F_pfree(m, v10)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	return v78
L17:
	;
	goto L16
}
func F_gin_triconsistent_jsonb(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int64
	_ = v102
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = base.I32_wrap_i64(v13)
	switch v14&int32(_a_F_gin_triconsistent_jsonb_0) - int32(7) {
	case 0, 4:
		goto L3
	default:
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v102
L2:
	;
	if base.Ui32((v14-int32(9))&int32(_a_F_gin_triconsistent_jsonb_0)) <= base.Ui32(int32(1)) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	v19 = int32(0)
	v20 = int64(2)
	if v11 <= v19 {
		v102 = v20
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = v19
	goto L5
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v12))))
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v102 = int64(0)
	goto L1
L7:
	;
	v32 = v23 + int32(1)
	if v11 != v32 {
		v23 = v32
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
	v102 = v20
	goto L1
L11:
	;
	v41 = int32(0)
	if v11 <= v41 {
		v102 = v6
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.Ui32((v14-int32(15))&int32(_a_F_gin_triconsistent_jsonb_0)) <= base.Ui32(int32(1)) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v44 = v41
	goto L15
L15:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v12))))
	if base.Ui32(int32(2)) <= base.Ui32((v51-int32(1))&int32(255)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v102 = int64(2)
	goto L1
L17:
	;
	v59 = v44 + int32(1)
	if v11 != v59 {
		v44 = v59
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
	v102 = v6
	goto L1
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = F_execute_jsp_gin_node(m, v90, v12)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L26
	} else {
		goto L30
	}
L22:
	;
	if int32(0) < v11 {
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
	v76 = m.ExcPending
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v102 = int64(2)
	goto L1
L26:
	;
	return int64(0)
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14 & int32(_a_F_gin_triconsistent_jsonb_0)
	F_errmsg_internal(m, int32(_a_F_gin_triconsistent_jsonb_1), v9)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_gin_triconsistent_jsonb_2), int32(1073), int32(_a_F_gin_triconsistent_jsonb_3))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
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
	if v91 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v95 = int32(2)
	goto L33
L32:
	;
	v95 = v91
	goto L33
L33:
	;
	v102 = base.I64_extend_i32_s(v95)
	goto L1
}
func F_gin_triconsistent_jsonb_path(m *base.Module, l0 int32) int64 {
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
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = base.I32_wrap_i64(v13)
	if v14&int32(_a_F_gin_triconsistent_jsonb_path_0) == int32(7) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v71
L2:
	;
	v19 = int32(0)
	v20 = int64(2)
	if v11 <= v19 {
		v71 = v20
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if base.Ui32((v14-int32(15))&int32(_a_F_gin_triconsistent_jsonb_path_0)) <= base.Ui32(int32(1)) {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v23 = v19
	goto L6
L6:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v12))))
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v71 = int64(0)
	goto L1
L8:
	;
	v32 = v23 + int32(1)
	if v11 != v32 {
		v23 = v32
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
	v71 = v20
	goto L1
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = F_execute_jsp_gin_node(m, v63, v12)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L21
	}
L13:
	;
	if int32(0) < v11 {
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
	v49 = m.ExcPending
	if v49 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v71 = int64(2)
	goto L1
L17:
	;
	return int64(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14 & int32(_a_F_gin_triconsistent_jsonb_path_0)
	F_errmsg_internal(m, int32(_a_F_gin_triconsistent_jsonb_path_1), v9)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_gin_triconsistent_jsonb_path_2), int32(1315), int32(_a_F_gin_triconsistent_jsonb_path_3))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
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
	if v64 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v68 = int32(2)
	goto L24
L23:
	;
	v68 = v64
	goto L24
L24:
	;
	v71 = base.I64_extend_i32_s(v68)
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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
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
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v240 int32
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v265 int32
	_ = v265
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v285 int32
	_ = v285
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v335 int32
	_ = v335
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
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
	var v366 int64
	_ = v366
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v413 int32
	_ = v413
	var v415 int64
	_ = v415
	var v417 int64
	_ = v417
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
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
	v479 = m.ExcPending
	if v479 != 0 {
		goto L12
	} else {
		goto L81
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L12
	} else {
		goto L76
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
	v83 = v24 + v59<<(uint(int32(3))%32) + v62*int32(100) + int32(28)
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
	v211 = int32(1)
	v212 = v62 + v211
	v213 = base.I32_extend16_s(v212)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+6)))
	v225 = int32(4)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v215+v217*(v213-v211)<<(uint(int32(2))%32)+v225-v225)))
	goto L37
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
	*(*int32)(unsafe.Add(mBase, uint32(v111+v114<<(uint(int32(3))%32)+int32(200))+24)) = v113
	goto L16
L16:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v123 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v123 < v132 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L8
L18:
	;
	v136 = v122 + int32(28)
	v143 = v123
	v144 = v132
	v146 = v123
	goto L22
L19:
	;
	v200 = v123
	v207 = v132
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v122)+16)) = v200
	goto L17
L21:
	;
	v200 = v194
	v207 = v173
	goto L20
L22:
	;
	v152 = v136 + v132<<(uint(int32(3))%32) + v143*int32(100)
	v155 = v136 + v143<<(uint(int32(3))%32)
	if v132 != v144 {
		v173 = v144
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v194 = v132
	goto L21
L24:
	;
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+2)))
	if v174 <= int32(0) {
		v194 = v143
		goto L21
	} else {
		goto L32
	}
L25:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+7)))
	if v157 != int32(118) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v173 = v143
	goto L24
L27:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
	if v160 != int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+6)))
	if v163&int32(6) != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+2)))
	if v166 <= int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+90)))
	if v169 != int32(118) {
		v173 = v132
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+90)))
	if v177 == int32(118) {
		v194 = v143
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+5)))
	v186 = (v146 + v180 - int32(1)) & (int32(0) - v180)
	if int32(_a_F_initGinState_1) < v186 {
		v194 = v143
		goto L21
	} else {
		goto L34
	}
L34:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v155))) = uint16(v186)
	v192 = v143 + int32(1)
	if v192 != v132 {
		v143 = v192
		v144 = v173
		v146 = v186 + v174
		goto L22
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	v276 = v62 * int32(28)
	v277 = l0 + int32(1036) + v276
	v279 = F_index_getprocinfo(m, l1, v213, int32(2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L12
	} else {
		goto L46
	}
L37:
	;
	if v229 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v232 = v54 + v62*int32(28)
	v234 = F_index_getprocinfo(m, l1, v213, int32(1))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v83)+68))
	v251 = F_lookup_type_cache(m, v249, int32(64))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L12
	} else {
		goto L43
	}
L41:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v234)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+16)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v234)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+24)) = v240
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v234)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+8)) = v242
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v234)))
	*(*int64)(unsafe.Add(mBase, uint32(v232))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v232)+20)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v232)+16)) = int32(0)
	goto L42
L42:
	;
	goto L36
L43:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)+108))
	if v253 == int32(0) {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v258 = v54 + v62*int32(28)
	v260 = v251 + int32(104)
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v260)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+16)) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v260)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+24)) = v265
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v260)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+8)) = v267
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v260)))
	*(*int64)(unsafe.Add(mBase, uint32(v258))) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v258)+20)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v258)+16)) = int32(0)
	goto L45
L45:
	;
	goto L36
L46:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v279)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v277)+16)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v279)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+24)) = v285
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v279)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v277)+8)) = v287
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v279)))
	*(*int64)(unsafe.Add(mBase, uint32(v277))) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v277)+20)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v277)+16)) = int32(0)
	goto L47
L47:
	;
	v294 = v276 + (l0 + int32(1932))
	v296 = F_index_getprocinfo(m, l1, v213, int32(3))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v296)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v294)+16)) = v300
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+24)) = v302
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v296)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v294)+8)) = v304
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v296)))
	*(*int64)(unsafe.Add(mBase, uint32(v294))) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v294)+20)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v294)+16)) = int32(0)
	goto L49
L49:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313)+6)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v312+v314*(v213-int32(1))<<(uint(int32(2))%32)+int32(24)-int32(4))))
	goto L50
L50:
	;
	if v326 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v327 = v276 + (l0 + int32(3724))
	v329 = F_index_getprocinfo(m, l1, v213, int32(6))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L12
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v346)+6)))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v345+v347*(v213-int32(1))<<(uint(int32(2))%32)+int32(16)-int32(4))))
	goto L56
L54:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v329)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v327)+16)) = v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v327)+24)) = v335
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v329)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v327)+8)) = v337
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v329)))
	*(*int64)(unsafe.Add(mBase, uint32(v327))) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v327)+20)) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v327)+16)) = int32(0)
	goto L55
L55:
	;
	goto L53
L56:
	;
	if v359 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v360 = v276 + (l0 + int32(2828))
	v362 = F_index_getprocinfo(m, l1, v213, int32(4))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L12
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v377 = l0 + v276
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v377+int32(2832))))
	if v380 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v362)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v360)+16)) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v362)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+24)) = v368
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v362)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v360)+8)) = v370
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v362)))
	*(*int64)(unsafe.Add(mBase, uint32(v360))) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v360)+20)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v360)+16)) = int32(0)
	goto L61
L61:
	;
	goto L59
L62:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v377+int32(3728))))
	if v385 == int32(0) {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v391)+6)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v390+v392*(v213-int32(1))<<(uint(int32(2))%32)+int32(20)-int32(4))))
	goto L66
L65:
	;
	goto L64
L66:
	;
	if v404 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v405 = v276 + (l0 + int32(_a_F_initGinState_2))
	v407 = F_index_getprocinfo(m, l1, v213, int32(5))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L12
	} else {
		goto L70
	}
L68:
	;
	v424 = int32(0)
	goto L69
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62+(l0+int32(_a_F_initGinState_3))))) = uint8(v424)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v427+v85)))
	if v429 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_initGinState[0]))
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v407)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v405)+16)) = v411
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v407)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v405)+24)) = v413
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v407)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v405)+8)) = v415
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
	*(*int64)(unsafe.Add(mBase, uint32(v405))) = v417
	*(*int32)(unsafe.Add(mBase, uint32(v405)+20)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v405)+16)) = int32(0)
	goto L71
L71:
	;
	v424 = int32(1)
	goto L69
L72:
	;
	v431 = v429
	goto L74
L73:
	;
	v431 = int32(100)
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85+(l0+int32(_a_F_initGinState_4))))) = v431
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v212 < v433 {
		v59 = v433
		v62 = v212
		goto L6
	} else {
		goto L75
	}
L75:
	;
	goto L7
L76:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v83)+68))
	v465 = F_format_type_be(m, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v465
	F_errmsg(m, int32(_a_F_initGinState_5), v22)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L12
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_initGinState_6), int32(156), int32(_a_F_initGinState_7))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L12
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
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = int64(25769803780)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v480 + int32(4)
	F_errmsg_internal(m, int32(_a_F_initGinState_8), v22+int32(16))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L12
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_initGinState_6), int32(193), int32(_a_F_initGinState_7))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L12
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
