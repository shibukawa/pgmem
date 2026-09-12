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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v127 int32
	_ = v127
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
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
	var v216 int32
	_ = v216
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	if l1 == v5 {
		v216 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(32)
	return v216
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v203
	v208 = F_list_make1_impl(m, int32(1), v15+int32(8))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L27
	} else {
		goto L55
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
	if v175 == int32(0) {
		v216 = v174
		goto L1
	} else {
		goto L54
	}
L6:
	;
	v174 = v5
	v175 = v5
	goto L5
L7:
	;
	goto L8
L8:
	;
	v30 = v5
	v31 = v5
	v34 = v5
	goto L9
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v34<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	goto L12
L10:
	;
	v174 = v158
	v175 = v159
	goto L5
L11:
	;
	v165 = v34 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v165 < v166 {
		v30 = v158
		v31 = v159
		v34 = v165
		goto L9
	} else {
		goto L53
	}
L12:
	;
	if v41 != int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v45 == int32(0) {
		v158 = v30
		v159 = v31
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v138 = F_RestrictInfoIsTidQual(m, l0, v40, l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L27
	} else {
		goto L44
	}
L16:
	;
	v48 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 < v50 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v48
	v62 = v48
	goto L20
L18:
	;
	v127 = v48
	goto L19
L19:
	;
	if v127 == int32(0) {
		v158 = v30
		v159 = v31
		goto L11
	} else {
		goto L39
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v62<<(uint(int32(2))%32))))
	if v69 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v127 = v113
	goto L19
L22:
	;
	if v110 == int32(0) {
		v158 = v30
		v159 = v31
		goto L11
	} else {
		goto L36
	}
L23:
	;
	v99 = F_RestrictInfoIsTidQual(m, l0, v69, l2)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L27
	} else {
		goto L33
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v72 != int32(21) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v75 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v79 = F_TidQualFromRestrictInfoList(m, l0, v76, l2, v15+int32(31))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	if v83 != int32(1) {
		v110 = v79
		goto L22
	} else {
		goto L29
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errmsg_internal(m, int32(358416), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(497587), int32(318), int32(76101))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L27
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
	if v99 == int32(0) {
		v158 = v30
		v159 = v31
		goto L11
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v69
	v108 = F_list_make1_impl(m, int32(1), v15+int32(12))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v110 = v108
	goto L22
L36:
	;
	v113 = F_list_concat(m, v61, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v116 = v62 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v116 < v117 {
		v61 = v113
		v62 = v116
		goto L20
	} else {
		goto L38
	}
L38:
	;
	goto L21
L39:
	;
	if v30 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v158 = v127
	v159 = v31
	goto L11
L41:
	;
	goto L42
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v136 <= v135 {
		v158 = v30
		v159 = v31
		goto L11
	} else {
		goto L43
	}
L43:
	;
	v158 = v127
	v159 = v31
	goto L11
L44:
	;
	if v138 == int32(0) {
		v158 = v30
		v159 = v31
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v142 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v31 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v145 != int32(58) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	if v148 == v149 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v151 = v31
	goto L52
L51:
	;
	v151 = v40
	goto L52
L52:
	;
	v158 = v30
	v159 = v151
	goto L11
L53:
	;
	goto L10
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v175
	v202 = v15 + int32(16)
	goto L3
L55:
	;
	v216 = v208
	goto L1
}
func F_TidRangeEval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
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
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v250 int32
	_ = v250
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+28)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v2
	v20 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)) = uint16(v20)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(-1)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v24 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v250
L2:
	;
	v250 = int32(0)
	goto L1
L3:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v230
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+124)) = uint16(v232)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+126)) = v234
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+130)) = uint16(v236)
	v250 = int32(1)
	goto L1
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v27 <= int32(0) {
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
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v36<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v46 = int32(4515120)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v49
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v45, v15, v13+int32(15))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v47
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v60 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	switch v61 {
	case 0:
		goto L12
	case 1:
		goto L13
	default:
		goto L11
	}
L11:
	;
	v217 = v36 + int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v217 < v218 {
		v36 = v217
		goto L6
	} else {
		goto L49
	}
L12:
	;
	v137 = v13 + int32(12)
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v137))) = uint16(v138)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v140
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)))
	if v142 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v63 = v13 + int32(12)
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v63))) = uint16(v64)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)))
	if v68 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v72 = v13 + int32(8)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+2)))
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72))))
	v77 = v73 | v74<<(uint(int32(16))%32)
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	if v78 == int32(65535) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v103 = v13 + int32(8)
	v105 = v13 + int32(24)
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+2)))
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103))))
	v111 = int32(16)
	v113 = v109 | v110<<(uint(v111)%32)
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+2)))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105))))
	v118 = v114 | v115<<(uint(v111)%32)
	if base.Ui32(v113) < base.Ui32(v118) {
		v129 = int32(-1)
		goto L25
	} else {
		goto L26
	}
L17:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v93)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+2)) = uint16(v91)
	v97 = int32(base.Ui32(v91) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v72))) = uint16(v97)
	goto L16
L18:
	;
	v82 = v77 + int32(1)
	if v82 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v91 = v77
	v93 = v78 + int32(1)
	goto L17
L21:
	;
	v84 = v82
	goto L23
L22:
	;
	v84 = int32(-1)
	goto L23
L23:
	;
	v85 = int32(0)
	v91 = v84
	v93 = v85 - base.B2i32(v82 == v85)
	goto L17
L24:
	;
	if v129 <= int32(0) {
		goto L11
	} else {
		goto L29
	}
L25:
	;
	goto L24
L26:
	;
	if base.Ui32(v118) < base.Ui32(v113) {
		v129 = int32(1)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)))
	if base.Ui32(v123) < base.Ui32(v124) {
		v129 = int32(-1)
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v129 = base.B2i32(base.Ui32(v124) < base.Ui32(v123))
	goto L25
L29:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+28)) = uint16(v132)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v134
	goto L11
L30:
	;
	v146 = v13 + int32(8)
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+2)))
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146))))
	v151 = v147 | v148<<(uint(int32(16))%32)
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+4)))
	if v152 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v178 = v13 + int32(8)
	v179 = int32(16)
	v180 = v13 + v179
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+2)))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178))))
	v188 = v184 | v185<<(uint(v179)%32)
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+2)))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180))))
	v193 = v189 | v190<<(uint(v179)%32)
	if base.Ui32(v188) < base.Ui32(v193) {
		v204 = int32(-1)
		goto L44
	} else {
		goto L45
	}
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v146)+4)) = uint16(v166)
	*(*uint16)(unsafe.Add(mBase, uint32(v146)+2)) = uint16(v165)
	v171 = int32(base.Ui32(v165) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v146))) = uint16(v171)
	goto L32
L34:
	;
	if v151 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v165 = v151
	v166 = v152 - int32(1)
	goto L33
L37:
	;
	v157 = int32(-1)
	goto L39
L38:
	;
	v157 = int32(0)
	goto L39
L39:
	;
	v159 = v151 - int32(1)
	if base.Ui32(v159) <= base.Ui32(v151) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v162 = v159
	goto L42
L41:
	;
	v162 = int32(0)
	goto L42
L42:
	;
	v165 = v162
	v166 = v157
	goto L33
L43:
	;
	if int32(0) <= v204 {
		goto L11
	} else {
		goto L48
	}
L44:
	;
	goto L43
L45:
	;
	if base.Ui32(v193) < base.Ui32(v188) {
		v204 = int32(1)
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+4)))
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+4)))
	if base.Ui32(v198) < base.Ui32(v199) {
		v204 = int32(-1)
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v204 = base.B2i32(base.Ui32(v199) < base.Ui32(v198))
	goto L44
L48:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137))))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)) = uint16(v207)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v209
	goto L11
L49:
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
