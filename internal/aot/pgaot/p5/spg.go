package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_spgUpdateNodeLink(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v3 = l2
	v4 = l3
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = int32(base.Ui32(v12)>>(uint(int32(3))%32)) & int32(_a_F_spgUpdateNodeLink_0)
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)) = uint16(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+2)) = uint16(v3)
	v60 = int32(16)
	v61 = int32(base.Ui32(v3) >> (uint(v60) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v23))) = uint16(v61)
	m.G0 = v10 + v60
	return
L2:
	;
	v23 = l0 + int32(base.Ui32(v12)>>(uint(int32(16))%32)) + int32(8)
	v27 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if l1 == v27 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L4
L7:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+6)))
	v36 = v27 + int32(1)
	if v36 != v16 {
		v23 = v23 + v31&int32(_a_F_spgUpdateNodeLink_0)
		v27 = v36
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
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	F_errmsg_internal(m, int32(_a_F_spgUpdateNodeLink_1), v10)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_spgUpdateNodeLink_2), int32(68), int32(_a_F_spgUpdateNodeLink_3))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
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
func F_spg_kd_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935500989)
	return int32(0)
}
func F_spg_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v11 = v9 & int32(_a_F_spg_mask_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v11)
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(24)) <= base.Ui32(v13) {
		F_mask_unused_space(m, l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_spg_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int64
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
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
	var v537 int32
	_ = v537
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
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
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
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v706 int32
	_ = v706
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int64
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v984 int64
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1084 int32
	_ = v1084
	var v1085 int32
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
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1440 int32
	_ = v1440
	var v1441 int64
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1589 int32
	_ = v1589
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int64
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int64
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1732 int64
	_ = v1732
	var v1737 int32
	_ = v1737
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1779 int32
	_ = v1779
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1861 int32
	_ = v1861
	var v1863 int32
	_ = v1863
	var v1869 int32
	_ = v1869
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1890 int32
	_ = v1890
	var v1908 int32
	_ = v1908
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2037 int32
	_ = v2037
	var v2041 int32
	_ = v2041
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2070 int32
	_ = v2070
	var v2076 int32
	_ = v2076
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2091 int32
	_ = v2091
	var v2096 int32
	_ = v2096
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2109 int32
	_ = v2109
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2124 int32
	_ = v2124
	var v2129 int32
	_ = v2129
	var v2133 int32
	_ = v2133
	var v2137 int32
	_ = v2137
	var v2142 int32
	_ = v2142
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(192)
	m.G0 = v19
	v21 = int32(_a_F_spg_redo_0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+48)))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_redo[0])) = v27
	v30 = v24 & int32(240)
	switch int32(base.Ui32(v30-int32(16)) >> (uint(int32(4)) % 32)) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	case 5:
		goto L9
	case 6:
		goto L8
	case 7:
		goto L7
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L20
	} else {
		goto L482
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L20
	} else {
		goto L479
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2100 = m.ExcPending
	if v2100 != 0 {
		goto L20
	} else {
		goto L476
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L20
	} else {
		goto L473
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L20
	} else {
		goto L470
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_redo[0])) = v22
	v2061 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[1]))
	F_MemoryContextReset(m, v2061)
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L20
	} else {
		goto L469
	}
L7:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1716 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1718 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[2]))
	if base.Ui32(int32(2)) <= base.Ui32(v1718) {
		goto L422
	} else {
		goto L423
	}
L8:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1670 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1674 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(108))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L20
	} else {
		goto L410
	}
L9:
	;
	v1441 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+8))
	v1444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1442)+12)))
	v1446 = v19 + int32(108)
	base.MemoryFill(m, v1446, int32(0), int32(84))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)) = uint8(v1444)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v1443
	v1453 = F_palloc0(m, int32(16))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L20
	} else {
		goto L383
	}
L10:
	;
	v984 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v987 = int32(0)
	F_XLogRecGetBlockTag(m, l0, int32(2), v987, v987, v19+int32(92))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L20
	} else {
		goto L259
	}
L11:
	;
	v848 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v851 = v849 + int32(6)
	v852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v849)+10)))
	v853 = v851 + v852
	v854 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v853)+4)))
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849)+5)))
	if v855 != 0 {
		goto L218
	} else {
		goto L219
	}
L12:
	;
	v481 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+24)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v482)+12))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+16)))
	base.MemoryFill(m, v19+int32(108), int32(0), int32(84))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)) = uint8(v485)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v484
	v494 = F_palloc0(m, int32(16))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L20
	} else {
		goto L121
	}
L13:
	;
	v215 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v218 = int32(0)
	F_XLogRecGetBlockTag(m, l0, int32(1), v218, v218, v19+int32(100))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L20
	} else {
		goto L61
	}
L14:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+10))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v38 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v151 != 0 {
		goto L43
	} else {
		goto L44
	}
L16:
	;
	if v84 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v42 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v81 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(108))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L27
	}
L20:
	;
	return
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v42
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	v51 = (v45<<(uint(int32(3))%32) | int32(4)) & int32(12)
	if v42 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v84 = v42
	goto L16
L23:
	;
	F_PageInit(m, v69, int32(_a_F_spg_redo_1), int32(8))
	mBase = m.M
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
	v74 = v69 + v73
	v75 = int32(_a_F_spg_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v74)+6)) = uint16(v75)
	*(*uint16)(unsafe.Add(mBase, uint32(v74))) = uint16(v51)
	goto L22
L24:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(v42^int32(-1))<<(uint(int32(2))%32))))
	v69 = v61
	goto L23
L25:
	;
	goto L26
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v69 = v63 + v42<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L27:
	;
	if v81 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	v84 = v83
	goto L16
L29:
	;
	v104 = v36 + int32(10)
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
	if v105 != v106 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+(v84^int32(-1))<<(uint(int32(2))%32))))
	v102 = v94
	goto L29
L31:
	;
	goto L32
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v102 = v96 + v84<<(uint(int32(13))%32) + int32(-8192)
	goto L29
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v102))) = base.I64_rotr(v35, int64(32))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L20
	} else {
		goto L42
	}
L34:
	;
	F_addOrReplaceTuple(m, v102, v104, int32(base.Ui32(v37)>>(uint(int32(2))%32)), v105)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L20
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_PageIndexTupleDelete(m, v102, v105)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L39
	}
L37:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
	if v112 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v102+v112<<(uint(int32(2))%32))+20))
	v121 = v102 + v118&int32(_a_F_spg_redo_3)
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+4)))
	v128 = v122&int32(_a_F_spg_redo_4) | v125&int32(_a_F_spg_redo_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+4)) = uint16(v128)
	goto L33
L39:
	;
	v133 = int32(base.Ui32(v37) >> (uint(int32(2)) % 32))
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	v136 = F_PageAddItemExtended(m, v102, v104, v133, v134, int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	if v136 != v138 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	goto L33
L42:
	;
	goto L15
L43:
	;
	F_UnlockReleaseBuffer(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L20
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
	if v154 == int32(0) {
		goto L6
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v160 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(108))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	if v160 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v164 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v164, v164, v164, v19+int32(104))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L20
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v210 == int32(0) {
		goto L6
	} else {
		goto L59
	}
L52:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v171 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189+v190<<(uint(int32(2))%32))+20))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+8)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	F_spgUpdateNodeLink(m, v194&int32(_a_F_spg_redo_3)+v189, v198, v199, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L20
	} else {
		goto L57
	}
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175+(v171^int32(-1))<<(uint(int32(2))%32))))
	v189 = v181
	goto L53
L55:
	;
	goto L56
L56:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v189 = v183 + v171<<(uint(int32(13))%32) + int32(-8192)
	goto L53
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v189))) = base.I64_rotr(v35, int64(32))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	goto L51
L59:
	;
	F_UnlockReleaseBuffer(m, v210)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L20
	} else {
		goto L60
	}
L60:
	;
	goto L6
L61:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+16)))
	base.MemoryFill(m, v19+int32(108), int32(0), int32(84))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)) = uint8(v225)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v224
	v234 = F_palloc0(m, int32(16))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v234
	v238 = v216 + int32(20)
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
	v240 = int32(1)
	v242 = v238 + v239<<(uint(v240)%32)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+3)))
	if v246 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v247 = v240
	goto L65
L64:
	;
	v247 = v239 + v240
	goto L65
L65:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+2)))
	if v248 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v368 != 0 {
		goto L88
	} else {
		goto L89
	}
L67:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v296 < int32(0) {
		goto L80
	} else {
		goto L81
	}
L68:
	;
	v252 = F_XLogInitBufferForRedo(m, l0, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L20
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v291 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(104))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L20
	} else {
		goto L77
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v252
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+4)))
	v261 = (v255<<(uint(int32(3))%32) | int32(4)) & int32(12)
	if v252 < int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L67
L73:
	;
	F_PageInit(m, v279, int32(_a_F_spg_redo_1), int32(8))
	mBase = m.M
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+16)))
	v284 = v279 + v283
	v285 = int32(_a_F_spg_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v284)+6)) = uint16(v285)
	*(*uint16)(unsafe.Add(mBase, uint32(v284))) = uint16(v261)
	goto L72
L74:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265+(v252^int32(-1))<<(uint(int32(2))%32))))
	v279 = v271
	goto L73
L75:
	;
	goto L76
L76:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v279 = v273 + v252<<(uint(int32(13))%32) + int32(-8192)
	goto L73
L77:
	;
	if v291 != 0 {
		goto L66
	} else {
		goto L78
	}
L78:
	;
	goto L67
L79:
	;
	v318 = v247<<(uint(int32(1))%32) + v242
	v320 = int32(0)
	goto L83
L80:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v300+(v296^int32(-1))<<(uint(int32(2))%32))))
	v314 = v306
	goto L79
L81:
	;
	goto L82
L82:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v314 = v308 + v296<<(uint(int32(13))%32) + int32(-8192)
	goto L79
L83:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v335 = int32(base.Ui32(v333) >> (uint(int32(2)) % 32))
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242+v320<<(uint(int32(1))%32)))))
	F_addOrReplaceTuple(m, v314, v318, v335, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L20
	} else {
		goto L85
	}
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v314))) = base.I64_rotr(v215, int64(32))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L20
	} else {
		goto L87
	}
L85:
	;
	v344 = v320 + int32(1)
	if v344 != v247 {
		v318 = v318 + v335
		v320 = v344
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	goto L66
L88:
	;
	F_UnlockReleaseBuffer(m, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L20
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v374 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L20
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	if v374 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v380 < int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v422 != 0 {
		goto L105
	} else {
		goto L106
	}
L96:
	;
	v399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
	v401 = int32(1)
	if v225&v401 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v384+(v380^int32(-1))<<(uint(int32(2))%32))))
	v398 = v390
	goto L96
L98:
	;
	goto L99
L99:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v398 = v392 + v380<<(uint(int32(13))%32) + int32(-8192)
	goto L96
L100:
	;
	v404 = int32(3)
	goto L102
L101:
	;
	v404 = v401
	goto L102
L102:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	v412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242+v247<<(uint(int32(1))%32)-int32(2)))))
	F_spgPageIndexMultiDelete(m, v19+int32(108), v398, v238, v399, v404, int32(3), v406, v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L20
	} else {
		goto L103
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v398))) = base.I64_rotr(v215, int64(32))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L20
	} else {
		goto L104
	}
L104:
	;
	goto L95
L105:
	;
	F_UnlockReleaseBuffer(m, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L20
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v428 = F_XLogReadBufferForRedo(m, l0, int32(2), v19+int32(104))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L20
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	if v428 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v432 < int32(0) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	goto L112
L112:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v476 == int32(0) {
		goto L6
	} else {
		goto L119
	}
L113:
	;
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+6)))
	v452 = int32(2)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v450+v451<<(uint(v452)%32))+20))
	v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242+v247<<(uint(int32(1))%32)-v452))))
	F_spgUpdateNodeLink(m, v455&int32(_a_F_spg_redo_3)+v450, v459, v460, v466)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L20
	} else {
		goto L117
	}
L114:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v436+(v432^int32(-1))<<(uint(int32(2))%32))))
	v450 = v442
	goto L113
L115:
	;
	goto L116
L116:
	;
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v450 = v444 + v432<<(uint(int32(13))%32) + int32(-8192)
	goto L113
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v450))) = base.I64_rotr(v215, int64(32))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L20
	} else {
		goto L118
	}
L118:
	;
	goto L112
L119:
	;
	F_UnlockReleaseBuffer(m, v476)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L20
	} else {
		goto L120
	}
L120:
	;
	goto L6
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v494
	v498 = v482 + int32(20)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+72))
	if int32(0) < v500 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v551 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v551, v551, v551, v19+int32(100))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L20
	} else {
		goto L141
	}
L123:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+128)))
	if v503 != 0 {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v507 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L20
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	if v507 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v511 < int32(0) {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	goto L130
L130:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v546 == int32(0) {
		goto L6
	} else {
		goto L139
	}
L131:
	;
	v530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482))))
	F_PageIndexTupleDelete(m, v529, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L20
	} else {
		goto L135
	}
L132:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v515+(v511^int32(-1))<<(uint(int32(2))%32))))
	v529 = v521
	goto L131
L133:
	;
	goto L134
L134:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v529 = v523 + v511<<(uint(int32(13))%32) + int32(-8192)
	goto L131
L135:
	;
	v533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482))))
	v535 = F_PageAddItemExtended(m, v529, v498, v483, v533, int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L20
	} else {
		goto L136
	}
L136:
	;
	v537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482))))
	if v535 != v537 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v529))) = base.I64_rotr(v481, int64(32))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L20
	} else {
		goto L138
	}
L138:
	;
	goto L130
L139:
	;
	F_UnlockReleaseBuffer(m, v546)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L20
	} else {
		goto L140
	}
L140:
	;
	goto L6
L141:
	;
	v559 = int32(0)
	F_XLogRecGetBlockTag(m, l0, int32(1), v559, v559, v19+int32(96))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L20
	} else {
		goto L142
	}
L142:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+4)))
	if v565 == int32(1) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v650 != 0 {
		goto L166
	} else {
		goto L167
	}
L144:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v605 < int32(0) {
		goto L157
	} else {
		goto L158
	}
L145:
	;
	v569 = F_XLogInitBufferForRedo(m, l0, int32(1))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L20
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v602 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(104))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L20
	} else {
		goto L154
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v569
	v572 = int32(0)
	if v569 < v572 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L144
L150:
	;
	F_PageInit(m, v590, int32(_a_F_spg_redo_1), int32(8))
	mBase = m.M
	v594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v590)+16)))
	v595 = v590 + v594
	v596 = int32(_a_F_spg_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v595)+6)) = uint16(v596)
	*(*uint16)(unsafe.Add(mBase, uint32(v595))) = uint16(v572)
	goto L149
L151:
	;
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v576+(v569^int32(-1))<<(uint(int32(2))%32))))
	v590 = v582
	goto L150
L152:
	;
	goto L153
L153:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v590 = v584 + v569<<(uint(int32(13))%32) + int32(-8192)
	goto L150
L154:
	;
	if v602 != 0 {
		goto L143
	} else {
		goto L155
	}
L155:
	;
	goto L144
L156:
	;
	v624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+2)))
	F_addOrReplaceTuple(m, v623, v498, v483, v624)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L20
	} else {
		goto L160
	}
L157:
	;
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v609+(v605^int32(-1))<<(uint(int32(2))%32))))
	v623 = v615
	goto L156
L158:
	;
	goto L159
L159:
	;
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v623 = v617 + v605<<(uint(int32(13))%32) + int32(-8192)
	goto L156
L160:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+5)))
	if v627 == int32(1) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v630 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+6)))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v623+v630<<(uint(int32(2))%32))+20))
	v638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+8)))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v640 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+2)))
	F_spgUpdateNodeLink(m, v623+v634&int32(_a_F_spg_redo_3), v638, v639, v640)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L20
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623))) = base.I64_rotr(v481, int64(32))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L20
	} else {
		goto L165
	}
L164:
	;
	goto L163
L165:
	;
	goto L143
L166:
	;
	F_UnlockReleaseBuffer(m, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L20
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v656 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L20
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	if v656 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v660 < int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	goto L173
L173:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v791 != 0 {
		goto L201
	} else {
		goto L202
	}
L174:
	;
	if v485&int32(1) != 0 {
		goto L179
	} else {
		goto L180
	}
L175:
	;
	v664 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v664+(v660^int32(-1))<<(uint(int32(2))%32))))
	v678 = v670
	goto L174
L176:
	;
	goto L177
L177:
	;
	v672 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v678 = v672 + v660<<(uint(int32(13))%32) + int32(-8192)
	goto L174
L178:
	;
	v744 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482))))
	F_PageIndexTupleDelete(m, v678, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L20
	} else {
		goto L190
	}
L179:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(108))+72))
	*(*int32)(unsafe.Add(mBase, uint32(v687))) = int32(67)
	v693 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v687)+4)))
	v695 = v693 & int32(_a_F_spg_redo_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v687)+4)) = uint16(v695)
	goto L184
L180:
	;
	goto L181
L181:
	;
	v713 = v19 + int32(108)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v716 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+2)))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v713)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = int32(65)
	v724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v718)+4)))
	v726 = v724 & int32(_a_F_spg_redo_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v718)+4)) = uint16(v726)
	goto L187
L182:
	;
	v743 = v687
	goto L178
L184:
	;
	goto L185
L185:
	;
	v706 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v687)+10)) = uint16(v706)
	*(*int32)(unsafe.Add(mBase, uint32(v687)+6)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v687)+12)) = v706
	goto L182
L186:
	;
	v743 = v718
	goto L178
L187:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v718)+10)) = uint16(v716)
	*(*uint16)(unsafe.Add(mBase, uint32(v718)+8)) = uint16(v715)
	v733 = int32(base.Ui32(v715) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v718)+6)) = uint16(v733)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v713)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v718)+12)) = v735
	goto L186
L190:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	v750 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482))))
	v752 = F_PageAddItemExtended(m, v678, v743, int32(base.Ui32(v747)>>(uint(int32(2))%32)), v750, int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L20
	} else {
		goto L191
	}
L191:
	;
	v754 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482))))
	if v752 != v754 {
		goto L3
	} else {
		goto L192
	}
L192:
	;
	v756 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v678)+16)))
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)))
	if v760 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v761 = int32(4)
	goto L195
L194:
	;
	v761 = int32(2)
	goto L195
L195:
	;
	v762 = v678 + v756 + v761
	v763 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v762))))
	v765 = v763 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v762))) = uint16(v765)
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+5)))
	if v767 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v770 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+6)))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v678+v770<<(uint(int32(2))%32))+20))
	v778 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+8)))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v780 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+2)))
	F_spgUpdateNodeLink(m, v678+v774&int32(_a_F_spg_redo_3), v778, v779, v780)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L20
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v678))) = base.I64_rotr(v481, int64(32))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v786)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L20
	} else {
		goto L200
	}
L199:
	;
	goto L198
L200:
	;
	goto L173
L201:
	;
	F_UnlockReleaseBuffer(m, v791)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L20
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+5)))
	if v794 != int32(2) {
		goto L6
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	v800 = F_XLogReadBufferForRedo(m, l0, int32(2), v19+int32(104))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L20
	} else {
		goto L206
	}
L206:
	;
	if v800 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v804 < int32(0) {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	goto L209
L209:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v843 == int32(0) {
		goto L6
	} else {
		goto L216
	}
L210:
	;
	v823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+6)))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v822+v823<<(uint(int32(2))%32))+20))
	v831 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+8)))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v833 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+2)))
	F_spgUpdateNodeLink(m, v827&int32(_a_F_spg_redo_3)+v822, v831, v832, v833)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L20
	} else {
		goto L214
	}
L211:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v808+(v804^int32(-1))<<(uint(int32(2))%32))))
	v822 = v814
	goto L210
L212:
	;
	goto L213
L213:
	;
	v816 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v822 = v816 + v804<<(uint(int32(13))%32) + int32(-8192)
	goto L210
L214:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v822))) = base.I64_rotr(v481, int64(32))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v839)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L20
	} else {
		goto L215
	}
L215:
	;
	goto L209
L216:
	;
	F_UnlockReleaseBuffer(m, v843)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L20
	} else {
		goto L217
	}
L217:
	;
	goto L6
L218:
	;
	v934 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(108))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L20
	} else {
		goto L241
	}
L219:
	;
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849)+4)))
	if v856 == int32(1) {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v925 == int32(0) {
		goto L218
	} else {
		goto L239
	}
L221:
	;
	if v896 < int32(0) {
		goto L234
	} else {
		goto L235
	}
L222:
	;
	v860 = F_XLogInitBufferForRedo(m, l0, int32(1))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L20
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v893 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(108))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L20
	} else {
		goto L231
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v860
	v863 = int32(0)
	if v860 < v863 {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	v896 = v860
	goto L221
L227:
	;
	F_PageInit(m, v881, int32(_a_F_spg_redo_1), int32(8))
	mBase = m.M
	v885 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v881)+16)))
	v886 = v881 + v885
	v887 = int32(_a_F_spg_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v886)+6)) = uint16(v887)
	*(*uint16)(unsafe.Add(mBase, uint32(v886))) = uint16(v863)
	goto L226
L228:
	;
	v867 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v867+(v860^int32(-1))<<(uint(int32(2))%32))))
	v881 = v873
	goto L227
L229:
	;
	goto L230
L230:
	;
	v875 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v881 = v875 + v860<<(uint(int32(13))%32) + int32(-8192)
	goto L227
L231:
	;
	if v893 != 0 {
		goto L220
	} else {
		goto L232
	}
L232:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	v896 = v895
	goto L221
L233:
	;
	v915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v849)+2)))
	F_addOrReplaceTuple(m, v914, v853, v854, v915)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L20
	} else {
		goto L237
	}
L234:
	;
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v900+(v896^int32(-1))<<(uint(int32(2))%32))))
	v914 = v906
	goto L233
L235:
	;
	goto L236
L236:
	;
	v908 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v914 = v908 + v896<<(uint(int32(13))%32) + int32(-8192)
	goto L233
L237:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v914))) = base.I64_rotr(v848, int64(32))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v921)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L20
	} else {
		goto L238
	}
L238:
	;
	goto L220
L239:
	;
	F_UnlockReleaseBuffer(m, v925)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L20
	} else {
		goto L240
	}
L240:
	;
	goto L218
L241:
	;
	if v934 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v938 < int32(0) {
		goto L246
	} else {
		goto L247
	}
L243:
	;
	goto L244
L244:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v979 == int32(0) {
		goto L6
	} else {
		goto L257
	}
L245:
	;
	v957 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v849))))
	F_PageIndexTupleDelete(m, v956, v957)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L20
	} else {
		goto L249
	}
L246:
	;
	v942 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v942+(v938^int32(-1))<<(uint(int32(2))%32))))
	v956 = v948
	goto L245
L247:
	;
	goto L248
L248:
	;
	v950 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v956 = v950 + v938<<(uint(int32(13))%32) + int32(-8192)
	goto L245
L249:
	;
	v960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v849))))
	v962 = F_PageAddItemExtended(m, v956, v851, v852, v960, int32(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L20
	} else {
		goto L250
	}
L250:
	;
	v964 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v849))))
	if v962 != v964 {
		goto L2
	} else {
		goto L251
	}
L251:
	;
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849)+5)))
	if v966 == int32(1) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v969 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v849)+2)))
	F_addOrReplaceTuple(m, v956, v853, v854, v969)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L20
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v956))) = base.I64_rotr(v848, int64(32))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v975)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L20
	} else {
		goto L256
	}
L255:
	;
	goto L254
L256:
	;
	goto L244
L257:
	;
	F_UnlockReleaseBuffer(m, v979)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L20
	} else {
		goto L258
	}
L258:
	;
	goto L6
L259:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v985)+20))
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+24)))
	base.MemoryFill(m, v19+int32(108), int32(0), int32(84))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)) = uint8(v994)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v993
	v1003 = F_palloc0(m, int32(16))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L20
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v1003
	v1007 = v985 + int32(28)
	v1008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+2)))
	v1009 = int32(1)
	v1011 = v1007 + v1008<<(uint(v1009)%32)
	v1012 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+4)))
	v1015 = v1011 + v1012<<(uint(v1009)%32)
	v1016 = v1015 + v1012
	v1017 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1016)+4)))
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985))))
	if v1018 == v1009 {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+72))
	if int32(0) < v1129 {
		goto L291
	} else {
		goto L292
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = int32(0)
	v1127 = v2
	goto L261
L263:
	;
	goto L264
L264:
	;
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+6)))
	if v1023 == int32(1) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1027 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L20
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v1084 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L20
	} else {
		goto L278
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v1027
	if v1027 < int32(0) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+11)))
	v1054 = (v1048<<(uint(int32(3))%32) | int32(4)) & int32(12)
	if v1027 < int32(0) {
		goto L275
	} else {
		goto L276
	}
L270:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1033+(v1027^int32(-1))<<(uint(int32(2))%32))))
	v1047 = v1039
	goto L269
L271:
	;
	goto L272
L272:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1047 = v1041 + v1027<<(uint(int32(13))%32) + int32(-8192)
	goto L269
L273:
	;
	v1127 = v1047
	goto L261
L274:
	;
	F_PageInit(m, v1072, int32(_a_F_spg_redo_1), int32(8))
	mBase = m.M
	v1076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1072)+16)))
	v1077 = v1072 + v1076
	v1078 = int32(_a_F_spg_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1077)+6)) = uint16(v1078)
	*(*uint16)(unsafe.Add(mBase, uint32(v1077))) = uint16(v1054)
	goto L273
L275:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1058+(v1027^int32(-1))<<(uint(int32(2))%32))))
	v1072 = v1064
	goto L274
L276:
	;
	goto L277
L277:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1072 = v1066 + v1027<<(uint(int32(13))%32) + int32(-8192)
	goto L274
L278:
	;
	if v1084 != 0 {
		v1127 = v2
		goto L261
	} else {
		goto L279
	}
L279:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1086 < int32(0) {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+2)))
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)))
	if v1106 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L281:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1090+(v1086^int32(-1))<<(uint(int32(2))%32))))
	v1104 = v1096
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1104 = v1098 + v1086<<(uint(int32(13))%32) + int32(-8192)
	goto L280
L284:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v19)+92))
	v1114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+8)))
	F_spgPageIndexMultiDelete(m, v19+int32(108), v1104, v1007, v1105, int32(1), int32(3), v1113, v1114)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L20
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1119 = int32(3)
	F_spgPageIndexMultiDelete(m, v19+int32(108), v1104, v1007, v1105, v1119, v1119, int32(-1), int32(0))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L20
	} else {
		goto L288
	}
L287:
	;
	v1127 = v1104
	goto L261
L288:
	;
	v1127 = v1104
	goto L261
L289:
	;
	v1218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+4)))
	if v1218 != 0 {
		goto L313
	} else {
		goto L314
	}
L290:
	;
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+7)))
	if v1135 == int32(1) {
		goto L295
	} else {
		goto L296
	}
L291:
	;
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1128)+128)))
	if v1132 != 0 {
		goto L290
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = int32(0)
	v1217 = v2
	goto L289
L294:
	;
	goto L293
L295:
	;
	v1139 = F_XLogInitBufferForRedo(m, l0, int32(1))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L20
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v1196 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(100))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L20
	} else {
		goto L308
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v1139
	if v1139 < int32(0) {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+11)))
	v1166 = (v1160<<(uint(int32(3))%32) | int32(4)) & int32(12)
	if v1139 < int32(0) {
		goto L305
	} else {
		goto L306
	}
L300:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1145+(v1139^int32(-1))<<(uint(int32(2))%32))))
	v1159 = v1151
	goto L299
L301:
	;
	goto L302
L302:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1159 = v1153 + v1139<<(uint(int32(13))%32) + int32(-8192)
	goto L299
L303:
	;
	v1217 = v1159
	goto L289
L304:
	;
	F_PageInit(m, v1184, int32(_a_F_spg_redo_1), int32(8))
	mBase = m.M
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1184)+16)))
	v1189 = v1184 + v1188
	v1190 = int32(_a_F_spg_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1189)+6)) = uint16(v1190)
	*(*uint16)(unsafe.Add(mBase, uint32(v1189))) = uint16(v1166)
	goto L303
L305:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1170+(v1139^int32(-1))<<(uint(int32(2))%32))))
	v1184 = v1176
	goto L304
L306:
	;
	goto L307
L307:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1184 = v1178 + v1139<<(uint(int32(13))%32) + int32(-8192)
	goto L304
L308:
	;
	if v1196 != 0 {
		v1217 = v2
		goto L289
	} else {
		goto L309
	}
L309:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	if v1198 < int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1202+(v1198^int32(-1))<<(uint(int32(2))%32))))
	v1217 = v1208
	goto L289
L311:
	;
	goto L312
L312:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1217 = v1210 + v1198<<(uint(int32(13))%32) + int32(-8192)
	goto L289
L313:
	;
	v1222 = v1016 + v1017
	v1224 = int32(0)
	v1228 = v1218
	goto L316
L314:
	;
	goto L315
L315:
	;
	if v1127 != 0 {
		goto L326
	} else {
		goto L327
	}
L316:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1222)))
	v1239 = int32(base.Ui32(v1237) >> (uint(int32(2)) % 32))
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1224+v1015))))
	if v1241 != 0 {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	goto L315
L318:
	;
	v1242 = v1217
	goto L320
L319:
	;
	v1242 = v1127
	goto L320
L320:
	;
	if v1242 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1011+v1224<<(uint(int32(1))%32)))))
	F_addOrReplaceTuple(m, v1242, v1222, v1239, v1246)
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L20
	} else {
		goto L324
	}
L322:
	;
	v1250 = v1228
	goto L323
L323:
	;
	v1253 = v1224 + int32(1)
	if base.Ui32(v1253) < base.Ui32(v1250&int32(_a_F_spg_redo_6)) {
		v1222 = v1222 + v1239
		v1224 = v1253
		v1228 = v1250
		goto L316
	} else {
		goto L325
	}
L324:
	;
	v1249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+4)))
	v1250 = v1249
	goto L323
L325:
	;
	goto L317
L326:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1127))) = base.I64_rotr(v984, int64(32))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v1276)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L20
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	if v1217 != 0 {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	goto L328
L330:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1217))) = base.I64_rotr(v984, int64(32))
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	F_MarkBufferDirty(m, v1282)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L20
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+10)))
	if v1285 == int32(1) {
		goto L336
	} else {
		goto L337
	}
L333:
	;
	goto L332
L334:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	if v1374 != 0 {
		goto L357
	} else {
		goto L358
	}
L335:
	;
	if v1329 < int32(0) {
		goto L348
	} else {
		goto L349
	}
L336:
	;
	v1289 = F_XLogInitBufferForRedo(m, l0, int32(2))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L20
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	v1326 = F_XLogReadBufferForRedo(m, l0, int32(2), v19+int32(96))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L20
	} else {
		goto L345
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v1289
	v1292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+11)))
	v1296 = v1292 << (uint(int32(3)) % 32) & int32(8)
	if v1289 < int32(0) {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	v1329 = v1289
	goto L335
L341:
	;
	F_PageInit(m, v1314, int32(_a_F_spg_redo_1), int32(8))
	mBase = m.M
	v1318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1314)+16)))
	v1319 = v1314 + v1318
	v1320 = int32(_a_F_spg_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1319)+6)) = uint16(v1320)
	*(*uint16)(unsafe.Add(mBase, uint32(v1319))) = uint16(v1296)
	goto L340
L342:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1300+(v1289^int32(-1))<<(uint(int32(2))%32))))
	v1314 = v1306
	goto L341
L343:
	;
	goto L344
L344:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1314 = v1308 + v1289<<(uint(int32(13))%32) + int32(-8192)
	goto L341
L345:
	;
	if v1326 != 0 {
		goto L334
	} else {
		goto L346
	}
L346:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v1329 = v1328
	goto L335
L347:
	;
	v1348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+8)))
	F_addOrReplaceTuple(m, v1347, v1016, v1017, v1348)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L20
	} else {
		goto L351
	}
L348:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1333+(v1329^int32(-1))<<(uint(int32(2))%32))))
	v1347 = v1339
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1347 = v1341 + v1329<<(uint(int32(13))%32) + int32(-8192)
	goto L347
L351:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+12)))
	if v1351 == int32(1) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+14)))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1347+v1354<<(uint(int32(2))%32))+20))
	v1362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+16)))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v19)+92))
	v1364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+8)))
	F_spgUpdateNodeLink(m, v1347+v1358&int32(_a_F_spg_redo_3), v1362, v1363, v1364)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L20
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1347))) = base.I64_rotr(v984, int64(32))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	F_MarkBufferDirty(m, v1370)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L20
	} else {
		goto L356
	}
L355:
	;
	goto L354
L356:
	;
	goto L334
L357:
	;
	F_UnlockReleaseBuffer(m, v1374)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L20
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1377 != 0 {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	goto L359
L361:
	;
	F_UnlockReleaseBuffer(m, v1377)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L20
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	if v1380 != 0 {
		goto L365
	} else {
		goto L366
	}
L364:
	;
	goto L363
L365:
	;
	F_UnlockReleaseBuffer(m, v1380)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L20
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+72))
	if v1384 < int32(3) {
		goto L6
	} else {
		goto L369
	}
L368:
	;
	goto L367
L369:
	;
	v1387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1383)+232)))
	if v1387 != int32(1) {
		goto L6
	} else {
		goto L370
	}
L370:
	;
	v1393 = F_XLogReadBufferForRedo(m, l0, int32(3), v19+int32(88))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L20
	} else {
		goto L371
	}
L371:
	;
	if v1393 == int32(0) {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v1397 < int32(0) {
		goto L376
	} else {
		goto L377
	}
L373:
	;
	goto L374
L374:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v1436 == int32(0) {
		goto L6
	} else {
		goto L381
	}
L375:
	;
	v1416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+14)))
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1415+v1416<<(uint(int32(2))%32))+20))
	v1424 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+16)))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v19)+92))
	v1426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+8)))
	F_spgUpdateNodeLink(m, v1420&int32(_a_F_spg_redo_3)+v1415, v1424, v1425, v1426)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L20
	} else {
		goto L379
	}
L376:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1401+(v1397^int32(-1))<<(uint(int32(2))%32))))
	v1415 = v1407
	goto L375
L377:
	;
	goto L378
L378:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1415 = v1409 + v1397<<(uint(int32(13))%32) + int32(-8192)
	goto L375
L379:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1415))) = base.I64_rotr(v984, int64(32))
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	F_MarkBufferDirty(m, v1432)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L20
	} else {
		goto L380
	}
L380:
	;
	goto L374
L381:
	;
	F_UnlockReleaseBuffer(m, v1436)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L20
	} else {
		goto L382
	}
L382:
	;
	goto L6
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v1453
	v1456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442))))
	v1457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442)+2)))
	v1458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442)+4)))
	v1459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442)+6)))
	v1463 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L20
	} else {
		goto L384
	}
L384:
	;
	if v1463 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1468 = v1442 + int32(16)
	v1469 = int32(1)
	v1471 = v1468 + v1456<<(uint(v1469)%32)
	v1474 = v1471 + v1457<<(uint(v1469)%32)
	v1476 = v1458 << (uint(v1469) % 32)
	v1477 = v1474 + v1476
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1478 < int32(0) {
		goto L389
	} else {
		goto L390
	}
L386:
	;
	goto L387
L387:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1664 == int32(0) {
		goto L6
	} else {
		goto L408
	}
L388:
	;
	v1497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442))))
	v1498 = int32(2)
	F_spgPageIndexMultiDelete(m, v1446, v1496, v1468, v1497, v1498, v1498, int32(-1), int32(0))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L20
	} else {
		goto L392
	}
L389:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1482+(v1478^int32(-1))<<(uint(int32(2))%32))))
	v1496 = v1488
	goto L388
L390:
	;
	goto L391
L391:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1496 = v1490 + v1478<<(uint(int32(13))%32) + int32(-8192)
	goto L388
L392:
	;
	v1506 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442)+2)))
	v1507 = int32(3)
	F_spgPageIndexMultiDelete(m, v19+int32(108), v1496, v1471, v1506, v1507, v1507, int32(-1), int32(0))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L20
	} else {
		goto L393
	}
L393:
	;
	v1514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442)+4)))
	if v1514 != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1516 = v1496 + int32(20)
	v1521 = int32(0)
	goto L397
L395:
	;
	v1554 = int32(0)
	goto L396
L396:
	;
	v1572 = int32(3)
	F_spgPageIndexMultiDelete(m, v19+int32(108), v1496, v1474, v1554, v1572, v1572, int32(-1), int32(0))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L20
	} else {
		goto L400
	}
L397:
	;
	v1534 = int32(1)
	v1535 = v1521 << (uint(v1534) % 32)
	v1537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1474+v1535))))
	v1538 = int32(2)
	v1540 = v1516 + v1537<<(uint(v1538)%32)
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1540)))
	v1543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1535+v1477))))
	v1546 = v1516 + v1543<<(uint(v1538)%32)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1546)))
	*(*int32)(unsafe.Add(mBase, uint32(v1540))) = v1547
	*(*int32)(unsafe.Add(mBase, uint32(v1546))) = v1541
	v1551 = v1521 + v1534
	v1552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442)+4)))
	if base.Ui32(v1551) < base.Ui32(v1552) {
		v1521 = v1551
		goto L397
	} else {
		goto L399
	}
L398:
	;
	v1554 = v1552
	goto L396
L399:
	;
	goto L398
L400:
	;
	v1578 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442)+6)))
	if v1578 != 0 {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1579 = v1477 + v1476
	v1589 = int32(0)
	goto L404
L402:
	;
	goto L403
L403:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1496))) = base.I64_rotr(v1441, int64(32))
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v1645)
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L20
	} else {
		goto L407
	}
L404:
	;
	v1602 = int32(1)
	v1603 = v1589 << (uint(v1602) % 32)
	v1605 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1579+v1603))))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1496+int32(20)+v1605<<(uint(int32(2))%32))))
	v1612 = v1496 + v1609&int32(_a_F_spg_redo_3)
	v1614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1579+v1459<<(uint(int32(1))%32)+v1603))))
	v1617 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1612)+4)))
	v1620 = v1614&int32(_a_F_spg_redo_4) | v1617&int32(_a_F_spg_redo_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v1612)+4)) = uint16(v1620)
	v1623 = v1589 + v1602
	v1624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1442)+6)))
	if base.Ui32(v1623) < base.Ui32(v1624) {
		v1589 = v1623
		goto L404
	} else {
		goto L406
	}
L405:
	;
	goto L403
L406:
	;
	goto L405
L407:
	;
	goto L387
L408:
	;
	F_UnlockReleaseBuffer(m, v1664)
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L20
	} else {
		goto L409
	}
L409:
	;
	goto L6
L410:
	;
	if v1674 == int32(0) {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v1680 < int32(0) {
		goto L415
	} else {
		goto L416
	}
L412:
	;
	goto L413
L413:
	;
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v1710 == int32(0) {
		goto L6
	} else {
		goto L420
	}
L414:
	;
	v1699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669))))
	F_PageIndexMultiDelete(m, v1698, v1669+int32(12), v1699)
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L20
	} else {
		goto L418
	}
L415:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1684+(v1680^int32(-1))<<(uint(int32(2))%32))))
	v1698 = v1690
	goto L414
L416:
	;
	goto L417
L417:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1698 = v1692 + v1680<<(uint(int32(13))%32) + int32(-8192)
	goto L414
L418:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1698))) = base.I64_rotr(v1670, int64(32))
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v1705)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L20
	} else {
		goto L419
	}
L419:
	;
	goto L413
L420:
	;
	F_UnlockReleaseBuffer(m, v1710)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L20
	} else {
		goto L421
	}
L421:
	;
	goto L6
L422:
	;
	v1721 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v1721, v19+int32(108), v1721, v1721)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L20
	} else {
		goto L425
	}
L423:
	;
	goto L424
L424:
	;
	v1742 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(108))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L20
	} else {
		goto L427
	}
L425:
	;
	v1728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1715)+8)))
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1715)+4))
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v19)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v1730
	v1732 = *(*int64)(unsafe.Add(mBase, uint32(v19)+108))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v1732
	F_ResolveRecoveryConflictWithSnapshot(m, v1729, v1728, v19+int32(72))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L20
	} else {
		goto L426
	}
L426:
	;
	goto L424
L427:
	;
	if v1742 == int32(0) {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v1746 < int32(0) {
		goto L432
	} else {
		goto L433
	}
L429:
	;
	goto L430
L430:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v2037 == int32(0) {
		goto L6
	} else {
		goto L467
	}
L431:
	;
	v1765 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1764)+16)))
	v1766 = v1765 + v1764
	v1767 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715))))
	if v1767 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L432:
	;
	v1750 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[3]))
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1750+(v1746^int32(-1))<<(uint(int32(2))%32))))
	v1764 = v1756
	goto L431
L433:
	;
	goto L434
L434:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, _c_F_spg_redo[4]))
	v1764 = v1758 + v1746<<(uint(int32(13))%32) + int32(-8192)
	goto L431
L435:
	;
	v1831 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1766)+2)))
	v1832 = v1831 - v1816
	*(*uint16)(unsafe.Add(mBase, uint32(v1766)+2)) = uint16(v1832)
	v1834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1766)+4)))
	v1835 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715))))
	v1836 = v1834 + v1835
	*(*uint16)(unsafe.Add(mBase, uint32(v1766)+4)) = uint16(v1836)
	v1838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715)+2)))
	if v1838 != 0 {
		goto L442
	} else {
		goto L443
	}
L436:
	;
	v1816 = int32(0)
	goto L435
L437:
	;
	goto L438
L438:
	;
	v1779 = int32(0)
	goto L439
L439:
	;
	v1792 = int32(1)
	v1795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715+int32(10)+v1779<<(uint(v1792)%32)))))
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1764+int32(20)+v1795<<(uint(int32(2))%32))))
	v1802 = v1764 + v1799&int32(_a_F_spg_redo_3)
	v1803 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1802)+10)) = uint16(v1803)
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+6)) = int32(-1)
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1802)))
	*(*int32)(unsafe.Add(mBase, uint32(v1802))) = v1807 | int32(3)
	v1812 = v1779 + v1792
	v1813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715))))
	if base.Ui32(v1812) < base.Ui32(v1813) {
		v1779 = v1812
		goto L439
	} else {
		goto L441
	}
L440:
	;
	v1816 = v1813
	goto L435
L441:
	;
	goto L440
L442:
	;
	v1839 = int32(0)
	v1840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1764)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1840) {
		goto L445
	} else {
		goto L446
	}
L443:
	;
	goto L444
L444:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1764))) = base.I64_rotr(v1716, int64(32))
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v2018)
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L20
	} else {
		goto L466
	}
L445:
	;
	v1850 = int32(base.Ui32(v1840+int32(_a_F_spg_redo_7))>>(uint(int32(2))%32)) & int32(_a_F_spg_redo_6)
	goto L447
L446:
	;
	v1850 = v1839
	goto L447
L447:
	;
	v1853 = F_palloc(m, v1850<<(uint(int32(1))%32))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L20
	} else {
		goto L448
	}
L448:
	;
	v1855 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715)+2)))
	if base.Ui32(v1855) <= base.Ui32(v1850) {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v1857 = v1850 - v1855
	v1861 = (v1857 + int32(1)) & int32(3)
	if v1861 != 0 {
		goto L452
	} else {
		goto L453
	}
L450:
	;
	v1974 = v1855
	goto L451
L451:
	;
	v1989 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1766)+4)))
	v1992 = v1850 - v1974 + int32(1)
	v1993 = v1989 - v1992
	*(*uint16)(unsafe.Add(mBase, uint32(v1766)+4)) = uint16(v1993)
	F_PageIndexMultiDelete(m, v1764, v1853, v1992)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L20
	} else {
		goto L464
	}
L452:
	;
	v1863 = v1855
	v1869 = v1839
	goto L455
L453:
	;
	v1890 = v1855
	goto L454
L454:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v1857) {
		goto L458
	} else {
		goto L459
	}
L455:
	;
	v1878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715)+2)))
	v1880 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1853+(v1863-v1878)<<(uint(v1880)%32)))) = uint16(v1863)
	v1885 = v1863 + v1880
	v1887 = v1869 + v1880
	if v1887 != v1861 {
		v1863 = v1885
		v1869 = v1887
		goto L455
	} else {
		goto L457
	}
L456:
	;
	v1890 = v1885
	goto L454
L457:
	;
	goto L456
L458:
	;
	v1908 = v1890
	goto L461
L459:
	;
	goto L460
L460:
	;
	v1972 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715)+2)))
	v1974 = v1972
	goto L451
L461:
	;
	v1923 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715)+2)))
	v1925 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1853+(v1908-v1923)<<(uint(v1925)%32)))) = uint16(v1908)
	v1930 = v1908 + v1925
	v1931 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1853+(v1930-v1931)<<(uint(v1925)%32)))) = uint16(v1930)
	v1938 = v1908 + int32(2)
	v1939 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1853+(v1938-v1939)<<(uint(v1925)%32)))) = uint16(v1938)
	v1946 = v1908 + int32(3)
	v1947 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1715)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1853+(v1946-v1947)<<(uint(v1925)%32)))) = uint16(v1946)
	if v1850 != v1946 {
		v1908 = v1908 + int32(4)
		goto L461
	} else {
		goto L463
	}
L462:
	;
	goto L460
L463:
	;
	goto L462
L464:
	;
	F_pfree(m, v1853)
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L20
	} else {
		goto L465
	}
L465:
	;
	goto L444
L466:
	;
	goto L430
L467:
	;
	F_UnlockReleaseBuffer(m, v2037)
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L20
	} else {
		goto L468
	}
L468:
	;
	goto L6
L469:
	;
	m.G0 = v19 + int32(192)
	return
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v133
	F_errmsg_internal(m, int32(_a_F_spg_redo_8), v19+int32(16))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L20
	} else {
		goto L471
	}
L471:
	;
	F_errfinish(m, int32(_a_F_spg_redo_9), int32(135), int32(_a_F_spg_redo_10))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L20
	} else {
		goto L472
	}
L472:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v483
	F_errmsg_internal(m, int32(_a_F_spg_redo_8), v19+int32(48))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L20
	} else {
		goto L474
	}
L474:
	;
	F_errfinish(m, int32(_a_F_spg_redo_9), int32(316), int32(_a_F_spg_redo_11))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L20
	} else {
		goto L475
	}
L475:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L476:
	;
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = int32(base.Ui32(v2101) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(_a_F_spg_redo_8), v19+int32(32))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L20
	} else {
		goto L477
	}
L477:
	;
	F_errfinish(m, int32(_a_F_spg_redo_9), int32(397), int32(_a_F_spg_redo_11))
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L20
	} else {
		goto L478
	}
L478:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v852
	F_errmsg_internal(m, int32(_a_F_spg_redo_8), v19-int32(-64))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L20
	} else {
		goto L480
	}
L480:
	;
	F_errfinish(m, int32(_a_F_spg_redo_9), int32(514), int32(_a_F_spg_redo_12))
	mBase = m.M
	v2129 = m.ExcPending
	if v2129 != 0 {
		goto L20
	} else {
		goto L481
	}
L481:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	F_errmsg_internal(m, int32(_a_F_spg_redo_13), v19)
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L20
	} else {
		goto L483
	}
L483:
	;
	F_errfinish(m, int32(_a_F_spg_redo_9), int32(968), int32(_a_F_spg_redo_14))
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L20
	} else {
		goto L484
	}
L484:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
