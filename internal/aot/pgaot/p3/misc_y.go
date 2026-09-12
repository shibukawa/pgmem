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
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v717 int32
	_ = v717
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v755 int32
	_ = v755
	var v762 int32
	_ = v762
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v835 int32
	_ = v835
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v873 int32
	_ = v873
	var v880 int32
	_ = v880
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v953 int32
	_ = v953
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v979 int32
	_ = v979
	var v991 int32
	_ = v991
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1024 int32
	_ = v1024
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1075 int32
	_ = v1075
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1108 int32
	_ = v1108
	var v1119 int32
	_ = v1119
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1199 int32
	_ = v1199
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1225 int32
	_ = v1225
	var v1237 int32
	_ = v1237
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
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
	var v1730 int32
	_ = v1730
	var v1733 int32
	_ = v1733
	var v1737 int32
	_ = v1737
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1892 int32
	_ = v1892
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1905 int32
	_ = v1905
	var v1911 int32
	_ = v1911
	var v1927 int32
	_ = v1927
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1962 int32
	_ = v1962
	var v1965 int32
	_ = v1965
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1987 int32
	_ = v1987
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2009 int32
	_ = v2009
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2038 int32
	_ = v2038
	var v2040 int32
	_ = v2040
	var v2045 int32
	_ = v2045
	var v2050 int32
	_ = v2050
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2090 int32
	_ = v2090
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v6
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v15 = F_find_among(m, l0, int32(4384688), int32(8))
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
	return v2090
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	goto L1
L6:
	;
	v2083 = F_slice_from_s(m, l0, int32(2), int32(2227329))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L8
	} else {
		goto L655
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
	v111 = F_slice_from_s(m, l0, int32(2), int32(2227337))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L45
	}
L12:
	;
	v105 = F_slice_from_s(m, l0, int32(2), int32(2227335))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L43
	}
L13:
	;
	v99 = F_slice_from_s(m, l0, int32(2), int32(2227333))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L41
	}
L14:
	;
	v93 = F_slice_from_s(m, l0, int32(2), int32(2227331))
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
	v35 = F_memcmp(m, v33+v30, int32(2227317), v25)
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
	v43 = F_slice_from_s(m, l0, int32(2), int32(2227319))
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
	v2090 = v43
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
	v57 = F_memcmp(m, v55+v52, int32(2227321), v47)
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
	v65 = F_slice_from_s(m, l0, int32(2), int32(2227323))
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
	v2090 = v65
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
	v79 = F_memcmp(m, v77+v74, int32(2227325), v69)
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
	v87 = F_slice_from_s(m, l0, int32(2), int32(2227327))
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
	v2090 = v87
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
	v2090 = v93
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
	v2090 = v99
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
	v2090 = v105
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
	v2090 = v111
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
	v267 = v252&int32(63) | (v210<<(uint(int32(18))%32)&int32(1835008) | v219<<(uint(int32(12))%32) | v235<<(uint(int32(6))%32))
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
	v267 = v210<<(uint(int32(12))%32)&int32(61440) | v219<<(uint(int32(6))%32) | v235
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
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v272)>>(uint(int32(3))%32)))+uint32(_consts[1295]))))
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
	v2090 = v302
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
	v435 = F_find_among(m, l0, int32(4384848), int32(40))
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
	v379 = F_memcmp(m, v377+v367, int32(2227364), v369)
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
	v426 = F_slice_from_s(m, l0, int32(2), int32(2227376))
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
	v399 = F_memcmp(m, v397+v387, int32(2227368), v389)
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
	v415 = F_memcmp(m, v413+v387, int32(2227372), v405)
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
		v2090 = v426
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
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
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
	v450 = F_memcmp(m, v448+v439, int32(2227378), v440)
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
	v466 = F_memcmp(m, v464+v439, int32(2227386), v456)
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
	v482 = F_memcmp(m, v480+v439, int32(2227394), v472)
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
	v503 = F_memcmp(m, v501+v439, int32(2227402), v493)
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
	v520 = F_memcmp(m, v518+v515, int32(2227410), v510)
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
	v529 = F_slice_from_s(m, l0, int32(2), int32(2227414))
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
	v2090 = v529
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
	v545 = F_memcmp(m, v543+v439, int32(2227416), v535)
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
	v556 = F_slice_from_s(m, l0, int32(3), int32(2227420))
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
		v2090 = v556
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
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1259
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1259
	v1264 = F_find_among_b(m, l0, int32(4385728), int32(79))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L8
	} else {
		goto L348
	}
L187:
	;
	if v621 < int32(0) {
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
	v576 = v568
	v578 = int32(3)
	goto L194
L193:
	;
	v621 = v606
	goto L187
L194:
	;
	if v569 <= v576 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	goto L193
L196:
	;
	v621 = int32(-1)
	goto L187
L197:
	;
	goto L198
L198:
	;
	v583 = v576 + int32(1)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567+v576))))
	if base.Ui32(v585) < base.Ui32(int32(192)) {
		v606 = v583
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v607 = int32(1)
	if v607 < v578 {
		v576 = v606
		v578 = v578 - v607
		goto L194
	} else {
		goto L206
	}
L200:
	;
	if v569 <= v583 {
		v606 = v583
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v592 = v583
	goto L202
L202:
	;
	v595 = int32(*(*int8)(unsafe.Add(mBase, uint32(v567+v592))))
	if int32(-65) < v595 {
		v606 = v592
		goto L199
	} else {
		goto L204
	}
L203:
	;
	v606 = v569
	goto L199
L204:
	;
	v599 = v592 + int32(1)
	if v599 != v569 {
		v592 = v599
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
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = v621
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v568
	v628 = v568 + int32(5)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v629 <= v628 {
		v644 = v568
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L218
L209:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631+v628))))
	if v633&int32(254) != int32(168) {
		v644 = v568
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v640 = F_find_among(m, l0, int32(4385648), int32(4))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L8
	} else {
		goto L211
	}
L211:
	;
	if v640 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v644 = v642
	goto L208
L213:
	;
	goto L214
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v568
	v644 = v568
	goto L208
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v644
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1024 = v644
	goto L293
L216:
	;
	if v762 != 0 {
		goto L215
	} else {
		goto L240
	}
L217:
	;
	v762 = v755
	goto L216
L218:
	;
	if v657 <= v656 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v755 = int32(0)
	goto L217
L220:
	;
	v762 = int32(-1)
	goto L216
L221:
	;
	goto L222
L222:
	;
	v673 = int32(1)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656+v658))))
	if base.Ui32(v675) < base.Ui32(int32(192)) {
		v732 = v675
		v733 = v673
		goto L223
	} else {
		goto L224
	}
L223:
	;
	if int32(1520) < v732 {
		v755 = v733
		goto L217
	} else {
		goto L236
	}
L224:
	;
	v679 = v656 + int32(1)
	if v679 == v657 {
		v732 = v675
		v733 = v673
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679+v658))))
	v684 = v682 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v675) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688+v658))))
	v700 = v698 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v675) {
		goto L232
	} else {
		goto L233
	}
L227:
	;
	v688 = v656 + int32(2)
	if v688 != v657 {
		goto L226
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v732 = v675<<(uint(int32(6))%32)&int32(1984) | v684
	v733 = int32(2)
	goto L223
L230:
	;
	goto L229
L231:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v658+v704))))
	v732 = v717&int32(63) | (v675<<(uint(int32(18))%32)&int32(1835008) | v684<<(uint(int32(12))%32) | v700<<(uint(int32(6))%32))
	v733 = int32(4)
	goto L223
L232:
	;
	v704 = v656 + int32(3)
	if v704 != v657 {
		goto L231
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v732 = v675<<(uint(int32(12))%32)&int32(61440) | v684<<(uint(int32(6))%32) | v700
	v733 = int32(3)
	goto L223
L235:
	;
	goto L234
L236:
	;
	v737 = v732 - int32(1489)
	if v737 < int32(0) {
		v755 = v733
		goto L217
	} else {
		goto L237
	}
L237:
	;
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v737)>>(uint(int32(3))%32)))+uint32(_consts[1296]))))
	if int32(base.Ui32(v743)>>(uint(v737&int32(7))%32))&int32(1) == int32(0) {
		v755 = v733
		goto L217
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v733 + v656
	goto L239
L239:
	;
	goto L219
L240:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L243
L241:
	;
	if v880 != 0 {
		goto L215
	} else {
		goto L265
	}
L242:
	;
	v880 = v873
	goto L241
L243:
	;
	if v775 <= v774 {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v873 = int32(0)
	goto L242
L245:
	;
	v880 = int32(-1)
	goto L241
L246:
	;
	goto L247
L247:
	;
	v791 = int32(1)
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v774+v776))))
	if base.Ui32(v793) < base.Ui32(int32(192)) {
		v850 = v793
		v851 = v791
		goto L248
	} else {
		goto L249
	}
L248:
	;
	if int32(1520) < v850 {
		v873 = v851
		goto L242
	} else {
		goto L261
	}
L249:
	;
	v797 = v774 + int32(1)
	if v797 == v775 {
		v850 = v793
		v851 = v791
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v797+v776))))
	v802 = v800 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v793) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806+v776))))
	v818 = v816 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v793) {
		goto L257
	} else {
		goto L258
	}
L252:
	;
	v806 = v774 + int32(2)
	if v806 != v775 {
		goto L251
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	v850 = v793<<(uint(int32(6))%32)&int32(1984) | v802
	v851 = int32(2)
	goto L248
L255:
	;
	goto L254
L256:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776+v822))))
	v850 = v835&int32(63) | (v793<<(uint(int32(18))%32)&int32(1835008) | v802<<(uint(int32(12))%32) | v818<<(uint(int32(6))%32))
	v851 = int32(4)
	goto L248
L257:
	;
	v822 = v774 + int32(3)
	if v822 != v775 {
		goto L256
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v850 = v793<<(uint(int32(12))%32)&int32(61440) | v802<<(uint(int32(6))%32) | v818
	v851 = int32(3)
	goto L248
L260:
	;
	goto L259
L261:
	;
	v855 = v850 - int32(1489)
	if v855 < int32(0) {
		v873 = v851
		goto L242
	} else {
		goto L262
	}
L262:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v855)>>(uint(int32(3))%32)))+uint32(_consts[1296]))))
	if int32(base.Ui32(v861)>>(uint(v855&int32(7))%32))&int32(1) == int32(0) {
		v873 = v851
		goto L242
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v851 + v774
	goto L264
L264:
	;
	goto L244
L265:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L268
L266:
	;
	if v998 != 0 {
		goto L215
	} else {
		goto L290
	}
L267:
	;
	v998 = v991
	goto L266
L268:
	;
	if v893 <= v892 {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v991 = int32(0)
	goto L267
L270:
	;
	v998 = int32(-1)
	goto L266
L271:
	;
	goto L272
L272:
	;
	v909 = int32(1)
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892+v894))))
	if base.Ui32(v911) < base.Ui32(int32(192)) {
		v968 = v911
		v969 = v909
		goto L273
	} else {
		goto L274
	}
L273:
	;
	if int32(1520) < v968 {
		v991 = v969
		goto L267
	} else {
		goto L286
	}
L274:
	;
	v915 = v892 + int32(1)
	if v915 == v893 {
		v968 = v911
		v969 = v909
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915+v894))))
	v920 = v918 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v911) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924+v894))))
	v936 = v934 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v911) {
		goto L282
	} else {
		goto L283
	}
L277:
	;
	v924 = v892 + int32(2)
	if v924 != v893 {
		goto L276
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v968 = v911<<(uint(int32(6))%32)&int32(1984) | v920
	v969 = int32(2)
	goto L273
L280:
	;
	goto L279
L281:
	;
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v894+v940))))
	v968 = v953&int32(63) | (v911<<(uint(int32(18))%32)&int32(1835008) | v920<<(uint(int32(12))%32) | v936<<(uint(int32(6))%32))
	v969 = int32(4)
	goto L273
L282:
	;
	v940 = v892 + int32(3)
	if v940 != v893 {
		goto L281
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v968 = v911<<(uint(int32(12))%32)&int32(61440) | v920<<(uint(int32(6))%32) | v936
	v969 = int32(3)
	goto L273
L285:
	;
	goto L284
L286:
	;
	v973 = v968 - int32(1489)
	if v973 < int32(0) {
		v991 = v969
		goto L267
	} else {
		goto L287
	}
L287:
	;
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v973)>>(uint(int32(3))%32)))+uint32(_consts[1296]))))
	if int32(base.Ui32(v979)>>(uint(v973&int32(7))%32))&int32(1) == int32(0) {
		v991 = v969
		goto L267
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v969 + v892
	goto L289
L289:
	;
	goto L269
L290:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v999)+4)) = v1000
	goto L186
L291:
	;
	if v1119 < int32(0) {
		goto L186
	} else {
		goto L316
	}
L292:
	;
	v1119 = v1091
	goto L291
L293:
	;
	if v1015 <= v1024 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1119 = int32(-1)
	goto L291
L296:
	;
	goto L297
L297:
	;
	v1031 = int32(1)
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024+v1016))))
	if base.Ui32(v1033) < base.Ui32(int32(192)) {
		v1090 = v1033
		v1091 = v1031
		goto L298
	} else {
		goto L299
	}
L298:
	;
	if int32(1522) < v1090 {
		goto L311
	} else {
		goto L312
	}
L299:
	;
	v1037 = v1024 + int32(1)
	if v1037 == v1015 {
		v1090 = v1033
		v1091 = v1031
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037+v1016))))
	v1042 = v1040 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1033) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1046+v1016))))
	v1058 = v1056 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1033) {
		goto L307
	} else {
		goto L308
	}
L302:
	;
	v1046 = v1024 + int32(2)
	if v1046 != v1015 {
		goto L301
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v1090 = v1033<<(uint(int32(6))%32)&int32(1984) | v1042
	v1091 = int32(2)
	goto L298
L305:
	;
	goto L304
L306:
	;
	v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1016+v1062))))
	v1090 = v1075&int32(63) | (v1033<<(uint(int32(18))%32)&int32(1835008) | v1042<<(uint(int32(12))%32) | v1058<<(uint(int32(6))%32))
	v1091 = int32(4)
	goto L298
L307:
	;
	v1062 = v1024 + int32(3)
	if v1062 != v1015 {
		goto L306
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v1090 = v1033<<(uint(int32(12))%32)&int32(61440) | v1042<<(uint(int32(6))%32) | v1058
	v1091 = int32(3)
	goto L298
L310:
	;
	goto L309
L311:
	;
	v1108 = v1091 + v1024
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1108
	v1024 = v1108
	goto L293
L312:
	;
	v1095 = v1090 - int32(1488)
	if v1095 < int32(0) {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v1101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1095)>>(uint(int32(3))%32)))+uint32(_consts[1297]))))
	if int32(base.Ui32(v1101)>>(uint(v1095&int32(7))%32))&int32(1) != 0 {
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
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L321
L318:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1247)))
	if v1249 < v1248 {
		goto L344
	} else {
		goto L345
	}
L319:
	;
	if v1244 == int32(0) {
		goto L317
	} else {
		goto L343
	}
L320:
	;
	v1244 = v1237
	goto L319
L321:
	;
	if v1139 <= v1138 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1237 = int32(0)
	goto L320
L323:
	;
	v1244 = int32(-1)
	goto L319
L324:
	;
	goto L325
L325:
	;
	v1155 = int32(1)
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138+v1140))))
	if base.Ui32(v1157) < base.Ui32(int32(192)) {
		v1214 = v1157
		v1215 = v1155
		goto L326
	} else {
		goto L327
	}
L326:
	;
	if int32(1522) < v1214 {
		v1237 = v1215
		goto L320
	} else {
		goto L339
	}
L327:
	;
	v1161 = v1138 + int32(1)
	if v1161 == v1139 {
		v1214 = v1157
		v1215 = v1155
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1161+v1140))))
	v1166 = v1164 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1157) {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	v1180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1170+v1140))))
	v1182 = v1180 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1157) {
		goto L335
	} else {
		goto L336
	}
L330:
	;
	v1170 = v1138 + int32(2)
	if v1170 != v1139 {
		goto L329
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1214 = v1157<<(uint(int32(6))%32)&int32(1984) | v1166
	v1215 = int32(2)
	goto L326
L333:
	;
	goto L332
L334:
	;
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1140+v1186))))
	v1214 = v1199&int32(63) | (v1157<<(uint(int32(18))%32)&int32(1835008) | v1166<<(uint(int32(12))%32) | v1182<<(uint(int32(6))%32))
	v1215 = int32(4)
	goto L326
L335:
	;
	v1186 = v1138 + int32(3)
	if v1186 != v1139 {
		goto L334
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	v1214 = v1157<<(uint(int32(12))%32)&int32(61440) | v1166<<(uint(int32(6))%32) | v1182
	v1215 = int32(3)
	goto L326
L338:
	;
	goto L337
L339:
	;
	v1219 = v1214 - int32(1488)
	if v1219 < int32(0) {
		v1237 = v1215
		goto L320
	} else {
		goto L340
	}
L340:
	;
	v1225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1219)>>(uint(int32(3))%32)))+uint32(_consts[1297]))))
	if int32(base.Ui32(v1225)>>(uint(v1219&int32(7))%32))&int32(1) == int32(0) {
		v1237 = v1215
		goto L320
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1215 + v1138
	goto L342
L342:
	;
	goto L322
L343:
	;
	goto L318
L344:
	;
	v1251 = v1248
	goto L346
L345:
	;
	v1251 = v1249
	goto L346
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1247)+4)) = v1251
	goto L186
L347:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1766
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1766
	v1770 = v1766 - int32(1)
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1770 <= v1771 {
		goto L571
	} else {
		goto L572
	}
L348:
	;
	if v1264 == int32(0) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1268
	switch v1264 - int32(1) {
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
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1700 = int32(2)
	v1702 = int32(0)
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1704-v1705 < v1700 {
		v1715 = v1702
		goto L554
	} else {
		goto L555
	}
L351:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+4))
	if v1268 < v1691 {
		goto L347
	} else {
		goto L549
	}
L352:
	;
	v1686 = F_slice_from_s(m, l0, int32(10), int32(2228206))
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L8
	} else {
		goto L547
	}
L353:
	;
	v1680 = F_slice_from_s(m, l0, int32(8), int32(2228198))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L8
	} else {
		goto L545
	}
L354:
	;
	v1674 = F_slice_from_s(m, l0, int32(6), int32(2228192))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L8
	} else {
		goto L543
	}
L355:
	;
	v1668 = F_slice_from_s(m, l0, int32(12), int32(2228180))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L8
	} else {
		goto L541
	}
L356:
	;
	v1662 = F_slice_from_s(m, l0, int32(6), int32(2228174))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L8
	} else {
		goto L539
	}
L357:
	;
	v1656 = F_slice_from_s(m, l0, int32(6), int32(2228168))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L8
	} else {
		goto L537
	}
L358:
	;
	v1650 = F_slice_from_s(m, l0, int32(10), int32(2228158))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L8
	} else {
		goto L535
	}
L359:
	;
	v1644 = F_slice_from_s(m, l0, int32(10), int32(2228148))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L8
	} else {
		goto L533
	}
L360:
	;
	v1638 = F_slice_from_s(m, l0, int32(10), int32(2228138))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L8
	} else {
		goto L531
	}
L361:
	;
	v1632 = F_slice_from_s(m, l0, int32(8), int32(2228130))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L8
	} else {
		goto L529
	}
L362:
	;
	v1626 = F_slice_from_s(m, l0, int32(8), int32(2228122))
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L8
	} else {
		goto L527
	}
L363:
	;
	v1620 = F_slice_from_s(m, l0, int32(8), int32(2228114))
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L8
	} else {
		goto L525
	}
L364:
	;
	v1614 = F_slice_from_s(m, l0, int32(8), int32(2228106))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L8
	} else {
		goto L523
	}
L365:
	;
	v1608 = F_slice_from_s(m, l0, int32(8), int32(2228098))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L8
	} else {
		goto L521
	}
L366:
	;
	v1602 = F_slice_from_s(m, l0, int32(8), int32(2228090))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L8
	} else {
		goto L519
	}
L367:
	;
	v1596 = F_slice_from_s(m, l0, int32(6), int32(2228084))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L8
	} else {
		goto L517
	}
L368:
	;
	v1590 = F_slice_from_s(m, l0, int32(6), int32(2228078))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L8
	} else {
		goto L515
	}
L369:
	;
	v1584 = F_slice_from_s(m, l0, int32(8), int32(2228070))
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L8
	} else {
		goto L513
	}
L370:
	;
	v1578 = F_slice_from_s(m, l0, int32(6), int32(2228064))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L8
	} else {
		goto L511
	}
L371:
	;
	v1572 = F_slice_from_s(m, l0, int32(8), int32(2228056))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L8
	} else {
		goto L509
	}
L372:
	;
	v1566 = F_slice_from_s(m, l0, int32(6), int32(2228050))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L8
	} else {
		goto L507
	}
L373:
	;
	v1560 = F_slice_from_s(m, l0, int32(6), int32(2228044))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L8
	} else {
		goto L505
	}
L374:
	;
	v1554 = F_slice_from_s(m, l0, int32(6), int32(2228038))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L8
	} else {
		goto L503
	}
L375:
	;
	v1548 = F_slice_from_s(m, l0, int32(6), int32(2228032))
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L8
	} else {
		goto L501
	}
L376:
	;
	v1542 = F_slice_from_s(m, l0, int32(8), int32(2228024))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L8
	} else {
		goto L499
	}
L377:
	;
	v1536 = F_slice_from_s(m, l0, int32(6), int32(2228018))
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L8
	} else {
		goto L497
	}
L378:
	;
	v1530 = F_slice_from_s(m, l0, int32(4), int32(2228014))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L8
	} else {
		goto L495
	}
L379:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1463)+4))
	if v1464 <= v1268 {
		goto L473
	} else {
		goto L474
	}
L380:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1288)+4))
	if v1268 < v1289 {
		goto L347
	} else {
		goto L389
	}
L381:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+4))
	if v1268 < v1280 {
		goto L347
	} else {
		goto L386
	}
L382:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+4))
	if v1268 < v1273 {
		goto L347
	} else {
		goto L383
	}
L383:
	;
	v1275 = F_slice_del(m, l0)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L8
	} else {
		goto L384
	}
L384:
	;
	if int32(0) <= v1275 {
		goto L347
	} else {
		goto L385
	}
L385:
	;
	v2090 = v1275
	goto L4
L386:
	;
	v1284 = F_slice_from_s(m, l0, int32(4), int32(2227794))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L8
	} else {
		goto L387
	}
L387:
	;
	if int32(0) <= v1284 {
		goto L347
	} else {
		goto L388
	}
L388:
	;
	v2090 = v1284
	goto L4
L389:
	;
	v1291 = F_slice_del(m, l0)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L8
	} else {
		goto L390
	}
L390:
	;
	if v1291 < int32(0) {
		v2090 = v1291
		goto L4
	} else {
		goto L391
	}
L391:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1295
	v1299 = F_find_among_b(m, l0, int32(4387312), int32(26))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L8
	} else {
		goto L392
	}
L392:
	;
	if v1299 == int32(0) {
		goto L347
	} else {
		goto L393
	}
L393:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1303
	switch v1299 - int32(1) {
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
	v1459 = F_slice_from_s(m, l0, int32(8), int32(2227982))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L8
	} else {
		goto L470
	}
L395:
	;
	v1453 = F_slice_from_s(m, l0, int32(6), int32(2227976))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L8
	} else {
		goto L468
	}
L396:
	;
	v1447 = F_slice_from_s(m, l0, int32(12), int32(2227964))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L8
	} else {
		goto L466
	}
L397:
	;
	v1441 = F_slice_from_s(m, l0, int32(6), int32(2227958))
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L8
	} else {
		goto L464
	}
L398:
	;
	v1435 = F_slice_from_s(m, l0, int32(6), int32(2227952))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L8
	} else {
		goto L462
	}
L399:
	;
	v1429 = F_slice_from_s(m, l0, int32(10), int32(2227942))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L8
	} else {
		goto L460
	}
L400:
	;
	v1423 = F_slice_from_s(m, l0, int32(10), int32(2227932))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L8
	} else {
		goto L458
	}
L401:
	;
	v1417 = F_slice_from_s(m, l0, int32(10), int32(2227922))
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L8
	} else {
		goto L456
	}
L402:
	;
	v1411 = F_slice_from_s(m, l0, int32(8), int32(2227914))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L8
	} else {
		goto L454
	}
L403:
	;
	v1405 = F_slice_from_s(m, l0, int32(8), int32(2227906))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L8
	} else {
		goto L452
	}
L404:
	;
	v1399 = F_slice_from_s(m, l0, int32(8), int32(2227898))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L8
	} else {
		goto L450
	}
L405:
	;
	v1393 = F_slice_from_s(m, l0, int32(8), int32(2227890))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L8
	} else {
		goto L448
	}
L406:
	;
	v1387 = F_slice_from_s(m, l0, int32(8), int32(2227882))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L8
	} else {
		goto L446
	}
L407:
	;
	v1381 = F_slice_from_s(m, l0, int32(8), int32(2227874))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L8
	} else {
		goto L444
	}
L408:
	;
	v1375 = F_slice_from_s(m, l0, int32(8), int32(2227866))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L8
	} else {
		goto L442
	}
L409:
	;
	v1369 = F_slice_from_s(m, l0, int32(6), int32(2227860))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L8
	} else {
		goto L440
	}
L410:
	;
	v1363 = F_slice_from_s(m, l0, int32(6), int32(2227854))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L8
	} else {
		goto L438
	}
L411:
	;
	v1357 = F_slice_from_s(m, l0, int32(8), int32(2227846))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L8
	} else {
		goto L436
	}
L412:
	;
	v1351 = F_slice_from_s(m, l0, int32(6), int32(2227840))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L8
	} else {
		goto L434
	}
L413:
	;
	v1345 = F_slice_from_s(m, l0, int32(8), int32(2227832))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L8
	} else {
		goto L432
	}
L414:
	;
	v1339 = F_slice_from_s(m, l0, int32(6), int32(2227826))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L8
	} else {
		goto L430
	}
L415:
	;
	v1333 = F_slice_from_s(m, l0, int32(6), int32(2227820))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L8
	} else {
		goto L428
	}
L416:
	;
	v1327 = F_slice_from_s(m, l0, int32(6), int32(2227814))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L8
	} else {
		goto L426
	}
L417:
	;
	v1321 = F_slice_from_s(m, l0, int32(6), int32(2227808))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L8
	} else {
		goto L424
	}
L418:
	;
	v1315 = F_slice_from_s(m, l0, int32(6), int32(2227802))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L8
	} else {
		goto L422
	}
L419:
	;
	v1309 = F_slice_from_s(m, l0, int32(4), int32(2227798))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L8
	} else {
		goto L420
	}
L420:
	;
	if int32(0) <= v1309 {
		goto L347
	} else {
		goto L421
	}
L421:
	;
	v2090 = v1309
	goto L4
L422:
	;
	if int32(0) <= v1315 {
		goto L347
	} else {
		goto L423
	}
L423:
	;
	v2090 = v1315
	goto L4
L424:
	;
	if int32(0) <= v1321 {
		goto L347
	} else {
		goto L425
	}
L425:
	;
	v2090 = v1321
	goto L4
L426:
	;
	if int32(0) <= v1327 {
		goto L347
	} else {
		goto L427
	}
L427:
	;
	v2090 = v1327
	goto L4
L428:
	;
	if int32(0) <= v1333 {
		goto L347
	} else {
		goto L429
	}
L429:
	;
	v2090 = v1333
	goto L4
L430:
	;
	if int32(0) <= v1339 {
		goto L347
	} else {
		goto L431
	}
L431:
	;
	v2090 = v1339
	goto L4
L432:
	;
	if int32(0) <= v1345 {
		goto L347
	} else {
		goto L433
	}
L433:
	;
	v2090 = v1345
	goto L4
L434:
	;
	if int32(0) <= v1351 {
		goto L347
	} else {
		goto L435
	}
L435:
	;
	v2090 = v1351
	goto L4
L436:
	;
	if int32(0) <= v1357 {
		goto L347
	} else {
		goto L437
	}
L437:
	;
	v2090 = v1357
	goto L4
L438:
	;
	if int32(0) <= v1363 {
		goto L347
	} else {
		goto L439
	}
L439:
	;
	v2090 = v1363
	goto L4
L440:
	;
	if int32(0) <= v1369 {
		goto L347
	} else {
		goto L441
	}
L441:
	;
	v2090 = v1369
	goto L4
L442:
	;
	if int32(0) <= v1375 {
		goto L347
	} else {
		goto L443
	}
L443:
	;
	v2090 = v1375
	goto L4
L444:
	;
	if int32(0) <= v1381 {
		goto L347
	} else {
		goto L445
	}
L445:
	;
	v2090 = v1381
	goto L4
L446:
	;
	if int32(0) <= v1387 {
		goto L347
	} else {
		goto L447
	}
L447:
	;
	v2090 = v1387
	goto L4
L448:
	;
	if int32(0) <= v1393 {
		goto L347
	} else {
		goto L449
	}
L449:
	;
	v2090 = v1393
	goto L4
L450:
	;
	if int32(0) <= v1399 {
		goto L347
	} else {
		goto L451
	}
L451:
	;
	v2090 = v1399
	goto L4
L452:
	;
	if int32(0) <= v1405 {
		goto L347
	} else {
		goto L453
	}
L453:
	;
	v2090 = v1405
	goto L4
L454:
	;
	if int32(0) <= v1411 {
		goto L347
	} else {
		goto L455
	}
L455:
	;
	v2090 = v1411
	goto L4
L456:
	;
	if int32(0) <= v1417 {
		goto L347
	} else {
		goto L457
	}
L457:
	;
	v2090 = v1417
	goto L4
L458:
	;
	if int32(0) <= v1423 {
		goto L347
	} else {
		goto L459
	}
L459:
	;
	v2090 = v1423
	goto L4
L460:
	;
	if int32(0) <= v1429 {
		goto L347
	} else {
		goto L461
	}
L461:
	;
	v2090 = v1429
	goto L4
L462:
	;
	if int32(0) <= v1435 {
		goto L347
	} else {
		goto L463
	}
L463:
	;
	v2090 = v1435
	goto L4
L464:
	;
	if int32(0) <= v1441 {
		goto L347
	} else {
		goto L465
	}
L465:
	;
	v2090 = v1441
	goto L4
L466:
	;
	if int32(0) <= v1447 {
		goto L347
	} else {
		goto L467
	}
L467:
	;
	v2090 = v1447
	goto L4
L468:
	;
	if int32(0) <= v1453 {
		goto L347
	} else {
		goto L469
	}
L469:
	;
	v2090 = v1453
	goto L4
L470:
	;
	if int32(0) <= v1459 {
		goto L347
	} else {
		goto L471
	}
L471:
	;
	v2090 = v1459
	goto L4
L472:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1477
	v1479 = int32(8)
	v1481 = int32(0)
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1477-v1484 < v1479 {
		v1494 = v1481
		goto L481
	} else {
		goto L482
	}
L473:
	;
	v1466 = F_slice_del(m, l0)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L8
	} else {
		goto L476
	}
L474:
	;
	goto L475
L475:
	;
	v1472 = F_slice_from_s(m, l0, int32(2), int32(2227990))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L8
	} else {
		goto L478
	}
L476:
	;
	if int32(0) <= v1466 {
		goto L472
	} else {
		goto L477
	}
L477:
	;
	v2090 = v1466
	goto L4
L478:
	;
	if v1472 < int32(0) {
		v2090 = v1472
		goto L4
	} else {
		goto L479
	}
L479:
	;
	goto L472
L480:
	;
	if v1494 == int32(0) {
		goto L347
	} else {
		goto L484
	}
L481:
	;
	goto L480
L482:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1490 = F_memcmp(m, v1487+v1477-v1479, int32(2227992), v1479)
	mBase = m.M
	if v1490 != 0 {
		v1494 = v1481
		goto L481
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1477 - v1479
	v1494 = int32(1)
	goto L481
L484:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1499 = int32(4)
	v1501 = int32(0)
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1497-v1504 < v1499 {
		v1514 = v1501
		goto L487
	} else {
		goto L488
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1520
	v1524 = F_slice_from_s(m, l0, int32(10), int32(2228004))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L8
	} else {
		goto L493
	}
L486:
	;
	if v1514 != 0 {
		goto L490
	} else {
		goto L491
	}
L487:
	;
	goto L486
L488:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1510 = F_memcmp(m, v1507+v1497-v1499, int32(2228000), v1499)
	mBase = m.M
	if v1510 != 0 {
		v1514 = v1501
		goto L487
	} else {
		goto L489
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1497 - v1499
	v1514 = int32(1)
	goto L487
L490:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1520 = v1515
	goto L485
L491:
	;
	goto L492
L492:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1518 = v1516 + (v1497 - v1498)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1518
	v1520 = v1518
	goto L485
L493:
	;
	if int32(0) <= v1524 {
		goto L347
	} else {
		goto L494
	}
L494:
	;
	v2090 = v1524
	goto L4
L495:
	;
	if int32(0) <= v1530 {
		goto L347
	} else {
		goto L496
	}
L496:
	;
	v2090 = v1530
	goto L4
L497:
	;
	if int32(0) <= v1536 {
		goto L347
	} else {
		goto L498
	}
L498:
	;
	v2090 = v1536
	goto L4
L499:
	;
	if int32(0) <= v1542 {
		goto L347
	} else {
		goto L500
	}
L500:
	;
	v2090 = v1542
	goto L4
L501:
	;
	if int32(0) <= v1548 {
		goto L347
	} else {
		goto L502
	}
L502:
	;
	v2090 = v1548
	goto L4
L503:
	;
	if int32(0) <= v1554 {
		goto L347
	} else {
		goto L504
	}
L504:
	;
	v2090 = v1554
	goto L4
L505:
	;
	if int32(0) <= v1560 {
		goto L347
	} else {
		goto L506
	}
L506:
	;
	v2090 = v1560
	goto L4
L507:
	;
	if int32(0) <= v1566 {
		goto L347
	} else {
		goto L508
	}
L508:
	;
	v2090 = v1566
	goto L4
L509:
	;
	if int32(0) <= v1572 {
		goto L347
	} else {
		goto L510
	}
L510:
	;
	v2090 = v1572
	goto L4
L511:
	;
	if int32(0) <= v1578 {
		goto L347
	} else {
		goto L512
	}
L512:
	;
	v2090 = v1578
	goto L4
L513:
	;
	if int32(0) <= v1584 {
		goto L347
	} else {
		goto L514
	}
L514:
	;
	v2090 = v1584
	goto L4
L515:
	;
	if int32(0) <= v1590 {
		goto L347
	} else {
		goto L516
	}
L516:
	;
	v2090 = v1590
	goto L4
L517:
	;
	if int32(0) <= v1596 {
		goto L347
	} else {
		goto L518
	}
L518:
	;
	v2090 = v1596
	goto L4
L519:
	;
	if int32(0) <= v1602 {
		goto L347
	} else {
		goto L520
	}
L520:
	;
	v2090 = v1602
	goto L4
L521:
	;
	if int32(0) <= v1608 {
		goto L347
	} else {
		goto L522
	}
L522:
	;
	v2090 = v1608
	goto L4
L523:
	;
	if int32(0) <= v1614 {
		goto L347
	} else {
		goto L524
	}
L524:
	;
	v2090 = v1614
	goto L4
L525:
	;
	if int32(0) <= v1620 {
		goto L347
	} else {
		goto L526
	}
L526:
	;
	v2090 = v1620
	goto L4
L527:
	;
	if int32(0) <= v1626 {
		goto L347
	} else {
		goto L528
	}
L528:
	;
	v2090 = v1626
	goto L4
L529:
	;
	if int32(0) <= v1632 {
		goto L347
	} else {
		goto L530
	}
L530:
	;
	v2090 = v1632
	goto L4
L531:
	;
	if int32(0) <= v1638 {
		goto L347
	} else {
		goto L532
	}
L532:
	;
	v2090 = v1638
	goto L4
L533:
	;
	if int32(0) <= v1644 {
		goto L347
	} else {
		goto L534
	}
L534:
	;
	v2090 = v1644
	goto L4
L535:
	;
	if int32(0) <= v1650 {
		goto L347
	} else {
		goto L536
	}
L536:
	;
	v2090 = v1650
	goto L4
L537:
	;
	if int32(0) <= v1656 {
		goto L347
	} else {
		goto L538
	}
L538:
	;
	v2090 = v1656
	goto L4
L539:
	;
	if int32(0) <= v1662 {
		goto L347
	} else {
		goto L540
	}
L540:
	;
	v2090 = v1662
	goto L4
L541:
	;
	if int32(0) <= v1668 {
		goto L347
	} else {
		goto L542
	}
L542:
	;
	v2090 = v1668
	goto L4
L543:
	;
	if int32(0) <= v1674 {
		goto L347
	} else {
		goto L544
	}
L544:
	;
	v2090 = v1674
	goto L4
L545:
	;
	if int32(0) <= v1680 {
		goto L347
	} else {
		goto L546
	}
L546:
	;
	v2090 = v1680
	goto L4
L547:
	;
	if int32(0) <= v1686 {
		goto L347
	} else {
		goto L548
	}
L548:
	;
	v2090 = v1686
	goto L4
L549:
	;
	v1695 = F_slice_from_s(m, l0, int32(2), int32(2228216))
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L8
	} else {
		goto L550
	}
L550:
	;
	if int32(0) <= v1695 {
		goto L347
	} else {
		goto L551
	}
L551:
	;
	v2090 = v1695
	goto L4
L552:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1754 = v1753 - v1719
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1754
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	if v1754 < v1757 {
		goto L347
	} else {
		goto L568
	}
L553:
	;
	if v1715 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L554:
	;
	goto L553
L555:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1711 = F_memcmp(m, v1708+v1704-v1700, int32(2228218), v1700)
	mBase = m.M
	if v1711 != 0 {
		v1715 = v1702
		goto L554
	} else {
		goto L556
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1704 - v1700
	v1715 = int32(1)
	goto L554
L557:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1719 = v1699 - v1268
	v1720 = v1718 - v1719
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1720
	v1722 = int32(2)
	v1724 = int32(0)
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1720-v1727 < v1722 {
		v1737 = v1724
		goto L561
	} else {
		goto L562
	}
L558:
	;
	goto L559
L559:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1741)+4))
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1743+int32(6) < v1742 {
		goto L347
	} else {
		goto L565
	}
L560:
	;
	if v1737 == int32(0) {
		goto L552
	} else {
		goto L564
	}
L561:
	;
	goto L560
L562:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1733 = F_memcmp(m, v1730+v1720-v1722, int32(2228220), v1722)
	mBase = m.M
	if v1733 != 0 {
		v1737 = v1724
		goto L561
	} else {
		goto L563
	}
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1720 - v1722
	v1737 = int32(1)
	goto L561
L564:
	;
	goto L559
L565:
	;
	v1749 = F_slice_from_s(m, l0, int32(4), int32(2228222))
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L8
	} else {
		goto L566
	}
L566:
	;
	if int32(0) <= v1749 {
		goto L347
	} else {
		goto L567
	}
L567:
	;
	v2090 = v1749
	goto L4
L568:
	;
	v1759 = F_slice_del(m, l0)
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L8
	} else {
		goto L569
	}
L569:
	;
	if v1759 < int32(0) {
		v2090 = v1759
		goto L4
	} else {
		goto L570
	}
L570:
	;
	goto L347
L571:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1941
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1941
	v1946 = F_find_among_b(m, l0, int32(4387968), int32(9))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L8
	} else {
		goto L611
	}
L572:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1773+v1770))))
	if v1775&int32(224) != int32(128) {
		goto L571
	} else {
		goto L573
	}
L573:
	;
	if int32(1)<<(uint(v1775)%32)&int32(285474816) == int32(0) {
		goto L571
	} else {
		goto L574
	}
L574:
	;
	v1788 = F_find_among_b(m, l0, int32(4387840), int32(6))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L8
	} else {
		goto L575
	}
L575:
	;
	if v1788 == int32(0) {
		goto L571
	} else {
		goto L576
	}
L576:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1792
	switch v1788 - int32(1) {
	case 0:
		goto L578
	case 1:
		goto L577
	default:
		goto L571
	}
L577:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1803)+4))
	if v1792 < v1804 {
		goto L571
	} else {
		goto L582
	}
L578:
	;
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1796)+4))
	if v1792 < v1797 {
		goto L571
	} else {
		goto L579
	}
L579:
	;
	v1799 = F_slice_del(m, l0)
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L8
	} else {
		goto L580
	}
L580:
	;
	if int32(0) <= v1799 {
		goto L571
	} else {
		goto L581
	}
L581:
	;
	v2090 = v1799
	goto L4
L582:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L585
L583:
	;
	if v1934 != 0 {
		goto L571
	} else {
		goto L607
	}
L584:
	;
	v1934 = v1927
	goto L583
L585:
	;
	if v1822 <= v1823 {
		v1927 = int32(-1)
		goto L584
	} else {
		goto L587
	}
L586:
	;
	v1927 = int32(0)
	goto L584
L587:
	;
	v1840 = int32(1)
	v1841 = v1822 - v1840
	v1843 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1819+v1841))))
	v1845 = v1843 & int32(255)
	if v1841 == v1823 {
		v1900 = v1845
		v1901 = v1840
		goto L588
	} else {
		goto L589
	}
L588:
	;
	if int32(1520) < v1900 {
		goto L597
	} else {
		goto L598
	}
L589:
	;
	if int32(0) <= v1843 {
		v1900 = v1845
		v1901 = v1840
		goto L588
	} else {
		goto L590
	}
L590:
	;
	v1851 = v1845 & int32(63)
	v1853 = v1822 - int32(2)
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1819+v1853))))
	v1857 = v1855 << (uint(int32(6)) % 32)
	if base.B2i32(v1853 != v1823)&base.B2i32(base.Ui32(v1855) < base.Ui32(int32(192))) == int32(0) {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v1900 = v1857&int32(1984) | v1851
	v1901 = int32(2)
	goto L588
L592:
	;
	goto L593
L593:
	;
	v1870 = v1857&int32(4032) | v1851
	v1872 = v1822 - int32(3)
	v1874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1819+v1872))))
	if base.B2i32(v1872 != v1823)&base.B2i32(base.Ui32(v1874) < base.Ui32(int32(224))) == int32(0) {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v1900 = v1874<<(uint(int32(12))%32)&int32(61440) | v1870
	v1901 = int32(3)
	goto L588
L595:
	;
	goto L596
L596:
	;
	v1892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822+(v1819-int32(4))))))
	v1900 = v1874<<(uint(int32(12))%32)&int32(258048) | v1892&int32(7)<<(uint(int32(18))%32) | v1870
	v1901 = int32(4)
	goto L588
L597:
	;
	v1934 = v1901
	goto L583
L598:
	;
	goto L599
L599:
	;
	v1905 = v1900 - int32(1489)
	if v1905 < int32(0) {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	v1934 = v1901
	goto L583
L601:
	;
	goto L602
L602:
	;
	v1911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1905)>>(uint(int32(3))%32)))+uint32(_consts[1296]))))
	if int32(base.Ui32(v1911)>>(uint(v1905&int32(7))%32))&int32(1) == int32(0) {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v1934 = v1901
	goto L583
L604:
	;
	goto L605
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1822 - v1901
	goto L606
L606:
	;
	goto L586
L607:
	;
	v1935 = F_slice_del(m, l0)
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L8
	} else {
		goto L608
	}
L608:
	;
	if v1935 < int32(0) {
		v2090 = v1935
		goto L4
	} else {
		goto L609
	}
L609:
	;
	goto L571
L610:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1965 = v1962
	goto L617
L611:
	;
	if v1946 == int32(0) {
		goto L610
	} else {
		goto L612
	}
L612:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1950
	if v1946 != int32(1) {
		goto L610
	} else {
		goto L613
	}
L613:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1954)+4))
	if v1950 < v1955 {
		goto L610
	} else {
		goto L614
	}
L614:
	;
	v1957 = F_slice_del(m, l0)
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L8
	} else {
		goto L615
	}
L615:
	;
	if v1957 < int32(0) {
		v2090 = v1957
		goto L4
	} else {
		goto L616
	}
L616:
	;
	goto L610
L617:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1965
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1965
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1971 = v1970 - v1965
	v1972 = int32(2)
	v1974 = int32(0)
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1965-v1977 < v1972 {
		v1987 = v1974
		goto L621
	} else {
		goto L622
	}
L618:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2078
	v2090 = int32(1)
	goto L4
L619:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2021 = v2020 - v1971
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2021
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L636
L620:
	;
	if v1987 == int32(0) {
		goto L624
	} else {
		goto L625
	}
L621:
	;
	goto L620
L622:
	;
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1983 = F_memcmp(m, v1980+v1965-v1972, int32(2228226), v1972)
	mBase = m.M
	if v1983 != 0 {
		v1987 = v1974
		goto L621
	} else {
		goto L623
	}
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1965 - v1972
	v1987 = int32(1)
	goto L621
L624:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1992 = v1990 + (v1965 - v1970)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1992
	v1994 = int32(3)
	v1996 = int32(0)
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1992-v1999 < v1994 {
		v2009 = v1996
		goto L628
	} else {
		goto L629
	}
L625:
	;
	goto L626
L626:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2012
	v2014 = F_slice_del(m, l0)
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L8
	} else {
		goto L632
	}
L627:
	;
	if v2009 == int32(0) {
		goto L619
	} else {
		goto L631
	}
L628:
	;
	goto L627
L629:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2005 = F_memcmp(m, v2002+v1992-v1994, int32(2228228), v1994)
	mBase = m.M
	if v2005 != 0 {
		v2009 = v1996
		goto L628
	} else {
		goto L630
	}
L630:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1992 - v1994
	v2009 = int32(1)
	goto L628
L631:
	;
	goto L626
L632:
	;
	if v2014 < int32(0) {
		v2090 = v2014
		goto L4
	} else {
		goto L633
	}
L633:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1965 = v2018 - v1971
	goto L617
L634:
	;
	if int32(0) <= v2075 {
		v1965 = v2075
		goto L617
	} else {
		goto L654
	}
L636:
	;
	goto L637
L637:
	;
	goto L638
L638:
	;
	v2031 = v2021
	v2033 = int32(1)
	goto L641
L640:
	;
	v2075 = v2057
	goto L634
L641:
	;
	if v2031 <= v2024 {
		goto L643
	} else {
		goto L644
	}
L642:
	;
	goto L640
L643:
	;
	v2075 = int32(-1)
	goto L634
L644:
	;
	goto L645
L645:
	;
	v2038 = v2031 - int32(1)
	v2040 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2023+v2038))))
	if int32(0) <= v2040 {
		v2057 = v2038
		goto L646
	} else {
		goto L647
	}
L646:
	;
	v2061 = int32(1)
	if v2061 < v2033 {
		v2031 = v2057
		v2033 = v2033 - v2061
		goto L641
	} else {
		goto L653
	}
L647:
	;
	if v2038 <= v2024 {
		v2057 = v2038
		goto L646
	} else {
		goto L648
	}
L648:
	;
	v2045 = v2038
	goto L649
L649:
	;
	v2050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2023+v2045))))
	if base.Ui32(int32(191)) < base.Ui32(v2050) {
		v2057 = v2045
		goto L646
	} else {
		goto L651
	}
L650:
	;
	v2057 = v2024
	goto L646
L651:
	;
	v2054 = v2045 - int32(1)
	if v2024 < v2054 {
		v2045 = v2054
		goto L649
	} else {
		goto L652
	}
L652:
	;
	goto L650
L653:
	;
	goto L642
L654:
	;
	goto L618
L655:
	;
	if v2083 < int32(0) {
		v2090 = v2083
		goto L4
	} else {
		goto L656
	}
L656:
	;
	goto L5
}
func F_yy_fatal_error_4(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_errmsg_internal(m, int32(206059), v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errfinish(m, int32(314384), int32(37), int32(80779))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
