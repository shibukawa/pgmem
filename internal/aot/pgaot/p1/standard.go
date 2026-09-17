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
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
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
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
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
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v613 int32
	_ = v613
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v770 int32
	_ = v770
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v984 int32
	_ = v984
	var v990 int32
	_ = v990
	var v999 int32
	_ = v999
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1081 int32
	_ = v1081
	var v1096 int32
	_ = v1096
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1130 int32
	_ = v1130
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1164 int32
	_ = v1164
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1265 int32
	_ = v1265
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1380 int32
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1401 int32
	_ = v1401
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1425 int32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1464 int32
	_ = v1464
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1509 int32
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1555 int32
	_ = v1555
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1634 int32
	_ = v1634
	var v1639 int32
	_ = v1639
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1655 int32
	_ = v1655
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
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
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
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
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+25)))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v193
	v196 = int32(_a_F_standard_ExecutorStart_0)
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v193)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4])) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+88)) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+84))
	if v204 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v208 = F_palloc0(m, v205*int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L20
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+56)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+96)) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(4)) <= base.Ui32(v215-int32(2)) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+92)) = v208
	goto L38
L40:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v259 = F_RegisterSnapshot(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L20
	} else {
		goto L57
	}
L41:
	;
	v246 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L20
	} else {
		goto L53
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L20
	} else {
		goto L50
	}
L43:
	;
	if v215 != int32(1) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v228 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L20
	} else {
		goto L49
	}
L46:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+72))
	if v223 != 0 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+25)))
	if v224 != 0 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v256 = l1 | int32(32)
	goto L40
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+64)) = v228
	v256 = l1
	goto L40
L50:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v235
	F_errmsg_internal(m, int32(_a_F_standard_ExecutorStart_1), v33)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(238), int32(_a_F_standard_ExecutorStart_3))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v193)+64)) = v246
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+25)))
	if v252&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v255 = l1
	goto L56
L55:
	;
	v255 = l1 | int32(32)
	goto L56
L56:
	;
	v256 = v255
	goto L40
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+8)) = v259
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v263 = F_RegisterSnapshot(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+128)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v193)+12)) = v263
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+132)) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+176)) = v270
	if v256&int32(33) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v276 = int32(_a_F_standard_ExecutorStart_4)
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[5])) = v278 + int32(1)
	goto L62
L60:
	;
	goto L61
L61:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+36))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v283)+44))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283)+52))
	v289 = F_ExecCheckPermissions(m, v286, v287, int32(1))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L20
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v283)+52))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v283)+48))
	v293 = F_bms_copy(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	F_ExecInitRangeTable(m, v285, v286, v291, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+36)) = v283
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v283)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v285)+40)) = v298
	v300 = m.G0
	v302 = v300 - int32(16)
	m.G0 = v302
	if v298 == int32(0) {
		v1209 = l0
		v1210 = v256
		v1214 = v285
		v1217 = v33
		v1219 = v283
		v1222 = v302
		v1232 = v284
		v1234 = v197
		v1235 = v282
		goto L66
	} else {
		goto L67
	}
L66:
	;
	m.G0 = v1222 + int32(16)
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+72))
	if v1242 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L67:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if v307 <= int32(0) {
		v1209 = l0
		v1210 = v256
		v1214 = v285
		v1217 = v33
		v1219 = v283
		v1222 = v302
		v1232 = v284
		v1234 = v197
		v1235 = v282
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v310 = l0
	v311 = v256
	v315 = v285
	v318 = v33
	v320 = v283
	v323 = v302
	v328 = v298
	v333 = v284
	v334 = v3
	v335 = v197
	v336 = v282
	goto L69
L69:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v340+v334<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+12)) = int32(0)
	v347 = F_CreateExprContext(m, v315)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L20
	} else {
		goto L71
	}
L70:
	;
	v1209 = v310
	v1210 = v311
	v1214 = v315
	v1217 = v318
	v1219 = v320
	v1222 = v323
	v1232 = v333
	v1234 = v335
	v1235 = v336
	goto L66
L71:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v315)+76))
	if v349 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v315)+100))
	v354 = F_CreatePartitionDirectory(m, v352, int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L20
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v357 = int32(0)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	if v359 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+76)) = v354
	goto L74
L76:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	v361 = v360
	goto L78
L77:
	;
	v361 = v357
	goto L78
L78:
	;
	v366 = F_palloc(m, v361<<(uint(int32(2))%32)+int32(24))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L20
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v366))) = v347
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	v372 = F_bms_copy(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L20
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+20)) = v361
	v375 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v366)+16)) = uint16(v375)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+8)) = v372
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4]))
	v384 = F_AllocSetContextCreateInternal(m, v379, int32(_a_F_standard_ExecutorStart_5), v375, int32(_a_F_standard_ExecutorStart_6), int32(_a_F_standard_ExecutorStart_7))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L20
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+12)) = v384
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	if v387 == int32(0) {
		v1164 = v357
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v315)+44))
	v1183 = F_lappend(m, v1182, v366)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L20
	} else {
		goto L193
	}
L83:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
	if v390 <= int32(0) {
		v1164 = v357
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v408 = v357
	v416 = int32(0)
	goto L85
L85:
	;
	v428 = v416 << (uint(int32(2)) % 32)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v387)+12))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v428+v429)))
	if v431 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v1164 = v1130
	goto L82
L87:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	v433 = v432
	goto L89
L88:
	;
	v433 = int32(0)
	goto L89
L89:
	;
	v439 = F_palloc(m, v433*int32(120)|int32(4))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L20
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v428+(v366+int32(24))))) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v439))) = v433
	if v431 == int32(0) {
		v1130 = v408
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v1149 = v416 + int32(1)
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
	if v1149 < v1150 {
		v408 = v1130
		v416 = v1149
		goto L85
	} else {
		goto L192
	}
L92:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v445 <= int32(0) {
		v1130 = v408
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v463 = v408
	v467 = int32(0)
	goto L94
L94:
	;
	v483 = v439 + int32(4) + v467*int32(120)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v484+v467<<(uint(int32(2))%32))))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	v490 = F_ExecGetRangeTableRelation(m, v315, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L20
	} else {
		goto L96
	}
L95:
	;
	v1130 = v1096
	goto L91
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = v490
	v493 = F_RelationGetPartitionKey(m, v490)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L20
	} else {
		goto L97
	}
L97:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v315)+76))
	v496 = F_PartitionDirectoryLookup(m, v495, v490)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L20
	} else {
		goto L98
	}
L98:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+4)) = v498
	v502 = F_palloc(m, v498<<(uint(int32(2))%32))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L20
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483)+8)) = v502
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	if v505 != v506 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v488)+8))
	v890 = F_bms_copy(m, v889)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L20
	} else {
		goto L149
	}
L101:
	;
	v588 = F_palloc(m, v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L20
	} else {
		goto L125
	}
L102:
	;
	v587 = v505 << (uint(int32(2)) % 32)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v488)+28))
	v513 = v505 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v513) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	if v575 != 0 {
		v587 = v513
		goto L101
	} else {
		goto L123
	}
L106:
	;
	v575 = int32(0)
	goto L105
L107:
	;
	v549 = v544
	v550 = v545
	v551 = v546
	goto L117
L108:
	;
	if (v510|v511)&int32(3) != 0 {
		v544 = v510
		v545 = v511
		v546 = v513
		goto L107
	} else {
		goto L111
	}
L109:
	;
	v537 = v510
	v538 = v511
	v539 = v513
	goto L110
L110:
	;
	if v539 == int32(0) {
		goto L106
	} else {
		goto L116
	}
L111:
	;
	v521 = v510
	v522 = v511
	v523 = v513
	goto L112
L112:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v521)))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	if v526 != v527 {
		v544 = v521
		v545 = v522
		v546 = v523
		goto L107
	} else {
		goto L114
	}
L113:
	;
	v537 = v532
	v538 = v530
	v539 = v534
	goto L110
L114:
	;
	v529 = int32(4)
	v530 = v522 + v529
	v532 = v521 + v529
	v534 = v523 - v529
	if base.Ui32(int32(3)) < base.Ui32(v534) {
		v521 = v532
		v522 = v530
		v523 = v534
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v544 = v537
	v545 = v538
	v546 = v539
	goto L107
L117:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549))))
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	if v554 == v555 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v575 = v554 - v555
	goto L105
L119:
	;
	v557 = int32(1)
	v562 = v551 - v557
	if v562 != 0 {
		v549 = v549 + v557
		v550 = v550 + v557
		v551 = v562
		goto L117
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	goto L118
L122:
	;
	goto L106
L123:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v488)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+12)) = v576
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v488)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+16)) = v578
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v582 = v580 << (uint(int32(2)) % 32)
	if v582 == int32(0) {
		goto L100
	} else {
		goto L124
	}
L124:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v488)+16))
	base.MemoryCopy(m, v502, v585, v582)
	goto L100
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483)+12)) = v588
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v594 = F_palloc(m, v591<<(uint(int32(2))%32))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L20
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483)+16)) = v594
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	if v597 <= int32(0) {
		goto L100
	} else {
		goto L127
	}
L127:
	;
	v600 = int32(0)
	v604 = v600
	v613 = v600
	goto L128
L128:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	if v632 <= v604 {
		v702 = v604
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L100
L130:
	;
	v706 = v702
	goto L137
L131:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v488)+28))
	v637 = v604
	goto L132
L132:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v634+v637<<(uint(int32(2))%32))))
	if v668 != 0 {
		v702 = v637
		goto L130
	} else {
		goto L134
	}
L133:
	;
	v702 = v632
	goto L130
L134:
	;
	v670 = v637 + int32(1)
	if v670 != v632 {
		v637 = v670
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v856 = v613 + int32(1)
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	if v856 < v857 {
		v604 = v827
		v613 = v856
		goto L128
	} else {
		goto L148
	}
L137:
	;
	if v632 <= v706 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v812 = v613 << (uint(int32(2)) % 32)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v483)+12))
	v815 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v812+v813))) = v815
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v483)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v817+v812))) = v815
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v821+v812))) = int32(0)
	v827 = v706
	goto L136
L139:
	;
	v770 = v706
	goto L142
L140:
	;
	v734 = int32(2)
	v735 = v706 << (uint(v734) % 32)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v488)+28))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v735+v736)))
	v740 = v613 << (uint(v734) % 32)
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v740+v741)))
	if v738 != v743 {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v483)+8))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v488)+16))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v747+v735)))
	*(*int32)(unsafe.Add(mBase, uint32(v745+v740))) = v749
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v483)+12))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v488)+20))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v753+v735)))
	*(*int32)(unsafe.Add(mBase, uint32(v751+v740))) = v755
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v488)+24))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v759+v735)))
	*(*int32)(unsafe.Add(mBase, uint32(v757+v740))) = v761
	v827 = v706 + int32(1)
	goto L136
L142:
	;
	v798 = v770 + int32(1)
	if v798 < v632 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	goto L138
L144:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v488)+28))
	v801 = int32(2)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v800+v798<<(uint(v801)%32))))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v805+v613<<(uint(v801)%32))))
	if v804 != v809 {
		v770 = v798
		goto L142
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	goto L143
L147:
	;
	v706 = v798
	goto L137
L148:
	;
	goto L129
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483)+20)) = v890
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v488)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+24)) = v893
	if v893 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v488)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+28)) = v908
	if v908 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v347)+64))
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897)+128)))
	if v898&int32(2) != 0 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	F_InitPartitionPruneContext(m, v483+int32(32), v893, v496, v493, int32(0), v347)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L20
	} else {
		goto L153
	}
L153:
	;
	v906 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+16)) = uint8(v906)
	goto L150
L154:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v488)+40))
	v920 = F_bms_add_members(m, v918, v919)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L20
	} else {
		goto L157
	}
L155:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v347)+64))
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912)+128)))
	if v913&int32(2) != 0 {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v916 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+17)) = uint8(v916)
	goto L154
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+4)) = v920
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v488)+32))
	if v923 == int32(0) {
		v1096 = v463
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v1115 = v467 + int32(1)
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v1115 < v1116 {
		v463 = v1096
		v467 = v1115
		goto L94
	} else {
		goto L191
	}
L159:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+16)))
	if v926 != 0 {
		v1096 = v463
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v483)+20))
	if v927 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	if v984 < int32(0) {
		v1096 = v463
		goto L158
	} else {
		goto L172
	}
L162:
	;
	v984 = base.I32_ctz(v970) | v971<<(uint(int32(5))%32)
	goto L161
L163:
	;
	v984 = int32(-2)
	goto L161
L164:
	;
	v937 = base.I32_div_s(int32(0), int32(32))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v927)+4))
	if v938 <= v937 {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v941 = v927 + int32(8)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v941+v937<<(uint(int32(2))%32))))
	v948 = v945 & int32(-1)
	if v948 != 0 {
		v970 = v948
		v971 = v937
		goto L162
	} else {
		goto L166
	}
L166:
	;
	v950 = v937 + int32(1)
	if v950 == v938 {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v953 = v950
	goto L168
L168:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v941+v953<<(uint(int32(2))%32))))
	if v960 != 0 {
		v970 = v960
		v971 = v953
		goto L162
	} else {
		goto L170
	}
L169:
	;
	goto L163
L170:
	;
	v962 = v953 + int32(1)
	if v962 != v938 {
		v953 = v962
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	v990 = v984
	v999 = v463
	goto L173
L173:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1017+v990<<(uint(int32(2))%32))))
	if v1021 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v1096 = v1024
	goto L158
L175:
	;
	v1022 = F_bms_add_member(m, v999, v1021)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L20
	} else {
		goto L178
	}
L176:
	;
	v1024 = v999
	goto L177
L177:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v483)+20))
	if v1025 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L178:
	;
	v1024 = v1022
	goto L177
L179:
	;
	if int32(0) <= v1081 {
		v990 = v1081
		v999 = v1024
		goto L173
	} else {
		goto L190
	}
L180:
	;
	v1081 = base.I32_ctz(v1067) | v1068<<(uint(int32(5))%32)
	goto L179
L181:
	;
	v1081 = int32(-2)
	goto L179
L182:
	;
	v1032 = v990 + int32(1)
	v1034 = base.I32_div_s(v1032, int32(32))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+4))
	if v1035 <= v1034 {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v1038 = v1025 + int32(8)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1038+v1034<<(uint(int32(2))%32))))
	v1045 = v1042 & (int32(-1) << (uint(v1032) % 32))
	if v1045 != 0 {
		v1067 = v1045
		v1068 = v1034
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v1047 = v1034 + int32(1)
	if v1047 == v1035 {
		goto L181
	} else {
		goto L185
	}
L185:
	;
	v1050 = v1047
	goto L186
L186:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1038+v1050<<(uint(int32(2))%32))))
	if v1057 != 0 {
		v1067 = v1057
		v1068 = v1050
		goto L180
	} else {
		goto L188
	}
L187:
	;
	goto L181
L188:
	;
	v1059 = v1050 + int32(1)
	if v1059 != v1035 {
		v1050 = v1059
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	goto L174
L191:
	;
	goto L95
L192:
	;
	goto L86
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+44)) = v1183
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+16)))
	if v1186 != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v315)+52))
	v1198 = F_bms_add_members(m, v1197, v1196)
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L20
	} else {
		goto L199
	}
L195:
	;
	v1190 = F_ExecFindMatchingSubPlans(m, v366, int32(1), v323+int32(12))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L20
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+12)) = v1164
	v1195 = int32(0)
	v1196 = v1164
	goto L194
L198:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	v1195 = v1190
	v1196 = v1192
	goto L194
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+52)) = v1198
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v315)+48))
	v1202 = F_lappend(m, v1201, v1195)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L20
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+48)) = v1202
	v1206 = v334 + int32(1)
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v1206 < v1207 {
		v334 = v1206
		goto L69
	} else {
		goto L201
	}
L201:
	;
	goto L70
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+44)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+36)) = v1713
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4])) = v1234
	m.G0 = v1217 + int32(128)
	return
L203:
	;
	v1664 = int32(0)
	goto L281
L204:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L20
	} else {
		goto L277
	}
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L20
	} else {
		goto L273
	}
L206:
	;
	v1509 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1214)+156)) = v1509
	*(*int32)(unsafe.Add(mBase, uint32(v1214)+104)) = v1509
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+64))
	if v1514 == v1509 {
		goto L257
	} else {
		goto L258
	}
L207:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+20))
	v1248 = F_palloc0(m, v1245<<(uint(int32(2))%32))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L20
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1214)+28)) = v1248
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+72))
	if v1251 == int32(0) {
		goto L206
	} else {
		goto L209
	}
L209:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1251)+4))
	if v1254 <= int32(0) {
		goto L206
	} else {
		goto L210
	}
L210:
	;
	v1265 = int32(0)
	goto L211
L211:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1251)+12))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1288+v1265<<(uint(int32(2))%32))))
	v1293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292)+32)))
	if v1293 != 0 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	goto L206
L213:
	;
	v1476 = v1265 + int32(1)
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1251)+4))
	if v1476 < v1477 {
		v1265 = v1476
		goto L211
	} else {
		goto L256
	}
L214:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+16))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+12))
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+4))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1295+v1296<<(uint(int32(2))%32)-int32(4))))
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+12))
	if v1303 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+52))
	v1307 = F_bms_is_member(m, v1296, v1306)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L20
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+16))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+16))
	if base.Ui32(int32(5)) <= base.Ui32(v1312) {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	if v1307 == int32(0) {
		goto L213
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	v1440 = F_palloc(m, int32(44))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L20
	} else {
		goto L255
	}
L221:
	;
	if v1312 == int32(5) {
		v1437 = int32(0)
		goto L220
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+4))
	v1336 = F_ExecGetRangeTableRelation(m, v1214, v1335)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L20
	} else {
		goto L228
	}
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L20
	} else {
		goto L225
	}
L225:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+16)) = v1322
	F_errmsg_internal(m, int32(_a_F_standard_ExecutorStart_8), v1217+int32(16))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L20
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(918), int32(_a_F_standard_ExecutorStart_9))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L20
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	if v1336 == int32(0) {
		v1437 = int32(0)
		goto L220
	} else {
		goto L229
	}
L229:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+48))
	v1341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1340)+119)))
	switch v1341 - int32(83) {
	case 0:
		goto L235
	default:
		goto L204
	case 19:
		goto L231
	case 26:
		goto L232
	case 29, 31:
		v1437 = v1336
		goto L220
	case 33:
		goto L234
	case 35:
		goto L233
	}
L230:
	;
	v1437 = v1336
	goto L220
L231:
	;
	v1432 = F_GetFdwRoutineForRelation(m, v1336, int32(0))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L20
	} else {
		goto L253
	}
L232:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+16))
	if v1407 == int32(4) {
		goto L230
	} else {
		goto L248
	}
L233:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L20
	} else {
		goto L244
	}
L234:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L20
	} else {
		goto L240
	}
L235:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L20
	} else {
		goto L236
	}
L236:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L20
	} else {
		goto L237
	}
L237:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+48)) = v1351 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_10), v1217+int32(48))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L20
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1197), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L20
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L20
	} else {
		goto L241
	}
L241:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+64)) = v1372 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_12), v1217-int32(-64))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L20
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1204), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L20
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+80)) = v1393 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_13), v1217+int32(80))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1211), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L20
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L20
	} else {
		goto L249
	}
L249:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L20
	} else {
		goto L250
	}
L250:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+96)) = v1417 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_14), v1217+int32(96))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L20
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1219), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L20
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+108))
	if v1434 == int32(0) {
		goto L205
	} else {
		goto L254
	}
L254:
	;
	goto L230
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+4)) = v1311
	*(*int32)(unsafe.Add(mBase, uint32(v1440))) = v1437
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+8)) = v1444
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+12)) = v1446
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+16)) = v1448
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+20)) = v1450
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+24)) = v1452
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+28))
	v1455 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+40)) = v1455
	*(*uint16)(unsafe.Add(mBase, uint32(v1440)+38)) = uint16(v1455)
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+34)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1440)+32)) = uint8(v1455)
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+28)) = v1454
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1464+v1444<<(uint(int32(2))%32)-int32(4)))) = v1440
	goto L213
L256:
	;
	goto L212
L257:
	;
	v1606 = F_ExecInitNode(m, v1232, v1214, v1210)
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L20
	} else {
		goto L269
	}
L258:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+4))
	if v1518 <= int32(0) {
		goto L257
	} else {
		goto L259
	}
L259:
	;
	v1522 = v1210 & int32(-29)
	v1528 = v1509
	v1529 = int32(1)
	goto L260
L260:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+12))
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1555+v1528<<(uint(int32(2))%32))))
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+68))
	v1561 = F_bms_is_member(m, v1529, v1560)
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L20
	} else {
		goto L262
	}
L261:
	;
	goto L257
L262:
	;
	if v1561 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1563 = v1522 | int32(4)
	goto L265
L264:
	;
	v1563 = v1522
	goto L265
L265:
	;
	v1564 = F_ExecInitNode(m, v1559, v1214, v1563)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L20
	} else {
		goto L266
	}
L266:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+144))
	v1567 = F_lappend(m, v1566, v1564)
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L20
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1214)+144)) = v1567
	v1570 = int32(1)
	v1573 = v1528 + v1570
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+4))
	if v1573 < v1574 {
		v1528 = v1573
		v1529 = v1529 + v1570
		goto L260
	} else {
		goto L268
	}
L268:
	;
	goto L261
L269:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1606)+56))
	if v1235 != int32(1) {
		v1713 = v1608
		goto L202
	} else {
		goto L270
	}
L270:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+44))
	if v1611 == int32(0) {
		v1713 = v1608
		goto L202
	} else {
		goto L271
	}
L271:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v1614 <= int32(0) {
		v1713 = v1608
		goto L202
	} else {
		goto L272
	}
L272:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	goto L203
L273:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L20
	} else {
		goto L274
	}
L274:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+112)) = v1626 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_15), v1217+int32(112))
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L20
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1228), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L20
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L20
	} else {
		goto L278
	}
L278:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+32)) = v1647 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_16), v1217+int32(32))
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L20
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1234), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L20
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1617+v1664<<(uint(int32(2))%32))))
	v1695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694)+26)))
	if v1695 != int32(1) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	v1703 = F_ExecInitExtraTupleSlot(m, v1214, int32(0), int32(_a_F_standard_ExecutorStart_17))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L20
	} else {
		goto L287
	}
L283:
	;
	v1699 = v1664 + int32(1)
	if v1699 != v1614 {
		v1664 = v1699
		goto L281
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	goto L282
L286:
	;
	v1713 = v1608
	goto L202
L287:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1606)+4))
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1705)+44))
	v1707 = F_ExecInitJunkFilter(m, v1706, v1703)
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L20
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1214)+60)) = v1707
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1707)+8))
	v1713 = v1710
	goto L202
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
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v860 int32
	_ = v860
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v959 int32
	_ = v959
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1053 int32
	_ = v1053
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1168 int32
	_ = v1168
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
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
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1303 int32
	_ = v1303
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1349 int32
	_ = v1349
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
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
		goto L271
	} else {
		goto L272
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
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
	goto L1
L24:
	;
	return
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v85
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v90
	goto L1
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
	goto L1
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
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
	v107 = F_make_pathkeys_for_sortclauses_extended(m, l0, l0+int32(256), v23, v102, int32(0), v20+int32(10), v102)
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v107
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)))
	if v110 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
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
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v128 == int32(0) {
		v276 = v3
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v285 = int32(0)
	if v276 == v285 {
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
	if v860 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L68:
	;
	goto L82
L69:
	;
	v320 = int32(0)
	goto L68
L70:
	;
	goto L71
L71:
	;
	v292 = int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v293 <= v292 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v296 = v292
	goto L74
L73:
	;
	v296 = v293
	goto L74
L74:
	;
	v300 = int32(0)
	v302 = v285
	goto L75
L75:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v276+int32(8)+v300<<(uint(int32(2))%32))))
	if v308 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v320 = v311
	goto L68
L77:
	;
	v311 = v302 + base.I32_popcnt(v308)
	goto L79
L78:
	;
	v311 = v302
	goto L79
L79:
	;
	v313 = v300 + int32(1)
	if v313 != v296 {
		v300 = v313
		v302 = v311
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	if v320 <= int32(0) {
		v860 = v3
		goto L67
	} else {
		goto L94
	}
L82:
	;
	goto L81
L94:
	;
	v368 = v276
	v371 = int32(0)
	v373 = v3
	goto L95
L95:
	;
	v377 = int32(0)
	if v368 == v377 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	if v769 == int32(0) {
		v860 = v806
		goto L67
	} else {
		goto L238
	}
L97:
	;
	if int32(0) <= v435 {
		goto L108
	} else {
		goto L109
	}
L98:
	;
	v435 = base.I32_ctz(v421) | v422<<(uint(int32(5))%32)
	goto L97
L99:
	;
	v435 = int32(-2)
	goto L97
L100:
	;
	v388 = base.I32_div_s(int32(0), int32(32))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	if v389 <= v388 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v392 = v368 + int32(8)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392+v388<<(uint(int32(2))%32))))
	v399 = v396 & int32(-1)
	if v399 != 0 {
		v421 = v399
		v422 = v388
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v401 = v388 + int32(1)
	if v401 == v389 {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	v404 = v401
	goto L104
L104:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v392+v404<<(uint(int32(2))%32))))
	if v411 != 0 {
		v421 = v411
		v422 = v404
		goto L98
	} else {
		goto L106
	}
L105:
	;
	goto L99
L106:
	;
	v413 = v404 + int32(1)
	if v413 != v389 {
		v404 = v413
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v440 = v377
	v442 = v377
	v443 = v435
	v446 = v368
	goto L111
L109:
	;
	v679 = v377
	v681 = v377
	v685 = v368
	goto L110
L110:
	;
	v694 = F_bms_del_members(m, v685, v681)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L24
	} else {
		goto L178
	}
L111:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456+v443<<(uint(int32(2))%32))))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+40))
	if v464 != 0 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v679 = v604
	v681 = v606
	v685 = v610
	goto L110
L113:
	;
	if v610 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L114:
	;
	if v440 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L115:
	;
	v466 = v464
	goto L117
L116:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v463)+36))
	v466 = v465
	goto L117
L117:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463)+32))
	v468 = F_make_pathkeys_for_sortclauses(m, l0, v466, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L24
	} else {
		goto L118
	}
L118:
	;
	if v468 == int32(0) {
		goto L114
	} else {
		goto L119
	}
L119:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	if v472 <= int32(0) {
		goto L114
	} else {
		goto L120
	}
L120:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v468)+12))
	v483 = int32(0)
	goto L121
L121:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v475+v483<<(uint(int32(2))%32))))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+41)))
	if v499 != int32(1) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v505 = F_bms_del_member(m, v446, v443)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L24
	} else {
		goto L127
	}
L123:
	;
	v503 = v483 + int32(1)
	if v503 != v472 {
		v483 = v503
		goto L121
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	goto L114
L127:
	;
	v604 = v440
	v606 = v442
	v610 = v505
	goto L113
L128:
	;
	v600 = F_bms_add_member(m, v442, v443)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L24
	} else {
		goto L165
	}
L129:
	;
	v598 = v597
	goto L128
L130:
	;
	if v107 == int32(0) {
		v597 = v468
		goto L129
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if v107 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v528 = F_list_copy(m, v107)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L24
	} else {
		goto L134
	}
L134:
	;
	v530 = F_append_pathkeys(m, v528, v468)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L24
	} else {
		goto L135
	}
L135:
	;
	v598 = v530
	goto L128
L136:
	;
	v532 = F_list_copy(m, v107)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L24
	} else {
		goto L139
	}
L137:
	;
	v536 = v468
	goto L138
L138:
	;
	if v440 == v536 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v534 = F_append_pathkeys(m, v532, v468)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L24
	} else {
		goto L140
	}
L140:
	;
	v536 = v534
	goto L138
L141:
	;
	switch v596 {
	case 0, 1:
		v598 = v440
		goto L128
	case 2:
		v597 = v536
		goto L129
	default:
		v604 = v440
		v606 = v442
		v610 = v446
		goto L113
	}
L142:
	;
	v596 = int32(0)
	goto L141
L143:
	;
	goto L144
L144:
	;
	v545 = int32(0)
	goto L147
L145:
	;
	if v585 != 0 {
		goto L162
	} else {
		goto L163
	}
L146:
	;
	v580 = int32(0)
	if v567 != 0 {
		goto L159
	} else {
		goto L160
	}
L147:
	;
	v549 = int32(0)
	if v440 == v549 {
		v559 = v549
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v596 = int32(3)
	goto L141
L149:
	;
	if v536 != 0 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v553 <= v545 {
		v559 = int32(0)
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	v559 = v555 + v545<<(uint(int32(2))%32)
	goto L149
L152:
	;
	v565 = int32(0)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v536)+12))
	if base.B2i32(v559 == v565)|base.B2i32(v567 == v565) != 0 {
		goto L146
	} else {
		goto L157
	}
L153:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if v545 < v560 {
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v562 = int32(0)
	v585 = base.B2i32(v559 == v562)
	v587 = v562
	goto L145
L156:
	;
	goto L155
L157:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v567+v545<<(uint(int32(2))%32))))
	if v575 == v577 {
		v545 = v545 + int32(1)
		goto L147
	} else {
		goto L158
	}
L158:
	;
	goto L148
L159:
	;
	v584 = int32(2)
	goto L161
L160:
	;
	v584 = v580
	goto L161
L161:
	;
	v585 = base.B2i32(v559 == v580)
	v587 = v584
	goto L145
L162:
	;
	v589 = v587
	goto L164
L163:
	;
	v589 = int32(1)
	goto L164
L164:
	;
	v596 = v589
	goto L141
L165:
	;
	v604 = v598
	v606 = v600
	v610 = v446
	goto L113
L166:
	;
	if int32(0) <= v674 {
		v440 = v604
		v442 = v606
		v443 = v674
		v446 = v610
		goto L111
	} else {
		goto L177
	}
L167:
	;
	v674 = base.I32_ctz(v660) | v661<<(uint(int32(5))%32)
	goto L166
L168:
	;
	v674 = int32(-2)
	goto L166
L169:
	;
	v625 = v443 + int32(1)
	v627 = base.I32_div_s(v625, int32(32))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v610)+4))
	if v628 <= v627 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v631 = v610 + int32(8)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v631+v627<<(uint(int32(2))%32))))
	v638 = v635 & (int32(-1) << (uint(v625) % 32))
	if v638 != 0 {
		v660 = v638
		v661 = v627
		goto L167
	} else {
		goto L171
	}
L171:
	;
	v640 = v627 + int32(1)
	if v640 == v628 {
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v643 = v640
	goto L173
L173:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v631+v643<<(uint(int32(2))%32))))
	if v650 != 0 {
		v660 = v650
		v661 = v643
		goto L167
	} else {
		goto L175
	}
L174:
	;
	goto L168
L175:
	;
	v652 = v643 + int32(1)
	if v652 != v628 {
		v643 = v652
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	goto L112
L178:
	;
	v696 = int32(0)
	if v681 == v696 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v732 = int32(0)
	if v373 == v732 {
		goto L193
	} else {
		goto L194
	}
L180:
	;
	v731 = int32(0)
	goto L179
L181:
	;
	goto L182
L182:
	;
	v703 = int32(1)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	if v704 <= v703 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v707 = v703
	goto L185
L184:
	;
	v707 = v704
	goto L185
L185:
	;
	v711 = int32(0)
	v713 = v696
	goto L186
L186:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v681+int32(8)+v711<<(uint(int32(2))%32))))
	if v719 != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v731 = v722
	goto L179
L188:
	;
	v722 = v713 + base.I32_popcnt(v719)
	goto L190
L189:
	;
	v722 = v713
	goto L190
L190:
	;
	v724 = v711 + int32(1)
	if v724 != v707 {
		v711 = v724
		v713 = v722
		goto L186
	} else {
		goto L191
	}
L191:
	;
	goto L187
L192:
	;
	v768 = base.B2i32(v767 < v731)
	if v767 < v731 {
		goto L205
	} else {
		goto L206
	}
L193:
	;
	v767 = int32(0)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v739 = int32(1)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v740 <= v739 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v743 = v739
	goto L198
L197:
	;
	v743 = v740
	goto L198
L198:
	;
	v747 = int32(0)
	v749 = v732
	goto L199
L199:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v373+int32(8)+v747<<(uint(int32(2))%32))))
	if v755 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v767 = v758
	goto L192
L201:
	;
	v758 = v749 + base.I32_popcnt(v755)
	goto L203
L202:
	;
	v758 = v749
	goto L203
L203:
	;
	v760 = v747 + int32(1)
	if v760 != v743 {
		v747 = v760
		v749 = v758
		goto L199
	} else {
		goto L204
	}
L204:
	;
	goto L200
L205:
	;
	v769 = v679
	goto L207
L206:
	;
	v769 = v371
	goto L207
L207:
	;
	v770 = int32(0)
	if v694 == v770 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v767 < v731 {
		goto L221
	} else {
		goto L222
	}
L209:
	;
	v805 = int32(0)
	goto L208
L210:
	;
	goto L211
L211:
	;
	v777 = int32(1)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v778 <= v777 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v781 = v777
	goto L214
L213:
	;
	v781 = v778
	goto L214
L214:
	;
	v785 = int32(0)
	v787 = v770
	goto L215
L215:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v694+int32(8)+v785<<(uint(int32(2))%32))))
	if v793 != 0 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v805 = v796
	goto L208
L217:
	;
	v796 = v787 + base.I32_popcnt(v793)
	goto L219
L218:
	;
	v796 = v787
	goto L219
L219:
	;
	v798 = v785 + int32(1)
	if v798 != v781 {
		v785 = v798
		v787 = v796
		goto L215
	} else {
		goto L220
	}
L220:
	;
	goto L216
L221:
	;
	v806 = v681
	goto L223
L222:
	;
	v806 = v373
	goto L223
L223:
	;
	v807 = int32(0)
	if v806 == v807 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if v842 < v805 {
		v368 = v694
		v371 = v769
		v373 = v806
		goto L95
	} else {
		goto L237
	}
L225:
	;
	v842 = int32(0)
	goto L224
L226:
	;
	goto L227
L227:
	;
	v814 = int32(1)
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v806)+4))
	if v815 <= v814 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v818 = v814
	goto L230
L229:
	;
	v818 = v815
	goto L230
L230:
	;
	v822 = int32(0)
	v824 = v807
	goto L231
L231:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v806+int32(8)+v822<<(uint(int32(2))%32))))
	if v830 != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v842 = v833
	goto L224
L233:
	;
	v833 = v824 + base.I32_popcnt(v830)
	goto L235
L234:
	;
	v833 = v824
	goto L235
L235:
	;
	v835 = v822 + int32(1)
	if v835 != v818 {
		v822 = v835
		v824 = v833
		goto L231
	} else {
		goto L236
	}
L236:
	;
	goto L232
L237:
	;
	goto L96
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v769
	v860 = v806
	goto L67
L239:
	;
	if v920 < int32(0) {
		goto L1
	} else {
		goto L250
	}
L240:
	;
	v920 = base.I32_ctz(v906) | v907<<(uint(int32(5))%32)
	goto L239
L241:
	;
	v920 = int32(-2)
	goto L239
L242:
	;
	v873 = base.I32_div_s(int32(0), int32(32))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	if v874 <= v873 {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v877 = v860 + int32(8)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v877+v873<<(uint(int32(2))%32))))
	v884 = v881 & int32(-1)
	if v884 != 0 {
		v906 = v884
		v907 = v873
		goto L240
	} else {
		goto L244
	}
L244:
	;
	v886 = v873 + int32(1)
	if v886 == v874 {
		goto L241
	} else {
		goto L245
	}
L245:
	;
	v889 = v886
	goto L246
L246:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v877+v889<<(uint(int32(2))%32))))
	if v896 != 0 {
		v906 = v896
		v907 = v889
		goto L240
	} else {
		goto L248
	}
L247:
	;
	goto L241
L248:
	;
	v898 = v889 + int32(1)
	if v898 != v874 {
		v889 = v898
		goto L246
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	v927 = v920
	goto L251
L251:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)+12))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v941+v927<<(uint(int32(2))%32))))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v945)+4))
	if v946 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	goto L1
L253:
	;
	if v860 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L254:
	;
	v949 = int32(0)
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v946)+4))
	if v950 <= v949 {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v959 = v949
	goto L256
L256:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v946)+12))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v970+v959<<(uint(int32(2))%32))))
	v975 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v974)+51)) = uint8(v975)
	v978 = v959 + v975
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v946)+4))
	if v978 < v979 {
		v959 = v978
		goto L256
	} else {
		goto L258
	}
L257:
	;
	goto L253
L258:
	;
	goto L257
L259:
	;
	if int32(0) <= v1053 {
		v927 = v1053
		goto L251
	} else {
		goto L270
	}
L260:
	;
	v1053 = base.I32_ctz(v1039) | v1040<<(uint(int32(5))%32)
	goto L259
L261:
	;
	v1053 = int32(-2)
	goto L259
L262:
	;
	v1004 = v927 + int32(1)
	v1006 = base.I32_div_s(v1004, int32(32))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	if v1007 <= v1006 {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1010 = v860 + int32(8)
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1010+v1006<<(uint(int32(2))%32))))
	v1017 = v1014 & (int32(-1) << (uint(v1004) % 32))
	if v1017 != 0 {
		v1039 = v1017
		v1040 = v1006
		goto L260
	} else {
		goto L264
	}
L264:
	;
	v1019 = v1006 + int32(1)
	if v1019 == v1007 {
		goto L261
	} else {
		goto L265
	}
L265:
	;
	v1022 = v1019
	goto L266
L266:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1010+v1022<<(uint(int32(2))%32))))
	if v1029 != 0 {
		v1039 = v1029
		v1040 = v1022
		goto L260
	} else {
		goto L268
	}
L267:
	;
	goto L261
L268:
	;
	v1031 = v1022 + int32(1)
	if v1031 != v1007 {
		v1022 = v1031
		goto L266
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	goto L252
L271:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1075)))
	v1077 = F_make_pathkeys_for_window(m, l0, v1076, v23)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L24
	} else {
		goto L274
	}
L272:
	;
	v1080 = int32(0)
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1080
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v24)+120))
	if v1082 != 0 {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	v1080 = v1077
	goto L273
L275:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v24)+124))
	v1102 = F_make_pathkeys_for_sortclauses(m, l0, v1101, v23)
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L24
	} else {
		goto L284
	}
L276:
	;
	v1083 = F_list_copy(m, v1082)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L24
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L275
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v1083
	v1089 = int32(0)
	v1093 = F_make_pathkeys_for_sortclauses_extended(m, l0, l0+int32(260), v23, int32(1), v1089, v20+int32(9), v1089)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L24
	} else {
		goto L280
	}
L280:
	;
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)))
	if v1096 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1097 = v1093
	goto L283
L282:
	;
	v1097 = int32(0)
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v1097
	goto L275
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1102
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1105 != 0 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v1377
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v1379 != 0 {
		goto L356
	} else {
		goto L357
	}
L286:
	;
	v1106 = int32(0)
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+32))
	v1109 = F_copyObjectImpl(m, v1108)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L24
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v1377 = int32(0)
	goto L285
L289:
	;
	if v1109 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+12))
	v1112 = v1111
	goto L292
L291:
	;
	v1112 = v1106
	goto L292
L292:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+20))
	if v1113 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+12))
	v1115 = v1114
	goto L295
L294:
	;
	v1115 = v1106
	goto L295
L295:
	;
	if v23 == int32(0) {
		v1331 = v1109
		goto L296
	} else {
		goto L297
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v1331
	v1349 = int32(0)
	v1354 = F_make_pathkeys_for_sortclauses_extended(m, l0, v20+int32(4), v23, v1349, v1349, v20+int32(3), v1349)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L24
	} else {
		goto L351
	}
L297:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v1118 <= int32(0) {
		v1331 = v1109
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1125 = v1118
	v1126 = v1115
	v1127 = v1112
	v1128 = int32(0)
	goto L299
L299:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1139+v1128<<(uint(int32(2))%32))))
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1143)+26)))
	if v1144 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v1331 = v1109
	goto L296
L301:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1127)))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1126)))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+4))
	v1150 = F_exprType(m, v1149)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L24
	} else {
		goto L304
	}
L302:
	;
	v1319 = v1125
	v1320 = v1126
	v1321 = v1127
	goto L303
L303:
	;
	v1327 = v1128 + int32(1)
	if v1327 < v1319 {
		v1125 = v1319
		v1126 = v1320
		v1127 = v1321
		v1128 = v1327
		goto L299
	} else {
		goto L350
	}
L304:
	;
	if v1148 != v1150 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1331 = int32(0)
	goto L296
L306:
	;
	goto L307
L307:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+20))
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+12))
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+4))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+12))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+4))
	v1159 = int32(0)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+16))
	if v1168 == v1159 {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1147)+4)) = v1292
	v1303 = v1127 + int32(4)
	if base.Ui32(v1303) < base.Ui32(v1157+v1158<<(uint(int32(2))%32)) {
		goto L344
	} else {
		goto L345
	}
L309:
	;
	if v23 == int32(0) {
		v1288 = int32(1)
		goto L312
	} else {
		goto L313
	}
L310:
	;
	v1292 = v1168
	goto L311
L311:
	;
	goto L308
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1143)+16)) = v1288
	v1292 = v1288
	goto L311
L313:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v1175 <= int32(0) {
		v1288 = int32(1)
		goto L312
	} else {
		goto L314
	}
L314:
	;
	v1178 = int32(0)
	if v1178 < v1175 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1181 = v1175
	goto L317
L316:
	;
	v1181 = v1178
	goto L317
L317:
	;
	v1183 = v1181 & int32(3)
	v1184 = int32(0)
	if int32(4) <= v1175 {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	v1288 = v1266 + int32(1)
	goto L312
L319:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1193 = v1184
	v1194 = int32(0)
	v1195 = v1159
	goto L322
L320:
	;
	v1230 = v1184
	v1232 = v1159
	goto L321
L321:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1241 = int32(0)
	v1243 = v1230
	v1245 = v1232
	goto L338
L322:
	;
	v1204 = v1189 + v1195<<(uint(int32(2))%32)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+12))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+16))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+8))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+16))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+4))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+16))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1204)))
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+16))
	if base.Ui32(v1193) < base.Ui32(v1212) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	if v1183 == int32(0) {
		v1266 = v1220
		goto L318
	} else {
		goto L337
	}
L324:
	;
	v1214 = v1212
	goto L326
L325:
	;
	v1214 = v1193
	goto L326
L326:
	;
	if base.Ui32(v1214) < base.Ui32(v1210) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1216 = v1210
	goto L329
L328:
	;
	v1216 = v1214
	goto L329
L329:
	;
	if base.Ui32(v1216) < base.Ui32(v1208) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1218 = v1208
	goto L332
L331:
	;
	v1218 = v1216
	goto L332
L332:
	;
	if base.Ui32(v1218) < base.Ui32(v1206) {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1220 = v1206
	goto L335
L334:
	;
	v1220 = v1218
	goto L335
L335:
	;
	v1221 = int32(4)
	v1222 = v1195 + v1221
	v1224 = v1194 + v1221
	if v1224 != v1181&int32(2147483644) {
		v1193 = v1220
		v1194 = v1224
		v1195 = v1222
		goto L322
	} else {
		goto L336
	}
L336:
	;
	goto L323
L337:
	;
	v1230 = v1220
	v1232 = v1222
	goto L321
L338:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1239+v1245<<(uint(int32(2))%32))))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+16))
	if base.Ui32(v1243) < base.Ui32(v1256) {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	v1266 = v1258
	goto L318
L340:
	;
	v1258 = v1256
	goto L342
L341:
	;
	v1258 = v1243
	goto L342
L342:
	;
	v1259 = int32(1)
	v1262 = v1241 + v1259
	if v1262 != v1183 {
		v1241 = v1262
		v1243 = v1258
		v1245 = v1245 + v1259
		goto L338
	} else {
		goto L343
	}
L343:
	;
	goto L339
L344:
	;
	v1309 = v1303
	goto L346
L345:
	;
	v1309 = int32(0)
	goto L346
L346:
	;
	v1311 = v1126 + int32(4)
	if base.Ui32(v1311) < base.Ui32(v1155+v1156<<(uint(int32(2))%32)) {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1317 = v1311
	goto L349
L348:
	;
	v1317 = int32(0)
	goto L349
L349:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v1319 = v1318
	v1320 = v1317
	v1321 = v1309
	goto L303
L350:
	;
	goto L300
L351:
	;
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
	if v1357 != 0 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1358 = v1354
	goto L354
L353:
	;
	v1358 = int32(0)
	goto L354
L354:
	;
	v1377 = v1358
	goto L285
L355:
	;
	m.G0 = v20 + int32(16)
	return
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1379
	goto L355
L357:
	;
	goto L358
L358:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v1381 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1381
	goto L355
L360:
	;
	goto L361
L361:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v1383 != 0 {
		goto L366
	} else {
		goto L367
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1395
	goto L355
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1383
	goto L355
L364:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+4))
	if v1396 <= v1397 {
		goto L362
	} else {
		goto L375
	}
L365:
	;
	if v1377 != 0 {
		goto L372
	} else {
		goto L373
	}
L366:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+4))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1385 != 0 {
		v1395 = v1385
		v1396 = v1384
		goto L364
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1389 != 0 {
		v1395 = v1389
		v1396 = int32(0)
		goto L364
	} else {
		goto L371
	}
L369:
	;
	if int32(0) < v1384 {
		goto L363
	} else {
		goto L370
	}
L370:
	;
	goto L365
L371:
	;
	goto L365
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1377
	goto L355
L373:
	;
	goto L374
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	goto L355
L375:
	;
	goto L363
}
