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
	var v52 int32
	_ = v52
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
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
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
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
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v357 = m.G0
	v359 = v357 - int32(16)
	m.G0 = v359
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v361 == int32(0) {
		v469 = v348
		goto L102
	} else {
		goto L103
	}
L5:
	;
	return v343
L6:
	;
	v341 = F_expression_tree_mutator_impl(m, l0, int32(904), l1)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L30
	} else {
		goto L98
	}
L7:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v319 + int32(1)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v324)
	v328 = F_query_tree_mutator_impl(m, l0, int32(904), l1, int32(4))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L30
	} else {
		goto L97
	}
L8:
	;
	if v15 == int32(67) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v154 = F_expression_tree_mutator_impl(m, l0, int32(904), l1)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L30
	} else {
		goto L60
	}
L11:
	;
	if v15 != int32(6) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v22 != v23 {
		v343 = l0
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = int32(2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(v29)%32)-int32(4))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v35 != v29 {
		v343 = l0
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v38 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v47 = int32(0)
	v50 = v3
	v52 = v3
	goto L18
L16:
	;
	goto L17
L17:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v126+v38<<(uint(int32(2))%32)-int32(4))))
	v133 = F_copyObjectImpl(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L30
	} else {
		goto L47
	}
L18:
	;
	v55 = int32(0)
	if v41 == v55 {
		v65 = v55
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v43 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v59 <= v47 {
		v65 = int32(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v65 = v61 + v47<<(uint(int32(2))%32)
	goto L20
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v94 != 0 {
		goto L32
	} else {
		goto L33
	}
L24:
	;
	v80 = F_palloc0(m, int32(24))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v68 = int32(0)
	v76 = v68
	v77 = v68
	goto L24
L26:
	;
	goto L27
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if base.B2i32(v65 == int32(0))|base.B2i32(v72 <= v47) != 0 {
		v76 = v50
		v77 = v52
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if v75 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v76 = v50
	v77 = v52
	goto L24
L30:
	;
	return int32(0)
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(36)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v87
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v92
	v348 = v80
	goto L4
L32:
	;
	v95 = F_copyObjectImpl(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	v119 = v50
	v120 = v52
	goto L34
L34:
	;
	v47 = v47 + int32(1)
	v50 = v119
	v52 = v120
	goto L18
L35:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v97 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_IncrementVarSublevelsUp(m, v95, v97, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L30
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v101 == int32(6) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+44)) = v104
	goto L42
L41:
	;
	goto L42
L42:
	;
	v109 = F_flatten_join_alias_vars_mutator(m, v95, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L30
	} else {
		goto L43
	}
L43:
	;
	v111 = F_lappend(m, v52, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v75+v47<<(uint(int32(2))%32))))
	v114 = F_copyObjectImpl(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L30
	} else {
		goto L45
	}
L45:
	;
	v116 = F_lappend(m, v50, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L30
	} else {
		goto L46
	}
L46:
	;
	v119 = v116
	v120 = v111
	goto L34
L47:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v135 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_IncrementVarSublevelsUp(m, v133, v135, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L30
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v139 == int32(6) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+44)) = v142
	goto L54
L53:
	;
	goto L54
L54:
	;
	v144 = F_flatten_join_alias_vars_mutator(m, v133, l1)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L30
	} else {
		goto L55
	}
L55:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v146 != int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v348 = v144
	goto L4
L57:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v149 != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v150 = F_checkExprHasSubLink(m, v144)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L30
	} else {
		goto L59
	}
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v150)
	goto L56
L60:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v156 != v157 {
		v343 = v154
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v160 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	if v161 == v160 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if int32(0) <= v218 {
		goto L73
	} else {
		goto L74
	}
L63:
	;
	v218 = base.I32_ctz(v204) | v205<<(uint(int32(5))%32)
	goto L62
L64:
	;
	v218 = int32(-2)
	goto L62
L65:
	;
	v171 = base.I32_div_s(int32(0), int32(32))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v172 <= v171 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v175 = v161 + int32(8)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v171<<(uint(int32(2))%32))))
	v182 = v179 & int32(-1)
	if v182 != 0 {
		v204 = v182
		v205 = v171
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v184 = v171 + int32(1)
	if v184 == v172 {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v187 = v184
	goto L69
L69:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175+v187<<(uint(int32(2))%32))))
	if v194 != 0 {
		v204 = v194
		v205 = v187
		goto L63
	} else {
		goto L71
	}
L70:
	;
	goto L64
L71:
	;
	v196 = v187 + int32(1)
	if v196 != v172 {
		v187 = v196
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v222 = v218
	v223 = v160
	goto L76
L74:
	;
	v309 = v160
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v309
	return v154
L76:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v159)+52))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v233 = int32(2)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v232+v222<<(uint(v233)%32)-int32(4))))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	if v239 == v233 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v309 = v248
	goto L75
L78:
	;
	if v161 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L79:
	;
	v242 = F_get_relids_for_join(m, v159, v222)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L30
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v246 = F_bms_add_member(m, v223, v222)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L30
	} else {
		goto L84
	}
L82:
	;
	v244 = F_bms_join(m, v223, v242)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L30
	} else {
		goto L83
	}
L83:
	;
	v248 = v244
	goto L78
L84:
	;
	v248 = v246
	goto L78
L85:
	;
	if int32(0) <= v304 {
		v222 = v304
		v223 = v248
		goto L76
	} else {
		goto L96
	}
L86:
	;
	v304 = base.I32_ctz(v290) | v291<<(uint(int32(5))%32)
	goto L85
L87:
	;
	v304 = int32(-2)
	goto L85
L88:
	;
	v255 = v222 + int32(1)
	v257 = base.I32_div_s(v255, int32(32))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v258 <= v257 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v261 = v161 + int32(8)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261+v257<<(uint(int32(2))%32))))
	v268 = v265 & (int32(-1) << (uint(v255) % 32))
	if v268 != 0 {
		v290 = v268
		v291 = v257
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v270 = v257 + int32(1)
	if v270 == v258 {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	v273 = v270
	goto L92
L92:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v261+v273<<(uint(int32(2))%32))))
	if v280 != 0 {
		v290 = v280
		v291 = v273
		goto L86
	} else {
		goto L94
	}
L93:
	;
	goto L87
L94:
	;
	v282 = v273 + int32(1)
	if v282 != v258 {
		v273 = v282
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	goto L77
L97:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+39)))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v332 = v330 | v331
	*(*uint8)(unsafe.Add(mBase, uint32(v328)+39)) = uint8(v332)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v323)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v335 - int32(1)
	return v328
L98:
	;
	v343 = v341
	goto L5
L99:
	;
	return v469
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L30
	} else {
		goto L149
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L30
	} else {
		goto L146
	}
L102:
	;
	m.G0 = v359 + int32(16)
	goto L99
L103:
	;
	v364 = int32(0)
	if v348 == v364 {
		v431 = v364
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v437 != 0 {
		goto L132
	} else {
		goto L133
	}
L105:
	;
	v437 = v431
	goto L104
L106:
	;
	v369 = v348
	goto L107
L107:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	switch v374 - int32(6) {
	case 0:
		goto L113
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 24, 25, 26, 27, 28, 29, 30, 31:
		v431 = v364
		goto L105
	case 9:
		goto L112
	case 21, 22, 23:
		goto L111
	case 32:
		goto L110
	default:
		goto L114
	}
L108:
	;
	v431 = v364
	goto L105
L109:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	if v426 != 0 {
		v369 = v426
		goto L107
	} else {
		goto L131
	}
L110:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	if v396 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L111:
	;
	v425 = v369 + int32(4)
	goto L109
L112:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v369)+16))
	if v387 != int32(2) {
		v431 = v364
		goto L105
	} else {
		goto L118
	}
L113:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v369)+28))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v383 != v384 {
		v431 = v364
		goto L105
	} else {
		goto L117
	}
L114:
	;
	if v374 != int32(319) {
		v431 = v364
		goto L105
	} else {
		goto L115
	}
L115:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v369)+20))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v379 != v380 {
		v431 = v364
		goto L105
	} else {
		goto L116
	}
L116:
	;
	v437 = int32(1)
	goto L104
L117:
	;
	v437 = int32(1)
	goto L104
L118:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v369)+28))
	if v390 == int32(0) {
		v431 = v364
		goto L105
	} else {
		goto L119
	}
L119:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v425 = v393
	goto L109
L120:
	;
	v437 = int32(1)
	goto L104
L121:
	;
	goto L122
L122:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v401 <= int32(0) {
		v431 = int32(1)
		goto L105
	} else {
		goto L123
	}
L123:
	;
	v404 = int32(0)
	if v404 < v401 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v408 = v401
	goto L126
L125:
	;
	v408 = v404
	goto L126
L126:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	v410 = v404
	goto L127
L127:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v409+v410<<(uint(int32(2))%32))))
	v419 = F_is_standard_join_alias_expression(m, v418, l0)
	mBase = m.M
	if v419 == int32(0) {
		v431 = v419
		goto L105
	} else {
		goto L129
	}
L128:
	;
	v431 = v419
	goto L105
L129:
	;
	v423 = v410 + int32(1)
	if v423 != v408 {
		v410 = v423
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	goto L108
L132:
	;
	F_adjust_standard_join_alias_expression(m, v348, l0)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L30
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	if v356 == int32(0) {
		goto L100
	} else {
		goto L136
	}
L135:
	;
	v469 = v348
	goto L102
L136:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v359)+12)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v359)+8)) = v356
	v445 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v359)+4)) = v445
	v451 = F_query_or_expression_tree_walker_impl(m, v348, int32(896), v359+int32(4), v445)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L30
	} else {
		goto L137
	}
L137:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v453 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v461 = v453
	goto L140
L139:
	;
	if v442 != 0 {
		goto L101
	} else {
		goto L141
	}
L140:
	;
	v462 = F_make_placeholder_expr(m, v356, v348, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L30
	} else {
		goto L144
	}
L141:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v456 = F_get_relids_for_join(m, v454, v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L30
	} else {
		goto L142
	}
L142:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v459 = F_bms_del_member(m, v456, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L30
	} else {
		goto L143
	}
L143:
	;
	v461 = v459
	goto L140
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462)+20)) = v442
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v466 = F_bms_copy(m, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L30
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462)+12)) = v466
	v469 = v462
	goto L102
L146:
	;
	F_errmsg_internal(m, int32(_a_F_flatten_join_alias_vars_mutator_0), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L30
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_flatten_join_alias_vars_mutator_1), int32(1200), int32(_a_F_flatten_join_alias_vars_mutator_2))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L30
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errmsg_internal(m, int32(_a_F_flatten_join_alias_vars_mutator_0), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L30
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_flatten_join_alias_vars_mutator_1), int32(1215), int32(_a_F_flatten_join_alias_vars_mutator_2))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L30
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_generate_join_implied_equalities(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
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
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v353 int32
	_ = v353
	var v366 int32
	_ = v366
	v6 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v14))|base.B2i32(int32(1)<<(uint(v14)%32)&int32(44) == v6) == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v27 = F_bms_union(m, l2, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v34 = l1
	v35 = v13
	goto L3
L3:
	;
	if l4 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return int32(0)
L5:
	;
	v32 = F_add_outer_joins_to_relids(m, l0, v27, l4, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v34 = v32
	v35 = v26
	goto L3
L7:
	;
	if v193 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L8:
	;
	v184 = F_get_common_eclass_indexes(m, l0, v35, l2)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L41
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	if v38 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v34 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if v97 <= int32(0) {
		v193 = v6
		goto L7
	} else {
		goto L22
	}
L12:
	;
	v97 = base.I32_ctz(v83) | v84<<(uint(int32(5))%32)
	goto L11
L13:
	;
	v97 = int32(-2)
	goto L11
L14:
	;
	v50 = base.I32_div_s(int32(0), int32(32))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v51 <= v50 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v54 = v34 + int32(8)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v50<<(uint(int32(2))%32))))
	v61 = v58 & int32(-1)
	if v61 != 0 {
		v83 = v61
		v84 = v50
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v63 = v50 + int32(1)
	if v63 == v51 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v66 = v63
	goto L18
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v54+v66<<(uint(int32(2))%32))))
	if v73 != 0 {
		v83 = v73
		v84 = v66
		goto L12
	} else {
		goto L20
	}
L19:
	;
	goto L13
L20:
	;
	v75 = v66 + int32(1)
	if v75 != v51 {
		v66 = v75
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v104 = v97
	v107 = v6
	goto L23
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v104 == v112 {
		v125 = v107
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v193 = v125
	goto L7
L25:
	;
	if v34 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v104<<(uint(int32(2))%32))))
	if v118 == int32(0) {
		v125 = v107
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+136))
	v122 = F_bms_add_members(m, v107, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v125 = v122
	goto L25
L29:
	;
	if int32(0) < v181 {
		v104 = v181
		v107 = v125
		goto L23
	} else {
		goto L40
	}
L30:
	;
	v181 = base.I32_ctz(v167) | v168<<(uint(int32(5))%32)
	goto L29
L31:
	;
	v181 = int32(-2)
	goto L29
L32:
	;
	v132 = v104 + int32(1)
	v134 = base.I32_div_s(v132, int32(32))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v135 <= v134 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v138 = v34 + int32(8)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138+v134<<(uint(int32(2))%32))))
	v145 = v142 & (int32(-1) << (uint(v132) % 32))
	if v145 != 0 {
		v167 = v145
		v168 = v134
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v147 = v134 + int32(1)
	if v147 == v135 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v150 = v147
	goto L36
L36:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v138+v150<<(uint(int32(2))%32))))
	if v157 != 0 {
		v167 = v157
		v168 = v150
		goto L30
	} else {
		goto L38
	}
L37:
	;
	goto L31
L38:
	;
	v159 = v150 + int32(1)
	if v159 != v135 {
		v150 = v159
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L24
L41:
	;
	v193 = v184
	goto L7
L42:
	;
	if int32(0) <= v254 {
		goto L53
	} else {
		goto L54
	}
L43:
	;
	v254 = base.I32_ctz(v240) | v241<<(uint(int32(5))%32)
	goto L42
L44:
	;
	v254 = int32(-2)
	goto L42
L45:
	;
	v207 = base.I32_div_s(int32(0), int32(32))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v208 <= v207 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v211 = v193 + int32(8)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+v207<<(uint(int32(2))%32))))
	v218 = v215 & int32(-1)
	if v218 != 0 {
		v240 = v218
		v241 = v207
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v220 = v207 + int32(1)
	if v220 == v208 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v223 = v220
	goto L49
L49:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v211+v223<<(uint(int32(2))%32))))
	if v230 != 0 {
		v240 = v230
		v241 = v223
		goto L43
	} else {
		goto L51
	}
L50:
	;
	goto L44
L51:
	;
	v232 = v223 + int32(1)
	if v232 != v208 {
		v223 = v232
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v261 = v254
	v267 = v6
	goto L56
L54:
	;
	v366 = v6
	goto L55
L55:
	;
	return v366
L56:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270+v261<<(uint(int32(2))%32))))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+40)))
	if v275 != 0 {
		v297 = v267
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v366 = v297
	goto L55
L58:
	;
	if v193 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L59:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	if v276 == int32(0) {
		v297 = v267
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v279 < int32(2) {
		v297 = v267
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+42)))
	if v282 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v294 = F_list_concat(m, v267, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L69
	}
L63:
	;
	v285 = F_generate_join_implied_equalities_normal(m, l0, v274, l1, l2, v13)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v291 = F_generate_join_implied_equalities_broken(m, l0, v274, v34, l2, v35, l3)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+42)))
	if v287 != int32(1) {
		v293 = v285
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v293 = v291
	goto L62
L69:
	;
	v297 = v294
	goto L58
L70:
	;
	if int32(0) <= v353 {
		v261 = v353
		v267 = v297
		goto L56
	} else {
		goto L81
	}
L71:
	;
	v353 = base.I32_ctz(v339) | v340<<(uint(int32(5))%32)
	goto L70
L72:
	;
	v353 = int32(-2)
	goto L70
L73:
	;
	v304 = v261 + int32(1)
	v306 = base.I32_div_s(v304, int32(32))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v307 <= v306 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v310 = v193 + int32(8)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v310+v306<<(uint(int32(2))%32))))
	v317 = v314 & (int32(-1) << (uint(v304) % 32))
	if v317 != 0 {
		v339 = v317
		v340 = v306
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v319 = v306 + int32(1)
	if v319 == v307 {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v322 = v319
	goto L77
L77:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v310+v322<<(uint(int32(2))%32))))
	if v329 != 0 {
		v339 = v329
		v340 = v322
		goto L71
	} else {
		goto L79
	}
L78:
	;
	goto L72
L79:
	;
	v331 = v322 + int32(1)
	if v331 != v307 {
		v322 = v331
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	goto L57
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
	var v78 int32
	_ = v78
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
	var v134 int32
	_ = v134
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
	var v188 int32
	_ = v188
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
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	v7 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l5)+224))
	v248 = F_adjust_appendrel_attrs_multilevel(m, l0, v211, l5, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L57
	} else {
		goto L65
	}
L2:
	;
	return v242
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v12 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v242 = int32(0)
	goto L2
L6:
	;
	v17 = int32(0)
	v22 = v7
	goto L9
L7:
	;
	v211 = v7
	goto L8
L8:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v215&int32(-2) != int32(2) {
		goto L60
	} else {
		goto L61
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v17<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	v32 = int32(0)
	if v31 == v32 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v211 = v200
	goto L8
L11:
	;
	v202 = v17 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v202 < v203 {
		v17 = v202
		v22 = v200
		goto L9
	} else {
		goto L59
	}
L12:
	;
	if v85 == int32(0) {
		v200 = v22
		goto L11
	} else {
		goto L26
	}
L13:
	;
	v85 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	if l2 == int32(0) {
		v78 = v32
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v85 = v78
	goto L12
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v42 < v41 {
		v78 = v32
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v44 = int32(1)
	if v41 <= v44 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v47 = v44
	goto L21
L20:
	;
	v47 = v41
	goto L21
L21:
	;
	v48 = int32(8)
	v53 = int32(0)
	goto L22
L22:
	;
	v60 = v53 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v31+v48+v60)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2+v48+v60)))
	v67 = v62 & (v64 ^ int32(-1))
	v69 = base.B2i32(v67 == int32(0))
	if v67 != 0 {
		v78 = v69
		goto L16
	} else {
		goto L24
	}
L23:
	;
	v78 = v69
	goto L16
L24:
	;
	v71 = v53 + int32(1)
	if v71 != v47 {
		v53 = v71
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v88 = int32(0)
	if v31 == v88 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v141 != 0 {
		v200 = v22
		goto L11
	} else {
		goto L41
	}
L28:
	;
	v141 = int32(1)
	goto L27
L29:
	;
	goto L30
L30:
	;
	if l3 == int32(0) {
		v134 = v88
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v141 = v134
	goto L27
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v98 < v97 {
		v134 = v88
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v100 = int32(1)
	if v97 <= v100 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v103 = v100
	goto L36
L35:
	;
	v103 = v97
	goto L36
L36:
	;
	v104 = int32(8)
	v109 = int32(0)
	goto L37
L37:
	;
	v116 = v109 << (uint(int32(2)) % 32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v31+v104+v116)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l3+v104+v116)))
	v123 = v118 & (v120 ^ int32(-1))
	v125 = base.B2i32(v123 == int32(0))
	if v123 != 0 {
		v134 = v125
		goto L31
	} else {
		goto L39
	}
L38:
	;
	v134 = v125
	goto L31
L39:
	;
	v127 = v109 + int32(1)
	if v127 != v103 {
		v109 = v127
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v142 = int32(0)
	if v31 == v142 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v195 != 0 {
		v200 = v22
		goto L11
	} else {
		goto L56
	}
L43:
	;
	v195 = int32(1)
	goto L42
L44:
	;
	goto L45
L45:
	;
	if l4 == int32(0) {
		v188 = v142
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v195 = v188
	goto L42
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v152 < v151 {
		v188 = v142
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v154 = int32(1)
	if v151 <= v154 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v157 = v154
	goto L51
L50:
	;
	v157 = v151
	goto L51
L51:
	;
	v158 = int32(8)
	v163 = int32(0)
	goto L52
L52:
	;
	v170 = v163 << (uint(int32(2)) % 32)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v31+v158+v170)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l4+v158+v170)))
	v177 = v172 & (v174 ^ int32(-1))
	v179 = base.B2i32(v177 == int32(0))
	if v177 != 0 {
		v188 = v179
		goto L46
	} else {
		goto L54
	}
L53:
	;
	v188 = v179
	goto L46
L54:
	;
	v181 = v163 + int32(1)
	if v181 != v157 {
		v163 = v181
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v196 = F_lappend(m, v22, v30)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	v200 = v196
	goto L11
L59:
	;
	goto L10
L60:
	;
	if base.B2i32(v211 == int32(0))|base.B2i32(v215 != int32(5)) != 0 {
		v242 = v211
		goto L2
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v211 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L1
L64:
	;
	goto L5
L65:
	;
	return v248
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
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
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
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	v6 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v7 != 0 {
		v389 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v389
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v8 != 0 {
		v389 = v6
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
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v147 == int32(0) {
		goto L38
	} else {
		goto L39
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
	v139 = v18 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v139 < v140 {
		v18 = v139
		goto L7
	} else {
		goto L37
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
		v73 = v27
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v80 = v73
	goto L10
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v37 < v36 {
		v73 = v27
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
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v26+v43+v55)))
	v62 = v57 & (v59 ^ int32(-1))
	v64 = base.B2i32(v62 == int32(0))
	if v62 != 0 {
		v73 = v64
		goto L14
	} else {
		goto L22
	}
L21:
	;
	v73 = v64
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
	if base.B2i32(v83 == v85)|base.B2i32(v84 == v85) != 0 {
		v131 = base.B2i32(v83|v84 == v85)
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v131 != 0 {
		goto L9
	} else {
		goto L36
	}
L26:
	;
	goto L25
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v99 != v100 {
		v131 = int32(0)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v102 = int32(1)
	if v99 <= v102 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v105 = v102
	goto L31
L30:
	;
	v105 = v99
	goto L31
L31:
	;
	v106 = int32(8)
	v111 = int32(0)
	goto L32
L32:
	;
	v119 = v111 << (uint(int32(2)) % 32)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v83+v106+v119)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v84+v106+v119)))
	v124 = base.B2i32(v121 == v123)
	if v121 != v123 {
		v131 = v124
		goto L26
	} else {
		goto L34
	}
L33:
	;
	v131 = v124
	goto L26
L34:
	;
	v127 = v111 + int32(1)
	if v127 != v105 {
		v111 = v127
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	return int32(1)
L37:
	;
	goto L8
L38:
	;
	v389 = int32(0)
	goto L1
L39:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v150 <= int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v157 = int32(0)
	goto L41
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	v160 = int32(2)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v157<<(uint(v160)%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+20))
	if v164 == v160 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L38
L43:
	;
	v378 = v157 + int32(1)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v378 < v379 {
		v157 = v378
		goto L41
	} else {
		goto L105
	}
L44:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v169 = int32(0)
	if v167 == v169 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v222 != 0 {
		goto L59
	} else {
		goto L60
	}
L46:
	;
	v222 = int32(1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	if v168 == int32(0) {
		v215 = v169
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v222 = v215
	goto L45
L50:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v179 < v178 {
		v215 = v169
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v181 = int32(1)
	if v178 <= v181 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v184 = v181
	goto L54
L53:
	;
	v184 = v178
	goto L54
L54:
	;
	v185 = int32(8)
	v190 = int32(0)
	goto L55
L55:
	;
	v197 = v190 << (uint(int32(2)) % 32)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v167+v185+v197)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v168+v185+v197)))
	v204 = v199 & (v201 ^ int32(-1))
	v206 = base.B2i32(v204 == int32(0))
	if v204 != 0 {
		v215 = v206
		goto L49
	} else {
		goto L57
	}
L56:
	;
	v215 = v206
	goto L49
L57:
	;
	v208 = v190 + int32(1)
	if v208 != v184 {
		v190 = v208
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v225 = int32(0)
	if v223 == v225 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	goto L61
L61:
	;
	v279 = int32(1)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v282 = int32(0)
	if base.B2i32(v280 == v282)|base.B2i32(v281 == v282) != 0 {
		v327 = v282
		goto L78
	} else {
		goto L79
	}
L62:
	;
	if v278 != 0 {
		goto L43
	} else {
		goto L76
	}
L63:
	;
	v278 = int32(1)
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v224 == int32(0) {
		v271 = v225
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v278 = v271
	goto L62
L67:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v235 < v234 {
		v271 = v225
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v237 = int32(1)
	if v234 <= v237 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v240 = v237
	goto L71
L70:
	;
	v240 = v234
	goto L71
L71:
	;
	v241 = int32(8)
	v246 = int32(0)
	goto L72
L72:
	;
	v253 = v246 << (uint(int32(2)) % 32)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v223+v241+v253)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v224+v241+v253)))
	v260 = v255 & (v257 ^ int32(-1))
	v262 = base.B2i32(v260 == int32(0))
	if v260 != 0 {
		v271 = v262
		goto L66
	} else {
		goto L74
	}
L73:
	;
	v271 = v262
	goto L66
L74:
	;
	v264 = v246 + int32(1)
	if v264 != v240 {
		v246 = v264
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	goto L61
L77:
	;
	if v327 != 0 {
		v389 = v279
		goto L1
	} else {
		goto L90
	}
L78:
	;
	goto L77
L79:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v292 < v293 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v295 = v292
	goto L82
L81:
	;
	v295 = v293
	goto L82
L82:
	;
	if v295 <= int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v298 = int32(1)
	goto L85
L84:
	;
	v298 = v295
	goto L85
L85:
	;
	v299 = int32(8)
	v304 = int32(0)
	goto L86
L86:
	;
	v311 = v304 << (uint(int32(2)) % 32)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v281+v299+v311)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v280+v299+v311)))
	v316 = v313 & v315
	v318 = base.B2i32(v316 != int32(0))
	if v316 != 0 {
		v327 = v318
		goto L78
	} else {
		goto L88
	}
L87:
	;
	v327 = v318
	goto L78
L88:
	;
	v320 = v304 + int32(1)
	if v320 != v298 {
		v304 = v320
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v330 = int32(0)
	if base.B2i32(v328 == v330)|base.B2i32(v329 == v330) != 0 {
		v375 = v330
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v375 != 0 {
		v389 = v279
		goto L1
	} else {
		goto L104
	}
L92:
	;
	goto L91
L93:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v340 < v341 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v343 = v340
	goto L96
L95:
	;
	v343 = v341
	goto L96
L96:
	;
	if v343 <= int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v346 = int32(1)
	goto L99
L98:
	;
	v346 = v343
	goto L99
L99:
	;
	v347 = int32(8)
	v352 = int32(0)
	goto L100
L100:
	;
	v359 = v352 << (uint(int32(2)) % 32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v329+v347+v359)))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v328+v347+v359)))
	v364 = v361 & v363
	v366 = base.B2i32(v364 != int32(0))
	if v364 != 0 {
		v375 = v366
		goto L92
	} else {
		goto L102
	}
L101:
	;
	v375 = v366
	goto L92
L102:
	;
	v368 = v352 + int32(1)
	if v368 != v346 {
		v352 = v368
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	goto L43
L105:
	;
	goto L42
}
