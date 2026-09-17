package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___newlocale(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v415 int64
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v609 int64
	_ = v609
	var v611 int64
	_ = v611
	var v613 int64
	_ = v613
	var v625 int32
	_ = v625
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v15 = int32(0)
	goto L3
L1:
	;
	m.G0 = v10 + int32(32)
	return v625
L2:
	;
	v625 = int32(0)
	goto L1
L3:
	;
	v20 = v15 << (uint(int32(2)) % 32)
	v26 = int32(1) << (uint(v15) % 32) & l0
	v27 = int32(0)
	if v26|base.B2i32(l2 == v27) == v27 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L64
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+(v10+int32(8))))) = v152
	if v152 == int32(-1) {
		goto L2
	} else {
		goto L62
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2+v20)))
	v152 = v33
	goto L5
L7:
	;
	goto L8
L8:
	;
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v35 = l1
	goto L11
L10:
	;
	v35 = int32(_a_F___newlocale_0)
	goto L11
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v39 != 0 {
		v53 = v35
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v152 = v151
	goto L5
L13:
	;
	v56 = int32(0)
	goto L28
L14:
	;
	v41 = F_getenv(m, int32(_a_F___newlocale_1))
	mBase = m.M
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v42 != 0 {
		v53 = v41
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v47 = F_getenv(m, v15*int32(12)+int32(_a_F___newlocale_2))
	mBase = m.M
	if v47 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v48 != 0 {
		v53 = v47
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v50 = F_getenv(m, int32(_a_F___newlocale_3))
	mBase = m.M
	if v50 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v51 != 0 {
		v53 = v50
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v53 = int32(_a_F___newlocale_4)
	goto L13
L26:
	;
	goto L25
L27:
	;
	v75 = int32(_a_F___newlocale_4)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v76 == int32(46) {
		v83 = v75
		goto L38
	} else {
		goto L39
	}
L28:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v56))))
	v61 = int32(0)
	if base.B2i32(v60 == v61)|base.B2i32(v60 == int32(47)) == v61 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v74 = v56
	goto L27
L30:
	;
	v68 = int32(23)
	v70 = v56 + int32(1)
	if v70 != v68 {
		v56 = v70
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	v74 = v68
	goto L27
L34:
	;
	v151 = v143
	goto L12
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F___newlocale[0]))
	if v104 != 0 {
		goto L49
	} else {
		goto L50
	}
L36:
	;
	if v15 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L37:
	;
	v89 = F_strcmp(m, v87, int32(_a_F___newlocale_4))
	mBase = m.M
	if v89 == int32(0) {
		v94 = v87
		goto L36
	} else {
		goto L43
	}
L38:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v84 == int32(0) {
		v94 = v83
		goto L36
	} else {
		goto L42
	}
L39:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v74))))
	if v80 != 0 {
		v83 = v75
		goto L38
	} else {
		goto L40
	}
L40:
	;
	if v76 != int32(67) {
		v87 = v53
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v83 = v53
	goto L38
L42:
	;
	v87 = v83
	goto L37
L43:
	;
	v93 = F_strcmp(m, v87, int32(_a_F___newlocale_5))
	mBase = m.M
	if v93 != 0 {
		goto L35
	} else {
		goto L44
	}
L44:
	;
	v94 = v87
	goto L36
L45:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v98 == int32(46) {
		v143 = int32(_a_F___newlocale_6)
		goto L34
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v151 = int32(0)
	goto L12
L48:
	;
	goto L47
L49:
	;
	v107 = v104
	goto L52
L50:
	;
	goto L51
L51:
	;
	v122 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v122 != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v112 = F_strcmp(m, v87, v107+int32(8))
	mBase = m.M
	if v112 == int32(0) {
		v143 = v107
		goto L34
	} else {
		goto L54
	}
L53:
	;
	goto L51
L54:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+32))
	if v115 != 0 {
		v107 = v115
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v124 = *(*int64)(unsafe.Add(mBase, _c_F___newlocale[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v122))) = v124
	v127 = v122 + int32(8)
	v128 = F___memcpy(m, v127, v87, v74)
	mBase = m.M
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v127+v74))) = uint8(v130)
	v132 = int32(_a_F___newlocale_7)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F___newlocale[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = v133
	*(*int32)(unsafe.Add(mBase, _c_F___newlocale[0])) = v122
	goto L58
L57:
	;
	goto L58
L58:
	;
	if v15|v122 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v140 = v122
	goto L61
L60:
	;
	v140 = int32(_a_F___newlocale_6)
	goto L61
L61:
	;
	v143 = v140
	goto L34
L62:
	;
	v157 = v15 + int32(1)
	if v157 != int32(6) {
		v15 = v157
		goto L3
	} else {
		goto L63
	}
L63:
	;
	goto L4
L64:
	;
	if base.B2i32(l2 != int32(0))&base.B2i32(l2 != int32(_a_F___newlocale_8))&base.B2i32(l2 != int32(_a_F___newlocale_9))&base.B2i32(l2 != int32(_a_F___newlocale_10))&base.B2i32(l2 != int32(_a_F___newlocale_11)) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v176 = int32(_a_F___newlocale_8)
	v178 = v10 + int32(8)
	v180 = int32(24)
	goto L71
L66:
	;
	v604 = l2
	goto L67
L67:
	;
	v609 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v604)+16)) = v609
	v611 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v604)+8)) = v611
	v613 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v604))) = v613
	v625 = v604
	goto L1
L68:
	;
	if v242 == int32(0) {
		v625 = v176
		goto L1
	} else {
		goto L86
	}
L69:
	;
	v242 = int32(0)
	goto L68
L70:
	;
	v216 = v211
	v217 = v212
	v218 = v213
	goto L80
L71:
	;
	if (v178|v176)&int32(3) != 0 {
		v211 = v178
		v212 = v176
		v213 = v180
		goto L70
	} else {
		goto L74
	}
L73:
	;
	if v201 == int32(0) {
		goto L69
	} else {
		goto L79
	}
L74:
	;
	v188 = v178
	v189 = v176
	v190 = v180
	goto L75
L75:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v193 != v194 {
		v211 = v188
		v212 = v189
		v213 = v190
		goto L70
	} else {
		goto L77
	}
L76:
	;
	goto L73
L77:
	;
	v196 = int32(4)
	v197 = v189 + v196
	v199 = v188 + v196
	v201 = v190 - v196
	if base.Ui32(int32(3)) < base.Ui32(v201) {
		v188 = v199
		v189 = v197
		v190 = v201
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v211 = v199
	v212 = v197
	v213 = v201
	goto L70
L80:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v221 == v222 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v242 = v221 - v222
	goto L68
L82:
	;
	v224 = int32(1)
	v229 = v218 - v224
	if v229 != 0 {
		v216 = v216 + v224
		v217 = v217 + v224
		v218 = v229
		goto L80
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	goto L69
L86:
	;
	v245 = int32(_a_F___newlocale_9)
	v247 = int32(24)
	goto L90
L87:
	;
	if v309 == int32(0) {
		v625 = v245
		goto L1
	} else {
		goto L105
	}
L88:
	;
	v309 = int32(0)
	goto L87
L89:
	;
	v283 = v278
	v284 = v279
	v285 = v280
	goto L99
L90:
	;
	if (v178|v245)&int32(3) != 0 {
		v278 = v178
		v279 = v245
		v280 = v247
		goto L89
	} else {
		goto L93
	}
L92:
	;
	if v268 == int32(0) {
		goto L88
	} else {
		goto L98
	}
L93:
	;
	v255 = v178
	v256 = v245
	v257 = v247
	goto L94
L94:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v260 != v261 {
		v278 = v255
		v279 = v256
		v280 = v257
		goto L89
	} else {
		goto L96
	}
L95:
	;
	goto L92
L96:
	;
	v263 = int32(4)
	v264 = v256 + v263
	v266 = v255 + v263
	v268 = v257 - v263
	if base.Ui32(int32(3)) < base.Ui32(v268) {
		v255 = v266
		v256 = v264
		v257 = v268
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v278 = v266
	v279 = v264
	v280 = v268
	goto L89
L99:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v288 == v289 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v309 = v288 - v289
	goto L87
L101:
	;
	v291 = int32(1)
	v296 = v285 - v291
	if v296 != 0 {
		v283 = v283 + v291
		v284 = v284 + v291
		v285 = v296
		goto L99
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	goto L100
L104:
	;
	goto L88
L105:
	;
	v312 = int32(0)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, _c_F___newlocale[2])))
	if v314 == v312 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v320 = v312
	goto L109
L107:
	;
	goto L108
L108:
	;
	v462 = int32(_a_F___newlocale_10)
	v464 = v10 + int32(8)
	v466 = int32(24)
	goto L165
L109:
	;
	v326 = int32(_a_F___newlocale_0)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F___newlocale[3])))
	if v330 != 0 {
		v344 = v326
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v449 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F___newlocale[2])) = uint8(v449)
	v453 = *(*int32)(unsafe.Add(mBase, _c_F___newlocale[4]))
	*(*int32)(unsafe.Add(mBase, _c_F___newlocale[5])) = v453
	goto L108
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320<<(uint(int32(2))%32))+uint32(_c_F___newlocale[4]))) = v442
	v445 = v320 + int32(1)
	if v445 != int32(6) {
		v320 = v445
		goto L109
	} else {
		goto L161
	}
L112:
	;
	v347 = int32(0)
	goto L127
L113:
	;
	v332 = F_getenv(m, int32(_a_F___newlocale_1))
	mBase = m.M
	if v332 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v333 != 0 {
		v344 = v332
		goto L112
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v338 = F_getenv(m, v320*int32(12)+int32(_a_F___newlocale_2))
	mBase = m.M
	if v338 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	if v339 != 0 {
		v344 = v338
		goto L112
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v341 = F_getenv(m, int32(_a_F___newlocale_3))
	mBase = m.M
	if v341 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v342 != 0 {
		v344 = v341
		goto L112
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v344 = int32(_a_F___newlocale_4)
	goto L112
L125:
	;
	goto L124
L126:
	;
	v366 = int32(_a_F___newlocale_4)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	if v367 == int32(46) {
		v374 = v366
		goto L137
	} else {
		goto L138
	}
L127:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v347))))
	v352 = int32(0)
	if base.B2i32(v351 == v352)|base.B2i32(v351 == int32(47)) == v352 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v365 = v347
	goto L126
L129:
	;
	v359 = int32(23)
	v361 = v347 + int32(1)
	if v361 != v359 {
		v347 = v361
		goto L127
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	goto L128
L132:
	;
	v365 = v359
	goto L126
L133:
	;
	v442 = v434
	goto L111
L134:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F___newlocale[0]))
	if v395 != 0 {
		goto L148
	} else {
		goto L149
	}
L135:
	;
	if v320 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L136:
	;
	v380 = F_strcmp(m, v378, int32(_a_F___newlocale_4))
	mBase = m.M
	if v380 == int32(0) {
		v385 = v378
		goto L135
	} else {
		goto L142
	}
L137:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)))
	if v375 == int32(0) {
		v385 = v374
		goto L135
	} else {
		goto L141
	}
L138:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v365))))
	if v371 != 0 {
		v374 = v366
		goto L137
	} else {
		goto L139
	}
L139:
	;
	if v367 != int32(67) {
		v378 = v344
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v374 = v344
	goto L137
L141:
	;
	v378 = v374
	goto L136
L142:
	;
	v384 = F_strcmp(m, v378, int32(_a_F___newlocale_5))
	mBase = m.M
	if v384 != 0 {
		goto L134
	} else {
		goto L143
	}
L143:
	;
	v385 = v378
	goto L135
L144:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+1)))
	if v389 == int32(46) {
		v434 = int32(_a_F___newlocale_6)
		goto L133
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v442 = int32(0)
	goto L111
L147:
	;
	goto L146
L148:
	;
	v398 = v395
	goto L151
L149:
	;
	goto L150
L150:
	;
	v413 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v413 != 0 {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v403 = F_strcmp(m, v378, v398+int32(8))
	mBase = m.M
	if v403 == int32(0) {
		v434 = v398
		goto L133
	} else {
		goto L153
	}
L152:
	;
	goto L150
L153:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v398)+32))
	if v406 != 0 {
		v398 = v406
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v415 = *(*int64)(unsafe.Add(mBase, _c_F___newlocale[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v413))) = v415
	v418 = v413 + int32(8)
	v419 = F___memcpy(m, v418, v378, v365)
	mBase = m.M
	v421 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v418+v365))) = uint8(v421)
	v423 = int32(_a_F___newlocale_7)
	v424 = *(*int32)(unsafe.Add(mBase, _c_F___newlocale[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+32)) = v424
	*(*int32)(unsafe.Add(mBase, _c_F___newlocale[0])) = v413
	goto L157
L156:
	;
	goto L157
L157:
	;
	if v320|v413 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v431 = v413
	goto L160
L159:
	;
	v431 = int32(_a_F___newlocale_6)
	goto L160
L160:
	;
	v434 = v431
	goto L133
L161:
	;
	goto L110
L162:
	;
	if v528 == int32(0) {
		v625 = v462
		goto L1
	} else {
		goto L180
	}
L163:
	;
	v528 = int32(0)
	goto L162
L164:
	;
	v502 = v497
	v503 = v498
	v504 = v499
	goto L174
L165:
	;
	if (v464|v462)&int32(3) != 0 {
		v497 = v464
		v498 = v462
		v499 = v466
		goto L164
	} else {
		goto L168
	}
L167:
	;
	if v487 == int32(0) {
		goto L163
	} else {
		goto L173
	}
L168:
	;
	v474 = v464
	v475 = v462
	v476 = v466
	goto L169
L169:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	if v479 != v480 {
		v497 = v474
		v498 = v475
		v499 = v476
		goto L164
	} else {
		goto L171
	}
L170:
	;
	goto L167
L171:
	;
	v482 = int32(4)
	v483 = v475 + v482
	v485 = v474 + v482
	v487 = v476 - v482
	if base.Ui32(int32(3)) < base.Ui32(v487) {
		v474 = v485
		v475 = v483
		v476 = v487
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v497 = v485
	v498 = v483
	v499 = v487
	goto L164
L174:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
	if v507 == v508 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v528 = v507 - v508
	goto L162
L176:
	;
	v510 = int32(1)
	v515 = v504 - v510
	if v515 != 0 {
		v502 = v502 + v510
		v503 = v503 + v510
		v504 = v515
		goto L174
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	goto L175
L179:
	;
	goto L163
L180:
	;
	v531 = int32(_a_F___newlocale_11)
	v533 = int32(24)
	goto L184
L181:
	;
	if v595 == int32(0) {
		v625 = v531
		goto L1
	} else {
		goto L199
	}
L182:
	;
	v595 = int32(0)
	goto L181
L183:
	;
	v569 = v564
	v570 = v565
	v571 = v566
	goto L193
L184:
	;
	if (v464|v531)&int32(3) != 0 {
		v564 = v464
		v565 = v531
		v566 = v533
		goto L183
	} else {
		goto L187
	}
L186:
	;
	if v554 == int32(0) {
		goto L182
	} else {
		goto L192
	}
L187:
	;
	v541 = v464
	v542 = v531
	v543 = v533
	goto L188
L188:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	if v546 != v547 {
		v564 = v541
		v565 = v542
		v566 = v543
		goto L183
	} else {
		goto L190
	}
L189:
	;
	goto L186
L190:
	;
	v549 = int32(4)
	v550 = v542 + v549
	v552 = v541 + v549
	v554 = v543 - v549
	if base.Ui32(int32(3)) < base.Ui32(v554) {
		v541 = v552
		v542 = v550
		v543 = v554
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v564 = v552
	v565 = v550
	v566 = v554
	goto L183
L193:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569))))
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570))))
	if v574 == v575 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v595 = v574 - v575
	goto L181
L195:
	;
	v577 = int32(1)
	v582 = v571 - v577
	if v582 != 0 {
		v569 = v569 + v577
		v570 = v570 + v577
		v571 = v582
		goto L193
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	goto L194
L198:
	;
	goto L182
L199:
	;
	v599 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v599 == int32(0) {
		goto L2
	} else {
		goto L200
	}
L200:
	;
	v604 = v599
	goto L67
}
func F_name_matches_visible_ENR(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v4 = int32(0)
	if v3 == v4 {
		v36 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.B2i32(v36 != int32(0))
L2:
	;
	goto L1
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v9 == int32(0) {
		v36 = v4
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v12 <= int32(0) {
		v36 = v4
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v17 = int32(0)
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15+v17<<(uint(int32(2))%32))))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = F_strcmp(m, v26, l1)
	mBase = m.M
	if v27 == int32(0) {
		v36 = v25
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v36 = int32(0)
	goto L2
L8:
	;
	v31 = v17 + int32(1)
	if v12 != v31 {
		v17 = v31
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_nameeqtext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = F_strlen(m, v8)
	mBase = m.M
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v15 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v45 != int32(950) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v21 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v32 = int32(1)
	if v15&v32 != 0 {
		v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v24 = int32(16)
	goto L9
L8:
	;
	v24 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = int32(4)
	goto L12
L11:
	;
	v31 = v24
	goto L12
L12:
	;
	v44 = v31
	goto L3
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v151 != v10 {
		goto L51
	} else {
		goto L52
	}
L15:
	;
	v140 = int32(1)
	if v15&v140 != 0 {
		goto L47
	} else {
		goto L48
	}
L16:
	;
	if v45 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v44 != v14 {
		v150 = int32(0)
		goto L14
	} else {
		goto L25
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_nameeqtext_0), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(_a_F_nameeqtext_1), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_nameeqtext_2), int32(1648), int32(_a_F_nameeqtext_3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
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
	v70 = int32(1)
	if v15&v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v74 = v70
	goto L28
L27:
	;
	v74 = int32(4)
	goto L28
L28:
	;
	v75 = v10 + v74
	if base.Ui32(int32(4)) <= base.Ui32(v14) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v150 = base.B2i32(v137 == int32(0))
	goto L14
L30:
	;
	v137 = int32(0)
	goto L29
L31:
	;
	v111 = v106
	v112 = v107
	v113 = v108
	goto L41
L32:
	;
	if (v8|v75)&int32(3) != 0 {
		v106 = v8
		v107 = v75
		v108 = v14
		goto L31
	} else {
		goto L35
	}
L33:
	;
	v99 = v8
	v100 = v75
	v101 = v14
	goto L34
L34:
	;
	if v101 == int32(0) {
		goto L30
	} else {
		goto L40
	}
L35:
	;
	v83 = v8
	v84 = v75
	v85 = v14
	goto L36
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v88 != v89 {
		v106 = v83
		v107 = v84
		v108 = v85
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v99 = v94
	v100 = v92
	v101 = v96
	goto L34
L38:
	;
	v91 = int32(4)
	v92 = v84 + v91
	v94 = v83 + v91
	v96 = v85 - v91
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v83 = v94
		v84 = v92
		v85 = v96
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v106 = v99
	v107 = v100
	v108 = v101
	goto L31
L41:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v116 == v117 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v137 = v116 - v117
	goto L29
L43:
	;
	v119 = int32(1)
	v124 = v113 - v119
	if v124 != 0 {
		v111 = v111 + v119
		v112 = v112 + v119
		v113 = v124
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	goto L30
L47:
	;
	v144 = v140
	goto L49
L48:
	;
	v144 = int32(4)
	goto L49
L49:
	;
	v146 = F_varstr_cmp(m, v8, v14, v10+v144, v44, v45)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v150 = base.B2i32(v146 == int32(0))
	goto L14
L51:
	;
	F_pfree(m, v10)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return v150
L54:
	;
	goto L53
}
func F_namegetext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(1542), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v6^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_namegt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.B2i32(int32(0) < v61)
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v55 = F_strlen(m, v5)
	mBase = m.M
	v56 = F_strlen(m, v4)
	mBase = m.M
	v57 = F_varstr_cmp(m, v5, v55, v4, v56, v6)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v61 = v46 - v47
	goto L1
L7:
	;
	goto L8
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L13
L10:
	;
	v42 = v4
	v46 = int32(0)
	goto L11
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L5
L12:
	;
	v42 = v37
	v46 = v39
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(v19 != v21)|base.B2i32(v21 == int32(0)) != 0 {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v37 = v31
	v39 = int32(0)
	goto L12
L15:
	;
	v27 = v18 - int32(1)
	if v27 == int32(0) {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v30 = int32(1)
	v31 = v17 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v32 != 0 {
		v16 = v16 + v30
		v17 = v31
		v18 = v27
		v19 = v32
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	return int32(0)
L19:
	;
	v61 = v57
	goto L1
}
func F_nameiclike(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v12 = F_DirectFunctionCall1Coll(m, int32(1436), int32(0), v4)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_pg_detoast_datum_packed(m, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v17 = F_Generic_Text_IC_like(m, v14, v6, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v17 == int32(1))
				}
			}
		}
	}
}
func F_nameicnlike(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v12 = F_DirectFunctionCall1Coll(m, int32(1436), int32(0), v4)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_pg_detoast_datum_packed(m, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v17 = F_Generic_Text_IC_like(m, v14, v6, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v17 != int32(1))
				}
			}
		}
	}
}
func F_nameicregexeq(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13947(m, l0, int32(27))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_namein(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_strlen(m, v4)
	mBase = m.M
	if int32(64) <= v5 {
		v9 = F_pg_mbcliplen(m, v4, v5, int32(63))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = v9
			v15 = F_palloc0(m, int32(64))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v13 != 0 {
					base.MemoryCopy(m, v15, v4, v13)
				} else {
				}
				return v15
			}
		}
	} else {
		v13 = v5
		v15 = F_palloc0(m, int32(64))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v13 != 0 {
				base.MemoryCopy(m, v15, v4, v13)
			} else {
			}
			return v15
		}
	}
}
func F_namelttext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(1542), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v6) >> (uint(int32(31)) % 32))
	}
}
func F_namene(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.B2i32(v61 != int32(0))
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v55 = F_strlen(m, v5)
	mBase = m.M
	v56 = F_strlen(m, v4)
	mBase = m.M
	v57 = F_varstr_cmp(m, v5, v55, v4, v56, v6)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v61 = v46 - v47
	goto L1
L7:
	;
	goto L8
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L13
L10:
	;
	v42 = v4
	v46 = int32(0)
	goto L11
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L5
L12:
	;
	v42 = v37
	v46 = v39
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(v19 != v21)|base.B2i32(v21 == int32(0)) != 0 {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v37 = v31
	v39 = int32(0)
	goto L12
L15:
	;
	v27 = v18 - int32(1)
	if v27 == int32(0) {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v30 = int32(1)
	v31 = v17 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v32 != 0 {
		v16 = v16 + v30
		v17 = v31
		v18 = v27
		v19 = v32
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	return int32(0)
L19:
	;
	v61 = v57
	goto L1
}
func F_namenlike(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v14 = F_strlen(m, v6)
		mBase = m.M
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v19 = v17 & v12
		if v19 != 0 {
			v20 = v13
		} else {
			v20 = v8 + int32(4)
		}
		if v17 == int32(1) {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v26 == int32(18) {
				v29 = int32(16)
			} else {
				v29 = int32(0)
			}
			if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v36 = int32(4)
			} else {
				v36 = v29
			}
			v47 = v36
		} else {
			v37 = int32(1)
			if v19 != 0 {
				v47 = int32(base.Ui32(v17)>>(uint(v37)%32)) - v37
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v49 = F_GenericMatchText(m, v6, v14, v20, v47, v48)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v49 != int32(1))
		}
	}
}
func F_namestrcmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	if l0|l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(-1)
	goto L3
L2:
	;
	v8 = int32(0)
	goto L3
L3:
	;
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v9 = int32(1)
	goto L6
L5:
	;
	v9 = v8
	goto L6
L6:
	;
	v10 = int32(0)
	if base.B2i32(l0 == v10)|base.B2i32(l1 == v10) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v61 = v9
	goto L9
L8:
	;
	goto L12
L9:
	;
	return v61
L10:
	;
	v61 = v52 - v53
	goto L9
L12:
	;
	goto L13
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v21 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v22 = l0
	v23 = l1
	v24 = int32(64)
	v25 = v21
	goto L18
L15:
	;
	v48 = l1
	v52 = int32(0)
	goto L16
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	goto L10
L17:
	;
	v48 = v43
	v52 = v45
	goto L16
L18:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if base.B2i32(v25 != v27)|base.B2i32(v27 == int32(0)) != 0 {
		v43 = v23
		v45 = v25
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v43 = v37
	v45 = int32(0)
	goto L17
L20:
	;
	v33 = v24 - int32(1)
	if v33 == int32(0) {
		v43 = v23
		v45 = v25
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v36 = int32(1)
	v37 = v23 + v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v38 != 0 {
		v22 = v22 + v36
		v23 = v37
		v24 = v33
		v25 = v38
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
}
func F_new_head_cell(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v7 <= v6 {
		v9 = int32(1)
		v11 = int32(16)
		v13 = v6 + v9
		if v13 <= v11 {
			v16 = v11
		} else {
			v16 = v13
		}
		if v16&(v16-int32(1)) != 0 {
			v23 = v9 << (uint(int32(32)-base.I32_clz(v16)) % 32)
		} else {
			v23 = v16
		}
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v26 = l0 + int32(16)
		if v24 == v26 {
			v28 = F_GetMemoryChunkContext(m, l0)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v32 = F_MemoryContextAlloc(m, v28, v23<<(uint(int32(2))%32))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v32
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v37 = v35 << (uint(int32(2)) % 32)
					if v37 == int32(0) {
					} else {
						base.MemoryCopy(m, v32, v26, v37)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v54 = v49
					v56 = v54 << (uint(int32(2)) % 32)
					if v56 != 0 {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						base.MemoryCopy(m, v57+int32(4), v57, v56)
					} else {
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62 + int32(1)
					return
				}
			}
		} else {
			v43 = F_repalloc(m, v24, v23<<(uint(int32(2))%32))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v54 = v49
				v56 = v54 << (uint(int32(2)) % 32)
				if v56 != 0 {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					base.MemoryCopy(m, v57+int32(4), v57, v56)
				} else {
				}
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62 + int32(1)
				return
			}
		}
	} else {
		v54 = v6
		v56 = v54 << (uint(int32(2)) % 32)
		if v56 != 0 {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			base.MemoryCopy(m, v57+int32(4), v57, v56)
		} else {
		}
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62 + int32(1)
		return
	}
}
