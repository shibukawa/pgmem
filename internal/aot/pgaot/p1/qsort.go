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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int64
	_ = v362
	var v364 int64
	_ = v364
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int64
	_ = v451
	var v453 int64
	_ = v453
	var v456 int32
	_ = v456
	var v457 int64
	_ = v457
	var v459 int64
	_ = v459
	var v461 int64
	_ = v461
	var v463 int64
	_ = v463
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v511 int32
	_ = v511
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int64
	_ = v528
	var v530 int64
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int64
	_ = v535
	var v537 int64
	_ = v537
	var v539 int64
	_ = v539
	var v541 int64
	_ = v541
	var v544 int32
	_ = v544
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int64
	_ = v599
	var v601 int64
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int64
	_ = v606
	var v608 int64
	_ = v608
	var v610 int64
	_ = v610
	var v612 int64
	_ = v612
	var v615 int32
	_ = v615
	var v638 int32
	_ = v638
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int64
	_ = v656
	var v658 int64
	_ = v658
	var v661 int32
	_ = v661
	var v662 int64
	_ = v662
	var v664 int64
	_ = v664
	var v666 int64
	_ = v666
	var v668 int64
	_ = v668
	var v670 int32
	_ = v670
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = l0
	v23 = l1
	goto L1
L1:
	;
	v40 = v22 + int32(16)
	v42 = v23
	goto L3
L2:
	;
	m.G0 = v20 + int32(16)
	return
L3:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v59 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L2
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v64 = v22 + v42<<(uint(int32(4))%32)
	if base.Ui32(v42) <= base.Ui32(int32(6)) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	goto L4
L11:
	;
	if base.Ui32(v64) <= base.Ui32(v40) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v64) <= base.Ui32(v40) {
		goto L10
	} else {
		goto L39
	}
L14:
	;
	v81 = v40
	goto L15
L15:
	;
	if base.Ui32(v81) <= base.Ui32(v22) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v174 = v81 + int32(16)
	if base.Ui32(v174) < base.Ui32(v64) {
		v81 = v174
		goto L15
	} else {
		goto L38
	}
L18:
	;
	v90 = v81
	goto L19
L19:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+8)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90-int32(8)))))
	if v106 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L17
L21:
	;
	v135 = int32(8)
	v136 = v20 + v135
	v138 = v90 + v135
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v136))) = v139
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v141
	v144 = v90 - int32(16)
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
	*(*int64)(unsafe.Add(mBase, uint32(v90))) = v145
	v148 = v144 + v135
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v149
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v144))) = v151
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	*(*int64)(unsafe.Add(mBase, uint32(v148))) = v153
	if base.Ui32(v22) < base.Ui32(v144) {
		v90 = v144
		goto L19
	} else {
		goto L37
	}
L22:
	;
	if v103&int32(1) != 0 {
		goto L17
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v103&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v111 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L17
L27:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v116 != 0 {
		goto L21
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v90-int32(12))))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v122 = m.T0[v121].(func(*base.Module, int32, int32, int32) int32)(m, v119, v120, l2)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L31
	}
L30:
	;
	goto L17
L31:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v124 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v122 < int32(0) {
		goto L21
	} else {
		goto L35
	}
L33:
	;
	v131 = v122
	goto L34
L34:
	;
	if v131 <= int32(0) {
		goto L17
	} else {
		goto L36
	}
L35:
	;
	v131 = int32(0) - v122
	goto L34
L36:
	;
	goto L21
L37:
	;
	goto L20
L38:
	;
	goto L16
L39:
	;
	v180 = v40
	goto L40
L40:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v195 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v240 = v22 + v42<<(uint(int32(3))%32)&int32(-16)
	if v42 != int32(7) {
		goto L64
	} else {
		goto L65
	}
L42:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+8)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180-int32(8)))))
	if v201 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L44
L46:
	;
	goto L41
L47:
	;
	v233 = v180 + int32(16)
	if base.Ui32(v233) < base.Ui32(v64) {
		v180 = v233
		goto L40
	} else {
		goto L63
	}
L48:
	;
	if v198&int32(1) != 0 {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v198&int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v206 == int32(0) {
		goto L46
	} else {
		goto L52
	}
L52:
	;
	goto L47
L53:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v211 == int32(0) {
		goto L47
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v180-int32(12))))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v219 = m.T0[v218].(func(*base.Module, int32, int32, int32) int32)(m, v216, v217, l2)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L57
	}
L56:
	;
	goto L46
L57:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v221 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v219 < int32(0) {
		goto L46
	} else {
		goto L61
	}
L59:
	;
	v228 = v219
	goto L60
L60:
	;
	if int32(0) < v228 {
		goto L46
	} else {
		goto L62
	}
L61:
	;
	v228 = int32(0) - v219
	goto L60
L62:
	;
	goto L47
L63:
	;
	goto L10
L64:
	;
	v244 = v64 - int32(16)
	if base.Ui32(v42) < base.Ui32(int32(41)) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v272 = v240
	goto L66
L66:
	;
	v277 = int32(8)
	v278 = v20 + v277
	v280 = v22 + v277
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v280)))
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v281
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v283
	v286 = v272 + v277
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v286)))
	*(*int64)(unsafe.Add(mBase, uint32(v280))) = v287
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v272)))
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v289
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	*(*int64)(unsafe.Add(mBase, uint32(v286))) = v291
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v272))) = v293
	v296 = v64 - int32(16)
	v300 = v296
	v301 = v40
	v305 = v40
	v309 = v296
	goto L75
L67:
	;
	v270 = F_qsort_ssup_med3(m, v268, v265, v266, l2)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L8
	} else {
		goto L74
	}
L68:
	;
	v265 = v240
	v266 = v244
	v268 = v22
	goto L67
L69:
	;
	goto L70
L70:
	;
	v248 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
	v250 = v248 << (uint(int32(4)) % 32)
	v253 = v248 << (uint(int32(5)) % 32)
	v255 = F_qsort_ssup_med3(m, v22, v22+v250, v22+v253, l2)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	v259 = F_qsort_ssup_med3(m, v240-v250, v240, v240+v250, l2)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v263 = F_qsort_ssup_med3(m, v244-v253, v244-v250, v244, l2)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	v265 = v259
	v266 = v263
	v268 = v255
	goto L67
L74:
	;
	v272 = v270
	goto L66
L75:
	;
	if base.Ui32(v300) < base.Ui32(v301) {
		v392 = v301
		v396 = v305
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if base.Ui32(v392) <= base.Ui32(v300) {
		goto L105
	} else {
		goto L106
	}
L78:
	;
	v319 = v301
	v323 = v305
	goto L79
L79:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+8)))
	if v333 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v392 = v386
	v396 = v379
	goto L77
L81:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v382 != 0 {
		goto L99
	} else {
		goto L100
	}
L82:
	;
	v360 = int32(8)
	v361 = v323 + v360
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v361)))
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v362
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v323)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v364
	v367 = v319 + v360
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v367)))
	*(*int64)(unsafe.Add(mBase, uint32(v361))) = v368
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v319)))
	*(*int64)(unsafe.Add(mBase, uint32(v323))) = v370
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	*(*int64)(unsafe.Add(mBase, uint32(v367))) = v372
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v319))) = v374
	v379 = v323 + int32(16)
	goto L81
L83:
	;
	if v332&int32(1) != 0 {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v332&int32(1) != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v338 != 0 {
		v379 = v323
		goto L81
	} else {
		goto L87
	}
L87:
	;
	v392 = v319
	v396 = v323
	goto L77
L88:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v341 == int32(0) {
		v379 = v323
		goto L81
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v347 = m.T0[v346].(func(*base.Module, int32, int32, int32) int32)(m, v344, v345, l2)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L92
	}
L91:
	;
	v392 = v319
	v396 = v323
	goto L77
L92:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v349 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v347 < int32(0) {
		v392 = v319
		v396 = v323
		goto L77
	} else {
		goto L96
	}
L94:
	;
	v356 = v347
	goto L95
L95:
	;
	if int32(0) < v356 {
		v392 = v319
		v396 = v323
		goto L77
	} else {
		goto L97
	}
L96:
	;
	v356 = int32(0) - v347
	goto L95
L97:
	;
	if v356 != 0 {
		v379 = v323
		goto L81
	} else {
		goto L98
	}
L98:
	;
	goto L82
L99:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L8
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v386 = v319 + int32(16)
	if base.Ui32(v386) <= base.Ui32(v300) {
		v319 = v386
		v323 = v379
		goto L79
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	goto L80
L104:
	;
	v654 = int32(8)
	v655 = v392 + v654
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v655)))
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v656
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v392)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v658
	v661 = v409 + v654
	v662 = *(*int64)(unsafe.Add(mBase, uint32(v661)))
	*(*int64)(unsafe.Add(mBase, uint32(v655))) = v662
	v664 = *(*int64)(unsafe.Add(mBase, uint32(v409)))
	*(*int64)(unsafe.Add(mBase, uint32(v392))) = v664
	v666 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	*(*int64)(unsafe.Add(mBase, uint32(v661))) = v666
	v668 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v409))) = v668
	v670 = int32(16)
	v300 = v409 - v670
	v301 = v392 + v670
	v305 = v396
	v309 = v418
	goto L75
L105:
	;
	v409 = v300
	v418 = v309
	goto L108
L106:
	;
	v480 = v300
	v489 = v309
	goto L107
L107:
	;
	v495 = int32(4)
	v496 = (v396 - v22) >> (uint(v495) % 32)
	v499 = (v392 - v396) >> (uint(v495) % 32)
	if v496 < v499 {
		goto L133
	} else {
		goto L134
	}
L108:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+8)))
	if v424 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v480 = v475
	v489 = v469
	goto L107
L110:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v471 != 0 {
		goto L128
	} else {
		goto L129
	}
L111:
	;
	v449 = int32(8)
	v450 = v409 + v449
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v450)))
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v451
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v409)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v453
	v456 = v418 + v449
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v456)))
	*(*int64)(unsafe.Add(mBase, uint32(v450))) = v457
	v459 = *(*int64)(unsafe.Add(mBase, uint32(v418)))
	*(*int64)(unsafe.Add(mBase, uint32(v409))) = v459
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	*(*int64)(unsafe.Add(mBase, uint32(v456))) = v461
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v418))) = v463
	v469 = v418 - int32(16)
	goto L110
L112:
	;
	if v423&int32(1) != 0 {
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if v423&int32(1) != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v429 != 0 {
		goto L104
	} else {
		goto L116
	}
L116:
	;
	v469 = v418
	goto L110
L117:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v432 != 0 {
		v469 = v418
		goto L110
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v436 = m.T0[v435].(func(*base.Module, int32, int32, int32) int32)(m, v433, v434, l2)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L8
	} else {
		goto L121
	}
L120:
	;
	goto L104
L121:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v438 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	if v436 < int32(0) {
		v469 = v418
		goto L110
	} else {
		goto L125
	}
L123:
	;
	v445 = v436
	goto L124
L124:
	;
	if v445 < int32(0) {
		goto L104
	} else {
		goto L126
	}
L125:
	;
	v445 = int32(0) - v436
	goto L124
L126:
	;
	if v445 != 0 {
		v469 = v418
		goto L110
	} else {
		goto L127
	}
L127:
	;
	goto L111
L128:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v475 = v409 - int32(16)
	if base.Ui32(v392) <= base.Ui32(v475) {
		v409 = v475
		v418 = v469
		goto L108
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	goto L109
L133:
	;
	v501 = v496
	goto L135
L134:
	;
	v501 = v499
	goto L135
L135:
	;
	if v501 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v511 = int32(0)
	goto L139
L137:
	;
	goto L138
L138:
	;
	v564 = int32(4)
	v565 = (v489 - v480) >> (uint(v564) % 32)
	v570 = (v64-v489)>>(uint(v564)%32) - int32(1)
	if v565 < v570 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v524 = v511 << (uint(int32(4)) % 32)
	v525 = v22 + v524
	v526 = int32(8)
	v527 = v525 + v526
	v528 = *(*int64)(unsafe.Add(mBase, uint32(v527)))
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v528
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v525)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v530
	v532 = v524 + (v392 - v501<<(uint(int32(4))%32))
	v534 = v532 + v526
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v534)))
	*(*int64)(unsafe.Add(mBase, uint32(v527))) = v535
	v537 = *(*int64)(unsafe.Add(mBase, uint32(v532)))
	*(*int64)(unsafe.Add(mBase, uint32(v525))) = v537
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	*(*int64)(unsafe.Add(mBase, uint32(v534))) = v539
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v532))) = v541
	v544 = v511 + int32(1)
	if v544 != v501 {
		v511 = v544
		goto L139
	} else {
		goto L141
	}
L140:
	;
	goto L138
L141:
	;
	goto L140
L142:
	;
	v572 = v565
	goto L144
L143:
	;
	v572 = v570
	goto L144
L144:
	;
	if v572 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v580 = int32(0)
	goto L148
L146:
	;
	goto L147
L147:
	;
	if base.Ui32(v499) <= base.Ui32(v565) {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v595 = v580 << (uint(int32(4)) % 32)
	v596 = v392 + v595
	v597 = int32(8)
	v598 = v596 + v597
	v599 = *(*int64)(unsafe.Add(mBase, uint32(v598)))
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v599
	v601 = *(*int64)(unsafe.Add(mBase, uint32(v596)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v601
	v603 = v595 + (v64 - v572<<(uint(int32(4))%32))
	v605 = v603 + v597
	v606 = *(*int64)(unsafe.Add(mBase, uint32(v605)))
	*(*int64)(unsafe.Add(mBase, uint32(v598))) = v606
	v608 = *(*int64)(unsafe.Add(mBase, uint32(v603)))
	*(*int64)(unsafe.Add(mBase, uint32(v596))) = v608
	v610 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	*(*int64)(unsafe.Add(mBase, uint32(v605))) = v610
	v612 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v603))) = v612
	v615 = v580 + int32(1)
	if v615 != v572 {
		v580 = v615
		goto L148
	} else {
		goto L150
	}
L149:
	;
	goto L147
L150:
	;
	goto L149
L151:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v499) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v565) {
		goto L159
	} else {
		goto L160
	}
L154:
	;
	F_qsort_ssup(m, v22, v499, l2)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L8
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	if base.Ui32(v565) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L158
	}
L157:
	;
	goto L156
L158:
	;
	v22 = v64 - v565<<(uint(int32(4))%32)
	v23 = v565
	goto L1
L159:
	;
	F_qsort_ssup(m, v64-v565<<(uint(int32(4))%32), v565, l2)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L8
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	if base.Ui32(int32(1)) < base.Ui32(v499) {
		v42 = v499
		goto L3
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	goto L10
}
