package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyReadLine(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
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
	var v81 int32
	_ = v81
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int64
	_ = v638
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v665 int32
	_ = v665
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v766 int32
	_ = v766
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v791 int32
	_ = v791
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v843 int32
	_ = v843
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int64
	_ = v923
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1041 int32
	_ = v1041
	v3 = int32(0)
	v20 = l0 + int32(288)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v3
	goto L1
L1:
	;
	v28 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+304)) = uint8(v28)
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v31 != v34 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v41 = v3
	v42 = v3
	goto L4
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v48 = v43
	v49 = v45
	v50 = v45
	v52 = v3
	v56 = v3
	v58 = v3
	v63 = v3
	goto L17
L5:
	;
	v36 = v31
	goto L7
L6:
	;
	v36 = int32(0)
	goto L7
L7:
	;
	v41 = base.I32_extend8_s(v34)
	v42 = base.I32_extend8_s(v36)
	goto L4
L8:
	;
	v1041 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+304)) = uint8(v1041)
	return v1025
L9:
	;
	if v582 < v986 {
		goto L292
	} else {
		goto L293
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v975
	v986 = v979
	goto L9
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L27
	} else {
		goto L288
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L27
	} else {
		goto L284
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L27
	} else {
		goto L280
	}
L14:
	;
	v895 = int32(1)
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v896 != v895 {
		v1025 = v895
		goto L8
	} else {
		goto L275
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v874
	goto L14
L16:
	;
	if v599 <= v582 {
		goto L14
	} else {
		goto L273
	}
L17:
	;
	if v52&int32(1) != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800+v44))))
	switch v811 - int32(10) {
	case 0, 3:
		goto L258
	default:
		goto L259
	}
L19:
	;
	v597 = int32(1)
	v599 = v584 + v597
	v601 = int32(*(*int8)(unsafe.Add(mBase, uint32(v584+v44))))
	if l1 == int32(0) {
		v645 = v58
		v646 = v597
		v647 = v63
		goto L182
	} else {
		goto L183
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v50 <= v49 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if v48 <= v50 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v581 = v48
	v582 = v49
	v584 = v50
	v588 = v56
	goto L19
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v80 == v81 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v78 = v48
	v79 = v49
	v80 = v67
	goto L23
L25:
	;
	goto L26
L26:
	;
	F_appendBinaryStringInfo(m, v20, v49+v67, v50-v49)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v50
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v78 = v77
	v79 = v50
	v80 = v76
	goto L23
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v79
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L32
L32:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v103 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v525 = v523 - v524
	if v78-v79 < v525 {
		goto L166
	} else {
		goto L167
	}
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	if v106 == v107 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v257 == v258 {
		goto L87
	} else {
		goto L88
	}
L38:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v109 != int32(1) {
		goto L34
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v115 = v114 + v107
	v116 = v106 - v107
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v117) <= base.Ui32(int32(41)) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)) = uint8(v112)
	goto L34
L42:
	;
	if v234 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L43:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v117*int32(28))+uint32(_consts[354])))
	v125 = m.T0[v124].(func(*base.Module, int32, int32) int32)(m, v115, v116)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L27
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v127 = int32(0)
	v131 = base.B2i32(v116 != v127)
	if v115&int32(3) == v127 {
		v157 = v115
		v159 = v116
		v160 = v131
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v234 = v125
	goto L42
L47:
	;
	if v230 != 0 {
		goto L73
	} else {
		goto L74
	}
L48:
	;
	v230 = int32(0)
	goto L47
L49:
	;
	v208 = v201
	v210 = v203
	goto L67
L50:
	;
	if v160 == int32(0) {
		goto L48
	} else {
		goto L58
	}
L51:
	;
	if v116 == int32(0) {
		v157 = v115
		v159 = v116
		v160 = v131
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v140 = v115
	v142 = v116
	goto L53
L53:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v145 == int32(0) {
		v201 = v140
		v203 = v142
		goto L49
	} else {
		goto L55
	}
L54:
	;
	v157 = v152
	v159 = v148
	v160 = v150
	goto L50
L55:
	;
	v147 = int32(1)
	v148 = v142 - v147
	v149 = int32(0)
	v150 = base.B2i32(v148 != v149)
	v152 = v140 + v147
	if v152&int32(3) == v149 {
		v157 = v152
		v159 = v148
		v160 = v150
		goto L50
	} else {
		goto L56
	}
L56:
	;
	if v148 != 0 {
		v140 = v152
		v142 = v148
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v164 == int32(0) {
		v194 = v157
		v196 = v159
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if v196 == int32(0) {
		goto L48
	} else {
		goto L66
	}
L60:
	;
	if base.Ui32(v159) < base.Ui32(int32(4)) {
		v194 = v157
		v196 = v159
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v174 = v157
	v176 = v159
	goto L62
L62:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v181 = v180 ^ int32(0)
	v184 = int32(-2139062144)
	if (int32(16843008)-v181|v181)&v184 != v184 {
		v201 = v174
		v203 = v176
		goto L49
	} else {
		goto L64
	}
L63:
	;
	v194 = v189
	v196 = v191
	goto L59
L64:
	;
	v188 = int32(4)
	v189 = v174 + v188
	v191 = v176 - v188
	if base.Ui32(int32(3)) < base.Ui32(v191) {
		v174 = v189
		v176 = v191
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v201 = v194
	v203 = v196
	goto L49
L67:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if int32(0) == v213 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L48
L69:
	;
	v230 = v208
	goto L47
L70:
	;
	goto L71
L71:
	;
	v215 = int32(1)
	v218 = v210 - v215
	if v218 != 0 {
		v208 = v208 + v215
		v210 = v218
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	v232 = v230 - v115
	goto L75
L74:
	;
	v232 = v116
	goto L75
L75:
	;
	v234 = v232
	goto L42
L76:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v237 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v254 + v234
	goto L34
L79:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v240) <= base.Ui32(int32(41)) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	goto L81
L81:
	;
	v252 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+321)) = uint8(v252)
	goto L34
L82:
	;
	if v116 < v250 {
		goto L34
	} else {
		goto L86
	}
L83:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v240*int32(28))+uint32(_consts[355])))
	v250 = v249
	goto L85
L84:
	;
	v250 = int32(1)
	goto L85
L85:
	;
	goto L82
L86:
	;
	goto L81
L87:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v260 != int32(1) {
		goto L34
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v267 = v265 - v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v266 <= int32(0) {
		v419 = v268
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v263 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)) = uint8(v263)
	goto L34
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v267
	v421 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v421
	*(*uint8)(unsafe.Add(mBase, uint32(v267+v419))) = uint8(v421)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v434 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	goto L140
L92:
	;
	if v267 <= int32(0) {
		v419 = v268
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v273 = v268 + v266
	if v268 == v273 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v419 = v418
	goto L91
L95:
	;
	goto L94
L96:
	;
	v277 = v268 + v267
	if base.Ui32(v273-v277) <= base.Ui32(int32(0)-v267<<(uint(int32(1))%32)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v284 = F___memcpy(m, v268, v273, v267)
	mBase = m.M
	goto L94
L98:
	;
	goto L99
L99:
	;
	v287 = (v268 ^ v273) & int32(3)
	if base.Ui32(v268) < base.Ui32(v273) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	if v389 == int32(0) {
		goto L95
	} else {
		goto L136
	}
L101:
	;
	if base.Ui32(v367) <= base.Ui32(int32(3)) {
		v388 = v366
		v389 = v367
		v390 = v368
		goto L100
	} else {
		goto L132
	}
L102:
	;
	if v287 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	if v287 != 0 {
		v349 = v267
		goto L115
	} else {
		goto L116
	}
L105:
	;
	v388 = v273
	v389 = v267
	v390 = v268
	goto L100
L106:
	;
	goto L107
L107:
	;
	if v268&int32(3) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v366 = v273
	v367 = v267
	v368 = v268
	goto L101
L109:
	;
	goto L110
L110:
	;
	v294 = v273
	v295 = v267
	v296 = v268
	goto L111
L111:
	;
	if v295 == int32(0) {
		goto L95
	} else {
		goto L113
	}
L112:
	;
	v366 = v303
	v367 = v305
	v368 = v307
	goto L101
L113:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	*(*uint8)(unsafe.Add(mBase, uint32(v296))) = uint8(v300)
	v302 = int32(1)
	v303 = v294 + v302
	v305 = v295 - v302
	v307 = v296 + v302
	if v307&int32(3) != 0 {
		v294 = v303
		v295 = v305
		v296 = v307
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	if v349 == int32(0) {
		goto L95
	} else {
		goto L128
	}
L116:
	;
	if v277&int32(3) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v314 = v267
	goto L120
L118:
	;
	v329 = v267
	goto L119
L119:
	;
	if base.Ui32(v329) <= base.Ui32(int32(3)) {
		v349 = v329
		goto L115
	} else {
		goto L124
	}
L120:
	;
	if v314 == int32(0) {
		goto L95
	} else {
		goto L122
	}
L121:
	;
	v329 = v320
	goto L119
L122:
	;
	v320 = v314 - int32(1)
	v321 = v268 + v320
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273+v320))))
	*(*uint8)(unsafe.Add(mBase, uint32(v321))) = uint8(v323)
	if v321&int32(3) != 0 {
		v314 = v320
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v336 = v329
	goto L125
L125:
	;
	v340 = v336 - int32(4)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v273+v340)))
	*(*int32)(unsafe.Add(mBase, uint32(v268+v340))) = v343
	if base.Ui32(int32(3)) < base.Ui32(v340) {
		v336 = v340
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v349 = v340
	goto L115
L127:
	;
	goto L126
L128:
	;
	v356 = v349
	goto L129
L129:
	;
	v360 = v356 - int32(1)
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273+v360))))
	*(*uint8)(unsafe.Add(mBase, uint32(v268+v360))) = uint8(v363)
	if v360 != 0 {
		v356 = v360
		goto L129
	} else {
		goto L131
	}
L130:
	;
	goto L95
L131:
	;
	goto L130
L132:
	;
	v373 = v366
	v374 = v367
	v375 = v368
	goto L133
L133:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	*(*int32)(unsafe.Add(mBase, uint32(v375))) = v377
	v379 = int32(4)
	v380 = v373 + v379
	v382 = v375 + v379
	v384 = v374 - v379
	if base.Ui32(int32(3)) < base.Ui32(v384) {
		v373 = v380
		v374 = v384
		v375 = v382
		goto L133
	} else {
		goto L135
	}
L134:
	;
	v388 = v380
	v389 = v384
	v390 = v382
	goto L100
L135:
	;
	goto L134
L136:
	;
	v395 = v388
	v396 = v389
	v397 = v390
	goto L137
L137:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	*(*uint8)(unsafe.Add(mBase, uint32(v397))) = uint8(v399)
	v401 = int32(1)
	v406 = v396 - v401
	if v406 != 0 {
		v395 = v395 + v401
		v396 = v406
		v397 = v397 + v401
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L95
L139:
	;
	goto L138
L140:
	;
	v437 = v429 - v428
	v438 = v426 + v427
	v442 = F_pg_do_encoding_conversion_buf(m, v431, v432, v435, v428+v430, v437, v438, int32(65537)-v426, int32(1))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L27
	} else {
		goto L141
	}
L141:
	;
	if v442 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if base.B2i32(v446 == int32(0))&base.B2i32(v437 < int32(16)) != 0 {
		goto L34
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v454 + v442
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	if v438&int32(3) == int32(0) {
		v481 = v438
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v452 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+321)) = uint8(v452)
	goto L34
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v457 + v514
	goto L34
L147:
	;
	v514 = v506 - v438
	goto L146
L148:
	;
	v485 = v481
	goto L157
L149:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v465 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v514 = int32(0)
	goto L146
L151:
	;
	goto L152
L152:
	;
	v470 = v438
	goto L153
L153:
	;
	v474 = v470 + int32(1)
	if v474&int32(3) == int32(0) {
		v481 = v474
		goto L148
	} else {
		goto L155
	}
L154:
	;
	v506 = v474
	goto L147
L155:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	if v479 != 0 {
		v470 = v474
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	v494 = int32(-2139062144)
	if (int32(16843008)-v491|v491)&v494 == v494 {
		v485 = v485 + int32(4)
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v500 = v485
	goto L160
L159:
	;
	goto L158
L160:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	if v504 != 0 {
		v500 = v500 + int32(1)
		goto L160
	} else {
		goto L162
	}
L161:
	;
	v506 = v500
	goto L147
L162:
	;
	goto L161
L163:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L27
	} else {
		goto L181
	}
L164:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	F_report_invalid_encoding(m, v571, v572+v523, v531-v523)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L27
	} else {
		goto L180
	}
L165:
	;
	if v525 <= int32(0) {
		goto L14
	} else {
		goto L179
	}
L166:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	v568 = v527
	goto L165
L167:
	;
	goto L168
L168:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+321)))
	if v528 == int32(1) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v532 == int32(0) {
		goto L164
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	if v564 == int32(0) {
		goto L163
	} else {
		goto L178
	}
L172:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v541 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	goto L173
L173:
	;
	v549 = F_pg_do_encoding_conversion_buf(m, v538, v539, v542, v536+v537, v531-v536, v523+v535, int32(65537)-v523, int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L27
	} else {
		goto L174
	}
L174:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L27
	} else {
		goto L175
	}
L175:
	;
	F_errmsg_internal(m, int32(199718), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L27
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(468032), int32(579), int32(200085))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L27
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	v568 = int32(1)
	goto L165
L179:
	;
	v581 = v523
	v582 = v524
	v584 = v524
	v588 = v568
	goto L19
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	goto L32
L182:
	;
	v648 = int32(0)
	v650 = v601 & int32(255)
	switch v650 - int32(10) {
	case 0:
		goto L194
	case 1, 2:
		v48 = v581
		v49 = v582
		v50 = v599
		v52 = v648
		v56 = v588
		v58 = v645
		v63 = v647
		goto L17
	case 3:
		goto L195
	default:
		goto L193
	}
L183:
	;
	v604 = int32(1)
	v605 = int32(0)
	if (base.B2i32(v601 != int32(13))|base.B2i32(v599 < v581)|v588)&v604 == v605 {
		v48 = v581
		v49 = v582
		v50 = v584
		v52 = v604
		v56 = v605
		goto L17
	} else {
		goto L184
	}
L184:
	;
	v615 = base.B2i32(v601 == v42)
	v617 = v63 ^ v615&v58
	v618 = v615 & v617
	if v601 == v41 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v623 = v617 ^ v58 ^ int32(-1)
	goto L187
L186:
	;
	v623 = v58
	goto L187
L187:
	;
	v624 = int32(0)
	if v623&int32(1) == v624 {
		v645 = v624
		v646 = v597
		v647 = v618
		goto L182
	} else {
		goto L188
	}
L188:
	;
	v629 = int32(1)
	v630 = int32(0)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v633 == v629 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v636 = int32(10)
	goto L191
L190:
	;
	v636 = int32(13)
	goto L191
L191:
	;
	if v636 != v601 {
		v645 = v629
		v646 = v630
		v647 = v618
		goto L182
	} else {
		goto L192
	}
L192:
	;
	v638 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v638 + int64(1)
	v645 = v629
	v646 = v630
	v647 = v618
	goto L182
L193:
	;
	if v650 != int32(92) {
		v48 = v581
		v49 = v582
		v50 = v599
		v52 = v648
		v56 = v588
		v58 = v645
		v63 = v647
		goto L17
	} else {
		goto L243
	}
L194:
	;
	if v646 == int32(0) {
		v48 = v581
		v49 = v582
		v50 = v599
		v52 = v648
		v56 = v588
		v58 = v645
		v63 = v647
		goto L17
	} else {
		goto L228
	}
L195:
	;
	if v646 == int32(0) {
		v48 = v581
		v49 = v582
		v50 = v599
		v52 = v648
		v56 = v588
		v58 = v645
		v63 = v647
		goto L17
	} else {
		goto L196
	}
L196:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v655 {
	case 0, 3:
		goto L198
	case 1:
		goto L197
	default:
		v986 = v599
		goto L9
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L27
	} else {
		goto L217
	}
L198:
	;
	v656 = int32(1)
	v657 = int32(0)
	if (base.B2i32(v599 < v581)|v588)&v656 == v657 {
		v48 = v581
		v49 = v582
		v50 = v584
		v52 = v656
		v56 = v657
		v58 = v645
		v63 = v647
		goto L17
	} else {
		goto L199
	}
L199:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v44))))
	if v665 == int32(10) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v975 = int32(3)
	v979 = v584 + int32(2)
	goto L10
L201:
	;
	goto L202
L202:
	;
	if v655 != int32(3) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v975 = int32(2)
	v979 = v599
	goto L10
L204:
	;
	goto L205
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L27
	} else {
		goto L206
	}
L206:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L27
	} else {
		goto L207
	}
L207:
	;
	if l1 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v683 = int32(473485)
	goto L210
L209:
	;
	v683 = int32(473447)
	goto L210
L210:
	;
	F_errmsg(m, v683, int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L27
	} else {
		goto L211
	}
L211:
	;
	if l1 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v689 = int32(563734)
	goto L214
L213:
	;
	v689 = int32(563785)
	goto L214
L214:
	;
	F_errhint(m, v689, int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L27
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(468032), int32(1398), int32(60451))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L27
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L27
	} else {
		goto L218
	}
L218:
	;
	if l1 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v707 = int32(473485)
	goto L221
L220:
	;
	v707 = int32(473447)
	goto L221
L221:
	;
	F_errmsg(m, v707, int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L27
	} else {
		goto L222
	}
L222:
	;
	if l1 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v713 = int32(563734)
	goto L225
L224:
	;
	v713 = int32(563785)
	goto L225
L225:
	;
	F_errhint(m, v713, int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L27
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(468032), int32(1415), int32(60451))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L27
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v724&int32(-2) != int32(2) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v975 = int32(1)
	v979 = v599
	goto L10
L230:
	;
	goto L231
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L27
	} else {
		goto L232
	}
L232:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L27
	} else {
		goto L233
	}
L233:
	;
	if l1 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v739 = int32(473554)
	goto L236
L235:
	;
	v739 = int32(473524)
	goto L236
L236:
	;
	F_errmsg(m, v739, int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L27
	} else {
		goto L237
	}
L237:
	;
	if l1 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v745 = int32(585342)
	goto L240
L239:
	;
	v745 = int32(585385)
	goto L240
L240:
	;
	F_errhint(m, v745, int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L27
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(468032), int32(1431), int32(60451))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L27
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	if l1 != 0 {
		v48 = v581
		v49 = v582
		v50 = v599
		v52 = v648
		v56 = v588
		v58 = v645
		v63 = v647
		goto L17
	} else {
		goto L244
	}
L244:
	;
	v756 = int32(1)
	v757 = int32(0)
	v758 = base.B2i32(v599 < v581)
	if (v758|v588)&v756 == v757 {
		v48 = v581
		v49 = v582
		v50 = v584
		v52 = v756
		v56 = v757
		v58 = v645
		v63 = v647
		goto L17
	} else {
		goto L245
	}
L245:
	;
	if v758 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v766 = int32(1)
	if (v588^v766)&v766 == int32(0) {
		goto L16
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v774 = v584 + int32(2)
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v44))))
	if v776 != int32(46) {
		v48 = v581
		v49 = v582
		v50 = v774
		v52 = int32(0)
		v56 = v588
		v58 = v645
		v63 = v647
		goto L17
	} else {
		goto L250
	}
L249:
	;
	goto L248
L250:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v779 == int32(3) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v782 = int32(1)
	v783 = int32(0)
	if (base.B2i32(v774 < v581)|v588)&v782 == v783 {
		v48 = v581
		v49 = v582
		v50 = v584
		v52 = v782
		v56 = v783
		v58 = v645
		v63 = v647
		goto L17
	} else {
		goto L254
	}
L252:
	;
	v800 = v774
	goto L253
L253:
	;
	v802 = int32(1)
	v803 = int32(0)
	if (base.B2i32(v800 < v581)|v588)&v802 == v803 {
		v48 = v581
		v49 = v582
		v50 = v584
		v52 = v802
		v56 = v803
		v58 = v645
		v63 = v647
		goto L17
	} else {
		goto L257
	}
L254:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v774+v44))))
	if v791 == int32(10) {
		goto L13
	} else {
		goto L255
	}
L255:
	;
	if v791 != int32(13) {
		goto L12
	} else {
		goto L256
	}
L256:
	;
	v800 = v584 + int32(3)
	goto L253
L257:
	;
	goto L18
L258:
	;
	if (base.B2i32(v779 == int32(1))|base.B2i32(v779 == int32(3)))&base.B2i32(v811 != int32(10)) != 0 {
		goto L11
	} else {
		goto L264
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L27
	} else {
		goto L260
	}
L260:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L27
	} else {
		goto L261
	}
L261:
	;
	F_errmsg(m, int32(351055), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L27
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(468032), int32(1484), int32(60451))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L27
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	if base.B2i32(v779 == int32(2))&base.B2i32(v811 != int32(13)) != 0 {
		goto L11
	} else {
		goto L265
	}
L265:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if int32(0) < v843 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L27
	} else {
		goto L269
	}
L267:
	;
	if v582 < v584 {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v874 = v800 + int32(1)
	goto L15
L269:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L27
	} else {
		goto L270
	}
L270:
	;
	F_errmsg(m, int32(351055), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L27
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(468032), int32(1500), int32(60451))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L27
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	F_appendBinaryStringInfo(m, v20, v866+v582, v599-v582)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L27
	} else {
		goto L274
	}
L274:
	;
	v874 = v599
	goto L15
L275:
	;
	goto L276
L276:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v919 = F_CopyGetData(m, l0, v917, int32(65536))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L27
	} else {
		goto L278
	}
L277:
	;
	v923 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v923
	*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = v923
	v1025 = v895
	goto L8
L278:
	;
	if int32(0) < v919 {
		goto L276
	} else {
		goto L279
	}
L279:
	;
	goto L277
L280:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L27
	} else {
		goto L281
	}
L281:
	;
	F_errmsg(m, int32(358937), int32(0))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L27
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(468032), int32(1469), int32(60451))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L27
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L27
	} else {
		goto L285
	}
L285:
	;
	F_errmsg(m, int32(351055), int32(0))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L27
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(468032), int32(1473), int32(60451))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L27
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L27
	} else {
		goto L289
	}
L289:
	;
	F_errmsg(m, int32(358937), int32(0))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L27
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(468032), int32(1491), int32(60451))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L27
	} else {
		goto L291
	}
L291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L292:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	F_appendBinaryStringInfo(m, v20, v989+v582, v986-v582)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L27
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v995 = int32(0)
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v996 - int32(1) {
	case 0:
		goto L298
	case 1:
		goto L297
	case 2:
		goto L296
	default:
		v1025 = v995
		goto L8
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v986
	goto L294
L296:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v1017 = v1015 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v1017
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v1021 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1019+v1017))) = uint8(v1021)
	v1025 = v995
	goto L8
L297:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v1009 = v1007 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v1009
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v1013 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1011+v1009))) = uint8(v1013)
	v1025 = v995
	goto L8
L298:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v1001 = v999 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v1001
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v1005 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1003+v1001))) = uint8(v1005)
	v1025 = v995
	goto L8
}
func F_line_construct_pp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v24 float64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 float64
	_ = v28
	var v31 int64
	_ = v31
	var v38 float64
	_ = v38
	var v45 int64
	_ = v45
	var v50 float64
	_ = v50
	var v66 int32
	_ = v66
	var v72 float64
	_ = v72
	var v76 int64
	_ = v76
	var v77 float64
	_ = v77
	var v80 int64
	_ = v80
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v118 float64
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_palloc(m, int32(24))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
		if base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v24 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
			v26 = int64(9223372036854775807)
			v27 = base.I64_reinterpret_f64(v24) & v26
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			v31 = base.I64_reinterpret_f64(v28) & v26
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v31) {
				v66 = base.B2i32(base.Ui64(v27) < base.Ui64(int64(9218868437227405313)))
				if base.F64_ne(v18, v24) != 0 {
					v118 = F_point_sl(m, v12, v11)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						F_line_construct(m, v14, v12, v118)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							return v14
						}
					}
				} else {
					if v66 == int32(0) {
						v118 = F_point_sl(m, v12, v11)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							F_line_construct(m, v14, v12, v118)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								return v14
							}
						}
					} else {
						v72 = v28
						v76 = v31
						v77 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
						v80 = base.I64_reinterpret_f64(v77) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v76) {
							if base.Ui64(v80) <= base.Ui64(int64(9218868437227405312)) {
								v118 = F_point_sl(m, v12, v11)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_line_construct(m, v14, v12, v118)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										return v14
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(109985), int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(463487), int32(1124), int32(220035))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
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
							if base.F64_ne(v77, v72) != 0 {
								v118 = F_point_sl(m, v12, v11)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_line_construct(m, v14, v12, v118)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										return v14
									}
								}
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v80) {
									v118 = F_point_sl(m, v12, v11)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_line_construct(m, v14, v12, v118)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											return v14
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(109985), int32(0))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(463487), int32(1124), int32(220035))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
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
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v27) {
					v118 = F_point_sl(m, v12, v11)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						F_line_construct(m, v14, v12, v118)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							return v14
						}
					}
				} else {
					v38 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					if base.Ui64(base.I64_reinterpret_f64(v38)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
						if base.F64_ne(v18, v24) != 0 {
							if base.F64_le(base.F64_abs(base.F64_sub(v18, v24)), float64(1e-06)) == int32(0) {
								v118 = F_point_sl(m, v12, v11)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_line_construct(m, v14, v12, v118)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										return v14
									}
								}
							} else {
								if base.F64_eq(v28, v38) != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(109985), int32(0))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(463487), int32(1124), int32(220035))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
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
									if base.F64_le(base.F64_abs(base.F64_sub(v28, v38)), float64(1e-06)) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(109985), int32(0))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(463487), int32(1124), int32(220035))
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
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
										v118 = F_point_sl(m, v12, v11)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											F_line_construct(m, v14, v12, v118)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												return v14
											}
										}
									}
								}
							}
						} else {
							if base.F64_eq(v28, v38) != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(109985), int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(463487), int32(1124), int32(220035))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
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
								if base.F64_le(base.F64_abs(base.F64_sub(v28, v38)), float64(1e-06)) != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(109985), int32(0))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(463487), int32(1124), int32(220035))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
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
									v118 = F_point_sl(m, v12, v11)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_line_construct(m, v14, v12, v118)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											return v14
										}
									}
								}
							}
						}
					} else {
						v66 = int32(1)
						if base.F64_ne(v18, v24) != 0 {
							v118 = F_point_sl(m, v12, v11)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								F_line_construct(m, v14, v12, v118)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									return v14
								}
							}
						} else {
							if v66 == int32(0) {
								v118 = F_point_sl(m, v12, v11)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_line_construct(m, v14, v12, v118)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										return v14
									}
								}
							} else {
								v72 = v28
								v76 = v31
								v77 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
								v80 = base.I64_reinterpret_f64(v77) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v76) {
									if base.Ui64(v80) <= base.Ui64(int64(9218868437227405312)) {
										v118 = F_point_sl(m, v12, v11)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											F_line_construct(m, v14, v12, v118)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												return v14
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(109985), int32(0))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(463487), int32(1124), int32(220035))
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return int32(0)
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
									if base.F64_ne(v77, v72) != 0 {
										v118 = F_point_sl(m, v12, v11)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											F_line_construct(m, v14, v12, v118)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												return v14
											}
										}
									} else {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v80) {
											v118 = F_point_sl(m, v12, v11)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												F_line_construct(m, v14, v12, v118)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													return v14
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(109985), int32(0))
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(463487), int32(1124), int32(220035))
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return int32(0)
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
		} else {
			v45 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui64(v45&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
				v118 = F_point_sl(m, v12, v11)
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					F_line_construct(m, v14, v12, v118)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						return v14
					}
				}
			} else {
				v50 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
				v72 = v50
				v76 = base.I64_reinterpret_f64(v50) & int64(9223372036854775807)
				v77 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v80 = base.I64_reinterpret_f64(v77) & int64(9223372036854775807)
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v76) {
					if base.Ui64(v80) <= base.Ui64(int64(9218868437227405312)) {
						v118 = F_point_sl(m, v12, v11)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							F_line_construct(m, v14, v12, v118)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								return v14
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(109985), int32(0))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(463487), int32(1124), int32(220035))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
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
					if base.F64_ne(v77, v72) != 0 {
						v118 = F_point_sl(m, v12, v11)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							F_line_construct(m, v14, v12, v118)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								return v14
							}
						}
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v80) {
							v118 = F_point_sl(m, v12, v11)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								F_line_construct(m, v14, v12, v118)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									return v14
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(109985), int32(0))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(463487), int32(1124), int32(220035))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
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
func F_line_horizontal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	return base.F64_le(base.F64_abs(v3), float64(1e-06))
}
func F_line_vertical(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	return base.F64_le(base.F64_abs(v3), float64(1e-06))
}
