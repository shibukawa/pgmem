package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CheckPointGuts(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int64
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
	var v198 int64
	_ = v198
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v290 int32
	_ = v290
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int64
	_ = v371
	var v372 int32
	_ = v372
	var v373 int64
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int64
	_ = v382
	var v395 int32
	_ = v395
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int64
	_ = v510
	var v511 int64
	_ = v511
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int64
	_ = v794
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int64
	_ = v805
	var v807 int64
	_ = v807
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v864 int32
	_ = v864
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int64
	_ = v1010
	var v1011 int64
	_ = v1011
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int64
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1062 int32
	_ = v1062
	var v1070 int64
	_ = v1070
	var v1074 int64
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1106 int64
	_ = v1106
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int64
	_ = v1130
	var v1132 int64
	_ = v1132
	var v1148 int64
	_ = v1148
	var v1163 int64
	_ = v1163
	var v1190 int32
	_ = v1190
	var v1191 int64
	_ = v1191
	var v1194 int64
	_ = v1194
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1238 int32
	_ = v1238
	var v1240 int64
	_ = v1240
	var v1242 int64
	_ = v1242
	var v1258 int64
	_ = v1258
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1287 int32
	_ = v1287
	var v1288 int64
	_ = v1288
	var v1290 int64
	_ = v1290
	var v1295 int64
	_ = v1295
	var v1308 int64
	_ = v1308
	var v1319 int64
	_ = v1319
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1404 int64
	_ = v1404
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1453 int32
	_ = v1453
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1518 int32
	_ = v1518
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1531 int64
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1541 int64
	_ = v1541
	var v1544 int64
	_ = v1544
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1558 int32
	_ = v1558
	var v1560 int64
	_ = v1560
	var v1566 int32
	_ = v1566
	var v1567 float64
	_ = v1567
	var v1568 float64
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1581 int64
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1612 float64
	_ = v1612
	var v1613 float64
	_ = v1613
	var v1615 float64
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1631 int64
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int64
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int64
	_ = v1635
	var v1637 int64
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1647 float64
	_ = v1647
	var v1653 int32
	_ = v1653
	var v1657 int64
	_ = v1657
	var v1659 int64
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1666 float64
	_ = v1666
	var v1670 float64
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1682 int32
	_ = v1682
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1742 int32
	_ = v1742
	var v1750 int32
	_ = v1750
	var v1754 int32
	_ = v1754
	var v1758 int32
	_ = v1758
	var v1761 int32
	_ = v1761
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int64
	_ = v1821
	var v1822 int64
	_ = v1822
	var v1832 int64
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1859 int32
	_ = v1859
	var v1876 int32
	_ = v1876
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1929 int64
	_ = v1929
	var v1930 int64
	_ = v1930
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1957 int32
	_ = v1957
	var v1958 int64
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1973 int32
	_ = v1973
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1983 int32
	_ = v1983
	var v1989 int32
	_ = v1989
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
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2026 int32
	_ = v2026
	var v2044 int32
	_ = v2044
	var v2045 int64
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2080 int32
	_ = v2080
	var v2088 int32
	_ = v2088
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2150 int64
	_ = v2150
	var v2155 int32
	_ = v2155
	var v2159 int64
	_ = v2159
	var v2162 int64
	_ = v2162
	var v2168 int64
	_ = v2168
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2192 int32
	_ = v2192
	var v2197 int32
	_ = v2197
	var v2198 int64
	_ = v2198
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2211 int64
	_ = v2211
	var v2212 int64
	_ = v2212
	var v2219 int32
	_ = v2219
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2230 int32
	_ = v2230
	var v2233 int32
	_ = v2233
	var v2237 int64
	_ = v2237
	var v2238 int64
	_ = v2238
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2255 int32
	_ = v2255
	var v2259 int64
	_ = v2259
	var v2260 int64
	_ = v2260
	var v2273 int32
	_ = v2273
	var v2281 int32
	_ = v2281
	var v2285 int32
	_ = v2285
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2298 int32
	_ = v2298
	var v2303 int32
	_ = v2303
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2313 int64
	_ = v2313
	var v2314 int64
	_ = v2314
	var v2324 int32
	_ = v2324
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2331 int32
	_ = v2331
	var v2335 int32
	_ = v2335
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2374 int32
	_ = v2374
	var v2375 int64
	_ = v2375
	var v2377 int64
	_ = v2377
	var v2383 int32
	_ = v2383
	var v2384 int64
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2390 int64
	_ = v2390
	var v2397 int32
	_ = v2397
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2413 int32
	_ = v2413
	var v2416 int32
	_ = v2416
	var v2419 int32
	_ = v2419
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2438 int32
	_ = v2438
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2451 int32
	_ = v2451
	var v2460 int32
	_ = v2460
	var v2462 int32
	_ = v2462
	var v2470 int32
	_ = v2470
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2481 int32
	_ = v2481
	var v2486 int32
	_ = v2486
	var v2491 int32
	_ = v2491
	var v2495 int32
	_ = v2495
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2511 int64
	_ = v2511
	var v2516 int32
	_ = v2516
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2533 int32
	_ = v2533
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2557 int32
	_ = v2557
	var v2559 int32
	_ = v2559
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2575 int32
	_ = v2575
	var v2580 int32
	_ = v2580
	var v2605 int32
	_ = v2605
	var v2607 int32
	_ = v2607
	var v2615 int32
	_ = v2615
	var v2620 int32
	_ = v2620
	var v2624 int32
	_ = v2624
	var v2626 int32
	_ = v2626
	var v2634 int32
	_ = v2634
	var v2639 int32
	_ = v2639
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2653 int32
	_ = v2653
	var v2658 int32
	_ = v2658
	v3 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v24 = F_LWLockAcquire(m, v20+int32(3200), int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v27+int32(3200))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = m.G0
	v34 = v32 - int32(1040)
	m.G0 = v34
	v38 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v38 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_0), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v49 = int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v56 = F_LWLockAcquire(m, v52+int32(_a_F_CheckPointGuts_1), v49)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_2), int32(2327), int32(_a_F_CheckPointGuts_3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[1]))
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[2]))
	if int32(0) < v59+v61 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	m.G0 = v34 + int32(1040)
	v183 = m.G0
	v185 = v183 - int32(1152)
	m.G0 = v185
	v187 = F_GetRedoRecPtr(m)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L37
	}
L12:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[3]))
	v69 = v59
	v70 = v66
	v71 = v61
	v73 = v3
	v74 = v3
	goto L15
L13:
	;
	goto L14
L14:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v157+int32(_a_F_CheckPointGuts_1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L36
	}
L15:
	;
	v87 = v70 + v73*int32(296)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v88 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v145+int32(_a_F_CheckPointGuts_1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L33
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(_a_F_CheckPointGuts_4)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v87 + int32(24)
	v99 = F_pg_sprintf(m, v34+int32(16), int32(_a_F_CheckPointGuts_5), v34)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v136 = v69
	v137 = v70
	v138 = v71
	v139 = v74
	goto L19
L19:
	;
	v141 = v73 + int32(1)
	if v141 < v136+v138 {
		v69 = v136
		v70 = v137
		v71 = v138
		v73 = v141
		v74 = v139
		goto L15
	} else {
		goto L32
	}
L20:
	;
	if l1&v49 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v87)+280))
	F_SaveSlotToPath(m, v87, v34+int32(16), int32(15))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v87)+88))
	if v103 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v108 = base.AtomicRmwXchg32(m, v87, int32(0), int32(1))
	if v108 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_s_lock(m, v87, int32(_a_F_CheckPointGuts_6))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v87)+112))
	if v112 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v118 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v87))), uint32(v118))
	goto L21
L29:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v87)+264))
	if base.Ui64(v113) <= base.Ui64(v114) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v116 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v87)+12)) = uint16(v116)
	goto L28
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[2]))
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[3]))
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[1]))
	v136 = v135
	v137 = v133
	v138 = v129
	v139 = base.B2i32(v121 != v122) | v74
	goto L19
L32:
	;
	goto L16
L33:
	;
	if v139&int32(1) == int32(0) {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
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
	v189 = F_ReplicationSlotsComputeLogicalRestartLSN(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v192 = F_AllocateDir(m, int32(_a_F_CheckPointGuts_7))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v195 = F_ReadDir(m, v192, int32(_a_F_CheckPointGuts_7))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v195 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if base.Ui64(v187) < base.Ui64(v189) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	F_FreeDir(m, v192)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L86
	}
L44:
	;
	v198 = v187
	goto L46
L45:
	;
	v198 = v189
	goto L46
L46:
	;
	v205 = v195
	goto L47
L47:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+19)))
	if v219 != int32(46) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L43
L49:
	;
	v342 = F_ReadDir(m, v192, int32(_a_F_CheckPointGuts_7))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L84
	}
L50:
	;
	v232 = v205 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+84)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v185)+80)) = int32(_a_F_CheckPointGuts_7)
	v237 = v185 + int32(96)
	v242 = F_pg_snprintf(m, v237, int32(1045), int32(_a_F_CheckPointGuts_5), v185+int32(80))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L55
	}
L51:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+20)))
	if v222 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+20)))
	if v225 != int32(46) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+21)))
	if v228 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v247 = F_get_dirent_type(m, v237, v205, int32(0), int32(14))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_8), v334, int32(_a_F_CheckPointGuts_9))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L83
	}
L57:
	;
	if v247&int32(-3) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v253 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+52)) = v185 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+48)) = v185 + int32(92)
	v273 = F_sscanf(m, v232, int32(_a_F_CheckPointGuts_10), v185+int32(48))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	if v253 == int32(0) {
		goto L49
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+64)) = v237
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_11), v185-int32(-64))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v334 = int32(2013)
	goto L56
L64:
	;
	if v273 != int32(2) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v279 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v292 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v185)+88)))
	v293 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v185)+92)))
	if base.Ui64(v198-int64(1)) < base.Ui64(v292|v293<<(uint(int64(32))%64)) {
		goto L49
	} else {
		goto L71
	}
L68:
	;
	if v279 == int32(0) {
		goto L49
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+32)) = v185 + int32(96)
	F_errmsg(m, int32(_a_F_CheckPointGuts_12), v185+int32(32))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v334 = int32(2029)
	goto L56
L71:
	;
	v300 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v300 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+16)) = v185 + int32(96)
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_13), v185+int32(16))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v316 = v185 + int32(96)
	v317 = F_unlink(m, v316)
	mBase = m.M
	if int32(0) <= v317 {
		goto L49
	} else {
		goto L78
	}
L76:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_8), int32(2038), int32(_a_F_CheckPointGuts_9))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v322 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v322 == int32(0) {
		goto L49
	} else {
		goto L80
	}
L80:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v316
	F_errmsg(m, int32(_a_F_CheckPointGuts_14), v185)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v334 = int32(2050)
	goto L56
L83:
	;
	goto L49
L84:
	;
	if v342 != 0 {
		v205 = v342
		goto L47
	} else {
		goto L85
	}
L85:
	;
	goto L48
L86:
	;
	m.G0 = v185 + int32(1152)
	v367 = m.G0
	v369 = v367 - int32(1216)
	m.G0 = v369
	v371 = F_GetRedoRecPtr(m)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v373 = F_ReplicationSlotsComputeLogicalRestartLSN(m)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v376 = F_AllocateDir(m, int32(_a_F_CheckPointGuts_15))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L93
	}
L89:
	;
	v697 = m.G0
	v699 = v697 - int32(112)
	m.G0 = v699
	*(*int32)(unsafe.Add(mBase, uint32(v699)+108)) = int32(307747550)
	v704 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	if v704 != 0 {
		goto L183
	} else {
		goto L184
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L175
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L171
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L168
	}
L93:
	;
	v379 = F_ReadDir(m, v376, int32(_a_F_CheckPointGuts_15))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v379 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if base.Ui64(v371) < base.Ui64(v373) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	F_FreeDir(m, v376)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L166
	}
L98:
	;
	v382 = v371
	goto L100
L99:
	;
	v382 = v373
	goto L100
L100:
	;
	v395 = v379
	goto L101
L101:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+19)))
	if v407 != int32(46) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L97
L103:
	;
	v616 = F_ReadDir(m, v376, int32(_a_F_CheckPointGuts_15))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L164
	}
L104:
	;
	v420 = v395 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+132)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v369)+128)) = int32(_a_F_CheckPointGuts_15)
	v425 = v369 + int32(160)
	v430 = F_pg_snprintf(m, v425, int32(1044), int32(_a_F_CheckPointGuts_5), v369+int32(128))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L109
	}
L105:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+20)))
	if v410 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+20)))
	if v413 != int32(46) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+21)))
	if v416 == int32(0) {
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	v434 = F_get_dirent_type(m, v425, v395, int32(0), int32(14))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v434&int32(-3) != 0 {
		goto L103
	} else {
		goto L111
	}
L111:
	;
	v438 = int32(_a_F_CheckPointGuts_16)
	goto L114
L112:
	;
	if v476-v477 != 0 {
		goto L103
	} else {
		goto L125
	}
L114:
	;
	goto L115
L115:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	if v445 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v446 = v420
	v447 = v438
	v448 = int32(4)
	v449 = v445
	goto L120
L117:
	;
	v472 = v438
	v476 = int32(0)
	goto L118
L118:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472))))
	goto L112
L119:
	;
	v472 = v467
	v476 = v469
	goto L118
L120:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447))))
	if base.B2i32(v449 != v451)|base.B2i32(v451 == int32(0)) != 0 {
		v467 = v447
		v469 = v449
		goto L119
	} else {
		goto L122
	}
L121:
	;
	v467 = v461
	v469 = int32(0)
	goto L119
L122:
	;
	v457 = v448 - int32(1)
	if v457 == int32(0) {
		v467 = v447
		v469 = v449
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v460 = int32(1)
	v461 = v447 + v460
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+1)))
	if v462 != 0 {
		v446 = v446 + v460
		v447 = v461
		v448 = v457
		v449 = v462
		goto L120
	} else {
		goto L124
	}
L124:
	;
	goto L121
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369+int32(116)))) = v369 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v369+int32(112)))) = v369 + int32(148)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+108)) = v369 + int32(136)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+104)) = v369 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+100)) = v369 + int32(152)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+96)) = v369 + int32(156)
	v506 = F_sscanf(m, v420, int32(_a_F_CheckPointGuts_17), v369+int32(96))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v506 != int32(6) {
		goto L92
	} else {
		goto L127
	}
L127:
	;
	v510 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v369)+136)))
	v511 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v369)+140)))
	if base.Ui64(v510|v511<<(uint(int64(32))%64)) <= base.Ui64(v382-int64(1)) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v518 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v554 = v369 + int32(160)
	v556 = F_OpenTransientFile(m, v554, int32(2))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L142
	}
L131:
	;
	if v518 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+64)) = v425
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_18), v369-int32(-64))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v532 = v369 + int32(160)
	v533 = F_unlink(m, v532)
	mBase = m.M
	if int32(0) <= v533 {
		goto L103
	} else {
		goto L137
	}
L135:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_19), int32(1213), int32(_a_F_CheckPointGuts_20))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
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
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+48)) = v532
	F_errmsg(m, int32(_a_F_CheckPointGuts_14), v369+int32(48))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_19), int32(1217), int32(_a_F_CheckPointGuts_20))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
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
	if v556 < int32(0) {
		goto L91
	} else {
		goto L143
	}
L143:
	;
	v561 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = int32(167772196)
	v566 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])))
	if v566 != int32(1) {
		v580 = int32(0)
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v607))) = int32(0)
	v610 = F_CloseTransientFile(m, v556)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L162
	}
L145:
	;
	if v580 == int32(0) {
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
	v571 = F_fsync(m, v556)
	mBase = m.M
	if v571 != int32(-1) {
		v580 = v571
		goto L146
	} else {
		goto L150
	}
L149:
	;
	v580 = int32(-1)
	goto L146
L150:
	;
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	if v575 == int32(27) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[8])))
	if v586 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v589 = F_errstart(m, v587, int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L157
	}
L154:
	;
	v587 = int32(21)
	goto L156
L155:
	;
	v587 = int32(24)
	goto L156
L156:
	;
	goto L153
L157:
	;
	if v589 == int32(0) {
		goto L144
	} else {
		goto L158
	}
L158:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+32)) = v554
	F_errmsg(m, int32(_a_F_CheckPointGuts_21), v369+int32(32))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_19), int32(1243), int32(_a_F_CheckPointGuts_20))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	goto L144
L162:
	;
	if v610 != 0 {
		goto L90
	} else {
		goto L163
	}
L163:
	;
	goto L103
L164:
	;
	if v616 != 0 {
		v395 = v616
		goto L101
	} else {
		goto L165
	}
L165:
	;
	goto L102
L166:
	;
	F_fsync_fname(m, int32(_a_F_CheckPointGuts_15), int32(1))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	m.G0 = v369 + int32(1216)
	goto L89
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+80)) = v420
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_22), v369+int32(80))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_19), int32(1207), int32(_a_F_CheckPointGuts_20))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
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
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v369 + int32(160)
	F_errmsg(m, int32(_a_F_CheckPointGuts_23), v369)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_19), int32(1232), int32(_a_F_CheckPointGuts_20))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
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
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+16)) = v369 + int32(160)
	F_errmsg(m, int32(_a_F_CheckPointGuts_24), v369+int32(16))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_19), int32(1249), int32(_a_F_CheckPointGuts_20))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
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
	v1005 = m.G0
	v1006 = int32(16)
	v1007 = v1005 - v1006
	m.G0 = v1007
	F_gettimeofday(m, v1007)
	mBase = m.M
	v1010 = *(*int64)(unsafe.Add(mBase, uint32(v1007)))
	v1011 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1007)+8)))
	m.G0 = v1007 + v1006
	goto L251
L180:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L247
	}
L181:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L243
	}
L182:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L239
	}
L183:
	;
	v706 = F_unlink(m, int32(_a_F_CheckPointGuts_25))
	mBase = m.M
	if v706 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	goto L185
L185:
	;
	m.G0 = v699 + int32(112)
	goto L179
L186:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	if v710 != int32(44) {
		goto L182
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v715 = F_OpenTransientFile(m, int32(_a_F_CheckPointGuts_25), int32(193))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L190
	}
L189:
	;
	goto L188
L190:
	;
	if v715 < int32(0) {
		goto L181
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])) = int32(0)
	v724 = int32(4)
	v725 = F_write(m, v715, v699+int32(108), v724)
	mBase = m.M
	if v725 != v724 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	if v729 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L194
L194:
	;
	v757 = m.Env.Pgmem_crc32c(m, int32(-1), v699+int32(108), int32(4))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v699)+104)) = v757
	v760 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v764 = F_LWLockAcquire(m, v760+int32(_a_F_CheckPointGuts_26), int32(1))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L202
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])) = int32(51)
	goto L197
L196:
	;
	goto L197
L197:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+64)) = int32(_a_F_CheckPointGuts_25)
	F_errmsg(m, int32(_a_F_CheckPointGuts_27), v699-int32(-64))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_28), int32(657), int32(_a_F_CheckPointGuts_29))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
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
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	if int32(0) < v767 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[9]))
	v775 = int32(0)
	v776 = v757
	v778 = v767
	v779 = v771
	goto L206
L204:
	;
	v864 = v757
	goto L205
L205:
	;
	v879 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v879+int32(_a_F_CheckPointGuts_26))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L225
	}
L206:
	;
	v792 = v779 + v775<<(uint(int32(6))%32)
	v793 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v792))))
	if v793 != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v864 = v854
	goto L205
L208:
	;
	v794 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v699)+96)) = v794
	*(*int64)(unsafe.Add(mBase, uint32(v699)+88)) = v794
	v799 = v792 + int32(44)
	v801 = F_LWLockAcquire(m, v799, int32(1))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	v854 = v776
	v855 = v778
	v856 = v779
	goto L210
L210:
	;
	v858 = v775 + int32(1)
	if v858 < v855 {
		v775 = v858
		v776 = v854
		v778 = v855
		v779 = v856
		goto L206
	} else {
		goto L224
	}
L211:
	;
	v803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v792))))
	*(*uint16)(unsafe.Add(mBase, uint32(v699)+88)) = uint16(v803)
	v805 = *(*int64)(unsafe.Add(mBase, uint32(v792)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v699)+96)) = v805
	v807 = *(*int64)(unsafe.Add(mBase, uint32(v792)+16))
	F_LWLockRelease(m, v799)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_XLogFlush(m, v807)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])) = int32(0)
	v817 = int32(16)
	v818 = F_write(m, v715, v699+int32(88), v817)
	mBase = m.M
	if v818 != v817 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	if v822 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	v849 = m.Env.Pgmem_crc32c(m, v776, v699+int32(88), int32(16))
	mBase = m.M
	v851 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[9]))
	v853 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	v854 = v849
	v855 = v853
	v856 = v851
	goto L210
L217:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])) = int32(51)
	goto L219
L218:
	;
	goto L219
L219:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+48)) = int32(_a_F_CheckPointGuts_25)
	F_errmsg(m, int32(_a_F_CheckPointGuts_27), v699+int32(48))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_28), int32(699), int32(_a_F_CheckPointGuts_29))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v699)+104)) = v864 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])) = int32(0)
	v892 = int32(4)
	v893 = F_write(m, v715, v699+int32(104), v892)
	mBase = m.M
	if v893 != v892 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v897 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	if v897 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	goto L228
L228:
	;
	v921 = F_CloseTransientFile(m, v715)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L236
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])) = int32(51)
	goto L231
L230:
	;
	goto L231
L231:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+32)) = int32(_a_F_CheckPointGuts_25)
	F_errmsg(m, int32(_a_F_CheckPointGuts_27), v699+int32(32))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_28), int32(718), int32(_a_F_CheckPointGuts_29))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
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
	if v921 != 0 {
		goto L180
	} else {
		goto L237
	}
L237:
	;
	v926 = F_durable_rename(m, int32(_a_F_CheckPointGuts_25), int32(_a_F_CheckPointGuts_30), int32(24))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
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
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+80)) = int32(_a_F_CheckPointGuts_25)
	F_errmsg(m, int32(_a_F_CheckPointGuts_14), v699+int32(80))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_28), int32(633), int32(_a_F_CheckPointGuts_29))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
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
	v972 = m.ExcPending
	if v972 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699))) = int32(_a_F_CheckPointGuts_25)
	F_errmsg(m, int32(_a_F_CheckPointGuts_31), v699)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_28), int32(645), int32(_a_F_CheckPointGuts_29))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
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
	v988 = m.ExcPending
	if v988 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+16)) = int32(_a_F_CheckPointGuts_25)
	F_errmsg(m, int32(_a_F_CheckPointGuts_24), v699+int32(16))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_28), int32(725), int32(_a_F_CheckPointGuts_29))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
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
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[10])) = v1011 + v1010*int64(1000000) - int64(946684800000000)
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_32))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_33))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_34))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_35))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_36))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v1041 = F_LWLockAcquire(m, v1037+int32(_a_F_CheckPointGuts_37), int32(0))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[11]))
	v1045 = *(*int64)(unsafe.Add(mBase, uint32(v1044)))
	if v1045 < int64(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1090 = int32(0)
	v1094 = m.G0
	v1096 = v1094 - int32(_a_F_CheckPointGuts_38)
	m.G0 = v1096
	v1099 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[12]))
	if v1099 <= v1090 {
		goto L273
	} else {
		goto L274
	}
L259:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v1049+int32(_a_F_CheckPointGuts_37))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1044)+12))
	if v1054 != 0 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L258
L263:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v1076+int32(_a_F_CheckPointGuts_37))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L270
	}
L264:
	;
	v1055 = int32(10)
	v1062 = base.I32_wrap_i64(v1045) << (uint(v1055) % 32)
	if (v1054&int32(-1024)-v1062-int32(1023))&(v1054-v1062) < int32(0) {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1044))) = int64(-1)
	v1074 = v1045
	goto L263
L267:
	;
	v1070 = base.I64_extend_i32_u(int32(base.Ui32(v1054) >> (uint(v1055) % 32)))
	goto L269
L268:
	;
	v1070 = v1045
	goto L269
L269:
	;
	v1074 = v1070
	goto L263
L270:
	;
	F_SimpleLruTruncate(m, int32(_a_F_CheckPointGuts_39), v1074)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_39))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	goto L258
L273:
	;
	m.G0 = v1096 + int32(_a_F_CheckPointGuts_38)
	v1816 = m.G0
	v1817 = int32(16)
	v1818 = v1816 - v1817
	m.G0 = v1818
	F_gettimeofday(m, v1818)
	mBase = m.M
	v1821 = *(*int64)(unsafe.Add(mBase, uint32(v1818)))
	v1822 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1818)+8)))
	m.G0 = v1818 + v1817
	goto L417
L274:
	;
	if l1&int32(19) != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1106 = int64(-8388609)
	goto L277
L276:
	;
	v1106 = int64(-2155872257)
	goto L277
L277:
	;
	v1114 = v1090
	v1116 = v1090
	goto L278
L278:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[13]))
	v1129 = v1126 + v1114*int32(56)
	v1130 = int64(4194304)
	v1132 = base.AtomicRmwOr64(m, v1129, int32(24), v1130)
	if v1132&v1130 != int64(0) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	if v1287 == int32(0) {
		goto L273
	} else {
		goto L318
	}
L280:
	;
	v1148 = v1132
	goto L283
L281:
	;
	v1258 = v1132
	goto L282
L282:
	;
	if v1258|v1106 == int64(-1) {
		goto L304
	} else {
		goto L305
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1096)+28)) = int32(_a_F_CheckPointGuts_40)
	*(*int32)(unsafe.Add(mBase, uint32(v1096)+24)) = int32(_a_F_CheckPointGuts_41)
	*(*int32)(unsafe.Add(mBase, uint32(v1096)+20)) = int32(_a_F_CheckPointGuts_42)
	*(*int32)(unsafe.Add(mBase, uint32(v1096)+16)) = int32(0)
	v1163 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1096)+8)) = v1163
	if v1148&int64(4194304) != v1163 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	v1258 = v1242
	goto L282
L285:
	;
	goto L288
L286:
	;
	goto L287
L287:
	;
	v1220 = int32(_a_F_CheckPointGuts_43)
	v1221 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14]))
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1096+int32(8))+8))
	if v1223 == int32(0) {
		goto L295
	} else {
		goto L296
	}
L288:
	;
	F_perform_spin_delay(m, v1096+int32(8))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L1
	} else {
		goto L290
	}
L289:
	;
	goto L287
L290:
	;
	v1191 = int64(0)
	v1194 = base.AtomicRmwCmpxchg64(m, v1129, int32(24), v1191, v1191)
	if v1194&int64(4194304) != v1191 {
		goto L288
	} else {
		goto L291
	}
L291:
	;
	goto L289
L292:
	;
	v1240 = int64(4194304)
	v1242 = base.AtomicRmwOr64(m, v1129, int32(24), v1240)
	if v1242&v1240 != int64(0) {
		v1148 = v1242
		goto L283
	} else {
		goto L303
	}
L293:
	;
	goto L292
L294:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14])) = v1238
	goto L293
L295:
	;
	if int32(999) < v1221 {
		goto L293
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	if v1221 < int32(11) {
		goto L293
	} else {
		goto L302
	}
L298:
	;
	v1228 = int32(900)
	if v1228 <= v1221 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1231 = v1228
	goto L301
L300:
	;
	v1231 = v1221
	goto L301
L301:
	;
	v1238 = v1231 + int32(100)
	goto L294
L302:
	;
	v1238 = v1221 - int32(1)
	goto L294
L303:
	;
	goto L284
L304:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	v1273 = v1270 + v1116*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v1273)+16)) = v1114
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1129)))
	*(*int32)(unsafe.Add(mBase, uint32(v1273))) = v1275
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1273)+4)) = v1277
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1273)+8)) = v1279
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1273)+12)) = v1281
	v1287 = v1116 + int32(1)
	v1288 = int64(1073741824)
	goto L306
L305:
	;
	v1287 = v1116
	v1288 = int64(0)
	goto L306
L306:
	;
	v1290 = v1258 | int64(4194304)
	v1295 = base.AtomicRmwCmpxchg64(m, v1129, int32(24), v1290, v1288|v1258&int64(-4194305))
	if v1295 != v1290 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1308 = v1295
	goto L310
L308:
	;
	goto L309
L309:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[16]))
	if v1340 != 0 {
		goto L313
	} else {
		goto L314
	}
L310:
	;
	v1319 = base.AtomicRmwCmpxchg64(m, v1129, int32(24), v1308, v1308&int64(-4194305)|v1288)
	if v1308 != v1319 {
		v1308 = v1319
		goto L310
	} else {
		goto L312
	}
L311:
	;
	goto L309
L312:
	;
	goto L311
L313:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	v1344 = v1114 + int32(1)
	v1346 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[12]))
	if v1344 < v1346 {
		v1114 = v1344
		v1116 = v1287
		goto L278
	} else {
		goto L317
	}
L316:
	;
	goto L315
L317:
	;
	goto L279
L318:
	;
	v1350 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1096)+12)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v1096)+8)) = int32(_a_F_CheckPointGuts_44)
	v1355 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	F_sort_checkpoint_bufferids(m, v1355, v1287)
	mBase = m.M
	if v1350 < v1287 {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	F_binaryheap_build(m, v1487)
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L1
	} else {
		goto L350
	}
L320:
	;
	v1359 = int32(0)
	v1363 = v1090
	v1364 = v1090
	v1366 = v1359
	v1368 = v1359
	goto L323
L321:
	;
	goto L322
L322:
	;
	v1476 = int32(0)
	v1480 = F_binaryheap_allocate(m, v1476, int32(1164), v1476)
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L1
	} else {
		goto L349
	}
L323:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1381+v1366*int32(20))))
	if v1385 == v1364 {
		goto L326
	} else {
		goto L327
	}
L324:
	;
	v1437 = int32(0)
	v1440 = F_binaryheap_allocate(m, v1421, int32(1164), v1437)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L343
	}
L325:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+24)) = v1426 + int32(1)
	v1431 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[16]))
	if v1431 != 0 {
		goto L338
	} else {
		goto L339
	}
L326:
	;
	v1387 = v1364
	goto L328
L327:
	;
	v1387 = int32(0)
	goto L328
L328:
	;
	if v1387 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1391 = v1363 + int32(1)
	v1393 = v1391 * int32(40)
	if v1368 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L330:
	;
	goto L331
L331:
	;
	v1416 = int32(40)
	v1421 = v1363
	v1422 = v1364
	v1424 = v1368
	v1425 = v1368 + v1363*v1416 - v1416
	goto L325
L332:
	;
	v1403 = v1400 + v1363*int32(40)
	v1404 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1403)+32)) = v1404
	*(*int64)(unsafe.Add(mBase, uint32(v1403))) = v1404
	*(*int64)(unsafe.Add(mBase, uint32(v1403)+24)) = v1404
	*(*int64)(unsafe.Add(mBase, uint32(v1403)+16)) = v1404
	*(*int64)(unsafe.Add(mBase, uint32(v1403)+8)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+32)) = v1366
	*(*int32)(unsafe.Add(mBase, uint32(v1403))) = v1385
	v1421 = v1391
	v1422 = v1385
	v1424 = v1400
	v1425 = v1403
	goto L325
L333:
	;
	v1396 = F_palloc(m, v1393)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L1
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v1398 = F_repalloc(m, v1368, v1393)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L1
	} else {
		goto L337
	}
L336:
	;
	v1400 = v1396
	goto L332
L337:
	;
	v1400 = v1398
	goto L332
L338:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L1
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	v1435 = v1366 + int32(1)
	if v1435 != v1287 {
		v1363 = v1421
		v1364 = v1422
		v1366 = v1435
		v1368 = v1424
		goto L323
	} else {
		goto L342
	}
L341:
	;
	goto L340
L342:
	;
	goto L324
L343:
	;
	if v1421 <= int32(0) {
		v1487 = v1440
		v1489 = v1424
		goto L319
	} else {
		goto L344
	}
L344:
	;
	v1453 = v1437
	goto L345
L345:
	;
	v1465 = v1424 + v1453*int32(40)
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v1465)+16)) = base.F64_div(base.F64_convert_i32_u(v1287), base.F64_convert_i32_s(v1466))
	F_binaryheap_add_unordered(m, v1440, base.I64_extend_i32_u(v1465))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L1
	} else {
		goto L347
	}
L346:
	;
	v1487 = v1440
	v1489 = v1424
	goto L319
L347:
	;
	v1474 = v1453 + int32(1)
	if v1474 != v1421 {
		v1453 = v1474
		goto L345
	} else {
		goto L348
	}
L348:
	;
	goto L346
L349:
	;
	v1487 = v1480
	v1489 = v1476
	goto L319
L350:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1487)))
	if v1502 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	F_IssuePendingWritebacks(m, v1096+int32(8), int32(3))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L1
	} else {
		goto L414
	}
L352:
	;
	v1761 = int32(0)
	goto L351
L353:
	;
	goto L354
L354:
	;
	v1507 = int32(0)
	v1511 = v1507
	v1518 = v1507
	goto L355
L355:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[13]))
	v1530 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	v1531 = *(*int64)(unsafe.Add(mBase, uint32(v1487)+24))
	v1532 = base.I32_wrap_i64(v1531)
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+32))
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1530+v1533*int32(20))+16))
	v1541 = int64(0)
	v1544 = base.AtomicRmwCmpxchg64(m, v1528+v1537*int32(56), int32(24), v1541, v1541)
	if v1544&int64(1073741824) == v1541 {
		v1566 = v1511
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v1761 = v1566
	goto L351
L357:
	;
	v1567 = *(*float64)(unsafe.Add(mBase, uint32(v1532)+16))
	v1568 = *(*float64)(unsafe.Add(mBase, uint32(v1532)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1532)+8)) = base.F64_add(v1567, v1568)
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+28))
	v1572 = int32(1)
	v1573 = v1571 + v1572
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+28)) = v1573
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+32)) = v1575 + v1572
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+24))
	if v1579 == v1573 {
		goto L362
	} else {
		goto L363
	}
L358:
	;
	v1552 = F_SyncOneBuffer(m, v1537, int32(0), v1096+int32(8))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	if v1552&int32(1) == int32(0) {
		v1566 = v1511
		goto L357
	} else {
		goto L360
	}
L360:
	;
	v1558 = int32(_a_F_CheckPointGuts_45)
	v1560 = *(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[17])) = v1560 + int64(1)
	v1566 = v1511 + int32(1)
	goto L357
L361:
	;
	v1588 = v1518 + int32(1)
	v1592 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[18]))
	if v1592 != int32(11) {
		goto L367
	} else {
		goto L368
	}
L362:
	;
	v1581 = F_binaryheap_remove_first(m, v1487)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L1
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	F_binaryheap_replace_first(m, v1487, v1531&int64(4294967295))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L1
	} else {
		goto L366
	}
L365:
	;
	goto L361
L366:
	;
	goto L361
L367:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1487)))
	if v1758 != 0 {
		v1511 = v1566
		v1518 = v1588
		goto L355
	} else {
		goto L413
	}
L368:
	;
	if l1&int32(4) != 0 {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	v1750 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[16]))
	if v1750 == int32(0) {
		goto L367
	} else {
		goto L411
	}
L370:
	;
	v1733 = int32(_a_F_CheckPointGuts_46)
	v1735 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[19]))
	v1737 = v1735 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[19])) = v1737
	if int32(0) < v1737 {
		goto L369
	} else {
		goto L409
	}
L371:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[20]))
	if v1598 != 0 {
		goto L370
	} else {
		goto L372
	}
L372:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[21]))
	if v1600 != 0 {
		goto L370
	} else {
		goto L373
	}
L373:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[22]))
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	if v1603&int32(4) != 0 {
		goto L370
	} else {
		goto L374
	}
L374:
	;
	v1606 = m.G0
	v1608 = v1606 - int32(16)
	m.G0 = v1608
	v1612 = *(*float64)(unsafe.Add(mBase, _c_F_CheckPointGuts[23]))
	v1613 = base.F64_mul(base.F64_div(base.F64_convert_i32_s(v1588), base.F64_convert_i32_s(v1287)), v1612)
	v1615 = *(*float64)(unsafe.Add(mBase, _c_F_CheckPointGuts[24]))
	if base.F64_lt(v1613, v1615) != 0 {
		v1675 = int32(0)
		goto L375
	} else {
		goto L376
	}
L375:
	;
	m.G0 = v1608 + int32(16)
	if v1675 == int32(0) {
		goto L370
	} else {
		goto L391
	}
L376:
	;
	v1619 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[25])))
	if v1619 == int32(1) {
		goto L379
	} else {
		goto L380
	}
L377:
	;
	v1637 = *(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[26]))
	v1641 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[27]))
	v1645 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[28]))
	v1647 = base.F64_div(base.F64_div(base.F64_convert_i64_u(v1635-v1637), base.F64_convert_i32_s(v1641)), base.F64_convert_i32_s(v1645))
	if base.F64_lt(v1613, v1647) == int32(0) {
		goto L387
	} else {
		goto L388
	}
L378:
	;
	if v1629 != 0 {
		goto L382
	} else {
		goto L383
	}
L379:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[29]))
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+308))
	v1627 = base.B2i32(v1625 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[25])) = uint8(v1627)
	v1629 = v1627
	goto L381
L380:
	;
	v1629 = int32(0)
	goto L381
L381:
	;
	goto L378
L382:
	;
	v1631 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L1
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	v1633 = F_GetInsertRecPtr(m)
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L1
	} else {
		goto L386
	}
L385:
	;
	v1635 = v1631
	goto L377
L386:
	;
	v1635 = v1633
	goto L377
L387:
	;
	F_gettimeofday(m, v1608)
	mBase = m.M
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+8))
	v1657 = *(*int64)(unsafe.Add(mBase, uint32(v1608)))
	v1659 = *(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[30]))
	v1664 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[31]))
	v1666 = base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v1653), float64(1e+06)), base.F64_convert_i64_s(v1657-v1659)), base.F64_convert_i32_s(v1664))
	if base.F64_lt(v1613, v1666) == int32(0) {
		v1675 = int32(1)
		goto L375
	} else {
		goto L390
	}
L388:
	;
	v1670 = v1647
	goto L389
L389:
	;
	*(*float64)(unsafe.Add(mBase, _c_F_CheckPointGuts[24])) = v1670
	v1675 = int32(0)
	goto L375
L390:
	;
	v1670 = v1666
	goto L389
L391:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[32]))
	if v1682 != 0 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[32])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L1
	} else {
		goto L404
	}
L395:
	;
	F_SyncRepUpdateSyncStandbysDefined(m)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	v1695 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	if v1695 != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_47), int32(0))
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L1
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	goto L394
L402:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_48), int32(1515), int32(_a_F_CheckPointGuts_49))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	goto L401
L404:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[19])) = int32(1000)
	F_CheckArchiveTimeout(m)
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	v1716 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[33]))
	v1720 = F_WaitLatch(m, v1716, int32(41), int32(100), int32(150994945))
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[33]))
	v1724 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1723))) = v1724
	v1729 = base.AtomicRmwOr32(m, v1724, int32(_a_F_CheckPointGuts_50), v1724)
	goto L408
L408:
	;
	goto L369
L409:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[19])) = int32(1000)
	goto L369
L411:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	goto L367
L413:
	;
	goto L356
L414:
	;
	F_pfree(m, v1489)
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	F_pfree(m, v1487)
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	v1786 = int32(_a_F_CheckPointGuts_51)
	v1788 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[34]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[34])) = v1788 + v1761
	goto L273
L417:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[35])) = v1822 + v1821*int64(1000000) - int64(946684800000000)
	v1832 = int64(0)
	v1833 = int32(0)
	v1835 = m.G0
	v1837 = v1835 - int32(1152)
	m.G0 = v1837
	v1840 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[36]))
	if v1840 != 0 {
		goto L419
	} else {
		goto L420
	}
L418:
	;
	v2308 = m.G0
	v2309 = int32(16)
	v2310 = v2308 - v2309
	m.G0 = v2310
	F_gettimeofday(m, v2310)
	mBase = m.M
	v2313 = *(*int64)(unsafe.Add(mBase, uint32(v2310)))
	v2314 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2310)+8)))
	m.G0 = v2310 + v2309
	goto L520
L419:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L1
	} else {
		goto L422
	}
L420:
	;
	goto L421
L421:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2294 = m.ExcPending
	if v2294 != 0 {
		goto L1
	} else {
		goto L517
	}
L422:
	;
	v1844 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])))
	if v1844 == int32(0) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1901 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])) = uint8(v1901)
	v1903 = int32(_a_F_CheckPointGuts_52)
	v1905 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[38])))
	v1907 = v1905 + v1901
	*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[38])) = uint16(v1907)
	v1910 = v1837 + int32(1116)
	v1912 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[36]))
	F_hash_seq_init(m, v1910, v1912)
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L1
	} else {
		goto L432
	}
L424:
	;
	v1848 = v1837 + int32(1116)
	v1850 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[36]))
	F_hash_seq_init(m, v1848, v1850)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	v1853 = F_hash_seq_search(m, v1848)
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	if v1853 == int32(0) {
		goto L423
	} else {
		goto L427
	}
L427:
	;
	v1859 = v1853
	goto L428
L428:
	;
	v1876 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[38])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1859)+24)) = uint16(v1876)
	v1880 = F_hash_seq_search(m, v1837+int32(1116))
	mBase = m.M
	v1881 = m.ExcPending
	if v1881 != 0 {
		goto L1
	} else {
		goto L430
	}
L429:
	;
	goto L423
L430:
	;
	if v1880 != 0 {
		v1859 = v1880
		goto L428
	} else {
		goto L431
	}
L431:
	;
	goto L429
L432:
	;
	v1915 = F_hash_seq_search(m, v1910)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L1
	} else {
		goto L434
	}
L433:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L1
	} else {
		goto L514
	}
L434:
	;
	if v1915 != 0 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1920 = v1915
	v1922 = int32(10)
	v1925 = v1833
	v1929 = v1832
	v1930 = v1832
	goto L438
L436:
	;
	v2255 = v1833
	v2259 = v1832
	v2260 = v1832
	goto L437
L437:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[39])) = v2260
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[40])) = v2259
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[41])) = v2255
	v2273 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])) = uint8(v2273)
	m.G0 = v1837 + int32(1152)
	goto L418
L438:
	;
	v1936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+24)))
	v1938 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[38])))
	if v1936 != v1938 {
		goto L440
	} else {
		goto L441
	}
L439:
	;
	v2255 = v2233
	v2259 = v2237
	v2260 = v2238
	goto L437
L440:
	;
	v1941 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])))
	if v1941 != int32(1) {
		v2204 = v1922
		v2207 = v1925
		v2211 = v1929
		v2212 = v1930
		goto L443
	} else {
		goto L444
	}
L441:
	;
	v2230 = v1922
	v2233 = v1925
	v2237 = v1929
	v2238 = v1930
	goto L442
L442:
	;
	v2246 = F_hash_seq_search(m, v1837+int32(1116))
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L1
	} else {
		goto L512
	}
L443:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[36]))
	v2222 = F_hash_search(m, v2219, v1920, int32(2), int32(0))
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L1
	} else {
		goto L510
	}
L444:
	;
	v1945 = v1922 - int32(1)
	if v1945 <= int32(0) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L1
	} else {
		goto L448
	}
L446:
	;
	v1951 = v1945
	goto L447
L447:
	;
	v1952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1920)+26)))
	if v1952 != 0 {
		v2204 = v1951
		v2207 = v1925
		v2211 = v1929
		v2212 = v1930
		goto L443
	} else {
		goto L449
	}
L448:
	;
	v1951 = int32(10)
	goto L447
L449:
	;
	F___clock_gettime(m, int32(1), v1837+int32(1136))
	mBase = m.M
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1837)+1144))
	v1958 = *(*int64)(unsafe.Add(mBase, uint32(v1837)+1136))
	v1960 = v1837 + int32(80)
	v1961 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1920))))
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v1961*int32(12))+uint32(_c_F_CheckPointGuts[42])))
	v1967 = m.T0[v1966].(func(*base.Module, int32, int32) int32)(m, v1920, v1960)
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L1
	} else {
		goto L451
	}
L450:
	;
	v2155 = int32(1)
	F___clock_gettime(m, v2155, v1837+int32(1136))
	mBase = m.M
	v2159 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1837)+1144)))
	v2162 = *(*int64)(unsafe.Add(mBase, uint32(v1837)+1136))
	v2168 = base.I64_div_s(v2159-base.I64_extend_i32_s(v2140)+(v2162-v2150)*int64(1000000000), int64(1000))
	v2171 = v1925 + v2155
	v2173 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[43])))
	if v2173 != v2155 {
		goto L501
	} else {
		goto L502
	}
L451:
	;
	if v1967 == int32(0) {
		v2140 = v1957
		v2141 = v1951
		v2150 = v1958
		goto L450
	} else {
		goto L452
	}
L452:
	;
	v1973 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	if v1973 == int32(44) {
		goto L455
	} else {
		goto L456
	}
L453:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		goto L1
	} else {
		goto L471
	}
L454:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_53), v2012, int32(_a_F_CheckPointGuts_54))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L470
	}
L455:
	;
	v1978 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L1
	} else {
		goto L458
	}
L456:
	;
	goto L457
L457:
	;
	v1994 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[8])))
	if v1994 != 0 {
		goto L463
	} else {
		goto L464
	}
L458:
	;
	if v1978 == int32(0) {
		goto L453
	} else {
		goto L459
	}
L459:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+48)) = v1960
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_55), v1837+int32(48))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L1
	} else {
		goto L461
	}
L461:
	;
	v2012 = int32(453)
	goto L454
L462:
	;
	v1997 = F_errstart(m, v1995, int32(0))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L1
	} else {
		goto L466
	}
L463:
	;
	v1995 = int32(21)
	goto L465
L464:
	;
	v1995 = int32(24)
	goto L465
L465:
	;
	goto L462
L466:
	;
	if v1997 == int32(0) {
		goto L453
	} else {
		goto L467
	}
L467:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+64)) = v1837 + int32(80)
	F_errmsg(m, int32(_a_F_CheckPointGuts_21), v1837-int32(-64))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v2012 = int32(448)
	goto L454
L470:
	;
	goto L453
L471:
	;
	v2019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1920)+26)))
	if v2019 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v2026 = int32(1)
	goto L475
L473:
	;
	goto L474
L474:
	;
	v2204 = int32(10)
	v2207 = v1925
	v2211 = v1929
	v2212 = v1930
	goto L443
L475:
	;
	F___clock_gettime(m, int32(1), v1837+int32(1136))
	mBase = m.M
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v1837)+1144))
	v2045 = *(*int64)(unsafe.Add(mBase, uint32(v1837)+1136))
	v2048 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1920))))
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v2048*int32(12))+uint32(_c_F_CheckPointGuts[42])))
	v2054 = m.T0[v2053].(func(*base.Module, int32, int32) int32)(m, v1920, v1837+int32(80))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L1
	} else {
		goto L477
	}
L476:
	;
	goto L474
L477:
	;
	if v2054 == int32(0) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v2140 = v2044
	v2141 = int32(10)
	v2150 = v2045
	goto L450
L479:
	;
	goto L480
L480:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	v2064 = int32(0)
	if base.B2i32(v2061 == int32(44))&base.B2i32(v2026 <= v2064) == v2064 {
		goto L483
	} else {
		goto L484
	}
L481:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L1
	} else {
		goto L499
	}
L482:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_53), v2107, int32(_a_F_CheckPointGuts_54))
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L1
	} else {
		goto L498
	}
L483:
	;
	v2072 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[8])))
	if v2072 != 0 {
		goto L487
	} else {
		goto L488
	}
L484:
	;
	goto L485
L485:
	;
	v2092 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L1
	} else {
		goto L494
	}
L486:
	;
	v2075 = F_errstart(m, v2073, int32(0))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L1
	} else {
		goto L490
	}
L487:
	;
	v2073 = int32(21)
	goto L489
L488:
	;
	v2073 = int32(24)
	goto L489
L489:
	;
	goto L486
L490:
	;
	if v2075 == int32(0) {
		goto L481
	} else {
		goto L491
	}
L491:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2080 = m.ExcPending
	if v2080 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+16)) = v1837 + int32(80)
	F_errmsg(m, int32(_a_F_CheckPointGuts_21), v1837+int32(16))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	v2107 = int32(448)
	goto L482
L494:
	;
	if v2092 == int32(0) {
		goto L481
	} else {
		goto L495
	}
L495:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+32)) = v1837 + int32(80)
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_55), v1837+int32(32))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v2107 = int32(453)
	goto L482
L498:
	;
	goto L481
L499:
	;
	v2115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1920)+26)))
	if v2115 != int32(1) {
		v2026 = v2026 + int32(1)
		goto L475
	} else {
		goto L500
	}
L500:
	;
	goto L476
L501:
	;
	if base.Ui64(v1929) < base.Ui64(v2168) {
		goto L507
	} else {
		goto L508
	}
L502:
	;
	v2178 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	if v2178 == int32(0) {
		goto L501
	} else {
		goto L504
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1837))) = v2171
	*(*float64)(unsafe.Add(mBase, uint32(v1837)+8)) = base.F64_div(base.F64_convert_i64_u(v2168), float64(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+4)) = v1837 + int32(80)
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_56), v1837)
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_53), int32(433), int32(_a_F_CheckPointGuts_54))
	mBase = m.M
	v2197 = m.ExcPending
	if v2197 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	goto L501
L507:
	;
	v2198 = v2168
	goto L509
L508:
	;
	v2198 = v1929
	goto L509
L509:
	;
	v2204 = v2141
	v2207 = v2171
	v2211 = v2198
	v2212 = v1930 + v2168
	goto L443
L510:
	;
	if v2222 == int32(0) {
		goto L433
	} else {
		goto L511
	}
L511:
	;
	v2230 = v2204
	v2233 = v2207
	v2237 = v2211
	v2238 = v2212
	goto L442
L512:
	;
	if v2246 != 0 {
		v1920 = v2246
		v1922 = v2230
		v1925 = v2233
		v1929 = v2237
		v1930 = v2238
		goto L438
	} else {
		goto L513
	}
L513:
	;
	goto L439
L514:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_57), int32(0))
	mBase = m.M
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_53), int32(466), int32(_a_F_CheckPointGuts_54))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L517:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_58), int32(0))
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_53), int32(309), int32(_a_F_CheckPointGuts_54))
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L1
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
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[44])) = v2314 + v2313*int64(1000000) - int64(946684800000000)
	v2324 = int32(0)
	v2326 = m.G0
	v2328 = v2326 - int32(1152)
	m.G0 = v2328
	v2331 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[45]))
	if v2331 <= v2324 {
		goto L525
	} else {
		goto L526
	}
L521:
	;
	return
L522:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2643 = m.ExcPending
	if v2643 != 0 {
		goto L1
	} else {
		goto L591
	}
L523:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L1
	} else {
		goto L587
	}
L524:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L1
	} else {
		goto L583
	}
L525:
	;
	m.G0 = v2328 + int32(1152)
	goto L521
L526:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v2339 = F_LWLockAcquire(m, v2335+int32(2304), int32(1))
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[46]))
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v2342)+4))
	if int32(0) < v2343 {
		goto L528
	} else {
		goto L529
	}
L528:
	;
	v2348 = v2342
	v2349 = v2324
	v2353 = v2324
	goto L531
L529:
	;
	v2533 = v2324
	goto L530
L530:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v2549+int32(2304))
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L1
	} else {
		goto L576
	}
L531:
	;
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v2348+v2353<<(uint(int32(2))%32))+8))
	v2368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+48)))
	if v2368 == int32(0) {
		goto L534
	} else {
		goto L535
	}
L532:
	;
	v2533 = v2522
	goto L530
L533:
	;
	v2527 = v2353 + int32(1)
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2521)+4))
	if v2527 < v2528 {
		v2348 = v2521
		v2349 = v2522
		v2353 = v2527
		goto L531
	} else {
		goto L575
	}
L534:
	;
	v2371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+50)))
	if v2371 != int32(1) {
		v2521 = v2348
		v2522 = v2349
		goto L533
	} else {
		goto L537
	}
L535:
	;
	goto L536
L536:
	;
	v2374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2367)+49)))
	if v2374 != 0 {
		v2521 = v2348
		v2522 = v2349
		goto L533
	} else {
		goto L538
	}
L537:
	;
	goto L536
L538:
	;
	v2375 = *(*int64)(unsafe.Add(mBase, uint32(v2367)+24))
	if base.Ui64(l0) < base.Ui64(v2375) {
		v2521 = v2348
		v2522 = v2349
		goto L533
	} else {
		goto L539
	}
L539:
	;
	v2377 = *(*int64)(unsafe.Add(mBase, uint32(v2367)+16))
	F_XlogReadTwoPhaseData(m, v2377, v2328+int32(120), v2328+int32(116))
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	v2384 = *(*int64)(unsafe.Add(mBase, uint32(v2367)+32))
	v2385 = int32(-1)
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2328)+120))
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v2328)+116))
	v2388 = m.Env.Pgmem_crc32c(m, v2385, v2386, v2387)
	mBase = m.M
	v2390 = int64(base.Ui64(v2384) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2328)+96)) = uint32(v2390)
	*(*uint32)(unsafe.Add(mBase, uint32(v2328)+100)) = uint32(v2384)
	*(*int32)(unsafe.Add(mBase, uint32(v2328)+124)) = v2388 ^ v2385
	v2397 = v2328 + int32(128)
	v2402 = F_pg_snprintf(m, v2397, int32(1024), int32(_a_F_CheckPointGuts_59), v2328+int32(96))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	v2405 = F_OpenTransientFile(m, v2397, int32(577))
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	if v2405 < int32(0) {
		goto L524
	} else {
		goto L543
	}
L543:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])) = int32(0)
	v2413 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v2413))) = int32(167772226)
	v2416 = F_write(m, v2405, v2386, v2387)
	mBase = m.M
	if v2416 != v2387 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v2419 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	if v2419 == int32(0) {
		goto L547
	} else {
		goto L548
	}
L545:
	;
	goto L546
L546:
	;
	v2446 = int32(4)
	v2447 = F_write(m, v2405, v2328+int32(124), v2446)
	mBase = m.M
	if v2447 != v2446 {
		goto L554
	} else {
		goto L555
	}
L547:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])) = int32(51)
	goto L549
L548:
	;
	goto L549
L549:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2328)+80)) = v2328 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_60), v2328+int32(80))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L1
	} else {
		goto L552
	}
L552:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_61), int32(1778), int32(_a_F_CheckPointGuts_62))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L554:
	;
	v2451 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	if v2451 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L555:
	;
	goto L556
L556:
	;
	v2476 = int32(_a_F_CheckPointGuts_63)
	v2477 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[5]))
	v2478 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2477))) = v2478
	v2481 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v2481))) = int32(167772225)
	v2486 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])))
	if v2486 != int32(1) {
		v2500 = v2478
		goto L565
	} else {
		goto L566
	}
L557:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])) = int32(51)
	goto L559
L558:
	;
	goto L559
L559:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L1
	} else {
		goto L560
	}
L560:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2328)+64)) = v2328 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_60), v2328-int32(-64))
	mBase = m.M
	v2470 = m.ExcPending
	if v2470 != 0 {
		goto L1
	} else {
		goto L562
	}
L562:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_61), int32(1787), int32(_a_F_CheckPointGuts_62))
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L1
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
	if v2500 != 0 {
		goto L523
	} else {
		goto L571
	}
L565:
	;
	goto L564
L566:
	;
	goto L567
L567:
	;
	v2491 = F_fsync(m, v2405)
	mBase = m.M
	if v2491 != int32(-1) {
		v2500 = v2491
		goto L565
	} else {
		goto L569
	}
L568:
	;
	v2500 = int32(-1)
	goto L565
L569:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[7]))
	if v2495 == int32(27) {
		goto L567
	} else {
		goto L570
	}
L570:
	;
	goto L568
L571:
	;
	v2502 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v2502))) = int32(0)
	v2505 = F_CloseTransientFile(m, v2405)
	mBase = m.M
	v2506 = m.ExcPending
	if v2506 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	if v2505 != 0 {
		goto L522
	} else {
		goto L573
	}
L573:
	;
	v2507 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2367)+49)) = uint8(v2507)
	v2510 = v2367 + int32(16)
	v2511 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2510)+8)) = v2511
	*(*int64)(unsafe.Add(mBase, uint32(v2510))) = v2511
	F_pfree(m, v2386)
	mBase = m.M
	v2516 = m.ExcPending
	if v2516 != 0 {
		goto L1
	} else {
		goto L574
	}
L574:
	;
	v2520 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[46]))
	v2521 = v2520
	v2522 = v2349 + int32(1)
	goto L533
L575:
	;
	goto L532
L576:
	;
	F_fsync_fname(m, int32(_a_F_CheckPointGuts_64), int32(1))
	mBase = m.M
	v2557 = m.ExcPending
	if v2557 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	v2559 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[43])))
	if base.B2i32(v2559 != int32(1))|base.B2i32(v2533 <= int32(0)) != 0 {
		goto L525
	} else {
		goto L578
	}
L578:
	;
	v2567 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2568 = m.ExcPending
	if v2568 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	if v2567 == int32(0) {
		goto L525
	} else {
		goto L580
	}
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2328))) = v2533
	F_errmsg_plural(m, int32(_a_F_CheckPointGuts_65), int32(_a_F_CheckPointGuts_66), v2533, v2328)
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_61), int32(1898), int32(_a_F_CheckPointGuts_67))
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	goto L525
L583:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L1
	} else {
		goto L584
	}
L584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2328)+16)) = v2328 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_68), v2328+int32(16))
	mBase = m.M
	v2615 = m.ExcPending
	if v2615 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_61), int32(1766), int32(_a_F_CheckPointGuts_62))
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L1
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2328)+48)) = v2328 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_21), v2328+int32(48))
	mBase = m.M
	v2634 = m.ExcPending
	if v2634 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_61), int32(1799), int32(_a_F_CheckPointGuts_62))
	mBase = m.M
	v2639 = m.ExcPending
	if v2639 != 0 {
		goto L1
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v2645 = m.ExcPending
	if v2645 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2328)+32)) = v2328 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_24), v2328+int32(32))
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_61), int32(1805), int32(_a_F_CheckPointGuts_62))
	mBase = m.M
	v2658 = m.ExcPending
	if v2658 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_point_dt(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v91 float64
	_ = v91
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v107 float64
	_ = v107
	var v115 float64
	_ = v115
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v134 float64
	_ = v134
	var v136 float64
	_ = v136
	var v143 float64
	_ = v143
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v10 = base.F64_sub(v8, v9)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v10), v12)|base.F64_eq(base.F64_abs(v8), v12)|base.F64_eq(base.F64_abs(v9), v12) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v28 = v10
	goto L3
L3:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return float64(0)
L5:
	;
	v28 = v24
	goto L3
L6:
	;
	v39 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = base.F64_sub(v39, v40)
	v43 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v41), v43)|base.F64_eq(base.F64_abs(v39), v43)|base.F64_eq(base.F64_abs(v40), v43) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v31 != int32(453) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v34 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	return float64(0)
L10:
	;
	v55 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v57 = v41
	goto L12
L12:
	;
	if l2 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v57 = v55
	goto L12
L14:
	;
	v76 = m.G0
	v78 = v76 - int32(32)
	m.G0 = v78
	v80 = base.F64_abs(v28)
	v81 = base.F64_abs(v57)
	v84 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v80)) < base.Ui64(base.I64_reinterpret_f64(v81)))
	if base.Ui64(base.I64_reinterpret_f64(v80)) < base.Ui64(base.I64_reinterpret_f64(v81)) {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v60 != int32(453) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v63 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	return float64(0)
L18:
	;
	return v143
L19:
	;
	m.G0 = v78 + int32(32)
	goto L18
L20:
	;
	v85 = v80
	goto L22
L21:
	;
	v85 = v81
	goto L22
L22:
	;
	v86 = base.I64_reinterpret_f64(v85)
	v88 = int64(base.Ui64(v86) >> (uint(int64(52)) % 64))
	if v88 == int64(2047) {
		v143 = v85
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if base.Ui64(base.I64_reinterpret_f64(v80)) < base.Ui64(base.I64_reinterpret_f64(v81)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v91 = v81
	goto L26
L25:
	;
	v91 = v80
	goto L26
L26:
	;
	if v86 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v143 = v91
	goto L19
L28:
	;
	v94 = base.I64_reinterpret_f64(v91)
	v96 = int64(base.Ui64(v94) >> (uint(int64(52)) % 64))
	if v96 == int64(2047) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if int32(65) <= base.I32_wrap_i64(v96)-base.I32_wrap_i64(v88) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v143 = base.F64_add(v80, v81)
	goto L19
L31:
	;
	goto L32
L32:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v94) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_sq(m, v78+int32(24), v78+int32(16), v120)
	mBase = m.M
	F_sq(m, v78+int32(8), v78, v121)
	mBase = m.M
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v78)))
	v132 = *(*float64)(unsafe.Add(mBase, uint32(v78)+16))
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v78)+8))
	v136 = *(*float64)(unsafe.Add(mBase, uint32(v78)+24))
	v143 = base.F64_mul(v122, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v131, v132), v134), v136)))
	goto L19
L34:
	;
	v107 = float64(1.90109156629516e-211)
	v120 = base.F64_mul(v91, v107)
	v121 = base.F64_mul(v85, v107)
	v122 = float64(5.260135901548374e+210)
	goto L33
L35:
	;
	goto L36
L36:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v86) {
		v120 = v91
		v121 = v85
		v122 = float64(1)
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v115 = float64(5.260135901548374e+210)
	v120 = base.F64_mul(v91, v115)
	v121 = base.F64_mul(v85, v115)
	v122 = float64(1.90109156629516e-211)
	goto L33
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
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v84 float64
	_ = v84
	var v97 float64
	_ = v97
	var v98 int32
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v117 float64
	_ = v117
	var v118 int32
	_ = v118
	var v119 float64
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 float64
	_ = v131
	var v136 float64
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v163 int32
	_ = v163
	v12 = int32(0)
	v16 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v18 = base.F64_sub(v16, v17)
	v20 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v18), v20)|base.F64_eq(base.F64_abs(v16), v20)|base.F64_eq(base.F64_abs(v17), v20) == v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v37 = v18
	goto L3
L3:
	;
	v38 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v39 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = base.F64_sub(v38, v39)
	v42 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v40), v42)|base.F64_eq(base.F64_abs(v38), v42)|base.F64_eq(base.F64_abs(v39), v42) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v37 = v33
	goto L3
L6:
	;
	v55 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v57 = v40
	goto L8
L8:
	;
	v58 = int32(2)
	if l1 < v58 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v57 = v55
	goto L8
L10:
	;
	return v163
L11:
	;
	v143 = F_lseg_crossing(m, v37, v57, v136, v131)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L28
	}
L12:
	;
	v131 = v57
	v136 = v37
	v142 = v12
	goto L11
L13:
	;
	goto L14
L14:
	;
	v66 = v37
	v67 = v57
	v73 = int32(1)
	v76 = v12
	goto L15
L15:
	;
	v79 = l2 + v73<<(uint(int32(4))%32)
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v79)))
	v81 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v82 = base.F64_sub(v80, v81)
	v84 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v82), v84)|base.F64_eq(base.F64_abs(v80), v84)|base.F64_eq(base.F64_abs(v81), v84) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v131 = v119
	v136 = v99
	v142 = v124
	goto L11
L17:
	;
	v97 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	v99 = v82
	goto L19
L19:
	;
	v100 = *(*float64)(unsafe.Add(mBase, uint32(v79)+8))
	v101 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v102 = base.F64_sub(v100, v101)
	v104 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v102), v104)|base.F64_eq(base.F64_abs(v100), v104)|base.F64_eq(base.F64_abs(v101), v104) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v99 = v97
	goto L19
L21:
	;
	v117 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v119 = v102
	goto L23
L23:
	;
	v120 = F_lseg_crossing(m, v99, v119, v66, v67)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	v119 = v117
	goto L23
L25:
	;
	if v120 == int32(2147483647) {
		v163 = v58
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v124 = v120 + v76
	v126 = v73 + int32(1)
	if v126 != l1 {
		v66 = v99
		v67 = v119
		v73 = v126
		v76 = v124
		goto L15
	} else {
		goto L27
	}
L27:
	;
	goto L16
L28:
	;
	if v143 == int32(2147483647) {
		v163 = v58
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v163 = base.B2i32(v142 != int32(0)-v143)
	goto L10
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
	var v27 float64
	_ = v27
	var v39 float64
	_ = v39
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v63 float64
	_ = v63
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v77 float64
	_ = v77
	var v78 int32
	_ = v78
	var v81 float64
	_ = v81
	var v83 float64
	_ = v83
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v94 float64
	_ = v94
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v109 float64
	_ = v109
	v3 = float64(0)
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_eq(v10, v11) != 0 {
		v109 = v3
		return v109
	} else {
		v13 = base.F64_sub(v10, v11)
		v14 = base.F64_abs(v13)
		if base.F64_le(v14, float64(1e-06)) != 0 {
			v109 = v3
			return v109
		} else {
			v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_eq(v18, v19)|base.F64_le(base.F64_abs(base.F64_sub(v18, v19)), float64(1e-06)) != 0 {
				v109 = math.Float64frombits(uint64(0x7ff0000000000000))
				return v109
			} else {
				v27 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(base.F64_abs(v10), v27)|base.F64_ne(v14, v27)|base.F64_eq(base.F64_abs(v11), v27) == int32(0) {
					v39 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return float64(0)
					} else {
						v43 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						v44 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
						v45 = v43
						v46 = v39
						v47 = v44
						v49 = math.Float64frombits(uint64(0x7ff0000000000000))
						v51 = base.F64_sub(v45, v47)
						if base.F64_eq(base.F64_abs(v45), v49)|base.F64_ne(base.F64_abs(v51), v49)|base.F64_eq(base.F64_abs(v47), v49) == int32(0) {
							v63 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return float64(0)
							} else {
								v65 = v63
								if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v46)&int64(9223372036854775807)))|base.F64_ne(v65, float64(0)) == int32(0) {
									v77 = F_float_zero_divide_error_ext(m, int32(0))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return float64(0)
									} else {
										return v77
									}
								} else {
									v81 = math.Float64frombits(uint64(0x7ff0000000000000))
									v83 = base.F64_div(v46, v65)
									if base.F64_eq(base.F64_abs(v46), v81)|base.F64_ne(base.F64_abs(v83), v81) == int32(0) {
										v91 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return float64(0)
										} else {
											return v91
										}
									} else {
										v94 = float64(0)
										if base.F64_eq(v46, v94)|base.F64_ne(v83, v94)|base.F64_eq(base.F64_abs(v65), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											v109 = v83
											return v109
										} else {
											v104 = F_float_underflow_error_ext(m, int32(0))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return float64(0)
											} else {
												v109 = v104
												return v109
											}
										}
									}
								}
							}
						} else {
							v65 = v51
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v46)&int64(9223372036854775807)))|base.F64_ne(v65, float64(0)) == int32(0) {
								v77 = F_float_zero_divide_error_ext(m, int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return float64(0)
								} else {
									return v77
								}
							} else {
								v81 = math.Float64frombits(uint64(0x7ff0000000000000))
								v83 = base.F64_div(v46, v65)
								if base.F64_eq(base.F64_abs(v46), v81)|base.F64_ne(base.F64_abs(v83), v81) == int32(0) {
									v91 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return float64(0)
									} else {
										return v91
									}
								} else {
									v94 = float64(0)
									if base.F64_eq(v46, v94)|base.F64_ne(v83, v94)|base.F64_eq(base.F64_abs(v65), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v109 = v83
										return v109
									} else {
										v104 = F_float_underflow_error_ext(m, int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return float64(0)
										} else {
											v109 = v104
											return v109
										}
									}
								}
							}
						}
					}
				} else {
					v45 = v19
					v46 = v13
					v47 = v18
					v49 = math.Float64frombits(uint64(0x7ff0000000000000))
					v51 = base.F64_sub(v45, v47)
					if base.F64_eq(base.F64_abs(v45), v49)|base.F64_ne(base.F64_abs(v51), v49)|base.F64_eq(base.F64_abs(v47), v49) == int32(0) {
						v63 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return float64(0)
						} else {
							v65 = v63
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v46)&int64(9223372036854775807)))|base.F64_ne(v65, float64(0)) == int32(0) {
								v77 = F_float_zero_divide_error_ext(m, int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return float64(0)
								} else {
									return v77
								}
							} else {
								v81 = math.Float64frombits(uint64(0x7ff0000000000000))
								v83 = base.F64_div(v46, v65)
								if base.F64_eq(base.F64_abs(v46), v81)|base.F64_ne(base.F64_abs(v83), v81) == int32(0) {
									v91 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return float64(0)
									} else {
										return v91
									}
								} else {
									v94 = float64(0)
									if base.F64_eq(v46, v94)|base.F64_ne(v83, v94)|base.F64_eq(base.F64_abs(v65), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v109 = v83
										return v109
									} else {
										v104 = F_float_underflow_error_ext(m, int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return float64(0)
										} else {
											v109 = v104
											return v109
										}
									}
								}
							}
						}
					} else {
						v65 = v51
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v46)&int64(9223372036854775807)))|base.F64_ne(v65, float64(0)) == int32(0) {
							v77 = F_float_zero_divide_error_ext(m, int32(0))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return float64(0)
							} else {
								return v77
							}
						} else {
							v81 = math.Float64frombits(uint64(0x7ff0000000000000))
							v83 = base.F64_div(v46, v65)
							if base.F64_eq(base.F64_abs(v46), v81)|base.F64_ne(base.F64_abs(v83), v81) == int32(0) {
								v91 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return float64(0)
								} else {
									return v91
								}
							} else {
								v94 = float64(0)
								if base.F64_eq(v46, v94)|base.F64_ne(v83, v94)|base.F64_eq(base.F64_abs(v65), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v109 = v83
									return v109
								} else {
									v104 = F_float_underflow_error_ext(m, int32(0))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return float64(0)
									} else {
										v109 = v104
										return v109
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
func F_point_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = v9 + int32(16)
	F_initStringInfo(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		F_appendStringInfoChar(m, v13, int32(40))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
			v23 = F_float8out_internal(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = F_float8out_internal(m, v21)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v23
					F_appendStringInfo(m, v13, int32(_a_F_point_out_0), v9)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_pfree(m, v23)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v25)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								F_appendStringInfoChar(m, v13, int32(41))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int64(0)
								} else {
									v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)))
									m.G0 = v9 + int32(32)
									return v39
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_point_sub(m *base.Module, l0 int32) int64 {
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
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_palloc(m, int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_sub(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19)|base.F64_eq(base.F64_abs(v16), v19) == int32(0) {
			v32 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
			} else {
				v34 = v32
				v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				v36 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
				v37 = base.F64_sub(v35, v36)
				v39 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39)|base.F64_eq(base.F64_abs(v36), v39) != 0 {
					v52 = v37
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
					return base.I64_extend_i32_u(v11)
				} else {
					v50 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						v52 = v50
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
						return base.I64_extend_i32_u(v11)
					}
				}
			}
		} else {
			v34 = v17
			v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v36 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v37 = base.F64_sub(v35, v36)
			v39 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39)|base.F64_eq(base.F64_abs(v36), v39) != 0 {
				v52 = v37
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
				return base.I64_extend_i32_u(v11)
			} else {
				v50 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					v52 = v50
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
					return base.I64_extend_i32_u(v11)
				}
			}
		}
	}
}
