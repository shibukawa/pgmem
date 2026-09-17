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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
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
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
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
	var v561 int32
	_ = v561
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
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v715 int32
	_ = v715
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v753 int32
	_ = v753
	var v760 int32
	_ = v760
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v833 int32
	_ = v833
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v951 int32
	_ = v951
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1022 int32
	_ = v1022
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1073 int32
	_ = v1073
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1117 int32
	_ = v1117
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1197 int32
	_ = v1197
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1223 int32
	_ = v1223
	var v1235 int32
	_ = v1235
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1763 int32
	_ = v1763
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1887 int32
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1909 int32
	_ = v1909
	var v1925 int32
	_ = v1925
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1939 int32
	_ = v1939
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2044 int32
	_ = v2044
	var v2049 int32
	_ = v2049
	var v2053 int32
	_ = v2053
	var v2056 int32
	_ = v2056
	var v2060 int32
	_ = v2060
	var v2074 int32
	_ = v2074
	var v2077 int32
	_ = v2077
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2089 int32
	_ = v2089
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v6
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v15 = F_find_among(m, l0, int32(_a_F_yiddish_UTF_8_stem_0), int32(8))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170
	v9 = v170
	goto L1
L4:
	;
	return v2089
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	goto L1
L6:
	;
	v2082 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_1))
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L8
	} else {
		goto L652
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L49
L8:
	;
	return int32(0)
L9:
	;
	if v15 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v21
	switch v15 - int32(1) {
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
	v111 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_2))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L45
	}
L12:
	;
	v105 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_3))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L43
	}
L13:
	;
	v99 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_4))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L41
	}
L14:
	;
	v93 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_5))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L39
	}
L15:
	;
	v69 = int32(2)
	v71 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v73-v74 < v69 {
		v83 = v71
		goto L33
	} else {
		goto L34
	}
L16:
	;
	v47 = int32(2)
	v49 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v51-v52 < v47 {
		v61 = v49
		goto L26
	} else {
		goto L27
	}
L17:
	;
	v25 = int32(2)
	v27 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v29-v30 < v25 {
		v39 = v27
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v39 != 0 {
		goto L7
	} else {
		goto L22
	}
L19:
	;
	goto L18
L20:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = F_memcmp(m, v33+v30, int32(_a_F_yiddish_UTF_8_stem_6), v25)
	mBase = m.M
	if v35 != 0 {
		v39 = v27
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 + v30
	v39 = int32(1)
	goto L19
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
	v43 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_7))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	if int32(0) <= v43 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v2089 = v43
	goto L4
L25:
	;
	if v61 != 0 {
		goto L7
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = F_memcmp(m, v55+v52, int32(_a_F_yiddish_UTF_8_stem_8), v47)
	mBase = m.M
	if v57 != 0 {
		v61 = v49
		goto L26
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v47 + v52
	v61 = int32(1)
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
	v65 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_9))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	if int32(0) <= v65 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v2089 = v65
	goto L4
L32:
	;
	if v83 != 0 {
		goto L7
	} else {
		goto L36
	}
L33:
	;
	goto L32
L34:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = F_memcmp(m, v77+v74, int32(_a_F_yiddish_UTF_8_stem_10), v69)
	mBase = m.M
	if v79 != 0 {
		v83 = v71
		goto L33
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69 + v74
	v83 = int32(1)
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
	v87 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_11))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	if int32(0) <= v87 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v2089 = v87
	goto L4
L39:
	;
	if int32(0) <= v93 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v2089 = v93
	goto L4
L41:
	;
	if int32(0) <= v99 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v2089 = v99
	goto L4
L43:
	;
	if int32(0) <= v105 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v2089 = v105
	goto L4
L45:
	;
	if int32(0) <= v111 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v2089 = v111
	goto L4
L47:
	;
	if int32(0) <= v170 {
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
	v125 = v9
	v127 = int32(1)
	goto L54
L53:
	;
	v170 = v155
	goto L47
L54:
	;
	if v118 <= v125 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v170 = int32(-1)
	goto L47
L57:
	;
	goto L58
L58:
	;
	v132 = v125 + int32(1)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v125))))
	if base.Ui32(v134) < base.Ui32(int32(192)) {
		v155 = v132
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v156 = int32(1)
	if v156 < v127 {
		v125 = v155
		v127 = v127 - v156
		goto L54
	} else {
		goto L66
	}
L60:
	;
	if v118 <= v132 {
		v155 = v132
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v141 = v132
	goto L62
L62:
	;
	v144 = int32(*(*int8)(unsafe.Add(mBase, uint32(v117+v141))))
	if int32(-65) < v144 {
		v155 = v141
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v155 = v118
	goto L59
L64:
	;
	v148 = v141 + int32(1)
	if v148 != v118 {
		v141 = v148
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
	v175 = v6
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v175
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L72
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+4)) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v367
	v369 = int32(4)
	v371 = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v373-v367 < v369 {
		v383 = v371
		goto L122
	} else {
		goto L123
	}
L70:
	;
	if v297 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L71:
	;
	v297 = v290
	goto L70
L72:
	;
	if v192 <= v175 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v290 = int32(0)
	goto L71
L74:
	;
	v297 = int32(-1)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v208 = int32(1)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v193))))
	if base.Ui32(v210) < base.Ui32(int32(192)) {
		v267 = v210
		v268 = v208
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if int32(1474) < v267 {
		v290 = v268
		goto L71
	} else {
		goto L90
	}
L78:
	;
	v214 = v175 + int32(1)
	if v214 == v192 {
		v267 = v210
		v268 = v208
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214+v193))))
	v219 = v217 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v210) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223+v193))))
	v235 = v233 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v210) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v223 = v175 + int32(2)
	if v223 != v192 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v267 = v210<<(uint(int32(6))%32)&int32(1984) | v219
	v268 = int32(2)
	goto L77
L84:
	;
	goto L83
L85:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v239))))
	v267 = v252&int32(63) | (v210<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v219<<(uint(int32(12))%32) | v235<<(uint(int32(6))%32))
	v268 = int32(4)
	goto L77
L86:
	;
	v239 = v175 + int32(3)
	if v239 != v192 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v267 = v210<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v219<<(uint(int32(6))%32) | v235
	v268 = int32(3)
	goto L77
L89:
	;
	goto L88
L90:
	;
	v272 = v267 - int32(1456)
	if v272 < int32(0) {
		v290 = v268
		goto L71
	} else {
		goto L91
	}
L91:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v272)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[0]))))
	if int32(base.Ui32(v278)>>(uint(v272&int32(7))%32))&int32(1) == int32(0) {
		v290 = v268
		goto L71
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v268 + v175
	goto L93
L93:
	;
	goto L73
L94:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v300
	v302 = F_slice_del(m, l0)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L8
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v175
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L101
L97:
	;
	if int32(0) <= v302 {
		goto L68
	} else {
		goto L98
	}
L98:
	;
	v2089 = v302
	goto L4
L99:
	;
	if int32(0) <= v360 {
		v175 = v360
		goto L68
	} else {
		goto L119
	}
L101:
	;
	goto L102
L102:
	;
	goto L103
L103:
	;
	v315 = v175
	v317 = int32(1)
	goto L106
L105:
	;
	v360 = v345
	goto L99
L106:
	;
	if v308 <= v315 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L105
L108:
	;
	v360 = int32(-1)
	goto L99
L109:
	;
	goto L110
L110:
	;
	v322 = v315 + int32(1)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307+v315))))
	if base.Ui32(v324) < base.Ui32(int32(192)) {
		v345 = v322
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v346 = int32(1)
	if v346 < v317 {
		v315 = v345
		v317 = v317 - v346
		goto L106
	} else {
		goto L118
	}
L112:
	;
	if v308 <= v322 {
		v345 = v322
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v331 = v322
	goto L114
L114:
	;
	v334 = int32(*(*int8)(unsafe.Add(mBase, uint32(v307+v331))))
	if int32(-65) < v334 {
		v345 = v331
		goto L111
	} else {
		goto L116
	}
L115:
	;
	v345 = v308
	goto L111
L116:
	;
	v338 = v331 + int32(1)
	if v338 != v308 {
		v331 = v338
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	goto L107
L119:
	;
	goto L69
L120:
	;
	v435 = F_find_among(m, l0, int32(_a_F_yiddish_UTF_8_stem_14), int32(40))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L8
	} else {
		goto L145
	}
L121:
	;
	if v383 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	goto L121
L123:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v379 = F_memcmp(m, v377+v367, int32(_a_F_yiddish_UTF_8_stem_15), v369)
	mBase = m.M
	if v379 != 0 {
		v383 = v371
		goto L122
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v369 + v367
	v383 = int32(1)
	goto L122
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367
	v432 = v367
	goto L120
L126:
	;
	goto L127
L127:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v387
	v389 = int32(4)
	v391 = int32(0)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v393-v387 < v389 {
		v403 = v391
		goto L131
	} else {
		goto L132
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v387
	v426 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_16))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L8
	} else {
		goto L141
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367
	v432 = v367
	goto L120
L130:
	;
	if v403 != 0 {
		goto L129
	} else {
		goto L134
	}
L131:
	;
	goto L130
L132:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v399 = F_memcmp(m, v397+v387, int32(_a_F_yiddish_UTF_8_stem_17), v389)
	mBase = m.M
	if v399 != 0 {
		v403 = v391
		goto L131
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v389 + v387
	v403 = int32(1)
	goto L131
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v387
	v405 = int32(4)
	v407 = int32(0)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v409-v387 < v405 {
		v419 = v407
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v419 != 0 {
		goto L129
	} else {
		goto L139
	}
L136:
	;
	goto L135
L137:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v415 = F_memcmp(m, v413+v387, int32(_a_F_yiddish_UTF_8_stem_18), v405)
	mBase = m.M
	if v415 != 0 {
		v419 = v407
		goto L136
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v405 + v387
	v419 = int32(1)
	goto L136
L139:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v387 < v420 {
		goto L128
	} else {
		goto L140
	}
L140:
	;
	goto L129
L141:
	;
	if v426 < int32(0) {
		v2089 = v426
		goto L4
	} else {
		goto L142
	}
L142:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v432 = v430
	goto L120
L143:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L189
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v561
	goto L143
L145:
	;
	if v435 == int32(0) {
		v561 = v432
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v440 = int32(8)
	v442 = int32(0)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v444-v439 < v440 {
		v454 = v442
		goto L150
	} else {
		goto L151
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v493 = int32(8)
	v495 = int32(0)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v497-v439 < v493 {
		v507 = v495
		goto L166
	} else {
		goto L167
	}
L148:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v489 < v490 {
		goto L147
	} else {
		goto L164
	}
L149:
	;
	if v454 != 0 {
		goto L148
	} else {
		goto L153
	}
L150:
	;
	goto L149
L151:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v450 = F_memcmp(m, v448+v439, int32(_a_F_yiddish_UTF_8_stem_19), v440)
	mBase = m.M
	if v450 != 0 {
		v454 = v442
		goto L150
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v440 + v439
	v454 = int32(1)
	goto L150
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v456 = int32(8)
	v458 = int32(0)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v460-v439 < v456 {
		v470 = v458
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if v470 != 0 {
		goto L148
	} else {
		goto L158
	}
L155:
	;
	goto L154
L156:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v466 = F_memcmp(m, v464+v439, int32(_a_F_yiddish_UTF_8_stem_20), v456)
	mBase = m.M
	if v466 != 0 {
		v470 = v458
		goto L155
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v456 + v439
	v470 = int32(1)
	goto L155
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v472 = int32(8)
	v474 = int32(0)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v476-v439 < v472 {
		v486 = v474
		goto L160
	} else {
		goto L161
	}
L159:
	;
	if v486 == int32(0) {
		goto L147
	} else {
		goto L163
	}
L160:
	;
	goto L159
L161:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v482 = F_memcmp(m, v480+v439, int32(_a_F_yiddish_UTF_8_stem_21), v472)
	mBase = m.M
	if v482 != 0 {
		v486 = v474
		goto L160
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v472 + v439
	v486 = int32(1)
	goto L160
L163:
	;
	goto L148
L164:
	;
	v561 = v439
	goto L144
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	if v507 != 0 {
		goto L143
	} else {
		goto L169
	}
L166:
	;
	goto L165
L167:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v503 = F_memcmp(m, v501+v439, int32(_a_F_yiddish_UTF_8_stem_22), v493)
	mBase = m.M
	if v503 != 0 {
		v507 = v495
		goto L166
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v493 + v439
	v507 = int32(1)
	goto L166
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v439
	v510 = int32(4)
	v512 = int32(0)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v514-v515 < v510 {
		v524 = v512
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v524 != 0 {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	goto L170
L172:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v520 = F_memcmp(m, v518+v515, int32(_a_F_yiddish_UTF_8_stem_23), v510)
	mBase = m.M
	if v520 != 0 {
		v524 = v512
		goto L171
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v510 + v515
	v524 = int32(1)
	goto L171
L174:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v525
	v529 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_24))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L8
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v535 = int32(4)
	v537 = int32(0)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v539-v439 < v535 {
		v549 = v537
		goto L180
	} else {
		goto L181
	}
L177:
	;
	if int32(0) <= v529 {
		goto L143
	} else {
		goto L178
	}
L178:
	;
	v2089 = v529
	goto L4
L179:
	;
	if v549 == int32(0) {
		v561 = v432
		goto L144
	} else {
		goto L183
	}
L180:
	;
	goto L179
L181:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v545 = F_memcmp(m, v543+v439, int32(_a_F_yiddish_UTF_8_stem_25), v535)
	mBase = m.M
	if v545 != 0 {
		v549 = v537
		goto L180
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v535 + v439
	v549 = int32(1)
	goto L180
L183:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v552
	v556 = F_slice_from_s(m, l0, int32(3), int32(_a_F_yiddish_UTF_8_stem_26))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L8
	} else {
		goto L184
	}
L184:
	;
	if v556 < int32(0) {
		v2089 = v556
		goto L4
	} else {
		goto L185
	}
L185:
	;
	goto L143
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1257
	v1262 = F_find_among_b(m, l0, int32(_a_F_yiddish_UTF_8_stem_27), int32(79))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L8
	} else {
		goto L348
	}
L187:
	;
	if v619 < int32(0) {
		goto L186
	} else {
		goto L207
	}
L189:
	;
	goto L190
L190:
	;
	goto L191
L191:
	;
	v574 = v566
	v576 = int32(3)
	goto L194
L193:
	;
	v619 = v604
	goto L187
L194:
	;
	if v567 <= v574 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	goto L193
L196:
	;
	v619 = int32(-1)
	goto L187
L197:
	;
	goto L198
L198:
	;
	v581 = v574 + int32(1)
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+v574))))
	if base.Ui32(v583) < base.Ui32(int32(192)) {
		v604 = v581
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v605 = int32(1)
	if v605 < v576 {
		v574 = v604
		v576 = v576 - v605
		goto L194
	} else {
		goto L206
	}
L200:
	;
	if v567 <= v581 {
		v604 = v581
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v590 = v581
	goto L202
L202:
	;
	v593 = int32(*(*int8)(unsafe.Add(mBase, uint32(v565+v590))))
	if int32(-65) < v593 {
		v604 = v590
		goto L199
	} else {
		goto L204
	}
L203:
	;
	v604 = v567
	goto L199
L204:
	;
	v597 = v590 + int32(1)
	if v597 != v567 {
		v590 = v597
		goto L202
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	goto L195
L207:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = v619
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v566
	v626 = v566 + int32(5)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v627 <= v626 {
		v642 = v566
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L218
L209:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629+v626))))
	if v631&int32(254) != int32(168) {
		v642 = v566
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v638 = F_find_among(m, l0, int32(_a_F_yiddish_UTF_8_stem_28), int32(4))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L8
	} else {
		goto L211
	}
L211:
	;
	if v638 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v642 = v640
	goto L208
L213:
	;
	goto L214
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v566
	v642 = v566
	goto L208
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v642
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1022 = v642
	goto L293
L216:
	;
	if v760 != 0 {
		goto L215
	} else {
		goto L240
	}
L217:
	;
	v760 = v753
	goto L216
L218:
	;
	if v655 <= v654 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v753 = int32(0)
	goto L217
L220:
	;
	v760 = int32(-1)
	goto L216
L221:
	;
	goto L222
L222:
	;
	v671 = int32(1)
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654+v656))))
	if base.Ui32(v673) < base.Ui32(int32(192)) {
		v730 = v673
		v731 = v671
		goto L223
	} else {
		goto L224
	}
L223:
	;
	if int32(1520) < v730 {
		v753 = v731
		goto L217
	} else {
		goto L236
	}
L224:
	;
	v677 = v654 + int32(1)
	if v677 == v655 {
		v730 = v673
		v731 = v671
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677+v656))))
	v682 = v680 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v673) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686+v656))))
	v698 = v696 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v673) {
		goto L232
	} else {
		goto L233
	}
L227:
	;
	v686 = v654 + int32(2)
	if v686 != v655 {
		goto L226
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v730 = v673<<(uint(int32(6))%32)&int32(1984) | v682
	v731 = int32(2)
	goto L223
L230:
	;
	goto L229
L231:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656+v702))))
	v730 = v715&int32(63) | (v673<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v682<<(uint(int32(12))%32) | v698<<(uint(int32(6))%32))
	v731 = int32(4)
	goto L223
L232:
	;
	v702 = v654 + int32(3)
	if v702 != v655 {
		goto L231
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v730 = v673<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v682<<(uint(int32(6))%32) | v698
	v731 = int32(3)
	goto L223
L235:
	;
	goto L234
L236:
	;
	v735 = v730 - int32(1489)
	if v735 < int32(0) {
		v753 = v731
		goto L217
	} else {
		goto L237
	}
L237:
	;
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v735)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[1]))))
	if int32(base.Ui32(v741)>>(uint(v735&int32(7))%32))&int32(1) == int32(0) {
		v753 = v731
		goto L217
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v731 + v654
	goto L239
L239:
	;
	goto L219
L240:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L243
L241:
	;
	if v878 != 0 {
		goto L215
	} else {
		goto L265
	}
L242:
	;
	v878 = v871
	goto L241
L243:
	;
	if v773 <= v772 {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v871 = int32(0)
	goto L242
L245:
	;
	v878 = int32(-1)
	goto L241
L246:
	;
	goto L247
L247:
	;
	v789 = int32(1)
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772+v774))))
	if base.Ui32(v791) < base.Ui32(int32(192)) {
		v848 = v791
		v849 = v789
		goto L248
	} else {
		goto L249
	}
L248:
	;
	if int32(1520) < v848 {
		v871 = v849
		goto L242
	} else {
		goto L261
	}
L249:
	;
	v795 = v772 + int32(1)
	if v795 == v773 {
		v848 = v791
		v849 = v789
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v795+v774))))
	v800 = v798 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v791) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804+v774))))
	v816 = v814 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v791) {
		goto L257
	} else {
		goto L258
	}
L252:
	;
	v804 = v772 + int32(2)
	if v804 != v773 {
		goto L251
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	v848 = v791<<(uint(int32(6))%32)&int32(1984) | v800
	v849 = int32(2)
	goto L248
L255:
	;
	goto L254
L256:
	;
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v774+v820))))
	v848 = v833&int32(63) | (v791<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v800<<(uint(int32(12))%32) | v816<<(uint(int32(6))%32))
	v849 = int32(4)
	goto L248
L257:
	;
	v820 = v772 + int32(3)
	if v820 != v773 {
		goto L256
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v848 = v791<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v800<<(uint(int32(6))%32) | v816
	v849 = int32(3)
	goto L248
L260:
	;
	goto L259
L261:
	;
	v853 = v848 - int32(1489)
	if v853 < int32(0) {
		v871 = v849
		goto L242
	} else {
		goto L262
	}
L262:
	;
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v853)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[1]))))
	if int32(base.Ui32(v859)>>(uint(v853&int32(7))%32))&int32(1) == int32(0) {
		v871 = v849
		goto L242
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v849 + v772
	goto L264
L264:
	;
	goto L244
L265:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L268
L266:
	;
	if v996 != 0 {
		goto L215
	} else {
		goto L290
	}
L267:
	;
	v996 = v989
	goto L266
L268:
	;
	if v891 <= v890 {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v989 = int32(0)
	goto L267
L270:
	;
	v996 = int32(-1)
	goto L266
L271:
	;
	goto L272
L272:
	;
	v907 = int32(1)
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v890+v892))))
	if base.Ui32(v909) < base.Ui32(int32(192)) {
		v966 = v909
		v967 = v907
		goto L273
	} else {
		goto L274
	}
L273:
	;
	if int32(1520) < v966 {
		v989 = v967
		goto L267
	} else {
		goto L286
	}
L274:
	;
	v913 = v890 + int32(1)
	if v913 == v891 {
		v966 = v909
		v967 = v907
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v913+v892))))
	v918 = v916 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v909) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922+v892))))
	v934 = v932 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v909) {
		goto L282
	} else {
		goto L283
	}
L277:
	;
	v922 = v890 + int32(2)
	if v922 != v891 {
		goto L276
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v966 = v909<<(uint(int32(6))%32)&int32(1984) | v918
	v967 = int32(2)
	goto L273
L280:
	;
	goto L279
L281:
	;
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892+v938))))
	v966 = v951&int32(63) | (v909<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v918<<(uint(int32(12))%32) | v934<<(uint(int32(6))%32))
	v967 = int32(4)
	goto L273
L282:
	;
	v938 = v890 + int32(3)
	if v938 != v891 {
		goto L281
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v966 = v909<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v918<<(uint(int32(6))%32) | v934
	v967 = int32(3)
	goto L273
L285:
	;
	goto L284
L286:
	;
	v971 = v966 - int32(1489)
	if v971 < int32(0) {
		v989 = v967
		goto L267
	} else {
		goto L287
	}
L287:
	;
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v971)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[1]))))
	if int32(base.Ui32(v977)>>(uint(v971&int32(7))%32))&int32(1) == int32(0) {
		v989 = v967
		goto L267
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v967 + v890
	goto L289
L289:
	;
	goto L269
L290:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v997)+4)) = v998
	goto L186
L291:
	;
	if v1117 < int32(0) {
		goto L186
	} else {
		goto L316
	}
L292:
	;
	v1117 = v1089
	goto L291
L293:
	;
	if v1013 <= v1022 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1117 = int32(-1)
	goto L291
L296:
	;
	goto L297
L297:
	;
	v1029 = int32(1)
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022+v1014))))
	if base.Ui32(v1031) < base.Ui32(int32(192)) {
		v1088 = v1031
		v1089 = v1029
		goto L298
	} else {
		goto L299
	}
L298:
	;
	if int32(1522) < v1088 {
		goto L311
	} else {
		goto L312
	}
L299:
	;
	v1035 = v1022 + int32(1)
	if v1035 == v1013 {
		v1088 = v1031
		v1089 = v1029
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035+v1014))))
	v1040 = v1038 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1031) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v1054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1044+v1014))))
	v1056 = v1054 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1031) {
		goto L307
	} else {
		goto L308
	}
L302:
	;
	v1044 = v1022 + int32(2)
	if v1044 != v1013 {
		goto L301
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v1088 = v1031<<(uint(int32(6))%32)&int32(1984) | v1040
	v1089 = int32(2)
	goto L298
L305:
	;
	goto L304
L306:
	;
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014+v1060))))
	v1088 = v1073&int32(63) | (v1031<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v1040<<(uint(int32(12))%32) | v1056<<(uint(int32(6))%32))
	v1089 = int32(4)
	goto L298
L307:
	;
	v1060 = v1022 + int32(3)
	if v1060 != v1013 {
		goto L306
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v1088 = v1031<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v1040<<(uint(int32(6))%32) | v1056
	v1089 = int32(3)
	goto L298
L310:
	;
	goto L309
L311:
	;
	v1106 = v1089 + v1022
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1106
	v1022 = v1106
	goto L293
L312:
	;
	v1093 = v1088 - int32(1488)
	if v1093 < int32(0) {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1093)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[2]))))
	if int32(base.Ui32(v1099)>>(uint(v1093&int32(7))%32))&int32(1) != 0 {
		goto L292
	} else {
		goto L314
	}
L314:
	;
	goto L311
L316:
	;
	goto L317
L317:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L321
L318:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1245)))
	if v1247 < v1246 {
		goto L344
	} else {
		goto L345
	}
L319:
	;
	if v1242 == int32(0) {
		goto L317
	} else {
		goto L343
	}
L320:
	;
	v1242 = v1235
	goto L319
L321:
	;
	if v1137 <= v1136 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1235 = int32(0)
	goto L320
L323:
	;
	v1242 = int32(-1)
	goto L319
L324:
	;
	goto L325
L325:
	;
	v1153 = int32(1)
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1136+v1138))))
	if base.Ui32(v1155) < base.Ui32(int32(192)) {
		v1212 = v1155
		v1213 = v1153
		goto L326
	} else {
		goto L327
	}
L326:
	;
	if int32(1522) < v1212 {
		v1235 = v1213
		goto L320
	} else {
		goto L339
	}
L327:
	;
	v1159 = v1136 + int32(1)
	if v1159 == v1137 {
		v1212 = v1155
		v1213 = v1153
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1159+v1138))))
	v1164 = v1162 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1155) {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168+v1138))))
	v1180 = v1178 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1155) {
		goto L335
	} else {
		goto L336
	}
L330:
	;
	v1168 = v1136 + int32(2)
	if v1168 != v1137 {
		goto L329
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1212 = v1155<<(uint(int32(6))%32)&int32(1984) | v1164
	v1213 = int32(2)
	goto L326
L333:
	;
	goto L332
L334:
	;
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138+v1184))))
	v1212 = v1197&int32(63) | (v1155<<(uint(int32(18))%32)&int32(_a_F_yiddish_UTF_8_stem_12) | v1164<<(uint(int32(12))%32) | v1180<<(uint(int32(6))%32))
	v1213 = int32(4)
	goto L326
L335:
	;
	v1184 = v1136 + int32(3)
	if v1184 != v1137 {
		goto L334
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	v1212 = v1155<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v1164<<(uint(int32(6))%32) | v1180
	v1213 = int32(3)
	goto L326
L338:
	;
	goto L337
L339:
	;
	v1217 = v1212 - int32(1488)
	if v1217 < int32(0) {
		v1235 = v1213
		goto L320
	} else {
		goto L340
	}
L340:
	;
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1217)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[2]))))
	if int32(base.Ui32(v1223)>>(uint(v1217&int32(7))%32))&int32(1) == int32(0) {
		v1235 = v1213
		goto L320
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1213 + v1136
	goto L342
L342:
	;
	goto L322
L343:
	;
	goto L318
L344:
	;
	v1249 = v1246
	goto L346
L345:
	;
	v1249 = v1247
	goto L346
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+4)) = v1249
	goto L186
L347:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1763
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1763
	v1767 = v1763 - int32(1)
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1767 <= v1768 {
		goto L571
	} else {
		goto L572
	}
L348:
	;
	if v1262 == int32(0) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1266
	switch v1262 - int32(1) {
	case 0:
		goto L382
	case 1:
		goto L381
	case 2:
		goto L380
	case 3:
		goto L379
	case 4:
		goto L378
	case 5:
		goto L377
	case 6:
		goto L376
	case 7:
		goto L375
	case 8:
		goto L374
	case 9:
		goto L373
	case 10:
		goto L372
	case 11:
		goto L371
	case 12:
		goto L370
	case 13:
		goto L369
	case 14:
		goto L368
	case 15:
		goto L367
	case 16:
		goto L366
	case 17:
		goto L365
	case 18:
		goto L364
	case 19:
		goto L363
	case 20:
		goto L362
	case 21:
		goto L361
	case 22:
		goto L360
	case 23:
		goto L359
	case 24:
		goto L358
	case 25:
		goto L357
	case 26:
		goto L356
	case 27:
		goto L355
	case 28:
		goto L354
	case 29:
		goto L353
	case 30:
		goto L352
	case 31:
		goto L351
	case 32:
		goto L350
	default:
		goto L347
	}
L350:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1698 = int32(2)
	v1700 = int32(0)
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1702-v1703 < v1698 {
		v1713 = v1700
		goto L554
	} else {
		goto L555
	}
L351:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+4))
	if v1266 < v1689 {
		goto L347
	} else {
		goto L549
	}
L352:
	;
	v1684 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_29))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L8
	} else {
		goto L547
	}
L353:
	;
	v1678 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_30))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L8
	} else {
		goto L545
	}
L354:
	;
	v1672 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_31))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L8
	} else {
		goto L543
	}
L355:
	;
	v1666 = F_slice_from_s(m, l0, int32(12), int32(_a_F_yiddish_UTF_8_stem_32))
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L8
	} else {
		goto L541
	}
L356:
	;
	v1660 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_33))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L8
	} else {
		goto L539
	}
L357:
	;
	v1654 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_34))
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L8
	} else {
		goto L537
	}
L358:
	;
	v1648 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_35))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L8
	} else {
		goto L535
	}
L359:
	;
	v1642 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_36))
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L8
	} else {
		goto L533
	}
L360:
	;
	v1636 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_37))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L8
	} else {
		goto L531
	}
L361:
	;
	v1630 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_38))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L8
	} else {
		goto L529
	}
L362:
	;
	v1624 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_39))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L8
	} else {
		goto L527
	}
L363:
	;
	v1618 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_40))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L8
	} else {
		goto L525
	}
L364:
	;
	v1612 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_41))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L8
	} else {
		goto L523
	}
L365:
	;
	v1606 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_42))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L8
	} else {
		goto L521
	}
L366:
	;
	v1600 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_43))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L8
	} else {
		goto L519
	}
L367:
	;
	v1594 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_44))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L8
	} else {
		goto L517
	}
L368:
	;
	v1588 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_45))
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L8
	} else {
		goto L515
	}
L369:
	;
	v1582 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_46))
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L8
	} else {
		goto L513
	}
L370:
	;
	v1576 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_47))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L8
	} else {
		goto L511
	}
L371:
	;
	v1570 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_48))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L8
	} else {
		goto L509
	}
L372:
	;
	v1564 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_49))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L8
	} else {
		goto L507
	}
L373:
	;
	v1558 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_50))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L8
	} else {
		goto L505
	}
L374:
	;
	v1552 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_51))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L8
	} else {
		goto L503
	}
L375:
	;
	v1546 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_52))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L8
	} else {
		goto L501
	}
L376:
	;
	v1540 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_53))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L8
	} else {
		goto L499
	}
L377:
	;
	v1534 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_54))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L8
	} else {
		goto L497
	}
L378:
	;
	v1528 = F_slice_from_s(m, l0, int32(4), int32(_a_F_yiddish_UTF_8_stem_55))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L8
	} else {
		goto L495
	}
L379:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+4))
	if v1462 <= v1266 {
		goto L473
	} else {
		goto L474
	}
L380:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+4))
	if v1266 < v1287 {
		goto L347
	} else {
		goto L389
	}
L381:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+4))
	if v1266 < v1278 {
		goto L347
	} else {
		goto L386
	}
L382:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+4))
	if v1266 < v1271 {
		goto L347
	} else {
		goto L383
	}
L383:
	;
	v1273 = F_slice_del(m, l0)
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L8
	} else {
		goto L384
	}
L384:
	;
	if int32(0) <= v1273 {
		goto L347
	} else {
		goto L385
	}
L385:
	;
	v2089 = v1273
	goto L4
L386:
	;
	v1282 = F_slice_from_s(m, l0, int32(4), int32(_a_F_yiddish_UTF_8_stem_56))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L8
	} else {
		goto L387
	}
L387:
	;
	if int32(0) <= v1282 {
		goto L347
	} else {
		goto L388
	}
L388:
	;
	v2089 = v1282
	goto L4
L389:
	;
	v1289 = F_slice_del(m, l0)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L8
	} else {
		goto L390
	}
L390:
	;
	if v1289 < int32(0) {
		v2089 = v1289
		goto L4
	} else {
		goto L391
	}
L391:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1293
	v1297 = F_find_among_b(m, l0, int32(_a_F_yiddish_UTF_8_stem_57), int32(26))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L8
	} else {
		goto L392
	}
L392:
	;
	if v1297 == int32(0) {
		goto L347
	} else {
		goto L393
	}
L393:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1301
	switch v1297 - int32(1) {
	case 0:
		goto L419
	case 1:
		goto L418
	case 2:
		goto L417
	case 3:
		goto L416
	case 4:
		goto L415
	case 5:
		goto L414
	case 6:
		goto L413
	case 7:
		goto L412
	case 8:
		goto L411
	case 9:
		goto L410
	case 10:
		goto L409
	case 11:
		goto L408
	case 12:
		goto L407
	case 13:
		goto L406
	case 14:
		goto L405
	case 15:
		goto L404
	case 16:
		goto L403
	case 17:
		goto L402
	case 18:
		goto L401
	case 19:
		goto L400
	case 20:
		goto L399
	case 21:
		goto L398
	case 22:
		goto L397
	case 23:
		goto L396
	case 24:
		goto L395
	case 25:
		goto L394
	default:
		goto L347
	}
L394:
	;
	v1457 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_58))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L8
	} else {
		goto L470
	}
L395:
	;
	v1451 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_59))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L8
	} else {
		goto L468
	}
L396:
	;
	v1445 = F_slice_from_s(m, l0, int32(12), int32(_a_F_yiddish_UTF_8_stem_60))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L8
	} else {
		goto L466
	}
L397:
	;
	v1439 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_61))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L8
	} else {
		goto L464
	}
L398:
	;
	v1433 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_62))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L8
	} else {
		goto L462
	}
L399:
	;
	v1427 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_63))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L8
	} else {
		goto L460
	}
L400:
	;
	v1421 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_64))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L8
	} else {
		goto L458
	}
L401:
	;
	v1415 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_65))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L8
	} else {
		goto L456
	}
L402:
	;
	v1409 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_66))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L8
	} else {
		goto L454
	}
L403:
	;
	v1403 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_67))
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L8
	} else {
		goto L452
	}
L404:
	;
	v1397 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_68))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L8
	} else {
		goto L450
	}
L405:
	;
	v1391 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_69))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L8
	} else {
		goto L448
	}
L406:
	;
	v1385 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_70))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L8
	} else {
		goto L446
	}
L407:
	;
	v1379 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_71))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L8
	} else {
		goto L444
	}
L408:
	;
	v1373 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_72))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L8
	} else {
		goto L442
	}
L409:
	;
	v1367 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_73))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L8
	} else {
		goto L440
	}
L410:
	;
	v1361 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_74))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L8
	} else {
		goto L438
	}
L411:
	;
	v1355 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_75))
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L8
	} else {
		goto L436
	}
L412:
	;
	v1349 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_76))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L8
	} else {
		goto L434
	}
L413:
	;
	v1343 = F_slice_from_s(m, l0, int32(8), int32(_a_F_yiddish_UTF_8_stem_77))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L8
	} else {
		goto L432
	}
L414:
	;
	v1337 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_78))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L8
	} else {
		goto L430
	}
L415:
	;
	v1331 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_79))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L8
	} else {
		goto L428
	}
L416:
	;
	v1325 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_80))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L8
	} else {
		goto L426
	}
L417:
	;
	v1319 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_81))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L8
	} else {
		goto L424
	}
L418:
	;
	v1313 = F_slice_from_s(m, l0, int32(6), int32(_a_F_yiddish_UTF_8_stem_82))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L8
	} else {
		goto L422
	}
L419:
	;
	v1307 = F_slice_from_s(m, l0, int32(4), int32(_a_F_yiddish_UTF_8_stem_83))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L8
	} else {
		goto L420
	}
L420:
	;
	if int32(0) <= v1307 {
		goto L347
	} else {
		goto L421
	}
L421:
	;
	v2089 = v1307
	goto L4
L422:
	;
	if int32(0) <= v1313 {
		goto L347
	} else {
		goto L423
	}
L423:
	;
	v2089 = v1313
	goto L4
L424:
	;
	if int32(0) <= v1319 {
		goto L347
	} else {
		goto L425
	}
L425:
	;
	v2089 = v1319
	goto L4
L426:
	;
	if int32(0) <= v1325 {
		goto L347
	} else {
		goto L427
	}
L427:
	;
	v2089 = v1325
	goto L4
L428:
	;
	if int32(0) <= v1331 {
		goto L347
	} else {
		goto L429
	}
L429:
	;
	v2089 = v1331
	goto L4
L430:
	;
	if int32(0) <= v1337 {
		goto L347
	} else {
		goto L431
	}
L431:
	;
	v2089 = v1337
	goto L4
L432:
	;
	if int32(0) <= v1343 {
		goto L347
	} else {
		goto L433
	}
L433:
	;
	v2089 = v1343
	goto L4
L434:
	;
	if int32(0) <= v1349 {
		goto L347
	} else {
		goto L435
	}
L435:
	;
	v2089 = v1349
	goto L4
L436:
	;
	if int32(0) <= v1355 {
		goto L347
	} else {
		goto L437
	}
L437:
	;
	v2089 = v1355
	goto L4
L438:
	;
	if int32(0) <= v1361 {
		goto L347
	} else {
		goto L439
	}
L439:
	;
	v2089 = v1361
	goto L4
L440:
	;
	if int32(0) <= v1367 {
		goto L347
	} else {
		goto L441
	}
L441:
	;
	v2089 = v1367
	goto L4
L442:
	;
	if int32(0) <= v1373 {
		goto L347
	} else {
		goto L443
	}
L443:
	;
	v2089 = v1373
	goto L4
L444:
	;
	if int32(0) <= v1379 {
		goto L347
	} else {
		goto L445
	}
L445:
	;
	v2089 = v1379
	goto L4
L446:
	;
	if int32(0) <= v1385 {
		goto L347
	} else {
		goto L447
	}
L447:
	;
	v2089 = v1385
	goto L4
L448:
	;
	if int32(0) <= v1391 {
		goto L347
	} else {
		goto L449
	}
L449:
	;
	v2089 = v1391
	goto L4
L450:
	;
	if int32(0) <= v1397 {
		goto L347
	} else {
		goto L451
	}
L451:
	;
	v2089 = v1397
	goto L4
L452:
	;
	if int32(0) <= v1403 {
		goto L347
	} else {
		goto L453
	}
L453:
	;
	v2089 = v1403
	goto L4
L454:
	;
	if int32(0) <= v1409 {
		goto L347
	} else {
		goto L455
	}
L455:
	;
	v2089 = v1409
	goto L4
L456:
	;
	if int32(0) <= v1415 {
		goto L347
	} else {
		goto L457
	}
L457:
	;
	v2089 = v1415
	goto L4
L458:
	;
	if int32(0) <= v1421 {
		goto L347
	} else {
		goto L459
	}
L459:
	;
	v2089 = v1421
	goto L4
L460:
	;
	if int32(0) <= v1427 {
		goto L347
	} else {
		goto L461
	}
L461:
	;
	v2089 = v1427
	goto L4
L462:
	;
	if int32(0) <= v1433 {
		goto L347
	} else {
		goto L463
	}
L463:
	;
	v2089 = v1433
	goto L4
L464:
	;
	if int32(0) <= v1439 {
		goto L347
	} else {
		goto L465
	}
L465:
	;
	v2089 = v1439
	goto L4
L466:
	;
	if int32(0) <= v1445 {
		goto L347
	} else {
		goto L467
	}
L467:
	;
	v2089 = v1445
	goto L4
L468:
	;
	if int32(0) <= v1451 {
		goto L347
	} else {
		goto L469
	}
L469:
	;
	v2089 = v1451
	goto L4
L470:
	;
	if int32(0) <= v1457 {
		goto L347
	} else {
		goto L471
	}
L471:
	;
	v2089 = v1457
	goto L4
L472:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1475
	v1477 = int32(8)
	v1479 = int32(0)
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1475-v1482 < v1477 {
		v1492 = v1479
		goto L481
	} else {
		goto L482
	}
L473:
	;
	v1464 = F_slice_del(m, l0)
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L8
	} else {
		goto L476
	}
L474:
	;
	goto L475
L475:
	;
	v1470 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_84))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L8
	} else {
		goto L478
	}
L476:
	;
	if int32(0) <= v1464 {
		goto L472
	} else {
		goto L477
	}
L477:
	;
	v2089 = v1464
	goto L4
L478:
	;
	if v1470 < int32(0) {
		v2089 = v1470
		goto L4
	} else {
		goto L479
	}
L479:
	;
	goto L472
L480:
	;
	if v1492 == int32(0) {
		goto L347
	} else {
		goto L484
	}
L481:
	;
	goto L480
L482:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1488 = F_memcmp(m, v1485+v1475-v1477, int32(_a_F_yiddish_UTF_8_stem_85), v1477)
	mBase = m.M
	if v1488 != 0 {
		v1492 = v1479
		goto L481
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1475 - v1477
	v1492 = int32(1)
	goto L481
L484:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1497 = int32(4)
	v1499 = int32(0)
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1495-v1502 < v1497 {
		v1512 = v1499
		goto L487
	} else {
		goto L488
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1518
	v1522 = F_slice_from_s(m, l0, int32(10), int32(_a_F_yiddish_UTF_8_stem_86))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L8
	} else {
		goto L493
	}
L486:
	;
	if v1512 != 0 {
		goto L490
	} else {
		goto L491
	}
L487:
	;
	goto L486
L488:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1508 = F_memcmp(m, v1505+v1495-v1497, int32(_a_F_yiddish_UTF_8_stem_87), v1497)
	mBase = m.M
	if v1508 != 0 {
		v1512 = v1499
		goto L487
	} else {
		goto L489
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1495 - v1497
	v1512 = int32(1)
	goto L487
L490:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1518 = v1513
	goto L485
L491:
	;
	goto L492
L492:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1516 = v1514 + (v1495 - v1496)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1516
	v1518 = v1516
	goto L485
L493:
	;
	if int32(0) <= v1522 {
		goto L347
	} else {
		goto L494
	}
L494:
	;
	v2089 = v1522
	goto L4
L495:
	;
	if int32(0) <= v1528 {
		goto L347
	} else {
		goto L496
	}
L496:
	;
	v2089 = v1528
	goto L4
L497:
	;
	if int32(0) <= v1534 {
		goto L347
	} else {
		goto L498
	}
L498:
	;
	v2089 = v1534
	goto L4
L499:
	;
	if int32(0) <= v1540 {
		goto L347
	} else {
		goto L500
	}
L500:
	;
	v2089 = v1540
	goto L4
L501:
	;
	if int32(0) <= v1546 {
		goto L347
	} else {
		goto L502
	}
L502:
	;
	v2089 = v1546
	goto L4
L503:
	;
	if int32(0) <= v1552 {
		goto L347
	} else {
		goto L504
	}
L504:
	;
	v2089 = v1552
	goto L4
L505:
	;
	if int32(0) <= v1558 {
		goto L347
	} else {
		goto L506
	}
L506:
	;
	v2089 = v1558
	goto L4
L507:
	;
	if int32(0) <= v1564 {
		goto L347
	} else {
		goto L508
	}
L508:
	;
	v2089 = v1564
	goto L4
L509:
	;
	if int32(0) <= v1570 {
		goto L347
	} else {
		goto L510
	}
L510:
	;
	v2089 = v1570
	goto L4
L511:
	;
	if int32(0) <= v1576 {
		goto L347
	} else {
		goto L512
	}
L512:
	;
	v2089 = v1576
	goto L4
L513:
	;
	if int32(0) <= v1582 {
		goto L347
	} else {
		goto L514
	}
L514:
	;
	v2089 = v1582
	goto L4
L515:
	;
	if int32(0) <= v1588 {
		goto L347
	} else {
		goto L516
	}
L516:
	;
	v2089 = v1588
	goto L4
L517:
	;
	if int32(0) <= v1594 {
		goto L347
	} else {
		goto L518
	}
L518:
	;
	v2089 = v1594
	goto L4
L519:
	;
	if int32(0) <= v1600 {
		goto L347
	} else {
		goto L520
	}
L520:
	;
	v2089 = v1600
	goto L4
L521:
	;
	if int32(0) <= v1606 {
		goto L347
	} else {
		goto L522
	}
L522:
	;
	v2089 = v1606
	goto L4
L523:
	;
	if int32(0) <= v1612 {
		goto L347
	} else {
		goto L524
	}
L524:
	;
	v2089 = v1612
	goto L4
L525:
	;
	if int32(0) <= v1618 {
		goto L347
	} else {
		goto L526
	}
L526:
	;
	v2089 = v1618
	goto L4
L527:
	;
	if int32(0) <= v1624 {
		goto L347
	} else {
		goto L528
	}
L528:
	;
	v2089 = v1624
	goto L4
L529:
	;
	if int32(0) <= v1630 {
		goto L347
	} else {
		goto L530
	}
L530:
	;
	v2089 = v1630
	goto L4
L531:
	;
	if int32(0) <= v1636 {
		goto L347
	} else {
		goto L532
	}
L532:
	;
	v2089 = v1636
	goto L4
L533:
	;
	if int32(0) <= v1642 {
		goto L347
	} else {
		goto L534
	}
L534:
	;
	v2089 = v1642
	goto L4
L535:
	;
	if int32(0) <= v1648 {
		goto L347
	} else {
		goto L536
	}
L536:
	;
	v2089 = v1648
	goto L4
L537:
	;
	if int32(0) <= v1654 {
		goto L347
	} else {
		goto L538
	}
L538:
	;
	v2089 = v1654
	goto L4
L539:
	;
	if int32(0) <= v1660 {
		goto L347
	} else {
		goto L540
	}
L540:
	;
	v2089 = v1660
	goto L4
L541:
	;
	if int32(0) <= v1666 {
		goto L347
	} else {
		goto L542
	}
L542:
	;
	v2089 = v1666
	goto L4
L543:
	;
	if int32(0) <= v1672 {
		goto L347
	} else {
		goto L544
	}
L544:
	;
	v2089 = v1672
	goto L4
L545:
	;
	if int32(0) <= v1678 {
		goto L347
	} else {
		goto L546
	}
L546:
	;
	v2089 = v1678
	goto L4
L547:
	;
	if int32(0) <= v1684 {
		goto L347
	} else {
		goto L548
	}
L548:
	;
	v2089 = v1684
	goto L4
L549:
	;
	v1693 = F_slice_from_s(m, l0, int32(2), int32(_a_F_yiddish_UTF_8_stem_88))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L8
	} else {
		goto L550
	}
L550:
	;
	if int32(0) <= v1693 {
		goto L347
	} else {
		goto L551
	}
L551:
	;
	v2089 = v1693
	goto L4
L552:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1752 = v1751 - v1717
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1752
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+4))
	if v1752 < v1755 {
		goto L347
	} else {
		goto L568
	}
L553:
	;
	if v1713 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L554:
	;
	goto L553
L555:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1709 = F_memcmp(m, v1706+v1702-v1698, int32(_a_F_yiddish_UTF_8_stem_89), v1698)
	mBase = m.M
	if v1709 != 0 {
		v1713 = v1700
		goto L554
	} else {
		goto L556
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1702 - v1698
	v1713 = int32(1)
	goto L554
L557:
	;
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1717 = v1697 - v1266
	v1718 = v1716 - v1717
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1718
	v1720 = int32(2)
	v1722 = int32(0)
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1718-v1725 < v1720 {
		v1735 = v1722
		goto L561
	} else {
		goto L562
	}
L558:
	;
	goto L559
L559:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+4))
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1741+int32(6) < v1740 {
		goto L347
	} else {
		goto L565
	}
L560:
	;
	if v1735 == int32(0) {
		goto L552
	} else {
		goto L564
	}
L561:
	;
	goto L560
L562:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1731 = F_memcmp(m, v1728+v1718-v1720, int32(_a_F_yiddish_UTF_8_stem_90), v1720)
	mBase = m.M
	if v1731 != 0 {
		v1735 = v1722
		goto L561
	} else {
		goto L563
	}
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1718 - v1720
	v1735 = int32(1)
	goto L561
L564:
	;
	goto L559
L565:
	;
	v1747 = F_slice_from_s(m, l0, int32(4), int32(_a_F_yiddish_UTF_8_stem_91))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L8
	} else {
		goto L566
	}
L566:
	;
	if int32(0) <= v1747 {
		goto L347
	} else {
		goto L567
	}
L567:
	;
	v2089 = v1747
	goto L4
L568:
	;
	v1757 = F_slice_del(m, l0)
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L8
	} else {
		goto L569
	}
L569:
	;
	if v1757 < int32(0) {
		v2089 = v1757
		goto L4
	} else {
		goto L570
	}
L570:
	;
	goto L347
L571:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1939
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1939
	v1944 = F_find_among_b(m, l0, int32(_a_F_yiddish_UTF_8_stem_92), int32(9))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L8
	} else {
		goto L609
	}
L572:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1770+v1767))))
	if base.B2i32(v1772&int32(224) != int32(128))|base.B2i32(int32(1)<<(uint(v1772)%32)&int32(285474816) == int32(0)) != 0 {
		goto L571
	} else {
		goto L573
	}
L573:
	;
	v1786 = F_find_among_b(m, l0, int32(_a_F_yiddish_UTF_8_stem_93), int32(6))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L8
	} else {
		goto L574
	}
L574:
	;
	if v1786 == int32(0) {
		goto L571
	} else {
		goto L575
	}
L575:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1790
	switch v1786 - int32(1) {
	case 0:
		goto L577
	case 1:
		goto L576
	default:
		goto L571
	}
L576:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1801)+4))
	if v1790 < v1802 {
		goto L571
	} else {
		goto L581
	}
L577:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+4))
	if v1790 < v1795 {
		goto L571
	} else {
		goto L578
	}
L578:
	;
	v1797 = F_slice_del(m, l0)
	mBase = m.M
	v1798 = m.ExcPending
	if v1798 != 0 {
		goto L8
	} else {
		goto L579
	}
L579:
	;
	if int32(0) <= v1797 {
		goto L571
	} else {
		goto L580
	}
L580:
	;
	v2089 = v1797
	goto L4
L581:
	;
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L584
L582:
	;
	if v1932 != 0 {
		goto L571
	} else {
		goto L605
	}
L583:
	;
	v1932 = v1925
	goto L582
L584:
	;
	if v1816 <= v1817 {
		v1925 = int32(-1)
		goto L583
	} else {
		goto L586
	}
L585:
	;
	v1925 = int32(0)
	goto L583
L586:
	;
	v1834 = int32(1)
	v1835 = v1816 - v1834
	v1837 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1818+v1835))))
	v1839 = v1837 & int32(255)
	if base.B2i32(v1835 == v1817)|base.B2i32(int32(0) <= v1837) != 0 {
		v1897 = v1839
		v1901 = v1834
		goto L587
	} else {
		goto L588
	}
L587:
	;
	if int32(1520) < v1897 {
		goto L595
	} else {
		goto L596
	}
L588:
	;
	v1846 = v1839 & int32(63)
	v1848 = v1816 - int32(2)
	v1850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1818+v1848))))
	v1852 = v1850 << (uint(int32(6)) % 32)
	if base.B2i32(v1848 != v1817)&base.B2i32(base.Ui32(v1850) < base.Ui32(int32(192))) == int32(0) {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v1897 = v1852&int32(1984) | v1846
	v1901 = int32(2)
	goto L587
L590:
	;
	goto L591
L591:
	;
	v1865 = v1852&int32(4032) | v1846
	v1867 = v1816 - int32(3)
	v1869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1818+v1867))))
	if base.B2i32(v1867 != v1817)&base.B2i32(base.Ui32(v1869) < base.Ui32(int32(224))) == int32(0) {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	v1897 = v1869<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_13) | v1865
	v1901 = int32(3)
	goto L587
L593:
	;
	goto L594
L594:
	;
	v1887 = int32(4)
	v1889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1816+v1818-v1887))))
	v1897 = v1869<<(uint(int32(12))%32)&int32(_a_F_yiddish_UTF_8_stem_94) | v1889&int32(7)<<(uint(int32(18))%32) | v1865
	v1901 = v1887
	goto L587
L595:
	;
	v1932 = v1901
	goto L582
L596:
	;
	goto L597
L597:
	;
	v1903 = v1897 - int32(1489)
	if v1903 < int32(0) {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v1932 = v1901
	goto L582
L599:
	;
	goto L600
L600:
	;
	v1909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1903)>>(uint(int32(3))%32)))+uint32(_c_F_yiddish_UTF_8_stem[1]))))
	if int32(base.Ui32(v1909)>>(uint(v1903&int32(7))%32))&int32(1) == int32(0) {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v1932 = v1901
	goto L582
L602:
	;
	goto L603
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1816 - v1901
	goto L604
L604:
	;
	goto L585
L605:
	;
	v1933 = F_slice_del(m, l0)
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L8
	} else {
		goto L606
	}
L606:
	;
	if v1933 < int32(0) {
		v2089 = v1933
		goto L4
	} else {
		goto L607
	}
L607:
	;
	goto L571
L608:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1963 = v1960
	goto L615
L609:
	;
	if v1944 == int32(0) {
		goto L608
	} else {
		goto L610
	}
L610:
	;
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1948
	if v1944 != int32(1) {
		goto L608
	} else {
		goto L611
	}
L611:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1952)+4))
	if v1948 < v1953 {
		goto L608
	} else {
		goto L612
	}
L612:
	;
	v1955 = F_slice_del(m, l0)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L8
	} else {
		goto L613
	}
L613:
	;
	if v1955 < int32(0) {
		v2089 = v1955
		goto L4
	} else {
		goto L614
	}
L614:
	;
	goto L608
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1963
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1963
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1969 = v1968 - v1963
	v1970 = int32(2)
	v1972 = int32(0)
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1963-v1975 < v1970 {
		v1985 = v1972
		goto L619
	} else {
		goto L620
	}
L616:
	;
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2077
	v2089 = int32(1)
	goto L4
L617:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2019 = v2018 - v1969
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2019
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L634
L618:
	;
	if v1985 == int32(0) {
		goto L622
	} else {
		goto L623
	}
L619:
	;
	goto L618
L620:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1981 = F_memcmp(m, v1978+v1963-v1970, int32(_a_F_yiddish_UTF_8_stem_95), v1970)
	mBase = m.M
	if v1981 != 0 {
		v1985 = v1972
		goto L619
	} else {
		goto L621
	}
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1963 - v1970
	v1985 = int32(1)
	goto L619
L622:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1990 = v1988 + (v1963 - v1968)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1990
	v1992 = int32(3)
	v1994 = int32(0)
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1990-v1997 < v1992 {
		v2007 = v1994
		goto L626
	} else {
		goto L627
	}
L623:
	;
	goto L624
L624:
	;
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2010
	v2012 = F_slice_del(m, l0)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L8
	} else {
		goto L630
	}
L625:
	;
	if v2007 == int32(0) {
		goto L617
	} else {
		goto L629
	}
L626:
	;
	goto L625
L627:
	;
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2003 = F_memcmp(m, v2000+v1990-v1992, int32(_a_F_yiddish_UTF_8_stem_96), v1992)
	mBase = m.M
	if v2003 != 0 {
		v2007 = v1994
		goto L626
	} else {
		goto L628
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1990 - v1992
	v2007 = int32(1)
	goto L626
L629:
	;
	goto L624
L630:
	;
	if v2012 < int32(0) {
		v2089 = v2012
		goto L4
	} else {
		goto L631
	}
L631:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1963 = v2016 - v1969
	goto L615
L632:
	;
	if int32(0) <= v2074 {
		v1963 = v2074
		goto L615
	} else {
		goto L651
	}
L634:
	;
	goto L635
L635:
	;
	goto L636
L636:
	;
	v2029 = v2019
	v2031 = int32(1)
	goto L639
L638:
	;
	v2074 = v2056
	goto L632
L639:
	;
	if v2029 <= v2022 {
		goto L641
	} else {
		goto L642
	}
L640:
	;
	goto L638
L641:
	;
	v2074 = int32(-1)
	goto L632
L642:
	;
	goto L643
L643:
	;
	v2036 = v2029 - int32(1)
	v2038 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2021+v2036))))
	if base.B2i32(int32(0) <= v2038)|base.B2i32(v2036 <= v2022) != 0 {
		v2056 = v2036
		goto L644
	} else {
		goto L645
	}
L644:
	;
	v2060 = int32(1)
	if v2060 < v2031 {
		v2029 = v2056
		v2031 = v2031 - v2060
		goto L639
	} else {
		goto L650
	}
L645:
	;
	v2044 = v2036
	goto L646
L646:
	;
	v2049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2021+v2044))))
	if base.Ui32(int32(191)) < base.Ui32(v2049) {
		v2056 = v2044
		goto L644
	} else {
		goto L648
	}
L647:
	;
	v2056 = v2022
	goto L644
L648:
	;
	v2053 = v2044 - int32(1)
	if v2022 < v2053 {
		v2044 = v2053
		goto L646
	} else {
		goto L649
	}
L649:
	;
	goto L647
L650:
	;
	goto L640
L651:
	;
	goto L616
L652:
	;
	if v2082 < int32(0) {
		v2089 = v2082
		goto L4
	} else {
		goto L653
	}
L653:
	;
	goto L5
}
func F_yy_fatal_error_4(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	Fn14024(m, l0, int32(_a_F_yy_fatal_error_4_0), int32(37), int32(_a_F_yy_fatal_error_4_1), int32(_a_F_yy_fatal_error_4_2))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
