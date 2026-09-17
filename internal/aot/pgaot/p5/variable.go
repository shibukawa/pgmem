package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_variable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	v18 = m.G0
	v20 = v18 - int32(144)
	m.G0 = v20
	v22 = l0
	v34 = int32(0)
	goto L12
L1:
	;
	m.G0 = v20 + int32(144)
	return v651
L2:
	;
	if v411 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L3:
	;
	v596 = int32(0)
	if base.B2i32(v228 == v596)|base.B2i32(v584 == v596) != 0 {
		goto L2
	} else {
		goto L149
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L51
	} else {
		goto L146
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L51
	} else {
		goto L143
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L51
	} else {
		goto L140
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L51
	} else {
		goto L137
	}
L8:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+35)))
	if v414 != int32(1) {
		v584 = v413
		goto L3
	} else {
		goto L112
	}
L9:
	;
	v397 = base.B2i32(v395 == int32(0))
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
	if v399|v397 != 0 {
		v409 = v395
		v411 = v397
		v413 = int32(1)
		goto L8
	} else {
		goto L111
	}
L10:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	if v380 < v377 {
		goto L7
	} else {
		goto L107
	}
L11:
	;
	v377 = v343
	v379 = v22 + int32(32)
	goto L10
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v360 = v192 & int32(_a_F_get_variable_0)
	if v360 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v42 = v40
	goto L16
L15:
	;
	v42 = int32(0)
	goto L16
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v44 = v43 + v34
	if v42 <= v44 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v51 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v58 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v58 = v56
	v59 = int32(8)
	goto L18
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	if v54 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v58 = v51
	v59 = int32(40)
	goto L18
L22:
	;
	goto L13
L23:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v334 != 0 {
		goto L22
	} else {
		goto L97
	}
L24:
	;
	v238 = base.I32_extend16_s(v192)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	if v240 != 0 {
		goto L65
	} else {
		goto L66
	}
L25:
	;
	v233 = int32(0)
	F_resolve_special_varno(m, v22, l2, int32(1486), v233)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L51
	} else {
		goto L63
	}
L26:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v62 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v65 < v58 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22+v59))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v70 == int32(0) {
		v192 = v69
		v196 = v58
		v199 = v62
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v206 = int32(1)
	v209 = (v196 - v206) << (uint(int32(2)) % 32)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209+v210)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v214+v209)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	switch v217 - v206 {
	case 0:
		goto L62
	case 1:
		goto L61
	default:
		goto L60
	}
L30:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v73 == int32(0) {
		v192 = v69
		v196 = v58
		v199 = v62
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73+v58<<(uint(int32(2))%32))))
	if v79 == int32(0) {
		v192 = v69
		v196 = v58
		v199 = v62
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v85 = v83 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v82+v85-int32(4))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v90 != 0 {
		v192 = v69
		v196 = v58
		v199 = v62
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if int32(0) < v69 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+24))
	if v93 < v69 {
		v192 = v69
		v196 = v58
		v199 = v62
		goto L29
	} else {
		goto L37
	}
L35:
	;
	v104 = v69
	goto L36
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v85+v73)))
	if v106 == int32(0) {
		v170 = v104
		v171 = v83
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95+v69<<(uint(int32(1))%32)-int32(2)))))
	if v101 == int32(0) {
		v192 = v69
		v196 = v58
		v199 = v62
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v104 = v101
	goto L36
L39:
	;
	v182 = F_bms_is_member(m, v171, v70)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v82+v109<<(uint(int32(2))%32)-int32(4))))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	if v116 != 0 {
		v170 = v104
		v171 = v83
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v122 = v104
	v125 = v109
	v128 = v106
	goto L42
L42:
	;
	v134 = base.I32_extend16_s(v122)
	if int32(0) < v134 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v170 = v148
	v171 = v125
	goto L39
L44:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v128)+24))
	if v137 < v134 {
		v192 = v69
		v196 = v58
		v199 = v62
		goto L29
	} else {
		goto L47
	}
L45:
	;
	v148 = v122
	goto L46
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v73+v125<<(uint(int32(2))%32))))
	if v152 == int32(0) {
		v170 = v148
		v171 = v125
		goto L39
	} else {
		goto L49
	}
L47:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v128)+28))
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139+v134<<(uint(int32(1))%32)-int32(2)))))
	if v145 == int32(0) {
		v192 = v69
		v196 = v58
		v199 = v62
		goto L29
	} else {
		goto L48
	}
L48:
	;
	v148 = v145
	goto L46
L49:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v82+v155<<(uint(int32(2))%32)-int32(4))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	if v162 == int32(0) {
		v122 = v148
		v125 = v155
		v128 = v152
		goto L42
	} else {
		goto L50
	}
L50:
	;
	goto L43
L51:
	;
	return int32(0)
L52:
	;
	if v182 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v186 = v171
	goto L55
L54:
	;
	v186 = v58
	goto L55
L55:
	;
	if v182 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v187 = v170
	goto L58
L57:
	;
	v187 = v69
	goto L58
L58:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v192 = v187
	v196 = v186
	v199 = v188
	goto L29
L59:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v212)+12))
	switch v229 - int32(1) {
	case 0, 5:
		goto L24
	case 1:
		goto L23
	default:
		goto L22
	}
L60:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v227 = v225 + v209
	goto L59
L61:
	;
	v227 = v50 + int32(28)
	goto L59
L62:
	;
	v227 = v50 + int32(24)
	goto L59
L63:
	;
	v651 = v233
	goto L1
L64:
	;
	if v229 != int32(2) {
		goto L22
	} else {
		goto L96
	}
L65:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v243 = v241
	goto L67
L66:
	;
	v243 = int32(0)
	goto L67
L67:
	;
	if v238 <= v243 {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	if v245 == int32(0) {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v50)+60))
	if v248 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	if v286 == int32(0) {
		goto L5
	} else {
		goto L83
	}
L71:
	;
	goto L70
L72:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v252 <= int32(0) {
		v286 = int32(0)
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v286 = int32(0)
	goto L71
L75:
	;
	v255 = int32(0)
	if v255 < v252 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v258 = v252
	goto L78
L77:
	;
	v258 = v255
	goto L78
L78:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v263 = int32(0)
	goto L79
L79:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v259+v263<<(uint(int32(2))%32))))
	v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v271)+8)))
	if v272 == v238&int32(_a_F_get_variable_0) {
		v286 = v271
		goto L71
	} else {
		goto L81
	}
L80:
	;
	goto L74
L81:
	;
	v275 = v263 + int32(1)
	if v275 != v258 {
		v263 = v275
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	base.MemoryCopy(m, v20-int32(-64), v50, int32(80))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
	v297 = F_lcons(m, v295, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L51
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+44)) = v297
	F_set_deparse_plan(m, v50, v290)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L51
	} else {
		goto L85
	}
L85:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if v303 != int32(6) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_appendStringInfoChar(m, v67, int32(40))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L51
	} else {
		goto L89
	}
L87:
	;
	v310 = v302
	goto L88
L88:
	;
	F_get_rule_expr(m, v310, l2, int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L51
	} else {
		goto L90
	}
L89:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	v310 = v309
	goto L88
L90:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	if v315 != int32(6) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	F_appendStringInfoChar(m, v67, int32(41))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L51
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
	v322 = F_list_delete_first(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L51
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	base.MemoryCopy(m, v50, v20-int32(-64), int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+44)) = v322
	v651 = int32(0)
	goto L1
L96:
	;
	goto L23
L97:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v212)+52))
	if v335 == int32(0) {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	if base.I32_extend16_s(v192) <= int32(0) {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	v343 = v192 & int32(_a_F_get_variable_0)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v341+v343<<(uint(int32(2))%32)-int32(4))))
	if v349 == int32(0) {
		goto L11
	} else {
		goto L100
	}
L100:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	if v352 != int32(6) {
		goto L11
	} else {
		goto L101
	}
L101:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v22 = v349
	v34 = v355 + v34
	goto L12
L102:
	;
	v364 = int32(1)
	v409 = int32(0)
	v411 = v364
	v413 = v364
	goto L8
L103:
	;
	goto L104
L104:
	;
	v367 = v22 + int32(32)
	v368 = base.I32_extend16_s(v192)
	if int32(0) < v368 {
		v377 = v360
		v379 = v367
		goto L10
	} else {
		goto L105
	}
L105:
	;
	v371 = F_get_rte_attribute_name(m, v212, v368)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L51
	} else {
		goto L106
	}
L106:
	;
	v394 = v367
	v395 = v371
	goto L9
L107:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v382+v377<<(uint(int32(2))%32)-int32(4))))
	if v388 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v390 = v388
	goto L110
L109:
	;
	v390 = int32(_a_F_get_variable_1)
	goto L110
L110:
	;
	v394 = v379
	v395 = v390
	goto L9
L111:
	;
	v403 = int32(0)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v409 = v395
	v411 = v403
	v413 = base.B2i32(v404 != v403)
	goto L8
L112:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+34)))
	if (v417|v413)&int32(1) != 0 {
		v584 = v413
		goto L3
	} else {
		goto L113
	}
L113:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v421 == int32(0) {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	if v424 <= int32(0) {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	v427 = int32(0)
	v432 = v427
	v434 = v427
	goto L116
L116:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v421)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446+v434<<(uint(int32(2))%32))))
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+26)))
	if v451 != 0 {
		v505 = v432
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L2
L118:
	;
	v510 = v434 + int32(1)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	if v510 < v511 {
		v432 = v505
		v434 = v510
		goto L116
	} else {
		goto L136
	}
L119:
	;
	v453 = v432 + int32(1)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v454 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v505 = v453
	goto L118
L121:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if base.B2i32(v475 == int32(0))|base.B2i32(v475 != v478) != 0 {
		v496 = v475
		v497 = v478
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v450)+12))
	if v468 == int32(0) {
		goto L120
	} else {
		goto L125
	}
L123:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	if v457 < v453 {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v471 = v454 + v457<<(uint(int32(4))%32) + v432*int32(100) + int32(24)
	goto L121
L125:
	;
	v471 = v468
	goto L121
L126:
	;
	if v496-v497 != 0 {
		goto L120
	} else {
		goto L133
	}
L127:
	;
	goto L126
L128:
	;
	v481 = v471
	v482 = v409
	goto L129
L129:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+1)))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+1)))
	if v486 == int32(0) {
		v496 = v486
		v497 = v485
		goto L127
	} else {
		goto L131
	}
L130:
	;
	v496 = v486
	v497 = v485
	goto L127
L131:
	;
	v489 = int32(1)
	if v486 == v485 {
		v481 = v481 + v489
		v482 = v482 + v489
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	v500 = F_equal(m, v22, v499)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L51
	} else {
		goto L134
	}
L134:
	;
	if v500 != 0 {
		v505 = v453
		goto L118
	} else {
		goto L135
	}
L135:
	;
	v584 = int32(1)
	goto L3
L136:
	;
	goto L117
L137:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v377
	F_errmsg_internal(m, int32(_a_F_get_variable_2), v20+int32(16))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L51
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_get_variable_3), int32(_a_F_get_variable_4), int32(_a_F_get_variable_5))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L51
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errmsg_internal(m, int32(_a_F_get_variable_6), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L51
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_get_variable_3), int32(_a_F_get_variable_7), int32(_a_F_get_variable_5))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L51
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v238
	F_errmsg_internal(m, int32(_a_F_get_variable_2), v20+int32(32))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L51
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_get_variable_3), int32(_a_F_get_variable_8), int32(_a_F_get_variable_5))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L51
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v566
	F_errmsg_internal(m, int32(_a_F_get_variable_9), v20+int32(48))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L51
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_get_variable_3), int32(_a_F_get_variable_10), int32(_a_F_get_variable_5))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L51
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
	v601 = F_quote_identifier(m, v228)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L51
	} else {
		goto L150
	}
L150:
	;
	F_appendStringInfoString(m, v67, v601)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L51
	} else {
		goto L151
	}
L151:
	;
	F_appendStringInfoChar(m, v67, int32(46))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L51
	} else {
		goto L152
	}
L152:
	;
	goto L2
L153:
	;
	v627 = F_quote_identifier(m, v409)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L51
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	F_appendStringInfoChar(m, v67, int32(42))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L51
	} else {
		goto L158
	}
L156:
	;
	F_appendStringInfoString(m, v67, v627)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L51
	} else {
		goto L157
	}
L157:
	;
	v651 = v409
	goto L1
L158:
	;
	if l1 == int32(0) {
		v651 = v409
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v638 = F_format_type_with_typemod(m, v636, v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L51
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v638
	F_appendStringInfo(m, v67, int32(_a_F_get_variable_11), v20)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L51
	} else {
		goto L161
	}
L161:
	;
	v651 = v409
	goto L1
}
func F_variable_paramref_hook(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(int32(-268435456)) < base.Ui32(v12-int32(268435456)) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		if v19 < v12 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v22 != 0 {
				v23 = int32(2)
				v27 = F_repalloc0(m, v22, v19<<(uint(v23)%32), v12<<(uint(v23)%32))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v35 = v27
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v12
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v48 = v43 + v12<<(uint(int32(2))%32) - int32(4)
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
					if v49 != 0 {
						if v49 != int32(2278) {
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							if v52 != int32(41) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(705)
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(705)
					}
					v58 = F_palloc0(m, int32(28))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v12
						*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(8)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
						*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v63
						v67 = F_get_typcollation(m, v63)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v67
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v70
							m.G0 = v10 + int32(16)
							return v58
						}
					}
				}
			} else {
				v33 = F_palloc0(m, v12<<(uint(int32(2))%32))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = v33
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v12
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v48 = v43 + v12<<(uint(int32(2))%32) - int32(4)
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
					if v49 != 0 {
						if v49 != int32(2278) {
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							if v52 != int32(41) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(705)
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(705)
					}
					v58 = F_palloc0(m, int32(28))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v12
						*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(8)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
						*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v63
						v67 = F_get_typcollation(m, v63)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v67
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v70
							m.G0 = v10 + int32(16)
							return v58
						}
					}
				}
			}
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
			v48 = v43 + v12<<(uint(int32(2))%32) - int32(4)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
			if v49 != 0 {
				if v49 != int32(2278) {
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					if v52 != int32(41) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(705)
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(705)
			}
			v58 = F_palloc0(m, int32(28))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v12
				*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(8)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
				*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v63
				v67 = F_get_typcollation(m, v63)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v67
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v70
					m.G0 = v10 + int32(16)
					return v58
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685636))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
				F_errmsg(m, int32(_a_F_variable_paramref_hook_0), v10)
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					F_parser_errposition(m, l0, v87)
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_variable_paramref_hook_1), int32(144), int32(_a_F_variable_paramref_hook_2))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
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
