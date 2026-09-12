package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecHashJoin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 float64
	_ = v568
	var v572 int32
	_ = v572
	var v575 float64
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 float64
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v806 float64
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v909 int32
	_ = v909
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1109 int32
	_ = v1109
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	F_MemoryContextReset(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v41 = v20
	goto L3
L3:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v49 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	switch v52 - int32(1) {
	case 0:
		goto L20
	case 1:
		v106 = v41
		goto L19
	case 2:
		v351 = v41
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	case 5:
		goto L15
	default:
		goto L12
	}
L8:
	;
	goto L7
L9:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v1109 != 0 {
		goto L271
	} else {
		goto L272
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v1079
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810)+48)) = v844
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v810)))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v810)+120))
	F_MemoryContextReset(m, v943)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L241
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L238
	}
L13:
	;
	m.G0 = v18 + int32(32)
	return v909
L14:
	;
	v909 = int32(0)
	goto L13
L15:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)+44))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v810)+48))
	if int32(0) < v812 {
		goto L212
	} else {
		goto L213
	}
L16:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v639 = v636
	goto L176
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+168)))
	if v581 != 0 {
		goto L3
	} else {
		goto L164
	}
L18:
	;
	v358 = m.G0
	v360 = v358 - int32(16)
	m.G0 = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v362 != 0 {
		v379 = v362
		goto L108
	} else {
		goto L109
	}
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+48))
	if v108 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v55 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v87 = F_ExecHashTableCreate(m, v22)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L39
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	goto L21
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v56 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v60 = *(*float64)(unsafe.Add(mBase, uint32(v59)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v61)+16))
	if base.F64_lt(v60, v62) == int32(0) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v67 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+169)))
	if v66 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v71 = m.T0[v70].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v71
	if v71 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v82 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+169)) = uint8(v82)
	goto L21
L35:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
	if v74&int32(2) == int32(0) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+169)) = uint8(v79)
	v909 = v79
	goto L13
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v22)+104)) = v87
	v91 = F_MultiExecProcNode(m, v22)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v87)+64))
	if base.F64_eq(v93, float64(0)) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v96 == int32(0) {
		goto L14
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+56)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+169)) = uint8(v103)
	v106 = v87
	goto L19
L44:
	;
	goto L43
L45:
	;
	if v205&int32(2) != 0 {
		goto L9
	} else {
		goto L81
	}
L46:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v111 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	goto L48
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
	if v185 <= v108 {
		goto L9
	} else {
		goto L76
	}
L49:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
	if v128&int32(2) != 0 {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v119 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
	if v114&int32(2) != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	v127 = v111
	goto L49
L53:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v123 = m.T0[v122].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	if v123 == int32(0) {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v127 = v123
	goto L49
L59:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v1079 = v131
	goto L10
L60:
	;
	goto L61
L61:
	;
	v134 = v127
	goto L62
L62:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v134
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	F_MemoryContextReset(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	v1079 = v161
	goto L10
L64:
	;
	v152 = int32(4476144)
	v153 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v156
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	v161 = m.T0[v160].(func(*base.Module, int32, int32, int32) int32)(m, v154, v147, v18+int32(28))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v153
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+28)))
	if v165 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v161
	v169 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+169)) = uint8(v169)
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)))
	v205 = v171
	v206 = v134
	goto L45
L67:
	;
	goto L68
L68:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v172 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v176 = m.T0[v175].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	if v176 == int32(0) {
		v1079 = v161
		goto L10
	} else {
		goto L74
	}
L74:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+4)))
	if v180&int32(2) == int32(0) {
		v134 = v176
		goto L62
	} else {
		goto L75
	}
L75:
	;
	goto L63
L76:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v107)+92))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187+v108<<(uint(int32(2))%32))))
	if v191 == int32(0) {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v197 = F_ExecHashJoinGetSavedTuple(m, v191, v18+int32(24), v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v197 == int32(0) {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+4)))
	if v201&int32(2) != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	v205 = v201
	v206 = v197
	goto L45
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v206
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+168)) = uint8(v222)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v224
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v106)+44))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(132)))) = (v230 - int32(1)) & v224
	if base.Ui32(int32(2)) <= base.Ui32(v229) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+24)))
	if v251 != int32(1) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v243 = (v229 - int32(1)) & base.I32_rotr(v224, v240)
	goto L85
L84:
	;
	v243 = v222
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(20)))) = v243
	goto L82
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v292
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	if v296 == v297 {
		goto L96
	} else {
		goto L97
	}
L87:
	;
	v292 = int32(-1)
	goto L86
L88:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v106)+28))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v106)+32))
	v257 = v255 - int32(1)
	v258 = v245 & v257
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v254+v258<<(uint(int32(2))%32))))
	if v262 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v265 = v258
	v268 = v262
	goto L90
L90:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v245 == v271 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L87
L92:
	;
	v292 = v265
	goto L86
L93:
	;
	goto L94
L94:
	;
	v275 = (v265 + int32(1)) & v257
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v254+v275<<(uint(int32(2))%32))))
	if v279 != 0 {
		v265 = v275
		v268 = v279
		goto L90
	} else {
		goto L95
	}
L95:
	;
	goto L91
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(3)
	v351 = v106
	goto L18
L97:
	;
	if v292 != int32(-1) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v303 = F_ExecFetchSlotMinimalTuple(m, v206, v18+int32(19))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v106)+92))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v307
	v311 = v305 + v306<<(uint(int32(2))%32)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	if v312 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v315 = int32(4476144)
	v316 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v106)+124))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v318
	v321 = F_BufFileCreateTemp(m, int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	v326 = v312
	goto L102
L102:
	;
	F_BufFileWrite(m, v326, v18+int32(28), int32(4))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v321
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v316
	v326 = v321
	goto L102
L104:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	F_BufFileWrite(m, v326, v303, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+19)))
	if v336 != int32(1) {
		v41 = v106
		goto L3
	} else {
		goto L106
	}
L106:
	;
	F_pfree(m, v303)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v41 = v106
	goto L3
L108:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	if v381 != 0 {
		goto L114
	} else {
		goto L115
	}
L109:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v364 != int32(-1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363)+28))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367+v364<<(uint(int32(2))%32))))
	v379 = v371 + int32(4)
	goto L108
L111:
	;
	goto L112
L112:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v363)+20))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v379 = v374 + v375<<(uint(int32(2))%32)
	goto L108
L113:
	;
	m.G0 = v360 + int32(16)
	if v468 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L114:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v385 = v381
	goto L117
L115:
	;
	goto L116
L116:
	;
	v468 = int32(0)
	goto L113
L117:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v399 != v382 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L116
L119:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v436 != 0 {
		v385 = v436
		goto L117
	} else {
		goto L130
	}
L120:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v405 = F_ExecStoreMinimalTuple(m, v385+int32(8), v403, int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v405
	if v383 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v385
	v468 = int32(1)
	goto L113
L123:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	F_MemoryContextReset(m, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v413 = int32(4476144)
	v414 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v416
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v383)+20))
	v421 = m.T0[v420].(func(*base.Module, int32, int32, int32) int32)(m, v383, v25, v360+int32(15))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L127
	}
L126:
	;
	goto L122
L127:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v414
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	F_MemoryContextReset(m, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v421 == int32(0) {
		goto L119
	} else {
		goto L129
	}
L129:
	;
	goto L122
L130:
	;
	goto L118
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(4)
	v41 = v351
	goto L3
L132:
	;
	goto L133
L133:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v476 == int32(6) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v480 = int32(*(*int16)(unsafe.Add(mBase, uint32(v479)+18)))
	if v480 < int32(0) {
		v41 = v351
		goto L3
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	if v24 != 0 {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	goto L136
L138:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v572 == int32(0) {
		v41 = v351
		goto L3
	} else {
		goto L163
	}
L139:
	;
	v483 = int32(4476144)
	v484 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v486
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v491 = m.T0[v490].(func(*base.Module, int32, int32, int32) int32)(m, v24, v25, v18+int32(28))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v499 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+168)) = uint8(v499)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v503 = v501 + int32(18)
	v504 = int32(*(*int16)(unsafe.Add(mBase, uint32(v503))))
	if int32(0) <= v504 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v484
	if v491 == int32(0) {
		goto L138
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v508 = v504 | int32(32768)
	*(*uint16)(unsafe.Add(mBase, uint32(v503))) = uint16(v508)
	goto L146
L145:
	;
	goto L146
L146:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v510 == int32(5) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	v41 = v351
	goto L3
L148:
	;
	goto L149
L149:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v515 == int32(1) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	goto L152
L151:
	;
	goto L152
L152:
	;
	if v510 == int32(7) {
		v41 = v351
		goto L3
	} else {
		goto L153
	}
L153:
	;
	if v23 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v565 == int32(0) {
		v41 = v351
		goto L3
	} else {
		goto L162
	}
L155:
	;
	v522 = int32(4476144)
	v523 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v525
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v530 = m.T0[v529].(func(*base.Module, int32, int32, int32) int32)(m, v23, v25, v18+int32(28))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+72))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v538)+16))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	m.T0[v542].(func(*base.Module, int32))(m, v540)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v523
	if v530 == int32(0) {
		goto L154
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v545 = int32(4476144)
	v546 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v539)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v548
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v538)+24))
	v554 = m.T0[v553].(func(*base.Module, int32, int32, int32) int32)(m, v538+int32(4), v539, int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v546
	v558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+4)))
	v560 = v558 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v540)+4)) = uint16(v560)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v540)+12))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	*(*uint16)(unsafe.Add(mBase, uint32(v540)+6)) = uint16(v563)
	v909 = v540
	goto L13
L162:
	;
	v568 = *(*float64)(unsafe.Add(mBase, uint32(v565)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v565)+248)) = base.F64_add(v568, float64(1))
	v41 = v351
	goto L3
L163:
	;
	v575 = *(*float64)(unsafe.Add(mBase, uint32(v572)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v572)+240)) = base.F64_add(v575, float64(1))
	v41 = v351
	goto L3
L164:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v582 == int32(0) {
		goto L3
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v582
	if v23 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v629 == int32(0) {
		goto L3
	} else {
		goto L174
	}
L167:
	;
	v586 = int32(4476144)
	v587 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v589
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v594 = m.T0[v593].(func(*base.Module, int32, int32, int32) int32)(m, v23, v25, v18+int32(28))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)+72))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v602)+16))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)+8))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+12))
	m.T0[v606].(func(*base.Module, int32))(m, v604)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L172
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v587
	if v594 == int32(0) {
		goto L166
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	v609 = int32(4476144)
	v610 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v612
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v602)+24))
	v618 = m.T0[v617].(func(*base.Module, int32, int32, int32) int32)(m, v602+int32(4), v603, int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v610
	v622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+4)))
	v624 = v622 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v604)+4)) = uint16(v624)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v604)+12))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	*(*uint16)(unsafe.Add(mBase, uint32(v604)+6)) = uint16(v627)
	v909 = v604
	goto L13
L174:
	;
	v632 = *(*float64)(unsafe.Add(mBase, uint32(v629)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v629)+248)) = base.F64_add(v632, float64(1))
	goto L3
L175:
	;
	if v753 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L176:
	;
	if v639 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v683 != 0 {
		goto L186
	} else {
		goto L187
	}
L179:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v639)))
	v683 = v653
	goto L178
L180:
	;
	goto L181
L181:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	if v654 < v655 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v637)+20))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v657+v654<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v654 + int32(1)
	v683 = v661
	goto L178
L183:
	;
	goto L184
L184:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v637)+36))
	if v667 <= v666 {
		v753 = int32(0)
		goto L175
	} else {
		goto L185
	}
L185:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v637)+28))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v637)+40))
	v671 = int32(2)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v670+v666<<(uint(v671)%32))))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v669+v674<<(uint(v671)%32))))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v666 + int32(1)
	v683 = v679
	goto L178
L186:
	;
	v686 = v683
	goto L189
L187:
	;
	goto L188
L188:
	;
	v731 = int32(0)
	v733 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v733 == v731 {
		v639 = v731
		goto L176
	} else {
		goto L197
	}
L189:
	;
	v700 = int32(*(*int16)(unsafe.Add(mBase, uint32(v686)+18)))
	if int32(0) <= v700 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	goto L188
L191:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v707 = F_ExecStoreMinimalTuple(m, v686+int32(8), v705, int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v686)))
	if v715 != 0 {
		v686 = v715
		goto L189
	} else {
		goto L196
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v707
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	F_MemoryContextReset(m, v710)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v686
	v753 = int32(1)
	goto L175
L196:
	;
	goto L190
L197:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v639 = v731
	goto L176
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(6)
	goto L3
L200:
	;
	goto L201
L201:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v758
	if v23 != 0 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v803 == int32(0) {
		goto L3
	} else {
		goto L210
	}
L203:
	;
	v760 = int32(4476144)
	v761 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v763
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v768 = m.T0[v767].(func(*base.Module, int32, int32, int32) int32)(m, v23, v25, v18+int32(28))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v776)+72))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v776)+16))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v778)+8))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v779)+12))
	m.T0[v780].(func(*base.Module, int32))(m, v778)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v761
	if v768 == int32(0) {
		goto L202
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	v783 = int32(4476144)
	v784 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v777)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v786
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v776)+24))
	v792 = m.T0[v791].(func(*base.Module, int32, int32, int32) int32)(m, v776+int32(4), v777, int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v784
	v796 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v778)+4)))
	v798 = v796 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v778)+4)) = uint16(v798)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v778)+12))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v800)))
	*(*uint16)(unsafe.Add(mBase, uint32(v778)+6)) = uint16(v801)
	v909 = v778
	goto L13
L210:
	;
	v806 = *(*float64)(unsafe.Add(mBase, uint32(v803)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v803)+248)) = base.F64_add(v806, float64(1))
	goto L3
L211:
	;
	v839 = v812 + int32(1)
	if v811 <= v839 {
		goto L14
	} else {
		goto L219
	}
L212:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v810)+92))
	v817 = v812 << (uint(int32(2)) % 32)
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v815+v817)))
	if v819 != 0 {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	v827 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v810)+28)) = v827
	*(*uint8)(unsafe.Add(mBase, uint32(v810)+24)) = uint8(v827)
	*(*int32)(unsafe.Add(mBase, uint32(v810)+108)) = v827
	*(*int64)(unsafe.Add(mBase, uint32(v810)+36)) = int64(0)
	goto L211
L215:
	;
	F_BufFileClose(m, v819)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L218
	}
L216:
	;
	v823 = v815
	goto L217
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v823+v817))) = int32(0)
	goto L211
L218:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v810)+92))
	v823 = v822
	goto L217
L219:
	;
	v844 = v839
	goto L220
L220:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v810)+88))
	v858 = v844 << (uint(int32(2)) % 32)
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v856+v858)))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v810)+92))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v861+v858)))
	if v863 != 0 {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L14
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v858+v875))) = int32(0)
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v810)+92))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v879+v858)))
	if v881 != 0 {
		goto L233
	} else {
		goto L234
	}
L223:
	;
	if v860 != 0 {
		goto L11
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	if v860 == int32(0) {
		v875 = v856
		goto L222
	} else {
		goto L229
	}
L226:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v864 != 0 {
		goto L11
	} else {
		goto L227
	}
L227:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v810)+56))
	if v811 == v865 {
		v875 = v856
		goto L222
	} else {
		goto L228
	}
L228:
	;
	goto L11
L229:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v869 != 0 {
		goto L11
	} else {
		goto L230
	}
L230:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v810)+52))
	if v811 != v870 {
		goto L11
	} else {
		goto L231
	}
L231:
	;
	F_BufFileClose(m, v860)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v810)+88))
	v875 = v874
	goto L222
L233:
	;
	F_BufFileClose(m, v881)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	v885 = v879
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v885+v858))) = int32(0)
	v890 = v844 + int32(1)
	if v890 != v811 {
		v844 = v890
		goto L220
	} else {
		goto L237
	}
L236:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v810)+92))
	v885 = v884
	goto L235
L237:
	;
	goto L221
L238:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v931
	F_errmsg_internal(m, int32(478421), v18)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(491075), int32(672), int32(298034))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	v946 = int32(4476144)
	v947 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v810)+120))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v949
	v953 = F_palloc0(m, v942<<(uint(int32(2))%32))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v955 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v810)+96)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v810)+20)) = v953
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v947
	*(*int32)(unsafe.Add(mBase, uint32(v810)+128)) = v955
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v810)+88))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v962+v858)))
	if v964 != 0 {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L267
	}
L244:
	;
	v965 = int32(0)
	v968 = F_BufFileSeek(m, v964, v965, int64(0), v965)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v810)+92))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1034+v858)))
	if v1036 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L247:
	;
	if v968 != 0 {
		goto L243
	} else {
		goto L248
	}
L248:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v973 = F_ExecHashJoinGetSavedTuple(m, v964, v18+int32(28), v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	if v973 != 0 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v978 = v973
	goto L253
L251:
	;
	goto L252
L252:
	;
	F_BufFileClose(m, v964)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L258
	}
L253:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	F_ExecHashTableInsert(m, v810, v978, v990)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L255
	}
L254:
	;
	goto L252
L255:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v996 = F_ExecHashJoinGetSavedTuple(m, v964, v18+int32(28), v995)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	if v996 != 0 {
		v978 = v996
		goto L253
	} else {
		goto L257
	}
L257:
	;
	goto L254
L258:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v810)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v1015+v858))) = int32(0)
	goto L246
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	goto L3
L260:
	;
	v1039 = int32(0)
	v1042 = F_BufFileSeek(m, v1036, v1039, int64(0), v1039)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	if v1042 == int32(0) {
		goto L259
	} else {
		goto L262
	}
L262:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	F_errmsg(m, int32(382733), int32(0))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(491075), int32(1260), int32(322495))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_errmsg(m, int32(382733), int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(491075), int32(1230), int32(322495))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(5)
	v41 = v106
	goto L3
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(6)
	v41 = v106
	goto L3
}
func F_ExecHashTableCreate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v65 float64
	_ = v65
	var v67 int32
	_ = v67
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v94 float64
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 float64
	_ = v126
	var v130 float64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 float64
	_ = v140
	var v142 float64
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v168 float64
	_ = v168
	var v170 int32
	_ = v170
	var v174 float64
	_ = v174
	var v175 float64
	_ = v175
	var v178 float64
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 float64
	_ = v215
	var v217 float64
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v241 float64
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v282 float64
	_ = v282
	var v284 float64
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v402 int64
	_ = v402
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
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
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v684 float64
	_ = v684
	var v687 int32
	_ = v687
	var v688 float32
	_ = v688
	var v691 float32
	_ = v691
	var v694 float32
	_ = v694
	var v697 float32
	_ = v697
	var v699 float64
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v718 float64
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v732 float64
	_ = v732
	var v736 float32
	_ = v736
	var v738 float64
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v757 float64
	_ = v757
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v921 int32
	_ = v921
	var v940 int32
	_ = v940
	var v956 int32
	_ = v956
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+36)))
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = v19 + int32(88)
	goto L3
L2:
	;
	v26 = v22 + int32(24)
	goto L3
L3:
	;
	v27 = *(*float64)(unsafe.Add(mBase, uint32(v26)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	v35 = v32 - int32(1)
	goto L6
L5:
	;
	v35 = int32(0)
	goto L6
L6:
	;
	v36 = int32(0)
	v37 = base.B2i32(v28 != v36)
	v39 = base.B2i32(v31 != v36)
	v41 = v15 + int32(-40)
	v47 = v15 + int32(-52)
	v61 = (v29 + int32(7)) & int32(-8)
	v65 = *(*float64)(unsafe.Add(mBase, _consts[433]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	v71 = base.F64_mul(base.F64_mul(v65, base.F64_convert_i32_s(v67)), float64(1024))
	v72 = float64(4.294967295e+09)
	if base.F64_lt(v71, v72) != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v382 = int32(1073741823)
	if v382 <= v380 {
		goto L118
	} else {
		goto L119
	}
L8:
	;
	if base.F64_le(v27, float64(0)) != 0 {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	if v39 == int32(0) {
		v103 = v83
		goto L8
	} else {
		goto L16
	}
L10:
	;
	v75 = v71
	goto L12
L11:
	;
	v75 = v72
	goto L12
L12:
	;
	if base.F64_lt(v75, float64(4.294967296e+09))&base.F64_ge(v75, float64(0)) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v81 = base.I32_trunc_f64_u(v75)
	v83 = v81
	goto L9
L14:
	;
	goto L15
L15:
	;
	v83 = int32(0)
	goto L9
L16:
	;
	v90 = base.F64_mul(base.F64_convert_i32_s(v35+int32(1)), base.F64_convert_i32_u(v83))
	v91 = float64(4.294967295e+09)
	if base.F64_lt(v90, v91) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v94 = v90
	goto L19
L18:
	;
	v94 = v91
	goto L19
L19:
	;
	if base.F64_lt(v94, float64(4.294967296e+09))&base.F64_ge(v94, float64(0)) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v100 = base.I32_trunc_f64_u(v94)
	v103 = v100
	goto L8
L21:
	;
	goto L22
L22:
	;
	v103 = int32(0)
	goto L8
L23:
	;
	v105 = float64(1000)
	goto L25
L24:
	;
	v105 = v27
	goto L25
L25:
	;
	v108 = v61 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v103
	if v28 != v36 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v111 = base.I32_div_u_s(v103, v108)
	v112 = int32(50)
	v113 = base.I32_div_u_s(v111, v112)
	if base.Ui32(v112) <= base.Ui32(v111) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v120 = v103
	v121 = int32(0)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v121
	v124 = int32(1)
	v126 = base.F64_mul(v105, base.F64_convert_i32_s(v61+int32(24)))
	v130 = base.F64_ceil(v105)
	v132 = int32(268435455)
	v134 = int32(base.Ui32(v120) >> (uint(int32(2)) % 32))
	if base.Ui32(v132) <= base.Ui32(v134) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v118 = v113 * v108
	goto L31
L30:
	;
	v118 = int32(0)
	goto L31
L31:
	;
	v120 = v103 - v118
	v121 = v113
	goto L28
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(-44)))) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(-48)))) = v371
	goto L7
L33:
	;
	if v148 <= int32(1024) {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v137 = v132
	goto L36
L35:
	;
	v137 = v134
	goto L36
L36:
	;
	v139 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v137)) % 32))
	v140 = base.F64_convert_i32_u(v139)
	if base.F64_gt(v140, v130) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v142 = v130
	goto L39
L38:
	;
	v142 = v140
	goto L39
L39:
	;
	if base.F64_lt(base.F64_abs(v142), float64(2.147483648e+09)) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v146 = base.I32_trunc_f64_s(v142)
	v148 = v146
	goto L33
L41:
	;
	goto L42
L42:
	;
	v148 = int32(-2147483648)
	goto L33
L43:
	;
	v151 = int32(1024)
	goto L45
L44:
	;
	v151 = v148
	goto L45
L45:
	;
	if v151&(v151-int32(1)) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v158 = v124 << (uint(int32(32)-base.I32_clz(v151)) % 32)
	goto L48
L47:
	;
	v158 = v151
	goto L48
L48:
	;
	if base.F64_lt(base.F64_convert_i32_u(v120), base.F64_add(v126, base.F64_convert_i32_u(v158<<(uint(int32(2))%32)))) == int32(0) {
		v365 = v158
		v371 = v124
		goto L32
	} else {
		goto L49
	}
L49:
	;
	if v31 != v36 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v168 = *(*float64)(unsafe.Add(mBase, _consts[433]))
	v170 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	v174 = base.F64_mul(base.F64_mul(v168, base.F64_convert_i32_s(v170)), float64(1024))
	v175 = float64(4.294967295e+09)
	if base.F64_lt(v174, v175) != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v241 = v140
	v242 = v120
	v246 = v139
	goto L52
L52:
	;
	v248 = int32(1)
	v253 = v61 + int32(28)
	if base.Ui32(v253) < base.Ui32(v242) {
		goto L85
	} else {
		goto L86
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v186
	if v28 != v36 {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	v178 = v174
	goto L56
L55:
	;
	v178 = v175
	goto L56
L56:
	;
	if base.F64_lt(v178, float64(4.294967296e+09))&base.F64_ge(v178, float64(0)) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v184 = base.I32_trunc_f64_u(v178)
	v186 = v184
	goto L53
L58:
	;
	goto L59
L59:
	;
	v186 = int32(0)
	goto L53
L60:
	;
	v188 = base.I32_div_u_s(v186, v108)
	v189 = int32(50)
	v190 = base.I32_div_u_s(v188, v189)
	if base.Ui32(v189) <= base.Ui32(v188) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v197 = v186
	v198 = int32(0)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v198
	v201 = int32(1)
	v207 = int32(268435455)
	v209 = int32(base.Ui32(v197) >> (uint(int32(2)) % 32))
	if base.Ui32(v207) <= base.Ui32(v209) {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v195 = v190 * v108
	goto L65
L64:
	;
	v195 = int32(0)
	goto L65
L65:
	;
	v197 = v186 - v195
	v198 = v190
	goto L62
L66:
	;
	if v223 <= int32(1024) {
		goto L76
	} else {
		goto L77
	}
L67:
	;
	v212 = v207
	goto L69
L68:
	;
	v212 = v209
	goto L69
L69:
	;
	v214 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v212)) % 32))
	v215 = base.F64_convert_i32_u(v214)
	if base.F64_gt(v215, v130) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v217 = v130
	goto L72
L71:
	;
	v217 = v215
	goto L72
L72:
	;
	if base.F64_lt(base.F64_abs(v217), float64(2.147483648e+09)) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v221 = base.I32_trunc_f64_s(v217)
	v223 = v221
	goto L66
L74:
	;
	goto L75
L75:
	;
	v223 = int32(-2147483648)
	goto L66
L76:
	;
	v226 = int32(1024)
	goto L78
L77:
	;
	v226 = v223
	goto L78
L78:
	;
	if v226&(v226-int32(1)) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v233 = v201 << (uint(int32(32)-base.I32_clz(v226)) % 32)
	goto L81
L80:
	;
	v233 = v226
	goto L81
L81:
	;
	if base.F64_lt(base.F64_convert_i32_u(v197), base.F64_add(v126, base.F64_convert_i32_u(v233<<(uint(int32(2))%32)))) == int32(0) {
		v365 = v233
		v371 = v201
		goto L32
	} else {
		goto L82
	}
L82:
	;
	v241 = v215
	v242 = v197
	v246 = v214
	goto L52
L83:
	;
	v365 = v353
	v371 = v345
	goto L32
L84:
	;
	if v290 <= int32(2) {
		goto L103
	} else {
		goto L104
	}
L85:
	;
	v255 = int32(1)
	v257 = base.I32_div_u_s(v242, v253)
	if v257&(v257-v255) != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v265 = v248
	goto L87
L87:
	;
	if base.Ui32(v265) < base.Ui32(v246) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v264 = v255 << (uint(int32(32)-base.I32_clz(v257)) % 32)
	goto L90
L89:
	;
	v264 = v257
	goto L90
L90:
	;
	v265 = v264
	goto L87
L91:
	;
	v269 = v265
	goto L93
L92:
	;
	v269 = v246
	goto L93
L93:
	;
	if v269&(v269-int32(1)) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v276 = int32(1) << (uint(int32(32)-base.I32_clz(v269)) % 32)
	goto L96
L95:
	;
	v276 = v269
	goto L96
L96:
	;
	v282 = base.F64_ceil(base.F64_div(v126, base.F64_convert_i32_u(v242-v276<<(uint(int32(2))%32))))
	if base.F64_gt(v241, v282) != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v284 = v282
	goto L99
L98:
	;
	v284 = v241
	goto L99
L99:
	;
	if base.F64_lt(base.F64_abs(v284), float64(2.147483648e+09)) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v288 = base.I32_trunc_f64_s(v284)
	v290 = v288
	goto L84
L101:
	;
	goto L102
L102:
	;
	v290 = int32(-2147483648)
	goto L84
L103:
	;
	v293 = int32(2)
	goto L105
L104:
	;
	v293 = v290
	goto L105
L105:
	;
	if v293&(v293-int32(1)) != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v300 = v248 << (uint(int32(32)-base.I32_clz(v293)) % 32)
	goto L108
L107:
	;
	v300 = v293
	goto L108
L108:
	;
	if v300 < int32(2) {
		v345 = v300
		v353 = v276
		goto L83
	} else {
		goto L109
	}
L109:
	;
	if base.Ui32(int32(134217727)) < base.Ui32(v276) {
		v345 = v300
		v353 = v276
		goto L83
	} else {
		goto L110
	}
L110:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v307 = v300
	v315 = v276
	v317 = v305
	goto L111
L111:
	;
	if v317 < int32(0) {
		v345 = v307
		v353 = v315
		goto L83
	} else {
		goto L113
	}
L112:
	;
	v365 = v339
	v371 = v337
	goto L32
L113:
	;
	if base.Ui32(v307) < base.Ui32(int32(base.Ui32(v317)>>(uint(int32(13))%32))) {
		v345 = v307
		v353 = v315
		goto L83
	} else {
		goto L114
	}
L114:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v329 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v328 << (uint(v329) % 32)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v334 = v332 << (uint(v329) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v334
	v337 = int32(base.Ui32(v307) >> (uint(v329) % 32))
	v339 = v315 << (uint(v329) % 32)
	if base.Ui32(v307) < base.Ui32(int32(4)) {
		v365 = v339
		v371 = v337
		goto L32
	} else {
		goto L115
	}
L115:
	;
	if base.Ui32(v315) < base.Ui32(int32(67108864)) {
		v307 = v337
		v315 = v339
		v317 = v334
		goto L111
	} else {
		goto L116
	}
L116:
	;
	goto L112
L117:
	;
	v395 = F_palloc(m, int32(152))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L124
	} else {
		goto L125
	}
L118:
	;
	v385 = v382
	goto L120
L119:
	;
	v385 = v380
	goto L120
L120:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v385) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v393 = int32(32) - base.I32_clz(v385-int32(1))
	goto L123
L122:
	;
	v393 = int32(0)
	goto L123
L123:
	;
	goto L117
L124:
	;
	return int32(0)
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+12)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v395)+8)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v380
	v402 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v395)+28)) = v402
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+24)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+20)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v395)+16)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v395)+4)) = v393
	*(*int64)(unsafe.Add(mBase, uint32(v395)+36)) = v402
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+104)) = v404
	v415 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+60)) = uint8(v415)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+56)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v395)+52)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v395)+48)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v395)+44)) = v412
	*(*int64)(unsafe.Add(mBase, uint32(v395)+64)) = v402
	*(*int64)(unsafe.Add(mBase, uint32(v395)+72)) = v402
	*(*int64)(unsafe.Add(mBase, uint32(v395)+80)) = v402
	*(*int64)(unsafe.Add(mBase, uint32(v395)+88)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v395)+96)) = v404
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v395)+128)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v395)+108)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v395)+100)) = v432
	v441 = base.I32_div_u_s(v432<<(uint(v415)%32), int32(100))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+112)) = v441
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+140)) = v443
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+144)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v395)+136)) = v446
	v451 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v456 = F_AllocSetContextCreateInternal(m, v451, int32(62088), v404, int32(8192), int32(8388608))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+116)) = v456
	v463 = F_AllocSetContextCreateInternal(m, v456, int32(61951), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L124
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+120)) = v463
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v395)+116))
	v471 = F_AllocSetContextCreateInternal(m, v466, int32(61854), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L124
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+124)) = v471
	v475 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	if int32(2) <= v412 {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	m.G0 = v17 - int32(-64)
	return v395
L130:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v395)+120))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v624
	v628 = F_palloc0(m, v380<<(uint(int32(2))%32))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L124
	} else {
		goto L158
	}
L131:
	;
	v504 = v501 + int32(56)
	v505 = F_BarrierAttach(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L124
	} else {
		goto L140
	}
L132:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v395)+140))
	if v478 != 0 {
		v501 = v478
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v395)+140))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v475
	if v496 == int32(0) {
		goto L130
	} else {
		goto L139
	}
L135:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v395)+116))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v471
	v483 = v412 << (uint(int32(2)) % 32)
	v484 = F_palloc0(m, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L124
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+88)) = v484
	v487 = F_palloc0(m, v483)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L124
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+92)) = v487
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v479
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L124
	} else {
		goto L138
	}
L138:
	;
	goto L134
L139:
	;
	v501 = v496
	goto L131
L140:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	if v507 != 0 {
		goto L129
	} else {
		goto L141
	}
L141:
	;
	v509 = F_BarrierArriveAndWait(m, v504, int32(134217746))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L124
	} else {
		goto L142
	}
L142:
	;
	if v509 == int32(0) {
		goto L129
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+32)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v501)+8)) = v412
	v515 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v501)+20)) = v515
	F_ExecParallelHashJoinSetUpBatches(m, v395, v412)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L124
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+16)) = v380
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v395)+144))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v395)+136))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v395)+140))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+16))
	v529 = F_dsa_allocate_extended(m, v523, v525<<(uint(int32(2))%32), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L124
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v529
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v395)+136))
	v533 = F_dsa_get_address(m, v532, v529)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L124
	} else {
		goto L146
	}
L146:
	;
	if v525 <= int32(0) {
		goto L129
	} else {
		goto L147
	}
L147:
	;
	v538 = v525 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v525) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v544 = v515
	v549 = int32(0)
	goto L151
L149:
	;
	v582 = v515
	goto L150
L150:
	;
	if v538 == int32(0) {
		goto L129
	} else {
		goto L154
	}
L151:
	;
	v560 = v533 + v544<<(uint(int32(2))%32)
	v561 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v560))) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v560)+4)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v560)+8)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v560)+12)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v560)+16)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v560)+20)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v560)+24)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v560)+28)) = v561
	v577 = int32(8)
	v578 = v544 + v577
	v580 = v549 + v577
	if v580 != v525&int32(-8) {
		v544 = v578
		v549 = v580
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v582 = v578
	goto L150
L153:
	;
	goto L152
L154:
	;
	v599 = v582
	v601 = int32(0)
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533+v599<<(uint(int32(2))%32)))) = int32(0)
	v618 = int32(1)
	v621 = v601 + v618
	if v621 != v538 {
		v599 = v599 + v618
		v601 = v621
		goto L155
	} else {
		goto L157
	}
L156:
	;
	goto L129
L157:
	;
	goto L156
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+20)) = v628
	if v412 < int32(2) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v475
	goto L129
L160:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v633 <= int32(0) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v636 == int32(0) {
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v640 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+80)))
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+82)))
	v642 = F_SearchSysCache3(m, int32(65), v636, v640, v641)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L124
	} else {
		goto L163
	}
L163:
	;
	if v642 == int32(0) {
		goto L159
	} else {
		goto L164
	}
L164:
	;
	v651 = F_get_attstatsslot(m, v15+int32(-36), v642, int32(1), int32(0), int32(3))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L124
	} else {
		goto L165
	}
L165:
	;
	if v651 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	if v633 < v653 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	goto L168
L168:
	;
	F_ReleaseCatCache(m, v642)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L124
	} else {
		goto L211
	}
L169:
	;
	F_free_attstatsslot(m, v15+int32(-36))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L124
	} else {
		goto L210
	}
L170:
	;
	v655 = v633
	goto L172
L171:
	;
	v655 = v653
	goto L172
L172:
	;
	if v655 <= int32(0) {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v659 = v655 & int32(3)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	if base.Ui32(v655) < base.Ui32(int32(4)) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v659 != 0 {
		goto L181
	} else {
		goto L182
	}
L175:
	;
	v707 = int32(0)
	v718 = float64(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v669 = int32(0)
	v673 = v669
	v678 = v669
	v684 = float64(0)
	goto L178
L178:
	;
	v687 = v661 + v673<<(uint(int32(2))%32)
	v688 = *(*float32)(unsafe.Add(mBase, uint32(v687)))
	v691 = *(*float32)(unsafe.Add(mBase, uint32(v687)+4))
	v694 = *(*float32)(unsafe.Add(mBase, uint32(v687)+8))
	v697 = *(*float32)(unsafe.Add(mBase, uint32(v687)+12))
	v699 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v684, base.F64_promote_f32(v688)), base.F64_promote_f32(v691)), base.F64_promote_f32(v694)), base.F64_promote_f32(v697))
	v700 = int32(4)
	v701 = v673 + v700
	v703 = v678 + v700
	if v703 != v655&int32(2147483644) {
		v673 = v701
		v678 = v703
		v684 = v699
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v707 = v701
	v718 = v699
	goto L174
L180:
	;
	goto L179
L181:
	;
	v721 = v707
	v724 = int32(0)
	v732 = v718
	goto L184
L182:
	;
	v757 = v718
	goto L183
L183:
	;
	if base.F64_lt(v757, float64(0.01)) != 0 {
		goto L169
	} else {
		goto L187
	}
L184:
	;
	v736 = *(*float32)(unsafe.Add(mBase, uint32(v661+v721<<(uint(int32(2))%32))))
	v738 = base.F64_add(v732, base.F64_promote_f32(v736))
	v739 = int32(1)
	v742 = v724 + v739
	if v742 != v659 {
		v721 = v721 + v739
		v724 = v742
		v732 = v738
		goto L184
	} else {
		goto L186
	}
L185:
	;
	v757 = v738
	goto L183
L186:
	;
	goto L185
L187:
	;
	v760 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+24)) = uint8(v760)
	v765 = v655 + v760
	if v765&v655 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v770 = v760 << (uint(int32(32)-base.I32_clz(v765)) % 32)
	goto L190
L189:
	;
	v770 = v765
	goto L190
L190:
	;
	v772 = v770 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+32)) = v772
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v395)+120))
	v776 = v770 << (uint(int32(4)) % 32)
	v777 = F_MemoryContextAllocZero(m, v774, v776)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L124
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+28)) = v777
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v395)+120))
	v782 = v655 << (uint(int32(2)) % 32)
	v783 = F_MemoryContextAllocZero(m, v780, v782)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L124
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+40)) = v783
	v786 = v776 + v782
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v395)+96))
	v788 = v786 + v787
	*(*int32)(unsafe.Add(mBase, uint32(v395)+96)) = v788
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v395)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+108)) = v790 + v786
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v395)+104))
	if base.Ui32(v793) < base.Ui32(v788) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+104)) = v788
	goto L195
L194:
	;
	goto L195
L195:
	;
	v799 = v772 - int32(1)
	v809 = int32(0)
	goto L196
L196:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v817+v809<<(uint(int32(2))%32))))
	v822 = F_FunctionCall1Coll(m, v815, v816, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L124
	} else {
		goto L198
	}
L197:
	;
	goto L169
L198:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v395+int32(28))))
	v825 = v799 & v822
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v824+v825<<(uint(int32(2))%32))))
	if v829 != 0 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v921 = v809 + int32(1)
	if v921 != v655 {
		v809 = v921
		goto L196
	} else {
		goto L209
	}
L200:
	;
	v832 = v825
	v833 = v829
	goto L203
L201:
	;
	v855 = v825
	goto L202
L202:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v395)+120))
	v869 = F_MemoryContextAlloc(m, v867, int32(8))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L124
	} else {
		goto L207
	}
L203:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v833)))
	if v844 == v822 {
		goto L199
	} else {
		goto L205
	}
L204:
	;
	v855 = v848
	goto L202
L205:
	;
	v848 = (v832 + int32(1)) & v799
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v824+v848<<(uint(int32(2))%32))))
	if v852 != 0 {
		v832 = v848
		v833 = v852
		goto L203
	} else {
		goto L206
	}
L206:
	;
	goto L204
L207:
	;
	v871 = int32(2)
	v872 = v855 << (uint(v871) % 32)
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v395)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v872+v873))) = v869
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v395)+28))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v876+v872)))
	*(*int32)(unsafe.Add(mBase, uint32(v878))) = v822
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v395)+28))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v880+v872)))
	*(*int32)(unsafe.Add(mBase, uint32(v882)+4)) = int32(0)
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v395)+40))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v395)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v885+v886<<(uint(v871)%32)))) = v855
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v395)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+36)) = v891 + int32(1)
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v395)+96))
	v896 = int32(8)
	v897 = v895 + v896
	*(*int32)(unsafe.Add(mBase, uint32(v395)+96)) = v897
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v395)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+108)) = v899 + v896
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v395)+104))
	if base.Ui32(v897) <= base.Ui32(v903) {
		goto L199
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+104)) = v897
	goto L199
L209:
	;
	goto L197
L210:
	;
	goto L168
L211:
	;
	goto L159
}
func F__hash_addovflpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
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
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v733 int32
	_ = v733
	var v742 int32
	_ = v742
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1085 int64
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1105 int64
	_ = v1105
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1135 int32
	_ = v1135
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1149 int32
	_ = v1149
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1177 int32
	_ = v1177
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1191 int32
	_ = v1191
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	F_LockBuffer(m, l2, int32(2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F__hash_checkpage(m, l0, l2, int32(3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) <= l2 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_LockBuffer(m, l1, int32(2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L27
	}
L5:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+16)))
	v56 = v55 + v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v57 == int32(-1) {
		v123 = l2
		v124 = l3
		v132 = v56
		goto L4
	} else {
		goto L9
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v54 = v40 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(l2^int32(-1))<<(uint(int32(2))%32))))
	v54 = v53
	goto L5
L9:
	;
	if l3 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v71 = v57
	goto L16
L11:
	;
	F_UnlockReleaseBuffer(m, l2)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L10
L15:
	;
	goto L10
L16:
	;
	v93 = F__hash_getbuf(m, l0, v71, int32(2), int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+16)))
	v114 = v113 + v112
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v115 == int32(-1) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if v93 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+(v93^int32(-1))<<(uint(int32(2))%32))))
	v112 = v104
	goto L18
L21:
	;
	goto L22
L22:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v112 = v106 + v93<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L23:
	;
	v123 = v93
	v124 = int32(0)
	v132 = v114
	goto L4
L24:
	;
	goto L25
L25:
	;
	F_UnlockReleaseBuffer(m, v93)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v71 = v115
	goto L16
L27:
	;
	F__hash_checkpage(m, l0, l1, int32(8))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if l1 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v170 = v168 + int32(44)
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170))))
	v174 = int32(1)
	v175 = v171<<(uint(int32(3))%32) - v174
	v177 = v168 + int32(76)
	v179 = v168 + int32(60)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v183 = v177 + v180<<(uint(int32(2))%32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v186 = v184 - v174
	v187 = v175 & v186
	v189 = v168 - int32(-64)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v192 = v168 + int32(46)
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192))))
	v194 = int32(base.Ui32(v190) >> (uint(v193) % 32))
	v195 = int32(base.Ui32(v186) >> (uint(v193) % 32))
	if base.Ui32(v194) <= base.Ui32(v195) {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154+(l1^int32(-1))<<(uint(int32(2))%32))))
	v168 = v160
	goto L29
L31:
	;
	goto L32
L32:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v168 = v162 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L29
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L231
	}
L34:
	;
	F_MarkBufferDirty(m, v933)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L147
	}
L35:
	;
	v199 = v175 & v190
	v208 = int32(base.Ui32(v199) >> (uint(int32(5)) % 32))
	v209 = v175
	v210 = v199 & int32(-32)
	v211 = v187
	v217 = v195
	v220 = v194
	goto L38
L36:
	;
	v612 = v175
	v614 = v187
	v615 = v180
	v617 = v183
	v622 = v184
	goto L37
L37:
	;
	if v612 == v614 {
		goto L108
	} else {
		goto L109
	}
L38:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v168+int32(468)+v220<<(uint(int32(2))%32))))
	F_LockBuffer(m, l1, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v612 = v591
	v614 = v599
	v615 = v592
	v617 = v595
	v622 = v596
	goto L37
L40:
	;
	if v217 == v220 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v236 = v211
	goto L43
L42:
	;
	v236 = v209
	goto L43
L43:
	;
	v239 = F__hash_getbuf(m, l0, v231, int32(2), int32(4))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	if base.Ui32(v210) <= base.Ui32(v236) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	if v239 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244+(v239^int32(-1))<<(uint(int32(2))%32))))
	v258 = v250
	goto L44
L47:
	;
	goto L48
L48:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v258 = v252 + v239<<(uint(int32(13))%32) + int32(-8192)
	goto L44
L49:
	;
	v261 = v258 + int32(24)
	v266 = v208
	v268 = v210
	goto L52
L50:
	;
	goto L51
L51:
	;
	F_UnlockReleaseBuffer(m, v239)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L105
	}
L52:
	;
	v288 = v261 + v266<<(uint(int32(2))%32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v289 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L51
L54:
	;
	F_LockBuffer(m, l1, int32(2))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v556 = v268 + int32(32)
	if base.Ui32(v556) <= base.Ui32(v236) {
		v266 = v266 + int32(1)
		v268 = v556
		goto L52
	} else {
		goto L104
	}
L57:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v302 = int32(1)
	v305 = int32(0)
	goto L62
L58:
	;
	v366 = v268 + v365
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v366
	v368 = int32(1)
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+46)))
	v371 = v220<<(uint(v369)%32) + v366
	v373 = v371 + v368
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v168)+60))
	if base.Ui32(v375) < base.Ui32(int32(2)) {
		v414 = v368
		goto L72
	} else {
		goto L73
	}
L59:
	;
	v365 = v305 | int32(1)
	goto L58
L60:
	;
	v365 = v305 | int32(2)
	goto L58
L61:
	;
	v365 = v305 | int32(3)
	goto L58
L62:
	;
	if v302&v295 == int32(0) {
		v365 = v305
		goto L58
	} else {
		goto L64
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L69
	}
L64:
	;
	if v302<<(uint(int32(1))%32)&v295 == int32(0) {
		goto L59
	} else {
		goto L65
	}
L65:
	;
	if v302<<(uint(int32(2))%32)&v295 == int32(0) {
		goto L60
	} else {
		goto L66
	}
L66:
	;
	if v302<<(uint(int32(3))%32)&v295 == int32(0) {
		goto L61
	} else {
		goto L67
	}
L67:
	;
	v340 = int32(4)
	v343 = v305 + v340
	if v343 != int32(32) {
		v302 = v302 << (uint(v340) % 32)
		v305 = v343
		goto L62
	} else {
		goto L68
	}
L68:
	;
	goto L63
L69:
	;
	F_errmsg_internal(m, int32(102515), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(492006), int32(461), int32(102176))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	if base.Ui32(v414) <= base.Ui32(int32(9)) {
		goto L80
	} else {
		goto L81
	}
L73:
	;
	v382 = v368
	goto L74
L74:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v177+v382<<(uint(int32(2))%32))))
	if base.Ui32(v373) <= base.Ui32(v405) {
		v414 = v382
		goto L72
	} else {
		goto L76
	}
L75:
	;
	v414 = v375
	goto L72
L76:
	;
	v408 = v382 + int32(1)
	if v408 != v375 {
		v382 = v408
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v537 = int32(4470804)
	v539 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v540 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v539 + v540
	v547 = v261 + int32(base.Ui32(v366)>>(uint(int32(3))%32))&int32(536870908)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v548 | v540<<(uint(v365)%32)
	v927 = v463
	v929 = int32(0)
	v933 = v239
	v936 = v368
	v938 = v371
	v947 = v239
	goto L34
L79:
	;
	v457 = v456 + v373
	if v457 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v456 = int32(1) << (uint(v414) % 32)
	goto L79
L81:
	;
	goto L82
L82:
	;
	v442 = v414 - int32(10)
	v443 = int32(2)
	v445 = int32(512) << (uint(int32(base.Ui32(v442)>>(uint(v443)%32))) % 32)
	v456 = v445>>(uint(v443)%32)*(v442&int32(3)+int32(1)) + v445
	goto L79
L83:
	;
	v460 = int32(0)
	v463 = F_ReadBufferExtended(m, l0, v460, v457, int32(1), v460)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L101
	}
L86:
	;
	if v482&int32(3) != 0 {
		goto L93
	} else {
		goto L94
	}
L87:
	;
	if v463 < int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v468+(v463^int32(-1))<<(uint(int32(2))%32))))
	v482 = v474
	goto L86
L89:
	;
	goto L90
L90:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v482 = v476 + v463<<(uint(int32(13))%32) + int32(-8192)
	goto L86
L91:
	;
	goto L78
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+10)) = int32(1572864)
	v515 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v482)+18)) = uint16(v515)
	v521 = int32(8176)
	*(*uint16)(unsafe.Add(mBase, uint32(v482)+16)) = uint16(v521)
	*(*uint16)(unsafe.Add(mBase, uint32(v482)+14)) = uint16(v521)
	goto L91
L93:
	;
	v509 = F___memset(m, v482, int32(0), int32(8192))
	mBase = m.M
	goto L92
L94:
	;
	goto L93
L101:
	;
	F_errmsg_internal(m, int32(510452), int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(493850), int32(140), int32(334512))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	goto L53
L105:
	;
	F_LockBuffer(m, l1, int32(2))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170))))
	v590 = int32(1)
	v591 = v587<<(uint(int32(3))%32) - v590
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v595 = v177 + v592<<(uint(int32(2))%32)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)))
	v598 = v596 - v590
	v599 = v591 & v598
	v600 = int32(0)
	v603 = v220 + v590
	v604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192))))
	v605 = int32(base.Ui32(v598) >> (uint(v604) % 32))
	if base.Ui32(v603) <= base.Ui32(v605) {
		v208 = v600
		v209 = v591
		v210 = v600
		v211 = v599
		v217 = v605
		v220 = v603
		goto L38
	} else {
		goto L107
	}
L107:
	;
	goto L39
L108:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v168)+68))
	if base.Ui32(int32(1024)) <= base.Ui32(v633) {
		goto L33
	} else {
		goto L111
	}
L109:
	;
	v733 = int32(0)
	v742 = v622
	goto L110
L110:
	;
	v751 = int32(1)
	v752 = v742 + v751
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v168)+60))
	if base.Ui32(v754) < base.Ui32(int32(2)) {
		v793 = v751
		goto L123
	} else {
		goto L124
	}
L111:
	;
	v636 = int32(1)
	v638 = v622 + v636
	if base.Ui32(v615) < base.Ui32(int32(2)) {
		v677 = v636
		goto L112
	} else {
		goto L113
	}
L112:
	;
	if base.Ui32(v677) <= base.Ui32(int32(9)) {
		goto L119
	} else {
		goto L120
	}
L113:
	;
	v645 = v636
	goto L114
L114:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v177+v645<<(uint(int32(2))%32))))
	if base.Ui32(v638) <= base.Ui32(v668) {
		v677 = v645
		goto L112
	} else {
		goto L116
	}
L115:
	;
	v677 = v615
	goto L112
L116:
	;
	v671 = v645 + int32(1)
	if v671 != v615 {
		v645 = v671
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v721 = F__hash_getnewbuf(m, l0, v718+v638, int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L122
	}
L119:
	;
	v718 = int32(1) << (uint(v677) % 32)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v704 = v677 - int32(10)
	v705 = int32(2)
	v707 = int32(512) << (uint(int32(base.Ui32(v704)>>(uint(v705)%32))) % 32)
	v718 = v707>>(uint(v705)%32)*(v704&int32(3)+int32(1)) + v707
	goto L118
L122:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	v733 = v721
	v742 = v723 + base.B2i32(v721 != int32(0))
	goto L110
L123:
	;
	if base.Ui32(v793) <= base.Ui32(int32(9)) {
		goto L130
	} else {
		goto L131
	}
L124:
	;
	v761 = v751
	goto L125
L125:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v177+v761<<(uint(int32(2))%32))))
	if base.Ui32(v752) <= base.Ui32(v784) {
		v793 = v761
		goto L123
	} else {
		goto L127
	}
L126:
	;
	v793 = v754
	goto L123
L127:
	;
	v787 = v761 + int32(1)
	if v787 != v754 {
		v761 = v787
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v838 = F__hash_getnewbuf(m, l0, v835+v752, int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L133
	}
L130:
	;
	v835 = int32(1) << (uint(v793) % 32)
	goto L129
L131:
	;
	goto L132
L132:
	;
	v821 = v793 - int32(10)
	v822 = int32(2)
	v824 = int32(512) << (uint(int32(base.Ui32(v821)>>(uint(v822)%32))) % 32)
	v835 = v824>>(uint(v822)%32)*(v821&int32(3)+int32(1)) + v824
	goto L129
L133:
	;
	v840 = int32(4470804)
	v842 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v843 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v842 + v843
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	*(*int32)(unsafe.Add(mBase, uint32(v617))) = v846 + v843
	if v733 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v850 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+44)))
	if v733 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	goto L136
L136:
	;
	v927 = v838
	v929 = v733
	v933 = l1
	v936 = int32(0)
	v938 = v742
	v947 = int32(0)
	goto L34
L137:
	;
	v869 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v868)+16)))
	v870 = v869 + v868
	*(*int64)(unsafe.Add(mBase, uint32(v870)+8)) = int64(-36028775544127489)
	*(*int64)(unsafe.Add(mBase, uint32(v870))) = int64(-1)
	v879 = F__emscripten_memset_bulkmem(m, v868+int32(24), base.I32_extend8_s(int32(255)), v850)
	mBase = m.M
	goto L141
L138:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v854+(v733^int32(-1))<<(uint(int32(2))%32))))
	v868 = v860
	goto L137
L139:
	;
	goto L140
L140:
	;
	v862 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v868 = v862 + v733<<(uint(int32(13))%32) + int32(-8192)
	goto L137
L141:
	;
	v881 = v850 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v868)+12)) = uint16(v881)
	F_MarkBufferDirty(m, v733)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	if v733 < int32(0) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v905 = v168 + int32(68)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)))
	*(*int32)(unsafe.Add(mBase, uint32(v168+v906<<(uint(int32(2))%32))+468)) = v903
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v905)))
	v912 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v905))) = v911 + v912
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	*(*int32)(unsafe.Add(mBase, uint32(v617))) = v915 + v912
	goto L136
L144:
	;
	v888 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v888+(v733^int32(-1))<<(uint(int32(6))%32))+16))
	v903 = v894
	goto L143
L145:
	;
	goto L146
L146:
	;
	v896 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v896+v733<<(uint(int32(6))%32)+int32(-64))+16))
	v903 = v902
	goto L143
L147:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v190 == v950 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v938 + int32(1)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	if v927 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L150
L152:
	;
	v975 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v974)+16)))
	if v123 < int32(0) {
		goto L157
	} else {
		goto L158
	}
L153:
	;
	v960 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v960+(v927^int32(-1))<<(uint(int32(2))%32))))
	v974 = v966
	goto L152
L154:
	;
	goto L155
L155:
	;
	v968 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v974 = v968 + v927<<(uint(int32(13))%32) + int32(-8192)
	goto L152
L156:
	;
	v995 = v974 + v975
	*(*int32)(unsafe.Add(mBase, uint32(v995)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v995))) = v994
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v995)+12)) = int32(-8388607)
	*(*int32)(unsafe.Add(mBase, uint32(v995)+8)) = v999
	F_MarkBufferDirty(m, v927)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L160
	}
L157:
	;
	v979 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v979+(v123^int32(-1))<<(uint(int32(6))%32))+16))
	v994 = v985
	goto L156
L158:
	;
	goto L159
L159:
	;
	v987 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v987+v123<<(uint(int32(6))%32)+int32(-64))+16))
	v994 = v993
	goto L156
L160:
	;
	if v927 < int32(0) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v1023
	F_MarkBufferDirty(m, v123)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L165
	}
L162:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1008+(v927^int32(-1))<<(uint(int32(6))%32))+16))
	v1023 = v1014
	goto L161
L163:
	;
	goto L164
L164:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1016+v927<<(uint(int32(6))%32)+int32(-64))+16))
	v1023 = v1022
	goto L161
L165:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027)+118)))
	if v1028 != int32(112) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1198 = int32(4470804)
	v1200 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1200 - int32(1)
	if v124 != 0 {
		goto L217
	} else {
		goto L218
	}
L167:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if v1032 <= int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1035 != 0 {
		goto L166
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+10)) = uint8(v936)
	v1038 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+44)))
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)) = uint16(v1038)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1036 != 0 {
		goto L166
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	F_XLogRegisterData(m, v27+int32(8), int32(3))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_XLogRegisterBuffer(m, int32(0), v927, int32(6))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_XLogRegisterBufData(m, int32(0), v132+int32(8), int32(4))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_XLogRegisterBuffer(m, int32(1), v123, int32(8))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	if v947 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	F_XLogRegisterBuffer(m, int32(2), v947, int32(8))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	if v929 != 0 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	F_XLogRegisterBufData(m, int32(2), v27+int32(12), int32(4))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	F_XLogRegisterBuffer(m, int32(3), v929, int32(6))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L1
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_XLogRegisterBuffer(m, int32(4), l1, int32(8))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L187
	}
L186:
	;
	goto L185
L187:
	;
	v1079 = int32(4)
	F_XLogRegisterBufData(m, v1079, v189, v1079)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v1085 = F_XLogInsert(m, int32(12), int32(48))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	if v927 < int32(0) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v1105 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v1104))) = base.I64_rotr(v1085, v1105)
	v1110 = base.I32_wrap_i64(int64(base.Ui64(v1085) >> (uint(v1105) % 64)))
	v1111 = base.I32_wrap_i64(v1085)
	if v123 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1090+(v927^int32(-1))<<(uint(int32(2))%32))))
	v1104 = v1096
	goto L190
L192:
	;
	goto L193
L193:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1104 = v1098 + v927<<(uint(int32(13))%32) + int32(-8192)
	goto L190
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1129)+4)) = v1111
	*(*int32)(unsafe.Add(mBase, uint32(v1129))) = v1110
	if v947 != 0 {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1115+(v123^int32(-1))<<(uint(int32(2))%32))))
	v1129 = v1121
	goto L194
L196:
	;
	goto L197
L197:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1129 = v1123 + v123<<(uint(int32(13))%32) + int32(-8192)
	goto L194
L198:
	;
	if v947 < int32(0) {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	goto L200
L200:
	;
	if v929 != 0 {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+4)) = v1111
	*(*int32)(unsafe.Add(mBase, uint32(v1149))) = v1110
	goto L200
L202:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1135+(v947^int32(-1))<<(uint(int32(2))%32))))
	v1149 = v1141
	goto L201
L203:
	;
	goto L204
L204:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1149 = v1143 + v947<<(uint(int32(13))%32) + int32(-8192)
	goto L201
L205:
	;
	if v929 < int32(0) {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	goto L207
L207:
	;
	if l1 < int32(0) {
		goto L213
	} else {
		goto L214
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1170)+4)) = v1111
	*(*int32)(unsafe.Add(mBase, uint32(v1170))) = v1110
	goto L207
L209:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1156+(v929^int32(-1))<<(uint(int32(2))%32))))
	v1170 = v1162
	goto L208
L210:
	;
	goto L211
L211:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1170 = v1164 + v929<<(uint(int32(13))%32) + int32(-8192)
	goto L208
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1191)+4)) = v1111
	*(*int32)(unsafe.Add(mBase, uint32(v1191))) = v1110
	goto L166
L213:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1177+(l1^int32(-1))<<(uint(int32(2))%32))))
	v1191 = v1183
	goto L212
L214:
	;
	goto L215
L215:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1191 = v1185 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L212
L216:
	;
	if v947 != 0 {
		goto L222
	} else {
		goto L223
	}
L217:
	;
	F_LockBuffer(m, v123, int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L1
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	F_UnlockReleaseBuffer(m, v123)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L1
	} else {
		goto L221
	}
L220:
	;
	goto L216
L221:
	;
	goto L216
L222:
	;
	F_UnlockReleaseBuffer(m, v947)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	F_LockBuffer(m, l1, int32(0))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L1
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	if v929 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	F_UnlockReleaseBuffer(m, v929)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	m.G0 = v27 + int32(16)
	return v927
L230:
	;
	goto L229
L231:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v1227 + int32(4)
	F_errmsg(m, int32(673095), v27)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(492006), int32(285), int32(402614))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__hash_binsearch(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	v8 = int32(1)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v10) < base.Ui32(int32(25)) {
		v19 = v8
	} else {
		v19 = int32(base.Ui32(v10+int32(262120))>>(uint(int32(2))%32)) + v8
	}
	if base.Ui32(int32(2)) <= base.Ui32(v19&int32(65535)) {
		v28 = v8
		v29 = v19
		for {
			v33 = int32(65535)
			v39 = int32(base.Ui32(v29&v33+v28&v33) >> (uint(int32(1)) % 32))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v39<<(uint(int32(2))%32)+(l0+int32(24))-int32(4))))
			v48 = l0 + v45&int32(32767)
			v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+6)))
			if int32(0) <= v51 {
				v54 = int32(8)
			} else {
				v54 = int32(16)
			}
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v48+v54)))
			v57 = base.B2i32(base.Ui32(v56) < base.Ui32(l1))
			if base.Ui32(v56) < base.Ui32(l1) {
				v58 = v29
			} else {
				v58 = v39
			}
			if base.Ui32(v56) < base.Ui32(l1) {
				v63 = v39 + int32(1)
			} else {
				v63 = v28
			}
			if base.Ui32(v63&int32(65535)) < base.Ui32(v58&int32(65535)) {
				v28 = v63
				v29 = v58
				continue
			} else {
				break
			}
			break
		}
		v69 = v63
	} else {
		v69 = v8
	}
	return v69 & int32(65535)
}
func F__hash_checkpage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	if l1 < int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+(l1^int32(-1))<<(uint(int32(2))%32))))
		v28 = v20
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[1]))
		v28 = v22 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+14)))
	if v29 != 0 {
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+19)))
		v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
		if (v30<<(uint(int32(8))%32)-v33)&int32(65535) != int32(16) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if l1 < int32(0) {
						v111 = *(*int32)(unsafe.Add(mBase, _consts[8]))
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v111+(l1^int32(-1))<<(uint(int32(6))%32))+16))
						v126 = v117
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v119+l1<<(uint(int32(6))%32)+int32(-64))+16))
						v126 = v125
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v126
					*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v107 + int32(4)
					F_errmsg(m, int32(47471), v9-int32(-64))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return
					} else {
						F_errhint(m, int32(560299), int32(0))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(491967), int32(237), int32(402660))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
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
			if l2 == int32(0) {
				m.G0 = v9 + int32(80)
				return
			} else {
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+v33)+12)))
				if l2&v42 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return
					} else {
						F_errcode(m, int32(33557032))
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return
						} else {
							v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							if l1 < int32(0) {
								v156 = *(*int32)(unsafe.Add(mBase, _consts[8]))
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v156+(l1^int32(-1))<<(uint(int32(6))%32))+16))
								v171 = v162
							} else {
								v164 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								v170 = *(*int32)(unsafe.Add(mBase, uint32(v164+l1<<(uint(int32(6))%32)+int32(-64))+16))
								v171 = v170
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v171
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v152 + int32(4)
							F_errmsg(m, int32(47471), v9+int32(16))
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return
							} else {
								F_errhint(m, int32(560299), int32(0))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return
								} else {
									F_errfinish(m, int32(491967), int32(249), int32(402660))
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
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
					if l2 != int32(8) {
						m.G0 = v9 + int32(80)
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
						if v48 != int32(105121344) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return
							} else {
								F_errcode(m, int32(33557032))
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return
								} else {
									v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v197 + int32(4)
									F_errmsg(m, int32(28391), v9+int32(48))
									mBase = m.M
									v205 = m.ExcPending
									if v205 != 0 {
										return
									} else {
										F_errfinish(m, int32(491967), int32(263), int32(402660))
										mBase = m.M
										v210 = m.ExcPending
										if v210 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
							if v51 != int32(4) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v214 = m.ExcPending
								if v214 != 0 {
									return
								} else {
									F_errcode(m, int32(33557032))
									mBase = m.M
									v217 = m.ExcPending
									if v217 != 0 {
										return
									} else {
										v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v218 + int32(4)
										F_errmsg(m, int32(268674), v9+int32(32))
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return
										} else {
											F_errhint(m, int32(560299), int32(0))
											mBase = m.M
											v230 = m.ExcPending
											if v230 != 0 {
												return
											} else {
												F_errfinish(m, int32(491967), int32(270), int32(402660))
												mBase = m.M
												v235 = m.ExcPending
												if v235 != 0 {
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
								m.G0 = v9 + int32(80)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if l1 < int32(0) {
					v68 = *(*int32)(unsafe.Add(mBase, _consts[8]))
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v68+(l1^int32(-1))<<(uint(int32(6))%32))+16))
					v83 = v74
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+l1<<(uint(int32(6))%32)+int32(-64))+16))
					v83 = v82
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v64 + int32(4)
				F_errmsg(m, int32(47418), v9)
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					F_errhint(m, int32(560299), int32(0))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						F_errfinish(m, int32(491967), int32(226), int32(402660))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
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
func F__hash_get_indextuple_hashkey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if int32(0) <= v4 {
		v7 = int32(8)
	} else {
		v7 = int32(16)
	}
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0+v7)))
	return v9
}
func F__hash_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	if l1 != int32(-1) {
		v7 = F_ReadBuffer(m, l0, l1)
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if l2 != int32(-1) {
				F_LockBuffer(m, v7, l2)
				v14 = m.ExcPending
				if v14 != 0 {
					return int32(0)
				} else {
					F__hash_checkpage(m, l0, v7, l3)
					v16 = m.ExcPending
					if v16 != 0 {
						return int32(0)
					} else {
						return v7
					}
				}
			} else {
				F__hash_checkpage(m, l0, v7, l3)
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					return v7
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(510452), int32(0))
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493850), int32(75), int32(334529))
				v30 = m.ExcPending
				if v30 != 0 {
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
func F__hash_getnewbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
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
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = F_RelationGetNumberOfBlocksInFork(m, l0, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if l1 != int32(-1) {
			if base.Ui32(v10) < base.Ui32(l1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int32(0)
				} else {
					v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v171 + int32(4)
					F_errmsg_internal(m, int32(673136), v8)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493850), int32(207), int32(334485))
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if l1 == v10 {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = l0
					v23 = *(*int64)(unsafe.Add(mBase, uint32(v8)+36))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v23
					v29 = F_ExtendBufferedRel(m, v8+int32(24), l2, int32(0), int32(9))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 < int32(0) {
							v34 = *(*int32)(unsafe.Add(mBase, _consts[8]))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v34+(v29^int32(-1))<<(uint(int32(6))%32))+16))
							v49 = v40
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+v29<<(uint(int32(6))%32)+int32(-64))+16))
							v49 = v48
						}
						if v49 == l1 {
							v90 = v29
							if v90 < int32(0) {
								v94 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v94+(v90^int32(-1))<<(uint(int32(2))%32))))
								v108 = v100
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, _consts[1]))
								v108 = v102 + v90<<(uint(int32(13))%32) + int32(-8192)
							}
							if v108&int32(3) != 0 {
							} else {
							}
							v135 = F___memset(m, v108, int32(0), int32(8192))
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v108)+10)) = int32(1572864)
							v141 = int32(8196)
							*(*uint16)(unsafe.Add(mBase, uint32(v108)+18)) = uint16(v141)
							v147 = int32(8176)
							*(*uint16)(unsafe.Add(mBase, uint32(v108)+16)) = uint16(v147)
							*(*uint16)(unsafe.Add(mBase, uint32(v108)+14)) = uint16(v147)
							m.G0 = v8 + int32(48)
							return v90
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								if v29 < int32(0) {
									v58 = *(*int32)(unsafe.Add(mBase, _consts[8]))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+(v29^int32(-1))<<(uint(int32(6))%32))+16))
									v73 = v64
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, _consts[9]))
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v66+v29<<(uint(int32(6))%32)+int32(-64))+16))
									v73 = v72
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v73
								F_errmsg_internal(m, int32(52816), v8+int32(16))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493850), int32(216), int32(334485))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
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
					v88 = F_ReadBufferExtended(m, l0, l2, l1, int32(1), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						v90 = v88
						if v90 < int32(0) {
							v94 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v94+(v90^int32(-1))<<(uint(int32(2))%32))))
							v108 = v100
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, _consts[1]))
							v108 = v102 + v90<<(uint(int32(13))%32) + int32(-8192)
						}
						if v108&int32(3) != 0 {
						} else {
						}
						v135 = F___memset(m, v108, int32(0), int32(8192))
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v108)+10)) = int32(1572864)
						v141 = int32(8196)
						*(*uint16)(unsafe.Add(mBase, uint32(v108)+18)) = uint16(v141)
						v147 = int32(8176)
						*(*uint16)(unsafe.Add(mBase, uint32(v108)+16)) = uint16(v147)
						*(*uint16)(unsafe.Add(mBase, uint32(v108)+14)) = uint16(v147)
						m.G0 = v8 + int32(48)
						return v90
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(510452), int32(0))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493850), int32(204), int32(334485))
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
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
func F__hash_kill_items(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v163 int32
	_ = v163
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	v2 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v2
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v17 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	if v29 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_LockBuffer(m, v20, int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v25 = int32(1)
	v27 = F__hash_getbuf(m, v15, v24, v25, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	v29 = v20
	goto L2
L8:
	;
	v29 = v27
	goto L2
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+(v29^int32(-1))<<(uint(int32(2))%32))))
	v47 = v39
	goto L1
L10:
	;
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v47 = v41 + v29<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L12:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v20 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L13:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+16)))
	v51 = v47 + v50
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v56) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v64 = int32(base.Ui32(v56+int32(262120)) >> (uint(int32(2)) % 32))
	goto L16
L15:
	;
	v64 = int32(0)
	goto L16
L16:
	;
	v66 = v64 & int32(65535)
	v71 = v2
	v80 = v2
	goto L18
L17:
	;
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+12)))
	v183 = v181 | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+12)) = uint16(v183)
	F_MarkBufferDirtyHint(m, v29, int32(1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L6
	} else {
		goto L37
	}
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v71<<(uint(int32(2))%32))))
	v88 = v16 + int32(52) + v85<<(uint(int32(3))%32)
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+6)))
	if base.Ui32(v66) < base.Ui32(v89) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v80 == int32(0) {
		goto L12
	} else {
		goto L36
	}
L20:
	;
	v163 = v71 + int32(1)
	if v163 != v17 {
		v71 = v163
		goto L18
	} else {
		goto L35
	}
L21:
	;
	v91 = v89
	goto L22
L22:
	;
	v111 = v91&int32(65535)<<(uint(int32(2))%32) + (v47 + int32(24)) - int32(4)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v115 = v47 + v112&int32(32767)
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+2)))
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115))))
	v118 = int32(16)
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+2)))
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88))))
	if v116|v117<<(uint(v118)%32) == v121|v122<<(uint(v118)%32) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v140 | int32(98304)
	v144 = int32(1)
	v146 = v71 + v144
	if v146 != v17 {
		v71 = v146
		v80 = v144
		goto L18
	} else {
		goto L34
	}
L24:
	;
	if v132 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	goto L24
L26:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+4)))
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)))
	if v128 == v129 {
		v132 = int32(1)
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v132 = int32(0)
	goto L25
L29:
	;
	goto L28
L30:
	;
	v136 = v91 + int32(1)
	if base.Ui32(v136&int32(65535)) <= base.Ui32(v66) {
		v91 = v136
		goto L22
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L23
L33:
	;
	goto L20
L34:
	;
	goto L17
L35:
	;
	goto L19
L36:
	;
	goto L17
L37:
	;
	goto L12
L38:
	;
	F_UnlockReleaseBuffer(m, v29)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L44
	}
L39:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v205 != v202 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_LockBuffer(m, v202, int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	return
L44:
	;
	return
}
func F__hash_next(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if l1 == int32(1) {
		v17 = v13 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v17
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
		if v17 <= v19 {
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
			v91 = v12 + v88<<(uint(int32(3))%32)
			v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
			v100 = int32(1)
			m.G0 = v9 + int32(16)
			return v100
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			if int32(0) < v21 {
				F__hash_kill_items(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
					if v28 == int32(-1) {
						F__hash_dropscanbuf(m, v12)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v75 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
							*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
							v100 = v75
							m.G0 = v9 + int32(16)
							return v100
						}
					} else {
						v31 = int32(1)
						v33 = F__hash_getbuf(m, v11, v28, v31, v31)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33
							v39 = F__hash_readpage(m, l0, v9+int32(12), int32(1))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								if v39 == int32(0) {
									F__hash_dropscanbuf(m, v12)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v75 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
										*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
										*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
										v100 = v75
										m.G0 = v9 + int32(16)
										return v100
									}
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
									v91 = v12 + v88<<(uint(int32(3))%32)
									v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
									*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
									v100 = int32(1)
									m.G0 = v9 + int32(16)
									return v100
								}
							}
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
				if v28 == int32(-1) {
					F__hash_dropscanbuf(m, v12)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v75 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
						*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
						v100 = v75
						m.G0 = v9 + int32(16)
						return v100
					}
				} else {
					v31 = int32(1)
					v33 = F__hash_getbuf(m, v11, v28, v31, v31)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33
						v39 = F__hash_readpage(m, l0, v9+int32(12), int32(1))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 == int32(0) {
								F__hash_dropscanbuf(m, v12)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v75 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
									*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
									*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
									v100 = v75
									m.G0 = v9 + int32(16)
									return v100
								}
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
								v91 = v12 + v88<<(uint(int32(3))%32)
								v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
								*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
								v100 = int32(1)
								m.G0 = v9 + int32(16)
								return v100
							}
						}
					}
				}
			}
		}
	} else {
		v44 = v13 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v44
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
		if v46 <= v44 {
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
			v91 = v12 + v88<<(uint(int32(3))%32)
			v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
			v100 = int32(1)
			m.G0 = v9 + int32(16)
			return v100
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			if int32(0) < v48 {
				F__hash_kill_items(m, l0)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
					if v53 == int32(-1) {
						F__hash_dropscanbuf(m, v12)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v75 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
							*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
							v100 = v75
							m.G0 = v9 + int32(16)
							return v100
						}
					} else {
						v58 = F__hash_getbuf(m, v11, v53, int32(1), int32(3))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v58
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							if v61 != v58 {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if v58 != v63 {
									v69 = F__hash_readpage(m, l0, v9+int32(12), l1)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										if v69 != 0 {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
											v91 = v12 + v88<<(uint(int32(3))%32)
											v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
											v100 = int32(1)
											m.G0 = v9 + int32(16)
											return v100
										} else {
											F__hash_dropscanbuf(m, v12)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
												*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
												*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
												v100 = v75
												m.G0 = v9 + int32(16)
												return v100
											}
										}
									}
								} else {
									F_ReleaseBuffer(m, v58)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v69 = F__hash_readpage(m, l0, v9+int32(12), l1)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											if v69 != 0 {
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
												v91 = v12 + v88<<(uint(int32(3))%32)
												v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
												*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
												v100 = int32(1)
												m.G0 = v9 + int32(16)
												return v100
											} else {
												F__hash_dropscanbuf(m, v12)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													v75 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
													*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
													*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
													v100 = v75
													m.G0 = v9 + int32(16)
													return v100
												}
											}
										}
									}
								}
							} else {
								F_ReleaseBuffer(m, v58)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v69 = F__hash_readpage(m, l0, v9+int32(12), l1)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										if v69 != 0 {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
											v91 = v12 + v88<<(uint(int32(3))%32)
											v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
											v100 = int32(1)
											m.G0 = v9 + int32(16)
											return v100
										} else {
											F__hash_dropscanbuf(m, v12)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
												*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
												*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
												v100 = v75
												m.G0 = v9 + int32(16)
												return v100
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
				if v53 == int32(-1) {
					F__hash_dropscanbuf(m, v12)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v75 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
						*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
						v100 = v75
						m.G0 = v9 + int32(16)
						return v100
					}
				} else {
					v58 = F__hash_getbuf(m, v11, v53, int32(1), int32(3))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v58
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						if v61 != v58 {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v58 != v63 {
								v69 = F__hash_readpage(m, l0, v9+int32(12), l1)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									if v69 != 0 {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
										v91 = v12 + v88<<(uint(int32(3))%32)
										v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
										*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
										v100 = int32(1)
										m.G0 = v9 + int32(16)
										return v100
									} else {
										F__hash_dropscanbuf(m, v12)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											v75 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
											*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
											*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
											v100 = v75
											m.G0 = v9 + int32(16)
											return v100
										}
									}
								}
							} else {
								F_ReleaseBuffer(m, v58)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v69 = F__hash_readpage(m, l0, v9+int32(12), l1)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										if v69 != 0 {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
											v91 = v12 + v88<<(uint(int32(3))%32)
											v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
											v100 = int32(1)
											m.G0 = v9 + int32(16)
											return v100
										} else {
											F__hash_dropscanbuf(m, v12)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
												*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
												*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
												v100 = v75
												m.G0 = v9 + int32(16)
												return v100
											}
										}
									}
								}
							}
						} else {
							F_ReleaseBuffer(m, v58)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v69 = F__hash_readpage(m, l0, v9+int32(12), l1)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									if v69 != 0 {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
										v91 = v12 + v88<<(uint(int32(3))%32)
										v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+56)))
										*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v92)
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v94
										v100 = int32(1)
										m.G0 = v9 + int32(16)
										return v100
									} else {
										F__hash_dropscanbuf(m, v12)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											v75 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v75
											*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(-1)
											*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(-4294967296)
											v100 = v75
											m.G0 = v9 + int32(16)
											return v100
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
func F__hash_readpage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v612 int32
	_ = v612
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v839 int32
	_ = v839
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v904 int32
	_ = v904
	var v915 int32
	_ = v915
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1079 int32
	_ = v1079
	var v1097 int32
	_ = v1097
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1174 int32
	_ = v1174
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1194 int32
	_ = v1194
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1224 int32
	_ = v1224
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1239 int32
	_ = v1239
	var v1250 int32
	_ = v1250
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1414 int32
	_ = v1414
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1445 int32
	_ = v1445
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1513 int32
	_ = v1513
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v17
	F__hash_checkpage(m, v16, v17, int32(3))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v17 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v41
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+16)))
	v44 = v41 + v43
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v17
	if v17 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(v17^int32(-1))<<(uint(int32(2))%32))))
	v41 = v33
	goto L3
L5:
	;
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v41 = v35 + v17<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if l2 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50+(v17^int32(-1))<<(uint(int32(6))%32))+16))
	v65 = v56
	goto L7
L9:
	;
	goto L10
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+v17<<(uint(int32(6))%32)+int32(-64))+16))
	v65 = v64
	goto L7
L11:
	;
	m.G0 = v13 + int32(16)
	return v1513
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v1466
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v1467
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v1478 == v1479 {
		goto L343
	} else {
		goto L344
	}
L13:
	;
	v1463 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v1463
	v1513 = v1463
	goto L11
L14:
	;
	v70 = int32(0)
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
	if base.Ui32(v75) < base.Ui32(int32(25)) {
		v133 = v70
		goto L19
	} else {
		goto L20
	}
L15:
	;
	goto L16
L16:
	;
	v828 = int32(1)
	v830 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
	if base.Ui32(v830) < base.Ui32(int32(25)) {
		goto L198
	} else {
		goto L199
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(-1)
	goto L13
L18:
	;
	v140 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l2 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v139 = v133 & int32(65535)
	goto L18
L20:
	;
	v81 = int32(base.Ui32(v75+int32(262120)) >> (uint(int32(2)) % 32))
	if v81&int32(65535) == int32(0) {
		v133 = v70
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v90 = v70
	v91 = v81
	goto L22
L22:
	;
	v95 = int32(65535)
	v100 = int32(1)
	v103 = int32(base.Ui32(v90&v95+v91&v95+v100) >> (uint(v100) % 32))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v103<<(uint(int32(2))%32)+(v41+int32(24))-int32(4))))
	v114 = v41 + v111&int32(32767)
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114)+6)))
	if int32(0) <= v117 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v133 = v127
	goto L19
L24:
	;
	v120 = int32(8)
	goto L26
L25:
	;
	v120 = int32(16)
	goto L26
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v114+v120)))
	v123 = base.B2i32(base.Ui32(v67) < base.Ui32(v122))
	if base.Ui32(v67) < base.Ui32(v122) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v124 = v103 - v100
	goto L29
L28:
	;
	v124 = v91
	goto L29
L29:
	;
	if base.Ui32(v67) < base.Ui32(v122) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v127 = v90
	goto L32
L31:
	;
	v127 = v103
	goto L32
L32:
	;
	if base.Ui32(v127&int32(65535)) < base.Ui32(v124&int32(65535)) {
		v90 = v127
		v91 = v124
		goto L22
	} else {
		goto L33
	}
L33:
	;
	goto L23
L34:
	;
	v340 = v338 & int32(65535)
	if v340 == int32(408) {
		goto L77
	} else {
		goto L78
	}
L35:
	;
	v338 = v322
	goto L34
L36:
	;
	v248 = v139
	v252 = int32(408)
	goto L62
L37:
	;
	if v139 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v158) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v338 = int32(408)
	goto L34
L41:
	;
	goto L42
L42:
	;
	goto L36
L43:
	;
	v166 = int32(base.Ui32(v158+int32(262120)) >> (uint(int32(2)) % 32))
	goto L45
L44:
	;
	v166 = int32(0)
	goto L45
L45:
	;
	v168 = v166 & int32(65535)
	if base.Ui32(v168) < base.Ui32(v139) {
		v322 = v140
		goto L35
	} else {
		goto L46
	}
L46:
	;
	v176 = v139
	v180 = v140
	goto L47
L47:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+12)))
	v190 = v176
	goto L49
L48:
	;
	v322 = v240
	goto L35
L49:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v190&int32(65535)<<(uint(int32(2))%32)+(v41+int32(24))-int32(4))))
	v209 = v41 + v206&int32(32767)
	if v185&int32(1) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v229 = F__hash_get_indextuple_hashkey(m, v209)
	mBase = m.M
	if v228 != v229 {
		v322 = v180
		goto L35
	} else {
		goto L60
	}
L51:
	;
	goto L50
L52:
	;
	v224 = v190 + int32(1)
	if base.Ui32(v224&int32(65535)) <= base.Ui32(v168) {
		v190 = v224
		goto L49
	} else {
		goto L59
	}
L53:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	if v216 != int32(1) {
		goto L51
	} else {
		goto L57
	}
L54:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+13)))
	if v212 != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+7)))
	if v213&int32(32) != 0 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v219 = int32(98304)
	if v206&v219 != v219 {
		goto L51
	} else {
		goto L58
	}
L58:
	;
	goto L52
L59:
	;
	v322 = v180
	goto L35
L60:
	;
	v233 = v147 + int32(52) + v180<<(uint(int32(3))%32)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v234
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v233)+6)) = uint16(v190)
	*(*uint16)(unsafe.Add(mBase, uint32(v233)+4)) = uint16(v236)
	v239 = int32(1)
	v240 = v180 + v239
	v242 = v190 + v239
	if base.Ui32(v242&int32(65535)) <= base.Ui32(v168) {
		v176 = v242
		v180 = v240
		goto L47
	} else {
		goto L61
	}
L61:
	;
	goto L48
L62:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+12)))
	v262 = v248
	goto L64
L63:
	;
	v322 = v303
	goto L35
L64:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v262&int32(65535)<<(uint(int32(2))%32)+(v41+int32(24))-int32(4))))
	v281 = v41 + v278&int32(32767)
	if v257&int32(1) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v300 = F__hash_get_indextuple_hashkey(m, v281)
	mBase = m.M
	if v299 != v300 {
		v322 = v252
		goto L35
	} else {
		goto L75
	}
L66:
	;
	goto L65
L67:
	;
	v296 = v262 - int32(1)
	if v296&int32(65535) != 0 {
		v262 = v296
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	if v288 != int32(1) {
		goto L66
	} else {
		goto L72
	}
L69:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+13)))
	if v284 != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+7)))
	if v285&int32(32) != 0 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	v291 = int32(98304)
	if v278&v291 != v291 {
		goto L66
	} else {
		goto L73
	}
L73:
	;
	goto L67
L74:
	;
	v322 = v252
	goto L35
L75:
	;
	v302 = int32(1)
	v303 = v252 - v302
	v306 = v147 + int32(52) + v303<<(uint(int32(3))%32)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v307
	v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+6)) = uint16(v262)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+4)) = uint16(v309)
	v313 = v262 - v302
	if v313&int32(65535) != 0 {
		v248 = v313
		v252 = v303
		goto L62
	} else {
		goto L76
	}
L76:
	;
	goto L63
L77:
	;
	v345 = v44
	v348 = v17
	v352 = int32(-1)
	goto L80
L78:
	;
	v803 = v340
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v803
	v808 = int32(407)
	v1466 = v808
	v1467 = v808
	goto L12
L80:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if int32(0) < v354 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v803 = v794
	goto L79
L82:
	;
	F__hash_kill_items(m, l0)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v359 != v360 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L84
L86:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	if v369 != v348 {
		goto L93
	} else {
		goto L94
	}
L87:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v359 != v362 {
		v365 = v352
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v365 = v364
	goto L86
L90:
	;
	goto L89
L91:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v382 != 0 {
		goto L99
	} else {
		goto L100
	}
L92:
	;
	F_UnlockReleaseBuffer(m, v348)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L98
	}
L93:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	if v348 != v371 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v373 = int32(0)
	F_LockBuffer(m, v348, v373)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	v380 = v373
	goto L91
L98:
	;
	v380 = int32(1)
	goto L91
L99:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v380 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L101
L103:
	;
	if v493 == int32(0) {
		goto L17
	} else {
		goto L131
	}
L104:
	;
	v387 = F__hash_getbuf(m, v367, v366, int32(1), int32(3))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+12)))
	if v418 != int32(1) {
		goto L17
	} else {
		goto L117
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v387
	if v387 < int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v407
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v407)+16)))
	v410 = v407 + v409
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v410
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	if v412 != v387 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v393+(v387^int32(-1))<<(uint(int32(2))%32))))
	v407 = v399
	goto L108
L110:
	;
	goto L111
L111:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v407 = v401 + v387<<(uint(int32(13))%32) + int32(-8192)
	goto L108
L112:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	if v387 != v414 {
		v490 = v410
		v493 = v387
		goto L103
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_ReleaseBuffer(m, v387)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	v490 = v410
	v493 = v387
	goto L103
L117:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+13)))
	if v421 != int32(1) {
		goto L17
	} else {
		goto L118
	}
L118:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v424
	F_LockBuffer(m, v424, int32(1))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v424 < int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v446
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v446)+16)))
	v449 = v446 + v448
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v451 != int32(-1) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v432+(v424^int32(-1))<<(uint(int32(2))%32))))
	v446 = v438
	goto L120
L122:
	;
	goto L123
L123:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v446 = v440 + v424<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L124:
	;
	goto L127
L125:
	;
	v478 = v449
	v481 = v424
	goto L126
L126:
	;
	v487 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v368)+13)) = uint8(v487)
	v490 = v478
	v493 = v481
	goto L103
L127:
	;
	F__hash_readnext(m, l0, v13+int32(12), v13+int32(8), v13+int32(4))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v478 = v472
	v481 = v476
	goto L126
L129:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	if v473 != int32(-1) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v493
	if v493 < int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v520
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v524 = int32(0)
	v529 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v522)+12)))
	if base.Ui32(v529) < base.Ui32(int32(25)) {
		v587 = v524
		goto L137
	} else {
		goto L138
	}
L133:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v505+(v493^int32(-1))<<(uint(int32(6))%32))+16))
	v520 = v511
	goto L132
L134:
	;
	goto L135
L135:
	;
	v513 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v513+v493<<(uint(int32(6))%32)+int32(-64))+16))
	v520 = v519
	goto L132
L136:
	;
	v594 = int32(0)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l2 != int32(1) {
		goto L155
	} else {
		goto L156
	}
L137:
	;
	v593 = v587 & int32(65535)
	goto L136
L138:
	;
	v535 = int32(base.Ui32(v529+int32(262120)) >> (uint(int32(2)) % 32))
	if v535&int32(65535) == int32(0) {
		v587 = v524
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v544 = v524
	v545 = v535
	goto L140
L140:
	;
	v549 = int32(65535)
	v554 = int32(1)
	v557 = int32(base.Ui32(v544&v549+v545&v549+v554) >> (uint(v554) % 32))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v557<<(uint(int32(2))%32)+(v522+int32(24))-int32(4))))
	v568 = v522 + v565&int32(32767)
	v571 = int32(*(*int16)(unsafe.Add(mBase, uint32(v568)+6)))
	if int32(0) <= v571 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v587 = v581
	goto L137
L142:
	;
	v574 = int32(8)
	goto L144
L143:
	;
	v574 = int32(16)
	goto L144
L144:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v568+v574)))
	v577 = base.B2i32(base.Ui32(v523) < base.Ui32(v576))
	if base.Ui32(v523) < base.Ui32(v576) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v578 = v557 - v554
	goto L147
L146:
	;
	v578 = v545
	goto L147
L147:
	;
	if base.Ui32(v523) < base.Ui32(v576) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v581 = v544
	goto L150
L149:
	;
	v581 = v557
	goto L150
L150:
	;
	if base.Ui32(v581&int32(65535)) < base.Ui32(v578&int32(65535)) {
		v544 = v581
		v545 = v578
		goto L140
	} else {
		goto L151
	}
L151:
	;
	goto L141
L152:
	;
	v794 = v792 & int32(65535)
	if v794 == int32(408) {
		v345 = v490
		v348 = v493
		v352 = v365
		goto L80
	} else {
		goto L195
	}
L153:
	;
	v792 = v776
	goto L152
L154:
	;
	v702 = v593
	v706 = int32(408)
	goto L180
L155:
	;
	if v593 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	goto L157
L157:
	;
	v612 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v522)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v612) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	v792 = int32(408)
	goto L152
L159:
	;
	goto L160
L160:
	;
	goto L154
L161:
	;
	v620 = int32(base.Ui32(v612+int32(262120)) >> (uint(int32(2)) % 32))
	goto L163
L162:
	;
	v620 = int32(0)
	goto L163
L163:
	;
	v622 = v620 & int32(65535)
	if base.Ui32(v622) < base.Ui32(v593) {
		v776 = v594
		goto L153
	} else {
		goto L164
	}
L164:
	;
	v630 = v593
	v634 = v594
	goto L165
L165:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+12)))
	v644 = v630
	goto L167
L166:
	;
	v776 = v694
	goto L153
L167:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v644&int32(65535)<<(uint(int32(2))%32)+(v522+int32(24))-int32(4))))
	v663 = v522 + v660&int32(32767)
	if v639&int32(1) == int32(0) {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
	v683 = F__hash_get_indextuple_hashkey(m, v663)
	mBase = m.M
	if v682 != v683 {
		v776 = v634
		goto L153
	} else {
		goto L178
	}
L169:
	;
	goto L168
L170:
	;
	v678 = v644 + int32(1)
	if base.Ui32(v678&int32(65535)) <= base.Ui32(v622) {
		v644 = v678
		goto L167
	} else {
		goto L177
	}
L171:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	if v670 != int32(1) {
		goto L169
	} else {
		goto L175
	}
L172:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+13)))
	if v666 != 0 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+7)))
	if v667&int32(32) != 0 {
		goto L170
	} else {
		goto L174
	}
L174:
	;
	goto L171
L175:
	;
	v673 = int32(98304)
	if v660&v673 != v673 {
		goto L169
	} else {
		goto L176
	}
L176:
	;
	goto L170
L177:
	;
	v776 = v634
	goto L153
L178:
	;
	v687 = v601 + int32(52) + v634<<(uint(int32(3))%32)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	*(*int32)(unsafe.Add(mBase, uint32(v687))) = v688
	v690 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v663)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v687)+6)) = uint16(v644)
	*(*uint16)(unsafe.Add(mBase, uint32(v687)+4)) = uint16(v690)
	v693 = int32(1)
	v694 = v634 + v693
	v696 = v644 + v693
	if base.Ui32(v696&int32(65535)) <= base.Ui32(v622) {
		v630 = v696
		v634 = v694
		goto L165
	} else {
		goto L179
	}
L179:
	;
	goto L166
L180:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+12)))
	v716 = v702
	goto L182
L181:
	;
	v776 = v757
	goto L153
L182:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v716&int32(65535)<<(uint(int32(2))%32)+(v522+int32(24))-int32(4))))
	v735 = v522 + v732&int32(32767)
	if v711&int32(1) == int32(0) {
		goto L186
	} else {
		goto L187
	}
L183:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
	v754 = F__hash_get_indextuple_hashkey(m, v735)
	mBase = m.M
	if v753 != v754 {
		v776 = v706
		goto L153
	} else {
		goto L193
	}
L184:
	;
	goto L183
L185:
	;
	v750 = v716 - int32(1)
	if v750&int32(65535) != 0 {
		v716 = v750
		goto L182
	} else {
		goto L192
	}
L186:
	;
	v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	if v742 != int32(1) {
		goto L184
	} else {
		goto L190
	}
L187:
	;
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+13)))
	if v738 != 0 {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735)+7)))
	if v739&int32(32) != 0 {
		goto L185
	} else {
		goto L189
	}
L189:
	;
	goto L186
L190:
	;
	v745 = int32(98304)
	if v732&v745 != v745 {
		goto L184
	} else {
		goto L191
	}
L191:
	;
	goto L185
L192:
	;
	v776 = v706
	goto L153
L193:
	;
	v756 = int32(1)
	v757 = v706 - v756
	v760 = v601 + int32(52) + v757<<(uint(int32(3))%32)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v735)))
	*(*int32)(unsafe.Add(mBase, uint32(v760))) = v761
	v763 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v735)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v760)+6)) = uint16(v716)
	*(*uint16)(unsafe.Add(mBase, uint32(v760)+4)) = uint16(v763)
	v767 = v716 - v756
	if v767&int32(65535) != 0 {
		v702 = v767
		v706 = v757
		goto L180
	} else {
		goto L194
	}
L194:
	;
	goto L181
L195:
	;
	goto L81
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v1123
	goto L13
L197:
	;
	v897 = int32(0)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L220
L198:
	;
	v839 = v828
	goto L200
L199:
	;
	v839 = int32(base.Ui32(v830+int32(262120))>>(uint(int32(2))%32)) + v828
	goto L200
L200:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v839&int32(65535)) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v848 = v828
	v849 = v839
	goto L204
L202:
	;
	v889 = v828
	goto L203
L203:
	;
	v895 = v889 & int32(65535)
	goto L197
L204:
	;
	v853 = int32(65535)
	v859 = int32(base.Ui32(v849&v853+v848&v853) >> (uint(int32(1)) % 32))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v859<<(uint(int32(2))%32)+(v41+int32(24))-int32(4))))
	v868 = v41 + v865&int32(32767)
	v871 = int32(*(*int16)(unsafe.Add(mBase, uint32(v868)+6)))
	if int32(0) <= v871 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v889 = v883
	goto L203
L206:
	;
	v874 = int32(8)
	goto L208
L207:
	;
	v874 = int32(16)
	goto L208
L208:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v868+v874)))
	v877 = base.B2i32(base.Ui32(v876) < base.Ui32(v67))
	if base.Ui32(v876) < base.Ui32(v67) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v878 = v849
	goto L211
L210:
	;
	v878 = v859
	goto L211
L211:
	;
	if base.Ui32(v876) < base.Ui32(v67) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v883 = v859 + int32(1)
	goto L214
L213:
	;
	v883 = v848
	goto L214
L214:
	;
	if base.Ui32(v883&int32(65535)) < base.Ui32(v878&int32(65535)) {
		v848 = v883
		v849 = v878
		goto L204
	} else {
		goto L215
	}
L215:
	;
	goto L205
L216:
	;
	v1097 = v1079 & int32(65535)
	if v1097 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L217:
	;
	goto L216
L220:
	;
	goto L221
L221:
	;
	v915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v915) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v923 = int32(base.Ui32(v915+int32(262120)) >> (uint(int32(2)) % 32))
	goto L227
L226:
	;
	v923 = int32(0)
	goto L227
L227:
	;
	v925 = v923 & int32(65535)
	if base.Ui32(v925) < base.Ui32(v895) {
		v1079 = v897
		goto L217
	} else {
		goto L228
	}
L228:
	;
	v933 = v895
	v937 = v897
	goto L229
L229:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904)+12)))
	v947 = v933
	goto L231
L230:
	;
	v1079 = v997
	goto L217
L231:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v947&int32(65535)<<(uint(int32(2))%32)+(v41+int32(24))-int32(4))))
	v966 = v41 + v963&int32(32767)
	if v942&int32(1) == int32(0) {
		goto L235
	} else {
		goto L236
	}
L232:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v904)))
	v986 = F__hash_get_indextuple_hashkey(m, v966)
	mBase = m.M
	if v985 != v986 {
		v1079 = v937
		goto L217
	} else {
		goto L242
	}
L233:
	;
	goto L232
L234:
	;
	v981 = v947 + int32(1)
	if base.Ui32(v981&int32(65535)) <= base.Ui32(v925) {
		v947 = v981
		goto L231
	} else {
		goto L241
	}
L235:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	if v973 != int32(1) {
		goto L233
	} else {
		goto L239
	}
L236:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904)+13)))
	if v969 != 0 {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966)+7)))
	if v970&int32(32) != 0 {
		goto L234
	} else {
		goto L238
	}
L238:
	;
	goto L235
L239:
	;
	v976 = int32(98304)
	if v963&v976 != v976 {
		goto L233
	} else {
		goto L240
	}
L240:
	;
	goto L234
L241:
	;
	v1079 = v937
	goto L217
L242:
	;
	v990 = v904 + int32(52) + v937<<(uint(int32(3))%32)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v966)))
	*(*int32)(unsafe.Add(mBase, uint32(v990))) = v991
	v993 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v966)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v990)+6)) = uint16(v947)
	*(*uint16)(unsafe.Add(mBase, uint32(v990)+4)) = uint16(v993)
	v996 = int32(1)
	v997 = v937 + v996
	v999 = v947 + v996
	if base.Ui32(v999&int32(65535)) <= base.Ui32(v925) {
		v933 = v999
		v937 = v997
		goto L229
	} else {
		goto L243
	}
L243:
	;
	goto L230
L259:
	;
	goto L262
L260:
	;
	v1436 = v1097
	goto L261
L261:
	;
	v1445 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v1445
	v1466 = v1445
	v1467 = v1436 - int32(1)
	goto L12
L262:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if int32(0) < v1110 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1436 = v1432
	goto L261
L264:
	;
	F__hash_kill_items(m, l0)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1115 = int32(-1)
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v1116 == v1117 {
		v1123 = v1115
		goto L268
	} else {
		goto L269
	}
L267:
	;
	goto L266
L268:
	;
	F__hash_readnext(m, l0, v13+int32(12), v13+int32(8), v13+int32(4))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v1116 == v1119 {
		v1123 = v1115
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1121)))
	v1123 = v1122
	goto L268
L271:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v1132 == int32(0) {
		goto L196
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v1132
	if v1132 < int32(0) {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v1154
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v1163 = int32(1)
	v1165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1156)+12)))
	if base.Ui32(v1165) < base.Ui32(int32(25)) {
		goto L278
	} else {
		goto L279
	}
L274:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1139+(v1132^int32(-1))<<(uint(int32(6))%32))+16))
	v1154 = v1145
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1147+v1132<<(uint(int32(6))%32)+int32(-64))+16))
	v1154 = v1153
	goto L273
L277:
	;
	v1232 = int32(0)
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L300
L278:
	;
	v1174 = v1163
	goto L280
L279:
	;
	v1174 = int32(base.Ui32(v1165+int32(262120))>>(uint(int32(2))%32)) + v1163
	goto L280
L280:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1174&int32(65535)) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1183 = v1163
	v1184 = v1174
	goto L284
L282:
	;
	v1224 = v1163
	goto L283
L283:
	;
	v1230 = v1224 & int32(65535)
	goto L277
L284:
	;
	v1188 = int32(65535)
	v1194 = int32(base.Ui32(v1184&v1188+v1183&v1188) >> (uint(int32(1)) % 32))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1194<<(uint(int32(2))%32)+(v1156+int32(24))-int32(4))))
	v1203 = v1156 + v1200&int32(32767)
	v1206 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1203)+6)))
	if int32(0) <= v1206 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	v1224 = v1218
	goto L283
L286:
	;
	v1209 = int32(8)
	goto L288
L287:
	;
	v1209 = int32(16)
	goto L288
L288:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1203+v1209)))
	v1212 = base.B2i32(base.Ui32(v1211) < base.Ui32(v1157))
	if base.Ui32(v1211) < base.Ui32(v1157) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1213 = v1184
	goto L291
L290:
	;
	v1213 = v1194
	goto L291
L291:
	;
	if base.Ui32(v1211) < base.Ui32(v1157) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1218 = v1194 + int32(1)
	goto L294
L293:
	;
	v1218 = v1183
	goto L294
L294:
	;
	if base.Ui32(v1218&int32(65535)) < base.Ui32(v1213&int32(65535)) {
		v1183 = v1218
		v1184 = v1213
		goto L284
	} else {
		goto L295
	}
L295:
	;
	goto L285
L296:
	;
	v1432 = v1414 & int32(65535)
	if v1432 == int32(0) {
		goto L262
	} else {
		goto L339
	}
L297:
	;
	goto L296
L300:
	;
	goto L301
L301:
	;
	v1250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1156)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1250) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1258 = int32(base.Ui32(v1250+int32(262120)) >> (uint(int32(2)) % 32))
	goto L307
L306:
	;
	v1258 = int32(0)
	goto L307
L307:
	;
	v1260 = v1258 & int32(65535)
	if base.Ui32(v1260) < base.Ui32(v1230) {
		v1414 = v1232
		goto L297
	} else {
		goto L308
	}
L308:
	;
	v1268 = v1230
	v1272 = v1232
	goto L309
L309:
	;
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+12)))
	v1282 = v1268
	goto L311
L310:
	;
	v1414 = v1332
	goto L297
L311:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1282&int32(65535)<<(uint(int32(2))%32)+(v1156+int32(24))-int32(4))))
	v1301 = v1156 + v1298&int32(32767)
	if v1277&int32(1) == int32(0) {
		goto L315
	} else {
		goto L316
	}
L312:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1239)))
	v1321 = F__hash_get_indextuple_hashkey(m, v1301)
	mBase = m.M
	if v1320 != v1321 {
		v1414 = v1272
		goto L297
	} else {
		goto L322
	}
L313:
	;
	goto L312
L314:
	;
	v1316 = v1282 + int32(1)
	if base.Ui32(v1316&int32(65535)) <= base.Ui32(v1260) {
		v1282 = v1316
		goto L311
	} else {
		goto L321
	}
L315:
	;
	v1308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	if v1308 != int32(1) {
		goto L313
	} else {
		goto L319
	}
L316:
	;
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+13)))
	if v1304 != 0 {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301)+7)))
	if v1305&int32(32) != 0 {
		goto L314
	} else {
		goto L318
	}
L318:
	;
	goto L315
L319:
	;
	v1311 = int32(98304)
	if v1298&v1311 != v1311 {
		goto L313
	} else {
		goto L320
	}
L320:
	;
	goto L314
L321:
	;
	v1414 = v1272
	goto L297
L322:
	;
	v1325 = v1239 + int32(52) + v1272<<(uint(int32(3))%32)
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1301)))
	*(*int32)(unsafe.Add(mBase, uint32(v1325))) = v1326
	v1328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1301)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1325)+6)) = uint16(v1282)
	*(*uint16)(unsafe.Add(mBase, uint32(v1325)+4)) = uint16(v1328)
	v1331 = int32(1)
	v1332 = v1272 + v1331
	v1334 = v1282 + v1331
	if base.Ui32(v1334&int32(65535)) <= base.Ui32(v1260) {
		v1268 = v1334
		v1272 = v1332
		goto L309
	} else {
		goto L323
	}
L323:
	;
	goto L310
L339:
	;
	goto L263
L340:
	;
	v1513 = int32(1)
	goto L11
L341:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1482)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v1493
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1482)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v1495
	F_UnlockReleaseBuffer(m, v1478)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L348
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(-1)
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1485)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v1488
	F_LockBuffer(m, v1478, int32(0))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L1
	} else {
		goto L347
	}
L343:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v1485 = v1481
	goto L342
L344:
	;
	goto L345
L345:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v1478 != v1483 {
		goto L341
	} else {
		goto L346
	}
L346:
	;
	v1485 = v1482
	goto L342
L347:
	;
	goto L340
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	goto L340
}
func F_get_hash_memory_limit(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v5 int32
	_ = v5
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v13 float64
	_ = v13
	var v19 int32
	_ = v19
	v3 = *(*float64)(unsafe.Add(mBase, _consts[433]))
	v5 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	v9 = base.F64_mul(base.F64_mul(v3, base.F64_convert_i32_s(v5)), float64(1024))
	v10 = float64(4.294967295e+09)
	if base.F64_lt(v9, v10) != 0 {
		v13 = v9
	} else {
		v13 = v10
	}
	if base.F64_lt(v13, float64(4.294967296e+09))&base.F64_ge(v13, float64(0)) != 0 {
		v19 = base.I32_trunc_f64_u(v13)
		return v19
	} else {
		return int32(0)
	}
}
func F_hash_agg_entry_size(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	if l2 != 0 {
		v4 = int32(1)
		if l2&(l2-v4) != 0 {
			v12 = v4 << (uint(int32(32)-base.I32_clz(l2)) % 32)
		} else {
			v12 = l2
		}
		v16 = v12 + int32(8)
	} else {
		v16 = int32(0)
	}
	return v16 + ((l1+int32(23))&int32(-8) + l0<<(uint(int32(3))%32)) + int32(12)
}
func F_hash_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	if v8 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v20
				F_errmsg(m, int32(25373), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(489238), int32(3937), int32(81579))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		m.G0 = v5 + int32(16)
		return int32(0)
	}
}
func F_hash_corrupted(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v7 == int32(1) {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v14
			F_errmsg_internal(m, int32(440253), v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errfinish(m, int32(492463), int32(1790), int32(440029))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v28
			F_errmsg_internal(m, int32(440253), v5+int32(16))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				F_errfinish(m, int32(492463), int32(1792), int32(440029))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
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
func F_hash_ltree_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int64
	_ = v346
	var v349 int64
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v378 int64
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v380 = F_Int64GetDatum(m, v378)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L61
	}
L4:
	;
	v21 = v9 + int32(8)
	v22 = v15
	v24 = int64(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v367 != v9 {
		goto L57
	} else {
		goto L58
	}
L7:
	;
	v27 = v21 + int32(2)
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	v34 = v28 - int32(1636608432)
	if v14 == int64(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v9 == v360 {
		v378 = v349
		goto L3
	} else {
		goto L54
	}
L9:
	;
	v344 = F_Int64GetDatum(m, base.I64_extend_i32_u(v334)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v326^v334-base.I32_rotl(v334, int32(24))))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L52
	}
L10:
	;
	if v27&int32(3) != 0 {
		goto L26
	} else {
		goto L27
	}
L11:
	;
	v71 = v34
	v73 = v34
	v75 = v34
	goto L10
L12:
	;
	goto L13
L13:
	;
	v38 = v34 + base.I32_wrap_i64(v14)
	v39 = v38 + v34
	v43 = int32(4)
	v45 = base.I32_wrap_i64(int64(base.Ui64(v14)>>(uint(int64(32))%64))) ^ base.I32_rotl(v34, v43)
	v49 = v38 - v45 ^ base.I32_rotl(v45, int32(6))
	v53 = v39 - v49 ^ base.I32_rotl(v49, int32(8))
	v54 = v45 + v39
	v55 = v49 + v54
	v56 = v53 + v55
	v60 = v54 - v53 ^ base.I32_rotl(v53, int32(16))
	v64 = v55 - v60 ^ base.I32_rotl(v60, int32(19))
	v69 = v60 + v56
	v71 = v69
	v73 = v56 - v64 ^ base.I32_rotl(v64, v43)
	v75 = v64 + v69
	goto L10
L14:
	;
	v312 = int32(14)
	v314 = v308 ^ v309 - base.I32_rotl(v308, v312)
	v318 = v314 ^ v307 - base.I32_rotl(v314, int32(11))
	v322 = v318 ^ v308 - base.I32_rotl(v318, int32(25))
	v326 = v322 ^ v314 - base.I32_rotl(v322, int32(16))
	v330 = v326 ^ v318 - base.I32_rotl(v326, int32(4))
	v334 = v330 ^ v322 - base.I32_rotl(v330, v312)
	goto L9
L15:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v307 = v299 + v302
	v308 = v300
	v309 = v301
	goto L14
L16:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	v299 = v295<<(uint(int32(8))%32) + v292
	v300 = v293
	v301 = v294
	goto L15
L17:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
	v292 = v288<<(uint(int32(16))%32) + v285
	v293 = v286
	v294 = v287
	goto L16
L18:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
	v285 = v281<<(uint(int32(24))%32) + v132
	v286 = v279
	v287 = v280
	goto L17
L19:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
	v279 = v275 + v277
	v280 = v276
	goto L18
L20:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
	v275 = v271<<(uint(int32(8))%32) + v269
	v276 = v270
	goto L19
L21:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+6)))
	v269 = v265<<(uint(int32(16))%32) + v263
	v270 = v264
	goto L20
L22:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+7)))
	v263 = v259<<(uint(int32(24))%32) + v133
	v264 = v258
	goto L21
L23:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+8)))
	v258 = v254<<(uint(int32(8))%32) + v253
	goto L22
L24:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+9)))
	v253 = v249<<(uint(int32(16))%32) + v248
	goto L23
L25:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+10)))
	v248 = v244<<(uint(int32(24))%32) + v134
	goto L24
L26:
	;
	if base.Ui32(int32(11)) < base.Ui32(v28) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v28) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v80 = v27
	v81 = v28
	v83 = v71
	v84 = v75
	v85 = v73
	goto L32
L30:
	;
	v129 = v27
	v130 = v28
	v132 = v71
	v133 = v75
	v134 = v73
	goto L31
L31:
	;
	switch v130 - int32(1) {
	case 0:
		v299 = v132
		v300 = v133
		v301 = v134
		goto L15
	case 1:
		v292 = v132
		v293 = v133
		v294 = v134
		goto L16
	case 2:
		v285 = v132
		v286 = v133
		v287 = v134
		goto L17
	case 3:
		v279 = v133
		v280 = v134
		goto L18
	case 4:
		v275 = v133
		v276 = v134
		goto L19
	case 5:
		v269 = v133
		v270 = v134
		goto L20
	case 6:
		v263 = v133
		v264 = v134
		goto L21
	case 7:
		v258 = v134
		goto L22
	case 8:
		v253 = v134
		goto L23
	case 9:
		v248 = v134
		goto L24
	case 10:
		goto L25
	default:
		v307 = v132
		v308 = v133
		v309 = v134
		goto L14
	}
L32:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v88 = v87 + v84
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v92 = v91 + v85
	v94 = int32(4)
	v96 = v89 + v83 - v92 ^ base.I32_rotl(v92, v94)
	v100 = v88 - v96 ^ base.I32_rotl(v96, int32(6))
	v101 = v92 + v88
	v102 = v96 + v101
	v103 = v100 + v102
	v107 = v101 - v100 ^ base.I32_rotl(v100, int32(8))
	v111 = v102 - v107 ^ base.I32_rotl(v107, int32(16))
	v115 = v103 - v111 ^ base.I32_rotl(v111, int32(19))
	v116 = v107 + v103
	v117 = v111 + v116
	v118 = v115 + v117
	v122 = v116 - v115 ^ base.I32_rotl(v115, v94)
	v123 = int32(12)
	v124 = v80 + v123
	v126 = v81 - v123
	if base.Ui32(int32(11)) < base.Ui32(v126) {
		v80 = v124
		v81 = v126
		v83 = v117
		v84 = v118
		v85 = v122
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v129 = v124
	v130 = v126
	v132 = v117
	v133 = v118
	v134 = v122
	goto L31
L34:
	;
	goto L33
L35:
	;
	v140 = v27
	v141 = v28
	v143 = v71
	v144 = v75
	v145 = v73
	goto L38
L36:
	;
	v189 = v27
	v190 = v28
	v192 = v71
	v193 = v75
	v194 = v73
	goto L37
L37:
	;
	switch v190 - int32(1) {
	case 0:
		v241 = v192
		goto L41
	case 1:
		v236 = v192
		goto L42
	case 2:
		goto L43
	case 3:
		v229 = v193
		goto L44
	case 4:
		v226 = v193
		goto L45
	case 5:
		v221 = v193
		goto L46
	case 6:
		goto L47
	case 7:
		v212 = v194
		goto L48
	case 8:
		v207 = v194
		goto L49
	case 9:
		v202 = v194
		goto L50
	case 10:
		goto L51
	default:
		v307 = v192
		v308 = v193
		v309 = v194
		goto L14
	}
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v148 = v147 + v144
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v140)+8))
	v152 = v151 + v145
	v154 = int32(4)
	v156 = v149 + v143 - v152 ^ base.I32_rotl(v152, v154)
	v160 = v148 - v156 ^ base.I32_rotl(v156, int32(6))
	v161 = v152 + v148
	v162 = v156 + v161
	v163 = v160 + v162
	v167 = v161 - v160 ^ base.I32_rotl(v160, int32(8))
	v171 = v162 - v167 ^ base.I32_rotl(v167, int32(16))
	v175 = v163 - v171 ^ base.I32_rotl(v171, int32(19))
	v176 = v167 + v163
	v177 = v171 + v176
	v178 = v175 + v177
	v182 = v176 - v175 ^ base.I32_rotl(v175, v154)
	v183 = int32(12)
	v184 = v140 + v183
	v186 = v141 - v183
	if base.Ui32(int32(11)) < base.Ui32(v186) {
		v140 = v184
		v141 = v186
		v143 = v177
		v144 = v178
		v145 = v182
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v189 = v184
	v190 = v186
	v192 = v177
	v193 = v178
	v194 = v182
	goto L37
L40:
	;
	goto L39
L41:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v307 = v241 + v242
	v308 = v193
	v309 = v194
	goto L14
L42:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	v241 = v237<<(uint(int32(8))%32) + v236
	goto L41
L43:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+2)))
	v236 = v232<<(uint(int32(16))%32) + v192
	goto L42
L44:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v307 = v230 + v192
	v308 = v229
	v309 = v194
	goto L14
L45:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+4)))
	v229 = v226 + v227
	goto L44
L46:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+5)))
	v226 = v222<<(uint(int32(8))%32) + v221
	goto L45
L47:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+6)))
	v221 = v217<<(uint(int32(16))%32) + v193
	goto L46
L48:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v307 = v213 + v192
	v308 = v215 + v193
	v309 = v212
	goto L14
L49:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+8)))
	v212 = v208<<(uint(int32(8))%32) + v207
	goto L48
L50:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+9)))
	v207 = v203<<(uint(int32(16))%32) + v202
	goto L49
L51:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+10)))
	v202 = v198<<(uint(int32(24))%32) + v194
	goto L50
L52:
	;
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v344)))
	v349 = v346 + v24*int64(31)
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	v356 = int32(1)
	if base.Ui32(v356) < base.Ui32(v22) {
		v21 = v21 + (v350+int32(9))&int32(131064)
		v22 = v22 - v356
		v24 = v349
		goto L7
	} else {
		goto L53
	}
L53:
	;
	goto L8
L54:
	;
	F_pfree(m, v9)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v364 = F_Int64GetDatum(m, v349)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	return v364
L57:
	;
	F_pfree(m, v9)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v378 = v14 + int64(1)
	goto L3
L60:
	;
	goto L59
L61:
	;
	return v380
}
func F_hash_multirange(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L40
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L37
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+200))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	if v38 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 == v22 {
		v35 = v24
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v28 = F_lookup_type_cache(m, v22, int32(65536))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
	if v30 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v28
	v35 = v28
	goto L5
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v43 = F_lookup_type_cache(m, v41, int32(128))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v48 = v37
	goto L14
L14:
	;
	v49 = int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if int32(0) < v50 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+136))
	if v45 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v48 = v43
	goto L14
L17:
	;
	v56 = v48 + int32(132)
	v59 = int32(0)
	v63 = v49
	goto L20
L18:
	;
	v147 = v49
	goto L19
L19:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v154 != v18 {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(8)+v70<<(uint(int32(2))%32)+v59))))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
	F_multirange_get_bounds(m, v76, v18, v59, v15+int32(40), v15+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v147 = v138
	goto L19
L22:
	;
	v83 = int32(0)
	if v75&int32(41) == v83 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+208))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v91 = F_FunctionCall1Coll(m, v56, v89, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v93 = v83
	goto L25
L25:
	;
	if v75&int32(81) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v93 = v91
	goto L25
L27:
	;
	v103 = int32(0)
	goto L29
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+208))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v101 = F_FunctionCall1Coll(m, v56, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v108 = int32(711645284)
	v111 = v75 - int32(1636608428) ^ v108 - int32(1455628627)
	v116 = v111 ^ int32(-1636608428) - base.I32_rotl(v111, int32(25))
	v121 = v116 ^ v108 - base.I32_rotl(v116, int32(16))
	v125 = v121 ^ v111 - base.I32_rotl(v121, int32(4))
	v129 = v125 ^ v116 - base.I32_rotl(v125, int32(14))
	goto L31
L30:
	;
	v103 = v101
	goto L29
L31:
	;
	v135 = int32(1)
	v138 = v63*int32(31) + (v103 ^ base.I32_rotl(v129^v121-base.I32_rotl(v129, int32(24))^v93, v135))
	v140 = v59 + v135
	if v140 != v50 {
		v59 = v140
		v63 = v138
		goto L20
	} else {
		goto L32
	}
L32:
	;
	goto L21
L33:
	;
	F_pfree(m, v18)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	m.G0 = v15 + int32(48)
	return v147
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
	F_errmsg_internal(m, int32(366213), v15)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(488725), int32(558), int32(394659))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v183 = F_format_type_be(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v183
	F_errmsg(m, int32(187948), v15+int32(16))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(488725), int32(2807), int32(396746))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hash_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = m.T0[v6].(func(*base.Module, int32, int32) int32)(m, l1, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_hash_search_with_hash_value(m, l0, l1, v7, l2, l3)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_hash_seq_init_with_hash_value(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v4)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	if v16 == v4 {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
		if int32(100) <= v20 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return
			} else {
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v81
				F_errmsg_internal(m, int32(688301), v9)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					F_errfinish(m, int32(492463), int32(1872), int32(281579))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v20<<(uint(int32(2))%32))+uint32(_consts[1188]))) = l1
			v29 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
			v31 = int32(4469936)
			v32 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
			*(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_consts[1189]))) = v30
			*(*int32)(unsafe.Add(mBase, _consts[1187])) = v32 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
			v45 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v45)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+396))
			v49 = v48 & l2
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+392))
			if base.Ui32(v50) < base.Ui32(v49) {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+400))
				v54 = v52 & v49
			} else {
				v54 = v49
			}
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+int32(base.Ui32(v54)>>(uint(v56)%32))<<(uint(int32(2))%32))))
			if v61 == int32(0) {
				F_hash_corrupted(m, l1)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v61+(v64-int32(1))&v54<<(uint(int32(2))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72
				m.G0 = v9 + int32(16)
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
		v45 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v45)
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+396))
		v49 = v48 & l2
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+392))
		if base.Ui32(v50) < base.Ui32(v49) {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+400))
			v54 = v52 & v49
		} else {
			v54 = v49
		}
		v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+int32(base.Ui32(v54)>>(uint(v56)%32))<<(uint(int32(2))%32))))
		if v61 == int32(0) {
			F_hash_corrupted(m, l1)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v61+(v64-int32(1))&v54<<(uint(int32(2))%32))))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_hash_seq_search(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v17 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L49
	} else {
		goto L56
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L49
	} else {
		goto L53
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L49
	} else {
		goto L50
	}
L4:
	;
	m.G0 = v14 + int32(48)
	return v269
L5:
	;
	v22 = v16
	goto L9
L6:
	;
	goto L7
L7:
	;
	if v16 != 0 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+34)))
	if v42 != 0 {
		v269 = int32(0)
		goto L4
	} else {
		goto L13
	}
L9:
	;
	if v22 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v269 = v22 + int32(8)
	goto L4
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v35 != v36 {
		v22 = v33
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	v46 = v44
	goto L14
L14:
	;
	v57 = v46 - int32(1)
	if v57 < int32(0) {
		goto L3
	} else {
		goto L16
	}
L15:
	;
	v67 = v44 - int32(1)
	v68 = int32(2)
	v69 = v67 << (uint(v68) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1188])))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1188]))) = v72
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1189])))
	*(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(v68)%32))+uint32(_consts[1189]))) = v80
	*(*int32)(unsafe.Add(mBase, _consts[1187])) = v67
	v269 = int32(0)
	goto L4
L16:
	;
	v61 = v57 << (uint(int32(2)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[1188])))
	if v64 != v41 {
		v46 = v57
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v85
	if v85 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+392))
	if base.Ui32(v98) < base.Ui32(v95) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v89 + int32(1)
	goto L23
L22:
	;
	goto L23
L23:
	;
	v269 = v16 + int32(8)
	goto L4
L24:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+34)))
	if v101 != 0 {
		v269 = int32(0)
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v96)+44))
	v146 = int32(base.Ui32(v95) >> (uint(v145) % 32))
	v147 = int32(2)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144+v146<<(uint(v147)%32))))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
	v154 = (v151 - int32(1)) & v95
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v150+v154<<(uint(v147)%32))))
	if v158 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	v105 = v103
	goto L28
L28:
	;
	v116 = v105 - int32(1)
	if v116 < int32(0) {
		goto L2
	} else {
		goto L30
	}
L29:
	;
	v126 = v103 - int32(1)
	v127 = int32(2)
	v128 = v126 << (uint(v127) % 32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+uint32(_consts[1188])))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[1188]))) = v131
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v128)+uint32(_consts[1189])))
	*(*int32)(unsafe.Add(mBase, uint32(v116<<(uint(v127)%32))+uint32(_consts[1189]))) = v139
	*(*int32)(unsafe.Add(mBase, _consts[1187])) = v126
	v269 = int32(0)
	goto L4
L30:
	;
	v120 = v116 << (uint(int32(2)) % 32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[1188])))
	if v123 != v96 {
		v105 = v116
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v162 = v154
	v164 = v95
	v167 = v146
	v168 = v150
	goto L35
L33:
	;
	v242 = v95
	v244 = v158
	goto L34
L34:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v242 + base.B2i32(v250 == int32(0))
	v269 = v244 + int32(8)
	goto L4
L35:
	;
	v173 = v164 + int32(1)
	if base.Ui32(v98) < base.Ui32(v173) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v242 = v173
	v244 = v236
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v173
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+34)))
	if v177 != 0 {
		v269 = int32(0)
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v221 = v162 + int32(1)
	if v151 <= v221 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	v181 = v179
	goto L41
L41:
	;
	v192 = v181 - int32(1)
	if v192 < int32(0) {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v202 = v179 - int32(1)
	v203 = int32(2)
	v204 = v202 << (uint(v203) % 32)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_consts[1188])))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+uint32(_consts[1188]))) = v207
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_consts[1189])))
	*(*int32)(unsafe.Add(mBase, uint32(v192<<(uint(v203)%32))+uint32(_consts[1189]))) = v215
	*(*int32)(unsafe.Add(mBase, _consts[1187])) = v202
	v269 = int32(0)
	goto L4
L43:
	;
	v196 = v192 << (uint(int32(2)) % 32)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v196)+uint32(_consts[1188])))
	if v199 != v96 {
		v181 = v192
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v224 = v167 + int32(1)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v144+v224<<(uint(int32(2))%32))))
	v230 = int32(0)
	v231 = v224
	v232 = v228
	goto L47
L46:
	;
	v230 = v221
	v231 = v167
	v232 = v168
	goto L47
L47:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232+v230<<(uint(int32(2))%32))))
	if v236 == int32(0) {
		v162 = v230
		v164 = v173
		v167 = v231
		v168 = v232
		goto L35
	} else {
		goto L48
	}
L48:
	;
	goto L36
L49:
	;
	return int32(0)
L50:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v280
	F_errmsg_internal(m, int32(697946), v14)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(492463), int32(1896), int32(281577))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v294
	F_errmsg_internal(m, int32(697946), v14+int32(16))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(492463), int32(1896), int32(281577))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L49
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v310
	F_errmsg_internal(m, int32(697946), v14+int32(32))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(492463), int32(1896), int32(281577))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L49
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hash_seq_term(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+34)))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	v16 = v15
	goto L5
L3:
	;
	goto L4
L4:
	;
	m.G0 = v8 + int32(16)
	return
L5:
	;
	v22 = v16 - int32(1)
	if v22 < int32(0) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v32 = v15 - int32(1)
	v33 = int32(2)
	v34 = v32 << (uint(v33) % 32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_consts[1188])))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1188]))) = v37
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_consts[1189])))
	*(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(v33)%32))+uint32(_consts[1189]))) = v45
	*(*int32)(unsafe.Add(mBase, _consts[1187])) = v32
	goto L4
L7:
	;
	v26 = v22 << (uint(int32(2)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1188])))
	if v29 != v10 {
		v16 = v22
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
	F_errmsg_internal(m, int32(697946), v8)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(492463), int32(1896), int32(281577))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
