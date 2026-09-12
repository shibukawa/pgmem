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
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
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
	var v494 int32
	_ = v494
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
	var v516 int32
	_ = v516
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
	var v534 int32
	_ = v534
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v913 int32
	_ = v913
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1042 int32
	_ = v1042
	var v1051 int32
	_ = v1051
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1104 int32
	_ = v1104
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1140 int32
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1184 int32
	_ = v1184
	var v1194 int32
	_ = v1194
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1239 int32
	_ = v1239
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
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1304 int32
	_ = v1304
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1373 int32
	_ = v1373
	var v1378 int32
	_ = v1378
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1395 int32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1497 int32
	_ = v1497
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1541 int32
	_ = v1541
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1562 int32
	_ = v1562
	var v1567 int32
	_ = v1567
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1583 int32
	_ = v1583
	var v1588 int32
	_ = v1588
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1606 int32
	_ = v1606
	var v1611 int32
	_ = v1611
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1683 int32
	_ = v1683
	var v1690 int32
	_ = v1690
	v6 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(1120)
	m.G0 = v31
	v34 = base.B2i32(l3 == int32(4538))
	v36 = base.B2i32(l3 == int32(5080))
	v38 = base.B2i32(l3 == int32(5078))
	v40 = base.B2i32(l3 == int32(5079))
	v42 = base.B2i32(l3 == int32(4537))
	v44 = base.B2i32(l3 == int32(3500))
	v46 = base.B2i32(l3 == int32(2776))
	if l2 <= v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v506 != 0 {
		goto L191
	} else {
		goto L192
	}
L2:
	;
	v505 = v6
	v506 = v6
	v507 = v6
	v508 = v6
	v509 = v6
	v510 = v6
	v511 = v6
	v512 = v6
	v513 = v6
	v514 = v6
	v515 = v6
	v516 = v6
	v517 = v36
	v518 = v38
	v519 = v34
	v520 = v40
	v521 = v42
	v522 = v44
	v523 = v46
	goto L1
L3:
	;
	goto L4
L4:
	;
	v57 = v6
	v58 = v6
	v59 = v6
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
	v70 = v36
	v71 = v38
	v72 = v34
	v73 = v40
	v74 = v42
	v75 = v44
	v76 = v46
	goto L5
L5:
	;
	v78 = v57 << (uint(int32(2)) % 32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0+v78)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1+v78)))
	if v82 <= int32(3830) {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v505 = v474
	v506 = v475
	v507 = v476
	v508 = v477
	v509 = v478
	v510 = v479
	v511 = v480
	v512 = v481
	v513 = v482
	v514 = v483
	v515 = v484
	v516 = v485
	v517 = v486
	v518 = v487
	v519 = v488
	v520 = v489
	v521 = v490
	v522 = v491
	v523 = v492
	goto L1
L7:
	;
	v494 = v57 + int32(1)
	if v494 != l2 {
		v57 = v494
		v58 = v474
		v59 = v475
		v60 = v476
		v61 = v477
		v62 = v478
		v63 = v479
		v64 = v480
		v65 = v481
		v66 = v482
		v67 = v483
		v68 = v484
		v69 = v485
		v70 = v486
		v71 = v487
		v72 = v488
		v73 = v489
		v74 = v490
		v75 = v491
		v76 = v492
		goto L5
	} else {
		goto L189
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(704)+v64<<(uint(int32(2))%32)))) = v456
	v469 = int32(1)
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = v469
	v478 = v62
	v479 = v63
	v480 = v64 + v469
	v481 = v65
	v482 = v458
	v483 = v67
	v484 = v459
	v485 = v69
	v486 = v460
	v487 = v461
	v488 = v72
	v489 = v462
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L9:
	;
	v388 = int32(1)
	if v80 == int32(705) {
		goto L160
	} else {
		goto L161
	}
L10:
	;
	v322 = int32(1)
	if v80 == int32(705) {
		goto L131
	} else {
		goto L132
	}
L11:
	;
	v288 = int32(1)
	if v80 == int32(705) {
		goto L115
	} else {
		goto L116
	}
L12:
	;
	v283 = int32(1)
	if v80 == int32(705) {
		v474 = v58
		v475 = v59
		v476 = v60
		v477 = v283
		v478 = v62
		v479 = v63
		v480 = v64
		v481 = v65
		v482 = v66
		v483 = v67
		v484 = v68
		v485 = v69
		v486 = v70
		v487 = v71
		v488 = v72
		v489 = v282
		v490 = v74
		v491 = v75
		v492 = v76
		goto L7
	} else {
		goto L113
	}
L13:
	;
	v282 = int32(1)
	goto L12
L14:
	;
	v236 = int32(1)
	v238 = v59 + v236
	if v80 == int32(705) {
		goto L92
	} else {
		goto L93
	}
L15:
	;
	if v80 == int32(3831) {
		goto L74
	} else {
		goto L75
	}
L16:
	;
	v154 = v59 + int32(1)
	if v80 == int32(705) {
		goto L53
	} else {
		goto L54
	}
L17:
	;
	v111 = v59 + int32(1)
	if v80 == int32(705) {
		goto L32
	} else {
		goto L33
	}
L18:
	;
	if v82 != int32(3500) {
		v108 = v75
		v109 = v76
		goto L17
	} else {
		goto L31
	}
L19:
	;
	switch v82 - int32(2277) {
	case 0:
		goto L16
	case 1, 2, 3, 4, 5:
		v474 = v58
		v475 = v59
		v476 = v60
		v477 = v61
		v478 = v62
		v479 = v63
		v480 = v64
		v481 = v65
		v482 = v66
		v483 = v67
		v484 = v68
		v485 = v69
		v486 = v70
		v487 = v71
		v488 = v72
		v489 = v73
		v490 = v74
		v491 = v75
		v492 = v76
		goto L7
	case 6:
		goto L22
	default:
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	switch v82 - int32(5077) {
	case 0:
		v282 = v73
		goto L12
	case 1:
		goto L11
	case 2:
		goto L13
	case 3:
		goto L10
	default:
		goto L27
	}
L22:
	;
	if v82 != int32(2776) {
		goto L18
	} else {
		goto L26
	}
L23:
	;
	if v82 == int32(2776) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v82 != int32(3500) {
		v474 = v58
		v475 = v59
		v476 = v60
		v477 = v61
		v478 = v62
		v479 = v63
		v480 = v64
		v481 = v65
		v482 = v66
		v483 = v67
		v484 = v68
		v485 = v69
		v486 = v70
		v487 = v71
		v488 = v72
		v489 = v73
		v490 = v74
		v491 = v75
		v492 = v76
		goto L7
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v108 = v75
	v109 = int32(1)
	goto L17
L27:
	;
	switch v82 - int32(4537) {
	case 0:
		goto L14
	case 1:
		goto L9
	default:
		goto L28
	}
L28:
	;
	if v82 != int32(3831) {
		v474 = v58
		v475 = v59
		v476 = v60
		v477 = v61
		v478 = v62
		v479 = v63
		v480 = v64
		v481 = v65
		v482 = v66
		v483 = v67
		v484 = v68
		v485 = v69
		v486 = v70
		v487 = v71
		v488 = v72
		v489 = v73
		v490 = v74
		v491 = v75
		v492 = v76
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v101 = v59 + int32(1)
	if v80 != int32(705) {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v474 = v58
	v475 = v101
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = int32(1)
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L31:
	;
	v108 = int32(1)
	v109 = v76
	goto L17
L32:
	;
	v474 = v58
	v475 = v111
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = int32(1)
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v108
	v492 = v109
	goto L7
L33:
	;
	goto L34
L34:
	;
	if v82 == v80 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v117 = l4
	goto L37
L36:
	;
	v117 = int32(0)
	goto L37
L37:
	;
	if v117 != 0 {
		v474 = v58
		v475 = v111
		v476 = v60
		v477 = v61
		v478 = v62
		v479 = v63
		v480 = v64
		v481 = v65
		v482 = v66
		v483 = v67
		v484 = v68
		v485 = v69
		v486 = v70
		v487 = v71
		v488 = v72
		v489 = v73
		v490 = v74
		v491 = v108
		v492 = v109
		goto L7
	} else {
		goto L38
	}
L38:
	;
	if v58 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v474 = v80
	v475 = v111
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v108
	v492 = v109
	goto L7
L40:
	;
	goto L41
L41:
	;
	if v80 == v58 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v474 = v80
	v475 = v111
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v108
	v492 = v109
	goto L7
L43:
	;
	goto L44
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return int32(0)
L46:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+480)) = int32(102188)
	F_errmsg(m, int32(416183), v31+int32(480))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v137 = F_format_type_be(m, v58)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v139 = F_format_type_be(m, v80)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+468)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v31)+464)) = v137
	F_errdetail(m, int32(189661), v31+int32(464))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(522368), int32(2192), int32(23367))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L45
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
	v474 = v58
	v475 = v154
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = int32(1)
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L54:
	;
	goto L55
L55:
	;
	if v80 == int32(2277) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v161 = l4
	goto L58
L57:
	;
	v161 = int32(0)
	goto L58
L58:
	;
	if v161 != 0 {
		v474 = v58
		v475 = v154
		v476 = v60
		v477 = v61
		v478 = v62
		v479 = v63
		v480 = v64
		v481 = v65
		v482 = v66
		v483 = v67
		v484 = v68
		v485 = v69
		v486 = v70
		v487 = v71
		v488 = v72
		v489 = v73
		v490 = v74
		v491 = v75
		v492 = v76
		goto L7
	} else {
		goto L59
	}
L59:
	;
	v162 = F_getBaseType(m, v80)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L45
	} else {
		goto L60
	}
L60:
	;
	if v65 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v474 = v58
	v475 = v154
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v162
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L62:
	;
	goto L63
L63:
	;
	if v162 == v65 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v474 = v58
	v475 = v154
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v162
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L45
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L45
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+512)) = int32(24402)
	F_errmsg(m, int32(416183), v31+int32(512))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L45
	} else {
		goto L69
	}
L69:
	;
	v181 = F_format_type_be(m, v65)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L45
	} else {
		goto L70
	}
L70:
	;
	v183 = F_format_type_be(m, v162)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L45
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+500)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v31)+496)) = v181
	F_errdetail(m, int32(189661), v31+int32(496))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L45
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(522368), int32(2212), int32(23367))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L45
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
	v200 = l4
	goto L76
L75:
	;
	v200 = int32(0)
	goto L76
L76:
	;
	if v200 != 0 {
		v474 = v58
		v475 = v101
		v476 = v60
		v477 = v61
		v478 = v62
		v479 = v63
		v480 = v64
		v481 = v65
		v482 = v66
		v483 = v67
		v484 = v68
		v485 = v69
		v486 = v70
		v487 = v71
		v488 = v72
		v489 = v73
		v490 = v74
		v491 = v75
		v492 = v76
		goto L7
	} else {
		goto L77
	}
L77:
	;
	v201 = F_getBaseType(m, v80)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L45
	} else {
		goto L78
	}
L78:
	;
	if v60 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v474 = v58
	v475 = v101
	v476 = v201
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L80:
	;
	goto L81
L81:
	;
	if v201 == v60 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v474 = v58
	v475 = v101
	v476 = v201
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L83:
	;
	goto L84
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L45
	} else {
		goto L85
	}
L85:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L45
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = int32(418338)
	F_errmsg(m, int32(416183), v31+int32(544))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L45
	} else {
		goto L87
	}
L87:
	;
	v220 = F_format_type_be(m, v60)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L45
	} else {
		goto L88
	}
L88:
	;
	v222 = F_format_type_be(m, v201)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L45
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v220
	F_errdetail(m, int32(189661), v31+int32(528))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L45
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(522368), int32(2232), int32(23367))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L45
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v474 = v58
	v475 = v238
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = int32(1)
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v236
	v491 = v75
	v492 = v76
	goto L7
L93:
	;
	goto L94
L94:
	;
	if v80 == int32(4537) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v245 = l4
	goto L97
L96:
	;
	v245 = int32(0)
	goto L97
L97:
	;
	if v245 != 0 {
		v474 = v58
		v475 = v238
		v476 = v60
		v477 = v61
		v478 = v62
		v479 = v63
		v480 = v64
		v481 = v65
		v482 = v66
		v483 = v67
		v484 = v68
		v485 = v69
		v486 = v70
		v487 = v71
		v488 = v72
		v489 = v73
		v490 = v236
		v491 = v75
		v492 = v76
		goto L7
	} else {
		goto L98
	}
L98:
	;
	v246 = F_getBaseType(m, v80)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L45
	} else {
		goto L99
	}
L99:
	;
	if v63 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v474 = v58
	v475 = v238
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v246
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v236
	v491 = v75
	v492 = v76
	goto L7
L101:
	;
	goto L102
L102:
	;
	if v246 == v63 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v474 = v58
	v475 = v238
	v476 = v60
	v477 = v61
	v478 = v62
	v479 = v246
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v236
	v491 = v75
	v492 = v76
	goto L7
L104:
	;
	goto L105
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L45
	} else {
		goto L106
	}
L106:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L45
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+576)) = int32(418347)
	F_errmsg(m, int32(416183), v31+int32(576))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L45
	} else {
		goto L108
	}
L108:
	;
	v265 = F_format_type_be(m, v63)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L45
	} else {
		goto L109
	}
L109:
	;
	v267 = F_format_type_be(m, v246)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L45
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+564)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v31)+560)) = v265
	F_errdetail(m, int32(189661), v31+int32(560))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L45
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(522368), int32(2253), int32(23367))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L45
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	if l4&base.B2i32(v82 == v80) != 0 {
		v474 = v58
		v475 = v59
		v476 = v60
		v477 = v283
		v478 = v62
		v479 = v63
		v480 = v64
		v481 = v65
		v482 = v66
		v483 = v67
		v484 = v68
		v485 = v69
		v486 = v70
		v487 = v71
		v488 = v72
		v489 = v282
		v490 = v74
		v491 = v75
		v492 = v76
		goto L7
	} else {
		goto L114
	}
L114:
	;
	v456 = v80
	v458 = v66
	v459 = v68
	v460 = v70
	v461 = v71
	v462 = v282
	goto L8
L115:
	;
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = int32(1)
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v288
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L116:
	;
	goto L117
L117:
	;
	if base.B2i32(v80 == int32(5078))&l4 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = int32(1)
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v288
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L119:
	;
	goto L120
L120:
	;
	v296 = F_getBaseType(m, v80)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L45
	} else {
		goto L121
	}
L121:
	;
	v298 = F_get_element_type(m, v296)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L45
	} else {
		goto L122
	}
L122:
	;
	if v298 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v456 = v298
	v458 = v66
	v459 = v68
	v460 = v70
	v461 = v288
	v462 = v73
	goto L8
L124:
	;
	goto L125
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L45
	} else {
		goto L126
	}
L126:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L45
	} else {
		goto L127
	}
L127:
	;
	v307 = F_format_type_be(m, v296)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L45
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+596)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v31)+592)) = int32(24454)
	F_errmsg(m, int32(198263), v31+int32(592))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L45
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(522368), int32(2286), int32(23367))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L45
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = int32(1)
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v322
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L132:
	;
	goto L133
L133:
	;
	if base.B2i32(v80 == int32(5080))&l4 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = int32(1)
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v322
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L135:
	;
	goto L136
L136:
	;
	v330 = F_getBaseType(m, v80)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L45
	} else {
		goto L137
	}
L137:
	;
	if v66 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	if v330 == v66 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	v364 = F_get_range_subtype(m, v330)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L45
	} else {
		goto L151
	}
L141:
	;
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = int32(1)
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v322
	v487 = v71
	v488 = v72
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L142:
	;
	goto L143
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L45
	} else {
		goto L144
	}
L144:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L45
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+640)) = int32(418949)
	F_errmsg(m, int32(416183), v31+int32(640))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L45
	} else {
		goto L146
	}
L146:
	;
	v348 = F_format_type_be(m, v66)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L45
	} else {
		goto L147
	}
L147:
	;
	v350 = F_format_type_be(m, v330)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L45
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+628)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v31)+624)) = v348
	F_errdetail(m, int32(189661), v31+int32(624))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L45
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(522368), int32(2308), int32(23367))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L45
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	if v364 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v456 = v364
	v458 = v330
	v459 = v364
	v460 = v322
	v461 = v71
	v462 = v73
	goto L8
L153:
	;
	goto L154
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L45
	} else {
		goto L155
	}
L155:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L45
	} else {
		goto L156
	}
L156:
	;
	v373 = F_format_type_be(m, v330)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L45
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+612)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v31)+608)) = int32(418949)
	F_errmsg(m, int32(198370), v31+int32(608))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L45
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(522368), int32(2319), int32(23367))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L45
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = int32(1)
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v388
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L161:
	;
	goto L162
L162:
	;
	if base.B2i32(v80 == int32(4538))&l4 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = int32(1)
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v388
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L164:
	;
	goto L165
L165:
	;
	v396 = F_getBaseType(m, v80)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L45
	} else {
		goto L166
	}
L166:
	;
	if v62 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	if v396 == v62 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L169
L169:
	;
	v430 = F_get_multirange_range(m, v396)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L45
	} else {
		goto L180
	}
L170:
	;
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = int32(1)
	v478 = v62
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v67
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v388
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L171:
	;
	goto L172
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L45
	} else {
		goto L173
	}
L173:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L45
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+688)) = int32(418361)
	F_errmsg(m, int32(416183), v31+int32(688))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L45
	} else {
		goto L175
	}
L175:
	;
	v414 = F_format_type_be(m, v62)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L45
	} else {
		goto L176
	}
L176:
	;
	v416 = F_format_type_be(m, v396)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L45
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+676)) = v416
	*(*int32)(unsafe.Add(mBase, uint32(v31)+672)) = v414
	F_errdetail(m, int32(189661), v31+int32(672))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L45
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(522368), int32(2342), int32(23367))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L45
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
	if v430 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v474 = v58
	v475 = v59
	v476 = v60
	v477 = int32(1)
	v478 = v396
	v479 = v63
	v480 = v64
	v481 = v65
	v482 = v66
	v483 = v430
	v484 = v68
	v485 = v69
	v486 = v70
	v487 = v71
	v488 = v388
	v489 = v73
	v490 = v74
	v491 = v75
	v492 = v76
	goto L7
L182:
	;
	goto L183
L183:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L45
	} else {
		goto L184
	}
L184:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L45
	} else {
		goto L185
	}
L185:
	;
	v440 = F_format_type_be(m, v396)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L45
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+660)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v31)+656)) = int32(418361)
	F_errmsg(m, int32(198312), v31+int32(656))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L45
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(522368), int32(2353), int32(23367))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L45
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	goto L6
L190:
	;
	m.G0 = v31 + int32(1120)
	return v1690
L191:
	;
	if v506 == int32(0) {
		v828 = v505
		v829 = v507
		v830 = v510
		v831 = v512
		goto L205
	} else {
		goto L206
	}
L192:
	;
	if v508&int32(1) != 0 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v1690 = l3
	goto L190
L194:
	;
	if l3 == int32(4537) {
		goto L512
	} else {
		goto L513
	}
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L45
	} else {
		goto L508
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L45
	} else {
		goto L504
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L45
	} else {
		goto L498
	}
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L45
	} else {
		goto L494
	}
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L45
	} else {
		goto L488
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L45
	} else {
		goto L484
	}
L201:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L45
	} else {
		goto L479
	}
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L45
	} else {
		goto L475
	}
L203:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L45
	} else {
		goto L470
	}
L204:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L45
	} else {
		goto L465
	}
L205:
	;
	v832 = int32(0)
	if v508&int32(1) == v832 {
		goto L318
	} else {
		goto L319
	}
L206:
	;
	if v512 == int32(0) {
		v573 = v505
		goto L212
	} else {
		goto L213
	}
L207:
	;
	v788 = base.B2i32(v784 == int32(2283))
	if (v788|(v523^int32(-1)))&int32(1) != 0 {
		goto L304
	} else {
		goto L305
	}
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L45
	} else {
		goto L299
	}
L209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L45
	} else {
		goto L294
	}
L210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L45
	} else {
		goto L287
	}
L211:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L45
	} else {
		goto L282
	}
L212:
	;
	if v510 != 0 {
		goto L240
	} else {
		goto L241
	}
L213:
	;
	if v512 == int32(2277) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	if v505 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L215:
	;
	if v506 != int32(1) {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	goto L217
L217:
	;
	v564 = F_get_element_type(m, v512)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L45
	} else {
		goto L232
	}
L218:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L45
	} else {
		goto L228
	}
L219:
	;
	v534 = int32(2283)
	if l3 <= int32(3499) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	if l3 == int32(2283) {
		goto L218
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	if l3 == int32(3500) {
		goto L218
	} else {
		goto L225
	}
L223:
	;
	if l3 == int32(2776) {
		goto L218
	} else {
		goto L224
	}
L224:
	;
	v568 = v534
	goto L214
L225:
	;
	if l3 == int32(4537) {
		goto L218
	} else {
		goto L226
	}
L226:
	;
	if l3 != int32(3831) {
		v568 = v534
		goto L214
	} else {
		goto L227
	}
L227:
	;
	goto L218
L228:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L45
	} else {
		goto L229
	}
L229:
	;
	F_errmsg(m, int32(100804), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L45
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(522368), int32(2388), int32(23367))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L45
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	if v564 == int32(0) {
		goto L211
	} else {
		goto L233
	}
L233:
	;
	v568 = v564
	goto L214
L234:
	;
	v573 = v568
	goto L212
L235:
	;
	goto L236
L236:
	;
	if v568 != v505 {
		goto L210
	} else {
		goto L237
	}
L237:
	;
	v573 = v505
	goto L212
L238:
	;
	if v573 != 0 {
		v784 = v573
		v785 = int32(0)
		v786 = v613
		goto L207
	} else {
		goto L274
	}
L239:
	;
	v624 = F_get_range_subtype(m, v622)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L45
	} else {
		goto L261
	}
L240:
	;
	v574 = F_get_multirange_range(m, v510)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L45
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v613 = int32(0)
	if base.B2i32(v507 != v613)&v521 != 0 {
		goto L256
	} else {
		goto L257
	}
L243:
	;
	if v574 == int32(0) {
		goto L209
	} else {
		goto L244
	}
L244:
	;
	if v507 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v622 = v574
	v623 = v510
	goto L239
L246:
	;
	goto L247
L247:
	;
	if v574 == v507 {
		v622 = v507
		v623 = v510
		goto L239
	} else {
		goto L248
	}
L248:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L45
	} else {
		goto L249
	}
L249:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L45
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+420)) = int32(418338)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+416)) = int32(418347)
	F_errmsg(m, int32(206708), v31+int32(416))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L45
	} else {
		goto L251
	}
L251:
	;
	v597 = F_format_type_be(m, v510)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L45
	} else {
		goto L252
	}
L252:
	;
	v599 = F_format_type_be(m, v507)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L45
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+404)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v31)+400)) = v597
	F_errdetail(m, int32(189661), v31+int32(400))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L45
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(522368), int32(2449), int32(23367))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L45
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
	v617 = F_get_range_multirange(m, v507)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L45
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	if v507 == int32(0) {
		goto L238
	} else {
		goto L260
	}
L259:
	;
	v622 = v507
	v623 = v617
	goto L239
L260:
	;
	v622 = v507
	v623 = v613
	goto L239
L261:
	;
	if v624 == int32(0) {
		goto L208
	} else {
		goto L262
	}
L262:
	;
	if v573 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v784 = v624
	v785 = v622
	v786 = v623
	goto L207
L264:
	;
	goto L265
L265:
	;
	if v624 == v573 {
		v784 = v573
		v785 = v622
		v786 = v623
		goto L207
	} else {
		goto L266
	}
L266:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L45
	} else {
		goto L267
	}
L267:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L45
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+356)) = int32(102188)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+352)) = int32(418338)
	F_errmsg(m, int32(206708), v31+int32(352))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L45
	} else {
		goto L269
	}
L269:
	;
	v647 = F_format_type_be(m, v622)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L45
	} else {
		goto L270
	}
L270:
	;
	v649 = F_format_type_be(m, v573)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L45
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+340)) = v649
	*(*int32)(unsafe.Add(mBase, uint32(v31)+336)) = v647
	F_errdetail(m, int32(189661), v31+int32(336))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L45
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(522368), int32(2488), int32(23367))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L45
	} else {
		goto L273
	}
L273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L274:
	;
	if l4 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v828 = int32(2283)
	v829 = int32(3831)
	v830 = int32(4537)
	v831 = int32(2277)
	goto L205
L276:
	;
	goto L277
L277:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L45
	} else {
		goto L278
	}
L278:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L45
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+368)) = int32(255208)
	F_errmsg(m, int32(199038), v31+int32(368))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L45
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(522368), int32(2510), int32(23367))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L45
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L45
	} else {
		goto L283
	}
L283:
	;
	v694 = F_format_type_be(m, v512)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L45
	} else {
		goto L284
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v694
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(24402)
	F_errmsg(m, int32(198263), v31)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L45
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(522368), int32(2398), int32(23367))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L45
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L287:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L45
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+452)) = int32(102188)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+448)) = int32(24402)
	F_errmsg(m, int32(206708), v31+int32(448))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L45
	} else {
		goto L289
	}
L289:
	;
	v723 = F_format_type_be(m, v512)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L45
	} else {
		goto L290
	}
L290:
	;
	v725 = F_format_type_be(m, v505)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L45
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+436)) = v725
	*(*int32)(unsafe.Add(mBase, uint32(v31)+432)) = v723
	F_errdetail(m, int32(189661), v31+int32(432))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L45
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(522368), int32(2418), int32(23367))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L45
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L45
	} else {
		goto L295
	}
L295:
	;
	v746 = F_format_type_be(m, v510)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L45
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+388)) = v746
	*(*int32)(unsafe.Add(mBase, uint32(v31)+384)) = int32(418347)
	F_errmsg(m, int32(198312), v31+int32(384))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L45
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(522368), int32(2433), int32(23367))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L45
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L45
	} else {
		goto L300
	}
L300:
	;
	v768 = F_format_type_be(m, v622)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L45
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = int32(418338)
	F_errmsg(m, int32(198370), v31+int32(16))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L45
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(522368), int32(2469), int32(23367))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L45
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	if (v522^int32(-1)|v788)&int32(1) != 0 {
		v828 = v784
		v829 = v785
		v830 = v786
		v831 = v512
		goto L205
	} else {
		goto L313
	}
L305:
	;
	v794 = F_get_base_element_type(m, v784)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L45
	} else {
		goto L306
	}
L306:
	;
	if v794 == int32(0) {
		goto L304
	} else {
		goto L307
	}
L307:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L45
	} else {
		goto L308
	}
L308:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L45
	} else {
		goto L309
	}
L309:
	;
	v805 = F_format_type_be(m, v784)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L45
	} else {
		goto L310
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+320)) = v805
	F_errmsg(m, int32(213090), v31+int32(320))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L45
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(522368), int32(2524), int32(23367))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L45
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	v823 = F_type_is_enum(m, v784)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L45
	} else {
		goto L314
	}
L314:
	;
	if v823 == int32(0) {
		goto L204
	} else {
		goto L315
	}
L315:
	;
	v828 = v784
	v829 = v785
	v830 = v786
	v831 = v512
	goto L205
L316:
	;
	if l3 <= int32(3499) {
		goto L446
	} else {
		goto L447
	}
L317:
	;
	if (v516^int32(-1)|base.B2i32(l2 <= int32(0)))&int32(1) != 0 {
		v1304 = v1146
		v1310 = v1152
		v1313 = v1155
		v1316 = v831
		v1318 = v1160
		goto L316
	} else {
		goto L407
	}
L318:
	;
	v1146 = int32(0)
	v1152 = v832
	v1155 = v509
	v1160 = v513
	goto L317
L319:
	;
	goto L320
L320:
	;
	if v509 != 0 {
		goto L323
	} else {
		goto L324
	}
L321:
	;
	if int32(0) < v894 {
		goto L342
	} else {
		goto L343
	}
L322:
	;
	v893 = v892
	v894 = v511
	v895 = v513
	v896 = v515
	v897 = v517
	goto L321
L323:
	;
	if v513 != 0 {
		goto L326
	} else {
		goto L327
	}
L324:
	;
	goto L325
L325:
	;
	v884 = int32(0)
	if v519&base.B2i32(v513 != v884) == v884 {
		v892 = v884
		goto L322
	} else {
		goto L339
	}
L326:
	;
	if v513 == v514 {
		v892 = v509
		goto L322
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	v871 = F_get_range_subtype(m, v514)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L45
	} else {
		goto L337
	}
L329:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L45
	} else {
		goto L330
	}
L330:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L45
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+292)) = int32(418949)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+288)) = int32(418361)
	F_errmsg(m, int32(206708), v31+int32(288))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L45
	} else {
		goto L332
	}
L332:
	;
	v855 = F_format_type_be(m, v509)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L45
	} else {
		goto L333
	}
L333:
	;
	v857 = F_format_type_be(m, v513)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L45
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v855
	F_errdetail(m, int32(189661), v31+int32(272))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L45
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(522368), int32(2555), int32(23367))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L45
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	if v871 == int32(0) {
		goto L203
	} else {
		goto L338
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(704)+v511<<(uint(int32(2))%32)))) = v871
	v881 = int32(1)
	v893 = v509
	v894 = v511 + v881
	v895 = v514
	v896 = v871
	v897 = v881
	goto L321
L339:
	;
	v890 = F_get_range_multirange(m, v513)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L45
	} else {
		goto L340
	}
L340:
	;
	v892 = v890
	goto L322
L341:
	;
	if l2 <= int32(0) {
		v1304 = v1002
		v1310 = v1008
		v1313 = v1011
		v1316 = v831
		v1318 = v1016
		goto L316
	} else {
		goto L382
	}
L342:
	;
	v903 = F_select_common_type_from_oids(m, v894, v31+int32(704), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L45
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	if l4 != 0 {
		goto L377
	} else {
		goto L378
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1116)) = v903
	v913 = int32(0)
	goto L346
L346:
	;
	v944 = F_can_coerce_type(m, int32(1), v31+int32(704)+v913<<(uint(int32(2))%32), v31+int32(1116), int32(0))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L45
	} else {
		goto L348
	}
L347:
	;
	if v944 == int32(0) {
		goto L202
	} else {
		goto L353
	}
L348:
	;
	if v944 != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v947 = v913 + int32(1)
	if v947 != v894 {
		v913 = v947
		goto L346
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	goto L347
L352:
	;
	goto L351
L353:
	;
	if v518&int32(1) != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v955 = F_get_array_type(m, v903)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L45
	} else {
		goto L357
	}
L355:
	;
	v959 = int32(0)
	goto L356
L356:
	;
	if v897 != 0 {
		goto L359
	} else {
		goto L360
	}
L357:
	;
	if v955 == int32(0) {
		goto L201
	} else {
		goto L358
	}
L358:
	;
	v959 = v955
	goto L356
L359:
	;
	if v895 == int32(0) {
		goto L200
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	if v519 != 0 {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	if v903 != v896 {
		goto L199
	} else {
		goto L363
	}
L363:
	;
	goto L361
L364:
	;
	if v893 == int32(0) {
		goto L198
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	if v520 == int32(0) {
		v1002 = v959
		v1008 = v903
		v1011 = v893
		v1016 = v895
		goto L341
	} else {
		goto L369
	}
L367:
	;
	if v903 != v896 {
		goto L197
	} else {
		goto L368
	}
L368:
	;
	goto L366
L369:
	;
	v968 = F_get_base_element_type(m, v903)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L45
	} else {
		goto L370
	}
L370:
	;
	if v968 == int32(0) {
		v1002 = v959
		v1008 = v903
		v1011 = v893
		v1016 = v895
		goto L341
	} else {
		goto L371
	}
L371:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L45
	} else {
		goto L372
	}
L372:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L45
	} else {
		goto L373
	}
L373:
	;
	v979 = F_format_type_be(m, v903)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L45
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+176)) = v979
	F_errmsg(m, int32(213139), v31+int32(176))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L45
	} else {
		goto L375
	}
L375:
	;
	F_errfinish(m, int32(522368), int32(2658), int32(23367))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L45
	} else {
		goto L376
	}
L376:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L377:
	;
	v1002 = int32(5078)
	v1008 = int32(5077)
	v1011 = int32(4538)
	v1016 = int32(5080)
	goto L341
L378:
	;
	goto L379
L379:
	;
	if v897 != 0 {
		goto L196
	} else {
		goto L380
	}
L380:
	;
	if v519 != 0 {
		goto L195
	} else {
		goto L381
	}
L381:
	;
	v1002 = int32(1009)
	v1008 = int32(25)
	v1011 = v893
	v1016 = v895
	goto L341
L382:
	;
	v1028 = int32(1)
	v1030 = int32(0)
	if l2 != v1028 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1042 = v1030
	v1051 = int32(0)
	goto L386
L384:
	;
	v1104 = v1030
	goto L385
L385:
	;
	if l2&v1028 == int32(0) {
		v1146 = v1002
		v1152 = v1008
		v1155 = v1011
		v1160 = v1016
		goto L317
	} else {
		goto L401
	}
L386:
	;
	v1066 = l1 + v1042<<(uint(int32(2))%32)
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1066)))
	if v1067&int32(-3) == int32(5077) {
		v1076 = v1008
		goto L389
	} else {
		goto L390
	}
L387:
	;
	v1104 = v1094
	goto L385
L388:
	;
	v1080 = v1066 + int32(4)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1080)))
	if v1081&int32(-3) == int32(5077) {
		v1090 = v1008
		goto L395
	} else {
		goto L396
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1066))) = v1076
	goto L388
L390:
	;
	switch v1067 - int32(5078) {
	case 0:
		v1076 = v1002
		goto L389
	case 1:
		goto L388
	case 2:
		goto L392
	default:
		goto L391
	}
L391:
	;
	if v1067 != int32(4538) {
		goto L388
	} else {
		goto L393
	}
L392:
	;
	v1076 = v1016
	goto L389
L393:
	;
	v1076 = v1011
	goto L389
L394:
	;
	v1093 = int32(2)
	v1094 = v1042 + v1093
	v1096 = v1051 + v1093
	if v1096 != l2&int32(2147483646) {
		v1042 = v1094
		v1051 = v1096
		goto L386
	} else {
		goto L400
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080))) = v1090
	goto L394
L396:
	;
	switch v1081 - int32(5078) {
	case 0:
		v1090 = v1002
		goto L395
	case 1:
		goto L394
	case 2:
		goto L397
	default:
		goto L398
	}
L397:
	;
	v1090 = v1016
	goto L395
L398:
	;
	if v1081 == int32(4538) {
		v1090 = v1011
		goto L395
	} else {
		goto L399
	}
L399:
	;
	goto L394
L400:
	;
	goto L387
L401:
	;
	v1130 = l1 + v1104<<(uint(int32(2))%32)
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)))
	if v1131&int32(-3) == int32(5077) {
		v1140 = v1008
		goto L402
	} else {
		goto L403
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1130))) = v1140
	v1146 = v1002
	v1152 = v1008
	v1155 = v1011
	v1160 = v1016
	goto L317
L403:
	;
	switch v1131 - int32(5078) {
	case 0:
		v1140 = v1002
		goto L402
	case 1:
		v1146 = v1002
		v1152 = v1008
		v1155 = v1011
		v1160 = v1016
		goto L317
	case 2:
		goto L404
	default:
		goto L405
	}
L404:
	;
	v1140 = v1016
	goto L402
L405:
	;
	if v1131 == int32(4538) {
		v1140 = v1011
		goto L402
	} else {
		goto L406
	}
L406:
	;
	v1146 = v1002
	v1152 = v1008
	v1155 = v1011
	v1160 = v1016
	goto L317
L407:
	;
	v1184 = int32(0)
	v1194 = v831
	goto L408
L408:
	;
	v1207 = v1184 << (uint(int32(2)) % 32)
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0+v1207)))
	if v1209 != int32(705) {
		v1296 = v1194
		goto L410
	} else {
		goto L411
	}
L409:
	;
	v1304 = v1146
	v1310 = v1152
	v1313 = v1155
	v1316 = v1296
	v1318 = v1160
	goto L316
L410:
	;
	v1298 = v1184 + int32(1)
	if v1298 != l2 {
		v1184 = v1298
		v1194 = v1296
		goto L408
	} else {
		goto L443
	}
L411:
	;
	v1212 = l1 + v1207
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1212)))
	if v1213 <= int32(3499) {
		goto L416
	} else {
		goto L417
	}
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1212))) = v1290
	v1296 = v1291
	goto L410
L413:
	;
	if v1213 != int32(2776) {
		v1296 = v1194
		goto L410
	} else {
		goto L442
	}
L414:
	;
	if v829 != 0 {
		v1290 = v829
		v1291 = v1194
		goto L412
	} else {
		goto L437
	}
L415:
	;
	if v1194 != 0 {
		v1290 = v1194
		v1291 = v1194
		goto L412
	} else {
		goto L429
	}
L416:
	;
	switch v1213 - int32(2277) {
	case 0:
		goto L415
	case 1, 2, 3, 4, 5:
		v1296 = v1194
		goto L410
	case 6:
		v1290 = v828
		v1291 = v1194
		goto L412
	default:
		goto L413
	}
L417:
	;
	goto L418
L418:
	;
	if v1213 == int32(3500) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1290 = v828
	v1291 = v1194
	goto L412
L420:
	;
	goto L421
L421:
	;
	if v1213 == int32(3831) {
		goto L414
	} else {
		goto L422
	}
L422:
	;
	if v1213 != int32(4537) {
		v1296 = v1194
		goto L410
	} else {
		goto L423
	}
L423:
	;
	if v830 != 0 {
		v1290 = v830
		v1291 = v1194
		goto L412
	} else {
		goto L424
	}
L424:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L45
	} else {
		goto L425
	}
L425:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L45
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = int32(255208)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+160)) = int32(418347)
	F_errmsg(m, int32(198972), v31+int32(160))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L45
	} else {
		goto L427
	}
L427:
	;
	F_errfinish(m, int32(522368), int32(2765), int32(23367))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L45
	} else {
		goto L428
	}
L428:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L429:
	;
	v1245 = F_get_array_type(m, v828)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L45
	} else {
		goto L430
	}
L430:
	;
	if v1245 != 0 {
		v1290 = v1245
		v1291 = v1245
		goto L412
	} else {
		goto L431
	}
L431:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L45
	} else {
		goto L432
	}
L432:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L45
	} else {
		goto L433
	}
L433:
	;
	v1254 = F_format_type_be(m, v828)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L45
	} else {
		goto L434
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1254
	F_errmsg(m, int32(203239), v31+int32(128))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L45
	} else {
		goto L435
	}
L435:
	;
	F_errfinish(m, int32(522368), int32(2741), int32(23367))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L45
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L45
	} else {
		goto L438
	}
L438:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L45
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+148)) = int32(255208)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+144)) = int32(418338)
	F_errmsg(m, int32(198972), v31+int32(144))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L45
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(522368), int32(2753), int32(23367))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L45
	} else {
		goto L441
	}
L441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L442:
	;
	v1290 = v828
	v1291 = v1194
	goto L412
L443:
	;
	goto L409
L444:
	;
	if l3 != int32(2776) {
		goto L194
	} else {
		goto L464
	}
L445:
	;
	if v1316 != 0 {
		v1690 = v1316
		goto L190
	} else {
		goto L456
	}
L446:
	;
	switch l3 - int32(2277) {
	case 0:
		goto L445
	case 1, 2, 3, 4, 5:
		goto L194
	case 6:
		v1690 = v828
		goto L190
	default:
		goto L444
	}
L447:
	;
	goto L448
L448:
	;
	if l3 == int32(3500) {
		v1690 = v828
		goto L190
	} else {
		goto L449
	}
L449:
	;
	if l3 != int32(3831) {
		goto L194
	} else {
		goto L450
	}
L450:
	;
	if v829 != 0 {
		v1690 = v829
		goto L190
	} else {
		goto L451
	}
L451:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L45
	} else {
		goto L452
	}
L452:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L45
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = int32(255208)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = int32(418338)
	F_errmsg_internal(m, int32(198972), v31+int32(112))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L45
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(522368), int32(2801), int32(23367))
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L45
	} else {
		goto L455
	}
L455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L456:
	;
	v1357 = F_get_array_type(m, v828)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L45
	} else {
		goto L457
	}
L457:
	;
	if v1357 != 0 {
		v1690 = v1357
		goto L190
	} else {
		goto L458
	}
L458:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L45
	} else {
		goto L459
	}
L459:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L45
	} else {
		goto L460
	}
L460:
	;
	v1366 = F_format_type_be(m, v828)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L45
	} else {
		goto L461
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1366
	F_errmsg(m, int32(203239), v31+int32(96))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L45
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(522368), int32(2788), int32(23367))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L45
	} else {
		goto L463
	}
L463:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L464:
	;
	v1690 = v828
	goto L190
L465:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L45
	} else {
		goto L466
	}
L466:
	;
	v1388 = F_format_type_be(m, v784)
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L45
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v1388
	F_errmsg(m, int32(213198), v31+int32(304))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L45
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(522368), int32(2534), int32(23367))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L45
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
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L45
	} else {
		goto L471
	}
L471:
	;
	v1408 = F_format_type_be(m, v509)
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L45
	} else {
		goto L472
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v1408
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = int32(418361)
	F_errmsg(m, int32(198312), v31+int32(256))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L45
	} else {
		goto L473
	}
L473:
	;
	F_errfinish(m, int32(522368), int32(2566), int32(23367))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L45
	} else {
		goto L474
	}
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L45
	} else {
		goto L476
	}
L476:
	;
	F_errmsg(m, int32(385508), int32(0))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L45
	} else {
		goto L477
	}
L477:
	;
	F_errfinish(m, int32(522368), int32(2594), int32(23367))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L45
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
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L45
	} else {
		goto L480
	}
L480:
	;
	v1446 = F_format_type_be(m, v903)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L45
	} else {
		goto L481
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v1446
	F_errmsg(m, int32(203239), v31+int32(32))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L45
	} else {
		goto L482
	}
L482:
	;
	F_errfinish(m, int32(522368), int32(2603), int32(23367))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L45
	} else {
		goto L483
	}
L483:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L484:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L45
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(255208)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = int32(418949)
	F_errmsg(m, int32(198972), v31+int32(48))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L45
	} else {
		goto L486
	}
L486:
	;
	F_errfinish(m, int32(522368), int32(2613), int32(23367))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L45
	} else {
		goto L487
	}
L487:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L488:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L45
	} else {
		goto L489
	}
L489:
	;
	v1487 = F_format_type_be(m, v895)
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L45
	} else {
		goto L490
	}
L490:
	;
	v1489 = F_format_type_be(m, v903)
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L45
	} else {
		goto L491
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+212)) = v1489
	*(*int32)(unsafe.Add(mBase, uint32(v31)+208)) = v1487
	F_errmsg(m, int32(202513), v31+int32(208))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L45
	} else {
		goto L492
	}
L492:
	;
	F_errfinish(m, int32(522368), int32(2624), int32(23367))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L45
	} else {
		goto L493
	}
L493:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L494:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L45
	} else {
		goto L495
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+68)) = int32(255208)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = int32(418361)
	F_errmsg(m, int32(198972), v31-int32(-64))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L45
	} else {
		goto L496
	}
L496:
	;
	F_errfinish(m, int32(522368), int32(2634), int32(23367))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L45
	} else {
		goto L497
	}
L497:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L45
	} else {
		goto L499
	}
L499:
	;
	v1531 = F_format_type_be(m, v893)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L45
	} else {
		goto L500
	}
L500:
	;
	v1533 = F_format_type_be(m, v903)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L45
	} else {
		goto L501
	}
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+196)) = v1533
	*(*int32)(unsafe.Add(mBase, uint32(v31)+192)) = v1531
	F_errmsg(m, int32(202444), v31+int32(192))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L45
	} else {
		goto L502
	}
L502:
	;
	F_errfinish(m, int32(522368), int32(2645), int32(23367))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L45
	} else {
		goto L503
	}
L503:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L504:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L45
	} else {
		goto L505
	}
L505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = int32(255208)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = int32(418949)
	F_errmsg(m, int32(198972), v31+int32(224))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L45
	} else {
		goto L506
	}
L506:
	;
	F_errfinish(m, int32(522368), int32(2684), int32(23367))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L45
	} else {
		goto L507
	}
L507:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L508:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L45
	} else {
		goto L509
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = int32(255208)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = int32(418361)
	F_errmsg(m, int32(198972), v31+int32(240))
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L45
	} else {
		goto L510
	}
L510:
	;
	F_errfinish(m, int32(522368), int32(2689), int32(23367))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L45
	} else {
		goto L511
	}
L511:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L512:
	;
	if v830 != 0 {
		v1690 = v830
		goto L190
	} else {
		goto L515
	}
L513:
	;
	goto L514
L514:
	;
	switch l3 - int32(5077) {
	case 0, 2:
		goto L521
	default:
		goto L520
	}
L515:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L45
	} else {
		goto L516
	}
L516:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L45
	} else {
		goto L517
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = int32(255208)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = int32(418347)
	F_errmsg_internal(m, int32(198972), v31+int32(80))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L45
	} else {
		goto L518
	}
L518:
	;
	F_errfinish(m, int32(522368), int32(2813), int32(23367))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L45
	} else {
		goto L519
	}
L519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L520:
	;
	if l3 == int32(5078) {
		goto L527
	} else {
		goto L528
	}
L521:
	;
	if v1310 != 0 {
		v1690 = v1310
		goto L190
	} else {
		goto L522
	}
L522:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L45
	} else {
		goto L523
	}
L523:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L45
	} else {
		goto L524
	}
L524:
	;
	F_errmsg_internal(m, int32(386972), int32(0))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L45
	} else {
		goto L525
	}
L525:
	;
	F_errfinish(m, int32(522368), int32(2825), int32(23367))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L45
	} else {
		goto L526
	}
L526:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L527:
	;
	if v1304 != 0 {
		v1690 = v1304
		goto L190
	} else {
		goto L530
	}
L528:
	;
	goto L529
L529:
	;
	if l3 == int32(5080) {
		goto L535
	} else {
		goto L536
	}
L530:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L45
	} else {
		goto L531
	}
L531:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L45
	} else {
		goto L532
	}
L532:
	;
	F_errmsg_internal(m, int32(383768), int32(0))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L45
	} else {
		goto L533
	}
L533:
	;
	F_errfinish(m, int32(522368), int32(2836), int32(23367))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L45
	} else {
		goto L534
	}
L534:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L535:
	;
	if v1318 != 0 {
		v1690 = v1318
		goto L190
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	if l3 != int32(4538) {
		v1690 = l3
		goto L190
	} else {
		goto L543
	}
L538:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L45
	} else {
		goto L539
	}
L539:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L45
	} else {
		goto L540
	}
L540:
	;
	F_errmsg_internal(m, int32(387211), int32(0))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L45
	} else {
		goto L541
	}
L541:
	;
	F_errfinish(m, int32(522368), int32(2847), int32(23367))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L45
	} else {
		goto L542
	}
L542:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L543:
	;
	if v1313 != 0 {
		v1690 = v1313
		goto L190
	} else {
		goto L544
	}
L544:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L45
	} else {
		goto L545
	}
L545:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L45
	} else {
		goto L546
	}
L546:
	;
	F_errmsg_internal(m, int32(387064), int32(0))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L45
	} else {
		goto L547
	}
L547:
	;
	F_errfinish(m, int32(522368), int32(2858), int32(23367))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L45
	} else {
		goto L548
	}
L548:
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
