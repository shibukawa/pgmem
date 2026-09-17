package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReplaceVarFromTargetList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v296 int32
	_ = v296
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v17 == v7 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(32)
	return v296
L2:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v277 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_expandRTE(m, l1, v20, int32(0), v22, v23, base.B2i32(v24 != int32(2249)), v15+int32(20), v15+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if l2 != 0 {
		goto L27
	} else {
		goto L28
	}
L6:
	;
	return int32(0)
L7:
	;
	v36 = F_palloc0(m, int32(24))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = int64(36)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v40
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v46 == int32(2249) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v49 = v44
	goto L11
L10:
	;
	v49 = int32(0)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v53 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v56 <= int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v69 = v7
	v70 = v7
	goto L14
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v69<<(uint(int32(2))%32))))
	if v75 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L2
L16:
	;
	v87 = F_lappend(m, v86, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L22
	}
L17:
	;
	v85 = int32(0)
	v86 = v70
	goto L16
L18:
	;
	goto L19
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v79 != int32(6) {
		v85 = v75
		v86 = v70
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v82 = F_ReplaceVarFromTargetList(m, v75, l1, l2, l3, l4, l5)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v85 = v82
	v86 = v84
	goto L16
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v87
	v91 = v69 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v91 < v92 {
		v69 = v91
		v70 = v87
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v172 = F_copyObjectImpl(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L53
	}
L25:
	;
	if v131 != 0 {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	goto L25
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v97 <= int32(0) {
		v131 = int32(0)
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v131 = int32(0)
	goto L26
L30:
	;
	v100 = int32(0)
	if v100 < v97 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v103 = v97
	goto L33
L32:
	;
	v103 = v100
	goto L33
L33:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v108 = int32(0)
	goto L34
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v104+v108<<(uint(int32(2))%32))))
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+8)))
	if v117 == v17&int32(_a_F_ReplaceVarFromTargetList_0) {
		v131 = v116
		goto L26
	} else {
		goto L36
	}
L35:
	;
	goto L29
L36:
	;
	v120 = v108 + int32(1)
	if v120 != v103 {
		v108 = v120
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+26)))
	if v133 != int32(1) {
		goto L24
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	switch l4 - int32(1) {
	case 0:
		goto L44
	case 1:
		goto L43
	default:
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L48
	}
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_get_typlenbyval(m, v143, v15+int32(20), v15+int32(16))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L6
	} else {
		goto L46
	}
L44:
	;
	v138 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = l5
	v296 = v138
	goto L1
L46:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v153 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+20)))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+16)))
	v155 = F_coerce_null_to_domain(m, v150, v151, v152, v153, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v296 = v155
	goto L1
L48:
	;
	v161 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v161
	F_errmsg_internal(m, int32(_a_F_ReplaceVarFromTargetList_1), v15)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_ReplaceVarFromTargetList_2), int32(1902), int32(_a_F_ReplaceVarFromTargetList_3))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L6
	} else {
		goto L86
	}
L52:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v188 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	if v172 == int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v176 != int32(8) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v181 = F_expression_tree_walker_impl(m, v172, int32(1056), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L6
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v185 == int32(3) {
		goto L51
	} else {
		goto L60
	}
L58:
	;
	if v181 == int32(0) {
		goto L52
	} else {
		goto L59
	}
L59:
	;
	goto L51
L60:
	;
	goto L52
L61:
	;
	v296 = v172
	goto L1
L62:
	;
	goto L63
L63:
	;
	if l3 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v188
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l3
	if v172 == v192 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L6
	} else {
		goto L83
	}
L67:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v220 != int32(6) {
		goto L78
	} else {
		goto L79
	}
L68:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v197 != int32(67) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v217 = F_expression_tree_walker_impl(m, v172, int32(1057), v15+int32(20))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L6
	} else {
		goto L77
	}
L70:
	;
	if v197 != int32(6) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(1)
	v212 = F_query_tree_walker_impl(m, v172, int32(1057), v15+int32(20), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L76
	}
L73:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v202 != l3 {
		goto L67
	} else {
		goto L74
	}
L74:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v172)+28))
	if v204 != 0 {
		goto L67
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+32)) = v188
	goto L67
L76:
	;
	goto L67
L77:
	;
	goto L67
L78:
	;
	v227 = F_palloc0(m, int32(16))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L82
	}
L79:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v223 != l3 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v172)+28))
	if v225 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v296 = v172
	goto L1
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v227))) = int64(61)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+12)) = v172
	*(*uint8)(unsafe.Add(mBase, uint32(v227)+8)) = uint8(base.B2i32(v231 == int32(1)))
	v296 = v227
	goto L1
L83:
	;
	F_errmsg_internal(m, int32(_a_F_ReplaceVarFromTargetList_4), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_ReplaceVarFromTargetList_2), int32(1933), int32(_a_F_ReplaceVarFromTargetList_3))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(_a_F_ReplaceVarFromTargetList_5), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_ReplaceVarFromTargetList_2), int32(1923), int32(_a_F_ReplaceVarFromTargetList_3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	v296 = v36
	goto L1
L91:
	;
	goto L92
L92:
	;
	v281 = F_palloc0(m, int32(16))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v281))) = int64(61)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = v36
	*(*uint8)(unsafe.Add(mBase, uint32(v281)+8)) = uint8(base.B2i32(v285 == int32(1)))
	v296 = v281
	goto L1
}
func F_cmp_var(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v383 int32
	_ = v383
	var v393 int32
	_ = v393
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v10 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	if v9 == int32(_a_F_cmp_var_0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = int32(1)
	goto L9
L8:
	;
	v22 = int32(-1)
	goto L9
L9:
	;
	return v22
L10:
	;
	if v24 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v24 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v29 = int32(-1)
	goto L15
L14:
	;
	v29 = int32(1)
	goto L15
L15:
	;
	return v29
L16:
	;
	if v9 == int32(_a_F_cmp_var_0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v9 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L19:
	;
	return int32(1)
L20:
	;
	goto L21
L21:
	;
	v41 = int32(0)
	if base.B2i32(v31 < v33)&base.B2i32(v41 < v11) == v41 {
		v72 = v33
		v76 = v41
		goto L24
	} else {
		goto L25
	}
L22:
	;
	return v214
L23:
	;
	v214 = v204
	goto L22
L24:
	;
	if base.B2i32(v10 <= int32(0))|base.B2i32(v31 <= v72) != 0 {
		v108 = v31
		v110 = v41
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v53 = v33
	v57 = v41
	goto L26
L26:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v57<<(uint(int32(1))%32)))))
	if v63 != 0 {
		v204 = int32(1)
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v72 = v67
	v76 = v65
	goto L24
L28:
	;
	v64 = int32(1)
	v65 = v57 + v64
	v67 = v53 - v64
	if v67 <= v31 {
		v72 = v67
		v76 = v65
		goto L24
	} else {
		goto L29
	}
L29:
	;
	if v65 < v11 {
		v53 = v67
		v57 = v65
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	if v72 != v108 {
		v150 = v76
		v151 = v110
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v89 = v31
	v91 = v41
	goto L33
L33:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v91<<(uint(int32(1))%32)))))
	if v96 != 0 {
		v204 = int32(-1)
		goto L23
	} else {
		goto L35
	}
L34:
	;
	v108 = v100
	v110 = v98
	goto L31
L35:
	;
	v97 = int32(1)
	v98 = v91 + v97
	v100 = v89 - v97
	if v100 <= v72 {
		v108 = v100
		v110 = v98
		goto L31
	} else {
		goto L36
	}
L36:
	;
	if v98 < v10 {
		v89 = v100
		v91 = v98
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	if v11 < v150 {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	v119 = v76
	v120 = v110
	goto L40
L40:
	;
	if base.B2i32(v11 <= v119)|base.B2i32(v10 <= v120) != 0 {
		v150 = v119
		v151 = v120
		goto L38
	} else {
		goto L42
	}
L41:
	;
	if base.I32_extend16_s(v136) < base.I32_extend16_s(v134) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v125 = int32(1)
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v119<<(uint(v125)%32)))))
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120<<(uint(v125)%32)+v32))))
	if v134 == v136 {
		v119 = v119 + v125
		v120 = v120 + v125
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v143 = int32(1)
	goto L46
L45:
	;
	v143 = int32(-1)
	goto L46
L46:
	;
	v214 = v143
	goto L22
L47:
	;
	v154 = v150
	goto L49
L48:
	;
	v154 = v11
	goto L49
L49:
	;
	v161 = v150
	goto L50
L50:
	;
	if v154 == v161 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v204 = v187
	goto L23
L52:
	;
	if v10 < v151 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v187 = int32(1)
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v161<<(uint(v187)%32)))))
	if v193 == int32(0) {
		v161 = v161 + v187
		goto L50
	} else {
		goto L64
	}
L55:
	;
	v166 = v151
	goto L57
L56:
	;
	v166 = v10
	goto L57
L57:
	;
	v174 = v151
	goto L58
L58:
	;
	if v166 == v174 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v204 = int32(-1)
	goto L23
L60:
	;
	v214 = int32(0)
	goto L22
L61:
	;
	goto L62
L62:
	;
	v178 = int32(1)
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v174<<(uint(v178)%32)+v32))))
	if v183 == int32(0) {
		v174 = v174 + v178
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	goto L51
L65:
	;
	return int32(-1)
L66:
	;
	goto L67
L67:
	;
	v220 = int32(0)
	if base.B2i32(v33 < v31)&base.B2i32(v220 < v10) == v220 {
		v251 = v31
		v255 = v220
		goto L70
	} else {
		goto L71
	}
L68:
	;
	return v393
L69:
	;
	v393 = v383
	goto L68
L70:
	;
	if base.B2i32(v11 <= int32(0))|base.B2i32(v33 <= v251) != 0 {
		v287 = v33
		v289 = v220
		goto L77
	} else {
		goto L78
	}
L71:
	;
	v232 = v31
	v236 = v220
	goto L72
L72:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v236<<(uint(int32(1))%32)))))
	if v242 != 0 {
		v383 = int32(1)
		goto L69
	} else {
		goto L74
	}
L73:
	;
	v251 = v246
	v255 = v244
	goto L70
L74:
	;
	v243 = int32(1)
	v244 = v236 + v243
	v246 = v232 - v243
	if v246 <= v33 {
		v251 = v246
		v255 = v244
		goto L70
	} else {
		goto L75
	}
L75:
	;
	if v244 < v10 {
		v232 = v246
		v236 = v244
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	if v251 != v287 {
		v329 = v255
		v330 = v289
		goto L84
	} else {
		goto L85
	}
L78:
	;
	v268 = v33
	v270 = v220
	goto L79
L79:
	;
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v270<<(uint(int32(1))%32)))))
	if v275 != 0 {
		v383 = int32(-1)
		goto L69
	} else {
		goto L81
	}
L80:
	;
	v287 = v279
	v289 = v277
	goto L77
L81:
	;
	v276 = int32(1)
	v277 = v270 + v276
	v279 = v268 - v276
	if v279 <= v251 {
		v287 = v279
		v289 = v277
		goto L77
	} else {
		goto L82
	}
L82:
	;
	if v277 < v11 {
		v268 = v279
		v270 = v277
		goto L79
	} else {
		goto L83
	}
L83:
	;
	goto L80
L84:
	;
	if v10 < v329 {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	v298 = v255
	v299 = v289
	goto L86
L86:
	;
	if base.B2i32(v10 <= v298)|base.B2i32(v11 <= v299) != 0 {
		v329 = v298
		v330 = v299
		goto L84
	} else {
		goto L88
	}
L87:
	;
	if base.I32_extend16_s(v315) < base.I32_extend16_s(v313) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v304 = int32(1)
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v298<<(uint(v304)%32)))))
	v315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299<<(uint(v304)%32)+v34))))
	if v313 == v315 {
		v298 = v298 + v304
		v299 = v299 + v304
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v322 = int32(1)
	goto L92
L91:
	;
	v322 = int32(-1)
	goto L92
L92:
	;
	v393 = v322
	goto L68
L93:
	;
	v333 = v329
	goto L95
L94:
	;
	v333 = v10
	goto L95
L95:
	;
	v340 = v329
	goto L96
L96:
	;
	if v333 == v340 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v383 = v366
	goto L69
L98:
	;
	if v11 < v330 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v366 = int32(1)
	v372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v340<<(uint(v366)%32)))))
	if v372 == int32(0) {
		v340 = v340 + v366
		goto L96
	} else {
		goto L110
	}
L101:
	;
	v345 = v330
	goto L103
L102:
	;
	v345 = v11
	goto L103
L103:
	;
	v353 = v330
	goto L104
L104:
	;
	if v345 == v353 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v383 = int32(-1)
	goto L69
L106:
	;
	v393 = int32(0)
	goto L68
L107:
	;
	goto L108
L108:
	;
	v357 = int32(1)
	v362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v353<<(uint(v357)%32)+v34))))
	if v362 == int32(0) {
		v353 = v353 + v357
		goto L104
	} else {
		goto L109
	}
L109:
	;
	goto L105
L110:
	;
	goto L97
}
func F_lookup_var_attr_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
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
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v378 int32
	_ = v378
	var v387 int32
	_ = v387
	v5 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l0 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v51 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v23 = int32(1)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v24 <= v23 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = v23
	goto L7
L6:
	;
	v27 = v24
	goto L7
L7:
	;
	v31 = int32(0)
	v33 = v5
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v31<<(uint(int32(2))%32))))
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v51 = v42
	goto L1
L10:
	;
	v42 = v33 + base.I32_popcnt(v39)
	goto L12
L11:
	;
	v42 = v33
	goto L12
L12:
	;
	v44 = v31 + int32(1)
	if v44 != v27 {
		v31 = v44
		v33 = v42
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v54 = v52
	goto L16
L15:
	;
	v54 = int32(0)
	goto L16
L16:
	;
	v58 = F_palloc(m, (v54+v51)<<(uint(int32(2))%32))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	if l0 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	m.G0 = v14 + int32(16)
	return v387
L20:
	;
	F_pfree(m, v58)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L17
	} else {
		goto L85
	}
L21:
	;
	if int32(0) <= v118 {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v118 = base.I32_ctz(v104) | v105<<(uint(int32(5))%32)
	goto L21
L23:
	;
	v118 = int32(-2)
	goto L21
L24:
	;
	v71 = base.I32_div_s(int32(0), int32(32))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v72 <= v71 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v75 = l0 + int32(8)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v71<<(uint(int32(2))%32))))
	v82 = v79 & int32(-1)
	if v82 != 0 {
		v104 = v82
		v105 = v71
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v84 = v71 + int32(1)
	if v84 == v72 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v87 = v84
	goto L28
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75+v87<<(uint(int32(2))%32))))
	if v94 != 0 {
		v104 = v94
		v105 = v87
		goto L22
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v96 = v87 + int32(1)
	if v96 != v72 {
		v87 = v96
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v127 = v118
	v129 = v5
	goto L35
L33:
	;
	v229 = v5
	goto L34
L34:
	;
	if l1 == int32(0) {
		v387 = v58
		goto L19
	} else {
		goto L56
	}
L35:
	;
	v134 = v58 + v129<<(uint(int32(2))%32)
	v135 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v135
	if l2 <= v135 {
		goto L20
	} else {
		goto L37
	}
L36:
	;
	v229 = v162
	goto L34
L37:
	;
	v144 = v135
	goto L38
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l3+v144<<(uint(int32(2))%32))))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+224))
	if v155 != v127 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v154
	v162 = v129 + int32(1)
	if l0 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v158 = v144 + int32(1)
	if l2 != v158 {
		v144 = v158
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	goto L20
L44:
	;
	if int32(0) <= v218 {
		v127 = v218
		v129 = v162
		goto L35
	} else {
		goto L55
	}
L45:
	;
	v218 = base.I32_ctz(v204) | v205<<(uint(int32(5))%32)
	goto L44
L46:
	;
	v218 = int32(-2)
	goto L44
L47:
	;
	v169 = v127 + int32(1)
	v171 = base.I32_div_s(v169, int32(32))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v172 <= v171 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v175 = l0 + int32(8)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v171<<(uint(int32(2))%32))))
	v182 = v179 & (int32(-1) << (uint(v169) % 32))
	if v182 != 0 {
		v204 = v182
		v205 = v171
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v184 = v171 + int32(1)
	if v184 == v172 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v187 = v184
	goto L51
L51:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175+v187<<(uint(int32(2))%32))))
	if v194 != 0 {
		v204 = v194
		v205 = v187
		goto L45
	} else {
		goto L53
	}
L52:
	;
	goto L46
L53:
	;
	v196 = v187 + int32(1)
	if v196 != v172 {
		v187 = v196
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L36
L56:
	;
	v234 = int32(0)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v235 <= v234 {
		v387 = v58
		goto L19
	} else {
		goto L57
	}
L57:
	;
	v244 = v234
	v246 = v229
	goto L58
L58:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249+v244<<(uint(int32(2))%32))))
	v255 = F_palloc0(m, int32(248))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L17
	} else {
		goto L61
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L17
	} else {
		goto L82
	}
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = int32(-1)
	v259 = F_exprType(m, v253)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L17
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = v259
	v262 = F_exprTypmod(m, v253)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+8)) = v262
	v265 = F_exprCollation(m, v253)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L17
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+16)) = v265
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v271 = F_SearchSysCacheCopy(m, int32(82), v269, int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L17
	} else {
		goto L65
	}
L65:
	;
	if v271 == int32(0) {
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)+16))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+22)))
	v277 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+224)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v255)+20)) = v277
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+184)) = v281
	v283 = v275 + v276
	*(*int32)(unsafe.Add(mBase, uint32(v255)+12)) = v283
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v255)+204)) = uint16(v285)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+214)) = uint8(v287)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+188)) = v281
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+219)) = uint8(v289)
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v255)+206)) = uint16(v292)
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+215)) = uint8(v294)
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+192)) = v281
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+220)) = uint8(v296)
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v255)+208)) = uint16(v299)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+216)) = uint8(v301)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+196)) = v281
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+221)) = uint8(v303)
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v255)+210)) = uint16(v306)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+217)) = uint8(v308)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+200)) = v281
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+222)) = uint8(v310)
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v255)+212)) = uint16(v313)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+218)) = uint8(v315)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+223)) = uint8(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v283)+124))
	if v319 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58+v246<<(uint(int32(2))%32)))) = v338
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+232))
	*(*int32)(unsafe.Add(mBase, uint32(v338)+232)) = v344
	v346 = int32(1)
	v349 = v244 + v346
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v349 < v350 {
		v244 = v349
		v246 = v246 + v346
		goto L58
	} else {
		goto L81
	}
L68:
	;
	F_pfree(m, v271)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L17
	} else {
		goto L79
	}
L69:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v255)+24))
	if v327 == int32(0) {
		goto L68
	} else {
		goto L77
	}
L70:
	;
	v321 = F_OidFunctionCall1Coll(m, v319, int32(0), v255)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L17
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v323 = F_std_typanalyze(m, v255)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L17
	} else {
		goto L75
	}
L73:
	;
	if v321 != 0 {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L68
L75:
	;
	if v323 == int32(0) {
		goto L68
	} else {
		goto L76
	}
L76:
	;
	goto L69
L77:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v255)+28))
	if int32(0) < v330 {
		v338 = v255
		goto L67
	} else {
		goto L78
	}
L78:
	;
	goto L68
L79:
	;
	F_pfree(m, v255)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	v338 = int32(0)
	goto L67
L81:
	;
	v387 = v58
	goto L19
L82:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v356
	F_errmsg_internal(m, int32(_a_F_lookup_var_attr_stats_0), v14)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L17
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_lookup_var_attr_stats_1), int32(554), int32(_a_F_lookup_var_attr_stats_2))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L17
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
	v387 = int32(0)
	goto L19
}
func F_pull_var_clause_walker(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v4 - int32(6) {
		case 0:
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v7 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_pull_var_clause_walker_0), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pull_var_clause_walker_1), int32(680), int32(_a_F_pull_var_clause_walker_2))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v149 = F_lappend(m, v148, l0)
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v149
					return int32(0)
				}
			}
		case 1, 2:
			v128 = F_expression_tree_walker_impl(m, l0, int32(903), l1)
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				v132 = v128
				return v132
			}
		case 3:
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			if v8 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_pull_var_clause_walker_3), int32(0))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pull_var_clause_walker_1), int32(687), int32(_a_F_pull_var_clause_walker_2))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v9&int32(1) != 0 {
					v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v149 = F_lappend(m, v148, l0)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v149
						return int32(0)
					}
				} else {
					if v9&int32(2) != 0 {
						v128 = F_expression_tree_walker_impl(m, l0, int32(903), l1)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							v132 = v128
							return v132
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_pull_var_clause_walker_4), int32(0))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pull_var_clause_walker_1), int32(699), int32(_a_F_pull_var_clause_walker_2))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
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
		case 4:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v29 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_pull_var_clause_walker_5), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pull_var_clause_walker_1), int32(704), int32(_a_F_pull_var_clause_walker_2))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v30&int32(1) != 0 {
					v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v149 = F_lappend(m, v148, l0)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v149
						return int32(0)
					}
				} else {
					if v30&int32(2) != 0 {
						v128 = F_expression_tree_walker_impl(m, l0, int32(903), l1)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							v132 = v128
							return v132
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_pull_var_clause_walker_6), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pull_var_clause_walker_1), int32(716), int32(_a_F_pull_var_clause_walker_2))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
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
		case 5:
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v48&int32(4) != 0 {
				v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v149 = F_lappend(m, v148, l0)
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v149
					return int32(0)
				}
			} else {
				if v48&int32(8) != 0 {
					v128 = F_expression_tree_walker_impl(m, l0, int32(903), l1)
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int32(0)
					} else {
						v132 = v128
						return v132
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pull_var_clause_walker_7), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pull_var_clause_walker_1), int32(732), int32(_a_F_pull_var_clause_walker_2))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
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
		default:
			if v4 == int32(319) {
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v107 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pull_var_clause_walker_8), int32(0))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pull_var_clause_walker_1), int32(737), int32(_a_F_pull_var_clause_walker_2))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v108&int32(16) != 0 {
						v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v149 = F_lappend(m, v148, l0)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v149
							return int32(0)
						}
					} else {
						if v108&int32(32) != 0 {
							v128 = F_expression_tree_walker_impl(m, l0, int32(903), l1)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								v132 = v128
								return v132
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_pull_var_clause_walker_9), int32(0))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pull_var_clause_walker_1), int32(749), int32(_a_F_pull_var_clause_walker_2))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
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
				v128 = F_expression_tree_walker_impl(m, l0, int32(903), l1)
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int32(0)
				} else {
					v132 = v128
					return v132
				}
			}
		}
	} else {
		v132 = int32(0)
		return v132
	}
}
func F_var_eq_const(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 float64
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 float32
	_ = v22
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 float64
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 float32
	_ = v83
	var v85 float32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v117 int32
	_ = v117
	var v118 float64
	_ = v118
	var v121 float64
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 float64
	_ = v128
	var v131 int32
	_ = v131
	var v138 float64
	_ = v138
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v147 float64
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
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
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v219 float32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v238 float64
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 float64
	_ = v262
	var v267 int32
	_ = v267
	var v268 float32
	_ = v268
	var v271 float32
	_ = v271
	var v274 float32
	_ = v274
	var v277 float32
	_ = v277
	var v279 float64
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v297 float64
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v310 float64
	_ = v310
	var v316 float32
	_ = v316
	var v318 float64
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v334 float64
	_ = v334
	var v339 float64
	_ = v339
	var v347 float64
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 float64
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 float32
	_ = v359
	var v361 float32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v388 float64
	_ = v388
	var v389 float64
	_ = v389
	var v393 int32
	_ = v393
	var v394 float64
	_ = v394
	var v397 float64
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 float64
	_ = v404
	var v407 int32
	_ = v407
	var v414 float64
	_ = v414
	var v417 float64
	_ = v417
	var v418 int32
	_ = v418
	var v423 float64
	_ = v423
	var v424 int32
	_ = v424
	var v426 float64
	_ = v426
	var v430 float64
	_ = v430
	var v433 int32
	_ = v433
	var v439 float32
	_ = v439
	var v440 float64
	_ = v440
	var v455 float64
	_ = v455
	var v460 int32
	_ = v460
	var v472 float64
	_ = v472
	var v478 float64
	_ = v478
	var v496 float64
	_ = v496
	v11 = float64(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	if l4 != 0 {
		v496 = v11
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v16 + int32(112)
	return v496
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
	v22 = *(*float32)(unsafe.Add(mBase, uint32(v19+v20)+8))
	v26 = base.F64_promote_f32(v22)
	goto L5
L4:
	;
	v26 = float64(0)
	goto L5
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v27 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l6 != 0 {
		goto L133
	} else {
		goto L134
	}
L7:
	;
	if v18 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v30)+120))
	if base.F64_ge(v33, float64(1)) == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v472 = base.F64_div(float64(1), v33)
	goto L6
L11:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v155 = F_get_attstatsslot(m, v16+int32(76), v151, int32(1), int32(0), int32(3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L14
	} else {
		goto L60
	}
L12:
	;
	v73 = v16 + int32(40)
	v74 = int32(0)
	v75 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v74)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v79 != 0 {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	v44 = F_get_opcode(m, l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return float64(0)
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v48 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	if v44 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v51 = F_get_func_leakproof(m, v44)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v51 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v55 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	if v55 == int32(0) {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v59 = F_get_func_name(m, v44)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v59
	F_errmsg_internal(m, int32(_a_F_var_eq_const_0), v16)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_var_eq_const_1), int32(_a_F_var_eq_const_2), int32(_a_F_var_eq_const_3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	v472 = base.F64_div(float64(1), v147)
	goto L6
L26:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v117 != 0 {
		goto L40
	} else {
		goto L41
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
	v82 = v80 + v81
	v83 = *(*float32)(unsafe.Add(mBase, uint32(v82)+8))
	v85 = *(*float32)(unsafe.Add(mBase, uint32(v82)+16))
	v112 = base.F64_promote_f32(v85)
	v113 = base.F64_promote_f32(v83)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v87 == int32(16) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v112 = float64(2)
	v113 = v75
	goto L26
L31:
	;
	goto L32
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v91 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v98 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	if v94 != int32(5) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v112 = float64(-1)
	v113 = v75
	goto L26
L36:
	;
	v112 = float64(0)
	v113 = v75
	goto L26
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v101 != int32(6) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+8)))
	switch v105 - int32(_a_F_var_eq_const_4) {
	case 0:
		goto L39
	default:
		goto L36
	case 5:
		v112 = float64(-1)
		v113 = v75
		goto L26
	}
L39:
	;
	v112 = float64(1)
	v113 = v75
	goto L26
L40:
	;
	v118 = base.F64_neg(base.F64_sub(float64(1), v113))
	goto L42
L41:
	;
	v118 = v112
	goto L42
L42:
	;
	if base.F64_gt(v118, float64(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v121 = F_clamp_row_est(m, v118)
	mBase = m.M
	v147 = v121
	goto L25
L44:
	;
	goto L45
L45:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v122 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v125 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v125)
	v147 = float64(200)
	goto L25
L47:
	;
	goto L48
L48:
	;
	v128 = *(*float64)(unsafe.Add(mBase, uint32(v122)+120))
	if base.F64_le(v128, float64(0)) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v131)
	v147 = float64(200)
	goto L25
L50:
	;
	goto L51
L51:
	;
	if base.F64_lt(v118, float64(0)) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v138 = F_clamp_row_est(m, base.F64_mul(v128, base.F64_neg(v118)))
	mBase = m.M
	v147 = v138
	goto L25
L53:
	;
	goto L54
L54:
	;
	if base.F64_lt(v128, float64(200)) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v141 = F_clamp_row_est(m, v128)
	mBase = m.M
	v147 = v141
	goto L25
L56:
	;
	goto L57
L57:
	;
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v142)
	v147 = float64(200)
	goto L25
L58:
	;
	F_free_attstatsslot(m, v16+int32(76))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L14
	} else {
		goto L132
	}
L59:
	;
	v238 = float64(0)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	if v239 <= int32(0) {
		v334 = v238
		goto L79
	} else {
		goto L80
	}
L60:
	;
	if v155 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v160 = v16 + int32(12)
	F_fmgr_info(m, v44, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L14
	} else {
		goto L62
	}
L62:
	;
	v163 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+72)) = uint8(v163)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+64)) = uint8(v163)
	v167 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+58)) = uint16(v167)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v163)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v160
	if l5 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	if v177 <= int32(0) {
		goto L59
	} else {
		goto L67
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = l3
	goto L63
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = l3
	goto L63
L67:
	;
	v185 = int32(0)
	goto L68
L68:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194+v185<<(uint(int32(2))%32))))
	if l5 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L59
L70:
	;
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v201)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v207 = m.T0[v206].(func(*base.Module, int32) int32)(m, v16+int32(40))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L14
	} else {
		goto L74
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v198
	goto L70
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v198
	goto L70
L74:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)))
	v210 = int32(0)
	if v209|base.B2i32(v207 == v210) == v210 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v219 = *(*float32)(unsafe.Add(mBase, uint32(v215+v185<<(uint(int32(2))%32))))
	v455 = base.F64_promote_f32(v219)
	goto L58
L76:
	;
	goto L77
L77:
	;
	v222 = v185 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	if v222 < v223 {
		v185 = v222
		goto L68
	} else {
		goto L78
	}
L78:
	;
	goto L69
L79:
	;
	v339 = base.F64_sub(base.F64_sub(float64(1), v334), v26)
	if base.F64_lt(v339, float64(0)) != 0 {
		v347 = v11
		goto L91
	} else {
		goto L92
	}
L80:
	;
	v243 = v239 & int32(3)
	v244 = int32(0)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	if base.Ui32(int32(4)) <= base.Ui32(v239) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v254 = int32(0)
	v256 = v244
	v262 = v238
	goto L84
L82:
	;
	v291 = v244
	v297 = v238
	goto L83
L83:
	;
	v301 = v244
	v304 = v291
	v310 = v297
	goto L88
L84:
	;
	v267 = v245 + v256<<(uint(int32(2))%32)
	v268 = *(*float32)(unsafe.Add(mBase, uint32(v267)))
	v271 = *(*float32)(unsafe.Add(mBase, uint32(v267)+4))
	v274 = *(*float32)(unsafe.Add(mBase, uint32(v267)+8))
	v277 = *(*float32)(unsafe.Add(mBase, uint32(v267)+12))
	v279 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v262, base.F64_promote_f32(v268)), base.F64_promote_f32(v271)), base.F64_promote_f32(v274)), base.F64_promote_f32(v277))
	v280 = int32(4)
	v281 = v256 + v280
	v283 = v254 + v280
	if v283 != v239&int32(2147483644) {
		v254 = v283
		v256 = v281
		v262 = v279
		goto L84
	} else {
		goto L86
	}
L85:
	;
	if v243 == int32(0) {
		v334 = v279
		goto L79
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v291 = v281
	v297 = v279
	goto L83
L88:
	;
	v316 = *(*float32)(unsafe.Add(mBase, uint32(v245+v304<<(uint(int32(2))%32))))
	v318 = base.F64_add(v310, base.F64_promote_f32(v316))
	v319 = int32(1)
	v322 = v301 + v319
	if v322 != v243 {
		v301 = v322
		v304 = v304 + v319
		v310 = v318
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v334 = v318
	goto L79
L90:
	;
	goto L89
L91:
	;
	v349 = v16 + int32(40)
	v350 = int32(0)
	v351 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v349))) = uint8(v350)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v355 != 0 {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	if base.F64_gt(v339, float64(1)) == int32(0) {
		v347 = v339
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v347 = float64(1)
	goto L91
L94:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	v426 = base.F64_sub(v423, base.F64_convert_i32_s(v424))
	if base.F64_gt(v426, float64(1)) != 0 {
		goto L127
	} else {
		goto L128
	}
L95:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v393 != 0 {
		goto L109
	} else {
		goto L110
	}
L96:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+22)))
	v358 = v356 + v357
	v359 = *(*float32)(unsafe.Add(mBase, uint32(v358)+8))
	v361 = *(*float32)(unsafe.Add(mBase, uint32(v358)+16))
	v388 = base.F64_promote_f32(v361)
	v389 = base.F64_promote_f32(v359)
	goto L95
L97:
	;
	goto L98
L98:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v363 == int32(16) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v388 = float64(2)
	v389 = v351
	goto L95
L100:
	;
	goto L101
L101:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v367 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v374 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v367)+76))
	if v370 != int32(5) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v388 = float64(-1)
	v389 = v351
	goto L95
L105:
	;
	v388 = float64(0)
	v389 = v351
	goto L95
L106:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	if v377 != int32(6) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v374)+8)))
	switch v381 - int32(_a_F_var_eq_const_4) {
	case 0:
		goto L108
	default:
		goto L105
	case 5:
		v388 = float64(-1)
		v389 = v351
		goto L95
	}
L108:
	;
	v388 = float64(1)
	v389 = v351
	goto L95
L109:
	;
	v394 = base.F64_neg(base.F64_sub(float64(1), v389))
	goto L111
L110:
	;
	v394 = v388
	goto L111
L111:
	;
	if base.F64_gt(v394, float64(0)) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v397 = F_clamp_row_est(m, v394)
	mBase = m.M
	v423 = v397
	goto L94
L113:
	;
	goto L114
L114:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v398 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v401 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v349))) = uint8(v401)
	v423 = float64(200)
	goto L94
L116:
	;
	goto L117
L117:
	;
	v404 = *(*float64)(unsafe.Add(mBase, uint32(v398)+120))
	if base.F64_le(v404, float64(0)) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v407 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v349))) = uint8(v407)
	v423 = float64(200)
	goto L94
L119:
	;
	goto L120
L120:
	;
	if base.F64_lt(v394, float64(0)) != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v414 = F_clamp_row_est(m, base.F64_mul(v404, base.F64_neg(v394)))
	mBase = m.M
	v423 = v414
	goto L94
L122:
	;
	goto L123
L123:
	;
	if base.F64_lt(v404, float64(200)) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v417 = F_clamp_row_est(m, v404)
	mBase = m.M
	v423 = v417
	goto L94
L125:
	;
	goto L126
L126:
	;
	v418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v349))) = uint8(v418)
	v423 = float64(200)
	goto L94
L127:
	;
	v430 = base.F64_div(v347, v426)
	goto L129
L128:
	;
	v430 = v347
	goto L129
L129:
	;
	if v424 <= int32(0) {
		v455 = v430
		goto L58
	} else {
		goto L130
	}
L130:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v439 = *(*float32)(unsafe.Add(mBase, uint32(v433+v424<<(uint(int32(2))%32)-int32(4))))
	v440 = base.F64_promote_f32(v439)
	if base.F64_gt(v430, v440) == int32(0) {
		v455 = v430
		goto L58
	} else {
		goto L131
	}
L131:
	;
	v455 = v440
	goto L58
L132:
	;
	v472 = v455
	goto L6
L133:
	;
	v478 = base.F64_sub(base.F64_sub(float64(1), v472), v26)
	goto L135
L134:
	;
	v478 = v472
	goto L135
L135:
	;
	if base.F64_lt(v478, float64(0)) != 0 {
		v496 = float64(0)
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if base.F64_gt(v478, float64(1)) == int32(0) {
		v496 = v478
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v496 = float64(1)
	goto L1
}
