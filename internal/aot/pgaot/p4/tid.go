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
	v4 = F_ExecScan(m, l0, int32(811), int32(812))
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
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
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
	var v65 int32
	_ = v65
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
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
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
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
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v237 int32
	_ = v237
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+28)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v2
	v19 = int32(_a_F_TidRangeEval_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v19)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(-1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v23 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v237
L2:
	;
	v237 = int32(0)
	goto L1
L3:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+124)) = uint16(v218)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+126)) = v222
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+130)) = uint16(v224)
	v237 = int32(1)
	goto L1
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v26 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v36 = v2
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v36<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = int32(_a_F_TidRangeEval_1)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_TidRangeEval[0]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_TidRangeEval[0])) = v47
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v52 = m.T0[v51].(func(*base.Module, int32, int32, int32) int64)(m, v43, v14, v12+int32(15))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_TidRangeEval[0])) = v45
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v58 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v59 = base.I32_wrap_i64(v52)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	switch v60 {
	case 0:
		goto L12
	case 1:
		goto L13
	default:
		goto L11
	}
L11:
	;
	v206 = v36 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v206 < v207 {
		v36 = v206
		goto L6
	} else {
		goto L46
	}
L12:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)) = uint16(v133)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v135
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+8)))
	if v137 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)) = uint16(v61)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+8)))
	if v65 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v69 = v12 + int32(8)
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+2)))
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69))))
	v74 = v70 | v71<<(uint(int32(16))%32)
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	if v75 == int32(_a_F_TidRangeEval_0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v100 = v12 + int32(8)
	v102 = v12 + int32(24)
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100))))
	v108 = int32(16)
	v110 = v106 | v107<<(uint(v108)%32)
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+2)))
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102))))
	v115 = v111 | v112<<(uint(v108)%32)
	if base.Ui32(v110) < base.Ui32(v115) {
		v126 = int32(-1)
		goto L25
	} else {
		goto L26
	}
L17:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v90)
	*(*uint16)(unsafe.Add(mBase, uint32(v69)+2)) = uint16(v88)
	v94 = int32(base.Ui32(v88) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v69))) = uint16(v94)
	goto L16
L18:
	;
	v79 = v74 + int32(1)
	if v79 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v88 = v74
	v90 = v75 + int32(1)
	goto L17
L21:
	;
	v81 = v79
	goto L23
L22:
	;
	v81 = int32(-1)
	goto L23
L23:
	;
	v82 = int32(0)
	v88 = v81
	v90 = v82 - base.B2i32(v79 == v82)
	goto L17
L24:
	;
	if v126 <= int32(0) {
		goto L11
	} else {
		goto L29
	}
L25:
	;
	goto L24
L26:
	;
	if base.Ui32(v115) < base.Ui32(v110) {
		v126 = int32(1)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+4)))
	if base.Ui32(v120) < base.Ui32(v121) {
		v126 = int32(-1)
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v126 = base.B2i32(base.Ui32(v121) < base.Ui32(v120))
	goto L25
L29:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+28)) = uint16(v129)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v131
	goto L11
L30:
	;
	v141 = v12 + int32(8)
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+2)))
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141))))
	v146 = v142 | v143<<(uint(int32(16))%32)
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+4)))
	if v147 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v169 = v12 + int32(8)
	v170 = int32(16)
	v171 = v12 + v170
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+2)))
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169))))
	v179 = v175 | v176<<(uint(v170)%32)
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+2)))
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171))))
	v184 = v180 | v181<<(uint(v170)%32)
	if base.Ui32(v179) < base.Ui32(v184) {
		v195 = int32(-1)
		goto L41
	} else {
		goto L42
	}
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+4)) = uint16(v159)
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+2)) = uint16(v158)
	v163 = int32(base.Ui32(v158) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v141))) = uint16(v163)
	goto L32
L34:
	;
	if v146 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v158 = v146
	v159 = v147 - int32(1)
	goto L33
L37:
	;
	v152 = int32(-1)
	goto L39
L38:
	;
	v152 = int32(0)
	goto L39
L39:
	;
	v158 = v146 - base.B2i32(v146 != int32(0))
	v159 = v152
	goto L33
L40:
	;
	if int32(0) <= v195 {
		goto L11
	} else {
		goto L45
	}
L41:
	;
	goto L40
L42:
	;
	if base.Ui32(v184) < base.Ui32(v179) {
		v195 = int32(1)
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+4)))
	if base.Ui32(v189) < base.Ui32(v190) {
		v195 = int32(-1)
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v195 = base.B2i32(base.Ui32(v190) < base.Ui32(v189))
	goto L41
L45:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v200
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
