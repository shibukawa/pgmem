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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
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
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
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
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v775 int32
	_ = v775
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v838 int32
	_ = v838
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v1009 int32
	_ = v1009
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1087 int32
	_ = v1087
	var v1106 int32
	_ = v1106
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1140 int32
	_ = v1140
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1174 int32
	_ = v1174
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
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
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1276 int32
	_ = v1276
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1409 int32
	_ = v1409
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1472 int32
	_ = v1472
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
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
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1707 int32
	_ = v1707
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1731 int32
	_ = v1731
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
	v200 = F_CreateExecutorState(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L20
	} else {
		goto L35
	}
L2:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	if v55 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[1]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if v44 != 0 {
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
	v46 = int32(1)
	goto L9
L8:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+76)))
	v46 = v45
	goto L9
L9:
	;
	goto L6
L10:
	;
	if v46&int32(1) != 0 {
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
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v155 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v59 <= v58 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v65 = v58
	goto L16
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v65<<(uint(int32(2))%32))))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v96)+16))
	if v97&int64(-3) == int64(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L13
L18:
	;
	v122 = v65 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v122 < v123 {
		v65 = v122
		goto L16
	} else {
		goto L27
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v103 = F_get_rel_namespace(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[2]))
	goto L22
L22:
	;
	if base.B2i32(v107 != int32(0))&base.B2i32(v103 == v107) != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v112 = F_CreateCommandTag(m, v54)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112<<(uint(int32(3))%32))+uint32(_c_F_standard_ExecutorStart[3])))
	goto L25
L25:
	;
	F_PreventCommandIfReadOnly(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
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
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+25)))
	if v158 != int32(1) {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v161 = F_CreateCommandTag(m, v54)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161<<(uint(int32(3))%32))+uint32(_c_F_standard_ExecutorStart[3])))
	goto L33
L33:
	;
	F_PreventCommandIfParallelMode(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	goto L1
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v200
	v203 = int32(_a_F_standard_ExecutorStart_0)
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4]))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v200)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4])) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+88)) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+84))
	if v211 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v215 = F_palloc0(m, v212*int32(12))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L20
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+56)) = v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+96)) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(4)) <= base.Ui32(v222-int32(2)) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+92)) = v215
	goto L38
L40:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v266 = F_RegisterSnapshot(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L20
	} else {
		goto L57
	}
L41:
	;
	v253 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L20
	} else {
		goto L53
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L20
	} else {
		goto L50
	}
L43:
	;
	if v222 != int32(1) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v235 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L20
	} else {
		goto L49
	}
L46:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+72))
	if v230 != 0 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+25)))
	if v231 != 0 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v263 = l1 | int32(32)
	goto L40
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+64)) = v235
	v263 = l1
	goto L40
L50:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v242
	F_errmsg_internal(m, int32(_a_F_standard_ExecutorStart_1), v33)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(238), int32(_a_F_standard_ExecutorStart_3))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v200)+64)) = v253
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+25)))
	if v259&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v262 = l1
	goto L56
L55:
	;
	v262 = l1 | int32(32)
	goto L56
L56:
	;
	v263 = v262
	goto L40
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+8)) = v266
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v270 = F_RegisterSnapshot(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+128)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v270
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+132)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+176)) = v277
	if v263&int32(33) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v283 = int32(_a_F_standard_ExecutorStart_4)
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[5])) = v285 + int32(1)
	goto L62
L60:
	;
	goto L61
L61:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+36))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v290)+44))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290)+52))
	v296 = F_ExecCheckPermissions(m, v293, v294, int32(1))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L20
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v290)+52))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v290)+48))
	v300 = F_bms_copy(m, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	F_ExecInitRangeTable(m, v292, v293, v298, v300)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+36)) = v290
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v290)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v292)+40)) = v305
	v307 = m.G0
	v309 = v307 - int32(16)
	m.G0 = v309
	if v305 == int32(0) {
		v1217 = l0
		v1218 = v263
		v1220 = v292
		v1225 = v33
		v1229 = v290
		v1234 = v309
		v1244 = v291
		v1245 = v204
		v1246 = v289
		goto L66
	} else {
		goto L67
	}
L66:
	;
	m.G0 = v1234 + int32(16)
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+72))
	if v1250 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L67:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	if v314 <= int32(0) {
		v1217 = l0
		v1218 = v263
		v1220 = v292
		v1225 = v33
		v1229 = v290
		v1234 = v309
		v1244 = v291
		v1245 = v204
		v1246 = v289
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v317 = l0
	v318 = v263
	v320 = v292
	v321 = v305
	v325 = v33
	v329 = v290
	v334 = v309
	v339 = v3
	v344 = v291
	v345 = v204
	v346 = v289
	goto L69
L69:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v321)+12))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v347+v339<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+12)) = int32(0)
	v354 = F_CreateExprContext(m, v320)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L20
	} else {
		goto L71
	}
L70:
	;
	v1217 = v317
	v1218 = v318
	v1220 = v320
	v1225 = v325
	v1229 = v329
	v1234 = v334
	v1244 = v344
	v1245 = v345
	v1246 = v346
	goto L66
L71:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v320)+76))
	if v356 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v320)+100))
	v361 = F_CreatePartitionDirectory(m, v359, int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L20
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v364 = int32(0)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v351)+8))
	if v366 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+76)) = v361
	goto L74
L76:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	v368 = v367
	goto L78
L77:
	;
	v368 = v364
	goto L78
L78:
	;
	v373 = F_palloc(m, v368<<(uint(int32(2))%32)+int32(24))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L20
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v373))) = v354
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v379 = F_bms_copy(m, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L20
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+20)) = v368
	v382 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v373)+16)) = uint16(v382)
	*(*int32)(unsafe.Add(mBase, uint32(v373)+8)) = v379
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4]))
	v391 = F_AllocSetContextCreateInternal(m, v386, int32(_a_F_standard_ExecutorStart_5), v382, int32(_a_F_standard_ExecutorStart_6), int32(_a_F_standard_ExecutorStart_7))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L20
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+12)) = v391
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v351)+8))
	if v394 == int32(0) {
		v1174 = v364
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v320)+44))
	v1189 = F_lappend(m, v1188, v373)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L20
	} else {
		goto L196
	}
L83:
	;
	v397 = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v398 <= v397 {
		v1174 = v364
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v419 = v364
	v427 = v397
	goto L85
L85:
	;
	v435 = v427 << (uint(int32(2)) % 32)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v435+v436)))
	if v438 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v1174 = v1140
	goto L82
L87:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v440 = v439
	goto L89
L88:
	;
	v440 = int32(0)
	goto L89
L89:
	;
	v446 = F_palloc(m, v440*int32(120)|int32(4))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L20
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435+(v373+int32(24))))) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v446))) = v440
	if v438 == int32(0) {
		v1140 = v419
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v1155 = v427 + int32(1)
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v1155 < v1156 {
		v419 = v1140
		v427 = v1155
		goto L85
	} else {
		goto L195
	}
L92:
	;
	v452 = int32(0)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v453 <= v452 {
		v1140 = v419
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v474 = v419
	v479 = v452
	goto L94
L94:
	;
	v490 = v446 + int32(4) + v479*int32(120)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v438)+12))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v491+v479<<(uint(int32(2))%32))))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	v497 = F_ExecGetRangeTableRelation(m, v320, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L20
	} else {
		goto L96
	}
L95:
	;
	v1140 = v1106
	goto L91
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490))) = v497
	v500 = F_RelationGetPartitionKey(m, v497)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L20
	} else {
		goto L97
	}
L97:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v320)+76))
	v503 = F_PartitionDirectoryLookup(m, v502, v497)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L20
	} else {
		goto L98
	}
L98:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	*(*int32)(unsafe.Add(mBase, uint32(v490)+4)) = v505
	v509 = F_palloc(m, v505<<(uint(int32(2))%32))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L20
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490)+8)) = v509
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v495)+12))
	if v512 != v513 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v495)+8))
	v896 = F_bms_copy(m, v895)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L20
	} else {
		goto L152
	}
L101:
	;
	v594 = F_palloc(m, v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L20
	} else {
		goto L128
	}
L102:
	;
	v593 = v512 << (uint(int32(2)) % 32)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v503)+8))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v520 = v512 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v520) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	if v582 != 0 {
		v593 = v520
		goto L101
	} else {
		goto L123
	}
L106:
	;
	v582 = int32(0)
	goto L105
L107:
	;
	v556 = v551
	v557 = v552
	v558 = v553
	goto L117
L108:
	;
	if (v517|v518)&int32(3) != 0 {
		v551 = v517
		v552 = v518
		v553 = v520
		goto L107
	} else {
		goto L111
	}
L109:
	;
	v544 = v517
	v545 = v518
	v546 = v520
	goto L110
L110:
	;
	if v546 == int32(0) {
		goto L106
	} else {
		goto L116
	}
L111:
	;
	v528 = v517
	v529 = v518
	v530 = v520
	goto L112
L112:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	if v533 != v534 {
		v551 = v528
		v552 = v529
		v553 = v530
		goto L107
	} else {
		goto L114
	}
L113:
	;
	v544 = v539
	v545 = v537
	v546 = v541
	goto L110
L114:
	;
	v536 = int32(4)
	v537 = v529 + v536
	v539 = v528 + v536
	v541 = v530 - v536
	if base.Ui32(int32(3)) < base.Ui32(v541) {
		v528 = v539
		v529 = v537
		v530 = v541
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v551 = v544
	v552 = v545
	v553 = v546
	goto L107
L117:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	if v561 == v562 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v582 = v561 - v562
	goto L105
L119:
	;
	v564 = int32(1)
	v569 = v558 - v564
	if v569 != 0 {
		v556 = v556 + v564
		v557 = v557 + v564
		v558 = v569
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
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v495)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v490)+12)) = v583
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v495)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v490)+16)) = v585
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v495)+16))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v495)+12))
	v590 = v588 << (uint(int32(2)) % 32)
	if v590 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L100
L125:
	;
	v591 = F__emscripten_memcpy_bulkmem(m, v509, v587, v590)
	mBase = m.M
	goto L127
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490)+12)) = v594
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	v600 = F_palloc(m, v597<<(uint(int32(2))%32))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L20
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490)+16)) = v600
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	if v603 <= int32(0) {
		goto L100
	} else {
		goto L130
	}
L130:
	;
	v606 = int32(0)
	v615 = v606
	v622 = v606
	goto L131
L131:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v495)+12))
	if v638 <= v615 {
		v708 = v615
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L100
L133:
	;
	v711 = v708
	goto L140
L134:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v648 = v615
	goto L135
L135:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v640+v648<<(uint(int32(2))%32))))
	if v674 != 0 {
		v708 = v648
		goto L133
	} else {
		goto L137
	}
L136:
	;
	v708 = v638
	goto L133
L137:
	;
	v676 = v648 + int32(1)
	if v676 != v638 {
		v648 = v676
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v862 = v622 + int32(1)
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	if v862 < v863 {
		v615 = v838
		v622 = v862
		goto L131
	} else {
		goto L151
	}
L140:
	;
	if v638 <= v711 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v818 = v622 << (uint(int32(2)) % 32)
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v490)+12))
	v821 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v818+v819))) = v821
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v490)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v823+v818))) = v821
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v490)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v827+v818))) = int32(0)
	v838 = v711
	goto L139
L142:
	;
	v775 = v711
	goto L145
L143:
	;
	v740 = int32(2)
	v741 = v711 << (uint(v740) % 32)
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v741+v742)))
	v746 = v622 << (uint(v740) % 32)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v503)+8))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v746+v747)))
	if v744 != v749 {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v490)+8))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v495)+16))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v753+v741)))
	*(*int32)(unsafe.Add(mBase, uint32(v751+v746))) = v755
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v490)+12))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v495)+20))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v759+v741)))
	*(*int32)(unsafe.Add(mBase, uint32(v757+v746))) = v761
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v490)+16))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v495)+24))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v765+v741)))
	*(*int32)(unsafe.Add(mBase, uint32(v763+v746))) = v767
	v838 = v711 + int32(1)
	goto L139
L145:
	;
	v804 = v775 + int32(1)
	if v804 < v638 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L141
L147:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v807 = int32(2)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v806+v804<<(uint(v807)%32))))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v503)+8))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v811+v622<<(uint(v807)%32))))
	if v810 != v815 {
		v775 = v804
		goto L145
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	goto L146
L150:
	;
	v711 = v804
	goto L140
L151:
	;
	goto L132
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490)+20)) = v896
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v495)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v490)+24)) = v899
	if v899 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v495)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v490)+28)) = v914
	if v914 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v354)+64))
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903)+128)))
	if v904&int32(2) != 0 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	F_InitPartitionPruneContext(m, v490+int32(32), v899, v503, v500, int32(0), v354)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L20
	} else {
		goto L156
	}
L156:
	;
	v912 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v373)+16)) = uint8(v912)
	goto L153
L157:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v495)+40))
	v926 = F_bms_add_members(m, v924, v925)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L20
	} else {
		goto L160
	}
L158:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v354)+64))
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918)+128)))
	if v919&int32(2) != 0 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v922 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v373)+17)) = uint8(v922)
	goto L157
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+4)) = v926
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v495)+32))
	if v929 == int32(0) {
		v1106 = v474
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v1121 = v479 + int32(1)
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v1121 < v1122 {
		v474 = v1106
		v479 = v1121
		goto L94
	} else {
		goto L194
	}
L162:
	;
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+16)))
	if v932 != 0 {
		v1106 = v474
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v490)+20))
	if v933 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	if v990 < int32(0) {
		v1106 = v474
		goto L161
	} else {
		goto L175
	}
L165:
	;
	v990 = base.I32_ctz(v976) | v977<<(uint(int32(5))%32)
	goto L164
L166:
	;
	v990 = int32(-2)
	goto L164
L167:
	;
	v943 = base.I32_div_s(int32(0), int32(32))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	if v944 <= v943 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v947 = v933 + int32(8)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v947+v943<<(uint(int32(2))%32))))
	v954 = v951 & int32(-1)
	if v954 != 0 {
		v976 = v954
		v977 = v943
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v956 = v943 + int32(1)
	if v956 == v944 {
		goto L166
	} else {
		goto L170
	}
L170:
	;
	v959 = v956
	goto L171
L171:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v947+v959<<(uint(int32(2))%32))))
	if v966 != 0 {
		v976 = v966
		v977 = v959
		goto L165
	} else {
		goto L173
	}
L172:
	;
	goto L166
L173:
	;
	v968 = v959 + int32(1)
	if v968 != v944 {
		v959 = v968
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v995 = v990
	v1009 = v474
	goto L176
L176:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v490)+16))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1023+v995<<(uint(int32(2))%32))))
	if v1027 != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v1106 = v1030
	goto L161
L178:
	;
	v1028 = F_bms_add_member(m, v1009, v1027)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L20
	} else {
		goto L181
	}
L179:
	;
	v1030 = v1009
	goto L180
L180:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v490)+20))
	if v1031 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	v1030 = v1028
	goto L180
L182:
	;
	if int32(0) <= v1087 {
		v995 = v1087
		v1009 = v1030
		goto L176
	} else {
		goto L193
	}
L183:
	;
	v1087 = base.I32_ctz(v1073) | v1074<<(uint(int32(5))%32)
	goto L182
L184:
	;
	v1087 = int32(-2)
	goto L182
L185:
	;
	v1038 = v995 + int32(1)
	v1040 = base.I32_div_s(v1038, int32(32))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1031)+4))
	if v1041 <= v1040 {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v1044 = v1031 + int32(8)
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1044+v1040<<(uint(int32(2))%32))))
	v1051 = v1048 & (int32(-1) << (uint(v1038) % 32))
	if v1051 != 0 {
		v1073 = v1051
		v1074 = v1040
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v1053 = v1040 + int32(1)
	if v1053 == v1041 {
		goto L184
	} else {
		goto L188
	}
L188:
	;
	v1056 = v1053
	goto L189
L189:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1044+v1056<<(uint(int32(2))%32))))
	if v1063 != 0 {
		v1073 = v1063
		v1074 = v1056
		goto L183
	} else {
		goto L191
	}
L190:
	;
	goto L184
L191:
	;
	v1065 = v1056 + int32(1)
	if v1065 != v1041 {
		v1056 = v1065
		goto L189
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	goto L177
L194:
	;
	goto L95
L195:
	;
	goto L86
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+44)) = v1189
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+16)))
	if v1192 == int32(1) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v320)+52))
	v1206 = F_bms_add_members(m, v1205, v1204)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L20
	} else {
		goto L202
	}
L198:
	;
	v1198 = F_ExecFindMatchingSubPlans(m, v373, int32(1), v334+int32(12))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L20
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+12)) = v1174
	v1203 = int32(0)
	v1204 = v1174
	goto L197
L201:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	v1203 = v1198
	v1204 = v1200
	goto L197
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+52)) = v1206
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v320)+48))
	v1210 = F_lappend(m, v1209, v1203)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L20
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+48)) = v1210
	v1214 = v339 + int32(1)
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	if v1214 < v1215 {
		v339 = v1214
		goto L69
	} else {
		goto L204
	}
L204:
	;
	goto L70
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+44)) = v1614
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+36)) = v1731
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorStart[4])) = v1245
	m.G0 = v1225 + int32(128)
	return
L206:
	;
	v1671 = int32(0)
	goto L284
L207:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L20
	} else {
		goto L280
	}
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L20
	} else {
		goto L276
	}
L209:
	;
	v1517 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+156)) = v1517
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+104)) = v1517
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+64))
	if v1522 == v1517 {
		goto L260
	} else {
		goto L261
	}
L210:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+20))
	v1256 = F_palloc0(m, v1253<<(uint(int32(2))%32))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L20
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+28)) = v1256
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+72))
	if v1259 == int32(0) {
		goto L209
	} else {
		goto L212
	}
L212:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+4))
	if v1262 <= int32(0) {
		goto L209
	} else {
		goto L213
	}
L213:
	;
	v1276 = int32(0)
	goto L214
L214:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+12))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1296+v1276<<(uint(int32(2))%32))))
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300)+32)))
	if v1301 != 0 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	goto L209
L216:
	;
	v1484 = v1276 + int32(1)
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+4))
	if v1484 < v1485 {
		v1276 = v1484
		goto L214
	} else {
		goto L259
	}
L217:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+16))
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+12))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+4))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1303+v1304<<(uint(int32(2))%32)-int32(4))))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+12))
	if v1311 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+52))
	v1315 = F_bms_is_member(m, v1304, v1314)
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L20
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+16))
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+16))
	if base.Ui32(int32(5)) <= base.Ui32(v1320) {
		goto L224
	} else {
		goto L225
	}
L221:
	;
	if v1315 == int32(0) {
		goto L216
	} else {
		goto L222
	}
L222:
	;
	goto L220
L223:
	;
	v1448 = F_palloc(m, int32(44))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L20
	} else {
		goto L258
	}
L224:
	;
	if v1320 == int32(5) {
		v1446 = int32(0)
		goto L223
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+4))
	v1344 = F_ExecGetRangeTableRelation(m, v1220, v1343)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L20
	} else {
		goto L231
	}
L227:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L20
	} else {
		goto L228
	}
L228:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+16)) = v1330
	F_errmsg_internal(m, int32(_a_F_standard_ExecutorStart_8), v1225+int32(16))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L20
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(918), int32(_a_F_standard_ExecutorStart_9))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L20
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	if v1344 == int32(0) {
		v1446 = int32(0)
		goto L223
	} else {
		goto L232
	}
L232:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+48))
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1348)+119)))
	switch v1349 - int32(83) {
	case 0:
		goto L238
	default:
		goto L207
	case 19:
		goto L234
	case 26:
		goto L235
	case 29, 31:
		v1446 = v1344
		goto L223
	case 33:
		goto L237
	case 35:
		goto L236
	}
L233:
	;
	v1446 = v1344
	goto L223
L234:
	;
	v1440 = F_GetFdwRoutineForRelation(m, v1344, int32(0))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L20
	} else {
		goto L256
	}
L235:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+16))
	if v1415 == int32(4) {
		goto L233
	} else {
		goto L251
	}
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L20
	} else {
		goto L247
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L20
	} else {
		goto L243
	}
L238:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L20
	} else {
		goto L239
	}
L239:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L20
	} else {
		goto L240
	}
L240:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+48)) = v1359 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_10), v1225+int32(48))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L20
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1197), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L20
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L20
	} else {
		goto L244
	}
L244:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+64)) = v1380 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_12), v1225-int32(-64))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1204), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L20
	} else {
		goto L248
	}
L248:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+80)) = v1401 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_13), v1225+int32(80))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L20
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1211), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L20
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L20
	} else {
		goto L252
	}
L252:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L20
	} else {
		goto L253
	}
L253:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+96)) = v1425 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_14), v1225+int32(96))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L20
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1219), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L20
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+108))
	if v1442 == int32(0) {
		goto L208
	} else {
		goto L257
	}
L257:
	;
	goto L233
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+4)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v1448))) = v1446
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+8)) = v1452
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+12)) = v1454
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+16)) = v1456
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+20)) = v1458
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+24)) = v1460
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+28))
	v1463 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+40)) = v1463
	*(*uint16)(unsafe.Add(mBase, uint32(v1448)+38)) = uint16(v1463)
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+34)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1448)+32)) = uint8(v1463)
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+28)) = v1462
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1472+v1452<<(uint(int32(2))%32)-int32(4)))) = v1448
	goto L216
L259:
	;
	goto L215
L260:
	;
	v1614 = F_ExecInitNode(m, v1244, v1220, v1218)
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L20
	} else {
		goto L272
	}
L261:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+4))
	if v1525 <= int32(0) {
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1529 = v1218 & int32(-29)
	v1535 = v1517
	v1537 = int32(1)
	goto L263
L263:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+12))
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1563+v1535<<(uint(int32(2))%32))))
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+68))
	v1569 = F_bms_is_member(m, v1537, v1568)
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L20
	} else {
		goto L265
	}
L264:
	;
	goto L260
L265:
	;
	if v1569 != 0 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1571 = v1529 | int32(4)
	goto L268
L267:
	;
	v1571 = v1529
	goto L268
L268:
	;
	v1572 = F_ExecInitNode(m, v1567, v1220, v1571)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L20
	} else {
		goto L269
	}
L269:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+144))
	v1575 = F_lappend(m, v1574, v1572)
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L20
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+144)) = v1575
	v1578 = int32(1)
	v1581 = v1535 + v1578
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+4))
	if v1581 < v1582 {
		v1535 = v1581
		v1537 = v1537 + v1578
		goto L263
	} else {
		goto L271
	}
L271:
	;
	goto L264
L272:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1614)+56))
	if v1246 != int32(1) {
		v1731 = v1616
		goto L205
	} else {
		goto L273
	}
L273:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+44))
	if v1619 == int32(0) {
		v1731 = v1616
		goto L205
	} else {
		goto L274
	}
L274:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1619)+4))
	if v1622 <= int32(0) {
		v1731 = v1616
		goto L205
	} else {
		goto L275
	}
L275:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1619)+12))
	goto L206
L276:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L20
	} else {
		goto L277
	}
L277:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+112)) = v1634 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_15), v1225+int32(112))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L20
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1228), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L20
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L20
	} else {
		goto L281
	}
L281:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1225)+32)) = v1655 + int32(4)
	F_errmsg(m, int32(_a_F_standard_ExecutorStart_16), v1225+int32(32))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L20
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(_a_F_standard_ExecutorStart_2), int32(1234), int32(_a_F_standard_ExecutorStart_11))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L20
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1625+v1671<<(uint(int32(2))%32))))
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1702)+26)))
	if v1703 != int32(1) {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	v1711 = F_ExecInitExtraTupleSlot(m, v1220, int32(0), int32(_a_F_standard_ExecutorStart_17))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L20
	} else {
		goto L290
	}
L286:
	;
	v1707 = v1671 + int32(1)
	if v1707 != v1622 {
		v1671 = v1707
		goto L284
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	goto L285
L289:
	;
	v1731 = v1616
	goto L205
L290:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1614)+4))
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+44))
	v1715 = F_ExecInitJunkFilter(m, v1714, v1711)
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L20
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+60)) = v1715
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1715)+8))
	v1731 = v1718
	goto L205
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
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
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
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
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
	var v479 int32
	_ = v479
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
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v671 int32
	_ = v671
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v858 int32
	_ = v858
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v917 int32
	_ = v917
	var v924 int32
	_ = v924
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1050 int32
	_ = v1050
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1163 int32
	_ = v1163
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1296 int32
	_ = v1296
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1331 int32
	_ = v1331
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
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
		goto L31
	} else {
		goto L32
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
	v78 = int32(0)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+45)))
	v86 = F_make_pathkeys_for_sortclauses_extended(m, l0, v20+int32(12), v23, v78, v82, v20+int32(11), v78)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v86
	if v86 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v90 = v89
	goto L28
L27:
	;
	v90 = v78
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v90
	goto L1
L29:
	;
	v285 = int32(0)
	if v274 == v285 {
		goto L71
	} else {
		goto L72
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
	goto L1
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
	if v97 <= int32(0) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v100 = int32(0)
	v103 = int32(1)
	v108 = F_make_pathkeys_for_sortclauses_extended(m, l0, l0+int32(256), v23, v103, v100, v20+int32(10), v103)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L24
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v108
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)))
	if v111 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
	goto L1
L37:
	;
	goto L38
L38:
	;
	if v108 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v117 = v116
	goto L41
L40:
	;
	v117 = v100
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
	if v119 <= int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_qp_callback[0])))
	if v123 != int32(1) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v126 == int32(0) {
		v274 = v3
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v129 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v274 = v3
	goto L29
L46:
	;
	goto L47
L47:
	;
	v136 = v3
	v138 = v3
	v139 = v129
	goto L48
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149+v136<<(uint(int32(2))%32))))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+50)))
	if v157 != int32(110) {
		v252 = v138
		v253 = v139
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v274 = v252
	goto L29
L50:
	;
	v264 = v136 + int32(1)
	if v264 < v253 {
		v136 = v264
		v138 = v252
		v139 = v253
		goto L48
	} else {
		goto L68
	}
L51:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+40))
	if v160 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v156)+36))
	if v163 == int32(0) {
		v252 = v138
		v253 = v139
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v156)+44))
	if v166 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v243 = F_bms_add_member(m, v138, v136)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L24
	} else {
		goto L67
	}
L57:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v156)+32))
	if v169 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v172 <= int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v185 = int32(0)
	goto L60
L60:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v175+v185<<(uint(int32(2))%32))))
	v200 = v197
	goto L62
L61:
	;
	goto L56
L62:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	if v216 == int32(27) {
		v200 = v215
		goto L62
	} else {
		goto L64
	}
L63:
	;
	if base.Ui32(int32(1)) < base.Ui32(v216-int32(6)) {
		v252 = v138
		v253 = v139
		goto L50
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v224 = v185 + int32(1)
	if v172 != v224 {
		v185 = v224
		goto L60
	} else {
		goto L66
	}
L66:
	;
	goto L61
L67:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v252 = v243
	v253 = v245
	goto L50
L68:
	;
	goto L49
L69:
	;
	if v858 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L70:
	;
	goto L84
L71:
	;
	v321 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v293 = int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v294 <= v293 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v297 = v293
	goto L76
L75:
	;
	v297 = v294
	goto L76
L76:
	;
	v301 = int32(0)
	v303 = v285
	goto L77
L77:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v274+int32(8)+v301<<(uint(int32(2))%32))))
	if v309 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v321 = v312
	goto L70
L79:
	;
	v312 = v303 + base.I32_popcnt(v309)
	goto L81
L80:
	;
	v312 = v303
	goto L81
L81:
	;
	v314 = v301 + int32(1)
	if v314 != v297 {
		v301 = v314
		v303 = v312
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	if v321 <= int32(0) {
		v858 = v285
		goto L69
	} else {
		goto L96
	}
L84:
	;
	goto L83
L96:
	;
	v366 = v274
	v374 = v285
	v375 = v3
	goto L97
L97:
	;
	v377 = int32(0)
	if v366 == v377 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	if v766 == int32(0) {
		v858 = v803
		goto L69
	} else {
		goto L238
	}
L99:
	;
	if int32(0) <= v435 {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v435 = base.I32_ctz(v421) | v422<<(uint(int32(5))%32)
	goto L99
L101:
	;
	v435 = int32(-2)
	goto L99
L102:
	;
	v388 = base.I32_div_s(int32(0), int32(32))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	if v389 <= v388 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v392 = v366 + int32(8)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392+v388<<(uint(int32(2))%32))))
	v399 = v396 & int32(-1)
	if v399 != 0 {
		v421 = v399
		v422 = v388
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v401 = v388 + int32(1)
	if v401 == v389 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v404 = v401
	goto L106
L106:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v392+v404<<(uint(int32(2))%32))))
	if v411 != 0 {
		v421 = v411
		v422 = v404
		goto L100
	} else {
		goto L108
	}
L107:
	;
	goto L101
L108:
	;
	v413 = v404 + int32(1)
	if v413 != v389 {
		v404 = v413
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v443 = v435
	v444 = v366
	v446 = v377
	v449 = v377
	goto L113
L111:
	;
	v680 = v366
	v682 = v377
	v685 = v377
	goto L112
L112:
	;
	v691 = F_bms_del_members(m, v680, v685)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L24
	} else {
		goto L178
	}
L113:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456+v443<<(uint(int32(2))%32))))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+40))
	if v464 != 0 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v680 = v605
	v682 = v607
	v685 = v610
	goto L112
L115:
	;
	if v605 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L116:
	;
	if v446 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L117:
	;
	v466 = v464
	goto L119
L118:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v463)+36))
	v466 = v465
	goto L119
L119:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463)+32))
	v468 = F_make_pathkeys_for_sortclauses(m, l0, v466, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L24
	} else {
		goto L120
	}
L120:
	;
	if v468 == int32(0) {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	if v472 <= int32(0) {
		goto L116
	} else {
		goto L122
	}
L122:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v468)+12))
	v479 = int32(0)
	goto L123
L123:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v475+v479<<(uint(int32(2))%32))))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+41)))
	if v499 != int32(1) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v505 = F_bms_del_member(m, v444, v443)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L24
	} else {
		goto L129
	}
L125:
	;
	v503 = v479 + int32(1)
	if v503 != v472 {
		v479 = v503
		goto L123
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	goto L116
L129:
	;
	v605 = v505
	v607 = v446
	v610 = v449
	goto L115
L130:
	;
	v597 = F_bms_add_member(m, v449, v443)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L24
	} else {
		goto L165
	}
L131:
	;
	v596 = v594
	goto L130
L132:
	;
	if v108 == int32(0) {
		v594 = v468
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	if v108 != 0 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v528 = F_list_copy(m, v108)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L24
	} else {
		goto L136
	}
L136:
	;
	v530 = F_append_pathkeys(m, v528, v468)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L24
	} else {
		goto L137
	}
L137:
	;
	v596 = v530
	goto L130
L138:
	;
	v532 = F_list_copy(m, v108)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L24
	} else {
		goto L141
	}
L139:
	;
	v536 = v468
	goto L140
L140:
	;
	if v446 == v536 {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	v534 = F_append_pathkeys(m, v532, v468)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L24
	} else {
		goto L142
	}
L142:
	;
	v536 = v534
	goto L140
L143:
	;
	switch v593 {
	case 0, 1:
		v596 = v446
		goto L130
	case 2:
		v594 = v536
		goto L131
	default:
		v605 = v444
		v607 = v446
		v610 = v449
		goto L115
	}
L144:
	;
	v593 = int32(0)
	goto L143
L145:
	;
	goto L146
L146:
	;
	v544 = int32(0)
	goto L149
L147:
	;
	if v583 != 0 {
		goto L162
	} else {
		goto L163
	}
L148:
	;
	v577 = int32(0)
	v583 = base.B2i32(v557 == v577)
	v585 = base.B2i32(v566 != v577) << (uint(int32(1)) % 32)
	goto L147
L149:
	;
	v547 = int32(0)
	if v446 == v547 {
		v557 = v547
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v593 = int32(3)
	goto L143
L151:
	;
	if v536 != 0 {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	if v551 <= v544 {
		v557 = int32(0)
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v446)+12))
	v557 = v553 + v544<<(uint(int32(2))%32)
	goto L151
L154:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v536)+12))
	v566 = v563 + v544<<(uint(int32(2))%32)
	if v557 == int32(0) {
		goto L148
	} else {
		goto L159
	}
L155:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if v544 < v558 {
		goto L154
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v560 = int32(0)
	v583 = base.B2i32(v557 == v560)
	v585 = v560
	goto L147
L158:
	;
	goto L157
L159:
	;
	if v566 == int32(0) {
		goto L148
	} else {
		goto L160
	}
L160:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	if v573 == v574 {
		v544 = v544 + int32(1)
		goto L149
	} else {
		goto L161
	}
L161:
	;
	goto L150
L162:
	;
	v587 = v585
	goto L164
L163:
	;
	v587 = int32(1)
	goto L164
L164:
	;
	v593 = v587
	goto L143
L165:
	;
	v605 = v444
	v607 = v596
	v610 = v597
	goto L115
L166:
	;
	if int32(0) <= v671 {
		v443 = v671
		v444 = v605
		v446 = v607
		v449 = v610
		goto L113
	} else {
		goto L177
	}
L167:
	;
	v671 = base.I32_ctz(v657) | v658<<(uint(int32(5))%32)
	goto L166
L168:
	;
	v671 = int32(-2)
	goto L166
L169:
	;
	v622 = v443 + int32(1)
	v624 = base.I32_div_s(v622, int32(32))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	if v625 <= v624 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v628 = v605 + int32(8)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628+v624<<(uint(int32(2))%32))))
	v635 = v632 & (int32(-1) << (uint(v622) % 32))
	if v635 != 0 {
		v657 = v635
		v658 = v624
		goto L167
	} else {
		goto L171
	}
L171:
	;
	v637 = v624 + int32(1)
	if v637 == v625 {
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v640 = v637
	goto L173
L173:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v628+v640<<(uint(int32(2))%32))))
	if v647 != 0 {
		v657 = v647
		v658 = v640
		goto L167
	} else {
		goto L175
	}
L174:
	;
	goto L168
L175:
	;
	v649 = v640 + int32(1)
	if v649 != v625 {
		v640 = v649
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	goto L114
L178:
	;
	v693 = int32(0)
	if v685 == v693 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v729 = int32(0)
	if v374 == v729 {
		goto L193
	} else {
		goto L194
	}
L180:
	;
	v728 = int32(0)
	goto L179
L181:
	;
	goto L182
L182:
	;
	v700 = int32(1)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v685)+4))
	if v701 <= v700 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v704 = v700
	goto L185
L184:
	;
	v704 = v701
	goto L185
L185:
	;
	v708 = int32(0)
	v710 = v693
	goto L186
L186:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v685+int32(8)+v708<<(uint(int32(2))%32))))
	if v716 != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v728 = v719
	goto L179
L188:
	;
	v719 = v710 + base.I32_popcnt(v716)
	goto L190
L189:
	;
	v719 = v710
	goto L190
L190:
	;
	v721 = v708 + int32(1)
	if v721 != v704 {
		v708 = v721
		v710 = v719
		goto L186
	} else {
		goto L191
	}
L191:
	;
	goto L187
L192:
	;
	v765 = base.B2i32(v764 < v728)
	if v764 < v728 {
		goto L205
	} else {
		goto L206
	}
L193:
	;
	v764 = int32(0)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v736 = int32(1)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v737 <= v736 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v740 = v736
	goto L198
L197:
	;
	v740 = v737
	goto L198
L198:
	;
	v744 = int32(0)
	v746 = v729
	goto L199
L199:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v374+int32(8)+v744<<(uint(int32(2))%32))))
	if v752 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v764 = v755
	goto L192
L201:
	;
	v755 = v746 + base.I32_popcnt(v752)
	goto L203
L202:
	;
	v755 = v746
	goto L203
L203:
	;
	v757 = v744 + int32(1)
	if v757 != v740 {
		v744 = v757
		v746 = v755
		goto L199
	} else {
		goto L204
	}
L204:
	;
	goto L200
L205:
	;
	v766 = v682
	goto L207
L206:
	;
	v766 = v375
	goto L207
L207:
	;
	v767 = int32(0)
	if v691 == v767 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v764 < v728 {
		goto L221
	} else {
		goto L222
	}
L209:
	;
	v802 = int32(0)
	goto L208
L210:
	;
	goto L211
L211:
	;
	v774 = int32(1)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	if v775 <= v774 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v778 = v774
	goto L214
L213:
	;
	v778 = v775
	goto L214
L214:
	;
	v782 = int32(0)
	v784 = v767
	goto L215
L215:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v691+int32(8)+v782<<(uint(int32(2))%32))))
	if v790 != 0 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v802 = v793
	goto L208
L217:
	;
	v793 = v784 + base.I32_popcnt(v790)
	goto L219
L218:
	;
	v793 = v784
	goto L219
L219:
	;
	v795 = v782 + int32(1)
	if v795 != v778 {
		v782 = v795
		v784 = v793
		goto L215
	} else {
		goto L220
	}
L220:
	;
	goto L216
L221:
	;
	v803 = v685
	goto L223
L222:
	;
	v803 = v374
	goto L223
L223:
	;
	v804 = int32(0)
	if v803 == v804 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if v839 < v802 {
		v366 = v691
		v374 = v803
		v375 = v766
		goto L97
	} else {
		goto L237
	}
L225:
	;
	v839 = int32(0)
	goto L224
L226:
	;
	goto L227
L227:
	;
	v811 = int32(1)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	if v812 <= v811 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v815 = v811
	goto L230
L229:
	;
	v815 = v812
	goto L230
L230:
	;
	v819 = int32(0)
	v821 = v804
	goto L231
L231:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v803+int32(8)+v819<<(uint(int32(2))%32))))
	if v827 != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v839 = v830
	goto L224
L233:
	;
	v830 = v821 + base.I32_popcnt(v827)
	goto L235
L234:
	;
	v830 = v821
	goto L235
L235:
	;
	v832 = v819 + int32(1)
	if v832 != v815 {
		v819 = v832
		v821 = v830
		goto L231
	} else {
		goto L236
	}
L236:
	;
	goto L232
L237:
	;
	goto L98
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v766
	v858 = v803
	goto L69
L239:
	;
	if v917 < int32(0) {
		goto L1
	} else {
		goto L250
	}
L240:
	;
	v917 = base.I32_ctz(v903) | v904<<(uint(int32(5))%32)
	goto L239
L241:
	;
	v917 = int32(-2)
	goto L239
L242:
	;
	v870 = base.I32_div_s(int32(0), int32(32))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	if v871 <= v870 {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v874 = v858 + int32(8)
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v874+v870<<(uint(int32(2))%32))))
	v881 = v878 & int32(-1)
	if v881 != 0 {
		v903 = v881
		v904 = v870
		goto L240
	} else {
		goto L244
	}
L244:
	;
	v883 = v870 + int32(1)
	if v883 == v871 {
		goto L241
	} else {
		goto L245
	}
L245:
	;
	v886 = v883
	goto L246
L246:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v874+v886<<(uint(int32(2))%32))))
	if v893 != 0 {
		v903 = v893
		v904 = v886
		goto L240
	} else {
		goto L248
	}
L247:
	;
	goto L241
L248:
	;
	v895 = v886 + int32(1)
	if v895 != v871 {
		v886 = v895
		goto L246
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	v924 = v917
	goto L251
L251:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v937)+12))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v938+v924<<(uint(int32(2))%32))))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v942)+4))
	if v943 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	goto L1
L253:
	;
	if v858 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L254:
	;
	v946 = int32(0)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	if v947 <= v946 {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v952 = v946
	goto L256
L256:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v943)+12))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v967+v952<<(uint(int32(2))%32))))
	v972 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v971)+51)) = uint8(v972)
	v975 = v952 + v972
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	if v975 < v976 {
		v952 = v975
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
	if int32(0) <= v1050 {
		v924 = v1050
		goto L251
	} else {
		goto L270
	}
L260:
	;
	v1050 = base.I32_ctz(v1036) | v1037<<(uint(int32(5))%32)
	goto L259
L261:
	;
	v1050 = int32(-2)
	goto L259
L262:
	;
	v1001 = v924 + int32(1)
	v1003 = base.I32_div_s(v1001, int32(32))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	if v1004 <= v1003 {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1007 = v858 + int32(8)
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1007+v1003<<(uint(int32(2))%32))))
	v1014 = v1011 & (int32(-1) << (uint(v1001) % 32))
	if v1014 != 0 {
		v1036 = v1014
		v1037 = v1003
		goto L260
	} else {
		goto L264
	}
L264:
	;
	v1016 = v1003 + int32(1)
	if v1016 == v1004 {
		goto L261
	} else {
		goto L265
	}
L265:
	;
	v1019 = v1016
	goto L266
L266:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1007+v1019<<(uint(int32(2))%32))))
	if v1026 != 0 {
		v1036 = v1026
		v1037 = v1019
		goto L260
	} else {
		goto L268
	}
L267:
	;
	goto L261
L268:
	;
	v1028 = v1019 + int32(1)
	if v1028 != v1004 {
		v1019 = v1028
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
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1070)))
	v1072 = F_make_pathkeys_for_window(m, l0, v1071, v23)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L24
	} else {
		goto L274
	}
L272:
	;
	v1075 = int32(0)
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1075
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v24)+120))
	if v1077 != 0 {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	v1075 = v1072
	goto L273
L275:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v24)+124))
	v1097 = F_make_pathkeys_for_sortclauses(m, l0, v1096, v23)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L24
	} else {
		goto L284
	}
L276:
	;
	v1078 = F_list_copy(m, v1077)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v1078
	v1084 = int32(0)
	v1088 = F_make_pathkeys_for_sortclauses_extended(m, l0, l0+int32(260), v23, int32(1), v1084, v20+int32(9), v1084)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L24
	} else {
		goto L280
	}
L280:
	;
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)))
	if v1091 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1092 = v1088
	goto L283
L282:
	;
	v1092 = int32(0)
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v1092
	goto L275
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1097
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1100 != 0 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v1371
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v1373 != 0 {
		goto L357
	} else {
		goto L358
	}
L286:
	;
	v1101 = int32(0)
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+32))
	v1104 = F_copyObjectImpl(m, v1103)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L24
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v1371 = int32(0)
	goto L285
L289:
	;
	if v1104 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+12))
	v1107 = v1106
	goto L292
L291:
	;
	v1107 = v1101
	goto L292
L292:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+20))
	if v1108 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+12))
	v1110 = v1109
	goto L295
L294:
	;
	v1110 = v1101
	goto L295
L295:
	;
	if v23 == int32(0) {
		v1331 = v1104
		goto L296
	} else {
		goto L297
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v1331
	v1343 = int32(0)
	v1348 = F_make_pathkeys_for_sortclauses_extended(m, l0, v20+int32(4), v23, v1343, v1343, v20+int32(3), v1343)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L24
	} else {
		goto L352
	}
L297:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v1113 <= int32(0) {
		v1331 = v1104
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1119 = int32(0)
	v1121 = v1110
	v1122 = v1107
	v1124 = v1113
	goto L299
L299:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1134+v1119<<(uint(int32(2))%32))))
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138)+26)))
	if v1139 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v1331 = v1104
	goto L296
L301:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1122)))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1121)))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1138)+4))
	v1145 = F_exprType(m, v1144)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L24
	} else {
		goto L304
	}
L302:
	;
	v1314 = v1121
	v1315 = v1122
	v1317 = v1124
	goto L303
L303:
	;
	v1321 = v1119 + int32(1)
	if v1321 < v1317 {
		v1119 = v1321
		v1121 = v1314
		v1122 = v1315
		v1124 = v1317
		goto L299
	} else {
		goto L351
	}
L304:
	;
	if v1143 != v1145 {
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
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+20))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+12))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+4))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+12))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+4))
	v1154 = int32(0)
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1138)+16))
	if v1163 == v1154 {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1142)+4)) = v1285
	v1296 = v1122 + int32(4)
	if base.Ui32(v1296) < base.Ui32(v1152+v1153<<(uint(int32(2))%32)) {
		goto L345
	} else {
		goto L346
	}
L309:
	;
	if v23 == int32(0) {
		v1281 = int32(1)
		goto L312
	} else {
		goto L313
	}
L310:
	;
	v1285 = v1163
	goto L311
L311:
	;
	goto L308
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1138)+16)) = v1281
	v1285 = v1281
	goto L311
L313:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v1170 <= int32(0) {
		v1281 = int32(1)
		goto L312
	} else {
		goto L314
	}
L314:
	;
	v1173 = int32(0)
	if v1173 < v1170 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1176 = v1170
	goto L317
L316:
	;
	v1176 = v1173
	goto L317
L317:
	;
	v1178 = v1176 & int32(3)
	v1179 = int32(0)
	if int32(4) <= v1170 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1188 = v1179
	v1189 = v1154
	v1190 = int32(0)
	goto L321
L319:
	;
	v1223 = v1179
	v1224 = v1154
	goto L320
L320:
	;
	if v1178 != 0 {
		goto L336
	} else {
		goto L337
	}
L321:
	;
	v1199 = v1184 + v1189<<(uint(int32(2))%32)
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+12))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+16))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+8))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+16))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+4))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+16))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1199)))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+16))
	if base.Ui32(v1188) < base.Ui32(v1207) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1223 = v1215
	v1224 = v1217
	goto L320
L323:
	;
	v1209 = v1207
	goto L325
L324:
	;
	v1209 = v1188
	goto L325
L325:
	;
	if base.Ui32(v1209) < base.Ui32(v1205) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1211 = v1205
	goto L328
L327:
	;
	v1211 = v1209
	goto L328
L328:
	;
	if base.Ui32(v1211) < base.Ui32(v1203) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1213 = v1203
	goto L331
L330:
	;
	v1213 = v1211
	goto L331
L331:
	;
	if base.Ui32(v1213) < base.Ui32(v1201) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1215 = v1201
	goto L334
L333:
	;
	v1215 = v1213
	goto L334
L334:
	;
	v1216 = int32(4)
	v1217 = v1189 + v1216
	v1219 = v1190 + v1216
	if v1219 != v1176&int32(2147483644) {
		v1188 = v1215
		v1189 = v1217
		v1190 = v1219
		goto L321
	} else {
		goto L335
	}
L335:
	;
	goto L322
L336:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1234 = int32(0)
	v1236 = v1223
	v1237 = v1224
	goto L339
L337:
	;
	v1259 = v1223
	goto L338
L338:
	;
	v1281 = v1259 + int32(1)
	goto L312
L339:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1232+v1237<<(uint(int32(2))%32))))
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1248)+16))
	if base.Ui32(v1236) < base.Ui32(v1249) {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	v1259 = v1251
	goto L338
L341:
	;
	v1251 = v1249
	goto L343
L342:
	;
	v1251 = v1236
	goto L343
L343:
	;
	v1252 = int32(1)
	v1255 = v1234 + v1252
	if v1255 != v1178 {
		v1234 = v1255
		v1236 = v1251
		v1237 = v1237 + v1252
		goto L339
	} else {
		goto L344
	}
L344:
	;
	goto L340
L345:
	;
	v1302 = v1296
	goto L347
L346:
	;
	v1302 = int32(0)
	goto L347
L347:
	;
	v1304 = v1121 + int32(4)
	if base.Ui32(v1304) < base.Ui32(v1150+v1151<<(uint(int32(2))%32)) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1310 = v1304
	goto L350
L349:
	;
	v1310 = int32(0)
	goto L350
L350:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v1314 = v1310
	v1315 = v1302
	v1317 = v1311
	goto L303
L351:
	;
	goto L300
L352:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
	if v1351 != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1352 = v1348
	goto L355
L354:
	;
	v1352 = int32(0)
	goto L355
L355:
	;
	v1371 = v1352
	goto L285
L356:
	;
	m.G0 = v20 + int32(16)
	return
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1373
	goto L356
L358:
	;
	goto L359
L359:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v1375 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1375
	goto L356
L361:
	;
	goto L362
L362:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v1377 != 0 {
		goto L367
	} else {
		goto L368
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1389
	goto L356
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1377
	goto L356
L365:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+4))
	if v1390 <= v1391 {
		goto L363
	} else {
		goto L376
	}
L366:
	;
	if v1371 != 0 {
		goto L373
	} else {
		goto L374
	}
L367:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1379 != 0 {
		v1389 = v1379
		v1390 = v1378
		goto L365
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1383 != 0 {
		v1389 = v1383
		v1390 = int32(0)
		goto L365
	} else {
		goto L372
	}
L370:
	;
	if int32(0) < v1378 {
		goto L364
	} else {
		goto L371
	}
L371:
	;
	goto L366
L372:
	;
	goto L366
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1371
	goto L356
L374:
	;
	goto L375
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	goto L356
L376:
	;
	goto L364
}
