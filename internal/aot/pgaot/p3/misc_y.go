package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_yiddish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v714 int32
	_ = v714
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v752 int32
	_ = v752
	var v759 int32
	_ = v759
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v832 int32
	_ = v832
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v870 int32
	_ = v870
	var v877 int32
	_ = v877
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v950 int32
	_ = v950
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v976 int32
	_ = v976
	var v988 int32
	_ = v988
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1070 int32
	_ = v1070
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1141 int32
	_ = v1141
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1192 int32
	_ = v1192
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1218 int32
	_ = v1218
	var v1226 int32
	_ = v1226
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
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
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1721 int32
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1891 int32
	_ = v1891
	var v1907 int32
	_ = v1907
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1921 int32
	_ = v1921
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1957 int32
	_ = v1957
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1985 int32
	_ = v1985
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2034 int32
	_ = v2034
	var v2037 int32
	_ = v2037
	var v2041 int32
	_ = v2041
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2070 int32
	_ = v2070
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v6
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v16 = F_find_among(m, l0, int32(_a_F_yiddish_UTF_8_stem_0), int32(8), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v171
	v9 = v171
	goto L1
L4:
	;
	return v2070
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	goto L1
L6:
	;
	v2063 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_1))
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L8
	} else {
		goto L639
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L49
L8:
	;
	return int32(0)
L9:
	;
	if v16 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v22
	switch v16 - int32(1) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L6
	case 4:
		goto L14
	case 5:
		goto L13
	case 6:
		goto L12
	case 7:
		goto L11
	default:
		goto L5
	}
L11:
	;
	v112 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L45
	}
L12:
	;
	v106 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_3))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L43
	}
L13:
	;
	v100 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_4))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L41
	}
L14:
	;
	v94 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_5))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L39
	}
L15:
	;
	v70 = int32(2)
	v72 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v74-v75 < v70 {
		v84 = v72
		goto L33
	} else {
		goto L34
	}
L16:
	;
	v48 = int32(2)
	v50 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v52-v53 < v48 {
		v62 = v50
		goto L26
	} else {
		goto L27
	}
L17:
	;
	v26 = int32(2)
	v28 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30-v31 < v26 {
		v40 = v28
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v40 != 0 {
		goto L7
	} else {
		goto L22
	}
L19:
	;
	goto L18
L20:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = F_memcmp(m, v34+v31, int32(_a_F_yiddish_UTF_8_stem_6), v26)
	mBase = m.M
	if v36 != 0 {
		v40 = v28
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26 + v31
	v40 = int32(1)
	goto L19
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22
	v44 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_7))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	if int32(0) <= v44 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v2070 = v44
	goto L4
L25:
	;
	if v62 != 0 {
		goto L7
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = F_memcmp(m, v56+v53, int32(_a_F_yiddish_UTF_8_stem_8), v48)
	mBase = m.M
	if v58 != 0 {
		v62 = v50
		goto L26
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48 + v53
	v62 = int32(1)
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22
	v66 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_9))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	if int32(0) <= v66 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v2070 = v66
	goto L4
L32:
	;
	if v84 != 0 {
		goto L7
	} else {
		goto L36
	}
L33:
	;
	goto L32
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = F_memcmp(m, v78+v75, int32(_a_F_yiddish_UTF_8_stem_10), v70)
	mBase = m.M
	if v80 != 0 {
		v84 = v72
		goto L33
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v70 + v75
	v84 = int32(1)
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22
	v88 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_11))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	if int32(0) <= v88 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v2070 = v88
	goto L4
L39:
	;
	if int32(0) <= v94 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v2070 = v94
	goto L4
L41:
	;
	if int32(0) <= v100 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v2070 = v100
	goto L4
L43:
	;
	if int32(0) <= v106 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v2070 = v106
	goto L4
L45:
	;
	if int32(0) <= v112 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v2070 = v112
	goto L4
L47:
	;
	if int32(0) <= v171 {
		goto L3
	} else {
		goto L67
	}
L49:
	;
	goto L50
L50:
	;
	goto L51
L51:
	;
	v126 = v9
	v128 = int32(1)
	goto L54
L53:
	;
	v171 = v156
	goto L47
L54:
	;
	if v119 <= v126 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v171 = int32(-1)
	goto L47
L57:
	;
	goto L58
L58:
	;
	v133 = v126 + int32(1)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v126))))
	if base.Ui32(v135) < base.Ui32(int32(192)) {
		v156 = v133
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v157 = int32(1)
	if v157 < v128 {
		v126 = v156
		v128 = v128 - v157
		goto L54
	} else {
		goto L66
	}
L60:
	;
	if v119 <= v133 {
		v156 = v133
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v142 = v133
	goto L62
L62:
	;
	v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v118+v142))))
	if int32(-65) < v145 {
		v156 = v142
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v156 = v119
	goto L59
L64:
	;
	v149 = v142 + int32(1)
	if v149 != v119 {
		v142 = v149
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L55
L67:
	;
	v176 = v6
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L72
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v365
	v367 = int32(4)
	v369 = int32(0)
	if v365-v6 < v367 {
		v381 = v369
		goto L122
	} else {
		goto L123
	}
L70:
	;
	if v298 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L71:
	;
	v298 = v291
	goto L70
L72:
	;
	if v193 <= v176 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v291 = int32(0)
	goto L71
L74:
	;
	v298 = int32(-1)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v209 = int32(1)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v194))))
	if base.Ui32(v211) < base.Ui32(int32(192)) {
		v268 = v211
		v269 = v209
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if int32(1474) < v268 {
		v291 = v269
		goto L71
	} else {
		goto L90
	}
L78:
	;
	v215 = v176 + int32(1)
	if v215 == v193 {
		v268 = v211
		v269 = v209
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215+v194))))
	v220 = v218 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v211) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v194))))
	v236 = v234 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v211) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v224 = v176 + int32(2)
	if v224 != v193 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v268 = v211<<(uint(int32(6))%32)&int32(1984) | v220
	v269 = int32(2)
	goto L77
L84:
	;
	goto L83
L85:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v240))))
	v268 = v253&int32(63) | (v211<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v220<<(uint(int32(12))%32) | v236<<(uint(int32(6))%32))
	v269 = int32(4)
	goto L77
L86:
	;
	v240 = v176 + int32(3)
	if v240 != v193 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v268 = v211<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v220<<(uint(int32(6))%32) | v236
	v269 = int32(3)
	goto L77
L89:
	;
	goto L88
L90:
	;
	v273 = v268 - int32(1456)
	if v273 < int32(0) {
		v291 = v269
		goto L71
	} else {
		goto L91
	}
L91:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v273)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[0]))))
	if int32(base.Ui32(v279)>>(uint(v273&int32(7))%32))&int32(1) == int32(0) {
		v291 = v269
		goto L71
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 + v176
	goto L93
L93:
	;
	goto L73
L94:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v301
	v303 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v303 {
		goto L68
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L100
L97:
	;
	v2070 = v303
	goto L4
L98:
	;
	if int32(0) <= v360 {
		v176 = v360
		goto L68
	} else {
		goto L118
	}
L100:
	;
	goto L101
L101:
	;
	goto L102
L102:
	;
	v315 = v176
	v317 = int32(1)
	goto L105
L104:
	;
	v360 = v345
	goto L98
L105:
	;
	if v308 <= v315 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L104
L107:
	;
	v360 = int32(-1)
	goto L98
L108:
	;
	goto L109
L109:
	;
	v322 = v315 + int32(1)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307+v315))))
	if base.Ui32(v324) < base.Ui32(int32(192)) {
		v345 = v322
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v346 = int32(1)
	if v346 < v317 {
		v315 = v345
		v317 = v317 - v346
		goto L105
	} else {
		goto L117
	}
L111:
	;
	if v308 <= v322 {
		v345 = v322
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v331 = v322
	goto L113
L113:
	;
	v334 = int32(*(*int8)(unsafe.Add(mBase, uint32(v307+v331))))
	if int32(-65) < v334 {
		v345 = v331
		goto L110
	} else {
		goto L115
	}
L114:
	;
	v345 = v308
	goto L110
L115:
	;
	v338 = v331 + int32(1)
	if v338 != v308 {
		v331 = v338
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	goto L106
L118:
	;
	goto L69
L119:
	;
	v435 = F_find_among(m, l0, int32(_a_F_yiddish_UTF_8_stem_14), int32(40), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L8
	} else {
		goto L141
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v431 = v6
	goto L119
L121:
	;
	if v381 == int32(0) {
		goto L120
	} else {
		goto L125
	}
L122:
	;
	goto L121
L123:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v377 = F_memcmp(m, v375+v6, int32(_a_F_yiddish_UTF_8_stem_15), v367)
	mBase = m.M
	if v377 != 0 {
		v381 = v369
		goto L122
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367 + v6
	v381 = int32(1)
	goto L122
L125:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v384
	v386 = int32(4)
	v388 = int32(0)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v390-v384 < v386 {
		v400 = v388
		goto L127
	} else {
		goto L128
	}
L126:
	;
	if v400 != 0 {
		goto L120
	} else {
		goto L130
	}
L127:
	;
	goto L126
L128:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v396 = F_memcmp(m, v394+v384, int32(_a_F_yiddish_UTF_8_stem_16), v386)
	mBase = m.M
	if v396 != 0 {
		v400 = v388
		goto L127
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v386 + v384
	v400 = int32(1)
	goto L127
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v384
	v402 = int32(4)
	v404 = int32(0)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v406-v384 < v402 {
		v416 = v404
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v416 != 0 {
		goto L120
	} else {
		goto L135
	}
L132:
	;
	goto L131
L133:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v412 = F_memcmp(m, v410+v384, int32(_a_F_yiddish_UTF_8_stem_17), v402)
	mBase = m.M
	if v412 != 0 {
		v416 = v404
		goto L132
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v402 + v384
	v416 = int32(1)
	goto L132
L135:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v417 <= v384 {
		goto L120
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v384
	v422 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_18))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L8
	} else {
		goto L137
	}
L137:
	;
	if v422 < int32(0) {
		v2070 = v422
		goto L4
	} else {
		goto L138
	}
L138:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v431 = v426
	goto L119
L139:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L185
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v560
	goto L139
L141:
	;
	if v435 == int32(0) {
		v560 = v431
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v440 = int32(8)
	v442 = int32(0)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v444-v439 < v440 {
		v454 = v442
		goto L146
	} else {
		goto L147
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v493 = int32(8)
	v495 = int32(0)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v497-v439 < v493 {
		v507 = v495
		goto L162
	} else {
		goto L163
	}
L144:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v489 < v490 {
		goto L143
	} else {
		goto L160
	}
L145:
	;
	if v454 != 0 {
		goto L144
	} else {
		goto L149
	}
L146:
	;
	goto L145
L147:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v450 = F_memcmp(m, v448+v439, int32(_a_F_yiddish_UTF_8_stem_19), v440)
	mBase = m.M
	if v450 != 0 {
		v454 = v442
		goto L146
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v440 + v439
	v454 = int32(1)
	goto L146
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v456 = int32(8)
	v458 = int32(0)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v460-v439 < v456 {
		v470 = v458
		goto L151
	} else {
		goto L152
	}
L150:
	;
	if v470 != 0 {
		goto L144
	} else {
		goto L154
	}
L151:
	;
	goto L150
L152:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v466 = F_memcmp(m, v464+v439, int32(_a_F_yiddish_UTF_8_stem_20), v456)
	mBase = m.M
	if v466 != 0 {
		v470 = v458
		goto L151
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v456 + v439
	v470 = int32(1)
	goto L151
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v472 = int32(8)
	v474 = int32(0)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v476-v439 < v472 {
		v486 = v474
		goto L156
	} else {
		goto L157
	}
L155:
	;
	if v486 == int32(0) {
		goto L143
	} else {
		goto L159
	}
L156:
	;
	goto L155
L157:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v482 = F_memcmp(m, v480+v439, int32(_a_F_yiddish_UTF_8_stem_21), v472)
	mBase = m.M
	if v482 != 0 {
		v486 = v474
		goto L156
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v472 + v439
	v486 = int32(1)
	goto L156
L159:
	;
	goto L144
L160:
	;
	v560 = v439
	goto L140
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	if v507 != 0 {
		goto L139
	} else {
		goto L165
	}
L162:
	;
	goto L161
L163:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v503 = F_memcmp(m, v501+v439, int32(_a_F_yiddish_UTF_8_stem_22), v493)
	mBase = m.M
	if v503 != 0 {
		v507 = v495
		goto L162
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v493 + v439
	v507 = int32(1)
	goto L162
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v439
	v510 = int32(4)
	v512 = int32(0)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v514-v515 < v510 {
		v524 = v512
		goto L167
	} else {
		goto L168
	}
L166:
	;
	if v524 != 0 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	goto L166
L168:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v520 = F_memcmp(m, v518+v515, int32(_a_F_yiddish_UTF_8_stem_23), v510)
	mBase = m.M
	if v520 != 0 {
		v524 = v512
		goto L167
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v510 + v515
	v524 = int32(1)
	goto L167
L170:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v525
	v529 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_24))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L8
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v535 = int32(4)
	v537 = int32(0)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v539-v439 < v535 {
		v549 = v537
		goto L176
	} else {
		goto L177
	}
L173:
	;
	if int32(0) <= v529 {
		goto L139
	} else {
		goto L174
	}
L174:
	;
	v2070 = v529
	goto L4
L175:
	;
	if v549 == int32(0) {
		v560 = v431
		goto L140
	} else {
		goto L179
	}
L176:
	;
	goto L175
L177:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v545 = F_memcmp(m, v543+v439, int32(_a_F_yiddish_UTF_8_stem_25), v535)
	mBase = m.M
	if v545 != 0 {
		v549 = v537
		goto L176
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v535 + v439
	v549 = int32(1)
	goto L176
L179:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v552
	v556 = F_slice_from_s(m, l0, int32(3), int32(_a_F_yiddish_UTF_8_stem_26))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L8
	} else {
		goto L180
	}
L180:
	;
	if v556 < int32(0) {
		v2070 = v556
		goto L4
	} else {
		goto L181
	}
L181:
	;
	goto L139
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1249
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1249
	v1255 = F_find_among_b(m, l0, int32(_a_F_yiddish_UTF_8_stem_27), int32(79), int32(0))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L8
	} else {
		goto L343
	}
L183:
	;
	if v619 < int32(0) {
		goto L182
	} else {
		goto L203
	}
L185:
	;
	goto L186
L186:
	;
	goto L187
L187:
	;
	v574 = v566
	v576 = int32(3)
	goto L190
L189:
	;
	v619 = v604
	goto L183
L190:
	;
	if v567 <= v574 {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	goto L189
L192:
	;
	v619 = int32(-1)
	goto L183
L193:
	;
	goto L194
L194:
	;
	v581 = v574 + int32(1)
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+v574))))
	if base.Ui32(v583) < base.Ui32(int32(192)) {
		v604 = v581
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v605 = int32(1)
	if v605 < v576 {
		v574 = v604
		v576 = v576 - v605
		goto L190
	} else {
		goto L202
	}
L196:
	;
	if v567 <= v581 {
		v604 = v581
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v590 = v581
	goto L198
L198:
	;
	v593 = int32(*(*int8)(unsafe.Add(mBase, uint32(v565+v590))))
	if int32(-65) < v593 {
		v604 = v590
		goto L195
	} else {
		goto L200
	}
L199:
	;
	v604 = v567
	goto L195
L200:
	;
	v597 = v590 + int32(1)
	if v597 != v567 {
		v590 = v597
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	goto L191
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v566
	v624 = v566 + int32(5)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v625 <= v624 {
		v641 = v566
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L215
L205:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627+v624))))
	if v629&int32(254) != int32(168) {
		v641 = v566
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v637 = F_find_among(m, l0, int32(_a_F_yiddish_UTF_8_stem_28), int32(4), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L8
	} else {
		goto L207
	}
L207:
	;
	if v637 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v641 = v639
	goto L204
L209:
	;
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v566
	v641 = v566
	goto L204
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1244
	goto L182
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v641
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1019 = v641
	goto L290
L213:
	;
	if v759 != 0 {
		goto L212
	} else {
		goto L237
	}
L214:
	;
	v759 = v752
	goto L213
L215:
	;
	if v654 <= v653 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v752 = int32(0)
	goto L214
L217:
	;
	v759 = int32(-1)
	goto L213
L218:
	;
	goto L219
L219:
	;
	v670 = int32(1)
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653+v655))))
	if base.Ui32(v672) < base.Ui32(int32(192)) {
		v729 = v672
		v730 = v670
		goto L220
	} else {
		goto L221
	}
L220:
	;
	if int32(1520) < v729 {
		v752 = v730
		goto L214
	} else {
		goto L233
	}
L221:
	;
	v676 = v653 + int32(1)
	if v676 == v654 {
		v729 = v672
		v730 = v670
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676+v655))))
	v681 = v679 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v672) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685+v655))))
	v697 = v695 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v672) {
		goto L229
	} else {
		goto L230
	}
L224:
	;
	v685 = v653 + int32(2)
	if v685 != v654 {
		goto L223
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v729 = v672<<(uint(int32(6))%32)&int32(1984) | v681
	v730 = int32(2)
	goto L220
L227:
	;
	goto L226
L228:
	;
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655+v701))))
	v729 = v714&int32(63) | (v672<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v681<<(uint(int32(12))%32) | v697<<(uint(int32(6))%32))
	v730 = int32(4)
	goto L220
L229:
	;
	v701 = v653 + int32(3)
	if v701 != v654 {
		goto L228
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v729 = v672<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v681<<(uint(int32(6))%32) | v697
	v730 = int32(3)
	goto L220
L232:
	;
	goto L231
L233:
	;
	v734 = v729 - int32(1489)
	if v734 < int32(0) {
		v752 = v730
		goto L214
	} else {
		goto L234
	}
L234:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v734)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[1]))))
	if int32(base.Ui32(v740)>>(uint(v734&int32(7))%32))&int32(1) == int32(0) {
		v752 = v730
		goto L214
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v730 + v653
	goto L236
L236:
	;
	goto L216
L237:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L240
L238:
	;
	if v877 != 0 {
		goto L212
	} else {
		goto L262
	}
L239:
	;
	v877 = v870
	goto L238
L240:
	;
	if v772 <= v771 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v870 = int32(0)
	goto L239
L242:
	;
	v877 = int32(-1)
	goto L238
L243:
	;
	goto L244
L244:
	;
	v788 = int32(1)
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771+v773))))
	if base.Ui32(v790) < base.Ui32(int32(192)) {
		v847 = v790
		v848 = v788
		goto L245
	} else {
		goto L246
	}
L245:
	;
	if int32(1520) < v847 {
		v870 = v848
		goto L239
	} else {
		goto L258
	}
L246:
	;
	v794 = v771 + int32(1)
	if v794 == v772 {
		v847 = v790
		v848 = v788
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794+v773))))
	v799 = v797 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v790) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803+v773))))
	v815 = v813 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v790) {
		goto L254
	} else {
		goto L255
	}
L249:
	;
	v803 = v771 + int32(2)
	if v803 != v772 {
		goto L248
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v847 = v790<<(uint(int32(6))%32)&int32(1984) | v799
	v848 = int32(2)
	goto L245
L252:
	;
	goto L251
L253:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773+v819))))
	v847 = v832&int32(63) | (v790<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v799<<(uint(int32(12))%32) | v815<<(uint(int32(6))%32))
	v848 = int32(4)
	goto L245
L254:
	;
	v819 = v771 + int32(3)
	if v819 != v772 {
		goto L253
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v847 = v790<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v799<<(uint(int32(6))%32) | v815
	v848 = int32(3)
	goto L245
L257:
	;
	goto L256
L258:
	;
	v852 = v847 - int32(1489)
	if v852 < int32(0) {
		v870 = v848
		goto L239
	} else {
		goto L259
	}
L259:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v852)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[1]))))
	if int32(base.Ui32(v858)>>(uint(v852&int32(7))%32))&int32(1) == int32(0) {
		v870 = v848
		goto L239
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v848 + v771
	goto L261
L261:
	;
	goto L241
L262:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L265
L263:
	;
	if v995 != 0 {
		goto L212
	} else {
		goto L287
	}
L264:
	;
	v995 = v988
	goto L263
L265:
	;
	if v890 <= v889 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v988 = int32(0)
	goto L264
L267:
	;
	v995 = int32(-1)
	goto L263
L268:
	;
	goto L269
L269:
	;
	v906 = int32(1)
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v889+v891))))
	if base.Ui32(v908) < base.Ui32(int32(192)) {
		v965 = v908
		v966 = v906
		goto L270
	} else {
		goto L271
	}
L270:
	;
	if int32(1520) < v965 {
		v988 = v966
		goto L264
	} else {
		goto L283
	}
L271:
	;
	v912 = v889 + int32(1)
	if v912 == v890 {
		v965 = v908
		v966 = v906
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912+v891))))
	v917 = v915 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v908) {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921+v891))))
	v933 = v931 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v908) {
		goto L279
	} else {
		goto L280
	}
L274:
	;
	v921 = v889 + int32(2)
	if v921 != v890 {
		goto L273
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v965 = v908<<(uint(int32(6))%32)&int32(1984) | v917
	v966 = int32(2)
	goto L270
L277:
	;
	goto L276
L278:
	;
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891+v937))))
	v965 = v950&int32(63) | (v908<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v917<<(uint(int32(12))%32) | v933<<(uint(int32(6))%32))
	v966 = int32(4)
	goto L270
L279:
	;
	v937 = v889 + int32(3)
	if v937 != v890 {
		goto L278
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v965 = v908<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v917<<(uint(int32(6))%32) | v933
	v966 = int32(3)
	goto L270
L282:
	;
	goto L281
L283:
	;
	v970 = v965 - int32(1489)
	if v970 < int32(0) {
		v988 = v966
		goto L264
	} else {
		goto L284
	}
L284:
	;
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v970)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[1]))))
	if int32(base.Ui32(v976)>>(uint(v970&int32(7))%32))&int32(1) == int32(0) {
		v988 = v966
		goto L264
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v966 + v889
	goto L286
L286:
	;
	goto L266
L287:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1244 = v996
	goto L211
L288:
	;
	if v1114 < int32(0) {
		goto L182
	} else {
		goto L313
	}
L289:
	;
	v1114 = v1086
	goto L288
L290:
	;
	if v1010 <= v1019 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1114 = int32(-1)
	goto L288
L293:
	;
	goto L294
L294:
	;
	v1026 = int32(1)
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019+v1011))))
	if base.Ui32(v1028) < base.Ui32(int32(192)) {
		v1085 = v1028
		v1086 = v1026
		goto L295
	} else {
		goto L296
	}
L295:
	;
	if int32(1522) < v1085 {
		goto L308
	} else {
		goto L309
	}
L296:
	;
	v1032 = v1019 + int32(1)
	if v1032 == v1010 {
		v1085 = v1028
		v1086 = v1026
		goto L295
	} else {
		goto L297
	}
L297:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032+v1011))))
	v1037 = v1035 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1028) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1041+v1011))))
	v1053 = v1051 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1028) {
		goto L304
	} else {
		goto L305
	}
L299:
	;
	v1041 = v1019 + int32(2)
	if v1041 != v1010 {
		goto L298
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v1085 = v1028<<(uint(int32(6))%32)&int32(1984) | v1037
	v1086 = int32(2)
	goto L295
L302:
	;
	goto L301
L303:
	;
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1011+v1057))))
	v1085 = v1070&int32(63) | (v1028<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v1037<<(uint(int32(12))%32) | v1053<<(uint(int32(6))%32))
	v1086 = int32(4)
	goto L295
L304:
	;
	v1057 = v1019 + int32(3)
	if v1057 != v1010 {
		goto L303
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1085 = v1028<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v1037<<(uint(int32(6))%32) | v1053
	v1086 = int32(3)
	goto L295
L307:
	;
	goto L306
L308:
	;
	v1103 = v1086 + v1019
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1103
	v1019 = v1103
	goto L290
L309:
	;
	v1090 = v1085 - int32(1488)
	if v1090 < int32(0) {
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1090)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[2]))))
	if int32(base.Ui32(v1096)>>(uint(v1090&int32(7))%32))&int32(1) != 0 {
		goto L289
	} else {
		goto L311
	}
L311:
	;
	goto L308
L313:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1118 = v1117 + v1114
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1118
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1141 = v1118
	goto L316
L314:
	;
	if v1237 < int32(0) {
		goto L182
	} else {
		goto L338
	}
L315:
	;
	v1237 = v1208
	goto L314
L316:
	;
	if v1132 <= v1141 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1237 = int32(-1)
	goto L314
L319:
	;
	goto L320
L320:
	;
	v1148 = int32(1)
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1141+v1133))))
	if base.Ui32(v1150) < base.Ui32(int32(192)) {
		v1207 = v1150
		v1208 = v1148
		goto L321
	} else {
		goto L322
	}
L321:
	;
	if int32(1522) < v1207 {
		goto L315
	} else {
		goto L334
	}
L322:
	;
	v1154 = v1141 + int32(1)
	if v1154 == v1132 {
		v1207 = v1150
		v1208 = v1148
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1154+v1133))))
	v1159 = v1157 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1150) {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163+v1133))))
	v1175 = v1173 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1150) {
		goto L330
	} else {
		goto L331
	}
L325:
	;
	v1163 = v1141 + int32(2)
	if v1163 != v1132 {
		goto L324
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1207 = v1150<<(uint(int32(6))%32)&int32(1984) | v1159
	v1208 = int32(2)
	goto L321
L328:
	;
	goto L327
L329:
	;
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1133+v1179))))
	v1207 = v1192&int32(63) | (v1150<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v1159<<(uint(int32(12))%32) | v1175<<(uint(int32(6))%32))
	v1208 = int32(4)
	goto L321
L330:
	;
	v1179 = v1141 + int32(3)
	if v1179 != v1132 {
		goto L329
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1207 = v1150<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v1159<<(uint(int32(6))%32) | v1175
	v1208 = int32(3)
	goto L321
L333:
	;
	goto L332
L334:
	;
	v1212 = v1207 - int32(1488)
	if v1212 < int32(0) {
		goto L315
	} else {
		goto L335
	}
L335:
	;
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1212)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[2]))))
	if int32(base.Ui32(v1218)>>(uint(v1212&int32(7))%32))&int32(1) == int32(0) {
		goto L315
	} else {
		goto L336
	}
L336:
	;
	v1226 = v1208 + v1141
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1226
	v1141 = v1226
	goto L316
L338:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v619 < v1240 {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1242 = v1240
	goto L341
L340:
	;
	v1242 = v619
	goto L341
L341:
	;
	v1244 = v1242
	goto L211
L342:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1747
	v1751 = v1747 - int32(1)
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1751 <= v1752 {
		goto L562
	} else {
		goto L563
	}
L343:
	;
	if v1255 == int32(0) {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1259
	switch v1255 - int32(1) {
	case 0:
		goto L377
	case 1:
		goto L376
	case 2:
		goto L375
	case 3:
		goto L374
	case 4:
		goto L373
	case 5:
		goto L372
	case 6:
		goto L371
	case 7:
		goto L370
	case 8:
		goto L369
	case 9:
		goto L368
	case 10:
		goto L367
	case 11:
		goto L366
	case 12:
		goto L365
	case 13:
		goto L364
	case 14:
		goto L363
	case 15:
		goto L362
	case 16:
		goto L361
	case 17:
		goto L360
	case 18:
		goto L359
	case 19:
		goto L358
	case 20:
		goto L357
	case 21:
		goto L356
	case 22:
		goto L355
	case 23:
		goto L354
	case 24:
		goto L353
	case 25:
		goto L352
	case 26:
		goto L351
	case 27:
		goto L350
	case 28:
		goto L349
	case 29:
		goto L348
	case 30:
		goto L347
	case 31:
		goto L346
	case 32:
		goto L345
	default:
		goto L342
	}
L345:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1684 = int32(2)
	v1686 = int32(0)
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1688-v1689 < v1684 {
		v1699 = v1686
		goto L546
	} else {
		goto L547
	}
L346:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1259 < v1675 {
		goto L342
	} else {
		goto L541
	}
L347:
	;
	v1671 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_29))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L8
	} else {
		goto L539
	}
L348:
	;
	v1665 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_30))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L8
	} else {
		goto L537
	}
L349:
	;
	v1659 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_31))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L8
	} else {
		goto L535
	}
L350:
	;
	v1653 = F_slice_from_s(m, l0, int32(12), int32(_a_F_yiddish_UTF_8_stem_32))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L8
	} else {
		goto L533
	}
L351:
	;
	v1647 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_33))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L8
	} else {
		goto L531
	}
L352:
	;
	v1641 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_34))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L8
	} else {
		goto L529
	}
L353:
	;
	v1635 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_35))
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L8
	} else {
		goto L527
	}
L354:
	;
	v1629 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_36))
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L8
	} else {
		goto L525
	}
L355:
	;
	v1623 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_37))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L8
	} else {
		goto L523
	}
L356:
	;
	v1617 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_38))
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L8
	} else {
		goto L521
	}
L357:
	;
	v1611 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_39))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L8
	} else {
		goto L519
	}
L358:
	;
	v1605 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_40))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L8
	} else {
		goto L517
	}
L359:
	;
	v1599 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_41))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L8
	} else {
		goto L515
	}
L360:
	;
	v1593 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_42))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L8
	} else {
		goto L513
	}
L361:
	;
	v1587 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_43))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L8
	} else {
		goto L511
	}
L362:
	;
	v1581 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_44))
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L8
	} else {
		goto L509
	}
L363:
	;
	v1575 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_45))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L8
	} else {
		goto L507
	}
L364:
	;
	v1569 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_46))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L8
	} else {
		goto L505
	}
L365:
	;
	v1563 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_47))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L8
	} else {
		goto L503
	}
L366:
	;
	v1557 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_48))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L8
	} else {
		goto L501
	}
L367:
	;
	v1551 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_49))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L8
	} else {
		goto L499
	}
L368:
	;
	v1545 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_50))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L8
	} else {
		goto L497
	}
L369:
	;
	v1539 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_51))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L8
	} else {
		goto L495
	}
L370:
	;
	v1533 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_52))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L8
	} else {
		goto L493
	}
L371:
	;
	v1527 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_53))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L8
	} else {
		goto L491
	}
L372:
	;
	v1521 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_54))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L8
	} else {
		goto L489
	}
L373:
	;
	v1515 = F_slice_from_s(m, l0, int32(4), int32(_a_F_yiddish_UTF_8_stem_55))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L8
	} else {
		goto L487
	}
L374:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1450 <= v1259 {
		goto L466
	} else {
		goto L467
	}
L375:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1259 < v1276 {
		goto L342
	} else {
		goto L383
	}
L376:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1259 < v1268 {
		goto L342
	} else {
		goto L380
	}
L377:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1259 < v1263 {
		goto L342
	} else {
		goto L378
	}
L378:
	;
	v1265 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1265 {
		goto L342
	} else {
		goto L379
	}
L379:
	;
	v2070 = v1265
	goto L4
L380:
	;
	v1272 = F_slice_from_s(m, l0, int32(4), int32(_a_F_yiddish_UTF_8_stem_56))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L8
	} else {
		goto L381
	}
L381:
	;
	if int32(0) <= v1272 {
		goto L342
	} else {
		goto L382
	}
L382:
	;
	v2070 = v1272
	goto L4
L383:
	;
	v1278 = F_slice_del(m, l0)
	mBase = m.M
	if v1278 < int32(0) {
		v2070 = v1278
		goto L4
	} else {
		goto L384
	}
L384:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1281
	v1286 = F_find_among_b(m, l0, int32(_a_F_yiddish_UTF_8_stem_57), int32(26), int32(0))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L8
	} else {
		goto L385
	}
L385:
	;
	if v1286 == int32(0) {
		goto L342
	} else {
		goto L386
	}
L386:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1290
	switch v1286 - int32(1) {
	case 0:
		goto L412
	case 1:
		goto L411
	case 2:
		goto L410
	case 3:
		goto L409
	case 4:
		goto L408
	case 5:
		goto L407
	case 6:
		goto L406
	case 7:
		goto L405
	case 8:
		goto L404
	case 9:
		goto L403
	case 10:
		goto L402
	case 11:
		goto L401
	case 12:
		goto L400
	case 13:
		goto L399
	case 14:
		goto L398
	case 15:
		goto L397
	case 16:
		goto L396
	case 17:
		goto L395
	case 18:
		goto L394
	case 19:
		goto L393
	case 20:
		goto L392
	case 21:
		goto L391
	case 22:
		goto L390
	case 23:
		goto L389
	case 24:
		goto L388
	case 25:
		goto L387
	default:
		goto L342
	}
L387:
	;
	v1446 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_58))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L8
	} else {
		goto L463
	}
L388:
	;
	v1440 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_59))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L8
	} else {
		goto L461
	}
L389:
	;
	v1434 = F_slice_from_s(m, l0, int32(12), int32(_a_F_yiddish_UTF_8_stem_60))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L8
	} else {
		goto L459
	}
L390:
	;
	v1428 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_61))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L8
	} else {
		goto L457
	}
L391:
	;
	v1422 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_62))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L8
	} else {
		goto L455
	}
L392:
	;
	v1416 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_63))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L8
	} else {
		goto L453
	}
L393:
	;
	v1410 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_64))
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L8
	} else {
		goto L451
	}
L394:
	;
	v1404 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_65))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L8
	} else {
		goto L449
	}
L395:
	;
	v1398 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_66))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L8
	} else {
		goto L447
	}
L396:
	;
	v1392 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_67))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L8
	} else {
		goto L445
	}
L397:
	;
	v1386 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_68))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L8
	} else {
		goto L443
	}
L398:
	;
	v1380 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_69))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L8
	} else {
		goto L441
	}
L399:
	;
	v1374 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_70))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L8
	} else {
		goto L439
	}
L400:
	;
	v1368 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_71))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L8
	} else {
		goto L437
	}
L401:
	;
	v1362 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_72))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L8
	} else {
		goto L435
	}
L402:
	;
	v1356 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_73))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L8
	} else {
		goto L433
	}
L403:
	;
	v1350 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_74))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L8
	} else {
		goto L431
	}
L404:
	;
	v1344 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_75))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L8
	} else {
		goto L429
	}
L405:
	;
	v1338 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_76))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L8
	} else {
		goto L427
	}
L406:
	;
	v1332 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_77))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L8
	} else {
		goto L425
	}
L407:
	;
	v1326 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_78))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L8
	} else {
		goto L423
	}
L408:
	;
	v1320 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_79))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L8
	} else {
		goto L421
	}
L409:
	;
	v1314 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_80))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L8
	} else {
		goto L419
	}
L410:
	;
	v1308 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_81))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L8
	} else {
		goto L417
	}
L411:
	;
	v1302 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_82))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L8
	} else {
		goto L415
	}
L412:
	;
	v1296 = F_slice_from_s(m, l0, int32(4), int32(_a_F_yiddish_UTF_8_stem_83))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L8
	} else {
		goto L413
	}
L413:
	;
	if int32(0) <= v1296 {
		goto L342
	} else {
		goto L414
	}
L414:
	;
	v2070 = v1296
	goto L4
L415:
	;
	if int32(0) <= v1302 {
		goto L342
	} else {
		goto L416
	}
L416:
	;
	v2070 = v1302
	goto L4
L417:
	;
	if int32(0) <= v1308 {
		goto L342
	} else {
		goto L418
	}
L418:
	;
	v2070 = v1308
	goto L4
L419:
	;
	if int32(0) <= v1314 {
		goto L342
	} else {
		goto L420
	}
L420:
	;
	v2070 = v1314
	goto L4
L421:
	;
	if int32(0) <= v1320 {
		goto L342
	} else {
		goto L422
	}
L422:
	;
	v2070 = v1320
	goto L4
L423:
	;
	if int32(0) <= v1326 {
		goto L342
	} else {
		goto L424
	}
L424:
	;
	v2070 = v1326
	goto L4
L425:
	;
	if int32(0) <= v1332 {
		goto L342
	} else {
		goto L426
	}
L426:
	;
	v2070 = v1332
	goto L4
L427:
	;
	if int32(0) <= v1338 {
		goto L342
	} else {
		goto L428
	}
L428:
	;
	v2070 = v1338
	goto L4
L429:
	;
	if int32(0) <= v1344 {
		goto L342
	} else {
		goto L430
	}
L430:
	;
	v2070 = v1344
	goto L4
L431:
	;
	if int32(0) <= v1350 {
		goto L342
	} else {
		goto L432
	}
L432:
	;
	v2070 = v1350
	goto L4
L433:
	;
	if int32(0) <= v1356 {
		goto L342
	} else {
		goto L434
	}
L434:
	;
	v2070 = v1356
	goto L4
L435:
	;
	if int32(0) <= v1362 {
		goto L342
	} else {
		goto L436
	}
L436:
	;
	v2070 = v1362
	goto L4
L437:
	;
	if int32(0) <= v1368 {
		goto L342
	} else {
		goto L438
	}
L438:
	;
	v2070 = v1368
	goto L4
L439:
	;
	if int32(0) <= v1374 {
		goto L342
	} else {
		goto L440
	}
L440:
	;
	v2070 = v1374
	goto L4
L441:
	;
	if int32(0) <= v1380 {
		goto L342
	} else {
		goto L442
	}
L442:
	;
	v2070 = v1380
	goto L4
L443:
	;
	if int32(0) <= v1386 {
		goto L342
	} else {
		goto L444
	}
L444:
	;
	v2070 = v1386
	goto L4
L445:
	;
	if int32(0) <= v1392 {
		goto L342
	} else {
		goto L446
	}
L446:
	;
	v2070 = v1392
	goto L4
L447:
	;
	if int32(0) <= v1398 {
		goto L342
	} else {
		goto L448
	}
L448:
	;
	v2070 = v1398
	goto L4
L449:
	;
	if int32(0) <= v1404 {
		goto L342
	} else {
		goto L450
	}
L450:
	;
	v2070 = v1404
	goto L4
L451:
	;
	if int32(0) <= v1410 {
		goto L342
	} else {
		goto L452
	}
L452:
	;
	v2070 = v1410
	goto L4
L453:
	;
	if int32(0) <= v1416 {
		goto L342
	} else {
		goto L454
	}
L454:
	;
	v2070 = v1416
	goto L4
L455:
	;
	if int32(0) <= v1422 {
		goto L342
	} else {
		goto L456
	}
L456:
	;
	v2070 = v1422
	goto L4
L457:
	;
	if int32(0) <= v1428 {
		goto L342
	} else {
		goto L458
	}
L458:
	;
	v2070 = v1428
	goto L4
L459:
	;
	if int32(0) <= v1434 {
		goto L342
	} else {
		goto L460
	}
L460:
	;
	v2070 = v1434
	goto L4
L461:
	;
	if int32(0) <= v1440 {
		goto L342
	} else {
		goto L462
	}
L462:
	;
	v2070 = v1440
	goto L4
L463:
	;
	if int32(0) <= v1446 {
		goto L342
	} else {
		goto L464
	}
L464:
	;
	v2070 = v1446
	goto L4
L465:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1462
	v1464 = int32(8)
	v1466 = int32(0)
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1462-v1469 < v1464 {
		v1479 = v1466
		goto L473
	} else {
		goto L474
	}
L466:
	;
	v1452 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1452 {
		goto L465
	} else {
		goto L469
	}
L467:
	;
	goto L468
L468:
	;
	v1457 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_84))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L8
	} else {
		goto L470
	}
L469:
	;
	v2070 = v1452
	goto L4
L470:
	;
	if v1457 < int32(0) {
		v2070 = v1457
		goto L4
	} else {
		goto L471
	}
L471:
	;
	goto L465
L472:
	;
	if v1479 == int32(0) {
		goto L342
	} else {
		goto L476
	}
L473:
	;
	goto L472
L474:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1475 = F_memcmp(m, v1472+v1462-v1464, int32(_a_F_yiddish_UTF_8_stem_85), v1464)
	mBase = m.M
	if v1475 != 0 {
		v1479 = v1466
		goto L473
	} else {
		goto L475
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1462 - v1464
	v1479 = int32(1)
	goto L473
L476:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1484 = int32(4)
	v1486 = int32(0)
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1482-v1489 < v1484 {
		v1499 = v1486
		goto L479
	} else {
		goto L480
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1505
	v1509 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_86))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L8
	} else {
		goto L485
	}
L478:
	;
	if v1499 != 0 {
		goto L482
	} else {
		goto L483
	}
L479:
	;
	goto L478
L480:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1495 = F_memcmp(m, v1492+v1482-v1484, int32(_a_F_yiddish_UTF_8_stem_87), v1484)
	mBase = m.M
	if v1495 != 0 {
		v1499 = v1486
		goto L479
	} else {
		goto L481
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1482 - v1484
	v1499 = int32(1)
	goto L479
L482:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1505 = v1500
	goto L477
L483:
	;
	goto L484
L484:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1503 = v1501 + (v1482 - v1483)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1503
	v1505 = v1503
	goto L477
L485:
	;
	if int32(0) <= v1509 {
		goto L342
	} else {
		goto L486
	}
L486:
	;
	v2070 = v1509
	goto L4
L487:
	;
	if int32(0) <= v1515 {
		goto L342
	} else {
		goto L488
	}
L488:
	;
	v2070 = v1515
	goto L4
L489:
	;
	if int32(0) <= v1521 {
		goto L342
	} else {
		goto L490
	}
L490:
	;
	v2070 = v1521
	goto L4
L491:
	;
	if int32(0) <= v1527 {
		goto L342
	} else {
		goto L492
	}
L492:
	;
	v2070 = v1527
	goto L4
L493:
	;
	if int32(0) <= v1533 {
		goto L342
	} else {
		goto L494
	}
L494:
	;
	v2070 = v1533
	goto L4
L495:
	;
	if int32(0) <= v1539 {
		goto L342
	} else {
		goto L496
	}
L496:
	;
	v2070 = v1539
	goto L4
L497:
	;
	if int32(0) <= v1545 {
		goto L342
	} else {
		goto L498
	}
L498:
	;
	v2070 = v1545
	goto L4
L499:
	;
	if int32(0) <= v1551 {
		goto L342
	} else {
		goto L500
	}
L500:
	;
	v2070 = v1551
	goto L4
L501:
	;
	if int32(0) <= v1557 {
		goto L342
	} else {
		goto L502
	}
L502:
	;
	v2070 = v1557
	goto L4
L503:
	;
	if int32(0) <= v1563 {
		goto L342
	} else {
		goto L504
	}
L504:
	;
	v2070 = v1563
	goto L4
L505:
	;
	if int32(0) <= v1569 {
		goto L342
	} else {
		goto L506
	}
L506:
	;
	v2070 = v1569
	goto L4
L507:
	;
	if int32(0) <= v1575 {
		goto L342
	} else {
		goto L508
	}
L508:
	;
	v2070 = v1575
	goto L4
L509:
	;
	if int32(0) <= v1581 {
		goto L342
	} else {
		goto L510
	}
L510:
	;
	v2070 = v1581
	goto L4
L511:
	;
	if int32(0) <= v1587 {
		goto L342
	} else {
		goto L512
	}
L512:
	;
	v2070 = v1587
	goto L4
L513:
	;
	if int32(0) <= v1593 {
		goto L342
	} else {
		goto L514
	}
L514:
	;
	v2070 = v1593
	goto L4
L515:
	;
	if int32(0) <= v1599 {
		goto L342
	} else {
		goto L516
	}
L516:
	;
	v2070 = v1599
	goto L4
L517:
	;
	if int32(0) <= v1605 {
		goto L342
	} else {
		goto L518
	}
L518:
	;
	v2070 = v1605
	goto L4
L519:
	;
	if int32(0) <= v1611 {
		goto L342
	} else {
		goto L520
	}
L520:
	;
	v2070 = v1611
	goto L4
L521:
	;
	if int32(0) <= v1617 {
		goto L342
	} else {
		goto L522
	}
L522:
	;
	v2070 = v1617
	goto L4
L523:
	;
	if int32(0) <= v1623 {
		goto L342
	} else {
		goto L524
	}
L524:
	;
	v2070 = v1623
	goto L4
L525:
	;
	if int32(0) <= v1629 {
		goto L342
	} else {
		goto L526
	}
L526:
	;
	v2070 = v1629
	goto L4
L527:
	;
	if int32(0) <= v1635 {
		goto L342
	} else {
		goto L528
	}
L528:
	;
	v2070 = v1635
	goto L4
L529:
	;
	if int32(0) <= v1641 {
		goto L342
	} else {
		goto L530
	}
L530:
	;
	v2070 = v1641
	goto L4
L531:
	;
	if int32(0) <= v1647 {
		goto L342
	} else {
		goto L532
	}
L532:
	;
	v2070 = v1647
	goto L4
L533:
	;
	if int32(0) <= v1653 {
		goto L342
	} else {
		goto L534
	}
L534:
	;
	v2070 = v1653
	goto L4
L535:
	;
	if int32(0) <= v1659 {
		goto L342
	} else {
		goto L536
	}
L536:
	;
	v2070 = v1659
	goto L4
L537:
	;
	if int32(0) <= v1665 {
		goto L342
	} else {
		goto L538
	}
L538:
	;
	v2070 = v1665
	goto L4
L539:
	;
	if int32(0) <= v1671 {
		goto L342
	} else {
		goto L540
	}
L540:
	;
	v2070 = v1671
	goto L4
L541:
	;
	v1679 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_88))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L8
	} else {
		goto L542
	}
L542:
	;
	if int32(0) <= v1679 {
		goto L342
	} else {
		goto L543
	}
L543:
	;
	v2070 = v1679
	goto L4
L544:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1737 = v1736 - v1703
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1737
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1737 < v1739 {
		goto L342
	} else {
		goto L560
	}
L545:
	;
	if v1699 == int32(0) {
		goto L549
	} else {
		goto L550
	}
L546:
	;
	goto L545
L547:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1695 = F_memcmp(m, v1692+v1688-v1684, int32(_a_F_yiddish_UTF_8_stem_89), v1684)
	mBase = m.M
	if v1695 != 0 {
		v1699 = v1686
		goto L546
	} else {
		goto L548
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1688 - v1684
	v1699 = int32(1)
	goto L546
L549:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1703 = v1683 - v1259
	v1704 = v1702 - v1703
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1704
	v1706 = int32(2)
	v1708 = int32(0)
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1704-v1711 < v1706 {
		v1721 = v1708
		goto L553
	} else {
		goto L554
	}
L550:
	;
	goto L551
L551:
	;
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1726+int32(6) < v1725 {
		goto L342
	} else {
		goto L557
	}
L552:
	;
	if v1721 == int32(0) {
		goto L544
	} else {
		goto L556
	}
L553:
	;
	goto L552
L554:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1717 = F_memcmp(m, v1714+v1704-v1706, int32(_a_F_yiddish_UTF_8_stem_90), v1706)
	mBase = m.M
	if v1717 != 0 {
		v1721 = v1708
		goto L553
	} else {
		goto L555
	}
L555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1704 - v1706
	v1721 = int32(1)
	goto L553
L556:
	;
	goto L551
L557:
	;
	v1732 = F_slice_from_s(m, l0, int32(4), int32(_a_F_yiddish_UTF_8_stem_91))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L8
	} else {
		goto L558
	}
L558:
	;
	if int32(0) <= v1732 {
		goto L342
	} else {
		goto L559
	}
L559:
	;
	v2070 = v1732
	goto L4
L560:
	;
	v1741 = F_slice_del(m, l0)
	mBase = m.M
	if v1741 < int32(0) {
		v2070 = v1741
		goto L4
	} else {
		goto L561
	}
L561:
	;
	goto L342
L562:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1921
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1921
	v1927 = F_find_among_b(m, l0, int32(_a_F_yiddish_UTF_8_stem_92), int32(9), int32(0))
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L8
	} else {
		goto L598
	}
L563:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1754+v1751))))
	if base.B2i32(v1756&int32(224) != int32(128))|base.B2i32(int32(1)<<(uint(v1756)%32)&int32(285474816) == int32(0)) != 0 {
		goto L562
	} else {
		goto L564
	}
L564:
	;
	v1771 = F_find_among_b(m, l0, int32(_a_F_yiddish_UTF_8_stem_93), int32(6), int32(0))
	mBase = m.M
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L8
	} else {
		goto L565
	}
L565:
	;
	if v1771 == int32(0) {
		goto L562
	} else {
		goto L566
	}
L566:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1775
	switch v1771 - int32(1) {
	case 0:
		goto L568
	case 1:
		goto L567
	default:
		goto L562
	}
L567:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1775 < v1784 {
		goto L562
	} else {
		goto L571
	}
L568:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1775 < v1779 {
		goto L562
	} else {
		goto L569
	}
L569:
	;
	v1781 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1781 {
		goto L562
	} else {
		goto L570
	}
L570:
	;
	v2070 = v1781
	goto L4
L571:
	;
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L574
L572:
	;
	if v1914 != 0 {
		goto L562
	} else {
		goto L595
	}
L573:
	;
	v1914 = v1907
	goto L572
L574:
	;
	if v1798 <= v1799 {
		v1907 = int32(-1)
		goto L573
	} else {
		goto L576
	}
L575:
	;
	v1907 = int32(0)
	goto L573
L576:
	;
	v1816 = int32(1)
	v1817 = v1798 - v1816
	v1819 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1800+v1817))))
	v1821 = v1819 & int32(255)
	if base.B2i32(v1817 == v1799)|base.B2i32(int32(0) <= v1819) != 0 {
		v1879 = v1821
		v1883 = v1816
		goto L577
	} else {
		goto L578
	}
L577:
	;
	if int32(1520) < v1879 {
		goto L585
	} else {
		goto L586
	}
L578:
	;
	v1828 = v1821 & int32(63)
	v1830 = v1798 - int32(2)
	v1832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1800+v1830))))
	v1834 = v1832 << (uint(int32(6)) % 32)
	if base.B2i32(v1830 != v1799)&base.B2i32(base.Ui32(v1832) < base.Ui32(int32(192))) == int32(0) {
		goto L579
	} else {
		goto L580
	}
L579:
	;
	v1879 = v1834&int32(1984) | v1828
	v1883 = int32(2)
	goto L577
L580:
	;
	goto L581
L581:
	;
	v1847 = v1834&int32(4032) | v1828
	v1849 = v1798 - int32(3)
	v1851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1800+v1849))))
	if base.B2i32(v1849 != v1799)&base.B2i32(base.Ui32(v1851) < base.Ui32(int32(224))) == int32(0) {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v1879 = v1851<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v1847
	v1883 = int32(3)
	goto L577
L583:
	;
	goto L584
L584:
	;
	v1869 = int32(4)
	v1871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1798+v1800-v1869))))
	v1879 = v1851<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_94) | v1871&int32(7)<<(uint(int32(18))%32) | v1847
	v1883 = v1869
	goto L577
L585:
	;
	v1914 = v1883
	goto L572
L586:
	;
	goto L587
L587:
	;
	v1885 = v1879 - int32(1489)
	if v1885 < int32(0) {
		goto L588
	} else {
		goto L589
	}
L588:
	;
	v1914 = v1883
	goto L572
L589:
	;
	goto L590
L590:
	;
	v1891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1885)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[1]))))
	if int32(base.Ui32(v1891)>>(uint(v1885&int32(7))%32))&int32(1) == int32(0) {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v1914 = v1883
	goto L572
L592:
	;
	goto L593
L593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1798 - v1883
	goto L594
L594:
	;
	goto L575
L595:
	;
	v1915 = F_slice_del(m, l0)
	mBase = m.M
	if v1915 < int32(0) {
		v2070 = v1915
		goto L4
	} else {
		goto L596
	}
L596:
	;
	goto L562
L597:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1945 = v1942
	goto L603
L598:
	;
	if v1927 == int32(0) {
		goto L597
	} else {
		goto L599
	}
L599:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1931
	if v1927 != int32(1) {
		goto L597
	} else {
		goto L600
	}
L600:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1931 < v1935 {
		goto L597
	} else {
		goto L601
	}
L601:
	;
	v1937 = F_slice_del(m, l0)
	mBase = m.M
	if v1937 < int32(0) {
		v2070 = v1937
		goto L4
	} else {
		goto L602
	}
L602:
	;
	goto L597
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1945
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1945
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1951 = v1950 - v1945
	v1952 = int32(2)
	v1954 = int32(0)
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1945-v1957 < v1952 {
		v1967 = v1954
		goto L607
	} else {
		goto L608
	}
L604:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2058
	v2070 = int32(1)
	goto L4
L605:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2000 = v1999 - v1951
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2000
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L621
L606:
	;
	if v1967 == int32(0) {
		goto L610
	} else {
		goto L611
	}
L607:
	;
	goto L606
L608:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1963 = F_memcmp(m, v1960+v1945-v1952, int32(_a_F_yiddish_UTF_8_stem_95), v1952)
	mBase = m.M
	if v1963 != 0 {
		v1967 = v1954
		goto L607
	} else {
		goto L609
	}
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1945 - v1952
	v1967 = int32(1)
	goto L607
L610:
	;
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1972 = v1970 + (v1945 - v1950)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1972
	v1974 = int32(3)
	v1976 = int32(0)
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1972-v1979 < v1974 {
		v1989 = v1976
		goto L614
	} else {
		goto L615
	}
L611:
	;
	goto L612
L612:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1992
	v1994 = F_slice_del(m, l0)
	mBase = m.M
	if v1994 < int32(0) {
		v2070 = v1994
		goto L4
	} else {
		goto L618
	}
L613:
	;
	if v1989 == int32(0) {
		goto L605
	} else {
		goto L617
	}
L614:
	;
	goto L613
L615:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1985 = F_memcmp(m, v1982+v1972-v1974, int32(_a_F_yiddish_UTF_8_stem_96), v1974)
	mBase = m.M
	if v1985 != 0 {
		v1989 = v1976
		goto L614
	} else {
		goto L616
	}
L616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1972 - v1974
	v1989 = int32(1)
	goto L614
L617:
	;
	goto L612
L618:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1945 = v1997 - v1951
	goto L603
L619:
	;
	if int32(0) <= v2055 {
		v1945 = v2055
		goto L603
	} else {
		goto L638
	}
L621:
	;
	goto L622
L622:
	;
	goto L623
L623:
	;
	v2010 = v2000
	v2012 = int32(1)
	goto L626
L625:
	;
	v2055 = v2037
	goto L619
L626:
	;
	if v2010 <= v2003 {
		goto L628
	} else {
		goto L629
	}
L627:
	;
	goto L625
L628:
	;
	v2055 = int32(-1)
	goto L619
L629:
	;
	goto L630
L630:
	;
	v2017 = v2010 - int32(1)
	v2019 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2002+v2017))))
	if base.B2i32(int32(0) <= v2019)|base.B2i32(v2017 <= v2003) != 0 {
		v2037 = v2017
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v2041 = int32(1)
	if v2041 < v2012 {
		v2010 = v2037
		v2012 = v2012 - v2041
		goto L626
	} else {
		goto L637
	}
L632:
	;
	v2025 = v2017
	goto L633
L633:
	;
	v2030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2002+v2025))))
	if base.Ui32(int32(191)) < base.Ui32(v2030) {
		v2037 = v2025
		goto L631
	} else {
		goto L635
	}
L634:
	;
	v2037 = v2003
	goto L631
L635:
	;
	v2034 = v2025 - int32(1)
	if v2003 < v2034 {
		v2025 = v2034
		goto L633
	} else {
		goto L636
	}
L636:
	;
	goto L634
L637:
	;
	goto L627
L638:
	;
	goto L604
L639:
	;
	if v2063 < int32(0) {
		v2070 = v2063
		goto L4
	} else {
		goto L640
	}
L640:
	;
	goto L5
}
func F_yy_fatal_error_4(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	Fn14413(m, l0, int32(_a_F_yy_fatal_error_4_0), int32(37), int32(_a_F_yy_fatal_error_4_1), int32(_a_F_yy_fatal_error_4_2))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
