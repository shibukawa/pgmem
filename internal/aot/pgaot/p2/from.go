package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BeginCopyFrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int64
	_ = v589
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
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
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v771 int32
	_ = v771
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v925 int32
	_ = v925
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1176 int32
	_ = v1176
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1216 int64
	_ = v1216
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1383 int64
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1463 int32
	_ = v1463
	var v1468 int32
	_ = v1468
	v5 = l4
	v19 = m.G0
	v21 = v19 - int32(272)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+264)) = v24
	v27 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyFrom[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+256)) = v27
	v30 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyFrom[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+240)) = v30
	v33 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyFrom[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = v33
	v36 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyFrom[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+224)) = v36
	v39 = F_palloc0(m, int32(368))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[5]))
	v49 = F_AllocSetContextCreateInternal(m, v44, int32(_a_F_BeginCopyFrom_0), int32(0), int32(_a_F_BeginCopyFrom_1), int32(_a_F_BeginCopyFrom_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+212)) = v49
	v52 = int32(_a_F_BeginCopyFrom_3)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[5])) = v49
	v57 = v39 + int32(56)
	F_ProcessCopyOptions(m, l0, v57, int32(1), l7)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = l1
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v39)+60))
	if v65 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v68 = int32(_a_F_BeginCopyFrom_4)
	goto L7
L6:
	;
	v68 = int32(_a_F_BeginCopyFrom_5)
	goto L7
L7:
	;
	if v65 == int32(2) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v71 = int32(_a_F_BeginCopyFrom_6)
	goto L10
L9:
	;
	v71 = v68
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v74 = F_CopyGetAttnums(m, v73, l1, l6)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v78 = base.I32_extend16_s(v77)
	v79 = F_palloc0(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+124)) = v79
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+120)))
	if v82 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v321 = F_palloc0(m, v78)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L57
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+228)) = int32(0)
	goto L13
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L53
	}
L16:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v39)+148))
	if v195 == int32(0) {
		goto L14
	} else {
		goto L42
	}
L17:
	;
	if v78 == int32(0) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v39)+116))
	if v89 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L20:
	;
	base.MemoryFill(m, v79, int32(1), v78)
	goto L16
L21:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v93 = F_CopyGetAttnums(m, v73, v92, v89)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v93 == int32(0) {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v97 <= int32(0) {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v102 = int32(0)
	goto L25
L25:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v102<<(uint(int32(2))%32))))
	v125 = v123 - int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v128 = int32(0)
	if v127 == v128 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L16
L27:
	;
	if v166 == int32(0) {
		goto L15
	} else {
		goto L40
	}
L28:
	;
	v166 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v134 <= int32(0) {
		v160 = v128
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v166 = v160
	goto L27
L32:
	;
	v137 = int32(0)
	if v137 < v134 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v140 = v134
	goto L35
L34:
	;
	v140 = v137
	goto L35
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v143 = int32(0)
	goto L36
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v141+v143<<(uint(int32(2))%32))))
	v152 = base.B2i32(v151 == v123)
	if v151 == v123 {
		v160 = v152
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v160 = v152
	goto L31
L38:
	;
	v154 = v143 + int32(1)
	if v154 != v140 {
		v143 = v154
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v39)+124))
	v171 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v169+v125))) = uint8(v171)
	v174 = v102 + v171
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v174 < v175 {
		v102 = v174
		goto L25
	} else {
		goto L41
	}
L41:
	;
	goto L26
L42:
	;
	v199 = F_palloc0(m, int32(12))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+228)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = int32(453)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v39)+228))
	v205 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v204)+4)) = uint8(v205)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v39)+148))
	if base.Ui32(int32(2)) <= base.Ui32(v207-int32(1)) {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v39)+228))
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+5)) = uint8(v213)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v39)+148))
	if v215 != int32(2) {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v219 = F_palloc0_mul(m, int32(1), v78)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+256)) = v219
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	if v222 == int32(0) {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v225 <= int32(0) {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	v230 = int32(0)
	goto L49
L49:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v230<<(uint(int32(2))%32))))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v73+v247<<(uint(int32(3))%32)+v255*int32(100)-int32(4))))
	v262 = F_DomainHasConstraints(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	goto L13
L51:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v39)+256))
	v266 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v255+v264-v266))) = uint8(v262)
	v270 = v230 + v266
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v270 < v271 {
		v230 = v270
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	F_errcode(m, int32(_a_F_BeginCopyFrom_7))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = int32(_a_F_BeginCopyFrom_8)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v73 + v126<<(uint(int32(3))%32) + v125*int32(100) + int32(32)
	F_errmsg(m, int32(_a_F_BeginCopyFrom_9), v21+int32(112))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_10), int32(1618), int32(_a_F_BeginCopyFrom_11))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+140)) = v321
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+136)))
	if v324 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v934 = F_palloc0(m, v933)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L215
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L202
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L198
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L194
	}
L62:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+144)))
	if v437 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L63:
	;
	if v78 == int32(0) {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v39)+132))
	if v331 == int32(0) {
		goto L62
	} else {
		goto L67
	}
L66:
	;
	base.MemoryFill(m, v321, int32(1), v78)
	goto L62
L67:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v335 = F_CopyGetAttnums(m, v73, v334, v331)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v335 == int32(0) {
		goto L62
	} else {
		goto L69
	}
L69:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v339 <= int32(0) {
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v344 = int32(0)
	goto L71
L71:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361+v344<<(uint(int32(2))%32))))
	v367 = v365 - int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v370 = int32(0)
	if v369 == v370 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L62
L73:
	;
	if v408 == int32(0) {
		goto L61
	} else {
		goto L86
	}
L74:
	;
	v408 = int32(0)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v376 <= int32(0) {
		v402 = v370
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v408 = v402
	goto L73
L78:
	;
	v379 = int32(0)
	if v379 < v376 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v382 = v376
	goto L81
L80:
	;
	v382 = v379
	goto L81
L81:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v385 = int32(0)
	goto L82
L82:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v383+v385<<(uint(int32(2))%32))))
	v394 = base.B2i32(v393 == v365)
	if v393 == v365 {
		v402 = v394
		goto L77
	} else {
		goto L84
	}
L83:
	;
	v402 = v394
	goto L77
L84:
	;
	v396 = v385 + int32(1)
	if v396 != v382 {
		v385 = v396
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v39)+140))
	v413 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v411+v367))) = uint8(v413)
	v416 = v344 + v413
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v416 < v417 {
		v344 = v416
		goto L71
	} else {
		goto L87
	}
L87:
	;
	goto L72
L88:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v547 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L89:
	;
	v440 = F_palloc0(m, v78)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+176)) = v440
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v39)+168))
	v445 = F_CopyGetAttnums(m, v73, v443, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v445 == int32(0) {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	if v449 <= int32(0) {
		goto L88
	} else {
		goto L93
	}
L93:
	;
	v454 = int32(0)
	goto L94
L94:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v445)+12))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v471+v454<<(uint(int32(2))%32))))
	v477 = v475 - int32(1)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v480 = int32(0)
	if v479 == v480 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L88
L96:
	;
	if v518 == int32(0) {
		goto L60
	} else {
		goto L109
	}
L97:
	;
	v518 = int32(0)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v486 <= int32(0) {
		v512 = v480
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v518 = v512
	goto L96
L101:
	;
	v489 = int32(0)
	if v489 < v486 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v492 = v486
	goto L104
L103:
	;
	v492 = v489
	goto L104
L104:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v479)+12))
	v495 = int32(0)
	goto L105
L105:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v493+v495<<(uint(int32(2))%32))))
	v504 = base.B2i32(v503 == v475)
	if v503 == v475 {
		v512 = v504
		goto L100
	} else {
		goto L107
	}
L106:
	;
	v512 = v504
	goto L100
L107:
	;
	v506 = v495 + int32(1)
	if v506 != v492 {
		v495 = v506
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v39)+176))
	v523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v521+v477))) = uint8(v523)
	v526 = v454 + v523
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	if v526 < v527 {
		v454 = v526
		goto L94
	} else {
		goto L110
	}
L110:
	;
	goto L95
L111:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[6]))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+4))
	goto L114
L112:
	;
	v553 = v547
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v553
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[7]))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	goto L118
L114:
	;
	v553 = v552
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+180)) = l2
	v582 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v582
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)+48))
	v589 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+192)) = v589
	*(*int64)(unsafe.Add(mBase, uint32(v39)+200)) = v589
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+208)) = uint8(v582)
	v595 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+252)) = uint8(v595)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+184)) = v588 + int32(4)
	v601 = F_palloc(m, int32(_a_F_BeginCopyFrom_12))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L126
	}
L116:
	;
	v567 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+24)) = uint8(v567)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[7]))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	goto L123
L117:
	;
	v565 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+24)) = uint8(v565)
	goto L115
L118:
	;
	if v557 == v553 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if v559 == int32(0) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[7]))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	goto L121
L121:
	;
	if v564 != 0 {
		goto L116
	} else {
		goto L122
	}
L122:
	;
	goto L117
L123:
	;
	v573 = F_FindDefaultConversionProc(m, v569, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+28)) = v573
	if v573 == int32(0) {
		goto L59
	} else {
		goto L125
	}
L125:
	;
	goto L115
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+344)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+340)) = v601
	v606 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+352)) = uint8(v606)
	F_initStringInfo(m, v39+int32(280))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if l0 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+264)) = v612
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+268)) = v614
	goto L130
L129:
	;
	goto L130
L130:
	;
	v618 = F_palloc(m, v78*int32(28))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v621 = v77 << (uint(int32(16)) % 32) >> (uint(int32(14)) % 32)
	v622 = F_palloc(m, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v624 = F_palloc(m, v621)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v626 = F_palloc(m, v621)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	if v78 <= int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v916 = int32(0)
	v925 = v582
	goto L58
L136:
	;
	goto L137
L137:
	;
	v634 = int32(0)
	v639 = int32(1)
	v643 = v582
	goto L138
L138:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v657 = v73 + v651<<(uint(int32(3))%32) + v639*int32(100)
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657)+19)))
	if v658 != 0 {
		v806 = v634
		v809 = v643
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v916 = v806
	v925 = v809
	goto L58
L140:
	;
	if v639 != v78 {
		v634 = v806
		v639 = v639 + int32(1)
		v643 = v809
		goto L138
	} else {
		goto L193
	}
L141:
	;
	v660 = v657 - int32(72)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)+68))
	v663 = v639 - int32(1)
	v668 = v663 << (uint(int32(2)) % 32)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	m.T0[v671].(func(*base.Module, int32, int32, int32, int32))(m, v39, v661, v618+v663*int32(28), v622+v668)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v674 = v626 + v668
	v675 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v674))) = v675
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	if v677 == v675 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v681 = int32(0)
	if v680 == v681 {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	goto L145
L145:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+90)))
	if v720 != 0 {
		v806 = v634
		v809 = v643
		goto L140
	} else {
		goto L160
	}
L146:
	;
	if v719 != 0 {
		v806 = v634
		v809 = v643
		goto L140
	} else {
		goto L159
	}
L147:
	;
	v719 = int32(0)
	goto L146
L148:
	;
	goto L149
L149:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	if v687 <= int32(0) {
		v713 = v681
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v719 = v713
	goto L146
L151:
	;
	v690 = int32(0)
	if v690 < v687 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v693 = v687
	goto L154
L153:
	;
	v693 = v690
	goto L154
L154:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v680)+12))
	v696 = int32(0)
	goto L155
L155:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v694+v696<<(uint(int32(2))%32))))
	v705 = base.B2i32(v704 == v639)
	if v704 == v639 {
		v713 = v705
		goto L150
	} else {
		goto L157
	}
L156:
	;
	v713 = v705
	goto L150
L157:
	;
	v707 = v696 + int32(1)
	if v707 != v693 {
		v696 = v707
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	goto L145
L160:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v722 = F_build_column_default(m, v721, v639)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	if v722 == int32(0) {
		v806 = v634
		v809 = v643
		goto L140
	} else {
		goto L162
	}
L162:
	;
	v726 = F_expression_planner(m, v722)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v729 = F_ExecInitExpr(m, v726, int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v674))) = v729
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v733 = int32(0)
	if v732 == v733 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	if v771 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L166:
	;
	v771 = int32(0)
	goto L165
L167:
	;
	goto L168
L168:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v732)+4))
	if v739 <= int32(0) {
		v765 = v733
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v771 = v765
	goto L165
L170:
	;
	v742 = int32(0)
	if v742 < v739 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v745 = v739
	goto L173
L172:
	;
	v745 = v742
	goto L173
L173:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v732)+12))
	v748 = int32(0)
	goto L174
L174:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v746+v748<<(uint(int32(2))%32))))
	v757 = base.B2i32(v756 == v639)
	if v756 == v639 {
		v765 = v757
		goto L169
	} else {
		goto L176
	}
L175:
	;
	v765 = v757
	goto L169
L176:
	;
	v759 = v748 + int32(1)
	if v759 != v745 {
		v748 = v759
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v624+base.I32_extend16_s(v634)<<(uint(int32(2))%32)))) = v663
	v781 = v634 + int32(1)
	goto L180
L179:
	;
	v781 = v634
	goto L180
L180:
	;
	if v643&int32(1) != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v806 = v781
	v809 = int32(1)
	goto L140
L182:
	;
	goto L183
L183:
	;
	v785 = int32(0)
	if v726 == v785 {
		v805 = v785
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v806 = v781
	v809 = v805
	goto L140
L185:
	;
	v791 = F_check_functions_in_node(m, v726, int32(910), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	if v791 != 0 {
		v805 = int32(1)
		goto L184
	} else {
		goto L187
	}
L187:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	if v793 == int32(67) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v797 = int32(0)
	v799 = F_query_tree_walker_impl(m, v726, int32(911), v797, v797)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v803 = F_expression_tree_walker_impl(m, v726, int32(911), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L192
	}
L191:
	;
	v805 = v799
	goto L184
L192:
	;
	v805 = v803
	goto L184
L193:
	;
	goto L139
L194:
	;
	F_errcode(m, int32(_a_F_BeginCopyFrom_7))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = int32(_a_F_BeginCopyFrom_13)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v73 + v368<<(uint(int32(3))%32) + v367*int32(100) + int32(32)
	F_errmsg(m, int32(_a_F_BeginCopyFrom_9), v21+int32(96))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_10), int32(1678), int32(_a_F_BeginCopyFrom_11))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	F_errcode(m, int32(_a_F_BeginCopyFrom_7))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v73 + v478<<(uint(int32(3))%32) + v477*int32(100) + int32(32)
	F_errmsg_internal(m, int32(_a_F_BeginCopyFrom_14), v21+int32(80))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_10), int32(1702), int32(_a_F_BeginCopyFrom_11))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if base.B2i32(v875 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v875)) != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[7]))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)+4))
	goto L208
L205:
	;
	v887 = int32(_a_F_BeginCopyFrom_15)
	goto L207
L206:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v875<<(uint(int32(3))%32))+uint32(_c_F_BeginCopyFrom[8])))
	v887 = v886
	goto L207
L207:
	;
	goto L204
L208:
	;
	if base.B2i32(v890 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v890)) != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v902
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v887
	F_errmsg(m, int32(_a_F_BeginCopyFrom_16), v21-int32(-64))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L213
	}
L210:
	;
	v902 = int32(_a_F_BeginCopyFrom_15)
	goto L212
L211:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v890<<(uint(int32(3))%32))+uint32(_c_F_BeginCopyFrom[8])))
	v902 = v901
	goto L212
L212:
	;
	goto L209
L213:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_10), int32(1732), int32(_a_F_BeginCopyFrom_11))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+248)) = v934
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	if v938 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v938)+56))
	v941 = v939
	goto L218
L217:
	;
	v941 = int32(0)
	goto L218
L218:
	;
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[9]))
	if v944 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v989 = v925 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+260)) = uint8(v989)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+244)) = v626
	*(*int32)(unsafe.Add(mBase, uint32(v39)+240)) = v624
	*(*int32)(unsafe.Add(mBase, uint32(v39)+224)) = v622
	*(*int32)(unsafe.Add(mBase, uint32(v39)+220)) = v618
	*(*int64)(unsafe.Add(mBase, uint32(v39)+360)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+216)) = uint16(v916)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+44)) = uint8(v5)
	if l5 != 0 {
		goto L226
	} else {
		goto L227
	}
L220:
	;
	goto L219
L221:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BeginCopyFrom[10])))
	if v948&int32(1) == int32(0) {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v953 = int32(_a_F_BeginCopyFrom_17)
	v955 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11]))
	v956 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11])) = v955 + v956
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v944)))
	*(*int32)(unsafe.Add(mBase, uint32(v944))) = v959 + v956
	v963 = int32(0)
	v965 = int32(_a_F_BeginCopyFrom_18)
	v966 = base.AtomicRmwOr32(m, v963, v965, v963)
	*(*int32)(unsafe.Add(mBase, uint32(v944)+220)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v944)+224)) = v941
	base.MemoryFill(m, v944+int32(232), v963, int32(160))
	v977 = base.AtomicRmwOr32(m, v963, v965, v963)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v944)))
	*(*int32)(unsafe.Add(mBase, uint32(v944))) = v978 + v956
	v984 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11])) = v984 - v956
	goto L220
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L308
	}
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L304
	}
L225:
	;
	goto L288
L226:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = int64(4)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = int32(2)
	goto L225
L227:
	;
	goto L228
L228:
	;
	if l3 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = int64(3)
	v1009 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[12]))
	if v1009 == int32(2) {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L231
L231:
	;
	v1126 = F_pstrdup(m, l3)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L254
	}
L232:
	;
	v1012 = m.G0
	v1014 = v1012 - int32(16)
	m.G0 = v1014
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	if v1016 != 0 {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	goto L234
L234:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v1124
	goto L225
L235:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+4))
	v1019 = v1017
	goto L237
L236:
	;
	v1019 = int32(0)
	goto L237
L237:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v39)+60))
	F_pq_beginmessage(m, v1014, int32(71))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_enlargeStringInfo(m, v1014, int32(1))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+4))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	v1030 = int32(1)
	v1031 = base.B2i32(v1020 == v1030)
	*(*uint8)(unsafe.Add(mBase, uint32(v1027+v1028))) = uint8(v1031)
	*(*int32)(unsafe.Add(mBase, uint32(v1014)+4)) = v1027 + v1030
	F_enlargeStringInfo(m, v1014, int32(2))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+4))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	v1042 = int32(8)
	v1048 = v1019<<(uint(v1042)%32) | int32(base.Ui32(v1019&int32(_a_F_BeginCopyFrom_19))>>(uint(v1042)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v1039+v1040))) = uint16(v1048)
	*(*int32)(unsafe.Add(mBase, uint32(v1014)+4)) = v1039 + int32(2)
	if int32(0) < v1019 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1055 = int32(0)
	if v1020 == v1030 {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	goto L243
L243:
	;
	F_pq_endmessage(m, v1014)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L1
	} else {
		goto L251
	}
L244:
	;
	v1058 = int32(256)
	goto L246
L245:
	;
	v1058 = v1055
	goto L246
L246:
	;
	v1060 = v1055
	goto L247
L247:
	;
	F_enlargeStringInfo(m, v1014, int32(2))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L1
	} else {
		goto L249
	}
L248:
	;
	goto L243
L249:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+4))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1080+v1081))) = uint16(v1058)
	*(*int32)(unsafe.Add(mBase, uint32(v1014)+4)) = v1080 + int32(2)
	v1088 = v1060 + int32(1)
	if v1088 != v1019 {
		v1060 = v1088
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = int32(1)
	v1112 = F_makeStringInfo(m)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v1112
	v1116 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[14]))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+4))
	v1118 = m.T0[v1117].(func(*base.Module) int32)(m)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	m.G0 = v1014 + int32(16)
	goto L225
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v1126
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+44)))
	if v1129 == int32(1) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = int64(2)
	v1135 = F_OpenPipeStream(m, v1126, int32(_a_F_BeginCopyFrom_20))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = int64(1)
	v1157 = F_AllocateFile(m, v1126, int32(_a_F_BeginCopyFrom_20))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L264
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v1135
	if v1135 != 0 {
		goto L225
	} else {
		goto L259
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1144
	F_errmsg(m, int32(_a_F_BeginCopyFrom_21), v21)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_10), int32(1888), int32(_a_F_BeginCopyFrom_11))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v1157
	if v1157 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[15]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+60))
	if v1193 < int32(0) {
		goto L277
	} else {
		goto L278
	}
L268:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v1170
	F_errmsg(m, int32(_a_F_BeginCopyFrom_22), v21+int32(16))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	if base.B2i32(v1163 != int32(44))&base.B2i32(v1163 != int32(2)) == int32(0) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	F_errhint(m, int32(_a_F_BeginCopyFrom_23), int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_10), int32(1907), int32(_a_F_BeginCopyFrom_11))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L275
	}
L274:
	;
	goto L273
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	if v1200 < int32(0) {
		goto L281
	} else {
		goto L282
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[15])) = int32(8)
	v1200 = int32(-1)
	goto L279
L278:
	;
	v1200 = v1193
	goto L279
L279:
	;
	goto L276
L280:
	;
	if v1210 != 0 {
		goto L224
	} else {
		goto L284
	}
L281:
	;
	v1206 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v1210 = v1206
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1209 = F___fstatat(m, v1200, int32(_a_F_BeginCopyFrom_15), v21+int32(128), int32(_a_F_BeginCopyFrom_24))
	mBase = m.M
	v1210 = v1209
	goto L280
L284:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v21)+132))
	if v1211&int32(_a_F_BeginCopyFrom_25) == int32(_a_F_BeginCopyFrom_26) {
		goto L223
	} else {
		goto L285
	}
L285:
	;
	v1216 = *(*int64)(unsafe.Add(mBase, uint32(v21)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+240)) = v1216
	goto L225
L286:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+4))
	m.T0[v1423].(func(*base.Module, int32, int32))(m, v39, v73)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L1
	} else {
		goto L303
	}
L287:
	;
	goto L286
L288:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[9]))
	if v1250 == int32(0) {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BeginCopyFrom[10])))
	if v1254&int32(1) == int32(0) {
		goto L287
	} else {
		goto L290
	}
L290:
	;
	v1259 = int32(_a_F_BeginCopyFrom_17)
	v1261 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11]))
	v1262 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11])) = v1261 + v1262
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1250)))
	*(*int32)(unsafe.Add(mBase, uint32(v1250))) = v1265 + v1262
	v1269 = int32(0)
	v1272 = base.AtomicRmwOr32(m, v1269, int32(_a_F_BeginCopyFrom_18), v1269)
	goto L292
L291:
	;
	v1399 = int32(0)
	v1402 = base.AtomicRmwOr32(m, v1399, int32(_a_F_BeginCopyFrom_18), v1399)
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1250)))
	v1404 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1250))) = v1403 + v1404
	v1407 = int32(_a_F_BeginCopyFrom_17)
	v1409 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11])) = v1409 - v1404
	goto L287
L292:
	;
	goto L294
L294:
	;
	goto L295
L295:
	;
	v1364 = int32(0)
	v1367 = int32(0)
	goto L300
L300:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(256)+v1367<<(uint(int32(2))%32))))
	v1377 = int32(3)
	v1383 = *(*int64)(unsafe.Add(mBase, uint32(v21+int32(224)+v1367<<(uint(v1377)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1250+int32(232)+v1376<<(uint(v1377)%32)))) = v1383
	v1385 = int32(1)
	v1388 = v1364 + v1385
	if v1388 != int32(3) {
		v1364 = v1388
		v1367 = v1367 + v1385
		goto L300
	} else {
		goto L302
	}
L301:
	;
	goto L291
L302:
	;
	goto L301
L303:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[5])) = v53
	m.G0 = v21 + int32(272)
	return v39
L304:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v1438
	F_errmsg(m, int32(_a_F_BeginCopyFrom_27), v21+int32(48))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_10), int32(1914), int32(_a_F_BeginCopyFrom_11))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v1457
	F_errmsg(m, int32(_a_F_BeginCopyFrom_28), v21+int32(32))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_10), int32(1919), int32(_a_F_BeginCopyFrom_11))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyFromBinaryOneRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v170 int32
	_ = v170
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
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int64
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
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
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int64
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v504 int64
	_ = v504
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+4)))
	v31 = v29
	goto L3
L2:
	;
	v31 = int32(0)
	goto L3
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v35 + int64(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v39-v40 <= int32(1) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L21
	} else {
		goto L132
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L21
	} else {
		goto L128
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L21
	} else {
		goto L124
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L21
	} else {
		goto L120
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L21
	} else {
		goto L116
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v110 + int32(1)
	goto L4
L10:
	;
	m.G0 = v25 + int32(16)
	return v517
L11:
	;
	v131 = int32(_a_F_CopyFromBinaryOneRow_0)
	if v113&v131 == v131 {
		goto L32
	} else {
		goto L33
	}
L12:
	;
	v47 = v40
	v50 = int32(0)
	v52 = v39
	v53 = v25 + int32(10)
	goto L16
L13:
	;
	goto L14
L14:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103+v40))))
	v107 = v40 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v107
	v110 = v107
	v113 = v105
	v114 = v103
	v115 = v39
	goto L11
L15:
	;
	v517 = int32(0)
	goto L10
L16:
	;
	if v47 == v52 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v89 != int32(2) {
		goto L15
	} else {
		goto L31
	}
L18:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v76 = v47
	v77 = v52
	goto L20
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v80 = int32(2) - v50
	v81 = v77 - v76
	if v80 < v81 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	return int32(0)
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v73 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v76 = v75
	v77 = v74
	goto L20
L24:
	;
	v83 = v80
	goto L26
L25:
	;
	v83 = v81
	goto L26
L26:
	;
	if v83 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	base.MemoryCopy(m, v53, v76+v78, v83)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v86 = v76 + v83
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v86
	v89 = v50 + v83
	if v89 < int32(2) {
		v47 = v86
		v50 = v89
		v52 = v77
		v53 = v53 + v83
		goto L16
	} else {
		goto L30
	}
L30:
	;
	goto L17
L31:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+10)))
	v110 = v86
	v113 = v94
	v114 = v78
	v115 = v77
	goto L11
L32:
	;
	v135 = int32(0)
	if v135 < v115-v110 {
		goto L9
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v187 = int32(8)
	v194 = base.I32_extend16_s(v113<<(uint(v187)%32) | int32(base.Ui32(v113&int32(_a_F_CopyFromBinaryOneRow_1))>>(uint(v187)%32)))
	if v31 != v194 {
		goto L8
	} else {
		goto L51
	}
L35:
	;
	v142 = v110
	v145 = v135
	v146 = v114
	v147 = v115
	v148 = v25 + int32(9)
	goto L36
L36:
	;
	if v142 == v147 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v517 = int32(0)
	goto L10
L38:
	;
	goto L37
L39:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	v170 = v142
	v171 = v146
	v172 = v147
	goto L41
L41:
	;
	v174 = int32(1) - v145
	v175 = v172 - v170
	if v174 < v175 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v166 != 0 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v170 = v169
	v171 = v167
	v172 = v168
	goto L41
L44:
	;
	v177 = v174
	goto L46
L45:
	;
	v177 = v175
	goto L46
L46:
	;
	if v177 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	base.MemoryCopy(m, v148, v170+v171, v177)
	goto L49
L48:
	;
	goto L49
L49:
	;
	v180 = v170 + v177
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v180
	v183 = v177 + v145
	if v183 <= int32(0) {
		v142 = v180
		v145 = v183
		v146 = v171
		v147 = v172
		v148 = v177 + v148
		goto L36
	} else {
		goto L50
	}
L50:
	;
	goto L4
L51:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v196 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v517 = int32(1)
	goto L10
L53:
	;
	goto L54
L54:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v201 <= int32(0) {
		v517 = int32(1)
		goto L10
	} else {
		goto L55
	}
L55:
	;
	v205 = l0 + int32(280)
	v218 = int32(0)
	goto L56
L56:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v234 = int32(2)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v233+v218<<(uint(v234)%32))))
	v240 = v34 + v229<<(uint(int32(3))%32) + v237*int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v240 - int32(68)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v246 = v237 - int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v32+v246<<(uint(v234)%32))))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v253-v254 < int32(4) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v517 = v511
	goto L10
L58:
	;
	v356 = l3 + v246
	v359 = v33 + v246*int32(28)
	if v338 == int32(-1) {
		goto L83
	} else {
		goto L84
	}
L59:
	;
	v260 = v254
	v263 = v25 + int32(12)
	v265 = int32(0)
	v266 = v253
	goto L63
L60:
	;
	goto L61
L61:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v328+v254)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v254 + int32(4)
	v338 = v330
	goto L58
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L21
	} else {
		goto L78
	}
L63:
	;
	if v260 == v266 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v300 != int32(4) {
		goto L62
	} else {
		goto L77
	}
L65:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L21
	} else {
		goto L68
	}
L66:
	;
	v287 = v260
	v288 = v266
	goto L67
L67:
	;
	v290 = int32(4) - v265
	v291 = v288 - v287
	if v290 < v291 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v284 != 0 {
		goto L62
	} else {
		goto L69
	}
L69:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v287 = v286
	v288 = v285
	goto L67
L70:
	;
	v293 = v290
	goto L72
L71:
	;
	v293 = v291
	goto L72
L72:
	;
	if v293 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	base.MemoryCopy(m, v263, v294+v287, v293)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v297 = v287 + v293
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v297
	v300 = v293 + v265
	if v300 < int32(4) {
		v260 = v297
		v263 = v263 + v293
		v265 = v300
		v266 = v288
		goto L63
	} else {
		goto L76
	}
L76:
	;
	goto L64
L77:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v338 = v305
	goto L58
L78:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L21
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_2), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(2313), int32(_a_F_CopyFromBinaryOneRow_4))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L21
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2+v246<<(uint(int32(3))%32)))) = v504
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = int32(0)
	v511 = int32(1)
	v513 = v218 + v511
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v513 < v514 {
		v218 = v513
		goto L56
	} else {
		goto L115
	}
L83:
	;
	v362 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v356))) = uint8(v362)
	v365 = F_ReceiveFunctionCall(m, v359, int32(0), v250, v244)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L21
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v369 = int32(16711935)
	v375 = base.I32_rotr(v338, int32(24))&v369 | base.I32_rotr(v338&v369, int32(8))
	if v375 < int32(0) {
		goto L7
	} else {
		goto L87
	}
L86:
	;
	v504 = v365
	goto L82
L87:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v379 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v378))) = uint8(v379)
	*(*int32)(unsafe.Add(mBase, uint32(v205)+12)) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v205)+4)) = v379
	goto L88
L88:
	;
	F_enlargeStringInfo(m, v205, v375)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L21
	} else {
		goto L89
	}
L89:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v375 <= v389-v390 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v375
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v474 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v472+v375))) = uint8(v474)
	v476 = F_ReceiveFunctionCall(m, v359, v205, v250, v244)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L21
	} else {
		goto L113
	}
L91:
	;
	if v375 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v403 = v390
	v405 = int32(0)
	v406 = v387
	goto L97
L94:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	base.MemoryCopy(m, v387, v393+v390, v375)
	goto L96
L95:
	;
	goto L96
L96:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v396 + v375
	goto L90
L97:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v403 == v421 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	if v375 != v445 {
		goto L6
	} else {
		goto L112
	}
L99:
	;
	goto L98
L100:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L21
	} else {
		goto L103
	}
L101:
	;
	v428 = v421
	v429 = v403
	goto L102
L102:
	;
	v430 = v375 - v405
	v431 = v428 - v429
	if v430 < v431 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v425 != 0 {
		v445 = v405
		goto L99
	} else {
		goto L104
	}
L104:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v428 = v427
	v429 = v426
	goto L102
L105:
	;
	v433 = v430
	goto L107
L106:
	;
	v433 = v431
	goto L107
L107:
	;
	if v433 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	base.MemoryCopy(m, v406, v434+v429, v433)
	goto L110
L109:
	;
	goto L110
L110:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v438 = v437 + v433
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v438
	v441 = v433 + v405
	if v441 < v375 {
		v403 = v438
		v405 = v441
		v406 = v433 + v406
		goto L97
	} else {
		goto L111
	}
L111:
	;
	v445 = v441
	goto L99
L112:
	;
	goto L90
L113:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v478 != v479 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v481 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v356))) = uint8(v481)
	v504 = v476
	goto L82
L115:
	;
	goto L57
L116:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L21
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v194
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_5), v25)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L21
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(1233), int32(_a_F_CopyFromBinaryOneRow_6))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L21
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L21
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_7), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(2322), int32(_a_F_CopyFromBinaryOneRow_4))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L21
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L21
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_2), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L21
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(2332), int32(_a_F_CopyFromBinaryOneRow_4))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L21
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L21
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_8), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L21
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(2345), int32(_a_F_CopyFromBinaryOneRow_4))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L21
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L21
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_9), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L21
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(1225), int32(_a_F_CopyFromBinaryOneRow_6))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L21
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyFromErrorCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v145 int32
	_ = v145
	v7 = m.G0
	v9 = v7 - int32(176)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v11 == int32(1) {
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
			F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_0), v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				m.G0 = v9 + int32(176)
				return
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		if v23 == int32(1) {
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
				if v22 != 0 {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v31
					*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v30
					F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_1), v9+int32(32))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v9 + int32(176)
						return
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v30
					F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_2), v9+int32(16))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						m.G0 = v9 + int32(176)
						return
					}
				}
			}
		} else {
			if v22 != 0 {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				if v47 != 0 {
					v48 = F_strlen(m, v47)
					mBase = m.M
					if v48 <= int32(100) {
						v51 = F_pstrdup(m, v47)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v64 = v51
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
								v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v71
								*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v70
								*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v69
								F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_3), v9+int32(144))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									F_pfree(m, v64)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										m.G0 = v9 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v54 = F_pg_mbcliplen(m, v47, v48, int32(100))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v58 = F_palloc(m, v54+int32(4))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								if v54 != 0 {
									base.MemoryCopy(m, v58, v47, v54)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v58+v54))) = int32(_a_F_CopyFromErrorCallback_4)
								v64 = v58
								F_set_errcontext_domain(m, int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
									v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v64
									*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v71
									*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v70
									*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v69
									F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_3), v9+int32(144))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										F_pfree(m, v64)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											m.G0 = v9 + int32(176)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
						v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v88
						*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v87
						*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v86
						F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_5), v9+int32(112))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							m.G0 = v9 + int32(176)
							return
						}
					}
				}
			} else {
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
				if v97 == int32(1) {
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
					v101 = F_strlen(m, v100)
					mBase = m.M
					if v101 <= int32(100) {
						v104 = F_pstrdup(m, v100)
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							v117 = v104
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
								v123 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v117
								*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v123
								*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v122
								F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_6), v9-int32(-64))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return
								} else {
									F_pfree(m, v117)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										m.G0 = v9 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v107 = F_pg_mbcliplen(m, v100, v101, int32(100))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return
						} else {
							v111 = F_palloc(m, v107+int32(4))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return
							} else {
								if v107 != 0 {
									base.MemoryCopy(m, v111, v100, v107)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v111+v107))) = int32(_a_F_CopyFromErrorCallback_4)
								v117 = v111
								F_set_errcontext_domain(m, int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
									v123 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v117
									*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v123
									*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v122
									F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_6), v9-int32(-64))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return
									} else {
										F_pfree(m, v117)
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return
										} else {
											m.G0 = v9 + int32(176)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return
					} else {
						v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
						v138 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v138
						*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v137
						F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_2), v9+int32(96))
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return
						} else {
							m.G0 = v9 + int32(176)
							return
						}
					}
				}
			}
		}
	}
}
func F_CopyFromTextLikeInFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_getTypeInputInfo(m, l1, v7+int32(12), l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		F_fmgr_info(m, v13, l2)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_RemoveFromWaitQueue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+384))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	*(*int64)(unsafe.Add(mBase, uint32(l0)+388)) = int64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v21 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v20 - v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v24 - v21
	v30 = v10 + v12<<(uint(int32(2))%32)
	v32 = v30 + int32(44)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = v33 - v21
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	if v35 == v37 {
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v39 & base.I32_rotl(int32(-2), v12)
	} else {
	}
	v44 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+416)) = v44
	v46 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(l0)+384)) = v46
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(v44)%32))+uint32(_c_F_RemoveFromWaitQueue[0])))
	F_CleanUpLock(m, v10, v9, v52, l1, int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		return
	} else {
		return
	}
}
