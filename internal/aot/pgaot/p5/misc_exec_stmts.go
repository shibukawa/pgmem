package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_exec_stmts(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int64
	_ = v396
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int64
	_ = v425
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
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
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
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
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v597 int32
	_ = v597
	var v605 int32
	_ = v605
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
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
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v749 int32
	_ = v749
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
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v863 int32
	_ = v863
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v914 int64
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1314 int32
	_ = v1314
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1619 int32
	_ = v1619
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1833 int32
	_ = v1833
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1849 int32
	_ = v1849
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2004 int32
	_ = v2004
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2165 int32
	_ = v2165
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2179 int32
	_ = v2179
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2206 int32
	_ = v2206
	var v2209 int32
	_ = v2209
	var v2213 int32
	_ = v2213
	var v2218 int32
	_ = v2218
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2233 int32
	_ = v2233
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2262 int32
	_ = v2262
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2279 int32
	_ = v2279
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2292 int32
	_ = v2292
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2316 int32
	_ = v2316
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2337 int32
	_ = v2337
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2379 int32
	_ = v2379
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2392 int32
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2407 int32
	_ = v2407
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2445 int32
	_ = v2445
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2508 int32
	_ = v2508
	var v2512 int32
	_ = v2512
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2533 int32
	_ = v2533
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2545 int32
	_ = v2545
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2550 int64
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2552 int32
	_ = v2552
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2560 int32
	_ = v2560
	var v2561 int32
	_ = v2561
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2581 int32
	_ = v2581
	var v2583 int32
	_ = v2583
	var v2584 int64
	_ = v2584
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2596 int32
	_ = v2596
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2616 int32
	_ = v2616
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2656 int32
	_ = v2656
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2662 int32
	_ = v2662
	var v2663 int64
	_ = v2663
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2687 int32
	_ = v2687
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2699 int32
	_ = v2699
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2703 int64
	_ = v2703
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2711 int32
	_ = v2711
	var v2713 int32
	_ = v2713
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2731 int32
	_ = v2731
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2740 int32
	_ = v2740
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2754 int32
	_ = v2754
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2772 int32
	_ = v2772
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2782 int32
	_ = v2782
	var v2784 int32
	_ = v2784
	var v2809 int32
	_ = v2809
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2828 int32
	_ = v2828
	var v2831 int32
	_ = v2831
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2846 int32
	_ = v2846
	var v2847 int32
	_ = v2847
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2860 int32
	_ = v2860
	var v2865 int32
	_ = v2865
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2875 int32
	_ = v2875
	var v2877 int32
	_ = v2877
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2882 int32
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2895 int32
	_ = v2895
	var v2900 int32
	_ = v2900
	var v2925 int32
	_ = v2925
	var v2928 int32
	_ = v2928
	var v2935 int32
	_ = v2935
	var v2943 int32
	_ = v2943
	var v2948 int32
	_ = v2948
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2975 int32
	_ = v2975
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3005 int32
	_ = v3005
	var v3008 int32
	_ = v3008
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3036 int32
	_ = v3036
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3042 int32
	_ = v3042
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3048 int32
	_ = v3048
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3062 int32
	_ = v3062
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3066 int32
	_ = v3066
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3086 int32
	_ = v3086
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3103 int32
	_ = v3103
	var v3111 int32
	_ = v3111
	var v3119 int32
	_ = v3119
	var v3127 int32
	_ = v3127
	var v3130 int32
	_ = v3130
	var v3133 int32
	_ = v3133
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3136 int32
	_ = v3136
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3141 int32
	_ = v3141
	var v3147 int32
	_ = v3147
	var v3153 int32
	_ = v3153
	var v3159 int32
	_ = v3159
	var v3162 int32
	_ = v3162
	var v3165 int32
	_ = v3165
	var v3168 int32
	_ = v3168
	var v3171 int32
	_ = v3171
	var v3174 int32
	_ = v3174
	var v3179 int32
	_ = v3179
	var v3181 int32
	_ = v3181
	var v3184 int32
	_ = v3184
	var v3187 int32
	_ = v3187
	var v3189 int32
	_ = v3189
	var v3194 int32
	_ = v3194
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3197 int32
	_ = v3197
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3204 int32
	_ = v3204
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3213 int32
	_ = v3213
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3235 int32
	_ = v3235
	var v3237 int32
	_ = v3237
	var v3238 int32
	_ = v3238
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3255 int32
	_ = v3255
	var v3261 int32
	_ = v3261
	var v3264 int32
	_ = v3264
	var v3270 int32
	_ = v3270
	var v3274 int32
	_ = v3274
	var v3279 int32
	_ = v3279
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3289 int32
	_ = v3289
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3322 int32
	_ = v3322
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3332 int32
	_ = v3332
	var v3335 int32
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3338 int32
	_ = v3338
	var v3339 int32
	_ = v3339
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3342 int64
	_ = v3342
	var v3349 int32
	_ = v3349
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3360 int32
	_ = v3360
	var v3363 int32
	_ = v3363
	var v3367 int32
	_ = v3367
	var v3371 int32
	_ = v3371
	var v3376 int32
	_ = v3376
	var v3380 int32
	_ = v3380
	var v3383 int32
	_ = v3383
	var v3387 int32
	_ = v3387
	var v3392 int32
	_ = v3392
	var v3396 int32
	_ = v3396
	var v3399 int32
	_ = v3399
	var v3403 int32
	_ = v3403
	var v3408 int32
	_ = v3408
	var v3412 int32
	_ = v3412
	var v3413 int32
	_ = v3413
	var v3414 int32
	_ = v3414
	var v3421 int32
	_ = v3421
	var v3426 int32
	_ = v3426
	var v3428 int64
	_ = v3428
	var v3430 int32
	_ = v3430
	var v3434 int32
	_ = v3434
	var v3437 int32
	_ = v3437
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3443 int32
	_ = v3443
	var v3450 int32
	_ = v3450
	var v3454 int32
	_ = v3454
	var v3455 int32
	_ = v3455
	var v3458 int32
	_ = v3458
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3464 int32
	_ = v3464
	var v3467 int32
	_ = v3467
	var v3471 int32
	_ = v3471
	var v3477 int32
	_ = v3477
	var v3482 int32
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3491 int32
	_ = v3491
	var v3492 int32
	_ = v3492
	var v3493 int32
	_ = v3493
	var v3497 int32
	_ = v3497
	var v3500 int32
	_ = v3500
	var v3504 int32
	_ = v3504
	var v3510 int32
	_ = v3510
	var v3515 int32
	_ = v3515
	var v3516 int32
	_ = v3516
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3519 int32
	_ = v3519
	var v3521 int32
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3527 int32
	_ = v3527
	var v3530 int32
	_ = v3530
	var v3532 int32
	_ = v3532
	var v3536 int32
	_ = v3536
	var v3538 int32
	_ = v3538
	var v3540 int32
	_ = v3540
	var v3542 int32
	_ = v3542
	var v3543 int32
	_ = v3543
	var v3546 int32
	_ = v3546
	var v3547 int32
	_ = v3547
	var v3549 int32
	_ = v3549
	var v3550 int32
	_ = v3550
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3555 int32
	_ = v3555
	var v3556 int32
	_ = v3556
	var v3560 int32
	_ = v3560
	var v3561 int32
	_ = v3561
	var v3564 int32
	_ = v3564
	var v3567 int32
	_ = v3567
	var v3572 int32
	_ = v3572
	var v3573 int32
	_ = v3573
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3580 int32
	_ = v3580
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3585 int32
	_ = v3585
	var v3586 int32
	_ = v3586
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3591 int32
	_ = v3591
	var v3592 int32
	_ = v3592
	var v3594 int32
	_ = v3594
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3599 int32
	_ = v3599
	var v3601 int32
	_ = v3601
	var v3603 int32
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3606 int32
	_ = v3606
	var v3607 int32
	_ = v3607
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3627 int32
	_ = v3627
	var v3631 int32
	_ = v3631
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3642 int32
	_ = v3642
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3650 int32
	_ = v3650
	var v3654 int32
	_ = v3654
	var v3656 int32
	_ = v3656
	var v3657 int32
	_ = v3657
	var v3658 int32
	_ = v3658
	var v3659 int32
	_ = v3659
	var v3660 int32
	_ = v3660
	var v3665 int32
	_ = v3665
	var v3667 int32
	_ = v3667
	var v3668 int32
	_ = v3668
	var v3669 int32
	_ = v3669
	var v3670 int32
	_ = v3670
	var v3674 int32
	_ = v3674
	var v3675 int32
	_ = v3675
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3684 int32
	_ = v3684
	var v3688 int32
	_ = v3688
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3698 int32
	_ = v3698
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3702 int32
	_ = v3702
	var v3703 int32
	_ = v3703
	var v3705 int32
	_ = v3705
	var v3706 int32
	_ = v3706
	var v3707 int32
	_ = v3707
	var v3710 int32
	_ = v3710
	var v3711 int32
	_ = v3711
	var v3714 int32
	_ = v3714
	var v3718 int32
	_ = v3718
	var v3723 int32
	_ = v3723
	var v3724 int32
	_ = v3724
	var v3725 int32
	_ = v3725
	var v3726 int32
	_ = v3726
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3734 int32
	_ = v3734
	var v3736 int32
	_ = v3736
	var v3737 int32
	_ = v3737
	var v3739 int32
	_ = v3739
	var v3742 int32
	_ = v3742
	var v3744 int32
	_ = v3744
	var v3746 int32
	_ = v3746
	var v3747 int32
	_ = v3747
	var v3748 int32
	_ = v3748
	var v3752 int32
	_ = v3752
	var v3753 int32
	_ = v3753
	var v3755 int32
	_ = v3755
	var v3757 int32
	_ = v3757
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3760 int32
	_ = v3760
	var v3764 int32
	_ = v3764
	var v3766 int64
	_ = v3766
	var v3770 int32
	_ = v3770
	var v3771 int32
	_ = v3771
	var v3772 int32
	_ = v3772
	var v3773 int32
	_ = v3773
	var v3775 int32
	_ = v3775
	var v3776 int32
	_ = v3776
	var v3778 int32
	_ = v3778
	var v3781 int32
	_ = v3781
	var v3782 int32
	_ = v3782
	var v3784 int32
	_ = v3784
	var v3786 int32
	_ = v3786
	var v3788 int32
	_ = v3788
	var v3790 int32
	_ = v3790
	var v3792 int64
	_ = v3792
	var v3795 int64
	_ = v3795
	var v3797 int32
	_ = v3797
	var v3798 int32
	_ = v3798
	var v3799 int32
	_ = v3799
	var v3803 int32
	_ = v3803
	var v3809 int32
	_ = v3809
	var v3810 int32
	_ = v3810
	var v3811 int32
	_ = v3811
	var v3815 int32
	_ = v3815
	var v3816 int32
	_ = v3816
	var v3819 int32
	_ = v3819
	var v3820 int32
	_ = v3820
	var v3822 int32
	_ = v3822
	var v3823 int32
	_ = v3823
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3830 int32
	_ = v3830
	var v3831 int32
	_ = v3831
	var v3835 int32
	_ = v3835
	var v3837 int32
	_ = v3837
	var v3842 int32
	_ = v3842
	var v3845 int32
	_ = v3845
	var v3846 int32
	_ = v3846
	var v3851 int32
	_ = v3851
	var v3854 int32
	_ = v3854
	var v3859 int32
	_ = v3859
	var v3860 int32
	_ = v3860
	var v3864 int32
	_ = v3864
	var v3869 int32
	_ = v3869
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3876 int32
	_ = v3876
	var v3880 int32
	_ = v3880
	var v3885 int32
	_ = v3885
	var v3889 int32
	_ = v3889
	var v3893 int32
	_ = v3893
	var v3898 int32
	_ = v3898
	var v3902 int32
	_ = v3902
	var v3903 int32
	_ = v3903
	var v3909 int32
	_ = v3909
	var v3914 int32
	_ = v3914
	var v3918 int32
	_ = v3918
	var v3921 int32
	_ = v3921
	var v3922 int32
	_ = v3922
	var v3926 int32
	_ = v3926
	var v3932 int32
	_ = v3932
	var v3937 int32
	_ = v3937
	var v3941 int32
	_ = v3941
	var v3942 int32
	_ = v3942
	var v3943 int32
	_ = v3943
	var v3944 int32
	_ = v3944
	var v3951 int32
	_ = v3951
	var v3956 int32
	_ = v3956
	var v3960 int32
	_ = v3960
	var v3964 int32
	_ = v3964
	var v3969 int32
	_ = v3969
	var v3973 int32
	_ = v3973
	var v3977 int32
	_ = v3977
	var v3982 int32
	_ = v3982
	var v3986 int32
	_ = v3986
	var v3989 int32
	_ = v3989
	var v3993 int32
	_ = v3993
	var v3998 int32
	_ = v3998
	var v4002 int32
	_ = v4002
	var v4005 int32
	_ = v4005
	var v4009 int32
	_ = v4009
	var v4014 int32
	_ = v4014
	var v4018 int32
	_ = v4018
	var v4021 int32
	_ = v4021
	var v4025 int32
	_ = v4025
	var v4030 int32
	_ = v4030
	var v4034 int32
	_ = v4034
	var v4037 int32
	_ = v4037
	var v4041 int32
	_ = v4041
	var v4046 int32
	_ = v4046
	var v4050 int32
	_ = v4050
	var v4053 int32
	_ = v4053
	var v4057 int32
	_ = v4057
	var v4062 int32
	_ = v4062
	var v4066 int32
	_ = v4066
	var v4069 int32
	_ = v4069
	var v4075 int32
	_ = v4075
	var v4080 int32
	_ = v4080
	var v4084 int32
	_ = v4084
	var v4087 int32
	_ = v4087
	var v4091 int32
	_ = v4091
	var v4096 int32
	_ = v4096
	var v4100 int32
	_ = v4100
	var v4103 int32
	_ = v4103
	var v4107 int32
	_ = v4107
	var v4112 int32
	_ = v4112
	var v4116 int32
	_ = v4116
	var v4118 int32
	_ = v4118
	var v4119 int32
	_ = v4119
	var v4120 int32
	_ = v4120
	var v4126 int32
	_ = v4126
	var v4131 int32
	_ = v4131
	var v4135 int32
	_ = v4135
	var v4138 int32
	_ = v4138
	var v4142 int32
	_ = v4142
	var v4147 int32
	_ = v4147
	var v4151 int32
	_ = v4151
	var v4154 int32
	_ = v4154
	var v4155 int32
	_ = v4155
	var v4156 int32
	_ = v4156
	var v4162 int32
	_ = v4162
	var v4167 int32
	_ = v4167
	var v4171 int32
	_ = v4171
	var v4174 int32
	_ = v4174
	var v4175 int32
	_ = v4175
	var v4176 int32
	_ = v4176
	var v4183 int32
	_ = v4183
	var v4188 int32
	_ = v4188
	var v4192 int32
	_ = v4192
	var v4195 int32
	_ = v4195
	var v4199 int32
	_ = v4199
	var v4204 int32
	_ = v4204
	var v4208 int32
	_ = v4208
	var v4211 int32
	_ = v4211
	var v4215 int32
	_ = v4215
	var v4220 int32
	_ = v4220
	var v4224 int32
	_ = v4224
	var v4227 int32
	_ = v4227
	var v4231 int32
	_ = v4231
	var v4236 int32
	_ = v4236
	var v4240 int32
	_ = v4240
	var v4243 int32
	_ = v4243
	var v4247 int32
	_ = v4247
	var v4252 int32
	_ = v4252
	var v4256 int32
	_ = v4256
	var v4259 int32
	_ = v4259
	var v4263 int32
	_ = v4263
	var v4268 int32
	_ = v4268
	var v4272 int32
	_ = v4272
	var v4275 int32
	_ = v4275
	var v4279 int32
	_ = v4279
	var v4284 int32
	_ = v4284
	var v4288 int32
	_ = v4288
	var v4291 int32
	_ = v4291
	var v4295 int32
	_ = v4295
	var v4300 int32
	_ = v4300
	var v4304 int32
	_ = v4304
	var v4307 int32
	_ = v4307
	var v4311 int32
	_ = v4311
	var v4316 int32
	_ = v4316
	var v4320 int32
	_ = v4320
	var v4323 int32
	_ = v4323
	var v4327 int32
	_ = v4327
	var v4332 int32
	_ = v4332
	var v4336 int32
	_ = v4336
	var v4339 int32
	_ = v4339
	var v4343 int32
	_ = v4343
	var v4348 int32
	_ = v4348
	var v4352 int32
	_ = v4352
	var v4353 int32
	_ = v4353
	var v4354 int32
	_ = v4354
	var v4361 int32
	_ = v4361
	var v4366 int32
	_ = v4366
	var v4368 int32
	_ = v4368
	var v4372 int32
	_ = v4372
	var v4376 int32
	_ = v4376
	var v4381 int32
	_ = v4381
	var v4385 int32
	_ = v4385
	var v4389 int32
	_ = v4389
	var v4394 int32
	_ = v4394
	var v4398 int32
	_ = v4398
	var v4401 int32
	_ = v4401
	var v4405 int32
	_ = v4405
	var v4410 int32
	_ = v4410
	var v4414 int32
	_ = v4414
	var v4417 int32
	_ = v4417
	var v4424 int32
	_ = v4424
	var v4429 int32
	_ = v4429
	var v4433 int32
	_ = v4433
	var v4436 int32
	_ = v4436
	var v4443 int32
	_ = v4443
	var v4448 int32
	_ = v4448
	var v4452 int32
	_ = v4452
	var v4455 int32
	_ = v4455
	var v4462 int32
	_ = v4462
	var v4467 int32
	_ = v4467
	var v4471 int32
	_ = v4471
	var v4474 int32
	_ = v4474
	var v4481 int32
	_ = v4481
	var v4486 int32
	_ = v4486
	var v4490 int32
	_ = v4490
	var v4493 int32
	_ = v4493
	var v4500 int32
	_ = v4500
	var v4505 int32
	_ = v4505
	var v4509 int32
	_ = v4509
	var v4512 int32
	_ = v4512
	var v4519 int32
	_ = v4519
	var v4524 int32
	_ = v4524
	var v4528 int32
	_ = v4528
	var v4531 int32
	_ = v4531
	var v4538 int32
	_ = v4538
	var v4543 int32
	_ = v4543
	var v4547 int32
	_ = v4547
	var v4550 int32
	_ = v4550
	var v4557 int32
	_ = v4557
	var v4562 int32
	_ = v4562
	var v4566 int32
	_ = v4566
	var v4569 int32
	_ = v4569
	var v4576 int32
	_ = v4576
	var v4581 int32
	_ = v4581
	var v4585 int32
	_ = v4585
	var v4586 int32
	_ = v4586
	var v4592 int32
	_ = v4592
	var v4597 int32
	_ = v4597
	var v4601 int32
	_ = v4601
	var v4604 int32
	_ = v4604
	var v4608 int32
	_ = v4608
	var v4613 int32
	_ = v4613
	var v4617 int32
	_ = v4617
	var v4620 int32
	_ = v4620
	var v4624 int32
	_ = v4624
	var v4629 int32
	_ = v4629
	var v4633 int32
	_ = v4633
	var v4636 int32
	_ = v4636
	var v4642 int32
	_ = v4642
	var v4647 int32
	_ = v4647
	var v4651 int32
	_ = v4651
	var v4654 int32
	_ = v4654
	var v4658 int32
	_ = v4658
	var v4663 int32
	_ = v4663
	var v4667 int32
	_ = v4667
	var v4670 int32
	_ = v4670
	var v4674 int32
	_ = v4674
	var v4679 int32
	_ = v4679
	var v4683 int32
	_ = v4683
	var v4685 int32
	_ = v4685
	var v4686 int32
	_ = v4686
	var v4687 int32
	_ = v4687
	var v4693 int32
	_ = v4693
	var v4698 int32
	_ = v4698
	var v4702 int32
	_ = v4702
	var v4705 int32
	_ = v4705
	var v4706 int32
	_ = v4706
	var v4712 int32
	_ = v4712
	var v4717 int32
	_ = v4717
	var v4721 int32
	_ = v4721
	var v4724 int32
	_ = v4724
	var v4730 int32
	_ = v4730
	var v4735 int32
	_ = v4735
	var v4739 int32
	_ = v4739
	var v4742 int32
	_ = v4742
	var v4746 int32
	_ = v4746
	var v4751 int32
	_ = v4751
	var v4755 int32
	_ = v4755
	var v4758 int32
	_ = v4758
	var v4759 int32
	_ = v4759
	var v4765 int32
	_ = v4765
	var v4770 int32
	_ = v4770
	var v4774 int32
	_ = v4774
	var v4777 int32
	_ = v4777
	var v4783 int32
	_ = v4783
	var v4788 int32
	_ = v4788
	var v4792 int32
	_ = v4792
	var v4823 int32
	_ = v4823
	var v4827 int32
	_ = v4827
	var v4833 int32
	_ = v4833
	var v4836 int32
	_ = v4836
	var v4839 int32
	_ = v4839
	var v4840 int32
	_ = v4840
	var v4843 int32
	_ = v4843
	var v4844 int32
	_ = v4844
	var v4847 int32
	_ = v4847
	var v4854 int32
	_ = v4854
	var v4855 int32
	_ = v4855
	var v4858 int32
	_ = v4858
	var v4861 int32
	_ = v4861
	var v4865 int32
	_ = v4865
	var v4868 int32
	_ = v4868
	var v4892 int32
	_ = v4892
	var v4893 int32
	_ = v4893
	var v4897 int32
	_ = v4897
	var v4898 int32
	_ = v4898
	var v4901 int32
	_ = v4901
	var v4902 int32
	_ = v4902
	var v4906 int32
	_ = v4906
	var v4907 int32
	_ = v4907
	var v4912 int32
	_ = v4912
	var v4915 int32
	_ = v4915
	var v4918 int32
	_ = v4918
	var v4919 int32
	_ = v4919
	var v4922 int32
	_ = v4922
	var v4923 int32
	_ = v4923
	var v4926 int32
	_ = v4926
	var v4933 int32
	_ = v4933
	var v4934 int32
	_ = v4934
	var v4965 int32
	_ = v4965
	var v4997 int32
	_ = v4997
	var v5000 int32
	_ = v5000
	var v5001 int32
	_ = v5001
	var v5002 int32
	_ = v5002
	var v5005 int32
	_ = v5005
	var v5006 int32
	_ = v5006
	var v5009 int32
	_ = v5009
	var v5013 int32
	_ = v5013
	var v5015 int32
	_ = v5015
	var v5017 int32
	_ = v5017
	var v5019 int32
	_ = v5019
	var v5021 int32
	_ = v5021
	var v5026 int32
	_ = v5026
	var v5060 int32
	_ = v5060
	var v5061 int32
	_ = v5061
	var v5062 int32
	_ = v5062
	var v5063 int32
	_ = v5063
	var v5067 int32
	_ = v5067
	var v5070 int32
	_ = v5070
	var v5074 int32
	_ = v5074
	var v5078 int32
	_ = v5078
	var v5083 int32
	_ = v5083
	var v5144 int32
	_ = v5144
	var v5145 int32
	_ = v5145
	var v5146 int32
	_ = v5146
	var v5147 int32
	_ = v5147
	var v5177 int32
	_ = v5177
	var v5179 int32
	_ = v5179
	var v5180 int32
	_ = v5180
	var v5183 int32
	_ = v5183
	var v5186 int32
	_ = v5186
	var v5188 int32
	_ = v5188
	var v5191 int32
	_ = v5191
	var v5219 int32
	_ = v5219
	var v5220 int32
	_ = v5220
	var v5223 int32
	_ = v5223
	var v5227 int32
	_ = v5227
	var v5229 int32
	_ = v5229
	var v5231 int32
	_ = v5231
	var v5235 int32
	_ = v5235
	var v5236 int32
	_ = v5236
	var v5272 int32
	_ = v5272
	v3 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(640)
	m.G0 = v32
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v32 + int32(640)
	return v5272
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v34
	v5272 = int32(0)
	goto L1
L3:
	;
	v63 = v47
	v79 = v3
	goto L11
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 <= int32(0) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[0]))
	if v49 == int32(0) {
		v5272 = v3
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v39 = l0 + int32(20)
	v41 = l0 + int32(16)
	v43 = l0 + int32(12)
	v45 = v32 + int32(616)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[1]))
	goto L3
L8:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v5272 = v3
	goto L1
L11:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v79<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v91 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v34
	v5272 = v5191
	goto L1
L13:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[0]))
	if v101 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v94 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	m.T0[v94].(func(*base.Module, int32, int32))(m, l0, v89)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v104 {
	case 0:
		goto L89
	case 1:
		goto L116
	case 2:
		goto L112
	case 3:
		goto L111
	case 4:
		goto L110
	case 5:
		goto L109
	case 6:
		goto L108
	case 7:
		goto L107
	case 8:
		goto L106
	case 9:
		goto L105
	case 10:
		goto L104
	case 11:
		goto L103
	case 12:
		goto L102
	case 13:
		goto L101
	case 14:
		goto L100
	case 15:
		goto L99
	case 16:
		goto L98
	case 17:
		goto L97
	case 18:
		goto L96
	case 19:
		goto L113
	case 20:
		goto L95
	case 21:
		goto L94
	case 22:
		goto L93
	case 23:
		goto L115
	case 24:
		goto L114
	case 25:
		goto L92
	case 26:
		goto L91
	default:
		goto L90
	}
L20:
	;
	goto L19
L21:
	;
	v5219 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[1]))
	v5220 = *(*int32)(unsafe.Add(mBase, uint32(v5219)))
	if v5220 == int32(0) {
		v5231 = v5219
		goto L1504
	} else {
		goto L1505
	}
L22:
	;
	v5177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v5177 != 0 {
		goto L1498
	} else {
		goto L1499
	}
L23:
	;
	v5145 = *(*int32)(unsafe.Add(mBase, uint32(v5144)))
	v5146 = F_exec_stmts(m, l0, v5145)
	mBase = m.M
	v5147 = m.ExcPending
	if v5147 != 0 {
		goto L9
	} else {
		goto L1497
	}
L24:
	;
	v5144 = v89 + int32(24)
	goto L23
L25:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v5067 = m.ExcPending
	if v5067 != 0 {
		goto L9
	} else {
		goto L1492
	}
L26:
	;
	v5061 = *(*int32)(unsafe.Add(mBase, uint32(v5060)))
	v5062 = F_exec_stmts(m, l0, v5061)
	mBase = m.M
	v5063 = m.ExcPending
	if v5063 != 0 {
		goto L9
	} else {
		goto L1491
	}
L27:
	;
	if v1108 != 0 {
		goto L1478
	} else {
		goto L1479
	}
L28:
	;
	v4965 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v4965
	v5191 = v4965
	goto L21
L29:
	;
	v4902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4902 == int32(0) {
		goto L1466
	} else {
		goto L1467
	}
L30:
	;
	v4892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v4893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v4897 = *(*int32)(unsafe.Add(mBase, uint32(v4892+v4893<<(uint(int32(2))%32))))
	v4898 = int32(0)
	F_assign_simple_var(m, l0, v4897, v4868, v4898, v4898)
	mBase = m.M
	v4901 = m.ExcPending
	if v4901 != 0 {
		goto L9
	} else {
		goto L1465
	}
L31:
	;
	v4861 = int32(0)
	v4865 = v4861
	v4868 = v4861
	goto L30
L32:
	;
	v4823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4823 == int32(0) {
		goto L1449
	} else {
		goto L1450
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
	F_plpgsql_create_econtext(m, l0)
	mBase = m.M
	v4792 = m.ExcPending
	if v4792 != 0 {
		goto L9
	} else {
		goto L1448
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4774 = m.ExcPending
	if v4774 != 0 {
		goto L9
	} else {
		goto L1444
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4755 = m.ExcPending
	if v4755 != 0 {
		goto L9
	} else {
		goto L1440
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4739 = m.ExcPending
	if v4739 != 0 {
		goto L9
	} else {
		goto L1436
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4721 = m.ExcPending
	if v4721 != 0 {
		goto L9
	} else {
		goto L1432
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4702 = m.ExcPending
	if v4702 != 0 {
		goto L9
	} else {
		goto L1428
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4683 = m.ExcPending
	if v4683 != 0 {
		goto L9
	} else {
		goto L1424
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4667 = m.ExcPending
	if v4667 != 0 {
		goto L9
	} else {
		goto L1420
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4651 = m.ExcPending
	if v4651 != 0 {
		goto L9
	} else {
		goto L1416
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4633 = m.ExcPending
	if v4633 != 0 {
		goto L9
	} else {
		goto L1412
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4617 = m.ExcPending
	if v4617 != 0 {
		goto L9
	} else {
		goto L1408
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4601 = m.ExcPending
	if v4601 != 0 {
		goto L9
	} else {
		goto L1404
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4585 = m.ExcPending
	if v4585 != 0 {
		goto L9
	} else {
		goto L1401
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4566 = m.ExcPending
	if v4566 != 0 {
		goto L9
	} else {
		goto L1397
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4547 = m.ExcPending
	if v4547 != 0 {
		goto L9
	} else {
		goto L1393
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4528 = m.ExcPending
	if v4528 != 0 {
		goto L9
	} else {
		goto L1389
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4509 = m.ExcPending
	if v4509 != 0 {
		goto L9
	} else {
		goto L1385
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4490 = m.ExcPending
	if v4490 != 0 {
		goto L9
	} else {
		goto L1381
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4471 = m.ExcPending
	if v4471 != 0 {
		goto L9
	} else {
		goto L1377
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4452 = m.ExcPending
	if v4452 != 0 {
		goto L9
	} else {
		goto L1373
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4433 = m.ExcPending
	if v4433 != 0 {
		goto L9
	} else {
		goto L1369
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4414 = m.ExcPending
	if v4414 != 0 {
		goto L9
	} else {
		goto L1365
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4398 = m.ExcPending
	if v4398 != 0 {
		goto L9
	} else {
		goto L1361
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4385 = m.ExcPending
	if v4385 != 0 {
		goto L9
	} else {
		goto L1358
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4372 = m.ExcPending
	if v4372 != 0 {
		goto L9
	} else {
		goto L1355
	}
L58:
	;
	F_ReThrowError(m, v2720)
	mBase = m.M
	v4368 = m.ExcPending
	if v4368 != 0 {
		goto L9
	} else {
		goto L1354
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4352 = m.ExcPending
	if v4352 != 0 {
		goto L9
	} else {
		goto L1350
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4336 = m.ExcPending
	if v4336 != 0 {
		goto L9
	} else {
		goto L1346
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4320 = m.ExcPending
	if v4320 != 0 {
		goto L9
	} else {
		goto L1342
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4304 = m.ExcPending
	if v4304 != 0 {
		goto L9
	} else {
		goto L1338
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4288 = m.ExcPending
	if v4288 != 0 {
		goto L9
	} else {
		goto L1334
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4272 = m.ExcPending
	if v4272 != 0 {
		goto L9
	} else {
		goto L1330
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4256 = m.ExcPending
	if v4256 != 0 {
		goto L9
	} else {
		goto L1326
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4240 = m.ExcPending
	if v4240 != 0 {
		goto L9
	} else {
		goto L1322
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4224 = m.ExcPending
	if v4224 != 0 {
		goto L9
	} else {
		goto L1318
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4208 = m.ExcPending
	if v4208 != 0 {
		goto L9
	} else {
		goto L1314
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4192 = m.ExcPending
	if v4192 != 0 {
		goto L9
	} else {
		goto L1310
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		goto L9
	} else {
		goto L1306
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4151 = m.ExcPending
	if v4151 != 0 {
		goto L9
	} else {
		goto L1301
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4135 = m.ExcPending
	if v4135 != 0 {
		goto L9
	} else {
		goto L1297
	}
L73:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4116 = m.ExcPending
	if v4116 != 0 {
		goto L9
	} else {
		goto L1293
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4100 = m.ExcPending
	if v4100 != 0 {
		goto L9
	} else {
		goto L1289
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4084 = m.ExcPending
	if v4084 != 0 {
		goto L9
	} else {
		goto L1285
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4066 = m.ExcPending
	if v4066 != 0 {
		goto L9
	} else {
		goto L1281
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4050 = m.ExcPending
	if v4050 != 0 {
		goto L9
	} else {
		goto L1277
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4034 = m.ExcPending
	if v4034 != 0 {
		goto L9
	} else {
		goto L1273
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4018 = m.ExcPending
	if v4018 != 0 {
		goto L9
	} else {
		goto L1269
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4002 = m.ExcPending
	if v4002 != 0 {
		goto L9
	} else {
		goto L1265
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3986 = m.ExcPending
	if v3986 != 0 {
		goto L9
	} else {
		goto L1261
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3973 = m.ExcPending
	if v3973 != 0 {
		goto L9
	} else {
		goto L1258
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3960 = m.ExcPending
	if v3960 != 0 {
		goto L9
	} else {
		goto L1255
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3941 = m.ExcPending
	if v3941 != 0 {
		goto L9
	} else {
		goto L1251
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3918 = m.ExcPending
	if v3918 != 0 {
		goto L9
	} else {
		goto L1247
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3902 = m.ExcPending
	if v3902 != 0 {
		goto L9
	} else {
		goto L1244
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3889 = m.ExcPending
	if v3889 != 0 {
		goto L9
	} else {
		goto L1241
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		goto L9
	} else {
		goto L1238
	}
L89:
	;
	v3870 = F_exec_stmt_block(m, l0, v89)
	mBase = m.M
	v3871 = m.ExcPending
	if v3871 != 0 {
		goto L9
	} else {
		goto L1237
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v34
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3859 = m.ExcPending
	if v3859 != 0 {
		goto L9
	} else {
		goto L1234
	}
L91:
	;
	v3846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v3846 == int32(1) {
		goto L1229
	} else {
		goto L1230
	}
L92:
	;
	v3837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v3837 == int32(1) {
		goto L1223
	} else {
		goto L1224
	}
L93:
	;
	v3810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3811 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3815 = *(*int32)(unsafe.Add(mBase, uint32(v3810+v3811<<(uint(int32(2))%32))))
	v3816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3815)+44)))
	if v3816 == int32(1) {
		goto L35
	} else {
		goto L1217
	}
L94:
	;
	v3689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3690 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v3694 = *(*int32)(unsafe.Add(mBase, uint32(v3689+v3690<<(uint(int32(2))%32))))
	v3695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3694)+44)))
	if v3695 == int32(1) {
		goto L38
	} else {
		goto L1181
	}
L95:
	;
	v3553 = int32(0)
	v3555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3556 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3560 = *(*int32)(unsafe.Add(mBase, uint32(v3555+v3556<<(uint(int32(2))%32))))
	v3561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3560)+44)))
	if v3561 == v3553 {
		goto L1124
	} else {
		goto L1125
	}
L96:
	;
	v3542 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v3543 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v3546 = F_exec_dynquery_with_params(m, l0, v3542, v3543, int32(0), int32(4))
	mBase = m.M
	v3547 = m.ExcPending
	if v3547 != 0 {
		goto L9
	} else {
		goto L1121
	}
L97:
	;
	v3286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v3286 == int32(0) {
		goto L1036
	} else {
		goto L1037
	}
L98:
	;
	F_exec_stmt_execsql(m, l0, v89)
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L9
	} else {
		goto L1035
	}
L99:
	;
	v3184 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[2])))
	if v3184 != int32(1) {
		goto L1004
	} else {
		goto L1005
	}
L100:
	;
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v2717 != 0 {
		goto L847
	} else {
		goto L848
	}
L101:
	;
	v2530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v2530 == int32(0) {
		goto L791
	} else {
		goto L792
	}
L102:
	;
	v2279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v2279 == int32(0) {
		goto L67
	} else {
		goto L704
	}
L103:
	;
	v2171 = int32(2)
	v2172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v2172 != 0 {
		v5191 = v2171
		goto L21
	} else {
		goto L672
	}
L104:
	;
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v2133 != 0 {
		goto L655
	} else {
		goto L656
	}
L105:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1856 = F_exec_eval_expr(m, l0, v1849, v32+int32(596), v32+int32(608), v32+int32(636))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L9
	} else {
		goto L578
	}
L106:
	;
	v1706 = int32(0)
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1708+v1709<<(uint(int32(2))%32))))
	v1714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1713)+44)))
	if v1714 == v1706 {
		goto L521
	} else {
		goto L522
	}
L107:
	;
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1698 = F_exec_run_select(m, l0, v1695, v32+int32(608))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L9
	} else {
		goto L518
	}
L108:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+4))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1446+v1448<<(uint(int32(2))%32))))
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v1455 = v32 + int32(600)
	v1460 = F_exec_eval_expr(m, l0, v1453, v1455, v32+int32(608), v32+int32(636))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L9
	} else {
		goto L427
	}
L109:
	;
	goto L387
L110:
	;
	goto L373
L111:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1061 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L112:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v957 = v32 + int32(600)
	v962 = F_exec_eval_expr(m, l0, v955, v957, v32+int32(608), v32+int32(636))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L9
	} else {
		goto L290
	}
L113:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v457 == int32(1) {
		goto L190
	} else {
		goto L191
	}
L114:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+24))
	if v145 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L115:
	;
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v118 = F_exec_run_select(m, l0, v116, v115)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L118
	}
L116:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105+v106<<(uint(int32(2))%32))))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	F_exec_assign_expr(m, l0, v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L117
	}
L117:
	;
	v5191 = int32(0)
	goto L21
L118:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120+v121<<(uint(int32(2))%32))))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v129 = int32(0)
	F_assign_simple_var(m, l0, v125, base.B2i32(v126 != int64(0)), v129, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v133 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_SPI_freetuptable(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v136 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v138 == v136 {
		v5191 = v115
		goto L21
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
	F_MemoryContextReset(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	v5191 = v115
	goto L21
L126:
	;
	F_exec_prepare_plan(m, l0, v144, int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+16)))
	if v151 != int32(1) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	goto L128
L130:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[3]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+56))
	v396 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+608)) = v396
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v396
	*(*int64)(unsafe.Add(mBase, uint32(v32)+624)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v32)+608)) = v392
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+612)) = uint8(v403)
	v405 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+613)) = uint8(v405)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+628)) = v407
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v144)+24))
	v412 = F_SPI_execute_plan_extended(m, v409, v32+int32(608))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L169
	}
L131:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+20)) = v144
	v392 = v390
	goto L130
L132:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	if v359 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L133:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v144)+28))
	if v357 != 0 {
		goto L131
	} else {
		goto L160
	}
L134:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v154 != 0 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v155 = int32(_a_F_exec_stmts_1)
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v144)+24))
	v162 = F_SPI_plan_get_cached_plan(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L136
	}
L136:
	;
	if v162 == int32(0) {
		goto L88
	} else {
		goto L137
	}
L137:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v166 == int32(0) {
		goto L88
	} else {
		goto L138
	}
L138:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v169 != int32(1) {
		goto L88
	} else {
		goto L139
	}
L139:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+88))
	if v174 == int32(0) {
		goto L87
	} else {
		goto L140
	}
L140:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v177 != int32(213) {
		goto L87
	} else {
		goto L141
	}
L141:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v183 = F_SearchSysCache1(m, int32(47), v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L142
	}
L142:
	;
	if v183 == int32(0) {
		goto L86
	} else {
		goto L143
	}
L143:
	;
	v193 = F_get_func_arg_info(m, v183, v32+int32(608), v32+int32(636), v32+int32(600))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	F_ReleaseCatCache(m, v183)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+48))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v199
	v202 = F_palloc0(m, int32(40))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = int32(_a_F_exec_stmts_2)
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = int32(1)
	v212 = F_palloc(m, v193<<(uint(int32(2))%32))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+36)) = v212
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v217
	v219 = int32(0)
	if v219 < v193 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v225 = int32(0)
	v227 = v219
	goto L151
L149:
	;
	v295 = v219
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+28)) = v295
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[5]))
	F_ReleaseCachedPlan(m, v162, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L9
	} else {
		goto L159
	}
L151:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v32)+600))
	if v252 == int32(0) {
		v285 = v227
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v295 = v285
	goto L150
L153:
	;
	v289 = v225 + int32(1)
	if v289 != v193 {
		v225 = v289
		v227 = v285
		goto L151
	} else {
		goto L158
	}
L154:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+v252))))
	v258 = v256 - int32(98)
	v259 = int32(0)
	if base.B2i32(v258 == v259)|base.B2i32(v258 == int32(13)) == v259 {
		v285 = v227
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v267 = v227 << (uint(int32(2)) % 32)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267+v269)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v272 != int32(8) {
		goto L132
	} else {
		goto L156
	}
L156:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	v277 = v275 - int32(1)
	F_exec_check_assignable(m, l0, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v202)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v280+v267))) = v277
	v285 = v227 + int32(1)
	goto L153
L158:
	;
	goto L152
L159:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v202
	goto L133
L160:
	;
	v392 = int32(0)
	goto L130
L161:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L9
	} else {
		goto L165
	}
L162:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v359+v225<<(uint(int32(2))%32))))
	if v365 == int32(0) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	if v368 != 0 {
		goto L85
	} else {
		goto L164
	}
L164:
	;
	goto L161
L165:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L9
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v225 + int32(1)
	F_errmsg(m, int32(_a_F_exec_stmts_3), v32+int32(48))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2388), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L9
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	if v412 < int32(0) {
		goto L84
	} else {
		goto L170
	}
L170:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[3]))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+56))
	if v418 != v395 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
	F_plpgsql_create_econtext(m, l0)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L9
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v425 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmts[6]))
	if base.Ui64(int64(1)) < base.Ui64(v425) {
		goto L82
	} else {
		goto L175
	}
L174:
	;
	goto L173
L175:
	;
	if base.I32_wrap_i64(v425) == int32(1) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+16)))
	if v431 == int32(0) {
		goto L83
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v443 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	F_exec_move_row(m, l0, v434, v438, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	F_SPI_freetuptable(m, v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L9
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v446 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v446
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v449 != 0 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	goto L183
L185:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+20))
	F_MemoryContextReset(m, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L9
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v454 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	F_SPI_freetuptable(m, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L9
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	v5191 = v446
	goto L21
L190:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v460 == int32(0) {
		goto L81
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v463 == int32(0) {
		goto L22
	} else {
		goto L194
	}
L193:
	;
	goto L192
L194:
	;
	v466 = int32(0)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v467 <= v466 {
		goto L22
	} else {
		goto L195
	}
L195:
	;
	v472 = v466
	goto L196
L196:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v463)+12))
	v501 = int32(2)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500+v472<<(uint(v501)%32))))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v499+v505<<(uint(v501)%32))))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	switch v510 {
	case 0:
		goto L199
	case 1:
		goto L212
	case 2:
		goto L201
	case 3:
		goto L211
	case 4:
		goto L210
	case 5:
		goto L209
	case 6:
		goto L208
	case 7:
		goto L207
	case 8:
		goto L206
	case 9:
		goto L205
	case 10:
		goto L204
	case 11:
		goto L203
	case 12:
		goto L202
	default:
		goto L200
	}
L197:
	;
	goto L22
L198:
	;
	v952 = v472 + int32(1)
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v952 < v953 {
		v472 = v952
		goto L196
	} else {
		goto L289
	}
L199:
	;
	v914 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v915 = F_Int64GetDatum(m, v914)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L9
	} else {
		goto L287
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L9
	} else {
		goto L284
	}
L201:
	;
	v750 = int32(_a_F_exec_stmts_1)
	v751 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v754
	v757 = int32(_a_F_exec_stmts_6)
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[8]))
	v760 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[8])) = v759 + v760
	v763 = int32(_a_F_exec_stmts_7)
	v765 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9]))
	v767 = v765 + v760
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9])) = v767
	if v767 < int32(5) {
		goto L266
	} else {
		goto L267
	}
L202:
	;
	v731 = int32(_a_F_exec_stmts_1)
	v732 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)+60))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v736)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v737
	if v734 != 0 {
		goto L260
	} else {
		goto L261
	}
L203:
	;
	v712 = int32(_a_F_exec_stmts_1)
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+64))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v718
	if v715 != 0 {
		goto L255
	} else {
		goto L256
	}
L204:
	;
	v693 = int32(_a_F_exec_stmts_1)
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)+32))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v699
	if v696 != 0 {
		goto L250
	} else {
		goto L251
	}
L205:
	;
	v674 = int32(_a_F_exec_stmts_1)
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v676)+72))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v680
	if v677 != 0 {
		goto L245
	} else {
		goto L246
	}
L206:
	;
	v655 = int32(_a_F_exec_stmts_1)
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)+76))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v661
	if v658 != 0 {
		goto L240
	} else {
		goto L241
	}
L207:
	;
	v636 = int32(_a_F_exec_stmts_1)
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)+68))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v642
	if v639 != 0 {
		goto L235
	} else {
		goto L236
	}
L208:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v575)+28))
	v577 = int32(_a_F_exec_stmts_8)
	v578 = int32(63)
	v580 = int32(48)
	v581 = v576&v578 + v580
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[10])) = uint8(v581)
	v589 = int32(base.Ui32(v576)>>(uint(int32(24))%32))&v578 + v580
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[11])) = uint8(v589)
	v597 = int32(base.Ui32(v576)>>(uint(int32(18))%32))&v578 + v580
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[12])) = uint8(v597)
	v605 = int32(base.Ui32(v576)>>(uint(int32(12))%32))&v578 + v580
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[13])) = uint8(v605)
	v613 = int32(base.Ui32(v576)>>(uint(int32(6))%32))&v578 + v580
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[14])) = uint8(v613)
	v616 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[15])) = uint8(v616)
	goto L229
L209:
	;
	v556 = int32(_a_F_exec_stmts_1)
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)+44))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v562
	if v559 != 0 {
		goto L224
	} else {
		goto L225
	}
L210:
	;
	v537 = int32(_a_F_exec_stmts_1)
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+36))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v543
	if v540 != 0 {
		goto L219
	} else {
		goto L220
	}
L211:
	;
	v518 = int32(_a_F_exec_stmts_1)
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)+48))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v524
	if v521 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)+36))
	F_exec_assign_value(m, l0, v509, v512, int32(0), int32(26), int32(-1))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L9
	} else {
		goto L213
	}
L213:
	;
	goto L198
L214:
	;
	v527 = v521
	goto L216
L215:
	;
	v527 = int32(_a_F_exec_stmts_9)
	goto L216
L216:
	;
	v528 = F_cstring_to_text(m, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L9
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v519
	F_exec_assign_value(m, l0, v509, v528, int32(0), int32(25), int32(-1))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L9
	} else {
		goto L218
	}
L218:
	;
	goto L198
L219:
	;
	v546 = v540
	goto L221
L220:
	;
	v546 = int32(_a_F_exec_stmts_9)
	goto L221
L221:
	;
	v547 = F_cstring_to_text(m, v546)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L9
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v538
	F_exec_assign_value(m, l0, v509, v547, int32(0), int32(25), int32(-1))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L9
	} else {
		goto L223
	}
L223:
	;
	goto L198
L224:
	;
	v565 = v559
	goto L226
L225:
	;
	v565 = int32(_a_F_exec_stmts_9)
	goto L226
L226:
	;
	v566 = F_cstring_to_text(m, v565)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L9
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v557
	F_exec_assign_value(m, l0, v509, v566, int32(0), int32(25), int32(-1))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L9
	} else {
		goto L228
	}
L228:
	;
	goto L198
L229:
	;
	v619 = int32(_a_F_exec_stmts_1)
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v623
	goto L230
L230:
	;
	goto L232
L232:
	;
	v627 = F_cstring_to_text(m, v577)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L9
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v620
	F_exec_assign_value(m, l0, v509, v627, int32(0), int32(25), int32(-1))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L9
	} else {
		goto L234
	}
L234:
	;
	goto L198
L235:
	;
	v645 = v639
	goto L237
L236:
	;
	v645 = int32(_a_F_exec_stmts_9)
	goto L237
L237:
	;
	v646 = F_cstring_to_text(m, v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L9
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v637
	F_exec_assign_value(m, l0, v509, v646, int32(0), int32(25), int32(-1))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L9
	} else {
		goto L239
	}
L239:
	;
	goto L198
L240:
	;
	v664 = v658
	goto L242
L241:
	;
	v664 = int32(_a_F_exec_stmts_9)
	goto L242
L242:
	;
	v665 = F_cstring_to_text(m, v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L9
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v656
	F_exec_assign_value(m, l0, v509, v665, int32(0), int32(25), int32(-1))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L9
	} else {
		goto L244
	}
L244:
	;
	goto L198
L245:
	;
	v683 = v677
	goto L247
L246:
	;
	v683 = int32(_a_F_exec_stmts_9)
	goto L247
L247:
	;
	v684 = F_cstring_to_text(m, v683)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L9
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v675
	F_exec_assign_value(m, l0, v509, v684, int32(0), int32(25), int32(-1))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L9
	} else {
		goto L249
	}
L249:
	;
	goto L198
L250:
	;
	v702 = v696
	goto L252
L251:
	;
	v702 = int32(_a_F_exec_stmts_9)
	goto L252
L252:
	;
	v703 = F_cstring_to_text(m, v702)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v694
	F_exec_assign_value(m, l0, v509, v703, int32(0), int32(25), int32(-1))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L9
	} else {
		goto L254
	}
L254:
	;
	goto L198
L255:
	;
	v721 = v715
	goto L257
L256:
	;
	v721 = int32(_a_F_exec_stmts_9)
	goto L257
L257:
	;
	v722 = F_cstring_to_text(m, v721)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L9
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v713
	F_exec_assign_value(m, l0, v509, v722, int32(0), int32(25), int32(-1))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L9
	} else {
		goto L259
	}
L259:
	;
	goto L198
L260:
	;
	v740 = v734
	goto L262
L261:
	;
	v740 = int32(_a_F_exec_stmts_9)
	goto L262
L262:
	;
	v741 = F_cstring_to_text(m, v740)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L9
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v732
	F_exec_assign_value(m, l0, v509, v741, int32(0), int32(25), int32(-1))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L9
	} else {
		goto L264
	}
L264:
	;
	goto L198
L265:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v884)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v885
	if v839 != 0 {
		goto L279
	} else {
		goto L280
	}
L266:
	;
	v771 = int32(100)
	v772 = v767 * v771
	base.MemoryFill(m, v772+int32(_a_F_exec_stmts_10), int32(0), v771)
	v781 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v772)+uint32(_c_F_exec_stmts[17]))) = v781
	v786 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v772)+uint32(_c_F_exec_stmts[18]))) = v786
	v790 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[19]))
	if v790 != 0 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9])) = int32(-1)
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L9
	} else {
		goto L276
	}
L269:
	;
	v795 = v790
	goto L272
L270:
	;
	v839 = int32(0)
	v842 = v765
	v863 = v759
	goto L271
L271:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[8])) = v863
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9])) = v842
	goto L265
L272:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v795)+8))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	m.T0[v821].(func(*base.Module, int32))(m, v820)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L9
	} else {
		goto L274
	}
L273:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9]))
	v827 = int32(1)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v772)+uint32(_c_F_exec_stmts[20])))
	v831 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[8]))
	v839 = v829
	v842 = v826 - v827
	v863 = v831 - v827
	goto L271
L274:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v795)))
	if v824 != 0 {
		v795 = v824
		goto L272
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_11), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L9
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_12), int32(762), int32(_a_F_exec_stmts_13))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L9
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
	v888 = v839
	goto L281
L280:
	;
	v888 = int32(_a_F_exec_stmts_9)
	goto L281
L281:
	;
	v889 = F_cstring_to_text(m, v888)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L9
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v751
	F_exec_assign_value(m, l0, v509, v889, int32(0), int32(25), int32(-1))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L9
	} else {
		goto L283
	}
L283:
	;
	goto L198
L284:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = v902
	F_errmsg_internal(m, int32(_a_F_exec_stmts_14), v32+int32(80))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L9
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2510), int32(_a_F_exec_stmts_15))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L9
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
	F_exec_assign_value(m, l0, v509, v915, int32(0), int32(20), int32(-1))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L9
	} else {
		goto L288
	}
L288:
	;
	goto L198
L289:
	;
	goto L197
L290:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v968 = F_exec_cast_value(m, l0, v962, v957, v964, v965, int32(16), int32(-1))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L9
	} else {
		goto L291
	}
L291:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v970 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	F_SPI_freetuptable(m, v970)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L9
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v975 != 0 {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	goto L294
L296:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v975)+20))
	F_MemoryContextReset(m, v976)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L9
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	v982 = int32(0)
	if v981|base.B2i32(v968 == v982) == v982 {
		v5144 = v89 + int32(16)
		goto L23
	} else {
		goto L300
	}
L299:
	;
	goto L298
L300:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v987 == int32(0) {
		goto L24
	} else {
		goto L301
	}
L301:
	;
	v990 = int32(0)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v987)+4))
	if v991 <= v990 {
		goto L24
	} else {
		goto L302
	}
L302:
	;
	v996 = v990
	goto L303
L303:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v987)+12))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1023+v996<<(uint(int32(2))%32))))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+4))
	v1030 = v32 + int32(600)
	v1035 = F_exec_eval_expr(m, l0, v1028, v1030, v32+int32(608), v32+int32(636))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L9
	} else {
		goto L305
	}
L304:
	;
	v5144 = v1027 + int32(8)
	goto L23
L305:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v1041 = F_exec_cast_value(m, l0, v1035, v1030, v1037, v1038, int32(16), int32(-1))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L9
	} else {
		goto L306
	}
L306:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1043 != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	F_SPI_freetuptable(m, v1043)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L9
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1048 != 0 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	goto L309
L311:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+20))
	F_MemoryContextReset(m, v1049)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L9
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v1041 != 0 {
		goto L315
	} else {
		goto L316
	}
L314:
	;
	goto L313
L315:
	;
	v1054 = v1052
	goto L317
L316:
	;
	v1054 = int32(1)
	goto L317
L317:
	;
	if v1054 != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1056 = v996 + int32(1)
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v987)+4))
	if v1057 <= v1056 {
		goto L24
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	goto L304
L321:
	;
	v996 = v1056
	goto L303
L322:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v1111 == int32(0) {
		goto L27
	} else {
		goto L340
	}
L323:
	;
	v1108 = int32(0)
	goto L322
L324:
	;
	goto L325
L325:
	;
	v1071 = F_exec_eval_expr(m, l0, v1061, v32+int32(600), v32+int32(608), v32+int32(636))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L9
	} else {
		goto L326
	}
L326:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1075+v1076<<(uint(int32(2))%32))))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+24))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	if v1074 == v1082 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	F_exec_assign_value(m, l0, v1080, v1071, v1092, v1074, v1073)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L9
	} else {
		goto L333
	}
L328:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+24))
	if v1084 == v1073 {
		goto L327
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+44))
	v1089 = F_plpgsql_build_datatype(m, v1074, v1073, v1087, int32(0))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L9
	} else {
		goto L332
	}
L331:
	;
	goto L330
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080)+24)) = v1089
	goto L327
L333:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1095 != 0 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	F_SPI_freetuptable(m, v1095)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L9
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	v1098 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1098
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1100 == v1098 {
		v1108 = v1080
		goto L322
	} else {
		goto L338
	}
L337:
	;
	goto L336
L338:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+20))
	F_MemoryContextReset(m, v1103)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L9
	} else {
		goto L339
	}
L339:
	;
	v1108 = v1080
	goto L322
L340:
	;
	v1114 = int32(0)
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+4))
	if v1115 <= v1114 {
		goto L27
	} else {
		goto L341
	}
L341:
	;
	v1120 = v1114
	goto L342
L342:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+12))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1147+v1120<<(uint(int32(2))%32))))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+4))
	v1154 = v32 + int32(600)
	v1159 = F_exec_eval_expr(m, l0, v1152, v1154, v32+int32(608), v32+int32(636))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L9
	} else {
		goto L344
	}
L343:
	;
	if v1108 != 0 {
		goto L361
	} else {
		goto L362
	}
L344:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v1165 = F_exec_cast_value(m, l0, v1159, v1154, v1161, v1162, int32(16), int32(-1))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L9
	} else {
		goto L345
	}
L345:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1167 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	F_SPI_freetuptable(m, v1167)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L9
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1172 != 0 {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	goto L348
L350:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+20))
	F_MemoryContextReset(m, v1173)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L9
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v1165 != 0 {
		goto L354
	} else {
		goto L355
	}
L353:
	;
	goto L352
L354:
	;
	v1178 = v1176
	goto L356
L355:
	;
	v1178 = int32(1)
	goto L356
L356:
	;
	if v1178 != 0 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1180 = v1120 + int32(1)
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+4))
	if v1181 <= v1180 {
		goto L27
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	goto L343
L360:
	;
	v1120 = v1180
	goto L342
L361:
	;
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+45)))
	if v1183 != int32(1) {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	goto L363
L363:
	;
	v5060 = v1151 + int32(8)
	goto L26
L364:
	;
	v1205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+48)) = v1205
	v1207 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1108)+44)) = uint16(v1207)
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+40)) = v1205
	goto L363
L365:
	;
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+44)))
	if v1186 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+40))
	F_pfree(m, v1201)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L9
	} else {
		goto L372
	}
L367:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+24))
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1187)+12)))
	if v1188 != int32(_a_F_exec_stmts_16) {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+40))
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191))))
	if v1192 != int32(1) {
		goto L366
	} else {
		goto L369
	}
L369:
	;
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191)+1)))
	if v1195 != int32(3) {
		goto L366
	} else {
		goto L370
	}
L370:
	;
	F_DeleteExpandedObject(m, v1191)
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L9
	} else {
		goto L371
	}
L371:
	;
	goto L364
L372:
	;
	goto L364
L373:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1244 = F_exec_stmts(m, l0, v1243)
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L9
	} else {
		goto L376
	}
L375:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1248 == int32(0) {
		goto L373
	} else {
		goto L377
	}
L376:
	;
	switch v1244 - int32(1) {
	case 0:
		goto L29
	case 1:
		v5191 = v1244
		goto L21
	case 2:
		goto L375
	default:
		goto L373
	}
L377:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1251 == int32(0) {
		v5191 = v1244
		goto L21
	} else {
		goto L378
	}
L378:
	;
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1251))))
	v1259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248))))
	if base.B2i32(v1256 == int32(0))|base.B2i32(v1256 != v1259) != 0 {
		v1277 = v1256
		v1278 = v1259
		goto L380
	} else {
		goto L381
	}
L379:
	;
	if v1277-v1278 != 0 {
		v5191 = v1244
		goto L21
	} else {
		goto L386
	}
L380:
	;
	goto L379
L381:
	;
	v1262 = v1251
	v1263 = v1248
	goto L382
L382:
	;
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1263)+1)))
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1262)+1)))
	if v1267 == int32(0) {
		v1277 = v1267
		v1278 = v1266
		goto L380
	} else {
		goto L384
	}
L383:
	;
	v1277 = v1267
	v1278 = v1266
	goto L380
L384:
	;
	v1270 = int32(1)
	if v1267 == v1266 {
		v1262 = v1262 + v1270
		v1263 = v1263 + v1270
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
	goto L373
L387:
	;
	v1314 = int32(0)
	goto L389
L389:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1343 = v32 + int32(600)
	v1348 = F_exec_eval_expr(m, l0, v1341, v1343, v32+int32(608), v32+int32(636))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L9
	} else {
		goto L391
	}
L391:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v1354 = F_exec_cast_value(m, l0, v1348, v1343, v1350, v1351, int32(16), int32(-1))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L9
	} else {
		goto L392
	}
L392:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1356 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	F_SPI_freetuptable(m, v1356)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L9
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1361 != 0 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	goto L395
L397:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+20))
	F_MemoryContextReset(m, v1362)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L9
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v1365|base.B2i32(v1354 == int32(0)) != 0 {
		v5191 = v1314
		goto L21
	} else {
		goto L401
	}
L400:
	;
	goto L399
L401:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v1370 = F_exec_stmts(m, l0, v1369)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L9
	} else {
		goto L404
	}
L402:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1410 == int32(0) {
		goto L387
	} else {
		goto L417
	}
L403:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1374 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L404:
	;
	switch v1370 - int32(1) {
	case 0:
		goto L403
	case 1:
		v5191 = v1370
		goto L21
	case 2:
		goto L402
	default:
		v1314 = v1370
		goto L389
	}
L405:
	;
	v5191 = int32(0)
	goto L21
L406:
	;
	goto L407
L407:
	;
	v1378 = int32(1)
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1379 == int32(0) {
		v5191 = v1378
		goto L21
	} else {
		goto L408
	}
L408:
	;
	v1384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1379))))
	v1387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1374))))
	if base.B2i32(v1384 == int32(0))|base.B2i32(v1384 != v1387) != 0 {
		v1405 = v1384
		v1406 = v1387
		goto L410
	} else {
		goto L411
	}
L409:
	;
	if v1405-v1406 == int32(0) {
		goto L28
	} else {
		goto L416
	}
L410:
	;
	goto L409
L411:
	;
	v1390 = v1379
	v1391 = v1374
	goto L412
L412:
	;
	v1394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391)+1)))
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1390)+1)))
	if v1395 == int32(0) {
		v1405 = v1395
		v1406 = v1394
		goto L410
	} else {
		goto L414
	}
L413:
	;
	v1405 = v1395
	v1406 = v1394
	goto L410
L414:
	;
	v1398 = int32(1)
	if v1395 == v1394 {
		v1390 = v1390 + v1398
		v1391 = v1391 + v1398
		goto L412
	} else {
		goto L415
	}
L415:
	;
	goto L413
L416:
	;
	v5191 = v1378
	goto L21
L417:
	;
	v1413 = int32(3)
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1414 == int32(0) {
		v5191 = v1413
		goto L21
	} else {
		goto L418
	}
L418:
	;
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1414))))
	v1422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1410))))
	if base.B2i32(v1419 == int32(0))|base.B2i32(v1419 != v1422) != 0 {
		v1440 = v1419
		v1441 = v1422
		goto L420
	} else {
		goto L421
	}
L419:
	;
	if v1440-v1441 != 0 {
		v5191 = v1413
		goto L21
	} else {
		goto L426
	}
L420:
	;
	goto L419
L421:
	;
	v1425 = v1414
	v1426 = v1410
	goto L422
L422:
	;
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426)+1)))
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1425)+1)))
	if v1430 == int32(0) {
		v1440 = v1430
		v1441 = v1429
		goto L420
	} else {
		goto L424
	}
L423:
	;
	v1440 = v1430
	v1441 = v1429
	goto L420
L424:
	;
	v1433 = int32(1)
	if v1430 == v1429 {
		v1425 = v1425 + v1433
		v1426 = v1426 + v1433
		goto L422
	} else {
		goto L425
	}
L425:
	;
	goto L423
L426:
	;
	v1443 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1443
	v1314 = v1443
	goto L389
L427:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1452)+24))
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+4))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+24))
	v1467 = F_exec_cast_value(m, l0, v1460, v1455, v1462, v1463, v1465, v1466)
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L9
	} else {
		goto L428
	}
L428:
	;
	v1469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v1469 == int32(1) {
		goto L80
	} else {
		goto L429
	}
L429:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1472 != 0 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	F_SPI_freetuptable(m, v1472)
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L9
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1477 != 0 {
		goto L434
	} else {
		goto L435
	}
L433:
	;
	goto L432
L434:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+20))
	F_MemoryContextReset(m, v1478)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L9
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1483 = v32 + int32(600)
	v1488 = F_exec_eval_expr(m, l0, v1481, v1483, v32+int32(608), v32+int32(636))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L9
	} else {
		goto L438
	}
L437:
	;
	goto L436
L438:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1452)+24))
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1492)+4))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1492)+24))
	v1495 = F_exec_cast_value(m, l0, v1488, v1483, v1490, v1491, v1493, v1494)
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L9
	} else {
		goto L439
	}
L439:
	;
	v1497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v1497 == int32(1) {
		goto L79
	} else {
		goto L440
	}
L440:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1500 != 0 {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	F_SPI_freetuptable(m, v1500)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L9
	} else {
		goto L444
	}
L442:
	;
	goto L443
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1505 != 0 {
		goto L445
	} else {
		goto L446
	}
L444:
	;
	goto L443
L445:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+20))
	F_MemoryContextReset(m, v1506)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L9
	} else {
		goto L448
	}
L446:
	;
	goto L447
L447:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v1509 == int32(0) {
		goto L450
	} else {
		goto L451
	}
L448:
	;
	goto L447
L449:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	if v1544 != 0 {
		goto L466
	} else {
		goto L467
	}
L450:
	;
	v1543 = int32(1)
	goto L449
L451:
	;
	goto L452
L452:
	;
	v1514 = v32 + int32(600)
	v1519 = F_exec_eval_expr(m, l0, v1509, v1514, v32+int32(608), v32+int32(636))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L9
	} else {
		goto L453
	}
L453:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1452)+24))
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1523)+4))
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1523)+24))
	v1526 = F_exec_cast_value(m, l0, v1519, v1514, v1521, v1522, v1524, v1525)
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L9
	} else {
		goto L454
	}
L454:
	;
	v1528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v1528 == int32(1) {
		goto L78
	} else {
		goto L455
	}
L455:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1531 != 0 {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	F_SPI_freetuptable(m, v1531)
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L9
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1536 != 0 {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	goto L458
L460:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+20))
	F_MemoryContextReset(m, v1537)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L9
	} else {
		goto L463
	}
L461:
	;
	goto L462
L462:
	;
	if v1526 <= int32(0) {
		goto L77
	} else {
		goto L464
	}
L463:
	;
	goto L462
L464:
	;
	v1543 = v1526
	goto L449
L465:
	;
	v1547 = int32(0)
	F_assign_simple_var(m, l0, v1452, v1467, v1547, v1547)
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L9
	} else {
		goto L471
	}
L466:
	;
	if v1495 <= v1467 {
		goto L465
	} else {
		goto L469
	}
L467:
	;
	goto L468
L468:
	;
	if v1495 < v1467 {
		goto L31
	} else {
		goto L470
	}
L469:
	;
	goto L31
L470:
	;
	goto L465
L471:
	;
	v1551 = int32(1)
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	v1553 = F_exec_stmts(m, l0, v1552)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L9
	} else {
		goto L474
	}
L472:
	;
	v1598 = v1543 ^ int32(2147483647)
	v1600 = v1543 | int32(-2147483648)
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	if v1601 != 0 {
		goto L486
	} else {
		goto L487
	}
L473:
	;
	v1557 = int32(0)
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1558 == v1557 {
		v1594 = v1557
		goto L472
	} else {
		goto L475
	}
L474:
	;
	switch v1553 - int32(1) {
	case 0:
		goto L32
	case 1:
		v4865 = v1553
		v4868 = v1551
		goto L30
	case 2:
		goto L473
	default:
		v1594 = v1553
		goto L472
	}
L475:
	;
	v1561 = int32(3)
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1562 == int32(0) {
		v4865 = v1561
		v4868 = v1551
		goto L30
	} else {
		goto L476
	}
L476:
	;
	v1567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1562))))
	v1570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1558))))
	if base.B2i32(v1567 == int32(0))|base.B2i32(v1567 != v1570) != 0 {
		v1588 = v1567
		v1589 = v1570
		goto L478
	} else {
		goto L479
	}
L477:
	;
	if v1588-v1589 != 0 {
		v4865 = v1561
		v4868 = v1551
		goto L30
	} else {
		goto L484
	}
L478:
	;
	goto L477
L479:
	;
	v1573 = v1562
	v1574 = v1558
	goto L480
L480:
	;
	v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1574)+1)))
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1573)+1)))
	if v1578 == int32(0) {
		v1588 = v1578
		v1589 = v1577
		goto L478
	} else {
		goto L482
	}
L481:
	;
	v1588 = v1578
	v1589 = v1577
	goto L478
L482:
	;
	v1581 = int32(1)
	if v1578 == v1577 {
		v1573 = v1573 + v1581
		v1574 = v1574 + v1581
		goto L480
	} else {
		goto L483
	}
L483:
	;
	goto L481
L484:
	;
	v1591 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1591
	v1594 = v1591
	goto L472
L485:
	;
	v1611 = v1594
	v1613 = v1608
	v1619 = v1607
	goto L491
L486:
	;
	if v1467 < v1600 {
		v4865 = v1594
		v4868 = v1551
		goto L30
	} else {
		goto L489
	}
L487:
	;
	goto L488
L488:
	;
	if v1598 < v1467 {
		v4865 = v1594
		v4868 = v1551
		goto L30
	} else {
		goto L490
	}
L489:
	;
	v1607 = int32(1)
	v1608 = v1467 - v1543
	goto L485
L490:
	;
	v1607 = v1601
	v1608 = v1467 + v1543
	goto L485
L491:
	;
	if v1619 != 0 {
		goto L494
	} else {
		goto L495
	}
L493:
	;
	v1640 = int32(0)
	F_assign_simple_var(m, l0, v1452, v1613, v1640, v1640)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L9
	} else {
		goto L499
	}
L494:
	;
	if v1495 <= v1613 {
		goto L493
	} else {
		goto L497
	}
L495:
	;
	goto L496
L496:
	;
	if v1495 < v1613 {
		v4865 = v1611
		v4868 = v1551
		goto L30
	} else {
		goto L498
	}
L497:
	;
	v4865 = v1611
	v4868 = v1551
	goto L30
L498:
	;
	goto L493
L499:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	v1645 = F_exec_stmts(m, l0, v1644)
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L9
	} else {
		goto L502
	}
L500:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	if v1689 != 0 {
		goto L513
	} else {
		goto L514
	}
L501:
	;
	v1649 = int32(0)
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1650 == v1649 {
		v1686 = v1649
		goto L500
	} else {
		goto L503
	}
L502:
	;
	switch v1645 - int32(1) {
	case 0:
		goto L32
	case 1:
		v4865 = v1645
		v4868 = v1551
		goto L30
	case 2:
		goto L501
	default:
		v1686 = v1645
		goto L500
	}
L503:
	;
	v1653 = int32(3)
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1654 == int32(0) {
		v4865 = v1653
		v4868 = v1551
		goto L30
	} else {
		goto L504
	}
L504:
	;
	v1659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654))))
	v1662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1650))))
	if base.B2i32(v1659 == int32(0))|base.B2i32(v1659 != v1662) != 0 {
		v1680 = v1659
		v1681 = v1662
		goto L506
	} else {
		goto L507
	}
L505:
	;
	if v1680-v1681 != 0 {
		v4865 = v1653
		v4868 = v1551
		goto L30
	} else {
		goto L512
	}
L506:
	;
	goto L505
L507:
	;
	v1665 = v1654
	v1666 = v1650
	goto L508
L508:
	;
	v1669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1666)+1)))
	v1670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1665)+1)))
	if v1670 == int32(0) {
		v1680 = v1670
		v1681 = v1669
		goto L506
	} else {
		goto L510
	}
L509:
	;
	v1680 = v1670
	v1681 = v1669
	goto L506
L510:
	;
	v1673 = int32(1)
	if v1670 == v1669 {
		v1665 = v1665 + v1673
		v1666 = v1666 + v1673
		goto L508
	} else {
		goto L511
	}
L511:
	;
	goto L509
L512:
	;
	v1683 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1683
	v1686 = v1683
	goto L500
L513:
	;
	if v1613 < v1600 {
		v4865 = v1686
		v4868 = v1551
		goto L30
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	if v1598 < v1613 {
		v4865 = v1686
		v4868 = v1551
		goto L30
	} else {
		goto L517
	}
L516:
	;
	v1611 = v1686
	v1613 = v1613 - v1543
	v1619 = int32(1)
	goto L491
L517:
	;
	v1611 = v1686
	v1613 = v1613 + v1543
	v1619 = v1689
	goto L491
L518:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1702 = F_exec_for_query(m, l0, v89, v1700, int32(1))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L9
	} else {
		goto L519
	}
L519:
	;
	F_SPI_cursor_close(m, v1700)
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L9
	} else {
		goto L520
	}
L520:
	;
	v5191 = v1702
	goto L21
L521:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v1717 == int32(0) {
		goto L524
	} else {
		goto L525
	}
L522:
	;
	v1740 = v1706
	v1741 = v1706
	goto L523
L523:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v1743 != 0 {
		goto L532
	} else {
		goto L533
	}
L524:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v1725 = F_AllocSetContextCreateInternal(m, v1720, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L9
	} else {
		goto L527
	}
L525:
	;
	v1728 = v1717
	goto L526
L526:
	;
	v1729 = int32(_a_F_exec_stmts_1)
	v1730 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1728
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+40))
	v1734 = F_text_to_cstring(m, v1733)
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L9
	} else {
		goto L528
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v1725
	v1728 = v1725
	goto L526
L528:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1730
	v1738 = F_GetPortalByName(m, v1734)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L9
	} else {
		goto L529
	}
L529:
	;
	if v1738 != 0 {
		goto L76
	} else {
		goto L530
	}
L530:
	;
	v1740 = v1728
	v1741 = v1734
	goto L523
L531:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+28))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1773)+24))
	if v1774 == int32(0) {
		goto L538
	} else {
		goto L539
	}
L532:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+32))
	if v1744 < int32(0) {
		goto L75
	} else {
		goto L535
	}
L533:
	;
	goto L534
L534:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+32))
	if int32(0) <= v1768 {
		goto L74
	} else {
		goto L537
	}
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+608)) = int32(16)
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v1754 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+626)) = uint8(v1754)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+620)) = v1743
	*(*int32)(unsafe.Add(mBase, uint32(v32)+612)) = v1753
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1758+v1744<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+628)) = v1762
	F_exec_stmt_execsql(m, l0, v32+int32(608))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L9
	} else {
		goto L536
	}
L536:
	;
	goto L531
L537:
	;
	goto L531
L538:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+36))
	F_exec_prepare_plan(m, l0, v1773, v1777)
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L9
	} else {
		goto L541
	}
L539:
	;
	goto L540
L540:
	;
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v1773)+28))
	if v1780 == int32(0) {
		goto L543
	} else {
		goto L544
	}
L541:
	;
	goto L540
L542:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1773)+24))
	v1788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	v1789 = F_SPI_cursor_open_internal(m, v1741, v1787, v1786, v1788)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L9
	} else {
		goto L546
	}
L543:
	;
	v1786 = int32(0)
	goto L542
L544:
	;
	goto L545
L545:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1784)+20)) = v1773
	v1786 = v1784
	goto L542
L546:
	;
	if v1789 == int32(0) {
		goto L73
	} else {
		goto L547
	}
L547:
	;
	if v1741 == int32(0) {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	F_exec_check_assignable(m, l0, v1795)
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L9
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1805 != 0 {
		goto L554
	} else {
		goto L555
	}
L551:
	;
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1789)))
	v1799 = F_cstring_to_text(m, v1798)
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L9
	} else {
		goto L552
	}
L552:
	;
	F_assign_simple_var(m, l0, v1713, v1799, int32(0), int32(1))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L9
	} else {
		goto L553
	}
L553:
	;
	goto L550
L554:
	;
	F_SPI_freetuptable(m, v1805)
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L9
	} else {
		goto L557
	}
L555:
	;
	goto L556
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1810 != 0 {
		goto L558
	} else {
		goto L559
	}
L557:
	;
	goto L556
L558:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1810)+20))
	F_MemoryContextReset(m, v1811)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L9
	} else {
		goto L561
	}
L559:
	;
	goto L560
L560:
	;
	if v1740 != 0 {
		goto L562
	} else {
		goto L563
	}
L561:
	;
	goto L560
L562:
	;
	F_MemoryContextReset(m, v1740)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L9
	} else {
		goto L565
	}
L563:
	;
	goto L564
L564:
	;
	v1817 = F_exec_for_query(m, l0, v89, v1789, int32(0))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L9
	} else {
		goto L566
	}
L565:
	;
	goto L564
L566:
	;
	F_SPI_cursor_close(m, v1789)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L9
	} else {
		goto L567
	}
L567:
	;
	if v1741 != 0 {
		v5191 = v1817
		goto L21
	} else {
		goto L568
	}
L568:
	;
	v1821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1713)+45)))
	if v1821 != int32(1) {
		goto L569
	} else {
		goto L570
	}
L569:
	;
	v1843 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+48)) = v1843
	v1845 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1713)+44)) = uint16(v1845)
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+40)) = v1843
	v5191 = v1817
	goto L21
L570:
	;
	v1824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1713)+44)))
	if v1824 != 0 {
		goto L571
	} else {
		goto L572
	}
L571:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+40))
	F_pfree(m, v1839)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L9
	} else {
		goto L577
	}
L572:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+24))
	v1826 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1825)+12)))
	if v1826 != int32(_a_F_exec_stmts_16) {
		goto L571
	} else {
		goto L573
	}
L573:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+40))
	v1830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1829))))
	if v1830 != int32(1) {
		goto L571
	} else {
		goto L574
	}
L574:
	;
	v1833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1829)+1)))
	if v1833 != int32(3) {
		goto L571
	} else {
		goto L575
	}
L575:
	;
	F_DeleteExpandedObject(m, v1829)
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L9
	} else {
		goto L576
	}
L576:
	;
	goto L569
L577:
	;
	goto L569
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+600)) = v1856
	v1859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+596)))
	if v1859 == int32(1) {
		goto L72
	} else {
		goto L579
	}
L579:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v1862 == int32(0) {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v1870 = F_AllocSetContextCreateInternal(m, v1865, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L9
	} else {
		goto L583
	}
L581:
	;
	v1872 = v1862
	goto L582
L582:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1872
	v1876 = int32(_a_F_exec_stmts_1)
	v1877 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1872
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1881 = F_get_element_type(m, v1880)
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L9
	} else {
		goto L584
	}
L583:
	;
	v1872 = v1870
	goto L582
L584:
	;
	if v1881 == int32(0) {
		goto L71
	} else {
		goto L585
	}
L585:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v32)+600))
	v1886 = F_pg_detoast_datum_copy(m, v1885)
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L9
	} else {
		goto L586
	}
L586:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1888 != 0 {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	F_SPI_freetuptable(m, v1888)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L9
	} else {
		goto L590
	}
L588:
	;
	goto L589
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1893 != 0 {
		goto L591
	} else {
		goto L592
	}
L590:
	;
	goto L589
L591:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+20))
	F_MemoryContextReset(m, v1894)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L9
	} else {
		goto L594
	}
L592:
	;
	goto L593
L593:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v1897 < int32(0) {
		goto L70
	} else {
		goto L595
	}
L594:
	;
	goto L593
L595:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v1886)+4))
	if v1900 < v1897 {
		goto L70
	} else {
		goto L596
	}
L596:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1905 = int32(2)
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1903+v1904<<(uint(v1905)%32))))
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v1908)))
	if base.Ui32(v1905) <= base.Ui32(v1909-int32(1)) {
		goto L597
	} else {
		goto L598
	}
L597:
	;
	v1914 = F_plpgsql_exec_get_datum_type(m, l0, v1908)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L9
	} else {
		goto L600
	}
L598:
	;
	v1919 = v1897
	v1920 = int32(0)
	goto L599
L599:
	;
	v1921 = int32(0)
	if base.B2i32(v1920 == v1921)&base.B2i32(v1921 < v1919) != 0 {
		goto L69
	} else {
		goto L602
	}
L600:
	;
	v1916 = F_get_element_type(m, v1914)
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L9
	} else {
		goto L601
	}
L601:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v1919 = v1918
	v1920 = v1916
	goto L599
L602:
	;
	if v1920 != 0 {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v1927 = v1919
	goto L605
L604:
	;
	v1927 = int32(1)
	goto L605
L605:
	;
	if v1927 == int32(0) {
		goto L68
	} else {
		goto L606
	}
L606:
	;
	v1931 = F_array_create_iterator(m, v1886, v1919, int32(0))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L9
	} else {
		goto L607
	}
L607:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if int32(0) < v1937 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	v1940 = v32 + int32(608)
	goto L610
L609:
	;
	v1940 = v1886 + int32(12)
	goto L610
L610:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1940)))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v1943 = int32(0)
	v1949 = F_array_iterate(m, v1931, v32+int32(600), v32+int32(596))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L9
	} else {
		goto L612
	}
L611:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1877
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v2117
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2117)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v2119
	F_MemoryContextReset(m, v1872)
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L9
	} else {
		goto L653
	}
L612:
	;
	if v1949 == int32(0) {
		v2088 = v1943
		v2090 = v1943
		goto L611
	} else {
		goto L613
	}
L613:
	;
	goto L614
L614:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1877
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v32)+600))
	v1985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+596)))
	F_exec_assign_value(m, l0, v1908, v1984, v1985, v1941, v1942)
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L9
	} else {
		goto L616
	}
L615:
	;
	v2088 = v2075
	v2090 = v1988
	goto L611
L616:
	;
	v1988 = int32(1)
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if int32(0) < v1989 {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v32)+600))
	F_pfree(m, v1992)
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L9
	} else {
		goto L620
	}
L618:
	;
	goto L619
L619:
	;
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v1996 = F_exec_stmts(m, l0, v1995)
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L9
	} else {
		goto L624
	}
L620:
	;
	goto L619
L621:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1872
	v2084 = F_array_iterate(m, v1931, v32+int32(600), v32+int32(596))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L9
	} else {
		goto L651
	}
L622:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2038 == int32(0) {
		goto L639
	} else {
		goto L640
	}
L623:
	;
	v2000 = int32(0)
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2001 == v2000 {
		v2088 = v2000
		v2090 = v1988
		goto L611
	} else {
		goto L625
	}
L624:
	;
	switch v1996 - int32(1) {
	case 0:
		goto L623
	case 1:
		v2088 = v1996
		v2090 = v1988
		goto L611
	case 2:
		goto L622
	default:
		v2075 = v1996
		goto L621
	}
L625:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2004 == int32(0) {
		goto L626
	} else {
		goto L627
	}
L626:
	;
	v2088 = int32(1)
	v2090 = v1988
	goto L611
L627:
	;
	goto L628
L628:
	;
	v2010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2004))))
	v2013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2001))))
	if base.B2i32(v2010 == int32(0))|base.B2i32(v2010 != v2013) != 0 {
		v2031 = v2010
		v2032 = v2013
		goto L630
	} else {
		goto L631
	}
L629:
	;
	if v2031-v2032 != 0 {
		goto L636
	} else {
		goto L637
	}
L630:
	;
	goto L629
L631:
	;
	v2016 = v2004
	v2017 = v2001
	goto L632
L632:
	;
	v2020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2017)+1)))
	v2021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2016)+1)))
	if v2021 == int32(0) {
		v2031 = v2021
		v2032 = v2020
		goto L630
	} else {
		goto L634
	}
L633:
	;
	v2031 = v2021
	v2032 = v2020
	goto L630
L634:
	;
	v2024 = int32(1)
	if v2021 == v2020 {
		v2016 = v2016 + v2024
		v2017 = v2017 + v2024
		goto L632
	} else {
		goto L635
	}
L635:
	;
	goto L633
L636:
	;
	v2088 = int32(1)
	v2090 = v1988
	goto L611
L637:
	;
	goto L638
L638:
	;
	v2035 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2035
	v2088 = v2035
	v2090 = v1988
	goto L611
L639:
	;
	v2075 = int32(0)
	goto L621
L640:
	;
	goto L641
L641:
	;
	v2042 = int32(3)
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2043 == int32(0) {
		v2088 = v2042
		v2090 = v1988
		goto L611
	} else {
		goto L642
	}
L642:
	;
	v2048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2043))))
	v2051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2038))))
	if base.B2i32(v2048 == int32(0))|base.B2i32(v2048 != v2051) != 0 {
		v2069 = v2048
		v2070 = v2051
		goto L644
	} else {
		goto L645
	}
L643:
	;
	if v2069-v2070 != 0 {
		v2088 = v2042
		v2090 = v1988
		goto L611
	} else {
		goto L650
	}
L644:
	;
	goto L643
L645:
	;
	v2054 = v2043
	v2055 = v2038
	goto L646
L646:
	;
	v2058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2055)+1)))
	v2059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2054)+1)))
	if v2059 == int32(0) {
		v2069 = v2059
		v2070 = v2058
		goto L644
	} else {
		goto L648
	}
L647:
	;
	v2069 = v2059
	v2070 = v2058
	goto L644
L648:
	;
	v2062 = int32(1)
	if v2059 == v2058 {
		v2054 = v2054 + v2062
		v2055 = v2055 + v2062
		goto L646
	} else {
		goto L649
	}
L649:
	;
	goto L647
L650:
	;
	v2072 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2072
	v2075 = v2072
	goto L621
L651:
	;
	if v2084 != 0 {
		goto L614
	} else {
		goto L652
	}
L652:
	;
	goto L615
L653:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v2123+v2124<<(uint(int32(2))%32))))
	v2129 = int32(0)
	F_assign_simple_var(m, l0, v2128, v2090, v2129, v2129)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L9
	} else {
		goto L654
	}
L654:
	;
	v5191 = v2088
	goto L21
L655:
	;
	v2135 = v32 + int32(600)
	v2140 = F_exec_eval_expr(m, l0, v2133, v2135, v32+int32(608), v32+int32(636))
	mBase = m.M
	v2141 = m.ExcPending
	if v2141 != 0 {
		goto L9
	} else {
		goto L658
	}
L656:
	;
	goto L657
L657:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2165
	v2169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v2169 != 0 {
		goto L669
	} else {
		goto L670
	}
L658:
	;
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v2146 = F_exec_cast_value(m, l0, v2140, v2135, v2142, v2143, int32(16), int32(-1))
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L9
	} else {
		goto L659
	}
L659:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2148 != 0 {
		goto L660
	} else {
		goto L661
	}
L660:
	;
	F_SPI_freetuptable(m, v2148)
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L9
	} else {
		goto L663
	}
L661:
	;
	goto L662
L662:
	;
	v2151 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2151
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2154 != 0 {
		goto L664
	} else {
		goto L665
	}
L663:
	;
	goto L662
L664:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v2154)+20))
	F_MemoryContextReset(m, v2155)
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L9
	} else {
		goto L667
	}
L665:
	;
	goto L666
L666:
	;
	v2158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v2158|base.B2i32(v2146 == int32(0)) != 0 {
		v5191 = v2151
		goto L21
	} else {
		goto L668
	}
L667:
	;
	goto L666
L668:
	;
	goto L657
L669:
	;
	v2170 = int32(1)
	goto L671
L670:
	;
	v2170 = int32(3)
	goto L671
L671:
	;
	v5191 = v2170
	goto L21
L672:
	;
	v2173 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v2173
	v2175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v2175)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v2173
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v2173 <= v2179 {
		goto L673
	} else {
		goto L674
	}
L673:
	;
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v2182+v2179<<(uint(int32(2))%32))))
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v2186)))
	switch v2187 {
	case 0:
		goto L678
	case 1, 2:
		goto L677
	default:
		goto L676
	case 4:
		goto L679
	}
L674:
	;
	goto L675
L675:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2239 != 0 {
		goto L690
	} else {
		goto L691
	}
L676:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L9
	} else {
		goto L687
	}
L677:
	;
	F_exec_eval_datum(m, l0, v2186, v39, v32+int32(608), v43, v41)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L9
	} else {
		goto L686
	}
L678:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v2186)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2190
	v2192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2186)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v2192)
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v2186)+24))
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2194)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2195
	v2197 = int32(1)
	v2199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2192&v2197|base.B2i32(v2199 != v2197) != 0 {
		v5191 = v2171
		goto L21
	} else {
		goto L681
	}
L679:
	;
	F_plpgsql_fulfill_promise(m, l0, v2186)
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L9
	} else {
		goto L680
	}
L680:
	;
	goto L678
L681:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L9
	} else {
		goto L682
	}
L682:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L9
	} else {
		goto L683
	}
L683:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_20), int32(0))
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		goto L9
	} else {
		goto L684
	}
L684:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3255), int32(_a_F_exec_stmts_21))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L9
	} else {
		goto L685
	}
L685:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L686:
	;
	v5191 = v2171
	goto L21
L687:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2186)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+160)) = v2227
	F_errmsg_internal(m, int32(_a_F_exec_stmts_22), v32+int32(160))
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L9
	} else {
		goto L688
	}
L688:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3275), int32(_a_F_exec_stmts_21))
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L9
	} else {
		goto L689
	}
L689:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L690:
	;
	v2242 = F_exec_eval_expr(m, l0, v2239, v41, v39, v32+int32(608))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L9
	} else {
		goto L693
	}
L691:
	;
	goto L692
L692:
	;
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2268 != int32(2278) {
		v5191 = v2171
		goto L21
	} else {
		goto L702
	}
L693:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2242
	v2245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2245 != int32(1) {
		v5191 = v2171
		goto L21
	} else {
		goto L694
	}
L694:
	;
	v2248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v2248 != 0 {
		v5191 = v2171
		goto L21
	} else {
		goto L695
	}
L695:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v2250 = F_type_is_rowtype(m, v2249)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L9
	} else {
		goto L696
	}
L696:
	;
	if v2250 != 0 {
		v5191 = v2171
		goto L21
	} else {
		goto L697
	}
L697:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L9
	} else {
		goto L698
	}
L698:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L9
	} else {
		goto L699
	}
L699:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_20), int32(0))
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		goto L9
	} else {
		goto L700
	}
L700:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3298), int32(_a_F_exec_stmts_21))
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
		goto L9
	} else {
		goto L701
	}
L701:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L702:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2271)+65)))
	if v2272 == int32(112) {
		v5191 = v2171
		goto L21
	} else {
		goto L703
	}
L703:
	;
	v2275 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v2275)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(2278)
	v5191 = v2171
	goto L21
L704:
	;
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2282 == int32(0) {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	F_exec_init_tuple_store(m, l0)
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L9
	} else {
		goto L708
	}
L706:
	;
	goto L707
L707:
	;
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v2287)))
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if int32(0) <= v2289 {
		goto L710
	} else {
		goto L711
	}
L708:
	;
	goto L707
L709:
	;
	v2518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2518 != 0 {
		goto L785
	} else {
		goto L786
	}
L710:
	;
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2292+v2289<<(uint(int32(2))%32))))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2296)))
	switch v2297 {
	case 0:
		v2301 = v2288
		goto L716
	case 1:
		goto L714
	case 2:
		goto L715
	default:
		goto L713
	case 4:
		goto L717
	}
L711:
	;
	goto L712
L712:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2413 == int32(0) {
		goto L62
	} else {
		goto L754
	}
L713:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		goto L9
	} else {
		goto L751
	}
L714:
	;
	v2382 = int32(_a_F_exec_stmts_1)
	v2383 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2386
	v2388 = F_make_tuple_from_row(m, l0, v2296, v2287)
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L9
	} else {
		goto L748
	}
L715:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+36))
	if v2342 == int32(0) {
		goto L729
	} else {
		goto L730
	}
L716:
	;
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+40))
	v2303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2296)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+636)) = uint8(v2303)
	if v2288 != int32(1) {
		goto L66
	} else {
		goto L719
	}
L717:
	;
	F_plpgsql_fulfill_promise(m, l0, v2296)
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L9
	} else {
		goto L718
	}
L718:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2287)))
	v2301 = v2300
	goto L716
L719:
	;
	v2309 = v2287 + v2301<<(uint(int32(4))%32)
	if v2303&int32(1) != 0 {
		v2326 = v2302
		goto L720
	} else {
		goto L721
	}
L720:
	;
	v2328 = v32 + int32(636)
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+24))
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v2329)+4))
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(v2329)+24))
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v2309)+88))
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v2309)+96))
	v2334 = F_exec_cast_value(m, l0, v2326, v2328, v2330, v2331, v2332, v2333)
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L9
	} else {
		goto L727
	}
L721:
	;
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+24))
	v2313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2312)+12)))
	if v2313 != int32(_a_F_exec_stmts_16) {
		v2326 = v2302
		goto L720
	} else {
		goto L722
	}
L722:
	;
	v2316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2302))))
	if v2316 != int32(1) {
		v2325 = v2302
		goto L724
	} else {
		goto L725
	}
L723:
	;
	v2326 = v2325
	goto L720
L724:
	;
	goto L723
L725:
	;
	v2319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2302)+1)))
	if v2319 != int32(3) {
		v2325 = v2302
		goto L724
	} else {
		goto L726
	}
L726:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v2302)+2))
	v2325 = v2322 + int32(18)
	goto L724
L727:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+608)) = v2334
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_putvalues(m, v2337, v2287, v32+int32(608), v2328)
	mBase = m.M
	v2341 = m.ExcPending
	if v2341 != 0 {
		goto L9
	} else {
		goto L728
	}
L728:
	;
	goto L709
L729:
	;
	F_instantiate_empty_record_variable(m, l0, v2296)
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L9
	} else {
		goto L732
	}
L730:
	;
	v2348 = v2342
	goto L731
L731:
	;
	v2349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2348)+28)))
	if v2349&int32(5) == int32(0) {
		goto L733
	} else {
		goto L734
	}
L732:
	;
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+36))
	v2348 = v2347
	goto L731
L733:
	;
	F_deconstruct_expanded_record(m, v2348)
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L9
	} else {
		goto L736
	}
L734:
	;
	v2357 = v2348
	goto L735
L735:
	;
	v2358 = int32(_a_F_exec_stmts_1)
	v2359 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v2361)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2362
	v2364 = *(*int32)(unsafe.Add(mBase, uint32(v2357)+44))
	if v2364 != 0 {
		goto L737
	} else {
		goto L738
	}
L736:
	;
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+36))
	v2357 = v2356
	goto L735
L737:
	;
	v2367 = v2364
	goto L739
L738:
	;
	v2365 = F_expanded_record_fetch_tupdesc(m, v2357)
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L9
	} else {
		goto L740
	}
L739:
	;
	v2369 = F_convert_tuples_by_position(m, v2367, v2287, int32(_a_F_exec_stmts_23))
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L9
	} else {
		goto L741
	}
L740:
	;
	v2367 = v2365
	goto L739
L741:
	;
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+36))
	v2372 = F_expanded_record_get_tuple(m, v2371)
	mBase = m.M
	v2373 = m.ExcPending
	if v2373 != 0 {
		goto L9
	} else {
		goto L742
	}
L742:
	;
	if v2369 != 0 {
		goto L743
	} else {
		goto L744
	}
L743:
	;
	v2374 = F_execute_attr_map_tuple(m, v2372, v2369)
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L9
	} else {
		goto L746
	}
L744:
	;
	v2376 = v2372
	goto L745
L745:
	;
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_puttuple(m, v2377, v2376)
	mBase = m.M
	v2379 = m.ExcPending
	if v2379 != 0 {
		goto L9
	} else {
		goto L747
	}
L746:
	;
	v2376 = v2374
	goto L745
L747:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2359
	goto L709
L748:
	;
	if v2388 == int32(0) {
		goto L65
	} else {
		goto L749
	}
L749:
	;
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_puttuple(m, v2392, v2388)
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L9
	} else {
		goto L750
	}
L750:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2383
	goto L709
L751:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2296)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+176)) = v2401
	F_errmsg_internal(m, int32(_a_F_exec_stmts_22), v32+int32(176))
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L9
	} else {
		goto L752
	}
L752:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3444), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L9
	} else {
		goto L753
	}
L753:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L754:
	;
	v2422 = F_exec_eval_expr(m, l0, v2413, v32+int32(607), v32+int32(600), v32+int32(596))
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L9
	} else {
		goto L755
	}
L755:
	;
	v2424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2424 == int32(1) {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v2427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+607)))
	if v2427 == int32(0) {
		goto L759
	} else {
		goto L760
	}
L757:
	;
	goto L758
L758:
	;
	if v2288 != int32(1) {
		goto L63
	} else {
		goto L782
	}
L759:
	;
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(v32)+600))
	v2431 = F_type_is_rowtype(m, v2430)
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L9
	} else {
		goto L762
	}
L760:
	;
	goto L761
L761:
	;
	v2478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+20))
	v2482 = F_MemoryContextAllocZero(m, v2479, v2288<<(uint(int32(2))%32))
	mBase = m.M
	v2483 = m.ExcPending
	if v2483 != 0 {
		goto L9
	} else {
		goto L776
	}
L762:
	;
	if v2431 == int32(0) {
		goto L64
	} else {
		goto L763
	}
L763:
	;
	v2435 = int32(_a_F_exec_stmts_1)
	v2436 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v2438)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2439
	v2441 = F_pg_detoast_datum(m, v2422)
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L9
	} else {
		goto L764
	}
L764:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v2441)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+624)) = v2441
	v2445 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+620)) = v2445
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+616)) = uint16(v2445)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+612)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+608)) = int32(base.Ui32(v2443) >> (uint(int32(2)) % 32))
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+8))
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+4))
	v2456 = F_lookup_rowtype_tupdesc(m, v2454, v2455)
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L9
	} else {
		goto L765
	}
L765:
	;
	v2459 = F_convert_tuples_by_position(m, v2456, v2287, int32(_a_F_exec_stmts_25))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L9
	} else {
		goto L766
	}
L766:
	;
	if v2459 != 0 {
		goto L767
	} else {
		goto L768
	}
L767:
	;
	v2463 = F_execute_attr_map_tuple(m, v32+int32(608), v2459)
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L9
	} else {
		goto L770
	}
L768:
	;
	v2467 = v32 + int32(608)
	goto L769
L769:
	;
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_puttuple(m, v2468, v2467)
	mBase = m.M
	v2470 = m.ExcPending
	if v2470 != 0 {
		goto L9
	} else {
		goto L771
	}
L770:
	;
	v2467 = v2463
	goto L769
L771:
	;
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v2456)+12))
	if int32(0) <= v2471 {
		goto L772
	} else {
		goto L773
	}
L772:
	;
	F_DecrTupleDescRefCount(m, v2456)
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L9
	} else {
		goto L775
	}
L773:
	;
	goto L774
L774:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2436
	goto L709
L775:
	;
	goto L774
L776:
	;
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v2484)+20))
	v2486 = F_MemoryContextAlloc(m, v2485, v2288)
	mBase = m.M
	v2487 = m.ExcPending
	if v2487 != 0 {
		goto L9
	} else {
		goto L777
	}
L777:
	;
	if v2288 != 0 {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	base.MemoryFill(m, v2486, int32(1), v2288)
	goto L780
L779:
	;
	goto L780
L780:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_putvalues(m, v2490, v2287, v2482, v2486)
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L9
	} else {
		goto L781
	}
L781:
	;
	goto L709
L782:
	;
	v2496 = v32 + int32(607)
	v2497 = *(*int32)(unsafe.Add(mBase, uint32(v32)+600))
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(v32)+596))
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(v2287)))
	v2502 = v2287 + v2499<<(uint(int32(4))%32)
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(v2502)+88))
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(v2502)+96))
	v2505 = F_exec_cast_value(m, l0, v2422, v2496, v2497, v2498, v2503, v2504)
	mBase = m.M
	v2506 = m.ExcPending
	if v2506 != 0 {
		goto L9
	} else {
		goto L783
	}
L783:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+636)) = v2505
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_putvalues(m, v2508, v2287, v32+int32(636), v2496)
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L9
	} else {
		goto L784
	}
L784:
	;
	goto L709
L785:
	;
	F_SPI_freetuptable(m, v2518)
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L9
	} else {
		goto L788
	}
L786:
	;
	goto L787
L787:
	;
	v2521 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2521
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2524 == v2521 {
		v5191 = v2521
		goto L21
	} else {
		goto L789
	}
L788:
	;
	goto L787
L789:
	;
	v2527 = *(*int32)(unsafe.Add(mBase, uint32(v2524)+20))
	F_MemoryContextReset(m, v2527)
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L9
	} else {
		goto L790
	}
L790:
	;
	v5191 = v2521
	goto L21
L791:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v2538 = F_AllocSetContextCreateInternal(m, v2533, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L9
	} else {
		goto L794
	}
L792:
	;
	v2541 = v2530
	goto L793
L793:
	;
	v2542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v2542 == int32(0) {
		goto L61
	} else {
		goto L795
	}
L794:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v2538
	v2541 = v2538
	goto L793
L795:
	;
	v2545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2545 != 0 {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	v2549 = v2545
	goto L798
L797:
	;
	F_exec_init_tuple_store(m, l0)
	mBase = m.M
	v2547 = m.ExcPending
	if v2547 != 0 {
		goto L9
	} else {
		goto L799
	}
L798:
	;
	v2550 = *(*int64)(unsafe.Add(mBase, uint32(v2549)+40))
	v2551 = int32(_a_F_exec_stmts_1)
	v2552 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2541
	v2556 = F_CreateDestReceiver(m, int32(6))
	mBase = m.M
	v2557 = m.ExcPending
	if v2557 != 0 {
		goto L9
	} else {
		goto L800
	}
L799:
	;
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v2549 = v2548
	goto L798
L800:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2560 = int32(0)
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+36)) = int32(_a_F_exec_stmts_26)
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+32)) = v2561
	*(*uint8)(unsafe.Add(mBase, uint32(v2556)+28)) = uint8(v2560)
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+24)) = v2559
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+20)) = v2558
	goto L801
L801:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2552
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2570 != 0 {
		goto L803
	} else {
		goto L804
	}
L802:
	;
	v2687 = *(*int32)(unsafe.Add(mBase, uint32(v2556)+12))
	m.T0[v2687].(func(*base.Module, int32))(m, v2556)
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L9
	} else {
		goto L836
	}
L803:
	;
	v2571 = *(*int32)(unsafe.Add(mBase, uint32(v2570)+24))
	if v2571 == int32(0) {
		goto L806
	} else {
		goto L807
	}
L804:
	;
	goto L805
L805:
	;
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v2629 = F_exec_eval_expr(m, l0, v2622, v32+int32(607), v32+int32(636), v32+int32(600))
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L9
	} else {
		goto L820
	}
L806:
	;
	F_exec_prepare_plan(m, l0, v2570, int32(2048))
	mBase = m.M
	v2576 = m.ExcPending
	if v2576 != 0 {
		goto L9
	} else {
		goto L809
	}
L807:
	;
	goto L808
L808:
	;
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(v2570)+28))
	if v2577 == int32(0) {
		goto L811
	} else {
		goto L812
	}
L809:
	;
	goto L808
L810:
	;
	v2584 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+608)) = v2584
	*(*int64)(unsafe.Add(mBase, uint32(v32)+624)) = v2584
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v2584
	*(*int32)(unsafe.Add(mBase, uint32(v32)+608)) = v2583
	v2591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+624)) = v2556
	v2593 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+614)) = uint8(v2593)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+612)) = uint8(v2591)
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(v2570)+24))
	v2599 = F_SPI_execute_plan_extended(m, v2596, v32+int32(608))
	mBase = m.M
	v2600 = m.ExcPending
	if v2600 != 0 {
		goto L9
	} else {
		goto L814
	}
L811:
	;
	v2583 = int32(0)
	goto L810
L812:
	;
	goto L813
L813:
	;
	v2581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v2581)+20)) = v2570
	v2583 = v2581
	goto L810
L814:
	;
	if int32(0) <= v2599 {
		goto L802
	} else {
		goto L815
	}
L815:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2606 = m.ExcPending
	if v2606 != 0 {
		goto L9
	} else {
		goto L816
	}
L816:
	;
	v2607 = *(*int32)(unsafe.Add(mBase, uint32(v2570)))
	v2608 = F_SPI_result_code_string(m, v2599)
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L9
	} else {
		goto L817
	}
L817:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+212)) = v2608
	*(*int32)(unsafe.Add(mBase, uint32(v32)+208)) = v2607
	F_errmsg_internal(m, int32(_a_F_exec_stmts_27), v32+int32(208))
	mBase = m.M
	v2616 = m.ExcPending
	if v2616 != 0 {
		goto L9
	} else {
		goto L818
	}
L818:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3608), int32(_a_F_exec_stmts_28))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L9
	} else {
		goto L819
	}
L819:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L820:
	;
	v2631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+607)))
	if v2631 == int32(1) {
		goto L60
	} else {
		goto L821
	}
L821:
	;
	v2634 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v2635 = int32(_a_F_exec_stmts_1)
	v2636 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2639 = *(*int32)(unsafe.Add(mBase, uint32(v2638)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2639
	F_getTypeOutputInfo(m, v2634, v32+int32(608), v32+int32(596))
	mBase = m.M
	v2646 = m.ExcPending
	if v2646 != 0 {
		goto L9
	} else {
		goto L822
	}
L822:
	;
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v2648 = F_OidOutputFunctionCall(m, v2647, v2629)
	mBase = m.M
	v2649 = m.ExcPending
	if v2649 != 0 {
		goto L9
	} else {
		goto L823
	}
L823:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2636
	v2652 = F_MemoryContextStrdup(m, v2541, v2648)
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L9
	} else {
		goto L824
	}
L824:
	;
	v2654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2654 != 0 {
		goto L825
	} else {
		goto L826
	}
L825:
	;
	F_SPI_freetuptable(m, v2654)
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L9
	} else {
		goto L828
	}
L826:
	;
	goto L827
L827:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2659 != 0 {
		goto L829
	} else {
		goto L830
	}
L828:
	;
	goto L827
L829:
	;
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v2659)+20))
	F_MemoryContextReset(m, v2660)
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		goto L9
	} else {
		goto L832
	}
L830:
	;
	goto L831
L831:
	;
	v2663 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+624)) = v2663
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v2663
	*(*int64)(unsafe.Add(mBase, uint32(v32)+608)) = v2663
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v2670 = F_exec_eval_using_params(m, l0, v2669)
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L9
	} else {
		goto L833
	}
L832:
	;
	goto L831
L833:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+608)) = v2670
	v2673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+624)) = v2556
	v2675 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+614)) = uint8(v2675)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+612)) = uint8(v2673)
	v2680 = F_SPI_execute_extended(m, v2652, v32+int32(608))
	mBase = m.M
	v2681 = m.ExcPending
	if v2681 != 0 {
		goto L9
	} else {
		goto L834
	}
L834:
	;
	if v2680 < int32(0) {
		goto L59
	} else {
		goto L835
	}
L835:
	;
	goto L802
L836:
	;
	v2690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2690 != 0 {
		goto L837
	} else {
		goto L838
	}
L837:
	;
	F_SPI_freetuptable(m, v2690)
	mBase = m.M
	v2692 = m.ExcPending
	if v2692 != 0 {
		goto L9
	} else {
		goto L840
	}
L838:
	;
	goto L839
L839:
	;
	v2693 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2693
	v2696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2696 != 0 {
		goto L841
	} else {
		goto L842
	}
L840:
	;
	goto L839
L841:
	;
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+20))
	F_MemoryContextReset(m, v2697)
	mBase = m.M
	v2699 = m.ExcPending
	if v2699 != 0 {
		goto L9
	} else {
		goto L844
	}
L842:
	;
	goto L843
L843:
	;
	F_MemoryContextReset(m, v2541)
	mBase = m.M
	v2701 = m.ExcPending
	if v2701 != 0 {
		goto L9
	} else {
		goto L845
	}
L844:
	;
	goto L843
L845:
	;
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v2703 = *(*int64)(unsafe.Add(mBase, uint32(v2702)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v2703 - v2550
	v2706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2707 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v2706+v2707<<(uint(int32(2))%32))))
	v2713 = int32(0)
	F_assign_simple_var(m, l0, v2711, base.B2i32(v2550 != v2703), v2713, v2713)
	mBase = m.M
	v2716 = m.ExcPending
	if v2716 != 0 {
		goto L9
	} else {
		goto L846
	}
L846:
	;
	v5191 = v2693
	goto L21
L847:
	;
	v2737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v2737 == int32(0) {
		goto L856
	} else {
		goto L857
	}
L848:
	;
	v2718 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v2718 != 0 {
		goto L847
	} else {
		goto L849
	}
L849:
	;
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v2719 != 0 {
		goto L847
	} else {
		goto L850
	}
L850:
	;
	v2720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2720 != 0 {
		goto L58
	} else {
		goto L851
	}
L851:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L9
	} else {
		goto L852
	}
L852:
	;
	F_errcode(m, int32(33557120))
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L9
	} else {
		goto L853
	}
L853:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_29), int32(0))
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L9
	} else {
		goto L854
	}
L854:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3749), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L9
	} else {
		goto L855
	}
L855:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L856:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v2745 = F_AllocSetContextCreateInternal(m, v2740, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v2746 = m.ExcPending
	if v2746 != 0 {
		goto L9
	} else {
		goto L859
	}
L857:
	;
	v2749 = v2717
	v2750 = v2737
	goto L858
L858:
	;
	v2751 = int32(0)
	if v2749 == v2751 {
		goto L861
	} else {
		goto L862
	}
L859:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v2745
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v2749 = v2748
	v2750 = v2745
	goto L858
L860:
	;
	v2764 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v2764 != 0 {
		goto L866
	} else {
		goto L867
	}
L861:
	;
	v2754 = int32(0)
	v2762 = v2754
	v2763 = v2754
	goto L860
L862:
	;
	goto L863
L863:
	;
	v2757 = F_plpgsql_recognize_err_condition(m, v2749, int32(1))
	mBase = m.M
	v2758 = m.ExcPending
	if v2758 != 0 {
		goto L9
	} else {
		goto L864
	}
L864:
	;
	v2759 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v2760 = F_MemoryContextStrdup(m, v2750, v2759)
	mBase = m.M
	v2761 = m.ExcPending
	if v2761 != 0 {
		goto L9
	} else {
		goto L865
	}
L865:
	;
	v2762 = v2757
	v2763 = v2760
	goto L860
L866:
	;
	v2765 = int32(_a_F_exec_stmts_1)
	v2766 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2750
	F_initStringInfo(m, v32+int32(608))
	mBase = m.M
	v2772 = m.ExcPending
	if v2772 != 0 {
		goto L9
	} else {
		goto L869
	}
L867:
	;
	v2900 = v2751
	goto L868
L868:
	;
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v2925 == int32(0) {
		goto L906
	} else {
		goto L907
	}
L869:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2766
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	if v2775 != 0 {
		goto L870
	} else {
		goto L871
	}
L870:
	;
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v2775)+12))
	v2778 = v2776
	goto L872
L871:
	;
	v2778 = int32(0)
	goto L872
L872:
	;
	v2779 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v2782 = v2779
	v2784 = v2778
	goto L873
L873:
	;
	v2809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2782))))
	if v2809 != int32(37) {
		goto L877
	} else {
		goto L878
	}
L874:
	;
	if v2784 != 0 {
		goto L56
	} else {
		goto L904
	}
L875:
	;
	goto L874
L876:
	;
	v2782 = v2888 + int32(1)
	v2784 = v2889
	goto L873
L877:
	;
	if v2809 == int32(0) {
		goto L875
	} else {
		goto L880
	}
L878:
	;
	goto L879
L879:
	;
	v2819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2782)+1)))
	if v2819 == int32(37) {
		goto L882
	} else {
		goto L883
	}
L880:
	;
	F_appendStringInfoChar(m, v32+int32(608), base.I32_extend8_s(v2809))
	mBase = m.M
	v2818 = m.ExcPending
	if v2818 != 0 {
		goto L9
	} else {
		goto L881
	}
L881:
	;
	v2888 = v2782
	v2889 = v2784
	goto L876
L882:
	;
	F_appendStringInfoChar(m, v32+int32(608), int32(37))
	mBase = m.M
	v2828 = m.ExcPending
	if v2828 != 0 {
		goto L9
	} else {
		goto L885
	}
L883:
	;
	goto L884
L884:
	;
	if v2784 == int32(0) {
		goto L57
	} else {
		goto L886
	}
L885:
	;
	v2888 = v2782 + int32(1)
	v2889 = v2784
	goto L876
L886:
	;
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(v2784)))
	v2838 = F_exec_eval_expr(m, l0, v2831, v32+int32(595), v32+int32(600), v32+int32(596))
	mBase = m.M
	v2839 = m.ExcPending
	if v2839 != 0 {
		goto L9
	} else {
		goto L887
	}
L887:
	;
	v2840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+595)))
	if v2840 != 0 {
		goto L889
	} else {
		goto L890
	}
L888:
	;
	F_appendStringInfoString(m, v32+int32(608), v2860)
	mBase = m.M
	v2865 = m.ExcPending
	if v2865 != 0 {
		goto L9
	} else {
		goto L894
	}
L889:
	;
	v2860 = int32(_a_F_exec_stmts_31)
	goto L888
L890:
	;
	goto L891
L891:
	;
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(v32)+600))
	v2843 = int32(_a_F_exec_stmts_1)
	v2844 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2846 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2847 = *(*int32)(unsafe.Add(mBase, uint32(v2846)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2847
	F_getTypeOutputInfo(m, v2842, v32+int32(636), v32+int32(607))
	mBase = m.M
	v2854 = m.ExcPending
	if v2854 != 0 {
		goto L9
	} else {
		goto L892
	}
L892:
	;
	v2855 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v2856 = F_OidOutputFunctionCall(m, v2855, v2838)
	mBase = m.M
	v2857 = m.ExcPending
	if v2857 != 0 {
		goto L9
	} else {
		goto L893
	}
L893:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2844
	v2860 = v2856
	goto L888
L894:
	;
	v2867 = v2784 + int32(4)
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2868)+12))
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(v2868)+4))
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2875 != 0 {
		goto L895
	} else {
		goto L896
	}
L895:
	;
	F_SPI_freetuptable(m, v2875)
	mBase = m.M
	v2877 = m.ExcPending
	if v2877 != 0 {
		goto L9
	} else {
		goto L898
	}
L896:
	;
	goto L897
L897:
	;
	if base.Ui32(v2867) < base.Ui32(v2869+v2870<<(uint(int32(2))%32)) {
		goto L899
	} else {
		goto L900
	}
L898:
	;
	goto L897
L899:
	;
	v2879 = v2867
	goto L901
L900:
	;
	v2879 = int32(0)
	goto L901
L901:
	;
	v2880 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2880
	v2882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2882 == v2880 {
		v2888 = v2782
		v2889 = v2879
		goto L876
	} else {
		goto L902
	}
L902:
	;
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(v2882)+20))
	F_MemoryContextReset(m, v2885)
	mBase = m.M
	v2887 = m.ExcPending
	if v2887 != 0 {
		goto L9
	} else {
		goto L903
	}
L903:
	;
	v2888 = v2782
	v2889 = v2879
	goto L876
L904:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v2900 = v2895
	goto L868
L905:
	;
	if v3062 == int32(0) {
		goto L954
	} else {
		goto L955
	}
L906:
	;
	v2928 = int32(0)
	v3057 = v2900
	v3058 = v2763
	v3062 = v2762
	v3063 = v2928
	v3064 = v2928
	v3065 = v2928
	v3066 = v2928
	v3069 = v2928
	v3070 = v2928
	v3071 = v2928
	goto L905
L907:
	;
	goto L908
L908:
	;
	v2935 = int32(0)
	v2943 = *(*int32)(unsafe.Add(mBase, uint32(v2925)+4))
	if v2943 <= v2935 {
		v3057 = v2900
		v3058 = v2763
		v3062 = v2762
		v3063 = v2935
		v3064 = v2935
		v3065 = v2935
		v3066 = v2935
		v3069 = v2935
		v3070 = v2935
		v3071 = v2935
		goto L905
	} else {
		goto L909
	}
L909:
	;
	v2948 = v2935
	v2950 = v2900
	v2951 = v2763
	v2955 = v2762
	v2956 = v2935
	v2957 = v2935
	v2958 = v2935
	v2959 = v2935
	v2962 = v2935
	v2963 = v2935
	v2964 = v2935
	goto L910
L910:
	;
	v2975 = *(*int32)(unsafe.Add(mBase, uint32(v2925)+12))
	v2979 = *(*int32)(unsafe.Add(mBase, uint32(v2975+v2948<<(uint(int32(2))%32))))
	v2980 = *(*int32)(unsafe.Add(mBase, uint32(v2979)+4))
	v2987 = F_exec_eval_expr(m, l0, v2980, v32+int32(607), v32+int32(636), v32+int32(600))
	mBase = m.M
	v2988 = m.ExcPending
	if v2988 != 0 {
		goto L9
	} else {
		goto L912
	}
L911:
	;
	v3057 = v3030
	v3058 = v3031
	v3062 = v3032
	v3063 = v3033
	v3064 = v3034
	v3065 = v3035
	v3066 = v3036
	v3069 = v3037
	v3070 = v3038
	v3071 = v3039
	goto L905
L912:
	;
	v2989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+607)))
	if v2989 != 0 {
		goto L55
	} else {
		goto L913
	}
L913:
	;
	v2990 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v2991 = int32(_a_F_exec_stmts_1)
	v2992 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2995 = *(*int32)(unsafe.Add(mBase, uint32(v2994)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2995
	F_getTypeOutputInfo(m, v2990, v32+int32(608), v32+int32(596))
	mBase = m.M
	v3002 = m.ExcPending
	if v3002 != 0 {
		goto L9
	} else {
		goto L914
	}
L914:
	;
	v3003 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v3004 = F_OidOutputFunctionCall(m, v3003, v2987)
	mBase = m.M
	v3005 = m.ExcPending
	if v3005 != 0 {
		goto L9
	} else {
		goto L915
	}
L915:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2992
	v3008 = *(*int32)(unsafe.Add(mBase, uint32(v2979)))
	switch v3008 {
	case 0:
		goto L925
	case 1:
		goto L924
	case 2:
		goto L923
	case 3:
		goto L922
	case 4:
		goto L921
	case 5:
		goto L920
	case 6:
		goto L919
	case 7:
		goto L918
	case 8:
		goto L917
	default:
		goto L45
	}
L916:
	;
	v3040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3040 != 0 {
		goto L945
	} else {
		goto L946
	}
L917:
	;
	if v2958 != 0 {
		goto L46
	} else {
		goto L943
	}
L918:
	;
	if v2957 != 0 {
		goto L47
	} else {
		goto L941
	}
L919:
	;
	if v2956 != 0 {
		goto L48
	} else {
		goto L939
	}
L920:
	;
	if v2959 != 0 {
		goto L49
	} else {
		goto L937
	}
L921:
	;
	if v2962 != 0 {
		goto L50
	} else {
		goto L935
	}
L922:
	;
	if v2963 != 0 {
		goto L51
	} else {
		goto L933
	}
L923:
	;
	if v2964 != 0 {
		goto L52
	} else {
		goto L931
	}
L924:
	;
	if v2950 != 0 {
		goto L53
	} else {
		goto L929
	}
L925:
	;
	if v2955 != 0 {
		goto L54
	} else {
		goto L926
	}
L926:
	;
	v3010 = F_plpgsql_recognize_err_condition(m, v3004, int32(1))
	mBase = m.M
	v3011 = m.ExcPending
	if v3011 != 0 {
		goto L9
	} else {
		goto L927
	}
L927:
	;
	v3012 = F_MemoryContextStrdup(m, v2750, v3004)
	mBase = m.M
	v3013 = m.ExcPending
	if v3013 != 0 {
		goto L9
	} else {
		goto L928
	}
L928:
	;
	v3030 = v2950
	v3031 = v3012
	v3032 = v3010
	v3033 = v2956
	v3034 = v2957
	v3035 = v2958
	v3036 = v2959
	v3037 = v2962
	v3038 = v2963
	v3039 = v2964
	goto L916
L929:
	;
	v3014 = F_MemoryContextStrdup(m, v2750, v3004)
	mBase = m.M
	v3015 = m.ExcPending
	if v3015 != 0 {
		goto L9
	} else {
		goto L930
	}
L930:
	;
	v3030 = v3014
	v3031 = v2951
	v3032 = v2955
	v3033 = v2956
	v3034 = v2957
	v3035 = v2958
	v3036 = v2959
	v3037 = v2962
	v3038 = v2963
	v3039 = v2964
	goto L916
L931:
	;
	v3016 = F_MemoryContextStrdup(m, v2750, v3004)
	mBase = m.M
	v3017 = m.ExcPending
	if v3017 != 0 {
		goto L9
	} else {
		goto L932
	}
L932:
	;
	v3030 = v2950
	v3031 = v2951
	v3032 = v2955
	v3033 = v2956
	v3034 = v2957
	v3035 = v2958
	v3036 = v2959
	v3037 = v2962
	v3038 = v2963
	v3039 = v3016
	goto L916
L933:
	;
	v3018 = F_MemoryContextStrdup(m, v2750, v3004)
	mBase = m.M
	v3019 = m.ExcPending
	if v3019 != 0 {
		goto L9
	} else {
		goto L934
	}
L934:
	;
	v3030 = v2950
	v3031 = v2951
	v3032 = v2955
	v3033 = v2956
	v3034 = v2957
	v3035 = v2958
	v3036 = v2959
	v3037 = v2962
	v3038 = v3018
	v3039 = v2964
	goto L916
L935:
	;
	v3020 = F_MemoryContextStrdup(m, v2750, v3004)
	mBase = m.M
	v3021 = m.ExcPending
	if v3021 != 0 {
		goto L9
	} else {
		goto L936
	}
L936:
	;
	v3030 = v2950
	v3031 = v2951
	v3032 = v2955
	v3033 = v2956
	v3034 = v2957
	v3035 = v2958
	v3036 = v2959
	v3037 = v3020
	v3038 = v2963
	v3039 = v2964
	goto L916
L937:
	;
	v3022 = F_MemoryContextStrdup(m, v2750, v3004)
	mBase = m.M
	v3023 = m.ExcPending
	if v3023 != 0 {
		goto L9
	} else {
		goto L938
	}
L938:
	;
	v3030 = v2950
	v3031 = v2951
	v3032 = v2955
	v3033 = v2956
	v3034 = v2957
	v3035 = v2958
	v3036 = v3022
	v3037 = v2962
	v3038 = v2963
	v3039 = v2964
	goto L916
L939:
	;
	v3024 = F_MemoryContextStrdup(m, v2750, v3004)
	mBase = m.M
	v3025 = m.ExcPending
	if v3025 != 0 {
		goto L9
	} else {
		goto L940
	}
L940:
	;
	v3030 = v2950
	v3031 = v2951
	v3032 = v2955
	v3033 = v3024
	v3034 = v2957
	v3035 = v2958
	v3036 = v2959
	v3037 = v2962
	v3038 = v2963
	v3039 = v2964
	goto L916
L941:
	;
	v3026 = F_MemoryContextStrdup(m, v2750, v3004)
	mBase = m.M
	v3027 = m.ExcPending
	if v3027 != 0 {
		goto L9
	} else {
		goto L942
	}
L942:
	;
	v3030 = v2950
	v3031 = v2951
	v3032 = v2955
	v3033 = v2956
	v3034 = v3026
	v3035 = v2958
	v3036 = v2959
	v3037 = v2962
	v3038 = v2963
	v3039 = v2964
	goto L916
L943:
	;
	v3028 = F_MemoryContextStrdup(m, v2750, v3004)
	mBase = m.M
	v3029 = m.ExcPending
	if v3029 != 0 {
		goto L9
	} else {
		goto L944
	}
L944:
	;
	v3030 = v2950
	v3031 = v2951
	v3032 = v2955
	v3033 = v2956
	v3034 = v2957
	v3035 = v3028
	v3036 = v2959
	v3037 = v2962
	v3038 = v2963
	v3039 = v2964
	goto L916
L945:
	;
	F_SPI_freetuptable(m, v3040)
	mBase = m.M
	v3042 = m.ExcPending
	if v3042 != 0 {
		goto L9
	} else {
		goto L948
	}
L946:
	;
	goto L947
L947:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v3045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3045 != 0 {
		goto L949
	} else {
		goto L950
	}
L948:
	;
	goto L947
L949:
	;
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v3045)+20))
	F_MemoryContextReset(m, v3046)
	mBase = m.M
	v3048 = m.ExcPending
	if v3048 != 0 {
		goto L9
	} else {
		goto L952
	}
L950:
	;
	goto L951
L951:
	;
	v3050 = v2948 + int32(1)
	v3051 = *(*int32)(unsafe.Add(mBase, uint32(v2925)+4))
	if v3050 < v3051 {
		v2948 = v3050
		v2950 = v3030
		v2951 = v3031
		v2955 = v3032
		v2956 = v3033
		v2957 = v3034
		v2958 = v3035
		v2959 = v3036
		v2962 = v3037
		v2963 = v3038
		v2964 = v3039
		goto L910
	} else {
		goto L953
	}
L952:
	;
	goto L951
L953:
	;
	goto L911
L954:
	;
	v3086 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if int32(20) < v3086 {
		goto L957
	} else {
		goto L958
	}
L955:
	;
	v3090 = v3062
	goto L956
L956:
	;
	if v3057 != 0 {
		v3135 = v3057
		goto L960
	} else {
		goto L961
	}
L957:
	;
	v3089 = int32(16777248)
	goto L959
L958:
	;
	v3089 = int32(0)
	goto L959
L959:
	;
	v3090 = v3089
	goto L956
L960:
	;
	v3136 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3138 = F_errstart(m, v3136, int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3139 = m.ExcPending
	if v3139 != 0 {
		goto L9
	} else {
		goto L965
	}
L961:
	;
	if v3058 != 0 {
		v3135 = v3058
		goto L960
	} else {
		goto L962
	}
L962:
	;
	v3091 = int32(_a_F_exec_stmts_8)
	v3092 = int32(63)
	v3094 = int32(48)
	v3095 = v3090&v3092 + v3094
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[10])) = uint8(v3095)
	v3103 = int32(base.Ui32(v3090)>>(uint(int32(24))%32))&v3092 + v3094
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[11])) = uint8(v3103)
	v3111 = int32(base.Ui32(v3090)>>(uint(int32(18))%32))&v3092 + v3094
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[12])) = uint8(v3111)
	v3119 = int32(base.Ui32(v3090)>>(uint(int32(12))%32))&v3092 + v3094
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[13])) = uint8(v3119)
	v3127 = int32(base.Ui32(v3090)>>(uint(int32(6))%32))&v3092 + v3094
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[14])) = uint8(v3127)
	v3130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[15])) = uint8(v3130)
	goto L963
L963:
	;
	v3133 = F_MemoryContextStrdup(m, v2750, v3091)
	mBase = m.M
	v3134 = m.ExcPending
	if v3134 != 0 {
		goto L9
	} else {
		goto L964
	}
L964:
	;
	v3135 = v3133
	goto L960
L965:
	;
	if v3138 != 0 {
		goto L966
	} else {
		goto L967
	}
L966:
	;
	if v3090 != 0 {
		goto L969
	} else {
		goto L970
	}
L967:
	;
	goto L968
L968:
	;
	F_MemoryContextReset(m, v2750)
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		goto L9
	} else {
		goto L1003
	}
L969:
	;
	F_errcode(m, v3090)
	mBase = m.M
	v3141 = m.ExcPending
	if v3141 != 0 {
		goto L9
	} else {
		goto L972
	}
L970:
	;
	goto L971
L971:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+256)) = v3135
	F_errmsg_internal(m, int32(_a_F_exec_stmts_32), v32+int32(256))
	mBase = m.M
	v3147 = m.ExcPending
	if v3147 != 0 {
		goto L9
	} else {
		goto L973
	}
L972:
	;
	goto L971
L973:
	;
	if v3071 != 0 {
		goto L974
	} else {
		goto L975
	}
L974:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+240)) = v3071
	F_errdetail_internal(m, int32(_a_F_exec_stmts_32), v32+int32(240))
	mBase = m.M
	v3153 = m.ExcPending
	if v3153 != 0 {
		goto L9
	} else {
		goto L977
	}
L975:
	;
	goto L976
L976:
	;
	if v3070 != 0 {
		goto L978
	} else {
		goto L979
	}
L977:
	;
	goto L976
L978:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+224)) = v3070
	F_errhint(m, int32(_a_F_exec_stmts_32), v32+int32(224))
	mBase = m.M
	v3159 = m.ExcPending
	if v3159 != 0 {
		goto L9
	} else {
		goto L981
	}
L979:
	;
	goto L980
L980:
	;
	if v3069 != 0 {
		goto L982
	} else {
		goto L983
	}
L981:
	;
	goto L980
L982:
	;
	F_err_generic_string(m, int32(99), v3069)
	mBase = m.M
	v3162 = m.ExcPending
	if v3162 != 0 {
		goto L9
	} else {
		goto L985
	}
L983:
	;
	goto L984
L984:
	;
	if v3066 != 0 {
		goto L986
	} else {
		goto L987
	}
L985:
	;
	goto L984
L986:
	;
	F_err_generic_string(m, int32(110), v3066)
	mBase = m.M
	v3165 = m.ExcPending
	if v3165 != 0 {
		goto L9
	} else {
		goto L989
	}
L987:
	;
	goto L988
L988:
	;
	if v3063 != 0 {
		goto L990
	} else {
		goto L991
	}
L989:
	;
	goto L988
L990:
	;
	F_err_generic_string(m, int32(100), v3063)
	mBase = m.M
	v3168 = m.ExcPending
	if v3168 != 0 {
		goto L9
	} else {
		goto L993
	}
L991:
	;
	goto L992
L992:
	;
	if v3064 != 0 {
		goto L994
	} else {
		goto L995
	}
L993:
	;
	goto L992
L994:
	;
	F_err_generic_string(m, int32(116), v3064)
	mBase = m.M
	v3171 = m.ExcPending
	if v3171 != 0 {
		goto L9
	} else {
		goto L997
	}
L995:
	;
	goto L996
L996:
	;
	if v3065 != 0 {
		goto L998
	} else {
		goto L999
	}
L997:
	;
	goto L996
L998:
	;
	F_err_generic_string(m, int32(115), v3065)
	mBase = m.M
	v3174 = m.ExcPending
	if v3174 != 0 {
		goto L9
	} else {
		goto L1001
	}
L999:
	;
	goto L1000
L1000:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3923), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v3179 = m.ExcPending
	if v3179 != 0 {
		goto L9
	} else {
		goto L1002
	}
L1001:
	;
	goto L1000
L1002:
	;
	goto L968
L1003:
	;
	v5191 = int32(0)
	goto L21
L1004:
	;
	v5191 = int32(0)
	goto L21
L1005:
	;
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3189 = v32 + int32(600)
	v3194 = F_exec_eval_expr(m, l0, v3187, v3189, v32+int32(608), v32+int32(636))
	mBase = m.M
	v3195 = m.ExcPending
	if v3195 != 0 {
		goto L9
	} else {
		goto L1006
	}
L1006:
	;
	v3196 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v3197 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v3200 = F_exec_cast_value(m, l0, v3194, v3189, v3196, v3197, int32(16), int32(-1))
	mBase = m.M
	v3201 = m.ExcPending
	if v3201 != 0 {
		goto L9
	} else {
		goto L1007
	}
L1007:
	;
	v3202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3202 != 0 {
		goto L1008
	} else {
		goto L1009
	}
L1008:
	;
	F_SPI_freetuptable(m, v3202)
	mBase = m.M
	v3204 = m.ExcPending
	if v3204 != 0 {
		goto L9
	} else {
		goto L1011
	}
L1009:
	;
	goto L1010
L1010:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v3207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3207 != 0 {
		goto L1012
	} else {
		goto L1013
	}
L1011:
	;
	goto L1010
L1012:
	;
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(v3207)+20))
	F_MemoryContextReset(m, v3208)
	mBase = m.M
	v3210 = m.ExcPending
	if v3210 != 0 {
		goto L9
	} else {
		goto L1015
	}
L1013:
	;
	goto L1014
L1014:
	;
	v3211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v3200 != 0 {
		goto L1016
	} else {
		goto L1017
	}
L1015:
	;
	goto L1014
L1016:
	;
	v3213 = v3211
	goto L1018
L1017:
	;
	v3213 = int32(1)
	goto L1018
L1018:
	;
	if v3213 == int32(0) {
		goto L1004
	} else {
		goto L1019
	}
L1019:
	;
	v3216 = int32(0)
	v3217 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v3217 == v3216 {
		v3255 = v3216
		goto L1020
	} else {
		goto L1021
	}
L1020:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3261 = m.ExcPending
	if v3261 != 0 {
		goto L9
	} else {
		goto L1026
	}
L1021:
	;
	v3226 = F_exec_eval_expr(m, l0, v3217, v32+int32(600), v32+int32(608), v32+int32(636))
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L9
	} else {
		goto L1022
	}
L1022:
	;
	v3228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v3228 != 0 {
		v3255 = v3216
		goto L1020
	} else {
		goto L1023
	}
L1023:
	;
	v3229 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v3230 = m.G0
	v3232 = v3230 - int32(16)
	m.G0 = v3232
	v3234 = int32(_a_F_exec_stmts_1)
	v3235 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v3237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v3238 = *(*int32)(unsafe.Add(mBase, uint32(v3237)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3238
	F_getTypeOutputInfo(m, v3229, v3232+int32(12), v3232+int32(11))
	mBase = m.M
	v3245 = m.ExcPending
	if v3245 != 0 {
		goto L9
	} else {
		goto L1024
	}
L1024:
	;
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(v3232)+12))
	v3247 = F_OidOutputFunctionCall(m, v3246, v3226)
	mBase = m.M
	v3248 = m.ExcPending
	if v3248 != 0 {
		goto L9
	} else {
		goto L1025
	}
L1025:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3235
	m.G0 = v3232 + int32(16)
	v3255 = v3247
	goto L1020
L1026:
	;
	F_errcode(m, int32(67108896))
	mBase = m.M
	v3264 = m.ExcPending
	if v3264 != 0 {
		goto L9
	} else {
		goto L1027
	}
L1027:
	;
	if v3255 != 0 {
		goto L1029
	} else {
		goto L1030
	}
L1028:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3968), int32(_a_F_exec_stmts_33))
	mBase = m.M
	v3279 = m.ExcPending
	if v3279 != 0 {
		goto L9
	} else {
		goto L1034
	}
L1029:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+432)) = v3255
	F_errmsg_internal(m, int32(_a_F_exec_stmts_32), v32+int32(432))
	mBase = m.M
	v3270 = m.ExcPending
	if v3270 != 0 {
		goto L9
	} else {
		goto L1032
	}
L1030:
	;
	goto L1031
L1031:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_34), int32(0))
	mBase = m.M
	v3274 = m.ExcPending
	if v3274 != 0 {
		goto L9
	} else {
		goto L1033
	}
L1032:
	;
	goto L1028
L1033:
	;
	goto L1028
L1034:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1035:
	;
	v5191 = int32(0)
	goto L21
L1036:
	;
	v3289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3294 = F_AllocSetContextCreateInternal(m, v3289, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v3295 = m.ExcPending
	if v3295 != 0 {
		goto L9
	} else {
		goto L1039
	}
L1037:
	;
	v3297 = v3286
	goto L1038
L1038:
	;
	v3298 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3305 = F_exec_eval_expr(m, l0, v3298, v32+int32(607), v32+int32(636), v32+int32(600))
	mBase = m.M
	v3306 = m.ExcPending
	if v3306 != 0 {
		goto L9
	} else {
		goto L1040
	}
L1039:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v3294
	v3297 = v3294
	goto L1038
L1040:
	;
	v3307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+607)))
	if v3307 == int32(1) {
		goto L44
	} else {
		goto L1041
	}
L1041:
	;
	v3310 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v3311 = int32(_a_F_exec_stmts_1)
	v3312 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v3314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v3315 = *(*int32)(unsafe.Add(mBase, uint32(v3314)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3315
	F_getTypeOutputInfo(m, v3310, v32+int32(608), v32+int32(596))
	mBase = m.M
	v3322 = m.ExcPending
	if v3322 != 0 {
		goto L9
	} else {
		goto L1042
	}
L1042:
	;
	v3323 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v3324 = F_OidOutputFunctionCall(m, v3323, v3305)
	mBase = m.M
	v3325 = m.ExcPending
	if v3325 != 0 {
		goto L9
	} else {
		goto L1043
	}
L1043:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3312
	v3328 = F_MemoryContextStrdup(m, v3297, v3324)
	mBase = m.M
	v3329 = m.ExcPending
	if v3329 != 0 {
		goto L9
	} else {
		goto L1044
	}
L1044:
	;
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3330 != 0 {
		goto L1045
	} else {
		goto L1046
	}
L1045:
	;
	F_SPI_freetuptable(m, v3330)
	mBase = m.M
	v3332 = m.ExcPending
	if v3332 != 0 {
		goto L9
	} else {
		goto L1048
	}
L1046:
	;
	goto L1047
L1047:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v3335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3335 != 0 {
		goto L1049
	} else {
		goto L1050
	}
L1048:
	;
	goto L1047
L1049:
	;
	v3336 = *(*int32)(unsafe.Add(mBase, uint32(v3335)+20))
	F_MemoryContextReset(m, v3336)
	mBase = m.M
	v3338 = m.ExcPending
	if v3338 != 0 {
		goto L9
	} else {
		goto L1052
	}
L1050:
	;
	goto L1051
L1051:
	;
	v3339 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v3340 = F_exec_eval_using_params(m, l0, v3339)
	mBase = m.M
	v3341 = m.ExcPending
	if v3341 != 0 {
		goto L9
	} else {
		goto L1053
	}
L1052:
	;
	goto L1051
L1053:
	;
	v3342 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+608)) = v3342
	*(*int64)(unsafe.Add(mBase, uint32(v32)+624)) = v3342
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v3342
	*(*int32)(unsafe.Add(mBase, uint32(v32)+608)) = v3340
	v3349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+612)) = uint8(v3349)
	v3353 = F_SPI_execute_extended(m, v3328, v32+int32(608))
	mBase = m.M
	v3354 = m.ExcPending
	if v3354 != 0 {
		goto L9
	} else {
		goto L1059
	}
L1054:
	;
	v3428 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmts[6]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v3428
	v3430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+16)))
	if v3430 != int32(1) {
		goto L1077
	} else {
		goto L1078
	}
L1055:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3412 = m.ExcPending
	if v3412 != 0 {
		goto L9
	} else {
		goto L1073
	}
L1056:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3396 = m.ExcPending
	if v3396 != 0 {
		goto L9
	} else {
		goto L1069
	}
L1057:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3380 = m.ExcPending
	if v3380 != 0 {
		goto L9
	} else {
		goto L1065
	}
L1058:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3360 = m.ExcPending
	if v3360 != 0 {
		goto L9
	} else {
		goto L1060
	}
L1059:
	;
	switch v3353 + int32(8) {
	case 0:
		goto L1056
	default:
		goto L1055
	case 6:
		goto L1057
	case 8, 12, 13, 15, 16, 17, 19, 20, 21, 22, 26, 27:
		goto L1054
	case 14:
		goto L1058
	}
L1060:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3363 = m.ExcPending
	if v3363 != 0 {
		goto L9
	} else {
		goto L1061
	}
L1061:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_35), int32(0))
	mBase = m.M
	v3367 = m.ExcPending
	if v3367 != 0 {
		goto L9
	} else {
		goto L1062
	}
L1062:
	;
	F_errhint(m, int32(_a_F_exec_stmts_36), int32(0))
	mBase = m.M
	v3371 = m.ExcPending
	if v3371 != 0 {
		goto L9
	} else {
		goto L1063
	}
L1063:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_37), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3376 = m.ExcPending
	if v3376 != 0 {
		goto L9
	} else {
		goto L1064
	}
L1064:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1065:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3383 = m.ExcPending
	if v3383 != 0 {
		goto L9
	} else {
		goto L1066
	}
L1066:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_39), int32(0))
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L9
	} else {
		goto L1067
	}
L1067:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_40), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L9
	} else {
		goto L1068
	}
L1068:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1069:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L9
	} else {
		goto L1070
	}
L1070:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_41), int32(0))
	mBase = m.M
	v3403 = m.ExcPending
	if v3403 != 0 {
		goto L9
	} else {
		goto L1071
	}
L1071:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_42), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L9
	} else {
		goto L1072
	}
L1072:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1073:
	;
	v3413 = F_SPI_result_code_string(m, v3353)
	mBase = m.M
	v3414 = m.ExcPending
	if v3414 != 0 {
		goto L9
	} else {
		goto L1074
	}
L1074:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+452)) = v3413
	*(*int32)(unsafe.Add(mBase, uint32(v32)+448)) = v3328
	F_errmsg_internal(m, int32(_a_F_exec_stmts_43), v32+int32(448))
	mBase = m.M
	v3421 = m.ExcPending
	if v3421 != 0 {
		goto L9
	} else {
		goto L1075
	}
L1075:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_44), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L9
	} else {
		goto L1076
	}
L1076:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1077:
	;
	v3536 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	F_SPI_freetuptable(m, v3536)
	mBase = m.M
	v3538 = m.ExcPending
	if v3538 != 0 {
		goto L9
	} else {
		goto L1119
	}
L1078:
	;
	v3434 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	if v3434 == int32(0) {
		goto L43
	} else {
		goto L1079
	}
L1079:
	;
	v3437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v3439 = *(*int32)(unsafe.Add(mBase, uint32(v3438)+4))
	v3443 = *(*int32)(unsafe.Add(mBase, uint32(v3437+v3439<<(uint(int32(2))%32))))
	if base.Ui64(v3428) <= base.Ui64(int64(1)) {
		goto L1082
	} else {
		goto L1083
	}
L1080:
	;
	v3519 = *(*int32)(unsafe.Add(mBase, uint32(v3434)))
	F_exec_move_row(m, l0, v3443, v3518, v3519)
	mBase = m.M
	v3521 = m.ExcPending
	if v3521 != 0 {
		goto L9
	} else {
		goto L1112
	}
L1081:
	;
	v3516 = *(*int32)(unsafe.Add(mBase, uint32(v3434)+4))
	v3517 = *(*int32)(unsafe.Add(mBase, uint32(v3516)))
	v3518 = v3517
	goto L1080
L1082:
	;
	if base.I32_wrap_i64(v3428) == int32(1) {
		goto L1081
	} else {
		goto L1085
	}
L1083:
	;
	goto L1084
L1084:
	;
	v3483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+17)))
	if v3483 != int32(1) {
		goto L1081
	} else {
		goto L1099
	}
L1085:
	;
	v3450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+17)))
	if v3450 != int32(1) {
		v3518 = int32(0)
		goto L1080
	} else {
		goto L1086
	}
L1086:
	;
	v3454 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3454)+492)))
	if v3455 == int32(1) {
		goto L1087
	} else {
		goto L1088
	}
L1087:
	;
	v3458 = F_format_preparedparamsdata(m, l0, v3340)
	mBase = m.M
	v3459 = m.ExcPending
	if v3459 != 0 {
		goto L9
	} else {
		goto L1090
	}
L1088:
	;
	v3460 = int32(0)
	goto L1089
L1089:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3464 = m.ExcPending
	if v3464 != 0 {
		goto L9
	} else {
		goto L1091
	}
L1090:
	;
	v3460 = v3458
	goto L1089
L1091:
	;
	F_errcode(m, int32(33554464))
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		goto L9
	} else {
		goto L1092
	}
L1092:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_45), int32(0))
	mBase = m.M
	v3471 = m.ExcPending
	if v3471 != 0 {
		goto L9
	} else {
		goto L1093
	}
L1093:
	;
	if v3460 != 0 {
		goto L1094
	} else {
		goto L1095
	}
L1094:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+480)) = v3460
	F_errdetail_internal(m, int32(_a_F_exec_stmts_46), v32+int32(480))
	mBase = m.M
	v3477 = m.ExcPending
	if v3477 != 0 {
		goto L9
	} else {
		goto L1097
	}
L1095:
	;
	goto L1096
L1096:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_47), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3482 = m.ExcPending
	if v3482 != 0 {
		goto L9
	} else {
		goto L1098
	}
L1097:
	;
	goto L1096
L1098:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1099:
	;
	v3487 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3487)+492)))
	if v3488 == int32(1) {
		goto L1100
	} else {
		goto L1101
	}
L1100:
	;
	v3491 = F_format_preparedparamsdata(m, l0, v3340)
	mBase = m.M
	v3492 = m.ExcPending
	if v3492 != 0 {
		goto L9
	} else {
		goto L1103
	}
L1101:
	;
	v3493 = int32(0)
	goto L1102
L1102:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3497 = m.ExcPending
	if v3497 != 0 {
		goto L9
	} else {
		goto L1104
	}
L1103:
	;
	v3493 = v3491
	goto L1102
L1104:
	;
	F_errcode(m, int32(50331680))
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		goto L9
	} else {
		goto L1105
	}
L1105:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_48), int32(0))
	mBase = m.M
	v3504 = m.ExcPending
	if v3504 != 0 {
		goto L9
	} else {
		goto L1106
	}
L1106:
	;
	if v3493 != 0 {
		goto L1107
	} else {
		goto L1108
	}
L1107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+464)) = v3493
	F_errdetail_internal(m, int32(_a_F_exec_stmts_46), v32+int32(464))
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L9
	} else {
		goto L1110
	}
L1108:
	;
	goto L1109
L1109:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_49), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3515 = m.ExcPending
	if v3515 != 0 {
		goto L9
	} else {
		goto L1111
	}
L1110:
	;
	goto L1109
L1111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1112:
	;
	v3522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3522 != 0 {
		goto L1113
	} else {
		goto L1114
	}
L1113:
	;
	F_SPI_freetuptable(m, v3522)
	mBase = m.M
	v3524 = m.ExcPending
	if v3524 != 0 {
		goto L9
	} else {
		goto L1116
	}
L1114:
	;
	goto L1115
L1115:
	;
	v3525 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v3525
	v3527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3527 == v3525 {
		goto L1077
	} else {
		goto L1117
	}
L1116:
	;
	goto L1115
L1117:
	;
	v3530 = *(*int32)(unsafe.Add(mBase, uint32(v3527)+20))
	F_MemoryContextReset(m, v3530)
	mBase = m.M
	v3532 = m.ExcPending
	if v3532 != 0 {
		goto L9
	} else {
		goto L1118
	}
L1118:
	;
	goto L1077
L1119:
	;
	F_MemoryContextReset(m, v3297)
	mBase = m.M
	v3540 = m.ExcPending
	if v3540 != 0 {
		goto L9
	} else {
		goto L1120
	}
L1120:
	;
	v5191 = int32(0)
	goto L21
L1121:
	;
	v3549 = F_exec_for_query(m, l0, v89, v3546, int32(1))
	mBase = m.M
	v3550 = m.ExcPending
	if v3550 != 0 {
		goto L9
	} else {
		goto L1122
	}
L1122:
	;
	F_SPI_cursor_close(m, v3546)
	mBase = m.M
	v3552 = m.ExcPending
	if v3552 != 0 {
		goto L9
	} else {
		goto L1123
	}
L1123:
	;
	v5191 = v3549
	goto L21
L1124:
	;
	v3564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v3564 == int32(0) {
		goto L1127
	} else {
		goto L1128
	}
L1125:
	;
	v3588 = v3553
	v3589 = v3553
	goto L1126
L1126:
	;
	v3590 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	if v3590 != 0 {
		goto L1135
	} else {
		goto L1136
	}
L1127:
	;
	v3567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3572 = F_AllocSetContextCreateInternal(m, v3567, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v3573 = m.ExcPending
	if v3573 != 0 {
		goto L9
	} else {
		goto L1130
	}
L1128:
	;
	v3575 = v3564
	goto L1129
L1129:
	;
	v3576 = int32(_a_F_exec_stmts_1)
	v3577 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3575
	v3580 = *(*int32)(unsafe.Add(mBase, uint32(v3560)+40))
	v3581 = F_text_to_cstring(m, v3580)
	mBase = m.M
	v3582 = m.ExcPending
	if v3582 != 0 {
		goto L9
	} else {
		goto L1131
	}
L1130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v3572
	v3575 = v3572
	goto L1129
L1131:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3577
	v3585 = F_GetPortalByName(m, v3581)
	mBase = m.M
	v3586 = m.ExcPending
	if v3586 != 0 {
		goto L9
	} else {
		goto L1132
	}
L1132:
	;
	if v3585 != 0 {
		goto L42
	} else {
		goto L1133
	}
L1133:
	;
	v3588 = v3575
	v3589 = v3581
	goto L1126
L1134:
	;
	v3650 = *(*int32)(unsafe.Add(mBase, uint32(v3647)+28))
	if v3650 == int32(0) {
		goto L1160
	} else {
		goto L1161
	}
L1135:
	;
	v3591 = *(*int32)(unsafe.Add(mBase, uint32(v3590)+24))
	if v3591 != 0 {
		v3647 = v3590
		goto L1134
	} else {
		goto L1138
	}
L1136:
	;
	goto L1137
L1137:
	;
	v3595 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v3595 != 0 {
		goto L1140
	} else {
		goto L1141
	}
L1138:
	;
	v3592 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	F_exec_prepare_plan(m, l0, v3590, v3592)
	mBase = m.M
	v3594 = m.ExcPending
	if v3594 != 0 {
		goto L9
	} else {
		goto L1139
	}
L1139:
	;
	v3647 = v3590
	goto L1134
L1140:
	;
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	v3597 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v3598 = F_exec_dynquery_with_params(m, l0, v3595, v3596, v3589, v3597)
	mBase = m.M
	v3599 = m.ExcPending
	if v3599 != 0 {
		goto L9
	} else {
		goto L1143
	}
L1141:
	;
	goto L1142
L1142:
	;
	v3612 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v3612 != 0 {
		goto L1151
	} else {
		goto L1152
	}
L1143:
	;
	if v3589 != 0 {
		goto L1144
	} else {
		goto L1145
	}
L1144:
	;
	v5191 = int32(0)
	goto L21
L1145:
	;
	goto L1146
L1146:
	;
	v3601 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	F_exec_check_assignable(m, l0, v3601)
	mBase = m.M
	v3603 = m.ExcPending
	if v3603 != 0 {
		goto L9
	} else {
		goto L1147
	}
L1147:
	;
	v3605 = *(*int32)(unsafe.Add(mBase, uint32(v3598)))
	v3606 = F_cstring_to_text(m, v3605)
	mBase = m.M
	v3607 = m.ExcPending
	if v3607 != 0 {
		goto L9
	} else {
		goto L1148
	}
L1148:
	;
	F_assign_simple_var(m, l0, v3560, v3606, int32(0), int32(1))
	mBase = m.M
	v3611 = m.ExcPending
	if v3611 != 0 {
		goto L9
	} else {
		goto L1149
	}
L1149:
	;
	v5191 = int32(0)
	goto L21
L1150:
	;
	v3642 = *(*int32)(unsafe.Add(mBase, uint32(v3560)+28))
	v3643 = *(*int32)(unsafe.Add(mBase, uint32(v3642)+24))
	if v3643 != 0 {
		v3647 = v3642
		goto L1134
	} else {
		goto L1157
	}
L1151:
	;
	v3613 = *(*int32)(unsafe.Add(mBase, uint32(v3560)+32))
	if v3613 < int32(0) {
		goto L41
	} else {
		goto L1154
	}
L1152:
	;
	goto L1153
L1153:
	;
	v3637 = *(*int32)(unsafe.Add(mBase, uint32(v3560)+32))
	if int32(0) <= v3637 {
		goto L40
	} else {
		goto L1156
	}
L1154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+608)) = int32(16)
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v3623 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+626)) = uint8(v3623)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+620)) = v3612
	*(*int32)(unsafe.Add(mBase, uint32(v32)+612)) = v3622
	v3627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3631 = *(*int32)(unsafe.Add(mBase, uint32(v3627+v3613<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+628)) = v3631
	F_exec_stmt_execsql(m, l0, v32+int32(608))
	mBase = m.M
	v3636 = m.ExcPending
	if v3636 != 0 {
		goto L9
	} else {
		goto L1155
	}
L1155:
	;
	goto L1150
L1156:
	;
	goto L1150
L1157:
	;
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v3560)+36))
	F_exec_prepare_plan(m, l0, v3642, v3644)
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		goto L9
	} else {
		goto L1158
	}
L1158:
	;
	v3647 = v3642
	goto L1134
L1159:
	;
	v3657 = *(*int32)(unsafe.Add(mBase, uint32(v3647)+24))
	v3658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	v3659 = F_SPI_cursor_open_internal(m, v3589, v3657, v3656, v3658)
	mBase = m.M
	v3660 = m.ExcPending
	if v3660 != 0 {
		goto L9
	} else {
		goto L1163
	}
L1160:
	;
	v3656 = int32(0)
	goto L1159
L1161:
	;
	goto L1162
L1162:
	;
	v3654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3654)+20)) = v3647
	v3656 = v3654
	goto L1159
L1163:
	;
	if v3659 == int32(0) {
		goto L39
	} else {
		goto L1164
	}
L1164:
	;
	if v3589 == int32(0) {
		goto L1165
	} else {
		goto L1166
	}
L1165:
	;
	v3665 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	F_exec_check_assignable(m, l0, v3665)
	mBase = m.M
	v3667 = m.ExcPending
	if v3667 != 0 {
		goto L9
	} else {
		goto L1168
	}
L1166:
	;
	goto L1167
L1167:
	;
	v3675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3675 != 0 {
		goto L1171
	} else {
		goto L1172
	}
L1168:
	;
	v3668 = *(*int32)(unsafe.Add(mBase, uint32(v3659)))
	v3669 = F_cstring_to_text(m, v3668)
	mBase = m.M
	v3670 = m.ExcPending
	if v3670 != 0 {
		goto L9
	} else {
		goto L1169
	}
L1169:
	;
	F_assign_simple_var(m, l0, v3560, v3669, int32(0), int32(1))
	mBase = m.M
	v3674 = m.ExcPending
	if v3674 != 0 {
		goto L9
	} else {
		goto L1170
	}
L1170:
	;
	goto L1167
L1171:
	;
	F_SPI_freetuptable(m, v3675)
	mBase = m.M
	v3677 = m.ExcPending
	if v3677 != 0 {
		goto L9
	} else {
		goto L1174
	}
L1172:
	;
	goto L1173
L1173:
	;
	v3678 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v3678
	v3681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3681 != 0 {
		goto L1175
	} else {
		goto L1176
	}
L1174:
	;
	goto L1173
L1175:
	;
	v3682 = *(*int32)(unsafe.Add(mBase, uint32(v3681)+20))
	F_MemoryContextReset(m, v3682)
	mBase = m.M
	v3684 = m.ExcPending
	if v3684 != 0 {
		goto L9
	} else {
		goto L1178
	}
L1176:
	;
	goto L1177
L1177:
	;
	if v3588 == int32(0) {
		v5191 = v3678
		goto L21
	} else {
		goto L1179
	}
L1178:
	;
	goto L1177
L1179:
	;
	F_MemoryContextReset(m, v3588)
	mBase = m.M
	v3688 = m.ExcPending
	if v3688 != 0 {
		goto L9
	} else {
		goto L1180
	}
L1180:
	;
	v5191 = v3678
	goto L21
L1181:
	;
	v3698 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v3699 = int32(_a_F_exec_stmts_1)
	v3700 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v3702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v3703 = *(*int32)(unsafe.Add(mBase, uint32(v3702)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3703
	v3705 = *(*int32)(unsafe.Add(mBase, uint32(v3694)+40))
	v3706 = F_text_to_cstring(m, v3705)
	mBase = m.M
	v3707 = m.ExcPending
	if v3707 != 0 {
		goto L9
	} else {
		goto L1182
	}
L1182:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3700
	v3710 = F_GetPortalByName(m, v3706)
	mBase = m.M
	v3711 = m.ExcPending
	if v3711 != 0 {
		goto L9
	} else {
		goto L1183
	}
L1183:
	;
	if v3710 == int32(0) {
		goto L37
	} else {
		goto L1184
	}
L1184:
	;
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v3714 == int32(0) {
		v3746 = v3698
		goto L1185
	} else {
		goto L1186
	}
L1185:
	;
	v3747 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v3748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+32)))
	if v3748 == int32(0) {
		goto L1197
	} else {
		goto L1198
	}
L1186:
	;
	v3718 = v32 + int32(600)
	v3723 = F_exec_eval_expr(m, l0, v3714, v3718, v32+int32(608), v32+int32(636))
	mBase = m.M
	v3724 = m.ExcPending
	if v3724 != 0 {
		goto L9
	} else {
		goto L1187
	}
L1187:
	;
	v3725 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v3726 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v3729 = F_exec_cast_value(m, l0, v3723, v3718, v3725, v3726, int32(23), int32(-1))
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L9
	} else {
		goto L1188
	}
L1188:
	;
	v3731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+600)))
	if v3731 == int32(1) {
		goto L36
	} else {
		goto L1189
	}
L1189:
	;
	v3734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3734 != 0 {
		goto L1190
	} else {
		goto L1191
	}
L1190:
	;
	F_SPI_freetuptable(m, v3734)
	mBase = m.M
	v3736 = m.ExcPending
	if v3736 != 0 {
		goto L9
	} else {
		goto L1193
	}
L1191:
	;
	goto L1192
L1192:
	;
	v3737 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v3737
	v3739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3739 == v3737 {
		v3746 = v3729
		goto L1185
	} else {
		goto L1194
	}
L1193:
	;
	goto L1192
L1194:
	;
	v3742 = *(*int32)(unsafe.Add(mBase, uint32(v3739)+20))
	F_MemoryContextReset(m, v3742)
	mBase = m.M
	v3744 = m.ExcPending
	if v3744 != 0 {
		goto L9
	} else {
		goto L1195
	}
L1195:
	;
	v3746 = v3729
	goto L1185
L1196:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v3795
	v3797 = int32(0)
	v3798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v3803 = *(*int32)(unsafe.Add(mBase, uint32(v3798+v3799<<(uint(int32(2))%32))))
	F_assign_simple_var(m, l0, v3803, base.B2i32(v3795 != int64(0)), v3797, v3797)
	mBase = m.M
	v3809 = m.ExcPending
	if v3809 != 0 {
		goto L9
	} else {
		goto L1216
	}
L1197:
	;
	v3752 = F_CreateDestReceiver(m, int32(5))
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L9
	} else {
		goto L1200
	}
L1198:
	;
	goto L1199
L1199:
	;
	v3788 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[21]))
	F__SPI_cursor_operation(m, v3710, v3747, v3746, v3788)
	mBase = m.M
	v3790 = m.ExcPending
	if v3790 != 0 {
		goto L9
	} else {
		goto L1215
	}
L1200:
	;
	F__SPI_cursor_operation(m, v3710, v3747, v3746, v3752)
	mBase = m.M
	v3755 = m.ExcPending
	if v3755 != 0 {
		goto L9
	} else {
		goto L1201
	}
L1201:
	;
	v3757 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	v3758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3759 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3760 = *(*int32)(unsafe.Add(mBase, uint32(v3759)+4))
	v3764 = *(*int32)(unsafe.Add(mBase, uint32(v3758+v3760<<(uint(int32(2))%32))))
	v3766 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmts[6]))
	if v3766 == int64(0) {
		goto L1202
	} else {
		goto L1203
	}
L1202:
	;
	v3772 = int32(0)
	goto L1204
L1203:
	;
	v3770 = *(*int32)(unsafe.Add(mBase, uint32(v3757)+4))
	v3771 = *(*int32)(unsafe.Add(mBase, uint32(v3770)))
	v3772 = v3771
	goto L1204
L1204:
	;
	v3773 = *(*int32)(unsafe.Add(mBase, uint32(v3757)))
	F_exec_move_row(m, l0, v3764, v3772, v3773)
	mBase = m.M
	v3775 = m.ExcPending
	if v3775 != 0 {
		goto L9
	} else {
		goto L1205
	}
L1205:
	;
	v3776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3776 != 0 {
		goto L1206
	} else {
		goto L1207
	}
L1206:
	;
	F_SPI_freetuptable(m, v3776)
	mBase = m.M
	v3778 = m.ExcPending
	if v3778 != 0 {
		goto L9
	} else {
		goto L1209
	}
L1207:
	;
	goto L1208
L1208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v3781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3781 != 0 {
		goto L1210
	} else {
		goto L1211
	}
L1209:
	;
	goto L1208
L1210:
	;
	v3782 = *(*int32)(unsafe.Add(mBase, uint32(v3781)+20))
	F_MemoryContextReset(m, v3782)
	mBase = m.M
	v3784 = m.ExcPending
	if v3784 != 0 {
		goto L9
	} else {
		goto L1213
	}
L1211:
	;
	goto L1212
L1212:
	;
	F_SPI_freetuptable(m, v3757)
	mBase = m.M
	v3786 = m.ExcPending
	if v3786 != 0 {
		goto L9
	} else {
		goto L1214
	}
L1213:
	;
	goto L1212
L1214:
	;
	v3795 = v3766
	goto L1196
L1215:
	;
	v3792 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmts[6]))
	v3795 = v3792
	goto L1196
L1216:
	;
	v5191 = v3797
	goto L21
L1217:
	;
	v3819 = int32(_a_F_exec_stmts_1)
	v3820 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v3822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v3823 = *(*int32)(unsafe.Add(mBase, uint32(v3822)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3823
	v3825 = *(*int32)(unsafe.Add(mBase, uint32(v3815)+40))
	v3826 = F_text_to_cstring(m, v3825)
	mBase = m.M
	v3827 = m.ExcPending
	if v3827 != 0 {
		goto L9
	} else {
		goto L1218
	}
L1218:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3820
	v3830 = F_GetPortalByName(m, v3826)
	mBase = m.M
	v3831 = m.ExcPending
	if v3831 != 0 {
		goto L9
	} else {
		goto L1219
	}
L1219:
	;
	if v3830 == int32(0) {
		goto L34
	} else {
		goto L1220
	}
L1220:
	;
	F_SPI_cursor_close(m, v3830)
	mBase = m.M
	v3835 = m.ExcPending
	if v3835 != 0 {
		goto L9
	} else {
		goto L1221
	}
L1221:
	;
	v5191 = int32(0)
	goto L21
L1222:
	;
	goto L33
L1223:
	;
	F__SPI_commit(m, int32(1))
	mBase = m.M
	v3842 = m.ExcPending
	if v3842 != 0 {
		goto L9
	} else {
		goto L1226
	}
L1224:
	;
	goto L1225
L1225:
	;
	F__SPI_commit(m, int32(0))
	mBase = m.M
	v3845 = m.ExcPending
	if v3845 != 0 {
		goto L9
	} else {
		goto L1227
	}
L1226:
	;
	goto L1222
L1227:
	;
	goto L1222
L1228:
	;
	goto L33
L1229:
	;
	F__SPI_rollback(m, int32(1))
	mBase = m.M
	v3851 = m.ExcPending
	if v3851 != 0 {
		goto L9
	} else {
		goto L1232
	}
L1230:
	;
	goto L1231
L1231:
	;
	F__SPI_rollback(m, int32(0))
	mBase = m.M
	v3854 = m.ExcPending
	if v3854 != 0 {
		goto L9
	} else {
		goto L1233
	}
L1232:
	;
	goto L1228
L1233:
	;
	goto L1228
L1234:
	;
	v3860 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v3860
	F_errmsg_internal(m, int32(_a_F_exec_stmts_50), v32)
	mBase = m.M
	v3864 = m.ExcPending
	if v3864 != 0 {
		goto L9
	} else {
		goto L1235
	}
L1235:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2138), int32(_a_F_exec_stmts_51))
	mBase = m.M
	v3869 = m.ExcPending
	if v3869 != 0 {
		goto L9
	} else {
		goto L1236
	}
L1236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1237:
	;
	v5191 = v3870
	goto L21
L1238:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_52), int32(0))
	mBase = m.M
	v3880 = m.ExcPending
	if v3880 != 0 {
		goto L9
	} else {
		goto L1239
	}
L1239:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2315), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v3885 = m.ExcPending
	if v3885 != 0 {
		goto L9
	} else {
		goto L1240
	}
L1240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1241:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_52), int32(0))
	mBase = m.M
	v3893 = m.ExcPending
	if v3893 != 0 {
		goto L9
	} else {
		goto L1242
	}
L1242:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2319), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v3898 = m.ExcPending
	if v3898 != 0 {
		goto L9
	} else {
		goto L1243
	}
L1243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1244:
	;
	v3903 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v3903
	F_errmsg_internal(m, int32(_a_F_exec_stmts_53), v32+int32(16))
	mBase = m.M
	v3909 = m.ExcPending
	if v3909 != 0 {
		goto L9
	} else {
		goto L1245
	}
L1245:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2327), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v3914 = m.ExcPending
	if v3914 != 0 {
		goto L9
	} else {
		goto L1246
	}
L1246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1247:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3921 = m.ExcPending
	if v3921 != 0 {
		goto L9
	} else {
		goto L1248
	}
L1248:
	;
	v3922 = *(*int32)(unsafe.Add(mBase, uint32(v32)+636))
	v3926 = *(*int32)(unsafe.Add(mBase, uint32(v3922+v225<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v3926
	F_errmsg(m, int32(_a_F_exec_stmts_54), v32-int32(-64))
	mBase = m.M
	v3932 = m.ExcPending
	if v3932 != 0 {
		goto L9
	} else {
		goto L1249
	}
L1249:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2383), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v3937 = m.ExcPending
	if v3937 != 0 {
		goto L9
	} else {
		goto L1250
	}
L1250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1251:
	;
	v3942 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v3943 = F_SPI_result_code_string(m, v412)
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L9
	} else {
		goto L1252
	}
L1252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v3943
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v3942
	F_errmsg_internal(m, int32(_a_F_exec_stmts_27), v32+int32(32))
	mBase = m.M
	v3951 = m.ExcPending
	if v3951 != 0 {
		goto L9
	} else {
		goto L1253
	}
L1253:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2246), int32(_a_F_exec_stmts_55))
	mBase = m.M
	v3956 = m.ExcPending
	if v3956 != 0 {
		goto L9
	} else {
		goto L1254
	}
L1254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1255:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_56), int32(0))
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		goto L9
	} else {
		goto L1256
	}
L1256:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2270), int32(_a_F_exec_stmts_55))
	mBase = m.M
	v3969 = m.ExcPending
	if v3969 != 0 {
		goto L9
	} else {
		goto L1257
	}
L1257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1258:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_57), int32(0))
	mBase = m.M
	v3977 = m.ExcPending
	if v3977 != 0 {
		goto L9
	} else {
		goto L1259
	}
L1259:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2275), int32(_a_F_exec_stmts_55))
	mBase = m.M
	v3982 = m.ExcPending
	if v3982 != 0 {
		goto L9
	} else {
		goto L1260
	}
L1260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1261:
	;
	F_errcode(m, int32(33557120))
	mBase = m.M
	v3989 = m.ExcPending
	if v3989 != 0 {
		goto L9
	} else {
		goto L1262
	}
L1262:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_58), int32(0))
	mBase = m.M
	v3993 = m.ExcPending
	if v3993 != 0 {
		goto L9
	} else {
		goto L1263
	}
L1263:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2423), int32(_a_F_exec_stmts_15))
	mBase = m.M
	v3998 = m.ExcPending
	if v3998 != 0 {
		goto L9
	} else {
		goto L1264
	}
L1264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1265:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4005 = m.ExcPending
	if v4005 != 0 {
		goto L9
	} else {
		goto L1266
	}
L1266:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_59), int32(0))
	mBase = m.M
	v4009 = m.ExcPending
	if v4009 != 0 {
		goto L9
	} else {
		goto L1267
	}
L1267:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2723), int32(_a_F_exec_stmts_60))
	mBase = m.M
	v4014 = m.ExcPending
	if v4014 != 0 {
		goto L9
	} else {
		goto L1268
	}
L1268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1269:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4021 = m.ExcPending
	if v4021 != 0 {
		goto L9
	} else {
		goto L1270
	}
L1270:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_61), int32(0))
	mBase = m.M
	v4025 = m.ExcPending
	if v4025 != 0 {
		goto L9
	} else {
		goto L1271
	}
L1271:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2739), int32(_a_F_exec_stmts_60))
	mBase = m.M
	v4030 = m.ExcPending
	if v4030 != 0 {
		goto L9
	} else {
		goto L1272
	}
L1272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1273:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4037 = m.ExcPending
	if v4037 != 0 {
		goto L9
	} else {
		goto L1274
	}
L1274:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_62), int32(0))
	mBase = m.M
	v4041 = m.ExcPending
	if v4041 != 0 {
		goto L9
	} else {
		goto L1275
	}
L1275:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2757), int32(_a_F_exec_stmts_60))
	mBase = m.M
	v4046 = m.ExcPending
	if v4046 != 0 {
		goto L9
	} else {
		goto L1276
	}
L1276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1277:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v4053 = m.ExcPending
	if v4053 != 0 {
		goto L9
	} else {
		goto L1278
	}
L1278:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_63), int32(0))
	mBase = m.M
	v4057 = m.ExcPending
	if v4057 != 0 {
		goto L9
	} else {
		goto L1279
	}
L1279:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2763), int32(_a_F_exec_stmts_60))
	mBase = m.M
	v4062 = m.ExcPending
	if v4062 != 0 {
		goto L9
	} else {
		goto L1280
	}
L1280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1281:
	;
	F_errcode(m, int32(50462852))
	mBase = m.M
	v4069 = m.ExcPending
	if v4069 != 0 {
		goto L9
	} else {
		goto L1282
	}
L1282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v1734
	F_errmsg(m, int32(_a_F_exec_stmts_64), v32+int32(112))
	mBase = m.M
	v4075 = m.ExcPending
	if v4075 != 0 {
		goto L9
	} else {
		goto L1283
	}
L1283:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2897), int32(_a_F_exec_stmts_65))
	mBase = m.M
	v4080 = m.ExcPending
	if v4080 != 0 {
		goto L9
	} else {
		goto L1284
	}
L1284:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1285:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4087 = m.ExcPending
	if v4087 != 0 {
		goto L9
	} else {
		goto L1286
	}
L1286:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_66), int32(0))
	mBase = m.M
	v4091 = m.ExcPending
	if v4091 != 0 {
		goto L9
	} else {
		goto L1287
	}
L1287:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2920), int32(_a_F_exec_stmts_65))
	mBase = m.M
	v4096 = m.ExcPending
	if v4096 != 0 {
		goto L9
	} else {
		goto L1288
	}
L1288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1289:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4103 = m.ExcPending
	if v4103 != 0 {
		goto L9
	} else {
		goto L1290
	}
L1290:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_67), int32(0))
	mBase = m.M
	v4107 = m.ExcPending
	if v4107 != 0 {
		goto L9
	} else {
		goto L1291
	}
L1291:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2939), int32(_a_F_exec_stmts_65))
	mBase = m.M
	v4112 = m.ExcPending
	if v4112 != 0 {
		goto L9
	} else {
		goto L1292
	}
L1292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1293:
	;
	v4118 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[22]))
	v4119 = F_SPI_result_code_string(m, v4118)
	mBase = m.M
	v4120 = m.ExcPending
	if v4120 != 0 {
		goto L9
	} else {
		goto L1294
	}
L1294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v4119
	F_errmsg_internal(m, int32(_a_F_exec_stmts_68), v32+int32(96))
	mBase = m.M
	v4126 = m.ExcPending
	if v4126 != 0 {
		goto L9
	} else {
		goto L1295
	}
L1295:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2961), int32(_a_F_exec_stmts_65))
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L9
	} else {
		goto L1296
	}
L1296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1297:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4138 = m.ExcPending
	if v4138 != 0 {
		goto L9
	} else {
		goto L1298
	}
L1298:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_69), int32(0))
	mBase = m.M
	v4142 = m.ExcPending
	if v4142 != 0 {
		goto L9
	} else {
		goto L1299
	}
L1299:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3030), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4147 = m.ExcPending
	if v4147 != 0 {
		goto L9
	} else {
		goto L1300
	}
L1300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1301:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4154 = m.ExcPending
	if v4154 != 0 {
		goto L9
	} else {
		goto L1302
	}
L1302:
	;
	v4155 = F_format_type_be(m, v1880)
	mBase = m.M
	v4156 = m.ExcPending
	if v4156 != 0 {
		goto L9
	} else {
		goto L1303
	}
L1303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+128)) = v4155
	F_errmsg(m, int32(_a_F_exec_stmts_71), v32+int32(128))
	mBase = m.M
	v4162 = m.ExcPending
	if v4162 != 0 {
		goto L9
	} else {
		goto L1304
	}
L1304:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3046), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4167 = m.ExcPending
	if v4167 != 0 {
		goto L9
	} else {
		goto L1305
	}
L1305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1306:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v4174 = m.ExcPending
	if v4174 != 0 {
		goto L9
	} else {
		goto L1307
	}
L1307:
	;
	v4175 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v4176 = *(*int32)(unsafe.Add(mBase, uint32(v1886)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+148)) = v4176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+144)) = v4175
	F_errmsg(m, int32(_a_F_exec_stmts_72), v32+int32(144))
	mBase = m.M
	v4183 = m.ExcPending
	if v4183 != 0 {
		goto L9
	} else {
		goto L1308
	}
L1308:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3063), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4188 = m.ExcPending
	if v4188 != 0 {
		goto L9
	} else {
		goto L1309
	}
L1309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1310:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4195 = m.ExcPending
	if v4195 != 0 {
		goto L9
	} else {
		goto L1311
	}
L1311:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_73), int32(0))
	mBase = m.M
	v4199 = m.ExcPending
	if v4199 != 0 {
		goto L9
	} else {
		goto L1312
	}
L1312:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3089), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4204 = m.ExcPending
	if v4204 != 0 {
		goto L9
	} else {
		goto L1313
	}
L1313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1314:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4211 = m.ExcPending
	if v4211 != 0 {
		goto L9
	} else {
		goto L1315
	}
L1315:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_74), int32(0))
	mBase = m.M
	v4215 = m.ExcPending
	if v4215 != 0 {
		goto L9
	} else {
		goto L1316
	}
L1316:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3093), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4220 = m.ExcPending
	if v4220 != 0 {
		goto L9
	} else {
		goto L1317
	}
L1317:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1318:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4227 = m.ExcPending
	if v4227 != 0 {
		goto L9
	} else {
		goto L1319
	}
L1319:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_75), int32(0))
	mBase = m.M
	v4231 = m.ExcPending
	if v4231 != 0 {
		goto L9
	} else {
		goto L1320
	}
L1320:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3337), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4236 = m.ExcPending
	if v4236 != 0 {
		goto L9
	} else {
		goto L1321
	}
L1321:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1322:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4243 = m.ExcPending
	if v4243 != 0 {
		goto L9
	} else {
		goto L1323
	}
L1323:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_76), int32(0))
	mBase = m.M
	v4247 = m.ExcPending
	if v4247 != 0 {
		goto L9
	} else {
		goto L1324
	}
L1324:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3378), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4252 = m.ExcPending
	if v4252 != 0 {
		goto L9
	} else {
		goto L1325
	}
L1325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1326:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4259 = m.ExcPending
	if v4259 != 0 {
		goto L9
	} else {
		goto L1327
	}
L1327:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_23), int32(0))
	mBase = m.M
	v4263 = m.ExcPending
	if v4263 != 0 {
		goto L9
	} else {
		goto L1328
	}
L1328:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3437), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4268 = m.ExcPending
	if v4268 != 0 {
		goto L9
	} else {
		goto L1329
	}
L1329:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1330:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4275 = m.ExcPending
	if v4275 != 0 {
		goto L9
	} else {
		goto L1331
	}
L1331:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_20), int32(0))
	mBase = m.M
	v4279 = m.ExcPending
	if v4279 != 0 {
		goto L9
	} else {
		goto L1332
	}
L1332:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3473), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4284 = m.ExcPending
	if v4284 != 0 {
		goto L9
	} else {
		goto L1333
	}
L1333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1334:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4291 = m.ExcPending
	if v4291 != 0 {
		goto L9
	} else {
		goto L1335
	}
L1335:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_76), int32(0))
	mBase = m.M
	v4295 = m.ExcPending
	if v4295 != 0 {
		goto L9
	} else {
		goto L1336
	}
L1336:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3510), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4300 = m.ExcPending
	if v4300 != 0 {
		goto L9
	} else {
		goto L1337
	}
L1337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1338:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4307 = m.ExcPending
	if v4307 != 0 {
		goto L9
	} else {
		goto L1339
	}
L1339:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_77), int32(0))
	mBase = m.M
	v4311 = m.ExcPending
	if v4311 != 0 {
		goto L9
	} else {
		goto L1340
	}
L1340:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3529), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4316 = m.ExcPending
	if v4316 != 0 {
		goto L9
	} else {
		goto L1341
	}
L1341:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1342:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4323 = m.ExcPending
	if v4323 != 0 {
		goto L9
	} else {
		goto L1343
	}
L1343:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_78), int32(0))
	mBase = m.M
	v4327 = m.ExcPending
	if v4327 != 0 {
		goto L9
	} else {
		goto L1344
	}
L1344:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3557), int32(_a_F_exec_stmts_28))
	mBase = m.M
	v4332 = m.ExcPending
	if v4332 != 0 {
		goto L9
	} else {
		goto L1345
	}
L1345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1346:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4339 = m.ExcPending
	if v4339 != 0 {
		goto L9
	} else {
		goto L1347
	}
L1347:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_79), int32(0))
	mBase = m.M
	v4343 = m.ExcPending
	if v4343 != 0 {
		goto L9
	} else {
		goto L1348
	}
L1348:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3630), int32(_a_F_exec_stmts_28))
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L9
	} else {
		goto L1349
	}
L1349:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1350:
	;
	v4353 = F_SPI_result_code_string(m, v2680)
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L9
	} else {
		goto L1351
	}
L1351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+196)) = v4353
	*(*int32)(unsafe.Add(mBase, uint32(v32)+192)) = v2652
	F_errmsg_internal(m, int32(_a_F_exec_stmts_43), v32+int32(192))
	mBase = m.M
	v4361 = m.ExcPending
	if v4361 != 0 {
		goto L9
	} else {
		goto L1352
	}
L1352:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3651), int32(_a_F_exec_stmts_28))
	mBase = m.M
	v4366 = m.ExcPending
	if v4366 != 0 {
		goto L9
	} else {
		goto L1353
	}
L1353:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1355:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_80), int32(0))
	mBase = m.M
	v4376 = m.ExcPending
	if v4376 != 0 {
		goto L9
	} else {
		goto L1356
	}
L1356:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3798), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4381 = m.ExcPending
	if v4381 != 0 {
		goto L9
	} else {
		goto L1357
	}
L1357:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1358:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_80), int32(0))
	mBase = m.M
	v4389 = m.ExcPending
	if v4389 != 0 {
		goto L9
	} else {
		goto L1359
	}
L1359:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3822), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4394 = m.ExcPending
	if v4394 != 0 {
		goto L9
	} else {
		goto L1360
	}
L1360:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1361:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4401 = m.ExcPending
	if v4401 != 0 {
		goto L9
	} else {
		goto L1362
	}
L1362:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_81), int32(0))
	mBase = m.M
	v4405 = m.ExcPending
	if v4405 != 0 {
		goto L9
	} else {
		goto L1363
	}
L1363:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3843), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4410 = m.ExcPending
	if v4410 != 0 {
		goto L9
	} else {
		goto L1364
	}
L1364:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1365:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4417 = m.ExcPending
	if v4417 != 0 {
		goto L9
	} else {
		goto L1366
	}
L1366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+288)) = int32(_a_F_exec_stmts_82)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(288))
	mBase = m.M
	v4424 = m.ExcPending
	if v4424 != 0 {
		goto L9
	} else {
		goto L1367
	}
L1367:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3854), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4429 = m.ExcPending
	if v4429 != 0 {
		goto L9
	} else {
		goto L1368
	}
L1368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1369:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4436 = m.ExcPending
	if v4436 != 0 {
		goto L9
	} else {
		goto L1370
	}
L1370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+304)) = int32(_a_F_exec_stmts_84)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(304))
	mBase = m.M
	v4443 = m.ExcPending
	if v4443 != 0 {
		goto L9
	} else {
		goto L1371
	}
L1371:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3859), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4448 = m.ExcPending
	if v4448 != 0 {
		goto L9
	} else {
		goto L1372
	}
L1372:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1373:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4455 = m.ExcPending
	if v4455 != 0 {
		goto L9
	} else {
		goto L1374
	}
L1374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+320)) = int32(_a_F_exec_stmts_85)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(320))
	mBase = m.M
	v4462 = m.ExcPending
	if v4462 != 0 {
		goto L9
	} else {
		goto L1375
	}
L1375:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3862), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4467 = m.ExcPending
	if v4467 != 0 {
		goto L9
	} else {
		goto L1376
	}
L1376:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1377:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4474 = m.ExcPending
	if v4474 != 0 {
		goto L9
	} else {
		goto L1378
	}
L1378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+336)) = int32(_a_F_exec_stmts_86)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(336))
	mBase = m.M
	v4481 = m.ExcPending
	if v4481 != 0 {
		goto L9
	} else {
		goto L1379
	}
L1379:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3865), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4486 = m.ExcPending
	if v4486 != 0 {
		goto L9
	} else {
		goto L1380
	}
L1380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1381:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4493 = m.ExcPending
	if v4493 != 0 {
		goto L9
	} else {
		goto L1382
	}
L1382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+352)) = int32(_a_F_exec_stmts_87)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(352))
	mBase = m.M
	v4500 = m.ExcPending
	if v4500 != 0 {
		goto L9
	} else {
		goto L1383
	}
L1383:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3868), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4505 = m.ExcPending
	if v4505 != 0 {
		goto L9
	} else {
		goto L1384
	}
L1384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1385:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4512 = m.ExcPending
	if v4512 != 0 {
		goto L9
	} else {
		goto L1386
	}
L1386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+368)) = int32(_a_F_exec_stmts_88)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(368))
	mBase = m.M
	v4519 = m.ExcPending
	if v4519 != 0 {
		goto L9
	} else {
		goto L1387
	}
L1387:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3871), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4524 = m.ExcPending
	if v4524 != 0 {
		goto L9
	} else {
		goto L1388
	}
L1388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1389:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4531 = m.ExcPending
	if v4531 != 0 {
		goto L9
	} else {
		goto L1390
	}
L1390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = int32(_a_F_exec_stmts_89)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(384))
	mBase = m.M
	v4538 = m.ExcPending
	if v4538 != 0 {
		goto L9
	} else {
		goto L1391
	}
L1391:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3874), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4543 = m.ExcPending
	if v4543 != 0 {
		goto L9
	} else {
		goto L1392
	}
L1392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1393:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4550 = m.ExcPending
	if v4550 != 0 {
		goto L9
	} else {
		goto L1394
	}
L1394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = int32(_a_F_exec_stmts_90)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(400))
	mBase = m.M
	v4557 = m.ExcPending
	if v4557 != 0 {
		goto L9
	} else {
		goto L1395
	}
L1395:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3877), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4562 = m.ExcPending
	if v4562 != 0 {
		goto L9
	} else {
		goto L1396
	}
L1396:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1397:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4569 = m.ExcPending
	if v4569 != 0 {
		goto L9
	} else {
		goto L1398
	}
L1398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+416)) = int32(_a_F_exec_stmts_91)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(416))
	mBase = m.M
	v4576 = m.ExcPending
	if v4576 != 0 {
		goto L9
	} else {
		goto L1399
	}
L1399:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3880), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4581 = m.ExcPending
	if v4581 != 0 {
		goto L9
	} else {
		goto L1400
	}
L1400:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1401:
	;
	v4586 = *(*int32)(unsafe.Add(mBase, uint32(v2979)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+272)) = v4586
	F_errmsg_internal(m, int32(_a_F_exec_stmts_92), v32+int32(272))
	mBase = m.M
	v4592 = m.ExcPending
	if v4592 != 0 {
		goto L9
	} else {
		goto L1402
	}
L1402:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3883), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4597 = m.ExcPending
	if v4597 != 0 {
		goto L9
	} else {
		goto L1403
	}
L1403:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1404:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4604 = m.ExcPending
	if v4604 != 0 {
		goto L9
	} else {
		goto L1405
	}
L1405:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_79), int32(0))
	mBase = m.M
	v4608 = m.ExcPending
	if v4608 != 0 {
		goto L9
	} else {
		goto L1406
	}
L1406:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_93), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v4613 = m.ExcPending
	if v4613 != 0 {
		goto L9
	} else {
		goto L1407
	}
L1407:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1408:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4620 = m.ExcPending
	if v4620 != 0 {
		goto L9
	} else {
		goto L1409
	}
L1409:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_94), int32(0))
	mBase = m.M
	v4624 = m.ExcPending
	if v4624 != 0 {
		goto L9
	} else {
		goto L1410
	}
L1410:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_95), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v4629 = m.ExcPending
	if v4629 != 0 {
		goto L9
	} else {
		goto L1411
	}
L1411:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1412:
	;
	F_errcode(m, int32(50462852))
	mBase = m.M
	v4636 = m.ExcPending
	if v4636 != 0 {
		goto L9
	} else {
		goto L1413
	}
L1413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+512)) = v3581
	F_errmsg(m, int32(_a_F_exec_stmts_64), v32+int32(512))
	mBase = m.M
	v4642 = m.ExcPending
	if v4642 != 0 {
		goto L9
	} else {
		goto L1414
	}
L1414:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_96), int32(_a_F_exec_stmts_97))
	mBase = m.M
	v4647 = m.ExcPending
	if v4647 != 0 {
		goto L9
	} else {
		goto L1415
	}
L1415:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1416:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4654 = m.ExcPending
	if v4654 != 0 {
		goto L9
	} else {
		goto L1417
	}
L1417:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_66), int32(0))
	mBase = m.M
	v4658 = m.ExcPending
	if v4658 != 0 {
		goto L9
	} else {
		goto L1418
	}
L1418:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_98), int32(_a_F_exec_stmts_97))
	mBase = m.M
	v4663 = m.ExcPending
	if v4663 != 0 {
		goto L9
	} else {
		goto L1419
	}
L1419:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1420:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4670 = m.ExcPending
	if v4670 != 0 {
		goto L9
	} else {
		goto L1421
	}
L1421:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_67), int32(0))
	mBase = m.M
	v4674 = m.ExcPending
	if v4674 != 0 {
		goto L9
	} else {
		goto L1422
	}
L1422:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_99), int32(_a_F_exec_stmts_97))
	mBase = m.M
	v4679 = m.ExcPending
	if v4679 != 0 {
		goto L9
	} else {
		goto L1423
	}
L1423:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1424:
	;
	v4685 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[22]))
	v4686 = F_SPI_result_code_string(m, v4685)
	mBase = m.M
	v4687 = m.ExcPending
	if v4687 != 0 {
		goto L9
	} else {
		goto L1425
	}
L1425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+496)) = v4686
	F_errmsg_internal(m, int32(_a_F_exec_stmts_68), v32+int32(496))
	mBase = m.M
	v4693 = m.ExcPending
	if v4693 != 0 {
		goto L9
	} else {
		goto L1426
	}
L1426:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_100), int32(_a_F_exec_stmts_97))
	mBase = m.M
	v4698 = m.ExcPending
	if v4698 != 0 {
		goto L9
	} else {
		goto L1427
	}
L1427:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1428:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4705 = m.ExcPending
	if v4705 != 0 {
		goto L9
	} else {
		goto L1429
	}
L1429:
	;
	v4706 = *(*int32)(unsafe.Add(mBase, uint32(v3694)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+528)) = v4706
	F_errmsg(m, int32(_a_F_exec_stmts_101), v32+int32(528))
	mBase = m.M
	v4712 = m.ExcPending
	if v4712 != 0 {
		goto L9
	} else {
		goto L1430
	}
L1430:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_102), int32(_a_F_exec_stmts_103))
	mBase = m.M
	v4717 = m.ExcPending
	if v4717 != 0 {
		goto L9
	} else {
		goto L1431
	}
L1431:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1432:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v4724 = m.ExcPending
	if v4724 != 0 {
		goto L9
	} else {
		goto L1433
	}
L1433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v3706
	F_errmsg(m, int32(_a_F_exec_stmts_104), v32+int32(544))
	mBase = m.M
	v4730 = m.ExcPending
	if v4730 != 0 {
		goto L9
	} else {
		goto L1434
	}
L1434:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_105), int32(_a_F_exec_stmts_103))
	mBase = m.M
	v4735 = m.ExcPending
	if v4735 != 0 {
		goto L9
	} else {
		goto L1435
	}
L1435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1436:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4742 = m.ExcPending
	if v4742 != 0 {
		goto L9
	} else {
		goto L1437
	}
L1437:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_106), int32(0))
	mBase = m.M
	v4746 = m.ExcPending
	if v4746 != 0 {
		goto L9
	} else {
		goto L1438
	}
L1438:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_107), int32(_a_F_exec_stmts_103))
	mBase = m.M
	v4751 = m.ExcPending
	if v4751 != 0 {
		goto L9
	} else {
		goto L1439
	}
L1439:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1440:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4758 = m.ExcPending
	if v4758 != 0 {
		goto L9
	} else {
		goto L1441
	}
L1441:
	;
	v4759 = *(*int32)(unsafe.Add(mBase, uint32(v3815)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v4759
	F_errmsg(m, int32(_a_F_exec_stmts_101), v32+int32(560))
	mBase = m.M
	v4765 = m.ExcPending
	if v4765 != 0 {
		goto L9
	} else {
		goto L1442
	}
L1442:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_108), int32(_a_F_exec_stmts_109))
	mBase = m.M
	v4770 = m.ExcPending
	if v4770 != 0 {
		goto L9
	} else {
		goto L1443
	}
L1443:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1444:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v4777 = m.ExcPending
	if v4777 != 0 {
		goto L9
	} else {
		goto L1445
	}
L1445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+576)) = v3826
	F_errmsg(m, int32(_a_F_exec_stmts_104), v32+int32(576))
	mBase = m.M
	v4783 = m.ExcPending
	if v4783 != 0 {
		goto L9
	} else {
		goto L1446
	}
L1446:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_110), int32(_a_F_exec_stmts_109))
	mBase = m.M
	v4788 = m.ExcPending
	if v4788 != 0 {
		goto L9
	} else {
		goto L1447
	}
L1447:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1448:
	;
	v5191 = int32(0)
	goto L21
L1449:
	;
	v4865 = int32(0)
	v4868 = v1551
	goto L30
L1450:
	;
	goto L1451
L1451:
	;
	v4827 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v4827 == int32(0) {
		goto L1452
	} else {
		goto L1453
	}
L1452:
	;
	v4865 = int32(1)
	v4868 = v1551
	goto L30
L1453:
	;
	goto L1454
L1454:
	;
	v4833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4827))))
	v4836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4823))))
	if base.B2i32(v4833 == int32(0))|base.B2i32(v4833 != v4836) != 0 {
		v4854 = v4833
		v4855 = v4836
		goto L1456
	} else {
		goto L1457
	}
L1455:
	;
	if v4854-v4855 != 0 {
		goto L1462
	} else {
		goto L1463
	}
L1456:
	;
	goto L1455
L1457:
	;
	v4839 = v4827
	v4840 = v4823
	goto L1458
L1458:
	;
	v4843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4840)+1)))
	v4844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4839)+1)))
	if v4844 == int32(0) {
		v4854 = v4844
		v4855 = v4843
		goto L1456
	} else {
		goto L1460
	}
L1459:
	;
	v4854 = v4844
	v4855 = v4843
	goto L1456
L1460:
	;
	v4847 = int32(1)
	if v4844 == v4843 {
		v4839 = v4839 + v4847
		v4840 = v4840 + v4847
		goto L1458
	} else {
		goto L1461
	}
L1461:
	;
	goto L1459
L1462:
	;
	v4865 = int32(1)
	v4868 = v1551
	goto L30
L1463:
	;
	goto L1464
L1464:
	;
	v4858 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v4858
	v4865 = v4858
	v4868 = v1551
	goto L30
L1465:
	;
	v5191 = v4865
	goto L21
L1466:
	;
	v5191 = int32(0)
	goto L21
L1467:
	;
	goto L1468
L1468:
	;
	v4906 = int32(1)
	v4907 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v4907 == int32(0) {
		v5191 = v4906
		goto L21
	} else {
		goto L1469
	}
L1469:
	;
	v4912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4907))))
	v4915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4902))))
	if base.B2i32(v4912 == int32(0))|base.B2i32(v4912 != v4915) != 0 {
		v4933 = v4912
		v4934 = v4915
		goto L1471
	} else {
		goto L1472
	}
L1470:
	;
	if v4933-v4934 != 0 {
		v5191 = v4906
		goto L21
	} else {
		goto L1477
	}
L1471:
	;
	goto L1470
L1472:
	;
	v4918 = v4907
	v4919 = v4902
	goto L1473
L1473:
	;
	v4922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4919)+1)))
	v4923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4918)+1)))
	if v4923 == int32(0) {
		v4933 = v4923
		v4934 = v4922
		goto L1471
	} else {
		goto L1475
	}
L1474:
	;
	v4933 = v4923
	v4934 = v4922
	goto L1471
L1475:
	;
	v4926 = int32(1)
	if v4923 == v4922 {
		v4918 = v4918 + v4926
		v4919 = v4919 + v4926
		goto L1473
	} else {
		goto L1476
	}
L1476:
	;
	goto L1474
L1477:
	;
	goto L28
L1478:
	;
	v4997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+45)))
	if v4997 != int32(1) {
		goto L1481
	} else {
		goto L1482
	}
L1479:
	;
	goto L1480
L1480:
	;
	v5026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+24)))
	if v5026 == int32(0) {
		goto L25
	} else {
		goto L1490
	}
L1481:
	;
	v5019 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+48)) = v5019
	v5021 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1108)+44)) = uint16(v5021)
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+40)) = v5019
	goto L1480
L1482:
	;
	v5000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+44)))
	if v5000 != 0 {
		goto L1483
	} else {
		goto L1484
	}
L1483:
	;
	v5015 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+40))
	F_pfree(m, v5015)
	mBase = m.M
	v5017 = m.ExcPending
	if v5017 != 0 {
		goto L9
	} else {
		goto L1489
	}
L1484:
	;
	v5001 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+24))
	v5002 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5001)+12)))
	if v5002 != int32(_a_F_exec_stmts_16) {
		goto L1483
	} else {
		goto L1485
	}
L1485:
	;
	v5005 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+40))
	v5006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5005))))
	if v5006 != int32(1) {
		goto L1483
	} else {
		goto L1486
	}
L1486:
	;
	v5009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5005)+1)))
	if v5009 != int32(3) {
		goto L1483
	} else {
		goto L1487
	}
L1487:
	;
	F_DeleteExpandedObject(m, v5005)
	mBase = m.M
	v5013 = m.ExcPending
	if v5013 != 0 {
		goto L9
	} else {
		goto L1488
	}
L1488:
	;
	goto L1481
L1489:
	;
	goto L1481
L1490:
	;
	v5060 = v89 + int32(28)
	goto L26
L1491:
	;
	v5191 = v5062
	goto L21
L1492:
	;
	F_errcode(m, int32(2))
	mBase = m.M
	v5070 = m.ExcPending
	if v5070 != 0 {
		goto L9
	} else {
		goto L1493
	}
L1493:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_111), int32(0))
	mBase = m.M
	v5074 = m.ExcPending
	if v5074 != 0 {
		goto L9
	} else {
		goto L1494
	}
L1494:
	;
	F_errhint(m, int32(_a_F_exec_stmts_112), int32(0))
	mBase = m.M
	v5078 = m.ExcPending
	if v5078 != 0 {
		goto L9
	} else {
		goto L1495
	}
L1495:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2630), int32(_a_F_exec_stmts_113))
	mBase = m.M
	v5083 = m.ExcPending
	if v5083 != 0 {
		goto L9
	} else {
		goto L1496
	}
L1496:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1497:
	;
	v5191 = v5146
	goto L21
L1498:
	;
	F_SPI_freetuptable(m, v5177)
	mBase = m.M
	v5179 = m.ExcPending
	if v5179 != 0 {
		goto L9
	} else {
		goto L1501
	}
L1499:
	;
	goto L1500
L1500:
	;
	v5180 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v5180
	v5183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v5183 == v5180 {
		v5191 = v5180
		goto L21
	} else {
		goto L1502
	}
L1501:
	;
	goto L1500
L1502:
	;
	v5186 = *(*int32)(unsafe.Add(mBase, uint32(v5183)+20))
	F_MemoryContextReset(m, v5186)
	mBase = m.M
	v5188 = m.ExcPending
	if v5188 != 0 {
		goto L9
	} else {
		goto L1503
	}
L1503:
	;
	v5191 = v5180
	goto L21
L1504:
	;
	if v5191 == int32(0) {
		goto L1508
	} else {
		goto L1509
	}
L1505:
	;
	v5223 = *(*int32)(unsafe.Add(mBase, uint32(v5220)+16))
	if v5223 == int32(0) {
		v5231 = v5219
		goto L1504
	} else {
		goto L1506
	}
L1506:
	;
	m.T0[v5223].(func(*base.Module, int32, int32))(m, l0, v89)
	mBase = m.M
	v5227 = m.ExcPending
	if v5227 != 0 {
		goto L9
	} else {
		goto L1507
	}
L1507:
	;
	v5229 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[1]))
	v5231 = v5229
	goto L1504
L1508:
	;
	v5235 = v79 + int32(1)
	v5236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5236 <= v5235 {
		goto L2
	} else {
		goto L1511
	}
L1509:
	;
	goto L1510
L1510:
	;
	goto L12
L1511:
	;
	v63 = v5231
	v79 = v5235
	goto L11
}
