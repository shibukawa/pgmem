package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GinInitBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v2 = l1
	if l0 < int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_GinInitBuffer[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+(l0^int32(-1))<<(uint(int32(2))%32))))
		v20 = v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_GinInitBuffer[1]))
		v20 = v14 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	if v20&int32(3) != 0 {
	} else {
	}
	v47 = F___memset(m, v20, int32(0), int32(_a_F_GinInitBuffer_0))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v20)+10)) = int32(_a_F_GinInitBuffer_1)
	v53 = int32(_a_F_GinInitBuffer_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+18)) = uint16(v53)
	v59 = int32(_a_F_GinInitBuffer_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v59)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)) = uint16(v59)
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v63 = v20 + v62
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v2)
	return
}
func F_freeGinBtreeStack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v4 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v8 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_ReleaseBuffer(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_pfree(m, v4)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	if v7 != 0 {
		v4 = v7
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L5
}
func F_ginBuildCallbackParallel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v14 = m.G0
	v15 = int32(16)
	v16 = v14 - v15
	m.G0 = v16
	v18 = int32(_a_F_ginBuildCallbackParallel_0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0])) = v21
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v32 = v28 | v29<<(uint(v15)%32)
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[2]))))
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[3]))))
	v37 = v33 | v34<<(uint(v15)%32)
	if base.Ui32(v32) < base.Ui32(v37) {
		v48 = int32(-1)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v48 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	if base.Ui32(v37) < base.Ui32(v32) {
		v48 = int32(1)
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[4]))))
	if base.Ui32(v42) < base.Ui32(v43) {
		v48 = int32(-1)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v48 = base.B2i32(base.Ui32(v43) < base.Ui32(v42))
	goto L2
L6:
	;
	F_ginFlushBuildState(m, l5, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[3]))) = v53
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[4]))) = uint16(v55)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if int32(0) < v58 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v68 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[5])))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[6])))
	if base.Ui32(v130<<(uint(int32(10))%32)) <= base.Ui32(v129) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v68))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2+v68<<(uint(int32(2))%32))))
	v83 = int32(_a_F_ginBuildCallbackParallel_0)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[7])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0])) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[8])))
	v90 = v68 + int32(1)
	v92 = v90 & int32(_a_F_ginBuildCallbackParallel_1)
	v97 = F_ginExtractEntries(m, v88, v92, v82, v78, v16+int32(8), v16+int32(12))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0])) = v84
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	F_ginInsertBAEntries(m, l5+int32(_a_F_ginBuildCallbackParallel_2), l1, v92, v97, v101, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[9])))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[9]))) = base.F64_add(v105, base.F64_convert_i32_s(v106))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_ginBuildCallbackParallel[7])))
	F_MemoryContextReset(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v90 < v114 {
		v68 = v90
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	F_ginFlushBuildState(m, l5, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginBuildCallbackParallel[0])) = v19
	m.G0 = v16 + int32(16)
	return
L23:
	;
	goto L22
}
func F_ginCompressPostingList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v112 int64
	_ = v112
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v133 int64
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	var v149 int32
	_ = v149
	var v157 int64
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = l2 & int32(-2)
	v21 = F_palloc(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v27
	v30 = v21 + int32(8)
	v31 = int32(1)
	if l1 < int32(2) {
		v189 = v30
		v193 = v31
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v201 = v189 - v30
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v201)
	v204 = v201 & int32(_a_F_ginCompressPostingList_0)
	if v204 != (v204+int32(1))&int32(_a_F_ginCompressPostingList_1) {
		goto L29
	} else {
		goto L30
	}
L4:
	;
	v34 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v35 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v21)+2)))
	v38 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	v46 = v30
	v50 = v31
	v56 = v34 | (v35<<(uint(int64(11))%64) | v38<<(uint(int64(27))%64))
	goto L5
L5:
	;
	v60 = l0 + v50*int32(6)
	v61 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)))
	v64 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v68 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)))
	v69 = v61<<(uint(int64(11))%64) | v64<<(uint(int64(27))%64) | v68
	v70 = v69 - v56
	v71 = v20 + v21 - v46
	if int32(7) <= v71 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v189 = v183
	v193 = l1
	goto L3
L7:
	;
	v185 = v50 + int32(1)
	if v185 != l1 {
		v46 = v183
		v50 = v185
		v56 = v69
		goto L5
	} else {
		goto L28
	}
L8:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v70) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v118 = v17 + int32(9)
	if base.Ui64(int64(128)) <= base.Ui64(v70) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v78 = v46
	v88 = v70
	goto L14
L12:
	;
	v102 = v46
	v112 = v70
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v112)
	v183 = v102 + int32(1)
	goto L7
L14:
	;
	v92 = base.I32_wrap_i64(v88) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v92)
	v95 = v78 + int32(1)
	v99 = int64(base.Ui64(v88) >> (uint(int64(7)) % 64))
	if base.Ui64(int64(16383)) < base.Ui64(v88) {
		v78 = v95
		v88 = v99
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v102 = v95
	v112 = v99
	goto L13
L16:
	;
	goto L15
L17:
	;
	v125 = v118
	v133 = v70
	goto L20
L18:
	;
	v149 = v118
	v157 = v70
	goto L19
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v157)
	v162 = v149 - v17 - int32(8)
	if v71 < v162 {
		v189 = v46
		v193 = v50
		goto L3
	} else {
		goto L23
	}
L20:
	;
	v137 = base.I32_wrap_i64(v133) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v137)
	v140 = v125 + int32(1)
	v144 = int64(base.Ui64(v133) >> (uint(int64(7)) % 64))
	if base.Ui64(int64(16383)) < base.Ui64(v133) {
		v125 = v140
		v133 = v144
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v149 = v140
	v157 = v144
	goto L19
L22:
	;
	goto L21
L23:
	;
	if v162 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v183 = v167 + v162
	goto L7
L25:
	;
	v166 = F__emscripten_memcpy_bulkmem(m, v46, v17+int32(9), v162)
	mBase = m.M
	v167 = v166
	goto L27
L26:
	;
	v167 = v46
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L6
L29:
	;
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v204+v30))) = uint8(v211)
	goto L31
L30:
	;
	goto L31
L31:
	;
	if l3 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v193
	goto L34
L33:
	;
	goto L34
L34:
	;
	m.G0 = v17 + int32(16)
	return v21
}
func F_ginFindLeafPage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
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
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
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
	v9 = F_palloc(m, int32(24))
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_ReadBuffer(m, v15, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(1)
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_CheckForSerializableConflictIn(m, v21, int32(0), v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v29 = v9
	goto L8
L7:
	;
	goto L6
L8:
	;
	v33 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+8)) = uint16(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v35 < v33 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return v29
L10:
	;
	F_LockBuffer(m, v35, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v35^int32(-1))<<(uint(int32(2))%32))))
	v53 = v45
	goto L10
L12:
	;
	goto L13
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[1]))
	v53 = v47 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L14:
	;
	if v35 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v75 = int32(1)
	if l1 != 0 {
		v102 = v75
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[0]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+(v35^int32(-1))<<(uint(int32(2))%32))))
	v74 = v66
	goto L15
L17:
	;
	goto L18
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[1]))
	v74 = v68 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L15
L19:
	;
	if l1 != 0 {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+16)))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74+v76)+6)))
	if v78&int32(2) == int32(0) {
		v102 = v75
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_LockBuffer(m, v35, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v86 = int32(2)
	F_LockBuffer(m, v35, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+16)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v90)+6)))
	if v92&int32(2) != 0 {
		v102 = v86
		goto L19
	} else {
		goto L24
	}
L24:
	;
	F_LockBuffer(m, v35, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v98 = int32(1)
	F_LockBuffer(m, v35, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v102 = v98
	goto L19
L27:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v112 != 0 {
		v171 = v53
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+16)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v103)+6)))
	if v105&int32(64) == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_ginFinishOldSplit(m, l0, v29, v102)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+16)))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v176)+6)))
	if v178&int32(2) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L32:
	;
	v115 = v53
	goto L33
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v120 == v121 {
		v171 = v115
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v171 = v156
	goto L31
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v124 = m.T0[v123].(func(*base.Module, int32, int32) int32)(m, l0, v115)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v124 == int32(0) {
		v171 = v115
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+16)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v115+v128)))
	if v130 == int32(-1) {
		v171 = v115
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v135 = F_ginStepRight(m, v133, v134, v102)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v135
	if v135 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if l1 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[0]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142+(v135^int32(-1))<<(uint(int32(2))%32))))
	v156 = v148
	goto L40
L42:
	;
	goto L43
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_ginFindLeafPage[1]))
	v156 = v150 + v135<<(uint(int32(13))%32) + int32(-8192)
	goto L40
L44:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v166 == int32(0) {
		v115 = v156
		goto L33
	} else {
		goto L48
	}
L45:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+16)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v157)+6)))
	if v159&int32(64) == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	F_ginFinishOldSplit(m, l0, v29, v102)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	goto L34
L49:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v184 = m.T0[v183].(func(*base.Module, int32, int32) int32)(m, l0, v29)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L9
L52:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	F_LockBuffer(m, v186, int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if l1 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v184
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v193 = F_ReleaseAndReadBuffer(m, v191, v192, v184)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v197 = F_palloc(m, int32(24))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v193
	goto L8
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v197)+20)) = v29
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v202 = F_ReadBuffer(m, v201, v184)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+4)) = v202
	v29 = v197
	goto L8
}
func F_ginFinishOldSplit(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15 + int32(4)
			F_errmsg_internal(m, int32(_a_F_ginFinishOldSplit_0), v8)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ginFinishOldSplit_1), int32(783), int32(_a_F_ginFinishOldSplit_2))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					if l2 == int32(1) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						F_LockBuffer(m, v32, int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							F_LockBuffer(m, v36, int32(2))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if v40 < int32(0) {
									v44 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishOldSplit[0]))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(v40^int32(-1))<<(uint(int32(2))%32))))
									v58 = v50
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishOldSplit[1]))
									v58 = v52 + v40<<(uint(int32(13))%32) + int32(-8192)
								}
								v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v58)+6)))
								if v61&int32(64) == int32(0) {
									m.G0 = v8 + int32(16)
									return
								} else {
									v67 = int32(0)
									F_ginFinishSplit(m, l0, l1, v67, v67)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v67 = int32(0)
						F_ginFinishSplit(m, l0, l1, v67, v67)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			if l2 == int32(1) {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				F_LockBuffer(m, v32, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					F_LockBuffer(m, v36, int32(2))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v40 < int32(0) {
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishOldSplit[0]))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(v40^int32(-1))<<(uint(int32(2))%32))))
							v58 = v50
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, _c_F_ginFinishOldSplit[1]))
							v58 = v52 + v40<<(uint(int32(13))%32) + int32(-8192)
						}
						v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v58)+6)))
						if v61&int32(64) == int32(0) {
							m.G0 = v8 + int32(16)
							return
						} else {
							v67 = int32(0)
							F_ginFinishSplit(m, l0, l1, v67, v67)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v67 = int32(0)
				F_ginFinishSplit(m, l0, l1, v67, v67)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
func F_ginInsertBAEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
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
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
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
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
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
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v411 int32
	_ = v411
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
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
	var v480 int32
	_ = v480
	var v502 int32
	_ = v502
	v3 = l2
	v21 = m.G0
	v23 = v21 - int32(48)
	m.G0 = v23
	if int32(0) < l5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = int32(1)
	v29 = int32(base.Ui32(l5)>>(uint(v27)%32)) | l5
	v32 = int32(base.Ui32(v29)>>(uint(int32(2))%32)) | v29
	v33 = int32(4)
	v35 = int32(base.Ui32(v32)>>(uint(v33)%32)) | v32
	v38 = int32(base.Ui32(v35)>>(uint(int32(8))%32)) | v35
	v66 = int32(base.Ui32(v38)>>(uint(int32(17))%32)) | int32(base.Ui32(v38)>>(uint(v27)%32)) + v27
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v23 + int32(48)
	return
L4:
	;
	v71 = v66 - int32(1)
	if base.Ui32(v71) < base.Ui32(l5) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v85 = v71
	goto L9
L7:
	;
	goto L8
L8:
	;
	v502 = int32(1)
	if base.Ui32(v502) < base.Ui32(v66) {
		v66 = int32(base.Ui32(v66) >> (uint(v502) % 32))
		goto L4
	} else {
		goto L139
	}
L9:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3+v85<<(uint(int32(2))%32))))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v85))))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+30)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+28)) = uint8(v100)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v98
	v106 = v23 + int32(8)
	v108 = v23 + int32(7)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v111 == int32(_a_F_ginInsertBAEntries_0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L8
L11:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)))
	if v434 == int32(1) {
		goto L125
	} else {
		goto L126
	}
L12:
	;
	v164 = int32(0)
	v178 = int32(0)
	goto L14
L13:
	;
	v122 = v111
	goto L15
L14:
	;
	v179 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v179)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v110)+24))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v183 = m.T0[v182].(func(*base.Module, int32) int32)(m, v181)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L17
	} else {
		goto L27
	}
L15:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v110)+24))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v137 = m.T0[v136].(func(*base.Module, int32, int32, int32) int32)(m, v106, v122, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v164 = int32(base.Ui32(v137) >> (uint(int32(31)) % 32))
	v178 = v122
	goto L14
L17:
	;
	return
L18:
	;
	if v137 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v110)+24))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	m.T0[v142].(func(*base.Module, int32, int32, int32))(m, v122, v106, v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v137 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v145 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v145)
	v433 = v122
	goto L11
L23:
	;
	v151 = int32(4)
	goto L25
L24:
	;
	v151 = int32(8)
	goto L25
L25:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v122+v151)))
	if v153 != int32(_a_F_ginInsertBAEntries_0) {
		v122 = v153
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v178
	v186 = int32(_a_F_ginInsertBAEntries_0)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v186
	v190 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v190)
	v192 = int32(16)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v198 = v196 - v192
	if v198 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v178 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v199 = F__emscripten_memcpy_bulkmem(m, v183+v192, v23+int32(24), v198)
	mBase = m.M
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v183 == v204 {
		v399 = v204
		goto L39
	} else {
		goto L40
	}
L33:
	;
	if v164 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v183
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = v183
	goto L32
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = v183
	goto L32
L39:
	;
	v411 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v399))) = uint8(v411)
	v433 = v183
	goto L11
L40:
	;
	v207 = v183 + int32(12)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v209 != int32(1) {
		v399 = v204
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v218 = v207
	v219 = v208
	v221 = v183
	goto L42
L42:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v233 == v219 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v399 = v383
	goto L39
L44:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v382 == v383 {
		v399 = v383
		goto L39
	} else {
		goto L123
	}
L45:
	;
	v382 = v373
	goto L44
L46:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	if v236 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v304 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L49:
	;
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v239)
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v239)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v245 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v244))) = uint8(v245)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v382 = v248
	goto L44
L50:
	;
	goto L51
L51:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	if v249 != v221 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v277 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v277)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	v281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v281)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+12))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+4)) = v286
	if v286 != int32(_a_F_ginInsertBAEntries_0) {
		goto L72
	} else {
		goto L73
	}
L53:
	;
	v273 = v219
	v274 = v221
	goto L52
L54:
	;
	goto L55
L55:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+8)) = v251
	if v251 != int32(_a_F_ginInsertBAEntries_0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+12)) = v219
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v257 = v256
	goto L58
L57:
	;
	v257 = v232
	goto L58
L58:
	;
	if v221 != int32(_a_F_ginInsertBAEntries_0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+12)) = v257
	goto L61
L60:
	;
	goto L61
L61:
	;
	if v257 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+4)) = v219
	if v219 == int32(_a_F_ginInsertBAEntries_0) {
		goto L69
	} else {
		goto L70
	}
L63:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v261 == v219 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v249
	goto L62
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257)+4)) = v249
	goto L62
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257)+8)) = v249
	goto L62
L69:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertBAEntries[0]))
	v273 = v270
	v274 = int32(_a_F_ginInsertBAEntries_0)
	goto L52
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+12)) = v249
	v273 = v249
	v274 = v219
	goto L52
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+12)) = v284
	goto L74
L73:
	;
	goto L74
L74:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	if v285 != int32(_a_F_ginInsertBAEntries_0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+12)) = v291
	goto L77
L76:
	;
	goto L77
L77:
	;
	if v291 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+8)) = v284
	if v284 == int32(_a_F_ginInsertBAEntries_0) {
		v373 = v274
		goto L45
	} else {
		goto L85
	}
L79:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	if v295 == v284 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v285
	goto L78
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v285
	goto L78
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v285
	goto L78
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = v285
	v373 = v274
	goto L45
L86:
	;
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v307)
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v307)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+12))
	v313 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v312))) = uint8(v313)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	v382 = v316
	goto L44
L87:
	;
	goto L88
L88:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v317 != v221 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v345 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v341))) = uint8(v345)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v349 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v348))) = uint8(v349)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+8))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v352)+8)) = v354
	if v354 != int32(_a_F_ginInsertBAEntries_0) {
		goto L109
	} else {
		goto L110
	}
L90:
	;
	v341 = v219
	v342 = v221
	goto L89
L91:
	;
	goto L92
L92:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v317)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+4)) = v319
	if v319 != int32(_a_F_ginInsertBAEntries_0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+12)) = v219
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v325 = v324
	goto L95
L94:
	;
	v325 = v232
	goto L95
L95:
	;
	if v221 != int32(_a_F_ginInsertBAEntries_0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+12)) = v325
	goto L98
L97:
	;
	goto L98
L98:
	;
	if v325 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+8)) = v219
	if v219 == int32(_a_F_ginInsertBAEntries_0) {
		goto L106
	} else {
		goto L107
	}
L100:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	if v329 == v219 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v317
	goto L99
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325)+8)) = v317
	goto L99
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325)+4)) = v317
	goto L99
L106:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertBAEntries[0]))
	v341 = v338
	v342 = int32(_a_F_ginInsertBAEntries_0)
	goto L89
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+12)) = v317
	v341 = v317
	v342 = v219
	goto L89
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+12)) = v352
	goto L111
L110:
	;
	goto L111
L111:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v352)+12))
	if v353 != int32(_a_F_ginInsertBAEntries_0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+12)) = v359
	goto L114
L113:
	;
	goto L114
L114:
	;
	if v359 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+4)) = v352
	if v352 == int32(_a_F_ginInsertBAEntries_0) {
		v373 = v342
		goto L45
	} else {
		goto L122
	}
L116:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v363 == v352 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v353
	goto L115
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+4)) = v353
	goto L115
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+8)) = v353
	goto L115
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352)+12)) = v353
	v373 = v342
	goto L45
L123:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v382)+12))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v388 == int32(1) {
		v218 = v382 + int32(12)
		v219 = v387
		v221 = v382
		goto L42
	} else {
		goto L124
	}
L124:
	;
	goto L43
L125:
	;
	if v100&int32(255) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v480 = v85 + v66<<(uint(int32(1))%32)
	if base.Ui32(v480) < base.Ui32(l5) {
		v85 = v480
		goto L9
	} else {
		goto L138
	}
L128:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
	v443 = v442 + (v3-v27)<<(uint(v33)%32)
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+26)))
	if v444 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	v460 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v433)+24)) = uint8(v460)
	*(*int64)(unsafe.Add(mBase, uint32(v433)+32)) = int64(4294967301)
	v465 = F_palloc(m, int32(30))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L17
	} else {
		goto L136
	}
L131:
	;
	v448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v443)+24)))
	v449 = F_datumCopy(m, v98, int32(0), v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L17
	} else {
		goto L134
	}
L132:
	;
	v456 = v98
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433)+16)) = v456
	goto L130
L134:
	;
	v451 = F_GetMemoryChunkSpace(m, v449)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L17
	} else {
		goto L135
	}
L135:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v451 + v453
	v456 = v449
	goto L133
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433)+28)) = v465
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v465)+4)) = uint16(v468)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v465))) = v470
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v433)+28))
	v473 = F_GetMemoryChunkSpace(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L17
	} else {
		goto L137
	}
L137:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v473 + v475
	goto L127
L138:
	;
	goto L10
L139:
	;
	goto L5
}
func F_ginInsertItemPointers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v13 = v10 + int32(56)
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(35)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(37)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(38)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l0
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)) = uint8(v44)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+65)) = uint8(base.B2i32(l4 != v6))
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v54 = v10 + int32(74)
	v57 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(80)
	return
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v66 = v63 + v57*int32(6)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v67
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)) = uint16(v69)
	v72 = v10 + int32(12)
	v77 = F_ginFindLeafPage(m, v72, int32(0), int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
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
	F_ginInsertValue(m, v72, v77, v10, l4)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if base.Ui32(v81) < base.Ui32(v82) {
		v57 = v81
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_ginReadTuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	var v54 int32
	_ = v54
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v14 = v12 << (uint(v8) % 32)
	v18 = l0 + (v11 | v14&int32(2147418112))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v14 < int32(0) {
		if v19 != 0 {
			v24 = F_ginPostingListDecode(m, v18, v9+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				if v28 == v19 {
					v54 = v24
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
					m.G0 = v9 + int32(16)
					return v54
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v35
						F_errmsg_internal(m, int32(_a_F_ginReadTuple_0), v9)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ginReadTuple_1), int32(177), int32(_a_F_ginReadTuple_2))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
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
		} else {
			v46 = F_palloc(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v54 = v46
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
				m.G0 = v9 + int32(16)
				return v54
			}
		}
	} else {
		v49 = v19 * int32(6)
		v50 = F_palloc(m, v49)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v49 != 0 {
				v52 = F__emscripten_memcpy_bulkmem(m, v50, v18, v49)
				mBase = m.M
			} else {
			}
			v54 = v50
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
			m.G0 = v9 + int32(16)
			return v54
		}
	}
}
func F_ginScanBeginPostingTree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	v4 = int32(0)
	v6 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v6
	v9 = l0 + int32(44)
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v6
	v13 = l0 + int32(52)
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(35)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(37)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(38)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(42)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+60)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v37)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v37)
	v44 = F_ginFindLeafPage(m, l0, v37, v4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		return int32(0)
	} else {
		return v44
	}
}
func F_ginUpdateStats(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v13 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_LockBuffer(m, v13, int32(2))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v13 < int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[0]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v13^int32(-1))<<(uint(int32(2))%32))))
				v35 = v27
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[1]))
				v35 = v29 + v13<<(uint(int32(13))%32) + int32(-8192)
			}
			v36 = int32(_a_F_ginUpdateStats_0)
			v38 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
			*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v38 + int32(1)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v42
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v44
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v46
			v48 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
			v49 = int32(80)
			*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)) = uint16(v49)
			*(*int64)(unsafe.Add(mBase, uint32(v35-int32(-64)))) = v48
			F_MarkBufferDirty(m, v13)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+118)))
				if v57 != int32(112) {
					F_UnlockReleaseBuffer(m, v13)
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return
					} else {
						v115 = int32(_a_F_ginUpdateStats_0)
						v117 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v117 - int32(1)
						m.G0 = v10 + int32(96)
						return
					}
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[3]))
					if v61 <= int32(0) {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						if v64 != 0 {
							F_UnlockReleaseBuffer(m, v13)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								v115 = int32(_a_F_ginUpdateStats_0)
								v117 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v117 - int32(1)
								m.G0 = v10 + int32(96)
								return
							}
						} else {
							if l2 != 0 {
								F_UnlockReleaseBuffer(m, v13)
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return
								} else {
									v115 = int32(_a_F_ginUpdateStats_0)
									v117 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
									*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v117 - int32(1)
									m.G0 = v10 + int32(96)
									return
								}
							} else {
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v65 == int32(0) {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v68
									v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v70
									*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = int64(-1)
									v77 = v35 + int32(24)
									v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v78
									v80 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v80
									v82 = *(*int64)(unsafe.Add(mBase, uint32(v77)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v82
									v84 = *(*int64)(unsafe.Add(mBase, uint32(v77)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v84
									v88 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = v88
									v90 = *(*int64)(unsafe.Add(mBase, uint32(v77)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v90
									v92 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v92
									F_XLogBeginInsert(m)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										F_XLogRegisterData(m, v10+int32(8), int32(88))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											F_XLogRegisterBuffer(m, int32(0), v13, int32(14))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												v107 = F_XLogInsert(m, int32(13), int32(96))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v35))) = base.I64_rotr(v107, int64(32))
													F_UnlockReleaseBuffer(m, v13)
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return
													} else {
														v115 = int32(_a_F_ginUpdateStats_0)
														v117 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
														*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v117 - int32(1)
														m.G0 = v10 + int32(96)
														return
													}
												}
											}
										}
									}
								} else {
									F_UnlockReleaseBuffer(m, v13)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										v115 = int32(_a_F_ginUpdateStats_0)
										v117 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
										*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v117 - int32(1)
										m.G0 = v10 + int32(96)
										return
									}
								}
							}
						}
					} else {
						if l2 != 0 {
							F_UnlockReleaseBuffer(m, v13)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								v115 = int32(_a_F_ginUpdateStats_0)
								v117 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v117 - int32(1)
								m.G0 = v10 + int32(96)
								return
							}
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v68
							v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v70
							*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = int64(-1)
							v77 = v35 + int32(24)
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v78
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v77)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v82
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v77)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v84
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = v88
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v77)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v90
							v92 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v92
							F_XLogBeginInsert(m)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								F_XLogRegisterData(m, v10+int32(8), int32(88))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									F_XLogRegisterBuffer(m, int32(0), v13, int32(14))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return
									} else {
										v107 = F_XLogInsert(m, int32(13), int32(96))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v35))) = base.I64_rotr(v107, int64(32))
											F_UnlockReleaseBuffer(m, v13)
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return
											} else {
												v115 = int32(_a_F_ginUpdateStats_0)
												v117 = *(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2]))
												*(*int32)(unsafe.Add(mBase, _c_F_ginUpdateStats[2])) = v117 - int32(1)
												m.G0 = v10 + int32(96)
												return
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
func F_gin_btree_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2))) = uint8(v3)
	return int32(1)
}
func F_gin_cmp_prefix(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v16 = v9 + int32(1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v21 = int32(1)
			v22 = v20 & v21
			if v20 == v21 {
				v25 = int32(4)
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v27&int32(254) == int32(2) {
					v36 = v25
				} else {
					v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
				}
				if v27 == int32(1) {
					v39 = v25
				} else {
					v39 = v36
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v22 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v22 != 0 {
				v51 = v16
			} else {
				v51 = v9 + int32(4)
			}
			v52 = int32(1)
			v53 = v18 + v52
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			v58 = v56 & v52
			if v58 != 0 {
				v59 = v53
			} else {
				v59 = v18 + int32(4)
			}
			if v56 == int32(1) {
				v62 = int32(4)
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v64&int32(254) == int32(2) {
					v73 = v62
				} else {
					v73 = base.B2i32(v64 == int32(18)) << (uint(v62) % 32)
				}
				if v64 == int32(1) {
					v76 = v62
				} else {
					v76 = v73
				}
				v87 = v76
			} else {
				v77 = int32(1)
				if v58 != 0 {
					v87 = int32(base.Ui32(v56)>>(uint(v77)%32)) - v77
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v87 = int32(base.Ui32(v81)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v50 == int32(0) {
				v113 = int32(0)
			} else {
				if v87 == int32(0) {
					v113 = base.B2i32(int32(0) < v50)
				} else {
					if base.Ui32(v50) < base.Ui32(v87) {
						v102 = v50
					} else {
						v102 = v87
					}
					v103 = F_memcmp(m, v51, v59, v102)
					mBase = m.M
					if v103 != 0 {
						v113 = v103
					} else {
						v113 = base.B2i32(v87 < v50)
					}
				}
			}
			v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v114 != v9 {
				F_pfree(m, v9)
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v118 != v18 {
						F_pfree(m, v18)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							if v113 < int32(0) {
								v125 = int32(1)
							} else {
								v125 = v113
							}
							return v125
						}
					} else {
						if v113 < int32(0) {
							v125 = int32(1)
						} else {
							v125 = v113
						}
						return v125
					}
				}
			} else {
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v118 != v18 {
					F_pfree(m, v18)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						if v113 < int32(0) {
							v125 = int32(1)
						} else {
							v125 = v113
						}
						return v125
					}
				} else {
					if v113 < int32(0) {
						v125 = int32(1)
					} else {
						v125 = v113
					}
					return v125
				}
			}
		}
	}
}
func F_gin_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
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
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	switch int32(base.Ui32(v14-int32(32)) >> (uint(int32(4)) % 32)) {
	case 0:
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		if v21&int32(2) != 0 {
			v24 = int32(84)
		} else {
			v24 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v24
		if v21&int32(1) != 0 {
			v30 = int32(84)
		} else {
			v30 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v30
		F_appendStringInfo(m, l0, int32(_a_F_gin_desc_0), v10-int32(-64))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v37&int32(2) == int32(0) {
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
				v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
				v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)))
				v46 = int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v44 | v45<<(uint(v46)%32)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v42 | v43<<(uint(v46)%32)
				F_appendStringInfo(m, l0, int32(_a_F_gin_desc_1), v10+int32(48))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+105)))
					if v62 == int32(1) {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+106)))
						if v65 == int32(1) {
							F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_2))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						} else {
							F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_3))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						}
					} else {
						v74 = int32(0)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
						if v78 < v74 {
							v100 = v74
						} else {
							v84 = v77 + int32(76)
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
							if v85 != int32(1) {
								v100 = v74
							} else {
								v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+43)))
								if v88 == int32(0) {
									v100 = v74
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
									v100 = v98
								}
							}
						}
						v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
						if v104&int32(1) == int32(0) {
							v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
							if v111 != 0 {
								v112 = int32(84)
							} else {
								v112 = int32(70)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v112
							F_appendStringInfo(m, l0, int32(_a_F_gin_desc_4), v10+int32(16))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						} else {
							if v104&int32(2) != 0 {
								F_desc_recompress_leaf(m, l0, v100)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									m.G0 = v10 + int32(112)
									return
								}
							} else {
								v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
								v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
								v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+8)))
								v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+6)))
								v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+10)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v127
								v129 = int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v125 | v126<<(uint(v129)%32)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v123 | v124<<(uint(v129)%32)
								F_appendStringInfo(m, l0, int32(_a_F_gin_desc_5), v10+int32(32))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return
								} else {
									m.G0 = v10 + int32(112)
									return
								}
							}
						}
					}
				}
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+105)))
				if v62 == int32(1) {
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+106)))
					if v65 == int32(1) {
						F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							m.G0 = v10 + int32(112)
							return
						}
					} else {
						F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_3))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							m.G0 = v10 + int32(112)
							return
						}
					}
				} else {
					v74 = int32(0)
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
					if v78 < v74 {
						v100 = v74
					} else {
						v84 = v77 + int32(76)
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
						if v85 != int32(1) {
							v100 = v74
						} else {
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+43)))
							if v88 == int32(0) {
								v100 = v74
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
								v100 = v98
							}
						}
					}
					v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
					if v104&int32(1) == int32(0) {
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
						if v111 != 0 {
							v112 = int32(84)
						} else {
							v112 = int32(70)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v112
						F_appendStringInfo(m, l0, int32(_a_F_gin_desc_4), v10+int32(16))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return
						} else {
							m.G0 = v10 + int32(112)
							return
						}
					} else {
						if v104&int32(2) != 0 {
							F_desc_recompress_leaf(m, l0, v100)
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						} else {
							v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
							v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
							v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+8)))
							v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+6)))
							v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+10)))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v127
							v129 = int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v125 | v126<<(uint(v129)%32)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v123 | v124<<(uint(v129)%32)
							F_appendStringInfo(m, l0, int32(_a_F_gin_desc_5), v10+int32(32))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						}
					}
				}
			}
		}
	case 1:
		v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+24)))
		if v144&int32(4) != 0 {
			v147 = int32(84)
		} else {
			v147 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v147
		F_appendStringInfo(m, l0, int32(_a_F_gin_desc_6), v10+int32(96))
		mBase = m.M
		v153 = m.ExcPending
		if v153 != 0 {
			return
		} else {
			v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+24)))
			if v156&int32(2) != 0 {
				v159 = int32(84)
			} else {
				v159 = int32(70)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v159
			if v156&int32(1) != 0 {
				v165 = int32(84)
			} else {
				v165 = int32(70)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v165
			F_appendStringInfo(m, l0, int32(_a_F_gin_desc_7), v10+int32(80))
			mBase = m.M
			v171 = m.ExcPending
			if v171 != 0 {
				return
			} else {
				m.G0 = v10 + int32(112)
				return
			}
		}
	default:
		m.G0 = v10 + int32(112)
		return
	case 6:
		v216 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v216
		F_appendStringInfo(m, l0, int32(_a_F_gin_desc_8), v10)
		mBase = m.M
		v220 = m.ExcPending
		if v220 != 0 {
			return
		} else {
			m.G0 = v10 + int32(112)
			return
		}
	case 7:
		v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+105)))
		if v172 == int32(1) {
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+106)))
			if v175 == int32(1) {
				F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_2))
				mBase = m.M
				v180 = m.ExcPending
				if v180 != 0 {
					return
				} else {
					m.G0 = v10 + int32(112)
					return
				}
			} else {
				F_appendStringInfoString(m, l0, int32(_a_F_gin_desc_3))
				mBase = m.M
				v183 = m.ExcPending
				if v183 != 0 {
					return
				} else {
					m.G0 = v10 + int32(112)
					return
				}
			}
		} else {
			v184 = int32(0)
			v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+72))
			if v188 < v184 {
				v210 = v184
			} else {
				v194 = v187 + int32(76)
				v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
				if v195 != int32(1) {
					v210 = v184
				} else {
					v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+43)))
					if v198 == int32(0) {
						v210 = v184
					} else {
						v208 = *(*int32)(unsafe.Add(mBase, uint32(v194)+44))
						v210 = v208
					}
				}
			}
			F_desc_recompress_leaf(m, l0, v210)
			mBase = m.M
			v215 = m.ExcPending
			if v215 != 0 {
				return
			} else {
				m.G0 = v10 + int32(112)
				return
			}
		}
	}
}
func F_gin_extract_query_bpchar(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(_a_F_gin_extract_query_bpchar_0), int32(2422))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_date(m *base.Module, l0 int32) int32 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(1443)
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
					v46 = v14 & int32(_a_F_gin_extract_query_date_0)
					switch v46 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(-2147483648)
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
							F_errmsg_internal(m, int32(_a_F_gin_extract_query_date_1), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_extract_query_date_2), int32(97), int32(_a_F_gin_extract_query_date_3))
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
func F_gin_extract_query_inet(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(_a_F_gin_extract_query_inet_0), int32(2343))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_extract_query_macaddr(m *base.Module, l0 int32) int32 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2281)
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
					v46 = v14 & int32(_a_F_gin_extract_query_macaddr_0)
					switch v46 - int32(1) {
					case 0, 1:
						v70 = F_palloc0(m, int32(6))
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
							F_errmsg_internal(m, int32(_a_F_gin_extract_query_macaddr_1), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_extract_query_macaddr_2), int32(97), int32(_a_F_gin_extract_query_macaddr_3))
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
func F_gin_extract_query_money(m *base.Module, l0 int32) int32 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2134)
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
					v46 = v14 & int32(_a_F_gin_extract_query_money_0)
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
							F_errmsg_internal(m, int32(_a_F_gin_extract_query_money_1), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_extract_query_money_2), int32(97), int32(_a_F_gin_extract_query_money_3))
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
func F_gin_extract_query_oid(m *base.Module, l0 int32) int32 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2115)
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
					v46 = v14 & int32(_a_F_gin_extract_query_oid_0)
					switch v46 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
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
							F_errmsg_internal(m, int32(_a_F_gin_extract_query_oid_1), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_extract_query_oid_2), int32(97), int32(_a_F_gin_extract_query_oid_3))
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
func F_gin_extract_query_trgm(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
		if base.Ui32(int32(11)) < base.Ui32(v22) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v295 = m.ExcPending
			if v295 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
				F_errmsg_internal(m, int32(_a_F_gin_extract_query_trgm_0), v15)
				mBase = m.M
				v301 = m.ExcPending
				if v301 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_gin_extract_query_trgm_1), int32(139), int32(_a_F_gin_extract_query_trgm_2))
					mBase = m.M
					v308 = m.ExcPending
					if v308 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = int32(1) << (uint(v22) % 32)
			if v28&int32(2690) == int32(0) {
				if v28&int32(24) != 0 {
					v182 = int32(1)
					v183 = v18 + v182
					v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					v188 = v186 & v182
					if v188 != 0 {
						v189 = v183
					} else {
						v189 = v18 + int32(4)
					}
					if v186 == int32(1) {
						v192 = int32(4)
						v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
						if v194&int32(254) == int32(2) {
							v203 = v192
						} else {
							v203 = base.B2i32(v194 == int32(18)) << (uint(v192) % 32)
						}
						if v194 == int32(1) {
							v206 = v192
						} else {
							v206 = v203
						}
						v217 = v206
					} else {
						v207 = int32(1)
						if v188 != 0 {
							v217 = int32(base.Ui32(v186)>>(uint(v207)%32)) - v207
						} else {
							v211 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
							v217 = int32(base.Ui32(v211)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v218 = F_generate_wildcard_trgm(m, v189, v217)
					mBase = m.M
					v219 = m.ExcPending
					if v219 != 0 {
						return int32(0)
					} else {
						v224 = v218
						v232 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
						v236 = int32(base.Ui32(v232)>>(uint(int32(2))%32)) - int32(5)
						v237 = int32(3)
						v238 = base.I32_div_u_s(v236, v237)
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v238
						if base.Ui32(v236) < base.Ui32(v237) {
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
							v314 = int32(0)
							m.G0 = v15 + int32(16)
							return v314
						} else {
							v242 = int32(1)
							if base.Ui32(v238) <= base.Ui32(v242) {
								v245 = v242
							} else {
								v245 = v238
							}
							v251 = F_palloc(m, v238<<(uint(int32(2))%32))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								v253 = v224 + int32(5)
								v254 = int32(0)
								for {
									v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+2)))
									v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
									v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
									*(*int32)(unsafe.Add(mBase, uint32(v251+v254<<(uint(int32(2))%32)))) = v268 | (v269<<(uint(int32(8))%32) | v272<<(uint(int32(16))%32))
									v281 = v254 + int32(1)
									if v281 != v245 {
										v253 = v253 + int32(3)
										v254 = v281
										continue
									} else {
										break
									}
									break
								}
								v314 = v251
								m.G0 = v15 + int32(16)
								return v314
							}
						}
					}
				} else {
					if int32(1)<<(uint(v22)%32)&int32(96) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v295 = m.ExcPending
						if v295 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
							F_errmsg_internal(m, int32(_a_F_gin_extract_query_trgm_0), v15)
							mBase = m.M
							v301 = m.ExcPending
							if v301 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_extract_query_trgm_1), int32(139), int32(_a_F_gin_extract_query_trgm_2))
								mBase = m.M
								v308 = m.ExcPending
								if v308 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_gin_extract_query_trgm[0]))
						v47 = F_createTrgmNFA(m, v18, v42, v15+int32(12), v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							if v47 == int32(0) {
								v284 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v26))) = v284
								*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
								v314 = v284
								m.G0 = v15 + int32(16)
								return v314
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
								v55 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(5)
								if base.Ui32(v55) < base.Ui32(int32(3)) {
									v284 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v26))) = v284
									*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
									v314 = v284
									m.G0 = v15 + int32(16)
									return v314
								} else {
									v59 = base.I32_div_u_s(v55, int32(3))
									v62 = F_palloc(m, v59<<(uint(int32(2))%32))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v41))) = v62
										v65 = int32(1)
										if base.Ui32(v59) <= base.Ui32(v65) {
											v68 = v65
										} else {
											v68 = v59
										}
										v70 = v68 & int32(7)
										v71 = int32(0)
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
										if base.Ui32(int32(24)) <= base.Ui32(v55) {
											v81 = v71
											v86 = int32(0)
											for {
												v93 = v62 + v81<<(uint(int32(2))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v93))) = v72
												*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v72
												*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v72
												*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v72
												*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v72
												*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v72
												*(*int32)(unsafe.Add(mBase, uint32(v93)+24)) = v72
												*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v72
												v102 = int32(8)
												v103 = v81 + v102
												v105 = v86 + v102
												if v105 != v68&int32(2147483640) {
													v81 = v103
													v86 = v105
													continue
												} else {
													break
												}
												break
											}
											v109 = v103
										} else {
											v109 = v71
										}
										if v70 == int32(0) {
											v224 = v47
										} else {
											v123 = v109
											v124 = v71
											for {
												*(*int32)(unsafe.Add(mBase, uint32(v62+v123<<(uint(int32(2))%32)))) = v72
												v137 = int32(1)
												v140 = v124 + v137
												if v140 != v70 {
													v123 = v123 + v137
													v124 = v140
													continue
												} else {
													break
												}
												break
											}
											v224 = v47
										}
										v232 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
										v236 = int32(base.Ui32(v232)>>(uint(int32(2))%32)) - int32(5)
										v237 = int32(3)
										v238 = base.I32_div_u_s(v236, v237)
										*(*int32)(unsafe.Add(mBase, uint32(v26))) = v238
										if base.Ui32(v236) < base.Ui32(v237) {
											*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
											v314 = int32(0)
											m.G0 = v15 + int32(16)
											return v314
										} else {
											v242 = int32(1)
											if base.Ui32(v238) <= base.Ui32(v242) {
												v245 = v242
											} else {
												v245 = v238
											}
											v251 = F_palloc(m, v238<<(uint(int32(2))%32))
											mBase = m.M
											v252 = m.ExcPending
											if v252 != 0 {
												return int32(0)
											} else {
												v253 = v224 + int32(5)
												v254 = int32(0)
												for {
													v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+2)))
													v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
													v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
													*(*int32)(unsafe.Add(mBase, uint32(v251+v254<<(uint(int32(2))%32)))) = v268 | (v269<<(uint(int32(8))%32) | v272<<(uint(int32(16))%32))
													v281 = v254 + int32(1)
													if v281 != v245 {
														v253 = v253 + int32(3)
														v254 = v281
														continue
													} else {
														break
													}
													break
												}
												v314 = v251
												m.G0 = v15 + int32(16)
												return v314
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v142 = int32(1)
				v143 = v18 + v142
				v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				v148 = v146 & v142
				if v148 != 0 {
					v149 = v143
				} else {
					v149 = v18 + int32(4)
				}
				if v146 == int32(1) {
					v152 = int32(4)
					v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
					if v154&int32(254) == int32(2) {
						v163 = v152
					} else {
						v163 = base.B2i32(v154 == int32(18)) << (uint(v152) % 32)
					}
					if v154 == int32(1) {
						v166 = v152
					} else {
						v166 = v163
					}
					v167 = F_generate_trgm(m, v149, v166)
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						v224 = v167
						v232 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
						v236 = int32(base.Ui32(v232)>>(uint(int32(2))%32)) - int32(5)
						v237 = int32(3)
						v238 = base.I32_div_u_s(v236, v237)
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v238
						if base.Ui32(v236) < base.Ui32(v237) {
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
							v314 = int32(0)
							m.G0 = v15 + int32(16)
							return v314
						} else {
							v242 = int32(1)
							if base.Ui32(v238) <= base.Ui32(v242) {
								v245 = v242
							} else {
								v245 = v238
							}
							v251 = F_palloc(m, v238<<(uint(int32(2))%32))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								v253 = v224 + int32(5)
								v254 = int32(0)
								for {
									v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+2)))
									v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
									v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
									*(*int32)(unsafe.Add(mBase, uint32(v251+v254<<(uint(int32(2))%32)))) = v268 | (v269<<(uint(int32(8))%32) | v272<<(uint(int32(16))%32))
									v281 = v254 + int32(1)
									if v281 != v245 {
										v253 = v253 + int32(3)
										v254 = v281
										continue
									} else {
										break
									}
									break
								}
								v314 = v251
								m.G0 = v15 + int32(16)
								return v314
							}
						}
					}
				} else {
					if v148 != 0 {
						v169 = int32(1)
						v173 = F_generate_trgm(m, v149, int32(base.Ui32(v146)>>(uint(v169)%32))-v169)
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int32(0)
						} else {
							v224 = v173
							v232 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
							v236 = int32(base.Ui32(v232)>>(uint(int32(2))%32)) - int32(5)
							v237 = int32(3)
							v238 = base.I32_div_u_s(v236, v237)
							*(*int32)(unsafe.Add(mBase, uint32(v26))) = v238
							if base.Ui32(v236) < base.Ui32(v237) {
								*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
								v314 = int32(0)
								m.G0 = v15 + int32(16)
								return v314
							} else {
								v242 = int32(1)
								if base.Ui32(v238) <= base.Ui32(v242) {
									v245 = v242
								} else {
									v245 = v238
								}
								v251 = F_palloc(m, v238<<(uint(int32(2))%32))
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return int32(0)
								} else {
									v253 = v224 + int32(5)
									v254 = int32(0)
									for {
										v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+2)))
										v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
										v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
										*(*int32)(unsafe.Add(mBase, uint32(v251+v254<<(uint(int32(2))%32)))) = v268 | (v269<<(uint(int32(8))%32) | v272<<(uint(int32(16))%32))
										v281 = v254 + int32(1)
										if v281 != v245 {
											v253 = v253 + int32(3)
											v254 = v281
											continue
										} else {
											break
										}
										break
									}
									v314 = v251
									m.G0 = v15 + int32(16)
									return v314
								}
							}
						}
					} else {
						v175 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						v180 = F_generate_trgm(m, v149, int32(base.Ui32(v175)>>(uint(int32(2))%32))-int32(4))
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return int32(0)
						} else {
							v224 = v180
							v232 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
							v236 = int32(base.Ui32(v232)>>(uint(int32(2))%32)) - int32(5)
							v237 = int32(3)
							v238 = base.I32_div_u_s(v236, v237)
							*(*int32)(unsafe.Add(mBase, uint32(v26))) = v238
							if base.Ui32(v236) < base.Ui32(v237) {
								*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
								v314 = int32(0)
								m.G0 = v15 + int32(16)
								return v314
							} else {
								v242 = int32(1)
								if base.Ui32(v238) <= base.Ui32(v242) {
									v245 = v242
								} else {
									v245 = v238
								}
								v251 = F_palloc(m, v238<<(uint(int32(2))%32))
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return int32(0)
								} else {
									v253 = v224 + int32(5)
									v254 = int32(0)
									for {
										v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+2)))
										v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
										v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
										*(*int32)(unsafe.Add(mBase, uint32(v251+v254<<(uint(int32(2))%32)))) = v268 | (v269<<(uint(int32(8))%32) | v272<<(uint(int32(16))%32))
										v281 = v254 + int32(1)
										if v281 != v245 {
											v253 = v253 + int32(3)
											v254 = v281
											continue
										} else {
											break
										}
										break
									}
									v314 = v251
									m.G0 = v15 + int32(16)
									return v314
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_gin_extract_tsquery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	v2 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v2
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21 <= v2 {
		v229 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v236 != v17 {
		goto L32
	} else {
		goto L33
	}
L2:
	;
	v27 = v17 + int32(8)
	v28 = F_tsquery_requires_match(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v32 = int32(0)
	goto L7
L6:
	;
	v32 = int32(2)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v34 <= int32(0) {
		v139 = v2
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v139
	v152 = v139 << (uint(int32(2)) % 32)
	v153 = F_palloc(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L20
	}
L9:
	;
	v38 = v34 & int32(3)
	v39 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v34) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = v39
	v47 = v2
	v48 = int32(0)
	goto L13
L11:
	;
	v98 = v39
	v99 = v2
	goto L12
L12:
	;
	if v38 == int32(0) {
		v139 = v99
		goto L8
	} else {
		goto L16
	}
L13:
	;
	v58 = int32(12)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v46*v58))))
	v62 = int32(1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+(v46|v62)*v58))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+(v46|int32(2))*v58))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+(v46|int32(3))*v58))))
	v91 = v47 + base.B2i32(v61 == v62) + base.B2i32(v70 == v62) + base.B2i32(v79 == v62) + base.B2i32(v88 == v62)
	v92 = int32(4)
	v93 = v46 + v92
	v95 = v48 + v92
	if v95 != v34&int32(2147483644) {
		v46 = v93
		v47 = v91
		v48 = v95
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v98 = v93
	v99 = v91
	goto L12
L15:
	;
	goto L14
L16:
	;
	v113 = v98
	v114 = v99
	v119 = v2
	goto L17
L17:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v113*int32(12)))))
	v129 = int32(1)
	v131 = v114 + base.B2i32(v128 == v129)
	v135 = v119 + v129
	if v135 != v38 {
		v113 = v113 + v129
		v114 = v131
		v119 = v135
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v139 = v131
	goto L8
L19:
	;
	goto L18
L20:
	;
	v155 = F_palloc(m, v139)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v155
	v158 = F_palloc(m, v152)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v158
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v164 = F_palloc0(m, v161<<(uint(int32(2))%32))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v166 <= int32(0) {
		v229 = v153
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v171 = int32(0)
	v172 = v166
	v178 = v2
	goto L25
L25:
	;
	v185 = v27 + v178*int32(12)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v186 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v229 = v153
	goto L1
L27:
	;
	v190 = v171 << (uint(int32(2)) % 32)
	v192 = int32(12)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	v201 = F_cstring_to_text_with_len(m, v27+v172*v192+int32(base.Ui32(v195)>>(uint(v192)%32)), v195&int32(4095))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	v217 = v171
	v218 = v172
	goto L29
L29:
	;
	v221 = v178 + int32(1)
	if v221 < v218 {
		v171 = v217
		v172 = v218
		v178 = v221
		goto L25
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153+v190))) = v201
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v171+v155))) = uint8(v205)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v207+v190))) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v164+v178<<(uint(int32(2))%32)))) = v171
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v217 = v171 + int32(1)
	v218 = v214
	goto L29
L31:
	;
	goto L26
L32:
	;
	F_pfree(m, v17)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	return v229
L35:
	;
	goto L34
}
func F_gin_extract_value_int2(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(4))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v5
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
		return v7
	}
}
func F_gin_trgm_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
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
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v235 int32
	_ = v235
	var v236 float64
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v366 float32
	_ = v366
	var v373 int32
	_ = v373
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v21)
	if base.Ui32(int32(11)) < base.Ui32(v19) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v373
L2:
	;
	v373 = base.F64_le(v236, base.F64_promote_f32(base.F32_div(v366, base.F32_convert_i32_s(v17))))
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L41
	} else {
		goto L64
	}
L4:
	;
	if int32(1)<<(uint(v19)%32)&int32(642) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v17 <= int32(0) {
		v373 = v31
		goto L1
	} else {
		goto L59
	}
L6:
	;
	v31 = int32(1)
	v33 = v31 << (uint(v19) % 32)
	if v33&int32(2072) != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v236 = F_index_strategy_get_limit(m, v19)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L41
	} else {
		goto L42
	}
L9:
	;
	if v33&int32(96) == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v17 <= int32(0) {
		v373 = v31
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v43 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v56 = F___memset(m, v53, v43, v55)
	mBase = m.M
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v60 = F___memset(m, v57, v43, v59)
	mBase = m.M
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 < v61 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v373 = v235
	goto L1
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v68 = v43
	v70 = v43
	goto L16
L14:
	;
	goto L15
L15:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v146 = v135
	v148 = v43
	goto L28
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v64+v70<<(uint(int32(2))%32))))
	v81 = v80 + v68
	if v80 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v120 = v70 + int32(1)
	if v120 != v61 {
		v68 = v81
		v70 = v120
		goto L16
	} else {
		goto L26
	}
L19:
	;
	v86 = v68
	goto L20
L20:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v86))))
	if v97 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v105 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v103+v70))) = uint8(v105)
	goto L18
L22:
	;
	v101 = v86 + int32(1)
	if v101 < v81 {
		v86 = v101
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
	goto L18
L26:
	;
	goto L17
L27:
	;
	goto L12
L28:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v137+v148<<(uint(int32(2))%32))))
	v160 = v140 + v157<<(uint(int32(3))%32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if int32(0) < v161 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v235 = int32(0)
	goto L27
L30:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v169 = int32(0)
	v171 = v146
	goto L33
L31:
	;
	v211 = v146
	goto L32
L32:
	;
	v220 = v148 + int32(1)
	if v220 < v211 {
		v146 = v211
		v148 = v220
		goto L28
	} else {
		goto L40
	}
L33:
	;
	v181 = v165 + v169<<(uint(int32(3))%32)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v182))))
	if v184 != int32(1) {
		v202 = v171
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v211 = v202
	goto L32
L35:
	;
	v205 = v169 + int32(1)
	if v205 != v161 {
		v169 = v205
		v171 = v202
		goto L33
	} else {
		goto L39
	}
L36:
	;
	v187 = int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v188 == v187 {
		v235 = v187
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v191 = v188 + v134
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v192 != 0 {
		v202 = v171
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v193)
	*(*int32)(unsafe.Add(mBase, uint32(v137+v171<<(uint(int32(2))%32)))) = v188
	v202 = v171 + v193
	goto L35
L39:
	;
	goto L34
L40:
	;
	goto L29
L41:
	;
	return int32(0)
L42:
	;
	if int32(0) < v17 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v242 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v17) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if v17 != 0 {
		v366 = float32(0)
		goto L2
	} else {
		goto L58
	}
L46:
	;
	v248 = v242
	v249 = v2
	v255 = v2
	goto L49
L47:
	;
	v273 = v242
	v274 = v2
	goto L48
L48:
	;
	v285 = v17 & int32(3)
	if v285 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v259 = v248 + v18
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+2)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+3)))
	v267 = v249 + v260 + v262 + v264 + v266
	v268 = int32(4)
	v269 = v248 + v268
	v271 = v255 + v268
	if v271 != v17&int32(2147483644) {
		v248 = v269
		v249 = v267
		v255 = v271
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v273 = v269
	v274 = v267
	goto L48
L51:
	;
	goto L50
L52:
	;
	v286 = v273
	v287 = v274
	v290 = v242
	goto L55
L53:
	;
	v306 = v274
	goto L54
L54:
	;
	v366 = base.F32_convert_i32_u(v306)
	goto L2
L55:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v18))))
	v299 = v287 + v298
	v300 = int32(1)
	v303 = v290 + v300
	if v303 != v285 {
		v286 = v286 + v300
		v287 = v299
		v290 = v303
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v306 = v299
	goto L54
L57:
	;
	goto L56
L58:
	;
	v373 = v2
	goto L1
L59:
	;
	v320 = int32(0)
	goto L60
L60:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v18))))
	if v332 != int32(1) {
		v373 = v332
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v373 = v332
	goto L1
L62:
	;
	v336 = v320 + int32(1)
	if v336 != v17 {
		v320 = v336
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v19
	F_errmsg_internal(m, int32(_a_F_gin_trgm_consistent_0), v14)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L41
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_gin_trgm_consistent_1), int32(254), int32(_a_F_gin_trgm_consistent_2))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L41
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_tsquery_triconsistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 <= int32(0) {
		v31 = int32(0)
		m.G0 = v7 + int32(16)
		return v31
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
		v18 = v9 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v20
		v24 = F_TS_execute_ternary(m, v18, v7+int32(4))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v31 = base.I32_extend8_s(v24)
			m.G0 = v7 + int32(16)
			return v31
		}
	}
}
func F_gin_xlog_cleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_gin_xlog_cleanup[0]))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_gin_xlog_cleanup[0])) = int32(0)
		return
	}
}
