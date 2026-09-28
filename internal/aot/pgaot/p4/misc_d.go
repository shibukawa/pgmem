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
	var v44 int32
	_ = v44
	var v46 int64
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
	var v71 int32
	_ = v71
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
	var v98 int32
	_ = v98
	var v100 int64
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
	var v148 int32
	_ = v148
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
	return v148
L3:
	;
	if v11 == int32(0) {
		v148 = v2
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v148 = int32(1)
	goto L2
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[0]))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_DeadLockCheckRecurse[1]))
	if v20 <= v18 {
		v148 = int32(1)
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v44
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v48
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
		v148 = v33
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
	v71 = v50
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
	v97 = v91 + v23*v87 + v71*v87
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+16)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v97)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = v102
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
		v148 = v33
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
	v119 = v71 + v115
	if v11 != v119 {
		v70 = v116
		v71 = v119
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
	var v50 int32
	_ = v50
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v110 int64
	_ = v110
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v143 int64
	_ = v143
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v173 int64
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v188 int64
	_ = v188
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v209 int32
	_ = v209
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v245 int64
	_ = v245
	var v252 int64
	_ = v252
	var v264 int32
	_ = v264
	var v265 int64
	_ = v265
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v288 int64
	_ = v288
	var v295 int64
	_ = v295
	var v307 int32
	_ = v307
	var v308 int64
	_ = v308
	var v311 int64
	_ = v311
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v335 int64
	_ = v335
	var v342 int64
	_ = v342
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v367 int64
	_ = v367
	var v370 int64
	_ = v370
	var v371 int64
	_ = v371
	var v373 int64
	_ = v373
	var v374 int64
	_ = v374
	var v378 int64
	_ = v378
	var v385 int64
	_ = v385
	var v397 int32
	_ = v397
	var v398 int64
	_ = v398
	var v401 int64
	_ = v401
	var v404 int64
	_ = v404
	var v405 int64
	_ = v405
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v414 int64
	_ = v414
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v420 int64
	_ = v420
	var v421 int64
	_ = v421
	var v425 int64
	_ = v425
	var v432 int64
	_ = v432
	var v444 int32
	_ = v444
	var v445 int64
	_ = v445
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v454 int64
	_ = v454
	var v455 int64
	_ = v455
	var v457 int64
	_ = v457
	var v460 int64
	_ = v460
	var v461 int64
	_ = v461
	var v463 int64
	_ = v463
	var v464 int64
	_ = v464
	var v468 int64
	_ = v468
	var v475 int64
	_ = v475
	var v487 int32
	_ = v487
	var v488 int64
	_ = v488
	var v491 int64
	_ = v491
	var v494 int64
	_ = v494
	var v495 int64
	_ = v495
	var v501 int64
	_ = v501
	var v502 int64
	_ = v502
	var v504 int64
	_ = v504
	var v507 int64
	_ = v507
	var v508 int64
	_ = v508
	var v510 int64
	_ = v510
	var v511 int64
	_ = v511
	var v515 int64
	_ = v515
	var v522 int64
	_ = v522
	var v534 int32
	_ = v534
	var v535 int64
	_ = v535
	var v537 int64
	_ = v537
	var v538 int64
	_ = v538
	var v544 int64
	_ = v544
	var v545 int64
	_ = v545
	var v547 int64
	_ = v547
	var v550 int64
	_ = v550
	var v551 int64
	_ = v551
	var v553 int64
	_ = v553
	var v554 int64
	_ = v554
	var v558 int64
	_ = v558
	var v565 int64
	_ = v565
	var v577 int32
	_ = v577
	var v578 int64
	_ = v578
	var v579 int64
	_ = v579
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v587 int64
	_ = v587
	var v593 int64
	_ = v593
	var v594 int64
	_ = v594
	var v596 int64
	_ = v596
	var v599 int64
	_ = v599
	var v600 int64
	_ = v600
	var v602 int64
	_ = v602
	var v603 int64
	_ = v603
	var v607 int64
	_ = v607
	var v614 int64
	_ = v614
	var v626 int32
	_ = v626
	var v628 int64
	_ = v628
	var v629 int64
	_ = v629
	var v635 int64
	_ = v635
	var v636 int64
	_ = v636
	var v638 int64
	_ = v638
	var v641 int64
	_ = v641
	var v642 int64
	_ = v642
	var v644 int64
	_ = v644
	var v645 int64
	_ = v645
	var v649 int64
	_ = v649
	var v656 int64
	_ = v656
	var v668 int32
	_ = v668
	var v669 int64
	_ = v669
	var v671 int64
	_ = v671
	var v672 int64
	_ = v672
	var v673 int64
	_ = v673
	var v674 int64
	_ = v674
	var v682 int64
	_ = v682
	var v688 int64
	_ = v688
	var v689 int64
	_ = v689
	var v691 int64
	_ = v691
	var v694 int64
	_ = v694
	var v695 int64
	_ = v695
	var v697 int64
	_ = v697
	var v698 int64
	_ = v698
	var v702 int64
	_ = v702
	var v709 int64
	_ = v709
	var v721 int32
	_ = v721
	var v723 int64
	_ = v723
	var v724 int64
	_ = v724
	var v730 int64
	_ = v730
	var v731 int64
	_ = v731
	var v733 int64
	_ = v733
	var v736 int64
	_ = v736
	var v737 int64
	_ = v737
	var v739 int64
	_ = v739
	var v740 int64
	_ = v740
	var v744 int64
	_ = v744
	var v751 int64
	_ = v751
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int64
	_ = v766
	var v767 int64
	_ = v767
	var v768 int64
	_ = v768
	var v769 int64
	_ = v769
	var v772 int64
	_ = v772
	var v773 int64
	_ = v773
	var v776 int64
	_ = v776
	var v778 int64
	_ = v778
	var v779 int64
	_ = v779
	var v780 int64
	_ = v780
	var v782 int64
	_ = v782
	var v784 int64
	_ = v784
	var v786 int64
	_ = v786
	var v787 int64
	_ = v787
	var v789 int64
	_ = v789
	var v791 int64
	_ = v791
	var v796 int64
	_ = v796
	var v808 int64
	_ = v808
	var v810 int64
	_ = v810
	var v812 int64
	_ = v812
	var v815 int64
	_ = v815
	var v816 int64
	_ = v816
	var v818 int64
	_ = v818
	var v823 int64
	_ = v823
	var v825 int64
	_ = v825
	var v831 int64
	_ = v831
	var v833 int64
	_ = v833
	var v844 int64
	_ = v844
	var v849 int64
	_ = v849
	var v850 int64
	_ = v850
	var v852 int64
	_ = v852
	var v856 int64
	_ = v856
	var v858 int64
	_ = v858
	var v862 int64
	_ = v862
	var v866 int64
	_ = v866
	var v868 int64
	_ = v868
	var v870 int64
	_ = v870
	var v872 int64
	_ = v872
	var v886 int64
	_ = v886
	var v890 int64
	_ = v890
	var v892 int64
	_ = v892
	var v900 int64
	_ = v900
	var v905 int64
	_ = v905
	var v909 int64
	_ = v909
	var v914 int64
	_ = v914
	var v920 int64
	_ = v920
	var v925 int64
	_ = v925
	var v928 int64
	_ = v928
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int64
	_ = v936
	var v937 int64
	_ = v937
	var v945 int64
	_ = v945
	var v950 int64
	_ = v950
	var v951 int64
	_ = v951
	var v953 int64
	_ = v953
	var v956 int64
	_ = v956
	var v957 int64
	_ = v957
	var v959 int64
	_ = v959
	var v960 int64
	_ = v960
	var v964 int64
	_ = v964
	var v971 int64
	_ = v971
	var v984 int32
	_ = v984
	var v989 int64
	_ = v989
	var v991 int64
	_ = v991
	var v992 int64
	_ = v992
	var v999 int32
	_ = v999
	var v1002 int64
	_ = v1002
	var v1004 int64
	_ = v1004
	var v1006 int64
	_ = v1006
	var v1011 int64
	_ = v1011
	var v1012 int64
	_ = v1012
	var v1014 int64
	_ = v1014
	var v1017 int64
	_ = v1017
	var v1018 int64
	_ = v1018
	var v1020 int64
	_ = v1020
	var v1021 int64
	_ = v1021
	var v1025 int64
	_ = v1025
	var v1032 int64
	_ = v1032
	var v1045 int64
	_ = v1045
	var v1047 int64
	_ = v1047
	var v1048 int64
	_ = v1048
	var v1055 int64
	_ = v1055
	var v1056 int64
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int64
	_ = v1059
	var v1060 int64
	_ = v1060
	var v1062 int64
	_ = v1062
	var v1063 int64
	_ = v1063
	var v1071 int64
	_ = v1071
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1105 int64
	_ = v1105
	var v1109 int64
	_ = v1109
	var v1110 int64
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1128 int64
	_ = v1128
	var v1136 int64
	_ = v1136
	var v1137 int64
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1143 int64
	_ = v1143
	var v1144 int64
	_ = v1144
	var v1149 int64
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1152 int64
	_ = v1152
	var v1155 int64
	_ = v1155
	var v1156 int64
	_ = v1156
	var v1158 int64
	_ = v1158
	var v1159 int64
	_ = v1159
	var v1163 int64
	_ = v1163
	var v1170 int64
	_ = v1170
	var v1181 int64
	_ = v1181
	var v1182 int64
	_ = v1182
	var v1183 int64
	_ = v1183
	var v1185 int64
	_ = v1185
	var v1190 int64
	_ = v1190
	var v1192 int64
	_ = v1192
	var v1197 int64
	_ = v1197
	var v1200 int64
	_ = v1200
	var v1201 int64
	_ = v1201
	var v1202 int64
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1205 int64
	_ = v1205
	var v1206 int64
	_ = v1206
	var v1211 int64
	_ = v1211
	var v1214 int64
	_ = v1214
	var v1217 int64
	_ = v1217
	var v1220 int64
	_ = v1220
	var v1221 int64
	_ = v1221
	var v1225 int64
	_ = v1225
	var v1232 int64
	_ = v1232
	var v1243 int64
	_ = v1243
	var v1244 int64
	_ = v1244
	var v1249 int64
	_ = v1249
	var v1252 int64
	_ = v1252
	var v1255 int64
	_ = v1255
	var v1258 int64
	_ = v1258
	var v1259 int64
	_ = v1259
	var v1263 int64
	_ = v1263
	var v1270 int64
	_ = v1270
	var v1282 int64
	_ = v1282
	var v1283 int64
	_ = v1283
	var v1287 int64
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1292 int64
	_ = v1292
	var v1295 int64
	_ = v1295
	var v1298 int64
	_ = v1298
	var v1300 int64
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1306 int64
	_ = v1306
	var v1309 int64
	_ = v1309
	var v1312 int64
	_ = v1312
	var v1314 int64
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1320 int64
	_ = v1320
	var v1325 int64
	_ = v1325
	var v1335 int64
	_ = v1335
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
	v50 = base.I32_wrap_i64(int64(base.Ui64(l2)>>(uint(v37)%64))) & v40
	if base.B2i32(base.Ui32(int32(-32767)) < base.Ui32(v41-v40))&base.B2i32(base.Ui32(int32(-32766)) <= base.Ui32(v50-v40)) != 0 {
		v206 = l1
		v208 = l3
		v209 = v6
		v211 = v31
		v212 = v33
		v215 = v28 + int32(288)
		v217 = v211 | int64(281474976710656)
		v222 = v217<<(uint(int64(15))%64) | int64(base.Ui64(v208)>>(uint(int64(49))%64))
		v223 = int64(0)
		v225 = int64(8432131802713292800) - v222
		v231 = int64(32)
		v232 = int64(base.Ui64(v225) >> (uint(v231) % 64))
		v234 = int64(base.Ui64(v222) >> (uint(v231) % 64))
		v237 = int64(4294967295)
		v238 = v225 & v237
		v240 = v222 & v237
		v241 = v238 * v240
		v245 = int64(base.Ui64(v241)>>(uint(v231)%64)) + v238*v234
		v252 = v240*v232 + v245&v237
		*(*int64)(unsafe.Add(mBase, uint32(v215)+8)) = v222*v223 + v223*v225 + v232*v234 + int64(base.Ui64(v245)>>(uint(v231)%64)) + int64(base.Ui64(v252)>>(uint(v231)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v215))) = v241&v237 | v252<<(uint(v231)%64)
		v264 = v28 + int32(272)
		v265 = int64(0)
		v266 = *(*int64)(unsafe.Add(mBase, uint32(v28)+296))
		v267 = v265 - v266
		v274 = int64(32)
		v275 = int64(base.Ui64(v225) >> (uint(v274) % 64))
		v277 = int64(base.Ui64(v267) >> (uint(v274) % 64))
		v280 = int64(4294967295)
		v281 = v225 & v280
		v283 = v267 & v280
		v284 = v281 * v283
		v288 = int64(base.Ui64(v284)>>(uint(v274)%64)) + v281*v277
		v295 = v283*v275 + v288&v280
		*(*int64)(unsafe.Add(mBase, uint32(v264)+8)) = v267*v265 + v265*v225 + v275*v277 + int64(base.Ui64(v288)>>(uint(v274)%64)) + int64(base.Ui64(v295)>>(uint(v274)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v264))) = v284&v280 | v295<<(uint(v274)%64)
		v307 = v28 + int32(256)
		v308 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
		v311 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
		v314 = v308<<(uint(int64(1))%64) | int64(base.Ui64(v311)>>(uint(int64(63))%64))
		v315 = int64(0)
		v321 = int64(32)
		v322 = int64(base.Ui64(v222) >> (uint(v321) % 64))
		v324 = int64(base.Ui64(v314) >> (uint(v321) % 64))
		v327 = int64(4294967295)
		v328 = v222 & v327
		v330 = v314 & v327
		v331 = v328 * v330
		v335 = int64(base.Ui64(v331)>>(uint(v321)%64)) + v328*v324
		v342 = v330*v322 + v335&v327
		*(*int64)(unsafe.Add(mBase, uint32(v307)+8)) = v314*v315 + v315*v222 + v322*v324 + int64(base.Ui64(v335)>>(uint(v321)%64)) + int64(base.Ui64(v342)>>(uint(v321)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v307))) = v331&v327 | v342<<(uint(v321)%64)
		v354 = v28 + int32(240)
		v355 = int64(0)
		v357 = *(*int64)(unsafe.Add(mBase, uint32(v28)+264))
		v358 = v355 - v357
		v364 = int64(32)
		v365 = int64(base.Ui64(v358) >> (uint(v364) % 64))
		v367 = int64(base.Ui64(v314) >> (uint(v364) % 64))
		v370 = int64(4294967295)
		v371 = v358 & v370
		v373 = v314 & v370
		v374 = v371 * v373
		v378 = int64(base.Ui64(v374)>>(uint(v364)%64)) + v371*v367
		v385 = v373*v365 + v378&v370
		*(*int64)(unsafe.Add(mBase, uint32(v354)+8)) = v314*v355 + v355*v358 + v365*v367 + int64(base.Ui64(v378)>>(uint(v364)%64)) + int64(base.Ui64(v385)>>(uint(v364)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v354))) = v374&v370 | v385<<(uint(v364)%64)
		v397 = v28 + int32(224)
		v398 = *(*int64)(unsafe.Add(mBase, uint32(v28)+248))
		v401 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
		v404 = v398<<(uint(int64(1))%64) | int64(base.Ui64(v401)>>(uint(int64(63))%64))
		v405 = int64(0)
		v411 = int64(32)
		v412 = int64(base.Ui64(v222) >> (uint(v411) % 64))
		v414 = int64(base.Ui64(v404) >> (uint(v411) % 64))
		v417 = int64(4294967295)
		v418 = v222 & v417
		v420 = v404 & v417
		v421 = v418 * v420
		v425 = int64(base.Ui64(v421)>>(uint(v411)%64)) + v418*v414
		v432 = v420*v412 + v425&v417
		*(*int64)(unsafe.Add(mBase, uint32(v397)+8)) = v404*v405 + v405*v222 + v412*v414 + int64(base.Ui64(v425)>>(uint(v411)%64)) + int64(base.Ui64(v432)>>(uint(v411)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v397))) = v421&v417 | v432<<(uint(v411)%64)
		v444 = v28 + int32(208)
		v445 = int64(0)
		v447 = *(*int64)(unsafe.Add(mBase, uint32(v28)+232))
		v448 = v445 - v447
		v454 = int64(32)
		v455 = int64(base.Ui64(v448) >> (uint(v454) % 64))
		v457 = int64(base.Ui64(v404) >> (uint(v454) % 64))
		v460 = int64(4294967295)
		v461 = v448 & v460
		v463 = v404 & v460
		v464 = v461 * v463
		v468 = int64(base.Ui64(v464)>>(uint(v454)%64)) + v461*v457
		v475 = v463*v455 + v468&v460
		*(*int64)(unsafe.Add(mBase, uint32(v444)+8)) = v404*v445 + v445*v448 + v455*v457 + int64(base.Ui64(v468)>>(uint(v454)%64)) + int64(base.Ui64(v475)>>(uint(v454)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v444))) = v464&v460 | v475<<(uint(v454)%64)
		v487 = v28 + int32(192)
		v488 = *(*int64)(unsafe.Add(mBase, uint32(v28)+216))
		v491 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
		v494 = v488<<(uint(int64(1))%64) | int64(base.Ui64(v491)>>(uint(int64(63))%64))
		v495 = int64(0)
		v501 = int64(32)
		v502 = int64(base.Ui64(v222) >> (uint(v501) % 64))
		v504 = int64(base.Ui64(v494) >> (uint(v501) % 64))
		v507 = int64(4294967295)
		v508 = v222 & v507
		v510 = v494 & v507
		v511 = v508 * v510
		v515 = int64(base.Ui64(v511)>>(uint(v501)%64)) + v508*v504
		v522 = v510*v502 + v515&v507
		*(*int64)(unsafe.Add(mBase, uint32(v487)+8)) = v494*v495 + v495*v222 + v502*v504 + int64(base.Ui64(v515)>>(uint(v501)%64)) + int64(base.Ui64(v522)>>(uint(v501)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v487))) = v511&v507 | v522<<(uint(v501)%64)
		v534 = v28 + int32(176)
		v535 = int64(0)
		v537 = *(*int64)(unsafe.Add(mBase, uint32(v28)+200))
		v538 = v535 - v537
		v544 = int64(32)
		v545 = int64(base.Ui64(v538) >> (uint(v544) % 64))
		v547 = int64(base.Ui64(v494) >> (uint(v544) % 64))
		v550 = int64(4294967295)
		v551 = v538 & v550
		v553 = v494 & v550
		v554 = v551 * v553
		v558 = int64(base.Ui64(v554)>>(uint(v544)%64)) + v551*v547
		v565 = v553*v545 + v558&v550
		*(*int64)(unsafe.Add(mBase, uint32(v534)+8)) = v494*v535 + v535*v538 + v545*v547 + int64(base.Ui64(v558)>>(uint(v544)%64)) + int64(base.Ui64(v565)>>(uint(v544)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v534))) = v554&v550 | v565<<(uint(v544)%64)
		v577 = v28 + int32(160)
		v578 = int64(0)
		v579 = *(*int64)(unsafe.Add(mBase, uint32(v28)+184))
		v580 = int64(1)
		v582 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
		v587 = v579<<(uint(v580)%64) | int64(base.Ui64(v582)>>(uint(int64(63))%64)) - v580
		v593 = int64(32)
		v594 = int64(base.Ui64(v587) >> (uint(v593) % 64))
		v596 = int64(base.Ui64(v222) >> (uint(v593) % 64))
		v599 = int64(4294967295)
		v600 = v587 & v599
		v602 = v222 & v599
		v603 = v600 * v602
		v607 = int64(base.Ui64(v603)>>(uint(v593)%64)) + v600*v596
		v614 = v602*v594 + v607&v599
		*(*int64)(unsafe.Add(mBase, uint32(v577)+8)) = v222*v578 + v578*v587 + v594*v596 + int64(base.Ui64(v607)>>(uint(v593)%64)) + int64(base.Ui64(v614)>>(uint(v593)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v577))) = v603&v599 | v614<<(uint(v593)%64)
		v626 = v28 + int32(144)
		v628 = v208 << (uint(int64(15)) % 64)
		v629 = int64(0)
		v635 = int64(32)
		v636 = int64(base.Ui64(v587) >> (uint(v635) % 64))
		v638 = int64(base.Ui64(v628) >> (uint(v635) % 64))
		v641 = int64(4294967295)
		v642 = v587 & v641
		v644 = v628 & v641
		v645 = v642 * v644
		v649 = int64(base.Ui64(v645)>>(uint(v635)%64)) + v642*v638
		v656 = v644*v636 + v649&v641
		*(*int64)(unsafe.Add(mBase, uint32(v626)+8)) = v628*v629 + v629*v587 + v636*v638 + int64(base.Ui64(v649)>>(uint(v635)%64)) + int64(base.Ui64(v656)>>(uint(v635)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v626))) = v645&v641 | v656<<(uint(v635)%64)
		v668 = v28 + int32(112)
		v669 = int64(0)
		v671 = *(*int64)(unsafe.Add(mBase, uint32(v28)+168))
		v672 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
		v673 = *(*int64)(unsafe.Add(mBase, uint32(v28)+152))
		v674 = v672 + v673
		v682 = v669 - (v671 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v674) < base.Ui64(v672))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(1)) < base.Ui64(v674))))
		v688 = int64(32)
		v689 = int64(base.Ui64(v682) >> (uint(v688) % 64))
		v691 = int64(base.Ui64(v587) >> (uint(v688) % 64))
		v694 = int64(4294967295)
		v695 = v682 & v694
		v697 = v587 & v694
		v698 = v695 * v697
		v702 = int64(base.Ui64(v698)>>(uint(v688)%64)) + v695*v691
		v709 = v697*v689 + v702&v694
		*(*int64)(unsafe.Add(mBase, uint32(v668)+8)) = v587*v669 + v669*v682 + v689*v691 + int64(base.Ui64(v702)>>(uint(v688)%64)) + int64(base.Ui64(v709)>>(uint(v688)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v668))) = v698&v694 | v709<<(uint(v688)%64)
		v721 = v28 + int32(128)
		v723 = int64(1) - v674
		v724 = int64(0)
		v730 = int64(32)
		v731 = int64(base.Ui64(v587) >> (uint(v730) % 64))
		v733 = int64(base.Ui64(v723) >> (uint(v730) % 64))
		v736 = int64(4294967295)
		v737 = v587 & v736
		v739 = v723 & v736
		v740 = v737 * v739
		v744 = int64(base.Ui64(v740)>>(uint(v730)%64)) + v737*v733
		v751 = v739*v731 + v744&v736
		*(*int64)(unsafe.Add(mBase, uint32(v721)+8)) = v723*v724 + v724*v587 + v731*v733 + int64(base.Ui64(v744)>>(uint(v730)%64)) + int64(base.Ui64(v751)>>(uint(v730)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v721))) = v740&v736 | v751<<(uint(v730)%64)
		v763 = v209 + (v50 - v41)
		v765 = v763 + int32(_a_F___divtf3_1)
		v766 = *(*int64)(unsafe.Add(mBase, uint32(v28)+112))
		v767 = int64(1)
		v768 = v766 << (uint(v767) % 64)
		v769 = *(*int64)(unsafe.Add(mBase, uint32(v28)+136))
		v772 = *(*int64)(unsafe.Add(mBase, uint32(v28)+128))
		v773 = int64(63)
		v776 = v768 + (v769<<(uint(v767)%64) | int64(base.Ui64(v772)>>(uint(v773)%64)))
		v778 = v776 - int64(13927)
		v779 = int64(32)
		v780 = int64(base.Ui64(v778) >> (uint(v779) % 64))
		v782 = v212 | int64(281474976710656)
		v784 = v782 << (uint(v767) % 64)
		v786 = int64(base.Ui64(v784) >> (uint(v779) % 64))
		v787 = v780 * v786
		v789 = v206 << (uint(v767) % 64)
		v791 = int64(base.Ui64(v789) >> (uint(v779) % 64))
		v796 = *(*int64)(unsafe.Add(mBase, uint32(v28)+120))
		v808 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v778) < base.Ui64(v776))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v776) < base.Ui64(v768))) + (v796<<(uint(v767)%64) | int64(base.Ui64(v766)>>(uint(v773)%64)) + int64(base.Ui64(v769)>>(uint(v773)%64)))) - v767
		v810 = int64(base.Ui64(v808) >> (uint(v779) % 64))
		v812 = v787 + v791*v810
		v815 = int64(4294967295)
		v816 = v808 & v815
		v818 = int64(base.Ui64(v206) >> (uint(v773) % 64))
		v823 = (v818 | v212<<(uint(v767)%64)) & v815
		v825 = v812 + v816*v823
		v831 = v786 * v816
		v833 = v831 + v823*v810
		v844 = v825 + v833<<(uint(v779)%64)
		v849 = v778 & v815
		v850 = v849 * v823
		v852 = v850 + v780*v791
		v856 = v789 & int64(4294967294)
		v858 = v852 + v816*v856
		v862 = v844 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v852) < base.Ui64(v850))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v858) < base.Ui64(v852))))
		v866 = v786 * v849
		v868 = v866 + v856*v810
		v870 = v868 + v780*v823
		v872 = v870 + v791*v816
		v886 = v862 + (int64(base.Ui64(v872)>>(uint(v779)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v872) < base.Ui64(v870)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v868) < base.Ui64(v866)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v870) < base.Ui64(v868)))))<<(uint(v779)%64))
		v890 = v780 * v856
		v892 = v890 + v791*v849
		v900 = v858 + (int64(base.Ui64(v892)>>(uint(v779)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v892) < base.Ui64(v890)))<<(uint(v779)%64))
		v905 = v900 + v872<<(uint(v779)%64)
		v909 = v886 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v900) < base.Ui64(v858))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v905) < base.Ui64(v900))))
		v914 = v892 << (uint(v779) % 64)
		v920 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v914+v856*v849) < base.Ui64(v914))) ^ int64(-1)
		v925 = v909 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v920) < base.Ui64(v905))&base.B2i32(v920 != v905))
		v928 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v812) < base.Ui64(v787))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v825) < base.Ui64(v812))) + v786*v810 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v833) < base.Ui64(v831)))<<(uint(v779)%64) | int64(base.Ui64(v833)>>(uint(v779)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v844) < base.Ui64(v825))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v862) < base.Ui64(v844))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v886) < base.Ui64(v862))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v909) < base.Ui64(v886))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v925) < base.Ui64(v909)))
		if base.Ui64(v928) <= base.Ui64(int64(562949953421311)) {
			v933 = v28 + int32(80)
			v935 = base.B2i32(base.Ui64(v928) < base.Ui64(int64(281474976710656)))
			v936 = base.I64_extend_i32_u(v935)
			v937 = v925 << (uint(v936) % 64)
			v945 = v928<<(uint(v936)%64) | int64(base.Ui64(int64(base.Ui64(v925)>>(uint(int64(1))%64)))>>(uint(base.I64_extend_i32_u(v935^int32(63)))%64))
			v950 = int64(32)
			v951 = int64(base.Ui64(v208) >> (uint(v950) % 64))
			v953 = int64(base.Ui64(v937) >> (uint(v950) % 64))
			v956 = int64(4294967295)
			v957 = v208 & v956
			v959 = v937 & v956
			v960 = v957 * v959
			v964 = int64(base.Ui64(v960)>>(uint(v950)%64)) + v957*v953
			v971 = v959*v951 + v964&v956
			*(*int64)(unsafe.Add(mBase, uint32(v933)+8)) = v937*v217 + v945*v208 + v951*v953 + int64(base.Ui64(v964)>>(uint(v950)%64)) + int64(base.Ui64(v971)>>(uint(v950)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v933))) = v960&v956 | v971<<(uint(v950)%64)
			if base.Ui64(v928) < base.Ui64(int64(281474976710656)) {
				v984 = v763 + int32(_a_F___divtf3_2)
			} else {
				v984 = v765
			}
			v989 = *(*int64)(unsafe.Add(mBase, uint32(v28)+88))
			v991 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
			v992 = int64(0)
			v1055 = v789
			v1056 = v945
			v1057 = v984 - int32(1)
			v1059 = v206<<(uint(int64(49))%64) - v989 - base.I64_extend_i32_u(base.B2i32(v991 != v992))
			v1060 = v937
			v1062 = v784 | v818
			v1063 = v992 - v991
		} else {
			v999 = v28 + int32(96)
			v1002 = int64(1)
			v1004 = v928<<(uint(int64(63))%64) | int64(base.Ui64(v925)>>(uint(v1002)%64))
			v1006 = int64(base.Ui64(v928) >> (uint(v1002) % 64))
			v1011 = int64(32)
			v1012 = int64(base.Ui64(v208) >> (uint(v1011) % 64))
			v1014 = int64(base.Ui64(v1004) >> (uint(v1011) % 64))
			v1017 = int64(4294967295)
			v1018 = v208 & v1017
			v1020 = v1004 & v1017
			v1021 = v1018 * v1020
			v1025 = int64(base.Ui64(v1021)>>(uint(v1011)%64)) + v1018*v1014
			v1032 = v1020*v1012 + v1025&v1017
			*(*int64)(unsafe.Add(mBase, uint32(v999)+8)) = v1004*v217 + v1006*v208 + v1012*v1014 + int64(base.Ui64(v1025)>>(uint(v1011)%64)) + int64(base.Ui64(v1032)>>(uint(v1011)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v999))) = v1021&v1017 | v1032<<(uint(v1011)%64)
			v1045 = *(*int64)(unsafe.Add(mBase, uint32(v28)+104))
			v1047 = *(*int64)(unsafe.Add(mBase, uint32(v28)+96))
			v1048 = int64(0)
			v1055 = v206
			v1056 = v1006
			v1057 = v765
			v1059 = v206<<(uint(int64(48))%64) - v1045 - base.I64_extend_i32_u(base.B2i32(v1047 != v1048))
			v1060 = v1004
			v1062 = v782
			v1063 = v1048 - v1047
		}
		if int32(_a_F___divtf3_0) <= v1057 {
			v1325 = int64(0)
			v1335 = v36 | int64(9223090561878065152)
		} else {
			if int32(0) < v1057 {
				v1071 = int64(1)
				v1197 = v1059<<(uint(v1071)%64) | int64(base.Ui64(v1063)>>(uint(int64(63))%64))
				v1200 = v1056&int64(281474976710655) | base.I64_extend_i32_u(v1057)<<(uint(int64(48))%64)
				v1201 = v1060
				v1202 = v1063 << (uint(v1071) % 64)
				v1204 = v28 + int32(16)
				v1205 = int64(3)
				v1206 = int64(0)
				v1211 = int64(32)
				v1214 = int64(base.Ui64(v208) >> (uint(v1211) % 64))
				v1217 = int64(4294967295)
				v1220 = v208 & v1217
				v1221 = v1205 * v1220
				v1225 = int64(base.Ui64(v1221)>>(uint(v1211)%64)) + v1205*v1214
				v1232 = v1220*v1206 + v1225&v1217
				*(*int64)(unsafe.Add(mBase, uint32(v1204)+8)) = v208*v1206 + v217*v1205 + v1206*v1214 + int64(base.Ui64(v1225)>>(uint(v1211)%64)) + int64(base.Ui64(v1232)>>(uint(v1211)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v1204))) = v1221&v1217 | v1232<<(uint(v1211)%64)
				v1243 = int64(5)
				v1244 = int64(0)
				v1249 = int64(32)
				v1252 = int64(base.Ui64(v208) >> (uint(v1249) % 64))
				v1255 = int64(4294967295)
				v1258 = v208 & v1255
				v1259 = v1243 * v1258
				v1263 = int64(base.Ui64(v1259)>>(uint(v1249)%64)) + v1243*v1252
				v1270 = v1258*v1244 + v1263&v1255
				*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v208*v1244 + v217*v1243 + v1244*v1252 + int64(base.Ui64(v1263)>>(uint(v1249)%64)) + int64(base.Ui64(v1270)>>(uint(v1249)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1259&v1255 | v1270<<(uint(v1249)%64)
				v1282 = v1201 & int64(1)
				v1283 = v1282 + v1202
				v1287 = v1197 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1283) < base.Ui64(v1282)))
				if v1287 == v217 {
					v1290 = base.B2i32(base.Ui64(v208) < base.Ui64(v1283))
				} else {
					v1290 = base.B2i32(base.Ui64(v217) < base.Ui64(v1287))
				}
				v1292 = v1201 + base.I64_extend_i32_u(v1290)
				v1295 = v1200 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1292) < base.Ui64(v1201)))
				v1298 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
				v1300 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
				if v1287 == v1300 {
					v1303 = base.B2i32(base.Ui64(v1298) < base.Ui64(v1283))
				} else {
					v1303 = base.B2i32(base.Ui64(v1300) < base.Ui64(v1287))
				}
				v1306 = v1292 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1295) < base.Ui64(int64(9223090561878065152)))&v1303)
				v1309 = v1295 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1306) < base.Ui64(v1292)))
				v1312 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
				v1314 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
				if v1287 == v1314 {
					v1317 = base.B2i32(base.Ui64(v1312) < base.Ui64(v1283))
				} else {
					v1317 = base.B2i32(base.Ui64(v1314) < base.Ui64(v1287))
				}
				v1320 = v1306 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1309) < base.Ui64(int64(9223090561878065152)))&v1317)
				v1325 = v1320
				v1335 = v1309 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1320) < base.Ui64(v1306))) | v36
			} else {
				if v1057 <= int32(-113) {
					v1325 = int64(0)
					v1335 = v36
				} else {
					v1088 = v28 - int32(-64)
					v1090 = int32(1) - v1057
					if v1090&int32(64) != 0 {
						v1109 = int64(base.Ui64(v1056) >> (uint(base.I64_extend_i32_u(v1090+int32(-64))) % 64))
						v1110 = int64(0)
					} else {
						if v1090 == int32(0) {
							v1109 = v1060
							v1110 = v1056
						} else {
							v1105 = base.I64_extend_i32_u(v1090)
							v1109 = v1056<<(uint(base.I64_extend_i32_u(int32(64)-v1090))%64) | int64(base.Ui64(v1060)>>(uint(v1105)%64))
							v1110 = int64(base.Ui64(v1056) >> (uint(v1105) % 64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v1088))) = v1109
					*(*int64)(unsafe.Add(mBase, uint32(v1088)+8)) = v1110
					v1115 = v28 + int32(48)
					v1117 = v1057 + int32(112)
					if v1117&int32(64) != 0 {
						v1136 = int64(0)
						v1137 = v1055 << (uint(base.I64_extend_i32_u(v1057+int32(48))) % 64)
					} else {
						if v1117 == int32(0) {
							v1136 = v1055
							v1137 = v1062
						} else {
							v1128 = base.I64_extend_i32_u(v1117)
							v1136 = v1055 << (uint(v1128) % 64)
							v1137 = v1062<<(uint(v1128)%64) | int64(base.Ui64(v1055)>>(uint(base.I64_extend_i32_u(int32(64)-v1117))%64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v1115))) = v1136
					*(*int64)(unsafe.Add(mBase, uint32(v1115)+8)) = v1137
					v1142 = v28 + int32(32)
					v1143 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
					v1144 = *(*int64)(unsafe.Add(mBase, uint32(v28)+72))
					v1149 = int64(32)
					v1150 = int64(base.Ui64(v1143) >> (uint(v1149) % 64))
					v1152 = int64(base.Ui64(v208) >> (uint(v1149) % 64))
					v1155 = int64(4294967295)
					v1156 = v1143 & v1155
					v1158 = v208 & v1155
					v1159 = v1156 * v1158
					v1163 = int64(base.Ui64(v1159)>>(uint(v1149)%64)) + v1156*v1152
					v1170 = v1158*v1150 + v1163&v1155
					*(*int64)(unsafe.Add(mBase, uint32(v1142)+8)) = v208*v1144 + v217*v1143 + v1150*v1152 + int64(base.Ui64(v1163)>>(uint(v1149)%64)) + int64(base.Ui64(v1170)>>(uint(v1149)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v1142))) = v1159&v1155 | v1170<<(uint(v1149)%64)
					v1181 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
					v1182 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
					v1183 = int64(1)
					v1185 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
					v1190 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
					v1192 = v1185 << (uint(v1183) % 64)
					v1197 = v1181 - (v1182<<(uint(v1183)%64) | int64(base.Ui64(v1185)>>(uint(int64(63))%64))) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1190) < base.Ui64(v1192)))
					v1200 = v1144
					v1201 = v1143
					v1202 = v1190 - v1192
					v1204 = v28 + int32(16)
					v1205 = int64(3)
					v1206 = int64(0)
					v1211 = int64(32)
					v1214 = int64(base.Ui64(v208) >> (uint(v1211) % 64))
					v1217 = int64(4294967295)
					v1220 = v208 & v1217
					v1221 = v1205 * v1220
					v1225 = int64(base.Ui64(v1221)>>(uint(v1211)%64)) + v1205*v1214
					v1232 = v1220*v1206 + v1225&v1217
					*(*int64)(unsafe.Add(mBase, uint32(v1204)+8)) = v208*v1206 + v217*v1205 + v1206*v1214 + int64(base.Ui64(v1225)>>(uint(v1211)%64)) + int64(base.Ui64(v1232)>>(uint(v1211)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v1204))) = v1221&v1217 | v1232<<(uint(v1211)%64)
					v1243 = int64(5)
					v1244 = int64(0)
					v1249 = int64(32)
					v1252 = int64(base.Ui64(v208) >> (uint(v1249) % 64))
					v1255 = int64(4294967295)
					v1258 = v208 & v1255
					v1259 = v1243 * v1258
					v1263 = int64(base.Ui64(v1259)>>(uint(v1249)%64)) + v1243*v1252
					v1270 = v1258*v1244 + v1263&v1255
					*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v208*v1244 + v217*v1243 + v1244*v1252 + int64(base.Ui64(v1263)>>(uint(v1249)%64)) + int64(base.Ui64(v1270)>>(uint(v1249)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1259&v1255 | v1270<<(uint(v1249)%64)
					v1282 = v1201 & int64(1)
					v1283 = v1282 + v1202
					v1287 = v1197 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1283) < base.Ui64(v1282)))
					if v1287 == v217 {
						v1290 = base.B2i32(base.Ui64(v208) < base.Ui64(v1283))
					} else {
						v1290 = base.B2i32(base.Ui64(v217) < base.Ui64(v1287))
					}
					v1292 = v1201 + base.I64_extend_i32_u(v1290)
					v1295 = v1200 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1292) < base.Ui64(v1201)))
					v1298 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
					v1300 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
					if v1287 == v1300 {
						v1303 = base.B2i32(base.Ui64(v1298) < base.Ui64(v1283))
					} else {
						v1303 = base.B2i32(base.Ui64(v1300) < base.Ui64(v1287))
					}
					v1306 = v1292 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1295) < base.Ui64(int64(9223090561878065152)))&v1303)
					v1309 = v1295 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1306) < base.Ui64(v1292)))
					v1312 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
					v1314 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
					if v1287 == v1314 {
						v1317 = base.B2i32(base.Ui64(v1312) < base.Ui64(v1283))
					} else {
						v1317 = base.B2i32(base.Ui64(v1314) < base.Ui64(v1287))
					}
					v1320 = v1306 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1309) < base.Ui64(int64(9223090561878065152)))&v1317)
					v1325 = v1320
					v1335 = v1309 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1320) < base.Ui64(v1306))) | v36
				}
			}
		}
	} else {
		v59 = l2 & int64(9223372036854775807)
		v60 = int64(9223090561878065152)
		if v59 == v60 {
			v64 = base.B2i32(l1 == int64(0))
		} else {
			v64 = base.B2i32(base.Ui64(v59) < base.Ui64(v60))
		}
		if v64 == int32(0) {
			v1325 = l1
			v1335 = l2 | int64(140737488355328)
		} else {
			v72 = l4 & int64(9223372036854775807)
			v73 = int64(9223090561878065152)
			if v72 == v73 {
				v77 = base.B2i32(l3 == int64(0))
			} else {
				v77 = base.B2i32(base.Ui64(v72) < base.Ui64(v73))
			}
			if v77 == int32(0) {
				v1325 = l3
				v1335 = l4 | int64(140737488355328)
			} else {
				if l1|(v59^int64(9223090561878065152)) == int64(0) {
					if l3|(v72^int64(9223090561878065152)) == int64(0) {
						v1325 = int64(0)
						v1335 = int64(9223231299366420480)
					} else {
						v1325 = int64(0)
						v1335 = v36 | int64(9223090561878065152)
					}
				} else {
					if l3|(v72^int64(9223090561878065152)) == int64(0) {
						v1325 = int64(0)
						v1335 = v36
					} else {
						if l1|v59 == int64(0) {
							if v72|l3 == int64(0) {
								v110 = int64(9223231299366420480)
							} else {
								v110 = v36
							}
							v1325 = int64(0)
							v1335 = v110
						} else {
							if v72|l3 == int64(0) {
								v1325 = int64(0)
								v1335 = v36 | int64(9223090561878065152)
							} else {
								if base.Ui64(v59) <= base.Ui64(int64(281474976710655)) {
									v121 = v28 + int32(320)
									v123 = base.B2i32(v33 == int64(0))
									if v33 == int64(0) {
										v124 = l1
									} else {
										v124 = v33
									}
									if v33 == int64(0) {
										v128 = int64(64)
									} else {
										v128 = int64(0)
									}
									v130 = base.I32_wrap_i64(base.I64_clz(v124) + v128)
									v132 = v130 - int32(15)
									if v132&int32(64) != 0 {
										v151 = int64(0)
										v152 = l1 << (uint(base.I64_extend_i32_u(v132+int32(-64))) % 64)
									} else {
										if v132 == int32(0) {
											v151 = l1
											v152 = v33
										} else {
											v143 = base.I64_extend_i32_u(v132)
											v151 = l1 << (uint(v143) % 64)
											v152 = v33<<(uint(v143)%64) | int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v132))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v121))) = v151
									*(*int64)(unsafe.Add(mBase, uint32(v121)+8)) = v152
									v158 = *(*int64)(unsafe.Add(mBase, uint32(v28)+328))
									v159 = *(*int64)(unsafe.Add(mBase, uint32(v28)+320))
									v160 = v159
									v161 = int32(16) - v130
									v162 = v158
								} else {
									v160 = l1
									v161 = v6
									v162 = v33
								}
								if base.Ui64(int64(281474976710655)) < base.Ui64(v72) {
									v206 = v160
									v208 = l3
									v209 = v161
									v211 = v31
									v212 = v162
								} else {
									v166 = v28 + int32(304)
									v168 = base.B2i32(v31 == int64(0))
									if v31 == int64(0) {
										v169 = l3
									} else {
										v169 = v31
									}
									if v31 == int64(0) {
										v173 = int64(64)
									} else {
										v173 = int64(0)
									}
									v175 = base.I32_wrap_i64(base.I64_clz(v169) + v173)
									v177 = v175 - int32(15)
									if v177&int32(64) != 0 {
										v196 = int64(0)
										v197 = l3 << (uint(base.I64_extend_i32_u(v177+int32(-64))) % 64)
									} else {
										if v177 == int32(0) {
											v196 = l3
											v197 = v31
										} else {
											v188 = base.I64_extend_i32_u(v177)
											v196 = l3 << (uint(v188) % 64)
											v197 = v31<<(uint(v188)%64) | int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v177))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v166))) = v196
									*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v197
									v204 = *(*int64)(unsafe.Add(mBase, uint32(v28)+312))
									v205 = *(*int64)(unsafe.Add(mBase, uint32(v28)+304))
									v206 = v160
									v208 = v205
									v209 = v161 + v175 - int32(16)
									v211 = v204
									v212 = v162
								}
								v215 = v28 + int32(288)
								v217 = v211 | int64(281474976710656)
								v222 = v217<<(uint(int64(15))%64) | int64(base.Ui64(v208)>>(uint(int64(49))%64))
								v223 = int64(0)
								v225 = int64(8432131802713292800) - v222
								v231 = int64(32)
								v232 = int64(base.Ui64(v225) >> (uint(v231) % 64))
								v234 = int64(base.Ui64(v222) >> (uint(v231) % 64))
								v237 = int64(4294967295)
								v238 = v225 & v237
								v240 = v222 & v237
								v241 = v238 * v240
								v245 = int64(base.Ui64(v241)>>(uint(v231)%64)) + v238*v234
								v252 = v240*v232 + v245&v237
								*(*int64)(unsafe.Add(mBase, uint32(v215)+8)) = v222*v223 + v223*v225 + v232*v234 + int64(base.Ui64(v245)>>(uint(v231)%64)) + int64(base.Ui64(v252)>>(uint(v231)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v215))) = v241&v237 | v252<<(uint(v231)%64)
								v264 = v28 + int32(272)
								v265 = int64(0)
								v266 = *(*int64)(unsafe.Add(mBase, uint32(v28)+296))
								v267 = v265 - v266
								v274 = int64(32)
								v275 = int64(base.Ui64(v225) >> (uint(v274) % 64))
								v277 = int64(base.Ui64(v267) >> (uint(v274) % 64))
								v280 = int64(4294967295)
								v281 = v225 & v280
								v283 = v267 & v280
								v284 = v281 * v283
								v288 = int64(base.Ui64(v284)>>(uint(v274)%64)) + v281*v277
								v295 = v283*v275 + v288&v280
								*(*int64)(unsafe.Add(mBase, uint32(v264)+8)) = v267*v265 + v265*v225 + v275*v277 + int64(base.Ui64(v288)>>(uint(v274)%64)) + int64(base.Ui64(v295)>>(uint(v274)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v264))) = v284&v280 | v295<<(uint(v274)%64)
								v307 = v28 + int32(256)
								v308 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
								v311 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
								v314 = v308<<(uint(int64(1))%64) | int64(base.Ui64(v311)>>(uint(int64(63))%64))
								v315 = int64(0)
								v321 = int64(32)
								v322 = int64(base.Ui64(v222) >> (uint(v321) % 64))
								v324 = int64(base.Ui64(v314) >> (uint(v321) % 64))
								v327 = int64(4294967295)
								v328 = v222 & v327
								v330 = v314 & v327
								v331 = v328 * v330
								v335 = int64(base.Ui64(v331)>>(uint(v321)%64)) + v328*v324
								v342 = v330*v322 + v335&v327
								*(*int64)(unsafe.Add(mBase, uint32(v307)+8)) = v314*v315 + v315*v222 + v322*v324 + int64(base.Ui64(v335)>>(uint(v321)%64)) + int64(base.Ui64(v342)>>(uint(v321)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v307))) = v331&v327 | v342<<(uint(v321)%64)
								v354 = v28 + int32(240)
								v355 = int64(0)
								v357 = *(*int64)(unsafe.Add(mBase, uint32(v28)+264))
								v358 = v355 - v357
								v364 = int64(32)
								v365 = int64(base.Ui64(v358) >> (uint(v364) % 64))
								v367 = int64(base.Ui64(v314) >> (uint(v364) % 64))
								v370 = int64(4294967295)
								v371 = v358 & v370
								v373 = v314 & v370
								v374 = v371 * v373
								v378 = int64(base.Ui64(v374)>>(uint(v364)%64)) + v371*v367
								v385 = v373*v365 + v378&v370
								*(*int64)(unsafe.Add(mBase, uint32(v354)+8)) = v314*v355 + v355*v358 + v365*v367 + int64(base.Ui64(v378)>>(uint(v364)%64)) + int64(base.Ui64(v385)>>(uint(v364)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v354))) = v374&v370 | v385<<(uint(v364)%64)
								v397 = v28 + int32(224)
								v398 = *(*int64)(unsafe.Add(mBase, uint32(v28)+248))
								v401 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
								v404 = v398<<(uint(int64(1))%64) | int64(base.Ui64(v401)>>(uint(int64(63))%64))
								v405 = int64(0)
								v411 = int64(32)
								v412 = int64(base.Ui64(v222) >> (uint(v411) % 64))
								v414 = int64(base.Ui64(v404) >> (uint(v411) % 64))
								v417 = int64(4294967295)
								v418 = v222 & v417
								v420 = v404 & v417
								v421 = v418 * v420
								v425 = int64(base.Ui64(v421)>>(uint(v411)%64)) + v418*v414
								v432 = v420*v412 + v425&v417
								*(*int64)(unsafe.Add(mBase, uint32(v397)+8)) = v404*v405 + v405*v222 + v412*v414 + int64(base.Ui64(v425)>>(uint(v411)%64)) + int64(base.Ui64(v432)>>(uint(v411)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v397))) = v421&v417 | v432<<(uint(v411)%64)
								v444 = v28 + int32(208)
								v445 = int64(0)
								v447 = *(*int64)(unsafe.Add(mBase, uint32(v28)+232))
								v448 = v445 - v447
								v454 = int64(32)
								v455 = int64(base.Ui64(v448) >> (uint(v454) % 64))
								v457 = int64(base.Ui64(v404) >> (uint(v454) % 64))
								v460 = int64(4294967295)
								v461 = v448 & v460
								v463 = v404 & v460
								v464 = v461 * v463
								v468 = int64(base.Ui64(v464)>>(uint(v454)%64)) + v461*v457
								v475 = v463*v455 + v468&v460
								*(*int64)(unsafe.Add(mBase, uint32(v444)+8)) = v404*v445 + v445*v448 + v455*v457 + int64(base.Ui64(v468)>>(uint(v454)%64)) + int64(base.Ui64(v475)>>(uint(v454)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v444))) = v464&v460 | v475<<(uint(v454)%64)
								v487 = v28 + int32(192)
								v488 = *(*int64)(unsafe.Add(mBase, uint32(v28)+216))
								v491 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
								v494 = v488<<(uint(int64(1))%64) | int64(base.Ui64(v491)>>(uint(int64(63))%64))
								v495 = int64(0)
								v501 = int64(32)
								v502 = int64(base.Ui64(v222) >> (uint(v501) % 64))
								v504 = int64(base.Ui64(v494) >> (uint(v501) % 64))
								v507 = int64(4294967295)
								v508 = v222 & v507
								v510 = v494 & v507
								v511 = v508 * v510
								v515 = int64(base.Ui64(v511)>>(uint(v501)%64)) + v508*v504
								v522 = v510*v502 + v515&v507
								*(*int64)(unsafe.Add(mBase, uint32(v487)+8)) = v494*v495 + v495*v222 + v502*v504 + int64(base.Ui64(v515)>>(uint(v501)%64)) + int64(base.Ui64(v522)>>(uint(v501)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v487))) = v511&v507 | v522<<(uint(v501)%64)
								v534 = v28 + int32(176)
								v535 = int64(0)
								v537 = *(*int64)(unsafe.Add(mBase, uint32(v28)+200))
								v538 = v535 - v537
								v544 = int64(32)
								v545 = int64(base.Ui64(v538) >> (uint(v544) % 64))
								v547 = int64(base.Ui64(v494) >> (uint(v544) % 64))
								v550 = int64(4294967295)
								v551 = v538 & v550
								v553 = v494 & v550
								v554 = v551 * v553
								v558 = int64(base.Ui64(v554)>>(uint(v544)%64)) + v551*v547
								v565 = v553*v545 + v558&v550
								*(*int64)(unsafe.Add(mBase, uint32(v534)+8)) = v494*v535 + v535*v538 + v545*v547 + int64(base.Ui64(v558)>>(uint(v544)%64)) + int64(base.Ui64(v565)>>(uint(v544)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v534))) = v554&v550 | v565<<(uint(v544)%64)
								v577 = v28 + int32(160)
								v578 = int64(0)
								v579 = *(*int64)(unsafe.Add(mBase, uint32(v28)+184))
								v580 = int64(1)
								v582 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
								v587 = v579<<(uint(v580)%64) | int64(base.Ui64(v582)>>(uint(int64(63))%64)) - v580
								v593 = int64(32)
								v594 = int64(base.Ui64(v587) >> (uint(v593) % 64))
								v596 = int64(base.Ui64(v222) >> (uint(v593) % 64))
								v599 = int64(4294967295)
								v600 = v587 & v599
								v602 = v222 & v599
								v603 = v600 * v602
								v607 = int64(base.Ui64(v603)>>(uint(v593)%64)) + v600*v596
								v614 = v602*v594 + v607&v599
								*(*int64)(unsafe.Add(mBase, uint32(v577)+8)) = v222*v578 + v578*v587 + v594*v596 + int64(base.Ui64(v607)>>(uint(v593)%64)) + int64(base.Ui64(v614)>>(uint(v593)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v577))) = v603&v599 | v614<<(uint(v593)%64)
								v626 = v28 + int32(144)
								v628 = v208 << (uint(int64(15)) % 64)
								v629 = int64(0)
								v635 = int64(32)
								v636 = int64(base.Ui64(v587) >> (uint(v635) % 64))
								v638 = int64(base.Ui64(v628) >> (uint(v635) % 64))
								v641 = int64(4294967295)
								v642 = v587 & v641
								v644 = v628 & v641
								v645 = v642 * v644
								v649 = int64(base.Ui64(v645)>>(uint(v635)%64)) + v642*v638
								v656 = v644*v636 + v649&v641
								*(*int64)(unsafe.Add(mBase, uint32(v626)+8)) = v628*v629 + v629*v587 + v636*v638 + int64(base.Ui64(v649)>>(uint(v635)%64)) + int64(base.Ui64(v656)>>(uint(v635)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v626))) = v645&v641 | v656<<(uint(v635)%64)
								v668 = v28 + int32(112)
								v669 = int64(0)
								v671 = *(*int64)(unsafe.Add(mBase, uint32(v28)+168))
								v672 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
								v673 = *(*int64)(unsafe.Add(mBase, uint32(v28)+152))
								v674 = v672 + v673
								v682 = v669 - (v671 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v674) < base.Ui64(v672))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(1)) < base.Ui64(v674))))
								v688 = int64(32)
								v689 = int64(base.Ui64(v682) >> (uint(v688) % 64))
								v691 = int64(base.Ui64(v587) >> (uint(v688) % 64))
								v694 = int64(4294967295)
								v695 = v682 & v694
								v697 = v587 & v694
								v698 = v695 * v697
								v702 = int64(base.Ui64(v698)>>(uint(v688)%64)) + v695*v691
								v709 = v697*v689 + v702&v694
								*(*int64)(unsafe.Add(mBase, uint32(v668)+8)) = v587*v669 + v669*v682 + v689*v691 + int64(base.Ui64(v702)>>(uint(v688)%64)) + int64(base.Ui64(v709)>>(uint(v688)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v668))) = v698&v694 | v709<<(uint(v688)%64)
								v721 = v28 + int32(128)
								v723 = int64(1) - v674
								v724 = int64(0)
								v730 = int64(32)
								v731 = int64(base.Ui64(v587) >> (uint(v730) % 64))
								v733 = int64(base.Ui64(v723) >> (uint(v730) % 64))
								v736 = int64(4294967295)
								v737 = v587 & v736
								v739 = v723 & v736
								v740 = v737 * v739
								v744 = int64(base.Ui64(v740)>>(uint(v730)%64)) + v737*v733
								v751 = v739*v731 + v744&v736
								*(*int64)(unsafe.Add(mBase, uint32(v721)+8)) = v723*v724 + v724*v587 + v731*v733 + int64(base.Ui64(v744)>>(uint(v730)%64)) + int64(base.Ui64(v751)>>(uint(v730)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v721))) = v740&v736 | v751<<(uint(v730)%64)
								v763 = v209 + (v50 - v41)
								v765 = v763 + int32(_a_F___divtf3_1)
								v766 = *(*int64)(unsafe.Add(mBase, uint32(v28)+112))
								v767 = int64(1)
								v768 = v766 << (uint(v767) % 64)
								v769 = *(*int64)(unsafe.Add(mBase, uint32(v28)+136))
								v772 = *(*int64)(unsafe.Add(mBase, uint32(v28)+128))
								v773 = int64(63)
								v776 = v768 + (v769<<(uint(v767)%64) | int64(base.Ui64(v772)>>(uint(v773)%64)))
								v778 = v776 - int64(13927)
								v779 = int64(32)
								v780 = int64(base.Ui64(v778) >> (uint(v779) % 64))
								v782 = v212 | int64(281474976710656)
								v784 = v782 << (uint(v767) % 64)
								v786 = int64(base.Ui64(v784) >> (uint(v779) % 64))
								v787 = v780 * v786
								v789 = v206 << (uint(v767) % 64)
								v791 = int64(base.Ui64(v789) >> (uint(v779) % 64))
								v796 = *(*int64)(unsafe.Add(mBase, uint32(v28)+120))
								v808 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v778) < base.Ui64(v776))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v776) < base.Ui64(v768))) + (v796<<(uint(v767)%64) | int64(base.Ui64(v766)>>(uint(v773)%64)) + int64(base.Ui64(v769)>>(uint(v773)%64)))) - v767
								v810 = int64(base.Ui64(v808) >> (uint(v779) % 64))
								v812 = v787 + v791*v810
								v815 = int64(4294967295)
								v816 = v808 & v815
								v818 = int64(base.Ui64(v206) >> (uint(v773) % 64))
								v823 = (v818 | v212<<(uint(v767)%64)) & v815
								v825 = v812 + v816*v823
								v831 = v786 * v816
								v833 = v831 + v823*v810
								v844 = v825 + v833<<(uint(v779)%64)
								v849 = v778 & v815
								v850 = v849 * v823
								v852 = v850 + v780*v791
								v856 = v789 & int64(4294967294)
								v858 = v852 + v816*v856
								v862 = v844 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v852) < base.Ui64(v850))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v858) < base.Ui64(v852))))
								v866 = v786 * v849
								v868 = v866 + v856*v810
								v870 = v868 + v780*v823
								v872 = v870 + v791*v816
								v886 = v862 + (int64(base.Ui64(v872)>>(uint(v779)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v872) < base.Ui64(v870)))+(base.I64_extend_i32_u(base.B2i32(base.Ui64(v868) < base.Ui64(v866)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v870) < base.Ui64(v868)))))<<(uint(v779)%64))
								v890 = v780 * v856
								v892 = v890 + v791*v849
								v900 = v858 + (int64(base.Ui64(v892)>>(uint(v779)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v892) < base.Ui64(v890)))<<(uint(v779)%64))
								v905 = v900 + v872<<(uint(v779)%64)
								v909 = v886 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v900) < base.Ui64(v858))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v905) < base.Ui64(v900))))
								v914 = v892 << (uint(v779) % 64)
								v920 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v914+v856*v849) < base.Ui64(v914))) ^ int64(-1)
								v925 = v909 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v920) < base.Ui64(v905))&base.B2i32(v920 != v905))
								v928 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v812) < base.Ui64(v787))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v825) < base.Ui64(v812))) + v786*v810 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v833) < base.Ui64(v831)))<<(uint(v779)%64) | int64(base.Ui64(v833)>>(uint(v779)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v844) < base.Ui64(v825))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v862) < base.Ui64(v844))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v886) < base.Ui64(v862))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v909) < base.Ui64(v886))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v925) < base.Ui64(v909)))
								if base.Ui64(v928) <= base.Ui64(int64(562949953421311)) {
									v933 = v28 + int32(80)
									v935 = base.B2i32(base.Ui64(v928) < base.Ui64(int64(281474976710656)))
									v936 = base.I64_extend_i32_u(v935)
									v937 = v925 << (uint(v936) % 64)
									v945 = v928<<(uint(v936)%64) | int64(base.Ui64(int64(base.Ui64(v925)>>(uint(int64(1))%64)))>>(uint(base.I64_extend_i32_u(v935^int32(63)))%64))
									v950 = int64(32)
									v951 = int64(base.Ui64(v208) >> (uint(v950) % 64))
									v953 = int64(base.Ui64(v937) >> (uint(v950) % 64))
									v956 = int64(4294967295)
									v957 = v208 & v956
									v959 = v937 & v956
									v960 = v957 * v959
									v964 = int64(base.Ui64(v960)>>(uint(v950)%64)) + v957*v953
									v971 = v959*v951 + v964&v956
									*(*int64)(unsafe.Add(mBase, uint32(v933)+8)) = v937*v217 + v945*v208 + v951*v953 + int64(base.Ui64(v964)>>(uint(v950)%64)) + int64(base.Ui64(v971)>>(uint(v950)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v933))) = v960&v956 | v971<<(uint(v950)%64)
									if base.Ui64(v928) < base.Ui64(int64(281474976710656)) {
										v984 = v763 + int32(_a_F___divtf3_2)
									} else {
										v984 = v765
									}
									v989 = *(*int64)(unsafe.Add(mBase, uint32(v28)+88))
									v991 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
									v992 = int64(0)
									v1055 = v789
									v1056 = v945
									v1057 = v984 - int32(1)
									v1059 = v206<<(uint(int64(49))%64) - v989 - base.I64_extend_i32_u(base.B2i32(v991 != v992))
									v1060 = v937
									v1062 = v784 | v818
									v1063 = v992 - v991
								} else {
									v999 = v28 + int32(96)
									v1002 = int64(1)
									v1004 = v928<<(uint(int64(63))%64) | int64(base.Ui64(v925)>>(uint(v1002)%64))
									v1006 = int64(base.Ui64(v928) >> (uint(v1002) % 64))
									v1011 = int64(32)
									v1012 = int64(base.Ui64(v208) >> (uint(v1011) % 64))
									v1014 = int64(base.Ui64(v1004) >> (uint(v1011) % 64))
									v1017 = int64(4294967295)
									v1018 = v208 & v1017
									v1020 = v1004 & v1017
									v1021 = v1018 * v1020
									v1025 = int64(base.Ui64(v1021)>>(uint(v1011)%64)) + v1018*v1014
									v1032 = v1020*v1012 + v1025&v1017
									*(*int64)(unsafe.Add(mBase, uint32(v999)+8)) = v1004*v217 + v1006*v208 + v1012*v1014 + int64(base.Ui64(v1025)>>(uint(v1011)%64)) + int64(base.Ui64(v1032)>>(uint(v1011)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v999))) = v1021&v1017 | v1032<<(uint(v1011)%64)
									v1045 = *(*int64)(unsafe.Add(mBase, uint32(v28)+104))
									v1047 = *(*int64)(unsafe.Add(mBase, uint32(v28)+96))
									v1048 = int64(0)
									v1055 = v206
									v1056 = v1006
									v1057 = v765
									v1059 = v206<<(uint(int64(48))%64) - v1045 - base.I64_extend_i32_u(base.B2i32(v1047 != v1048))
									v1060 = v1004
									v1062 = v782
									v1063 = v1048 - v1047
								}
								if int32(_a_F___divtf3_0) <= v1057 {
									v1325 = int64(0)
									v1335 = v36 | int64(9223090561878065152)
								} else {
									if int32(0) < v1057 {
										v1071 = int64(1)
										v1197 = v1059<<(uint(v1071)%64) | int64(base.Ui64(v1063)>>(uint(int64(63))%64))
										v1200 = v1056&int64(281474976710655) | base.I64_extend_i32_u(v1057)<<(uint(int64(48))%64)
										v1201 = v1060
										v1202 = v1063 << (uint(v1071) % 64)
										v1204 = v28 + int32(16)
										v1205 = int64(3)
										v1206 = int64(0)
										v1211 = int64(32)
										v1214 = int64(base.Ui64(v208) >> (uint(v1211) % 64))
										v1217 = int64(4294967295)
										v1220 = v208 & v1217
										v1221 = v1205 * v1220
										v1225 = int64(base.Ui64(v1221)>>(uint(v1211)%64)) + v1205*v1214
										v1232 = v1220*v1206 + v1225&v1217
										*(*int64)(unsafe.Add(mBase, uint32(v1204)+8)) = v208*v1206 + v217*v1205 + v1206*v1214 + int64(base.Ui64(v1225)>>(uint(v1211)%64)) + int64(base.Ui64(v1232)>>(uint(v1211)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v1204))) = v1221&v1217 | v1232<<(uint(v1211)%64)
										v1243 = int64(5)
										v1244 = int64(0)
										v1249 = int64(32)
										v1252 = int64(base.Ui64(v208) >> (uint(v1249) % 64))
										v1255 = int64(4294967295)
										v1258 = v208 & v1255
										v1259 = v1243 * v1258
										v1263 = int64(base.Ui64(v1259)>>(uint(v1249)%64)) + v1243*v1252
										v1270 = v1258*v1244 + v1263&v1255
										*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v208*v1244 + v217*v1243 + v1244*v1252 + int64(base.Ui64(v1263)>>(uint(v1249)%64)) + int64(base.Ui64(v1270)>>(uint(v1249)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1259&v1255 | v1270<<(uint(v1249)%64)
										v1282 = v1201 & int64(1)
										v1283 = v1282 + v1202
										v1287 = v1197 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1283) < base.Ui64(v1282)))
										if v1287 == v217 {
											v1290 = base.B2i32(base.Ui64(v208) < base.Ui64(v1283))
										} else {
											v1290 = base.B2i32(base.Ui64(v217) < base.Ui64(v1287))
										}
										v1292 = v1201 + base.I64_extend_i32_u(v1290)
										v1295 = v1200 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1292) < base.Ui64(v1201)))
										v1298 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
										v1300 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
										if v1287 == v1300 {
											v1303 = base.B2i32(base.Ui64(v1298) < base.Ui64(v1283))
										} else {
											v1303 = base.B2i32(base.Ui64(v1300) < base.Ui64(v1287))
										}
										v1306 = v1292 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1295) < base.Ui64(int64(9223090561878065152)))&v1303)
										v1309 = v1295 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1306) < base.Ui64(v1292)))
										v1312 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
										v1314 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
										if v1287 == v1314 {
											v1317 = base.B2i32(base.Ui64(v1312) < base.Ui64(v1283))
										} else {
											v1317 = base.B2i32(base.Ui64(v1314) < base.Ui64(v1287))
										}
										v1320 = v1306 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1309) < base.Ui64(int64(9223090561878065152)))&v1317)
										v1325 = v1320
										v1335 = v1309 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1320) < base.Ui64(v1306))) | v36
									} else {
										if v1057 <= int32(-113) {
											v1325 = int64(0)
											v1335 = v36
										} else {
											v1088 = v28 - int32(-64)
											v1090 = int32(1) - v1057
											if v1090&int32(64) != 0 {
												v1109 = int64(base.Ui64(v1056) >> (uint(base.I64_extend_i32_u(v1090+int32(-64))) % 64))
												v1110 = int64(0)
											} else {
												if v1090 == int32(0) {
													v1109 = v1060
													v1110 = v1056
												} else {
													v1105 = base.I64_extend_i32_u(v1090)
													v1109 = v1056<<(uint(base.I64_extend_i32_u(int32(64)-v1090))%64) | int64(base.Ui64(v1060)>>(uint(v1105)%64))
													v1110 = int64(base.Ui64(v1056) >> (uint(v1105) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v1088))) = v1109
											*(*int64)(unsafe.Add(mBase, uint32(v1088)+8)) = v1110
											v1115 = v28 + int32(48)
											v1117 = v1057 + int32(112)
											if v1117&int32(64) != 0 {
												v1136 = int64(0)
												v1137 = v1055 << (uint(base.I64_extend_i32_u(v1057+int32(48))) % 64)
											} else {
												if v1117 == int32(0) {
													v1136 = v1055
													v1137 = v1062
												} else {
													v1128 = base.I64_extend_i32_u(v1117)
													v1136 = v1055 << (uint(v1128) % 64)
													v1137 = v1062<<(uint(v1128)%64) | int64(base.Ui64(v1055)>>(uint(base.I64_extend_i32_u(int32(64)-v1117))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v1115))) = v1136
											*(*int64)(unsafe.Add(mBase, uint32(v1115)+8)) = v1137
											v1142 = v28 + int32(32)
											v1143 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
											v1144 = *(*int64)(unsafe.Add(mBase, uint32(v28)+72))
											v1149 = int64(32)
											v1150 = int64(base.Ui64(v1143) >> (uint(v1149) % 64))
											v1152 = int64(base.Ui64(v208) >> (uint(v1149) % 64))
											v1155 = int64(4294967295)
											v1156 = v1143 & v1155
											v1158 = v208 & v1155
											v1159 = v1156 * v1158
											v1163 = int64(base.Ui64(v1159)>>(uint(v1149)%64)) + v1156*v1152
											v1170 = v1158*v1150 + v1163&v1155
											*(*int64)(unsafe.Add(mBase, uint32(v1142)+8)) = v208*v1144 + v217*v1143 + v1150*v1152 + int64(base.Ui64(v1163)>>(uint(v1149)%64)) + int64(base.Ui64(v1170)>>(uint(v1149)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v1142))) = v1159&v1155 | v1170<<(uint(v1149)%64)
											v1181 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
											v1182 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
											v1183 = int64(1)
											v1185 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
											v1190 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
											v1192 = v1185 << (uint(v1183) % 64)
											v1197 = v1181 - (v1182<<(uint(v1183)%64) | int64(base.Ui64(v1185)>>(uint(int64(63))%64))) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1190) < base.Ui64(v1192)))
											v1200 = v1144
											v1201 = v1143
											v1202 = v1190 - v1192
											v1204 = v28 + int32(16)
											v1205 = int64(3)
											v1206 = int64(0)
											v1211 = int64(32)
											v1214 = int64(base.Ui64(v208) >> (uint(v1211) % 64))
											v1217 = int64(4294967295)
											v1220 = v208 & v1217
											v1221 = v1205 * v1220
											v1225 = int64(base.Ui64(v1221)>>(uint(v1211)%64)) + v1205*v1214
											v1232 = v1220*v1206 + v1225&v1217
											*(*int64)(unsafe.Add(mBase, uint32(v1204)+8)) = v208*v1206 + v217*v1205 + v1206*v1214 + int64(base.Ui64(v1225)>>(uint(v1211)%64)) + int64(base.Ui64(v1232)>>(uint(v1211)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v1204))) = v1221&v1217 | v1232<<(uint(v1211)%64)
											v1243 = int64(5)
											v1244 = int64(0)
											v1249 = int64(32)
											v1252 = int64(base.Ui64(v208) >> (uint(v1249) % 64))
											v1255 = int64(4294967295)
											v1258 = v208 & v1255
											v1259 = v1243 * v1258
											v1263 = int64(base.Ui64(v1259)>>(uint(v1249)%64)) + v1243*v1252
											v1270 = v1258*v1244 + v1263&v1255
											*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v208*v1244 + v217*v1243 + v1244*v1252 + int64(base.Ui64(v1263)>>(uint(v1249)%64)) + int64(base.Ui64(v1270)>>(uint(v1249)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1259&v1255 | v1270<<(uint(v1249)%64)
											v1282 = v1201 & int64(1)
											v1283 = v1282 + v1202
											v1287 = v1197 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1283) < base.Ui64(v1282)))
											if v1287 == v217 {
												v1290 = base.B2i32(base.Ui64(v208) < base.Ui64(v1283))
											} else {
												v1290 = base.B2i32(base.Ui64(v217) < base.Ui64(v1287))
											}
											v1292 = v1201 + base.I64_extend_i32_u(v1290)
											v1295 = v1200 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1292) < base.Ui64(v1201)))
											v1298 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
											v1300 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
											if v1287 == v1300 {
												v1303 = base.B2i32(base.Ui64(v1298) < base.Ui64(v1283))
											} else {
												v1303 = base.B2i32(base.Ui64(v1300) < base.Ui64(v1287))
											}
											v1306 = v1292 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1295) < base.Ui64(int64(9223090561878065152)))&v1303)
											v1309 = v1295 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1306) < base.Ui64(v1292)))
											v1312 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
											v1314 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
											if v1287 == v1314 {
												v1317 = base.B2i32(base.Ui64(v1312) < base.Ui64(v1283))
											} else {
												v1317 = base.B2i32(base.Ui64(v1314) < base.Ui64(v1287))
											}
											v1320 = v1306 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1309) < base.Ui64(int64(9223090561878065152)))&v1317)
											v1325 = v1320
											v1335 = v1309 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1320) < base.Ui64(v1306))) | v36
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
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1325
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v1335
	m.G0 = v28 + int32(336)
	return
}
func F_dacosh(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 float64
	_ = v33
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v45 int64
	_ = v45
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v67 int64
	_ = v67
	var v72 int32
	_ = v72
	var v85 float64
	_ = v85
	var v90 float64
	_ = v90
	var v104 float64
	_ = v104
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v117 float64
	_ = v117
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v155 float64
	_ = v155
	var v162 float64
	_ = v162
	var v166 float64
	_ = v166
	var v174 float64
	_ = v174
	var v175 float64
	_ = v175
	var v179 float64
	_ = v179
	v3 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.F64_lt(v3, float64(1)) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_dacosh_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_dacosh_1), int32(2744), int32(_a_F_dacosh_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v29 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v3))>>(uint(int64(52))%64))) & int32(2047)
		if base.Ui32(v29) <= base.Ui32(int32(1023)) {
			v33 = base.F64_add(v3, float64(-1))
			v38 = base.F64_add(v33, base.F64_sqrt(base.F64_add(base.F64_mul(v33, v33), base.F64_add(v33, v33))))
			v39 = float64(0)
			v45 = base.I64_reinterpret_f64(v38)
			if v45 <= int64(4601133429810003967) {
				if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v45) {
					if base.F64_eq(v38, float64(-1)) != 0 {
						v155 = math.Float64frombits(uint64(0xfff0000000000000))
						v162 = v155
					} else {
						v162 = base.F64_div(base.F64_sub(v38, v38), float64(0))
					}
				} else {
					if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v45)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
						v162 = v38
					} else {
						if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v45) {
							v65 = float64(1)
							v66 = base.F64_add(v38, v65)
							v67 = base.I64_reinterpret_f64(v66)
							v72 = base.I32_wrap_i64(int64(base.Ui64(v67)>>(uint(int64(32))%64))) + int32(_a_F_dacosh_3)
							if base.Ui32(int32(1074790399)) < base.Ui32(v72) {
								v85 = base.F64_add(base.F64_sub(v38, v66), v65)
							} else {
								v85 = base.F64_sub(v38, base.F64_add(v66, float64(-1)))
							}
							if base.Ui32(v72) <= base.Ui32(int32(1129316351)) {
								v90 = base.F64_div(v85, v66)
							} else {
								v90 = float64(0)
							}
							v104 = base.F64_convert_i32_s(int32(base.Ui32(v72)>>(uint(int32(20))%32)) - int32(1023))
							v108 = base.F64_add(base.F64_reinterpret_i64(v67&int64(4294967295)|base.I64_extend_i32_u(v72&int32(_a_F_dacosh_4)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v109 = v104
							v110 = base.F64_add(base.F64_mul(v104, float64(1.9082149292705877e-10)), v90)
						} else {
							v108 = v38
							v109 = v39
							v110 = v39
						}
						v117 = base.F64_div(v108, base.F64_add(v108, float64(2)))
						v120 = base.F64_mul(v108, base.F64_mul(v108, float64(0.5)))
						v121 = base.F64_mul(v117, v117)
						v122 = base.F64_mul(v121, v121)
						v155 = base.F64_add(base.F64_mul(v109, float64(0.6931471803691238)), base.F64_add(v108, base.F64_sub(base.F64_add(base.F64_mul(v117, base.F64_add(v120, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v121, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v110), v120)))
						v162 = v155
					}
				}
			} else {
				if base.Ui64(int64(9218868437227405311)) < base.Ui64(v45) {
					v162 = v38
				} else {
					v65 = float64(1)
					v66 = base.F64_add(v38, v65)
					v67 = base.I64_reinterpret_f64(v66)
					v72 = base.I32_wrap_i64(int64(base.Ui64(v67)>>(uint(int64(32))%64))) + int32(_a_F_dacosh_3)
					if base.Ui32(int32(1074790399)) < base.Ui32(v72) {
						v85 = base.F64_add(base.F64_sub(v38, v66), v65)
					} else {
						v85 = base.F64_sub(v38, base.F64_add(v66, float64(-1)))
					}
					if base.Ui32(v72) <= base.Ui32(int32(1129316351)) {
						v90 = base.F64_div(v85, v66)
					} else {
						v90 = float64(0)
					}
					v104 = base.F64_convert_i32_s(int32(base.Ui32(v72)>>(uint(int32(20))%32)) - int32(1023))
					v108 = base.F64_add(base.F64_reinterpret_i64(v67&int64(4294967295)|base.I64_extend_i32_u(v72&int32(_a_F_dacosh_4)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
					v109 = v104
					v110 = base.F64_add(base.F64_mul(v104, float64(1.9082149292705877e-10)), v90)
					v117 = base.F64_div(v108, base.F64_add(v108, float64(2)))
					v120 = base.F64_mul(v108, base.F64_mul(v108, float64(0.5)))
					v121 = base.F64_mul(v117, v117)
					v122 = base.F64_mul(v121, v121)
					v155 = base.F64_add(base.F64_mul(v109, float64(0.6931471803691238)), base.F64_add(v108, base.F64_sub(base.F64_add(base.F64_mul(v117, base.F64_add(v120, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v121, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v110), v120)))
					v162 = v155
				}
			}
			v179 = v162
		} else {
			if base.Ui32(v29) <= base.Ui32(int32(1048)) {
				v166 = float64(-1)
				v174 = F_log(m, base.F64_add(base.F64_add(v3, v3), base.F64_div(v166, base.F64_add(v3, base.F64_sqrt(base.F64_add(base.F64_mul(v3, v3), v166))))))
				mBase = m.M
				v179 = v174
			} else {
				v175 = F_log(m, v3)
				mBase = m.M
				v179 = base.F64_add(v175, float64(0.6931471805599453))
			}
		}
		return base.I64_reinterpret_f64(v179)
	}
}
func F_danish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v258 int32
	_ = v258
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
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
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v313 < v315 {
		goto L78
	} else {
		goto L79
	}
L2:
	;
	if v60 < int32(0) {
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
	v15 = v8
	v17 = int32(3)
	goto L9
L8:
	;
	v60 = v45
	goto L2
L9:
	;
	if v5 <= v15 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v60 = int32(-1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v22 = v15 + int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v15))))
	if base.Ui32(v24) < base.Ui32(int32(192)) {
		v45 = v22
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = int32(1)
	if v46 < v17 {
		v15 = v45
		v17 = v17 - v46
		goto L9
	} else {
		goto L21
	}
L15:
	;
	if v5 <= v22 {
		v45 = v22
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v31 = v22
	goto L17
L17:
	;
	v34 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7+v31))))
	if int32(-65) < v34 {
		v45 = v31
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v45 = v5
	goto L14
L19:
	;
	v38 = v31 + int32(1)
	if v38 != v5 {
		v31 = v38
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v85 = v8
	goto L25
L23:
	;
	if v180 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v180 = v152
	goto L23
L25:
	;
	if v76 <= v85 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v180 = int32(-1)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v92 = int32(1)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v77))))
	if base.Ui32(v94) < base.Ui32(int32(192)) {
		v151 = v94
		v152 = v92
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if int32(248) < v151 {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v98 = v85 + int32(1)
	if v98 == v76 {
		v151 = v94
		v152 = v92
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v77))))
	v103 = v101 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v94) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v77))))
	v119 = v117 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v94) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v107 = v85 + int32(2)
	if v107 != v76 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v151 = v94<<(uint(int32(6))%32)&int32(1984) | v103
	v152 = int32(2)
	goto L30
L37:
	;
	goto L36
L38:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v123))))
	v151 = v136&int32(63) | (v94<<(uint(int32(18))%32)&int32(_a_F_danish_UTF_8_stem_0) | v103<<(uint(int32(12))%32) | v119<<(uint(int32(6))%32))
	v152 = int32(4)
	goto L30
L39:
	;
	v123 = v85 + int32(3)
	if v123 != v76 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v151 = v94<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v103<<(uint(int32(6))%32) | v119
	v152 = int32(3)
	goto L30
L42:
	;
	goto L41
L43:
	;
	v169 = v152 + v85
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v169
	v85 = v169
	goto L25
L44:
	;
	v156 = v151 - int32(97)
	if v156 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v156)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[0]))))
	if int32(base.Ui32(v162)>>(uint(v156&int32(7))%32))&int32(1) != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v184 = v183 + v180
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v184
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v207 = v184
	goto L51
L49:
	;
	if v303 < int32(0) {
		goto L1
	} else {
		goto L73
	}
L50:
	;
	v303 = v274
	goto L49
L51:
	;
	if v198 <= v207 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v303 = int32(-1)
	goto L49
L54:
	;
	goto L55
L55:
	;
	v214 = int32(1)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v199))))
	if base.Ui32(v216) < base.Ui32(int32(192)) {
		v273 = v216
		v274 = v214
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if int32(248) < v273 {
		goto L50
	} else {
		goto L69
	}
L57:
	;
	v220 = v207 + int32(1)
	if v220 == v198 {
		v273 = v216
		v274 = v214
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220+v199))))
	v225 = v223 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v216) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+v199))))
	v241 = v239 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v216) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v229 = v207 + int32(2)
	if v229 != v198 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v273 = v216<<(uint(int32(6))%32)&int32(1984) | v225
	v274 = int32(2)
	goto L56
L63:
	;
	goto L62
L64:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v245))))
	v273 = v258&int32(63) | (v216<<(uint(int32(18))%32)&int32(_a_F_danish_UTF_8_stem_0) | v225<<(uint(int32(12))%32) | v241<<(uint(int32(6))%32))
	v274 = int32(4)
	goto L56
L65:
	;
	v245 = v207 + int32(3)
	if v245 != v198 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v273 = v216<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v225<<(uint(int32(6))%32) | v241
	v274 = int32(3)
	goto L56
L68:
	;
	goto L67
L69:
	;
	v278 = v273 - int32(97)
	if v278 < int32(0) {
		goto L50
	} else {
		goto L70
	}
L70:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v278)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[0]))))
	if int32(base.Ui32(v284)>>(uint(v278&int32(7))%32))&int32(1) == int32(0) {
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v292 = v274 + v207
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v292
	v207 = v292
	goto L51
L73:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v307 = v306 + v303
	if v60 < v307 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v309 = v307
	goto L76
L75:
	;
	v309 = v60
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v309
	goto L1
L77:
	;
	return v768
L78:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v489
	v491 = F_r_consonant_pair_2(m, l0)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L83
	} else {
		goto L114
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v315
	if v313 <= v315 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	goto L78
L81:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v322 = int32(1)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v313-v322))))
	if base.B2i32(v324&int32(224) != int32(96))|base.B2i32(v322<<(uint(v324)%32)&int32(_a_F_danish_UTF_8_stem_2) == int32(0)) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v339 = F_find_among_b(m, l0, int32(_a_F_danish_UTF_8_stem_3), int32(32), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return int32(0)
L84:
	;
	if v339 == int32(0) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v346
	switch v339 - int32(1) {
	case 0:
		goto L87
	case 1:
		goto L86
	default:
		goto L78
	}
L86:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L91
L87:
	;
	v350 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v350 {
		goto L78
	} else {
		goto L88
	}
L88:
	;
	v768 = v350
	goto L77
L89:
	;
	if v481 != 0 {
		goto L78
	} else {
		goto L112
	}
L90:
	;
	v481 = v474
	goto L89
L91:
	;
	if v365 <= v366 {
		v474 = int32(-1)
		goto L90
	} else {
		goto L93
	}
L92:
	;
	v474 = int32(0)
	goto L90
L93:
	;
	v383 = int32(1)
	v384 = v365 - v383
	v386 = int32(*(*int8)(unsafe.Add(mBase, uint32(v367+v384))))
	v388 = v386 & int32(255)
	if base.B2i32(v384 == v366)|base.B2i32(int32(0) <= v386) != 0 {
		v446 = v388
		v450 = v383
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if int32(229) < v446 {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	v395 = v388 & int32(63)
	v397 = v365 - int32(2)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367+v397))))
	v401 = v399 << (uint(int32(6)) % 32)
	if base.B2i32(v397 != v366)&base.B2i32(base.Ui32(v399) < base.Ui32(int32(192))) == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v446 = v401&int32(1984) | v395
	v450 = int32(2)
	goto L94
L97:
	;
	goto L98
L98:
	;
	v414 = v401&int32(4032) | v395
	v416 = v365 - int32(3)
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367+v416))))
	if base.B2i32(v416 != v366)&base.B2i32(base.Ui32(v418) < base.Ui32(int32(224))) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v446 = v418<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v414
	v450 = int32(3)
	goto L94
L100:
	;
	goto L101
L101:
	;
	v436 = int32(4)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365+v367-v436))))
	v446 = v418<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_4) | v438&int32(7)<<(uint(int32(18))%32) | v414
	v450 = v436
	goto L94
L102:
	;
	v481 = v450
	goto L89
L103:
	;
	goto L104
L104:
	;
	v452 = v446 - int32(97)
	if v452 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v481 = v450
	goto L89
L106:
	;
	goto L107
L107:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v452)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[1]))))
	if int32(base.Ui32(v458)>>(uint(v452&int32(7))%32))&int32(1) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v481 = v450
	goto L89
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365 - v450
	goto L111
L111:
	;
	goto L92
L112:
	;
	v482 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v482 {
		goto L78
	} else {
		goto L113
	}
L113:
	;
	v768 = v482
	goto L77
L114:
	;
	if v491 < int32(0) {
		v768 = v491
		goto L77
	} else {
		goto L115
	}
L115:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v495
	v498 = int32(2)
	v500 = int32(0)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v495-v503 < v498 {
		v513 = v500
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v540
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v540 < v542 {
		goto L128
	} else {
		goto L129
	}
L117:
	;
	if v513 == int32(0) {
		goto L116
	} else {
		goto L121
	}
L118:
	;
	goto L117
L119:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v509 = F_memcmp(m, v506+v495-v498, int32(_a_F_danish_UTF_8_stem_5), v498)
	mBase = m.M
	if v509 != 0 {
		v513 = v500
		goto L118
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v495 - v498
	v513 = int32(1)
	goto L118
L121:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v516
	v518 = int32(2)
	v520 = int32(0)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v516-v523 < v518 {
		v533 = v520
		goto L123
	} else {
		goto L124
	}
L122:
	;
	if v533 == int32(0) {
		goto L116
	} else {
		goto L126
	}
L123:
	;
	goto L122
L124:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v529 = F_memcmp(m, v526+v516-v518, int32(_a_F_danish_UTF_8_stem_6), v518)
	mBase = m.M
	if v529 != 0 {
		v533 = v520
		goto L123
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v516 - v518
	v533 = int32(1)
	goto L123
L126:
	;
	v536 = F_slice_del(m, l0)
	mBase = m.M
	if v536 < int32(0) {
		v768 = v536
		goto L77
	} else {
		goto L127
	}
L127:
	;
	goto L116
L128:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v593
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v593 < v595 {
		goto L142
	} else {
		goto L143
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v540
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v542
	v548 = v540 - int32(1)
	if v548 <= v542 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v545
	goto L128
L131:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550+v548))))
	if base.B2i32(v552&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v552)%32)&int32(_a_F_danish_UTF_8_stem_7) == int32(0)) != 0 {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v567 = F_find_among_b(m, l0, int32(_a_F_danish_UTF_8_stem_8), int32(5), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L83
	} else {
		goto L133
	}
L133:
	;
	if v567 == int32(0) {
		goto L130
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v545
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v572
	switch v567 - int32(1) {
	case 0:
		goto L136
	case 1:
		goto L135
	default:
		goto L128
	}
L135:
	;
	v585 = F_slice_from_s(m, l0, int32(4), int32(_a_F_danish_UTF_8_stem_9))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L83
	} else {
		goto L140
	}
L136:
	;
	v576 = F_slice_del(m, l0)
	mBase = m.M
	if v576 < int32(0) {
		v768 = v576
		goto L77
	} else {
		goto L137
	}
L137:
	;
	v579 = F_r_consonant_pair_2(m, l0)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L83
	} else {
		goto L138
	}
L138:
	;
	if int32(0) <= v579 {
		goto L128
	} else {
		goto L139
	}
L139:
	;
	v768 = v579
	goto L77
L140:
	;
	if int32(0) <= v585 {
		goto L128
	} else {
		goto L141
	}
L141:
	;
	v768 = v585
	goto L77
L142:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v765
	v768 = int32(1)
	goto L77
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v593
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v595
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L146
L144:
	;
	if v728 != 0 {
		goto L167
	} else {
		goto L168
	}
L145:
	;
	v728 = v721
	goto L144
L146:
	;
	if v612 <= v595 {
		v721 = int32(-1)
		goto L145
	} else {
		goto L148
	}
L147:
	;
	v721 = int32(0)
	goto L145
L148:
	;
	v630 = int32(1)
	v631 = v612 - v630
	v633 = int32(*(*int8)(unsafe.Add(mBase, uint32(v614+v631))))
	v635 = v633 & int32(255)
	if base.B2i32(v631 == v595)|base.B2i32(int32(0) <= v633) != 0 {
		v693 = v635
		v697 = v630
		goto L149
	} else {
		goto L150
	}
L149:
	;
	if int32(122) < v693 {
		goto L157
	} else {
		goto L158
	}
L150:
	;
	v642 = v635 & int32(63)
	v644 = v612 - int32(2)
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v644))))
	v648 = v646 << (uint(int32(6)) % 32)
	if base.B2i32(v644 != v595)&base.B2i32(base.Ui32(v646) < base.Ui32(int32(192))) == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v693 = v648&int32(1984) | v642
	v697 = int32(2)
	goto L149
L152:
	;
	goto L153
L153:
	;
	v661 = v648&int32(4032) | v642
	v663 = v612 - int32(3)
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v663))))
	if base.B2i32(v663 != v595)&base.B2i32(base.Ui32(v665) < base.Ui32(int32(224))) == int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v693 = v665<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v661
	v697 = int32(3)
	goto L149
L155:
	;
	goto L156
L156:
	;
	v683 = int32(4)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612+v614-v683))))
	v693 = v665<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_4) | v685&int32(7)<<(uint(int32(18))%32) | v661
	v697 = v683
	goto L149
L157:
	;
	v728 = v697
	goto L144
L158:
	;
	goto L159
L159:
	;
	v699 = v693 - int32(98)
	if v699 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v728 = v697
	goto L144
L161:
	;
	goto L162
L162:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v699)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[2]))))
	if int32(base.Ui32(v705)>>(uint(v699&int32(7))%32))&int32(1) == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v728 = v697
	goto L144
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v612 - v697
	goto L166
L166:
	;
	goto L147
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v598
	goto L142
L168:
	;
	goto L169
L169:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v730
	v734 = F_slice_to(m, l0, l0+int32(32))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L83
	} else {
		goto L170
	}
L170:
	;
	if v734 < int32(0) {
		v768 = v734
		goto L77
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v598
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v740 = int32(0)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v739-int32(4))))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v746-v598 < v745 {
		v757 = v740
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v757 == int32(0) {
		goto L142
	} else {
		goto L176
	}
L173:
	;
	goto L172
L174:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v753 = F_memcmp(m, v750+v746-v745, v739, v745)
	mBase = m.M
	if v753 != 0 {
		v757 = v740
		goto L173
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v746 - v745
	v757 = int32(1)
	goto L173
L176:
	;
	v760 = F_slice_del(m, l0)
	mBase = m.M
	if v760 < int32(0) {
		v768 = v760
		goto L77
	} else {
		goto L177
	}
L177:
	;
	goto L142
}
func F_datanh(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v19 float64
	_ = v19
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v30 float64
	_ = v30
	var v38 float64
	_ = v38
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v49 int64
	_ = v49
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v71 int64
	_ = v71
	var v76 int32
	_ = v76
	var v89 float64
	_ = v89
	var v94 float64
	_ = v94
	var v108 float64
	_ = v108
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v121 float64
	_ = v121
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v159 float64
	_ = v159
	var v166 float64
	_ = v166
	var v169 float64
	_ = v169
	var v174 float64
	_ = v174
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.F64_gt(base.F64_abs(v5), float64(1)) == int32(0) {
		if base.F64_eq(v5, float64(-1)) != 0 {
			return int64(-4503599627370496)
		} else {
			if base.F64_eq(v5, float64(1)) != 0 {
				return int64(9218868437227405312)
			} else {
				v19 = base.F64_abs(v5)
				v20 = base.I64_reinterpret_f64(v5)
				v25 = base.I32_wrap_i64(int64(base.Ui64(v20)>>(uint(int64(52))%64))) & int32(2047)
				if base.Ui32(v25) <= base.Ui32(int32(1021)) {
					if base.Ui32(v25) < base.Ui32(int32(991)) {
						v169 = v19
					} else {
						v30 = base.F64_add(v19, v19)
						v42 = base.F64_add(v30, base.F64_div(base.F64_mul(v19, v30), base.F64_sub(float64(1), v19)))
						v43 = float64(0)
						v49 = base.I64_reinterpret_f64(v42)
						if v49 <= int64(4601133429810003967) {
							if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v49) {
								if base.F64_eq(v42, float64(-1)) != 0 {
									v159 = math.Float64frombits(uint64(0xfff0000000000000))
									v166 = v159
								} else {
									v166 = base.F64_div(base.F64_sub(v42, v42), float64(0))
								}
							} else {
								if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v49)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
									v166 = v42
								} else {
									if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v49) {
										v69 = float64(1)
										v70 = base.F64_add(v42, v69)
										v71 = base.I64_reinterpret_f64(v70)
										v76 = base.I32_wrap_i64(int64(base.Ui64(v71)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
										if base.Ui32(int32(1074790399)) < base.Ui32(v76) {
											v89 = base.F64_add(base.F64_sub(v42, v70), v69)
										} else {
											v89 = base.F64_sub(v42, base.F64_add(v70, float64(-1)))
										}
										if base.Ui32(v76) <= base.Ui32(int32(1129316351)) {
											v94 = base.F64_div(v89, v70)
										} else {
											v94 = float64(0)
										}
										v108 = base.F64_convert_i32_s(int32(base.Ui32(v76)>>(uint(int32(20))%32)) - int32(1023))
										v112 = base.F64_add(base.F64_reinterpret_i64(v71&int64(4294967295)|base.I64_extend_i32_u(v76&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
										v113 = v108
										v114 = base.F64_add(base.F64_mul(v108, float64(1.9082149292705877e-10)), v94)
									} else {
										v112 = v42
										v113 = v43
										v114 = v43
									}
									v121 = base.F64_div(v112, base.F64_add(v112, float64(2)))
									v124 = base.F64_mul(v112, base.F64_mul(v112, float64(0.5)))
									v125 = base.F64_mul(v121, v121)
									v126 = base.F64_mul(v125, v125)
									v159 = base.F64_add(base.F64_mul(v113, float64(0.6931471803691238)), base.F64_add(v112, base.F64_sub(base.F64_add(base.F64_mul(v121, base.F64_add(v124, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v125, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v114), v124)))
									v166 = v159
								}
							}
						} else {
							if base.Ui64(int64(9218868437227405311)) < base.Ui64(v49) {
								v166 = v42
							} else {
								v69 = float64(1)
								v70 = base.F64_add(v42, v69)
								v71 = base.I64_reinterpret_f64(v70)
								v76 = base.I32_wrap_i64(int64(base.Ui64(v71)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
								if base.Ui32(int32(1074790399)) < base.Ui32(v76) {
									v89 = base.F64_add(base.F64_sub(v42, v70), v69)
								} else {
									v89 = base.F64_sub(v42, base.F64_add(v70, float64(-1)))
								}
								if base.Ui32(v76) <= base.Ui32(int32(1129316351)) {
									v94 = base.F64_div(v89, v70)
								} else {
									v94 = float64(0)
								}
								v108 = base.F64_convert_i32_s(int32(base.Ui32(v76)>>(uint(int32(20))%32)) - int32(1023))
								v112 = base.F64_add(base.F64_reinterpret_i64(v71&int64(4294967295)|base.I64_extend_i32_u(v76&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
								v113 = v108
								v114 = base.F64_add(base.F64_mul(v108, float64(1.9082149292705877e-10)), v94)
								v121 = base.F64_div(v112, base.F64_add(v112, float64(2)))
								v124 = base.F64_mul(v112, base.F64_mul(v112, float64(0.5)))
								v125 = base.F64_mul(v121, v121)
								v126 = base.F64_mul(v125, v125)
								v159 = base.F64_add(base.F64_mul(v113, float64(0.6931471803691238)), base.F64_add(v112, base.F64_sub(base.F64_add(base.F64_mul(v121, base.F64_add(v124, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v125, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v114), v124)))
								v166 = v159
							}
						}
						v169 = base.F64_mul(v166, float64(0.5))
					}
				} else {
					v38 = base.F64_div(v19, base.F64_sub(float64(1), v19))
					v42 = base.F64_add(v38, v38)
					v43 = float64(0)
					v49 = base.I64_reinterpret_f64(v42)
					if v49 <= int64(4601133429810003967) {
						if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v49) {
							if base.F64_eq(v42, float64(-1)) != 0 {
								v159 = math.Float64frombits(uint64(0xfff0000000000000))
								v166 = v159
							} else {
								v166 = base.F64_div(base.F64_sub(v42, v42), float64(0))
							}
						} else {
							if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v49)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
								v166 = v42
							} else {
								if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v49) {
									v69 = float64(1)
									v70 = base.F64_add(v42, v69)
									v71 = base.I64_reinterpret_f64(v70)
									v76 = base.I32_wrap_i64(int64(base.Ui64(v71)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
									if base.Ui32(int32(1074790399)) < base.Ui32(v76) {
										v89 = base.F64_add(base.F64_sub(v42, v70), v69)
									} else {
										v89 = base.F64_sub(v42, base.F64_add(v70, float64(-1)))
									}
									if base.Ui32(v76) <= base.Ui32(int32(1129316351)) {
										v94 = base.F64_div(v89, v70)
									} else {
										v94 = float64(0)
									}
									v108 = base.F64_convert_i32_s(int32(base.Ui32(v76)>>(uint(int32(20))%32)) - int32(1023))
									v112 = base.F64_add(base.F64_reinterpret_i64(v71&int64(4294967295)|base.I64_extend_i32_u(v76&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
									v113 = v108
									v114 = base.F64_add(base.F64_mul(v108, float64(1.9082149292705877e-10)), v94)
								} else {
									v112 = v42
									v113 = v43
									v114 = v43
								}
								v121 = base.F64_div(v112, base.F64_add(v112, float64(2)))
								v124 = base.F64_mul(v112, base.F64_mul(v112, float64(0.5)))
								v125 = base.F64_mul(v121, v121)
								v126 = base.F64_mul(v125, v125)
								v159 = base.F64_add(base.F64_mul(v113, float64(0.6931471803691238)), base.F64_add(v112, base.F64_sub(base.F64_add(base.F64_mul(v121, base.F64_add(v124, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v125, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v114), v124)))
								v166 = v159
							}
						}
					} else {
						if base.Ui64(int64(9218868437227405311)) < base.Ui64(v49) {
							v166 = v42
						} else {
							v69 = float64(1)
							v70 = base.F64_add(v42, v69)
							v71 = base.I64_reinterpret_f64(v70)
							v76 = base.I32_wrap_i64(int64(base.Ui64(v71)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
							if base.Ui32(int32(1074790399)) < base.Ui32(v76) {
								v89 = base.F64_add(base.F64_sub(v42, v70), v69)
							} else {
								v89 = base.F64_sub(v42, base.F64_add(v70, float64(-1)))
							}
							if base.Ui32(v76) <= base.Ui32(int32(1129316351)) {
								v94 = base.F64_div(v89, v70)
							} else {
								v94 = float64(0)
							}
							v108 = base.F64_convert_i32_s(int32(base.Ui32(v76)>>(uint(int32(20))%32)) - int32(1023))
							v112 = base.F64_add(base.F64_reinterpret_i64(v71&int64(4294967295)|base.I64_extend_i32_u(v76&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v113 = v108
							v114 = base.F64_add(base.F64_mul(v108, float64(1.9082149292705877e-10)), v94)
							v121 = base.F64_div(v112, base.F64_add(v112, float64(2)))
							v124 = base.F64_mul(v112, base.F64_mul(v112, float64(0.5)))
							v125 = base.F64_mul(v121, v121)
							v126 = base.F64_mul(v125, v125)
							v159 = base.F64_add(base.F64_mul(v113, float64(0.6931471803691238)), base.F64_add(v112, base.F64_sub(base.F64_add(base.F64_mul(v121, base.F64_add(v124, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v125, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, base.F64_add(base.F64_mul(v126, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v114), v124)))
							v166 = v159
						}
					}
					v169 = base.F64_mul(v166, float64(0.5))
				}
				if v20 < int64(0) {
					v174 = base.F64_neg(v169)
				} else {
					v174 = v169
				}
				return base.I64_reinterpret_f64(v174)
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v182 = m.ExcPending
		if v182 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v185 = m.ExcPending
			if v185 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_datanh_2), int32(0))
				mBase = m.M
				v189 = m.ExcPending
				if v189 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_datanh_3), int32(2768), int32(_a_F_datanh_4))
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int64(0)
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
func F_date2timestamptz_safe(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v135 int64
	_ = v135
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if l0 == int32(-2147483648) {
		v161 = int64(-9223372036854775807 - 1)
		m.G0 = v7 + int32(48)
		return v161
	} else {
		if l0 == int32(2147483647) {
			v161 = int64(9223372036854775807)
			m.G0 = v7 + int32(48)
			return v161
		} else {
			if int32(106751983) <= l0 {
				v17 = int64(9223372036854775807)
				v18 = F_errsave_start(m, l1)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					if v18 == int32(0) {
						v161 = v17
						m.G0 = v7 + int32(48)
						return v161
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_date2timestamptz_safe_0), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, l1, int32(_a_F_date2timestamptz_safe_1), int32(698), int32(_a_F_date2timestamptz_safe_2))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int64(0)
								} else {
									v161 = v17
									m.G0 = v7 + int32(48)
									return v161
								}
							}
						}
					}
				}
			} else {
				v47 = l0 + int32(_a_F_date2timestamptz_safe_3)
				v48 = int32(_a_F_date2timestamptz_safe_4)
				v49 = base.I32_div_u_s(v47, v48)
				v50 = int32(3)
				v56 = int32(2)
				v61 = base.I32_div_u_s((v49*int32(1073595727)+v47)<<(uint(v56)%32)|v50, v48)
				v64 = l0 + int32(_a_F_date2timestamptz_safe_5) + v49*v50 + v61 + int32(_a_F_date2timestamptz_safe_6)
				v65 = int32(1461)
				v66 = base.I32_div_u_s(v64, v65)
				v69 = v66*int32(-1461) + v64
				v71 = v69 << (uint(v56) % 32)
				if base.Ui32(v65) <= base.Ui32(v71) {
					v77 = base.I32_rem_u_s(v69+int32(305), int32(365))
					v82 = v77
				} else {
					v81 = base.I32_rem_u_s(v69+int32(306), int32(366))
					v82 = v81
				}
				v84 = base.I32_div_u_s(v71, int32(1461))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v84 + v66<<(uint(int32(2))%32) - int32(_a_F_date2timestamptz_safe_7)
				v92 = v82 + int32(123)
				v96 = int32(base.Ui32(v92*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v92 - int32(base.Ui32(v96*int32(_a_F_date2timestamptz_safe_8))>>(uint(int32(8))%32))
				v106 = base.I32_rem_u_s(v96+int32(10), int32(12))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(20)))) = v106 + int32(1)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v117 = *(*int32)(unsafe.Add(mBase, _c_F_date2timestamptz_safe[0]))
				v119 = m.G0
				v120 = int32(16)
				v121 = v119 - v120
				m.G0 = v121
				v125 = F_DetermineTimeZoneOffsetInternal(m, v7+int32(4), v117, v121+int32(8))
				mBase = m.M
				m.G0 = v121 + v120
				v135 = base.I64_extend_i32_s(v125)*int64(1000000) + base.I64_extend_i32_s(l0)*int64(86400000000)
				if base.Ui64(v135+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v161 = v135
					m.G0 = v7 + int32(48)
					return v161
				} else {
					if v135 < int64(-211813488000000000) {
						v144 = int64(-9223372036854775807 - 1)
					} else {
						v144 = int64(9223372036854775807)
					}
					v145 = F_errsave_start(m, l1)
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int64(0)
					} else {
						if v145 == int32(0) {
							v161 = v144
							m.G0 = v7 + int32(48)
							return v161
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_date2timestamptz_safe_0), int32(0))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, l1, int32(_a_F_date2timestamptz_safe_1), int32(723), int32(_a_F_date2timestamptz_safe_2))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int64(0)
									} else {
										v161 = v144
										m.G0 = v7 + int32(48)
										return v161
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
func F_dbase_redo(m *base.Module, l0 int32) {
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
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
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
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v177 int64
	_ = v177
	var v192 int64
	_ = v192
	var v201 int64
	_ = v201
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v225 int64
	_ = v225
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v281 int64
	_ = v281
	var v282 int32
	_ = v282
	var v284 int64
	_ = v284
	var v291 int64
	_ = v291
	var v293 int32
	_ = v293
	var v294 int64
	_ = v294
	var v300 int64
	_ = v300
	var v312 int64
	_ = v312
	var v318 int64
	_ = v318
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int64
	_ = v390
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v422 int64
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v597 int64
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v618 int32
	_ = v618
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
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+48)))
	v19 = v17 & int32(240)
	switch v19 - int32(16) {
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
	v767 = m.ExcPending
	if v767 != 0 {
		goto L8
	} else {
		goto L203
	}
L2:
	;
	m.G0 = v14 + int32(160)
	return
L3:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	v684 = F_GetDatabasePath(m, v682, v683)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L8
	} else {
		goto L175
	}
L4:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L8
	} else {
		goto L172
	}
L5:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[0]))
	if base.Ui32(int32(2)) <= base.Ui32(v435) {
		goto L110
	} else {
		goto L111
	}
L6:
	;
	if v19 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v25 = F_GetDatabasePath(m, v23, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v29 = F_GetDatabasePath(m, v27, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v60 = F_pstrdup(m, v29)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L21
	}
L11:
	;
	v35 = F___fstatat(m, int32(-100), v29, v14-int32(-64), int32(0))
	mBase = m.M
	goto L12
L12:
	;
	if v35 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	if v36&int32(_a_F_dbase_redo_0) != int32(_a_F_dbase_redo_1) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v41 = F_rmtree(m, v29)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v41 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v45 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	if v45 == int32(0) {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v29
	F_errmsg(m, int32(_a_F_dbase_redo_2), v14+int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3334), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v65 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v116 = F___fstatat(m, int32(-100), v60, v14-int32(-64), int32(0))
	mBase = m.M
	goto L44
L23:
	;
	v66 = F_strlen(m, v60)
	mBase = m.M
	v69 = v66 + v60
	goto L26
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	v73 = v69 - int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if base.B2i32(v74 == int32(47))&base.B2i32(base.Ui32(v60) < base.Ui32(v73)) != 0 {
		v69 = v73
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v80 = v73
	goto L29
L28:
	;
	goto L27
L29:
	;
	if base.Ui32(v60) < base.Ui32(v80) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v92 = v80
	goto L35
L31:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v86 != int32(47) {
		v80 = v80 - int32(1)
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L33
L35:
	;
	if base.Ui32(v60) < base.Ui32(v92) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v60 == v92 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v96 = v92 - int32(1)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v97 == int32(47) {
		v92 = v96
		goto L35
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	goto L36
L40:
	;
	goto L39
L41:
	;
	v105 = v60 + base.B2i32(v65 == int32(47))
	goto L43
L42:
	;
	v105 = v92
	goto L43
L43:
	;
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v106)
	goto L25
L44:
	;
	if v116 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[1]))
	if v120 != int32(44) {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_pfree(m, v60)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	F_recovery_create_dbdir(m, v60, int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v132 = F___fstatat(m, int32(-100), v25, v14-int32(-64), int32(0))
	mBase = m.M
	goto L52
L51:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v143 = m.G0
	v145 = v143 - int32(32)
	m.G0 = v145
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[2]))
	if int32(0) < v148 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	if int32(0) <= v132 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[1]))
	if v136 != int32(44) {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	F_recovery_create_dbdir(m, v25, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	v158 = int32(0)
	goto L59
L57:
	;
	goto L58
L58:
	;
	m.G0 = v145 + int32(32)
	v422 = F_EmitProcSignalBarrier(m, int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L8
	} else {
		goto L105
	}
L59:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[3]))
	v166 = v163 + v158*int32(56)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v167 != v142 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L58
L61:
	;
	v403 = v158 + int32(1)
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[2]))
	if v403 < v405 {
		v158 = v403
		goto L59
	} else {
		goto L104
	}
L62:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[4]))
	F_ResourceOwnerEnlarge(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	v175 = int64(4194304)
	v177 = base.AtomicRmwOr64(m, v166, int32(24), v175)
	if v177&v175 != int64(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v192 = v177
	goto L68
L66:
	;
	v281 = v177
	goto L67
L67:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v284 = int64(25165824)
	if base.B2i32(v282 != v142)|base.B2i32(v281&v284 != v284) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+28)) = int32(_a_F_dbase_redo_5)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+24)) = int32(_a_F_dbase_redo_6)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+20)) = int32(_a_F_dbase_redo_7)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+16)) = int32(0)
	v201 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v145)+8)) = v201
	if v192&int64(4194304) != v201 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v281 = v266
	goto L67
L70:
	;
	goto L73
L71:
	;
	goto L72
L72:
	;
	v244 = int32(_a_F_dbase_redo_8)
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[5]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v145+int32(8))+8))
	if v247 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L73:
	;
	F_perform_spin_delay(m, v145+int32(8))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L75
	}
L74:
	;
	goto L72
L75:
	;
	v222 = int64(0)
	v225 = base.AtomicRmwCmpxchg64(m, v166, int32(24), v222, v222)
	if v225&int64(4194304) != v222 {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v264 = int64(4194304)
	v266 = base.AtomicRmwOr64(m, v166, int32(24), v264)
	if v266&v264 != int64(0) {
		v192 = v266
		goto L68
	} else {
		goto L88
	}
L78:
	;
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[5])) = v262
	goto L78
L80:
	;
	if int32(999) < v245 {
		goto L78
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	if v245 < int32(11) {
		goto L78
	} else {
		goto L87
	}
L83:
	;
	v252 = int32(900)
	if v252 <= v245 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v255 = v252
	goto L86
L85:
	;
	v255 = v245
	goto L86
L86:
	;
	v262 = v255 + int32(100)
	goto L79
L87:
	;
	v262 = v245 - int32(1)
	goto L79
L88:
	;
	goto L69
L89:
	;
	v291 = int64(0)
	v293 = int32(24)
	v294 = base.AtomicRmwCmpxchg64(m, v166, v293, v291, v291)
	v300 = base.AtomicRmwCmpxchg64(m, v166, v293, v294, v294&int64(-4194305)+int64(1))
	if v294 != v300 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v390 = base.AtomicRmwSub64(m, v166, int32(24), int64(4194304))
	goto L61
L92:
	;
	v312 = v300
	goto L95
L93:
	;
	goto L94
L94:
	;
	v331 = int32(_a_F_dbase_redo_9)
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[6]))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v166)+20))
	v338 = int32(1)
	v339 = v337 + v338
	*(*int32)(unsafe.Add(mBase, uint32(v332<<(uint(int32(2))%32))+uint32(_c_F_dbase_redo[7]))) = v339
	v342 = v332 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v342)+uint32(_c_F_dbase_redo[8]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v342)+uint32(_c_F_dbase_redo[9]))) = v339
	*(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[6])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[10])) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v342)+uint32(_c_F_dbase_redo[11]))) = v338
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[4]))
	F_ResourceOwnerRemember(m, v360, base.I64_extend_i32_s(v339), int32(_a_F_dbase_redo_10))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L8
	} else {
		goto L98
	}
L95:
	;
	v318 = base.AtomicRmwCmpxchg64(m, v166, int32(24), v312, v312&int64(-4194305)+int64(1))
	if v312 != v318 {
		v312 = v318
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L94
L97:
	;
	goto L96
L98:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v166)+20))
	v367 = v365 + int32(1)
	F_BufferLockAcquire(m, v367, v166, int32(2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	F_FlushBuffer(m, v166, int32(0), int32(3))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	F_BufferLockUnlock(m, v367, v166)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[4]))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v166)+20))
	F_ResourceOwnerForget(m, v378, base.I64_extend_i32_s(v379+int32(1)), int32(_a_F_dbase_redo_10))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_UnpinBufferNoOwner(m, v166)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	goto L61
L104:
	;
	goto L60
L105:
	;
	F_WaitForProcSignalBarrier(m, v422)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	F_copydir(m, v25, v29, int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	F_pfree(m, v25)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	F_pfree(m, v29)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	goto L2
L110:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	F_LockSharedObjectForSession(m, v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L8
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	F_ReplicationSlotsDropDBSlots(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L8
	} else {
		goto L123
	}
L113:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v442 = F_CountDBBackends(m, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	if int32(0) < v442 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	goto L118
L116:
	;
	goto L117
L117:
	;
	goto L112
L118:
	;
	F_SignalRecoveryConflictWithDatabase(m, v441, int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L8
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	F_pg_usleep(m, int32(_a_F_dbase_redo_11))
	mBase = m.M
	v462 = F_CountDBBackends(m, v441)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L8
	} else {
		goto L121
	}
L121:
	;
	if int32(0) < v462 {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	goto L119
L123:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	F_DropDatabaseBuffers(m, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L8
	} else {
		goto L124
	}
L124:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	F_ForgetDatabaseSyncRequests(m, v494)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v498 = m.G0
	v500 = v498 - int32(112)
	m.G0 = v500
	F_smgrdestroyall(m)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[12]))
	if v505 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v597 = F_EmitProcSignalBarrier(m, int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L8
	} else {
		goto L153
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L8
	} else {
		goto L150
	}
L129:
	;
	m.G0 = v500 + int32(112)
	goto L127
L130:
	;
	v509 = v500 + int32(92)
	F_hash_seq_init(m, v509, v505)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L8
	} else {
		goto L131
	}
L131:
	;
	v512 = F_hash_seq_search(m, v509)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L8
	} else {
		goto L132
	}
L132:
	;
	if v512 == int32(0) {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	v516 = v512
	goto L134
L134:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v497 == v527 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L129
L136:
	;
	v531 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L8
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v567 = F_hash_seq_search(m, v500+int32(92))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L8
	} else {
		goto L148
	}
L139:
	;
	if v531 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	v535 = v500 + int32(20)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	F_GetRelationPath(m, v535, v536, v537, v538, int32(-1), v540)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L8
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[12]))
	v559 = F_hash_search(m, v556, v516, int32(2), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L8
	} else {
		goto L146
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v500)+4)) = v535
	F_errmsg_internal(m, int32(_a_F_dbase_redo_12), v500)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L8
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_13), int32(212), int32(_a_F_dbase_redo_14))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L8
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	if v559 == int32(0) {
		goto L128
	} else {
		goto L147
	}
L147:
	;
	goto L138
L148:
	;
	if v567 != 0 {
		v516 = v567
		goto L134
	} else {
		goto L149
	}
L149:
	;
	goto L135
L150:
	;
	F_errmsg_internal(m, int32(_a_F_dbase_redo_15), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L8
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_13), int32(217), int32(_a_F_dbase_redo_14))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L8
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_WaitForProcSignalBarrier(m, v597)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L8
	} else {
		goto L154
	}
L154:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	if int32(0) < v601 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v607 = int32(0)
	goto L158
L156:
	;
	goto L157
L157:
	;
	v662 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[0]))
	if base.Ui32(v662) < base.Ui32(int32(2)) {
		goto L2
	} else {
		goto L170
	}
L158:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v433+int32(8)+v607<<(uint(int32(2))%32))))
	v623 = F_GetDatabasePath(m, v618, v622)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L8
	} else {
		goto L161
	}
L159:
	;
	goto L157
L160:
	;
	F_pfree(m, v623)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L8
	} else {
		goto L168
	}
L161:
	;
	v625 = F_rmtree(m, v623)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L8
	} else {
		goto L162
	}
L162:
	;
	if v625 != 0 {
		goto L160
	} else {
		goto L163
	}
L163:
	;
	v629 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L8
	} else {
		goto L164
	}
L164:
	;
	if v629 == int32(0) {
		goto L160
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v623
	F_errmsg(m, int32(_a_F_dbase_redo_2), v14+int32(48))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L8
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3448), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L8
	} else {
		goto L167
	}
L167:
	;
	goto L160
L168:
	;
	v647 = v607 + int32(1)
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	if v647 < v648 {
		v607 = v647
		goto L158
	} else {
		goto L169
	}
L169:
	;
	goto L159
L170:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	F_UnlockSharedObjectForSession(m, v665)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L8
	} else {
		goto L171
	}
L171:
	;
	goto L2
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v19
	F_errmsg_internal(m, int32(_a_F_dbase_redo_16), v14)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L8
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3465), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L8
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	v686 = F_pstrdup(m, v684)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L8
	} else {
		goto L176
	}
L176:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686))))
	if v691 != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	F_recovery_create_dbdir(m, v686, int32(1))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L8
	} else {
		goto L199
	}
L178:
	;
	v692 = F_strlen(m, v686)
	mBase = m.M
	v695 = v692 + v686
	goto L181
L179:
	;
	goto L180
L180:
	;
	goto L177
L181:
	;
	v699 = v695 - int32(1)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if base.B2i32(v700 == int32(47))&base.B2i32(base.Ui32(v686) < base.Ui32(v699)) != 0 {
		v695 = v699
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v706 = v699
	goto L184
L183:
	;
	goto L182
L184:
	;
	if base.Ui32(v686) < base.Ui32(v706) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v718 = v706
	goto L190
L186:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
	if v712 != int32(47) {
		v706 = v706 - int32(1)
		goto L184
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	goto L185
L189:
	;
	goto L188
L190:
	;
	if base.Ui32(v686) < base.Ui32(v718) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	if v686 == v718 {
		goto L196
	} else {
		goto L197
	}
L192:
	;
	v722 = v718 - int32(1)
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722))))
	if v723 == int32(47) {
		v718 = v722
		goto L190
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	goto L191
L195:
	;
	goto L194
L196:
	;
	v731 = v686 + base.B2i32(v691 == int32(47))
	goto L198
L197:
	;
	v731 = v718
	goto L198
L198:
	;
	v732 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v731))) = uint8(v732)
	goto L180
L199:
	;
	F_pfree(m, v686)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L8
	} else {
		goto L200
	}
L200:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	F_CreateDirAndVersionFile(m, v684, v743, v744, int32(1))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L8
	} else {
		goto L201
	}
L201:
	;
	F_pfree(m, v684)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L8
	} else {
		goto L202
	}
L202:
	;
	goto L2
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v29
	F_errmsg(m, int32(_a_F_dbase_redo_17), v14+int32(16))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L8
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3348), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L8
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_decompile_conbin(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	v12 = v10 + v11
	v16 = F_heap_getattr_6(m, l0, int32(28), l1, v8+int32(15))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v20 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v27
				F_errmsg_internal(m, int32(_a_F_decompile_conbin_0), v8)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_decompile_conbin_1), int32(_a_F_decompile_conbin_2), int32(_a_F_decompile_conbin_3))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+80)))
			v40 = F_DirectFunctionCall2Coll(m, int32(623), int32(0), v16, v39)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v43 = F_text_to_cstring(m, base.I32_wrap_i64(v40))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v43
				}
			}
		}
	}
}
func F_decompress_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v5 = int32(1)
	if base.Ui32(v4-v5) <= base.Ui32(v5) {
		v10 = F_palloc0(m, int32(_a_F_decompress_init_0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_decompress_init_1)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = base.B2i32(v16 == int32(1))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v10
			v24 = int32(0)
			return v24
		}
	} else {
		v24 = int32(-102)
		return v24
	}
}
func F_decompress_read(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v12 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return v148
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v133
	if l2 < v135 {
		goto L42
	} else {
		goto L43
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v133 = v15
	v135 = v12
	goto L2
L4:
	;
	goto L5
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v19 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v148 = int32(0)
	goto L1
L9:
	;
	F_initStringInfo(m, v10+int32(16))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v106 = v19
	goto L11
L11:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v111 = l0 + int32(24)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = m.Env.Pgmem_zstream_read(m, v109, v111, v112)
	mBase = m.M
	if v113 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L12:
	;
	return int32(0)
L13:
	;
	v31 = F_pullf_read(m, l1, int32(_a_F_decompress_read_0), v10+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v69 = m.Env.Pgmem_inflate_all(m, v66, v67, v68)
	mBase = m.M
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v72 != 0 {
		goto L27
	} else {
		goto L28
	}
L15:
	;
	if int32(0) <= v31 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v39 = v31
	goto L19
L17:
	;
	v60 = v31
	goto L18
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L25
	}
L19:
	;
	if v39 == int32(0) {
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v60 = v52
	goto L18
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_appendBinaryStringInfo(m, v10+int32(16), v46, v39)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v52 = F_pullf_read(m, l1, int32(_a_F_decompress_read_0), v10+int32(12))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	if int32(0) <= v52 {
		v39 = v52
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v148 = v60
	goto L1
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L30
	}
L27:
	;
	base.MemoryFill(m, v70, int32(0), v72)
	goto L29
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	if v69 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_px_debug(m, int32(_a_F_decompress_read_1), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_decompress_read[0]))
	F_ResourceOwnerEnlarge(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L12
	} else {
		goto L35
	}
L34:
	;
	v148 = int32(-100)
	goto L1
L35:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_decompress_read[1]))
	v91 = F_MemoryContextAlloc(m, v89, int32(8))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v69
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_decompress_read[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v95
	F_ResourceOwnerRemember(m, v95, base.I64_extend_i32_u(v91), int32(_a_F_decompress_read_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v91
	v106 = v91
	goto L11
L38:
	;
	v148 = int32(-100)
	goto L1
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v111
	if v113 != 0 {
		v133 = v111
		v135 = v113
		goto L2
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(1)
	goto L8
L42:
	;
	v138 = l2
	goto L44
L43:
	;
	v138 = v135
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v135 - v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v141 + v138
	v148 = v138
	goto L1
}
func F_decrypt_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	return int32(_a_F_decrypt_init_0)
}
func F_decrypt_read(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_pullf_read(m, l1, l2, v9+int32(12))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if int32(0) < v13 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			v20 = F_pgp_cfb_decrypt(m, l0, v19, v13, l4)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = l4
				m.G0 = v9 + int32(16)
				return v13
			}
		} else {
			m.G0 = v9 + int32(16)
			return v13
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
func F_des_setkey(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v55 int32
	_ = v55
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v763 int64
	_ = v763
	var v787 int32
	_ = v787
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v842 int64
	_ = v842
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v986 int32
	_ = v986
	var v999 int32
	_ = v999
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1347 int32
	_ = v1347
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1373 int32
	_ = v1373
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1544 int32
	_ = v1544
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1605 int32
	_ = v1605
	var v1613 int32
	_ = v1613
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1682 int32
	_ = v1682
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1708 int32
	_ = v1708
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1721 int32
	_ = v1721
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1734 int32
	_ = v1734
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1747 int32
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1760 int32
	_ = v1760
	var v1765 int32
	_ = v1765
	var v1769 int32
	_ = v1769
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1798 int32
	_ = v1798
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1862 int32
	_ = v1862
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1905 int32
	_ = v1905
	var v1907 int32
	_ = v1907
	var v1915 int32
	_ = v1915
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2019 int32
	_ = v2019
	var v2022 int32
	_ = v2022
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2041 int32
	_ = v2041
	var v2047 int32
	_ = v2047
	v2 = int32(0)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_des_setkey[0])))
	if v17 == v2 {
		v20 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[1])) = v20
		*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[2])) = v20
		*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[3])) = v20
		*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[4])) = v20
		v55 = v20
		for {
			v86 = v55&int32(32) + int32(_a_F_des_setkey_0) + int32(base.Ui32(v55)>>(uint(int32(1))%32))&int32(15)
			v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+16)))
			*(*uint8)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_des_setkey[5]))) = uint8(v87)
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
			*(*uint8)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_des_setkey[6]))) = uint8(v89)
			v92 = v55 + int32(2)
			if v92 != int32(64) {
				v55 = v92
				continue
			} else {
				break
			}
			break
		}
		v95 = v20
		for {
			v127 = v95&int32(32) + int32(_a_F_des_setkey_0) + int32(base.Ui32(v95)>>(uint(int32(1))%32))&int32(15)
			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+80)))
			*(*uint8)(unsafe.Add(mBase, uint32(v95)+uint32(_c_F_des_setkey[7]))) = uint8(v128)
			v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+64)))
			*(*uint8)(unsafe.Add(mBase, uint32(v95)+uint32(_c_F_des_setkey[8]))) = uint8(v130)
			v133 = v95 + int32(2)
			if v133 != int32(64) {
				v95 = v133
				continue
			} else {
				break
			}
			break
		}
		v137 = int32(0)
		for {
			v169 = v137&int32(32) + int32(_a_F_des_setkey_0) + int32(base.Ui32(v137)>>(uint(int32(1))%32))&int32(15)
			v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+144)))
			*(*uint8)(unsafe.Add(mBase, uint32(v137)+uint32(_c_F_des_setkey[9]))) = uint8(v170)
			v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+128)))
			*(*uint8)(unsafe.Add(mBase, uint32(v137)+uint32(_c_F_des_setkey[10]))) = uint8(v172)
			v175 = v137 + int32(2)
			if v175 != int32(64) {
				v137 = v175
				continue
			} else {
				break
			}
			break
		}
		v179 = int32(0)
		for {
			v211 = v179&int32(32) + int32(_a_F_des_setkey_0) + int32(base.Ui32(v179)>>(uint(int32(1))%32))&int32(15)
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+208)))
			*(*uint8)(unsafe.Add(mBase, uint32(v179)+uint32(_c_F_des_setkey[11]))) = uint8(v212)
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+192)))
			*(*uint8)(unsafe.Add(mBase, uint32(v179)+uint32(_c_F_des_setkey[12]))) = uint8(v214)
			v217 = v179 + int32(2)
			if v217 != int32(64) {
				v179 = v217
				continue
			} else {
				break
			}
			break
		}
		v221 = int32(0)
		for {
			v253 = v221&int32(32) + int32(_a_F_des_setkey_0) + int32(base.Ui32(v221)>>(uint(int32(1))%32))&int32(15)
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+272)))
			*(*uint8)(unsafe.Add(mBase, uint32(v221)+uint32(_c_F_des_setkey[13]))) = uint8(v254)
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+256)))
			*(*uint8)(unsafe.Add(mBase, uint32(v221)+uint32(_c_F_des_setkey[14]))) = uint8(v256)
			v259 = v221 + int32(2)
			if v259 != int32(64) {
				v221 = v259
				continue
			} else {
				break
			}
			break
		}
		v263 = int32(0)
		for {
			v295 = v263&int32(32) + int32(_a_F_des_setkey_0) + int32(base.Ui32(v263)>>(uint(int32(1))%32))&int32(15)
			v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+336)))
			*(*uint8)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_des_setkey[15]))) = uint8(v296)
			v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+320)))
			*(*uint8)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_des_setkey[16]))) = uint8(v298)
			v301 = v263 + int32(2)
			if v301 != int32(64) {
				v263 = v301
				continue
			} else {
				break
			}
			break
		}
		v305 = int32(0)
		for {
			v337 = v305&int32(32) + int32(_a_F_des_setkey_0) + int32(base.Ui32(v305)>>(uint(int32(1))%32))&int32(15)
			v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+400)))
			*(*uint8)(unsafe.Add(mBase, uint32(v305)+uint32(_c_F_des_setkey[17]))) = uint8(v338)
			v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+384)))
			*(*uint8)(unsafe.Add(mBase, uint32(v305)+uint32(_c_F_des_setkey[18]))) = uint8(v340)
			v343 = v305 + int32(2)
			if v343 != int32(64) {
				v305 = v343
				continue
			} else {
				break
			}
			break
		}
		v347 = int32(0)
		for {
			v379 = v347&int32(32) + int32(_a_F_des_setkey_0) + int32(base.Ui32(v347)>>(uint(int32(1))%32))&int32(15)
			v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+464)))
			*(*uint8)(unsafe.Add(mBase, uint32(v347)+uint32(_c_F_des_setkey[19]))) = uint8(v380)
			v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+448)))
			*(*uint8)(unsafe.Add(mBase, uint32(v347)+uint32(_c_F_des_setkey[20]))) = uint8(v382)
			v385 = v347 + int32(2)
			if v385 != int32(64) {
				v347 = v385
				continue
			} else {
				break
			}
			break
		}
		v390 = v20
		for {
			v413 = v390<<(uint(int32(6))%32) + int32(_a_F_des_setkey_1)
			v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+uint32(_c_F_des_setkey[6]))))
			v418 = v416 << (uint(int32(4)) % 32)
			v420 = int32(0)
			for {
				v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420)+uint32(_c_F_des_setkey[8]))))
				v446 = v418 | v445
				*(*uint8)(unsafe.Add(mBase, uint32(v420+v413))) = uint8(v446)
				v449 = v420 | int32(1)
				v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449)+uint32(_c_F_des_setkey[8]))))
				v454 = v418 | v453
				*(*uint8)(unsafe.Add(mBase, uint32(v413+v449))) = uint8(v454)
				v457 = v420 | int32(2)
				v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+uint32(_c_F_des_setkey[8]))))
				v462 = v418 | v461
				*(*uint8)(unsafe.Add(mBase, uint32(v413+v457))) = uint8(v462)
				v465 = v420 | int32(3)
				v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+uint32(_c_F_des_setkey[8]))))
				v470 = v418 | v469
				*(*uint8)(unsafe.Add(mBase, uint32(v413+v465))) = uint8(v470)
				v473 = v420 + int32(4)
				if v473 != int32(64) {
					v420 = v473
					continue
				} else {
					break
				}
				break
			}
			v477 = v390 + int32(1)
			if v477 != int32(64) {
				v390 = v477
				continue
			} else {
				break
			}
			break
		}
		v483 = int32(0)
		for {
			v508 = v483<<(uint(int32(6))%32) + int32(_a_F_des_setkey_2)
			v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_des_setkey[10]))))
			v513 = v511 << (uint(int32(4)) % 32)
			v514 = int32(0)
			for {
				v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+uint32(_c_F_des_setkey[12]))))
				v540 = v513 | v539
				*(*uint8)(unsafe.Add(mBase, uint32(v514+v508))) = uint8(v540)
				v543 = v514 | int32(1)
				v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+uint32(_c_F_des_setkey[12]))))
				v548 = v513 | v547
				*(*uint8)(unsafe.Add(mBase, uint32(v508+v543))) = uint8(v548)
				v551 = v514 | int32(2)
				v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+uint32(_c_F_des_setkey[12]))))
				v556 = v513 | v555
				*(*uint8)(unsafe.Add(mBase, uint32(v508+v551))) = uint8(v556)
				v559 = v514 | int32(3)
				v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+uint32(_c_F_des_setkey[12]))))
				v564 = v513 | v563
				*(*uint8)(unsafe.Add(mBase, uint32(v508+v559))) = uint8(v564)
				v567 = v514 + int32(4)
				if v567 != int32(64) {
					v514 = v567
					continue
				} else {
					break
				}
				break
			}
			v571 = v483 + int32(1)
			if v571 != int32(64) {
				v483 = v571
				continue
			} else {
				break
			}
			break
		}
		v577 = int32(0)
		for {
			v602 = v577<<(uint(int32(6))%32) + int32(_a_F_des_setkey_3)
			v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+uint32(_c_F_des_setkey[14]))))
			v607 = v605 << (uint(int32(4)) % 32)
			v608 = int32(0)
			for {
				v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608)+uint32(_c_F_des_setkey[16]))))
				v634 = v607 | v633
				*(*uint8)(unsafe.Add(mBase, uint32(v608+v602))) = uint8(v634)
				v637 = v608 | int32(1)
				v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637)+uint32(_c_F_des_setkey[16]))))
				v642 = v607 | v641
				*(*uint8)(unsafe.Add(mBase, uint32(v602+v637))) = uint8(v642)
				v645 = v608 | int32(2)
				v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+uint32(_c_F_des_setkey[16]))))
				v650 = v607 | v649
				*(*uint8)(unsafe.Add(mBase, uint32(v602+v645))) = uint8(v650)
				v653 = v608 | int32(3)
				v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+uint32(_c_F_des_setkey[16]))))
				v658 = v607 | v657
				*(*uint8)(unsafe.Add(mBase, uint32(v602+v653))) = uint8(v658)
				v661 = v608 + int32(4)
				if v661 != int32(64) {
					v608 = v661
					continue
				} else {
					break
				}
				break
			}
			v665 = v577 + int32(1)
			if v665 != int32(64) {
				v577 = v665
				continue
			} else {
				break
			}
			break
		}
		v671 = int32(0)
		for {
			v696 = v671<<(uint(int32(6))%32) + int32(_a_F_des_setkey_4)
			v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671)+uint32(_c_F_des_setkey[18]))))
			v701 = v699 << (uint(int32(4)) % 32)
			v702 = int32(0)
			for {
				v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702)+uint32(_c_F_des_setkey[20]))))
				v728 = v701 | v727
				*(*uint8)(unsafe.Add(mBase, uint32(v702+v696))) = uint8(v728)
				v731 = v702 | int32(1)
				v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+uint32(_c_F_des_setkey[20]))))
				v736 = v701 | v735
				*(*uint8)(unsafe.Add(mBase, uint32(v696+v731))) = uint8(v736)
				v739 = v702 | int32(2)
				v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739)+uint32(_c_F_des_setkey[20]))))
				v744 = v701 | v743
				*(*uint8)(unsafe.Add(mBase, uint32(v696+v739))) = uint8(v744)
				v747 = v702 | int32(3)
				v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747)+uint32(_c_F_des_setkey[20]))))
				v752 = v701 | v751
				*(*uint8)(unsafe.Add(mBase, uint32(v696+v747))) = uint8(v752)
				v755 = v702 + int32(4)
				if v755 != int32(64) {
					v702 = v755
					continue
				} else {
					break
				}
				break
			}
			v759 = v671 + int32(1)
			if v759 != int32(64) {
				v671 = v759
				continue
			} else {
				break
			}
			break
		}
		v763 = int64(-1)
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[21])) = v763
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[22])) = v763
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[23])) = v763
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[24])) = v763
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[25])) = v763
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[26])) = v763
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[27])) = v763
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[28])) = v763
		v787 = int32(0)
		for {
			v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787)+uint32(_c_F_des_setkey[29]))))
			v814 = int32(1)
			v815 = v813 - v814
			*(*uint8)(unsafe.Add(mBase, uint32(v787)+uint32(_c_F_des_setkey[30]))) = uint8(v815)
			v817 = int32(255)
			*(*uint8)(unsafe.Add(mBase, uint32(v815&v817)+uint32(_c_F_des_setkey[31]))) = uint8(v787)
			v823 = v787 | v814
			v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823)+uint32(_c_F_des_setkey[29]))))
			v830 = v828 - v814
			*(*uint8)(unsafe.Add(mBase, uint32(v823)+uint32(_c_F_des_setkey[30]))) = uint8(v830)
			*(*uint8)(unsafe.Add(mBase, uint32(v830&v817)+uint32(_c_F_des_setkey[31]))) = uint8(v823)
			v838 = v787 + int32(2)
			if v838 != int32(64) {
				v787 = v838
				continue
			} else {
				break
			}
			break
		}
		v842 = int64(-1)
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[32])) = v842
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[33])) = v842
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[34])) = v842
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[35])) = v842
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[36])) = v842
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[37])) = v842
		*(*int64)(unsafe.Add(mBase, _c_F_des_setkey[38])) = v842
		v862 = int32(0)
		v865 = v862
		for {
			v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865)+uint32(_c_F_des_setkey[39]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v888)+uint32(_c_F_des_setkey[40]))) = uint8(v865)
			v893 = v865 | int32(1)
			v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893)+uint32(_c_F_des_setkey[39]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v896)+uint32(_c_F_des_setkey[40]))) = uint8(v893)
			v901 = v865 | int32(2)
			v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v901)+uint32(_c_F_des_setkey[39]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v904)+uint32(_c_F_des_setkey[40]))) = uint8(v901)
			v909 = v865 | int32(3)
			v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909)+uint32(_c_F_des_setkey[39]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v912)+uint32(_c_F_des_setkey[40]))) = uint8(v909)
			v917 = v865 + int32(4)
			if v917 != int32(56) {
				v865 = v917
				continue
			} else {
				break
			}
			break
		}
		v920 = v862
		for {
			v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+uint32(_c_F_des_setkey[41]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v944)+uint32(_c_F_des_setkey[42]))) = uint8(v920)
			v949 = v920 | int32(1)
			v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v949)+uint32(_c_F_des_setkey[41]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v952)+uint32(_c_F_des_setkey[42]))) = uint8(v949)
			v957 = v920 | int32(2)
			v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v957)+uint32(_c_F_des_setkey[41]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v960)+uint32(_c_F_des_setkey[42]))) = uint8(v957)
			v965 = v920 | int32(3)
			v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+uint32(_c_F_des_setkey[41]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v968)+uint32(_c_F_des_setkey[42]))) = uint8(v965)
			v973 = v920 + int32(4)
			if v973 != int32(48) {
				v920 = v973
				continue
			} else {
				break
			}
			break
		}
		v986 = v20
		for {
			v999 = v986 << (uint(int32(10)) % 32)
			v1009 = v986 << (uint(int32(3)) % 32)
			v1012 = int32(0)
			for {
				v1034 = v1012 << (uint(int32(2)) % 32)
				v1035 = v999 + int32(_a_F_des_setkey_5) + v1034
				v1036 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v1035))) = v1036
				v1038 = v1034 + (v999 + int32(_a_F_des_setkey_6))
				*(*int32)(unsafe.Add(mBase, uint32(v1038))) = v1036
				v1041 = v1034 + (v999 + int32(_a_F_des_setkey_7))
				*(*int32)(unsafe.Add(mBase, uint32(v1041))) = v1036
				v1044 = v1034 + (v999 + int32(_a_F_des_setkey_8))
				*(*int32)(unsafe.Add(mBase, uint32(v1044))) = v1036
				v1052 = v1036
				v1054 = v1036
				v1056 = v1036
				v1059 = v1036
				v1061 = v1036
				for {
					v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+uint32(_c_F_des_setkey[43]))))
					if v1012&v1076 == int32(0) {
						v1117 = v1054
						v1118 = v1056
						v1120 = v1059
						v1121 = v1061
					} else {
						v1080 = v1052 + v1009
						v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1080)+uint32(_c_F_des_setkey[31]))))
						v1085 = v1083 << (uint(int32(2)) % 32)
						if base.Ui32(v1083) <= base.Ui32(int32(31)) {
							v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+uint32(_c_F_des_setkey[44])))
							v1091 = v1059 | v1090
							*(*int32)(unsafe.Add(mBase, uint32(v1035))) = v1091
							v1098 = v1091
							v1099 = v1061
						} else {
							v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1085+int32(_a_F_des_setkey_9)-int32(128))))
							v1096 = v1061 | v1095
							*(*int32)(unsafe.Add(mBase, uint32(v1038))) = v1096
							v1098 = v1059
							v1099 = v1096
						}
						v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1080)+uint32(_c_F_des_setkey[30]))))
						v1104 = v1102 << (uint(int32(2)) % 32)
						if base.Ui32(v1102) <= base.Ui32(int32(31)) {
							v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+uint32(_c_F_des_setkey[44])))
							v1110 = v1056 | v1109
							*(*int32)(unsafe.Add(mBase, uint32(v1041))) = v1110
							v1117 = v1054
							v1118 = v1110
							v1120 = v1098
							v1121 = v1099
						} else {
							v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1104+int32(_a_F_des_setkey_9)-int32(128))))
							v1115 = v1054 | v1114
							*(*int32)(unsafe.Add(mBase, uint32(v1044))) = v1115
							v1117 = v1115
							v1118 = v1056
							v1120 = v1098
							v1121 = v1099
						}
					}
					v1125 = v1052 + int32(1)
					if v1125 != int32(8) {
						v1052 = v1125
						v1054 = v1117
						v1056 = v1118
						v1059 = v1120
						v1061 = v1121
						continue
					} else {
						break
					}
					break
				}
				v1129 = v1012 + int32(1)
				if v1129 != int32(256) {
					v1012 = v1129
					continue
				} else {
					break
				}
				break
			}
			v1133 = v986 << (uint(int32(9)) % 32)
			v1144 = v986 * int32(7)
			v1145 = int32(0)
			for {
				v1168 = v1145 << (uint(int32(2)) % 32)
				v1169 = v1133 + int32(_a_F_des_setkey_10) + v1168
				v1170 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1170
				v1172 = v1168 + (v1133 + int32(_a_F_des_setkey_11))
				*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1170
				v1178 = v1145 & int32(64)
				if v1178 == v1170 {
					v1199 = v1170
					v1200 = v1170
				} else {
					v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009)+uint32(_c_F_des_setkey[28]))))
					if v1183 == int32(255) {
						v1199 = v1170
						v1200 = v1170
					} else {
						v1187 = v1183 << (uint(int32(2)) % 32)
						if base.Ui32(v1183) <= base.Ui32(int32(27)) {
							v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+uint32(_c_F_des_setkey[45])))
							*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1192
							v1199 = v1170
							v1200 = v1192
						} else {
							v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1187+int32(_a_F_des_setkey_12)-int32(112))))
							*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1196
							v1199 = v1196
							v1200 = int32(0)
						}
					}
				}
				v1203 = v1145 & int32(32)
				if v1203 == int32(0) {
					v1225 = v1199
					v1226 = v1200
				} else {
					v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009)+uint32(_c_F_des_setkey[46]))))
					if v1208 == int32(255) {
						v1225 = v1199
						v1226 = v1200
					} else {
						v1212 = v1208 << (uint(int32(2)) % 32)
						if base.Ui32(int32(28)) <= base.Ui32(v1208) {
							v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1212+int32(_a_F_des_setkey_12)-int32(112))))
							v1220 = v1199 | v1219
							*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1220
							v1225 = v1220
							v1226 = v1200
						} else {
							v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+uint32(_c_F_des_setkey[45])))
							v1223 = v1200 | v1222
							*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1223
							v1225 = v1199
							v1226 = v1223
						}
					}
				}
				v1230 = v1145 & int32(16)
				if v1230 == int32(0) {
					v1252 = v1225
					v1253 = v1226
				} else {
					v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009)+uint32(_c_F_des_setkey[47]))))
					if v1235 == int32(255) {
						v1252 = v1225
						v1253 = v1226
					} else {
						v1239 = v1235 << (uint(int32(2)) % 32)
						if base.Ui32(int32(28)) <= base.Ui32(v1235) {
							v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1239+int32(_a_F_des_setkey_12)-int32(112))))
							v1247 = v1225 | v1246
							*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1247
							v1252 = v1247
							v1253 = v1226
						} else {
							v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+uint32(_c_F_des_setkey[45])))
							v1250 = v1226 | v1249
							*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1250
							v1252 = v1225
							v1253 = v1250
						}
					}
				}
				v1257 = v1145 & int32(8)
				if v1257 == int32(0) {
					v1279 = v1252
					v1280 = v1253
				} else {
					v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009)+uint32(_c_F_des_setkey[48]))))
					if v1262 == int32(255) {
						v1279 = v1252
						v1280 = v1253
					} else {
						v1266 = v1262 << (uint(int32(2)) % 32)
						if base.Ui32(int32(28)) <= base.Ui32(v1262) {
							v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1266+int32(_a_F_des_setkey_12)-int32(112))))
							v1274 = v1252 | v1273
							*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1274
							v1279 = v1274
							v1280 = v1253
						} else {
							v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+uint32(_c_F_des_setkey[45])))
							v1277 = v1253 | v1276
							*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1277
							v1279 = v1252
							v1280 = v1277
						}
					}
				}
				v1284 = v1145 & int32(4)
				if v1284 == int32(0) {
					v1306 = v1279
					v1307 = v1280
				} else {
					v1289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009)+uint32(_c_F_des_setkey[49]))))
					if v1289 == int32(255) {
						v1306 = v1279
						v1307 = v1280
					} else {
						v1293 = v1289 << (uint(int32(2)) % 32)
						if base.Ui32(int32(28)) <= base.Ui32(v1289) {
							v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1293+int32(_a_F_des_setkey_12)-int32(112))))
							v1301 = v1279 | v1300
							*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1301
							v1306 = v1301
							v1307 = v1280
						} else {
							v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+uint32(_c_F_des_setkey[45])))
							v1304 = v1280 | v1303
							*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1304
							v1306 = v1279
							v1307 = v1304
						}
					}
				}
				v1311 = v1145 & int32(2)
				if v1311 == int32(0) {
					v1333 = v1306
					v1334 = v1307
				} else {
					v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009)+uint32(_c_F_des_setkey[50]))))
					if v1316 == int32(255) {
						v1333 = v1306
						v1334 = v1307
					} else {
						v1320 = v1316 << (uint(int32(2)) % 32)
						if base.Ui32(int32(28)) <= base.Ui32(v1316) {
							v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1320+int32(_a_F_des_setkey_12)-int32(112))))
							v1328 = v1306 | v1327
							*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1328
							v1333 = v1328
							v1334 = v1307
						} else {
							v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+uint32(_c_F_des_setkey[45])))
							v1331 = v1307 | v1330
							*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1331
							v1333 = v1306
							v1334 = v1331
						}
					}
				}
				v1338 = v1145 & int32(1)
				if v1338 == int32(0) {
				} else {
					v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009)+uint32(_c_F_des_setkey[51]))))
					if v1343 == int32(255) {
					} else {
						v1347 = v1343 << (uint(int32(2)) % 32)
						if base.Ui32(int32(28)) <= base.Ui32(v1343) {
							v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1347+int32(_a_F_des_setkey_12)-int32(112))))
							*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1333 | v1354
						} else {
							v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+uint32(_c_F_des_setkey[45])))
							*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1334 | v1357
						}
					}
				}
				v1362 = int32(0)
				v1363 = v1168 + (v1133 + int32(_a_F_des_setkey_13))
				*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1362
				v1366 = v1168 + (v1133 + int32(_a_F_des_setkey_14))
				*(*int32)(unsafe.Add(mBase, uint32(v1366))) = v1362
				if v1178 == v1362 {
					v1389 = v1362
					v1392 = int32(0)
					v1393 = v1389
				} else {
					v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_des_setkey[38]))))
					if v1373 == int32(255) {
						v1389 = v1362
						v1392 = int32(0)
						v1393 = v1389
					} else {
						v1377 = v1373 << (uint(int32(2)) % 32)
						if base.Ui32(v1373) <= base.Ui32(int32(23)) {
							v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+uint32(_c_F_des_setkey[52])))
							*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1382
							v1392 = v1382
							v1393 = v1362
						} else {
							v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1377+int32(_a_F_des_setkey_15)-int32(96))))
							*(*int32)(unsafe.Add(mBase, uint32(v1366))) = v1386
							v1389 = v1386
							v1392 = int32(0)
							v1393 = v1389
						}
					}
				}
				if v1203 == int32(0) {
					v1416 = v1392
					v1417 = v1393
				} else {
					v1399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_des_setkey[53]))))
					if v1399 == int32(255) {
						v1416 = v1392
						v1417 = v1393
					} else {
						v1403 = v1399 << (uint(int32(2)) % 32)
						if base.Ui32(int32(24)) <= base.Ui32(v1399) {
							v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1403+int32(_a_F_des_setkey_15)-int32(96))))
							v1411 = v1393 | v1410
							*(*int32)(unsafe.Add(mBase, uint32(v1366))) = v1411
							v1416 = v1392
							v1417 = v1411
						} else {
							v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+uint32(_c_F_des_setkey[52])))
							v1414 = v1392 | v1413
							*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1414
							v1416 = v1414
							v1417 = v1393
						}
					}
				}
				if v1230 == int32(0) {
					v1441 = v1416
					v1442 = v1417
				} else {
					v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_des_setkey[54]))))
					if v1424 == int32(255) {
						v1441 = v1416
						v1442 = v1417
					} else {
						v1428 = v1424 << (uint(int32(2)) % 32)
						if base.Ui32(int32(24)) <= base.Ui32(v1424) {
							v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1428+int32(_a_F_des_setkey_15)-int32(96))))
							v1436 = v1417 | v1435
							*(*int32)(unsafe.Add(mBase, uint32(v1366))) = v1436
							v1441 = v1416
							v1442 = v1436
						} else {
							v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+uint32(_c_F_des_setkey[52])))
							v1439 = v1416 | v1438
							*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1439
							v1441 = v1439
							v1442 = v1417
						}
					}
				}
				if v1257 == int32(0) {
					v1466 = v1441
					v1467 = v1442
				} else {
					v1449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_des_setkey[55]))))
					if v1449 == int32(255) {
						v1466 = v1441
						v1467 = v1442
					} else {
						v1453 = v1449 << (uint(int32(2)) % 32)
						if base.Ui32(int32(24)) <= base.Ui32(v1449) {
							v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1453+int32(_a_F_des_setkey_15)-int32(96))))
							v1461 = v1442 | v1460
							*(*int32)(unsafe.Add(mBase, uint32(v1366))) = v1461
							v1466 = v1441
							v1467 = v1461
						} else {
							v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1453)+uint32(_c_F_des_setkey[52])))
							v1464 = v1441 | v1463
							*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1464
							v1466 = v1464
							v1467 = v1442
						}
					}
				}
				if v1284 == int32(0) {
					v1491 = v1466
					v1492 = v1467
				} else {
					v1474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_des_setkey[56]))))
					if v1474 == int32(255) {
						v1491 = v1466
						v1492 = v1467
					} else {
						v1478 = v1474 << (uint(int32(2)) % 32)
						if base.Ui32(int32(24)) <= base.Ui32(v1474) {
							v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1478+int32(_a_F_des_setkey_15)-int32(96))))
							v1486 = v1467 | v1485
							*(*int32)(unsafe.Add(mBase, uint32(v1366))) = v1486
							v1491 = v1466
							v1492 = v1486
						} else {
							v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+uint32(_c_F_des_setkey[52])))
							v1489 = v1466 | v1488
							*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1489
							v1491 = v1489
							v1492 = v1467
						}
					}
				}
				if v1311 == int32(0) {
					v1516 = v1491
					v1517 = v1492
				} else {
					v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_des_setkey[57]))))
					if v1499 == int32(255) {
						v1516 = v1491
						v1517 = v1492
					} else {
						v1503 = v1499 << (uint(int32(2)) % 32)
						if base.Ui32(int32(24)) <= base.Ui32(v1499) {
							v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1503+int32(_a_F_des_setkey_15)-int32(96))))
							v1511 = v1492 | v1510
							*(*int32)(unsafe.Add(mBase, uint32(v1366))) = v1511
							v1516 = v1491
							v1517 = v1511
						} else {
							v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+uint32(_c_F_des_setkey[52])))
							v1514 = v1491 | v1513
							*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1514
							v1516 = v1514
							v1517 = v1492
						}
					}
				}
				if v1338 == int32(0) {
				} else {
					v1524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_des_setkey[58]))))
					if v1524 == int32(255) {
					} else {
						v1528 = v1524 << (uint(int32(2)) % 32)
						if base.Ui32(int32(24)) <= base.Ui32(v1524) {
							v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1528+int32(_a_F_des_setkey_15)-int32(96))))
							*(*int32)(unsafe.Add(mBase, uint32(v1366))) = v1517 | v1535
						} else {
							v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1528)+uint32(_c_F_des_setkey[52])))
							*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1516 | v1538
						}
					}
				}
				v1544 = v1145 + int32(1)
				if v1544 != int32(128) {
					v1145 = v1544
					continue
				} else {
					break
				}
				break
			}
			v1548 = v986 + int32(1)
			if v1548 != int32(8) {
				v986 = v1548
				continue
			} else {
				break
			}
			break
		}
		v1552 = int32(0)
		for {
			v1576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1552)+uint32(_c_F_des_setkey[59]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v1576)+uint32(_c_F_des_setkey[60]))) = uint8(v1552)
			v1581 = v1552 | int32(1)
			v1584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1581)+uint32(_c_F_des_setkey[59]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v1584)+uint32(_c_F_des_setkey[60]))) = uint8(v1581)
			v1589 = v1552 | int32(2)
			v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589)+uint32(_c_F_des_setkey[59]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v1592)+uint32(_c_F_des_setkey[60]))) = uint8(v1589)
			v1597 = v1552 | int32(3)
			v1600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1597)+uint32(_c_F_des_setkey[59]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v1600)+uint32(_c_F_des_setkey[60]))) = uint8(v1597)
			v1605 = v1552 + int32(4)
			if v1605 != int32(32) {
				v1552 = v1605
				continue
			} else {
				break
			}
			break
		}
		v1613 = int32(0)
		for {
			v1636 = v1613 << (uint(int32(3)) % 32)
			v1638 = int32(0)
			for {
				v1662 = v1613<<(uint(int32(10))%32) + int32(_a_F_des_setkey_16) + v1638<<(uint(int32(2))%32)
				v1663 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1663
				if v1638&int32(128) != 0 {
					v1670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1636)+uint32(_c_F_des_setkey[61]))))
					v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1670<<(uint(int32(2))%32))+uint32(_c_F_des_setkey[44])))
					*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1675
					v1677 = v1675
				} else {
					v1677 = v1663
				}
				if v1638&int32(64) != 0 {
					v1682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1636)+uint32(_c_F_des_setkey[62]))))
					v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1682<<(uint(int32(2))%32))+uint32(_c_F_des_setkey[44])))
					v1688 = v1677 | v1687
					*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1688
					v1690 = v1688
				} else {
					v1690 = v1677
				}
				if v1638&int32(32) != 0 {
					v1695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1636)+uint32(_c_F_des_setkey[63]))))
					v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1695<<(uint(int32(2))%32))+uint32(_c_F_des_setkey[44])))
					v1701 = v1690 | v1700
					*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1701
					v1703 = v1701
				} else {
					v1703 = v1690
				}
				if v1638&int32(16) != 0 {
					v1708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1636)+uint32(_c_F_des_setkey[64]))))
					v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1708<<(uint(int32(2))%32))+uint32(_c_F_des_setkey[44])))
					v1714 = v1703 | v1713
					*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1714
					v1716 = v1714
				} else {
					v1716 = v1703
				}
				if v1638&int32(8) != 0 {
					v1721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1636)+uint32(_c_F_des_setkey[65]))))
					v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1721<<(uint(int32(2))%32))+uint32(_c_F_des_setkey[44])))
					v1727 = v1716 | v1726
					*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1727
					v1729 = v1727
				} else {
					v1729 = v1716
				}
				if v1638&int32(4) != 0 {
					v1734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1636)+uint32(_c_F_des_setkey[66]))))
					v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1734<<(uint(int32(2))%32))+uint32(_c_F_des_setkey[44])))
					v1740 = v1729 | v1739
					*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1740
					v1742 = v1740
				} else {
					v1742 = v1729
				}
				if v1638&int32(2) != 0 {
					v1747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1636)+uint32(_c_F_des_setkey[67]))))
					v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1747<<(uint(int32(2))%32))+uint32(_c_F_des_setkey[44])))
					v1753 = v1742 | v1752
					*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1753
					v1755 = v1753
				} else {
					v1755 = v1742
				}
				if v1638&int32(1) != 0 {
					v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1636)+uint32(_c_F_des_setkey[68]))))
					v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1760<<(uint(int32(2))%32))+uint32(_c_F_des_setkey[44])))
					*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1755 | v1765
				} else {
				}
				v1769 = v1638 + int32(1)
				if v1769 != int32(256) {
					v1638 = v1769
					continue
				} else {
					break
				}
				break
			}
			v1773 = v1613 + int32(1)
			if v1773 != int32(4) {
				v1613 = v1773
				continue
			} else {
				break
			}
			break
		}
		v1777 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_des_setkey[0])) = uint8(v1777)
	} else {
	}
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1780 = int32(24)
	v1782 = int32(16711935)
	v1786 = int32(8)
	v1788 = base.I32_rotr(v1779, v1780)&v1782 | base.I32_rotr(v1779&v1782, v1786)
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1798 = base.I32_rotr(v1789, v1780)&v1782 | base.I32_rotr(v1789&v1782, v1786)
	if v1788|v1798 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[2])) = v1788
		*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[1])) = v1798
		v1812 = int32(7)
		v1814 = int32(508)
		v1815 = int32(base.Ui32(v1788)>>(uint(v1812)%32)) & v1814
		v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+uint32(_c_F_des_setkey[69])))
		v1819 = int32(15)
		v1822 = int32(base.Ui32(v1788)>>(uint(v1819)%32)) & v1814
		v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+uint32(_c_F_des_setkey[70])))
		v1826 = int32(23)
		v1829 = int32(base.Ui32(v1788)>>(uint(v1826)%32)) & v1814
		v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+uint32(_c_F_des_setkey[71])))
		v1836 = int32(base.Ui32(v1798)>>(uint(v1812)%32)) & v1814
		v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+uint32(_c_F_des_setkey[72])))
		v1843 = int32(base.Ui32(v1798)>>(uint(v1819)%32)) & v1814
		v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+uint32(_c_F_des_setkey[73])))
		v1850 = int32(base.Ui32(v1798)>>(uint(v1826)%32)) & v1814
		v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1850)+uint32(_c_F_des_setkey[74])))
		v1856 = int32(1)
		v1859 = v1798 << (uint(v1856) % 32) & v1814
		v1862 = *(*int32)(unsafe.Add(mBase, uint32(v1859)+uint32(_c_F_des_setkey[75])))
		v1870 = v1788 << (uint(v1856) % 32) & v1814
		v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1870)+uint32(_c_F_des_setkey[76])))
		v1874 = v1818 | (v1825 | (v1832 | (v1839 | (v1846 | v1853) | v1862))) | v1873
		v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1870)+uint32(_c_F_des_setkey[77])))
		v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+uint32(_c_F_des_setkey[78])))
		v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+uint32(_c_F_des_setkey[79])))
		v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+uint32(_c_F_des_setkey[80])))
		v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1859)+uint32(_c_F_des_setkey[81])))
		v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+uint32(_c_F_des_setkey[82])))
		v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+uint32(_c_F_des_setkey[83])))
		v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1850)+uint32(_c_F_des_setkey[84])))
		v1905 = v1877 | (v1880 | (v1883 | (v1886 | (v1889 | (v1892 | (v1895 | v1898))))))
		v1907 = int32(0)
		v1915 = v2
		for {
			v1922 = int32(2)
			v1923 = v1907 << (uint(v1922) % 32)
			v1929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1907)+uint32(_c_F_des_setkey[85]))))
			v1930 = v1915 + v1929
			v1931 = int32(28) - v1930
			v1934 = int32(base.Ui32(v1905)>>(uint(v1931)%32)) | v1905<<(uint(v1930)%32)
			v1935 = int32(12)
			v1937 = int32(508)
			v1938 = int32(base.Ui32(v1934)>>(uint(v1935)%32)) & v1937
			v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+uint32(_c_F_des_setkey[86])))
			v1942 = int32(19)
			v1945 = int32(base.Ui32(v1934)>>(uint(v1942)%32)) & v1937
			v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+uint32(_c_F_des_setkey[87])))
			v1950 = int32(5)
			v1953 = int32(base.Ui32(v1934)>>(uint(v1950)%32)) & v1937
			v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+uint32(_c_F_des_setkey[88])))
			v1958 = int32(127)
			v1961 = v1934 & v1958 << (uint(v1922) % 32)
			v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+uint32(_c_F_des_setkey[89])))
			v1968 = v1874<<(uint(v1930)%32) | int32(base.Ui32(v1874)>>(uint(v1931)%32))
			v1972 = int32(base.Ui32(v1968)>>(uint(v1942)%32)) & v1937
			v1975 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+uint32(_c_F_des_setkey[90])))
			v1980 = int32(base.Ui32(v1968)>>(uint(v1935)%32)) & v1937
			v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+uint32(_c_F_des_setkey[91])))
			v1988 = int32(base.Ui32(v1968)>>(uint(v1950)%32)) & v1937
			v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1988)+uint32(_c_F_des_setkey[92])))
			v1996 = v1968 & v1958 << (uint(v1922) % 32)
			v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+uint32(_c_F_des_setkey[93])))
			v2000 = v1941 | v1948 | v1956 | v1964 | v1975 | v1983 | v1991 | v1999
			*(*int32)(unsafe.Add(mBase, uint32(v1923)+uint32(_c_F_des_setkey[94]))) = v2000
			v2005 = (int32(15) - v1907) << (uint(v1922) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v2005)+uint32(_c_F_des_setkey[95]))) = v2000
			v2013 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+uint32(_c_F_des_setkey[96])))
			v2016 = *(*int32)(unsafe.Add(mBase, uint32(v1988)+uint32(_c_F_des_setkey[97])))
			v2019 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+uint32(_c_F_des_setkey[98])))
			v2022 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+uint32(_c_F_des_setkey[99])))
			v2025 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+uint32(_c_F_des_setkey[100])))
			v2028 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+uint32(_c_F_des_setkey[101])))
			v2031 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+uint32(_c_F_des_setkey[102])))
			v2034 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+uint32(_c_F_des_setkey[103])))
			v2041 = v2013 | (v2016 | (v2019 | (v2022 | (v2025 | (v2028 | (v2031 | v2034))))))
			*(*int32)(unsafe.Add(mBase, uint32(v2005)+uint32(_c_F_des_setkey[104]))) = v2041
			*(*int32)(unsafe.Add(mBase, uint32(v1923)+uint32(_c_F_des_setkey[105]))) = v2041
			v2047 = v1907 + int32(1)
			if v2047 != int32(16) {
				v1907 = v2047
				v1915 = v1930
				continue
			} else {
				break
			}
			break
		}
	} else {
		v1803 = *(*int32)(unsafe.Add(mBase, _c_F_des_setkey[1]))
		if v1798 != v1803 {
			*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[2])) = v1788
			*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[1])) = v1798
			v1812 = int32(7)
			v1814 = int32(508)
			v1815 = int32(base.Ui32(v1788)>>(uint(v1812)%32)) & v1814
			v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+uint32(_c_F_des_setkey[69])))
			v1819 = int32(15)
			v1822 = int32(base.Ui32(v1788)>>(uint(v1819)%32)) & v1814
			v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+uint32(_c_F_des_setkey[70])))
			v1826 = int32(23)
			v1829 = int32(base.Ui32(v1788)>>(uint(v1826)%32)) & v1814
			v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+uint32(_c_F_des_setkey[71])))
			v1836 = int32(base.Ui32(v1798)>>(uint(v1812)%32)) & v1814
			v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+uint32(_c_F_des_setkey[72])))
			v1843 = int32(base.Ui32(v1798)>>(uint(v1819)%32)) & v1814
			v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+uint32(_c_F_des_setkey[73])))
			v1850 = int32(base.Ui32(v1798)>>(uint(v1826)%32)) & v1814
			v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1850)+uint32(_c_F_des_setkey[74])))
			v1856 = int32(1)
			v1859 = v1798 << (uint(v1856) % 32) & v1814
			v1862 = *(*int32)(unsafe.Add(mBase, uint32(v1859)+uint32(_c_F_des_setkey[75])))
			v1870 = v1788 << (uint(v1856) % 32) & v1814
			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1870)+uint32(_c_F_des_setkey[76])))
			v1874 = v1818 | (v1825 | (v1832 | (v1839 | (v1846 | v1853) | v1862))) | v1873
			v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1870)+uint32(_c_F_des_setkey[77])))
			v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+uint32(_c_F_des_setkey[78])))
			v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+uint32(_c_F_des_setkey[79])))
			v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+uint32(_c_F_des_setkey[80])))
			v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1859)+uint32(_c_F_des_setkey[81])))
			v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+uint32(_c_F_des_setkey[82])))
			v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+uint32(_c_F_des_setkey[83])))
			v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1850)+uint32(_c_F_des_setkey[84])))
			v1905 = v1877 | (v1880 | (v1883 | (v1886 | (v1889 | (v1892 | (v1895 | v1898))))))
			v1907 = int32(0)
			v1915 = v2
			for {
				v1922 = int32(2)
				v1923 = v1907 << (uint(v1922) % 32)
				v1929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1907)+uint32(_c_F_des_setkey[85]))))
				v1930 = v1915 + v1929
				v1931 = int32(28) - v1930
				v1934 = int32(base.Ui32(v1905)>>(uint(v1931)%32)) | v1905<<(uint(v1930)%32)
				v1935 = int32(12)
				v1937 = int32(508)
				v1938 = int32(base.Ui32(v1934)>>(uint(v1935)%32)) & v1937
				v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+uint32(_c_F_des_setkey[86])))
				v1942 = int32(19)
				v1945 = int32(base.Ui32(v1934)>>(uint(v1942)%32)) & v1937
				v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+uint32(_c_F_des_setkey[87])))
				v1950 = int32(5)
				v1953 = int32(base.Ui32(v1934)>>(uint(v1950)%32)) & v1937
				v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+uint32(_c_F_des_setkey[88])))
				v1958 = int32(127)
				v1961 = v1934 & v1958 << (uint(v1922) % 32)
				v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+uint32(_c_F_des_setkey[89])))
				v1968 = v1874<<(uint(v1930)%32) | int32(base.Ui32(v1874)>>(uint(v1931)%32))
				v1972 = int32(base.Ui32(v1968)>>(uint(v1942)%32)) & v1937
				v1975 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+uint32(_c_F_des_setkey[90])))
				v1980 = int32(base.Ui32(v1968)>>(uint(v1935)%32)) & v1937
				v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+uint32(_c_F_des_setkey[91])))
				v1988 = int32(base.Ui32(v1968)>>(uint(v1950)%32)) & v1937
				v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1988)+uint32(_c_F_des_setkey[92])))
				v1996 = v1968 & v1958 << (uint(v1922) % 32)
				v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+uint32(_c_F_des_setkey[93])))
				v2000 = v1941 | v1948 | v1956 | v1964 | v1975 | v1983 | v1991 | v1999
				*(*int32)(unsafe.Add(mBase, uint32(v1923)+uint32(_c_F_des_setkey[94]))) = v2000
				v2005 = (int32(15) - v1907) << (uint(v1922) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(v2005)+uint32(_c_F_des_setkey[95]))) = v2000
				v2013 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+uint32(_c_F_des_setkey[96])))
				v2016 = *(*int32)(unsafe.Add(mBase, uint32(v1988)+uint32(_c_F_des_setkey[97])))
				v2019 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+uint32(_c_F_des_setkey[98])))
				v2022 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+uint32(_c_F_des_setkey[99])))
				v2025 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+uint32(_c_F_des_setkey[100])))
				v2028 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+uint32(_c_F_des_setkey[101])))
				v2031 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+uint32(_c_F_des_setkey[102])))
				v2034 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+uint32(_c_F_des_setkey[103])))
				v2041 = v2013 | (v2016 | (v2019 | (v2022 | (v2025 | (v2028 | (v2031 | v2034))))))
				*(*int32)(unsafe.Add(mBase, uint32(v2005)+uint32(_c_F_des_setkey[104]))) = v2041
				*(*int32)(unsafe.Add(mBase, uint32(v1923)+uint32(_c_F_des_setkey[105]))) = v2041
				v2047 = v1907 + int32(1)
				if v2047 != int32(16) {
					v1907 = v2047
					v1915 = v1930
					continue
				} else {
					break
				}
				break
			}
		} else {
			v1806 = *(*int32)(unsafe.Add(mBase, _c_F_des_setkey[2]))
			if v1788 == v1806 {
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[2])) = v1788
				*(*int32)(unsafe.Add(mBase, _c_F_des_setkey[1])) = v1798
				v1812 = int32(7)
				v1814 = int32(508)
				v1815 = int32(base.Ui32(v1788)>>(uint(v1812)%32)) & v1814
				v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+uint32(_c_F_des_setkey[69])))
				v1819 = int32(15)
				v1822 = int32(base.Ui32(v1788)>>(uint(v1819)%32)) & v1814
				v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+uint32(_c_F_des_setkey[70])))
				v1826 = int32(23)
				v1829 = int32(base.Ui32(v1788)>>(uint(v1826)%32)) & v1814
				v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+uint32(_c_F_des_setkey[71])))
				v1836 = int32(base.Ui32(v1798)>>(uint(v1812)%32)) & v1814
				v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+uint32(_c_F_des_setkey[72])))
				v1843 = int32(base.Ui32(v1798)>>(uint(v1819)%32)) & v1814
				v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+uint32(_c_F_des_setkey[73])))
				v1850 = int32(base.Ui32(v1798)>>(uint(v1826)%32)) & v1814
				v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1850)+uint32(_c_F_des_setkey[74])))
				v1856 = int32(1)
				v1859 = v1798 << (uint(v1856) % 32) & v1814
				v1862 = *(*int32)(unsafe.Add(mBase, uint32(v1859)+uint32(_c_F_des_setkey[75])))
				v1870 = v1788 << (uint(v1856) % 32) & v1814
				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1870)+uint32(_c_F_des_setkey[76])))
				v1874 = v1818 | (v1825 | (v1832 | (v1839 | (v1846 | v1853) | v1862))) | v1873
				v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1870)+uint32(_c_F_des_setkey[77])))
				v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+uint32(_c_F_des_setkey[78])))
				v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+uint32(_c_F_des_setkey[79])))
				v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+uint32(_c_F_des_setkey[80])))
				v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1859)+uint32(_c_F_des_setkey[81])))
				v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+uint32(_c_F_des_setkey[82])))
				v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+uint32(_c_F_des_setkey[83])))
				v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1850)+uint32(_c_F_des_setkey[84])))
				v1905 = v1877 | (v1880 | (v1883 | (v1886 | (v1889 | (v1892 | (v1895 | v1898))))))
				v1907 = int32(0)
				v1915 = v2
				for {
					v1922 = int32(2)
					v1923 = v1907 << (uint(v1922) % 32)
					v1929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1907)+uint32(_c_F_des_setkey[85]))))
					v1930 = v1915 + v1929
					v1931 = int32(28) - v1930
					v1934 = int32(base.Ui32(v1905)>>(uint(v1931)%32)) | v1905<<(uint(v1930)%32)
					v1935 = int32(12)
					v1937 = int32(508)
					v1938 = int32(base.Ui32(v1934)>>(uint(v1935)%32)) & v1937
					v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+uint32(_c_F_des_setkey[86])))
					v1942 = int32(19)
					v1945 = int32(base.Ui32(v1934)>>(uint(v1942)%32)) & v1937
					v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+uint32(_c_F_des_setkey[87])))
					v1950 = int32(5)
					v1953 = int32(base.Ui32(v1934)>>(uint(v1950)%32)) & v1937
					v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+uint32(_c_F_des_setkey[88])))
					v1958 = int32(127)
					v1961 = v1934 & v1958 << (uint(v1922) % 32)
					v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+uint32(_c_F_des_setkey[89])))
					v1968 = v1874<<(uint(v1930)%32) | int32(base.Ui32(v1874)>>(uint(v1931)%32))
					v1972 = int32(base.Ui32(v1968)>>(uint(v1942)%32)) & v1937
					v1975 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+uint32(_c_F_des_setkey[90])))
					v1980 = int32(base.Ui32(v1968)>>(uint(v1935)%32)) & v1937
					v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+uint32(_c_F_des_setkey[91])))
					v1988 = int32(base.Ui32(v1968)>>(uint(v1950)%32)) & v1937
					v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1988)+uint32(_c_F_des_setkey[92])))
					v1996 = v1968 & v1958 << (uint(v1922) % 32)
					v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+uint32(_c_F_des_setkey[93])))
					v2000 = v1941 | v1948 | v1956 | v1964 | v1975 | v1983 | v1991 | v1999
					*(*int32)(unsafe.Add(mBase, uint32(v1923)+uint32(_c_F_des_setkey[94]))) = v2000
					v2005 = (int32(15) - v1907) << (uint(v1922) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(v2005)+uint32(_c_F_des_setkey[95]))) = v2000
					v2013 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+uint32(_c_F_des_setkey[96])))
					v2016 = *(*int32)(unsafe.Add(mBase, uint32(v1988)+uint32(_c_F_des_setkey[97])))
					v2019 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+uint32(_c_F_des_setkey[98])))
					v2022 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+uint32(_c_F_des_setkey[99])))
					v2025 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+uint32(_c_F_des_setkey[100])))
					v2028 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+uint32(_c_F_des_setkey[101])))
					v2031 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+uint32(_c_F_des_setkey[102])))
					v2034 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+uint32(_c_F_des_setkey[103])))
					v2041 = v2013 | (v2016 | (v2019 | (v2022 | (v2025 | (v2028 | (v2031 | v2034))))))
					*(*int32)(unsafe.Add(mBase, uint32(v2005)+uint32(_c_F_des_setkey[104]))) = v2041
					*(*int32)(unsafe.Add(mBase, uint32(v1923)+uint32(_c_F_des_setkey[105]))) = v2041
					v2047 = v1907 + int32(1)
					if v2047 != int32(16) {
						v1907 = v2047
						v1915 = v1930
						continue
					} else {
						break
					}
					break
				}
			}
		}
	}
	return
}
func F_die(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_die[0])))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_die[1])) = int32(4)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_die[2]))
	v24 = int32(0)
	v27 = base.AtomicRmwOr32(m, v24, int32(_a_F_die_0), v24)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v6 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_die[3])) = v6
	*(*int32)(unsafe.Add(mBase, _c_F_die[4])) = v6
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_die[5]))
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_die[5])) = v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_die[6])) = v17
	goto L1
L4:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_die[7])))
	if v76 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
	v31 = int32(0)
	v34 = base.AtomicRmwOr32(m, v31, int32(_a_F_die_0), v31)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v35 == v31 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v38 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_die[8]))
	if v42 == v38 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = m.G0
	v46 = v44 - int32(16)
	m.G0 = v46
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_die[9]))
	if v49 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v72 = F_pgmem_kill(m, v38, int32(23))
	mBase = m.M
	goto L5
L12:
	;
	m.G0 = v46 + int32(16)
	goto L4
L13:
	;
	v52 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+15)) = uint8(v52)
	goto L14
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_die[10]))
	v60 = F_write(m, v56, v46+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v60 {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_die[11]))
	if v64 == int32(27) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_die[12]))
	if v80 == int32(2) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	goto L18
}
func F_digest_finish(m *base.Module, l0 int32, l1 int32) {
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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = m.Env.Pgmem_hash_final(m, v4, l1, v5)
	mBase = m.M
	if v6 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_digest_finish_0), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_digest_finish_1), int32(204), int32(_a_F_digest_finish_2))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		return
	}
}
func F_digest_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = m.Env.Pgmem_hash_reset(m, v3)
	mBase = m.M
	if v4 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_digest_reset_0), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_digest_reset_1), int32(186), int32(_a_F_digest_reset_2))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		return
	}
}
func F_digest_result_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	return v3
}
func F_digest_update(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = m.Env.Pgmem_hash_update(m, v5, l1, l2)
	mBase = m.M
	if v6 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_digest_update_0), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_digest_update_1), int32(195), int32(_a_F_digest_update_2))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		return
	}
}
func F_distance_taxicab(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 float64
	_ = v53
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v88 int32
	_ = v88
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v111 int32
	_ = v111
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	var v134 int32
	_ = v134
	var v141 float64
	_ = v141
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v167 float64
	_ = v167
	var v168 int32
	_ = v168
	var v183 int32
	_ = v183
	var v184 float64
	_ = v184
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v194 int32
	_ = v194
	var v204 float64
	_ = v204
	var v206 float64
	_ = v206
	var v209 int32
	_ = v209
	var v217 float64
	_ = v217
	var v218 float64
	_ = v218
	var v219 float64
	_ = v219
	var v221 int32
	_ = v221
	var v228 float64
	_ = v228
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	v2 = float64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v26 = F_pg_detoast_datum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			v29 = int32(2147483647)
			v30 = v28 & v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			v33 = v31 & v29
			v34 = base.B2i32(base.Ui32(v30) < base.Ui32(v33))
			if base.Ui32(v30) < base.Ui32(v33) {
				v35 = v26
			} else {
				v35 = v21
			}
			if base.Ui32(v30) < base.Ui32(v33) {
				v36 = v21
			} else {
				v36 = v26
			}
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
			v39 = v37 & int32(2147483647)
			if v39 == int32(0) {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v141 = v2
				v143 = v42
			} else {
				v43 = int32(8)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v53 = v2
				v58 = int32(0)
				for {
					v68 = v58 << (uint(int32(3)) % 32)
					v69 = v35 + v43 + v68
					v70 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
					v71 = v68 + (v36 + v43)
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v71)))
					if int32(0) <= v47 {
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v69+v47<<(uint(int32(3))%32))))
						v79 = v78
					} else {
						v79 = v70
					}
					if int32(0) <= v37 {
						v85 = *(*float64)(unsafe.Add(mBase, uint32(v71+v37<<(uint(int32(3))%32))))
						v86 = v85
					} else {
						v86 = v72
					}
					v88 = int32(0)
					if base.B2i32(base.F64_le(v79, v86) == v88)|base.B2i32(base.F64_le(v70, v72) == v88)|(base.B2i32(base.F64_le(v79, v72) == v88)|base.B2i32(base.F64_ge(v86, v70) == v88)) == v88 {
						if base.F64_gt(v86, v72) != 0 {
							v105 = v72
						} else {
							v105 = v86
						}
						if base.F64_lt(v79, v70) != 0 {
							v107 = v70
						} else {
							v107 = v79
						}
						v130 = base.F64_sub(v105, v107)
					} else {
						v111 = int32(0)
						if base.B2i32(base.F64_gt(v79, v86) == v111)|base.B2i32(base.F64_gt(v70, v72) == v111)|(base.B2i32(base.F64_gt(v79, v72) == v111)|base.B2i32(base.F64_lt(v86, v70) == v111)) != 0 {
							v130 = float64(0)
						} else {
							if base.F64_gt(v79, v70) != 0 {
								v126 = v70
							} else {
								v126 = v79
							}
							if base.F64_lt(v86, v72) != 0 {
								v128 = v72
							} else {
								v128 = v86
							}
							v130 = base.F64_sub(v126, v128)
						}
					}
					v132 = base.F64_add(v53, base.F64_abs(v130))
					v134 = v58 + int32(1)
					if v134 != v39 {
						v53 = v132
						v58 = v134
						continue
					} else {
						break
					}
					break
				}
				v141 = v132
				v143 = v47
			}
			v156 = v143 & int32(2147483647)
			if base.Ui32(v39) < base.Ui32(v156) {
				v167 = v141
				v168 = v39
				for {
					v183 = v35 + int32(8) + v168<<(uint(int32(3))%32)
					v184 = *(*float64)(unsafe.Add(mBase, uint32(v183)))
					if base.B2i32(v143 < int32(0)) == int32(0) {
						v190 = *(*float64)(unsafe.Add(mBase, uint32(v183+v156<<(uint(int32(3))%32))))
						v191 = v190
					} else {
						v191 = v184
					}
					v192 = float64(0)
					v194 = int32(0)
					if base.B2i32(base.F64_le(v184, v192) == v194)|base.B2i32(base.F64_le(v191, v192) == v194) == v194 {
						if base.F64_lt(v191, v184) != 0 {
							v204 = v184
						} else {
							v204 = v191
						}
						v218 = base.F64_abs(v204)
					} else {
						v206 = float64(0)
						v209 = int32(0)
						if base.B2i32(base.F64_gt(v184, v206) == v209)|base.B2i32(base.F64_gt(v191, v206) == v209) != 0 {
							v218 = v206
						} else {
							if base.F64_gt(v191, v184) != 0 {
								v217 = v184
							} else {
								v217 = v191
							}
							v218 = v217
						}
					}
					v219 = base.F64_add(v167, v218)
					v221 = v168 + int32(1)
					if v221 != v156 {
						v167 = v219
						v168 = v221
						continue
					} else {
						break
					}
					break
				}
				v228 = v219
			} else {
				v228 = v141
			}
			v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if base.Ui32(v30) < base.Ui32(v33) {
				if v242 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v246 = m.ExcPending
					if v246 != 0 {
						return int64(0)
					} else {
						v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v26 != v247 {
							F_pfree(m, v26)
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return int64(0)
							} else {
								return base.I64_reinterpret_f64(v228)
							}
						} else {
							return base.I64_reinterpret_f64(v228)
						}
					}
				} else {
					v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v26 != v247 {
						F_pfree(m, v26)
						mBase = m.M
						v255 = m.ExcPending
						if v255 != 0 {
							return int64(0)
						} else {
							return base.I64_reinterpret_f64(v228)
						}
					} else {
						return base.I64_reinterpret_f64(v228)
					}
				}
			} else {
				if v242 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v251 = m.ExcPending
					if v251 != 0 {
						return int64(0)
					} else {
						v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v26 == v252 {
							return base.I64_reinterpret_f64(v228)
						} else {
							F_pfree(m, v26)
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return int64(0)
							} else {
								return base.I64_reinterpret_f64(v228)
							}
						}
					}
				} else {
					v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v26 == v252 {
						return base.I64_reinterpret_f64(v228)
					} else {
						F_pfree(m, v26)
						mBase = m.M
						v255 = m.ExcPending
						if v255 != 0 {
							return int64(0)
						} else {
							return base.I64_reinterpret_f64(v228)
						}
					}
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	v5 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v18 = F_palloc(m, l1+int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if int32(0) < l1 {
			if l1 != int32(1) {
				v44 = v5
				v48 = v5
				for {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v44))))
					if base.Ui32((v51-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v60 = v51 | int32(32)
					} else {
						v60 = v51
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v18+v44))) = uint8(v60)
					v63 = v44 | int32(1)
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v63))))
					if base.Ui32((v66-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v75 = v66 | int32(32)
					} else {
						v75 = v66
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v18+v63))) = uint8(v75)
					v77 = int32(2)
					v78 = v44 + v77
					v80 = v48 + v77
					if v80 != l1&int32(2147483646) {
						v44 = v78
						v48 = v80
						continue
					} else {
						break
					}
					break
				}
				if l1&int32(1) == int32(0) {
				} else {
					v90 = v78
					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v90))))
					if base.Ui32((v97-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v106 = v97 | int32(32)
					} else {
						v106 = v97
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v18+v90))) = uint8(v106)
				}
			} else {
				v90 = v5
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v90))))
				if base.Ui32((v97-int32(65))&int32(255)) < base.Ui32(int32(26)) {
					v106 = v97 | int32(32)
				} else {
					v106 = v97
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v18+v90))) = uint8(v106)
			}
			v120 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1+v18))) = uint8(v120)
			if base.B2i32(l3 == v120)|base.B2i32(base.Ui32(l1) < base.Ui32(int32(64))) != 0 {
				m.G0 = v14 + int32(16)
				return v18
			} else {
				v128 = F_pg_mbcliplen(m, v18, l1, int32(63))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int32(0)
				} else {
					if l2 == int32(0) {
						v153 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v128+v18))) = uint8(v153)
						m.G0 = v14 + int32(16)
						return v18
					} else {
						v134 = F_errstart(m, int32(18), int32(0))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return int32(0)
						} else {
							if v134 == int32(0) {
								v153 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v128+v18))) = uint8(v153)
								m.G0 = v14 + int32(16)
								return v18
							} else {
								F_errcode(m, int32(34103428))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v18
									*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = v18
									F_errmsg(m, int32(_a_F_downcase_identifier_0), v14)
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_downcase_identifier_1), int32(102), int32(_a_F_downcase_identifier_2))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											v153 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v128+v18))) = uint8(v153)
											m.G0 = v14 + int32(16)
											return v18
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v155 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v155)
			m.G0 = v14 + int32(16)
			return v18
		}
	}
}
func F_dsnowball_lexize(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v10 = F_str_tolower(m, v7, v8, int32(100))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v16 = F_palloc0_mul(m, int32(8), int32(2))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			if int32(1001) <= v8 {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v10
				return base.I64_extend_i32_u(v16)
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				if v23 != 0 {
					v26 = F_searchstoplist(m, v6+int32(4), v10)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						if v26 == int32(0) {
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
							if v34 != int32(1) {
								v44 = v10
								v45 = int32(_a_F_dsnowball_lexize_0)
								v46 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0]))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v48
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
								v51 = F_strlen(m, v44)
								mBase = m.M
								v52 = F_SN_set_current(m, v50, v51, v44)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
									v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v54)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v46
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
										if v61 == int32(0) {
											v81 = v44
											v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
											if v83 != int32(1) {
												v93 = v81
												*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
												return base.I64_extend_i32_u(v16)
											} else {
												v86 = F_strlen(m, v81)
												mBase = m.M
												v88 = F_pg_any_to_server(m, v81, v86, int32(6))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int64(0)
												} else {
													if v81 == v88 {
														v93 = v81
														*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
														return base.I64_extend_i32_u(v16)
													} else {
														F_pfree(m, v81)
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int64(0)
														} else {
															v93 = v88
															*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
															return base.I64_extend_i32_u(v16)
														}
													}
												}
											}
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
											if v64 == int32(0) {
												v81 = v44
												v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
												if v83 != int32(1) {
													v93 = v81
													*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
													return base.I64_extend_i32_u(v16)
												} else {
													v86 = F_strlen(m, v81)
													mBase = m.M
													v88 = F_pg_any_to_server(m, v81, v86, int32(6))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int64(0)
													} else {
														if v81 == v88 {
															v93 = v81
															*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
															return base.I64_extend_i32_u(v16)
														} else {
															F_pfree(m, v81)
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int64(0)
															} else {
																v93 = v88
																*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																return base.I64_extend_i32_u(v16)
															}
														}
													}
												}
											} else {
												v69 = F_repalloc(m, v44, v64+int32(1))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int64(0)
												} else {
													v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
													v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
													if v72 != 0 {
														v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
														base.MemoryCopy(m, v69, v73, v72)
													} else {
													}
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
													v78 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v69+v76))) = uint8(v78)
													v81 = v69
													v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
													if v83 != int32(1) {
														v93 = v81
														*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
														return base.I64_extend_i32_u(v16)
													} else {
														v86 = F_strlen(m, v81)
														mBase = m.M
														v88 = F_pg_any_to_server(m, v81, v86, int32(6))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int64(0)
														} else {
															if v81 == v88 {
																v93 = v81
																*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																return base.I64_extend_i32_u(v16)
															} else {
																F_pfree(m, v81)
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return int64(0)
																} else {
																	v93 = v88
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																	return base.I64_extend_i32_u(v16)
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
								v37 = F_strlen(m, v10)
								mBase = m.M
								v39 = F_pg_server_to_any(m, v10, v37, int32(6))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int64(0)
								} else {
									if v10 == v39 {
										v44 = v10
										v45 = int32(_a_F_dsnowball_lexize_0)
										v46 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0]))
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
										*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v48
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
										v51 = F_strlen(m, v44)
										mBase = m.M
										v52 = F_SN_set_current(m, v50, v51, v44)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int64(0)
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
											v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v54)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v46
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
												v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
												if v61 == int32(0) {
													v81 = v44
													v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
													if v83 != int32(1) {
														v93 = v81
														*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
														return base.I64_extend_i32_u(v16)
													} else {
														v86 = F_strlen(m, v81)
														mBase = m.M
														v88 = F_pg_any_to_server(m, v81, v86, int32(6))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int64(0)
														} else {
															if v81 == v88 {
																v93 = v81
																*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																return base.I64_extend_i32_u(v16)
															} else {
																F_pfree(m, v81)
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return int64(0)
																} else {
																	v93 = v88
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																	return base.I64_extend_i32_u(v16)
																}
															}
														}
													}
												} else {
													v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
													if v64 == int32(0) {
														v81 = v44
														v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
														if v83 != int32(1) {
															v93 = v81
															*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
															return base.I64_extend_i32_u(v16)
														} else {
															v86 = F_strlen(m, v81)
															mBase = m.M
															v88 = F_pg_any_to_server(m, v81, v86, int32(6))
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int64(0)
															} else {
																if v81 == v88 {
																	v93 = v81
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																	return base.I64_extend_i32_u(v16)
																} else {
																	F_pfree(m, v81)
																	mBase = m.M
																	v92 = m.ExcPending
																	if v92 != 0 {
																		return int64(0)
																	} else {
																		v93 = v88
																		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																		return base.I64_extend_i32_u(v16)
																	}
																}
															}
														}
													} else {
														v69 = F_repalloc(m, v44, v64+int32(1))
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int64(0)
														} else {
															v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
															if v72 != 0 {
																v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
																base.MemoryCopy(m, v69, v73, v72)
															} else {
															}
															v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
															v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
															v78 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v69+v76))) = uint8(v78)
															v81 = v69
															v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
															if v83 != int32(1) {
																v93 = v81
																*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																return base.I64_extend_i32_u(v16)
															} else {
																v86 = F_strlen(m, v81)
																mBase = m.M
																v88 = F_pg_any_to_server(m, v81, v86, int32(6))
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return int64(0)
																} else {
																	if v81 == v88 {
																		v93 = v81
																		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																		return base.I64_extend_i32_u(v16)
																	} else {
																		F_pfree(m, v81)
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return int64(0)
																		} else {
																			v93 = v88
																			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																			return base.I64_extend_i32_u(v16)
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
										F_pfree(m, v10)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return int64(0)
										} else {
											v44 = v39
											v45 = int32(_a_F_dsnowball_lexize_0)
											v46 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0]))
											v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
											*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v48
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
											v51 = F_strlen(m, v44)
											mBase = m.M
											v52 = F_SN_set_current(m, v50, v51, v44)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int64(0)
											} else {
												v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
												v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
												v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v54)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v46
													v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
													v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
													if v61 == int32(0) {
														v81 = v44
														v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
														if v83 != int32(1) {
															v93 = v81
															*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
															return base.I64_extend_i32_u(v16)
														} else {
															v86 = F_strlen(m, v81)
															mBase = m.M
															v88 = F_pg_any_to_server(m, v81, v86, int32(6))
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int64(0)
															} else {
																if v81 == v88 {
																	v93 = v81
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																	return base.I64_extend_i32_u(v16)
																} else {
																	F_pfree(m, v81)
																	mBase = m.M
																	v92 = m.ExcPending
																	if v92 != 0 {
																		return int64(0)
																	} else {
																		v93 = v88
																		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																		return base.I64_extend_i32_u(v16)
																	}
																}
															}
														}
													} else {
														v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
														if v64 == int32(0) {
															v81 = v44
															v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
															if v83 != int32(1) {
																v93 = v81
																*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																return base.I64_extend_i32_u(v16)
															} else {
																v86 = F_strlen(m, v81)
																mBase = m.M
																v88 = F_pg_any_to_server(m, v81, v86, int32(6))
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return int64(0)
																} else {
																	if v81 == v88 {
																		v93 = v81
																		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																		return base.I64_extend_i32_u(v16)
																	} else {
																		F_pfree(m, v81)
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return int64(0)
																		} else {
																			v93 = v88
																			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																			return base.I64_extend_i32_u(v16)
																		}
																	}
																}
															}
														} else {
															v69 = F_repalloc(m, v44, v64+int32(1))
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int64(0)
															} else {
																v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
																if v72 != 0 {
																	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
																	base.MemoryCopy(m, v69, v73, v72)
																} else {
																}
																v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
																v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
																v78 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v69+v76))) = uint8(v78)
																v81 = v69
																v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
																if v83 != int32(1) {
																	v93 = v81
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																	return base.I64_extend_i32_u(v16)
																} else {
																	v86 = F_strlen(m, v81)
																	mBase = m.M
																	v88 = F_pg_any_to_server(m, v81, v86, int32(6))
																	mBase = m.M
																	v89 = m.ExcPending
																	if v89 != 0 {
																		return int64(0)
																	} else {
																		if v81 == v88 {
																			v93 = v81
																			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																			return base.I64_extend_i32_u(v16)
																		} else {
																			F_pfree(m, v81)
																			mBase = m.M
																			v92 = m.ExcPending
																			if v92 != 0 {
																				return int64(0)
																			} else {
																				v93 = v88
																				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v93
																				return base.I64_extend_i32_u(v16)
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
							F_pfree(m, v10)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v16)
							}
						}
					}
				} else {
					F_pfree(m, v10)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v16)
					}
				}
			}
		}
	}
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
func F_dutch_ISO_8859_1_close_env(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	if l0 != 0 {
		v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		F_lose_s(m, v2)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			F_SN_delete_env(m, l0)
			mBase = m.M
			v6 = m.ExcPending
			if v6 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_SN_delete_env(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_dxsyn_lexize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v107 int64
	_ = v107
	v2 = int32(0)
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v14 == v2 {
		v107 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v107
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v18 == int32(0) {
		v107 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = F_pnstrdup(m, v21, v14)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	v27 = F_str_tolower(m, v22, v14, int32(100))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v27
	F_pfree(m, v22)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v34 = int32(8)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v40 = F_bsearch(m, v12+v34, v36, v37, v34, int32(_a_F_dxsyn_lexize_0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	F_pfree(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v40 == int32(0) {
		v107 = v9
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v49 = F_palloc(m, int32(8))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v51 = v47
	v54 = v49
	v57 = v2
	goto L12
L12:
	;
	v62 = F_find_word(m, v51, v12+int32(4))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v91<<(uint(int32(3))%32))+4)) = int32(0)
	v107 = base.I64_extend_i32_u(v89)
	goto L1
L14:
	;
	if v62 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = v57 << (uint(int32(3)) % 32)
	v68 = F_repalloc(m, v54, v65+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	v89 = v54
	v91 = v57
	goto L17
L17:
	;
	goto L13
L18:
	;
	if v51 != v47 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+11)))
	if v87 != 0 {
		v51 = v83
		v54 = v68
		v57 = v85
		goto L12
	} else {
		goto L24
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v75 = F_pnstrdup(m, v62, v73-v62)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+9)))
	if v71 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v83 = v72
	v85 = v57
	goto L19
L23:
	;
	v77 = v68 + v65
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v75
	v83 = v73
	v85 = v57 + int32(1)
	goto L19
L24:
	;
	v89 = v68
	v91 = v85
	goto L17
}
