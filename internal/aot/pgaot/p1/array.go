package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ArrayGetNItems(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 <= v3 {
		v80 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v80
L2:
	;
	v16 = v3
	v19 = int64(1)
	goto L6
L3:
	;
	v80 = int32(-1)
	goto L1
L4:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L11
	} else {
		goto L23
	}
L5:
	;
	v50 = base.I32_wrap_i64(v36)
	if base.Ui32(v50) < base.Ui32(int32(268435456)) {
		v80 = v50
		goto L1
	} else {
		goto L20
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1+v16<<(uint(int32(2))%32))))
	if v23 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v45 = F_errsave_start(m, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L18
	}
L8:
	;
	v27 = F_errsave_start(m, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v36 = base.I64_extend_i32_u(v23) * base.I64_extend32_s(v19)
	if base.Ui64(v36+int64(2147483648)) < base.Ui64(int64(4294967296)) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	return int32(0)
L12:
	;
	if v27 == int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v61 = int32(84)
	goto L4
L14:
	;
	v42 = v16 + int32(1)
	if v42 == l0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L7
L17:
	;
	v16 = v42
	v19 = v36
	goto L6
L18:
	;
	if v45 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v61 = int32(93)
	goto L4
L20:
	;
	v54 = F_errsave_start(m, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	if v54 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v61 = int32(100)
	goto L4
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(268435455)
	F_errmsg(m, int32(661140), v9)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	F_errsave_finish(m, int32(0), int32(490769), v61, int32(407200))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	goto L3
}
func F_array_agg_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
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
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
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
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
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
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = v10 + int32(12)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v43 = v40
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
	v40 = v36
	goto L2
L4:
	;
	v32 = int32(0)
	if v13 == v32 {
		v40 = v32
		goto L2
	} else {
		goto L14
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v18 - int32(429) {
	case 0:
		goto L7
	case 1:
		goto L6
	default:
		goto L4
	}
L6:
	;
	if v13 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v13 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v43 = int32(1)
	goto L1
L9:
	;
	goto L10
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v35 = v25
	v36 = int32(1)
	goto L3
L11:
	;
	v43 = int32(2)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+368))
	v35 = v30
	v36 = int32(2)
	goto L3
L14:
	;
	v35 = v32
	v36 = v2
	goto L3
L15:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v44 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L31
	} else {
		goto L69
	}
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v48 = v47
	goto L20
L19:
	;
	v48 = v2
	goto L20
L20:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v49 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	m.G0 = v10 + int32(16)
	return v214
L22:
	;
	if v48 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v52 != 0 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v48 != 0 {
		v214 = v48
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
	v214 = int32(0)
	goto L21
L28:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v63 = F_initArrayResultWithSize(m, v59, v60, int32(0), v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v123 <= int32(0) {
		v214 = v48
		goto L21
	} else {
		goto L47
	}
L31:
	;
	return int32(0)
L32:
	;
	v67 = int32(4487040)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if int32(0) < v72 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v76 = int32(0)
	goto L36
L34:
	;
	v110 = v72
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v68
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v110 != 0 {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	v83 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v76))))
	if v86 == v83 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v110 = v106
	goto L35
L38:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v76<<(uint(int32(2))%32))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+26)))
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+24)))
	v96 = F_datumCopy(m, v93, v94, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L31
	} else {
		goto L41
	}
L39:
	;
	v98 = v83
	goto L40
L40:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v99+v76<<(uint(int32(2))%32)))) = v98
	v105 = v76 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v105 < v106 {
		v76 = v105
		goto L36
	} else {
		goto L42
	}
L41:
	;
	v98 = v96
	goto L40
L42:
	;
	goto L37
L43:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v121
	v214 = v63
	goto L21
L44:
	;
	v119 = F__emscripten_memcpy_bulkmem(m, v117, v118, v110)
	mBase = m.M
	goto L46
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v127 = int32(4487040)
	v128 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v131
	v133 = v123 + v129
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v134 < v133 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v199 != 0 {
		goto L66
	} else {
		goto L67
	}
L49:
	;
	v136 = int32(1)
	if v133&(v133-v136) != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v161 = int32(0)
	goto L58
L52:
	;
	v144 = v136 << (uint(int32(32)-base.I32_clz(v133)) % 32)
	goto L54
L53:
	;
	v144 = v133
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v149 = F_repalloc(m, v146, v144<<(uint(int32(2))%32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L31
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v149
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v154 = F_repalloc(m, v152, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L31
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v157 <= int32(0) {
		v199 = v157
		goto L48
	} else {
		goto L57
	}
L57:
	;
	goto L51
L58:
	;
	v168 = int32(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v161))))
	if v171 == v168 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v199 = v195
	goto L48
L60:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174+v161<<(uint(int32(2))%32))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
	v180 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
	v181 = F_datumCopy(m, v178, v179, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L31
	} else {
		goto L63
	}
L61:
	;
	v183 = v168
	goto L62
L62:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v185 = int32(2)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v184+v161<<(uint(v185)%32)+v188<<(uint(v185)%32)))) = v183
	v194 = v161 + int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v194 < v195 {
		v161 = v194
		goto L58
	} else {
		goto L64
	}
L63:
	;
	v183 = v181
	goto L62
L64:
	;
	goto L59
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v133
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v128
	v214 = v48
	goto L21
L66:
	;
	v208 = F__emscripten_memcpy_bulkmem(m, v204+v205, v207, v199)
	mBase = m.M
	goto L68
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	F_errmsg_internal(m, int32(60895), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L31
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(491534), int32(609), int32(371598))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_agg_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v50 int32
	_ = v50
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L19
	} else {
		goto L76
	}
L2:
	;
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	v43 = int32(0)
	goto L2
L5:
	;
	goto L3
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v18 - int32(429) {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L5
	}
L7:
	;
	goto L12
L8:
	;
	goto L9
L9:
	;
	v43 = int32(1)
	goto L2
L12:
	;
	v43 = int32(2)
	goto L2
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v45 = F_pg_detoast_datum_packed(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	v232 = m.ExcPending
	if v232 != 0 {
		goto L19
	} else {
		goto L73
	}
L19:
	;
	return int32(0)
L20:
	;
	v49 = int32(1)
	v50 = v45 + v49
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v55 = v53 & v49
	if v55 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v56 = v50
	goto L23
L22:
	;
	v56 = v45 + int32(4)
	goto L23
L23:
	;
	if v53 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v56
	v92 = F_pq_getmsgint(m, v11+int32(16), int32(4))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L19
	} else {
		goto L35
	}
L25:
	;
	v59 = int32(4)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v61&int32(254) == int32(2) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v74 = int32(1)
	if v55 != 0 {
		v84 = int32(base.Ui32(v53)>>(uint(v74)%32)) - v74
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v70 = v59
	goto L30
L29:
	;
	v70 = base.B2i32(v61 == int32(18)) << (uint(v59) % 32)
	goto L30
L30:
	;
	if v61 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v73 = v59
	goto L33
L32:
	;
	v73 = v70
	goto L33
L33:
	;
	v84 = v73
	goto L24
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
	goto L24
L35:
	;
	v96 = F_pq_getmsgint64(m, v11+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v101 = base.I32_wrap_i64(v96)
	v102 = F_initArrayResultWithSize(m, v92, v99, int32(0), v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+16)) = v101
	v108 = F_pq_getmsgint(m, v11+int32(16), int32(2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v102)+24)) = uint16(v108)
	v113 = F_pq_getmsgbyte(m, v11+int32(16))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v102)+26)) = uint8(base.B2i32(v113 != int32(0)))
	v120 = F_pq_getmsgbyte(m, v11+int32(16))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v102)+27)) = uint8(v120)
	v125 = F_pq_getmsgbytes(m, v11+int32(16), v101)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	if v101 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+26)))
	if v130 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v128 = F__emscripten_memcpy_bulkmem(m, v127, v125, v101)
	mBase = m.M
	goto L45
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	F_pq_getmsgend(m, v11+int32(16))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L19
	} else {
		goto L72
	}
L47:
	;
	v136 = v101 << (uint(int32(2)) % 32)
	v137 = F_pq_getmsgbytes(m, v11+int32(16), v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L19
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	if v143 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v136 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L46
L52:
	;
	v140 = F__emscripten_memcpy_bulkmem(m, v139, v137, v136)
	mBase = m.M
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+20))
	v148 = F_MemoryContextAlloc(m, v146, int32(32))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L19
	} else {
		goto L58
	}
L56:
	;
	v161 = v143
	goto L57
L57:
	;
	if v96 <= int64(0) {
		goto L46
	} else {
		goto L61
	}
L58:
	;
	F_getTypeBinaryInputInfo(m, v92, v11, v148+int32(28))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	F_fmgr_info_cxt(m, v154, v148, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+16)) = v148
	v161 = v148
	goto L57
L61:
	;
	v167 = int32(0)
	goto L62
L62:
	;
	v173 = int32(0)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v167))))
	if v176 == v173 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L46
L64:
	;
	v182 = F_pq_getmsgint(m, v11+int32(16), int32(4))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L19
	} else {
		goto L67
	}
L65:
	;
	v202 = v173
	goto L66
L66:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v204+v167<<(uint(int32(2))%32)))) = v202
	v210 = v167 + int32(1)
	if base.I64_extend_i32_s(v210) < v96 {
		v167 = v210
		goto L62
	} else {
		goto L71
	}
L67:
	;
	if v182 < int32(0) {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v186-v187 < v182 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v192 + v187
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v182 + v187
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v161)+28))
	v200 = F_ReceiveFunctionCall(m, v161, v11, v198, int32(-1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	v202 = v200
	goto L66
L71:
	;
	goto L63
L72:
	;
	m.G0 = v11 + int32(32)
	return v102
L73:
	;
	F_errmsg_internal(m, int32(60895), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L19
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(491534), int32(797), int32(339896))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L19
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
	F_errcode(m, int32(50462850))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L19
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(402125), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L19
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(491534), int32(873), int32(339896))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_bitmap_copy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v84 int32
	_ = v84
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	if l4 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = int32(1) << (uint(l1&int32(7)) % 32)
	v17 = base.I32_div_s(l1, int32(8))
	v18 = l0 + v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v114)
	goto L1
L4:
	;
	v23 = v19
	v26 = l4
	v27 = v15
	v28 = v18
	goto L7
L5:
	;
	goto L6
L6:
	;
	v57 = base.I32_div_s(l3, int32(8))
	v58 = l2 + v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v60 = int32(1) << (uint(l3&int32(7)) % 32)
	v61 = v19
	v64 = l4
	v65 = v15
	v66 = v18
	v67 = v58
	v68 = v59
	goto L15
L7:
	;
	v31 = v23 | v27
	v32 = int32(1)
	v33 = v26 - v32
	v35 = v27 << (uint(v32) % 32)
	if v35 == int32(256) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v46 != int32(1) {
		v114 = v45
		v119 = v47
		goto L3
	} else {
		goto L14
	}
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v31)
	if v33 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v45 = v31
	v46 = v35
	v47 = v28
	goto L11
L11:
	;
	if base.Ui32(int32(1)) < base.Ui32(v26) {
		v23 = v45
		v26 = v33
		v27 = v46
		v28 = v47
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v41 = int32(1)
	v43 = v28 + v41
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v45 = v44
	v46 = v41
	v47 = v43
	goto L11
L13:
	;
	goto L8
L14:
	;
	goto L1
L15:
	;
	if v60&v68 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v89 == int32(1) {
		goto L1
	} else {
		goto L30
	}
L17:
	;
	v74 = v61 | v65
	goto L19
L18:
	;
	v74 = v61 & (v65 ^ int32(-1))
	goto L19
L19:
	;
	v75 = int32(1)
	v76 = v64 - v75
	v78 = v65 << (uint(v75) % 32)
	if v78 == int32(256) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v74)
	if v76 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v88 = v74
	v89 = v78
	v90 = v66
	goto L22
L22:
	;
	v92 = v60 << (uint(int32(1)) % 32)
	if v92 == int32(256) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v84 = int32(1)
	v86 = v66 + v84
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v88 = v87
	v89 = v84
	v90 = v86
	goto L22
L24:
	;
	goto L16
L25:
	;
	if v76 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v101 = v92
	v102 = v67
	v103 = v68
	goto L27
L27:
	;
	if base.Ui32(int32(1)) < base.Ui32(v64) {
		v60 = v101
		v61 = v88
		v64 = v76
		v65 = v89
		v66 = v90
		v67 = v102
		v68 = v103
		goto L15
	} else {
		goto L29
	}
L28:
	;
	v97 = int32(1)
	v98 = v67 + v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v101 = v97
	v102 = v98
	v103 = v99
	goto L27
L29:
	;
	goto L24
L30:
	;
	v114 = v88
	v119 = v90
	goto L3
}
func F_array_fill(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v6 != int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_pg_detoast_datum(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v14 == int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v18 = v17
			} else {
				v18 = int32(0)
			}
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = F_get_fn_expr_argtype(m, v19, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(64748), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(491427), int32(6067), int32(302427))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = F_array_fill_internal(m, v10, int32(0), v18, v14, v21, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						return v26
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67108994))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(300596), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491427), int32(6050), int32(302427))
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
}
func F_array_get_slice(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
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
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v182 int32
	_ = v182
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
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
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v388 int32
	_ = v388
	var v389 int64
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v461 int32
	_ = v461
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
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
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v831 int32
	_ = v831
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1009 int32
	_ = v1009
	v21 = m.G0
	v23 = v21 - int32(160)
	m.G0 = v23
	if l6 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v23 + int32(160)
	return v1009
L2:
	;
	v997 = F_palloc0(m, int32(16))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L7
	} else {
		goto L172
	}
L3:
	;
	if v32 <= v161 {
		goto L35
	} else {
		goto L36
	}
L4:
	;
	v27 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L31
	}
L7:
	;
	return int32(0)
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v32 < l1 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if base.Ui32(v32-int32(7)) < base.Ui32(int32(-6)) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v40 = v27 + int32(16)
	v42 = v32 << (uint(int32(2)) % 32)
	v43 = v40 + v42
	v44 = int32(0)
	if l1 <= v44 {
		v161 = v44
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v53 = v44
	goto L12
L12:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v53))))
	if v68 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v161 = l1
	goto L3
L14:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v53))))
	if v91 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v53<<(uint(int32(2))%32)))) = v83
	v88 = v83
	goto L14
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v43+v53<<(uint(int32(2))%32))))
	v83 = v74
	goto L15
L17:
	;
	goto L18
L18:
	;
	v76 = v53 << (uint(int32(2)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3+v76)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v43)))
	if v80 <= v78 {
		v88 = v78
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v83 = v80
	goto L15
L20:
	;
	if v125 < v123 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v115 = v53 << (uint(int32(2)) % 32)
	v119 = v111 + v113 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2+v115))) = v119
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115+l3)))
	v123 = v122
	v125 = v119
	goto L20
L22:
	;
	v95 = v53 << (uint(int32(2)) % 32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v43+v95)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v40)))
	v111 = v99
	v113 = v97
	goto L21
L23:
	;
	goto L24
L24:
	;
	v101 = v53 << (uint(int32(2)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2+v101)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101+v43)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+v40)))
	if v103 < v105+v107 {
		v123 = v88
		v125 = v103
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v111 = v107
	v113 = v105
	goto L21
L26:
	;
	v129 = F_palloc0(m, int32(16))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v137 = v53 + int32(1)
	if v137 != l1 {
		v53 = v137
		goto L12
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v129))) = int64(64)
	v1009 = v129
	goto L1
L30:
	;
	goto L13
L31:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(442526), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(491427), int32(2067), int32(415654))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v242 = int32(0)
	if v32 <= v242 {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	v182 = v161
	goto L37
L37:
	;
	v197 = v182 << (uint(int32(2)) % 32)
	v198 = l3 + v197
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v197+v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v200
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v197+v40)))
	v207 = v200 + v204 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v197+l2))) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	if v207 < v209 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v215 = F_palloc0(m, int32(16))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L42
	}
L39:
	;
	goto L38
L40:
	;
	v212 = v182 + int32(1)
	if v32 != v212 {
		v182 = v212
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L35
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v215)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v215))) = int64(64)
	v1009 = v215
	goto L1
L43:
	;
	v319 = v32 << (uint(int32(3)) % 32)
	v321 = v319 + int32(23)
	if v38 != 0 {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	goto L43
L45:
	;
	v248 = int32(1)
	if v32 != v248 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v258 = v242
	v259 = v242
	goto L49
L47:
	;
	v293 = v242
	goto L48
L48:
	;
	if v32&v248 == int32(0) {
		goto L44
	} else {
		goto L52
	}
L49:
	;
	v262 = int32(2)
	v263 = v258 << (uint(v262) % 32)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263+l2)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v263+l3)))
	v270 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23+v263))) = v266 - v268 + v270
	v274 = v263 | int32(4)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v274+l2)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v274+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v23+v274))) = v277 - v279 + v270
	v285 = v258 + v262
	v287 = v259 + v262
	if v287 != v32&int32(2147483646) {
		v258 = v285
		v259 = v287
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v293 = v285
	goto L48
L51:
	;
	goto L50
L52:
	;
	v300 = v293 << (uint(int32(2)) % 32)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v300+l2)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v300+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v23+v300))) = v303 - v305 + int32(1)
	goto L44
L53:
	;
	v324 = v38
	goto L55
L54:
	;
	v324 = v321 & int32(-8)
	goto L55
L55:
	;
	v325 = v27 + v324
	if v38 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v328 = v319 + v40
	goto L58
L57:
	;
	v328 = int32(0)
	goto L58
L58:
	;
	v329 = F_array_slice_size(m, v325, v328, v32, v40, v43, l3, l2, l7, l8)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	if v328 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v345 = v344 + v329
	v346 = F_palloc0(m, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L7
	} else {
		goto L65
	}
L61:
	;
	v331 = F_ArrayGetNItems(m, v32, v23)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L7
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v343 = int32(0)
	v344 = v321 & int32(120)
	goto L60
L64:
	;
	v336 = base.I32_div_s(v331+int32(7), int32(8))
	v339 = (v321 + v336) & int32(-8)
	v343 = v339
	v344 = v339
	goto L60
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v346)+8)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v346)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v345 << (uint(int32(2)) % 32)
	v355 = v346 + int32(16)
	if v42 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if int32(0) < v32 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v356 = F__emscripten_memcpy_bulkmem(m, v355, v23, v42)
	mBase = m.M
	v357 = v356
	goto L69
L68:
	;
	v357 = v355
	goto L69
L69:
	;
	goto L66
L70:
	;
	v360 = v357 + v42
	v361 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(v32) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v461 = v343
	goto L72
L72:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	v473 = int32(0)
	v484 = v32 - int32(1)
	if v484 < v473 {
		v562 = v473
		goto L83
	} else {
		goto L84
	}
L73:
	;
	v367 = int32(0)
	v372 = v361
	goto L76
L74:
	;
	v407 = v361
	goto L75
L75:
	;
	v421 = v361
	v427 = v407
	goto L79
L76:
	;
	v388 = v360 + v372<<(uint(int32(2))%32)
	v389 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v388))) = v389
	*(*int64)(unsafe.Add(mBase, uint32(v388)+8)) = v389
	*(*int64)(unsafe.Add(mBase, uint32(v388)+16)) = v389
	*(*int64)(unsafe.Add(mBase, uint32(v388)+24)) = v389
	v397 = int32(8)
	v398 = v372 + v397
	v400 = v367 + v397
	if v400 != 0 {
		v367 = v400
		v372 = v398
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v407 = v398
	goto L75
L78:
	;
	goto L77
L79:
	;
	v444 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v360+v427<<(uint(int32(2))%32)))) = v444
	v449 = v421 + v444
	if v449 != v32 {
		v421 = v449
		v427 = v427 + v444
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	v461 = v451
	goto L72
L81:
	;
	goto L80
L82:
	;
	v570 = F_array_seek(m, v325, v473, v328, v562, l7, l8)
	mBase = m.M
	v572 = v23 + int32(128)
	v576 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v572+v32<<(uint(v576)%32)-int32(4)))) = int32(1)
	v584 = v32 - v576
	if v584 < int32(0) {
		goto L94
	} else {
		goto L95
	}
L83:
	;
	goto L82
L84:
	;
	v487 = int32(1)
	if v484 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v32&v487 == int32(0) {
		v562 = v539
		goto L83
	} else {
		goto L92
	}
L86:
	;
	v539 = v473
	v540 = v484
	v541 = v487
	goto L85
L87:
	;
	goto L88
L88:
	;
	v498 = v473
	v499 = v484
	v500 = v487
	v501 = v473
	goto L89
L89:
	;
	v506 = int32(2)
	v507 = v499 << (uint(v506) % 32)
	v509 = v507 - int32(4)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l3+v509)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v43+v509)))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v507+v40)))
	v517 = v516 * v500
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v507+l3)))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v507+v43)))
	v526 = (v511-v513)*v517 + ((v520-v522)*v500 + v498)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v40+v509)))
	v529 = v528 * v517
	v531 = v499 - v506
	v533 = v501 + v506
	if v533 != v32&int32(-2) {
		v498 = v526
		v499 = v531
		v500 = v529
		v501 = v533
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v539 = v526
	v540 = v531
	v541 = v529
	goto L85
L91:
	;
	goto L90
L92:
	;
	v550 = v540 << (uint(int32(2)) % 32)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l3+v550)))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v550+v43)))
	v562 = (v552-v554)*v541 + v539
	goto L83
L93:
	;
	v644 = v23 + int32(96)
	v645 = int32(0)
	if v32 <= v645 {
		goto L104
	} else {
		goto L105
	}
L94:
	;
	goto L93
L95:
	;
	if v32&int32(1) == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v591 = int32(2)
	v597 = v32<<(uint(v591)%32) - int32(4)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v40+v597)))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v572+v597)))
	*(*int32)(unsafe.Add(mBase, uint32(v572+v584<<(uint(v591)%32)))) = v599 * v601
	v606 = v32 - int32(3)
	goto L98
L97:
	;
	v606 = v584
	goto L98
L98:
	;
	if v584 == int32(0) {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v612 = v606
	goto L100
L100:
	;
	v615 = int32(2)
	v616 = v612 << (uint(v615) % 32)
	v619 = v616 + int32(4)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v40+v619)))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v572+v619)))
	v624 = v621 * v623
	*(*int32)(unsafe.Add(mBase, uint32(v572+v616))) = v624
	v627 = v612 - int32(1)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v616+v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v572+v627<<(uint(v615)%32)))) = v632 * v624
	if v627 != 0 {
		v612 = v612 - v615
		goto L100
	} else {
		goto L102
	}
L101:
	;
	goto L94
L102:
	;
	goto L101
L103:
	;
	v722 = v23 - int32(-64)
	v724 = v23 + int32(128)
	v726 = v23 + int32(96)
	v727 = int32(0)
	v733 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v722+v32<<(uint(v733)%32)-int32(4)))) = v727
	v741 = v32 - v733
	if v727 <= v741 {
		goto L114
	} else {
		goto L115
	}
L104:
	;
	goto L103
L105:
	;
	v651 = int32(1)
	if v32 != v651 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v661 = v645
	v662 = v645
	goto L109
L107:
	;
	v696 = v645
	goto L108
L108:
	;
	if v32&v651 == int32(0) {
		goto L104
	} else {
		goto L112
	}
L109:
	;
	v665 = int32(2)
	v666 = v661 << (uint(v665) % 32)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v666+l2)))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v666+l3)))
	v673 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v644+v666))) = v669 - v671 + v673
	v677 = v666 | int32(4)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v677+l2)))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v677+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v644+v677))) = v680 - v682 + v673
	v688 = v661 + v665
	v690 = v662 + v665
	if v690 != v32&int32(2147483646) {
		v661 = v688
		v662 = v690
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v696 = v688
	goto L108
L111:
	;
	goto L110
L112:
	;
	v703 = v696 << (uint(int32(2)) % 32)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v703+l2)))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v703+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v644+v703))) = v706 - v708 + int32(1)
	goto L104
L113:
	;
	v851 = F__emscripten_memset_bulkmem(m, v23+int32(32), base.I32_extend8_s(int32(0)), v42)
	mBase = m.M
	goto L129
L114:
	;
	v748 = v741
	v750 = v727
	goto L117
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	v755 = v748 << (uint(int32(2)) % 32)
	v756 = v722 + v755
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v724+v755)))
	v759 = int32(1)
	v760 = v758 - v759
	*(*int32)(unsafe.Add(mBase, uint32(v756))) = v760
	v763 = v748 + v759
	if v32 <= v763 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L116
L119:
	;
	v831 = int32(1)
	if int32(0) < v748 {
		v748 = v748 - v831
		v750 = v750 + v831
		goto L117
	} else {
		goto L128
	}
L120:
	;
	if v750&int32(1) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v769 = int32(2)
	v770 = v763 << (uint(v769) % 32)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v726+v770)))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v724+v770)))
	v778 = v760 - (v772-int32(1))*v776
	*(*int32)(unsafe.Add(mBase, uint32(v756))) = v778
	v782 = v748 + v769
	v783 = v778
	goto L123
L122:
	;
	v782 = v763
	v783 = v760
	goto L123
L123:
	;
	if v750 == int32(0) {
		goto L119
	} else {
		goto L124
	}
L124:
	;
	v790 = v782
	v791 = v783
	goto L125
L125:
	;
	v796 = int32(2)
	v797 = v790 << (uint(v796) % 32)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v726+v797)))
	v800 = int32(1)
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v724+v797)))
	v805 = v791 - (v799-v800)*v803
	*(*int32)(unsafe.Add(mBase, uint32(v756))) = v805
	v808 = v797 + int32(4)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v726+v808)))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v724+v808)))
	v816 = v805 - (v810-v800)*v814
	*(*int32)(unsafe.Add(mBase, uint32(v756))) = v816
	v819 = v790 + v796
	if v819 != v32 {
		v790 = v819
		v791 = v816
		goto L125
	} else {
		goto L127
	}
L126:
	;
	goto L119
L127:
	;
	goto L126
L128:
	;
	goto L118
L129:
	;
	v853 = v472 << (uint(int32(3)) % 32)
	if v461 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v856 = v357 + v853
	goto L132
L131:
	;
	v856 = int32(0)
	goto L132
L132:
	;
	if v461 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v863 = v461
	goto L135
L134:
	;
	v863 = (v853 + int32(23)) & int32(-8)
	goto L135
L135:
	;
	v865 = v32 - int32(1)
	v867 = v346 + v863
	v871 = v570
	v875 = v473
	v877 = v562
	goto L136
L136:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(-64)+v865<<(uint(int32(2))%32))))
	if v890 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v1009 = v346
	goto L1
L138:
	;
	v898 = F_array_seek(m, v896, v895, v328, int32(1), l7, l8)
	mBase = m.M
	v899 = v898 - v896
	if v899 != 0 {
		goto L143
	} else {
		goto L144
	}
L139:
	;
	v895 = v877
	v896 = v871
	goto L138
L140:
	;
	goto L141
L141:
	;
	v894 = F_array_seek(m, v871, v877, v328, v890, l7, l8)
	mBase = m.M
	v895 = v890 + v877
	v896 = v894
	goto L138
L142:
	;
	if v856 != 0 {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v900 = F__emscripten_memcpy_bulkmem(m, v867, v896, v899)
	mBase = m.M
	v901 = v900
	goto L145
L144:
	;
	v901 = v867
	goto L145
L145:
	;
	goto L142
L146:
	;
	v905 = int32(1) << (uint(v875&int32(7)) % 32)
	v907 = base.I32_div_s(v875, int32(8))
	v908 = v856 + v907
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	if v328 != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	goto L148
L148:
	;
	v930 = int32(1)
	v937 = v23 + int32(32)
	v939 = v23 + int32(96)
	if v32 <= int32(0) {
		goto L157
	} else {
		goto L158
	}
L149:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v908))) = uint8(v925)
	goto L148
L150:
	;
	v915 = base.I32_div_s(v895, int32(8))
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328+v915))))
	if int32(base.Ui32(v917)>>(uint(v895&int32(7))%32))&int32(1) != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	v925 = v909 | v905
	goto L149
L153:
	;
	v923 = v909 | v905
	goto L155
L154:
	;
	v923 = v909 & (v905 ^ int32(-1))
	goto L155
L155:
	;
	v925 = v923
	goto L149
L156:
	;
	if v993 != int32(-1) {
		v865 = v993
		v867 = v901 + v899
		v871 = v899 + v896
		v875 = v875 + v930
		v877 = v895 + v930
		goto L136
	} else {
		goto L171
	}
L157:
	;
	v993 = int32(-1)
	goto L156
L158:
	;
	goto L159
L159:
	;
	v945 = int32(1)
	v946 = v32 - v945
	v948 = v946 << (uint(int32(2)) % 32)
	v949 = v937 + v948
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v949)))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v939+v948)))
	v955 = base.I32_rem_s(v950+v945, v954)
	*(*int32)(unsafe.Add(mBase, uint32(v949))) = v955
	if v946 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v993 = v983
	goto L156
L161:
	;
	v957 = v946
	v960 = v955
	goto L164
L162:
	;
	goto L163
L163:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	if v981 != 0 {
		goto L168
	} else {
		goto L169
	}
L164:
	;
	if v960 != 0 {
		v983 = v957
		goto L160
	} else {
		goto L166
	}
L165:
	;
	goto L163
L166:
	;
	v962 = int32(1)
	v963 = v957 - v962
	v965 = v963 << (uint(int32(2)) % 32)
	v966 = v937 + v965
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v966)))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v939+v965)))
	v972 = base.I32_rem_s(v967+v962, v971)
	*(*int32)(unsafe.Add(mBase, uint32(v966))) = v972
	if v963 != 0 {
		v957 = v963
		v960 = v972
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v982 = int32(0)
	goto L170
L169:
	;
	v982 = int32(-1)
	goto L170
L170:
	;
	v983 = v982
	goto L160
L171:
	;
	goto L137
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v997)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v997)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v997))) = int64(64)
	v1009 = v997
	goto L1
}
func F_array_gt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v2)
	}
}
func F_array_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = F_array_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if int32(0) < v4 {
			v10 = int32(20)
		} else {
			v10 = int32(28)
		}
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_array_lower(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v13 == int32(-1) {
			v16 = int32(28)
		} else {
			v16 = int32(4)
		}
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v7+v16)))
		if base.Ui32(v18-int32(7)) <= base.Ui32(int32(-7)) {
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = int32(0)
			if base.B2i32(v28 < v27)&base.B2i32(base.Ui32(v27) <= base.Ui32(v18)) == v28 {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int32(0)
			} else {
				if v13 == int32(-1) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
					v47 = v40
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v47 = v7 + v41<<(uint(int32(2))%32) + int32(16)
				}
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+v27<<(uint(int32(2))%32)-int32(4))))
				return v53
			}
		}
	}
}
func F_array_position(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v183
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L8
	} else {
		goto L66
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L62
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L57
	}
L5:
	;
	m.G0 = v15 + int32(16)
	goto L1
L6:
	;
	v179 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v179)
	v183 = int32(0)
	goto L5
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(2) <= v24 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v24 != int32(1) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v29 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20+v42)+16))
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v46 == int32(3) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v32 = F_array_contains_nulls(m, v20)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v41 = v39
	v42 = int32(4)
	goto L12
L16:
	;
	if v32 == int32(0) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v41 = v2
	v42 = v36 << (uint(int32(2)) % 32)
	goto L12
L18:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v49 == int32(1) {
		goto L3
	} else {
		goto L21
	}
L19:
	;
	v53 = v45
	goto L20
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v55 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v53 = v52
	goto L20
L22:
	;
	v97 = v45 - int32(1)
	v99 = F_array_create_iterator(m, v20, int32(0), v94)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L34
	}
L23:
	;
	F_get_typlenbyvalalign(m, v43, v71+int32(4), v71+int32(6), v71+int32(7))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L29
	}
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v60 = F_MemoryContextAlloc(m, v58, int32(48))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v69 == v43 {
		v94 = v55
		goto L22
	} else {
		goto L28
	}
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v60
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v43 ^ int32(-1)
	v71 = v65
	goto L23
L28:
	;
	v71 = v55
	goto L23
L29:
	;
	v81 = F_lookup_type_cache(m, v43, int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+80))
	if v83 == int32(0) {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v43
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+80))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	F_fmgr_info_cxt(m, v87, v71+int32(20), v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v94 = v71
	goto L22
L33:
	;
	F_array_free_iterator(m, v99)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L51
	}
L34:
	;
	v105 = F_array_iterate(m, v99, v15+int32(12), v15+int32(11))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v105 == int32(0) {
		v150 = v97
		v158 = v2
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v112 = v97
	goto L37
L37:
	;
	v124 = v112 + int32(1)
	if v124 < v53 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v150 = v124
	v158 = v2
	goto L33
L39:
	;
	v147 = F_array_iterate(m, v99, v15+int32(12), v15+int32(11))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L49
	}
L40:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+11)))
	if v29|v126&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v150 = v124
	v158 = int32(1)
	goto L33
L42:
	;
	if v29 == int32(0) {
		goto L39
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v137 = F_FunctionCall2Coll(m, v94+int32(20), v18, v41, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L8
	} else {
		goto L47
	}
L45:
	;
	if v126&int32(1) == int32(0) {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L41
L47:
	;
	if v137 == int32(0) {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L41
L49:
	;
	if v147 != 0 {
		v112 = v124
		goto L37
	} else {
		goto L50
	}
L50:
	;
	goto L38
L51:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v163 != v20 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_pfree(m, v20)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v158 != 0 {
		v183 = v150
		goto L5
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	goto L6
L57:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	v204 = F_format_type_be(m, v43)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v204
	F_errmsg(m, int32(188458), v15)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(491534), int32(1411), int32(244349))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(301281), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(491534), int32(1377), int32(244349))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(439263), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(491534), int32(1348), int32(244349))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_shuffle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		if v11 <= int32(0) {
			v33 = v7
			return v33
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			if v14 < int32(2) {
				v33 = v7
				return v33
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				if v19 != 0 {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					if v20 == v17 {
						v28 = v19
						v29 = v14
						v31 = F_array_shuffle_n(m, v7, v29, int32(1), v17, v28)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = v31
							return v33
						}
					} else {
						v23 = F_lookup_type_cache(m, v17, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v23
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
							v28 = v23
							v29 = v27
							v31 = F_array_shuffle_n(m, v7, v29, int32(1), v17, v28)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = v31
								return v33
							}
						}
					}
				} else {
					v23 = F_lookup_type_cache(m, v17, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v23
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
						v28 = v23
						v29 = v27
						v31 = F_array_shuffle_n(m, v7, v29, int32(1), v17, v28)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = v31
							return v33
						}
					}
				}
			}
		}
	}
}
func F_array_sort_order(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = base.B2i32(v8 != int32(0))
		v11 = F_array_sort_internal(m, v4, v10, v10, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_array_sort_order_nulls_first(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = int32(0)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v13 = F_array_sort_internal(m, v3, base.B2i32(v7 != v8), base.B2i32(v10 != v8), l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v13
		}
	}
}
func F_array_subscript_assign(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v44 int32
	_ = v44
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	if v13 <= int32(0) {
		if v8&int32(1) != 0 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v26 = F_construct_empty_array(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v29 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v29)
				v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
				v32 = v26
				v33 = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
				v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
				v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
				v42 = F_array_set_element(m, v32, v34, v12+int32(12), v37, v38, v33, v39, v40, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
					return
				}
			}
		} else {
			v32 = v10
			v33 = v13
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
			v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
			v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
			v42 = F_array_set_element(m, v32, v34, v12+int32(12), v37, v38, v33, v39, v40, v41)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
				return
			}
		}
	} else {
		if v8&int32(1) != 0 {
			return
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
			if v20 != 0 {
				return
			} else {
				if v8&int32(1) == int32(0) {
					v32 = v10
					v33 = v13
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
					v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
					v42 = F_array_set_element(m, v32, v34, v12+int32(12), v37, v38, v33, v39, v40, v41)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
						return
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v26 = F_construct_empty_array(m, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v29 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v29)
						v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
						v32 = v26
						v33 = v31
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
						v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
						v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+9)))
						v42 = F_array_set_element(m, v32, v34, v12+int32(12), v37, v38, v33, v39, v40, v41)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
							return
						}
					}
				}
			}
		}
	}
}
func F_array_subscript_fetch_slice(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+6)))
	v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18)+9)))
	v29 = F_array_get_slice(m, v15, v17, v18+int32(12), v18+int32(36), v23, v24, v25, v26, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v31))) = v29
		return
	}
}
func F_estimate_array_length(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
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
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 float32
	_ = v76
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v87 float64
	_ = v87
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v96 int32
	_ = v96
	var v98 float64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v115 float64
	_ = v115
	v5 = float64(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return v115
L2:
	;
	v115 = float64(10)
	goto L1
L3:
	;
	v13 = l1
	goto L6
L4:
	;
	if l0 == int32(0) {
		goto L2
	} else {
		goto L24
	}
L5:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)))
	if v42 != 0 {
		goto L4
	} else {
		goto L20
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	switch v17 - int32(7) {
	case 0:
		goto L8
	default:
		goto L4
	case 20:
		goto L9
	case 22:
		goto L10
	case 28:
		goto L5
	}
L7:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
	if v30 != 0 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v29 != 0 {
		v13 = v29
		goto L6
	} else {
		goto L13
	}
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 != int32(27) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 != int32(34) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	goto L2
L14:
	;
	v115 = v5
	goto L1
L15:
	;
	goto L16
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v32 = F_pg_detoast_datum(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return float64(0)
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v39 = F_ArrayGetNItems(m, v36, v32+int32(16))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v115 = base.F64_convert_i32_s(v39)
	goto L1
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v43 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v115 = v5
	goto L1
L22:
	;
	goto L23
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v115 = base.F64_convert_i32_s(v46)
	goto L1
L24:
	;
	F_examine_variable(m, l0, v13, int32(0), v8+int32(48))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	if v56 == int32(0) {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v64 = F_get_attstatsslot(m, v8+int32(12), v56, int32(5), int32(0), int32(2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	if v64 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	if v66 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v98 = v5
	goto L30
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	if v99 != 0 {
		goto L40
	} else {
		goto L41
	}
L31:
	;
	v92 = float64(0)
	goto L33
L32:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v70+v66<<(uint(int32(2))%32)-int32(4))))
	v77 = base.F64_promote_f32(v76)
	v79 = float64(1e+100)
	if base.F64_gt(v77, v79) != 0 {
		v91 = v79
		goto L35
	} else {
		goto L36
	}
L33:
	;
	F_free_attstatsslot(m, v8+int32(12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L39
	}
L34:
	;
	v92 = v91
	goto L33
L35:
	;
	goto L34
L36:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v77)&int64(9223372036854775807)) {
		v91 = v79
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v87 = float64(1)
	if base.F64_le(v77, v87) != 0 {
		v91 = v87
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v91 = base.F64_nearest(v77)
	goto L35
L39:
	;
	v98 = v92
	goto L30
L40:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	m.T0[v100].(func(*base.Module, int32))(m, v99)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L17
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if base.F64_gt(v98, float64(0)) != 0 {
		v115 = v98
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	goto L2
}
func F_getArrayIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v15 = F_executeItemOptUnwrapTarget(m, l0, l1, l2, v8+int32(8), v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)) = uint8(v19)
		v21 = int32(2)
		if v15 == v21 {
			v97 = v21
			m.G0 = v8 + int32(16)
			return v97
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			if v24 == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				if v27 == int32(0) {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
					if v40 != int32(1) {
						v97 = v21
						m.G0 = v8 + int32(16)
						return v97
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(51118210))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(344310), int32(0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496877), int32(3474), int32(29154))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					if v30 != int32(1) {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
						if v40 != int32(1) {
							v97 = v21
							m.G0 = v8 + int32(16)
							return v97
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(51118210))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(344310), int32(0))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(496877), int32(3474), int32(29154))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
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
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v35 = v34
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
						if v36 == int32(2) {
							v59 = int32(0)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
							v64 = F_DirectFunctionCall2Coll(m, int32(1442), v59, v62, v59)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v66 = F_pg_detoast_datum(m, v64)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v70 = F_numeric_int4_opt_error(m, v66, v8+int32(7))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v70
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
										if v73 != int32(1) {
											v97 = v59
											m.G0 = v8 + int32(16)
											return v97
										} else {
											v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
											if v77 != int32(1) {
												v97 = int32(2)
												m.G0 = v8 + int32(16)
												return v97
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(51118210))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(399451), int32(0))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(496877), int32(3486), int32(29154))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
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
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
							if v40 != int32(1) {
								v97 = v21
								m.G0 = v8 + int32(16)
								return v97
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(51118210))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(344310), int32(0))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496877), int32(3474), int32(29154))
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
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
			} else {
				v35 = v24
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
				if v36 == int32(2) {
					v59 = int32(0)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					v64 = F_DirectFunctionCall2Coll(m, int32(1442), v59, v62, v59)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = F_pg_detoast_datum(m, v64)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v70 = F_numeric_int4_opt_error(m, v66, v8+int32(7))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v70
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
								if v73 != int32(1) {
									v97 = v59
									m.G0 = v8 + int32(16)
									return v97
								} else {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
									if v77 != int32(1) {
										v97 = int32(2)
										m.G0 = v8 + int32(16)
										return v97
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(51118210))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(399451), int32(0))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496877), int32(3486), int32(29154))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
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
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
					if v40 != int32(1) {
						v97 = v21
						m.G0 = v8 + int32(16)
						return v97
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(51118210))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(344310), int32(0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496877), int32(3474), int32(29154))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
func F_get_array_element_end(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 < v10 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = int32(1)
		v15 = v10 - v14
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v15))))
		if v17 != v14 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v20 == int32(0) {
				return int32(0)
			} else {
				v24 = v15 << (uint(int32(2)) % 32)
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v24+v25)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v24+v20)))
				if v27 != v29 {
					return int32(0)
				} else {
					if v10 < v11 {
						v33 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v10+v13))) = uint8(v33)
						return v33
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v37 == int32(0) {
							return int32(0)
						} else {
							if l1 != 0 {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
								if v41 != 0 {
									v48 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
									return int32(0)
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
									v44 = F_cstring_to_text_with_len(m, v37, v42-v37)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v48 = v44
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
										return int32(0)
									}
								}
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
								v44 = F_cstring_to_text_with_len(m, v37, v42-v37)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v48 = v44
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
									return int32(0)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_get_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v13 < v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return int32(0)
L2:
	;
	v17 = v13 << (uint(int32(2)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v17+v18))) = int32(-1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22+v17)))
	if base.Ui32(v24) < base.Ui32(int32(-2147483647)) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if v14|v13 != 0 {
		goto L1
	} else {
		goto L51
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = m.G0
	v32 = v30 - int32(80)
	m.G0 = v32
	if v27 == int32(4488504) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	m.G0 = v32 + int32(80)
	if v115 != 0 {
		goto L46
	} else {
		goto L47
	}
L7:
	;
	v115 = int32(16)
	goto L6
L8:
	;
	goto L9
L9:
	;
	goto L11
L10:
	;
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+68)) = uint8(v42)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v44 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	if v48 != int32(5) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v40 = F__emscripten_memcpy_bulkmem(m, v32+int32(12), v27, int32(68))
	mBase = m.M
	goto L13
L13:
	;
	goto L10
L14:
	;
	v51 = int32(11)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v61 = F_json_lex(m, v32+int32(12))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v55 = int32(6)
	goto L19
L18:
	;
	v55 = v51
	goto L19
L19:
	;
	if v48 == int32(12) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v58 = v51
	goto L22
L21:
	;
	v58 = v55
	goto L22
L22:
	;
	v115 = v58
	goto L6
L23:
	;
	return int32(0)
L24:
	;
	if v61 != 0 {
		v115 = v61
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	if v65 == int32(6) {
		v106 = v2
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v110 = F_json_lex(m, v32+int32(12))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L23
	} else {
		goto L44
	}
L27:
	;
	v73 = v2
	goto L28
L28:
	;
	v78 = F_parse_array_element(m, v32+int32(12), int32(1839104))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v115 = v97
	goto L6
L30:
	;
	if v78 != 0 {
		v115 = v78
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v81 = v73 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	if v82 != int32(7) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v82 == int32(6) {
		v106 = v81
		goto L26
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v97 = F_json_lex(m, v32+int32(12))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L23
	} else {
		goto L42
	}
L35:
	;
	v87 = int32(11)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v90 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v91 = int32(7)
	goto L38
L37:
	;
	v91 = v87
	goto L38
L38:
	;
	if v82 == int32(12) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v94 = v87
	goto L41
L40:
	;
	v94 = v91
	goto L41
L41:
	;
	v115 = v94
	goto L6
L42:
	;
	if v97 == int32(0) {
		v73 = v81
		goto L28
	} else {
		goto L43
	}
L43:
	;
	goto L29
L44:
	;
	if v110 != 0 {
		v115 = v110
		goto L6
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12)))) = v106
	v115 = int32(0)
	goto L6
L46:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_json_errsave_error(m, v115, v124, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L23
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v133 = v130 + v13<<(uint(int32(2))%32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v128 < int32(0)-v134 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v128 + v134
	goto L1
L51:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v140
	goto L1
}
func F_initArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v4 = l3
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 == int32(0) {
		v13 = F_get_element_type(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v13 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = F_format_type_be(m, l0)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
							F_errmsg(m, int32(364671), v9)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(491427), int32(5525), int32(205894))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
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
				v19 = v13
				if v4 != 0 {
					v24 = F_AllocSetContextCreateInternal(m, l2, int32(205913), int32(0), int32(8192), int32(8388608))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = v24
						v28 = F_MemoryContextAllocZero(m, v26, int32(92))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v4)
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v19
							*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = l0
							m.G0 = v9 + int32(16)
							return v28
						}
					}
				} else {
					v26 = l2
					v28 = F_MemoryContextAllocZero(m, v26, int32(92))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v4)
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = l0
						m.G0 = v9 + int32(16)
						return v28
					}
				}
			}
		}
	} else {
		v19 = l1
		if v4 != 0 {
			v24 = F_AllocSetContextCreateInternal(m, l2, int32(205913), int32(0), int32(8192), int32(8388608))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = v24
				v28 = F_MemoryContextAllocZero(m, v26, int32(92))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v4)
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v19
					*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = l0
					m.G0 = v9 + int32(16)
					return v28
				}
			}
		} else {
			v26 = l2
			v28 = F_MemoryContextAllocZero(m, v26, int32(92))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v4)
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = l0
				m.G0 = v9 + int32(16)
				return v28
			}
		}
	}
}
func F_makeArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
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
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
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
	var v166 int32
	_ = v166
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	v4 = int32(0)
	v14 = int32(4487040)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = l1
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v18 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v15
	if l2 != 0 {
		goto L50
	} else {
		goto L51
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v23 = F_palloc0(m, int32(16))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v33 = l0 + int32(32)
	v34 = F_ArrayGetNItems(m, v18, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = int64(64)
	v181 = v23
	goto L1
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v38 = l0 + int32(56)
	F_ArrayCheckBounds(m, v36, v33, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v43 = v41 << (uint(int32(3)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = v61 + v44
	v63 = F_palloc0(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L13
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v50 = base.I32_div_s(v46+int32(7), int32(8))
	v55 = (v43 + v50 + int32(23)) & int32(-8)
	v60 = v55
	v61 = v55
	goto L9
L11:
	;
	goto L12
L12:
	;
	v60 = v4
	v61 = (v43 + int32(23)) & int32(-8)
	goto L9
L13:
	;
	v65 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v62 << (uint(v65) % 32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v71
	v74 = v63 + int32(16)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v77 = v75 << (uint(v65) % 32)
	if v77 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v80 = int32(2)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v85 = v83 << (uint(v80) % 32)
	if v85 != 0 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v78 = F__emscripten_memcpy_bulkmem(m, v74, v33, v77)
	mBase = m.M
	v79 = v78
	goto L17
L16:
	;
	v79 = v74
	goto L17
L17:
	;
	goto L14
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v88 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v86 = F__emscripten_memcpy_bulkmem(m, v79+v68<<(uint(v80)%32), v38, v85)
	mBase = m.M
	goto L21
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	v96 = v88
	goto L24
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L24
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v99 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v102 == int32(0) {
		v181 = v63
		goto L1
	} else {
		goto L29
	}
L26:
	;
	v100 = F__emscripten_memcpy_bulkmem(m, v96+v63, v98, v99)
	mBase = m.M
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v105 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v111 = v79 + v106<<(uint(int32(3))%32)
	goto L32
L31:
	;
	v111 = int32(0)
	goto L32
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v112 <= int32(0) {
		v181 = v63
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v117 = int32(1)
	v120 = v117
	v122 = v116
	v124 = v117
	v125 = v112
	v127 = v111
	v128 = v102
	v129 = v115
	goto L34
L34:
	;
	if v124&v129 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v151 == int32(1) {
		v181 = v63
		goto L1
	} else {
		goto L49
	}
L36:
	;
	v137 = v120 | v122
	goto L38
L37:
	;
	v137 = v122 & (v120 ^ int32(-1))
	goto L38
L38:
	;
	v138 = int32(1)
	v139 = v125 - v138
	v141 = v120 << (uint(v138) % 32)
	if v141 == int32(256) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v137)
	if v139 == int32(0) {
		v181 = v63
		goto L1
	} else {
		goto L42
	}
L40:
	;
	v151 = v141
	v152 = v137
	v153 = v127
	goto L41
L41:
	;
	v155 = v124 << (uint(int32(1)) % 32)
	if v155 == int32(256) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v147 = int32(1)
	v148 = v127 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v151 = v147
	v152 = v149
	v153 = v148
	goto L41
L43:
	;
	goto L35
L44:
	;
	if v139 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	v164 = v155
	v165 = v128
	v166 = v129
	goto L46
L46:
	;
	if base.Ui32(int32(1)) < base.Ui32(v125) {
		v120 = v151
		v122 = v152
		v124 = v164
		v125 = v139
		v127 = v153
		v128 = v165
		v129 = v166
		goto L34
	} else {
		goto L48
	}
L47:
	;
	v160 = int32(1)
	v161 = v128 + v160
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v164 = v160
	v165 = v161
	v166 = v162
	goto L46
L48:
	;
	goto L43
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v152)
	v181 = v63
	goto L1
L50:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_MemoryContextDelete(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	return v181
L53:
	;
	goto L52
}
