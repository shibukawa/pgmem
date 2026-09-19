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
	var v46 int64
	_ = v46
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v86 float64
	_ = v86
	var v91 float64
	_ = v91
	var v105 float64
	_ = v105
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v118 float64
	_ = v118
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v156 float64
	_ = v156
	var v163 float64
	_ = v163
	var v167 float64
	_ = v167
	var v175 float64
	_ = v175
	var v176 float64
	_ = v176
	var v180 float64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
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
			v46 = base.I64_reinterpret_f64(v39)
			if v46 <= int64(4601133429810003967) {
				if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v46) {
					if base.F64_eq(v39, float64(-1)) != 0 {
						v156 = math.Float64frombits(uint64(0xfff0000000000000))
						v163 = v156
					} else {
						v163 = base.F64_div(base.F64_sub(v39, v39), float64(0))
					}
				} else {
					if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v46)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
						v163 = v39
					} else {
						if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v46) {
							v66 = float64(1)
							v67 = base.F64_add(v39, v66)
							v68 = base.I64_reinterpret_f64(v67)
							v73 = base.I32_wrap_i64(int64(base.Ui64(v68)>>(uint(int64(32))%64))) + int32(_a_F_dacosh_3)
							if base.Ui32(int32(1074790399)) < base.Ui32(v73) {
								v86 = base.F64_add(base.F64_sub(v39, v67), v66)
							} else {
								v86 = base.F64_sub(v39, base.F64_add(v67, float64(-1)))
							}
							if base.Ui32(v73) <= base.Ui32(int32(1129316351)) {
								v91 = base.F64_div(v86, v67)
							} else {
								v91 = float64(0)
							}
							v105 = base.F64_convert_i32_s(int32(base.Ui32(v73)>>(uint(int32(20))%32)) - int32(1023))
							v109 = base.F64_add(base.F64_reinterpret_i64(v68&int64(4294967295)|base.I64_extend_i32_u(v73&int32(_a_F_dacosh_4)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v110 = v105
							v111 = base.F64_add(base.F64_mul(v105, float64(1.9082149292705877e-10)), v91)
						} else {
							v109 = v39
							v110 = v40
							v111 = v40
						}
						v118 = base.F64_div(v109, base.F64_add(v109, float64(2)))
						v121 = base.F64_mul(v109, base.F64_mul(v109, float64(0.5)))
						v122 = base.F64_mul(v118, v118)
						v123 = base.F64_mul(v122, v122)
						v156 = base.F64_add(base.F64_mul(v110, float64(0.6931471803691238)), base.F64_add(v109, base.F64_sub(base.F64_add(base.F64_mul(v118, base.F64_add(v121, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v122, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v111), v121)))
						v163 = v156
					}
				}
			} else {
				if base.Ui64(int64(9218868437227405311)) < base.Ui64(v46) {
					v163 = v39
				} else {
					v66 = float64(1)
					v67 = base.F64_add(v39, v66)
					v68 = base.I64_reinterpret_f64(v67)
					v73 = base.I32_wrap_i64(int64(base.Ui64(v68)>>(uint(int64(32))%64))) + int32(_a_F_dacosh_3)
					if base.Ui32(int32(1074790399)) < base.Ui32(v73) {
						v86 = base.F64_add(base.F64_sub(v39, v67), v66)
					} else {
						v86 = base.F64_sub(v39, base.F64_add(v67, float64(-1)))
					}
					if base.Ui32(v73) <= base.Ui32(int32(1129316351)) {
						v91 = base.F64_div(v86, v67)
					} else {
						v91 = float64(0)
					}
					v105 = base.F64_convert_i32_s(int32(base.Ui32(v73)>>(uint(int32(20))%32)) - int32(1023))
					v109 = base.F64_add(base.F64_reinterpret_i64(v68&int64(4294967295)|base.I64_extend_i32_u(v73&int32(_a_F_dacosh_4)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
					v110 = v105
					v111 = base.F64_add(base.F64_mul(v105, float64(1.9082149292705877e-10)), v91)
					v118 = base.F64_div(v109, base.F64_add(v109, float64(2)))
					v121 = base.F64_mul(v109, base.F64_mul(v109, float64(0.5)))
					v122 = base.F64_mul(v118, v118)
					v123 = base.F64_mul(v122, v122)
					v156 = base.F64_add(base.F64_mul(v110, float64(0.6931471803691238)), base.F64_add(v109, base.F64_sub(base.F64_add(base.F64_mul(v118, base.F64_add(v121, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v122, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v111), v121)))
					v163 = v156
				}
			}
			v180 = v163
		} else {
			if base.Ui32(v30) <= base.Ui32(int32(1048)) {
				v167 = float64(-1)
				v175 = F_log(m, base.F64_add(base.F64_add(v4, v4), base.F64_div(v167, base.F64_add(v4, base.F64_sqrt(base.F64_add(base.F64_mul(v4, v4), v167))))))
				mBase = m.M
				v180 = v175
			} else {
				v176 = F_log(m, v4)
				mBase = m.M
				v180 = base.F64_add(v176, float64(0.6931471805599453))
			}
		}
		v181 = F_Float8GetDatum(m, v180)
		mBase = m.M
		v182 = m.ExcPending
		if v182 != 0 {
			return int32(0)
		} else {
			return v181
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
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
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
	return v781
L78:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v495
	v497 = F_r_consonant_pair_2(m, l0)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L83
	} else {
		goto L116
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
	v328 = int32(1)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v318-v328))))
	if base.B2i32(v330&int32(224) != int32(96))|base.B2i32(v328<<(uint(v330)%32)&int32(_a_F_danish_UTF_8_stem_2) == int32(0)) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v344 = F_find_among_b(m, l0, int32(_a_F_danish_UTF_8_stem_3), int32(32))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return int32(0)
L84:
	;
	if v344 == int32(0) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v351
	switch v344 - int32(1) {
	case 0:
		goto L87
	case 1:
		goto L86
	default:
		goto L78
	}
L86:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L92
L87:
	;
	v355 = F_slice_del(m, l0)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L83
	} else {
		goto L88
	}
L88:
	;
	if int32(0) <= v355 {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v781 = v355
	goto L77
L90:
	;
	if v487 != 0 {
		goto L78
	} else {
		goto L113
	}
L91:
	;
	v487 = v480
	goto L90
L92:
	;
	if v371 <= v372 {
		v480 = int32(-1)
		goto L91
	} else {
		goto L94
	}
L93:
	;
	v480 = int32(0)
	goto L91
L94:
	;
	v389 = int32(1)
	v390 = v371 - v389
	v392 = int32(*(*int8)(unsafe.Add(mBase, uint32(v373+v390))))
	v394 = v392 & int32(255)
	if base.B2i32(v390 == v372)|base.B2i32(int32(0) <= v392) != 0 {
		v452 = v394
		v456 = v389
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if int32(229) < v452 {
		goto L103
	} else {
		goto L104
	}
L96:
	;
	v401 = v394 & int32(63)
	v403 = v371 - int32(2)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v403))))
	v407 = v405 << (uint(int32(6)) % 32)
	if base.B2i32(v403 != v372)&base.B2i32(base.Ui32(v405) < base.Ui32(int32(192))) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v452 = v407&int32(1984) | v401
	v456 = int32(2)
	goto L95
L98:
	;
	goto L99
L99:
	;
	v420 = v407&int32(4032) | v401
	v422 = v371 - int32(3)
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v422))))
	if base.B2i32(v422 != v372)&base.B2i32(base.Ui32(v424) < base.Ui32(int32(224))) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v452 = v424<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v420
	v456 = int32(3)
	goto L95
L101:
	;
	goto L102
L102:
	;
	v442 = int32(4)
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371+v373-v442))))
	v452 = v424<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_4) | v444&int32(7)<<(uint(int32(18))%32) | v420
	v456 = v442
	goto L95
L103:
	;
	v487 = v456
	goto L90
L104:
	;
	goto L105
L105:
	;
	v458 = v452 - int32(97)
	if v458 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v487 = v456
	goto L90
L107:
	;
	goto L108
L108:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v458)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[1]))))
	if int32(base.Ui32(v464)>>(uint(v458&int32(7))%32))&int32(1) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v487 = v456
	goto L90
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v371 - v456
	goto L112
L112:
	;
	goto L93
L113:
	;
	v488 = F_slice_del(m, l0)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L83
	} else {
		goto L114
	}
L114:
	;
	if int32(0) <= v488 {
		goto L78
	} else {
		goto L115
	}
L115:
	;
	v781 = v488
	goto L77
L116:
	;
	if v497 < int32(0) {
		v781 = v497
		goto L77
	} else {
		goto L117
	}
L117:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v501
	v504 = int32(2)
	v506 = int32(0)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v501-v509 < v504 {
		v519 = v506
		goto L120
	} else {
		goto L121
	}
L118:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v547
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	if v547 < v550 {
		goto L131
	} else {
		goto L132
	}
L119:
	;
	if v519 == int32(0) {
		goto L118
	} else {
		goto L123
	}
L120:
	;
	goto L119
L121:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v515 = F_memcmp(m, v512+v501-v504, int32(_a_F_danish_UTF_8_stem_5), v504)
	mBase = m.M
	if v515 != 0 {
		v519 = v506
		goto L120
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v501 - v504
	v519 = int32(1)
	goto L120
L123:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v522
	v524 = int32(2)
	v526 = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v522-v529 < v524 {
		v539 = v526
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v539 == int32(0) {
		goto L118
	} else {
		goto L128
	}
L125:
	;
	goto L124
L126:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v535 = F_memcmp(m, v532+v522-v524, int32(_a_F_danish_UTF_8_stem_6), v524)
	mBase = m.M
	if v535 != 0 {
		v539 = v526
		goto L125
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v522 - v524
	v539 = int32(1)
	goto L125
L128:
	;
	v542 = F_slice_del(m, l0)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L83
	} else {
		goto L129
	}
L129:
	;
	if v542 < int32(0) {
		v781 = v542
		goto L77
	} else {
		goto L130
	}
L130:
	;
	goto L118
L131:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v601
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	if v601 < v604 {
		goto L146
	} else {
		goto L147
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v547
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v550
	v556 = v547 - int32(1)
	if v556 <= v550 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v553
	goto L131
L134:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558+v556))))
	if base.B2i32(v560&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v560)%32)&int32(_a_F_danish_UTF_8_stem_7) == int32(0)) != 0 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v574 = F_find_among_b(m, l0, int32(_a_F_danish_UTF_8_stem_8), int32(5))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L83
	} else {
		goto L136
	}
L136:
	;
	if v574 == int32(0) {
		goto L133
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v553
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v579
	switch v574 - int32(1) {
	case 0:
		goto L139
	case 1:
		goto L138
	default:
		goto L131
	}
L138:
	;
	v593 = F_slice_from_s(m, l0, int32(4), int32(_a_F_danish_UTF_8_stem_9))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L83
	} else {
		goto L144
	}
L139:
	;
	v583 = F_slice_del(m, l0)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L83
	} else {
		goto L140
	}
L140:
	;
	if v583 < int32(0) {
		v781 = v583
		goto L77
	} else {
		goto L141
	}
L141:
	;
	v587 = F_r_consonant_pair_2(m, l0)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L83
	} else {
		goto L142
	}
L142:
	;
	if int32(0) <= v587 {
		goto L131
	} else {
		goto L143
	}
L143:
	;
	v781 = v587
	goto L77
L144:
	;
	if int32(0) <= v593 {
		goto L131
	} else {
		goto L145
	}
L145:
	;
	v781 = v593
	goto L77
L146:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v778
	v781 = int32(1)
	goto L77
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v601
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v604
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L150
L148:
	;
	if v737 != 0 {
		goto L171
	} else {
		goto L172
	}
L149:
	;
	v737 = v730
	goto L148
L150:
	;
	if v621 <= v604 {
		v730 = int32(-1)
		goto L149
	} else {
		goto L152
	}
L151:
	;
	v730 = int32(0)
	goto L149
L152:
	;
	v639 = int32(1)
	v640 = v621 - v639
	v642 = int32(*(*int8)(unsafe.Add(mBase, uint32(v623+v640))))
	v644 = v642 & int32(255)
	if base.B2i32(v640 == v604)|base.B2i32(int32(0) <= v642) != 0 {
		v702 = v644
		v706 = v639
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if int32(122) < v702 {
		goto L161
	} else {
		goto L162
	}
L154:
	;
	v651 = v644 & int32(63)
	v653 = v621 - int32(2)
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623+v653))))
	v657 = v655 << (uint(int32(6)) % 32)
	if base.B2i32(v653 != v604)&base.B2i32(base.Ui32(v655) < base.Ui32(int32(192))) == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v702 = v657&int32(1984) | v651
	v706 = int32(2)
	goto L153
L156:
	;
	goto L157
L157:
	;
	v670 = v657&int32(4032) | v651
	v672 = v621 - int32(3)
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623+v672))))
	if base.B2i32(v672 != v604)&base.B2i32(base.Ui32(v674) < base.Ui32(int32(224))) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v702 = v674<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_1) | v670
	v706 = int32(3)
	goto L153
L159:
	;
	goto L160
L160:
	;
	v692 = int32(4)
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v623-v692))))
	v702 = v674<<(uint(int32(12))%32)&int32(_a_F_danish_UTF_8_stem_4) | v694&int32(7)<<(uint(int32(18))%32) | v670
	v706 = v692
	goto L153
L161:
	;
	v737 = v706
	goto L148
L162:
	;
	goto L163
L163:
	;
	v708 = v702 - int32(98)
	if v708 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v737 = v706
	goto L148
L165:
	;
	goto L166
L166:
	;
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v708)>>(uint(int32(3))%32)))+uint32(_c_F_danish_UTF_8_stem[2]))))
	if int32(base.Ui32(v714)>>(uint(v708&int32(7))%32))&int32(1) == int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v737 = v706
	goto L148
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v621 - v706
	goto L170
L170:
	;
	goto L151
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v607
	goto L146
L172:
	;
	goto L173
L173:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v739
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	v743 = F_slice_to(m, l0, v742)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L83
	} else {
		goto L174
	}
L174:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v745))) = v743
	if v743 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	return int32(-1)
L176:
	;
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v607
	v752 = int32(0)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v743-int32(4))))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v758-v607 < v757 {
		v769 = v752
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v769 == int32(0) {
		goto L146
	} else {
		goto L182
	}
L179:
	;
	goto L178
L180:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v765 = F_memcmp(m, v762+v758-v757, v743, v757)
	mBase = m.M
	if v765 != 0 {
		v769 = v752
		goto L179
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v758 - v757
	v769 = int32(1)
	goto L179
L182:
	;
	v772 = F_slice_del(m, l0)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L83
	} else {
		goto L183
	}
L183:
	;
	if v772 < int32(0) {
		v781 = v772
		goto L77
	} else {
		goto L184
	}
L184:
	;
	goto L146
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
	var v56 int64
	_ = v56
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v96 float64
	_ = v96
	var v101 float64
	_ = v101
	var v115 float64
	_ = v115
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v128 float64
	_ = v128
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v166 float64
	_ = v166
	var v173 float64
	_ = v173
	var v176 float64
	_ = v176
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
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
						v176 = v26
					} else {
						v37 = base.F64_add(v26, v26)
						v49 = base.F64_add(v37, base.F64_div(base.F64_mul(v26, v37), base.F64_sub(float64(1), v26)))
						v50 = float64(0)
						v56 = base.I64_reinterpret_f64(v49)
						if v56 <= int64(4601133429810003967) {
							if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v56) {
								if base.F64_eq(v49, float64(-1)) != 0 {
									v166 = math.Float64frombits(uint64(0xfff0000000000000))
									v173 = v166
								} else {
									v173 = base.F64_div(base.F64_sub(v49, v49), float64(0))
								}
							} else {
								if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v56)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
									v173 = v49
								} else {
									if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v56) {
										v76 = float64(1)
										v77 = base.F64_add(v49, v76)
										v78 = base.I64_reinterpret_f64(v77)
										v83 = base.I32_wrap_i64(int64(base.Ui64(v78)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
										if base.Ui32(int32(1074790399)) < base.Ui32(v83) {
											v96 = base.F64_add(base.F64_sub(v49, v77), v76)
										} else {
											v96 = base.F64_sub(v49, base.F64_add(v77, float64(-1)))
										}
										if base.Ui32(v83) <= base.Ui32(int32(1129316351)) {
											v101 = base.F64_div(v96, v77)
										} else {
											v101 = float64(0)
										}
										v115 = base.F64_convert_i32_s(int32(base.Ui32(v83)>>(uint(int32(20))%32)) - int32(1023))
										v119 = base.F64_add(base.F64_reinterpret_i64(v78&int64(4294967295)|base.I64_extend_i32_u(v83&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
										v120 = v115
										v121 = base.F64_add(base.F64_mul(v115, float64(1.9082149292705877e-10)), v101)
									} else {
										v119 = v49
										v120 = v50
										v121 = v50
									}
									v128 = base.F64_div(v119, base.F64_add(v119, float64(2)))
									v131 = base.F64_mul(v119, base.F64_mul(v119, float64(0.5)))
									v132 = base.F64_mul(v128, v128)
									v133 = base.F64_mul(v132, v132)
									v166 = base.F64_add(base.F64_mul(v120, float64(0.6931471803691238)), base.F64_add(v119, base.F64_sub(base.F64_add(base.F64_mul(v128, base.F64_add(v131, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v132, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v121), v131)))
									v173 = v166
								}
							}
						} else {
							if base.Ui64(int64(9218868437227405311)) < base.Ui64(v56) {
								v173 = v49
							} else {
								v76 = float64(1)
								v77 = base.F64_add(v49, v76)
								v78 = base.I64_reinterpret_f64(v77)
								v83 = base.I32_wrap_i64(int64(base.Ui64(v78)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
								if base.Ui32(int32(1074790399)) < base.Ui32(v83) {
									v96 = base.F64_add(base.F64_sub(v49, v77), v76)
								} else {
									v96 = base.F64_sub(v49, base.F64_add(v77, float64(-1)))
								}
								if base.Ui32(v83) <= base.Ui32(int32(1129316351)) {
									v101 = base.F64_div(v96, v77)
								} else {
									v101 = float64(0)
								}
								v115 = base.F64_convert_i32_s(int32(base.Ui32(v83)>>(uint(int32(20))%32)) - int32(1023))
								v119 = base.F64_add(base.F64_reinterpret_i64(v78&int64(4294967295)|base.I64_extend_i32_u(v83&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
								v120 = v115
								v121 = base.F64_add(base.F64_mul(v115, float64(1.9082149292705877e-10)), v101)
								v128 = base.F64_div(v119, base.F64_add(v119, float64(2)))
								v131 = base.F64_mul(v119, base.F64_mul(v119, float64(0.5)))
								v132 = base.F64_mul(v128, v128)
								v133 = base.F64_mul(v132, v132)
								v166 = base.F64_add(base.F64_mul(v120, float64(0.6931471803691238)), base.F64_add(v119, base.F64_sub(base.F64_add(base.F64_mul(v128, base.F64_add(v131, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v132, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v121), v131)))
								v173 = v166
							}
						}
						v176 = base.F64_mul(v173, float64(0.5))
					}
				} else {
					v45 = base.F64_div(v26, base.F64_sub(float64(1), v26))
					v49 = base.F64_add(v45, v45)
					v50 = float64(0)
					v56 = base.I64_reinterpret_f64(v49)
					if v56 <= int64(4601133429810003967) {
						if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v56) {
							if base.F64_eq(v49, float64(-1)) != 0 {
								v166 = math.Float64frombits(uint64(0xfff0000000000000))
								v173 = v166
							} else {
								v173 = base.F64_div(base.F64_sub(v49, v49), float64(0))
							}
						} else {
							if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v56)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
								v173 = v49
							} else {
								if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v56) {
									v76 = float64(1)
									v77 = base.F64_add(v49, v76)
									v78 = base.I64_reinterpret_f64(v77)
									v83 = base.I32_wrap_i64(int64(base.Ui64(v78)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
									if base.Ui32(int32(1074790399)) < base.Ui32(v83) {
										v96 = base.F64_add(base.F64_sub(v49, v77), v76)
									} else {
										v96 = base.F64_sub(v49, base.F64_add(v77, float64(-1)))
									}
									if base.Ui32(v83) <= base.Ui32(int32(1129316351)) {
										v101 = base.F64_div(v96, v77)
									} else {
										v101 = float64(0)
									}
									v115 = base.F64_convert_i32_s(int32(base.Ui32(v83)>>(uint(int32(20))%32)) - int32(1023))
									v119 = base.F64_add(base.F64_reinterpret_i64(v78&int64(4294967295)|base.I64_extend_i32_u(v83&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
									v120 = v115
									v121 = base.F64_add(base.F64_mul(v115, float64(1.9082149292705877e-10)), v101)
								} else {
									v119 = v49
									v120 = v50
									v121 = v50
								}
								v128 = base.F64_div(v119, base.F64_add(v119, float64(2)))
								v131 = base.F64_mul(v119, base.F64_mul(v119, float64(0.5)))
								v132 = base.F64_mul(v128, v128)
								v133 = base.F64_mul(v132, v132)
								v166 = base.F64_add(base.F64_mul(v120, float64(0.6931471803691238)), base.F64_add(v119, base.F64_sub(base.F64_add(base.F64_mul(v128, base.F64_add(v131, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v132, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v121), v131)))
								v173 = v166
							}
						}
					} else {
						if base.Ui64(int64(9218868437227405311)) < base.Ui64(v56) {
							v173 = v49
						} else {
							v76 = float64(1)
							v77 = base.F64_add(v49, v76)
							v78 = base.I64_reinterpret_f64(v77)
							v83 = base.I32_wrap_i64(int64(base.Ui64(v78)>>(uint(int64(32))%64))) + int32(_a_F_datanh_0)
							if base.Ui32(int32(1074790399)) < base.Ui32(v83) {
								v96 = base.F64_add(base.F64_sub(v49, v77), v76)
							} else {
								v96 = base.F64_sub(v49, base.F64_add(v77, float64(-1)))
							}
							if base.Ui32(v83) <= base.Ui32(int32(1129316351)) {
								v101 = base.F64_div(v96, v77)
							} else {
								v101 = float64(0)
							}
							v115 = base.F64_convert_i32_s(int32(base.Ui32(v83)>>(uint(int32(20))%32)) - int32(1023))
							v119 = base.F64_add(base.F64_reinterpret_i64(v78&int64(4294967295)|base.I64_extend_i32_u(v83&int32(_a_F_datanh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v120 = v115
							v121 = base.F64_add(base.F64_mul(v115, float64(1.9082149292705877e-10)), v101)
							v128 = base.F64_div(v119, base.F64_add(v119, float64(2)))
							v131 = base.F64_mul(v119, base.F64_mul(v119, float64(0.5)))
							v132 = base.F64_mul(v128, v128)
							v133 = base.F64_mul(v132, v132)
							v166 = base.F64_add(base.F64_mul(v120, float64(0.6931471803691238)), base.F64_add(v119, base.F64_sub(base.F64_add(base.F64_mul(v128, base.F64_add(v131, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v132, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, base.F64_add(base.F64_mul(v133, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v121), v131)))
							v173 = v166
						}
					}
					v176 = base.F64_mul(v173, float64(0.5))
				}
				if v27 < int64(0) {
					v181 = base.F64_neg(v176)
				} else {
					v181 = v176
				}
				v182 = F_Float8GetDatum(m, v181)
				mBase = m.M
				v183 = m.ExcPending
				if v183 != 0 {
					return int32(0)
				} else {
					return v182
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v188 = m.ExcPending
		if v188 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v191 = m.ExcPending
			if v191 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_datanh_2), int32(0))
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_datanh_3), int32(2727), int32(_a_F_datanh_4))
					mBase = m.M
					v200 = m.ExcPending
					if v200 != 0 {
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
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
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
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v314 int64
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int64
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
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
	v643 = m.ExcPending
	if v643 != 0 {
		goto L8
	} else {
		goto L191
	}
L2:
	;
	m.G0 = v12 + int32(160)
	return
L3:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	v562 = F_GetDatabasePath(m, v560, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L8
	} else {
		goto L163
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L8
	} else {
		goto L160
	}
L5:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[0]))
	if base.Ui32(int32(2)) <= base.Ui32(v327) {
		goto L98
	} else {
		goto L99
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
	v114 = F___fstatat(m, int32(-100), v58, v12-int32(-64), int32(0))
	mBase = m.M
	goto L44
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
	if base.B2i32(v72 == int32(47))&base.B2i32(base.Ui32(v58) < base.Ui32(v71)) != 0 {
		v67 = v71
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v78 = v71
	goto L29
L28:
	;
	goto L27
L29:
	;
	if base.Ui32(v58) < base.Ui32(v78) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v90 = v78
	goto L35
L31:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v84 != int32(47) {
		v78 = v78 - int32(1)
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
	if base.Ui32(v58) < base.Ui32(v90) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v58 == v90 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v94 = v90 - int32(1)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v95 == int32(47) {
		v90 = v94
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
	v103 = v58 + base.B2i32(v63 == int32(47))
	goto L43
L42:
	;
	v103 = v90
	goto L43
L43:
	;
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v104)
	goto L25
L44:
	;
	if v114 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[1]))
	if v118 != int32(44) {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_pfree(m, v58)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	F_recovery_create_dbdir(m, v58, int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v130 = F___fstatat(m, int32(-100), v23, v12-int32(-64), int32(0))
	mBase = m.M
	goto L52
L51:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v141 = m.G0
	v143 = v141 - int32(32)
	m.G0 = v143
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[2]))
	if int32(0) < v146 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	if int32(0) <= v130 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[1]))
	if v134 != int32(44) {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	F_recovery_create_dbdir(m, v23, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	v156 = int32(0)
	goto L59
L57:
	;
	goto L58
L58:
	;
	m.G0 = v143 + int32(32)
	v314 = F_EmitProcSignalBarrier(m)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L8
	} else {
		goto L93
	}
L59:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[3]))
	v162 = v159 + v156<<(uint(int32(6))%32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v163 != v140 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L58
L61:
	;
	v298 = v156 + int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[2]))
	if v298 < v300 {
		v156 = v298
		goto L59
	} else {
		goto L92
	}
L62:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[4]))
	F_ResourceOwnerEnlarge(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143)+28)) = int32(_a_F_dbase_redo_5)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+24)) = int32(_a_F_dbase_redo_6)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+20)) = int32(_a_F_dbase_redo_7)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v143)+8)) = int64(0)
	v181 = int32(_a_F_dbase_redo_8)
	v183 = base.AtomicRmwOr32(m, v162, int32(24), v181)
	if v183&v181 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	goto L68
L66:
	;
	v207 = v183
	goto L67
L67:
	;
	v216 = int32(_a_F_dbase_redo_9)
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[5]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v143+int32(8))+8))
	if v219 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	F_perform_spin_delay(m, v143+int32(8))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L70
	}
L69:
	;
	v207 = v201
	goto L67
L70:
	;
	v199 = int32(_a_F_dbase_redo_8)
	v201 = base.AtomicRmwOr32(m, v162, int32(24), v199)
	if v201&v199 != 0 {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v236 = int32(25165824)
	if v207&v236 != v236 {
		goto L83
	} else {
		goto L84
	}
L73:
	;
	goto L72
L74:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[5])) = v234
	goto L73
L75:
	;
	if int32(999) < v217 {
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	if v217 < int32(11) {
		goto L73
	} else {
		goto L82
	}
L78:
	;
	v224 = int32(900)
	if v224 <= v217 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v227 = v224
	goto L81
L80:
	;
	v227 = v217
	goto L81
L81:
	;
	v234 = v227 + int32(100)
	goto L74
L82:
	;
	v234 = v217 - int32(1)
	goto L74
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+24)) = v207 & int32(-4194305)
	goto L61
L84:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v240 != v140 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v162)+24))
	v243 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+24)) = (v242 + v243) & int32(-4194305)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	v249 = int32(_a_F_dbase_redo_10)
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[6])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = v243
	v257 = v248 + v243
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v257
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[4]))
	F_ResourceOwnerRemember(m, v260, v257, int32(_a_F_dbase_redo_11))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v265 = v162 + int32(48)
	v267 = F_LWLockAcquire(m, v265, int32(1))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	F_FlushBuffer(m, v162, int32(0), int32(3))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	F_LWLockRelease(m, v265)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[4]))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	F_ResourceOwnerForget(m, v276, v277+int32(1), int32(_a_F_dbase_redo_11))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_UnpinBufferNoOwner(m, v162)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	goto L61
L92:
	;
	goto L60
L93:
	;
	F_WaitForProcSignalBarrier(m, v314)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	F_copydir(m, v23, v27, int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	F_pfree(m, v23)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_pfree(m, v27)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	goto L2
L98:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	F_LockSharedObjectForSession(m, v330)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L8
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	F_ReplicationSlotsDropDBSlots(m, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L8
	} else {
		goto L111
	}
L101:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v334 = F_CountDBBackends(m, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	if int32(0) < v334 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	goto L106
L104:
	;
	goto L105
L105:
	;
	goto L100
L106:
	;
	F_CancelDBBackends(m, v333, int32(7), int32(1))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L8
	} else {
		goto L108
	}
L107:
	;
	goto L105
L108:
	;
	F_pg_usleep(m, int32(_a_F_dbase_redo_12))
	mBase = m.M
	v353 = F_CountDBBackends(m, v333)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	if int32(0) < v353 {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	goto L107
L111:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	F_DropDatabaseBuffers(m, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	F_ForgetDatabaseSyncRequests(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L8
	} else {
		goto L113
	}
L113:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v385 = m.G0
	v387 = v385 - int32(112)
	m.G0 = v387
	F_smgrdestroyall(m)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[7]))
	if v392 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v479 = F_EmitProcSignalBarrier(m)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L8
	} else {
		goto L141
	}
L116:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L8
	} else {
		goto L138
	}
L117:
	;
	m.G0 = v387 + int32(112)
	goto L115
L118:
	;
	v396 = v387 + int32(92)
	F_hash_seq_init(m, v396, v392)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L8
	} else {
		goto L119
	}
L119:
	;
	v399 = F_hash_seq_search(m, v396)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L8
	} else {
		goto L120
	}
L120:
	;
	if v399 == int32(0) {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v403 = v399
	goto L122
L122:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v384 == v412 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L117
L124:
	;
	v416 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L8
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v452 = F_hash_seq_search(m, v387+int32(92))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L8
	} else {
		goto L136
	}
L127:
	;
	if v416 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v403)+16))
	v420 = v387 + int32(20)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	F_GetRelationPath(m, v420, v421, v422, v423, int32(-1), v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L8
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[7]))
	v444 = F_hash_search(m, v441, v403, int32(2), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L8
	} else {
		goto L134
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v420
	F_errmsg_internal(m, int32(_a_F_dbase_redo_13), v387)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L8
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_14), int32(212), int32(_a_F_dbase_redo_15))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	if v444 == int32(0) {
		goto L116
	} else {
		goto L135
	}
L135:
	;
	goto L126
L136:
	;
	if v452 != 0 {
		v403 = v452
		goto L122
	} else {
		goto L137
	}
L137:
	;
	goto L123
L138:
	;
	F_errmsg_internal(m, int32(_a_F_dbase_redo_16), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_14), int32(217), int32(_a_F_dbase_redo_15))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L8
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_WaitForProcSignalBarrier(m, v479)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L8
	} else {
		goto L142
	}
L142:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if int32(0) < v483 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v489 = int32(0)
	goto L146
L144:
	;
	goto L145
L145:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_dbase_redo[0]))
	if base.Ui32(v540) < base.Ui32(int32(2)) {
		goto L2
	} else {
		goto L158
	}
L146:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v325+int32(8)+v489<<(uint(int32(2))%32))))
	v503 = F_GetDatabasePath(m, v498, v502)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L8
	} else {
		goto L149
	}
L147:
	;
	goto L145
L148:
	;
	F_pfree(m, v503)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L8
	} else {
		goto L156
	}
L149:
	;
	v505 = F_rmtree(m, v503)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L8
	} else {
		goto L150
	}
L150:
	;
	if v505 != 0 {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v509 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L8
	} else {
		goto L152
	}
L152:
	;
	if v509 == int32(0) {
		goto L148
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v503
	F_errmsg(m, int32(_a_F_dbase_redo_2), v12+int32(48))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L8
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3454), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L8
	} else {
		goto L155
	}
L155:
	;
	goto L148
L156:
	;
	v527 = v489 + int32(1)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if v527 < v528 {
		v489 = v527
		goto L146
	} else {
		goto L157
	}
L157:
	;
	goto L147
L158:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	F_UnlockSharedObjectForSession(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L8
	} else {
		goto L159
	}
L159:
	;
	goto L2
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v17
	F_errmsg_internal(m, int32(_a_F_dbase_redo_17), v12)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L8
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3471), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L8
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	v564 = F_pstrdup(m, v562)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L8
	} else {
		goto L164
	}
L164:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	if v569 != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	F_recovery_create_dbdir(m, v564, int32(1))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L8
	} else {
		goto L187
	}
L166:
	;
	v570 = F_strlen(m, v564)
	mBase = m.M
	v573 = v570 + v564
	goto L169
L167:
	;
	goto L168
L168:
	;
	goto L165
L169:
	;
	v577 = v573 - int32(1)
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	if base.B2i32(v578 == int32(47))&base.B2i32(base.Ui32(v564) < base.Ui32(v577)) != 0 {
		v573 = v577
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v584 = v577
	goto L172
L171:
	;
	goto L170
L172:
	;
	if base.Ui32(v564) < base.Ui32(v584) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v596 = v584
	goto L178
L174:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584))))
	if v590 != int32(47) {
		v584 = v584 - int32(1)
		goto L172
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	goto L173
L177:
	;
	goto L176
L178:
	;
	if base.Ui32(v564) < base.Ui32(v596) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v564 == v596 {
		goto L184
	} else {
		goto L185
	}
L180:
	;
	v600 = v596 - int32(1)
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	if v601 == int32(47) {
		v596 = v600
		goto L178
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	goto L179
L183:
	;
	goto L182
L184:
	;
	v609 = v564 + base.B2i32(v569 == int32(47))
	goto L186
L185:
	;
	v609 = v596
	goto L186
L186:
	;
	v610 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v609))) = uint8(v610)
	goto L168
L187:
	;
	F_pfree(m, v564)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L8
	} else {
		goto L188
	}
L188:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	F_CreateDirAndVersionFile(m, v562, v621, v622, int32(1))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L8
	} else {
		goto L189
	}
L189:
	;
	F_pfree(m, v562)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L8
	} else {
		goto L190
	}
L190:
	;
	goto L2
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v27
	F_errmsg(m, int32(_a_F_dbase_redo_18), v12+int32(16))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L8
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_dbase_redo_3), int32(3354), int32(_a_F_dbase_redo_4))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L8
	} else {
		goto L193
	}
L193:
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
			v39 = F_DirectFunctionCall2Coll(m, int32(578), int32(0), v15, v38)
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
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
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
	return v147
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v132
	if l2 < v134 {
		goto L42
	} else {
		goto L43
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v132 = v15
	v134 = v12
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
	v147 = int32(0)
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
	v105 = v19
	goto L11
L11:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v110 = l0 + int32(24)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v112 = m.Env.Pgmem_zstream_read(m, v108, v110, v111)
	mBase = m.M
	if v112 < int32(0) {
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
	v147 = v60
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
	v147 = int32(-100)
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
	F_ResourceOwnerRemember(m, v95, v91, int32(_a_F_decompress_read_2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v91
	v105 = v91
	goto L11
L38:
	;
	v147 = int32(-100)
	goto L1
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v110
	if v112 != 0 {
		v132 = v110
		v134 = v112
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
	v137 = l2
	goto L44
L43:
	;
	v137 = v134
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v134 - v137
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v140 + v137
	v147 = v137
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
func F_die(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_die[0])))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_die[1])) = v7
	*(*int32)(unsafe.Add(mBase, _c_F_die[2])) = v7
	goto L3
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_die[3])) = int32(4)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_die[4]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_die[5])))
	if v61 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v20 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v23 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_die[6]))
	if v27 == v23 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_die[7]))
	if v34 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v57 = F_pgmem_kill(m, v23, int32(23))
	mBase = m.M
	goto L5
L12:
	;
	m.G0 = v31 + int32(16)
	goto L4
L13:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+15)) = uint8(v37)
	goto L14
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_die[8]))
	v45 = F_write(m, v41, v31+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v45 {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_die[9]))
	if v49 == int32(27) {
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
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_die[10]))
	if v65 == int32(2) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
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
func F_distance_taxicab(m *base.Module, l0 int32) int32 {
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	v2 = float64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v26 = F_pg_detoast_datum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
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
			v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if base.Ui32(v30) < base.Ui32(v33) {
				if v242 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v246 = m.ExcPending
					if v246 != 0 {
						return int32(0)
					} else {
						v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v26 != v247 {
							F_pfree(m, v26)
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return int32(0)
							} else {
								v256 = F_Float8GetDatum(m, v228)
								mBase = m.M
								v257 = m.ExcPending
								if v257 != 0 {
									return int32(0)
								} else {
									return v256
								}
							}
						} else {
							v256 = F_Float8GetDatum(m, v228)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								return v256
							}
						}
					}
				} else {
					v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v26 != v247 {
						F_pfree(m, v26)
						mBase = m.M
						v255 = m.ExcPending
						if v255 != 0 {
							return int32(0)
						} else {
							v256 = F_Float8GetDatum(m, v228)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								return v256
							}
						}
					} else {
						v256 = F_Float8GetDatum(m, v228)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							return v256
						}
					}
				}
			} else {
				if v242 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v251 = m.ExcPending
					if v251 != 0 {
						return int32(0)
					} else {
						v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v26 == v252 {
							v256 = F_Float8GetDatum(m, v228)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								return v256
							}
						} else {
							F_pfree(m, v26)
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return int32(0)
							} else {
								v256 = F_Float8GetDatum(m, v228)
								mBase = m.M
								v257 = m.ExcPending
								if v257 != 0 {
									return int32(0)
								} else {
									return v256
								}
							}
						}
					}
				} else {
					v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v26 == v252 {
						v256 = F_Float8GetDatum(m, v228)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							return v256
						}
					} else {
						F_pfree(m, v26)
						mBase = m.M
						v255 = m.ExcPending
						if v255 != 0 {
							return int32(0)
						} else {
							v256 = F_Float8GetDatum(m, v228)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								return v256
							}
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
	var v43 int32
	_ = v43
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = F_str_tolower(m, v7, v8, int32(100))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = F_palloc0(m, int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if int32(1001) <= v8 {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v10
				return v15
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				if v21 != 0 {
					v24 = F_searchstoplist(m, v6+int32(4), v10)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if v24 == int32(0) {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
							if v31 != int32(1) {
								v41 = v10
								v42 = int32(_a_F_dsnowball_lexize_0)
								v43 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0]))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v45
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
								v48 = F_strlen(m, v41)
								mBase = m.M
								v49 = F_SN_set_current(m, v47, v48, v41)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
									v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v51)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v43
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
										if v58 == int32(0) {
											v77 = v41
											v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
											if v80 != int32(1) {
												*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
												return v15
											} else {
												v85 = F_strlen(m, v77)
												mBase = m.M
												v87 = F_pg_any_to_server(m, v77, v85, int32(6))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int32(0)
												} else {
													if v77 == v87 {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
														return v15
													} else {
														F_pfree(m, v77)
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v87
															return v15
														}
													}
												}
											}
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
											if v61 == int32(0) {
												v77 = v41
												v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
												if v80 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
													return v15
												} else {
													v85 = F_strlen(m, v77)
													mBase = m.M
													v87 = F_pg_any_to_server(m, v77, v85, int32(6))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														if v77 == v87 {
															*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
															return v15
														} else {
															F_pfree(m, v77)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v87
																return v15
															}
														}
													}
												}
											} else {
												v66 = F_repalloc(m, v41, v61+int32(1))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
													if v69 != 0 {
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
														base.MemoryCopy(m, v66, v70, v69)
													} else {
													}
													v72 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v66+v73))) = uint8(v75)
													v77 = v66
													v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
													if v80 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
														return v15
													} else {
														v85 = F_strlen(m, v77)
														mBase = m.M
														v87 = F_pg_any_to_server(m, v77, v85, int32(6))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															if v77 == v87 {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																return v15
															} else {
																F_pfree(m, v77)
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v87
																	return v15
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
								v34 = F_strlen(m, v10)
								mBase = m.M
								v36 = F_pg_server_to_any(m, v10, v34, int32(6))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									if v10 == v36 {
										v41 = v10
										v42 = int32(_a_F_dsnowball_lexize_0)
										v43 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0]))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
										*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v45
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
										v48 = F_strlen(m, v41)
										mBase = m.M
										v49 = F_SN_set_current(m, v47, v48, v41)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
											v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v51)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v43
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
												if v58 == int32(0) {
													v77 = v41
													v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
													if v80 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
														return v15
													} else {
														v85 = F_strlen(m, v77)
														mBase = m.M
														v87 = F_pg_any_to_server(m, v77, v85, int32(6))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															if v77 == v87 {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																return v15
															} else {
																F_pfree(m, v77)
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v87
																	return v15
																}
															}
														}
													}
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
													if v61 == int32(0) {
														v77 = v41
														v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
														if v80 != int32(1) {
															*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
															return v15
														} else {
															v85 = F_strlen(m, v77)
															mBase = m.M
															v87 = F_pg_any_to_server(m, v77, v85, int32(6))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																if v77 == v87 {
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																	return v15
																} else {
																	F_pfree(m, v77)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v87
																		return v15
																	}
																}
															}
														}
													} else {
														v66 = F_repalloc(m, v41, v61+int32(1))
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
															if v69 != 0 {
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
																base.MemoryCopy(m, v66, v70, v69)
															} else {
															}
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v66+v73))) = uint8(v75)
															v77 = v66
															v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
															if v80 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																return v15
															} else {
																v85 = F_strlen(m, v77)
																mBase = m.M
																v87 = F_pg_any_to_server(m, v77, v85, int32(6))
																mBase = m.M
																v88 = m.ExcPending
																if v88 != 0 {
																	return int32(0)
																} else {
																	if v77 == v87 {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																		return v15
																	} else {
																		F_pfree(m, v77)
																		mBase = m.M
																		v93 = m.ExcPending
																		if v93 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v87
																			return v15
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
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											v41 = v36
											v42 = int32(_a_F_dsnowball_lexize_0)
											v43 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0]))
											v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
											*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v45
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
											v48 = F_strlen(m, v41)
											mBase = m.M
											v49 = F_SN_set_current(m, v47, v48, v41)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int32(0)
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
												v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
												v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v51)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_dsnowball_lexize[0])) = v43
													v57 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
													v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
													if v58 == int32(0) {
														v77 = v41
														v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
														if v80 != int32(1) {
															*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
															return v15
														} else {
															v85 = F_strlen(m, v77)
															mBase = m.M
															v87 = F_pg_any_to_server(m, v77, v85, int32(6))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																if v77 == v87 {
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																	return v15
																} else {
																	F_pfree(m, v77)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v87
																		return v15
																	}
																}
															}
														}
													} else {
														v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
														if v61 == int32(0) {
															v77 = v41
															v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
															if v80 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																return v15
															} else {
																v85 = F_strlen(m, v77)
																mBase = m.M
																v87 = F_pg_any_to_server(m, v77, v85, int32(6))
																mBase = m.M
																v88 = m.ExcPending
																if v88 != 0 {
																	return int32(0)
																} else {
																	if v77 == v87 {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																		return v15
																	} else {
																		F_pfree(m, v77)
																		mBase = m.M
																		v93 = m.ExcPending
																		if v93 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v87
																			return v15
																		}
																	}
																}
															}
														} else {
															v66 = F_repalloc(m, v41, v61+int32(1))
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int32(0)
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
																v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
																if v69 != 0 {
																	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
																	base.MemoryCopy(m, v66, v70, v69)
																} else {
																}
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
																v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
																v75 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v66+v73))) = uint8(v75)
																v77 = v66
																v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
																if v80 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																	return v15
																} else {
																	v85 = F_strlen(m, v77)
																	mBase = m.M
																	v87 = F_pg_any_to_server(m, v77, v85, int32(6))
																	mBase = m.M
																	v88 = m.ExcPending
																	if v88 != 0 {
																		return int32(0)
																	} else {
																		if v77 == v87 {
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
																			return v15
																		} else {
																			F_pfree(m, v77)
																			mBase = m.M
																			v93 = m.ExcPending
																			if v93 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v87
																				return v15
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
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								return v15
							}
						}
					}
				} else {
					F_pfree(m, v10)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						return v15
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
func F_dxsyn_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v13 == v2 {
		v98 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v98
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v17 == int32(0) {
		v98 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pnstrdup(m, v20, v13)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v26 = F_str_tolower(m, v21, v13, int32(100))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v26
	F_pfree(m, v21)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
	v33 = int32(8)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v39 = F_bsearch(m, v11+v33, v35, v36, v33, int32(_a_F_dxsyn_lexize_0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_pfree(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v39 == int32(0) {
		v98 = v2
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v48 = F_palloc(m, int32(8))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v50 = v46
	v52 = v48
	v55 = v2
	goto L12
L12:
	;
	v60 = F_find_word(m, v50, v11+int32(4))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87+v88<<(uint(int32(3))%32))+4)) = int32(0)
	v98 = v87
	goto L1
L14:
	;
	if v60 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = v55 << (uint(int32(3)) % 32)
	v66 = F_repalloc(m, v52, v63+int32(16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	v87 = v52
	v88 = v55
	goto L17
L17:
	;
	goto L13
L18:
	;
	if v50 != v46 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+11)))
	if v85 != 0 {
		v50 = v81
		v52 = v66
		v55 = v82
		goto L12
	} else {
		goto L24
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v73 = F_pnstrdup(m, v60, v71-v60)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)))
	if v69 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v81 = v70
	v82 = v55
	goto L19
L23:
	;
	v75 = v66 + v63
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v73
	v81 = v71
	v82 = v55 + int32(1)
	goto L19
L24:
	;
	v87 = v66
	v88 = v82
	goto L17
}
