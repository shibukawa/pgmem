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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
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
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
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
	if v15 != int32(321) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v381 = m.G0
	v383 = v381 - int32(16)
	m.G0 = v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v385 == int32(0) {
		v493 = v372
		goto L114
	} else {
		goto L115
	}
L5:
	;
	return v367
L6:
	;
	v365 = F_expression_tree_mutator_impl(m, l0, int32(955), l1)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L34
	} else {
		goto L110
	}
L7:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v343 + int32(1)
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v348)
	v352 = F_query_tree_mutator_impl(m, l0, int32(955), l1, int32(4))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L34
	} else {
		goto L109
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
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v156) < base.Ui32(v155) {
		goto L64
	} else {
		goto L65
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
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return l0
L14:
	;
	goto L15
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = int32(2)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(v30)%32)-int32(4))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	if v36 != v30 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return l0
L17:
	;
	goto L18
L18:
	;
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v40 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v49 = int32(0)
	v52 = v3
	v54 = v3
	goto L22
L20:
	;
	goto L21
L21:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v128+v40<<(uint(int32(2))%32)-int32(4))))
	v135 = F_copyObjectImpl(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L34
	} else {
		goto L51
	}
L22:
	;
	v57 = int32(0)
	if v43 == v57 {
		v67 = v57
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v45 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v61 <= v49 {
		v67 = int32(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v67 = v63 + v49<<(uint(int32(2))%32)
	goto L24
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v96 != 0 {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	v82 = F_palloc0(m, int32(24))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v70 = int32(0)
	v78 = v70
	v79 = v70
	goto L28
L30:
	;
	goto L31
L31:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if base.B2i32(v67 == int32(0))|base.B2i32(v74 <= v49) != 0 {
		v78 = v52
		v79 = v54
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if v77 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v78 = v52
	v79 = v54
	goto L28
L34:
	;
	return int32(0)
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(36)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = v94
	v372 = v82
	goto L4
L36:
	;
	v97 = F_copyObjectImpl(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L34
	} else {
		goto L39
	}
L37:
	;
	v121 = v52
	v122 = v54
	goto L38
L38:
	;
	v49 = v49 + int32(1)
	v52 = v121
	v54 = v122
	goto L22
L39:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v99 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_IncrementVarSublevelsUp(m, v97, v99, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L34
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v103 == int32(6) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+44)) = v106
	goto L46
L45:
	;
	goto L46
L46:
	;
	v111 = F_flatten_join_alias_vars_mutator(m, v97, l1)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L34
	} else {
		goto L47
	}
L47:
	;
	v113 = F_lappend(m, v54, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L34
	} else {
		goto L48
	}
L48:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v77+v49<<(uint(int32(2))%32))))
	v116 = F_copyObjectImpl(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L34
	} else {
		goto L49
	}
L49:
	;
	v118 = F_lappend(m, v52, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L34
	} else {
		goto L50
	}
L50:
	;
	v121 = v118
	v122 = v113
	goto L38
L51:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v137 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_IncrementVarSublevelsUp(m, v135, v137, int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L34
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v141 == int32(6) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+44)) = v144
	goto L58
L57:
	;
	goto L58
L58:
	;
	v146 = F_flatten_join_alias_vars_mutator(m, v135, l1)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L34
	} else {
		goto L59
	}
L59:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v148 != int32(1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v372 = v146
	goto L4
L61:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v151 != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v152 = F_checkExprHasSubLink(m, v146)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L34
	} else {
		goto L63
	}
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v152)
	goto L60
L64:
	;
	return l0
L65:
	;
	goto L66
L66:
	;
	v160 = int32(0)
	if base.B2i32(v156 != v155)|base.B2i32(v156 <= v160) == v160 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v180 != v181 {
		v367 = v179
		goto L5
	} else {
		goto L73
	}
L68:
	;
	v166 = F_palloc0(m, int32(24))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L34
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v177 = F_expression_tree_mutator_impl(m, l0, int32(955), l1)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L34
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = int32(321)
	v170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v170
	v172 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v166)+16)) = v172
	v174 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v166))) = v174
	v179 = v166
	goto L67
L72:
	;
	v179 = v177
	goto L67
L73:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v184 = int32(0)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	if v185 == v184 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	if int32(0) <= v242 {
		goto L85
	} else {
		goto L86
	}
L75:
	;
	v242 = base.I32_ctz(v228) | v229<<(uint(int32(5))%32)
	goto L74
L76:
	;
	v242 = int32(-2)
	goto L74
L77:
	;
	v193 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v196 <= v193 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v199 = v185 + int32(8)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v206 = v203 & int32(-1)
	if v206 != 0 {
		v228 = v206
		v229 = v193
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v207 = int32(1)
	if v207 == v196 {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v211 = v207
	goto L81
L81:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v199+v211<<(uint(int32(2))%32))))
	if v218 != 0 {
		v228 = v218
		v229 = v211
		goto L75
	} else {
		goto L83
	}
L82:
	;
	goto L76
L83:
	;
	v220 = v211 + int32(1)
	if v220 != v196 {
		v211 = v220
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v245 = v242
	v246 = v184
	goto L88
L86:
	;
	v332 = v184
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v332
	return v179
L88:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v183)+52))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v257 = int32(2)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256+v245<<(uint(v257)%32)-int32(4))))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	if v263 == v257 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v332 = v272
	goto L87
L90:
	;
	if v185 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L91:
	;
	v266 = F_get_relids_for_join(m, v183, v245)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L34
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v270 = F_bms_add_member(m, v246, v245)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L34
	} else {
		goto L96
	}
L94:
	;
	v268 = F_bms_join(m, v246, v266)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L34
	} else {
		goto L95
	}
L95:
	;
	v272 = v268
	goto L90
L96:
	;
	v272 = v270
	goto L90
L97:
	;
	if int32(0) <= v328 {
		v245 = v328
		v246 = v272
		goto L88
	} else {
		goto L108
	}
L98:
	;
	v328 = base.I32_ctz(v314) | v315<<(uint(int32(5))%32)
	goto L97
L99:
	;
	v328 = int32(-2)
	goto L97
L100:
	;
	v279 = v245 + int32(1)
	v281 = int32(base.Ui32(v279) >> (uint(int32(5)) % 32))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v282 <= v281 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v285 = v185 + int32(8)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285+v281<<(uint(int32(2))%32))))
	v292 = v289 & (int32(-1) << (uint(v279) % 32))
	if v292 != 0 {
		v314 = v292
		v315 = v281
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v294 = v281 + int32(1)
	if v294 == v282 {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	v297 = v294
	goto L104
L104:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v285+v297<<(uint(int32(2))%32))))
	if v304 != 0 {
		v314 = v304
		v315 = v297
		goto L98
	} else {
		goto L106
	}
L105:
	;
	goto L99
L106:
	;
	v306 = v297 + int32(1)
	if v306 != v282 {
		v297 = v306
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	goto L89
L109:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+39)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v356 = v354 | v355
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+39)) = uint8(v356)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v347)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v359 - int32(1)
	return v352
L110:
	;
	v367 = v365
	goto L5
L111:
	;
	return v493
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L34
	} else {
		goto L161
	}
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L34
	} else {
		goto L158
	}
L114:
	;
	m.G0 = v383 + int32(16)
	goto L111
L115:
	;
	v388 = int32(0)
	if v372 == v388 {
		v455 = v388
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v461 != 0 {
		goto L144
	} else {
		goto L145
	}
L117:
	;
	v461 = v455
	goto L116
L118:
	;
	v393 = v372
	goto L119
L119:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	switch v398 - int32(6) {
	case 0:
		goto L125
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 24, 25, 26, 27, 28, 29, 30, 31:
		v455 = v388
		goto L117
	case 9:
		goto L124
	case 21, 22, 23:
		goto L123
	case 32:
		goto L122
	default:
		goto L126
	}
L120:
	;
	v455 = v388
	goto L117
L121:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	if v450 != 0 {
		v393 = v450
		goto L119
	} else {
		goto L143
	}
L122:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	if v420 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L123:
	;
	v449 = v393 + int32(4)
	goto L121
L124:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
	if v411 != int32(2) {
		v455 = v388
		goto L117
	} else {
		goto L130
	}
L125:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v393)+28))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v407 != v408 {
		v455 = v388
		goto L117
	} else {
		goto L129
	}
L126:
	;
	if v398 != int32(321) {
		v455 = v388
		goto L117
	} else {
		goto L127
	}
L127:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v393)+20))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v403 != v404 {
		v455 = v388
		goto L117
	} else {
		goto L128
	}
L128:
	;
	v461 = int32(1)
	goto L116
L129:
	;
	v461 = int32(1)
	goto L116
L130:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v393)+28))
	if v414 == int32(0) {
		v455 = v388
		goto L117
	} else {
		goto L131
	}
L131:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v414)+12))
	v449 = v417
	goto L121
L132:
	;
	v461 = int32(1)
	goto L116
L133:
	;
	goto L134
L134:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v425 <= int32(0) {
		v455 = int32(1)
		goto L117
	} else {
		goto L135
	}
L135:
	;
	v428 = int32(0)
	if v428 < v425 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v432 = v425
	goto L138
L137:
	;
	v432 = v428
	goto L138
L138:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	v434 = v428
	goto L139
L139:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v433+v434<<(uint(int32(2))%32))))
	v443 = F_is_standard_join_alias_expression(m, v442, l0)
	mBase = m.M
	if v443 == int32(0) {
		v455 = v443
		goto L117
	} else {
		goto L141
	}
L140:
	;
	v455 = v443
	goto L117
L141:
	;
	v447 = v434 + int32(1)
	if v447 != v432 {
		v434 = v447
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	goto L120
L144:
	;
	F_adjust_standard_join_alias_expression(m, v372, l0)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L34
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	if v380 == int32(0) {
		goto L112
	} else {
		goto L148
	}
L147:
	;
	v493 = v372
	goto L114
L148:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+12)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v383)+8)) = v380
	v469 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v383)+4)) = v469
	v475 = F_query_or_expression_tree_walker_impl(m, v372, int32(947), v383+int32(4), v469)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L34
	} else {
		goto L149
	}
L149:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v477 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v485 = v477
	goto L152
L151:
	;
	if v466 != 0 {
		goto L113
	} else {
		goto L153
	}
L152:
	;
	v486 = F_make_placeholder_expr(m, v380, v372, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L34
	} else {
		goto L156
	}
L153:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v480 = F_get_relids_for_join(m, v478, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L34
	} else {
		goto L154
	}
L154:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v483 = F_bms_del_member(m, v480, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L34
	} else {
		goto L155
	}
L155:
	;
	v485 = v483
	goto L152
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v486)+20)) = v466
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v490 = F_bms_copy(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L34
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v486)+12)) = v490
	v493 = v486
	goto L114
L158:
	;
	F_errmsg_internal(m, int32(_a_F_flatten_join_alias_vars_mutator_0), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L34
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_flatten_join_alias_vars_mutator_1), int32(1249), int32(_a_F_flatten_join_alias_vars_mutator_2))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L34
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errmsg_internal(m, int32(_a_F_flatten_join_alias_vars_mutator_0), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L34
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_flatten_join_alias_vars_mutator_1), int32(1264), int32(_a_F_flatten_join_alias_vars_mutator_2))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L34
	} else {
		goto L163
	}
L163:
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+252))
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
	v48 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v51 <= v48 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v54 = v34 + int32(8)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v61 = v58 & int32(-1)
	if v61 != 0 {
		v83 = v61
		v84 = v48
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v62 = int32(1)
	if v62 == v51 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v66 = v62
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
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
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
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v104<<(uint(int32(2))%32))))
	if v118 == int32(0) {
		v125 = v107
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+144))
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
	v134 = int32(base.Ui32(v132) >> (uint(int32(5)) % 32))
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
	v205 = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v208 <= v205 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v211 = v193 + int32(8)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v218 = v215 & int32(-1)
	if v218 != 0 {
		v240 = v218
		v241 = v205
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v219 = int32(1)
	if v219 == v208 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v223 = v219
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
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
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
	v306 = int32(base.Ui32(v304) >> (uint(int32(5)) % 32))
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
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l5)+248))
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v8 != 0 {
		v389 = v6
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
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
