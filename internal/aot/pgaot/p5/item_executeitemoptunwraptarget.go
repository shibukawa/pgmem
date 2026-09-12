package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_executeItemOptUnwrapTarget(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int64
	_ = v516
	var v518 int64
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v577 int32
	_ = v577
	var v593 int32
	_ = v593
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v884 int64
	_ = v884
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1030 int64
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1152 int64
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1397 int64
	_ = v1397
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1496 int32
	_ = v1496
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1532 int64
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
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
	var v1555 int32
	_ = v1555
	var v1564 int32
	_ = v1564
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
	var v1575 int32
	_ = v1575
	var v1578 int64
	_ = v1578
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1597 int32
	_ = v1597
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1615 int32
	_ = v1615
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1659 int32
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1705 int32
	_ = v1705
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1737 int32
	_ = v1737
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1789 int64
	_ = v1789
	var v1790 int64
	_ = v1790
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1838 int32
	_ = v1838
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1935 int32
	_ = v1935
	var v1941 int32
	_ = v1941
	var v1943 int64
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1950 int32
	_ = v1950
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1960 int64
	_ = v1960
	var v1962 int64
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1978 int32
	_ = v1978
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2022 int32
	_ = v2022
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2052 int32
	_ = v2052
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2079 int32
	_ = v2079
	var v2084 int32
	_ = v2084
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2095 int32
	_ = v2095
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2118 int32
	_ = v2118
	var v2125 int32
	_ = v2125
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2137 int32
	_ = v2137
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2150 int32
	_ = v2150
	var v2154 int32
	_ = v2154
	var v2157 int64
	_ = v2157
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2222 int64
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2268 int32
	_ = v2268
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2285 int32
	_ = v2285
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2313 int32
	_ = v2313
	var v2317 int32
	_ = v2317
	var v2322 int32
	_ = v2322
	var v2331 int32
	_ = v2331
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2346 int32
	_ = v2346
	var v2351 int32
	_ = v2351
	var v2355 int32
	_ = v2355
	var v2359 int32
	_ = v2359
	var v2364 int32
	_ = v2364
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2398 int32
	_ = v2398
	var v2403 int32
	_ = v2403
	var v2408 int32
	_ = v2408
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2425 int32
	_ = v2425
	var v2428 int32
	_ = v2428
	var v2437 int32
	_ = v2437
	var v2444 int32
	_ = v2444
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2457 int32
	_ = v2457
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2467 int32
	_ = v2467
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2480 int32
	_ = v2480
	var v2491 int32
	_ = v2491
	var v2498 int32
	_ = v2498
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2511 int32
	_ = v2511
	var v2516 int32
	_ = v2516
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2531 int32
	_ = v2531
	var v2536 int32
	_ = v2536
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2544 int32
	_ = v2544
	var v2546 int32
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2551 int32
	_ = v2551
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2558 int64
	_ = v2558
	var v2561 int32
	_ = v2561
	var v2568 int32
	_ = v2568
	var v2571 int64
	_ = v2571
	var v2574 int64
	_ = v2574
	var v2575 int64
	_ = v2575
	var v2578 int64
	_ = v2578
	var v2579 int64
	_ = v2579
	var v2581 int64
	_ = v2581
	var v2582 int64
	_ = v2582
	var v2585 int64
	_ = v2585
	var v2591 int64
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2596 int32
	_ = v2596
	var v2599 int32
	_ = v2599
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2621 int32
	_ = v2621
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2631 int32
	_ = v2631
	var v2633 int32
	_ = v2633
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2649 int32
	_ = v2649
	var v2652 int64
	_ = v2652
	var v2655 int64
	_ = v2655
	var v2656 int64
	_ = v2656
	var v2659 int64
	_ = v2659
	var v2660 int64
	_ = v2660
	var v2662 int64
	_ = v2662
	var v2663 int64
	_ = v2663
	var v2666 int64
	_ = v2666
	var v2676 int32
	_ = v2676
	var v2687 int32
	_ = v2687
	var v2694 int32
	_ = v2694
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2707 int32
	_ = v2707
	var v2712 int32
	_ = v2712
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2726 int32
	_ = v2726
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2747 int32
	_ = v2747
	var v2750 int64
	_ = v2750
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2755 int64
	_ = v2755
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2766 int32
	_ = v2766
	var v2773 int32
	_ = v2773
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2785 int32
	_ = v2785
	var v2790 int32
	_ = v2790
	var v2791 int64
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2808 int32
	_ = v2808
	var v2814 int32
	_ = v2814
	var v2817 int32
	_ = v2817
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2827 int32
	_ = v2827
	var v2832 int32
	_ = v2832
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2845 int32
	_ = v2845
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2855 int32
	_ = v2855
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2876 int32
	_ = v2876
	var v2881 int32
	_ = v2881
	var v2884 int32
	_ = v2884
	var v2885 int32
	_ = v2885
	var v2886 int32
	_ = v2886
	var v2889 int32
	_ = v2889
	var v2891 int32
	_ = v2891
	var v2897 int32
	_ = v2897
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2904 int32
	_ = v2904
	var v2912 int32
	_ = v2912
	var v2916 int32
	_ = v2916
	var v2926 int32
	_ = v2926
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2945 int32
	_ = v2945
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2957 int64
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2965 int32
	_ = v2965
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2971 int32
	_ = v2971
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2979 int32
	_ = v2979
	var v2984 int32
	_ = v2984
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2989 int32
	_ = v2989
	var v2993 int32
	_ = v2993
	var v2996 int64
	_ = v2996
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3001 int64
	_ = v3001
	var v3008 int32
	_ = v3008
	var v3009 int32
	_ = v3009
	var v3012 int32
	_ = v3012
	var v3016 int64
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3021 int32
	_ = v3021
	var v3025 int32
	_ = v3025
	var v3028 int32
	_ = v3028
	var v3036 int32
	_ = v3036
	var v3037 int32
	_ = v3037
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3050 int32
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3054 int32
	_ = v3054
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3063 int32
	_ = v3063
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3068 int32
	_ = v3068
	var v3069 int32
	_ = v3069
	var v3075 int32
	_ = v3075
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3088 int32
	_ = v3088
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3101 int32
	_ = v3101
	var v3106 int32
	_ = v3106
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3118 int32
	_ = v3118
	var v3123 int32
	_ = v3123
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3142 int32
	_ = v3142
	var v3145 int64
	_ = v3145
	var v3151 float64
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3156 int32
	_ = v3156
	var v3162 int32
	_ = v3162
	var v3165 int32
	_ = v3165
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3177 int32
	_ = v3177
	var v3182 int32
	_ = v3182
	var v3183 float64
	_ = v3183
	var v3190 int32
	_ = v3190
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3200 int32
	_ = v3200
	var v3203 int64
	_ = v3203
	var v3209 float64
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3214 int32
	_ = v3214
	var v3220 int32
	_ = v3220
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3235 int32
	_ = v3235
	var v3240 int32
	_ = v3240
	var v3241 float64
	_ = v3241
	var v3248 int32
	_ = v3248
	var v3254 int32
	_ = v3254
	var v3260 int32
	_ = v3260
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3272 int32
	_ = v3272
	var v3277 int32
	_ = v3277
	var v3282 int32
	_ = v3282
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3291 int32
	_ = v3291
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3302 int32
	_ = v3302
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3314 int32
	_ = v3314
	var v3319 int32
	_ = v3319
	var v3323 int32
	_ = v3323
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3335 int32
	_ = v3335
	var v3340 int32
	_ = v3340
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3345 int32
	_ = v3345
	var v3346 int32
	_ = v3346
	var v3348 int32
	_ = v3348
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3362 int32
	_ = v3362
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3372 int32
	_ = v3372
	var v3375 int32
	_ = v3375
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3378 int32
	_ = v3378
	var v3384 int32
	_ = v3384
	var v3389 int32
	_ = v3389
	var v3391 int32
	_ = v3391
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3403 int32
	_ = v3403
	var v3404 int32
	_ = v3404
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3410 int32
	_ = v3410
	var v3411 int32
	_ = v3411
	var v3412 int32
	_ = v3412
	var v3413 int32
	_ = v3413
	var v3422 int32
	_ = v3422
	var v3427 int32
	_ = v3427
	var v3431 int32
	_ = v3431
	var v3436 int32
	_ = v3436
	var v3438 int32
	_ = v3438
	var v3442 int32
	_ = v3442
	var v3448 int32
	_ = v3448
	var v3451 int32
	_ = v3451
	var v3457 int32
	_ = v3457
	var v3461 int32
	_ = v3461
	var v3463 int32
	_ = v3463
	var v3471 int32
	_ = v3471
	var v3473 int32
	_ = v3473
	var v3475 int32
	_ = v3475
	var v3476 int32
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3478 int32
	_ = v3478
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3490 int32
	_ = v3490
	var v3491 int32
	_ = v3491
	var v3493 int32
	_ = v3493
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3501 int32
	_ = v3501
	var v3502 int32
	_ = v3502
	var v3503 int64
	_ = v3503
	var v3505 int32
	_ = v3505
	var v3506 int32
	_ = v3506
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3516 int32
	_ = v3516
	var v3517 int32
	_ = v3517
	var v3520 int32
	_ = v3520
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3528 int32
	_ = v3528
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3535 int32
	_ = v3535
	var v3537 int32
	_ = v3537
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3543 int32
	_ = v3543
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3547 int32
	_ = v3547
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3564 int32
	_ = v3564
	var v3565 int32
	_ = v3565
	var v3566 int32
	_ = v3566
	var v3568 int32
	_ = v3568
	var v3575 int32
	_ = v3575
	var v3578 int32
	_ = v3578
	var v3582 int32
	_ = v3582
	var v3587 int32
	_ = v3587
	var v3591 int32
	_ = v3591
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
	var v3604 int32
	_ = v3604
	var v3609 int32
	_ = v3609
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3626 int32
	_ = v3626
	var v3628 int32
	_ = v3628
	var v3634 int32
	_ = v3634
	var v3635 int32
	_ = v3635
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3640 int32
	_ = v3640
	var v3641 int32
	_ = v3641
	var v3642 int32
	_ = v3642
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3651 int32
	_ = v3651
	var v3654 int32
	_ = v3654
	var v3658 int32
	_ = v3658
	var v3663 int32
	_ = v3663
	var v3668 int32
	_ = v3668
	v6 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(992)
	m.G0 = v24
	F_check_stack_depth(m)
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
	v31 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v34 {
	case 0, 1, 2, 3, 28:
		goto L42
	case 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 30, 41, 42:
		goto L41
	case 14:
		goto L29
	case 15:
		goto L40
	case 16:
		goto L39
	case 17:
		goto L38
	case 18:
		goto L37
	case 19:
		goto L36
	case 20:
		goto L35
	case 21:
		goto L34
	case 22:
		goto L33
	case 23:
		goto L32
	case 24:
		goto L9
	case 25:
		goto L10
	case 26:
		goto L11
	case 27:
		goto L12
	case 29:
		goto L13
	case 31:
		goto L14
	case 32:
		goto L15
	case 33:
		goto L16
	case 34:
		goto L17
	case 35:
		goto L18
	case 36:
		goto L19
	case 37, 45, 50, 51, 52, 53:
		goto L20
	case 38:
		goto L21
	default:
		goto L28
	case 40:
		goto L22
	case 43:
		goto L23
	case 44:
		goto L24
	case 46, 48:
		goto L25
	case 47:
		goto L26
	case 49:
		goto L27
	}
L6:
	;
	goto L5
L7:
	;
	m.G0 = v24 + int32(992)
	return v3668
L8:
	;
	v3643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3643 != 0 {
		v3668 = v329
		goto L7
	} else {
		goto L1143
	}
L9:
	;
	v3612 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v3613 = m.ExcPending
	if v3613 != 0 {
		goto L1
	} else {
		goto L1132
	}
L10:
	;
	v3524 = F_JsonbType(m, l2)
	mBase = m.M
	v3525 = m.ExcPending
	if v3525 != 0 {
		goto L1
	} else {
		goto L1092
	}
L11:
	;
	v3520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3522 = F_executeNextItem(m, l0, l1, int32(0), v3520, l3, int32(1))
	mBase = m.M
	v3523 = m.ExcPending
	if v3523 != 0 {
		goto L1
	} else {
		goto L1090
	}
L12:
	;
	v3503 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v3505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3506 = *(*int32)(unsafe.Add(mBase, uint32(v3505)))
	if v3506 == int32(18) {
		goto L1086
	} else {
		goto L1087
	}
L13:
	;
	if l4 != 0 {
		goto L1076
	} else {
		goto L1077
	}
L14:
	;
	v3406 = F_palloc(m, int32(20))
	mBase = m.M
	v3407 = m.ExcPending
	if v3407 != 0 {
		goto L1
	} else {
		goto L1054
	}
L15:
	;
	v3350 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v3350 != int32(18) {
		goto L1038
	} else {
		goto L1039
	}
L16:
	;
	v3348 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1291), l3)
	mBase = m.M
	v3349 = m.ExcPending
	if v3349 != 0 {
		goto L1
	} else {
		goto L1036
	}
L17:
	;
	v3345 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1437), l3)
	mBase = m.M
	v3346 = m.ExcPending
	if v3346 != 0 {
		goto L1
	} else {
		goto L1035
	}
L18:
	;
	v3342 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1436), l3)
	mBase = m.M
	v3343 = m.ExcPending
	if v3343 != 0 {
		goto L1
	} else {
		goto L1034
	}
L19:
	;
	if l4 == int32(0) {
		goto L974
	} else {
		goto L975
	}
L20:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l4 == int32(0) {
		goto L668
	} else {
		goto L669
	}
L21:
	;
	if l4 == int32(0) {
		goto L582
	} else {
		goto L583
	}
L22:
	;
	v1667 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L1
	} else {
		goto L568
	}
L23:
	;
	if l4 == int32(0) {
		goto L525
	} else {
		goto L526
	}
L24:
	;
	if l4 == int32(0) {
		goto L478
	} else {
		goto L479
	}
L25:
	;
	if l4 == int32(0) {
		goto L354
	} else {
		goto L355
	}
L26:
	;
	if l4 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L27:
	;
	if l4 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L265
	}
L29:
	;
	v675 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1422), l3)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L264
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L259
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L256
	}
L32:
	;
	v325 = F_JsonbType(m, l2)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L160
	}
L33:
	;
	v271 = F_JsonbType(m, l2)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L135
	}
L34:
	;
	v227 = F_JsonbType(m, l2)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L114
	}
L35:
	;
	v225 = F_executeUnaryArithmExpr(m, l0, l1, l2, int32(1421), l3)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L113
	}
L36:
	;
	v222 = F_executeUnaryArithmExpr(m, l0, l1, l2, int32(0), l3)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L112
	}
L37:
	;
	v219 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1420), l3)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L111
	}
L38:
	;
	v216 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1419), l3)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L110
	}
L39:
	;
	v213 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1418), l3)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L109
	}
L40:
	;
	v210 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1417), l3)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L108
	}
L41:
	;
	v181 = F_executeBoolItem(m, l0, l1, l2, int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L98
	}
L42:
	;
	v37 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if l3 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v37 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	if v37 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v40 != int32(28) {
		v3668 = int32(0)
		goto L7
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v49 = F_palloc(m, int32(20))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v51 = v24 + int32(908)
	goto L50
L50:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v53 {
	case 0:
		goto L53
	case 1:
		goto L56
	case 2:
		goto L57
	case 3:
		goto L58
	default:
		goto L54
	case 28:
		goto L55
	}
L51:
	;
	v51 = v49
	goto L50
L52:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v126 {
		goto L77
	} else {
		goto L78
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
	goto L52
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L74
	}
L55:
	;
	v74 = v24 + int32(936)
	if v74 != 0 {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(1)
	v68 = v51 + int32(4)
	if v68 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(2)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v63
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(3)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)) = uint8(base.B2i32(v57 != int32(0)))
	goto L52
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v71
	goto L52
L60:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v69
	goto L62
L61:
	;
	goto L62
L62:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L59
L63:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v78 == int32(0) {
		goto L30
	} else {
		goto L67
	}
L64:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75
	goto L66
L65:
	;
	goto L66
L66:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L63
L67:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v24)+936))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = m.T0[v86].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v78, v77, v81, v24+int32(952), v24+int32(888))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v87 == int32(0) {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v24)+888))
	if v91 <= int32(0) {
		goto L52
	} else {
		goto L70
	}
L70:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v24)+960))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v24)+952))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v91
	if v101 == int32(18) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v106 = v100
	goto L73
L72:
	;
	v106 = int32(0)
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v106
	goto L52
L74:
	;
	F_errmsg_internal(m, int32(351255), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(476631), int32(2982), int32(276693))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v132 = F_executeItemOptUnwrapTarget(m, l0, v24+int32(752), v51, l3, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if l3 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v52
	v3668 = v132
	goto L7
L81:
	;
	if v37 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	goto L83
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v52
	v3668 = int32(0)
	goto L7
L84:
	;
	v148 = m.G0
	v150 = v148 - int32(16)
	m.G0 = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v152 != 0 {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v147 = v51
	goto L84
L86:
	;
	goto L87
L87:
	;
	v139 = F_palloc(m, int32(20))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+16)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v143
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v139))) = v145
	v147 = v139
	goto L84
L89:
	;
	m.G0 = v150 + int32(16)
	goto L83
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v150)+12)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v147
	v159 = F_list_make2_impl(m, v150+int32(4), v150)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v164 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v159
	goto L89
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v147
	goto L89
L95:
	;
	goto L96
L96:
	;
	v168 = F_lappend(m, v164, v147)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v168
	goto L89
L98:
	;
	v185 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if l3 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v189 = int32(0)
	if v185 == v189 {
		v3668 = v189
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v181 != int32(2) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)) = uint8(base.B2i32(v181 == int32(1)))
	v200 = int32(3)
	goto L106
L105:
	;
	v200 = int32(0)
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+952)) = v200
	v207 = F_executeNextItem(m, l0, l1, v24+int32(752), v24+int32(952), l3, int32(1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v3668 = v207
	goto L7
L108:
	;
	v3668 = v210
	goto L7
L109:
	;
	v3668 = v213
	goto L7
L110:
	;
	v3668 = v216
	goto L7
L111:
	;
	v3668 = v219
	goto L7
L112:
	;
	v3668 = v222
	goto L7
L113:
	;
	v3668 = v225
	goto L7
L114:
	;
	if v227 == int32(16) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v232 = v24 + int32(752)
	v236 = F_jspGetNext(m, l1, v232)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v242 = int32(1)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v243 == v242 {
		goto L123
	} else {
		goto L124
	}
L118:
	;
	if v236 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v238 = v232
	goto L121
L120:
	;
	v238 = int32(0)
	goto L121
L121:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v240 = F_executeItemUnwrapTargetArray(m, l0, v238, l2, l3, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v3668 = v240
	goto L7
L123:
	;
	v248 = F_executeNextItem(m, l0, l1, int32(0), l2, l3, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v250 != 0 {
		v3668 = v242
		goto L7
	} else {
		goto L127
	}
L126:
	;
	v3668 = v248
	goto L7
L127:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v251 != int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v3668 = int32(2)
	goto L7
L129:
	;
	goto L130
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(23498), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(476631), int32(849), int32(100760))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	if v271 == int32(17) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v277 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	if l4 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L139:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v279 != int32(18) {
		goto L31
	} else {
		goto L140
	}
L140:
	;
	if v277 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v285 = v24 + int32(752)
	goto L143
L142:
	;
	v285 = int32(0)
	goto L143
L143:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v287 = int32(1)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v292 = F_executeAnyItem(m, l0, v285, v286, l3, v287, v287, v287, int32(0), v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v3668 = v292
	goto L7
L145:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v303 != 0 {
		goto L150
	} else {
		goto L151
	}
L146:
	;
	v296 = F_JsonbType(m, l2)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v296 != int32(16) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v301 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v3668 = v301
	goto L7
L150:
	;
	v3668 = int32(1)
	goto L7
L151:
	;
	goto L152
L152:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v305 != int32(1) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v3668 = int32(2)
	goto L7
L154:
	;
	goto L155
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errcode(m, int32(319553666))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(103251), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(476631), int32(872), int32(100760))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
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
	if v325 != int32(16) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v329 = int32(1)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v330 != v329 {
		goto L8
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v334 != int32(18) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L163
L165:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v353 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L169
	}
L166:
	;
	v348 = int32(-1)
	goto L165
L167:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	if v338&int32(1342177280) != int32(1073741824) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v348 = v338 & int32(268435455)
	goto L165
L169:
	;
	v357 = base.B2i32(v348 < int32(0))
	if v348 < int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v358 = int32(1)
	goto L172
L171:
	;
	v358 = v348
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v358
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v360 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v374 = v6
	goto L176
L174:
	;
	v618 = int32(1)
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v349
	v3668 = v618
	goto L7
L176:
	;
	v389 = int32(2)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v396 = v374 << (uint(int32(3)) % 32)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v396+v397)))
	F_jspInitByBuffer(m, v24+int32(952), v394, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L178
	}
L177:
	;
	v618 = v593
	goto L175
L178:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v402+v396)+4))
	if v404 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v24+int32(908), v405, v404)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v414 = F_getArrayIndex(m, l0, v24+int32(952), l2, v24+int32(900))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L185
	}
L182:
	;
	goto L181
L183:
	;
	v611 = v374 + int32(1)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v611 < v612 {
		v374 = v611
		goto L176
	} else {
		goto L255
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v349
	v3668 = v577
	goto L7
L185:
	;
	if v414 == int32(2) {
		v577 = v389
		goto L184
	} else {
		goto L186
	}
L186:
	;
	if v404 != int32(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v430 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L188:
	;
	v422 = F_getArrayIndex(m, l0, v24+int32(908), l2, v24+int32(988))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v427
	v429 = v427
	goto L187
L191:
	;
	if v422 == int32(2) {
		v577 = v389
		goto L184
	} else {
		goto L192
	}
L192:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	v429 = v426
	goto L187
L193:
	;
	if v460 < v358 {
		goto L208
	} else {
		goto L209
	}
L194:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	v460 = v433
	goto L193
L195:
	;
	goto L196
L196:
	;
	if v429 < int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v440 != int32(1) {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	if v436 < v429 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	if v436 < v358 {
		v460 = v436
		goto L193
	} else {
		goto L200
	}
L200:
	;
	goto L197
L201:
	;
	v3668 = int32(2)
	goto L7
L202:
	;
	goto L203
L203:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errcode(m, int32(51118210))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errmsg(m, int32(162409), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(476631), int32(921), int32(100760))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	v462 = v460
	goto L210
L209:
	;
	v462 = v358 - int32(1)
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v462
	v464 = int32(0)
	if v464 < v429 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v467 = v429
	goto L213
L212:
	;
	v467 = v464
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = v467
	v469 = int32(1)
	if v462 < v467 {
		v593 = v469
		goto L183
	} else {
		goto L214
	}
L214:
	;
	v475 = v469
	v479 = v467
	v480 = v462
	goto L216
L215:
	;
	if l3 != 0 {
		v593 = v562
		goto L183
	} else {
		goto L253
	}
L216:
	;
	if v357 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	if v554 == int32(2) {
		v577 = v389
		goto L184
	} else {
		goto L252
	}
L218:
	;
	v558 = v479 + int32(1)
	if v558 <= v556 {
		v475 = v554
		v479 = v558
		v480 = v556
		goto L216
	} else {
		goto L251
	}
L219:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v495 = F_getIthJsonbValueFromContainer(m, v494, v479)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	v499 = l2
	goto L221
L221:
	;
	if base.B2i32(l3 != int32(0))|v353 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	if v495 == int32(0) {
		v554 = v475
		v556 = v480
		goto L218
	} else {
		goto L223
	}
L223:
	;
	v499 = v495
	goto L221
L224:
	;
	v3668 = int32(0)
	goto L7
L225:
	;
	goto L226
L226:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v503 <= int32(0) {
		goto L230
	} else {
		goto L231
	}
L227:
	;
	v554 = int32(0)
	v556 = v552
	goto L218
L228:
	;
	v546 = F_lappend(m, v535, v520)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L250
	}
L229:
	;
	if l3 != 0 {
		v554 = v544
		v556 = v480
		goto L218
	} else {
		goto L248
	}
L230:
	;
	if l3 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	goto L232
L232:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v540 = F_executeItemOptUnwrapTarget(m, l0, v24+int32(752), v499, l3, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L246
	}
L233:
	;
	v544 = int32(0)
	goto L229
L234:
	;
	goto L235
L235:
	;
	if int32(0) <= v348 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v521 != 0 {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	v520 = v499
	goto L236
L238:
	;
	goto L239
L239:
	;
	v512 = F_palloc(m, int32(20))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v499)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v512)+16)) = v514
	v516 = *(*int64)(unsafe.Add(mBase, uint32(v499)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v512)+8)) = v516
	v518 = *(*int64)(unsafe.Add(mBase, uint32(v499)))
	*(*int64)(unsafe.Add(mBase, uint32(v512))) = v518
	v520 = v512
	goto L236
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+888)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v24)+936)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v520
	v530 = F_list_make2_impl(m, v24+int32(44), v24+int32(40))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v535 != 0 {
		goto L228
	} else {
		goto L245
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v530
	v552 = v480
	goto L227
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v520
	v552 = v480
	goto L227
L246:
	;
	if v540 == int32(2) {
		v577 = v389
		goto L184
	} else {
		goto L247
	}
L247:
	;
	v544 = v540
	goto L229
L248:
	;
	if v544 != 0 {
		v554 = v544
		v556 = v480
		goto L218
	} else {
		goto L249
	}
L249:
	;
	v562 = int32(0)
	goto L215
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v546
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	v552 = v549
	goto L227
L251:
	;
	goto L217
L252:
	;
	v562 = v554
	goto L215
L253:
	;
	if v562 != 0 {
		v593 = v562
		goto L183
	} else {
		goto L254
	}
L254:
	;
	v577 = int32(0)
	goto L184
L255:
	;
	goto L177
L256:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v640
	F_errmsg_internal(m, int32(462308), v24+int32(32))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(476631), int32(858), int32(100760))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v24)+936))
	v661 = F_pnstrdup(m, v77, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v661
	F_errmsg(m, int32(679576), v24+int32(16))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(476631), int32(3158), int32(377811))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	v3668 = v675
	goto L7
L265:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v681
	F_errmsg_internal(m, int32(462701), v24)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(476631), int32(1663), int32(100760))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
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
	v700 = F_JsonbType(m, l2)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L280
	}
L269:
	;
	v693 = F_JsonbType(m, l2)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	if v693 != int32(16) {
		goto L268
	} else {
		goto L271
	}
L271:
	;
	v698 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	v3668 = v698
	goto L7
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v755
	if v755&int32(3) == int32(0) {
		v780 = v755
		goto L296
	} else {
		goto L297
	}
L274:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v753 = F_pnstrdup(m, v751, v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L293
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v726 != int32(1) {
		v3668 = int32(2)
		goto L7
	} else {
		goto L287
	}
L277:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v719 = F_JsonEncodeDateTime(m, v24+int32(752), v715, v716, l2+int32(16))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L285
	}
L278:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v711 != 0 {
		goto L282
	} else {
		goto L283
	}
L279:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v707 = F_DirectFunctionCall1Coll(m, int32(617), int32(0), v706)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L281
	}
L280:
	;
	switch v700 - int32(1) {
	case 0:
		goto L274
	case 1:
		goto L279
	case 2:
		goto L278
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
		goto L275
	default:
		goto L276
	case 31:
		goto L277
	}
L281:
	;
	v755 = v707
	goto L273
L282:
	;
	v712 = int32(327357)
	goto L284
L283:
	;
	v712 = int32(343932)
	goto L284
L284:
	;
	v755 = v712
	goto L273
L285:
	;
	v723 = F_pstrdup(m, v24+int32(752))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	v755 = v723
	goto L273
L287:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v737 = F_jspOperationName(m, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+736)) = v737
	F_errmsg(m, int32(329090), v24+int32(736))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(476631), int32(1648), int32(100760))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	goto L275
L293:
	;
	v755 = v753
	goto L273
L294:
	;
	v814 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+952)) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v24)+956)) = v813
	v821 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(952), l3, v814)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L311
	}
L295:
	;
	v813 = v805 - v755
	goto L294
L296:
	;
	v784 = v780
	goto L305
L297:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	if v764 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v813 = int32(0)
	goto L294
L299:
	;
	goto L300
L300:
	;
	v769 = v755
	goto L301
L301:
	;
	v773 = v769 + int32(1)
	if v773&int32(3) == int32(0) {
		v780 = v773
		goto L296
	} else {
		goto L303
	}
L302:
	;
	v805 = v773
	goto L295
L303:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773))))
	if v778 != 0 {
		v769 = v773
		goto L301
	} else {
		goto L304
	}
L304:
	;
	goto L302
L305:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v784)))
	v793 = int32(-2139062144)
	if (int32(16843008)-v790|v790)&v793 == v793 {
		v784 = v784 + int32(4)
		goto L305
	} else {
		goto L307
	}
L306:
	;
	v799 = v784
	goto L308
L307:
	;
	goto L306
L308:
	;
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799))))
	if v803 != 0 {
		v799 = v799 + int32(1)
		goto L308
	} else {
		goto L310
	}
L309:
	;
	v805 = v799
	goto L295
L310:
	;
	goto L309
L311:
	;
	v3668 = v821
	goto L7
L312:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v832 - int32(1) {
	case 0:
		goto L321
	case 1:
		goto L322
	default:
		goto L320
	}
L313:
	;
	v825 = F_JsonbType(m, l2)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	if v825 != int32(16) {
		goto L312
	} else {
		goto L315
	}
L315:
	;
	v830 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v3668 = v830
	goto L7
L317:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L349
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(2)
	v934 = F_DirectFunctionCall1Coll(m, int32(1413), int32(0), v929)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L346
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+908)) = v838
	v929 = v838
	goto L318
L320:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v903 != int32(1) {
		v3668 = int32(2)
		goto L7
	} else {
		goto L340
	}
L321:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v878 = F_pnstrdup(m, v876, v877)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L334
	}
L322:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v838 = F_numeric_int4_opt_error(m, v835, v24+int32(952))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+952)))
	if v840 != int32(1) {
		goto L319
	} else {
		goto L324
	}
L324:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v843 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v3668 = int32(2)
	goto L7
L326:
	;
	goto L327
L327:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v857 = F_DirectFunctionCall1Coll(m, int32(617), int32(0), v856)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v860 = F_jspOperationName(m, v859)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+712)) = int32(214368)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+708)) = v860
	*(*int32)(unsafe.Add(mBase, uint32(v24)+704)) = v857
	F_errmsg(m, int32(179756), v24+int32(704))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(476631), int32(1563), int32(100760))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	v881 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v881
	v884 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v884
	v892 = F_DirectInputFunctionCallSafe(m, int32(1423), v878, int32(-1), v24+int32(952), v24+int32(908))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L336
	}
L335:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v898 == int32(1) {
		goto L317
	} else {
		goto L339
	}
L336:
	;
	if v892 == int32(0) {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v896 != 0 {
		goto L335
	} else {
		goto L338
	}
L338:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v929 = v897
	goto L318
L339:
	;
	v3668 = int32(2)
	goto L7
L340:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v914 = F_jspOperationName(m, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+688)) = v914
	F_errmsg(m, int32(329215), v24+int32(688))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(476631), int32(1593), int32(100760))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L346:
	;
	v936 = F_pg_detoast_datum(m, v934)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+756)) = v936
	v943 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(752), l3, int32(1))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	v3668 = v943
	goto L7
L349:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v953 = F_jspOperationName(m, v952)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+728)) = int32(214368)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+724)) = v953
	*(*int32)(unsafe.Add(mBase, uint32(v24)+720)) = v878
	F_errmsg(m, int32(179756), v24+int32(720))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	F_errfinish(m, int32(476631), int32(1585), int32(100760))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L354:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v978 - int32(1) {
	case 0:
		goto L366
	case 1:
		goto L367
	default:
		goto L365
	}
L355:
	;
	v971 = F_JsonbType(m, l2)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	if v971 != int32(16) {
		goto L354
	} else {
		goto L357
	}
L357:
	;
	v976 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	v3668 = v976
	goto L7
L359:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L1
	} else {
		goto L473
	}
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L470
	}
L361:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L1
	} else {
		goto L467
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+956)) = v1314
	*(*int32)(unsafe.Add(mBase, uint32(v24)+952)) = int32(2)
	v1324 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(952), l3, int32(1))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L1
	} else {
		goto L466
	}
L363:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1142 != int32(46) {
		v1314 = v1140
		goto L362
	} else {
		goto L420
	}
L364:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L415
	}
L365:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1094 != int32(1) {
		goto L407
	} else {
		goto L408
	}
L366:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v1027
	v1030 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v1030
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1035 = F_pnstrdup(m, v1033, v1034)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L387
	}
L367:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v982 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v981)+4)))
	goto L369
L368:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1019 != int32(46) {
		v1314 = v981
		goto L362
	} else {
		goto L383
	}
L369:
	;
	if base.B2i32(v982 == int32(49152)) == int32(0) {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v987 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v981)+4)))
	goto L373
L371:
	;
	goto L372
L372:
	;
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v994 != int32(1) {
		goto L375
	} else {
		goto L376
	}
L373:
	;
	if base.B2i32(v987&int32(57343) == int32(53248)) == int32(0) {
		goto L368
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	v3668 = int32(2)
	goto L7
L376:
	;
	goto L377
L377:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1006 = F_jspOperationName(m, v1005)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+592)) = v1006
	F_errmsg(m, int32(641694), v24+int32(592))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	F_errfinish(m, int32(476631), int32(1414), int32(100760))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L383:
	;
	v1024 = F_DirectFunctionCall1Coll(m, int32(617), int32(0), v981)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	v1140 = v981
	v1141 = v1024
	goto L363
L385:
	;
	v3668 = int32(2)
	goto L7
L386:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v1075 = F_pg_detoast_datum(m, v1074)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L399
	}
L387:
	;
	v1042 = F_DirectInputFunctionCallSafe(m, int32(408), v1035, int32(-1), v24+int32(752), v24+int32(908))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	if v1042 != 0 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	if v1044 != int32(1) {
		goto L386
	} else {
		goto L392
	}
L390:
	;
	goto L391
L391:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1047 != int32(1) {
		goto L385
	} else {
		goto L393
	}
L392:
	;
	goto L391
L393:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1058 = F_jspOperationName(m, v1057)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+664)) = int32(468051)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+660)) = v1058
	*(*int32)(unsafe.Add(mBase, uint32(v24)+656)) = v1035
	F_errmsg(m, int32(179756), v24+int32(656))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(476631), int32(1439), int32(100760))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L399:
	;
	v1077 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)))
	goto L400
L400:
	;
	if base.B2i32(v1077 == int32(49152)) == int32(0) {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1082 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)))
	goto L404
L402:
	;
	goto L403
L403:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1089 == int32(1) {
		goto L364
	} else {
		goto L406
	}
L404:
	;
	if base.B2i32(v1082&int32(57343) == int32(53248)) == int32(0) {
		v1140 = v1075
		v1141 = v1035
		goto L363
	} else {
		goto L405
	}
L405:
	;
	goto L403
L406:
	;
	goto L385
L407:
	;
	v3668 = int32(2)
	goto L7
L408:
	;
	goto L409
L409:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1106 = F_jspOperationName(m, v1105)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+576)) = v1106
	F_errmsg(m, int32(329215), v24+int32(576))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	F_errfinish(m, int32(476631), int32(1455), int32(100760))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L415:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1127 = F_jspOperationName(m, v1126)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L417
	}
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+672)) = v1127
	F_errmsg(m, int32(641694), v24+int32(672))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L1
	} else {
		goto L418
	}
L418:
	;
	F_errfinish(m, int32(476631), int32(1446), int32(100760))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L419
	}
L419:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L420:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1145 == int32(0) {
		v1314 = v1140
		goto L362
	} else {
		goto L421
	}
L421:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+896)) = v1149
	v1152 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+888)) = v1152
	F_jspGetArg(m, l1, v24+int32(752))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L422
	}
L422:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v24)+752))
	if v1158 != int32(2) {
		goto L361
	} else {
		goto L423
	}
L423:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(752))+12))
	v1166 = F_numeric_int4_opt_error(m, v1163, v24+int32(984))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+984)))
	if v1168 == int32(1) {
		goto L427
	} else {
		goto L428
	}
L425:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	v1310 = F_pg_detoast_datum(m, v1309)
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L1
	} else {
		goto L464
	}
L426:
	;
	v3668 = int32(2)
	goto L7
L427:
	;
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1171 != int32(1) {
		goto L426
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1195 == int32(0) {
		v1239 = v6
		goto L436
	} else {
		goto L437
	}
L430:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L432
	}
L432:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1182 = F_jspOperationName(m, v1181)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+608)) = v1182
	F_errmsg(m, int32(214163), v24+int32(608))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(476631), int32(1487), int32(100760))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L436:
	;
	v1241 = v24 + int32(908)
	if int32(0) <= v1166 {
		goto L449
	} else {
		goto L450
	}
L437:
	;
	F_jspGetRightArg(m, l1, v24+int32(752))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v24)+752))
	if v1202 != int32(2) {
		goto L360
	} else {
		goto L439
	}
L439:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(752))+12))
	v1210 = F_numeric_int4_opt_error(m, v1207, v24+int32(984))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L440
	}
L440:
	;
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+984)))
	if v1212 != int32(1) {
		v1239 = v1210
		goto L436
	} else {
		goto L441
	}
L441:
	;
	v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1215 != int32(1) {
		goto L426
	} else {
		goto L442
	}
L442:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1226 = F_jspOperationName(m, v1225)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+640)) = v1226
	F_errmsg(m, int32(214236), v24+int32(640))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(476631), int32(1501), int32(100760))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = v24 + int32(908)
	v1263 = v24 + int32(936)
	if int32(0) <= v1239 {
		goto L453
	} else {
		goto L454
	}
L449:
	;
	v1251 = v1166
	v1252 = int32(0)
	goto L451
L450:
	;
	v1246 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1241))) = uint8(v1246)
	v1251 = int32(0) - v1166
	v1252 = int32(1)
	goto L451
L451:
	;
	v1254 = F_pg_ultoa_n(m, v1251, v1241+v1252)
	mBase = m.M
	v1257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1241+(v1254+v1252)))) = uint8(v1257)
	goto L448
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+904)) = v24 + int32(936)
	v1291 = F_construct_array_builtin(m, v24+int32(900), int32(2), int32(2275))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L456
	}
L453:
	;
	v1273 = v1239
	v1274 = int32(0)
	goto L455
L454:
	;
	v1268 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1263))) = uint8(v1268)
	v1273 = int32(0) - v1239
	v1274 = int32(1)
	goto L455
L455:
	;
	v1276 = F_pg_ultoa_n(m, v1273, v1263+v1274)
	mBase = m.M
	v1279 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1263+(v1276+v1274)))) = uint8(v1279)
	goto L452
L456:
	;
	v1293 = F_DirectFunctionCall1Coll(m, int32(1424), int32(0), v1291)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	v1299 = F_DirectInputFunctionCallSafe(m, int32(408), v1141, v1293, v24+int32(888), v24+int32(988))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	if v1299 != 0 {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+892)))
	if v1301 != int32(1) {
		goto L425
	} else {
		goto L462
	}
L460:
	;
	goto L461
L461:
	;
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1304 == int32(1) {
		goto L359
	} else {
		goto L463
	}
L462:
	;
	goto L461
L463:
	;
	goto L426
L464:
	;
	F_pfree(m, v1291)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	v1314 = v1310
	goto L362
L466:
	;
	v3668 = v1324
	goto L7
L467:
	;
	F_errmsg_internal(m, int32(259158), int32(0))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(476631), int32(1479), int32(100760))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(378477), int32(0))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	F_errfinish(m, int32(476631), int32(1493), int32(100760))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L1
	} else {
		goto L472
	}
L472:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L473:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L474
	}
L474:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1360 = F_jspOperationName(m, v1359)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L475
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+632)) = int32(468051)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+628)) = v1360
	*(*int32)(unsafe.Add(mBase, uint32(v24)+624)) = v1141
	F_errmsg(m, int32(179756), v24+int32(624))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	F_errfinish(m, int32(476631), int32(1528), int32(100760))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L1
	} else {
		goto L477
	}
L477:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L478:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1385 - int32(1) {
	case 0:
		goto L488
	case 1:
		goto L489
	case 2:
		goto L485
	default:
		goto L487
	}
L479:
	;
	v1378 = F_JsonbType(m, l2)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	if v1378 != int32(16) {
		goto L478
	} else {
		goto L481
	}
L481:
	;
	v1383 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	v3668 = v1383
	goto L7
L483:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)) = uint8(v1507)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(3)
	v1515 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(752), l3, int32(1))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L1
	} else {
		goto L524
	}
L484:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+936)) = uint8(v1505)
	v1507 = v1505
	goto L483
L485:
	;
	v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v1505 = v1503
	goto L484
L486:
	;
	v3668 = int32(2)
	goto L7
L487:
	;
	v1477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1477 != int32(1) {
		goto L516
	} else {
		goto L517
	}
L488:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1442 = F_pnstrdup(m, v1440, v1441)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L1
	} else {
		goto L503
	}
L489:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1391 = F_DirectFunctionCall1Coll(m, int32(617), int32(0), v1390)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v1394
	v1397 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v1397
	v1405 = F_DirectInputFunctionCallSafe(m, int32(1423), v1391, int32(-1), v24+int32(952), v24+int32(908))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L492
	}
L491:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v1505 = base.B2i32(v1437 != int32(0))
	goto L484
L492:
	;
	if v1405 != 0 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v1407 != int32(1) {
		goto L491
	} else {
		goto L496
	}
L494:
	;
	goto L495
L495:
	;
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1410 != int32(1) {
		goto L486
	} else {
		goto L497
	}
L496:
	;
	goto L495
L497:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1421 = F_jspOperationName(m, v1420)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+552)) = int32(270013)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+548)) = v1421
	*(*int32)(unsafe.Add(mBase, uint32(v24)+544)) = v1391
	F_errmsg(m, int32(179756), v24+int32(544))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(476631), int32(1357), int32(100760))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L1
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
	v1446 = F_strlen(m, v1442)
	mBase = m.M
	v1447 = F_parse_bool_with_len(m, v1442, v1446, v24+int32(936))
	mBase = m.M
	goto L504
L504:
	;
	if v1447 != 0 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v1448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+936)))
	v1507 = v1448
	goto L483
L506:
	;
	goto L507
L507:
	;
	v1449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1449 != int32(1) {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3668 = int32(2)
	goto L7
L509:
	;
	goto L510
L510:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1461 = F_jspOperationName(m, v1460)
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+568)) = int32(270013)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+564)) = v1461
	*(*int32)(unsafe.Add(mBase, uint32(v24)+560)) = v1442
	F_errmsg(m, int32(179756), v24+int32(560))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(476631), int32(1377), int32(100760))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L516:
	;
	v3668 = int32(2)
	goto L7
L517:
	;
	goto L518
L518:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1489 = F_jspOperationName(m, v1488)
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+528)) = v1489
	F_errmsg(m, int32(329291), v24+int32(528))
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	F_errfinish(m, int32(476631), int32(1386), int32(100760))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L524:
	;
	v3668 = v1515
	goto L7
L525:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1526 - int32(1) {
	case 0:
		goto L534
	case 1:
		goto L535
	default:
		goto L533
	}
L526:
	;
	v1519 = F_JsonbType(m, l2)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	if v1519 != int32(16) {
		goto L525
	} else {
		goto L528
	}
L528:
	;
	v1524 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	v3668 = v1524
	goto L7
L530:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L1
	} else {
		goto L563
	}
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(2)
	v1630 = F_DirectFunctionCall1Coll(m, int32(1414), int32(0), v1624)
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L1
	} else {
		goto L560
	}
L532:
	;
	v1621 = F_Int64GetDatum(m, v1532)
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L1
	} else {
		goto L559
	}
L533:
	;
	v1597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1597 != int32(1) {
		v3668 = int32(2)
		goto L7
	} else {
		goto L553
	}
L534:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1572 = F_pnstrdup(m, v1570, v1571)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L1
	} else {
		goto L547
	}
L535:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1532 = F_numeric_int8_opt_error(m, v1529, v24+int32(952))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	v1534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+952)))
	if v1534 != int32(1) {
		goto L532
	} else {
		goto L537
	}
L537:
	;
	v1537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1537 == int32(0) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v3668 = int32(2)
	goto L7
L539:
	;
	goto L540
L540:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1551 = F_DirectFunctionCall1Coll(m, int32(617), int32(0), v1550)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1554 = F_jspOperationName(m, v1553)
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+504)) = int32(84881)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+500)) = v1554
	*(*int32)(unsafe.Add(mBase, uint32(v24)+496)) = v1551
	F_errmsg(m, int32(179756), v24+int32(496))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L1
	} else {
		goto L545
	}
L545:
	;
	F_errfinish(m, int32(476631), int32(1283), int32(100760))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L547:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v1575
	v1578 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v1578
	v1586 = F_DirectInputFunctionCallSafe(m, int32(546), v1572, int32(-1), v24+int32(952), v24+int32(908))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L1
	} else {
		goto L549
	}
L548:
	;
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1592 == int32(1) {
		goto L530
	} else {
		goto L552
	}
L549:
	;
	if v1586 == int32(0) {
		goto L548
	} else {
		goto L550
	}
L550:
	;
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v1590 != 0 {
		goto L548
	} else {
		goto L551
	}
L551:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v1624 = v1591
	goto L531
L552:
	;
	v3668 = int32(2)
	goto L7
L553:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L1
	} else {
		goto L554
	}
L554:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L1
	} else {
		goto L555
	}
L555:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1608 = F_jspOperationName(m, v1607)
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+480)) = v1608
	F_errmsg(m, int32(329215), v24+int32(480))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	F_errfinish(m, int32(476631), int32(1313), int32(100760))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L1
	} else {
		goto L558
	}
L558:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L559:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+908)) = v1621
	v1624 = v1621
	goto L531
L560:
	;
	v1632 = F_pg_detoast_datum(m, v1630)
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+756)) = v1632
	v1639 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(752), l3, int32(1))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L1
	} else {
		goto L562
	}
L562:
	;
	v3668 = v1639
	goto L7
L563:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1649 = F_jspOperationName(m, v1648)
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+520)) = int32(84881)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+516)) = v1649
	*(*int32)(unsafe.Add(mBase, uint32(v24)+512)) = v1572
	F_errmsg(m, int32(179756), v24+int32(512))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	F_errfinish(m, int32(476631), int32(1305), int32(100760))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L568:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(0) <= v1669 {
		goto L569
	} else {
		goto L570
	}
L569:
	;
	v1672 = int32(0)
	if base.B2i32(l3 == v1672)&(v1667^int32(1)) != 0 {
		v3668 = v1672
		goto L7
	} else {
		goto L572
	}
L570:
	;
	goto L571
L571:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L1
	} else {
		goto L579
	}
L572:
	;
	if v1667 == int32(0) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v1685 = F_palloc(m, int32(20))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L576
	}
L574:
	;
	v1687 = v24 + int32(952)
	goto L575
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1687))) = int32(2)
	v1691 = F_int64_to_numeric(m, base.I64_extend_i32_s(v1669-int32(1)))
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L1
	} else {
		goto L577
	}
L576:
	;
	v1687 = v1685
	goto L575
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1687)+4)) = v1691
	v1696 = F_executeNextItem(m, l0, l1, v24+int32(752), v1687, l3, v1667)
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	v3668 = v1696
	goto L7
L579:
	;
	F_errmsg_internal(m, int32(79587), int32(0))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	F_errfinish(m, int32(476631), int32(1241), int32(100760))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L582:
	;
	v1720 = m.G0
	v1722 = v1720 - int32(224)
	m.G0 = v1722
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v1724 == int32(18) {
		goto L591
	} else {
		goto L592
	}
L583:
	;
	v1713 = F_JsonbType(m, l2)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L1
	} else {
		goto L584
	}
L584:
	;
	if v1713 != int32(16) {
		goto L582
	} else {
		goto L585
	}
L585:
	;
	v1718 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	v3668 = v1718
	goto L7
L587:
	;
	v3668 = v2022
	goto L7
L588:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L1
	} else {
		goto L663
	}
L589:
	;
	m.G0 = v1722 + int32(224)
	goto L587
L590:
	;
	if v1728&int32(268435455) == int32(0) {
		goto L604
	} else {
		goto L605
	}
L591:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1727)))
	if v1728&int32(536870912) != 0 {
		goto L590
	} else {
		goto L594
	}
L592:
	;
	goto L593
L593:
	;
	v1737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1737 != int32(1) {
		goto L596
	} else {
		goto L597
	}
L594:
	;
	if v1728&int32(1073741824) == int32(0) {
		goto L588
	} else {
		goto L595
	}
L595:
	;
	goto L593
L596:
	;
	v2022 = int32(2)
	goto L589
L597:
	;
	goto L598
L598:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	F_errcode(m, int32(319553666))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1749 = F_jspOperationName(m, v1748)
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+16)) = v1749
	F_errmsg(m, int32(103318), v1722+int32(16))
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	F_errfinish(m, int32(476631), int32(2840), int32(402770))
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L1
	} else {
		goto L603
	}
L603:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L604:
	;
	v2022 = int32(1)
	goto L589
L605:
	;
	goto L606
L606:
	;
	v1769 = F_jspGetNext(m, l1, v1722+int32(188))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+116)) = int32(20889)
	*(*int64)(unsafe.Add(mBase, uint32(v1722)+108)) = int64(12884901889)
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+96)) = int32(330269)
	*(*int64)(unsafe.Add(mBase, uint32(v1722)+88)) = int64(21474836481)
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+76)) = int32(416708)
	*(*int64)(unsafe.Add(mBase, uint32(v1722)+68)) = int64(8589934593)
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v1783 == int32(18) {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1789 = base.I64_extend_i32_s(v1727 - v1786)
	goto L610
L609:
	;
	v1789 = int64(0)
	goto L610
L610:
	;
	v1790 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+128)) = int32(2)
	v1796 = F_int64_to_numeric(m, v1790*int64(10000000000)+v1789)
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L1
	} else {
		goto L611
	}
L611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+132)) = v1796
	v1799 = F_JsonbIteratorInit(m, v1727)
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+64)) = v1799
	v1802 = int32(1)
	v1808 = F_JsonbIteratorNext(m, v1722-int32(-64), v1722+int32(168), v1802)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	if v1808 == int32(0) {
		v2022 = v1802
		goto L589
	} else {
		goto L614
	}
L614:
	;
	v1817 = v1802
	v1819 = v1808
	goto L615
L615:
	;
	if v1819 != int32(1) {
		v2005 = v1817
		goto L618
	} else {
		goto L619
	}
L616:
	;
	v2022 = v2017
	goto L589
L617:
	;
	goto L616
L618:
	;
	v2014 = F_JsonbIteratorNext(m, v1722-int32(-64), v1722+int32(168), int32(1))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L661
	}
L619:
	;
	v1838 = int32(0)
	if base.B2i32(l3 != int32(0))|v1769 == v1838 {
		v2017 = v1838
		goto L617
	} else {
		goto L620
	}
L620:
	;
	v1846 = F_JsonbIteratorNext(m, v1722-int32(-64), v1722+int32(148), int32(1))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L1
	} else {
		goto L621
	}
L621:
	;
	v1848 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+40)) = v1848
	v1854 = F_pushJsonbValue(m, v1722+int32(40), int32(6), v1848)
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	v1861 = F_pushJsonbValue(m, v1722+int32(40), int32(1), v1722+int32(108))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L623
	}
L623:
	;
	v1868 = F_pushJsonbValue(m, v1722+int32(40), int32(2), v1722+int32(168))
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	v1875 = F_pushJsonbValue(m, v1722+int32(40), int32(1), v1722+int32(88))
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	v1882 = F_pushJsonbValue(m, v1722+int32(40), int32(2), v1722+int32(148))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	v1889 = F_pushJsonbValue(m, v1722+int32(40), int32(1), v1722+int32(68))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	v1896 = F_pushJsonbValue(m, v1722+int32(40), int32(2), v1722+int32(128))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	v1902 = F_pushJsonbValue(m, v1722+int32(40), int32(7), int32(0))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L1
	} else {
		goto L629
	}
L629:
	;
	v1904 = F_JsonbValueToJsonb(m, v1902)
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L1
	} else {
		goto L630
	}
L630:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+44)) = int32(18)
	v1909 = v1904 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+52)) = v1909
	v1911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1904))))
	if v1911 == int32(1) {
		goto L632
	} else {
		goto L633
	}
L631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+48)) = v1941
	v1943 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1909
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1945
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1945 + int32(1)
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1950 <= int32(0) {
		goto L643
	} else {
		goto L644
	}
L632:
	;
	v1914 = int32(4)
	v1916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1904)+1)))
	if v1916&int32(254) == int32(2) {
		goto L635
	} else {
		goto L636
	}
L633:
	;
	goto L634
L634:
	;
	v1929 = int32(1)
	if v1911&v1929 != 0 {
		v1941 = int32(base.Ui32(v1911)>>(uint(v1929)%32)) - v1929
		goto L631
	} else {
		goto L641
	}
L635:
	;
	v1925 = v1914
	goto L637
L636:
	;
	v1925 = base.B2i32(v1916 == int32(18)) << (uint(v1914) % 32)
	goto L637
L637:
	;
	if v1916 == int32(1) {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v1928 = v1914
	goto L640
L639:
	;
	v1928 = v1925
	goto L640
L640:
	;
	v1941 = v1928
	goto L631
L641:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1904)))
	v1941 = int32(base.Ui32(v1935)>>(uint(int32(2))%32)) - int32(4)
	goto L631
L642:
	;
	if l3 != 0 {
		v2005 = v2000
		goto L618
	} else {
		goto L659
	}
L643:
	;
	if l3 == int32(0) {
		goto L646
	} else {
		goto L647
	}
L644:
	;
	goto L645
L645:
	;
	v1993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v1994 = F_executeItemOptUnwrapTarget(m, l0, v1722+int32(188), v1722+int32(44), l3, v1993)
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L1
	} else {
		goto L657
	}
L646:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v1943
	v2000 = int32(0)
	goto L642
L647:
	;
	v1956 = F_palloc(m, int32(20))
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L1
	} else {
		goto L648
	}
L648:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1722)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1956)+16)) = v1958
	v1960 = *(*int64)(unsafe.Add(mBase, uint32(v1722)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v1956)+8)) = v1960
	v1962 = *(*int64)(unsafe.Add(mBase, uint32(v1722)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v1956))) = v1962
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1964 != 0 {
		goto L649
	} else {
		goto L650
	}
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+216)) = v1956
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+220)) = v1964
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+36)) = v1964
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+32)) = v1956
	v1973 = F_list_make2_impl(m, v1722+int32(36), v1722+int32(32))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L652
	}
L650:
	;
	goto L651
L651:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v1978 == int32(0) {
		goto L653
	} else {
		goto L654
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v1973
	goto L646
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1956
	goto L646
L654:
	;
	goto L655
L655:
	;
	v1982 = F_lappend(m, v1978, v1956)
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L1
	} else {
		goto L656
	}
L656:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v1982
	goto L646
L657:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v1943
	v1997 = int32(2)
	if v1994 == v1997 {
		v2017 = v1997
		goto L617
	} else {
		goto L658
	}
L658:
	;
	v2000 = v1994
	goto L642
L659:
	;
	v2002 = int32(0)
	if v2000 == v2002 {
		v2017 = v2002
		goto L617
	} else {
		goto L660
	}
L660:
	;
	v2005 = v2000
	goto L618
L661:
	;
	if v2014 != 0 {
		v1817 = v2005
		v1819 = v2014
		goto L615
	} else {
		goto L662
	}
L662:
	;
	v2022 = v2005
	goto L589
L663:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v1727)))
	*(*int32)(unsafe.Add(mBase, uint32(v1722))) = v2048
	F_errmsg_internal(m, int32(27340), v1722)
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	F_errfinish(m, int32(476631), int32(3629), int32(353884))
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L1
	} else {
		goto L665
	}
L665:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L666:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2145 = F_cstring_to_text_with_len(m, v2143, v2144)
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L1
	} else {
		goto L691
	}
L667:
	;
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2118 != int32(1) {
		goto L683
	} else {
		goto L684
	}
L668:
	;
	v2108 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2108
	*(*int32)(unsafe.Add(mBase, uint32(v24)+984)) = int32(0)
	if v2058 == int32(1) {
		goto L666
	} else {
		goto L682
	}
L669:
	;
	switch v2058 - int32(16) {
	case 0:
		goto L671
	default:
		goto L668
	case 2:
		goto L672
	}
L670:
	;
	v2101 = int32(1)
	v2104 = int32(0)
	v2106 = F_executeAnyItem(m, l0, l1, v2063, l3, v2101, v2101, v2101, v2104, v2104)
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L1
	} else {
		goto L681
	}
L671:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L1
	} else {
		goto L678
	}
L672:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v2063)))
	if v2064&int32(536870912) != 0 {
		goto L667
	} else {
		goto L673
	}
L673:
	;
	if v2064&int32(1073741824) != 0 {
		goto L670
	} else {
		goto L674
	}
L674:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L1
	} else {
		goto L675
	}
L675:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2063)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+448)) = v2073
	F_errmsg_internal(m, int32(27340), v24+int32(448))
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L1
	} else {
		goto L676
	}
L676:
	;
	F_errfinish(m, int32(476631), int32(3629), int32(353884))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L1
	} else {
		goto L677
	}
L677:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L678:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+464)) = v2089
	F_errmsg_internal(m, int32(462926), v24+int32(464))
	mBase = m.M
	v2095 = m.ExcPending
	if v2095 != 0 {
		goto L1
	} else {
		goto L679
	}
L679:
	;
	F_errfinish(m, int32(476631), int32(1680), int32(24373))
	mBase = m.M
	v2100 = m.ExcPending
	if v2100 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L681:
	;
	v3668 = v2106
	goto L7
L682:
	;
	goto L667
L683:
	;
	v3668 = int32(2)
	goto L7
L684:
	;
	goto L685
L685:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		goto L1
	} else {
		goto L687
	}
L687:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2130 = F_jspOperationName(m, v2129)
	mBase = m.M
	v2131 = m.ExcPending
	if v2131 != 0 {
		goto L1
	} else {
		goto L688
	}
L688:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+432)) = v2130
	F_errmsg(m, int32(313978), v24+int32(432))
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	F_errfinish(m, int32(476631), int32(2357), int32(402792))
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L691:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v2147 - int32(37) {
	case 0:
		goto L698
	default:
		goto L697
	case 8:
		v2216 = v2108
		goto L696
	}
L692:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v2425 - int32(37) {
	case 0:
		v3025 = v2408
		goto L762
	default:
		goto L770
	case 8:
		goto L775
	case 13:
		goto L774
	case 14:
		goto L773
	case 15:
		goto L772
	case 16:
		goto L771
	}
L693:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L1
	} else {
		goto L754
	}
L694:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L1
	} else {
		goto L750
	}
L695:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L1
	} else {
		goto L747
	}
L696:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	v2222 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	v2226 = int32(0)
	goto L717
L697:
	;
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2192 == int32(0) {
		v2216 = v2108
		goto L696
	} else {
		goto L711
	}
L698:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2150 == int32(0) {
		v2216 = v2108
		goto L696
	} else {
		goto L699
	}
L699:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v2154
	v2157 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2157
	F_jspGetArg(m, l1, v24+int32(952))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L1
	} else {
		goto L700
	}
L700:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v24)+952))
	if v2163 != int32(1) {
		goto L695
	} else {
		goto L701
	}
L701:
	;
	v2167 = v24 + int32(952)
	v2169 = v24 + int32(936)
	if v2169 != 0 {
		goto L703
	} else {
		goto L704
	}
L702:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v24)+936))
	v2174 = F_cstring_to_text_with_len(m, v2172, v2173)
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L1
	} else {
		goto L706
	}
L703:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v2167)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2169))) = v2170
	goto L705
L704:
	;
	goto L705
L705:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2167)+12))
	goto L702
L706:
	;
	v2185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2185 != 0 {
		goto L707
	} else {
		goto L708
	}
L707:
	;
	v2186 = int32(0)
	goto L709
L708:
	;
	v2186 = v24 + int32(752)
	goto L709
L709:
	;
	v2187 = F_parse_datetime(m, v2145, v2174, v24+int32(900), v24+int32(988), v24+int32(984), v2186)
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	v2189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	v2408 = v2187
	v2410 = v2108
	v2412 = v2189 ^ int32(1)
	goto L692
L711:
	;
	F_jspGetArg(m, l1, v24+int32(952))
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L1
	} else {
		goto L712
	}
L712:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v24)+952))
	if v2199 != int32(2) {
		goto L694
	} else {
		goto L713
	}
L713:
	;
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(952))+12))
	v2207 = F_numeric_int4_opt_error(m, v2204, v24+int32(752))
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L1
	} else {
		goto L714
	}
L714:
	;
	v2209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+752)))
	if v2209 != int32(1) {
		v2216 = v2207
		goto L696
	} else {
		goto L715
	}
L715:
	;
	v2212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2212 == int32(1) {
		goto L693
	} else {
		goto L716
	}
L716:
	;
	v3668 = int32(2)
	goto L7
L717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(760)))) = v2218
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2222
	v2248 = v2226 << (uint(int32(2)) % 32)
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v2248)+uint32(_consts[1078])))
	if v2251 == int32(0) {
		goto L719
	} else {
		goto L720
	}
L718:
	;
	v2288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2289 == int32(37) {
		goto L726
	} else {
		goto L727
	}
L719:
	;
	v2254 = int32(4442576)
	v2255 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2258 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2258
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v2248)+uint32(_consts[1079])))
	v2263 = F_cstring_to_text(m, v2262)
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L1
	} else {
		goto L722
	}
L720:
	;
	v2268 = v2251
	goto L721
L721:
	;
	v2279 = F_parse_datetime(m, v2145, v2268, v24+int32(900), v24+int32(988), v24+int32(984), v24+int32(752))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L1
	} else {
		goto L723
	}
L722:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2248)+uint32(_consts[1078]))) = v2263
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2255
	v2268 = v2263
	goto L721
L723:
	;
	v2281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	if v2281 == int32(0) {
		v2408 = v2279
		v2410 = v2216
		v2412 = int32(1)
		goto L692
	} else {
		goto L724
	}
L724:
	;
	v2285 = v2226 + int32(1)
	if v2285 != int32(13) {
		v2226 = v2285
		goto L717
	} else {
		goto L725
	}
L725:
	;
	goto L718
L726:
	;
	if v2288&int32(1) == int32(0) {
		goto L729
	} else {
		goto L730
	}
L727:
	;
	goto L728
L728:
	;
	if v2288&int32(1) == int32(0) {
		goto L738
	} else {
		goto L739
	}
L729:
	;
	v3668 = int32(2)
	goto L7
L730:
	;
	goto L731
L731:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L1
	} else {
		goto L732
	}
L732:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L1
	} else {
		goto L733
	}
L733:
	;
	v2304 = F_text_to_cstring(m, v2145)
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+164)) = v2304
	*(*int32)(unsafe.Add(mBase, uint32(v24)+160)) = int32(356581)
	F_errmsg(m, int32(685515), v24+int32(160))
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	F_errhint(m, int32(542642), int32(0))
	mBase = m.M
	v2317 = m.ExcPending
	if v2317 != 0 {
		goto L1
	} else {
		goto L736
	}
L736:
	;
	F_errfinish(m, int32(476631), int32(2485), int32(402792))
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L1
	} else {
		goto L737
	}
L737:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L738:
	;
	v3668 = int32(2)
	goto L7
L739:
	;
	goto L740
L740:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L1
	} else {
		goto L741
	}
L741:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2336 = F_jspOperationName(m, v2335)
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L1
	} else {
		goto L743
	}
L743:
	;
	v2338 = F_text_to_cstring(m, v2145)
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L1
	} else {
		goto L744
	}
L744:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+180)) = v2338
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = v2336
	F_errmsg(m, int32(685515), v24+int32(176))
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L1
	} else {
		goto L745
	}
L745:
	;
	F_errfinish(m, int32(476631), int32(2490), int32(402792))
	mBase = m.M
	v2351 = m.ExcPending
	if v2351 != 0 {
		goto L1
	} else {
		goto L746
	}
L746:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L747:
	;
	F_errmsg_internal(m, int32(88636), int32(0))
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	F_errfinish(m, int32(476631), int32(2383), int32(402792))
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L750:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2370 = F_jspOperationName(m, v2369)
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L1
	} else {
		goto L751
	}
L751:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+416)) = v2370
	F_errmsg_internal(m, int32(88048), v24+int32(416))
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L1
	} else {
		goto L752
	}
L752:
	;
	F_errfinish(m, int32(476631), int32(2442), int32(402792))
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L1
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
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L1
	} else {
		goto L755
	}
L755:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2391 = F_jspOperationName(m, v2390)
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L1
	} else {
		goto L756
	}
L756:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+400)) = v2391
	F_errmsg(m, int32(214158), v24+int32(400))
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L1
	} else {
		goto L757
	}
L757:
	;
	F_errfinish(m, int32(476631), int32(2450), int32(402792))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L1
	} else {
		goto L758
	}
L758:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L759:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3111 = m.ExcPending
	if v3111 != 0 {
		goto L1
	} else {
		goto L971
	}
L760:
	;
	v3081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3081 != int32(1) {
		goto L963
	} else {
		goto L964
	}
L761:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3063 = m.ExcPending
	if v3063 != 0 {
		goto L1
	} else {
		goto L958
	}
L762:
	;
	F_pfree(m, v2145)
	mBase = m.M
	v3028 = m.ExcPending
	if v3028 != 0 {
		goto L1
	} else {
		goto L945
	}
L763:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1184)
	v3025 = v3021
	goto L762
L764:
	;
	v3016 = *(*int64)(unsafe.Add(mBase, uint32(v24)+888))
	v3017 = F_Int64GetDatum(m, v3016)
	mBase = m.M
	v3018 = m.ExcPending
	if v3018 != 0 {
		goto L1
	} else {
		goto L944
	}
L765:
	;
	v3668 = int32(2)
	goto L7
L766:
	;
	if v2410 == int32(-1) {
		v3021 = v2989
		goto L763
	} else {
		goto L939
	}
L767:
	;
	v2986 = F_DirectFunctionCall1Coll(m, v2984, int32(0), v2408)
	mBase = m.M
	v2987 = m.ExcPending
	if v2987 != 0 {
		goto L1
	} else {
		goto L938
	}
L768:
	;
	v2971 = *(*int32)(unsafe.Add(mBase, _consts[330]))
	v2973 = m.G0
	v2974 = int32(16)
	v2975 = v2973 - v2974
	m.G0 = v2975
	v2979 = F_DetermineTimeZoneOffsetInternal(m, v24+int32(752), v2971, v2975+int32(8))
	mBase = m.M
	m.G0 = v2975 + v2974
	goto L937
L769:
	;
	v2951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2951, int32(224999), int32(7437))
	mBase = m.M
	v2955 = m.ExcPending
	if v2955 != 0 {
		goto L1
	} else {
		goto L934
	}
L770:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L1
	} else {
		goto L931
	}
L771:
	;
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2799 <= int32(1183) {
		goto L910
	} else {
		goto L911
	}
L772:
	;
	v2676 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2676 <= int32(1183) {
		goto L870
	} else {
		goto L871
	}
L773:
	;
	v2599 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2599 <= int32(1183) {
		goto L841
	} else {
		goto L842
	}
L774:
	;
	v2480 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2480 <= int32(1183) {
		goto L804
	} else {
		goto L805
	}
L775:
	;
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2428 <= int32(1183) {
		goto L781
	} else {
		goto L782
	}
L776:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1082)
	v3025 = v2476
	goto L762
L777:
	;
	v2474 = F_DirectFunctionCall1Coll(m, v2472, int32(0), v2408)
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L1
	} else {
		goto L796
	}
L778:
	;
	if v2428 != int32(1114) {
		goto L759
	} else {
		goto L795
	}
L779:
	;
	v2463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2463, int32(7437), int32(339894))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L1
	} else {
		goto L794
	}
L780:
	;
	v2437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2437 != int32(1) {
		goto L786
	} else {
		goto L787
	}
L781:
	;
	switch v2428 - int32(1082) {
	case 0:
		v2476 = v2408
		goto L776
	case 1:
		goto L780
	default:
		goto L778
	}
L782:
	;
	goto L783
L783:
	;
	if v2428 == int32(1184) {
		goto L779
	} else {
		goto L784
	}
L784:
	;
	if v2428 != int32(1266) {
		goto L759
	} else {
		goto L785
	}
L785:
	;
	goto L780
L786:
	;
	v3668 = int32(2)
	goto L7
L787:
	;
	goto L788
L788:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2444 = m.ExcPending
	if v2444 != 0 {
		goto L1
	} else {
		goto L789
	}
L789:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L1
	} else {
		goto L790
	}
L790:
	;
	v2448 = F_text_to_cstring(m, v2145)
	mBase = m.M
	v2449 = m.ExcPending
	if v2449 != 0 {
		goto L1
	} else {
		goto L791
	}
L791:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+228)) = v2448
	*(*int32)(unsafe.Add(mBase, uint32(v24)+224)) = int32(339894)
	F_errmsg(m, int32(685515), v24+int32(224))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L1
	} else {
		goto L792
	}
L792:
	;
	F_errfinish(m, int32(476631), int32(2517), int32(402792))
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L1
	} else {
		goto L793
	}
L793:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L794:
	;
	v2472 = int32(1425)
	goto L777
L795:
	;
	v2472 = int32(1426)
	goto L777
L796:
	;
	v2476 = v2474
	goto L776
L797:
	;
	if v2410 != int32(-1) {
		goto L823
	} else {
		goto L824
	}
L798:
	;
	v2548 = F_DirectFunctionCall1Coll(m, v2546, int32(0), v2408)
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L1
	} else {
		goto L822
	}
L799:
	;
	v2541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2541, v2540, int32(357895))
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L1
	} else {
		goto L821
	}
L800:
	;
	v2539 = int32(1429)
	v2540 = int32(7437)
	goto L799
L801:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		goto L1
	} else {
		goto L818
	}
L802:
	;
	if v2480 == int32(1114) {
		v2546 = int32(1428)
		goto L798
	} else {
		goto L817
	}
L803:
	;
	v2491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2491 != int32(1) {
		goto L809
	} else {
		goto L810
	}
L804:
	;
	switch v2480 - int32(1082) {
	case 0:
		goto L803
	case 1:
		v2551 = v2408
		goto L797
	default:
		goto L802
	}
L805:
	;
	goto L806
L806:
	;
	if v2480 == int32(1184) {
		goto L800
	} else {
		goto L807
	}
L807:
	;
	if v2480 != int32(1266) {
		goto L801
	} else {
		goto L808
	}
L808:
	;
	v2539 = int32(1427)
	v2540 = int32(7588)
	goto L799
L809:
	;
	v3668 = int32(2)
	goto L7
L810:
	;
	goto L811
L811:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2498 = m.ExcPending
	if v2498 != 0 {
		goto L1
	} else {
		goto L812
	}
L812:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2501 = m.ExcPending
	if v2501 != 0 {
		goto L1
	} else {
		goto L813
	}
L813:
	;
	v2502 = F_text_to_cstring(m, v2145)
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L1
	} else {
		goto L814
	}
L814:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+260)) = v2502
	*(*int32)(unsafe.Add(mBase, uint32(v24)+256)) = int32(357895)
	F_errmsg(m, int32(685515), v24+int32(256))
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L1
	} else {
		goto L815
	}
L815:
	;
	F_errfinish(m, int32(476631), int32(2545), int32(402792))
	mBase = m.M
	v2516 = m.ExcPending
	if v2516 != 0 {
		goto L1
	} else {
		goto L816
	}
L816:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L817:
	;
	goto L801
L818:
	;
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+240)) = v2525
	F_errmsg_internal(m, int32(420741), v24+int32(240))
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L1
	} else {
		goto L819
	}
L819:
	;
	F_errfinish(m, int32(476631), int32(2566), int32(402792))
	mBase = m.M
	v2536 = m.ExcPending
	if v2536 != 0 {
		goto L1
	} else {
		goto L820
	}
L820:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L821:
	;
	v2546 = v2539
	goto L798
L822:
	;
	v2551 = v2548
	goto L797
L823:
	;
	v2556 = F_anytime_typmod_check(m, int32(0), v2410)
	mBase = m.M
	v2557 = m.ExcPending
	if v2557 != 0 {
		goto L1
	} else {
		goto L826
	}
L824:
	;
	v2596 = v2551
	goto L825
L825:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1083)
	v3025 = v2596
	goto L762
L826:
	;
	v2558 = *(*int64)(unsafe.Add(mBase, uint32(v2551)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2558
	v2561 = v24 + int32(752)
	if base.Ui32(v2556) <= base.Ui32(int32(6)) {
		goto L828
	} else {
		goto L829
	}
L827:
	;
	v2591 = *(*int64)(unsafe.Add(mBase, uint32(v24)+752))
	v2592 = F_Int64GetDatum(m, v2591)
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L1
	} else {
		goto L835
	}
L828:
	;
	v2568 = v2556 << (uint(int32(3)) % 32)
	v2571 = *(*int64)(unsafe.Add(mBase, uint32(v2568)+uint32(_consts[1039])))
	v2574 = *(*int64)(unsafe.Add(mBase, uint32(v2568)+uint32(_consts[1040])))
	v2575 = *(*int64)(unsafe.Add(mBase, uint32(v2561)))
	if int64(0) <= v2575 {
		goto L832
	} else {
		goto L833
	}
L829:
	;
	goto L830
L830:
	;
	goto L827
L831:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2561))) = v2585
	goto L830
L832:
	;
	v2578 = v2574 + v2575
	v2579 = base.I64_rem_s(v2578, v2571)
	v2585 = v2578 - v2579
	goto L831
L833:
	;
	goto L834
L834:
	;
	v2581 = v2574 - v2575
	v2582 = base.I64_rem_s(v2581, v2571)
	v2585 = v2582 - v2581
	goto L831
L835:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2556
	v2596 = v2592
	goto L825
L836:
	;
	if v2410 != int32(-1) {
		goto L852
	} else {
		goto L853
	}
L837:
	;
	v2635 = F_DirectFunctionCall1Coll(m, v2633, int32(0), v2408)
	mBase = m.M
	v2636 = m.ExcPending
	if v2636 != 0 {
		goto L1
	} else {
		goto L851
	}
L838:
	;
	v2627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2627, int32(357895), int32(7588))
	mBase = m.M
	v2631 = m.ExcPending
	if v2631 != 0 {
		goto L1
	} else {
		goto L850
	}
L839:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2614 = m.ExcPending
	if v2614 != 0 {
		goto L1
	} else {
		goto L847
	}
L840:
	;
	if v2599 == int32(1114) {
		goto L760
	} else {
		goto L846
	}
L841:
	;
	switch v2599 - int32(1082) {
	case 0:
		goto L760
	case 1:
		goto L838
	default:
		goto L840
	}
L842:
	;
	goto L843
L843:
	;
	if v2599 == int32(1184) {
		v2633 = int32(1430)
		goto L837
	} else {
		goto L844
	}
L844:
	;
	if v2599 != int32(1266) {
		goto L839
	} else {
		goto L845
	}
L845:
	;
	v2637 = v2408
	goto L836
L846:
	;
	goto L839
L847:
	;
	v2615 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+272)) = v2615
	F_errmsg_internal(m, int32(420741), v24+int32(272))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L1
	} else {
		goto L848
	}
L848:
	;
	F_errfinish(m, int32(476631), int32(2613), int32(402792))
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L1
	} else {
		goto L849
	}
L849:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L850:
	;
	v2633 = int32(1431)
	goto L837
L851:
	;
	v2637 = v2635
	goto L836
L852:
	;
	v2641 = F_anytime_typmod_check(m, int32(1), v2410)
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L1
	} else {
		goto L855
	}
L853:
	;
	goto L854
L854:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1266)
	v3025 = v2637
	goto L762
L855:
	;
	if base.Ui32(v2641) <= base.Ui32(int32(6)) {
		goto L857
	} else {
		goto L858
	}
L856:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2641
	goto L854
L857:
	;
	v2649 = v2641 << (uint(int32(3)) % 32)
	v2652 = *(*int64)(unsafe.Add(mBase, uint32(v2649)+uint32(_consts[1039])))
	v2655 = *(*int64)(unsafe.Add(mBase, uint32(v2649)+uint32(_consts[1040])))
	v2656 = *(*int64)(unsafe.Add(mBase, uint32(v2637)))
	if int64(0) <= v2656 {
		goto L861
	} else {
		goto L862
	}
L858:
	;
	goto L859
L859:
	;
	goto L856
L860:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2637))) = v2666
	goto L859
L861:
	;
	v2659 = v2655 + v2656
	v2660 = base.I64_rem_s(v2659, v2652)
	v2666 = v2659 - v2660
	goto L860
L862:
	;
	goto L863
L863:
	;
	v2662 = v2655 - v2656
	v2663 = base.I64_rem_s(v2662, v2652)
	v2666 = v2663 - v2662
	goto L860
L864:
	;
	if v2410 != int32(-1) {
		goto L889
	} else {
		goto L890
	}
L865:
	;
	v2740 = F_DirectFunctionCall1Coll(m, v2738, int32(0), v2408)
	mBase = m.M
	v2741 = m.ExcPending
	if v2741 != 0 {
		goto L1
	} else {
		goto L888
	}
L866:
	;
	v2732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2732, int32(7437), int32(224999))
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L1
	} else {
		goto L887
	}
L867:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2719 = m.ExcPending
	if v2719 != 0 {
		goto L1
	} else {
		goto L884
	}
L868:
	;
	if v2676 == int32(1114) {
		v2742 = v2408
		goto L864
	} else {
		goto L883
	}
L869:
	;
	v2687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2687 != int32(1) {
		goto L875
	} else {
		goto L876
	}
L870:
	;
	switch v2676 - int32(1082) {
	case 0:
		v2738 = int32(1432)
		goto L865
	case 1:
		goto L869
	default:
		goto L868
	}
L871:
	;
	goto L872
L872:
	;
	if v2676 == int32(1184) {
		goto L866
	} else {
		goto L873
	}
L873:
	;
	if v2676 != int32(1266) {
		goto L867
	} else {
		goto L874
	}
L874:
	;
	goto L869
L875:
	;
	v3668 = int32(2)
	goto L7
L876:
	;
	goto L877
L877:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2694 = m.ExcPending
	if v2694 != 0 {
		goto L1
	} else {
		goto L878
	}
L878:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L1
	} else {
		goto L879
	}
L879:
	;
	v2698 = F_text_to_cstring(m, v2145)
	mBase = m.M
	v2699 = m.ExcPending
	if v2699 != 0 {
		goto L1
	} else {
		goto L880
	}
L880:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+340)) = v2698
	*(*int32)(unsafe.Add(mBase, uint32(v24)+336)) = int32(224999)
	F_errmsg(m, int32(685515), v24+int32(336))
	mBase = m.M
	v2707 = m.ExcPending
	if v2707 != 0 {
		goto L1
	} else {
		goto L881
	}
L881:
	;
	F_errfinish(m, int32(476631), int32(2649), int32(402792))
	mBase = m.M
	v2712 = m.ExcPending
	if v2712 != 0 {
		goto L1
	} else {
		goto L882
	}
L882:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L883:
	;
	goto L867
L884:
	;
	v2720 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+304)) = v2720
	F_errmsg_internal(m, int32(420741), v24+int32(304))
	mBase = m.M
	v2726 = m.ExcPending
	if v2726 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	F_errfinish(m, int32(476631), int32(2660), int32(402792))
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L887:
	;
	v2738 = int32(1433)
	goto L865
L888:
	;
	v2742 = v2740
	goto L864
L889:
	;
	v2747 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v2747
	v2750 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2750
	v2753 = F_anytimestamp_typmod_check(m, int32(0), v2410)
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		goto L1
	} else {
		goto L892
	}
L890:
	;
	v2796 = v2742
	goto L891
L891:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1114)
	v3025 = v2796
	goto L762
L892:
	;
	v2755 = *(*int64)(unsafe.Add(mBase, uint32(v2742)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+936)) = v2755
	F_AdjustTimestampForTypmod(m, v24+int32(936), v2753, v24+int32(752))
	mBase = m.M
	v2762 = m.ExcPending
	if v2762 != 0 {
		goto L1
	} else {
		goto L893
	}
L893:
	;
	v2763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	if v2763 == int32(1) {
		goto L894
	} else {
		goto L895
	}
L894:
	;
	v2766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2766 == int32(0) {
		goto L897
	} else {
		goto L898
	}
L895:
	;
	goto L896
L896:
	;
	v2791 = *(*int64)(unsafe.Add(mBase, uint32(v24)+936))
	v2792 = F_Int64GetDatum(m, v2791)
	mBase = m.M
	v2793 = m.ExcPending
	if v2793 != 0 {
		goto L1
	} else {
		goto L905
	}
L897:
	;
	v3668 = int32(2)
	goto L7
L898:
	;
	goto L899
L899:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2773 = m.ExcPending
	if v2773 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2776 = m.ExcPending
	if v2776 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2778 = F_jspOperationName(m, v2777)
	mBase = m.M
	v2779 = m.ExcPending
	if v2779 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+320)) = v2778
	F_errmsg(m, int32(414589), v24+int32(320))
	mBase = m.M
	v2785 = m.ExcPending
	if v2785 != 0 {
		goto L1
	} else {
		goto L903
	}
L903:
	;
	F_errfinish(m, int32(476631), int32(2679), int32(402792))
	mBase = m.M
	v2790 = m.ExcPending
	if v2790 != 0 {
		goto L1
	} else {
		goto L904
	}
L904:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L905:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2753
	v2796 = v2792
	goto L891
L906:
	;
	v2851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2851, int32(339894), int32(7437))
	mBase = m.M
	v2855 = m.ExcPending
	if v2855 != 0 {
		goto L1
	} else {
		goto L925
	}
L907:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L1
	} else {
		goto L922
	}
L908:
	;
	if v2799 == int32(1114) {
		goto L769
	} else {
		goto L921
	}
L909:
	;
	v2808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2808 != int32(1) {
		goto L765
	} else {
		goto L915
	}
L910:
	;
	switch v2799 - int32(1082) {
	case 0:
		goto L906
	case 1:
		goto L909
	default:
		goto L908
	}
L911:
	;
	goto L912
L912:
	;
	if v2799 == int32(1184) {
		v2989 = v2408
		goto L766
	} else {
		goto L913
	}
L913:
	;
	if v2799 != int32(1266) {
		goto L907
	} else {
		goto L914
	}
L914:
	;
	goto L909
L915:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L1
	} else {
		goto L916
	}
L916:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L1
	} else {
		goto L917
	}
L917:
	;
	v2818 = F_text_to_cstring(m, v2145)
	mBase = m.M
	v2819 = m.ExcPending
	if v2819 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+388)) = v2818
	*(*int32)(unsafe.Add(mBase, uint32(v24)+384)) = int32(7688)
	F_errmsg(m, int32(685515), v24+int32(384))
	mBase = m.M
	v2827 = m.ExcPending
	if v2827 != 0 {
		goto L1
	} else {
		goto L919
	}
L919:
	;
	F_errfinish(m, int32(476631), int32(2720), int32(402792))
	mBase = m.M
	v2832 = m.ExcPending
	if v2832 != 0 {
		goto L1
	} else {
		goto L920
	}
L920:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L921:
	;
	goto L907
L922:
	;
	v2839 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+352)) = v2839
	F_errmsg_internal(m, int32(420741), v24+int32(352))
	mBase = m.M
	v2845 = m.ExcPending
	if v2845 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	F_errfinish(m, int32(476631), int32(2741), int32(402792))
	mBase = m.M
	v2850 = m.ExcPending
	if v2850 != 0 {
		goto L1
	} else {
		goto L924
	}
L924:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L925:
	;
	v2867 = v2408 + int32(2483589)
	v2868 = int32(146097)
	v2869 = base.I32_div_u_s(v2867, v2868)
	v2870 = int32(3)
	v2876 = int32(2)
	v2881 = base.I32_div_u_s((v2869*int32(1073595727)+v2867)<<(uint(v2876)%32)|v2870, v2868)
	v2884 = v2408 + int32(2451545) + v2869*v2870 + v2881 + int32(32104)
	v2885 = int32(1461)
	v2886 = base.I32_div_u_s(v2884, v2885)
	v2889 = v2886*int32(-1461) + v2884
	v2891 = v2889 << (uint(v2876) % 32)
	if base.Ui32(v2885) <= base.Ui32(v2891) {
		goto L928
	} else {
		goto L929
	}
L926:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = int64(0)
	v2967 = int32(1434)
	goto L768
L927:
	;
	v2904 = base.I32_div_u_s(v2891, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(772)))) = v2904 + v2886<<(uint(int32(2))%32) - int32(4800)
	v2912 = v2902 + int32(123)
	v2916 = int32(base.Ui32(v2912*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(764)))) = v2912 - int32(base.Ui32(v2916*int32(7834))>>(uint(int32(8))%32))
	v2926 = base.I32_rem_u_s(v2916+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(768)))) = v2926 + int32(1)
	goto L926
L928:
	;
	v2897 = base.I32_rem_u_s(v2889+int32(305), int32(365))
	v2902 = v2897
	goto L927
L929:
	;
	goto L930
L930:
	;
	v2901 = base.I32_rem_u_s(v2889+int32(306), int32(366))
	v2902 = v2901
	goto L927
L931:
	;
	v2939 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+192)) = v2939
	F_errmsg_internal(m, int32(462701), v24+int32(192))
	mBase = m.M
	v2945 = m.ExcPending
	if v2945 != 0 {
		goto L1
	} else {
		goto L932
	}
L932:
	;
	F_errfinish(m, int32(476631), int32(2771), int32(402792))
	mBase = m.M
	v2950 = m.ExcPending
	if v2950 != 0 {
		goto L1
	} else {
		goto L933
	}
L933:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L934:
	;
	v2956 = int32(1435)
	v2957 = *(*int64)(unsafe.Add(mBase, uint32(v2408)))
	v2958 = int32(0)
	v2965 = F_timestamp2tm(m, v2957, v2958, v24+int32(752), v24+int32(948), v2958, v2958)
	mBase = m.M
	v2966 = m.ExcPending
	if v2966 != 0 {
		goto L1
	} else {
		goto L935
	}
L935:
	;
	if v2965 != 0 {
		v2984 = v2956
		goto L767
	} else {
		goto L936
	}
L936:
	;
	v2967 = v2956
	goto L768
L937:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+984)) = v2979
	v2984 = v2967
	goto L767
L938:
	;
	v2989 = v2986
	goto L766
L939:
	;
	v2993 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+944)) = v2993
	v2996 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+936)) = v2996
	v2999 = F_anytimestamp_typmod_check(m, int32(1), v2410)
	mBase = m.M
	v3000 = m.ExcPending
	if v3000 != 0 {
		goto L1
	} else {
		goto L940
	}
L940:
	;
	v3001 = *(*int64)(unsafe.Add(mBase, uint32(v2989)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+888)) = v3001
	F_AdjustTimestampForTypmod(m, v24+int32(888), v2999, v24+int32(936))
	mBase = m.M
	v3008 = m.ExcPending
	if v3008 != 0 {
		goto L1
	} else {
		goto L941
	}
L941:
	;
	v3009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+940)))
	if v3009 != int32(1) {
		goto L764
	} else {
		goto L942
	}
L942:
	;
	v3012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3012 != 0 {
		goto L761
	} else {
		goto L943
	}
L943:
	;
	goto L765
L944:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2999
	v3021 = v3017
	goto L763
L945:
	;
	if v2412&int32(1) == int32(0) {
		goto L946
	} else {
		goto L947
	}
L946:
	;
	v3668 = int32(2)
	goto L7
L947:
	;
	goto L948
L948:
	;
	v3036 = F_jspGetNext(m, l1, v24+int32(952))
	mBase = m.M
	v3037 = m.ExcPending
	if v3037 != 0 {
		goto L1
	} else {
		goto L949
	}
L949:
	;
	if l3 != 0 {
		goto L950
	} else {
		goto L951
	}
L950:
	;
	if v3036 == int32(0) {
		goto L953
	} else {
		goto L954
	}
L951:
	;
	if v3036 != 0 {
		goto L950
	} else {
		goto L952
	}
L952:
	;
	v3668 = int32(0)
	goto L7
L953:
	;
	v3044 = F_palloc(m, int32(20))
	mBase = m.M
	v3045 = m.ExcPending
	if v3045 != 0 {
		goto L1
	} else {
		goto L956
	}
L954:
	;
	v3046 = v24 + int32(908)
	goto L955
L955:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3046)+4)) = v3025
	*(*int32)(unsafe.Add(mBase, uint32(v3046))) = int32(32)
	v3050 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v3046)+8)) = v3050
	v3052 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	*(*int32)(unsafe.Add(mBase, uint32(v3046)+12)) = v3052
	v3054 = *(*int32)(unsafe.Add(mBase, uint32(v24)+984))
	*(*int32)(unsafe.Add(mBase, uint32(v3046)+16)) = v3054
	v3058 = F_executeNextItem(m, l0, l1, v24+int32(952), v3046, l3, v3036)
	mBase = m.M
	v3059 = m.ExcPending
	if v3059 != 0 {
		goto L1
	} else {
		goto L957
	}
L956:
	;
	v3046 = v3044
	goto L955
L957:
	;
	v3668 = v3058
	goto L7
L958:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v3066 = m.ExcPending
	if v3066 != 0 {
		goto L1
	} else {
		goto L959
	}
L959:
	;
	v3067 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3068 = F_jspOperationName(m, v3067)
	mBase = m.M
	v3069 = m.ExcPending
	if v3069 != 0 {
		goto L1
	} else {
		goto L960
	}
L960:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+368)) = v3068
	F_errmsg(m, int32(414589), v24+int32(368))
	mBase = m.M
	v3075 = m.ExcPending
	if v3075 != 0 {
		goto L1
	} else {
		goto L961
	}
L961:
	;
	F_errfinish(m, int32(476631), int32(2760), int32(402792))
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L1
	} else {
		goto L962
	}
L962:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L963:
	;
	v3668 = int32(2)
	goto L7
L964:
	;
	goto L965
L965:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3088 = m.ExcPending
	if v3088 != 0 {
		goto L1
	} else {
		goto L966
	}
L966:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L1
	} else {
		goto L967
	}
L967:
	;
	v3092 = F_text_to_cstring(m, v2145)
	mBase = m.M
	v3093 = m.ExcPending
	if v3093 != 0 {
		goto L1
	} else {
		goto L968
	}
L968:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+292)) = v3092
	*(*int32)(unsafe.Add(mBase, uint32(v24)+288)) = int32(7721)
	F_errmsg(m, int32(685515), v24+int32(288))
	mBase = m.M
	v3101 = m.ExcPending
	if v3101 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	F_errfinish(m, int32(476631), int32(2598), int32(402792))
	mBase = m.M
	v3106 = m.ExcPending
	if v3106 != 0 {
		goto L1
	} else {
		goto L970
	}
L970:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L971:
	;
	v3112 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+208)) = v3112
	F_errmsg_internal(m, int32(420741), v24+int32(208))
	mBase = m.M
	v3118 = m.ExcPending
	if v3118 != 0 {
		goto L1
	} else {
		goto L972
	}
L972:
	;
	F_errfinish(m, int32(476631), int32(2530), int32(402792))
	mBase = m.M
	v3123 = m.ExcPending
	if v3123 != 0 {
		goto L1
	} else {
		goto L973
	}
L973:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L974:
	;
	v3133 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v3133 - int32(1) {
	case 0:
		goto L984
	case 1:
		goto L985
	default:
		goto L983
	}
L975:
	;
	v3126 = F_JsonbType(m, l2)
	mBase = m.M
	v3127 = m.ExcPending
	if v3127 != 0 {
		goto L1
	} else {
		goto L976
	}
L976:
	;
	if v3126 != int32(16) {
		goto L974
	} else {
		goto L977
	}
L977:
	;
	v3131 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		goto L1
	} else {
		goto L978
	}
L978:
	;
	v3668 = v3131
	goto L7
L979:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3323 = m.ExcPending
	if v3323 != 0 {
		goto L1
	} else {
		goto L1029
	}
L980:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3302 = m.ExcPending
	if v3302 != 0 {
		goto L1
	} else {
		goto L1024
	}
L981:
	;
	v3297 = F_executeNextItem(m, l0, l1, int32(0), v3291, l3, int32(1))
	mBase = m.M
	v3298 = m.ExcPending
	if v3298 != 0 {
		goto L1
	} else {
		goto L1023
	}
L982:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(2)
	v3282 = F_Float8GetDatum(m, v3209)
	mBase = m.M
	v3283 = m.ExcPending
	if v3283 != 0 {
		goto L1
	} else {
		goto L1020
	}
L983:
	;
	v3254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3254 != int32(1) {
		v3668 = int32(2)
		goto L7
	} else {
		goto L1014
	}
L984:
	;
	v3195 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3197 = F_pnstrdup(m, v3195, v3196)
	mBase = m.M
	v3198 = m.ExcPending
	if v3198 != 0 {
		goto L1
	} else {
		goto L1000
	}
L985:
	;
	v3138 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3139 = F_DirectFunctionCall1Coll(m, int32(617), int32(0), v3138)
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		goto L1
	} else {
		goto L986
	}
L986:
	;
	v3142 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v3142
	v3145 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v3145
	v3151 = F_float8in_internal(m, v3139, int32(0), int32(259141), v3139, v24+int32(952))
	mBase = m.M
	v3152 = m.ExcPending
	if v3152 != 0 {
		goto L1
	} else {
		goto L987
	}
L987:
	;
	v3153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v3153 == int32(1) {
		goto L989
	} else {
		goto L990
	}
L988:
	;
	v3668 = int32(2)
	goto L7
L989:
	;
	v3156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3156 != int32(1) {
		goto L988
	} else {
		goto L992
	}
L990:
	;
	goto L991
L991:
	;
	v3183 = base.F64_abs(v3151)
	if base.F64_ne(v3183, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v3183)) < base.Ui64(int64(9218868437227405313))) != 0 {
		v3291 = l2
		goto L981
	} else {
		goto L998
	}
L992:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3162 = m.ExcPending
	if v3162 != 0 {
		goto L1
	} else {
		goto L993
	}
L993:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3165 = m.ExcPending
	if v3165 != 0 {
		goto L1
	} else {
		goto L994
	}
L994:
	;
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3167 = F_jspOperationName(m, v3166)
	mBase = m.M
	v3168 = m.ExcPending
	if v3168 != 0 {
		goto L1
	} else {
		goto L995
	}
L995:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = int32(259141)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+100)) = v3167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v3139
	F_errmsg(m, int32(179756), v24+int32(96))
	mBase = m.M
	v3177 = m.ExcPending
	if v3177 != 0 {
		goto L1
	} else {
		goto L996
	}
L996:
	;
	F_errfinish(m, int32(476631), int32(1166), int32(100760))
	mBase = m.M
	v3182 = m.ExcPending
	if v3182 != 0 {
		goto L1
	} else {
		goto L997
	}
L997:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L998:
	;
	v3190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3190 == int32(1) {
		goto L980
	} else {
		goto L999
	}
L999:
	;
	goto L988
L1000:
	;
	v3200 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v3200
	v3203 = *(*int64)(unsafe.Add(mBase, _consts[1077]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v3203
	v3209 = F_float8in_internal(m, v3197, int32(0), int32(259141), v3197, v24+int32(952))
	mBase = m.M
	v3210 = m.ExcPending
	if v3210 != 0 {
		goto L1
	} else {
		goto L1001
	}
L1001:
	;
	v3211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v3211 == int32(1) {
		goto L1003
	} else {
		goto L1004
	}
L1002:
	;
	v3668 = int32(2)
	goto L7
L1003:
	;
	v3214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3214 != int32(1) {
		goto L1002
	} else {
		goto L1006
	}
L1004:
	;
	goto L1005
L1005:
	;
	v3241 = base.F64_abs(v3209)
	if base.F64_ne(v3241, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v3241)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L982
	} else {
		goto L1012
	}
L1006:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3220 = m.ExcPending
	if v3220 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1007:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1008:
	;
	v3224 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3225 = F_jspOperationName(m, v3224)
	mBase = m.M
	v3226 = m.ExcPending
	if v3226 != 0 {
		goto L1
	} else {
		goto L1009
	}
L1009:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = int32(259141)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+132)) = v3225
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v3197
	F_errmsg(m, int32(179756), v24+int32(128))
	mBase = m.M
	v3235 = m.ExcPending
	if v3235 != 0 {
		goto L1
	} else {
		goto L1010
	}
L1010:
	;
	F_errfinish(m, int32(476631), int32(1192), int32(100760))
	mBase = m.M
	v3240 = m.ExcPending
	if v3240 != 0 {
		goto L1
	} else {
		goto L1011
	}
L1011:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1012:
	;
	v3248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3248 == int32(1) {
		goto L979
	} else {
		goto L1013
	}
L1013:
	;
	goto L1002
L1014:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3260 = m.ExcPending
	if v3260 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1015:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	v3264 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3265 = F_jspOperationName(m, v3264)
	mBase = m.M
	v3266 = m.ExcPending
	if v3266 != 0 {
		goto L1
	} else {
		goto L1017
	}
L1017:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v3265
	F_errmsg(m, int32(329215), v24+int32(80))
	mBase = m.M
	v3272 = m.ExcPending
	if v3272 != 0 {
		goto L1
	} else {
		goto L1018
	}
L1018:
	;
	F_errfinish(m, int32(476631), int32(1210), int32(100760))
	mBase = m.M
	v3277 = m.ExcPending
	if v3277 != 0 {
		goto L1
	} else {
		goto L1019
	}
L1019:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1020:
	;
	v3284 = F_DirectFunctionCall1Coll(m, int32(1416), int32(0), v3282)
	mBase = m.M
	v3285 = m.ExcPending
	if v3285 != 0 {
		goto L1
	} else {
		goto L1021
	}
L1021:
	;
	v3286 = F_pg_detoast_datum(m, v3284)
	mBase = m.M
	v3287 = m.ExcPending
	if v3287 != 0 {
		goto L1
	} else {
		goto L1022
	}
L1022:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+756)) = v3286
	v3291 = v24 + int32(752)
	goto L981
L1023:
	;
	v3668 = v3297
	goto L7
L1024:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		goto L1
	} else {
		goto L1025
	}
L1025:
	;
	v3306 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3307 = F_jspOperationName(m, v3306)
	mBase = m.M
	v3308 = m.ExcPending
	if v3308 != 0 {
		goto L1
	} else {
		goto L1026
	}
L1026:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v3307
	F_errmsg(m, int32(641694), v24+int32(112))
	mBase = m.M
	v3314 = m.ExcPending
	if v3314 != 0 {
		goto L1
	} else {
		goto L1027
	}
L1027:
	;
	F_errfinish(m, int32(476631), int32(1171), int32(100760))
	mBase = m.M
	v3319 = m.ExcPending
	if v3319 != 0 {
		goto L1
	} else {
		goto L1028
	}
L1028:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1029:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3326 = m.ExcPending
	if v3326 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	v3327 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3328 = F_jspOperationName(m, v3327)
	mBase = m.M
	v3329 = m.ExcPending
	if v3329 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1031:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v3328
	F_errmsg(m, int32(641694), v24+int32(144))
	mBase = m.M
	v3335 = m.ExcPending
	if v3335 != 0 {
		goto L1
	} else {
		goto L1032
	}
L1032:
	;
	F_errfinish(m, int32(476631), int32(1197), int32(100760))
	mBase = m.M
	v3340 = m.ExcPending
	if v3340 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1033:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1034:
	;
	v3668 = v3342
	goto L7
L1035:
	;
	v3668 = v3345
	goto L7
L1036:
	;
	v3668 = v3348
	goto L7
L1037:
	;
	v3393 = F_palloc(m, int32(20))
	mBase = m.M
	v3394 = m.ExcPending
	if v3394 != 0 {
		goto L1
	} else {
		goto L1051
	}
L1038:
	;
	v3362 = int32(1)
	v3363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v3363 != 0 {
		v3391 = v3362
		goto L1037
	} else {
		goto L1041
	}
L1039:
	;
	v3353 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(v3353)))
	if v3354&int32(1342177280) != int32(1073741824) {
		goto L1038
	} else {
		goto L1040
	}
L1040:
	;
	v3391 = v3354 & int32(268435455)
	goto L1037
L1041:
	;
	v3364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3364 != 0 {
		v3668 = v3362
		goto L7
	} else {
		goto L1042
	}
L1042:
	;
	v3365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3365 != int32(1) {
		goto L1043
	} else {
		goto L1044
	}
L1043:
	;
	v3668 = int32(2)
	goto L7
L1044:
	;
	goto L1045
L1045:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3372 = m.ExcPending
	if v3372 != 0 {
		goto L1
	} else {
		goto L1046
	}
L1046:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v3375 = m.ExcPending
	if v3375 != 0 {
		goto L1
	} else {
		goto L1047
	}
L1047:
	;
	v3376 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3377 = F_jspOperationName(m, v3376)
	mBase = m.M
	v3378 = m.ExcPending
	if v3378 != 0 {
		goto L1
	} else {
		goto L1048
	}
L1048:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v3377
	F_errmsg(m, int32(23563), v24-int32(-64))
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L1
	} else {
		goto L1049
	}
L1049:
	;
	F_errfinish(m, int32(476631), int32(1113), int32(100760))
	mBase = m.M
	v3389 = m.ExcPending
	if v3389 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1050:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1051:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3393))) = int32(2)
	v3398 = F_int64_to_numeric(m, base.I64_extend_i32_u(v3391))
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L1
	} else {
		goto L1052
	}
L1052:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3393)+4)) = v3398
	v3401 = int32(0)
	v3403 = F_executeNextItem(m, l0, l1, v3401, v3393, l3, v3401)
	mBase = m.M
	v3404 = m.ExcPending
	if v3404 != 0 {
		goto L1
	} else {
		goto L1053
	}
L1053:
	;
	v3668 = v3403
	goto L7
L1054:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3406))) = int32(1)
	v3410 = F_JsonbTypeName(m, l2)
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		goto L1
	} else {
		goto L1055
	}
L1055:
	;
	v3412 = F_pstrdup(m, v3410)
	mBase = m.M
	v3413 = m.ExcPending
	if v3413 != 0 {
		goto L1
	} else {
		goto L1056
	}
L1056:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3406)+8)) = v3412
	if v3412&int32(3) == int32(0) {
		v3438 = v3412
		goto L1059
	} else {
		goto L1060
	}
L1057:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3406)+4)) = v3471
	v3473 = int32(0)
	v3475 = F_executeNextItem(m, l0, l1, v3473, v3406, l3, v3473)
	mBase = m.M
	v3476 = m.ExcPending
	if v3476 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1058:
	;
	v3471 = v3463 - v3412
	goto L1057
L1059:
	;
	v3442 = v3438
	goto L1068
L1060:
	;
	v3422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3412))))
	if v3422 == int32(0) {
		goto L1061
	} else {
		goto L1062
	}
L1061:
	;
	v3471 = int32(0)
	goto L1057
L1062:
	;
	goto L1063
L1063:
	;
	v3427 = v3412
	goto L1064
L1064:
	;
	v3431 = v3427 + int32(1)
	if v3431&int32(3) == int32(0) {
		v3438 = v3431
		goto L1059
	} else {
		goto L1066
	}
L1065:
	;
	v3463 = v3431
	goto L1058
L1066:
	;
	v3436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3431))))
	if v3436 != 0 {
		v3427 = v3431
		goto L1064
	} else {
		goto L1067
	}
L1067:
	;
	goto L1065
L1068:
	;
	v3448 = *(*int32)(unsafe.Add(mBase, uint32(v3442)))
	v3451 = int32(-2139062144)
	if (int32(16843008)-v3448|v3448)&v3451 == v3451 {
		v3442 = v3442 + int32(4)
		goto L1068
	} else {
		goto L1070
	}
L1069:
	;
	v3457 = v3442
	goto L1071
L1070:
	;
	goto L1069
L1071:
	;
	v3461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3457))))
	if v3461 != 0 {
		v3457 = v3457 + int32(1)
		goto L1071
	} else {
		goto L1073
	}
L1072:
	;
	v3463 = v3457
	goto L1058
L1073:
	;
	goto L1072
L1074:
	;
	v3668 = v3475
	goto L7
L1075:
	;
	v3501 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3502 = m.ExcPending
	if v3502 != 0 {
		goto L1
	} else {
		goto L1085
	}
L1076:
	;
	v3477 = F_JsonbType(m, l2)
	mBase = m.M
	v3478 = m.ExcPending
	if v3478 != 0 {
		goto L1
	} else {
		goto L1079
	}
L1077:
	;
	goto L1078
L1078:
	;
	F_jspGetArg(m, l1, v24+int32(752))
	mBase = m.M
	v3484 = m.ExcPending
	if v3484 != 0 {
		goto L1
	} else {
		goto L1081
	}
L1079:
	;
	if v3477 == int32(16) {
		goto L1075
	} else {
		goto L1080
	}
L1080:
	;
	goto L1078
L1081:
	;
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
	v3490 = F_executeBoolItem(m, l0, v24+int32(752), l2, int32(0))
	mBase = m.M
	v3491 = m.ExcPending
	if v3491 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1082:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3485
	v3493 = int32(1)
	if v3490 != v3493 {
		v3668 = v3493
		goto L7
	} else {
		goto L1083
	}
L1083:
	;
	v3498 = F_executeNextItem(m, l0, l1, int32(0), l2, l3, int32(1))
	mBase = m.M
	v3499 = m.ExcPending
	if v3499 != 0 {
		goto L1
	} else {
		goto L1084
	}
L1084:
	;
	v3668 = v3498
	goto L7
L1085:
	;
	v3668 = v3501
	goto L7
L1086:
	;
	v3509 = *(*int32)(unsafe.Add(mBase, uint32(v3505)+8))
	v3510 = v3509
	goto L1088
L1087:
	;
	v3510 = int32(0)
	goto L1088
L1088:
	;
	v3511 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3511
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3510
	v3516 = F_executeNextItem(m, l0, l1, v3511, v3505, l3, int32(1))
	mBase = m.M
	v3517 = m.ExcPending
	if v3517 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1089:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v3503
	v3668 = v3516
	goto L7
L1090:
	;
	v3668 = v3522
	goto L7
L1091:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3591 = m.ExcPending
	if v3591 != 0 {
		goto L1
	} else {
		goto L1127
	}
L1092:
	;
	if v3524 == int32(17) {
		goto L1093
	} else {
		goto L1094
	}
L1093:
	;
	v3528 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = v3528
	v3532 = v24 + int32(756)
	if v3532 != 0 {
		goto L1097
	} else {
		goto L1098
	}
L1094:
	;
	goto L1095
L1095:
	;
	if l4 == int32(0) {
		goto L1112
	} else {
		goto L1113
	}
L1096:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v3535
	v3537 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3541 = F_findJsonbValueFromContainer(m, v3537, int32(536870912), v24+int32(752))
	mBase = m.M
	v3542 = m.ExcPending
	if v3542 != 0 {
		goto L1
	} else {
		goto L1100
	}
L1097:
	;
	v3533 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3532))) = v3533
	goto L1099
L1098:
	;
	goto L1099
L1099:
	;
	v3535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L1096
L1100:
	;
	if v3541 != 0 {
		goto L1101
	} else {
		goto L1102
	}
L1101:
	;
	v3543 = int32(0)
	v3545 = F_executeNextItem(m, l0, l1, v3543, v3541, l3, v3543)
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L1
	} else {
		goto L1104
	}
L1102:
	;
	goto L1103
L1103:
	;
	v3552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3552 != 0 {
		v3668 = v3528
		goto L7
	} else {
		goto L1110
	}
L1104:
	;
	if l3 != 0 {
		goto L1105
	} else {
		goto L1106
	}
L1105:
	;
	v3547 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v3547 <= int32(0) {
		v3668 = v3545
		goto L7
	} else {
		goto L1108
	}
L1106:
	;
	goto L1107
L1107:
	;
	F_pfree(m, v3541)
	mBase = m.M
	v3551 = m.ExcPending
	if v3551 != 0 {
		goto L1
	} else {
		goto L1109
	}
L1108:
	;
	goto L1107
L1109:
	;
	v3668 = v3545
	goto L7
L1110:
	;
	v3553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3553 == int32(1) {
		goto L1091
	} else {
		goto L1111
	}
L1111:
	;
	v3668 = int32(2)
	goto L7
L1112:
	;
	v3566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3566 != 0 {
		goto L1117
	} else {
		goto L1118
	}
L1113:
	;
	v3559 = F_JsonbType(m, l2)
	mBase = m.M
	v3560 = m.ExcPending
	if v3560 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1114:
	;
	if v3559 != int32(16) {
		goto L1112
	} else {
		goto L1115
	}
L1115:
	;
	v3564 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3565 = m.ExcPending
	if v3565 != 0 {
		goto L1
	} else {
		goto L1116
	}
L1116:
	;
	v3668 = v3564
	goto L7
L1117:
	;
	v3668 = int32(1)
	goto L7
L1118:
	;
	goto L1119
L1119:
	;
	v3568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3568 != int32(1) {
		goto L1120
	} else {
		goto L1121
	}
L1120:
	;
	v3668 = int32(2)
	goto L7
L1121:
	;
	goto L1122
L1122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3575 = m.ExcPending
	if v3575 != 0 {
		goto L1
	} else {
		goto L1123
	}
L1123:
	;
	F_errcode(m, int32(285999234))
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L1
	} else {
		goto L1124
	}
L1124:
	;
	F_errmsg(m, int32(103193), int32(0))
	mBase = m.M
	v3582 = m.ExcPending
	if v3582 != 0 {
		goto L1
	} else {
		goto L1125
	}
L1125:
	;
	F_errfinish(m, int32(476631), int32(1054), int32(100760))
	mBase = m.M
	v3587 = m.ExcPending
	if v3587 != 0 {
		goto L1
	} else {
		goto L1126
	}
L1126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1127:
	;
	F_errcode(m, int32(285999234))
	mBase = m.M
	v3594 = m.ExcPending
	if v3594 != 0 {
		goto L1
	} else {
		goto L1128
	}
L1128:
	;
	v3595 = *(*int32)(unsafe.Add(mBase, uint32(v24)+760))
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(v24)+756))
	v3597 = F_pnstrdup(m, v3595, v3596)
	mBase = m.M
	v3598 = m.ExcPending
	if v3598 != 0 {
		goto L1
	} else {
		goto L1129
	}
L1129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v3597
	F_errmsg(m, int32(650964), v24+int32(48))
	mBase = m.M
	v3604 = m.ExcPending
	if v3604 != 0 {
		goto L1
	} else {
		goto L1130
	}
L1130:
	;
	F_errfinish(m, int32(476631), int32(1044), int32(100760))
	mBase = m.M
	v3609 = m.ExcPending
	if v3609 != 0 {
		goto L1
	} else {
		goto L1131
	}
L1131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1132:
	;
	v3615 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v3615 != 0 {
		v3626 = int32(1)
		goto L1133
	} else {
		goto L1134
	}
L1133:
	;
	v3628 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v3628 != int32(18) {
		v3668 = v3626
		goto L7
	} else {
		goto L1138
	}
L1134:
	;
	v3616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	v3617 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v3617)
	v3622 = F_executeNextItem(m, l0, l1, v24+int32(752), l2, l3, v3617)
	mBase = m.M
	v3623 = m.ExcPending
	if v3623 != 0 {
		goto L1
	} else {
		goto L1135
	}
L1135:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v3616)
	if l3 != 0 {
		v3626 = v3622
		goto L1133
	} else {
		goto L1136
	}
L1136:
	;
	if v3622 != 0 {
		v3626 = v3622
		goto L1133
	} else {
		goto L1137
	}
L1137:
	;
	v3668 = int32(0)
	goto L7
L1138:
	;
	if v3612 != 0 {
		goto L1139
	} else {
		goto L1140
	}
L1139:
	;
	v3634 = v24 + int32(752)
	goto L1141
L1140:
	;
	v3634 = int32(0)
	goto L1141
L1141:
	;
	v3635 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3636 = int32(1)
	v3637 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3638 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v3641 = F_executeAnyItem(m, l0, v3634, v3635, l3, v3636, v3637, v3638, v3636, v3640)
	mBase = m.M
	v3642 = m.ExcPending
	if v3642 != 0 {
		goto L1
	} else {
		goto L1142
	}
L1142:
	;
	v3668 = v3641
	goto L7
L1143:
	;
	v3644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3644 != int32(1) {
		goto L1144
	} else {
		goto L1145
	}
L1144:
	;
	v3668 = int32(2)
	goto L7
L1145:
	;
	goto L1146
L1146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3651 = m.ExcPending
	if v3651 != 0 {
		goto L1
	} else {
		goto L1147
	}
L1147:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v3654 = m.ExcPending
	if v3654 != 0 {
		goto L1
	} else {
		goto L1148
	}
L1148:
	;
	F_errmsg(m, int32(23442), int32(0))
	mBase = m.M
	v3658 = m.ExcPending
	if v3658 != 0 {
		goto L1
	} else {
		goto L1149
	}
L1149:
	;
	F_errfinish(m, int32(476631), int32(978), int32(100760))
	mBase = m.M
	v3663 = m.ExcPending
	if v3663 != 0 {
		goto L1
	} else {
		goto L1150
	}
L1150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
