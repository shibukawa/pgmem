package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__bt_advance_array_keys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v834 int32
	_ = v834
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v865 int32
	_ = v865
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v936 int32
	_ = v936
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v987 int32
	_ = v987
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1103 int32
	_ = v1103
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1135 int32
	_ = v1135
	v8 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(16)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	if l6 == v8 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	m.G0 = v29 + int32(16)
	return v1135
L2:
	;
	v1109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v1109)
	v1111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+17)) = uint8(v1111)
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+18)))
	if v1114 != v1109 {
		v1135 = v1111
		goto L1
	} else {
		goto L300
	}
L3:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1064 == int32(1) {
		goto L294
	} else {
		goto L295
	}
L4:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+18)))
	if v993 != 0 {
		v1103 = v987
		goto L2
	} else {
		goto L283
	}
L5:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v942 == l2 {
		v1058 = v936
		goto L3
	} else {
		goto L275
	}
L6:
	;
	v883 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l5 + v883
	v893 = F__bt_check_compare(m, l0, v33, l2, l3, l4, int32(0), l6^v883, v29+int32(15), v29+int32(8))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L16
	} else {
		goto L263
	}
L7:
	;
	if l6 != 0 {
		goto L253
	} else {
		goto L254
	}
L8:
	;
	v865 = int32(1)
	v877 = v8
	goto L6
L9:
	;
	v76 = int32(1)
	v86 = int32(0)
	v88 = v8
	v89 = v8
	v93 = v8
	v96 = v76
	v98 = v8
	v100 = v76
	goto L21
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v38 = v36 - int32(1)
	if v38 <= l5 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+26)) = v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v68 <= v65 {
		v865 = int32(1)
		v877 = v8
		goto L6
	} else {
		goto L20
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v60 <= int32(0) {
		goto L8
	} else {
		goto L19
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v38*int32(48)))))
	if v44&int32(32) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v38
	v48 = int32(0)
	v54 = F__bt_check_compare(m, l0, v33, l2, l3, l4, v48, v48, v29+int32(15), v29+int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	if v54 == int32(0) {
		v1135 = v8
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	v74 = v32 + int32(4)
	goto L9
L20:
	;
	v74 = v32 + int32(4)
	goto L9
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v107 = v104 + v86*int32(48)
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+6)))
	if v108 == int32(3) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	if v477&int32(1) == int32(0) {
		v834 = v480
		goto L7
	} else {
		goto L161
	}
L23:
	;
	if v86 < l5 {
		v477 = v89
		v480 = v93
		v481 = v96
		v483 = v100
		goto L33
	} else {
		goto L34
	}
L24:
	;
	v111 = int32(0)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v112&int32(32) == v111 {
		v135 = v111
		v136 = v88
		v137 = v98
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v123 = int32(0)
	switch v33 + int32(1) {
	case 0:
		goto L29
	default:
		v135 = v123
		v136 = v88
		v137 = v98
		goto L23
	case 2:
		goto L30
	}
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v135 = v117 + v88<<(uint(int32(5))%32)
	v136 = v88 + int32(1)
	v137 = v98
	goto L23
L28:
	;
	v135 = v123
	v136 = v88
	v137 = int32(1)
	goto L23
L29:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
	if v129&int32(1) == int32(0) {
		v135 = v123
		v136 = v88
		v137 = v98
		goto L23
	} else {
		goto L32
	}
L30:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
	if v126&int32(2) != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v135 = v123
	v136 = v88
	v137 = v98
	goto L23
L32:
	;
	goto L28
L33:
	;
	v487 = v86 + int32(1)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v487 < v488 {
		v86 = v487
		v88 = v136
		v89 = v477
		v93 = v480
		v96 = v481
		v98 = v137
		v100 = v483
		goto L21
	} else {
		goto L160
	}
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v141 = v139 & int32(196608)
	if v141 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if l5 != v86 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v107)+4)))
	if v144 <= l3 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+18)) = uint8(v146)
	goto L35
L38:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+6)))
	if v152 != int32(3) {
		v477 = v89
		v480 = v93
		v481 = v96
		v483 = v100
		goto L33
	} else {
		goto L41
	}
L39:
	;
	if v135 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v149 = int32(0)
	v477 = int32(1)
	v480 = v93
	v481 = v149
	v483 = v149
	goto L33
L41:
	;
	if v135|v141 == int32(0) {
		v477 = v89
		v480 = v93
		v481 = v96
		v483 = v100
		goto L33
	} else {
		goto L42
	}
L42:
	;
	if v89&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v135 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if v96&int32(1) != 0 {
		goto L66
	} else {
		goto L67
	}
L46:
	;
	v477 = int32(1)
	v480 = v93
	v481 = v96
	v483 = v100
	goto L33
L47:
	;
	goto L48
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v163 != int32(-1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v33 != int32(-1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+18)))
	if v180 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v171 = v163 - int32(1)
	goto L54
L53:
	;
	v171 = int32(0)
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173+v171<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+44)) = v177
	v477 = int32(1)
	v480 = v93
	v481 = v96
	v483 = v100
	goto L33
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+44)) = int32(0)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v191 = v189 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v191
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+19)))
	if v193 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	if v181 == int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v181)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	if v33 == int32(-1) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	if base.B2i32(v189&int32(33554432) == int32(0)) == base.B2i32(v33 == int32(-1)) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v191 | int32(65)
	v477 = int32(1)
	v480 = v93
	v481 = v96
	v483 = v100
	goto L33
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v191 | int32(524288)
	v477 = int32(1)
	v480 = v93
	v481 = v96
	v483 = v100
	goto L33
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v191 | int32(1048576)
	v477 = int32(1)
	v480 = v93
	v481 = v96
	v483 = v100
	goto L33
L65:
	;
	v277 = F_index_getattr_2(m, l2, v219, l4, v29+int32(15))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L16
	} else {
		goto L87
	}
L66:
	;
	v219 = int32(*(*int16)(unsafe.Add(mBase, uint32(v107)+4)))
	if v219 <= l3 {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v222 = int32(0)
	if v135 == v222 {
		v477 = v222
		v480 = v93
		v481 = v96
		v483 = v100
		goto L33
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v225 != int32(-1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v228 = int32(1)
	if v33 != v228 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+18)))
	if v241 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v233 = v225 - v228
	goto L76
L75:
	;
	v233 = int32(0)
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235+v233<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+44)) = v239
	v477 = v222
	v480 = v93
	v481 = v96
	v483 = v100
	goto L33
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+44)) = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v252 = v250 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v252
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+19)))
	if v254 != int32(1) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	if v242 == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	F_pfree(m, v242)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if v33 == int32(1) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	if base.B2i32(v250&int32(33554432) == int32(0)) == base.B2i32(v33 == int32(1)) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v252 | int32(65)
	v477 = v222
	v480 = v93
	v481 = v96
	v483 = v100
	goto L33
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v252 | int32(524288)
	v477 = v222
	v480 = v93
	v481 = v96
	v483 = v100
	goto L33
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v252 | int32(1048576)
	v477 = v222
	v480 = v93
	v481 = v96
	v483 = v100
	goto L33
L87:
	;
	if v135 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if l6 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L89:
	;
	v344 = v339
	v345 = int32(0)
	goto L88
L90:
	;
	v280 = l6 & base.B2i32(l5 == v86)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v281 == int32(-1) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v299 = int32(1)
	v300 = v298 & v299
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	if v301 == v299 {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	F__bt_binsrch_skiparray_skey(m, v280, v33, v277, v284, v135, v107, v29+int32(8))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L16
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	v296 = F__bt_binsrch_array_skey(m, v289+v86*int32(28), v280, v33, v277, v293, v135, v107, v29+int32(8))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L16
	} else {
		goto L97
	}
L96:
	;
	v339 = v284
	goto L89
L97:
	;
	v344 = v293
	v345 = v296
	goto L88
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v334
	v339 = v301
	goto L89
L99:
	;
	if v300 != 0 {
		v334 = int32(0)
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v300 != 0 {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	if v298&int32(33554432) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v309 = int32(-1)
	goto L105
L104:
	;
	v309 = int32(1)
	goto L105
L105:
	;
	v334 = v309
	goto L98
L106:
	;
	if v298&int32(33554432) != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	v321 = F_FunctionCall2Coll(m, v315+v86*int32(28), v319, v277, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L16
	} else {
		goto L112
	}
L109:
	;
	v314 = int32(1)
	goto L111
L110:
	;
	v314 = int32(-1)
	goto L111
L111:
	;
	v334 = v314
	goto L98
L112:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
	if v323&int32(1) == int32(0) {
		v334 = v321
		goto L98
	} else {
		goto L113
	}
L113:
	;
	v329 = int32(0)
	if v321 < v329 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v333 = int32(1)
	goto L116
L115:
	;
	v333 = v329 - v321
	goto L116
L116:
	;
	v334 = v333
	goto L98
L117:
	;
	v477 = v373
	v480 = int32(1)
	v481 = int32(0)
	v483 = v375
	goto L33
L118:
	;
	v472 = int32(1)
	v477 = v373
	v480 = v472
	v481 = v472
	v483 = v375
	goto L33
L119:
	;
	if l6 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L120:
	;
	if v135 == int32(0) {
		v477 = v373
		v480 = v93
		v481 = v376
		v483 = v375
		goto L33
	} else {
		goto L126
	}
L121:
	;
	if v346 != 0 {
		goto L119
	} else {
		goto L125
	}
L122:
	;
	if v141 == int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v351 = int32(0)
	v352 = int32(1)
	v362 = base.B2i32(v33 == v352)&base.B2i32(v351 < v346) | base.B2i32(v33 == int32(-1))&base.B2i32(v346 < v351)
	if v346 == v351 {
		v373 = v362
		v374 = v351
		v375 = v100
		v376 = v352
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v368 = int32(0)
	v373 = v362
	v374 = int32(base.Ui32(v346) >> (uint(int32(31)) % 32))
	v375 = v368
	v376 = v368
	goto L120
L125:
	;
	v370 = int32(0)
	v373 = v370
	v374 = v370
	v375 = v100
	v376 = int32(1)
	goto L120
L126:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v379 == int32(-1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+18)))
	if v376 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	if v449 == v345 {
		v477 = v373
		v480 = v93
		v481 = v376
		v483 = v375
		goto L33
	} else {
		goto L155
	}
L130:
	;
	if v382&int32(1) != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	if v344 != 0 {
		goto L143
	} else {
		goto L144
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+44)) = int32(0)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v397 = v395 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v397
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+19)))
	if v399 != int32(1) {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	if v387 == int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	F_pfree(m, v387)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L16
	} else {
		goto L136
	}
L136:
	;
	goto L133
L137:
	;
	if v374 != 0 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	if v374 == base.B2i32(v395&int32(33554432) == int32(0)) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v397 | int32(65)
	goto L117
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v397 | int32(524288)
	goto L117
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v397 | int32(1048576)
	goto L117
L143:
	;
	if v382&int32(1) != 0 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L145
L145:
	;
	if v382&int32(1) != 0 {
		goto L150
	} else {
		goto L151
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+44)) = int32(0)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v426&int32(-7864386) | int32(65)
	goto L118
L147:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	if v418 == int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	F_pfree(m, v418)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L16
	} else {
		goto L149
	}
L149:
	;
	goto L146
L150:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v440 & int32(-7864386)
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+18)))
	v445 = int32(*(*int16)(unsafe.Add(mBase, uint32(v135)+16)))
	v446 = F_datumCopy(m, v277, v444, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L16
	} else {
		goto L154
	}
L151:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	if v434 == int32(0) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	F_pfree(m, v434)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L16
	} else {
		goto L153
	}
L153:
	;
	goto L150
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+44)) = v446
	goto L118
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v345
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v452+v345<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+44)) = v456
	v477 = v373
	v480 = v93
	v481 = v376
	v483 = v375
	goto L33
L156:
	;
	v1135 = int32(0)
	goto L1
L157:
	;
	goto L158
L158:
	;
	v461 = int32(1)
	if v93&v461 == int32(0) {
		v865 = v461
		v877 = v137
		goto L6
	} else {
		goto L159
	}
L159:
	;
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
	v468 = v466 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)) = uint16(v468)
	v865 = v461
	v877 = v137
	goto L6
L160:
	;
	goto L22
L161:
	;
	v494 = int32(0)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+12))
	v498 = v496 - int32(1)
	if v494 <= v498 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v503 = base.B2i32(v33 == int32(1))
	v512 = v498
	v519 = v480
	goto L165
L163:
	;
	goto L164
L164:
	;
	F__bt_start_array_keys(m, l0, int32(0)-v33)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L16
	} else {
		goto L252
	}
L165:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v495)+8))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v495)+20))
	v534 = v531 + v512<<(uint(int32(5))%32)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v534)))
	v538 = v530 + v535*int32(48)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if v539 == int32(-1) {
		goto L171
	} else {
		goto L172
	}
L166:
	;
	goto L164
L167:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if v732 != int32(-1) {
		goto L235
	} else {
		goto L236
	}
L168:
	;
	v648 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)) = uint8(v648)
	v650 = int32(1)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if v651&int32(524288) != 0 {
		v731 = v650
		goto L167
	} else {
		goto L205
	}
L169:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	if v636 <= int32(0) {
		v731 = v519
		goto L167
	} else {
		goto L204
	}
L170:
	;
	if v547&int32(1) != 0 {
		goto L180
	} else {
		goto L181
	}
L171:
	;
	if v503 == int32(0) {
		goto L168
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	if v33 != int32(1) {
		goto L169
	} else {
		goto L178
	}
L174:
	;
	v544 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)) = uint8(v544)
	v546 = int32(1)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if v547&int32(1048576) != 0 {
		v731 = v546
		goto L167
	} else {
		goto L175
	}
L175:
	;
	if v547&int32(33554433) == int32(1) {
		v731 = v546
		goto L167
	} else {
		goto L176
	}
L176:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	if v554 != 0 {
		goto L170
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v547 | int32(2097152)
	v834 = v546
	goto L7
L178:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	if v539-int32(1) <= v560 {
		v731 = v519
		goto L167
	} else {
		goto L179
	}
L179:
	;
	v565 = v560 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v534)+12)) = v565
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v567+v565<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = v571
	v834 = v519
	goto L7
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v547 & int32(-1048642)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+18)))
	v581 = int32(*(*int16)(unsafe.Add(mBase, uint32(v534)+16)))
	v582 = F_datumCopy(m, v579, v580, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L16
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v538)+44))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v554)+12))
	v589 = m.T0[v588].(func(*base.Module, int32, int32, int32) int32)(m, v501, v585, v29+int32(8))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L16
	} else {
		goto L184
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = v582
	v834 = v546
	goto L7
L184:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
	if v591 == int32(1) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+19)))
	if v594 != int32(1) {
		v731 = v546
		goto L167
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v534)+28))
	if v616 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L188:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if v597&int32(33554432) != 0 {
		v731 = v546
		goto L167
	} else {
		goto L189
	}
L189:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+18)))
	if v600 != 0 {
		v608 = v597
		goto L190
	} else {
		goto L191
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v608&int32(-7864386) | int32(65)
	v834 = v546
	goto L7
L191:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v538)+44))
	if v601 == int32(0) {
		v608 = v597
		goto L190
	} else {
		goto L192
	}
L192:
	;
	F_pfree(m, v601)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L16
	} else {
		goto L193
	}
L193:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	v608 = v606
	goto L190
L194:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+18)))
	if v628 != 0 {
		goto L200
	} else {
		goto L201
	}
L195:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v616)+12))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v616)+44))
	v623 = F_FunctionCall2Coll(m, v616+int32(16), v621, v589, v622)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L16
	} else {
		goto L196
	}
L196:
	;
	if v623 != 0 {
		goto L194
	} else {
		goto L197
	}
L197:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+18)))
	if v625 != 0 {
		v731 = v546
		goto L167
	} else {
		goto L198
	}
L198:
	;
	F_pfree(m, v589)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L16
	} else {
		goto L199
	}
L199:
	;
	v731 = v546
	goto L167
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = v589
	v834 = v546
	goto L7
L201:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v538)+44))
	if v629 == int32(0) {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	F_pfree(m, v629)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L16
	} else {
		goto L203
	}
L203:
	;
	goto L200
L204:
	;
	v640 = v636 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v534)+12)) = v640
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v642+v640<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = v646
	v834 = v519
	goto L7
L205:
	;
	v654 = int32(33554433)
	if v651&v654 == v654 {
		v731 = v650
		goto L167
	} else {
		goto L206
	}
L206:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	if v658 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v651 | int32(4194304)
	v834 = v650
	goto L7
L208:
	;
	goto L209
L209:
	;
	if v651&int32(1) != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v651 & int32(-524354)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+18)))
	v672 = int32(*(*int16)(unsafe.Add(mBase, uint32(v534)+16)))
	v673 = F_datumCopy(m, v670, v671, v672)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L16
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v538)+44))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v658)+8))
	v680 = m.T0[v679].(func(*base.Module, int32, int32, int32) int32)(m, v501, v676, v29+int32(8))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L16
	} else {
		goto L214
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = v673
	v834 = v650
	goto L7
L214:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
	if v682 == int32(1) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+19)))
	if v685 != int32(1) {
		v731 = v650
		goto L167
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v534)+24))
	if v709 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L218:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if v688&int32(33554432) == int32(0) {
		v731 = v650
		goto L167
	} else {
		goto L219
	}
L219:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+18)))
	if v693 != 0 {
		v701 = v688
		goto L220
	} else {
		goto L221
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v701&int32(-7864386) | int32(65)
	v834 = v650
	goto L7
L221:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v538)+44))
	if v694 == int32(0) {
		v701 = v688
		goto L220
	} else {
		goto L222
	}
L222:
	;
	F_pfree(m, v694)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L16
	} else {
		goto L223
	}
L223:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	v701 = v699
	goto L220
L224:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+18)))
	if v721 != 0 {
		goto L230
	} else {
		goto L231
	}
L225:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v709)+12))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v709)+44))
	v716 = F_FunctionCall2Coll(m, v709+int32(16), v714, v680, v715)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L16
	} else {
		goto L226
	}
L226:
	;
	if v716 != 0 {
		goto L224
	} else {
		goto L227
	}
L227:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+18)))
	if v718 != 0 {
		v731 = v650
		goto L167
	} else {
		goto L228
	}
L228:
	;
	F_pfree(m, v680)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L16
	} else {
		goto L229
	}
L229:
	;
	v731 = v650
	goto L167
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = v680
	v834 = v650
	goto L7
L231:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v538)+44))
	if v722 == int32(0) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	F_pfree(m, v722)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L16
	} else {
		goto L233
	}
L233:
	;
	goto L230
L234:
	;
	if int32(0) < v512 {
		v512 = v512 - int32(1)
		v519 = v731
		goto L165
	} else {
		goto L251
	}
L235:
	;
	if v33 == int32(1) {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	goto L237
L237:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+18)))
	if v746 != 0 {
		goto L241
	} else {
		goto L242
	}
L238:
	;
	v738 = int32(0)
	goto L240
L239:
	;
	v738 = v732 - int32(1)
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534)+12)) = v738
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v740+v738<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = v744
	goto L234
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = int32(0)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	v757 = v755 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v757
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+19)))
	if v759 != int32(1) {
		goto L245
	} else {
		goto L246
	}
L242:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v538)+44))
	if v747 == int32(0) {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	F_pfree(m, v747)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L16
	} else {
		goto L244
	}
L244:
	;
	goto L241
L245:
	;
	if v33 == int32(1) {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	if v503 == base.B2i32(v755&int32(33554432) == int32(0)) {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v757 | int32(65)
	goto L234
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v757 | int32(524288)
	goto L234
L249:
	;
	goto L250
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v757 | int32(1048576)
	goto L234
L251:
	;
	goto L166
L252:
	;
	v814 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v814)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+17)) = uint8(v814)
	v1135 = v494
	goto L1
L253:
	;
	v845 = int32(1)
	if v834&v845 != 0 {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	goto L255
L255:
	;
	if v483 != 0 {
		v865 = v481
		v877 = v137
		goto L6
	} else {
		goto L261
	}
L256:
	;
	v848 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
	v849 = int32(1)
	v850 = v848 + v849
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)) = uint16(v850)
	if v481&v849 != 0 {
		v865 = v845
		v877 = v137
		goto L6
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	if v481&int32(1) != 0 {
		v865 = v845
		v877 = v137
		goto L6
	} else {
		goto L260
	}
L259:
	;
	v936 = v137
	goto L5
L260:
	;
	v936 = v137
	goto L5
L261:
	;
	v1135 = int32(0)
	goto L1
L262:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	if v903 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L263:
	;
	if v893 == int32(0) {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+18)))
	if v897 != 0 {
		goto L262
	} else {
		goto L265
	}
L265:
	;
	v898 = int32(1)
	if l1 == int32(0) {
		v1135 = v898
		goto L1
	} else {
		goto L266
	}
L266:
	;
	v901 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v901)
	v1135 = v898
	goto L1
L267:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v908 = F__bt_advance_array_keys(m, l0, l1, l2, l3, l4, v906, int32(1))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L16
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	if l6 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1135 = int32(0)
	goto L1
L271:
	;
	v1135 = int32(0)
	goto L1
L272:
	;
	goto L273
L273:
	;
	if v865&int32(1) != 0 {
		v987 = v877
		goto L4
	} else {
		goto L274
	}
L274:
	;
	v936 = v877
	goto L5
L275:
	;
	if v942 == int32(0) {
		v987 = v936
		goto L4
	} else {
		goto L276
	}
L276:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942)+7)))
	if v946&int32(32) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v961 = int32(0)
	v965 = F__bt_tuple_before_array_skeys(m, l0, v33, v942, l4, v960, v961, v961, v32+int32(18))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L16
	} else {
		goto L281
	}
L278:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v31)+192))
	v958 = int32(*(*int16)(unsafe.Add(mBase, uint32(v957)+8)))
	v960 = v958
	goto L277
L279:
	;
	v951 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v942)+4)))
	if v951&int32(8192) != 0 {
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v960 = v951 & int32(4095)
	goto L277
L281:
	;
	if v965 != 0 {
		v1058 = v936
		goto L3
	} else {
		goto L282
	}
L282:
	;
	v987 = v936
	goto L4
L283:
	;
	if v987 == int32(0) {
		v1103 = v987
		goto L2
	} else {
		goto L284
	}
L284:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v996 == int32(0) {
		v1103 = v987
		goto L2
	} else {
		goto L285
	}
L285:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v999)+52))
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996)+7)))
	if v1001&int32(32) == int32(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1017 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v1017
	v1027 = F__bt_check_compare(m, l0, v1017-v33, v996, v1015, v1000, v1017, v1017, v29+int32(15), v29+int32(8))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L16
	} else {
		goto L290
	}
L287:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v999)+192))
	v1013 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1012)+8)))
	v1015 = v1013
	goto L286
L288:
	;
	v1006 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v996)+4)))
	if v1006&int32(8192) != 0 {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v1015 = v1006 & int32(4095)
	goto L286
L290:
	;
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	if v1029 != 0 {
		v1103 = v987
		goto L2
	} else {
		goto L291
	}
L291:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+8))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v1035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1030+v1031*int32(48))+6)))
	if v1035 == int32(3) {
		v1103 = v987
		goto L2
	} else {
		goto L292
	}
L292:
	;
	v1058 = v987
	goto L3
L293:
	;
	v1072 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v1072)
	v1075 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+17)) = uint8(v1075)
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v1077 == v1072 {
		v1135 = v1072
		goto L1
	} else {
		goto L298
	}
L294:
	;
	v1067 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+30)))
	if v1067 < int32(4) {
		goto L293
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v1070 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+18)) = uint8(v1070)
	v1103 = v1058
	goto L2
L297:
	;
	goto L296
L298:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
	F__bt_parallel_primscan_schedule(m, l0, v1080)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L16
	} else {
		goto L299
	}
L299:
	;
	v1135 = v1072
	goto L1
L300:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+19)) = uint8(v1103)
	if v33 != int32(1) {
		v1135 = v1111
		goto L1
	} else {
		goto L301
	}
L301:
	;
	v1120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v1122 = v1120 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v1122)
	v1135 = v1111
	goto L1
}
func F__bt_buildadd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v281 int64
	_ = v281
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l3
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = int32(4)
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+14)))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v32 = v30 - v31
	if v32 <= v29 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v45 = (v39&int32(8191) + int32(7)) & int32(16376)
	if base.Ui32(int32(2705)) <= base.Ui32(v45) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v35 = v29
	goto L9
L8:
	;
	v35 = v32
	goto L9
L9:
	;
	v37 = v35 - int32(4)
	goto L6
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F__bt_check_third_page(m, v48, v49, base.B2i32(v38 == int32(0)), v28, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(base.B2i32(v38 == int32(0))<<(uint(int32(3))%32)+v45) <= base.Ui32(v37) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	v332 = F_PageAddItemExtended(m, v321, v319, v327, v318&int32(65535), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L65
	}
L15:
	;
	v318 = v315
	v319 = v316
	v321 = v28
	v325 = v27
	v327 = v317
	goto L14
L16:
	;
	v315 = v273
	v316 = l2
	v317 = v45
	goto L15
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L62
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L59
	}
L19:
	;
	if v26 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui32(v60) <= base.Ui32(v37+v24) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v67 = F_smgr_bulk_get_buf(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	if base.Ui32(v26) < base.Ui32(int32(3)) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	F_PageInit(m, v67, int32(8192), int32(16))
	mBase = m.M
	goto L26
L26:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v73 = v67 + v72
	v74 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+14)) = uint16(v74)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+12)) = uint16(base.B2i32(v65 == v74))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v65
	*(*int64)(unsafe.Add(mBase, uint32(v73))) = int64(0)
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	v83 = int32(4)
	v84 = v82 + v83
	*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)) = uint16(v84)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v86 + int32(1)
	v91 = v28 + int32(24)
	v96 = v91 + v26<<(uint(int32(2))%32) - v83
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v100 = v28 + v97&int32(32767)
	if v38 == v74 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v116 = F_PageAddItemExtended(m, v67, v113, v112, int32(2), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L31
	}
L28:
	;
	v112 = int32(base.Ui32(v97) >> (uint(int32(17)) % 32))
	v113 = v100
	goto L27
L29:
	;
	goto L30
L30:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(537395200)
	v109 = int32(8)
	v112 = v109
	v113 = v18 + v109
	goto L27
L31:
	;
	if v116 == int32(0) {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v120
	v122 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v122
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v126 = v124 - int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)) = uint16(v126)
	if v38 == v122 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v130 = int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32((v26-v130)&int32(65535)<<(uint(int32(2))%32)+v91-int32(4))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v146 = F__bt_truncate(m, v131, v28+v141&int32(32767), v100, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	v162 = v100
	goto L35
L35:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v163 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+6)))
	v151 = F_PageIndexTupleOverwrite(m, v28, v130, v146, v148&int32(8191))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v151 == int32(0) {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	F_pfree(m, v146)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v162 = v28 + v157&int32(32767)
	goto L35
L40:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v168 = v166 + int32(1)
	v170 = F_palloc0(m, int32(32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = base.I32_rotr(v27, int32(16))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__bt_buildadd(m, l0, v232, v233, int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L49
	}
L43:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v173 = F_smgr_bulk_get_buf(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_PageInit(m, v173, int32(8192), int32(16))
	mBase = m.M
	goto L45
L45:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+16)))
	v179 = v173 + v178
	v180 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v179)+14)) = uint16(v180)
	*(*uint16)(unsafe.Add(mBase, uint32(v179)+12)) = uint16(base.B2i32(v168 == v180))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v168
	*(*int64)(unsafe.Add(mBase, uint32(v179))) = int64(0)
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+12)))
	v190 = v188 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+12)) = uint16(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v173
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v194 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v193 + v194
	*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v170)+16)) = v180
	*(*uint16)(unsafe.Add(mBase, uint32(v170)+12)) = uint16(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v193
	if v168 != 0 {
		v219 = int32(2457)
		goto L46
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v170
	goto L42
L47:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+180))
	if v208 == int32(0) {
		v219 = int32(819)
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v217 = base.I32_div_s(int32(819200)-v212<<(uint(int32(13))%32), int32(100))
	v219 = v217
	goto L46
L49:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_pfree(m, v237)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v240 = F_CopyIndexTuple(m, v162)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v240
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v28+v244)+4)) = v86
	v247 = v243 + v67
	*(*int32)(unsafe.Add(mBase, uint32(v247)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v27
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_smgr_bulk_write(m, v251, v27, v28, int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v318 = int32(3)
	v319 = l2
	v321 = v67
	v325 = v86
	v327 = v45
	goto L14
L53:
	;
	v259 = F_palloc0(m, int32(8))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v273 = v26 + int32(1)
	if v38 == int32(0) {
		goto L16
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v259
	v262 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v259)+6)) = uint16(v262)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v265 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v264)+4)) = uint16(v265)
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264)+6)))
	v269 = v267 | int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v264)+6)) = uint16(v269)
	goto L55
L57:
	;
	v276 = int32(2)
	if v273&int32(65535) != v276 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(537395200)
	v285 = int32(8)
	v315 = v276
	v316 = v18 + v285
	v317 = v285
	goto L15
L59:
	;
	F_errmsg_internal(m, int32(427626), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(517238), int32(735), int32(244715))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errmsg_internal(m, int32(427585), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(517238), int32(938), int32(485877))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	if v332 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v318)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v321
	m.G0 = v18 + int32(16)
	return
L67:
	;
	goto L68
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errmsg_internal(m, int32(427626), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(517238), int32(735), int32(244715))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_checkpage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 < int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l1^int32(-1))<<(uint(int32(2))%32))))
		v26 = v18
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[1]))
		v26 = v20 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+14)))
	if v27 != 0 {
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+19)))
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
		if (v28<<(uint(int32(8))%32)-v31)&int32(65535) != int32(16) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if l1 < int32(0) {
						v94 = *(*int32)(unsafe.Add(mBase, _consts[8]))
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v94+(l1^int32(-1))<<(uint(int32(6))%32))+16))
						v109 = v100
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v102+l1<<(uint(int32(6))%32)+int32(-64))+16))
						v109 = v108
					}
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v109
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v90 + int32(4)
					F_errmsg(m, int32(52001), v7+int32(16))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return
					} else {
						F_errhint(m, int32(611072), int32(0))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return
						} else {
							F_errfinish(m, int32(525205), int32(824), int32(427017))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
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
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if l1 < int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[8]))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(l1^int32(-1))<<(uint(int32(6))%32))+16))
					v66 = v57
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+l1<<(uint(int32(6))%32)+int32(-64))+16))
					v66 = v65
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v66
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v47 + int32(4)
				F_errmsg(m, int32(51948), v7)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					F_errhint(m, int32(611072), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_errfinish(m, int32(525205), int32(813), int32(427017))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
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
}
func F__bt_end_vacuum_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	v9 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v13 = F_LWLockAcquire(m, v9+int32(2560), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v17 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v63+int32(2560))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L11
	}
L4:
	;
	v21 = v16 + int32(12)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v24 = int32(0)
	goto L5
L5:
	;
	v33 = v21 + v24*int32(12)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 != v22 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v53 = v24 + int32(1)
	if v53 != v17 {
		v24 = v53
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v36 != v37 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = int32(12)
	v43 = v17*v39 + v21 - v39
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v48 - int32(1)
	goto L3
L10:
	;
	goto L6
L11:
	;
	return
}
func F__bt_find_extreme_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+208))
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14+v15<<(uint(int32(2))%32)-int32(4))))
	v23 = F_get_opfamily_member(m, v21, l2, l2, base.I32_extend16_s(l3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L23
	}
L2:
	;
	return int32(0)
L3:
	;
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = F_get_opcode(m, v23)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L20
	}
L7:
	;
	if v27 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_fmgr_info(m, v27, v11+int32(20))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if int32(2) <= l5 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = v35
	v41 = int32(1)
	goto L13
L11:
	;
	v61 = v35
	goto L12
L12:
	;
	m.G0 = v11 + int32(48)
	return v61
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v52 = l4 + v41<<(uint(int32(2))%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = F_FunctionCall2Coll(m, v11+int32(20), v49, v53, v39)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	v61 = v57
	goto L12
L15:
	;
	if v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v57 = v56
	goto L18
L17:
	;
	v57 = v39
	goto L18
L18:
	;
	v59 = v41 + int32(1)
	if v59 != l5 {
		v39 = v57
		v41 = v59
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v13)+208))
	v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77+v78<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
	F_errmsg_internal(m, int32(42587), v11)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(518228), int32(2605), int32(102811))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v23
	F_errmsg_internal(m, int32(46874), v11+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(518228), int32(2608), int32(102811))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_form_posting(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v7&int32(8192) == int32(0) {
		v24 = v7 & int32(8191)
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
		if v12&int32(32) == int32(0) {
			v24 = v7 & int32(8191)
		} else {
			v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
			v24 = v17 | v18<<(uint(int32(16))%32)
		}
	}
	v26 = l2 * int32(6)
	if int32(1) < l2 {
		v34 = (v24 + v26 + int32(7)) & int32(-8)
	} else {
		v34 = v24
	}
	v35 = F_palloc0(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		if v24 != 0 {
			v39 = F__emscripten_memcpy_bulkmem(m, v35, l0, v24)
			mBase = m.M
			v40 = v39
		} else {
			v40 = v35
		}
		v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+6)))
		v44 = v41&int32(-8192) | v34
		if int32(2) <= l2 {
			*(*uint16)(unsafe.Add(mBase, uint32(v40)+2)) = uint16(v24)
			v48 = int32(8192)
			v49 = l2 | v48
			*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)) = uint16(v49)
			v52 = v44 | v48
			*(*uint16)(unsafe.Add(mBase, uint32(v40)+6)) = uint16(v52)
			v55 = int32(base.Ui32(v24) >> (uint(int32(16)) % 32))
			*(*uint16)(unsafe.Add(mBase, uint32(v40))) = uint16(v55)
			if v26 != 0 {
				v58 = F__emscripten_memcpy_bulkmem(m, v40+v24, l1, v26)
				mBase = m.M
			} else {
			}
			return v40
		} else {
			v62 = v44 & int32(57343)
			*(*uint16)(unsafe.Add(mBase, uint32(v40)+6)) = uint16(v62)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v40))) = v64
			v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)) = uint16(v66)
			return v40
		}
	}
}
func F__bt_getmeta(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 < int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l1^int32(-1))<<(uint(int32(2))%32))))
		v26 = v18
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[1]))
		v26 = v20 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v26)+12)))
	if v29&int32(8) == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v60 + int32(4)
				F_errmsg(m, int32(430945), v7)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(525205), int32(158), int32(530512))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v35 = v26 + int32(24)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
		if v36 != int32(340322) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v60 + int32(4)
					F_errmsg(m, int32(430945), v7)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(525205), int32(158), int32(530512))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v40 = v26 + int32(28)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
			if base.Ui32(v41-int32(5)) <= base.Ui32(int32(-4)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33557032))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(8589934596)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v79
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v80 + int32(4)
						F_errmsg(m, int32(496133), v7+int32(16))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525205), int32(167), int32(530512))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				m.G0 = v7 + int32(32)
				return v35
			}
		}
	}
}
func F__bt_getstackbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v224 int32
	_ = v224
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v23 = v15
	v26 = v16
	goto L1
L1:
	;
	goto L3
L2:
	;
	return int32(0)
L3:
	;
	v46 = F__bt_getbuf(m, l0, v26, int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v70&int32(20) != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v69 = v68 + v67
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+12)))
	if v70&int32(128) != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	return int32(0)
L7:
	;
	if v46 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(v46^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L5
L9:
	;
	goto L10
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v67 = v61 + v46<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F__bt_finish_split(m, l0, l1, v46, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L4
L14:
	;
	goto L3
L15:
	;
	F__bt_relbuf(m, v46)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L43
	}
L16:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v79) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v197)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v26
	return v46
L18:
	;
	v87 = int32(base.Ui32(v79+int32(262120)) >> (uint(int32(2)) % 32))
	goto L20
L19:
	;
	v87 = int32(0)
	goto L20
L20:
	;
	v88 = int32(1)
	v91 = v23 & int32(65535)
	if v76 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v94 = int32(2)
	goto L23
L22:
	;
	v94 = v88
	goto L23
L23:
	;
	if base.Ui32(v94) < base.Ui32(v91) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v96 = v91
	goto L26
L25:
	;
	v96 = v94
	goto L26
L26:
	;
	v98 = v87 & int32(65535)
	if base.Ui32(v98) < base.Ui32(v96) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v100 = v87 + v88
	goto L29
L28:
	;
	v100 = v96
	goto L29
L29:
	;
	if base.Ui32(v100&int32(65535)) <= base.Ui32(v98) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v110 = v100
	goto L33
L31:
	;
	goto L32
L32:
	;
	v163 = v100
	goto L39
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v110&int32(65535)<<(uint(int32(2))%32)+(v67+int32(24))-int32(4))))
	v130 = v67 + v127&int32(32767)
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130))))
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+2)))
	if l3 == v131<<(uint(int32(16))%32)|v134 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v197 = v110
	goto L17
L36:
	;
	goto L37
L37:
	;
	v138 = v110 + int32(1)
	if base.Ui32(v138&int32(65535)) <= base.Ui32(v98) {
		v110 = v138
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v173 = v163 - int32(1)
	v175 = v173 & int32(65535)
	if base.Ui32(v175) < base.Ui32(v94) {
		goto L15
	} else {
		goto L41
	}
L40:
	;
	v197 = v173
	goto L17
L41:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v175<<(uint(int32(2))%32)+(v67+int32(24))-int32(4))))
	v185 = v67 + v182&int32(32767)
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185))))
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+2)))
	if v186<<(uint(int32(16))%32)|v189 != l3 {
		v163 = v173
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	if v76 != 0 {
		v23 = int32(0)
		v26 = v76
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L2
}
func F__bt_next(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
	if l1 == int32(1) {
		v10 = v6 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+100)) = v10
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+96))
		if v10 <= v12 {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
			v37 = v5 + v34*int32(10)
			v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+108)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v38)
			v41 = v37 + int32(104)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v42
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
			if v44 != 0 {
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+8)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44 + v45
			} else {
			}
			return int32(1)
		} else {
			v15 = F__bt_steppage(m, l0, int32(1))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v15 != 0 {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
					v37 = v5 + v34*int32(10)
					v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+108)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v38)
					v41 = v37 + int32(104)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v42
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
					if v44 != 0 {
						v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+8)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44 + v45
					} else {
					}
					return int32(1)
				} else {
					return int32(0)
				}
			}
		}
	} else {
		v22 = v6 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+100)) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v5)+92))
		if v24 <= v22 {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
			v37 = v5 + v34*int32(10)
			v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+108)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v38)
			v41 = v37 + int32(104)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v42
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
			if v44 != 0 {
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+8)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44 + v45
			} else {
			}
			return int32(1)
		} else {
			v26 = F__bt_steppage(m, l0, l1)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 != 0 {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
					v37 = v5 + v34*int32(10)
					v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+108)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v38)
					v41 = v37 + int32(104)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v42
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
					if v44 != 0 {
						v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+8)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44 + v45
					} else {
					}
					return int32(1)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F__bt_readfirstpage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+17)))
	if v12 == int32(1) {
		v15 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+17)) = uint8(v15)
		v17 = int32(257)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+88)) = uint16(v17)
	} else {
		if l2 == int32(1) {
			v21 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(v7)+88)) = uint16(v21)
		} else {
			v23 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v7)+88)) = uint16(v23)
		}
	}
	v26 = F__bt_readpage(m, l0, l2, l1, int32(1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		if v26 != 0 {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+40)))
			if v32 == int32(0) {
				F__bt_unlockbuf(m, v31)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return int32(1)
				}
			} else {
				v39 = F_BufferGetLSNAtomic(m, v31)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v7)+72)) = v39
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
					F__bt_relbuf(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = int32(0)
						return int32(1)
					}
				}
			}
		} else {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
			F__bt_unlockbuf(m, v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v52 = F__bt_steppage(m, l0, l2)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					return v52
				}
			}
		}
	}
}
func F__bt_set_startikey(m *base.Module, l0 int32, l1 int32) {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
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
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v27 = v25 + int32(24)
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v29 = int32(2)
	v32 = int32(4)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(v29)%32)-v32)))
	v35 = int32(32767)
	v37 = v25 + v34&v35
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38<<(uint(v29)%32)+v27-v32)))
	v47 = v25 + v44&v35
	v48 = int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+10)))
	if v50 <= int32(0) {
		v99 = v48
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v113 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v57 = v48
	goto L5
L5:
	;
	v73 = F_index_getattr_2(m, v47, v57, v24, v17+int32(15))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v99 = v50 + int32(1)
	goto L3
L7:
	;
	return
L8:
	;
	v77 = F_index_getattr_2(m, v37, v57, v24, v17+int32(14))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)))
	if v79 != v80 {
		v99 = v57
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v79 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v88 = v57<<(uint(int32(4))%32) + (v24 + int32(20)) - int32(16)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+6)))
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
	v91 = F_datum_image_eq(m, v73, v77, v89, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v57 != v50 {
		v57 = v57 + int32(1)
		goto L5
	} else {
		goto L16
	}
L14:
	;
	if v91 == int32(0) {
		v99 = v57
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	goto L6
L17:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v388 != 0 {
		goto L1
	} else {
		goto L102
	}
L18:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v357
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v364)
	if v364 != int32(1) {
		goto L1
	} else {
		goto L101
	}
L19:
	;
	v357 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v117 = int32(0)
	v124 = v113
	v126 = v117
	v127 = v117
	v129 = v117
	goto L23
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v340
	v344 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v344)
	v375 = l1 + int32(13)
	v387 = l1 + int32(16)
	goto L17
L23:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v137 = v134 + v127*int32(48)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if v138&int32(196608) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	if v323&int32(1) == int32(0) {
		v357 = v330
		goto L18
	} else {
		goto L100
	}
L25:
	;
	v330 = v127 + int32(1)
	if v330 < v325 {
		v124 = v325
		v126 = v326
		v127 = v330
		v129 = v327
		goto L23
	} else {
		goto L99
	}
L26:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v323 = v321
	v325 = v322
	v326 = v319
	v327 = v321
	goto L25
L27:
	;
	v319 = v126
	v321 = v129
	goto L26
L28:
	;
	v319 = v206
	v321 = v129
	goto L26
L29:
	;
	if v129&int32(1) != 0 {
		v340 = v127
		goto L22
	} else {
		goto L98
	}
L30:
	;
	if v138&int32(4) != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+6)))
	if v145 != int32(3) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+19)))
	if v241 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L33:
	;
	if v129&int32(1) != 0 {
		v340 = v127
		goto L22
	} else {
		goto L68
	}
L34:
	;
	v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+4)))
	if v99 < v148 {
		goto L29
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v138&int32(32) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L37:
	;
	v152 = F_index_getattr_2(m, v47, v148, v24, v17+int32(13))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+4)))
	v157 = F_index_getattr_2(m, v37, v154, v24, v17+int32(12))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+13)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v160&int32(1) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v159&int32(1) != 0 {
		goto L29
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v159&int32(1) != 0 {
		goto L29
	} else {
		goto L45
	}
L43:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	if v165 != 0 {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L27
L45:
	;
	v169 = v137 + int32(16)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	v172 = F_FunctionCall2Coll(m, v169, v170, v152, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	if v172 == int32(0) {
		goto L29
	} else {
		goto L47
	}
L47:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	if v176 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	v179 = F_FunctionCall2Coll(m, v169, v177, v157, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	if v179 != 0 {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	goto L33
L51:
	;
	v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+4)))
	if v99 <= v185 {
		goto L29
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v206 = v126 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v210 = v207 + v126<<(uint(int32(5))%32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v211 == int32(-1) {
		goto L32
	} else {
		goto L63
	}
L54:
	;
	v189 = F_index_getattr_2(m, v47, v185, v24, v17+int32(13))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+13)))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v192&int32(1) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v191&int32(1) != 0 {
		goto L27
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v191&int32(1) != 0 {
		goto L29
	} else {
		goto L60
	}
L59:
	;
	goto L33
L60:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	v203 = F_FunctionCall2Coll(m, v137+int32(16), v201, v189, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	if v203 != 0 {
		goto L27
	} else {
		goto L62
	}
L62:
	;
	goto L33
L63:
	;
	v214 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+4)))
	if v99 <= v214 {
		goto L29
	} else {
		goto L64
	}
L64:
	;
	v218 = F_index_getattr_2(m, v47, v214, v24, v17+int32(13))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v224 = int32(0)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+13)))
	v229 = F__bt_binsrch_array_skey(m, v220+v127*int32(28), v224, v224, v218, v226, v210, v137, v17+int32(8))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v232 == int32(0) {
		v319 = v206
		v321 = int32(1)
		goto L26
	} else {
		goto L67
	}
L67:
	;
	goto L33
L68:
	;
	v357 = v127
	goto L18
L69:
	;
	v323 = v129
	v325 = v124
	v326 = v206
	v327 = v129
	goto L25
L70:
	;
	goto L71
L71:
	;
	v244 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+4)))
	if v99 < v244 {
		goto L29
	} else {
		goto L72
	}
L72:
	;
	v248 = F_index_getattr_2(m, v47, v244, v24, v17+int32(13))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	v250 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+4)))
	v253 = F_index_getattr_2(m, v37, v250, v24, v17+int32(12))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+19)))
	if v255 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+19)))
	if v278 != 0 {
		goto L28
	} else {
		goto L86
	}
L76:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+13)))
	if v256 != 0 {
		goto L29
	} else {
		goto L77
	}
L77:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v210)+24))
	if v257 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)+44))
	v262 = F_FunctionCall2Coll(m, v257+int32(16), v260, v248, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L7
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	if v266 == int32(0) {
		goto L75
	} else {
		goto L83
	}
L81:
	;
	if v262 == int32(0) {
		goto L29
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266)+44))
	v273 = F_FunctionCall2Coll(m, v266+int32(16), v271, v248, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	if v273 == int32(0) {
		goto L29
	} else {
		goto L85
	}
L85:
	;
	goto L75
L86:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	if v279 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	if v129&int32(1) == int32(0) {
		v357 = v127
		goto L18
	} else {
		goto L97
	}
L88:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v210)+24))
	if v280 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v280)+44))
	v285 = F_FunctionCall2Coll(m, v280+int32(16), v283, v253, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L7
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	if v289 == int32(0) {
		goto L28
	} else {
		goto L94
	}
L92:
	;
	if v285 == int32(0) {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v289)+44))
	v296 = F_FunctionCall2Coll(m, v289+int32(16), v294, v253, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	if v296 != 0 {
		goto L28
	} else {
		goto L96
	}
L96:
	;
	goto L87
L97:
	;
	v340 = v127
	goto L22
L98:
	;
	v357 = v127
	goto L18
L99:
	;
	goto L24
L100:
	;
	v340 = v330
	goto L22
L101:
	;
	v375 = l1 + int32(13)
	v387 = l1 + int32(16)
	goto L17
L102:
	;
	v389 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v375))) = uint8(v389)
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v389
	goto L1
}
func F__bt_start_array_keys(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v3 < v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	v91 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+18)) = uint16(v91)
	return
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v23 = v20 + v17<<(uint(int32(5))%32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v27 = v19 + v24*int32(48)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v28 != int32(-1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v81 = v17 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v81 < v82 {
		v17 = v81
		goto L4
	} else {
		goto L24
	}
L7:
	;
	v31 = int32(1)
	if l1 != v31 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+18)))
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v36 = v28 - v31
	goto L12
L11:
	;
	v36 = int32(0)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v36<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v42
	goto L6
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v55 = v53 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v55
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+19)))
	if v57 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v45 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	goto L13
L18:
	;
	if l1 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if base.B2i32(v53&int32(33554432) == int32(0)) == base.B2i32(l1 == int32(1)) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v55 | int32(65)
	goto L6
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v55 | int32(524288)
	goto L6
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v55 | int32(1048576)
	goto L6
L24:
	;
	goto L5
}
func F_bt_page_items(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_bt_page_items_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_bt_report_duplicate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v117 int64
	_ = v117
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	v13 = m.G0
	v15 = v13 - int32(144)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+2)))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17))))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v18 | v19<<(uint(int32(16))%32)
	v26 = int32(703531)
	v29 = F_psprintf(m, v26, v15+int32(128))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return
	} else {
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
		v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
		v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v33
		*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v31 | v32<<(uint(int32(16))%32)
		v41 = F_psprintf(m, v26, v15+int32(112))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v44
			*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v43
			v49 = F_psprintf(m, v26, v15+int32(96))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if l3 != v51 {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l3
					v64 = F_psprintf(m, int32(703530), v15+int32(80))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v67 = v64
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v69 < int32(0) {
							v80 = int32(793540)
							if l5 < int32(0) {
								v92 = int32(793540)
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									F_errcode(m, int32(33557032))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
										F_errmsg(m, int32(726753), v15+int32(32))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return
										} else {
											v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
											*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
											v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
											*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
											*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
											*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
											*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
											F_errdetail(m, int32(688016), v15)
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return
											} else {
												F_errfinish(m, int32(525315), int32(906), int32(376378))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
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
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l5
								v90 = F_psprintf(m, int32(53017), v15+int32(48))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									v92 = v90
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
											F_errmsg(m, int32(726753), v15+int32(32))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
												v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
												*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
												*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
												*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
												F_errdetail(m, int32(688016), v15)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													F_errfinish(m, int32(525315), int32(906), int32(376378))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
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
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v69
							v78 = F_psprintf(m, int32(53017), v15-int32(-64))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								v80 = v78
								if l5 < int32(0) {
									v92 = int32(793540)
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
											F_errmsg(m, int32(726753), v15+int32(32))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
												v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
												*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
												*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
												*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
												F_errdetail(m, int32(688016), v15)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													F_errfinish(m, int32(525315), int32(906), int32(376378))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
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
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l5
									v90 = F_psprintf(m, int32(53017), v15+int32(48))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										v92 = v90
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											F_errcode(m, int32(33557032))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
												F_errmsg(m, int32(726753), v15+int32(32))
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return
												} else {
													v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
													v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
													*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
													*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
													*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
													F_errdetail(m, int32(688016), v15)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														F_errfinish(m, int32(525315), int32(906), int32(376378))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
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
								}
							}
						}
					}
				} else {
					v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					if l4 != v54 {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l3
						v64 = F_psprintf(m, int32(703530), v15+int32(80))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v67 = v64
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v69 < int32(0) {
								v80 = int32(793540)
								if l5 < int32(0) {
									v92 = int32(793540)
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
											F_errmsg(m, int32(726753), v15+int32(32))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
												v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
												*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
												*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
												*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
												F_errdetail(m, int32(688016), v15)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													F_errfinish(m, int32(525315), int32(906), int32(376378))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
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
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l5
									v90 = F_psprintf(m, int32(53017), v15+int32(48))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										v92 = v90
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											F_errcode(m, int32(33557032))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
												F_errmsg(m, int32(726753), v15+int32(32))
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return
												} else {
													v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
													v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
													*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
													*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
													*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
													F_errdetail(m, int32(688016), v15)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														F_errfinish(m, int32(525315), int32(906), int32(376378))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
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
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v69
								v78 = F_psprintf(m, int32(53017), v15-int32(-64))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									v80 = v78
									if l5 < int32(0) {
										v92 = int32(793540)
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											F_errcode(m, int32(33557032))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
												F_errmsg(m, int32(726753), v15+int32(32))
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return
												} else {
													v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
													v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
													*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
													*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
													*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
													F_errdetail(m, int32(688016), v15)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														F_errfinish(m, int32(525315), int32(906), int32(376378))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
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
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l5
										v90 = F_psprintf(m, int32(53017), v15+int32(48))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											v92 = v90
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												F_errcode(m, int32(33557032))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
													F_errmsg(m, int32(726753), v15+int32(32))
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return
													} else {
														v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
														*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
														*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
														v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
														*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
														*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
														*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
														F_errdetail(m, int32(688016), v15)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															F_errfinish(m, int32(525315), int32(906), int32(376378))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
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
									}
								}
							}
						}
					} else {
						v67 = int32(793540)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v69 < int32(0) {
							v80 = int32(793540)
							if l5 < int32(0) {
								v92 = int32(793540)
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									F_errcode(m, int32(33557032))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
										F_errmsg(m, int32(726753), v15+int32(32))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return
										} else {
											v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
											*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
											v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
											*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
											*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
											*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
											*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
											F_errdetail(m, int32(688016), v15)
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return
											} else {
												F_errfinish(m, int32(525315), int32(906), int32(376378))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
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
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l5
								v90 = F_psprintf(m, int32(53017), v15+int32(48))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									v92 = v90
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
											F_errmsg(m, int32(726753), v15+int32(32))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
												v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
												*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
												*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
												*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
												F_errdetail(m, int32(688016), v15)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													F_errfinish(m, int32(525315), int32(906), int32(376378))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
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
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v69
							v78 = F_psprintf(m, int32(53017), v15-int32(-64))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								v80 = v78
								if l5 < int32(0) {
									v92 = int32(793540)
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
											F_errmsg(m, int32(726753), v15+int32(32))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
												v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
												*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
												*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
												*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
												F_errdetail(m, int32(688016), v15)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													F_errfinish(m, int32(525315), int32(906), int32(376378))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
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
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l5
									v90 = F_psprintf(m, int32(53017), v15+int32(48))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										v92 = v90
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											F_errcode(m, int32(33557032))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v101 + int32(4)
												F_errmsg(m, int32(726753), v15+int32(32))
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return
												} else {
													v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*uint32)(unsafe.Add(mBase, uint32(v15)+28)) = uint32(v112)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v41
													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
													v117 = int64(base.Ui64(v112) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v117)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v92
													*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v67
													*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
													*(*int32)(unsafe.Add(mBase, uint32(v15))) = v49
													F_errdetail(m, int32(688016), v15)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														F_errfinish(m, int32(525315), int32(906), int32(376378))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
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
								}
							}
						}
					}
				}
			}
		}
	}
}
