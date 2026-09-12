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
	v48 = *(*int32)(unsafe.Add(mBase, _consts[969]))
	v50 = *(*int32)(unsafe.Add(mBase, _consts[970]))
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
	F_parse_format(m, v23, v13, int32(1648816), v27, int32(1648432), int32(2), l1)
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
	*(*int32)(unsafe.Add(mBase, _consts[970])) = v162
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[971])))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+976))
	v80 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+976)) = v79 >> (uint(v80) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[972])))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+976)) = v86 >> (uint(v80) % 32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[973])))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+976)) = v93 >> (uint(v80) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[974])))
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
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v123<<(uint(int32(2))%32))+uint32(_consts[971])))
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
	F_parse_format(m, v703, v13, int32(1648816), v717, int32(1648432), int32(2), v703+int32(980))
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
	*(*int32)(unsafe.Add(mBase, _consts[970])) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v195)+976)) = v406
	v735 = v195
	goto L24
L30:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v181<<(uint(int32(2))%32))+uint32(_consts[971])))
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
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255)+uint32(_consts[971])))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+976))
	v260 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v258)+976)) = v259 >> (uint(v260) % 32)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255)+uint32(_consts[972])))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+976)) = v266 >> (uint(v260) % 32)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v255)+uint32(_consts[973])))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v272)+976)) = v273 >> (uint(v260) % 32)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v255)+uint32(_consts[974])))
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
	*(*int32)(unsafe.Add(mBase, _consts[970])) = int32(1073741823)
	goto L47
L57:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v301<<(uint(int32(2))%32))+uint32(_consts[971])))
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
	v358 = *(*int32)(unsafe.Add(mBase, _consts[971]))
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
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v364<<(uint(int32(2))%32))+uint32(_consts[971])))
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
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v387<<(uint(int32(2))%32))+uint32(_consts[971])))
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
	v427 = *(*int32)(unsafe.Add(mBase, _consts[969]))
	*(*int32)(unsafe.Add(mBase, uint32(v427<<(uint(int32(2))%32))+uint32(_consts[971]))) = v424
	v433 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+975)) = uint8(v433)
	v436 = v424 + int32(900)
	goto L81
L78:
	;
	v552 = int32(4473792)
	v554 = *(*int32)(unsafe.Add(mBase, _consts[970]))
	v555 = int32(1)
	v556 = v554 + v555
	*(*int32)(unsafe.Add(mBase, _consts[970])) = v556
	*(*int32)(unsafe.Add(mBase, uint32(v424)+976)) = v556
	v559 = int32(4473700)
	v561 = *(*int32)(unsafe.Add(mBase, _consts[969]))
	*(*int32)(unsafe.Add(mBase, _consts[969])) = v561 + v555
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
	v696 = int32(4473792)
	v698 = *(*int32)(unsafe.Add(mBase, _consts[970]))
	v700 = v698 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[970])) = v700
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
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v306 int32
	_ = v306
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v716 int32
	_ = v716
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v932 int32
	_ = v932
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1004 int32
	_ = v1004
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1257 int32
	_ = v1257
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1366 int32
	_ = v1366
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1594 int32
	_ = v1594
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1631 int32
	_ = v1631
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1646 int32
	_ = v1646
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1779 int64
	_ = v1779
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1878 int32
	_ = v1878
	var v1913 int32
	_ = v1913
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1946 int32
	_ = v1946
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v2007 int32
	_ = v2007
	var v2045 int32
	_ = v2045
	var v2053 int32
	_ = v2053
	var v2063 int32
	_ = v2063
	var v2101 int32
	_ = v2101
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2151 int32
	_ = v2151
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2171 int32
	_ = v2171
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2193 int32
	_ = v2193
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2217 int32
	_ = v2217
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2240 int32
	_ = v2240
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2257 int32
	_ = v2257
	var v2263 int32
	_ = v2263
	var v2269 int32
	_ = v2269
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2295 int32
	_ = v2295
	var v2301 int32
	_ = v2301
	var v2316 int32
	_ = v2316
	var v2320 int32
	_ = v2320
	var v2324 int32
	_ = v2324
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2337 int32
	_ = v2337
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2371 int32
	_ = v2371
	var v2376 int32
	_ = v2376
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2387 int32
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
	var v2405 int32
	_ = v2405
	var v2407 int32
	_ = v2407
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2425 int32
	_ = v2425
	var v2429 int32
	_ = v2429
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2443 int32
	_ = v2443
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2448 int32
	_ = v2448
	var v2459 int64
	_ = v2459
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2472 int64
	_ = v2472
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2484 int32
	_ = v2484
	var v2489 int32
	_ = v2489
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2500 int32
	_ = v2500
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2542 int32
	_ = v2542
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2561 int32
	_ = v2561
	var v2572 int64
	_ = v2572
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2585 int64
	_ = v2585
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2599 int32
	_ = v2599
	var v2601 int32
	_ = v2601
	var v2604 int32
	_ = v2604
	var v2606 int32
	_ = v2606
	var v2609 int32
	_ = v2609
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2626 int32
	_ = v2626
	var v2628 int32
	_ = v2628
	var v2631 int32
	_ = v2631
	var v2633 int32
	_ = v2633
	var v2636 int32
	_ = v2636
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2652 int32
	_ = v2652
	var v2655 int32
	_ = v2655
	var v2657 int32
	_ = v2657
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2681 int32
	_ = v2681
	var v2686 int32
	_ = v2686
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2704 int32
	_ = v2704
	var v2729 int32
	_ = v2729
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2752 int32
	_ = v2752
	var v2768 int32
	_ = v2768
	var v2769 int32
	_ = v2769
	var v2780 int32
	_ = v2780
	var v2785 int32
	_ = v2785
	var v2795 int32
	_ = v2795
	var v2820 int32
	_ = v2820
	var v2835 int32
	_ = v2835
	var v2845 int32
	_ = v2845
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2873 int32
	_ = v2873
	var v2875 int32
	_ = v2875
	var v2976 int32
	_ = v2976
	var v2979 int32
	_ = v2979
	var v2983 int32
	_ = v2983
	var v2988 int32
	_ = v2988
	var v2992 int32
	_ = v2992
	var v2995 int32
	_ = v2995
	var v2999 int32
	_ = v2999
	var v3004 int32
	_ = v3004
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
	v2992 = m.ExcPending
	if v2992 != 0 {
		goto L71
	} else {
		goto L816
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2976 = m.ExcPending
	if v2976 != 0 {
		goto L71
	} else {
		goto L812
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
	v397 = int32(730275)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v398 == int32(0) {
		goto L67
	} else {
		goto L68
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
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v343 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v343)
	v345 = int32(0)
	v355 = v345
	v356 = v345
	v365 = v9
	v366 = v9
	v396 = v341 + v342 - int32(1)
	goto L32
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v169 = int32(34)
	if v164&v169 != v169 {
		v306 = v9
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
	v355 = l5
	v356 = l6
	v365 = v306
	v366 = v165
	v396 = v166 + v167 - base.B2i32(l5|v165 != int32(0))
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
		v306 = v9
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
		v306 = v191
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
	v282 = F_strlen(m, l3)
	mBase = m.M
	v284 = v282 - int32(1)
	v285 = v280 - l5
	if v284 < v285 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v287 = v284
	goto L62
L61:
	;
	v287 = v285
	goto L62
L62:
	;
	v288 = l3 + v287
	if base.Ui32(v191) < base.Ui32(v288) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v290 = v288
	goto L65
L64:
	;
	v290 = v191
	goto L65
L65:
	;
	v306 = v290
	goto L46
L66:
	;
	v448 = int32(1)
	v450 = l3 + (l7 ^ v448)
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v451 == v448 {
		goto L103
	} else {
		goto L104
	}
L67:
	;
	v443 = int32(654124)
	v444 = int32(654126)
	v445 = int32(654114)
	v446 = int32(654088)
	v447 = v397
	goto L66
L68:
	;
	goto L69
L69:
	;
	v405 = F_PGLC_localeconv(m)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v405)+32))
	if v411 != 0 {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	return
L72:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v405)+36))
	if v407 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407))))
	if v408 != 0 {
		v410 = v407
		goto L70
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v410 = int32(654114)
	goto L70
L76:
	;
	goto L75
L77:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	if v415 != 0 {
		goto L83
	} else {
		goto L84
	}
L78:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v412 != 0 {
		v414 = v411
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v414 = int32(654126)
	goto L77
L81:
	;
	goto L80
L82:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v420&int32(4) != 0 {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	if v416 != 0 {
		v418 = v415
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v418 = int32(654088)
	goto L82
L86:
	;
	goto L85
L87:
	;
	v423 = v418
	goto L89
L88:
	;
	v423 = int32(654088)
	goto L89
L89:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v424 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v405)+16))
	if v435 == int32(0) {
		v443 = v434
		v444 = v414
		v445 = v410
		v446 = v423
		v447 = v397
		goto L66
	} else {
		goto L99
	}
L91:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
	if v425 != 0 {
		v434 = v424
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	if v427 != int32(44) {
		v434 = int32(654124)
		goto L90
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+1)))
	if v432 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v433 = int32(654124)
	goto L98
L97:
	;
	v433 = int32(654088)
	goto L98
L98:
	;
	v434 = v433
	goto L90
L99:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	if v438 == int32(0) {
		v443 = v434
		v444 = v414
		v445 = v410
		v446 = v423
		v447 = v397
		goto L66
	} else {
		goto L100
	}
L100:
	;
	v443 = v434
	v444 = v414
	v445 = v410
	v446 = v423
	v447 = v435
	goto L66
L101:
	;
	v2869 = v2835 - int32(1)
	v2870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2869))))
	if v2870 == int32(46) {
		goto L809
	} else {
		goto L810
	}
L102:
	;
	if l7 == int32(0) {
		v2835 = v2785
		v2845 = v2795
		goto L101
	} else {
		goto L807
	}
L103:
	;
	v2780 = l2
	v2785 = v450
	v2795 = v9
	goto L102
L104:
	;
	goto L105
L105:
	;
	v455 = base.B2i32(v356 == int32(45))
	if v356 == int32(45) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v456 = v445
	goto L108
L107:
	;
	v456 = v444
	goto L108
L108:
	;
	v460 = base.B2i32(v356 == int32(43))
	if v356 == int32(43) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v461 = int32(32)
	goto L111
L110:
	;
	v461 = int32(62)
	goto L111
L111:
	;
	if v356 == int32(43) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v464 = int32(32)
	goto L114
L113:
	;
	v464 = int32(60)
	goto L114
L114:
	;
	v465 = l2 + l4
	v471 = l0
	v479 = l2
	v480 = v451
	v484 = v450
	v485 = v9
	v487 = v366
	v489 = v9
	v494 = v9
	v495 = v9
	v502 = v9
	goto L115
L115:
	;
	if l7 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	v2780 = v2729
	v2785 = v2734
	v2795 = v2744
	goto L102
L117:
	;
	v2768 = v471 + int32(12)
	v2769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2768))))
	if v2769 != int32(1) {
		v471 = v2768
		v479 = v2729
		v480 = v2769
		v484 = v2734
		v485 = v2735
		v487 = v2737
		v489 = v2739
		v494 = v2744
		v495 = v2745
		v502 = v2752
		goto L115
	} else {
		goto L806
	}
L118:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v471)+8))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)+8))
	switch v610 {
	case 0:
		goto L158
	case 1, 2, 3, 6:
		goto L159
	default:
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	case 9:
		goto L157
	case 10:
		goto L156
	case 11:
		goto L152
	case 12:
		goto L151
	case 14, 30:
		goto L155
	case 15:
		goto L150
	case 18:
		goto L153
	case 34:
		goto L154
	}
L119:
	;
	v606 = F_pg_mblen_range(m, v479, v465)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L71
	} else {
		goto L147
	}
L120:
	;
	if base.Ui32(v465) <= base.Ui32(v479) {
		v2835 = v484
		v2845 = v494
		goto L101
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if v480&int32(255) == int32(2) {
		goto L118
	} else {
		goto L125
	}
L123:
	;
	if v480&int32(255) != int32(2) {
		goto L119
	} else {
		goto L124
	}
L124:
	;
	goto L118
L125:
	;
	v529 = v471 + int32(1)
	if (v529^v479)&int32(3) != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	v604 = F_strlen(m, v479)
	mBase = m.M
	v2729 = v604 + v479
	v2734 = v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L127:
	;
	goto L126
L128:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v584))) = uint8(v583)
	if v583&int32(255) == int32(0) {
		goto L127
	} else {
		goto L143
	}
L129:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529))))
	v582 = v529
	v583 = v535
	v584 = v479
	goto L128
L130:
	;
	goto L131
L131:
	;
	if v529&int32(3) != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v539 = v529
	v541 = v479
	goto L135
L133:
	;
	v553 = v529
	v555 = v479
	goto L134
L134:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	v560 = int32(-2139062144)
	if (int32(16843008)-v557|v557)&v560 != v560 {
		v582 = v553
		v583 = v557
		v584 = v555
		goto L128
	} else {
		goto L139
	}
L135:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539))))
	*(*uint8)(unsafe.Add(mBase, uint32(v541))) = uint8(v542)
	if v542 == int32(0) {
		goto L127
	} else {
		goto L137
	}
L136:
	;
	v553 = v549
	v555 = v547
	goto L134
L137:
	;
	v546 = int32(1)
	v547 = v541 + v546
	v549 = v539 + v546
	if v549&int32(3) != 0 {
		v539 = v549
		v541 = v547
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v565 = v553
	v566 = v557
	v567 = v555
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567))) = v566
	v569 = int32(4)
	v570 = v567 + v569
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	v573 = v565 + v569
	v577 = int32(-2139062144)
	if (v571|(int32(16843008)-v571))&v577 == v577 {
		v565 = v573
		v566 = v571
		v567 = v570
		goto L140
	} else {
		goto L142
	}
L141:
	;
	v582 = v573
	v583 = v571
	v584 = v570
	goto L128
L142:
	;
	goto L141
L143:
	;
	v591 = v582
	v593 = v584
	goto L144
L144:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v593)+1)) = uint8(v594)
	v596 = int32(1)
	if v594 != 0 {
		v591 = v591 + v596
		v593 = v593 + v596
		goto L144
	} else {
		goto L146
	}
L145:
	;
	goto L127
L146:
	;
	goto L145
L147:
	;
	v2729 = v606 + v479
	v2734 = v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L148:
	;
	v2729 = v2681 + int32(1)
	v2734 = v2686
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v2696
	v2745 = v2697
	v2752 = v2704
	goto L117
L149:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1480)
	v2681 = v1333
	v2686 = v1335
	v2696 = v1336
	v2697 = v1337
	v2704 = v1338
	goto L148
L150:
	;
	if l7 != 0 {
		goto L794
	} else {
		goto L795
	}
L151:
	;
	if l7 != 0 {
		goto L778
	} else {
		goto L779
	}
L152:
	;
	if l7 != 0 {
		goto L762
	} else {
		goto L763
	}
L153:
	;
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2484&int32(1024) != 0 {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L721
	}
L154:
	;
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2371&int32(1024) != 0 {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L680
	}
L155:
	;
	if l7 != 0 {
		goto L557
	} else {
		goto L558
	}
L156:
	;
	if l7 != 0 {
		goto L522
	} else {
		goto L523
	}
L157:
	;
	v1502 = F_strlen(m, v443)
	mBase = m.M
	if l7 != 0 {
		goto L471
	} else {
		goto L472
	}
L158:
	;
	if l7 != 0 {
		goto L458
	} else {
		goto L459
	}
L159:
	;
	if l7 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v611&int32(1024) != 0 {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if base.Ui32(v465) <= base.Ui32(v479) {
		v2681 = v479
		v2686 = v484
		v2696 = v494
		v2697 = v495
		v2704 = v502
		goto L148
	} else {
		goto L328
	}
L163:
	;
	if v487 != 0 {
		v734 = v479
		v736 = v487
		goto L164
	} else {
		goto L165
	}
L164:
	;
	if int32(1)<<(uint(v610)%32)&int32(78) != 0 {
		goto L214
	} else {
		goto L215
	}
L165:
	;
	v615 = v611 & int32(8)
	if v489 < v355 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	if v611&int32(64) != 0 {
		goto L178
	} else {
		goto L179
	}
L167:
	;
	v617 = int32(0)
	if v615 == v617 {
		v734 = v479
		v736 = v617
		goto L164
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	if v615 != 0 {
		goto L166
	} else {
		goto L172
	}
L170:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v620 == v489 {
		goto L166
	} else {
		goto L171
	}
L171:
	;
	v734 = v479
	v736 = v617
	goto L164
L172:
	;
	if l3 != v484 {
		goto L166
	} else {
		goto L173
	}
L173:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v623 != int32(48) {
		goto L166
	} else {
		goto L174
	}
L174:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v626 == int32(0) {
		goto L166
	} else {
		goto L175
	}
L175:
	;
	v629 = int32(0)
	if v365 == v629 {
		v734 = v479
		v736 = v629
		goto L164
	} else {
		goto L176
	}
L176:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	if v632 != int32(46) {
		v734 = v479
		v736 = v629
		goto L164
	} else {
		goto L177
	}
L177:
	;
	goto L166
L178:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v638 != int32(-1) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	if v611&int32(128) != 0 {
		goto L206
	} else {
		goto L207
	}
L181:
	;
	v734 = v479
	v736 = int32(0)
	goto L164
L182:
	;
	goto L183
L183:
	;
	if (v456^v479)&int32(3) != 0 {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	v716 = F_strlen(m, v479)
	mBase = m.M
	v734 = v716 + v479
	v736 = int32(1)
	goto L164
L185:
	;
	goto L184
L186:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v696))) = uint8(v695)
	if v695&int32(255) == int32(0) {
		goto L185
	} else {
		goto L201
	}
L187:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	v694 = v456
	v695 = v647
	v696 = v479
	goto L186
L188:
	;
	goto L189
L189:
	;
	if v456&int32(3) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v651 = v456
	v653 = v479
	goto L193
L191:
	;
	v665 = v456
	v667 = v479
	goto L192
L192:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v665)))
	v672 = int32(-2139062144)
	if (int32(16843008)-v669|v669)&v672 != v672 {
		v694 = v665
		v695 = v669
		v696 = v667
		goto L186
	} else {
		goto L197
	}
L193:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	*(*uint8)(unsafe.Add(mBase, uint32(v653))) = uint8(v654)
	if v654 == int32(0) {
		goto L185
	} else {
		goto L195
	}
L194:
	;
	v665 = v661
	v667 = v659
	goto L192
L195:
	;
	v658 = int32(1)
	v659 = v653 + v658
	v661 = v651 + v658
	if v661&int32(3) != 0 {
		v651 = v661
		v653 = v659
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	v677 = v665
	v678 = v669
	v679 = v667
	goto L198
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v679))) = v678
	v681 = int32(4)
	v682 = v679 + v681
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v677)+4))
	v685 = v677 + v681
	v689 = int32(-2139062144)
	if (v683|(int32(16843008)-v683))&v689 == v689 {
		v677 = v685
		v678 = v683
		v679 = v682
		goto L198
	} else {
		goto L200
	}
L199:
	;
	v694 = v685
	v695 = v683
	v696 = v682
	goto L186
L200:
	;
	goto L199
L201:
	;
	v703 = v694
	v705 = v696
	goto L202
L202:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v705)+1)) = uint8(v706)
	v708 = int32(1)
	if v706 != 0 {
		v703 = v703 + v708
		v705 = v705 + v708
		goto L202
	} else {
		goto L204
	}
L203:
	;
	goto L185
L204:
	;
	goto L203
L205:
	;
	v731 = int32(1)
	v734 = v479 + v731
	v736 = v731
	goto L164
L206:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v464)
	goto L205
L207:
	;
	goto L208
L208:
	;
	switch v356 - int32(43) {
	case 0:
		goto L210
	default:
		v734 = v479
		v736 = int32(0)
		goto L164
	case 2:
		goto L209
	}
L209:
	;
	v728 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v728)
	goto L205
L210:
	;
	if v611&int32(32) != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v734 = v479
	v736 = int32(1)
	goto L164
L212:
	;
	goto L213
L213:
	;
	v726 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v726)
	goto L205
L214:
	;
	v744 = base.B2i32(base.Ui32(v610) <= base.Ui32(int32(6)))
	goto L216
L215:
	;
	v744 = int32(0)
	goto L216
L216:
	;
	if v744 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v2729 = v734
	v2734 = v484
	v2735 = int32(0)
	v2737 = v736
	v2739 = v489 + int32(1)
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L218:
	;
	goto L219
L219:
	;
	if v489 < v355 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v986 = int32(1)
	v990 = v396 + base.B2i32(v355 != int32(0)) + int32(base.Ui32(v985)>>(uint(v986)%32))&v986
	if v983 == v365 {
		goto L295
	} else {
		goto L296
	}
L221:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v489 < v752 {
		goto L225
	} else {
		goto L226
	}
L222:
	;
	goto L223
L223:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	if v770 == int32(46) {
		goto L230
	} else {
		goto L231
	}
L224:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v734))) = uint8(v764)
	v979 = v734 + int32(1)
	v983 = v484
	v984 = v765
	goto L220
L225:
	;
	v760 = int32(32)
	v761 = int32(0)
	if v751&v760 != 0 {
		v979 = v734
		v983 = v484
		v984 = v761
		goto L220
	} else {
		goto L228
	}
L226:
	;
	if v751&int32(8) == int32(0) {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v764 = int32(48)
	v765 = int32(1)
	goto L224
L228:
	;
	v764 = v760
	v765 = v761
	goto L224
L229:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	v979 = v973
	v983 = v484 + base.B2i32(v975 != int32(0))
	v984 = v974
	goto L220
L230:
	;
	if v365 != 0 {
		goto L234
	} else {
		goto L235
	}
L231:
	;
	goto L232
L232:
	;
	if v365 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L233:
	;
	v853 = int32(0)
	if v769&int32(32) == v853 {
		v973 = v734
		v974 = v853
		goto L229
	} else {
		goto L259
	}
L234:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	if v773 == int32(46) {
		goto L233
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if (v446^v734)&int32(3) != 0 {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	goto L236
L238:
	;
	v850 = F_strlen(m, v734)
	mBase = m.M
	v973 = v850 + v734
	v974 = int32(0)
	goto L229
L239:
	;
	goto L238
L240:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v830))) = uint8(v829)
	if v829&int32(255) == int32(0) {
		goto L239
	} else {
		goto L255
	}
L241:
	;
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	v828 = v446
	v829 = v781
	v830 = v734
	goto L240
L242:
	;
	goto L243
L243:
	;
	if v446&int32(3) != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v785 = v446
	v787 = v734
	goto L247
L245:
	;
	v799 = v446
	v801 = v734
	goto L246
L246:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	v806 = int32(-2139062144)
	if (int32(16843008)-v803|v803)&v806 != v806 {
		v828 = v799
		v829 = v803
		v830 = v801
		goto L240
	} else {
		goto L251
	}
L247:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v785))))
	*(*uint8)(unsafe.Add(mBase, uint32(v787))) = uint8(v788)
	if v788 == int32(0) {
		goto L239
	} else {
		goto L249
	}
L248:
	;
	v799 = v795
	v801 = v793
	goto L246
L249:
	;
	v792 = int32(1)
	v793 = v787 + v792
	v795 = v785 + v792
	if v795&int32(3) != 0 {
		v785 = v795
		v787 = v793
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v811 = v799
	v812 = v803
	v813 = v801
	goto L252
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v812
	v815 = int32(4)
	v816 = v813 + v815
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	v819 = v811 + v815
	v823 = int32(-2139062144)
	if (v817|(int32(16843008)-v817))&v823 == v823 {
		v811 = v819
		v812 = v817
		v813 = v816
		goto L252
	} else {
		goto L254
	}
L253:
	;
	v828 = v819
	v829 = v817
	v830 = v816
	goto L240
L254:
	;
	goto L253
L255:
	;
	v837 = v828
	v839 = v830
	goto L256
L256:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v839)+1)) = uint8(v840)
	v842 = int32(1)
	if v840 != 0 {
		v837 = v837 + v842
		v839 = v839 + v842
		goto L256
	} else {
		goto L258
	}
L257:
	;
	goto L239
L258:
	;
	goto L257
L259:
	;
	if (v446^v734)&int32(3) != 0 {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v932 = F_strlen(m, v734)
	mBase = m.M
	v973 = v932 + v734
	v974 = v853
	goto L229
L261:
	;
	goto L260
L262:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v912))) = uint8(v911)
	if v911&int32(255) == int32(0) {
		goto L261
	} else {
		goto L277
	}
L263:
	;
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	v910 = v446
	v911 = v863
	v912 = v734
	goto L262
L264:
	;
	goto L265
L265:
	;
	if v446&int32(3) != 0 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v867 = v446
	v869 = v734
	goto L269
L267:
	;
	v881 = v446
	v883 = v734
	goto L268
L268:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	v888 = int32(-2139062144)
	if (int32(16843008)-v885|v885)&v888 != v888 {
		v910 = v881
		v911 = v885
		v912 = v883
		goto L262
	} else {
		goto L273
	}
L269:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867))))
	*(*uint8)(unsafe.Add(mBase, uint32(v869))) = uint8(v870)
	if v870 == int32(0) {
		goto L261
	} else {
		goto L271
	}
L270:
	;
	v881 = v877
	v883 = v875
	goto L268
L271:
	;
	v874 = int32(1)
	v875 = v869 + v874
	v877 = v867 + v874
	if v877&int32(3) != 0 {
		v867 = v877
		v869 = v875
		goto L269
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	v893 = v881
	v894 = v885
	v895 = v883
	goto L274
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v895))) = v894
	v897 = int32(4)
	v898 = v895 + v897
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v893)+4))
	v901 = v893 + v897
	v905 = int32(-2139062144)
	if (v899|(int32(16843008)-v899))&v905 == v905 {
		v893 = v901
		v894 = v899
		v895 = v898
		goto L274
	} else {
		goto L276
	}
L275:
	;
	v910 = v901
	v911 = v899
	v912 = v898
	goto L262
L276:
	;
	goto L275
L277:
	;
	v919 = v910
	v921 = v912
	goto L278
L278:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v921)+1)) = uint8(v922)
	v924 = int32(1)
	if v922 != 0 {
		v919 = v919 + v924
		v921 = v921 + v924
		goto L278
	} else {
		goto L280
	}
L279:
	;
	goto L261
L280:
	;
	goto L279
L281:
	;
	if v769&int32(8) != 0 {
		goto L285
	} else {
		goto L286
	}
L282:
	;
	if v610 == int32(2) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	if base.Ui32(v365) < base.Ui32(v484) {
		v973 = v734
		v974 = int32(0)
		goto L229
	} else {
		goto L284
	}
L284:
	;
	goto L281
L285:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v734))) = uint8(v770)
	v970 = int32(1)
	v973 = v734 + v970
	v974 = v970
	goto L229
L286:
	;
	if l3 != v484 {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v944 != int32(48) {
		goto L285
	} else {
		goto L288
	}
L288:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v947 == int32(0) {
		goto L285
	} else {
		goto L289
	}
L289:
	;
	if v769&int32(32) == int32(0) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v954 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v734))) = uint8(v954)
	v973 = v734 + int32(1)
	v974 = int32(0)
	goto L229
L291:
	;
	goto L292
L292:
	;
	v959 = int32(0)
	if v365 == v959 {
		v973 = v734
		v974 = v959
		goto L229
	} else {
		goto L293
	}
L293:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	if v962 != int32(46) {
		v973 = v734
		v974 = v959
		goto L229
	} else {
		goto L294
	}
L294:
	;
	v965 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v734))) = uint8(v965)
	v973 = v734 + int32(1)
	v974 = v959
	goto L229
L295:
	;
	v992 = v489
	goto L297
L296:
	;
	v992 = v990
	goto L297
L297:
	;
	v994 = v489 + int32(1)
	if v365 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v995 = v992
	goto L300
L299:
	;
	v995 = v990
	goto L300
L300:
	;
	if v994 != v995 {
		v2729 = v979
		v2734 = v983
		v2735 = v984
		v2737 = v736
		v2739 = v994
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L301
	}
L301:
	;
	if v736 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	if v985&int32(64) == int32(0) {
		v2729 = v979
		v2734 = v983
		v2735 = v984
		v2737 = v736
		v2739 = v994
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L305
	}
L303:
	;
	if v985&int32(128) == int32(0) {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v979))) = uint8(v461)
	v1004 = int32(1)
	v2729 = v979 + v1004
	v2734 = v983
	v2735 = v984
	v2737 = v1004
	v2739 = v994
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L305:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1011 != int32(1) {
		v2729 = v979
		v2734 = v983
		v2735 = v984
		v2737 = v736
		v2739 = v994
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L306
	}
L306:
	;
	if (v456^v979)&int32(3) != 0 {
		goto L310
	} else {
		goto L311
	}
L307:
	;
	v1088 = F_strlen(m, v979)
	mBase = m.M
	v2729 = v1088 + v979
	v2734 = v983
	v2735 = v984
	v2737 = v736
	v2739 = v994
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L308:
	;
	goto L307
L309:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1068))) = uint8(v1067)
	if v1067&int32(255) == int32(0) {
		goto L308
	} else {
		goto L324
	}
L310:
	;
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	v1066 = v456
	v1067 = v1019
	v1068 = v979
	goto L309
L311:
	;
	goto L312
L312:
	;
	if v456&int32(3) != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1023 = v456
	v1025 = v979
	goto L316
L314:
	;
	v1037 = v456
	v1039 = v979
	goto L315
L315:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1037)))
	v1044 = int32(-2139062144)
	if (int32(16843008)-v1041|v1041)&v1044 != v1044 {
		v1066 = v1037
		v1067 = v1041
		v1068 = v1039
		goto L309
	} else {
		goto L320
	}
L316:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1023))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1025))) = uint8(v1026)
	if v1026 == int32(0) {
		goto L308
	} else {
		goto L318
	}
L317:
	;
	v1037 = v1033
	v1039 = v1031
	goto L315
L318:
	;
	v1030 = int32(1)
	v1031 = v1025 + v1030
	v1033 = v1023 + v1030
	if v1033&int32(3) != 0 {
		v1023 = v1033
		v1025 = v1031
		goto L316
	} else {
		goto L319
	}
L319:
	;
	goto L317
L320:
	;
	v1049 = v1037
	v1050 = v1041
	v1051 = v1039
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1051))) = v1050
	v1053 = int32(4)
	v1054 = v1051 + v1053
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1049)+4))
	v1057 = v1049 + v1053
	v1061 = int32(-2139062144)
	if (v1055|(int32(16843008)-v1055))&v1061 == v1061 {
		v1049 = v1057
		v1050 = v1055
		v1051 = v1054
		goto L321
	} else {
		goto L323
	}
L322:
	;
	v1066 = v1057
	v1067 = v1055
	v1068 = v1054
	goto L309
L323:
	;
	goto L322
L324:
	;
	v1075 = v1066
	v1077 = v1068
	goto L325
L325:
	;
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1077)+1)) = uint8(v1078)
	v1080 = int32(1)
	if v1078 != 0 {
		v1075 = v1075 + v1080
		v1077 = v1077 + v1080
		goto L325
	} else {
		goto L327
	}
L326:
	;
	goto L308
L327:
	;
	goto L326
L328:
	;
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v1094 = v479 + base.B2i32(v1091 == int32(32))
	if base.Ui32(v465) <= base.Ui32(v1094) {
		v2681 = v1094
		v2686 = v484
		v2696 = v494
		v2697 = v495
		v2704 = v502
		goto L148
	} else {
		goto L329
	}
L329:
	;
	if v610&int32(-2) != int32(2) {
		v1237 = v1094
		goto L330
	} else {
		goto L331
	}
L330:
	;
	if base.Ui32(v465) <= base.Ui32(v1237) {
		v2681 = v1237
		v2686 = v484
		v2696 = v494
		v2697 = v495
		v2704 = v502
		goto L148
	} else {
		goto L380
	}
L331:
	;
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1100 != int32(32) {
		v1237 = v1094
		goto L330
	} else {
		goto L332
	}
L332:
	;
	if v502 != int32(0)-v494 {
		v1237 = v1094
		goto L330
	} else {
		goto L333
	}
L333:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1106&int32(64) == int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1094))))
	if v1220 != int32(45) {
		goto L374
	} else {
		goto L375
	}
L335:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1111 != int32(-1) {
		goto L334
	} else {
		goto L336
	}
L336:
	;
	v1114 = F_strlen(m, v445)
	mBase = m.M
	if v1114 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1167 = F_strlen(m, v444)
	mBase = m.M
	if v1167 == int32(0) {
		v1237 = v1094
		goto L330
	} else {
		goto L355
	}
L338:
	;
	if base.Ui32(l2+(l4-v1114)) < base.Ui32(v1094) {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	if v1114 == int32(0) {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	if v1163 != 0 {
		goto L337
	} else {
		goto L354
	}
L341:
	;
	v1163 = int32(0)
	goto L340
L342:
	;
	goto L343
L343:
	;
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1094))))
	if v1125 != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1126 = v1094
	v1127 = v445
	v1128 = v1114
	v1129 = v1125
	goto L348
L345:
	;
	v1151 = v445
	v1155 = int32(0)
	goto L346
L346:
	;
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151))))
	v1163 = v1155 - v1156
	goto L340
L347:
	;
	v1151 = v1146
	v1155 = v1148
	goto L346
L348:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127))))
	if v1129 != v1131 {
		v1146 = v1127
		v1148 = v1129
		goto L347
	} else {
		goto L350
	}
L349:
	;
	v1146 = v1140
	v1148 = int32(0)
	goto L347
L350:
	;
	if v1131 == int32(0) {
		v1146 = v1127
		v1148 = v1129
		goto L347
	} else {
		goto L351
	}
L351:
	;
	v1136 = v1128 - int32(1)
	if v1136 == int32(0) {
		v1146 = v1127
		v1148 = v1129
		goto L347
	} else {
		goto L352
	}
L352:
	;
	v1139 = int32(1)
	v1140 = v1127 + v1139
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126)+1)))
	if v1141 != 0 {
		v1126 = v1126 + v1139
		v1127 = v1140
		v1128 = v1136
		v1129 = v1141
		goto L348
	} else {
		goto L353
	}
L353:
	;
	goto L349
L354:
	;
	v1164 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1164)
	v1237 = v1094 + v1114
	goto L330
L355:
	;
	if base.Ui32(l2+(l4-v1167)) < base.Ui32(v1094) {
		v1237 = v1094
		goto L330
	} else {
		goto L356
	}
L356:
	;
	if v1167 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	if v1216 != 0 {
		v1237 = v1094
		goto L330
	} else {
		goto L371
	}
L358:
	;
	v1216 = int32(0)
	goto L357
L359:
	;
	goto L360
L360:
	;
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1094))))
	if v1178 != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1179 = v1094
	v1180 = v444
	v1181 = v1167
	v1182 = v1178
	goto L365
L362:
	;
	v1204 = v444
	v1208 = int32(0)
	goto L363
L363:
	;
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1204))))
	v1216 = v1208 - v1209
	goto L357
L364:
	;
	v1204 = v1199
	v1208 = v1201
	goto L363
L365:
	;
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180))))
	if v1182 != v1184 {
		v1199 = v1180
		v1201 = v1182
		goto L364
	} else {
		goto L367
	}
L366:
	;
	v1199 = v1193
	v1201 = int32(0)
	goto L364
L367:
	;
	if v1184 == int32(0) {
		v1199 = v1180
		v1201 = v1182
		goto L364
	} else {
		goto L368
	}
L368:
	;
	v1189 = v1181 - int32(1)
	if v1189 == int32(0) {
		v1199 = v1180
		v1201 = v1182
		goto L364
	} else {
		goto L369
	}
L369:
	;
	v1192 = int32(1)
	v1193 = v1180 + v1192
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179)+1)))
	if v1194 != 0 {
		v1179 = v1179 + v1192
		v1180 = v1193
		v1181 = v1189
		v1182 = v1194
		goto L365
	} else {
		goto L370
	}
L370:
	;
	goto L366
L371:
	;
	v1217 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1217)
	v1237 = v1094 + v1167
	goto L330
L372:
	;
	v1237 = v1094 + int32(1)
	goto L330
L373:
	;
	if v1220 != int32(43) {
		v1237 = v1094
		goto L330
	} else {
		goto L379
	}
L374:
	;
	if v1106&int32(128) == int32(0) {
		goto L373
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	v1229 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1229)
	goto L372
L377:
	;
	if v1220 != int32(60) {
		goto L373
	} else {
		goto L378
	}
L378:
	;
	goto L376
L379:
	;
	v1233 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1233)
	goto L372
L380:
	;
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1237))))
	if base.Ui32((v1241-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	if base.Ui32(v465) <= base.Ui32(v1333) {
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	} else {
		goto L414
	}
L382:
	;
	if v495 != 0 {
		goto L385
	} else {
		goto L386
	}
L383:
	;
	goto L384
L384:
	;
	v1263 = int32(0)
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1264&int32(2) == v1263 {
		v1333 = v1237
		v1334 = v1263
		v1335 = v484
		v1336 = v494
		v1337 = v495
		v1338 = v502
		goto L381
	} else {
		goto L389
	}
L385:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v494 == v1248 {
		v2681 = v1237
		v2686 = v484
		v2696 = v494
		v2697 = v495
		v2704 = v502
		goto L148
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v1241)
	v1257 = int32(1)
	v1333 = v1237
	v1334 = v1257
	v1335 = v484 + v1257
	v1336 = v494
	v1337 = int32(0)
	v1338 = v502 + v1257
	goto L381
L388:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v1241)
	v1251 = int32(1)
	v1333 = v1237
	v1334 = v1251
	v1335 = v484 + v1251
	v1336 = v494 + v1251
	v1337 = v495
	v1338 = v502
	goto L381
L389:
	;
	if v495 != 0 {
		v1333 = v1237
		v1334 = v1263
		v1335 = v484
		v1336 = v494
		v1337 = v495
		v1338 = v502
		goto L381
	} else {
		goto L390
	}
L390:
	;
	v1269 = F_strlen(m, v446)
	mBase = m.M
	if v1269 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1272 = int32(0)
	v1333 = v1237
	v1334 = v1272
	v1335 = v484
	v1336 = v494
	v1337 = v1272
	v1338 = v502
	goto L381
L392:
	;
	goto L393
L393:
	;
	v1274 = int32(0)
	if base.Ui32(l2+(l4-v1269)) < base.Ui32(v1237) {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1333 = v1237
	v1334 = int32(0)
	v1335 = v484
	v1336 = v494
	v1337 = v1274
	v1338 = v502
	goto L381
L395:
	;
	goto L396
L396:
	;
	if v1269 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	if v1322 != 0 {
		goto L411
	} else {
		goto L412
	}
L398:
	;
	v1322 = int32(0)
	goto L397
L399:
	;
	goto L400
L400:
	;
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1237))))
	if v1284 != 0 {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1285 = v1237
	v1286 = v446
	v1287 = v1269
	v1288 = v1284
	goto L405
L402:
	;
	v1310 = v446
	v1314 = int32(0)
	goto L403
L403:
	;
	v1315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1310))))
	v1322 = v1314 - v1315
	goto L397
L404:
	;
	v1310 = v1305
	v1314 = v1307
	goto L403
L405:
	;
	v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1286))))
	if v1288 != v1290 {
		v1305 = v1286
		v1307 = v1288
		goto L404
	} else {
		goto L407
	}
L406:
	;
	v1305 = v1299
	v1307 = int32(0)
	goto L404
L407:
	;
	if v1290 == int32(0) {
		v1305 = v1286
		v1307 = v1288
		goto L404
	} else {
		goto L408
	}
L408:
	;
	v1295 = v1287 - int32(1)
	if v1295 == int32(0) {
		v1305 = v1286
		v1307 = v1288
		goto L404
	} else {
		goto L409
	}
L409:
	;
	v1298 = int32(1)
	v1299 = v1286 + v1298
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1285)+1)))
	if v1300 != 0 {
		v1285 = v1285 + v1298
		v1286 = v1299
		v1287 = v1295
		v1288 = v1300
		goto L405
	} else {
		goto L410
	}
L410:
	;
	goto L406
L411:
	;
	v1333 = v1237
	v1334 = int32(0)
	v1335 = v484
	v1336 = v494
	v1337 = v1274
	v1338 = v502
	goto L381
L412:
	;
	goto L413
L413:
	;
	v1324 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v1324)
	v1326 = int32(1)
	v1333 = v1237 + v1269 - v1326
	v1334 = v1326
	v1335 = v484 + v1326
	v1336 = v494
	v1337 = v1326
	v1338 = v502
	goto L381
L414:
	;
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1340 != int32(32) {
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	} else {
		goto L415
	}
L415:
	;
	if v1336+v1338 <= int32(0) {
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	} else {
		goto L416
	}
L416:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1348 = v1346 & int32(64)
	v1349 = int32(0)
	if base.B2i32(v1348 == v1349)|(v1334^int32(1)) == v1349 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1357 = v1333 + int32(1)
	if base.Ui32(v465) <= base.Ui32(v1357) {
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	} else {
		goto L420
	}
L418:
	;
	goto L419
L419:
	;
	if v1334|(v1348|base.B2i32(v1346&int32(768) == int32(0))) != 0 {
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	} else {
		goto L457
	}
L420:
	;
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1357))))
	if base.Ui32((v1359-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	} else {
		goto L421
	}
L421:
	;
	v1366 = F_strlen(m, v445)
	mBase = m.M
	if v1366 == int32(0) {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1419 = F_strlen(m, v444)
	mBase = m.M
	if v1419 == int32(0) {
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	} else {
		goto L440
	}
L423:
	;
	if base.Ui32(l2+(l4-v1366)) < base.Ui32(v1357) {
		goto L422
	} else {
		goto L424
	}
L424:
	;
	if v1366 == int32(0) {
		goto L426
	} else {
		goto L427
	}
L425:
	;
	if v1415 != 0 {
		goto L422
	} else {
		goto L439
	}
L426:
	;
	v1415 = int32(0)
	goto L425
L427:
	;
	goto L428
L428:
	;
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1357))))
	if v1377 != 0 {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v1378 = v1357
	v1379 = v445
	v1380 = v1366
	v1381 = v1377
	goto L433
L430:
	;
	v1403 = v445
	v1407 = int32(0)
	goto L431
L431:
	;
	v1408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403))))
	v1415 = v1407 - v1408
	goto L425
L432:
	;
	v1403 = v1398
	v1407 = v1400
	goto L431
L433:
	;
	v1383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1379))))
	if v1381 != v1383 {
		v1398 = v1379
		v1400 = v1381
		goto L432
	} else {
		goto L435
	}
L434:
	;
	v1398 = v1392
	v1400 = int32(0)
	goto L432
L435:
	;
	if v1383 == int32(0) {
		v1398 = v1379
		v1400 = v1381
		goto L432
	} else {
		goto L436
	}
L436:
	;
	v1388 = v1380 - int32(1)
	if v1388 == int32(0) {
		v1398 = v1379
		v1400 = v1381
		goto L432
	} else {
		goto L437
	}
L437:
	;
	v1391 = int32(1)
	v1392 = v1379 + v1391
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378)+1)))
	if v1393 != 0 {
		v1378 = v1378 + v1391
		v1379 = v1392
		v1380 = v1388
		v1381 = v1393
		goto L433
	} else {
		goto L438
	}
L438:
	;
	goto L434
L439:
	;
	v1417 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1417)
	v2681 = v1333 + v1366
	v2686 = v1335
	v2696 = v1336
	v2697 = v1337
	v2704 = v1338
	goto L148
L440:
	;
	if base.Ui32(l2+(l4-v1419)) < base.Ui32(v1357) {
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	} else {
		goto L441
	}
L441:
	;
	if v1419 == int32(0) {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	if v1468 != 0 {
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	} else {
		goto L456
	}
L443:
	;
	v1468 = int32(0)
	goto L442
L444:
	;
	goto L445
L445:
	;
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1357))))
	if v1430 != 0 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1431 = v1357
	v1432 = v444
	v1433 = v1419
	v1434 = v1430
	goto L450
L447:
	;
	v1456 = v444
	v1460 = int32(0)
	goto L448
L448:
	;
	v1461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1456))))
	v1468 = v1460 - v1461
	goto L442
L449:
	;
	v1456 = v1451
	v1460 = v1453
	goto L448
L450:
	;
	v1436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1432))))
	if v1434 != v1436 {
		v1451 = v1432
		v1453 = v1434
		goto L449
	} else {
		goto L452
	}
L451:
	;
	v1451 = v1445
	v1453 = int32(0)
	goto L449
L452:
	;
	if v1436 == int32(0) {
		v1451 = v1432
		v1453 = v1434
		goto L449
	} else {
		goto L453
	}
L453:
	;
	v1441 = v1433 - int32(1)
	if v1441 == int32(0) {
		v1451 = v1432
		v1453 = v1434
		goto L449
	} else {
		goto L454
	}
L454:
	;
	v1444 = int32(1)
	v1445 = v1432 + v1444
	v1446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1431)+1)))
	if v1446 != 0 {
		v1431 = v1431 + v1444
		v1432 = v1445
		v1433 = v1441
		v1434 = v1446
		goto L450
	} else {
		goto L455
	}
L455:
	;
	goto L451
L456:
	;
	v1472 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1472)
	v2681 = v1357 + v1419 - int32(1)
	v2686 = v1335
	v2696 = v1336
	v2697 = v1337
	v2704 = v1338
	goto L148
L457:
	;
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1333))))
	switch v1480 - int32(43) {
	case 0, 2:
		goto L149
	default:
		v2681 = v1333
		v2686 = v1335
		v2696 = v1336
		v2697 = v1337
		v2704 = v1338
		goto L148
	}
L458:
	;
	if v485 == int32(0) {
		goto L461
	} else {
		goto L462
	}
L459:
	;
	goto L460
L460:
	;
	if v485 != 0 {
		goto L467
	} else {
		goto L468
	}
L461:
	;
	v1485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1485&int32(32) != 0 {
		goto L464
	} else {
		goto L465
	}
L462:
	;
	goto L463
L463:
	;
	v1491 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v1491)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L464:
	;
	v2729 = v479
	v2734 = v484
	v2735 = int32(0)
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L465:
	;
	goto L466
L466:
	;
	v1489 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v1489)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L467:
	;
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v1499 == int32(44) {
		v2681 = v479
		v2686 = v484
		v2696 = v494
		v2697 = v495
		v2704 = v502
		goto L148
	} else {
		goto L470
	}
L468:
	;
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1493&int32(32) == int32(0) {
		goto L467
	} else {
		goto L469
	}
L469:
	;
	v2729 = v479
	v2734 = v484
	v2735 = int32(0)
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L470:
	;
	v2729 = v479
	v2734 = v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L471:
	;
	if v485 == int32(0) {
		goto L474
	} else {
		goto L475
	}
L472:
	;
	goto L473
L473:
	;
	if v485 != 0 {
		goto L503
	} else {
		goto L504
	}
L474:
	;
	v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1505&int32(32) != 0 {
		goto L477
	} else {
		goto L478
	}
L475:
	;
	goto L476
L476:
	;
	if (v443^v479)&int32(3) != 0 {
		goto L485
	} else {
		goto L486
	}
L477:
	;
	v2729 = v479
	v2734 = v484
	v2735 = int32(0)
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L478:
	;
	goto L479
L479:
	;
	v1510 = F_pg_mbstrlen(m, v443)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L71
	} else {
		goto L480
	}
L480:
	;
	v1513 = F__emscripten_memset_bulkmem(m, v479, base.I32_extend8_s(int32(32)), v1510)
	mBase = m.M
	goto L481
L481:
	;
	v2681 = v1513 + v1510 - int32(1)
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L482:
	;
	v2681 = v479 + v1502 - int32(1)
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L483:
	;
	goto L482
L484:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1571))) = uint8(v1570)
	if v1570&int32(255) == int32(0) {
		goto L483
	} else {
		goto L499
	}
L485:
	;
	v1522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	v1569 = v443
	v1570 = v1522
	v1571 = v479
	goto L484
L486:
	;
	goto L487
L487:
	;
	if v443&int32(3) != 0 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v1526 = v443
	v1528 = v479
	goto L491
L489:
	;
	v1540 = v443
	v1542 = v479
	goto L490
L490:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1540)))
	v1547 = int32(-2139062144)
	if (int32(16843008)-v1544|v1544)&v1547 != v1547 {
		v1569 = v1540
		v1570 = v1544
		v1571 = v1542
		goto L484
	} else {
		goto L495
	}
L491:
	;
	v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1526))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1528))) = uint8(v1529)
	if v1529 == int32(0) {
		goto L483
	} else {
		goto L493
	}
L492:
	;
	v1540 = v1536
	v1542 = v1534
	goto L490
L493:
	;
	v1533 = int32(1)
	v1534 = v1528 + v1533
	v1536 = v1526 + v1533
	if v1536&int32(3) != 0 {
		v1526 = v1536
		v1528 = v1534
		goto L491
	} else {
		goto L494
	}
L494:
	;
	goto L492
L495:
	;
	v1552 = v1540
	v1553 = v1544
	v1554 = v1542
	goto L496
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1554))) = v1553
	v1556 = int32(4)
	v1557 = v1554 + v1556
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+4))
	v1560 = v1552 + v1556
	v1564 = int32(-2139062144)
	if (v1558|(int32(16843008)-v1558))&v1564 == v1564 {
		v1552 = v1560
		v1553 = v1558
		v1554 = v1557
		goto L496
	} else {
		goto L498
	}
L497:
	;
	v1569 = v1560
	v1570 = v1558
	v1571 = v1557
	goto L484
L498:
	;
	goto L497
L499:
	;
	v1578 = v1569
	v1580 = v1571
	goto L500
L500:
	;
	v1581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1578)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1580)+1)) = uint8(v1581)
	v1583 = int32(1)
	if v1581 != 0 {
		v1578 = v1578 + v1583
		v1580 = v1580 + v1583
		goto L500
	} else {
		goto L502
	}
L501:
	;
	goto L483
L502:
	;
	goto L501
L503:
	;
	if base.Ui32(l2+(l4-v1502)) < base.Ui32(v479) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L506
	}
L504:
	;
	v1594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1594&int32(32) == int32(0) {
		goto L503
	} else {
		goto L505
	}
L505:
	;
	v2729 = v479
	v2734 = v484
	v2735 = int32(0)
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L506:
	;
	if v1502 == int32(0) {
		goto L508
	} else {
		goto L509
	}
L507:
	;
	if v1646 != 0 {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L521
	}
L508:
	;
	v1646 = int32(0)
	goto L507
L509:
	;
	goto L510
L510:
	;
	v1608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v1608 != 0 {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v1609 = v479
	v1610 = v443
	v1611 = v1502
	v1612 = v1608
	goto L515
L512:
	;
	v1634 = v443
	v1638 = int32(0)
	goto L513
L513:
	;
	v1639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1634))))
	v1646 = v1638 - v1639
	goto L507
L514:
	;
	v1634 = v1629
	v1638 = v1631
	goto L513
L515:
	;
	v1614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1610))))
	if v1612 != v1614 {
		v1629 = v1610
		v1631 = v1612
		goto L514
	} else {
		goto L517
	}
L516:
	;
	v1629 = v1623
	v1631 = int32(0)
	goto L514
L517:
	;
	if v1614 == int32(0) {
		v1629 = v1610
		v1631 = v1612
		goto L514
	} else {
		goto L518
	}
L518:
	;
	v1619 = v1611 - int32(1)
	if v1619 == int32(0) {
		v1629 = v1610
		v1631 = v1612
		goto L514
	} else {
		goto L519
	}
L519:
	;
	v1622 = int32(1)
	v1623 = v1610 + v1622
	v1624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1609)+1)))
	if v1624 != 0 {
		v1609 = v1609 + v1622
		v1610 = v1623
		v1611 = v1619
		v1612 = v1624
		goto L515
	} else {
		goto L520
	}
L520:
	;
	goto L516
L521:
	;
	v2681 = v479 + v1502 - int32(1)
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L522:
	;
	if (v447^v479)&int32(3) != 0 {
		goto L528
	} else {
		goto L529
	}
L523:
	;
	goto L524
L524:
	;
	v1728 = F_pg_mbstrlen(m, v447)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L71
	} else {
		goto L546
	}
L525:
	;
	v1724 = F_strlen(m, v447)
	mBase = m.M
	v2681 = v479 + v1724 - int32(1)
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L526:
	;
	goto L525
L527:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1704))) = uint8(v1703)
	if v1703&int32(255) == int32(0) {
		goto L526
	} else {
		goto L542
	}
L528:
	;
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447))))
	v1702 = v447
	v1703 = v1655
	v1704 = v479
	goto L527
L529:
	;
	goto L530
L530:
	;
	if v447&int32(3) != 0 {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	v1659 = v447
	v1661 = v479
	goto L534
L532:
	;
	v1673 = v447
	v1675 = v479
	goto L533
L533:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1673)))
	v1680 = int32(-2139062144)
	if (int32(16843008)-v1677|v1677)&v1680 != v1680 {
		v1702 = v1673
		v1703 = v1677
		v1704 = v1675
		goto L527
	} else {
		goto L538
	}
L534:
	;
	v1662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1659))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1661))) = uint8(v1662)
	if v1662 == int32(0) {
		goto L526
	} else {
		goto L536
	}
L535:
	;
	v1673 = v1669
	v1675 = v1667
	goto L533
L536:
	;
	v1666 = int32(1)
	v1667 = v1661 + v1666
	v1669 = v1659 + v1666
	if v1669&int32(3) != 0 {
		v1659 = v1669
		v1661 = v1667
		goto L534
	} else {
		goto L537
	}
L537:
	;
	goto L535
L538:
	;
	v1685 = v1673
	v1686 = v1677
	v1687 = v1675
	goto L539
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1687))) = v1686
	v1689 = int32(4)
	v1690 = v1687 + v1689
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1685)+4))
	v1693 = v1685 + v1689
	v1697 = int32(-2139062144)
	if (v1691|(int32(16843008)-v1691))&v1697 == v1697 {
		v1685 = v1693
		v1686 = v1691
		v1687 = v1690
		goto L539
	} else {
		goto L541
	}
L540:
	;
	v1702 = v1693
	v1703 = v1691
	v1704 = v1690
	goto L527
L541:
	;
	goto L540
L542:
	;
	v1711 = v1702
	v1713 = v1704
	goto L543
L543:
	;
	v1714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1711)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1713)+1)) = uint8(v1714)
	v1716 = int32(1)
	if v1714 != 0 {
		v1711 = v1711 + v1716
		v1713 = v1713 + v1716
		goto L543
	} else {
		goto L545
	}
L544:
	;
	goto L526
L545:
	;
	goto L544
L546:
	;
	if v1728 <= int32(0) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L547
	}
L547:
	;
	v1740 = v479
	v1741 = v1728
	goto L548
L548:
	;
	if base.Ui32(v465) <= base.Ui32(v1740) {
		v2729 = v1740
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L550
	}
L549:
	;
	v2729 = v1790
	v2734 = v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L550:
	;
	v1779 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1740))))
	if base.Ui64(v1779) <= base.Ui64(int64(63)) {
		goto L551
	} else {
		goto L552
	}
L551:
	;
	if int64(1)<<(uint(v1779)%64)&int64(288080842570334209) != int64(0) {
		v2729 = v1740
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	v1788 = F_pg_mblen_range(m, v1740, v465)
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L71
	} else {
		goto L555
	}
L554:
	;
	goto L553
L555:
	;
	v1790 = v1788 + v1740
	v1791 = int32(1)
	if base.Ui32(v1791) < base.Ui32(v1741) {
		v1740 = v1790
		v1741 = v1741 - v1791
		goto L548
	} else {
		goto L556
	}
L556:
	;
	goto L549
L557:
	;
	if v610 != int32(30) {
		v1878 = v484
		goto L560
	} else {
		goto L561
	}
L558:
	;
	goto L559
L559:
	;
	if base.Ui32(v465) <= base.Ui32(v479) {
		v2063 = v479
		v2101 = v479
		goto L600
	} else {
		goto L601
	}
L560:
	;
	v1913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1913&int32(32) != 0 {
		goto L575
	} else {
		goto L576
	}
L561:
	;
	if v484 == int32(0) {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	v1878 = int32(0)
	goto L560
L563:
	;
	goto L564
L564:
	;
	v1800 = F_strlen(m, v484)
	mBase = m.M
	v1801 = F_pnstrdup(m, v484, v1800)
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L71
	} else {
		goto L565
	}
L565:
	;
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1801))))
	if v1803 == int32(0) {
		v1878 = v1801
		goto L560
	} else {
		goto L566
	}
L566:
	;
	v1815 = v1801
	v1816 = v1803
	goto L567
L567:
	;
	v1852 = int32(255)
	v1853 = v1816 & v1852
	if base.Ui32((v1853-int32(65))&v1852) < base.Ui32(int32(26)) {
		goto L570
	} else {
		goto L571
	}
L568:
	;
	v1878 = v1801
	goto L560
L569:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1815))) = uint8(v1862)
	v1865 = v1815 + int32(1)
	v1866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865))))
	if v1866 != 0 {
		v1815 = v1865
		v1816 = v1866
		goto L567
	} else {
		goto L573
	}
L570:
	;
	v1862 = v1853 | int32(32)
	goto L572
L571:
	;
	v1862 = v1853
	goto L572
L572:
	;
	goto L569
L573:
	;
	goto L568
L574:
	;
	v1994 = F_strlen(m, v479)
	mBase = m.M
	v2681 = v1994 + v479 - int32(1)
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L575:
	;
	if (v1878^v479)&int32(3) != 0 {
		goto L581
	} else {
		goto L582
	}
L576:
	;
	goto L577
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v1878
	v1992 = F_pg_sprintf(m, v479, int32(174966), v49)
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L71
	} else {
		goto L599
	}
L578:
	;
	goto L574
L579:
	;
	goto L578
L580:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1970))) = uint8(v1969)
	if v1969&int32(255) == int32(0) {
		goto L579
	} else {
		goto L595
	}
L581:
	;
	v1921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1878))))
	v1968 = v1878
	v1969 = v1921
	v1970 = v479
	goto L580
L582:
	;
	goto L583
L583:
	;
	if v1878&int32(3) != 0 {
		goto L584
	} else {
		goto L585
	}
L584:
	;
	v1925 = v1878
	v1927 = v479
	goto L587
L585:
	;
	v1939 = v1878
	v1941 = v479
	goto L586
L586:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1939)))
	v1946 = int32(-2139062144)
	if (int32(16843008)-v1943|v1943)&v1946 != v1946 {
		v1968 = v1939
		v1969 = v1943
		v1970 = v1941
		goto L580
	} else {
		goto L591
	}
L587:
	;
	v1928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1925))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1927))) = uint8(v1928)
	if v1928 == int32(0) {
		goto L579
	} else {
		goto L589
	}
L588:
	;
	v1939 = v1935
	v1941 = v1933
	goto L586
L589:
	;
	v1932 = int32(1)
	v1933 = v1927 + v1932
	v1935 = v1925 + v1932
	if v1935&int32(3) != 0 {
		v1925 = v1935
		v1927 = v1933
		goto L587
	} else {
		goto L590
	}
L590:
	;
	goto L588
L591:
	;
	v1951 = v1939
	v1952 = v1943
	v1953 = v1941
	goto L592
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1953))) = v1952
	v1955 = int32(4)
	v1956 = v1953 + v1955
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1951)+4))
	v1959 = v1951 + v1955
	v1963 = int32(-2139062144)
	if (v1957|(int32(16843008)-v1957))&v1963 == v1963 {
		v1951 = v1959
		v1952 = v1957
		v1953 = v1956
		goto L592
	} else {
		goto L594
	}
L593:
	;
	v1968 = v1959
	v1969 = v1957
	v1970 = v1956
	goto L580
L594:
	;
	goto L593
L595:
	;
	v1977 = v1968
	v1979 = v1970
	goto L596
L596:
	;
	v1980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1977)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1979)+1)) = uint8(v1980)
	v1982 = int32(1)
	if v1980 != 0 {
		v1977 = v1977 + v1982
		v1979 = v1979 + v1982
		goto L596
	} else {
		goto L598
	}
L597:
	;
	goto L579
L598:
	;
	goto L597
L599:
	;
	goto L574
L600:
	;
	v2111 = v2063
	v2112 = int32(0)
	v2113 = v2101
	goto L609
L601:
	;
	v2007 = v479
	goto L602
L602:
	;
	v2045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2007))))
	if base.Ui32(v2045-int32(9)) < base.Ui32(int32(5)) {
		goto L604
	} else {
		goto L605
	}
L603:
	;
	v2063 = v2053
	v2101 = v465
	goto L600
L604:
	;
	v2053 = v2007 + int32(1)
	if v2053 != v465 {
		v2007 = v2053
		goto L602
	} else {
		goto L607
	}
L605:
	;
	if v2045 == int32(32) {
		goto L604
	} else {
		goto L606
	}
L606:
	;
	v2063 = v2007
	v2101 = v2007
	goto L600
L607:
	;
	goto L603
L608:
	;
	v2198 = int32(1)
	v2200 = int32(0)
	v2217 = v2200
	v2233 = v2200
	v2234 = v2200
	v2235 = v2200
	v2236 = v2198
	v2240 = v2200
	v2246 = v2200
	v2247 = v2200
	goto L626
L609:
	;
	if base.Ui32(v465) <= base.Ui32(v2113) {
		goto L611
	} else {
		goto L612
	}
L610:
	;
	if v2112 == int32(0) {
		goto L5
	} else {
		goto L625
	}
L611:
	;
	goto L610
L612:
	;
	v2151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2113))))
	if base.Ui32((v2151-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L621
	} else {
		goto L622
	}
L613:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(97)+v2112))) = uint8(v2162)
	*(*int32)(unsafe.Add(mBase, uint32(v49+int32(32)+v2112<<(uint(int32(2))%32)))) = v2171
	v2182 = int32(15)
	v2183 = int32(1)
	v2184 = v2111 + v2183
	v2186 = v2112 + v2183
	if v2186 != v2182 {
		v2111 = v2184
		v2112 = v2186
		v2113 = v2184
		goto L609
	} else {
		goto L624
	}
L614:
	;
	v2171 = int32(1000)
	goto L613
L615:
	;
	v2171 = int32(500)
	goto L613
L616:
	;
	v2171 = int32(100)
	goto L613
L617:
	;
	v2171 = int32(50)
	goto L613
L618:
	;
	v2171 = int32(10)
	goto L613
L619:
	;
	v2171 = int32(5)
	goto L613
L620:
	;
	switch v2162 - int32(67) {
	case 0:
		goto L616
	case 1:
		goto L615
	default:
		goto L611
	case 6:
		v2171 = int32(1)
		goto L613
	case 9:
		goto L617
	case 10:
		goto L614
	case 19:
		goto L619
	case 21:
		goto L618
	}
L621:
	;
	v2160 = v2151 - int32(32)
	goto L623
L622:
	;
	v2160 = v2151
	goto L623
L623:
	;
	v2162 = v2160 & int32(255)
	goto L620
L624:
	;
	v2193 = v2184
	v2197 = v2182
	goto L608
L625:
	;
	v2193 = v2111
	v2197 = v2112
	goto L608
L626:
	;
	v2257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(97)+v2217))))
	v2263 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(32)+v2217<<(uint(int32(2))%32))))
	if base.B2i32(v2246 <= v2263)&v2247 != 0 {
		goto L5
	} else {
		goto L628
	}
L627:
	;
	if v2355 < int32(0) {
		goto L5
	} else {
		goto L678
	}
L628:
	;
	if int32(4) < v2263 {
		goto L629
	} else {
		goto L630
	}
L629:
	;
	v2269 = v2233
	goto L631
L630:
	;
	v2269 = int32(0)
	goto L631
L631:
	;
	if v2269 != 0 {
		goto L5
	} else {
		goto L632
	}
L632:
	;
	if int32(49) < v2263 {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v2273 = v2234
	goto L635
L634:
	;
	v2273 = int32(0)
	goto L635
L635:
	;
	if v2273 != 0 {
		goto L5
	} else {
		goto L636
	}
L636:
	;
	if int32(499) < v2263 {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v2277 = v2235
	goto L639
L638:
	;
	v2277 = int32(0)
	goto L639
L639:
	;
	if v2277 != 0 {
		goto L5
	} else {
		goto L640
	}
L640:
	;
	switch v2257 - int32(68) {
	case 0:
		goto L642
	default:
		v2286 = v2233
		v2287 = v2234
		v2288 = v2235
		goto L641
	case 8:
		goto L643
	case 18:
		goto L644
	}
L641:
	;
	if v2197-v2198 <= v2217 {
		goto L646
	} else {
		goto L647
	}
L642:
	;
	v2286 = v2233
	v2287 = v2234
	v2288 = v2235 + int32(1)
	goto L641
L643:
	;
	v2286 = v2233
	v2287 = v2234 + int32(1)
	v2288 = v2235
	goto L641
L644:
	;
	v2286 = v2233 + int32(1)
	v2287 = v2234
	v2288 = v2235
	goto L641
L645:
	;
	v2355 = v2345 + v2240
	v2357 = v2346 + int32(1)
	if v2357 < v2197 {
		v2217 = v2357
		v2233 = v2347
		v2234 = v2348
		v2235 = v2349
		v2236 = v2350
		v2240 = v2355
		v2246 = v2353
		v2247 = v2354
		goto L626
	} else {
		goto L677
	}
L646:
	;
	v2345 = v2263
	v2346 = v2217
	v2347 = v2286
	v2348 = v2287
	v2349 = v2288
	v2350 = v2236
	v2353 = v2246
	v2354 = v2247
	goto L645
L647:
	;
	goto L648
L648:
	;
	v2291 = v2217 + int32(1)
	v2295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2291+(v49+int32(97))))))
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(32)+v2291<<(uint(int32(2))%32))))
	if v2263 < v2301 {
		goto L649
	} else {
		goto L650
	}
L649:
	;
	switch v2257 - int32(67) {
	case 0:
		goto L653
	default:
		goto L5
	case 6:
		goto L655
	case 21:
		goto L654
	}
L650:
	;
	goto L651
L651:
	;
	if v2295 != v2257 {
		goto L673
	} else {
		goto L674
	}
L652:
	;
	if int32(1) < v2236 {
		goto L5
	} else {
		goto L656
	}
L653:
	;
	switch v2295 - int32(68) {
	case 0, 9:
		goto L652
	default:
		goto L5
	}
L654:
	;
	switch v2295 - int32(67) {
	case 0, 9:
		goto L652
	default:
		goto L5
	}
L655:
	;
	switch v2295 - int32(86) {
	case 0, 2:
		goto L652
	default:
		goto L5
	}
L656:
	;
	if int32(4) < v2301 {
		goto L657
	} else {
		goto L658
	}
L657:
	;
	v2316 = v2286
	goto L659
L658:
	;
	v2316 = int32(0)
	goto L659
L659:
	;
	if v2316 != 0 {
		goto L5
	} else {
		goto L660
	}
L660:
	;
	if int32(49) < v2301 {
		goto L661
	} else {
		goto L662
	}
L661:
	;
	v2320 = v2287
	goto L663
L662:
	;
	v2320 = int32(0)
	goto L663
L663:
	;
	if v2320 != 0 {
		goto L5
	} else {
		goto L664
	}
L664:
	;
	if int32(499) < v2301 {
		goto L665
	} else {
		goto L666
	}
L665:
	;
	v2324 = v2288
	goto L667
L666:
	;
	v2324 = int32(0)
	goto L667
L667:
	;
	if v2324 != 0 {
		goto L5
	} else {
		goto L668
	}
L668:
	;
	switch v2295 - int32(68) {
	case 0:
		goto L670
	default:
		v2333 = v2286
		v2334 = v2287
		v2335 = v2288
		goto L669
	case 8:
		goto L671
	case 18:
		goto L672
	}
L669:
	;
	v2337 = int32(1)
	v2345 = v2301 - v2263
	v2346 = v2291
	v2347 = v2333
	v2348 = v2334
	v2349 = v2335
	v2350 = v2337
	v2353 = v2263
	v2354 = v2337
	goto L645
L670:
	;
	v2333 = v2286
	v2334 = v2287
	v2335 = v2288 + int32(1)
	goto L669
L671:
	;
	v2333 = v2286
	v2334 = v2287 + int32(1)
	v2335 = v2288
	goto L669
L672:
	;
	v2333 = v2286 + int32(1)
	v2334 = v2287
	v2335 = v2288
	goto L669
L673:
	;
	v2345 = v2263
	v2346 = v2217
	v2347 = v2286
	v2348 = v2287
	v2349 = v2288
	v2350 = int32(1)
	v2353 = v2246
	v2354 = v2247
	goto L645
L674:
	;
	goto L675
L675:
	;
	v2342 = v2236 + int32(1)
	if int32(3) < v2342 {
		goto L5
	} else {
		goto L676
	}
L676:
	;
	v2345 = v2263
	v2346 = v2217
	v2347 = v2286
	v2348 = v2287
	v2349 = v2288
	v2350 = v2342
	v2353 = v2246
	v2354 = v2247
	goto L645
L677:
	;
	goto L627
L678:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v2355
	v2365 = F_pg_sprintf(m, v484, int32(487662), v49+int32(16))
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L71
	} else {
		goto L679
	}
L679:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2365
	v2729 = v2193
	v2734 = v2365 + v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L680:
	;
	if v2371&int32(2) != 0 {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L681
	}
L681:
	;
	v2376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v2376 == int32(35) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L682
	}
L682:
	;
	if v356 == int32(45) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L683
	}
L683:
	;
	if l7 != 0 {
		goto L684
	} else {
		goto L685
	}
L684:
	;
	v2380 = F_get_th(m, l3, int32(2))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L71
	} else {
		goto L687
	}
L685:
	;
	goto L686
L686:
	;
	if base.Ui32(v465) <= base.Ui32(v479) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L709
	}
L687:
	;
	if (v2380^v479)&int32(3) != 0 {
		goto L691
	} else {
		goto L692
	}
L688:
	;
	v2681 = v479 + int32(1)
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L689:
	;
	goto L688
L690:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2436))) = uint8(v2435)
	if v2435&int32(255) == int32(0) {
		goto L689
	} else {
		goto L705
	}
L691:
	;
	v2387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2380))))
	v2434 = v2380
	v2435 = v2387
	v2436 = v479
	goto L690
L692:
	;
	goto L693
L693:
	;
	if v2380&int32(3) != 0 {
		goto L694
	} else {
		goto L695
	}
L694:
	;
	v2391 = v2380
	v2393 = v479
	goto L697
L695:
	;
	v2405 = v2380
	v2407 = v479
	goto L696
L696:
	;
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(v2405)))
	v2412 = int32(-2139062144)
	if (int32(16843008)-v2409|v2409)&v2412 != v2412 {
		v2434 = v2405
		v2435 = v2409
		v2436 = v2407
		goto L690
	} else {
		goto L701
	}
L697:
	;
	v2394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2391))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2393))) = uint8(v2394)
	if v2394 == int32(0) {
		goto L689
	} else {
		goto L699
	}
L698:
	;
	v2405 = v2401
	v2407 = v2399
	goto L696
L699:
	;
	v2398 = int32(1)
	v2399 = v2393 + v2398
	v2401 = v2391 + v2398
	if v2401&int32(3) != 0 {
		v2391 = v2401
		v2393 = v2399
		goto L697
	} else {
		goto L700
	}
L700:
	;
	goto L698
L701:
	;
	v2417 = v2405
	v2418 = v2409
	v2419 = v2407
	goto L702
L702:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2419))) = v2418
	v2421 = int32(4)
	v2422 = v2419 + v2421
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(v2417)+4))
	v2425 = v2417 + v2421
	v2429 = int32(-2139062144)
	if (v2423|(int32(16843008)-v2423))&v2429 == v2429 {
		v2417 = v2425
		v2418 = v2423
		v2419 = v2422
		goto L702
	} else {
		goto L704
	}
L703:
	;
	v2434 = v2425
	v2435 = v2423
	v2436 = v2422
	goto L690
L704:
	;
	goto L703
L705:
	;
	v2443 = v2434
	v2445 = v2436
	goto L706
L706:
	;
	v2446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2443)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2445)+1)) = uint8(v2446)
	v2448 = int32(1)
	if v2446 != 0 {
		v2443 = v2443 + v2448
		v2445 = v2445 + v2448
		goto L706
	} else {
		goto L708
	}
L707:
	;
	goto L689
L708:
	;
	goto L707
L709:
	;
	v2459 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if base.Ui64(v2459) <= base.Ui64(int64(63)) {
		goto L710
	} else {
		goto L711
	}
L710:
	;
	if int64(1)<<(uint(v2459)%64)&int64(288080842570334209) != int64(0) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L713
	}
L711:
	;
	goto L712
L712:
	;
	v2468 = F_pg_mblen_range(m, v479, v465)
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L71
	} else {
		goto L714
	}
L713:
	;
	goto L712
L714:
	;
	v2470 = v2468 + v479
	if base.Ui32(v465) <= base.Ui32(v2470) {
		v2729 = v2470
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L715
	}
L715:
	;
	v2472 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2470))))
	if base.Ui64(v2472) <= base.Ui64(int64(63)) {
		goto L716
	} else {
		goto L717
	}
L716:
	;
	if int64(1)<<(uint(v2472)%64)&int64(288080842570334209) != int64(0) {
		v2729 = v2470
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L719
	}
L717:
	;
	goto L718
L718:
	;
	v2481 = F_pg_mblen_range(m, v2470, v465)
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L71
	} else {
		goto L720
	}
L719:
	;
	goto L718
L720:
	;
	v2729 = v2481 + v2470
	v2734 = v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L721:
	;
	if v2484&int32(2) != 0 {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L722
	}
L722:
	;
	v2489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v2489 == int32(35) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L723
	}
L723:
	;
	if v356 == int32(45) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L724
	}
L724:
	;
	if l7 != 0 {
		goto L725
	} else {
		goto L726
	}
L725:
	;
	v2493 = F_get_th(m, l3, int32(1))
	mBase = m.M
	v2494 = m.ExcPending
	if v2494 != 0 {
		goto L71
	} else {
		goto L728
	}
L726:
	;
	goto L727
L727:
	;
	if base.Ui32(v465) <= base.Ui32(v479) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L750
	}
L728:
	;
	if (v2493^v479)&int32(3) != 0 {
		goto L732
	} else {
		goto L733
	}
L729:
	;
	v2681 = v479 + int32(1)
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L730:
	;
	goto L729
L731:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2549))) = uint8(v2548)
	if v2548&int32(255) == int32(0) {
		goto L730
	} else {
		goto L746
	}
L732:
	;
	v2500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2493))))
	v2547 = v2493
	v2548 = v2500
	v2549 = v479
	goto L731
L733:
	;
	goto L734
L734:
	;
	if v2493&int32(3) != 0 {
		goto L735
	} else {
		goto L736
	}
L735:
	;
	v2504 = v2493
	v2506 = v479
	goto L738
L736:
	;
	v2518 = v2493
	v2520 = v479
	goto L737
L737:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v2518)))
	v2525 = int32(-2139062144)
	if (int32(16843008)-v2522|v2522)&v2525 != v2525 {
		v2547 = v2518
		v2548 = v2522
		v2549 = v2520
		goto L731
	} else {
		goto L742
	}
L738:
	;
	v2507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2504))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2506))) = uint8(v2507)
	if v2507 == int32(0) {
		goto L730
	} else {
		goto L740
	}
L739:
	;
	v2518 = v2514
	v2520 = v2512
	goto L737
L740:
	;
	v2511 = int32(1)
	v2512 = v2506 + v2511
	v2514 = v2504 + v2511
	if v2514&int32(3) != 0 {
		v2504 = v2514
		v2506 = v2512
		goto L738
	} else {
		goto L741
	}
L741:
	;
	goto L739
L742:
	;
	v2530 = v2518
	v2531 = v2522
	v2532 = v2520
	goto L743
L743:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2532))) = v2531
	v2534 = int32(4)
	v2535 = v2532 + v2534
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(v2530)+4))
	v2538 = v2530 + v2534
	v2542 = int32(-2139062144)
	if (v2536|(int32(16843008)-v2536))&v2542 == v2542 {
		v2530 = v2538
		v2531 = v2536
		v2532 = v2535
		goto L743
	} else {
		goto L745
	}
L744:
	;
	v2547 = v2538
	v2548 = v2536
	v2549 = v2535
	goto L731
L745:
	;
	goto L744
L746:
	;
	v2556 = v2547
	v2558 = v2549
	goto L747
L747:
	;
	v2559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2556)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2558)+1)) = uint8(v2559)
	v2561 = int32(1)
	if v2559 != 0 {
		v2556 = v2556 + v2561
		v2558 = v2558 + v2561
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
	v2572 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if base.Ui64(v2572) <= base.Ui64(int64(63)) {
		goto L751
	} else {
		goto L752
	}
L751:
	;
	if int64(1)<<(uint(v2572)%64)&int64(288080842570334209) != int64(0) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L754
	}
L752:
	;
	goto L753
L753:
	;
	v2581 = F_pg_mblen_range(m, v479, v465)
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L71
	} else {
		goto L755
	}
L754:
	;
	goto L753
L755:
	;
	v2583 = v2581 + v479
	if base.Ui32(v465) <= base.Ui32(v2583) {
		v2729 = v2583
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L756
	}
L756:
	;
	v2585 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2583))))
	if base.Ui64(v2585) <= base.Ui64(int64(63)) {
		goto L757
	} else {
		goto L758
	}
L757:
	;
	if int64(1)<<(uint(v2585)%64)&int64(288080842570334209) != int64(0) {
		v2729 = v2583
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L760
	}
L758:
	;
	goto L759
L759:
	;
	v2594 = F_pg_mblen_range(m, v2583, v465)
	mBase = m.M
	v2595 = m.ExcPending
	if v2595 != 0 {
		goto L71
	} else {
		goto L761
	}
L760:
	;
	goto L759
L761:
	;
	v2729 = v2594 + v2583
	v2734 = v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L762:
	;
	if v356 == int32(45) {
		goto L765
	} else {
		goto L766
	}
L763:
	;
	goto L764
L764:
	;
	v2606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v2606 == int32(45) {
		goto L769
	} else {
		goto L770
	}
L765:
	;
	v2599 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v2599)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L766:
	;
	goto L767
L767:
	;
	v2601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v2601&int32(32) != 0 {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L768
	}
L768:
	;
	v2604 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v2604)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L769:
	;
	v2609 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2609)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L770:
	;
	goto L771
L771:
	;
	if base.Ui32(v465) <= base.Ui32(v479) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L772
	}
L772:
	;
	if base.Ui32(v2606) <= base.Ui32(int32(63)) {
		goto L773
	} else {
		goto L774
	}
L773:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v2606))%64)&int64(288080842570334209) != int64(0) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L776
	}
L774:
	;
	goto L775
L775:
	;
	v2621 = F_pg_mblen_range(m, v479, v465)
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L71
	} else {
		goto L777
	}
L776:
	;
	goto L775
L777:
	;
	v2729 = v2621 + v479
	v2734 = v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L778:
	;
	if v356 == int32(43) {
		goto L781
	} else {
		goto L782
	}
L779:
	;
	goto L780
L780:
	;
	v2633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v2633 == int32(43) {
		goto L785
	} else {
		goto L786
	}
L781:
	;
	v2626 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v2626)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L782:
	;
	goto L783
L783:
	;
	v2628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v2628&int32(32) != 0 {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L784
	}
L784:
	;
	v2631 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v2631)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L785:
	;
	v2636 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2636)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L786:
	;
	goto L787
L787:
	;
	if base.Ui32(v465) <= base.Ui32(v479) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L788
	}
L788:
	;
	if base.Ui32(v2633) <= base.Ui32(int32(63)) {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v2633))%64)&int64(288080842570334209) != int64(0) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L792
	}
L790:
	;
	goto L791
L791:
	;
	v2648 = F_pg_mblen_range(m, v479, v465)
	mBase = m.M
	v2649 = m.ExcPending
	if v2649 != 0 {
		goto L71
	} else {
		goto L793
	}
L792:
	;
	goto L791
L793:
	;
	v2729 = v2648 + v479
	v2734 = v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L794:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v356)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L795:
	;
	goto L796
L796:
	;
	v2652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	switch v2652 - int32(43) {
	case 0:
		goto L798
	default:
		goto L797
	case 2:
		goto L799
	}
L797:
	;
	if base.Ui32(v465) <= base.Ui32(v479) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L800
	}
L798:
	;
	v2657 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2657)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L799:
	;
	v2655 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2655)
	v2681 = v479
	v2686 = v484
	v2696 = v494
	v2697 = v495
	v2704 = v502
	goto L148
L800:
	;
	if base.Ui32(v2652) <= base.Ui32(int32(63)) {
		goto L801
	} else {
		goto L802
	}
L801:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v2652))%64)&int64(288080842570334209) != int64(0) {
		v2729 = v479
		v2734 = v484
		v2735 = v485
		v2737 = v487
		v2739 = v489
		v2744 = v494
		v2745 = v495
		v2752 = v502
		goto L117
	} else {
		goto L804
	}
L802:
	;
	goto L803
L803:
	;
	v2669 = F_pg_mblen_range(m, v479, v465)
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		goto L71
	} else {
		goto L805
	}
L804:
	;
	goto L803
L805:
	;
	v2729 = v2669 + v479
	v2734 = v484
	v2735 = v485
	v2737 = v487
	v2739 = v489
	v2744 = v494
	v2745 = v495
	v2752 = v502
	goto L117
L806:
	;
	goto L116
L807:
	;
	v2820 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2780))) = uint8(v2820)
	goto L6
L808:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2845
	goto L6
L809:
	;
	v2873 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2869))) = uint8(v2873)
	goto L808
L810:
	;
	goto L811
L811:
	;
	v2875 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2835))) = uint8(v2875)
	goto L808
L812:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v2979 = m.ExcPending
	if v2979 != 0 {
		goto L71
	} else {
		goto L813
	}
L813:
	;
	F_errmsg(m, int32(310214), int32(0))
	mBase = m.M
	v2983 = m.ExcPending
	if v2983 != 0 {
		goto L71
	} else {
		goto L814
	}
L814:
	;
	F_errfinish(m, int32(497568), int32(6113), int32(209681))
	mBase = m.M
	v2988 = m.ExcPending
	if v2988 != 0 {
		goto L71
	} else {
		goto L815
	}
L815:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L816:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2995 = m.ExcPending
	if v2995 != 0 {
		goto L71
	} else {
		goto L817
	}
L817:
	;
	F_errmsg(m, int32(64715), int32(0))
	mBase = m.M
	v2999 = m.ExcPending
	if v2999 != 0 {
		goto L71
	} else {
		goto L818
	}
L818:
	;
	F_errfinish(m, int32(497568), int32(5836), int32(209681))
	mBase = m.M
	v3004 = m.ExcPending
	if v3004 != 0 {
		goto L71
	} else {
		goto L819
	}
L819:
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
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v465 int32
	_ = v465
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
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
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v691 int32
	_ = v691
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v863 int32
	_ = v863
	var v873 int32
	_ = v873
	v4 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(1040)
	m.G0 = v24
	v26 = F_strlen(m, l1)
	mBase = m.M
	v32 = F__emscripten_memset_bulkmem(m, v24+int32(528), base.I32_extend8_s(v4), int32(512))
	mBase = m.M
	goto L1
L1:
	;
	v38 = F__emscripten_memset_bulkmem(m, v24+int32(16), base.I32_extend8_s(int32(0)), int32(512))
	mBase = m.M
	goto L2
L2:
	;
	if int32(256) < v26 {
		v873 = v4
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v24 + int32(1040)
	return v873
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = F_palloc(m, int32(4096))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(0)
	v50 = F_FindWord(m, l0, l1, int32(741336), l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v50 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v52 = F_pstrdup(m, l1)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v60 = v43
	goto L10
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v61 == int32(0) {
		v303 = v60
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v55 = v43 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v52
	v60 = v55
	goto L10
L12:
	;
	if v41 == int32(0) {
		v844 = v303
		goto L69
	} else {
		goto L70
	}
L13:
	;
	v67 = v61
	v68 = v60
	v74 = v4
	goto L14
L14:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v85&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v303 = v290
	goto L12
L16:
	;
	v202 = v68
	v203 = int32(0)
	goto L45
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if base.Ui32(int32(255)) < base.Ui32(v88) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v96 = v67
	goto L19
L19:
	;
	if v26 < v74 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v179 = v67 + int32(4)
	v186 = v74
	goto L16
L21:
	;
	goto L22
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if v93 == int32(0) {
		v303 = v68
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v96 = v93
	goto L19
L24:
	;
	v98 = v74
	goto L26
L25:
	;
	v98 = v26
	goto L26
L26:
	;
	v102 = v96
	v109 = v74
	goto L27
L27:
	;
	if v109 == v98 {
		v303 = v68
		goto L12
	} else {
		goto L29
	}
L28:
	;
	v303 = v68
	goto L12
L29:
	;
	v122 = v102 + int32(4)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v128 = v122 + int32(base.Ui32(v123)>>(uint(int32(1))%32))*int32(12)
	if base.Ui32(v128) <= base.Ui32(v122) {
		v303 = v68
		goto L12
	} else {
		goto L30
	}
L30:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v109))))
	v137 = v122
	v138 = v128
	goto L31
L31:
	;
	v154 = int32(12)
	v155 = base.I32_div_s(v138-v137, v154)
	v160 = v137 + int32(base.Ui32(v155)>>(uint(int32(1))%32))*v154
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v163 = v161 & int32(255)
	if v131 == v163 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L28
L33:
	;
	v166 = v109 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v161) {
		v179 = v160
		v186 = v166
		goto L16
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v172 = base.B2i32(base.Ui32(v163) < base.Ui32(v131))
	if base.Ui32(v163) < base.Ui32(v131) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v169 != 0 {
		v102 = v169
		v109 = v166
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v303 = v68
	goto L12
L38:
	;
	v173 = v160 + int32(12)
	goto L40
L39:
	;
	v173 = v137
	goto L40
L40:
	;
	if base.Ui32(v163) < base.Ui32(v131) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v174 = v138
	goto L43
L42:
	;
	v174 = v160
	goto L43
L43:
	;
	if base.Ui32(v173) < base.Ui32(v174) {
		v137 = v173
		v138 = v174
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L32
L45:
	;
	v220 = v203 << (uint(int32(2)) % 32)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220+v221)))
	v227 = F_CheckAffix(m, l1, v26, v223, l2, v24+int32(528), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L48
	}
L46:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	if v298 != 0 {
		v67 = v298
		v68 = v290
		v74 = v186
		goto L14
	} else {
		goto L68
	}
L47:
	;
	v293 = v203 + int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if base.Ui32(v293) < base.Ui32(int32(base.Ui32(v294)>>(uint(int32(8))%32))) {
		v202 = v290
		v203 = v293
		goto L45
	} else {
		goto L67
	}
L48:
	;
	if v227 == int32(0) {
		v290 = v202
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233+v220)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v237 = F_FindWord(m, l0, v24+int32(528), v236, l2)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	if v237 == int32(0) {
		v290 = v202
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v241 = int32(0)
	if int32(4088) < v202-v43 {
		v286 = v241
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v290 = v202 + v286<<(uint(int32(2))%32)
	goto L47
L53:
	;
	if v202 != v43 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v247 = v24 + int32(528)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v202-int32(4))))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	if v254 == int32(0) {
		v273 = v253
		v274 = v254
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	v280 = F_pstrdup(m, v24+int32(528))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L66
	}
L57:
	;
	if v274-v273 == int32(0) {
		v286 = v241
		goto L52
	} else {
		goto L65
	}
L58:
	;
	goto L57
L59:
	;
	if v253 != v254 {
		v273 = v253
		v274 = v254
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v258 = v247
	v259 = v250
	goto L61
L61:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+1)))
	if v263 == int32(0) {
		v273 = v262
		v274 = v263
		goto L58
	} else {
		goto L63
	}
L62:
	;
	v273 = v262
	v274 = v263
	goto L58
L63:
	;
	v266 = int32(1)
	if v262 == v263 {
		v258 = v258 + v266
		v259 = v259 + v266
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L56
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v280
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = int32(0)
	v286 = int32(1)
	goto L52
L67:
	;
	goto L46
L68:
	;
	goto L15
L69:
	;
	if v844 != v43 {
		v873 = v43
		goto L3
	} else {
		goto L187
	}
L70:
	;
	v327 = v303
	v332 = v41
	v338 = v4
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v346&int32(1) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v844 = v816
	goto L69
L73:
	;
	v465 = v327
	v479 = int32(0)
	goto L102
L74:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	if base.Ui32(int32(255)) < base.Ui32(v349) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v357 = v332
	goto L76
L76:
	;
	if v26 < v338 {
		goto L81
	} else {
		goto L82
	}
L77:
	;
	v451 = v332 + int32(4)
	v454 = v338
	goto L73
L78:
	;
	goto L79
L79:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	if v354 == int32(0) {
		v844 = v327
		goto L69
	} else {
		goto L80
	}
L80:
	;
	v357 = v354
	goto L76
L81:
	;
	v359 = v338
	goto L83
L82:
	;
	v359 = v26
	goto L83
L83:
	;
	v369 = v357
	v375 = v338
	goto L84
L84:
	;
	if v359 == v375 {
		v844 = v327
		goto L69
	} else {
		goto L86
	}
L85:
	;
	v844 = v327
	goto L69
L86:
	;
	v383 = v369 + int32(4)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v389 = v383 + int32(base.Ui32(v384)>>(uint(int32(1))%32))*int32(12)
	if base.Ui32(v389) <= base.Ui32(v383) {
		v844 = v327
		goto L69
	} else {
		goto L87
	}
L87:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v26+(v375^int32(-1))))))
	v398 = v383
	v400 = v389
	goto L88
L88:
	;
	v417 = int32(12)
	v418 = base.I32_div_s(v400-v398, v417)
	v423 = v398 + int32(base.Ui32(v418)>>(uint(int32(1))%32))*v417
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	v426 = v424 & int32(255)
	if v394 == v426 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L85
L90:
	;
	v429 = v375 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v424) {
		v451 = v423
		v454 = v429
		goto L73
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v435 = base.B2i32(base.Ui32(v426) < base.Ui32(v394))
	if base.Ui32(v426) < base.Ui32(v394) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v423)+8))
	if v432 != 0 {
		v369 = v432
		v375 = v429
		goto L84
	} else {
		goto L94
	}
L94:
	;
	v844 = v327
	goto L69
L95:
	;
	v436 = v423 + int32(12)
	goto L97
L96:
	;
	v436 = v398
	goto L97
L97:
	;
	if base.Ui32(v426) < base.Ui32(v394) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v437 = v400
	goto L100
L99:
	;
	v437 = v423
	goto L100
L100:
	;
	if base.Ui32(v436) < base.Ui32(v437) {
		v398 = v436
		v400 = v437
		goto L88
	} else {
		goto L101
	}
L101:
	;
	goto L89
L102:
	;
	v483 = v479 << (uint(int32(2)) % 32)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v483+v484)))
	v491 = F_CheckAffix(m, l1, v26, v486, l2, v24+int32(528), v24+int32(12))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L5
	} else {
		goto L105
	}
L103:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v451)+8))
	if v839 != 0 {
		v327 = v816
		v332 = v839
		v338 = v454
		goto L71
	} else {
		goto L186
	}
L104:
	;
	v834 = v479 + int32(1)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	if base.Ui32(v834) < base.Ui32(int32(base.Ui32(v835)>>(uint(int32(8))%32))) {
		v465 = v816
		v479 = v834
		goto L102
	} else {
		goto L185
	}
L105:
	;
	if v491 == int32(0) {
		v816 = v465
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v497+v483)))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v501 = F_FindWord(m, l0, v24+int32(528), v500, l2)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	if v501 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v503 = int32(0)
	if int32(4088) < v465-v43 {
		v548 = v503
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v553 = v465
	goto L110
L110:
	;
	v556 = F_strlen(m, v24+int32(528))
	mBase = m.M
	v557 = int32(0)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v558 == v557 {
		v816 = v553
		goto L104
	} else {
		goto L126
	}
L111:
	;
	v553 = v465 + v548<<(uint(int32(2))%32)
	goto L110
L112:
	;
	if v465 != v43 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v509 = v24 + int32(528)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v465-int32(4))))
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
	if v516 == int32(0) {
		v535 = v515
		v536 = v516
		goto L117
	} else {
		goto L118
	}
L114:
	;
	goto L115
L115:
	;
	v542 = F_pstrdup(m, v24+int32(528))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L5
	} else {
		goto L125
	}
L116:
	;
	if v536-v535 == int32(0) {
		v548 = v503
		goto L111
	} else {
		goto L124
	}
L117:
	;
	goto L116
L118:
	;
	if v515 != v516 {
		v535 = v515
		v536 = v516
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v520 = v509
	v521 = v512
	goto L120
L120:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+1)))
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+1)))
	if v525 == int32(0) {
		v535 = v524
		v536 = v525
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v535 = v524
	v536 = v525
	goto L117
L122:
	;
	v528 = int32(1)
	if v524 == v525 {
		v520 = v520 + v528
		v521 = v521 + v528
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	goto L115
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v465))) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v465)+4)) = int32(0)
	v548 = int32(1)
	goto L111
L126:
	;
	v564 = v558
	v565 = v553
	v577 = v557
	goto L127
L127:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	if v582&int32(1) != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v816 = v802
	goto L104
L129:
	;
	v701 = v565
	v702 = int32(0)
	goto L158
L130:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v564)+4))
	if base.Ui32(int32(255)) < base.Ui32(v585) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v593 = v564
	goto L132
L132:
	;
	if v556 < v577 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	v678 = v564 + int32(4)
	v691 = v577
	goto L129
L134:
	;
	goto L135
L135:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v564)+12))
	if v590 == int32(0) {
		v816 = v565
		goto L104
	} else {
		goto L136
	}
L136:
	;
	v593 = v590
	goto L132
L137:
	;
	v595 = v577
	goto L139
L138:
	;
	v595 = v556
	goto L139
L139:
	;
	v599 = v593
	v612 = v577
	goto L140
L140:
	;
	if v612 == v595 {
		v816 = v565
		goto L104
	} else {
		goto L142
	}
L141:
	;
	v816 = v565
	goto L104
L142:
	;
	v619 = v599 + int32(4)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	v625 = v619 + int32(base.Ui32(v620)>>(uint(int32(1))%32))*int32(12)
	if base.Ui32(v625) <= base.Ui32(v619) {
		v816 = v565
		goto L104
	} else {
		goto L143
	}
L143:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(528)+v612))))
	v636 = v619
	v637 = v625
	goto L144
L144:
	;
	v653 = int32(12)
	v654 = base.I32_div_s(v637-v636, v653)
	v659 = v636 + int32(base.Ui32(v654)>>(uint(int32(1))%32))*v653
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	v662 = v660 & int32(255)
	if v630 == v662 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L141
L146:
	;
	v665 = v612 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v660) {
		v678 = v659
		v691 = v665
		goto L129
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v671 = base.B2i32(base.Ui32(v662) < base.Ui32(v630))
	if base.Ui32(v662) < base.Ui32(v630) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v659)+8))
	if v668 != 0 {
		v599 = v668
		v612 = v665
		goto L140
	} else {
		goto L150
	}
L150:
	;
	v816 = v565
	goto L104
L151:
	;
	v672 = v659 + int32(12)
	goto L153
L152:
	;
	v672 = v636
	goto L153
L153:
	;
	if base.Ui32(v662) < base.Ui32(v630) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v673 = v637
	goto L156
L155:
	;
	v673 = v659
	goto L156
L156:
	;
	if base.Ui32(v672) < base.Ui32(v673) {
		v636 = v672
		v637 = v673
		goto L144
	} else {
		goto L157
	}
L157:
	;
	goto L145
L158:
	;
	v721 = v702 << (uint(int32(2)) % 32)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v721+v722)))
	v729 = F_CheckAffix(m, v24+int32(528), v556, v724, l2, v24+int32(16), v24+int32(12))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L5
	} else {
		goto L161
	}
L159:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v678)+8))
	if v811 != 0 {
		v564 = v811
		v565 = v802
		v577 = v691
		goto L127
	} else {
		goto L184
	}
L160:
	;
	v806 = v702 + int32(1)
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	if base.Ui32(v806) < base.Ui32(int32(base.Ui32(v807)>>(uint(int32(8))%32))) {
		v701 = v802
		v702 = v806
		goto L158
	} else {
		goto L183
	}
L161:
	;
	if v729 == int32(0) {
		v802 = v701
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v736+v721)))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v738)+4))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v740+v483)))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)+4))
	if v739&v743&int32(128) != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v748 = int32(741336)
	goto L165
L164:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v738)))
	v748 = v747
	goto L165
L165:
	;
	v749 = F_FindWord(m, l0, v24+int32(16), v748, l2)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L5
	} else {
		goto L166
	}
L166:
	;
	if v749 == int32(0) {
		v802 = v701
		goto L160
	} else {
		goto L167
	}
L167:
	;
	v753 = int32(0)
	if int32(4088) < v701-v43 {
		v798 = v753
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v802 = v701 + v798<<(uint(int32(2))%32)
	goto L160
L169:
	;
	if v701 != v43 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v759 = v24 + int32(16)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v701-int32(4))))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759))))
	if v766 == int32(0) {
		v785 = v765
		v786 = v766
		goto L174
	} else {
		goto L175
	}
L171:
	;
	goto L172
L172:
	;
	v792 = F_pstrdup(m, v24+int32(16))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L182
	}
L173:
	;
	if v786-v785 == int32(0) {
		v798 = v753
		goto L168
	} else {
		goto L181
	}
L174:
	;
	goto L173
L175:
	;
	if v765 != v766 {
		v785 = v765
		v786 = v766
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v770 = v759
	v771 = v762
	goto L177
L177:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771)+1)))
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770)+1)))
	if v775 == int32(0) {
		v785 = v774
		v786 = v775
		goto L174
	} else {
		goto L179
	}
L178:
	;
	v785 = v774
	v786 = v775
	goto L174
L179:
	;
	v778 = int32(1)
	if v774 == v775 {
		v770 = v770 + v778
		v771 = v771 + v778
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	goto L172
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v701))) = v792
	*(*int32)(unsafe.Add(mBase, uint32(v701)+4)) = int32(0)
	v798 = int32(1)
	goto L168
L183:
	;
	goto L159
L184:
	;
	goto L128
L185:
	;
	goto L103
L186:
	;
	goto L72
L187:
	;
	F_pfree(m, v43)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	v873 = int32(0)
	goto L3
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
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
	v16 = F_pg_snprintf(m, v7+int32(16), int32(20), int32(38544), v7)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = F_strlen(m, v9)
		mBase = m.M
		if int32(64) <= v16+v20 {
			v26 = F_pg_mbcliplen(m, v9, v20, int32(63)-v16)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = v26
				v30 = F_palloc0(m, int32(64))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if v28 != 0 {
						v32 = F__emscripten_memcpy_bulkmem(m, v30, v9, v28)
						mBase = m.M
						v33 = v32
					} else {
						v33 = v30
					}
					if v16 != 0 {
						v37 = F__emscripten_memcpy_bulkmem(m, v33+v28, v7+int32(16), v16)
						mBase = m.M
					} else {
					}
					m.G0 = v7 + int32(48)
					return v33
				}
			}
		} else {
			v28 = v20
			v30 = F_palloc0(m, int32(64))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v28 != 0 {
					v32 = F__emscripten_memcpy_bulkmem(m, v30, v9, v28)
					mBase = m.M
					v33 = v32
				} else {
					v33 = v30
				}
				if v16 != 0 {
					v37 = F__emscripten_memcpy_bulkmem(m, v33+v28, v7+int32(16), v16)
					mBase = m.M
				} else {
				}
				m.G0 = v7 + int32(48)
				return v33
			}
		}
	}
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1558), v3, v4, v5)
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
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
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
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	v2 = F_strlen(m, l0)
	mBase = m.M
	v8 = v2 - int32(1636608432)
	if l0&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v2) {
			v117 = l0
			v118 = v2
			v119 = v8
			v120 = v8
			v121 = v8
			for {
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
				v124 = v123 + v120
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
				v128 = v127 + v121
				v130 = int32(4)
				v132 = v125 + v119 - v128 ^ base.I32_rotl(v128, v130)
				v136 = v124 - v132 ^ base.I32_rotl(v132, int32(6))
				v137 = v128 + v124
				v138 = v132 + v137
				v139 = v136 + v138
				v143 = v137 - v136 ^ base.I32_rotl(v136, int32(8))
				v147 = v138 - v143 ^ base.I32_rotl(v143, int32(16))
				v151 = v139 - v147 ^ base.I32_rotl(v147, int32(19))
				v152 = v143 + v139
				v153 = v147 + v152
				v154 = v151 + v153
				v158 = v152 - v151 ^ base.I32_rotl(v151, v130)
				v159 = int32(12)
				v160 = v117 + v159
				v162 = v118 - v159
				if base.Ui32(int32(11)) < base.Ui32(v162) {
					v117 = v160
					v118 = v162
					v119 = v153
					v120 = v154
					v121 = v158
					continue
				} else {
					break
				}
				break
			}
			v165 = v160
			v166 = v162
			v167 = v153
			v168 = v154
			v169 = v158
		} else {
			v165 = l0
			v166 = v2
			v167 = v8
			v168 = v8
			v169 = v8
		}
		switch v166 - int32(1) {
		case 0:
			v228 = v167
			v229 = v168
			v230 = v169
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 1:
			v221 = v167
			v222 = v168
			v223 = v169
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 2:
			v214 = v167
			v215 = v168
			v216 = v169
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 3:
			v208 = v168
			v209 = v169
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 4:
			v204 = v168
			v205 = v169
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 5:
			v198 = v168
			v199 = v169
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 6:
			v192 = v168
			v193 = v169
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 7:
			v187 = v169
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 8:
			v182 = v169
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 9:
			v177 = v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 10:
			v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+10)))
			v177 = v173<<(uint(int32(24))%32) + v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		default:
			v235 = v167
			v236 = v168
			v237 = v169
		}
	} else {
		if base.Ui32(v2) < base.Ui32(int32(12)) {
			v63 = l0
			v64 = v2
			v65 = v8
			v66 = v8
			v67 = v8
		} else {
			v15 = l0
			v16 = v2
			v17 = v8
			v18 = v8
			v19 = v8
			for {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v22 = v21 + v18
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v26 = v25 + v19
				v28 = int32(4)
				v30 = v23 + v17 - v26 ^ base.I32_rotl(v26, v28)
				v34 = v22 - v30 ^ base.I32_rotl(v30, int32(6))
				v35 = v26 + v22
				v36 = v30 + v35
				v37 = v34 + v36
				v41 = v35 - v34 ^ base.I32_rotl(v34, int32(8))
				v45 = v36 - v41 ^ base.I32_rotl(v41, int32(16))
				v49 = v37 - v45 ^ base.I32_rotl(v45, int32(19))
				v50 = v41 + v37
				v51 = v45 + v50
				v52 = v49 + v51
				v56 = v50 - v49 ^ base.I32_rotl(v49, v28)
				v57 = int32(12)
				v58 = v15 + v57
				v60 = v16 - v57
				if base.Ui32(int32(11)) < base.Ui32(v60) {
					v15 = v58
					v16 = v60
					v17 = v51
					v18 = v52
					v19 = v56
					continue
				} else {
					break
				}
				break
			}
			v63 = v58
			v64 = v60
			v65 = v51
			v66 = v52
			v67 = v56
		}
		switch v64 - int32(1) {
		case 0:
			v114 = v65
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 1:
			v109 = v65
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 2:
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
			v109 = v105<<(uint(int32(16))%32) + v65
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 3:
			v102 = v66
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 4:
			v99 = v66
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 5:
			v94 = v66
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
			v99 = v95<<(uint(int32(8))%32) + v94
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 6:
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
			v94 = v90<<(uint(int32(16))%32) + v66
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
			v99 = v95<<(uint(int32(8))%32) + v94
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 7:
			v85 = v67
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 8:
			v80 = v67
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 9:
			v75 = v67
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+9)))
			v80 = v76<<(uint(int32(16))%32) + v75
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 10:
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+10)))
			v75 = v71<<(uint(int32(24))%32) + v67
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+9)))
			v80 = v76<<(uint(int32(16))%32) + v75
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		default:
			v235 = v65
			v236 = v66
			v237 = v67
		}
	}
	v240 = int32(14)
	v242 = v236 ^ v237 - base.I32_rotl(v236, v240)
	v246 = v242 ^ v235 - base.I32_rotl(v242, int32(11))
	v250 = v246 ^ v236 - base.I32_rotl(v246, int32(25))
	v254 = v250 ^ v242 - base.I32_rotl(v250, int32(16))
	v258 = v254 ^ v246 - base.I32_rotl(v254, int32(4))
	v262 = v258 ^ v250 - base.I32_rotl(v258, v240)
	return v262 ^ v254 - base.I32_rotl(v262, int32(24))
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
	v61 = *(*int32)(unsafe.Add(mBase, _consts[661]))
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
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+uint32(_consts[662]))))
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
	v281 = *(*int32)(unsafe.Add(mBase, _consts[661]))
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
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+uint32(_consts[662]))))
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
	v443 = *(*int32)(unsafe.Add(mBase, _consts[661]))
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
	v456 = *(*int32)(unsafe.Add(mBase, _consts[663]))
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
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+uint32(_consts[662]))))
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
	v880 = *(*int32)(unsafe.Add(mBase, _consts[661]))
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
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878)+uint32(_consts[662]))))
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
	v947 = *(*int32)(unsafe.Add(mBase, _consts[661]))
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
	v955 = *(*int32)(unsafe.Add(mBase, _consts[663]))
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
	v1242 = *(*int32)(unsafe.Add(mBase, _consts[661]))
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
	v1255 = *(*int32)(unsafe.Add(mBase, _consts[663]))
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
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+uint32(_consts[662]))))
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
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v467 int32
	_ = v467
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v546 int32
	_ = v546
	var v555 int32
	_ = v555
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v820 int32
	_ = v820
	var v856 int32
	_ = v856
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
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
	if v856 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L2:
	;
	v856 = int32(0)
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
	v804 = v20 + int32(4096) + v792
	v806 = v804
	goto L234
L22:
	;
	v792 = v71
	v793 = v787
	v794 = int32(0)
	v796 = int32(0)
	goto L21
L23:
	;
	v69 = int32(1)
	v71 = v51 + v69
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+(v20+int32(4096))))))
	v76 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v76)
	if v75 != v76 {
		v787 = v69
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
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v20)+uint32(_consts[1229]))))
	if v84 == int32(47) {
		v787 = v69
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v87 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v87)
	v787 = int32(2)
	goto L22
L28:
	;
	v495 = v494 + v54
	if base.Ui32(int32(4095)) < base.Ui32(v495) {
		goto L13
	} else {
		goto L139
	}
L29:
	;
	v493 = v51
	v494 = v187
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
	v792 = v51 + int32(1)
	v793 = v54
	v794 = v55
	v796 = v57
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
	v493 = v207
	v494 = v187 + v206
	goto L28
L63:
	;
	v225 = F_getcwd(m, v20+int32(4096), int32(4097))
	mBase = m.M
	if v225 == int32(0) {
		v856 = v215
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v484 = F_strlen(m, v20)
	mBase = m.M
	v486 = v484 + int32(1)
	v487 = F_emscripten_builtin_malloc(m, v486)
	mBase = m.M
	if v487 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L66:
	;
	v228 = int32(0)
	v231 = F_strlen(m, v20+int32(4096))
	mBase = m.M
	if v57 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v233 = v231
	v237 = v228
	v239 = v57
	goto L70
L68:
	;
	v284 = v231
	v288 = v228
	goto L69
L69:
	;
	v296 = v54 - v288
	if v54 == v288 {
		v311 = v284
		goto L81
	} else {
		goto L82
	}
L70:
	;
	v248 = v239 - int32(1)
	v250 = v233
	goto L73
L71:
	;
	v284 = v282
	v288 = v279
	goto L69
L72:
	;
	v277 = v237 + int32(2)
	if base.Ui32(v277) < base.Ui32(v54) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	if base.Ui32(v250) < base.Ui32(int32(2)) {
		v273 = base.B2i32(v233 != int32(0))
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v273 = v250
	goto L72
L75:
	;
	v265 = v250 - int32(1)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+(v20+int32(4096))))))
	if v269 != int32(47) {
		v250 = v265
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v279 = v237 + int32(3)
	goto L79
L78:
	;
	v279 = v277
	goto L79
L79:
	;
	v282 = v273 - base.B2i32(base.Ui32(int32(1)) < base.Ui32(v250))
	if v248 != 0 {
		v233 = v282
		v237 = v279
		v239 = v248
		goto L70
	} else {
		goto L80
	}
L80:
	;
	goto L71
L81:
	;
	if base.Ui32(v311+v296-int32(4095)) < base.Ui32(int32(-4096)) {
		goto L13
	} else {
		goto L84
	}
L82:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v20+int32(4095)))))
	if v301 == int32(47) {
		v311 = v284
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v307 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(4096)+v284))) = uint8(v307)
	v311 = v284 + int32(1)
	goto L81
L84:
	;
	v317 = v311 + v20
	v318 = v20 + v288
	v320 = v296 + int32(1)
	if v317 == v318 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v311 != 0 {
		goto L132
	} else {
		goto L133
	}
L86:
	;
	goto L85
L87:
	;
	v324 = v317 + v320
	if base.Ui32(v318-v324) <= base.Ui32(int32(0)-v320<<(uint(int32(1))%32)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v331 = F___memcpy(m, v317, v318, v320)
	mBase = m.M
	goto L85
L89:
	;
	goto L90
L90:
	;
	v334 = (v317 ^ v318) & int32(3)
	if base.Ui32(v317) < base.Ui32(v318) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	if v436 == int32(0) {
		goto L86
	} else {
		goto L127
	}
L92:
	;
	if base.Ui32(v414) <= base.Ui32(int32(3)) {
		v435 = v413
		v436 = v414
		v437 = v415
		goto L91
	} else {
		goto L123
	}
L93:
	;
	if v334 != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if v334 != 0 {
		v396 = v320
		goto L106
	} else {
		goto L107
	}
L96:
	;
	v435 = v318
	v436 = v320
	v437 = v317
	goto L91
L97:
	;
	goto L98
L98:
	;
	if v317&int32(3) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v413 = v318
	v414 = v320
	v415 = v317
	goto L92
L100:
	;
	goto L101
L101:
	;
	v341 = v318
	v342 = v320
	v343 = v317
	goto L102
L102:
	;
	if v342 == int32(0) {
		goto L86
	} else {
		goto L104
	}
L103:
	;
	v413 = v350
	v414 = v352
	v415 = v354
	goto L92
L104:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	*(*uint8)(unsafe.Add(mBase, uint32(v343))) = uint8(v347)
	v349 = int32(1)
	v350 = v341 + v349
	v352 = v342 - v349
	v354 = v343 + v349
	if v354&int32(3) != 0 {
		v341 = v350
		v342 = v352
		v343 = v354
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	if v396 == int32(0) {
		goto L86
	} else {
		goto L119
	}
L107:
	;
	if v324&int32(3) != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v361 = v320
	goto L111
L109:
	;
	v376 = v320
	goto L110
L110:
	;
	if base.Ui32(v376) <= base.Ui32(int32(3)) {
		v396 = v376
		goto L106
	} else {
		goto L115
	}
L111:
	;
	if v361 == int32(0) {
		goto L86
	} else {
		goto L113
	}
L112:
	;
	v376 = v367
	goto L110
L113:
	;
	v367 = v361 - int32(1)
	v368 = v317 + v367
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318+v367))))
	*(*uint8)(unsafe.Add(mBase, uint32(v368))) = uint8(v370)
	if v368&int32(3) != 0 {
		v361 = v367
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v383 = v376
	goto L116
L116:
	;
	v387 = v383 - int32(4)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v318+v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v317+v387))) = v390
	if base.Ui32(int32(3)) < base.Ui32(v387) {
		v383 = v387
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v396 = v387
	goto L106
L118:
	;
	goto L117
L119:
	;
	v403 = v396
	goto L120
L120:
	;
	v407 = v403 - int32(1)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318+v407))))
	*(*uint8)(unsafe.Add(mBase, uint32(v317+v407))) = uint8(v410)
	if v407 != 0 {
		v403 = v407
		goto L120
	} else {
		goto L122
	}
L121:
	;
	goto L86
L122:
	;
	goto L121
L123:
	;
	v420 = v413
	v421 = v414
	v422 = v415
	goto L124
L124:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = v424
	v426 = int32(4)
	v427 = v420 + v426
	v429 = v422 + v426
	v431 = v421 - v426
	if base.Ui32(int32(3)) < base.Ui32(v431) {
		v420 = v427
		v421 = v431
		v422 = v429
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v435 = v427
	v436 = v431
	v437 = v429
	goto L91
L126:
	;
	goto L125
L127:
	;
	v442 = v435
	v443 = v436
	v444 = v437
	goto L128
L128:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	*(*uint8)(unsafe.Add(mBase, uint32(v444))) = uint8(v446)
	v448 = int32(1)
	v453 = v443 - v448
	if v453 != 0 {
		v442 = v442 + v448
		v443 = v453
		v444 = v444 + v448
		goto L128
	} else {
		goto L130
	}
L129:
	;
	goto L86
L130:
	;
	goto L129
L131:
	;
	goto L65
L132:
	;
	v467 = F__emscripten_memcpy_bulkmem(m, v20, v20+int32(4096), v311)
	mBase = m.M
	goto L134
L133:
	;
	goto L134
L134:
	;
	goto L131
L135:
	;
	v856 = v492
	goto L1
L136:
	;
	v492 = int32(0)
	goto L135
L137:
	;
	goto L138
L138:
	;
	v491 = F___memcpy(m, v487, v20, v486)
	mBase = m.M
	v492 = v491
	goto L135
L139:
	;
	if v494 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v505 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v495))) = uint8(v505)
	v507 = v493 + v494
	v508 = int32(1)
	if v187 != int32(2) {
		v532 = v508
		goto L148
	} else {
		goto L149
	}
L141:
	;
	v502 = F__emscripten_memcpy_bulkmem(m, v20+v54, v20+int32(4096)+v493, v494)
	mBase = m.M
	goto L143
L142:
	;
	goto L143
L143:
	;
	goto L140
L144:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v780 == int32(47) {
		goto L228
	} else {
		goto L229
	}
L145:
	;
	v583 = v61 + int32(1)
	if v583 == int32(40) {
		goto L173
	} else {
		goto L174
	}
L146:
	;
	if v187 != 0 {
		goto L170
	} else {
		goto L171
	}
L147:
	;
	v555 = v54
	goto L163
L148:
	;
	v536 = F_readlink(m, v20, v20+int32(4096), v507)
	mBase = m.M
	if v536 == v507 {
		goto L13
	} else {
		goto L156
	}
L149:
	;
	v513 = v507 + (v20 + int32(4096))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513-int32(2)))))
	if v516 != int32(46) {
		v532 = v508
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513-int32(1)))))
	if v521 != int32(46) {
		v532 = v508
		goto L148
	} else {
		goto L151
	}
L151:
	;
	if base.Ui32(v54) <= base.Ui32(v57*int32(3)) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v792 = v507
	v793 = v495
	v794 = v55
	v796 = v57 + int32(1)
	goto L21
L153:
	;
	goto L154
L154:
	;
	v529 = int32(0)
	if v55 == v529 {
		goto L147
	} else {
		goto L155
	}
L155:
	;
	v532 = v529
	goto L148
L156:
	;
	if v536 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(44)
	goto L2
L158:
	;
	goto L159
L159:
	;
	if int32(0) <= v536 {
		goto L145
	} else {
		goto L160
	}
L160:
	;
	v546 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v546 != int32(28) {
		goto L2
	} else {
		goto L161
	}
L161:
	;
	if v532 != 0 {
		goto L146
	} else {
		goto L162
	}
L162:
	;
	goto L147
L163:
	;
	v564 = int32(0)
	if v555 == v564 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v555 != int32(1) {
		goto L144
	} else {
		goto L169
	}
L165:
	;
	v792 = v507
	v793 = int32(0)
	v794 = v564
	v796 = v57
	goto L21
L166:
	;
	goto L167
L167:
	;
	v569 = v555 - int32(1)
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v569))))
	if v571 != int32(47) {
		v555 = v569
		goto L163
	} else {
		goto L168
	}
L168:
	;
	goto L164
L169:
	;
	v792 = v507
	v793 = int32(1)
	v794 = v564
	v796 = v57
	goto L21
L170:
	;
	v577 = v495
	goto L172
L171:
	;
	v577 = v54
	goto L172
L172:
	;
	v581 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20+int32(4096)+v507))))
	v792 = v507
	v793 = v577
	v794 = v581
	v796 = v57
	goto L21
L173:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(32)
	goto L2
L174:
	;
	goto L175
L175:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v536+int32(4095)))))
	if v592 == int32(47) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v598 = v507
	goto L179
L177:
	;
	v617 = v507
	goto L178
L178:
	;
	v629 = v617 - v536
	v631 = v20 + int32(4096)
	v632 = v629 + v631
	if v632 == v631 {
		goto L183
	} else {
		goto L184
	}
L179:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598+(v20+int32(4096))))))
	if v613 == int32(47) {
		v598 = v598 + int32(1)
		goto L179
	} else {
		goto L181
	}
L180:
	;
	v617 = v598
	goto L178
L181:
	;
	goto L180
L182:
	;
	v51 = v629
	v61 = v583
	goto L19
L183:
	;
	goto L182
L184:
	;
	v638 = v632 + v536
	if base.Ui32(v631-v638) <= base.Ui32(int32(0)-v536<<(uint(int32(1))%32)) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v645 = F___memcpy(m, v632, v631, v536)
	mBase = m.M
	goto L182
L186:
	;
	goto L187
L187:
	;
	v648 = (v632 ^ v631) & int32(3)
	if base.Ui32(v632) < base.Ui32(v631) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	if v750 == int32(0) {
		goto L183
	} else {
		goto L224
	}
L189:
	;
	if base.Ui32(v728) <= base.Ui32(int32(3)) {
		v749 = v727
		v750 = v728
		v751 = v729
		goto L188
	} else {
		goto L220
	}
L190:
	;
	if v648 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	goto L192
L192:
	;
	if v648 != 0 {
		v710 = v536
		goto L203
	} else {
		goto L204
	}
L193:
	;
	v749 = v631
	v750 = v536
	v751 = v632
	goto L188
L194:
	;
	goto L195
L195:
	;
	if v632&int32(3) == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v727 = v631
	v728 = v536
	v729 = v632
	goto L189
L197:
	;
	goto L198
L198:
	;
	v655 = v631
	v656 = v536
	v657 = v632
	goto L199
L199:
	;
	if v656 == int32(0) {
		goto L183
	} else {
		goto L201
	}
L200:
	;
	v727 = v664
	v728 = v666
	v729 = v668
	goto L189
L201:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655))))
	*(*uint8)(unsafe.Add(mBase, uint32(v657))) = uint8(v661)
	v663 = int32(1)
	v664 = v655 + v663
	v666 = v656 - v663
	v668 = v657 + v663
	if v668&int32(3) != 0 {
		v655 = v664
		v656 = v666
		v657 = v668
		goto L199
	} else {
		goto L202
	}
L202:
	;
	goto L200
L203:
	;
	if v710 == int32(0) {
		goto L183
	} else {
		goto L216
	}
L204:
	;
	if v638&int32(3) != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v675 = v536
	goto L208
L206:
	;
	v690 = v536
	goto L207
L207:
	;
	if base.Ui32(v690) <= base.Ui32(int32(3)) {
		v710 = v690
		goto L203
	} else {
		goto L212
	}
L208:
	;
	if v675 == int32(0) {
		goto L183
	} else {
		goto L210
	}
L209:
	;
	v690 = v681
	goto L207
L210:
	;
	v681 = v675 - int32(1)
	v682 = v632 + v681
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631+v681))))
	*(*uint8)(unsafe.Add(mBase, uint32(v682))) = uint8(v684)
	if v682&int32(3) != 0 {
		v675 = v681
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v697 = v690
	goto L213
L213:
	;
	v701 = v697 - int32(4)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v631+v701)))
	*(*int32)(unsafe.Add(mBase, uint32(v632+v701))) = v704
	if base.Ui32(int32(3)) < base.Ui32(v701) {
		v697 = v701
		goto L213
	} else {
		goto L215
	}
L214:
	;
	v710 = v701
	goto L203
L215:
	;
	goto L214
L216:
	;
	v717 = v710
	goto L217
L217:
	;
	v721 = v717 - int32(1)
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631+v721))))
	*(*uint8)(unsafe.Add(mBase, uint32(v632+v721))) = uint8(v724)
	if v721 != 0 {
		v717 = v721
		goto L217
	} else {
		goto L219
	}
L218:
	;
	goto L183
L219:
	;
	goto L218
L220:
	;
	v734 = v727
	v735 = v728
	v736 = v729
	goto L221
L221:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	*(*int32)(unsafe.Add(mBase, uint32(v736))) = v738
	v740 = int32(4)
	v741 = v734 + v740
	v743 = v736 + v740
	v745 = v735 - v740
	if base.Ui32(int32(3)) < base.Ui32(v745) {
		v734 = v741
		v735 = v745
		v736 = v743
		goto L221
	} else {
		goto L223
	}
L222:
	;
	v749 = v741
	v750 = v745
	v751 = v743
	goto L188
L223:
	;
	goto L222
L224:
	;
	v756 = v749
	v757 = v750
	v758 = v751
	goto L225
L225:
	;
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756))))
	*(*uint8)(unsafe.Add(mBase, uint32(v758))) = uint8(v760)
	v762 = int32(1)
	v767 = v757 - v762
	if v767 != 0 {
		v756 = v756 + v762
		v757 = v767
		v758 = v758 + v762
		goto L225
	} else {
		goto L227
	}
L226:
	;
	goto L183
L227:
	;
	goto L226
L228:
	;
	v783 = int32(2)
	goto L230
L229:
	;
	v783 = v569
	goto L230
L230:
	;
	if v555 != int32(2) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v786 = v569
	goto L233
L232:
	;
	v786 = v783
	goto L233
L233:
	;
	v792 = v507
	v793 = v786
	v794 = v564
	v796 = v57
	goto L21
L234:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806))))
	if v820 == int32(47) {
		v806 = v806 + int32(1)
		goto L234
	} else {
		goto L236
	}
L235:
	;
	v51 = v792 + (v806 - v804)
	v54 = v793
	v55 = v794
	v57 = v796
	goto L19
L236:
	;
	goto L235
L237:
	;
	m.G0 = v16 + int32(16)
	return v1009
L238:
	;
	v873 = int32(-1)
	v876 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	goto L250
L241:
	;
	return int32(0)
L242:
	;
	if v876 == int32(0) {
		v1009 = v873
		goto L237
	} else {
		goto L243
	}
L243:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L241
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg_internal(m, int32(293589), v16)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L241
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(499438), int32(253), int32(321259))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L241
	} else {
		goto L246
	}
L246:
	;
	v1009 = v873
	goto L237
L247:
	;
	F_emscripten_builtin_free(m, v856)
	mBase = m.M
	v1009 = v2
	goto L237
L248:
	;
	v1005 = F_strlen(m, v994)
	mBase = m.M
	goto L247
L250:
	;
	goto L251
L251:
	;
	v899 = int32(1023)
	if (l0^v856)&int32(3) != 0 {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	v998 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v995))) = uint8(v998)
	goto L248
L253:
	;
	v979 = v974
	v980 = v975
	v981 = v976
	goto L275
L254:
	;
	if v969 == int32(0) {
		v994 = v967
		v995 = v968
		goto L252
	} else {
		goto L274
	}
L255:
	;
	v967 = v856
	v968 = l0
	v969 = v899
	goto L254
L256:
	;
	goto L257
L257:
	;
	if v856&int32(3) == int32(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	if v936 == int32(0) {
		v994 = v933
		v995 = v934
		goto L252
	} else {
		goto L267
	}
L259:
	;
	v933 = v856
	v934 = l0
	v935 = v899
	v936 = int32(1)
	goto L258
L260:
	;
	goto L261
L261:
	;
	v912 = v856
	v913 = l0
	v914 = v899
	goto L262
L262:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912))))
	*(*uint8)(unsafe.Add(mBase, uint32(v913))) = uint8(v916)
	if v916 == int32(0) {
		v974 = v912
		v975 = v913
		v976 = v914
		goto L253
	} else {
		goto L264
	}
L263:
	;
	v933 = v927
	v934 = v921
	v935 = v923
	v936 = v925
	goto L258
L264:
	;
	v920 = int32(1)
	v921 = v913 + v920
	v923 = v914 - v920
	v924 = int32(0)
	v925 = base.B2i32(v923 != v924)
	v927 = v912 + v920
	if v927&int32(3) == v924 {
		v933 = v927
		v934 = v921
		v935 = v923
		v936 = v925
		goto L258
	} else {
		goto L265
	}
L265:
	;
	if v923 != 0 {
		v912 = v927
		v913 = v921
		v914 = v923
		goto L262
	} else {
		goto L266
	}
L266:
	;
	goto L263
L267:
	;
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933))))
	if v939 == int32(0) {
		v967 = v933
		v968 = v934
		v969 = v935
		goto L254
	} else {
		goto L268
	}
L268:
	;
	if base.Ui32(v935) < base.Ui32(int32(4)) {
		v967 = v933
		v968 = v934
		v969 = v935
		goto L254
	} else {
		goto L269
	}
L269:
	;
	v945 = v933
	v946 = v934
	v947 = v935
	goto L270
L270:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
	v953 = int32(-2139062144)
	if (int32(16843008)-v950|v950)&v953 != v953 {
		v974 = v945
		v975 = v946
		v976 = v947
		goto L253
	} else {
		goto L272
	}
L271:
	;
	v967 = v961
	v968 = v959
	v969 = v963
	goto L254
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v946))) = v950
	v958 = int32(4)
	v959 = v946 + v958
	v961 = v945 + v958
	v963 = v947 - v958
	if base.Ui32(int32(3)) < base.Ui32(v963) {
		v945 = v961
		v946 = v959
		v947 = v963
		goto L270
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	v974 = v967
	v975 = v968
	v976 = v969
	goto L253
L275:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979))))
	*(*uint8)(unsafe.Add(mBase, uint32(v980))) = uint8(v983)
	if v983 == int32(0) {
		v994 = v979
		v995 = v980
		goto L252
	} else {
		goto L277
	}
L276:
	;
	v994 = v990
	v995 = v988
	goto L252
L277:
	;
	v987 = int32(1)
	v988 = v980 + v987
	v990 = v979 + v987
	v992 = v981 - v987
	if v992 != 0 {
		v979 = v990
		v980 = v988
		v981 = v992
		goto L275
	} else {
		goto L278
	}
L278:
	;
	goto L276
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
					F_errmsg_internal(m, int32(483395), v8+int32(-48))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return float64(0)
					} else {
						F_errfinish(m, int32(494201), int32(1741), int32(306287))
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
						F_errmsg_internal(m, int32(483395), v10)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return float64(0)
						} else {
							F_errfinish(m, int32(494201), int32(1769), int32(306287))
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
							F_errmsg_internal(m, int32(483395), v10)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(494201), int32(1769), int32(306287))
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
								F_errmsg_internal(m, int32(483395), v10)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(494201), int32(1769), int32(306287))
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
