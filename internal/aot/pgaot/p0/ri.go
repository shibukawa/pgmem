package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ri_restrict(m *base.Module, l0 int32, l1 int32) {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v273 int32
	_ = v273
	var v284 int32
	_ = v284
	var v301 int32
	_ = v301
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v702 int32
	_ = v702
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v802 int32
	_ = v802
	var v810 int32
	_ = v810
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v918 int32
	_ = v918
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v969 int32
	_ = v969
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1021 int32
	_ = v1021
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1064 int32
	_ = v1064
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1170 int32
	_ = v1170
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1195 int32
	_ = v1195
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1262 int32
	_ = v1262
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1302 int32
	_ = v1302
	var v1310 int32
	_ = v1310
	var v1316 int32
	_ = v1316
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1435 int32
	_ = v1435
	var v1442 int32
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1451 int32
	_ = v1451
	v17 = m.G0
	v19 = v17 - int32(1264)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_ri_FetchConstraintInfo(m, v21, v22, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	v28 = F_table_open(m, v26, int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L1
	} else {
		goto L271
	}
L5:
	;
	F_sequence_close(m, v28, int32(2))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L270
	}
L6:
	;
	v1407 = int32(1)
	v1411 = F_ri_PerformCheck(m, v24, v19+int32(984), v1388, v28, v31, v30, int32(0), l1^v1407, v1407, int32(5))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L1
	} else {
		goto L267
	}
L7:
	;
	if l1 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L169
	}
L9:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L108
	}
L10:
	;
	v516 = int32(6)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v516 = int32(5)
	goto L9
L14:
	;
	goto L15
L15:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+972)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+968)) = v40
	v46 = F_ri_FetchPreparedPlan(m, v19+int32(968))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v46 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_initStringInfo(m, v19+int32(416))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v490 = v46
	goto L20
L20:
	;
	v503 = int32(5)
	v506 = int32(0)
	v510 = F_ri_PerformCheck(m, v24, v19+int32(968), v490, v28, v31, v30, v506, v506, int32(1), v503)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L104
	}
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+119)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
	v57 = F_get_namespace_name(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v59 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+992)) = uint8(v59)
	v63 = v57
	v66 = v19 + int32(992)
	goto L23
L23:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v79 != int32(34) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v96 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+1)) = uint16(v96)
	v99 = v19 + int32(992)
	if v99&int32(3) == int32(0) {
		v123 = v99
		goto L33
	} else {
		goto L34
	}
L25:
	;
	goto L24
L26:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v91)
	v63 = v63 + int32(1)
	v66 = v92
	goto L23
L27:
	;
	if v79 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v86 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)) = uint8(v86)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v91 = v88
	v92 = v66 + int32(2)
	goto L26
L30:
	;
	v91 = v79
	v92 = v66 + int32(1)
	goto L26
L31:
	;
	v159 = v156 + (v19 + int32(992))
	v160 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v159))) = uint8(v160)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v164 = v159 + int32(1)
	v165 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v165)
	v169 = v164
	v172 = v162 + int32(4)
	goto L48
L32:
	;
	v156 = v148 - v99
	goto L31
L33:
	;
	v127 = v123
	goto L42
L34:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v107 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v156 = int32(0)
	goto L31
L36:
	;
	goto L37
L37:
	;
	v112 = v99
	goto L38
L38:
	;
	v116 = v112 + int32(1)
	if v116&int32(3) == int32(0) {
		v123 = v116
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v148 = v116
	goto L32
L40:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v121 != 0 {
		v112 = v116
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v136 = int32(-2139062144)
	if (int32(16843008)-v133|v133)&v136 == v136 {
		v127 = v127 + int32(4)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v142 = v127
	goto L45
L44:
	;
	goto L43
L45:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v146 != 0 {
		v142 = v142 + int32(1)
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v148 = v142
	goto L32
L47:
	;
	goto L46
L48:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v185 != int32(34) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v202 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v169)+1)) = uint16(v202)
	if v55&int32(255) == int32(112) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	goto L49
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v197)
	v169 = v198
	v172 = v172 + int32(1)
	goto L48
L52:
	;
	if v185 == int32(0) {
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v192 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)) = uint8(v192)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	v197 = v194
	v198 = v169 + int32(2)
	goto L51
L55:
	;
	v197 = v185
	v198 = v169 + int32(1)
	goto L51
L56:
	;
	v210 = int32(728204)
	goto L58
L57:
	;
	v210 = int32(714654)
	goto L58
L58:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v211 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if int32(0) < v301 {
		goto L74
	} else {
		goto L75
	}
L60:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v218 = int32(*(*int16)(unsafe.Add(mBase, uint32(v214<<(uint(int32(1))%32)+v24)+170)))
	v219 = F_attnumAttName(m, v31, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v19)+228)) = v19 + int32(992)
	F_appendStringInfo(m, v19+int32(416), int32(29108), v19+int32(224))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L73
	}
L63:
	;
	v221 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+704)) = uint8(v221)
	v225 = v219
	v228 = v19 + int32(704)
	goto L64
L64:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v241 != int32(34) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v258 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v228)+1)) = uint16(v258)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+212)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v19)+216)) = v19 + int32(992)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+208)) = v19 + int32(704)
	F_appendStringInfo(m, v19+int32(416), int32(29066), v19+int32(208))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L72
	}
L66:
	;
	goto L65
L67:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v254))) = uint8(v253)
	v225 = v225 + int32(1)
	v228 = v254
	goto L64
L68:
	;
	if v241 == int32(0) {
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v248 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)) = uint8(v248)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	v253 = v250
	v254 = v228 + int32(2)
	goto L67
L71:
	;
	v253 = v241
	v254 = v228 + int32(1)
	goto L67
L72:
	;
	goto L59
L73:
	;
	goto L59
L74:
	;
	v316 = int32(0)
	v320 = int32(527647)
	goto L77
L75:
	;
	goto L76
L76:
	;
	F_appendStringInfoString(m, v19+int32(416), int32(29129))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L93
	}
L77:
	;
	v328 = v24 + int32(172) + v316<<(uint(int32(1))%32)
	v329 = int32(*(*int16)(unsafe.Add(mBase, uint32(v328))))
	v330 = F_attnumTypeId(m, v31, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L76
L79:
	;
	v332 = int32(*(*int16)(unsafe.Add(mBase, uint32(v328))))
	v333 = F_attnumAttName(m, v31, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v335 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+704)) = uint8(v335)
	v339 = v333
	v342 = v19 + int32(704)
	goto L81
L81:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if v355 != int32(34) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v372 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v342)+1)) = uint16(v372)
	v375 = v316 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v375
	v382 = F_pg_sprintf(m, v19+int32(272), int32(456777), v19+int32(192))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L89
	}
L83:
	;
	goto L82
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v368))) = uint8(v367)
	v339 = v339 + int32(1)
	v342 = v368
	goto L81
L85:
	;
	if v355 == int32(0) {
		goto L83
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v362 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+1)) = uint8(v362)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	v367 = v364
	v368 = v342 + int32(2)
	goto L84
L88:
	;
	v367 = v355
	v368 = v342 + int32(1)
	goto L84
L89:
	;
	v385 = v316 << (uint(int32(2)) % 32)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(428)+v385)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v320
	F_appendStringInfo(m, v19+int32(416), int32(709145), v19+int32(176))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_generate_operator_clause(m, v19+int32(416), v19+int32(704), v330, v387, v19+int32(272), v330)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(560)+v385))) = v330
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v375 < v409 {
		v316 = v375
		v320 = int32(530859)
		goto L77
	} else {
		goto L92
	}
L92:
	;
	goto L78
L93:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v432 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v437 = int32(*(*int16)(unsafe.Add(mBase, uint32(v433<<(uint(int32(1))%32)+v24)+234)))
	v438 = F_attnumTypeId(m, v28, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v19)+416))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v485 = F_ri_PlanCheck(m, v479, v480, v19+int32(560), v19+int32(968), v28, v31)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L103
	}
L97:
	;
	F_appendStringInfoString(m, v19+int32(416), int32(715549))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v445
	v452 = F_pg_sprintf(m, v19+int32(272), int32(456777), v19+int32(160))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v24)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = int32(728204)
	F_appendStringInfo(m, v19+int32(416), int32(709145), v19+int32(144))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_generate_operator_clause(m, v19+int32(416), v19+int32(272), v438, v454, int32(330126), int32(4537))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_appendStringInfoString(m, v19+int32(416), int32(646922))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L96
L103:
	;
	v490 = v485
	goto L20
L104:
	;
	v512 = F_SPI_finish(m)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v512 != int32(2) {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	if v510 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v516 = v503
	goto L9
L108:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+988)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v19)+984)) = v535
	v540 = F_ri_FetchPreparedPlan(m, v19+int32(984))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	if v540 != 0 {
		v1388 = v540
		goto L6
	} else {
		goto L110
	}
L110:
	;
	F_initStringInfo(m, v19+int32(968))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+119)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v546)+68))
	v549 = F_get_namespace_name(m, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v551 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+704)) = uint8(v551)
	v555 = v549
	v558 = v19 + int32(704)
	goto L113
L113:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555))))
	if v571 != int32(34) {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v588 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v558)+1)) = uint16(v588)
	v591 = v19 + int32(704)
	if v591&int32(3) == int32(0) {
		v615 = v591
		goto L123
	} else {
		goto L124
	}
L115:
	;
	goto L114
L116:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v584))) = uint8(v583)
	v555 = v555 + int32(1)
	v558 = v584
	goto L113
L117:
	;
	if v571 == int32(0) {
		goto L115
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v578 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v558)+1)) = uint8(v578)
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555))))
	v583 = v580
	v584 = v558 + int32(2)
	goto L116
L120:
	;
	v583 = v571
	v584 = v558 + int32(1)
	goto L116
L121:
	;
	v651 = v648 + (v19 + int32(704))
	v652 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v651))) = uint8(v652)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	v656 = v651 + int32(1)
	v657 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v656))) = uint8(v657)
	v661 = v656
	v664 = v654 + int32(4)
	goto L138
L122:
	;
	v648 = v640 - v591
	goto L121
L123:
	;
	v619 = v615
	goto L132
L124:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591))))
	if v599 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v648 = int32(0)
	goto L121
L126:
	;
	goto L127
L127:
	;
	v604 = v591
	goto L128
L128:
	;
	v608 = v604 + int32(1)
	if v608&int32(3) == int32(0) {
		v615 = v608
		goto L123
	} else {
		goto L130
	}
L129:
	;
	v640 = v608
	goto L122
L130:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	if v613 != 0 {
		v604 = v608
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	v628 = int32(-2139062144)
	if (int32(16843008)-v625|v625)&v628 == v628 {
		v619 = v619 + int32(4)
		goto L132
	} else {
		goto L134
	}
L133:
	;
	v634 = v619
	goto L135
L134:
	;
	goto L133
L135:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	if v638 != 0 {
		v634 = v634 + int32(1)
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v640 = v634
	goto L122
L137:
	;
	goto L136
L138:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	if v677 != int32(34) {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v694 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v661)+1)) = uint16(v694)
	if v547&int32(255) == int32(112) {
		goto L146
	} else {
		goto L147
	}
L140:
	;
	goto L139
L141:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v690))) = uint8(v689)
	v661 = v690
	v664 = v664 + int32(1)
	goto L138
L142:
	;
	if v677 == int32(0) {
		goto L140
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v684 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v661)+1)) = uint8(v684)
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	v689 = v686
	v690 = v661 + int32(2)
	goto L141
L145:
	;
	v689 = v677
	v690 = v661 + int32(1)
	goto L141
L146:
	;
	v702 = int32(728204)
	goto L148
L147:
	;
	v702 = int32(714654)
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = v19 + int32(704)
	F_appendStringInfo(m, v19+int32(968), int32(29108), v19+int32(128))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v714 <= int32(0) {
		v846 = v714
		goto L7
	} else {
		goto L150
	}
L150:
	;
	v731 = int32(0)
	v735 = int32(527647)
	goto L151
L151:
	;
	v742 = v731 << (uint(int32(1)) % 32)
	v744 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(172)+v742))))
	v745 = F_attnumTypeId(m, v31, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v747 = v742 + (v24 + int32(236))
	v748 = int32(*(*int16)(unsafe.Add(mBase, uint32(v747))))
	v749 = F_attnumTypeId(m, v28, v748)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v751 = int32(*(*int16)(unsafe.Add(mBase, uint32(v747))))
	v752 = F_attnumAttName(m, v28, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v754 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+560)) = uint8(v754)
	v758 = v752
	v761 = v19 + int32(560)
	goto L156
L156:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758))))
	if v774 != int32(34) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v826))) = uint8(v825)
	v758 = v758 + int32(1)
	v761 = v826
	goto L156
L159:
	;
	if v774 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v820 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v761)+1)) = uint8(v820)
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758))))
	v825 = v822
	v826 = v761 + int32(2)
	goto L158
L162:
	;
	v779 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v761)+1)) = uint16(v779)
	v782 = v731 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v782
	v789 = F_pg_sprintf(m, v19+int32(400), int32(456777), v19+int32(112))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v825 = v774
	v826 = v761 + int32(1)
	goto L158
L165:
	;
	v792 = v731 << (uint(int32(2)) % 32)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(300)+v792)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v735
	F_appendStringInfo(m, v19+int32(968), int32(709145), v19+int32(96))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_generate_operator_clause(m, v19+int32(968), v19+int32(400), v745, v794, v19+int32(560), v749)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(272)+v792))) = v745
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v782 < v816 {
		v731 = v782
		v735 = int32(530859)
		goto L151
	} else {
		goto L168
	}
L168:
	;
	v846 = v816
	goto L7
L169:
	;
	F_errmsg_internal(m, int32(444443), int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(482656), int32(625), int32(318551))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	F_appendStringInfoString(m, v19+int32(968), int32(29129))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L1
	} else {
		goto L265
	}
L173:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+165)))
	if v861&int32(1) == int32(0) {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v867 = v24 + int32(172)
	v873 = int32(*(*int16)(unsafe.Add(mBase, uint32(v867+v846<<(uint(int32(1))%32)-int32(2)))))
	v874 = F_attnumTypeId(m, v31, v873)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v877 = v24 + int32(236)
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v884 = int32(*(*int16)(unsafe.Add(mBase, uint32(v877+v878<<(uint(int32(1))%32)-int32(2)))))
	v885 = F_attnumTypeId(m, v28, v884)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+119)))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v895 = int32(*(*int16)(unsafe.Add(mBase, uint32(v889<<(uint(int32(1))%32)+v877-int32(2)))))
	v896 = F_attnumAttName(m, v28, v895)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v898 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+560)) = uint8(v898)
	v902 = v896
	v905 = v19 + int32(560)
	goto L178
L178:
	;
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	if v918 != int32(34) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	v935 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v905)+1)) = uint16(v935)
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v937
	v944 = F_pg_sprintf(m, v19+int32(400), int32(456777), v19+int32(80))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L186
	}
L180:
	;
	goto L179
L181:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v931))) = uint8(v930)
	v902 = v902 + int32(1)
	v905 = v931
	goto L178
L182:
	;
	if v918 == int32(0) {
		goto L180
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v925 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v905)+1)) = uint8(v925)
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	v930 = v927
	v931 = v905 + int32(2)
	goto L181
L185:
	;
	v930 = v918
	v931 = v905 + int32(1)
	goto L181
L186:
	;
	F_appendStringInfoString(m, v19+int32(968), int32(656670))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	F_initStringInfo(m, v19+int32(256))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_appendStringInfoChar(m, v19+int32(256), int32(40))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v24)+692))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = int32(728204)
	F_appendStringInfo(m, v19+int32(256), int32(709145), v19-int32(-64))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	F_generate_operator_clause(m, v19+int32(256), v19+int32(560), v885, v960, v19+int32(400), v874)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_appendStringInfoChar(m, v19+int32(256), int32(41))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_initStringInfo(m, v19+int32(240))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_appendStringInfoString(m, v19+int32(240), int32(715429))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v998 = int32(*(*int16)(unsafe.Add(mBase, uint32(v992<<(uint(int32(1))%32)+v867-int32(2)))))
	v999 = F_attnumAttName(m, v31, v998)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v1001 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+416)) = uint8(v1001)
	v1005 = v999
	v1008 = v19 + int32(416)
	goto L196
L196:
	;
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005))))
	if v1021 != int32(34) {
		goto L200
	} else {
		goto L201
	}
L197:
	;
	v1038 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v1008)+1)) = uint16(v1038)
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+68))
	v1042 = F_get_namespace_name(m, v1041)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L204
	}
L198:
	;
	goto L197
L199:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1034))) = uint8(v1033)
	v1005 = v1005 + int32(1)
	v1008 = v1034
	goto L196
L200:
	;
	if v1021 == int32(0) {
		goto L198
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v1028 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v1008)+1)) = uint8(v1028)
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005))))
	v1033 = v1030
	v1034 = v1008 + int32(2)
	goto L199
L203:
	;
	v1033 = v1021
	v1034 = v1008 + int32(1)
	goto L199
L204:
	;
	v1044 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+992)) = uint8(v1044)
	v1048 = v1042
	v1051 = v19 + int32(992)
	goto L205
L205:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048))))
	if v1064 != int32(34) {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	v1081 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v1051)+1)) = uint16(v1081)
	v1084 = v19 + int32(992)
	if v1084&int32(3) == int32(0) {
		v1108 = v1084
		goto L215
	} else {
		goto L216
	}
L207:
	;
	goto L206
L208:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1077))) = uint8(v1076)
	v1048 = v1048 + int32(1)
	v1051 = v1077
	goto L205
L209:
	;
	if v1064 == int32(0) {
		goto L207
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v1071 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v1051)+1)) = uint8(v1071)
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048))))
	v1076 = v1073
	v1077 = v1051 + int32(2)
	goto L208
L212:
	;
	v1076 = v1064
	v1077 = v1051 + int32(1)
	goto L208
L213:
	;
	v1144 = v1141 + (v19 + int32(992))
	v1145 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1144))) = uint8(v1145)
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v1149 = v1144 + int32(1)
	v1150 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v1149))) = uint8(v1150)
	v1154 = v1149
	v1157 = v1147 + int32(4)
	goto L230
L214:
	;
	v1141 = v1133 - v1084
	goto L213
L215:
	;
	v1112 = v1108
	goto L224
L216:
	;
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1084))))
	if v1092 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1141 = int32(0)
	goto L213
L218:
	;
	goto L219
L219:
	;
	v1097 = v1084
	goto L220
L220:
	;
	v1101 = v1097 + int32(1)
	if v1101&int32(3) == int32(0) {
		v1108 = v1101
		goto L215
	} else {
		goto L222
	}
L221:
	;
	v1133 = v1101
	goto L214
L222:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101))))
	if v1106 != 0 {
		v1097 = v1101
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1112)))
	v1121 = int32(-2139062144)
	if (int32(16843008)-v1118|v1118)&v1121 == v1121 {
		v1112 = v1112 + int32(4)
		goto L224
	} else {
		goto L226
	}
L225:
	;
	v1127 = v1112
	goto L227
L226:
	;
	goto L225
L227:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127))))
	if v1131 != 0 {
		v1127 = v1127 + int32(1)
		goto L227
	} else {
		goto L229
	}
L228:
	;
	v1133 = v1127
	goto L214
L229:
	;
	goto L228
L230:
	;
	v1170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157))))
	if v1170 != int32(34) {
		goto L234
	} else {
		goto L235
	}
L231:
	;
	v1187 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v1154)+1)) = uint16(v1187)
	if v888&int32(255) == int32(112) {
		goto L238
	} else {
		goto L239
	}
L232:
	;
	goto L231
L233:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1183))) = uint8(v1182)
	v1154 = v1183
	v1157 = v1157 + int32(1)
	goto L230
L234:
	;
	if v1170 == int32(0) {
		goto L232
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1177 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v1154)+1)) = uint8(v1177)
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157))))
	v1182 = v1179
	v1183 = v1154 + int32(2)
	goto L233
L237:
	;
	v1182 = v1170
	v1183 = v1154 + int32(1)
	goto L233
L238:
	;
	v1195 = int32(728204)
	goto L240
L239:
	;
	v1195 = int32(714654)
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v1195
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v19 + int32(992)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v19 + int32(416)
	F_appendStringInfo(m, v19+int32(240), int32(26308), v19+int32(48))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if int32(0) < v1210 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1223 = int32(0)
	v1227 = int32(527647)
	goto L245
L243:
	;
	goto L244
L244:
	;
	F_appendStringInfoString(m, v19+int32(240), int32(653735))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L1
	} else {
		goto L261
	}
L245:
	;
	v1235 = v867 + v1223<<(uint(int32(1))%32)
	v1236 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1235))))
	v1237 = F_attnumTypeId(m, v31, v1236)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L247
	}
L246:
	;
	goto L244
L247:
	;
	v1239 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1235))))
	v1240 = F_attnumAttName(m, v31, v1239)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	v1242 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+560)) = uint8(v1242)
	v1246 = v1240
	v1249 = v19 + int32(560)
	goto L249
L249:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246))))
	if v1262 != int32(34) {
		goto L253
	} else {
		goto L254
	}
L250:
	;
	v1279 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v1249)+1)) = uint16(v1279)
	v1282 = v1223 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1282
	v1289 = F_pg_sprintf(m, v19+int32(400), int32(456777), v19+int32(32))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L1
	} else {
		goto L257
	}
L251:
	;
	goto L250
L252:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1275))) = uint8(v1274)
	v1246 = v1246 + int32(1)
	v1249 = v1275
	goto L249
L253:
	;
	if v1262 == int32(0) {
		goto L251
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v1269 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v1249)+1)) = uint8(v1269)
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246))))
	v1274 = v1271
	v1275 = v1249 + int32(2)
	goto L252
L256:
	;
	v1274 = v1262
	v1275 = v1249 + int32(1)
	goto L252
L257:
	;
	v1292 = v1223 << (uint(int32(2)) % 32)
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(428)+v1292)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1227
	F_appendStringInfo(m, v19+int32(240), int32(709145), v19+int32(16))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	F_generate_operator_clause(m, v19+int32(240), v19+int32(400), v1237, v1294, v19+int32(560), v1237)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(272)+v1292))) = v1237
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	if v1282 < v1316 {
		v1223 = v1282
		v1227 = int32(530859)
		goto L245
	} else {
		goto L260
	}
L260:
	;
	goto L246
L261:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v24)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(728204)
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v19)+240))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
	F_appendStringInfo(m, v19+int32(968), int32(709145), v19)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_generate_operator_clause(m, v19+int32(968), v1343, v885, v1339, v1342, int32(4537))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	F_appendStringInfoString(m, v19+int32(968), int32(648085))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	goto L172
L265:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v19)+968))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v24)+168))
	v1386 = F_ri_PlanCheck(m, v1380, v1381, v19+int32(272), v19+int32(984), v28, v31)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	v1388 = v1386
	goto L6
L267:
	;
	v1413 = F_SPI_finish(m)
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	if v1413 != int32(2) {
		goto L4
	} else {
		goto L269
	}
L269:
	;
	goto L5
L270:
	;
	m.G0 = v19 + int32(1264)
	return
L271:
	;
	F_errmsg_internal(m, int32(444443), int32(0))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(482656), int32(901), int32(106444))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
