package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___newlocale(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v409 int64
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
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
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v603 int64
	_ = v603
	var v605 int64
	_ = v605
	var v607 int64
	_ = v607
	var v618 int32
	_ = v618
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v14 = int32(0)
	goto L3
L1:
	;
	m.G0 = v9 + int32(32)
	return v618
L2:
	;
	v618 = int32(0)
	goto L1
L3:
	;
	v19 = int32(1) << (uint(v14) % 32) & l0
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L64
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(8)+v14<<(uint(int32(2))%32)))) = v146
	if v146 == int32(-1) {
		goto L2
	} else {
		goto L62
	}
L6:
	;
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v19 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2+v14<<(uint(int32(2))%32))))
	v146 = v30
	goto L5
L9:
	;
	v32 = l1
	goto L11
L10:
	;
	v32 = int32(717063)
	goto L11
L11:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v36 != 0 {
		v50 = v32
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v146 = v145
	goto L5
L13:
	;
	v53 = int32(0)
	goto L28
L14:
	;
	v38 = F_getenv(m, int32(511138))
	mBase = m.M
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39 != 0 {
		v50 = v38
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v44 = F_getenv(m, v14*int32(12)+int32(4038704))
	mBase = m.M
	if v44 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v45 != 0 {
		v50 = v44
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v47 = F_getenv(m, int32(514257))
	mBase = m.M
	if v47 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v48 != 0 {
		v50 = v47
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v50 = int32(526902)
	goto L13
L26:
	;
	goto L25
L27:
	;
	v69 = int32(526902)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v70 == int32(46) {
		v77 = v69
		goto L38
	} else {
		goto L39
	}
L28:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v53))))
	if v57 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v68 = v53
	goto L27
L30:
	;
	goto L29
L31:
	;
	if v57 == int32(47) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v62 = int32(23)
	v64 = v53 + int32(1)
	if v64 != v62 {
		v53 = v64
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v68 = v62
	goto L27
L34:
	;
	v145 = v137
	goto L12
L35:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[1110]))
	if v98 != 0 {
		goto L49
	} else {
		goto L50
	}
L36:
	;
	if v14 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L37:
	;
	v83 = F_strcmp(m, v81, int32(526902))
	mBase = m.M
	if v83 == int32(0) {
		v88 = v81
		goto L36
	} else {
		goto L43
	}
L38:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v78 == int32(0) {
		v88 = v77
		goto L36
	} else {
		goto L42
	}
L39:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v68))))
	if v74 != 0 {
		v77 = v69
		goto L38
	} else {
		goto L40
	}
L40:
	;
	if v70 != int32(67) {
		v81 = v50
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v77 = v50
	goto L38
L42:
	;
	v81 = v77
	goto L37
L43:
	;
	v87 = F_strcmp(m, v81, int32(487784))
	mBase = m.M
	if v87 != 0 {
		goto L35
	} else {
		goto L44
	}
L44:
	;
	v88 = v81
	goto L36
L45:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v92 == int32(46) {
		v137 = int32(4038612)
		goto L34
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v145 = int32(0)
	goto L12
L48:
	;
	goto L47
L49:
	;
	v101 = v98
	goto L52
L50:
	;
	goto L51
L51:
	;
	v116 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v116 != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v106 = F_strcmp(m, v81, v101+int32(8))
	mBase = m.M
	if v106 == int32(0) {
		v137 = v101
		goto L34
	} else {
		goto L54
	}
L53:
	;
	goto L51
L54:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+32))
	if v109 != 0 {
		v101 = v109
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v118 = *(*int64)(unsafe.Add(mBase, _consts[1111]))
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = v118
	v121 = v116 + int32(8)
	v122 = F___memcpy(m, v121, v81, v68)
	mBase = m.M
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v121+v68))) = uint8(v124)
	v126 = int32(4608080)
	v127 = *(*int32)(unsafe.Add(mBase, _consts[1110]))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = v127
	*(*int32)(unsafe.Add(mBase, _consts[1110])) = v116
	goto L58
L57:
	;
	goto L58
L58:
	;
	if v14|v116 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v134 = v116
	goto L61
L60:
	;
	v134 = int32(4038612)
	goto L61
L61:
	;
	v137 = v134
	goto L34
L62:
	;
	v151 = v14 + int32(1)
	if v151 != int32(6) {
		v14 = v151
		goto L3
	} else {
		goto L63
	}
L63:
	;
	goto L4
L64:
	;
	if base.B2i32(l2 != int32(0))&base.B2i32(l2 != int32(4038648))&base.B2i32(l2 != int32(4038672))&base.B2i32(l2 != int32(4608084))&base.B2i32(l2 != int32(4608108)) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v170 = int32(4038648)
	v172 = v9 + int32(8)
	v174 = int32(24)
	goto L71
L66:
	;
	v599 = l2
	goto L67
L67:
	;
	v603 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v599))) = v603
	v605 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v599)+16)) = v605
	v607 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v599)+8)) = v607
	v618 = v599
	goto L1
L68:
	;
	if v236 == int32(0) {
		v618 = v170
		goto L1
	} else {
		goto L86
	}
L69:
	;
	v236 = int32(0)
	goto L68
L70:
	;
	v210 = v205
	v211 = v206
	v212 = v207
	goto L80
L71:
	;
	if (v172|v170)&int32(3) != 0 {
		v205 = v172
		v206 = v170
		v207 = v174
		goto L70
	} else {
		goto L74
	}
L73:
	;
	if v195 == int32(0) {
		goto L69
	} else {
		goto L79
	}
L74:
	;
	v182 = v172
	v183 = v170
	v184 = v174
	goto L75
L75:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v187 != v188 {
		v205 = v182
		v206 = v183
		v207 = v184
		goto L70
	} else {
		goto L77
	}
L76:
	;
	goto L73
L77:
	;
	v190 = int32(4)
	v191 = v183 + v190
	v193 = v182 + v190
	v195 = v184 - v190
	if base.Ui32(int32(3)) < base.Ui32(v195) {
		v182 = v193
		v183 = v191
		v184 = v195
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v205 = v193
	v206 = v191
	v207 = v195
	goto L70
L80:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v215 == v216 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v236 = v215 - v216
	goto L68
L82:
	;
	v218 = int32(1)
	v223 = v212 - v218
	if v223 != 0 {
		v210 = v210 + v218
		v211 = v211 + v218
		v212 = v223
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
	v239 = int32(4038672)
	v241 = v9 + int32(8)
	v243 = int32(24)
	goto L90
L87:
	;
	if v305 == int32(0) {
		v618 = v239
		goto L1
	} else {
		goto L105
	}
L88:
	;
	v305 = int32(0)
	goto L87
L89:
	;
	v279 = v274
	v280 = v275
	v281 = v276
	goto L99
L90:
	;
	if (v241|v239)&int32(3) != 0 {
		v274 = v241
		v275 = v239
		v276 = v243
		goto L89
	} else {
		goto L93
	}
L92:
	;
	if v264 == int32(0) {
		goto L88
	} else {
		goto L98
	}
L93:
	;
	v251 = v241
	v252 = v239
	v253 = v243
	goto L94
L94:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	if v256 != v257 {
		v274 = v251
		v275 = v252
		v276 = v253
		goto L89
	} else {
		goto L96
	}
L95:
	;
	goto L92
L96:
	;
	v259 = int32(4)
	v260 = v252 + v259
	v262 = v251 + v259
	v264 = v253 - v259
	if base.Ui32(int32(3)) < base.Ui32(v264) {
		v251 = v262
		v252 = v260
		v253 = v264
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v274 = v262
	v275 = v260
	v276 = v264
	goto L89
L99:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if v284 == v285 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v305 = v284 - v285
	goto L87
L101:
	;
	v287 = int32(1)
	v292 = v281 - v287
	if v292 != 0 {
		v279 = v279 + v287
		v280 = v280 + v287
		v281 = v292
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
	v308 = int32(0)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1112])))
	if v310 == v308 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v316 = v308
	goto L109
L107:
	;
	goto L108
L108:
	;
	v455 = int32(4608084)
	v457 = v9 + int32(8)
	v459 = int32(24)
	goto L165
L109:
	;
	v323 = int32(717063)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1113])))
	if v327 != 0 {
		v341 = v323
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v443 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1112])) = uint8(v443)
	v447 = *(*int32)(unsafe.Add(mBase, _consts[1114]))
	*(*int32)(unsafe.Add(mBase, _consts[1115])) = v447
	goto L108
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316<<(uint(int32(2))%32))+uint32(_consts[1114]))) = v436
	v439 = v316 + int32(1)
	if v439 != int32(6) {
		v316 = v439
		goto L109
	} else {
		goto L161
	}
L112:
	;
	v344 = int32(0)
	goto L127
L113:
	;
	v329 = F_getenv(m, int32(511138))
	mBase = m.M
	if v329 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	if v330 != 0 {
		v341 = v329
		goto L112
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v335 = F_getenv(m, v316*int32(12)+int32(4038704))
	mBase = m.M
	if v335 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	if v336 != 0 {
		v341 = v335
		goto L112
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v338 = F_getenv(m, int32(514257))
	mBase = m.M
	if v338 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	if v339 != 0 {
		v341 = v338
		goto L112
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v341 = int32(526902)
	goto L112
L125:
	;
	goto L124
L126:
	;
	v360 = int32(526902)
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v361 == int32(46) {
		v368 = v360
		goto L137
	} else {
		goto L138
	}
L127:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341+v344))))
	if v348 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v359 = v344
	goto L126
L129:
	;
	goto L128
L130:
	;
	if v348 == int32(47) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v353 = int32(23)
	v355 = v344 + int32(1)
	if v355 != v353 {
		v344 = v355
		goto L127
	} else {
		goto L132
	}
L132:
	;
	v359 = v353
	goto L126
L133:
	;
	v436 = v428
	goto L111
L134:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _consts[1110]))
	if v389 != 0 {
		goto L148
	} else {
		goto L149
	}
L135:
	;
	if v316 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L136:
	;
	v374 = F_strcmp(m, v372, int32(526902))
	mBase = m.M
	if v374 == int32(0) {
		v379 = v372
		goto L135
	} else {
		goto L142
	}
L137:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+1)))
	if v369 == int32(0) {
		v379 = v368
		goto L135
	} else {
		goto L141
	}
L138:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341+v359))))
	if v365 != 0 {
		v368 = v360
		goto L137
	} else {
		goto L139
	}
L139:
	;
	if v361 != int32(67) {
		v372 = v341
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v368 = v341
	goto L137
L141:
	;
	v372 = v368
	goto L136
L142:
	;
	v378 = F_strcmp(m, v372, int32(487784))
	mBase = m.M
	if v378 != 0 {
		goto L134
	} else {
		goto L143
	}
L143:
	;
	v379 = v372
	goto L135
L144:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+1)))
	if v383 == int32(46) {
		v428 = int32(4038612)
		goto L133
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v436 = int32(0)
	goto L111
L147:
	;
	goto L146
L148:
	;
	v392 = v389
	goto L151
L149:
	;
	goto L150
L150:
	;
	v407 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v407 != 0 {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v397 = F_strcmp(m, v372, v392+int32(8))
	mBase = m.M
	if v397 == int32(0) {
		v428 = v392
		goto L133
	} else {
		goto L153
	}
L152:
	;
	goto L150
L153:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v392)+32))
	if v400 != 0 {
		v392 = v400
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v409 = *(*int64)(unsafe.Add(mBase, _consts[1111]))
	*(*int64)(unsafe.Add(mBase, uint32(v407))) = v409
	v412 = v407 + int32(8)
	v413 = F___memcpy(m, v412, v372, v359)
	mBase = m.M
	v415 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v412+v359))) = uint8(v415)
	v417 = int32(4608080)
	v418 = *(*int32)(unsafe.Add(mBase, _consts[1110]))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+32)) = v418
	*(*int32)(unsafe.Add(mBase, _consts[1110])) = v407
	goto L157
L156:
	;
	goto L157
L157:
	;
	if v316|v407 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v425 = v407
	goto L160
L159:
	;
	v425 = int32(4038612)
	goto L160
L160:
	;
	v428 = v425
	goto L133
L161:
	;
	goto L110
L162:
	;
	if v521 == int32(0) {
		v618 = v455
		goto L1
	} else {
		goto L180
	}
L163:
	;
	v521 = int32(0)
	goto L162
L164:
	;
	v495 = v490
	v496 = v491
	v497 = v492
	goto L174
L165:
	;
	if (v457|v455)&int32(3) != 0 {
		v490 = v457
		v491 = v455
		v492 = v459
		goto L164
	} else {
		goto L168
	}
L167:
	;
	if v480 == int32(0) {
		goto L163
	} else {
		goto L173
	}
L168:
	;
	v467 = v457
	v468 = v455
	v469 = v459
	goto L169
L169:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	if v472 != v473 {
		v490 = v467
		v491 = v468
		v492 = v469
		goto L164
	} else {
		goto L171
	}
L170:
	;
	goto L167
L171:
	;
	v475 = int32(4)
	v476 = v468 + v475
	v478 = v467 + v475
	v480 = v469 - v475
	if base.Ui32(int32(3)) < base.Ui32(v480) {
		v467 = v478
		v468 = v476
		v469 = v480
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v490 = v478
	v491 = v476
	v492 = v480
	goto L164
L174:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	if v500 == v501 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v521 = v500 - v501
	goto L162
L176:
	;
	v503 = int32(1)
	v508 = v497 - v503
	if v508 != 0 {
		v495 = v495 + v503
		v496 = v496 + v503
		v497 = v508
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
	v524 = int32(4608108)
	v526 = v9 + int32(8)
	v528 = int32(24)
	goto L184
L181:
	;
	if v590 == int32(0) {
		v618 = v524
		goto L1
	} else {
		goto L199
	}
L182:
	;
	v590 = int32(0)
	goto L181
L183:
	;
	v564 = v559
	v565 = v560
	v566 = v561
	goto L193
L184:
	;
	if (v526|v524)&int32(3) != 0 {
		v559 = v526
		v560 = v524
		v561 = v528
		goto L183
	} else {
		goto L187
	}
L186:
	;
	if v549 == int32(0) {
		goto L182
	} else {
		goto L192
	}
L187:
	;
	v536 = v526
	v537 = v524
	v538 = v528
	goto L188
L188:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	if v541 != v542 {
		v559 = v536
		v560 = v537
		v561 = v538
		goto L183
	} else {
		goto L190
	}
L189:
	;
	goto L186
L190:
	;
	v544 = int32(4)
	v545 = v537 + v544
	v547 = v536 + v544
	v549 = v538 - v544
	if base.Ui32(int32(3)) < base.Ui32(v549) {
		v536 = v547
		v537 = v545
		v538 = v549
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v559 = v547
	v560 = v545
	v561 = v549
	goto L183
L193:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	if v569 == v570 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v590 = v569 - v570
	goto L181
L195:
	;
	v572 = int32(1)
	v577 = v566 - v572
	if v577 != 0 {
		v564 = v564 + v572
		v565 = v565 + v572
		v566 = v577
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
	v594 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v594 == int32(0) {
		goto L2
	} else {
		goto L200
	}
L200:
	;
	v599 = v594
	goto L67
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
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
	if v8&int32(3) == int32(0) {
		v37 = v8
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v71 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v70 = v62 - v8
	goto L3
L5:
	;
	v41 = v37
	goto L14
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v70 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v26 = v8
	goto L10
L10:
	;
	v30 = v26 + int32(1)
	if v30&int32(3) == int32(0) {
		v37 = v30
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v62 = v30
	goto L4
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 != 0 {
		v26 = v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v50 = int32(-2139062144)
	if (int32(16843008)-v47|v47)&v50 == v50 {
		v41 = v41 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v56 = v41
	goto L17
L16:
	;
	goto L15
L17:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 != 0 {
		v56 = v56 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v62 = v56
	goto L4
L19:
	;
	goto L18
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v102 != int32(950) {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v74 = int32(4)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v76&int32(254) == int32(2) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v89 = int32(1)
	if v71&v89 != 0 {
		v101 = int32(base.Ui32(v71)>>(uint(v89)%32)) - v89
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v85 = v74
	goto L26
L25:
	;
	v85 = base.B2i32(v76 == int32(18)) << (uint(v74) % 32)
	goto L26
L26:
	;
	if v76 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = v74
	goto L29
L28:
	;
	v88 = v85
	goto L29
L29:
	;
	v101 = v88
	goto L20
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v101 = int32(base.Ui32(v95)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L31:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v208 != v10 {
		goto L68
	} else {
		goto L69
	}
L32:
	;
	v197 = int32(1)
	if v71&v197 != 0 {
		goto L64
	} else {
		goto L65
	}
L33:
	;
	if v102 != 0 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v70 != v101 {
		v207 = int32(0)
		goto L31
	} else {
		goto L42
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(233937), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errhint(m, int32(536375), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(478374), int32(1648), int32(99187))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v127 = int32(1)
	if v71&v127 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v131 = v127
	goto L45
L44:
	;
	v131 = int32(4)
	goto L45
L45:
	;
	v132 = v10 + v131
	if base.Ui32(int32(4)) <= base.Ui32(v70) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v207 = base.B2i32(v194 == int32(0))
	goto L31
L47:
	;
	v194 = int32(0)
	goto L46
L48:
	;
	v168 = v163
	v169 = v164
	v170 = v165
	goto L58
L49:
	;
	if (v8|v132)&int32(3) != 0 {
		v163 = v8
		v164 = v132
		v165 = v70
		goto L48
	} else {
		goto L52
	}
L50:
	;
	v156 = v8
	v157 = v132
	v158 = v70
	goto L51
L51:
	;
	if v158 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L52:
	;
	v140 = v8
	v141 = v132
	v142 = v70
	goto L53
L53:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v145 != v146 {
		v163 = v140
		v164 = v141
		v165 = v142
		goto L48
	} else {
		goto L55
	}
L54:
	;
	v156 = v151
	v157 = v149
	v158 = v153
	goto L51
L55:
	;
	v148 = int32(4)
	v149 = v141 + v148
	v151 = v140 + v148
	v153 = v142 - v148
	if base.Ui32(int32(3)) < base.Ui32(v153) {
		v140 = v151
		v141 = v149
		v142 = v153
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v163 = v156
	v164 = v157
	v165 = v158
	goto L48
L58:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v173 == v174 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v194 = v173 - v174
	goto L46
L60:
	;
	v176 = int32(1)
	v181 = v170 - v176
	if v181 != 0 {
		v168 = v168 + v176
		v169 = v169 + v176
		v170 = v181
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	goto L47
L64:
	;
	v201 = v197
	goto L66
L65:
	;
	v201 = int32(4)
	goto L66
L66:
	;
	v203 = F_varstr_cmp(m, v8, v70, v10+v201, v101, v102)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v207 = base.B2i32(v203 == int32(0))
	goto L31
L68:
	;
	F_pfree(m, v10)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	return v207
L71:
	;
	goto L70
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1557), v3, v4, v5)
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
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
	return base.B2i32(int32(0) < v172)
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	if v5&int32(3) == int32(0) {
		v77 = v5
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v172 = v45 - v46
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
	v41 = v4
	v45 = int32(0)
	goto L11
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	goto L5
L12:
	;
	v41 = v36
	v45 = v38
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 != v21 {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v36 = v30
	v38 = int32(0)
	goto L12
L15:
	;
	if v21 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v26 = v18 - int32(1)
	if v26 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v29 = int32(1)
	v30 = v17 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v31 != 0 {
		v16 = v16 + v29
		v17 = v30
		v18 = v26
		v19 = v31
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v4&int32(3) == int32(0) {
		v134 = v4
		goto L38
	} else {
		goto L39
	}
L20:
	;
	v110 = v102 - v5
	goto L19
L21:
	;
	v81 = v77
	goto L30
L22:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v61 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v110 = int32(0)
	goto L19
L24:
	;
	goto L25
L25:
	;
	v66 = v5
	goto L26
L26:
	;
	v70 = v66 + int32(1)
	if v70&int32(3) == int32(0) {
		v77 = v70
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v102 = v70
	goto L20
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v75 != 0 {
		v66 = v70
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 == v90 {
		v81 = v81 + int32(4)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v96 = v81
	goto L33
L32:
	;
	goto L31
L33:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 != 0 {
		v96 = v96 + int32(1)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v102 = v96
	goto L20
L35:
	;
	goto L34
L36:
	;
	v168 = F_varstr_cmp(m, v5, v110, v4, v167, v6)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v167 = v159 - v4
	goto L36
L38:
	;
	v138 = v134
	goto L47
L39:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v118 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v167 = int32(0)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v123 = v4
	goto L43
L43:
	;
	v127 = v123 + int32(1)
	if v127&int32(3) == int32(0) {
		v134 = v127
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v159 = v127
	goto L37
L45:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v132 != 0 {
		v123 = v127
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v147 = int32(-2139062144)
	if (int32(16843008)-v144|v144)&v147 == v147 {
		v138 = v138 + int32(4)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v153 = v138
	goto L50
L49:
	;
	goto L48
L50:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v157 != 0 {
		v153 = v153 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v159 = v153
	goto L37
L52:
	;
	goto L51
L53:
	;
	return int32(0)
L54:
	;
	v172 = v168
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
		v12 = F_DirectFunctionCall1Coll(m, int32(1451), int32(0), v4)
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
		v12 = F_DirectFunctionCall1Coll(m, int32(1451), int32(0), v4)
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
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
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
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v5&int32(3) == int32(0) {
		v34 = v5
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = F_RE_compile_and_cache(m, v7, int32(27), v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v67 = v59 - v5
	goto L3
L5:
	;
	v38 = v34
	goto L14
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v67 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v23 = v5
	goto L10
L10:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v59 = v27
	goto L4
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v53 = v38
	goto L17
L16:
	;
	goto L15
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v59 = v53
	goto L4
L19:
	;
	goto L18
L20:
	;
	v76 = F_palloc(m, v67<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v78 = F_pg_mb2wchar_with_len(m, v5, v76, v67)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v80 = int32(0)
	v83 = F_RE_wchar_execute(m, v76, v78, v80, v80, v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_pfree(m, v76)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	return v83
}
func F_namein(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3&int32(3) == int32(0) {
		v27 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if int32(64) <= v60 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v60 = v52 - v3
	goto L1
L3:
	;
	v31 = v27
	goto L12
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v60 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v16 = v3
	goto L8
L8:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v52 = v20
	goto L2
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v46 = v31
	goto L15
L14:
	;
	goto L13
L15:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v52 = v46
	goto L2
L17:
	;
	goto L16
L18:
	;
	v64 = F_pg_mbcliplen(m, v3, v60, int32(63))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v68 = v60
	goto L20
L20:
	;
	v70 = F_palloc0(m, int32(64))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L21
	} else {
		goto L23
	}
L21:
	;
	return int32(0)
L22:
	;
	v68 = v64
	goto L20
L23:
	;
	if v68 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	return v73
L25:
	;
	v72 = F__emscripten_memcpy_bulkmem(m, v70, v3, v68)
	mBase = m.M
	v73 = v72
	goto L27
L26:
	;
	v73 = v70
	goto L27
L27:
	;
	goto L24
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1557), v3, v4, v5)
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
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
	return base.B2i32(v172 != int32(0))
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	if v5&int32(3) == int32(0) {
		v77 = v5
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v172 = v45 - v46
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
	v41 = v4
	v45 = int32(0)
	goto L11
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	goto L5
L12:
	;
	v41 = v36
	v45 = v38
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 != v21 {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v36 = v30
	v38 = int32(0)
	goto L12
L15:
	;
	if v21 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v26 = v18 - int32(1)
	if v26 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v29 = int32(1)
	v30 = v17 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v31 != 0 {
		v16 = v16 + v29
		v17 = v30
		v18 = v26
		v19 = v31
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v4&int32(3) == int32(0) {
		v134 = v4
		goto L38
	} else {
		goto L39
	}
L20:
	;
	v110 = v102 - v5
	goto L19
L21:
	;
	v81 = v77
	goto L30
L22:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v61 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v110 = int32(0)
	goto L19
L24:
	;
	goto L25
L25:
	;
	v66 = v5
	goto L26
L26:
	;
	v70 = v66 + int32(1)
	if v70&int32(3) == int32(0) {
		v77 = v70
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v102 = v70
	goto L20
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v75 != 0 {
		v66 = v70
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 == v90 {
		v81 = v81 + int32(4)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v96 = v81
	goto L33
L32:
	;
	goto L31
L33:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 != 0 {
		v96 = v96 + int32(1)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v102 = v96
	goto L20
L35:
	;
	goto L34
L36:
	;
	v168 = F_varstr_cmp(m, v5, v110, v4, v167, v6)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v167 = v159 - v4
	goto L36
L38:
	;
	v138 = v134
	goto L47
L39:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v118 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v167 = int32(0)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v123 = v4
	goto L43
L43:
	;
	v127 = v123 + int32(1)
	if v127&int32(3) == int32(0) {
		v134 = v127
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v159 = v127
	goto L37
L45:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v132 != 0 {
		v123 = v127
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v147 = int32(-2139062144)
	if (int32(16843008)-v144|v144)&v147 == v147 {
		v138 = v138 + int32(4)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v153 = v138
	goto L50
L49:
	;
	goto L48
L50:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v157 != 0 {
		v153 = v153 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v159 = v153
	goto L37
L52:
	;
	goto L51
L53:
	;
	return int32(0)
L54:
	;
	v172 = v168
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
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = v8 + int32(1)
	if v6&int32(3) == int32(0) {
		v37 = v6
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v75 = v73 & int32(1)
	if v75 != 0 {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v70 = v62 - v6
	goto L3
L5:
	;
	v41 = v37
	goto L14
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v70 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v26 = v6
	goto L10
L10:
	;
	v30 = v26 + int32(1)
	if v30&int32(3) == int32(0) {
		v37 = v30
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v62 = v30
	goto L4
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 != 0 {
		v26 = v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v50 = int32(-2139062144)
	if (int32(16843008)-v47|v47)&v50 == v50 {
		v41 = v41 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v56 = v41
	goto L17
L16:
	;
	goto L15
L17:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 != 0 {
		v56 = v56 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v62 = v56
	goto L4
L19:
	;
	goto L18
L20:
	;
	v76 = v13
	goto L22
L21:
	;
	v76 = v8 + int32(4)
	goto L22
L22:
	;
	if v73 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v106 = F_GenericMatchText(m, v6, v70, v76, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L34
	}
L24:
	;
	v79 = int32(4)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v81&int32(254) == int32(2) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v94 = int32(1)
	if v75 != 0 {
		v104 = int32(base.Ui32(v73)>>(uint(v94)%32)) - v94
		goto L23
	} else {
		goto L33
	}
L27:
	;
	v90 = v79
	goto L29
L28:
	;
	v90 = base.B2i32(v81 == int32(18)) << (uint(v79) % 32)
	goto L29
L29:
	;
	if v81 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v93 = v79
	goto L32
L31:
	;
	v93 = v90
	goto L32
L32:
	;
	v104 = v93
	goto L23
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v104 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
	goto L23
L34:
	;
	return base.B2i32(v106 != int32(1))
}
func F_namestrcmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
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
	if l0 == int32(0) {
		v59 = v9
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return v59
L8:
	;
	if l1 == int32(0) {
		v59 = v9
		goto L7
	} else {
		goto L9
	}
L9:
	;
	goto L12
L10:
	;
	v59 = v50 - v51
	goto L7
L12:
	;
	goto L13
L13:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v20 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v21 = l0
	v22 = l1
	v23 = int32(64)
	v24 = v20
	goto L18
L15:
	;
	v46 = l1
	v50 = int32(0)
	goto L16
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	goto L10
L17:
	;
	v46 = v41
	v50 = v43
	goto L16
L18:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v24 != v26 {
		v41 = v22
		v43 = v24
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v41 = v35
	v43 = int32(0)
	goto L17
L20:
	;
	if v26 == int32(0) {
		v41 = v22
		v43 = v24
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v31 = v23 - int32(1)
	if v31 == int32(0) {
		v41 = v22
		v43 = v24
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v34 = int32(1)
	v35 = v22 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v36 != 0 {
		v21 = v21 + v34
		v22 = v35
		v23 = v31
		v24 = v36
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
}
func F_new_head_cell(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v199 int32
	_ = v199
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 <= v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(1)
	v10 = int32(16)
	v12 = v5 + v8
	if v12 <= v10 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v47 = v5
	goto L3
L3:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v52 = v50 + int32(4)
	v54 = v47 << (uint(int32(2)) % 32)
	if v52 == v50 {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v15 = v10
	goto L6
L5:
	;
	v15 = v12
	goto L6
L6:
	;
	if v15&(v15-int32(1)) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = v8 << (uint(int32(32)-base.I32_clz(v15)) % 32)
	goto L9
L8:
	;
	v22 = v15
	goto L9
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v25 = l0 + int32(16)
	if v23 == v25 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v47 = v46
	goto L3
L11:
	;
	v27 = F_GetMemoryChunkContext(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v41 = F_repalloc(m, v23, v22<<(uint(int32(2))%32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L21
	}
L14:
	;
	return
L15:
	;
	v31 = F_MemoryContextAlloc(m, v27, v22<<(uint(int32(2))%32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = v34 << (uint(int32(2)) % 32)
	if v36 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L10
L18:
	;
	v37 = F__emscripten_memcpy_bulkmem(m, v31, v25, v36)
	mBase = m.M
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v41
	goto L10
L22:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v199 + int32(1)
	return
L23:
	;
	goto L22
L24:
	;
	v58 = v52 + v54
	if base.Ui32(v50-v58) <= base.Ui32(int32(0)-v54<<(uint(int32(1))%32)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v65 = F___memcpy(m, v52, v50, v54)
	mBase = m.M
	goto L22
L26:
	;
	goto L27
L27:
	;
	v68 = (v52 ^ v50) & int32(3)
	if base.Ui32(v52) < base.Ui32(v50) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v170 == int32(0) {
		goto L23
	} else {
		goto L64
	}
L29:
	;
	if base.Ui32(v148) <= base.Ui32(int32(3)) {
		v169 = v147
		v170 = v148
		v171 = v149
		goto L28
	} else {
		goto L60
	}
L30:
	;
	if v68 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if v68 != 0 {
		v130 = v54
		goto L43
	} else {
		goto L44
	}
L33:
	;
	v169 = v50
	v170 = v54
	v171 = v52
	goto L28
L34:
	;
	goto L35
L35:
	;
	if v52&int32(3) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v147 = v50
	v148 = v54
	v149 = v52
	goto L29
L37:
	;
	goto L38
L38:
	;
	v75 = v50
	v76 = v54
	v77 = v52
	goto L39
L39:
	;
	if v76 == int32(0) {
		goto L23
	} else {
		goto L41
	}
L40:
	;
	v147 = v84
	v148 = v86
	v149 = v88
	goto L29
L41:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v81)
	v83 = int32(1)
	v84 = v75 + v83
	v86 = v76 - v83
	v88 = v77 + v83
	if v88&int32(3) != 0 {
		v75 = v84
		v76 = v86
		v77 = v88
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	if v130 == int32(0) {
		goto L23
	} else {
		goto L56
	}
L44:
	;
	if v58&int32(3) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v95 = v54
	goto L48
L46:
	;
	v110 = v54
	goto L47
L47:
	;
	if base.Ui32(v110) <= base.Ui32(int32(3)) {
		v130 = v110
		goto L43
	} else {
		goto L52
	}
L48:
	;
	if v95 == int32(0) {
		goto L23
	} else {
		goto L50
	}
L49:
	;
	v110 = v101
	goto L47
L50:
	;
	v101 = v95 - int32(1)
	v102 = v52 + v101
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v101))))
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v104)
	if v102&int32(3) != 0 {
		v95 = v101
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v117 = v110
	goto L53
L53:
	;
	v121 = v117 - int32(4)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v50+v121)))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v121))) = v124
	if base.Ui32(int32(3)) < base.Ui32(v121) {
		v117 = v121
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v130 = v121
	goto L43
L55:
	;
	goto L54
L56:
	;
	v137 = v130
	goto L57
L57:
	;
	v141 = v137 - int32(1)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v141))))
	*(*uint8)(unsafe.Add(mBase, uint32(v52+v141))) = uint8(v144)
	if v141 != 0 {
		v137 = v141
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L23
L59:
	;
	goto L58
L60:
	;
	v154 = v147
	v155 = v148
	v156 = v149
	goto L61
L61:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v158
	v160 = int32(4)
	v161 = v154 + v160
	v163 = v156 + v160
	v165 = v155 - v160
	if base.Ui32(int32(3)) < base.Ui32(v165) {
		v154 = v161
		v155 = v165
		v156 = v163
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v169 = v161
	v170 = v165
	v171 = v163
	goto L28
L63:
	;
	goto L62
L64:
	;
	v176 = v169
	v177 = v170
	v178 = v171
	goto L65
L65:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v180)
	v182 = int32(1)
	v187 = v177 - v182
	if v187 != 0 {
		v176 = v176 + v182
		v177 = v187
		v178 = v178 + v182
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L23
L67:
	;
	goto L66
}
