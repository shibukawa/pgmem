package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAgg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v54 int32
	_ = v54
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v79 int32
	_ = v79
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v210 int64
	_ = v210
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
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
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v842 int64
	_ = v842
	var v844 int64
	_ = v844
	var v845 int64
	_ = v845
	var v849 int64
	_ = v849
	var v861 int32
	_ = v861
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
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
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1020 float64
	_ = v1020
	var v1041 int32
	_ = v1041
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1112 int32
	_ = v1112
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1198 int64
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1254 int32
	_ = v1254
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1351 float64
	_ = v1351
	var v1355 int32
	_ = v1355
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
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1373 float64
	_ = v1373
	var v1374 float64
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1382 int32
	_ = v1382
	var v1383 float64
	_ = v1383
	var v1396 int32
	_ = v1396
	var v1397 float64
	_ = v1397
	var v1403 float64
	_ = v1403
	var v1409 float64
	_ = v1409
	var v1411 float64
	_ = v1411
	var v1414 float64
	_ = v1414
	var v1417 float64
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1430 int32
	_ = v1430
	var v1438 int32
	_ = v1438
	var v1440 float64
	_ = v1440
	var v1445 int64
	_ = v1445
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1477 int32
	_ = v1477
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
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
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1658 int32
	_ = v1658
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1697 int32
	_ = v1697
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1751 int32
	_ = v1751
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1828 int64
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1859 int32
	_ = v1859
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1893 int32
	_ = v1893
	var v1894 float64
	_ = v1894
	var v1895 float64
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1964 int64
	_ = v1964
	var v1966 int64
	_ = v1966
	var v1967 int64
	_ = v1967
	var v1971 int64
	_ = v1971
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v2001 int32
	_ = v2001
	var v2002 int64
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2047 int32
	_ = v2047
	var v2052 int32
	_ = v2052
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2077 int32
	_ = v2077
	var v2095 int32
	_ = v2095
	var v2102 int32
	_ = v2102
	var v2113 int32
	_ = v2113
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[0]))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)))
	if v28 != 0 {
		v2095 = v20
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	m.G0 = v2113 + int32(16)
	return v2102
L7:
	;
	v2102 = int32(0)
	v2113 = v2095
	goto L6
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	switch v30 {
	case 0, 1:
		goto L10
	case 2:
		goto L11
	case 3:
		v1057 = v20
		goto L9
	default:
		v2095 = v20
		goto L7
	}
L9:
	;
	v1061 = m.G0
	v1063 = v1061 - int32(80)
	m.G0 = v1063
	v1066 = l0 + int32(288)
	v1068 = l0 + int32(280)
	goto L219
L10:
	;
	v268 = int32(1)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v269 <= v268 {
		goto L57
	} else {
		goto L58
	}
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)))
	if v31 != 0 {
		v1057 = v20
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v33 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v98 != 0 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	if v33 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v39 = v33
	goto L16
L16:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	if v54&int32(2) != 0 {
		goto L13
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v39
	F_lookup_hash_entries(m, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v60 = int32(_a_F_ExecAgg_0)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+28))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v70 = m.T0[v69].(func(*base.Module, int32, int32, int32) int32)(m, v63, v65, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v61
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	F_MemoryContextReset(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v78 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v78 != 0 {
		v39 = v78
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if int32(0) < v99 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v159 = v97
	goto L26
L26:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v174&int32(-2) != int32(2) {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	v105 = int32(0)
	v107 = v97
	goto L30
L28:
	;
	v137 = v97
	v150 = v98
	goto L29
L29:
	;
	F_pfree(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L34
	}
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v123 = v120 + v105*int32(24)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	F_hashagg_spill_finish(m, l0, v123, v105)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v137 = v127
	v150 = v132
	goto L29
L32:
	;
	v127 = v124 + v107
	v129 = v105 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v129 < v130 {
		v105 = v129
		v107 = v127
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	v159 = v137
	goto L26
L35:
	;
	v222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)) = uint8(v222)
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v224
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v228
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v234 = v230 + int32(4)
	v238 = int32(-1)
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v232)))
	if v239 == int64(0) {
		v261 = v238
		goto L49
	} else {
		goto L50
	}
L36:
	;
	goto L35
L37:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v181 = F_MemoryContextMemAllocated(m, v179, int32(1))
	mBase = m.M
	goto L39
L39:
	;
	goto L40
L40:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v189 = int32(1)
	v190 = F_MemoryContextMemAllocated(m, v188, v189)
	mBase = m.M
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	v195 = F_MemoryContextMemAllocated(m, v193, v189)
	mBase = m.M
	v196 = v181 + v159<<(uint(int32(13))%32) + v190 + v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v197) < base.Ui32(v196) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v196
	goto L43
L42:
	;
	goto L43
L43:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v200 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v210 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v210 == int64(0) {
		goto L36
	} else {
		goto L47
	}
L45:
	;
	v203 = F_LogicalTapeSetBlocks(m, v200)
	mBase = m.M
	v205 = v203 << (uint(int64(3)) % 64)
	v206 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v205) <= base.Ui64(v206) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v205
	goto L44
L47:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v195), base.F64_convert_i64_u(v210)), float64(12))
	goto L36
L48:
	;
	v1057 = v20
	goto L9
L49:
	;
	v264 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+8)) = uint8(v264)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v261
	goto L48
L50:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v232)+20))
	v244 = int32(0)
	goto L51
L51:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v242+v244*int32(12))+4))
	if v252 != int32(1) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v261 = v238
	goto L49
L53:
	;
	v261 = v244
	goto L49
L54:
	;
	goto L55
L55:
	;
	v256 = v244 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v256)) < base.Ui64(v239) {
		v244 = v256
		goto L51
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	v272 = v268
	goto L59
L58:
	;
	v272 = v269
	goto L59
L59:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v280 = v272
	v285 = v273
	goto L60
L60:
	;
	F_ReScanExprContext(m, v278)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v2095 = v20
	goto L7
L62:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v299 < v280 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v303 = v299 + int32(1)
	goto L65
L64:
	;
	v303 = v280
	goto L65
L65:
	;
	if int32(0) <= v299 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v306 = v303
	goto L68
L67:
	;
	v306 = v280
	goto L68
L68:
	;
	if int32(0) < v306 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v311 = int32(0)
	goto L72
L70:
	;
	v341 = v299
	goto L71
L71:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)))
	v355 = int32(1)
	v358 = v280 - v355
	if base.B2i32(v354 != v355)|base.B2i32(v341 < v358) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326+v311<<(uint(int32(2))%32))))
	F_ReScanExprContext(m, v330)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L74
	}
L73:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v341 = v336
	goto L71
L74:
	;
	v334 = v311 + int32(1)
	if v334 != v306 {
		v311 = v334
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v465
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)))
	if v467 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L77:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v363 < v364-int32(1) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v449 = int32(0)
	if base.B2i32(v341 < v449)|base.B2i32(v358 <= v341) != 0 {
		v460 = v280
		v461 = v358
		v462 = v306
		v463 = v285
		v464 = v449
		goto L76
	} else {
		goto L107
	}
L80:
	;
	F_initialize_phase(m, l0, v363+int32(1))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v386 == int32(3) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(-1)
	v374 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v374)
	v376 = int32(1)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	if v378 <= v376 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v381 = v376
	goto L86
L85:
	;
	v381 = v378
	goto L86
L86:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v377)+20))
	v460 = v381
	v461 = v381 - int32(1)
	v462 = v381
	v463 = v384
	v464 = int32(0)
	goto L76
L87:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v389 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v447 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v447)
	v2095 = v20
	goto L7
L90:
	;
	F_tuplesort_end(m, v389)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v394 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = int32(0)
	goto L92
L94:
	;
	F_tuplesort_end(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)) = uint8(v399)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v409 = v405 + int32(4)
	v413 = int32(-1)
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
	if v414 == int64(0) {
		v436 = v413
		goto L99
	} else {
		goto L100
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = int32(0)
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = int32(0)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v445
	v1057 = v20
	goto L9
L99:
	;
	v439 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v409)+8)) = uint8(v439)
	*(*int32)(unsafe.Add(mBase, uint32(v409)+4)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v409))) = v436
	goto L98
L100:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v407)+20))
	v419 = int32(0)
	goto L101
L101:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v417+v419*int32(12))+4))
	if v427 != int32(1) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v436 = v413
	goto L99
L103:
	;
	v436 = v419
	goto L99
L104:
	;
	goto L105
L105:
	;
	v431 = v419 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v431)) < base.Ui64(v414) {
		v419 = v431
		goto L101
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+8))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v455+v341<<(uint(int32(2))%32))+4))
	v460 = v280
	v461 = v358
	v462 = v306
	v463 = v285
	v464 = v459
	goto L76
L108:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)))
	if v1041 == int32(0) {
		v280 = v460
		v285 = v463
		goto L60
	} else {
		goto L216
	}
L109:
	;
	F_prepare_projection_slot(m, l0, v958, v943)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L4
	} else {
		goto L205
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(0)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v520 != 0 {
		goto L123
	} else {
		goto L124
	}
L111:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v463)+72))
	if v470 == int32(0) {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v513 = v511 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v513
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v943 = v513
	v958 = v515
	goto L109
L114:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if base.B2i32(v473 == int32(-1))|base.B2i32(v461 <= v473)|base.B2i32(v464 <= int32(0)) != 0 {
		goto L110
	} else {
		goto L115
	}
L115:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+16))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v482+v464<<(uint(int32(2))%32)-int32(4))))
	if v488 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v277)+20))
	F_MemoryContextReset(m, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v494 = int32(_a_F_ExecAgg_0)
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v277)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v497
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v488)+20))
	v502 = m.T0[v501].(func(*base.Module, int32, int32, int32) int32)(m, v488, v277, v20+int32(13))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L120
	}
L119:
	;
	goto L110
L120:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v495
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v277)+20))
	F_MemoryContextReset(m, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	if v502 != 0 {
		goto L110
	} else {
		goto L122
	}
L122:
	;
	goto L113
L123:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v596 = int32(0)
	goto L141
L124:
	;
	v521 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L126
	}
L125:
	;
	if int32(0) < v269 {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	if v521 == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+4)))
	if v525&int32(2) != 0 {
		goto L125
	} else {
		goto L128
	}
L128:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v521)+8))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+44))
	v530 = m.T0[v529].(func(*base.Module, int32) int32)(m, v521)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = v530
	goto L123
L130:
	;
	v535 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v535)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v541 = v537
	goto L133
L131:
	;
	goto L132
L132:
	;
	v569 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v569)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v463)+72))
	if v571 != 0 {
		v2095 = v20
		goto L7
	} else {
		goto L140
	}
L133:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v538)+8))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v556+v541<<(uint(int32(2))%32))))
	if int32(0) < v560 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v567 < v460 {
		goto L123
	} else {
		goto L139
	}
L135:
	;
	v564 = v541 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v564
	if v564 < v460 {
		v541 = v564
		goto L133
	} else {
		goto L138
	}
L136:
	;
	v567 = v541
	goto L137
L137:
	;
	goto L134
L138:
	;
	v567 = v564
	goto L137
L139:
	;
	goto L108
L140:
	;
	goto L123
L141:
	;
	v610 = v596 << (uint(int32(2)) % 32)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v275+v610)))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v613+v610)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v596
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v615
	v618 = int32(0)
	if v618 < v590 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v669 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L143:
	;
	v623 = v618
	goto L146
L144:
	;
	goto L145
L145:
	;
	v667 = v596 + int32(1)
	if v667 != v462 {
		v596 = v667
		goto L141
	} else {
		goto L150
	}
L146:
	;
	F_initialize_aggregate(m, l0, v589+v623*int32(224), v612+v623<<(uint(int32(3))%32))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L4
	} else {
		goto L148
	}
L147:
	;
	goto L145
L148:
	;
	v647 = v623 + int32(1)
	if v647 != v590 {
		v623 = v647
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	goto L142
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+12)) = v274
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v943 = v940
	v958 = v274
	goto L109
L152:
	;
	F_ExecForceStoreHeapTuple(m, v669, v274, int32(1))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+12)) = v274
	goto L154
L154:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v695 != int32(3) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v720)+8))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v917)+44))
	v919 = m.T0[v918].(func(*base.Module, int32) int32)(m, v720)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L4
	} else {
		goto L204
	}
L156:
	;
	v703 = int32(_a_F_ExecAgg_0)
	v704 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+28))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v708)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v709
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v706)+20))
	v713 = m.T0[v712].(func(*base.Module, int32, int32, int32) int32)(m, v706, v708, int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L4
	} else {
		goto L160
	}
L157:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v698 != int32(1) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	F_lookup_hash_entries(m, l0)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	goto L156
L160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v704
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v277)+20))
	F_MemoryContextReset(m, v717)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v720 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+12)) = v720
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v463)+72))
	if v887 == int32(0) {
		goto L154
	} else {
		goto L199
	}
L163:
	;
	if v720 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720)+4)))
	if v722&int32(2) == int32(0) {
		goto L162
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v727 != int32(3) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L166
L168:
	;
	if int32(0) < v269 {
		goto L196
	} else {
		goto L197
	}
L169:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v730 != int32(1) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v733 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v813&int32(-2) != int32(2) {
		goto L184
	} else {
		goto L185
	}
L172:
	;
	v798 = int32(0)
	goto L171
L173:
	;
	goto L174
L174:
	;
	v737 = int32(0)
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v737 < v739 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v744 = v737
	v746 = v737
	goto L178
L176:
	;
	v776 = v737
	v789 = v733
	goto L177
L177:
	;
	F_pfree(m, v789)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L4
	} else {
		goto L182
	}
L178:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v762 = v759 + v744*int32(24)
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v762)))
	F_hashagg_spill_finish(m, l0, v762, v744)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L4
	} else {
		goto L180
	}
L179:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v776 = v766
	v789 = v771
	goto L177
L180:
	;
	v766 = v746 + v763
	v768 = v744 + int32(1)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v768 < v769 {
		v744 = v768
		v746 = v766
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	v798 = v776
	goto L171
L183:
	;
	v861 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v861)
	goto L168
L184:
	;
	goto L183
L185:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v820 = F_MemoryContextMemAllocated(m, v818, int32(1))
	mBase = m.M
	goto L187
L187:
	;
	goto L188
L188:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v828 = int32(1)
	v829 = F_MemoryContextMemAllocated(m, v827, v828)
	mBase = m.M
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+20))
	v834 = F_MemoryContextMemAllocated(m, v832, v828)
	mBase = m.M
	v835 = v820 + v798<<(uint(int32(13))%32) + v829 + v834
	v836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v836) < base.Ui32(v835) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v835
	goto L191
L190:
	;
	goto L191
L191:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v839 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v849 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v849 == int64(0) {
		goto L184
	} else {
		goto L195
	}
L193:
	;
	v842 = F_LogicalTapeSetBlocks(m, v839)
	mBase = m.M
	v844 = v842 << (uint(int64(3)) % 64)
	v845 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v844) <= base.Ui64(v845) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v844
	goto L192
L195:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v834), base.F64_convert_i64_u(v849)), float64(12))
	goto L184
L196:
	;
	v882 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v882)
	goto L151
L197:
	;
	goto L198
L198:
	;
	v884 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v884)
	goto L151
L199:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v463)+80))
	if v890 <= int32(0) {
		goto L154
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v274
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v894)+16))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v463)+80))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v895+v896<<(uint(int32(2))%32)-int32(4))))
	if v902 == int32(0) {
		goto L154
	} else {
		goto L201
	}
L201:
	;
	v905 = int32(_a_F_ExecAgg_0)
	v906 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v277)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v908
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v902)+20))
	v913 = m.T0[v912].(func(*base.Module, int32, int32, int32) int32)(m, v902, v277, v20+int32(14))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v906
	if v913 != 0 {
		goto L154
	} else {
		goto L203
	}
L203:
	;
	goto L155
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = v919
	goto L151
L205:
	;
	v962 = v943 << (uint(int32(2)) % 32)
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v962+v963)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v943
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v965
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v962+v275)))
	F_finalize_aggregates(m, l0, v276, v969)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v972 != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1017 == int32(0) {
		goto L108
	} else {
		goto L215
	}
L208:
	;
	v973 = int32(_a_F_ExecAgg_0)
	v974 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v977
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v972)+20))
	v982 = m.T0[v981].(func(*base.Module, int32, int32, int32) int32)(m, v972, v976, v20+int32(15))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L4
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v990)+72))
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v990)+16))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v992)+8))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v993)+12))
	m.T0[v994].(func(*base.Module, int32))(m, v992)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L4
	} else {
		goto L213
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v974
	if v982 == int32(0) {
		goto L207
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	v997 = int32(_a_F_ExecAgg_0)
	v998 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v991)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1000
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v990)+24))
	v1006 = m.T0[v1005].(func(*base.Module, int32, int32, int32) int32)(m, v990+int32(4), v991, int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L4
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v998
	v1010 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v992)+4)))
	v1012 = v1010 & int32(_a_F_ExecAgg_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v992)+4)) = uint16(v1012)
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v992)+12))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	*(*uint16)(unsafe.Add(mBase, uint32(v992)+6)) = uint16(v1015)
	v2102 = v992
	v2113 = v20
	goto L6
L215:
	;
	v1020 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+240)) = base.F64_add(v1020, float64(1))
	goto L108
L216:
	;
	goto L61
L217:
	;
	if v1561 == int32(0) {
		v2095 = v1057
		goto L7
	} else {
		goto L448
	}
L218:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L4
	} else {
		goto L444
	}
L219:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1098 = v1086 + v1087*int32(52)
	goto L224
L220:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L4
	} else {
		goto L440
	}
L221:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+320)) = int64(0)
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1599)))
	if v1600 != int32(3) {
		goto L336
	} else {
		goto L337
	}
L222:
	;
	m.G0 = v1063 + int32(80)
	goto L217
L223:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+72))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+16))
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1535)+8))
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+12))
	m.T0[v1537].(func(*base.Module, int32))(m, v1535)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L4
	} else {
		goto L333
	}
L224:
	;
	v1112 = v1098 + int32(4)
	goto L226
L225:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v1355 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L226:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+16))
	v1133 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[0]))
	if v1133 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	goto L225
L228:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L4
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1130)))
	v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112)+8)))
	v1143 = v1140
	goto L234
L231:
	;
	goto L230
L232:
	;
	goto L227
L233:
	;
	if v1175 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L234:
	;
	if v1143&int32(1) != 0 {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v1175 = v1158
	goto L233
L236:
	;
	v1175 = int32(0)
	goto L233
L237:
	;
	goto L238
L238:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+20))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+12))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1112)))
	v1154 = v1150 & (v1151 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v1112))) = v1154
	v1158 = v1149 + v1151*int32(12)
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+12))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	v1162 = v1159 & (v1160 ^ v1154)
	if v1162 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1165 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1112)+8)) = uint8(v1165)
	goto L241
L240:
	;
	goto L241
L241:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+4))
	if v1169 != int32(1) {
		v1143 = base.B2i32(v1162 == int32(0))
		goto L234
	} else {
		goto L242
	}
L242:
	;
	goto L235
L243:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1180 = v1178 + int32(1)
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v1181 <= v1180 {
		goto L232
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+20))
	F_MemoryContextReset(m, v1227)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L4
	} else {
		goto L256
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v1180
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1184
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1189 = v1186 + v1180*int32(52)
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1189)))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	v1193 = v1189 + int32(4)
	v1197 = int32(-1)
	v1198 = *(*int64)(unsafe.Add(mBase, uint32(v1191)))
	if v1198 == int64(0) {
		v1220 = v1197
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v1098 = v1189
	goto L224
L248:
	;
	v1223 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1193)+8)) = uint8(v1223)
	*(*int32)(unsafe.Add(mBase, uint32(v1193)+4)) = v1220
	*(*int32)(unsafe.Add(mBase, uint32(v1193))) = v1220
	goto L247
L249:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+20))
	v1203 = int32(0)
	goto L250
L250:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1201+v1203*int32(12))+4))
	if v1211 != int32(1) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v1220 = v1197
	goto L248
L252:
	;
	v1220 = v1203
	goto L248
L253:
	;
	goto L254
L254:
	;
	v1215 = v1203 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1215)) < base.Ui64(v1198) {
		v1203 = v1215
		goto L250
	} else {
		goto L255
	}
L255:
	;
	goto L251
L256:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1175)))
	v1232 = F_ExecStoreMinimalTuple(m, v1230, v1131, int32(0))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+12))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1234)))
	v1236 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1131)+6)))
	if v1236 < v1235 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	F_slot_getsomeattrs_int(m, v1131, v1235)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L4
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+8))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+12))
	m.T0[v1241].(func(*base.Module, int32))(m, v1091)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L4
	} else {
		goto L262
	}
L261:
	;
	goto L260
L262:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+12))
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1244)))
	if v1245 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+20))
	base.MemoryFill(m, v1246, int32(1), v1245)
	goto L265
L264:
	;
	goto L265
L265:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+32))
	if int32(0) < v1249 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1254 = int32(0)
	goto L269
L267:
	;
	goto L268
L268:
	;
	v1314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091)+4)))
	v1316 = v1314 & int32(_a_F_ExecAgg_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1091)+4)) = uint16(v1316)
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+12))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1318)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1091)+6)) = uint16(v1319)
	goto L272
L269:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+16))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+40))
	v1272 = int32(1)
	v1275 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1271+v1254<<(uint(v1272)%32)))))
	v1277 = v1275 - v1272
	v1278 = int32(2)
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+16))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1281+v1254<<(uint(v1278)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1270+v1277<<(uint(v1278)%32)))) = v1285
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+20))
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+20))
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1289+v1254))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1287+v1277))) = uint8(v1291)
	v1294 = v1254 + v1272
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+32))
	if v1294 < v1295 {
		v1254 = v1294
		goto L269
	} else {
		goto L271
	}
L270:
	;
	goto L268
L271:
	;
	goto L270
L272:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+32))
	if v1321 != 0 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1175)))
	v1325 = v1322 - v1321
	goto L275
L274:
	;
	v1325 = int32(0)
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+12)) = v1091
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	F_prepare_projection_slot(m, l0, v1091, v1327)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L4
	} else {
		goto L276
	}
L276:
	;
	F_finalize_aggregates(m, l0, v1092, v1325)
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1332 == int32(0) {
		goto L223
	} else {
		goto L278
	}
L278:
	;
	v1335 = int32(_a_F_ExecAgg_0)
	v1336 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1339
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+20))
	v1344 = m.T0[v1343].(func(*base.Module, int32, int32, int32) int32)(m, v1332, v1338, v1063+int32(48))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L4
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1336
	if v1344 != 0 {
		goto L223
	} else {
		goto L280
	}
L280:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1348 == int32(0) {
		goto L226
	} else {
		goto L281
	}
L281:
	;
	v1351 = *(*float64)(unsafe.Add(mBase, uint32(v1348)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v1348)+240)) = base.F64_add(v1351, float64(1))
	goto L226
L282:
	;
	v1358 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v1358)
	v1561 = int32(0)
	goto L222
L283:
	;
	goto L284
L284:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+12))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+4))
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1362+v1363<<(uint(int32(2))%32)-int32(4))))
	v1370 = F_list_delete_last(m, v1355)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L4
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v1370
	v1373 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v1374 = *(*float64)(unsafe.Add(mBase, uint32(v1369)+24))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+4))
	v1382 = F_get_hash_memory_limit(m)
	mBase = m.M
	v1383 = base.F64_convert_i32_u(v1382)
	if base.F64_ge(v1383, base.F64_mul(v1373, v1374)) != 0 {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v1457 = v1455 << (uint(int32(2)) % 32)
	if v1452&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v1457)) == int32(0) {
		goto L315
	} else {
		goto L316
	}
L287:
	;
	goto L291
L288:
	;
	goto L289
L289:
	;
	v1396 = F_get_hash_memory_limit(m)
	mBase = m.M
	v1397 = base.F64_convert_i32_u(v1396)
	v1403 = base.F64_mul(base.F64_add(base.F64_mul(v1397, float64(0.25)), float64(-8192)), float64(0.0001220703125))
	v1409 = base.F64_add(base.F64_div(base.F64_mul(v1373, base.F64_mul(v1374, float64(1.5))), v1397), float64(1))
	if base.F64_gt(v1409, v1403) != 0 {
		goto L293
	} else {
		goto L294
	}
L291:
	;
	goto L292
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1068))) = v1382
	*(*int64)(unsafe.Add(mBase, uint32(v1066))) = base.I64_trunc_sat_f64_u(base.F64_div(v1383, v1373))
	goto L286
L293:
	;
	v1411 = v1403
	goto L295
L294:
	;
	v1411 = v1409
	goto L295
L295:
	;
	if base.F64_lt(v1411, float64(4)) != 0 {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1414 = float64(4)
	goto L298
L297:
	;
	v1414 = v1411
	goto L298
L298:
	;
	if base.F64_gt(v1414, float64(1024)) != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1417 = float64(1024)
	goto L301
L300:
	;
	v1417 = v1414
	goto L301
L301:
	;
	v1419 = F_my_log2(m, base.I32_trunc_sat_f64_s(v1417))
	mBase = m.M
	if int32(31) < v1375+v1419 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1423 = int32(32) - v1375
	goto L304
L303:
	;
	v1423 = v1419
	goto L304
L304:
	;
	goto L306
L306:
	;
	goto L307
L307:
	;
	v1430 = int32(_a_F_ExecAgg_2)<<(uint(v1423)%32) - int32(-8192)
	if base.Ui32(v1430<<(uint(int32(2))%32)) < base.Ui32(v1382) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1438 = v1382 - v1430
	goto L310
L309:
	;
	v1438 = base.I32_trunc_sat_f64_u(base.F64_mul(v1383, float64(0.75)))
	goto L310
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1068))) = v1438
	v1440 = base.F64_convert_i32_u(v1438)
	if base.F64_gt(v1440, v1373) != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1445 = base.I64_trunc_sat_f64_u(base.F64_div(v1440, v1373))
	goto L313
L312:
	;
	v1445 = int64(1)
	goto L313
L313:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1066))) = v1445
	goto L286
L314:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_ReScanExprContext(m, v1485)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L4
	} else {
		goto L323
	}
L315:
	;
	if v1457 == int32(0) {
		goto L314
	} else {
		goto L318
	}
L316:
	;
	v1477 = v1457
	goto L317
L317:
	;
	if v1477 == int32(0) {
		goto L314
	} else {
		goto L322
	}
L318:
	;
	v1467 = v1457 + v1452
	v1469 = v1452 + int32(4)
	if base.Ui32(v1469) < base.Ui32(v1467) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1471 = v1467
	goto L321
L320:
	;
	v1471 = v1469
	goto L321
L321:
	;
	v1477 = (v1452^int32(-1)+v1471)&int32(-4) + int32(4)
	goto L317
L322:
	;
	base.MemoryFill(m, v1452, int32(0), v1477)
	goto L314
L323:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	F_MemoryContextReset(m, v1488)
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L4
	} else {
		goto L324
	}
L324:
	;
	v1491 = int32(0)
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v1492 <= v1491 {
		goto L221
	} else {
		goto L325
	}
L325:
	;
	v1496 = v1491
	goto L326
L326:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1512+v1496*int32(52))))
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1516)))
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1518)))
	v1521 = v1519 * int32(12)
	if v1521 != 0 {
		goto L329
	} else {
		goto L330
	}
L327:
	;
	goto L221
L328:
	;
	v1528 = v1496 + int32(1)
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v1528 < v1529 {
		v1496 = v1528
		goto L326
	} else {
		goto L332
	}
L329:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1518)+20))
	base.MemoryFill(m, v1522, int32(0), v1521)
	goto L331
L330:
	;
	goto L331
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1518)+8)) = int32(0)
	goto L328
L332:
	;
	goto L327
L333:
	;
	v1540 = int32(_a_F_ExecAgg_0)
	v1541 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1534)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1543
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+24))
	v1549 = m.T0[v1548].(func(*base.Module, int32, int32, int32) int32)(m, v1533+int32(4), v1534, int32(0))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L4
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1541
	v1553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1535)+4)))
	v1555 = v1553 & int32(_a_F_ExecAgg_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1535)+4)) = uint16(v1555)
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1535)+12))
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1557)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1535)+6)) = uint16(v1558)
	v1561 = v1535
	goto L222
L335:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v1611
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1613
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1618 != int32(2) {
		goto L339
	} else {
		goto L340
	}
L336:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v1610 = v1603
	goto L335
L337:
	;
	goto L338
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(1)
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v1606 + int32(48)
	v1610 = v1606
	goto L335
L339:
	;
	v1621 = int32(48)
	goto L341
L340:
	;
	v1621 = int32(0)
	goto L341
L341:
	;
	v1622 = v1610 + v1621
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+44))
	if v1623 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)))
	v1627 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v1627)
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(_a_F_ExecAgg_3)
	v1635 = F_ExecBuildAggTrans(m, l0, v1622, int32(0), v1627, v1627)
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L4
	} else {
		goto L345
	}
L343:
	;
	v1641 = v1623
	goto L344
L344:
	;
	v1645 = v1615 + v1611*int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+28)) = v1641
	v1658 = int32(0)
	goto L347
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+44)) = v1635
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v1626)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1629
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+44))
	v1641 = v1640
	goto L344
L346:
	;
	goto L220
L347:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1645)))
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+16))
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v1668 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1063)+47)) = uint8(v1668)
	v1670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)))
	v1672 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[0]))
	if v1672 != 0 {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+8))
	F_LogicalTapeClose(m, v1918)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L4
	} else {
		goto L412
	}
L349:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L4
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+8))
	v1679 = F_LogicalTapeRead(m, v1675, v1063+int32(72), int32(4))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L4
	} else {
		goto L354
	}
L352:
	;
	goto L351
L353:
	;
	goto L348
L354:
	;
	if v1679 != int32(4) {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	if v1679 == int32(0) {
		goto L353
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+72))
	v1707 = F_LogicalTapeRead(m, v1675, v1063+int32(76), int32(4))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L4
	} else {
		goto L363
	}
L358:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L4
	} else {
		goto L359
	}
L359:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L4
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+8)) = v1679
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+4)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1063))) = v1675
	F_errmsg_internal(m, int32(_a_F_ExecAgg_4), v1063)
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L4
	} else {
		goto L361
	}
L361:
	;
	F_errfinish(m, int32(_a_F_ExecAgg_5), int32(3131), int32(_a_F_ExecAgg_6))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L4
	} else {
		goto L362
	}
L362:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L363:
	;
	if v1707 != int32(4) {
		goto L346
	} else {
		goto L364
	}
L364:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+76))
	v1712 = F_palloc(m, v1711)
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L4
	} else {
		goto L365
	}
L365:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1712))) = v1714
	v1716 = int32(4)
	v1720 = F_LogicalTapeRead(m, v1675, v1712+v1716, v1714-v1716)
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L4
	} else {
		goto L366
	}
L366:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+76))
	if v1720 != v1722-int32(4) {
		goto L218
	} else {
		goto L367
	}
L367:
	;
	v1727 = F_ExecStoreMinimalTuple(m, v1712, v1667, int32(1))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L4
	} else {
		goto L368
	}
L368:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v1729)+12)) = v1667
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+36))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1732)+12))
	v1734 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1733)+6)))
	if v1734 < v1731 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	F_slot_getsomeattrs_int(m, v1733, v1731)
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L4
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	if v1670 != 0 {
		goto L373
	} else {
		goto L374
	}
L372:
	;
	goto L371
L373:
	;
	v1741 = int32(0)
	goto L375
L374:
	;
	v1741 = v1063 + int32(47)
	goto L375
L375:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+8))
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+12))
	m.T0[v1743].(func(*base.Module, int32))(m, v1666)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L4
	} else {
		goto L376
	}
L376:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+32))
	if int32(0) < v1746 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1751 = int32(0)
	goto L380
L378:
	;
	goto L379
L379:
	;
	v1811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1666)+4)))
	v1813 = v1811 & int32(_a_F_ExecAgg_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1666)+4)) = uint16(v1813)
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+12))
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1815)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1666)+6)) = uint16(v1816)
	goto L383
L380:
	;
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+16))
	v1768 = int32(2)
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+16))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+40))
	v1773 = int32(1)
	v1776 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1772+v1751<<(uint(v1773)%32)))))
	v1778 = v1776 - v1773
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1771+v1778<<(uint(v1768)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1767+v1751<<(uint(v1768)%32)))) = v1782
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+20))
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+20))
	v1788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1786+v1778))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1784+v1751))) = uint8(v1788)
	v1791 = v1751 + v1773
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+32))
	if v1791 < v1792 {
		v1751 = v1791
		goto L380
	} else {
		goto L382
	}
L381:
	;
	goto L379
L382:
	;
	goto L381
L383:
	;
	v1818 = m.G0
	v1820 = v1818 - int32(16)
	m.G0 = v1820
	v1822 = int32(_a_F_ExecAgg_0)
	v1823 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v1665)+40)) = v1666
	v1828 = *(*int64)(unsafe.Add(mBase, uint32(v1665)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v1665)+44)) = v1828
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1665)))
	if v1741 != 0 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1823
	m.G0 = v1820 + int32(16)
	if v1853 != 0 {
		goto L395
	} else {
		goto L396
	}
L385:
	;
	v1833 = F_tuplehash_insert_hash_internal(m, v1830, v1703, v1820+int32(15))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L4
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	v1851 = F_tuplehash_lookup_hash_internal(m, v1830, v1703)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L4
	} else {
		goto L393
	}
L388:
	;
	v1835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1820)+15)))
	if v1835 == int32(1) {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1838 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1741))) = uint8(v1838)
	v1853 = v1833
	goto L384
L390:
	;
	goto L391
L391:
	;
	v1840 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1741))) = uint8(v1840)
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1843
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+32))
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+8))
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1846)+48))
	v1848 = m.T0[v1847].(func(*base.Module, int32, int32) int32)(m, v1666, v1845)
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L4
	} else {
		goto L392
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1833))) = v1848
	v1853 = v1833
	goto L384
L393:
	;
	v1853 = v1851
	goto L384
L394:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+20))
	F_MemoryContextReset(m, v1915)
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L4
	} else {
		goto L411
	}
L395:
	;
	v1859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063)+47)))
	if v1859 == int32(1) {
		goto L398
	} else {
		goto L399
	}
L396:
	;
	goto L397
L397:
	;
	if v1658 == int32(0) {
		goto L406
	} else {
		goto L407
	}
L398:
	;
	F_initialize_hash_entry(m, l0, v1665, v1853)
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L4
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+32))
	if v1869 != 0 {
		goto L402
	} else {
		goto L403
	}
L401:
	;
	goto L400
L402:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1853)))
	v1873 = v1870 - v1869
	goto L404
L403:
	;
	v1873 = int32(0)
	goto L404
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1864+v1865<<(uint(int32(2))%32)))) = v1873
	v1875 = int32(_a_F_ExecAgg_0)
	v1876 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1877)+28))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1880)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1881
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+20))
	v1885 = m.T0[v1884].(func(*base.Module, int32, int32, int32) int32)(m, v1878, v1880, int32(0))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L4
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1876
	v1913 = v1658
	goto L394
L406:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+4))
	v1894 = *(*float64)(unsafe.Add(mBase, uint32(v1369)+24))
	v1895 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	F_hashagg_spill_init(m, v1063+int32(48), v1361, v1893, v1894, v1895)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L4
	} else {
		goto L409
	}
L407:
	;
	goto L408
L408:
	;
	F_hashagg_spill_tuple(m, l0, v1063+int32(48), v1667, v1703)
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L4
	} else {
		goto L410
	}
L409:
	;
	goto L408
L410:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	*(*int32)(unsafe.Add(mBase, uint32(v1902+v1903<<(uint(int32(2))%32)))) = int32(0)
	v1913 = int32(1)
	goto L394
L411:
	;
	v1658 = v1913
	goto L347
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v1923
	if v1658 != 0 {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	F_hashagg_spill_finish(m, l0, v1063+int32(48), v1928)
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L4
	} else {
		goto L416
	}
L414:
	;
	v1933 = int32(0)
	goto L415
L415:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1935&int32(-2) != int32(2) {
		goto L418
	} else {
		goto L419
	}
L416:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+48))
	v1933 = v1931
	goto L415
L417:
	;
	v1983 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v1983)
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v1985
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1987
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	v1993 = v1989 + v1990*int32(52)
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1993)))
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1994)))
	v1997 = v1993 + int32(4)
	v2001 = int32(-1)
	v2002 = *(*int64)(unsafe.Add(mBase, uint32(v1995)))
	if v2002 == int64(0) {
		v2024 = v2001
		goto L431
	} else {
		goto L432
	}
L418:
	;
	goto L417
L419:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v1942 = F_MemoryContextMemAllocated(m, v1940, int32(1))
	mBase = m.M
	goto L420
L420:
	;
	goto L422
L422:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v1950 = int32(1)
	v1951 = F_MemoryContextMemAllocated(m, v1949, v1950)
	mBase = m.M
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+20))
	v1956 = F_MemoryContextMemAllocated(m, v1954, v1950)
	mBase = m.M
	v1957 = v1942 + (v1933<<(uint(int32(13))%32) - int32(-8192)) + v1951 + v1956
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v1958) < base.Ui32(v1957) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v1957
	goto L425
L424:
	;
	goto L425
L425:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v1961 == int32(0) {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1971 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v1971 == int64(0) {
		goto L418
	} else {
		goto L429
	}
L427:
	;
	v1964 = F_LogicalTapeSetBlocks(m, v1961)
	mBase = m.M
	v1966 = v1964 << (uint(int64(3)) % 64)
	v1967 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v1966) <= base.Ui64(v1967) {
		goto L426
	} else {
		goto L428
	}
L428:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v1966
	goto L426
L429:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v1956), base.F64_convert_i64_u(v1971)), float64(12))
	goto L418
L430:
	;
	F_pfree(m, v1369)
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L4
	} else {
		goto L439
	}
L431:
	;
	v2027 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1997)+8)) = uint8(v2027)
	*(*int32)(unsafe.Add(mBase, uint32(v1997)+4)) = v2024
	*(*int32)(unsafe.Add(mBase, uint32(v1997))) = v2024
	goto L430
L432:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v1995)+20))
	v2007 = int32(0)
	goto L433
L433:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v2005+v2007*int32(12))+4))
	if v2015 != int32(1) {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	v2024 = v2001
	goto L431
L435:
	;
	v2024 = v2007
	goto L431
L436:
	;
	goto L437
L437:
	;
	v2019 = v2007 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v2019)) < base.Ui64(v2002) {
		v2007 = v2019
		goto L433
	} else {
		goto L438
	}
L438:
	;
	goto L434
L439:
	;
	goto L219
L440:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L4
	} else {
		goto L441
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+40)) = v1707
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+36)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+32)) = v1675
	F_errmsg_internal(m, int32(_a_F_ExecAgg_4), v1063+int32(32))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L4
	} else {
		goto L442
	}
L442:
	;
	F_errfinish(m, int32(_a_F_ExecAgg_5), int32(3140), int32(_a_F_ExecAgg_6))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L4
	} else {
		goto L443
	}
L443:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L444:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L4
	} else {
		goto L445
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+16)) = v1675
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+24)) = v1720
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+20)) = v2061 - int32(4)
	F_errmsg_internal(m, int32(_a_F_ExecAgg_4), v1063+int32(16))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L4
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(_a_F_ExecAgg_5), int32(3152), int32(_a_F_ExecAgg_6))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L4
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
	v2077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1561)+4)))
	if v2077&int32(2) == int32(0) {
		v2102 = v1561
		v2113 = v1057
		goto L6
	} else {
		goto L449
	}
L449:
	;
	v2095 = v1057
	goto L7
}
