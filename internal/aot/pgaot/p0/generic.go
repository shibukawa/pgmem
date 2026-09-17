package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enforce_generic_type_consistency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
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
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
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
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v538 int32
	_ = v538
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v921 int32
	_ = v921
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1033 int32
	_ = v1033
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1122 int32
	_ = v1122
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1143 int32
	_ = v1143
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1162 int32
	_ = v1162
	var v1181 int32
	_ = v1181
	var v1189 int32
	_ = v1189
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1235 int32
	_ = v1235
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
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1311 int32
	_ = v1311
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1391 int32
	_ = v1391
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1413 int32
	_ = v1413
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1493 int32
	_ = v1493
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1514 int32
	_ = v1514
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1579 int32
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1631 int32
	_ = v1631
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1643 int32
	_ = v1643
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1661 int32
	_ = v1661
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1687 int32
	_ = v1687
	v6 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(1120)
	m.G0 = v31
	v34 = base.B2i32(l3 == int32(_a_F_enforce_generic_type_consistency_0))
	v36 = base.B2i32(l3 == int32(_a_F_enforce_generic_type_consistency_1))
	v38 = base.B2i32(l3 == int32(_a_F_enforce_generic_type_consistency_2))
	v40 = base.B2i32(l3 == int32(_a_F_enforce_generic_type_consistency_3))
	v42 = base.B2i32(l3 == int32(_a_F_enforce_generic_type_consistency_4))
	v44 = base.B2i32(l3 == int32(3500))
	v46 = base.B2i32(l3 == int32(2776))
	if v6 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v57 = v6
	v58 = v6
	v59 = v44
	v60 = v6
	v61 = v6
	v62 = v6
	v63 = v6
	v64 = v6
	v65 = v6
	v66 = v6
	v67 = v6
	v68 = v6
	v69 = v6
	v70 = v34
	v71 = v36
	v72 = v6
	v73 = v38
	v74 = v40
	v75 = v42
	v76 = v46
	goto L4
L2:
	;
	v505 = v6
	v506 = v6
	v507 = v44
	v508 = v6
	v509 = v6
	v510 = v6
	v511 = v6
	v512 = v6
	v513 = v6
	v514 = v6
	v515 = v6
	v517 = v6
	v518 = v34
	v519 = v36
	v520 = v6
	v521 = v38
	v522 = v40
	v523 = v42
	v524 = v46
	goto L3
L3:
	;
	if v511|v506&int32(1) == int32(0) {
		goto L187
	} else {
		goto L188
	}
L4:
	;
	v78 = v68 << (uint(int32(2)) % 32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0+v78)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1+v78)))
	if v82 <= int32(3830) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v505 = v475
	v506 = v476
	v507 = v477
	v508 = v478
	v509 = v479
	v510 = v480
	v511 = v481
	v512 = v482
	v513 = v483
	v514 = v484
	v515 = v485
	v517 = v486
	v518 = v487
	v519 = v488
	v520 = v489
	v521 = v490
	v522 = v491
	v523 = v492
	v524 = v493
	goto L3
L6:
	;
	v495 = v68 + int32(1)
	if v495 != l2 {
		v57 = v475
		v58 = v476
		v59 = v477
		v60 = v478
		v61 = v479
		v62 = v480
		v63 = v481
		v64 = v482
		v65 = v483
		v66 = v484
		v67 = v485
		v68 = v495
		v69 = v486
		v70 = v487
		v71 = v488
		v72 = v489
		v73 = v490
		v74 = v491
		v75 = v492
		v76 = v493
		goto L4
	} else {
		goto L185
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(704)+v67<<(uint(int32(2))%32)))) = v456
	v470 = int32(1)
	v475 = v57
	v476 = v470
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v459
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67 + v470
	v486 = v460
	v487 = v70
	v488 = v461
	v489 = v72
	v490 = v462
	v491 = v463
	v492 = v75
	v493 = v76
	goto L6
L8:
	;
	v389 = int32(1)
	if v80 == int32(705) {
		goto L156
	} else {
		goto L157
	}
L9:
	;
	v323 = int32(1)
	if v80 == int32(705) {
		goto L127
	} else {
		goto L128
	}
L10:
	;
	v289 = int32(1)
	if v80 == int32(705) {
		goto L113
	} else {
		goto L114
	}
L11:
	;
	if l4&base.B2i32(v80 == v82)|base.B2i32(v80 == int32(705)) != 0 {
		v475 = v57
		v476 = int32(1)
		v477 = v59
		v478 = v60
		v479 = v61
		v480 = v62
		v481 = v63
		v482 = v64
		v483 = v65
		v484 = v66
		v485 = v67
		v486 = v69
		v487 = v70
		v488 = v71
		v489 = v72
		v490 = v73
		v491 = v282
		v492 = v75
		v493 = v76
		goto L6
	} else {
		goto L112
	}
L12:
	;
	v282 = int32(1)
	goto L11
L13:
	;
	v236 = int32(1)
	v238 = v63 + v236
	if v80 == int32(705) {
		goto L91
	} else {
		goto L92
	}
L14:
	;
	if v80 == int32(3831) {
		goto L73
	} else {
		goto L74
	}
L15:
	;
	v154 = v63 + int32(1)
	if v80 == int32(705) {
		goto L52
	} else {
		goto L53
	}
L16:
	;
	v111 = v63 + int32(1)
	if v80 == int32(705) {
		goto L31
	} else {
		goto L32
	}
L17:
	;
	if v82 != int32(3500) {
		v108 = v59
		v109 = v76
		goto L16
	} else {
		goto L30
	}
L18:
	;
	switch v82 - int32(2277) {
	case 0:
		goto L15
	case 1, 2, 3, 4, 5:
		v475 = v57
		v476 = v58
		v477 = v59
		v478 = v60
		v479 = v61
		v480 = v62
		v481 = v63
		v482 = v64
		v483 = v65
		v484 = v66
		v485 = v67
		v486 = v69
		v487 = v70
		v488 = v71
		v489 = v72
		v490 = v73
		v491 = v74
		v492 = v75
		v493 = v76
		goto L6
	case 6:
		goto L21
	default:
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	switch v82 - int32(_a_F_enforce_generic_type_consistency_5) {
	case 0:
		v282 = v74
		goto L11
	case 1:
		goto L10
	case 2:
		goto L12
	case 3:
		goto L9
	default:
		goto L26
	}
L21:
	;
	if v82 != int32(2776) {
		goto L17
	} else {
		goto L25
	}
L22:
	;
	if v82 == int32(2776) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v82 != int32(3500) {
		v475 = v57
		v476 = v58
		v477 = v59
		v478 = v60
		v479 = v61
		v480 = v62
		v481 = v63
		v482 = v64
		v483 = v65
		v484 = v66
		v485 = v67
		v486 = v69
		v487 = v70
		v488 = v71
		v489 = v72
		v490 = v73
		v491 = v74
		v492 = v75
		v493 = v76
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v108 = v59
	v109 = int32(1)
	goto L16
L26:
	;
	switch v82 - int32(_a_F_enforce_generic_type_consistency_4) {
	case 0:
		goto L13
	case 1:
		goto L8
	default:
		goto L27
	}
L27:
	;
	if v82 != int32(3831) {
		v475 = v57
		v476 = v58
		v477 = v59
		v478 = v60
		v479 = v61
		v480 = v62
		v481 = v63
		v482 = v64
		v483 = v65
		v484 = v66
		v485 = v67
		v486 = v69
		v487 = v70
		v488 = v71
		v489 = v72
		v490 = v73
		v491 = v74
		v492 = v75
		v493 = v76
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v101 = v63 + int32(1)
	if v80 != int32(705) {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v475 = v57
	v476 = v58
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v101
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = int32(1)
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L30:
	;
	v108 = int32(1)
	v109 = v76
	goto L16
L31:
	;
	v475 = v57
	v476 = v58
	v477 = v108
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v111
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = int32(1)
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v109
	goto L6
L32:
	;
	goto L33
L33:
	;
	if v80 == v82 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v117 = l4
	goto L36
L35:
	;
	v117 = int32(0)
	goto L36
L36:
	;
	if v117 != 0 {
		v475 = v57
		v476 = v58
		v477 = v108
		v478 = v60
		v479 = v61
		v480 = v62
		v481 = v111
		v482 = v64
		v483 = v65
		v484 = v66
		v485 = v67
		v486 = v69
		v487 = v70
		v488 = v71
		v489 = v72
		v490 = v73
		v491 = v74
		v492 = v75
		v493 = v109
		goto L6
	} else {
		goto L37
	}
L37:
	;
	if v57 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v475 = v80
	v476 = v58
	v477 = v108
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v111
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v109
	goto L6
L39:
	;
	goto L40
L40:
	;
	if v80 == v57 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v475 = v80
	v476 = v58
	v477 = v108
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v111
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v109
	goto L6
L42:
	;
	goto L43
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+480)) = int32(_a_F_enforce_generic_type_consistency_6)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_7), v31+int32(480))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v137 = F_format_type_be(m, v57)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v139 = F_format_type_be(m, v80)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+468)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v31)+464)) = v137
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(464))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2192), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v475 = v57
	v476 = v58
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v154
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = int32(1)
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L53:
	;
	goto L54
L54:
	;
	if v80 == int32(2277) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v161 = l4
	goto L57
L56:
	;
	v161 = int32(0)
	goto L57
L57:
	;
	if v161 != 0 {
		v475 = v57
		v476 = v58
		v477 = v59
		v478 = v60
		v479 = v61
		v480 = v62
		v481 = v154
		v482 = v64
		v483 = v65
		v484 = v66
		v485 = v67
		v486 = v69
		v487 = v70
		v488 = v71
		v489 = v72
		v490 = v73
		v491 = v74
		v492 = v75
		v493 = v76
		goto L6
	} else {
		goto L58
	}
L58:
	;
	v162 = F_getBaseType(m, v80)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L44
	} else {
		goto L59
	}
L59:
	;
	if v64 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v475 = v57
	v476 = v58
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v154
	v482 = v162
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L61:
	;
	goto L62
L62:
	;
	if v162 == v64 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v475 = v57
	v476 = v58
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v154
	v482 = v162
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L64:
	;
	goto L65
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L44
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L44
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+512)) = int32(_a_F_enforce_generic_type_consistency_11)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_7), v31+int32(512))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L44
	} else {
		goto L68
	}
L68:
	;
	v181 = F_format_type_be(m, v64)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L44
	} else {
		goto L69
	}
L69:
	;
	v183 = F_format_type_be(m, v162)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L44
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+500)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v31)+496)) = v181
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(496))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L44
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2212), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L44
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v200 = l4
	goto L75
L74:
	;
	v200 = int32(0)
	goto L75
L75:
	;
	if v200 != 0 {
		v475 = v57
		v476 = v58
		v477 = v59
		v478 = v60
		v479 = v61
		v480 = v62
		v481 = v101
		v482 = v64
		v483 = v65
		v484 = v66
		v485 = v67
		v486 = v69
		v487 = v70
		v488 = v71
		v489 = v72
		v490 = v73
		v491 = v74
		v492 = v75
		v493 = v76
		goto L6
	} else {
		goto L76
	}
L76:
	;
	v201 = F_getBaseType(m, v80)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L44
	} else {
		goto L77
	}
L77:
	;
	if v60 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v475 = v57
	v476 = v58
	v477 = v59
	v478 = v201
	v479 = v61
	v480 = v62
	v481 = v101
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L79:
	;
	goto L80
L80:
	;
	if v201 == v60 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v475 = v57
	v476 = v58
	v477 = v59
	v478 = v201
	v479 = v61
	v480 = v62
	v481 = v101
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L82:
	;
	goto L83
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L44
	} else {
		goto L84
	}
L84:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L44
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = int32(_a_F_enforce_generic_type_consistency_12)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_7), v31+int32(544))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L44
	} else {
		goto L86
	}
L86:
	;
	v220 = F_format_type_be(m, v60)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L44
	} else {
		goto L87
	}
L87:
	;
	v222 = F_format_type_be(m, v201)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L44
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v220
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(528))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L44
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2232), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L44
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v475 = v57
	v476 = v58
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v238
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = int32(1)
	v490 = v73
	v491 = v74
	v492 = v236
	v493 = v76
	goto L6
L92:
	;
	goto L93
L93:
	;
	if v80 == int32(_a_F_enforce_generic_type_consistency_4) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v245 = l4
	goto L96
L95:
	;
	v245 = int32(0)
	goto L96
L96:
	;
	if v245 != 0 {
		v475 = v57
		v476 = v58
		v477 = v59
		v478 = v60
		v479 = v61
		v480 = v62
		v481 = v238
		v482 = v64
		v483 = v65
		v484 = v66
		v485 = v67
		v486 = v69
		v487 = v70
		v488 = v71
		v489 = v72
		v490 = v73
		v491 = v74
		v492 = v236
		v493 = v76
		goto L6
	} else {
		goto L97
	}
L97:
	;
	v246 = F_getBaseType(m, v80)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L44
	} else {
		goto L98
	}
L98:
	;
	if v65 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v475 = v57
	v476 = v58
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v238
	v482 = v64
	v483 = v246
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v236
	v493 = v76
	goto L6
L100:
	;
	goto L101
L101:
	;
	if v246 == v65 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v475 = v57
	v476 = v58
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v238
	v482 = v64
	v483 = v246
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v236
	v493 = v76
	goto L6
L103:
	;
	goto L104
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L44
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L44
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+576)) = int32(_a_F_enforce_generic_type_consistency_13)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_7), v31+int32(576))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L44
	} else {
		goto L107
	}
L107:
	;
	v265 = F_format_type_be(m, v65)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L44
	} else {
		goto L108
	}
L108:
	;
	v267 = F_format_type_be(m, v246)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L44
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+564)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v31)+560)) = v265
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(560))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L44
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2253), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L44
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	v456 = v80
	v459 = v62
	v460 = v69
	v461 = v71
	v462 = v73
	v463 = v282
	goto L7
L113:
	;
	v475 = v57
	v476 = int32(1)
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v289
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L114:
	;
	goto L115
L115:
	;
	if base.B2i32(v80 == int32(_a_F_enforce_generic_type_consistency_2))&l4 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v475 = v57
	v476 = int32(1)
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v71
	v489 = v72
	v490 = v289
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L117:
	;
	goto L118
L118:
	;
	v297 = F_getBaseType(m, v80)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L44
	} else {
		goto L119
	}
L119:
	;
	v299 = F_get_element_type(m, v297)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L44
	} else {
		goto L120
	}
L120:
	;
	if v299 != 0 {
		v456 = v299
		v459 = v62
		v460 = v69
		v461 = v71
		v462 = v289
		v463 = v74
		goto L7
	} else {
		goto L121
	}
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L44
	} else {
		goto L122
	}
L122:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L44
	} else {
		goto L123
	}
L123:
	;
	v308 = F_format_type_be(m, v297)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L44
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+596)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+592)) = int32(_a_F_enforce_generic_type_consistency_14)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_15), v31+int32(592))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L44
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2286), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L44
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	v475 = v57
	v476 = int32(1)
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v323
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L128:
	;
	goto L129
L129:
	;
	if base.B2i32(v80 == int32(_a_F_enforce_generic_type_consistency_1))&l4 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v475 = v57
	v476 = int32(1)
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v323
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L131:
	;
	goto L132
L132:
	;
	v331 = F_getBaseType(m, v80)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L44
	} else {
		goto L133
	}
L133:
	;
	if v62 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	if v331 == v62 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v365 = F_get_range_subtype(m, v331)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L44
	} else {
		goto L147
	}
L137:
	;
	v475 = v57
	v476 = int32(1)
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v70
	v488 = v323
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L138:
	;
	goto L139
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L44
	} else {
		goto L140
	}
L140:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L44
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+640)) = int32(_a_F_enforce_generic_type_consistency_16)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_7), v31+int32(640))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L44
	} else {
		goto L142
	}
L142:
	;
	v349 = F_format_type_be(m, v62)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L44
	} else {
		goto L143
	}
L143:
	;
	v351 = F_format_type_be(m, v331)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L44
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+628)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v31)+624)) = v349
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(624))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L44
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2308), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L44
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	if v365 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v456 = v365
	v459 = v331
	v460 = v365
	v461 = v323
	v462 = v73
	v463 = v74
	goto L7
L149:
	;
	goto L150
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L44
	} else {
		goto L151
	}
L151:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L44
	} else {
		goto L152
	}
L152:
	;
	v374 = F_format_type_be(m, v331)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L44
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+612)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v31)+608)) = int32(_a_F_enforce_generic_type_consistency_16)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_17), v31+int32(608))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L44
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2319), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L44
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	v475 = v57
	v476 = int32(1)
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v389
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L157:
	;
	goto L158
L158:
	;
	if base.B2i32(v80 == int32(_a_F_enforce_generic_type_consistency_0))&l4 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v475 = v57
	v476 = int32(1)
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v389
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L160:
	;
	goto L161
L161:
	;
	v397 = F_getBaseType(m, v80)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L44
	} else {
		goto L162
	}
L162:
	;
	if v61 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	if v397 == v61 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L165
L165:
	;
	v431 = F_get_multirange_range(m, v397)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L44
	} else {
		goto L176
	}
L166:
	;
	v475 = v57
	v476 = int32(1)
	v477 = v59
	v478 = v60
	v479 = v61
	v480 = v62
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v66
	v485 = v67
	v486 = v69
	v487 = v389
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L167:
	;
	goto L168
L168:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L44
	} else {
		goto L169
	}
L169:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L44
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+688)) = int32(_a_F_enforce_generic_type_consistency_18)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_7), v31+int32(688))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L44
	} else {
		goto L171
	}
L171:
	;
	v415 = F_format_type_be(m, v61)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L44
	} else {
		goto L172
	}
L172:
	;
	v417 = F_format_type_be(m, v397)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L44
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+676)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(v31)+672)) = v415
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(672))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L44
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2342), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L44
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
	if v431 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v475 = v57
	v476 = int32(1)
	v477 = v59
	v478 = v60
	v479 = v397
	v480 = v62
	v481 = v63
	v482 = v64
	v483 = v65
	v484 = v431
	v485 = v67
	v486 = v69
	v487 = v389
	v488 = v71
	v489 = v72
	v490 = v73
	v491 = v74
	v492 = v75
	v493 = v76
	goto L6
L178:
	;
	goto L179
L179:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L44
	} else {
		goto L180
	}
L180:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L44
	} else {
		goto L181
	}
L181:
	;
	v441 = F_format_type_be(m, v397)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L44
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+660)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v31)+656)) = int32(_a_F_enforce_generic_type_consistency_18)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_19), v31+int32(656))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L44
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2353), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L44
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	goto L5
L186:
	;
	m.G0 = v31 + int32(1120)
	return v1687
L187:
	;
	v1687 = l3
	goto L186
L188:
	;
	goto L189
L189:
	;
	if v511 == int32(0) {
		v835 = v505
		v836 = v508
		v837 = v512
		v838 = v513
		goto L201
	} else {
		goto L202
	}
L190:
	;
	if l3 == int32(_a_F_enforce_generic_type_consistency_4) {
		goto L507
	} else {
		goto L508
	}
L191:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L44
	} else {
		goto L503
	}
L192:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L44
	} else {
		goto L499
	}
L193:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L44
	} else {
		goto L493
	}
L194:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L44
	} else {
		goto L489
	}
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L44
	} else {
		goto L483
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L44
	} else {
		goto L479
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L44
	} else {
		goto L474
	}
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L44
	} else {
		goto L470
	}
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L44
	} else {
		goto L465
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L44
	} else {
		goto L460
	}
L201:
	;
	v839 = int32(0)
	if v506&int32(1) == v839 {
		goto L311
	} else {
		goto L312
	}
L202:
	;
	if v512 == int32(0) {
		v579 = v505
		goto L208
	} else {
		goto L209
	}
L203:
	;
	v794 = base.B2i32(v790 == int32(2283))
	if (v794|(v524^int32(-1)))&int32(1) != 0 {
		goto L298
	} else {
		goto L299
	}
L204:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L44
	} else {
		goto L293
	}
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L44
	} else {
		goto L288
	}
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L44
	} else {
		goto L281
	}
L207:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L44
	} else {
		goto L276
	}
L208:
	;
	if v513 != 0 {
		goto L234
	} else {
		goto L235
	}
L209:
	;
	if v512 == int32(2277) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	if v505 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L211:
	;
	if v511 != int32(1) {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	v570 = F_get_element_type(m, v512)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L44
	} else {
		goto L226
	}
L214:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L44
	} else {
		goto L222
	}
L215:
	;
	v538 = int32(2283)
	if l3 <= int32(3499) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	if base.B2i32(l3 == int32(2283))|base.B2i32(l3 == int32(2776)) != 0 {
		goto L214
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	if base.B2i32(l3 == int32(3500))|base.B2i32(l3 == int32(_a_F_enforce_generic_type_consistency_4)) != 0 {
		goto L214
	} else {
		goto L220
	}
L219:
	;
	v574 = v538
	goto L210
L220:
	;
	if l3 != int32(3831) {
		v574 = v538
		goto L210
	} else {
		goto L221
	}
L221:
	;
	goto L214
L222:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L44
	} else {
		goto L223
	}
L223:
	;
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_20), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L44
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2388), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L44
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	if v570 == int32(0) {
		goto L207
	} else {
		goto L227
	}
L227:
	;
	v574 = v570
	goto L210
L228:
	;
	v579 = v574
	goto L208
L229:
	;
	goto L230
L230:
	;
	if v574 != v505 {
		goto L206
	} else {
		goto L231
	}
L231:
	;
	v579 = v505
	goto L208
L232:
	;
	if v579 != 0 {
		v790 = v579
		v791 = int32(0)
		v792 = v619
		goto L203
	} else {
		goto L268
	}
L233:
	;
	v630 = F_get_range_subtype(m, v628)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L44
	} else {
		goto L255
	}
L234:
	;
	v580 = F_get_multirange_range(m, v513)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L44
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v619 = int32(0)
	if base.B2i32(v508 != v619)&v523 != 0 {
		goto L250
	} else {
		goto L251
	}
L237:
	;
	if v580 == int32(0) {
		goto L205
	} else {
		goto L238
	}
L238:
	;
	if v508 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v628 = v580
	v629 = v513
	goto L233
L240:
	;
	goto L241
L241:
	;
	if v580 == v508 {
		v628 = v508
		v629 = v513
		goto L233
	} else {
		goto L242
	}
L242:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L44
	} else {
		goto L243
	}
L243:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L44
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+420)) = int32(_a_F_enforce_generic_type_consistency_12)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+416)) = int32(_a_F_enforce_generic_type_consistency_13)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_21), v31+int32(416))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L44
	} else {
		goto L245
	}
L245:
	;
	v603 = F_format_type_be(m, v513)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L44
	} else {
		goto L246
	}
L246:
	;
	v605 = F_format_type_be(m, v508)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L44
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+404)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+400)) = v603
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(400))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L44
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2449), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L44
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
	v623 = F_get_range_multirange(m, v508)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L44
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	if v508 == int32(0) {
		goto L232
	} else {
		goto L254
	}
L253:
	;
	v628 = v508
	v629 = v623
	goto L233
L254:
	;
	v628 = v508
	v629 = v619
	goto L233
L255:
	;
	if v630 == int32(0) {
		goto L204
	} else {
		goto L256
	}
L256:
	;
	if v579 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v790 = v630
	v791 = v628
	v792 = v629
	goto L203
L258:
	;
	goto L259
L259:
	;
	if v630 == v579 {
		v790 = v579
		v791 = v628
		v792 = v629
		goto L203
	} else {
		goto L260
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L44
	} else {
		goto L261
	}
L261:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L44
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+356)) = int32(_a_F_enforce_generic_type_consistency_6)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+352)) = int32(_a_F_enforce_generic_type_consistency_12)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_21), v31+int32(352))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L44
	} else {
		goto L263
	}
L263:
	;
	v653 = F_format_type_be(m, v628)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L44
	} else {
		goto L264
	}
L264:
	;
	v655 = F_format_type_be(m, v579)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L44
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+340)) = v655
	*(*int32)(unsafe.Add(mBase, uint32(v31)+336)) = v653
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(336))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L44
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2488), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L44
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	if l4 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v835 = int32(2283)
	v836 = int32(3831)
	v837 = int32(2277)
	v838 = int32(_a_F_enforce_generic_type_consistency_4)
	goto L201
L270:
	;
	goto L271
L271:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L44
	} else {
		goto L272
	}
L272:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L44
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+368)) = int32(_a_F_enforce_generic_type_consistency_22)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_23), v31+int32(368))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L44
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2510), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L44
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L44
	} else {
		goto L277
	}
L277:
	;
	v700 = F_format_type_be(m, v512)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L44
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v700
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(_a_F_enforce_generic_type_consistency_11)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_15), v31)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L44
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2398), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L44
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L44
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+452)) = int32(_a_F_enforce_generic_type_consistency_6)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+448)) = int32(_a_F_enforce_generic_type_consistency_11)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_21), v31+int32(448))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L44
	} else {
		goto L283
	}
L283:
	;
	v729 = F_format_type_be(m, v512)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L44
	} else {
		goto L284
	}
L284:
	;
	v731 = F_format_type_be(m, v505)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L44
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+436)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v31)+432)) = v729
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(432))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L44
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2418), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L44
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L44
	} else {
		goto L289
	}
L289:
	;
	v752 = F_format_type_be(m, v513)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L44
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+388)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v31)+384)) = int32(_a_F_enforce_generic_type_consistency_13)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_19), v31+int32(384))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L44
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2433), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L44
	} else {
		goto L292
	}
L292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L293:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L44
	} else {
		goto L294
	}
L294:
	;
	v774 = F_format_type_be(m, v628)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L44
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = int32(_a_F_enforce_generic_type_consistency_12)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_17), v31+int32(16))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L44
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2469), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L44
	} else {
		goto L297
	}
L297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	if (v507^int32(-1)|v794)&int32(1) != 0 {
		v835 = v790
		v836 = v791
		v837 = v512
		v838 = v792
		goto L201
	} else {
		goto L307
	}
L299:
	;
	v800 = F_get_base_element_type(m, v790)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L44
	} else {
		goto L300
	}
L300:
	;
	if v800 == int32(0) {
		goto L298
	} else {
		goto L301
	}
L301:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L44
	} else {
		goto L302
	}
L302:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L44
	} else {
		goto L303
	}
L303:
	;
	v811 = F_format_type_be(m, v790)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L44
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+320)) = v811
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_24), v31+int32(320))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L44
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2524), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L44
	} else {
		goto L306
	}
L306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L307:
	;
	v829 = F_type_is_enum(m, v790)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L44
	} else {
		goto L308
	}
L308:
	;
	if v829 == int32(0) {
		goto L200
	} else {
		goto L309
	}
L309:
	;
	v835 = v790
	v836 = v791
	v837 = v512
	v838 = v792
	goto L201
L310:
	;
	if v520 != 0 {
		goto L400
	} else {
		goto L401
	}
L311:
	;
	v1154 = int32(0)
	v1155 = v839
	v1157 = v509
	v1162 = v510
	goto L310
L312:
	;
	goto L313
L313:
	;
	if v509 != 0 {
		goto L316
	} else {
		goto L317
	}
L314:
	;
	if int32(0) < v902 {
		goto L335
	} else {
		goto L336
	}
L315:
	;
	v900 = v899
	v901 = v510
	v902 = v515
	v903 = v517
	v904 = v519
	goto L314
L316:
	;
	if v510 != 0 {
		goto L319
	} else {
		goto L320
	}
L317:
	;
	goto L318
L318:
	;
	v891 = int32(0)
	if v518&base.B2i32(v510 != v891) == v891 {
		v899 = v891
		goto L315
	} else {
		goto L332
	}
L319:
	;
	if v510 == v514 {
		v899 = v509
		goto L315
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v878 = F_get_range_subtype(m, v514)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L44
	} else {
		goto L330
	}
L322:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L44
	} else {
		goto L323
	}
L323:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L44
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+292)) = int32(_a_F_enforce_generic_type_consistency_16)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+288)) = int32(_a_F_enforce_generic_type_consistency_18)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_21), v31+int32(288))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L44
	} else {
		goto L325
	}
L325:
	;
	v862 = F_format_type_be(m, v509)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L44
	} else {
		goto L326
	}
L326:
	;
	v864 = F_format_type_be(m, v510)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L44
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v864
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v862
	F_errdetail(m, int32(_a_F_enforce_generic_type_consistency_8), v31+int32(272))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L44
	} else {
		goto L328
	}
L328:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2555), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L44
	} else {
		goto L329
	}
L329:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L330:
	;
	if v878 == int32(0) {
		goto L199
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(704)+v515<<(uint(int32(2))%32)))) = v878
	v888 = int32(1)
	v900 = v509
	v901 = v514
	v902 = v515 + v888
	v903 = v878
	v904 = v888
	goto L314
L332:
	;
	v897 = F_get_range_multirange(m, v510)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L44
	} else {
		goto L333
	}
L333:
	;
	v899 = v897
	goto L315
L334:
	;
	if l2 <= int32(0) {
		v1154 = v1012
		v1155 = v1013
		v1157 = v1015
		v1162 = v1020
		goto L310
	} else {
		goto L375
	}
L335:
	;
	v910 = F_select_common_type_from_oids(m, v902, v31+int32(704), int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L44
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	if l4 != 0 {
		goto L370
	} else {
		goto L371
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1116)) = v910
	v921 = int32(0)
	goto L339
L339:
	;
	v951 = F_can_coerce_type(m, int32(1), v31+int32(704)+v921<<(uint(int32(2))%32), v31+int32(1116), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L44
	} else {
		goto L341
	}
L340:
	;
	if v951 == int32(0) {
		goto L198
	} else {
		goto L346
	}
L341:
	;
	if v951 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v954 = v921 + int32(1)
	if v954 != v902 {
		v921 = v954
		goto L339
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	goto L340
L345:
	;
	goto L344
L346:
	;
	if v521 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v960 = F_get_array_type(m, v910)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L44
	} else {
		goto L350
	}
L348:
	;
	v964 = int32(0)
	goto L349
L349:
	;
	if v904 != 0 {
		goto L352
	} else {
		goto L353
	}
L350:
	;
	if v960 == int32(0) {
		goto L197
	} else {
		goto L351
	}
L351:
	;
	v964 = v960
	goto L349
L352:
	;
	if v901 == int32(0) {
		goto L196
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	if v518 != 0 {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	if v910 != v903 {
		goto L195
	} else {
		goto L356
	}
L356:
	;
	goto L354
L357:
	;
	if v900 == int32(0) {
		goto L194
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	if v522 == int32(0) {
		v1012 = v964
		v1013 = v910
		v1015 = v900
		v1020 = v901
		goto L334
	} else {
		goto L362
	}
L360:
	;
	if v910 != v903 {
		goto L193
	} else {
		goto L361
	}
L361:
	;
	goto L359
L362:
	;
	v973 = F_get_base_element_type(m, v910)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L44
	} else {
		goto L363
	}
L363:
	;
	if v973 == int32(0) {
		v1012 = v964
		v1013 = v910
		v1015 = v900
		v1020 = v901
		goto L334
	} else {
		goto L364
	}
L364:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L44
	} else {
		goto L365
	}
L365:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L44
	} else {
		goto L366
	}
L366:
	;
	v984 = F_format_type_be(m, v910)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L44
	} else {
		goto L367
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+176)) = v984
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_25), v31+int32(176))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L44
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2658), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L44
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	v1012 = int32(_a_F_enforce_generic_type_consistency_2)
	v1013 = int32(_a_F_enforce_generic_type_consistency_5)
	v1015 = int32(_a_F_enforce_generic_type_consistency_0)
	v1020 = int32(_a_F_enforce_generic_type_consistency_1)
	goto L334
L371:
	;
	goto L372
L372:
	;
	if v904 != 0 {
		goto L192
	} else {
		goto L373
	}
L373:
	;
	if v518 != 0 {
		goto L191
	} else {
		goto L374
	}
L374:
	;
	v1012 = int32(1009)
	v1013 = int32(25)
	v1015 = v900
	v1020 = v901
	goto L334
L375:
	;
	v1033 = int32(0)
	if l2 != int32(1) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1055 = int32(0)
	v1060 = v1033
	goto L379
L377:
	;
	v1122 = v1033
	goto L378
L378:
	;
	v1133 = l1 + v1122<<(uint(int32(2))%32)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	if v1134&int32(-3) == int32(_a_F_enforce_generic_type_consistency_5) {
		v1143 = v1013
		goto L395
	} else {
		goto L396
	}
L379:
	;
	v1071 = l1 + v1060<<(uint(int32(2))%32)
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1071)))
	if v1072&int32(-3) == int32(_a_F_enforce_generic_type_consistency_5) {
		v1081 = v1013
		goto L382
	} else {
		goto L383
	}
L380:
	;
	if l2&int32(1) == int32(0) {
		v1154 = v1012
		v1155 = v1013
		v1157 = v1015
		v1162 = v1020
		goto L310
	} else {
		goto L394
	}
L381:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+4))
	if v1084&int32(-3) == int32(_a_F_enforce_generic_type_consistency_5) {
		v1093 = v1013
		goto L388
	} else {
		goto L389
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1071))) = v1081
	goto L381
L383:
	;
	switch v1072 - int32(_a_F_enforce_generic_type_consistency_2) {
	case 0:
		v1081 = v1012
		goto L382
	case 1:
		goto L381
	case 2:
		goto L385
	default:
		goto L384
	}
L384:
	;
	if v1072 != int32(_a_F_enforce_generic_type_consistency_0) {
		goto L381
	} else {
		goto L386
	}
L385:
	;
	v1081 = v1020
	goto L382
L386:
	;
	v1081 = v1015
	goto L382
L387:
	;
	v1096 = int32(2)
	v1097 = v1060 + v1096
	v1099 = v1055 + v1096
	if v1099 != l2&int32(2147483646) {
		v1055 = v1099
		v1060 = v1097
		goto L379
	} else {
		goto L393
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1071)+4)) = v1093
	goto L387
L389:
	;
	switch v1084 - int32(_a_F_enforce_generic_type_consistency_2) {
	case 0:
		v1093 = v1012
		goto L388
	case 1:
		goto L387
	case 2:
		goto L390
	default:
		goto L391
	}
L390:
	;
	v1093 = v1020
	goto L388
L391:
	;
	if v1084 == int32(_a_F_enforce_generic_type_consistency_0) {
		v1093 = v1015
		goto L388
	} else {
		goto L392
	}
L392:
	;
	goto L387
L393:
	;
	goto L380
L394:
	;
	v1122 = v1097
	goto L378
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1133))) = v1143
	v1154 = v1012
	v1155 = v1013
	v1157 = v1015
	v1162 = v1020
	goto L310
L396:
	;
	switch v1134 - int32(_a_F_enforce_generic_type_consistency_2) {
	case 0:
		v1143 = v1012
		goto L395
	case 1:
		v1154 = v1012
		v1155 = v1013
		v1157 = v1015
		v1162 = v1020
		goto L310
	case 2:
		goto L397
	default:
		goto L398
	}
L397:
	;
	v1143 = v1020
	goto L395
L398:
	;
	if v1134 == int32(_a_F_enforce_generic_type_consistency_0) {
		v1143 = v1015
		goto L395
	} else {
		goto L399
	}
L399:
	;
	v1154 = v1012
	v1155 = v1013
	v1157 = v1015
	v1162 = v1020
	goto L310
L400:
	;
	v1181 = int32(0)
	v1189 = v837
	goto L403
L401:
	;
	v1311 = v837
	goto L402
L402:
	;
	if l3 <= int32(3499) {
		goto L441
	} else {
		goto L442
	}
L403:
	;
	v1203 = v1181 << (uint(int32(2)) % 32)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0+v1203)))
	if v1205 != int32(705) {
		v1292 = v1189
		goto L405
	} else {
		goto L406
	}
L404:
	;
	v1311 = v1292
	goto L402
L405:
	;
	v1294 = v1181 + int32(1)
	if v1294 != l2 {
		v1181 = v1294
		v1189 = v1292
		goto L403
	} else {
		goto L438
	}
L406:
	;
	v1208 = l1 + v1203
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1208)))
	if v1209 <= int32(3499) {
		goto L411
	} else {
		goto L412
	}
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1208))) = v1286
	v1292 = v1287
	goto L405
L408:
	;
	if v1209 != int32(2776) {
		v1292 = v1189
		goto L405
	} else {
		goto L437
	}
L409:
	;
	if v836 != 0 {
		v1286 = v836
		v1287 = v1189
		goto L407
	} else {
		goto L432
	}
L410:
	;
	if v1189 != 0 {
		v1286 = v1189
		v1287 = v1189
		goto L407
	} else {
		goto L424
	}
L411:
	;
	switch v1209 - int32(2277) {
	case 0:
		goto L410
	case 1, 2, 3, 4, 5:
		v1292 = v1189
		goto L405
	case 6:
		v1286 = v835
		v1287 = v1189
		goto L407
	default:
		goto L408
	}
L412:
	;
	goto L413
L413:
	;
	if v1209 == int32(3500) {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1286 = v835
	v1287 = v1189
	goto L407
L415:
	;
	goto L416
L416:
	;
	if v1209 == int32(3831) {
		goto L409
	} else {
		goto L417
	}
L417:
	;
	if v1209 != int32(_a_F_enforce_generic_type_consistency_4) {
		v1292 = v1189
		goto L405
	} else {
		goto L418
	}
L418:
	;
	if v838 != 0 {
		v1286 = v838
		v1287 = v1189
		goto L407
	} else {
		goto L419
	}
L419:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L44
	} else {
		goto L420
	}
L420:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L44
	} else {
		goto L421
	}
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = int32(_a_F_enforce_generic_type_consistency_22)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+160)) = int32(_a_F_enforce_generic_type_consistency_13)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_26), v31+int32(160))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L44
	} else {
		goto L422
	}
L422:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2765), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L44
	} else {
		goto L423
	}
L423:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L424:
	;
	v1241 = F_get_array_type(m, v835)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L44
	} else {
		goto L425
	}
L425:
	;
	if v1241 != 0 {
		v1286 = v1241
		v1287 = v1241
		goto L407
	} else {
		goto L426
	}
L426:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L44
	} else {
		goto L427
	}
L427:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L44
	} else {
		goto L428
	}
L428:
	;
	v1250 = F_format_type_be(m, v835)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L44
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1250
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_27), v31+int32(128))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L44
	} else {
		goto L430
	}
L430:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2741), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L44
	} else {
		goto L431
	}
L431:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L432:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L44
	} else {
		goto L433
	}
L433:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L44
	} else {
		goto L434
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+148)) = int32(_a_F_enforce_generic_type_consistency_22)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+144)) = int32(_a_F_enforce_generic_type_consistency_12)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_26), v31+int32(144))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L44
	} else {
		goto L435
	}
L435:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2753), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L44
	} else {
		goto L436
	}
L436:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L437:
	;
	v1286 = v835
	v1287 = v1189
	goto L407
L438:
	;
	goto L404
L439:
	;
	if l3 != int32(2776) {
		goto L190
	} else {
		goto L459
	}
L440:
	;
	if v1311 != 0 {
		v1687 = v1311
		goto L186
	} else {
		goto L451
	}
L441:
	;
	switch l3 - int32(2277) {
	case 0:
		goto L440
	case 1, 2, 3, 4, 5:
		goto L190
	case 6:
		v1687 = v835
		goto L186
	default:
		goto L439
	}
L442:
	;
	goto L443
L443:
	;
	if l3 == int32(3500) {
		v1687 = v835
		goto L186
	} else {
		goto L444
	}
L444:
	;
	if l3 != int32(3831) {
		goto L190
	} else {
		goto L445
	}
L445:
	;
	if v836 != 0 {
		v1687 = v836
		goto L186
	} else {
		goto L446
	}
L446:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L44
	} else {
		goto L447
	}
L447:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L44
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = int32(_a_F_enforce_generic_type_consistency_22)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = int32(_a_F_enforce_generic_type_consistency_12)
	F_errmsg_internal(m, int32(_a_F_enforce_generic_type_consistency_26), v31+int32(112))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L44
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2801), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L44
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	v1353 = F_get_array_type(m, v835)
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L44
	} else {
		goto L452
	}
L452:
	;
	if v1353 != 0 {
		v1687 = v1353
		goto L186
	} else {
		goto L453
	}
L453:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L44
	} else {
		goto L454
	}
L454:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L44
	} else {
		goto L455
	}
L455:
	;
	v1362 = F_format_type_be(m, v835)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L44
	} else {
		goto L456
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1362
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_27), v31+int32(96))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L44
	} else {
		goto L457
	}
L457:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2788), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L44
	} else {
		goto L458
	}
L458:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L459:
	;
	v1687 = v835
	goto L186
L460:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L44
	} else {
		goto L461
	}
L461:
	;
	v1384 = F_format_type_be(m, v790)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L44
	} else {
		goto L462
	}
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v1384
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_28), v31+int32(304))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L44
	} else {
		goto L463
	}
L463:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2534), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L44
	} else {
		goto L464
	}
L464:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L465:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L44
	} else {
		goto L466
	}
L466:
	;
	v1404 = F_format_type_be(m, v509)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L44
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = int32(_a_F_enforce_generic_type_consistency_18)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_19), v31+int32(256))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L44
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2566), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L44
	} else {
		goto L469
	}
L469:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L470:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L44
	} else {
		goto L471
	}
L471:
	;
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_29), int32(0))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L44
	} else {
		goto L472
	}
L472:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2594), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L44
	} else {
		goto L473
	}
L473:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L474:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L44
	} else {
		goto L475
	}
L475:
	;
	v1442 = F_format_type_be(m, v910)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L44
	} else {
		goto L476
	}
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v1442
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_27), v31+int32(32))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L44
	} else {
		goto L477
	}
L477:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2603), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L44
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L44
	} else {
		goto L480
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(_a_F_enforce_generic_type_consistency_22)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = int32(_a_F_enforce_generic_type_consistency_16)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_26), v31+int32(48))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L44
	} else {
		goto L481
	}
L481:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2613), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L44
	} else {
		goto L482
	}
L482:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L483:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L44
	} else {
		goto L484
	}
L484:
	;
	v1483 = F_format_type_be(m, v901)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L44
	} else {
		goto L485
	}
L485:
	;
	v1485 = F_format_type_be(m, v910)
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L44
	} else {
		goto L486
	}
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+212)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v31)+208)) = v1483
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_30), v31+int32(208))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L44
	} else {
		goto L487
	}
L487:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2624), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L44
	} else {
		goto L488
	}
L488:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L489:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L44
	} else {
		goto L490
	}
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+68)) = int32(_a_F_enforce_generic_type_consistency_22)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = int32(_a_F_enforce_generic_type_consistency_18)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_26), v31-int32(-64))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L44
	} else {
		goto L491
	}
L491:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2634), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L44
	} else {
		goto L492
	}
L492:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L493:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L44
	} else {
		goto L494
	}
L494:
	;
	v1527 = F_format_type_be(m, v900)
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L44
	} else {
		goto L495
	}
L495:
	;
	v1529 = F_format_type_be(m, v910)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L44
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+196)) = v1529
	*(*int32)(unsafe.Add(mBase, uint32(v31)+192)) = v1527
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_31), v31+int32(192))
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L44
	} else {
		goto L497
	}
L497:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2645), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L44
	} else {
		goto L498
	}
L498:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L499:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L44
	} else {
		goto L500
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = int32(_a_F_enforce_generic_type_consistency_22)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = int32(_a_F_enforce_generic_type_consistency_16)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_26), v31+int32(224))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L44
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2684), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L44
	} else {
		goto L502
	}
L502:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L503:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L44
	} else {
		goto L504
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = int32(_a_F_enforce_generic_type_consistency_22)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = int32(_a_F_enforce_generic_type_consistency_18)
	F_errmsg(m, int32(_a_F_enforce_generic_type_consistency_26), v31+int32(240))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L44
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2689), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L44
	} else {
		goto L506
	}
L506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L507:
	;
	if v838 != 0 {
		v1687 = v838
		goto L186
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	switch l3 - int32(_a_F_enforce_generic_type_consistency_5) {
	case 0, 2:
		goto L516
	default:
		goto L515
	}
L510:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L44
	} else {
		goto L511
	}
L511:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L44
	} else {
		goto L512
	}
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = int32(_a_F_enforce_generic_type_consistency_22)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = int32(_a_F_enforce_generic_type_consistency_13)
	F_errmsg_internal(m, int32(_a_F_enforce_generic_type_consistency_26), v31+int32(80))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L44
	} else {
		goto L513
	}
L513:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2813), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L44
	} else {
		goto L514
	}
L514:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L515:
	;
	if l3 == int32(_a_F_enforce_generic_type_consistency_2) {
		goto L522
	} else {
		goto L523
	}
L516:
	;
	if v1155 != 0 {
		v1687 = v1155
		goto L186
	} else {
		goto L517
	}
L517:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L44
	} else {
		goto L518
	}
L518:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L44
	} else {
		goto L519
	}
L519:
	;
	F_errmsg_internal(m, int32(_a_F_enforce_generic_type_consistency_32), int32(0))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L44
	} else {
		goto L520
	}
L520:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2825), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L44
	} else {
		goto L521
	}
L521:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L522:
	;
	if v1154 != 0 {
		v1687 = v1154
		goto L186
	} else {
		goto L525
	}
L523:
	;
	goto L524
L524:
	;
	if l3 == int32(_a_F_enforce_generic_type_consistency_1) {
		goto L530
	} else {
		goto L531
	}
L525:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L44
	} else {
		goto L526
	}
L526:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L44
	} else {
		goto L527
	}
L527:
	;
	F_errmsg_internal(m, int32(_a_F_enforce_generic_type_consistency_33), int32(0))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L44
	} else {
		goto L528
	}
L528:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2836), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L44
	} else {
		goto L529
	}
L529:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L530:
	;
	if v1162 != 0 {
		v1687 = v1162
		goto L186
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	if l3 != int32(_a_F_enforce_generic_type_consistency_0) {
		v1687 = l3
		goto L186
	} else {
		goto L538
	}
L533:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L44
	} else {
		goto L534
	}
L534:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L44
	} else {
		goto L535
	}
L535:
	;
	F_errmsg_internal(m, int32(_a_F_enforce_generic_type_consistency_34), int32(0))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L44
	} else {
		goto L536
	}
L536:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2847), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L44
	} else {
		goto L537
	}
L537:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L538:
	;
	if v1157 != 0 {
		v1687 = v1157
		goto L186
	} else {
		goto L539
	}
L539:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L44
	} else {
		goto L540
	}
L540:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L44
	} else {
		goto L541
	}
L541:
	;
	F_errmsg_internal(m, int32(_a_F_enforce_generic_type_consistency_35), int32(0))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L44
	} else {
		goto L542
	}
L542:
	;
	F_errfinish(m, int32(_a_F_enforce_generic_type_consistency_9), int32(2858), int32(_a_F_enforce_generic_type_consistency_10))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L44
	} else {
		goto L543
	}
L543:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_generic_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	F_mask_unused_space(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
