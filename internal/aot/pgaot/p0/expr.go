package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecPushExprSetupSteps(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v475 int64
	_ = v475
	var v477 int64
	_ = v477
	var v479 int64
	_ = v479
	var v481 int64
	_ = v481
	var v483 int64
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v631 int64
	_ = v631
	var v633 int64
	_ = v633
	var v635 int64
	_ = v635
	var v637 int64
	_ = v637
	var v639 int64
	_ = v639
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v752 int32
	_ = v752
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v787 int64
	_ = v787
	var v789 int64
	_ = v789
	var v791 int64
	_ = v791
	var v793 int64
	_ = v793
	var v795 int64
	_ = v795
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v11
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	if v21 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+2)))
	if v175 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L2:
	;
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(2)
	v30 = v9 + int32(8)
	v34 = m.G0
	v36 = v34 - int32(16)
	m.G0 = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v24)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	if v128 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	m.G0 = v36 + int32(16)
	goto L3
L5:
	;
	v128 = int32(1)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = v103
	v117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v102
	if v103 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L5
	} else {
		goto L30
	}
L7:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)) = uint8(v112)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = int64(0)
	goto L5
L8:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)))
	if base.B2i32(v102 == int32(0))|base.B2i32(v106 != int32(1)) != 0 {
		goto L7
	} else {
		goto L28
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(base.B2i32(v42 != int32(0)))
	v102 = v41
	v103 = v42
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v38 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	switch v48 - int32(2) {
	case 0:
		goto L15
	case 1:
		goto L14
	default:
		goto L13
	}
L13:
	;
	if base.Ui32(int32(2)) < base.Ui32(v48-int32(4)) {
		goto L7
	} else {
		goto L26
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+101)))
	if v72 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+102)))
	if v52 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v51 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+98)))
	if v55 != int32(1) {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v38)+88))
	if v58 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v61 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v61)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
	v102 = v63
	v103 = v58
	goto L8
L20:
	;
	v69 = F_ExecGetResultSlotOps(m, v51, v36+int32(15))
	mBase = m.M
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
	v102 = v70
	v103 = v69
	goto L8
L21:
	;
	if v71 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L22:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+97)))
	if v75 != int32(1) {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v38)+84))
	if v78 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v81)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	v102 = v83
	v103 = v78
	goto L8
L25:
	;
	v89 = F_ExecGetResultSlotOps(m, v71, v36+int32(15))
	mBase = m.M
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	v102 = v90
	v103 = v89
	goto L8
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+100)))
	if v97 != int32(1) {
		v102 = v96
		v103 = v95
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v100)
	v102 = v96
	v103 = v95
	goto L8
L28:
	;
	if v103 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	goto L7
L30:
	;
	v128 = int32(0)
	goto L4
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v134 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v156 + int32(1)
	v162 = v155 + v156*int32(40)
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v162)+32)) = v163
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v162)+24)) = v165
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v162)+16)) = v167
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v162)+8)) = v169
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v162))) = v171
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v153
	v155 = v153
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v140 = F_palloc(m, int32(640))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v142 != v134 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	return
L38:
	;
	v153 = v140
	goto L33
L39:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v155 = v144
	goto L32
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v134 << (uint(int32(1)) % 32)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v151 = F_repalloc(m, v148, v134*int32(80))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v153 = v151
	goto L33
L43:
	;
	v331 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v331 <= int32(0) {
		goto L84
	} else {
		goto L85
	}
L44:
	;
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v178)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v186 = v9 + int32(8)
	v190 = m.G0
	v192 = v190 - int32(16)
	m.G0 = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+15)) = uint8(v178)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v186)+24))
	if v197 != 0 {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	if v284 == int32(0) {
		goto L43
	} else {
		goto L73
	}
L46:
	;
	m.G0 = v192 + int32(16)
	goto L45
L47:
	;
	v284 = int32(1)
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186)+28)) = v259
	v273 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+20)) = uint8(v273)
	*(*int32)(unsafe.Add(mBase, uint32(v186)+24)) = v258
	if v259 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L47
	} else {
		goto L72
	}
L49:
	;
	v268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+20)) = uint8(v268)
	*(*int64)(unsafe.Add(mBase, uint32(v186)+24)) = int64(0)
	goto L47
L50:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+15)))
	if base.B2i32(v258 == int32(0))|base.B2i32(v262 != int32(1)) != 0 {
		goto L49
	} else {
		goto L70
	}
L51:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+15)) = uint8(base.B2i32(v198 != int32(0)))
	v258 = v197
	v259 = v198
	goto L50
L52:
	;
	goto L53
L53:
	;
	if v194 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	switch v204 - int32(2) {
	case 0:
		goto L57
	case 1:
		goto L56
	default:
		goto L55
	}
L55:
	;
	if base.Ui32(int32(2)) < base.Ui32(v204-int32(4)) {
		goto L49
	} else {
		goto L68
	}
L56:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v194)+36))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+101)))
	if v228 != int32(1) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v194)+40))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+102)))
	if v208 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v207 == int32(0) {
		goto L49
	} else {
		goto L62
	}
L59:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+98)))
	if v211 != int32(1) {
		goto L49
	} else {
		goto L60
	}
L60:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v194)+88))
	if v214 == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v217 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+15)) = uint8(v217)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v207)+56))
	v258 = v219
	v259 = v214
	goto L50
L62:
	;
	v225 = F_ExecGetResultSlotOps(m, v207, v192+int32(15))
	mBase = m.M
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v207)+56))
	v258 = v226
	v259 = v225
	goto L50
L63:
	;
	if v227 == int32(0) {
		goto L49
	} else {
		goto L67
	}
L64:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+97)))
	if v231 != int32(1) {
		goto L49
	} else {
		goto L65
	}
L65:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v194)+84))
	if v234 == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v237 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+15)) = uint8(v237)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v227)+56))
	v258 = v239
	v259 = v234
	goto L50
L67:
	;
	v245 = F_ExecGetResultSlotOps(m, v227, v192+int32(15))
	mBase = m.M
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v227)+56))
	v258 = v246
	v259 = v245
	goto L50
L68:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v194)+80))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v194)+76))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+100)))
	if v253 != int32(1) {
		v258 = v252
		v259 = v251
		goto L50
	} else {
		goto L69
	}
L69:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+15)) = uint8(v256)
	v258 = v252
	v259 = v251
	goto L50
L70:
	;
	if v259 != 0 {
		goto L48
	} else {
		goto L71
	}
L71:
	;
	goto L49
L72:
	;
	v284 = int32(0)
	goto L46
L73:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v290 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v312 + int32(1)
	v318 = v311 + v312*int32(40)
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v318)+32)) = v319
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v318)+24)) = v321
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v318)+16)) = v323
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v318)+8)) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v318))) = v327
	goto L43
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v309
	v311 = v309
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v296 = F_palloc(m, int32(640))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L37
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v298 != v290 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v309 = v296
	goto L75
L80:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v311 = v300
	goto L74
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v290 << (uint(int32(1)) % 32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v307 = F_repalloc(m, v304, v290*int32(80))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L37
	} else {
		goto L83
	}
L83:
	;
	v309 = v307
	goto L75
L84:
	;
	v487 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v487 <= int32(0) {
		goto L125
	} else {
		goto L126
	}
L85:
	;
	v334 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v334)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v342 = v9 + int32(8)
	v346 = m.G0
	v348 = v346 - int32(16)
	m.G0 = v348
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v348)+15)) = uint8(v334)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v342)+24))
	if v353 != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	if v440 == int32(0) {
		goto L84
	} else {
		goto L114
	}
L87:
	;
	m.G0 = v348 + int32(16)
	goto L86
L88:
	;
	v440 = int32(1)
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342)+28)) = v415
	v429 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+20)) = uint8(v429)
	*(*int32)(unsafe.Add(mBase, uint32(v342)+24)) = v414
	if v415 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L88
	} else {
		goto L113
	}
L90:
	;
	v424 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+20)) = uint8(v424)
	*(*int64)(unsafe.Add(mBase, uint32(v342)+24)) = int64(0)
	goto L88
L91:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+15)))
	if base.B2i32(v414 == int32(0))|base.B2i32(v418 != int32(1)) != 0 {
		goto L90
	} else {
		goto L111
	}
L92:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v342)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v348)+15)) = uint8(base.B2i32(v354 != int32(0)))
	v414 = v353
	v415 = v354
	goto L91
L93:
	;
	goto L94
L94:
	;
	if v350 == int32(0) {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	switch v360 - int32(2) {
	case 0:
		goto L98
	case 1:
		goto L97
	default:
		goto L96
	}
L96:
	;
	if base.Ui32(int32(2)) < base.Ui32(v360-int32(4)) {
		goto L90
	} else {
		goto L109
	}
L97:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v350)+36))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+101)))
	if v384 != int32(1) {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v350)+40))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+102)))
	if v364 != int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v363 == int32(0) {
		goto L90
	} else {
		goto L103
	}
L100:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+98)))
	if v367 != int32(1) {
		goto L90
	} else {
		goto L101
	}
L101:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v350)+88))
	if v370 == int32(0) {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v373 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v348)+15)) = uint8(v373)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v363)+56))
	v414 = v375
	v415 = v370
	goto L91
L103:
	;
	v381 = F_ExecGetResultSlotOps(m, v363, v348+int32(15))
	mBase = m.M
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v363)+56))
	v414 = v382
	v415 = v381
	goto L91
L104:
	;
	if v383 == int32(0) {
		goto L90
	} else {
		goto L108
	}
L105:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+97)))
	if v387 != int32(1) {
		goto L90
	} else {
		goto L106
	}
L106:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v350)+84))
	if v390 == int32(0) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v393 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v348)+15)) = uint8(v393)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v383)+56))
	v414 = v395
	v415 = v390
	goto L91
L108:
	;
	v401 = F_ExecGetResultSlotOps(m, v383, v348+int32(15))
	mBase = m.M
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v383)+56))
	v414 = v402
	v415 = v401
	goto L91
L109:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v350)+80))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v350)+76))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+100)))
	if v409 != int32(1) {
		v414 = v408
		v415 = v407
		goto L91
	} else {
		goto L110
	}
L110:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v348)+15)) = uint8(v412)
	v414 = v408
	v415 = v407
	goto L91
L111:
	;
	if v415 != 0 {
		goto L89
	} else {
		goto L112
	}
L112:
	;
	goto L90
L113:
	;
	v440 = int32(0)
	goto L87
L114:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v446 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v468 + int32(1)
	v474 = v467 + v468*int32(40)
	v475 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v474)+32)) = v475
	v477 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v474)+24)) = v477
	v479 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v474)+16)) = v479
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v474)+8)) = v481
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v474))) = v483
	goto L84
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v465
	v467 = v465
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v452 = F_palloc(m, int32(640))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L37
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v454 != v446 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v465 = v452
	goto L116
L121:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v467 = v456
	goto L115
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v446 << (uint(int32(1)) % 32)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v463 = F_repalloc(m, v460, v446*int32(80))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L37
	} else {
		goto L124
	}
L124:
	;
	v465 = v463
	goto L116
L125:
	;
	v643 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
	if v643 <= int32(0) {
		goto L166
	} else {
		goto L167
	}
L126:
	;
	v490 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v490)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(5)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v498 = v9 + int32(8)
	v502 = m.G0
	v504 = v502 - int32(16)
	m.G0 = v504
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v504)+15)) = uint8(v490)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v498)+24))
	if v509 != 0 {
		goto L133
	} else {
		goto L134
	}
L127:
	;
	if v596 == int32(0) {
		goto L125
	} else {
		goto L155
	}
L128:
	;
	m.G0 = v504 + int32(16)
	goto L127
L129:
	;
	v596 = int32(1)
	goto L128
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v498)+28)) = v571
	v585 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+20)) = uint8(v585)
	*(*int32)(unsafe.Add(mBase, uint32(v498)+24)) = v570
	if v571 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L129
	} else {
		goto L154
	}
L131:
	;
	v580 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+20)) = uint8(v580)
	*(*int64)(unsafe.Add(mBase, uint32(v498)+24)) = int64(0)
	goto L129
L132:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504)+15)))
	if base.B2i32(v570 == int32(0))|base.B2i32(v574 != int32(1)) != 0 {
		goto L131
	} else {
		goto L152
	}
L133:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v498)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v504)+15)) = uint8(base.B2i32(v510 != int32(0)))
	v570 = v509
	v571 = v510
	goto L132
L134:
	;
	goto L135
L135:
	;
	if v506 == int32(0) {
		goto L131
	} else {
		goto L136
	}
L136:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	switch v516 - int32(2) {
	case 0:
		goto L139
	case 1:
		goto L138
	default:
		goto L137
	}
L137:
	;
	if base.Ui32(int32(2)) < base.Ui32(v516-int32(4)) {
		goto L131
	} else {
		goto L150
	}
L138:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v506)+36))
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+101)))
	if v540 != int32(1) {
		goto L145
	} else {
		goto L146
	}
L139:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v506)+40))
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+102)))
	if v520 != int32(1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if v519 == int32(0) {
		goto L131
	} else {
		goto L144
	}
L141:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+98)))
	if v523 != int32(1) {
		goto L131
	} else {
		goto L142
	}
L142:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v506)+88))
	if v526 == int32(0) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v529 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v504)+15)) = uint8(v529)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v519)+56))
	v570 = v531
	v571 = v526
	goto L132
L144:
	;
	v537 = F_ExecGetResultSlotOps(m, v519, v504+int32(15))
	mBase = m.M
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v519)+56))
	v570 = v538
	v571 = v537
	goto L132
L145:
	;
	if v539 == int32(0) {
		goto L131
	} else {
		goto L149
	}
L146:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+97)))
	if v543 != int32(1) {
		goto L131
	} else {
		goto L147
	}
L147:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v506)+84))
	if v546 == int32(0) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v549 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v504)+15)) = uint8(v549)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v539)+56))
	v570 = v551
	v571 = v546
	goto L132
L149:
	;
	v557 = F_ExecGetResultSlotOps(m, v539, v504+int32(15))
	mBase = m.M
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v539)+56))
	v570 = v558
	v571 = v557
	goto L132
L150:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v506)+80))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v506)+76))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+100)))
	if v565 != int32(1) {
		v570 = v564
		v571 = v563
		goto L132
	} else {
		goto L151
	}
L151:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v504)+15)) = uint8(v568)
	v570 = v564
	v571 = v563
	goto L132
L152:
	;
	if v571 != 0 {
		goto L130
	} else {
		goto L153
	}
L153:
	;
	goto L131
L154:
	;
	v596 = int32(0)
	goto L128
L155:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v602 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v624 + int32(1)
	v630 = v623 + v624*int32(40)
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v630)+32)) = v631
	v633 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v630)+24)) = v633
	v635 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v630)+16)) = v635
	v637 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v630)+8)) = v637
	v639 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v630))) = v639
	goto L125
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v621
	v623 = v621
	goto L156
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v608 = F_palloc(m, int32(640))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L37
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v610 != v602 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v621 = v608
	goto L157
L162:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v623 = v612
	goto L156
L163:
	;
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v602 << (uint(int32(1)) % 32)
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v619 = F_repalloc(m, v616, v602*int32(80))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L37
	} else {
		goto L165
	}
L165:
	;
	v621 = v619
	goto L157
L166:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v799 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L167:
	;
	v646 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v646)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v654 = v9 + int32(8)
	v658 = m.G0
	v660 = v658 - int32(16)
	m.G0 = v660
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v660)+15)) = uint8(v646)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v654)+24))
	if v665 != 0 {
		goto L174
	} else {
		goto L175
	}
L168:
	;
	if v752 == int32(0) {
		goto L166
	} else {
		goto L196
	}
L169:
	;
	m.G0 = v660 + int32(16)
	goto L168
L170:
	;
	v752 = int32(1)
	goto L169
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v654)+28)) = v727
	v741 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v654)+20)) = uint8(v741)
	*(*int32)(unsafe.Add(mBase, uint32(v654)+24)) = v726
	if v727 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L170
	} else {
		goto L195
	}
L172:
	;
	v736 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v654)+20)) = uint8(v736)
	*(*int64)(unsafe.Add(mBase, uint32(v654)+24)) = int64(0)
	goto L170
L173:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+15)))
	if base.B2i32(v726 == int32(0))|base.B2i32(v730 != int32(1)) != 0 {
		goto L172
	} else {
		goto L193
	}
L174:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v654)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v660)+15)) = uint8(base.B2i32(v666 != int32(0)))
	v726 = v665
	v727 = v666
	goto L173
L175:
	;
	goto L176
L176:
	;
	if v662 == int32(0) {
		goto L172
	} else {
		goto L177
	}
L177:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v654)))
	switch v672 - int32(2) {
	case 0:
		goto L180
	case 1:
		goto L179
	default:
		goto L178
	}
L178:
	;
	if base.Ui32(int32(2)) < base.Ui32(v672-int32(4)) {
		goto L172
	} else {
		goto L191
	}
L179:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v662)+36))
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+101)))
	if v696 != int32(1) {
		goto L186
	} else {
		goto L187
	}
L180:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v662)+40))
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+102)))
	if v676 != int32(1) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if v675 == int32(0) {
		goto L172
	} else {
		goto L185
	}
L182:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+98)))
	if v679 != int32(1) {
		goto L172
	} else {
		goto L183
	}
L183:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v662)+88))
	if v682 == int32(0) {
		goto L181
	} else {
		goto L184
	}
L184:
	;
	v685 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v660)+15)) = uint8(v685)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v675)+56))
	v726 = v687
	v727 = v682
	goto L173
L185:
	;
	v693 = F_ExecGetResultSlotOps(m, v675, v660+int32(15))
	mBase = m.M
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v675)+56))
	v726 = v694
	v727 = v693
	goto L173
L186:
	;
	if v695 == int32(0) {
		goto L172
	} else {
		goto L190
	}
L187:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+97)))
	if v699 != int32(1) {
		goto L172
	} else {
		goto L188
	}
L188:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v662)+84))
	if v702 == int32(0) {
		goto L186
	} else {
		goto L189
	}
L189:
	;
	v705 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v660)+15)) = uint8(v705)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v695)+56))
	v726 = v707
	v727 = v702
	goto L173
L190:
	;
	v713 = F_ExecGetResultSlotOps(m, v695, v660+int32(15))
	mBase = m.M
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v695)+56))
	v726 = v714
	v727 = v713
	goto L173
L191:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v662)+80))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v662)+76))
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+100)))
	if v721 != int32(1) {
		v726 = v720
		v727 = v719
		goto L173
	} else {
		goto L192
	}
L192:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v660)+15)) = uint8(v724)
	v726 = v720
	v727 = v719
	goto L173
L193:
	;
	if v727 != 0 {
		goto L171
	} else {
		goto L194
	}
L194:
	;
	goto L172
L195:
	;
	v752 = int32(0)
	goto L169
L196:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v758 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v780 + int32(1)
	v786 = v779 + v780*int32(40)
	v787 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v786)+32)) = v787
	v789 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v786)+24)) = v789
	v791 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v786)+16)) = v791
	v793 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v786)+8)) = v793
	v795 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v786))) = v795
	goto L166
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v777
	v779 = v777
	goto L197
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v764 = F_palloc(m, int32(640))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L37
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v766 != v758 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v777 = v764
	goto L198
L203:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v779 = v768
	goto L197
L204:
	;
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v758 << (uint(int32(1)) % 32)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v775 = F_repalloc(m, v772, v758*int32(80))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L37
	} else {
		goto L206
	}
L206:
	;
	v777 = v775
	goto L198
L207:
	;
	m.G0 = v9 + int32(48)
	return
L208:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v799)+4))
	if v802 <= int32(0) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v811 = int32(0)
	goto L210
L210:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v799)+12))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v816+v811<<(uint(int32(2))%32))))
	F_ExecInitSubPlanExpr(m, v820, l0, l0+int32(8), l0+int32(5))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L37
	} else {
		goto L212
	}
L211:
	;
	goto L207
L212:
	;
	v824 = v811 + int32(1)
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v799)+4))
	if v824 < v825 {
		v811 = v824
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
}
func F_assign_expr_collations(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
	v15 = F_assign_collations_walker(m, l1, v6+int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v6 + int32(32)
		return
	}
}
func F_checkExprIsVarFree(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_contain_vars_of_level(m, l1, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_checkExprIsVarFree_0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
					F_errmsg(m, int32(_a_F_checkExprIsVarFree_1), v7)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v24 = F_locate_var_of_level(m, l1, int32(0))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_parser_errposition(m, l0, v24)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_checkExprIsVarFree_2), int32(1935), int32(_a_F_checkExprIsVarFree_3))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
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
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_exprInputCollation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v2 = int32(0)
	if l0 == v2 {
		v30 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = v7 - int32(9)
		if base.Ui32(int32(30)) < base.Ui32(v9) {
			v30 = v2
		} else {
			v13 = int32(1) << (uint(v9) % 32)
			if v13&int32(3904) == int32(0) {
				if v13&int32(5) != 0 {
					v25 = int32(16)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25+l0)))
					v30 = v27
				} else {
					if v9 != int32(30) {
						v30 = v2
					} else {
						v25 = int32(12)
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25+l0)))
						v30 = v27
					}
				}
			} else {
				v25 = int32(24)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v25+l0)))
				v30 = v27
			}
		}
	}
	return v30
}
func F_fix_scan_expr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 float64
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v169 float64
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 float64
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 float64
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	if l0 == int32(0) {
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
	v16 = l0
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	switch v27 - int32(6) {
	case 0:
		goto L12
	default:
		v123 = v27
		goto L9
	case 2:
		goto L11
	case 3:
		goto L10
	}
L5:
	;
	return int32(0)
L6:
	;
	if v219 != 0 {
		v16 = v219
		goto L4
	} else {
		goto L56
	}
L7:
	;
	v216 = F_copyObjectImpl(m, v108)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L13
	} else {
		goto L55
	}
L8:
	;
	return v206
L9:
	;
	if v123 != int32(24) {
		goto L36
	} else {
		goto L37
	}
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+276))
	if v65 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v61 = F_fix_param_node(m, v60, v16)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L13
	} else {
		goto L19
	}
L12:
	;
	v31 = F_palloc(m, int32(48))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v35
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+40)) = v37
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if int32(0) <= v47 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v50 + v47
	goto L17
L16:
	;
	goto L17
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if v53 == int32(0) {
		v206 = v31
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = v56 + v53
	return v31
L19:
	;
	return v61
L20:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v123 = v120
	goto L9
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v68 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v71 != int32(1) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v74 <= int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v82 = int32(0)
	v84 = v74
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v82<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v96 == v97 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
	if v108 != 0 {
		goto L7
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v101 = F_equal(m, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L31
	}
L29:
	;
	v104 = v84
	goto L30
L30:
	;
	v106 = v82 + int32(1)
	if v106 < v104 {
		v82 = v106
		v84 = v104
		goto L25
	} else {
		goto L33
	}
L31:
	;
	if v101 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v104 = v103
	goto L30
L33:
	;
	goto L20
L34:
	;
	goto L20
L35:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_fix_expr_common(m, v198, v16)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L13
	} else {
		goto L53
	}
L36:
	;
	if v123 != int32(319) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v147 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v149 = int32(0)
	v152 = v149
	v154 = v149
	v159 = float64(0)
	goto L44
L39:
	;
	if v123 != int32(58) {
		goto L35
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v219 = v145
	goto L6
L42:
	;
	v138 = F_copyObjectImpl(m, v16)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v140 + v141
	return v138
L44:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v154<<(uint(int32(2))%32))))
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v167)+56))
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v167)+64))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v148)+360))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v170+v171-v173))) = uint8(v173)
	v178 = base.F64_add(v168, base.F64_mul(v147, v169))
	v180 = int32(0)
	v184 = base.B2i32(base.F64_ge(v159, v178) == v180) & base.B2i32(v152 != v180)
	if v184 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v148)+364))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v194 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v191+v192-v194))) = uint8(v194)
	v219 = v186
	goto L6
L46:
	;
	v185 = v159
	goto L48
L47:
	;
	v185 = v178
	goto L48
L48:
	;
	if v184 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v186 = v152
	goto L51
L50:
	;
	v186 = v167
	goto L51
L51:
	;
	v188 = v154 + int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v188 < v189 {
		v152 = v186
		v154 = v188
		v159 = v185
		goto L44
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	v202 = F_expression_tree_mutator_impl(m, v16, int32(839), l1)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	v206 = v202
	goto L8
L55:
	;
	return v216
L56:
	;
	goto L5
}
func F_format_expr_params(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v323 int32
	_ = v323
	var v337 int32
	_ = v337
	var v348 int32
	_ = v348
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(_a_F_format_expr_params_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v19
	v22 = v10 + int32(-24)
	F_initStringInfo(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v348 = int32(0)
	goto L3
L3:
	;
	m.G0 = v12 - int32(-64)
	return v348
L4:
	;
	return int32(0)
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v27 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v16
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v348 = v337
	goto L3
L7:
	;
	if v84 < int32(0) {
		goto L6
	} else {
		goto L18
	}
L8:
	;
	v84 = base.I32_ctz(v70) | v71<<(uint(int32(5))%32)
	goto L7
L9:
	;
	v84 = int32(-2)
	goto L7
L10:
	;
	v37 = base.I32_div_s(int32(0), int32(32))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v38 <= v37 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v41 = v27 + int32(8)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v37<<(uint(int32(2))%32))))
	v48 = v45 & int32(-1)
	if v48 != 0 {
		v70 = v48
		v71 = v37
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v50 = v37 + int32(1)
	if v50 == v38 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v53 = v50
	goto L14
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41+v53<<(uint(int32(2))%32))))
	if v60 != 0 {
		v70 = v60
		v71 = v53
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L9
L16:
	;
	v62 = v53 + int32(1)
	if v62 != v38 {
		v53 = v62
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v84<<(uint(int32(2))%32))))
	F_exec_eval_datum(m, l0, v91, v10+int32(-32), v10+int32(-40), v10+int32(-28), v10+int32(-33))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_format_expr_params_1)
	F_appendStringInfo(m, v22, int32(_a_F_format_expr_params_2), v10+int32(-48))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)))
	if v111 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v143 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v116 = int32(_a_F_format_expr_params_0)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v120
	F_getTypeOutputInfo(m, v115, v10+int32(-4), v10+int32(-5))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_appendStringInfoString(m, v10+int32(-24), int32(_a_F_format_expr_params_3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L28
	}
L25:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v129 = F_OidOutputFunctionCall(m, v128, v114)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v117
	F_appendStringInfoStringQuoted(m, v22, v129, int32(-1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	goto L21
L29:
	;
	if v199 < int32(0) {
		goto L6
	} else {
		goto L40
	}
L30:
	;
	v199 = base.I32_ctz(v185) | v186<<(uint(int32(5))%32)
	goto L29
L31:
	;
	v199 = int32(-2)
	goto L29
L32:
	;
	v150 = v84 + int32(1)
	v152 = base.I32_div_s(v150, int32(32))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v153 <= v152 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v156 = v143 + int32(8)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v152<<(uint(int32(2))%32))))
	v163 = v160 & (int32(-1) << (uint(v150) % 32))
	if v163 != 0 {
		v185 = v163
		v186 = v152
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v165 = v152 + int32(1)
	if v165 == v153 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v168 = v165
	goto L36
L36:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v156+v168<<(uint(int32(2))%32))))
	if v175 != 0 {
		v185 = v175
		v186 = v168
		goto L30
	} else {
		goto L38
	}
L37:
	;
	goto L31
L38:
	;
	v177 = v168 + int32(1)
	if v177 != v153 {
		v168 = v177
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v206 = v199
	goto L41
L41:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+v206<<(uint(int32(2))%32))))
	F_exec_eval_datum(m, l0, v215, v10+int32(-32), v10+int32(-40), v10+int32(-28), v10+int32(-33))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	goto L6
L43:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_format_expr_params_4)
	v231 = v10 + int32(-24)
	F_appendStringInfo(m, v231, int32(_a_F_format_expr_params_2), v12)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)))
	if v235 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v267 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	F_appendStringInfoString(m, v231, int32(_a_F_format_expr_params_3))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v243 = int32(_a_F_format_expr_params_0)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v247
	F_getTypeOutputInfo(m, v242, v10+int32(-4), v10+int32(-5))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	goto L45
L50:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v256 = F_OidOutputFunctionCall(m, v255, v241)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v244
	F_appendStringInfoStringQuoted(m, v10+int32(-24), v256, int32(-1))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	if int32(0) <= v323 {
		v206 = v323
		goto L41
	} else {
		goto L64
	}
L54:
	;
	v323 = base.I32_ctz(v309) | v310<<(uint(int32(5))%32)
	goto L53
L55:
	;
	v323 = int32(-2)
	goto L53
L56:
	;
	v274 = v206 + int32(1)
	v276 = base.I32_div_s(v274, int32(32))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	if v277 <= v276 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v280 = v267 + int32(8)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v280+v276<<(uint(int32(2))%32))))
	v287 = v284 & (int32(-1) << (uint(v274) % 32))
	if v287 != 0 {
		v309 = v287
		v310 = v276
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v289 = v276 + int32(1)
	if v289 == v277 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v292 = v289
	goto L60
L60:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v280+v292<<(uint(int32(2))%32))))
	if v299 != 0 {
		v309 = v299
		v310 = v292
		goto L54
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v301 = v292 + int32(1)
	if v301 != v277 {
		v292 = v301
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L42
}
