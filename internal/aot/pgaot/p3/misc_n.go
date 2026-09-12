package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NUM_cache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v715 int64
	_ = v715
	var v717 int32
	_ = v717
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	v5 = int32(0)
	v13 = F_text_to_cstring(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if base.Ui32(int32(75)) <= base.Ui32(l0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = int32(12)
	v23 = F_palloc(m, l0*v19+v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[960]))
	v50 = *(*int32)(unsafe.Add(mBase, _consts[961]))
	if int32(2147483646) <= v50 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v25)
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v27
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v29
	F_parse_format(m, v23, v13, int32(1640624), v27, int32(1640240), int32(2), l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_pfree(m, v13)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	return v23
L9:
	;
	if v48 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v169 = v50
	goto L11
L11:
	;
	if v48 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v162 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, _consts[961])) = v162
	v169 = v162
	goto L11
L13:
	;
	v56 = v48 & int32(3)
	v57 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v48) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = v57
	v68 = v5
	goto L17
L15:
	;
	v109 = v57
	goto L16
L16:
	;
	if v56 == int32(0) {
		goto L12
	} else {
		goto L20
	}
L17:
	;
	v75 = v62 << (uint(int32(2)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[962])))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+976))
	v80 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+976)) = v79 >> (uint(v80) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[963])))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+976)) = v86 >> (uint(v80) % 32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[964])))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+976)) = v93 >> (uint(v80) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[965])))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+976)) = v100 >> (uint(v80) % 32)
	v104 = int32(4)
	v105 = v62 + v104
	v107 = v68 + v104
	if v107 != v48&int32(2147483644) {
		v62 = v105
		v68 = v107
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v109 = v105
	goto L16
L19:
	;
	goto L18
L20:
	;
	v123 = v109
	v132 = v5
	goto L21
L21:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v123<<(uint(int32(2))%32))+uint32(_consts[962])))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+976))
	v141 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+976)) = v140 >> (uint(v141) % 32)
	v147 = v132 + v141
	if v147 != v56 {
		v123 = v123 + v141
		v132 = v147
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L12
L23:
	;
	goto L22
L24:
	;
	v747 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v747)
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v735)+992))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v749
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v735)+988))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v751
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v735)+980))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v753
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v735)+984))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v755
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v735)+996))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v757
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v735)+1012))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v759
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v735)+1000))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v761
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v735)+1004))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v763
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v735)+1008))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v765
	F_pfree(m, v13)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L143
	}
L25:
	;
	v715 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v703)+980)) = v715
	v717 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v703)+1012)) = v717
	*(*int64)(unsafe.Add(mBase, uint32(v703)+1004)) = v715
	*(*int64)(unsafe.Add(mBase, uint32(v703)+996)) = v715
	*(*int64)(unsafe.Add(mBase, uint32(v703)+988)) = v715
	F_parse_format(m, v703, v13, int32(1640624), v717, int32(1640240), int32(2), v703+int32(980))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L142
	}
L26:
	;
	v577 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v565)+975)) = uint8(v577)
	v580 = v565 + int32(900)
	goto L113
L27:
	;
	v422 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v424 = F_MemoryContextAllocZero(m, v422, int32(1016))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L77
	}
L28:
	;
	v181 = int32(0)
	goto L30
L29:
	;
	v406 = v169 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[961])) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v195)+976)) = v406
	v735 = v195
	goto L24
L30:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v181<<(uint(int32(2))%32))+uint32(_consts[962])))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+975)))
	if v196 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if int32(2147483646) <= v169 {
		goto L45
	} else {
		goto L46
	}
L32:
	;
	v200 = v195 + int32(900)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v204 == int32(0) {
		v223 = v203
		v224 = v204
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v229 = v181 + int32(1)
	if v229 != v48 {
		v181 = v229
		goto L30
	} else {
		goto L44
	}
L35:
	;
	if v224-v223 == int32(0) {
		goto L29
	} else {
		goto L43
	}
L36:
	;
	goto L35
L37:
	;
	if v203 != v204 {
		v223 = v203
		v224 = v204
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v208 = v200
	v209 = v13
	goto L39
L39:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if v213 == int32(0) {
		v223 = v212
		v224 = v213
		goto L36
	} else {
		goto L41
	}
L40:
	;
	v223 = v212
	v224 = v213
	goto L36
L41:
	;
	v216 = int32(1)
	if v212 == v213 {
		v208 = v208 + v216
		v209 = v209 + v216
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L34
L44:
	;
	goto L31
L45:
	;
	v234 = v48 & int32(3)
	v235 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v48) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v48 < int32(20) {
		goto L27
	} else {
		goto L60
	}
L48:
	;
	v242 = v235
	v251 = int32(0)
	goto L51
L49:
	;
	v289 = v235
	goto L50
L50:
	;
	if v234 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v255 = v242 << (uint(int32(2)) % 32)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255)+uint32(_consts[962])))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+976))
	v260 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v258)+976)) = v259 >> (uint(v260) % 32)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255)+uint32(_consts[963])))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+976)) = v266 >> (uint(v260) % 32)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v255)+uint32(_consts[964])))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v272)+976)) = v273 >> (uint(v260) % 32)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v255)+uint32(_consts[965])))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+976)) = v280 >> (uint(v260) % 32)
	v284 = int32(4)
	v285 = v242 + v284
	v287 = v251 + v284
	if v287 != v48&int32(2147483644) {
		v242 = v285
		v251 = v287
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v289 = v285
	goto L50
L53:
	;
	goto L52
L54:
	;
	v301 = v289
	v305 = v235
	goto L57
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[961])) = int32(1073741823)
	goto L47
L57:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v301<<(uint(int32(2))%32))+uint32(_consts[962])))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+976))
	v319 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+976)) = v318 >> (uint(v319) % 32)
	v325 = v305 + v319
	if v325 != v234 {
		v301 = v301 + v319
		v305 = v325
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L56
L59:
	;
	goto L58
L60:
	;
	v356 = int32(1)
	v358 = *(*int32)(unsafe.Add(mBase, _consts[962]))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+975)))
	if v359 != v356 {
		v565 = v358
		goto L26
	} else {
		goto L61
	}
L61:
	;
	v362 = v358
	v364 = v356
	goto L62
L62:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v364<<(uint(int32(2))%32))+uint32(_consts[962])))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+975)))
	if v379 != int32(1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v565 = v378
	goto L26
L65:
	;
	goto L66
L66:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v378)+976))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v362)+976))
	if v382 < v383 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v385 = v378
	goto L69
L68:
	;
	v385 = v362
	goto L69
L69:
	;
	v387 = v364 + int32(1)
	if v387 == int32(20) {
		v565 = v385
		goto L26
	} else {
		goto L70
	}
L70:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v387<<(uint(int32(2))%32))+uint32(_consts[962])))
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+975)))
	if v395 != int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v565 = v394
	goto L26
L72:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394)+976))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v385)+976))
	if v398 < v399 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v401 = v394
	goto L76
L75:
	;
	v401 = v385
	goto L76
L76:
	;
	v362 = v401
	v364 = v364 + int32(2)
	goto L62
L77:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _consts[960]))
	*(*int32)(unsafe.Add(mBase, uint32(v427<<(uint(int32(2))%32))+uint32(_consts[962]))) = v424
	v433 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+975)) = uint8(v433)
	v436 = v424 + int32(900)
	goto L81
L78:
	;
	v552 = int32(4455104)
	v554 = *(*int32)(unsafe.Add(mBase, _consts[961]))
	v555 = int32(1)
	v556 = v554 + v555
	*(*int32)(unsafe.Add(mBase, _consts[961])) = v556
	*(*int32)(unsafe.Add(mBase, uint32(v424)+976)) = v556
	v559 = int32(4455012)
	v561 = *(*int32)(unsafe.Add(mBase, _consts[960]))
	*(*int32)(unsafe.Add(mBase, _consts[960])) = v561 + v555
	v703 = v424
	goto L25
L79:
	;
	v549 = F_strlen(m, v538)
	mBase = m.M
	goto L78
L81:
	;
	goto L82
L82:
	;
	v443 = int32(74)
	if (v436^v13)&int32(3) != 0 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v542 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v539))) = uint8(v542)
	goto L79
L84:
	;
	v523 = v518
	v524 = v519
	v525 = v520
	goto L106
L85:
	;
	if v513 == int32(0) {
		v538 = v511
		v539 = v512
		goto L83
	} else {
		goto L105
	}
L86:
	;
	v511 = v13
	v512 = v436
	v513 = v443
	goto L85
L87:
	;
	goto L88
L88:
	;
	if v13&int32(3) == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v480 == int32(0) {
		v538 = v477
		v539 = v478
		goto L83
	} else {
		goto L98
	}
L90:
	;
	v477 = v13
	v478 = v436
	v479 = v443
	v480 = int32(1)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v456 = v13
	v457 = v436
	v458 = v443
	goto L93
L93:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	*(*uint8)(unsafe.Add(mBase, uint32(v457))) = uint8(v460)
	if v460 == int32(0) {
		v518 = v456
		v519 = v457
		v520 = v458
		goto L84
	} else {
		goto L95
	}
L94:
	;
	v477 = v471
	v478 = v465
	v479 = v467
	v480 = v469
	goto L89
L95:
	;
	v464 = int32(1)
	v465 = v457 + v464
	v467 = v458 - v464
	v468 = int32(0)
	v469 = base.B2i32(v467 != v468)
	v471 = v456 + v464
	if v471&int32(3) == v468 {
		v477 = v471
		v478 = v465
		v479 = v467
		v480 = v469
		goto L89
	} else {
		goto L96
	}
L96:
	;
	if v467 != 0 {
		v456 = v471
		v457 = v465
		v458 = v467
		goto L93
	} else {
		goto L97
	}
L97:
	;
	goto L94
L98:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if v483 == int32(0) {
		v511 = v477
		v512 = v478
		v513 = v479
		goto L85
	} else {
		goto L99
	}
L99:
	;
	if base.Ui32(v479) < base.Ui32(int32(4)) {
		v511 = v477
		v512 = v478
		v513 = v479
		goto L85
	} else {
		goto L100
	}
L100:
	;
	v489 = v477
	v490 = v478
	v491 = v479
	goto L101
L101:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	v497 = int32(-2139062144)
	if (int32(16843008)-v494|v494)&v497 != v497 {
		v518 = v489
		v519 = v490
		v520 = v491
		goto L84
	} else {
		goto L103
	}
L102:
	;
	v511 = v505
	v512 = v503
	v513 = v507
	goto L85
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490))) = v494
	v502 = int32(4)
	v503 = v490 + v502
	v505 = v489 + v502
	v507 = v491 - v502
	if base.Ui32(int32(3)) < base.Ui32(v507) {
		v489 = v505
		v490 = v503
		v491 = v507
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v518 = v511
	v519 = v512
	v520 = v513
	goto L84
L106:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v527)
	if v527 == int32(0) {
		v538 = v523
		v539 = v524
		goto L83
	} else {
		goto L108
	}
L107:
	;
	v538 = v534
	v539 = v532
	goto L83
L108:
	;
	v531 = int32(1)
	v532 = v524 + v531
	v534 = v523 + v531
	v536 = v525 - v531
	if v536 != 0 {
		v523 = v534
		v524 = v532
		v525 = v536
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v696 = int32(4455104)
	v698 = *(*int32)(unsafe.Add(mBase, _consts[961]))
	v700 = v698 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[961])) = v700
	*(*int32)(unsafe.Add(mBase, uint32(v565)+976)) = v700
	v703 = v565
	goto L25
L111:
	;
	v693 = F_strlen(m, v682)
	mBase = m.M
	goto L110
L113:
	;
	goto L114
L114:
	;
	v587 = int32(74)
	if (v580^v13)&int32(3) != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v686 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v683))) = uint8(v686)
	goto L111
L116:
	;
	v667 = v662
	v668 = v663
	v669 = v664
	goto L138
L117:
	;
	if v657 == int32(0) {
		v682 = v655
		v683 = v656
		goto L115
	} else {
		goto L137
	}
L118:
	;
	v655 = v13
	v656 = v580
	v657 = v587
	goto L117
L119:
	;
	goto L120
L120:
	;
	if v13&int32(3) == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v624 == int32(0) {
		v682 = v621
		v683 = v622
		goto L115
	} else {
		goto L130
	}
L122:
	;
	v621 = v13
	v622 = v580
	v623 = v587
	v624 = int32(1)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v600 = v13
	v601 = v580
	v602 = v587
	goto L125
L125:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	*(*uint8)(unsafe.Add(mBase, uint32(v601))) = uint8(v604)
	if v604 == int32(0) {
		v662 = v600
		v663 = v601
		v664 = v602
		goto L116
	} else {
		goto L127
	}
L126:
	;
	v621 = v615
	v622 = v609
	v623 = v611
	v624 = v613
	goto L121
L127:
	;
	v608 = int32(1)
	v609 = v601 + v608
	v611 = v602 - v608
	v612 = int32(0)
	v613 = base.B2i32(v611 != v612)
	v615 = v600 + v608
	if v615&int32(3) == v612 {
		v621 = v615
		v622 = v609
		v623 = v611
		v624 = v613
		goto L121
	} else {
		goto L128
	}
L128:
	;
	if v611 != 0 {
		v600 = v615
		v601 = v609
		v602 = v611
		goto L125
	} else {
		goto L129
	}
L129:
	;
	goto L126
L130:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	if v627 == int32(0) {
		v655 = v621
		v656 = v622
		v657 = v623
		goto L117
	} else {
		goto L131
	}
L131:
	;
	if base.Ui32(v623) < base.Ui32(int32(4)) {
		v655 = v621
		v656 = v622
		v657 = v623
		goto L117
	} else {
		goto L132
	}
L132:
	;
	v633 = v621
	v634 = v622
	v635 = v623
	goto L133
L133:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	v641 = int32(-2139062144)
	if (int32(16843008)-v638|v638)&v641 != v641 {
		v662 = v633
		v663 = v634
		v664 = v635
		goto L116
	} else {
		goto L135
	}
L134:
	;
	v655 = v649
	v656 = v647
	v657 = v651
	goto L117
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v634))) = v638
	v646 = int32(4)
	v647 = v634 + v646
	v649 = v633 + v646
	v651 = v635 - v646
	if base.Ui32(int32(3)) < base.Ui32(v651) {
		v633 = v649
		v634 = v647
		v635 = v651
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v662 = v655
	v663 = v656
	v664 = v657
	goto L116
L138:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	*(*uint8)(unsafe.Add(mBase, uint32(v668))) = uint8(v671)
	if v671 == int32(0) {
		v682 = v667
		v683 = v668
		goto L115
	} else {
		goto L140
	}
L139:
	;
	v682 = v678
	v683 = v676
	goto L115
L140:
	;
	v675 = int32(1)
	v676 = v668 + v675
	v678 = v667 + v675
	v680 = v669 - v675
	if v680 != 0 {
		v667 = v678
		v668 = v676
		v669 = v680
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v733 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v703)+975)) = uint8(v733)
	v735 = v703
	goto L24
L143:
	;
	return v735
}
func F_NUM_processor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v237 int32
	_ = v237
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v362 int32
	_ = v362
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v884 int32
	_ = v884
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1212 int32
	_ = v1212
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1284 int32
	_ = v1284
	var v1291 int32
	_ = v1291
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1436 int32
	_ = v1436
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1506 int32
	_ = v1506
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1601 int32
	_ = v1601
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1615 int32
	_ = v1615
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1677 int32
	_ = v1677
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1705 int32
	_ = v1705
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1744 int32
	_ = v1744
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1861 int32
	_ = v1861
	var v1863 int32
	_ = v1863
	var v1877 int32
	_ = v1877
	var v1882 int32
	_ = v1882
	var v1886 int32
	_ = v1886
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1903 int32
	_ = v1903
	var v1906 int32
	_ = v1906
	var v1912 int32
	_ = v1912
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1926 int32
	_ = v1926
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1986 int32
	_ = v1986
	var v1991 int32
	_ = v1991
	var v1995 int32
	_ = v1995
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2006 int32
	_ = v2006
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2027 int32
	_ = v2027
	var v2035 int32
	_ = v2035
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2052 int32
	_ = v2052
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2067 int32
	_ = v2067
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2084 int32
	_ = v2084
	var v2088 int32
	_ = v2088
	var v2096 int32
	_ = v2096
	var v2101 int32
	_ = v2101
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2115 int32
	_ = v2115
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2145 int32
	_ = v2145
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2194 int32
	_ = v2194
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2236 int32
	_ = v2236
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2255 int32
	_ = v2255
	var v2266 int32
	_ = v2266
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2291 int32
	_ = v2291
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2318 int32
	_ = v2318
	var v2327 int32
	_ = v2327
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2352 int32
	_ = v2352
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2369 int32
	_ = v2369
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2403 int32
	_ = v2403
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2417 int32
	_ = v2417
	var v2419 int32
	_ = v2419
	var v2423 int32
	_ = v2423
	var v2429 int32
	_ = v2429
	var v2432 int32
	_ = v2432
	var v2438 int32
	_ = v2438
	var v2442 int32
	_ = v2442
	var v2444 int32
	_ = v2444
	var v2452 int32
	_ = v2452
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2507 int64
	_ = v2507
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2535 int32
	_ = v2535
	var v2540 int32
	_ = v2540
	var v2544 int32
	_ = v2544
	var v2549 int32
	_ = v2549
	var v2551 int32
	_ = v2551
	var v2555 int32
	_ = v2555
	var v2561 int32
	_ = v2561
	var v2564 int32
	_ = v2564
	var v2570 int32
	_ = v2570
	var v2574 int32
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2646 int32
	_ = v2646
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2662 int32
	_ = v2662
	var v2697 int32
	_ = v2697
	var v2705 int32
	_ = v2705
	var v2709 int32
	_ = v2709
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2719 int32
	_ = v2719
	var v2723 int32
	_ = v2723
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2730 int32
	_ = v2730
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2743 int32
	_ = v2743
	var v2747 int32
	_ = v2747
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2766 int32
	_ = v2766
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2785 int32
	_ = v2785
	var v2790 int32
	_ = v2790
	var v2794 int32
	_ = v2794
	var v2799 int32
	_ = v2799
	var v2801 int32
	_ = v2801
	var v2805 int32
	_ = v2805
	var v2811 int32
	_ = v2811
	var v2814 int32
	_ = v2814
	var v2820 int32
	_ = v2820
	var v2824 int32
	_ = v2824
	var v2826 int32
	_ = v2826
	var v2834 int32
	_ = v2834
	var v2847 int32
	_ = v2847
	var v2885 int32
	_ = v2885
	var v2893 int32
	_ = v2893
	var v2903 int32
	_ = v2903
	var v2941 int32
	_ = v2941
	var v2951 int32
	_ = v2951
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2991 int32
	_ = v2991
	var v3000 int32
	_ = v3000
	var v3002 int32
	_ = v3002
	var v3011 int32
	_ = v3011
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3026 int32
	_ = v3026
	var v3033 int32
	_ = v3033
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3040 int32
	_ = v3040
	var v3057 int32
	_ = v3057
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3080 int32
	_ = v3080
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3097 int32
	_ = v3097
	var v3103 int32
	_ = v3103
	var v3109 int32
	_ = v3109
	var v3113 int32
	_ = v3113
	var v3117 int32
	_ = v3117
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3131 int32
	_ = v3131
	var v3135 int32
	_ = v3135
	var v3141 int32
	_ = v3141
	var v3156 int32
	_ = v3156
	var v3160 int32
	_ = v3160
	var v3164 int32
	_ = v3164
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3177 int32
	_ = v3177
	var v3182 int32
	_ = v3182
	var v3185 int32
	_ = v3185
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3193 int32
	_ = v3193
	var v3194 int32
	_ = v3194
	var v3195 int32
	_ = v3195
	var v3197 int32
	_ = v3197
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3211 int32
	_ = v3211
	var v3216 int32
	_ = v3216
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3227 int32
	_ = v3227
	var v3231 int32
	_ = v3231
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3241 int32
	_ = v3241
	var v3245 int32
	_ = v3245
	var v3247 int32
	_ = v3247
	var v3249 int32
	_ = v3249
	var v3252 int32
	_ = v3252
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3261 int32
	_ = v3261
	var v3262 int32
	_ = v3262
	var v3263 int32
	_ = v3263
	var v3265 int32
	_ = v3265
	var v3269 int32
	_ = v3269
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3283 int32
	_ = v3283
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3288 int32
	_ = v3288
	var v3299 int64
	_ = v3299
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3312 int64
	_ = v3312
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3329 int32
	_ = v3329
	var v3333 int32
	_ = v3333
	var v3334 int32
	_ = v3334
	var v3340 int32
	_ = v3340
	var v3344 int32
	_ = v3344
	var v3346 int32
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3354 int32
	_ = v3354
	var v3358 int32
	_ = v3358
	var v3360 int32
	_ = v3360
	var v3362 int32
	_ = v3362
	var v3365 int32
	_ = v3365
	var v3370 int32
	_ = v3370
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3376 int32
	_ = v3376
	var v3378 int32
	_ = v3378
	var v3382 int32
	_ = v3382
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3396 int32
	_ = v3396
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3412 int64
	_ = v3412
	var v3421 int32
	_ = v3421
	var v3422 int32
	_ = v3422
	var v3423 int32
	_ = v3423
	var v3425 int64
	_ = v3425
	var v3434 int32
	_ = v3434
	var v3435 int32
	_ = v3435
	var v3439 int32
	_ = v3439
	var v3441 int32
	_ = v3441
	var v3444 int32
	_ = v3444
	var v3446 int32
	_ = v3446
	var v3449 int32
	_ = v3449
	var v3461 int32
	_ = v3461
	var v3462 int32
	_ = v3462
	var v3466 int32
	_ = v3466
	var v3468 int32
	_ = v3468
	var v3471 int32
	_ = v3471
	var v3473 int32
	_ = v3473
	var v3476 int32
	_ = v3476
	var v3488 int32
	_ = v3488
	var v3489 int32
	_ = v3489
	var v3492 int32
	_ = v3492
	var v3495 int32
	_ = v3495
	var v3497 int32
	_ = v3497
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3521 int32
	_ = v3521
	var v3526 int32
	_ = v3526
	var v3536 int32
	_ = v3536
	var v3537 int32
	_ = v3537
	var v3544 int32
	_ = v3544
	var v3569 int32
	_ = v3569
	var v3574 int32
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3577 int32
	_ = v3577
	var v3579 int32
	_ = v3579
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3592 int32
	_ = v3592
	var v3608 int32
	_ = v3608
	var v3609 int32
	_ = v3609
	var v3620 int32
	_ = v3620
	var v3625 int32
	_ = v3625
	var v3635 int32
	_ = v3635
	var v3660 int32
	_ = v3660
	var v3675 int32
	_ = v3675
	var v3685 int32
	_ = v3685
	var v3709 int32
	_ = v3709
	var v3710 int32
	_ = v3710
	var v3713 int32
	_ = v3713
	var v3715 int32
	_ = v3715
	var v3816 int32
	_ = v3816
	var v3819 int32
	_ = v3819
	var v3823 int32
	_ = v3823
	var v3828 int32
	_ = v3828
	var v3832 int32
	_ = v3832
	var v3835 int32
	_ = v3835
	var v3839 int32
	_ = v3839
	var v3844 int32
	_ = v3844
	v9 = int32(0)
	v47 = m.G0
	v49 = v47 - int32(112)
	m.G0 = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v51 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v51 - int32(1)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v55&int32(16384) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3832 = m.ExcPending
	if v3832 != 0 {
		goto L88
	} else {
		goto L1071
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3816 = m.ExcPending
	if v3816 != 0 {
		goto L88
	} else {
		goto L1067
	}
L6:
	;
	m.G0 = v49 + int32(112)
	return
L7:
	;
	if l7 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if l7 != 0 {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	if (l3^l2)&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L6
L12:
	;
	goto L11
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v113)
	if v113&int32(255) == int32(0) {
		goto L12
	} else {
		goto L28
	}
L14:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v112 = l3
	v113 = v65
	v114 = l2
	goto L13
L15:
	;
	goto L16
L16:
	;
	if l3&int32(3) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v69 = l3
	v71 = l2
	goto L20
L18:
	;
	v83 = l3
	v85 = l2
	goto L19
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 != v90 {
		v112 = v83
		v113 = v87
		v114 = v85
		goto L13
	} else {
		goto L24
	}
L20:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v72)
	if v72 == int32(0) {
		goto L12
	} else {
		goto L22
	}
L21:
	;
	v83 = v79
	v85 = v77
	goto L19
L22:
	;
	v76 = int32(1)
	v77 = v71 + v76
	v79 = v69 + v76
	if v79&int32(3) != 0 {
		v69 = v79
		v71 = v77
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v95 = v83
	v96 = v87
	v97 = v85
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v96
	v99 = int32(4)
	v100 = v97 + v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v103 = v95 + v99
	v107 = int32(-2139062144)
	if (v101|(int32(16843008)-v101))&v107 == v107 {
		v95 = v103
		v96 = v101
		v97 = v100
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v112 = v103
	v113 = v101
	v114 = v100
	goto L13
L27:
	;
	goto L26
L28:
	;
	v121 = v112
	v123 = v114
	goto L29
L29:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)) = uint8(v124)
	v126 = int32(1)
	if v124 != 0 {
		v121 = v121 + v126
		v123 = v123 + v126
		goto L29
	} else {
		goto L31
	}
L30:
	;
	goto L12
L31:
	;
	goto L30
L32:
	;
	v453 = int32(722216)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v454 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L33:
	;
	v135 = v55 & int32(768)
	if v135 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v399 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v399)
	v401 = int32(0)
	v411 = v401
	v412 = v401
	v421 = v9
	v422 = v9
	v452 = v397 + v398 - int32(1)
	goto L32
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v169 = int32(34)
	if v164&v169 != v169 {
		v362 = v9
		goto L46
	} else {
		goto L47
	}
L37:
	;
	if v135 == int32(512) {
		v164 = v55
		v165 = v9
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if l6 == int32(45) {
		v148 = v55
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v164 = v55
	v165 = int32(1)
	goto L36
L41:
	;
	v155 = base.B2i32(l6 == int32(43)) & base.B2i32(v148&int32(96) == int32(32))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v156 != int32(-1) {
		v164 = v148
		v165 = v155
		goto L36
	} else {
		goto L44
	}
L42:
	;
	if v55&int32(32) == int32(0) {
		v148 = v55
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v146 = v55 & int32(-17281)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v146
	v148 = v146
	goto L41
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v159 != v160 {
		v164 = v148
		v165 = v155
		goto L36
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	v164 = v148
	v165 = v155
	goto L36
L46:
	;
	v411 = l5
	v412 = l6
	v421 = v362
	v422 = v165
	v452 = v166 + v167 - base.B2i32(l5|v165 != int32(0))
	goto L32
L47:
	;
	v173 = int32(46)
	v174 = F___strchrnul(m, l3, v173)
	mBase = m.M
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v176 == v173 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v180 == int32(0) {
		v362 = v9
		goto L46
	} else {
		goto L52
	}
L49:
	;
	v180 = v174
	goto L51
L50:
	;
	v180 = int32(0)
	goto L51
L51:
	;
	goto L48
L52:
	;
	v191 = v180
	goto L53
L53:
	;
	v237 = v191
	goto L55
L54:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v280 <= l5 {
		v362 = v191
		goto L46
	} else {
		goto L59
	}
L55:
	;
	v276 = v237 + int32(1)
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if v277 == int32(48) {
		v237 = v276
		goto L55
	} else {
		goto L57
	}
L56:
	;
	if v277 != 0 {
		v191 = v276
		goto L53
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	goto L54
L59:
	;
	if l3&int32(3) == int32(0) {
		v305 = l3
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v340 = v338 - int32(1)
	v341 = v280 - l5
	if v340 < v341 {
		goto L77
	} else {
		goto L78
	}
L61:
	;
	v338 = v330 - l3
	goto L60
L62:
	;
	v309 = v305
	goto L71
L63:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v289 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v338 = int32(0)
	goto L60
L65:
	;
	goto L66
L66:
	;
	v294 = l3
	goto L67
L67:
	;
	v298 = v294 + int32(1)
	if v298&int32(3) == int32(0) {
		v305 = v298
		goto L62
	} else {
		goto L69
	}
L68:
	;
	v330 = v298
	goto L61
L69:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v303 != 0 {
		v294 = v298
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v318 = int32(-2139062144)
	if (int32(16843008)-v315|v315)&v318 == v318 {
		v309 = v309 + int32(4)
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v324 = v309
	goto L74
L73:
	;
	goto L72
L74:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	if v328 != 0 {
		v324 = v324 + int32(1)
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v330 = v324
	goto L61
L76:
	;
	goto L75
L77:
	;
	v343 = v340
	goto L79
L78:
	;
	v343 = v341
	goto L79
L79:
	;
	v344 = l3 + v343
	if base.Ui32(v191) < base.Ui32(v344) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v346 = v344
	goto L82
L81:
	;
	v346 = v191
	goto L82
L82:
	;
	v362 = v346
	goto L46
L83:
	;
	v504 = int32(1)
	v506 = l3 + (l7 ^ v504)
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v507 == v504 {
		goto L120
	} else {
		goto L121
	}
L84:
	;
	v499 = int32(646195)
	v500 = int32(646197)
	v501 = int32(646185)
	v502 = int32(646159)
	v503 = v453
	goto L83
L85:
	;
	goto L86
L86:
	;
	v461 = F_PGLC_localeconv(m)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v461)+32))
	if v467 != 0 {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	return
L89:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v461)+36))
	if v463 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	if v464 != 0 {
		v466 = v463
		goto L87
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v466 = int32(646185)
	goto L87
L93:
	;
	goto L92
L94:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	if v471 != 0 {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
	if v468 != 0 {
		v470 = v467
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v470 = int32(646197)
	goto L94
L98:
	;
	goto L97
L99:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v476&int32(4) != 0 {
		goto L104
	} else {
		goto L105
	}
L100:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	if v472 != 0 {
		v474 = v471
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v474 = int32(646159)
	goto L99
L103:
	;
	goto L102
L104:
	;
	v479 = v474
	goto L106
L105:
	;
	v479 = int32(646159)
	goto L106
L106:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v480 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v461)+16))
	if v491 == int32(0) {
		v499 = v490
		v500 = v470
		v501 = v466
		v502 = v479
		v503 = v453
		goto L83
	} else {
		goto L116
	}
L108:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v481 != 0 {
		v490 = v480
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v483 != int32(44) {
		v490 = int32(646195)
		goto L107
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+1)))
	if v488 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v489 = int32(646195)
	goto L115
L114:
	;
	v489 = int32(646159)
	goto L115
L115:
	;
	v490 = v489
	goto L107
L116:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491))))
	if v494 == int32(0) {
		v499 = v490
		v500 = v470
		v501 = v466
		v502 = v479
		v503 = v453
		goto L83
	} else {
		goto L117
	}
L117:
	;
	v499 = v490
	v500 = v470
	v501 = v466
	v502 = v479
	v503 = v491
	goto L83
L118:
	;
	v3709 = v3675 - int32(1)
	v3710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3709))))
	if v3710 == int32(46) {
		goto L1064
	} else {
		goto L1065
	}
L119:
	;
	if l7 == int32(0) {
		v3675 = v3625
		v3685 = v3635
		goto L118
	} else {
		goto L1062
	}
L120:
	;
	v3620 = l2
	v3625 = v506
	v3635 = v9
	goto L119
L121:
	;
	goto L122
L122:
	;
	v511 = base.B2i32(v412 == int32(45))
	if v412 == int32(45) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v512 = v501
	goto L125
L124:
	;
	v512 = v500
	goto L125
L125:
	;
	v516 = base.B2i32(v412 == int32(43))
	if v412 == int32(43) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v517 = int32(32)
	goto L128
L127:
	;
	v517 = int32(62)
	goto L128
L128:
	;
	if v412 == int32(43) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v520 = int32(32)
	goto L131
L130:
	;
	v520 = int32(60)
	goto L131
L131:
	;
	v521 = l2 + l4
	v527 = l0
	v535 = l2
	v536 = v507
	v540 = v506
	v541 = v9
	v543 = v422
	v545 = v9
	v550 = v9
	v551 = v9
	v558 = v9
	goto L132
L132:
	;
	if l7 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	v3620 = v3569
	v3625 = v3574
	v3635 = v3584
	goto L119
L134:
	;
	v3608 = v527 + int32(12)
	v3609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3608))))
	if v3609 != int32(1) {
		v527 = v3608
		v535 = v3569
		v536 = v3609
		v540 = v3574
		v541 = v3575
		v543 = v3577
		v545 = v3579
		v550 = v3584
		v551 = v3585
		v558 = v3592
		goto L132
	} else {
		goto L1061
	}
L135:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v527)+8))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+8))
	switch v722 {
	case 0:
		goto L192
	case 1, 2, 3, 6:
		goto L193
	default:
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	case 9:
		goto L191
	case 10:
		goto L190
	case 11:
		goto L186
	case 12:
		goto L185
	case 14, 30:
		goto L189
	case 15:
		goto L184
	case 18:
		goto L187
	case 34:
		goto L188
	}
L136:
	;
	v718 = F_pg_mblen_range(m, v535, v521)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L88
	} else {
		goto L181
	}
L137:
	;
	if base.Ui32(v521) <= base.Ui32(v535) {
		v3675 = v540
		v3685 = v550
		goto L118
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	if v536&int32(255) == int32(2) {
		goto L135
	} else {
		goto L142
	}
L140:
	;
	if v536&int32(255) != int32(2) {
		goto L136
	} else {
		goto L141
	}
L141:
	;
	goto L135
L142:
	;
	v585 = v527 + int32(1)
	if (v585^v535)&int32(3) != 0 {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	if v535&int32(3) == int32(0) {
		v683 = v535
		goto L166
	} else {
		goto L167
	}
L144:
	;
	goto L143
L145:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v640))) = uint8(v639)
	if v639&int32(255) == int32(0) {
		goto L144
	} else {
		goto L160
	}
L146:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
	v638 = v585
	v639 = v591
	v640 = v535
	goto L145
L147:
	;
	goto L148
L148:
	;
	if v585&int32(3) != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v595 = v585
	v597 = v535
	goto L152
L150:
	;
	v609 = v585
	v611 = v535
	goto L151
L151:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	v616 = int32(-2139062144)
	if (int32(16843008)-v613|v613)&v616 != v616 {
		v638 = v609
		v639 = v613
		v640 = v611
		goto L145
	} else {
		goto L156
	}
L152:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	*(*uint8)(unsafe.Add(mBase, uint32(v597))) = uint8(v598)
	if v598 == int32(0) {
		goto L144
	} else {
		goto L154
	}
L153:
	;
	v609 = v605
	v611 = v603
	goto L151
L154:
	;
	v602 = int32(1)
	v603 = v597 + v602
	v605 = v595 + v602
	if v605&int32(3) != 0 {
		v595 = v605
		v597 = v603
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v621 = v609
	v622 = v613
	v623 = v611
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = v622
	v625 = int32(4)
	v626 = v623 + v625
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v621)+4))
	v629 = v621 + v625
	v633 = int32(-2139062144)
	if (v627|(int32(16843008)-v627))&v633 == v633 {
		v621 = v629
		v622 = v627
		v623 = v626
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v638 = v629
	v639 = v627
	v640 = v626
	goto L145
L159:
	;
	goto L158
L160:
	;
	v647 = v638
	v649 = v640
	goto L161
L161:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+1)) = uint8(v650)
	v652 = int32(1)
	if v650 != 0 {
		v647 = v647 + v652
		v649 = v649 + v652
		goto L161
	} else {
		goto L163
	}
L162:
	;
	goto L144
L163:
	;
	goto L162
L164:
	;
	v3569 = v716 + v535
	v3574 = v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L165:
	;
	v716 = v708 - v535
	goto L164
L166:
	;
	v687 = v683
	goto L175
L167:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v667 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v716 = int32(0)
	goto L164
L169:
	;
	goto L170
L170:
	;
	v672 = v535
	goto L171
L171:
	;
	v676 = v672 + int32(1)
	if v676&int32(3) == int32(0) {
		v683 = v676
		goto L166
	} else {
		goto L173
	}
L172:
	;
	v708 = v676
	goto L165
L173:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	if v681 != 0 {
		v672 = v676
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v687)))
	v696 = int32(-2139062144)
	if (int32(16843008)-v693|v693)&v696 == v696 {
		v687 = v687 + int32(4)
		goto L175
	} else {
		goto L177
	}
L176:
	;
	v702 = v687
	goto L178
L177:
	;
	goto L176
L178:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702))))
	if v706 != 0 {
		v702 = v702 + int32(1)
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v708 = v702
	goto L165
L180:
	;
	goto L179
L181:
	;
	v3569 = v718 + v535
	v3574 = v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L182:
	;
	v3569 = v3521 + int32(1)
	v3574 = v3526
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v3536
	v3585 = v3537
	v3592 = v3544
	goto L134
L183:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2096)
	v3521 = v1837
	v3526 = v1839
	v3536 = v1840
	v3537 = v1841
	v3544 = v1842
	goto L182
L184:
	;
	if l7 != 0 {
		goto L1049
	} else {
		goto L1050
	}
L185:
	;
	if l7 != 0 {
		goto L1033
	} else {
		goto L1034
	}
L186:
	;
	if l7 != 0 {
		goto L1017
	} else {
		goto L1018
	}
L187:
	;
	v3324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v3324&int32(1024) != 0 {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L976
	}
L188:
	;
	v3211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v3211&int32(1024) != 0 {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L935
	}
L189:
	;
	if l7 != 0 {
		goto L778
	} else {
		goto L779
	}
L190:
	;
	if l7 != 0 {
		goto L726
	} else {
		goto L727
	}
L191:
	;
	if v499&int32(3) == int32(0) {
		v2141 = v499
		goto L660
	} else {
		goto L661
	}
L192:
	;
	if l7 != 0 {
		goto L645
	} else {
		goto L646
	}
L193:
	;
	if l7 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v723&int32(1024) != 0 {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	if base.Ui32(v521) <= base.Ui32(v535) {
		v3521 = v535
		v3526 = v540
		v3536 = v550
		v3537 = v551
		v3544 = v558
		goto L182
	} else {
		goto L430
	}
L197:
	;
	if v543 != 0 {
		v902 = v535
		v904 = v543
		goto L198
	} else {
		goto L199
	}
L198:
	;
	if int32(1)<<(uint(v722)%32)&int32(78) != 0 {
		goto L265
	} else {
		goto L266
	}
L199:
	;
	v727 = v723 & int32(8)
	if v545 < v411 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	if v723&int32(64) != 0 {
		goto L212
	} else {
		goto L213
	}
L201:
	;
	v729 = int32(0)
	if v727 == v729 {
		v902 = v535
		v904 = v729
		goto L198
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	if v727 != 0 {
		goto L200
	} else {
		goto L206
	}
L204:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v732 == v545 {
		goto L200
	} else {
		goto L205
	}
L205:
	;
	v902 = v535
	v904 = v729
	goto L198
L206:
	;
	if l3 != v540 {
		goto L200
	} else {
		goto L207
	}
L207:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v735 != int32(48) {
		goto L200
	} else {
		goto L208
	}
L208:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v738 == int32(0) {
		goto L200
	} else {
		goto L209
	}
L209:
	;
	v741 = int32(0)
	if v421 == v741 {
		v902 = v535
		v904 = v741
		goto L198
	} else {
		goto L210
	}
L210:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	if v744 != int32(46) {
		v902 = v535
		v904 = v741
		goto L198
	} else {
		goto L211
	}
L211:
	;
	goto L200
L212:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v750 != int32(-1) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	if v723&int32(128) != 0 {
		goto L257
	} else {
		goto L258
	}
L215:
	;
	v902 = v535
	v904 = int32(0)
	goto L198
L216:
	;
	goto L217
L217:
	;
	if (v512^v535)&int32(3) != 0 {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	if v535&int32(3) == int32(0) {
		v851 = v535
		goto L241
	} else {
		goto L242
	}
L219:
	;
	goto L218
L220:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v808))) = uint8(v807)
	if v807&int32(255) == int32(0) {
		goto L219
	} else {
		goto L235
	}
L221:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	v806 = v512
	v807 = v759
	v808 = v535
	goto L220
L222:
	;
	goto L223
L223:
	;
	if v512&int32(3) != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v763 = v512
	v765 = v535
	goto L227
L225:
	;
	v777 = v512
	v779 = v535
	goto L226
L226:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v777)))
	v784 = int32(-2139062144)
	if (int32(16843008)-v781|v781)&v784 != v784 {
		v806 = v777
		v807 = v781
		v808 = v779
		goto L220
	} else {
		goto L231
	}
L227:
	;
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763))))
	*(*uint8)(unsafe.Add(mBase, uint32(v765))) = uint8(v766)
	if v766 == int32(0) {
		goto L219
	} else {
		goto L229
	}
L228:
	;
	v777 = v773
	v779 = v771
	goto L226
L229:
	;
	v770 = int32(1)
	v771 = v765 + v770
	v773 = v763 + v770
	if v773&int32(3) != 0 {
		v763 = v773
		v765 = v771
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	v789 = v777
	v790 = v781
	v791 = v779
	goto L232
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v791))) = v790
	v793 = int32(4)
	v794 = v791 + v793
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v789)+4))
	v797 = v789 + v793
	v801 = int32(-2139062144)
	if (v795|(int32(16843008)-v795))&v801 == v801 {
		v789 = v797
		v790 = v795
		v791 = v794
		goto L232
	} else {
		goto L234
	}
L233:
	;
	v806 = v797
	v807 = v795
	v808 = v794
	goto L220
L234:
	;
	goto L233
L235:
	;
	v815 = v806
	v817 = v808
	goto L236
L236:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v817)+1)) = uint8(v818)
	v820 = int32(1)
	if v818 != 0 {
		v815 = v815 + v820
		v817 = v817 + v820
		goto L236
	} else {
		goto L238
	}
L237:
	;
	goto L219
L238:
	;
	goto L237
L239:
	;
	v902 = v884 + v535
	v904 = int32(1)
	goto L198
L240:
	;
	v884 = v876 - v535
	goto L239
L241:
	;
	v855 = v851
	goto L250
L242:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v835 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v884 = int32(0)
	goto L239
L244:
	;
	goto L245
L245:
	;
	v840 = v535
	goto L246
L246:
	;
	v844 = v840 + int32(1)
	if v844&int32(3) == int32(0) {
		v851 = v844
		goto L241
	} else {
		goto L248
	}
L247:
	;
	v876 = v844
	goto L240
L248:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844))))
	if v849 != 0 {
		v840 = v844
		goto L246
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v855)))
	v864 = int32(-2139062144)
	if (int32(16843008)-v861|v861)&v864 == v864 {
		v855 = v855 + int32(4)
		goto L250
	} else {
		goto L252
	}
L251:
	;
	v870 = v855
	goto L253
L252:
	;
	goto L251
L253:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870))))
	if v874 != 0 {
		v870 = v870 + int32(1)
		goto L253
	} else {
		goto L255
	}
L254:
	;
	v876 = v870
	goto L240
L255:
	;
	goto L254
L256:
	;
	v899 = int32(1)
	v902 = v535 + v899
	v904 = v899
	goto L198
L257:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v520)
	goto L256
L258:
	;
	goto L259
L259:
	;
	switch v412 - int32(43) {
	case 0:
		goto L261
	default:
		v902 = v535
		v904 = int32(0)
		goto L198
	case 2:
		goto L260
	}
L260:
	;
	v896 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v896)
	goto L256
L261:
	;
	if v723&int32(32) != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v902 = v535
	v904 = int32(1)
	goto L198
L263:
	;
	goto L264
L264:
	;
	v894 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v894)
	goto L256
L265:
	;
	v912 = base.B2i32(base.Ui32(v722) <= base.Ui32(int32(6)))
	goto L267
L266:
	;
	v912 = int32(0)
	goto L267
L267:
	;
	if v912 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v3569 = v902
	v3574 = v540
	v3575 = int32(0)
	v3577 = v904
	v3579 = v545 + int32(1)
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L269:
	;
	goto L270
L270:
	;
	if v545 < v411 {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1266 = int32(1)
	v1270 = v452 + base.B2i32(v411 != int32(0)) + int32(base.Ui32(v1265)>>(uint(v1266)%32))&v1266
	if v1263 == v421 {
		goto L380
	} else {
		goto L381
	}
L272:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v545 < v920 {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	goto L274
L274:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540))))
	if v938 == int32(46) {
		goto L281
	} else {
		goto L282
	}
L275:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v902))) = uint8(v932)
	v1259 = v902 + int32(1)
	v1263 = v540
	v1264 = v933
	goto L271
L276:
	;
	v928 = int32(32)
	v929 = int32(0)
	if v919&v928 != 0 {
		v1259 = v902
		v1263 = v540
		v1264 = v929
		goto L271
	} else {
		goto L279
	}
L277:
	;
	if v919&int32(8) == int32(0) {
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v932 = int32(48)
	v933 = int32(1)
	goto L275
L279:
	;
	v932 = v928
	v933 = v929
	goto L275
L280:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540))))
	v1259 = v1253
	v1263 = v540 + base.B2i32(v1255 != int32(0))
	v1264 = v1254
	goto L271
L281:
	;
	if v421 != 0 {
		goto L285
	} else {
		goto L286
	}
L282:
	;
	goto L283
L283:
	;
	if v421 == int32(0) {
		goto L366
	} else {
		goto L367
	}
L284:
	;
	v1077 = int32(0)
	if v937&int32(32) == v1077 {
		v1253 = v902
		v1254 = v1077
		goto L280
	} else {
		goto L327
	}
L285:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	if v941 == int32(46) {
		goto L284
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	if (v502^v902)&int32(3) != 0 {
		goto L292
	} else {
		goto L293
	}
L288:
	;
	goto L287
L289:
	;
	if v902&int32(3) == int32(0) {
		v1041 = v902
		goto L312
	} else {
		goto L313
	}
L290:
	;
	goto L289
L291:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v998))) = uint8(v997)
	if v997&int32(255) == int32(0) {
		goto L290
	} else {
		goto L306
	}
L292:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	v996 = v502
	v997 = v949
	v998 = v902
	goto L291
L293:
	;
	goto L294
L294:
	;
	if v502&int32(3) != 0 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v953 = v502
	v955 = v902
	goto L298
L296:
	;
	v967 = v502
	v969 = v902
	goto L297
L297:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v967)))
	v974 = int32(-2139062144)
	if (int32(16843008)-v971|v971)&v974 != v974 {
		v996 = v967
		v997 = v971
		v998 = v969
		goto L291
	} else {
		goto L302
	}
L298:
	;
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v953))))
	*(*uint8)(unsafe.Add(mBase, uint32(v955))) = uint8(v956)
	if v956 == int32(0) {
		goto L290
	} else {
		goto L300
	}
L299:
	;
	v967 = v963
	v969 = v961
	goto L297
L300:
	;
	v960 = int32(1)
	v961 = v955 + v960
	v963 = v953 + v960
	if v963&int32(3) != 0 {
		v953 = v963
		v955 = v961
		goto L298
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	v979 = v967
	v980 = v971
	v981 = v969
	goto L303
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v981))) = v980
	v983 = int32(4)
	v984 = v981 + v983
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v979)+4))
	v987 = v979 + v983
	v991 = int32(-2139062144)
	if (v985|(int32(16843008)-v985))&v991 == v991 {
		v979 = v987
		v980 = v985
		v981 = v984
		goto L303
	} else {
		goto L305
	}
L304:
	;
	v996 = v987
	v997 = v985
	v998 = v984
	goto L291
L305:
	;
	goto L304
L306:
	;
	v1005 = v996
	v1007 = v998
	goto L307
L307:
	;
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1007)+1)) = uint8(v1008)
	v1010 = int32(1)
	if v1008 != 0 {
		v1005 = v1005 + v1010
		v1007 = v1007 + v1010
		goto L307
	} else {
		goto L309
	}
L308:
	;
	goto L290
L309:
	;
	goto L308
L310:
	;
	v1253 = v1074 + v902
	v1254 = int32(0)
	goto L280
L311:
	;
	v1074 = v1066 - v902
	goto L310
L312:
	;
	v1045 = v1041
	goto L321
L313:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	if v1025 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1074 = int32(0)
	goto L310
L315:
	;
	goto L316
L316:
	;
	v1030 = v902
	goto L317
L317:
	;
	v1034 = v1030 + int32(1)
	if v1034&int32(3) == int32(0) {
		v1041 = v1034
		goto L312
	} else {
		goto L319
	}
L318:
	;
	v1066 = v1034
	goto L311
L319:
	;
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	if v1039 != 0 {
		v1030 = v1034
		goto L317
	} else {
		goto L320
	}
L320:
	;
	goto L318
L321:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1045)))
	v1054 = int32(-2139062144)
	if (int32(16843008)-v1051|v1051)&v1054 == v1054 {
		v1045 = v1045 + int32(4)
		goto L321
	} else {
		goto L323
	}
L322:
	;
	v1060 = v1045
	goto L324
L323:
	;
	goto L322
L324:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060))))
	if v1064 != 0 {
		v1060 = v1060 + int32(1)
		goto L324
	} else {
		goto L326
	}
L325:
	;
	v1066 = v1060
	goto L311
L326:
	;
	goto L325
L327:
	;
	if (v502^v902)&int32(3) != 0 {
		goto L331
	} else {
		goto L332
	}
L328:
	;
	if v902&int32(3) == int32(0) {
		v1179 = v902
		goto L351
	} else {
		goto L352
	}
L329:
	;
	goto L328
L330:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1136))) = uint8(v1135)
	if v1135&int32(255) == int32(0) {
		goto L329
	} else {
		goto L345
	}
L331:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	v1134 = v502
	v1135 = v1087
	v1136 = v902
	goto L330
L332:
	;
	goto L333
L333:
	;
	if v502&int32(3) != 0 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1091 = v502
	v1093 = v902
	goto L337
L335:
	;
	v1105 = v502
	v1107 = v902
	goto L336
L336:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1105)))
	v1112 = int32(-2139062144)
	if (int32(16843008)-v1109|v1109)&v1112 != v1112 {
		v1134 = v1105
		v1135 = v1109
		v1136 = v1107
		goto L330
	} else {
		goto L341
	}
L337:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1091))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1093))) = uint8(v1094)
	if v1094 == int32(0) {
		goto L329
	} else {
		goto L339
	}
L338:
	;
	v1105 = v1101
	v1107 = v1099
	goto L336
L339:
	;
	v1098 = int32(1)
	v1099 = v1093 + v1098
	v1101 = v1091 + v1098
	if v1101&int32(3) != 0 {
		v1091 = v1101
		v1093 = v1099
		goto L337
	} else {
		goto L340
	}
L340:
	;
	goto L338
L341:
	;
	v1117 = v1105
	v1118 = v1109
	v1119 = v1107
	goto L342
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1119))) = v1118
	v1121 = int32(4)
	v1122 = v1119 + v1121
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+4))
	v1125 = v1117 + v1121
	v1129 = int32(-2139062144)
	if (v1123|(int32(16843008)-v1123))&v1129 == v1129 {
		v1117 = v1125
		v1118 = v1123
		v1119 = v1122
		goto L342
	} else {
		goto L344
	}
L343:
	;
	v1134 = v1125
	v1135 = v1123
	v1136 = v1122
	goto L330
L344:
	;
	goto L343
L345:
	;
	v1143 = v1134
	v1145 = v1136
	goto L346
L346:
	;
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1143)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1145)+1)) = uint8(v1146)
	v1148 = int32(1)
	if v1146 != 0 {
		v1143 = v1143 + v1148
		v1145 = v1145 + v1148
		goto L346
	} else {
		goto L348
	}
L347:
	;
	goto L329
L348:
	;
	goto L347
L349:
	;
	v1253 = v1212 + v902
	v1254 = v1077
	goto L280
L350:
	;
	v1212 = v1204 - v902
	goto L349
L351:
	;
	v1183 = v1179
	goto L360
L352:
	;
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	if v1163 == int32(0) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1212 = int32(0)
	goto L349
L354:
	;
	goto L355
L355:
	;
	v1168 = v902
	goto L356
L356:
	;
	v1172 = v1168 + int32(1)
	if v1172&int32(3) == int32(0) {
		v1179 = v1172
		goto L351
	} else {
		goto L358
	}
L357:
	;
	v1204 = v1172
	goto L350
L358:
	;
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1172))))
	if v1177 != 0 {
		v1168 = v1172
		goto L356
	} else {
		goto L359
	}
L359:
	;
	goto L357
L360:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1183)))
	v1192 = int32(-2139062144)
	if (int32(16843008)-v1189|v1189)&v1192 == v1192 {
		v1183 = v1183 + int32(4)
		goto L360
	} else {
		goto L362
	}
L361:
	;
	v1198 = v1183
	goto L363
L362:
	;
	goto L361
L363:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1198))))
	if v1202 != 0 {
		v1198 = v1198 + int32(1)
		goto L363
	} else {
		goto L365
	}
L364:
	;
	v1204 = v1198
	goto L350
L365:
	;
	goto L364
L366:
	;
	if v937&int32(8) != 0 {
		goto L370
	} else {
		goto L371
	}
L367:
	;
	if v722 == int32(2) {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	if base.Ui32(v421) < base.Ui32(v540) {
		v1253 = v902
		v1254 = int32(0)
		goto L280
	} else {
		goto L369
	}
L369:
	;
	goto L366
L370:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v902))) = uint8(v938)
	v1250 = int32(1)
	v1253 = v902 + v1250
	v1254 = v1250
	goto L280
L371:
	;
	if l3 != v540 {
		goto L370
	} else {
		goto L372
	}
L372:
	;
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1224 != int32(48) {
		goto L370
	} else {
		goto L373
	}
L373:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1227 == int32(0) {
		goto L370
	} else {
		goto L374
	}
L374:
	;
	if v937&int32(32) == int32(0) {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1234 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v902))) = uint8(v1234)
	v1253 = v902 + int32(1)
	v1254 = int32(0)
	goto L280
L376:
	;
	goto L377
L377:
	;
	v1239 = int32(0)
	if v421 == v1239 {
		v1253 = v902
		v1254 = v1239
		goto L280
	} else {
		goto L378
	}
L378:
	;
	v1242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	if v1242 != int32(46) {
		v1253 = v902
		v1254 = v1239
		goto L280
	} else {
		goto L379
	}
L379:
	;
	v1245 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v902))) = uint8(v1245)
	v1253 = v902 + int32(1)
	v1254 = v1239
	goto L280
L380:
	;
	v1272 = v545
	goto L382
L381:
	;
	v1272 = v1270
	goto L382
L382:
	;
	v1274 = v545 + int32(1)
	if v421 != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1275 = v1272
	goto L385
L384:
	;
	v1275 = v1270
	goto L385
L385:
	;
	if v1274 != v1275 {
		v3569 = v1259
		v3574 = v1263
		v3575 = v1264
		v3577 = v904
		v3579 = v1274
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L386
	}
L386:
	;
	if v904 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	if v1265&int32(64) == int32(0) {
		v3569 = v1259
		v3574 = v1263
		v3575 = v1264
		v3577 = v904
		v3579 = v1274
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L390
	}
L388:
	;
	if v1265&int32(128) == int32(0) {
		goto L387
	} else {
		goto L389
	}
L389:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1259))) = uint8(v517)
	v1284 = int32(1)
	v3569 = v1259 + v1284
	v3574 = v1263
	v3575 = v1264
	v3577 = v1284
	v3579 = v1274
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L390:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1291 != int32(1) {
		v3569 = v1259
		v3574 = v1263
		v3575 = v1264
		v3577 = v904
		v3579 = v1274
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L391
	}
L391:
	;
	if (v512^v1259)&int32(3) != 0 {
		goto L395
	} else {
		goto L396
	}
L392:
	;
	if v1259&int32(3) == int32(0) {
		v1391 = v1259
		goto L415
	} else {
		goto L416
	}
L393:
	;
	goto L392
L394:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1348))) = uint8(v1347)
	if v1347&int32(255) == int32(0) {
		goto L393
	} else {
		goto L409
	}
L395:
	;
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	v1346 = v512
	v1347 = v1299
	v1348 = v1259
	goto L394
L396:
	;
	goto L397
L397:
	;
	if v512&int32(3) != 0 {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v1303 = v512
	v1305 = v1259
	goto L401
L399:
	;
	v1317 = v512
	v1319 = v1259
	goto L400
L400:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1317)))
	v1324 = int32(-2139062144)
	if (int32(16843008)-v1321|v1321)&v1324 != v1324 {
		v1346 = v1317
		v1347 = v1321
		v1348 = v1319
		goto L394
	} else {
		goto L405
	}
L401:
	;
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1305))) = uint8(v1306)
	if v1306 == int32(0) {
		goto L393
	} else {
		goto L403
	}
L402:
	;
	v1317 = v1313
	v1319 = v1311
	goto L400
L403:
	;
	v1310 = int32(1)
	v1311 = v1305 + v1310
	v1313 = v1303 + v1310
	if v1313&int32(3) != 0 {
		v1303 = v1313
		v1305 = v1311
		goto L401
	} else {
		goto L404
	}
L404:
	;
	goto L402
L405:
	;
	v1329 = v1317
	v1330 = v1321
	v1331 = v1319
	goto L406
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1331))) = v1330
	v1333 = int32(4)
	v1334 = v1331 + v1333
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+4))
	v1337 = v1329 + v1333
	v1341 = int32(-2139062144)
	if (v1335|(int32(16843008)-v1335))&v1341 == v1341 {
		v1329 = v1337
		v1330 = v1335
		v1331 = v1334
		goto L406
	} else {
		goto L408
	}
L407:
	;
	v1346 = v1337
	v1347 = v1335
	v1348 = v1334
	goto L394
L408:
	;
	goto L407
L409:
	;
	v1355 = v1346
	v1357 = v1348
	goto L410
L410:
	;
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1355)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1357)+1)) = uint8(v1358)
	v1360 = int32(1)
	if v1358 != 0 {
		v1355 = v1355 + v1360
		v1357 = v1357 + v1360
		goto L410
	} else {
		goto L412
	}
L411:
	;
	goto L393
L412:
	;
	goto L411
L413:
	;
	v3569 = v1424 + v1259
	v3574 = v1263
	v3575 = v1264
	v3577 = v904
	v3579 = v1274
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L414:
	;
	v1424 = v1416 - v1259
	goto L413
L415:
	;
	v1395 = v1391
	goto L424
L416:
	;
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1259))))
	if v1375 == int32(0) {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1424 = int32(0)
	goto L413
L418:
	;
	goto L419
L419:
	;
	v1380 = v1259
	goto L420
L420:
	;
	v1384 = v1380 + int32(1)
	if v1384&int32(3) == int32(0) {
		v1391 = v1384
		goto L415
	} else {
		goto L422
	}
L421:
	;
	v1416 = v1384
	goto L414
L422:
	;
	v1389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1384))))
	if v1389 != 0 {
		v1380 = v1384
		goto L420
	} else {
		goto L423
	}
L423:
	;
	goto L421
L424:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1395)))
	v1404 = int32(-2139062144)
	if (int32(16843008)-v1401|v1401)&v1404 == v1404 {
		v1395 = v1395 + int32(4)
		goto L424
	} else {
		goto L426
	}
L425:
	;
	v1410 = v1395
	goto L427
L426:
	;
	goto L425
L427:
	;
	v1414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1410))))
	if v1414 != 0 {
		v1410 = v1410 + int32(1)
		goto L427
	} else {
		goto L429
	}
L428:
	;
	v1416 = v1410
	goto L414
L429:
	;
	goto L428
L430:
	;
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	v1430 = v535 + base.B2i32(v1427 == int32(32))
	if base.Ui32(v521) <= base.Ui32(v1430) {
		v3521 = v1430
		v3526 = v540
		v3536 = v550
		v3537 = v551
		v3544 = v558
		goto L182
	} else {
		goto L431
	}
L431:
	;
	if v722&int32(-2) != int32(2) {
		v1685 = v1430
		goto L432
	} else {
		goto L433
	}
L432:
	;
	if base.Ui32(v521) <= base.Ui32(v1685) {
		v3521 = v1685
		v3526 = v540
		v3536 = v550
		v3537 = v551
		v3544 = v558
		goto L182
	} else {
		goto L516
	}
L433:
	;
	v1436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1436 != int32(32) {
		v1685 = v1430
		goto L432
	} else {
		goto L434
	}
L434:
	;
	if v558 != int32(0)-v550 {
		v1685 = v1430
		goto L432
	} else {
		goto L435
	}
L435:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1442&int32(64) == int32(0) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v1668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1430))))
	if v1668 != int32(45) {
		goto L510
	} else {
		goto L511
	}
L437:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1447 != int32(-1) {
		goto L436
	} else {
		goto L438
	}
L438:
	;
	if v501&int32(3) == int32(0) {
		v1473 = v501
		goto L442
	} else {
		goto L443
	}
L439:
	;
	if v500&int32(3) == int32(0) {
		v1582 = v500
		goto L476
	} else {
		goto L477
	}
L440:
	;
	if v1506 == int32(0) {
		goto L439
	} else {
		goto L457
	}
L441:
	;
	v1506 = v1498 - v501
	goto L440
L442:
	;
	v1477 = v1473
	goto L451
L443:
	;
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if v1457 == int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v1506 = int32(0)
	goto L440
L445:
	;
	goto L446
L446:
	;
	v1462 = v501
	goto L447
L447:
	;
	v1466 = v1462 + int32(1)
	if v1466&int32(3) == int32(0) {
		v1473 = v1466
		goto L442
	} else {
		goto L449
	}
L448:
	;
	v1498 = v1466
	goto L441
L449:
	;
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1466))))
	if v1471 != 0 {
		v1462 = v1466
		goto L447
	} else {
		goto L450
	}
L450:
	;
	goto L448
L451:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1477)))
	v1486 = int32(-2139062144)
	if (int32(16843008)-v1483|v1483)&v1486 == v1486 {
		v1477 = v1477 + int32(4)
		goto L451
	} else {
		goto L453
	}
L452:
	;
	v1492 = v1477
	goto L454
L453:
	;
	goto L452
L454:
	;
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492))))
	if v1496 != 0 {
		v1492 = v1492 + int32(1)
		goto L454
	} else {
		goto L456
	}
L455:
	;
	v1498 = v1492
	goto L441
L456:
	;
	goto L455
L457:
	;
	if base.Ui32(l2+(l4-v1506)) < base.Ui32(v1430) {
		goto L439
	} else {
		goto L458
	}
L458:
	;
	if v1506 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	if v1555 != 0 {
		goto L439
	} else {
		goto L473
	}
L460:
	;
	v1555 = int32(0)
	goto L459
L461:
	;
	goto L462
L462:
	;
	v1517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1430))))
	if v1517 != 0 {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v1518 = v1430
	v1519 = v501
	v1520 = v1506
	v1521 = v1517
	goto L467
L464:
	;
	v1543 = v501
	v1547 = int32(0)
	goto L465
L465:
	;
	v1548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1543))))
	v1555 = v1547 - v1548
	goto L459
L466:
	;
	v1543 = v1538
	v1547 = v1540
	goto L465
L467:
	;
	v1523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1519))))
	if v1521 != v1523 {
		v1538 = v1519
		v1540 = v1521
		goto L466
	} else {
		goto L469
	}
L468:
	;
	v1538 = v1532
	v1540 = int32(0)
	goto L466
L469:
	;
	if v1523 == int32(0) {
		v1538 = v1519
		v1540 = v1521
		goto L466
	} else {
		goto L470
	}
L470:
	;
	v1528 = v1520 - int32(1)
	if v1528 == int32(0) {
		v1538 = v1519
		v1540 = v1521
		goto L466
	} else {
		goto L471
	}
L471:
	;
	v1531 = int32(1)
	v1532 = v1519 + v1531
	v1533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+1)))
	if v1533 != 0 {
		v1518 = v1518 + v1531
		v1519 = v1532
		v1520 = v1528
		v1521 = v1533
		goto L467
	} else {
		goto L472
	}
L472:
	;
	goto L468
L473:
	;
	v1556 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1556)
	v1685 = v1430 + v1506
	goto L432
L474:
	;
	if v1615 == int32(0) {
		v1685 = v1430
		goto L432
	} else {
		goto L491
	}
L475:
	;
	v1615 = v1607 - v500
	goto L474
L476:
	;
	v1586 = v1582
	goto L485
L477:
	;
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	if v1566 == int32(0) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v1615 = int32(0)
	goto L474
L479:
	;
	goto L480
L480:
	;
	v1571 = v500
	goto L481
L481:
	;
	v1575 = v1571 + int32(1)
	if v1575&int32(3) == int32(0) {
		v1582 = v1575
		goto L476
	} else {
		goto L483
	}
L482:
	;
	v1607 = v1575
	goto L475
L483:
	;
	v1580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1575))))
	if v1580 != 0 {
		v1571 = v1575
		goto L481
	} else {
		goto L484
	}
L484:
	;
	goto L482
L485:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1586)))
	v1595 = int32(-2139062144)
	if (int32(16843008)-v1592|v1592)&v1595 == v1595 {
		v1586 = v1586 + int32(4)
		goto L485
	} else {
		goto L487
	}
L486:
	;
	v1601 = v1586
	goto L488
L487:
	;
	goto L486
L488:
	;
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1601))))
	if v1605 != 0 {
		v1601 = v1601 + int32(1)
		goto L488
	} else {
		goto L490
	}
L489:
	;
	v1607 = v1601
	goto L475
L490:
	;
	goto L489
L491:
	;
	if base.Ui32(l2+(l4-v1615)) < base.Ui32(v1430) {
		v1685 = v1430
		goto L432
	} else {
		goto L492
	}
L492:
	;
	if v1615 == int32(0) {
		goto L494
	} else {
		goto L495
	}
L493:
	;
	if v1664 != 0 {
		v1685 = v1430
		goto L432
	} else {
		goto L507
	}
L494:
	;
	v1664 = int32(0)
	goto L493
L495:
	;
	goto L496
L496:
	;
	v1626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1430))))
	if v1626 != 0 {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v1627 = v1430
	v1628 = v500
	v1629 = v1615
	v1630 = v1626
	goto L501
L498:
	;
	v1652 = v500
	v1656 = int32(0)
	goto L499
L499:
	;
	v1657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1652))))
	v1664 = v1656 - v1657
	goto L493
L500:
	;
	v1652 = v1647
	v1656 = v1649
	goto L499
L501:
	;
	v1632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1628))))
	if v1630 != v1632 {
		v1647 = v1628
		v1649 = v1630
		goto L500
	} else {
		goto L503
	}
L502:
	;
	v1647 = v1641
	v1649 = int32(0)
	goto L500
L503:
	;
	if v1632 == int32(0) {
		v1647 = v1628
		v1649 = v1630
		goto L500
	} else {
		goto L504
	}
L504:
	;
	v1637 = v1629 - int32(1)
	if v1637 == int32(0) {
		v1647 = v1628
		v1649 = v1630
		goto L500
	} else {
		goto L505
	}
L505:
	;
	v1640 = int32(1)
	v1641 = v1628 + v1640
	v1642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1627)+1)))
	if v1642 != 0 {
		v1627 = v1627 + v1640
		v1628 = v1641
		v1629 = v1637
		v1630 = v1642
		goto L501
	} else {
		goto L506
	}
L506:
	;
	goto L502
L507:
	;
	v1665 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1665)
	v1685 = v1430 + v1615
	goto L432
L508:
	;
	v1685 = v1430 + int32(1)
	goto L432
L509:
	;
	if v1668 != int32(43) {
		v1685 = v1430
		goto L432
	} else {
		goto L515
	}
L510:
	;
	if v1442&int32(128) == int32(0) {
		goto L509
	} else {
		goto L513
	}
L511:
	;
	goto L512
L512:
	;
	v1677 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1677)
	goto L508
L513:
	;
	if v1668 != int32(60) {
		goto L509
	} else {
		goto L514
	}
L514:
	;
	goto L512
L515:
	;
	v1681 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1681)
	goto L508
L516:
	;
	v1689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1685))))
	if base.Ui32((v1689-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L518
	} else {
		goto L519
	}
L517:
	;
	if base.Ui32(v521) <= base.Ui32(v1837) {
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	} else {
		goto L567
	}
L518:
	;
	if v551 != 0 {
		goto L521
	} else {
		goto L522
	}
L519:
	;
	goto L520
L520:
	;
	v1711 = int32(0)
	v1712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1712&int32(2) == v1711 {
		v1837 = v1685
		v1838 = v1711
		v1839 = v540
		v1840 = v550
		v1841 = v551
		v1842 = v558
		goto L517
	} else {
		goto L525
	}
L521:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v550 == v1696 {
		v3521 = v1685
		v3526 = v540
		v3536 = v550
		v3537 = v551
		v3544 = v558
		goto L182
	} else {
		goto L524
	}
L522:
	;
	goto L523
L523:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v540))) = uint8(v1689)
	v1705 = int32(1)
	v1837 = v1685
	v1838 = v1705
	v1839 = v540 + v1705
	v1840 = v550
	v1841 = int32(0)
	v1842 = v558 + v1705
	goto L517
L524:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v540))) = uint8(v1689)
	v1699 = int32(1)
	v1837 = v1685
	v1838 = v1699
	v1839 = v540 + v1699
	v1840 = v550 + v1699
	v1841 = v551
	v1842 = v558
	goto L517
L525:
	;
	if v551 != 0 {
		v1837 = v1685
		v1838 = v1711
		v1839 = v540
		v1840 = v550
		v1841 = v551
		v1842 = v558
		goto L517
	} else {
		goto L526
	}
L526:
	;
	if v502&int32(3) == int32(0) {
		v1740 = v502
		goto L529
	} else {
		goto L530
	}
L527:
	;
	if v1773 == int32(0) {
		goto L544
	} else {
		goto L545
	}
L528:
	;
	v1773 = v1765 - v502
	goto L527
L529:
	;
	v1744 = v1740
	goto L538
L530:
	;
	v1724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	if v1724 == int32(0) {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	v1773 = int32(0)
	goto L527
L532:
	;
	goto L533
L533:
	;
	v1729 = v502
	goto L534
L534:
	;
	v1733 = v1729 + int32(1)
	if v1733&int32(3) == int32(0) {
		v1740 = v1733
		goto L529
	} else {
		goto L536
	}
L535:
	;
	v1765 = v1733
	goto L528
L536:
	;
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1733))))
	if v1738 != 0 {
		v1729 = v1733
		goto L534
	} else {
		goto L537
	}
L537:
	;
	goto L535
L538:
	;
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1744)))
	v1753 = int32(-2139062144)
	if (int32(16843008)-v1750|v1750)&v1753 == v1753 {
		v1744 = v1744 + int32(4)
		goto L538
	} else {
		goto L540
	}
L539:
	;
	v1759 = v1744
	goto L541
L540:
	;
	goto L539
L541:
	;
	v1763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759))))
	if v1763 != 0 {
		v1759 = v1759 + int32(1)
		goto L541
	} else {
		goto L543
	}
L542:
	;
	v1765 = v1759
	goto L528
L543:
	;
	goto L542
L544:
	;
	v1776 = int32(0)
	v1837 = v1685
	v1838 = v1776
	v1839 = v540
	v1840 = v550
	v1841 = v1776
	v1842 = v558
	goto L517
L545:
	;
	goto L546
L546:
	;
	v1778 = int32(0)
	if base.Ui32(l2+(l4-v1773)) < base.Ui32(v1685) {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v1837 = v1685
	v1838 = int32(0)
	v1839 = v540
	v1840 = v550
	v1841 = v1778
	v1842 = v558
	goto L517
L548:
	;
	goto L549
L549:
	;
	if v1773 == int32(0) {
		goto L551
	} else {
		goto L552
	}
L550:
	;
	if v1826 != 0 {
		goto L564
	} else {
		goto L565
	}
L551:
	;
	v1826 = int32(0)
	goto L550
L552:
	;
	goto L553
L553:
	;
	v1788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1685))))
	if v1788 != 0 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v1789 = v1685
	v1790 = v502
	v1791 = v1773
	v1792 = v1788
	goto L558
L555:
	;
	v1814 = v502
	v1818 = int32(0)
	goto L556
L556:
	;
	v1819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814))))
	v1826 = v1818 - v1819
	goto L550
L557:
	;
	v1814 = v1809
	v1818 = v1811
	goto L556
L558:
	;
	v1794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1790))))
	if v1792 != v1794 {
		v1809 = v1790
		v1811 = v1792
		goto L557
	} else {
		goto L560
	}
L559:
	;
	v1809 = v1803
	v1811 = int32(0)
	goto L557
L560:
	;
	if v1794 == int32(0) {
		v1809 = v1790
		v1811 = v1792
		goto L557
	} else {
		goto L561
	}
L561:
	;
	v1799 = v1791 - int32(1)
	if v1799 == int32(0) {
		v1809 = v1790
		v1811 = v1792
		goto L557
	} else {
		goto L562
	}
L562:
	;
	v1802 = int32(1)
	v1803 = v1790 + v1802
	v1804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1789)+1)))
	if v1804 != 0 {
		v1789 = v1789 + v1802
		v1790 = v1803
		v1791 = v1799
		v1792 = v1804
		goto L558
	} else {
		goto L563
	}
L563:
	;
	goto L559
L564:
	;
	v1837 = v1685
	v1838 = int32(0)
	v1839 = v540
	v1840 = v550
	v1841 = v1778
	v1842 = v558
	goto L517
L565:
	;
	goto L566
L566:
	;
	v1828 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v540))) = uint8(v1828)
	v1830 = int32(1)
	v1837 = v1685 + v1773 - v1830
	v1838 = v1830
	v1839 = v540 + v1830
	v1840 = v550
	v1841 = v1830
	v1842 = v558
	goto L517
L567:
	;
	v1844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1844 != int32(32) {
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	} else {
		goto L568
	}
L568:
	;
	if v1840+v1842 <= int32(0) {
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	} else {
		goto L569
	}
L569:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1852 = v1850 & int32(64)
	v1853 = int32(0)
	if base.B2i32(v1852 == v1853)|(v1838^int32(1)) == v1853 {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	v1861 = v1837 + int32(1)
	if base.Ui32(v521) <= base.Ui32(v1861) {
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	} else {
		goto L573
	}
L571:
	;
	goto L572
L572:
	;
	if v1838|(v1852|base.B2i32(v1850&int32(768) == int32(0))) != 0 {
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	} else {
		goto L644
	}
L573:
	;
	v1863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1861))))
	if base.Ui32((v1863-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	} else {
		goto L574
	}
L574:
	;
	if v501&int32(3) == int32(0) {
		v1893 = v501
		goto L578
	} else {
		goto L579
	}
L575:
	;
	if v500&int32(3) == int32(0) {
		v2002 = v500
		goto L612
	} else {
		goto L613
	}
L576:
	;
	if v1926 == int32(0) {
		goto L575
	} else {
		goto L593
	}
L577:
	;
	v1926 = v1918 - v501
	goto L576
L578:
	;
	v1897 = v1893
	goto L587
L579:
	;
	v1877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if v1877 == int32(0) {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v1926 = int32(0)
	goto L576
L581:
	;
	goto L582
L582:
	;
	v1882 = v501
	goto L583
L583:
	;
	v1886 = v1882 + int32(1)
	if v1886&int32(3) == int32(0) {
		v1893 = v1886
		goto L578
	} else {
		goto L585
	}
L584:
	;
	v1918 = v1886
	goto L577
L585:
	;
	v1891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1886))))
	if v1891 != 0 {
		v1882 = v1886
		goto L583
	} else {
		goto L586
	}
L586:
	;
	goto L584
L587:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1897)))
	v1906 = int32(-2139062144)
	if (int32(16843008)-v1903|v1903)&v1906 == v1906 {
		v1897 = v1897 + int32(4)
		goto L587
	} else {
		goto L589
	}
L588:
	;
	v1912 = v1897
	goto L590
L589:
	;
	goto L588
L590:
	;
	v1916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1912))))
	if v1916 != 0 {
		v1912 = v1912 + int32(1)
		goto L590
	} else {
		goto L592
	}
L591:
	;
	v1918 = v1912
	goto L577
L592:
	;
	goto L591
L593:
	;
	if base.Ui32(l2+(l4-v1926)) < base.Ui32(v1861) {
		goto L575
	} else {
		goto L594
	}
L594:
	;
	if v1926 == int32(0) {
		goto L596
	} else {
		goto L597
	}
L595:
	;
	if v1975 != 0 {
		goto L575
	} else {
		goto L609
	}
L596:
	;
	v1975 = int32(0)
	goto L595
L597:
	;
	goto L598
L598:
	;
	v1937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1861))))
	if v1937 != 0 {
		goto L599
	} else {
		goto L600
	}
L599:
	;
	v1938 = v1861
	v1939 = v501
	v1940 = v1926
	v1941 = v1937
	goto L603
L600:
	;
	v1963 = v501
	v1967 = int32(0)
	goto L601
L601:
	;
	v1968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1963))))
	v1975 = v1967 - v1968
	goto L595
L602:
	;
	v1963 = v1958
	v1967 = v1960
	goto L601
L603:
	;
	v1943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1939))))
	if v1941 != v1943 {
		v1958 = v1939
		v1960 = v1941
		goto L602
	} else {
		goto L605
	}
L604:
	;
	v1958 = v1952
	v1960 = int32(0)
	goto L602
L605:
	;
	if v1943 == int32(0) {
		v1958 = v1939
		v1960 = v1941
		goto L602
	} else {
		goto L606
	}
L606:
	;
	v1948 = v1940 - int32(1)
	if v1948 == int32(0) {
		v1958 = v1939
		v1960 = v1941
		goto L602
	} else {
		goto L607
	}
L607:
	;
	v1951 = int32(1)
	v1952 = v1939 + v1951
	v1953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1938)+1)))
	if v1953 != 0 {
		v1938 = v1938 + v1951
		v1939 = v1952
		v1940 = v1948
		v1941 = v1953
		goto L603
	} else {
		goto L608
	}
L608:
	;
	goto L604
L609:
	;
	v1977 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1977)
	v3521 = v1837 + v1926
	v3526 = v1839
	v3536 = v1840
	v3537 = v1841
	v3544 = v1842
	goto L182
L610:
	;
	if v2035 == int32(0) {
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	} else {
		goto L627
	}
L611:
	;
	v2035 = v2027 - v500
	goto L610
L612:
	;
	v2006 = v2002
	goto L621
L613:
	;
	v1986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	if v1986 == int32(0) {
		goto L614
	} else {
		goto L615
	}
L614:
	;
	v2035 = int32(0)
	goto L610
L615:
	;
	goto L616
L616:
	;
	v1991 = v500
	goto L617
L617:
	;
	v1995 = v1991 + int32(1)
	if v1995&int32(3) == int32(0) {
		v2002 = v1995
		goto L612
	} else {
		goto L619
	}
L618:
	;
	v2027 = v1995
	goto L611
L619:
	;
	v2000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1995))))
	if v2000 != 0 {
		v1991 = v1995
		goto L617
	} else {
		goto L620
	}
L620:
	;
	goto L618
L621:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v2006)))
	v2015 = int32(-2139062144)
	if (int32(16843008)-v2012|v2012)&v2015 == v2015 {
		v2006 = v2006 + int32(4)
		goto L621
	} else {
		goto L623
	}
L622:
	;
	v2021 = v2006
	goto L624
L623:
	;
	goto L622
L624:
	;
	v2025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2021))))
	if v2025 != 0 {
		v2021 = v2021 + int32(1)
		goto L624
	} else {
		goto L626
	}
L625:
	;
	v2027 = v2021
	goto L611
L626:
	;
	goto L625
L627:
	;
	if base.Ui32(l2+(l4-v2035)) < base.Ui32(v1861) {
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	} else {
		goto L628
	}
L628:
	;
	if v2035 == int32(0) {
		goto L630
	} else {
		goto L631
	}
L629:
	;
	if v2084 != 0 {
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	} else {
		goto L643
	}
L630:
	;
	v2084 = int32(0)
	goto L629
L631:
	;
	goto L632
L632:
	;
	v2046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1861))))
	if v2046 != 0 {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v2047 = v1861
	v2048 = v500
	v2049 = v2035
	v2050 = v2046
	goto L637
L634:
	;
	v2072 = v500
	v2076 = int32(0)
	goto L635
L635:
	;
	v2077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2072))))
	v2084 = v2076 - v2077
	goto L629
L636:
	;
	v2072 = v2067
	v2076 = v2069
	goto L635
L637:
	;
	v2052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2048))))
	if v2050 != v2052 {
		v2067 = v2048
		v2069 = v2050
		goto L636
	} else {
		goto L639
	}
L638:
	;
	v2067 = v2061
	v2069 = int32(0)
	goto L636
L639:
	;
	if v2052 == int32(0) {
		v2067 = v2048
		v2069 = v2050
		goto L636
	} else {
		goto L640
	}
L640:
	;
	v2057 = v2049 - int32(1)
	if v2057 == int32(0) {
		v2067 = v2048
		v2069 = v2050
		goto L636
	} else {
		goto L641
	}
L641:
	;
	v2060 = int32(1)
	v2061 = v2048 + v2060
	v2062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2047)+1)))
	if v2062 != 0 {
		v2047 = v2047 + v2060
		v2048 = v2061
		v2049 = v2057
		v2050 = v2062
		goto L637
	} else {
		goto L642
	}
L642:
	;
	goto L638
L643:
	;
	v2088 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2088)
	v3521 = v1861 + v2035 - int32(1)
	v3526 = v1839
	v3536 = v1840
	v3537 = v1841
	v3544 = v1842
	goto L182
L644:
	;
	v2096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1837))))
	switch v2096 - int32(43) {
	case 0, 2:
		goto L183
	default:
		v3521 = v1837
		v3526 = v1839
		v3536 = v1840
		v3537 = v1841
		v3544 = v1842
		goto L182
	}
L645:
	;
	if v541 == int32(0) {
		goto L648
	} else {
		goto L649
	}
L646:
	;
	goto L647
L647:
	;
	if v541 != 0 {
		goto L654
	} else {
		goto L655
	}
L648:
	;
	v2101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v2101&int32(32) != 0 {
		goto L651
	} else {
		goto L652
	}
L649:
	;
	goto L650
L650:
	;
	v2107 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v2107)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L651:
	;
	v3569 = v535
	v3574 = v540
	v3575 = int32(0)
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L652:
	;
	goto L653
L653:
	;
	v2105 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v2105)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L654:
	;
	v2115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v2115 == int32(44) {
		v3521 = v535
		v3526 = v540
		v3536 = v550
		v3537 = v551
		v3544 = v558
		goto L182
	} else {
		goto L657
	}
L655:
	;
	v2109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v2109&int32(32) == int32(0) {
		goto L654
	} else {
		goto L656
	}
L656:
	;
	v3569 = v535
	v3574 = v540
	v3575 = int32(0)
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L657:
	;
	v3569 = v535
	v3574 = v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L658:
	;
	if l7 != 0 {
		goto L675
	} else {
		goto L676
	}
L659:
	;
	v2174 = v2166 - v499
	goto L658
L660:
	;
	v2145 = v2141
	goto L669
L661:
	;
	v2125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	if v2125 == int32(0) {
		goto L662
	} else {
		goto L663
	}
L662:
	;
	v2174 = int32(0)
	goto L658
L663:
	;
	goto L664
L664:
	;
	v2130 = v499
	goto L665
L665:
	;
	v2134 = v2130 + int32(1)
	if v2134&int32(3) == int32(0) {
		v2141 = v2134
		goto L660
	} else {
		goto L667
	}
L666:
	;
	v2166 = v2134
	goto L659
L667:
	;
	v2139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2134))))
	if v2139 != 0 {
		v2130 = v2134
		goto L665
	} else {
		goto L668
	}
L668:
	;
	goto L666
L669:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v2145)))
	v2154 = int32(-2139062144)
	if (int32(16843008)-v2151|v2151)&v2154 == v2154 {
		v2145 = v2145 + int32(4)
		goto L669
	} else {
		goto L671
	}
L670:
	;
	v2160 = v2145
	goto L672
L671:
	;
	goto L670
L672:
	;
	v2164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2160))))
	if v2164 != 0 {
		v2160 = v2160 + int32(1)
		goto L672
	} else {
		goto L674
	}
L673:
	;
	v2166 = v2160
	goto L659
L674:
	;
	goto L673
L675:
	;
	if v541 == int32(0) {
		goto L678
	} else {
		goto L679
	}
L676:
	;
	goto L677
L677:
	;
	if v541 != 0 {
		goto L707
	} else {
		goto L708
	}
L678:
	;
	v2177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v2177&int32(32) != 0 {
		goto L681
	} else {
		goto L682
	}
L679:
	;
	goto L680
L680:
	;
	if (v499^v535)&int32(3) != 0 {
		goto L689
	} else {
		goto L690
	}
L681:
	;
	v3569 = v535
	v3574 = v540
	v3575 = int32(0)
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L682:
	;
	goto L683
L683:
	;
	v2182 = F_pg_mbstrlen(m, v499)
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L88
	} else {
		goto L684
	}
L684:
	;
	v2185 = F__emscripten_memset_bulkmem(m, v535, base.I32_extend8_s(int32(32)), v2182)
	mBase = m.M
	goto L685
L685:
	;
	v3521 = v2185 + v2182 - int32(1)
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L686:
	;
	v3521 = v535 + v2174 - int32(1)
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L687:
	;
	goto L686
L688:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2243))) = uint8(v2242)
	if v2242&int32(255) == int32(0) {
		goto L687
	} else {
		goto L703
	}
L689:
	;
	v2194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	v2241 = v499
	v2242 = v2194
	v2243 = v535
	goto L688
L690:
	;
	goto L691
L691:
	;
	if v499&int32(3) != 0 {
		goto L692
	} else {
		goto L693
	}
L692:
	;
	v2198 = v499
	v2200 = v535
	goto L695
L693:
	;
	v2212 = v499
	v2214 = v535
	goto L694
L694:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v2212)))
	v2219 = int32(-2139062144)
	if (int32(16843008)-v2216|v2216)&v2219 != v2219 {
		v2241 = v2212
		v2242 = v2216
		v2243 = v2214
		goto L688
	} else {
		goto L699
	}
L695:
	;
	v2201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2198))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2200))) = uint8(v2201)
	if v2201 == int32(0) {
		goto L687
	} else {
		goto L697
	}
L696:
	;
	v2212 = v2208
	v2214 = v2206
	goto L694
L697:
	;
	v2205 = int32(1)
	v2206 = v2200 + v2205
	v2208 = v2198 + v2205
	if v2208&int32(3) != 0 {
		v2198 = v2208
		v2200 = v2206
		goto L695
	} else {
		goto L698
	}
L698:
	;
	goto L696
L699:
	;
	v2224 = v2212
	v2225 = v2216
	v2226 = v2214
	goto L700
L700:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2226))) = v2225
	v2228 = int32(4)
	v2229 = v2226 + v2228
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2224)+4))
	v2232 = v2224 + v2228
	v2236 = int32(-2139062144)
	if (v2230|(int32(16843008)-v2230))&v2236 == v2236 {
		v2224 = v2232
		v2225 = v2230
		v2226 = v2229
		goto L700
	} else {
		goto L702
	}
L701:
	;
	v2241 = v2232
	v2242 = v2230
	v2243 = v2229
	goto L688
L702:
	;
	goto L701
L703:
	;
	v2250 = v2241
	v2252 = v2243
	goto L704
L704:
	;
	v2253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2250)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2252)+1)) = uint8(v2253)
	v2255 = int32(1)
	if v2253 != 0 {
		v2250 = v2250 + v2255
		v2252 = v2252 + v2255
		goto L704
	} else {
		goto L706
	}
L705:
	;
	goto L687
L706:
	;
	goto L705
L707:
	;
	if base.Ui32(l2+(l4-v2174)) < base.Ui32(v535) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L710
	}
L708:
	;
	v2266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v2266&int32(32) == int32(0) {
		goto L707
	} else {
		goto L709
	}
L709:
	;
	v3569 = v535
	v3574 = v540
	v3575 = int32(0)
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L710:
	;
	if v2174 == int32(0) {
		goto L712
	} else {
		goto L713
	}
L711:
	;
	if v2318 != 0 {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L725
	}
L712:
	;
	v2318 = int32(0)
	goto L711
L713:
	;
	goto L714
L714:
	;
	v2280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v2280 != 0 {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	v2281 = v535
	v2282 = v499
	v2283 = v2174
	v2284 = v2280
	goto L719
L716:
	;
	v2306 = v499
	v2310 = int32(0)
	goto L717
L717:
	;
	v2311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2306))))
	v2318 = v2310 - v2311
	goto L711
L718:
	;
	v2306 = v2301
	v2310 = v2303
	goto L717
L719:
	;
	v2286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2282))))
	if v2284 != v2286 {
		v2301 = v2282
		v2303 = v2284
		goto L718
	} else {
		goto L721
	}
L720:
	;
	v2301 = v2295
	v2303 = int32(0)
	goto L718
L721:
	;
	if v2286 == int32(0) {
		v2301 = v2282
		v2303 = v2284
		goto L718
	} else {
		goto L722
	}
L722:
	;
	v2291 = v2283 - int32(1)
	if v2291 == int32(0) {
		v2301 = v2282
		v2303 = v2284
		goto L718
	} else {
		goto L723
	}
L723:
	;
	v2294 = int32(1)
	v2295 = v2282 + v2294
	v2296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2281)+1)))
	if v2296 != 0 {
		v2281 = v2281 + v2294
		v2282 = v2295
		v2283 = v2291
		v2284 = v2296
		goto L719
	} else {
		goto L724
	}
L724:
	;
	goto L720
L725:
	;
	v3521 = v535 + v2174 - int32(1)
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L726:
	;
	if (v503^v535)&int32(3) != 0 {
		goto L732
	} else {
		goto L733
	}
L727:
	;
	goto L728
L728:
	;
	v2456 = F_pg_mbstrlen(m, v503)
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L88
	} else {
		goto L767
	}
L729:
	;
	if v503&int32(3) == int32(0) {
		v2419 = v503
		goto L752
	} else {
		goto L753
	}
L730:
	;
	goto L729
L731:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2376))) = uint8(v2375)
	if v2375&int32(255) == int32(0) {
		goto L730
	} else {
		goto L746
	}
L732:
	;
	v2327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
	v2374 = v503
	v2375 = v2327
	v2376 = v535
	goto L731
L733:
	;
	goto L734
L734:
	;
	if v503&int32(3) != 0 {
		goto L735
	} else {
		goto L736
	}
L735:
	;
	v2331 = v503
	v2333 = v535
	goto L738
L736:
	;
	v2345 = v503
	v2347 = v535
	goto L737
L737:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(v2345)))
	v2352 = int32(-2139062144)
	if (int32(16843008)-v2349|v2349)&v2352 != v2352 {
		v2374 = v2345
		v2375 = v2349
		v2376 = v2347
		goto L731
	} else {
		goto L742
	}
L738:
	;
	v2334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2331))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2333))) = uint8(v2334)
	if v2334 == int32(0) {
		goto L730
	} else {
		goto L740
	}
L739:
	;
	v2345 = v2341
	v2347 = v2339
	goto L737
L740:
	;
	v2338 = int32(1)
	v2339 = v2333 + v2338
	v2341 = v2331 + v2338
	if v2341&int32(3) != 0 {
		v2331 = v2341
		v2333 = v2339
		goto L738
	} else {
		goto L741
	}
L741:
	;
	goto L739
L742:
	;
	v2357 = v2345
	v2358 = v2349
	v2359 = v2347
	goto L743
L743:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2359))) = v2358
	v2361 = int32(4)
	v2362 = v2359 + v2361
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v2357)+4))
	v2365 = v2357 + v2361
	v2369 = int32(-2139062144)
	if (v2363|(int32(16843008)-v2363))&v2369 == v2369 {
		v2357 = v2365
		v2358 = v2363
		v2359 = v2362
		goto L743
	} else {
		goto L745
	}
L744:
	;
	v2374 = v2365
	v2375 = v2363
	v2376 = v2362
	goto L731
L745:
	;
	goto L744
L746:
	;
	v2383 = v2374
	v2385 = v2376
	goto L747
L747:
	;
	v2386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2383)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2385)+1)) = uint8(v2386)
	v2388 = int32(1)
	if v2386 != 0 {
		v2383 = v2383 + v2388
		v2385 = v2385 + v2388
		goto L747
	} else {
		goto L749
	}
L748:
	;
	goto L730
L749:
	;
	goto L748
L750:
	;
	v3521 = v535 + v2452 - int32(1)
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L751:
	;
	v2452 = v2444 - v503
	goto L750
L752:
	;
	v2423 = v2419
	goto L761
L753:
	;
	v2403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
	if v2403 == int32(0) {
		goto L754
	} else {
		goto L755
	}
L754:
	;
	v2452 = int32(0)
	goto L750
L755:
	;
	goto L756
L756:
	;
	v2408 = v503
	goto L757
L757:
	;
	v2412 = v2408 + int32(1)
	if v2412&int32(3) == int32(0) {
		v2419 = v2412
		goto L752
	} else {
		goto L759
	}
L758:
	;
	v2444 = v2412
	goto L751
L759:
	;
	v2417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2412))))
	if v2417 != 0 {
		v2408 = v2412
		goto L757
	} else {
		goto L760
	}
L760:
	;
	goto L758
L761:
	;
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(v2423)))
	v2432 = int32(-2139062144)
	if (int32(16843008)-v2429|v2429)&v2432 == v2432 {
		v2423 = v2423 + int32(4)
		goto L761
	} else {
		goto L763
	}
L762:
	;
	v2438 = v2423
	goto L764
L763:
	;
	goto L762
L764:
	;
	v2442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2438))))
	if v2442 != 0 {
		v2438 = v2438 + int32(1)
		goto L764
	} else {
		goto L766
	}
L765:
	;
	v2444 = v2438
	goto L751
L766:
	;
	goto L765
L767:
	;
	if v2456 <= int32(0) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L768
	}
L768:
	;
	v2468 = v535
	v2469 = v2456
	goto L769
L769:
	;
	if base.Ui32(v521) <= base.Ui32(v2468) {
		v3569 = v2468
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L771
	}
L770:
	;
	v3569 = v2518
	v3574 = v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L771:
	;
	v2507 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2468))))
	if base.Ui64(v2507) <= base.Ui64(int64(63)) {
		goto L772
	} else {
		goto L773
	}
L772:
	;
	if int64(1)<<(uint(v2507)%64)&int64(288080842570334209) != int64(0) {
		v3569 = v2468
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L775
	}
L773:
	;
	goto L774
L774:
	;
	v2516 = F_pg_mblen_range(m, v2468, v521)
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L88
	} else {
		goto L776
	}
L775:
	;
	goto L774
L776:
	;
	v2518 = v2516 + v2468
	v2519 = int32(1)
	if base.Ui32(v2519) < base.Ui32(v2469) {
		v2468 = v2518
		v2469 = v2469 - v2519
		goto L769
	} else {
		goto L777
	}
L777:
	;
	goto L770
L778:
	;
	if v722 != int32(30) {
		v2662 = v540
		goto L781
	} else {
		goto L782
	}
L779:
	;
	goto L780
L780:
	;
	if base.Ui32(v521) <= base.Ui32(v535) {
		v2903 = v535
		v2941 = v535
		goto L855
	} else {
		goto L856
	}
L781:
	;
	v2697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v2697&int32(32) != 0 {
		goto L813
	} else {
		goto L814
	}
L782:
	;
	if v540 == int32(0) {
		goto L783
	} else {
		goto L784
	}
L783:
	;
	v2662 = int32(0)
	goto L781
L784:
	;
	goto L785
L785:
	;
	if v540&int32(3) == int32(0) {
		v2551 = v540
		goto L788
	} else {
		goto L789
	}
L786:
	;
	v2585 = F_pnstrdup(m, v540, v2584)
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
		goto L88
	} else {
		goto L803
	}
L787:
	;
	v2584 = v2576 - v540
	goto L786
L788:
	;
	v2555 = v2551
	goto L797
L789:
	;
	v2535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540))))
	if v2535 == int32(0) {
		goto L790
	} else {
		goto L791
	}
L790:
	;
	v2584 = int32(0)
	goto L786
L791:
	;
	goto L792
L792:
	;
	v2540 = v540
	goto L793
L793:
	;
	v2544 = v2540 + int32(1)
	if v2544&int32(3) == int32(0) {
		v2551 = v2544
		goto L788
	} else {
		goto L795
	}
L794:
	;
	v2576 = v2544
	goto L787
L795:
	;
	v2549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2544))))
	if v2549 != 0 {
		v2540 = v2544
		goto L793
	} else {
		goto L796
	}
L796:
	;
	goto L794
L797:
	;
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(v2555)))
	v2564 = int32(-2139062144)
	if (int32(16843008)-v2561|v2561)&v2564 == v2564 {
		v2555 = v2555 + int32(4)
		goto L797
	} else {
		goto L799
	}
L798:
	;
	v2570 = v2555
	goto L800
L799:
	;
	goto L798
L800:
	;
	v2574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2570))))
	if v2574 != 0 {
		v2570 = v2570 + int32(1)
		goto L800
	} else {
		goto L802
	}
L801:
	;
	v2576 = v2570
	goto L787
L802:
	;
	goto L801
L803:
	;
	v2587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2585))))
	if v2587 == int32(0) {
		v2662 = v2585
		goto L781
	} else {
		goto L804
	}
L804:
	;
	v2599 = v2585
	v2600 = v2587
	goto L805
L805:
	;
	v2636 = int32(255)
	v2637 = v2600 & v2636
	if base.Ui32((v2637-int32(65))&v2636) < base.Ui32(int32(26)) {
		goto L808
	} else {
		goto L809
	}
L806:
	;
	v2662 = v2585
	goto L781
L807:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2599))) = uint8(v2646)
	v2649 = v2599 + int32(1)
	v2650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2649))))
	if v2650 != 0 {
		v2599 = v2649
		v2600 = v2650
		goto L805
	} else {
		goto L811
	}
L808:
	;
	v2646 = v2637 | int32(32)
	goto L810
L809:
	;
	v2646 = v2637
	goto L810
L810:
	;
	goto L807
L811:
	;
	goto L806
L812:
	;
	if v535&int32(3) == int32(0) {
		v2801 = v535
		goto L840
	} else {
		goto L841
	}
L813:
	;
	if (v2662^v535)&int32(3) != 0 {
		goto L819
	} else {
		goto L820
	}
L814:
	;
	goto L815
L815:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v2662
	v2776 = F_pg_sprintf(m, v535, int32(172615), v49)
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L88
	} else {
		goto L837
	}
L816:
	;
	goto L812
L817:
	;
	goto L816
L818:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2754))) = uint8(v2753)
	if v2753&int32(255) == int32(0) {
		goto L817
	} else {
		goto L833
	}
L819:
	;
	v2705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2662))))
	v2752 = v2662
	v2753 = v2705
	v2754 = v535
	goto L818
L820:
	;
	goto L821
L821:
	;
	if v2662&int32(3) != 0 {
		goto L822
	} else {
		goto L823
	}
L822:
	;
	v2709 = v2662
	v2711 = v535
	goto L825
L823:
	;
	v2723 = v2662
	v2725 = v535
	goto L824
L824:
	;
	v2727 = *(*int32)(unsafe.Add(mBase, uint32(v2723)))
	v2730 = int32(-2139062144)
	if (int32(16843008)-v2727|v2727)&v2730 != v2730 {
		v2752 = v2723
		v2753 = v2727
		v2754 = v2725
		goto L818
	} else {
		goto L829
	}
L825:
	;
	v2712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2709))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2711))) = uint8(v2712)
	if v2712 == int32(0) {
		goto L817
	} else {
		goto L827
	}
L826:
	;
	v2723 = v2719
	v2725 = v2717
	goto L824
L827:
	;
	v2716 = int32(1)
	v2717 = v2711 + v2716
	v2719 = v2709 + v2716
	if v2719&int32(3) != 0 {
		v2709 = v2719
		v2711 = v2717
		goto L825
	} else {
		goto L828
	}
L828:
	;
	goto L826
L829:
	;
	v2735 = v2723
	v2736 = v2727
	v2737 = v2725
	goto L830
L830:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2737))) = v2736
	v2739 = int32(4)
	v2740 = v2737 + v2739
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(v2735)+4))
	v2743 = v2735 + v2739
	v2747 = int32(-2139062144)
	if (v2741|(int32(16843008)-v2741))&v2747 == v2747 {
		v2735 = v2743
		v2736 = v2741
		v2737 = v2740
		goto L830
	} else {
		goto L832
	}
L831:
	;
	v2752 = v2743
	v2753 = v2741
	v2754 = v2740
	goto L818
L832:
	;
	goto L831
L833:
	;
	v2761 = v2752
	v2763 = v2754
	goto L834
L834:
	;
	v2764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2761)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2763)+1)) = uint8(v2764)
	v2766 = int32(1)
	if v2764 != 0 {
		v2761 = v2761 + v2766
		v2763 = v2763 + v2766
		goto L834
	} else {
		goto L836
	}
L835:
	;
	goto L817
L836:
	;
	goto L835
L837:
	;
	goto L812
L838:
	;
	v3521 = v2834 + v535 - int32(1)
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L839:
	;
	v2834 = v2826 - v535
	goto L838
L840:
	;
	v2805 = v2801
	goto L849
L841:
	;
	v2785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v2785 == int32(0) {
		goto L842
	} else {
		goto L843
	}
L842:
	;
	v2834 = int32(0)
	goto L838
L843:
	;
	goto L844
L844:
	;
	v2790 = v535
	goto L845
L845:
	;
	v2794 = v2790 + int32(1)
	if v2794&int32(3) == int32(0) {
		v2801 = v2794
		goto L840
	} else {
		goto L847
	}
L846:
	;
	v2826 = v2794
	goto L839
L847:
	;
	v2799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2794))))
	if v2799 != 0 {
		v2790 = v2794
		goto L845
	} else {
		goto L848
	}
L848:
	;
	goto L846
L849:
	;
	v2811 = *(*int32)(unsafe.Add(mBase, uint32(v2805)))
	v2814 = int32(-2139062144)
	if (int32(16843008)-v2811|v2811)&v2814 == v2814 {
		v2805 = v2805 + int32(4)
		goto L849
	} else {
		goto L851
	}
L850:
	;
	v2820 = v2805
	goto L852
L851:
	;
	goto L850
L852:
	;
	v2824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2820))))
	if v2824 != 0 {
		v2820 = v2820 + int32(1)
		goto L852
	} else {
		goto L854
	}
L853:
	;
	v2826 = v2820
	goto L839
L854:
	;
	goto L853
L855:
	;
	v2951 = v2903
	v2952 = int32(0)
	v2953 = v2941
	goto L864
L856:
	;
	v2847 = v535
	goto L857
L857:
	;
	v2885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2847))))
	if base.Ui32(v2885-int32(9)) < base.Ui32(int32(5)) {
		goto L859
	} else {
		goto L860
	}
L858:
	;
	v2903 = v2893
	v2941 = v521
	goto L855
L859:
	;
	v2893 = v2847 + int32(1)
	if v2893 != v521 {
		v2847 = v2893
		goto L857
	} else {
		goto L862
	}
L860:
	;
	if v2885 == int32(32) {
		goto L859
	} else {
		goto L861
	}
L861:
	;
	v2903 = v2847
	v2941 = v2847
	goto L855
L862:
	;
	goto L858
L863:
	;
	v3038 = int32(1)
	v3040 = int32(0)
	v3057 = v3040
	v3073 = v3040
	v3074 = v3040
	v3075 = v3040
	v3076 = v3038
	v3080 = v3040
	v3086 = v3040
	v3087 = v3040
	goto L881
L864:
	;
	if base.Ui32(v521) <= base.Ui32(v2953) {
		goto L866
	} else {
		goto L867
	}
L865:
	;
	if v2952 == int32(0) {
		goto L5
	} else {
		goto L880
	}
L866:
	;
	goto L865
L867:
	;
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2953))))
	if base.Ui32((v2991-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L876
	} else {
		goto L877
	}
L868:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(97)+v2952))) = uint8(v3002)
	*(*int32)(unsafe.Add(mBase, uint32(v49+int32(32)+v2952<<(uint(int32(2))%32)))) = v3011
	v3022 = int32(15)
	v3023 = int32(1)
	v3024 = v2951 + v3023
	v3026 = v2952 + v3023
	if v3026 != v3022 {
		v2951 = v3024
		v2952 = v3026
		v2953 = v3024
		goto L864
	} else {
		goto L879
	}
L869:
	;
	v3011 = int32(1000)
	goto L868
L870:
	;
	v3011 = int32(500)
	goto L868
L871:
	;
	v3011 = int32(100)
	goto L868
L872:
	;
	v3011 = int32(50)
	goto L868
L873:
	;
	v3011 = int32(10)
	goto L868
L874:
	;
	v3011 = int32(5)
	goto L868
L875:
	;
	switch v3002 - int32(67) {
	case 0:
		goto L871
	case 1:
		goto L870
	default:
		goto L866
	case 6:
		v3011 = int32(1)
		goto L868
	case 9:
		goto L872
	case 10:
		goto L869
	case 19:
		goto L874
	case 21:
		goto L873
	}
L876:
	;
	v3000 = v2991 - int32(32)
	goto L878
L877:
	;
	v3000 = v2991
	goto L878
L878:
	;
	v3002 = v3000 & int32(255)
	goto L875
L879:
	;
	v3033 = v3024
	v3037 = v3022
	goto L863
L880:
	;
	v3033 = v2951
	v3037 = v2952
	goto L863
L881:
	;
	v3097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(97)+v3057))))
	v3103 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(32)+v3057<<(uint(int32(2))%32))))
	if base.B2i32(v3086 <= v3103)&v3087 != 0 {
		goto L5
	} else {
		goto L883
	}
L882:
	;
	if v3195 < int32(0) {
		goto L5
	} else {
		goto L933
	}
L883:
	;
	if int32(4) < v3103 {
		goto L884
	} else {
		goto L885
	}
L884:
	;
	v3109 = v3073
	goto L886
L885:
	;
	v3109 = int32(0)
	goto L886
L886:
	;
	if v3109 != 0 {
		goto L5
	} else {
		goto L887
	}
L887:
	;
	if int32(49) < v3103 {
		goto L888
	} else {
		goto L889
	}
L888:
	;
	v3113 = v3074
	goto L890
L889:
	;
	v3113 = int32(0)
	goto L890
L890:
	;
	if v3113 != 0 {
		goto L5
	} else {
		goto L891
	}
L891:
	;
	if int32(499) < v3103 {
		goto L892
	} else {
		goto L893
	}
L892:
	;
	v3117 = v3075
	goto L894
L893:
	;
	v3117 = int32(0)
	goto L894
L894:
	;
	if v3117 != 0 {
		goto L5
	} else {
		goto L895
	}
L895:
	;
	switch v3097 - int32(68) {
	case 0:
		goto L897
	default:
		v3126 = v3073
		v3127 = v3074
		v3128 = v3075
		goto L896
	case 8:
		goto L898
	case 18:
		goto L899
	}
L896:
	;
	if v3037-v3038 <= v3057 {
		goto L901
	} else {
		goto L902
	}
L897:
	;
	v3126 = v3073
	v3127 = v3074
	v3128 = v3075 + int32(1)
	goto L896
L898:
	;
	v3126 = v3073
	v3127 = v3074 + int32(1)
	v3128 = v3075
	goto L896
L899:
	;
	v3126 = v3073 + int32(1)
	v3127 = v3074
	v3128 = v3075
	goto L896
L900:
	;
	v3195 = v3185 + v3080
	v3197 = v3186 + int32(1)
	if v3197 < v3037 {
		v3057 = v3197
		v3073 = v3187
		v3074 = v3188
		v3075 = v3189
		v3076 = v3190
		v3080 = v3195
		v3086 = v3193
		v3087 = v3194
		goto L881
	} else {
		goto L932
	}
L901:
	;
	v3185 = v3103
	v3186 = v3057
	v3187 = v3126
	v3188 = v3127
	v3189 = v3128
	v3190 = v3076
	v3193 = v3086
	v3194 = v3087
	goto L900
L902:
	;
	goto L903
L903:
	;
	v3131 = v3057 + int32(1)
	v3135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3131+(v49+int32(97))))))
	v3141 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(32)+v3131<<(uint(int32(2))%32))))
	if v3103 < v3141 {
		goto L904
	} else {
		goto L905
	}
L904:
	;
	switch v3097 - int32(67) {
	case 0:
		goto L908
	default:
		goto L5
	case 6:
		goto L910
	case 21:
		goto L909
	}
L905:
	;
	goto L906
L906:
	;
	if v3135 != v3097 {
		goto L928
	} else {
		goto L929
	}
L907:
	;
	if int32(1) < v3076 {
		goto L5
	} else {
		goto L911
	}
L908:
	;
	switch v3135 - int32(68) {
	case 0, 9:
		goto L907
	default:
		goto L5
	}
L909:
	;
	switch v3135 - int32(67) {
	case 0, 9:
		goto L907
	default:
		goto L5
	}
L910:
	;
	switch v3135 - int32(86) {
	case 0, 2:
		goto L907
	default:
		goto L5
	}
L911:
	;
	if int32(4) < v3141 {
		goto L912
	} else {
		goto L913
	}
L912:
	;
	v3156 = v3126
	goto L914
L913:
	;
	v3156 = int32(0)
	goto L914
L914:
	;
	if v3156 != 0 {
		goto L5
	} else {
		goto L915
	}
L915:
	;
	if int32(49) < v3141 {
		goto L916
	} else {
		goto L917
	}
L916:
	;
	v3160 = v3127
	goto L918
L917:
	;
	v3160 = int32(0)
	goto L918
L918:
	;
	if v3160 != 0 {
		goto L5
	} else {
		goto L919
	}
L919:
	;
	if int32(499) < v3141 {
		goto L920
	} else {
		goto L921
	}
L920:
	;
	v3164 = v3128
	goto L922
L921:
	;
	v3164 = int32(0)
	goto L922
L922:
	;
	if v3164 != 0 {
		goto L5
	} else {
		goto L923
	}
L923:
	;
	switch v3135 - int32(68) {
	case 0:
		goto L925
	default:
		v3173 = v3126
		v3174 = v3127
		v3175 = v3128
		goto L924
	case 8:
		goto L926
	case 18:
		goto L927
	}
L924:
	;
	v3177 = int32(1)
	v3185 = v3141 - v3103
	v3186 = v3131
	v3187 = v3173
	v3188 = v3174
	v3189 = v3175
	v3190 = v3177
	v3193 = v3103
	v3194 = v3177
	goto L900
L925:
	;
	v3173 = v3126
	v3174 = v3127
	v3175 = v3128 + int32(1)
	goto L924
L926:
	;
	v3173 = v3126
	v3174 = v3127 + int32(1)
	v3175 = v3128
	goto L924
L927:
	;
	v3173 = v3126 + int32(1)
	v3174 = v3127
	v3175 = v3128
	goto L924
L928:
	;
	v3185 = v3103
	v3186 = v3057
	v3187 = v3126
	v3188 = v3127
	v3189 = v3128
	v3190 = int32(1)
	v3193 = v3086
	v3194 = v3087
	goto L900
L929:
	;
	goto L930
L930:
	;
	v3182 = v3076 + int32(1)
	if int32(3) < v3182 {
		goto L5
	} else {
		goto L931
	}
L931:
	;
	v3185 = v3103
	v3186 = v3057
	v3187 = v3126
	v3188 = v3127
	v3189 = v3128
	v3190 = v3182
	v3193 = v3086
	v3194 = v3087
	goto L900
L932:
	;
	goto L882
L933:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v3195
	v3205 = F_pg_sprintf(m, v540, int32(480823), v49+int32(16))
	mBase = m.M
	v3206 = m.ExcPending
	if v3206 != 0 {
		goto L88
	} else {
		goto L934
	}
L934:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3205
	v3569 = v3033
	v3574 = v3205 + v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L935:
	;
	if v3211&int32(2) != 0 {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L936
	}
L936:
	;
	v3216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v3216 == int32(35) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L937
	}
L937:
	;
	if v412 == int32(45) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L938
	}
L938:
	;
	if l7 != 0 {
		goto L939
	} else {
		goto L940
	}
L939:
	;
	v3220 = F_get_th(m, l3, int32(2))
	mBase = m.M
	v3221 = m.ExcPending
	if v3221 != 0 {
		goto L88
	} else {
		goto L942
	}
L940:
	;
	goto L941
L941:
	;
	if base.Ui32(v521) <= base.Ui32(v535) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L964
	}
L942:
	;
	if (v3220^v535)&int32(3) != 0 {
		goto L946
	} else {
		goto L947
	}
L943:
	;
	v3521 = v535 + int32(1)
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L944:
	;
	goto L943
L945:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3276))) = uint8(v3275)
	if v3275&int32(255) == int32(0) {
		goto L944
	} else {
		goto L960
	}
L946:
	;
	v3227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3220))))
	v3274 = v3220
	v3275 = v3227
	v3276 = v535
	goto L945
L947:
	;
	goto L948
L948:
	;
	if v3220&int32(3) != 0 {
		goto L949
	} else {
		goto L950
	}
L949:
	;
	v3231 = v3220
	v3233 = v535
	goto L952
L950:
	;
	v3245 = v3220
	v3247 = v535
	goto L951
L951:
	;
	v3249 = *(*int32)(unsafe.Add(mBase, uint32(v3245)))
	v3252 = int32(-2139062144)
	if (int32(16843008)-v3249|v3249)&v3252 != v3252 {
		v3274 = v3245
		v3275 = v3249
		v3276 = v3247
		goto L945
	} else {
		goto L956
	}
L952:
	;
	v3234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3231))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3233))) = uint8(v3234)
	if v3234 == int32(0) {
		goto L944
	} else {
		goto L954
	}
L953:
	;
	v3245 = v3241
	v3247 = v3239
	goto L951
L954:
	;
	v3238 = int32(1)
	v3239 = v3233 + v3238
	v3241 = v3231 + v3238
	if v3241&int32(3) != 0 {
		v3231 = v3241
		v3233 = v3239
		goto L952
	} else {
		goto L955
	}
L955:
	;
	goto L953
L956:
	;
	v3257 = v3245
	v3258 = v3249
	v3259 = v3247
	goto L957
L957:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3259))) = v3258
	v3261 = int32(4)
	v3262 = v3259 + v3261
	v3263 = *(*int32)(unsafe.Add(mBase, uint32(v3257)+4))
	v3265 = v3257 + v3261
	v3269 = int32(-2139062144)
	if (v3263|(int32(16843008)-v3263))&v3269 == v3269 {
		v3257 = v3265
		v3258 = v3263
		v3259 = v3262
		goto L957
	} else {
		goto L959
	}
L958:
	;
	v3274 = v3265
	v3275 = v3263
	v3276 = v3262
	goto L945
L959:
	;
	goto L958
L960:
	;
	v3283 = v3274
	v3285 = v3276
	goto L961
L961:
	;
	v3286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3283)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3285)+1)) = uint8(v3286)
	v3288 = int32(1)
	if v3286 != 0 {
		v3283 = v3283 + v3288
		v3285 = v3285 + v3288
		goto L961
	} else {
		goto L963
	}
L962:
	;
	goto L944
L963:
	;
	goto L962
L964:
	;
	v3299 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if base.Ui64(v3299) <= base.Ui64(int64(63)) {
		goto L965
	} else {
		goto L966
	}
L965:
	;
	if int64(1)<<(uint(v3299)%64)&int64(288080842570334209) != int64(0) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L968
	}
L966:
	;
	goto L967
L967:
	;
	v3308 = F_pg_mblen_range(m, v535, v521)
	mBase = m.M
	v3309 = m.ExcPending
	if v3309 != 0 {
		goto L88
	} else {
		goto L969
	}
L968:
	;
	goto L967
L969:
	;
	v3310 = v3308 + v535
	if base.Ui32(v521) <= base.Ui32(v3310) {
		v3569 = v3310
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L970
	}
L970:
	;
	v3312 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3310))))
	if base.Ui64(v3312) <= base.Ui64(int64(63)) {
		goto L971
	} else {
		goto L972
	}
L971:
	;
	if int64(1)<<(uint(v3312)%64)&int64(288080842570334209) != int64(0) {
		v3569 = v3310
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L974
	}
L972:
	;
	goto L973
L973:
	;
	v3321 = F_pg_mblen_range(m, v3310, v521)
	mBase = m.M
	v3322 = m.ExcPending
	if v3322 != 0 {
		goto L88
	} else {
		goto L975
	}
L974:
	;
	goto L973
L975:
	;
	v3569 = v3321 + v3310
	v3574 = v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L976:
	;
	if v3324&int32(2) != 0 {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L977
	}
L977:
	;
	v3329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v3329 == int32(35) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L978
	}
L978:
	;
	if v412 == int32(45) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L979
	}
L979:
	;
	if l7 != 0 {
		goto L980
	} else {
		goto L981
	}
L980:
	;
	v3333 = F_get_th(m, l3, int32(1))
	mBase = m.M
	v3334 = m.ExcPending
	if v3334 != 0 {
		goto L88
	} else {
		goto L983
	}
L981:
	;
	goto L982
L982:
	;
	if base.Ui32(v521) <= base.Ui32(v535) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1005
	}
L983:
	;
	if (v3333^v535)&int32(3) != 0 {
		goto L987
	} else {
		goto L988
	}
L984:
	;
	v3521 = v535 + int32(1)
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L985:
	;
	goto L984
L986:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3389))) = uint8(v3388)
	if v3388&int32(255) == int32(0) {
		goto L985
	} else {
		goto L1001
	}
L987:
	;
	v3340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3333))))
	v3387 = v3333
	v3388 = v3340
	v3389 = v535
	goto L986
L988:
	;
	goto L989
L989:
	;
	if v3333&int32(3) != 0 {
		goto L990
	} else {
		goto L991
	}
L990:
	;
	v3344 = v3333
	v3346 = v535
	goto L993
L991:
	;
	v3358 = v3333
	v3360 = v535
	goto L992
L992:
	;
	v3362 = *(*int32)(unsafe.Add(mBase, uint32(v3358)))
	v3365 = int32(-2139062144)
	if (int32(16843008)-v3362|v3362)&v3365 != v3365 {
		v3387 = v3358
		v3388 = v3362
		v3389 = v3360
		goto L986
	} else {
		goto L997
	}
L993:
	;
	v3347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3344))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3346))) = uint8(v3347)
	if v3347 == int32(0) {
		goto L985
	} else {
		goto L995
	}
L994:
	;
	v3358 = v3354
	v3360 = v3352
	goto L992
L995:
	;
	v3351 = int32(1)
	v3352 = v3346 + v3351
	v3354 = v3344 + v3351
	if v3354&int32(3) != 0 {
		v3344 = v3354
		v3346 = v3352
		goto L993
	} else {
		goto L996
	}
L996:
	;
	goto L994
L997:
	;
	v3370 = v3358
	v3371 = v3362
	v3372 = v3360
	goto L998
L998:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3372))) = v3371
	v3374 = int32(4)
	v3375 = v3372 + v3374
	v3376 = *(*int32)(unsafe.Add(mBase, uint32(v3370)+4))
	v3378 = v3370 + v3374
	v3382 = int32(-2139062144)
	if (v3376|(int32(16843008)-v3376))&v3382 == v3382 {
		v3370 = v3378
		v3371 = v3376
		v3372 = v3375
		goto L998
	} else {
		goto L1000
	}
L999:
	;
	v3387 = v3378
	v3388 = v3376
	v3389 = v3375
	goto L986
L1000:
	;
	goto L999
L1001:
	;
	v3396 = v3387
	v3398 = v3389
	goto L1002
L1002:
	;
	v3399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3396)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3398)+1)) = uint8(v3399)
	v3401 = int32(1)
	if v3399 != 0 {
		v3396 = v3396 + v3401
		v3398 = v3398 + v3401
		goto L1002
	} else {
		goto L1004
	}
L1003:
	;
	goto L985
L1004:
	;
	goto L1003
L1005:
	;
	v3412 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if base.Ui64(v3412) <= base.Ui64(int64(63)) {
		goto L1006
	} else {
		goto L1007
	}
L1006:
	;
	if int64(1)<<(uint(v3412)%64)&int64(288080842570334209) != int64(0) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1009
	}
L1007:
	;
	goto L1008
L1008:
	;
	v3421 = F_pg_mblen_range(m, v535, v521)
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L88
	} else {
		goto L1010
	}
L1009:
	;
	goto L1008
L1010:
	;
	v3423 = v3421 + v535
	if base.Ui32(v521) <= base.Ui32(v3423) {
		v3569 = v3423
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1011
	}
L1011:
	;
	v3425 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3423))))
	if base.Ui64(v3425) <= base.Ui64(int64(63)) {
		goto L1012
	} else {
		goto L1013
	}
L1012:
	;
	if int64(1)<<(uint(v3425)%64)&int64(288080842570334209) != int64(0) {
		v3569 = v3423
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1015
	}
L1013:
	;
	goto L1014
L1014:
	;
	v3434 = F_pg_mblen_range(m, v3423, v521)
	mBase = m.M
	v3435 = m.ExcPending
	if v3435 != 0 {
		goto L88
	} else {
		goto L1016
	}
L1015:
	;
	goto L1014
L1016:
	;
	v3569 = v3434 + v3423
	v3574 = v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L1017:
	;
	if v412 == int32(45) {
		goto L1020
	} else {
		goto L1021
	}
L1018:
	;
	goto L1019
L1019:
	;
	v3446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v3446 == int32(45) {
		goto L1024
	} else {
		goto L1025
	}
L1020:
	;
	v3439 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v3439)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L1021:
	;
	goto L1022
L1022:
	;
	v3441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v3441&int32(32) != 0 {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1023
	}
L1023:
	;
	v3444 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v3444)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L1024:
	;
	v3449 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v3449)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L1025:
	;
	goto L1026
L1026:
	;
	if base.Ui32(v521) <= base.Ui32(v535) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1027
	}
L1027:
	;
	if base.Ui32(v3446) <= base.Ui32(int32(63)) {
		goto L1028
	} else {
		goto L1029
	}
L1028:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v3446))%64)&int64(288080842570334209) != int64(0) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1031
	}
L1029:
	;
	goto L1030
L1030:
	;
	v3461 = F_pg_mblen_range(m, v535, v521)
	mBase = m.M
	v3462 = m.ExcPending
	if v3462 != 0 {
		goto L88
	} else {
		goto L1032
	}
L1031:
	;
	goto L1030
L1032:
	;
	v3569 = v3461 + v535
	v3574 = v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L1033:
	;
	if v412 == int32(43) {
		goto L1036
	} else {
		goto L1037
	}
L1034:
	;
	goto L1035
L1035:
	;
	v3473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v3473 == int32(43) {
		goto L1040
	} else {
		goto L1041
	}
L1036:
	;
	v3466 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v3466)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L1037:
	;
	goto L1038
L1038:
	;
	v3468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v3468&int32(32) != 0 {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1039
	}
L1039:
	;
	v3471 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v3471)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L1040:
	;
	v3476 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v3476)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L1041:
	;
	goto L1042
L1042:
	;
	if base.Ui32(v521) <= base.Ui32(v535) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1043
	}
L1043:
	;
	if base.Ui32(v3473) <= base.Ui32(int32(63)) {
		goto L1044
	} else {
		goto L1045
	}
L1044:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v3473))%64)&int64(288080842570334209) != int64(0) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1047
	}
L1045:
	;
	goto L1046
L1046:
	;
	v3488 = F_pg_mblen_range(m, v535, v521)
	mBase = m.M
	v3489 = m.ExcPending
	if v3489 != 0 {
		goto L88
	} else {
		goto L1048
	}
L1047:
	;
	goto L1046
L1048:
	;
	v3569 = v3488 + v535
	v3574 = v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L1049:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v412)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L1050:
	;
	goto L1051
L1051:
	;
	v3492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	switch v3492 - int32(43) {
	case 0:
		goto L1053
	default:
		goto L1052
	case 2:
		goto L1054
	}
L1052:
	;
	if base.Ui32(v521) <= base.Ui32(v535) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1055
	}
L1053:
	;
	v3497 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v3497)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L1054:
	;
	v3495 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v3495)
	v3521 = v535
	v3526 = v540
	v3536 = v550
	v3537 = v551
	v3544 = v558
	goto L182
L1055:
	;
	if base.Ui32(v3492) <= base.Ui32(int32(63)) {
		goto L1056
	} else {
		goto L1057
	}
L1056:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v3492))%64)&int64(288080842570334209) != int64(0) {
		v3569 = v535
		v3574 = v540
		v3575 = v541
		v3577 = v543
		v3579 = v545
		v3584 = v550
		v3585 = v551
		v3592 = v558
		goto L134
	} else {
		goto L1059
	}
L1057:
	;
	goto L1058
L1058:
	;
	v3509 = F_pg_mblen_range(m, v535, v521)
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L88
	} else {
		goto L1060
	}
L1059:
	;
	goto L1058
L1060:
	;
	v3569 = v3509 + v535
	v3574 = v540
	v3575 = v541
	v3577 = v543
	v3579 = v545
	v3584 = v550
	v3585 = v551
	v3592 = v558
	goto L134
L1061:
	;
	goto L133
L1062:
	;
	v3660 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3620))) = uint8(v3660)
	goto L6
L1063:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v3685
	goto L6
L1064:
	;
	v3713 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3709))) = uint8(v3713)
	goto L1063
L1065:
	;
	goto L1066
L1066:
	;
	v3715 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3675))) = uint8(v3715)
	goto L1063
L1067:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v3819 = m.ExcPending
	if v3819 != 0 {
		goto L88
	} else {
		goto L1068
	}
L1068:
	;
	F_errmsg(m, int32(305887), int32(0))
	mBase = m.M
	v3823 = m.ExcPending
	if v3823 != 0 {
		goto L88
	} else {
		goto L1069
	}
L1069:
	;
	F_errfinish(m, int32(490497), int32(6113), int32(207231))
	mBase = m.M
	v3828 = m.ExcPending
	if v3828 != 0 {
		goto L88
	} else {
		goto L1070
	}
L1070:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1071:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3835 = m.ExcPending
	if v3835 != 0 {
		goto L88
	} else {
		goto L1072
	}
L1072:
	;
	F_errmsg(m, int32(64142), int32(0))
	mBase = m.M
	v3839 = m.ExcPending
	if v3839 != 0 {
		goto L88
	} else {
		goto L1073
	}
L1073:
	;
	F_errfinish(m, int32(490497), int32(5836), int32(207231))
	mBase = m.M
	v3844 = m.ExcPending
	if v3844 != 0 {
		goto L88
	} else {
		goto L1074
	}
L1074:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_NormalizeSubWord(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v521 int32
	_ = v521
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v803 int32
	_ = v803
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v975 int32
	_ = v975
	var v985 int32
	_ = v985
	v4 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(1040)
	m.G0 = v24
	if l1&int32(3) == v4 {
		v49 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v88 = F__emscripten_memset_bulkmem(m, v24+int32(528), base.I32_extend8_s(int32(0)), int32(512))
	mBase = m.M
	goto L18
L2:
	;
	v82 = v74 - l1
	goto L1
L3:
	;
	v53 = v49
	goto L12
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v33 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v82 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v38 = l1
	goto L8
L8:
	;
	v42 = v38 + int32(1)
	if v42&int32(3) == int32(0) {
		v49 = v42
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v74 = v42
	goto L2
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v47 != 0 {
		v38 = v42
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v62 = int32(-2139062144)
	if (int32(16843008)-v59|v59)&v62 == v62 {
		v53 = v53 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v68 = v53
	goto L15
L14:
	;
	goto L13
L15:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v72 != 0 {
		v68 = v68 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v74 = v68
	goto L2
L17:
	;
	goto L16
L18:
	;
	v94 = F__emscripten_memset_bulkmem(m, v24+int32(16), base.I32_extend8_s(int32(0)), int32(512))
	mBase = m.M
	goto L19
L19:
	;
	if int32(256) < v82 {
		v985 = v4
		goto L20
	} else {
		goto L21
	}
L20:
	;
	m.G0 = v24 + int32(1040)
	return v985
L21:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v99 = F_palloc(m, int32(4096))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(0)
	v106 = F_FindWord(m, l0, l1, int32(733277), l2)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v106 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v108 = F_pstrdup(m, l1)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v116 = v99
	goto L27
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v117 == int32(0) {
		v359 = v116
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v111 = v99 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v108
	v116 = v111
	goto L27
L29:
	;
	if v97 == int32(0) {
		v956 = v359
		goto L86
	} else {
		goto L87
	}
L30:
	;
	v123 = v117
	v124 = v116
	v130 = v4
	goto L31
L31:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v141&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v359 = v346
	goto L29
L33:
	;
	v258 = v124
	v259 = int32(0)
	goto L62
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if base.Ui32(int32(255)) < base.Ui32(v144) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v152 = v123
	goto L36
L36:
	;
	if v82 < v130 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v235 = v123 + int32(4)
	v242 = v130
	goto L33
L38:
	;
	goto L39
L39:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	if v149 == int32(0) {
		v359 = v124
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v152 = v149
	goto L36
L41:
	;
	v154 = v130
	goto L43
L42:
	;
	v154 = v82
	goto L43
L43:
	;
	v158 = v152
	v165 = v130
	goto L44
L44:
	;
	if v165 == v154 {
		v359 = v124
		goto L29
	} else {
		goto L46
	}
L45:
	;
	v359 = v124
	goto L29
L46:
	;
	v178 = v158 + int32(4)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v184 = v178 + int32(base.Ui32(v179)>>(uint(int32(1))%32))*int32(12)
	if base.Ui32(v184) <= base.Ui32(v178) {
		v359 = v124
		goto L29
	} else {
		goto L47
	}
L47:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v165))))
	v193 = v178
	v194 = v184
	goto L48
L48:
	;
	v210 = int32(12)
	v211 = base.I32_div_s(v194-v193, v210)
	v216 = v193 + int32(base.Ui32(v211)>>(uint(int32(1))%32))*v210
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v219 = v217 & int32(255)
	if v187 == v219 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L45
L50:
	;
	v222 = v165 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v217) {
		v235 = v216
		v242 = v222
		goto L33
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v228 = base.B2i32(base.Ui32(v219) < base.Ui32(v187))
	if base.Ui32(v219) < base.Ui32(v187) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	if v225 != 0 {
		v158 = v225
		v165 = v222
		goto L44
	} else {
		goto L54
	}
L54:
	;
	v359 = v124
	goto L29
L55:
	;
	v229 = v216 + int32(12)
	goto L57
L56:
	;
	v229 = v193
	goto L57
L57:
	;
	if base.Ui32(v219) < base.Ui32(v187) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v230 = v194
	goto L60
L59:
	;
	v230 = v216
	goto L60
L60:
	;
	if base.Ui32(v229) < base.Ui32(v230) {
		v193 = v229
		v194 = v230
		goto L48
	} else {
		goto L61
	}
L61:
	;
	goto L49
L62:
	;
	v276 = v259 << (uint(int32(2)) % 32)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v276+v277)))
	v283 = F_CheckAffix(m, l1, v82, v279, l2, v24+int32(528), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L22
	} else {
		goto L65
	}
L63:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v235)+8))
	if v354 != 0 {
		v123 = v354
		v124 = v346
		v130 = v242
		goto L31
	} else {
		goto L85
	}
L64:
	;
	v349 = v259 + int32(1)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	if base.Ui32(v349) < base.Ui32(int32(base.Ui32(v350)>>(uint(int32(8))%32))) {
		v258 = v346
		v259 = v349
		goto L62
	} else {
		goto L84
	}
L65:
	;
	if v283 == int32(0) {
		v346 = v258
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v289+v276)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v293 = F_FindWord(m, l0, v24+int32(528), v292, l2)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	if v293 == int32(0) {
		v346 = v258
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v297 = int32(0)
	if int32(4088) < v258-v99 {
		v342 = v297
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v346 = v258 + v342<<(uint(int32(2))%32)
	goto L64
L70:
	;
	if v258 != v99 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v303 = v24 + int32(528)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v258-int32(4))))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v310 == int32(0) {
		v329 = v309
		v330 = v310
		goto L75
	} else {
		goto L76
	}
L72:
	;
	goto L73
L73:
	;
	v336 = F_pstrdup(m, v24+int32(528))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L22
	} else {
		goto L83
	}
L74:
	;
	if v330-v329 == int32(0) {
		v342 = v297
		goto L69
	} else {
		goto L82
	}
L75:
	;
	goto L74
L76:
	;
	if v309 != v310 {
		v329 = v309
		v330 = v310
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v314 = v303
	v315 = v306
	goto L78
L78:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	if v319 == int32(0) {
		v329 = v318
		v330 = v319
		goto L75
	} else {
		goto L80
	}
L79:
	;
	v329 = v318
	v330 = v319
	goto L75
L80:
	;
	v322 = int32(1)
	if v318 == v319 {
		v314 = v314 + v322
		v315 = v315 + v322
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L73
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v258)+4)) = int32(0)
	v342 = int32(1)
	goto L69
L84:
	;
	goto L63
L85:
	;
	goto L32
L86:
	;
	if v956 != v99 {
		v985 = v99
		goto L20
	} else {
		goto L221
	}
L87:
	;
	v383 = v359
	v388 = v97
	v394 = v4
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if v402&int32(1) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v956 = v928
	goto L86
L90:
	;
	v521 = v383
	v535 = int32(0)
	goto L119
L91:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	if base.Ui32(int32(255)) < base.Ui32(v405) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v413 = v388
	goto L93
L93:
	;
	if v82 < v394 {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	v507 = v388 + int32(4)
	v510 = v394
	goto L90
L95:
	;
	goto L96
L96:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v388)+12))
	if v410 == int32(0) {
		v956 = v383
		goto L86
	} else {
		goto L97
	}
L97:
	;
	v413 = v410
	goto L93
L98:
	;
	v415 = v394
	goto L100
L99:
	;
	v415 = v82
	goto L100
L100:
	;
	v425 = v413
	v431 = v394
	goto L101
L101:
	;
	if v415 == v431 {
		v956 = v383
		goto L86
	} else {
		goto L103
	}
L102:
	;
	v956 = v383
	goto L86
L103:
	;
	v439 = v425 + int32(4)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v445 = v439 + int32(base.Ui32(v440)>>(uint(int32(1))%32))*int32(12)
	if base.Ui32(v445) <= base.Ui32(v439) {
		v956 = v383
		goto L86
	} else {
		goto L104
	}
L104:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v82+(v431^int32(-1))))))
	v454 = v439
	v456 = v445
	goto L105
L105:
	;
	v473 = int32(12)
	v474 = base.I32_div_s(v456-v454, v473)
	v479 = v454 + int32(base.Ui32(v474)>>(uint(int32(1))%32))*v473
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	v482 = v480 & int32(255)
	if v450 == v482 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L102
L107:
	;
	v485 = v431 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v480) {
		v507 = v479
		v510 = v485
		goto L90
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v491 = base.B2i32(base.Ui32(v482) < base.Ui32(v450))
	if base.Ui32(v482) < base.Ui32(v450) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v479)+8))
	if v488 != 0 {
		v425 = v488
		v431 = v485
		goto L101
	} else {
		goto L111
	}
L111:
	;
	v956 = v383
	goto L86
L112:
	;
	v492 = v479 + int32(12)
	goto L114
L113:
	;
	v492 = v454
	goto L114
L114:
	;
	if base.Ui32(v482) < base.Ui32(v450) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v493 = v456
	goto L117
L116:
	;
	v493 = v479
	goto L117
L117:
	;
	if base.Ui32(v492) < base.Ui32(v493) {
		v454 = v492
		v456 = v493
		goto L105
	} else {
		goto L118
	}
L118:
	;
	goto L106
L119:
	;
	v539 = v535 << (uint(int32(2)) % 32)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v539+v540)))
	v547 = F_CheckAffix(m, l1, v82, v542, l2, v24+int32(528), v24+int32(12))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L22
	} else {
		goto L122
	}
L120:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	if v951 != 0 {
		v383 = v928
		v388 = v951
		v394 = v510
		goto L88
	} else {
		goto L220
	}
L121:
	;
	v946 = v535 + int32(1)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	if base.Ui32(v946) < base.Ui32(int32(base.Ui32(v947)>>(uint(int32(8))%32))) {
		v521 = v928
		v535 = v946
		goto L119
	} else {
		goto L219
	}
L122:
	;
	if v547 == int32(0) {
		v928 = v521
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v553+v539)))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	v557 = F_FindWord(m, l0, v24+int32(528), v556, l2)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L22
	} else {
		goto L124
	}
L124:
	;
	if v557 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v559 = int32(0)
	if int32(4088) < v521-v99 {
		v604 = v559
		goto L128
	} else {
		goto L129
	}
L126:
	;
	v609 = v521
	goto L127
L127:
	;
	v611 = v24 + int32(528)
	if v611&int32(3) == int32(0) {
		v635 = v611
		goto L145
	} else {
		goto L146
	}
L128:
	;
	v609 = v521 + v604<<(uint(int32(2))%32)
	goto L127
L129:
	;
	if v521 != v99 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v565 = v24 + int32(528)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v521-int32(4))))
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	if v572 == int32(0) {
		v591 = v571
		v592 = v572
		goto L134
	} else {
		goto L135
	}
L131:
	;
	goto L132
L132:
	;
	v598 = F_pstrdup(m, v24+int32(528))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L22
	} else {
		goto L142
	}
L133:
	;
	if v592-v591 == int32(0) {
		v604 = v559
		goto L128
	} else {
		goto L141
	}
L134:
	;
	goto L133
L135:
	;
	if v571 != v572 {
		v591 = v571
		v592 = v572
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v576 = v565
	v577 = v568
	goto L137
L137:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+1)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
	if v581 == int32(0) {
		v591 = v580
		v592 = v581
		goto L134
	} else {
		goto L139
	}
L138:
	;
	v591 = v580
	v592 = v581
	goto L134
L139:
	;
	v584 = int32(1)
	if v580 == v581 {
		v576 = v576 + v584
		v577 = v577 + v584
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	goto L132
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521))) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v521)+4)) = int32(0)
	v604 = int32(1)
	goto L128
L143:
	;
	v669 = int32(0)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v670 == v669 {
		v928 = v609
		goto L121
	} else {
		goto L160
	}
L144:
	;
	v668 = v660 - v611
	goto L143
L145:
	;
	v639 = v635
	goto L154
L146:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
	if v619 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v668 = int32(0)
	goto L143
L148:
	;
	goto L149
L149:
	;
	v624 = v611
	goto L150
L150:
	;
	v628 = v624 + int32(1)
	if v628&int32(3) == int32(0) {
		v635 = v628
		goto L145
	} else {
		goto L152
	}
L151:
	;
	v660 = v628
	goto L144
L152:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628))))
	if v633 != 0 {
		v624 = v628
		goto L150
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v639)))
	v648 = int32(-2139062144)
	if (int32(16843008)-v645|v645)&v648 == v648 {
		v639 = v639 + int32(4)
		goto L154
	} else {
		goto L156
	}
L155:
	;
	v654 = v639
	goto L157
L156:
	;
	goto L155
L157:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	if v658 != 0 {
		v654 = v654 + int32(1)
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v660 = v654
	goto L144
L159:
	;
	goto L158
L160:
	;
	v676 = v670
	v677 = v609
	v689 = v669
	goto L161
L161:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	if v694&int32(1) != 0 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v928 = v914
	goto L121
L163:
	;
	v813 = v677
	v814 = int32(0)
	goto L192
L164:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	if base.Ui32(int32(255)) < base.Ui32(v697) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v705 = v676
	goto L166
L166:
	;
	if v668 < v689 {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v790 = v676 + int32(4)
	v803 = v689
	goto L163
L168:
	;
	goto L169
L169:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v676)+12))
	if v702 == int32(0) {
		v928 = v677
		goto L121
	} else {
		goto L170
	}
L170:
	;
	v705 = v702
	goto L166
L171:
	;
	v707 = v689
	goto L173
L172:
	;
	v707 = v668
	goto L173
L173:
	;
	v711 = v705
	v724 = v689
	goto L174
L174:
	;
	if v724 == v707 {
		v928 = v677
		goto L121
	} else {
		goto L176
	}
L175:
	;
	v928 = v677
	goto L121
L176:
	;
	v731 = v711 + int32(4)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v737 = v731 + int32(base.Ui32(v732)>>(uint(int32(1))%32))*int32(12)
	if base.Ui32(v737) <= base.Ui32(v731) {
		v928 = v677
		goto L121
	} else {
		goto L177
	}
L177:
	;
	v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(528)+v724))))
	v748 = v731
	v749 = v737
	goto L178
L178:
	;
	v765 = int32(12)
	v766 = base.I32_div_s(v749-v748, v765)
	v771 = v748 + int32(base.Ui32(v766)>>(uint(int32(1))%32))*v765
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v771)))
	v774 = v772 & int32(255)
	if v742 == v774 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L175
L180:
	;
	v777 = v724 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v772) {
		v790 = v771
		v803 = v777
		goto L163
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v783 = base.B2i32(base.Ui32(v774) < base.Ui32(v742))
	if base.Ui32(v774) < base.Ui32(v742) {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v771)+8))
	if v780 != 0 {
		v711 = v780
		v724 = v777
		goto L174
	} else {
		goto L184
	}
L184:
	;
	v928 = v677
	goto L121
L185:
	;
	v784 = v771 + int32(12)
	goto L187
L186:
	;
	v784 = v748
	goto L187
L187:
	;
	if base.Ui32(v774) < base.Ui32(v742) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v785 = v749
	goto L190
L189:
	;
	v785 = v771
	goto L190
L190:
	;
	if base.Ui32(v784) < base.Ui32(v785) {
		v748 = v784
		v749 = v785
		goto L178
	} else {
		goto L191
	}
L191:
	;
	goto L179
L192:
	;
	v833 = v814 << (uint(int32(2)) % 32)
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v790)+4))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v833+v834)))
	v841 = F_CheckAffix(m, v24+int32(528), v668, v836, l2, v24+int32(16), v24+int32(12))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L22
	} else {
		goto L195
	}
L193:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v790)+8))
	if v923 != 0 {
		v676 = v923
		v677 = v914
		v689 = v803
		goto L161
	} else {
		goto L218
	}
L194:
	;
	v918 = v814 + int32(1)
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	if base.Ui32(v918) < base.Ui32(int32(base.Ui32(v919)>>(uint(int32(8))%32))) {
		v813 = v914
		v814 = v918
		goto L192
	} else {
		goto L217
	}
L195:
	;
	if v841 == int32(0) {
		v914 = v813
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v790)+4))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v848+v833)))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v850)+4))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v852+v539)))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+4))
	if v851&v855&int32(128) != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v860 = int32(733277)
	goto L199
L198:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v850)))
	v860 = v859
	goto L199
L199:
	;
	v861 = F_FindWord(m, l0, v24+int32(16), v860, l2)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L22
	} else {
		goto L200
	}
L200:
	;
	if v861 == int32(0) {
		v914 = v813
		goto L194
	} else {
		goto L201
	}
L201:
	;
	v865 = int32(0)
	if int32(4088) < v813-v99 {
		v910 = v865
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v914 = v813 + v910<<(uint(int32(2))%32)
	goto L194
L203:
	;
	if v813 != v99 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v871 = v24 + int32(16)
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v813-int32(4))))
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874))))
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871))))
	if v878 == int32(0) {
		v897 = v877
		v898 = v878
		goto L208
	} else {
		goto L209
	}
L205:
	;
	goto L206
L206:
	;
	v904 = F_pstrdup(m, v24+int32(16))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L22
	} else {
		goto L216
	}
L207:
	;
	if v898-v897 == int32(0) {
		v910 = v865
		goto L202
	} else {
		goto L215
	}
L208:
	;
	goto L207
L209:
	;
	if v877 != v878 {
		v897 = v877
		v898 = v878
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v882 = v871
	v883 = v874
	goto L211
L211:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v883)+1)))
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v882)+1)))
	if v887 == int32(0) {
		v897 = v886
		v898 = v887
		goto L208
	} else {
		goto L213
	}
L212:
	;
	v897 = v886
	v898 = v887
	goto L208
L213:
	;
	v890 = int32(1)
	if v886 == v887 {
		v882 = v882 + v890
		v883 = v883 + v890
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	goto L206
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v904
	*(*int32)(unsafe.Add(mBase, uint32(v813)+4)) = int32(0)
	v910 = int32(1)
	goto L202
L217:
	;
	goto L193
L218:
	;
	goto L162
L219:
	;
	goto L120
L220:
	;
	goto L89
L221:
	;
	F_pfree(m, v99)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L22
	} else {
		goto L222
	}
L222:
	;
	v985 = int32(0)
	goto L20
}
func F_NumRelids(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v3 = F_pull_varnos(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v8 = F_bms_del_members(m, v3, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = int32(0)
			if v8 == v10 {
				v45 = int32(0)
			} else {
				v17 = int32(1)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				if v18 <= v17 {
					v21 = v17
				} else {
					v21 = v18
				}
				v25 = int32(0)
				v27 = v10
				for {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8)+v25<<(uint(int32(2))%32))))
					if v33 != 0 {
						v36 = v27 + base.I32_popcnt(v33)
					} else {
						v36 = v27
					}
					v38 = v25 + int32(1)
					if v38 != v21 {
						v25 = v38
						v27 = v36
						continue
					} else {
						break
					}
					break
				}
				v45 = v36
			}
			F_bms_free(m, v8)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				return v45
			}
		}
	}
}
func F_nameconcatoid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
	v16 = F_pg_snprintf(m, v7+int32(16), int32(20), int32(38146), v7)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v9&int32(3) == int32(0) {
		v43 = v9
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if int32(64) <= v16+v76 {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v76 = v68 - v9
	goto L3
L5:
	;
	v47 = v43
	goto L14
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v27 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v76 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v32 = v9
	goto L10
L10:
	;
	v36 = v32 + int32(1)
	if v36&int32(3) == int32(0) {
		v43 = v36
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v68 = v36
	goto L4
L12:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v41 != 0 {
		v32 = v36
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v56 = int32(-2139062144)
	if (int32(16843008)-v53|v53)&v56 == v56 {
		v47 = v47 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v62 = v47
	goto L17
L16:
	;
	goto L15
L17:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v66 != 0 {
		v62 = v62 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v68 = v62
	goto L4
L19:
	;
	goto L18
L20:
	;
	v82 = F_pg_mbcliplen(m, v9, v76, int32(63)-v16)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v84 = v76
	goto L22
L22:
	;
	v86 = F_palloc0(m, int32(64))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v84 = v82
	goto L22
L24:
	;
	if v84 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v16 != 0 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v88 = F__emscripten_memcpy_bulkmem(m, v86, v9, v84)
	mBase = m.M
	v89 = v88
	goto L28
L27:
	;
	v89 = v86
	goto L28
L28:
	;
	goto L25
L29:
	;
	m.G0 = v7 + int32(48)
	return v89
L30:
	;
	v93 = F__emscripten_memcpy_bulkmem(m, v89+v84, v7+int32(16), v16)
	mBase = m.M
	goto L32
L31:
	;
	goto L32
L32:
	;
	goto L29
}
func F_namegttext(m *base.Module, l0 int32) int32 {
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
		return base.B2i32(int32(0) < v6)
	}
}
func F_namehashfast(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
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
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
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
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	if l0&int32(3) == int32(0) {
		v25 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v64 = v58 - int32(1636608432)
	if l0&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v58 = v50 - l0
	goto L1
L3:
	;
	v29 = v25
	goto L12
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v58 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v14 = l0
	goto L8
L8:
	;
	v18 = v14 + int32(1)
	if v18&int32(3) == int32(0) {
		v25 = v18
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v50 = v18
	goto L2
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v23 != 0 {
		v14 = v18
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v38 = int32(-2139062144)
	if (int32(16843008)-v35|v35)&v38 == v38 {
		v29 = v29 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v44 = v29
	goto L15
L14:
	;
	goto L13
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v48 != 0 {
		v44 = v44 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v50 = v44
	goto L2
L17:
	;
	goto L16
L18:
	;
	return v318 ^ v310 - base.I32_rotl(v318, int32(24))
L19:
	;
	v296 = int32(14)
	v298 = v292 ^ v293 - base.I32_rotl(v292, v296)
	v302 = v298 ^ v291 - base.I32_rotl(v298, int32(11))
	v306 = v302 ^ v292 - base.I32_rotl(v302, int32(25))
	v310 = v306 ^ v298 - base.I32_rotl(v306, int32(16))
	v314 = v310 ^ v302 - base.I32_rotl(v310, int32(4))
	v318 = v314 ^ v306 - base.I32_rotl(v314, v296)
	goto L18
L20:
	;
	switch v222 - int32(1) {
	case 0:
		v284 = v223
		v285 = v224
		v286 = v225
		goto L47
	case 1:
		v277 = v223
		v278 = v224
		v279 = v225
		goto L48
	case 2:
		v270 = v223
		v271 = v224
		v272 = v225
		goto L49
	case 3:
		v264 = v224
		v265 = v225
		goto L50
	case 4:
		v260 = v224
		v261 = v225
		goto L51
	case 5:
		v254 = v224
		v255 = v225
		goto L52
	case 6:
		v248 = v224
		v249 = v225
		goto L53
	case 7:
		v243 = v225
		goto L54
	case 8:
		v238 = v225
		goto L55
	case 9:
		v233 = v225
		goto L56
	case 10:
		goto L57
	default:
		v291 = v223
		v292 = v224
		v293 = v225
		goto L19
	}
L21:
	;
	v173 = l0
	v174 = v58
	v175 = v64
	v176 = v64
	v177 = v64
	goto L44
L22:
	;
	if base.Ui32(int32(11)) < base.Ui32(v58) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v58) < base.Ui32(int32(12)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v221 = l0
	v222 = v58
	v223 = v64
	v224 = v64
	v225 = v64
	goto L20
L26:
	;
	switch v120 - int32(1) {
	case 0:
		v170 = v121
		goto L33
	case 1:
		v165 = v121
		goto L34
	case 2:
		goto L35
	case 3:
		v158 = v122
		goto L36
	case 4:
		v155 = v122
		goto L37
	case 5:
		v150 = v122
		goto L38
	case 6:
		goto L39
	case 7:
		v141 = v123
		goto L40
	case 8:
		v136 = v123
		goto L41
	case 9:
		v131 = v123
		goto L42
	case 10:
		goto L43
	default:
		v291 = v121
		v292 = v122
		v293 = v123
		goto L19
	}
L27:
	;
	v119 = l0
	v120 = v58
	v121 = v64
	v122 = v64
	v123 = v64
	goto L26
L28:
	;
	goto L29
L29:
	;
	v71 = l0
	v72 = v58
	v73 = v64
	v74 = v64
	v75 = v64
	goto L30
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v78 = v77 + v74
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v82 = v81 + v75
	v84 = int32(4)
	v86 = v79 + v73 - v82 ^ base.I32_rotl(v82, v84)
	v90 = v78 - v86 ^ base.I32_rotl(v86, int32(6))
	v91 = v82 + v78
	v92 = v86 + v91
	v93 = v90 + v92
	v97 = v91 - v90 ^ base.I32_rotl(v90, int32(8))
	v101 = v92 - v97 ^ base.I32_rotl(v97, int32(16))
	v105 = v93 - v101 ^ base.I32_rotl(v101, int32(19))
	v106 = v97 + v93
	v107 = v101 + v106
	v108 = v105 + v107
	v112 = v106 - v105 ^ base.I32_rotl(v105, v84)
	v113 = int32(12)
	v114 = v71 + v113
	v116 = v72 - v113
	if base.Ui32(int32(11)) < base.Ui32(v116) {
		v71 = v114
		v72 = v116
		v73 = v107
		v74 = v108
		v75 = v112
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v119 = v114
	v120 = v116
	v121 = v107
	v122 = v108
	v123 = v112
	goto L26
L32:
	;
	goto L31
L33:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v291 = v170 + v171
	v292 = v122
	v293 = v123
	goto L19
L34:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	v170 = v166<<(uint(int32(8))%32) + v165
	goto L33
L35:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
	v165 = v161<<(uint(int32(16))%32) + v121
	goto L34
L36:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v291 = v159 + v121
	v292 = v158
	v293 = v123
	goto L19
L37:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
	v158 = v155 + v156
	goto L36
L38:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+5)))
	v155 = v151<<(uint(int32(8))%32) + v150
	goto L37
L39:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
	v150 = v146<<(uint(int32(16))%32) + v122
	goto L38
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v291 = v142 + v121
	v292 = v144 + v122
	v293 = v141
	goto L19
L41:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+8)))
	v141 = v137<<(uint(int32(8))%32) + v136
	goto L40
L42:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+9)))
	v136 = v132<<(uint(int32(16))%32) + v131
	goto L41
L43:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+10)))
	v131 = v127<<(uint(int32(24))%32) + v123
	goto L42
L44:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v180 = v179 + v176
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v184 = v183 + v177
	v186 = int32(4)
	v188 = v181 + v175 - v184 ^ base.I32_rotl(v184, v186)
	v192 = v180 - v188 ^ base.I32_rotl(v188, int32(6))
	v193 = v184 + v180
	v194 = v188 + v193
	v195 = v192 + v194
	v199 = v193 - v192 ^ base.I32_rotl(v192, int32(8))
	v203 = v194 - v199 ^ base.I32_rotl(v199, int32(16))
	v207 = v195 - v203 ^ base.I32_rotl(v203, int32(19))
	v208 = v199 + v195
	v209 = v203 + v208
	v210 = v207 + v209
	v214 = v208 - v207 ^ base.I32_rotl(v207, v186)
	v215 = int32(12)
	v216 = v173 + v215
	v218 = v174 - v215
	if base.Ui32(int32(11)) < base.Ui32(v218) {
		v173 = v216
		v174 = v218
		v175 = v209
		v176 = v210
		v177 = v214
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v221 = v216
	v222 = v218
	v223 = v209
	v224 = v210
	v225 = v214
	goto L20
L46:
	;
	goto L45
L47:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	v291 = v284 + v287
	v292 = v285
	v293 = v286
	goto L19
L48:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	v284 = v280<<(uint(int32(8))%32) + v277
	v285 = v278
	v286 = v279
	goto L47
L49:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+2)))
	v277 = v273<<(uint(int32(16))%32) + v270
	v278 = v271
	v279 = v272
	goto L48
L50:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+3)))
	v270 = v266<<(uint(int32(24))%32) + v223
	v271 = v264
	v272 = v265
	goto L49
L51:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+4)))
	v264 = v260 + v262
	v265 = v261
	goto L50
L52:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+5)))
	v260 = v256<<(uint(int32(8))%32) + v254
	v261 = v255
	goto L51
L53:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+6)))
	v254 = v250<<(uint(int32(16))%32) + v248
	v255 = v249
	goto L52
L54:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+7)))
	v248 = v244<<(uint(int32(24))%32) + v224
	v249 = v243
	goto L53
L55:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+8)))
	v243 = v239<<(uint(int32(8))%32) + v238
	goto L54
L56:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+9)))
	v238 = v234<<(uint(int32(16))%32) + v233
	goto L55
L57:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+10)))
	v233 = v229<<(uint(int32(24))%32) + v225
	goto L56
}
func F_nameout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pstrdup(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_newcolor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v7 == int32(0) {
		v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
		if v10 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v14 = v11 + v10*int32(24)
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v15)
			v92 = v14
			*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = int32(0)
			v98 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v92)+12)) = v98
			v100 = int32(65535)
			*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)) = uint16(v100)
			*(*int64)(unsafe.Add(mBase, uint32(v92))) = v98
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v107 = base.I32_div_s(v92-v104, int32(24))
			return base.I32_extend16_s(v107)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v17) < base.Ui32(v18-int32(1)) {
				v23 = v17 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v92 = v25 + v23*int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = int32(0)
				v98 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v92)+12)) = v98
				v100 = int32(65535)
				*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)) = uint16(v100)
				*(*int64)(unsafe.Add(mBase, uint32(v92))) = v98
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v107 = base.I32_div_s(v92-v104, int32(24))
				return base.I32_extend16_s(v107)
			} else {
				if v17 == int32(32767) {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(101)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					if v34 != 0 {
						v36 = v34
					} else {
						v36 = int32(20)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v36
					return int32(-1)
				} else {
					v40 = int32(32768)
					v42 = v18 << (uint(int32(1)) % 32)
					if base.Ui32(v40) <= base.Ui32(v42) {
						v45 = v40
					} else {
						v45 = v42
					}
					v47 = v45 * int32(24)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v50 = l0 + int32(108)
					if v48 == v50 {
						v53 = F_palloc_extended(m, v47, int32(2))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v53 == int32(0) {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v67)+24)) = int32(101)
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
								if v71 != 0 {
									v73 = v71
								} else {
									v73 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v73
								return int32(-1)
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v61 = v59 * int32(24)
								if v61 != 0 {
									v62 = F__emscripten_memcpy_bulkmem(m, v53, v50, v61)
									mBase = m.M
								} else {
								}
								v82 = v53
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v82
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v87 = v85 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v87
								v92 = v82 + v87*int32(24)
								*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = int32(0)
								v98 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v92)+12)) = v98
								v100 = int32(65535)
								*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)) = uint16(v100)
								*(*int64)(unsafe.Add(mBase, uint32(v92))) = v98
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v107 = base.I32_div_s(v92-v104, int32(24))
								return base.I32_extend16_s(v107)
							}
						}
					} else {
						v64 = F_repalloc_extended(m, v48, v47)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							if v64 != 0 {
								v82 = v64
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v82
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v87 = v85 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v87
								v92 = v82 + v87*int32(24)
								*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = int32(0)
								v98 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v92)+12)) = v98
								v100 = int32(65535)
								*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)) = uint16(v100)
								*(*int64)(unsafe.Add(mBase, uint32(v92))) = v98
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v107 = base.I32_div_s(v92-v104, int32(24))
								return base.I32_extend16_s(v107)
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v67)+24)) = int32(101)
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
								if v71 != 0 {
									v73 = v71
								} else {
									v73 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v73
								return int32(-1)
							}
						}
					}
				}
			}
		}
	} else {
		return int32(-1)
	}
}
func F_newdfa(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
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
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = int32(base.Ui32(v9+int32(31)) >> (uint(int32(5)) % 32))
	v15 = v9 << (uint(int32(1)) % 32)
	if base.Ui32(int32(20)) < base.Ui32(v15) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = int32(0)
	if v137&int32(32) != 0 {
		goto L51
	} else {
		goto L52
	}
L2:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+69)) = uint8(v113)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+68)) = uint8(base.B2i32(l3 == v113))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+36)) = v112 + int32(3916)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+32)) = v112 + int32(1516)
	v125 = v112 + int32(1352)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+24)) = v125
	v128 = v112 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v125 + v15<<(uint(int32(2))%32)
	v135 = v112
	v136 = v128
	goto L1
L3:
	;
	v35 = F_palloc_extended(m, int32(72), int32(2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L13
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(15) < v18 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if l3 != 0 {
		v112 = l3
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v23 = F_palloc_extended(m, int32(8716), int32(2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	if v23 != 0 {
		v112 = v23
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = v27
	goto L12
L11:
	;
	v29 = int32(12)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v29
	return int32(0)
L13:
	;
	if v35 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v39 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v48 = F_palloc_extended(m, v9<<(uint(int32(6))%32), int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L20
	}
L17:
	;
	v41 = v39
	goto L19
L18:
	;
	v41 = int32(12)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v41
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v48
	v54 = int32(2)
	v57 = F_palloc_extended(m, v13*(v15|int32(1))<<(uint(v54)%32), v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v57
	v61 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v57 + v15*v13<<(uint(v61)%32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v70 = F_palloc_extended(m, v9*v65<<(uint(int32(3))%32), v61)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v78 = F_palloc_extended(m, v9*v73<<(uint(int32(4))%32), int32(2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v80 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+68)) = uint16(v80)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+36)) = v78
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if v83 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	if v84 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	if v92 != 0 {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	F_pfree(m, v83)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L31
	}
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	if v87 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if v78 != 0 {
		v135 = v35
		v136 = v83
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	goto L26
L32:
	;
	F_pfree(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	if v95 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_pfree(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v35)+36))
	if v98 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_pfree(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+68)))
	if v101 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_pfree(m, v35)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v106 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v108 = v106
	goto L50
L49:
	;
	v108 = int32(12)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v108
	return int32(0)
L51:
	;
	v143 = int32(7)
	goto L53
L52:
	;
	v143 = v15
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+56)) = v136
	*(*int64)(unsafe.Add(mBase, uint32(v135)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v135)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v135)+16)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v135)+60)) = int64(4294967295)
	return v135
}
func F_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v500 int32
	_ = v500
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v718 int32
	_ = v718
	var v738 int32
	_ = v738
	var v758 int32
	_ = v758
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v858 int32
	_ = v858
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1194 int32
	_ = v1194
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1306 int32
	_ = v1306
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 != 0 {
		v1306 = v2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1312 != int32(91) {
		goto L385
	} else {
		goto L386
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	return int32(1)
L3:
	;
	return v1306
L4:
	;
	goto L17
L5:
	;
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v1234&int32(2) == int32(0) {
		goto L369
	} else {
		goto L370
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L7:
	;
	if base.Ui32(v158) < base.Ui32(v124) {
		goto L5
	} else {
		goto L365
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(36)
	return int32(1)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(94)
	return int32(1)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(46)
	return int32(1)
L11:
	;
	if v124-v158 < int32(21) {
		goto L351
	} else {
		goto L352
	}
L12:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1140 == int32(40) {
		goto L348
	} else {
		goto L349
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967336)
	return int32(1)
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1130 != 0 {
		goto L345
	} else {
		goto L346
	}
L15:
	;
	if base.Ui32(v1020) <= base.Ui32(v1031) {
		goto L331
	} else {
		goto L332
	}
L16:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1076)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+8)) = v1077 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(8589934668)
	return int32(1)
L17:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v16 != int32(110) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1067)+8)) = v1068 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(12884901964)
	return int32(1)
L19:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v18&int32(32) == int32(0) {
		v122 = v29
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v18&int32(1024) == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(65)
	return int32(1)
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v123) < base.Ui32(v124) {
		goto L54
	} else {
		goto L55
	}
L23:
	;
	if base.Ui32(int32(5)) < base.Ui32(v29) {
		v122 = v29
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if int32(1)<<(uint(v29)%32)&int32(54) == int32(0) {
		v122 = v29
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = v47
	v51 = v46
	goto L27
L26:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v122 = v121
	goto L22
L27:
	;
	if base.Ui32(v51) <= base.Ui32(v49) {
		v94 = v49
		v96 = v51
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if base.Ui32(v94) < base.Ui32(v96) {
		goto L43
	} else {
		goto L44
	}
L30:
	;
	v55 = v49
	v57 = v51
	goto L31
L31:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v61 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	switch v61 - int32(1) {
	case 0:
		goto L36
	case 1:
		goto L35
	case 2:
		goto L34
	default:
		goto L37
	}
L32:
	;
	v94 = v90
	v96 = v85
	goto L29
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v84 == int32(0) {
		v94 = v86
		v96 = v85
		goto L29
	} else {
		goto L40
	}
L34:
	;
	if base.Ui32(int32(255)) < base.Ui32(v59) {
		v94 = v55
		v96 = v57
		goto L29
	} else {
		goto L39
	}
L35:
	;
	v75 = F_iswspace(m, v59)
	mBase = m.M
	v84 = v75
	goto L33
L36:
	;
	v71 = F_pg_u_isspace(m, v59)
	mBase = m.M
	v84 = v71
	goto L33
L37:
	;
	if base.Ui32(int32(127)) < base.Ui32(v59) {
		v94 = v55
		v96 = v57
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+uint32(_consts[653]))))
	v84 = int32(base.Ui32(v68) >> (uint(int32(7)) % 32))
	goto L33
L39:
	;
	v81 = F___isspace(m, v59)
	mBase = m.M
	v84 = base.B2i32(v81 != int32(0))
	goto L33
L40:
	;
	v90 = v86 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
	if base.Ui32(v90) < base.Ui32(v85) {
		v55 = v90
		v57 = v85
		goto L31
	} else {
		goto L41
	}
L41:
	;
	goto L32
L42:
	;
	v110 = v94
	goto L50
L43:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v99 == int32(35) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v94 != v47 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v104 | int32(128)
	goto L49
L48:
	;
	goto L49
L49:
	;
	goto L26
L50:
	;
	v115 = v110 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115
	if base.Ui32(v96) <= base.Ui32(v115) {
		v49 = v115
		v51 = v96
		goto L27
	} else {
		goto L52
	}
L51:
	;
	v49 = v115
	v51 = v96
	goto L27
L52:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v118 != int32(10) {
		v110 = v115
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v158 = v123 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	switch v122 - int32(2) {
	case 0:
		goto L75
	case 1:
		goto L74
	case 2, 3:
		goto L73
	case 4:
		goto L72
	case 5:
		goto L71
	case 6:
		goto L70
	case 7:
		goto L69
	default:
		goto L68
	}
L55:
	;
	if base.Ui32(int32(9)) < base.Ui32(v122) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v129 = int32(1) << (uint(v122) % 32)
	if v129&int32(960) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v129&int32(14) != 0 {
		goto L2
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v150 != 0 {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	if v129&int32(48) == int32(0) {
		goto L54
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v142 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v144 = v142
	goto L64
L63:
	;
	v144 = int32(9)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v144
	return int32(0)
L65:
	;
	v152 = v150
	goto L67
L66:
	;
	v152 = int32(7)
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v152
	return int32(0)
L68:
	;
	if v160 != int32(40) {
		goto L246
	} else {
		goto L247
	}
L69:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L242
	} else {
		goto L243
	}
L70:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L238
	} else {
		goto L239
	}
L71:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L234
	} else {
		goto L235
	}
L72:
	;
	switch v160 - int32(91) {
	case 0:
		goto L198
	case 1:
		goto L199
	case 2:
		goto L200
	default:
		goto L197
	}
L73:
	;
	v500 = v160 - int32(48)
	if base.Ui32(int32(10)) <= base.Ui32(v500) {
		goto L171
	} else {
		goto L172
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L75:
	;
	switch v160 - int32(36) {
	case 0:
		goto L78
	default:
		goto L77
	case 6:
		goto L82
	case 10:
		goto L80
	case 55:
		goto L81
	case 56:
		goto L76
	case 58:
		goto L79
	}
L76:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L144
	} else {
		goto L145
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L78:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v259&int32(32) != 0 {
		goto L106
	} else {
		goto L107
	}
L79:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v237 != int32(40) {
		goto L102
	} else {
		goto L103
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(46)
	return int32(1)
L81:
	;
	if v124-v158 < int32(21) {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	switch v165 - int32(94) {
	case 0, 16:
		goto L84
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L83
	default:
		goto L85
	}
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967338)
	return int32(1)
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(180388626544)
	return int32(1)
L85:
	;
	if v165 != int32(40) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L98
	} else {
		goto L99
	}
L88:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v181 != int32(91) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	if v184 != int32(58) {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	switch v187 - int32(60) {
	case 0, 2:
		goto L91
	default:
		goto L87
	}
L91:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	if v190 != int32(58) {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v123)+20))
	if v193 != int32(93) {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v123)+24))
	if v196 != int32(93) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(28)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v203 | int32(128)
	v207 = int32(60)
	if v187 == v207 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v211 = v207
	goto L97
L96:
	;
	v211 = int32(62)
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v211
	return int32(1)
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967387)
	return int32(1)
L99:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v219 != int32(94) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(91)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	return int32(1)
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(403726925936)
	return int32(1)
L102:
	;
	if v237 != int32(110) {
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v247 | int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(94)
	return int32(1)
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(94)
	return int32(1)
L106:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v269 = v267
	v271 = v266
	goto L110
L107:
	;
	v343 = v158
	v344 = v124
	goto L108
L108:
	;
	if base.Ui32(v344) <= base.Ui32(v343) {
		goto L137
	} else {
		goto L138
	}
L109:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v343 = v341
	v344 = v342
	goto L108
L110:
	;
	if base.Ui32(v271) <= base.Ui32(v269) {
		v314 = v269
		v316 = v271
		goto L112
	} else {
		goto L113
	}
L112:
	;
	if base.Ui32(v314) < base.Ui32(v316) {
		goto L126
	} else {
		goto L127
	}
L113:
	;
	v275 = v269
	v277 = v271
	goto L114
L114:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v281 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	switch v281 - int32(1) {
	case 0:
		goto L119
	case 1:
		goto L118
	case 2:
		goto L117
	default:
		goto L120
	}
L115:
	;
	v314 = v310
	v316 = v305
	goto L112
L116:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v304 == int32(0) {
		v314 = v306
		v316 = v305
		goto L112
	} else {
		goto L123
	}
L117:
	;
	if base.Ui32(int32(255)) < base.Ui32(v279) {
		v314 = v275
		v316 = v277
		goto L112
	} else {
		goto L122
	}
L118:
	;
	v295 = F_iswspace(m, v279)
	mBase = m.M
	v304 = v295
	goto L116
L119:
	;
	v291 = F_pg_u_isspace(m, v279)
	mBase = m.M
	v304 = v291
	goto L116
L120:
	;
	if base.Ui32(int32(127)) < base.Ui32(v279) {
		v314 = v275
		v316 = v277
		goto L112
	} else {
		goto L121
	}
L121:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+uint32(_consts[653]))))
	v304 = int32(base.Ui32(v288) >> (uint(int32(7)) % 32))
	goto L116
L122:
	;
	v301 = F___isspace(m, v279)
	mBase = m.M
	v304 = base.B2i32(v301 != int32(0))
	goto L116
L123:
	;
	v310 = v306 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v310
	if base.Ui32(v310) < base.Ui32(v305) {
		v275 = v310
		v277 = v305
		goto L114
	} else {
		goto L124
	}
L124:
	;
	goto L115
L125:
	;
	v330 = v314
	goto L133
L126:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	if v319 == int32(35) {
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	if v314 != v267 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = v324 | int32(128)
	goto L132
L131:
	;
	goto L132
L132:
	;
	goto L109
L133:
	;
	v335 = v330 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
	if base.Ui32(v316) <= base.Ui32(v335) {
		v269 = v335
		v271 = v316
		goto L110
	} else {
		goto L135
	}
L134:
	;
	v269 = v335
	v271 = v316
	goto L110
L135:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	if v338 != int32(10) {
		v330 = v335
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(36)
	return int32(1)
L138:
	;
	goto L139
L139:
	;
	if v344-v343 < int32(5) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(154618822768)
	return int32(1)
L141:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	if v353 != int32(92) {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	if v356 != int32(41) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v359)+8)) = v360 | int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(36)
	return int32(1)
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v380 != 0 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	switch v389 - int32(40) {
	case 0:
		goto L155
	case 1:
		goto L154
	default:
		goto L150
	case 9, 10, 11, 12, 13, 14, 15, 16, 17:
		goto L151
	case 20:
		goto L153
	case 22:
		goto L152
	case 83:
		goto L156
	}
L147:
	;
	v382 = v380
	goto L149
L148:
	;
	v382 = int32(5)
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v382
	return int32(0)
L150:
	;
	v441 = int32(0)
	v443 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	switch v443 - int32(1) {
	case 0:
		goto L161
	case 1:
		goto L160
	case 2:
		goto L159
	default:
		goto L162
	}
L151:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
	v431 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v429)+8)) = v430 | v431
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v389 - int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(98)
	return v431
L152:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v420)+8)) = v421 | int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(62)
	return int32(1)
L153:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v411)+8)) = v412 | int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(60)
	return int32(1)
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(176093659177)
	return int32(1)
L155:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967336)
	return int32(1)
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(5)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+8)) = v395 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(123)
	return int32(1)
L157:
	;
	if v477 != 0 {
		goto L165
	} else {
		goto L166
	}
L158:
	;
	v477 = v475
	goto L157
L159:
	;
	if base.Ui32(int32(255)) < base.Ui32(v389) {
		v475 = v441
		goto L158
	} else {
		goto L164
	}
L160:
	;
	v466 = F_iswalnum(m, v389)
	mBase = m.M
	v477 = v466
	goto L157
L161:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _consts[654]))
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+16)))
	v462 = F_pg_u_isalnum(m, v389, (v457^int32(-1))&int32(1))
	mBase = m.M
	v477 = v462
	goto L157
L162:
	;
	if base.Ui32(int32(127)) < base.Ui32(v389) {
		v475 = v441
		goto L158
	} else {
		goto L163
	}
L163:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+uint32(_consts[653]))))
	v477 = base.B2i32(v450&int32(3) != int32(0))
	goto L157
L164:
	;
	v472 = F_isalnum(m, v389)
	mBase = m.M
	v475 = base.B2i32(v472 != int32(0))
	goto L158
L165:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v478)+8)) = v479 | int32(16)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+8)) = v484 | int32(256)
	goto L167
L166:
	;
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L168:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L190
	} else {
		goto L191
	}
L169:
	;
	if v122 == int32(4) {
		goto L180
	} else {
		goto L181
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(44)
	return int32(1)
L171:
	;
	if v160 == int32(44) {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v500
	return int32(1)
L174:
	;
	if v160 == int32(92) {
		goto L168
	} else {
		goto L175
	}
L175:
	;
	if v160 == int32(125) {
		goto L169
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v511 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v513 = v511
	goto L179
L178:
	;
	v513 = int32(10)
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v513
	return int32(0)
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v557 != 0 {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967421)
	return int32(1)
L184:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v531&int32(2) == int32(0) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v536 != int32(63) {
		goto L183
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v542)+8)) = v543 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(125)
	return int32(1)
L187:
	;
	v559 = v557
	goto L189
L188:
	;
	v559 = int32(10)
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v559
	return int32(0)
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v580 != 0 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	if v122 != int32(5) {
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v566 != int32(125) {
		goto L190
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967421)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	return int32(1)
L194:
	;
	v582 = v580
	goto L196
L195:
	;
	v582 = int32(10)
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v582
	return int32(0)
L197:
	;
	if v160 == int32(45) {
		goto L1
	} else {
		goto L233
	}
L198:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L223
	} else {
		goto L224
	}
L199:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v605)+8)) = v606 | int32(64)
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v610&int32(2) == int32(0) {
		goto L204
	} else {
		goto L205
	}
L200:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v588 == int32(91) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(399431958640)
	return int32(1)
L202:
	;
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(93)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v599 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(2) - v598&v599
	return v599
L204:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(395136991344)
	return int32(1)
L205:
	;
	goto L206
L206:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v619)+8)) = v620 | int32(128)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v625) <= base.Ui32(v624) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v629 != 0 {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	goto L209
L209:
	;
	v635 = F_lexescape(m, l0)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v631 = v629
	goto L212
L211:
	;
	v631 = int32(5)
	goto L212
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v631
	return int32(0)
L213:
	;
	return int32(0)
L214:
	;
	if v635 == int32(0) {
		v1306 = v2
		goto L3
	} else {
		goto L215
	}
L215:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v643 = v641 - int32(99)
	if base.Ui32(v643) <= base.Ui32(int32(16)) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v646 = int32(1)
	if v646<<(uint(v643)%32)&int32(73729) != 0 {
		v1306 = v646
		goto L3
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v654 != 0 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	goto L218
L220:
	;
	v656 = v654
	goto L222
L221:
	;
	v656 = int32(5)
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v656
	return int32(0)
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v663 != 0 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	switch v672 - int32(46) {
	case 0:
		goto L232
	default:
		goto L229
	case 12:
		goto L230
	case 15:
		goto L231
	}
L226:
	;
	v665 = v663
	goto L228
L227:
	;
	v665 = int32(7)
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v665
	return int32(0)
L229:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(390842024048)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v158
	return int32(1)
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(9)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v694)+8)) = v695 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(67)
	return int32(1)
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(8)
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v683)+8)) = v684 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(69)
	return int32(1)
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(73)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(7)
	return int32(1)
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L235:
	;
	if v160 != int32(46) {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v718 != int32(93) {
		goto L234
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(197568495704)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	return int32(1)
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L239:
	;
	if v160 != int32(61) {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v738 != int32(93) {
		goto L238
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(261993005144)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	return int32(1)
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L243:
	;
	if v160 != int32(58) {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v758 != int32(93) {
		goto L242
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(249108103256)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	return int32(1)
L246:
	;
	switch v160 - int32(36) {
	case 0:
		goto L8
	default:
		goto L6
	case 5:
		goto L12
	case 6:
		goto L252
	case 7:
		goto L251
	case 10:
		goto L10
	case 27:
		goto L250
	case 55:
		goto L11
	case 56:
		goto L7
	case 58:
		goto L9
	case 87:
		goto L249
	case 88:
		goto L253
	}
L247:
	;
	goto L248
L248:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L13
	} else {
		goto L310
	}
L249:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v858&int32(32) != 0 {
		goto L267
	} else {
		goto L268
	}
L250:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L262
	} else {
		goto L263
	}
L251:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L258
	} else {
		goto L259
	}
L252:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(124)
	return int32(1)
L254:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967338)
	return int32(1)
L255:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v784&int32(2) == int32(0) {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v789 != int32(63) {
		goto L254
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v795)+8)) = v796 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(42)
	return int32(1)
L258:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967339)
	return int32(1)
L259:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v809&int32(2) == int32(0) {
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v814 != int32(63) {
		goto L258
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v820)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+8)) = v821 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(43)
	return int32(1)
L262:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967359)
	return int32(1)
L263:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v834&int32(2) == int32(0) {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v839 != int32(63) {
		goto L262
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v845)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v845)+8)) = v846 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(63)
	return int32(1)
L266:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v991)+8))
	v993 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v991)+8)) = v992 | v993
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(123)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v993
	return int32(1)
L267:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v868 = v866
	v870 = v865
	goto L271
L268:
	;
	v942 = v158
	v943 = v124
	goto L269
L269:
	;
	if base.Ui32(v942) < base.Ui32(v943) {
		goto L298
	} else {
		goto L299
	}
L270:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v942 = v940
	v943 = v941
	goto L269
L271:
	;
	if base.Ui32(v870) <= base.Ui32(v868) {
		v913 = v868
		v915 = v870
		goto L273
	} else {
		goto L274
	}
L273:
	;
	if base.Ui32(v913) < base.Ui32(v915) {
		goto L287
	} else {
		goto L288
	}
L274:
	;
	v874 = v868
	v876 = v870
	goto L275
L275:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v874)))
	v880 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	switch v880 - int32(1) {
	case 0:
		goto L280
	case 1:
		goto L279
	case 2:
		goto L278
	default:
		goto L281
	}
L276:
	;
	v913 = v909
	v915 = v904
	goto L273
L277:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v903 == int32(0) {
		v913 = v905
		v915 = v904
		goto L273
	} else {
		goto L284
	}
L278:
	;
	if base.Ui32(int32(255)) < base.Ui32(v878) {
		v913 = v874
		v915 = v876
		goto L273
	} else {
		goto L283
	}
L279:
	;
	v894 = F_iswspace(m, v878)
	mBase = m.M
	v903 = v894
	goto L277
L280:
	;
	v890 = F_pg_u_isspace(m, v878)
	mBase = m.M
	v903 = v890
	goto L277
L281:
	;
	if base.Ui32(int32(127)) < base.Ui32(v878) {
		v913 = v874
		v915 = v876
		goto L273
	} else {
		goto L282
	}
L282:
	;
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878)+uint32(_consts[653]))))
	v903 = int32(base.Ui32(v887) >> (uint(int32(7)) % 32))
	goto L277
L283:
	;
	v900 = F___isspace(m, v878)
	mBase = m.M
	v903 = base.B2i32(v900 != int32(0))
	goto L277
L284:
	;
	v909 = v905 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v909
	if base.Ui32(v909) < base.Ui32(v904) {
		v874 = v909
		v876 = v904
		goto L275
	} else {
		goto L285
	}
L285:
	;
	goto L276
L286:
	;
	v929 = v913
	goto L294
L287:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v913)))
	if v918 == int32(35) {
		goto L286
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	if v913 != v866 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	goto L289
L291:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v922)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v922)+8)) = v923 | int32(128)
	goto L293
L292:
	;
	goto L293
L293:
	;
	goto L270
L294:
	;
	v934 = v929 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v934
	if base.Ui32(v915) <= base.Ui32(v934) {
		v868 = v934
		v870 = v915
		goto L271
	} else {
		goto L296
	}
L295:
	;
	v868 = v934
	v870 = v915
	goto L271
L296:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v934)))
	if v937 != int32(10) {
		v929 = v934
		goto L294
	} else {
		goto L297
	}
L297:
	;
	goto L295
L298:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v942)))
	v947 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	switch v947 - int32(1) {
	case 0:
		goto L304
	case 1:
		goto L303
	case 2:
		goto L302
	default:
		goto L305
	}
L299:
	;
	goto L300
L300:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v977)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v977)+8)) = v978 | int32(8)
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v982)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v982)+8)) = v983 | int32(256)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(528280977520)
	return int32(1)
L301:
	;
	if v976 != 0 {
		goto L266
	} else {
		goto L309
	}
L302:
	;
	if base.Ui32(v945) <= base.Ui32(int32(255)) {
		goto L306
	} else {
		goto L307
	}
L303:
	;
	v965 = F_isdigit(m, v945)
	mBase = m.M
	v976 = v965
	goto L301
L304:
	;
	v955 = *(*int32)(unsafe.Add(mBase, _consts[654]))
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+16)))
	v961 = F_pg_u_isdigit(m, v945, (v956^int32(-1))&int32(1))
	mBase = m.M
	v976 = v961
	goto L301
L305:
	;
	v976 = base.B2i32(base.Ui32(v945-int32(48)) < base.Ui32(int32(10)))
	goto L301
L306:
	;
	v971 = F_isdigit(m, v945)
	mBase = m.M
	v975 = base.B2i32(v971 != int32(0))
	goto L308
L307:
	;
	v975 = int32(0)
	goto L308
L308:
	;
	v976 = v975
	goto L301
L309:
	;
	goto L300
L310:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1003&int32(2) == int32(0) {
		goto L13
	} else {
		goto L311
	}
L311:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v1008 != int32(63) {
		goto L13
	} else {
		goto L312
	}
L312:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+8)) = v1012 | int32(128)
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1018 = v1016 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1018
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v1020) <= base.Ui32(v1018) {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1024 != 0 {
		goto L316
	} else {
		goto L317
	}
L314:
	;
	goto L315
L315:
	;
	v1031 = v1016 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1031
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1018)))
	if v1033 != int32(35) {
		goto L320
	} else {
		goto L321
	}
L316:
	;
	v1026 = v1024
	goto L318
L317:
	;
	v1026 = int32(13)
	goto L318
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1026
	return int32(0)
L319:
	;
	goto L18
L320:
	;
	switch v1033 - int32(33) {
	case 0:
		goto L16
	default:
		goto L14
	case 25:
		goto L323
	case 27:
		goto L15
	case 28:
		goto L319
	}
L321:
	;
	goto L322
L322:
	;
	if base.Ui32(v1020) <= base.Ui32(v1031) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(40)
	return int32(1)
L324:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1064 == int32(0) {
		goto L17
	} else {
		goto L330
	}
L325:
	;
	v1044 = v1031
	goto L326
L326:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1044)))
	v1052 = v1044 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1052
	if v1050 == int32(41) {
		goto L324
	} else {
		goto L328
	}
L327:
	;
	goto L324
L328:
	;
	if base.Ui32(v1052) < base.Ui32(v1020) {
		v1044 = v1052
		goto L326
	} else {
		goto L329
	}
L329:
	;
	goto L327
L330:
	;
	v1306 = v2
	goto L3
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1088 != 0 {
		goto L334
	} else {
		goto L335
	}
L332:
	;
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1016 + int32(12)
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+8))
	if v1097 != int32(33) {
		goto L338
	} else {
		goto L339
	}
L334:
	;
	v1090 = v1088
	goto L336
L335:
	;
	v1090 = int32(13)
	goto L336
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1090
	return int32(0)
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1122 != 0 {
		goto L342
	} else {
		goto L343
	}
L338:
	;
	if v1097 != int32(61) {
		goto L337
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1111)+8)) = v1112 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(76)
	return int32(1)
L341:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1102)+8)) = v1103 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967372)
	return int32(1)
L342:
	;
	v1124 = v1122
	goto L344
L343:
	;
	v1124 = int32(13)
	goto L344
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1124
	return int32(0)
L345:
	;
	v1132 = v1130
	goto L347
L346:
	;
	v1132 = int32(13)
	goto L347
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1132
	return int32(0)
L348:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1143)+8)) = v1144 | int32(256)
	goto L350
L349:
	;
	goto L350
L350:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(176093659177)
	return int32(1)
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L362
	} else {
		goto L363
	}
L352:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v1156 != int32(91) {
		goto L351
	} else {
		goto L353
	}
L353:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	if v1159 != int32(58) {
		goto L351
	} else {
		goto L354
	}
L354:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	switch v1162 - int32(60) {
	case 0, 2:
		goto L355
	default:
		goto L351
	}
L355:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	if v1165 != int32(58) {
		goto L351
	} else {
		goto L356
	}
L356:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v123)+20))
	if v1168 != int32(93) {
		goto L351
	} else {
		goto L357
	}
L357:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v123)+24))
	if v1171 != int32(93) {
		goto L351
	} else {
		goto L358
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(28)
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1177)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1177)+8)) = v1178 | int32(128)
	v1182 = int32(60)
	if v1162 == v1182 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1186 = v1182
	goto L361
L360:
	;
	v1186 = int32(62)
	goto L361
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1186
	return int32(1)
L362:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967387)
	return int32(1)
L363:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v1194 != int32(94) {
		goto L362
	} else {
		goto L364
	}
L364:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(91)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(8)
	return int32(1)
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1223 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1225 = v1223
	goto L368
L367:
	;
	v1225 = int32(5)
	goto L368
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1225
	return int32(0)
L369:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v1240 = int32(0)
	v1242 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	switch v1242 - int32(1) {
	case 0:
		goto L376
	case 1:
		goto L375
	case 2:
		goto L374
	default:
		goto L377
	}
L370:
	;
	goto L371
L371:
	;
	v1298 = F_lexescape(m, l0)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L213
	} else {
		goto L383
	}
L372:
	;
	if v1276 != 0 {
		goto L380
	} else {
		goto L381
	}
L373:
	;
	v1276 = v1274
	goto L372
L374:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1239) {
		v1274 = v1240
		goto L373
	} else {
		goto L379
	}
L375:
	;
	v1265 = F_iswalnum(m, v1239)
	mBase = m.M
	v1276 = v1265
	goto L372
L376:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, _consts[654]))
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1255)+16)))
	v1261 = F_pg_u_isalnum(m, v1239, (v1256^int32(-1))&int32(1))
	mBase = m.M
	v1276 = v1261
	goto L372
L377:
	;
	if base.Ui32(int32(127)) < base.Ui32(v1239) {
		v1274 = v1240
		goto L373
	} else {
		goto L378
	}
L378:
	;
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+uint32(_consts[653]))))
	v1276 = base.B2i32(v1249&int32(3) != int32(0))
	goto L372
L379:
	;
	v1271 = F_isalnum(m, v1239)
	mBase = m.M
	v1274 = base.B2i32(v1271 != int32(0))
	goto L373
L380:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+8)) = v1278 | int32(16)
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1282)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1282)+8)) = v1283 | int32(256)
	goto L382
L381:
	;
	goto L382
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1290 + int32(4)
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1290)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1294
	return int32(1)
L383:
	;
	v1306 = v1298
	goto L3
L384:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(193273528402)
	return int32(1)
L385:
	;
	if base.Ui32(v124) <= base.Ui32(v158) {
		goto L384
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(193273528432)
	return int32(1)
L388:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v1316 != int32(93) {
		goto L384
	} else {
		goto L389
	}
L389:
	;
	goto L387
}
func F_nfanode(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v9 = F_newnfa(m, l0, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v13 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	F_dupnfa(m, v9, v16, v17, v18, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v105 = v4
	goto L5
L5:
	;
	return v105
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v25 != 0 {
		v32 = v25
		v33 = v4
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	F_specialcolors(m, v9)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v28 != 0 {
		v32 = v28
		v33 = v4
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v29 = F_optimize(m, v9)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = v31
	v33 = v29
	goto L7
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v42 != 0 {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	if v32 != 0 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v37 = v32
	goto L15
L15:
	;
	if v37 != 0 {
		goto L12
	} else {
		goto L18
	}
L16:
	;
	F_makesearch(m, l0, v9)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v37 = v36
	goto L15
L18:
	;
	F_compact(m, v9, l1+int32(36))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	v43 = v42
	goto L23
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v69 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+136))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+136)) = v51 + v52*int32(-36) - int32(8)
	F_pfree(m, v43)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	if v49 != 0 {
		v43 = v49
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v70 = v69
	goto L30
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(0)
	F_pfree(m, v9)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L34
	}
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+136))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+136)) = v78 + v79*int32(-40) - int32(8)
	F_pfree(m, v70)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	if v76 != 0 {
		v70 = v76
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v105 = v33
	goto L5
}
func F_nfatree(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	v17 = F_nfanode(m, l0, l1, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L9
	}
L4:
	;
	v8 = F_nfatree(m, l0, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	if v12 != 0 {
		v7 = v12
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	return v17
}
func F_normalize_exec_path(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
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
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v523 int32
	_ = v523
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v602 int32
	_ = v602
	var v611 int32
	_ = v611
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v876 int32
	_ = v876
	var v912 int32
	_ = v912
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = m.G0
	v20 = v18 - int32(8208)
	m.G0 = v20
	if l0 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v20 + int32(8208)
	if v912 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L2:
	;
	v912 = int32(0)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v27 = int32(4097)
	v30 = F_memchr(m, l0, int32(0), v27)
	mBase = m.M
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v32 = v30 - l0
	goto L9
L8:
	;
	v32 = v27
	goto L9
L9:
	;
	goto L6
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(44)
	goto L2
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(int32(4095)) < base.Ui32(v32) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(37)
	goto L2
L14:
	;
	v40 = int32(4096)
	v41 = v40 - v32
	v46 = v32 + int32(1)
	if v46 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v51 = v41
	v54 = int32(0)
	v55 = v2
	v57 = v2
	v61 = v2
	goto L19
L16:
	;
	v47 = F__emscripten_memcpy_bulkmem(m, v41+(v20+v40), l0, v46)
	mBase = m.M
	goto L18
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	v65 = v20 + int32(4096) + v51
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v66 == int32(47) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v860 = v20 + int32(4096) + v848
	v862 = v860
	goto L251
L22:
	;
	v848 = v71
	v849 = v843
	v850 = int32(0)
	v852 = int32(0)
	goto L21
L23:
	;
	v69 = int32(1)
	v71 = v51 + v69
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+(v20+int32(4096))))))
	v76 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v76)
	if v75 != v76 {
		v843 = v69
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L34
L26:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v20)+uint32(_consts[1220]))))
	if v84 == int32(47) {
		v843 = v69
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v87 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v87)
	v843 = int32(2)
	goto L22
L28:
	;
	v551 = v550 + v54
	if base.Ui32(int32(4095)) < base.Ui32(v551) {
		goto L13
	} else {
		goto L156
	}
L29:
	;
	v549 = v51
	v550 = v187
	goto L28
L30:
	;
	v187 = v176 - v65
	if v187|v55 != 0 {
		goto L54
	} else {
		goto L55
	}
L31:
	;
	goto L30
L32:
	;
	v166 = v161
	goto L50
L33:
	;
	v161 = v153
	goto L32
L34:
	;
	if v65&int32(3) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v100 = v65
	goto L40
L38:
	;
	v113 = v65
	goto L39
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v122 = int32(-2139062144)
	if (int32(16843008)-v119|v119)&v122 != v122 {
		v153 = v113
		goto L33
	} else {
		goto L45
	}
L40:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v105 == int32(0) {
		v176 = v100
		goto L31
	} else {
		goto L42
	}
L41:
	;
	v113 = v110
	goto L39
L42:
	;
	if int32(47) == v105 {
		v176 = v100
		goto L31
	} else {
		goto L43
	}
L43:
	;
	v110 = v100 + int32(1)
	if v110&int32(3) != 0 {
		v100 = v110
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v128 = v113
	v130 = v119
	goto L46
L46:
	;
	v134 = v130 ^ int32(791621423)
	v137 = int32(-2139062144)
	if (int32(16843008)-v134|v134)&v137 != v137 {
		v153 = v128
		goto L33
	} else {
		goto L48
	}
L47:
	;
	v161 = v143
	goto L32
L48:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v143 = v128 + int32(4)
	v147 = int32(-2139062144)
	if (v141|(int32(16843008)-v141))&v147 == v147 {
		v128 = v143
		v130 = v141
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v168 == int32(0) {
		v176 = v166
		goto L31
	} else {
		goto L52
	}
L51:
	;
	v176 = v166
	goto L31
L52:
	;
	if v168 != int32(47) {
		v166 = v166 + int32(1)
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	if v187 != int32(1) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v54))) = uint8(v215)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v219 != int32(47) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	if v54 == int32(0) {
		goto L29
	} else {
		goto L60
	}
L58:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v191 != int32(46) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v848 = v51 + int32(1)
	v849 = v54
	v850 = v55
	v852 = v57
	goto L21
L60:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v54-int32(1)))))
	if v201 == int32(47) {
		goto L29
	} else {
		goto L61
	}
L61:
	;
	if v51 == int32(0) {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	v206 = int32(1)
	v207 = v51 - v206
	v211 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v207+(v20+int32(4096))))) = uint8(v211)
	v549 = v207
	v550 = v187 + v206
	goto L28
L63:
	;
	v225 = F_getcwd(m, v20+int32(4096), int32(4097))
	mBase = m.M
	if v225 == int32(0) {
		v912 = v215
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v540 = F_strlen(m, v20)
	mBase = m.M
	v542 = v540 + int32(1)
	v543 = F_emscripten_builtin_malloc(m, v542)
	mBase = m.M
	if v543 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L66:
	;
	v228 = int32(0)
	v230 = v20 + int32(4096)
	if v230&int32(3) == v228 {
		v254 = v230
		goto L69
	} else {
		goto L70
	}
L67:
	;
	if v57 != 0 {
		goto L84
	} else {
		goto L85
	}
L68:
	;
	v287 = v279 - v230
	goto L67
L69:
	;
	v258 = v254
	goto L78
L70:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[211]))))
	if v238 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v287 = int32(0)
	goto L67
L72:
	;
	goto L73
L73:
	;
	v243 = v230
	goto L74
L74:
	;
	v247 = v243 + int32(1)
	if v247&int32(3) == int32(0) {
		v254 = v247
		goto L69
	} else {
		goto L76
	}
L75:
	;
	v279 = v247
	goto L68
L76:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	if v252 != 0 {
		v243 = v247
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v267 = int32(-2139062144)
	if (int32(16843008)-v264|v264)&v267 == v267 {
		v258 = v258 + int32(4)
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v273 = v258
	goto L81
L80:
	;
	goto L79
L81:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v277 != 0 {
		v273 = v273 + int32(1)
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v279 = v273
	goto L68
L83:
	;
	goto L82
L84:
	;
	v289 = v287
	v293 = v228
	v295 = v57
	goto L87
L85:
	;
	v340 = v287
	v344 = v228
	goto L86
L86:
	;
	v352 = v54 - v344
	if v54 == v344 {
		v367 = v340
		goto L98
	} else {
		goto L99
	}
L87:
	;
	v304 = v295 - int32(1)
	v306 = v289
	goto L90
L88:
	;
	v340 = v338
	v344 = v335
	goto L86
L89:
	;
	v333 = v293 + int32(2)
	if base.Ui32(v333) < base.Ui32(v54) {
		goto L94
	} else {
		goto L95
	}
L90:
	;
	if base.Ui32(v306) < base.Ui32(int32(2)) {
		v329 = base.B2i32(v289 != int32(0))
		goto L89
	} else {
		goto L92
	}
L91:
	;
	v329 = v306
	goto L89
L92:
	;
	v321 = v306 - int32(1)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321+(v20+int32(4096))))))
	if v325 != int32(47) {
		v306 = v321
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v335 = v293 + int32(3)
	goto L96
L95:
	;
	v335 = v333
	goto L96
L96:
	;
	v338 = v329 - base.B2i32(base.Ui32(int32(1)) < base.Ui32(v306))
	if v304 != 0 {
		v289 = v338
		v293 = v335
		v295 = v304
		goto L87
	} else {
		goto L97
	}
L97:
	;
	goto L88
L98:
	;
	if base.Ui32(v367+v352-int32(4095)) < base.Ui32(int32(-4096)) {
		goto L13
	} else {
		goto L101
	}
L99:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340+v20+int32(4095)))))
	if v357 == int32(47) {
		v367 = v340
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v363 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(4096)+v340))) = uint8(v363)
	v367 = v340 + int32(1)
	goto L98
L101:
	;
	v373 = v367 + v20
	v374 = v20 + v344
	v376 = v352 + int32(1)
	if v373 == v374 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v367 != 0 {
		goto L149
	} else {
		goto L150
	}
L103:
	;
	goto L102
L104:
	;
	v380 = v373 + v376
	if base.Ui32(v374-v380) <= base.Ui32(int32(0)-v376<<(uint(int32(1))%32)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v387 = F___memcpy(m, v373, v374, v376)
	mBase = m.M
	goto L102
L106:
	;
	goto L107
L107:
	;
	v390 = (v373 ^ v374) & int32(3)
	if base.Ui32(v373) < base.Ui32(v374) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	if v492 == int32(0) {
		goto L103
	} else {
		goto L144
	}
L109:
	;
	if base.Ui32(v470) <= base.Ui32(int32(3)) {
		v491 = v469
		v492 = v470
		v493 = v471
		goto L108
	} else {
		goto L140
	}
L110:
	;
	if v390 != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if v390 != 0 {
		v452 = v376
		goto L123
	} else {
		goto L124
	}
L113:
	;
	v491 = v374
	v492 = v376
	v493 = v373
	goto L108
L114:
	;
	goto L115
L115:
	;
	if v373&int32(3) == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v469 = v374
	v470 = v376
	v471 = v373
	goto L109
L117:
	;
	goto L118
L118:
	;
	v397 = v374
	v398 = v376
	v399 = v373
	goto L119
L119:
	;
	if v398 == int32(0) {
		goto L103
	} else {
		goto L121
	}
L120:
	;
	v469 = v406
	v470 = v408
	v471 = v410
	goto L109
L121:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	*(*uint8)(unsafe.Add(mBase, uint32(v399))) = uint8(v403)
	v405 = int32(1)
	v406 = v397 + v405
	v408 = v398 - v405
	v410 = v399 + v405
	if v410&int32(3) != 0 {
		v397 = v406
		v398 = v408
		v399 = v410
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	if v452 == int32(0) {
		goto L103
	} else {
		goto L136
	}
L124:
	;
	if v380&int32(3) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v417 = v376
	goto L128
L126:
	;
	v432 = v376
	goto L127
L127:
	;
	if base.Ui32(v432) <= base.Ui32(int32(3)) {
		v452 = v432
		goto L123
	} else {
		goto L132
	}
L128:
	;
	if v417 == int32(0) {
		goto L103
	} else {
		goto L130
	}
L129:
	;
	v432 = v423
	goto L127
L130:
	;
	v423 = v417 - int32(1)
	v424 = v373 + v423
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+v423))))
	*(*uint8)(unsafe.Add(mBase, uint32(v424))) = uint8(v426)
	if v424&int32(3) != 0 {
		v417 = v423
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v439 = v432
	goto L133
L133:
	;
	v443 = v439 - int32(4)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v374+v443)))
	*(*int32)(unsafe.Add(mBase, uint32(v373+v443))) = v446
	if base.Ui32(int32(3)) < base.Ui32(v443) {
		v439 = v443
		goto L133
	} else {
		goto L135
	}
L134:
	;
	v452 = v443
	goto L123
L135:
	;
	goto L134
L136:
	;
	v459 = v452
	goto L137
L137:
	;
	v463 = v459 - int32(1)
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+v463))))
	*(*uint8)(unsafe.Add(mBase, uint32(v373+v463))) = uint8(v466)
	if v463 != 0 {
		v459 = v463
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L103
L139:
	;
	goto L138
L140:
	;
	v476 = v469
	v477 = v470
	v478 = v471
	goto L141
L141:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v476)))
	*(*int32)(unsafe.Add(mBase, uint32(v478))) = v480
	v482 = int32(4)
	v483 = v476 + v482
	v485 = v478 + v482
	v487 = v477 - v482
	if base.Ui32(int32(3)) < base.Ui32(v487) {
		v476 = v483
		v477 = v487
		v478 = v485
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v491 = v483
	v492 = v487
	v493 = v485
	goto L108
L143:
	;
	goto L142
L144:
	;
	v498 = v491
	v499 = v492
	v500 = v493
	goto L145
L145:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	*(*uint8)(unsafe.Add(mBase, uint32(v500))) = uint8(v502)
	v504 = int32(1)
	v509 = v499 - v504
	if v509 != 0 {
		v498 = v498 + v504
		v499 = v509
		v500 = v500 + v504
		goto L145
	} else {
		goto L147
	}
L146:
	;
	goto L103
L147:
	;
	goto L146
L148:
	;
	goto L65
L149:
	;
	v523 = F__emscripten_memcpy_bulkmem(m, v20, v20+int32(4096), v367)
	mBase = m.M
	goto L151
L150:
	;
	goto L151
L151:
	;
	goto L148
L152:
	;
	v912 = v548
	goto L1
L153:
	;
	v548 = int32(0)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v547 = F___memcpy(m, v543, v20, v542)
	mBase = m.M
	v548 = v547
	goto L152
L156:
	;
	if v550 != 0 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v561 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v551))) = uint8(v561)
	v563 = v549 + v550
	v564 = int32(1)
	if v187 != int32(2) {
		v588 = v564
		goto L165
	} else {
		goto L166
	}
L158:
	;
	v558 = F__emscripten_memcpy_bulkmem(m, v20+v54, v20+int32(4096)+v549, v550)
	mBase = m.M
	goto L160
L159:
	;
	goto L160
L160:
	;
	goto L157
L161:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v836 == int32(47) {
		goto L245
	} else {
		goto L246
	}
L162:
	;
	v639 = v61 + int32(1)
	if v639 == int32(40) {
		goto L190
	} else {
		goto L191
	}
L163:
	;
	if v187 != 0 {
		goto L187
	} else {
		goto L188
	}
L164:
	;
	v611 = v54
	goto L180
L165:
	;
	v592 = F_readlink(m, v20, v20+int32(4096), v563)
	mBase = m.M
	if v592 == v563 {
		goto L13
	} else {
		goto L173
	}
L166:
	;
	v569 = v563 + (v20 + int32(4096))
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569-int32(2)))))
	if v572 != int32(46) {
		v588 = v564
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569-int32(1)))))
	if v577 != int32(46) {
		v588 = v564
		goto L165
	} else {
		goto L168
	}
L168:
	;
	if base.Ui32(v54) <= base.Ui32(v57*int32(3)) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v848 = v563
	v849 = v551
	v850 = v55
	v852 = v57 + int32(1)
	goto L21
L170:
	;
	goto L171
L171:
	;
	v585 = int32(0)
	if v55 == v585 {
		goto L164
	} else {
		goto L172
	}
L172:
	;
	v588 = v585
	goto L165
L173:
	;
	if v592 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(44)
	goto L2
L175:
	;
	goto L176
L176:
	;
	if int32(0) <= v592 {
		goto L162
	} else {
		goto L177
	}
L177:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v602 != int32(28) {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	if v588 != 0 {
		goto L163
	} else {
		goto L179
	}
L179:
	;
	goto L164
L180:
	;
	v620 = int32(0)
	if v611 == v620 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	if v611 != int32(1) {
		goto L161
	} else {
		goto L186
	}
L182:
	;
	v848 = v563
	v849 = int32(0)
	v850 = v620
	v852 = v57
	goto L21
L183:
	;
	goto L184
L184:
	;
	v625 = v611 - int32(1)
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v625))))
	if v627 != int32(47) {
		v611 = v625
		goto L180
	} else {
		goto L185
	}
L185:
	;
	goto L181
L186:
	;
	v848 = v563
	v849 = int32(1)
	v850 = v620
	v852 = v57
	goto L21
L187:
	;
	v633 = v551
	goto L189
L188:
	;
	v633 = v54
	goto L189
L189:
	;
	v637 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20+int32(4096)+v563))))
	v848 = v563
	v849 = v633
	v850 = v637
	v852 = v57
	goto L21
L190:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(32)
	goto L2
L191:
	;
	goto L192
L192:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v592+int32(4095)))))
	if v648 == int32(47) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v654 = v563
	goto L196
L194:
	;
	v673 = v563
	goto L195
L195:
	;
	v685 = v673 - v592
	v687 = v20 + int32(4096)
	v688 = v685 + v687
	if v688 == v687 {
		goto L200
	} else {
		goto L201
	}
L196:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654+(v20+int32(4096))))))
	if v669 == int32(47) {
		v654 = v654 + int32(1)
		goto L196
	} else {
		goto L198
	}
L197:
	;
	v673 = v654
	goto L195
L198:
	;
	goto L197
L199:
	;
	v51 = v685
	v61 = v639
	goto L19
L200:
	;
	goto L199
L201:
	;
	v694 = v688 + v592
	if base.Ui32(v687-v694) <= base.Ui32(int32(0)-v592<<(uint(int32(1))%32)) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v701 = F___memcpy(m, v688, v687, v592)
	mBase = m.M
	goto L199
L203:
	;
	goto L204
L204:
	;
	v704 = (v688 ^ v687) & int32(3)
	if base.Ui32(v688) < base.Ui32(v687) {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	if v806 == int32(0) {
		goto L200
	} else {
		goto L241
	}
L206:
	;
	if base.Ui32(v784) <= base.Ui32(int32(3)) {
		v805 = v783
		v806 = v784
		v807 = v785
		goto L205
	} else {
		goto L237
	}
L207:
	;
	if v704 != 0 {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	goto L209
L209:
	;
	if v704 != 0 {
		v766 = v592
		goto L220
	} else {
		goto L221
	}
L210:
	;
	v805 = v687
	v806 = v592
	v807 = v688
	goto L205
L211:
	;
	goto L212
L212:
	;
	if v688&int32(3) == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v783 = v687
	v784 = v592
	v785 = v688
	goto L206
L214:
	;
	goto L215
L215:
	;
	v711 = v687
	v712 = v592
	v713 = v688
	goto L216
L216:
	;
	if v712 == int32(0) {
		goto L200
	} else {
		goto L218
	}
L217:
	;
	v783 = v720
	v784 = v722
	v785 = v724
	goto L206
L218:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711))))
	*(*uint8)(unsafe.Add(mBase, uint32(v713))) = uint8(v717)
	v719 = int32(1)
	v720 = v711 + v719
	v722 = v712 - v719
	v724 = v713 + v719
	if v724&int32(3) != 0 {
		v711 = v720
		v712 = v722
		v713 = v724
		goto L216
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	if v766 == int32(0) {
		goto L200
	} else {
		goto L233
	}
L221:
	;
	if v694&int32(3) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v731 = v592
	goto L225
L223:
	;
	v746 = v592
	goto L224
L224:
	;
	if base.Ui32(v746) <= base.Ui32(int32(3)) {
		v766 = v746
		goto L220
	} else {
		goto L229
	}
L225:
	;
	if v731 == int32(0) {
		goto L200
	} else {
		goto L227
	}
L226:
	;
	v746 = v737
	goto L224
L227:
	;
	v737 = v731 - int32(1)
	v738 = v688 + v737
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687+v737))))
	*(*uint8)(unsafe.Add(mBase, uint32(v738))) = uint8(v740)
	if v738&int32(3) != 0 {
		v731 = v737
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v753 = v746
	goto L230
L230:
	;
	v757 = v753 - int32(4)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v687+v757)))
	*(*int32)(unsafe.Add(mBase, uint32(v688+v757))) = v760
	if base.Ui32(int32(3)) < base.Ui32(v757) {
		v753 = v757
		goto L230
	} else {
		goto L232
	}
L231:
	;
	v766 = v757
	goto L220
L232:
	;
	goto L231
L233:
	;
	v773 = v766
	goto L234
L234:
	;
	v777 = v773 - int32(1)
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687+v777))))
	*(*uint8)(unsafe.Add(mBase, uint32(v688+v777))) = uint8(v780)
	if v777 != 0 {
		v773 = v777
		goto L234
	} else {
		goto L236
	}
L235:
	;
	goto L200
L236:
	;
	goto L235
L237:
	;
	v790 = v783
	v791 = v784
	v792 = v785
	goto L238
L238:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	*(*int32)(unsafe.Add(mBase, uint32(v792))) = v794
	v796 = int32(4)
	v797 = v790 + v796
	v799 = v792 + v796
	v801 = v791 - v796
	if base.Ui32(int32(3)) < base.Ui32(v801) {
		v790 = v797
		v791 = v801
		v792 = v799
		goto L238
	} else {
		goto L240
	}
L239:
	;
	v805 = v797
	v806 = v801
	v807 = v799
	goto L205
L240:
	;
	goto L239
L241:
	;
	v812 = v805
	v813 = v806
	v814 = v807
	goto L242
L242:
	;
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812))))
	*(*uint8)(unsafe.Add(mBase, uint32(v814))) = uint8(v816)
	v818 = int32(1)
	v823 = v813 - v818
	if v823 != 0 {
		v812 = v812 + v818
		v813 = v823
		v814 = v814 + v818
		goto L242
	} else {
		goto L244
	}
L243:
	;
	goto L200
L244:
	;
	goto L243
L245:
	;
	v839 = int32(2)
	goto L247
L246:
	;
	v839 = v625
	goto L247
L247:
	;
	if v611 != int32(2) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v842 = v625
	goto L250
L249:
	;
	v842 = v839
	goto L250
L250:
	;
	v848 = v563
	v849 = v842
	v850 = v620
	v852 = v57
	goto L21
L251:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862))))
	if v876 == int32(47) {
		v862 = v862 + int32(1)
		goto L251
	} else {
		goto L253
	}
L252:
	;
	v51 = v848 + (v862 - v860)
	v54 = v849
	v55 = v850
	v57 = v852
	goto L19
L253:
	;
	goto L252
L254:
	;
	m.G0 = v16 + int32(16)
	return v1065
L255:
	;
	v929 = int32(-1)
	v932 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	goto L257
L257:
	;
	goto L267
L258:
	;
	return int32(0)
L259:
	;
	if v932 == int32(0) {
		v1065 = v929
		goto L254
	} else {
		goto L260
	}
L260:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L258
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg_internal(m, int32(289553), v16)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L258
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(492341), int32(253), int32(316852))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L258
	} else {
		goto L263
	}
L263:
	;
	v1065 = v929
	goto L254
L264:
	;
	F_emscripten_builtin_free(m, v912)
	mBase = m.M
	v1065 = v2
	goto L254
L265:
	;
	v1061 = F_strlen(m, v1050)
	mBase = m.M
	goto L264
L267:
	;
	goto L268
L268:
	;
	v955 = int32(1023)
	if (l0^v912)&int32(3) != 0 {
		goto L272
	} else {
		goto L273
	}
L269:
	;
	v1054 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1051))) = uint8(v1054)
	goto L265
L270:
	;
	v1035 = v1030
	v1036 = v1031
	v1037 = v1032
	goto L292
L271:
	;
	if v1025 == int32(0) {
		v1050 = v1023
		v1051 = v1024
		goto L269
	} else {
		goto L291
	}
L272:
	;
	v1023 = v912
	v1024 = l0
	v1025 = v955
	goto L271
L273:
	;
	goto L274
L274:
	;
	if v912&int32(3) == int32(0) {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	if v992 == int32(0) {
		v1050 = v989
		v1051 = v990
		goto L269
	} else {
		goto L284
	}
L276:
	;
	v989 = v912
	v990 = l0
	v991 = v955
	v992 = int32(1)
	goto L275
L277:
	;
	goto L278
L278:
	;
	v968 = v912
	v969 = l0
	v970 = v955
	goto L279
L279:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v968))))
	*(*uint8)(unsafe.Add(mBase, uint32(v969))) = uint8(v972)
	if v972 == int32(0) {
		v1030 = v968
		v1031 = v969
		v1032 = v970
		goto L270
	} else {
		goto L281
	}
L280:
	;
	v989 = v983
	v990 = v977
	v991 = v979
	v992 = v981
	goto L275
L281:
	;
	v976 = int32(1)
	v977 = v969 + v976
	v979 = v970 - v976
	v980 = int32(0)
	v981 = base.B2i32(v979 != v980)
	v983 = v968 + v976
	if v983&int32(3) == v980 {
		v989 = v983
		v990 = v977
		v991 = v979
		v992 = v981
		goto L275
	} else {
		goto L282
	}
L282:
	;
	if v979 != 0 {
		v968 = v983
		v969 = v977
		v970 = v979
		goto L279
	} else {
		goto L283
	}
L283:
	;
	goto L280
L284:
	;
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989))))
	if v995 == int32(0) {
		v1023 = v989
		v1024 = v990
		v1025 = v991
		goto L271
	} else {
		goto L285
	}
L285:
	;
	if base.Ui32(v991) < base.Ui32(int32(4)) {
		v1023 = v989
		v1024 = v990
		v1025 = v991
		goto L271
	} else {
		goto L286
	}
L286:
	;
	v1001 = v989
	v1002 = v990
	v1003 = v991
	goto L287
L287:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1001)))
	v1009 = int32(-2139062144)
	if (int32(16843008)-v1006|v1006)&v1009 != v1009 {
		v1030 = v1001
		v1031 = v1002
		v1032 = v1003
		goto L270
	} else {
		goto L289
	}
L288:
	;
	v1023 = v1017
	v1024 = v1015
	v1025 = v1019
	goto L271
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1002))) = v1006
	v1014 = int32(4)
	v1015 = v1002 + v1014
	v1017 = v1001 + v1014
	v1019 = v1003 - v1014
	if base.Ui32(int32(3)) < base.Ui32(v1019) {
		v1001 = v1017
		v1002 = v1015
		v1003 = v1019
		goto L287
	} else {
		goto L290
	}
L290:
	;
	goto L288
L291:
	;
	v1030 = v1023
	v1031 = v1024
	v1032 = v1025
	goto L270
L292:
	;
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1036))) = uint8(v1039)
	if v1039 == int32(0) {
		v1050 = v1035
		v1051 = v1036
		goto L269
	} else {
		goto L294
	}
L293:
	;
	v1050 = v1046
	v1051 = v1044
	goto L269
L294:
	;
	v1043 = int32(1)
	v1044 = v1036 + v1043
	v1046 = v1035 + v1043
	v1048 = v1037 - v1043
	if v1048 != 0 {
		v1035 = v1046
		v1036 = v1044
		v1037 = v1048
		goto L292
	} else {
		goto L295
	}
L295:
	;
	goto L293
}
func F_nulltestsel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 float32
	_ = v22
	var v23 float64
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 float64
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v82 float64
	_ = v82
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	F_examine_variable(m, l0, l2, l3, v8+int32(-32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return float64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
		if v18 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
			v22 = *(*float32)(unsafe.Add(mBase, uint32(v19+v20)+8))
			v23 = base.F64_promote_f32(v22)
			switch l1 {
			case 0:
				v68 = v23
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
				m.T0[v69].(func(*base.Module, int32))(m, v18)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return float64(0)
				} else {
					if base.F64_lt(v68, float64(0)) != 0 {
						v82 = float64(0)
					} else {
						if base.F64_gt(v68, float64(1)) == int32(0) {
							v82 = v68
						} else {
							v82 = float64(1)
						}
					}
					m.G0 = v10 - int32(-64)
					return v82
				}
			case 1:
				v68 = base.F64_sub(float64(1), v23)
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
				m.T0[v69].(func(*base.Module, int32))(m, v18)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return float64(0)
				} else {
					if base.F64_lt(v68, float64(0)) != 0 {
						v82 = float64(0)
					} else {
						if base.F64_gt(v68, float64(1)) == int32(0) {
							v82 = v68
						} else {
							v82 = float64(1)
						}
					}
					m.G0 = v10 - int32(-64)
					return v82
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return float64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
					F_errmsg_internal(m, int32(476556), v8+int32(-48))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return float64(0)
					} else {
						F_errfinish(m, int32(487199), int32(1741), int32(302251))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
			if v39 == int32(0) {
				switch l1 {
				case 0:
					v82 = float64(0.005)
					m.G0 = v10 - int32(-64)
					return v82
				case 1:
					v82 = float64(0.995)
					m.G0 = v10 - int32(-64)
					return v82
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return float64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
						F_errmsg_internal(m, int32(476556), v10)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return float64(0)
						} else {
							F_errfinish(m, int32(487199), int32(1769), int32(302251))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
				if v42 != int32(6) {
					switch l1 {
					case 0:
						v82 = float64(0.005)
						m.G0 = v10 - int32(-64)
						return v82
					case 1:
						v82 = float64(0.995)
						m.G0 = v10 - int32(-64)
						return v82
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return float64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
							F_errmsg_internal(m, int32(476556), v10)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(487199), int32(1769), int32(302251))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+8)))
					if int32(0) <= v45 {
						switch l1 {
						case 0:
							v82 = float64(0.005)
							m.G0 = v10 - int32(-64)
							return v82
						case 1:
							v82 = float64(0.995)
							m.G0 = v10 - int32(-64)
							return v82
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return float64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
								F_errmsg_internal(m, int32(476556), v10)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(487199), int32(1769), int32(302251))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if l1 != 0 {
							v50 = float64(1)
						} else {
							v50 = float64(0)
						}
						v82 = v50
						m.G0 = v10 - int32(-64)
						return v82
					}
				}
			}
		}
	}
}
func F_numericvar_serialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_enlargeStringInfo(m, l0, int32(4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = int32(24)
	v15 = int32(65280)
	v17 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v10+v11))) = v6<<(uint(v13)%32) | v6&v15<<(uint(v17)%32) | (int32(base.Ui32(v6)>>(uint(v17)%32))&v15 | int32(base.Ui32(v6)>>(uint(v13)%32)))
	v29 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_enlargeStringInfo(m, l0, v29)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(24)
	v41 = int32(65280)
	v43 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v36+v37))) = v32<<(uint(v39)%32) | v32&v41<<(uint(v43)%32) | (int32(base.Ui32(v32)>>(uint(v43)%32))&v41 | int32(base.Ui32(v32)>>(uint(v39)%32)))
	v55 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36 + v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_enlargeStringInfo(m, l0, v55)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = int32(24)
	v67 = int32(65280)
	v69 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v62+v63))) = v58<<(uint(v65)%32) | v58&v67<<(uint(v69)%32) | (int32(base.Ui32(v58)>>(uint(v69)%32))&v67 | int32(base.Ui32(v58)>>(uint(v65)%32)))
	v81 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62 + v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_enlargeStringInfo(m, l0, v81)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = int32(24)
	v93 = int32(65280)
	v95 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v88+v89))) = v84<<(uint(v91)%32) | v84&v93<<(uint(v95)%32) | (int32(base.Ui32(v84)>>(uint(v95)%32))&v93 | int32(base.Ui32(v84)>>(uint(v91)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88 + int32(4)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v110 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v116 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	return
L9:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119+v116<<(uint(int32(1))%32)))))
	F_enlargeStringInfo(m, l0, int32(2))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v130 = int32(8)
	v134 = v123<<(uint(v130)%32) | int32(base.Ui32(v123)>>(uint(v130)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v127+v128))) = uint16(v134)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v127 + int32(2)
	v140 = v116 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v140 < v141 {
		v116 = v140
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_numericvar_to_int128(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
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
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v207 int32
	_ = v207
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v230 int64
	_ = v230
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v244 int64
	_ = v244
	var v251 int64
	_ = v251
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v268 int64
	_ = v268
	var v272 int64
	_ = v272
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v279 int64
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int64
	_ = v287
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v292 int64
	_ = v292
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v352 int64
	_ = v352
	var v353 int64
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v371 int64
	_ = v371
	var v388 int64
	_ = v388
	var v393 int32
	_ = v393
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = F_palloc(m, v20<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v25))) = uint16(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v27 < v29 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = v29 << (uint(int32(1)) % 32)
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = v39 << (uint(int32(2)) % 32)
	if v41+int32(4) < int32(0) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v37 = F__emscripten_memcpy_bulkmem(m, v25+int32(2), v34, v36)
	mBase = m.M
	goto L9
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	m.G0 = v18 + int32(32)
	return
L11:
	;
	v388 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v388
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v388
	F_pfree(m, v25)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L69
	}
L12:
	;
	v47 = v25 + int32(2)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = base.I32_div_s(v41+int32(7), int32(4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v53 <= v52 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if int32(0) < v112 {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v112 = v53
	v113 = v47
	v117 = v39
	goto L13
L15:
	;
	goto L16
L16:
	;
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47+v52<<(uint(int32(1))%32)))))
	if int32(5000) <= v58 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v52
	goto L20
L18:
	;
	v89 = v52
	goto L19
L19:
	;
	if int32(0) <= v89 {
		v112 = v52
		v113 = v47
		v117 = v39
		goto L13
	} else {
		goto L26
	}
L20:
	;
	v76 = int32(1)
	v78 = v25 + v61<<(uint(v76)%32)
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78))))
	v83 = base.B2i32(int32(9998) < v81)
	if int32(9998) < v81 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v89 = v88
	goto L19
L22:
	;
	v84 = int32(-9999)
	goto L24
L23:
	;
	v84 = v76
	goto L24
L24:
	;
	v85 = v84 + v81
	*(*uint16)(unsafe.Add(mBase, uint32(v78))) = uint16(v85)
	v88 = v61 - int32(1)
	if int32(9998) < v81 {
		v61 = v88
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v106 = int32(1)
	v112 = v52 + v106
	v113 = v25
	v117 = v39 + v106
	goto L13
L27:
	;
	v197 = base.I64_extend16_s(base.I64_extend_i32_u(v181))
	v199 = v197 >> (uint(int64(63)) % 64)
	if v188 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L28:
	;
	v129 = v112
	v130 = v113
	v134 = v117
	goto L31
L29:
	;
	goto L30
L30:
	;
	if v112 == int32(0) {
		goto L11
	} else {
		goto L41
	}
L31:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130))))
	if v142 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v157 = v129
	goto L37
L33:
	;
	v145 = int32(1)
	if v145 < v129 {
		v129 = v129 - v145
		v130 = v130 + int32(2)
		v134 = v134 - v145
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L11
L37:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130-int32(2)+v157<<(uint(int32(1))%32)))))
	if v173 != 0 {
		v181 = v142
		v183 = v157
		v184 = v130
		v188 = v134
		goto L27
	} else {
		goto L39
	}
L38:
	;
	goto L11
L39:
	;
	v174 = int32(1)
	if v174 < v157 {
		v157 = v157 - v174
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113))))
	v181 = v180
	v183 = v112
	v184 = v113
	v188 = v117
	goto L27
L42:
	;
	F_pfree(m, v25)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L62
	}
L43:
	;
	v352 = v197
	v353 = v199
	goto L42
L44:
	;
	goto L45
L45:
	;
	v207 = int32(1)
	v219 = v199
	v220 = v197
	goto L46
L46:
	;
	v223 = v18 + int32(16)
	v224 = int64(10000)
	v225 = int64(0)
	v230 = int64(32)
	v233 = int64(base.Ui64(v220) >> (uint(v230) % 64))
	v236 = int64(4294967295)
	v239 = v220 & v236
	v240 = v224 * v239
	v244 = int64(base.Ui64(v240)>>(uint(v230)%64)) + v224*v233
	v251 = v239*v225 + v244&v236
	*(*int64)(unsafe.Add(mBase, uint32(v223)+8)) = v220*v225 + v219*v224 + v225*v233 + int64(base.Ui64(v244)>>(uint(v230)%64)) + int64(base.Ui64(v251)>>(uint(v230)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v223))) = v240&v236 | v251<<(uint(v230)%64)
	goto L48
L47:
	;
	v352 = v276
	v353 = v277
	goto L42
L48:
	;
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v18+int32(24))))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	if v183 <= v207 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v279 = int64(0)
	v283 = m.G0
	v284 = int32(16)
	v285 = v283 - v284
	m.G0 = v285
	v287 = int64(63)
	v288 = v277 >> (uint(v287) % 64)
	v289 = v288 ^ v276
	v292 = v289 + int64(base.Ui64(v277)>>(uint(v287)%64))
	F___udivmodti4(m, v285, v292, base.I64_extend_i32_u(base.B2i32(base.Ui64(v292) < base.Ui64(v289)))+(v277^v288), int64(10000), base.I64_extend_i32_u(int32(0))+v279)
	mBase = m.M
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v285)+8))
	v309 = v279 ^ v288
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v285)))
	v311 = v309 ^ v310
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v311 - v309
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v308 ^ v309 - v309 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v311) < base.Ui64(v309)))
	m.G0 = v285 + v284
	goto L53
L50:
	;
	v276 = v263
	v277 = v262
	goto L49
L51:
	;
	goto L52
L52:
	;
	v268 = int64(*(*int16)(unsafe.Add(mBase, uint32(v184+v207<<(uint(int32(1))%32)))))
	v272 = v268 + v263
	v276 = v272
	v277 = v262 + v268>>(uint(int64(63))%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v272) < base.Ui64(v263)))
	goto L49
L53:
	;
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v323^v220|(v325^v219) == int64(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v340 = v207 + int32(1)
	if v340 <= v188 {
		v207 = v340
		v219 = v277
		v220 = v276
		goto L46
	} else {
		goto L61
	}
L55:
	;
	if v48 != int32(16384) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_pfree(m, v25)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L60
	}
L57:
	;
	if v276|(v277^int64(-9223372036854775807-1)) != int64(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	if int64(0) <= v219 {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	goto L10
L61:
	;
	goto L47
L62:
	;
	v362 = base.B2i32(v48 == int32(16384))
	if v48 == int32(16384) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v363 = int64(0) - v352
	goto L65
L64:
	;
	v363 = v352
	goto L65
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v363
	v365 = int64(0)
	if v48 == int32(16384) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v371 = v365 - (v353 + base.I64_extend_i32_u(base.B2i32(v352 != v365)))
	goto L68
L67:
	;
	v371 = v353
	goto L68
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v371
	goto L10
L69:
	;
	goto L10
}
func F_numst(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	v6 = l1 + int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v7 != 0 {
		v9 = v7
		v10 = v6
		for {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v10
			v14 = v10 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			if v15 != 0 {
				v17 = v15
				v18 = v14
				for {
					v19 = F_numst(m, v17, v18)
					mBase = m.M
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
					if v20 != 0 {
						v17 = v20
						v18 = v19
						continue
					} else {
						break
					}
					break
				}
				v23 = v19
			} else {
				v23 = v14
			}
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
			if v24 != 0 {
				v9 = v24
				v10 = v23
				continue
			} else {
				break
			}
			break
		}
		v27 = v23
	} else {
		v27 = v6
	}
	return v27
}
