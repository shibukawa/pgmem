package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecTidRangeScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(765), int32(766))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_TidQualFromRestrictInfoList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	if l1 == v5 {
		v217 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(32)
	return v217
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v21 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v203
	v208 = F_list_make1_impl(m, int32(1), v15+int32(8))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L26
	} else {
		goto L54
	}
L4:
	;
	v185 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v185)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v40
	v202 = v15 + int32(20)
	goto L3
L5:
	;
	v31 = v5
	v32 = v5
	v35 = v5
	goto L8
L6:
	;
	v175 = v5
	v176 = v5
	goto L7
L7:
	;
	if v176 == int32(0) {
		v217 = v175
		goto L1
	} else {
		goto L53
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v35<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	goto L11
L9:
	;
	v175 = v159
	v176 = v160
	goto L7
L10:
	;
	v165 = v35 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v165 < v166 {
		v31 = v159
		v32 = v160
		v35 = v165
		goto L8
	} else {
		goto L52
	}
L11:
	;
	if v41 != int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v45 == int32(0) {
		v159 = v31
		v160 = v32
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v138 = F_RestrictInfoIsTidQual(m, l0, v40, l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L26
	} else {
		goto L43
	}
L15:
	;
	v48 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 < v50 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v59 = v48
	v63 = v48
	goto L19
L17:
	;
	v125 = v48
	goto L18
L18:
	;
	if v125 == int32(0) {
		v159 = v31
		v160 = v32
		goto L10
	} else {
		goto L38
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v63<<(uint(int32(2))%32))))
	if v69 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v125 = v113
	goto L18
L21:
	;
	if v110 == int32(0) {
		v159 = v31
		v160 = v32
		goto L10
	} else {
		goto L35
	}
L22:
	;
	v99 = F_RestrictInfoIsTidQual(m, l0, v69, l2)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L26
	} else {
		goto L32
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v72 != int32(21) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v75 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v79 = F_TidQualFromRestrictInfoList(m, l0, v76, l2, v15+int32(31))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return int32(0)
L27:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	if v83 != int32(1) {
		v110 = v79
		goto L21
	} else {
		goto L28
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	F_errmsg_internal(m, int32(_a_F_TidQualFromRestrictInfoList_0), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_TidQualFromRestrictInfoList_1), int32(318), int32(_a_F_TidQualFromRestrictInfoList_2))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	if v99 == int32(0) {
		v159 = v31
		v160 = v32
		goto L10
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v69
	v108 = F_list_make1_impl(m, int32(1), v15+int32(12))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v110 = v108
	goto L21
L35:
	;
	v113 = F_list_concat(m, v59, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L26
	} else {
		goto L36
	}
L36:
	;
	v116 = v63 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v116 < v117 {
		v59 = v113
		v63 = v116
		goto L19
	} else {
		goto L37
	}
L37:
	;
	goto L20
L38:
	;
	if v31 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v159 = v125
	v160 = v32
	goto L10
L40:
	;
	goto L41
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v136 <= v135 {
		v159 = v31
		v160 = v32
		goto L10
	} else {
		goto L42
	}
L42:
	;
	v159 = v125
	v160 = v32
	goto L10
L43:
	;
	if v138 == int32(0) {
		v159 = v31
		v160 = v32
		goto L10
	} else {
		goto L44
	}
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v142 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v32 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v145 != int32(58) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	if v148 == v149 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v151 = v32
	goto L51
L50:
	;
	v151 = v40
	goto L51
L51:
	;
	v159 = v31
	v160 = v151
	goto L10
L52:
	;
	goto L9
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v176
	v202 = v15 + int32(16)
	goto L3
L54:
	;
	v217 = v208
	goto L1
}
func F_TidRangeEval(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+28)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v2
	v18 = int32(_a_F_TidRangeEval_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v18)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(-1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v22 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v232
L2:
	;
	v232 = int32(0)
	goto L1
L3:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+124)) = uint16(v214)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+126)) = v218
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+130)) = uint16(v220)
	v232 = int32(1)
	goto L1
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v35 = v2
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v35<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = int32(_a_F_TidRangeEval_1)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_TidRangeEval[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_TidRangeEval[0])) = v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v50 = m.T0[v49].(func(*base.Module, int32, int32, int32) int32)(m, v41, v13, v11+int32(15))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TidRangeEval[0])) = v43
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v56 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	switch v57 {
	case 0:
		goto L12
	case 1:
		goto L13
	default:
		goto L11
	}
L11:
	;
	v203 = v35 + int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v203 < v204 {
		v35 = v203
		goto L6
	} else {
		goto L46
	}
L12:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v130)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v132
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+8)))
	if v134 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v58)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v60
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+8)))
	if v62 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = v11 + int32(8)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66))))
	v71 = v67 | v68<<(uint(int32(16))%32)
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
	if v72 == int32(_a_F_TidRangeEval_0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v97 = v11 + int32(8)
	v99 = v11 + int32(24)
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+2)))
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	v105 = int32(16)
	v107 = v103 | v104<<(uint(v105)%32)
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+2)))
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
	v112 = v108 | v109<<(uint(v105)%32)
	if base.Ui32(v107) < base.Ui32(v112) {
		v123 = int32(-1)
		goto L25
	} else {
		goto L26
	}
L17:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v87)
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v85)
	v91 = int32(base.Ui32(v85) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v91)
	goto L16
L18:
	;
	v76 = v71 + int32(1)
	if v76 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v85 = v71
	v87 = v72 + int32(1)
	goto L17
L21:
	;
	v78 = v76
	goto L23
L22:
	;
	v78 = int32(-1)
	goto L23
L23:
	;
	v79 = int32(0)
	v85 = v78
	v87 = v79 - base.B2i32(v76 == v79)
	goto L17
L24:
	;
	if v123 <= int32(0) {
		goto L11
	} else {
		goto L29
	}
L25:
	;
	goto L24
L26:
	;
	if base.Ui32(v112) < base.Ui32(v107) {
		v123 = int32(1)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)))
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)))
	if base.Ui32(v117) < base.Ui32(v118) {
		v123 = int32(-1)
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v123 = base.B2i32(base.Ui32(v118) < base.Ui32(v117))
	goto L25
L29:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+28)) = uint16(v126)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v128
	goto L11
L30:
	;
	v138 = v11 + int32(8)
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+2)))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138))))
	v143 = v139 | v140<<(uint(int32(16))%32)
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+4)))
	if v144 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v166 = v11 + int32(8)
	v167 = int32(16)
	v168 = v11 + v167
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+2)))
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166))))
	v176 = v172 | v173<<(uint(v167)%32)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+2)))
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168))))
	v181 = v177 | v178<<(uint(v167)%32)
	if base.Ui32(v176) < base.Ui32(v181) {
		v192 = int32(-1)
		goto L41
	} else {
		goto L42
	}
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+4)) = uint16(v156)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+2)) = uint16(v155)
	v160 = int32(base.Ui32(v155) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v138))) = uint16(v160)
	goto L32
L34:
	;
	if v143 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v155 = v143
	v156 = v144 - int32(1)
	goto L33
L37:
	;
	v149 = int32(-1)
	goto L39
L38:
	;
	v149 = int32(0)
	goto L39
L39:
	;
	v155 = v143 - base.B2i32(v143 != int32(0))
	v156 = v149
	goto L33
L40:
	;
	if int32(0) <= v192 {
		goto L11
	} else {
		goto L45
	}
L41:
	;
	goto L40
L42:
	;
	if base.Ui32(v181) < base.Ui32(v176) {
		v192 = int32(1)
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+4)))
	if base.Ui32(v186) < base.Ui32(v187) {
		v192 = int32(-1)
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v192 = base.B2i32(base.Ui32(v187) < base.Ui32(v186))
	goto L41
L45:
	;
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v197
	goto L11
L46:
	;
	goto L7
}
func F_TidStoreDestroy(m *base.Module, l0 int32) {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
		F_shared_ts_free_recurse(m, v4, v7, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			F_dsa_free(m, v14, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_pfree(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_dsa_detach(m, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
		F_MemoryContextReset(m, v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			F_pfree(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_pfree(m, v4)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_MemoryContextDelete(m, v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
