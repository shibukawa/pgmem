package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecBuildGroupingEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
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
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v472 int64
	_ = v472
	var v474 int64
	_ = v474
	var v476 int64
	_ = v476
	var v478 int64
	_ = v478
	var v480 int64
	_ = v480
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int64
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int64
	_ = v619
	var v622 int32
	_ = v622
	var v623 int64
	_ = v623
	var v626 int32
	_ = v626
	var v627 int64
	_ = v627
	var v629 int64
	_ = v629
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v675 int64
	_ = v675
	var v677 int64
	_ = v677
	var v679 int64
	_ = v679
	var v681 int64
	_ = v681
	var v683 int64
	_ = v683
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v724 int64
	_ = v724
	var v726 int64
	_ = v726
	var v728 int64
	_ = v728
	var v730 int64
	_ = v730
	var v732 int64
	_ = v732
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v769 int64
	_ = v769
	var v771 int64
	_ = v771
	var v773 int64
	_ = v773
	var v775 int64
	_ = v775
	var v777 int64
	_ = v777
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v802 int32
	_ = v802
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v888 int64
	_ = v888
	var v890 int64
	_ = v890
	var v892 int64
	_ = v892
	var v894 int64
	_ = v894
	var v896 int64
	_ = v896
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v911 int32
	_ = v911
	v10 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(48)
	m.G0 = v26
	v29 = F_palloc0(m, int32(68))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(380)
	v35 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v35
	if l4 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v26 + int32(48)
	return v911
L4:
	;
	v911 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = l8
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)) = uint8(v47)
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v49
	v52 = v29 + int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v52
	v55 = v29 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v55
	if l4 <= v49 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(2)
	v200 = v26 + int32(8)
	v201 = int32(0)
	v204 = m.G0
	v206 = v204 - int32(16)
	m.G0 = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+15)) = uint8(v201)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v200)+24))
	if v211 != 0 {
		goto L43
	} else {
		goto L44
	}
L8:
	;
	v179 = int32(-1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v61 = l4 & int32(3)
	if base.Ui32(l4) < base.Ui32(int32(4)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v61 == int32(0) {
		v179 = v119
		goto L7
	} else {
		goto L30
	}
L12:
	;
	v119 = int32(-1)
	v122 = v10
	goto L11
L13:
	;
	goto L14
L14:
	;
	v76 = int32(-1)
	v79 = v10
	v82 = v10
	goto L15
L15:
	;
	v93 = l5 + v79<<(uint(int32(1))%32)
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93))))
	if v94 < v76 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v119 = v105
	v122 = v107
	goto L11
L17:
	;
	v96 = v76
	goto L19
L18:
	;
	v96 = v94
	goto L19
L19:
	;
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+2)))
	if v97 < v96 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v99 = v96
	goto L22
L21:
	;
	v99 = v97
	goto L22
L22:
	;
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+4)))
	if v100 < v99 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v102 = v99
	goto L25
L24:
	;
	v102 = v100
	goto L25
L25:
	;
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+6)))
	if v103 < v102 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v105 = v102
	goto L28
L27:
	;
	v105 = v103
	goto L28
L28:
	;
	v106 = int32(4)
	v107 = v79 + v106
	v109 = v82 + v106
	if v109 != l4&int32(2147483644) {
		v76 = v105
		v79 = v107
		v82 = v109
		goto L15
	} else {
		goto L29
	}
L29:
	;
	goto L16
L30:
	;
	v145 = v119
	v148 = v122
	v150 = int32(0)
	goto L31
L31:
	;
	v163 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5+v148<<(uint(int32(1))%32)))))
	if v163 < v145 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v179 = v165
	goto L7
L33:
	;
	v165 = v145
	goto L35
L34:
	;
	v165 = v163
	goto L35
L35:
	;
	v166 = int32(1)
	v169 = v150 + v166
	if v169 != v61 {
		v145 = v165
		v148 = v148 + v166
		v150 = v169
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	if v293 != 0 {
		goto L65
	} else {
		goto L66
	}
L38:
	;
	m.G0 = v206 + int32(16)
	goto L37
L39:
	;
	v293 = int32(1)
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+28)) = v269
	v282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v200)+20)) = uint8(v282)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = v268
	if v269 != int32(1575684) {
		goto L39
	} else {
		goto L64
	}
L41:
	;
	v277 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v200)+20)) = uint8(v277)
	*(*int64)(unsafe.Add(mBase, uint32(v200)+24)) = int64(0)
	goto L39
L42:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+15)))
	if v270 != int32(1) {
		goto L41
	} else {
		goto L61
	}
L43:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v200)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+15)) = uint8(base.B2i32(v212 != int32(0)))
	v268 = v211
	v269 = v212
	goto L42
L44:
	;
	goto L45
L45:
	;
	if v208 == int32(0) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	switch v218 - int32(2) {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2, 3, 4:
		goto L47
	default:
		goto L41
	}
L47:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v208)+80))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v208)+76))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+100)))
	if v263 != int32(1) {
		v268 = v262
		v269 = v261
		goto L42
	} else {
		goto L60
	}
L48:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v208)+36))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+101)))
	if v242 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v208)+40))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+102)))
	if v222 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v221 == int32(0) {
		goto L41
	} else {
		goto L54
	}
L51:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+98)))
	if v225 != int32(1) {
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v208)+88))
	if v228 == int32(0) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+15)) = uint8(v231)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v221)+56))
	v268 = v233
	v269 = v228
	goto L42
L54:
	;
	v239 = F_ExecGetResultSlotOps(m, v221, v206+int32(15))
	mBase = m.M
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v221)+56))
	v268 = v240
	v269 = v239
	goto L42
L55:
	;
	if v241 == int32(0) {
		goto L41
	} else {
		goto L59
	}
L56:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+97)))
	if v245 != int32(1) {
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v208)+84))
	if v248 == int32(0) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v251 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+15)) = uint8(v251)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v241)+56))
	v268 = v253
	v269 = v248
	goto L42
L59:
	;
	v259 = F_ExecGetResultSlotOps(m, v241, v206+int32(15))
	mBase = m.M
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v241)+56))
	v268 = v260
	v269 = v259
	goto L42
L60:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+15)) = uint8(v266)
	v268 = v262
	v269 = v261
	goto L42
L61:
	;
	if v268 == int32(0) {
		goto L41
	} else {
		goto L62
	}
L62:
	;
	if v269 != 0 {
		goto L40
	} else {
		goto L63
	}
L63:
	;
	goto L41
L64:
	;
	v293 = int32(0)
	goto L38
L65:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v297 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = l1
	v340 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+28)) = uint8(v340)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(3)
	v346 = v26 + int32(8)
	v350 = m.G0
	v352 = v350 - int32(16)
	m.G0 = v352
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+15)) = uint8(v340)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v346)+24))
	if v357 != 0 {
		goto L84
	} else {
		goto L85
	}
L68:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v319 + int32(1)
	v325 = v318 + v319*int32(40)
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v26)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v325)+32)) = v326
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v325)+24)) = v328
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v325)+16)) = v330
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v26)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v325)+8)) = v332
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v325))) = v334
	goto L67
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v316
	v318 = v316
	goto L68
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(16)
	v303 = F_palloc(m, int32(640))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v305 != v297 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v316 = v303
	goto L69
L74:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v318 = v307
	goto L68
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v297 << (uint(int32(1)) % 32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v314 = F_repalloc(m, v311, v297*int32(80))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v316 = v314
	goto L69
L78:
	;
	if v439 != 0 {
		goto L106
	} else {
		goto L107
	}
L79:
	;
	m.G0 = v352 + int32(16)
	goto L78
L80:
	;
	v439 = int32(1)
	goto L79
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+28)) = v415
	v428 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+20)) = uint8(v428)
	*(*int32)(unsafe.Add(mBase, uint32(v346)+24)) = v414
	if v415 != int32(1575684) {
		goto L80
	} else {
		goto L105
	}
L82:
	;
	v423 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+20)) = uint8(v423)
	*(*int64)(unsafe.Add(mBase, uint32(v346)+24)) = int64(0)
	goto L80
L83:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+15)))
	if v416 != int32(1) {
		goto L82
	} else {
		goto L102
	}
L84:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v346)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+15)) = uint8(base.B2i32(v358 != int32(0)))
	v414 = v357
	v415 = v358
	goto L83
L85:
	;
	goto L86
L86:
	;
	if v354 == int32(0) {
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	switch v364 - int32(2) {
	case 0:
		goto L90
	case 1:
		goto L89
	case 2, 3, 4:
		goto L88
	default:
		goto L82
	}
L88:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v354)+80))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v354)+76))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+100)))
	if v409 != int32(1) {
		v414 = v408
		v415 = v407
		goto L83
	} else {
		goto L101
	}
L89:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v354)+36))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+101)))
	if v388 != int32(1) {
		goto L96
	} else {
		goto L97
	}
L90:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v354)+40))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+102)))
	if v368 != int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v367 == int32(0) {
		goto L82
	} else {
		goto L95
	}
L92:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+98)))
	if v371 != int32(1) {
		goto L82
	} else {
		goto L93
	}
L93:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v354)+88))
	if v374 == int32(0) {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v377 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+15)) = uint8(v377)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v367)+56))
	v414 = v379
	v415 = v374
	goto L83
L95:
	;
	v385 = F_ExecGetResultSlotOps(m, v367, v352+int32(15))
	mBase = m.M
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v367)+56))
	v414 = v386
	v415 = v385
	goto L83
L96:
	;
	if v387 == int32(0) {
		goto L82
	} else {
		goto L100
	}
L97:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+97)))
	if v391 != int32(1) {
		goto L82
	} else {
		goto L98
	}
L98:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v354)+84))
	if v394 == int32(0) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v397 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+15)) = uint8(v397)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v387)+56))
	v414 = v399
	v415 = v394
	goto L83
L100:
	;
	v405 = F_ExecGetResultSlotOps(m, v387, v352+int32(15))
	mBase = m.M
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v387)+56))
	v414 = v406
	v415 = v405
	goto L83
L101:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+15)) = uint8(v412)
	v414 = v408
	v415 = v407
	goto L83
L102:
	;
	if v414 == int32(0) {
		goto L82
	} else {
		goto L103
	}
L103:
	;
	if v415 != 0 {
		goto L81
	} else {
		goto L104
	}
L104:
	;
	goto L82
L105:
	;
	v439 = int32(0)
	goto L79
L106:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v443 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	goto L108
L108:
	;
	v485 = l4 - int32(1)
	if v485 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L109:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v465 + int32(1)
	v471 = v464 + v465*int32(40)
	v472 = *(*int64)(unsafe.Add(mBase, uint32(v26)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v471)+32)) = v472
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v471)+24)) = v474
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v471)+16)) = v476
	v478 = *(*int64)(unsafe.Add(mBase, uint32(v26)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v471)+8)) = v478
	v480 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v471))) = v480
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v462
	v464 = v462
	goto L109
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(16)
	v449 = F_palloc(m, int32(640))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v451 != v443 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v462 = v449
	goto L110
L115:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v464 = v453
	goto L109
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v443 << (uint(int32(1)) % 32)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v460 = F_repalloc(m, v457, v443*int32(80))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v462 = v460
	goto L110
L119:
	;
	v855 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v855
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = int64(0)
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v859 == v855 {
		goto L185
	} else {
		goto L186
	}
L120:
	;
	v488 = int32(88)
	v495 = int32(0)
	v501 = v485
	goto L121
L121:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v521 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5+v501<<(uint(int32(1))%32)))))
	v523 = v501 << (uint(int32(2)) % 32)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l7+v523)))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l6+v523)))
	v530 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v532 = F_object_aclcheck(m, int32(1255), v528, v530, int64(128))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L123
	}
L122:
	;
	if v784 == int32(0) {
		goto L119
	} else {
		goto L178
	}
L123:
	;
	if v532 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v535 = F_get_func_name(m, v528)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v542 = v521 - int32(1)
	v544 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	if v544 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	F_aclcheck_error(m, v532, int32(19), v535)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	F_RunFunctionExecuteHook(m, v528)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v550 = F_palloc0(m, int32(28))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	v553 = F_palloc0(m, int32(36))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_fmgr_info(m, v528, v550)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v557 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v550)+24)) = v557
	v559 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v553)+18)) = uint16(v559)
	*(*uint8)(unsafe.Add(mBase, uint32(v553)+16)) = uint8(v557)
	*(*int32)(unsafe.Add(mBase, uint32(v553)+12)) = v525
	*(*int64)(unsafe.Add(mBase, uint32(v553)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v553))) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v542
	v571 = v542 * int32(100)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v571+(v517<<(uint(int32(4))%32)+(l0+v488)))))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v574
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v553 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v553 + int32(20)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v584 == v557 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v606 + int32(1)
	v610 = int32(40)
	v612 = v605 + v606*v610
	v614 = v26 + v610
	v615 = *(*int64)(unsafe.Add(mBase, uint32(v614)))
	*(*int64)(unsafe.Add(mBase, uint32(v612)+32)) = v615
	v617 = int32(32)
	v618 = v26 + v617
	v619 = *(*int64)(unsafe.Add(mBase, uint32(v618)))
	*(*int64)(unsafe.Add(mBase, uint32(v612)+24)) = v619
	v622 = v26 + int32(24)
	v623 = *(*int64)(unsafe.Add(mBase, uint32(v622)))
	*(*int64)(unsafe.Add(mBase, uint32(v612)+16)) = v623
	v626 = v26 + int32(16)
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v626)))
	*(*int64)(unsafe.Add(mBase, uint32(v612)+8)) = v627
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v612))) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(8)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l1+v488+v516<<(uint(int32(4))%32)+v571)))
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = v553 + v617
	v640 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v618))) = v640
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v553 + int32(28)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v646 == v640 {
		goto L148
	} else {
		goto L149
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v603
	v605 = v603
	goto L136
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(16)
	v590 = F_palloc(m, int32(640))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v592 != v584 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v603 = v590
	goto L137
L142:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v605 = v594
	goto L136
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v584 << (uint(int32(1)) % 32)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v601 = F_repalloc(m, v598, v584*int32(80))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v603 = v601
	goto L137
L146:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v668 + int32(1)
	v674 = v667 + v668*int32(40)
	v675 = *(*int64)(unsafe.Add(mBase, uint32(v614)))
	*(*int64)(unsafe.Add(mBase, uint32(v674)+32)) = v675
	v677 = *(*int64)(unsafe.Add(mBase, uint32(v618)))
	*(*int64)(unsafe.Add(mBase, uint32(v674)+24)) = v677
	v679 = *(*int64)(unsafe.Add(mBase, uint32(v622)))
	*(*int64)(unsafe.Add(mBase, uint32(v674)+16)) = v679
	v681 = *(*int64)(unsafe.Add(mBase, uint32(v626)))
	*(*int64)(unsafe.Add(mBase, uint32(v674)+8)) = v681
	v683 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v674))) = v683
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v553
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(62)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	*(*int32)(unsafe.Add(mBase, uint32(v618))) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(2)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v695 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v665
	v667 = v665
	goto L146
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(16)
	v652 = F_palloc(m, int32(640))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v654 != v646 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v665 = v652
	goto L147
L152:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v667 = v656
	goto L146
L153:
	;
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v646 << (uint(int32(1)) % 32)
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v663 = F_repalloc(m, v660, v646*int32(80))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v665 = v663
	goto L147
L156:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v717 + int32(1)
	v723 = v716 + v717*int32(40)
	v724 = *(*int64)(unsafe.Add(mBase, uint32(v614)))
	*(*int64)(unsafe.Add(mBase, uint32(v723)+32)) = v724
	v726 = *(*int64)(unsafe.Add(mBase, uint32(v618)))
	*(*int64)(unsafe.Add(mBase, uint32(v723)+24)) = v726
	v728 = *(*int64)(unsafe.Add(mBase, uint32(v622)))
	*(*int64)(unsafe.Add(mBase, uint32(v723)+16)) = v728
	v730 = *(*int64)(unsafe.Add(mBase, uint32(v626)))
	*(*int64)(unsafe.Add(mBase, uint32(v723)+8)) = v730
	v732 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v723))) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v55
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v740 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v714
	v716 = v714
	goto L156
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(16)
	v701 = F_palloc(m, int32(640))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v703 != v695 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v714 = v701
	goto L157
L162:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v716 = v705
	goto L156
L163:
	;
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v695 << (uint(int32(1)) % 32)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v712 = F_repalloc(m, v709, v695*int32(80))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v714 = v712
	goto L157
L166:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v763 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v762 + v763
	v768 = v761 + v762*int32(40)
	v769 = *(*int64)(unsafe.Add(mBase, uint32(v614)))
	*(*int64)(unsafe.Add(mBase, uint32(v768)+32)) = v769
	v771 = *(*int64)(unsafe.Add(mBase, uint32(v618)))
	*(*int64)(unsafe.Add(mBase, uint32(v768)+24)) = v771
	v773 = *(*int64)(unsafe.Add(mBase, uint32(v622)))
	*(*int64)(unsafe.Add(mBase, uint32(v768)+16)) = v773
	v775 = *(*int64)(unsafe.Add(mBase, uint32(v626)))
	*(*int64)(unsafe.Add(mBase, uint32(v768)+8)) = v775
	v777 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v768))) = v777
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v784 = F_lappend_int(m, v495, v781-v763)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L176
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v759
	v761 = v759
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(16)
	v746 = F_palloc(m, int32(640))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v748 != v740 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v759 = v746
	goto L167
L172:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v761 = v750
	goto L166
L173:
	;
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v740 << (uint(int32(1)) % 32)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v757 = F_repalloc(m, v754, v740*int32(80))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v759 = v757
	goto L167
L176:
	;
	if int32(0) < v501 {
		v495 = v784
		v501 = v501 - v763
		goto L121
	} else {
		goto L177
	}
L177:
	;
	goto L122
L178:
	;
	v790 = int32(0)
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v784)+4))
	if v791 <= v790 {
		goto L119
	} else {
		goto L179
	}
L179:
	;
	v802 = v790
	goto L180
L180:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v784)+12))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v818+v802<<(uint(int32(2))%32))))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v817+v822*int32(40))+16)) = v826
	v829 = v802 + int32(1)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v784)+4))
	if v829 < v830 {
		v802 = v829
		goto L180
	} else {
		goto L182
	}
L181:
	;
	goto L119
L182:
	;
	goto L181
L183:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v881 + int32(1)
	v887 = v880 + v881*int32(40)
	v888 = *(*int64)(unsafe.Add(mBase, uint32(v26)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v887)+32)) = v888
	v890 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v887)+24)) = v890
	v892 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v887)+16)) = v892
	v894 = *(*int64)(unsafe.Add(mBase, uint32(v26)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v887)+8)) = v894
	v896 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v887))) = v896
	v898 = F_jit_compile_expr(m, v29)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L193
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v878
	v880 = v878
	goto L183
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(16)
	v865 = F_palloc(m, int32(640))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v867 != v859 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v878 = v865
	goto L184
L189:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v880 = v869
	goto L183
L190:
	;
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v859 << (uint(int32(1)) % 32)
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v876 = F_repalloc(m, v873, v859*int32(80))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v878 = v876
	goto L184
L193:
	;
	if v898 != 0 {
		v911 = v29
		goto L3
	} else {
		goto L194
	}
L194:
	;
	F_ExecReadyInterpretedExpr(m, v29)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v911 = v29
	goto L3
}
func F_flatten_grouping_sets(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	v4 = int32(0)
	F_check_stack_depth(m)
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
	if l0 == int32(0) {
		v172 = v4
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v172
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == int32(1) {
		v128 = l0
		v129 = l1
		v130 = l2
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v135 = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v136 <= v135 {
		v172 = v135
		goto L3
	} else {
		goto L60
	}
L6:
	;
	if v14 != int32(107) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v73 != 0 {
		goto L42
	} else {
		goto L43
	}
L8:
	;
	if v14 != int32(36) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if l2 != 0 {
		goto L34
	} else {
		goto L35
	}
L11:
	;
	return l0
L12:
	;
	goto L13
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v22 != int32(2) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return l0
L15:
	;
	goto L16
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v26 == int32(0) {
		v172 = v4
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v32 = v26
	goto L19
L19:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v39 != int32(36) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v172 = v4
	goto L3
L21:
	;
	if v39 == int32(107) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v48 != int32(2) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v66 = v32
	v70 = int32(1)
	goto L7
L25:
	;
	goto L26
L26:
	;
	if v39 != int32(1) {
		v172 = v32
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v128 = v32
	v129 = int32(0)
	v130 = int32(0)
	goto L5
L28:
	;
	return v32
L29:
	;
	goto L30
L30:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v52 != 0 {
		v32 = v52
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L20
L33:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v63 == v62 {
		v172 = v62
		goto L3
	} else {
		goto L39
	}
L34:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v55)
	if l1 != 0 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if l1 == int32(0) {
		v66 = l0
		v70 = int32(1)
		goto L7
	} else {
		goto L38
	}
L37:
	;
	v66 = l0
	v70 = v55
	goto L7
L38:
	;
	goto L33
L39:
	;
	v66 = l0
	v70 = v62
	goto L7
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if base.B2i32(v119 == int32(4))&v70 != 0 {
		goto L56
	} else {
		goto L57
	}
L41:
	;
	v83 = int32(0)
	v84 = v74
	goto L46
L42:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v74 < v75 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v114 = int32(0)
	goto L40
L45:
	;
	goto L44
L46:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v84<<(uint(int32(2))%32))))
	v93 = int32(0)
	v95 = F_flatten_grouping_sets(m, v92, v93, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v114 = v107
	goto L40
L48:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v97 != int32(107) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v109 = v84 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v109 < v110 {
		v83 = v107
		v84 = v109
		goto L46
	} else {
		goto L55
	}
L50:
	;
	v105 = F_lappend(m, v83, v95)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L54
	}
L51:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v100 != int32(4) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v103 = F_list_concat(m, v83, v95)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v107 = v103
	goto L49
L54:
	;
	v107 = v105
	goto L49
L55:
	;
	goto L47
L56:
	;
	return v114
L57:
	;
	goto L58
L58:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v125 = F_makeGroupingSet(m, v119, v114, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	return v125
L60:
	;
	v143 = int32(0)
	v144 = v135
	goto L61
L61:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147+v143<<(uint(int32(2))%32))))
	v152 = F_flatten_grouping_sets(m, v151, v129, v130)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	v172 = v163
	goto L3
L63:
	;
	v165 = v143 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v165 < v166 {
		v143 = v165
		v144 = v163
		goto L61
	} else {
		goto L71
	}
L64:
	;
	if v152 == int32(0) {
		v163 = v144
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if v156 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v159 = F_list_concat(m, v144, v152)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v161 = F_lappend(m, v144, v152)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v163 = v159
	goto L63
L70:
	;
	v163 = v161
	goto L63
L71:
	;
	goto L62
}
