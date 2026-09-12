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
	var v458 int32
	_ = v458
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v582 int64
	_ = v582
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v609 int32
	_ = v609
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
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
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v787 int32
	_ = v787
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int64
	_ = v867
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v985 int32
	_ = v985
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
	v985 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+304)) = uint8(v985)
	return v969
L9:
	;
	if v526 < v930 {
		goto L275
	} else {
		goto L276
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v919
	v930 = v923
	goto L9
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L27
	} else {
		goto L271
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L27
	} else {
		goto L267
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L27
	} else {
		goto L263
	}
L14:
	;
	v839 = int32(1)
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v840 != v839 {
		v969 = v839
		goto L8
	} else {
		goto L258
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v818
	goto L14
L16:
	;
	if v543 <= v526 {
		goto L14
	} else {
		goto L256
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
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744+v44))))
	switch v755 - int32(10) {
	case 0, 3:
		goto L241
	default:
		goto L242
	}
L19:
	;
	v541 = int32(1)
	v543 = v528 + v541
	v545 = int32(*(*int8)(unsafe.Add(mBase, uint32(v528+v44))))
	if l1 == int32(0) {
		v589 = v58
		v590 = v541
		v591 = v63
		goto L165
	} else {
		goto L166
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
	v525 = v48
	v526 = v49
	v528 = v50
	v532 = v56
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
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v469 = v467 - v468
	if v78-v79 < v469 {
		goto L149
	} else {
		goto L150
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
	v458 = F_strlen(m, v438)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v457 + v458
	goto L34
L145:
	;
	v452 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+321)) = uint8(v452)
	goto L34
L146:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L27
	} else {
		goto L164
	}
L147:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	F_report_invalid_encoding(m, v515, v516+v467, v475-v467)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L27
	} else {
		goto L163
	}
L148:
	;
	if v469 <= int32(0) {
		goto L14
	} else {
		goto L162
	}
L149:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	v512 = v471
	goto L148
L150:
	;
	goto L151
L151:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+321)))
	if v472 == int32(1) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v476 == int32(0) {
		goto L147
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	if v508 == int32(0) {
		goto L146
	} else {
		goto L161
	}
L155:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v485 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	goto L156
L156:
	;
	v493 = F_pg_do_encoding_conversion_buf(m, v482, v483, v486, v480+v481, v475-v480, v467+v479, int32(65537)-v467, int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L27
	} else {
		goto L157
	}
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L27
	} else {
		goto L158
	}
L158:
	;
	F_errmsg_internal(m, int32(212578), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L27
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(499867), int32(579), int32(212994))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L27
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	v512 = int32(1)
	goto L148
L162:
	;
	v525 = v467
	v526 = v468
	v528 = v468
	v532 = v512
	goto L19
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	goto L32
L165:
	;
	v592 = int32(0)
	v594 = v545 & int32(255)
	switch v594 - int32(10) {
	case 0:
		goto L177
	case 1, 2:
		v48 = v525
		v49 = v526
		v50 = v543
		v52 = v592
		v56 = v532
		v58 = v589
		v63 = v591
		goto L17
	case 3:
		goto L178
	default:
		goto L176
	}
L166:
	;
	v548 = int32(1)
	v549 = int32(0)
	if (base.B2i32(v545 != int32(13))|base.B2i32(v543 < v525)|v532)&v548 == v549 {
		v48 = v525
		v49 = v526
		v50 = v528
		v52 = v548
		v56 = v549
		goto L17
	} else {
		goto L167
	}
L167:
	;
	v559 = base.B2i32(v545 == v42)
	v561 = v63 ^ v559&v58
	v562 = v559 & v561
	if v545 == v41 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v567 = v561 ^ v58 ^ int32(-1)
	goto L170
L169:
	;
	v567 = v58
	goto L170
L170:
	;
	v568 = int32(0)
	if v567&int32(1) == v568 {
		v589 = v568
		v590 = v541
		v591 = v562
		goto L165
	} else {
		goto L171
	}
L171:
	;
	v573 = int32(1)
	v574 = int32(0)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v577 == v573 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v580 = int32(10)
	goto L174
L173:
	;
	v580 = int32(13)
	goto L174
L174:
	;
	if v580 != v545 {
		v589 = v573
		v590 = v574
		v591 = v562
		goto L165
	} else {
		goto L175
	}
L175:
	;
	v582 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v582 + int64(1)
	v589 = v573
	v590 = v574
	v591 = v562
	goto L165
L176:
	;
	if v594 != int32(92) {
		v48 = v525
		v49 = v526
		v50 = v543
		v52 = v592
		v56 = v532
		v58 = v589
		v63 = v591
		goto L17
	} else {
		goto L226
	}
L177:
	;
	if v590 == int32(0) {
		v48 = v525
		v49 = v526
		v50 = v543
		v52 = v592
		v56 = v532
		v58 = v589
		v63 = v591
		goto L17
	} else {
		goto L211
	}
L178:
	;
	if v590 == int32(0) {
		v48 = v525
		v49 = v526
		v50 = v543
		v52 = v592
		v56 = v532
		v58 = v589
		v63 = v591
		goto L17
	} else {
		goto L179
	}
L179:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v599 {
	case 0, 3:
		goto L181
	case 1:
		goto L180
	default:
		v930 = v543
		goto L9
	}
L180:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L27
	} else {
		goto L200
	}
L181:
	;
	v600 = int32(1)
	v601 = int32(0)
	if (base.B2i32(v543 < v525)|v532)&v600 == v601 {
		v48 = v525
		v49 = v526
		v50 = v528
		v52 = v600
		v56 = v601
		v58 = v589
		v63 = v591
		goto L17
	} else {
		goto L182
	}
L182:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v44))))
	if v609 == int32(10) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v919 = int32(3)
	v923 = v528 + int32(2)
	goto L10
L184:
	;
	goto L185
L185:
	;
	if v599 != int32(3) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v919 = int32(2)
	v923 = v543
	goto L10
L187:
	;
	goto L188
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L27
	} else {
		goto L189
	}
L189:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L27
	} else {
		goto L190
	}
L190:
	;
	if l1 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v627 = int32(505915)
	goto L193
L192:
	;
	v627 = int32(505877)
	goto L193
L193:
	;
	F_errmsg(m, v627, int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L27
	} else {
		goto L194
	}
L194:
	;
	if l1 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v633 = int32(614187)
	goto L197
L196:
	;
	v633 = int32(614238)
	goto L197
L197:
	;
	F_errhint(m, v633, int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L27
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(499867), int32(1398), int32(64309))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L27
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L27
	} else {
		goto L201
	}
L201:
	;
	if l1 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v651 = int32(505915)
	goto L204
L203:
	;
	v651 = int32(505877)
	goto L204
L204:
	;
	F_errmsg(m, v651, int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L27
	} else {
		goto L205
	}
L205:
	;
	if l1 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v657 = int32(614187)
	goto L208
L207:
	;
	v657 = int32(614238)
	goto L208
L208:
	;
	F_errhint(m, v657, int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L27
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(499867), int32(1415), int32(64309))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L27
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v668&int32(-2) != int32(2) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v919 = int32(1)
	v923 = v543
	goto L10
L213:
	;
	goto L214
L214:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L27
	} else {
		goto L215
	}
L215:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L27
	} else {
		goto L216
	}
L216:
	;
	if l1 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v683 = int32(505984)
	goto L219
L218:
	;
	v683 = int32(505954)
	goto L219
L219:
	;
	F_errmsg(m, v683, int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L27
	} else {
		goto L220
	}
L220:
	;
	if l1 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v689 = int32(635966)
	goto L223
L222:
	;
	v689 = int32(636009)
	goto L223
L223:
	;
	F_errhint(m, v689, int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L27
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(499867), int32(1431), int32(64309))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L27
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	if l1 != 0 {
		v48 = v525
		v49 = v526
		v50 = v543
		v52 = v592
		v56 = v532
		v58 = v589
		v63 = v591
		goto L17
	} else {
		goto L227
	}
L227:
	;
	v700 = int32(1)
	v701 = int32(0)
	v702 = base.B2i32(v543 < v525)
	if (v702|v532)&v700 == v701 {
		v48 = v525
		v49 = v526
		v50 = v528
		v52 = v700
		v56 = v701
		v58 = v589
		v63 = v591
		goto L17
	} else {
		goto L228
	}
L228:
	;
	if v702 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v710 = int32(1)
	if (v532^v710)&v710 == int32(0) {
		goto L16
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v718 = v528 + int32(2)
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v44))))
	if v720 != int32(46) {
		v48 = v525
		v49 = v526
		v50 = v718
		v52 = int32(0)
		v56 = v532
		v58 = v589
		v63 = v591
		goto L17
	} else {
		goto L233
	}
L232:
	;
	goto L231
L233:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v723 == int32(3) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v726 = int32(1)
	v727 = int32(0)
	if (base.B2i32(v718 < v525)|v532)&v726 == v727 {
		v48 = v525
		v49 = v526
		v50 = v528
		v52 = v726
		v56 = v727
		v58 = v589
		v63 = v591
		goto L17
	} else {
		goto L237
	}
L235:
	;
	v744 = v718
	goto L236
L236:
	;
	v746 = int32(1)
	v747 = int32(0)
	if (base.B2i32(v744 < v525)|v532)&v746 == v747 {
		v48 = v525
		v49 = v526
		v50 = v528
		v52 = v746
		v56 = v747
		v58 = v589
		v63 = v591
		goto L17
	} else {
		goto L240
	}
L237:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718+v44))))
	if v735 == int32(10) {
		goto L13
	} else {
		goto L238
	}
L238:
	;
	if v735 != int32(13) {
		goto L12
	} else {
		goto L239
	}
L239:
	;
	v744 = v528 + int32(3)
	goto L236
L240:
	;
	goto L18
L241:
	;
	if (base.B2i32(v723 == int32(1))|base.B2i32(v723 == int32(3)))&base.B2i32(v755 != int32(10)) != 0 {
		goto L11
	} else {
		goto L247
	}
L242:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L27
	} else {
		goto L243
	}
L243:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L27
	} else {
		goto L244
	}
L244:
	;
	F_errmsg(m, int32(374484), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L27
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(499867), int32(1484), int32(64309))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L27
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	if base.B2i32(v723 == int32(2))&base.B2i32(v755 != int32(13)) != 0 {
		goto L11
	} else {
		goto L248
	}
L248:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if int32(0) < v787 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L27
	} else {
		goto L252
	}
L250:
	;
	if v526 < v528 {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v818 = v744 + int32(1)
	goto L15
L252:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L27
	} else {
		goto L253
	}
L253:
	;
	F_errmsg(m, int32(374484), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L27
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(499867), int32(1500), int32(64309))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L27
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	F_appendBinaryStringInfo(m, v20, v810+v526, v543-v526)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L27
	} else {
		goto L257
	}
L257:
	;
	v818 = v543
	goto L15
L258:
	;
	goto L259
L259:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v863 = F_CopyGetData(m, l0, v861, int32(65536))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L27
	} else {
		goto L261
	}
L260:
	;
	v867 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v867
	*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = v867
	v969 = v839
	goto L8
L261:
	;
	if int32(0) < v863 {
		goto L259
	} else {
		goto L262
	}
L262:
	;
	goto L260
L263:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L27
	} else {
		goto L264
	}
L264:
	;
	F_errmsg(m, int32(383426), int32(0))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L27
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(499867), int32(1469), int32(64309))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L27
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L27
	} else {
		goto L268
	}
L268:
	;
	F_errmsg(m, int32(374484), int32(0))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L27
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(499867), int32(1473), int32(64309))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L27
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L27
	} else {
		goto L272
	}
L272:
	;
	F_errmsg(m, int32(383426), int32(0))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L27
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(499867), int32(1491), int32(64309))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L27
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	F_appendBinaryStringInfo(m, v20, v933+v526, v930-v526)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L27
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v939 = int32(0)
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v940 - int32(1) {
	case 0:
		goto L281
	case 1:
		goto L280
	case 2:
		goto L279
	default:
		v969 = v939
		goto L8
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v930
	goto L277
L279:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v961 = v959 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v961
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v965 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v963+v961))) = uint8(v965)
	v969 = v939
	goto L8
L280:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v953 = v951 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v953
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v957 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v955+v953))) = uint8(v957)
	v969 = v939
	goto L8
L281:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v945 = v943 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v945
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v949 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v947+v945))) = uint8(v949)
	v969 = v939
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
										F_errmsg(m, int32(119594), int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
											F_errmsg(m, int32(119594), int32(0))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
											F_errmsg(m, int32(119594), int32(0))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
												F_errmsg(m, int32(119594), int32(0))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
										F_errmsg(m, int32(119594), int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
											F_errmsg(m, int32(119594), int32(0))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
												F_errmsg(m, int32(119594), int32(0))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
													F_errmsg(m, int32(119594), int32(0))
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
								F_errmsg(m, int32(119594), int32(0))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
									F_errmsg(m, int32(119594), int32(0))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(494958), int32(1124), int32(234603))
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
