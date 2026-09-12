package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_btree_redo(m *base.Module, l0 int32) {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
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
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int64
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v517 int64
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int64
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int64
	_ = v549
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v633 int64
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int64
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int64
	_ = v775
	var v776 int32
	_ = v776
	var v777 int64
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v810 int64
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int64
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int64
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1106 int64
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int64
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int64
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v22 = int32(4562096)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+48)))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	v31 = v25 & int32(240)
	switch int32(base.Ui32(v31)>>(uint(int32(4))%32)) - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	case 5:
		goto L14
	case 6:
		goto L12
	case 7, 8:
		goto L10
	case 9:
		goto L9
	case 10:
		goto L11
	case 11:
		goto L13
	case 12:
		goto L8
	case 13:
		goto L7
	case 14:
		goto L6
	default:
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L20
	} else {
		goto L297
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L20
	} else {
		goto L294
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L20
	} else {
		goto L291
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
	v1189 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	F_MemoryContextReset(m, v1189)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L20
	} else {
		goto L290
	}
L5:
	;
	v1165 = int32(0)
	F_btree_xlog_insert(m, int32(1), v1165, v1165, l0)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L20
	} else {
		goto L289
	}
L6:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L20
	} else {
		goto L286
	}
L7:
	;
	F__bt_restore_meta(m, l0, int32(0))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L20
	} else {
		goto L285
	}
L8:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	if base.Ui32(v1132) < base.Ui32(int32(2)) {
		goto L4
	} else {
		goto L283
	}
L9:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v995 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v997 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L20
	} else {
		goto L247
	}
L10:
	;
	v775 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v777 = *(*int64)(unsafe.Add(mBase, uint32(v776)+16))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v776)+8))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v776)))
	if v780 != 0 {
		goto L193
	} else {
		goto L194
	}
L11:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v649 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v653 = F_XLogReadBufferForRedo(m, l0, int32(1), v18+int32(-4))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L20
	} else {
		goto L168
	}
L12:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v533 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v535 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	if base.Ui32(int32(2)) <= base.Ui32(v535) {
		goto L133
	} else {
		goto L134
	}
L13:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v434 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v435 = int32(0)
	v440 = F_XLogReadBufferForRedoExtended(m, l0, v435, v435, int32(1), v18+int32(-20))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L20
	} else {
		goto L103
	}
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v62 = F_XLogReadBufferForRedo(m, l0, int32(0), v18+int32(-20))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L26
	}
L15:
	;
	v52 = int32(1)
	F_btree_xlog_insert(m, v52, int32(0), v52, l0)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L20
	} else {
		goto L25
	}
L16:
	;
	F_btree_xlog_split(m, int32(0), l0)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L20
	} else {
		goto L24
	}
L17:
	;
	F_btree_xlog_split(m, int32(1), l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L20
	} else {
		goto L23
	}
L18:
	;
	v41 = int32(0)
	F_btree_xlog_insert(m, v41, int32(1), v41, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L20
	} else {
		goto L22
	}
L19:
	;
	v36 = int32(0)
	F_btree_xlog_insert(m, v36, v36, v36, l0)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	goto L4
L22:
	;
	goto L4
L23:
	;
	goto L4
L24:
	;
	goto L4
L25:
	;
	goto L4
L26:
	;
	if v62 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v66 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+72))
	if v70 < v66 {
		v92 = v66
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v428 == int32(0) {
		goto L4
	} else {
		goto L101
	}
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v96 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	goto L30
L32:
	;
	v76 = v69 + int32(76)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v77 != int32(1) {
		v92 = v66
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+43)))
	if v80 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v92 = v66
	goto L31
L35:
	;
	goto L36
L36:
	;
	goto L39
L39:
	;
	goto L40
L40:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v76)+44))
	v92 = v90
	goto L31
L41:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+16)))
	v117 = F_palloc(m, int32(1676))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L20
	} else {
		goto L45
	}
L42:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100+(v96^int32(-1))<<(uint(int32(2))%32))))
	v114 = v106
	goto L41
L43:
	;
	goto L44
L44:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v114 = v108 + v96<<(uint(int32(13))%32) + int32(-8192)
	goto L41
L45:
	;
	v119 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+20)) = v119
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)) = uint16(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v117)+4)) = int64(11613591568384)
	v127 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v127)
	v130 = F_palloc(m, int32(2704))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v132 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v117)+28)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v117)+24)) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v117)+36)) = v132
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+12)))
	v138 = v115 + v114
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = F_PageGetTempPageCopySpecial(m, v114)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v142 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v114)+24))
	v151 = F_PageAddItemExtended(m, v140, v114+v143&int32(32767), int32(base.Ui32(v143)>>(uint(int32(17))%32)), int32(1), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L20
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v139 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if v151 == int32(0) {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v158 = int32(2)
	goto L55
L54:
	;
	v158 = int32(1)
	goto L55
L55:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v137) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v166 = int32(base.Ui32(v137+int32(262120)) >> (uint(int32(2)) % 32))
	goto L58
L57:
	;
	v166 = int32(0)
	goto L58
L58:
	;
	v168 = v166 & int32(65535)
	if base.Ui32(v158) <= base.Ui32(v168) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v172 = v158
	goto L62
L60:
	;
	goto L61
L61:
	;
	F__bt_dedup_finish_pending(m, v140, v117)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L20
	} else {
		goto L95
	}
L62:
	;
	v190 = v172 & int32(65535)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190<<(uint(int32(2))%32)+(v114+int32(24))-int32(4))))
	v199 = v114 + v196&int32(32767)
	if v190 != v158 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L61
L64:
	;
	v370 = v172 + int32(1)
	if base.Ui32(v370&int32(65535)) <= base.Ui32(v168) {
		v172 = v370
		goto L62
	} else {
		goto L94
	}
L65:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v117)+40))
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57))))
	if v202 <= v201 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v307 = v158
	goto L67
L67:
	;
	v309 = v307 & int32(65535)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+7)))
	if v311&int32(32) != 0 {
		goto L90
	} else {
		goto L91
	}
L68:
	;
	F__bt_dedup_finish_pending(m, v140, v117)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L20
	} else {
		goto L86
	}
L69:
	;
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)))
	v207 = v92 + v201<<(uint(int32(2))%32)
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207))))
	if v204 != v208 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+2)))
	if v211 <= v210 {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v217 = int32(1)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+7)))
	if v218&int32(32) == int32(0) {
		v236 = v217
		v238 = v199
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if base.Ui32((v240+(v241+v236)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v239) {
		goto L64
	} else {
		goto L82
	}
L73:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v117)+20))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v117)+28))
	v250 = base.B2i32(base.Ui32((v240+(v241+v236)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v239))
	if v250 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+4)))
	if v223&int32(8192) == int32(0) {
		v236 = v217
		v238 = v199
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+2)))
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199))))
	v236 = v223 & int32(4095)
	v238 = v199 + (v230 | v231<<(uint(int32(16))%32))
	goto L73
L76:
	;
	goto L72
L77:
	;
	v283 = v117 + v281
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v284 + v282
	goto L76
L78:
	;
	if v241 <= int32(50) {
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+32)) = v257 + int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v117)+24))
	v262 = int32(6)
	v267 = F___memcpy(m, v261+v241*v262, v238, v236*v262)
	mBase = m.M
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v117)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v268 + v236
	v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+6)))
	v281 = int32(36)
	v282 = (v272&int32(8191)+int32(7))&int32(16376) | int32(4)
	goto L77
L81:
	;
	v281 = int32(4)
	v282 = int32(1)
	goto L77
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L20
	} else {
		goto L83
	}
L83:
	;
	F_errmsg_internal(m, int32(81535), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L20
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(524329), int32(526), int32(246394))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L20
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v307 = v172
	goto L67
L87:
	;
	goto L64
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+20)) = v347
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)) = uint16(v309)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v199
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+36)) = (v353&int32(8191)+int32(7))&int32(16376) | int32(4)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v117)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v117+v363<<(uint(int32(2))%32))+44)) = uint16(v309)
	goto L87
L89:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v117)+24))
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+2)))
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199))))
	v331 = int32(16)
	v336 = v314 & int32(4095)
	v339 = F___memcpy(m, v328, v199+(v329|v330<<(uint(v331)%32)), v336*int32(6))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v336
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+2)))
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199))))
	v347 = v341 | v342<<(uint(v331)%32)
	goto L88
L90:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+4)))
	if v314&int32(8192) != 0 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v117)+24))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v319
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v318)+4)) = uint16(v321)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = int32(1)
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+6)))
	v347 = v325 & int32(8191)
	goto L88
L93:
	;
	goto L92
L94:
	;
	goto L63
L95:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+12)))
	if v393&int32(64) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+16)))
	v397 = v140 + v396
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+12)))
	v400 = v398 & int32(65471)
	*(*uint16)(unsafe.Add(mBase, uint32(v397)+12)) = uint16(v400)
	goto L98
L97:
	;
	goto L98
L98:
	;
	F_PageRestoreTempPage(m, v140, v114)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L20
	} else {
		goto L99
	}
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v114))) = base.I64_rotr(v58, int64(32))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	F_MarkBufferDirty(m, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L20
	} else {
		goto L100
	}
L100:
	;
	goto L29
L101:
	;
	F_UnlockReleaseBuffer(m, v428)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L20
	} else {
		goto L102
	}
L102:
	;
	goto L4
L103:
	;
	if v440 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v444 = int32(0)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+72))
	if v448 < v444 {
		v470 = v444
		goto L108
	} else {
		goto L109
	}
L105:
	;
	goto L106
L106:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v527 == int32(0) {
		goto L4
	} else {
		goto L131
	}
L107:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v474 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L108:
	;
	goto L107
L109:
	;
	v454 = v447 + int32(76)
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	if v455 != int32(1) {
		v470 = v444
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+43)))
	if v458 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v470 = v444
	goto L108
L112:
	;
	goto L113
L113:
	;
	goto L116
L116:
	;
	goto L117
L117:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v454)+44))
	v470 = v468
	goto L108
L118:
	;
	v493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v433)+2)))
	if v493 != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v478+(v474^int32(-1))<<(uint(int32(2))%32))))
	v492 = v484
	goto L118
L120:
	;
	goto L121
L121:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v492 = v486 + v474<<(uint(int32(13))%32) + int32(-8192)
	goto L118
L122:
	;
	v494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v433))))
	v495 = int32(1)
	v497 = v470 + v494<<(uint(v495)%32)
	F_btree_xlog_updates(m, v492, v497, v497+v493<<(uint(v495)%32), v493)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L20
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v433))))
	if v504 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L124
L126:
	;
	F_PageIndexMultiDelete(m, v492, v470, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L20
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v507 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v492)+16)))
	v508 = v492 + v507
	v509 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v508)+14)) = uint16(v509)
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508)+12)))
	v513 = v511 & int32(65471)
	*(*uint16)(unsafe.Add(mBase, uint32(v508)+12)) = uint16(v513)
	*(*uint32)(unsafe.Add(mBase, uint32(v492)+4)) = uint32(v434)
	v517 = int64(base.Ui64(v434) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v492))) = uint32(v517)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	F_MarkBufferDirty(m, v519)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L20
	} else {
		goto L130
	}
L129:
	;
	goto L128
L130:
	;
	goto L106
L131:
	;
	F_UnlockReleaseBuffer(m, v527)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L20
	} else {
		goto L132
	}
L132:
	;
	goto L4
L133:
	;
	v538 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v538, v18+int32(-20), v538, v538)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L20
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v558 = F_XLogReadBufferForRedo(m, l0, int32(0), v18+int32(-20))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L20
	} else {
		goto L138
	}
L136:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+8)))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v547
	v549 = *(*int64)(unsafe.Add(mBase, uint32(v20)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v549
	F_ResolveRecoveryConflictWithSnapshot(m, v546, v545, v20)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L20
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	if v558 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v562 = int32(0)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+72))
	if v566 < v562 {
		v588 = v562
		goto L143
	} else {
		goto L144
	}
L140:
	;
	goto L141
L141:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v643 == int32(0) {
		goto L4
	} else {
		goto L166
	}
L142:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v592 < int32(0) {
		goto L154
	} else {
		goto L155
	}
L143:
	;
	goto L142
L144:
	;
	v572 = v565 + int32(76)
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572))))
	if v573 != int32(1) {
		v588 = v562
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+43)))
	if v576 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v588 = v562
	goto L143
L147:
	;
	goto L148
L148:
	;
	goto L151
L151:
	;
	goto L152
L152:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	v588 = v586
	goto L143
L153:
	;
	v611 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v532)+6)))
	if v611 != 0 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v596 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v596+(v592^int32(-1))<<(uint(int32(2))%32))))
	v610 = v602
	goto L153
L155:
	;
	goto L156
L156:
	;
	v604 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v610 = v604 + v592<<(uint(int32(13))%32) + int32(-8192)
	goto L153
L157:
	;
	v612 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v532)+4)))
	v613 = int32(1)
	v615 = v588 + v612<<(uint(v613)%32)
	F_btree_xlog_updates(m, v610, v615, v615+v611<<(uint(v613)%32), v611)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L20
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v532)+4)))
	if v622 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L159
L161:
	;
	F_PageIndexMultiDelete(m, v610, v588, v622)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L20
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v610)+16)))
	v626 = v610 + v625
	v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v626)+12)))
	v629 = v627 & int32(65471)
	*(*uint16)(unsafe.Add(mBase, uint32(v626)+12)) = uint16(v629)
	*(*uint32)(unsafe.Add(mBase, uint32(v610)+4)) = uint32(v533)
	v633 = int64(base.Ui64(v533) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v610))) = uint32(v633)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	F_MarkBufferDirty(m, v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L20
	} else {
		goto L165
	}
L164:
	;
	goto L163
L165:
	;
	goto L141
L166:
	;
	F_UnlockReleaseBuffer(m, v643)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L20
	} else {
		goto L167
	}
L167:
	;
	goto L4
L168:
	;
	if v653 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	if v657 < int32(0) {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	goto L171
L171:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	if v714 != 0 {
		goto L178
	} else {
		goto L179
	}
L172:
	;
	v677 = v675 + int32(24)
	v678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v648))))
	v679 = int32(2)
	v682 = int32(4)
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v677+v678<<(uint(v679)%32)-v682)))
	v685 = int32(32767)
	v691 = (v678 + int32(1)) & int32(65535)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v691<<(uint(v679)%32)+v677-v682)))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v675+v697&v685)))
	*(*int32)(unsafe.Add(mBase, uint32(v684&v685+v675))) = v701
	F_PageIndexTupleDelete(m, v675, v691)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L20
	} else {
		goto L176
	}
L173:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v661+(v657^int32(-1))<<(uint(int32(2))%32))))
	v675 = v667
	goto L172
L174:
	;
	goto L175
L175:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v675 = v669 + v657<<(uint(int32(13))%32) + int32(-8192)
	goto L172
L176:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v675))) = base.I64_rotr(v649, int64(32))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	F_MarkBufferDirty(m, v708)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L20
	} else {
		goto L177
	}
L177:
	;
	goto L171
L178:
	;
	F_UnlockReleaseBuffer(m, v714)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L20
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v718 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L20
	} else {
		goto L182
	}
L181:
	;
	goto L180
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v718
	if v718 < int32(0) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	F_PageInit(m, v738, int32(8192), int32(16))
	mBase = m.M
	goto L187
L184:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v724+(v718^int32(-1))<<(uint(int32(2))%32))))
	v738 = v730
	goto L183
L185:
	;
	goto L186
L186:
	;
	v732 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v738 = v732 + v718<<(uint(int32(13))%32) + int32(-8192)
	goto L183
L187:
	;
	v742 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v738)+16)))
	v743 = v738 + v742
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v648)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v743))) = v744
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v648)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v743)+8)) = int64(73014444032)
	*(*int32)(unsafe.Add(mBase, uint32(v743)+4)) = v746
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v648)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+46)) = uint16(v750)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = int32(537395200)
	v755 = int32(base.Ui32(v750) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+44)) = uint16(v755)
	v762 = F_PageAddItemExtended(m, v738, v18+int32(-20), int32(8), int32(1), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L20
	} else {
		goto L188
	}
L188:
	;
	if v762 == int32(0) {
		goto L2
	} else {
		goto L189
	}
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v738))) = base.I64_rotr(v649, int64(32))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	F_MarkBufferDirty(m, v769)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L20
	} else {
		goto L190
	}
L190:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	F_UnlockReleaseBuffer(m, v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L20
	} else {
		goto L191
	}
L191:
	;
	goto L4
L192:
	;
	v819 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L20
	} else {
		goto L204
	}
L193:
	;
	v784 = F_XLogReadBufferForRedo(m, l0, int32(1), v18+int32(-4))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L20
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = int32(0)
	goto L192
L196:
	;
	if v784 != 0 {
		goto L192
	} else {
		goto L197
	}
L197:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	if v786 < int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v805 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v804)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v805+v804)+4)) = v779
	*(*uint32)(unsafe.Add(mBase, uint32(v804)+4)) = uint32(v775)
	v810 = int64(base.Ui64(v775) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v804))) = uint32(v810)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	F_MarkBufferDirty(m, v812)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L20
	} else {
		goto L202
	}
L199:
	;
	v790 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v790+(v786^int32(-1))<<(uint(int32(2))%32))))
	v804 = v796
	goto L198
L200:
	;
	goto L201
L201:
	;
	v798 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v804 = v798 + v786<<(uint(int32(13))%32) + int32(-8192)
	goto L198
L202:
	;
	goto L192
L203:
	;
	F_PageInit(m, v838, int32(8192), int32(16))
	mBase = m.M
	goto L208
L204:
	;
	if v819 < int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v824+(v819^int32(-1))<<(uint(int32(2))%32))))
	v838 = v830
	goto L203
L206:
	;
	goto L207
L207:
	;
	v832 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v838 = v832 + v819<<(uint(int32(13))%32) + int32(-8192)
	goto L203
L208:
	;
	v842 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838)+16)))
	v843 = v838 + v842
	*(*int32)(unsafe.Add(mBase, uint32(v843)+8)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v843)+4)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v843))) = v780
	v847 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838)+16)))
	v848 = v838 + v847
	v849 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v848)+12)))
	v853 = v849&int32(65259) | int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v848)+12)) = uint16(v853)
	v855 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v838)+12)) = uint16(v855)
	*(*int64)(unsafe.Add(mBase, uint32(v838)+24)) = v777
	v858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v838)+14)) = uint16(v858)
	if v778 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v843)+12)))
	v864 = v862 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v843)+12)) = uint16(v864)
	goto L211
L210:
	;
	goto L211
L211:
	;
	v866 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v843)+14)) = uint16(v866)
	v868 = base.I32_wrap_i64(v775)
	*(*int32)(unsafe.Add(mBase, uint32(v838)+4)) = v868
	v872 = base.I32_wrap_i64(int64(base.Ui64(v775) >> (uint(int64(32)) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v838))) = v872
	F_MarkBufferDirty(m, v819)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L20
	} else {
		goto L212
	}
L212:
	;
	v879 = F_XLogReadBufferForRedo(m, l0, int32(2), v18+int32(-8))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L20
	} else {
		goto L213
	}
L213:
	;
	if v879 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if v883 < int32(0) {
		goto L218
	} else {
		goto L219
	}
L215:
	;
	goto L216
L216:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	if v911 != 0 {
		goto L222
	} else {
		goto L223
	}
L217:
	;
	v902 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v901)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v902+v901))) = v780
	*(*int32)(unsafe.Add(mBase, uint32(v901)+4)) = v868
	*(*int32)(unsafe.Add(mBase, uint32(v901))) = v872
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	F_MarkBufferDirty(m, v907)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L20
	} else {
		goto L221
	}
L218:
	;
	v887 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v887+(v883^int32(-1))<<(uint(int32(2))%32))))
	v901 = v893
	goto L217
L219:
	;
	goto L220
L220:
	;
	v895 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v901 = v895 + v883<<(uint(int32(13))%32) + int32(-8192)
	goto L217
L221:
	;
	goto L216
L222:
	;
	F_UnlockReleaseBuffer(m, v911)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L20
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if v914 != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L224
L226:
	;
	F_UnlockReleaseBuffer(m, v914)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L20
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	F_UnlockReleaseBuffer(m, v819)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L20
	} else {
		goto L230
	}
L229:
	;
	goto L228
L230:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v919)+72))
	if v920 < int32(3) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	if v31 != int32(144) {
		goto L4
	} else {
		goto L244
	}
L232:
	;
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+232)))
	if v923 != int32(1) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v927 = F_XLogInitBufferForRedo(m, l0, int32(3))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L20
	} else {
		goto L235
	}
L234:
	;
	F_PageInit(m, v946, int32(8192), int32(16))
	mBase = m.M
	goto L239
L235:
	;
	if v927 < int32(0) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v932 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v932+(v927^int32(-1))<<(uint(int32(2))%32))))
	v946 = v938
	goto L234
L237:
	;
	goto L238
L238:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v946 = v940 + v927<<(uint(int32(13))%32) + int32(-8192)
	goto L234
L239:
	;
	v950 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+16)))
	v951 = v946 + v950
	v952 = int32(17)
	*(*uint16)(unsafe.Add(mBase, uint32(v951)+12)) = uint16(v952)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v776)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v951))) = v954
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v776)+28))
	v957 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v951)+14)) = uint16(v957)
	*(*int32)(unsafe.Add(mBase, uint32(v951)+8)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v951)+4)) = v956
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v776)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+46)) = uint16(v962)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = int32(537395200)
	v967 = int32(base.Ui32(v962) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+44)) = uint16(v967)
	v974 = F_PageAddItemExtended(m, v946, v18+int32(-20), int32(8), int32(1), v957)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L20
	} else {
		goto L240
	}
L240:
	;
	if v974 == int32(0) {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v946)+4)) = v868
	*(*int32)(unsafe.Add(mBase, uint32(v946))) = v872
	F_MarkBufferDirty(m, v927)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L20
	} else {
		goto L242
	}
L242:
	;
	F_UnlockReleaseBuffer(m, v927)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L20
	} else {
		goto L243
	}
L243:
	;
	goto L231
L244:
	;
	F__bt_restore_meta(m, l0, int32(4))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	goto L4
L246:
	;
	F_PageInit(m, v1016, int32(8192), int32(16))
	mBase = m.M
	goto L251
L247:
	;
	if v997 < int32(0) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1002+(v997^int32(-1))<<(uint(int32(2))%32))))
	v1016 = v1008
	goto L246
L249:
	;
	goto L250
L250:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1016 = v1010 + v997<<(uint(int32(13))%32) + int32(-8192)
	goto L246
L251:
	;
	v1020 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1016)+16)))
	v1021 = v1016 + v1020
	v1022 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1021)+12)) = uint16(v1022)
	*(*int64)(unsafe.Add(mBase, uint32(v1021))) = int64(0)
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v994)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1021)+8)) = v1026
	if v1026 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1030 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v1021)+12)) = uint16(v1030)
	goto L254
L253:
	;
	goto L254
L254:
	;
	v1032 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1021)+14)) = uint16(v1032)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v994)+4))
	if v1034 == v1032 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1016))) = base.I64_rotr(v995, int64(32))
	F_MarkBufferDirty(m, v997)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L20
	} else {
		goto L280
	}
L256:
	;
	v1037 = int32(0)
	v1039 = v18 + int32(-4)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+72))
	if v1042 < v1037 {
		v1064 = v1037
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	F__bt_restore_page(m, v1016, v1067, v1068)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L20
	} else {
		goto L268
	}
L258:
	;
	v1067 = v1064
	goto L257
L259:
	;
	v1048 = v1041 + int32(76)
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048))))
	if v1049 != int32(1) {
		v1064 = v1037
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048)+43)))
	if v1052 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	if v1039 == int32(0) {
		v1064 = v1037
		goto L258
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	if v1039 != 0 {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v1057 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1039))) = v1057
	v1067 = v1057
	goto L257
L265:
	;
	v1060 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1048)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1039))) = v1060
	goto L267
L266:
	;
	goto L267
L267:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+44))
	v1064 = v1062
	goto L258
L268:
	;
	v1071 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1075 = F_XLogReadBufferForRedo(m, l0, int32(1), v18+int32(-20))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L20
	} else {
		goto L269
	}
L269:
	;
	if v1075 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v1079 < int32(0) {
		goto L274
	} else {
		goto L275
	}
L271:
	;
	goto L272
L272:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v1113 == int32(0) {
		goto L255
	} else {
		goto L278
	}
L273:
	;
	v1098 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1097)+16)))
	v1099 = v1098 + v1097
	v1100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1099)+12)))
	v1102 = v1100 & int32(65407)
	*(*uint16)(unsafe.Add(mBase, uint32(v1099)+12)) = uint16(v1102)
	*(*uint32)(unsafe.Add(mBase, uint32(v1097)+4)) = uint32(v1071)
	v1106 = int64(base.Ui64(v1071) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1097))) = uint32(v1106)
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	F_MarkBufferDirty(m, v1108)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L20
	} else {
		goto L277
	}
L274:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1083+(v1079^int32(-1))<<(uint(int32(2))%32))))
	v1097 = v1089
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1097 = v1091 + v1079<<(uint(int32(13))%32) + int32(-8192)
	goto L273
L277:
	;
	goto L272
L278:
	;
	F_UnlockReleaseBuffer(m, v1113)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L20
	} else {
		goto L279
	}
L279:
	;
	goto L255
L280:
	;
	F_UnlockReleaseBuffer(m, v997)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L20
	} else {
		goto L281
	}
L281:
	;
	F__bt_restore_meta(m, l0, int32(2))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L20
	} else {
		goto L282
	}
L282:
	;
	goto L4
L283:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(v1135)+16))
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1135)+24)))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v1138
	v1140 = *(*int64)(unsafe.Add(mBase, uint32(v1135)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v1140
	F_ResolveRecoveryConflictWithSnapshotFullXid(m, v1136, v1137, v18+int32(-48))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L20
	} else {
		goto L284
	}
L284:
	;
	goto L4
L285:
	;
	goto L4
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v31
	F_errmsg_internal(m, int32(57761), v18+int32(-32))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L20
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(524329), int32(1070), int32(255884))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L20
	} else {
		goto L288
	}
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L289:
	;
	goto L4
L290:
	;
	m.G0 = v20 - int32(-64)
	return
L291:
	;
	F_errmsg_internal(m, int32(21543), int32(0))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L20
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(524329), int32(508), int32(246394))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L20
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	F_errmsg_internal(m, int32(429185), int32(0))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L20
	} else {
		goto L295
	}
L295:
	;
	F_errfinish(m, int32(524329), int32(793), int32(487364))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L20
	} else {
		goto L296
	}
L296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L297:
	;
	F_errmsg_internal(m, int32(429185), int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L20
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(524329), int32(928), int32(427716))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L20
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_btree_xlog_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l3)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+96))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v68 = F_XLogReadBufferForRedo(m, l3, int32(0), v12+int32(8))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L16
	}
L2:
	;
	v20 = F_XLogReadBufferForRedo(m, l3, int32(1), v12+int32(12))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v24 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v58 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)))
	v44 = v43 + v42
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)))
	v47 = v45 & int32(65407)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v47)
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+4)) = uint32(v14)
	v51 = int64(base.Ui64(v14) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v42))) = uint32(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_MarkBufferDirty(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L12
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(v24^int32(-1))<<(uint(int32(2))%32))))
	v42 = v34
	goto L8
L10:
	;
	goto L11
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v42 = v36 + v24<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L12:
	;
	goto L7
L13:
	;
	F_UnlockReleaseBuffer(m, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L3
	} else {
		goto L61
	}
L16:
	;
	if v68 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v72 = int32(0)
	v74 = v12 + int32(4)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+96))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+72))
	if v77 < v72 {
		v99 = v72
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v194 != 0 {
		goto L53
	} else {
		goto L54
	}
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v103 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v102 = v99
	goto L20
L22:
	;
	v83 = v76 + int32(76)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v84 != int32(1) {
		v99 = v72
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+43)))
	if v87 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v74 == int32(0) {
		v99 = v72
		goto L21
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v74 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v92
	v102 = v92
	goto L20
L28:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v95
	goto L30
L29:
	;
	goto L30
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v83)+44))
	v99 = v97
	goto L21
L31:
	;
	if l2 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v103^int32(-1))<<(uint(int32(2))%32))))
	v121 = v113
	goto L31
L33:
	;
	goto L34
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v121 = v115 + v103<<(uint(int32(13))%32) + int32(-8192)
	goto L31
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = base.I64_rotr(v14, int64(32))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	F_MarkBufferDirty(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L52
	}
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16))))
	v127 = F_PageAddItemExtended(m, v121, v102, v124, v125, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102))))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v144 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v143 - v144
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16))))
	v155 = *(*int32)(unsafe.Add(mBase, uint32((v147-int32(1))&int32(65535)<<(uint(v144)%32)+v121)+20))
	v158 = v121 + v155&int32(32767)
	v161 = F_CopyIndexTuple(m, v102+v144)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L44
	}
L39:
	;
	if v127 != 0 {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errmsg_internal(m, int32(305882), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(524329), int32(191), int32(88108))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v163 = F__bt_swap_posting(m, v161, v158, v142)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+6)))
	v171 = (v165&int32(8191) + int32(7)) & int32(16376)
	if v171 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16))))
	v177 = F_PageAddItemExtended(m, v121, v161, v174, v175, int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L50
	}
L47:
	;
	v172 = F__emscripten_memcpy_bulkmem(m, v158, v163, v171)
	mBase = m.M
	goto L49
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	if v177 == int32(0) {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	goto L35
L52:
	;
	goto L19
L53:
	;
	F_UnlockReleaseBuffer(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if l1 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F__bt_restore_meta(m, l3, int32(2))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	m.G0 = v12 + int32(16)
	return
L60:
	;
	goto L59
L61:
	;
	F_errmsg_internal(m, int32(305845), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(524329), int32(230), int32(88108))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_btree_xlog_startup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v8 = F_AllocSetContextCreateInternal(m, v3, int32(65142), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[76])) = v8
		return
	}
}
