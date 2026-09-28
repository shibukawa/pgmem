package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParseFuncOrColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v488 int32
	_ = v488
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
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
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v803 int32
	_ = v803
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v902 int32
	_ = v902
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1098 int32
	_ = v1098
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1177 int32
	_ = v1177
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1216 int32
	_ = v1216
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1341 int32
	_ = v1341
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1439 int32
	_ = v1439
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1544 int32
	_ = v1544
	var v1553 int32
	_ = v1553
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1623 int32
	_ = v1623
	var v1629 int32
	_ = v1629
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1648 int32
	_ = v1648
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1666 int32
	_ = v1666
	var v1676 int32
	_ = v1676
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
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
	var v1742 int32
	_ = v1742
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1822 int32
	_ = v1822
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1866 int32
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1890 int32
	_ = v1890
	var v1894 int32
	_ = v1894
	var v1900 int32
	_ = v1900
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1983 int32
	_ = v1983
	var v1987 int32
	_ = v1987
	var v1991 int32
	_ = v1991
	var v1996 int32
	_ = v1996
	var v2000 int32
	_ = v2000
	var v2004 int32
	_ = v2004
	var v2009 int32
	_ = v2009
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2031 int32
	_ = v2031
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2046 int32
	_ = v2046
	var v2048 int32
	_ = v2048
	var v2053 int32
	_ = v2053
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2067 int32
	_ = v2067
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2106 int32
	_ = v2106
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2147 int32
	_ = v2147
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2196 int32
	_ = v2196
	var v2200 int32
	_ = v2200
	var v2203 int32
	_ = v2203
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2214 int32
	_ = v2214
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2232 int32
	_ = v2232
	var v2236 int32
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2254 int32
	_ = v2254
	var v2258 int32
	_ = v2258
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2272 int32
	_ = v2272
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2297 int32
	_ = v2297
	var v2301 int32
	_ = v2301
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2321 int32
	_ = v2321
	var v2325 int32
	_ = v2325
	var v2327 int32
	_ = v2327
	var v2332 int32
	_ = v2332
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2348 int32
	_ = v2348
	var v2350 int32
	_ = v2350
	var v2355 int32
	_ = v2355
	v8 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(1088)
	m.G0 = v35
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l4)+32))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+31)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+30)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+29)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+28)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v45 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v56 = v8
	v57 = v8
	v58 = v8
	v59 = v8
	v60 = v8
	v61 = v8
	v62 = v8
	v63 = v8
	v64 = int32(0)
	goto L3
L3:
	;
	if l2 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v48 = F_transformWhereClause(m, l0, v45, int32(8), int32(_a_F_ParseFuncOrColumn_0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v53 = int32(0)
	goto L6
L6:
	;
	v56 = v44
	v57 = v42
	v58 = v41
	v59 = v43
	v60 = v40
	v61 = v38
	v62 = v39
	v63 = v37
	v64 = v53
	goto L3
L7:
	;
	return int32(0)
L8:
	;
	v53 = v48
	goto L6
L9:
	;
	v501 = int32(0)
	v503 = int32(1)
	if v477|(base.B2i32(l1 == v501)|(v62|(v60|(v58|(l5|base.B2i32(v479 != v503)|base.B2i32(v56 != v501)|base.B2i32(v64 != v501)))|base.B2i32(v59 != v501)))&v503) != 0 {
		v543 = v8
		goto L107
	} else {
		goto L108
	}
L10:
	;
	v197 = int32(0)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v197 < v198 {
		goto L35
	} else {
		goto L36
	}
L11:
	;
	v194 = int32(0)
	v470 = v194
	v477 = v194
	v479 = v172
	v488 = int32(1)
	v500 = v194
	goto L9
L12:
	;
	v172 = v8
	goto L11
L13:
	;
	goto L14
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v67 <= int32(100) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v156 != 0 {
		goto L10
	} else {
		goto L34
	}
L16:
	;
	v70 = int32(0)
	v76 = l2
	v82 = v70
	v85 = v8
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L29
	}
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v106 <= v82 {
		v156 = v76
		v158 = v85
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v156 = v76
	v158 = v132
	goto L15
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v82<<(uint(int32(2))%32))))
	v113 = F_exprType(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35+int32(656)+v85<<(uint(int32(2))%32)))) = v113
	v129 = int32(1)
	v132 = v85 + v129
	if v76 != 0 {
		v82 = v82 + v129
		v85 = v132
		goto L19
	} else {
		goto L28
	}
L23:
	;
	if v113 != int32(2278) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if base.B2i32(v117 != int32(8))|(base.B2i32(l4 == v70)|v57) != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v121 = F_list_delete_nth_cell(m, v76, v82)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	if v121 != 0 {
		v76 = v121
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v156 = v121
	v158 = v85
	goto L15
L28:
	;
	goto L20
L29:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v140 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+592)) = v140
	F_errmsg_plural(m, int32(_a_F_ParseFuncOrColumn_1), int32(_a_F_ParseFuncOrColumn_2), v140, v35+int32(592))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(146), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v172 = v158
	goto L11
L35:
	;
	v213 = v8
	v218 = v8
	goto L38
L36:
	;
	v443 = v197
	goto L37
L37:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	v470 = v156
	v477 = v443
	v479 = v158
	v488 = v8
	v500 = v467
	goto L9
L38:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v233+v218<<(uint(int32(2))%32))))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	if v238 == int32(16) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v443 = v429
	goto L37
L40:
	;
	v431 = v218 + int32(1)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v431 < v432 {
		v213 = v429
		v218 = v431
		goto L38
	} else {
		goto L75
	}
L41:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v237)+8))
	v395 = F_lappend(m, v213, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L7
	} else {
		goto L74
	}
L42:
	;
	if v213 == int32(0) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v340 = int32(0)
	if v213 == v340 {
		v429 = v340
		goto L40
	} else {
		goto L68
	}
L45:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v243 <= int32(0) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v246 = int32(0)
	if v246 < v243 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v249 = v243
	goto L49
L48:
	;
	v249 = v246
	goto L49
L49:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v237)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	v261 = int32(0)
	goto L50
L50:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v251+v261<<(uint(int32(2))%32))))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if base.B2i32(v291 == int32(0))|base.B2i32(v291 != v294) != 0 {
		v312 = v291
		v313 = v294
		goto L53
	} else {
		goto L54
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L7
	} else {
		goto L63
	}
L52:
	;
	if v312-v313 != 0 {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	goto L52
L54:
	;
	v297 = v250
	v298 = v288
	goto L55
L55:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v302 == int32(0) {
		v312 = v302
		v313 = v301
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v312 = v302
	v313 = v301
	goto L53
L57:
	;
	v305 = int32(1)
	if v302 == v301 {
		v297 = v297 + v305
		v298 = v298 + v305
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v316 = v261 + int32(1)
	if v249 != v316 {
		v261 = v316
		goto L50
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L51
L62:
	;
	goto L41
L63:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v237)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+608)) = v325
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_5), v35+int32(608))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v237)+16))
	F_parser_errposition(m, l0, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(200), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_6), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	v354 = F_exprLocation(m, v237)
	mBase = m.M
	F_parser_errposition(m, l0, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(210), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v429 = v395
	goto L40
L75:
	;
	goto L39
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L7
	} else {
		goto L566
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2318 = m.ExcPending
	if v2318 != 0 {
		goto L7
	} else {
		goto L561
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2294 = m.ExcPending
	if v2294 != 0 {
		goto L7
	} else {
		goto L555
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L7
	} else {
		goto L550
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L7
	} else {
		goto L545
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L7
	} else {
		goto L539
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L7
	} else {
		goto L534
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L7
	} else {
		goto L529
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L7
	} else {
		goto L523
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L7
	} else {
		goto L518
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L7
	} else {
		goto L512
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+224)) = v1054
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_7), v35+int32(224))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L7
	} else {
		goto L508
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L7
	} else {
		goto L500
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+192)) = v1007
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_8), v35+int32(192))
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L7
	} else {
		goto L495
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L7
	} else {
		goto L489
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L7
	} else {
		goto L483
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L7
	} else {
		goto L480
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L7
	} else {
		goto L477
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L7
	} else {
		goto L469
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L7
	} else {
		goto L463
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L7
	} else {
		goto L457
	}
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L7
	} else {
		goto L454
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L7
	} else {
		goto L448
	}
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L7
	} else {
		goto L442
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L7
	} else {
		goto L436
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L7
	} else {
		goto L430
	}
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L7
	} else {
		goto L424
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L7
	} else {
		goto L418
	}
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L7
	} else {
		goto L411
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L7
	} else {
		goto L404
	}
L106:
	;
	m.G0 = v35 + int32(1088)
	return v1676
L107:
	;
	v545 = v35 + int32(612)
	*(*int32)(unsafe.Add(mBase, uint32(v545)+12)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v545)+4)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v545)+16)) = v545
	v551 = int32(_a_F_ParseFuncOrColumn_9)
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_ParseFuncOrColumn[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v545)+8)) = v552
	*(*int32)(unsafe.Add(mBase, _c_F_ParseFuncOrColumn[0])) = v35 + int32(620)
	goto L120
L108:
	;
	v522 = int32(1)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v523 != v522 {
		v543 = v8
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v35)+656))
	if v526 != int32(2249) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v529 = F_typeOrDomainTypeRelid(m, v526)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L7
	} else {
		goto L113
	}
L111:
	;
	v533 = v522
	goto L112
L112:
	;
	if l4 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v533 = base.B2i32(v529 != int32(0))
	goto L112
L114:
	;
	v543 = v533
	goto L107
L115:
	;
	goto L116
L116:
	;
	if v533 == int32(0) {
		v543 = v533
		goto L107
	} else {
		goto L117
	}
L117:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v539 = F_ParseComplexProjection(m, l0, v538, v500, l6)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	if v539 != 0 {
		v1676 = v539
		goto L106
	} else {
		goto L119
	}
L119:
	;
	v543 = int32(1)
	goto L107
L120:
	;
	v560 = int32(1)
	v579 = F_func_get_detail(m, l1, v470, v477, v479, v35+int32(656), v62^v560, v560, l5, v35+int32(632), v35+int32(1064), v35+int32(1068), v35+int32(647), v35+int32(640), v35+int32(636), v35+int32(652), v35+int32(648))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v545)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseFuncOrColumn[0])) = v582
	goto L122
L122:
	;
	if l5 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	if base.B2i32(v579 != int32(6))&base.B2i32(v623 != int32(2)) == int32(0) {
		goto L139
	} else {
		goto L140
	}
L124:
	;
	if v579&int32(3) != int32(2) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	if v579 == int32(3) {
		goto L105
	} else {
		goto L138
	}
L127:
	;
	v589 = v579 & int32(6)
	if v589 != int32(4) {
		v623 = v589
		goto L123
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L7
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	v602 = F_func_signature_string(m, l1, v479, v477, v35+int32(656))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v602
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_10), v35)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_11), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L7
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(297), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L7
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
	v623 = v579 & int32(6)
	goto L123
L139:
	;
	if v58 != 0 {
		goto L104
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if v623 == int32(2) {
		v1098 = v8
		goto L149
	} else {
		goto L150
	}
L142:
	;
	if v60 != 0 {
		goto L103
	} else {
		goto L143
	}
L143:
	;
	if v57 != 0 {
		goto L102
	} else {
		goto L144
	}
L144:
	;
	if v56 != 0 {
		goto L101
	} else {
		goto L145
	}
L145:
	;
	if v64 != 0 {
		goto L100
	} else {
		goto L146
	}
L146:
	;
	if v59 != 0 {
		goto L99
	} else {
		goto L147
	}
L147:
	;
	if v61 != 0 {
		goto L98
	} else {
		goto L148
	}
L148:
	;
	goto L141
L149:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v35)+648))
	if v1105 == int32(0) {
		v1177 = v479
		goto L281
	} else {
		goto L282
	}
L150:
	;
	switch v579 - int32(4) {
	case 0:
		goto L153
	case 1:
		goto L152
	default:
		goto L151
	}
L151:
	;
	if v579 == int32(6) {
		goto L242
	} else {
		goto L243
	}
L152:
	;
	if v59 == int32(0) {
		goto L90
	} else {
		goto L234
	}
L153:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v35)+1064))
	v638 = F_SearchSysCache1(m, int32(0), base.I64_extend_i32_u(v636))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L7
	} else {
		goto L154
	}
L154:
	;
	if v638 == int32(0) {
		goto L97
	} else {
		goto L155
	}
L155:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v638)+16))
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+22)))
	v644 = v642 + v643
	v645 = int32(*(*int16)(unsafe.Add(mBase, uint32(v644)+6)))
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644)+4)))
	F_ReleaseCatCache(m, v638)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L7
	} else {
		goto L156
	}
L156:
	;
	if v646 != int32(110) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	if v61 == int32(0) {
		v1098 = v646
		goto L149
	} else {
		goto L228
	}
L158:
	;
	if v57 == int32(0) {
		goto L96
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	if v57 != 0 {
		goto L91
	} else {
		goto L227
	}
L161:
	;
	if v59 != 0 {
		goto L95
	} else {
		goto L162
	}
L162:
	;
	if v56 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v655 = v653
	goto L165
L164:
	;
	v655 = int32(0)
	goto L165
L165:
	;
	v656 = v479 - v655
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v35)+636))
	if v657 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	if v488 != 0 {
		goto L209
	} else {
		goto L210
	}
L167:
	;
	if v646 != int32(104) {
		goto L157
	} else {
		goto L208
	}
L168:
	;
	if v645 == v656 {
		goto L167
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v696 = int32(1)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v35)+640))
	if v697 <= v696 {
		goto L180
	} else {
		goto L181
	}
L171:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L7
	} else {
		goto L172
	}
L172:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L7
	} else {
		goto L173
	}
L173:
	;
	v670 = F_func_signature_string(m, l1, v479, v477, v35+int32(656))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L7
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+336)) = v670
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v35+int32(336))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L7
	} else {
		goto L175
	}
L175:
	;
	v678 = F_NameListToString(m, l1)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L7
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+328)) = v656
	*(*int32)(unsafe.Add(mBase, uint32(v35)+324)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v35)+320)) = v678
	F_errhint_plural(m, int32(_a_F_ParseFuncOrColumn_13), int32(_a_F_ParseFuncOrColumn_14), v645, v35+int32(320))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L7
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(438), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L7
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	v700 = v696
	goto L182
L181:
	;
	v700 = v697
	goto L182
L182:
	;
	if v645 < v479-v700+int32(1) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	if v645 == v656 {
		goto L167
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	if v646 == int32(104) {
		goto L195
	} else {
		goto L196
	}
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L7
	} else {
		goto L187
	}
L187:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L7
	} else {
		goto L188
	}
L188:
	;
	v715 = F_func_signature_string(m, l1, v479, v477, v35+int32(656))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L7
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+368)) = v715
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v35+int32(368))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L7
	} else {
		goto L190
	}
L190:
	;
	v723 = F_NameListToString(m, l1)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L7
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+360)) = v656
	*(*int32)(unsafe.Add(mBase, uint32(v35)+356)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v35)+352)) = v723
	F_errhint_plural(m, int32(_a_F_ParseFuncOrColumn_13), int32(_a_F_ParseFuncOrColumn_14), v645, v35+int32(352))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L7
	} else {
		goto L192
	}
L192:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L7
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(469), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L7
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	if v697 == v655<<(uint(int32(1))%32) {
		goto L166
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	if v697 <= v655 {
		goto L94
	} else {
		goto L207
	}
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L7
	} else {
		goto L199
	}
L199:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L7
	} else {
		goto L200
	}
L200:
	;
	v755 = F_func_signature_string(m, l1, v479, v477, v35+int32(656))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L7
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+400)) = v755
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v35+int32(400))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L7
	} else {
		goto L202
	}
L202:
	;
	v763 = F_NameListToString(m, l1)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L7
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+384)) = v763
	*(*int32)(unsafe.Add(mBase, uint32(v35)+392)) = v655
	*(*int32)(unsafe.Add(mBase, uint32(v35)+388)) = v697 - v655
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_15), v35+int32(384))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L7
	} else {
		goto L204
	}
L204:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L7
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(494), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L7
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	goto L167
L208:
	;
	goto L166
L209:
	;
	v788 = int32(0)
	goto L211
L210:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	v788 = v787
	goto L211
L211:
	;
	v789 = v788 - v655
	v790 = v789 - v655
	if v790 < int32(0) {
		goto L93
	} else {
		goto L212
	}
L212:
	;
	if v789 <= v790 {
		goto L157
	} else {
		goto L213
	}
L213:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v35)+652))
	v803 = v790
	goto L214
L214:
	;
	v827 = int32(2)
	v828 = v803 << (uint(v827) % 32)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v794+v828)))
	v834 = (v803 - v790 + v789) << (uint(v827) % 32)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v794+v834)))
	if v830 != v836 {
		goto L92
	} else {
		goto L216
	}
L215:
	;
	goto L157
L216:
	;
	if v830 == int32(2276) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v841 = v840 + v834
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v841)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+1084)) = v842
	v844 = v840 + v828
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+1080)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v35)+316)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v35)+312)) = v845
	v853 = F_list_make2_impl(m, v35+int32(316), v35+int32(312))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L7
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v902 = v803 + int32(1)
	if v902 != v789 {
		v803 = v902
		goto L214
	} else {
		goto L226
	}
L220:
	;
	v857 = F_select_common_type(m, l0, v853, int32(_a_F_ParseFuncOrColumn_16), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L7
	} else {
		goto L221
	}
L221:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v841)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+1076)) = v859
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v844)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+1072)) = v861
	*(*int32)(unsafe.Add(mBase, uint32(v35)+308)) = v859
	*(*int32)(unsafe.Add(mBase, uint32(v35)+304)) = v861
	v869 = F_list_make2_impl(m, v35+int32(308), v35+int32(304))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L7
	} else {
		goto L222
	}
L222:
	;
	v871 = F_select_common_typmod(m, v869, v857)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L7
	} else {
		goto L223
	}
L223:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v844)))
	v875 = v35 + int32(656)
	v876 = v828 + v875
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v876)))
	v881 = F_coerce_type(m, l0, v873, v877, v857, v871, int32(0), int32(2), int32(-1))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L7
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844))) = v881
	*(*int32)(unsafe.Add(mBase, uint32(v876))) = v857
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v841)))
	v886 = v875 + v834
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)))
	v891 = F_coerce_type(m, l0, v885, v887, v857, v871, int32(0), int32(2), int32(-1))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L7
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v841))) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v886))) = v857
	goto L219
L226:
	;
	goto L215
L227:
	;
	goto L157
L228:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L7
	} else {
		goto L229
	}
L229:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L7
	} else {
		goto L230
	}
L230:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_17), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L7
	} else {
		goto L231
	}
L231:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L7
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(535), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L7
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
	if v57 == int32(0) {
		v1098 = v8
		goto L149
	} else {
		goto L235
	}
L235:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L7
	} else {
		goto L236
	}
L236:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L7
	} else {
		goto L237
	}
L237:
	;
	v967 = F_NameListToString(m, l1)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L7
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+496)) = v967
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_18), v35+int32(496))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L7
	} else {
		goto L239
	}
L239:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L7
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(554), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L7
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v35)+656))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v35)+1068))
	v991 = F_coerce_type(m, l0, v985, v986, v987, int32(-1), int32(3), int32(0), l6)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L7
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	if v579 == int32(1) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v1676 = v991
	goto L106
L246:
	;
	if l4 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	goto L248
L248:
	;
	if l4 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L249:
	;
	v1676 = int32(0)
	goto L106
L250:
	;
	goto L251
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L7
	} else {
		goto L252
	}
L252:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L7
	} else {
		goto L253
	}
L253:
	;
	v1007 = F_func_signature_string(m, l1, v479, v477, v35+int32(656))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L7
	} else {
		goto L254
	}
L254:
	;
	if l5 != 0 {
		goto L89
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+208)) = v1007
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_19), v35+int32(208))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L7
	} else {
		goto L256
	}
L256:
	;
	v1017 = F_errdetail(m, int32(_a_F_ParseFuncOrColumn_20), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L7
	} else {
		goto L257
	}
L257:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_21), int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L7
	} else {
		goto L258
	}
L258:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L7
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(595), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L7
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	v1676 = int32(0)
	goto L106
L262:
	;
	goto L263
L263:
	;
	if v543 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1033)))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+4))
	v1036 = F_ParseComplexProjection(m, l0, v1035, v500, l6)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L7
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	if v56 != 0 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	if v1036 != 0 {
		v1676 = v1036
		goto L106
	} else {
		goto L268
	}
L268:
	;
	goto L266
L269:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v57|base.B2i32(v1039 < int32(2)) == int32(0) {
		goto L88
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L7
	} else {
		goto L273
	}
L272:
	;
	goto L271
L273:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L7
	} else {
		goto L274
	}
L274:
	;
	v1054 = F_func_signature_string(m, l1, v479, v477, v35+int32(656))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L7
	} else {
		goto L275
	}
L275:
	;
	if l5 != 0 {
		goto L87
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v1054
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v35+int32(240))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L7
	} else {
		goto L277
	}
L277:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v35)+632))
	F_func_lookup_failure_details(m, v1062, v477, int32(0))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L7
	} else {
		goto L278
	}
L278:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L7
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(656), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L7
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	v1201 = int32(0)
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v35)+652))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v35)+1068))
	v1207 = F_enforce_generic_type_consistency(m, v35+int32(656), v1204, v1177, v1205, v1201)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L7
	} else {
		goto L292
	}
L282:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+4))
	if v1108 <= int32(0) {
		v1177 = v479
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1111 = int32(100)
	if v479 <= v1111 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1114 = v1111
	goto L286
L285:
	;
	v1114 = v479
	goto L286
L286:
	;
	v1121 = int32(0)
	v1125 = v479
	goto L287
L287:
	;
	if v1121 == v1114-v479 {
		goto L76
	} else {
		goto L289
	}
L288:
	;
	v1177 = v1164
	goto L281
L289:
	;
	v1152 = int32(2)
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+12))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1155+v1121<<(uint(v1152)%32))))
	v1160 = F_exprType(m, v1159)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L7
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35+int32(656)+v1125<<(uint(v1152)%32)))) = v1160
	v1163 = int32(1)
	v1164 = v1125 + v1163
	v1166 = v1121 + v1163
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+4))
	if v1166 < v1167 {
		v1121 = v1166
		v1125 = v1164
		goto L287
	} else {
		goto L291
	}
L291:
	;
	goto L288
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+1068)) = v1207
	if int32(0) < v1177 {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	if v488 != 0 {
		goto L313
	} else {
		goto L314
	}
L294:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L7
	} else {
		goto L308
	}
L295:
	;
	v1216 = v1201
	goto L298
L296:
	;
	goto L297
L297:
	;
	if v1207 != int32(2281) {
		goto L293
	} else {
		goto L302
	}
L298:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1204+v1216<<(uint(int32(2))%32))))
	if v1247 == int32(2281) {
		goto L294
	} else {
		goto L300
	}
L299:
	;
	goto L297
L300:
	;
	v1251 = v1216 + int32(1)
	if v1251 != v1177 {
		v1216 = v1251
		goto L298
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L7
	} else {
		goto L303
	}
L303:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L7
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+80)) = int32(_a_F_ParseFuncOrColumn_22)
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_23), v35+int32(80))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L7
	} else {
		goto L305
	}
L305:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L7
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(719), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L7
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L7
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+64)) = int32(_a_F_ParseFuncOrColumn_22)
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_24), v35-int32(-64))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L7
	} else {
		goto L310
	}
L310:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L7
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(712), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L7
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v35)+636))
	v1434 = int32(0)
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v35)+640))
	if base.B2i32(v1433 == int32(2276))|base.B2i32(v1439 <= v1434) == v1434 {
		goto L326
	} else {
		goto L327
	}
L314:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v1329 <= int32(0) {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v1341 = int32(0)
	goto L316
L316:
	;
	v1366 = v1341 << (uint(int32(2)) % 32)
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1366+(v35+int32(656)))))
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1366+v1204)))
	if v1370 == v1372 {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	goto L313
L318:
	;
	v1398 = v1341 + int32(1)
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v1398 < v1399 {
		v1341 = v1398
		goto L316
	} else {
		goto L325
	}
L319:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v1375 = v1374 + v1366
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1375)))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1376)))
	if v1377 == int32(16) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v1381 = int32(-1)
	v1385 = F_coerce_type(m, l0, v1380, v1370, v1372, v1381, int32(0), int32(2), v1381)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L7
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	v1388 = int32(-1)
	v1392 = F_coerce_type(m, l0, v1376, v1370, v1372, v1388, int32(0), int32(2), v1388)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L7
	} else {
		goto L324
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+4)) = v1385
	goto L318
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1375))) = v1392
	goto L318
L325:
	;
	goto L317
L326:
	;
	v1446 = F_palloc0(m, int32(36))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L7
	} else {
		goto L329
	}
L327:
	;
	v1483 = v470
	v1484 = v62 & base.B2i32(v1433 != v1434)
	v1486 = v1433
	goto L328
L328:
	;
	v1487 = int32(0)
	if base.B2i32(v1484 == v1487)|base.B2i32(v479 <= v1487)|base.B2i32(v1486 != int32(2276)) == v1487 {
		goto L342
	} else {
		goto L343
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1446))) = int32(35)
	v1450 = v479 - v1439
	v1451 = F_list_copy_tail(m, v470, v1450)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L7
	} else {
		goto L330
	}
L330:
	;
	v1453 = int32(0)
	if base.B2i32(v470 == v1453)|base.B2i32(v1450 <= v1453) != 0 {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1446)+16)) = v1451
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+12))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1465)))
	v1467 = F_exprType(m, v1466)
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L7
	} else {
		goto L338
	}
L332:
	;
	v1463 = int32(0)
	goto L334
L333:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v1450 < v1460 {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	goto L331
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470)+4)) = v1450
	goto L337
L336:
	;
	goto L337
L337:
	;
	v1463 = v470
	goto L334
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1446)+12)) = v1467
	v1470 = F_get_array_type(m, v1467)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L7
	} else {
		goto L339
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1446)+4)) = v1470
	if v1470 == int32(0) {
		goto L86
	} else {
		goto L340
	}
L340:
	;
	v1475 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1446)+20)) = uint8(v1475)
	v1477 = F_exprLocation(m, v1451)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1446)+32)) = v1477
	v1480 = F_lappend(m, v1463, v1446)
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L7
	} else {
		goto L341
	}
L341:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v35)+636))
	v1483 = v1480
	v1484 = int32(1)
	v1486 = v1482
	goto L328
L342:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v479<<(uint(int32(2))%32)+v35)+652))
	v1501 = F_get_base_element_type(m, v1500)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L7
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+647)))
	if v1505 == int32(1) {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	if v1501 == int32(0) {
		goto L85
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	F_check_srf_call_placement(m, l0, l3, l6)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L7
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	if v623 == int32(2) {
		goto L352
	} else {
		goto L353
	}
L350:
	;
	goto L349
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v1666
	v1676 = v1666
	goto L106
L352:
	;
	v1513 = F_palloc0(m, int32(36))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L7
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	if base.B2i32(v579 != int32(4))|v59 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1513))) = int32(15)
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v35)+1064))
	*(*int32)(unsafe.Add(mBase, uint32(v1513)+4)) = v1517
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v35)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v1513)+32)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v1513)+28)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v1513)+16)) = v63
	*(*uint8)(unsafe.Add(mBase, uint32(v1513)+13)) = uint8(v1484)
	*(*uint8)(unsafe.Add(mBase, uint32(v1513)+12)) = uint8(v1505)
	*(*int32)(unsafe.Add(mBase, uint32(v1513)+8)) = v1519
	if v1505 != 0 {
		v1666 = v1513
		goto L351
	} else {
		goto L356
	}
L356:
	;
	v1676 = v1513
	goto L106
L357:
	;
	F_transformAggregateCall(m, l0, v1532, v1483, v56, v60)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L7
	} else {
		goto L402
	}
L358:
	;
	v1532 = F_palloc0(m, int32(72))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L7
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1603 = F_palloc0(m, int32(48))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L7
	} else {
		goto L380
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1532))) = int32(9)
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v35)+1064))
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+4)) = v1536
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v35)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+68)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+64)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1532)+56)) = int64(-4294967296)
	v1544 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1532)+51)) = uint8(v1544)
	*(*uint8)(unsafe.Add(mBase, uint32(v1532)+50)) = uint8(v1098)
	*(*uint8)(unsafe.Add(mBase, uint32(v1532)+49)) = uint8(v1484)
	*(*uint8)(unsafe.Add(mBase, uint32(v1532)+48)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+44)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+20)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+8)) = v1538
	if v1483 != 0 {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	if v1505 != 0 {
		goto L83
	} else {
		goto L373
	}
L363:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+4))
	if v1553 < int32(100) {
		goto L362
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	if v57|v58 == int32(0) {
		goto L84
	} else {
		goto L372
	}
L366:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L7
	} else {
		goto L367
	}
L367:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L7
	} else {
		goto L368
	}
L368:
	;
	v1563 = int32(99)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+160)) = v1563
	F_errmsg_plural(m, int32(_a_F_ParseFuncOrColumn_25), int32(_a_F_ParseFuncOrColumn_26), v1563, v35+int32(160))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L7
	} else {
		goto L369
	}
L369:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L7
	} else {
		goto L370
	}
L370:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(844), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L7
	} else {
		goto L371
	}
L371:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L372:
	;
	goto L362
L373:
	;
	if v477 == int32(0) {
		goto L357
	} else {
		goto L374
	}
L374:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L7
	} else {
		goto L375
	}
L375:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L7
	} else {
		goto L376
	}
L376:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_27), int32(0))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L7
	} else {
		goto L377
	}
L377:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L7
	} else {
		goto L378
	}
L378:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(876), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L7
	} else {
		goto L379
	}
L379:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1603))) = int32(11)
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v35)+1064))
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+4)) = v1607
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v35)+1068))
	*(*uint8)(unsafe.Add(mBase, uint32(v1603)+37)) = uint8(base.B2i32(v579 == int32(4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1603)+36)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+20)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+8)) = v1609
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+40)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+24)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+44)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+28)) = int32(0)
	if v60 != 0 {
		goto L82
	} else {
		goto L381
	}
L381:
	;
	if v579 != int32(4) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	if v56 != 0 {
		goto L80
	} else {
		goto L394
	}
L383:
	;
	if v1483 != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+4))
	if v1623 < int32(100) {
		goto L382
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	if v58 == int32(0) {
		goto L81
	} else {
		goto L393
	}
L387:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L7
	} else {
		goto L388
	}
L388:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L7
	} else {
		goto L389
	}
L389:
	;
	v1633 = int32(99)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+128)) = v1633
	F_errmsg_plural(m, int32(_a_F_ParseFuncOrColumn_25), int32(_a_F_ParseFuncOrColumn_26), v1633, v35+int32(128))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L7
	} else {
		goto L390
	}
L390:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L7
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(923), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L7
	} else {
		goto L392
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	goto L382
L394:
	;
	if v64 != 0 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1654 = base.B2i32(v579 != int32(4))
	goto L397
L396:
	;
	v1654 = int32(0)
	goto L397
L397:
	;
	if v1654 != 0 {
		goto L79
	} else {
		goto L398
	}
L398:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v1655 != l3 {
		goto L78
	} else {
		goto L399
	}
L399:
	;
	if v1505 != 0 {
		goto L77
	} else {
		goto L400
	}
L400:
	;
	F_transformWindowFuncCall(m, l0, v1603, v59)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L7
	} else {
		goto L401
	}
L401:
	;
	v1676 = v1603
	goto L106
L402:
	;
	v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+647)))
	if v1661 != int32(1) {
		v1676 = v1532
		goto L106
	} else {
		goto L403
	}
L403:
	;
	v1666 = v1532
	goto L351
L404:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L7
	} else {
		goto L405
	}
L405:
	;
	v1713 = F_func_signature_string(m, l1, v479, v477, v35+int32(656))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L7
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+576)) = v1713
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_28), v35+int32(576))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L7
	} else {
		goto L407
	}
L407:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_29), int32(0))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L7
	} else {
		goto L408
	}
L408:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L7
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(307), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L7
	} else {
		goto L410
	}
L410:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L411:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L7
	} else {
		goto L412
	}
L412:
	;
	v1739 = F_NameListToString(m, l1)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L7
	} else {
		goto L413
	}
L413:
	;
	v1741 = F_NameListToString(m, l1)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L7
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v1739
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_30), v35+int32(16))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L7
	} else {
		goto L415
	}
L415:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L7
	} else {
		goto L416
	}
L416:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(323), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		goto L7
	} else {
		goto L417
	}
L417:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L418:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L7
	} else {
		goto L419
	}
L419:
	;
	v1764 = F_NameListToString(m, l1)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L7
	} else {
		goto L420
	}
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v1764
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_31), v35+int32(32))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L7
	} else {
		goto L421
	}
L421:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L7
	} else {
		goto L422
	}
L422:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(329), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L7
	} else {
		goto L423
	}
L423:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L424:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L7
	} else {
		goto L425
	}
L425:
	;
	v1786 = F_NameListToString(m, l1)
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L7
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v1786
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_32), v35+int32(48))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L7
	} else {
		goto L427
	}
L427:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L7
	} else {
		goto L428
	}
L428:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(335), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L7
	} else {
		goto L429
	}
L429:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L430:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L7
	} else {
		goto L431
	}
L431:
	;
	v1808 = F_NameListToString(m, l1)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L7
	} else {
		goto L432
	}
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+560)) = v1808
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_33), v35+int32(560))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L7
	} else {
		goto L433
	}
L433:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L7
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(341), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L7
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
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L7
	} else {
		goto L437
	}
L437:
	;
	v1830 = F_NameListToString(m, l1)
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L7
	} else {
		goto L438
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+544)) = v1830
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_34), v35+int32(544))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L7
	} else {
		goto L439
	}
L439:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L7
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(347), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L7
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
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L7
	} else {
		goto L443
	}
L443:
	;
	v1852 = F_NameListToString(m, l1)
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L7
	} else {
		goto L444
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+528)) = v1852
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_35), v35+int32(528))
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L7
	} else {
		goto L445
	}
L445:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L7
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(353), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L7
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
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L7
	} else {
		goto L449
	}
L449:
	;
	v1874 = F_NameListToString(m, l1)
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L7
	} else {
		goto L450
	}
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+516)) = v1874
	*(*int32)(unsafe.Add(mBase, uint32(v35)+512)) = int32(_a_F_ParseFuncOrColumn_36)
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_37), v35+int32(512))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L7
	} else {
		goto L451
	}
L451:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L7
	} else {
		goto L452
	}
L452:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(360), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L7
	} else {
		goto L453
	}
L453:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+272)) = v636
	F_errmsg_internal(m, int32(_a_F_ParseFuncOrColumn_38), v35+int32(272))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L7
	} else {
		goto L455
	}
L455:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(381), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L7
	} else {
		goto L456
	}
L456:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L457:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L7
	} else {
		goto L458
	}
L458:
	;
	v1913 = F_NameListToString(m, l1)
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L7
	} else {
		goto L459
	}
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+464)) = v1913
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_39), v35+int32(464))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L7
	} else {
		goto L460
	}
L460:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L7
	} else {
		goto L461
	}
L461:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(398), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L7
	} else {
		goto L462
	}
L462:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L463:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L7
	} else {
		goto L464
	}
L464:
	;
	v1935 = F_NameListToString(m, l1)
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L7
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+448)) = v1935
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_40), v35+int32(448))
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L7
	} else {
		goto L466
	}
L466:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L7
	} else {
		goto L467
	}
L467:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(404), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L7
	} else {
		goto L468
	}
L468:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L469:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L7
	} else {
		goto L470
	}
L470:
	;
	v1959 = F_func_signature_string(m, l1, v479, v477, v35+int32(656))
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L7
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+432)) = v1959
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v35+int32(432))
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L7
	} else {
		goto L472
	}
L472:
	;
	v1967 = F_NameListToString(m, l1)
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L7
	} else {
		goto L473
	}
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+420)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v35)+416)) = v1967
	F_errhint_plural(m, int32(_a_F_ParseFuncOrColumn_41), int32(_a_F_ParseFuncOrColumn_42), v645, v35+int32(416))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L7
	} else {
		goto L474
	}
L474:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L7
	} else {
		goto L475
	}
L475:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(510), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L7
	} else {
		goto L476
	}
L476:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L477:
	;
	F_errmsg_internal(m, int32(_a_F_ParseFuncOrColumn_43), int32(0))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L7
	} else {
		goto L478
	}
L478:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(1942), int32(_a_F_ParseFuncOrColumn_44))
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L7
	} else {
		goto L479
	}
L479:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L480:
	;
	F_errmsg_internal(m, int32(_a_F_ParseFuncOrColumn_45), int32(0))
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L7
	} else {
		goto L481
	}
L481:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(1955), int32(_a_F_ParseFuncOrColumn_44))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L7
	} else {
		goto L482
	}
L482:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L483:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L7
	} else {
		goto L484
	}
L484:
	;
	v2017 = F_NameListToString(m, l1)
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L7
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+288)) = v2017
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_46), v35+int32(288))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L7
	} else {
		goto L486
	}
L486:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L7
	} else {
		goto L487
	}
L487:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(528), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L7
	} else {
		goto L488
	}
L488:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L489:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L7
	} else {
		goto L490
	}
L490:
	;
	v2039 = F_NameListToString(m, l1)
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L7
	} else {
		goto L491
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+480)) = v2039
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_47), v35+int32(480))
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L7
	} else {
		goto L492
	}
L492:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L7
	} else {
		goto L493
	}
L493:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(547), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L7
	} else {
		goto L494
	}
L494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L495:
	;
	v2062 = F_errdetail(m, int32(_a_F_ParseFuncOrColumn_48), int32(0))
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L7
	} else {
		goto L496
	}
L496:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_21), int32(0))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L7
	} else {
		goto L497
	}
L497:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L7
	} else {
		goto L498
	}
L498:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(586), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L7
	} else {
		goto L499
	}
L499:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L500:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L7
	} else {
		goto L501
	}
L501:
	;
	v2084 = F_func_signature_string(m, l1, v479, v477, v35+int32(656))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L7
	} else {
		goto L502
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+256)) = v2084
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v35+int32(256))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L7
	} else {
		goto L503
	}
L503:
	;
	v2094 = F_errdetail(m, int32(_a_F_ParseFuncOrColumn_49), int32(0))
	mBase = m.M
	v2095 = m.ExcPending
	if v2095 != 0 {
		goto L7
	} else {
		goto L504
	}
L504:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_50), int32(0))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L7
	} else {
		goto L505
	}
L505:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L7
	} else {
		goto L506
	}
L506:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(637), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2106 = m.ExcPending
	if v2106 != 0 {
		goto L7
	} else {
		goto L507
	}
L507:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L508:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v35)+632))
	F_func_lookup_failure_details(m, v2113, v477, int32(1))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L7
	} else {
		goto L509
	}
L509:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L7
	} else {
		goto L510
	}
L510:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(647), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L7
	} else {
		goto L511
	}
L511:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L512:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L7
	} else {
		goto L513
	}
L513:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+12))
	v2132 = F_format_type_be(m, v2131)
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L7
	} else {
		goto L514
	}
L514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+96)) = v2132
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_51), v35+int32(96))
	mBase = m.M
	v2139 = m.ExcPending
	if v2139 != 0 {
		goto L7
	} else {
		goto L515
	}
L515:
	;
	v2140 = F_exprLocation(m, v1451)
	mBase = m.M
	F_parser_errposition(m, l0, v2140)
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L7
	} else {
		goto L516
	}
L516:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(758), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L7
	} else {
		goto L517
	}
L517:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L518:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L7
	} else {
		goto L519
	}
L519:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_52), int32(0))
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L7
	} else {
		goto L520
	}
L520:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+12))
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+4))
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2159+v2160<<(uint(int32(2))%32)-int32(4))))
	v2167 = F_exprLocation(m, v2166)
	mBase = m.M
	F_parser_errposition(m, l0, v2167)
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L7
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(784), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L7
	} else {
		goto L522
	}
L522:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L523:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L7
	} else {
		goto L524
	}
L524:
	;
	v2182 = F_NameListToString(m, l1)
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L7
	} else {
		goto L525
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+144)) = v2182
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_53), v35+int32(144))
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L7
	} else {
		goto L526
	}
L526:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2191 = m.ExcPending
	if v2191 != 0 {
		goto L7
	} else {
		goto L527
	}
L527:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(855), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L7
	} else {
		goto L528
	}
L528:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L529:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L7
	} else {
		goto L530
	}
L530:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_54), int32(0))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L7
	} else {
		goto L531
	}
L531:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L7
	} else {
		goto L532
	}
L532:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(861), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L7
	} else {
		goto L533
	}
L533:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L534:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L7
	} else {
		goto L535
	}
L535:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_55), int32(0))
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L7
	} else {
		goto L536
	}
L536:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L7
	} else {
		goto L537
	}
L537:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(910), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		goto L7
	} else {
		goto L538
	}
L538:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L539:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L7
	} else {
		goto L540
	}
L540:
	;
	v2240 = F_NameListToString(m, l1)
	mBase = m.M
	v2241 = m.ExcPending
	if v2241 != 0 {
		goto L7
	} else {
		goto L541
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+112)) = v2240
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_53), v35+int32(112))
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L7
	} else {
		goto L542
	}
L542:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L7
	} else {
		goto L543
	}
L543:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(934), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L7
	} else {
		goto L544
	}
L544:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L545:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L7
	} else {
		goto L546
	}
L546:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_56), int32(0))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L7
	} else {
		goto L547
	}
L547:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
		goto L7
	} else {
		goto L548
	}
L548:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(943), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L7
	} else {
		goto L549
	}
L549:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L550:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L7
	} else {
		goto L551
	}
L551:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_57), int32(0))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L7
	} else {
		goto L552
	}
L552:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L7
	} else {
		goto L553
	}
L553:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(952), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L7
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L7
	} else {
		goto L556
	}
L556:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_58), int32(0))
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L7
	} else {
		goto L557
	}
L557:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_59), int32(0))
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L7
	} else {
		goto L558
	}
L558:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v2307 = F_exprLocation(m, v2306)
	mBase = m.M
	F_parser_errposition(m, l0, v2307)
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L7
	} else {
		goto L559
	}
L559:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(963), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L7
	} else {
		goto L560
	}
L560:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L561:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L7
	} else {
		goto L562
	}
L562:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_60), int32(0))
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L7
	} else {
		goto L563
	}
L563:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L7
	} else {
		goto L564
	}
L564:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(969), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2332 = m.ExcPending
	if v2332 != 0 {
		goto L7
	} else {
		goto L565
	}
L565:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L566:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L7
	} else {
		goto L567
	}
L567:
	;
	v2340 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+176)) = v2340
	F_errmsg_plural(m, int32(_a_F_ParseFuncOrColumn_1), int32(_a_F_ParseFuncOrColumn_2), v2340, v35+int32(176))
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L7
	} else {
		goto L568
	}
L568:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L7
	} else {
		goto L569
	}
L569:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(679), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L7
	} else {
		goto L570
	}
L570:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_func_match_argtypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v5 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = l2
	v13 = v5
	goto L4
L2:
	;
	v33 = v5
	goto L3
L3:
	;
	return v33
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v19 = F_can_coerce_type(m, l0, l1, v11+int32(32), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v33 = v28
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v23
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v11
	v28 = v13 + int32(1)
	goto L10
L9:
	;
	v28 = v13
	goto L10
L10:
	;
	if v15 != 0 {
		v11 = v15
		v13 = v28
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L5
}
func F_func_parallel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_func_parallel_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_func_parallel_1), int32(2118), int32(_a_F_func_parallel_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29+v30)+102)))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
			}
		}
	}
}
func F_get_func_input_arg_names(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 == int64(0) {
		v121 = v4
		v123 = v4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L38
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L35
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v123
	m.G0 = v13 + int32(16)
	return v121
L4:
	;
	v18 = F_pg_detoast_datum(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 != int32(1) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v25 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v26 != int32(25) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_deconstruct_array_builtin(m, v18, int32(25), v13+int32(8), int32(0), v13+int32(12))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if l1 == int64(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v55 <= int32(0) {
		v121 = v4
		v123 = v4
		goto L3
	} else {
		goto L20
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v55 = v39
	v57 = v4
	goto L11
L13:
	;
	goto L14
L14:
	;
	v41 = F_pg_detoast_datum(m, base.I32_wrap_i64(l1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v43 != int32(1) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v46 != v47 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v50 != int32(18) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v55 = v46
	v57 = v41 + int32(24)
	goto L11
L20:
	;
	v62 = F_palloc(m, v55<<(uint(int32(2))%32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v64 <= int32(0) {
		v121 = v4
		v123 = v62
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v71 = v64
	v73 = int32(0)
	v74 = v4
	goto L23
L23:
	;
	if v57 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v121 = v110
	v123 = v62
	goto L3
L25:
	;
	v113 = v73 + int32(1)
	if v113 < v109 {
		v71 = v109
		v73 = v113
		v74 = v110
		goto L23
	} else {
		goto L34
	}
L26:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v57))))
	v81 = v79 - int32(98)
	if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v81))|base.B2i32(int32(1)<<(uint(v81)%32)&int32(_a_F_get_func_input_arg_names_0) == int32(0)) != 0 {
		v109 = v71
		v110 = v74
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v73<<(uint(int32(3))%32))))
	v100 = F_text_to_cstring(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v103 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v104 = v100
	goto L33
L32:
	;
	v104 = int32(0)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62+v74<<(uint(int32(2))%32)))) = v104
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v109 = v108
	v110 = v74 + int32(1)
	goto L25
L34:
	;
	goto L24
L35:
	;
	F_errmsg_internal(m, int32(_a_F_get_func_input_arg_names_1), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_get_func_input_arg_names_2), int32(1552), int32(_a_F_get_func_input_arg_names_3))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v148
	F_errmsg_internal(m, int32(_a_F_get_func_input_arg_names_4), v13)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_get_func_input_arg_names_2), int32(1562), int32(_a_F_get_func_input_arg_names_3))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_makeFuncCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = F_palloc0(m, int32(40))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l2
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(76)
		return v7
	}
}
