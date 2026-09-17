package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CheckPointGuts(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v196 int32
	_ = v196
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int64
	_ = v359
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v383 int32
	_ = v383
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
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
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int64
	_ = v497
	var v498 int64
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int64
	_ = v779
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int64
	_ = v790
	var v792 int64
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int64
	_ = v993
	var v994 int64
	_ = v994
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int64
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1082 int64
	_ = v1082
	var v1087 int64
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1119 int32
	_ = v1119
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1138 int32
	_ = v1138
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1190 int32
	_ = v1190
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1223 int32
	_ = v1223
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1317 int64
	_ = v1317
	var v1329 int32
	_ = v1329
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1363 int32
	_ = v1363
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1465 int64
	_ = v1465
	var v1471 int32
	_ = v1471
	var v1472 float64
	_ = v1472
	var v1473 float64
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1515 float64
	_ = v1515
	var v1516 float64
	_ = v1516
	var v1518 float64
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1534 int64
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int64
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int64
	_ = v1538
	var v1540 int64
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1548 int32
	_ = v1548
	var v1550 float64
	_ = v1550
	var v1556 int32
	_ = v1556
	var v1560 int64
	_ = v1560
	var v1562 int64
	_ = v1562
	var v1567 int32
	_ = v1567
	var v1569 float64
	_ = v1569
	var v1573 float64
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1585 int32
	_ = v1585
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1649 int32
	_ = v1649
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1718 int64
	_ = v1718
	var v1719 int64
	_ = v1719
	var v1729 int64
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1772 int32
	_ = v1772
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1824 int64
	_ = v1824
	var v1825 int64
	_ = v1825
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1851 int32
	_ = v1851
	var v1852 int64
	_ = v1852
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1867 int32
	_ = v1867
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1920 int32
	_ = v1920
	var v1937 int32
	_ = v1937
	var v1938 int64
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1954 int32
	_ = v1954
	var v1957 int32
	_ = v1957
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1973 int32
	_ = v1973
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1990 int32
	_ = v1990
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2042 int64
	_ = v2042
	var v2046 int32
	_ = v2046
	var v2050 int64
	_ = v2050
	var v2053 int64
	_ = v2053
	var v2059 int64
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2064 int32
	_ = v2064
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2083 int32
	_ = v2083
	var v2088 int32
	_ = v2088
	var v2089 int64
	_ = v2089
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2102 int64
	_ = v2102
	var v2103 int64
	_ = v2103
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2127 int64
	_ = v2127
	var v2128 int64
	_ = v2128
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2142 int32
	_ = v2142
	var v2148 int64
	_ = v2148
	var v2149 int64
	_ = v2149
	var v2161 int32
	_ = v2161
	var v2169 int32
	_ = v2169
	var v2173 int32
	_ = v2173
	var v2178 int32
	_ = v2178
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2191 int32
	_ = v2191
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2201 int64
	_ = v2201
	var v2202 int64
	_ = v2202
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2223 int32
	_ = v2223
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2261 int32
	_ = v2261
	var v2262 int64
	_ = v2262
	var v2264 int64
	_ = v2264
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2279 int64
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2286 int64
	_ = v2286
	var v2293 int64
	_ = v2293
	var v2298 int64
	_ = v2298
	var v2301 int64
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2335 int32
	_ = v2335
	var v2337 int32
	_ = v2337
	var v2345 int32
	_ = v2345
	var v2350 int32
	_ = v2350
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2358 int32
	_ = v2358
	var v2367 int32
	_ = v2367
	var v2369 int32
	_ = v2369
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2388 int32
	_ = v2388
	var v2393 int32
	_ = v2393
	var v2398 int32
	_ = v2398
	var v2402 int32
	_ = v2402
	var v2407 int32
	_ = v2407
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2417 int32
	_ = v2417
	var v2418 int64
	_ = v2418
	var v2423 int32
	_ = v2423
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2441 int32
	_ = v2441
	var v2456 int32
	_ = v2456
	var v2460 int32
	_ = v2460
	var v2464 int32
	_ = v2464
	var v2466 int32
	_ = v2466
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2482 int32
	_ = v2482
	var v2487 int32
	_ = v2487
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2521 int32
	_ = v2521
	var v2526 int32
	_ = v2526
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2540 int32
	_ = v2540
	var v2545 int32
	_ = v2545
	var v2549 int32
	_ = v2549
	var v2551 int32
	_ = v2551
	var v2559 int32
	_ = v2559
	var v2564 int32
	_ = v2564
	v3 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v23 = F_LWLockAcquire(m, v19+int32(3200), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v26+int32(3200))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = m.G0
	v33 = v31 - int32(1040)
	m.G0 = v33
	v37 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v37 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_0), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v48 = int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v55 = F_LWLockAcquire(m, v51+int32(_a_F_CheckPointGuts_1), v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_2), int32(2124), int32(_a_F_CheckPointGuts_3))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[1]))
	if int32(0) < v58 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	m.G0 = v33 + int32(1040)
	v174 = m.G0
	v176 = v174 - int32(1152)
	m.G0 = v176
	v178 = F_GetRedoRecPtr(m)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L37
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[2]))
	v65 = v3
	v66 = v62
	v67 = v3
	v69 = v58
	goto L15
L13:
	;
	goto L14
L14:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v149+int32(_a_F_CheckPointGuts_1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L36
	}
L15:
	;
	v82 = v66 + v65*int32(288)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	if v83 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v137+int32(_a_F_CheckPointGuts_1))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L33
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(_a_F_CheckPointGuts_4)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v82 + int32(24)
	v94 = F_pg_sprintf(m, v33+int32(16), int32(_a_F_CheckPointGuts_5), v33)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v130 = v66
	v131 = v67
	v132 = v69
	goto L19
L19:
	;
	v134 = v65 + int32(1)
	if v134 < v132 {
		v65 = v134
		v66 = v130
		v67 = v131
		v69 = v132
		goto L15
	} else {
		goto L32
	}
L20:
	;
	if l1&v48 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v82)+104))
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v82)+280))
	F_SaveSlotToPath(m, v82, v33+int32(16), int32(15))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v82)+88))
	if v98 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(1)
	if v101 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_s_lock(m, v82, int32(_a_F_CheckPointGuts_2), int32(2162), int32(_a_F_CheckPointGuts_3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v82)+112))
	if v109 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
	goto L21
L29:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v82)+120))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v82)+264))
	if base.Ui64(v110) <= base.Ui64(v111) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v113 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+12)) = uint16(v113)
	goto L28
L31:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[1]))
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[2]))
	v130 = v129
	v131 = base.B2i32(v117 != v118) | v67
	v132 = v125
	goto L19
L32:
	;
	goto L16
L33:
	;
	if v131&int32(1) == int32(0) {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L11
L36:
	;
	goto L11
L37:
	;
	v180 = F_ReplicationSlotsComputeLogicalRestartLSN(m)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v183 = F_AllocateDir(m, int32(_a_F_CheckPointGuts_6))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v186 = F_ReadDir(m, v183, int32(_a_F_CheckPointGuts_6))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v186 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if base.Ui64(v178) < base.Ui64(v180) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	F_FreeDir(m, v183)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L86
	}
L44:
	;
	v189 = v178
	goto L46
L45:
	;
	v189 = v180
	goto L46
L46:
	;
	v196 = v186
	goto L47
L47:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+19)))
	if v209 != int32(46) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L43
L49:
	;
	v331 = F_ReadDir(m, v183, int32(_a_F_CheckPointGuts_6))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L84
	}
L50:
	;
	v222 = v196 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v176)+84)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v176)+80)) = int32(_a_F_CheckPointGuts_6)
	v227 = v176 + int32(96)
	v232 = F_pg_snprintf(m, v227, int32(1045), int32(_a_F_CheckPointGuts_5), v176+int32(80))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L55
	}
L51:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+20)))
	if v212 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+20)))
	if v215 != int32(46) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+21)))
	if v218 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v237 = F_get_dirent_type(m, v227, v196, int32(0), int32(14))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_7), v324, int32(_a_F_CheckPointGuts_8))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L83
	}
L57:
	;
	if v237&int32(-3) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v243 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+52)) = v176 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v176)+48)) = v176 + int32(92)
	v263 = F_sscanf(m, v222, int32(_a_F_CheckPointGuts_9), v176+int32(48))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	if v243 == int32(0) {
		goto L49
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+64)) = v227
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_10), v176-int32(-64))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v324 = int32(2009)
	goto L56
L64:
	;
	if v263 != int32(2) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v269 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v282 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v176)+88)))
	v283 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v176)+92)))
	if base.Ui64(v189-int64(1)) < base.Ui64(v282|v283<<(uint(int64(32))%64)) {
		goto L49
	} else {
		goto L71
	}
L68:
	;
	if v269 == int32(0) {
		goto L49
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+32)) = v176 + int32(96)
	F_errmsg(m, int32(_a_F_CheckPointGuts_11), v176+int32(32))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v324 = int32(2025)
	goto L56
L71:
	;
	v290 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v290 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+16)) = v176 + int32(96)
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_12), v176+int32(16))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v306 = v176 + int32(96)
	v307 = F_unlink(m, v306)
	mBase = m.M
	if int32(0) <= v307 {
		goto L49
	} else {
		goto L78
	}
L76:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_7), int32(2034), int32(_a_F_CheckPointGuts_8))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v312 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v312 == int32(0) {
		goto L49
	} else {
		goto L80
	}
L80:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v306
	F_errmsg(m, int32(_a_F_CheckPointGuts_13), v176)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v324 = int32(2046)
	goto L56
L83:
	;
	goto L49
L84:
	;
	if v331 != 0 {
		v196 = v331
		goto L47
	} else {
		goto L85
	}
L85:
	;
	goto L48
L86:
	;
	m.G0 = v176 + int32(1152)
	v355 = m.G0
	v357 = v355 - int32(1216)
	m.G0 = v357
	v359 = F_GetRedoRecPtr(m)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v361 = F_ReplicationSlotsComputeLogicalRestartLSN(m)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v364 = F_AllocateDir(m, int32(_a_F_CheckPointGuts_14))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L93
	}
L89:
	;
	v683 = m.G0
	v685 = v683 - int32(112)
	m.G0 = v685
	*(*int32)(unsafe.Add(mBase, uint32(v685)+108)) = int32(307747550)
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[3]))
	if v690 != 0 {
		goto L183
	} else {
		goto L184
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L175
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L171
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L168
	}
L93:
	;
	v367 = F_ReadDir(m, v364, int32(_a_F_CheckPointGuts_14))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v367 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if base.Ui64(v359) < base.Ui64(v361) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	F_FreeDir(m, v364)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L166
	}
L98:
	;
	v370 = v359
	goto L100
L99:
	;
	v370 = v361
	goto L100
L100:
	;
	v383 = v367
	goto L101
L101:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+19)))
	if v394 != int32(46) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L97
L103:
	;
	v603 = F_ReadDir(m, v364, int32(_a_F_CheckPointGuts_14))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L164
	}
L104:
	;
	v407 = v383 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v357)+132)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v357)+128)) = int32(_a_F_CheckPointGuts_14)
	v412 = v357 + int32(160)
	v417 = F_pg_snprintf(m, v412, int32(1044), int32(_a_F_CheckPointGuts_5), v357+int32(128))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L109
	}
L105:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+20)))
	if v397 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+20)))
	if v400 != int32(46) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+21)))
	if v403 == int32(0) {
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	v421 = F_get_dirent_type(m, v412, v383, int32(0), int32(14))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v421&int32(-3) != 0 {
		goto L103
	} else {
		goto L111
	}
L111:
	;
	v425 = int32(_a_F_CheckPointGuts_15)
	goto L114
L112:
	;
	if v463-v464 != 0 {
		goto L103
	} else {
		goto L125
	}
L114:
	;
	goto L115
L115:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407))))
	if v432 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v433 = v407
	v434 = v425
	v435 = int32(4)
	v436 = v432
	goto L120
L117:
	;
	v459 = v425
	v463 = int32(0)
	goto L118
L118:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459))))
	goto L112
L119:
	;
	v459 = v454
	v463 = v456
	goto L118
L120:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	if base.B2i32(v436 != v438)|base.B2i32(v438 == int32(0)) != 0 {
		v454 = v434
		v456 = v436
		goto L119
	} else {
		goto L122
	}
L121:
	;
	v454 = v448
	v456 = int32(0)
	goto L119
L122:
	;
	v444 = v435 - int32(1)
	if v444 == int32(0) {
		v454 = v434
		v456 = v436
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v447 = int32(1)
	v448 = v434 + v447
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+1)))
	if v449 != 0 {
		v433 = v433 + v447
		v434 = v448
		v435 = v444
		v436 = v449
		goto L120
	} else {
		goto L124
	}
L124:
	;
	goto L121
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357+int32(116)))) = v357 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v357+int32(112)))) = v357 + int32(148)
	*(*int32)(unsafe.Add(mBase, uint32(v357)+108)) = v357 + int32(136)
	*(*int32)(unsafe.Add(mBase, uint32(v357)+104)) = v357 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v357)+100)) = v357 + int32(152)
	*(*int32)(unsafe.Add(mBase, uint32(v357)+96)) = v357 + int32(156)
	v493 = F_sscanf(m, v407, int32(_a_F_CheckPointGuts_16), v357+int32(96))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v493 != int32(6) {
		goto L92
	} else {
		goto L127
	}
L127:
	;
	v497 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v357)+136)))
	v498 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v357)+140)))
	if base.Ui64(v497|v498<<(uint(int64(32))%64)) <= base.Ui64(v370-int64(1)) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v505 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v541 = v357 + int32(160)
	v543 = F_OpenTransientFile(m, v541, int32(2))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L142
	}
L131:
	;
	if v505 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+64)) = v412
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_17), v357-int32(-64))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v519 = v357 + int32(160)
	v520 = F_unlink(m, v519)
	mBase = m.M
	if int32(0) <= v520 {
		goto L103
	} else {
		goto L137
	}
L135:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1210), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+48)) = v519
	F_errmsg(m, int32(_a_F_CheckPointGuts_13), v357+int32(48))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1214), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	if v543 < int32(0) {
		goto L91
	} else {
		goto L143
	}
L143:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v548))) = int32(167772194)
	v553 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[5])))
	if v553 != int32(1) {
		v567 = int32(0)
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v594 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v594))) = int32(0)
	v597 = F_CloseTransientFile(m, v543)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L162
	}
L145:
	;
	if v567 == int32(0) {
		goto L144
	} else {
		goto L152
	}
L146:
	;
	goto L145
L147:
	;
	goto L148
L148:
	;
	v558 = F_fsync(m, v543)
	mBase = m.M
	if v558 != int32(-1) {
		v567 = v558
		goto L146
	} else {
		goto L150
	}
L149:
	;
	v567 = int32(-1)
	goto L146
L150:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v562 == int32(27) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])))
	if v573 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v576 = F_errstart(m, v574, int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L157
	}
L154:
	;
	v574 = int32(21)
	goto L156
L155:
	;
	v574 = int32(23)
	goto L156
L156:
	;
	goto L153
L157:
	;
	if v576 == int32(0) {
		goto L144
	} else {
		goto L158
	}
L158:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+32)) = v541
	F_errmsg(m, int32(_a_F_CheckPointGuts_20), v357+int32(32))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1240), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	goto L144
L162:
	;
	if v597 != 0 {
		goto L90
	} else {
		goto L163
	}
L163:
	;
	goto L103
L164:
	;
	if v603 != 0 {
		v383 = v603
		goto L101
	} else {
		goto L165
	}
L165:
	;
	goto L102
L166:
	;
	F_fsync_fname(m, int32(_a_F_CheckPointGuts_14), int32(1))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	m.G0 = v357 + int32(1216)
	goto L89
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+80)) = v407
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_21), v357+int32(80))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1204), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v357 + int32(160)
	F_errmsg(m, int32(_a_F_CheckPointGuts_22), v357)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1229), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+16)) = v357 + int32(160)
	F_errmsg(m, int32(_a_F_CheckPointGuts_23), v357+int32(16))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1246), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v988 = m.G0
	v989 = int32(16)
	v990 = v988 - v989
	m.G0 = v990
	F_gettimeofday(m, v990)
	mBase = m.M
	v993 = *(*int64)(unsafe.Add(mBase, uint32(v990)))
	v994 = int64(*(*int32)(unsafe.Add(mBase, uint32(v990)+8)))
	m.G0 = v990 + v989
	goto L251
L180:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L247
	}
L181:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L243
	}
L182:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L239
	}
L183:
	;
	v692 = F_unlink(m, int32(_a_F_CheckPointGuts_24))
	mBase = m.M
	if v692 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	goto L185
L185:
	;
	m.G0 = v685 + int32(112)
	goto L179
L186:
	;
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v696 != int32(44) {
		goto L182
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v701 = F_OpenTransientFile(m, int32(_a_F_CheckPointGuts_24), int32(193))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L190
	}
L189:
	;
	goto L188
L190:
	;
	if v701 < int32(0) {
		goto L181
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(0)
	v710 = int32(4)
	v711 = F_write(m, v701, v685+int32(108), v710)
	mBase = m.M
	if v711 != v710 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v715 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L194
L194:
	;
	v743 = m.Env.Pgmem_crc32c(m, int32(-1), v685+int32(108), int32(4))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v685)+104)) = v743
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v750 = F_LWLockAcquire(m, v746+int32(_a_F_CheckPointGuts_25), int32(1))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L202
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L197
L196:
	;
	goto L197
L197:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+64)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_26), v685-int32(-64))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(639), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	v753 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[3]))
	if int32(0) < v753 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[8]))
	v761 = int32(0)
	v762 = v743
	v763 = v757
	v764 = v753
	goto L206
L204:
	;
	v849 = v743
	goto L205
L205:
	;
	v863 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v863+int32(_a_F_CheckPointGuts_25))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L225
	}
L206:
	;
	v777 = v763 + v761*int32(56)
	v778 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v777))))
	if v778 != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v849 = v839
	goto L205
L208:
	;
	v779 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v685)+96)) = v779
	*(*int64)(unsafe.Add(mBase, uint32(v685)+88)) = v779
	v784 = v777 + int32(40)
	v786 = F_LWLockAcquire(m, v784, int32(1))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	v839 = v762
	v840 = v763
	v841 = v764
	goto L210
L210:
	;
	v843 = v761 + int32(1)
	if v843 < v841 {
		v761 = v843
		v762 = v839
		v763 = v840
		v764 = v841
		goto L206
	} else {
		goto L224
	}
L211:
	;
	v788 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v777))))
	*(*uint16)(unsafe.Add(mBase, uint32(v685)+88)) = uint16(v788)
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v777)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v685)+96)) = v790
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v777)+16))
	F_LWLockRelease(m, v784)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_XLogFlush(m, v792)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(0)
	v802 = int32(16)
	v803 = F_write(m, v701, v685+int32(88), v802)
	mBase = m.M
	if v803 != v802 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v807 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v807 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	v834 = m.Env.Pgmem_crc32c(m, v762, v685+int32(88), int32(16))
	mBase = m.M
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[8]))
	v838 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[3]))
	v839 = v834
	v840 = v836
	v841 = v838
	goto L210
L217:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L219
L218:
	;
	goto L219
L219:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+48)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_26), v685+int32(48))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(681), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	goto L207
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+104)) = v849 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(0)
	v876 = int32(4)
	v877 = F_write(m, v701, v685+int32(104), v876)
	mBase = m.M
	if v877 != v876 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v881 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v881 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	goto L228
L228:
	;
	v905 = F_CloseTransientFile(m, v701)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L236
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L231
L230:
	;
	goto L231
L231:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+32)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_26), v685+int32(32))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(700), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	if v905 != 0 {
		goto L180
	} else {
		goto L237
	}
L237:
	;
	v910 = F_durable_rename(m, int32(_a_F_CheckPointGuts_24), int32(_a_F_CheckPointGuts_29), int32(23))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	goto L185
L239:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+80)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_13), v685+int32(80))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(615), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_30), v685)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(627), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+16)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_23), v685+int32(16))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(707), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[9])) = v994 + v993*int64(1000000) - int64(946684800000000)
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_31))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_32))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_33))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_34))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_35))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v1024 = F_LWLockAcquire(m, v1020+int32(_a_F_CheckPointGuts_36), int32(0))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[10]))
	v1028 = *(*int64)(unsafe.Add(mBase, uint32(v1027)))
	if v1028 < int64(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1104 = int32(0)
	v1107 = m.G0
	v1109 = v1107 - int32(_a_F_CheckPointGuts_37)
	m.G0 = v1109
	v1112 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[11]))
	if v1112 <= v1104 {
		goto L281
	} else {
		goto L282
	}
L259:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v1032+int32(_a_F_CheckPointGuts_36))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+12))
	if v1037 != 0 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L258
L263:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v1089+int32(_a_F_CheckPointGuts_36))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L278
	}
L264:
	;
	v1040 = int32(4)
	v1041 = v1037&int32(-1024) | v1040
	v1044 = base.I32_wrap_i64(v1028) << (uint(int32(10)) % 32)
	v1046 = v1044 | v1040
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1046))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1041)) == int32(0) {
		goto L269
	} else {
		goto L270
	}
L265:
	;
	goto L266
L266:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1027))) = int64(-1)
	v1087 = v1028
	goto L263
L267:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[10]))
	v1082 = *(*int64)(unsafe.Add(mBase, uint32(v1081)))
	v1087 = v1082
	goto L263
L268:
	;
	if v1058 == int32(0) {
		goto L267
	} else {
		goto L272
	}
L269:
	;
	v1058 = base.B2i32(base.Ui32(v1041) < base.Ui32(v1046))
	goto L268
L270:
	;
	goto L271
L271:
	;
	v1058 = int32(base.Ui32(v1041-v1046) >> (uint(int32(31)) % 32))
	goto L268
L272:
	;
	v1062 = v1044 + int32(1027)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1062))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1041)) == int32(0) {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	if v1074 == int32(0) {
		goto L267
	} else {
		goto L277
	}
L274:
	;
	v1074 = base.B2i32(base.Ui32(v1041) < base.Ui32(v1062))
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1074 = int32(base.Ui32(v1041-v1062) >> (uint(int32(31)) % 32))
	goto L273
L277:
	;
	v1087 = base.I64_extend_i32_u(int32(base.Ui32(v1037) >> (uint(int32(10)) % 32)))
	goto L263
L278:
	;
	F_SimpleLruTruncate(m, int32(_a_F_CheckPointGuts_38), v1087)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_38))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	goto L258
L281:
	;
	m.G0 = v1109 + int32(_a_F_CheckPointGuts_37)
	v1713 = m.G0
	v1714 = int32(16)
	v1715 = v1713 - v1714
	m.G0 = v1715
	F_gettimeofday(m, v1715)
	mBase = m.M
	v1718 = *(*int64)(unsafe.Add(mBase, uint32(v1715)))
	v1719 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1715)+8)))
	m.G0 = v1715 + v1714
	goto L413
L282:
	;
	if l1&int32(19) != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1119 = int32(-8388609)
	goto L285
L284:
	;
	v1119 = int32(2139095039)
	goto L285
L285:
	;
	v1128 = v1104
	v1129 = v1104
	goto L286
L286:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+28)) = int32(_a_F_CheckPointGuts_39)
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+24)) = int32(_a_F_CheckPointGuts_40)
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+20)) = int32(_a_F_CheckPointGuts_41)
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1109)+8)) = int64(0)
	v1151 = v1138 + v1128<<(uint(int32(6))%32)
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+24))
	v1153 = int32(_a_F_CheckPointGuts_42)
	*(*int32)(unsafe.Add(mBase, uint32(v1151)+24)) = v1152 | v1153
	if v1152&v1153 != 0 {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	if v1247 == int32(0) {
		goto L281
	} else {
		goto L314
	}
L288:
	;
	goto L291
L289:
	;
	v1190 = v1152
	goto L290
L290:
	;
	v1205 = int32(_a_F_CheckPointGuts_43)
	v1206 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[13]))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1109+int32(8))+8))
	if v1208 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L291:
	;
	F_perform_spin_delay(m, v1109+int32(8))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L293
	}
L292:
	;
	v1190 = v1179
	goto L290
L293:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+24))
	v1180 = int32(_a_F_CheckPointGuts_42)
	*(*int32)(unsafe.Add(mBase, uint32(v1151)+24)) = v1179 | v1180
	if v1179&v1180 != 0 {
		goto L291
	} else {
		goto L294
	}
L294:
	;
	goto L292
L295:
	;
	if v1190|v1119 == int32(-1) {
		goto L306
	} else {
		goto L307
	}
L296:
	;
	goto L295
L297:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[13])) = v1223
	goto L296
L298:
	;
	if int32(999) < v1206 {
		goto L296
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	if v1206 < int32(11) {
		goto L296
	} else {
		goto L305
	}
L301:
	;
	v1213 = int32(900)
	if v1213 <= v1206 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1216 = v1213
	goto L304
L303:
	;
	v1216 = v1206
	goto L304
L304:
	;
	v1223 = v1216 + int32(100)
	goto L297
L305:
	;
	v1223 = v1206 - int32(1)
	goto L297
L306:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14]))
	v1232 = v1229 + v1129*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+16)) = v1128
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1151)))
	*(*int32)(unsafe.Add(mBase, uint32(v1232))) = v1234
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+4)) = v1236
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+8)) = v1238
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+12)) = v1240
	v1247 = v1129 + int32(1)
	v1248 = v1190 | int32(1077936128)
	goto L308
L307:
	;
	v1247 = v1129
	v1248 = v1190
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1151)+24)) = v1248 & int32(-4194305)
	v1253 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	if v1253 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v1257 = v1128 + int32(1)
	v1259 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[11]))
	if v1257 < v1259 {
		v1128 = v1257
		v1129 = v1247
		goto L286
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	goto L287
L314:
	;
	v1263 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+12)) = v1263
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+8)) = int32(_a_F_CheckPointGuts_44)
	v1269 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14]))
	F_sort_checkpoint_bufferids(m, v1269, v1247)
	mBase = m.M
	if v1263 < v1247 {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	F_binaryheap_build(m, v1400)
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L1
	} else {
		goto L346
	}
L316:
	;
	v1273 = int32(0)
	v1277 = v1263
	v1278 = v1104
	v1282 = v1273
	v1283 = v1273
	goto L319
L317:
	;
	goto L318
L318:
	;
	v1387 = int32(0)
	v1391 = F_binaryheap_allocate(m, v1387, int32(1077), v1387)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L345
	}
L319:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14]))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1294+v1282*int32(20))))
	if v1278 == v1298 {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	v1350 = int32(0)
	v1353 = F_binaryheap_allocate(m, v1334, int32(1077), v1350)
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L1
	} else {
		goto L339
	}
L321:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+24)) = v1339 + int32(1)
	v1344 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	if v1344 != 0 {
		goto L334
	} else {
		goto L335
	}
L322:
	;
	v1300 = v1278
	goto L324
L323:
	;
	v1300 = int32(0)
	goto L324
L324:
	;
	if v1300 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1304 = v1277 + int32(1)
	v1306 = v1304 * int32(40)
	if v1283 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L326:
	;
	goto L327
L327:
	;
	v1329 = int32(40)
	v1334 = v1277
	v1335 = v1278
	v1336 = v1283 + v1277*v1329 - v1329
	v1338 = v1283
	goto L321
L328:
	;
	v1316 = v1313 + v1277*int32(40)
	v1317 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1316)+32)) = v1317
	*(*int64)(unsafe.Add(mBase, uint32(v1316))) = v1317
	*(*int64)(unsafe.Add(mBase, uint32(v1316)+24)) = v1317
	*(*int64)(unsafe.Add(mBase, uint32(v1316)+16)) = v1317
	*(*int64)(unsafe.Add(mBase, uint32(v1316)+8)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v1316)+32)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v1316))) = v1298
	v1334 = v1304
	v1335 = v1298
	v1336 = v1316
	v1338 = v1313
	goto L321
L329:
	;
	v1309 = F_palloc(m, v1306)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L1
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	v1311 = F_repalloc(m, v1283, v1306)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L1
	} else {
		goto L333
	}
L332:
	;
	v1313 = v1309
	goto L328
L333:
	;
	v1313 = v1311
	goto L328
L334:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L1
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	v1348 = v1282 + int32(1)
	if v1348 != v1247 {
		v1277 = v1334
		v1278 = v1335
		v1282 = v1348
		v1283 = v1338
		goto L319
	} else {
		goto L338
	}
L337:
	;
	goto L336
L338:
	;
	goto L320
L339:
	;
	if v1334 <= int32(0) {
		v1400 = v1353
		v1401 = v1338
		goto L315
	} else {
		goto L340
	}
L340:
	;
	v1363 = v1350
	goto L341
L341:
	;
	v1377 = v1338 + v1363*int32(40)
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v1377)+16)) = base.F64_div(base.F64_convert_i32_u(v1247), base.F64_convert_i32_s(v1378))
	F_binaryheap_add_unordered(m, v1353, v1377)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L1
	} else {
		goto L343
	}
L342:
	;
	v1400 = v1353
	v1401 = v1338
	goto L315
L343:
	;
	v1385 = v1363 + int32(1)
	if v1385 != v1334 {
		v1363 = v1385
		goto L341
	} else {
		goto L344
	}
L344:
	;
	goto L342
L345:
	;
	v1400 = v1391
	v1401 = v1387
	goto L315
L346:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1400)))
	if v1412 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	F_IssuePendingWritebacks(m, v1109+int32(8), int32(3))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L1
	} else {
		goto L410
	}
L348:
	;
	v1660 = int32(0)
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1417 = int32(0)
	v1421 = v1417
	v1423 = v1417
	goto L351
L351:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[12]))
	v1439 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14]))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+20))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+32))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1439+v1441*int32(20))+16))
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1437+v1445<<(uint(int32(6))%32))+24))
	if v1449&int32(1073741824) == int32(0) {
		v1471 = v1421
		goto L353
	} else {
		goto L354
	}
L352:
	;
	v1660 = v1471
	goto L347
L353:
	;
	v1472 = *(*float64)(unsafe.Add(mBase, uint32(v1440)+16))
	v1473 = *(*float64)(unsafe.Add(mBase, uint32(v1440)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1440)+8)) = base.F64_add(v1472, v1473)
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+28))
	v1477 = int32(1)
	v1478 = v1476 + v1477
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+28)) = v1478
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+32)) = v1480 + v1477
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+24))
	if v1484 == v1478 {
		goto L358
	} else {
		goto L359
	}
L354:
	;
	v1457 = F_SyncOneBuffer(m, v1445, int32(0), v1109+int32(8))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	if v1457&int32(1) == int32(0) {
		v1471 = v1421
		goto L353
	} else {
		goto L356
	}
L356:
	;
	v1463 = int32(_a_F_CheckPointGuts_45)
	v1465 = *(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[16])) = v1465 + int64(1)
	v1471 = v1421 + int32(1)
	goto L353
L357:
	;
	v1491 = v1423 + int32(1)
	v1495 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[17]))
	if v1495 != int32(11) {
		goto L363
	} else {
		goto L364
	}
L358:
	;
	v1486 = F_binaryheap_remove_first(m, v1400)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L1
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	F_binaryheap_replace_first(m, v1400, v1440)
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L1
	} else {
		goto L362
	}
L361:
	;
	goto L357
L362:
	;
	goto L357
L363:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1400)))
	if v1657 != 0 {
		v1421 = v1471
		v1423 = v1491
		goto L351
	} else {
		goto L409
	}
L364:
	;
	if l1&int32(4) != 0 {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	if v1649 == int32(0) {
		goto L363
	} else {
		goto L407
	}
L366:
	;
	v1632 = int32(_a_F_CheckPointGuts_46)
	v1634 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[18]))
	v1636 = v1634 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[18])) = v1636
	if int32(0) < v1636 {
		goto L365
	} else {
		goto L405
	}
L367:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[19]))
	if v1501 != 0 {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[20]))
	if v1503 != 0 {
		goto L366
	} else {
		goto L369
	}
L369:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[21]))
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+20))
	if v1506&int32(4) != 0 {
		goto L366
	} else {
		goto L370
	}
L370:
	;
	v1509 = m.G0
	v1511 = v1509 - int32(16)
	m.G0 = v1511
	v1515 = *(*float64)(unsafe.Add(mBase, _c_F_CheckPointGuts[22]))
	v1516 = base.F64_mul(base.F64_div(base.F64_convert_i32_s(v1491), base.F64_convert_i32_s(v1247)), v1515)
	v1518 = *(*float64)(unsafe.Add(mBase, _c_F_CheckPointGuts[23]))
	if base.F64_lt(v1516, v1518) != 0 {
		v1578 = int32(0)
		goto L371
	} else {
		goto L372
	}
L371:
	;
	m.G0 = v1511 + int32(16)
	if v1578 == int32(0) {
		goto L366
	} else {
		goto L387
	}
L372:
	;
	v1522 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[24])))
	if v1522 == int32(1) {
		goto L375
	} else {
		goto L376
	}
L373:
	;
	v1540 = *(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[25]))
	v1544 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[26]))
	v1548 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[27]))
	v1550 = base.F64_div(base.F64_div(base.F64_convert_i64_u(v1538-v1540), base.F64_convert_i32_s(v1544)), base.F64_convert_i32_s(v1548))
	if base.F64_lt(v1516, v1550) == int32(0) {
		goto L383
	} else {
		goto L384
	}
L374:
	;
	if v1532 != 0 {
		goto L378
	} else {
		goto L379
	}
L375:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[28]))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+316))
	v1530 = base.B2i32(v1528 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[24])) = uint8(v1530)
	v1532 = v1530
	goto L377
L376:
	;
	v1532 = int32(0)
	goto L377
L377:
	;
	goto L374
L378:
	;
	v1534 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L1
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v1536 = F_GetInsertRecPtr(m)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L1
	} else {
		goto L382
	}
L381:
	;
	v1538 = v1534
	goto L373
L382:
	;
	v1538 = v1536
	goto L373
L383:
	;
	F_gettimeofday(m, v1511)
	mBase = m.M
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+8))
	v1560 = *(*int64)(unsafe.Add(mBase, uint32(v1511)))
	v1562 = *(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[29]))
	v1567 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[30]))
	v1569 = base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v1556), float64(1e+06)), base.F64_convert_i64_s(v1560-v1562)), base.F64_convert_i32_s(v1567))
	if base.F64_lt(v1516, v1569) == int32(0) {
		v1578 = int32(1)
		goto L371
	} else {
		goto L386
	}
L384:
	;
	v1573 = v1550
	goto L385
L385:
	;
	*(*float64)(unsafe.Add(mBase, _c_F_CheckPointGuts[23])) = v1573
	v1578 = int32(0)
	goto L371
L386:
	;
	v1573 = v1569
	goto L385
L387:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[31]))
	if v1585 != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[31])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L1
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L1
	} else {
		goto L400
	}
L391:
	;
	F_SyncRepUpdateSyncStandbysDefined(m)
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	v1598 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	if v1598 != 0 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_47), int32(0))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L1
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	goto L390
L398:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_48), int32(1390), int32(_a_F_CheckPointGuts_49))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	goto L397
L400:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[18])) = int32(1000)
	F_CheckArchiveTimeout(m)
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[32]))
	v1623 = F_WaitLatch(m, v1619, int32(41), int32(100), int32(150994945))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v1626))) = int32(0)
	goto L404
L404:
	;
	goto L365
L405:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[18])) = int32(1000)
	goto L365
L407:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	goto L363
L409:
	;
	goto L352
L410:
	;
	F_pfree(m, v1401)
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	F_pfree(m, v1400)
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	v1684 = int32(_a_F_CheckPointGuts_50)
	v1686 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[33]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[33])) = v1686 + v1660
	goto L281
L413:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[34])) = v1719 + v1718*int64(1000000) - int64(946684800000000)
	v1729 = int64(0)
	v1730 = int32(0)
	v1732 = m.G0
	v1734 = v1732 - int32(1152)
	m.G0 = v1734
	v1737 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[35]))
	if v1737 != 0 {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	v2196 = m.G0
	v2197 = int32(16)
	v2198 = v2196 - v2197
	m.G0 = v2198
	F_gettimeofday(m, v2198)
	mBase = m.M
	v2201 = *(*int64)(unsafe.Add(mBase, uint32(v2198)))
	v2202 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2198)+8)))
	m.G0 = v2198 + v2197
	goto L516
L415:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L1
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L1
	} else {
		goto L513
	}
L418:
	;
	v1741 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[36])))
	if v1741 == int32(0) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1796 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[36])) = uint8(v1796)
	v1798 = int32(_a_F_CheckPointGuts_51)
	v1800 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])))
	v1802 = v1800 + v1796
	*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])) = uint16(v1802)
	v1805 = v1734 + int32(1116)
	v1807 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[35]))
	F_hash_seq_init(m, v1805, v1807)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L1
	} else {
		goto L428
	}
L420:
	;
	v1745 = v1734 + int32(1116)
	v1747 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[35]))
	F_hash_seq_init(m, v1745, v1747)
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	v1750 = F_hash_seq_search(m, v1745)
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L1
	} else {
		goto L422
	}
L422:
	;
	if v1750 == int32(0) {
		goto L419
	} else {
		goto L423
	}
L423:
	;
	v1756 = v1750
	goto L424
L424:
	;
	v1772 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1756)+24)) = uint16(v1772)
	v1776 = F_hash_seq_search(m, v1734+int32(1116))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L1
	} else {
		goto L426
	}
L425:
	;
	goto L419
L426:
	;
	if v1776 != 0 {
		v1756 = v1776
		goto L424
	} else {
		goto L427
	}
L427:
	;
	goto L425
L428:
	;
	v1810 = F_hash_seq_search(m, v1805)
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L1
	} else {
		goto L430
	}
L429:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L1
	} else {
		goto L510
	}
L430:
	;
	if v1810 != 0 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1815 = v1810
	v1817 = int32(10)
	v1818 = v1730
	v1824 = v1729
	v1825 = v1729
	goto L434
L432:
	;
	v2142 = v1730
	v2148 = v1729
	v2149 = v1729
	goto L433
L433:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[38])) = v2148
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[39])) = v2149
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[40])) = v2142
	v2161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[36])) = uint8(v2161)
	m.G0 = v1734 + int32(1152)
	goto L414
L434:
	;
	v1830 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1815)+24)))
	v1832 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])))
	if v1830 != v1832 {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	v2142 = v2121
	v2148 = v2127
	v2149 = v2128
	goto L433
L436:
	;
	v1835 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[5])))
	if v1835 != int32(1) {
		v2095 = v1817
		v2096 = v1818
		v2102 = v1824
		v2103 = v1825
		goto L439
	} else {
		goto L440
	}
L437:
	;
	v2120 = v1817
	v2121 = v1818
	v2127 = v1824
	v2128 = v1825
	goto L438
L438:
	;
	v2135 = F_hash_seq_search(m, v1734+int32(1116))
	mBase = m.M
	v2136 = m.ExcPending
	if v2136 != 0 {
		goto L1
	} else {
		goto L508
	}
L439:
	;
	v2109 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[35]))
	v2112 = F_hash_search(m, v2109, v1815, int32(2), int32(0))
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L1
	} else {
		goto L506
	}
L440:
	;
	v1839 = v1817 - int32(1)
	if v1839 <= int32(0) {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L1
	} else {
		goto L444
	}
L442:
	;
	v1845 = v1839
	goto L443
L443:
	;
	v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815)+26)))
	if v1846 != 0 {
		v2095 = v1845
		v2096 = v1818
		v2102 = v1824
		v2103 = v1825
		goto L439
	} else {
		goto L445
	}
L444:
	;
	v1845 = int32(10)
	goto L443
L445:
	;
	F___clock_gettime(m, int32(1), v1734+int32(1136))
	mBase = m.M
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+1144))
	v1852 = *(*int64)(unsafe.Add(mBase, uint32(v1734)+1136))
	v1854 = v1734 + int32(80)
	v1855 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1815))))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1855*int32(12))+uint32(_c_F_CheckPointGuts[41])))
	v1861 = m.T0[v1860].(func(*base.Module, int32, int32) int32)(m, v1815, v1854)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L447
	}
L446:
	;
	v2046 = int32(1)
	F___clock_gettime(m, v2046, v1734+int32(1136))
	mBase = m.M
	v2050 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1734)+1144)))
	v2053 = *(*int64)(unsafe.Add(mBase, uint32(v1734)+1136))
	v2059 = base.I64_div_s(v2050-base.I64_extend_i32_s(v2032)+(v2053-v2042)*int64(1000000000), int64(1000))
	v2062 = v1818 + v2046
	v2064 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[42])))
	if v2064 != v2046 {
		goto L497
	} else {
		goto L498
	}
L447:
	;
	if v1861 == int32(0) {
		v2032 = v1851
		v2033 = v1845
		v2042 = v1852
		goto L446
	} else {
		goto L448
	}
L448:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v1867 == int32(44) {
		goto L451
	} else {
		goto L452
	}
L449:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L1
	} else {
		goto L467
	}
L450:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), v1906, int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L466
	}
L451:
	;
	v1872 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	v1888 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])))
	if v1888 != 0 {
		goto L459
	} else {
		goto L460
	}
L454:
	;
	if v1872 == int32(0) {
		goto L449
	} else {
		goto L455
	}
L455:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1734)+48)) = v1854
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_54), v1734+int32(48))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	v1906 = int32(452)
	goto L450
L458:
	;
	v1891 = F_errstart(m, v1889, int32(0))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L1
	} else {
		goto L462
	}
L459:
	;
	v1889 = int32(21)
	goto L461
L460:
	;
	v1889 = int32(23)
	goto L461
L461:
	;
	goto L458
L462:
	;
	if v1891 == int32(0) {
		goto L449
	} else {
		goto L463
	}
L463:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1734)+64)) = v1734 + int32(80)
	F_errmsg(m, int32(_a_F_CheckPointGuts_20), v1734-int32(-64))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	v1906 = int32(447)
	goto L450
L466:
	;
	goto L449
L467:
	;
	v1913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815)+26)))
	if v1913 == int32(0) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v1920 = int32(1)
	goto L471
L469:
	;
	goto L470
L470:
	;
	v2095 = int32(10)
	v2096 = v1818
	v2102 = v1824
	v2103 = v1825
	goto L439
L471:
	;
	F___clock_gettime(m, int32(1), v1734+int32(1136))
	mBase = m.M
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+1144))
	v1938 = *(*int64)(unsafe.Add(mBase, uint32(v1734)+1136))
	v1941 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1815))))
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1941*int32(12))+uint32(_c_F_CheckPointGuts[41])))
	v1947 = m.T0[v1946].(func(*base.Module, int32, int32) int32)(m, v1815, v1734+int32(80))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L1
	} else {
		goto L473
	}
L472:
	;
	goto L470
L473:
	;
	if v1947 == int32(0) {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v2032 = v1937
	v2033 = int32(10)
	v2042 = v1938
	goto L446
L475:
	;
	goto L476
L476:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	v1957 = int32(0)
	if base.B2i32(v1954 == int32(44))&base.B2i32(v1920 <= v1957) == v1957 {
		goto L479
	} else {
		goto L480
	}
L477:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L1
	} else {
		goto L495
	}
L478:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), v2000, int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L1
	} else {
		goto L494
	}
L479:
	;
	v1965 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])))
	if v1965 != 0 {
		goto L483
	} else {
		goto L484
	}
L480:
	;
	goto L481
L481:
	;
	v1985 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L1
	} else {
		goto L490
	}
L482:
	;
	v1968 = F_errstart(m, v1966, int32(0))
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L1
	} else {
		goto L486
	}
L483:
	;
	v1966 = int32(21)
	goto L485
L484:
	;
	v1966 = int32(23)
	goto L485
L485:
	;
	goto L482
L486:
	;
	if v1968 == int32(0) {
		goto L477
	} else {
		goto L487
	}
L487:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1734)+16)) = v1734 + int32(80)
	F_errmsg(m, int32(_a_F_CheckPointGuts_20), v1734+int32(16))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	v2000 = int32(447)
	goto L478
L490:
	;
	if v1985 == int32(0) {
		goto L477
	} else {
		goto L491
	}
L491:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1734)+32)) = v1734 + int32(80)
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_54), v1734+int32(32))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	v2000 = int32(452)
	goto L478
L494:
	;
	goto L477
L495:
	;
	v2008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815)+26)))
	if v2008 != int32(1) {
		v1920 = v1920 + int32(1)
		goto L471
	} else {
		goto L496
	}
L496:
	;
	goto L472
L497:
	;
	if base.Ui64(v1825) < base.Ui64(v2059) {
		goto L503
	} else {
		goto L504
	}
L498:
	;
	v2069 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	if v2069 == int32(0) {
		goto L497
	} else {
		goto L500
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1734))) = v2062
	*(*float64)(unsafe.Add(mBase, uint32(v1734)+8)) = base.F64_div(base.F64_convert_i64_u(v2059), float64(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v1734)+4)) = v1734 + int32(80)
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_55), v1734)
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), int32(432), int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	goto L497
L503:
	;
	v2089 = v2059
	goto L505
L504:
	;
	v2089 = v1825
	goto L505
L505:
	;
	v2095 = v2033
	v2096 = v2062
	v2102 = v1824 + v2059
	v2103 = v2089
	goto L439
L506:
	;
	if v2112 == int32(0) {
		goto L429
	} else {
		goto L507
	}
L507:
	;
	v2120 = v2095
	v2121 = v2096
	v2127 = v2102
	v2128 = v2103
	goto L438
L508:
	;
	if v2135 != 0 {
		v1815 = v2135
		v1817 = v2120
		v1818 = v2121
		v1824 = v2127
		v1825 = v2128
		goto L434
	} else {
		goto L509
	}
L509:
	;
	goto L435
L510:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_56), int32(0))
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), int32(465), int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L513:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_57), int32(0))
	mBase = m.M
	v2186 = m.ExcPending
	if v2186 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), int32(308), int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v2191 = m.ExcPending
	if v2191 != 0 {
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
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[43])) = v2202 + v2201*int64(1000000) - int64(946684800000000)
	v2212 = int32(0)
	v2214 = m.G0
	v2216 = v2214 - int32(1152)
	m.G0 = v2216
	v2219 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[44]))
	if v2219 <= v2212 {
		goto L521
	} else {
		goto L522
	}
L517:
	;
	return
L518:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L1
	} else {
		goto L594
	}
L519:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L1
	} else {
		goto L590
	}
L520:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L1
	} else {
		goto L586
	}
L521:
	;
	m.G0 = v2216 + int32(1152)
	goto L517
L522:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v2227 = F_LWLockAcquire(m, v2223+int32(2304), int32(1))
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[45]))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2230)+4))
	if int32(0) < v2231 {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	v2236 = v2230
	v2237 = v2212
	v2239 = v2212
	goto L527
L525:
	;
	v2441 = v2212
	goto L526
L526:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v2456+int32(2304))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L1
	} else {
		goto L579
	}
L527:
	;
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(v2236+v2239<<(uint(int32(2))%32))+8))
	v2255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2254)+44)))
	if v2255 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L528:
	;
	v2441 = v2429
	goto L526
L529:
	;
	v2435 = v2239 + int32(1)
	v2436 = *(*int32)(unsafe.Add(mBase, uint32(v2428)+4))
	if v2435 < v2436 {
		v2236 = v2428
		v2237 = v2429
		v2239 = v2435
		goto L527
	} else {
		goto L578
	}
L530:
	;
	v2258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2254)+46)))
	if v2258 != int32(1) {
		v2428 = v2236
		v2429 = v2237
		goto L529
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	v2261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2254)+45)))
	if v2261 != 0 {
		v2428 = v2236
		v2429 = v2237
		goto L529
	} else {
		goto L534
	}
L533:
	;
	goto L532
L534:
	;
	v2262 = *(*int64)(unsafe.Add(mBase, uint32(v2254)+24))
	if base.Ui64(l0) < base.Ui64(v2262) {
		v2428 = v2236
		v2429 = v2237
		goto L529
	} else {
		goto L535
	}
L535:
	;
	v2264 = *(*int64)(unsafe.Add(mBase, uint32(v2254)+16))
	F_XlogReadTwoPhaseData(m, v2264, v2216+int32(120), v2216+int32(116))
	mBase = m.M
	v2270 = m.ExcPending
	if v2270 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+32))
	v2272 = int32(-1)
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2216)+120))
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v2216)+116))
	v2275 = m.Env.Pgmem_crc32c(m, v2272, v2273, v2274)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2216)+124)) = v2275 ^ v2272
	v2279 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L1
	} else {
		goto L537
	}
L537:
	;
	if base.Ui32(v2271) <= base.Ui32(int32(2)) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v2298 = base.I64_extend_i32_u(v2271)
	goto L540
L539:
	;
	v2286 = int64(base.Ui64(v2279) >> (uint(int64(32)) % 64))
	if base.Ui32(base.I32_wrap_i64(v2279)) < base.Ui32(v2271) {
		goto L541
	} else {
		goto L542
	}
L540:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v2216)+100)) = uint32(v2298)
	v2301 = int64(base.Ui64(v2298) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2216)+96)) = uint32(v2301)
	v2304 = v2216 + int32(128)
	v2309 = F_pg_snprintf(m, v2304, int32(1024), int32(_a_F_CheckPointGuts_58), v2216+int32(96))
	mBase = m.M
	v2310 = m.ExcPending
	if v2310 != 0 {
		goto L1
	} else {
		goto L544
	}
L541:
	;
	v2293 = (v2286 - int64(1)) & int64(4294967295)
	goto L543
L542:
	;
	v2293 = v2286
	goto L543
L543:
	;
	v2298 = base.I64_extend_i32_u(v2271) | v2293<<(uint(int64(32))%64)
	goto L540
L544:
	;
	v2312 = F_OpenTransientFile(m, v2304, int32(577))
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L1
	} else {
		goto L545
	}
L545:
	;
	if v2312 < int32(0) {
		goto L520
	} else {
		goto L546
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(0)
	v2320 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v2320))) = int32(167772224)
	v2323 = F_write(m, v2312, v2273, v2274)
	mBase = m.M
	if v2323 != v2274 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v2326 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v2326 == int32(0) {
		goto L550
	} else {
		goto L551
	}
L548:
	;
	goto L549
L549:
	;
	v2353 = int32(4)
	v2354 = F_write(m, v2312, v2216+int32(124), v2353)
	mBase = m.M
	if v2354 != v2353 {
		goto L557
	} else {
		goto L558
	}
L550:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L552
L551:
	;
	goto L552
L552:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L1
	} else {
		goto L554
	}
L554:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2216)+80)) = v2216 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_59), v2216+int32(80))
	mBase = m.M
	v2345 = m.ExcPending
	if v2345 != 0 {
		goto L1
	} else {
		goto L555
	}
L555:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1756), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L557:
	;
	v2358 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v2358 == int32(0) {
		goto L560
	} else {
		goto L561
	}
L558:
	;
	goto L559
L559:
	;
	v2383 = int32(_a_F_CheckPointGuts_62)
	v2384 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	v2385 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2384))) = v2385
	v2388 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v2388))) = int32(167772223)
	v2393 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[5])))
	if v2393 != int32(1) {
		v2407 = v2385
		goto L568
	} else {
		goto L569
	}
L560:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L562
L561:
	;
	goto L562
L562:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2367 = m.ExcPending
	if v2367 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2216)+64)) = v2216 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_59), v2216-int32(-64))
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1765), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L567:
	;
	if v2407 != 0 {
		goto L519
	} else {
		goto L574
	}
L568:
	;
	goto L567
L569:
	;
	goto L570
L570:
	;
	v2398 = F_fsync(m, v2312)
	mBase = m.M
	if v2398 != int32(-1) {
		v2407 = v2398
		goto L568
	} else {
		goto L572
	}
L571:
	;
	v2407 = int32(-1)
	goto L568
L572:
	;
	v2402 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v2402 == int32(27) {
		goto L570
	} else {
		goto L573
	}
L573:
	;
	goto L571
L574:
	;
	v2409 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v2409))) = int32(0)
	v2412 = F_CloseTransientFile(m, v2312)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	if v2412 != 0 {
		goto L518
	} else {
		goto L576
	}
L576:
	;
	v2414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2254)+45)) = uint8(v2414)
	v2417 = v2254 + int32(16)
	v2418 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2417)+8)) = v2418
	*(*int64)(unsafe.Add(mBase, uint32(v2417))) = v2418
	F_pfree(m, v2273)
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[45]))
	v2428 = v2427
	v2429 = v2237 + int32(1)
	goto L529
L578:
	;
	goto L528
L579:
	;
	F_fsync_fname(m, int32(_a_F_CheckPointGuts_63), int32(1))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	v2466 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[42])))
	if base.B2i32(v2466 != int32(1))|base.B2i32(v2441 <= int32(0)) != 0 {
		goto L521
	} else {
		goto L581
	}
L581:
	;
	v2474 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	if v2474 == int32(0) {
		goto L521
	} else {
		goto L583
	}
L583:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2216))) = v2441
	F_errmsg_plural(m, int32(_a_F_CheckPointGuts_64), int32(_a_F_CheckPointGuts_65), v2441, v2216)
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L1
	} else {
		goto L584
	}
L584:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1876), int32(_a_F_CheckPointGuts_66))
	mBase = m.M
	v2487 = m.ExcPending
	if v2487 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	goto L521
L586:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2216)+16)) = v2216 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_67), v2216+int32(16))
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1744), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L590:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2216)+48)) = v2216 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_20), v2216+int32(48))
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1777), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2545 = m.ExcPending
	if v2545 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L594:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2551 = m.ExcPending
	if v2551 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2216)+32)) = v2216 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_23), v2216+int32(32))
	mBase = m.M
	v2559 = m.ExcPending
	if v2559 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1783), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2564 = m.ExcPending
	if v2564 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_point_dt(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v15 float64
	_ = v15
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v31 float64
	_ = v31
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v64 float64
	_ = v64
	var v71 float64
	_ = v71
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v11 = base.F64_abs(base.F64_sub(v8, v9))
	if base.F64_eq(v11, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v15 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v8), v15)&base.F64_ne(base.F64_abs(v9), v15) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return float64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v21 = math.Float64frombits(uint64(0x7ff0000000000000))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_eq(base.F64_abs(base.F64_sub(v22, v23)), v21) != 0 {
				v36 = v22
				v37 = v23
				v38 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(base.F64_abs(v36), v38)|base.F64_eq(base.F64_abs(v37), v38) != 0 {
					v71 = v38
					return v71
				} else {
					F_float_overflow_error(m)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				v71 = v21
				return v71
			}
		}
	} else {
		v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
		v31 = base.F64_abs(base.F64_sub(v28, v29))
		if base.F64_ne(v31, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v48 = int64(9218868437227405312)
			if base.B2i32(base.Ui64(v48) < base.Ui64(base.I64_reinterpret_f64(v11)))|base.B2i32(base.Ui64(v48) < base.Ui64(base.I64_reinterpret_f64(v31))) != 0 {
				v71 = math.Float64frombits(uint64(0x7ff8000000000000))
				return v71
			} else {
				v54 = base.F64_gt(v31, v11)
				if v54 != 0 {
					v55 = v31
				} else {
					v55 = v11
				}
				if v54 != 0 {
					v56 = v11
				} else {
					v56 = v31
				}
				if base.F64_eq(v56, float64(0)) != 0 {
					v71 = v55
					return v71
				} else {
					v59 = base.F64_div(v56, v55)
					v64 = base.F64_mul(v55, base.F64_sqrt(base.F64_add(base.F64_mul(v59, v59), float64(1))))
					if base.F64_eq(base.F64_abs(v64), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if base.F64_eq(v64, float64(0)) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v71 = v64
							return v71
						}
					}
				}
			}
		} else {
			v36 = v28
			v37 = v29
			v38 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v36), v38)|base.F64_eq(base.F64_abs(v37), v38) != 0 {
				v71 = v38
				return v71
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_point_inside(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v20 float64
	_ = v20
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v48 int32
	_ = v48
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v147 int32
	_ = v147
	var v166 int32
	_ = v166
	v12 = int32(0)
	v16 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v18 = base.F64_sub(v16, v17)
	v20 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v18), v20)|base.F64_eq(base.F64_abs(v16), v20) == v12)&base.F64_ne(base.F64_abs(v17), v20) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L19
	}
L2:
	;
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = base.F64_sub(v32, v33)
	v36 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v34), v36)|base.F64_eq(base.F64_abs(v32), v36) == int32(0))&base.F64_ne(base.F64_abs(v33), v36) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v48 = int32(2)
	if l1 < v48 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v147
L5:
	;
	v127 = F_lseg_crossing(m, v18, v34, v118, v117)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L13
	} else {
		goto L17
	}
L6:
	;
	v117 = v34
	v118 = v18
	v126 = v12
	goto L5
L7:
	;
	goto L8
L8:
	;
	v55 = v18
	v56 = v34
	v63 = int32(1)
	v66 = v12
	goto L9
L9:
	;
	v69 = l2 + v63<<(uint(int32(4))%32)
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v72 = base.F64_sub(v70, v71)
	v74 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v72), v74)|base.F64_eq(base.F64_abs(v70), v74) == int32(0))&base.F64_ne(base.F64_abs(v71), v74) != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v117 = v88
	v118 = v72
	v126 = v108
	goto L5
L11:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(v69)+8))
	v87 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v88 = base.F64_sub(v86, v87)
	v90 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v88), v90)|base.F64_eq(base.F64_abs(v86), v90) == int32(0))&base.F64_ne(base.F64_abs(v87), v90) != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v102 = F_lseg_crossing(m, v72, v88, v55, v56)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v102 == int32(2147483647) {
		v147 = v48
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v108 = v102 + v66
	v110 = v63 + int32(1)
	if v110 != l1 {
		v55 = v72
		v56 = v88
		v63 = v110
		v66 = v108
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	if v127 == int32(2147483647) {
		v147 = v48
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v147 = base.B2i32(v126 != int32(0)-v127)
	goto L4
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_point_invsl(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v58 float64
	_ = v58
	var v68 float64
	_ = v68
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v3 = float64(0)
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_eq(v10, v11) != 0 {
		v68 = v3
		return v68
	} else {
		v13 = base.F64_sub(v10, v11)
		v14 = base.F64_abs(v13)
		if base.F64_le(v14, float64(1e-06)) != 0 {
			v68 = v3
			return v68
		} else {
			v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_eq(v18, v19)|base.F64_le(base.F64_abs(base.F64_sub(v18, v19)), float64(1e-06)) != 0 {
				v68 = math.Float64frombits(uint64(0x7ff0000000000000))
				return v68
			} else {
				v26 = math.Float64frombits(uint64(0x7ff0000000000000))
				v27 = base.F64_ne(v14, v26)
				if base.B2i32(v27|base.F64_eq(base.F64_abs(v10), v26) == int32(0))&base.F64_ne(base.F64_abs(v11), v26) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v39 = math.Float64frombits(uint64(0x7ff0000000000000))
					v41 = base.F64_sub(v19, v18)
					v42 = base.F64_abs(v41)
					if base.B2i32(base.F64_eq(base.F64_abs(v18), v39)|base.F64_ne(v42, v39) == int32(0))&base.F64_ne(base.F64_abs(v19), v39) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v41, float64(0)) != 0 {
							F_float_zero_divide_error(m)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v58 = base.F64_div(v13, v41)
							if v27&base.F64_eq(base.F64_abs(v58), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.F64_ne(v58, float64(0)) != 0 {
									v68 = v58
									return v68
								} else {
									if base.F64_ne(v42, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v68 = v58
										return v68
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_point_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = v8 + int32(16)
	F_initStringInfo(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_appendStringInfoChar(m, v12, int32(40))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
			v22 = F_float8out_internal(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_float8out_internal(m, v20)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
					F_appendStringInfo(m, v12, int32(_a_F_point_out_0), v8)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v22)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v24)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_appendStringInfoChar(m, v12, int32(41))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
									m.G0 = v8 + int32(32)
									return v38
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_point_sub(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v54 int32
	_ = v54
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_palloc(m, int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_sub(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.B2i32(base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19) == int32(0))&base.F64_ne(base.F64_abs(v16), v19) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v33 = base.F64_sub(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.B2i32(base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35) == int32(0))&base.F64_ne(base.F64_abs(v32), v35) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v33
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				return v11
			}
		}
	}
}
