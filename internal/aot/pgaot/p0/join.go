package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_flatten_join_alias_vars_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
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
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 != int32(319) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return v349
L5:
	;
	v347 = F_expression_tree_mutator_impl(m, l0, int32(903), l1)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L30
	} else {
		goto L100
	}
L6:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v325 + int32(1)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v330)
	v334 = F_query_tree_mutator_impl(m, l0, int32(903), l1, int32(4))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L30
	} else {
		goto L99
	}
L7:
	;
	if v15 == int32(67) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v160 = F_expression_tree_mutator_impl(m, l0, int32(903), l1)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L30
	} else {
		goto L62
	}
L10:
	;
	if v15 != int32(6) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v22 != v23 {
		v349 = l0
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = int32(2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(v29)%32)-int32(4))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v35 != v29 {
		v349 = l0
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v38 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v47 = int32(0)
	v50 = v3
	v51 = v3
	goto L17
L15:
	;
	goto L16
L16:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v128+v38<<(uint(int32(2))%32)-int32(4))))
	v135 = F_copyObjectImpl(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L30
	} else {
		goto L48
	}
L17:
	;
	v55 = int32(0)
	if v41 == v55 {
		v65 = v55
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v43 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v59 <= v47 {
		v65 = int32(0)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v65 = v61 + v47<<(uint(int32(2))%32)
	goto L19
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v100 != 0 {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v82 = F_palloc0(m, int32(24))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v68 = int32(0)
	v78 = v68
	v79 = v68
	goto L23
L25:
	;
	goto L26
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v70 <= v47 {
		v78 = v50
		v79 = v51
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v65 == int32(0) {
		v78 = v50
		v79 = v51
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v77 = v74 + v47<<(uint(int32(2))%32)
	if v77 != 0 {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v78 = v50
	v79 = v51
	goto L23
L30:
	;
	return int32(0)
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(36)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v97 = F_add_nullingrels_if_needed(m, v96, v82, l0)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	return v97
L33:
	;
	v101 = F_copyObjectImpl(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v122 = v50
	v123 = v51
	goto L35
L35:
	;
	v47 = v47 + int32(1)
	v50 = v122
	v51 = v123
	goto L17
L36:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v103 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_IncrementVarSublevelsUp(m, v101, v103, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L30
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v107 == int32(6) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+44)) = v110
	goto L43
L42:
	;
	goto L43
L43:
	;
	v112 = F_flatten_join_alias_vars_mutator(m, v101, l1)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v114 = F_lappend(m, v51, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L30
	} else {
		goto L45
	}
L45:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v117 = F_copyObjectImpl(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L30
	} else {
		goto L46
	}
L46:
	;
	v119 = F_lappend(m, v50, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L30
	} else {
		goto L47
	}
L47:
	;
	v122 = v119
	v123 = v114
	goto L35
L48:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v137 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_IncrementVarSublevelsUp(m, v135, v137, int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L30
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v141 == int32(6) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+44)) = v144
	goto L55
L54:
	;
	goto L55
L55:
	;
	v146 = F_flatten_join_alias_vars_mutator(m, v135, l1)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L30
	} else {
		goto L56
	}
L56:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v148 != int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v156 = F_add_nullingrels_if_needed(m, v155, v146, l0)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L30
	} else {
		goto L61
	}
L58:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v151 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v152 = F_checkExprHasSubLink(m, v146)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L30
	} else {
		goto L60
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v152)
	goto L57
L61:
	;
	return v156
L62:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v160)+20))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v162 != v163 {
		v349 = v160
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v166 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v167 == v166 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if int32(0) <= v224 {
		goto L75
	} else {
		goto L76
	}
L65:
	;
	v224 = base.I32_ctz(v210) | v211<<(uint(int32(5))%32)
	goto L64
L66:
	;
	v224 = int32(-2)
	goto L64
L67:
	;
	v177 = base.I32_div_s(int32(0), int32(32))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v178 <= v177 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v181 = v167 + int32(8)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181+v177<<(uint(int32(2))%32))))
	v188 = v185 & int32(-1)
	if v188 != 0 {
		v210 = v188
		v211 = v177
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v190 = v177 + int32(1)
	if v190 == v178 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v193 = v190
	goto L71
L71:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v181+v193<<(uint(int32(2))%32))))
	if v200 != 0 {
		v210 = v200
		v211 = v193
		goto L65
	} else {
		goto L73
	}
L72:
	;
	goto L66
L73:
	;
	v202 = v193 + int32(1)
	if v202 != v178 {
		v193 = v202
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v228 = v224
	v229 = v166
	goto L78
L76:
	;
	v315 = v166
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v315
	return v160
L78:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v239 = int32(2)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v238+v228<<(uint(v239)%32)-int32(4))))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	if v245 == v239 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v315 = v254
	goto L77
L80:
	;
	if v167 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L81:
	;
	v248 = F_get_relids_for_join(m, v165, v228)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L30
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v252 = F_bms_add_member(m, v229, v228)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L30
	} else {
		goto L86
	}
L84:
	;
	v250 = F_bms_join(m, v229, v248)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L30
	} else {
		goto L85
	}
L85:
	;
	v254 = v250
	goto L80
L86:
	;
	v254 = v252
	goto L80
L87:
	;
	if int32(0) <= v310 {
		v228 = v310
		v229 = v254
		goto L78
	} else {
		goto L98
	}
L88:
	;
	v310 = base.I32_ctz(v296) | v297<<(uint(int32(5))%32)
	goto L87
L89:
	;
	v310 = int32(-2)
	goto L87
L90:
	;
	v261 = v228 + int32(1)
	v263 = base.I32_div_s(v261, int32(32))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v264 <= v263 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v267 = v167 + int32(8)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267+v263<<(uint(int32(2))%32))))
	v274 = v271 & (int32(-1) << (uint(v261) % 32))
	if v274 != 0 {
		v296 = v274
		v297 = v263
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v276 = v263 + int32(1)
	if v276 == v264 {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v279 = v276
	goto L94
L94:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v267+v279<<(uint(int32(2))%32))))
	if v286 != 0 {
		v296 = v286
		v297 = v279
		goto L88
	} else {
		goto L96
	}
L95:
	;
	goto L89
L96:
	;
	v288 = v279 + int32(1)
	if v288 != v264 {
		v279 = v288
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	goto L79
L99:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+39)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v338 = v336 | v337
	*(*uint8)(unsafe.Add(mBase, uint32(v334)+39)) = uint8(v338)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v329)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v341 - int32(1)
	return v334
L100:
	;
	v349 = v347
	goto L4
}
func F_generate_join_implied_equalities(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v358 int32
	_ = v358
	var v371 int32
	_ = v371
	v6 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.Ui32(int32(5)) < base.Ui32(v15) {
		v32 = v14
		v33 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l4 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	if int32(1)<<(uint(v15)%32)&int32(44) == int32(0) {
		v32 = v14
		v33 = l1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v25 = F_bms_union(m, l2, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v30 = F_add_outer_joins_to_relids(m, l0, v25, l4, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v32 = v24
	v33 = v30
	goto L1
L7:
	;
	if v191 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L8:
	;
	v183 = F_get_common_eclass_indexes(m, l0, v32, l2)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L43
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	if v36 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v33 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if v95 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v95 = base.I32_ctz(v81) | v82<<(uint(int32(5))%32)
	goto L11
L13:
	;
	v95 = int32(-2)
	goto L11
L14:
	;
	v48 = base.I32_div_s(int32(0), int32(32))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v49 <= v48 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = v33 + int32(8)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v48<<(uint(int32(2))%32))))
	v59 = v56 & int32(-1)
	if v59 != 0 {
		v81 = v59
		v82 = v48
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v61 = v48 + int32(1)
	if v61 == v49 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v64 = v61
	goto L18
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v52+v64<<(uint(int32(2))%32))))
	if v71 != 0 {
		v81 = v71
		v82 = v64
		goto L12
	} else {
		goto L20
	}
L19:
	;
	goto L13
L20:
	;
	v73 = v64 + int32(1)
	if v73 != v49 {
		v64 = v73
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v191 = v6
	goto L7
L23:
	;
	goto L24
L24:
	;
	v102 = v95
	v104 = v6
	goto L25
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v102 == v111 {
		v124 = v104
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v191 = v124
	goto L7
L27:
	;
	if v33 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v102<<(uint(int32(2))%32))))
	if v117 == int32(0) {
		v124 = v104
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117)+136))
	v121 = F_bms_add_members(m, v104, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v124 = v121
	goto L27
L31:
	;
	if int32(0) < v180 {
		v102 = v180
		v104 = v124
		goto L25
	} else {
		goto L42
	}
L32:
	;
	v180 = base.I32_ctz(v166) | v167<<(uint(int32(5))%32)
	goto L31
L33:
	;
	v180 = int32(-2)
	goto L31
L34:
	;
	v131 = v102 + int32(1)
	v133 = base.I32_div_s(v131, int32(32))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v134 <= v133 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v137 = v33 + int32(8)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137+v133<<(uint(int32(2))%32))))
	v144 = v141 & (int32(-1) << (uint(v131) % 32))
	if v144 != 0 {
		v166 = v144
		v167 = v133
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v146 = v133 + int32(1)
	if v146 == v134 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v149 = v146
	goto L38
L38:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v137+v149<<(uint(int32(2))%32))))
	if v156 != 0 {
		v166 = v156
		v167 = v149
		goto L32
	} else {
		goto L40
	}
L39:
	;
	goto L33
L40:
	;
	v158 = v149 + int32(1)
	if v158 != v134 {
		v149 = v158
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L26
L43:
	;
	v191 = v183
	goto L7
L44:
	;
	if int32(0) <= v254 {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	v254 = base.I32_ctz(v240) | v241<<(uint(int32(5))%32)
	goto L44
L46:
	;
	v254 = int32(-2)
	goto L44
L47:
	;
	v207 = base.I32_div_s(int32(0), int32(32))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v208 <= v207 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v211 = v191 + int32(8)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+v207<<(uint(int32(2))%32))))
	v218 = v215 & int32(-1)
	if v218 != 0 {
		v240 = v218
		v241 = v207
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v220 = v207 + int32(1)
	if v220 == v208 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v223 = v220
	goto L51
L51:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v211+v223<<(uint(int32(2))%32))))
	if v230 != 0 {
		v240 = v230
		v241 = v223
		goto L45
	} else {
		goto L53
	}
L52:
	;
	goto L46
L53:
	;
	v232 = v223 + int32(1)
	if v232 != v208 {
		v223 = v232
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v261 = v254
	v267 = v6
	goto L58
L56:
	;
	v371 = v6
	goto L57
L57:
	;
	return v371
L58:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271+v261<<(uint(int32(2))%32))))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+40)))
	if v276 != 0 {
		v301 = v267
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v371 = v301
	goto L57
L60:
	;
	if v191 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L61:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
	if v277 == int32(0) {
		v301 = v267
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v280 < int32(2) {
		v301 = v267
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+42)))
	if v283 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v298 = F_list_concat(m, v267, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L73
	}
L65:
	;
	v295 = F_generate_join_implied_equalities_broken(m, l0, v275, v33, l2, v32, l3)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L72
	}
L66:
	;
	v286 = F_generate_join_implied_equalities_normal(m, l0, v275, l1, l2, v14)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v291 = int32(0)
	if v283 == v291 {
		v297 = v291
		goto L64
	} else {
		goto L71
	}
L69:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+42)))
	if v288&int32(1) != 0 {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v297 = v286
	goto L64
L71:
	;
	goto L65
L72:
	;
	v297 = v295
	goto L64
L73:
	;
	v301 = v298
	goto L60
L74:
	;
	if int32(0) <= v358 {
		v261 = v358
		v267 = v301
		goto L58
	} else {
		goto L85
	}
L75:
	;
	v358 = base.I32_ctz(v344) | v345<<(uint(int32(5))%32)
	goto L74
L76:
	;
	v358 = int32(-2)
	goto L74
L77:
	;
	v309 = v261 + int32(1)
	v311 = base.I32_div_s(v309, int32(32))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v312 <= v311 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v315 = v191 + int32(8)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315+v311<<(uint(int32(2))%32))))
	v322 = v319 & (int32(-1) << (uint(v309) % 32))
	if v322 != 0 {
		v344 = v322
		v345 = v311
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v324 = v311 + int32(1)
	if v324 == v312 {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v327 = v324
	goto L81
L81:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v315+v327<<(uint(int32(2))%32))))
	if v334 != 0 {
		v344 = v334
		v345 = v327
		goto L75
	} else {
		goto L83
	}
L82:
	;
	goto L76
L83:
	;
	v336 = v327 + int32(1)
	if v336 != v312 {
		v327 = v336
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	goto L59
}
func F_generate_join_implied_equalities_broken(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	v7 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l5)+224))
	v247 = F_adjust_appendrel_attrs_multilevel(m, l0, v211, l5, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L58
	} else {
		goto L67
	}
L2:
	;
	return v241
L3:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v13 <= v12 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v241 = int32(0)
	goto L2
L6:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v215&int32(-2) != int32(2) {
		goto L61
	} else {
		goto L62
	}
L7:
	;
	v211 = v7
	goto L6
L8:
	;
	goto L9
L9:
	;
	v17 = v12
	v22 = v7
	goto L10
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v17<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	v32 = int32(0)
	if v31 == v32 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v211 = v200
	goto L6
L12:
	;
	v202 = v17 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v202 < v203 {
		v17 = v202
		v22 = v200
		goto L10
	} else {
		goto L60
	}
L13:
	;
	if v85 == int32(0) {
		v200 = v22
		goto L12
	} else {
		goto L27
	}
L14:
	;
	v85 = int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	if l2 == int32(0) {
		v76 = v32
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v85 = v76
	goto L13
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v42 < v41 {
		v76 = v32
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v44 = int32(1)
	if v41 <= v44 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v47 = v44
	goto L22
L21:
	;
	v47 = v41
	goto L22
L22:
	;
	v48 = int32(8)
	v53 = int32(0)
	goto L23
L23:
	;
	v60 = v53 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v31+v48+v60)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+(l2+v48))))
	v67 = v62 & (v64 ^ int32(-1))
	v69 = base.B2i32(v67 == int32(0))
	if v67 != 0 {
		v76 = v69
		goto L17
	} else {
		goto L25
	}
L24:
	;
	v76 = v69
	goto L17
L25:
	;
	v71 = v53 + int32(1)
	if v71 != v47 {
		v53 = v71
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v88 = int32(0)
	if v31 == v88 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v141 != 0 {
		v200 = v22
		goto L12
	} else {
		goto L42
	}
L29:
	;
	v141 = int32(1)
	goto L28
L30:
	;
	goto L31
L31:
	;
	if l3 == int32(0) {
		v132 = v88
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v141 = v132
	goto L28
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v98 < v97 {
		v132 = v88
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v100 = int32(1)
	if v97 <= v100 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v103 = v100
	goto L37
L36:
	;
	v103 = v97
	goto L37
L37:
	;
	v104 = int32(8)
	v109 = int32(0)
	goto L38
L38:
	;
	v116 = v109 << (uint(int32(2)) % 32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v31+v104+v116)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+(l3+v104))))
	v123 = v118 & (v120 ^ int32(-1))
	v125 = base.B2i32(v123 == int32(0))
	if v123 != 0 {
		v132 = v125
		goto L32
	} else {
		goto L40
	}
L39:
	;
	v132 = v125
	goto L32
L40:
	;
	v127 = v109 + int32(1)
	if v127 != v103 {
		v109 = v127
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v142 = int32(0)
	if v31 == v142 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v195 != 0 {
		v200 = v22
		goto L12
	} else {
		goto L57
	}
L44:
	;
	v195 = int32(1)
	goto L43
L45:
	;
	goto L46
L46:
	;
	if l4 == int32(0) {
		v186 = v142
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v195 = v186
	goto L43
L48:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v152 < v151 {
		v186 = v142
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v154 = int32(1)
	if v151 <= v154 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v157 = v154
	goto L52
L51:
	;
	v157 = v151
	goto L52
L52:
	;
	v158 = int32(8)
	v163 = int32(0)
	goto L53
L53:
	;
	v170 = v163 << (uint(int32(2)) % 32)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v31+v158+v170)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170+(l4+v158))))
	v177 = v172 & (v174 ^ int32(-1))
	v179 = base.B2i32(v177 == int32(0))
	if v177 != 0 {
		v186 = v179
		goto L47
	} else {
		goto L55
	}
L54:
	;
	v186 = v179
	goto L47
L55:
	;
	v181 = v163 + int32(1)
	if v181 != v157 {
		v163 = v181
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v196 = F_lappend(m, v22, v30)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	return int32(0)
L59:
	;
	v200 = v196
	goto L12
L60:
	;
	goto L11
L61:
	;
	if v215 != int32(5) {
		v241 = v211
		goto L2
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v211 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	if v211 == int32(0) {
		v241 = v211
		goto L2
	} else {
		goto L65
	}
L65:
	;
	goto L1
L66:
	;
	goto L5
L67:
	;
	return v247
}
func F_has_join_restriction(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	v6 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v7 != 0 {
		v386 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v386
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v8 != 0 {
		v386 = v6
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v146 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v12 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v18 = int32(0)
	goto L7
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v18<<(uint(int32(2))%32))))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v27 = int32(0)
	if v20 == v27 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L4
L9:
	;
	v138 = v18 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v138 < v139 {
		v18 = v138
		goto L7
	} else {
		goto L38
	}
L10:
	;
	if v80 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L11:
	;
	v80 = int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v26 == int32(0) {
		v71 = v27
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v80 = v71
	goto L10
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v37 < v36 {
		v71 = v27
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v39 = int32(1)
	if v36 <= v39 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v42 = v39
	goto L19
L18:
	;
	v42 = v36
	goto L19
L19:
	;
	v43 = int32(8)
	v48 = int32(0)
	goto L20
L20:
	;
	v55 = v48 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v20+v43+v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+(v26+v43))))
	v62 = v57 & (v59 ^ int32(-1))
	v64 = base.B2i32(v62 == int32(0))
	if v62 != 0 {
		v71 = v64
		goto L14
	} else {
		goto L22
	}
L21:
	;
	v71 = v64
	goto L14
L22:
	;
	v66 = v48 + int32(1)
	if v66 != v42 {
		v48 = v66
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v85 = int32(0)
	v92 = base.B2i32(v83|v84 == v85)
	if v83 == v85 {
		v131 = v92
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v131 != 0 {
		goto L9
	} else {
		goto L37
	}
L26:
	;
	goto L25
L27:
	;
	if v84 == int32(0) {
		v131 = v92
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v98 != v99 {
		v131 = int32(0)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v101 = int32(1)
	if v98 <= v101 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v104 = v101
	goto L32
L31:
	;
	v104 = v98
	goto L32
L32:
	;
	v105 = int32(8)
	v110 = int32(0)
	goto L33
L33:
	;
	v118 = v110 << (uint(int32(2)) % 32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v83+v105+v118)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+(v84+v105))))
	v123 = base.B2i32(v120 == v122)
	if v122 != v120 {
		v131 = v123
		goto L26
	} else {
		goto L35
	}
L34:
	;
	v131 = v123
	goto L26
L35:
	;
	v126 = v110 + int32(1)
	if v126 != v104 {
		v110 = v126
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	return int32(1)
L38:
	;
	goto L8
L39:
	;
	v386 = int32(0)
	goto L1
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v149 <= int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v156 = int32(0)
	goto L42
L42:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	v159 = int32(2)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158+v156<<(uint(v159)%32))))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	if v163 == v159 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L39
L44:
	;
	v375 = v156 + int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v375 < v376 {
		v156 = v375
		goto L42
	} else {
		goto L108
	}
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v168 = int32(0)
	if v166 == v168 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v221 != 0 {
		goto L60
	} else {
		goto L61
	}
L47:
	;
	v221 = int32(1)
	goto L46
L48:
	;
	goto L49
L49:
	;
	if v167 == int32(0) {
		v212 = v168
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v221 = v212
	goto L46
L51:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v178 < v177 {
		v212 = v168
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v180 = int32(1)
	if v177 <= v180 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v183 = v180
	goto L55
L54:
	;
	v183 = v177
	goto L55
L55:
	;
	v184 = int32(8)
	v189 = int32(0)
	goto L56
L56:
	;
	v196 = v189 << (uint(int32(2)) % 32)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v166+v184+v196)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196+(v167+v184))))
	v203 = v198 & (v200 ^ int32(-1))
	v205 = base.B2i32(v203 == int32(0))
	if v203 != 0 {
		v212 = v205
		goto L50
	} else {
		goto L58
	}
L57:
	;
	v212 = v205
	goto L50
L58:
	;
	v207 = v189 + int32(1)
	if v207 != v183 {
		v189 = v207
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v224 = int32(0)
	if v222 == v224 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	v278 = int32(1)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v281 = int32(0)
	if v279 == v281 {
		v322 = v281
		goto L79
	} else {
		goto L80
	}
L63:
	;
	if v277 != 0 {
		goto L44
	} else {
		goto L77
	}
L64:
	;
	v277 = int32(1)
	goto L63
L65:
	;
	goto L66
L66:
	;
	if v223 == int32(0) {
		v268 = v224
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v277 = v268
	goto L63
L68:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v234 < v233 {
		v268 = v224
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v236 = int32(1)
	if v233 <= v236 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v239 = v236
	goto L72
L71:
	;
	v239 = v233
	goto L72
L72:
	;
	v240 = int32(8)
	v245 = int32(0)
	goto L73
L73:
	;
	v252 = v245 << (uint(int32(2)) % 32)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v222+v240+v252)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252+(v223+v240))))
	v259 = v254 & (v256 ^ int32(-1))
	v261 = base.B2i32(v259 == int32(0))
	if v259 != 0 {
		v268 = v261
		goto L67
	} else {
		goto L75
	}
L74:
	;
	v268 = v261
	goto L67
L75:
	;
	v263 = v245 + int32(1)
	if v263 != v239 {
		v245 = v263
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L62
L78:
	;
	if v322 != 0 {
		v386 = v278
		goto L1
	} else {
		goto L92
	}
L79:
	;
	goto L78
L80:
	;
	if v280 == int32(0) {
		v322 = v281
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	if v290 < v291 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v293 = v290
	goto L84
L83:
	;
	v293 = v291
	goto L84
L84:
	;
	if v293 <= int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v296 = int32(1)
	goto L87
L86:
	;
	v296 = v293
	goto L87
L87:
	;
	v297 = int32(8)
	v302 = int32(0)
	goto L88
L88:
	;
	v309 = v302 << (uint(int32(2)) % 32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v280+v297+v309)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309+(v279+v297))))
	v314 = v311 & v313
	v316 = base.B2i32(v314 != int32(0))
	if v314 != 0 {
		v322 = v316
		goto L79
	} else {
		goto L90
	}
L89:
	;
	v322 = v316
	goto L79
L90:
	;
	v318 = v302 + int32(1)
	if v318 != v296 {
		v302 = v318
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v328 = int32(0)
	if v326 == v328 {
		v369 = v328
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v369 != 0 {
		v386 = v278
		goto L1
	} else {
		goto L107
	}
L94:
	;
	goto L93
L95:
	;
	if v327 == int32(0) {
		v369 = v328
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v337 < v338 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v340 = v337
	goto L99
L98:
	;
	v340 = v338
	goto L99
L99:
	;
	if v340 <= int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v343 = int32(1)
	goto L102
L101:
	;
	v343 = v340
	goto L102
L102:
	;
	v344 = int32(8)
	v349 = int32(0)
	goto L103
L103:
	;
	v356 = v349 << (uint(int32(2)) % 32)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v327+v344+v356)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356+(v326+v344))))
	v361 = v358 & v360
	v363 = base.B2i32(v361 != int32(0))
	if v361 != 0 {
		v369 = v363
		goto L94
	} else {
		goto L105
	}
L104:
	;
	v369 = v363
	goto L94
L105:
	;
	v365 = v349 + int32(1)
	if v365 != v343 {
		v349 = v365
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	goto L44
L108:
	;
	goto L43
}
