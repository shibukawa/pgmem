package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CreateStandaloneExprContext(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	v4 = F_palloc0(m, int32(72))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(382)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStandaloneExprContext[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v13
		v19 = F_AllocSetContextCreateInternal(m, v13, int32(_a_F_CreateStandaloneExprContext_0), int32(0), int32(_a_F_CreateStandaloneExprContext_1), int32(_a_F_CreateStandaloneExprContext_2))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v4)+32)) = v21
			v26 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+40)) = v26
			*(*int64)(unsafe.Add(mBase, uint32(v4)+64)) = v21
			v30 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+52)) = uint8(v30)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = v26
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+44)) = uint8(v30)
			return v4
		}
	}
}
func F_CreateTemplateTupleDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_palloc(m, l0*int32(116)+int32(20))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = int64(4294967295)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(-4294965047)
		return v7
	}
}
func F_create_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 float64
	_ = v109
	var v111 int32
	_ = v111
	var v114 float64
	_ = v114
	var v116 int32
	_ = v116
	var v122 float64
	_ = v122
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v139 float64
	_ = v139
	var v143 float64
	_ = v143
	var v151 float64
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v197 float64
	_ = v197
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
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
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v640 float64
	_ = v640
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 float64
	_ = v663
	var v665 float64
	_ = v665
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v788 float64
	_ = v788
	var v807 int32
	_ = v807
	var v824 float64
	_ = v824
	var v846 int32
	_ = v846
	var v888 float64
	_ = v888
	var v889 int32
	_ = v889
	var v898 int32
	_ = v898
	var v904 float64
	_ = v904
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v990 int32
	_ = v990
	var v991 float64
	_ = v991
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1015 float64
	_ = v1015
	var v1016 float64
	_ = v1016
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 float64
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 float64
	_ = v1039
	var v1041 float64
	_ = v1041
	var v1042 float64
	_ = v1042
	var v1046 float64
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1062 float64
	_ = v1062
	var v1087 float64
	_ = v1087
	var v1088 float64
	_ = v1088
	var v1096 float64
	_ = v1096
	var v1100 float64
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 float64
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1112 float64
	_ = v1112
	var v1113 float64
	_ = v1113
	var v1116 float64
	_ = v1116
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 float64
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int64
	_ = v1129
	var v1134 float64
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1145 int32
	_ = v1145
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 float64
	_ = v1191
	var v1192 float64
	_ = v1192
	var v1208 float64
	_ = v1208
	var v1229 float64
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int64
	_ = v1231
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 float64
	_ = v1293
	var v1294 float64
	_ = v1294
	var v1319 float64
	_ = v1319
	var v1331 float64
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1340 float64
	_ = v1340
	var v1341 float64
	_ = v1341
	var v1343 float64
	_ = v1343
	var v1346 float64
	_ = v1346
	var v1349 float64
	_ = v1349
	var v1353 float64
	_ = v1353
	var v1362 float64
	_ = v1362
	var v1366 float64
	_ = v1366
	var v1367 float64
	_ = v1367
	var v1373 float64
	_ = v1373
	var v1381 float64
	_ = v1381
	var v1385 float64
	_ = v1385
	var v1388 float64
	_ = v1388
	var v1394 float64
	_ = v1394
	var v1395 float64
	_ = v1395
	var v1396 float64
	_ = v1396
	var v1404 float64
	_ = v1404
	var v1408 float64
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 float64
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 float64
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int64
	_ = v1419
	var v1439 int32
	_ = v1439
	var v1440 float64
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1447 int32
	_ = v1447
	var v1455 float64
	_ = v1455
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1489 float64
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 float64
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1507 float64
	_ = v1507
	var v1534 float64
	_ = v1534
	var v1536 float64
	_ = v1536
	var v1544 float64
	_ = v1544
	var v1548 float64
	_ = v1548
	var v1562 float64
	_ = v1562
	var v1586 float64
	_ = v1586
	var v1588 float64
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1590 float64
	_ = v1590
	var v1602 float64
	_ = v1602
	var v1606 float64
	_ = v1606
	var v1607 float64
	_ = v1607
	var v1609 float64
	_ = v1609
	v12 = float64(0)
	v37 = m.G0
	v39 = v37 - int32(16)
	m.G0 = v39
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = l8
	v43 = F_palloc0(m, int32(112))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v43))) = int64(1541893259564)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v55 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v52, l9, v39+int32(12))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v55
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	v59 = l7 & v58
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+20)) = uint8(v59)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v62 != int32(1) {
		v70 = int32(0)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v72 = v70 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+21)) = uint8(v72)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v43)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+24)) = v74
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v43)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+76)) = uint8(v79)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+96)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v43)+88)) = v83
	v86 = m.G0
	v88 = v86 + int32(-64)
	m.G0 = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v92 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l3)+88))
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v99 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+40)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if v102 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v66 != int32(1) {
		v70 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v70 = v69
	goto L4
L7:
	;
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v108)))
	*(*float64)(unsafe.Add(mBase, uint32(v43)+32)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	if int32(0) < v111 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v108 = v102 + int32(8)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v108 = v105 + int32(16)
	goto L7
L11:
	;
	v114 = base.F64_convert_i32_u(v111)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_hashjoin_path[0])))
	if v116 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v43)+104)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v43)+100)) = v90
	v151 = base.F64_mul(base.F64_convert_i32_s(v91), base.F64_convert_i32_s(v90))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v152 == int32(295) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v122 = base.F64_add(base.F64_mul(v114, float64(-0.3)), float64(1))
	if base.F64_gt(v122, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v128 = v114
	goto L16
L16:
	;
	v130 = float64(1e+100)
	v131 = base.F64_div(v109, v128)
	if base.F64_gt(v131, v130) != 0 {
		v143 = v130
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v126 = v122
	goto L19
L18:
	;
	v126 = math.Float64frombits(uint64(0x8000000000000000))
	goto L19
L19:
	;
	v128 = base.F64_add(v126, v114)
	goto L16
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v43)+32)) = v143
	goto L13
L21:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v131)&int64(9223372036854775807)) {
		v143 = v130
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v139 = float64(1)
	if base.F64_le(v131, v139) != 0 {
		v143 = v139
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v143 = base.F64_nearest(v131)
	goto L20
L24:
	;
	v1087 = float64(1e+100)
	v1088 = base.F64_mul(v97, v1062)
	if base.F64_gt(v1088, v1087) != 0 {
		v1100 = v1087
		goto L194
	} else {
		goto L195
	}
L25:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v88))) = base.F64_div(float64(1), v151)
	v1062 = float64(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v88))) = int64(4607182418800017408)
	v162 = m.G0
	v164 = v162 - int32(16)
	m.G0 = v164
	v166 = F_list_copy(m, l10)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if l10 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	m.G0 = v164 + int32(16)
	if v846 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L30:
	;
	v846 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v171 < int32(2) {
		v846 = l10
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v166 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v88))) = base.F64_div(float64(1), v824)
	v846 = v807
	goto L29
L35:
	;
	v807 = int32(0)
	v824 = float64(1)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v180 = int32(0)
	v197 = float64(1)
	v211 = v166
	goto L38
L38:
	;
	v216 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v164)+12)) = v216
	v218 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = v218
	v224 = v180
	v225 = v216
	v226 = v218
	v230 = v218
	v255 = v211
	v256 = v218
	goto L40
L39:
	;
	v807 = v771
	v824 = v788
	goto L34
L40:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v230 < v260 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	if v612 != 0 {
		goto L105
	} else {
		goto L106
	}
L42:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262+v230<<(uint(int32(2))%32))))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+120)))
	if v269 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v576 = v224
	v578 = v226
	v607 = v255
	v608 = v256
	goto L44
L44:
	;
	goto L41
L45:
	;
	v270 = int32(48)
	goto L47
L46:
	;
	v270 = int32(44)
	goto L47
L47:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266+v270)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	if v269 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v291 = int32(0)
	if v272 == v291 {
		goto L62
	} else {
		goto L63
	}
L49:
	;
	v275 = int32(0)
	if v274 == v275 {
		v288 = v275
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v274 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v278 < int32(2) {
		v288 = v275
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v288 = v282
	goto L48
L54:
	;
	v288 = int32(0)
	goto L48
L55:
	;
	goto L56
L56:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v288 = v287
	goto L48
L57:
	;
	if v569 != 0 {
		v224 = v538
		v225 = v539
		v226 = v540
		v230 = v545 + int32(1)
		v255 = v569
		v256 = v570
		goto L40
	} else {
		goto L102
	}
L58:
	;
	v536 = F_list_delete_nth_cell(m, v255, v230)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L101
	}
L59:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v387 = F_remove_nulling_relids(m, v288, v385, int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L87
	}
L60:
	;
	v380 = F_lappend(m, v224, v266)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L86
	}
L61:
	;
	if v344 == int32(0) {
		goto L60
	} else {
		goto L77
	}
L62:
	;
	v344 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v299 = int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v300 <= v299 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v303 = v299
	goto L67
L66:
	;
	v303 = v300
	goto L67
L67:
	;
	v308 = int32(0)
	v311 = int32(-1)
	goto L69
L68:
	;
	v344 = v336
	goto L61
L69:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(8)+v308<<(uint(int32(2))%32))))
	if v318 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164+int32(12)))) = v328
	v336 = int32(1)
	goto L68
L71:
	;
	if int32(0) <= v311 {
		v336 = v291
		goto L68
	} else {
		goto L74
	}
L72:
	;
	v328 = v311
	goto L73
L73:
	;
	v330 = v308 + int32(1)
	if v330 != v303 {
		v308 = v330
		v311 = v328
		goto L69
	} else {
		goto L76
	}
L74:
	;
	if base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v318)) {
		v336 = v291
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v328 = base.I32_ctz(v318) | v308<<(uint(int32(5))%32)
	goto L73
L76:
	;
	goto L70
L77:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v349 = v347 << (uint(int32(2)) % 32)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v349+v350)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+112))
	if v353 == int32(0) {
		goto L60
	} else {
		goto L78
	}
L78:
	;
	if v225 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v358+v349)))
	if v360 == int32(0) {
		goto L60
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if base.B2i32(v225 != v347) == int32(0) {
		v382 = v226
		v383 = v225
		goto L59
	} else {
		goto L85
	}
L82:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+21)))
	v365 = v363 - int32(102)
	if base.Ui32(int32(12)) < base.Ui32(v365) {
		goto L60
	} else {
		goto L83
	}
L83:
	;
	if int32(1)<<(uint(v365)%32)&int32(_a_F_create_hashjoin_path_0) == int32(0) {
		goto L60
	} else {
		goto L84
	}
L84:
	;
	v382 = v352
	v383 = v347
	goto L59
L85:
	;
	v538 = v224
	v539 = v225
	v540 = v226
	v545 = v230
	v569 = v255
	v570 = v256
	goto L57
L86:
	;
	v498 = v380
	v499 = v225
	v500 = v226
	v530 = v256
	goto L58
L87:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	if v389 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v483 = F_palloc0(m, int32(24))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L98
	}
L89:
	;
	v392 = int32(0)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v393 <= v392 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v423 = v392
	goto L91
L91:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v432+v423<<(uint(int32(2))%32))))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	v438 = F_equal(m, v387, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	v538 = v224
	v539 = v383
	v540 = v382
	v545 = v230
	v569 = v255
	v570 = v256
	goto L57
L93:
	;
	if v438 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v443 = v423 + int32(1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v443 < v444 {
		v423 = v443
		goto L91
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	goto L92
L97:
	;
	goto L88
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = v387
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v486+v487<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+4)) = v491
	v493 = F_lappend(m, v389, v483)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = v493
	v496 = F_lappend(m, v256, v266)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v498 = v224
	v499 = v383
	v500 = v382
	v530 = v496
	goto L58
L101:
	;
	v538 = v498
	v539 = v499
	v540 = v500
	v545 = v230 - int32(1)
	v569 = v536
	v570 = v530
	goto L57
L102:
	;
	v576 = v538
	v578 = v540
	v607 = v569
	v608 = v570
	goto L44
L103:
	;
	if v607 != 0 {
		v180 = v771
		v197 = v788
		v211 = v607
		goto L38
	} else {
		goto L148
	}
L104:
	;
	v640 = v197
	goto L112
L105:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)+4))
	if int32(2) <= v613 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v616 = F_list_concat(m, v576, v608)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	F_list_free_deep(m, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_list_free(m, v608)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v771 = v616
	v788 = v197
	goto L103
L112:
	;
	v661 = F_estimate_multivariate_ndistinct(m, l0, v578, v164+int32(8), v164)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L114
	}
L113:
	;
	v672 = v576
	v673 = int32(0)
	goto L119
L114:
	;
	v663 = *(*float64)(unsafe.Add(mBase, uint32(v164)))
	if base.F64_lt(v640, v663) != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v665 = v663
	goto L117
L116:
	;
	v665 = v640
	goto L117
L117:
	;
	if v661 != 0 {
		v640 = v665
		goto L112
	} else {
		goto L118
	}
L118:
	;
	goto L113
L119:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v612)+4))
	if v673 < v704 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v771 = v672
	v788 = v640
	goto L103
L121:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v612)+12))
	v710 = v706 + v673<<(uint(int32(2))%32)
	goto L123
L122:
	;
	v710 = int32(0)
	goto L123
L123:
	;
	if v608 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v771 = v576
	v788 = v640
	goto L103
L125:
	;
	goto L126
L126:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	if v713 <= v673 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	goto L120
L128:
	;
	if v710 == int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v608)+12))
	v720 = v717 + v673<<(uint(int32(2))%32)
	if v720 == int32(0) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	v725 = int32(0)
	if v723 == v725 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v763 != 0 {
		goto L144
	} else {
		goto L145
	}
L132:
	;
	v763 = int32(0)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v723)+4))
	if v731 <= int32(0) {
		v756 = v725
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v763 = v756
	goto L131
L136:
	;
	v734 = int32(0)
	if v734 < v731 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v737 = v731
	goto L139
L138:
	;
	v737 = v734
	goto L139
L139:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v723)+12))
	v740 = int32(0)
	goto L140
L140:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v738+v740<<(uint(int32(2))%32))))
	v749 = base.B2i32(v748 == v724)
	if v748 == v724 {
		v756 = v749
		goto L135
	} else {
		goto L142
	}
L141:
	;
	v756 = v749
	goto L135
L142:
	;
	v751 = v740 + int32(1)
	if v751 != v737 {
		v740 = v751
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	v765 = F_lappend(m, v672, v764)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	v767 = v672
	goto L146
L146:
	;
	v672 = v767
	v673 = v673 + int32(1)
	goto L119
L147:
	;
	v767 = v765
	goto L146
L148:
	;
	goto L39
L149:
	;
	v1062 = float64(1)
	goto L24
L150:
	;
	goto L151
L151:
	;
	v888 = float64(1)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	if v889 <= int32(0) {
		v1062 = v888
		goto L24
	} else {
		goto L152
	}
L152:
	;
	v898 = int32(0)
	v904 = v888
	goto L153
L153:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v846)+12))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v929+v898<<(uint(int32(2))%32))))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v933)+48))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v935)+8))
	v937 = int32(0)
	if v934 == v937 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v1062 = v1046
	goto L24
L155:
	;
	v1041 = *(*float64)(unsafe.Add(mBase, uint32(v1036+v933)))
	v1042 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	if base.F64_lt(v1039, v1042) != 0 {
		goto L187
	} else {
		goto L188
	}
L156:
	;
	if v990 != 0 {
		goto L170
	} else {
		goto L171
	}
L157:
	;
	v990 = int32(1)
	goto L156
L158:
	;
	goto L159
L159:
	;
	if v936 == int32(0) {
		v981 = v937
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v990 = v981
	goto L156
L161:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v934)+4))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	if v947 < v946 {
		v981 = v937
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v949 = int32(1)
	if v946 <= v949 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v952 = v949
	goto L165
L164:
	;
	v952 = v946
	goto L165
L165:
	;
	v953 = int32(8)
	v958 = int32(0)
	goto L166
L166:
	;
	v965 = v958 << (uint(int32(2)) % 32)
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v934+v953+v965)))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v965+(v936+v953))))
	v972 = v967 & (v969 ^ int32(-1))
	v974 = base.B2i32(v972 == int32(0))
	if v972 != 0 {
		v981 = v974
		goto L160
	} else {
		goto L168
	}
L167:
	;
	v981 = v974
	goto L160
L168:
	;
	v976 = v958 + int32(1)
	if v976 != v952 {
		v958 = v976
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v991 = *(*float64)(unsafe.Add(mBase, uint32(v933)+136))
	if base.F64_lt(v991, float64(0)) == int32(0) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	v1016 = *(*float64)(unsafe.Add(mBase, uint32(v933)+128))
	if base.F64_lt(v1016, float64(0)) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L173:
	;
	v1036 = int32(152)
	v1039 = v991
	goto L155
L174:
	;
	goto L175
L175:
	;
	v999 = int32(0)
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+28))
	if v1001 == v999 {
		v1009 = v999
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v1010 = int32(152)
	F_estimate_hash_bucket_stats(m, l0, v1009, v151, v933+v1010, v933+int32(136))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	if v1004 < int32(2) {
		v1009 = v999
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+12))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+4))
	v1009 = v1008
	goto L176
L179:
	;
	v1015 = *(*float64)(unsafe.Add(mBase, uint32(v933)+136))
	v1036 = v1010
	v1039 = v1015
	goto L155
L180:
	;
	v1036 = int32(144)
	v1039 = v1016
	goto L155
L181:
	;
	goto L182
L182:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+28))
	if v1025 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+12))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1026)))
	v1029 = v1027
	goto L185
L184:
	;
	v1029 = int32(0)
	goto L185
L185:
	;
	v1030 = int32(144)
	F_estimate_hash_bucket_stats(m, l0, v1029, v151, v933+v1030, v933+int32(128))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v1035 = *(*float64)(unsafe.Add(mBase, uint32(v933)+128))
	v1036 = v1030
	v1039 = v1035
	goto L155
L187:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v88))) = v1039
	goto L189
L188:
	;
	goto L189
L189:
	;
	if base.F64_gt(v904, v1041) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1046 = v1041
	goto L192
L191:
	;
	v1046 = v904
	goto L192
L192:
	;
	v1048 = v898 + int32(1)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	if v1048 < v1049 {
		v898 = v1048
		v904 = v1046
		goto L153
	} else {
		goto L193
	}
L193:
	;
	goto L154
L194:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+32))
	v1106 = *(*float64)(unsafe.Add(mBase, _c_F_create_hashjoin_path[1]))
	v1108 = *(*int32)(unsafe.Add(mBase, _c_F_create_hashjoin_path[2]))
	v1112 = base.F64_mul(base.F64_mul(v1106, base.F64_convert_i32_s(v1108)), float64(1024))
	v1113 = float64(4.294967295e+09)
	if base.F64_lt(v1112, v1113) != 0 {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1088)&int64(9223372036854775807)) {
		v1100 = v1087
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v1096 = float64(1)
	if base.F64_le(v1088, v1096) != 0 {
		v1100 = v1096
		goto L194
	} else {
		goto L197
	}
L197:
	;
	v1100 = base.F64_nearest(v1088)
	goto L194
L198:
	;
	v1126 = *(*float64)(unsafe.Add(mBase, _c_F_create_hashjoin_path[3]))
	v1128 = v88 + int32(24)
	v1129 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1128))) = v1129
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v1129
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l0
	v1134 = float64(0)
	if l10 == int32(0) {
		v1208 = v1134
		v1229 = v1134
		goto L205
	} else {
		goto L206
	}
L199:
	;
	v1116 = v1112
	goto L201
L200:
	;
	v1116 = v1113
	goto L201
L201:
	;
	if base.F64_lt(v1116, float64(4.294967296e+09))&base.F64_ge(v1116, float64(0)) != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1122 = base.I32_trunc_f64_u(v1116)
	v1124 = v1122
	goto L198
L203:
	;
	goto L204
L204:
	;
	v1124 = int32(0)
	goto L198
L205:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v43)+88))
	v1231 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1128))) = v1231
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v1231
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l0
	if v1230 == int32(0) {
		v1319 = v12
		v1331 = float64(0)
		goto L212
	} else {
		goto L213
	}
L206:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1139 <= int32(0) {
		v1208 = v1134
		v1229 = float64(0)
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1145 = int32(0)
	goto L208
L208:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1178+v1145<<(uint(int32(2))%32))))
	v1185 = F_cost_qual_eval_walker(m, v1182, v88+int32(8))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L1
	} else {
		goto L210
	}
L209:
	;
	v1191 = *(*float64)(unsafe.Add(mBase, uint32(v88)+16))
	v1192 = *(*float64)(unsafe.Add(mBase, uint32(v88)+24))
	v1208 = v1191
	v1229 = v1192
	goto L205
L210:
	;
	v1188 = v1145 + int32(1)
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1188 < v1189 {
		v1145 = v1188
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if v1332&int32(-2) != int32(4) {
		goto L221
	} else {
		goto L222
	}
L213:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+4))
	if v1240 <= int32(0) {
		v1319 = v12
		v1331 = float64(0)
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v1247 = int32(0)
	goto L215
L215:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+12))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1280+v1247<<(uint(int32(2))%32))))
	v1287 = F_cost_qual_eval_walker(m, v1284, v88+int32(8))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L1
	} else {
		goto L217
	}
L216:
	;
	v1293 = *(*float64)(unsafe.Add(mBase, uint32(v88)+24))
	v1294 = *(*float64)(unsafe.Add(mBase, uint32(v88)+16))
	v1319 = v1293
	v1331 = v1294
	goto L212
L217:
	;
	v1290 = v1247 + int32(1)
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+4))
	if v1290 < v1291 {
		v1247 = v1290
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v1588 = *(*float64)(unsafe.Add(mBase, _c_F_create_hashjoin_path[4]))
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v1590 = *(*float64)(unsafe.Add(mBase, uint32(v1589)+24))
	if base.F64_lt(base.F64_convert_i32_u(v1124), base.F64_mul(v1100, base.F64_convert_i32_u((v1102+int32(7))&int32(-8)+int32(24)))) != 0 {
		goto L254
	} else {
		goto L255
	}
L220:
	;
	v1394 = float64(1e+100)
	v1395 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	v1396 = base.F64_mul(v97, v1395)
	if base.F64_gt(v1396, v1394) != 0 {
		v1408 = v1394
		goto L236
	} else {
		goto L237
	}
L221:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v1337 != int32(1) {
		goto L220
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v1340 = float64(1e+100)
	v1341 = *(*float64)(unsafe.Add(mBase, uint32(l4)+16))
	v1343 = base.F64_nearest(base.F64_mul(v99, v1341))
	v1346 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	v1349 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	v1353 = base.F64_mul(base.F64_mul(v97, v1346), base.F64_div(float64(2), base.F64_add(v1349, float64(1))))
	if base.F64_gt(v1353, v1340) != 0 {
		v1366 = v1340
		goto L225
	} else {
		goto L226
	}
L224:
	;
	goto L223
L225:
	;
	v1367 = base.F64_sub(v99, v1343)
	v1373 = base.F64_div(v97, v151)
	if base.F64_gt(v1373, float64(1e+100)) != 0 {
		v1385 = v1340
		goto L229
	} else {
		goto L230
	}
L226:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1353)&int64(9223372036854775807)) {
		v1366 = float64(1e+100)
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v1362 = float64(1)
	if base.F64_le(v1353, v1362) != 0 {
		v1366 = v1362
		goto L225
	} else {
		goto L228
	}
L228:
	;
	v1366 = base.F64_nearest(v1353)
	goto L225
L229:
	;
	if v1332 == int32(5) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1373)&int64(9223372036854775807)) {
		v1385 = v1340
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v1381 = float64(1)
	if base.F64_le(v1373, v1381) != 0 {
		v1385 = v1381
		goto L229
	} else {
		goto L232
	}
L232:
	;
	v1385 = base.F64_nearest(v1373)
	goto L229
L233:
	;
	v1388 = v1367
	goto L235
L234:
	;
	v1388 = v1343
	goto L235
L235:
	;
	v1562 = v1388
	v1586 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1229, v1367), v1385), float64(0.05)), base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1229, v1343), v1366), float64(0.5)), v92))
	goto L219
L236:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v43)+84))
	v1410 = *(*float64)(unsafe.Add(mBase, uint32(v1409)+32))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	v1412 = *(*float64)(unsafe.Add(mBase, uint32(v1411)+32))
	v1414 = v88 + int32(8)
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1411)+8))
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1415)+8))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+8))
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+8))
	v1419 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1414)+48)) = v1419
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+16)) = v1418
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+12)) = v1416
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+8)) = v1418
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+4)) = v1416
	*(*int32)(unsafe.Add(mBase, uint32(v1414))) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(v1414)+20)) = v1419
	*(*int64)(unsafe.Add(mBase, uint32(v1414)+28)) = v1419
	*(*int64)(unsafe.Add(mBase, uint32(v1414)+36)) = v1419
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+43)) = int32(0)
	goto L240
L237:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1396)&int64(9223372036854775807)) {
		v1408 = v1394
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1404 = float64(1)
	if base.F64_le(v1396, v1404) != 0 {
		v1408 = v1404
		goto L236
	} else {
		goto L239
	}
L239:
	;
	v1408 = base.F64_nearest(v1396)
	goto L236
L240:
	;
	if l10 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1534 = float64(1e+100)
	v1536 = base.F64_mul(v1410, base.F64_mul(v1412, v1507))
	if base.F64_gt(v1536, v1534) != 0 {
		v1548 = v1534
		goto L250
	} else {
		goto L251
	}
L242:
	;
	v1507 = float64(1)
	goto L241
L243:
	;
	goto L244
L244:
	;
	v1439 = int32(0)
	v1440 = float64(1)
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1441 <= v1439 {
		v1507 = v1440
		goto L241
	} else {
		goto L245
	}
L245:
	;
	v1447 = v1439
	v1455 = v1440
	goto L246
L246:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1480+v1447<<(uint(int32(2))%32))))
	v1485 = int32(0)
	v1489 = F_clause_selectivity(m, l0, v1484, v1485, v1485, v88+int32(8))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L1
	} else {
		goto L248
	}
L247:
	;
	v1507 = v1491
	goto L241
L248:
	;
	v1491 = base.F64_mul(v1455, v1489)
	v1493 = v1447 + int32(1)
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1493 < v1494 {
		v1447 = v1493
		v1455 = v1491
		goto L246
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	v1562 = v1548
	v1586 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v99, v1229), v1408), float64(0.5)), v92)
	goto L219
L251:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1536)&int64(9223372036854775807)) {
		v1548 = v1534
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v1544 = float64(1)
	if base.F64_le(v1536, v1544) != 0 {
		v1548 = v1544
		goto L250
	} else {
		goto L253
	}
L253:
	;
	v1548 = base.F64_nearest(v1536)
	goto L250
L254:
	;
	v1602 = base.F64_add(v93, v1126)
	goto L256
L255:
	;
	v1602 = v93
	goto L256
L256:
	;
	v1606 = *(*float64)(unsafe.Add(mBase, uint32(v1589)+16))
	v1607 = base.F64_add(base.F64_add(base.F64_add(v1602, v1208), base.F64_sub(v1331, v1208)), v1606)
	*(*float64)(unsafe.Add(mBase, uint32(v43)+48)) = v1607
	v1609 = *(*float64)(unsafe.Add(mBase, uint32(v43)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v43)+56)) = base.F64_add(v1607, base.F64_add(base.F64_mul(v1590, v1609), base.F64_add(base.F64_mul(base.F64_add(v1588, base.F64_sub(v1319, v1229)), v1562), v1586)))
	m.G0 = v88 - int32(-64)
	m.G0 = v39 + int32(16)
	return v43
}
func F_create_mergejoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 float64
	_ = v15
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v108 int32
	_ = v108
	var v111 float64
	_ = v111
	var v113 int32
	_ = v113
	var v119 float64
	_ = v119
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v136 float64
	_ = v136
	var v140 float64
	_ = v140
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v151 float64
	_ = v151
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 float64
	_ = v205
	var v206 float64
	_ = v206
	var v223 float64
	_ = v223
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 float64
	_ = v299
	var v300 float64
	_ = v300
	var v324 float64
	_ = v324
	var v333 float64
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 float64
	_ = v360
	var v361 int32
	_ = v361
	var v362 float64
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
	var v369 int64
	_ = v369
	var v388 int32
	_ = v388
	var v389 float64
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v407 float64
	_ = v407
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 float64
	_ = v434
	var v435 int32
	_ = v435
	var v436 float64
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v455 float64
	_ = v455
	var v475 float64
	_ = v475
	var v477 float64
	_ = v477
	var v485 float64
	_ = v485
	var v489 float64
	_ = v489
	var v491 float64
	_ = v491
	var v492 int32
	_ = v492
	var v493 float64
	_ = v493
	var v494 int32
	_ = v494
	var v499 float64
	_ = v499
	var v505 float64
	_ = v505
	var v508 float64
	_ = v508
	var v509 float64
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v518 float64
	_ = v518
	var v521 float64
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v648 int32
	_ = v648
	var v663 float64
	_ = v663
	var v678 int32
	_ = v678
	var v679 float64
	_ = v679
	var v681 float64
	_ = v681
	var v689 float64
	_ = v689
	var v690 float64
	_ = v690
	var v692 float64
	_ = v692
	v15 = float64(0)
	v31 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(16)
	m.G0 = v35
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = l7
	v39 = F_palloc0(m, int32(120))
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
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(1537598292267)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v51 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v48, l9, v35+int32(12))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+20)) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v51
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v56 != int32(1) {
		v63 = v31
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v65 = v63 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+21)) = uint8(v65)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = v67
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+76)) = uint8(v71)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+108)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v39)+104)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v39)+100)) = l11
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v39)+88)) = v75
	v81 = m.G0
	v83 = v81 + int32(-64)
	m.G0 = v83
	v85 = *(*float64)(unsafe.Add(mBase, uint32(l3)+72))
	v86 = *(*float64)(unsafe.Add(mBase, uint32(l3)+64))
	v87 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	v88 = *(*float64)(unsafe.Add(mBase, uint32(l3)+48))
	v89 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v96 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v59 != int32(1) {
		v63 = v31
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v63 = v62
	goto L4
L7:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v102)))
	*(*float64)(unsafe.Add(mBase, uint32(v39)+32)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if int32(0) < v108 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v102 = v96 + int32(8)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v102 = v99 + int32(16)
	goto L7
L11:
	;
	v111 = base.F64_convert_i32_u(v108)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_mergejoin_path[0])))
	if v113 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v145 = v81 + int32(-40)
	v146 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = l0
	v151 = float64(0)
	if v104 == int32(0) {
		v223 = v151
		v239 = v151
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v119 = base.F64_add(base.F64_mul(v111, float64(-0.3)), float64(1))
	if base.F64_gt(v119, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v125 = v111
	goto L16
L16:
	;
	v127 = float64(1e+100)
	v128 = base.F64_div(v106, v125)
	if base.F64_gt(v128, v127) != 0 {
		v140 = v127
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v123 = v119
	goto L19
L18:
	;
	v123 = math.Float64frombits(uint64(0x8000000000000000))
	goto L19
L19:
	;
	v125 = base.F64_add(v123, v111)
	goto L16
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v39)+32)) = v140
	goto L13
L21:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v128)&int64(9223372036854775807)) {
		v140 = v127
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v136 = float64(1)
	if base.F64_le(v128, v136) != 0 {
		v140 = v136
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v140 = base.F64_nearest(v128)
	goto L20
L24:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	v241 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v241
	*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = l0
	if v240 == int32(0) {
		v324 = v15
		v333 = float64(0)
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v156 <= int32(0) {
		v223 = v151
		v239 = float64(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v162 = int32(0)
	goto L27
L27:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v162<<(uint(int32(2))%32))))
	v199 = F_cost_qual_eval_walker(m, v196, v81+int32(-56))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v205 = *(*float64)(unsafe.Add(mBase, uint32(v83)+24))
	v206 = *(*float64)(unsafe.Add(mBase, uint32(v83)+16))
	v223 = v205
	v239 = v206
	goto L24
L29:
	;
	v202 = v162 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v202 < v203 {
		v162 = v202
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	if v334&int32(-2) != int32(4) {
		goto L40
	} else {
		goto L41
	}
L32:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v250 <= int32(0) {
		v324 = v15
		v333 = float64(0)
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v256 = int32(0)
	goto L34
L34:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v286+v256<<(uint(int32(2))%32))))
	v293 = F_cost_qual_eval_walker(m, v290, v81+int32(-56))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v299 = *(*float64)(unsafe.Add(mBase, uint32(v83)+24))
	v300 = *(*float64)(unsafe.Add(mBase, uint32(v83)+16))
	v324 = v299
	v333 = v300
	goto L31
L36:
	;
	v296 = v256 + int32(1)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v296 < v297 {
		v256 = v296
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)) = uint8(v357)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	v360 = *(*float64)(unsafe.Add(mBase, uint32(v359)+32))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v362 = *(*float64)(unsafe.Add(mBase, uint32(v361)+32))
	v364 = v81 + int32(-56)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	v369 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v364)+48)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v364)+16)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v364)+12)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v364)+8)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v364)+4)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(v364)+20)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v364)+28)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v364)+36)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v364)+43)) = int32(0)
	goto L51
L39:
	;
	v357 = int32(0)
	goto L38
L40:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v339 != int32(1) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	if v343 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v345 = v344
	goto L46
L45:
	;
	v345 = int32(0)
	goto L46
L46:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	if v347 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v350 = v348
	goto L49
L48:
	;
	v350 = int32(0)
	goto L49
L49:
	;
	if v350 == v345 {
		v357 = int32(1)
		goto L38
	} else {
		goto L50
	}
L50:
	;
	goto L39
L51:
	;
	if v104 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v475 = float64(1e+100)
	v477 = base.F64_mul(v360, base.F64_mul(v362, v455))
	if base.F64_gt(v477, v475) != 0 {
		v489 = v475
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v455 = float64(1)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v388 = int32(0)
	v389 = float64(1)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v390 <= v388 {
		v455 = v389
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v395 = v388
	v407 = v389
	goto L57
L57:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v425+v395<<(uint(int32(2))%32))))
	v430 = int32(0)
	v434 = F_clause_selectivity(m, l0, v429, v430, v430, v81+int32(-56))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v455 = v436
	goto L52
L59:
	;
	v436 = base.F64_mul(v407, v434)
	v438 = v395 + int32(1)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v438 < v439 {
		v395 = v438
		v407 = v436
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if base.F64_le(v93, float64(0)) != 0 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v477)&int64(9223372036854775807)) {
		v489 = v475
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v485 = float64(1)
	if base.F64_le(v477, v485) != 0 {
		v489 = v485
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v489 = base.F64_nearest(v477)
	goto L61
L65:
	;
	v491 = float64(1)
	goto L67
L66:
	;
	v491 = v93
	goto L67
L67:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)))
	v493 = float64(0)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v494 == int32(295) {
		v505 = v493
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v508 = base.F64_add(base.F64_div(v505, v87), float64(1))
	v509 = base.F64_mul(v89, v508)
	v510 = int32(1)
	if v492&v510 != 0 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	if v492&int32(1) != 0 {
		v505 = v493
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v499 = base.F64_sub(v489, v491)
	if base.F64_lt(v499, float64(0)) == int32(0) {
		v505 = v499
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v505 = float64(0)
	goto L68
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+113)) = uint8(v648)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v679 = *(*float64)(unsafe.Add(mBase, uint32(v678)+24))
	v681 = *(*float64)(unsafe.Add(mBase, _c_F_create_mergejoin_path[1]))
	v689 = *(*float64)(unsafe.Add(mBase, uint32(v678)+16))
	v690 = base.F64_add(base.F64_add(base.F64_sub(v333, v239), base.F64_add(base.F64_mul(v223, base.F64_add(base.F64_mul(v85, v508), v86)), base.F64_add(v91, v239))), v689)
	*(*float64)(unsafe.Add(mBase, uint32(v39)+48)) = v690
	v692 = *(*float64)(unsafe.Add(mBase, uint32(v39)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v39)+56)) = base.F64_add(v690, base.F64_add(base.F64_mul(v679, v692), base.F64_add(base.F64_mul(base.F64_add(v681, base.F64_sub(v324, v223)), v489), base.F64_add(base.F64_mul(v223, base.F64_add(base.F64_mul(base.F64_sub(v87, v85), v508), base.F64_sub(v88, v86))), base.F64_add(v90, v663)))))
	m.G0 = v83 - int32(-64)
	m.G0 = v35 + int32(16)
	return v39
L73:
	;
	v648 = int32(0)
	v663 = v509
	goto L72
L74:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_mergejoin_path[2])))
	v518 = *(*float64)(unsafe.Add(mBase, _c_F_create_mergejoin_path[3]))
	v521 = base.F64_add(base.F64_mul(base.F64_mul(v87, v518), v508), v89)
	if base.B2i32(v514 == int32(1))&base.F64_gt(v509, v521) != 0 {
		v648 = v510
		v663 = v521
		goto L72
	} else {
		goto L75
	}
L75:
	;
	if v103 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v526 = int32(0)
	v527 = l6
	goto L80
L77:
	;
	goto L78
L78:
	;
	if v514 == int32(0) {
		goto L73
	} else {
		goto L94
	}
L79:
	;
	if v591&int32(1) != 0 {
		goto L73
	} else {
		goto L93
	}
L80:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	switch v559 - int32(331) {
	case 0:
		goto L85
	default:
		v591 = v526
		goto L79
	case 3:
		goto L84
	case 4:
		goto L83
	case 10, 11:
		goto L87
	case 24:
		goto L86
	case 29, 31:
		goto L82
	}
L81:
	;
	v591 = int32(1)
	goto L79
L82:
	;
	goto L81
L83:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v527)+72))
	if v581 == int32(0) {
		v591 = v526
		goto L79
	} else {
		goto L91
	}
L84:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v527)+72))
	if v573 == int32(0) {
		v591 = v526
		goto L79
	} else {
		goto L89
	}
L85:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	if v569 != int32(301) {
		v591 = v526
		goto L79
	} else {
		goto L88
	}
L86:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+72)))
	v591 = int32(base.Ui32(v564&int32(2)) >> (uint(int32(1)) % 32))
	goto L79
L87:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v527)+72))
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+112)))
	v591 = v563
	goto L79
L88:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v527)+72))
	v527 = v572
	goto L80
L89:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v576 != int32(1) {
		v591 = v526
		goto L79
	} else {
		goto L90
	}
L90:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	v527 = v580
	goto L80
L91:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	if v584 != int32(1) {
		v591 = v526
		goto L79
	} else {
		goto L92
	}
L92:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v581)+12))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v527 = v588
	goto L80
L93:
	;
	v648 = v510
	v663 = v521
	goto L72
L94:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_create_mergejoin_path[4]))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+32))
	if base.F64_lt(base.F64_convert_i32_u(v597<<(uint(int32(10))%32)), base.F64_mul(v491, base.F64_convert_i32_u((v602+int32(7))&int32(-8)+int32(24)))) != 0 {
		v648 = v510
		v663 = v521
		goto L72
	} else {
		goto L95
	}
L95:
	;
	goto L73
}
func F_crlf_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	if l3 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v10 = l2 + l3
	v13 = l2
	goto L4
L4:
	;
	v17 = v10 - v13
	v18 = int32(0)
	v21 = base.B2i32(v17 != v18)
	if v13&int32(3) == v18 {
		v47 = v13
		v49 = v17
		v50 = v21
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return v163
L6:
	;
	goto L5
L7:
	;
	if base.Ui32(v10) <= base.Ui32(v134) {
		v157 = v133
		v158 = v134
		goto L43
	} else {
		goto L44
	}
L8:
	;
	if v120 != 0 {
		goto L34
	} else {
		goto L35
	}
L9:
	;
	v120 = int32(0)
	goto L8
L10:
	;
	v98 = v91
	v100 = v93
	goto L28
L11:
	;
	if v50 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L12:
	;
	if v17 == int32(0) {
		v47 = v13
		v49 = v17
		v50 = v21
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = v13
	v32 = v17
	goto L14
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 == int32(10) {
		v91 = v30
		v93 = v32
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v47 = v42
	v49 = v38
	v50 = v40
	goto L11
L16:
	;
	v37 = int32(1)
	v38 = v32 - v37
	v39 = int32(0)
	v40 = base.B2i32(v38 != v39)
	v42 = v30 + v37
	if v42&int32(3) == v39 {
		v47 = v42
		v49 = v38
		v50 = v40
		goto L11
	} else {
		goto L17
	}
L17:
	;
	if v38 != 0 {
		v30 = v42
		v32 = v38
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v54 == int32(10) {
		v84 = v47
		v86 = v49
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v86 == int32(0) {
		goto L9
	} else {
		goto L27
	}
L21:
	;
	if base.Ui32(v49) < base.Ui32(int32(4)) {
		v84 = v47
		v86 = v49
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v64 = v47
	v66 = v49
	goto L23
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = v70 ^ int32(168430090)
	v74 = int32(-2139062144)
	if (int32(16843008)-v71|v71)&v74 != v74 {
		v91 = v64
		v93 = v66
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v84 = v79
	v86 = v81
	goto L20
L25:
	;
	v78 = int32(4)
	v79 = v64 + v78
	v81 = v66 - v78
	if base.Ui32(int32(3)) < base.Ui32(v81) {
		v64 = v79
		v66 = v81
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v91 = v84
	v93 = v86
	goto L10
L28:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if int32(10) == v103 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L9
L30:
	;
	v120 = v98
	goto L8
L31:
	;
	goto L32
L32:
	;
	v105 = int32(1)
	v108 = v100 - v105
	if v108 != 0 {
		v98 = v98 + v105
		v100 = v108
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v121 = v120
	goto L36
L35:
	;
	v121 = v10
	goto L36
L36:
	;
	v122 = v121 - v13
	if v122 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v133 = int32(0)
	v134 = v13
	goto L7
L38:
	;
	goto L39
L39:
	;
	v126 = F_pushf_write(m, l0, v13, v122)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	return int32(0)
L41:
	;
	if v126 < int32(0) {
		v163 = v126
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v133 = v126
	v134 = v13 + v122
	goto L7
L43:
	;
	if base.Ui32(v158) < base.Ui32(v10) {
		v13 = v158
		goto L4
	} else {
		goto L54
	}
L44:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v136 != int32(10) {
		v157 = v133
		v158 = v134
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v141 = v134
	goto L46
L46:
	;
	v146 = F_pushf_write(m, l0, int32(_a_F_crlf_process_0), int32(2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L40
	} else {
		goto L48
	}
L47:
	;
	v157 = v146
	v158 = v10
	goto L43
L48:
	;
	if v146 < int32(0) {
		v157 = v146
		v158 = v141
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v151 = v141 + int32(1)
	if base.Ui32(v151) < base.Ui32(v10) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v153 != int32(10) {
		v157 = v146
		v158 = v151
		goto L43
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L47
L53:
	;
	v141 = v151
	goto L46
L54:
	;
	v163 = v157
	goto L6
}
func F_crosstab(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int64
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
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
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v232 int64
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v337 int32
	_ = v337
	var v353 int64
	_ = v353
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int64
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v430 int64
	_ = v430
	var v434 int32
	_ = v434
	var v454 int64
	_ = v454
	var v476 int64
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v504 int64
	_ = v504
	var v508 int32
	_ = v508
	var v526 int64
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v547 int64
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v607 int32
	_ = v607
	var v610 int64
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_pg_detoast_datum_packed(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = F_text_to_cstring(m, v26)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v32 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L180
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v35 != int32(383) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+12)))
	if v38&int32(2) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L176
	}
L10:
	;
	v48 = F_SPI_execute(m, v30, int32(1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L171
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L167
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L162
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L155
	}
L15:
	;
	m.G0 = v23 + int32(32)
	return int32(0)
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab[0]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v66 != int32(3) {
		goto L11
	} else {
		goto L23
	}
L17:
	;
	if v48 == int32(5) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v53 = *(*int64)(unsafe.Add(mBase, _c_F_crosstab[1]))
	if v53 != int64(0) {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v57 = F_SPI_finish(m)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(2)
	v61 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v61)
	goto L15
L23:
	;
	v72 = F_get_call_result_type(m, l0, int32(0), v23+int32(28))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v72 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v72 == int32(3) {
		goto L12
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 <= int32(1) {
		goto L13
	} else {
		goto L33
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_crosstab_0), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(442), int32(_a_F_crosstab_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v98 = int32(4)
	v100 = v94 + v95<<(uint(v98)%32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+96))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v105 = v65 + v102<<(uint(v98)%32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+96))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+88))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+88))
	if v107 != v108 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	if base.B2i32(v101 != v106)&base.B2i32(int32(0) <= v101) != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v105)+296))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v105)+288))
	v119 = int32(1)
	goto L36
L36:
	;
	v141 = v100 + int32(20) + v119*int32(100)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)+68))
	if v117 == v143 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v180 = int32(_a_F_crosstab_3)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab[2])) = v42
	v184 = F_CreateTupleDescCopy(m, v94)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L52
	}
L38:
	;
	v178 = v119 + int32(1)
	if v178 != v95 {
		v119 = v178
		goto L36
	} else {
		goto L51
	}
L39:
	;
	if v142 < int32(0) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	v148 = v143
	goto L41
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	if v142 == v116 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v148 = v117
	goto L41
L44:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_crosstab_4), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v160 = F_format_type_with_typemod(m, v117, v116)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v162 = F_format_type_with_typemod(m, v148, v142)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v119 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v160
	F_errdetail(m, int32(_a_F_crosstab_5), v23)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(1569), int32(_a_F_crosstab_6))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	goto L37
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v184
	v187 = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab[3]))
	v196 = F_tuplestore_begin_heap(m, int32(base.Ui32(v188&int32(4))>>(uint(int32(2))%32)), v187, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab[2])) = v181
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v201 = F_TupleDescGetAttInMetadata(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v205 = int32(2)
	v209 = int32(1)
	v218 = v187
	v220 = v209
	v232 = int64(0)
	goto L55
L55:
	;
	v235 = F_palloc0(m, v204<<(uint(v205)%32))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v612
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = int32(2)
	v617 = F_SPI_finish(m)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L154
	}
L57:
	;
	if v204 < int32(2) {
		v476 = v232
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v550 = int32(0)
	if v531 != 0 {
		goto L138
	} else {
		goto L139
	}
L59:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v531 = v529
	v547 = v526
	goto L58
L60:
	;
	F_pfree(m, v218)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L137
	}
L61:
	;
	v479 = F_BuildTupleFromCStrings(m, v201, v235)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L133
	}
L62:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239+base.I32_wrap_i64(v232)<<(uint(int32(2))%32))))
	v246 = F_SPI_getvalue(m, v244, v65, int32(1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v246 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v248 = F_pstrdup(m, v246)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	v251 = int32(0)
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v251
	if v220&int32(1) != 0 {
		goto L73
	} else {
		goto L74
	}
L67:
	;
	v251 = v248
	goto L66
L68:
	;
	v476 = v454 - int64(1)
	goto L61
L69:
	;
	F_pfree(m, v413)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L132
	}
L70:
	;
	v322 = v235 + int32(4)
	v324 = F_SPI_getvalue(m, v244, v65, int32(3))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L100
	}
L71:
	;
	if v251 != 0 {
		v454 = v232
		goto L68
	} else {
		goto L99
	}
L72:
	;
	if v251 == int32(0) {
		v413 = v246
		v430 = v232
		goto L69
	} else {
		goto L89
	}
L73:
	;
	if v246 == int32(0) {
		goto L71
	} else {
		goto L88
	}
L74:
	;
	if v246|v218 == int32(0) {
		v531 = v251
		v547 = v232
		goto L58
	} else {
		goto L75
	}
L75:
	;
	if v218 == int32(0) {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	if v246 == int32(0) {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v265 == int32(0) {
		v284 = v264
		v285 = v265
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v285-v284 != 0 {
		goto L72
	} else {
		goto L86
	}
L79:
	;
	goto L78
L80:
	;
	if v264 != v265 {
		v284 = v264
		v285 = v265
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v269 = v218
	v270 = v246
	goto L82
L82:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+1)))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	if v274 == int32(0) {
		v284 = v273
		v285 = v274
		goto L79
	} else {
		goto L84
	}
L83:
	;
	v284 = v273
	v285 = v274
	goto L79
L84:
	;
	v277 = int32(1)
	if v273 == v274 {
		v269 = v269 + v277
		v270 = v270 + v277
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	F_pfree(m, v246)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v504 = v232
	goto L60
L88:
	;
	goto L72
L89:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v296 == int32(0) {
		v315 = v295
		v316 = v296
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v316-v315 != 0 {
		v413 = v246
		v430 = v232
		goto L69
	} else {
		goto L98
	}
L91:
	;
	goto L90
L92:
	;
	if v295 != v296 {
		v315 = v295
		v316 = v296
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v300 = v246
	v301 = v251
	goto L94
L94:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v305 == int32(0) {
		v315 = v304
		v316 = v305
		goto L91
	} else {
		goto L96
	}
L95:
	;
	v315 = v304
	v316 = v305
	goto L91
L96:
	;
	v308 = int32(1)
	if v304 == v305 {
		v300 = v300 + v308
		v301 = v301 + v308
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v320 = int32(0)
	goto L70
L99:
	;
	v320 = int32(1)
	goto L70
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = v324
	if v320 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	F_pfree(m, v246)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v331 = v232 + base.I64_extend_i32_u(base.B2i32(v204 != v205))
	if v204 == int32(2) {
		v476 = v331
		goto L61
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	if base.Ui64(v53) <= base.Ui64(v331) {
		v476 = v331
		goto L61
	} else {
		goto L106
	}
L106:
	;
	v337 = int32(1)
	v353 = v331
	goto L107
L107:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v356+base.I32_wrap_i64(v353)<<(uint(int32(2))%32))))
	v363 = F_SPI_getvalue(m, v361, v65, int32(1))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	v476 = v408
	goto L61
L109:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	if v363 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v401 = F_SPI_getvalue(m, v361, v65, int32(3))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L125
	}
L111:
	;
	if v365 == int32(0) {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v365 == int32(0) {
		v413 = v363
		v430 = v353
		goto L69
	} else {
		goto L115
	}
L114:
	;
	v454 = v353
	goto L68
L115:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	if v375 == int32(0) {
		v394 = v374
		v395 = v375
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v395-v394 != 0 {
		v413 = v363
		v430 = v353
		goto L69
	} else {
		goto L124
	}
L117:
	;
	goto L116
L118:
	;
	if v374 != v375 {
		v394 = v374
		v395 = v375
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v379 = v363
	v380 = v365
	goto L120
L120:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+1)))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+1)))
	if v384 == int32(0) {
		v394 = v383
		v395 = v384
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v394 = v383
	v395 = v384
	goto L117
L122:
	;
	v387 = int32(1)
	if v383 == v384 {
		v379 = v379 + v387
		v380 = v380 + v387
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	goto L110
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322+v337<<(uint(int32(2))%32)))) = v401
	if v363 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_pfree(m, v363)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v408 = base.I64_extend_i32_u(base.B2i32(v337 < v204-v205)) + v353
	v410 = v337 + int32(1)
	if v204-v209 <= v410 {
		v476 = v408
		goto L61
	} else {
		goto L130
	}
L129:
	;
	goto L128
L130:
	;
	if base.Ui64(v408) < base.Ui64(v53) {
		v337 = v410
		v353 = v408
		goto L107
	} else {
		goto L131
	}
L131:
	;
	goto L108
L132:
	;
	v476 = v430 - int64(1)
	goto L61
L133:
	;
	F_tuplestore_puttuple(m, v196, v479)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_pfree(m, v479)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	if v218 == int32(0) {
		v526 = v476
		goto L59
	} else {
		goto L136
	}
L136:
	;
	v504 = v476
	goto L60
L137:
	;
	v526 = v504
	goto L59
L138:
	;
	v552 = F_pstrdup(m, v531)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	v554 = v550
	goto L140
L140:
	;
	if int32(0) < v204 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v554 = v552
	goto L140
L142:
	;
	v557 = v550
	goto L145
L143:
	;
	goto L144
L144:
	;
	F_pfree(m, v235)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L152
	}
L145:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v235+v557<<(uint(int32(2))%32))))
	if v580 != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L144
L147:
	;
	F_pfree(m, v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v584 = v557 + int32(1)
	if v584 != v204 {
		v557 = v584
		goto L145
	} else {
		goto L151
	}
L150:
	;
	goto L149
L151:
	;
	goto L146
L152:
	;
	v610 = v547 + int64(1)
	if base.Ui64(v610) < base.Ui64(v53) {
		v218 = v554
		v220 = int32(0)
		v232 = v610
		goto L55
	} else {
		goto L153
	}
L153:
	;
	goto L56
L154:
	;
	goto L15
L155:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_crosstab_4), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v655 = F_format_type_with_typemod(m, v108, v106)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v657 = F_format_type_with_typemod(m, v107, v101)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v657
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v655
	F_errdetail(m, int32(_a_F_crosstab_7), v23+int32(16))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(1547), int32(_a_F_crosstab_6))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errmsg(m, int32(_a_F_crosstab_4), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errdetail(m, int32(_a_F_crosstab_8), int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(1532), int32(_a_F_crosstab_6))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errmsg(m, int32(_a_F_crosstab_9), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(436), int32(_a_F_crosstab_2))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errmsg(m, int32(_a_F_crosstab_10), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	F_errdetail(m, int32(_a_F_crosstab_11), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(423), int32(_a_F_crosstab_2))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errmsg(m, int32(_a_F_crosstab_12), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(386), int32(_a_F_crosstab_2))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_errmsg(m, int32(_a_F_crosstab_13), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(382), int32(_a_F_crosstab_2))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
