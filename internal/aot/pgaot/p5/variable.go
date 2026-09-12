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
	var v32 int32
	_ = v32
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
	var v63 int32
	_ = v63
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
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
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
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
	var v438 int32
	_ = v438
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
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	v18 = m.G0
	v20 = v18 - int32(144)
	m.G0 = v20
	v22 = l0
	v32 = int32(0)
	goto L12
L1:
	;
	m.G0 = v20 + int32(144)
	return v646
L2:
	;
	if v410 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L3:
	;
	if v224 == int32(0) {
		goto L2
	} else {
		goto L161
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L53
	} else {
		goto L158
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L53
	} else {
		goto L155
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L53
	} else {
		goto L152
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L53
	} else {
		goto L149
	}
L8:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+35)))
	if v414 != int32(1) {
		v581 = v407
		goto L3
	} else {
		goto L123
	}
L9:
	;
	v398 = base.B2i32(v396 == int32(0))
	v399 = int32(1)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
	if v400 != 0 {
		v407 = v399
		v408 = v396
		v410 = v398
		goto L8
	} else {
		goto L121
	}
L10:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v380 < v375 {
		goto L7
	} else {
		goto L117
	}
L11:
	;
	v375 = v341
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
	v359 = v191 & int32(65535)
	if v359 == int32(0) {
		goto L112
	} else {
		goto L113
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
	v44 = v43 + v32
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
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v332 != 0 {
		goto L22
	} else {
		goto L107
	}
L24:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	if v235 != 0 {
		goto L67
	} else {
		goto L68
	}
L25:
	;
	v229 = int32(0)
	F_resolve_special_varno(m, v22, l2, int32(1501), v229)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L53
	} else {
		goto L65
	}
L26:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v62 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v65 = v63
	goto L29
L28:
	;
	v65 = int32(0)
	goto L29
L29:
	;
	if v65 < v58 {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22+v59))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v70 == int32(0) {
		v189 = v58
		v190 = v62
		v191 = v69
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v203 = int32(1)
	v206 = (v189 - v203) << (uint(int32(2)) % 32)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+v206)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210+v206)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	switch v213 - v203 {
	case 0:
		goto L64
	case 1:
		goto L63
	default:
		goto L62
	}
L32:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v73 == int32(0) {
		v189 = v58
		v190 = v62
		v191 = v69
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73+v58<<(uint(int32(2))%32))))
	if v79 == int32(0) {
		v189 = v58
		v190 = v62
		v191 = v69
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v84 = v82 - int32(4)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v87 = v85 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84+v87)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v90 != 0 {
		v189 = v58
		v190 = v62
		v191 = v69
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if int32(0) < v69 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+24))
	if v93 < v69 {
		v189 = v58
		v190 = v62
		v191 = v69
		goto L31
	} else {
		goto L39
	}
L37:
	;
	v104 = v69
	goto L38
L38:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v73+v87)))
	if v106 == int32(0) {
		v173 = v85
		v177 = v104
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95+v69<<(uint(int32(1))%32)-int32(2)))))
	if v101 == int32(0) {
		v189 = v58
		v190 = v62
		v191 = v69
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v104 = v101
	goto L38
L41:
	;
	v178 = F_bms_is_member(m, v173, v70)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L53
	} else {
		goto L54
	}
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v84+v109<<(uint(int32(2))%32))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	if v114 != 0 {
		v173 = v85
		v177 = v104
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v122 = v106
	v128 = v109
	v131 = v104
	goto L44
L44:
	;
	v132 = base.I32_extend16_s(v131)
	if int32(0) < v132 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v173 = v128
	v177 = v146
	goto L41
L46:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	if v135 < v132 {
		v189 = v58
		v190 = v62
		v191 = v69
		goto L31
	} else {
		goto L49
	}
L47:
	;
	v146 = v131
	goto L48
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v73+v128<<(uint(int32(2))%32))))
	if v150 == int32(0) {
		v173 = v128
		v177 = v146
		goto L41
	} else {
		goto L51
	}
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v122)+28))
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137+v132<<(uint(int32(1))%32)-int32(2)))))
	if v143 == int32(0) {
		v189 = v58
		v190 = v62
		v191 = v69
		goto L31
	} else {
		goto L50
	}
L50:
	;
	v146 = v143
	goto L48
L51:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v84+v153<<(uint(int32(2))%32))))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	if v158 == int32(0) {
		v122 = v150
		v128 = v153
		v131 = v146
		goto L44
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	return int32(0)
L54:
	;
	if v178 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v182 = v173
	goto L57
L56:
	;
	v182 = v58
	goto L57
L57:
	;
	if v178 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v183 = v177
	goto L60
L59:
	;
	v183 = v69
	goto L60
L60:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v189 = v182
	v190 = v184
	v191 = v183
	goto L31
L61:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	switch v225 - int32(1) {
	case 0, 5:
		goto L24
	case 1:
		goto L23
	default:
		goto L22
	}
L62:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v223 = v221 + v206
	goto L61
L63:
	;
	v223 = v50 + int32(28)
	goto L61
L64:
	;
	v223 = v50 + int32(24)
	goto L61
L65:
	;
	v646 = v229
	goto L1
L66:
	;
	if v225 != int32(2) {
		goto L22
	} else {
		goto L106
	}
L67:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v238 = v236
	goto L69
L68:
	;
	v238 = int32(0)
	goto L69
L69:
	;
	v239 = base.I32_extend16_s(v191)
	if v239 <= v238 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	if v241 == int32(0) {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v50)+60))
	if v244 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	if v282 == int32(0) {
		goto L5
	} else {
		goto L85
	}
L73:
	;
	goto L72
L74:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v248 <= int32(0) {
		v282 = int32(0)
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v282 = int32(0)
	goto L73
L77:
	;
	v251 = int32(0)
	if v251 < v248 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v254 = v248
	goto L80
L79:
	;
	v254 = v251
	goto L80
L80:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v259 = int32(0)
	goto L81
L81:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v255+v259<<(uint(int32(2))%32))))
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v267)+8)))
	if v268 == v239&int32(65535) {
		v282 = v267
		goto L73
	} else {
		goto L83
	}
L82:
	;
	goto L76
L83:
	;
	v271 = v259 + int32(1)
	if v271 != v254 {
		v259 = v271
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	goto L87
L86:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
	v294 = F_lcons(m, v292, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L53
	} else {
		goto L90
	}
L87:
	;
	v290 = F__emscripten_memcpy_bulkmem(m, v20-int32(-64), v50, int32(80))
	mBase = m.M
	goto L89
L89:
	;
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+44)) = v294
	F_set_deparse_plan(m, v50, v286)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L53
	} else {
		goto L91
	}
L91:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	if v300 != int32(6) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_appendStringInfoChar(m, v67, int32(40))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L53
	} else {
		goto L95
	}
L93:
	;
	v307 = v299
	goto L94
L94:
	;
	F_get_rule_expr(m, v307, l2, int32(1))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L53
	} else {
		goto L96
	}
L95:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v307 = v306
	goto L94
L96:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	if v312 != int32(6) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_appendStringInfoChar(m, v67, int32(41))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L53
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
	v319 = F_list_delete_first(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L53
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	goto L103
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324)+44)) = v319
	v646 = int32(0)
	goto L1
L103:
	;
	v324 = F__emscripten_memcpy_bulkmem(m, v50, v20-int32(-64), int32(80))
	mBase = m.M
	goto L105
L105:
	;
	goto L102
L106:
	;
	goto L23
L107:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v208)+52))
	if v333 == int32(0) {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	if base.I32_extend16_s(v191) <= int32(0) {
		goto L22
	} else {
		goto L109
	}
L109:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v341 = v191 & int32(65535)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v339+v341<<(uint(int32(2))%32)-int32(4))))
	if v347 == int32(0) {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	if v350 != int32(6) {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v22 = v347
	v32 = v353 + v32
	goto L12
L112:
	;
	v363 = int32(1)
	v407 = v363
	v408 = int32(0)
	v410 = v363
	goto L8
L113:
	;
	goto L114
L114:
	;
	v366 = v22 + int32(32)
	v367 = base.I32_extend16_s(v191)
	if int32(0) < v367 {
		v375 = v359
		v379 = v366
		goto L10
	} else {
		goto L115
	}
L115:
	;
	v370 = F_get_rte_attribute_name(m, v208, v367)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L53
	} else {
		goto L116
	}
L116:
	;
	v394 = v366
	v396 = v370
	goto L9
L117:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v382+v375<<(uint(int32(2))%32)-int32(4))))
	if v388 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v390 = v388
	goto L120
L119:
	;
	v390 = int32(533686)
	goto L120
L120:
	;
	v394 = v379
	v396 = v390
	goto L9
L121:
	;
	if v396 == int32(0) {
		v407 = v399
		v408 = v396
		v410 = v398
		goto L8
	} else {
		goto L122
	}
L122:
	;
	v403 = int32(0)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v407 = base.B2i32(v404 != v403)
	v408 = v396
	v410 = v403
	goto L8
L123:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+34)))
	if (v417|v407)&int32(1) != 0 {
		v581 = v407
		goto L3
	} else {
		goto L124
	}
L124:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v421 == int32(0) {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	if v424 <= int32(0) {
		goto L2
	} else {
		goto L126
	}
L126:
	;
	v427 = int32(0)
	v432 = v427
	v438 = v427
	goto L127
L127:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v421)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446+v432<<(uint(int32(2))%32))))
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+26)))
	if v451 != 0 {
		v505 = v438
		goto L129
	} else {
		goto L130
	}
L128:
	;
	goto L2
L129:
	;
	v509 = v432 + int32(1)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	if v509 < v510 {
		v432 = v509
		v438 = v505
		goto L127
	} else {
		goto L148
	}
L130:
	;
	v453 = v438 + int32(1)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v454 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v505 = v453
	goto L129
L132:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	if v476 == int32(0) {
		v495 = v475
		v496 = v476
		goto L138
	} else {
		goto L139
	}
L133:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v450)+12))
	if v468 == int32(0) {
		goto L131
	} else {
		goto L136
	}
L134:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	if v457 < v453 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v471 = v454 + v457<<(uint(int32(4))%32) + v438*int32(100) + int32(24)
	goto L132
L136:
	;
	v471 = v468
	goto L132
L137:
	;
	if v496-v495 != 0 {
		goto L131
	} else {
		goto L145
	}
L138:
	;
	goto L137
L139:
	;
	if v475 != v476 {
		v495 = v475
		v496 = v476
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v480 = v471
	v481 = v408
	goto L141
L141:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+1)))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+1)))
	if v485 == int32(0) {
		v495 = v484
		v496 = v485
		goto L138
	} else {
		goto L143
	}
L142:
	;
	v495 = v484
	v496 = v485
	goto L138
L143:
	;
	v488 = int32(1)
	if v484 == v485 {
		v480 = v480 + v488
		v481 = v481 + v488
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	v499 = F_equal(m, v22, v498)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L53
	} else {
		goto L146
	}
L146:
	;
	if v499 != 0 {
		v505 = v453
		goto L129
	} else {
		goto L147
	}
L147:
	;
	v581 = int32(1)
	goto L3
L148:
	;
	goto L128
L149:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v375
	F_errmsg_internal(m, int32(676230), v20+int32(16))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L53
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(482839), int32(7803), int32(387224))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L53
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errmsg_internal(m, int32(401145), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L53
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(482839), int32(7774), int32(387224))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L53
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v548
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v239
	F_errmsg_internal(m, int32(676230), v20+int32(32))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L53
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(482839), int32(7740), int32(387224))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L53
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v565
	F_errmsg_internal(m, int32(458301), v20+int32(48))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L53
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(482839), int32(7624), int32(387224))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L53
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
	if v581 == int32(0) {
		goto L2
	} else {
		goto L162
	}
L162:
	;
	v599 = F_quote_identifier(m, v224)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L53
	} else {
		goto L163
	}
L163:
	;
	F_appendStringInfoString(m, v67, v599)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L53
	} else {
		goto L164
	}
L164:
	;
	F_appendStringInfoChar(m, v67, int32(46))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L53
	} else {
		goto L165
	}
L165:
	;
	goto L2
L166:
	;
	v625 = F_quote_identifier(m, v408)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L53
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	F_appendStringInfoChar(m, v67, int32(42))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L53
	} else {
		goto L171
	}
L169:
	;
	F_appendStringInfoString(m, v67, v625)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L53
	} else {
		goto L170
	}
L170:
	;
	v646 = v408
	goto L1
L171:
	;
	if l1 == int32(0) {
		v646 = v408
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v636 = F_format_type_with_typemod(m, v634, v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L53
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v636
	F_appendStringInfo(m, v67, int32(173174), v20)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L53
	} else {
		goto L174
	}
L174:
	;
	v646 = v408
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
				F_errmsg(m, int32(456626), v10)
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
						F_errfinish(m, int32(486061), int32(144), int32(308234))
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
