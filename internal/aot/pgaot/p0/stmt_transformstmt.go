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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
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
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v291 int64
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int64
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int64
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v398 int64
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
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
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
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
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v730 int32
	_ = v730
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v822 int32
	_ = v822
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v913 int32
	_ = v913
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v973 int32
	_ = v973
	var v980 int32
	_ = v980
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
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
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1065 int32
	_ = v1065
	var v1085 int32
	_ = v1085
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1209 int32
	_ = v1209
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1337 int32
	_ = v1337
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int64
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1567 int32
	_ = v1567
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
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
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1753 int32
	_ = v1753
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1789 int32
	_ = v1789
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1817 int32
	_ = v1817
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1958 int32
	_ = v1958
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2044 int32
	_ = v2044
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2139 int32
	_ = v2139
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2190 int32
	_ = v2190
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2295 int32
	_ = v2295
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2319 int32
	_ = v2319
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2398 int32
	_ = v2398
	var v2400 int64
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2405 int32
	_ = v2405
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2411 int32
	_ = v2411
	var v2416 int64
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2440 int32
	_ = v2440
	var v2446 int32
	_ = v2446
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2454 int32
	_ = v2454
	var v2475 int32
	_ = v2475
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2502 int32
	_ = v2502
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2522 int32
	_ = v2522
	var v2550 int32
	_ = v2550
	var v2574 int32
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2590 int32
	_ = v2590
	var v2593 int32
	_ = v2593
	var v2597 int32
	_ = v2597
	var v2602 int32
	_ = v2602
	var v2606 int32
	_ = v2606
	var v2609 int32
	_ = v2609
	var v2613 int32
	_ = v2613
	var v2618 int32
	_ = v2618
	var v2622 int32
	_ = v2622
	var v2625 int32
	_ = v2625
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2637 int32
	_ = v2637
	var v2641 int32
	_ = v2641
	var v2645 int32
	_ = v2645
	var v2650 int32
	_ = v2650
	var v2654 int32
	_ = v2654
	var v2657 int32
	_ = v2657
	var v2666 int32
	_ = v2666
	var v2671 int32
	_ = v2671
	var v2675 int32
	_ = v2675
	var v2678 int32
	_ = v2678
	var v2687 int32
	_ = v2687
	var v2692 int32
	_ = v2692
	var v2696 int32
	_ = v2696
	var v2700 int32
	_ = v2700
	var v2705 int32
	_ = v2705
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2716 int32
	_ = v2716
	var v2721 int32
	_ = v2721
	var v2725 int32
	_ = v2725
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2734 int32
	_ = v2734
	var v2739 int32
	_ = v2739
	var v2741 int32
	_ = v2741
	var v2747 int32
	_ = v2747
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2756 int32
	_ = v2756
	var v2760 int32
	_ = v2760
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2769 int32
	_ = v2769
	var v2774 int32
	_ = v2774
	var v2776 int32
	_ = v2776
	var v2782 int32
	_ = v2782
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2791 int32
	_ = v2791
	var v2795 int32
	_ = v2795
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2809 int32
	_ = v2809
	var v2811 int32
	_ = v2811
	var v2817 int32
	_ = v2817
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2826 int32
	_ = v2826
	var v2830 int32
	_ = v2830
	var v2833 int32
	_ = v2833
	var v2837 int32
	_ = v2837
	var v2842 int32
	_ = v2842
	var v2846 int32
	_ = v2846
	var v2849 int32
	_ = v2849
	var v2853 int32
	_ = v2853
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2869 int32
	_ = v2869
	var v2873 int32
	_ = v2873
	var v2876 int32
	_ = v2876
	var v2880 int32
	_ = v2880
	var v2885 int32
	_ = v2885
	var v2889 int32
	_ = v2889
	var v2892 int32
	_ = v2892
	var v2896 int32
	_ = v2896
	var v2901 int32
	_ = v2901
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2912 int32
	_ = v2912
	var v2917 int32
	_ = v2917
	var v2921 int32
	_ = v2921
	var v2927 int32
	_ = v2927
	var v2932 int32
	_ = v2932
	var v2947 int32
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2952 int32
	_ = v2952
	var v2963 int32
	_ = v2963
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2975 int32
	_ = v2975
	var v2981 int32
	_ = v2981
	var v3005 int32
	_ = v3005
	var v3008 int32
	_ = v3008
	var v3009 int32
	_ = v3009
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3018 int32
	_ = v3018
	var v3022 int32
	_ = v3022
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3057 int32
	_ = v3057
	var v3059 int32
	_ = v3059
	var v3068 int32
	_ = v3068
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3097 int32
	_ = v3097
	var v3121 int32
	_ = v3121
	var v3124 int32
	_ = v3124
	var v3126 int32
	_ = v3126
	var v3130 int32
	_ = v3130
	var v3135 int32
	_ = v3135
	var v3138 int32
	_ = v3138
	var v3140 int32
	_ = v3140
	var v3142 int32
	_ = v3142
	var v3144 int32
	_ = v3144
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
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3170 int32
	_ = v3170
	var v3172 int32
	_ = v3172
	var v3183 int32
	_ = v3183
	var v3185 int32
	_ = v3185
	var v3186 int32
	_ = v3186
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3195 int32
	_ = v3195
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3201 int32
	_ = v3201
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3208 int32
	_ = v3208
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3213 int32
	_ = v3213
	var v3215 int32
	_ = v3215
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3224 int32
	_ = v3224
	var v3227 int32
	_ = v3227
	var v3229 int32
	_ = v3229
	var v3231 int32
	_ = v3231
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3239 int32
	_ = v3239
	var v3243 int32
	_ = v3243
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3250 int32
	_ = v3250
	var v3252 int32
	_ = v3252
	var v3257 int32
	_ = v3257
	var v3259 int32
	_ = v3259
	var v3263 int32
	_ = v3263
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3277 int32
	_ = v3277
	var v3279 int32
	_ = v3279
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
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3301 int32
	_ = v3301
	var v3304 int32
	_ = v3304
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3311 int32
	_ = v3311
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3318 int32
	_ = v3318
	var v3325 int32
	_ = v3325
	var v3328 int32
	_ = v3328
	var v3349 int32
	_ = v3349
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3357 int32
	_ = v3357
	var v3360 int32
	_ = v3360
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3383 int32
	_ = v3383
	var v3384 int32
	_ = v3384
	var v3386 int32
	_ = v3386
	var v3389 int32
	_ = v3389
	var v3393 int32
	_ = v3393
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3403 int32
	_ = v3403
	var v3422 int32
	_ = v3422
	var v3426 int32
	_ = v3426
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3434 int32
	_ = v3434
	var v3436 int32
	_ = v3436
	var v3441 int32
	_ = v3441
	var v3444 int32
	_ = v3444
	var v3448 int32
	_ = v3448
	var v3449 int32
	_ = v3449
	var v3451 int32
	_ = v3451
	var v3456 int32
	_ = v3456
	var v3457 int32
	_ = v3457
	var v3458 int32
	_ = v3458
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3476 int32
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3481 int32
	_ = v3481
	var v3485 int32
	_ = v3485
	var v3490 int32
	_ = v3490
	var v3491 int32
	_ = v3491
	var v3492 int32
	_ = v3492
	var v3496 int32
	_ = v3496
	var v3497 int32
	_ = v3497
	var v3500 int32
	_ = v3500
	var v3501 int32
	_ = v3501
	var v3503 int32
	_ = v3503
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3513 int32
	_ = v3513
	var v3514 int32
	_ = v3514
	var v3533 int32
	_ = v3533
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3543 int32
	_ = v3543
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3547 int32
	_ = v3547
	var v3548 int32
	_ = v3548
	var v3550 int32
	_ = v3550
	var v3551 int32
	_ = v3551
	var v3557 int32
	_ = v3557
	var v3560 int32
	_ = v3560
	var v3561 int32
	_ = v3561
	var v3580 int32
	_ = v3580
	var v3582 int32
	_ = v3582
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3592 int32
	_ = v3592
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3599 int32
	_ = v3599
	var v3600 int32
	_ = v3600
	var v3607 int32
	_ = v3607
	var v3628 int32
	_ = v3628
	var v3629 int32
	_ = v3629
	var v3631 int32
	_ = v3631
	var v3632 int32
	_ = v3632
	var v3641 int32
	_ = v3641
	var v3660 int32
	_ = v3660
	var v3661 int32
	_ = v3661
	var v3662 int32
	_ = v3662
	var v3664 int32
	_ = v3664
	var v3669 int32
	_ = v3669
	var v3693 int32
	_ = v3693
	var v3697 int32
	_ = v3697
	var v3699 int32
	_ = v3699
	var v3703 int32
	_ = v3703
	var v3704 int32
	_ = v3704
	var v3707 int32
	_ = v3707
	var v3709 int32
	_ = v3709
	var v3713 int32
	_ = v3713
	var v3716 int32
	_ = v3716
	var v3720 int32
	_ = v3720
	var v3724 int32
	_ = v3724
	var v3726 int32
	_ = v3726
	var v3729 int32
	_ = v3729
	var v3732 int32
	_ = v3732
	var v3734 int32
	_ = v3734
	var v3736 int32
	_ = v3736
	var v3739 int32
	_ = v3739
	var v3740 int32
	_ = v3740
	var v3743 int32
	_ = v3743
	var v3746 int32
	_ = v3746
	var v3750 int32
	_ = v3750
	var v3753 int32
	_ = v3753
	var v3757 int32
	_ = v3757
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3761 int32
	_ = v3761
	var v3766 int32
	_ = v3766
	var v3767 int32
	_ = v3767
	var v3776 int32
	_ = v3776
	var v3780 int32
	_ = v3780
	var v3781 int32
	_ = v3781
	var v3782 int32
	_ = v3782
	var v3784 int32
	_ = v3784
	var v3785 int32
	_ = v3785
	var v3786 int32
	_ = v3786
	var v3787 int32
	_ = v3787
	var v3788 int32
	_ = v3788
	var v3790 int32
	_ = v3790
	var v3791 int32
	_ = v3791
	var v3792 int32
	_ = v3792
	var v3794 int32
	_ = v3794
	var v3795 int32
	_ = v3795
	var v3796 int32
	_ = v3796
	var v3798 int32
	_ = v3798
	var v3800 int32
	_ = v3800
	var v3802 int32
	_ = v3802
	var v3803 int32
	_ = v3803
	var v3806 int32
	_ = v3806
	var v3810 int32
	_ = v3810
	var v3813 int32
	_ = v3813
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3824 int32
	_ = v3824
	var v3825 int32
	_ = v3825
	var v3828 int32
	_ = v3828
	var v3829 int32
	_ = v3829
	var v3832 int32
	_ = v3832
	var v3837 int32
	_ = v3837
	var v3838 int32
	_ = v3838
	var v3852 int32
	_ = v3852
	var v3857 int32
	_ = v3857
	var v3868 int32
	_ = v3868
	var v3872 int32
	_ = v3872
	var v3874 int32
	_ = v3874
	var v3875 int32
	_ = v3875
	var v3878 int32
	_ = v3878
	var v3879 int32
	_ = v3879
	var v3880 int32
	_ = v3880
	var v3881 int32
	_ = v3881
	var v3885 int32
	_ = v3885
	var v3886 int32
	_ = v3886
	var v3889 int32
	_ = v3889
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3900 int32
	_ = v3900
	var v3902 int32
	_ = v3902
	var v3905 int32
	_ = v3905
	var v3906 int32
	_ = v3906
	var v3908 int32
	_ = v3908
	var v3909 int32
	_ = v3909
	var v3910 int32
	_ = v3910
	var v3911 int32
	_ = v3911
	var v3913 int32
	_ = v3913
	var v3915 int32
	_ = v3915
	var v3918 int32
	_ = v3918
	var v3919 int32
	_ = v3919
	var v3921 int32
	_ = v3921
	var v3923 int32
	_ = v3923
	var v3924 int32
	_ = v3924
	var v3926 int32
	_ = v3926
	var v3927 int32
	_ = v3927
	var v3945 int32
	_ = v3945
	var v3984 int32
	_ = v3984
	var v3986 int32
	_ = v3986
	var v3987 int32
	_ = v3987
	var v3989 int32
	_ = v3989
	var v3992 int32
	_ = v3992
	var v3993 int32
	_ = v3993
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v3996 int32
	_ = v3996
	var v3997 int32
	_ = v3997
	var v3999 int32
	_ = v3999
	var v4005 int32
	_ = v4005
	var v4006 int32
	_ = v4006
	var v4008 int32
	_ = v4008
	var v4014 int32
	_ = v4014
	var v4022 int32
	_ = v4022
	var v4030 int32
	_ = v4030
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4037 int32
	_ = v4037
	var v4038 int32
	_ = v4038
	var v4041 int32
	_ = v4041
	var v4042 int32
	_ = v4042
	var v4044 int32
	_ = v4044
	var v4047 int32
	_ = v4047
	var v4050 int64
	_ = v4050
	var v4051 int32
	_ = v4051
	var v4052 int32
	_ = v4052
	var v4054 int32
	_ = v4054
	var v4055 int32
	_ = v4055
	var v4056 int32
	_ = v4056
	var v4059 int32
	_ = v4059
	var v4062 int32
	_ = v4062
	var v4063 int32
	_ = v4063
	var v4066 int32
	_ = v4066
	var v4084 int32
	_ = v4084
	var v4086 int32
	_ = v4086
	var v4103 int32
	_ = v4103
	var v4106 int32
	_ = v4106
	var v4107 int32
	_ = v4107
	var v4109 int32
	_ = v4109
	var v4122 int32
	_ = v4122
	var v4139 int32
	_ = v4139
	var v4140 int32
	_ = v4140
	var v4145 int32
	_ = v4145
	var v4153 int32
	_ = v4153
	var v4157 int32
	_ = v4157
	var v4162 int32
	_ = v4162
	var v4166 int32
	_ = v4166
	var v4169 int32
	_ = v4169
	var v4170 int32
	_ = v4170
	var v4171 int32
	_ = v4171
	var v4176 int32
	_ = v4176
	var v4181 int32
	_ = v4181
	var v4182 int64
	_ = v4182
	var v4186 int32
	_ = v4186
	var v4187 int32
	_ = v4187
	var v4188 int32
	_ = v4188
	var v4223 int32
	_ = v4223
	var v4226 int32
	_ = v4226
	var v4229 int32
	_ = v4229
	var v4232 int32
	_ = v4232
	var v4236 int32
	_ = v4236
	var v4240 int32
	_ = v4240
	var v4241 int32
	_ = v4241
	var v4243 int32
	_ = v4243
	var v4248 int32
	_ = v4248
	var v4252 int32
	_ = v4252
	var v4255 int32
	_ = v4255
	var v4259 int32
	_ = v4259
	var v4260 int32
	_ = v4260
	var v4262 int32
	_ = v4262
	var v4267 int32
	_ = v4267
	var v4271 int32
	_ = v4271
	var v4274 int32
	_ = v4274
	var v4275 int32
	_ = v4275
	var v4276 int32
	_ = v4276
	var v4284 int32
	_ = v4284
	var v4285 int32
	_ = v4285
	var v4287 int32
	_ = v4287
	var v4292 int32
	_ = v4292
	var v4296 int32
	_ = v4296
	var v4299 int32
	_ = v4299
	var v4306 int32
	_ = v4306
	var v4307 int32
	_ = v4307
	var v4309 int32
	_ = v4309
	var v4314 int32
	_ = v4314
	var v4318 int32
	_ = v4318
	var v4321 int32
	_ = v4321
	var v4328 int32
	_ = v4328
	var v4329 int32
	_ = v4329
	var v4331 int32
	_ = v4331
	var v4336 int32
	_ = v4336
	var v4340 int32
	_ = v4340
	var v4343 int32
	_ = v4343
	var v4347 int32
	_ = v4347
	var v4348 int32
	_ = v4348
	var v4350 int32
	_ = v4350
	var v4355 int32
	_ = v4355
	var v4356 int32
	_ = v4356
	var v4358 int32
	_ = v4358
	var v4364 int32
	_ = v4364
	var v4367 int32
	_ = v4367
	var v4368 int32
	_ = v4368
	var v4371 int32
	_ = v4371
	var v4372 int32
	_ = v4372
	var v4373 int32
	_ = v4373
	var v4374 int32
	_ = v4374
	var v4375 int32
	_ = v4375
	var v4378 int32
	_ = v4378
	var v4379 int32
	_ = v4379
	var v4380 int32
	_ = v4380
	var v4381 int32
	_ = v4381
	var v4382 int32
	_ = v4382
	var v4384 int32
	_ = v4384
	var v4385 int32
	_ = v4385
	var v4387 int32
	_ = v4387
	var v4388 int32
	_ = v4388
	var v4391 int32
	_ = v4391
	var v4393 int32
	_ = v4393
	var v4395 int32
	_ = v4395
	var v4397 int32
	_ = v4397
	var v4399 int32
	_ = v4399
	var v4433 int32
	_ = v4433
	var v4436 int32
	_ = v4436
	var v4437 int32
	_ = v4437
	var v4439 int32
	_ = v4439
	var v4441 int32
	_ = v4441
	var v4443 int32
	_ = v4443
	var v4444 int32
	_ = v4444
	var v4446 int32
	_ = v4446
	var v4448 int32
	_ = v4448
	var v4451 int32
	_ = v4451
	var v4452 int32
	_ = v4452
	var v4456 int32
	_ = v4456
	var v4457 int32
	_ = v4457
	var v4458 int32
	_ = v4458
	var v4460 int32
	_ = v4460
	var v4461 int32
	_ = v4461
	var v4462 int32
	_ = v4462
	var v4463 int32
	_ = v4463
	var v4464 int32
	_ = v4464
	var v4466 int32
	_ = v4466
	var v4469 int32
	_ = v4469
	var v4470 int32
	_ = v4470
	var v4479 int32
	_ = v4479
	var v4501 int32
	_ = v4501
	v3 = int32(0)
	v26 = int64(0)
	v28 = m.G0
	v30 = v28 - int32(176)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v32 - int32(137) {
	case 0:
		goto L33
	case 1:
		goto L32
	case 2:
		goto L31
	case 3:
		goto L30
	case 4:
		goto L29
	default:
		goto L22
	case 6:
		goto L28
	case 7:
		goto L27
	case 64:
		goto L26
	case 76:
		goto L23
	case 104:
		goto L25
	case 105:
		goto L24
	}
L1:
	;
	v4501 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4479)+24)) = uint8(v4501)
	*(*int32)(unsafe.Add(mBase, uint32(v4479)+8)) = int32(0)
	m.G0 = v30 + int32(176)
	return v4479
L2:
	;
	v3273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v3274 = int32(0)
	v3276 = F_setTargetTable(m, l0, v3273, v3274, v3274, v69)
	mBase = m.M
	v3277 = m.ExcPending
	if v3277 != 0 {
		goto L34
	} else {
		goto L699
	}
L3:
	;
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3183 != 0 {
		goto L676
	} else {
		goto L677
	}
L4:
	;
	if v1327 <= int32(0) {
		v3167 = v3
		v3168 = v1405
		v3170 = v3
		v3172 = v3
		goto L3
	} else {
		goto L642
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2921 = m.ExcPending
	if v2921 != 0 {
		goto L34
	} else {
		goto L639
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2905 = m.ExcPending
	if v2905 != 0 {
		goto L34
	} else {
		goto L636
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2889 = m.ExcPending
	if v2889 != 0 {
		goto L34
	} else {
		goto L632
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2873 = m.ExcPending
	if v2873 != 0 {
		goto L34
	} else {
		goto L628
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2846 = m.ExcPending
	if v2846 != 0 {
		goto L34
	} else {
		goto L622
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2830 = m.ExcPending
	if v2830 != 0 {
		goto L34
	} else {
		goto L618
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2795 = m.ExcPending
	if v2795 != 0 {
		goto L34
	} else {
		goto L609
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2760 = m.ExcPending
	if v2760 != 0 {
		goto L34
	} else {
		goto L600
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2725 = m.ExcPending
	if v2725 != 0 {
		goto L34
	} else {
		goto L591
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2709 = m.ExcPending
	if v2709 != 0 {
		goto L34
	} else {
		goto L587
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2696 = m.ExcPending
	if v2696 != 0 {
		goto L34
	} else {
		goto L584
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2675 = m.ExcPending
	if v2675 != 0 {
		goto L34
	} else {
		goto L580
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L34
	} else {
		goto L576
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L34
	} else {
		goto L573
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L34
	} else {
		goto L568
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2606 = m.ExcPending
	if v2606 != 0 {
		goto L34
	} else {
		goto L564
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L34
	} else {
		goto L560
	}
L22:
	;
	v2582 = F_palloc0(m, int32(168))
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L34
	} else {
		goto L559
	}
L23:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(v2315)+8))
	if v2316 == int32(0) {
		v2368 = v3
		v2370 = v2315
		goto L509
	} else {
		goto L510
	}
L24:
	;
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2281 = F_transformStmt(m, l0, v2280)
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L34
	} else {
		goto L497
	}
L25:
	;
	v2034 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+156)) = v2034
	*(*int32)(unsafe.Add(mBase, uint32(v30)+172)) = v2034
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v2038 == v2034 {
		goto L458
	} else {
		goto L459
	}
L26:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1998 = int32(6)
	if v1997&v1998 == v1998 {
		goto L17
	} else {
		goto L443
	}
L27:
	;
	v1887 = F_palloc0(m, int32(12))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L34
	} else {
		goto L425
	}
L28:
	;
	v1839 = F_palloc0(m, int32(168))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L34
	} else {
		goto L415
	}
L29:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v1255 != 0 {
		goto L268
	} else {
		goto L269
	}
L30:
	;
	v229 = m.G0
	v231 = v229 - int32(48)
	m.G0 = v231
	v234 = F_palloc0(m, int32(168))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L34
	} else {
		goto L91
	}
L31:
	;
	v163 = F_palloc0(m, int32(168))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L34
	} else {
		goto L74
	}
L32:
	;
	v89 = F_palloc0(m, int32(168))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L34
	} else {
		goto L57
	}
L33:
	;
	v36 = F_palloc0(m, int32(168))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return int32(0)
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(67)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(3)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v45 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+41)) = uint8(v46)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v49 = F_transformWithClause(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v57 == int32(0) {
		v69 = int64(1)
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = v49
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+42)) = uint8(v52)
	goto L38
L40:
	;
	if v42 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	switch v61 - int32(2) {
	case 0:
		v69 = int64(5)
		goto L40
	case 1:
		goto L43
	default:
		goto L42
	}
L42:
	;
	v69 = int64(1)
	goto L40
L43:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v66 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v67 = int64(5)
	goto L46
L45:
	;
	v67 = int64(1)
	goto L46
L46:
	;
	v69 = v67
	goto L40
L47:
	;
	v3269 = v3
	v3270 = v3
	v3271 = v3
	v3272 = v3
	goto L2
L48:
	;
	goto L49
L49:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v72 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v3269 = int32(1)
	v3270 = v80
	v3271 = v84
	v3272 = v81
	goto L2
L51:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+44))
	if v75 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	if v76 != 0 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	if v77 != 0 {
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v42)+60))
	if v78 != 0 {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v42)+64))
	if v79 != 0 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v3269 = v3
	v3270 = v3
	v3271 = v3
	v3272 = v3
	goto L2
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = int64(17179869251)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v93 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+41)) = uint8(v94)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v97 = F_transformWithClause(m, l0, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L34
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+16)))
	v106 = F_setTargetTable(m, l0, v102, v103, int32(1), int64(8))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L34
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+48)) = v97
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+42)) = uint8(v100)
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v106
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v110 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+120)) = int32(0)
	v123 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v109)+22)) = uint16(v123)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_transformFromClause(m, l0, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L34
	} else {
		goto L67
	}
L64:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v113 != int32(58) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+48))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+119)))
	if v118 == int32(118) {
		goto L21
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v128 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v109)+22)) = uint16(v128)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v133 = F_transformWhereClause(m, l0, v130, int32(6), int32(_a_F_transformStmt_0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L34
	} else {
		goto L68
	}
L68:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_transformReturningClause(m, l0, v89, v135, int32(24))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L34
	} else {
		goto L69
	}
L69:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+52)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+56)) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v144 = F_makeFromExpr(m, v143, v133)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L34
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+60)) = v144
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+39)) = uint8(v147)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+37)) = uint8(v149)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+38)) = uint8(v151)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+36)) = uint8(v153)
	F_assign_query_collations(m, l0, v89)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L34
	} else {
		goto L71
	}
L71:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v157 != int32(1) {
		v4479 = v89
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_parseCheckAggregates(m, l0, v89)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L34
	} else {
		goto L73
	}
L73:
	;
	v4479 = v89
	goto L1
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = int64(8589934659)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v167 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+41)) = uint8(v168)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v171 = F_transformWithClause(m, l0, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L34
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+16)))
	v180 = F_setTargetTable(m, l0, v176, v177, int32(1), int64(4))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L34
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+48)) = v171
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+42)) = uint8(v174)
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+32)) = v180
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v183 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v195 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v194)+22)) = uint16(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_transformFromClause(m, l0, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L34
	} else {
		goto L84
	}
L81:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v186 != int32(58) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+48))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+119)))
	if v191 == int32(118) {
		goto L20
	} else {
		goto L83
	}
L83:
	;
	goto L80
L84:
	;
	v200 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v194)+22)) = uint16(v200)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v205 = F_transformWhereClause(m, l0, v202, int32(6), int32(_a_F_transformStmt_0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L34
	} else {
		goto L85
	}
L85:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_transformReturningClause(m, l0, v163, v207, int32(24))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L34
	} else {
		goto L86
	}
L86:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v212 = F_transformUpdateTargetList(m, l0, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L34
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+76)) = v212
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+52)) = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+56)) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v220 = F_makeFromExpr(m, v219, v205)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L34
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+60)) = v220
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+38)) = uint8(v223)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+39)) = uint8(v225)
	F_assign_query_collations(m, l0, v163)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L34
	} else {
		goto L89
	}
L89:
	;
	v4479 = v163
	goto L1
L90:
	;
	v4479 = v234
	goto L1
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(67)
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+41)) = uint8(v238)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = int32(5)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v242 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+16)))
	v403 = F_setTargetTable(m, l0, v400, v401, int32(0), v398)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L34
	} else {
		goto L126
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L34
	} else {
		goto L122
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L34
	} else {
		goto L118
	}
L95:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+8)))
	if v243 == int32(1) {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v251 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+47)) = uint8(v251)
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+45)) = uint16(v251)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v255 == v251 {
		v398 = v26
		goto L92
	} else {
		goto L100
	}
L98:
	;
	v246 = F_transformWithClause(m, l0, v242)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L34
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+48)) = v246
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+42)) = uint8(v249)
	goto L97
L100:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v258 <= int32(0) {
		v398 = v26
		goto L92
	} else {
		goto L101
	}
L101:
	;
	v261 = int32(0)
	if v261 < v258 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v264 = v258
	goto L104
L103:
	;
	v264 = v261
	goto L104
L104:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v273 = v3
	v291 = v26
	goto L105
L105:
	;
	v293 = int32(2)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v265+v273<<(uint(v293)%32))))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+8))
	v299 = v297 - v293
	if base.B2i32(base.Ui32(v299) <= base.Ui32(int32(5)))&(int32(base.Ui32(int32(39))>>(uint(v299)%32))&int32(1)) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v398 = v337
	goto L92
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L34
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	v325 = v322 + (v231 + int32(45))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	if v326 == int32(1) {
		goto L93
	} else {
		goto L113
	}
L110:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_1), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L34
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(167), int32(_a_F_transformStmt_3))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L34
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
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v299<<(uint(int32(3))%32))+uint32(_c_F_transformStmt[0])))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v296)+16))
	if v332 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v335 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v325))) = uint8(v335)
	goto L116
L115:
	;
	goto L116
L116:
	;
	v337 = v291 | v331
	v339 = v273 + int32(1)
	if v339 != v264 {
		v273 = v339
		v291 = v337
		goto L105
	} else {
		goto L117
	}
L117:
	;
	goto L106
L118:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L34
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_transformStmt_4), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L34
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(129), int32(_a_F_transformStmt_3))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L34
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L34
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(_a_F_transformStmt_5), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L34
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(176), int32(_a_F_transformStmt_3))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L34
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+68)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v234)+32)) = v403
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+48))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+119)))
	v411 = v409 - int32(112)
	if int32(1)<<(uint(v411)%32)&int32(69) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v419 = base.B2i32(base.Ui32(v411) <= base.Ui32(int32(6)))
	goto L129
L128:
	;
	v419 = int32(0)
	goto L129
L129:
	;
	if v419 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L34
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+32)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v447
	v453 = F_list_make1_impl(m, int32(1), v231+int32(32))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L34
	} else {
		goto L138
	}
L133:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L34
	} else {
		goto L134
	}
L134:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v430 + int32(4)
	F_errmsg(m, int32(_a_F_transformStmt_6), v231)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L34
	} else {
		goto L135
	}
L135:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+48))
	v439 = int32(*(*int8)(unsafe.Add(mBase, uint32(v438)+119)))
	F_errdetail_relkind_not_supported(m, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L34
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(205), int32(_a_F_transformStmt_3))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L34
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_transformFromClause(m, l0, v453)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L34
	} else {
		goto L139
	}
L139:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v457 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
	v459 = v458
	goto L142
L141:
	;
	v459 = v3
	goto L142
L142:
	;
	v460 = int32(0)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v461 == v460 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+4))
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552))))
	if base.B2i32(v555 == int32(0))|base.B2i32(v555 != v558) != 0 {
		v576 = v555
		v577 = v558
		goto L160
	} else {
		goto L161
	}
L144:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L34
	} else {
		goto L156
	}
L145:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v464 <= int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v467 = int32(0)
	if v467 < v464 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v470 = v464
	goto L149
L148:
	;
	v470 = v467
	goto L149
L149:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v475 = v460
	goto L150
L150:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v471+v475<<(uint(int32(2))%32))))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	if v459 != v503 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L143
L152:
	;
	v506 = v475 + int32(1)
	if v470 != v506 {
		v475 = v506
		goto L150
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	goto L151
L155:
	;
	goto L144
L156:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_7), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L34
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_transformStmt_8), int32(535), int32(_a_F_transformStmt_9))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L34
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	if v576-v577 != 0 {
		goto L166
	} else {
		goto L167
	}
L160:
	;
	goto L159
L161:
	;
	v561 = v550
	v562 = v552
	goto L162
L162:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+1)))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561)+1)))
	if v566 == int32(0) {
		v576 = v566
		v577 = v565
		goto L160
	} else {
		goto L164
	}
L163:
	;
	v576 = v566
	v577 = v565
	goto L160
L164:
	;
	v569 = int32(1)
	if v566 == v565 {
		v561 = v561 + v569
		v562 = v562 + v569
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v579 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+76)) = v579
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+52)) = v581
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+56)) = v583
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v587 = int32(1)
	F_addNSItemToQuery(m, l0, v585, v579, v587, v587)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L34
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L34
	} else {
		goto L263
	}
L169:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v593 = F_transformExpr(m, l0, v591, int32(2))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L34
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+72)) = v593
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v598 = F_makeFromExpr(m, v596, int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L34
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+60)) = v598
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_transformReturningClause(m, l0, v234, v601, int32(25))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L34
	} else {
		goto L172
	}
L172:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v605 == int32(0) {
		v1209 = v3
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v1220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+38)) = uint8(v1220)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+64)) = v1209
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+39)) = uint8(v1223)
	F_assign_query_collations(m, l0, v234)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L34
	} else {
		goto L262
	}
L174:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	if v608 <= int32(0) {
		v1209 = v3
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v623 = v3
	v627 = v3
	goto L176
L176:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v605)+12))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v638+v623<<(uint(int32(2))%32))))
	v644 = F_palloc0(m, int32(28))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L34
	} else {
		goto L178
	}
L177:
	;
	v1209 = v1187
	goto L173
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = int32(54)
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v642)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v644)+8)) = v648
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v644)+4)) = v650
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)+12))
	v654 = int32(2)
	v657 = int32(4)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v653+v459<<(uint(v654)%32)-v657)))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v234)+32))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v653+v660<<(uint(v654)%32)-v657)))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v650 {
	case 0:
		goto L184
	case 1:
		goto L183
	default:
		goto L182
	}
L179:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v642)+16))
	v1029 = F_transformWhereClause(m, l0, v1026, int32(18), int32(_a_F_transformStmt_10))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L34
	} else {
		goto L230
	}
L180:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v973)+21)) = uint8(v980)
	*(*uint8)(unsafe.Add(mBase, uint32(v973)+20)) = uint8(v980)
	goto L179
L181:
	;
	v973 = v945
	v980 = int32(1)
	goto L180
L182:
	;
	if v667 == int32(0) {
		goto L179
	} else {
		goto L215
	}
L183:
	;
	if v667 == int32(0) {
		goto L179
	} else {
		goto L200
	}
L184:
	;
	if v667 == int32(0) {
		goto L179
	} else {
		goto L185
	}
L185:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	if v670 <= int32(0) {
		goto L179
	} else {
		goto L186
	}
L186:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v667)+12))
	v678 = int32(0)
	goto L188
L187:
	;
	if v715 == int32(0) {
		goto L179
	} else {
		goto L194
	}
L188:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v673+v678<<(uint(int32(2))%32))))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+4))
	if v666 != v706 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v711 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v705)+20)) = uint16(v711)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v715 = v713
	goto L187
L190:
	;
	v709 = v678 + int32(1)
	if v709 != v670 {
		v678 = v709
		goto L188
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	goto L189
L193:
	;
	v715 = v667
	goto L187
L194:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v715)+4))
	if v718 <= int32(0) {
		goto L179
	} else {
		goto L195
	}
L195:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v715)+12))
	v730 = int32(0)
	goto L196
L196:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v721+v730<<(uint(int32(2))%32))))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)+4))
	if v754 == v659 {
		v945 = v753
		goto L181
	} else {
		goto L198
	}
L197:
	;
	goto L179
L198:
	;
	v757 = v730 + int32(1)
	if v757 != v718 {
		v730 = v757
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	if v761 <= int32(0) {
		goto L179
	} else {
		goto L201
	}
L201:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v667)+12))
	v769 = int32(0)
	goto L203
L202:
	;
	if v806 == int32(0) {
		goto L179
	} else {
		goto L209
	}
L203:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v764+v769<<(uint(int32(2))%32))))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v796)+4))
	if v666 != v797 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v802 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v796)+20)) = uint16(v802)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v806 = v804
	goto L202
L205:
	;
	v800 = v769 + int32(1)
	if v800 != v761 {
		v769 = v800
		goto L203
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	goto L204
L208:
	;
	v806 = v667
	goto L202
L209:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v806)+4))
	if v809 <= int32(0) {
		goto L179
	} else {
		goto L210
	}
L210:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v806)+12))
	v813 = int32(0)
	v822 = v813
	goto L211
L211:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v812+v822<<(uint(int32(2))%32))))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v845)+4))
	if v846 == v659 {
		v973 = v845
		v980 = v813
		goto L180
	} else {
		goto L213
	}
L212:
	;
	goto L179
L213:
	;
	v849 = v822 + int32(1)
	if v849 != v809 {
		v822 = v849
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	if v853 <= int32(0) {
		goto L179
	} else {
		goto L216
	}
L216:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v667)+12))
	v861 = int32(0)
	goto L218
L217:
	;
	if v898 == int32(0) {
		goto L179
	} else {
		goto L224
	}
L218:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v856+v861<<(uint(int32(2))%32))))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v888)+4))
	if v666 != v889 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v894 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v888)+20)) = uint16(v894)
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v898 = v896
	goto L217
L220:
	;
	v892 = v861 + int32(1)
	if v892 != v853 {
		v861 = v892
		goto L218
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	goto L219
L223:
	;
	v898 = v667
	goto L217
L224:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v898)+4))
	if v901 <= int32(0) {
		goto L179
	} else {
		goto L225
	}
L225:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v898)+12))
	v913 = int32(0)
	goto L226
L226:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v904+v913<<(uint(int32(2))%32))))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	if v937 == v659 {
		v945 = v936
		goto L181
	} else {
		goto L228
	}
L227:
	;
	goto L179
L228:
	;
	v940 = v913 + int32(1)
	if v940 != v901 {
		v913 = v940
		goto L226
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v644)+16)) = v1029
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v644)+8))
	switch v1032 - int32(2) {
	case 0:
		goto L232
	case 1:
		goto L235
	case 2:
		goto L231
	default:
		goto L233
	case 5:
		goto L234
	}
L231:
	;
	v1187 = F_lappend(m, v627, v644)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L34
	} else {
		goto L260
	}
L232:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v642)+20))
	v1157 = F_transformUpdateTargetList(m, l0, v1156)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L34
	} else {
		goto L259
	}
L233:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L34
	} else {
		goto L256
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v644)+20)) = int32(0)
	goto L231
L235:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v642)+20))
	v1038 = F_checkInsertTargets(m, l0, v1035, v231+int32(36))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L34
	} else {
		goto L236
	}
L236:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v642)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v644)+12)) = v1040
	v1042 = int32(0)
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v642)+24))
	if v1044 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1047 = F_transformExpressionList(m, l0, v1044, int32(27), int32(1))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L34
	} else {
		goto L240
	}
L238:
	;
	v1054 = v1042
	goto L239
L239:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+12))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v1065 = v1042
	goto L242
L240:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v642)+20))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v1052 = F_transformInsertRow(m, l0, v1047, v1049, v1038, v1050, int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L34
	} else {
		goto L241
	}
L241:
	;
	v1054 = v1052
	goto L239
L242:
	;
	v1085 = int32(0)
	if v1054 == v1085 {
		v1096 = v1085
		goto L244
	} else {
		goto L245
	}
L244:
	;
	if v1038 == int32(0) {
		v1105 = v1085
		goto L247
	} else {
		goto L248
	}
L245:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+4))
	if v1090 <= v1065 {
		v1096 = int32(0)
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+12))
	v1096 = v1092 + v1065<<(uint(int32(2))%32)
	goto L244
L247:
	;
	if v1057 == int32(0) {
		goto L231
	} else {
		goto L250
	}
L248:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+4))
	if v1099 <= v1065 {
		v1105 = v1085
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+12))
	v1105 = v1101 + v1065<<(uint(int32(2))%32)
	goto L247
L250:
	;
	v1108 = int32(0)
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if base.B2i32(v1105 == v1108)|(base.B2i32(v1096 == v1108)|base.B2i32(v1112 <= v1065)) != 0 {
		goto L231
	} else {
		goto L251
	}
L251:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	if v1116 == int32(0) {
		goto L231
	} else {
		goto L252
	}
L252:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	v1123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1116+v1065<<(uint(int32(2))%32)))))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1105)))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+4))
	v1127 = F_makeTargetEntry(m, v1119, v1123, v1125, int32(0))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L34
	} else {
		goto L253
	}
L253:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v644)+20))
	v1130 = F_lappend(m, v1129, v1127)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L34
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v644)+20)) = v1130
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+32))
	v1136 = F_bms_add_member(m, v1133, v1123+int32(7))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L34
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1056)+32)) = v1136
	v1065 = v1065 + int32(1)
	goto L242
L256:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_1), int32(0))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L34
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(393), int32(_a_F_transformStmt_3))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L34
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
	*(*int32)(unsafe.Add(mBase, uint32(v644)+20)) = v1157
	goto L231
L260:
	;
	v1190 = v623 + int32(1)
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	if v1190 < v1191 {
		v623 = v1190
		v627 = v1187
		goto L176
	} else {
		goto L261
	}
L261:
	;
	goto L177
L262:
	;
	m.G0 = v231 + int32(48)
	goto L90
L263:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L34
	} else {
		goto L264
	}
L264:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1237)))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+16)) = v1239
	F_errmsg(m, int32(_a_F_transformStmt_11), v231+int32(16))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L34
	} else {
		goto L265
	}
L265:
	;
	v1248 = F_errdetail(m, int32(_a_F_transformStmt_12), int32(0))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L34
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_transformStmt_2), int32(224), int32(_a_F_transformStmt_3))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L34
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
	v1257 = F_palloc0(m, int32(168))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L34
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v1411 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L271:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1257))) = int64(4294967363)
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v1261 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1261)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1257)+41)) = uint8(v1262)
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1265 = F_transformWithClause(m, l0, v1264)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L34
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v1270 == int32(0) {
		v3167 = v3
		v3168 = v3
		v3170 = v3
		v3172 = v3
		goto L3
	} else {
		goto L276
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+48)) = v1265
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1257)+42)) = uint8(v1268)
	goto L274
L276:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+4))
	if v1273 <= int32(0) {
		v3167 = v3
		v3168 = v3
		v3170 = v3
		v3172 = v3
		goto L3
	} else {
		goto L277
	}
L277:
	;
	v1283 = v3
	v1285 = v3
	v1287 = int32(-1)
	v1289 = v3
	goto L278
L278:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+12))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1304+v1283<<(uint(int32(2))%32))))
	v1311 = F_transformExpressionList(m, l0, v1308, int32(26), int32(0))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L34
	} else {
		goto L280
	}
L279:
	;
	goto L4
L280:
	;
	if v1287 < int32(0) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	if v1311 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L282:
	;
	if v1311 != 0 {
		goto L285
	} else {
		goto L286
	}
L283:
	;
	goto L284
L284:
	;
	if v1311 != 0 {
		goto L289
	} else {
		goto L290
	}
L285:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4))
	v1317 = v1315
	goto L287
L286:
	;
	v1317 = int32(0)
	goto L287
L287:
	;
	v1320 = F_palloc0(m, v1317<<(uint(int32(2))%32))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L34
	} else {
		goto L288
	}
L288:
	;
	v1326 = v1320
	v1327 = v1317
	goto L281
L289:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4))
	v1324 = v1322
	goto L291
L290:
	;
	v1324 = int32(0)
	goto L291
L291:
	;
	if v1324 != v1287 {
		goto L19
	} else {
		goto L292
	}
L292:
	;
	v1326 = v1285
	v1327 = v1287
	goto L281
L293:
	;
	F_list_free(m, v1311)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L34
	} else {
		goto L300
	}
L294:
	;
	v1330 = int32(0)
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4))
	if v1331 <= v1330 {
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1337 = v1330
	goto L296
L296:
	;
	v1362 = v1337 << (uint(int32(2)) % 32)
	v1363 = v1326 + v1362
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1363)))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+12))
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1365+v1362)))
	v1368 = F_lappend(m, v1364, v1367)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L34
	} else {
		goto L298
	}
L297:
	;
	goto L293
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1368
	v1372 = v1337 + int32(1)
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4))
	if v1372 < v1373 {
		v1337 = v1372
		goto L296
	} else {
		goto L299
	}
L299:
	;
	goto L297
L300:
	;
	v1405 = F_lappend(m, v1289, int32(0))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L34
	} else {
		goto L301
	}
L301:
	;
	v1408 = v1283 + int32(1)
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+4))
	if v1408 < v1409 {
		v1283 = v1408
		v1285 = v1326
		v1287 = v1327
		v1289 = v1405
		goto L278
	} else {
		goto L302
	}
L302:
	;
	goto L279
L303:
	;
	v1415 = F_transformSelectStmt(m, l0, l1, int32(0))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L34
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1417 = m.G0
	v1419 = v1417 - int32(16)
	m.G0 = v1419
	v1422 = F_palloc0(m, int32(168))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L34
	} else {
		goto L308
	}
L306:
	;
	v4479 = v1415
	goto L1
L307:
	;
	v4479 = v1422
	goto L1
L308:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1422))) = int64(4294967363)
	v1427 = l1
	goto L309
L309:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+76))
	if v1453 != 0 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1453)+8))
	if v1455 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L311:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1453)+68))
	if v1454 != 0 {
		v1427 = v1453
		goto L309
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	goto L310
L314:
	;
	goto L313
L315:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v1460 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+60)) = v1460
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1463 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v1463
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+44)) = v1460
	if v1459 == v1463 {
		goto L318
	} else {
		goto L319
	}
L316:
	;
	goto L317
L317:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L34
	} else {
		goto L410
	}
L318:
	;
	if v1458 != 0 {
		goto L321
	} else {
		goto L322
	}
L319:
	;
	goto L320
L320:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L34
	} else {
		goto L402
	}
L321:
	;
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1458)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422)+41)) = uint8(v1471)
	v1473 = F_transformWithClause(m, l0, v1458)
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L34
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	v1480 = F_transformSetOperationTree(m, l0, l1, int32(1), int32(0))
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L34
	} else {
		goto L325
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+48)) = v1473
	v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422)+42)) = uint8(v1476)
	goto L323
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+144)) = v1480
	v1484 = v1480
	goto L326
L326:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1484)+12))
	if v1510 != 0 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+12))
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1510)+4))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1515+v1516<<(uint(int32(2))%32)-int32(4))))
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+76)) = int32(0)
	v1527 = v1422 + int32(76)
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+20))
	if v1528 != 0 {
		goto L332
	} else {
		goto L333
	}
L328:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1510)))
	if v1511 == int32(142) {
		v1484 = v1510
		goto L326
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	goto L327
L331:
	;
	goto L330
L332:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1528)+4))
	v1533 = v1529 << (uint(int32(5)) % 32)
	goto L334
L333:
	;
	v1533 = int32(0)
	goto L334
L334:
	;
	v1534 = F_palloc0(m, v1533)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L34
	} else {
		goto L335
	}
L335:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1523)+76))
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+28))
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+24))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+20))
	v1547 = v3
	v1552 = v3
	v1557 = v3
	goto L336
L336:
	;
	v1567 = int32(0)
	if v1539 == v1567 {
		v1578 = v1567
		goto L338
	} else {
		goto L339
	}
L337:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L34
	} else {
		goto L395
	}
L338:
	;
	if v1538 == int32(0) {
		v1587 = v1567
		goto L341
	} else {
		goto L342
	}
L339:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1539)+4))
	if v1572 <= v1547 {
		v1578 = int32(0)
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1539)+12))
	v1578 = v1574 + v1547<<(uint(int32(2))%32)
	goto L338
L341:
	;
	v1588 = int32(0)
	if v1537 == v1588 {
		v1599 = v1588
		goto L344
	} else {
		goto L345
	}
L342:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1538)+4))
	if v1581 <= v1547 {
		v1587 = v1567
		goto L341
	} else {
		goto L343
	}
L343:
	;
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1538)+12))
	v1587 = v1583 + v1547<<(uint(int32(2))%32)
	goto L341
L344:
	;
	if v1536 == int32(0) {
		v1608 = v1588
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1537)+4))
	if v1593 <= v1547 {
		v1599 = int32(0)
		goto L344
	} else {
		goto L346
	}
L346:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1537)+12))
	v1599 = v1595 + v1547<<(uint(int32(2))%32)
	goto L344
L347:
	;
	v1609 = int32(0)
	if v1608 != 0 {
		goto L351
	} else {
		goto L352
	}
L348:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+4))
	if v1602 <= v1547 {
		v1608 = v1588
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+12))
	v1608 = v1604 + v1547<<(uint(int32(2))%32)
	goto L347
L350:
	;
	goto L337
L351:
	;
	v1618 = base.B2i32(v1599 == v1609) | (base.B2i32(v1578 == v1609) | base.B2i32(v1587 == v1609))
	goto L353
L352:
	;
	v1618 = int32(1)
	goto L353
L353:
	;
	if v1618 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1619 = int32(0)
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1621 != 0 {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	goto L356
L356:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1599)))
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1587)))
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1578)))
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1608)))
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1714)+12))
	v1716 = F_pstrdup(m, v1715)
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L34
	} else {
		goto L388
	}
L357:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1621)+4))
	v1623 = v1622
	goto L359
L358:
	;
	v1623 = v1619
	goto L359
L359:
	;
	v1624 = int32(0)
	v1631 = F_addRangeTableEntryForJoin(m, l0, v1552, v1534, v1624, v1624, v1557, v1624, v1624, v1624, v1624, v1624)
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L34
	} else {
		goto L360
	}
L360:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1634 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1634
	F_addNSItemToQuery(m, l0, v1631, v1634, v1634, int32(1))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L34
	} else {
		goto L361
	}
L361:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1527)))
	if v1641 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1641)+4))
	v1643 = v1642
	goto L364
L363:
	;
	v1643 = v1619
	goto L364
L364:
	;
	v1645 = F_transformSortClause(m, l0, v1466, v1527, int32(0))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L34
	} else {
		goto L365
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+124)) = v1645
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1633
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1650 = int32(0)
	if base.B2i32(v1649 == v1650)|base.B2i32(v1623 <= v1650) != 0 {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1660
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+76))
	if v1662 != 0 {
		goto L373
	} else {
		goto L374
	}
L367:
	;
	v1660 = int32(0)
	goto L369
L368:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+4))
	if v1623 < v1657 {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	goto L366
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1649)+4)) = v1623
	goto L372
L371:
	;
	goto L372
L372:
	;
	v1660 = v1649
	goto L369
L373:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1662)+4))
	v1665 = v1663
	goto L375
L374:
	;
	v1665 = int32(0)
	goto L375
L375:
	;
	if v1665 != v1643 {
		goto L350
	} else {
		goto L376
	}
L376:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1670 = F_transformLimitClause(m, l0, v1465, int32(23), int32(_a_F_transformStmt_13), v1669)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L34
	} else {
		goto L377
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+128)) = v1670
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1676 = F_transformLimitClause(m, l0, v1462, int32(22), int32(_a_F_transformStmt_14), v1675)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L34
	} else {
		goto L378
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+132)) = v1676
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+136)) = v1679
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+52)) = v1681
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+56)) = v1683
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1687 = F_makeFromExpr(m, v1685, int32(0))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L34
	} else {
		goto L379
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+60)) = v1687
	v1690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422)+39)) = uint8(v1690)
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422)+37)) = uint8(v1692)
	v1694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422)+38)) = uint8(v1694)
	v1696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422)+36)) = uint8(v1696)
	F_assign_query_collations(m, l0, v1422)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L34
	} else {
		goto L380
	}
L380:
	;
	v1700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v1700 != 0 {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	m.G0 = v1419 + int32(16)
	goto L307
L382:
	;
	F_parseCheckAggregates(m, l0, v1422)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L34
	} else {
		goto L387
	}
L383:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+100))
	if v1701 != 0 {
		goto L382
	} else {
		goto L384
	}
L384:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+108))
	if v1702 != 0 {
		goto L382
	} else {
		goto L385
	}
L385:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+112))
	if v1703 == int32(0) {
		goto L381
	} else {
		goto L386
	}
L386:
	;
	goto L382
L387:
	;
	goto L381
L388:
	;
	v1718 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1714)+8)))
	v1720 = F_makeVar(m, v1516, v1718, v1713, v1712, v1711, int32(0))
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L34
	} else {
		goto L389
	}
L389:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1714)+4))
	v1723 = F_exprLocation(m, v1722)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1720)+44)) = v1723
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v1725 + int32(1)
	v1731 = F_makeTargetEntry(m, v1720, base.I32_extend16_s(v1725), v1716, int32(0))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L34
	} else {
		goto L390
	}
L390:
	;
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1527)))
	v1734 = F_lappend(m, v1733, v1731)
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L34
	} else {
		goto L391
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1527))) = v1734
	v1737 = F_lappend(m, v1557, v1720)
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L34
	} else {
		goto L392
	}
L392:
	;
	v1739 = F_makeString(m, v1716)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L34
	} else {
		goto L393
	}
L393:
	;
	v1741 = F_lappend(m, v1552, v1739)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L34
	} else {
		goto L394
	}
L394:
	;
	v1745 = v1534 + v1547<<(uint(int32(5))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1745))) = v1516
	v1747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1714)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v1745)+24)) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v1745)+16)) = v1711
	*(*int32)(unsafe.Add(mBase, uint32(v1745)+12)) = v1712
	*(*int32)(unsafe.Add(mBase, uint32(v1745)+8)) = v1713
	*(*uint16)(unsafe.Add(mBase, uint32(v1745)+4)) = uint16(v1747)
	v1753 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1714)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1745)+28)) = uint16(v1753)
	v1547 = v1547 + int32(1)
	v1552 = v1741
	v1557 = v1737
	goto L336
L395:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L34
	} else {
		goto L396
	}
L396:
	;
	F_errmsg(m, int32(_a_F_transformStmt_15), int32(0))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L34
	} else {
		goto L397
	}
L397:
	;
	v1770 = F_errdetail(m, int32(_a_F_transformStmt_16), int32(0))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L34
	} else {
		goto L398
	}
L398:
	;
	F_errhint(m, int32(_a_F_transformStmt_17), int32(0))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L34
	} else {
		goto L399
	}
L399:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1527)))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+12))
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1777+v1643<<(uint(int32(2))%32))))
	v1782 = F_exprLocation(m, v1781)
	mBase = m.M
	F_parser_errposition(m, l0, v1782)
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L34
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(2008), int32(_a_F_transformStmt_19))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L34
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L34
	} else {
		goto L403
	}
L403:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+12))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1797)))
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+8))
	v1801 = v1799 - int32(1)
	if base.Ui32(v1801) <= base.Ui32(int32(3)) {
		goto L405
	} else {
		goto L406
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1419))) = v1808
	F_errmsg(m, int32(_a_F_transformStmt_20), v1419)
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L34
	} else {
		goto L408
	}
L405:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1801<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v1808 = v1806
	goto L407
L406:
	;
	v1808 = int32(_a_F_transformStmt_21)
	goto L407
L407:
	;
	goto L404
L408:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(1866), int32(_a_F_transformStmt_19))
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L34
	} else {
		goto L409
	}
L409:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L410:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L34
	} else {
		goto L411
	}
L411:
	;
	F_errmsg(m, int32(_a_F_transformStmt_22), int32(0))
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L34
	} else {
		goto L412
	}
L412:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1453)+8))
	v1830 = F_exprLocation(m, v1829)
	mBase = m.M
	F_parser_errposition(m, l0, v1830)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L34
	} else {
		goto L413
	}
L413:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(1839), int32(_a_F_transformStmt_19))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L34
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
	v1841 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1839)+46)) = uint8(v1841)
	*(*int64)(unsafe.Add(mBase, uint32(v1839))) = int64(4294967363)
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1847 = F_transformExpr(m, l0, v1845, int32(14))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L34
	} else {
		goto L416
	}
L416:
	;
	v1850 = int32(0)
	v1852 = F_makeTargetEntry(m, v1847, int32(1), v1850, v1850)
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L34
	} else {
		goto L417
	}
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v1852
	*(*int32)(unsafe.Add(mBase, uint32(v30)+156)) = v1852
	v1859 = F_list_make1_impl(m, int32(1), v30+int32(8))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L34
	} else {
		goto L418
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1839)+76)) = v1859
	v1862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+81)))
	if v1862 == int32(1) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	F_resolveTargetListUnknowns(m, l0, v1859)
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L34
	} else {
		goto L422
	}
L420:
	;
	goto L421
L421:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1839)+52)) = v1867
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1839)+56)) = v1869
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1873 = F_makeFromExpr(m, v1871, int32(0))
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L34
	} else {
		goto L423
	}
L422:
	;
	goto L421
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1839)+60)) = v1873
	v1876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1839)+39)) = uint8(v1876)
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1839)+37)) = uint8(v1878)
	v1880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1839)+38)) = uint8(v1880)
	v1882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1839)+36)) = uint8(v1882)
	F_assign_query_collations(m, l0, v1839)
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L34
	} else {
		goto L424
	}
L424:
	;
	v4479 = v1839
	goto L1
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1887))) = int32(69)
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1894 = F_makeString(m, v1893)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L34
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v1894
	*(*int32)(unsafe.Add(mBase, uint32(v30)+172)) = v1894
	v1901 = F_list_make1_impl(m, int32(1), v30+int32(12))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L34
	} else {
		goto L427
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1887)+4)) = v1901
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1887)+8)) = v1904
	if v1892 < int32(2) {
		v1958 = v1891
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1983 = F_transformExpr(m, l0, v1887, int32(17))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L34
	} else {
		goto L441
	}
L429:
	;
	v1908 = F_list_copy(m, v1891)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L34
	} else {
		goto L430
	}
L430:
	;
	if v1908 == int32(0) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1958 = int32(0)
	goto L428
L432:
	;
	goto L433
L433:
	;
	v1916 = v1908
	v1917 = v1892
	goto L434
L434:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1916)+12))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1940)))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1941)))
	if v1942 != int32(476) {
		goto L18
	} else {
		goto L436
	}
L435:
	;
	v1958 = v1949
	goto L428
L436:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1887)+4))
	v1946 = F_lappend(m, v1945, v1941)
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L34
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1887)+4)) = v1946
	v1949 = F_list_delete_first(m, v1916)
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L34
	} else {
		goto L438
	}
L438:
	;
	if v1917 < int32(3) {
		v1958 = v1949
		goto L428
	} else {
		goto L439
	}
L439:
	;
	if v1949 != 0 {
		v1916 = v1949
		v1917 = v1917 - int32(1)
		goto L434
	} else {
		goto L440
	}
L440:
	;
	goto L435
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+164)) = v1958
	*(*int32)(unsafe.Add(mBase, uint32(v30)+160)) = v1983
	*(*int32)(unsafe.Add(mBase, uint32(v30)+156)) = l1
	v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+81)))
	v1989 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+81)) = uint8(v1989)
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1994 = F_transformSelectStmt(m, l0, v1991, v30+int32(156))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L34
	} else {
		goto L442
	}
L442:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+81)) = uint8(v1988)
	v4479 = v1994
	goto L1
L443:
	;
	v2002 = int32(24)
	if v1997&v2002 == v2002 {
		goto L16
	} else {
		goto L444
	}
L444:
	;
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2007 = F_transformStmt(m, l0, v2006)
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L34
	} else {
		goto L445
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v2007
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v2007)))
	if v2010 != int32(67) {
		goto L15
	} else {
		goto L446
	}
L446:
	;
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2007)+4))
	if v2013 != int32(1) {
		goto L15
	} else {
		goto L447
	}
L447:
	;
	v2016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2007)+42)))
	if v2016 == int32(1) {
		goto L14
	} else {
		goto L448
	}
L448:
	;
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v2007)+140))
	if v2019 != 0 {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v2020&int32(32) != 0 {
		goto L13
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	v2029 = F_palloc0(m, int32(168))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L34
	} else {
		goto L455
	}
L452:
	;
	if v2020&int32(2) != 0 {
		goto L12
	} else {
		goto L453
	}
L453:
	;
	if v2020&int32(8) != 0 {
		goto L11
	} else {
		goto L454
	}
L454:
	;
	goto L451
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2029)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2029))) = int64(25769803843)
	v4479 = v2029
	goto L1
L456:
	;
	v2275 = F_palloc0(m, int32(168))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L34
	} else {
		goto L496
	}
L457:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2244 = F_transformOptionalSelectInto(m, l0, v2243)
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		goto L34
	} else {
		goto L495
	}
L458:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v2041 == int32(0) {
		goto L457
	} else {
		goto L461
	}
L459:
	;
	goto L460
L460:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v2132)))
	if v2133 != int32(141) {
		goto L482
	} else {
		goto L483
	}
L461:
	;
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+4))
	if v2044 <= int32(0) {
		goto L457
	} else {
		goto L462
	}
L462:
	;
	v2050 = v3
	v2053 = v3
	goto L463
L463:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+12))
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v2074+v2050<<(uint(int32(2))%32))))
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v2078)+8))
	v2080 = int32(_a_F_transformStmt_23)
	v2083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2079))))
	v2086 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformStmt[2])))
	if base.B2i32(v2083 == int32(0))|base.B2i32(v2083 != v2086) != 0 {
		v2104 = v2083
		v2105 = v2086
		goto L466
	} else {
		goto L467
	}
L464:
	;
	if v2111&int32(1) == int32(0) {
		goto L457
	} else {
		goto L477
	}
L465:
	;
	if v2104-v2105 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L466:
	;
	goto L465
L467:
	;
	v2089 = v2079
	v2090 = v2080
	goto L468
L468:
	;
	v2093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2090)+1)))
	v2094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2089)+1)))
	if v2094 == int32(0) {
		v2104 = v2094
		v2105 = v2093
		goto L466
	} else {
		goto L470
	}
L469:
	;
	v2104 = v2094
	v2105 = v2093
	goto L466
L470:
	;
	v2097 = int32(1)
	if v2094 == v2093 {
		v2089 = v2089 + v2097
		v2090 = v2090 + v2097
		goto L468
	} else {
		goto L471
	}
L471:
	;
	goto L469
L472:
	;
	v2109 = F_defGetBoolean(m, v2078)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L34
	} else {
		goto L475
	}
L473:
	;
	v2111 = v2053
	goto L474
L474:
	;
	v2113 = v2050 + int32(1)
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+4))
	if v2113 < v2114 {
		v2050 = v2113
		v2053 = v2111
		goto L463
	} else {
		goto L476
	}
L475:
	;
	v2111 = v2109
	goto L474
L476:
	;
	goto L464
L477:
	;
	F_setup_parse_variable_parameters(m, l0, v30+int32(156), v30+int32(172))
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L34
	} else {
		goto L478
	}
L478:
	;
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2127 = F_transformOptionalSelectInto(m, l0, v2126)
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		goto L34
	} else {
		goto L479
	}
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2127
	F_check_variable_parameters(m, l0, v2127)
	mBase = m.M
	v2131 = m.ExcPending
	if v2131 != 0 {
		goto L34
	} else {
		goto L480
	}
L480:
	;
	goto L456
L481:
	;
	v2213 = F_transformStmt(m, l0, v2190)
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L34
	} else {
		goto L494
	}
L482:
	;
	v2190 = v2132
	goto L481
L483:
	;
	goto L484
L484:
	;
	v2139 = v2132
	goto L486
L485:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v2168)+8))
	if v2169 == int32(0) {
		goto L490
	} else {
		goto L491
	}
L486:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+68))
	if v2163 == int32(0) {
		v2168 = v2139
		goto L485
	} else {
		goto L488
	}
L487:
	;
	v2168 = int32(0)
	goto L485
L488:
	;
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+76))
	if v2166 != 0 {
		v2139 = v2166
		goto L486
	} else {
		goto L489
	}
L489:
	;
	goto L487
L490:
	;
	v2190 = v2132
	goto L481
L491:
	;
	goto L492
L492:
	;
	v2173 = F_palloc0(m, int32(20))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L34
	} else {
		goto L493
	}
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+4)) = v2132
	*(*int32)(unsafe.Add(mBase, uint32(v2173))) = int32(242)
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v2168)+8))
	v2179 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2173)+16)) = uint8(v2179)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+12)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+8)) = v2178
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+8)) = int32(0)
	v2190 = v2173
	goto L481
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2213
	goto L456
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2244
	goto L456
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2275)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2275))) = int64(25769803843)
	v4479 = v2275
	goto L1
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2281
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2284 == int32(23) {
		goto L498
	} else {
		goto L499
	}
L498:
	;
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2281)+42)))
	if v2287 == int32(1) {
		goto L10
	} else {
		goto L501
	}
L499:
	;
	goto L500
L500:
	;
	v2310 = F_palloc0(m, int32(168))
	mBase = m.M
	v2311 = m.ExcPending
	if v2311 != 0 {
		goto L34
	} else {
		goto L508
	}
L501:
	;
	v2292 = F_query_uses_temp_object(m, v2281, v30+int32(156))
	mBase = m.M
	v2293 = m.ExcPending
	if v2293 != 0 {
		goto L34
	} else {
		goto L502
	}
L502:
	;
	if v2292 != 0 {
		goto L9
	} else {
		goto L503
	}
L503:
	;
	v2295 = int32(0)
	v2297 = F_query_tree_walker_impl(m, v2281, int32(531), v2295, v2295)
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L34
	} else {
		goto L504
	}
L504:
	;
	if v2297 != 0 {
		goto L8
	} else {
		goto L505
	}
L505:
	;
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+4))
	v2301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2300)+17)))
	if v2301 == int32(117) {
		goto L7
	} else {
		goto L506
	}
L506:
	;
	v2304 = F_copyObjectImpl(m, v2281)
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L34
	} else {
		goto L507
	}
L507:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2306)+28)) = v2304
	goto L500
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2310)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2310))) = int64(25769803843)
	v4479 = v2310
	goto L1
L509:
	;
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+4))
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+36))
	v2395 = F_ParseFuncOrColumn(m, l0, v2391, v2368, v2392, v2370, int32(1), v2394)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L34
	} else {
		goto L517
	}
L510:
	;
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(v2316)+4))
	if v2319 <= int32(0) {
		v2368 = v3
		v2370 = v2315
		goto L509
	} else {
		goto L511
	}
L511:
	;
	v2325 = v3
	v2326 = v3
	goto L512
L512:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(v2316)+12))
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(v2349+v2325<<(uint(int32(2))%32))))
	v2355 = F_transformExpr(m, l0, v2353, int32(41))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L34
	} else {
		goto L514
	}
L513:
	;
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2368 = v2357
	v2370 = v2363
	goto L509
L514:
	;
	v2357 = F_lappend(m, v2326, v2355)
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
		goto L34
	} else {
		goto L515
	}
L515:
	;
	v2360 = v2325 + int32(1)
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v2316)+4))
	if v2360 < v2361 {
		v2325 = v2360
		v2326 = v2357
		goto L512
	} else {
		goto L516
	}
L516:
	;
	goto L513
L517:
	;
	F_assign_expr_collations(m, l0, v2395)
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L34
	} else {
		goto L518
	}
L518:
	;
	v2400 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2395)+4)))
	v2401 = F_SearchSysCache1(m, int32(47), v2400)
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L34
	} else {
		goto L519
	}
L519:
	;
	if v2401 == int32(0) {
		goto L6
	} else {
		goto L520
	}
L520:
	;
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+28))
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+8))
	v2408 = F_expand_function_arguments(m, v2405, int32(1), v2407, v2401)
	mBase = m.M
	v2409 = m.ExcPending
	if v2409 != 0 {
		goto L34
	} else {
		goto L521
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2395)+28)) = v2408
	v2411 = int32(0)
	v2416 = F_SysCacheGetAttr(m, int32(47), v2401, int32(22), v30+int32(156))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L34
	} else {
		goto L522
	}
L522:
	;
	v2418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+156)))
	if v2418 == int32(0) {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v2422 = F_pg_detoast_datum(m, base.I32_wrap_i64(v2416))
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L34
	} else {
		goto L526
	}
L524:
	;
	v2550 = v2411
	goto L525
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v2550
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v2395
	F_ReleaseCatCache(m, v2401)
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L34
	} else {
		goto L557
	}
L526:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+28))
	if v2424 != 0 {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v2424)+4))
	v2427 = v2425
	goto L529
L528:
	;
	v2427 = int32(0)
	goto L529
L529:
	;
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v2422)+4))
	if v2428 != int32(1) {
		goto L5
	} else {
		goto L530
	}
L530:
	;
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(v2422)+16))
	if v2431 != v2427 {
		goto L5
	} else {
		goto L531
	}
L531:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(v2422)+8))
	if v2433 != 0 {
		goto L5
	} else {
		goto L532
	}
L532:
	;
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(v2422)+12))
	if v2434 != int32(18) {
		goto L5
	} else {
		goto L533
	}
L533:
	;
	if v2424 == int32(0) {
		goto L535
	} else {
		goto L536
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2395)+28)) = v2516
	v2550 = v2522
	goto L525
L535:
	;
	v2516 = int32(0)
	v2522 = v2411
	goto L534
L536:
	;
	goto L537
L537:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v2424)+4))
	if v2440 <= int32(0) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v2516 = int32(0)
	v2522 = v2411
	goto L534
L539:
	;
	goto L540
L540:
	;
	v2446 = int32(0)
	v2448 = v2446
	v2451 = v2446
	v2454 = v2411
	goto L541
L541:
	;
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v2424)+12))
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2475+v2451<<(uint(int32(2))%32))))
	v2480 = v2451 + (v2422 + int32(24))
	v2481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2480))))
	switch v2481 - int32(98) {
	case 0:
		goto L546
	default:
		goto L545
	case 7, 20:
		goto L544
	case 13:
		goto L547
	}
L542:
	;
	v2516 = v2510
	v2522 = v2511
	goto L534
L543:
	;
	v2513 = v2451 + int32(1)
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v2424)+4))
	if v2513 < v2514 {
		v2448 = v2510
		v2451 = v2513
		v2454 = v2511
		goto L541
	} else {
		goto L556
	}
L544:
	;
	v2508 = F_lappend(m, v2448, v2479)
	mBase = m.M
	v2509 = m.ExcPending
	if v2509 != 0 {
		goto L34
	} else {
		goto L555
	}
L545:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2495 = m.ExcPending
	if v2495 != 0 {
		goto L34
	} else {
		goto L552
	}
L546:
	;
	v2486 = F_lappend(m, v2448, v2479)
	mBase = m.M
	v2487 = m.ExcPending
	if v2487 != 0 {
		goto L34
	} else {
		goto L549
	}
L547:
	;
	v2484 = F_lappend(m, v2454, v2479)
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L34
	} else {
		goto L548
	}
L548:
	;
	v2510 = v2448
	v2511 = v2484
	goto L543
L549:
	;
	v2488 = F_copyObjectImpl(m, v2479)
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		goto L34
	} else {
		goto L550
	}
L550:
	;
	v2490 = F_lappend(m, v2454, v2488)
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		goto L34
	} else {
		goto L551
	}
L551:
	;
	v2510 = v2486
	v2511 = v2490
	goto L543
L552:
	;
	v2496 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2480))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+128)) = v2496
	F_errmsg_internal(m, int32(_a_F_transformStmt_24), v30+int32(128))
	mBase = m.M
	v2502 = m.ExcPending
	if v2502 != 0 {
		goto L34
	} else {
		goto L553
	}
L553:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3345), int32(_a_F_transformStmt_25))
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L34
	} else {
		goto L554
	}
L554:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L555:
	;
	v2510 = v2508
	v2511 = v2454
	goto L543
L556:
	;
	goto L542
L557:
	;
	v2576 = F_palloc0(m, int32(168))
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L34
	} else {
		goto L558
	}
L558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2576)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2576))) = int64(25769803843)
	v4479 = v2576
	goto L1
L559:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2582)+28)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v2582))) = int64(25769803843)
	v4479 = v2582
	goto L1
L560:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L34
	} else {
		goto L561
	}
L561:
	;
	F_errmsg(m, int32(_a_F_transformStmt_26), int32(0))
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		goto L34
	} else {
		goto L562
	}
L562:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(597), int32(_a_F_transformStmt_27))
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L34
	} else {
		goto L563
	}
L563:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L564:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L34
	} else {
		goto L565
	}
L565:
	;
	F_errmsg(m, int32(_a_F_transformStmt_26), int32(0))
	mBase = m.M
	v2613 = m.ExcPending
	if v2613 != 0 {
		goto L34
	} else {
		goto L566
	}
L566:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(2557), int32(_a_F_transformStmt_28))
	mBase = m.M
	v2618 = m.ExcPending
	if v2618 != 0 {
		goto L34
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L34
	} else {
		goto L569
	}
L569:
	;
	F_errmsg(m, int32(_a_F_transformStmt_29), int32(0))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L34
	} else {
		goto L570
	}
L570:
	;
	v2630 = F_exprLocation(m, v1311)
	mBase = m.M
	F_parser_errposition(m, l0, v2630)
	mBase = m.M
	v2632 = m.ExcPending
	if v2632 != 0 {
		goto L34
	} else {
		goto L571
	}
L571:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(1645), int32(_a_F_transformStmt_30))
	mBase = m.M
	v2637 = m.ExcPending
	if v2637 != 0 {
		goto L34
	} else {
		goto L572
	}
L572:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L573:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_31), int32(0))
	mBase = m.M
	v2645 = m.ExcPending
	if v2645 != 0 {
		goto L34
	} else {
		goto L574
	}
L574:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(2866), int32(_a_F_transformStmt_32))
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L34
	} else {
		goto L575
	}
L575:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L576:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v2657 = m.ExcPending
	if v2657 != 0 {
		goto L34
	} else {
		goto L577
	}
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = int32(_a_F_transformStmt_33)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = int32(_a_F_transformStmt_34)
	F_errmsg(m, int32(_a_F_transformStmt_35), v30+int32(16))
	mBase = m.M
	v2666 = m.ExcPending
	if v2666 != 0 {
		goto L34
	} else {
		goto L578
	}
L578:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3033), int32(_a_F_transformStmt_36))
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L34
	} else {
		goto L579
	}
L579:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L580:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v2678 = m.ExcPending
	if v2678 != 0 {
		goto L34
	} else {
		goto L581
	}
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = int32(_a_F_transformStmt_37)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = int32(_a_F_transformStmt_38)
	F_errmsg(m, int32(_a_F_transformStmt_35), v30+int32(32))
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		goto L34
	} else {
		goto L582
	}
L582:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3041), int32(_a_F_transformStmt_36))
	mBase = m.M
	v2692 = m.ExcPending
	if v2692 != 0 {
		goto L34
	} else {
		goto L583
	}
L583:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L584:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_39), int32(0))
	mBase = m.M
	v2700 = m.ExcPending
	if v2700 != 0 {
		goto L34
	} else {
		goto L585
	}
L585:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3050), int32(_a_F_transformStmt_36))
	mBase = m.M
	v2705 = m.ExcPending
	if v2705 != 0 {
		goto L34
	} else {
		goto L586
	}
L586:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L587:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2712 = m.ExcPending
	if v2712 != 0 {
		goto L34
	} else {
		goto L588
	}
L588:
	;
	F_errmsg(m, int32(_a_F_transformStmt_40), int32(0))
	mBase = m.M
	v2716 = m.ExcPending
	if v2716 != 0 {
		goto L34
	} else {
		goto L589
	}
L589:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3060), int32(_a_F_transformStmt_36))
	mBase = m.M
	v2721 = m.ExcPending
	if v2721 != 0 {
		goto L34
	} else {
		goto L590
	}
L590:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L591:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2728 = m.ExcPending
	if v2728 != 0 {
		goto L34
	} else {
		goto L592
	}
L592:
	;
	v2729 = *(*int32)(unsafe.Add(mBase, uint32(v2007)+140))
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2729)+12))
	v2731 = *(*int32)(unsafe.Add(mBase, uint32(v2730)))
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(v2731)+8))
	v2734 = v2732 - int32(1)
	if base.Ui32(v2734) <= base.Ui32(int32(3)) {
		goto L594
	} else {
		goto L595
	}
L593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v2741
	F_errmsg(m, int32(_a_F_transformStmt_41), v30+int32(80))
	mBase = m.M
	v2747 = m.ExcPending
	if v2747 != 0 {
		goto L34
	} else {
		goto L597
	}
L594:
	;
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2734<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v2741 = v2739
	goto L596
L595:
	;
	v2741 = int32(_a_F_transformStmt_21)
	goto L596
L596:
	;
	goto L593
L597:
	;
	v2750 = F_errdetail(m, int32(_a_F_transformStmt_42), int32(0))
	mBase = m.M
	v2751 = m.ExcPending
	if v2751 != 0 {
		goto L34
	} else {
		goto L598
	}
L598:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3071), int32(_a_F_transformStmt_36))
	mBase = m.M
	v2756 = m.ExcPending
	if v2756 != 0 {
		goto L34
	} else {
		goto L599
	}
L599:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L600:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2763 = m.ExcPending
	if v2763 != 0 {
		goto L34
	} else {
		goto L601
	}
L601:
	;
	v2764 = *(*int32)(unsafe.Add(mBase, uint32(v2007)+140))
	v2765 = *(*int32)(unsafe.Add(mBase, uint32(v2764)+12))
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v2765)))
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(v2766)+8))
	v2769 = v2767 - int32(1)
	if base.Ui32(v2769) <= base.Ui32(int32(3)) {
		goto L603
	} else {
		goto L604
	}
L602:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v2776
	F_errmsg(m, int32(_a_F_transformStmt_43), v30-int32(-64))
	mBase = m.M
	v2782 = m.ExcPending
	if v2782 != 0 {
		goto L34
	} else {
		goto L606
	}
L603:
	;
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(v2769<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v2776 = v2774
	goto L605
L604:
	;
	v2776 = int32(_a_F_transformStmt_21)
	goto L605
L605:
	;
	goto L602
L606:
	;
	v2785 = F_errdetail(m, int32(_a_F_transformStmt_44), int32(0))
	mBase = m.M
	v2786 = m.ExcPending
	if v2786 != 0 {
		goto L34
	} else {
		goto L607
	}
L607:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3082), int32(_a_F_transformStmt_36))
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		goto L34
	} else {
		goto L608
	}
L608:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L609:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v2798 = m.ExcPending
	if v2798 != 0 {
		goto L34
	} else {
		goto L610
	}
L610:
	;
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(v2007)+140))
	v2800 = *(*int32)(unsafe.Add(mBase, uint32(v2799)+12))
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(v2800)))
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(v2801)+8))
	v2804 = v2802 - int32(1)
	if base.Ui32(v2804) <= base.Ui32(int32(3)) {
		goto L612
	} else {
		goto L613
	}
L611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v2811
	F_errmsg(m, int32(_a_F_transformStmt_45), v30+int32(48))
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L34
	} else {
		goto L615
	}
L612:
	;
	v2809 = *(*int32)(unsafe.Add(mBase, uint32(v2804<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v2811 = v2809
	goto L614
L613:
	;
	v2811 = int32(_a_F_transformStmt_21)
	goto L614
L614:
	;
	goto L611
L615:
	;
	v2820 = F_errdetail(m, int32(_a_F_transformStmt_46), int32(0))
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L34
	} else {
		goto L616
	}
L616:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3093), int32(_a_F_transformStmt_36))
	mBase = m.M
	v2826 = m.ExcPending
	if v2826 != 0 {
		goto L34
	} else {
		goto L617
	}
L617:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L618:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L34
	} else {
		goto L619
	}
L619:
	;
	F_errmsg(m, int32(_a_F_transformStmt_47), int32(0))
	mBase = m.M
	v2837 = m.ExcPending
	if v2837 != 0 {
		goto L34
	} else {
		goto L620
	}
L620:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3189), int32(_a_F_transformStmt_48))
	mBase = m.M
	v2842 = m.ExcPending
	if v2842 != 0 {
		goto L34
	} else {
		goto L621
	}
L621:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L622:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2849 = m.ExcPending
	if v2849 != 0 {
		goto L34
	} else {
		goto L623
	}
L623:
	;
	F_errmsg(m, int32(_a_F_transformStmt_49), int32(0))
	mBase = m.M
	v2853 = m.ExcPending
	if v2853 != 0 {
		goto L34
	} else {
		goto L624
	}
L624:
	;
	v2857 = F_getObjectDescription(m, v30+int32(156), int32(0))
	mBase = m.M
	v2858 = m.ExcPending
	if v2858 != 0 {
		goto L34
	} else {
		goto L625
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v2857
	v2863 = F_errdetail(m, int32(_a_F_transformStmt_50), v30+int32(96))
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
		goto L34
	} else {
		goto L626
	}
L626:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3201), int32(_a_F_transformStmt_48))
	mBase = m.M
	v2869 = m.ExcPending
	if v2869 != 0 {
		goto L34
	} else {
		goto L627
	}
L627:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L628:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2876 = m.ExcPending
	if v2876 != 0 {
		goto L34
	} else {
		goto L629
	}
L629:
	;
	F_errmsg(m, int32(_a_F_transformStmt_51), int32(0))
	mBase = m.M
	v2880 = m.ExcPending
	if v2880 != 0 {
		goto L34
	} else {
		goto L630
	}
L630:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3211), int32(_a_F_transformStmt_48))
	mBase = m.M
	v2885 = m.ExcPending
	if v2885 != 0 {
		goto L34
	} else {
		goto L631
	}
L631:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L632:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2892 = m.ExcPending
	if v2892 != 0 {
		goto L34
	} else {
		goto L633
	}
L633:
	;
	F_errmsg(m, int32(_a_F_transformStmt_52), int32(0))
	mBase = m.M
	v2896 = m.ExcPending
	if v2896 != 0 {
		goto L34
	} else {
		goto L634
	}
L634:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3223), int32(_a_F_transformStmt_48))
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L34
	} else {
		goto L635
	}
L635:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L636:
	;
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v2906
	F_errmsg_internal(m, int32(_a_F_transformStmt_53), v30+int32(112))
	mBase = m.M
	v2912 = m.ExcPending
	if v2912 != 0 {
		goto L34
	} else {
		goto L637
	}
L637:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3284), int32(_a_F_transformStmt_25))
	mBase = m.M
	v2917 = m.ExcPending
	if v2917 != 0 {
		goto L34
	} else {
		goto L638
	}
L638:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+144)) = v2427
	F_errmsg_internal(m, int32(_a_F_transformStmt_54), v30+int32(144))
	mBase = m.M
	v2927 = m.ExcPending
	if v2927 != 0 {
		goto L34
	} else {
		goto L640
	}
L640:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(3320), int32(_a_F_transformStmt_25))
	mBase = m.M
	v2932 = m.ExcPending
	if v2932 != 0 {
		goto L34
	} else {
		goto L641
	}
L641:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L642:
	;
	v2947 = v3
	v2949 = int32(0)
	v2950 = v3
	v2952 = v3
	goto L643
L643:
	;
	v2963 = int32(0)
	v2966 = v1326 + v2949<<(uint(int32(2))%32)
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v2966)))
	v2970 = F_select_common_type(m, l0, v2967, int32(_a_F_transformStmt_55), v2963)
	mBase = m.M
	v2971 = m.ExcPending
	if v2971 != 0 {
		goto L34
	} else {
		goto L645
	}
L644:
	;
	v3068 = int32(0)
	goto L661
L645:
	;
	v2972 = *(*int32)(unsafe.Add(mBase, uint32(v2966)))
	if v2972 == int32(0) {
		v3022 = v2963
		goto L646
	} else {
		goto L647
	}
L646:
	;
	v3046 = F_select_common_typmod(m, v3022, v2970)
	mBase = m.M
	v3047 = m.ExcPending
	if v3047 != 0 {
		goto L34
	} else {
		goto L655
	}
L647:
	;
	v2975 = *(*int32)(unsafe.Add(mBase, uint32(v2972)+4))
	if v2975 <= int32(0) {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	v3022 = v2972
	goto L646
L649:
	;
	goto L650
L650:
	;
	v2981 = v2963
	goto L651
L651:
	;
	v3005 = *(*int32)(unsafe.Add(mBase, uint32(v2972)+12))
	v3008 = v3005 + v2981<<(uint(int32(2))%32)
	v3009 = *(*int32)(unsafe.Add(mBase, uint32(v3008)))
	v3011 = F_coerce_to_common_type(m, l0, v3009, v2970, int32(_a_F_transformStmt_55))
	mBase = m.M
	v3012 = m.ExcPending
	if v3012 != 0 {
		goto L34
	} else {
		goto L653
	}
L652:
	;
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(v2966)))
	v3022 = v3018
	goto L646
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3008))) = v3011
	v3015 = v2981 + int32(1)
	v3016 = *(*int32)(unsafe.Add(mBase, uint32(v2972)+4))
	if v3015 < v3016 {
		v2981 = v3015
		goto L651
	} else {
		goto L654
	}
L654:
	;
	goto L652
L655:
	;
	v3048 = *(*int32)(unsafe.Add(mBase, uint32(v2966)))
	v3050 = F_select_common_collation(m, l0, v3048, int32(1))
	mBase = m.M
	v3051 = m.ExcPending
	if v3051 != 0 {
		goto L34
	} else {
		goto L656
	}
L656:
	;
	v3052 = F_lappend_oid(m, v2950, v2970)
	mBase = m.M
	v3053 = m.ExcPending
	if v3053 != 0 {
		goto L34
	} else {
		goto L657
	}
L657:
	;
	v3054 = F_lappend_int(m, v2952, v3046)
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L34
	} else {
		goto L658
	}
L658:
	;
	v3056 = F_lappend_oid(m, v2947, v3050)
	mBase = m.M
	v3057 = m.ExcPending
	if v3057 != 0 {
		goto L34
	} else {
		goto L659
	}
L659:
	;
	v3059 = v2949 + int32(1)
	if v3059 != v1327 {
		v2947 = v3056
		v2949 = v3059
		v2950 = v3052
		v2952 = v3054
		goto L643
	} else {
		goto L660
	}
L660:
	;
	goto L644
L661:
	;
	v3091 = v1326 + v3068<<(uint(int32(2))%32)
	v3092 = *(*int32)(unsafe.Add(mBase, uint32(v3091)))
	v3097 = int32(0)
	goto L663
L663:
	;
	v3121 = int32(0)
	if v3092 == v3121 {
		v3130 = v3121
		goto L665
	} else {
		goto L666
	}
L665:
	;
	if v1405 == int32(0) {
		goto L669
	} else {
		goto L670
	}
L666:
	;
	v3124 = *(*int32)(unsafe.Add(mBase, uint32(v3092)+4))
	if v3124 <= v3097 {
		v3130 = v3121
		goto L665
	} else {
		goto L667
	}
L667:
	;
	v3126 = *(*int32)(unsafe.Add(mBase, uint32(v3092)+12))
	v3130 = v3126 + v3097<<(uint(int32(2))%32)
	goto L665
L668:
	;
	v3148 = v3138 + v3097<<(uint(int32(2))%32)
	v3149 = *(*int32)(unsafe.Add(mBase, uint32(v3148)))
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v3130)))
	v3151 = F_lappend(m, v3149, v3150)
	mBase = m.M
	v3152 = m.ExcPending
	if v3152 != 0 {
		goto L34
	} else {
		goto L675
	}
L669:
	;
	v3140 = *(*int32)(unsafe.Add(mBase, uint32(v3091)))
	F_list_free(m, v3140)
	mBase = m.M
	v3142 = m.ExcPending
	if v3142 != 0 {
		goto L34
	} else {
		goto L673
	}
L670:
	;
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+4))
	if base.B2i32(v3130 == int32(0))|base.B2i32(v3135 <= v3097) != 0 {
		goto L669
	} else {
		goto L671
	}
L671:
	;
	v3138 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+12))
	if v3138 != 0 {
		goto L668
	} else {
		goto L672
	}
L672:
	;
	goto L669
L673:
	;
	v3144 = v3068 + int32(1)
	if v3144 != v1327 {
		v3068 = v3144
		goto L661
	} else {
		goto L674
	}
L674:
	;
	v3167 = v3056
	v3168 = v1405
	v3170 = v3052
	v3172 = v3054
	goto L3
L675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3148))) = v3151
	v3097 = v3097 + int32(1)
	goto L663
L676:
	;
	v3185 = F_contain_vars_of_level(m, v3168, int32(0))
	mBase = m.M
	v3186 = m.ExcPending
	if v3186 != 0 {
		goto L34
	} else {
		goto L679
	}
L677:
	;
	v3188 = int32(0)
	goto L678
L678:
	;
	v3189 = F_addRangeTableEntryForValues(m, l0, v3168, v3170, v3172, v3167, v3188)
	mBase = m.M
	v3190 = m.ExcPending
	if v3190 != 0 {
		goto L34
	} else {
		goto L680
	}
L679:
	;
	v3188 = v3185
	goto L678
L680:
	;
	v3191 = int32(1)
	F_addNSItemToQuery(m, l0, v3189, v3191, v3191, v3191)
	mBase = m.M
	v3195 = m.ExcPending
	if v3195 != 0 {
		goto L34
	} else {
		goto L681
	}
L681:
	;
	v3198 = F_expandNSItemAttrs(m, l0, v3189, int32(0), int32(-1))
	mBase = m.M
	v3199 = m.ExcPending
	if v3199 != 0 {
		goto L34
	} else {
		goto L682
	}
L682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+76)) = v3198
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v3205 = F_transformSortClause(m, l0, v3201, v1257+int32(76), int32(0))
	mBase = m.M
	v3206 = m.ExcPending
	if v3206 != 0 {
		goto L34
	} else {
		goto L683
	}
L683:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+124)) = v3205
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v3211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v3212 = F_transformLimitClause(m, l0, v3208, int32(23), int32(_a_F_transformStmt_13), v3211)
	mBase = m.M
	v3213 = m.ExcPending
	if v3213 != 0 {
		goto L34
	} else {
		goto L684
	}
L684:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+128)) = v3212
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v3219 = F_transformLimitClause(m, l0, v3215, int32(22), int32(_a_F_transformStmt_14), v3218)
	mBase = m.M
	v3220 = m.ExcPending
	if v3220 != 0 {
		goto L34
	} else {
		goto L685
	}
L685:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+132)) = v3219
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+136)) = v3222
	v3224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v3224 == int32(0) {
		goto L686
	} else {
		goto L687
	}
L686:
	;
	v3227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+52)) = v3227
	v3229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+56)) = v3229
	v3231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3233 = F_makeFromExpr(m, v3231, int32(0))
	mBase = m.M
	v3234 = m.ExcPending
	if v3234 != 0 {
		goto L34
	} else {
		goto L689
	}
L687:
	;
	goto L688
L688:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3243 = m.ExcPending
	if v3243 != 0 {
		goto L34
	} else {
		goto L691
	}
L689:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+60)) = v3233
	v3236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1257)+39)) = uint8(v3236)
	F_assign_query_collations(m, l0, v1257)
	mBase = m.M
	v3239 = m.ExcPending
	if v3239 != 0 {
		goto L34
	} else {
		goto L690
	}
L690:
	;
	v4479 = v1257
	goto L1
L691:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3246 = m.ExcPending
	if v3246 != 0 {
		goto L34
	} else {
		goto L692
	}
L692:
	;
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v3248 = *(*int32)(unsafe.Add(mBase, uint32(v3247)+12))
	v3249 = *(*int32)(unsafe.Add(mBase, uint32(v3248)))
	v3250 = *(*int32)(unsafe.Add(mBase, uint32(v3249)+8))
	v3252 = v3250 - int32(1)
	if base.Ui32(v3252) <= base.Ui32(int32(3)) {
		goto L694
	} else {
		goto L695
	}
L693:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v3259
	F_errmsg(m, int32(_a_F_transformStmt_56), v30)
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		goto L34
	} else {
		goto L697
	}
L694:
	;
	v3257 = *(*int32)(unsafe.Add(mBase, uint32(v3252<<(uint(int32(2))%32))+uint32(_c_F_transformStmt[1])))
	v3259 = v3257
	goto L696
L695:
	;
	v3259 = int32(_a_F_transformStmt_21)
	goto L696
L696:
	;
	goto L693
L697:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(1768), int32(_a_F_transformStmt_30))
	mBase = m.M
	v3268 = m.ExcPending
	if v3268 != 0 {
		goto L34
	} else {
		goto L698
	}
L698:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L699:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v3276
	v3279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v3282 = F_checkInsertTargets(m, l0, v3279, v30+int32(152))
	mBase = m.M
	v3283 = m.ExcPending
	if v3283 != 0 {
		goto L34
	} else {
		goto L700
	}
L700:
	;
	if v42 != 0 {
		goto L701
	} else {
		goto L702
	}
L701:
	;
	if v3269 != 0 {
		goto L707
	} else {
		goto L708
	}
L702:
	;
	v3641 = v3
	goto L703
L703:
	;
	v3660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3661 = *(*int32)(unsafe.Add(mBase, uint32(v3660)+12))
	v3662 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v3662
	v3664 = *(*int32)(unsafe.Add(mBase, uint32(v30)+152))
	v3669 = v3662
	goto L787
L704:
	;
	v3628 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v3629 = *(*int32)(unsafe.Add(mBase, uint32(v30)+152))
	v3631 = F_transformInsertRow(m, l0, v3607, v3628, v3282, v3629, int32(0))
	mBase = m.M
	v3632 = m.ExcPending
	if v3632 != 0 {
		goto L34
	} else {
		goto L786
	}
L705:
	;
	v3491 = *(*int32)(unsafe.Add(mBase, uint32(v3466)+12))
	v3492 = *(*int32)(unsafe.Add(mBase, uint32(v3491)))
	if v3492 == int32(0) {
		goto L765
	} else {
		goto L766
	}
L706:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3481 = m.ExcPending
	if v3481 != 0 {
		goto L34
	} else {
		goto L759
	}
L707:
	;
	v3284 = F_make_parsestate(m, l0)
	mBase = m.M
	v3285 = m.ExcPending
	if v3285 != 0 {
		goto L34
	} else {
		goto L710
	}
L708:
	;
	goto L709
L709:
	;
	v3386 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v3386 == int32(0) {
		goto L733
	} else {
		goto L734
	}
L710:
	;
	v3286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3284)+81)) = uint8(v3286)
	*(*int32)(unsafe.Add(mBase, uint32(v3284)+28)) = v3271
	*(*int64)(unsafe.Add(mBase, uint32(v3284)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3284)+12)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v3284)+8)) = v3272
	v3293 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3294 = F_transformStmt(m, v3284, v3293)
	mBase = m.M
	v3295 = m.ExcPending
	if v3295 != 0 {
		goto L34
	} else {
		goto L711
	}
L711:
	;
	F_free_parsestate(m, v3284)
	mBase = m.M
	v3297 = m.ExcPending
	if v3297 != 0 {
		goto L34
	} else {
		goto L712
	}
L712:
	;
	v3298 = *(*int32)(unsafe.Add(mBase, uint32(v3294)))
	if v3298 != int32(67) {
		goto L706
	} else {
		goto L713
	}
L713:
	;
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v3294)+4))
	if v3301 != int32(1) {
		goto L706
	} else {
		goto L714
	}
L714:
	;
	v3304 = int32(0)
	v3308 = F_addRangeTableEntryForSubquery(m, l0, v3294, v3304, v3304, v3304)
	mBase = m.M
	v3309 = m.ExcPending
	if v3309 != 0 {
		goto L34
	} else {
		goto L715
	}
L715:
	;
	v3311 = int32(0)
	F_addNSItemToQuery(m, l0, v3308, int32(1), v3311, v3311)
	mBase = m.M
	v3314 = m.ExcPending
	if v3314 != 0 {
		goto L34
	} else {
		goto L716
	}
L716:
	;
	v3315 = *(*int32)(unsafe.Add(mBase, uint32(v3294)+76))
	if v3315 == int32(0) {
		v3607 = v3304
		goto L704
	} else {
		goto L717
	}
L717:
	;
	v3318 = *(*int32)(unsafe.Add(mBase, uint32(v3315)+4))
	if v3318 <= int32(0) {
		v3607 = v3304
		goto L704
	} else {
		goto L718
	}
L718:
	;
	v3325 = int32(0)
	v3328 = v3304
	goto L719
L719:
	;
	v3349 = *(*int32)(unsafe.Add(mBase, uint32(v3315)+12))
	v3353 = *(*int32)(unsafe.Add(mBase, uint32(v3349+v3325<<(uint(int32(2))%32))))
	v3354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3353)+26)))
	if v3354 == int32(0) {
		goto L721
	} else {
		goto L722
	}
L720:
	;
	v3607 = v3379
	goto L704
L721:
	;
	v3357 = *(*int32)(unsafe.Add(mBase, uint32(v3353)+4))
	if v3357 == int32(0) {
		goto L725
	} else {
		goto L726
	}
L722:
	;
	v3379 = v3328
	goto L723
L723:
	;
	v3383 = v3325 + int32(1)
	v3384 = *(*int32)(unsafe.Add(mBase, uint32(v3315)+4))
	if v3383 < v3384 {
		v3325 = v3383
		v3328 = v3379
		goto L719
	} else {
		goto L732
	}
L724:
	;
	v3377 = F_lappend(m, v3328, v3376)
	mBase = m.M
	v3378 = m.ExcPending
	if v3378 != 0 {
		goto L34
	} else {
		goto L731
	}
L725:
	;
	v3370 = *(*int32)(unsafe.Add(mBase, uint32(v3308)+8))
	v3371 = F_makeVarFromTargetEntry(m, v3370, v3353)
	mBase = m.M
	v3372 = m.ExcPending
	if v3372 != 0 {
		goto L34
	} else {
		goto L730
	}
L726:
	;
	v3360 = *(*int32)(unsafe.Add(mBase, uint32(v3357)))
	if base.Ui32(int32(1)) < base.Ui32(v3360-int32(7)) {
		goto L725
	} else {
		goto L727
	}
L727:
	;
	v3365 = F_exprType(m, v3357)
	mBase = m.M
	v3366 = m.ExcPending
	if v3366 != 0 {
		goto L34
	} else {
		goto L728
	}
L728:
	;
	if v3365 != int32(705) {
		goto L725
	} else {
		goto L729
	}
L729:
	;
	v3369 = *(*int32)(unsafe.Add(mBase, uint32(v3353)+4))
	v3376 = v3369
	goto L724
L730:
	;
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(v3353)+4))
	v3374 = F_exprLocation(m, v3373)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3371)+44)) = v3374
	v3376 = v3371
	goto L724
L731:
	;
	v3379 = v3377
	goto L723
L732:
	;
	goto L720
L733:
	;
	v3472 = *(*int32)(unsafe.Add(mBase, uint32(v3386)+12))
	v3473 = *(*int32)(unsafe.Add(mBase, uint32(v3472)))
	v3476 = F_transformExpressionList(m, l0, v3473, int32(27), int32(1))
	mBase = m.M
	v3477 = m.ExcPending
	if v3477 != 0 {
		goto L34
	} else {
		goto L758
	}
L734:
	;
	v3389 = *(*int32)(unsafe.Add(mBase, uint32(v3386)+4))
	if v3389 < int32(2) {
		goto L733
	} else {
		goto L735
	}
L735:
	;
	v3393 = int32(0)
	v3399 = v3393
	v3401 = v3393
	v3403 = int32(-1)
	goto L736
L736:
	;
	v3422 = *(*int32)(unsafe.Add(mBase, uint32(v3386)+12))
	v3426 = *(*int32)(unsafe.Add(mBase, uint32(v3422+v3399<<(uint(int32(2))%32))))
	v3429 = F_transformExpressionList(m, l0, v3426, int32(26), int32(1))
	mBase = m.M
	v3430 = m.ExcPending
	if v3430 != 0 {
		goto L34
	} else {
		goto L738
	}
L737:
	;
	goto L705
L738:
	;
	if v3403 < int32(0) {
		goto L741
	} else {
		goto L742
	}
L739:
	;
	v3459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v3460 = *(*int32)(unsafe.Add(mBase, uint32(v30)+152))
	v3462 = F_transformInsertRow(m, l0, v3429, v3459, v3282, v3460, int32(1))
	mBase = m.M
	v3463 = m.ExcPending
	if v3463 != 0 {
		goto L34
	} else {
		goto L754
	}
L740:
	;
	v3457 = *(*int32)(unsafe.Add(mBase, uint32(v3429)+4))
	v3458 = v3457
	goto L739
L741:
	;
	if v3429 != 0 {
		goto L740
	} else {
		goto L744
	}
L742:
	;
	goto L743
L743:
	;
	if v3429 != 0 {
		goto L745
	} else {
		goto L746
	}
L744:
	;
	v3458 = int32(0)
	goto L739
L745:
	;
	v3434 = *(*int32)(unsafe.Add(mBase, uint32(v3429)+4))
	v3436 = v3434
	goto L747
L746:
	;
	v3436 = int32(0)
	goto L747
L747:
	;
	if v3436 == v3403 {
		v3458 = v3403
		goto L739
	} else {
		goto L748
	}
L748:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3441 = m.ExcPending
	if v3441 != 0 {
		goto L34
	} else {
		goto L749
	}
L749:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3444 = m.ExcPending
	if v3444 != 0 {
		goto L34
	} else {
		goto L750
	}
L750:
	;
	F_errmsg(m, int32(_a_F_transformStmt_29), int32(0))
	mBase = m.M
	v3448 = m.ExcPending
	if v3448 != 0 {
		goto L34
	} else {
		goto L751
	}
L751:
	;
	v3449 = F_exprLocation(m, v3429)
	mBase = m.M
	F_parser_errposition(m, l0, v3449)
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		goto L34
	} else {
		goto L752
	}
L752:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(899), int32(_a_F_transformStmt_57))
	mBase = m.M
	v3456 = m.ExcPending
	if v3456 != 0 {
		goto L34
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
	F_assign_list_collations(m, l0, v3462)
	mBase = m.M
	v3465 = m.ExcPending
	if v3465 != 0 {
		goto L34
	} else {
		goto L755
	}
L755:
	;
	v3466 = F_lappend(m, v3401, v3462)
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		goto L34
	} else {
		goto L756
	}
L756:
	;
	v3469 = v3399 + int32(1)
	v3470 = *(*int32)(unsafe.Add(mBase, uint32(v3386)+4))
	if v3469 < v3470 {
		v3399 = v3469
		v3401 = v3466
		v3403 = v3458
		goto L736
	} else {
		goto L757
	}
L757:
	;
	goto L737
L758:
	;
	v3607 = v3476
	goto L704
L759:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_58), int32(0))
	mBase = m.M
	v3485 = m.ExcPending
	if v3485 != 0 {
		goto L34
	} else {
		goto L760
	}
L760:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(800), int32(_a_F_transformStmt_57))
	mBase = m.M
	v3490 = m.ExcPending
	if v3490 != 0 {
		goto L34
	} else {
		goto L761
	}
L761:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L762:
	;
	v3580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3580 != 0 {
		goto L778
	} else {
		goto L779
	}
L763:
	;
	v3503 = int32(0)
	v3509 = v3503
	v3510 = v3496
	v3513 = v3503
	v3514 = v3503
	goto L769
L764:
	;
	v3501 = int32(0)
	v3557 = v3500
	v3560 = v3501
	v3561 = v3501
	goto L762
L765:
	;
	v3500 = int32(0)
	goto L764
L766:
	;
	goto L767
L767:
	;
	v3496 = int32(0)
	v3497 = *(*int32)(unsafe.Add(mBase, uint32(v3492)+4))
	if v3496 < v3497 {
		goto L763
	} else {
		goto L768
	}
L768:
	;
	v3500 = v3496
	goto L764
L769:
	;
	v3533 = *(*int32)(unsafe.Add(mBase, uint32(v3492)+12))
	v3537 = *(*int32)(unsafe.Add(mBase, uint32(v3533+v3509<<(uint(int32(2))%32))))
	v3538 = F_exprType(m, v3537)
	mBase = m.M
	v3539 = m.ExcPending
	if v3539 != 0 {
		goto L34
	} else {
		goto L771
	}
L770:
	;
	v3557 = v3540
	v3560 = v3544
	v3561 = v3547
	goto L762
L771:
	;
	v3540 = F_lappend_oid(m, v3510, v3538)
	mBase = m.M
	v3541 = m.ExcPending
	if v3541 != 0 {
		goto L34
	} else {
		goto L772
	}
L772:
	;
	v3542 = F_exprTypmod(m, v3537)
	mBase = m.M
	v3543 = m.ExcPending
	if v3543 != 0 {
		goto L34
	} else {
		goto L773
	}
L773:
	;
	v3544 = F_lappend_int(m, v3513, v3542)
	mBase = m.M
	v3545 = m.ExcPending
	if v3545 != 0 {
		goto L34
	} else {
		goto L774
	}
L774:
	;
	v3547 = F_lappend_oid(m, v3514, int32(0))
	mBase = m.M
	v3548 = m.ExcPending
	if v3548 != 0 {
		goto L34
	} else {
		goto L775
	}
L775:
	;
	v3550 = v3509 + int32(1)
	v3551 = *(*int32)(unsafe.Add(mBase, uint32(v3492)+4))
	if v3550 < v3551 {
		v3509 = v3550
		v3510 = v3540
		v3513 = v3544
		v3514 = v3547
		goto L769
	} else {
		goto L776
	}
L776:
	;
	goto L770
L777:
	;
	v3589 = F_addRangeTableEntryForValues(m, l0, v3466, v3557, v3560, v3561, v3588)
	mBase = m.M
	v3590 = m.ExcPending
	if v3590 != 0 {
		goto L34
	} else {
		goto L783
	}
L778:
	;
	v3582 = *(*int32)(unsafe.Add(mBase, uint32(v3580)+4))
	if v3582 == int32(1) {
		v3588 = int32(0)
		goto L777
	} else {
		goto L781
	}
L779:
	;
	goto L780
L780:
	;
	v3586 = F_contain_vars_of_level(m, v3466, int32(0))
	mBase = m.M
	v3587 = m.ExcPending
	if v3587 != 0 {
		goto L34
	} else {
		goto L782
	}
L781:
	;
	goto L780
L782:
	;
	v3588 = v3586
	goto L777
L783:
	;
	v3592 = int32(0)
	F_addNSItemToQuery(m, l0, v3589, int32(1), v3592, v3592)
	mBase = m.M
	v3595 = m.ExcPending
	if v3595 != 0 {
		goto L34
	} else {
		goto L784
	}
L784:
	;
	v3596 = int32(0)
	v3599 = F_expandNSItemVars(m, l0, v3589, v3596, int32(-1), v3596)
	mBase = m.M
	v3600 = m.ExcPending
	if v3600 != 0 {
		goto L34
	} else {
		goto L785
	}
L785:
	;
	v3607 = v3599
	goto L704
L786:
	;
	v3641 = v3631
	goto L703
L787:
	;
	v3693 = int32(0)
	if v3641 == v3693 {
		v3703 = v3693
		goto L789
	} else {
		goto L790
	}
L789:
	;
	v3704 = int32(0)
	if v3282 == v3704 {
		v3713 = v3704
		goto L792
	} else {
		goto L793
	}
L790:
	;
	v3697 = *(*int32)(unsafe.Add(mBase, uint32(v3641)+4))
	if v3697 <= v3669 {
		v3703 = int32(0)
		goto L789
	} else {
		goto L791
	}
L791:
	;
	v3699 = *(*int32)(unsafe.Add(mBase, uint32(v3641)+12))
	v3703 = v3699 + v3669<<(uint(int32(2))%32)
	goto L789
L792:
	;
	if v3664 == int32(0) {
		goto L796
	} else {
		goto L797
	}
L793:
	;
	v3707 = *(*int32)(unsafe.Add(mBase, uint32(v3282)+4))
	if v3707 <= v3669 {
		v3713 = v3704
		goto L792
	} else {
		goto L794
	}
L794:
	;
	v3709 = *(*int32)(unsafe.Add(mBase, uint32(v3282)+12))
	v3713 = v3709 + v3669<<(uint(int32(2))%32)
	goto L792
L795:
	;
	v4452 = *(*int32)(unsafe.Add(mBase, uint32(v3703)))
	v4456 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3724+v3669<<(uint(int32(2))%32)))))
	v4457 = *(*int32)(unsafe.Add(mBase, uint32(v3713)))
	v4458 = *(*int32)(unsafe.Add(mBase, uint32(v4457)+4))
	v4460 = F_makeTargetEntry(m, v4452, v4456, v4458, int32(0))
	mBase = m.M
	v4461 = m.ExcPending
	if v4461 != 0 {
		goto L34
	} else {
		goto L961
	}
L796:
	;
	v3726 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3726 == int32(0) {
		goto L801
	} else {
		goto L802
	}
L797:
	;
	v3716 = int32(0)
	v3720 = *(*int32)(unsafe.Add(mBase, uint32(v3664)+4))
	if base.B2i32(v3713 == v3716)|(base.B2i32(v3703 == v3716)|base.B2i32(v3720 <= v3669)) != 0 {
		goto L796
	} else {
		goto L798
	}
L798:
	;
	v3724 = *(*int32)(unsafe.Add(mBase, uint32(v3664)+12))
	if v3724 != 0 {
		goto L795
	} else {
		goto L799
	}
L799:
	;
	goto L796
L800:
	;
	v4433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4433 != 0 {
		goto L955
	} else {
		goto L956
	}
L801:
	;
	v3729 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v3729 == int32(0) {
		goto L800
	} else {
		goto L804
	}
L802:
	;
	goto L803
L803:
	;
	v3732 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v3732
	v3734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3736 = int32(1)
	F_addNSItemToQuery(m, l0, v3734, v3732, v3736, v3736)
	mBase = m.M
	v3739 = m.ExcPending
	if v3739 != 0 {
		goto L34
	} else {
		goto L805
	}
L804:
	;
	goto L803
L805:
	;
	v3740 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3740 == int32(0) {
		goto L800
	} else {
		goto L806
	}
L806:
	;
	v3743 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+4))
	if v3743 == int32(3) {
		goto L809
	} else {
		goto L810
	}
L807:
	;
	v3796 = m.G0
	v3798 = v3796 + int32(-64)
	m.G0 = v3798
	v3800 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+8))
	v3802 = v30 + int32(156)
	v3803 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3802))) = v3803
	v3806 = v30 + int32(172)
	*(*int32)(unsafe.Add(mBase, uint32(v3806))) = v3803
	v3810 = v30 + int32(168)
	*(*int32)(unsafe.Add(mBase, uint32(v3810))) = v3803
	v3813 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+4))
	if base.B2i32(v3813 != int32(2))&base.B2i32(v3813 != int32(3))|v3800 != 0 {
		goto L828
	} else {
		goto L829
	}
L808:
	;
	v3776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3780 = F_makeAlias(m, int32(_a_F_transformStmt_59), int32(0))
	mBase = m.M
	v3781 = m.ExcPending
	if v3781 != 0 {
		goto L34
	} else {
		goto L819
	}
L809:
	;
	v3746 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v3746 != 0 {
		goto L808
	} else {
		goto L812
	}
L810:
	;
	goto L811
L811:
	;
	v3767 = int32(0)
	if v3743&int32(-2) != int32(2) {
		v3792 = v3767
		v3794 = v3767
		v3795 = v3767
		goto L807
	} else {
		goto L818
	}
L812:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3750 = m.ExcPending
	if v3750 != 0 {
		goto L34
	} else {
		goto L813
	}
L813:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L34
	} else {
		goto L814
	}
L814:
	;
	F_errmsg(m, int32(_a_F_transformStmt_60), int32(0))
	mBase = m.M
	v3757 = m.ExcPending
	if v3757 != 0 {
		goto L34
	} else {
		goto L815
	}
L815:
	;
	v3758 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3759 = *(*int32)(unsafe.Add(mBase, uint32(v3758)+24))
	F_parser_errposition(m, l0, v3759)
	mBase = m.M
	v3761 = m.ExcPending
	if v3761 != 0 {
		goto L34
	} else {
		goto L816
	}
L816:
	;
	F_errfinish(m, int32(_a_F_transformStmt_18), int32(1053), int32(_a_F_transformStmt_57))
	mBase = m.M
	v3766 = m.ExcPending
	if v3766 != 0 {
		goto L34
	} else {
		goto L817
	}
L817:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L818:
	;
	goto L808
L819:
	;
	v3782 = int32(0)
	v3784 = F_addRangeTableEntryForRelation(m, l0, v3776, int32(3), v3780, v3782, v3782)
	mBase = m.M
	v3785 = m.ExcPending
	if v3785 != 0 {
		goto L34
	} else {
		goto L820
	}
L820:
	;
	v3786 = *(*int32)(unsafe.Add(mBase, uint32(v3784)+8))
	v3787 = *(*int32)(unsafe.Add(mBase, uint32(v3784)+4))
	v3788 = int32(99)
	*(*uint8)(unsafe.Add(mBase, uint32(v3787)+21)) = uint8(v3788)
	v3790 = F_BuildOnConflictExcludedTargetlist(m, v3776, v3786)
	mBase = m.M
	v3791 = m.ExcPending
	if v3791 != 0 {
		goto L34
	} else {
		goto L821
	}
L821:
	;
	v3792 = v3784
	v3794 = v3786
	v3795 = v3790
	goto L807
L822:
	;
	v4356 = int32(0)
	v4358 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+4))
	if v4358&int32(-2) == int32(2) {
		goto L944
	} else {
		goto L945
	}
L823:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4340 = m.ExcPending
	if v4340 != 0 {
		goto L34
	} else {
		goto L939
	}
L824:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4318 = m.ExcPending
	if v4318 != 0 {
		goto L34
	} else {
		goto L934
	}
L825:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4296 = m.ExcPending
	if v4296 != 0 {
		goto L34
	} else {
		goto L929
	}
L826:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4271 = m.ExcPending
	if v4271 != 0 {
		goto L34
	} else {
		goto L924
	}
L827:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4252 = m.ExcPending
	if v4252 != 0 {
		goto L34
	} else {
		goto L919
	}
L828:
	;
	v3820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3821 = *(*int32)(unsafe.Add(mBase, uint32(v3820)+56))
	goto L831
L829:
	;
	goto L830
L830:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4223 = m.ExcPending
	if v4223 != 0 {
		goto L34
	} else {
		goto L910
	}
L831:
	;
	if base.Ui32(v3821) < base.Ui32(int32(_a_F_transformStmt_61)) {
		goto L827
	} else {
		goto L832
	}
L832:
	;
	v3824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3825 = *(*int32)(unsafe.Add(mBase, uint32(v3824)+180))
	if v3825 == int32(0) {
		goto L833
	} else {
		goto L834
	}
L833:
	;
	if v3800 == int32(0) {
		goto L837
	} else {
		goto L838
	}
L834:
	;
	v3828 = *(*int32)(unsafe.Add(mBase, uint32(v3824)+48))
	v3829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3828)+119)))
	switch v3829 - int32(109) {
	case 0, 5:
		goto L835
	default:
		goto L833
	}
L835:
	;
	v3832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3825)+112)))
	if v3832 == int32(1) {
		goto L826
	} else {
		goto L836
	}
L836:
	;
	goto L833
L837:
	;
	m.G0 = v3798 - int32(-64)
	goto L822
L838:
	;
	v3837 = *(*int32)(unsafe.Add(mBase, uint32(v3800)+4))
	if v3837 != 0 {
		goto L839
	} else {
		goto L840
	}
L839:
	;
	v3838 = *(*int32)(unsafe.Add(mBase, uint32(v3837)+4))
	if int32(0) < v3838 {
		goto L842
	} else {
		goto L843
	}
L840:
	;
	goto L841
L841:
	;
	v3984 = *(*int32)(unsafe.Add(mBase, uint32(v3800)+8))
	if v3984 != 0 {
		goto L868
	} else {
		goto L869
	}
L842:
	;
	v3852 = v3
	v3857 = v3
	goto L845
L843:
	;
	v3945 = v3
	goto L844
L844:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3802))) = v3945
	goto L841
L845:
	;
	v3868 = *(*int32)(unsafe.Add(mBase, uint32(v3837)+12))
	v3872 = *(*int32)(unsafe.Add(mBase, uint32(v3868+v3852<<(uint(int32(2))%32))))
	v3874 = F_palloc0(m, int32(16))
	mBase = m.M
	v3875 = m.ExcPending
	if v3875 != 0 {
		goto L34
	} else {
		goto L847
	}
L846:
	;
	v3945 = v3923
	goto L844
L847:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3874))) = int32(60)
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+28))
	if v3878 != 0 {
		goto L825
	} else {
		goto L848
	}
L848:
	;
	v3879 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+32))
	if v3879 != 0 {
		goto L824
	} else {
		goto L849
	}
L849:
	;
	v3880 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+24))
	if v3880 != 0 {
		goto L823
	} else {
		goto L850
	}
L850:
	;
	v3881 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+8))
	if v3881 == int32(0) {
		goto L851
	} else {
		goto L852
	}
L851:
	;
	v3885 = F_palloc0(m, int32(12))
	mBase = m.M
	v3886 = m.ExcPending
	if v3886 != 0 {
		goto L34
	} else {
		goto L854
	}
L852:
	;
	v3902 = v3881
	goto L853
L853:
	;
	v3905 = F_transformExpr(m, l0, v3902, int32(32))
	mBase = m.M
	v3906 = m.ExcPending
	if v3906 != 0 {
		goto L34
	} else {
		goto L857
	}
L854:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3885))) = int32(69)
	v3889 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+4))
	v3890 = F_makeString(m, v3889)
	mBase = m.M
	v3891 = m.ExcPending
	if v3891 != 0 {
		goto L34
	} else {
		goto L855
	}
L855:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3798)+12)) = v3890
	*(*int32)(unsafe.Add(mBase, uint32(v3798)+60)) = v3890
	v3897 = F_list_make1_impl(m, int32(1), v3796+int32(-52))
	mBase = m.M
	v3898 = m.ExcPending
	if v3898 != 0 {
		goto L34
	} else {
		goto L856
	}
L856:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3885)+4)) = v3897
	v3900 = *(*int32)(unsafe.Add(mBase, uint32(v3800)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3885)+8)) = v3900
	v3902 = v3885
	goto L853
L857:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3874)+4)) = v3905
	v3908 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+16))
	if v3908 != 0 {
		goto L858
	} else {
		goto L859
	}
L858:
	;
	v3909 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+36))
	v3910 = F_LookupCollation(m, l0, v3908, v3909)
	mBase = m.M
	v3911 = m.ExcPending
	if v3911 != 0 {
		goto L34
	} else {
		goto L861
	}
L859:
	;
	v3913 = int32(0)
	goto L860
L860:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3874)+8)) = v3913
	v3915 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+20))
	if v3915 != 0 {
		goto L862
	} else {
		goto L863
	}
L861:
	;
	v3913 = v3910
	goto L860
L862:
	;
	v3918 = F_get_opclass_oid(m, int32(403), v3915, int32(0))
	mBase = m.M
	v3919 = m.ExcPending
	if v3919 != 0 {
		goto L34
	} else {
		goto L865
	}
L863:
	;
	v3921 = int32(0)
	goto L864
L864:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3874)+12)) = v3921
	v3923 = F_lappend(m, v3857, v3874)
	mBase = m.M
	v3924 = m.ExcPending
	if v3924 != 0 {
		goto L34
	} else {
		goto L866
	}
L865:
	;
	v3921 = v3918
	goto L864
L866:
	;
	v3926 = v3852 + int32(1)
	v3927 = *(*int32)(unsafe.Add(mBase, uint32(v3837)+4))
	if v3926 < v3927 {
		v3852 = v3926
		v3857 = v3923
		goto L845
	} else {
		goto L867
	}
L867:
	;
	goto L846
L868:
	;
	v3986 = F_transformExpr(m, l0, v3984, int32(33))
	mBase = m.M
	v3987 = m.ExcPending
	if v3987 != 0 {
		goto L34
	} else {
		goto L871
	}
L869:
	;
	goto L870
L870:
	;
	v3989 = *(*int32)(unsafe.Add(mBase, uint32(v3800)+12))
	if v3989 == int32(0) {
		goto L837
	} else {
		goto L872
	}
L871:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3806))) = v3986
	goto L870
L872:
	;
	v3992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3993 = *(*int32)(unsafe.Add(mBase, uint32(v3992)+12))
	v3994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3995 = *(*int32)(unsafe.Add(mBase, uint32(v3994)+56))
	v3996 = int32(0)
	v3997 = m.G0
	v3999 = v3997 - int32(192)
	m.G0 = v3999
	*(*int32)(unsafe.Add(mBase, uint32(v3810))) = v3996
	v4005 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L34
	} else {
		goto L874
	}
L873:
	;
	v4182 = *(*int64)(unsafe.Add(mBase, uint32(v3993)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3993)+16)) = v4182 | int64(2)
	v4186 = *(*int32)(unsafe.Add(mBase, uint32(v3993)+28))
	v4187 = F_bms_add_members(m, v4186, v4122)
	mBase = m.M
	v4188 = m.ExcPending
	if v4188 != 0 {
		goto L34
	} else {
		goto L909
	}
L874:
	;
	v4008 = v3999 + int32(16)
	F_ScanKeyInit(m, v4008, int32(9), int32(3), int32(184), base.I64_extend_i32_u(v3995))
	mBase = m.M
	v4014 = m.ExcPending
	if v4014 != 0 {
		goto L34
	} else {
		goto L875
	}
L875:
	;
	F_ScanKeyInit(m, v3999+int32(72), int32(10), int32(3), int32(184), int64(0))
	mBase = m.M
	v4022 = m.ExcPending
	if v4022 != 0 {
		goto L34
	} else {
		goto L876
	}
L876:
	;
	F_ScanKeyInit(m, v3999+int32(128), int32(2), int32(3), int32(62), base.I64_extend_i32_u(v3989))
	mBase = m.M
	v4030 = m.ExcPending
	if v4030 != 0 {
		goto L34
	} else {
		goto L877
	}
L877:
	;
	v4035 = F_systable_beginscan(m, v4005, int32(2665), int32(1), int32(0), int32(3), v4008)
	mBase = m.M
	v4036 = m.ExcPending
	if v4036 != 0 {
		goto L34
	} else {
		goto L881
	}
L878:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4166 = m.ExcPending
	if v4166 != 0 {
		goto L34
	} else {
		goto L904
	}
L879:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4153 = m.ExcPending
	if v4153 != 0 {
		goto L34
	} else {
		goto L901
	}
L880:
	;
	F_systable_endscan(m, v4035)
	mBase = m.M
	v4139 = m.ExcPending
	if v4139 != 0 {
		goto L34
	} else {
		goto L898
	}
L881:
	;
	v4037 = F_systable_getnext(m, v4035)
	mBase = m.M
	v4038 = m.ExcPending
	if v4038 != 0 {
		goto L34
	} else {
		goto L882
	}
L882:
	;
	if v4037 == int32(0) {
		v4122 = v3996
		goto L880
	} else {
		goto L883
	}
L883:
	;
	v4041 = *(*int32)(unsafe.Add(mBase, uint32(v4037)+16))
	v4042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4041)+22)))
	v4044 = *(*int32)(unsafe.Add(mBase, uint32(v4041+v4042)))
	*(*int32)(unsafe.Add(mBase, uint32(v3810))) = v4044
	v4047 = *(*int32)(unsafe.Add(mBase, uint32(v4005)+52))
	v4050 = F_heap_getattr_3(m, v4037, v4047, v3999+int32(15))
	mBase = m.M
	v4051 = m.ExcPending
	if v4051 != 0 {
		goto L34
	} else {
		goto L884
	}
L884:
	;
	v4052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3999)+15)))
	if v4052 != 0 {
		v4122 = int32(0)
		goto L880
	} else {
		goto L885
	}
L885:
	;
	v4054 = F_pg_detoast_datum(m, base.I32_wrap_i64(v4050))
	mBase = m.M
	v4055 = m.ExcPending
	if v4055 != 0 {
		goto L34
	} else {
		goto L886
	}
L886:
	;
	v4056 = *(*int32)(unsafe.Add(mBase, uint32(v4054)+4))
	if v4056 != int32(1) {
		goto L879
	} else {
		goto L887
	}
L887:
	;
	v4059 = *(*int32)(unsafe.Add(mBase, uint32(v4054)+16))
	if v4059 < int32(0) {
		goto L879
	} else {
		goto L888
	}
L888:
	;
	v4062 = *(*int32)(unsafe.Add(mBase, uint32(v4054)+8))
	if v4062 != 0 {
		goto L879
	} else {
		goto L889
	}
L889:
	;
	v4063 = *(*int32)(unsafe.Add(mBase, uint32(v4054)+12))
	if v4063 != int32(21) {
		goto L879
	} else {
		goto L890
	}
L890:
	;
	v4066 = int32(0)
	if v4059 == v4066 {
		goto L891
	} else {
		goto L892
	}
L891:
	;
	v4122 = int32(0)
	goto L880
L892:
	;
	goto L893
L893:
	;
	v4084 = int32(0)
	v4086 = v4066
	goto L894
L894:
	;
	v4103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4054+int32(24)+v4086<<(uint(int32(1))%32)))))
	v4106 = F_bms_add_member(m, v4084, v4103+int32(7))
	mBase = m.M
	v4107 = m.ExcPending
	if v4107 != 0 {
		goto L34
	} else {
		goto L896
	}
L895:
	;
	v4122 = v4106
	goto L880
L896:
	;
	v4109 = v4086 + int32(1)
	if v4109 != v4059 {
		v4084 = v4106
		v4086 = v4109
		goto L894
	} else {
		goto L897
	}
L897:
	;
	goto L895
L898:
	;
	v4140 = *(*int32)(unsafe.Add(mBase, uint32(v3810)))
	if v4140 == int32(0) {
		goto L878
	} else {
		goto L899
	}
L899:
	;
	F_relation_close(m, v4005, int32(1))
	mBase = m.M
	v4145 = m.ExcPending
	if v4145 != 0 {
		goto L34
	} else {
		goto L900
	}
L900:
	;
	m.G0 = v3999 + int32(192)
	goto L873
L901:
	;
	F_errmsg_internal(m, int32(_a_F_transformStmt_62), int32(0))
	mBase = m.M
	v4157 = m.ExcPending
	if v4157 != 0 {
		goto L34
	} else {
		goto L902
	}
L902:
	;
	F_errfinish(m, int32(_a_F_transformStmt_63), int32(1314), int32(_a_F_transformStmt_64))
	mBase = m.M
	v4162 = m.ExcPending
	if v4162 != 0 {
		goto L34
	} else {
		goto L903
	}
L903:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L904:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v4169 = m.ExcPending
	if v4169 != 0 {
		goto L34
	} else {
		goto L905
	}
L905:
	;
	v4170 = F_get_rel_name(m, v3995)
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		goto L34
	} else {
		goto L906
	}
L906:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3999)+4)) = v4170
	*(*int32)(unsafe.Add(mBase, uint32(v3999))) = v3989
	F_errmsg(m, int32(_a_F_transformStmt_65), v3999)
	mBase = m.M
	v4176 = m.ExcPending
	if v4176 != 0 {
		goto L34
	} else {
		goto L907
	}
L907:
	;
	F_errfinish(m, int32(_a_F_transformStmt_63), int32(1333), int32(_a_F_transformStmt_64))
	mBase = m.M
	v4181 = m.ExcPending
	if v4181 != 0 {
		goto L34
	} else {
		goto L908
	}
L908:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L909:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3993)+28)) = v4187
	goto L837
L910:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4226 = m.ExcPending
	if v4226 != 0 {
		goto L34
	} else {
		goto L911
	}
L911:
	;
	v4229 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+4))
	if v4229 == int32(2) {
		goto L912
	} else {
		goto L913
	}
L912:
	;
	v4232 = int32(_a_F_transformStmt_66)
	goto L914
L913:
	;
	v4232 = int32(_a_F_transformStmt_67)
	goto L914
L914:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3798))) = v4232
	F_errmsg(m, int32(_a_F_transformStmt_68), v3798)
	mBase = m.M
	v4236 = m.ExcPending
	if v4236 != 0 {
		goto L34
	} else {
		goto L915
	}
L915:
	;
	F_errhint(m, int32(_a_F_transformStmt_69), int32(0))
	mBase = m.M
	v4240 = m.ExcPending
	if v4240 != 0 {
		goto L34
	} else {
		goto L916
	}
L916:
	;
	v4241 = F_exprLocation(m, v3740)
	mBase = m.M
	F_parser_errposition(m, l0, v4241)
	mBase = m.M
	v4243 = m.ExcPending
	if v4243 != 0 {
		goto L34
	} else {
		goto L917
	}
L917:
	;
	F_errfinish(m, int32(_a_F_transformStmt_70), int32(3323), int32(_a_F_transformStmt_71))
	mBase = m.M
	v4248 = m.ExcPending
	if v4248 != 0 {
		goto L34
	} else {
		goto L918
	}
L918:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L919:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4255 = m.ExcPending
	if v4255 != 0 {
		goto L34
	} else {
		goto L920
	}
L920:
	;
	F_errmsg(m, int32(_a_F_transformStmt_72), int32(0))
	mBase = m.M
	v4259 = m.ExcPending
	if v4259 != 0 {
		goto L34
	} else {
		goto L921
	}
L921:
	;
	v4260 = F_exprLocation(m, v3740)
	mBase = m.M
	F_parser_errposition(m, l0, v4260)
	mBase = m.M
	v4262 = m.ExcPending
	if v4262 != 0 {
		goto L34
	} else {
		goto L922
	}
L922:
	;
	F_errfinish(m, int32(_a_F_transformStmt_70), int32(3334), int32(_a_F_transformStmt_71))
	mBase = m.M
	v4267 = m.ExcPending
	if v4267 != 0 {
		goto L34
	} else {
		goto L923
	}
L923:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L924:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4274 = m.ExcPending
	if v4274 != 0 {
		goto L34
	} else {
		goto L925
	}
L925:
	;
	v4275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4276 = *(*int32)(unsafe.Add(mBase, uint32(v4275)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v3798)+48)) = v4276 + int32(4)
	F_errmsg(m, int32(_a_F_transformStmt_73), v3796+int32(-16))
	mBase = m.M
	v4284 = m.ExcPending
	if v4284 != 0 {
		goto L34
	} else {
		goto L926
	}
L926:
	;
	v4285 = F_exprLocation(m, v3740)
	mBase = m.M
	F_parser_errposition(m, l0, v4285)
	mBase = m.M
	v4287 = m.ExcPending
	if v4287 != 0 {
		goto L34
	} else {
		goto L927
	}
L927:
	;
	F_errfinish(m, int32(_a_F_transformStmt_70), int32(3343), int32(_a_F_transformStmt_71))
	mBase = m.M
	v4292 = m.ExcPending
	if v4292 != 0 {
		goto L34
	} else {
		goto L928
	}
L928:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L929:
	;
	F_errcode(m, int32(_a_F_transformStmt_74))
	mBase = m.M
	v4299 = m.ExcPending
	if v4299 != 0 {
		goto L34
	} else {
		goto L930
	}
L930:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3798)+32)) = int32(_a_F_transformStmt_75)
	F_errmsg(m, int32(_a_F_transformStmt_76), v3796+int32(-32))
	mBase = m.M
	v4306 = m.ExcPending
	if v4306 != 0 {
		goto L34
	} else {
		goto L931
	}
L931:
	;
	v4307 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+36))
	F_parser_errposition(m, l0, v4307)
	mBase = m.M
	v4309 = m.ExcPending
	if v4309 != 0 {
		goto L34
	} else {
		goto L932
	}
L932:
	;
	F_errfinish(m, int32(_a_F_transformStmt_70), int32(3230), int32(_a_F_transformStmt_77))
	mBase = m.M
	v4314 = m.ExcPending
	if v4314 != 0 {
		goto L34
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
	F_errcode(m, int32(_a_F_transformStmt_74))
	mBase = m.M
	v4321 = m.ExcPending
	if v4321 != 0 {
		goto L34
	} else {
		goto L935
	}
L935:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3798)+16)) = int32(_a_F_transformStmt_78)
	F_errmsg(m, int32(_a_F_transformStmt_76), v3796+int32(-48))
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		goto L34
	} else {
		goto L936
	}
L936:
	;
	v4329 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+36))
	F_parser_errposition(m, l0, v4329)
	mBase = m.M
	v4331 = m.ExcPending
	if v4331 != 0 {
		goto L34
	} else {
		goto L937
	}
L937:
	;
	F_errfinish(m, int32(_a_F_transformStmt_70), int32(3236), int32(_a_F_transformStmt_77))
	mBase = m.M
	v4336 = m.ExcPending
	if v4336 != 0 {
		goto L34
	} else {
		goto L938
	}
L938:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L939:
	;
	F_errcode(m, int32(_a_F_transformStmt_74))
	mBase = m.M
	v4343 = m.ExcPending
	if v4343 != 0 {
		goto L34
	} else {
		goto L940
	}
L940:
	;
	F_errmsg(m, int32(_a_F_transformStmt_79), int32(0))
	mBase = m.M
	v4347 = m.ExcPending
	if v4347 != 0 {
		goto L34
	} else {
		goto L941
	}
L941:
	;
	v4348 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+36))
	F_parser_errposition(m, l0, v4348)
	mBase = m.M
	v4350 = m.ExcPending
	if v4350 != 0 {
		goto L34
	} else {
		goto L942
	}
L942:
	;
	F_errfinish(m, int32(_a_F_transformStmt_70), int32(3241), int32(_a_F_transformStmt_77))
	mBase = m.M
	v4355 = m.ExcPending
	if v4355 != 0 {
		goto L34
	} else {
		goto L943
	}
L943:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L944:
	;
	v4364 = int32(1)
	F_addNSItemToQuery(m, l0, v3792, int32(0), v4364, v4364)
	mBase = m.M
	v4367 = m.ExcPending
	if v4367 != 0 {
		goto L34
	} else {
		goto L947
	}
L945:
	;
	v4384 = v4356
	v4385 = v4356
	goto L946
L946:
	;
	v4387 = F_palloc0(m, int32(40))
	mBase = m.M
	v4388 = m.ExcPending
	if v4388 != 0 {
		goto L34
	} else {
		goto L954
	}
L947:
	;
	v4368 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+4))
	if v4368 == int32(2) {
		goto L948
	} else {
		goto L949
	}
L948:
	;
	v4371 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+16))
	v4372 = F_transformUpdateTargetList(m, l0, v4371)
	mBase = m.M
	v4373 = m.ExcPending
	if v4373 != 0 {
		goto L34
	} else {
		goto L951
	}
L949:
	;
	v4374 = v4356
	goto L950
L950:
	;
	v4375 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+20))
	v4378 = F_transformWhereClause(m, l0, v4375, int32(6), int32(_a_F_transformStmt_0))
	mBase = m.M
	v4379 = m.ExcPending
	if v4379 != 0 {
		goto L34
	} else {
		goto L952
	}
L951:
	;
	v4374 = v4372
	goto L950
L952:
	;
	v4380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4381 = F_list_delete_last(m, v4380)
	mBase = m.M
	v4382 = m.ExcPending
	if v4382 != 0 {
		goto L34
	} else {
		goto L953
	}
L953:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v4381
	v4384 = v4374
	v4385 = v4378
	goto L946
L954:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4387))) = int32(66)
	v4391 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4387)+4)) = v4391
	v4393 = *(*int32)(unsafe.Add(mBase, uint32(v30)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v4387)+8)) = v4393
	v4395 = *(*int32)(unsafe.Add(mBase, uint32(v30)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v4387)+12)) = v4395
	v4397 = *(*int32)(unsafe.Add(mBase, uint32(v30)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v4387)+16)) = v4397
	v4399 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v4387)+36)) = v3795
	*(*int32)(unsafe.Add(mBase, uint32(v4387)+32)) = v3794
	*(*int32)(unsafe.Add(mBase, uint32(v4387)+28)) = v4385
	*(*int32)(unsafe.Add(mBase, uint32(v4387)+24)) = v4384
	*(*int32)(unsafe.Add(mBase, uint32(v4387)+20)) = v4399
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v4387
	goto L800
L955:
	;
	F_transformReturningClause(m, l0, v36, v4433, int32(24))
	mBase = m.M
	v4436 = m.ExcPending
	if v4436 != 0 {
		goto L34
	} else {
		goto L958
	}
L956:
	;
	goto L957
L957:
	;
	v4437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v4437
	v4439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+56)) = v4439
	v4441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4443 = F_makeFromExpr(m, v4441, int32(0))
	mBase = m.M
	v4444 = m.ExcPending
	if v4444 != 0 {
		goto L34
	} else {
		goto L959
	}
L958:
	;
	goto L957
L959:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+60)) = v4443
	v4446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+38)) = uint8(v4446)
	v4448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+39)) = uint8(v4448)
	F_assign_query_collations(m, l0, v36)
	mBase = m.M
	v4451 = m.ExcPending
	if v4451 != 0 {
		goto L34
	} else {
		goto L960
	}
L960:
	;
	v4479 = v36
	goto L1
L961:
	;
	v4462 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v4463 = F_lappend(m, v4462, v4460)
	mBase = m.M
	v4464 = m.ExcPending
	if v4464 != 0 {
		goto L34
	} else {
		goto L962
	}
L962:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v4463
	v4466 = *(*int32)(unsafe.Add(mBase, uint32(v3661)+32))
	v4469 = F_bms_add_member(m, v4466, v4456+int32(7))
	mBase = m.M
	v4470 = m.ExcPending
	if v4470 != 0 {
		goto L34
	} else {
		goto L963
	}
L963:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3661)+32)) = v4469
	v3669 = v3669 + int32(1)
	goto L787
}
