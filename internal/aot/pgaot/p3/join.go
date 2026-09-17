package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fix_join_expr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
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
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 float64
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 float64
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 float64
	_ = v224
	var v225 float64
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 float64
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 float64
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v291 int32
	_ = v291
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v3 {
		v291 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v291
L2:
	;
	v19 = l0
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v31 != int32(319) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v291 = int32(0)
	goto L1
L5:
	;
	if v276 != 0 {
		v19 = v276
		goto L3
	} else {
		goto L84
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L21
	} else {
		goto L81
	}
L7:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v147 == int32(0))|base.B2i32(v31 == int32(7)) != 0 {
		goto L51
	} else {
		goto L52
	}
L8:
	;
	if v31 != int32(6) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v124 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	if v36 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v76 != 0 {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v72 = F_search_indexed_tlist_for_var(m, v19, v67, int32(-2), v70, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L21
	} else {
		goto L26
	}
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v37 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v64 == int32(0) {
		goto L12
	} else {
		goto L25
	}
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v38 == int32(0) {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v41 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v44 == v41 {
		v67 = v38
		goto L13
	} else {
		goto L20
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v52
	F_errmsg_internal(m, int32(_a_F_fix_join_expr_mutator_0), v15)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_fix_join_expr_mutator_1), int32(3148), int32(_a_F_fix_join_expr_mutator_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v67 = v64
	goto L13
L26:
	;
	if v72 != 0 {
		v291 = v72
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L12
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v80 = F_search_indexed_tlist_for_var(m, v19, v76, int32(-1), v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L21
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v83 == v84 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v80 != 0 {
		v291 = v80
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v87 = F_palloc(m, int32(48))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L21
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L21
	} else {
		goto L38
	}
L36:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v19)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+40)) = v89
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v19)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+32)) = v91
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+24)) = v93
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v101 + v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v87)+36))
	if v105 == int32(0) {
		v291 = v87
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+36)) = v108 + v105
	v291 = v87
	goto L1
L38:
	;
	F_errmsg_internal(m, int32(_a_F_fix_join_expr_mutator_3), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L21
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_fix_join_expr_mutator_1), int32(3186), int32(_a_F_fix_join_expr_mutator_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L21
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v135 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+8)))
	if v127 != int32(1) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v132 = F_search_indexed_tlist_for_phv(m, v19, v124, int32(-2), v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	if v132 != 0 {
		v291 = v132
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v276 = v146
	goto L5
L47:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+8)))
	if v138 != int32(1) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v143 = F_search_indexed_tlist_for_phv(m, v19, v135, int32(-1), v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	if v143 != 0 {
		v291 = v143
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v171 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+9)))
	if v153&int32(1) == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v159 = F_tlist_member(m, v19, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L21
	} else {
		goto L54
	}
L54:
	;
	if v159 == int32(0) {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v164 = F_makeVarFromTargetEntry(m, int32(-2), v159)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L21
	} else {
		goto L56
	}
L56:
	;
	v166 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v164)+40)) = uint16(v166)
	*(*int32)(unsafe.Add(mBase, uint32(v164)+36)) = v166
	v291 = v164
	goto L1
L57:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_fix_expr_common(m, v255, v19)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L21
	} else {
		goto L79
	}
L58:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v193 != int32(24) {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+9)))
	if v174 != int32(1) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v177 == int32(7) {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v181 = F_tlist_member(m, v19, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L21
	} else {
		goto L62
	}
L62:
	;
	if v181 == int32(0) {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v186 = F_makeVarFromTargetEntry(m, int32(-1), v181)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L21
	} else {
		goto L64
	}
L64:
	;
	v188 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v186)+40)) = uint16(v188)
	*(*int32)(unsafe.Add(mBase, uint32(v186)+36)) = v188
	v291 = v186
	goto L1
L65:
	;
	if v193 != int32(8) {
		goto L57
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v202 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v204 = int32(0)
	v207 = v204
	v209 = v204
	v215 = float64(0)
	goto L70
L68:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v199 = F_fix_param_node(m, v198, v19)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	v291 = v199
	goto L1
L70:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219+v209<<(uint(int32(2))%32))))
	v224 = *(*float64)(unsafe.Add(mBase, uint32(v223)+56))
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v223)+64))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v203)+360))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v229 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v226+v227-v229))) = uint8(v229)
	v234 = base.F64_add(v224, base.F64_mul(v202, v225))
	v236 = int32(0)
	v240 = base.B2i32(base.F64_ge(v215, v234) == v236) & base.B2i32(v207 != v236)
	if v240 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v203)+364))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	v250 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v247+v248-v250))) = uint8(v250)
	v276 = v242
	goto L5
L72:
	;
	v241 = v215
	goto L74
L73:
	;
	v241 = v234
	goto L74
L74:
	;
	if v240 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v242 = v207
	goto L77
L76:
	;
	v242 = v223
	goto L77
L77:
	;
	v244 = v209 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v244 < v245 {
		v207 = v242
		v209 = v244
		v215 = v241
		goto L70
	} else {
		goto L78
	}
L78:
	;
	goto L71
L79:
	;
	v259 = F_expression_tree_mutator_impl(m, v19, int32(835), l1)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	v291 = v259
	goto L1
L81:
	;
	F_errmsg_internal(m, int32(_a_F_fix_join_expr_mutator_4), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L21
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_fix_join_expr_mutator_1), int32(3145), int32(_a_F_fix_join_expr_mutator_2))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L21
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	goto L4
}
func F_get_join_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	v19 = m.G0
	v21 = v19 - int32(144)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v24 = F_list_member(m, v23, l7)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v24 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = int32(0)
	base.MemoryFill(m, v21+int32(12), v30, int32(132))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v30 < v33 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v21 + int32(144)
	return
L6:
	;
	v36 = int32(4)
	v43 = v21 + int32(16)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)))
	v57 = v44
	v58 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_get_index_paths(m, l0, l1, l2, v21+int32(12), l6)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L64
	}
L9:
	;
	v64 = v58 << (uint(int32(2)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l4+v36+v64)))
	if v66 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)) = uint8(v303)
	goto L8
L11:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v64+(l5+v36))))
	if v180 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v69 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v72 = v64 + v43
	v79 = int32(0)
	goto L14
L14:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v79<<(uint(int32(2))%32))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v99 = int32(0)
	if v98 == v99 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L11
L16:
	;
	if v152 != 0 {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v152 = int32(1)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if l7 == int32(0) {
		v145 = v99
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v152 = v145
	goto L16
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v109 < v108 {
		v145 = v99
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v111 = int32(1)
	if v108 <= v111 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v114 = v111
	goto L25
L24:
	;
	v114 = v108
	goto L25
L25:
	;
	v115 = int32(8)
	v120 = int32(0)
	goto L26
L26:
	;
	v127 = v120 << (uint(int32(2)) % 32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v98+v115+v127)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l7+v115+v127)))
	v134 = v129 & (v131 ^ int32(-1))
	v136 = base.B2i32(v134 == int32(0))
	if v134 != 0 {
		v145 = v136
		goto L20
	} else {
		goto L28
	}
L27:
	;
	v145 = v136
	goto L20
L28:
	;
	v138 = v120 + int32(1)
	if v138 != v114 {
		v120 = v138
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v154 = F_lappend(m, v153, v96)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v158 = v79 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v158 < v159 {
		v79 = v158
		goto L14
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v154
	goto L32
L34:
	;
	goto L15
L35:
	;
	v295 = v64 + v43
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v64+(l3+v36))))
	v299 = F_list_concat(m, v296, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L59
	}
L36:
	;
	v183 = int32(0)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v184 <= v183 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v192 = v183
	goto L38
L38:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v205+v192<<(uint(int32(2))%32))))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	v212 = int32(0)
	if v211 == v212 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v272 = v64 + v43
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v274 = F_lappend(m, v273, v209)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L58
	}
L40:
	;
	if v265 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L41:
	;
	v265 = int32(1)
	goto L40
L42:
	;
	goto L43
L43:
	;
	if l7 == int32(0) {
		v258 = v212
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v265 = v258
	goto L40
L45:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v222 < v221 {
		v258 = v212
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v224 = int32(1)
	if v221 <= v224 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v227 = v224
	goto L49
L48:
	;
	v227 = v221
	goto L49
L49:
	;
	v228 = int32(8)
	v233 = int32(0)
	goto L50
L50:
	;
	v240 = v233 << (uint(int32(2)) % 32)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v211+v228+v240)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l7+v228+v240)))
	v247 = v242 & (v244 ^ int32(-1))
	v249 = base.B2i32(v247 == int32(0))
	if v247 != 0 {
		v258 = v249
		goto L44
	} else {
		goto L52
	}
L51:
	;
	v258 = v249
	goto L44
L52:
	;
	v251 = v233 + int32(1)
	if v251 != v227 {
		v233 = v251
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v269 = v192 + int32(1)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v269 < v270 {
		v192 = v269
		goto L38
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L39
L57:
	;
	goto L35
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v274
	goto L35
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v299
	if v299 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v303 = int32(1)
	goto L62
L61:
	;
	v303 = v57
	goto L62
L62:
	;
	v305 = v58 + int32(1)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v305 < v306 {
		v57 = v303
		v58 = v305
		goto L9
	} else {
		goto L63
	}
L63:
	;
	goto L10
L64:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v332 = F_lappend(m, v331, l7)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v332
	goto L5
}
func F_join_clause_is_movable_into(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v5 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v160
L2:
	;
	if v59 == int32(0) {
		v160 = v4
		goto L1
	} else {
		goto L16
	}
L3:
	;
	v59 = int32(1)
	goto L2
L4:
	;
	goto L5
L5:
	;
	if l2 == int32(0) {
		v52 = v4
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v59 = v52
	goto L2
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v16 < v15 {
		v52 = v4
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v18 = int32(1)
	if v15 <= v18 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v21 = v18
	goto L11
L10:
	;
	v21 = v15
	goto L11
L11:
	;
	v22 = int32(8)
	v27 = int32(0)
	goto L12
L12:
	;
	v34 = v27 << (uint(int32(2)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v5+v22+v34)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2+v22+v34)))
	v41 = v36 & (v38 ^ int32(-1))
	v43 = base.B2i32(v41 == int32(0))
	if v41 != 0 {
		v52 = v43
		goto L6
	} else {
		goto L14
	}
L13:
	;
	v52 = v43
	goto L6
L14:
	;
	v45 = v27 + int32(1)
	if v45 != v21 {
		v27 = v45
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v63 = int32(0)
	if base.B2i32(l1 == v63)|base.B2i32(v62 == v63) != 0 {
		v108 = v63
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v108 == int32(0) {
		v160 = v4
		goto L1
	} else {
		goto L30
	}
L18:
	;
	goto L17
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v73 < v74 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v76 = v73
	goto L22
L21:
	;
	v76 = v74
	goto L22
L22:
	;
	if v76 <= int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v79 = int32(1)
	goto L25
L24:
	;
	v79 = v76
	goto L25
L25:
	;
	v80 = int32(8)
	v85 = int32(0)
	goto L26
L26:
	;
	v92 = v85 << (uint(int32(2)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v62+v80+v92)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1+v80+v92)))
	v97 = v94 & v96
	v99 = base.B2i32(v97 != int32(0))
	if v97 != 0 {
		v108 = v99
		goto L18
	} else {
		goto L28
	}
L27:
	;
	v108 = v99
	goto L18
L28:
	;
	v101 = v85 + int32(1)
	if v101 != v79 {
		v85 = v101
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v112 = int32(0)
	if base.B2i32(l1 == v112)|base.B2i32(v111 == v112) != 0 {
		v157 = v112
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v160 = v157 ^ int32(1)
	goto L1
L32:
	;
	goto L31
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v122 < v123 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v125 = v122
	goto L36
L35:
	;
	v125 = v123
	goto L36
L36:
	;
	if v125 <= int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v128 = int32(1)
	goto L39
L38:
	;
	v128 = v125
	goto L39
L39:
	;
	v129 = int32(8)
	v134 = int32(0)
	goto L40
L40:
	;
	v141 = v134 << (uint(int32(2)) % 32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v111+v129+v141)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1+v129+v141)))
	v146 = v143 & v145
	v148 = base.B2i32(v146 != int32(0))
	if v146 != 0 {
		v157 = v148
		goto L32
	} else {
		goto L42
	}
L41:
	;
	v157 = v148
	goto L32
L42:
	;
	v150 = v134 + int32(1)
	if v150 != v128 {
		v134 = v150
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
}
func F_join_path_components(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 != l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L7
L2:
	;
	goto L3
L3:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v130 != 0 {
		goto L35
	} else {
		goto L36
	}
L4:
	;
	goto L3
L5:
	;
	v127 = F_strlen(m, v116)
	mBase = m.M
	goto L4
L7:
	;
	goto L8
L8:
	;
	v17 = int32(1023)
	if (l0^l1)&int32(3) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v120)
	goto L5
L10:
	;
	v101 = v96
	v102 = v97
	v103 = v98
	goto L31
L11:
	;
	if v91 == int32(0) {
		v116 = v89
		v117 = v90
		goto L9
	} else {
		goto L30
	}
L12:
	;
	v89 = l1
	v90 = l0
	v91 = v17
	goto L11
L13:
	;
	goto L14
L14:
	;
	v21 = int32(0)
	if base.B2i32(l1&int32(3) == v21)|int32(0) == v21 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v57 == int32(0) {
		v116 = v54
		v117 = v55
		goto L9
	} else {
		goto L24
	}
L16:
	;
	v33 = l1
	v34 = l0
	v35 = v17
	goto L19
L17:
	;
	goto L18
L18:
	;
	v54 = l1
	v55 = l0
	v56 = v17
	v57 = int32(1)
	goto L15
L19:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v37)
	if v37 == int32(0) {
		v96 = v33
		v97 = v34
		v98 = v35
		goto L10
	} else {
		goto L21
	}
L20:
	;
	v54 = v48
	v55 = v42
	v56 = v44
	v57 = v46
	goto L15
L21:
	;
	v41 = int32(1)
	v42 = v34 + v41
	v44 = v35 - v41
	v45 = int32(0)
	v46 = base.B2i32(v44 != v45)
	v48 = v33 + v41
	if v48&int32(3) == v45 {
		v54 = v48
		v55 = v42
		v56 = v44
		v57 = v46
		goto L15
	} else {
		goto L22
	}
L22:
	;
	if v44 != 0 {
		v33 = v48
		v34 = v42
		v35 = v44
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if base.B2i32(v60 == int32(0))|base.B2i32(base.Ui32(v56) < base.Ui32(int32(4))) != 0 {
		v89 = v54
		v90 = v55
		v91 = v56
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v67 = v54
	v68 = v55
	v69 = v56
	goto L26
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v75 = int32(-2139062144)
	if (int32(16843008)-v72|v72)&v75 != v75 {
		v96 = v67
		v97 = v68
		v98 = v69
		goto L10
	} else {
		goto L28
	}
L27:
	;
	v89 = v83
	v90 = v81
	v91 = v85
	goto L11
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v72
	v80 = int32(4)
	v81 = v68 + v80
	v83 = v67 + v80
	v85 = v69 - v80
	if base.Ui32(int32(3)) < base.Ui32(v85) {
		v67 = v83
		v68 = v81
		v69 = v85
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v96 = v89
	v97 = v90
	v98 = v91
	goto L10
L31:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v105)
	if v105 == int32(0) {
		v116 = v101
		v117 = v102
		goto L9
	} else {
		goto L33
	}
L32:
	;
	v116 = v112
	v117 = v110
	goto L9
L33:
	;
	v109 = int32(1)
	v110 = v102 + v109
	v112 = v101 + v109
	v114 = v103 - v109
	if v114 != 0 {
		v101 = v112
		v102 = v110
		v103 = v114
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v131 = F_strlen(m, l0)
	mBase = m.M
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l2
	if v132 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	m.G0 = v8 + int32(16)
	return
L38:
	;
	v136 = int32(_a_F_join_path_components_0)
	goto L40
L39:
	;
	v136 = int32(_a_F_join_path_components_1)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v136
	v142 = F_pg_snprintf(m, l0+v131, int32(1024)-v131, int32(_a_F_join_path_components_2), v8)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return
L42:
	;
	goto L37
}
