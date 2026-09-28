package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsortCompareItemPointers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	v5 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v6 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v7 = int64(32)
	v9 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v10 = int64(48)
	v13 = v5 | (v6<<(uint(v7)%64) | v9<<(uint(v10)%64))
	v14 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v15 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v18 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v22 = v14 | (v15<<(uint(v7)%64) | v18<<(uint(v10)%64))
	return base.B2i32(base.Ui64(v22) < base.Ui64(v13)) - base.B2i32(base.Ui64(v13) < base.Ui64(v22))
}
func F_qsort_ssup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
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
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v296 int64
	_ = v296
	var v298 int64
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v359 int64
	_ = v359
	var v361 int64
	_ = v361
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v367 int64
	_ = v367
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int64
	_ = v427
	var v428 int64
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v443 int64
	_ = v443
	var v445 int64
	_ = v445
	var v447 int64
	_ = v447
	var v449 int64
	_ = v449
	var v451 int64
	_ = v451
	var v453 int64
	_ = v453
	var v455 int64
	_ = v455
	var v457 int64
	_ = v457
	var v459 int64
	_ = v459
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int64
	_ = v515
	var v517 int64
	_ = v517
	var v519 int64
	_ = v519
	var v521 int32
	_ = v521
	var v522 int64
	_ = v522
	var v524 int64
	_ = v524
	var v526 int64
	_ = v526
	var v528 int64
	_ = v528
	var v530 int64
	_ = v530
	var v532 int64
	_ = v532
	var v535 int32
	_ = v535
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int64
	_ = v582
	var v584 int64
	_ = v584
	var v586 int64
	_ = v586
	var v588 int32
	_ = v588
	var v589 int64
	_ = v589
	var v591 int64
	_ = v591
	var v593 int64
	_ = v593
	var v595 int64
	_ = v595
	var v597 int64
	_ = v597
	var v599 int64
	_ = v599
	var v602 int32
	_ = v602
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v630 int64
	_ = v630
	var v632 int64
	_ = v632
	var v634 int64
	_ = v634
	var v636 int64
	_ = v636
	var v638 int64
	_ = v638
	var v640 int64
	_ = v640
	var v642 int64
	_ = v642
	var v644 int64
	_ = v644
	var v646 int64
	_ = v646
	var v648 int32
	_ = v648
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = l0
	v20 = l1
	goto L1
L1:
	;
	v34 = v19 + int32(24)
	v36 = v20
	goto L3
L3:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_ssup[0]))
	if v50 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(v36) <= base.Ui32(int32(6)) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	v239 = v19 + int32(base.Ui32(v36)>>(uint(int32(1))%32))*int32(24)
	if v36 != int32(7) {
		goto L63
	} else {
		goto L64
	}
L11:
	;
	m.G0 = v17 + int32(32)
	return
L12:
	;
	if base.Ui32(v36) < base.Ui32(int32(2)) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v161 = v19 + v36*int32(24)
	v166 = v34
	goto L40
L15:
	;
	v71 = v34
	goto L16
L16:
	;
	if base.Ui32(v71) <= base.Ui32(v19) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L11
L18:
	;
	v157 = v71 + int32(24)
	if base.Ui32(v157) < base.Ui32(v19+v36*int32(24)) {
		v71 = v157
		goto L16
	} else {
		goto L39
	}
L19:
	;
	v78 = v71
	goto L20
L20:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+16)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78-int32(8)))))
	if v92 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L18
L22:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v78)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v121
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v123
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v125
	v128 = v78 - int32(24)
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v129
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v128)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v128)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+16)) = v133
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = v135
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+8)) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+16)) = v139
	if base.Ui32(v19) < base.Ui32(v128) {
		v78 = v128
		goto L20
	} else {
		goto L38
	}
L23:
	;
	if v89&int32(1) != 0 {
		goto L18
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v89&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v97 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v102 != 0 {
		goto L22
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v78-int32(16))))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v108 = m.T0[v107].(func(*base.Module, int64, int64, int32) int32)(m, v105, v106, l2)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L32
	}
L31:
	;
	goto L18
L32:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v110 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v108 < int32(0) {
		goto L22
	} else {
		goto L36
	}
L34:
	;
	v117 = v108
	goto L35
L35:
	;
	if v117 <= int32(0) {
		goto L18
	} else {
		goto L37
	}
L36:
	;
	v117 = int32(0) - v108
	goto L35
L37:
	;
	goto L22
L38:
	;
	goto L21
L39:
	;
	goto L17
L40:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_ssup[0]))
	if v177 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L11
L42:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+16)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166-int32(8)))))
	if v183 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L44
L46:
	;
	v215 = v166 + int32(24)
	if base.Ui32(v215) < base.Ui32(v161) {
		v166 = v215
		goto L40
	} else {
		goto L62
	}
L47:
	;
	if v180&int32(1) != 0 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v180&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v188 == int32(0) {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	goto L46
L52:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v193 == int32(0) {
		goto L46
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v166-int32(16))))
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v166)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v201 = m.T0[v200].(func(*base.Module, int64, int64, int32) int32)(m, v198, v199, l2)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L56
	}
L55:
	;
	goto L10
L56:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v203 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v201 < int32(0) {
		goto L10
	} else {
		goto L60
	}
L58:
	;
	v210 = v201
	goto L59
L59:
	;
	if int32(0) < v210 {
		goto L10
	} else {
		goto L61
	}
L60:
	;
	v210 = int32(0) - v201
	goto L59
L61:
	;
	goto L46
L62:
	;
	goto L41
L63:
	;
	v243 = v161 - int32(24)
	if base.Ui32(v36) < base.Ui32(int32(41)) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v278 = v239
	goto L65
L65:
	;
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v282
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v284
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v286
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v278)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v288
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v278)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v290
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v292
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v278)+16)) = v294
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v278)+8)) = v296
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v298
	v301 = v161 - int32(24)
	v305 = v34
	v306 = v301
	v308 = v34
	v310 = v301
	goto L74
L66:
	;
	v274 = F_qsort_ssup_med3(m, v268, v270, v269, l2)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L8
	} else {
		goto L73
	}
L67:
	;
	v268 = v19
	v269 = v243
	v270 = v239
	goto L66
L68:
	;
	goto L69
L69:
	;
	v247 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	v249 = v247 * int32(24)
	v254 = F_qsort_ssup_med3(m, v19, v19+v249, v19+v247*int32(48), l2)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v257 = v247 * int32(-24)
	v260 = F_qsort_ssup_med3(m, v239+v257, v239, v239+v249, l2)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	v266 = F_qsort_ssup_med3(m, v243+v247*int32(-48), v243+v257, v243, l2)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v268 = v254
	v269 = v266
	v270 = v260
	goto L66
L73:
	;
	v278 = v274
	goto L65
L74:
	;
	if base.Ui32(v306) < base.Ui32(v305) {
		v391 = v305
		v394 = v308
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if base.Ui32(v391) <= base.Ui32(v306) {
		goto L104
	} else {
		goto L105
	}
L77:
	;
	v320 = v305
	v323 = v308
	goto L78
L78:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+16)))
	if v332 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v391 = v386
	v394 = v380
	goto L76
L80:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_ssup[0]))
	if v382 != 0 {
		goto L98
	} else {
		goto L99
	}
L81:
	;
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v323)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v359
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v323)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v361
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v323)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v363
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v320)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v323)+16)) = v365
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v320)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v323)+8)) = v367
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v320)))
	*(*int64)(unsafe.Add(mBase, uint32(v323))) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+16)) = v371
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v373
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v375
	v380 = v323 + int32(24)
	goto L80
L82:
	;
	if v331&int32(1) != 0 {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v331&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v337 != 0 {
		v380 = v323
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v391 = v320
	v394 = v323
	goto L76
L87:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v340 == int32(0) {
		v380 = v323
		goto L80
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v320)+8))
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v346 = m.T0[v345].(func(*base.Module, int64, int64, int32) int32)(m, v343, v344, l2)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L8
	} else {
		goto L91
	}
L90:
	;
	v391 = v320
	v394 = v323
	goto L76
L91:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v348 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v346 < int32(0) {
		v391 = v320
		v394 = v323
		goto L76
	} else {
		goto L95
	}
L93:
	;
	v355 = v346
	goto L94
L94:
	;
	if int32(0) < v355 {
		v391 = v320
		v394 = v323
		goto L76
	} else {
		goto L96
	}
L95:
	;
	v355 = int32(0) - v346
	goto L94
L96:
	;
	if v355 != 0 {
		v380 = v323
		goto L80
	} else {
		goto L97
	}
L97:
	;
	goto L81
L98:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L8
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v386 = v320 + int32(24)
	if base.Ui32(v386) <= base.Ui32(v306) {
		v320 = v386
		v323 = v380
		goto L78
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	goto L79
L103:
	;
	v630 = *(*int64)(unsafe.Add(mBase, uint32(v391)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v630
	v632 = *(*int64)(unsafe.Add(mBase, uint32(v391)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v632
	v634 = *(*int64)(unsafe.Add(mBase, uint32(v391)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v634
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v407)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v391)+16)) = v636
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v407)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v391)+8)) = v638
	v640 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
	*(*int64)(unsafe.Add(mBase, uint32(v391))) = v640
	v642 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v407)+16)) = v642
	v644 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v407)+8)) = v644
	v646 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v407))) = v646
	v648 = int32(24)
	v305 = v391 + v648
	v306 = v407 - v648
	v308 = v394
	v310 = v411
	goto L74
L104:
	;
	v407 = v306
	v411 = v310
	goto L107
L105:
	;
	v476 = v306
	v480 = v310
	goto L106
L106:
	;
	v487 = int32(24)
	v488 = base.I32_div_s(v394-v19, v487)
	v491 = base.I32_div_s(v391-v394, v487)
	if v488 < v491 {
		goto L132
	} else {
		goto L133
	}
L107:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)))
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+16)))
	if v418 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v476 = v470
	v480 = v464
	goto L106
L109:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_ssup[0]))
	if v466 != 0 {
		goto L127
	} else {
		goto L128
	}
L110:
	;
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v407)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v443
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v407)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v445
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v447
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v411)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v407)+16)) = v449
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v411)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v407)+8)) = v451
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v411)))
	*(*int64)(unsafe.Add(mBase, uint32(v407))) = v453
	v455 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v411)+16)) = v455
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v411)+8)) = v457
	v459 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v411))) = v459
	v464 = v411 - int32(24)
	goto L109
L111:
	;
	if v417&int32(1) != 0 {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v417&int32(1) != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v423 != 0 {
		goto L103
	} else {
		goto L115
	}
L115:
	;
	v464 = v411
	goto L109
L116:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v426 != 0 {
		v464 = v411
		goto L109
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v427 = *(*int64)(unsafe.Add(mBase, uint32(v407)+8))
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v430 = m.T0[v429].(func(*base.Module, int64, int64, int32) int32)(m, v427, v428, l2)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L8
	} else {
		goto L120
	}
L119:
	;
	goto L103
L120:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v432 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if v430 < int32(0) {
		v464 = v411
		goto L109
	} else {
		goto L124
	}
L122:
	;
	v439 = v430
	goto L123
L123:
	;
	if v439 < int32(0) {
		goto L103
	} else {
		goto L125
	}
L124:
	;
	v439 = int32(0) - v430
	goto L123
L125:
	;
	if v439 != 0 {
		v464 = v411
		goto L109
	} else {
		goto L126
	}
L126:
	;
	goto L110
L127:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L8
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v470 = v407 - int32(24)
	if base.Ui32(v391) <= base.Ui32(v470) {
		v407 = v470
		v411 = v464
		goto L107
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	goto L108
L132:
	;
	v493 = v488
	goto L134
L133:
	;
	v493 = v491
	goto L134
L134:
	;
	if v493 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v508 = int32(0)
	goto L138
L136:
	;
	goto L137
L137:
	;
	v552 = int32(24)
	v553 = base.I32_div_s(v480-v476, v552)
	v556 = base.I32_div_s(v161-v480, v552)
	v558 = v556 - int32(1)
	if v553 < v558 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v513 = v508 * int32(24)
	v514 = v19 + v513
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v514)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v515
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v514)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v517
	v519 = *(*int64)(unsafe.Add(mBase, uint32(v514)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v519
	v521 = v513 + (v391 + v493*int32(-24))
	v522 = *(*int64)(unsafe.Add(mBase, uint32(v521)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v514)+16)) = v522
	v524 = *(*int64)(unsafe.Add(mBase, uint32(v521)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v514)+8)) = v524
	v526 = *(*int64)(unsafe.Add(mBase, uint32(v521)))
	*(*int64)(unsafe.Add(mBase, uint32(v514))) = v526
	v528 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v521)+16)) = v528
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v521)+8)) = v530
	v532 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v521))) = v532
	v535 = v508 + int32(1)
	if v535 != v493 {
		v508 = v535
		goto L138
	} else {
		goto L140
	}
L139:
	;
	goto L137
L140:
	;
	goto L139
L141:
	;
	v560 = v553
	goto L143
L142:
	;
	v560 = v558
	goto L143
L143:
	;
	if v560 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v571 = int32(0)
	goto L147
L145:
	;
	goto L146
L146:
	;
	if base.Ui32(v491) <= base.Ui32(v553) {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v580 = v571 * int32(24)
	v581 = v391 + v580
	v582 = *(*int64)(unsafe.Add(mBase, uint32(v581)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v582
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v581)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v584
	v586 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v586
	v588 = v580 + (v161 + v560*int32(-24))
	v589 = *(*int64)(unsafe.Add(mBase, uint32(v588)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v581)+16)) = v589
	v591 = *(*int64)(unsafe.Add(mBase, uint32(v588)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v581)+8)) = v591
	v593 = *(*int64)(unsafe.Add(mBase, uint32(v588)))
	*(*int64)(unsafe.Add(mBase, uint32(v581))) = v593
	v595 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v588)+16)) = v595
	v597 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v588)+8)) = v597
	v599 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v588))) = v599
	v602 = v571 + int32(1)
	if v602 != v560 {
		v571 = v602
		goto L147
	} else {
		goto L149
	}
L148:
	;
	goto L146
L149:
	;
	goto L148
L150:
	;
	F_qsort_ssup(m, v19, v491, l2)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L8
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	F_qsort_ssup(m, v161+v553*int32(-24), v553, l2)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L8
	} else {
		goto L154
	}
L153:
	;
	v19 = v161 + v553*int32(-24)
	v20 = v553
	goto L1
L154:
	;
	v36 = v491
	goto L3
}
