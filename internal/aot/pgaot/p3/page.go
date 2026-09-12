package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreePageManagerGetInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v757 int32
	_ = v757
	v4 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = int32(129)
	if base.Ui32(v23) <= base.Ui32(l1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(16)
	return v757
L2:
	;
	v26 = v23
	goto L4
L3:
	;
	v26 = l1
	goto L4
L4:
	;
	v28 = v26 - int32(1)
	if base.Ui32(int32(128)) < base.Ui32(v28) {
		v757 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = l0 - v31 + int32(1)
	v36 = l0 + int32(36)
	v40 = v28
	goto L7
L6:
	;
	if v137 != 0 {
		goto L36
	} else {
		goto L37
	}
L7:
	;
	v57 = v36 + v40<<(uint(int32(2))%32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v58 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v132
	goto L6
L9:
	;
	goto L8
L10:
	;
	v61 = v58 + v34 - int32(1)
	if v40 != int32(128) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	goto L12
L12:
	;
	v146 = v40 + int32(1)
	if v146 != int32(129) {
		v40 = v146
		goto L7
	} else {
		goto L35
	}
L13:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	if v132 != 0 {
		goto L30
	} else {
		goto L31
	}
L14:
	;
	if v101 == int32(0) {
		v757 = v4
		goto L1
	} else {
		goto L29
	}
L15:
	;
	v101 = v61
	goto L14
L16:
	;
	goto L17
L17:
	;
	v70 = v61
	v73 = v4
	goto L18
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if base.Ui32(v84) < base.Ui32(l1) {
		v89 = v73
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v101 = v89
	goto L14
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	if v90 == int32(0) {
		v101 = v89
		goto L14
	} else {
		goto L27
	}
L21:
	;
	if v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if base.Ui32(v86) <= base.Ui32(v84) {
		v89 = v73
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if l1 != v84 {
		v89 = v70
		goto L20
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v121 = v70
	goto L13
L27:
	;
	v93 = v90 + (v34 - int32(1))
	if v93 != 0 {
		v70 = v93
		v73 = v89
		goto L18
	} else {
		goto L28
	}
L28:
	;
	goto L19
L29:
	;
	v121 = v101
	goto L13
L30:
	;
	v137 = v34 + v132 - int32(1)
	goto L32
L31:
	;
	v137 = int32(0)
	goto L32
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v138 == int32(0) {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	v141 = v138 + v34
	if v141 == int32(1) {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+11)) = v132
	goto L6
L35:
	;
	v757 = v4
	goto L1
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+8)) = v152
	goto L38
L37:
	;
	goto L38
L38:
	;
	if v40 != int32(128) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v171 = int32(base.Ui32(v121-v34) >> (uint(int32(12)) % 32))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v172 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v167 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v167)
	goto L39
L41:
	;
	if v40+int32(1) != v161 {
		goto L39
	} else {
		goto L46
	}
L42:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v161 = v157
	goto L41
L43:
	;
	goto L44
L44:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v158 == v159 {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v161 = v158
	goto L41
L46:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v165 != 0 {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v171
	v757 = int32(1)
	goto L1
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v176 = v175 + l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v179 = v178 - l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v179
	if l1 == v178 {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = int32(1)
	if v228 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L52:
	;
	v182 = int32(129)
	if base.Ui32(v182) <= base.Ui32(v179) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v185 = v182
	goto L55
L54:
	;
	v185 = v179
	goto L55
L55:
	;
	v190 = v185<<(uint(int32(2))%32) + v36 - int32(4)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v194 = int32(1)
	v195 = l0 - v192 + v194
	v197 = v176 << (uint(int32(12)) % 32)
	v198 = v195 + v197
	*(*int32)(unsafe.Add(mBase, uint32(v198)+4)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = int32(-364896016)
	v202 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v202
	if v191 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v208 = v191 + v195 - v194
	goto L58
L57:
	;
	v208 = v202
	goto L58
L58:
	;
	if v208 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v213 = v208 - v195 + int32(1)
	goto L61
L60:
	;
	v213 = int32(0)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+12)) = v213
	v216 = v197 | int32(1)
	if v208 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v216
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v216
	goto L48
L65:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if l1 == v409 {
		goto L113
	} else {
		goto L114
	}
L66:
	;
	v392 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v392)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v392
	goto L65
L67:
	;
	v235 = int32(1)
	v236 = l0 - v229 + v235
	v239 = v236 + v228 - v235
	if v239 == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v243 == int32(430584521) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v246 = int32(1)
	v254 = v239
	v256 = v246
	goto L72
L70:
	;
	v328 = int32(2)
	v330 = v239
	goto L71
L71:
	;
	v337 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if base.Ui32(int32(509)) < base.Ui32(v339) {
		goto L94
	} else {
		goto L95
	}
L72:
	;
	v262 = v254 + int32(12)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v265 = int32(0)
	v268 = v263
	goto L74
L73:
	;
	v328 = v309 + int32(1)
	v330 = v319
	goto L71
L74:
	;
	if base.Ui32(v268) <= base.Ui32(v265) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if base.Ui32(v293) < base.Ui32(v263) {
		goto L87
	} else {
		goto L88
	}
L76:
	;
	goto L75
L77:
	;
	v293 = v265
	goto L76
L78:
	;
	goto L79
L79:
	;
	v279 = int32(1)
	v280 = int32(base.Ui32(v265+v268) >> (uint(v279) % 32))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v262+v280<<(uint(int32(3))%32))))
	v287 = base.B2i32(base.Ui32(v171) < base.Ui32(v286))
	if base.Ui32(v171) < base.Ui32(v286) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v288 = v265
	goto L82
L81:
	;
	v288 = v280 + v279
	goto L82
L82:
	;
	if base.Ui32(v171) < base.Ui32(v286) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v289 = v280
	goto L85
L84:
	;
	v289 = v268
	goto L85
L85:
	;
	if v171 != v286 {
		v265 = v288
		v268 = v289
		goto L74
	} else {
		goto L86
	}
L86:
	;
	v293 = v280
	goto L76
L87:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v262+v293<<(uint(int32(3))%32))))
	v303 = base.B2i32(v301 != v171)
	goto L89
L88:
	;
	v303 = int32(1)
	goto L89
L89:
	;
	if base.Ui32(int32(509)) < base.Ui32(v263) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v309 = v256 + int32(1)
	goto L92
L91:
	;
	v309 = int32(0)
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v309
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v254+(v293-base.B2i32(v293 != int32(0))&v303)<<(uint(int32(3))%32))+16))
	v319 = v236 - v246 + v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	if v320 == int32(430584521) {
		v254 = v319
		v256 = v309
		goto L72
	} else {
		goto L93
	}
L93:
	;
	goto L73
L94:
	;
	v342 = v328
	goto L96
L95:
	;
	v342 = v337
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v342
	v345 = v330 + int32(12)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v347 = v337
	v350 = v346
	goto L97
L97:
	;
	if base.Ui32(v350) <= base.Ui32(v347) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v330
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if base.Ui32(v375) < base.Ui32(v381) {
		goto L110
	} else {
		goto L111
	}
L99:
	;
	goto L98
L100:
	;
	v375 = v347
	goto L99
L101:
	;
	goto L102
L102:
	;
	v361 = int32(1)
	v362 = int32(base.Ui32(v347+v350) >> (uint(v361) % 32))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v345+v362<<(uint(int32(3))%32))))
	v369 = base.B2i32(base.Ui32(v171) < base.Ui32(v368))
	if base.Ui32(v171) < base.Ui32(v368) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v370 = v347
	goto L105
L104:
	;
	v370 = v362 + v361
	goto L105
L105:
	;
	if base.Ui32(v171) < base.Ui32(v368) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v371 = v362
	goto L108
L107:
	;
	v371 = v350
	goto L108
L108:
	;
	if v171 != v368 {
		v347 = v370
		v350 = v371
		goto L97
	} else {
		goto L109
	}
L109:
	;
	v375 = v362
	goto L99
L110:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v345+v375<<(uint(int32(3))%32))))
	v388 = base.B2i32(v171 == v386)
	goto L112
L111:
	;
	v388 = int32(0)
	goto L112
L112:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v388)
	goto L65
L113:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	if v422 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	goto L115
L115:
	;
	v549 = v408 + int32(12)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v553 = v549 + v550<<(uint(int32(3))%32)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	*(*int32)(unsafe.Add(mBase, uint32(v553))) = v554 + l1
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v553)+4)) = v557 - l1
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v550 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L116:
	;
	goto L48
L117:
	;
	F_FreePageBtreeRemovePage(m, l0, v408)
	mBase = m.M
	goto L116
L118:
	;
	goto L119
L119:
	;
	v427 = v422 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v408)+4)) = v427
	if base.Ui32(v427) <= base.Ui32(v411) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_FreePageBtreeConsolidate(m, l0, v408)
	mBase = m.M
	goto L116
L121:
	;
	v431 = v408 + int32(12)
	v432 = int32(3)
	v434 = v431 + v411<<(uint(v432)%32)
	v440 = F_memmove(m, v434, v434+int32(8), (v427-v411)<<(uint(v432)%32))
	mBase = m.M
	if v411 != 0 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v442 = l0 - v441
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v451 = v408
	goto L123
L123:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v451)+8))
	if v457 == int32(0) {
		goto L120
	} else {
		goto L125
	}
L124:
	;
	goto L120
L125:
	;
	v460 = v457 + v442
	if v460 == int32(0) {
		goto L120
	} else {
		goto L126
	}
L126:
	;
	v464 = v460 + int32(12)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	v469 = int32(0)
	v470 = v466
	goto L127
L127:
	;
	if base.Ui32(v470) <= base.Ui32(v469) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	if base.Ui32(v496) < base.Ui32(v466) {
		goto L140
	} else {
		goto L141
	}
L129:
	;
	goto L128
L130:
	;
	v496 = v469
	goto L129
L131:
	;
	goto L132
L132:
	;
	v482 = int32(1)
	v483 = int32(base.Ui32(v469+v470) >> (uint(v482) % 32))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v464+v483<<(uint(int32(3))%32))))
	v490 = base.B2i32(base.Ui32(v443) < base.Ui32(v489))
	if base.Ui32(v443) < base.Ui32(v489) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v491 = v469
	goto L135
L134:
	;
	v491 = v483 + v482
	goto L135
L135:
	;
	if base.Ui32(v443) < base.Ui32(v489) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v492 = v483
	goto L138
L137:
	;
	v492 = v470
	goto L138
L138:
	;
	if v443 != v489 {
		v469 = v491
		v470 = v492
		goto L127
	} else {
		goto L139
	}
L139:
	;
	v496 = v483
	goto L129
L140:
	;
	v502 = int32(0)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v460+v496<<(uint(int32(3))%32))+16))
	if v506 != 0 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v513 = int32(-1)
	goto L142
L142:
	;
	v514 = v513 + v496
	*(*int32)(unsafe.Add(mBase, uint32(v464+v514<<(uint(int32(3))%32)))) = v443
	if v514 == int32(0) {
		v451 = v460
		goto L123
	} else {
		goto L149
	}
L143:
	;
	v509 = v442 + v506
	goto L145
L144:
	;
	v509 = v502
	goto L145
L145:
	;
	if v509 != v451 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v511 = int32(-1)
	goto L148
L147:
	;
	v511 = v502
	goto L148
L148:
	;
	v513 = v511
	goto L142
L149:
	;
	goto L124
L150:
	;
	v563 = l0 - v560
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	v577 = v408
	goto L153
L151:
	;
	v674 = v560
	goto L152
L152:
	;
	v689 = int32(129)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v691 = v690 - l1
	if base.Ui32(v689) <= base.Ui32(v691) {
		goto L181
	} else {
		goto L182
	}
L153:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v577)+8))
	if v583 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v674 = v670
	goto L152
L155:
	;
	goto L154
L156:
	;
	v586 = v583 + v563
	if v586 == int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v590 = v586 + int32(12)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	v596 = int32(0)
	v597 = v592
	goto L158
L158:
	;
	if base.Ui32(v597) <= base.Ui32(v596) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	if base.Ui32(v627) < base.Ui32(v592) {
		goto L171
	} else {
		goto L172
	}
L160:
	;
	goto L159
L161:
	;
	v627 = v596
	goto L160
L162:
	;
	goto L163
L163:
	;
	v613 = int32(1)
	v614 = int32(base.Ui32(v596+v597) >> (uint(v613) % 32))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v590+v614<<(uint(int32(3))%32))))
	v621 = base.B2i32(base.Ui32(v564) < base.Ui32(v620))
	if base.Ui32(v564) < base.Ui32(v620) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v622 = v596
	goto L166
L165:
	;
	v622 = v614 + v613
	goto L166
L166:
	;
	if base.Ui32(v564) < base.Ui32(v620) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v623 = v614
	goto L169
L168:
	;
	v623 = v597
	goto L169
L169:
	;
	if v564 != v620 {
		v596 = v622
		v597 = v623
		goto L158
	} else {
		goto L170
	}
L170:
	;
	v627 = v614
	goto L160
L171:
	;
	v633 = int32(0)
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v586+v627<<(uint(int32(3))%32))+16))
	if v637 != 0 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v644 = int32(-1)
	goto L173
L173:
	;
	v645 = v644 + v627
	*(*int32)(unsafe.Add(mBase, uint32(v590+v645<<(uint(int32(3))%32)))) = v564
	if v645 == int32(0) {
		v577 = v586
		goto L153
	} else {
		goto L180
	}
L174:
	;
	v640 = v563 + v637
	goto L176
L175:
	;
	v640 = v633
	goto L176
L176:
	;
	if v640 != v577 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v642 = int32(-1)
	goto L179
L178:
	;
	v642 = v633
	goto L179
L179:
	;
	v644 = v642
	goto L173
L180:
	;
	goto L155
L181:
	;
	v694 = v689
	goto L183
L182:
	;
	v694 = v691
	goto L183
L183:
	;
	v699 = v694<<(uint(int32(2))%32) + v36 - int32(4)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	v702 = int32(1)
	v703 = l0 - v674 + v702
	v706 = (l1 + v171) << (uint(int32(12)) % 32)
	v707 = v703 + v706
	*(*int32)(unsafe.Add(mBase, uint32(v707)+4)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v707))) = int32(-364896016)
	v711 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v707)+8)) = v711
	if v700 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v717 = v700 + v703 - v702
	goto L186
L185:
	;
	v717 = v711
	goto L186
L186:
	;
	if v717 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v722 = v717 - v703 + int32(1)
	goto L189
L188:
	;
	v722 = int32(0)
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v707)+12)) = v722
	v725 = v706 | int32(1)
	if v717 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v717)+8)) = v725
	goto L192
L191:
	;
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699))) = v725
	goto L48
}
func F_PageGetFreeSpaceForMultipleTuples(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v5 = v3 - v4
	v7 = l1 << (uint(int32(2)) % 32)
	if v7 <= v5 {
		v11 = v5 - v7
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_PageRepairFragmentation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1808)
	m.G0 = v20
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v24) < base.Ui32(int32(24)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L38
	} else {
		goto L47
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L38
	} else {
		goto L43
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L38
	} else {
		goto L39
	}
L4:
	;
	if base.Ui32(v23) < base.Ui32(v24) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if base.Ui32(v22) < base.Ui32(v23) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v22) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if (v22+int32(7))&int32(32760) != v22 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v24) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v152 = int32(65535)
	v153 = v145 & v152
	if v153 != v43&v152 {
		goto L35
	} else {
		goto L36
	}
L10:
	;
	v43 = int32(base.Ui32(v24+int32(262120)) >> (uint(int32(2)) % 32))
	goto L12
L11:
	;
	v43 = int32(0)
	goto L12
L12:
	;
	v45 = v43 & int32(65535)
	if v45 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v22)
	v141 = v2
	v145 = v2
	goto L9
L14:
	;
	goto L15
L15:
	;
	v53 = int32(1)
	v59 = v53
	v60 = v20 + int32(48)
	v61 = v2
	v65 = v2
	v66 = v22
	v67 = v2
	v68 = v53
	goto L16
L16:
	;
	v76 = v59<<(uint(int32(2))%32) + (l0 + int32(24)) - int32(4)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77&int32(98304) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v122 = v20 + int32(48)
	v125 = base.I32_div_s(v111-v122, int32(6))
	if v122 == v111 {
		goto L31
	} else {
		goto L32
	}
L18:
	;
	if v59 != v45 {
		v59 = v59 + int32(1)
		v60 = v111
		v61 = v112
		v65 = v114
		v66 = v115
		v67 = v116
		v68 = v117
		goto L16
	} else {
		goto L30
	}
L19:
	;
	if base.Ui32(v77) < base.Ui32(int32(131072)) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = int32(0)
	v111 = v60
	v112 = v61 + int32(1)
	v114 = v65
	v115 = v66
	v116 = v67
	v117 = v68
	goto L18
L22:
	;
	v111 = v60
	v112 = v61
	v114 = v59
	v115 = v66
	v116 = v67
	v117 = v68
	goto L18
L23:
	;
	goto L24
L24:
	;
	v83 = v59 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v83)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v87 = v85 & int32(32767)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)) = uint16(v87)
	if base.Ui32(v87) < base.Ui32(v23) {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	if base.Ui32(v22) <= base.Ui32(v87) {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v97 = (int32(base.Ui32(v91)>>(uint(int32(17))%32)) + int32(7)) & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v97)
	if v66 < v87 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v105 = v66
	goto L29
L28:
	;
	v105 = v87
	goto L29
L29:
	;
	v111 = v60 + int32(6)
	v112 = v61
	v114 = v59
	v115 = v105
	v116 = v97 + v67
	v117 = base.B2i32(v87 < v66) & v68
	goto L18
L30:
	;
	goto L17
L31:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v22)
	v141 = v112
	v145 = v114
	goto L9
L32:
	;
	goto L33
L33:
	;
	v130 = v22 - v24
	if base.Ui32(v130) < base.Ui32(v116) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_compactify_tuples(m, v20+int32(48), v125, l0, v117)
	mBase = m.M
	v141 = v112
	v145 = v114
	goto L9
L35:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v158 = v45 - v153
	v161 = v157 - v158<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v161)
	v165 = v141 - v158
	goto L37
L36:
	;
	v165 = v141
	goto L37
L37:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v171 = v166&int32(65534) | base.B2i32(int32(0) < v165)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v171)
	m.G0 = v20 + int32(1808)
	return
L38:
	;
	return
L39:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v24
	F_errmsg(m, int32(54579), v20)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(476232), int32(730), int32(246583))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v201 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v201
	F_errmsg(m, int32(55466), v20+int32(32))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L38
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(476232), int32(759), int32(246583))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L38
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L38
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v116
	F_errmsg(m, int32(50430), v20+int32(16))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(476232), int32(789), int32(246583))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L38
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
