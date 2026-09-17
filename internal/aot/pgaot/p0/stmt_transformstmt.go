package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_transformStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int64
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
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
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v713 int32
	_ = v713
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v804 int32
	_ = v804
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v958 int32
	_ = v958
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
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
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1105 int32
	_ = v1105
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1233 int32
	_ = v1233
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int64
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1668 int32
	_ = v1668
	var v1673 int32
	_ = v1673
	var v1681 int32
	_ = v1681
	var v1689 int32
	_ = v1689
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1875 int32
	_ = v1875
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1911 int32
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1934 int32
	_ = v1934
	var v1939 int32
	_ = v1939
	var v1943 int32
	_ = v1943
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2086 int32
	_ = v2086
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2159 int32
	_ = v2159
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2172 int32
	_ = v2172
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2233 int32
	_ = v2233
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2266 int32
	_ = v2266
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2298 int32
	_ = v2298
	var v2300 int32
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
	var v2312 int32
	_ = v2312
	var v2336 int32
	_ = v2336
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2390 int32
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2398 int32
	_ = v2398
	var v2401 int32
	_ = v2401
	var v2404 int32
	_ = v2404
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2429 int32
	_ = v2429
	var v2432 int32
	_ = v2432
	var v2438 int32
	_ = v2438
	var v2442 int32
	_ = v2442
	var v2462 int32
	_ = v2462
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2527 int32
	_ = v2527
	var v2551 int32
	_ = v2551
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2578 int32
	_ = v2578
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2675 int32
	_ = v2675
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2705 int32
	_ = v2705
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2735 int32
	_ = v2735
	var v2739 int32
	_ = v2739
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2749 int32
	_ = v2749
	var v2754 int32
	_ = v2754
	var v2757 int32
	_ = v2757
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2784 int32
	_ = v2784
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2791 int32
	_ = v2791
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2797 int32
	_ = v2797
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2816 int32
	_ = v2816
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2825 int32
	_ = v2825
	var v2831 int32
	_ = v2831
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2840 int32
	_ = v2840
	var v2860 int32
	_ = v2860
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2887 int32
	_ = v2887
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2901 int32
	_ = v2901
	var v2908 int32
	_ = v2908
	var v2936 int32
	_ = v2936
	var v2959 int32
	_ = v2959
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2975 int32
	_ = v2975
	var v2978 int32
	_ = v2978
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2985 int32
	_ = v2985
	var v2990 int32
	_ = v2990
	var v2994 int32
	_ = v2994
	var v2998 int32
	_ = v2998
	var v3003 int32
	_ = v3003
	var v3007 int32
	_ = v3007
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
	var v3023 int32
	_ = v3023
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3030 int32
	_ = v3030
	var v3035 int32
	_ = v3035
	var v3039 int32
	_ = v3039
	var v3042 int32
	_ = v3042
	var v3051 int32
	_ = v3051
	var v3056 int32
	_ = v3056
	var v3060 int32
	_ = v3060
	var v3063 int32
	_ = v3063
	var v3072 int32
	_ = v3072
	var v3077 int32
	_ = v3077
	var v3081 int32
	_ = v3081
	var v3085 int32
	_ = v3085
	var v3090 int32
	_ = v3090
	var v3094 int32
	_ = v3094
	var v3097 int32
	_ = v3097
	var v3101 int32
	_ = v3101
	var v3106 int32
	_ = v3106
	var v3110 int32
	_ = v3110
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3119 int32
	_ = v3119
	var v3124 int32
	_ = v3124
	var v3126 int32
	_ = v3126
	var v3132 int32
	_ = v3132
	var v3136 int32
	_ = v3136
	var v3141 int32
	_ = v3141
	var v3145 int32
	_ = v3145
	var v3148 int32
	_ = v3148
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3154 int32
	_ = v3154
	var v3159 int32
	_ = v3159
	var v3161 int32
	_ = v3161
	var v3167 int32
	_ = v3167
	var v3171 int32
	_ = v3171
	var v3176 int32
	_ = v3176
	var v3180 int32
	_ = v3180
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3185 int32
	_ = v3185
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3189 int32
	_ = v3189
	var v3194 int32
	_ = v3194
	var v3196 int32
	_ = v3196
	var v3202 int32
	_ = v3202
	var v3206 int32
	_ = v3206
	var v3211 int32
	_ = v3211
	var v3215 int32
	_ = v3215
	var v3218 int32
	_ = v3218
	var v3222 int32
	_ = v3222
	var v3227 int32
	_ = v3227
	var v3231 int32
	_ = v3231
	var v3234 int32
	_ = v3234
	var v3238 int32
	_ = v3238
	var v3243 int32
	_ = v3243
	var v3247 int32
	_ = v3247
	var v3250 int32
	_ = v3250
	var v3254 int32
	_ = v3254
	var v3259 int32
	_ = v3259
	var v3263 int32
	_ = v3263
	var v3266 int32
	_ = v3266
	var v3270 int32
	_ = v3270
	var v3275 int32
	_ = v3275
	var v3279 int32
	_ = v3279
	var v3280 int32
	_ = v3280
	var v3286 int32
	_ = v3286
	var v3291 int32
	_ = v3291
	var v3295 int32
	_ = v3295
	var v3301 int32
	_ = v3301
	var v3306 int32
	_ = v3306
	var v3314 int32
	_ = v3314
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3326 int32
	_ = v3326
	var v3337 int32
	_ = v3337
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3346 int32
	_ = v3346
	var v3349 int32
	_ = v3349
	var v3355 int32
	_ = v3355
	var v3379 int32
	_ = v3379
	var v3382 int32
	_ = v3382
	var v3383 int32
	_ = v3383
	var v3385 int32
	_ = v3385
	var v3386 int32
	_ = v3386
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3392 int32
	_ = v3392
	var v3396 int32
	_ = v3396
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3422 int32
	_ = v3422
	var v3424 int32
	_ = v3424
	var v3425 int32
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3427 int32
	_ = v3427
	var v3428 int32
	_ = v3428
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3433 int32
	_ = v3433
	var v3442 int32
	_ = v3442
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3471 int32
	_ = v3471
	var v3495 int32
	_ = v3495
	var v3498 int32
	_ = v3498
	var v3500 int32
	_ = v3500
	var v3504 int32
	_ = v3504
	var v3509 int32
	_ = v3509
	var v3512 int32
	_ = v3512
	var v3514 int32
	_ = v3514
	var v3516 int32
	_ = v3516
	var v3518 int32
	_ = v3518
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3526 int32
	_ = v3526
	var v3541 int32
	_ = v3541
	var v3543 int32
	_ = v3543
	var v3544 int32
	_ = v3544
	var v3546 int32
	_ = v3546
	var v3557 int32
	_ = v3557
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3565 int32
	_ = v3565
	var v3569 int32
	_ = v3569
	var v3572 int32
	_ = v3572
	var v3573 int32
	_ = v3573
	var v3575 int32
	_ = v3575
	var v3579 int32
	_ = v3579
	var v3580 int32
	_ = v3580
	var v3582 int32
	_ = v3582
	var v3585 int32
	_ = v3585
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3589 int32
	_ = v3589
	var v3592 int32
	_ = v3592
	var v3593 int32
	_ = v3593
	var v3594 int32
	_ = v3594
	var v3596 int32
	_ = v3596
	var v3598 int32
	_ = v3598
	var v3601 int32
	_ = v3601
	var v3603 int32
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3607 int32
	_ = v3607
	var v3608 int32
	_ = v3608
	var v3610 int32
	_ = v3610
	var v3613 int32
	_ = v3613
	var v3617 int32
	_ = v3617
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3624 int32
	_ = v3624
	var v3626 int32
	_ = v3626
	var v3631 int32
	_ = v3631
	var v3633 int32
	_ = v3633
	var v3637 int32
	_ = v3637
	var v3642 int32
	_ = v3642
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3648 int32
	_ = v3648
	var v3650 int32
	_ = v3650
	var v3651 int32
	_ = v3651
	var v3653 int32
	_ = v3653
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
	var v3667 int32
	_ = v3667
	var v3668 int32
	_ = v3668
	var v3669 int32
	_ = v3669
	var v3671 int32
	_ = v3671
	var v3672 int32
	_ = v3672
	var v3675 int32
	_ = v3675
	var v3678 int32
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3685 int32
	_ = v3685
	var v3686 int32
	_ = v3686
	var v3688 int32
	_ = v3688
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3695 int32
	_ = v3695
	var v3702 int32
	_ = v3702
	var v3706 int32
	_ = v3706
	var v3726 int32
	_ = v3726
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3734 int32
	_ = v3734
	var v3737 int32
	_ = v3737
	var v3742 int32
	_ = v3742
	var v3743 int32
	_ = v3743
	var v3746 int32
	_ = v3746
	var v3747 int32
	_ = v3747
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3751 int32
	_ = v3751
	var v3753 int32
	_ = v3753
	var v3754 int32
	_ = v3754
	var v3755 int32
	_ = v3755
	var v3757 int32
	_ = v3757
	var v3760 int32
	_ = v3760
	var v3761 int32
	_ = v3761
	var v3763 int32
	_ = v3763
	var v3766 int32
	_ = v3766
	var v3770 int32
	_ = v3770
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3781 int32
	_ = v3781
	var v3799 int32
	_ = v3799
	var v3803 int32
	_ = v3803
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3811 int32
	_ = v3811
	var v3813 int32
	_ = v3813
	var v3818 int32
	_ = v3818
	var v3821 int32
	_ = v3821
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3828 int32
	_ = v3828
	var v3833 int32
	_ = v3833
	var v3834 int32
	_ = v3834
	var v3835 int32
	_ = v3835
	var v3836 int32
	_ = v3836
	var v3837 int32
	_ = v3837
	var v3839 int32
	_ = v3839
	var v3840 int32
	_ = v3840
	var v3842 int32
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3844 int32
	_ = v3844
	var v3846 int32
	_ = v3846
	var v3847 int32
	_ = v3847
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3853 int32
	_ = v3853
	var v3854 int32
	_ = v3854
	var v3858 int32
	_ = v3858
	var v3862 int32
	_ = v3862
	var v3867 int32
	_ = v3867
	var v3868 int32
	_ = v3868
	var v3869 int32
	_ = v3869
	var v3873 int32
	_ = v3873
	var v3874 int32
	_ = v3874
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3880 int32
	_ = v3880
	var v3886 int32
	_ = v3886
	var v3887 int32
	_ = v3887
	var v3889 int32
	_ = v3889
	var v3892 int32
	_ = v3892
	var v3910 int32
	_ = v3910
	var v3914 int32
	_ = v3914
	var v3915 int32
	_ = v3915
	var v3916 int32
	_ = v3916
	var v3917 int32
	_ = v3917
	var v3918 int32
	_ = v3918
	var v3919 int32
	_ = v3919
	var v3920 int32
	_ = v3920
	var v3921 int32
	_ = v3921
	var v3922 int32
	_ = v3922
	var v3924 int32
	_ = v3924
	var v3925 int32
	_ = v3925
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3934 int32
	_ = v3934
	var v3936 int32
	_ = v3936
	var v3939 int32
	_ = v3939
	var v3957 int32
	_ = v3957
	var v3959 int32
	_ = v3959
	var v3963 int32
	_ = v3963
	var v3964 int32
	_ = v3964
	var v3965 int32
	_ = v3965
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3969 int32
	_ = v3969
	var v3972 int32
	_ = v3972
	var v3973 int32
	_ = v3973
	var v3976 int32
	_ = v3976
	var v3977 int32
	_ = v3977
	var v3985 int32
	_ = v3985
	var v4005 int32
	_ = v4005
	var v4006 int32
	_ = v4006
	var v4008 int32
	_ = v4008
	var v4009 int32
	_ = v4009
	var v4019 int32
	_ = v4019
	var v4037 int32
	_ = v4037
	var v4038 int32
	_ = v4038
	var v4039 int32
	_ = v4039
	var v4041 int32
	_ = v4041
	var v4046 int32
	_ = v4046
	var v4070 int32
	_ = v4070
	var v4074 int32
	_ = v4074
	var v4076 int32
	_ = v4076
	var v4080 int32
	_ = v4080
	var v4081 int32
	_ = v4081
	var v4084 int32
	_ = v4084
	var v4086 int32
	_ = v4086
	var v4090 int32
	_ = v4090
	var v4093 int32
	_ = v4093
	var v4097 int32
	_ = v4097
	var v4101 int32
	_ = v4101
	var v4103 int32
	_ = v4103
	var v4106 int32
	_ = v4106
	var v4109 int32
	_ = v4109
	var v4112 int32
	_ = v4112
	var v4114 int32
	_ = v4114
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4121 int32
	_ = v4121
	var v4123 int32
	_ = v4123
	var v4126 int32
	_ = v4126
	var v4130 int32
	_ = v4130
	var v4131 int32
	_ = v4131
	var v4132 int32
	_ = v4132
	var v4134 int32
	_ = v4134
	var v4135 int32
	_ = v4135
	var v4136 int32
	_ = v4136
	var v4137 int32
	_ = v4137
	var v4138 int32
	_ = v4138
	var v4140 int32
	_ = v4140
	var v4141 int32
	_ = v4141
	var v4142 int32
	_ = v4142
	var v4144 int32
	_ = v4144
	var v4145 int32
	_ = v4145
	var v4146 int32
	_ = v4146
	var v4148 int32
	_ = v4148
	var v4150 int32
	_ = v4150
	var v4152 int32
	_ = v4152
	var v4153 int32
	_ = v4153
	var v4156 int32
	_ = v4156
	var v4160 int32
	_ = v4160
	var v4165 int32
	_ = v4165
	var v4171 int32
	_ = v4171
	var v4172 int32
	_ = v4172
	var v4175 int32
	_ = v4175
	var v4176 int32
	_ = v4176
	var v4179 int32
	_ = v4179
	var v4180 int32
	_ = v4180
	var v4183 int32
	_ = v4183
	var v4188 int32
	_ = v4188
	var v4189 int32
	_ = v4189
	var v4206 int32
	_ = v4206
	var v4208 int32
	_ = v4208
	var v4219 int32
	_ = v4219
	var v4223 int32
	_ = v4223
	var v4225 int32
	_ = v4225
	var v4226 int32
	_ = v4226
	var v4229 int32
	_ = v4229
	var v4230 int32
	_ = v4230
	var v4231 int32
	_ = v4231
	var v4235 int32
	_ = v4235
	var v4236 int32
	_ = v4236
	var v4239 int32
	_ = v4239
	var v4240 int32
	_ = v4240
	var v4241 int32
	_ = v4241
	var v4247 int32
	_ = v4247
	var v4248 int32
	_ = v4248
	var v4250 int32
	_ = v4250
	var v4252 int32
	_ = v4252
	var v4255 int32
	_ = v4255
	var v4256 int32
	_ = v4256
	var v4258 int32
	_ = v4258
	var v4259 int32
	_ = v4259
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4263 int32
	_ = v4263
	var v4265 int32
	_ = v4265
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4271 int32
	_ = v4271
	var v4273 int32
	_ = v4273
	var v4274 int32
	_ = v4274
	var v4276 int32
	_ = v4276
	var v4277 int32
	_ = v4277
	var v4295 int32
	_ = v4295
	var v4334 int32
	_ = v4334
	var v4336 int32
	_ = v4336
	var v4337 int32
	_ = v4337
	var v4339 int32
	_ = v4339
	var v4342 int32
	_ = v4342
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4345 int32
	_ = v4345
	var v4346 int32
	_ = v4346
	var v4347 int32
	_ = v4347
	var v4349 int32
	_ = v4349
	var v4355 int32
	_ = v4355
	var v4356 int32
	_ = v4356
	var v4358 int32
	_ = v4358
	var v4363 int32
	_ = v4363
	var v4371 int32
	_ = v4371
	var v4378 int32
	_ = v4378
	var v4383 int32
	_ = v4383
	var v4384 int32
	_ = v4384
	var v4385 int32
	_ = v4385
	var v4386 int32
	_ = v4386
	var v4389 int32
	_ = v4389
	var v4390 int32
	_ = v4390
	var v4392 int32
	_ = v4392
	var v4395 int32
	_ = v4395
	var v4398 int32
	_ = v4398
	var v4399 int32
	_ = v4399
	var v4400 int32
	_ = v4400
	var v4401 int32
	_ = v4401
	var v4402 int32
	_ = v4402
	var v4403 int32
	_ = v4403
	var v4406 int32
	_ = v4406
	var v4409 int32
	_ = v4409
	var v4410 int32
	_ = v4410
	var v4413 int32
	_ = v4413
	var v4428 int32
	_ = v4428
	var v4433 int32
	_ = v4433
	var v4450 int32
	_ = v4450
	var v4453 int32
	_ = v4453
	var v4454 int32
	_ = v4454
	var v4456 int32
	_ = v4456
	var v4466 int32
	_ = v4466
	var v4486 int32
	_ = v4486
	var v4487 int32
	_ = v4487
	var v4492 int32
	_ = v4492
	var v4500 int32
	_ = v4500
	var v4504 int32
	_ = v4504
	var v4509 int32
	_ = v4509
	var v4513 int32
	_ = v4513
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4518 int32
	_ = v4518
	var v4523 int32
	_ = v4523
	var v4528 int32
	_ = v4528
	var v4529 int64
	_ = v4529
	var v4533 int32
	_ = v4533
	var v4534 int32
	_ = v4534
	var v4535 int32
	_ = v4535
	var v4570 int32
	_ = v4570
	var v4573 int32
	_ = v4573
	var v4577 int32
	_ = v4577
	var v4581 int32
	_ = v4581
	var v4582 int32
	_ = v4582
	var v4584 int32
	_ = v4584
	var v4589 int32
	_ = v4589
	var v4593 int32
	_ = v4593
	var v4596 int32
	_ = v4596
	var v4600 int32
	_ = v4600
	var v4601 int32
	_ = v4601
	var v4603 int32
	_ = v4603
	var v4608 int32
	_ = v4608
	var v4612 int32
	_ = v4612
	var v4615 int32
	_ = v4615
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4625 int32
	_ = v4625
	var v4626 int32
	_ = v4626
	var v4628 int32
	_ = v4628
	var v4633 int32
	_ = v4633
	var v4637 int32
	_ = v4637
	var v4640 int32
	_ = v4640
	var v4644 int32
	_ = v4644
	var v4645 int32
	_ = v4645
	var v4647 int32
	_ = v4647
	var v4652 int32
	_ = v4652
	var v4656 int32
	_ = v4656
	var v4659 int32
	_ = v4659
	var v4663 int32
	_ = v4663
	var v4664 int32
	_ = v4664
	var v4666 int32
	_ = v4666
	var v4671 int32
	_ = v4671
	var v4672 int32
	_ = v4672
	var v4674 int32
	_ = v4674
	var v4677 int32
	_ = v4677
	var v4680 int32
	_ = v4680
	var v4683 int32
	_ = v4683
	var v4684 int32
	_ = v4684
	var v4685 int32
	_ = v4685
	var v4686 int32
	_ = v4686
	var v4687 int32
	_ = v4687
	var v4690 int32
	_ = v4690
	var v4691 int32
	_ = v4691
	var v4692 int32
	_ = v4692
	var v4693 int32
	_ = v4693
	var v4694 int32
	_ = v4694
	var v4696 int32
	_ = v4696
	var v4697 int32
	_ = v4697
	var v4699 int32
	_ = v4699
	var v4700 int32
	_ = v4700
	var v4703 int32
	_ = v4703
	var v4705 int32
	_ = v4705
	var v4707 int32
	_ = v4707
	var v4709 int32
	_ = v4709
	var v4743 int32
	_ = v4743
	var v4746 int32
	_ = v4746
	var v4747 int32
	_ = v4747
	var v4749 int32
	_ = v4749
	var v4751 int32
	_ = v4751
	var v4753 int32
	_ = v4753
	var v4754 int32
	_ = v4754
	var v4756 int32
	_ = v4756
	var v4758 int32
	_ = v4758
	var v4761 int32
	_ = v4761
	var v4762 int32
	_ = v4762
	var v4766 int32
	_ = v4766
	var v4767 int32
	_ = v4767
	var v4768 int32
	_ = v4768
	var v4770 int32
	_ = v4770
	var v4771 int32
	_ = v4771
	var v4772 int32
	_ = v4772
	var v4773 int32
	_ = v4773
	var v4774 int32
	_ = v4774
	var v4776 int32
	_ = v4776
	var v4779 int32
	_ = v4779
	var v4780 int32
	_ = v4780
	var v4789 int32
	_ = v4789
	var v4811 int32
	_ = v4811
	v3 = int32(0)
	v26 = int64(0)
	v28 = m.G0
	v30 = v28 - int32(208)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v32 - int32(137) {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	case 3:
		goto L29
	case 4:
		goto L28
	default:
		goto L21
	case 6:
		goto L27
	case 7:
		goto L26
	case 64:
		goto L25
	case 76:
		goto L22
	case 104:
		goto L24
	case 105:
		goto L23
	}
L1:
	;
	v4811 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4789)+24)) = uint8(v4811)
	*(*int32)(unsafe.Add(mBase, uint32(v4789)+8)) = int32(0)
	m.G0 = v30 + int32(208)
	return v4789
L2:
	;
	v3647 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v3648 = int32(0)
	v3650 = F_setTargetTable(m, l0, v3647, v3648, v3648, v66)
	mBase = m.M
	v3651 = m.ExcPending
	if v3651 != 0 {
		goto L33
	} else {
		goto L791
	}
L3:
	;
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3557 != 0 {
		goto L768
	} else {
		goto L769
	}
L4:
	;
	if v1222 <= int32(0) {
		v3541 = v1301
		v3543 = v3
		v3544 = v3
		v3546 = v3
		goto L3
	} else {
		goto L734
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3295 = m.ExcPending
	if v3295 != 0 {
		goto L33
	} else {
		goto L731
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3279 = m.ExcPending
	if v3279 != 0 {
		goto L33
	} else {
		goto L728
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		goto L33
	} else {
		goto L724
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3247 = m.ExcPending
	if v3247 != 0 {
		goto L33
	} else {
		goto L720
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3231 = m.ExcPending
	if v3231 != 0 {
		goto L33
	} else {
		goto L716
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L33
	} else {
		goto L712
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3180 = m.ExcPending
	if v3180 != 0 {
		goto L33
	} else {
		goto L703
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3145 = m.ExcPending
	if v3145 != 0 {
		goto L33
	} else {
		goto L694
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3110 = m.ExcPending
	if v3110 != 0 {
		goto L33
	} else {
		goto L685
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3094 = m.ExcPending
	if v3094 != 0 {
		goto L33
	} else {
		goto L681
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3081 = m.ExcPending
	if v3081 != 0 {
		goto L33
	} else {
		goto L678
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3060 = m.ExcPending
	if v3060 != 0 {
		goto L33
	} else {
		goto L674
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		goto L33
	} else {
		goto L670
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3007 = m.ExcPending
	if v3007 != 0 {
		goto L33
	} else {
		goto L662
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2994 = m.ExcPending
	if v2994 != 0 {
		goto L33
	} else {
		goto L659
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2975 = m.ExcPending
	if v2975 != 0 {
		goto L33
	} else {
		goto L654
	}
L21:
	;
	v2967 = F_palloc0(m, int32(168))
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L33
	} else {
		goto L653
	}
L22:
	;
	v2701 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(v2701)+8))
	if v2702 == int32(0) {
		v2754 = v3
		v2757 = v2701
		goto L603
	} else {
		goto L604
	}
L23:
	;
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2669 = F_transformStmt(m, l0, v2668)
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		goto L33
	} else {
		goto L591
	}
L24:
	;
	v2422 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+204)) = v2422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+200)) = v2422
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v2426 == v2422 {
		goto L552
	} else {
		goto L553
	}
L25:
	;
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2386 = int32(6)
	if v2385&v2386 == v2386 {
		goto L17
	} else {
		goto L537
	}
L26:
	;
	v2009 = F_palloc0(m, int32(168))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L33
	} else {
		goto L452
	}
L27:
	;
	v1961 = F_palloc0(m, int32(168))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L33
	} else {
		goto L442
	}
L28:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v1151 != 0 {
		goto L243
	} else {
		goto L244
	}
L29:
	;
	v206 = m.G0
	v208 = v206 - int32(48)
	m.G0 = v208
	v211 = F_palloc0(m, int32(168))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L33
	} else {
		goto L81
	}
L30:
	;
	v149 = F_palloc0(m, int32(168))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L33
	} else {
		goto L68
	}
L31:
	;
	v86 = F_palloc0(m, int32(168))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L33
	} else {
		goto L55
	}
L32:
	;
	v36 = F_palloc0(m, int32(168))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(67)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(3)
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)) = uint8(v45)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v47 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+41)) = uint8(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v51 = F_transformWithClause(m, l0, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v58 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = v51
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+42)) = uint8(v54)
	goto L37
L39:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v61 == int32(2) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v66 = int64(1)
	goto L41
L41:
	;
	if v42 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v64 = int64(5)
	goto L44
L43:
	;
	v64 = int64(1)
	goto L44
L44:
	;
	v66 = v64
	goto L41
L45:
	;
	v3643 = v3
	v3644 = v3
	v3645 = v3
	v3646 = v3
	goto L2
L46:
	;
	goto L47
L47:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v69 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v3643 = int32(1)
	v3644 = v81
	v3645 = v77
	v3646 = v78
	goto L2
L49:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+44))
	if v72 != 0 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	if v73 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	if v74 != 0 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+60))
	if v75 != 0 {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)+64))
	if v76 != 0 {
		goto L48
	} else {
		goto L54
	}
L54:
	;
	v3643 = v3
	v3644 = v3
	v3645 = v3
	v3646 = v3
	goto L2
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = int64(17179869251)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v90 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+41)) = uint8(v91)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v94 = F_transformWithClause(m, l0, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L33
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+16)))
	v103 = F_setTargetTable(m, l0, v99, v100, int32(1), int64(8))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L33
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+48)) = v94
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+42)) = uint8(v97)
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v103
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+120)) = int32(0)
	v109 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+22)) = uint16(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_transformFromClause(m, l0, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L33
	} else {
		goto L61
	}
L61:
	;
	v114 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+22)) = uint16(v114)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v119 = F_transformWhereClause(m, l0, v116, int32(6), int32(_a_F_transformStmt_0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L33
	} else {
		goto L62
	}
L62:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_transformReturningClause(m, l0, v86, v121, int32(24))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L33
	} else {
		goto L63
	}
L63:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+52)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+56)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v130 = F_makeFromExpr(m, v129, v119)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L33
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+60)) = v130
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+39)) = uint8(v133)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+37)) = uint8(v135)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+38)) = uint8(v137)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+36)) = uint8(v139)
	F_assign_query_collations(m, l0, v86)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L33
	} else {
		goto L65
	}
L65:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v143 != int32(1) {
		v4789 = v86
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_parseCheckAggregates(m, l0, v86)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L33
	} else {
		goto L67
	}
L67:
	;
	v4789 = v86
	goto L1
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v149))) = int64(8589934659)
	v153 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)) = uint8(v153)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v155 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+41)) = uint8(v156)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v159 = F_transformWithClause(m, l0, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L33
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+16)))
	v168 = F_setTargetTable(m, l0, v164, v165, int32(1), int64(4))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L33
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+48)) = v159
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+42)) = uint8(v162)
	goto L71
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+32)) = v168
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v172 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+22)) = uint16(v172)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_transformFromClause(m, l0, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L33
	} else {
		goto L74
	}
L74:
	;
	v177 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+22)) = uint16(v177)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v182 = F_transformWhereClause(m, l0, v179, int32(6), int32(_a_F_transformStmt_0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L33
	} else {
		goto L75
	}
L75:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_transformReturningClause(m, l0, v149, v184, int32(24))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L33
	} else {
		goto L76
	}
L76:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v189 = F_transformUpdateTargetList(m, l0, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L33
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+76)) = v189
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+52)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+56)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v197 = F_makeFromExpr(m, v196, v182)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L33
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+60)) = v197
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+38)) = uint8(v200)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+39)) = uint8(v202)
	F_assign_query_collations(m, l0, v149)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L33
	} else {
		goto L79
	}
L79:
	;
	v4789 = v149
	goto L1
L80:
	;
	v4789 = v211
	goto L1
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = int32(67)
	v215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+41)) = uint8(v215)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = int32(5)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v219 != 0 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+16)))
	v380 = F_setTargetTable(m, l0, v377, v378, int32(0), v375)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L33
	} else {
		goto L116
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L33
	} else {
		goto L112
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L33
	} else {
		goto L108
	}
L85:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+8)))
	if v220 == int32(1) {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+47)) = uint8(v228)
	*(*uint16)(unsafe.Add(mBase, uint32(v208)+45)) = uint16(v228)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v232 == v228 {
		v375 = v26
		goto L82
	} else {
		goto L90
	}
L88:
	;
	v223 = F_transformWithClause(m, l0, v219)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L33
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+48)) = v223
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+42)) = uint8(v226)
	goto L87
L90:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v235 <= int32(0) {
		v375 = v26
		goto L82
	} else {
		goto L91
	}
L91:
	;
	v238 = int32(0)
	if v238 < v235 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v241 = v235
	goto L94
L93:
	;
	v241 = v238
	goto L94
L94:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v249 = v3
	v268 = v26
	goto L95
L95:
	;
	v270 = int32(2)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v242+v249<<(uint(v270)%32))))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+8))
	v276 = v274 - v270
	if base.B2i32(base.Ui32(v276) <= base.Ui32(int32(5)))&(int32(base.Ui32(int32(39))>>(uint(v276)%32))&int32(1)) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v375 = v314
	goto L82
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L33
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v302 = v299 + (v208 + int32(45))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if v303 == int32(1) {
		goto L83
	} else {
		goto L103
	}
L100:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_1), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L33
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(167), int32(_a_F_transformStmt_3))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L33
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v276<<(uint(int32(3))%32))+uint32(_c_F_transformStmt[0])))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	if v309 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v312 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v302))) = uint8(v312)
	goto L106
L105:
	;
	goto L106
L106:
	;
	v314 = v268 | v308
	v316 = v249 + int32(1)
	if v316 != v241 {
		v249 = v316
		v268 = v314
		goto L95
	} else {
		goto L107
	}
L107:
	;
	goto L96
L108:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L33
	} else {
		goto L109
	}
L109:
	;
	F_errmsg(m, int32(_a_F_transformStmt_4), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L33
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(129), int32(_a_F_transformStmt_3))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L33
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L33
	} else {
		goto L113
	}
L113:
	;
	F_errmsg(m, int32(_a_F_transformStmt_5), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L33
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(176), int32(_a_F_transformStmt_3))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L33
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+68)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v211)+32)) = v380
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)+48))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+119)))
	v388 = v386 - int32(112)
	if int32(1)<<(uint(v388)%32)&int32(69) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v396 = base.B2i32(base.Ui32(v388) <= base.Ui32(int32(6)))
	goto L119
L118:
	;
	v396 = int32(0)
	goto L119
L119:
	;
	if v396 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L33
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+32)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v208)+40)) = v424
	v430 = F_list_make1_impl(m, int32(1), v208+int32(32))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L33
	} else {
		goto L128
	}
L123:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L33
	} else {
		goto L124
	}
L124:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v407 + int32(4)
	F_errmsg(m, int32(_a_F_transformStmt_6), v208)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L33
	} else {
		goto L125
	}
L125:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)+48))
	v416 = int32(*(*int8)(unsafe.Add(mBase, uint32(v415)+119)))
	F_errdetail_relkind_not_supported(m, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L33
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(205), int32(_a_F_transformStmt_3))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L33
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_transformFromClause(m, l0, v430)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L33
	} else {
		goto L129
	}
L129:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v434 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	v436 = v435
	goto L132
L131:
	;
	v436 = v3
	goto L132
L132:
	;
	v438 = F_GetNSItemByRangeTablePosn(m, l0, v436, int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L33
	} else {
		goto L133
	}
L133:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444))))
	if base.B2i32(v447 == int32(0))|base.B2i32(v447 != v450) != 0 {
		v468 = v447
		v469 = v450
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v468-v469 != 0 {
		goto L141
	} else {
		goto L142
	}
L135:
	;
	goto L134
L136:
	;
	v453 = v442
	v454 = v444
	goto L137
L137:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+1)))
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+1)))
	if v458 == int32(0) {
		v468 = v458
		v469 = v457
		goto L135
	} else {
		goto L139
	}
L138:
	;
	v468 = v458
	v469 = v457
	goto L135
L139:
	;
	v461 = int32(1)
	if v458 == v457 {
		v453 = v453 + v461
		v454 = v454 + v461
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v471 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+76)) = v471
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+52)) = v473
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+56)) = v475
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v479 = int32(1)
	F_addNSItemToQuery(m, l0, v477, v471, v479, v479)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L33
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L33
	} else {
		goto L238
	}
L144:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v485 = F_transformExpr(m, l0, v483, int32(2))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L33
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+72)) = v485
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v490 = F_makeFromExpr(m, v488, int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L33
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+60)) = v490
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_transformReturningClause(m, l0, v211, v493, int32(25))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L33
	} else {
		goto L147
	}
L147:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v497 == int32(0) {
		v1105 = v3
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+38)) = uint8(v1116)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+64)) = v1105
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+39)) = uint8(v1119)
	F_assign_query_collations(m, l0, v211)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L33
	} else {
		goto L237
	}
L149:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if v500 <= int32(0) {
		v1105 = v3
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v514 = v3
	v519 = v3
	goto L151
L151:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v497)+12))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v530+v514<<(uint(int32(2))%32))))
	v536 = F_palloc0(m, int32(28))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L33
	} else {
		goto L153
	}
L152:
	;
	v1105 = v1083
	goto L148
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = int32(54)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v536)+8)) = v540
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v536)+4)) = v542
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)+12))
	v546 = int32(2)
	v549 = int32(4)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v545+v436<<(uint(v546)%32)-v549)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v211)+32))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v545+v552<<(uint(v546)%32)-v549)))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v542 {
	case 0:
		goto L159
	case 1:
		goto L158
	default:
		goto L157
	}
L154:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v534)+16))
	v921 = F_transformWhereClause(m, l0, v918, int32(18), int32(_a_F_transformStmt_7))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L33
	} else {
		goto L205
	}
L155:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v865)+21)) = uint8(v870)
	*(*uint8)(unsafe.Add(mBase, uint32(v865)+20)) = uint8(v870)
	goto L154
L156:
	;
	v865 = v837
	v870 = int32(1)
	goto L155
L157:
	;
	if v559 == int32(0) {
		goto L154
	} else {
		goto L190
	}
L158:
	;
	if v559 == int32(0) {
		goto L154
	} else {
		goto L175
	}
L159:
	;
	if v559 == int32(0) {
		goto L154
	} else {
		goto L160
	}
L160:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	if v562 <= int32(0) {
		goto L154
	} else {
		goto L161
	}
L161:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v559)+12))
	v570 = int32(0)
	goto L163
L162:
	;
	if v607 == int32(0) {
		goto L154
	} else {
		goto L169
	}
L163:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v565+v570<<(uint(int32(2))%32))))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v558 != v598 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v603 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v597)+20)) = uint16(v603)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v607 = v605
	goto L162
L165:
	;
	v601 = v570 + int32(1)
	if v601 != v562 {
		v570 = v601
		goto L163
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	goto L164
L168:
	;
	v607 = v559
	goto L162
L169:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v607)+4))
	if v610 <= int32(0) {
		goto L154
	} else {
		goto L170
	}
L170:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v607)+12))
	v621 = int32(0)
	goto L171
L171:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v613+v621<<(uint(int32(2))%32))))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v645)+4))
	if v646 == v551 {
		v837 = v645
		goto L156
	} else {
		goto L173
	}
L172:
	;
	goto L154
L173:
	;
	v649 = v621 + int32(1)
	if v649 != v610 {
		v621 = v649
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	if v653 <= int32(0) {
		goto L154
	} else {
		goto L176
	}
L176:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v559)+12))
	v661 = int32(0)
	goto L178
L177:
	;
	if v698 == int32(0) {
		goto L154
	} else {
		goto L184
	}
L178:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v656+v661<<(uint(int32(2))%32))))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	if v558 != v689 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v694 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v688)+20)) = uint16(v694)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v698 = v696
	goto L177
L180:
	;
	v692 = v661 + int32(1)
	if v692 != v653 {
		v661 = v692
		goto L178
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	goto L179
L183:
	;
	v698 = v559
	goto L177
L184:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	if v701 <= int32(0) {
		goto L154
	} else {
		goto L185
	}
L185:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v698)+12))
	v705 = int32(0)
	v713 = v705
	goto L186
L186:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v704+v713<<(uint(int32(2))%32))))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)+4))
	if v738 == v551 {
		v865 = v737
		v870 = v705
		goto L155
	} else {
		goto L188
	}
L187:
	;
	goto L154
L188:
	;
	v741 = v713 + int32(1)
	if v741 != v701 {
		v713 = v741
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	if v745 <= int32(0) {
		goto L154
	} else {
		goto L191
	}
L191:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v559)+12))
	v753 = int32(0)
	goto L193
L192:
	;
	if v790 == int32(0) {
		goto L154
	} else {
		goto L199
	}
L193:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v748+v753<<(uint(int32(2))%32))))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)+4))
	if v558 != v781 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v786 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v780)+20)) = uint16(v786)
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v790 = v788
	goto L192
L195:
	;
	v784 = v753 + int32(1)
	if v784 != v745 {
		v753 = v784
		goto L193
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	goto L194
L198:
	;
	v790 = v559
	goto L192
L199:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v790)+4))
	if v793 <= int32(0) {
		goto L154
	} else {
		goto L200
	}
L200:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v790)+12))
	v804 = int32(0)
	goto L201
L201:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v796+v804<<(uint(int32(2))%32))))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	if v829 == v551 {
		v837 = v828
		goto L156
	} else {
		goto L203
	}
L202:
	;
	goto L154
L203:
	;
	v832 = v804 + int32(1)
	if v832 != v793 {
		v804 = v832
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+16)) = v921
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v536)+8))
	switch v924 - int32(2) {
	case 0:
		goto L207
	case 1:
		goto L210
	case 2:
		goto L206
	default:
		goto L208
	case 5:
		goto L209
	}
L206:
	;
	v1083 = F_lappend(m, v519, v536)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L33
	} else {
		goto L235
	}
L207:
	;
	v1050 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)) = uint8(v1050)
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	v1053 = F_transformUpdateTargetList(m, l0, v1052)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L33
	} else {
		goto L234
	}
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L33
	} else {
		goto L231
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+20)) = int32(0)
	goto L206
L210:
	;
	v927 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)) = uint8(v927)
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	v932 = F_checkInsertTargets(m, l0, v929, v208+int32(36))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L33
	} else {
		goto L211
	}
L211:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v536)+12)) = v934
	v936 = int32(0)
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v534)+24))
	if v938 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v941 = F_transformExpressionList(m, l0, v938, int32(27), int32(1))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L33
	} else {
		goto L215
	}
L213:
	;
	v948 = v936
	goto L214
L214:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v949)+12))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v208)+36))
	v958 = v936
	goto L217
L215:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v208)+36))
	v946 = F_transformInsertRow(m, l0, v941, v943, v932, v944, int32(0))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L33
	} else {
		goto L216
	}
L216:
	;
	v948 = v946
	goto L214
L217:
	;
	v979 = int32(0)
	if v948 == v979 {
		v990 = v979
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if v932 == int32(0) {
		v999 = v979
		goto L222
	} else {
		goto L223
	}
L220:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if v984 <= v958 {
		v990 = int32(0)
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v948)+12))
	v990 = v986 + v958<<(uint(int32(2))%32)
	goto L219
L222:
	;
	if v951 == int32(0) {
		goto L206
	} else {
		goto L225
	}
L223:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v932)+4))
	if v993 <= v958 {
		v999 = v979
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v932)+12))
	v999 = v995 + v958<<(uint(int32(2))%32)
	goto L222
L225:
	;
	v1002 = int32(0)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v951)+4))
	if base.B2i32(v999 == v1002)|(base.B2i32(v990 == v1002)|base.B2i32(v1006 <= v958)) != 0 {
		goto L206
	} else {
		goto L226
	}
L226:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v951)+12))
	if v1010 == int32(0) {
		goto L206
	} else {
		goto L227
	}
L227:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v990)))
	v1017 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1010+v958<<(uint(int32(2))%32)))))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v999)))
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1018)+4))
	v1021 = F_makeTargetEntry(m, v1013, v1017, v1019, int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L33
	} else {
		goto L228
	}
L228:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v536)+20))
	v1024 = F_lappend(m, v1023, v1021)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L33
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+20)) = v1024
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v950)+32))
	v1030 = F_bms_add_member(m, v1027, v1017+int32(7))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L33
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v950)+32)) = v1030
	v958 = v958 + int32(1)
	goto L217
L231:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_1), int32(0))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L33
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(398), int32(_a_F_transformStmt_3))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L33
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+20)) = v1053
	goto L206
L235:
	;
	v1086 = v514 + int32(1)
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if v1086 < v1087 {
		v514 = v1086
		v519 = v1083
		goto L151
	} else {
		goto L236
	}
L236:
	;
	goto L152
L237:
	;
	m.G0 = v208 + int32(48)
	goto L80
L238:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L33
	} else {
		goto L239
	}
L239:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+16)) = v1135
	F_errmsg(m, int32(_a_F_transformStmt_8), v208+int32(16))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L33
	} else {
		goto L240
	}
L240:
	;
	F_errdetail(m, int32(_a_F_transformStmt_9), int32(0))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L33
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(224), int32(_a_F_transformStmt_3))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L33
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
	v1153 = F_palloc0(m, int32(168))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L33
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v1307 == int32(0) {
		goto L278
	} else {
		goto L279
	}
L246:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1153))) = int64(4294967363)
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v1157 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1153)+41)) = uint8(v1158)
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1161 = F_transformWithClause(m, l0, v1160)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L33
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v1166 == int32(0) {
		v3541 = v3
		v3543 = v3
		v3544 = v3
		v3546 = v3
		goto L3
	} else {
		goto L251
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+48)) = v1161
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1153)+42)) = uint8(v1164)
	goto L249
L251:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+4))
	if v1169 <= int32(0) {
		v3541 = v3
		v3543 = v3
		v3544 = v3
		v3546 = v3
		goto L3
	} else {
		goto L252
	}
L252:
	;
	v1179 = v3
	v1181 = int32(-1)
	v1182 = v3
	v1184 = v3
	goto L253
L253:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+12))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1200+v1179<<(uint(int32(2))%32))))
	v1207 = F_transformExpressionList(m, l0, v1204, int32(26), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L33
	} else {
		goto L255
	}
L254:
	;
	goto L4
L255:
	;
	if v1181 < int32(0) {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v1207 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L257:
	;
	if v1207 != 0 {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L259
L259:
	;
	if v1207 != 0 {
		goto L264
	} else {
		goto L265
	}
L260:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+4))
	v1213 = v1211
	goto L262
L261:
	;
	v1213 = int32(0)
	goto L262
L262:
	;
	v1216 = F_palloc0(m, v1213<<(uint(int32(2))%32))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L33
	} else {
		goto L263
	}
L263:
	;
	v1222 = v1213
	v1223 = v1216
	goto L256
L264:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+4))
	v1220 = v1218
	goto L266
L265:
	;
	v1220 = int32(0)
	goto L266
L266:
	;
	if v1220 != v1181 {
		goto L20
	} else {
		goto L267
	}
L267:
	;
	v1222 = v1181
	v1223 = v1182
	goto L256
L268:
	;
	F_list_free(m, v1207)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L33
	} else {
		goto L275
	}
L269:
	;
	v1226 = int32(0)
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+4))
	if v1227 <= v1226 {
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v1233 = v1226
	goto L271
L271:
	;
	v1258 = v1233 << (uint(int32(2)) % 32)
	v1259 = v1223 + v1258
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1259)))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+12))
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1261+v1258)))
	v1264 = F_lappend(m, v1260, v1263)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L33
	} else {
		goto L273
	}
L272:
	;
	goto L268
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1259))) = v1264
	v1268 = v1233 + int32(1)
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+4))
	if v1268 < v1269 {
		v1233 = v1268
		goto L271
	} else {
		goto L274
	}
L274:
	;
	goto L272
L275:
	;
	v1301 = F_lappend(m, v1184, int32(0))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L33
	} else {
		goto L276
	}
L276:
	;
	v1304 = v1179 + int32(1)
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+4))
	if v1304 < v1305 {
		v1179 = v1304
		v1181 = v1222
		v1182 = v1223
		v1184 = v1301
		goto L253
	} else {
		goto L277
	}
L277:
	;
	goto L254
L278:
	;
	v1311 = F_palloc0(m, int32(168))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L33
	} else {
		goto L282
	}
L279:
	;
	goto L280
L280:
	;
	v1539 = m.G0
	v1541 = v1539 - int32(16)
	m.G0 = v1541
	v1544 = F_palloc0(m, int32(168))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L33
	} else {
		goto L335
	}
L281:
	;
	v4789 = v1311
	goto L1
L282:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1311))) = int64(4294967363)
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v1315 != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+41)) = uint8(v1316)
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1319 = F_transformWithClause(m, l0, v1318)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L33
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1324 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+48)) = v1319
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+42)) = uint8(v1322)
	goto L285
L287:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v1327
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1329
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_transformFromClause(m, l0, v1331)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L33
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L33
	} else {
		goto L329
	}
L290:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1336 = F_transformTargetList(m, l0, v1334, int32(14))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L33
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+76)) = v1336
	F_markTargetListOrigins(m, l0, v1336)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L33
	} else {
		goto L292
	}
L292:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1344 = F_transformWhereClause(m, l0, v1341, int32(6), int32(_a_F_transformStmt_0))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L33
	} else {
		goto L293
	}
L293:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v1349 = F_transformWhereClause(m, l0, v1346, int32(7), int32(_a_F_transformStmt_10))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L33
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+112)) = v1349
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v1354 = v1311 + int32(76)
	v1356 = F_transformSortClause(m, l0, v1352, v1354, int32(0))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L33
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+124)) = v1356
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1361 = v1311 + int32(108)
	v1364 = F_transformGroupClause(m, l0, v1359, v1361, v1354, v1356, int32(19), int32(0))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L33
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+100)) = v1364
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+104)) = uint8(v1367)
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1369 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1398 = F_transformLimitClause(m, l0, v1394, int32(23), int32(_a_F_transformStmt_11), v1397)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L33
	} else {
		goto L306
	}
L298:
	;
	v1372 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+40)) = uint8(v1372)
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+120)) = v1372
	goto L297
L299:
	;
	goto L300
L300:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+124))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+12))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)))
	if v1378 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1382 = F_transformDistinctClause(m, l0, v1354, v1376, int32(0))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L33
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	v1387 = F_transformDistinctOnClause(m, l0, v1369, v1354, v1376)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L33
	} else {
		goto L305
	}
L304:
	;
	v1384 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+40)) = uint8(v1384)
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+120)) = v1382
	goto L297
L305:
	;
	v1389 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+40)) = uint8(v1389)
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+120)) = v1387
	goto L297
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+128)) = v1398
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1405 = F_transformLimitClause(m, l0, v1401, int32(22), int32(_a_F_transformStmt_12), v1404)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L33
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+132)) = v1405
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+136)) = v1408
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1411 = F_transformWindowDefinitions(m, l0, v1410, v1354)
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L33
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+116)) = v1411
	v1414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+85)))
	if v1414 == int32(1) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1354)))
	F_resolveTargetListUnknowns(m, l0, v1417)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L33
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+52)) = v1420
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+56)) = v1422
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1425 = F_makeFromExpr(m, v1424, v1344)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L33
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+60)) = v1425
	v1428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+39)) = uint8(v1428)
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+37)) = uint8(v1430)
	v1432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+38)) = uint8(v1432)
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+36)) = uint8(v1434)
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v1436 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	F_assign_query_collations(m, l0, v1311)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L33
	} else {
		goto L321
	}
L315:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+4))
	if v1439 <= int32(0) {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v1444 = int32(0)
	goto L317
L317:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+12))
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1470+v1444<<(uint(int32(2))%32))))
	F_transformLockingClause(m, l0, v1311, v1474, int32(0))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L33
	} else {
		goto L319
	}
L318:
	;
	goto L314
L319:
	;
	v1479 = v1444 + int32(1)
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+4))
	if v1479 < v1480 {
		v1444 = v1479
		goto L317
	} else {
		goto L320
	}
L320:
	;
	goto L318
L321:
	;
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v1511 != 0 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	goto L281
L323:
	;
	F_parseCheckAggregates(m, l0, v1311)
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L33
	} else {
		goto L328
	}
L324:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+100))
	if v1512 != 0 {
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1361)))
	if v1513 != 0 {
		goto L323
	} else {
		goto L326
	}
L326:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+112))
	if v1514 == int32(0) {
		goto L322
	} else {
		goto L327
	}
L327:
	;
	goto L323
L328:
	;
	goto L322
L329:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L33
	} else {
		goto L330
	}
L330:
	;
	F_errmsg(m, int32(_a_F_transformStmt_13), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L33
	} else {
		goto L331
	}
L331:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1531 = F_exprLocation(m, v1530)
	mBase = m.M
	F_parser_errposition(m, l0, v1531)
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L33
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(1403), int32(_a_F_transformStmt_15))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L33
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
	v4789 = v1544
	goto L1
L335:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1544))) = int64(4294967363)
	v1549 = l1
	goto L336
L336:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+76))
	if v1575 != 0 {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1575)+8))
	if v1577 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L338:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1575)+68))
	if v1576 != 0 {
		v1549 = v1575
		goto L336
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	goto L337
L341:
	;
	goto L340
L342:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v1582 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+60)) = v1582
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1585 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v1585
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+44)) = v1582
	if v1581 == v1585 {
		goto L345
	} else {
		goto L346
	}
L343:
	;
	goto L344
L344:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L33
	} else {
		goto L437
	}
L345:
	;
	if v1580 != 0 {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	goto L347
L347:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L33
	} else {
		goto L429
	}
L348:
	;
	v1593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1580)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1544)+41)) = uint8(v1593)
	v1595 = F_transformWithClause(m, l0, v1580)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L33
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	v1602 = F_transformSetOperationTree(m, l0, l1, int32(1), int32(0))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L33
	} else {
		goto L352
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+48)) = v1595
	v1598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1544)+42)) = uint8(v1598)
	goto L350
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+144)) = v1602
	v1606 = v1602
	goto L353
L353:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1606)+12))
	if v1632 != 0 {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+12))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1632)+4))
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1637+v1638<<(uint(int32(2))%32)-int32(4))))
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1644)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+76)) = int32(0)
	v1649 = v1544 + int32(76)
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	if v1650 != 0 {
		goto L359
	} else {
		goto L360
	}
L355:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1632)))
	if v1633 == int32(142) {
		v1606 = v1632
		goto L353
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	goto L354
L358:
	;
	goto L357
L359:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1650)+4))
	v1655 = v1651 << (uint(int32(5)) % 32)
	goto L361
L360:
	;
	v1655 = int32(0)
	goto L361
L361:
	;
	v1656 = F_palloc0(m, v1655)
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L33
	} else {
		goto L362
	}
L362:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+76))
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+28))
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+24))
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	v1668 = v3
	v1673 = v3
	v1681 = v3
	goto L363
L363:
	;
	v1689 = int32(0)
	if v1661 == v1689 {
		v1700 = v1689
		goto L365
	} else {
		goto L366
	}
L364:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L33
	} else {
		goto L422
	}
L365:
	;
	if v1660 == int32(0) {
		v1709 = v1689
		goto L368
	} else {
		goto L369
	}
L366:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+4))
	if v1694 <= v1668 {
		v1700 = int32(0)
		goto L365
	} else {
		goto L367
	}
L367:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+12))
	v1700 = v1696 + v1668<<(uint(int32(2))%32)
	goto L365
L368:
	;
	v1710 = int32(0)
	if v1659 == v1710 {
		v1721 = v1710
		goto L371
	} else {
		goto L372
	}
L369:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1660)+4))
	if v1703 <= v1668 {
		v1709 = v1689
		goto L368
	} else {
		goto L370
	}
L370:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1660)+12))
	v1709 = v1705 + v1668<<(uint(int32(2))%32)
	goto L368
L371:
	;
	if v1658 == int32(0) {
		v1730 = v1710
		goto L374
	} else {
		goto L375
	}
L372:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1659)+4))
	if v1715 <= v1668 {
		v1721 = int32(0)
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1659)+12))
	v1721 = v1717 + v1668<<(uint(int32(2))%32)
	goto L371
L374:
	;
	v1731 = int32(0)
	if v1730 != 0 {
		goto L378
	} else {
		goto L379
	}
L375:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1658)+4))
	if v1724 <= v1668 {
		v1730 = v1710
		goto L374
	} else {
		goto L376
	}
L376:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1658)+12))
	v1730 = v1726 + v1668<<(uint(int32(2))%32)
	goto L374
L377:
	;
	goto L364
L378:
	;
	v1740 = base.B2i32(v1721 == v1731) | (base.B2i32(v1700 == v1731) | base.B2i32(v1709 == v1731))
	goto L380
L379:
	;
	v1740 = int32(1)
	goto L380
L380:
	;
	if v1740 != 0 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1741 = int32(0)
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1743 != 0 {
		goto L384
	} else {
		goto L385
	}
L382:
	;
	goto L383
L383:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1721)))
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1709)))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1700)))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1730)))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+12))
	v1838 = F_pstrdup(m, v1837)
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L33
	} else {
		goto L415
	}
L384:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1743)+4))
	v1745 = v1744
	goto L386
L385:
	;
	v1745 = v1741
	goto L386
L386:
	;
	v1746 = int32(0)
	v1753 = F_addRangeTableEntryForJoin(m, l0, v1673, v1656, v1746, v1746, v1681, v1746, v1746, v1746, v1746, v1746)
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L33
	} else {
		goto L387
	}
L387:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1756 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1756
	F_addNSItemToQuery(m, l0, v1753, v1756, v1756, int32(1))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L33
	} else {
		goto L388
	}
L388:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1649)))
	if v1763 != 0 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1763)+4))
	v1765 = v1764
	goto L391
L390:
	;
	v1765 = v1741
	goto L391
L391:
	;
	v1767 = F_transformSortClause(m, l0, v1588, v1649, int32(0))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L33
	} else {
		goto L392
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+124)) = v1767
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1755
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1772 = int32(0)
	if base.B2i32(v1771 == v1772)|base.B2i32(v1745 <= v1772) != 0 {
		goto L394
	} else {
		goto L395
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1782
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+76))
	if v1784 != 0 {
		goto L400
	} else {
		goto L401
	}
L394:
	;
	v1782 = int32(0)
	goto L396
L395:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+4))
	if v1745 < v1779 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	goto L393
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1771)+4)) = v1745
	goto L399
L398:
	;
	goto L399
L399:
	;
	v1782 = v1771
	goto L396
L400:
	;
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1784)+4))
	v1787 = v1785
	goto L402
L401:
	;
	v1787 = int32(0)
	goto L402
L402:
	;
	if v1787 != v1765 {
		goto L377
	} else {
		goto L403
	}
L403:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1792 = F_transformLimitClause(m, l0, v1587, int32(23), int32(_a_F_transformStmt_11), v1791)
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L33
	} else {
		goto L404
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+128)) = v1792
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1798 = F_transformLimitClause(m, l0, v1584, int32(22), int32(_a_F_transformStmt_12), v1797)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L33
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+132)) = v1798
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+136)) = v1801
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+52)) = v1803
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+56)) = v1805
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1809 = F_makeFromExpr(m, v1807, int32(0))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L33
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+60)) = v1809
	v1812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1544)+39)) = uint8(v1812)
	v1814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1544)+37)) = uint8(v1814)
	v1816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1544)+38)) = uint8(v1816)
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1544)+36)) = uint8(v1818)
	F_assign_query_collations(m, l0, v1544)
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L33
	} else {
		goto L407
	}
L407:
	;
	v1822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v1822 != 0 {
		goto L409
	} else {
		goto L410
	}
L408:
	;
	m.G0 = v1541 + int32(16)
	goto L334
L409:
	;
	F_parseCheckAggregates(m, l0, v1544)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L33
	} else {
		goto L414
	}
L410:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+100))
	if v1823 != 0 {
		goto L409
	} else {
		goto L411
	}
L411:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+108))
	if v1824 != 0 {
		goto L409
	} else {
		goto L412
	}
L412:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+112))
	if v1825 == int32(0) {
		goto L408
	} else {
		goto L413
	}
L413:
	;
	goto L409
L414:
	;
	goto L408
L415:
	;
	v1840 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1836)+8)))
	v1842 = F_makeVar(m, v1638, v1840, v1835, v1834, v1833, int32(0))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L33
	} else {
		goto L416
	}
L416:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+4))
	v1845 = F_exprLocation(m, v1844)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1842)+44)) = v1845
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v1847 + int32(1)
	v1853 = F_makeTargetEntry(m, v1842, base.I32_extend16_s(v1847), v1838, int32(0))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L33
	} else {
		goto L417
	}
L417:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1649)))
	v1856 = F_lappend(m, v1855, v1853)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L33
	} else {
		goto L418
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1649))) = v1856
	v1859 = F_lappend(m, v1681, v1842)
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L33
	} else {
		goto L419
	}
L419:
	;
	v1861 = F_makeString(m, v1838)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L33
	} else {
		goto L420
	}
L420:
	;
	v1863 = F_lappend(m, v1673, v1861)
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L33
	} else {
		goto L421
	}
L421:
	;
	v1867 = v1656 + v1668<<(uint(int32(5))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1867))) = v1638
	v1869 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1836)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v1867)+24)) = v1638
	*(*int32)(unsafe.Add(mBase, uint32(v1867)+16)) = v1833
	*(*int32)(unsafe.Add(mBase, uint32(v1867)+12)) = v1834
	*(*int32)(unsafe.Add(mBase, uint32(v1867)+8)) = v1835
	*(*uint16)(unsafe.Add(mBase, uint32(v1867)+4)) = uint16(v1869)
	v1875 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1836)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1867)+28)) = uint16(v1875)
	v1668 = v1668 + int32(1)
	v1673 = v1863
	v1681 = v1859
	goto L363
L422:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L33
	} else {
		goto L423
	}
L423:
	;
	F_errmsg(m, int32(_a_F_transformStmt_16), int32(0))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L33
	} else {
		goto L424
	}
L424:
	;
	F_errdetail(m, int32(_a_F_transformStmt_17), int32(0))
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L33
	} else {
		goto L425
	}
L425:
	;
	F_errhint(m, int32(_a_F_transformStmt_18), int32(0))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L33
	} else {
		goto L426
	}
L426:
	;
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1649)))
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v1898)+12))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1899+v1765<<(uint(int32(2))%32))))
	v1904 = F_exprLocation(m, v1903)
	mBase = m.M
	F_parser_errposition(m, l0, v1904)
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L33
	} else {
		goto L427
	}
L427:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(1959), int32(_a_F_transformStmt_19))
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L33
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L33
	} else {
		goto L430
	}
L430:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+12))
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v1919)))
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1920)+8))
	v1923 = v1921 - int32(1)
	if base.Ui32(v1923) <= base.Ui32(int32(3)) {
		goto L432
	} else {
		goto L433
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1541))) = v1930
	F_errmsg(m, int32(_a_F_transformStmt_20), v1541)
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L33
	} else {
		goto L435
	}
L432:
	;
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1923<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v1930 = v1928
	goto L434
L433:
	;
	v1930 = int32(_a_F_transformStmt_21)
	goto L434
L434:
	;
	goto L431
L435:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(1817), int32(_a_F_transformStmt_19))
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L33
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L33
	} else {
		goto L438
	}
L438:
	;
	F_errmsg(m, int32(_a_F_transformStmt_13), int32(0))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L33
	} else {
		goto L439
	}
L439:
	;
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1575)+8))
	v1952 = F_exprLocation(m, v1951)
	mBase = m.M
	F_parser_errposition(m, l0, v1952)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L33
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(1790), int32(_a_F_transformStmt_19))
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L33
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
	v1963 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1961)+46)) = uint8(v1963)
	*(*int64)(unsafe.Add(mBase, uint32(v1961))) = int64(4294967363)
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1969 = F_transformExpr(m, l0, v1967, int32(14))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L33
	} else {
		goto L443
	}
L443:
	;
	v1972 = int32(0)
	v1974 = F_makeTargetEntry(m, v1969, int32(1), v1972, v1972)
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L33
	} else {
		goto L444
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v1974
	*(*int32)(unsafe.Add(mBase, uint32(v30)+204)) = v1974
	v1981 = F_list_make1_impl(m, int32(1), v30+int32(12))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L33
	} else {
		goto L445
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1961)+76)) = v1981
	v1984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+85)))
	if v1984 == int32(1) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	F_resolveTargetListUnknowns(m, l0, v1981)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L33
	} else {
		goto L449
	}
L447:
	;
	goto L448
L448:
	;
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1961)+52)) = v1989
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1961)+56)) = v1991
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1995 = F_makeFromExpr(m, v1993, int32(0))
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L33
	} else {
		goto L450
	}
L449:
	;
	goto L448
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1961)+60)) = v1995
	v1998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1961)+39)) = uint8(v1998)
	v2000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1961)+37)) = uint8(v2000)
	v2002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1961)+38)) = uint8(v2002)
	v2004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1961)+36)) = uint8(v2004)
	F_assign_query_collations(m, l0, v1961)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L33
	} else {
		goto L451
	}
L451:
	;
	v4789 = v1961
	goto L1
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009))) = int32(67)
	v2014 = F_palloc0(m, int32(12))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L33
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2014))) = int32(69)
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2022 = F_makeString(m, v2021)
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L33
	} else {
		goto L454
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v2022
	*(*int32)(unsafe.Add(mBase, uint32(v30)+204)) = v2022
	v2029 = F_list_make1_impl(m, int32(1), v30+int32(60))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L33
	} else {
		goto L455
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2014)+4)) = v2029
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2014)+8)) = v2032
	if v2020 < int32(2) {
		v2086 = v2019
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v2111 = F_transformExpr(m, l0, v2014, int32(17))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L33
	} else {
		goto L469
	}
L457:
	;
	v2036 = F_list_copy(m, v2019)
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L33
	} else {
		goto L458
	}
L458:
	;
	if v2036 == int32(0) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v2086 = int32(0)
	goto L456
L460:
	;
	goto L461
L461:
	;
	v2044 = v2036
	v2045 = v2020
	goto L462
L462:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2044)+12))
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2068)))
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v2069)))
	if v2070 != int32(468) {
		goto L19
	} else {
		goto L464
	}
L463:
	;
	v2086 = v2077
	goto L456
L464:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2014)+4))
	v2074 = F_lappend(m, v2073, v2069)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L33
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2014)+4)) = v2074
	v2077 = F_list_delete_first(m, v2044)
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L33
	} else {
		goto L466
	}
L466:
	;
	if v2045 < int32(3) {
		v2086 = v2077
		goto L456
	} else {
		goto L467
	}
L467:
	;
	if v2077 != 0 {
		v2044 = v2077
		v2045 = v2045 - int32(1)
		goto L462
	} else {
		goto L468
	}
L468:
	;
	goto L463
L469:
	;
	v2113 = F_exprType(m, v2111)
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L33
	} else {
		goto L470
	}
L470:
	;
	v2115 = F_exprTypmod(m, v2111)
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L33
	} else {
		goto L471
	}
L471:
	;
	v2117 = F_exprCollation(m, v2111)
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L33
	} else {
		goto L472
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+4)) = int32(1)
	v2121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)) = uint8(v2121)
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v2123
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v2125
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+16))
	F_transformFromClause(m, l0, v2127)
	mBase = m.M
	v2129 = m.ExcPending
	if v2129 != 0 {
		goto L33
	} else {
		goto L473
	}
L473:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+12))
	v2132 = F_transformTargetList(m, l0, v2130, int32(14))
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L33
	} else {
		goto L476
	}
L474:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+12))
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2165)))
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+4))
	v2168 = F_exprType(m, v2167)
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L33
	} else {
		goto L487
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v2152
	F_errmsg_plural(m, int32(_a_F_transformStmt_22), int32(_a_F_transformStmt_23), v2152, v30+int32(16))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L33
	} else {
		goto L485
	}
L476:
	;
	if v2132 != 0 {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+4))
	if v2134 == int32(1) {
		goto L474
	} else {
		goto L480
	}
L478:
	;
	goto L479
L479:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		goto L33
	} else {
		goto L483
	}
L480:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L33
	} else {
		goto L481
	}
L481:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2143 = m.ExcPending
	if v2143 != 0 {
		goto L33
	} else {
		goto L482
	}
L482:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+4))
	v2152 = v2144
	goto L475
L483:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L33
	} else {
		goto L484
	}
L484:
	;
	v2152 = v3
	goto L475
L485:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(2843), int32(_a_F_transformStmt_24))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L33
	} else {
		goto L486
	}
L486:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(17)
	if v2086 != 0 {
		goto L489
	} else {
		goto L490
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+56)) = v2166
	*(*int32)(unsafe.Add(mBase, uint32(v30)+200)) = v2166
	v2210 = F_list_make1_impl(m, int32(1), v30+int32(56))
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L33
	} else {
		goto L505
	}
L489:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2086)+12))
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+4))
	v2177 = F_exprLocation(m, v2111)
	mBase = m.M
	v2178 = F_transformAssignmentIndirection(m, l0, v2111, v2172, int32(0), v2113, v2115, v2117, v2086, v2174, v2175, int32(2), v2177)
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L33
	} else {
		goto L492
	}
L490:
	;
	goto L491
L491:
	;
	if v2113 == v2168 {
		goto L493
	} else {
		goto L494
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2166)+4)) = v2178
	goto L488
L493:
	;
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+4))
	v2193 = int32(2)
	v2196 = F_coerce_to_target_type(m, l0, v2192, v2168, v2113, v2115, v2193, v2193, int32(-1))
	mBase = m.M
	v2197 = m.ExcPending
	if v2197 != 0 {
		goto L33
	} else {
		goto L503
	}
L494:
	;
	if v2113 != int32(2249) {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	v2184 = F_typeOrDomainTypeRelid(m, v2113)
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L33
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	if v2168 == int32(2249) {
		goto L488
	} else {
		goto L500
	}
L498:
	;
	if v2184 == int32(0) {
		goto L493
	} else {
		goto L499
	}
L499:
	;
	goto L497
L500:
	;
	v2190 = F_typeOrDomainTypeRelid(m, v2168)
	mBase = m.M
	v2191 = m.ExcPending
	if v2191 != 0 {
		goto L33
	} else {
		goto L501
	}
L501:
	;
	if v2190 != 0 {
		goto L488
	} else {
		goto L502
	}
L502:
	;
	goto L493
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2166)+4)) = v2196
	if v2196 == int32(0) {
		goto L18
	} else {
		goto L504
	}
L504:
	;
	goto L488
L505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+76)) = v2210
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+20))
	v2216 = F_transformWhereClause(m, l0, v2213, int32(6), int32(_a_F_transformStmt_0))
	mBase = m.M
	v2217 = m.ExcPending
	if v2217 != 0 {
		goto L33
	} else {
		goto L506
	}
L506:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+32))
	v2221 = F_transformWhereClause(m, l0, v2218, int32(7), int32(_a_F_transformStmt_10))
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L33
	} else {
		goto L507
	}
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+112)) = v2221
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+44))
	v2226 = v2009 + int32(76)
	v2228 = F_transformSortClause(m, l0, v2224, v2226, int32(0))
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L33
	} else {
		goto L508
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+124)) = v2228
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+24))
	v2233 = v2009 + int32(108)
	v2236 = F_transformGroupClause(m, l0, v2231, v2233, v2226, v2228, int32(19), int32(0))
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L33
	} else {
		goto L509
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+100)) = v2236
	v2239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2018)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2009)+104)) = uint8(v2239)
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+4))
	if v2241 == int32(0) {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+48))
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+56))
	v2270 = F_transformLimitClause(m, l0, v2266, int32(23), int32(_a_F_transformStmt_11), v2269)
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L33
	} else {
		goto L519
	}
L511:
	;
	v2244 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2009)+40)) = uint8(v2244)
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+120)) = v2244
	goto L510
L512:
	;
	goto L513
L513:
	;
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v2009)+124))
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v2241)+12))
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v2249)))
	if v2250 == int32(0) {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2254 = F_transformDistinctClause(m, l0, v2226, v2248, int32(0))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L33
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	v2259 = F_transformDistinctOnClause(m, l0, v2241, v2226, v2248)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L33
	} else {
		goto L518
	}
L517:
	;
	v2256 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2009)+40)) = uint8(v2256)
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+120)) = v2254
	goto L510
L518:
	;
	v2261 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2009)+40)) = uint8(v2261)
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+120)) = v2259
	goto L510
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+128)) = v2270
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+52))
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+56))
	v2277 = F_transformLimitClause(m, l0, v2273, int32(22), int32(_a_F_transformStmt_12), v2276)
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L33
	} else {
		goto L520
	}
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+132)) = v2277
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+136)) = v2280
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v2283 = F_transformWindowDefinitions(m, l0, v2282, v2226)
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L33
	} else {
		goto L521
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+116)) = v2283
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+52)) = v2286
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+56)) = v2288
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v2291 = F_makeFromExpr(m, v2290, v2216)
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L33
	} else {
		goto L522
	}
L522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+60)) = v2291
	v2294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2009)+39)) = uint8(v2294)
	v2296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2009)+37)) = uint8(v2296)
	v2298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2009)+38)) = uint8(v2298)
	v2300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2009)+36)) = uint8(v2300)
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+60))
	if v2302 == int32(0) {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	F_assign_query_collations(m, l0, v2009)
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L33
	} else {
		goto L530
	}
L524:
	;
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v2302)+4))
	if v2305 <= int32(0) {
		goto L523
	} else {
		goto L525
	}
L525:
	;
	v2312 = int32(0)
	goto L526
L526:
	;
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v2302)+12))
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2336+v2312<<(uint(int32(2))%32))))
	F_transformLockingClause(m, l0, v2009, v2340, int32(0))
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L33
	} else {
		goto L528
	}
L527:
	;
	goto L523
L528:
	;
	v2345 = v2312 + int32(1)
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v2302)+4))
	if v2345 < v2346 {
		v2312 = v2345
		goto L526
	} else {
		goto L529
	}
L529:
	;
	goto L527
L530:
	;
	v2377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v2377 != 0 {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	F_parseCheckAggregates(m, l0, v2009)
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L33
	} else {
		goto L536
	}
L532:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(v2009)+100))
	if v2378 != 0 {
		goto L531
	} else {
		goto L533
	}
L533:
	;
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2233)))
	if v2379 != 0 {
		goto L531
	} else {
		goto L534
	}
L534:
	;
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v2009)+112))
	if v2380 == int32(0) {
		v4789 = v2009
		goto L1
	} else {
		goto L535
	}
L535:
	;
	goto L531
L536:
	;
	v4789 = v2009
	goto L1
L537:
	;
	v2390 = int32(24)
	if v2385&v2390 == v2390 {
		goto L16
	} else {
		goto L538
	}
L538:
	;
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2395 = F_transformStmt(m, l0, v2394)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L33
	} else {
		goto L539
	}
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v2395
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(v2395)))
	if v2398 != int32(67) {
		goto L15
	} else {
		goto L540
	}
L540:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+4))
	if v2401 != int32(1) {
		goto L15
	} else {
		goto L541
	}
L541:
	;
	v2404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2395)+42)))
	if v2404 == int32(1) {
		goto L14
	} else {
		goto L542
	}
L542:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+140))
	if v2407 != 0 {
		goto L543
	} else {
		goto L544
	}
L543:
	;
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v2408&int32(32) != 0 {
		goto L13
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	v2417 = F_palloc0(m, int32(168))
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L33
	} else {
		goto L549
	}
L546:
	;
	if v2408&int32(2) != 0 {
		goto L12
	} else {
		goto L547
	}
L547:
	;
	if v2408&int32(8) != 0 {
		goto L11
	} else {
		goto L548
	}
L548:
	;
	goto L545
L549:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2417)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2417))) = int64(25769803843)
	v4789 = v2417
	goto L1
L550:
	;
	v2663 = F_palloc0(m, int32(168))
	mBase = m.M
	v2664 = m.ExcPending
	if v2664 != 0 {
		goto L33
	} else {
		goto L590
	}
L551:
	;
	v2631 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2632 = F_transformOptionalSelectInto(m, l0, v2631)
	mBase = m.M
	v2633 = m.ExcPending
	if v2633 != 0 {
		goto L33
	} else {
		goto L589
	}
L552:
	;
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v2429 == int32(0) {
		goto L551
	} else {
		goto L555
	}
L553:
	;
	goto L554
L554:
	;
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2521 = *(*int32)(unsafe.Add(mBase, uint32(v2520)))
	if v2521 != int32(141) {
		goto L576
	} else {
		goto L577
	}
L555:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(v2429)+4))
	if v2432 <= int32(0) {
		goto L551
	} else {
		goto L556
	}
L556:
	;
	v2438 = v3
	v2442 = v3
	goto L557
L557:
	;
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2429)+12))
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v2462+v2438<<(uint(int32(2))%32))))
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v2466)+8))
	v2468 = int32(_a_F_transformStmt_25)
	v2471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2467))))
	v2474 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformStmt[2])))
	if base.B2i32(v2471 == int32(0))|base.B2i32(v2471 != v2474) != 0 {
		v2492 = v2471
		v2493 = v2474
		goto L560
	} else {
		goto L561
	}
L558:
	;
	if v2499&int32(1) == int32(0) {
		goto L551
	} else {
		goto L571
	}
L559:
	;
	if v2492-v2493 == int32(0) {
		goto L566
	} else {
		goto L567
	}
L560:
	;
	goto L559
L561:
	;
	v2477 = v2467
	v2478 = v2468
	goto L562
L562:
	;
	v2481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2478)+1)))
	v2482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2477)+1)))
	if v2482 == int32(0) {
		v2492 = v2482
		v2493 = v2481
		goto L560
	} else {
		goto L564
	}
L563:
	;
	v2492 = v2482
	v2493 = v2481
	goto L560
L564:
	;
	v2485 = int32(1)
	if v2482 == v2481 {
		v2477 = v2477 + v2485
		v2478 = v2478 + v2485
		goto L562
	} else {
		goto L565
	}
L565:
	;
	goto L563
L566:
	;
	v2497 = F_defGetBoolean(m, v2466)
	mBase = m.M
	v2498 = m.ExcPending
	if v2498 != 0 {
		goto L33
	} else {
		goto L569
	}
L567:
	;
	v2499 = v2442
	goto L568
L568:
	;
	v2501 = v2438 + int32(1)
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v2429)+4))
	if v2501 < v2502 {
		v2438 = v2501
		v2442 = v2499
		goto L557
	} else {
		goto L570
	}
L569:
	;
	v2499 = v2497
	goto L568
L570:
	;
	goto L558
L571:
	;
	F_setup_parse_variable_parameters(m, l0, v30+int32(204), v30+int32(200))
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L33
	} else {
		goto L572
	}
L572:
	;
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2515 = F_transformOptionalSelectInto(m, l0, v2514)
	mBase = m.M
	v2516 = m.ExcPending
	if v2516 != 0 {
		goto L33
	} else {
		goto L573
	}
L573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2515
	F_check_variable_parameters(m, l0, v2515)
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		goto L33
	} else {
		goto L574
	}
L574:
	;
	goto L550
L575:
	;
	v2601 = F_transformStmt(m, l0, v2578)
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L33
	} else {
		goto L588
	}
L576:
	;
	v2578 = v2520
	goto L575
L577:
	;
	goto L578
L578:
	;
	v2527 = v2520
	goto L580
L579:
	;
	v2557 = *(*int32)(unsafe.Add(mBase, uint32(v2556)+8))
	if v2557 == int32(0) {
		goto L584
	} else {
		goto L585
	}
L580:
	;
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+68))
	if v2551 == int32(0) {
		v2556 = v2527
		goto L579
	} else {
		goto L582
	}
L581:
	;
	v2556 = int32(0)
	goto L579
L582:
	;
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+76))
	if v2554 != 0 {
		v2527 = v2554
		goto L580
	} else {
		goto L583
	}
L583:
	;
	goto L581
L584:
	;
	v2578 = v2520
	goto L575
L585:
	;
	goto L586
L586:
	;
	v2561 = F_palloc0(m, int32(20))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L33
	} else {
		goto L587
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2561)+4)) = v2520
	*(*int32)(unsafe.Add(mBase, uint32(v2561))) = int32(242)
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(v2556)+8))
	v2567 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2561)+16)) = uint8(v2567)
	*(*int32)(unsafe.Add(mBase, uint32(v2561)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v2561)+8)) = v2566
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+8)) = int32(0)
	v2578 = v2561
	goto L575
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2601
	goto L550
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2632
	goto L550
L590:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2663)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2663))) = int64(25769803843)
	v4789 = v2663
	goto L1
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2669
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2672 == int32(23) {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	v2675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2669)+42)))
	if v2675 == int32(1) {
		goto L10
	} else {
		goto L595
	}
L593:
	;
	goto L594
L594:
	;
	v2696 = F_palloc0(m, int32(168))
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L33
	} else {
		goto L602
	}
L595:
	;
	v2678 = F_isQueryUsingTempRelation(m, v2669)
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L33
	} else {
		goto L596
	}
L596:
	;
	if v2678 != 0 {
		goto L9
	} else {
		goto L597
	}
L597:
	;
	v2681 = int32(0)
	v2683 = F_query_tree_walker_impl(m, v2669, int32(496), v2681, v2681)
	mBase = m.M
	v2684 = m.ExcPending
	if v2684 != 0 {
		goto L33
	} else {
		goto L598
	}
L598:
	;
	if v2683 != 0 {
		goto L8
	} else {
		goto L599
	}
L599:
	;
	v2685 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2686 = *(*int32)(unsafe.Add(mBase, uint32(v2685)+4))
	v2687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2686)+17)))
	if v2687 == int32(117) {
		goto L7
	} else {
		goto L600
	}
L600:
	;
	v2690 = F_copyObjectImpl(m, v2669)
	mBase = m.M
	v2691 = m.ExcPending
	if v2691 != 0 {
		goto L33
	} else {
		goto L601
	}
L601:
	;
	v2692 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2692)+28)) = v2690
	goto L594
L602:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2696)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2696))) = int64(25769803843)
	v4789 = v2696
	goto L1
L603:
	;
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v2757)+4))
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v2757)+32))
	v2781 = F_ParseFuncOrColumn(m, l0, v2777, v2754, v2778, v2757, int32(1), v2780)
	mBase = m.M
	v2782 = m.ExcPending
	if v2782 != 0 {
		goto L33
	} else {
		goto L611
	}
L604:
	;
	v2705 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+4))
	if v2705 <= int32(0) {
		v2754 = v3
		v2757 = v2701
		goto L603
	} else {
		goto L605
	}
L605:
	;
	v2711 = v3
	v2712 = v3
	goto L606
L606:
	;
	v2735 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+12))
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2735+v2711<<(uint(int32(2))%32))))
	v2741 = F_transformExpr(m, l0, v2739, int32(41))
	mBase = m.M
	v2742 = m.ExcPending
	if v2742 != 0 {
		goto L33
	} else {
		goto L608
	}
L607:
	;
	v2749 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2754 = v2743
	v2757 = v2749
	goto L603
L608:
	;
	v2743 = F_lappend(m, v2712, v2741)
	mBase = m.M
	v2744 = m.ExcPending
	if v2744 != 0 {
		goto L33
	} else {
		goto L609
	}
L609:
	;
	v2746 = v2711 + int32(1)
	v2747 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+4))
	if v2746 < v2747 {
		v2711 = v2746
		v2712 = v2743
		goto L606
	} else {
		goto L610
	}
L610:
	;
	goto L607
L611:
	;
	F_assign_expr_collations(m, l0, v2781)
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L33
	} else {
		goto L612
	}
L612:
	;
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2781)+4))
	v2787 = F_SearchSysCache1(m, int32(47), v2786)
	mBase = m.M
	v2788 = m.ExcPending
	if v2788 != 0 {
		goto L33
	} else {
		goto L613
	}
L613:
	;
	if v2787 == int32(0) {
		goto L6
	} else {
		goto L614
	}
L614:
	;
	v2791 = *(*int32)(unsafe.Add(mBase, uint32(v2781)+28))
	v2793 = *(*int32)(unsafe.Add(mBase, uint32(v2781)+8))
	v2794 = F_expand_function_arguments(m, v2791, int32(1), v2793, v2787)
	mBase = m.M
	v2795 = m.ExcPending
	if v2795 != 0 {
		goto L33
	} else {
		goto L615
	}
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2781)+28)) = v2794
	v2797 = int32(0)
	v2802 = F_SysCacheGetAttr(m, int32(47), v2787, int32(22), v30+int32(204))
	mBase = m.M
	v2803 = m.ExcPending
	if v2803 != 0 {
		goto L33
	} else {
		goto L616
	}
L616:
	;
	v2804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+204)))
	if v2804 == int32(0) {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v2807 = F_pg_detoast_datum(m, v2802)
	mBase = m.M
	v2808 = m.ExcPending
	if v2808 != 0 {
		goto L33
	} else {
		goto L620
	}
L618:
	;
	v2936 = v2797
	goto L619
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v2936
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v2781
	F_ReleaseCatCache(m, v2787)
	mBase = m.M
	v2959 = m.ExcPending
	if v2959 != 0 {
		goto L33
	} else {
		goto L651
	}
L620:
	;
	v2809 = *(*int32)(unsafe.Add(mBase, uint32(v2781)+28))
	if v2809 != 0 {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v2809)+4))
	v2812 = v2810
	goto L623
L622:
	;
	v2812 = int32(0)
	goto L623
L623:
	;
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(v2807)+4))
	if v2813 != int32(1) {
		goto L5
	} else {
		goto L624
	}
L624:
	;
	v2816 = *(*int32)(unsafe.Add(mBase, uint32(v2807)+16))
	if v2816 != v2812 {
		goto L5
	} else {
		goto L625
	}
L625:
	;
	v2818 = *(*int32)(unsafe.Add(mBase, uint32(v2807)+8))
	if v2818 != 0 {
		goto L5
	} else {
		goto L626
	}
L626:
	;
	v2819 = *(*int32)(unsafe.Add(mBase, uint32(v2807)+12))
	if v2819 != int32(18) {
		goto L5
	} else {
		goto L627
	}
L627:
	;
	if v2809 == int32(0) {
		goto L629
	} else {
		goto L630
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2781)+28)) = v2901
	v2936 = v2908
	goto L619
L629:
	;
	v2901 = int32(0)
	v2908 = v2797
	goto L628
L630:
	;
	goto L631
L631:
	;
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v2809)+4))
	if v2825 <= int32(0) {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	v2901 = int32(0)
	v2908 = v2797
	goto L628
L633:
	;
	goto L634
L634:
	;
	v2831 = int32(0)
	v2833 = v2831
	v2836 = v2831
	v2840 = v2797
	goto L635
L635:
	;
	v2860 = *(*int32)(unsafe.Add(mBase, uint32(v2809)+12))
	v2864 = *(*int32)(unsafe.Add(mBase, uint32(v2860+v2836<<(uint(int32(2))%32))))
	v2865 = v2836 + (v2807 + int32(24))
	v2866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2865))))
	switch v2866 - int32(98) {
	case 0:
		goto L640
	default:
		goto L639
	case 7, 20:
		goto L638
	case 13:
		goto L641
	}
L636:
	;
	v2901 = v2895
	v2908 = v2896
	goto L628
L637:
	;
	v2898 = v2836 + int32(1)
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v2809)+4))
	if v2898 < v2899 {
		v2833 = v2895
		v2836 = v2898
		v2840 = v2896
		goto L635
	} else {
		goto L650
	}
L638:
	;
	v2893 = F_lappend(m, v2833, v2864)
	mBase = m.M
	v2894 = m.ExcPending
	if v2894 != 0 {
		goto L33
	} else {
		goto L649
	}
L639:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2880 = m.ExcPending
	if v2880 != 0 {
		goto L33
	} else {
		goto L646
	}
L640:
	;
	v2871 = F_lappend(m, v2833, v2864)
	mBase = m.M
	v2872 = m.ExcPending
	if v2872 != 0 {
		goto L33
	} else {
		goto L643
	}
L641:
	;
	v2869 = F_lappend(m, v2840, v2864)
	mBase = m.M
	v2870 = m.ExcPending
	if v2870 != 0 {
		goto L33
	} else {
		goto L642
	}
L642:
	;
	v2895 = v2833
	v2896 = v2869
	goto L637
L643:
	;
	v2873 = F_copyObjectImpl(m, v2864)
	mBase = m.M
	v2874 = m.ExcPending
	if v2874 != 0 {
		goto L33
	} else {
		goto L644
	}
L644:
	;
	v2875 = F_lappend(m, v2840, v2873)
	mBase = m.M
	v2876 = m.ExcPending
	if v2876 != 0 {
		goto L33
	} else {
		goto L645
	}
L645:
	;
	v2895 = v2871
	v2896 = v2875
	goto L637
L646:
	;
	v2881 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2865))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+160)) = v2881
	F_errmsg_internal(m, int32(_a_F_transformStmt_26), v30+int32(160))
	mBase = m.M
	v2887 = m.ExcPending
	if v2887 != 0 {
		goto L33
	} else {
		goto L647
	}
L647:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3336), int32(_a_F_transformStmt_27))
	mBase = m.M
	v2892 = m.ExcPending
	if v2892 != 0 {
		goto L33
	} else {
		goto L648
	}
L648:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L649:
	;
	v2895 = v2893
	v2896 = v2840
	goto L637
L650:
	;
	goto L636
L651:
	;
	v2961 = F_palloc0(m, int32(168))
	mBase = m.M
	v2962 = m.ExcPending
	if v2962 != 0 {
		goto L33
	} else {
		goto L652
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2961)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2961))) = int64(25769803843)
	v4789 = v2961
	goto L1
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2967)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2967))) = int64(25769803843)
	v4789 = v2967
	goto L1
L654:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2978 = m.ExcPending
	if v2978 != 0 {
		goto L33
	} else {
		goto L655
	}
L655:
	;
	F_errmsg(m, int32(_a_F_transformStmt_28), int32(0))
	mBase = m.M
	v2982 = m.ExcPending
	if v2982 != 0 {
		goto L33
	} else {
		goto L656
	}
L656:
	;
	v2983 = F_exprLocation(m, v1207)
	mBase = m.M
	F_parser_errposition(m, l0, v2983)
	mBase = m.M
	v2985 = m.ExcPending
	if v2985 != 0 {
		goto L33
	} else {
		goto L657
	}
L657:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(1596), int32(_a_F_transformStmt_29))
	mBase = m.M
	v2990 = m.ExcPending
	if v2990 != 0 {
		goto L33
	} else {
		goto L658
	}
L658:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L659:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_30), int32(0))
	mBase = m.M
	v2998 = m.ExcPending
	if v2998 != 0 {
		goto L33
	} else {
		goto L660
	}
L660:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(2800), int32(_a_F_transformStmt_24))
	mBase = m.M
	v3003 = m.ExcPending
	if v3003 != 0 {
		goto L33
	} else {
		goto L661
	}
L661:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L662:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v3010 = m.ExcPending
	if v3010 != 0 {
		goto L33
	} else {
		goto L663
	}
L663:
	;
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v3012 = F_format_type_be(m, v2113)
	mBase = m.M
	v3013 = m.ExcPending
	if v3013 != 0 {
		goto L33
	} else {
		goto L664
	}
L664:
	;
	v3014 = F_format_type_be(m, v2168)
	mBase = m.M
	v3015 = m.ExcPending
	if v3015 != 0 {
		goto L33
	} else {
		goto L665
	}
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v3014
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v3012
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v3011
	F_errmsg(m, int32(_a_F_transformStmt_31), v30+int32(32))
	mBase = m.M
	v3023 = m.ExcPending
	if v3023 != 0 {
		goto L33
	} else {
		goto L666
	}
L666:
	;
	F_errhint(m, int32(_a_F_transformStmt_32), int32(0))
	mBase = m.M
	v3027 = m.ExcPending
	if v3027 != 0 {
		goto L33
	} else {
		goto L667
	}
L667:
	;
	v3028 = F_exprLocation(m, v2192)
	mBase = m.M
	F_parser_errposition(m, l0, v3028)
	mBase = m.M
	v3030 = m.ExcPending
	if v3030 != 0 {
		goto L33
	} else {
		goto L668
	}
L668:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(2907), int32(_a_F_transformStmt_24))
	mBase = m.M
	v3035 = m.ExcPending
	if v3035 != 0 {
		goto L33
	} else {
		goto L669
	}
L669:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L670:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v3042 = m.ExcPending
	if v3042 != 0 {
		goto L33
	} else {
		goto L671
	}
L671:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = int32(_a_F_transformStmt_33)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = int32(_a_F_transformStmt_34)
	F_errmsg(m, int32(_a_F_transformStmt_35), v30-int32(-64))
	mBase = m.M
	v3051 = m.ExcPending
	if v3051 != 0 {
		goto L33
	} else {
		goto L672
	}
L672:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3028), int32(_a_F_transformStmt_36))
	mBase = m.M
	v3056 = m.ExcPending
	if v3056 != 0 {
		goto L33
	} else {
		goto L673
	}
L673:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L674:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v3063 = m.ExcPending
	if v3063 != 0 {
		goto L33
	} else {
		goto L675
	}
L675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+84)) = int32(_a_F_transformStmt_37)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = int32(_a_F_transformStmt_38)
	F_errmsg(m, int32(_a_F_transformStmt_35), v30+int32(80))
	mBase = m.M
	v3072 = m.ExcPending
	if v3072 != 0 {
		goto L33
	} else {
		goto L676
	}
L676:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3036), int32(_a_F_transformStmt_36))
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		goto L33
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
	F_errmsg_internal(m, int32(_a_F_transformStmt_39), int32(0))
	mBase = m.M
	v3085 = m.ExcPending
	if v3085 != 0 {
		goto L33
	} else {
		goto L679
	}
L679:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3045), int32(_a_F_transformStmt_36))
	mBase = m.M
	v3090 = m.ExcPending
	if v3090 != 0 {
		goto L33
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v3097 = m.ExcPending
	if v3097 != 0 {
		goto L33
	} else {
		goto L682
	}
L682:
	;
	F_errmsg(m, int32(_a_F_transformStmt_40), int32(0))
	mBase = m.M
	v3101 = m.ExcPending
	if v3101 != 0 {
		goto L33
	} else {
		goto L683
	}
L683:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3055), int32(_a_F_transformStmt_36))
	mBase = m.M
	v3106 = m.ExcPending
	if v3106 != 0 {
		goto L33
	} else {
		goto L684
	}
L684:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L685:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3113 = m.ExcPending
	if v3113 != 0 {
		goto L33
	} else {
		goto L686
	}
L686:
	;
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+140))
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v3114)+12))
	v3116 = *(*int32)(unsafe.Add(mBase, uint32(v3115)))
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(v3116)+8))
	v3119 = v3117 - int32(1)
	if base.Ui32(v3119) <= base.Ui32(int32(3)) {
		goto L688
	} else {
		goto L689
	}
L687:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+128)) = v3126
	F_errmsg(m, int32(_a_F_transformStmt_41), v30+int32(128))
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		goto L33
	} else {
		goto L691
	}
L688:
	;
	v3124 = *(*int32)(unsafe.Add(mBase, uint32(v3119<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v3126 = v3124
	goto L690
L689:
	;
	v3126 = int32(_a_F_transformStmt_21)
	goto L690
L690:
	;
	goto L687
L691:
	;
	F_errdetail(m, int32(_a_F_transformStmt_42), int32(0))
	mBase = m.M
	v3136 = m.ExcPending
	if v3136 != 0 {
		goto L33
	} else {
		goto L692
	}
L692:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3066), int32(_a_F_transformStmt_36))
	mBase = m.M
	v3141 = m.ExcPending
	if v3141 != 0 {
		goto L33
	} else {
		goto L693
	}
L693:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L694:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3148 = m.ExcPending
	if v3148 != 0 {
		goto L33
	} else {
		goto L695
	}
L695:
	;
	v3149 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+140))
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v3149)+12))
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(v3150)))
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(v3151)+8))
	v3154 = v3152 - int32(1)
	if base.Ui32(v3154) <= base.Ui32(int32(3)) {
		goto L697
	} else {
		goto L698
	}
L696:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v3161
	F_errmsg(m, int32(_a_F_transformStmt_43), v30+int32(112))
	mBase = m.M
	v3167 = m.ExcPending
	if v3167 != 0 {
		goto L33
	} else {
		goto L700
	}
L697:
	;
	v3159 = *(*int32)(unsafe.Add(mBase, uint32(v3154<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v3161 = v3159
	goto L699
L698:
	;
	v3161 = int32(_a_F_transformStmt_21)
	goto L699
L699:
	;
	goto L696
L700:
	;
	F_errdetail(m, int32(_a_F_transformStmt_44), int32(0))
	mBase = m.M
	v3171 = m.ExcPending
	if v3171 != 0 {
		goto L33
	} else {
		goto L701
	}
L701:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3077), int32(_a_F_transformStmt_36))
	mBase = m.M
	v3176 = m.ExcPending
	if v3176 != 0 {
		goto L33
	} else {
		goto L702
	}
L702:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L703:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v3183 = m.ExcPending
	if v3183 != 0 {
		goto L33
	} else {
		goto L704
	}
L704:
	;
	v3184 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+140))
	v3185 = *(*int32)(unsafe.Add(mBase, uint32(v3184)+12))
	v3186 = *(*int32)(unsafe.Add(mBase, uint32(v3185)))
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v3186)+8))
	v3189 = v3187 - int32(1)
	if base.Ui32(v3189) <= base.Ui32(int32(3)) {
		goto L706
	} else {
		goto L707
	}
L705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v3196
	F_errmsg(m, int32(_a_F_transformStmt_45), v30+int32(96))
	mBase = m.M
	v3202 = m.ExcPending
	if v3202 != 0 {
		goto L33
	} else {
		goto L709
	}
L706:
	;
	v3194 = *(*int32)(unsafe.Add(mBase, uint32(v3189<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v3196 = v3194
	goto L708
L707:
	;
	v3196 = int32(_a_F_transformStmt_21)
	goto L708
L708:
	;
	goto L705
L709:
	;
	F_errdetail(m, int32(_a_F_transformStmt_46), int32(0))
	mBase = m.M
	v3206 = m.ExcPending
	if v3206 != 0 {
		goto L33
	} else {
		goto L710
	}
L710:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3088), int32(_a_F_transformStmt_36))
	mBase = m.M
	v3211 = m.ExcPending
	if v3211 != 0 {
		goto L33
	} else {
		goto L711
	}
L711:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L712:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L33
	} else {
		goto L713
	}
L713:
	;
	F_errmsg(m, int32(_a_F_transformStmt_47), int32(0))
	mBase = m.M
	v3222 = m.ExcPending
	if v3222 != 0 {
		goto L33
	} else {
		goto L714
	}
L714:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3182), int32(_a_F_transformStmt_48))
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L33
	} else {
		goto L715
	}
L715:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L716:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3234 = m.ExcPending
	if v3234 != 0 {
		goto L33
	} else {
		goto L717
	}
L717:
	;
	F_errmsg(m, int32(_a_F_transformStmt_49), int32(0))
	mBase = m.M
	v3238 = m.ExcPending
	if v3238 != 0 {
		goto L33
	} else {
		goto L718
	}
L718:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3192), int32(_a_F_transformStmt_48))
	mBase = m.M
	v3243 = m.ExcPending
	if v3243 != 0 {
		goto L33
	} else {
		goto L719
	}
L719:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L720:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3250 = m.ExcPending
	if v3250 != 0 {
		goto L33
	} else {
		goto L721
	}
L721:
	;
	F_errmsg(m, int32(_a_F_transformStmt_50), int32(0))
	mBase = m.M
	v3254 = m.ExcPending
	if v3254 != 0 {
		goto L33
	} else {
		goto L722
	}
L722:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3202), int32(_a_F_transformStmt_48))
	mBase = m.M
	v3259 = m.ExcPending
	if v3259 != 0 {
		goto L33
	} else {
		goto L723
	}
L723:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L724:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3266 = m.ExcPending
	if v3266 != 0 {
		goto L33
	} else {
		goto L725
	}
L725:
	;
	F_errmsg(m, int32(_a_F_transformStmt_51), int32(0))
	mBase = m.M
	v3270 = m.ExcPending
	if v3270 != 0 {
		goto L33
	} else {
		goto L726
	}
L726:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3214), int32(_a_F_transformStmt_48))
	mBase = m.M
	v3275 = m.ExcPending
	if v3275 != 0 {
		goto L33
	} else {
		goto L727
	}
L727:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L728:
	;
	v3280 = *(*int32)(unsafe.Add(mBase, uint32(v2781)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+144)) = v3280
	F_errmsg_internal(m, int32(_a_F_transformStmt_52), v30+int32(144))
	mBase = m.M
	v3286 = m.ExcPending
	if v3286 != 0 {
		goto L33
	} else {
		goto L729
	}
L729:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3275), int32(_a_F_transformStmt_27))
	mBase = m.M
	v3291 = m.ExcPending
	if v3291 != 0 {
		goto L33
	} else {
		goto L730
	}
L730:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L731:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+176)) = v2812
	F_errmsg_internal(m, int32(_a_F_transformStmt_53), v30+int32(176))
	mBase = m.M
	v3301 = m.ExcPending
	if v3301 != 0 {
		goto L33
	} else {
		goto L732
	}
L732:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(3311), int32(_a_F_transformStmt_27))
	mBase = m.M
	v3306 = m.ExcPending
	if v3306 != 0 {
		goto L33
	} else {
		goto L733
	}
L733:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L734:
	;
	v3314 = int32(0)
	v3323 = v3
	v3324 = v3
	v3326 = v3
	goto L735
L735:
	;
	v3337 = int32(0)
	v3340 = v1223 + v3314<<(uint(int32(2))%32)
	v3341 = *(*int32)(unsafe.Add(mBase, uint32(v3340)))
	v3344 = F_select_common_type(m, l0, v3341, int32(_a_F_transformStmt_54), v3337)
	mBase = m.M
	v3345 = m.ExcPending
	if v3345 != 0 {
		goto L33
	} else {
		goto L737
	}
L736:
	;
	v3442 = int32(0)
	goto L753
L737:
	;
	v3346 = *(*int32)(unsafe.Add(mBase, uint32(v3340)))
	if v3346 == int32(0) {
		v3396 = v3337
		goto L738
	} else {
		goto L739
	}
L738:
	;
	v3420 = F_select_common_typmod(m, v3396, v3344)
	mBase = m.M
	v3421 = m.ExcPending
	if v3421 != 0 {
		goto L33
	} else {
		goto L747
	}
L739:
	;
	v3349 = *(*int32)(unsafe.Add(mBase, uint32(v3346)+4))
	if v3349 <= int32(0) {
		goto L740
	} else {
		goto L741
	}
L740:
	;
	v3396 = v3346
	goto L738
L741:
	;
	goto L742
L742:
	;
	v3355 = v3337
	goto L743
L743:
	;
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3346)+12))
	v3382 = v3379 + v3355<<(uint(int32(2))%32)
	v3383 = *(*int32)(unsafe.Add(mBase, uint32(v3382)))
	v3385 = F_coerce_to_common_type(m, l0, v3383, v3344, int32(_a_F_transformStmt_54))
	mBase = m.M
	v3386 = m.ExcPending
	if v3386 != 0 {
		goto L33
	} else {
		goto L745
	}
L744:
	;
	v3392 = *(*int32)(unsafe.Add(mBase, uint32(v3340)))
	v3396 = v3392
	goto L738
L745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3382))) = v3385
	v3389 = v3355 + int32(1)
	v3390 = *(*int32)(unsafe.Add(mBase, uint32(v3346)+4))
	if v3389 < v3390 {
		v3355 = v3389
		goto L743
	} else {
		goto L746
	}
L746:
	;
	goto L744
L747:
	;
	v3422 = *(*int32)(unsafe.Add(mBase, uint32(v3340)))
	v3424 = F_select_common_collation(m, l0, v3422, int32(1))
	mBase = m.M
	v3425 = m.ExcPending
	if v3425 != 0 {
		goto L33
	} else {
		goto L748
	}
L748:
	;
	v3426 = F_lappend_oid(m, v3323, v3344)
	mBase = m.M
	v3427 = m.ExcPending
	if v3427 != 0 {
		goto L33
	} else {
		goto L749
	}
L749:
	;
	v3428 = F_lappend_int(m, v3326, v3420)
	mBase = m.M
	v3429 = m.ExcPending
	if v3429 != 0 {
		goto L33
	} else {
		goto L750
	}
L750:
	;
	v3430 = F_lappend_oid(m, v3324, v3424)
	mBase = m.M
	v3431 = m.ExcPending
	if v3431 != 0 {
		goto L33
	} else {
		goto L751
	}
L751:
	;
	v3433 = v3314 + int32(1)
	if v3433 != v1222 {
		v3314 = v3433
		v3323 = v3426
		v3324 = v3430
		v3326 = v3428
		goto L735
	} else {
		goto L752
	}
L752:
	;
	goto L736
L753:
	;
	v3465 = v1223 + v3442<<(uint(int32(2))%32)
	v3466 = *(*int32)(unsafe.Add(mBase, uint32(v3465)))
	v3471 = int32(0)
	goto L755
L755:
	;
	v3495 = int32(0)
	if v3466 == v3495 {
		v3504 = v3495
		goto L757
	} else {
		goto L758
	}
L757:
	;
	if v1301 == int32(0) {
		goto L761
	} else {
		goto L762
	}
L758:
	;
	v3498 = *(*int32)(unsafe.Add(mBase, uint32(v3466)+4))
	if v3498 <= v3471 {
		v3504 = v3495
		goto L757
	} else {
		goto L759
	}
L759:
	;
	v3500 = *(*int32)(unsafe.Add(mBase, uint32(v3466)+12))
	v3504 = v3500 + v3471<<(uint(int32(2))%32)
	goto L757
L760:
	;
	v3522 = v3512 + v3471<<(uint(int32(2))%32)
	v3523 = *(*int32)(unsafe.Add(mBase, uint32(v3522)))
	v3524 = *(*int32)(unsafe.Add(mBase, uint32(v3504)))
	v3525 = F_lappend(m, v3523, v3524)
	mBase = m.M
	v3526 = m.ExcPending
	if v3526 != 0 {
		goto L33
	} else {
		goto L767
	}
L761:
	;
	v3514 = *(*int32)(unsafe.Add(mBase, uint32(v3465)))
	F_list_free(m, v3514)
	mBase = m.M
	v3516 = m.ExcPending
	if v3516 != 0 {
		goto L33
	} else {
		goto L765
	}
L762:
	;
	v3509 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+4))
	if base.B2i32(v3504 == int32(0))|base.B2i32(v3509 <= v3471) != 0 {
		goto L761
	} else {
		goto L763
	}
L763:
	;
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+12))
	if v3512 != 0 {
		goto L760
	} else {
		goto L764
	}
L764:
	;
	goto L761
L765:
	;
	v3518 = v3442 + int32(1)
	if v3518 != v1222 {
		v3442 = v3518
		goto L753
	} else {
		goto L766
	}
L766:
	;
	v3541 = v1301
	v3543 = v3426
	v3544 = v3430
	v3546 = v3428
	goto L3
L767:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3522))) = v3525
	v3471 = v3471 + int32(1)
	goto L755
L768:
	;
	v3559 = F_contain_vars_of_level(m, v3541, int32(0))
	mBase = m.M
	v3560 = m.ExcPending
	if v3560 != 0 {
		goto L33
	} else {
		goto L771
	}
L769:
	;
	v3562 = int32(0)
	goto L770
L770:
	;
	v3563 = F_addRangeTableEntryForValues(m, l0, v3541, v3543, v3546, v3544, v3562)
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		goto L33
	} else {
		goto L772
	}
L771:
	;
	v3562 = v3559
	goto L770
L772:
	;
	v3565 = int32(1)
	F_addNSItemToQuery(m, l0, v3563, v3565, v3565, v3565)
	mBase = m.M
	v3569 = m.ExcPending
	if v3569 != 0 {
		goto L33
	} else {
		goto L773
	}
L773:
	;
	v3572 = F_expandNSItemAttrs(m, l0, v3563, int32(0), int32(-1))
	mBase = m.M
	v3573 = m.ExcPending
	if v3573 != 0 {
		goto L33
	} else {
		goto L774
	}
L774:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+76)) = v3572
	v3575 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v3579 = F_transformSortClause(m, l0, v3575, v1153+int32(76), int32(0))
	mBase = m.M
	v3580 = m.ExcPending
	if v3580 != 0 {
		goto L33
	} else {
		goto L775
	}
L775:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+124)) = v3579
	v3582 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v3585 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v3586 = F_transformLimitClause(m, l0, v3582, int32(23), int32(_a_F_transformStmt_11), v3585)
	mBase = m.M
	v3587 = m.ExcPending
	if v3587 != 0 {
		goto L33
	} else {
		goto L776
	}
L776:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+128)) = v3586
	v3589 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v3592 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v3593 = F_transformLimitClause(m, l0, v3589, int32(22), int32(_a_F_transformStmt_12), v3592)
	mBase = m.M
	v3594 = m.ExcPending
	if v3594 != 0 {
		goto L33
	} else {
		goto L777
	}
L777:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+132)) = v3593
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+136)) = v3596
	v3598 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v3598 == int32(0) {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v3601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+52)) = v3601
	v3603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+56)) = v3603
	v3605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3607 = F_makeFromExpr(m, v3605, int32(0))
	mBase = m.M
	v3608 = m.ExcPending
	if v3608 != 0 {
		goto L33
	} else {
		goto L781
	}
L779:
	;
	goto L780
L780:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3617 = m.ExcPending
	if v3617 != 0 {
		goto L33
	} else {
		goto L783
	}
L781:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+60)) = v3607
	v3610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1153)+39)) = uint8(v3610)
	F_assign_query_collations(m, l0, v1153)
	mBase = m.M
	v3613 = m.ExcPending
	if v3613 != 0 {
		goto L33
	} else {
		goto L782
	}
L782:
	;
	v4789 = v1153
	goto L1
L783:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3620 = m.ExcPending
	if v3620 != 0 {
		goto L33
	} else {
		goto L784
	}
L784:
	;
	v3621 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v3621)+12))
	v3623 = *(*int32)(unsafe.Add(mBase, uint32(v3622)))
	v3624 = *(*int32)(unsafe.Add(mBase, uint32(v3623)+8))
	v3626 = v3624 - int32(1)
	if base.Ui32(v3626) <= base.Ui32(int32(3)) {
		goto L786
	} else {
		goto L787
	}
L785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v3633
	F_errmsg(m, int32(_a_F_transformStmt_55), v30)
	mBase = m.M
	v3637 = m.ExcPending
	if v3637 != 0 {
		goto L33
	} else {
		goto L789
	}
L786:
	;
	v3631 = *(*int32)(unsafe.Add(mBase, uint32(v3626<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v3633 = v3631
	goto L788
L787:
	;
	v3633 = int32(_a_F_transformStmt_21)
	goto L788
L788:
	;
	goto L785
L789:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(1719), int32(_a_F_transformStmt_29))
	mBase = m.M
	v3642 = m.ExcPending
	if v3642 != 0 {
		goto L33
	} else {
		goto L790
	}
L790:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L791:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v3650
	v3653 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v3656 = F_checkInsertTargets(m, l0, v3653, v30+int32(192))
	mBase = m.M
	v3657 = m.ExcPending
	if v3657 != 0 {
		goto L33
	} else {
		goto L792
	}
L792:
	;
	if v42 != 0 {
		goto L793
	} else {
		goto L794
	}
L793:
	;
	if v3643 != 0 {
		goto L799
	} else {
		goto L800
	}
L794:
	;
	v4019 = v3
	goto L795
L795:
	;
	v4037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v4038 = *(*int32)(unsafe.Add(mBase, uint32(v4037)+12))
	v4039 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v4039
	v4041 = *(*int32)(unsafe.Add(mBase, uint32(v30)+192))
	v4046 = v4039
	goto L880
L796:
	;
	v4005 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v4006 = *(*int32)(unsafe.Add(mBase, uint32(v30)+192))
	v4008 = F_transformInsertRow(m, l0, v3985, v4005, v3656, v4006, int32(0))
	mBase = m.M
	v4009 = m.ExcPending
	if v4009 != 0 {
		goto L33
	} else {
		goto L879
	}
L797:
	;
	v3868 = *(*int32)(unsafe.Add(mBase, uint32(v3843)+12))
	v3869 = *(*int32)(unsafe.Add(mBase, uint32(v3868)))
	if v3869 == int32(0) {
		goto L858
	} else {
		goto L859
	}
L798:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3858 = m.ExcPending
	if v3858 != 0 {
		goto L33
	} else {
		goto L852
	}
L799:
	;
	v3658 = F_make_parsestate(m, l0)
	mBase = m.M
	v3659 = m.ExcPending
	if v3659 != 0 {
		goto L33
	} else {
		goto L802
	}
L800:
	;
	goto L801
L801:
	;
	v3763 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v3763 == int32(0) {
		goto L826
	} else {
		goto L827
	}
L802:
	;
	v3660 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3658)+85)) = uint8(v3660)
	*(*int32)(unsafe.Add(mBase, uint32(v3658)+28)) = v3644
	*(*int64)(unsafe.Add(mBase, uint32(v3658)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3658)+12)) = v3645
	*(*int32)(unsafe.Add(mBase, uint32(v3658)+8)) = v3646
	v3667 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3668 = F_transformStmt(m, v3658, v3667)
	mBase = m.M
	v3669 = m.ExcPending
	if v3669 != 0 {
		goto L33
	} else {
		goto L803
	}
L803:
	;
	F_free_parsestate(m, v3658)
	mBase = m.M
	v3671 = m.ExcPending
	if v3671 != 0 {
		goto L33
	} else {
		goto L804
	}
L804:
	;
	v3672 = *(*int32)(unsafe.Add(mBase, uint32(v3668)))
	if v3672 != int32(67) {
		goto L798
	} else {
		goto L805
	}
L805:
	;
	v3675 = *(*int32)(unsafe.Add(mBase, uint32(v3668)+4))
	if v3675 != int32(1) {
		goto L798
	} else {
		goto L806
	}
L806:
	;
	v3678 = int32(0)
	v3681 = F_makeAlias(m, int32(_a_F_transformStmt_56), v3678)
	mBase = m.M
	v3682 = m.ExcPending
	if v3682 != 0 {
		goto L33
	} else {
		goto L807
	}
L807:
	;
	v3683 = int32(0)
	v3685 = F_addRangeTableEntryForSubquery(m, l0, v3668, v3681, v3683, v3683)
	mBase = m.M
	v3686 = m.ExcPending
	if v3686 != 0 {
		goto L33
	} else {
		goto L808
	}
L808:
	;
	v3688 = int32(0)
	F_addNSItemToQuery(m, l0, v3685, int32(1), v3688, v3688)
	mBase = m.M
	v3691 = m.ExcPending
	if v3691 != 0 {
		goto L33
	} else {
		goto L809
	}
L809:
	;
	v3692 = *(*int32)(unsafe.Add(mBase, uint32(v3668)+76))
	if v3692 == int32(0) {
		v3985 = v3678
		goto L796
	} else {
		goto L810
	}
L810:
	;
	v3695 = *(*int32)(unsafe.Add(mBase, uint32(v3692)+4))
	if v3695 <= int32(0) {
		v3985 = v3678
		goto L796
	} else {
		goto L811
	}
L811:
	;
	v3702 = int32(0)
	v3706 = v3678
	goto L812
L812:
	;
	v3726 = *(*int32)(unsafe.Add(mBase, uint32(v3692)+12))
	v3730 = *(*int32)(unsafe.Add(mBase, uint32(v3726+v3702<<(uint(int32(2))%32))))
	v3731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3730)+26)))
	if v3731 == int32(0) {
		goto L814
	} else {
		goto L815
	}
L813:
	;
	v3985 = v3757
	goto L796
L814:
	;
	v3734 = *(*int32)(unsafe.Add(mBase, uint32(v3730)+4))
	if v3734 == int32(0) {
		goto L818
	} else {
		goto L819
	}
L815:
	;
	v3757 = v3706
	goto L816
L816:
	;
	v3760 = v3702 + int32(1)
	v3761 = *(*int32)(unsafe.Add(mBase, uint32(v3692)+4))
	if v3760 < v3761 {
		v3702 = v3760
		v3706 = v3757
		goto L812
	} else {
		goto L825
	}
L817:
	;
	v3754 = F_lappend(m, v3706, v3753)
	mBase = m.M
	v3755 = m.ExcPending
	if v3755 != 0 {
		goto L33
	} else {
		goto L824
	}
L818:
	;
	v3747 = *(*int32)(unsafe.Add(mBase, uint32(v3685)+8))
	v3748 = F_makeVarFromTargetEntry(m, v3747, v3730)
	mBase = m.M
	v3749 = m.ExcPending
	if v3749 != 0 {
		goto L33
	} else {
		goto L823
	}
L819:
	;
	v3737 = *(*int32)(unsafe.Add(mBase, uint32(v3734)))
	if base.Ui32(int32(1)) < base.Ui32(v3737-int32(7)) {
		goto L818
	} else {
		goto L820
	}
L820:
	;
	v3742 = F_exprType(m, v3734)
	mBase = m.M
	v3743 = m.ExcPending
	if v3743 != 0 {
		goto L33
	} else {
		goto L821
	}
L821:
	;
	if v3742 != int32(705) {
		goto L818
	} else {
		goto L822
	}
L822:
	;
	v3746 = *(*int32)(unsafe.Add(mBase, uint32(v3730)+4))
	v3753 = v3746
	goto L817
L823:
	;
	v3750 = *(*int32)(unsafe.Add(mBase, uint32(v3730)+4))
	v3751 = F_exprLocation(m, v3750)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3748)+44)) = v3751
	v3753 = v3748
	goto L817
L824:
	;
	v3757 = v3754
	goto L816
L825:
	;
	goto L813
L826:
	;
	v3849 = *(*int32)(unsafe.Add(mBase, uint32(v3763)+12))
	v3850 = *(*int32)(unsafe.Add(mBase, uint32(v3849)))
	v3853 = F_transformExpressionList(m, l0, v3850, int32(27), int32(1))
	mBase = m.M
	v3854 = m.ExcPending
	if v3854 != 0 {
		goto L33
	} else {
		goto L851
	}
L827:
	;
	v3766 = *(*int32)(unsafe.Add(mBase, uint32(v3763)+4))
	if v3766 < int32(2) {
		goto L826
	} else {
		goto L828
	}
L828:
	;
	v3770 = int32(0)
	v3776 = v3770
	v3779 = v3770
	v3781 = int32(-1)
	goto L829
L829:
	;
	v3799 = *(*int32)(unsafe.Add(mBase, uint32(v3763)+12))
	v3803 = *(*int32)(unsafe.Add(mBase, uint32(v3799+v3776<<(uint(int32(2))%32))))
	v3806 = F_transformExpressionList(m, l0, v3803, int32(26), int32(1))
	mBase = m.M
	v3807 = m.ExcPending
	if v3807 != 0 {
		goto L33
	} else {
		goto L831
	}
L830:
	;
	goto L797
L831:
	;
	if v3781 < int32(0) {
		goto L834
	} else {
		goto L835
	}
L832:
	;
	v3836 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v3837 = *(*int32)(unsafe.Add(mBase, uint32(v30)+192))
	v3839 = F_transformInsertRow(m, l0, v3806, v3836, v3656, v3837, int32(1))
	mBase = m.M
	v3840 = m.ExcPending
	if v3840 != 0 {
		goto L33
	} else {
		goto L847
	}
L833:
	;
	v3834 = *(*int32)(unsafe.Add(mBase, uint32(v3806)+4))
	v3835 = v3834
	goto L832
L834:
	;
	if v3806 != 0 {
		goto L833
	} else {
		goto L837
	}
L835:
	;
	goto L836
L836:
	;
	if v3806 != 0 {
		goto L838
	} else {
		goto L839
	}
L837:
	;
	v3835 = int32(0)
	goto L832
L838:
	;
	v3811 = *(*int32)(unsafe.Add(mBase, uint32(v3806)+4))
	v3813 = v3811
	goto L840
L839:
	;
	v3813 = int32(0)
	goto L840
L840:
	;
	if v3813 == v3781 {
		v3835 = v3781
		goto L832
	} else {
		goto L841
	}
L841:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3818 = m.ExcPending
	if v3818 != 0 {
		goto L33
	} else {
		goto L842
	}
L842:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3821 = m.ExcPending
	if v3821 != 0 {
		goto L33
	} else {
		goto L843
	}
L843:
	;
	F_errmsg(m, int32(_a_F_transformStmt_28), int32(0))
	mBase = m.M
	v3825 = m.ExcPending
	if v3825 != 0 {
		goto L33
	} else {
		goto L844
	}
L844:
	;
	v3826 = F_exprLocation(m, v3806)
	mBase = m.M
	F_parser_errposition(m, l0, v3826)
	mBase = m.M
	v3828 = m.ExcPending
	if v3828 != 0 {
		goto L33
	} else {
		goto L845
	}
L845:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(871), int32(_a_F_transformStmt_57))
	mBase = m.M
	v3833 = m.ExcPending
	if v3833 != 0 {
		goto L33
	} else {
		goto L846
	}
L846:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L847:
	;
	F_assign_list_collations(m, l0, v3839)
	mBase = m.M
	v3842 = m.ExcPending
	if v3842 != 0 {
		goto L33
	} else {
		goto L848
	}
L848:
	;
	v3843 = F_lappend(m, v3779, v3839)
	mBase = m.M
	v3844 = m.ExcPending
	if v3844 != 0 {
		goto L33
	} else {
		goto L849
	}
L849:
	;
	v3846 = v3776 + int32(1)
	v3847 = *(*int32)(unsafe.Add(mBase, uint32(v3763)+4))
	if v3846 < v3847 {
		v3776 = v3846
		v3779 = v3843
		v3781 = v3835
		goto L829
	} else {
		goto L850
	}
L850:
	;
	goto L830
L851:
	;
	v3985 = v3853
	goto L796
L852:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_58), int32(0))
	mBase = m.M
	v3862 = m.ExcPending
	if v3862 != 0 {
		goto L33
	} else {
		goto L853
	}
L853:
	;
	F_errfinish(m, int32(_a_F_transformStmt_14), int32(772), int32(_a_F_transformStmt_57))
	mBase = m.M
	v3867 = m.ExcPending
	if v3867 != 0 {
		goto L33
	} else {
		goto L854
	}
L854:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L855:
	;
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3957 != 0 {
		goto L871
	} else {
		goto L872
	}
L856:
	;
	v3880 = int32(0)
	v3886 = v3880
	v3887 = v3873
	v3889 = v3880
	v3892 = v3880
	goto L862
L857:
	;
	v3878 = int32(0)
	v3934 = v3877
	v3936 = v3878
	v3939 = v3878
	goto L855
L858:
	;
	v3877 = int32(0)
	goto L857
L859:
	;
	goto L860
L860:
	;
	v3873 = int32(0)
	v3874 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+4))
	if v3873 < v3874 {
		goto L856
	} else {
		goto L861
	}
L861:
	;
	v3877 = v3873
	goto L857
L862:
	;
	v3910 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+12))
	v3914 = *(*int32)(unsafe.Add(mBase, uint32(v3910+v3886<<(uint(int32(2))%32))))
	v3915 = F_exprType(m, v3914)
	mBase = m.M
	v3916 = m.ExcPending
	if v3916 != 0 {
		goto L33
	} else {
		goto L864
	}
L863:
	;
	v3934 = v3917
	v3936 = v3921
	v3939 = v3924
	goto L855
L864:
	;
	v3917 = F_lappend_oid(m, v3887, v3915)
	mBase = m.M
	v3918 = m.ExcPending
	if v3918 != 0 {
		goto L33
	} else {
		goto L865
	}
L865:
	;
	v3919 = F_exprTypmod(m, v3914)
	mBase = m.M
	v3920 = m.ExcPending
	if v3920 != 0 {
		goto L33
	} else {
		goto L866
	}
L866:
	;
	v3921 = F_lappend_int(m, v3889, v3919)
	mBase = m.M
	v3922 = m.ExcPending
	if v3922 != 0 {
		goto L33
	} else {
		goto L867
	}
L867:
	;
	v3924 = F_lappend_oid(m, v3892, int32(0))
	mBase = m.M
	v3925 = m.ExcPending
	if v3925 != 0 {
		goto L33
	} else {
		goto L868
	}
L868:
	;
	v3927 = v3886 + int32(1)
	v3928 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+4))
	if v3927 < v3928 {
		v3886 = v3927
		v3887 = v3917
		v3889 = v3921
		v3892 = v3924
		goto L862
	} else {
		goto L869
	}
L869:
	;
	goto L863
L870:
	;
	v3966 = F_addRangeTableEntryForValues(m, l0, v3843, v3934, v3936, v3939, v3965)
	mBase = m.M
	v3967 = m.ExcPending
	if v3967 != 0 {
		goto L33
	} else {
		goto L876
	}
L871:
	;
	v3959 = *(*int32)(unsafe.Add(mBase, uint32(v3957)+4))
	if v3959 == int32(1) {
		v3965 = int32(0)
		goto L870
	} else {
		goto L874
	}
L872:
	;
	goto L873
L873:
	;
	v3963 = F_contain_vars_of_level(m, v3843, int32(0))
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		goto L33
	} else {
		goto L875
	}
L874:
	;
	goto L873
L875:
	;
	v3965 = v3963
	goto L870
L876:
	;
	v3969 = int32(0)
	F_addNSItemToQuery(m, l0, v3966, int32(1), v3969, v3969)
	mBase = m.M
	v3972 = m.ExcPending
	if v3972 != 0 {
		goto L33
	} else {
		goto L877
	}
L877:
	;
	v3973 = int32(0)
	v3976 = F_expandNSItemVars(m, l0, v3966, v3973, int32(-1), v3973)
	mBase = m.M
	v3977 = m.ExcPending
	if v3977 != 0 {
		goto L33
	} else {
		goto L878
	}
L878:
	;
	v3985 = v3976
	goto L796
L879:
	;
	v4019 = v4008
	goto L795
L880:
	;
	v4070 = int32(0)
	if v4019 == v4070 {
		v4080 = v4070
		goto L882
	} else {
		goto L883
	}
L882:
	;
	v4081 = int32(0)
	if v3656 == v4081 {
		v4090 = v4081
		goto L885
	} else {
		goto L886
	}
L883:
	;
	v4074 = *(*int32)(unsafe.Add(mBase, uint32(v4019)+4))
	if v4074 <= v4046 {
		v4080 = int32(0)
		goto L882
	} else {
		goto L884
	}
L884:
	;
	v4076 = *(*int32)(unsafe.Add(mBase, uint32(v4019)+12))
	v4080 = v4076 + v4046<<(uint(int32(2))%32)
	goto L882
L885:
	;
	if v4041 == int32(0) {
		goto L889
	} else {
		goto L890
	}
L886:
	;
	v4084 = *(*int32)(unsafe.Add(mBase, uint32(v3656)+4))
	if v4084 <= v4046 {
		v4090 = v4081
		goto L885
	} else {
		goto L887
	}
L887:
	;
	v4086 = *(*int32)(unsafe.Add(mBase, uint32(v3656)+12))
	v4090 = v4086 + v4046<<(uint(int32(2))%32)
	goto L885
L888:
	;
	v4762 = *(*int32)(unsafe.Add(mBase, uint32(v4080)))
	v4766 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4101+v4046<<(uint(int32(2))%32)))))
	v4767 = *(*int32)(unsafe.Add(mBase, uint32(v4090)))
	v4768 = *(*int32)(unsafe.Add(mBase, uint32(v4767)+4))
	v4770 = F_makeTargetEntry(m, v4762, v4766, v4768, int32(0))
	mBase = m.M
	v4771 = m.ExcPending
	if v4771 != 0 {
		goto L33
	} else {
		goto L1032
	}
L889:
	;
	v4103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v4103 == int32(0) {
		goto L894
	} else {
		goto L895
	}
L890:
	;
	v4093 = int32(0)
	v4097 = *(*int32)(unsafe.Add(mBase, uint32(v4041)+4))
	if base.B2i32(v4090 == v4093)|(base.B2i32(v4080 == v4093)|base.B2i32(v4097 <= v4046)) != 0 {
		goto L889
	} else {
		goto L891
	}
L891:
	;
	v4101 = *(*int32)(unsafe.Add(mBase, uint32(v4041)+12))
	if v4101 != 0 {
		goto L888
	} else {
		goto L892
	}
L892:
	;
	goto L889
L893:
	;
	v4743 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4743 != 0 {
		goto L1026
	} else {
		goto L1027
	}
L894:
	;
	v4106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4106 == int32(0) {
		goto L893
	} else {
		goto L897
	}
L895:
	;
	goto L896
L896:
	;
	v4109 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v4109
	v4112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v4114 = int32(1)
	F_addNSItemToQuery(m, l0, v4112, v4109, v4114, v4114)
	mBase = m.M
	v4117 = m.ExcPending
	if v4117 != 0 {
		goto L33
	} else {
		goto L898
	}
L897:
	;
	goto L896
L898:
	;
	v4118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v4118 == int32(0) {
		goto L893
	} else {
		goto L899
	}
L899:
	;
	v4121 = int32(0)
	v4123 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+4))
	if v4123 == int32(2) {
		goto L900
	} else {
		goto L901
	}
L900:
	;
	v4126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4130 = F_makeAlias(m, int32(_a_F_transformStmt_59), int32(0))
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L33
	} else {
		goto L903
	}
L901:
	;
	v4142 = v4121
	v4144 = v4109
	v4145 = v4121
	goto L902
L902:
	;
	v4146 = m.G0
	v4148 = v4146 - int32(32)
	m.G0 = v4148
	v4150 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+8))
	v4152 = v30 + int32(204)
	v4153 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4152))) = v4153
	v4156 = v30 + int32(200)
	*(*int32)(unsafe.Add(mBase, uint32(v4156))) = v4153
	v4160 = v30 + int32(196)
	*(*int32)(unsafe.Add(mBase, uint32(v4160))) = v4153
	v4165 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+4))
	if base.B2i32(v4150 == v4153)&base.B2i32(v4165 == int32(2)) == v4153 {
		goto L911
	} else {
		goto L912
	}
L903:
	;
	v4132 = int32(0)
	v4134 = F_addRangeTableEntryForRelation(m, l0, v4126, int32(3), v4130, v4132, v4132)
	mBase = m.M
	v4135 = m.ExcPending
	if v4135 != 0 {
		goto L33
	} else {
		goto L904
	}
L904:
	;
	v4136 = *(*int32)(unsafe.Add(mBase, uint32(v4134)+8))
	v4137 = *(*int32)(unsafe.Add(mBase, uint32(v4134)+4))
	v4138 = int32(99)
	*(*uint8)(unsafe.Add(mBase, uint32(v4137)+21)) = uint8(v4138)
	v4140 = F_BuildOnConflictExcludedTargetlist(m, v4126, v4136)
	mBase = m.M
	v4141 = m.ExcPending
	if v4141 != 0 {
		goto L33
	} else {
		goto L905
	}
L905:
	;
	v4142 = v4134
	v4144 = v4136
	v4145 = v4140
	goto L902
L906:
	;
	v4672 = int32(0)
	v4674 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+4))
	if v4674 == int32(2) {
		goto L1018
	} else {
		goto L1019
	}
L907:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4656 = m.ExcPending
	if v4656 != 0 {
		goto L33
	} else {
		goto L1013
	}
L908:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4637 = m.ExcPending
	if v4637 != 0 {
		goto L33
	} else {
		goto L1008
	}
L909:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4612 = m.ExcPending
	if v4612 != 0 {
		goto L33
	} else {
		goto L1003
	}
L910:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4593 = m.ExcPending
	if v4593 != 0 {
		goto L33
	} else {
		goto L998
	}
L911:
	;
	v4171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4172 = *(*int32)(unsafe.Add(mBase, uint32(v4171)+56))
	goto L914
L912:
	;
	goto L913
L913:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4570 = m.ExcPending
	if v4570 != 0 {
		goto L33
	} else {
		goto L992
	}
L914:
	;
	if base.Ui32(v4172) < base.Ui32(int32(_a_F_transformStmt_60)) {
		goto L910
	} else {
		goto L915
	}
L915:
	;
	v4175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4176 = *(*int32)(unsafe.Add(mBase, uint32(v4175)+180))
	if v4176 == int32(0) {
		goto L916
	} else {
		goto L917
	}
L916:
	;
	if v4150 == int32(0) {
		goto L920
	} else {
		goto L921
	}
L917:
	;
	v4179 = *(*int32)(unsafe.Add(mBase, uint32(v4175)+48))
	v4180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4179)+119)))
	switch v4180 - int32(109) {
	case 0, 5:
		goto L918
	default:
		goto L916
	}
L918:
	;
	v4183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4176)+104)))
	if v4183 == int32(1) {
		goto L909
	} else {
		goto L919
	}
L919:
	;
	goto L916
L920:
	;
	m.G0 = v4148 + int32(32)
	goto L906
L921:
	;
	v4188 = *(*int32)(unsafe.Add(mBase, uint32(v4150)+4))
	if v4188 != 0 {
		goto L922
	} else {
		goto L923
	}
L922:
	;
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+4))
	if int32(0) < v4189 {
		goto L925
	} else {
		goto L926
	}
L923:
	;
	goto L924
L924:
	;
	v4334 = *(*int32)(unsafe.Add(mBase, uint32(v4150)+8))
	if v4334 != 0 {
		goto L950
	} else {
		goto L951
	}
L925:
	;
	v4206 = v3
	v4208 = v3
	goto L928
L926:
	;
	v4295 = v3
	goto L927
L927:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4152))) = v4295
	goto L924
L928:
	;
	v4219 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+12))
	v4223 = *(*int32)(unsafe.Add(mBase, uint32(v4219+v4206<<(uint(int32(2))%32))))
	v4225 = F_palloc0(m, int32(16))
	mBase = m.M
	v4226 = m.ExcPending
	if v4226 != 0 {
		goto L33
	} else {
		goto L930
	}
L929:
	;
	v4295 = v4273
	goto L927
L930:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4225))) = int32(60)
	v4229 = *(*int32)(unsafe.Add(mBase, uint32(v4223)+28))
	if v4229 != 0 {
		goto L908
	} else {
		goto L931
	}
L931:
	;
	v4230 = *(*int32)(unsafe.Add(mBase, uint32(v4223)+32))
	if v4230 != 0 {
		goto L907
	} else {
		goto L932
	}
L932:
	;
	v4231 = *(*int32)(unsafe.Add(mBase, uint32(v4223)+8))
	if v4231 == int32(0) {
		goto L933
	} else {
		goto L934
	}
L933:
	;
	v4235 = F_palloc0(m, int32(12))
	mBase = m.M
	v4236 = m.ExcPending
	if v4236 != 0 {
		goto L33
	} else {
		goto L936
	}
L934:
	;
	v4252 = v4231
	goto L935
L935:
	;
	v4255 = F_transformExpr(m, l0, v4252, int32(32))
	mBase = m.M
	v4256 = m.ExcPending
	if v4256 != 0 {
		goto L33
	} else {
		goto L939
	}
L936:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4235))) = int32(69)
	v4239 = *(*int32)(unsafe.Add(mBase, uint32(v4223)+4))
	v4240 = F_makeString(m, v4239)
	mBase = m.M
	v4241 = m.ExcPending
	if v4241 != 0 {
		goto L33
	} else {
		goto L937
	}
L937:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4148)+12)) = v4240
	*(*int32)(unsafe.Add(mBase, uint32(v4148)+28)) = v4240
	v4247 = F_list_make1_impl(m, int32(1), v4148+int32(12))
	mBase = m.M
	v4248 = m.ExcPending
	if v4248 != 0 {
		goto L33
	} else {
		goto L938
	}
L938:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4235)+4)) = v4247
	v4250 = *(*int32)(unsafe.Add(mBase, uint32(v4150)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4235)+8)) = v4250
	v4252 = v4235
	goto L935
L939:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4225)+4)) = v4255
	v4258 = *(*int32)(unsafe.Add(mBase, uint32(v4223)+16))
	if v4258 != 0 {
		goto L940
	} else {
		goto L941
	}
L940:
	;
	v4259 = F_exprLocation(m, v4255)
	mBase = m.M
	v4260 = F_LookupCollation(m, l0, v4258, v4259)
	mBase = m.M
	v4261 = m.ExcPending
	if v4261 != 0 {
		goto L33
	} else {
		goto L943
	}
L941:
	;
	v4263 = int32(0)
	goto L942
L942:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4225)+8)) = v4263
	v4265 = *(*int32)(unsafe.Add(mBase, uint32(v4223)+20))
	if v4265 != 0 {
		goto L944
	} else {
		goto L945
	}
L943:
	;
	v4263 = v4260
	goto L942
L944:
	;
	v4268 = F_get_opclass_oid(m, int32(403), v4265, int32(0))
	mBase = m.M
	v4269 = m.ExcPending
	if v4269 != 0 {
		goto L33
	} else {
		goto L947
	}
L945:
	;
	v4271 = int32(0)
	goto L946
L946:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4225)+12)) = v4271
	v4273 = F_lappend(m, v4208, v4225)
	mBase = m.M
	v4274 = m.ExcPending
	if v4274 != 0 {
		goto L33
	} else {
		goto L948
	}
L947:
	;
	v4271 = v4268
	goto L946
L948:
	;
	v4276 = v4206 + int32(1)
	v4277 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+4))
	if v4276 < v4277 {
		v4206 = v4276
		v4208 = v4273
		goto L928
	} else {
		goto L949
	}
L949:
	;
	goto L929
L950:
	;
	v4336 = F_transformExpr(m, l0, v4334, int32(33))
	mBase = m.M
	v4337 = m.ExcPending
	if v4337 != 0 {
		goto L33
	} else {
		goto L953
	}
L951:
	;
	goto L952
L952:
	;
	v4339 = *(*int32)(unsafe.Add(mBase, uint32(v4150)+12))
	if v4339 == int32(0) {
		goto L920
	} else {
		goto L954
	}
L953:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4156))) = v4336
	goto L952
L954:
	;
	v4342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v4343 = *(*int32)(unsafe.Add(mBase, uint32(v4342)+12))
	v4344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4345 = *(*int32)(unsafe.Add(mBase, uint32(v4344)+56))
	v4346 = int32(0)
	v4347 = m.G0
	v4349 = v4347 - int32(160)
	m.G0 = v4349
	*(*int32)(unsafe.Add(mBase, uint32(v4160))) = v4346
	v4355 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v4356 = m.ExcPending
	if v4356 != 0 {
		goto L33
	} else {
		goto L956
	}
L955:
	;
	v4529 = *(*int64)(unsafe.Add(mBase, uint32(v4343)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4343)+16)) = v4529 | int64(2)
	v4533 = *(*int32)(unsafe.Add(mBase, uint32(v4343)+28))
	v4534 = F_bms_add_members(m, v4533, v4466)
	mBase = m.M
	v4535 = m.ExcPending
	if v4535 != 0 {
		goto L33
	} else {
		goto L991
	}
L956:
	;
	v4358 = v4349 + int32(16)
	F_ScanKeyInit(m, v4358, int32(9), int32(3), int32(184), v4345)
	mBase = m.M
	v4363 = m.ExcPending
	if v4363 != 0 {
		goto L33
	} else {
		goto L957
	}
L957:
	;
	F_ScanKeyInit(m, v4349-int32(-64), int32(10), int32(3), int32(184), int32(0))
	mBase = m.M
	v4371 = m.ExcPending
	if v4371 != 0 {
		goto L33
	} else {
		goto L958
	}
L958:
	;
	F_ScanKeyInit(m, v4349+int32(112), int32(2), int32(3), int32(62), v4339)
	mBase = m.M
	v4378 = m.ExcPending
	if v4378 != 0 {
		goto L33
	} else {
		goto L959
	}
L959:
	;
	v4383 = F_systable_beginscan(m, v4355, int32(2665), int32(1), int32(0), int32(3), v4358)
	mBase = m.M
	v4384 = m.ExcPending
	if v4384 != 0 {
		goto L33
	} else {
		goto L963
	}
L960:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4513 = m.ExcPending
	if v4513 != 0 {
		goto L33
	} else {
		goto L986
	}
L961:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4500 = m.ExcPending
	if v4500 != 0 {
		goto L33
	} else {
		goto L983
	}
L962:
	;
	F_systable_endscan(m, v4383)
	mBase = m.M
	v4486 = m.ExcPending
	if v4486 != 0 {
		goto L33
	} else {
		goto L980
	}
L963:
	;
	v4385 = F_systable_getnext(m, v4383)
	mBase = m.M
	v4386 = m.ExcPending
	if v4386 != 0 {
		goto L33
	} else {
		goto L964
	}
L964:
	;
	if v4385 == int32(0) {
		v4466 = v4346
		goto L962
	} else {
		goto L965
	}
L965:
	;
	v4389 = *(*int32)(unsafe.Add(mBase, uint32(v4385)+16))
	v4390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4389)+22)))
	v4392 = *(*int32)(unsafe.Add(mBase, uint32(v4389+v4390)))
	*(*int32)(unsafe.Add(mBase, uint32(v4160))) = v4392
	v4395 = *(*int32)(unsafe.Add(mBase, uint32(v4355)+52))
	v4398 = F_heap_getattr_3(m, v4385, v4395, v4349+int32(15))
	mBase = m.M
	v4399 = m.ExcPending
	if v4399 != 0 {
		goto L33
	} else {
		goto L966
	}
L966:
	;
	v4400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4349)+15)))
	if v4400 != 0 {
		v4466 = int32(0)
		goto L962
	} else {
		goto L967
	}
L967:
	;
	v4401 = F_pg_detoast_datum(m, v4398)
	mBase = m.M
	v4402 = m.ExcPending
	if v4402 != 0 {
		goto L33
	} else {
		goto L968
	}
L968:
	;
	v4403 = *(*int32)(unsafe.Add(mBase, uint32(v4401)+4))
	if v4403 != int32(1) {
		goto L961
	} else {
		goto L969
	}
L969:
	;
	v4406 = *(*int32)(unsafe.Add(mBase, uint32(v4401)+16))
	if v4406 < int32(0) {
		goto L961
	} else {
		goto L970
	}
L970:
	;
	v4409 = *(*int32)(unsafe.Add(mBase, uint32(v4401)+8))
	if v4409 != 0 {
		goto L961
	} else {
		goto L971
	}
L971:
	;
	v4410 = *(*int32)(unsafe.Add(mBase, uint32(v4401)+12))
	if v4410 != int32(21) {
		goto L961
	} else {
		goto L972
	}
L972:
	;
	v4413 = int32(0)
	if v4406 == v4413 {
		goto L973
	} else {
		goto L974
	}
L973:
	;
	v4466 = int32(0)
	goto L962
L974:
	;
	goto L975
L975:
	;
	v4428 = int32(0)
	v4433 = v4413
	goto L976
L976:
	;
	v4450 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4401+int32(24)+v4433<<(uint(int32(1))%32)))))
	v4453 = F_bms_add_member(m, v4428, v4450+int32(7))
	mBase = m.M
	v4454 = m.ExcPending
	if v4454 != 0 {
		goto L33
	} else {
		goto L978
	}
L977:
	;
	v4466 = v4453
	goto L962
L978:
	;
	v4456 = v4433 + int32(1)
	if v4456 != v4406 {
		v4428 = v4453
		v4433 = v4456
		goto L976
	} else {
		goto L979
	}
L979:
	;
	goto L977
L980:
	;
	v4487 = *(*int32)(unsafe.Add(mBase, uint32(v4160)))
	if v4487 == int32(0) {
		goto L960
	} else {
		goto L981
	}
L981:
	;
	F_relation_close(m, v4355, int32(1))
	mBase = m.M
	v4492 = m.ExcPending
	if v4492 != 0 {
		goto L33
	} else {
		goto L982
	}
L982:
	;
	m.G0 = v4349 + int32(160)
	goto L955
L983:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_61), int32(0))
	mBase = m.M
	v4504 = m.ExcPending
	if v4504 != 0 {
		goto L33
	} else {
		goto L984
	}
L984:
	;
	F_errfinish(m, int32(_a_F_transformStmt_62), int32(1309), int32(_a_F_transformStmt_63))
	mBase = m.M
	v4509 = m.ExcPending
	if v4509 != 0 {
		goto L33
	} else {
		goto L985
	}
L985:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L986:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v4516 = m.ExcPending
	if v4516 != 0 {
		goto L33
	} else {
		goto L987
	}
L987:
	;
	v4517 = F_get_rel_name(m, v4345)
	mBase = m.M
	v4518 = m.ExcPending
	if v4518 != 0 {
		goto L33
	} else {
		goto L988
	}
L988:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4349)+4)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v4349))) = v4339
	F_errmsg(m, int32(_a_F_transformStmt_64), v4349)
	mBase = m.M
	v4523 = m.ExcPending
	if v4523 != 0 {
		goto L33
	} else {
		goto L989
	}
L989:
	;
	F_errfinish(m, int32(_a_F_transformStmt_62), int32(1328), int32(_a_F_transformStmt_63))
	mBase = m.M
	v4528 = m.ExcPending
	if v4528 != 0 {
		goto L33
	} else {
		goto L990
	}
L990:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L991:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4343)+28)) = v4534
	goto L920
L992:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4573 = m.ExcPending
	if v4573 != 0 {
		goto L33
	} else {
		goto L993
	}
L993:
	;
	F_errmsg(m, int32(_a_F_transformStmt_65), int32(0))
	mBase = m.M
	v4577 = m.ExcPending
	if v4577 != 0 {
		goto L33
	} else {
		goto L994
	}
L994:
	;
	F_errhint(m, int32(_a_F_transformStmt_66), int32(0))
	mBase = m.M
	v4581 = m.ExcPending
	if v4581 != 0 {
		goto L33
	} else {
		goto L995
	}
L995:
	;
	v4582 = F_exprLocation(m, v4118)
	mBase = m.M
	F_parser_errposition(m, l0, v4582)
	mBase = m.M
	v4584 = m.ExcPending
	if v4584 != 0 {
		goto L33
	} else {
		goto L996
	}
L996:
	;
	F_errfinish(m, int32(_a_F_transformStmt_67), int32(3314), int32(_a_F_transformStmt_68))
	mBase = m.M
	v4589 = m.ExcPending
	if v4589 != 0 {
		goto L33
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v4596 = m.ExcPending
	if v4596 != 0 {
		goto L33
	} else {
		goto L999
	}
L999:
	;
	F_errmsg(m, int32(_a_F_transformStmt_69), int32(0))
	mBase = m.M
	v4600 = m.ExcPending
	if v4600 != 0 {
		goto L33
	} else {
		goto L1000
	}
L1000:
	;
	v4601 = F_exprLocation(m, v4118)
	mBase = m.M
	F_parser_errposition(m, l0, v4601)
	mBase = m.M
	v4603 = m.ExcPending
	if v4603 != 0 {
		goto L33
	} else {
		goto L1001
	}
L1001:
	;
	F_errfinish(m, int32(_a_F_transformStmt_67), int32(3325), int32(_a_F_transformStmt_68))
	mBase = m.M
	v4608 = m.ExcPending
	if v4608 != 0 {
		goto L33
	} else {
		goto L1002
	}
L1002:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1003:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4615 = m.ExcPending
	if v4615 != 0 {
		goto L33
	} else {
		goto L1004
	}
L1004:
	;
	v4616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4617 = *(*int32)(unsafe.Add(mBase, uint32(v4616)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4148)+16)) = v4617 + int32(4)
	F_errmsg(m, int32(_a_F_transformStmt_70), v4148+int32(16))
	mBase = m.M
	v4625 = m.ExcPending
	if v4625 != 0 {
		goto L33
	} else {
		goto L1005
	}
L1005:
	;
	v4626 = F_exprLocation(m, v4118)
	mBase = m.M
	F_parser_errposition(m, l0, v4626)
	mBase = m.M
	v4628 = m.ExcPending
	if v4628 != 0 {
		goto L33
	} else {
		goto L1006
	}
L1006:
	;
	F_errfinish(m, int32(_a_F_transformStmt_67), int32(3334), int32(_a_F_transformStmt_68))
	mBase = m.M
	v4633 = m.ExcPending
	if v4633 != 0 {
		goto L33
	} else {
		goto L1007
	}
L1007:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1008:
	;
	F_errcode(m, int32(_a_F_transformStmt_71))
	mBase = m.M
	v4640 = m.ExcPending
	if v4640 != 0 {
		goto L33
	} else {
		goto L1009
	}
L1009:
	;
	F_errmsg(m, int32(_a_F_transformStmt_72), int32(0))
	mBase = m.M
	v4644 = m.ExcPending
	if v4644 != 0 {
		goto L33
	} else {
		goto L1010
	}
L1010:
	;
	v4645 = F_exprLocation(m, v4150)
	mBase = m.M
	F_parser_errposition(m, l0, v4645)
	mBase = m.M
	v4647 = m.ExcPending
	if v4647 != 0 {
		goto L33
	} else {
		goto L1011
	}
L1011:
	;
	F_errfinish(m, int32(_a_F_transformStmt_67), int32(3228), int32(_a_F_transformStmt_73))
	mBase = m.M
	v4652 = m.ExcPending
	if v4652 != 0 {
		goto L33
	} else {
		goto L1012
	}
L1012:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1013:
	;
	F_errcode(m, int32(_a_F_transformStmt_71))
	mBase = m.M
	v4659 = m.ExcPending
	if v4659 != 0 {
		goto L33
	} else {
		goto L1014
	}
L1014:
	;
	F_errmsg(m, int32(_a_F_transformStmt_74), int32(0))
	mBase = m.M
	v4663 = m.ExcPending
	if v4663 != 0 {
		goto L33
	} else {
		goto L1015
	}
L1015:
	;
	v4664 = F_exprLocation(m, v4150)
	mBase = m.M
	F_parser_errposition(m, l0, v4664)
	mBase = m.M
	v4666 = m.ExcPending
	if v4666 != 0 {
		goto L33
	} else {
		goto L1016
	}
L1016:
	;
	F_errfinish(m, int32(_a_F_transformStmt_67), int32(3234), int32(_a_F_transformStmt_73))
	mBase = m.M
	v4671 = m.ExcPending
	if v4671 != 0 {
		goto L33
	} else {
		goto L1017
	}
L1017:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1018:
	;
	v4677 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)) = uint8(v4677)
	v4680 = int32(1)
	F_addNSItemToQuery(m, l0, v4142, v4677, v4680, v4680)
	mBase = m.M
	v4683 = m.ExcPending
	if v4683 != 0 {
		goto L33
	} else {
		goto L1021
	}
L1019:
	;
	v4696 = v4672
	v4697 = v4672
	goto L1020
L1020:
	;
	v4699 = F_palloc0(m, int32(36))
	mBase = m.M
	v4700 = m.ExcPending
	if v4700 != 0 {
		goto L33
	} else {
		goto L1025
	}
L1021:
	;
	v4684 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+12))
	v4685 = F_transformUpdateTargetList(m, l0, v4684)
	mBase = m.M
	v4686 = m.ExcPending
	if v4686 != 0 {
		goto L33
	} else {
		goto L1022
	}
L1022:
	;
	v4687 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+16))
	v4690 = F_transformWhereClause(m, l0, v4687, int32(6), int32(_a_F_transformStmt_0))
	mBase = m.M
	v4691 = m.ExcPending
	if v4691 != 0 {
		goto L33
	} else {
		goto L1023
	}
L1023:
	;
	v4692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4693 = F_list_delete_last(m, v4692)
	mBase = m.M
	v4694 = m.ExcPending
	if v4694 != 0 {
		goto L33
	} else {
		goto L1024
	}
L1024:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v4693
	v4696 = v4685
	v4697 = v4690
	goto L1020
L1025:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4699))) = int32(66)
	v4703 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4699)+4)) = v4703
	v4705 = *(*int32)(unsafe.Add(mBase, uint32(v30)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v4699)+8)) = v4705
	v4707 = *(*int32)(unsafe.Add(mBase, uint32(v30)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v4699)+12)) = v4707
	v4709 = *(*int32)(unsafe.Add(mBase, uint32(v30)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v4699)+32)) = v4145
	*(*int32)(unsafe.Add(mBase, uint32(v4699)+28)) = v4144
	*(*int32)(unsafe.Add(mBase, uint32(v4699)+24)) = v4697
	*(*int32)(unsafe.Add(mBase, uint32(v4699)+20)) = v4696
	*(*int32)(unsafe.Add(mBase, uint32(v4699)+16)) = v4709
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v4699
	goto L893
L1026:
	;
	F_transformReturningClause(m, l0, v36, v4743, int32(24))
	mBase = m.M
	v4746 = m.ExcPending
	if v4746 != 0 {
		goto L33
	} else {
		goto L1029
	}
L1027:
	;
	goto L1028
L1028:
	;
	v4747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v4747
	v4749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+56)) = v4749
	v4751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4753 = F_makeFromExpr(m, v4751, int32(0))
	mBase = m.M
	v4754 = m.ExcPending
	if v4754 != 0 {
		goto L33
	} else {
		goto L1030
	}
L1029:
	;
	goto L1028
L1030:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+60)) = v4753
	v4756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+38)) = uint8(v4756)
	v4758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+39)) = uint8(v4758)
	F_assign_query_collations(m, l0, v36)
	mBase = m.M
	v4761 = m.ExcPending
	if v4761 != 0 {
		goto L33
	} else {
		goto L1031
	}
L1031:
	;
	v4789 = v36
	goto L1
L1032:
	;
	v4772 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v4773 = F_lappend(m, v4772, v4770)
	mBase = m.M
	v4774 = m.ExcPending
	if v4774 != 0 {
		goto L33
	} else {
		goto L1033
	}
L1033:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v4773
	v4776 = *(*int32)(unsafe.Add(mBase, uint32(v4038)+32))
	v4779 = F_bms_add_member(m, v4776, v4766+int32(7))
	mBase = m.M
	v4780 = m.ExcPending
	if v4780 != 0 {
		goto L33
	} else {
		goto L1034
	}
L1034:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4038)+32)) = v4779
	v4046 = v4046 + int32(1)
	goto L880
}
