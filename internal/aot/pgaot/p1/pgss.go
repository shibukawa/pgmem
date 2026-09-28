package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pgss_ExecutorStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[0]))
	if int32(0) <= v5 {
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[1]))
		if v9 != int32(2) {
			if v9 != int32(1) {
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[2]))
				if v15 != 0 {
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
					if v17 == int64(0) {
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v20 | int32(2147483647)
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
			if v17 == int64(0) {
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v20 | int32(2147483647)
			}
		}
	}
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[3]))
	if v26 != 0 {
		m.T0[v26].(func(*base.Module, int32, int32))(m, l0, l1)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			return
		}
	} else {
		F_standard_ExecutorStart(m, l0, l1)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pgss_store(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 float64, l6 int64, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int64
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v215 int32
	_ = v215
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v319 int32
	_ = v319
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v384 int32
	_ = v384
	var v401 int32
	_ = v401
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v518 int32
	_ = v518
	var v529 int32
	_ = v529
	var v549 int32
	_ = v549
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
	var v575 int32
	_ = v575
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int64
	_ = v622
	var v623 int32
	_ = v623
	var v628 int64
	_ = v628
	var v632 int64
	_ = v632
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int64
	_ = v680
	var v681 int32
	_ = v681
	var v685 int64
	_ = v685
	var v689 int64
	_ = v689
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v976 int32
	_ = v976
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1150 int32
	_ = v1150
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1226 int64
	_ = v1226
	var v1228 int64
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int64
	_ = v1238
	var v1240 int64
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1245 float64
	_ = v1245
	var v1254 int32
	_ = v1254
	var v1255 float64
	_ = v1255
	var v1256 float64
	_ = v1256
	var v1259 float64
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1265 float64
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1270 float64
	_ = v1270
	var v1277 int32
	_ = v1277
	var v1278 float64
	_ = v1278
	var v1290 int32
	_ = v1290
	var v1291 float64
	_ = v1291
	var v1300 int64
	_ = v1300
	var v1303 int64
	_ = v1303
	var v1304 int64
	_ = v1304
	var v1307 int64
	_ = v1307
	var v1308 int64
	_ = v1308
	var v1311 int64
	_ = v1311
	var v1312 int64
	_ = v1312
	var v1315 int64
	_ = v1315
	var v1316 int64
	_ = v1316
	var v1319 int64
	_ = v1319
	var v1320 int64
	_ = v1320
	var v1323 int64
	_ = v1323
	var v1324 int64
	_ = v1324
	var v1327 int64
	_ = v1327
	var v1328 int64
	_ = v1328
	var v1331 int64
	_ = v1331
	var v1332 int64
	_ = v1332
	var v1335 int64
	_ = v1335
	var v1336 int64
	_ = v1336
	var v1339 int64
	_ = v1339
	var v1340 int64
	_ = v1340
	var v1343 float64
	_ = v1343
	var v1344 int64
	_ = v1344
	var v1346 float64
	_ = v1346
	var v1350 float64
	_ = v1350
	var v1351 int64
	_ = v1351
	var v1357 float64
	_ = v1357
	var v1358 int64
	_ = v1358
	var v1364 float64
	_ = v1364
	var v1365 int64
	_ = v1365
	var v1371 float64
	_ = v1371
	var v1372 int64
	_ = v1372
	var v1378 float64
	_ = v1378
	var v1379 int64
	_ = v1379
	var v1385 float64
	_ = v1385
	var v1389 int64
	_ = v1389
	var v1390 int64
	_ = v1390
	var v1393 int64
	_ = v1393
	var v1394 int64
	_ = v1394
	var v1397 int64
	_ = v1397
	var v1398 int64
	_ = v1398
	var v1401 int64
	_ = v1401
	var v1402 int64
	_ = v1402
	var v1405 int64
	_ = v1405
	var v1406 int64
	_ = v1406
	var v1409 float64
	_ = v1409
	var v1410 int64
	_ = v1410
	var v1412 float64
	_ = v1412
	var v1416 int64
	_ = v1416
	var v1419 float64
	_ = v1419
	var v1422 int64
	_ = v1422
	var v1426 float64
	_ = v1426
	var v1429 int64
	_ = v1429
	var v1432 float64
	_ = v1432
	var v1435 int64
	_ = v1435
	var v1439 float64
	_ = v1439
	var v1442 int64
	_ = v1442
	var v1445 float64
	_ = v1445
	var v1448 int64
	_ = v1448
	var v1452 float64
	_ = v1452
	var v1455 int64
	_ = v1455
	var v1458 float64
	_ = v1458
	var v1461 int64
	_ = v1461
	var v1465 float64
	_ = v1465
	var v1469 int64
	_ = v1469
	var v1473 int64
	_ = v1473
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int64
	_ = v1483
	var v1489 int32
	_ = v1489
	var v1494 int32
	_ = v1494
	var v1506 int32
	_ = v1506
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1543 int32
	_ = v1543
	v15 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(176)
	m.G0 = v32
	*(*int32)(unsafe.Add(mBase, uint32(v32)+144)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v32)+148)) = l2
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	goto L1
L1:
	;
	if l1 == int64(0) {
		v1543 = v32
		goto L2
	} else {
		goto L3
	}
L2:
	;
	m.G0 = v1543 + int32(176)
	return
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	if v42 == int32(0) {
		v1543 = v32
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	if v46 == int32(0) {
		v1543 = v32
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v50 = v32 + int32(148)
	v52 = v32 + int32(144)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v57 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v123 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+136)) = v123
	*(*int64)(unsafe.Add(mBase, uint32(v32)+128)) = v123
	*(*int64)(unsafe.Add(mBase, uint32(v32)+120)) = v123
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v32)+128)) = l1
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v134
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[5]))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+136)) = uint8(base.B2i32(v137 == int32(0)))
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v144 = F_LWLockAcquire(m, v142, int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v117
	goto L6
L8:
	;
	v75 = v71
	v78 = v72
	v79 = v73
	goto L15
L9:
	;
	v68 = F_strlen(m, v65)
	mBase = m.M
	if v68 <= int32(0) {
		v114 = v65
		v117 = v68
		v118 = v67
		goto L7
	} else {
		goto L14
	}
L10:
	;
	v65 = l0
	v67 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v61 = l0 + v57
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if int32(0) < v62 {
		v71 = v61
		v72 = v62
		v73 = v57
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v65 = v61
	v67 = v57
	goto L9
L14:
	;
	v71 = v65
	v72 = v68
	v73 = v67
	goto L8
L15:
	;
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v75))))
	v83 = F_scanner_isspace(m, v82)
	mBase = m.M
	if v83 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v114 = v108
	v117 = int32(0)
	v118 = v72 + v73
	goto L7
L17:
	;
	v89 = v78
	goto L20
L18:
	;
	goto L19
L19:
	;
	v105 = int32(1)
	v108 = v75 + v105
	if v105 < v78 {
		v75 = v108
		v78 = v78 - v105
		v79 = v79 + v105
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v75+v89-int32(1)))))
	v97 = F_scanner_isspace(m, v96)
	mBase = m.M
	if v97 == int32(0) {
		v114 = v75
		v117 = v89
		v118 = v79
		goto L7
	} else {
		goto L22
	}
L21:
	;
	v114 = v75
	v117 = int32(0)
	v118 = v79
	goto L7
L22:
	;
	v100 = int32(1)
	if v100 < v89 {
		v89 = v89 - v100
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L16
L25:
	;
	return
L26:
	;
	v146 = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	v153 = F_hash_search(m, v148, v32+int32(120), v146, v146)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	F_LWLockRelease(m, v1522)
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L25
	} else {
		goto L258
	}
L28:
	;
	if l10 != 0 {
		v1494 = v1191
		v1506 = v32
		goto L27
	} else {
		goto L221
	}
L29:
	;
	if v153 != 0 {
		v1191 = v146
		v1192 = v153
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if l10 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	F_LWLockRelease(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	v575 = v146
	goto L33
L33:
	;
	if v575 != 0 {
		goto L94
	} else {
		goto L95
	}
L34:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v32)+144))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	F_initStringInfoExt(m, v32+int32(152), v162+v163*int32(10))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v169 = m.G0
	v171 = v169 + int32(-64)
	m.G0 = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if v173 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v175 = F_palloc_mul(m, int32(12), v173)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L25
	} else {
		goto L39
	}
L37:
	;
	v401 = v15
	goto L38
L38:
	;
	m.G0 = v171 - int32(-64)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if int32(0) < v417 {
		goto L73
	} else {
		goto L74
	}
L39:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v179 = v177 * int32(12)
	if v179 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	base.MemoryCopy(m, v175, v180, v179)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if int32(2) <= v182 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pg_qsort(m, v175, v182, int32(12), int32(864))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L25
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v193 = F_scanner_init(m, v114, v169+int32(-56), int32(_a_F_pgss_store_0), int32(_a_F_pgss_store_1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L25
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if v195 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_scanner_finish(m, v193)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L25
	} else {
		goto L70
	}
L49:
	;
	v215 = v15
	goto L50
L50:
	;
	v229 = v175 + v215*int32(12)
	if v215 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	goto L48
L52:
	;
	v351 = v215 + int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if v351 < v352 {
		v215 = v351
		goto L50
	} else {
		goto L69
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v319
	goto L52
L54:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v229-int32(12))))
	if v231 == v234 {
		v319 = int32(-1)
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v238 = v175 + v215*int32(12)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+8)))
	if v239 != 0 {
		goto L52
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v241 = v240 - v159
	goto L59
L59:
	;
	v272 = v169 + int32(-60)
	v273 = F_core_yylex(m, v272, v171, v193)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L25
	} else {
		goto L61
	}
L60:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v114))))
	if v280 == int32(45) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if v273 == int32(0) {
		goto L48
	} else {
		goto L62
	}
L62:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	if v277 < v241 {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v283 = F_core_yylex(m, v272, v171, v193)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L25
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	v289 = F_strlen(m, v287+v241)
	mBase = m.M
	v319 = v289
	goto L53
L67:
	;
	if v283 == int32(0) {
		goto L48
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	goto L51
L70:
	;
	v401 = v175
	goto L38
L71:
	;
	F_appendBinaryStringInfo(m, v32+int32(152), v114+v549, v162-v549)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L25
	} else {
		goto L92
	}
L72:
	;
	F_pfree(m, v401)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L25
	} else {
		goto L91
	}
L73:
	;
	v421 = v417
	v424 = int32(0)
	v436 = v15
	v440 = v15
	v442 = v15
	v446 = v15
	goto L76
L74:
	;
	goto L75
L75:
	;
	if v401 == int32(0) {
		v549 = v15
		goto L71
	} else {
		goto L90
	}
L76:
	;
	v452 = v401 + v424*int32(12)
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452)+9)))
	if v453 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v518 = v491
	goto L72
L78:
	;
	v495 = v424 + int32(1)
	if v495 < v488 {
		v421 = v488
		v424 = v495
		v436 = v490
		v440 = v491
		v442 = v492
		v446 = v493
		goto L76
	} else {
		goto L89
	}
L79:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10)+24)))
	if v456 != int32(1) {
		v488 = v421
		v490 = v436
		v491 = v440
		v492 = v442
		v493 = v446
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if v459 < int32(0) {
		v488 = v421
		v490 = v436
		v491 = v440
		v492 = v442
		v493 = v446
		goto L78
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v464 = v32 + int32(152)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	v467 = v466 - v159
	F_appendBinaryStringInfo(m, v464, v114+v440, v467-(v436+v442))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L25
	} else {
		goto L84
	}
L84:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452)+8)))
	if v474 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v475 = int32(_a_F_pgss_store_2)
	goto L87
L86:
	;
	v475 = int32(_a_F_pgss_store_3)
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v475
	v478 = v446 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v471 + v478
	F_appendStringInfo(m, v464, int32(_a_F_pgss_store_4), v32+int32(96))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L25
	} else {
		goto L88
	}
L88:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v488 = v487
	v490 = v459
	v491 = v459 + v467
	v492 = v467
	v493 = v478
	goto L78
L89:
	;
	goto L77
L90:
	;
	v518 = v15
	goto L72
L91:
	;
	v549 = v518
	goto L71
L92:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v32)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+144)) = v565
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v32)+152))
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v571 = F_LWLockAcquire(m, v569, int32(1))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L25
	} else {
		goto L93
	}
L93:
	;
	v575 = v567
	goto L33
L94:
	;
	v602 = v575
	goto L96
L95:
	;
	v602 = v114
	goto L96
L96:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v32)+144))
	v608 = F_qtext_store(m, v602, v603, v32+int32(116), v32+int32(112))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L25
	} else {
		goto L97
	}
L97:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v614 = base.AtomicRmwXchg32(m, v611, int32(140), int32(1))
	if v614 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_s_lock(m, v611+int32(140), int32(_a_F_pgss_store_5))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L25
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v622 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v621)+144)))
	v623 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v621)+140)), uint32(v623))
	v628 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgss_store[6])))
	if base.Ui64(v628<<(uint(int64(9))%64)) <= base.Ui64(v622) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	v632 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v621)+136)))
	v637 = base.B2i32(base.Ui64(v628*v632<<(uint(int64(1))%64)) <= base.Ui64(v622))
	goto L104
L103:
	;
	v637 = v623
	goto L104
L104:
	;
	F_LWLockRelease(m, v621)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L25
	} else {
		goto L105
	}
L105:
	;
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v643 = F_LWLockAcquire(m, v641, int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L25
	} else {
		goto L106
	}
L106:
	;
	if v608 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v32)+116))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v32)+144))
	v664 = F_entry_alloc(m, v32+int32(120), v660, v661, v38, base.B2i32(l10 != int32(0)))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L25
	} else {
		goto L114
	}
L108:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)+152))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
	if v647 == v648 {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v32)+144))
	v654 = F_qtext_store(m, v602, v650, v32+int32(116), int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L25
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	if v654 == int32(0) {
		v1494 = v575
		v1506 = v32
		goto L27
	} else {
		goto L113
	}
L113:
	;
	goto L107
L114:
	;
	if v637 == int32(0) {
		v1191 = v575
		v1192 = v664
		goto L28
	} else {
		goto L115
	}
L115:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v672 = base.AtomicRmwXchg32(m, v669, int32(140), int32(1))
	if v672 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	F_s_lock(m, v669+int32(140), int32(_a_F_pgss_store_5))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L25
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v680 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v679)+144)))
	v681 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v679)+140)), uint32(v681))
	v685 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgss_store[6])))
	if base.Ui64(v680) < base.Ui64(v685<<(uint(int64(9))%64)) {
		v1191 = v575
		v1192 = v664
		goto L28
	} else {
		goto L120
	}
L119:
	;
	goto L118
L120:
	;
	v689 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v679)+136)))
	if base.Ui64(v680) < base.Ui64(v685*v689<<(uint(int64(1))%64)) {
		v1191 = v575
		v1192 = v664
		goto L28
	} else {
		goto L121
	}
L121:
	;
	v696 = F_qtext_load_file(m, v32+int32(172))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L25
	} else {
		goto L124
	}
L122:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1181)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v1181)+152)) = v1182 + int32(1)
	v1186 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1181)+140)), uint32(v1186))
	v1191 = v575
	v1192 = v664
	goto L28
L123:
	;
	F_s_lock(m, v1117+int32(140), int32(_a_F_pgss_store_5))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L25
	} else {
		goto L220
	}
L124:
	;
	if v696 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v700 = F_AllocateFile(m, int32(_a_F_pgss_store_6), int32(_a_F_pgss_store_7))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L25
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v1007 = v32 + int32(152)
	v1009 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	F_hash_seq_init(m, v1007, v1009)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L25
	} else {
		goto L199
	}
L128:
	;
	F_pfree(m, v696)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L25
	} else {
		goto L198
	}
L129:
	;
	if v700 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v706 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L25
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v725 = v32 + int32(152)
	v727 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	F_hash_seq_init(m, v725, v727)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L25
	} else {
		goto L138
	}
L133:
	;
	if v706 == int32(0) {
		goto L128
	} else {
		goto L134
	}
L134:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L25
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = int32(_a_F_pgss_store_6)
	F_errmsg(m, int32(_a_F_pgss_store_8), v32+int32(16))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L25
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_pgss_store_9), int32(2516), int32(_a_F_pgss_store_10))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L25
	} else {
		goto L137
	}
L137:
	;
	goto L128
L138:
	;
	v730 = F_hash_seq_search(m, v725)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L25
	} else {
		goto L140
	}
L139:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v700)+60))
	if v855 < int32(0) {
		goto L169
	} else {
		goto L170
	}
L140:
	;
	if v730 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v734 = int32(0)
	v841 = v734
	v842 = v734
	goto L139
L142:
	;
	goto L143
L143:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v32)+172))
	v737 = int32(0)
	v739 = v730
	v754 = v737
	v755 = v737
	goto L144
L144:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v739)+412))
	if v768 < int32(0) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v841 = v817
	v842 = v818
	goto L139
L146:
	;
	v824 = F_hash_seq_search(m, v32+int32(152))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L25
	} else {
		goto L165
	}
L147:
	;
	v783 = int32(1)
	v785 = v768 + v783
	v786 = F_fwrite(m, v696+v771, v783, v785, v700)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L25
	} else {
		goto L152
	}
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v739)+408)) = int64(-4294967296)
	v817 = v754
	v818 = v755
	goto L146
L149:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v739)+408))
	v772 = v771 + v768
	if base.Ui32(v736) <= base.Ui32(v772) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+v772))))
	if v775 == int32(0) {
		goto L147
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	if v786 != v785 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v791 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L25
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v739)+408)) = v754
	v817 = v754 + v785
	v818 = v755 + int32(1)
	goto L146
L156:
	;
	if v791 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L25
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	F_hash_seq_term(m, v32+int32(152))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L25
	} else {
		goto L163
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = int32(_a_F_pgss_store_6)
	F_errmsg(m, int32(_a_F_pgss_store_8), v32+int32(80))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L25
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_pgss_store_9), int32(2546), int32(_a_F_pgss_store_10))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L25
	} else {
		goto L162
	}
L162:
	;
	goto L159
L163:
	;
	v811 = F_FreeFile(m, v700)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L25
	} else {
		goto L164
	}
L164:
	;
	goto L128
L165:
	;
	if v824 != 0 {
		v739 = v824
		v754 = v817
		v755 = v818
		goto L144
	} else {
		goto L166
	}
L166:
	;
	goto L145
L167:
	;
	v887 = F_FreeFile(m, v700)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L25
	} else {
		goto L178
	}
L168:
	;
	v864 = F_ftruncate(m, v862, base.I64_extend_i32_u(v841))
	mBase = m.M
	if v864 == int32(0) {
		goto L167
	} else {
		goto L172
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_store[7])) = int32(8)
	v862 = int32(-1)
	goto L171
L170:
	;
	v862 = v855
	goto L171
L171:
	;
	goto L168
L172:
	;
	v869 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L25
	} else {
		goto L173
	}
L173:
	;
	if v869 == int32(0) {
		goto L167
	} else {
		goto L174
	}
L174:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L25
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = int32(_a_F_pgss_store_6)
	F_errmsg(m, int32(_a_F_pgss_store_11), v32-int32(-64))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L25
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_pgss_store_9), int32(2564), int32(_a_F_pgss_store_10))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L25
	} else {
		goto L177
	}
L177:
	;
	goto L167
L178:
	;
	if v887 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v891 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L25
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v911 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L25
	} else {
		goto L187
	}
L182:
	;
	if v891 == int32(0) {
		goto L128
	} else {
		goto L183
	}
L183:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L25
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = int32(_a_F_pgss_store_6)
	F_errmsg(m, int32(_a_F_pgss_store_8), v32+int32(48))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L25
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(_a_F_pgss_store_9), int32(2571), int32(_a_F_pgss_store_10))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L25
	} else {
		goto L186
	}
L186:
	;
	goto L128
L187:
	;
	if v911 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v914)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v915
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v841
	F_errmsg_internal(m, int32(_a_F_pgss_store_12), v32+int32(32))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L25
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v929 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v929)+144)) = v841
	if v842 <= int32(0) {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	F_errfinish(m, int32(_a_F_pgss_store_9), int32(2577), int32(_a_F_pgss_store_10))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L25
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	v935 = int32(1024)
	goto L195
L194:
	;
	v934 = base.I32_div_u_s(v841, v842)
	v935 = v934
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v929)+136)) = v935
	F_pfree(m, v696)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L25
	} else {
		goto L196
	}
L196:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v943 = base.AtomicRmwXchg32(m, v940, int32(140), int32(1))
	if v943 == int32(0) {
		goto L122
	} else {
		goto L197
	}
L197:
	;
	v1117 = v940
	goto L123
L198:
	;
	goto L127
L199:
	;
	v1012 = F_hash_seq_search(m, v1007)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L25
	} else {
		goto L200
	}
L200:
	;
	if v1012 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1014 = v1012
	goto L204
L202:
	;
	goto L203
L203:
	;
	v1078 = int32(_a_F_pgss_store_6)
	v1079 = F_unlink(m, v1078)
	mBase = m.M
	v1082 = F_AllocateFile(m, v1078, int32(_a_F_pgss_store_7))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L25
	} else {
		goto L209
	}
L204:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1014)+408)) = int64(-4294967296)
	v1047 = F_hash_seq_search(m, v32+int32(152))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L25
	} else {
		goto L206
	}
L205:
	;
	goto L203
L206:
	;
	if v1047 != 0 {
		v1014 = v1047
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v1107)+136)) = int32(1024)
	v1110 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1107)+144)) = v1110
	v1114 = base.AtomicRmwXchg32(m, v1107, int32(140), int32(1))
	if v1114 == v1110 {
		goto L122
	} else {
		goto L219
	}
L209:
	;
	if v1082 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1088 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L25
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v1104 = F_FreeFile(m, v1082)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L25
	} else {
		goto L218
	}
L213:
	;
	if v1088 == int32(0) {
		goto L208
	} else {
		goto L214
	}
L214:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L25
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(_a_F_pgss_store_6)
	F_errmsg(m, int32(_a_F_pgss_store_13), v32)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L25
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(_a_F_pgss_store_9), int32(2631), int32(_a_F_pgss_store_10))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L25
	} else {
		goto L217
	}
L217:
	;
	goto L208
L218:
	;
	goto L208
L219:
	;
	v1117 = v1107
	goto L123
L220:
	;
	goto L122
L221:
	;
	v1218 = int32(440)
	v1219 = v1192 + v1218
	v1222 = base.AtomicRmwXchg32(m, v1192, v1218, int32(1))
	if v1222 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	F_s_lock(m, v1219, int32(_a_F_pgss_store_5))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L25
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1226 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+24))
	v1228 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+32))
	if v1226 == int64(0)-v1228 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L224
L226:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+256)) = int64(4607182418800017408)
	goto L228
L227:
	;
	goto L228
L228:
	;
	v1234 = l4 << (uint(int32(3)) % 32)
	v1237 = v1234 + (v1192 + int32(24))
	v1238 = *(*int64)(unsafe.Add(mBase, uint32(v1237)))
	v1240 = v1238 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1237))) = v1240
	v1242 = v1234 + v1192
	v1244 = v1242 + int32(40)
	v1245 = *(*float64)(unsafe.Add(mBase, uint32(v1244)))
	*(*float64)(unsafe.Add(mBase, uint32(v1244))) = base.F64_add(l5, v1245)
	if v1238 == int64(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1300 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+120)) = v1300 + l6
	v1303 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+128))
	v1304 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+128)) = v1303 + v1304
	v1307 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+136))
	v1308 = *(*int64)(unsafe.Add(mBase, uint32(l7)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+136)) = v1307 + v1308
	v1311 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+144))
	v1312 = *(*int64)(unsafe.Add(mBase, uint32(l7)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+144)) = v1311 + v1312
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+152))
	v1316 = *(*int64)(unsafe.Add(mBase, uint32(l7)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+152)) = v1315 + v1316
	v1319 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+160))
	v1320 = *(*int64)(unsafe.Add(mBase, uint32(l7)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+160)) = v1319 + v1320
	v1323 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+168))
	v1324 = *(*int64)(unsafe.Add(mBase, uint32(l7)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+168)) = v1323 + v1324
	v1327 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+176))
	v1328 = *(*int64)(unsafe.Add(mBase, uint32(l7)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+176)) = v1327 + v1328
	v1331 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+184))
	v1332 = *(*int64)(unsafe.Add(mBase, uint32(l7)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+184)) = v1331 + v1332
	v1335 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+192))
	v1336 = *(*int64)(unsafe.Add(mBase, uint32(l7)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+192)) = v1335 + v1336
	v1339 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+200))
	v1340 = *(*int64)(unsafe.Add(mBase, uint32(l7)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+200)) = v1339 + v1340
	v1343 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+208))
	v1344 = *(*int64)(unsafe.Add(mBase, uint32(l7)+80))
	v1346 = float64(1e+06)
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+208)) = base.F64_add(v1343, base.F64_div(base.F64_convert_i64_s(v1344), v1346))
	v1350 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+216))
	v1351 = *(*int64)(unsafe.Add(mBase, uint32(l7)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+216)) = base.F64_add(v1350, base.F64_div(base.F64_convert_i64_s(v1351), v1346))
	v1357 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+224))
	v1358 = *(*int64)(unsafe.Add(mBase, uint32(l7)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+224)) = base.F64_add(v1357, base.F64_div(base.F64_convert_i64_s(v1358), v1346))
	v1364 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+232))
	v1365 = *(*int64)(unsafe.Add(mBase, uint32(l7)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+232)) = base.F64_add(v1364, base.F64_div(base.F64_convert_i64_s(v1365), v1346))
	v1371 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+240))
	v1372 = *(*int64)(unsafe.Add(mBase, uint32(l7)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+240)) = base.F64_add(v1371, base.F64_div(base.F64_convert_i64_s(v1372), v1346))
	v1378 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+248))
	v1379 = *(*int64)(unsafe.Add(mBase, uint32(l7)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+248)) = base.F64_add(v1378, base.F64_div(base.F64_convert_i64_s(v1379), v1346))
	v1385 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+256))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+256)) = base.F64_add(v1385, float64(1))
	v1389 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+264))
	v1390 = *(*int64)(unsafe.Add(mBase, uint32(l8)))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+264)) = v1389 + v1390
	v1393 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+272))
	v1394 = *(*int64)(unsafe.Add(mBase, uint32(l8)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+272)) = v1393 + v1394
	v1397 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+280))
	v1398 = *(*int64)(unsafe.Add(mBase, uint32(l8)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+280)) = v1397 + v1398
	v1401 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+288))
	v1402 = *(*int64)(unsafe.Add(mBase, uint32(l8)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+288)) = v1401 + v1402
	if l9 != 0 {
		goto L240
	} else {
		goto L241
	}
L230:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1242)+88)) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1242)+72)) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1242)+56)) = l5
	goto L229
L231:
	;
	goto L232
L232:
	;
	v1254 = v1242 + int32(88)
	v1255 = *(*float64)(unsafe.Add(mBase, uint32(v1254)))
	v1256 = base.F64_sub(l5, v1255)
	v1259 = base.F64_add(v1255, base.F64_div(v1256, base.F64_convert_i64_s(v1240)))
	*(*float64)(unsafe.Add(mBase, uint32(v1254))) = v1259
	v1262 = v1242 + int32(104)
	v1265 = *(*float64)(unsafe.Add(mBase, uint32(v1262)))
	*(*float64)(unsafe.Add(mBase, uint32(v1262))) = base.F64_add(base.F64_mul(v1256, base.F64_sub(l5, v1259)), v1265)
	v1269 = v1242 + int32(56)
	v1270 = *(*float64)(unsafe.Add(mBase, uint32(v1269)))
	if base.F64_ne(v1270, float64(0)) != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if base.F64_lt(l5, v1270) != 0 {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	v1277 = v1192 + l4<<(uint(int32(3))%32) + int32(72)
	v1278 = *(*float64)(unsafe.Add(mBase, uint32(v1277)))
	if base.F64_ne(v1278, float64(0)) != 0 {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1269))) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1277))) = l5
	goto L229
L236:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1269))) = l5
	goto L238
L237:
	;
	goto L238
L238:
	;
	v1290 = v1192 + l4<<(uint(int32(3))%32) + int32(72)
	v1291 = *(*float64)(unsafe.Add(mBase, uint32(v1290)))
	if base.F64_lt(v1291, l5) == int32(0) {
		goto L229
	} else {
		goto L239
	}
L239:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1290))) = l5
	goto L229
L240:
	;
	v1405 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+296))
	v1406 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l9))))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+296)) = v1405 + v1406
	v1409 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+304))
	v1410 = *(*int64)(unsafe.Add(mBase, uint32(l9)+8))
	v1412 = float64(1e+06)
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+304)) = base.F64_add(v1409, base.F64_div(base.F64_convert_i64_s(v1410), v1412))
	v1416 = *(*int64)(unsafe.Add(mBase, uint32(l9)+16))
	v1419 = base.F64_div(base.F64_convert_i64_s(v1416), v1412)
	if base.F64_ne(v1419, float64(0)) != 0 {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	goto L242
L242:
	;
	v1469 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+376))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+376)) = v1469 + base.I64_extend_i32_s(l11)
	v1473 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+384))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+384)) = v1473 + base.I64_extend_i32_s(l12)
	switch l13 - int32(3) {
	case 0:
		v1481 = int32(392)
		goto L256
	case 1:
		goto L257
	default:
		goto L255
	}
L243:
	;
	v1422 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+328)) = v1422 + int64(1)
	goto L245
L244:
	;
	goto L245
L245:
	;
	v1426 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+320))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+320)) = base.F64_add(v1419, v1426)
	v1429 = *(*int64)(unsafe.Add(mBase, uint32(l9)+24))
	v1432 = base.F64_div(base.F64_convert_i64_s(v1429), float64(1e+06))
	if base.F64_ne(v1432, float64(0)) != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1435 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+312)) = v1435 + int64(1)
	goto L248
L247:
	;
	goto L248
L248:
	;
	v1439 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+336))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+336)) = base.F64_add(v1432, v1439)
	v1442 = *(*int64)(unsafe.Add(mBase, uint32(l9)+32))
	v1445 = base.F64_div(base.F64_convert_i64_s(v1442), float64(1e+06))
	if base.F64_ne(v1445, float64(0)) != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1448 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+344)) = v1448 + int64(1)
	goto L251
L250:
	;
	goto L251
L251:
	;
	v1452 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+352))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+352)) = base.F64_add(v1445, v1452)
	v1455 = *(*int64)(unsafe.Add(mBase, uint32(l9)+40))
	v1458 = base.F64_div(base.F64_convert_i64_s(v1455), float64(1e+06))
	if base.F64_ne(v1458, float64(0)) != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1461 = *(*int64)(unsafe.Add(mBase, uint32(v1192)+360))
	*(*int64)(unsafe.Add(mBase, uint32(v1192)+360)) = v1461 + int64(1)
	goto L254
L253:
	;
	goto L254
L254:
	;
	v1465 = *(*float64)(unsafe.Add(mBase, uint32(v1192)+368))
	*(*float64)(unsafe.Add(mBase, uint32(v1192)+368)) = base.F64_add(v1458, v1465)
	goto L242
L255:
	;
	v1489 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1219))), uint32(v1489))
	v1494 = v1191
	v1506 = v32
	goto L27
L256:
	;
	v1482 = v1192 + v1481
	v1483 = *(*int64)(unsafe.Add(mBase, uint32(v1482)))
	*(*int64)(unsafe.Add(mBase, uint32(v1482))) = v1483 + int64(1)
	goto L255
L257:
	;
	v1481 = int32(400)
	goto L256
L258:
	;
	if v1494 == int32(0) {
		v1543 = v1506
		goto L2
	} else {
		goto L259
	}
L259:
	;
	F_pfree(m, v1494)
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L25
	} else {
		goto L260
	}
L260:
	;
	v1543 = v1506
	goto L2
}
