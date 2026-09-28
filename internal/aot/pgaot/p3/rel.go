package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_rel_agg_info(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v644 int32
	_ = v644
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
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v827 int32
	_ = v827
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int64
	_ = v1084
	var v1085 int64
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1227 int32
	_ = v1227
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1306 int32
	_ = v1306
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1386 float64
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1389 float64
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1393 float64
	_ = v1393
	var v1394 float64
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1402 int32
	_ = v1402
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v24))|base.B2i32(int32(1)<<(uint(v24)%32)&int32(44) == v4) == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(32)
	return v1402
L2:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+240))
	if v37 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v63 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v1402 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+236))
	v42 = F_adjust_appendrel_attrs_multilevel(m, l0, v41, l1, v36)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v46 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = v46
	if l2 == v46 {
		v1402 = v42
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v52 = int32(0)
	v54 = F_estimate_num_groups(m, l0, v50, v51, v52, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v42)+24)) = v54
	v58 = *(*float64)(unsafe.Add(mBase, _c_F_create_rel_agg_info[0]))
	v59 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v61 = base.F64_le(v58, base.F64_div(v59, v54))
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)) = uint8(v61)
	v1402 = v42
	goto L1
L12:
	;
	if int32(0) <= v120 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v120 = base.I32_ctz(v106) | v107<<(uint(int32(5))%32)
	goto L12
L14:
	;
	v120 = int32(-2)
	goto L12
L15:
	;
	v71 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v74 <= v71 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v77 = v63 + int32(8)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v84 = v81 & int32(-1)
	if v84 != 0 {
		v106 = v84
		v107 = v71
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v85 = int32(1)
	if v85 == v74 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v89 = v85
	goto L19
L19:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77+v89<<(uint(int32(2))%32))))
	if v96 != 0 {
		v106 = v96
		v107 = v89
		goto L13
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	v98 = v89 + int32(1)
	if v98 != v74 {
		v89 = v98
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v127 = v120
	goto L26
L24:
	;
	goto L25
L25:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v309 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L26:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v142) <= base.Ui32(v127) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L25
L28:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v231 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L29:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v148)+104))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v175 = int32(0)
	if v173 == v175 {
		goto L40
	} else {
		goto L41
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L36
	}
L31:
	;
	v145 = v127 << (uint(int32(2)) % 32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v145+v146)))
	if v148 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149+v145)))
	if v151 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	if v154 != int32(2) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151)+44))
	if v157 != 0 {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v127
	F_errmsg_internal(m, int32(_a_F_create_rel_agg_info_0), v22)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_create_rel_agg_info_1), int32(606), int32(_a_F_create_rel_agg_info_2))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	if v228 != 0 {
		goto L28
	} else {
		goto L53
	}
L40:
	;
	v228 = int32(1)
	goto L39
L41:
	;
	goto L42
L42:
	;
	if v174 == int32(0) {
		v221 = v175
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v228 = v221
	goto L39
L44:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v185 < v184 {
		v221 = v175
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v187 = int32(1)
	if v184 <= v187 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v190 = v187
	goto L48
L47:
	;
	v190 = v184
	goto L48
L48:
	;
	v191 = int32(8)
	v196 = int32(0)
	goto L49
L49:
	;
	v203 = v196 << (uint(int32(2)) % 32)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v173+v191+v203)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v174+v191+v203)))
	v210 = v205 & (v207 ^ int32(-1))
	v212 = base.B2i32(v210 == int32(0))
	if v210 != 0 {
		v221 = v212
		goto L43
	} else {
		goto L51
	}
L50:
	;
	v221 = v212
	goto L43
L51:
	;
	v214 = v196 + int32(1)
	if v214 != v190 {
		v196 = v214
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v1402 = int32(0)
	goto L1
L54:
	;
	if int32(0) <= v287 {
		v127 = v287
		goto L26
	} else {
		goto L65
	}
L55:
	;
	v287 = base.I32_ctz(v273) | v274<<(uint(int32(5))%32)
	goto L54
L56:
	;
	v287 = int32(-2)
	goto L54
L57:
	;
	v238 = v127 + int32(1)
	v240 = int32(base.Ui32(v238) >> (uint(int32(5)) % 32))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v241 <= v240 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v244 = v231 + int32(8)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+v240<<(uint(int32(2))%32))))
	v251 = v248 & (int32(-1) << (uint(v238) % 32))
	if v251 != 0 {
		v273 = v251
		v274 = v240
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v253 = v240 + int32(1)
	if v253 == v241 {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v256 = v253
	goto L61
L61:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v244+v256<<(uint(int32(2))%32))))
	if v263 != 0 {
		v273 = v263
		v274 = v256
		goto L55
	} else {
		goto L63
	}
L62:
	;
	goto L56
L63:
	;
	v265 = v256 + int32(1)
	if v265 != v241 {
		v256 = v265
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L27
L66:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	if v476 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L67:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if v312 <= int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v320 = int32(0)
	goto L69
L69:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335+v320<<(uint(int32(2))%32))))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+20))
	if v340&int32(-2) != int32(4) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L66
L71:
	;
	v453 = v320 + int32(1)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if v453 < v454 {
		v320 = v453
		goto L69
	} else {
		goto L102
	}
L72:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
	v347 = int32(0)
	if base.B2i32(v345 == v347)|base.B2i32(v346 == v347) != 0 {
		v392 = v347
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v392 == int32(0) {
		goto L71
	} else {
		goto L86
	}
L74:
	;
	goto L73
L75:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v357 < v358 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v360 = v357
	goto L78
L77:
	;
	v360 = v358
	goto L78
L78:
	;
	if v360 <= int32(1) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v363 = int32(1)
	goto L81
L80:
	;
	v363 = v360
	goto L81
L81:
	;
	v364 = int32(8)
	v369 = int32(0)
	goto L82
L82:
	;
	v376 = v369 << (uint(int32(2)) % 32)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v346+v364+v376)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v345+v364+v376)))
	v381 = v378 & v380
	v383 = base.B2i32(v381 != int32(0))
	if v381 != 0 {
		v392 = v383
		goto L74
	} else {
		goto L84
	}
L83:
	;
	v392 = v383
	goto L74
L84:
	;
	v385 = v369 + int32(1)
	if v385 != v363 {
		v369 = v385
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v397 = int32(0)
	if v395 == v397 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v450 != 0 {
		goto L71
	} else {
		goto L101
	}
L88:
	;
	v450 = int32(1)
	goto L87
L89:
	;
	goto L90
L90:
	;
	if v396 == int32(0) {
		v443 = v397
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v450 = v443
	goto L87
L92:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v407 < v406 {
		v443 = v397
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v409 = int32(1)
	if v406 <= v409 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v412 = v409
	goto L96
L95:
	;
	v412 = v406
	goto L96
L96:
	;
	v413 = int32(8)
	v418 = int32(0)
	goto L97
L97:
	;
	v425 = v418 << (uint(int32(2)) % 32)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v395+v413+v425)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v396+v413+v425)))
	v432 = v427 & (v429 ^ int32(-1))
	v434 = base.B2i32(v432 == int32(0))
	if v432 != 0 {
		v443 = v434
		goto L91
	} else {
		goto L99
	}
L98:
	;
	v443 = v434
	goto L91
L99:
	;
	v436 = v418 + int32(1)
	if v436 != v412 {
		v418 = v436
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v1402 = int32(0)
	goto L1
L102:
	;
	goto L70
L103:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v533 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L104:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if v479 <= int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v476)+12))
	v483 = int32(0)
	v489 = v483
	goto L106
L106:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v482+v489<<(uint(int32(2))%32))))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	if v508 == int32(321) {
		v1402 = v483
		goto L1
	} else {
		goto L108
	}
L107:
	;
	goto L103
L108:
	;
	v512 = v489 + int32(1)
	if v479 != v512 {
		v489 = v512
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v644 = F_create_empty_pathtarget(m)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L8
	} else {
		goto L133
	}
L111:
	;
	v536 = int32(0)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v537 <= v536 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v544 = v536
	goto L113
L113:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v559+v544<<(uint(int32(2))%32))))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+8))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v566 = int32(0)
	if v564 == v566 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v1402 = int32(0)
	goto L1
L115:
	;
	if v619 != 0 {
		goto L129
	} else {
		goto L130
	}
L116:
	;
	v619 = int32(1)
	goto L115
L117:
	;
	goto L118
L118:
	;
	if v565 == int32(0) {
		v612 = v566
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v619 = v612
	goto L115
L120:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v564)+4))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	if v576 < v575 {
		v612 = v566
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v578 = int32(1)
	if v575 <= v578 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v581 = v578
	goto L124
L123:
	;
	v581 = v575
	goto L124
L124:
	;
	v582 = int32(8)
	v587 = int32(0)
	goto L125
L125:
	;
	v594 = v587 << (uint(int32(2)) % 32)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v564+v582+v594)))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v565+v582+v594)))
	v601 = v596 & (v598 ^ int32(-1))
	v603 = base.B2i32(v601 == int32(0))
	if v601 != 0 {
		v612 = v603
		goto L119
	} else {
		goto L127
	}
L126:
	;
	v612 = v603
	goto L119
L127:
	;
	v605 = v587 + int32(1)
	if v605 != v581 {
		v587 = v605
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v621 = v544 + int32(1)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v621 < v622 {
		v544 = v621
		goto L113
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	goto L114
L132:
	;
	goto L110
L133:
	;
	v646 = F_create_empty_pathtarget(m)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	v648 = int32(0)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v650 == v648 {
		v764 = v648
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v778)+4))
	if v779 == int32(0) {
		v1402 = v648
		goto L1
	} else {
		goto L163
	}
L136:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v650)+4))
	if v653 <= int32(0) {
		v764 = v648
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v657 = v653 & int32(3)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v650)+12))
	v659 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v653) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v668 = v659
	v669 = v648
	v681 = v4
	goto L141
L139:
	;
	v713 = v659
	v714 = v648
	goto L140
L140:
	;
	v732 = v713
	v733 = v714
	v735 = v4
	goto L157
L141:
	;
	v685 = v658 + v668<<(uint(int32(2))%32)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v685)+12))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v686)+16))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v685)+8))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)+16))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v685)+4))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+16))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+16))
	if base.Ui32(v669) < base.Ui32(v693) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	if v657 == int32(0) {
		v764 = v701
		goto L135
	} else {
		goto L156
	}
L143:
	;
	v695 = v693
	goto L145
L144:
	;
	v695 = v669
	goto L145
L145:
	;
	if base.Ui32(v695) < base.Ui32(v691) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v697 = v691
	goto L148
L147:
	;
	v697 = v695
	goto L148
L148:
	;
	if base.Ui32(v697) < base.Ui32(v689) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v699 = v689
	goto L151
L150:
	;
	v699 = v697
	goto L151
L151:
	;
	if base.Ui32(v699) < base.Ui32(v687) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v701 = v687
	goto L154
L153:
	;
	v701 = v699
	goto L154
L154:
	;
	v702 = int32(4)
	v703 = v668 + v702
	v705 = v681 + v702
	if v705 != v653&int32(2147483644) {
		v668 = v703
		v669 = v701
		v681 = v705
		goto L141
	} else {
		goto L155
	}
L155:
	;
	goto L142
L156:
	;
	v713 = v703
	v714 = v701
	goto L140
L157:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v658+v732<<(uint(int32(2))%32))))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)+16))
	if base.Ui32(v733) < base.Ui32(v751) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v764 = v753
	goto L135
L159:
	;
	v753 = v751
	goto L161
L160:
	;
	v753 = v733
	goto L161
L161:
	;
	v754 = int32(1)
	v757 = v735 + v754
	if v757 != v657 {
		v732 = v732 + v754
		v733 = v753
		v735 = v757
		goto L157
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	v782 = int32(0)
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	if v783 <= v782 {
		v1402 = v648
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v786 = int32(0)
	v794 = v764
	v795 = v786
	v799 = v782
	v800 = v786
	v805 = v786
	goto L165
L165:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v779)+12))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v808+v805<<(uint(int32(2))%32))))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v813 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	if v1205 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L167:
	;
	v1215 = v805 + int32(1)
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	if v1215 < v1216 {
		v794 = v1200
		v795 = v1201
		v799 = v1205
		v800 = v1206
		v805 = v1215
		goto L165
	} else {
		goto L258
	}
L168:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v973 = F_bms_copy(m, v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L8
	} else {
		goto L196
	}
L169:
	;
	v816 = int32(0)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	if v817 <= v816 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v827 = v816
	goto L171
L171:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v813)+12))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v839+v827<<(uint(int32(2))%32))))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v843)+4))
	v845 = F_equal(m, v812, v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L8
	} else {
		goto L175
	}
L172:
	;
	goto L168
L173:
	;
	v950 = v827 + int32(1)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	if v950 < v951 {
		v827 = v950
		goto L171
	} else {
		goto L195
	}
L174:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v843)+8))
	if v914 == int32(0) {
		goto L168
	} else {
		goto L187
	}
L175:
	;
	if v845 != 0 {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v843)+12))
	if v847 == int32(0) {
		goto L173
	} else {
		goto L177
	}
L177:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v847)+36))
	v852 = F_bms_is_member(m, v850, v851)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L8
	} else {
		goto L178
	}
L178:
	;
	if v852 == int32(0) {
		goto L173
	} else {
		goto L179
	}
L179:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v843)+12))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v856)+16))
	if v857 == int32(0) {
		goto L173
	} else {
		goto L180
	}
L180:
	;
	v860 = int32(0)
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v857)+4))
	if v861 <= v860 {
		goto L173
	} else {
		goto L181
	}
L181:
	;
	v867 = v860
	goto L182
L182:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v857)+12))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v883+v867<<(uint(int32(2))%32))))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v887)+4))
	v889 = F_equal(m, v812, v888)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L8
	} else {
		goto L184
	}
L183:
	;
	goto L173
L184:
	;
	if v889 != 0 {
		goto L174
	} else {
		goto L185
	}
L185:
	;
	v892 = v867 + int32(1)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v857)+4))
	if v892 < v893 {
		v867 = v892
		goto L182
	} else {
		goto L186
	}
L186:
	;
	goto L183
L187:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v918 = F_get_sortgroupref_clause(m, v914, v917)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L8
	} else {
		goto L188
	}
L188:
	;
	F_add_column_to_pathtarget(m, v644, v812, v914)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L8
	} else {
		goto L189
	}
L189:
	;
	F_add_column_to_pathtarget(m, v646, v812, v914)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L8
	} else {
		goto L190
	}
L190:
	;
	v924 = F_list_member(m, v795, v918)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L8
	} else {
		goto L191
	}
L191:
	;
	if v924 != 0 {
		v1200 = v794
		v1201 = v795
		v1205 = v799
		v1206 = v800
		goto L167
	} else {
		goto L192
	}
L192:
	;
	v926 = F_lappend(m, v795, v918)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L8
	} else {
		goto L193
	}
L193:
	;
	v928 = F_lappend(m, v800, v812)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L8
	} else {
		goto L194
	}
L194:
	;
	v1200 = v794
	v1201 = v926
	v1205 = v799
	v1206 = v928
	goto L167
L195:
	;
	goto L172
L196:
	;
	v976 = F_bms_add_member(m, v973, int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L8
	} else {
		goto L197
	}
L197:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v978) < base.Ui32(v979) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v985)+92))
	v1003 = int32(*(*int16)(unsafe.Add(mBase, uint32(v812)+8)))
	v1004 = int32(*(*int16)(unsafe.Add(mBase, uint32(v985)+88)))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1002+(v1003-v1004)<<(uint(int32(2))%32))))
	if v1009 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L199:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v981+v978<<(uint(int32(2))%32))))
	if v985 != 0 {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L8
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v978
	F_errmsg_internal(m, int32(_a_F_create_rel_agg_info_0), v22+int32(16))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L8
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_create_rel_agg_info_1), int32(556), int32(_a_F_create_rel_agg_info_3))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L8
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	if v1064 != 0 {
		goto L220
	} else {
		goto L221
	}
L207:
	;
	v1064 = int32(0)
	goto L206
L208:
	;
	goto L209
L209:
	;
	v1017 = int32(1)
	if v976 == int32(0) {
		v1054 = v1017
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1064 = v1054
	goto L206
L211:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+4))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v976)+4))
	if v1021 < v1020 {
		v1054 = v1017
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v1023 = int32(1)
	if v1020 <= v1023 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1026 = v1023
	goto L215
L214:
	;
	v1026 = v1020
	goto L215
L215:
	;
	v1027 = int32(8)
	v1032 = int32(0)
	goto L216
L216:
	;
	v1039 = v1032 << (uint(int32(2)) % 32)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1009+v1027+v1039)))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v976+v1027+v1039)))
	v1046 = v1041 & (v1043 ^ int32(-1))
	v1048 = base.B2i32(v1046 != int32(0))
	if v1046 != 0 {
		v1054 = v1048
		goto L210
	} else {
		goto L218
	}
L217:
	;
	v1054 = v1048
	goto L210
L218:
	;
	v1050 = v1032 + int32(1)
	if v1050 != v1026 {
		v1032 = v1050
		goto L216
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	v1065 = int32(0)
	v1066 = F_exprType(m, v812)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L8
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v1121 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L223:
	;
	v1069 = F_lookup_type_cache(m, v1066, int32(512))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L8
	} else {
		goto L224
	}
L224:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1069)+36))
	if v1071 == int32(0) {
		v1402 = v1065
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1069)+40))
	if v1074 == int32(0) {
		v1402 = v1065
		goto L1
	} else {
		goto L226
	}
L226:
	;
	v1078 = F_get_opfamily_proc(m, v1071, v1074, v1074, int32(4))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L8
	} else {
		goto L227
	}
L227:
	;
	if v1078 == int32(0) {
		v1402 = v1065
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v1082 = F_exprCollation(m, v812)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L8
	} else {
		goto L229
	}
L229:
	;
	v1084 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1069)+40)))
	v1085 = F_OidFunctionCall1Coll(m, v1078, v1082, v1084)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L8
	} else {
		goto L230
	}
L230:
	;
	if v1085 == int64(0) {
		v1402 = v1065
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v1090 = F_palloc0(m, int32(20))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L8
	} else {
		goto L232
	}
L232:
	;
	v1093 = v794 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+4)) = v1093
	*(*int32)(unsafe.Add(mBase, uint32(v1090))) = int32(106)
	v1097 = F_exprType(m, v812)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L8
	} else {
		goto L233
	}
L233:
	;
	v1099 = int32(0)
	F_get_sort_group_operators(m, v1097, v1099, int32(1), v1099, v1090+int32(12), v1090+int32(8), v1099, v1090+int32(18))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L8
	} else {
		goto L234
	}
L234:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	F_add_column_to_pathtarget(m, v644, v812, v1111)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L8
	} else {
		goto L235
	}
L235:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	F_add_column_to_pathtarget(m, v646, v812, v1114)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L8
	} else {
		goto L236
	}
L236:
	;
	v1117 = F_lappend(m, v795, v1090)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L8
	} else {
		goto L237
	}
L237:
	;
	v1119 = F_lappend(m, v800, v812)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L8
	} else {
		goto L238
	}
L238:
	;
	v1200 = v1093
	v1201 = v1117
	v1205 = v799
	v1206 = v1119
	goto L167
L239:
	;
	v1193 = F_lappend(m, v799, v812)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L8
	} else {
		goto L257
	}
L240:
	;
	v1124 = int32(0)
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+4))
	if v1125 <= v1124 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v1131 = v1124
	goto L242
L242:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+12))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1148+v1131<<(uint(int32(2))%32))))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+8))
	v1154 = F_bms_is_member(m, v1147, v1153)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L8
	} else {
		goto L245
	}
L243:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v1170 = F_list_member(m, v1169, v812)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L8
	} else {
		goto L254
	}
L244:
	;
	goto L243
L245:
	;
	if v1154 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+4))
	v1158 = F_pull_var_clause(m, v1156, int32(42))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L8
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v1166 = v1131 + int32(1)
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+4))
	if v1166 < v1167 {
		v1131 = v1166
		goto L242
	} else {
		goto L253
	}
L249:
	;
	v1160 = F_list_member(m, v1158, v812)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L8
	} else {
		goto L250
	}
L250:
	;
	F_list_free(m, v1158)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L8
	} else {
		goto L251
	}
L251:
	;
	if v1160 != 0 {
		goto L244
	} else {
		goto L252
	}
L252:
	;
	goto L248
L253:
	;
	goto L239
L254:
	;
	if v1170 != 0 {
		goto L239
	} else {
		goto L255
	}
L255:
	;
	F_add_new_column_to_pathtarget(m, v646, v812)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L8
	} else {
		goto L256
	}
L256:
	;
	v1200 = v794
	v1201 = v795
	v1205 = v799
	v1206 = v800
	goto L167
L257:
	;
	v1200 = v794
	v1201 = v795
	v1205 = v1193
	v1206 = v800
	goto L167
L258:
	;
	goto L166
L259:
	;
	v1402 = int32(0)
	goto L1
L260:
	;
	if v1201 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L261:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+4))
	if v1220 <= int32(0) {
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1227 = int32(0)
	goto L263
L263:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = int32(0)
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1247 = int32(2)
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1243+v1227<<(uint(v1247)%32))))
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+4))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1246+v1251<<(uint(v1247)%32))))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+16))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+28))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v644)+4))
	v1261 = F_check_functional_grouping(m, v1256, v1251, v1257, v1258, v22+int32(28))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L8
	} else {
		goto L265
	}
L264:
	;
	goto L260
L265:
	;
	if v1261 == int32(0) {
		goto L259
	} else {
		goto L266
	}
L266:
	;
	F_add_new_column_to_pathtarget(m, v644, v1250)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L8
	} else {
		goto L267
	}
L267:
	;
	F_add_new_column_to_pathtarget(m, v646, v1250)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L8
	} else {
		goto L268
	}
L268:
	;
	v1270 = v1227 + int32(1)
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+4))
	if v1270 < v1271 {
		v1227 = v1270
		goto L263
	} else {
		goto L269
	}
L269:
	;
	goto L264
L270:
	;
	v1402 = int32(0)
	goto L1
L271:
	;
	goto L272
L272:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v1295 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1369 = F_set_pathtarget_cost_width(m, l0, v644)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L8
	} else {
		goto L291
	}
L274:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+4))
	if v1298 <= int32(0) {
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v1306 = int32(0)
	goto L276
L276:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+12))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1321+v1306<<(uint(int32(2))%32))))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1325)+4))
	v1327 = F_copyObjectImpl(m, v1326)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L8
	} else {
		goto L278
	}
L277:
	;
	goto L273
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+56)) = int32(6)
	goto L280
L279:
	;
	F_add_column_to_pathtarget(m, v644, v1327, int32(0))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L8
	} else {
		goto L289
	}
L280:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1327)+20))
	if v1334 == int32(2281) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	goto L279
L283:
	;
	v1337 = int32(17)
	goto L285
L284:
	;
	v1337 = v1334
	goto L285
L285:
	;
	goto L286
L286:
	;
	goto L288
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+8)) = v1337
	goto L282
L289:
	;
	v1347 = v1306 + int32(1)
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+4))
	if v1347 < v1348 {
		v1306 = v1347
		goto L276
	} else {
		goto L290
	}
L290:
	;
	goto L277
L291:
	;
	v1371 = F_set_pathtarget_cost_width(m, l0, v646)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L8
	} else {
		goto L292
	}
L292:
	;
	v1374 = F_palloc0(m, int32(40))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L8
	} else {
		goto L293
	}
L293:
	;
	v1376 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+20)) = v1376
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+16)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+12)) = v1201
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+8)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+4)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v1374))) = int32(271)
	if l2 == v1376 {
		v1402 = v1374
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v1386 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v1387 = int32(0)
	v1389 = F_estimate_num_groups(m, l0, v1206, v1386, v1387, v1387)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L8
	} else {
		goto L295
	}
L295:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1374)+24)) = v1389
	v1393 = *(*float64)(unsafe.Add(mBase, _c_F_create_rel_agg_info[0]))
	v1394 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v1396 = base.F64_le(v1393, base.F64_div(v1394, v1389))
	*(*uint8)(unsafe.Add(mBase, uint32(v1374)+32)) = uint8(v1396)
	v1402 = v1374
	goto L1
}
func F_get_rel_type_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+72))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func F_parseRelOptionsInternal(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 float64
	_ = v259
	var v260 float64
	_ = v260
	var v264 float64
	_ = v264
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 float64
	_ = v284
	var v285 float64
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	v17 = m.G0
	v19 = v17 - int32(224)
	m.G0 = v19
	v21 = base.I32_wrap_i64(l0)
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_deconstruct_array_builtin(m, v22, int32(25), v19+int32(216), int32(0), v19+int32(212))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+212))
	if int32(0) < v32 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L162
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L158
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L154
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L150
	}
L8:
	;
	v36 = l1 ^ int32(1)
	v49 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v19)+216))
	F_pfree(m, v555)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L145
	}
L11:
	;
	v53 = int32(0)
	if l3 <= v53 {
		v488 = v53
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v494 = int32(0)
	if base.B2i32(l1 == v494)|base.B2i32(v488 < l3) == v494 {
		goto L129
	} else {
		goto L130
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)+216))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v49<<(uint(int32(3))%32))))
	v61 = int32(4)
	v62 = v60 + v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v67 = int32(base.Ui32(v63)>>(uint(int32(2))%32)) - v61
	v78 = v53
	goto L15
L15:
	;
	v86 = l2 + v78<<(uint(int32(4))%32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	if v67 <= v88 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v488 = l3
	goto L13
L17:
	;
	v476 = v78 + int32(1)
	if v476 != l3 {
		v78 = v476
		goto L15
	} else {
		goto L128
	}
L18:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v62))))
	if v91 != int32(61) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v139 != 0 {
		goto L17
	} else {
		goto L33
	}
L21:
	;
	v139 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v100 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v101 = v62
	v102 = v94
	v103 = v88
	v104 = v100
	goto L28
L25:
	;
	v127 = v94
	v131 = int32(0)
	goto L26
L26:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v139 = v131 - v132
	goto L20
L27:
	;
	v127 = v122
	v131 = v124
	goto L26
L28:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if base.B2i32(v104 != v106)|base.B2i32(v106 == int32(0)) != 0 {
		v122 = v102
		v124 = v104
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v122 = v116
	v124 = int32(0)
	goto L27
L30:
	;
	v112 = v103 - int32(1)
	if v112 == int32(0) {
		v122 = v102
		v124 = v104
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v115 = int32(1)
	v116 = v102 + v115
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v117 != 0 {
		v101 = v101 + v115
		v102 = v116
		v103 = v112
		v104 = v117
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	if l1 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)))
	if v140&int32(1) != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v143 = v67 - v88
	v144 = F_palloc(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v147 = v143 - int32(1)
	if v147 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	base.MemoryCopy(m, v144, v62+v149+int32(1), v147)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v144+v147))) = uint8(v155)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
	switch v158 {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2:
		goto L50
	case 3:
		goto L49
	case 4:
		goto L48
	case 5:
		goto L47
	default:
		goto L46
	}
L42:
	;
	F_pfree(m, v144)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L127
	}
L43:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v454 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)) = uint8(v454)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v453
	goto L42
L44:
	;
	v449 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)) = uint8(v449)
	F_pfree(m, v144)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L126
	}
L45:
	;
	if v438&int32(1) == int32(0) {
		goto L42
	} else {
		goto L125
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L122
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v144
	if l1 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L48:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v157)+24))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	if v299 != 0 {
		goto L86
	} else {
		goto L87
	}
L49:
	;
	v247 = v86 + int32(8)
	v248 = int32(0)
	v250 = F_parse_real(m, v144, v247, v248, v248)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L74
	}
L50:
	;
	v201 = v86 + int32(8)
	v202 = int32(0)
	v204 = F_parse_int(m, v144, v201, v202, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L62
	}
L51:
	;
	v189 = F_strlen(m, v144)
	mBase = m.M
	v190 = F_parse_bool_with_len(m, v144, v189, v19+int32(223))
	mBase = m.M
	goto L59
L52:
	;
	v161 = F_strlen(m, v144)
	mBase = m.M
	v162 = F_parse_bool_with_len(m, v144, v161, v86+int32(8))
	mBase = m.M
	goto L53
L53:
	;
	if (v36|v162)&int32(1) != 0 {
		v438 = v162
		goto L45
	} else {
		goto L54
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v174
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_0), v19+int32(48))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1716), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+223)))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v191
	if (v190|v36)&int32(1) == int32(0) {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	if v190&int32(1) != 0 {
		goto L44
	} else {
		goto L61
	}
L61:
	;
	goto L42
L62:
	;
	if (v36|v204)&int32(1) == int32(0) {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	if l1 == int32(0) {
		v438 = v204
		goto L45
	} else {
		goto L64
	}
L64:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v157)+28))
	if v214 <= v213 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v157)+32))
	if v213 <= v216 {
		v438 = v204
		goto L45
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v144
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_3), v19+int32(96))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v157)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v234
	v239 = F_errdetail(m, int32(_a_F_parseRelOptionsInternal_4), v19+int32(80))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1750), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	if (v36|v250)&int32(1) == int32(0) {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	if l1 == int32(0) {
		v438 = v250
		goto L45
	} else {
		goto L76
	}
L76:
	;
	v259 = *(*float64)(unsafe.Add(mBase, uint32(v247)))
	v260 = *(*float64)(unsafe.Add(mBase, uint32(v157)+32))
	if base.F64_lt(v259, v260) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v264 = *(*float64)(unsafe.Add(mBase, uint32(v157)+40))
	if base.F64_gt(v259, v264) == int32(0) {
		v438 = v250
		goto L45
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v144
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_3), v19+int32(144))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v284 = *(*float64)(unsafe.Add(mBase, uint32(v157)+32))
	v285 = *(*float64)(unsafe.Add(mBase, uint32(v157)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+136)) = v285
	*(*float64)(unsafe.Add(mBase, uint32(v19)+128)) = v284
	v291 = F_errdetail(m, int32(_a_F_parseRelOptionsInternal_5), v19+int32(128))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1770), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
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
	v308 = v298
	v309 = v299
	goto L89
L87:
	;
	goto L88
L88:
	;
	if l1 != 0 {
		goto L106
	} else {
		goto L107
	}
L89:
	;
	v318 = v144
	v319 = v309
	goto L92
L90:
	;
	goto L88
L91:
	;
	if v356 == int32(0) {
		goto L43
	} else {
		goto L104
	}
L92:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v322 == v323 {
		v345 = v322
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v356 = int32(0)
	goto L91
L94:
	;
	v347 = int32(1)
	if v345 != 0 {
		v318 = v318 + v347
		v319 = v319 + v347
		goto L92
	} else {
		goto L103
	}
L95:
	;
	if base.Ui32((v322-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v333 = v322 | int32(32)
	goto L98
L97:
	;
	v333 = v322
	goto L98
L98:
	;
	if base.Ui32((v323-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v342 = v323 | int32(32)
	goto L101
L100:
	;
	v342 = v323
	goto L101
L101:
	;
	if v333 == v342 {
		v345 = v333
		goto L94
	} else {
		goto L102
	}
L102:
	;
	v356 = v333 - v342
	goto L91
L103:
	;
	goto L93
L104:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	if v359 != 0 {
		v308 = v308 + int32(8)
		v309 = v359
		goto L89
	} else {
		goto L105
	}
L105:
	;
	goto L90
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v157)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v406
	F_pfree(m, v144)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L117
	}
L109:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+196)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v386
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_6), v19+int32(192))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v157)+32))
	if v394 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v394
	F_errdetail_internal(m, int32(_a_F_parseRelOptionsInternal_7), v19+int32(176))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1794), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	v488 = v78
	goto L13
L118:
	;
	v419 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)) = uint8(v419)
	v488 = v78
	goto L13
L119:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v157)+32))
	if v413 == int32(0) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	m.T0[v413].(func(*base.Module, int32))(m, v144)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v426
	F_errmsg_internal(m, int32(_a_F_parseRelOptionsInternal_8), v19+int32(16))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1816), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	goto L44
L126:
	;
	v488 = v78
	goto L13
L127:
	;
	v488 = v78
	goto L13
L128:
	;
	goto L16
L129:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v19)+216))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500+v49<<(uint(int32(3))%32))))
	v505 = F_text_to_cstring(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v536 = v49 + int32(1)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v19)+212))
	if v536 < v537 {
		v49 = v536
		goto L11
	} else {
		goto L144
	}
L132:
	;
	v507 = int32(61)
	v508 = F___strchrnul(m, v505, v507)
	mBase = m.M
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	if v510 == v507 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v514 != 0 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v514 = v508
	goto L136
L135:
	;
	v514 = int32(0)
	goto L136
L136:
	;
	goto L133
L137:
	;
	v515 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v514))) = uint8(v515)
	goto L139
L138:
	;
	goto L139
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v505
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_9), v19+int32(32))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1587), int32(_a_F_parseRelOptionsInternal_10))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	goto L12
L145:
	;
	if v22 != v21 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	F_pfree(m, v22)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	m.G0 = v19 + int32(224)
	return
L149:
	;
	goto L148
L150:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v572
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_11), v19)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1700), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v590
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_0), v19-int32(-64))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1730), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+116)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v611
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_12), v19+int32(112))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1742), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+164)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v632
	F_errmsg(m, int32(_a_F_parseRelOptionsInternal_13), v19+int32(160))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_parseRelOptionsInternal_1), int32(1762), int32(_a_F_parseRelOptionsInternal_2))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_set_rel_width(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v96 int32
	_ = v96
	var v97 float64
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 float64
	_ = v158
	var v159 int32
	_ = v159
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v164 int32
	_ = v164
	var v165 float64
	_ = v165
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v238 int64
	_ = v238
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v272 int64
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v288 int64
	_ = v288
	var v294 int64
	_ = v294
	var v295 int64
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v324 int64
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int64
	_ = v331
	var v334 int64
	_ = v334
	var v349 int64
	_ = v349
	var v352 int32
	_ = v352
	var v353 int64
	_ = v353
	var v356 int64
	_ = v356
	v3 = int32(0)
	v13 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v38 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v37)+16)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v38
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v44 == int32(0) {
		v349 = v13
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v34 = v20 + v21<<(uint(int32(2))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v34 = v27 + v28<<(uint(int32(2))%32) - int32(4)
	goto L1
L5:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v353 = int64(1073741823)
	if v353 <= v349 {
		goto L63
	} else {
		goto L64
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if int32(0) < v47 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v51 = v18 + int32(16)
	v57 = v3
	v62 = v3
	v64 = v13
	goto L10
L8:
	;
	v192 = v3
	v194 = v13
	goto L9
L9:
	;
	if v192 == int32(0) {
		v349 = v194
		goto L5
	} else {
		goto L38
	}
L10:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v57<<(uint(int32(2))%32))))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72 != int32(6) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v192 = v173
	v194 = v175
	goto L9
L12:
	;
	v179 = v57 + int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v179 < v180 {
		v57 = v179
		v62 = v173
		v64 = v175
		goto L10
	} else {
		goto L37
	}
L13:
	;
	v143 = F_exprType(m, v71)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
	} else {
		goto L33
	}
L14:
	;
	if v72 != int32(321) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v101 != v102 {
		goto L13
	} else {
		goto L21
	}
L17:
	;
	v77 = F_find_placeholder_info(m, l0, v71)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v79 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+24)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l0
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v82
	v88 = F_cost_qual_eval_walker(m, v80, v18+int32(8))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v18)+16))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v91)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v91)+16)) = base.F64_add(v92, v93)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v97 = *(*float64)(unsafe.Add(mBase, uint32(v96)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v96)+24)) = base.F64_add(v90, v97)
	v173 = v62
	v175 = v64 + v79
	goto L12
L21:
	;
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+8)))
	if v104 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v173 = int32(1)
	v175 = v64
	goto L12
L23:
	;
	goto L24
L24:
	;
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+88)))
	v111 = (v104 - v108) << (uint(int32(2)) % 32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112)))
	if int32(0) < v114 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v173 = v62
	v175 = v64 + base.I64_extend_i32_u(v114)
	goto L12
L26:
	;
	goto L27
L27:
	;
	v119 = int32(0)
	if base.B2i32(v36 == v119)|base.B2i32(v104 <= v119) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	v136 = F_get_typavgwidth(m, v134, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L18
	} else {
		goto L32
	}
L29:
	;
	v124 = F_get_attavgwidth(m, v36, v104)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	if v124 <= int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v128+v111))) = v124
	v173 = v62
	v175 = v64 + base.I64_extend_i32_u(v124)
	goto L12
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v111))) = v136
	v173 = v62
	v175 = v64 + base.I64_extend_i32_s(v136)
	goto L12
L33:
	;
	v145 = F_exprTypmod(m, v71)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	v147 = F_get_typavgwidth(m, v143, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l0
	v150 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v150
	v156 = F_cost_qual_eval_walker(m, v71, v18+int32(8))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v158 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v160 = *(*float64)(unsafe.Add(mBase, uint32(v18)+16))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v159)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v159)+16)) = base.F64_add(v160, v161)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v164)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v164)+24)) = base.F64_add(v158, v165)
	v173 = v62
	v175 = v64 + base.I64_extend_i32_s(v147)
	goto L12
L37:
	;
	goto L11
L38:
	;
	if v36 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v327 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+88)))
	v331 = int64(1073741823)
	if v331 <= v324 {
		goto L60
	} else {
		goto L61
	}
L40:
	;
	v201 = int32(1)
	v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+90)))
	if v202 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v302 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+88)))
	v306 = F_get_relation_data_width(m, v36, v301-v302<<(uint(int32(2))%32))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L18
	} else {
		goto L59
	}
L43:
	;
	v324 = int64(24)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v206 = int32(2)
	v209 = base.I32_extend16_s(v202 + int32(1))
	if v209 <= v206 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v212 = v206
	goto L48
L47:
	;
	v212 = v209
	goto L48
L48:
	;
	v214 = v212 - int32(1)
	v216 = v214 & int32(3)
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+88)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v219 = int64(24)
	if int32(5) <= v209 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v225 = v201
	v230 = int32(0)
	v238 = v219
	goto L52
L50:
	;
	v259 = v201
	v272 = v219
	goto L51
L51:
	;
	v275 = v259
	v280 = int32(0)
	v288 = v272
	goto L56
L52:
	;
	v243 = v218 + (v225-v217)<<(uint(int32(2))%32)
	v244 = int64(*(*int32)(unsafe.Add(mBase, uint32(v243))))
	v246 = int64(*(*int32)(unsafe.Add(mBase, uint32(v243)+4)))
	v248 = int64(*(*int32)(unsafe.Add(mBase, uint32(v243)+8)))
	v250 = int64(*(*int32)(unsafe.Add(mBase, uint32(v243)+12)))
	v251 = v238 + v244 + v246 + v248 + v250
	v252 = int32(4)
	v253 = v225 + v252
	v255 = v230 + v252
	if v255 != v214&int32(-4) {
		v225 = v253
		v230 = v255
		v238 = v251
		goto L52
	} else {
		goto L54
	}
L53:
	;
	if v216 == int32(0) {
		v324 = v251
		goto L39
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v259 = v253
	v272 = v251
	goto L51
L56:
	;
	v294 = int64(*(*int32)(unsafe.Add(mBase, uint32(v218+(v275-v217)<<(uint(int32(2))%32)))))
	v295 = v288 + v294
	v296 = int32(1)
	v299 = v280 + v296
	if v299 != v216 {
		v275 = v275 + v296
		v280 = v299
		v288 = v295
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v324 = v295
	goto L39
L58:
	;
	goto L57
L59:
	;
	v324 = base.I64_extend_i32_s(v306) + int64(24)
	goto L39
L60:
	;
	v334 = v331
	goto L62
L61:
	;
	v334 = v324
	goto L62
L62:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v326-v327<<(uint(int32(2))%32)))) = uint32(v334)
	v349 = v194 + v324
	goto L5
L63:
	;
	v356 = v353
	goto L65
L64:
	;
	v356 = v349
	goto L65
L65:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v352)+32)) = uint32(v356)
	m.G0 = v18 + int32(32)
	return
}
func F_transformRelOptions(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v576 int32
	_ = v576
	var v585 int32
	_ = v585
	var v586 int64
	_ = v586
	var v587 int32
	_ = v587
	var v588 int64
	_ = v588
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	if l1 == v7 {
		v588 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(48)
	return v588
L2:
	;
	v24 = base.I32_wrap_i64(l0)
	if v24 == int32(0) {
		v246 = v7
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v251 <= int32(0) {
		v576 = v246
		goto L51
	} else {
		goto L52
	}
L4:
	;
	v27 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int64(0)
L6:
	;
	F_deconstruct_array_builtin(m, v27, int32(25), v20+int32(44), int32(0), v20+int32(40))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v39 <= int32(0) {
		v246 = v7
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v51 = v7
	v53 = v39
	v54 = v7
	goto L9
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v59+v51<<(uint(int32(3))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v64 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v246 = v226
	goto L3
L11:
	;
	v232 = v51 + int32(1)
	if v232 < v225 {
		v51 = v232
		v53 = v225
		v54 = v226
		goto L9
	} else {
		goto L50
	}
L12:
	;
	v67 = base.I32_wrap_i64(v63)
	v68 = int32(4)
	v69 = v67 + v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v85 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_transformRelOptions[0]))
	v211 = F_accumArrayResult(m, v54, v63, int32(0), int32(25), v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L49
	}
L15:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v75+v85<<(uint(int32(2))%32))))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if l2 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L14
L17:
	;
	v188 = v85 + int32(1)
	if v188 != v64 {
		v85 = v188
		goto L15
	} else {
		goto L48
	}
L18:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v132 = F_strlen(m, v131)
	mBase = m.M
	if int32(base.Ui32(v70)>>(uint(int32(2))%32))-v68 <= v132 {
		goto L17
	} else {
		goto L32
	}
L19:
	;
	if v98 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v98 == int32(0) {
		goto L17
	} else {
		goto L23
	}
L22:
	;
	goto L17
L23:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v107 == int32(0))|base.B2i32(v107 != v110) != 0 {
		v128 = v107
		v129 = v110
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v128-v129 != 0 {
		goto L17
	} else {
		goto L31
	}
L25:
	;
	goto L24
L26:
	;
	v113 = v98
	v114 = l2
	goto L27
L27:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	if v118 == int32(0) {
		v128 = v118
		v129 = v117
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v128 = v118
	v129 = v117
	goto L25
L29:
	;
	v121 = int32(1)
	if v118 == v117 {
		v113 = v113 + v121
		v114 = v114 + v121
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L18
L32:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v69))))
	if v135 != int32(61) {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	if v132 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v182 == int32(0) {
		v225 = v53
		v226 = v54
		goto L11
	} else {
		goto L47
	}
L35:
	;
	v182 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v143 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v144 = v69
	v145 = v131
	v146 = v132
	v147 = v143
	goto L42
L39:
	;
	v170 = v131
	v174 = int32(0)
	goto L40
L40:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	v182 = v174 - v175
	goto L34
L41:
	;
	v170 = v165
	v174 = v167
	goto L40
L42:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if base.B2i32(v147 != v149)|base.B2i32(v149 == int32(0)) != 0 {
		v165 = v145
		v167 = v147
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v165 = v159
	v167 = int32(0)
	goto L41
L44:
	;
	v155 = v146 - int32(1)
	if v155 == int32(0) {
		v165 = v145
		v167 = v147
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v158 = int32(1)
	v159 = v145 + v158
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	if v160 != 0 {
		v144 = v144 + v158
		v145 = v159
		v146 = v155
		v147 = v160
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	goto L17
L48:
	;
	goto L16
L49:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	v225 = v213
	v226 = v211
	goto L11
L50:
	;
	goto L10
L51:
	;
	if v576 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L52:
	;
	v266 = int32(0)
	v267 = v246
	goto L54
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L5
	} else {
		goto L129
	}
L54:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v272+v266<<(uint(int32(2))%32))))
	if l5 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L5
	} else {
		goto L125
	}
L56:
	;
	goto L55
L57:
	;
	v507 = v266 + int32(1)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v507 < v508 {
		v266 = v507
		v267 = v501
		goto L54
	} else {
		goto L124
	}
L58:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	if v277 == int32(0) {
		v501 = v267
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v296 != 0 {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_transformRelOptions_0), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_transformRelOptions_1), int32(1341), int32(_a_F_transformRelOptions_2))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	if v400 != 0 {
		goto L95
	} else {
		goto L96
	}
L67:
	;
	if l3 == int32(0) {
		goto L53
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if l2 != 0 {
		v501 = v267
		goto L57
	} else {
		goto L94
	}
L70:
	;
	v299 = int32(0)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v300 == v299 {
		goto L53
	} else {
		goto L71
	}
L71:
	;
	v311 = v300
	v313 = v299
	goto L72
L72:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if base.B2i32(v322 == int32(0))|base.B2i32(v322 != v325) != 0 {
		v343 = v322
		v344 = v325
		goto L75
	} else {
		goto L76
	}
L73:
	;
	if l2 == int32(0) {
		v501 = v267
		goto L57
	} else {
		goto L85
	}
L74:
	;
	if v343-v344 != 0 {
		goto L81
	} else {
		goto L82
	}
L75:
	;
	goto L74
L76:
	;
	v328 = v296
	v329 = v311
	goto L77
L77:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+1)))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+1)))
	if v333 == int32(0) {
		v343 = v333
		v344 = v332
		goto L75
	} else {
		goto L79
	}
L78:
	;
	v343 = v333
	v344 = v332
	goto L75
L79:
	;
	v336 = int32(1)
	if v333 == v332 {
		v328 = v328 + v336
		v329 = v329 + v336
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v347 = v313 + int32(1)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l3+v347<<(uint(int32(2))%32))))
	if v351 != 0 {
		v311 = v351
		v313 = v347
		goto L72
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	goto L73
L84:
	;
	goto L53
L85:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v356 == int32(0))|base.B2i32(v356 != v359) != 0 {
		v377 = v356
		v378 = v359
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v377-v378 == int32(0) {
		goto L66
	} else {
		goto L93
	}
L87:
	;
	goto L86
L88:
	;
	v362 = v296
	v363 = l2
	goto L89
L89:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+1)))
	if v367 == int32(0) {
		v377 = v367
		v378 = v366
		goto L87
	} else {
		goto L91
	}
L90:
	;
	v377 = v367
	v378 = v366
	goto L87
L91:
	;
	v370 = int32(1)
	if v367 == v366 {
		v362 = v362 + v370
		v363 = v363 + v370
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v501 = v267
	goto L57
L94:
	;
	goto L66
L95:
	;
	v401 = F_defGetString(m, v276)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
	} else {
		goto L98
	}
L96:
	;
	v404 = int32(_a_F_transformRelOptions_3)
	goto L97
L97:
	;
	v405 = int32(61)
	v406 = F___strchrnul(m, v399, v405)
	mBase = m.M
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if v408 == v405 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v404 = v401
	goto L97
L99:
	;
	if v412 != 0 {
		goto L56
	} else {
		goto L103
	}
L100:
	;
	v412 = v406
	goto L102
L101:
	;
	v412 = int32(0)
	goto L102
L102:
	;
	goto L99
L103:
	;
	if l4 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v463 = F_strlen(m, v399)
	mBase = m.M
	v464 = F_strlen(m, v404)
	mBase = m.M
	v465 = v463 + v464
	v468 = F_palloc(m, v465+int32(6))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L121
	}
L105:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v415 != 0 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v416 = int32(_a_F_transformRelOptions_4)
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	v422 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformRelOptions[1])))
	if base.B2i32(v419 == int32(0))|base.B2i32(v419 != v422) != 0 {
		v440 = v419
		v441 = v422
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v440-v441 != 0 {
		goto L104
	} else {
		goto L114
	}
L108:
	;
	goto L107
L109:
	;
	v425 = v399
	v426 = v416
	goto L110
L110:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+1)))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+1)))
	if v430 == int32(0) {
		v440 = v430
		v441 = v429
		goto L108
	} else {
		goto L112
	}
L111:
	;
	v440 = v430
	v441 = v429
	goto L108
L112:
	;
	v433 = int32(1)
	if v430 == v429 {
		v425 = v425 + v433
		v426 = v426 + v433
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v443 = F_defGetBoolean(m, v276)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	if v443 == int32(0) {
		v501 = v267
		goto L57
	} else {
		goto L116
	}
L116:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_transformRelOptions_5), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_transformRelOptions_1), int32(1419), int32(_a_F_transformRelOptions_2))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468))) = v465<<(uint(int32(2))%32) + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v399
	v480 = F_pg_sprintf(m, v468+int32(4), int32(_a_F_transformRelOptions_6), v20)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _c_F_transformRelOptions[0]))
	v487 = F_accumArrayResult(m, v267, base.I64_extend_i32_u(v468), int32(0), int32(25), v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v501 = v487
	goto L57
L124:
	;
	v576 = v501
	goto L51
L125:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v399
	F_errmsg(m, int32(_a_F_transformRelOptions_7), v20+int32(16))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_transformRelOptions_1), int32(1405), int32(_a_F_transformRelOptions_2))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v552
	F_errmsg(m, int32(_a_F_transformRelOptions_8), v20+int32(32))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_transformRelOptions_1), int32(1375), int32(_a_F_transformRelOptions_2))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	v588 = int64(0)
	goto L1
L134:
	;
	goto L135
L135:
	;
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_transformRelOptions[0]))
	v586 = F_makeArrayResult(m, v576, v585)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v588 = v586
	goto L1
}
