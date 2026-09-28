package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_finalize_grouping_exprs_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v322 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v322)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v325 = F_finalize_grouping_exprs_walker(m, v324, l1)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L23
	} else {
		goto L76
	}
L2:
	;
	return int32(0)
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v15 - int32(7) {
	case 0, 1:
		goto L2
	case 2:
		goto L7
	case 3:
		goto L6
	default:
		v246 = v15
		goto L5
	}
L4:
	;
	v305 = F_expression_tree_walker_impl(m, l0, int32(515), l1)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L23
	} else {
		goto L75
	}
L5:
	;
	if v246 != int32(67) {
		goto L4
	} else {
		goto L64
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v22 == v23 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v18 == v19 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v18 <= v19 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v232 = v22
	v241 = v23
	goto L12
L12:
	;
	if v241 < v232 {
		goto L2
	} else {
		goto L63
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v26 <= int32(0) {
		v211 = v3
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v218 = v22
	v224 = v3
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v224
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v232 = v218
	v241 = v228
	goto L12
L16:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v218 = v214
	v224 = v211
	goto L15
L17:
	;
	v30 = l1 + int32(12)
	v38 = v3
	v40 = v3
	goto L18
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v38<<(uint(int32(2))%32))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L23
	} else {
		goto L58
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v51 = F_flatten_join_alias_for_parser(m, v49, v47, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v55 = v47
	goto L22
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 == int32(6) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	return int32(0)
L24:
	;
	v55 = v51
	goto L22
L25:
	;
	goto L19
L26:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	if v161 == int32(0) {
		goto L25
	} else {
		goto L55
	}
L27:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v59 != v60 {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v106 != int32(1) {
		goto L25
	} else {
		goto L44
	}
L30:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v62 == int32(0) {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v65 <= int32(0) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v68 = int32(0)
	if v68 < v65 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v71 = v65
	goto L35
L34:
	;
	v71 = v68
	goto L35
L35:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v76 = int32(0)
	goto L36
L36:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v72+v76<<(uint(int32(2))%32))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v91 != int32(6) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L25
L38:
	;
	v104 = v76 + int32(1)
	if v104 != v71 {
		v76 = v104
		goto L36
	} else {
		goto L43
	}
L39:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v94 != v95 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+8)))
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+8)))
	if v97 != v98 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v90)+28))
	if v100 == int32(0) {
		v157 = v89
		goto L26
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	goto L37
L44:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v109 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v117 = v111 + v109<<(uint(int32(2))%32) - int32(4)
	goto L47
L46:
	;
	v117 = v30
	goto L47
L47:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v118 == int32(0) {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v121 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v122 <= v121 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	v127 = v121
	goto L50
L50:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137+v127<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v143 = F_equal(m, v55, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L23
	} else {
		goto L52
	}
L51:
	;
	goto L25
L52:
	;
	if v143 != 0 {
		v157 = v141
		goto L26
	} else {
		goto L53
	}
L53:
	;
	v146 = v127 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v146 < v147 {
		v127 = v146
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v164 = F_lappend_int(m, v40, v161)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	v167 = v38 + int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v167 < v168 {
		v38 = v167
		v40 = v164
		goto L18
	} else {
		goto L57
	}
L57:
	;
	v211 = v164
	goto L16
L58:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_finalize_grouping_exprs_walker_0), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v194 = F_exprLocation(m, v55)
	mBase = m.M
	F_parser_errposition(m, v193, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L23
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_finalize_grouping_exprs_walker_1), int32(1760), int32(_a_F_finalize_grouping_exprs_walker_2))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L23
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
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v246 = v243
	goto L5
L64:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v259 = int32(1)
	v260 = v258 + v259
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v260
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v262 != v259 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v285 = F_query_tree_walker_impl(m, l0, int32(515), l1, int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L23
	} else {
		goto L74
	}
L66:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v265 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v268 = v266
	goto L69
L68:
	;
	v268 = int32(0)
	goto L69
L69:
	;
	if v260 <= v268 {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v271 = F_copyObjectImpl(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L23
	} else {
		goto L71
	}
L71:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	F_IncrementVarSublevelsUp(m, v271, v273, int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L23
	} else {
		goto L72
	}
L72:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v278 = F_lappend(m, v277, v271)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L23
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v278
	goto L65
L74:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v287 - int32(1)
	return v285
L75:
	;
	return v305
L76:
	;
	v327 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v327)
	return v325
}
func F_gather_grouping_paths(m *base.Module, l0 int32, l1 int32) {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
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
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 float64
	_ = v159
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v163 int32
	_ = v163
	var v169 float64
	_ = v169
	var v173 float64
	_ = v173
	var v175 float64
	_ = v175
	var v177 float64
	_ = v177
	var v178 float64
	_ = v178
	var v187 float64
	_ = v187
	var v191 float64
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v19 = v17
	goto L3
L2:
	;
	v19 = int32(0)
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v20 < v19 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = F_list_copy_head(m, v16, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v24 = v16
	goto L6
L6:
	;
	F_generate_useful_gather_paths(m, l0, l1, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	v24 = v22
	goto L6
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if int32(0) < v29 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v42 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	m.G0 = v14 + int32(16)
	return
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v42<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
	v52 = v14 + int32(12)
	if v24 == v50 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L12
L15:
	;
	v202 = v42 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v202 < v203 {
		v42 = v202
		goto L13
	} else {
		goto L76
	}
L16:
	;
	if v130 != 0 {
		goto L15
	} else {
		goto L48
	}
L17:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v118
	v130 = int32(1)
	goto L16
L18:
	;
	if v24 != 0 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v24 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(0)
	v130 = int32(1)
	goto L16
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(0)
	v130 = int32(1)
	goto L16
L23:
	;
	goto L24
L24:
	;
	if v50 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v70
	v130 = v70
	goto L16
L26:
	;
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v74 = int32(0)
	if v74 < v73 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v77 = v73
	goto L30
L29:
	;
	v77 = v74
	goto L30
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v83 = int32(0)
	goto L31
L31:
	;
	if v83 < v78 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v94 = v90 + v83<<(uint(int32(2))%32)
	goto L35
L34:
	;
	v94 = int32(0)
	goto L35
L35:
	;
	if v83 == v77 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v77
	v130 = base.B2i32(v94 == int32(0))
	goto L16
L37:
	;
	goto L38
L38:
	;
	v100 = base.B2i32(v94 == int32(0))
	if v94 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v83
	v130 = v100
	goto L16
L40:
	;
	goto L41
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v104 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v83
	v130 = v100
	goto L16
L43:
	;
	goto L44
L44:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104+v83<<(uint(int32(2))%32))))
	if v108 != v112 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v83
	v130 = int32(0)
	goto L16
L46:
	;
	v83 = v83 + int32(1)
	goto L31
L48:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gather_grouping_paths[0])))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v49 == v33 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v142&int32(1) != 0 {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v142 = v132
	goto L49
L51:
	;
	goto L52
L52:
	;
	if v133 == int32(0) {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	v137 = int32(1)
	if v132&v137 == int32(0) {
		goto L15
	} else {
		goto L54
	}
L54:
	;
	v142 = v137
	goto L49
L55:
	;
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v155)+32))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+24))
	v161 = base.F64_convert_i32_s(v160)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gather_grouping_paths[1])))
	if v163 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L56:
	;
	v146 = v133
	goto L58
L57:
	;
	v146 = int32(0)
	goto L58
L58:
	;
	if v146 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v150 = F_create_sort_path(m, l1, v49, v24, float64(-1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v153 = F_create_incremental_sort_path(m, l0, l1, v49, v24, v133, float64(-1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L63
	}
L62:
	;
	v155 = v150
	goto L55
L63:
	;
	v155 = v153
	goto L55
L64:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14))) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v194 = F_create_gather_merge_path(m, l0, l1, v155, v193, v24, v14)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L74
	}
L65:
	;
	v169 = base.F64_add(base.F64_mul(v161, float64(-0.3)), float64(1))
	if base.F64_gt(v169, float64(0)) != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v175 = v161
	goto L67
L67:
	;
	v177 = float64(1e+100)
	v178 = base.F64_mul(v159, v175)
	if base.F64_gt(v178, v177)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v178)&int64(9223372036854775807))) != 0 {
		v191 = v177
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v173 = v169
	goto L70
L69:
	;
	v173 = math.Float64frombits(uint64(0x8000000000000000))
	goto L70
L70:
	;
	v175 = base.F64_add(v173, v161)
	goto L67
L71:
	;
	goto L64
L72:
	;
	v187 = float64(1)
	if base.F64_le(v178, v187) != 0 {
		v191 = v187
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v191 = base.F64_nearest(v178)
	goto L71
L74:
	;
	F_add_path(m, l1, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	goto L15
L76:
	;
	goto L14
}
