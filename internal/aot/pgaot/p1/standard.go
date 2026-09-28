package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_standard_ExecutorStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v501 int32
	_ = v501
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
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v630 int32
	_ = v630
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v784 int32
	_ = v784
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v998 int32
	_ = v998
	var v1004 int32
	_ = v1004
	var v1012 int32
	_ = v1012
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1095 int32
	_ = v1095
	var v1109 int32
	_ = v1109
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1143 int32
	_ = v1143
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1177 int32
	_ = v1177
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1278 int32
	_ = v1278
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1372 int32
	_ = v1372
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1414 int32
	_ = v1414
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1477 int32
	_ = v1477
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1647 int32
	_ = v1647
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1668 int32
	_ = v1668
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	v3 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(128)
	m.G0 = v33
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[0])))
	if v36 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v193 = F_CreateExecutorState(m)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L20
	} else {
		goto L35
	}
L2:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
	if v53 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[1]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v42 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	if l1&int32(1) != 0 {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	if l1&int32(1) != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	v45 = int32(1)
	goto L9
L8:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+76)))
	v45 = v44
	goto L9
L9:
	;
	goto L6
L10:
	;
	if v45&int32(1) != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L1
L12:
	;
	goto L2
L13:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v150 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v56 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v64 = v3
	goto L16
L16:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v64<<(uint(int32(2))%32))))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+16))
	if v94&int64(-3) == int64(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L13
L18:
	;
	v117 = v64 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v117 < v118 {
		v64 = v117
		goto L16
	} else {
		goto L27
	}
L19:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v100 = F_get_rel_namespace(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[2]))
	goto L22
L22:
	;
	if base.B2i32(v104 != int32(0))&base.B2i32(v100 == v104) != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v109 = F_CreateCommandTag(m, v52)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109<<(uint(int32(3))%32))+uint32(_c_F_standard_ExecutorStart[3])))
	goto L25
L25:
	;
	F_PreventCommandIfReadOnly(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	goto L17
L28:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+29)))
	if v153 != int32(1) {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v156 = F_CreateCommandTag(m, v52)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L20
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156<<(uint(int32(3))%32))+uint32(_c_F_standard_ExecutorStart[3])))
	goto L33
L33:
	;
	F_PreventCommandIfParallelMode(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	goto L1
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v193
	v196 = int32(_a_F_standard_ExecutorStart_0)
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v193)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4])) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+88)) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+96))
	if v204 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v207 = F_palloc0_mul(m, int32(24), v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L20
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+56)) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+96)) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(4)) <= base.Ui32(v214-int32(2)) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+92)) = v207
	goto L38
L40:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v258 = F_RegisterSnapshot(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L20
	} else {
		goto L57
	}
L41:
	;
	v245 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L20
	} else {
		goto L53
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L20
	} else {
		goto L50
	}
L43:
	;
	if v214 != int32(1) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v227 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L20
	} else {
		goto L49
	}
L46:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+80))
	if v222 != 0 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
	if v223 != 0 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v255 = l1 | int32(32)
	goto L40
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+64)) = v227
	v255 = l1
	goto L40
L50:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v234
	F_errmsg_internal(m, int32(_a_F_standard_ExecutorStart_1), v33)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(240), int32(_a_F_standard_ExecutorStart_3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+64)) = v245
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+29)))
	if v251&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v254 = l1
	goto L56
L55:
	;
	v254 = l1 | int32(32)
	goto L56
L56:
	;
	v255 = v254
	goto L40
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+8)) = v258
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v262 = F_RegisterSnapshot(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+128)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v193)+12)) = v262
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+132)) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+176)) = v269
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v271 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v273 = F_palloc0(m, int32(360))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L20
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v255&int32(33) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v275 = int32(1)
	v276 = v271 & v275
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v276)
	v281 = int32(base.Ui32(v271)>>(uint(int32(3))%32)) & v275
	*(*uint8)(unsafe.Add(mBase, uint32(v273)+2)) = uint8(v281)
	v286 = int32(base.Ui32(v271)>>(uint(v275)%32)) & v275
	*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)) = uint8(v286)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v273
	goto L61
L63:
	;
	v294 = int32(_a_F_standard_ExecutorStart_4)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[5])) = v296 + int32(1)
	goto L66
L64:
	;
	goto L65
L65:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+40))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+48))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301)+56))
	v307 = F_ExecCheckPermissions(m, v304, v305, int32(1))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L20
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v301)+56))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v301)+52))
	v311 = F_bms_copy(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	F_ExecInitRangeTable(m, v303, v304, v309, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L20
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+36)) = v301
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v301)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v303)+40)) = v316
	v318 = m.G0
	v320 = v318 - int32(16)
	m.G0 = v320
	if v316 == int32(0) {
		v1223 = l0
		v1224 = v255
		v1228 = v303
		v1231 = v33
		v1233 = v301
		v1235 = v320
		v1246 = v302
		v1248 = v197
		v1249 = v300
		goto L70
	} else {
		goto L71
	}
L70:
	;
	m.G0 = v1235 + int32(16)
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+80))
	if v1256 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L71:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	if v325 <= int32(0) {
		v1223 = l0
		v1224 = v255
		v1228 = v303
		v1231 = v33
		v1233 = v301
		v1235 = v320
		v1246 = v302
		v1248 = v197
		v1249 = v300
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v328 = l0
	v329 = v255
	v333 = v303
	v336 = v33
	v338 = v301
	v340 = v320
	v346 = v316
	v351 = v302
	v352 = v3
	v353 = v197
	v354 = v300
	goto L73
L73:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358+v352<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+12)) = int32(0)
	v365 = F_CreateExprContext(m, v333)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L20
	} else {
		goto L75
	}
L74:
	;
	v1223 = v328
	v1224 = v329
	v1228 = v333
	v1231 = v336
	v1233 = v338
	v1235 = v340
	v1246 = v351
	v1248 = v353
	v1249 = v354
	goto L70
L75:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v333)+76))
	if v367 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v333)+100))
	v372 = F_CreatePartitionDirectory(m, v370, int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L20
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v375 = int32(0)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	if v377 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+76)) = v372
	goto L78
L80:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	v379 = v378
	goto L82
L81:
	;
	v379 = v375
	goto L82
L82:
	;
	v384 = F_palloc(m, v379<<(uint(int32(2))%32)+int32(24))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L20
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v365
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	v390 = F_bms_copy(m, v389)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L20
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+20)) = v379
	v393 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v384)+16)) = uint16(v393)
	*(*int32)(unsafe.Add(mBase, uint32(v384)+8)) = v390
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4]))
	v402 = F_AllocSetContextCreateInternal(m, v397, int32(_a_F_standard_ExecutorStart_5), v393, int32(_a_F_standard_ExecutorStart_6), int32(_a_F_standard_ExecutorStart_7))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L20
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+12)) = v402
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	if v405 == int32(0) {
		v1177 = v375
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v333)+44))
	v1197 = F_lappend(m, v1196, v384)
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L20
	} else {
		goto L195
	}
L87:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v408 <= int32(0) {
		v1177 = v375
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v425 = v375
	v434 = int32(0)
	goto L89
L89:
	;
	v446 = v434 << (uint(int32(2)) % 32)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v446+v447)))
	if v449 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v1177 = v1143
	goto L86
L91:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	v451 = v450
	goto L93
L92:
	;
	v451 = int32(0)
	goto L93
L93:
	;
	v457 = F_palloc(m, v451*int32(120)|int32(4))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L20
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v446+(v384+int32(24))))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v451
	if v449 == int32(0) {
		v1143 = v425
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v1163 = v434 + int32(1)
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v1163 < v1164 {
		v425 = v1143
		v434 = v1163
		goto L89
	} else {
		goto L194
	}
L96:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v463 <= int32(0) {
		v1143 = v425
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v480 = v425
	v485 = int32(0)
	goto L98
L98:
	;
	v501 = v457 + int32(4) + v485*int32(120)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v502+v485<<(uint(int32(2))%32))))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	v508 = F_ExecGetRangeTableRelation(m, v333, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L20
	} else {
		goto L100
	}
L99:
	;
	v1143 = v1109
	goto L95
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501))) = v508
	v511 = F_RelationGetPartitionKey(m, v508)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L20
	} else {
		goto L101
	}
L101:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v333)+76))
	v514 = F_PartitionDirectoryLookup(m, v513, v508)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L20
	} else {
		goto L102
	}
L102:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+4)) = v516
	v519 = F_palloc_mul(m, int32(4), v516)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L20
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+8)) = v519
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	if v522 != v523 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v506)+8))
	v904 = F_bms_copy(m, v903)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L20
	} else {
		goto L151
	}
L105:
	;
	v603 = F_palloc_mul(m, int32(4), v522)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L20
	} else {
		goto L127
	}
L106:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v506)+28))
	v528 = v522 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v528) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	if v590 != 0 {
		goto L105
	} else {
		goto L125
	}
L108:
	;
	v590 = int32(0)
	goto L107
L109:
	;
	v564 = v559
	v565 = v560
	v566 = v561
	goto L119
L110:
	;
	if (v525|v526)&int32(3) != 0 {
		v559 = v525
		v560 = v526
		v561 = v528
		goto L109
	} else {
		goto L113
	}
L111:
	;
	v552 = v525
	v553 = v526
	v554 = v528
	goto L112
L112:
	;
	if v554 == int32(0) {
		goto L108
	} else {
		goto L118
	}
L113:
	;
	v536 = v525
	v537 = v526
	v538 = v528
	goto L114
L114:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	if v541 != v542 {
		v559 = v536
		v560 = v537
		v561 = v538
		goto L109
	} else {
		goto L116
	}
L115:
	;
	v552 = v547
	v553 = v545
	v554 = v549
	goto L112
L116:
	;
	v544 = int32(4)
	v545 = v537 + v544
	v547 = v536 + v544
	v549 = v538 - v544
	if base.Ui32(int32(3)) < base.Ui32(v549) {
		v536 = v547
		v537 = v545
		v538 = v549
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v559 = v552
	v560 = v553
	v561 = v554
	goto L109
L119:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	if v569 == v570 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v590 = v569 - v570
	goto L107
L121:
	;
	v572 = int32(1)
	v577 = v566 - v572
	if v577 != 0 {
		v564 = v564 + v572
		v565 = v565 + v572
		v566 = v577
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	goto L120
L124:
	;
	goto L108
L125:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v506)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+12)) = v591
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v506)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+16)) = v593
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v597 = v595 << (uint(int32(2)) % 32)
	if v597 == int32(0) {
		goto L104
	} else {
		goto L126
	}
L126:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v506)+16))
	base.MemoryCopy(m, v519, v600, v597)
	goto L104
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+12)) = v603
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	v608 = F_palloc_mul(m, int32(4), v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L20
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+16)) = v608
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	if v611 <= int32(0) {
		goto L104
	} else {
		goto L129
	}
L129:
	;
	v614 = int32(0)
	v618 = v614
	v630 = v614
	goto L130
L130:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	if v646 <= v618 {
		v716 = v618
		goto L132
	} else {
		goto L133
	}
L131:
	;
	goto L104
L132:
	;
	v720 = v716
	goto L139
L133:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v506)+28))
	v651 = v618
	goto L134
L134:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v648+v651<<(uint(int32(2))%32))))
	if v682 != 0 {
		v716 = v651
		goto L132
	} else {
		goto L136
	}
L135:
	;
	v716 = v646
	goto L132
L136:
	;
	v684 = v651 + int32(1)
	if v684 != v646 {
		v651 = v684
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v870 = v630 + int32(1)
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	if v870 < v871 {
		v618 = v841
		v630 = v870
		goto L130
	} else {
		goto L150
	}
L139:
	;
	if v646 <= v720 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v826 = v630 << (uint(int32(2)) % 32)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v501)+12))
	v829 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v826+v827))) = v829
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v501)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v831+v826))) = v829
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v501)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v835+v826))) = int32(0)
	v841 = v720
	goto L138
L141:
	;
	v784 = v720
	goto L144
L142:
	;
	v748 = int32(2)
	v749 = v720 << (uint(v748) % 32)
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v506)+28))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v749+v750)))
	v754 = v630 << (uint(v748) % 32)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v754+v755)))
	if v752 != v757 {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v501)+8))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v506)+16))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v761+v749)))
	*(*int32)(unsafe.Add(mBase, uint32(v759+v754))) = v763
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v501)+12))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v506)+20))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v767+v749)))
	*(*int32)(unsafe.Add(mBase, uint32(v765+v754))) = v769
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v501)+16))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v506)+24))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v773+v749)))
	*(*int32)(unsafe.Add(mBase, uint32(v771+v754))) = v775
	v841 = v720 + int32(1)
	goto L138
L144:
	;
	v812 = v784 + int32(1)
	if v812 < v646 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L140
L146:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v506)+28))
	v815 = int32(2)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v814+v812<<(uint(v815)%32))))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v819+v630<<(uint(v815)%32))))
	if v818 != v823 {
		v784 = v812
		goto L144
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	goto L145
L149:
	;
	v720 = v812
	goto L139
L150:
	;
	goto L131
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+20)) = v904
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v506)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+24)) = v907
	if v907 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v506)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+28)) = v922
	if v922 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v365)+76))
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911)+128)))
	if v912&int32(2) != 0 {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	F_InitPartitionPruneContext(m, v501+int32(32), v907, v514, v511, int32(0), v365)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L20
	} else {
		goto L155
	}
L155:
	;
	v920 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v384)+16)) = uint8(v920)
	goto L152
L156:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v506)+40))
	v934 = F_bms_add_members(m, v932, v933)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L20
	} else {
		goto L159
	}
L157:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v365)+76))
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926)+128)))
	if v927&int32(2) != 0 {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v930 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v384)+17)) = uint8(v930)
	goto L156
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+4)) = v934
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v506)+32))
	if v937 == int32(0) {
		v1109 = v480
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v1129 = v485 + int32(1)
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v1129 < v1130 {
		v480 = v1109
		v485 = v1129
		goto L98
	} else {
		goto L193
	}
L161:
	;
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+16)))
	if v940 != 0 {
		v1109 = v480
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v501)+20))
	if v941 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	if v998 < int32(0) {
		v1109 = v480
		goto L160
	} else {
		goto L174
	}
L164:
	;
	v998 = base.I32_ctz(v984) | v985<<(uint(int32(5))%32)
	goto L163
L165:
	;
	v998 = int32(-2)
	goto L163
L166:
	;
	v949 = int32(0)
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v941)+4))
	if v952 <= v949 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v955 = v941 + int32(8)
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v955)))
	v962 = v959 & int32(-1)
	if v962 != 0 {
		v984 = v962
		v985 = v949
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v963 = int32(1)
	if v963 == v952 {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v967 = v963
	goto L170
L170:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v955+v967<<(uint(int32(2))%32))))
	if v974 != 0 {
		v984 = v974
		v985 = v967
		goto L164
	} else {
		goto L172
	}
L171:
	;
	goto L165
L172:
	;
	v976 = v967 + int32(1)
	if v976 != v952 {
		v967 = v976
		goto L170
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
	;
	v1004 = v998
	v1012 = v480
	goto L175
L175:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v501)+16))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1031+v1004<<(uint(int32(2))%32))))
	if v1035 != 0 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v1109 = v1038
	goto L160
L177:
	;
	v1036 = F_bms_add_member(m, v1012, v1035)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L20
	} else {
		goto L180
	}
L178:
	;
	v1038 = v1012
	goto L179
L179:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v501)+20))
	if v1039 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v1038 = v1036
	goto L179
L181:
	;
	if int32(0) <= v1095 {
		v1004 = v1095
		v1012 = v1038
		goto L175
	} else {
		goto L192
	}
L182:
	;
	v1095 = base.I32_ctz(v1081) | v1082<<(uint(int32(5))%32)
	goto L181
L183:
	;
	v1095 = int32(-2)
	goto L181
L184:
	;
	v1046 = v1004 + int32(1)
	v1048 = int32(base.Ui32(v1046) >> (uint(int32(5)) % 32))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1039)+4))
	if v1049 <= v1048 {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v1052 = v1039 + int32(8)
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1052+v1048<<(uint(int32(2))%32))))
	v1059 = v1056 & (int32(-1) << (uint(v1046) % 32))
	if v1059 != 0 {
		v1081 = v1059
		v1082 = v1048
		goto L182
	} else {
		goto L186
	}
L186:
	;
	v1061 = v1048 + int32(1)
	if v1061 == v1049 {
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v1064 = v1061
	goto L188
L188:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1052+v1064<<(uint(int32(2))%32))))
	if v1071 != 0 {
		v1081 = v1071
		v1082 = v1064
		goto L182
	} else {
		goto L190
	}
L189:
	;
	goto L183
L190:
	;
	v1073 = v1064 + int32(1)
	if v1073 != v1049 {
		v1064 = v1073
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	goto L176
L193:
	;
	goto L99
L194:
	;
	goto L90
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+44)) = v1197
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+16)))
	if v1200 != 0 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v333)+52))
	v1212 = F_bms_add_members(m, v1211, v1210)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L20
	} else {
		goto L201
	}
L197:
	;
	v1204 = F_ExecFindMatchingSubPlans(m, v384, int32(1), v340+int32(12))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L20
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340)+12)) = v1177
	v1209 = int32(0)
	v1210 = v1177
	goto L196
L200:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	v1209 = v1204
	v1210 = v1206
	goto L196
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+52)) = v1212
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v333)+48))
	v1216 = F_lappend(m, v1215, v1209)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L20
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+48)) = v1216
	v1220 = v352 + int32(1)
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v1220 < v1221 {
		v352 = v1220
		goto L73
	} else {
		goto L203
	}
L203:
	;
	goto L74
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1223)+48)) = v1619
	*(*int32)(unsafe.Add(mBase, uint32(v1223)+40)) = v1726
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4])) = v1248
	m.G0 = v1231 + int32(128)
	return
L205:
	;
	v1677 = int32(0)
	goto L283
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L20
	} else {
		goto L279
	}
L207:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L20
	} else {
		goto L275
	}
L208:
	;
	v1522 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1228)+156)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v1228)+104)) = v1522
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+68))
	if v1527 == v1522 {
		goto L259
	} else {
		goto L260
	}
L209:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1228)+20))
	v1261 = F_palloc0_mul(m, int32(4), v1260)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L20
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1228)+28)) = v1261
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+80))
	if v1264 == int32(0) {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+4))
	if v1267 <= int32(0) {
		goto L208
	} else {
		goto L212
	}
L212:
	;
	v1278 = int32(0)
	goto L213
L213:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+12))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1301+v1278<<(uint(int32(2))%32))))
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305)+32)))
	if v1306 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	goto L208
L215:
	;
	v1489 = v1278 + int32(1)
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+4))
	if v1489 < v1490 {
		v1278 = v1489
		goto L213
	} else {
		goto L258
	}
L216:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1228)+16))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+12))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+4))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1308+v1309<<(uint(int32(2))%32)-int32(4))))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+12))
	if v1316 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1228)+52))
	v1320 = F_bms_is_member(m, v1309, v1319)
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L20
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+16))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+16))
	if base.Ui32(int32(5)) <= base.Ui32(v1325) {
		goto L223
	} else {
		goto L224
	}
L220:
	;
	if v1320 == int32(0) {
		goto L215
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v1453 = F_palloc(m, int32(44))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L20
	} else {
		goto L257
	}
L223:
	;
	if v1325 == int32(5) {
		v1450 = int32(0)
		goto L222
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+4))
	v1349 = F_ExecGetRangeTableRelation(m, v1228, v1348)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L20
	} else {
		goto L230
	}
L226:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L20
	} else {
		goto L227
	}
L227:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+16)) = v1335
	F_errmsg_internal(m, int32(_a_F_standard_ExecutorStart_8), v1231+int32(16))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L20
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(929), int32(_a_F_standard_ExecutorStart_9))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L20
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	if v1349 == int32(0) {
		v1450 = int32(0)
		goto L222
	} else {
		goto L231
	}
L231:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+48))
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1353)+119)))
	switch v1354 - int32(83) {
	case 0:
		goto L237
	default:
		goto L206
	case 19:
		goto L233
	case 26:
		goto L234
	case 29, 31:
		v1450 = v1349
		goto L222
	case 33:
		goto L236
	case 35:
		goto L235
	}
L232:
	;
	v1450 = v1349
	goto L222
L233:
	;
	v1445 = F_GetFdwRoutineForRelation(m, v1349, int32(0))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L20
	} else {
		goto L255
	}
L234:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+16))
	if v1420 == int32(4) {
		goto L232
	} else {
		goto L250
	}
L235:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L20
	} else {
		goto L246
	}
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L20
	} else {
		goto L242
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L20
	} else {
		goto L238
	}
L238:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L20
	} else {
		goto L239
	}
L239:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+48)) = v1364 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_10), v1231+int32(48))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L20
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1208), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L20
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L20
	} else {
		goto L243
	}
L243:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+64)) = v1385 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_12), v1231-int32(-64))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L20
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1215), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L20
	} else {
		goto L247
	}
L247:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+80)) = v1406 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_13), v1231+int32(80))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L20
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1222), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L20
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L20
	} else {
		goto L251
	}
L251:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L20
	} else {
		goto L252
	}
L252:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+96)) = v1430 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_14), v1231+int32(96))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L20
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1230), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L20
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+108))
	if v1447 == int32(0) {
		goto L207
	} else {
		goto L256
	}
L256:
	;
	goto L232
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+4)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v1453))) = v1450
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+8)) = v1457
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+12)) = v1459
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+16)) = v1461
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+20)) = v1463
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+24)) = v1465
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+28))
	v1468 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+40)) = v1468
	*(*uint16)(unsafe.Add(mBase, uint32(v1453)+38)) = uint16(v1468)
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+34)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1453)+32)) = uint8(v1468)
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+28)) = v1467
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1228)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1477+v1457<<(uint(int32(2))%32)-int32(4)))) = v1453
	goto L215
L258:
	;
	goto L214
L259:
	;
	v1619 = F_ExecInitNode(m, v1246, v1228, v1224)
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L20
	} else {
		goto L271
	}
L260:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+4))
	if v1531 <= int32(0) {
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v1535 = v1224 & int32(-29)
	v1541 = v1522
	v1542 = int32(1)
	goto L262
L262:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+12))
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1568+v1541<<(uint(int32(2))%32))))
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+76))
	v1574 = F_bms_is_member(m, v1542, v1573)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L20
	} else {
		goto L264
	}
L263:
	;
	goto L259
L264:
	;
	if v1574 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1576 = v1535 | int32(4)
	goto L267
L266:
	;
	v1576 = v1535
	goto L267
L267:
	;
	v1577 = F_ExecInitNode(m, v1572, v1228, v1576)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L20
	} else {
		goto L268
	}
L268:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1228)+144))
	v1580 = F_lappend(m, v1579, v1577)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L20
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1228)+144)) = v1580
	v1583 = int32(1)
	v1586 = v1541 + v1583
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+4))
	if v1586 < v1587 {
		v1541 = v1586
		v1542 = v1542 + v1583
		goto L262
	} else {
		goto L270
	}
L270:
	;
	goto L263
L271:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1619)+56))
	if v1249 != int32(1) {
		v1726 = v1621
		goto L204
	} else {
		goto L272
	}
L272:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+44))
	if v1624 == int32(0) {
		v1726 = v1621
		goto L204
	} else {
		goto L273
	}
L273:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+4))
	if v1627 <= int32(0) {
		v1726 = v1621
		goto L204
	} else {
		goto L274
	}
L274:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+12))
	goto L205
L275:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L20
	} else {
		goto L276
	}
L276:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+112)) = v1639 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_15), v1231+int32(112))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L20
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1239), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L20
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L20
	} else {
		goto L280
	}
L280:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+32)) = v1660 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_16), v1231+int32(32))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L20
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1245), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L20
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1630+v1677<<(uint(int32(2))%32))))
	v1708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1707)+26)))
	if v1708 != int32(1) {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	v1716 = F_ExecInitExtraTupleSlot(m, v1228, int32(0), int32(_a_F_standard_ExecutorStart_17))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L20
	} else {
		goto L289
	}
L285:
	;
	v1712 = v1677 + int32(1)
	if v1712 != v1627 {
		v1677 = v1712
		goto L283
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	goto L284
L288:
	;
	v1726 = v1621
	goto L204
L289:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1619)+4))
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+44))
	v1720 = F_ExecInitJunkFilter(m, v1719, v1716)
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L20
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1228)+60)) = v1720
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1720)+8))
	v1726 = v1723
	goto L204
}
func F_standard_qp_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v286 int64
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v308 int32
	_ = v308
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v317 int64
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int64
	_ = v322
	var v325 int64
	_ = v325
	var v330 int32
	_ = v330
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int64
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int64
	_ = v736
	var v737 int32
	_ = v737
	var v738 int64
	_ = v738
	var v739 int32
	_ = v739
	var v740 int64
	_ = v740
	var v741 int32
	_ = v741
	var v742 int64
	_ = v742
	var v743 int32
	_ = v743
	var v744 int64
	_ = v744
	var v748 int64
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int64
	_ = v753
	var v756 int64
	_ = v756
	var v761 int32
	_ = v761
	var v763 int64
	_ = v763
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int64
	_ = v782
	var v783 int32
	_ = v783
	var v784 int64
	_ = v784
	var v785 int32
	_ = v785
	var v786 int64
	_ = v786
	var v787 int32
	_ = v787
	var v788 int64
	_ = v788
	var v789 int32
	_ = v789
	var v790 int64
	_ = v790
	var v794 int64
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int64
	_ = v799
	var v802 int64
	_ = v802
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int64
	_ = v811
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int64
	_ = v830
	var v831 int32
	_ = v831
	var v832 int64
	_ = v832
	var v833 int32
	_ = v833
	var v834 int64
	_ = v834
	var v835 int32
	_ = v835
	var v836 int64
	_ = v836
	var v837 int32
	_ = v837
	var v838 int64
	_ = v838
	var v842 int64
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int64
	_ = v847
	var v850 int64
	_ = v850
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int64
	_ = v858
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int64
	_ = v877
	var v878 int32
	_ = v878
	var v879 int64
	_ = v879
	var v880 int32
	_ = v880
	var v881 int64
	_ = v881
	var v882 int32
	_ = v882
	var v883 int64
	_ = v883
	var v884 int32
	_ = v884
	var v885 int64
	_ = v885
	var v889 int64
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int64
	_ = v894
	var v897 int64
	_ = v897
	var v902 int32
	_ = v902
	var v920 int32
	_ = v920
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v980 int32
	_ = v980
	var v987 int32
	_ = v987
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1019 int32
	_ = v1019
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1113 int32
	_ = v1113
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1228 int32
	_ = v1228
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1409 int32
	_ = v1409
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v22 != 0 {
		goto L283
	} else {
		goto L284
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	if v94 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v31 = v29
	goto L7
L6:
	;
	v31 = int32(0)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v31
	if v31 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v77 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v77 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v41 <= int32(0) {
		v69 = int32(1)
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v77 = v69
	goto L8
L13:
	;
	v44 = int32(0)
	if v44 < v41 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = v41
	goto L16
L15:
	;
	v47 = v44
	goto L16
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v50 = int32(0)
	goto L17
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48+v50<<(uint(int32(2))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v60 = int32(0)
	v61 = base.B2i32(v59 != v60)
	if v59 == v60 {
		v69 = v61
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v69 = v61
	goto L12
L19:
	;
	v65 = v50 + int32(1)
	if v65 != v47 {
		v50 = v65
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v80 = int32(0)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+45)))
	v85 = F_make_pathkeys_for_sortclauses_extended(m, l0, v20+int32(12), v23, v80, v81, v20+int32(11), v80)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+180)) = int64(0)
	goto L1
L24:
	;
	return
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v85
	if v85 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v90 = v88
	goto L28
L27:
	;
	v90 = int32(0)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v90
	goto L1
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+180)) = int64(0)
	goto L1
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	if v97 <= int32(0) {
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v102 = int32(1)
	v107 = F_make_pathkeys_for_sortclauses_extended(m, l0, l0+int32(276), v23, v102, int32(0), v20+int32(10), v102)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L24
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v107
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)))
	if v110 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+180)) = int64(0)
	goto L1
L36:
	;
	goto L37
L37:
	;
	if v107 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v117 = v115
	goto L40
L39:
	;
	v117 = int32(0)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	if v119 <= int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_qp_callback[0])))
	if v123&int32(1) == int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v128 == int32(0) {
		v276 = v3
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v286 = int64(0)
	if v276 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v131 <= int32(0) {
		v276 = v3
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v137 = v131
	v138 = v3
	v142 = v3
	goto L46
L46:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v138<<(uint(int32(2))%32))))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+50)))
	if v159 != int32(110) {
		v251 = v137
		v256 = v142
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v276 = v256
	goto L43
L48:
	;
	v266 = v138 + int32(1)
	if v266 < v251 {
		v137 = v251
		v138 = v266
		v142 = v256
		goto L46
	} else {
		goto L66
	}
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+40))
	if v162 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v158)+36))
	if v165 == int32(0) {
		v251 = v137
		v256 = v142
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v158)+44))
	if v168 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v245 = F_bms_add_member(m, v142, v138)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L24
	} else {
		goto L65
	}
L55:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v158)+32))
	if v171 == int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v174 <= int32(0) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v181 = int32(0)
	goto L58
L58:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v177+v181<<(uint(int32(2))%32))))
	v206 = v199
	goto L60
L59:
	;
	goto L54
L60:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v218 == int32(27) {
		v206 = v217
		goto L60
	} else {
		goto L62
	}
L61:
	;
	if base.Ui32(int32(1)) < base.Ui32(v218-int32(6)) {
		v251 = v137
		v256 = v142
		goto L48
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v226 = v181 + int32(1)
	if v174 != v226 {
		v181 = v226
		goto L58
	} else {
		goto L64
	}
L64:
	;
	goto L59
L65:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v251 = v247
	v256 = v245
	goto L48
L66:
	;
	goto L47
L67:
	;
	if v920 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L68:
	;
	goto L84
L69:
	;
	v330 = int32(0)
	goto L68
L70:
	;
	goto L71
L71:
	;
	v291 = v276 + int32(8)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v292 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v330 = base.I32_popcnt(v295)
	goto L68
L73:
	;
	goto L74
L74:
	;
	v298 = v292 << (uint(int32(2)) % 32)
	if v298 <= int32(7) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v330 = base.I32_wrap_i64(v325)
	goto L68
L76:
	;
	if v298 == int32(0) {
		v325 = v286
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v322 = F_pg_popcount_optimized(m, v291, v298)
	mBase = m.M
	v325 = v322
	goto L75
L79:
	;
	v303 = v298
	v304 = v291
	v305 = v286
	goto L80
L80:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+3)))
	v307 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v306)+uint32(_c_F_standard_qp_callback[1]))))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+2)))
	v309 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v308)+uint32(_c_F_standard_qp_callback[1]))))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+1)))
	v311 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_standard_qp_callback[1]))))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	v313 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v312)+uint32(_c_F_standard_qp_callback[1]))))
	v317 = v307 + (v309 + (v311 + (v305 + v313)))
	v318 = int32(4)
	v321 = v303 - v318
	if v321 != 0 {
		v303 = v321
		v304 = v304 + v318
		v305 = v317
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v325 = v317
	goto L75
L82:
	;
	goto L81
L83:
	;
	if v330 <= int32(0) {
		v920 = v3
		goto L67
	} else {
		goto L98
	}
L84:
	;
	goto L83
L98:
	;
	v388 = v276
	v391 = int32(0)
	v393 = v3
	goto L99
L99:
	;
	v397 = int32(0)
	if v388 == v397 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	if v809 == int32(0) {
		v920 = v856
		goto L67
	} else {
		goto L250
	}
L101:
	;
	if int32(0) <= v455 {
		goto L112
	} else {
		goto L113
	}
L102:
	;
	v455 = base.I32_ctz(v441) | v442<<(uint(int32(5))%32)
	goto L101
L103:
	;
	v455 = int32(-2)
	goto L101
L104:
	;
	v406 = int32(0)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	if v409 <= v406 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v412 = v388 + int32(8)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v419 = v416 & int32(-1)
	if v419 != 0 {
		v441 = v419
		v442 = v406
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v420 = int32(1)
	if v420 == v409 {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v424 = v420
	goto L108
L108:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v412+v424<<(uint(int32(2))%32))))
	if v431 != 0 {
		v441 = v431
		v442 = v424
		goto L102
	} else {
		goto L110
	}
L109:
	;
	goto L103
L110:
	;
	v433 = v424 + int32(1)
	if v433 != v409 {
		v424 = v433
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v460 = v397
	v462 = v397
	v463 = v455
	v466 = v388
	goto L115
L113:
	;
	v699 = v397
	v701 = v397
	v705 = v388
	goto L114
L114:
	;
	v714 = F_bms_del_members(m, v705, v701)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L24
	} else {
		goto L182
	}
L115:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v476+v463<<(uint(int32(2))%32))))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+12))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+40))
	if v484 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v699 = v624
	v701 = v626
	v705 = v630
	goto L114
L117:
	;
	if v630 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L118:
	;
	if v460 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L119:
	;
	v486 = v484
	goto L121
L120:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v483)+36))
	v486 = v485
	goto L121
L121:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v483)+32))
	v488 = F_make_pathkeys_for_sortclauses(m, l0, v486, v487)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L24
	} else {
		goto L122
	}
L122:
	;
	if v488 == int32(0) {
		goto L118
	} else {
		goto L123
	}
L123:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v492 <= int32(0) {
		goto L118
	} else {
		goto L124
	}
L124:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v503 = int32(0)
	goto L125
L125:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v495+v503<<(uint(int32(2))%32))))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+41)))
	if v519 != int32(1) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v525 = F_bms_del_member(m, v466, v463)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L24
	} else {
		goto L131
	}
L127:
	;
	v523 = v503 + int32(1)
	if v523 != v492 {
		v503 = v523
		goto L125
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	goto L126
L130:
	;
	goto L118
L131:
	;
	v624 = v460
	v626 = v462
	v630 = v525
	goto L117
L132:
	;
	v620 = F_bms_add_member(m, v462, v463)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L24
	} else {
		goto L169
	}
L133:
	;
	v618 = v617
	goto L132
L134:
	;
	if v107 == int32(0) {
		v617 = v488
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	if v107 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v548 = F_list_copy(m, v107)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L24
	} else {
		goto L138
	}
L138:
	;
	v550 = F_append_pathkeys(m, v548, v488)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L24
	} else {
		goto L139
	}
L139:
	;
	v618 = v550
	goto L132
L140:
	;
	v552 = F_list_copy(m, v107)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L24
	} else {
		goto L143
	}
L141:
	;
	v556 = v488
	goto L142
L142:
	;
	if v460 == v556 {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v554 = F_append_pathkeys(m, v552, v488)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L24
	} else {
		goto L144
	}
L144:
	;
	v556 = v554
	goto L142
L145:
	;
	switch v616 {
	case 0, 1:
		v618 = v460
		goto L132
	case 2:
		v617 = v556
		goto L133
	default:
		v624 = v460
		v626 = v462
		v630 = v466
		goto L117
	}
L146:
	;
	v616 = int32(0)
	goto L145
L147:
	;
	goto L148
L148:
	;
	v565 = int32(0)
	goto L151
L149:
	;
	if v605 != 0 {
		goto L166
	} else {
		goto L167
	}
L150:
	;
	v600 = int32(0)
	if v587 != 0 {
		goto L163
	} else {
		goto L164
	}
L151:
	;
	v569 = int32(0)
	if v460 == v569 {
		v579 = v569
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v616 = int32(3)
	goto L145
L153:
	;
	if v556 != 0 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	if v573 <= v565 {
		v579 = int32(0)
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v460)+12))
	v579 = v575 + v565<<(uint(int32(2))%32)
	goto L153
L156:
	;
	v585 = int32(0)
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v556)+12))
	if base.B2i32(v579 == v585)|base.B2i32(v587 == v585) != 0 {
		goto L150
	} else {
		goto L161
	}
L157:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	if v565 < v580 {
		goto L156
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v582 = int32(0)
	v605 = base.B2i32(v579 == v582)
	v607 = v582
	goto L149
L160:
	;
	goto L159
L161:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v587+v565<<(uint(int32(2))%32))))
	if v595 == v597 {
		v565 = v565 + int32(1)
		goto L151
	} else {
		goto L162
	}
L162:
	;
	goto L152
L163:
	;
	v604 = int32(2)
	goto L165
L164:
	;
	v604 = v600
	goto L165
L165:
	;
	v605 = base.B2i32(v579 == v600)
	v607 = v604
	goto L149
L166:
	;
	v609 = v607
	goto L168
L167:
	;
	v609 = int32(1)
	goto L168
L168:
	;
	v616 = v609
	goto L145
L169:
	;
	v624 = v618
	v626 = v620
	v630 = v466
	goto L117
L170:
	;
	if int32(0) <= v694 {
		v460 = v624
		v462 = v626
		v463 = v694
		v466 = v630
		goto L115
	} else {
		goto L181
	}
L171:
	;
	v694 = base.I32_ctz(v680) | v681<<(uint(int32(5))%32)
	goto L170
L172:
	;
	v694 = int32(-2)
	goto L170
L173:
	;
	v645 = v463 + int32(1)
	v647 = int32(base.Ui32(v645) >> (uint(int32(5)) % 32))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	if v648 <= v647 {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v651 = v630 + int32(8)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v651+v647<<(uint(int32(2))%32))))
	v658 = v655 & (int32(-1) << (uint(v645) % 32))
	if v658 != 0 {
		v680 = v658
		v681 = v647
		goto L171
	} else {
		goto L175
	}
L175:
	;
	v660 = v647 + int32(1)
	if v660 == v648 {
		goto L172
	} else {
		goto L176
	}
L176:
	;
	v663 = v660
	goto L177
L177:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v651+v663<<(uint(int32(2))%32))))
	if v670 != 0 {
		v680 = v670
		v681 = v663
		goto L171
	} else {
		goto L179
	}
L178:
	;
	goto L172
L179:
	;
	v672 = v663 + int32(1)
	if v672 != v648 {
		v663 = v672
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	goto L116
L182:
	;
	v717 = int64(0)
	if v701 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v763 = int64(0)
	if v393 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L184:
	;
	v761 = int32(0)
	goto L183
L185:
	;
	goto L186
L186:
	;
	v722 = v701 + int32(8)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	if v723 == int32(1) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	v761 = base.I32_popcnt(v726)
	goto L183
L188:
	;
	goto L189
L189:
	;
	v729 = v723 << (uint(int32(2)) % 32)
	if v729 <= int32(7) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v761 = base.I32_wrap_i64(v756)
	goto L183
L191:
	;
	if v729 == int32(0) {
		v756 = v717
		goto L190
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v753 = F_pg_popcount_optimized(m, v722, v729)
	mBase = m.M
	v756 = v753
	goto L190
L194:
	;
	v734 = v729
	v735 = v722
	v736 = v717
	goto L195
L195:
	;
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735)+3)))
	v738 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v737)+uint32(_c_F_standard_qp_callback[1]))))
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735)+2)))
	v740 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v739)+uint32(_c_F_standard_qp_callback[1]))))
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735)+1)))
	v742 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v741)+uint32(_c_F_standard_qp_callback[1]))))
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735))))
	v744 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v743)+uint32(_c_F_standard_qp_callback[1]))))
	v748 = v738 + (v740 + (v742 + (v736 + v744)))
	v749 = int32(4)
	v752 = v734 - v749
	if v752 != 0 {
		v734 = v752
		v735 = v735 + v749
		v736 = v748
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v756 = v748
	goto L190
L197:
	;
	goto L196
L198:
	;
	v808 = base.B2i32(v807 < v761)
	if v807 < v761 {
		goto L213
	} else {
		goto L214
	}
L199:
	;
	v807 = int32(0)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v768 = v393 + int32(8)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v769 == int32(1) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v768)))
	v807 = base.I32_popcnt(v772)
	goto L198
L203:
	;
	goto L204
L204:
	;
	v775 = v769 << (uint(int32(2)) % 32)
	if v775 <= int32(7) {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v807 = base.I32_wrap_i64(v802)
	goto L198
L206:
	;
	if v775 == int32(0) {
		v802 = v763
		goto L205
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v799 = F_pg_popcount_optimized(m, v768, v775)
	mBase = m.M
	v802 = v799
	goto L205
L209:
	;
	v780 = v775
	v781 = v768
	v782 = v763
	goto L210
L210:
	;
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781)+3)))
	v784 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v783)+uint32(_c_F_standard_qp_callback[1]))))
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781)+2)))
	v786 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v785)+uint32(_c_F_standard_qp_callback[1]))))
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781)+1)))
	v788 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v787)+uint32(_c_F_standard_qp_callback[1]))))
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781))))
	v790 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v789)+uint32(_c_F_standard_qp_callback[1]))))
	v794 = v784 + (v786 + (v788 + (v782 + v790)))
	v795 = int32(4)
	v798 = v780 - v795
	if v798 != 0 {
		v780 = v798
		v781 = v781 + v795
		v782 = v794
		goto L210
	} else {
		goto L212
	}
L211:
	;
	v802 = v794
	goto L205
L212:
	;
	goto L211
L213:
	;
	v809 = v699
	goto L215
L214:
	;
	v809 = v391
	goto L215
L215:
	;
	v811 = int64(0)
	if v714 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	if v807 < v761 {
		goto L231
	} else {
		goto L232
	}
L217:
	;
	v855 = int32(0)
	goto L216
L218:
	;
	goto L219
L219:
	;
	v816 = v714 + int32(8)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	if v817 == int32(1) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	v855 = base.I32_popcnt(v820)
	goto L216
L221:
	;
	goto L222
L222:
	;
	v823 = v817 << (uint(int32(2)) % 32)
	if v823 <= int32(7) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v855 = base.I32_wrap_i64(v850)
	goto L216
L224:
	;
	if v823 == int32(0) {
		v850 = v811
		goto L223
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v847 = F_pg_popcount_optimized(m, v816, v823)
	mBase = m.M
	v850 = v847
	goto L223
L227:
	;
	v828 = v823
	v829 = v816
	v830 = v811
	goto L228
L228:
	;
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829)+3)))
	v832 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v831)+uint32(_c_F_standard_qp_callback[1]))))
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829)+2)))
	v834 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v833)+uint32(_c_F_standard_qp_callback[1]))))
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829)+1)))
	v836 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v835)+uint32(_c_F_standard_qp_callback[1]))))
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829))))
	v838 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v837)+uint32(_c_F_standard_qp_callback[1]))))
	v842 = v832 + (v834 + (v836 + (v830 + v838)))
	v843 = int32(4)
	v846 = v828 - v843
	if v846 != 0 {
		v828 = v846
		v829 = v829 + v843
		v830 = v842
		goto L228
	} else {
		goto L230
	}
L229:
	;
	v850 = v842
	goto L223
L230:
	;
	goto L229
L231:
	;
	v856 = v701
	goto L233
L232:
	;
	v856 = v393
	goto L233
L233:
	;
	v858 = int64(0)
	if v856 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	if v902 < v855 {
		v388 = v714
		v391 = v809
		v393 = v856
		goto L99
	} else {
		goto L249
	}
L235:
	;
	v902 = int32(0)
	goto L234
L236:
	;
	goto L237
L237:
	;
	v863 = v856 + int32(8)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v856)+4))
	if v864 == int32(1) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	v902 = base.I32_popcnt(v867)
	goto L234
L239:
	;
	goto L240
L240:
	;
	v870 = v864 << (uint(int32(2)) % 32)
	if v870 <= int32(7) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v902 = base.I32_wrap_i64(v897)
	goto L234
L242:
	;
	if v870 == int32(0) {
		v897 = v858
		goto L241
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v894 = F_pg_popcount_optimized(m, v863, v870)
	mBase = m.M
	v897 = v894
	goto L241
L245:
	;
	v875 = v870
	v876 = v863
	v877 = v858
	goto L246
L246:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876)+3)))
	v879 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v878)+uint32(_c_F_standard_qp_callback[1]))))
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876)+2)))
	v881 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v880)+uint32(_c_F_standard_qp_callback[1]))))
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876)+1)))
	v883 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v882)+uint32(_c_F_standard_qp_callback[1]))))
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876))))
	v885 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v884)+uint32(_c_F_standard_qp_callback[1]))))
	v889 = v879 + (v881 + (v883 + (v877 + v885)))
	v890 = int32(4)
	v893 = v875 - v890
	if v893 != 0 {
		v875 = v893
		v876 = v876 + v890
		v877 = v889
		goto L246
	} else {
		goto L248
	}
L247:
	;
	v897 = v889
	goto L241
L248:
	;
	goto L247
L249:
	;
	goto L100
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v809
	v920 = v856
	goto L67
L251:
	;
	if v980 < int32(0) {
		goto L1
	} else {
		goto L262
	}
L252:
	;
	v980 = base.I32_ctz(v966) | v967<<(uint(int32(5))%32)
	goto L251
L253:
	;
	v980 = int32(-2)
	goto L251
L254:
	;
	v931 = int32(0)
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v920)+4))
	if v934 <= v931 {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v937 = v920 + int32(8)
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	v944 = v941 & int32(-1)
	if v944 != 0 {
		v966 = v944
		v967 = v931
		goto L252
	} else {
		goto L256
	}
L256:
	;
	v945 = int32(1)
	if v945 == v934 {
		goto L253
	} else {
		goto L257
	}
L257:
	;
	v949 = v945
	goto L258
L258:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v937+v949<<(uint(int32(2))%32))))
	if v956 != 0 {
		v966 = v956
		v967 = v949
		goto L252
	} else {
		goto L260
	}
L259:
	;
	goto L253
L260:
	;
	v958 = v949 + int32(1)
	if v958 != v934 {
		v949 = v958
		goto L258
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	v987 = v980
	goto L263
L263:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+12))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v1001+v987<<(uint(int32(2))%32))))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1005)+4))
	if v1006 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	goto L1
L265:
	;
	if v920 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L266:
	;
	v1009 = int32(0)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+4))
	if v1010 <= v1009 {
		goto L265
	} else {
		goto L267
	}
L267:
	;
	v1019 = v1009
	goto L268
L268:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+12))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1030+v1019<<(uint(int32(2))%32))))
	v1035 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1034)+51)) = uint8(v1035)
	v1038 = v1019 + v1035
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+4))
	if v1038 < v1039 {
		v1019 = v1038
		goto L268
	} else {
		goto L270
	}
L269:
	;
	goto L265
L270:
	;
	goto L269
L271:
	;
	if int32(0) <= v1113 {
		v987 = v1113
		goto L263
	} else {
		goto L282
	}
L272:
	;
	v1113 = base.I32_ctz(v1099) | v1100<<(uint(int32(5))%32)
	goto L271
L273:
	;
	v1113 = int32(-2)
	goto L271
L274:
	;
	v1064 = v987 + int32(1)
	v1066 = int32(base.Ui32(v1064) >> (uint(int32(5)) % 32))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v920)+4))
	if v1067 <= v1066 {
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v1070 = v920 + int32(8)
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1070+v1066<<(uint(int32(2))%32))))
	v1077 = v1074 & (int32(-1) << (uint(v1064) % 32))
	if v1077 != 0 {
		v1099 = v1077
		v1100 = v1066
		goto L272
	} else {
		goto L276
	}
L276:
	;
	v1079 = v1066 + int32(1)
	if v1079 == v1067 {
		goto L273
	} else {
		goto L277
	}
L277:
	;
	v1082 = v1079
	goto L278
L278:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1070+v1082<<(uint(int32(2))%32))))
	if v1089 != 0 {
		v1099 = v1089
		v1100 = v1082
		goto L272
	} else {
		goto L280
	}
L279:
	;
	goto L273
L280:
	;
	v1091 = v1082 + int32(1)
	if v1091 != v1067 {
		v1082 = v1091
		goto L278
	} else {
		goto L281
	}
L281:
	;
	goto L279
L282:
	;
	goto L264
L283:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1135)))
	v1137 = F_make_pathkeys_for_window(m, l0, v1136, v23)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L24
	} else {
		goto L286
	}
L284:
	;
	v1140 = int32(0)
	goto L285
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v1140
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v24)+120))
	if v1142 != 0 {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	v1140 = v1137
	goto L285
L287:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v24)+124))
	v1162 = F_make_pathkeys_for_sortclauses(m, l0, v1161, v23)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L24
	} else {
		goto L296
	}
L288:
	;
	v1143 = F_list_copy(m, v1142)
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L24
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = int32(0)
	goto L287
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v1143
	v1149 = int32(0)
	v1153 = F_make_pathkeys_for_sortclauses_extended(m, l0, l0+int32(280), v23, int32(1), v1149, v20+int32(9), v1149)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L24
	} else {
		goto L292
	}
L292:
	;
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)))
	if v1156 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1157 = v1153
	goto L295
L294:
	;
	v1157 = int32(0)
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v1157
	goto L287
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v1162
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1165 != 0 {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v1437
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v1439 != 0 {
		goto L368
	} else {
		goto L369
	}
L298:
	;
	v1166 = int32(0)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+32))
	v1169 = F_copyObjectImpl(m, v1168)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L24
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1437 = int32(0)
	goto L297
L301:
	;
	if v1169 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+12))
	v1172 = v1171
	goto L304
L303:
	;
	v1172 = v1166
	goto L304
L304:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+20))
	if v1173 != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+12))
	v1175 = v1174
	goto L307
L306:
	;
	v1175 = v1166
	goto L307
L307:
	;
	if v23 == int32(0) {
		v1391 = v1169
		goto L308
	} else {
		goto L309
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v1391
	v1409 = int32(0)
	v1414 = F_make_pathkeys_for_sortclauses_extended(m, l0, v20+int32(4), v23, v1409, v1409, v20+int32(3), v1409)
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L24
	} else {
		goto L363
	}
L309:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v1178 <= int32(0) {
		v1391 = v1169
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v1185 = v1178
	v1186 = v1175
	v1187 = v1172
	v1188 = int32(0)
	goto L311
L311:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1199+v1188<<(uint(int32(2))%32))))
	v1204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203)+26)))
	if v1204 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	v1391 = v1169
	goto L308
L313:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1187)))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+4))
	v1210 = F_exprType(m, v1209)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L24
	} else {
		goto L316
	}
L314:
	;
	v1379 = v1185
	v1380 = v1186
	v1381 = v1187
	goto L315
L315:
	;
	v1387 = v1188 + int32(1)
	if v1387 < v1379 {
		v1185 = v1379
		v1186 = v1380
		v1187 = v1381
		v1188 = v1387
		goto L311
	} else {
		goto L362
	}
L316:
	;
	if v1208 != v1210 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1391 = int32(0)
	goto L308
L318:
	;
	goto L319
L319:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+20))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+12))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+4))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+12))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+4))
	v1219 = int32(0)
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+16))
	if v1228 == v1219 {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1207)+4)) = v1352
	v1363 = v1187 + int32(4)
	if base.Ui32(v1363) < base.Ui32(v1217+v1218<<(uint(int32(2))%32)) {
		goto L356
	} else {
		goto L357
	}
L321:
	;
	if v23 == int32(0) {
		v1348 = int32(1)
		goto L324
	} else {
		goto L325
	}
L322:
	;
	v1352 = v1228
	goto L323
L323:
	;
	goto L320
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1203)+16)) = v1348
	v1352 = v1348
	goto L323
L325:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v1235 <= int32(0) {
		v1348 = int32(1)
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v1238 = int32(0)
	if v1238 < v1235 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1241 = v1235
	goto L329
L328:
	;
	v1241 = v1238
	goto L329
L329:
	;
	v1243 = v1241 & int32(3)
	v1244 = int32(0)
	if int32(4) <= v1235 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1348 = v1326 + int32(1)
	goto L324
L331:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1253 = v1244
	v1254 = int32(0)
	v1255 = v1219
	goto L334
L332:
	;
	v1290 = v1244
	v1292 = v1219
	goto L333
L333:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1301 = int32(0)
	v1303 = v1290
	v1305 = v1292
	goto L350
L334:
	;
	v1264 = v1249 + v1255<<(uint(int32(2))%32)
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+12))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+16))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+8))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1267)+16))
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+4))
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+16))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1264)))
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+16))
	if base.Ui32(v1253) < base.Ui32(v1272) {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	if v1243 == int32(0) {
		v1326 = v1280
		goto L330
	} else {
		goto L349
	}
L336:
	;
	v1274 = v1272
	goto L338
L337:
	;
	v1274 = v1253
	goto L338
L338:
	;
	if base.Ui32(v1274) < base.Ui32(v1270) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1276 = v1270
	goto L341
L340:
	;
	v1276 = v1274
	goto L341
L341:
	;
	if base.Ui32(v1276) < base.Ui32(v1268) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1278 = v1268
	goto L344
L343:
	;
	v1278 = v1276
	goto L344
L344:
	;
	if base.Ui32(v1278) < base.Ui32(v1266) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1280 = v1266
	goto L347
L346:
	;
	v1280 = v1278
	goto L347
L347:
	;
	v1281 = int32(4)
	v1282 = v1255 + v1281
	v1284 = v1254 + v1281
	if v1284 != v1241&int32(2147483644) {
		v1253 = v1280
		v1254 = v1284
		v1255 = v1282
		goto L334
	} else {
		goto L348
	}
L348:
	;
	goto L335
L349:
	;
	v1290 = v1280
	v1292 = v1282
	goto L333
L350:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1299+v1305<<(uint(int32(2))%32))))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+16))
	if base.Ui32(v1303) < base.Ui32(v1316) {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v1326 = v1318
	goto L330
L352:
	;
	v1318 = v1316
	goto L354
L353:
	;
	v1318 = v1303
	goto L354
L354:
	;
	v1319 = int32(1)
	v1322 = v1301 + v1319
	if v1322 != v1243 {
		v1301 = v1322
		v1303 = v1318
		v1305 = v1305 + v1319
		goto L350
	} else {
		goto L355
	}
L355:
	;
	goto L351
L356:
	;
	v1369 = v1363
	goto L358
L357:
	;
	v1369 = int32(0)
	goto L358
L358:
	;
	v1371 = v1186 + int32(4)
	if base.Ui32(v1371) < base.Ui32(v1215+v1216<<(uint(int32(2))%32)) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1377 = v1371
	goto L361
L360:
	;
	v1377 = int32(0)
	goto L361
L361:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v1379 = v1378
	v1380 = v1377
	v1381 = v1369
	goto L315
L362:
	;
	goto L312
L363:
	;
	v1417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
	if v1417 != 0 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1418 = v1414
	goto L366
L365:
	;
	v1418 = int32(0)
	goto L366
L366:
	;
	v1437 = v1418
	goto L297
L367:
	;
	m.G0 = v20 + int32(16)
	return
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1439
	goto L367
L369:
	;
	goto L370
L370:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v1441 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1441
	goto L367
L372:
	;
	goto L373
L373:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v1443 != 0 {
		goto L378
	} else {
		goto L379
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1455
	goto L367
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1443
	goto L367
L376:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+4))
	if v1456 <= v1457 {
		goto L374
	} else {
		goto L387
	}
L377:
	;
	if v1437 != 0 {
		goto L384
	} else {
		goto L385
	}
L378:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+4))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v1445 != 0 {
		v1455 = v1445
		v1456 = v1444
		goto L376
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v1449 != 0 {
		v1455 = v1449
		v1456 = int32(0)
		goto L376
	} else {
		goto L383
	}
L381:
	;
	if int32(0) < v1444 {
		goto L375
	} else {
		goto L382
	}
L382:
	;
	goto L377
L383:
	;
	goto L377
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1437
	goto L367
L385:
	;
	goto L386
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
	goto L367
L387:
	;
	goto L375
}
