package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_DeadLockCheckRecurse(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
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
	var v119 int32
	_ = v119
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v2 = int32(0)
	v11 = F_TestConfiguration(m, l0)
	mBase = m.M
	if v2 <= v11 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L11
	} else {
		goto L26
	}
L2:
	;
	return v146
L3:
	;
	if v11 == int32(0) {
		v146 = v2
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v146 = int32(1)
	goto L2
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[0]))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[1]))
	if v20 <= v18 {
		v146 = int32(1)
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[2]))
	v24 = v23 + v11
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[3]))
	v27 = v24 + v26
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[4]))
	if v27 <= v29 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[2])) = v24
	goto L10
L9:
	;
	goto L10
L10:
	;
	v33 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[5]))
	v36 = int32(20)
	v38 = v35 + v18*v36
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[6]))
	v43 = v40 + v23*v36
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v48
	v50 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[0])) = v18 + v50
	v55 = F_DeadLockCheckRecurse(m, l0)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v55 == int32(0) {
		v146 = v33
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v61 = int32(_a_F_DeadLockCheckRecurse_0)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[0]))
	v64 = int32(1)
	v65 = v63 - v64
	*(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[0])) = v65
	if v11 != v64 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v70 = v65
	v72 = v50
	goto L17
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[2])) = v23
	goto L5
L17:
	;
	if v29 < v27 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	v80 = F_TestConfiguration(m, l0)
	mBase = m.M
	if v80 != v11 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v84 = v70
	goto L21
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[5]))
	v87 = int32(20)
	v89 = v86 + v84*v87
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[6]))
	v97 = v91 + v23*v87 + v72*v87
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+16)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v97)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v102
	*(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[0])) = v84 + int32(1)
	v108 = F_DeadLockCheckRecurse(m, l0)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L23
	}
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[0]))
	v84 = v83
	goto L21
L23:
	;
	if v108 == int32(0) {
		v146 = v33
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v112 = int32(_a_F_DeadLockCheckRecurse_0)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[0]))
	v115 = int32(1)
	v116 = v114 - v115
	*(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[0])) = v116
	v119 = v72 + v115
	if v11 != v119 {
		v70 = v116
		v72 = v119
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	F_errmsg_internal(m, int32(_a_F_DeadLockCheckRecurse_1), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_DeadLockCheckRecurse_2), int32(348), int32(_a_F_DeadLockCheckRecurse_3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DisownLatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	return
}
func F___divtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v109 int64
	_ = v109
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v142 int64
	_ = v142
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v187 int64
	_ = v187
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v244 int64
	_ = v244
	var v251 int64
	_ = v251
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v265 int64
	_ = v265
	var v266 int64
	_ = v266
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v279 int64
	_ = v279
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v287 int64
	_ = v287
	var v294 int64
	_ = v294
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v310 int64
	_ = v310
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v334 int64
	_ = v334
	var v341 int64
	_ = v341
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v366 int64
	_ = v366
	var v369 int64
	_ = v369
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v373 int64
	_ = v373
	var v377 int64
	_ = v377
	var v384 int64
	_ = v384
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v400 int64
	_ = v400
	var v403 int64
	_ = v403
	var v404 int64
	_ = v404
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v413 int64
	_ = v413
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v419 int64
	_ = v419
	var v420 int64
	_ = v420
	var v424 int64
	_ = v424
	var v431 int64
	_ = v431
	var v443 int32
	_ = v443
	var v444 int64
	_ = v444
	var v446 int64
	_ = v446
	var v447 int64
	_ = v447
	var v453 int64
	_ = v453
	var v454 int64
	_ = v454
	var v456 int64
	_ = v456
	var v459 int64
	_ = v459
	var v460 int64
	_ = v460
	var v462 int64
	_ = v462
	var v463 int64
	_ = v463
	var v467 int64
	_ = v467
	var v474 int64
	_ = v474
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v490 int64
	_ = v490
	var v493 int64
	_ = v493
	var v494 int64
	_ = v494
	var v500 int64
	_ = v500
	var v501 int64
	_ = v501
	var v503 int64
	_ = v503
	var v506 int64
	_ = v506
	var v507 int64
	_ = v507
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v514 int64
	_ = v514
	var v521 int64
	_ = v521
	var v533 int32
	_ = v533
	var v534 int64
	_ = v534
	var v536 int64
	_ = v536
	var v537 int64
	_ = v537
	var v543 int64
	_ = v543
	var v544 int64
	_ = v544
	var v546 int64
	_ = v546
	var v549 int64
	_ = v549
	var v550 int64
	_ = v550
	var v552 int64
	_ = v552
	var v553 int64
	_ = v553
	var v557 int64
	_ = v557
	var v564 int64
	_ = v564
	var v576 int32
	_ = v576
	var v577 int64
	_ = v577
	var v578 int64
	_ = v578
	var v579 int64
	_ = v579
	var v581 int64
	_ = v581
	var v586 int64
	_ = v586
	var v592 int64
	_ = v592
	var v593 int64
	_ = v593
	var v595 int64
	_ = v595
	var v598 int64
	_ = v598
	var v599 int64
	_ = v599
	var v601 int64
	_ = v601
	var v602 int64
	_ = v602
	var v606 int64
	_ = v606
	var v613 int64
	_ = v613
	var v625 int32
	_ = v625
	var v627 int64
	_ = v627
	var v628 int64
	_ = v628
	var v634 int64
	_ = v634
	var v635 int64
	_ = v635
	var v637 int64
	_ = v637
	var v640 int64
	_ = v640
	var v641 int64
	_ = v641
	var v643 int64
	_ = v643
	var v644 int64
	_ = v644
	var v648 int64
	_ = v648
	var v655 int64
	_ = v655
	var v667 int32
	_ = v667
	var v668 int64
	_ = v668
	var v670 int64
	_ = v670
	var v671 int64
	_ = v671
	var v672 int64
	_ = v672
	var v673 int64
	_ = v673
	var v681 int64
	_ = v681
	var v687 int64
	_ = v687
	var v688 int64
	_ = v688
	var v690 int64
	_ = v690
	var v693 int64
	_ = v693
	var v694 int64
	_ = v694
	var v696 int64
	_ = v696
	var v697 int64
	_ = v697
	var v701 int64
	_ = v701
	var v708 int64
	_ = v708
	var v720 int32
	_ = v720
	var v722 int64
	_ = v722
	var v723 int64
	_ = v723
	var v729 int64
	_ = v729
	var v730 int64
	_ = v730
	var v732 int64
	_ = v732
	var v735 int64
	_ = v735
	var v736 int64
	_ = v736
	var v738 int64
	_ = v738
	var v739 int64
	_ = v739
	var v743 int64
	_ = v743
	var v750 int64
	_ = v750
	var v762 int32
	_ = v762
	var v763 int64
	_ = v763
	var v764 int64
	_ = v764
	var v765 int64
	_ = v765
	var v766 int64
	_ = v766
	var v769 int64
	_ = v769
	var v770 int64
	_ = v770
	var v773 int64
	_ = v773
	var v775 int64
	_ = v775
	var v776 int64
	_ = v776
	var v777 int64
	_ = v777
	var v779 int64
	_ = v779
	var v781 int64
	_ = v781
	var v783 int64
	_ = v783
	var v784 int64
	_ = v784
	var v786 int64
	_ = v786
	var v788 int64
	_ = v788
	var v793 int64
	_ = v793
	var v805 int64
	_ = v805
	var v807 int64
	_ = v807
	var v809 int64
	_ = v809
	var v812 int64
	_ = v812
	var v813 int64
	_ = v813
	var v815 int64
	_ = v815
	var v820 int64
	_ = v820
	var v822 int64
	_ = v822
	var v828 int64
	_ = v828
	var v830 int64
	_ = v830
	var v841 int64
	_ = v841
	var v846 int64
	_ = v846
	var v847 int64
	_ = v847
	var v849 int64
	_ = v849
	var v853 int64
	_ = v853
	var v855 int64
	_ = v855
	var v859 int64
	_ = v859
	var v863 int64
	_ = v863
	var v865 int64
	_ = v865
	var v867 int64
	_ = v867
	var v869 int64
	_ = v869
	var v883 int64
	_ = v883
	var v887 int64
	_ = v887
	var v889 int64
	_ = v889
	var v897 int64
	_ = v897
	var v906 int64
	_ = v906
	var v909 int64
	_ = v909
	var v914 int32
	_ = v914
	var v919 int64
	_ = v919
	var v920 int64
	_ = v920
	var v922 int64
	_ = v922
	var v925 int64
	_ = v925
	var v926 int64
	_ = v926
	var v928 int64
	_ = v928
	var v929 int64
	_ = v929
	var v933 int64
	_ = v933
	var v940 int64
	_ = v940
	var v953 int64
	_ = v953
	var v955 int64
	_ = v955
	var v956 int64
	_ = v956
	var v965 int32
	_ = v965
	var v968 int64
	_ = v968
	var v970 int64
	_ = v970
	var v972 int64
	_ = v972
	var v977 int64
	_ = v977
	var v978 int64
	_ = v978
	var v980 int64
	_ = v980
	var v983 int64
	_ = v983
	var v984 int64
	_ = v984
	var v986 int64
	_ = v986
	var v987 int64
	_ = v987
	var v991 int64
	_ = v991
	var v998 int64
	_ = v998
	var v1011 int64
	_ = v1011
	var v1013 int64
	_ = v1013
	var v1014 int64
	_ = v1014
	var v1023 int64
	_ = v1023
	var v1024 int64
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int64
	_ = v1027
	var v1028 int64
	_ = v1028
	var v1029 int64
	_ = v1029
	var v1030 int64
	_ = v1030
	var v1038 int64
	_ = v1038
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1072 int64
	_ = v1072
	var v1076 int64
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1095 int64
	_ = v1095
	var v1103 int64
	_ = v1103
	var v1104 int64
	_ = v1104
	var v1109 int32
	_ = v1109
	var v1110 int64
	_ = v1110
	var v1111 int64
	_ = v1111
	var v1116 int64
	_ = v1116
	var v1117 int64
	_ = v1117
	var v1119 int64
	_ = v1119
	var v1122 int64
	_ = v1122
	var v1123 int64
	_ = v1123
	var v1125 int64
	_ = v1125
	var v1126 int64
	_ = v1126
	var v1130 int64
	_ = v1130
	var v1137 int64
	_ = v1137
	var v1148 int64
	_ = v1148
	var v1149 int64
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1152 int64
	_ = v1152
	var v1157 int64
	_ = v1157
	var v1159 int64
	_ = v1159
	var v1164 int64
	_ = v1164
	var v1165 int64
	_ = v1165
	var v1168 int64
	_ = v1168
	var v1169 int64
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int64
	_ = v1172
	var v1173 int64
	_ = v1173
	var v1178 int64
	_ = v1178
	var v1181 int64
	_ = v1181
	var v1184 int64
	_ = v1184
	var v1187 int64
	_ = v1187
	var v1188 int64
	_ = v1188
	var v1192 int64
	_ = v1192
	var v1199 int64
	_ = v1199
	var v1210 int64
	_ = v1210
	var v1211 int64
	_ = v1211
	var v1216 int64
	_ = v1216
	var v1219 int64
	_ = v1219
	var v1222 int64
	_ = v1222
	var v1225 int64
	_ = v1225
	var v1226 int64
	_ = v1226
	var v1230 int64
	_ = v1230
	var v1237 int64
	_ = v1237
	var v1249 int64
	_ = v1249
	var v1250 int64
	_ = v1250
	var v1254 int64
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1259 int64
	_ = v1259
	var v1262 int64
	_ = v1262
	var v1265 int64
	_ = v1265
	var v1267 int64
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1273 int64
	_ = v1273
	var v1276 int64
	_ = v1276
	var v1279 int64
	_ = v1279
	var v1281 int64
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1287 int64
	_ = v1287
	var v1292 int64
	_ = v1292
	var v1301 int64
	_ = v1301
	v6 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(336)
	m.G0 = v28
	v30 = int64(281474976710655)
	v31 = l4 & v30
	v33 = l2 & v30
	v36 = (l2 ^ l4) & int64(-9223372036854775807-1)
	v37 = int64(48)
	v40 = int32(_a_F___divtf3_0)
	v41 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(v37)%64))) & v40
	v46 = base.I32_wrap_i64(int64(base.Ui64(l2)>>(uint(v37)%64))) & v40
	if base.Ui32(int32(-32766)) <= base.Ui32(v46-v40) {
		if base.Ui32(int32(-32767)) < base.Ui32(v41-int32(_a_F___divtf3_0)) {
			v205 = l1
			v207 = l3
			v208 = v6
			v210 = v33
			v211 = v31
			v214 = v28 + int32(288)
			v216 = v211 | int64(281474976710656)
			v221 = v216<<(uint(int64(15))%64) | int64(base.Ui64(v207)>>(uint(int64(49))%64))
			v222 = int64(0)
			v224 = int64(8432131802713292800) - v221
			v230 = int64(32)
			v231 = int64(base.Ui64(v224) >> (uint(v230) % 64))
			v233 = int64(base.Ui64(v221) >> (uint(v230) % 64))
			v236 = int64(4294967295)
			v237 = v224 & v236
			v239 = v221 & v236
			v240 = v237 * v239
			v244 = int64(base.Ui64(v240)>>(uint(v230)%64)) + v237*v233
			v251 = v239*v231 + v244&v236
			*(*int64)(unsafe.Add(mBase, uint32(v214)+8)) = v221*v222 + v222*v224 + v231*v233 + int64(base.Ui64(v244)>>(uint(v230)%64)) + int64(base.Ui64(v251)>>(uint(v230)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v214))) = v240&v236 | v251<<(uint(v230)%64)
			v263 = v28 + int32(272)
			v264 = int64(0)
			v265 = *(*int64)(unsafe.Add(mBase, uint32(v28)+296))
			v266 = v264 - v265
			v273 = int64(32)
			v274 = int64(base.Ui64(v224) >> (uint(v273) % 64))
			v276 = int64(base.Ui64(v266) >> (uint(v273) % 64))
			v279 = int64(4294967295)
			v280 = v224 & v279
			v282 = v266 & v279
			v283 = v280 * v282
			v287 = int64(base.Ui64(v283)>>(uint(v273)%64)) + v280*v276
			v294 = v282*v274 + v287&v279
			*(*int64)(unsafe.Add(mBase, uint32(v263)+8)) = v266*v264 + v264*v224 + v274*v276 + int64(base.Ui64(v287)>>(uint(v273)%64)) + int64(base.Ui64(v294)>>(uint(v273)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v263))) = v283&v279 | v294<<(uint(v273)%64)
			v306 = v28 + int32(256)
			v307 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
			v310 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
			v313 = v307<<(uint(int64(1))%64) | int64(base.Ui64(v310)>>(uint(int64(63))%64))
			v314 = int64(0)
			v320 = int64(32)
			v321 = int64(base.Ui64(v221) >> (uint(v320) % 64))
			v323 = int64(base.Ui64(v313) >> (uint(v320) % 64))
			v326 = int64(4294967295)
			v327 = v221 & v326
			v329 = v313 & v326
			v330 = v327 * v329
			v334 = int64(base.Ui64(v330)>>(uint(v320)%64)) + v327*v323
			v341 = v329*v321 + v334&v326
			*(*int64)(unsafe.Add(mBase, uint32(v306)+8)) = v313*v314 + v314*v221 + v321*v323 + int64(base.Ui64(v334)>>(uint(v320)%64)) + int64(base.Ui64(v341)>>(uint(v320)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v306))) = v330&v326 | v341<<(uint(v320)%64)
			v353 = v28 + int32(240)
			v354 = int64(0)
			v356 = *(*int64)(unsafe.Add(mBase, uint32(v28)+264))
			v357 = v354 - v356
			v363 = int64(32)
			v364 = int64(base.Ui64(v357) >> (uint(v363) % 64))
			v366 = int64(base.Ui64(v313) >> (uint(v363) % 64))
			v369 = int64(4294967295)
			v370 = v357 & v369
			v372 = v313 & v369
			v373 = v370 * v372
			v377 = int64(base.Ui64(v373)>>(uint(v363)%64)) + v370*v366
			v384 = v372*v364 + v377&v369
			*(*int64)(unsafe.Add(mBase, uint32(v353)+8)) = v313*v354 + v354*v357 + v364*v366 + int64(base.Ui64(v377)>>(uint(v363)%64)) + int64(base.Ui64(v384)>>(uint(v363)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v353))) = v373&v369 | v384<<(uint(v363)%64)
			v396 = v28 + int32(224)
			v397 = *(*int64)(unsafe.Add(mBase, uint32(v28)+248))
			v400 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
			v403 = v397<<(uint(int64(1))%64) | int64(base.Ui64(v400)>>(uint(int64(63))%64))
			v404 = int64(0)
			v410 = int64(32)
			v411 = int64(base.Ui64(v221) >> (uint(v410) % 64))
			v413 = int64(base.Ui64(v403) >> (uint(v410) % 64))
			v416 = int64(4294967295)
			v417 = v221 & v416
			v419 = v403 & v416
			v420 = v417 * v419
			v424 = int64(base.Ui64(v420)>>(uint(v410)%64)) + v417*v413
			v431 = v419*v411 + v424&v416
			*(*int64)(unsafe.Add(mBase, uint32(v396)+8)) = v403*v404 + v404*v221 + v411*v413 + int64(base.Ui64(v424)>>(uint(v410)%64)) + int64(base.Ui64(v431)>>(uint(v410)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v396))) = v420&v416 | v431<<(uint(v410)%64)
			v443 = v28 + int32(208)
			v444 = int64(0)
			v446 = *(*int64)(unsafe.Add(mBase, uint32(v28)+232))
			v447 = v444 - v446
			v453 = int64(32)
			v454 = int64(base.Ui64(v447) >> (uint(v453) % 64))
			v456 = int64(base.Ui64(v403) >> (uint(v453) % 64))
			v459 = int64(4294967295)
			v460 = v447 & v459
			v462 = v403 & v459
			v463 = v460 * v462
			v467 = int64(base.Ui64(v463)>>(uint(v453)%64)) + v460*v456
			v474 = v462*v454 + v467&v459
			*(*int64)(unsafe.Add(mBase, uint32(v443)+8)) = v403*v444 + v444*v447 + v454*v456 + int64(base.Ui64(v467)>>(uint(v453)%64)) + int64(base.Ui64(v474)>>(uint(v453)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v443))) = v463&v459 | v474<<(uint(v453)%64)
			v486 = v28 + int32(192)
			v487 = *(*int64)(unsafe.Add(mBase, uint32(v28)+216))
			v490 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
			v493 = v487<<(uint(int64(1))%64) | int64(base.Ui64(v490)>>(uint(int64(63))%64))
			v494 = int64(0)
			v500 = int64(32)
			v501 = int64(base.Ui64(v221) >> (uint(v500) % 64))
			v503 = int64(base.Ui64(v493) >> (uint(v500) % 64))
			v506 = int64(4294967295)
			v507 = v221 & v506
			v509 = v493 & v506
			v510 = v507 * v509
			v514 = int64(base.Ui64(v510)>>(uint(v500)%64)) + v507*v503
			v521 = v509*v501 + v514&v506
			*(*int64)(unsafe.Add(mBase, uint32(v486)+8)) = v493*v494 + v494*v221 + v501*v503 + int64(base.Ui64(v514)>>(uint(v500)%64)) + int64(base.Ui64(v521)>>(uint(v500)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v486))) = v510&v506 | v521<<(uint(v500)%64)
			v533 = v28 + int32(176)
			v534 = int64(0)
			v536 = *(*int64)(unsafe.Add(mBase, uint32(v28)+200))
			v537 = v534 - v536
			v543 = int64(32)
			v544 = int64(base.Ui64(v537) >> (uint(v543) % 64))
			v546 = int64(base.Ui64(v493) >> (uint(v543) % 64))
			v549 = int64(4294967295)
			v550 = v537 & v549
			v552 = v493 & v549
			v553 = v550 * v552
			v557 = int64(base.Ui64(v553)>>(uint(v543)%64)) + v550*v546
			v564 = v552*v544 + v557&v549
			*(*int64)(unsafe.Add(mBase, uint32(v533)+8)) = v493*v534 + v534*v537 + v544*v546 + int64(base.Ui64(v557)>>(uint(v543)%64)) + int64(base.Ui64(v564)>>(uint(v543)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v533))) = v553&v549 | v564<<(uint(v543)%64)
			v576 = v28 + int32(160)
			v577 = int64(0)
			v578 = *(*int64)(unsafe.Add(mBase, uint32(v28)+184))
			v579 = int64(1)
			v581 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
			v586 = v578<<(uint(v579)%64) | int64(base.Ui64(v581)>>(uint(int64(63))%64)) - v579
			v592 = int64(32)
			v593 = int64(base.Ui64(v586) >> (uint(v592) % 64))
			v595 = int64(base.Ui64(v221) >> (uint(v592) % 64))
			v598 = int64(4294967295)
			v599 = v586 & v598
			v601 = v221 & v598
			v602 = v599 * v601
			v606 = int64(base.Ui64(v602)>>(uint(v592)%64)) + v599*v595
			v613 = v601*v593 + v606&v598
			*(*int64)(unsafe.Add(mBase, uint32(v576)+8)) = v221*v577 + v577*v586 + v593*v595 + int64(base.Ui64(v606)>>(uint(v592)%64)) + int64(base.Ui64(v613)>>(uint(v592)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v576))) = v602&v598 | v613<<(uint(v592)%64)
			v625 = v28 + int32(144)
			v627 = v207 << (uint(int64(15)) % 64)
			v628 = int64(0)
			v634 = int64(32)
			v635 = int64(base.Ui64(v586) >> (uint(v634) % 64))
			v637 = int64(base.Ui64(v627) >> (uint(v634) % 64))
			v640 = int64(4294967295)
			v641 = v586 & v640
			v643 = v627 & v640
			v644 = v641 * v643
			v648 = int64(base.Ui64(v644)>>(uint(v634)%64)) + v641*v637
			v655 = v643*v635 + v648&v640
			*(*int64)(unsafe.Add(mBase, uint32(v625)+8)) = v627*v628 + v628*v586 + v635*v637 + int64(base.Ui64(v648)>>(uint(v634)%64)) + int64(base.Ui64(v655)>>(uint(v634)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v625))) = v644&v640 | v655<<(uint(v634)%64)
			v667 = v28 + int32(112)
			v668 = int64(0)
			v670 = *(*int64)(unsafe.Add(mBase, uint32(v28)+168))
			v671 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
			v672 = *(*int64)(unsafe.Add(mBase, uint32(v28)+152))
			v673 = v671 + v672
			v681 = v668 - (v670 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v673) < base.Ui64(v671))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(1)) < base.Ui64(v673))))
			v687 = int64(32)
			v688 = int64(base.Ui64(v681) >> (uint(v687) % 64))
			v690 = int64(base.Ui64(v586) >> (uint(v687) % 64))
			v693 = int64(4294967295)
			v694 = v681 & v693
			v696 = v586 & v693
			v697 = v694 * v696
			v701 = int64(base.Ui64(v697)>>(uint(v687)%64)) + v694*v690
			v708 = v696*v688 + v701&v693
			*(*int64)(unsafe.Add(mBase, uint32(v667)+8)) = v586*v668 + v668*v681 + v688*v690 + int64(base.Ui64(v701)>>(uint(v687)%64)) + int64(base.Ui64(v708)>>(uint(v687)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v667))) = v697&v693 | v708<<(uint(v687)%64)
			v720 = v28 + int32(128)
			v722 = int64(1) - v673
			v723 = int64(0)
			v729 = int64(32)
			v730 = int64(base.Ui64(v586) >> (uint(v729) % 64))
			v732 = int64(base.Ui64(v722) >> (uint(v729) % 64))
			v735 = int64(4294967295)
			v736 = v586 & v735
			v738 = v722 & v735
			v739 = v736 * v738
			v743 = int64(base.Ui64(v739)>>(uint(v729)%64)) + v736*v732
			v750 = v738*v730 + v743&v735
			*(*int64)(unsafe.Add(mBase, uint32(v720)+8)) = v722*v723 + v723*v586 + v730*v732 + int64(base.Ui64(v743)>>(uint(v729)%64)) + int64(base.Ui64(v750)>>(uint(v729)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v720))) = v739&v735 | v750<<(uint(v729)%64)
			v762 = v208 + (v46 - v41)
			v763 = *(*int64)(unsafe.Add(mBase, uint32(v28)+112))
			v764 = int64(1)
			v765 = v763 << (uint(v764) % 64)
			v766 = *(*int64)(unsafe.Add(mBase, uint32(v28)+136))
			v769 = *(*int64)(unsafe.Add(mBase, uint32(v28)+128))
			v770 = int64(63)
			v773 = v765 + (v766<<(uint(v764)%64) | int64(base.Ui64(v769)>>(uint(v770)%64)))
			v775 = v773 - int64(13927)
			v776 = int64(32)
			v777 = int64(base.Ui64(v775) >> (uint(v776) % 64))
			v779 = v210 | int64(281474976710656)
			v781 = v779 << (uint(v764) % 64)
			v783 = int64(base.Ui64(v781) >> (uint(v776) % 64))
			v784 = v777 * v783
			v786 = v205 << (uint(v764) % 64)
			v788 = int64(base.Ui64(v786) >> (uint(v776) % 64))
			v793 = *(*int64)(unsafe.Add(mBase, uint32(v28)+120))
			v805 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v775) < base.Ui64(v773))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v773) < base.Ui64(v765))) + (v793<<(uint(v764)%64) | int64(base.Ui64(v763)>>(uint(v770)%64)) + int64(base.Ui64(v766)>>(uint(v770)%64)))) - v764
			v807 = int64(base.Ui64(v805) >> (uint(v776) % 64))
			v809 = v784 + v788*v807
			v812 = int64(4294967295)
			v813 = v805 & v812
			v815 = int64(base.Ui64(v205) >> (uint(v770) % 64))
			v820 = (v815 | v210<<(uint(v764)%64)) & v812
			v822 = v809 + v813*v820
			v828 = v783 * v813
			v830 = v828 + v820*v807
			v841 = v822 + v830<<(uint(v776)%64)
			v846 = v775 & v812
			v847 = v846 * v820
			v849 = v847 + v777*v788
			v853 = v786 & int64(4294967294)
			v855 = v849 + v813*v853
			v859 = v841 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v849) < base.Ui64(v847))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v855) < base.Ui64(v849))))
			v863 = v783 * v846
			v865 = v863 + v807*v853
			v867 = v865 + v777*v820
			v869 = v867 + v788*v813
			v883 = v859 + (int64(base.Ui64(v869)>>(uint(v776)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v869) < base.Ui64(v867)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v865) < base.Ui64(v863)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v867) < base.Ui64(v865)))))<<(uint(v776)%64))
			v887 = v777 * v853
			v889 = v887 + v788*v846
			v897 = v855 + (int64(base.Ui64(v889)>>(uint(v776)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v889) < base.Ui64(v887)))<<(uint(v776)%64))
			v906 = v883 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v897) < base.Ui64(v855))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v897+v869<<(uint(v776)%64)) < base.Ui64(v897))))
			v909 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v809) < base.Ui64(v784))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v822) < base.Ui64(v809))) + v783*v807 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v830) < base.Ui64(v828)))<<(uint(v776)%64) | int64(base.Ui64(v830)>>(uint(v776)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v841) < base.Ui64(v822))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v859) < base.Ui64(v841))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v883) < base.Ui64(v859))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v906) < base.Ui64(v883)))
			if base.Ui64(v909) <= base.Ui64(int64(562949953421311)) {
				v914 = v28 + int32(80)
				v919 = int64(32)
				v920 = int64(base.Ui64(v207) >> (uint(v919) % 64))
				v922 = int64(base.Ui64(v906) >> (uint(v919) % 64))
				v925 = int64(4294967295)
				v926 = v207 & v925
				v928 = v906 & v925
				v929 = v926 * v928
				v933 = int64(base.Ui64(v929)>>(uint(v919)%64)) + v926*v922
				v940 = v928*v920 + v933&v925
				*(*int64)(unsafe.Add(mBase, uint32(v914)+8)) = v906*v216 + v909*v207 + v920*v922 + int64(base.Ui64(v933)>>(uint(v919)%64)) + int64(base.Ui64(v940)>>(uint(v919)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v914))) = v929&v925 | v940<<(uint(v919)%64)
				v953 = *(*int64)(unsafe.Add(mBase, uint32(v28)+88))
				v955 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
				v956 = int64(0)
				v1023 = v906
				v1024 = v909
				v1025 = v762 + int32(_a_F___divtf3_1)
				v1027 = v205<<(uint(int64(49))%64) - v953 - base.I64_extend_i32_u(base.B2i32(v955 != v956))
				v1028 = v786
				v1029 = v781 | v815
				v1030 = v956 - v955
			} else {
				v965 = v28 + int32(96)
				v968 = int64(1)
				v970 = v909<<(uint(int64(63))%64) | int64(base.Ui64(v906)>>(uint(v968)%64))
				v972 = int64(base.Ui64(v909) >> (uint(v968) % 64))
				v977 = int64(32)
				v978 = int64(base.Ui64(v207) >> (uint(v977) % 64))
				v980 = int64(base.Ui64(v970) >> (uint(v977) % 64))
				v983 = int64(4294967295)
				v984 = v207 & v983
				v986 = v970 & v983
				v987 = v984 * v986
				v991 = int64(base.Ui64(v987)>>(uint(v977)%64)) + v984*v980
				v998 = v986*v978 + v991&v983
				*(*int64)(unsafe.Add(mBase, uint32(v965)+8)) = v970*v216 + v972*v207 + v978*v980 + int64(base.Ui64(v991)>>(uint(v977)%64)) + int64(base.Ui64(v998)>>(uint(v977)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v965))) = v987&v983 | v998<<(uint(v977)%64)
				v1011 = *(*int64)(unsafe.Add(mBase, uint32(v28)+104))
				v1013 = *(*int64)(unsafe.Add(mBase, uint32(v28)+96))
				v1014 = int64(0)
				v1023 = v970
				v1024 = v972
				v1025 = v762 + int32(_a_F___divtf3_2)
				v1027 = v205<<(uint(int64(48))%64) - v1011 - base.I64_extend_i32_u(base.B2i32(v1013 != v1014))
				v1028 = v205
				v1029 = v779
				v1030 = v1014 - v1013
			}
			if int32(_a_F___divtf3_0) <= v1025 {
				v1292 = int64(0)
				v1301 = v36 | int64(9223090561878065152)
			} else {
				if int32(0) < v1025 {
					v1038 = int64(1)
					v1164 = v1027<<(uint(v1038)%64) | int64(base.Ui64(v1030)>>(uint(int64(63))%64))
					v1165 = v1023
					v1168 = v1024&int64(281474976710655) | base.I64_extend_i32_u(v1025)<<(uint(int64(48))%64)
					v1169 = v1030 << (uint(v1038) % 64)
					v1171 = v28 + int32(16)
					v1172 = int64(3)
					v1173 = int64(0)
					v1178 = int64(32)
					v1181 = int64(base.Ui64(v207) >> (uint(v1178) % 64))
					v1184 = int64(4294967295)
					v1187 = v207 & v1184
					v1188 = v1172 * v1187
					v1192 = int64(base.Ui64(v1188)>>(uint(v1178)%64)) + v1172*v1181
					v1199 = v1187*v1173 + v1192&v1184
					*(*int64)(unsafe.Add(mBase, uint32(v1171)+8)) = v207*v1173 + v216*v1172 + v1173*v1181 + int64(base.Ui64(v1192)>>(uint(v1178)%64)) + int64(base.Ui64(v1199)>>(uint(v1178)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v1171))) = v1188&v1184 | v1199<<(uint(v1178)%64)
					v1210 = int64(5)
					v1211 = int64(0)
					v1216 = int64(32)
					v1219 = int64(base.Ui64(v207) >> (uint(v1216) % 64))
					v1222 = int64(4294967295)
					v1225 = v207 & v1222
					v1226 = v1210 * v1225
					v1230 = int64(base.Ui64(v1226)>>(uint(v1216)%64)) + v1210*v1219
					v1237 = v1225*v1211 + v1230&v1222
					*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v207*v1211 + v216*v1210 + v1211*v1219 + int64(base.Ui64(v1230)>>(uint(v1216)%64)) + int64(base.Ui64(v1237)>>(uint(v1216)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1226&v1222 | v1237<<(uint(v1216)%64)
					v1249 = v1165 & int64(1)
					v1250 = v1249 + v1169
					v1254 = v1164 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1250) < base.Ui64(v1249)))
					if v1254 == v216 {
						v1257 = base.B2i32(base.Ui64(v207) < base.Ui64(v1250))
					} else {
						v1257 = base.B2i32(base.Ui64(v216) < base.Ui64(v1254))
					}
					v1259 = v1165 + base.I64_extend_i32_u(v1257)
					v1262 = v1168 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1259) < base.Ui64(v1165)))
					v1265 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
					v1267 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
					if v1254 == v1267 {
						v1270 = base.B2i32(base.Ui64(v1265) < base.Ui64(v1250))
					} else {
						v1270 = base.B2i32(base.Ui64(v1267) < base.Ui64(v1254))
					}
					v1273 = v1259 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1262) < base.Ui64(int64(9223090561878065152)))&v1270)
					v1276 = v1262 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1273) < base.Ui64(v1259)))
					v1279 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
					v1281 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
					if v1254 == v1281 {
						v1284 = base.B2i32(base.Ui64(v1279) < base.Ui64(v1250))
					} else {
						v1284 = base.B2i32(base.Ui64(v1281) < base.Ui64(v1254))
					}
					v1287 = v1273 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1276) < base.Ui64(int64(9223090561878065152)))&v1284)
					v1292 = v1287
					v1301 = v1276 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1287) < base.Ui64(v1273))) | v36
				} else {
					if v1025 <= int32(-113) {
						v1292 = int64(0)
						v1301 = v36
					} else {
						v1055 = v28 - int32(-64)
						v1057 = int32(1) - v1025
						if v1057&int32(64) != 0 {
							v1076 = int64(base.Ui64(v1024) >> (uint(base.I64_extend_i32_u(v1057+int32(-64))) % 64))
							v1077 = int64(0)
						} else {
							if v1057 == int32(0) {
								v1076 = v1023
								v1077 = v1024
							} else {
								v1072 = base.I64_extend_i32_u(v1057)
								v1076 = v1024<<(uint(base.I64_extend_i32_u(int32(64)-v1057))%64) | int64(base.Ui64(v1023)>>(uint(v1072)%64))
								v1077 = int64(base.Ui64(v1024) >> (uint(v1072) % 64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v1055))) = v1076
						*(*int64)(unsafe.Add(mBase, uint32(v1055)+8)) = v1077
						v1082 = v28 + int32(48)
						v1084 = v1025 + int32(112)
						if v1084&int32(64) != 0 {
							v1103 = int64(0)
							v1104 = v1028 << (uint(base.I64_extend_i32_u(v1025+int32(48))) % 64)
						} else {
							if v1084 == int32(0) {
								v1103 = v1028
								v1104 = v1029
							} else {
								v1095 = base.I64_extend_i32_u(v1084)
								v1103 = v1028 << (uint(v1095) % 64)
								v1104 = v1029<<(uint(v1095)%64) | int64(base.Ui64(v1028)>>(uint(base.I64_extend_i32_u(int32(64)-v1084))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v1082))) = v1103
						*(*int64)(unsafe.Add(mBase, uint32(v1082)+8)) = v1104
						v1109 = v28 + int32(32)
						v1110 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
						v1111 = *(*int64)(unsafe.Add(mBase, uint32(v28)+72))
						v1116 = int64(32)
						v1117 = int64(base.Ui64(v1110) >> (uint(v1116) % 64))
						v1119 = int64(base.Ui64(v207) >> (uint(v1116) % 64))
						v1122 = int64(4294967295)
						v1123 = v1110 & v1122
						v1125 = v207 & v1122
						v1126 = v1123 * v1125
						v1130 = int64(base.Ui64(v1126)>>(uint(v1116)%64)) + v1123*v1119
						v1137 = v1125*v1117 + v1130&v1122
						*(*int64)(unsafe.Add(mBase, uint32(v1109)+8)) = v207*v1111 + v216*v1110 + v1117*v1119 + int64(base.Ui64(v1130)>>(uint(v1116)%64)) + int64(base.Ui64(v1137)>>(uint(v1116)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v1109))) = v1126&v1122 | v1137<<(uint(v1116)%64)
						v1148 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
						v1149 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
						v1150 = int64(1)
						v1152 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
						v1157 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
						v1159 = v1152 << (uint(v1150) % 64)
						v1164 = v1148 - (v1149<<(uint(v1150)%64) | int64(base.Ui64(v1152)>>(uint(int64(63))%64))) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1157) < base.Ui64(v1159)))
						v1165 = v1110
						v1168 = v1111
						v1169 = v1157 - v1159
						v1171 = v28 + int32(16)
						v1172 = int64(3)
						v1173 = int64(0)
						v1178 = int64(32)
						v1181 = int64(base.Ui64(v207) >> (uint(v1178) % 64))
						v1184 = int64(4294967295)
						v1187 = v207 & v1184
						v1188 = v1172 * v1187
						v1192 = int64(base.Ui64(v1188)>>(uint(v1178)%64)) + v1172*v1181
						v1199 = v1187*v1173 + v1192&v1184
						*(*int64)(unsafe.Add(mBase, uint32(v1171)+8)) = v207*v1173 + v216*v1172 + v1173*v1181 + int64(base.Ui64(v1192)>>(uint(v1178)%64)) + int64(base.Ui64(v1199)>>(uint(v1178)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v1171))) = v1188&v1184 | v1199<<(uint(v1178)%64)
						v1210 = int64(5)
						v1211 = int64(0)
						v1216 = int64(32)
						v1219 = int64(base.Ui64(v207) >> (uint(v1216) % 64))
						v1222 = int64(4294967295)
						v1225 = v207 & v1222
						v1226 = v1210 * v1225
						v1230 = int64(base.Ui64(v1226)>>(uint(v1216)%64)) + v1210*v1219
						v1237 = v1225*v1211 + v1230&v1222
						*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v207*v1211 + v216*v1210 + v1211*v1219 + int64(base.Ui64(v1230)>>(uint(v1216)%64)) + int64(base.Ui64(v1237)>>(uint(v1216)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1226&v1222 | v1237<<(uint(v1216)%64)
						v1249 = v1165 & int64(1)
						v1250 = v1249 + v1169
						v1254 = v1164 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1250) < base.Ui64(v1249)))
						if v1254 == v216 {
							v1257 = base.B2i32(base.Ui64(v207) < base.Ui64(v1250))
						} else {
							v1257 = base.B2i32(base.Ui64(v216) < base.Ui64(v1254))
						}
						v1259 = v1165 + base.I64_extend_i32_u(v1257)
						v1262 = v1168 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1259) < base.Ui64(v1165)))
						v1265 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
						v1267 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
						if v1254 == v1267 {
							v1270 = base.B2i32(base.Ui64(v1265) < base.Ui64(v1250))
						} else {
							v1270 = base.B2i32(base.Ui64(v1267) < base.Ui64(v1254))
						}
						v1273 = v1259 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1262) < base.Ui64(int64(9223090561878065152)))&v1270)
						v1276 = v1262 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1273) < base.Ui64(v1259)))
						v1279 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
						v1281 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
						if v1254 == v1281 {
							v1284 = base.B2i32(base.Ui64(v1279) < base.Ui64(v1250))
						} else {
							v1284 = base.B2i32(base.Ui64(v1281) < base.Ui64(v1254))
						}
						v1287 = v1273 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1276) < base.Ui64(int64(9223090561878065152)))&v1284)
						v1292 = v1287
						v1301 = v1276 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1287) < base.Ui64(v1273))) | v36
					}
				}
			}
		} else {
			v58 = l2 & int64(9223372036854775807)
			v59 = int64(9223090561878065152)
			if v58 == v59 {
				v63 = base.B2i32(l1 == int64(0))
			} else {
				v63 = base.B2i32(base.Ui64(v58) < base.Ui64(v59))
			}
			if v63 == int32(0) {
				v1292 = l1
				v1301 = l2 | int64(140737488355328)
			} else {
				v71 = l4 & int64(9223372036854775807)
				v72 = int64(9223090561878065152)
				if v71 == v72 {
					v76 = base.B2i32(l3 == int64(0))
				} else {
					v76 = base.B2i32(base.Ui64(v71) < base.Ui64(v72))
				}
				if v76 == int32(0) {
					v1292 = l3
					v1301 = l4 | int64(140737488355328)
				} else {
					if l1|(v58^int64(9223090561878065152)) == int64(0) {
						if l3|(v71^int64(9223090561878065152)) == int64(0) {
							v1292 = int64(0)
							v1301 = int64(9223231299366420480)
						} else {
							v1292 = int64(0)
							v1301 = v36 | int64(9223090561878065152)
						}
					} else {
						if l3|(v71^int64(9223090561878065152)) == int64(0) {
							v1292 = int64(0)
							v1301 = v36
						} else {
							if l1|v58 == int64(0) {
								if v71|l3 == int64(0) {
									v109 = int64(9223231299366420480)
								} else {
									v109 = v36
								}
								v1292 = int64(0)
								v1301 = v109
							} else {
								if v71|l3 == int64(0) {
									v1292 = int64(0)
									v1301 = v36 | int64(9223090561878065152)
								} else {
									if base.Ui64(v58) <= base.Ui64(int64(281474976710655)) {
										v120 = v28 + int32(320)
										v122 = base.B2i32(v33 == int64(0))
										if v33 == int64(0) {
											v123 = l1
										} else {
											v123 = v33
										}
										v129 = base.I32_wrap_i64(base.I64_clz(v123) + base.I64_extend_i32_u(v122<<(uint(int32(6))%32)))
										v131 = v129 - int32(15)
										if v131&int32(64) != 0 {
											v150 = int64(0)
											v151 = l1 << (uint(base.I64_extend_i32_u(v131+int32(-64))) % 64)
										} else {
											if v131 == int32(0) {
												v150 = l1
												v151 = v33
											} else {
												v142 = base.I64_extend_i32_u(v131)
												v150 = l1 << (uint(v142) % 64)
												v151 = v33<<(uint(v142)%64) | int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v131))%64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v120))) = v150
										*(*int64)(unsafe.Add(mBase, uint32(v120)+8)) = v151
										v157 = *(*int64)(unsafe.Add(mBase, uint32(v28)+328))
										v158 = *(*int64)(unsafe.Add(mBase, uint32(v28)+320))
										v159 = v158
										v160 = int32(16) - v129
										v161 = v157
									} else {
										v159 = l1
										v160 = v6
										v161 = v33
									}
									if base.Ui64(int64(281474976710655)) < base.Ui64(v71) {
										v205 = v159
										v207 = l3
										v208 = v160
										v210 = v161
										v211 = v31
									} else {
										v165 = v28 + int32(304)
										v167 = base.B2i32(v31 == int64(0))
										if v31 == int64(0) {
											v168 = l3
										} else {
											v168 = v31
										}
										v174 = base.I32_wrap_i64(base.I64_clz(v168) + base.I64_extend_i32_u(v167<<(uint(int32(6))%32)))
										v176 = v174 - int32(15)
										if v176&int32(64) != 0 {
											v195 = int64(0)
											v196 = l3 << (uint(base.I64_extend_i32_u(v176+int32(-64))) % 64)
										} else {
											if v176 == int32(0) {
												v195 = l3
												v196 = v31
											} else {
												v187 = base.I64_extend_i32_u(v176)
												v195 = l3 << (uint(v187) % 64)
												v196 = v31<<(uint(v187)%64) | int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v176))%64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v165))) = v195
										*(*int64)(unsafe.Add(mBase, uint32(v165)+8)) = v196
										v203 = *(*int64)(unsafe.Add(mBase, uint32(v28)+312))
										v204 = *(*int64)(unsafe.Add(mBase, uint32(v28)+304))
										v205 = v159
										v207 = v204
										v208 = v160 + v174 - int32(16)
										v210 = v161
										v211 = v203
									}
									v214 = v28 + int32(288)
									v216 = v211 | int64(281474976710656)
									v221 = v216<<(uint(int64(15))%64) | int64(base.Ui64(v207)>>(uint(int64(49))%64))
									v222 = int64(0)
									v224 = int64(8432131802713292800) - v221
									v230 = int64(32)
									v231 = int64(base.Ui64(v224) >> (uint(v230) % 64))
									v233 = int64(base.Ui64(v221) >> (uint(v230) % 64))
									v236 = int64(4294967295)
									v237 = v224 & v236
									v239 = v221 & v236
									v240 = v237 * v239
									v244 = int64(base.Ui64(v240)>>(uint(v230)%64)) + v237*v233
									v251 = v239*v231 + v244&v236
									*(*int64)(unsafe.Add(mBase, uint32(v214)+8)) = v221*v222 + v222*v224 + v231*v233 + int64(base.Ui64(v244)>>(uint(v230)%64)) + int64(base.Ui64(v251)>>(uint(v230)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v214))) = v240&v236 | v251<<(uint(v230)%64)
									v263 = v28 + int32(272)
									v264 = int64(0)
									v265 = *(*int64)(unsafe.Add(mBase, uint32(v28)+296))
									v266 = v264 - v265
									v273 = int64(32)
									v274 = int64(base.Ui64(v224) >> (uint(v273) % 64))
									v276 = int64(base.Ui64(v266) >> (uint(v273) % 64))
									v279 = int64(4294967295)
									v280 = v224 & v279
									v282 = v266 & v279
									v283 = v280 * v282
									v287 = int64(base.Ui64(v283)>>(uint(v273)%64)) + v280*v276
									v294 = v282*v274 + v287&v279
									*(*int64)(unsafe.Add(mBase, uint32(v263)+8)) = v266*v264 + v264*v224 + v274*v276 + int64(base.Ui64(v287)>>(uint(v273)%64)) + int64(base.Ui64(v294)>>(uint(v273)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v263))) = v283&v279 | v294<<(uint(v273)%64)
									v306 = v28 + int32(256)
									v307 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
									v310 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
									v313 = v307<<(uint(int64(1))%64) | int64(base.Ui64(v310)>>(uint(int64(63))%64))
									v314 = int64(0)
									v320 = int64(32)
									v321 = int64(base.Ui64(v221) >> (uint(v320) % 64))
									v323 = int64(base.Ui64(v313) >> (uint(v320) % 64))
									v326 = int64(4294967295)
									v327 = v221 & v326
									v329 = v313 & v326
									v330 = v327 * v329
									v334 = int64(base.Ui64(v330)>>(uint(v320)%64)) + v327*v323
									v341 = v329*v321 + v334&v326
									*(*int64)(unsafe.Add(mBase, uint32(v306)+8)) = v313*v314 + v314*v221 + v321*v323 + int64(base.Ui64(v334)>>(uint(v320)%64)) + int64(base.Ui64(v341)>>(uint(v320)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v306))) = v330&v326 | v341<<(uint(v320)%64)
									v353 = v28 + int32(240)
									v354 = int64(0)
									v356 = *(*int64)(unsafe.Add(mBase, uint32(v28)+264))
									v357 = v354 - v356
									v363 = int64(32)
									v364 = int64(base.Ui64(v357) >> (uint(v363) % 64))
									v366 = int64(base.Ui64(v313) >> (uint(v363) % 64))
									v369 = int64(4294967295)
									v370 = v357 & v369
									v372 = v313 & v369
									v373 = v370 * v372
									v377 = int64(base.Ui64(v373)>>(uint(v363)%64)) + v370*v366
									v384 = v372*v364 + v377&v369
									*(*int64)(unsafe.Add(mBase, uint32(v353)+8)) = v313*v354 + v354*v357 + v364*v366 + int64(base.Ui64(v377)>>(uint(v363)%64)) + int64(base.Ui64(v384)>>(uint(v363)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v353))) = v373&v369 | v384<<(uint(v363)%64)
									v396 = v28 + int32(224)
									v397 = *(*int64)(unsafe.Add(mBase, uint32(v28)+248))
									v400 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
									v403 = v397<<(uint(int64(1))%64) | int64(base.Ui64(v400)>>(uint(int64(63))%64))
									v404 = int64(0)
									v410 = int64(32)
									v411 = int64(base.Ui64(v221) >> (uint(v410) % 64))
									v413 = int64(base.Ui64(v403) >> (uint(v410) % 64))
									v416 = int64(4294967295)
									v417 = v221 & v416
									v419 = v403 & v416
									v420 = v417 * v419
									v424 = int64(base.Ui64(v420)>>(uint(v410)%64)) + v417*v413
									v431 = v419*v411 + v424&v416
									*(*int64)(unsafe.Add(mBase, uint32(v396)+8)) = v403*v404 + v404*v221 + v411*v413 + int64(base.Ui64(v424)>>(uint(v410)%64)) + int64(base.Ui64(v431)>>(uint(v410)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v396))) = v420&v416 | v431<<(uint(v410)%64)
									v443 = v28 + int32(208)
									v444 = int64(0)
									v446 = *(*int64)(unsafe.Add(mBase, uint32(v28)+232))
									v447 = v444 - v446
									v453 = int64(32)
									v454 = int64(base.Ui64(v447) >> (uint(v453) % 64))
									v456 = int64(base.Ui64(v403) >> (uint(v453) % 64))
									v459 = int64(4294967295)
									v460 = v447 & v459
									v462 = v403 & v459
									v463 = v460 * v462
									v467 = int64(base.Ui64(v463)>>(uint(v453)%64)) + v460*v456
									v474 = v462*v454 + v467&v459
									*(*int64)(unsafe.Add(mBase, uint32(v443)+8)) = v403*v444 + v444*v447 + v454*v456 + int64(base.Ui64(v467)>>(uint(v453)%64)) + int64(base.Ui64(v474)>>(uint(v453)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v443))) = v463&v459 | v474<<(uint(v453)%64)
									v486 = v28 + int32(192)
									v487 = *(*int64)(unsafe.Add(mBase, uint32(v28)+216))
									v490 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
									v493 = v487<<(uint(int64(1))%64) | int64(base.Ui64(v490)>>(uint(int64(63))%64))
									v494 = int64(0)
									v500 = int64(32)
									v501 = int64(base.Ui64(v221) >> (uint(v500) % 64))
									v503 = int64(base.Ui64(v493) >> (uint(v500) % 64))
									v506 = int64(4294967295)
									v507 = v221 & v506
									v509 = v493 & v506
									v510 = v507 * v509
									v514 = int64(base.Ui64(v510)>>(uint(v500)%64)) + v507*v503
									v521 = v509*v501 + v514&v506
									*(*int64)(unsafe.Add(mBase, uint32(v486)+8)) = v493*v494 + v494*v221 + v501*v503 + int64(base.Ui64(v514)>>(uint(v500)%64)) + int64(base.Ui64(v521)>>(uint(v500)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v486))) = v510&v506 | v521<<(uint(v500)%64)
									v533 = v28 + int32(176)
									v534 = int64(0)
									v536 = *(*int64)(unsafe.Add(mBase, uint32(v28)+200))
									v537 = v534 - v536
									v543 = int64(32)
									v544 = int64(base.Ui64(v537) >> (uint(v543) % 64))
									v546 = int64(base.Ui64(v493) >> (uint(v543) % 64))
									v549 = int64(4294967295)
									v550 = v537 & v549
									v552 = v493 & v549
									v553 = v550 * v552
									v557 = int64(base.Ui64(v553)>>(uint(v543)%64)) + v550*v546
									v564 = v552*v544 + v557&v549
									*(*int64)(unsafe.Add(mBase, uint32(v533)+8)) = v493*v534 + v534*v537 + v544*v546 + int64(base.Ui64(v557)>>(uint(v543)%64)) + int64(base.Ui64(v564)>>(uint(v543)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v533))) = v553&v549 | v564<<(uint(v543)%64)
									v576 = v28 + int32(160)
									v577 = int64(0)
									v578 = *(*int64)(unsafe.Add(mBase, uint32(v28)+184))
									v579 = int64(1)
									v581 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
									v586 = v578<<(uint(v579)%64) | int64(base.Ui64(v581)>>(uint(int64(63))%64)) - v579
									v592 = int64(32)
									v593 = int64(base.Ui64(v586) >> (uint(v592) % 64))
									v595 = int64(base.Ui64(v221) >> (uint(v592) % 64))
									v598 = int64(4294967295)
									v599 = v586 & v598
									v601 = v221 & v598
									v602 = v599 * v601
									v606 = int64(base.Ui64(v602)>>(uint(v592)%64)) + v599*v595
									v613 = v601*v593 + v606&v598
									*(*int64)(unsafe.Add(mBase, uint32(v576)+8)) = v221*v577 + v577*v586 + v593*v595 + int64(base.Ui64(v606)>>(uint(v592)%64)) + int64(base.Ui64(v613)>>(uint(v592)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v576))) = v602&v598 | v613<<(uint(v592)%64)
									v625 = v28 + int32(144)
									v627 = v207 << (uint(int64(15)) % 64)
									v628 = int64(0)
									v634 = int64(32)
									v635 = int64(base.Ui64(v586) >> (uint(v634) % 64))
									v637 = int64(base.Ui64(v627) >> (uint(v634) % 64))
									v640 = int64(4294967295)
									v641 = v586 & v640
									v643 = v627 & v640
									v644 = v641 * v643
									v648 = int64(base.Ui64(v644)>>(uint(v634)%64)) + v641*v637
									v655 = v643*v635 + v648&v640
									*(*int64)(unsafe.Add(mBase, uint32(v625)+8)) = v627*v628 + v628*v586 + v635*v637 + int64(base.Ui64(v648)>>(uint(v634)%64)) + int64(base.Ui64(v655)>>(uint(v634)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v625))) = v644&v640 | v655<<(uint(v634)%64)
									v667 = v28 + int32(112)
									v668 = int64(0)
									v670 = *(*int64)(unsafe.Add(mBase, uint32(v28)+168))
									v671 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
									v672 = *(*int64)(unsafe.Add(mBase, uint32(v28)+152))
									v673 = v671 + v672
									v681 = v668 - (v670 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v673) < base.Ui64(v671))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(1)) < base.Ui64(v673))))
									v687 = int64(32)
									v688 = int64(base.Ui64(v681) >> (uint(v687) % 64))
									v690 = int64(base.Ui64(v586) >> (uint(v687) % 64))
									v693 = int64(4294967295)
									v694 = v681 & v693
									v696 = v586 & v693
									v697 = v694 * v696
									v701 = int64(base.Ui64(v697)>>(uint(v687)%64)) + v694*v690
									v708 = v696*v688 + v701&v693
									*(*int64)(unsafe.Add(mBase, uint32(v667)+8)) = v586*v668 + v668*v681 + v688*v690 + int64(base.Ui64(v701)>>(uint(v687)%64)) + int64(base.Ui64(v708)>>(uint(v687)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v667))) = v697&v693 | v708<<(uint(v687)%64)
									v720 = v28 + int32(128)
									v722 = int64(1) - v673
									v723 = int64(0)
									v729 = int64(32)
									v730 = int64(base.Ui64(v586) >> (uint(v729) % 64))
									v732 = int64(base.Ui64(v722) >> (uint(v729) % 64))
									v735 = int64(4294967295)
									v736 = v586 & v735
									v738 = v722 & v735
									v739 = v736 * v738
									v743 = int64(base.Ui64(v739)>>(uint(v729)%64)) + v736*v732
									v750 = v738*v730 + v743&v735
									*(*int64)(unsafe.Add(mBase, uint32(v720)+8)) = v722*v723 + v723*v586 + v730*v732 + int64(base.Ui64(v743)>>(uint(v729)%64)) + int64(base.Ui64(v750)>>(uint(v729)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v720))) = v739&v735 | v750<<(uint(v729)%64)
									v762 = v208 + (v46 - v41)
									v763 = *(*int64)(unsafe.Add(mBase, uint32(v28)+112))
									v764 = int64(1)
									v765 = v763 << (uint(v764) % 64)
									v766 = *(*int64)(unsafe.Add(mBase, uint32(v28)+136))
									v769 = *(*int64)(unsafe.Add(mBase, uint32(v28)+128))
									v770 = int64(63)
									v773 = v765 + (v766<<(uint(v764)%64) | int64(base.Ui64(v769)>>(uint(v770)%64)))
									v775 = v773 - int64(13927)
									v776 = int64(32)
									v777 = int64(base.Ui64(v775) >> (uint(v776) % 64))
									v779 = v210 | int64(281474976710656)
									v781 = v779 << (uint(v764) % 64)
									v783 = int64(base.Ui64(v781) >> (uint(v776) % 64))
									v784 = v777 * v783
									v786 = v205 << (uint(v764) % 64)
									v788 = int64(base.Ui64(v786) >> (uint(v776) % 64))
									v793 = *(*int64)(unsafe.Add(mBase, uint32(v28)+120))
									v805 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v775) < base.Ui64(v773))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v773) < base.Ui64(v765))) + (v793<<(uint(v764)%64) | int64(base.Ui64(v763)>>(uint(v770)%64)) + int64(base.Ui64(v766)>>(uint(v770)%64)))) - v764
									v807 = int64(base.Ui64(v805) >> (uint(v776) % 64))
									v809 = v784 + v788*v807
									v812 = int64(4294967295)
									v813 = v805 & v812
									v815 = int64(base.Ui64(v205) >> (uint(v770) % 64))
									v820 = (v815 | v210<<(uint(v764)%64)) & v812
									v822 = v809 + v813*v820
									v828 = v783 * v813
									v830 = v828 + v820*v807
									v841 = v822 + v830<<(uint(v776)%64)
									v846 = v775 & v812
									v847 = v846 * v820
									v849 = v847 + v777*v788
									v853 = v786 & int64(4294967294)
									v855 = v849 + v813*v853
									v859 = v841 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v849) < base.Ui64(v847))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v855) < base.Ui64(v849))))
									v863 = v783 * v846
									v865 = v863 + v807*v853
									v867 = v865 + v777*v820
									v869 = v867 + v788*v813
									v883 = v859 + (int64(base.Ui64(v869)>>(uint(v776)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v869) < base.Ui64(v867)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v865) < base.Ui64(v863)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v867) < base.Ui64(v865)))))<<(uint(v776)%64))
									v887 = v777 * v853
									v889 = v887 + v788*v846
									v897 = v855 + (int64(base.Ui64(v889)>>(uint(v776)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v889) < base.Ui64(v887)))<<(uint(v776)%64))
									v906 = v883 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v897) < base.Ui64(v855))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v897+v869<<(uint(v776)%64)) < base.Ui64(v897))))
									v909 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v809) < base.Ui64(v784))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v822) < base.Ui64(v809))) + v783*v807 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v830) < base.Ui64(v828)))<<(uint(v776)%64) | int64(base.Ui64(v830)>>(uint(v776)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v841) < base.Ui64(v822))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v859) < base.Ui64(v841))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v883) < base.Ui64(v859))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v906) < base.Ui64(v883)))
									if base.Ui64(v909) <= base.Ui64(int64(562949953421311)) {
										v914 = v28 + int32(80)
										v919 = int64(32)
										v920 = int64(base.Ui64(v207) >> (uint(v919) % 64))
										v922 = int64(base.Ui64(v906) >> (uint(v919) % 64))
										v925 = int64(4294967295)
										v926 = v207 & v925
										v928 = v906 & v925
										v929 = v926 * v928
										v933 = int64(base.Ui64(v929)>>(uint(v919)%64)) + v926*v922
										v940 = v928*v920 + v933&v925
										*(*int64)(unsafe.Add(mBase, uint32(v914)+8)) = v906*v216 + v909*v207 + v920*v922 + int64(base.Ui64(v933)>>(uint(v919)%64)) + int64(base.Ui64(v940)>>(uint(v919)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v914))) = v929&v925 | v940<<(uint(v919)%64)
										v953 = *(*int64)(unsafe.Add(mBase, uint32(v28)+88))
										v955 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
										v956 = int64(0)
										v1023 = v906
										v1024 = v909
										v1025 = v762 + int32(_a_F___divtf3_1)
										v1027 = v205<<(uint(int64(49))%64) - v953 - base.I64_extend_i32_u(base.B2i32(v955 != v956))
										v1028 = v786
										v1029 = v781 | v815
										v1030 = v956 - v955
									} else {
										v965 = v28 + int32(96)
										v968 = int64(1)
										v970 = v909<<(uint(int64(63))%64) | int64(base.Ui64(v906)>>(uint(v968)%64))
										v972 = int64(base.Ui64(v909) >> (uint(v968) % 64))
										v977 = int64(32)
										v978 = int64(base.Ui64(v207) >> (uint(v977) % 64))
										v980 = int64(base.Ui64(v970) >> (uint(v977) % 64))
										v983 = int64(4294967295)
										v984 = v207 & v983
										v986 = v970 & v983
										v987 = v984 * v986
										v991 = int64(base.Ui64(v987)>>(uint(v977)%64)) + v984*v980
										v998 = v986*v978 + v991&v983
										*(*int64)(unsafe.Add(mBase, uint32(v965)+8)) = v970*v216 + v972*v207 + v978*v980 + int64(base.Ui64(v991)>>(uint(v977)%64)) + int64(base.Ui64(v998)>>(uint(v977)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v965))) = v987&v983 | v998<<(uint(v977)%64)
										v1011 = *(*int64)(unsafe.Add(mBase, uint32(v28)+104))
										v1013 = *(*int64)(unsafe.Add(mBase, uint32(v28)+96))
										v1014 = int64(0)
										v1023 = v970
										v1024 = v972
										v1025 = v762 + int32(_a_F___divtf3_2)
										v1027 = v205<<(uint(int64(48))%64) - v1011 - base.I64_extend_i32_u(base.B2i32(v1013 != v1014))
										v1028 = v205
										v1029 = v779
										v1030 = v1014 - v1013
									}
									if int32(_a_F___divtf3_0) <= v1025 {
										v1292 = int64(0)
										v1301 = v36 | int64(9223090561878065152)
									} else {
										if int32(0) < v1025 {
											v1038 = int64(1)
											v1164 = v1027<<(uint(v1038)%64) | int64(base.Ui64(v1030)>>(uint(int64(63))%64))
											v1165 = v1023
											v1168 = v1024&int64(281474976710655) | base.I64_extend_i32_u(v1025)<<(uint(int64(48))%64)
											v1169 = v1030 << (uint(v1038) % 64)
											v1171 = v28 + int32(16)
											v1172 = int64(3)
											v1173 = int64(0)
											v1178 = int64(32)
											v1181 = int64(base.Ui64(v207) >> (uint(v1178) % 64))
											v1184 = int64(4294967295)
											v1187 = v207 & v1184
											v1188 = v1172 * v1187
											v1192 = int64(base.Ui64(v1188)>>(uint(v1178)%64)) + v1172*v1181
											v1199 = v1187*v1173 + v1192&v1184
											*(*int64)(unsafe.Add(mBase, uint32(v1171)+8)) = v207*v1173 + v216*v1172 + v1173*v1181 + int64(base.Ui64(v1192)>>(uint(v1178)%64)) + int64(base.Ui64(v1199)>>(uint(v1178)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v1171))) = v1188&v1184 | v1199<<(uint(v1178)%64)
											v1210 = int64(5)
											v1211 = int64(0)
											v1216 = int64(32)
											v1219 = int64(base.Ui64(v207) >> (uint(v1216) % 64))
											v1222 = int64(4294967295)
											v1225 = v207 & v1222
											v1226 = v1210 * v1225
											v1230 = int64(base.Ui64(v1226)>>(uint(v1216)%64)) + v1210*v1219
											v1237 = v1225*v1211 + v1230&v1222
											*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v207*v1211 + v216*v1210 + v1211*v1219 + int64(base.Ui64(v1230)>>(uint(v1216)%64)) + int64(base.Ui64(v1237)>>(uint(v1216)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1226&v1222 | v1237<<(uint(v1216)%64)
											v1249 = v1165 & int64(1)
											v1250 = v1249 + v1169
											v1254 = v1164 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1250) < base.Ui64(v1249)))
											if v1254 == v216 {
												v1257 = base.B2i32(base.Ui64(v207) < base.Ui64(v1250))
											} else {
												v1257 = base.B2i32(base.Ui64(v216) < base.Ui64(v1254))
											}
											v1259 = v1165 + base.I64_extend_i32_u(v1257)
											v1262 = v1168 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1259) < base.Ui64(v1165)))
											v1265 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
											v1267 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
											if v1254 == v1267 {
												v1270 = base.B2i32(base.Ui64(v1265) < base.Ui64(v1250))
											} else {
												v1270 = base.B2i32(base.Ui64(v1267) < base.Ui64(v1254))
											}
											v1273 = v1259 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1262) < base.Ui64(int64(9223090561878065152)))&v1270)
											v1276 = v1262 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1273) < base.Ui64(v1259)))
											v1279 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
											v1281 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
											if v1254 == v1281 {
												v1284 = base.B2i32(base.Ui64(v1279) < base.Ui64(v1250))
											} else {
												v1284 = base.B2i32(base.Ui64(v1281) < base.Ui64(v1254))
											}
											v1287 = v1273 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1276) < base.Ui64(int64(9223090561878065152)))&v1284)
											v1292 = v1287
											v1301 = v1276 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1287) < base.Ui64(v1273))) | v36
										} else {
											if v1025 <= int32(-113) {
												v1292 = int64(0)
												v1301 = v36
											} else {
												v1055 = v28 - int32(-64)
												v1057 = int32(1) - v1025
												if v1057&int32(64) != 0 {
													v1076 = int64(base.Ui64(v1024) >> (uint(base.I64_extend_i32_u(v1057+int32(-64))) % 64))
													v1077 = int64(0)
												} else {
													if v1057 == int32(0) {
														v1076 = v1023
														v1077 = v1024
													} else {
														v1072 = base.I64_extend_i32_u(v1057)
														v1076 = v1024<<(uint(base.I64_extend_i32_u(int32(64)-v1057))%64) | int64(base.Ui64(v1023)>>(uint(v1072)%64))
														v1077 = int64(base.Ui64(v1024) >> (uint(v1072) % 64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v1055))) = v1076
												*(*int64)(unsafe.Add(mBase, uint32(v1055)+8)) = v1077
												v1082 = v28 + int32(48)
												v1084 = v1025 + int32(112)
												if v1084&int32(64) != 0 {
													v1103 = int64(0)
													v1104 = v1028 << (uint(base.I64_extend_i32_u(v1025+int32(48))) % 64)
												} else {
													if v1084 == int32(0) {
														v1103 = v1028
														v1104 = v1029
													} else {
														v1095 = base.I64_extend_i32_u(v1084)
														v1103 = v1028 << (uint(v1095) % 64)
														v1104 = v1029<<(uint(v1095)%64) | int64(base.Ui64(v1028)>>(uint(base.I64_extend_i32_u(int32(64)-v1084))%64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v1082))) = v1103
												*(*int64)(unsafe.Add(mBase, uint32(v1082)+8)) = v1104
												v1109 = v28 + int32(32)
												v1110 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
												v1111 = *(*int64)(unsafe.Add(mBase, uint32(v28)+72))
												v1116 = int64(32)
												v1117 = int64(base.Ui64(v1110) >> (uint(v1116) % 64))
												v1119 = int64(base.Ui64(v207) >> (uint(v1116) % 64))
												v1122 = int64(4294967295)
												v1123 = v1110 & v1122
												v1125 = v207 & v1122
												v1126 = v1123 * v1125
												v1130 = int64(base.Ui64(v1126)>>(uint(v1116)%64)) + v1123*v1119
												v1137 = v1125*v1117 + v1130&v1122
												*(*int64)(unsafe.Add(mBase, uint32(v1109)+8)) = v207*v1111 + v216*v1110 + v1117*v1119 + int64(base.Ui64(v1130)>>(uint(v1116)%64)) + int64(base.Ui64(v1137)>>(uint(v1116)%64))
												*(*int64)(unsafe.Add(mBase, uint32(v1109))) = v1126&v1122 | v1137<<(uint(v1116)%64)
												v1148 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
												v1149 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
												v1150 = int64(1)
												v1152 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
												v1157 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
												v1159 = v1152 << (uint(v1150) % 64)
												v1164 = v1148 - (v1149<<(uint(v1150)%64) | int64(base.Ui64(v1152)>>(uint(int64(63))%64))) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1157) < base.Ui64(v1159)))
												v1165 = v1110
												v1168 = v1111
												v1169 = v1157 - v1159
												v1171 = v28 + int32(16)
												v1172 = int64(3)
												v1173 = int64(0)
												v1178 = int64(32)
												v1181 = int64(base.Ui64(v207) >> (uint(v1178) % 64))
												v1184 = int64(4294967295)
												v1187 = v207 & v1184
												v1188 = v1172 * v1187
												v1192 = int64(base.Ui64(v1188)>>(uint(v1178)%64)) + v1172*v1181
												v1199 = v1187*v1173 + v1192&v1184
												*(*int64)(unsafe.Add(mBase, uint32(v1171)+8)) = v207*v1173 + v216*v1172 + v1173*v1181 + int64(base.Ui64(v1192)>>(uint(v1178)%64)) + int64(base.Ui64(v1199)>>(uint(v1178)%64))
												*(*int64)(unsafe.Add(mBase, uint32(v1171))) = v1188&v1184 | v1199<<(uint(v1178)%64)
												v1210 = int64(5)
												v1211 = int64(0)
												v1216 = int64(32)
												v1219 = int64(base.Ui64(v207) >> (uint(v1216) % 64))
												v1222 = int64(4294967295)
												v1225 = v207 & v1222
												v1226 = v1210 * v1225
												v1230 = int64(base.Ui64(v1226)>>(uint(v1216)%64)) + v1210*v1219
												v1237 = v1225*v1211 + v1230&v1222
												*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v207*v1211 + v216*v1210 + v1211*v1219 + int64(base.Ui64(v1230)>>(uint(v1216)%64)) + int64(base.Ui64(v1237)>>(uint(v1216)%64))
												*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1226&v1222 | v1237<<(uint(v1216)%64)
												v1249 = v1165 & int64(1)
												v1250 = v1249 + v1169
												v1254 = v1164 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1250) < base.Ui64(v1249)))
												if v1254 == v216 {
													v1257 = base.B2i32(base.Ui64(v207) < base.Ui64(v1250))
												} else {
													v1257 = base.B2i32(base.Ui64(v216) < base.Ui64(v1254))
												}
												v1259 = v1165 + base.I64_extend_i32_u(v1257)
												v1262 = v1168 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1259) < base.Ui64(v1165)))
												v1265 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
												v1267 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
												if v1254 == v1267 {
													v1270 = base.B2i32(base.Ui64(v1265) < base.Ui64(v1250))
												} else {
													v1270 = base.B2i32(base.Ui64(v1267) < base.Ui64(v1254))
												}
												v1273 = v1259 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1262) < base.Ui64(int64(9223090561878065152)))&v1270)
												v1276 = v1262 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1273) < base.Ui64(v1259)))
												v1279 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
												v1281 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
												if v1254 == v1281 {
													v1284 = base.B2i32(base.Ui64(v1279) < base.Ui64(v1250))
												} else {
													v1284 = base.B2i32(base.Ui64(v1281) < base.Ui64(v1254))
												}
												v1287 = v1273 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1276) < base.Ui64(int64(9223090561878065152)))&v1284)
												v1292 = v1287
												v1301 = v1276 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1287) < base.Ui64(v1273))) | v36
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
		v58 = l2 & int64(9223372036854775807)
		v59 = int64(9223090561878065152)
		if v58 == v59 {
			v63 = base.B2i32(l1 == int64(0))
		} else {
			v63 = base.B2i32(base.Ui64(v58) < base.Ui64(v59))
		}
		if v63 == int32(0) {
			v1292 = l1
			v1301 = l2 | int64(140737488355328)
		} else {
			v71 = l4 & int64(9223372036854775807)
			v72 = int64(9223090561878065152)
			if v71 == v72 {
				v76 = base.B2i32(l3 == int64(0))
			} else {
				v76 = base.B2i32(base.Ui64(v71) < base.Ui64(v72))
			}
			if v76 == int32(0) {
				v1292 = l3
				v1301 = l4 | int64(140737488355328)
			} else {
				if l1|(v58^int64(9223090561878065152)) == int64(0) {
					if l3|(v71^int64(9223090561878065152)) == int64(0) {
						v1292 = int64(0)
						v1301 = int64(9223231299366420480)
					} else {
						v1292 = int64(0)
						v1301 = v36 | int64(9223090561878065152)
					}
				} else {
					if l3|(v71^int64(9223090561878065152)) == int64(0) {
						v1292 = int64(0)
						v1301 = v36
					} else {
						if l1|v58 == int64(0) {
							if v71|l3 == int64(0) {
								v109 = int64(9223231299366420480)
							} else {
								v109 = v36
							}
							v1292 = int64(0)
							v1301 = v109
						} else {
							if v71|l3 == int64(0) {
								v1292 = int64(0)
								v1301 = v36 | int64(9223090561878065152)
							} else {
								if base.Ui64(v58) <= base.Ui64(int64(281474976710655)) {
									v120 = v28 + int32(320)
									v122 = base.B2i32(v33 == int64(0))
									if v33 == int64(0) {
										v123 = l1
									} else {
										v123 = v33
									}
									v129 = base.I32_wrap_i64(base.I64_clz(v123) + base.I64_extend_i32_u(v122<<(uint(int32(6))%32)))
									v131 = v129 - int32(15)
									if v131&int32(64) != 0 {
										v150 = int64(0)
										v151 = l1 << (uint(base.I64_extend_i32_u(v131+int32(-64))) % 64)
									} else {
										if v131 == int32(0) {
											v150 = l1
											v151 = v33
										} else {
											v142 = base.I64_extend_i32_u(v131)
											v150 = l1 << (uint(v142) % 64)
											v151 = v33<<(uint(v142)%64) | int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v131))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v120))) = v150
									*(*int64)(unsafe.Add(mBase, uint32(v120)+8)) = v151
									v157 = *(*int64)(unsafe.Add(mBase, uint32(v28)+328))
									v158 = *(*int64)(unsafe.Add(mBase, uint32(v28)+320))
									v159 = v158
									v160 = int32(16) - v129
									v161 = v157
								} else {
									v159 = l1
									v160 = v6
									v161 = v33
								}
								if base.Ui64(int64(281474976710655)) < base.Ui64(v71) {
									v205 = v159
									v207 = l3
									v208 = v160
									v210 = v161
									v211 = v31
								} else {
									v165 = v28 + int32(304)
									v167 = base.B2i32(v31 == int64(0))
									if v31 == int64(0) {
										v168 = l3
									} else {
										v168 = v31
									}
									v174 = base.I32_wrap_i64(base.I64_clz(v168) + base.I64_extend_i32_u(v167<<(uint(int32(6))%32)))
									v176 = v174 - int32(15)
									if v176&int32(64) != 0 {
										v195 = int64(0)
										v196 = l3 << (uint(base.I64_extend_i32_u(v176+int32(-64))) % 64)
									} else {
										if v176 == int32(0) {
											v195 = l3
											v196 = v31
										} else {
											v187 = base.I64_extend_i32_u(v176)
											v195 = l3 << (uint(v187) % 64)
											v196 = v31<<(uint(v187)%64) | int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v176))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v165))) = v195
									*(*int64)(unsafe.Add(mBase, uint32(v165)+8)) = v196
									v203 = *(*int64)(unsafe.Add(mBase, uint32(v28)+312))
									v204 = *(*int64)(unsafe.Add(mBase, uint32(v28)+304))
									v205 = v159
									v207 = v204
									v208 = v160 + v174 - int32(16)
									v210 = v161
									v211 = v203
								}
								v214 = v28 + int32(288)
								v216 = v211 | int64(281474976710656)
								v221 = v216<<(uint(int64(15))%64) | int64(base.Ui64(v207)>>(uint(int64(49))%64))
								v222 = int64(0)
								v224 = int64(8432131802713292800) - v221
								v230 = int64(32)
								v231 = int64(base.Ui64(v224) >> (uint(v230) % 64))
								v233 = int64(base.Ui64(v221) >> (uint(v230) % 64))
								v236 = int64(4294967295)
								v237 = v224 & v236
								v239 = v221 & v236
								v240 = v237 * v239
								v244 = int64(base.Ui64(v240)>>(uint(v230)%64)) + v237*v233
								v251 = v239*v231 + v244&v236
								*(*int64)(unsafe.Add(mBase, uint32(v214)+8)) = v221*v222 + v222*v224 + v231*v233 + int64(base.Ui64(v244)>>(uint(v230)%64)) + int64(base.Ui64(v251)>>(uint(v230)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v214))) = v240&v236 | v251<<(uint(v230)%64)
								v263 = v28 + int32(272)
								v264 = int64(0)
								v265 = *(*int64)(unsafe.Add(mBase, uint32(v28)+296))
								v266 = v264 - v265
								v273 = int64(32)
								v274 = int64(base.Ui64(v224) >> (uint(v273) % 64))
								v276 = int64(base.Ui64(v266) >> (uint(v273) % 64))
								v279 = int64(4294967295)
								v280 = v224 & v279
								v282 = v266 & v279
								v283 = v280 * v282
								v287 = int64(base.Ui64(v283)>>(uint(v273)%64)) + v280*v276
								v294 = v282*v274 + v287&v279
								*(*int64)(unsafe.Add(mBase, uint32(v263)+8)) = v266*v264 + v264*v224 + v274*v276 + int64(base.Ui64(v287)>>(uint(v273)%64)) + int64(base.Ui64(v294)>>(uint(v273)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v263))) = v283&v279 | v294<<(uint(v273)%64)
								v306 = v28 + int32(256)
								v307 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
								v310 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
								v313 = v307<<(uint(int64(1))%64) | int64(base.Ui64(v310)>>(uint(int64(63))%64))
								v314 = int64(0)
								v320 = int64(32)
								v321 = int64(base.Ui64(v221) >> (uint(v320) % 64))
								v323 = int64(base.Ui64(v313) >> (uint(v320) % 64))
								v326 = int64(4294967295)
								v327 = v221 & v326
								v329 = v313 & v326
								v330 = v327 * v329
								v334 = int64(base.Ui64(v330)>>(uint(v320)%64)) + v327*v323
								v341 = v329*v321 + v334&v326
								*(*int64)(unsafe.Add(mBase, uint32(v306)+8)) = v313*v314 + v314*v221 + v321*v323 + int64(base.Ui64(v334)>>(uint(v320)%64)) + int64(base.Ui64(v341)>>(uint(v320)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v306))) = v330&v326 | v341<<(uint(v320)%64)
								v353 = v28 + int32(240)
								v354 = int64(0)
								v356 = *(*int64)(unsafe.Add(mBase, uint32(v28)+264))
								v357 = v354 - v356
								v363 = int64(32)
								v364 = int64(base.Ui64(v357) >> (uint(v363) % 64))
								v366 = int64(base.Ui64(v313) >> (uint(v363) % 64))
								v369 = int64(4294967295)
								v370 = v357 & v369
								v372 = v313 & v369
								v373 = v370 * v372
								v377 = int64(base.Ui64(v373)>>(uint(v363)%64)) + v370*v366
								v384 = v372*v364 + v377&v369
								*(*int64)(unsafe.Add(mBase, uint32(v353)+8)) = v313*v354 + v354*v357 + v364*v366 + int64(base.Ui64(v377)>>(uint(v363)%64)) + int64(base.Ui64(v384)>>(uint(v363)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v353))) = v373&v369 | v384<<(uint(v363)%64)
								v396 = v28 + int32(224)
								v397 = *(*int64)(unsafe.Add(mBase, uint32(v28)+248))
								v400 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
								v403 = v397<<(uint(int64(1))%64) | int64(base.Ui64(v400)>>(uint(int64(63))%64))
								v404 = int64(0)
								v410 = int64(32)
								v411 = int64(base.Ui64(v221) >> (uint(v410) % 64))
								v413 = int64(base.Ui64(v403) >> (uint(v410) % 64))
								v416 = int64(4294967295)
								v417 = v221 & v416
								v419 = v403 & v416
								v420 = v417 * v419
								v424 = int64(base.Ui64(v420)>>(uint(v410)%64)) + v417*v413
								v431 = v419*v411 + v424&v416
								*(*int64)(unsafe.Add(mBase, uint32(v396)+8)) = v403*v404 + v404*v221 + v411*v413 + int64(base.Ui64(v424)>>(uint(v410)%64)) + int64(base.Ui64(v431)>>(uint(v410)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v396))) = v420&v416 | v431<<(uint(v410)%64)
								v443 = v28 + int32(208)
								v444 = int64(0)
								v446 = *(*int64)(unsafe.Add(mBase, uint32(v28)+232))
								v447 = v444 - v446
								v453 = int64(32)
								v454 = int64(base.Ui64(v447) >> (uint(v453) % 64))
								v456 = int64(base.Ui64(v403) >> (uint(v453) % 64))
								v459 = int64(4294967295)
								v460 = v447 & v459
								v462 = v403 & v459
								v463 = v460 * v462
								v467 = int64(base.Ui64(v463)>>(uint(v453)%64)) + v460*v456
								v474 = v462*v454 + v467&v459
								*(*int64)(unsafe.Add(mBase, uint32(v443)+8)) = v403*v444 + v444*v447 + v454*v456 + int64(base.Ui64(v467)>>(uint(v453)%64)) + int64(base.Ui64(v474)>>(uint(v453)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v443))) = v463&v459 | v474<<(uint(v453)%64)
								v486 = v28 + int32(192)
								v487 = *(*int64)(unsafe.Add(mBase, uint32(v28)+216))
								v490 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
								v493 = v487<<(uint(int64(1))%64) | int64(base.Ui64(v490)>>(uint(int64(63))%64))
								v494 = int64(0)
								v500 = int64(32)
								v501 = int64(base.Ui64(v221) >> (uint(v500) % 64))
								v503 = int64(base.Ui64(v493) >> (uint(v500) % 64))
								v506 = int64(4294967295)
								v507 = v221 & v506
								v509 = v493 & v506
								v510 = v507 * v509
								v514 = int64(base.Ui64(v510)>>(uint(v500)%64)) + v507*v503
								v521 = v509*v501 + v514&v506
								*(*int64)(unsafe.Add(mBase, uint32(v486)+8)) = v493*v494 + v494*v221 + v501*v503 + int64(base.Ui64(v514)>>(uint(v500)%64)) + int64(base.Ui64(v521)>>(uint(v500)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v486))) = v510&v506 | v521<<(uint(v500)%64)
								v533 = v28 + int32(176)
								v534 = int64(0)
								v536 = *(*int64)(unsafe.Add(mBase, uint32(v28)+200))
								v537 = v534 - v536
								v543 = int64(32)
								v544 = int64(base.Ui64(v537) >> (uint(v543) % 64))
								v546 = int64(base.Ui64(v493) >> (uint(v543) % 64))
								v549 = int64(4294967295)
								v550 = v537 & v549
								v552 = v493 & v549
								v553 = v550 * v552
								v557 = int64(base.Ui64(v553)>>(uint(v543)%64)) + v550*v546
								v564 = v552*v544 + v557&v549
								*(*int64)(unsafe.Add(mBase, uint32(v533)+8)) = v493*v534 + v534*v537 + v544*v546 + int64(base.Ui64(v557)>>(uint(v543)%64)) + int64(base.Ui64(v564)>>(uint(v543)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v533))) = v553&v549 | v564<<(uint(v543)%64)
								v576 = v28 + int32(160)
								v577 = int64(0)
								v578 = *(*int64)(unsafe.Add(mBase, uint32(v28)+184))
								v579 = int64(1)
								v581 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
								v586 = v578<<(uint(v579)%64) | int64(base.Ui64(v581)>>(uint(int64(63))%64)) - v579
								v592 = int64(32)
								v593 = int64(base.Ui64(v586) >> (uint(v592) % 64))
								v595 = int64(base.Ui64(v221) >> (uint(v592) % 64))
								v598 = int64(4294967295)
								v599 = v586 & v598
								v601 = v221 & v598
								v602 = v599 * v601
								v606 = int64(base.Ui64(v602)>>(uint(v592)%64)) + v599*v595
								v613 = v601*v593 + v606&v598
								*(*int64)(unsafe.Add(mBase, uint32(v576)+8)) = v221*v577 + v577*v586 + v593*v595 + int64(base.Ui64(v606)>>(uint(v592)%64)) + int64(base.Ui64(v613)>>(uint(v592)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v576))) = v602&v598 | v613<<(uint(v592)%64)
								v625 = v28 + int32(144)
								v627 = v207 << (uint(int64(15)) % 64)
								v628 = int64(0)
								v634 = int64(32)
								v635 = int64(base.Ui64(v586) >> (uint(v634) % 64))
								v637 = int64(base.Ui64(v627) >> (uint(v634) % 64))
								v640 = int64(4294967295)
								v641 = v586 & v640
								v643 = v627 & v640
								v644 = v641 * v643
								v648 = int64(base.Ui64(v644)>>(uint(v634)%64)) + v641*v637
								v655 = v643*v635 + v648&v640
								*(*int64)(unsafe.Add(mBase, uint32(v625)+8)) = v627*v628 + v628*v586 + v635*v637 + int64(base.Ui64(v648)>>(uint(v634)%64)) + int64(base.Ui64(v655)>>(uint(v634)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v625))) = v644&v640 | v655<<(uint(v634)%64)
								v667 = v28 + int32(112)
								v668 = int64(0)
								v670 = *(*int64)(unsafe.Add(mBase, uint32(v28)+168))
								v671 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
								v672 = *(*int64)(unsafe.Add(mBase, uint32(v28)+152))
								v673 = v671 + v672
								v681 = v668 - (v670 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v673) < base.Ui64(v671))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(1)) < base.Ui64(v673))))
								v687 = int64(32)
								v688 = int64(base.Ui64(v681) >> (uint(v687) % 64))
								v690 = int64(base.Ui64(v586) >> (uint(v687) % 64))
								v693 = int64(4294967295)
								v694 = v681 & v693
								v696 = v586 & v693
								v697 = v694 * v696
								v701 = int64(base.Ui64(v697)>>(uint(v687)%64)) + v694*v690
								v708 = v696*v688 + v701&v693
								*(*int64)(unsafe.Add(mBase, uint32(v667)+8)) = v586*v668 + v668*v681 + v688*v690 + int64(base.Ui64(v701)>>(uint(v687)%64)) + int64(base.Ui64(v708)>>(uint(v687)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v667))) = v697&v693 | v708<<(uint(v687)%64)
								v720 = v28 + int32(128)
								v722 = int64(1) - v673
								v723 = int64(0)
								v729 = int64(32)
								v730 = int64(base.Ui64(v586) >> (uint(v729) % 64))
								v732 = int64(base.Ui64(v722) >> (uint(v729) % 64))
								v735 = int64(4294967295)
								v736 = v586 & v735
								v738 = v722 & v735
								v739 = v736 * v738
								v743 = int64(base.Ui64(v739)>>(uint(v729)%64)) + v736*v732
								v750 = v738*v730 + v743&v735
								*(*int64)(unsafe.Add(mBase, uint32(v720)+8)) = v722*v723 + v723*v586 + v730*v732 + int64(base.Ui64(v743)>>(uint(v729)%64)) + int64(base.Ui64(v750)>>(uint(v729)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v720))) = v739&v735 | v750<<(uint(v729)%64)
								v762 = v208 + (v46 - v41)
								v763 = *(*int64)(unsafe.Add(mBase, uint32(v28)+112))
								v764 = int64(1)
								v765 = v763 << (uint(v764) % 64)
								v766 = *(*int64)(unsafe.Add(mBase, uint32(v28)+136))
								v769 = *(*int64)(unsafe.Add(mBase, uint32(v28)+128))
								v770 = int64(63)
								v773 = v765 + (v766<<(uint(v764)%64) | int64(base.Ui64(v769)>>(uint(v770)%64)))
								v775 = v773 - int64(13927)
								v776 = int64(32)
								v777 = int64(base.Ui64(v775) >> (uint(v776) % 64))
								v779 = v210 | int64(281474976710656)
								v781 = v779 << (uint(v764) % 64)
								v783 = int64(base.Ui64(v781) >> (uint(v776) % 64))
								v784 = v777 * v783
								v786 = v205 << (uint(v764) % 64)
								v788 = int64(base.Ui64(v786) >> (uint(v776) % 64))
								v793 = *(*int64)(unsafe.Add(mBase, uint32(v28)+120))
								v805 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v775) < base.Ui64(v773))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v773) < base.Ui64(v765))) + (v793<<(uint(v764)%64) | int64(base.Ui64(v763)>>(uint(v770)%64)) + int64(base.Ui64(v766)>>(uint(v770)%64)))) - v764
								v807 = int64(base.Ui64(v805) >> (uint(v776) % 64))
								v809 = v784 + v788*v807
								v812 = int64(4294967295)
								v813 = v805 & v812
								v815 = int64(base.Ui64(v205) >> (uint(v770) % 64))
								v820 = (v815 | v210<<(uint(v764)%64)) & v812
								v822 = v809 + v813*v820
								v828 = v783 * v813
								v830 = v828 + v820*v807
								v841 = v822 + v830<<(uint(v776)%64)
								v846 = v775 & v812
								v847 = v846 * v820
								v849 = v847 + v777*v788
								v853 = v786 & int64(4294967294)
								v855 = v849 + v813*v853
								v859 = v841 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v849) < base.Ui64(v847))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v855) < base.Ui64(v849))))
								v863 = v783 * v846
								v865 = v863 + v807*v853
								v867 = v865 + v777*v820
								v869 = v867 + v788*v813
								v883 = v859 + (int64(base.Ui64(v869)>>(uint(v776)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v869) < base.Ui64(v867)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v865) < base.Ui64(v863)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v867) < base.Ui64(v865)))))<<(uint(v776)%64))
								v887 = v777 * v853
								v889 = v887 + v788*v846
								v897 = v855 + (int64(base.Ui64(v889)>>(uint(v776)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v889) < base.Ui64(v887)))<<(uint(v776)%64))
								v906 = v883 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v897) < base.Ui64(v855))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v897+v869<<(uint(v776)%64)) < base.Ui64(v897))))
								v909 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v809) < base.Ui64(v784))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v822) < base.Ui64(v809))) + v783*v807 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v830) < base.Ui64(v828)))<<(uint(v776)%64) | int64(base.Ui64(v830)>>(uint(v776)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v841) < base.Ui64(v822))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v859) < base.Ui64(v841))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v883) < base.Ui64(v859))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v906) < base.Ui64(v883)))
								if base.Ui64(v909) <= base.Ui64(int64(562949953421311)) {
									v914 = v28 + int32(80)
									v919 = int64(32)
									v920 = int64(base.Ui64(v207) >> (uint(v919) % 64))
									v922 = int64(base.Ui64(v906) >> (uint(v919) % 64))
									v925 = int64(4294967295)
									v926 = v207 & v925
									v928 = v906 & v925
									v929 = v926 * v928
									v933 = int64(base.Ui64(v929)>>(uint(v919)%64)) + v926*v922
									v940 = v928*v920 + v933&v925
									*(*int64)(unsafe.Add(mBase, uint32(v914)+8)) = v906*v216 + v909*v207 + v920*v922 + int64(base.Ui64(v933)>>(uint(v919)%64)) + int64(base.Ui64(v940)>>(uint(v919)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v914))) = v929&v925 | v940<<(uint(v919)%64)
									v953 = *(*int64)(unsafe.Add(mBase, uint32(v28)+88))
									v955 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
									v956 = int64(0)
									v1023 = v906
									v1024 = v909
									v1025 = v762 + int32(_a_F___divtf3_1)
									v1027 = v205<<(uint(int64(49))%64) - v953 - base.I64_extend_i32_u(base.B2i32(v955 != v956))
									v1028 = v786
									v1029 = v781 | v815
									v1030 = v956 - v955
								} else {
									v965 = v28 + int32(96)
									v968 = int64(1)
									v970 = v909<<(uint(int64(63))%64) | int64(base.Ui64(v906)>>(uint(v968)%64))
									v972 = int64(base.Ui64(v909) >> (uint(v968) % 64))
									v977 = int64(32)
									v978 = int64(base.Ui64(v207) >> (uint(v977) % 64))
									v980 = int64(base.Ui64(v970) >> (uint(v977) % 64))
									v983 = int64(4294967295)
									v984 = v207 & v983
									v986 = v970 & v983
									v987 = v984 * v986
									v991 = int64(base.Ui64(v987)>>(uint(v977)%64)) + v984*v980
									v998 = v986*v978 + v991&v983
									*(*int64)(unsafe.Add(mBase, uint32(v965)+8)) = v970*v216 + v972*v207 + v978*v980 + int64(base.Ui64(v991)>>(uint(v977)%64)) + int64(base.Ui64(v998)>>(uint(v977)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v965))) = v987&v983 | v998<<(uint(v977)%64)
									v1011 = *(*int64)(unsafe.Add(mBase, uint32(v28)+104))
									v1013 = *(*int64)(unsafe.Add(mBase, uint32(v28)+96))
									v1014 = int64(0)
									v1023 = v970
									v1024 = v972
									v1025 = v762 + int32(_a_F___divtf3_2)
									v1027 = v205<<(uint(int64(48))%64) - v1011 - base.I64_extend_i32_u(base.B2i32(v1013 != v1014))
									v1028 = v205
									v1029 = v779
									v1030 = v1014 - v1013
								}
								if int32(_a_F___divtf3_0) <= v1025 {
									v1292 = int64(0)
									v1301 = v36 | int64(9223090561878065152)
								} else {
									if int32(0) < v1025 {
										v1038 = int64(1)
										v1164 = v1027<<(uint(v1038)%64) | int64(base.Ui64(v1030)>>(uint(int64(63))%64))
										v1165 = v1023
										v1168 = v1024&int64(281474976710655) | base.I64_extend_i32_u(v1025)<<(uint(int64(48))%64)
										v1169 = v1030 << (uint(v1038) % 64)
										v1171 = v28 + int32(16)
										v1172 = int64(3)
										v1173 = int64(0)
										v1178 = int64(32)
										v1181 = int64(base.Ui64(v207) >> (uint(v1178) % 64))
										v1184 = int64(4294967295)
										v1187 = v207 & v1184
										v1188 = v1172 * v1187
										v1192 = int64(base.Ui64(v1188)>>(uint(v1178)%64)) + v1172*v1181
										v1199 = v1187*v1173 + v1192&v1184
										*(*int64)(unsafe.Add(mBase, uint32(v1171)+8)) = v207*v1173 + v216*v1172 + v1173*v1181 + int64(base.Ui64(v1192)>>(uint(v1178)%64)) + int64(base.Ui64(v1199)>>(uint(v1178)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v1171))) = v1188&v1184 | v1199<<(uint(v1178)%64)
										v1210 = int64(5)
										v1211 = int64(0)
										v1216 = int64(32)
										v1219 = int64(base.Ui64(v207) >> (uint(v1216) % 64))
										v1222 = int64(4294967295)
										v1225 = v207 & v1222
										v1226 = v1210 * v1225
										v1230 = int64(base.Ui64(v1226)>>(uint(v1216)%64)) + v1210*v1219
										v1237 = v1225*v1211 + v1230&v1222
										*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v207*v1211 + v216*v1210 + v1211*v1219 + int64(base.Ui64(v1230)>>(uint(v1216)%64)) + int64(base.Ui64(v1237)>>(uint(v1216)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1226&v1222 | v1237<<(uint(v1216)%64)
										v1249 = v1165 & int64(1)
										v1250 = v1249 + v1169
										v1254 = v1164 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1250) < base.Ui64(v1249)))
										if v1254 == v216 {
											v1257 = base.B2i32(base.Ui64(v207) < base.Ui64(v1250))
										} else {
											v1257 = base.B2i32(base.Ui64(v216) < base.Ui64(v1254))
										}
										v1259 = v1165 + base.I64_extend_i32_u(v1257)
										v1262 = v1168 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1259) < base.Ui64(v1165)))
										v1265 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
										v1267 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
										if v1254 == v1267 {
											v1270 = base.B2i32(base.Ui64(v1265) < base.Ui64(v1250))
										} else {
											v1270 = base.B2i32(base.Ui64(v1267) < base.Ui64(v1254))
										}
										v1273 = v1259 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1262) < base.Ui64(int64(9223090561878065152)))&v1270)
										v1276 = v1262 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1273) < base.Ui64(v1259)))
										v1279 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
										v1281 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
										if v1254 == v1281 {
											v1284 = base.B2i32(base.Ui64(v1279) < base.Ui64(v1250))
										} else {
											v1284 = base.B2i32(base.Ui64(v1281) < base.Ui64(v1254))
										}
										v1287 = v1273 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1276) < base.Ui64(int64(9223090561878065152)))&v1284)
										v1292 = v1287
										v1301 = v1276 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1287) < base.Ui64(v1273))) | v36
									} else {
										if v1025 <= int32(-113) {
											v1292 = int64(0)
											v1301 = v36
										} else {
											v1055 = v28 - int32(-64)
											v1057 = int32(1) - v1025
											if v1057&int32(64) != 0 {
												v1076 = int64(base.Ui64(v1024) >> (uint(base.I64_extend_i32_u(v1057+int32(-64))) % 64))
												v1077 = int64(0)
											} else {
												if v1057 == int32(0) {
													v1076 = v1023
													v1077 = v1024
												} else {
													v1072 = base.I64_extend_i32_u(v1057)
													v1076 = v1024<<(uint(base.I64_extend_i32_u(int32(64)-v1057))%64) | int64(base.Ui64(v1023)>>(uint(v1072)%64))
													v1077 = int64(base.Ui64(v1024) >> (uint(v1072) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v1055))) = v1076
											*(*int64)(unsafe.Add(mBase, uint32(v1055)+8)) = v1077
											v1082 = v28 + int32(48)
											v1084 = v1025 + int32(112)
											if v1084&int32(64) != 0 {
												v1103 = int64(0)
												v1104 = v1028 << (uint(base.I64_extend_i32_u(v1025+int32(48))) % 64)
											} else {
												if v1084 == int32(0) {
													v1103 = v1028
													v1104 = v1029
												} else {
													v1095 = base.I64_extend_i32_u(v1084)
													v1103 = v1028 << (uint(v1095) % 64)
													v1104 = v1029<<(uint(v1095)%64) | int64(base.Ui64(v1028)>>(uint(base.I64_extend_i32_u(int32(64)-v1084))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v1082))) = v1103
											*(*int64)(unsafe.Add(mBase, uint32(v1082)+8)) = v1104
											v1109 = v28 + int32(32)
											v1110 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
											v1111 = *(*int64)(unsafe.Add(mBase, uint32(v28)+72))
											v1116 = int64(32)
											v1117 = int64(base.Ui64(v1110) >> (uint(v1116) % 64))
											v1119 = int64(base.Ui64(v207) >> (uint(v1116) % 64))
											v1122 = int64(4294967295)
											v1123 = v1110 & v1122
											v1125 = v207 & v1122
											v1126 = v1123 * v1125
											v1130 = int64(base.Ui64(v1126)>>(uint(v1116)%64)) + v1123*v1119
											v1137 = v1125*v1117 + v1130&v1122
											*(*int64)(unsafe.Add(mBase, uint32(v1109)+8)) = v207*v1111 + v216*v1110 + v1117*v1119 + int64(base.Ui64(v1130)>>(uint(v1116)%64)) + int64(base.Ui64(v1137)>>(uint(v1116)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v1109))) = v1126&v1122 | v1137<<(uint(v1116)%64)
											v1148 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
											v1149 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
											v1150 = int64(1)
											v1152 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
											v1157 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
											v1159 = v1152 << (uint(v1150) % 64)
											v1164 = v1148 - (v1149<<(uint(v1150)%64) | int64(base.Ui64(v1152)>>(uint(int64(63))%64))) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1157) < base.Ui64(v1159)))
											v1165 = v1110
											v1168 = v1111
											v1169 = v1157 - v1159
											v1171 = v28 + int32(16)
											v1172 = int64(3)
											v1173 = int64(0)
											v1178 = int64(32)
											v1181 = int64(base.Ui64(v207) >> (uint(v1178) % 64))
											v1184 = int64(4294967295)
											v1187 = v207 & v1184
											v1188 = v1172 * v1187
											v1192 = int64(base.Ui64(v1188)>>(uint(v1178)%64)) + v1172*v1181
											v1199 = v1187*v1173 + v1192&v1184
											*(*int64)(unsafe.Add(mBase, uint32(v1171)+8)) = v207*v1173 + v216*v1172 + v1173*v1181 + int64(base.Ui64(v1192)>>(uint(v1178)%64)) + int64(base.Ui64(v1199)>>(uint(v1178)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v1171))) = v1188&v1184 | v1199<<(uint(v1178)%64)
											v1210 = int64(5)
											v1211 = int64(0)
											v1216 = int64(32)
											v1219 = int64(base.Ui64(v207) >> (uint(v1216) % 64))
											v1222 = int64(4294967295)
											v1225 = v207 & v1222
											v1226 = v1210 * v1225
											v1230 = int64(base.Ui64(v1226)>>(uint(v1216)%64)) + v1210*v1219
											v1237 = v1225*v1211 + v1230&v1222
											*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v207*v1211 + v216*v1210 + v1211*v1219 + int64(base.Ui64(v1230)>>(uint(v1216)%64)) + int64(base.Ui64(v1237)>>(uint(v1216)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1226&v1222 | v1237<<(uint(v1216)%64)
											v1249 = v1165 & int64(1)
											v1250 = v1249 + v1169
											v1254 = v1164 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1250) < base.Ui64(v1249)))
											if v1254 == v216 {
												v1257 = base.B2i32(base.Ui64(v207) < base.Ui64(v1250))
											} else {
												v1257 = base.B2i32(base.Ui64(v216) < base.Ui64(v1254))
											}
											v1259 = v1165 + base.I64_extend_i32_u(v1257)
											v1262 = v1168 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1259) < base.Ui64(v1165)))
											v1265 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
											v1267 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
											if v1254 == v1267 {
												v1270 = base.B2i32(base.Ui64(v1265) < base.Ui64(v1250))
											} else {
												v1270 = base.B2i32(base.Ui64(v1267) < base.Ui64(v1254))
											}
											v1273 = v1259 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1262) < base.Ui64(int64(9223090561878065152)))&v1270)
											v1276 = v1262 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1273) < base.Ui64(v1259)))
											v1279 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
											v1281 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
											if v1254 == v1281 {
												v1284 = base.B2i32(base.Ui64(v1279) < base.Ui64(v1250))
											} else {
												v1284 = base.B2i32(base.Ui64(v1281) < base.Ui64(v1254))
											}
											v1287 = v1273 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1276) < base.Ui64(int64(9223090561878065152)))&v1284)
											v1292 = v1287
											v1301 = v1276 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1287) < base.Ui64(v1273))) | v36
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
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1292
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v1301
	m.G0 = v28 + int32(336)
	return
}
func F_dacosh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 float64
	_ = v34
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v47 int64
	_ = v47
	var v69 float64
	_ = v69
	var v70 int64
	_ = v70
	var v75 int32
	_ = v75
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v107 float64
	_ = v107
	var v109 float64
	_ = v109
	var v113 float64
	_ = v113
	var v118 float64
	_ = v118
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v161 float64
	_ = v161
	var v170 float64
	_ = v170
	var v174 float64
	_ = v174
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v187 float64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	if base.F64_lt(v4, float64(1)) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_dacosh_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_dacosh_1), int32(2703), int32(_a_F_dacosh_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v30 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v4))>>(uint(int64(52))%64))) & int32(2047)
		if base.Ui32(v30) <= base.Ui32(int32(1023)) {
			v34 = base.F64_add(v4, float64(-1))
			v39 = base.F64_add(v34, base.F64_sqrt(base.F64_add(base.F64_mul(v34, v34), base.F64_add(v34, v34))))
			v40 = float64(0)
			v47 = base.I64_reinterpret_f64(v39)
			if v47 <= int64(4601133429810003967) {
				if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v47) {
					if base.F64_eq(v39, float64(-1)) != 0 {
						v161 = math.Float64frombits(uint64(0xfff0000000000000))
						v170 = v161
					} else {
						v170 = base.F64_div(base.F64_sub(v39, v39), float64(0))
					}
				} else {
					if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v47)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
						v170 = v39
					} else {
						if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v47) {
							v69 = base.F64_add(v39, float64(1))
							v70 = base.I64_reinterpret_f64(v69)
							v75 = base.I32_wrap_i64(int64(base.Ui64(v70)>>(uint(int64(32))%64))) + int32(_a_F_dacosh_3)
							if base.Ui32(v75) <= base.Ui32(int32(1129316351)) {
								if base.Ui32(int32(1074790399)) < base.Ui32(v75) {
									v90 = base.F64_add(base.F64_sub(v39, v69), float64(1))
								} else {
									v90 = base.F64_sub(v39, base.F64_add(v69, float64(-1)))
								}
								v92 = base.F64_div(v90, v69)
							} else {
								v92 = v40
							}
							v107 = base.F64_add(base.F64_reinterpret_i64(v70&int64(4294967295)|base.I64_extend_i32_u(v75&int32(_a_F_dacosh_4)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v109 = v92
							v113 = base.F64_convert_i32_s(int32(base.Ui32(v75)>>(uint(int32(20))%32)) - int32(1023))
						} else {
							v107 = v39
							v109 = v40
							v113 = float64(0)
						}
						v118 = base.F64_div(v107, base.F64_add(v107, float64(2)))
						v121 = base.F64_mul(v107, base.F64_mul(v107, float64(0.5)))
						v122 = base.F64_mul(v118, v118)
						v123 = base.F64_mul(v122, v122)
						v161 = base.F64_add(base.F64_mul(v113, float64(0.6931471803691238)), base.F64_add(v107, base.F64_sub(base.F64_add(base.F64_mul(v118, base.F64_add(v121, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v122, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v113, float64(1.9082149292705877e-10)), v109)), v121)))
						v170 = v161
					}
				}
			} else {
				if base.Ui64(int64(9218868437227405311)) < base.Ui64(v47) {
					v170 = v39
				} else {
					v69 = base.F64_add(v39, float64(1))
					v70 = base.I64_reinterpret_f64(v69)
					v75 = base.I32_wrap_i64(int64(base.Ui64(v70)>>(uint(int64(32))%64))) + int32(_a_F_dacosh_3)
					if base.Ui32(v75) <= base.Ui32(int32(1129316351)) {
						if base.Ui32(int32(1074790399)) < base.Ui32(v75) {
							v90 = base.F64_add(base.F64_sub(v39, v69), float64(1))
						} else {
							v90 = base.F64_sub(v39, base.F64_add(v69, float64(-1)))
						}
						v92 = base.F64_div(v90, v69)
					} else {
						v92 = v40
					}
					v107 = base.F64_add(base.F64_reinterpret_i64(v70&int64(4294967295)|base.I64_extend_i32_u(v75&int32(_a_F_dacosh_4)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
					v109 = v92
					v113 = base.F64_convert_i32_s(int32(base.Ui32(v75)>>(uint(int32(20))%32)) - int32(1023))
					v118 = base.F64_div(v107, base.F64_add(v107, float64(2)))
					v121 = base.F64_mul(v107, base.F64_mul(v107, float64(0.5)))
					v122 = base.F64_mul(v118, v118)
					v123 = base.F64_mul(v122, v122)
					v161 = base.F64_add(base.F64_mul(v113, float64(0.6931471803691238)), base.F64_add(v107, base.F64_sub(base.F64_add(base.F64_mul(v118, base.F64_add(v121, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v122, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v113, float64(1.9082149292705877e-10)), v109)), v121)))
					v170 = v161
				}
			}
			v187 = v170
		} else {
			if base.Ui32(v30) <= base.Ui32(int32(1048)) {
				v174 = float64(-1)
				v182 = F_log(m, base.F64_add(base.F64_add(v4, v4), base.F64_div(v174, base.F64_add(v4, base.F64_sqrt(base.F64_add(base.F64_mul(v4, v4), v174))))))
				mBase = m.M
				v187 = v182
			} else {
				v183 = F_log(m, v4)
				mBase = m.M
				v187 = base.F64_add(v183, float64(0.6931471805599453))
			}
		}
		v188 = F_Float8GetDatum(m, v187)
		mBase = m.M
		v189 = m.ExcPending
		if v189 != 0 {
			return int32(0)
		} else {
			return v188
		}
	}
}
func F_danish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v693 int32
	_ = v693
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v318 < v321 {
		goto L78
	} else {
		goto L79
	}
L2:
	;
	if v62 < int32(0) {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	goto L5
L5:
	;
	goto L6
L6:
	;
	v17 = v10
	v19 = int32(3)
	goto L9
L8:
	;
	v62 = v47
	goto L2
L9:
	;
	if v7 <= v17 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v62 = int32(-1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v24 = v17 + int32(1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v17))))
	if base.Ui32(v26) < base.Ui32(int32(192)) {
		v47 = v24
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = int32(1)
	if v48 < v19 {
		v17 = v47
		v19 = v19 - v48
		goto L9
	} else {
		goto L21
	}
L15:
	;
	if v7 <= v24 {
		v47 = v24
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = v24
	goto L17
L17:
	;
	v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9+v33))))
	if int32(-65) < v36 {
		v47 = v33
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v47 = v7
	goto L14
L19:
	;
	v40 = v33 + int32(1)
	if v40 != v7 {
		v33 = v40
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L10
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v62
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = v10
	goto L25
L23:
	;
	if v184 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v184 = v156
	goto L23
L25:
	;
	if v80 <= v89 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v184 = int32(-1)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v96 = int32(1)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v81))))
	if base.Ui32(v98) < base.Ui32(int32(192)) {
		v155 = v98
		v156 = v96
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if int32(248) < v155 {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v102 = v89 + int32(1)
	if v102 == v80 {
		v155 = v98
		v156 = v96
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v81))))
	v107 = v105 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v98) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+v81))))
	v123 = v121 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v98) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v111 = v89 + int32(2)
	if v111 != v80 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v155 = v98<<(uint(int32(6))%32)&int32(1984) | v107
	v156 = int32(2)
	goto L30
L37:
	;
	goto L36
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v127))))
	v155 = v140&int32(63) | (v98<<(uint(int32(18))%32)&int32(_a_F_danish_UTF_8_stem_0) | v107<<(uint(int32(12))%32) | v123<<(uint(int32(6))%32))
	v156 = int32(4)
	goto L30
L39:
	;
	v127 = v89 + int32(3)
	if v127 != v80 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v155 = v98<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v107<<(uint(int32(6))%32) | v123
	v156 = int32(3)
	goto L30
L42:
	;
	goto L41
L43:
	;
	v173 = v156 + v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v173
	v89 = v173
	goto L25
L44:
	;
	v160 = v155 - int32(97)
	if v160 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v160)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[0]))))
	if int32(base.Ui32(v166)>>(uint(v160&int32(7))%32))&int32(1) != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = v198
	goto L51
L49:
	;
	if v304 < int32(0) {
		goto L1
	} else {
		goto L73
	}
L50:
	;
	v304 = v275
	goto L49
L51:
	;
	if v199 <= v208 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v304 = int32(-1)
	goto L49
L54:
	;
	goto L55
L55:
	;
	v215 = int32(1)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208+v200))))
	if base.Ui32(v217) < base.Ui32(int32(192)) {
		v274 = v217
		v275 = v215
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if int32(248) < v274 {
		goto L50
	} else {
		goto L69
	}
L57:
	;
	v221 = v208 + int32(1)
	if v221 == v199 {
		v274 = v217
		v275 = v215
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v200))))
	v226 = v224 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v217) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v200))))
	v242 = v240 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v217) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v230 = v208 + int32(2)
	if v230 != v199 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v274 = v217<<(uint(int32(6))%32)&int32(1984) | v226
	v275 = int32(2)
	goto L56
L63:
	;
	goto L62
L64:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v246))))
	v274 = v259&int32(63) | (v217<<(uint(int32(18))%32)&int32(_a_F_danish_UTF_8_stem_0) | v226<<(uint(int32(12))%32) | v242<<(uint(int32(6))%32))
	v275 = int32(4)
	goto L56
L65:
	;
	v246 = v208 + int32(3)
	if v246 != v199 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v274 = v217<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v226<<(uint(int32(6))%32) | v242
	v275 = int32(3)
	goto L56
L68:
	;
	goto L67
L69:
	;
	v279 = v274 - int32(97)
	if v279 < int32(0) {
		goto L50
	} else {
		goto L70
	}
L70:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v279)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[0]))))
	if int32(base.Ui32(v285)>>(uint(v279&int32(7))%32))&int32(1) == int32(0) {
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v293 = v275 + v208
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v293
	v208 = v293
	goto L51
L73:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v308 = v307 + v304
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	if v311 < v308 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v313 = v308
	goto L76
L75:
	;
	v313 = v311
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+4)) = v313
	goto L1
L77:
	;
	return v779
L78:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v494
	v496 = F_r_consonant_pair_2(m, l0)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L84
	} else {
		goto L118
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v321
	if v318 <= v321 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	goto L78
L81:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v318-int32(1)))))
	if v330&int32(224) != int32(96) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	if int32(1)<<(uint(v330)%32)&int32(_a_F_danish_UTF_8_stem_2) == int32(0) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v343 = F_find_among_b(m, l0, int32(_a_F_danish_UTF_8_stem_3), int32(32))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	return int32(0)
L85:
	;
	if v343 == int32(0) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v350
	switch v343 - int32(1) {
	case 0:
		goto L88
	case 1:
		goto L87
	default:
		goto L78
	}
L87:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L93
L88:
	;
	v354 = F_slice_del(m, l0)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L84
	} else {
		goto L89
	}
L89:
	;
	if int32(0) <= v354 {
		goto L78
	} else {
		goto L90
	}
L90:
	;
	v779 = v354
	goto L77
L91:
	;
	if v486 != 0 {
		goto L78
	} else {
		goto L115
	}
L92:
	;
	v486 = v479
	goto L91
L93:
	;
	if v374 <= v375 {
		v479 = int32(-1)
		goto L92
	} else {
		goto L95
	}
L94:
	;
	v479 = int32(0)
	goto L92
L95:
	;
	v392 = int32(1)
	v393 = v374 - v392
	v395 = int32(*(*int8)(unsafe.Add(mBase, uint32(v371+v393))))
	v397 = v395 & int32(255)
	if v393 == v375 {
		v452 = v397
		v453 = v392
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if int32(229) < v452 {
		goto L105
	} else {
		goto L106
	}
L97:
	;
	if int32(0) <= v395 {
		v452 = v397
		v453 = v392
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v403 = v397 & int32(63)
	v405 = v374 - int32(2)
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371+v405))))
	v409 = v407 << (uint(int32(6)) % 32)
	if base.B2i32(v405 != v375)&base.B2i32(base.Ui32(v407) < base.Ui32(int32(192))) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v452 = v409&int32(1984) | v403
	v453 = int32(2)
	goto L96
L100:
	;
	goto L101
L101:
	;
	v422 = v409&int32(4032) | v403
	v424 = v374 - int32(3)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371+v424))))
	if base.B2i32(v424 != v375)&base.B2i32(base.Ui32(v426) < base.Ui32(int32(224))) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v452 = v426<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v422
	v453 = int32(3)
	goto L96
L103:
	;
	goto L104
L104:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+(v371-int32(4))))))
	v452 = v426<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_4) | v444&int32(7)<<(uint(int32(18))%32) | v422
	v453 = int32(4)
	goto L96
L105:
	;
	v486 = v453
	goto L91
L106:
	;
	goto L107
L107:
	;
	v457 = v452 - int32(97)
	if v457 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v486 = v453
	goto L91
L109:
	;
	goto L110
L110:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[1]))))
	if int32(base.Ui32(v463)>>(uint(v457&int32(7))%32))&int32(1) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v486 = v453
	goto L91
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v374 - v453
	goto L114
L114:
	;
	goto L94
L115:
	;
	v487 = F_slice_del(m, l0)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L84
	} else {
		goto L116
	}
L116:
	;
	if int32(0) <= v487 {
		goto L78
	} else {
		goto L117
	}
L117:
	;
	v779 = v487
	goto L77
L118:
	;
	if v496 < int32(0) {
		v779 = v496
		goto L77
	} else {
		goto L119
	}
L119:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v500
	v503 = int32(2)
	v505 = int32(0)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v500-v508 < v503 {
		v518 = v505
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v546
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	if v546 < v549 {
		goto L133
	} else {
		goto L134
	}
L121:
	;
	if v518 == int32(0) {
		goto L120
	} else {
		goto L125
	}
L122:
	;
	goto L121
L123:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v514 = F_memcmp(m, v511+v500-v503, int32(_a_F_danish_UTF_8_stem_5), v503)
	mBase = m.M
	if v514 != 0 {
		v518 = v505
		goto L122
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v500 - v503
	v518 = int32(1)
	goto L122
L125:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v521
	v523 = int32(2)
	v525 = int32(0)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v521-v528 < v523 {
		v538 = v525
		goto L127
	} else {
		goto L128
	}
L126:
	;
	if v538 == int32(0) {
		goto L120
	} else {
		goto L130
	}
L127:
	;
	goto L126
L128:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v534 = F_memcmp(m, v531+v521-v523, int32(_a_F_danish_UTF_8_stem_6), v523)
	mBase = m.M
	if v534 != 0 {
		v538 = v525
		goto L127
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v521 - v523
	v538 = int32(1)
	goto L127
L130:
	;
	v541 = F_slice_del(m, l0)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L84
	} else {
		goto L131
	}
L131:
	;
	if v541 < int32(0) {
		v779 = v541
		goto L77
	} else {
		goto L132
	}
L132:
	;
	goto L120
L133:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v599
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+4))
	if v599 < v602 {
		goto L149
	} else {
		goto L150
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v546
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v549
	v555 = v546 - int32(1)
	if v555 <= v549 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v552
	goto L133
L136:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557+v555))))
	if v559&int32(224) != int32(96) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	if int32(1)<<(uint(v559)%32)&int32(_a_F_danish_UTF_8_stem_7) == int32(0) {
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v572 = F_find_among_b(m, l0, int32(_a_F_danish_UTF_8_stem_8), int32(5))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L84
	} else {
		goto L139
	}
L139:
	;
	if v572 == int32(0) {
		goto L135
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v552
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v577
	switch v572 - int32(1) {
	case 0:
		goto L142
	case 1:
		goto L141
	default:
		goto L133
	}
L141:
	;
	v591 = F_slice_from_s(m, l0, int32(4), int32(_a_F_danish_UTF_8_stem_9))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L84
	} else {
		goto L147
	}
L142:
	;
	v581 = F_slice_del(m, l0)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L84
	} else {
		goto L143
	}
L143:
	;
	if v581 < int32(0) {
		v779 = v581
		goto L77
	} else {
		goto L144
	}
L144:
	;
	v585 = F_r_consonant_pair_2(m, l0)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L84
	} else {
		goto L145
	}
L145:
	;
	if int32(0) <= v585 {
		goto L133
	} else {
		goto L146
	}
L146:
	;
	v779 = v585
	goto L77
L147:
	;
	if int32(0) <= v591 {
		goto L133
	} else {
		goto L148
	}
L148:
	;
	v779 = v591
	goto L77
L149:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v776
	v779 = int32(1)
	goto L77
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v599
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v602
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L153
L151:
	;
	if v735 != 0 {
		goto L175
	} else {
		goto L176
	}
L152:
	;
	v735 = v728
	goto L151
L153:
	;
	if v623 <= v602 {
		v728 = int32(-1)
		goto L152
	} else {
		goto L155
	}
L154:
	;
	v728 = int32(0)
	goto L152
L155:
	;
	v641 = int32(1)
	v642 = v623 - v641
	v644 = int32(*(*int8)(unsafe.Add(mBase, uint32(v620+v642))))
	v646 = v644 & int32(255)
	if v642 == v602 {
		v701 = v646
		v702 = v641
		goto L156
	} else {
		goto L157
	}
L156:
	;
	if int32(122) < v701 {
		goto L165
	} else {
		goto L166
	}
L157:
	;
	if int32(0) <= v644 {
		v701 = v646
		v702 = v641
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v652 = v646 & int32(63)
	v654 = v623 - int32(2)
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620+v654))))
	v658 = v656 << (uint(int32(6)) % 32)
	if base.B2i32(v654 != v602)&base.B2i32(base.Ui32(v656) < base.Ui32(int32(192))) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v701 = v658&int32(1984) | v652
	v702 = int32(2)
	goto L156
L160:
	;
	goto L161
L161:
	;
	v671 = v658&int32(4032) | v652
	v673 = v623 - int32(3)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620+v673))))
	if base.B2i32(v673 != v602)&base.B2i32(base.Ui32(v675) < base.Ui32(int32(224))) == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v701 = v675<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v671
	v702 = int32(3)
	goto L156
L163:
	;
	goto L164
L164:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623+(v620-int32(4))))))
	v701 = v675<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_4) | v693&int32(7)<<(uint(int32(18))%32) | v671
	v702 = int32(4)
	goto L156
L165:
	;
	v735 = v702
	goto L151
L166:
	;
	goto L167
L167:
	;
	v706 = v701 - int32(98)
	if v706 < int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v735 = v702
	goto L151
L169:
	;
	goto L170
L170:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v706)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[2]))))
	if int32(base.Ui32(v712)>>(uint(v706&int32(7))%32))&int32(1) == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v735 = v702
	goto L151
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v623 - v702
	goto L174
L174:
	;
	goto L154
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v605
	goto L149
L176:
	;
	goto L177
L177:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v737
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)))
	v741 = F_slice_to(m, l0, v740)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L84
	} else {
		goto L178
	}
L178:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v743))) = v741
	if v741 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	return int32(-1)
L180:
	;
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v605
	v750 = int32(0)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v741-int32(4))))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v756-v605 < v755 {
		v767 = v750
		goto L183
	} else {
		goto L184
	}
L182:
	;
	if v767 == int32(0) {
		goto L149
	} else {
		goto L186
	}
L183:
	;
	goto L182
L184:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v763 = F_memcmp(m, v760+v756-v755, v741, v755)
	mBase = m.M
	if v763 != 0 {
		v767 = v750
		goto L183
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v756 - v755
	v767 = int32(1)
	goto L183
L186:
	;
	v770 = F_slice_del(m, l0)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L84
	} else {
		goto L187
	}
L187:
	;
	if v770 < int32(0) {
		v779 = v770
		goto L77
	} else {
		goto L188
	}
L188:
	;
	goto L149
}
func F_datanh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 float64
	_ = v26
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v37 float64
	_ = v37
	var v45 float64
	_ = v45
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v57 int64
	_ = v57
	var v79 float64
	_ = v79
	var v80 int64
	_ = v80
	var v85 int32
	_ = v85
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v117 float64
	_ = v117
	var v119 float64
	_ = v119
	var v123 float64
	_ = v123
	var v128 float64
	_ = v128
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v171 float64
	_ = v171
	var v180 float64
	_ = v180
	var v183 float64
	_ = v183
	var v188 float64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	if base.F64_gt(base.F64_abs(v6), float64(1)) == int32(0) {
		if base.F64_eq(v6, float64(-1)) != 0 {
			v15 = F_Float8GetDatum(m, math.Float64frombits(uint64(0xfff0000000000000)))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v15
			}
		} else {
			if base.F64_eq(v6, float64(1)) != 0 {
				v23 = F_Float8GetDatum(m, math.Float64frombits(uint64(0x7ff0000000000000)))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			} else {
				v26 = base.F64_abs(v6)
				v27 = base.I64_reinterpret_f64(v6)
				v32 = base.I32_wrap_i64(int64(base.Ui64(v27)>>(uint(int64(52))%64))) & int32(2047)
				if base.Ui32(v32) <= base.Ui32(int32(1021)) {
					if base.Ui32(v32) < base.Ui32(int32(991)) {
						v183 = v26
					} else {
						v37 = base.F64_add(v26, v26)
						v49 = base.F64_add(v37, base.F64_div(base.F64_mul(v26, v37), base.F64_sub(float64(1), v26)))
						v50 = float64(0)
						v57 = base.I64_reinterpret_f64(v49)
						if v57 <= int64(4601133429810003967) {
							if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v57) {
								if base.F64_eq(v49, float64(-1)) != 0 {
									v171 = math.Float64frombits(uint64(0xfff0000000000000))
									v180 = v171
								} else {
									v180 = base.F64_div(base.F64_sub(v49, v49), float64(0))
								}
							} else {
								if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v57)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
									v180 = v49
								} else {
									if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v57) {
										v79 = base.F64_add(v49, float64(1))
										v80 = base.I64_reinterpret_f64(v79)
										v85 = base.I32_wrap_i64(int64(base.Ui64(v80)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
										if base.Ui32(v85) <= base.Ui32(int32(1129316351)) {
											if base.Ui32(int32(1074790399)) < base.Ui32(v85) {
												v100 = base.F64_add(base.F64_sub(v49, v79), float64(1))
											} else {
												v100 = base.F64_sub(v49, base.F64_add(v79, float64(-1)))
											}
											v102 = base.F64_div(v100, v79)
										} else {
											v102 = v50
										}
										v117 = base.F64_add(base.F64_reinterpret_i64(v80&int64(4294967295)|base.I64_extend_i32_u(v85&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
										v119 = v102
										v123 = base.F64_convert_i32_s(int32(base.Ui32(v85)>>(uint(int32(20))%32)) - int32(1023))
									} else {
										v117 = v49
										v119 = v50
										v123 = float64(0)
									}
									v128 = base.F64_div(v117, base.F64_add(v117, float64(2)))
									v131 = base.F64_mul(v117, base.F64_mul(v117, float64(0.5)))
									v132 = base.F64_mul(v128, v128)
									v133 = base.F64_mul(v132, v132)
									v171 = base.F64_add(base.F64_mul(v123, float64(0.6931471803691238)), base.F64_add(v117, base.F64_sub(base.F64_add(base.F64_mul(v128, base.F64_add(v131, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v132, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v123, float64(1.9082149292705877e-10)), v119)), v131)))
									v180 = v171
								}
							}
						} else {
							if base.Ui64(int64(9218868437227405311)) < base.Ui64(v57) {
								v180 = v49
							} else {
								v79 = base.F64_add(v49, float64(1))
								v80 = base.I64_reinterpret_f64(v79)
								v85 = base.I32_wrap_i64(int64(base.Ui64(v80)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
								if base.Ui32(v85) <= base.Ui32(int32(1129316351)) {
									if base.Ui32(int32(1074790399)) < base.Ui32(v85) {
										v100 = base.F64_add(base.F64_sub(v49, v79), float64(1))
									} else {
										v100 = base.F64_sub(v49, base.F64_add(v79, float64(-1)))
									}
									v102 = base.F64_div(v100, v79)
								} else {
									v102 = v50
								}
								v117 = base.F64_add(base.F64_reinterpret_i64(v80&int64(4294967295)|base.I64_extend_i32_u(v85&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
								v119 = v102
								v123 = base.F64_convert_i32_s(int32(base.Ui32(v85)>>(uint(int32(20))%32)) - int32(1023))
								v128 = base.F64_div(v117, base.F64_add(v117, float64(2)))
								v131 = base.F64_mul(v117, base.F64_mul(v117, float64(0.5)))
								v132 = base.F64_mul(v128, v128)
								v133 = base.F64_mul(v132, v132)
								v171 = base.F64_add(base.F64_mul(v123, float64(0.6931471803691238)), base.F64_add(v117, base.F64_sub(base.F64_add(base.F64_mul(v128, base.F64_add(v131, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v132, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v123, float64(1.9082149292705877e-10)), v119)), v131)))
								v180 = v171
							}
						}
						v183 = base.F64_mul(v180, float64(0.5))
					}
				} else {
					v45 = base.F64_div(v26, base.F64_sub(float64(1), v26))
					v49 = base.F64_add(v45, v45)
					v50 = float64(0)
					v57 = base.I64_reinterpret_f64(v49)
					if v57 <= int64(4601133429810003967) {
						if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v57) {
							if base.F64_eq(v49, float64(-1)) != 0 {
								v171 = math.Float64frombits(uint64(0xfff0000000000000))
								v180 = v171
							} else {
								v180 = base.F64_div(base.F64_sub(v49, v49), float64(0))
							}
						} else {
							if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v57)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
								v180 = v49
							} else {
								if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v57) {
									v79 = base.F64_add(v49, float64(1))
									v80 = base.I64_reinterpret_f64(v79)
									v85 = base.I32_wrap_i64(int64(base.Ui64(v80)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
									if base.Ui32(v85) <= base.Ui32(int32(1129316351)) {
										if base.Ui32(int32(1074790399)) < base.Ui32(v85) {
											v100 = base.F64_add(base.F64_sub(v49, v79), float64(1))
										} else {
											v100 = base.F64_sub(v49, base.F64_add(v79, float64(-1)))
										}
										v102 = base.F64_div(v100, v79)
									} else {
										v102 = v50
									}
									v117 = base.F64_add(base.F64_reinterpret_i64(v80&int64(4294967295)|base.I64_extend_i32_u(v85&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
									v119 = v102
									v123 = base.F64_convert_i32_s(int32(base.Ui32(v85)>>(uint(int32(20))%32)) - int32(1023))
								} else {
									v117 = v49
									v119 = v50
									v123 = float64(0)
								}
								v128 = base.F64_div(v117, base.F64_add(v117, float64(2)))
								v131 = base.F64_mul(v117, base.F64_mul(v117, float64(0.5)))
								v132 = base.F64_mul(v128, v128)
								v133 = base.F64_mul(v132, v132)
								v171 = base.F64_add(base.F64_mul(v123, float64(0.6931471803691238)), base.F64_add(v117, base.F64_sub(base.F64_add(base.F64_mul(v128, base.F64_add(v131, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v132, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v123, float64(1.9082149292705877e-10)), v119)), v131)))
								v180 = v171
							}
						}
					} else {
						if base.Ui64(int64(9218868437227405311)) < base.Ui64(v57) {
							v180 = v49
						} else {
							v79 = base.F64_add(v49, float64(1))
							v80 = base.I64_reinterpret_f64(v79)
							v85 = base.I32_wrap_i64(int64(base.Ui64(v80)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
							if base.Ui32(v85) <= base.Ui32(int32(1129316351)) {
								if base.Ui32(int32(1074790399)) < base.Ui32(v85) {
									v100 = base.F64_add(base.F64_sub(v49, v79), float64(1))
								} else {
									v100 = base.F64_sub(v49, base.F64_add(v79, float64(-1)))
								}
								v102 = base.F64_div(v100, v79)
							} else {
								v102 = v50
							}
							v117 = base.F64_add(base.F64_reinterpret_i64(v80&int64(4294967295)|base.I64_extend_i32_u(v85&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v119 = v102
							v123 = base.F64_convert_i32_s(int32(base.Ui32(v85)>>(uint(int32(20))%32)) - int32(1023))
							v128 = base.F64_div(v117, base.F64_add(v117, float64(2)))
							v131 = base.F64_mul(v117, base.F64_mul(v117, float64(0.5)))
							v132 = base.F64_mul(v128, v128)
							v133 = base.F64_mul(v132, v132)
							v171 = base.F64_add(base.F64_mul(v123, float64(0.6931471803691238)), base.F64_add(v117, base.F64_sub(base.F64_add(base.F64_mul(v128, base.F64_add(v131, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v132, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v123, float64(1.9082149292705877e-10)), v119)), v131)))
							v180 = v171
						}
					}
					v183 = base.F64_mul(v180, float64(0.5))
				}
				if v27 < int64(0) {
					v188 = base.F64_neg(v183)
				} else {
					v188 = v183
				}
				v189 = F_Float8GetDatum(m, v188)
				mBase = m.M
				v190 = m.ExcPending
				if v190 != 0 {
					return int32(0)
				} else {
					return v189
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v195 = m.ExcPending
		if v195 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v198 = m.ExcPending
			if v198 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_datanh_2), int32(0))
				mBase = m.M
				v202 = m.ExcPending
				if v202 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_datanh_3), int32(2727), int32(_a_F_datanh_4))
					mBase = m.M
					v207 = m.ExcPending
					if v207 != 0 {
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
func F_dbase_redo(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v315 int64
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
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
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int64
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	v17 = v15 & int32(240)
	switch v17 - int32(16) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L4
	case 16:
		goto L5
	default:
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L8
	} else {
		goto L198
	}
L2:
	;
	m.G0 = v12 + int32(160)
	return
L3:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	v566 = F_GetDatabasePath(m, v564, v565)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L8
	} else {
		goto L167
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L8
	} else {
		goto L164
	}
L5:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[0]))
	if base.Ui32(int32(2)) <= base.Ui32(v328) {
		goto L101
	} else {
		goto L102
	}
L6:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v23 = F_GetDatabasePath(m, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v27 = F_GetDatabasePath(m, v25, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v58 = F_pstrdup(m, v27)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L21
	}
L11:
	;
	v33 = F___fstatat(m, int32(-100), v27, v12-int32(-64), int32(0))
	mBase = m.M
	goto L12
L12:
	;
	if v33 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	if v34&int32(_a_F_dbase_redo_0) != int32(_a_F_dbase_redo_1) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v39 = F_rmtree(m, v27)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v39 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v43 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	if v43 == int32(0) {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v27
	F_errmsg(m, int32(_a_F_dbase_redo_2), v12+int32(32))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3340), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v63 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v113 = F___fstatat(m, int32(-100), v58, v12-int32(-64), int32(0))
	mBase = m.M
	goto L47
L23:
	;
	v64 = F_strlen(m, v58)
	mBase = m.M
	v67 = v64 + v58
	goto L26
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	v71 = v67 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v72 == int32(47) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v77 = v71
	goto L32
L28:
	;
	if base.Ui32(v58) < base.Ui32(v71) {
		v67 = v71
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L30
L32:
	;
	if base.Ui32(v58) < base.Ui32(v77) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v89 = v77
	goto L38
L34:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v83 != int32(47) {
		v77 = v77 - int32(1)
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	goto L36
L38:
	;
	if base.Ui32(v58) < base.Ui32(v89) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v58 == v89 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v93 = v89 - int32(1)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v94 == int32(47) {
		v89 = v93
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	goto L42
L44:
	;
	v102 = v58 + base.B2i32(v63 == int32(47))
	goto L46
L45:
	;
	v102 = v89
	goto L46
L46:
	;
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v103)
	goto L25
L47:
	;
	if v113 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[1]))
	if v117 != int32(44) {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_pfree(m, v58)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L53
	}
L51:
	;
	F_recovery_create_dbdir(m, v58, int32(1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v129 = F___fstatat(m, int32(-100), v23, v12-int32(-64), int32(0))
	mBase = m.M
	goto L55
L54:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v140 = m.G0
	v142 = v140 - int32(32)
	m.G0 = v142
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[2]))
	if int32(0) < v145 {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	if int32(0) <= v129 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[1]))
	if v133 != int32(44) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	F_recovery_create_dbdir(m, v23, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	v154 = int32(0)
	goto L62
L60:
	;
	goto L61
L61:
	;
	m.G0 = v142 + int32(32)
	v315 = F_EmitProcSignalBarrier(m)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L8
	} else {
		goto L96
	}
L62:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[3]))
	v161 = v158 + v154<<(uint(int32(6))%32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v162 != v139 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L61
L64:
	;
	v299 = v154 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[2]))
	if v299 < v301 {
		v154 = v299
		goto L62
	} else {
		goto L95
	}
L65:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[4]))
	F_ResourceOwnerEnlarge(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+28)) = int32(_a_F_dbase_redo_5)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+24)) = int32(_a_F_dbase_redo_6)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = int32(_a_F_dbase_redo_7)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v142)+8)) = int64(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	v181 = int32(_a_F_dbase_redo_8)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = v180 | v181
	if v180&v181 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	goto L71
L69:
	;
	v206 = v180
	goto L70
L70:
	;
	v217 = int32(_a_F_dbase_redo_9)
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[5]))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v142+int32(8))+8))
	if v220 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	F_perform_spin_delay(m, v142+int32(8))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L73
	}
L72:
	;
	v206 = v199
	goto L70
L73:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	v200 = int32(_a_F_dbase_redo_8)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = v199 | v200
	if v199&v200 != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v237 = int32(25165824)
	if v206&v237 != v237 {
		goto L86
	} else {
		goto L87
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[5])) = v235
	goto L76
L78:
	;
	if int32(999) < v218 {
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v218 < int32(11) {
		goto L76
	} else {
		goto L85
	}
L81:
	;
	v225 = int32(900)
	if v225 <= v218 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v228 = v225
	goto L84
L83:
	;
	v228 = v218
	goto L84
L84:
	;
	v235 = v228 + int32(100)
	goto L77
L85:
	;
	v235 = v218 - int32(1)
	goto L77
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = v206 & int32(-4194305)
	goto L64
L87:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v241 != v139 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	v244 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = (v243 + v244) & int32(-4194305)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	v250 = int32(_a_F_dbase_redo_10)
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[6])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v244
	v258 = v249 + v244
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v258
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[4]))
	F_ResourceOwnerRemember(m, v261, v258, int32(_a_F_dbase_redo_11))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v266 = v161 + int32(48)
	v268 = F_LWLockAcquire(m, v266, int32(1))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_FlushBuffer(m, v161, int32(0), int32(3))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	F_LWLockRelease(m, v266)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[4]))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	F_ResourceOwnerForget(m, v277, v278+int32(1), int32(_a_F_dbase_redo_11))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_UnpinBufferNoOwner(m, v161)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	goto L64
L95:
	;
	goto L63
L96:
	;
	F_WaitForProcSignalBarrier(m, v315)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	F_copydir(m, v23, v27, int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L8
	} else {
		goto L98
	}
L98:
	;
	F_pfree(m, v23)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	F_pfree(m, v27)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	goto L2
L101:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	F_LockSharedObjectForSession(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L8
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	F_ReplicationSlotsDropDBSlots(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L8
	} else {
		goto L115
	}
L104:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	v335 = F_CountDBBackends(m, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	if int32(0) < v335 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	goto L109
L107:
	;
	goto L108
L108:
	;
	goto L103
L109:
	;
	F_CancelDBBackends(m, v334, int32(7), int32(1))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L8
	} else {
		goto L111
	}
L110:
	;
	goto L108
L111:
	;
	F_pg_usleep(m, int32(_a_F_dbase_redo_12))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	v355 = F_CountDBBackends(m, v334)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L8
	} else {
		goto L113
	}
L113:
	;
	if int32(0) < v355 {
		goto L109
	} else {
		goto L114
	}
L114:
	;
	goto L110
L115:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	F_DropDatabaseBuffers(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L8
	} else {
		goto L116
	}
L116:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	F_ForgetDatabaseSyncRequests(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L8
	} else {
		goto L117
	}
L117:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	v387 = m.G0
	v389 = v387 - int32(112)
	m.G0 = v389
	F_smgrdestroyall(m)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L8
	} else {
		goto L118
	}
L118:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[7]))
	if v394 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v483 = F_EmitProcSignalBarrier(m)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L8
	} else {
		goto L145
	}
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
	} else {
		goto L142
	}
L121:
	;
	m.G0 = v389 + int32(112)
	goto L119
L122:
	;
	F_hash_seq_init(m, v389+int32(92), v394)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L8
	} else {
		goto L123
	}
L123:
	;
	v403 = F_hash_seq_search(m, v389+int32(92))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L8
	} else {
		goto L124
	}
L124:
	;
	if v403 == int32(0) {
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v407 = v403
	goto L126
L126:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v386 == v416 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L121
L128:
	;
	v420 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L8
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v456 = F_hash_seq_search(m, v389+int32(92))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L8
	} else {
		goto L140
	}
L131:
	;
	if v420 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v407)+8))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v407)+12))
	F_GetRelationPath(m, v389+int32(20), v425, v426, v427, int32(-1), v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L8
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[7]))
	v449 = F_hash_search(m, v446, v407, int32(2), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L8
	} else {
		goto L138
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v389 + int32(20)
	F_errmsg_internal(m, int32(_a_F_dbase_redo_13), v389)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L8
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_14), int32(212), int32(_a_F_dbase_redo_15))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L8
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	if v449 == int32(0) {
		goto L120
	} else {
		goto L139
	}
L139:
	;
	goto L130
L140:
	;
	if v456 != 0 {
		v407 = v456
		goto L126
	} else {
		goto L141
	}
L141:
	;
	goto L127
L142:
	;
	F_errmsg_internal(m, int32(_a_F_dbase_redo_16), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L8
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_14), int32(217), int32(_a_F_dbase_redo_15))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L8
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_WaitForProcSignalBarrier(m, v483)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L8
	} else {
		goto L146
	}
L146:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if int32(0) < v487 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v493 = int32(0)
	goto L150
L148:
	;
	goto L149
L149:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[0]))
	if base.Ui32(v544) < base.Ui32(int32(2)) {
		goto L2
	} else {
		goto L162
	}
L150:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v326+int32(8)+v493<<(uint(int32(2))%32))))
	v507 = F_GetDatabasePath(m, v502, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L8
	} else {
		goto L153
	}
L151:
	;
	goto L149
L152:
	;
	F_pfree(m, v507)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L8
	} else {
		goto L160
	}
L153:
	;
	v509 = F_rmtree(m, v507)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L8
	} else {
		goto L154
	}
L154:
	;
	if v509 != 0 {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v513 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	if v513 == int32(0) {
		goto L152
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v507
	F_errmsg(m, int32(_a_F_dbase_redo_2), v12+int32(48))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L8
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3454), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L8
	} else {
		goto L159
	}
L159:
	;
	goto L152
L160:
	;
	v531 = v493 + int32(1)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v531 < v532 {
		v493 = v531
		goto L150
	} else {
		goto L161
	}
L161:
	;
	goto L151
L162:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	F_UnlockSharedObjectForSession(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L8
	} else {
		goto L163
	}
L163:
	;
	goto L2
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v17
	F_errmsg_internal(m, int32(_a_F_dbase_redo_17), v12)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L8
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3471), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L8
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	v568 = F_pstrdup(m, v566)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L8
	} else {
		goto L168
	}
L168:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	if v573 != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	F_recovery_create_dbdir(m, v568, int32(1))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L8
	} else {
		goto L194
	}
L170:
	;
	v574 = F_strlen(m, v568)
	mBase = m.M
	v577 = v574 + v568
	goto L173
L171:
	;
	goto L172
L172:
	;
	goto L169
L173:
	;
	v581 = v577 - int32(1)
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581))))
	if v582 == int32(47) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v587 = v581
	goto L179
L175:
	;
	if base.Ui32(v568) < base.Ui32(v581) {
		v577 = v581
		goto L173
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	goto L174
L178:
	;
	goto L177
L179:
	;
	if base.Ui32(v568) < base.Ui32(v587) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v599 = v587
	goto L185
L181:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	if v593 != int32(47) {
		v587 = v587 - int32(1)
		goto L179
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	goto L180
L184:
	;
	goto L183
L185:
	;
	if base.Ui32(v568) < base.Ui32(v599) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	if v568 == v599 {
		goto L191
	} else {
		goto L192
	}
L187:
	;
	v603 = v599 - int32(1)
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
	if v604 == int32(47) {
		v599 = v603
		goto L185
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	goto L186
L190:
	;
	goto L189
L191:
	;
	v612 = v568 + base.B2i32(v573 == int32(47))
	goto L193
L192:
	;
	v612 = v599
	goto L193
L193:
	;
	v613 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v612))) = uint8(v613)
	goto L172
L194:
	;
	F_pfree(m, v568)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L8
	} else {
		goto L195
	}
L195:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	F_CreateDirAndVersionFile(m, v566, v624, v625, int32(1))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L8
	} else {
		goto L196
	}
L196:
	;
	F_pfree(m, v566)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L8
	} else {
		goto L197
	}
L197:
	;
	goto L2
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v27
	F_errmsg(m, int32(_a_F_dbase_redo_18), v12+int32(16))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L8
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3354), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L8
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_decompile_conbin(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
	v11 = v9 + v10
	v15 = F_heap_getattr_6(m, l0, int32(28), l1, v7+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v19 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v26
				F_errmsg_internal(m, int32(_a_F_decompile_conbin_0), v7)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_decompile_conbin_1), int32(_a_F_decompile_conbin_2), int32(_a_F_decompile_conbin_3))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
			v39 = F_DirectFunctionCall2Coll(m, int32(577), int32(0), v15, v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v41 = F_text_to_cstring(m, v39)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v41
				}
			}
		}
	}
}
func F_deparse_context_for(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_palloc0(m, int32(80))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = F_palloc0(m, int32(136))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(1)
			v20 = int32(114)
			*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l1
			v23 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(101)
			v28 = F_makeAlias(m, l0, v23)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v28
				v32 = int32(256)
				*(*uint16)(unsafe.Add(mBase, uint32(v16)+124)) = uint16(v32)
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16
				v41 = F_list_make1_impl(m, int32(1), v8+int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v43
					*(*int64)(unsafe.Add(mBase, uint32(v11)+12)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v41
					F_set_rtable_names(m, v11, v43, v43)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_set_simple_column_names(m, v11)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v11
							v57 = F_list_make1_impl(m, int32(1), v8)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(16)
								return v57
							}
						}
					}
				}
			}
		}
	}
}
func F_die(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_die[0])))
	if v3 == int32(0) {
		v7 = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_die[1])) = v7
		*(*int32)(unsafe.Add(mBase, _c_F_die[2])) = v7
	} else {
	}
	*(*int32)(unsafe.Add(mBase, _c_F_die[3])) = int32(4)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_die[4]))
	F_SetLatch(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_die[5])))
		if v20 == int32(0) {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_die[6]))
			if v24 == int32(2) {
				return
			} else {
				F_ProcessInterrupts(m)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_do_serialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v10 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v14 = F_pg_vsnprintf(m, v12, v13, l2, l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
					F_errmsg_internal(m, int32(_a_F_do_serialize_0), v8)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_do_serialize_1), int32(_a_F_do_serialize_2), int32(_a_F_do_serialize_3))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if base.Ui32(v18) <= base.Ui32(v14) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_do_serialize_4), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_do_serialize_1), int32(_a_F_do_serialize_5), int32(_a_F_do_serialize_3))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = v14 + int32(1)
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v21 + v22
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25 - v21
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_do_serialize_4), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_do_serialize_1), int32(_a_F_do_serialize_6), int32(_a_F_do_serialize_3))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
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
func F_downcase_identifier(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v17 = F_palloc(m, l1+int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if int32(0) < l1 {
			if l1 != int32(1) {
				v39 = v5
				v43 = v5
				for {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v39))))
					if base.Ui32((v47-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v56 = v47 | int32(32)
					} else {
						v56 = v47
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v39+v17))) = uint8(v56)
					v59 = v39 | int32(1)
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v59))))
					if base.Ui32((v62-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v71 = v62 | int32(32)
					} else {
						v71 = v62
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v17+v59))) = uint8(v71)
					v73 = int32(2)
					v74 = v39 + v73
					v76 = v43 + v73
					if v76 != l1&int32(2147483646) {
						v39 = v74
						v43 = v76
						continue
					} else {
						break
					}
					break
				}
				v82 = v74
			} else {
				v82 = v5
			}
			if l1&int32(1) != 0 {
				v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v82))))
				if base.Ui32((v92-int32(65))&int32(255)) < base.Ui32(int32(26)) {
					v101 = v92 | int32(32)
				} else {
					v101 = v92
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v82+v17))) = uint8(v101)
			} else {
			}
			v105 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1+v17))) = uint8(v105)
			if l3 == v105 {
				m.G0 = v13 + int32(16)
				return v17
			} else {
				if base.Ui32(l1) < base.Ui32(int32(64)) {
					m.G0 = v13 + int32(16)
					return v17
				} else {
					v112 = F_pg_mbcliplen(m, v17, l1, int32(63))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						if l2 == int32(0) {
							v137 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v112+v17))) = uint8(v137)
							m.G0 = v13 + int32(16)
							return v17
						} else {
							v118 = F_errstart(m, int32(18), int32(0))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								if v118 == int32(0) {
									v137 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v112+v17))) = uint8(v137)
									m.G0 = v13 + int32(16)
									return v17
								} else {
									F_errcode(m, int32(34103428))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v17
										*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v112
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = v17
										F_errmsg(m, int32(_a_F_downcase_identifier_0), v13)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_downcase_identifier_1), int32(102), int32(_a_F_downcase_identifier_2))
											mBase = m.M
											v135 = m.ExcPending
											if v135 != 0 {
												return int32(0)
											} else {
												v137 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v112+v17))) = uint8(v137)
												m.G0 = v13 + int32(16)
												return v17
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
			v139 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v139)
			m.G0 = v13 + int32(16)
			return v17
		}
	}
}
func F_dsnowball_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
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
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
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
	var v181 int32
	_ = v181
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
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = F_str_tolower(m, v7, v8, int32(100))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = F_palloc0(m, int32(16))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(1001) <= v8 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v10
	return v15
L5:
	;
	goto L6
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
	if v31 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v24 = F_searchstoplist(m, v6+int32(4), v10)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_pfree(m, v10)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	if v24 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	return v15
L14:
	;
	v98 = int32(_a_F_dsnowball_lexize_0)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v97&int32(3) == int32(0) {
		v126 = v97
		goto L42
	} else {
		goto L43
	}
L15:
	;
	v97 = v10
	goto L14
L16:
	;
	goto L17
L17:
	;
	if v10&int32(3) == int32(0) {
		v57 = v10
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v92 = F_pg_server_to_any(m, v10, v90, int32(6))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L35
	}
L19:
	;
	v90 = v82 - v10
	goto L18
L20:
	;
	v61 = v57
	goto L29
L21:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v41 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v90 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v46 = v10
	goto L25
L25:
	;
	v50 = v46 + int32(1)
	if v50&int32(3) == int32(0) {
		v57 = v50
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v82 = v50
	goto L19
L27:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v55 != 0 {
		v46 = v50
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v70 = int32(-2139062144)
	if (int32(16843008)-v67|v67)&v70 == v70 {
		v61 = v61 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v76 = v61
	goto L32
L31:
	;
	goto L30
L32:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 != 0 {
		v76 = v76 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v82 = v76
	goto L19
L34:
	;
	goto L33
L35:
	;
	if v10 == v92 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v97 = v10
	goto L14
L37:
	;
	goto L38
L38:
	;
	F_pfree(m, v10)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v97 = v92
	goto L14
L40:
	;
	v160 = F_SN_set_current(m, v102, v159, v97)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L57
	}
L41:
	;
	v159 = v151 - v97
	goto L40
L42:
	;
	v130 = v126
	goto L51
L43:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v110 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v159 = int32(0)
	goto L40
L45:
	;
	goto L46
L46:
	;
	v115 = v97
	goto L47
L47:
	;
	v119 = v115 + int32(1)
	if v119&int32(3) == int32(0) {
		v126 = v119
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v151 = v119
	goto L41
L49:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v124 != 0 {
		v115 = v119
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v139 = int32(-2139062144)
	if (int32(16843008)-v136|v136)&v139 == v139 {
		v130 = v130 + int32(4)
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v145 = v130
	goto L54
L53:
	;
	goto L52
L54:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v149 != 0 {
		v145 = v145 + int32(1)
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v151 = v145
	goto L41
L56:
	;
	goto L55
L57:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v164 = m.T0[v163].(func(*base.Module, int32) int32)(m, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v99
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v168 == int32(0) {
		v189 = v97
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
	if v190 != int32(1) {
		goto L67
	} else {
		goto L68
	}
L60:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	if v171 == int32(0) {
		v189 = v97
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v176 = F_repalloc(m, v97, v171+int32(1))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	if v180 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182+v184))) = uint8(v186)
	v189 = v176
	goto L59
L64:
	;
	v181 = F__emscripten_memcpy_bulkmem(m, v176, v179, v180)
	mBase = m.M
	v182 = v181
	goto L66
L65:
	;
	v182 = v176
	goto L66
L66:
	;
	goto L63
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v189
	return v15
L68:
	;
	goto L69
L69:
	;
	if v189&int32(3) == int32(0) {
		v218 = v189
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v253 = F_pg_any_to_server(m, v189, v251, int32(6))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L87
	}
L71:
	;
	v251 = v243 - v189
	goto L70
L72:
	;
	v222 = v218
	goto L81
L73:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v202 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v251 = int32(0)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v207 = v189
	goto L77
L77:
	;
	v211 = v207 + int32(1)
	if v211&int32(3) == int32(0) {
		v218 = v211
		goto L72
	} else {
		goto L79
	}
L78:
	;
	v243 = v211
	goto L71
L79:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v216 != 0 {
		v207 = v211
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v231 = int32(-2139062144)
	if (int32(16843008)-v228|v228)&v231 == v231 {
		v222 = v222 + int32(4)
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v237 = v222
	goto L84
L83:
	;
	goto L82
L84:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v241 != 0 {
		v237 = v237 + int32(1)
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v243 = v237
	goto L71
L86:
	;
	goto L85
L87:
	;
	if v189 == v253 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v189
	return v15
L89:
	;
	goto L90
L90:
	;
	F_pfree(m, v189)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v253
	return v15
}
func F_dupnfa(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	if l1 == l2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_dupnfa[0]))
	if v8 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l4
	F_duptraverse(m, l0, l1, l3)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L31
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v11 <= v12 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	F_createarc(m, l0, int32(110), int32(0), l3, l4)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L30
	}
L11:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v14 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v29 == int32(0) {
		goto L10
	} else {
		goto L22
	}
L14:
	;
	v19 = v14
	goto L15
L15:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v22 != l4 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v28 != 0 {
		v19 = v28
		goto L15
	} else {
		goto L21
	}
L18:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	if v24 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v25 == int32(110) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	goto L16
L22:
	;
	v34 = v29
	goto L23
L23:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v37 != l3 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L10
L25:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v43 != 0 {
		v34 = v43
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+4)))
	if v39 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v40 == int32(110) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	goto L24
L30:
	;
	return
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = int32(0)
	F_cleartraverse(m, l0, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	goto L1
}
